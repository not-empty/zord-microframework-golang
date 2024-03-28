package migrator

import (
	"ariga.io/atlas-go-sdk/atlasexec"
	"context"
	"fmt"
)

type Migrator struct {
	dsn string
}

func NewMigrator(dsn string) *Migrator {
	return &Migrator{
		dsn: dsn,
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
	defer func(workdir *atlasexec.WorkingDir) {
		err := workdir.Close()
		if err != nil {
			fmt.Printf("failed to load working directory: %v", err)
		}
	}(workdir)

	client, err := atlasexec.NewClient("tools/migrator", "atlas")
	if err != nil {
		fmt.Printf("failed to initialize client: %v", err)
		return
	}

	res, err := client.MigrateApply(context.Background(), &atlasexec.MigrateApplyParams{
		URL: "mysql://" + m.dsn,
	})
	if err != nil {
		fmt.Printf("failed to apply migrations: %v", err)
		return
	}
	fmt.Printf("Applied %d migrations\n", len(res.Applied))
}
