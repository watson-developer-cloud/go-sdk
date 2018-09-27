package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"regexp"

	validator "gopkg.in/go-playground/validator.v9"
)

// Validate single instance of Validate, it caches struct info
var Validate *validator.Validate

func init() {
	Validate = validator.New()
}

const (
	jsonMimePattern      = "(?i)^application\\/((json)|(merge\\-patch\\+json))(;.*)?$"
	jsonPatchMimePattern = "(?i)^application\\/json\\-patch\\+json(;.*)?$"
)

// isNil checks if the specified object is nil or not
func isNil(object interface{}) bool {
	if object == nil {
		return true
	}

	value := reflect.ValueOf(object)
	kind := value.Kind()
	if kind >= reflect.Chan && kind <= reflect.Slice && value.IsNil() {
		return true
	}

	return false
}

// ValidateNotNil - returns the specified error if 'object' is nil, nil otherwise
func ValidateNotNil(object interface{}, errorMsg string) error {
	if isNil(object) {
		return errors.New(errorMsg)
	}
	return nil
}

// ValidateStruct - validates 'param' (assumed to be a struct) according to the annotations attached to its fields
func ValidateStruct(param interface{}, paramName string) error {
	if param != nil {
		if err := Validate.Struct(param); err != nil {
			// If there were validation errors then return an error containing the field errors
			if fieldErrors, ok := err.(validator.ValidationErrors); ok {
				return fmt.Errorf("%s failed validation:\n%s", paramName, fieldErrors.Error())
			}
			return fmt.Errorf("An unexpected system error occurred while validating %s\n%s", paramName, err.Error())
		}
	}
	return nil
}

// StringPtr : return pointer to string literal
func StringPtr(literal string) *string {
	return &literal
}

// BoolPtr : return pointer to boolean literal
func BoolPtr(literal bool) *bool {
	return &literal
}

// Int64Ptr : return pointer to int64 literal
func Int64Ptr(literal int64) *int64 {
	return &literal
}

// IsJSONMimeType : Returns true iff the specified mimeType value represents a "JSON" mimetype.
func IsJSONMimeType(mimeType string) bool {
	if mimeType == "" {
		return false
	}
	matched, err := regexp.MatchString(jsonMimePattern, mimeType)
	if err != nil {
		return false
	}
	return matched
}

// IsJSONPatchMimeType : Returns true iff the specified mimeType value represents a "JSON Patch" mimetype.
func IsJSONPatchMimeType(mimeType string) bool {
	if mimeType == "" {
		return false
	}
	matched, err := regexp.MatchString(jsonPatchMimePattern, mimeType)
	if err != nil {
		return false
	}
	return matched
}

// StringNilMapper - de-references the parameter 's' and returns the result, or "" if 's' is nil
func StringNilMapper(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// PrettyPrint print pretty
func PrettyPrint(result interface{}, resultName string) {
	output, err := json.MarshalIndent(result, "", "    ")

	if err == nil {
		fmt.Printf("%v:\n%+v\n\n", resultName, string(output))
	}
}
