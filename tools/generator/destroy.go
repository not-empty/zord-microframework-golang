package generator

import (
	"go-skeleton/pkg/logger"
	"os"
)

type CodeDestroy struct {
	Logger     logger.ILogger
	config     *Config
	domainType string
}

func NewCodeDestroy(l logger.ILogger, domainType string) *CodeDestroy {
	return &CodeDestroy{
		Logger:     l,
		config:     GetConfig(l),
		domainType: domainType,
	}
}

func (cd *CodeDestroy) Handler(args []string) {
	stubs := GetStubsConfig(cd.Logger, cd.config, cd.domainType)
	replacers := GetReplacersConfig(cd.config, cd.domainType, args)

	for _, stub := range stubs {
		path := Replacer(stub.ToPath, replacers)
		if !stub.IsGenerated {
			for _, p := range stub.DeletePatterns {
				deletePattern := Replacer(p, replacers)
				err := RemoveFileLine(path, deletePattern)
				if err != nil {
					cd.Logger.Error(err)
				}
			}
			continue
		}

		if stub.UniqueDelete != "" {
			path = Replacer(stub.UniqueDelete, replacers)
		}

		if err := os.RemoveAll(path); err != nil {
			cd.Logger.Error(err)
		}
	}
}
