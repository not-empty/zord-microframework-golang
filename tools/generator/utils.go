package generator

import (
	"errors"
	"go-skeleton/application/services"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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

func Replacer(str string, replaces map[string]string) string {
	strReplaced := str
	for old, newValue := range replaces {
		strReplaced = strings.ReplaceAll(strReplaced, old, newValue)
	}
	return strReplaced
}

func ProcessFile(fromPath string, toPath string, replacers map[string]string) error {
	data, err := GetFileData(fromPath)
	if err != nil {
		return err
	}

	replData := Replacer(data, replacers)
	replPath := Replacer(toPath, replacers)

	err = os.WriteFile(replPath, []byte(replData), 0755)
	if err != nil {
		return err
	}

	return nil
}

func ProcessFolder(folderPath string, replacers map[string]string) error {
	replPath := Replacer(folderPath, replacers)
	if _, err := os.Stat(replPath); os.IsNotExist(err) {
		err := os.Mkdir(replPath, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func PascalCase(str string) string {
	strCap := strings.ReplaceAll(str, "_", " ")
	strCap = cases.Title(
		language.English,
	).String(
		strCap,
	)
	return strings.ReplaceAll(strCap, " ", "")
}

func CamelCase(str string) string {
	before, after, found := strings.Cut(str, "_")
	if !found {
		return str
	}

	after = strings.ReplaceAll(after, "_", " ")
	after = cases.Title(
		language.English,
	).String(
		after,
	)

	after = strings.ReplaceAll(after, " ", "")
	return before + after
}

func DefineFromToReplaceVariables(vars map[string]string, args []string, replacers map[string]string) map[string]string {
	replaced := map[string]string{}
	for varName, templ := range replacers {
		data, ok := vars[varName]
		if ok {
			replaced[templ] = data
			continue
		}
		replaced[varName] = templ
	}
	return replaced
}

func GetStubsConfig(l services.Logger, c *Config, domainType string) map[string]Stubs {
	stubs, ok := c.Stubs[domainType]
	if !ok {
		l.Error(errors.New("invalid domain type"))
	}
	return stubs
}

func GetReplacersConfig(l services.Logger, c *Config, domainType string, args []string) map[string]string {
	replacers, ok := c.Replacers[domainType]
	if !ok {
		return map[string]string{}
	}

	vars := map[string]string{
		"domain":           args[0],
		"domainPascalCase": PascalCase(args[0]),
		"domainCamelCase":  CamelCase(args[0]),
	}

	replacers = DefineFromToReplaceVariables(vars, args, replacers)

	for r, vl := range replacers {
		replacers[r] = Replacer(vl, replacers)
		if !strings.Contains(vl, "$repeat$") {
			continue
		}
		replacers[r] = Replacer(
			replacers[r],
			map[string]string{
				"$repeat$": r,
			},
		)
	}

	return replacers
}
