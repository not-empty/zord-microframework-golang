package idCreator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	idGenerator := NewIdCreator()
	ulidString := idGenerator.Create()
	assert.NotEmpty(t, ulidString)
}

func TestCreateUnique(t *testing.T) {
	ulidMap := make(map[string]bool)

	for i := 0; i < 1000; i++ {
		idGenerator := NewIdCreator()
		ulidString := idGenerator.Create()

		_, exists := ulidMap[ulidString]
		assert.False(t, exists)
	}
}
