package generator

import (
	"fmt"
	"os"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func (cg *CodeGenerator) ReadFromSchema(schema string, table string) {
	file, hclErr := cg.getHclFile(schema)
	if hclErr != nil {
		fmt.Println("Error validating files:", hclErr)
		return
	}
	for _, block := range file.Body().Blocks() {
		if table != "" && block.Labels()[0] != table {
			continue
		}
		err := cg.handleHclBlock(block)
		if err != nil {
			fmt.Println("Error validating files:", err)
			return
		}
	}
}

func (cg *CodeGenerator) getHclFile(schema string) (*hclwrite.File, error) {
	filepath := fmt.Sprintf("schemas/%s.my.hcl", schema)
	content, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}
	file, _ := hclwrite.ParseConfig(content, filepath, hcl.Pos{Line: 1, Column: 1})
	return file, err
}

func (cg *CodeGenerator) handleHclBlock(block *hclwrite.Block) error {
	if block.Type() == "schema" {
		return nil
	}
	if len(block.Labels()) == 0 {
		return nil
	}

	tableName := CamelCase(block.Labels()[0])
	rawTableName := block.Labels()[0]

	replacers := cg.generateDomainFromHclBlock(block, tableName, rawTableName)
	validateErr := cg.validateFiles(tableName)
	if validateErr != nil {
		return validateErr
	}
	stubs := GetStubsConfig(cg.Logger, cg.config, cg.domainType)
	cg.GenerateFilesFromStubs(stubs, replacers)
	return nil
}

func (cg *CodeGenerator) generateDomainFromHclBlock(block *hclwrite.Block, tableName string, rawTableName string) map[string]string {
	cg.needImportTime = new(bool)
	cg.primaryKey = new(string)
	cg.pkType = new(string)
	cg.isIntId = new(bool)
	*cg.needImportTime = false
	*cg.isIntId = false

	dbPk, structPk := cg.findPkOnBlocks(block.Body().Blocks())
	*cg.primaryKey = structPk

	domain := cg.generateDomainStruct(block.Body().Blocks(), tableName, cg.primaryKey, cg.pkType)

	// For GET Data struct, set pk and pkType so {{pkType}} is set
	getPk := dbPk
	getPkType := ""
	getDataType := cg.generateStruct(block.Body().Blocks(), &getPk, &getPkType, cg.generateDeclarationLine)

	createAttrData := cg.generateStruct(block.Body().Blocks(), &dbPk, nil, cg.generateAttributionLine)
	editAttrData := cg.generateStruct(block.Body().Blocks(), &dbPk, nil, cg.generateAttributionLine)
	replacers := GetReplacersConfig(cg.config, cg.domainType, []string{tableName, rawTableName})
	replacers["{{domainType}}"] = domain
	replacers["{{dataType}}"] = getDataType
	replacers["{{pkDbName}}"] = dbPk
	replacers["{{pkName}}"] = structPk
	replacers["{{pkType}}"] = getPkType
	replacers["{{createServiceData}}"] = createAttrData
	replacers["{{editServiceData}}"] = editAttrData
	if *cg.needImportTime {
		replacers["{{optionalImports}}"] = "\"time\""
	}
	replacers["{{idVar}}"] = "domain." + structPk + " = s.idCreator.Create()"
	return replacers
}

func (cg *CodeGenerator) generateDomainStruct(blocks []*hclwrite.Block, tableName string, pk, pkType *string) string {
	_, structPk := cg.findPkOnBlocks(blocks)
	*pk = structPk
	structString := "type " + PascalCase(tableName) + " struct {\n"
	structString += cg.generateStruct(blocks, pk, pkType, cg.generateDeclarationLine)
	structString += "\tclient string\n\tfilters *filters.Filters\n"
	structString += "}"
	return structString
}

func (cg *CodeGenerator) generateStruct(blocks []*hclwrite.Block, pk, pkType *string, strFormationFunc func(string, string, string, string) string) string {
	declarationString := ""
	for _, block := range blocks {
		if block.Type() == "column" {
			token, ok := block.Body().Attributes()["type"]
			if !ok {
				continue
			}
			tokenStr := string(token.Expr().BuildTokens(nil).Bytes())
			goType := cg.dbTypesToGoTypes(tokenStr)
			nullable, nullOk := block.Body().Attributes()["null"]
			isNullable := cg.verifyIsNullable(nullable, nullOk)
			if isNullable {
				goType = "*" + goType
			}

			if pk != nil && block.Labels()[0] == *pk {
				if pkType != nil {
					*pkType = fmt.Sprintf("%s string `param:\"id\"`\n", PascalCase(*pk))
				}
				continue
			}

			declarationString = strFormationFunc(
				declarationString,
				PascalCase(block.Labels()[0]),
				goType,
				block.Labels()[0],
			)
		}
	}
	return declarationString
}

func (cg *CodeGenerator) generateDeclarationLine(str, name, goType, dbTag string) string {
	if name == PascalCase(*cg.primaryKey) && strings.Contains(goType, "int") {
		return fmt.Sprintf(
			"%s	%s %s `db:\"%s\"`\n",
			str,
			name,
			"*string",
			dbTag,
		)
	}
	return fmt.Sprintf(
		"%s	%s %s `db:\"%s\"`\n",
		str,
		name,
		goType,
		dbTag,
	)
}

func (cg *CodeGenerator) generateAttributionLine(str, name, _, _ string) string {
	if name == PascalCase(*cg.primaryKey) {
		return str
	}
	return fmt.Sprintf(
		"%s\tdomain.%s = data.%s\n",
		str,
		name,
		name,
	)
}

// Utility function to strip 'column.' prefix
func stripColumnPrefix(col string) string {
	if strings.HasPrefix(col, "column.") {
		return strings.TrimPrefix(col, "column.")
	}
	return col
}

// Update findPkOnBlocks to use a sliding window over the tokens
func (cg *CodeGenerator) findPkOnBlocks(blocks []*hclwrite.Block) (string, string) {
	for _, block := range blocks {
		if block.Type() == "primary_key" {
			attr, ok := block.Body().Attributes()["columns"]
			if !ok {
				return "", ""
			}
			expr := attr.Expr()
			traversal := expr.BuildTokens(nil)
			var columns []string
			for i := 0; i < len(traversal)-2; i++ {
				if traversal[i].Type.String() == "TokenIdent" && string(traversal[i].Bytes) == "column" &&
					traversal[i+1].Type.String() == "TokenDot" &&
					traversal[i+2].Type.String() == "TokenIdent" {
					columns = append(columns, string(traversal[i+2].Bytes))
				}
			}
			if len(columns) > 0 {
				col := columns[0]
				return col, PascalCase(col)
			}
		}
	}
	return "", ""
}

func (cg *CodeGenerator) dbTypesToGoTypes(typo string) string {
	dbTypesMap := map[string]string{
		" bigint":     "int64",
		" bit":        " ",
		" char":       "string",
		" decimal":    "float64",
		" float":      "float32",
		" double":     "float64",
		" int":        "int",
		" longtext":   "string",
		" mediumint":  "int",
		" mediumtext": "string",
		" smallint":   "int16",
		" text":       "string",
		" time":       "string",
		" timestamp":  "string",
		" datetime":   "time.Time",
		" date":       "string",
		" tinyint":    "int8",
		" tinytext":   "string",
		" varbinary":  "string",
		" varchar":    "string",
		" json":       "string",
	}

	GolangType, ok := dbTypesMap[typo]
	if ok {
		if GolangType == "time.Time" {
			*cg.needImportTime = true
		}
		return GolangType
	}

	if strings.Contains(typo, "char") {
		return "string"
	}

	if strings.Contains(typo, "double") {
		return "float64"
	}

	if strings.Contains(typo, "year") {
		return "string"
	}

	if strings.Contains(typo, "decimal") {
		return "float64"
	}

	return typo
}

func (cg *CodeGenerator) verifyIsNullable(token *hclwrite.Attribute, ok bool) bool {
	if !ok {
		return false
	}
	value := token.Expr().BuildTokens(nil).Bytes()
	return string(value) == "true"
}
