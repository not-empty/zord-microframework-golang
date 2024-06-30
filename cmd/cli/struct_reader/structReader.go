package struct_reader

import (
	"go-skeleton/pkg/config"
	"go-skeleton/tools/structReader"

	"github.com/spf13/cobra"
)

type StructReader struct {
	GenerateFlags struct {
		Domain string
	}
	Config *config.Config
}

func NewStructReader(config *config.Config) *StructReader {
	return &StructReader{
		Config: config,
	}
}

func (sr *StructReader) DeclareCommands(cmd *cobra.Command) {
	generateSchemaFromDomain := &cobra.Command{
		Use:   "generate-schema-from-domain",
		Short: "create a HCL file based on domain struct",
		Run:   sr.Generate,
	}

	generateSchemaFromDomain.Flags().StringVar(&sr.GenerateFlags.Domain, "domain", "", "Generate to specific domain")

	cmd.AddCommand(generateSchemaFromDomain)
}

func (sr *StructReader) Generate(cmd *cobra.Command, args []string) {
	schema := sr.Config.ReadConfig("DB_DATABASE")
	structReader.GenerateHclFileFromDomain(schema, sr.GenerateFlags.Domain)
}
