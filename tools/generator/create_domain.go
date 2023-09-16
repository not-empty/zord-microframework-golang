package generator

import (
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

func (g *Generator) CreateDomain(domain string) error {
	domainCap := g.pascalCase(domain)

	targetPaths := map[string]string{
		"domain":       domainTo + domain,
		"services":     servicesTo + domain,
		"repositories": repositoriesTo + domain,
		"routes":       routesTo,
	}
	g.createFolders(targetPaths)

	stubs := "tools/generator/stubs/"

	filepath.Walk(stubs, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			folders := strings.Split(path, "/__")
			if len(folders) > 1 {
				fullPath := g.getFullFolderPath(targetPaths, folders[0], folders[1], "/stubs/")
				if _, err := os.Stat(fullPath); os.IsNotExist(err) {
					os.Mkdir(fullPath, 0755)
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
		g.Logger.Info(fullFilePath)
		g.createFile(
			domain,
			domainCap,
			path,
			fullFilePath,
			map[string]string{
				"{{domain}}":    domain,
				"{{domainCap}}": domainCap,
			},
		)
		return nil
	})

	oldDecData := g.getFileData(declarableDir)
	if strings.Contains(oldDecData, domain) {
		return nil
	}

	routesInstanceTemplate := fmt.Sprintf(
		"%sListRoutes := New%sRoutes(logger, Environment, MySql)\n	//{{codeGen1}}",
		domain,
		domainCap,
	)

	declarableTemplate := fmt.Sprintf(
		"\"%s\": %sListRoutes,\n		//{{codeGen2}}",
		domain,
		domain,
	)

	newData := g.replacer(
		oldDecData,
		map[string]string{
			"//{{codeGen1}}": routesInstanceTemplate,
			"//{{codeGen2}}": declarableTemplate,
		},
	)

	g.Logger.Info("ADD ROUTES: " + declarableDir)
	err := os.WriteFile(declarableDir, []byte(newData), 0755)
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
	g.removeFileLine(migratorTo, domain)
	g.removeFileLine(declarableDir, domain)

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
	for old, new := range replaces {
		strReplaced = strings.ReplaceAll(strReplaced, old, new)
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
			os.Mkdir(path, 0755)
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
