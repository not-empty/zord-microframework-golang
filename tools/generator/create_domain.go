package generator

import (
	"fmt"
	"go-skeleton/application/services"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var (
	domainTo         = "application/domain/"
	domainFrom       = "tools/generator/stubs/domain/"
	servicesTo       = "application/services/"
	servicesFrom     = "tools/generator/stubs/services/"
	repositoriesTo   = "pkg/repositories/"
	repositoriesFrom = "tools/generator/stubs/repositories/"
	routesTo         = "cmd/handlers/http/routes"
	routesFrom       = "tools/generator/stubs/routes/__{{domain}}.go.stub"
	declarableDir    = routesTo + "/declarable.go"
	migratorTo       = "tools/migrator/migrate.go"
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
	domainCap := strings.Title(domain)
	domainDir := domainTo + domain

	if _, err := os.Stat(domainDir); os.IsNotExist(err) {
		os.Mkdir(domainDir, 0755)
	}

	filepath.Walk(domainFrom, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		g.createFile(domain, domainCap, path, domainDir)
		return nil
	})

	servicesDir := servicesTo + domain
	if _, err := os.Stat(servicesDir); os.IsNotExist(err) {
		os.Mkdir(servicesDir, 0755)
	}

	filepath.Walk(servicesFrom, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			if len(strings.Split(path, "__")) == 1 {
				return nil
			}
			folderPath := "/" + strings.Split(path, "__")[1]
			folderPath = servicesDir + folderPath

			if _, err := os.Stat(folderPath); os.IsNotExist(err) {
				os.Mkdir(folderPath, 0755)
			}
			return err
		}
		g.createFile(domain, domainCap, path, servicesDir)
		return nil
	})

	repositoriesDir := repositoriesTo + domain
	if _, err := os.Stat(repositoriesDir); os.IsNotExist(err) {
		os.Mkdir(repositoriesDir, 0755)
	}

	filepath.Walk(repositoriesFrom, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			if len(strings.Split(path, "__")) == 1 {
				return nil
			}
			folderPath := "/" + strings.Split(path, "__")[1]
			folderPath = strings.ReplaceAll(folderPath, "{{domain}}", domain)
			folderPath = repositoriesDir + folderPath

			if _, err := os.Stat(folderPath); os.IsNotExist(err) {
				os.Mkdir(folderPath, 0755)
			}
			return err
		}
		g.createFile(domain, domainCap, path, repositoriesDir)
		return nil
	})

	g.createFile(domain, domainCap, routesFrom, routesTo)

	data, err := os.ReadFile(declarableDir)
	if err != nil {
		g.Logger.Error(err)
	}

	if strings.Contains(string(data), domain) {
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

	newFileData := strings.Replace(string(data), "//{{codeGen1}}", routesInstanceTemplate, 1)
	newFileData = strings.Replace(newFileData, "//{{codeGen2}}", declarableTemplate, 1)

	g.Logger.Info("ADD ROUTES: " + declarableDir)
	err = os.WriteFile(declarableDir, []byte(newFileData), 0755)
	if err != nil {
		g.Logger.Error(err)
	}

	importTemplate := fmt.Sprintf("\"go-skeleton/application/domain/%s\"\n	//{{codeGen1}}", domain)
	createTableTemplate := fmt.Sprintf("m.db.Db.Migrator().CreateTable(&%s.%s{})\n	//{{codeGen2}}", domain, domainCap)

	migratorData, err := os.ReadFile(migratorTo)
	if err != nil {
		g.Logger.Error(err)
	}

	if strings.Contains(string(migratorData), domain) {
		return nil
	}

	newMigFileData := strings.Replace(string(migratorData), "//{{codeGen1}}", importTemplate, 1)
	newMigFileData = strings.Replace(newMigFileData, "//{{codeGen2}}", createTableTemplate, 1)

	g.Logger.Info("ADD MIGRATOR: " + migratorTo)
	err = os.WriteFile(migratorTo, []byte(newMigFileData), 0755)
	if err != nil {
		g.Logger.Error(err)
	}

	return nil
}

func (g *Generator) DestroyDomain(domain string) error {
	migData, err := os.ReadFile(migratorTo)
	if err != nil {
		g.Logger.Error(err)
	}

	lines1 := strings.Split(string(migData), "\n")
	for i, l := range lines1 {
		if strings.Contains(string(l), domain) {
			lines1 = append(lines1[:i], lines1[i+1:]...)
		}
	}

	err = os.WriteFile(migratorTo, []byte(strings.Join(lines1, "\n")), 0755)
	if err != nil {
		g.Logger.Error(err)
	}

	data, err := os.ReadFile(declarableDir)
	if err != nil {
		g.Logger.Error(err)
	}

	lines2 := strings.Split(string(data), "\n")
	for i, l := range lines2 {
		if strings.Contains(string(l), domain) {
			lines2 = append(lines2[:i], lines2[i+1:]...)
		}
	}

	err = os.WriteFile(declarableDir, []byte(strings.Join(lines2, "\n")), 0755)
	if err != nil {
		g.Logger.Error(err)
	}

	destroyPaths := []string{
		routesTo + "/" + domain + ".go",
		repositoriesTo + domain,
		servicesTo + domain,
		domainTo + domain,
	}

	for _, p := range destroyPaths {
		if err = os.RemoveAll(p); err != nil {
			g.Logger.Error(err)
		}
	}

	return nil
}

func (g *Generator) createFile(domain string, domainCap string, from string, to string) {
	stubData, err := os.ReadFile(from)
	if err != nil {
		g.Logger.Error(err)
	}

	domainContent := strings.ReplaceAll(string(stubData), "{{domain}}", domain)
	domainContent = strings.ReplaceAll(string(domainContent), "{{domainCap}}", domainCap)

	fileName := strings.Split(from, "__")[1]
	fileName = strings.ReplaceAll(fileName, "{{domain}}", domain)
	fileName = strings.ReplaceAll(fileName, ".stub", "")

	err = os.WriteFile(to+"/"+fileName, []byte(domainContent), 0755)
	if err != nil {
		g.Logger.Error(err)
	}
	g.Logger.Info("CREATE FILE: " + to + "/" + fileName)
}
