package core

import (
	"io"
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

// IsObjectAReader : Returns true iff "obj" represents an instance of "io.ReadCloser"
func IsObjectAReader(obj interface{}) bool {
	_, ok := obj.(io.Reader)
	return ok
}

// IsObjectAString : Returns true iff "obj" is an instance of a "string"
func IsObjectAString(obj interface{}) bool {
	_, ok := obj.(string)
	return ok
}
