package core

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
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
func TestValidateStruct(t *testing.T) {
	type Address struct {
		Street string `validate:"required"`
		City   string `validate:"required"`
	}

	type User struct {
		FirstName *string    `json:"fname" validate:"required"`
		LastName  *string    `json:"lname" validate:"required"`
		Addresses []*Address `json:"address" validate:"dive"`
	}

	address := &Address{
		Street: "Eavesdown Docks",
		City:   "",
	}

	firstName := "Blossom"
	lastName := "Powerpuff"
	user := &User{
		FirstName: &firstName,
		LastName:  &lastName,
		Addresses: []*Address{address},
	}

	err := Validate.Struct(user)

	if err == nil {
		t.Errorf("Validator is incorrect, it should print an error")
	}
}
