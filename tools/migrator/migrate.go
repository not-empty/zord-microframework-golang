package migrator

import (
	"context"
	"fmt"
	"os"
	"strings"

	"ariga.io/atlas-go-sdk/atlasexec"
)

type Migrator struct {
	dsn          string
	dsnTest      string
	databaseName string
}

func NewMigrator(dsn string, dsnTest string, databaseName string) *Migrator {
	return &Migrator{
		dsn:          dsn,
		dsnTest:      dsnTest,
		databaseName: databaseName,
	}
}

func (m *Migrator) MigrateAllDomains(Tenant string) {
	workdir, err := atlasexec.NewWorkingDir(
		atlasexec.WithAtlasHCLPath("schemas/schema.my.hcl"),
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

	params := &atlasexec.SchemaApplyParams{
		DevURL: "mysql://" + m.dsnTest,
		To:     "file://" + workdir.Path(),
		URL:    "mysql://" + m.dsn,
		Vars: atlasexec.Vars{
			"database": m.databaseName,
		},
		Schema: []string{
			m.databaseName,
		},
	}

	if Tenant != "" {
		params.Vars = atlasexec.Vars{
			"database": Tenant,
		}
		params.Schema = []string{
			Tenant,
		}
	}

	res, err := client.SchemaApply(context.Background(), params)
	if err != nil {
		fmt.Printf("failed to apply migrations: %v", err)
		return
	}
	fmt.Printf("Applied %d changes:\n", len(res.Changes.Applied))
	fmt.Println(strings.Join(res.Changes.Applied, "\n"))
}

func (m *Migrator) Inspect() {
	workdir, err := atlasexec.NewWorkingDir(
		atlasexec.WithAtlasHCLPath("schemas/schema.my.hcl"),
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

func (m *Migrator) Generate(schemaName string) {
	workdir, err := atlasexec.NewWorkingDir(
		atlasexec.WithAtlasHCLPath("schemas/schema.my.hcl"),
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
	err = os.WriteFile("tools/migrator/schema/"+schemaName+".my.hcl", []byte(res), 0755)
	if err != nil {
		fmt.Println(err)
		return
	}
}
