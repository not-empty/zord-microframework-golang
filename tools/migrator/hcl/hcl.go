package hcl

import (
	"strings"

	"github.com/zclconf/go-cty/cty"

	"github.com/hashicorp/hcl/v2/hclwrite"
)

type AdditionalOpts struct {
	Alias      string
	Name       string
	Attributes map[string]string
}

type HCL struct{}

func NewHCL() *HCL {
	return &HCL{}
}

func (hcl *HCL) CreateAtlasHCLFromDomain(domain string, schema string, fields []string, dbDefinitions []string) string {
	f := hclwrite.NewEmptyFile()
	tableBody := hcl.AddBlock(f, "table", []string{domain})
	tableBody.Body().SetAttributeValue("schema", cty.StringVal("schema."+schema))
	hcl.SetFields(tableBody.Body(), fields, dbDefinitions)
	return hcl.RemoveQuotation(string(f.Bytes()))
}

func (hcl *HCL) AddBlock(f *hclwrite.File, name string, labels []string) *hclwrite.Block {
	return f.Body().AppendNewBlock(name, labels)
}

func (hcl *HCL) SetFields(body *hclwrite.Body, fields []string, dbDefinitions []string) {
	for i, name := range fields {
		metadata := strings.Split(dbDefinitions[i], ",")
		hcl.SetField(body, name, metadata)
	}
}

func (hcl *HCL) SetField(body *hclwrite.Body, fieldName string, metadata []string) {
	f := body.AppendNewBlock("column", []string{fieldName})
	for _, vl := range metadata {
		split := strings.Split(vl, "=")
		if len(split) > 1 {
			f.Body().SetAttributeValue(split[0], cty.StringVal(split[1]))
			continue
		}
		hcl.SetAdditionalAttr(body, fieldName, split[0])
	}
}

func (hcl *HCL) RemoveQuotation(hclStr string) string {
	noQuotationsAttr := hcl.getNoQuotationsAttributes()
	lines := strings.Split(hclStr, "\n")
	for i, l := range lines {
		newLine := l
		for _, nqOpt := range noQuotationsAttr {
			if strings.Contains(l, nqOpt) {
				newLine = strings.ReplaceAll(l, "\"", "")
			}
		}
		lines[i] = newLine
	}
	return strings.Join(lines, "\n")
}

func (hcl *HCL) SetAdditionalAttr(body *hclwrite.Body, fieldName string, opt string) {
	additionalOpts := hcl.getAdditionalAttributes(fieldName)
	for _, field := range additionalOpts {
		if opt == field.Alias {
			block := body.AppendNewBlock(field.Name, []string{})
			if field.Attributes != nil {
				for name, vl := range field.Attributes {
					block.Body().SetAttributeValue(name, cty.StringVal(vl))
				}
			}
		}
	}
}

func (hcl *HCL) getNoQuotationsAttributes() []string {
	return []string{
		"type =",
		"null =",
		"columns =",
		"schema =",
	}
}

func (hcl *HCL) getAdditionalAttributes(fielName string) []AdditionalOpts {
	return []AdditionalOpts{
		{
			Alias: "PK",
			Name:  "primary_key",
			Attributes: map[string]string{
				"columns": "[column." + fielName + "]",
			},
		},
	}
}
