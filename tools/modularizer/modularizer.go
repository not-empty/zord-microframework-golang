package modularizer

import (
	"encoding/json"
	"errors"
	"github.com/go-git/go-git/v5"
	"go-skeleton/pkg/logger"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	replacer       map[string]string
	filesToReplace []string
}

type Modularizer struct {
	Logger *logger.Logger
	Config *Config
}

var ModulesList = map[string]string{
	"grpc": "https://github.com/not-empty/zord-grpc-microframework-golang.git",
}
var IgnoreList []string
var tempFolder = "tools/modularizer/temp"

func NewModularizer(l *logger.Logger) *Modularizer {
	IgnoreList = append(
		IgnoreList,
		".git",
		"application",
		"config.json",
		"kernel",
	)
	return &Modularizer{
		Logger: l,
	}
}

func (m *Modularizer) GetModule(module string) error {
	module, exists := ModulesList[module]
	if !exists {
		err := errors.New("modulo inexistente ou não mapeado")
		m.Logger.Error(err)
		return err
	}
	_, err := git.PlainClone(tempFolder, false, &git.CloneOptions{
		URL: module,
	})
	if err != nil {
		m.Logger.Error(err)
		return err
	}

	jsonBytes, err := os.ReadFile(tempFolder + "/config.json")

	if err != nil {
		m.Logger.Info("Missing configuration")
		m.Logger.Error(err)
		return err
	}

	err = json.Unmarshal(jsonBytes, &m.Config)

	if err != nil {
		m.Logger.Info("Invalid configuration file")
		m.Logger.Error(err)
		return err
	}

	err = filepath.Walk(tempFolder, func(path string, info fs.FileInfo, err error) error {
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

	if err != nil {
		m.Logger.Error(err)
		return err
	}

	for _, file := range m.Config.filesToReplace {
		replaced := m.replacer(tempFolder+m.getFileData(file), m.Config.replacer)
		err = os.WriteFile(file, []byte(replaced), 0755)
		if err != nil {
			m.Logger.Error(err, "modularizer")
		}
	}

	if err := os.RemoveAll(tempFolder); err != nil {
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
