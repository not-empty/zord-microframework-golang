package registry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRegistry(t *testing.T) {
	r := NewRegistry()
	assert.NotNil(t, r)
	assert.NotNil(t, r.deps)
}

func TestProvideAndInject(t *testing.T) {
	r := NewRegistry()

	dep := "sampleDependency"
	r.Provide("sample", dep)

	injectedDep := r.Inject("sample")
	assert.Equal(t, dep, injectedDep)

	defer func() {
		r := recover()
		assert.Equal(t, r, "Invalid injectable")
	}()
	r.Inject("nonexistent")
}
