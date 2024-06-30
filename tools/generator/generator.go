package generator

import (
	"errors"
	"go-skeleton/pkg/logger"
	"io/fs"
	"os"
	"path/filepath"
)

var (
	configPath = "tools/generator/config.toml"
)

type CodeGenerator struct {
	Logger         *logger.Logger
	config         *Config
	validator      bool
	domainType     string
	needImportTime *bool
	primaryKey     *string
	pkType         *string
	isIntId        *bool
}

func NewCodeGenerator(logger *logger.Logger, validator bool, domainType string) *CodeGenerator {
	return &CodeGenerator{
		Logger:     logger,
		config:     GetConfig(logger),
		validator:  validator,
		domainType: domainType,
	}
}

func GetConfig(l *logger.Logger) *Config {
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
	replacers := GetReplacersConfig(cg.config, cg.domainType, args)
	domain := args[0]
	err := cg.validateFiles(domain)
	if err != nil {
		cg.Logger.Error(err)
		return
	}
	cg.GenerateFilesFromStubs(stubs, replacers)
}

func (cg *CodeGenerator) validateFiles(domain string) error {
	if FileExists("application/domain/"+domain+"/"+domain+".go") && cg.domainType == "crud" {
		return errors.New("domain already exists")
	}

	if FileExists("application/services/"+domain+"service.go") && cg.domainType == "usecase" {
		return errors.New("service already exists")
	}
	return nil
}

func (cg *CodeGenerator) GenerateFilesFromStubs(stubs map[string]Stubs, replacers map[string]string) {
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
