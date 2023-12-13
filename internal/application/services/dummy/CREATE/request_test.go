package dummy

import (
	"go-skeleton/pkg/validator"
	"testing"
)

func TestNewRequest(t *testing.T) {
	data := &Data{
		DummyId:   "",
		DummyName: "",
	}
	val := validator.NewValidator()
	NewRequest(data, val)
}

func TestValidate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		data := &Data{
			DummyName: "teste",
		}
		val := validator.NewValidator()
		str := NewRequest(data, val)
		res := str.Validate()
		if res != nil {
			t.Errorf("Validate() = %v, want %v", res, nil)
		}
	})

	t.Run("dummyCreateError", func(t *testing.T) {
		data := &Data{
			DummyName: "samuel",
		}
		val := validator.NewValidator()
		str := NewRequest(data, val)
		res := str.Validate()
		if res == nil {
			t.Errorf("Validate() = %v, want %v", res, nil)
		}
	})

	t.Run("validatorError", func(t *testing.T) {
		data := &Data{
			DummyName: "",
		}
		val := validator.NewValidator()
		str := NewRequest(data, val)
		res := str.Validate()
		if res == nil {
			t.Errorf("Validate() = %v, want %v", res, nil)
		}
	})
}
