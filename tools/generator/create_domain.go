package generator

import (
	"fmt"
	"go-skeleton/application/services"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
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
	domainDir := "application/domain/" + domain

	if _, err := os.Stat(domainDir); os.IsNotExist(err) {
		os.Mkdir(domainDir, 0755)
	}

	filepath.Walk("tools/generator/stubs/domain", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		g.createFile(domain, domainCap, path, domainDir)
		return nil
	})

	servicesDir := "application/services/" + domain
	if _, err := os.Stat(servicesDir); os.IsNotExist(err) {
		os.Mkdir(servicesDir, 0755)
	}

	filepath.Walk("tools/generator/stubs/services", func(path string, info fs.FileInfo, err error) error {
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

	repositoriesDir := "pkg/repositories/" + domain
	if _, err := os.Stat(repositoriesDir); os.IsNotExist(err) {
		os.Mkdir(repositoriesDir, 0755)
	}

	filepath.Walk("tools/generator/stubs/repositories", func(path string, info fs.FileInfo, err error) error {
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

	routesDir := "cmd/handlers/http/routes"
	g.createFile(domain, domainCap, "tools/generator/stubs/routes/__{{domain}}.go.stub", routesDir)

	declarableDir := routesDir + "/declarable.go"

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
