package generator

import (
	"os"

	"gopkg.in/yaml.v3"
)

func GetFileData(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func GetYamlConfig(filePath string) (Config, error) {
	data, err := GetFileData(filePath)
	if err != nil {
		return Config{}, err
	}
	c := Config{}
	err = yaml.Unmarshal([]byte(data), &c)
	if err != nil {
		return Config{}, err
	}
	return c, nil
}
