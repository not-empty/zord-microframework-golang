package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Todo: implement all validation cases
var unicodeVar rune = '日'
var testStruct = struct {
	Alpha           string `json:"alpha" validate:"alpha"`
	AlphaNum        string `json:"alphanum" validate:"alphanum"`
	AlphaNumUnicode string `json:"alphanumunicode" validate:"alphanumunicode"`
	// Datetime        string `json:"datetime" validate:"datetime"`
	Boolean      string `json:"boolean" validate:"boolean"`
	AlphaUnicode string `json:"alphaunicode" validate:"alphaunicode"`
	Required     string `json:"required" validate:"required"`
}{
	Alpha:           "abc",
	AlphaNum:        "abc123",
	AlphaNumUnicode: "abc123" + string(unicodeVar),
	AlphaUnicode:    "abc" + string(unicodeVar),
	Boolean:         "true",
	// Datetime:        "2020-01-02T00:00:00.0000-03:00",
	Required: "john_doe",
}

func TestValidator_ValidateStruct_ValidData(t *testing.T) {
	validator := NewValidator()
	validator.Boot()

	d := testStruct
	errorsList := validator.ValidateStruct(d)

	assert.Empty(t, errorsList)
}

func TestValidator_ValidateStruct_InvalidData(t *testing.T) {
	validator := NewValidator()

	d := testStruct
	d.Required = ""
	errorsList := validator.ValidateStruct(d)

	assert.NotEmpty(t, errorsList)
}

func TestValidator_translateError(t *testing.T) {
	errorData := &ErrorResponse{
		FailedField: "username",
		Tag:         "required",
		Value:       "",
	}
	validator := NewValidator()

	err := validator.translateError(errorData)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "The username field is required")
}
