package migrator

import (
	"context"
	"fmt"
	"go-skeleton/tools/migrator/hcl"
	"os"
	"strings"

	"github.com/fatih/structs"

	"ariga.io/atlas-go-sdk/atlasexec"
)

type Migrator struct {
	dsn     string
	dsnTest string
}

func NewMigrator(dsn string, dsnTest string) *Migrator {
	return &Migrator{
		dsn:     dsn,
		dsnTest: dsnTest,
	}
}

func (m *Migrator) MigrateAllDomains() {
	workdir, err := atlasexec.NewWorkingDir(
		atlasexec.WithAtlasHCLPath("tools/migrator/schema/schema.my.hcl"),
	)
	if err != nil {
		fmt.Printf("failed to load working directory: %v", err)
		return
	}
	defer workdir.Close()

	client, err := atlasexec.NewClient(workdir.Path(), "atlas")
	if err != nil {
		fmt.Printf("failed to initialize client: %v", err)
		return
	}

	res, err := client.SchemaApply(context.Background(), &atlasexec.SchemaApplyParams{
		DevURL: "mysql://" + m.dsnTest,
		To:     "file://" + workdir.Path(),
		URL:    "mysql://" + m.dsn,
	})
	if err != nil {
		fmt.Printf("failed to apply migrations: %v", err)
		return
	}
	fmt.Printf("Applied %d changes:\n", len(res.Changes.Applied))
	fmt.Println(strings.Join(res.Changes.Applied, "\n"))
}

func (m *Migrator) Inspect() {
	workdir, err := atlasexec.NewWorkingDir(
		atlasexec.WithAtlasHCLPath("tools/migrator/schema/schema.my.hcl"),
	)
	if err != nil {
		fmt.Printf("failed to load working directory: %v", err)
		return
	}
	defer workdir.Close()

	client, err := atlasexec.NewClient(workdir.Path(), "atlas")
	if err != nil {
		fmt.Printf("failed to initialize client: %v", err)
		return
	}

	res, err := client.SchemaInspect(context.Background(), &atlasexec.SchemaInspectParams{
		DevURL: "mysql://" + m.dsnTest,
		URL:    "mysql://" + m.dsn,
	})
	if err != nil {
		fmt.Printf("failed to inspect schema: %v", err)
		return
	}
	fmt.Println(res)
}

func (m *Migrator) Generate() {
	workdir, err := atlasexec.NewWorkingDir(
		atlasexec.WithAtlasHCLPath("tools/migrator/schema/schema.my.hcl"),
	)
	if err != nil {
		fmt.Printf("failed to load working directory: %v", err)
		return
	}
	defer workdir.Close()

	client, err := atlasexec.NewClient(workdir.Path(), "atlas")
	if err != nil {
		fmt.Printf("failed to initialize client: %v", err)
		return
	}

	res, err := client.SchemaInspect(context.Background(), &atlasexec.SchemaInspectParams{
		DevURL: "mysql://" + m.dsnTest,
		URL:    "mysql://" + m.dsn,
	})
	if err != nil {
		fmt.Printf("failed to inspect schema: %v", err)
		return
	}
	err = os.WriteFile("tools/migrator/generated/database.hcl", []byte(res), 0755)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (m *Migrator) MigrateFromDomains() {
	tables := GetTables()
	_hcl := hcl.NewHCL()
	for _, table := range tables {
		fields := structs.Fields(table)
		listFields := m.getFields("db", fields)
		dbDefinitions := m.getFields("zord_db", fields)
		hclConfig := _hcl.CreateAtlasHCLFromDomain(table.Schema(), "skeleton", listFields, dbDefinitions)
		fmt.Println(hclConfig)
		// Todo: adicionar ao arquivo de schemas
	}
}

func (m *Migrator) getFields(tag string, fields []*structs.Field) []string {
	var listFields []string
	for _, field := range fields {
		tag := field.Tag(tag)
		listFields = append(listFields, tag)
	}
	return listFields
}
