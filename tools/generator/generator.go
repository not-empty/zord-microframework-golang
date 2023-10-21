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

func (cg *CodeGenerator) DefineFromToReplaceVariables(args []string, replacers map[string]string) map[string]string {
	vars := map[string]string{
		"domain":           args[0],
		"domainPascalCase": PascalCase(args[0]),
		"domainCamelCase":  CamelCase(args[0]),
	}

	replaced := map[string]string{}
	for varName, templ := range replacers {
		data, ok := vars[varName]
		if ok {
			replaced[templ] = data
			continue
		}
		replaced[varName] = templ
	}

	return replaced
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
	stubs, ok := cg.config.Stubs[cg.domainType]
	if !ok {
		cg.Logger.Error(errors.New("invalid domain type"))
	}

	replacers, ok := cg.config.Replacers[cg.domainType]
	if ok {
		replacers = cg.DefineFromToReplaceVariables(args, replacers)
	} else {
		replacers = map[string]string{}
	}

	for name, stub := range stubs {
		err := ProcessFolder(stub.ToPath, replacers)
		if err != nil {
			cg.Logger.Error(err)
		}
		cg.WalkProcess(name, stub, replacers)
	}
}
