package core

import (
	"testing"
)

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
