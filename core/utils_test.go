package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestStringNilMapper(t *testing.T) {
	var s = "test string"
	assert.Equal(t, "", StringNilMapper(nil))
	assert.Equal(t, "test string", StringNilMapper(&s))
}

func TestValidateNotNil(t *testing.T) {
	var str *string
	assert.Nil(t, str)
	err := ValidateNotNil(str, "str should not be nil!")
	assert.NotNil(t, err, "Should have gotten an error for nil 'str' pointer")
	msg := err.Error()
	assert.Equal(t, "str should not be nil!", msg)

	type MyOperationOptions struct {
		Parameter1 *string
	}

	var options *MyOperationOptions
	assert.Nil(t, options, "options should be nil!")
	err = ValidateNotNil(options, "options param should not be nil")
	assert.NotNil(t, err, "Should have gotten an error for nil 'y' ptr")
	msg = err.Error()
	assert.Equal(t, "options param should not be nil", msg)
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

	goodStruct := &Address{
		Street: "Beltorre Drive",
		City:   "Georgetown, TX",
	}

	badStruct := &Address{
		Street: "Beltorre Drive",
	}

	assert.NotNil(t, ValidateStruct(user, "userPtr"), "Should have a validation error!")
	assert.Nil(t, ValidateStruct(nil, "nil ptr"), "nil pointer should validate cleanly!")
	assert.Nil(t, ValidateStruct(goodStruct, "goodStruct"), "Should not cause a validation error!")
	err := ValidateStruct(badStruct, "badStruct")
	assert.NotNil(t, err, "Should have a validation error!")
}

func TestCreateFormPartName(t *testing.T) {
	assert.Equal(t, "class1_positive_examples", CreateFormPartName("{classname}_positive_examples", "classname", "class1"))
	assert.Equal(t, "class2_positive_examples", CreateFormPartName("{classname}_positive_examples", "classname", "class2"))
	assert.Equal(t, "this_is_a_test", CreateFormPartName("this_is_{word}_test", "word", "a"))
	assert.Equal(t, "this_is_not_a_test", CreateFormPartName("this_is_{word}_test", "word", "not_a"))
}
