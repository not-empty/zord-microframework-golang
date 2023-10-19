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

func GetYamlConfig(filePath string) (map[interface{}]interface{}, error) {
	c := make(map[interface{}]interface{})
	data, err := GetFileData(filePath)
	if err != nil {
		return map[interface{}]interface{}{}, err
	}

	err = yaml.Unmarshal([]byte(data), &c)
	if err != nil {
		return map[interface{}]interface{}{}, err
	}
	return c, nil
}
