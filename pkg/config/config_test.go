package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	os.Setenv("ENVIRONMENT", "testing")
	c := NewConfig()
	c.LoadEnvs()
	assert.NotNil(t, c)
}

func TestReadConfig(t *testing.T) {
	key := "TEST_KEY"
	value := "test_value"
	os.Setenv(key, value)

	c := NewConfig()
	c.LoadEnvs()
	result := c.ReadConfig(key)

	assert.Equal(t, value, result)
	os.Unsetenv(key)
}

func TestReadNumberConfig(t *testing.T) {
	key := "TEST_NUMBER_KEY"
	value := "42"
	os.Setenv(key, value)

	c := NewConfig()
	c.LoadEnvs()
	result := c.ReadNumberConfig(key)
	expectedInt := int(42)

	assert.Equal(t, result, expectedInt)

	os.Unsetenv(key)
}

func TestReadArrayConfig(t *testing.T) {
	key := "TEST_ARRAY_KEY"
	value := "value1,value2,value3"
	os.Setenv(key, value)

	c := NewConfig()
	c.LoadEnvs()

	result := c.ReadArrayConfig(key)
	expectedResult := []string{"value1", "value2", "value3"}

	assert.Equal(t, len(result), len(expectedResult))

	for i, v := range result {
		assert.Equal(t, v, expectedResult[i])
	}
	os.Unsetenv(key)
}

func TestLoadEnvs(t *testing.T) {
	c := &Config{}
	err := c.LoadEnvs()
	assert.Nil(t, err)
}
