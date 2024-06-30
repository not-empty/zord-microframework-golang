package conf

import (
	"github.com/spf13/viper"
)

type ToolsConfig struct {
	Viper            *viper.Viper
	DomainsPath      string
	SchemasPath      string
	SupportedDrivers map[string]interface{}
}

func NewToolsConfig() *ToolsConfig {
	config := ToolsConfig{}
	config.initConfig()
	return &config
}

func (tc *ToolsConfig) initConfig() {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath("tools/conf")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	tc.DomainsPath = v.GetString("domainsPath")
	tc.SchemasPath = v.GetString("schemasPath")
	tc.SupportedDrivers = v.Get("supportedDrivers").(map[string]interface{})
	tc.Viper = v
}
