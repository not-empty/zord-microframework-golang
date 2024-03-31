package generator

import (
	"go-skeleton/pkg/logger"
	"os"
	"regexp"
)

type CodeDestroy struct {
	Logger     *logger.Logger
	config     *Config
	domainType string
}

func NewCodeDestroy(l *logger.Logger, domainType string) *CodeDestroy {
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
			if stub.DeleteRegex != "" {
				replaced := Replacer(stub.DeleteRegex, replacers)
				pattern := regexp.MustCompile(replaced)
				err := RemoveFromRegex(path, pattern)
				if err != nil {
					cd.Logger.Error(err)
				}
			}

			for _, p := range stub.DeleteLinePatterns {
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
