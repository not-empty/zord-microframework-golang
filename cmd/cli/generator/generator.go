package generator

import (
	"errors"
	"fmt"
	"go-skeleton/application/services"
	"go-skeleton/pkg"
	"go-skeleton/tools/generator"

	"github.com/spf13/cobra"
)

type Generator struct {
	Logger services.Logger
	Flags  Flags
}

type Flags struct {
	validator  bool
	domainType string
	domain     string
}

func NewGenerator(l services.Logger) *Generator {
	return &Generator{
		Logger: l,
	}
}

func (g *Generator) DeclareCommands(cmd *cobra.Command) {
	g.initFlags(cmd)
	createDomain := &cobra.Command{
		Use:    "create-domain",
		Short:  "Create a new domain service",
		PreRun: g.BootGenerator,
		Run:    g.CreateDomain,
	}

	createDomain.Flags().BoolVarP(&g.Flags.validator, "validator", "v", false, "Create domain with validation")
	createDomain.Flags().StringVar(&g.Flags.domainType, "type", "crud", "Define domain type: ['crud'|'<custom>']")

	destroyDomain := &cobra.Command{
		Use:    "destroy-domain",
		Short:  "Destroy a domain service",
		PreRun: g.BootGenerator,
		Run:    g.DestroyDomain,
	}

	cmd.AddCommand(
		createDomain,
		destroyDomain,
	)
}

func (g *Generator) CreateDomain(_ *cobra.Command, args []string) {
	if len(args) == 0 {
		g.Logger.Error(errors.New("empty args"))
	}
	generator.NewCodeGenerator(
		g.Logger,
		g.Flags.validator,
		g.Flags.domainType,
	).Handler(
		args,
	)
}

func (g *Generator) DestroyDomain(_ *cobra.Command, args []string) {
	if len(args) == 0 {
		g.Logger.Error(errors.New("empty args"))
	}
	generator.NewCodeDestroy(
		g.Logger,
		g.Flags.domainType,
	).Handler(
		args,
	)
}

func (g *Generator) BootGenerator(_ *cobra.Command, _ []string) {
	for index, dep := range pkg.CliDependencies {
		dep.Boot()
		pkg.Logger.Info(fmt.Sprintf("[Kernel.Kernel] Booting %s", index))
	}
}

func (g *Generator) initFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&g.Flags.domain, "domain", "", "Describe name to New Domain")
	cmd.MarkFlagsMutuallyExclusive("domain")
}
