package generator

type Stubs struct {
	ToPath      string
	FromPath    string
	IsDeletable bool
}

type Config struct {
	Replacers map[string]map[string]string
	Stubs     map[string]map[string]Stubs
}
