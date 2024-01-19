package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var unicodeVar rune = '日'
var testStruct = struct {
	Alpha           string `json:"alpha" validate:"alpha"`
	AlphaNum        string `json:"alphanum" validate:"alphanum"`
	AlphaNumUnicode string `json:"alphanumunicode" validate:"alphanumunicode"`
	AlphaUnicode    string `json:"alphaunicode" validate:"alphaunicode"`
	Boolean         string `json:"boolean" validate:"boolean"`
	Datetime        string `json:"datetime" validate:"datetime=2006-01-02"`
	Email           string `json:"email" validate:"email"`
	File            string `json:"file" validate:"file"`
	Gt              int    `json:"gt" validate:"gt=10"`
	Gte             int    `json:"gte" validate:"gte=10"`
	IntOrFloat      int    `json:"number" validate:"number"`
	IP              string `json:"ip" validate:"ip"`
	IPv4            string `json:"ipv4" validate:"ipv4"`
	IPv6            string `json:"ipv6" validate:"ipv6"`
	JSON            string `json:"json" validate:"json"`
	Lt              int    `json:"lt" validate:"lt=10"`
	Lte             int    `json:"lte" validate:"lte=10"`
	Max             int    `json:"max" validate:"max=10"`
	Min             int    `json:"min" validate:"min=10"`
	Numeric         string `json:"numeric" validate:"numeric"`
	Required        string `json:"required" validate:"required"`
	RequiredIf      string `json:"RequiredIfF" validate:"required_if=Required john_doe"`
	RequiredUnless  string `json:"required_unless" validate:"required_unless=RequiredIf 1"`
	RequiredWith    string `json:"required_with" validate:"required_with"`
	// RequiredWithAll string `json:"required_with_all" validate:"required_with_all=RequiredWith"` Todo: entender funcionamento
	// RequiredWithout    string `json:"required_without" validate:"required_without"`
	// RequiredWithoutAll string `json:"required_without_all" validate:"required_without_all"`
	// String             string `json:"string" validate:"string"`
	// Timezone string `json:"timezone" validate:"timezone"`
	// Unique   string `json:"unique" validate:"unique"`
	// URL      string `json:"url" validate:"url"`
	// ULID     string `json:"ulid" validate:"ulid"`
}{
	Alpha:           "abc",
	AlphaNum:        "abc123",
	AlphaNumUnicode: "abc123" + string(unicodeVar),
	AlphaUnicode:    "abc" + string(unicodeVar),
	Boolean:         "true",
	Datetime:        "2024-01-19",
	Email:           "test@gmail.com",
	File:            "validator.go",
	Gt:              11,
	Gte:             10,
	IntOrFloat:      10,
	IP:              "191.161.1.17",
	IPv4:            "191.161.1.17",
	IPv6:            "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
	JSON:            "{\"name\": \"Test\"}",
	Lt:              9,
	Lte:             10,
	Max:             9,
	Min:             11,
	Numeric:         "10",
	Required:        "john_doe",
	RequiredIf:      "test",
	RequiredUnless:  "test",
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
