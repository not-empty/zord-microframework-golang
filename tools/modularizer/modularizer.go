package modularizer

import (
	"errors"
	git "github.com/go-git/go-git/v5"
	"go-skeleton/pkg/logger"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type Modularizer struct {
	ModulesList map[string]string
	Logger      *logger.Logger
}

func NewModularizer(l *logger.Logger) *Modularizer {
	modulesList := map[string]string{
		"grpc": "https://github.com/not-empty/zord-microframework-golang.git",
	}
	return &Modularizer{
		ModulesList: modulesList,
		Logger:      l,
	}
}

func (m Modularizer) GetModule(module string) error {
	module, exists := m.ModulesList[module]
	if !exists {
		err := errors.New("modulo inexistente ou não mapeado")
		m.Logger.Error(err)
		return err
	}
	_, err := git.PlainClone("tools/modularizer/temp", false, &git.CloneOptions{
		URL: module,
	})
	if err != nil {
		m.Logger.Error(err)
		return err
	}

	filepath.Walk("tools/modularizer/temp", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			m.createFolder(path)
			return nil
		}
		return nil
	})

	if err := os.RemoveAll("tools/modularizer/temp"); err != nil {
		m.Logger.Error(err)
		return err
	}
	return nil
}

func (m *Modularizer) replacer(str string, replaces map[string]string) string {
	strReplaced := str
	for old, newCode := range replaces {
		strReplaced = strings.ReplaceAll(strReplaced, old, newCode)
	}
	return strReplaced
}

func (m *Modularizer) createFolder(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	}
}
