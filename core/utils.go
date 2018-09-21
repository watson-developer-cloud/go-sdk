package core

import (
	"io"
	"regexp"
)

const (
	jSONMimePattern      = "(?i)^application\\/((json)|(merge\\-patch\\+json))(;.*)?$"
	jsonPatchMimePattern = "(?i)^application\\/json\\-patch\\+json(;.*)?$"
)

// StringPtr : return pointer to string literal
func StringPtr(str string) *string {
	return &str
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
	matched, err := regexp.MatchString(jSONMimePattern, mimeType)
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
