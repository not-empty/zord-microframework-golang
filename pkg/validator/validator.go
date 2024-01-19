package validator

import (
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validationMessages = map[string]string{
	"alpha":                "The :attribute may only contain letters.",
	"alphanum":             "The :attribute may only contain letters and numbers.",
	"alphanumunicode":      "The :attribute may only contain letters, numbers and unicode characters.",
	"alphaunicode":         "The :attribute may only contain letters and unicode characters.",
	"boolean":              "The :attribute field must be true or false.",
	"datetime":             "The :attribute is not a valid datetime.",
	"email":                "The :attribute must be a valid email address.",
	"file":                 "The :attribute must be a file.",
	"gt":                   "The :attribute must be greater than :values.",
	"gte":                  "The :attribute must be greater than or equal :values.",
	"integer":              "The :attribute must be an integer.",
	"ip":                   "The :attribute must be a valid IP address.",
	"ipv4":                 "The :attribute must be a valid IPv4 address.",
	"ipv6":                 "The :attribute must be a valid IPv6 address.",
	"json":                 "The :attribute must be a valid JSON string.",
	"lt":                   "The :attribute must be less than :values.",
	"lte":                  "The :attribute must be less than or equal :values.",
	"max":                  "The :attribute may not be greater than :values.",
	"min":                  "The :attribute must be at least :values.",
	"numeric":              "The :attribute must be a number.",
	"required":             "The :attribute field is required.",
	"required_if":          "The :attribute field is required when :other is :values.",
	"required_unless":      "The :attribute field is required unless :other is in :values.",
	"required_with":        "The :attribute field is required when :values is present.",
	"required_with_all":    "The :attribute field is required when :values is present.",
	"required_without":     "The :attribute field is required when :values is not present.",
	"required_without_all": "The :attribute field is required when none of :values are present.",
	"string":               "The :attribute must be a string.",
	"timezone":             "The :attribute must be a valid zone.",
	"unique":               "The :attribute has already been taken.",
	"url":                  "The :attribute format is invalid.",
	"ulid":                 "The :attribute must be a valid ULID string.",
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

type Validator struct {
	validate           *validator.Validate
	validationMessages map[string]string
}

func NewValidator() *Validator {
	return &Validator{
		validate:           validator.New(),
		validationMessages: validationMessages,
	}
}

func (v *Validator) Boot() {
	v.validate.RegisterTagNameFunc(tagNameFunc)
}

func (v *Validator) ValidateStruct(modelData any) []error {
	var errorsList []error

	err := v.validate.Struct(modelData)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.Field()
			element.Tag = err.Tag()
			element.Value = err.Param()

			message := v.translateError(&element)
			errorsList = append(errorsList, message)
		}
	}

	return errorsList
}

func (v *Validator) translateError(errorData *ErrorResponse) error {
	var message string = "The :attribute with :values is not valid."

	if v, ok := v.validationMessages[errorData.Tag]; ok {
		message = v
	}

	message = strings.Replace(message, ":attribute", errorData.FailedField, 1)
	message = strings.Replace(message, ":values", errorData.Value, 1)
	erro := errors.New(message)
	return erro
}

func tagNameFunc(fld reflect.StructField) string {
	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

	if name == "-" {
		return ""
	}

	return name
}
