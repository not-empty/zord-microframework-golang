package generator

import (
	"os"
	"strings"
)

func CreateDomain(domain string) {
	domainDir := "application/domain/" + domain
	domainStub := "tools/generator/stubs/domain.stub"

	stubData, err := os.ReadFile(domainStub)
	if err != nil {
		panic(err)
	}

	// Todo: use another method
	domainCap := strings.Title(domain)
	domainContent := strings.ReplaceAll(string(stubData), "{{domain}}", domain)
	domainContent = strings.ReplaceAll(string(domainContent), "{{domainCap}}", domainCap)

	os.Mkdir(domainDir, 0755)
	os.WriteFile(domainDir+"/"+domain+".go", []byte(domainContent), 0755)
}
