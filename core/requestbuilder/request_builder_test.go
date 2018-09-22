package requestbuilder

import (
	"bytes"
	"testing"

	assert "github.com/stretchr/testify/assert"
)

type TestStructure struct {
	Name *string `json:"name"`
}

func setup() *Request {
	return NewRequest("GET")
}

func TestNewRequest(t *testing.T) {
	request := setup()
	if request.Method != "GET" {
		t.Errorf("Got incorrect method types")
	}
}

func TestConstructHTTPURL(t *testing.T) {
	endPoint := "https://gateway.watsonplatform.net/assistant/api"
	pathSegments := []string{"v1/workspaces", "message"}
	pathParameters := []string{"xxxxx"}
	request := setup()
	want := "https://gateway.watsonplatform.net/assistant/api/v1/workspaces/xxxxx/message"
	request.ConstructHTTPURL(endPoint, pathSegments, pathParameters)

	if request.URL.String() != want {
		t.Errorf("Invalid comstruction of url")
	}
}

func TestConstructHTTPURLWithNoPathParam(t *testing.T) {
	endPoint := "https://gateway.watsonplatform.net/assistant/api"
	pathSegments := []string{"v1/workspaces"}
	request := setup()
	want := "https://gateway.watsonplatform.net/assistant/api/v1/workspaces"
	request.ConstructHTTPURL(endPoint, pathSegments, nil)

	if request.URL.String() != want {
		t.Errorf("Invalid comstruction of url")
	}
}

func TestAddQuery(t *testing.T) {
	request := setup()
	request.AddQuery("VERSION", "2018-22-09")

	if len(request.Query) != 1 {
		t.Errorf("Didnt set the query pair")
	}
}

func TestAddHeader(t *testing.T) {
	request := setup()
	request.AddHeader("Content-Type", "application/json")

	if len(request.Header) != 1 {
		t.Errorf("Didnt set the header pair")
	}
}

func TestSetBodyContentJSON(t *testing.T) {
	name := "wonder woman"
	testStructure := &TestStructure{
		Name: &name,
	}
	body := make(map[string]interface{})
	body["name"] = testStructure.Name
	want := []byte(`{"name":"wonder woman"}`)

	request := setup()
	request.SetBodyContentJSON(body)
	buff := make([]byte, 19)
	request.Body.Read(buff)

	if !bytes.Equal(want, buff) {
		t.Errorf("Couldnt serialize")
	}
}

func TestBuild(t *testing.T) {
	endPoint := "https://gateway.watsonplatform.net/assistant/api"
	pathSegments := []string{"v1/workspaces", "message"}
	pathParameters := []string{"xxxxx"}
	wantURL := "https://gateway.watsonplatform.net/assistant/api/xxxxx/v1/workspaces?Version=2018-22-09"

	name := "wonder woman"
	testStructure := &TestStructure{
		Name: &name,
	}
	body := make(map[string]interface{})
	body["name"] = testStructure.Name

	request := NewRequest("POST").
		ConstructHTTPURL(endPoint, pathParameters, pathSegments).
		AddHeader("Content-Type", "Application/json").
		AddQuery("Version", "2018-22-09")

	request, _ = request.SetBodyContentJSON(body)
	req, err := request.Build()
	if err != nil {
		t.Errorf("Couldnt build successfully")
	}

	assert.Equal(t, req.URL.String(), wantURL)
	assert.Equal(t, req.Header["Content-Type"][0], "Application/json")
}
