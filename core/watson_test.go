package core

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httptest"
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
	service, _ := NewWatsonService(options, "watson")
	detailedResponse, _ := service.Request(req, new(Foo))
	assert.Equal(t, "wonder woman", *detailedResponse.Result.(*Foo).Name)
}

func TestDisableSSLverification(t *testing.T) {
	options := &ServiceOptions{
		URL:      "test.com",
		Username: "xxx",
		Password: "yyy",
	}
	service, _ := NewWatsonService(options, "watson")
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
	service, _ := NewWatsonService(options, "watson")

	service.Request(req, new(Foo))
}
