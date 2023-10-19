package generator

import (
	"fmt"
	"go-skeleton/application/services"
)

type Config struct {
	Replacers map[string]string `yaml:"replacers"`
	Stubs     map[string]struct {
		ToPath   string `yaml:"toPath"`
		FromPath string `yaml:"fromPath"`
	} `yaml:"stubs"`
}

type CodeGenerator struct {
	Logger services.Logger
	c      Config
}

func NewCodeGenerator(l services.Logger) *CodeGenerator {
	return &CodeGenerator{
		Logger: l,
		c:      GetConfig(l),
	}
}

func GetConfig(l services.Logger) Config {
	c, err := GetYamlConfig("tools/generator/config.yaml")
	if err != nil {
		l.Error(err)
	}
	return c
}

func (g *CodeGenerator) Handler(args []string) {
	for name, stub := range g.c.Stubs {
		g.Logger.Info("Creating: " + name)
		fmt.Println(stub.FromPath)
	}
}
