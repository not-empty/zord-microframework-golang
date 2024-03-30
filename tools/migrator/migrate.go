package migrator

import (
	"ariga.io/atlas-go-sdk/atlasexec"
	"context"
	"fmt"
	"strings"
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
