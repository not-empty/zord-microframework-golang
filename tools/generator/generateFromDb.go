package generator

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"os"
	"strings"
)

func (cg *CodeGenerator) ReadFromSchema(schema string) {
	file, hclErr := cg.getHclFile(schema)
	if hclErr != nil {
		fmt.Println("Error validating files:", hclErr)
		return
	}
	for _, block := range file.Body().Blocks() {
		err := cg.handleHclBlock(block)
		if err != nil {
			fmt.Println("Error validating files:", err)
			return
		}
	}
}

func (cg *CodeGenerator) getHclFile(schema string) (*hclwrite.File, error) {
	filepath := fmt.Sprintf("tools/migrator/%s.my.hcl", schema)
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
	tableName := block.Labels()[0]
	replacers := cg.generateDomainFromHclBlock(block, tableName)
	validateErr := cg.validateFiles(tableName)
	if validateErr != nil {
		return validateErr
	}
	stubs := GetStubsConfig(cg.Logger, cg.config, cg.domainType)
	cg.GenerateFilesFromStubs(stubs, replacers)
	return nil
}

func (cg *CodeGenerator) generateDomainFromHclBlock(block *hclwrite.Block, tableName string) map[string]string {
	cg.needImportTime = false
	domain := cg.generateDomainStruct(block.Body().Blocks(), tableName)
	replacers := GetReplacersConfig(cg.config, cg.domainType, []string{tableName})
	replacers["{{domainType}}"] = domain
	replacers["{{optionalImports}}"] = ""
	if cg.needImportTime {
		replacers["{{optionalImports}}"] = `"time"`
	}
	return replacers
}

func (cg *CodeGenerator) generateDomainStruct(blocks []*hclwrite.Block, tableName string) string {
	structString := "type " + PascalCase(tableName) + " struct {\n"
	for _, block := range blocks {
		if block.Type() == "column" {
			token, ok := block.Body().Attributes()["type"]
			if !ok {
				continue
			}
			structString = fmt.Sprintf(
				"%s	%s %s `db:\"%s\"`\n",
				structString,
				PascalCase(block.Labels()[0]),
				cg.dbTypesToGoTypes(string(token.Expr().BuildTokens(nil).Bytes())),
				block.Labels()[0],
			)
		}
	}
	structString += "}"
	return structString
}

func (cg *CodeGenerator) dbTypesToGoTypes(typo string) string {
	dbTypesMap := map[string]string{
		" bigint":     "int64",
		" bit":        "bool",
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
		" time":       "time.Time",
		" timestamp":  "time.Time",
		" datetime":   "time.Time",
		" date":       "time.Time",
		" tinyint":    "int8",
		" tinytext":   "string",
		" varbinary":  "string",
		" varchar":    "string",
	}

	GolangType, ok := dbTypesMap[typo]
	if ok {
		return GolangType
	}

	if strings.Contains(typo, "char") {
		return "string"
	}

	if strings.Contains(typo, "time") {
		cg.needImportTime = true
		return "time.Time"
	}
	return typo
}
