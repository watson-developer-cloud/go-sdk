package core

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Foo struct {
	Name *string `json:"name,omitempty"`
}

func TestRequestResponseAsJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		fmt.Fprintf(w, `{"name": "wonder woman"}`)
	}))
	defer server.Close()

	builder := NewRequestBuilder("GET").
		ConstructHTTPURL(server.URL, nil, nil).
		AddHeader("Content-Type", "Application/json").
		AddQuery("Version", "2018-22-09")
	req, _ := builder.Build()

	options := &ServiceOptions{
		URL:      server.URL,
		Username: "xxx",
		Password: "yyy",
	}
	service, _ := NewWatsonService(options, "watson", "watson")
	detailedResponse, _ := service.Request(req, new(Foo))
	assert.Equal(t, "wonder woman", *detailedResponse.Result.(*Foo).Name)
}

func TestIncorrectCreds(t *testing.T) {
	options := &ServiceOptions{
		URL:      "xxx",
		Username: "{yyy}",
		Password: "zzz",
	}
	_, serviceErr := NewWatsonService(options, "watson", "watson")
	assert.Equal(t, "The username shouldn't start or end with curly brackets or quotes. Be sure to remove any {} and \" characters surrounding your username", serviceErr.Error())
}

func TestIncorrectURL(t *testing.T) {
	options := &ServiceOptions{
		URL:      "{xxx}",
		Username: "yyy",
		Password: "zzz",
	}
	_, serviceErr := NewWatsonService(options, "watson", "watson")
	assert.Equal(t, "The URL shouldn't start or end with curly brackets or quotes. Be sure to remove any {} and \" characters surrounding your URL", serviceErr.Error())
}

func TestDisableSSLverification(t *testing.T) {
	options := &ServiceOptions{
		URL:      "test.com",
		Username: "xxx",
		Password: "yyy",
	}
	service, _ := NewWatsonService(options, "watson", "watson")
	assert.Nil(t, service.Client.Transport)
	service.DisableSSLVerification()
	assert.NotNil(t, service.Client.Transport)
}

func TestAuthentication(t *testing.T) {
	encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte("xxx:yyy"))
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		assert.Equal(t, "Basic "+encodedBasicAuth, r.Header["Authorization"][0])
	}))
	defer server.Close()

	builder := NewRequestBuilder("GET").
		ConstructHTTPURL(server.URL, nil, nil).
		AddQuery("Version", "2018-22-09")
	req, _ := builder.Build()

	options := &ServiceOptions{
		URL:      server.URL,
		Username: "xxx",
		Password: "yyy",
	}
	service, _ := NewWatsonService(options, "watson", "watson")

	service.Request(req, new(Foo))
}

func TestLoadingFromCredentialFile(t *testing.T) {
	pwd, _ := os.Getwd()
	credentialFilePath := path.Join(pwd, "/../resources/ibm-credentials.env")
	os.Setenv("IBM_CREDENTIALS_FILE", credentialFilePath)
	options := &ServiceOptions{}
	service, _ := NewWatsonService(options, "watson", "watson")
	assert.Equal(t, service.Options.IAMApiKey, "5678efgh")
	os.Unsetenv("IBM_CREDENTIALS_FILE")

	options2 := &ServiceOptions{IAMApiKey: "xxx"}
	service2, _ := NewWatsonService(options2, "watson", "watson")
	assert.Equal(t, service2.Options.IAMApiKey, "xxx")
}

func TestICPAuthentication(t *testing.T) {
	options := &ServiceOptions{
		IAMApiKey: "xxx",
	}
	service, _ := NewWatsonService(options, "watson", "watson")
	assert.Equal(t, "xxx", service.Options.IAMApiKey)

	options2 := &ServiceOptions{
		IAMApiKey: "icp-xxx",
	}
	service2, _ := NewWatsonService(options2, "watson", "watson")
	assert.Equal(t, "icp-xxx", service2.Options.Password)
}
