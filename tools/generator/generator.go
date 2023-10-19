package generator

import (
	"errors"
	"go-skeleton/application/services"
	"io/fs"
	"path/filepath"
)

var (
	configPath = "tools/generator/config.toml"
)

type CodeGenerator struct {
	Logger     services.Logger
	config     *Config
	service    string
	validator  bool
	domainType string
}

func NewCodeGenerator(logger services.Logger, service string, validator bool, domainType string) *CodeGenerator {
	return &CodeGenerator{
		Logger:     logger,
		config:     GetConfig(logger),
		service:    service,
		validator:  validator,
		domainType: domainType,
	}
}

func GetConfig(l services.Logger) *Config {
	c, err := GetTomlConfig(configPath)
	if err != nil {
		l.Error(err)
	}
	return c
}

func (cg *CodeGenerator) WalkProcess(name string, stub Stubs) {
	filepath.Walk(stub.FromPath, func(path string, info fs.FileInfo, e error) error {
		if info.IsDir() {
			err := ProcessFolder(stub.ToPath + info.Name())
			if err != nil {
				cg.Logger.Error(err)
			}
			return nil
		}
		err := ProcessFile(MountFilePath(path, stub.ToPath, name))
		if err != nil {
			cg.Logger.Error(err)
		}
		return nil
	})
}

func (cg *CodeGenerator) Handler() {
	stubs, ok := cg.config.Stubs[cg.domainType]
	if !ok {
		cg.Logger.Error(errors.New("invalid domain type"))
	}
	for name, stub := range stubs {
		cg.WalkProcess(name, stub)
	}
}
