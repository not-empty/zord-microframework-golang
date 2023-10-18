package generator

import (
	"go-skeleton/application/services"

	"github.com/spf13/cobra"
)

type CodeGenerator struct {
	Logger services.Logger
}

func NewCodeGenerator(l services.Logger) *CodeGenerator {
	return &CodeGenerator{
		Logger: l,
	}
}

func (g *CodeGenerator) Handler(cmd *cobra.Command) {
	g.getFlags(cmd)
}

func (g *CodeGenerator) getFlags(cmd *cobra.Command) {
	cmd.Flags().GetBool()
}

func (g *CodeGenerator) isInvalidDomain(domain string) bool {
	return domain == ""
}
