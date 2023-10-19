package generator

import (
	"fmt"
	"go-skeleton/application/services"
	"io/fs"
	"path/filepath"
)

type CodeGenerator struct {
	Logger services.Logger
	c      *Config
}

func NewCodeGenerator(l services.Logger) *CodeGenerator {
	return &CodeGenerator{
		Logger: l,
		c:      GetConfig(l),
	}
}

func GetConfig(l services.Logger) *Config {
	configPath := "tools/generator/config.toml"
	c, err := GetTomlConfig(configPath)
	if err != nil {
		l.Error(err)
	}
	return c
}

func (cg *CodeGenerator) ValidateArgs() {}

func (cg *CodeGenerator) StubHandler(stubs map[string]Stubs) {
	for name, stub := range stubs {
		cg.Logger.Info("Creating: " + name)
		err := filepath.Walk(stub.FromPath, func(path string, info fs.FileInfo, e error) error {
			fmt.Println(path)
			return nil
		})
		if err != nil {
			cg.Logger.Error(err)
		}
	}
}

func (cg *CodeGenerator) Handler(args []string) {
	cg.ValidateArgs()
	fmt.Println(cg.c)

	cg.StubHandler(cg.c.Stubs["crud"])
}
