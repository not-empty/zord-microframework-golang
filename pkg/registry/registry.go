package registry

type Registry struct {
	deps map[string]Dependency
}

type Dependency interface{}

func NewRegistry() *Registry {
	return &Registry{
		deps: make(map[string]Dependency),
	}
}

func (r *Registry) Provide(name string, dep Dependency) {
	r.deps[name] = dep
}

func (r *Registry) Inject(name string) Dependency {
	dep, ok := r.deps[name]
	if !ok {
		panic("Invalid injectable")
	}
	return dep
}
