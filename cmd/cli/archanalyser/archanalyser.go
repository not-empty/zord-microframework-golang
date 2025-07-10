package archanalyser

import (
	"fmt"
	"go-skeleton/tools/arch_analyser"
	"os"

	"github.com/spf13/cobra"
)

type ArchAnalyser struct{}

func NewArchAnalyser() *ArchAnalyser {
	return &ArchAnalyser{}
}

func (a *ArchAnalyser) DeclareCommands(cmd *cobra.Command) {
	cmd.AddCommand(
		&cobra.Command{
			Use:   "arch-analyse",
			Short: "Analisa as importações dos domínios para garantir regras de arquitetura",
			Run:   a.Analyse,
		},
	)
}

func (a *ArchAnalyser) Analyse(_ *cobra.Command, _ []string) {
	root, err := os.Getwd()
	if err != nil {
		fmt.Println("Erro ao obter diretório de trabalho:", err)
		os.Exit(1)
	}
	// Validação de importações
	err = arch_analyser.ValidateImports(root)
	if err != nil {
		fmt.Println("[ERRO]", err)
		os.Exit(1)
	}
	fmt.Println("Validação de importações concluída com sucesso!")

	// Validação de parâmetros de NewService
	err = arch_analyser.ValidateNewServiceParams(root)
	if err != nil {
		fmt.Println("[ERRO]", err)
		os.Exit(1)
	}
	fmt.Println("Validação dos parâmetros de NewService concluída com sucesso!")

	// Validação de queries no banco
	err = arch_analyser.ValidateDbQueriesInRepositories(root)
	if err != nil {
		fmt.Println("[ERRO]", err)
		os.Exit(1)
	}
	fmt.Println("Validação de queries no banco concluída com sucesso!")

	// Validação de dependências entre camadas
	err = arch_analyser.ValidateLayerDependencies(root)
	if err != nil {
		fmt.Println("[ERRO]", err)
		os.Exit(1)
	}
	fmt.Println("Validação de dependências entre camadas concluída com sucesso!")

	// Validação de circularidade de dependências
	err = arch_analyser.ValidateNoCircularDependencies(root)
	if err != nil {
		fmt.Println("[ERRO]", err)
		os.Exit(1)
	}
	fmt.Println("Validação de circularidade de dependências concluída com sucesso!")

	// Validação de uso correto de context
	err = arch_analyser.ValidateContextUsage(root)
	if err != nil {
		fmt.Println("[ERRO]", err)
		os.Exit(1)
	}
	fmt.Println("Validação de uso de context concluída com sucesso!")

	// Validação de pureza dos providers
	err = arch_analyser.ValidateProvidersArePure(root)
	if err != nil {
		fmt.Println("[ERRO]", err)
		os.Exit(1)
	}
	fmt.Println("Validação de pureza dos providers concluída com sucesso!")

	// Validação de implementação de interfaces de repositório
	err = arch_analyser.ValidateRepositoryInterfacesImplemented(root)
	if err != nil {
		fmt.Println("[ERRO]", err)
		os.Exit(1)
	}
	fmt.Println("Validação de implementação de interfaces de repositório concluída com sucesso!")

	// Validação de variáveis globais mutáveis
	err = arch_analyser.ValidateNoGlobalVars(root)
	if err != nil {
		fmt.Println("[ERRO]", err)
		os.Exit(1)
	}
	fmt.Println("Validação de variáveis globais mutáveis concluída com sucesso!")

	// Validação de orquestração
	err = arch_analyser.ValidateOrchestrationOnlyInServices(root)
	if err != nil {
		fmt.Println("[ERRO]", err)
		os.Exit(1)
	}
	fmt.Println("Validação de orquestração concluída com sucesso!")

	// Validação de nomenclatura e localização
	err = arch_analyser.ValidateNamingAndLocation(root)
	if err != nil {
		fmt.Println("[ERRO]", err)
		os.Exit(1)
	}
	fmt.Println("Validação de nomenclatura e localização concluída com sucesso!")

	// Validação de uso de pacotes externos
	err = arch_analyser.ValidateExternalPackagesUsage(root)
	if err != nil {
		fmt.Println("[ERRO]", err)
		os.Exit(1)
	}
	fmt.Println("Validação de uso de pacotes externos concluída com sucesso!")

	// Validação de uso de reflection
	err = arch_analyser.ValidateReflectionUsage(root)
	if err != nil {
		fmt.Println("[ERRO]", err)
		os.Exit(1)
	}
	fmt.Println("Validação de uso de reflection concluída com sucesso!")

	// Validação de type assertions (.())
	err = arch_analyser.ValidateTypeAssertions(root)
	if err != nil {
		fmt.Println("[ERRO]", err)
		os.Exit(1)
	}
	fmt.Println("Validação de type assertions (.()) concluída com sucesso!")
}
