package generator

import (
	"os"

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

func ProcessFile(path string) {}

func ProcessFolder(path string) {}
