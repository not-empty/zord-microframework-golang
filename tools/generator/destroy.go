package generator

import (
	"go-skeleton/application/services"
	"os"
)

type CodeDestroy struct {
	Logger     services.Logger
	config     *Config
	domainType string
}

func NewCodeDestroy(l services.Logger, domainType string) *CodeDestroy {
	return &CodeDestroy{
		Logger:     l,
		config:     GetConfig(l),
		domainType: domainType,
	}
}

func (cd *CodeDestroy) Handler(args []string) {
	stubs := GetStubsConfig(cd.Logger, cd.config, cd.domainType)
	replacers := GetReplacersConfig(cd.Logger, cd.config, cd.domainType, args)

	for _, stub := range stubs {
		path := Replacer(stub.ToPath, replacers)
		if !stub.IsDeletable {
			continue
		}
		if err := os.RemoveAll(path); err != nil {
			cd.Logger.Error(err)
		}
	}
}
