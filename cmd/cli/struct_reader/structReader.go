package struct_reader

import (
	"go-skeleton/tools/structReader"

	"github.com/spf13/cobra"
)

type StructReader struct {
	GenerateFlags struct {
		Domain string
	}
}

func NewStructReader() *StructReader {
	return &StructReader{}
}

func (sr *StructReader) DeclareCommands(cmd *cobra.Command) {
	generateSchemaFromDomain := &cobra.Command{
		Use:   "generate-schema-from-domain <schema_name>",
		Short: "create a HCL file based on domain struct",
		Run:   sr.Generate,
	}

	generateSchemaFromDomain.Flags().StringVar(&sr.GenerateFlags.Domain, "domain", "", "Generate to specific domain")

	cmd.AddCommand(generateSchemaFromDomain)
}

func (sr *StructReader) Generate(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		panic("empty schema name arg")
	}
	schema := args[0]
	structReader.GenerateHclFileFromDomain(schema, sr.GenerateFlags.Domain)
}
