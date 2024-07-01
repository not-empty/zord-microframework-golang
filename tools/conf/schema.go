package conf

import "fmt"

func (tc *ToolsConfig) MountSchemaHCLFilePath(schema string, driver string) string {
	prefix, ok := tc.SupportedDrivers[driver]
	if !ok {
		panic("invalid driver: " + driver)
	}
	return fmt.Sprintf("%s%s.%s.hcl", tc.SchemasPath, schema, prefix.(string))
}
