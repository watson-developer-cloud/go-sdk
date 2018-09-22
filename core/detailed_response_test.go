package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStructure struct {
	Name string `json:"name"`
}

func TestDetailedResponse(t *testing.T) {
	testStructure := TestStructure{
		Name: "wonder woman",
	}
	response := &DetailedResponse{
		StatusCode: 200,
		Result:     testStructure,
	}
	assert.Equal(t, response.GetResult(), testStructure)
	assert.Equal(t, response.GetStatusCode(), 200)
}
