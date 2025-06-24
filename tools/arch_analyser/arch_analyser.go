package arch_analyser

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

// getAllowedPkgs lê o arquivo TOML e retorna o set de pacotes permitidos
func getAllowedPkgs(configPath string) map[string]struct{} {
	type allowedConfig struct {
		Pkgs []string `toml:"pkgs"`
	}
	type rootConfig struct {
		Allowed allowedConfig `toml:"allowed"`
	}
	allowed := make(map[string]struct{})
	file, err := os.ReadFile(configPath)
	if err == nil {
		var conf rootConfig
		err = toml.Unmarshal(file, &conf)
		if err == nil {
			for _, pkg := range conf.Allowed.Pkgs {
				allowed[pkg] = struct{}{}
			}
			return allowed
		}
	}
	// fallback: lista padrão mínima
	for _, pkg := range []string{"fmt", "strings", "os", "io", "bytes", "time", "context", "errors"} {
		allowed[pkg] = struct{}{}
	}
	return allowed
}

// isAllowedInternalImport verifica se o arquivo pode importar um pacote internal
func isAllowedInternalImport(filePath, importPath string) bool {
	if strings.Contains(importPath, "/internal/") {
		if strings.Contains(filePath, string(os.PathSeparator)+"internal"+string(os.PathSeparator)) {
			return true // Permitido: arquivos em internal podem importar internal
		}
		if strings.Contains(filePath, string(os.PathSeparator)+"cmd"+string(os.PathSeparator)) {
			return true // Permitido: arquivos em cmd podem importar internal
		}
		return false // Bloqueado para outros
	}
	return false // Não é um import de internal
}

// isAllowedStdPkg verifica se o import é de um pacote padrão permitido
func isAllowedStdPkg(importPath string, allowedPkgs map[string]struct{}) bool {
	_, ok := allowedPkgs[importPath]
	return ok
}

// ValidateImports percorre os arquivos Go de internal e valida as importações proibidas.
func ValidateImports(root string) error {
	internalPath := filepath.Join(root, "internal")
	configPath := filepath.Join(filepath.Dir(internalPath), "tools", "arch_analyser", "allowed_imports.toml")
	allowedPkgs := getAllowedPkgs(configPath)

	return filepath.Walk(internalPath, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(path, ".go") {
			return err
		}
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, path, nil, parser.ImportsOnly)
		if err != nil {
			return err
		}
		for _, imp := range f.Imports {
			importPath := strings.Trim(imp.Path.Value, "\"")
			if isAllowedInternalImport(path, importPath) {
				continue
			}
			if isAllowedStdPkg(importPath, allowedPkgs) {
				continue
			}
			if strings.Contains(importPath, "/internal/") {
				return fmt.Errorf("%s importa %s, mas apenas arquivos em 'cmd' podem importar pacotes que contenham /internal/", path, importPath)
			}
			return fmt.Errorf("%s importa %s, que não é permitido (apenas pacotes padrão do Go ou, se em cmd, pacotes que contenham /internal/)", path, importPath)
		}
		return nil
	})
}

// ValidateNewServiceParams garante que todos os parâmetros de NewService são interfaces e não tipos externos
func ValidateNewServiceParams(root string) error {
	var errors []string
	internalPath := filepath.Join(root, "internal")

	_ = filepath.Walk(internalPath, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}
		fset := token.NewFileSet()
		fileAst, err := parser.ParseFile(fset, path, nil, 0)
		if err != nil {
			return nil
		}
		conf := types.Config{Importer: importer.Default()}
		infoTypes := &types.Info{
			Types: make(map[ast.Expr]types.TypeAndValue),
			Defs:  make(map[*ast.Ident]types.Object),
		}
		_, _ = conf.Check("pkg", fset, []*ast.File{fileAst}, infoTypes)

		for _, decl := range fileAst.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok || fn.Name.Name != "NewService" {
				continue
			}
			for _, param := range fn.Type.Params.List {
				// Pega o tipo do parâmetro
				typ := infoTypes.TypeOf(param.Type)
				if typ == nil {
					continue
				}
				// Verifica se é interface
				if _, ok := typ.Underlying().(*types.Interface); !ok {
					for _, name := range param.Names {
						errors = append(errors, path+": parâmetro '"+name.Name+"' de NewService não é interface")
					}
				}
				// Verifica se o tipo é externo (fora de internal ou stdlib)
				if named, ok := typ.(*types.Named); ok {
					pkgPath := named.Obj().Pkg()
					if pkgPath != nil {
						importPath := pkgPath.Path()
						if !strings.Contains(importPath, "/internal/") && !isStdLib(importPath) {
							for _, name := range param.Names {
								errors = append(errors, path+": parâmetro '"+name.Name+"' de NewService é de pacote externo: "+importPath)
							}
						}
					}
				}
			}
		}
		return nil
	})
	if len(errors) > 0 {
		return fmt.Errorf("Validação de parâmetros de NewService falhou:\n%s", strings.Join(errors, "\n"))
	}
	return nil
}

// isStdLib verifica se o importPath é da stdlib
func isStdLib(importPath string) bool {
	return !strings.Contains(importPath, ".")
}

// ValidateDbQueriesInRepositories garante que comandos de query no banco só estejam em repositories
func ValidateDbQueriesInRepositories(root string) error {
	var errors []string
	internalPath := filepath.Join(root, "internal")
	queryFuncs := []string{
		"Exec", "Query", "QueryRow", "Raw", "Find", "Where", "Select", "Update", "Delete", "Insert", "Scan", "Rows", "Transaction", "Begin", "Commit", "Rollback",
	}

	_ = filepath.Walk(internalPath, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}
		// Permitir apenas arquivos dentro de /internal/repositories/
		if strings.Contains(path, string(os.PathSeparator)+"repositories"+string(os.PathSeparator)) {
			return nil
		}
		content, err := os.ReadFile(path)
		if err != nil {
			return nil
		}
		code := string(content)
		for _, q := range queryFuncs {
			// Regex: ^|\W q(  => chamada de função simples (não método)
			funcCall := regexp.MustCompile(`(^|\W)` + q + `\s*\(`)
			if funcCall.MatchString(code) {
				// Agora, se for precedido por ponto (método), ignore
				methodCall := regexp.MustCompile(`\w+\.` + q + `\s*\(`)
				if methodCall.MatchString(code) {
					continue // É método, permitido
				}
				errors = append(errors, path+": uso de comando de query ('"+q+"') fora de repositories (chamada direta de função)")
			}
		}
		return nil
	})
	if len(errors) > 0 {
		return fmt.Errorf("Validação de queries no banco falhou:\n%s", strings.Join(errors, "\n"))
	}
	return nil
}

// getLayerAllowedImports lê o arquivo TOML de allowed imports por camada
func getLayerAllowedImports(configPath string) map[string][]string {
	type allowedConfig struct {
		Allowed []string `toml:"allowed"`
	}
	type rootConfig struct {
		Services     allowedConfig `toml:"services"`
		Domain       allowedConfig `toml:"domain"`
		Repositories allowedConfig `toml:"repositories"`
		Providers    allowedConfig `toml:"providers"`
		Context      allowedConfig `toml:"context"`
	}
	layerAllowed := make(map[string][]string)
	file, err := os.ReadFile(configPath)
	if err == nil {
		var conf rootConfig
		err = toml.Unmarshal(file, &conf)
		if err == nil {
			layerAllowed["services"] = conf.Services.Allowed
			layerAllowed["domain"] = conf.Domain.Allowed
			layerAllowed["repositories"] = conf.Repositories.Allowed
			layerAllowed["providers"] = conf.Providers.Allowed
			layerAllowed["context"] = conf.Context.Allowed
			return layerAllowed
		}
	}
	// fallback: regras padrão
	layerAllowed["services"] = []string{"internal/application/domain", "internal/repositories", "internal/application/providers", "internal/application/context"}
	layerAllowed["domain"] = []string{}
	layerAllowed["repositories"] = []string{"internal/application/domain"}
	layerAllowed["providers"] = []string{}
	layerAllowed["context"] = []string{}
	return layerAllowed
}

// Exemplo de função utilitária para formatar erro com linha
func errorWithPos(fset *token.FileSet, node ast.Node, path, msg string) string {
	if node != nil && fset != nil {
		pos := fset.Position(node.Pos())
		return fmt.Sprintf("%s:%d: %s", path, pos.Line, msg)
	}
	return fmt.Sprintf("%s: %s", path, msg)
}

func ValidateLayerDependencies(root string) error {
	var errors []string
	configPath := filepath.Join(root, "tools", "arch_analyser", "layer_allowed_imports.toml")
	layerRules := getLayerAllowedImports(configPath)
	layerPaths := map[string]string{
		"services":     filepath.Join(root, "internal/application/services"),
		"domain":       filepath.Join(root, "internal/application/domain"),
		"repositories": filepath.Join(root, "internal/repositories"),
		"providers":    filepath.Join(root, "internal/application/providers"),
		"context":      filepath.Join(root, "internal/application/context"),
	}

	for layer, allowedImports := range layerRules {
		layerPath := layerPaths[layer]
		filepath.Walk(layerPath, func(path string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() || !strings.HasSuffix(path, ".go") {
				return nil
			}
			fset := token.NewFileSet()
			f, err := parser.ParseFile(fset, path, nil, parser.ImportsOnly)
			if err != nil {
				return nil
			}
			for _, imp := range f.Imports {
				importPath := strings.Trim(imp.Path.Value, "\"")
				if isStdLib(importPath) {
					continue
				}
				allowed := false
				for _, allowedImport := range allowedImports {
					if importPath == allowedImport || strings.Contains(importPath, allowedImport) {
						allowed = true
						break
					}
				}
				if !allowed && len(allowedImports) > 0 {
					errors = append(errors, errorWithPos(fset, imp, path, "importa pacote proibido: "+importPath))
				}
				if len(allowedImports) == 0 && !isStdLib(importPath) {
					errors = append(errors, errorWithPos(fset, imp, path, "importa pacote proibido: "+importPath))
				}
			}
			return nil
		})
	}
	if len(errors) > 0 {
		return fmt.Errorf("Validação de dependências entre camadas falhou:\n%s", strings.Join(errors, "\n"))
	}
	return nil
}

// ValidateNoCircularDependencies garante que não existam ciclos de importação entre pacotes internos
func ValidateNoCircularDependencies(root string) error {
	// Executa o comando go list
	cmd := exec.Command("go", "list", "-json", "-deps", "./internal/...")
	cmd.Dir = root
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("Erro ao executar go list: %v", err)
	}

	type Package struct {
		ImportPath string   `json:"ImportPath"`
		Imports    []string `json:"Imports"`
	}

	dec := json.NewDecoder(strings.NewReader(string(output)))
	graph := make(map[string][]string)
	var pkgs []string
	for {
		var pkg Package
		if err := dec.Decode(&pkg); err != nil {
			break
		}
		if strings.Contains(pkg.ImportPath, "/internal/") {
			graph[pkg.ImportPath] = pkg.Imports
			pkgs = append(pkgs, pkg.ImportPath)
		}
	}

	// Detectar ciclos usando DFS
	visited := make(map[string]bool)
	stack := make(map[string]bool)
	var cycles []string
	var dfs func(string, []string)
	dfs = func(node string, path []string) {
		visited[node] = true
		stack[node] = true
		for _, dep := range graph[node] {
			if !strings.Contains(dep, "/internal/") {
				continue
			}
			if !visited[dep] {
				dfs(dep, append(path, dep))
			} else if stack[dep] {
				cycle := append(path, dep)
				cycles = append(cycles, strings.Join(cycle, " -> "))
			}
		}
		stack[node] = false
	}
	for _, pkg := range pkgs {
		if !visited[pkg] {
			dfs(pkg, []string{pkg})
		}
	}
	if len(cycles) > 0 {
		return fmt.Errorf("Ciclos de dependência detectados:\n%s", strings.Join(cycles, "\n"))
	}
	return nil
}

// ValidateContextUsage garante que apenas services e repositories importem/utilizem context
func ValidateContextUsage(root string) error {
	dirsPermitidos := []string{
		filepath.Join(root, "internal/application/services"),
		filepath.Join(root, "internal/repositories"),
	}
	internalPath := filepath.Join(root, "internal")
	var errors []string
	_ = filepath.Walk(internalPath, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}
		// Se está em services ou repositories, pode usar context
		for _, dir := range dirsPermitidos {
			if strings.HasPrefix(path, dir) {
				return nil
			}
		}
		content, err := os.ReadFile(path)
		if err != nil {
			return nil
		}
		code := string(content)
		if strings.Contains(code, "\"context\"") || strings.Contains(code, "context.Context") {
			errors = append(errors, path+": uso proibido de context/context.Context fora de services ou repositories")
		}
		return nil
	})
	if len(errors) > 0 {
		return fmt.Errorf("Validação de uso de context falhou:\n%s", strings.Join(errors, "\n"))
	}
	return nil
}

// ValidateProvidersArePure garante que providers não importam domain, services ou repositories
func ValidateProvidersArePure(root string) error {
	providersPath := filepath.Join(root, "internal/application/providers")
	var errors []string
	_ = filepath.Walk(providersPath, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, path, nil, parser.ImportsOnly)
		if err != nil {
			return nil
		}
		for _, imp := range f.Imports {
			importPath := strings.Trim(imp.Path.Value, "\"")
			if strings.Contains(importPath, "internal/application/domain") ||
				strings.Contains(importPath, "internal/application/services") ||
				strings.Contains(importPath, "internal/repositories") {
				errors = append(errors, path+": provider importa pacote proibido: "+importPath)
			}
		}
		return nil
	})
	if len(errors) > 0 {
		return fmt.Errorf("Validação de pureza dos providers falhou:\n%s", strings.Join(errors, "\n"))
	}
	return nil
}

// ValidateRepositoryInterfacesImplemented garante que toda interface de repositório em domain tenha implementação em repositories
func ValidateRepositoryInterfacesImplemented(root string) error {
	domainPath := filepath.Join(root, "internal/application/domain")
	repoPath := filepath.Join(root, "internal/repositories")
	var repoInterfaces []string
	var errors []string

	// 1. Encontrar interfaces em domain que terminam com Repository
	_ = filepath.Walk(domainPath, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}
		fset := token.NewFileSet()
		fileAst, err := parser.ParseFile(fset, path, nil, 0)
		if err != nil {
			return nil
		}
		for _, decl := range fileAst.Decls {
			gen, ok := decl.(*ast.GenDecl)
			if !ok || gen.Tok != token.TYPE {
				continue
			}
			for _, spec := range gen.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}
				if strings.HasSuffix(typeSpec.Name.Name, "Repository") && typeSpec.Name.Name != "Repository" {
					repoInterfaces = append(repoInterfaces, typeSpec.Name.Name+"::"+path)
				}
			}
		}
		return nil
	})

	// 2. Para cada interface, buscar implementação em repositories
	for _, ifaceWithPath := range repoInterfaces {
		parts := strings.SplitN(ifaceWithPath, "::", 2)
		iface := parts[0]
		ifacePath := parts[1]
		found := false
		_ = filepath.Walk(repoPath, func(path string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() || !strings.HasSuffix(path, ".go") {
				return nil
			}
			fset := token.NewFileSet()
			fileAst, err := parser.ParseFile(fset, path, nil, 0)
			if err != nil {
				return nil
			}
			conf := types.Config{Importer: importer.Default()}
			infoTypes := &types.Info{
				Types: make(map[ast.Expr]types.TypeAndValue),
				Defs:  make(map[*ast.Ident]types.Object),
			}
			_, _ = conf.Check("pkg", fset, []*ast.File{fileAst}, infoTypes)
			for _, decl := range fileAst.Decls {
				gen, ok := decl.(*ast.GenDecl)
				if !ok || gen.Tok != token.TYPE {
					continue
				}
				for _, spec := range gen.Specs {
					typeSpec, ok := spec.(*ast.TypeSpec)
					if !ok {
						continue
					}
					if _, ok := typeSpec.Type.(*ast.StructType); ok {
						// Verifica se struct implementa a interface (nomes iguais)
						for _, method := range typeSpec.Name.Obj.Decl.(*ast.TypeSpec).Type.(*ast.StructType).Fields.List {
							if method.Names != nil && method.Names[0].Name == iface {
								found = true
								return nil
							}
						}
					}
				}
			}
			return nil
		})
		if !found {
			errors = append(errors, "Interface de repositório sem implementação: '"+iface+"' encontrada em "+ifacePath)
		}
	}
	if len(errors) > 0 {
		return fmt.Errorf("Validação de implementação de interfaces de repositório falhou:\n%s", strings.Join(errors, "\n"))
	}
	return nil
}

// ValidateNoGlobalVars garante que não existam variáveis globais mutáveis fora de escopos controlados
func ValidateNoGlobalVars(root string) error {
	internalPath := filepath.Join(root, "internal")
	var errors []string
	_ = filepath.Walk(internalPath, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}
		fset := token.NewFileSet()
		fileAst, err := parser.ParseFile(fset, path, nil, 0)
		if err != nil {
			return nil
		}
		for _, decl := range fileAst.Decls {
			gen, ok := decl.(*ast.GenDecl)
			if !ok || gen.Tok != token.VAR {
				continue
			}
			for _, spec := range gen.Specs {
				vs, ok := spec.(*ast.ValueSpec)
				if !ok {
					continue
				}
				for _, name := range vs.Names {
					// Ignorar variáveis que começam com Err (padrão Go para erros imutáveis)
					if strings.HasPrefix(name.Name, "Err") {
						continue
					}
					errors = append(errors, errorWithPos(fset, name, path, "declaração de variável global mutável: "+name.Name))
				}
			}
		}
		return nil
	})
	if len(errors) > 0 {
		return fmt.Errorf("Validação de variáveis globais mutáveis falhou:\n%s", strings.Join(errors, "\n"))
	}
	return nil
}

// ValidateOrchestrationOnlyInServices garante que apenas services orquestrem múltiplos domínios ou repositórios
func ValidateOrchestrationOnlyInServices(root string) error {
	internalPath := filepath.Join(root, "internal")
	servicesPath := filepath.Join(root, "internal/application/services")
	var errors []string
	_ = filepath.Walk(internalPath, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}
		// Ignorar arquivos em services
		if strings.HasPrefix(path, servicesPath) {
			return nil
		}
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, path, nil, parser.ImportsOnly)
		if err != nil {
			return nil
		}
		domainImports := []string{}
		repoImports := []string{}
		for _, imp := range f.Imports {
			importPath := strings.Trim(imp.Path.Value, "\"")
			if strings.Contains(importPath, "internal/application/domain") {
				domainImports = append(domainImports, importPath)
			}
			if strings.Contains(importPath, "internal/repositories") {
				repoImports = append(repoImports, importPath)
			}
		}
		// Permitir explicitamente domínio + base_repository
		if len(domainImports) == 1 && len(repoImports) == 1 && strings.Contains(repoImports[0], "base_repository") {
			return nil
		}
		if len(domainImports)+len(repoImports) > 1 {
			errors = append(errors, errorWithPos(fset, f.Imports[0], path, "orquestração de múltiplos domínios/repositórios fora de services (exceto domínio + base_repository)"))
		}
		return nil
	})
	if len(errors) > 0 {
		return fmt.Errorf("Validação de orquestração falhou:\n%s", strings.Join(errors, "\n"))
	}
	return nil
}

// ValidateNamingAndLocation garante que arquivos estejam nas pastas corretas
func ValidateNamingAndLocation(root string) error {
	domainPath := filepath.Join(root, "internal/application/domain")
	servicesPath := filepath.Join(root, "internal/application/services")
	repoPath := filepath.Join(root, "internal/repositories")
	providersPath := filepath.Join(root, "internal/application/providers")
	contextPath := filepath.Join(root, "internal/application/context")
	var errors []string
	internalPath := filepath.Join(root, "internal")
	_ = filepath.Walk(internalPath, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}
		fset := token.NewFileSet()
		fileAst, err := parser.ParseFile(fset, path, nil, 0)
		if err != nil {
			return nil
		}
		for _, decl := range fileAst.Decls {
			// Domínio: struct PascalCase fora de domain
			if gen, ok := decl.(*ast.GenDecl); ok && gen.Tok == token.TYPE {
				for _, spec := range gen.Specs {
					typeSpec, ok := spec.(*ast.TypeSpec)
					if !ok {
						continue
					}
					if _, ok := typeSpec.Type.(*ast.StructType); ok {
						name := typeSpec.Name.Name
						// Verifica se struct implementa os métodos de Domain
						isDomain := false
						for _, decl2 := range fileAst.Decls {
							if fn, ok := decl2.(*ast.FuncDecl); ok && fn.Recv != nil && len(fn.Recv.List) > 0 {
								if starExpr, ok := fn.Recv.List[0].Type.(*ast.StarExpr); ok {
									if ident, ok := starExpr.X.(*ast.Ident); ok && ident.Name == name {
										if fn.Name.Name == "Schema" || fn.Name.Name == "GetFilters" || fn.Name.Name == "SoftDelete" {
											isDomain = true
										}
									}
								}
							}
						}
						if isDomain && !strings.HasPrefix(path, domainPath) {
							errors = append(errors, errorWithPos(fset, typeSpec, path, "struct de domínio fora de internal/application/domain: "+name))
						}
						if strings.HasSuffix(name, "Service") && !strings.HasPrefix(path, servicesPath) {
							errors = append(errors, errorWithPos(fset, typeSpec, path, "Service fora de internal/application/services: "+name))
						}
						if (strings.HasSuffix(name, "Repository") || strings.HasSuffix(name, "Repo")) && !strings.HasPrefix(path, repoPath) {
							errors = append(errors, errorWithPos(fset, typeSpec, path, "Repository fora de internal/repositories: "+name))
						}
						if strings.HasSuffix(name, "Provider") && !strings.HasPrefix(path, providersPath) {
							errors = append(errors, errorWithPos(fset, typeSpec, path, "Provider fora de internal/application/providers: "+name))
						}
						if strings.HasSuffix(name, "Context") && !strings.HasPrefix(path, contextPath) {
							errors = append(errors, errorWithPos(fset, typeSpec, path, "Context fora de internal/application/context: "+name))
						}
					}
				}
			}
			// Funções NewService fora de services
			if fn, ok := decl.(*ast.FuncDecl); ok {
				if fn.Name.Name == "NewService" && !strings.HasPrefix(path, servicesPath) {
					errors = append(errors, errorWithPos(fset, fn, path, "NewService fora de internal/application/services"))
				}
			}
		}
		return nil
	})
	if len(errors) > 0 {
		return fmt.Errorf("Validação de nomenclatura/localização falhou:\n%s", strings.Join(errors, "\n"))
	}
	return nil
}

// isPascalCase verifica se o nome está em PascalCase
func isPascalCase(s string) bool {
	if len(s) == 0 {
		return false
	}
	return strings.ToUpper(s[:1]) == s[:1] && !strings.Contains(s, "_")
}

// ValidateExternalPackagesUsage garante que apenas camadas permitidas importem pacotes externos
func ValidateExternalPackagesUsage(root string) error {
	// Defina aqui os pacotes externos permitidos por camada
	allowed := map[string][]string{
		"repositories": {"github.com/jmoiron/sqlx", "github.com/fatih/structs"},
		"services":     {"github.com/rs/zerolog"},
		// Adicione outros conforme necessário
	}
	layerPaths := map[string]string{
		"repositories": filepath.Join(root, "internal/repositories"),
		"services":     filepath.Join(root, "internal/application/services"),
		"domain":       filepath.Join(root, "internal/application/domain"),
		"providers":    filepath.Join(root, "internal/application/providers"),
		"context":      filepath.Join(root, "internal/application/context"),
	}
	var errors []string
	for layer, path := range layerPaths {
		filepath.Walk(path, func(fpath string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() || !strings.HasSuffix(fpath, ".go") {
				return nil
			}
			fset := token.NewFileSet()
			f, err := parser.ParseFile(fset, fpath, nil, parser.ImportsOnly)
			if err != nil {
				return nil
			}
			for _, imp := range f.Imports {
				importPath := strings.Trim(imp.Path.Value, "\"")
				if isStdLib(importPath) || strings.HasPrefix(importPath, "internal/") || strings.Contains(importPath, "/internal/") {
					continue
				}
				permitido := false
				for _, ext := range allowed[layer] {
					if importPath == ext {
						permitido = true
						break
					}
				}
				if !permitido && len(allowed[layer]) > 0 {
					errors = append(errors, errorWithPos(fset, imp, fpath, "importação de pacote externo não permitido para a camada '"+layer+"': "+importPath))
				}
				if len(allowed[layer]) == 0 && !isStdLib(importPath) && !strings.HasPrefix(importPath, "internal/") && !strings.Contains(importPath, "/internal/") {
					errors = append(errors, errorWithPos(fset, imp, fpath, "nenhum pacote externo permitido para a camada '"+layer+"': "+importPath))
				}
			}
			return nil
		})
	}
	if len(errors) > 0 {
		return fmt.Errorf("Validação de uso de pacotes externos falhou:\n%s", strings.Join(errors, "\n"))
	}
	return nil
}

// ValidateReflectionUsage garante que o uso de reflect seja restrito a camadas utilitárias ou explicitamente permitidas
func ValidateReflectionUsage(root string) error {
	providersPath := filepath.Join(root, "internal/application/providers")
	repoBasePath := filepath.Join(root, "internal/repositories/base_repository")
	internalPath := filepath.Join(root, "internal")
	var errors []string
	_ = filepath.Walk(internalPath, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}
		// Permitir apenas em providers e base_repository
		if strings.HasPrefix(path, providersPath) || strings.HasPrefix(path, repoBasePath) {
			return nil
		}
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, path, nil, parser.AllErrors)
		if err != nil {
			return nil
		}
		for _, imp := range f.Imports {
			importPath := strings.Trim(imp.Path.Value, "\"")
			if importPath == "reflect" {
				errors = append(errors, errorWithPos(fset, imp, path, "importação de reflect não permitida nesta camada"))
			}
		}
		// Busca por uso de reflect.
		content, err := os.ReadFile(path)
		if err == nil && strings.Contains(string(content), "reflect.") {
			errors = append(errors, path+": uso de reflect.* não permitido nesta camada")
		}
		return nil
	})
	if len(errors) > 0 {
		return fmt.Errorf("Validação de uso de reflection falhou:\n%s", strings.Join(errors, "\n"))
	}
	return nil
}

// ValidateTypeAssertions garante que type assertions (.()) só sejam usados em camadas permitidas
func ValidateTypeAssertions(root string) error {
	internalPath := filepath.Join(root, "internal")
	pkgPath := filepath.Join(root, "pkg")
	var errors []string
	checkDir := func(basePath string) {
		_ = filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() || !strings.HasSuffix(path, ".go") {
				return nil
			}
			content, err := os.ReadFile(path)
			if err != nil {
				return nil
			}
			lines := strings.Split(string(content), "\n")
			for i, line := range lines {
				idx := strings.Index(line, ".(")
				if idx != -1 {
					suffix := line[idx+3:]
					if strings.HasPrefix(strings.TrimSpace(suffix), "type") {
						continue
					}
					if strings.Contains(line, ", ok :=") || strings.Contains(line, ",ok :=") {
						continue
					}
					errors = append(errors, fmt.Sprintf("%s:%d: uso de type assertion (.()) não permitido (apenas type switch ou type assertion seguro com vírgula ok)", path, i+1))
				}
			}
			return nil
		})
	}
	checkDir(internalPath)
	checkDir(pkgPath)
	if len(errors) > 0 {
		return fmt.Errorf("Validação de type assertions (.()) falhou:\n%s", strings.Join(errors, "\n"))
	}
	return nil
}
