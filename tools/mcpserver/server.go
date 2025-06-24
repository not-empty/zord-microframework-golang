package mcpserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-skeleton/pkg/config"
	"go-skeleton/pkg/logger"
	"go-skeleton/tools/arch_analyser"
	"go-skeleton/tools/generator"
	"go-skeleton/tools/migrator"
	"os"
	"os/signal"
	"syscall"

	mcp "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"
)

type Server struct {
	server    *mcp.Server
	logger    *logger.Logger
	config    *config.Config
	generator *generator.CodeGenerator
	migrator  *migrator.Migrator
}

type CreateDomainArgs struct {
	Name       string `json:"name" jsonschema:"required,description=Name of the domain to create"`
	Validator  bool   `json:"validator" jsonschema:"description=Whether to create domain with validation"`
	DomainType string `json:"domainType" jsonschema:"description=Type of domain to create (default: crud)"`
}

type DestroyDomainArgs struct {
	Name       string `json:"name" jsonschema:"required,description=Name of the domain to destroy"`
	DomainType string `json:"domainType" jsonschema:"description=Type of domain to destroy (default: crud)"`
}

type CreateDomainFromSchemaArgs struct {
	SchemaName string `json:"schemaName" jsonschema:"required,description=Name of the schema to use"`
	TableName  string `json:"tableName" jsonschema:"description=Optional specific table to generate"`
}

type MigrateArgs struct {
	Tenant string `json:"tenant" jsonschema:"description=Optional tenant name for migration"`
}

type GenerateSchemaArgs struct {
	SchemaName   string `json:"schemaName" jsonschema:"required,description=Name for the generated schema"`
	DatabaseName string `json:"databaseName" jsonschema:"description=Optional specific database to generate from"`
}

type InspectArgs struct{}
type ListOfferingsArgs struct{}
type ProjectContextArgs struct{}

type ArchAnalyseArgs struct{}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func NewServer() *Server {
	conf := config.NewConfig()
	err := conf.LoadEnvs()
	if err != nil {
		os.Setenv("ENVIRONMENT", getEnvOrDefault("ENVIRONMENT", "development"))
		os.Setenv("APP", getEnvOrDefault("APP", "zord"))
		os.Setenv("VERSION", getEnvOrDefault("VERSION", "1.0.0"))
		os.Setenv("DB_USER", getEnvOrDefault("DB_USER", "root"))
		os.Setenv("DB_PASS", getEnvOrDefault("DB_PASS", "root"))
		os.Setenv("DB_URL", getEnvOrDefault("DB_URL", "localhost"))
		os.Setenv("DB_PORT", getEnvOrDefault("DB_PORT", "3306"))
		os.Setenv("DB_DATABASE", getEnvOrDefault("DB_DATABASE", "zord"))
		os.Setenv("DB_TEST_DATABASE", getEnvOrDefault("DB_TEST_DATABASE", "zord_test"))
	}

	l := logger.NewLogger(
		getEnvOrDefault("ENVIRONMENT", "development"),
		getEnvOrDefault("APP", "zord"),
		getEnvOrDefault("VERSION", "1.0.0"),
	)

	l.Boot()

	m := migrator.NewMigrator(
		fmt.Sprintf("%s:%s@%s:%s",
			getEnvOrDefault("DB_USER", "root"),
			getEnvOrDefault("DB_PASS", "root"),
			getEnvOrDefault("DB_URL", "localhost"),
			getEnvOrDefault("DB_PORT", "3306"),
		),
		fmt.Sprintf("%s:%s@%s:%s/%s",
			getEnvOrDefault("DB_USER", "root"),
			getEnvOrDefault("DB_PASS", "root"),
			getEnvOrDefault("DB_URL", "localhost"),
			getEnvOrDefault("DB_PORT", "3306"),
			getEnvOrDefault("DB_TEST_DATABASE", "zord_test"),
		),
		getEnvOrDefault("DB_DATABASE", "zord"),
	)

	return &Server{
		server:    mcp.NewServer(stdio.NewStdioServerTransport()),
		logger:    l,
		config:    conf,
		generator: generator.NewCodeGenerator(l, false, "crud"),
		migrator:  m,
	}
}

func (s *Server) RegisterTools() error {
	if err := s.server.RegisterTool("create-domain", "Create a new domain service", s.handleCreateDomain); err != nil {
		return fmt.Errorf("failed to register create-domain tool: %v", err)
	}
	if err := s.server.RegisterTool("destroy-domain", "Destroy a domain service", s.handleDestroyDomain); err != nil {
		return fmt.Errorf("failed to register destroy-domain tool: %v", err)
	}
	if err := s.server.RegisterTool("create-domain-from-schema", "Create domain from database schema", s.handleCreateDomainFromSchema); err != nil {
		return fmt.Errorf("failed to register create-domain-from-schema tool: %v", err)
	}
	if err := s.server.RegisterTool("migrate", "Migrate database schema", s.handleMigrate); err != nil {
		return fmt.Errorf("failed to register migrate tool: %v", err)
	}
	if err := s.server.RegisterTool("inspect", "Inspect database schema", s.handleInspect); err != nil {
		return fmt.Errorf("failed to register inspect tool: %v", err)
	}
	if err := s.server.RegisterTool("generate-schema-from-db", "Generate schema from database", s.handleGenerateSchema); err != nil {
		return fmt.Errorf("failed to register generate-schema-from-db tool: %v", err)
	}
	if err := s.server.RegisterTool("project-context", "Get full project context and usage prompt", s.handleProjectContext); err != nil {
		return fmt.Errorf("failed to register project-context tool: %v", err)
	}
	if err := s.server.RegisterTool("ListOfferings", "List available tools", s.handleListOfferings); err != nil {
		return fmt.Errorf("failed to register ListOfferings tool: %v", err)
	}
	if err := s.server.RegisterTool("arch-analyse", "Executa a análise de arquitetura do projeto", s.handleArchAnalyse); err != nil {
		return fmt.Errorf("failed to register arch-analyse tool: %v", err)
	}
	return nil
}

func (s *Server) Serve() error {
	if err := s.RegisterTools(); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to register tools:", err)
		return err
	}
	return s.server.Serve()
}

func (s *Server) handleCreateDomain(args CreateDomainArgs) (*mcp.ToolResponse, error) {
	gen := generator.NewCodeGenerator(s.logger, args.Validator, args.DomainType)
	gen.Handler([]string{args.Name})
	return mcp.NewToolResponse(mcp.NewTextContent(fmt.Sprintf("Domain %s created successfully", args.Name))), nil
}

func (s *Server) handleDestroyDomain(args DestroyDomainArgs) (*mcp.ToolResponse, error) {
	destroyer := generator.NewCodeDestroy(s.logger, args.DomainType)
	destroyer.Handler([]string{args.Name})
	return mcp.NewToolResponse(mcp.NewTextContent(fmt.Sprintf("Domain %s destroyed successfully", args.Name))), nil
}

func (s *Server) handleCreateDomainFromSchema(args CreateDomainFromSchemaArgs) (*mcp.ToolResponse, error) {
	s.generator.ReadFromSchema(args.SchemaName, args.TableName)
	return mcp.NewToolResponse(mcp.NewTextContent(fmt.Sprintf("Domain created from schema %s successfully", args.SchemaName))), nil
}

func (s *Server) handleMigrate(args MigrateArgs) (*mcp.ToolResponse, error) {
	s.migrator.MigrateAllDomains(args.Tenant)
	return mcp.NewToolResponse(mcp.NewTextContent("Migration completed successfully")), nil
}

func (s *Server) handleInspect(args InspectArgs) (*mcp.ToolResponse, error) {
	s.migrator.Inspect()
	return mcp.NewToolResponse(mcp.NewTextContent("Database inspection completed")), nil
}

func (s *Server) handleGenerateSchema(args GenerateSchemaArgs) (*mcp.ToolResponse, error) {
	s.migrator.Generate(args.SchemaName, args.DatabaseName)
	return mcp.NewToolResponse(mcp.NewTextContent(fmt.Sprintf("Schema %s generated successfully", args.SchemaName))), nil
}

func (s *Server) handleProjectContext(args ProjectContextArgs) (*mcp.ToolResponse, error) {
	prompt := `O projeto zord-microframework-golang é um microframework Go que segue a arquitetura hexagonal (ports & adapters), promovendo separação clara entre domínio, infraestrutura e interfaces. Ele é altamente modular, extensível e voltado para automação de tarefas de backend, geração de código e operações de banco de dados.

---

Estrutura do Projeto e Exemplos de Código

- cmd/: Entrypoints para aplicações CLI e servidores.

- internal/application/services/: Serviços de aplicação (application services) que encapsulam regras de negócio e orquestram operações de domínio e repositório.
  Exemplo:
    // Service de GET para Dummy
    type Service struct {
        services.BaseService
        response   *Response
        repository dummy.Repository
    }
    func (s *Service) Execute(request Request) {
        if err := request.Validate(); err != nil {
            s.BadRequest(err.Error())
            return
        }
        s.produceResponseRule(request.Data, request.Domain)
    }
    func (s *Service) produceResponseRule(data *Data, domain *dummy.Dummy) {
        dummyData, err := s.repository.Get(*domain, "id", data.ID)
        if err != nil {
            if err.Error() != "sql: no rows in result set" {
                s.InternalServerError("error on get data", err)
                return
            }
        }
        s.response = &Response{ Data: dummyData }
    }

- internal/application/domain/: Entidades de domínio e interfaces de repositório.
  Exemplo:
    type Dummy struct {
        ID        string
        DummyName string
        Email     string
    }
    type Repository interface {
        base_repository.BaseRepository[Dummy]
    }

- internal/repositories/: Implementações de repositórios, baseados em interfaces genéricas.
  Exemplo:
    type BaseRepository[dom Domain] interface {
        Get(domain dom, field string, value string) (*dom, error)
        Create(data dom, tx *sqlx.Tx, autoCommit bool) error
        List(domain dom, limit int, offset int) (*[]dom, error)
        // ...
    }
    // Implementação genérica
    func (repo *BaseRepo[Domain]) Get(Data Domain, field string, value string) (*Domain, error) {
        row := repo.Mysql.QueryRowx(
            fmt.Sprintf("SELECT %s FROM %s WHERE %s = ?", strings.Join(repo.fields, ", "), Data.Schema(), field),
            value,
        )
        err := row.StructScan(&Data)
        if err != nil {
            return nil, err
        }
        return &Data, nil
    }

- internal/application/context/: Contextos de execução, como multi-tenant.
  Exemplo:
    type PrepareContext struct {
        Tenant string
    }
    func (ctx *PrepareContext) SetContext(data Data) {
        data.SetClient(ctx.Tenant)
    }

- internal/application/providers/: Providers utilitários, como filtros e paginação.
  Exemplo:
    type Filters struct {
        ParsedData []Filter
    }
    type Filter struct {
        Field    string
        Operator string
        Value    string
        IsString bool
    }
    func (f *Filters) Parse(config map[string]string, data map[string]FilterData) error {
        // ...
    }

---

Exemplo de Fluxo Hexagonal

1. Service recebe um request, valida e orquestra.
2. Domain representa a entidade de negócio.
3. Repository executa operações de persistência (CRUD) usando interfaces genéricas.
4. Provider (ex: filtros) auxilia na construção de queries dinâmicas.
5. Context injeta informações como tenant para multi-tenant.

---

Exemplo de Comando para rodar um resource (serviço):

- Para rodar um serviço ou comando CLI:
  go run cmd/algum_comando.go --flag valor

- Para rodar testes:
  go test ./...

- Para rodar o projeto em modo desenvolvimento:
  go run cmd/main.go

---

Comandos/Recursos Disponíveis (exemplos):
- create-domain
- destroy-domain
- create-domain-from-schema
- migrate
- inspect
- generate-schema-from-db

---

Resumo:
- Arquitetura hexagonal real, com exemplos de services, domains, repositories, providers e contexts.
- Modular, extensível, pronto para automação e integração com IA.
- Fácil de rodar, testar e estender.`
	return mcp.NewToolResponse(mcp.NewTextContent(prompt)), nil
}

func (s *Server) handleListOfferings(args ListOfferingsArgs) (*mcp.ToolResponse, error) {
	response := map[string]interface{}{
		"name":        "zord-mcp",
		"description": "Zord Microframework MCP Server for code generation and database management",
		"tools": []map[string]interface{}{
			{
				"name":        "create-domain",
				"description": "Create a new domain service",
				"parameters": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"name": map[string]interface{}{
							"type":        "string",
							"description": "Name of the domain to create",
						},
						"validator": map[string]interface{}{
							"type":        "boolean",
							"description": "Whether to create domain with validation",
						},
						"domainType": map[string]interface{}{
							"type":        "string",
							"description": "Type of domain to create (default: crud)",
						},
					},
					"required": []string{"name"},
				},
			},
			{
				"name":        "destroy-domain",
				"description": "Destroy a domain service",
				"parameters": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"name": map[string]interface{}{
							"type":        "string",
							"description": "Name of the domain to destroy",
						},
						"domainType": map[string]interface{}{
							"type":        "string",
							"description": "Type of domain to destroy (default: crud)",
						},
					},
					"required": []string{"name"},
				},
			},
			{
				"name":        "create-domain-from-schema",
				"description": "Create domain from database schema",
				"parameters": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"schemaName": map[string]interface{}{
							"type":        "string",
							"description": "Name of the schema to use",
						},
						"tableName": map[string]interface{}{
							"type":        "string",
							"description": "Optional specific table to generate",
						},
					},
					"required": []string{"schemaName"},
				},
			},
			{
				"name":        "migrate",
				"description": "Migrate database schema",
				"parameters": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"tenant": map[string]interface{}{
							"type":        "string",
							"description": "Optional tenant name for migration",
						},
					},
				},
			},
			{
				"name":        "inspect",
				"description": "Inspect database schema",
				"parameters": map[string]interface{}{
					"type":       "object",
					"properties": map[string]interface{}{},
				},
			},
			{
				"name":        "generate-schema-from-db",
				"description": "Generate schema from database",
				"parameters": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"schemaName": map[string]interface{}{
							"type":        "string",
							"description": "Name for the generated schema",
						},
						"databaseName": map[string]interface{}{
							"type":        "string",
							"description": "Optional specific database to generate from",
						},
					},
					"required": []string{"schemaName"},
				},
			},
			{
				"name":        "project-context",
				"description": "Get full project context and usage prompt",
				"parameters": map[string]interface{}{
					"type":       "object",
					"properties": map[string]interface{}{},
				},
			},
		},
	}

	content, err := json.Marshal(response)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response: %v", err)
	}

	return mcp.NewToolResponse(mcp.NewTextContent(string(content))), nil
}

func (s *Server) handleArchAnalyse(args ArchAnalyseArgs) (*mcp.ToolResponse, error) {
	root, _ := os.Getwd()
	var buf bytes.Buffer

	// Chama as validações principais e captura o output
	validations := []struct {
		name string
		fn   func(string) error
	}{
		{"Validação de importações", arch_analyser.ValidateImports},
		{"Validação dos parâmetros de NewService", arch_analyser.ValidateNewServiceParams},
		{"Validação de queries no banco", arch_analyser.ValidateDbQueriesInRepositories},
		{"Validação de dependências entre camadas", arch_analyser.ValidateLayerDependencies},
		{"Validação de circularidade de dependências", arch_analyser.ValidateNoCircularDependencies},
		{"Validação de uso correto de context", arch_analyser.ValidateContextUsage},
		{"Validação de pureza dos providers", arch_analyser.ValidateProvidersArePure},
		{"Validação de implementação de interfaces de repositório", arch_analyser.ValidateRepositoryInterfacesImplemented},
		{"Validação de variáveis globais mutáveis", arch_analyser.ValidateNoGlobalVars},
		{"Validação de orquestração", arch_analyser.ValidateOrchestrationOnlyInServices},
		{"Validação de nomenclatura e localização", arch_analyser.ValidateNamingAndLocation},
		{"Validação de uso de pacotes externos", arch_analyser.ValidateExternalPackagesUsage},
		{"Validação de uso de reflection", arch_analyser.ValidateReflectionUsage},
		{"Validação de type assertions", arch_analyser.ValidateTypeAssertions},
	}
	allOk := true
	for _, v := range validations {
		err := v.fn(root)
		if err != nil {
			buf.WriteString("[ERRO] " + v.name + ":\n" + err.Error() + "\n")
			allOk = false
		} else {
			buf.WriteString(v.name + " concluída com sucesso!\n")
		}
	}
	if allOk {
		return mcp.NewToolResponse(mcp.NewTextContent(buf.String())), nil
	}
	return mcp.NewToolResponse(mcp.NewTextContent(buf.String())), fmt.Errorf("Uma ou mais validações falharam")
}

// Run inicializa e executa o servidor MCP
func Run() {
	transport := stdio.NewStdioServerTransport()
	server := mcp.NewServer(transport)

	s := &Server{
		server:    server,
		logger:    logger.NewLogger("development", "zord", "1.0.0"),
		config:    config.NewConfig(),
		generator: generator.NewCodeGenerator(logger.NewLogger("development", "zord", "1.0.0"), false, "crud"),
		migrator:  migrator.NewMigrator("root:root@localhost:3306", "root:root@localhost:3306/zord_test", "zord"),
	}

	if err := s.RegisterTools(); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to register tools:", err)
		os.Exit(1)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := server.Serve(); err != nil {
			fmt.Fprintln(os.Stderr, "Server error:", err)
			os.Exit(1)
		}
	}()

	<-sigChan
	fmt.Fprintln(os.Stderr, "Shutting down server...")
}
