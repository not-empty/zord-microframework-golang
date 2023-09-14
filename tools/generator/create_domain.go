package generator

import (
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

func (g *Generator) CreateDomain(domain string) {
	domainCap := strings.Title(domain)
	domainDir := "application/domain/" + domain

	if _, err := os.Stat(domainDir); os.IsNotExist(err) {
		os.Mkdir(domainDir, 0755)
	}

	filepath.Walk("tools/generator/stubs/domain", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return err
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
}

func (g *Generator) createFile(domain string, domainCap string, from string, to string) {
	stubData, err := os.ReadFile(from)
	if err != nil {
		panic(err)
	}

	domainContent := strings.ReplaceAll(string(stubData), "{{domain}}", domain)
	domainContent = strings.ReplaceAll(string(domainContent), "{{domainCap}}", domainCap)

	fileName := strings.Split(from, "__")[1]
	fileName = strings.ReplaceAll(fileName, "{{domain}}", domain)
	fileName = strings.ReplaceAll(fileName, ".stub", "")

	err = os.WriteFile(to+"/"+fileName, []byte(domainContent), 0755)
	if err != nil {
		panic(err)
	}
	g.Logger.Info("CREATE FILE: " + to + "/" + fileName)
}
