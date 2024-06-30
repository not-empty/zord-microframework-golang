package struct_reader

import (
	"github.com/spf13/cobra"
	"go-skeleton/tools/structReader"
)

type StructReader struct {
}

func NewStructReader() *StructReader {
	return &StructReader{}
}

func (sr *StructReader) DeclareCommands(cmd *cobra.Command) {
	cmd.AddCommand(
		&cobra.Command{
			Use:   "generate-schema-from-domain",
			Short: "criar HCL com base na struct de dominio",
			Run:   sr.Generate,
		},
	)
}

func (sr *StructReader) Generate(_ *cobra.Command, args []string) {
	structReader.GenerateHclFileFromDomain("")
}
