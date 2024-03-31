package generator

type Stubs struct {
	ToPath             string
	FromPath           string
	IsGenerated        bool
	UniqueDelete       string
	DeleteRegex        string
	DeleteLinePatterns []string
}

type Config struct {
	Replacers map[string]map[string]string
	Stubs     map[string]map[string]Stubs
}
