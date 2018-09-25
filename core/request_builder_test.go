package core

import (
	"bytes"
	"os"
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func setup() *RequestBuilder {
	return NewRequestBuilder("GET")
}

func TestNewRequestBuilder(t *testing.T) {
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
	testStructure := &TestStructure{
		Name: "wonder woman",
	}
	body := make(map[string]interface{})
	body["name"] = testStructure.Name
	want := []byte(`{"name":"wonder woman"}`)

	request := setup()
	request.SetBodyContentJSON(body)
	buff := make([]byte, 23)
	request.Body.Read(buff)

	if !bytes.Equal(want, buff) {
		t.Errorf("Couldnt serialize")
	}
}

func TestSetBodyContentString(t *testing.T) {
	request := setup()
	request.SetBodyContentString("hello GO SDK")

	if request.Body == nil {
		t.Errorf("Couldnt set content type as string")
	}
}

func TestBuildWithMultipartFormEmptyFileName(t *testing.T) {
	request := NewRequestBuilder("POST").
		ConstructHTTPURL("test.com", nil, nil).
		AddHeader("Content-Type", "Application/json").
		AddQuery("Version", "2018-22-09").
		AddFormData("hello1", "", "text/plain", "Hello GO SDK").
		AddFormData("hello2", "", "", "Hello GO SDK again")

	req, _ := request.Build()
	if req.Body == nil {
		t.Errorf("Couldnt build successfully")
	}
}

func TestBuildWithMultipartForm(t *testing.T) {
	json1 := make(map[string]interface{})
	json1["name1"] = "test name1"

	json2 := make(map[string]interface{})
	json2["name2"] = "test name2"

	request := NewRequestBuilder("POST").
		ConstructHTTPURL("test.com", nil, nil).
		AddHeader("Content-Type", "Application/json").
		AddQuery("Version", "2018-22-09").
		AddFormData("name1", "json1.json", "application/json", json1).
		AddFormData("name2", "json2.json", "application/json", json2).
		AddFormData("hello", "", "text/plain", "Hello GO SDK")

	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/resources/personality-v3.txt")
	if err != nil {
		t.Errorf("Could not open file")
	}
	request.AddFormData("personality", "personality.txt", "application/octet-stream", file)

	_, err = request.Build()
	if err != nil {
		t.Errorf("Couldnt build successfully")
	}
	defer file.Close()
}

func TestBuild(t *testing.T) {
	endPoint := "https://gateway.watsonplatform.net/assistant/api"
	pathSegments := []string{"v1/workspaces", "message"}
	pathParameters := []string{"xxxxx"}
	wantURL := "https://gateway.watsonplatform.net/assistant/api/xxxxx/v1/workspaces?Version=2018-22-09"

	testStructure := &TestStructure{
		Name: "wonder woman",
	}
	body := make(map[string]interface{})
	body["name"] = testStructure.Name

	request := NewRequestBuilder("POST").
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
