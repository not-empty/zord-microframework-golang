package generator

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-skeleton/pkg"
	"go-skeleton/pkg/logger"
	"go-skeleton/tools/generator"
)

type Generator struct {
	validator bool
	domain    string
}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) DeclareCommands(cmd *cobra.Command) {
	g.initFlags(cmd)
	createDomain := &cobra.Command{
		Use:    "create-domain",
		Short:  "Create a new domain service",
		PreRun: g.BootGenerator,
		Run:    g.CreateDomain,
	}
	createDomain.Flags().BoolVarP(&g.validator, "validator", "v", false, "Create domain with validation")
	cmd.AddCommand(createDomain)
	cmd.AddCommand(&cobra.Command{
		Use:    "destroy-domain",
		Short:  "Destroy a domain service",
		PreRun: g.BootGenerator,
		Run:    g.DestroyDomain,
	})
}

func (g *Generator) CreateDomain(_ *cobra.Command, args []string) {
	generatorInstance := generator.NewGenerator(pkg.CliDependencies["logger"].(*logger.Logger))
	if len(args) > 0 {
		g.domain = args[0]
	}
	err := generatorInstance.CreateDomain(g.domain, g.validator)
	pkg.Logger.Error(err)
}

func (g *Generator) DestroyDomain(_ *cobra.Command, args []string) {
	generatorInstance := generator.NewGenerator(pkg.CliDependencies["logger"].(*logger.Logger))
	if len(args) > 0 {
		g.domain = args[0]
	}
	err := generatorInstance.DestroyDomain(g.domain)
	pkg.Logger.Error(err)
}

func (g *Generator) BootGenerator(_ *cobra.Command, _ []string) {
	for index, dep := range pkg.CliDependencies {
		dep.Boot()
		pkg.Logger.Info(fmt.Sprintf("[Kernel.Kernel] Booting %s", index))
	}
}

func (g *Generator) initFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&g.domain, "domain", "", "Describe name to New Domain")
	cmd.MarkFlagsMutuallyExclusive("domain")
}
