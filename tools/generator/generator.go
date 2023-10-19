package generator

import (
	"fmt"
	"go-skeleton/application/services"
)

// type Config struct {
// 	replacers map[string]string
// 	stubs     map[string]struct {
// 		toPath   string
// 		fromPath string
// 	}
// }

type CodeGenerator struct {
	Logger services.Logger
	c      map[interface{}]interface{}
}

func NewCodeGenerator(l services.Logger) *CodeGenerator {
	return &CodeGenerator{
		Logger: l,
		c:      GetConfig(l),
	}
}

func GetConfig(l services.Logger) map[interface{}]interface{} {
	c, err := GetYamlConfig("tools/generator/config.yaml")
	if err != nil {
		l.Error(err)
	}
	return c
}

func (g *CodeGenerator) Handler(args []string) {
	fmt.Println("Handler.....")
	fmt.Println(g.c["replacers"])
}
