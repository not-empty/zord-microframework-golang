package dummy

import (
	"go-skeleton/pkg/validator"
	"testing"
)

type reqMock struct {
	Validate        func() error
	dummyCreateRule func() error
}

func TestNewRequest(t *testing.T) {
	data := &Data{
		DummyId:   "",
		DummyName: "",
	}
	val := validator.NewValidator()
	NewRequest(data, val)
}

func TestDummyCreateRule(t *testing.T) {
	data := &Data{
		DummyId:   "",
		DummyName: "",
	}
	val := validator.NewValidator()
	req := NewRequest(data, val)
	res := req.dummyCreateRule()
	if res != nil {
		t.Errorf("dummyCreateRule() = %v, want %v", res, nil)
	}
}

func TestValidate(t *testing.T) {
	data := &Data{
		DummyName: "teste",
	}
	val := validator.NewValidator()
	str := NewRequest(data, val)
	mock := &reqMock{}
	mock.dummyCreateRule = func() error {
		return nil
	}
	mock.Validate = str.Validate
	res := mock.Validate()
	if res != nil {
		t.Errorf("Validate() = %v, want %v", res, nil)
	}
}

func TestValidateError(t *testing.T) {
	data := &Data{
		DummyName: "",
	}
	val := validator.NewValidator()
	str := NewRequest(data, val)
	mock := &reqMock{}

	res := mock.Validate()
	if res == nil {
		t.Errorf("Validate() = %v, want %v", res, nil)
	}
}
