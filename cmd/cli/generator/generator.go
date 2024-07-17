package generator

import (
	"errors"
	"go-skeleton/pkg/config"
	"go-skeleton/pkg/logger"
	"go-skeleton/tools/generator"

	"github.com/spf13/cobra"
)

type Generator struct {
	Logger *logger.Logger
	Flags  Flags
}

type Flags struct {
	validator  bool
	domainType string
	domain     string
}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) DeclareCommands(cmd *cobra.Command) {
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

	destroyDomain.Flags().StringVar(&g.Flags.domainType, "type", "crud", "Define domain type: ['crud'|'<custom>']")

	cmd.AddCommand(
		createDomain,
		destroyDomain,
	)
}

func (g *Generator) DeclareDomainCreatorFromSchema(cmd *cobra.Command) {
	generateFromDb := &cobra.Command{
		Use:    "create-domain-from-schema",
		Short:  "create-domain-from-schema <schema name> <table name>",
		Long:   "this command will read the chosen schema and create a crud for each table (if you pass a specific table it will generate only this)",
		PreRun: g.BootGenerator,
		Run:    g.GenerateFromDb,
	}
	cmd.AddCommand(
		generateFromDb,
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

func (g *Generator) GenerateFromDb(_ *cobra.Command, args []string) {
	if len(args) == 0 {
		g.Logger.Error(errors.New("empty args"))
	}
	table := ""
	if len(args) > 1 {
		table = args[1]
	}
	generator.NewCodeGenerator(
		g.Logger,
		g.Flags.validator,
		g.Flags.domainType,
	).ReadFromSchema(args[0], table)
}

func (g *Generator) BootGenerator(_ *cobra.Command, _ []string) {
	conf := config.NewConfig()
	err := conf.LoadEnvs()
	if err != nil {
		panic(err)
	}

	l := logger.NewLogger(
		conf.ReadConfig("ENVIRONMENT"),
		conf.ReadConfig("APP"),
		conf.ReadConfig("VERSION"),
	)

	l.Boot()
	g.Logger = l
}
