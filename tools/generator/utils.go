package generator

import (
	"fmt"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
)

func GetFileData(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func GetTomlConfig(filePath string) (*Config, error) {
	data, err := GetFileData(filePath)
	if err != nil {
		return nil, err
	}
	c := Config{}
	_, err = toml.Decode(data, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func MountFilePath(fromPath string, toFolderPath string, separator string) string {
	return toFolderPath + strings.Split(fromPath, separator+"/")[1]
}

func ProcessFile(path string) error {
	fmt.Println(path)
	return nil
}

func ProcessFolder(folderPath string) error {
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		err := os.Mkdir(folderPath, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Generator) Replacer(str string, replaces map[string]string) string {
	strReplaced := str
	for old, newValue := range replaces {
		strReplaced = strings.ReplaceAll(strReplaced, old, newValue)
	}
	return strReplaced
}
