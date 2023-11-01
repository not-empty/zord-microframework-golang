package generator

import (
	"errors"
	"go-skeleton/internal/application/services"
	"io/fs"
	"os"
	"path/filepath"
)

var (
	configPath = "tools/generator/config.toml"
)

type CodeGenerator struct {
	Logger     services.Logger
	config     *Config
	validator  bool
	domainType string
}

func NewCodeGenerator(logger services.Logger, validator bool, domainType string) *CodeGenerator {
	return &CodeGenerator{
		Logger:     logger,
		config:     GetConfig(logger),
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

func (cg *CodeGenerator) WalkProcess(name string, stub Stubs, replacers map[string]string) {
	filepath.Walk(stub.FromPath, func(path string, info fs.FileInfo, e error) error {
		if name == info.Name() {
			return nil
		}
		if info.IsDir() {
			err := ProcessFolder(stub.ToPath+info.Name(), replacers)
			if err != nil {
				cg.Logger.Error(err)
			}
			return nil
		}
		err := ProcessFile(path, MountFilePath(path, stub.ToPath, name), replacers)
		if err != nil {
			cg.Logger.Error(err)
		}

		return nil
	})
}

func (cg *CodeGenerator) Handler(args []string) {
	stubs := GetStubsConfig(cg.Logger, cg.config, cg.domainType)
	replacers := GetReplacersConfig(cg.Logger, cg.config, cg.domainType, args)

	domain := args[0]
	if FileExists("application/domain/" + domain + "/" + domain + ".go") {
		cg.Logger.Error(errors.New("domain already exists"))
		return
	}

	for name, stub := range stubs {
		if !stub.IsGenerated {
			data, err := GetFileData(stub.ToPath)
			if err != nil {
				cg.Logger.Error(err)
			}

			replData := Replacer(data, replacers)
			err = os.WriteFile(stub.ToPath, []byte(replData), 0755)
			if err != nil {
				cg.Logger.Error(err)
			}

			continue
		}

		err := ProcessFolder(stub.ToPath, replacers)
		if err != nil {
			cg.Logger.Error(err)
		}
		cg.WalkProcess(name, stub, replacers)
	}
}
