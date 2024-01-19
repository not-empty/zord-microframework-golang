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

func TestDepsTypes(t *testing.T) {
	r := NewRegistry()

	type SampleStruct struct{}
	r.Provide("sample", SampleStruct{})

	assert.IsType(t, SampleStruct{}, r.Inject("sample"))
}

func TestProvideAndInject(t *testing.T) {
	r := NewRegistry()

	dep := "sampleDependency"
	r.Provide("sample", dep)

	injectedDep := r.Inject("sample")
	assert.Equal(t, dep, injectedDep)

	assert.Panics(t, func() { r.Inject("nonexistent") }, "Invalid injectable")
}
