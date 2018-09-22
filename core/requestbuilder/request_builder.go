package requestbuilder

/**
 * Copyright 2018 IBM All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

// HTTP methods
const (
	POST   = "POST"
	GET    = "GET"
	DELETE = "DELETE"
	PUT    = "PUT"
)

const (
	// ACCEPT header
	ACCEPT = "Accept"
	// APPLICATIONJSON header value
	APPLICATIONJSON = "application/json"
)

// A Request is an HTTP request to be sent to the service
type Request struct {
	Method string
	URL    *url.URL
	Header http.Header
	Body   io.Reader
	Query  map[string]string
}

// NewRequest : Initiates a new request
func NewRequest(method string) *Request {
	return &Request{
		Method: method,
		Header: make(http.Header),
		Query:  make(map[string]string),
	}
}

// ConstructHTTPURL creates a properly encoded URL with path parameters.
func (request *Request) ConstructHTTPURL(endPoint string, pathSegments []string, pathParameters []string) *Request {
	for i, pathSegment := range pathSegments {
		endPoint += "/" + pathSegment
		if pathParameters != nil && i < len(pathParameters) {
			endPoint += "/" + pathParameters[i]
		}
	}
	u, err := url.Parse(endPoint)
	if err != nil {
		panic(err)
	}
	request.URL = u
	return request
}

// AddQuery adds Query name and value
func (request *Request) AddQuery(name string, value string) *Request {
	request.Query[name] = value
	return request
}

// AddHeader adds header name and value
func (request *Request) AddHeader(name string, value string) *Request {
	request.Header[name] = []string{value}
	return request
}

// SetBodyContentJSON - set the body content from a JSON structure
func (request *Request) SetBodyContentJSON(bodyContent interface{}) (*Request, error) {
	request.Body = new(bytes.Buffer)
	err := json.NewEncoder(request.Body.(io.Writer)).Encode(bodyContent)
	return request, err
}

// Build the request
func (request *Request) Build() (*http.Request, error) {
	// URL
	url, err := url.Parse(request.URL.String())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// Create the request
	req, err := http.NewRequest(request.Method, url.String(), request.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// TODO: Take care of files

	// Headers
	req.Header = request.Header

	// Query
	query := req.URL.Query()
	for k, v := range request.Query {
		query.Add(k, v)
	}
	// Encode query
	req.URL.RawQuery = query.Encode()

	return req, nil
}

// TODO: Need to uncomment and write tests for below
// // SetBodyContentStream - set the body content from an io.Reader instance
// func (request *Request) SetBodyContentStream(bodyContent io.Reader) error {
// 	service.body = bodyContent
// 	return nil
// }

// // SetBodyContentString - set the body content from a string
// func (request *Request) SetBodyContentString(bodyContent string) error {
// 	service.body = strings.NewReader(bodyContent)
// 	return nil
// }

// // SetBodyContent - sets the body content from one of three different sources, based on the content type
// func (request *Request) SetBodyContent(contentType string, jsonContent interface{}, jsonPatchContent interface{},
// 	nonJSONContent interface{}) error {
// 	if contentType != "" {
// 		if IsJSONMimeType(contentType) {
// 			err := service.SetBodyContentJSON(jsonContent)
// 			if err != nil {
// 				return err
// 			}
// 		} else if IsJSONPatchMimeType(contentType) {
// 			err := service.SetBodyContentJSON(jsonPatchContent)
// 			if err != nil {
// 				return err
// 			}
// 		} else {
// 			// Set the non-JSON body content based on the type of value passed in,
// 			// which should be either a "string" or an "io.Reader"
// 			if IsObjectAString(nonJSONContent) {
// 				service.SetBodyContentString(nonJSONContent.(string))
// 			} else if IsObjectAReader(nonJSONContent) {
// 				service.SetBodyContentStream(nonJSONContent.(io.Reader))
// 			} else {
// 				return fmt.Errorf("Invalid type for non-JSON body content: %s", reflect.TypeOf(nonJSONContent).String())
// 			}
// 		}
// 	}

// 	return nil
// }
