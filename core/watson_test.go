package core

import (
	assert "github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestIsObjectAString(t *testing.T) {
	assert.True(t, IsObjectAString("a string"))

	reader, _ := os.Open("foo.txt")
	assert.True(t, IsObjectAReader(reader))

	assert.False(t, IsObjectAString(nil))
	assert.False(t, IsObjectAString(382636))

	assert.False(t, IsObjectAReader(nil))
	assert.False(t, IsObjectAReader("a string"))
}

func TestIsJSONMimeType(t *testing.T) {
	assert.True(t, IsJSONMimeType("application/json"))
	assert.True(t, IsJSONMimeType("APPlication/json"))
	assert.True(t, IsJSONMimeType("application/json;blah"))

	assert.False(t, IsJSONMimeType("application/json-patch+patch"))
	assert.False(t, IsJSONMimeType("YOapplication/jsonYO"))
}

func TestIsJSONPatchMimeType(t *testing.T) {
	assert.True(t, IsJSONPatchMimeType("application/json-patch+json"))
	assert.True(t, IsJSONPatchMimeType("APPlication/json-PATCH+json"))
	assert.True(t, IsJSONPatchMimeType("application/json-patch+json;charset=UTF8"))

	assert.False(t, IsJSONPatchMimeType("application/json"))
	assert.False(t, IsJSONPatchMimeType("YOapplication/json-patch+jsonYO"))
}
