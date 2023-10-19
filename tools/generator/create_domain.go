package generator

import (
	"errors"
	"fmt"
	"go-skeleton/application/services"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	domainTo       = "application/domain/"
	servicesTo     = "application/services/"
	repositoriesTo = "pkg/repositories/"
	routesTo       = "cmd/handlers/http/routes"
	declarableDir  = routesTo + "/declarable.go"
	migratorTo     = "tools/migrator/migrate.go"
)

type Generator struct {
	Logger services.Logger
}

func NewGenerator(l services.Logger) *Generator {
	return &Generator{
		Logger: l,
	}
}

func (g *Generator) CreateDomain(domain string, validator bool, service string) error {
	if g.isInvalidDomain(domain) {
		panic(errors.New("unable create a new domain, please check domain name"))
	}
	domainCap := g.pascalCase(domain)

	targetPaths := map[string]string{
		"domain":       domainTo + domain,
		"services":     servicesTo + domain,
		"repositories": repositoriesTo + domain,
		"routes":       routesTo,
	}
	// debug
	// g.createFolders(targetPaths)

	stubs := "tools/generator/stubs/"

	declarableTemplate := fmt.Sprintf(
		"\"%s\": %sListRoutes,\n		//{{codeGen2}}",
		domain,
		domain,
	)
	validatorRuleTemplate :=
		fmt.Sprintf(
			`errs := r.validator.ValidateStruct(r.%s)
			for _, err := range errs {
				if err != nil {
					return err
				}
			}`,
			domainCap,
		)

	useCustomService := false
	customService := ""
	customServiceCap := ""
	if service != "" {
		useCustomService = true
		customService = service
		customServiceCap = g.pascalCase(customService)
		//{{customService}}
		//{{customServiceCap}}
	}

	fmt.Println(customService)
	fmt.Println(customServiceCap)

	err := filepath.Walk(stubs, func(path string, info fs.FileInfo, e error) error {
		if useCustomService && !strings.Contains(path, "custom") {
			return nil
		}

		if info.IsDir() {
			folders := strings.Split(path, "/__")
			if useCustomService {
				folders = []string{
					strings.Split(path, "/custom")[0],
					strings.ToUpper(service),
				}
				fmt.Println(folders)
				fullPath := g.getFullFolderPath(targetPaths, folders[0], folders[1], "/stubs/")
				fmt.Println(fullPath)
				return nil
			}

			if len(folders) > 1 {
				fullPath := g.getFullFolderPath(targetPaths, folders[0], folders[1], "/stubs/")
				if _, err := os.Stat(fullPath); os.IsNotExist(err) {
					err = os.Mkdir(fullPath, 0755)
					if err != nil {
						return err
					}
				}
				g.Logger.Info("CREATE FOLDER: " + fullPath)
			}
			return nil
		}

		fullFilePath := g.getFullFilePath(targetPaths, path, "/stubs/")
		fullFilePath = g.replacer(
			fullFilePath,
			map[string]string{
				"{{domain}}": domain,
				".stub":      "",
				"__":         "",
			},
		)

		changeMap := map[string]string{
			"{{domain}}":          domain,
			"{{domainCap}}":       domainCap,
			"{{validatorImport}}": "\n	\"go-skeleton/pkg/validator\"",
			"{{validator}}":       "\nvalidator *validator.Validator",
			"{{,validator}}":      ", validator *validator.Validator",
			"{{validatorRule}}":   validatorRuleTemplate,
			"{{validatorInject}}": "validator: validator,",
			"{{hsValidator}}":     "hs.validator",
			"{{,hsValidator}}":    ", hs.validator",
		}

		if !validator {
			changeMap = map[string]string{
				"{{domain}}":          domain,
				"{{domainCap}}":       domainCap,
				"{{validatorImport}}": "",
				"{{validator}}":       "",
				"{{,validator}}":      "",
				"{{validatorRule}}":   "",
				"{{validatorInject}}": "",
				"{{hsValidator}}":     "",
				"{{,hsValidator}}":    "",
			}
		}

		g.Logger.Info(fullFilePath)
		err := g.createFile(
			domain,
			domainCap,
			path,
			fullFilePath,
			changeMap,
		)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	oldDecData := g.getFileData(declarableDir)
	if strings.Contains(oldDecData, domain) {
		return nil
	}

	routesString := "%sListRoutes := New%sRoutes(logger, Environment, MySql, idCreator)\n	//{{codeGen1}}"
	if validator {
		routesString = "%sListRoutes := New%sRoutes(logger, Environment, MySql, idCreator, validator)\n	//{{codeGen1}}"
	}

	routesInstanceTemplate := fmt.Sprintf(
		routesString,
		domain,
		domainCap,
	)

	newData := g.replacer(
		oldDecData,
		map[string]string{
			"//{{codeGen1}}":      routesInstanceTemplate,
			"//{{codeGen2}}":      declarableTemplate,
			"{{validatorImport}}": "\ngo-skeleton/pkg/validator",
			"{{validator}}":       "\nvalidator *validator.Validator",
			"{{,validator}}":      ", validator *validator.Validator",
			"{{validatorRule}}":   validatorRuleTemplate,
			"{{validatorInject}}": "validator: validator,",
			"{{hsValidator}}":     ", hs.validator",
		},
	)

	g.Logger.Info("ADD ROUTES: " + declarableDir)
	err = os.WriteFile(declarableDir, []byte(newData), 0755)
	if err != nil {
		g.Logger.Error(err)
		return err
	}

	oldMigData := g.getFileData(migratorTo)
	if strings.Contains(string(oldMigData), domain) {
		return nil
	}

	importTemplate := fmt.Sprintf("\"go-skeleton/application/domain/%s\"\n	//{{codeGen1}}", domain)
	createTableTemplate := fmt.Sprintf("m.db.Db.Migrator().CreateTable(&%s.%s{})\n	//{{codeGen2}}", domain, domainCap)

	newMigData := g.replacer(
		oldMigData,
		map[string]string{
			"//{{codeGen1}}": importTemplate,
			"//{{codeGen2}}": createTableTemplate,
		},
	)

	g.Logger.Info("ADD MIGRATOR: " + migratorTo)
	err = os.WriteFile(migratorTo, []byte(newMigData), 0755)
	if err != nil {
		g.Logger.Error(err)
		return err
	}

	return nil
}

func (g *Generator) removeFileLine(path string, search string) error {
	data := g.getFileData(path)
	lines := strings.Split(string(data), "\n")
	for i, l := range lines {
		if strings.Contains(string(l), search) {
			lines = append(lines[:i], lines[i+1:]...)
		}
	}
	err := os.WriteFile(path, []byte(strings.Join(lines, "\n")), 0755)
	if err != nil {
		g.Logger.Error(err)
		return err
	}
	return nil
}

func (g *Generator) DestroyDomain(domain string) error {
	if g.isInvalidDomain(domain) {
		panic(errors.New("Unable delete a domain, please check domain name for delete"))
	}

	err := g.removeFileLine(migratorTo, domain)
	if err != nil {
		return err
	}

	err = g.removeFileLine(declarableDir, domain)
	if err != nil {
		return err
	}

	destroyPaths := []string{
		routesTo + "/" + domain + ".go",
		repositoriesTo + domain,
		servicesTo + domain,
		domainTo + domain,
	}

	for _, p := range destroyPaths {
		if err := os.RemoveAll(p); err != nil {
			g.Logger.Error(err)
			return err
		}
	}

	return nil
}

func (g *Generator) pascalCase(str string) string {
	strCap := strings.ReplaceAll(str, "_", " ")
	strCap = cases.Title(
		language.English,
	).String(
		strCap,
	)
	return strings.ReplaceAll(strCap, " ", "")
}

func (g *Generator) getFileData(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		g.Logger.Error(err)
	}
	return string(data)
}

func (g *Generator) replacer(str string, replaces map[string]string) string {
	strReplaced := str
	for old, newValue := range replaces {
		strReplaced = strings.ReplaceAll(strReplaced, old, newValue)
	}
	return strReplaced
}

func (g *Generator) createFile(domain string, domainCap string, from string, to string, replacers map[string]string) error {
	data := g.getFileData(from)
	domainContent := g.replacer(data, replacers)

	err := os.WriteFile(to, []byte(domainContent), 0755)
	if err != nil {
		g.Logger.Error(err)
		return err
	}
	g.Logger.Info("CREATE FILE: " + to)
	return nil
}

func (g *Generator) createFolders(folders map[string]string) {
	for _, path := range folders {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			err := os.Mkdir(path, 0755)
			if err != nil {
				return
			}
		}
	}
}

func (g *Generator) getFullFolderPath(targetPaths map[string]string, base string, folder string, root string) string {
	rootLv := strings.Split(base, root)[1]
	pathTo := targetPaths[rootLv]
	fullPath := fmt.Sprintf("%s/%s", pathTo, folder)
	return fullPath
}

func (g *Generator) getFullFilePath(targetPath map[string]string, path string, root string) string {
	dir := strings.Split(strings.Split(path, root)[1], "/")[0]
	target := targetPath[dir]
	file := strings.Split(path, "/"+dir+"/")[1]
	return fmt.Sprintf("%s/%s", target, file)
}

func (g *Generator) isInvalidDomain(domain string) bool {
	if domain == "" {
		return true
	}
	return false
}
