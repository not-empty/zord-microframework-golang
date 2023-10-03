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
	Logger *logger.Logger
}

var ModulesList = map[string]string{
	"grpc": "https://github.com/not-empty/zord-grpc-microframework-golang.git",
}
var IgnoreList []string

func NewModularizer(l *logger.Logger) *Modularizer {
	IgnoreList = append(IgnoreList, ".git", "application")
	return &Modularizer{
		Logger: l,
	}
}

func (m Modularizer) GetModule(module string) error {
	module, exists := ModulesList[module]
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
		if m.isInIgnore(info.Name()) {
			return filepath.SkipDir
		}
		if info.IsDir() {
			name := m.getFullFilePath(path)
			if name == "" {
				return nil
			}
			m.createFolder(name)
			m.Logger.Info(path)
			return nil
		}
		err = m.createFile(path, m.getFullFilePath(path), make(map[string]string))
		if err != nil {
			m.Logger.Error(err)
			return err
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
		err = os.Mkdir(path, 0755)
		if err != nil {
			m.Logger.Error(err)
		}
	}
}

func (m *Modularizer) createFile(from string, to string, replacers map[string]string) error {
	data := m.getFileData(from)
	domainContent := m.replacer(data, replacers)

	err := os.WriteFile(to, []byte(domainContent), 0755)
	if err != nil {
		m.Logger.Error(err)
		return err
	}
	m.Logger.Info("CREATE FILE: " + to)
	return nil
}

func (m *Modularizer) getFileData(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		m.Logger.Error(err)
	}
	return string(data)
}

func (m *Modularizer) getFullFilePath(path string) string {
	pathArray := strings.Split(path, "/temp/")
	if len(pathArray) > 1 {
		return pathArray[1]
	}
	return ""
}

func (m *Modularizer) isInIgnore(file string) bool {
	for _, ignore := range IgnoreList {
		if file == ignore {

			return true
		}
	}
	return false
}
