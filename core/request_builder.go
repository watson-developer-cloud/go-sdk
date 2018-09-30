package core

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
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

// common HTTP methods
const (
	POST   = http.MethodPost
	GET    = http.MethodGet
	DELETE = http.MethodDelete
	PUT    = http.MethodPut
	PATCH  = http.MethodPatch
)

// common headers
const (
	Accept               = "Accept"
	ApplicationJSON      = "application/json"
	ContentDisposition   = "Content-Disposition"
	ContentType          = "Content-Type"
	FormURLEncodedHeader = "application/x-www-form-urlencoded"
)

// A FormData stores information for form data
type FormData struct {
	fileName    string
	contentType string
	contents    interface{}
}

// A RequestBuilder is an HTTP request to be sent to the service
type RequestBuilder struct {
	Method string
	URL    *url.URL
	Header http.Header
	Body   io.Reader
	Query  map[string]string
	Form   map[string]FormData
}

// NewRequestBuilder : Initiates a new request
func NewRequestBuilder(method string) *RequestBuilder {
	return &RequestBuilder{
		Method: method,
		Header: make(http.Header),
		Query:  make(map[string]string),
		Form:   make(map[string]FormData),
	}
}

// ConstructHTTPURL creates a properly encoded URL with path parameters.
func (requestBuilder *RequestBuilder) ConstructHTTPURL(endPoint string, pathSegments []string, pathParameters []string) *RequestBuilder {
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
	requestBuilder.URL = u
	return requestBuilder
}

// AddQuery adds Query name and value
func (requestBuilder *RequestBuilder) AddQuery(name string, value string) *RequestBuilder {
	requestBuilder.Query[name] = value
	return requestBuilder
}

// AddHeader adds header name and value
func (requestBuilder *RequestBuilder) AddHeader(name string, value string) *RequestBuilder {
	requestBuilder.Header[name] = []string{value}
	return requestBuilder
}

// AddFormData makes an entry for Form data
func (requestBuilder *RequestBuilder) AddFormData(fieldName string, fileName string, contentType string,
	contents interface{}) *RequestBuilder {
	if fileName == "" {
		if file, ok := contents.(*os.File); ok {
			if !((os.File{}) == *file) { // if file is not empty
				name := filepath.Base(file.Name())
				fileName = name
			}
		}
	}
	requestBuilder.Form[fieldName] = FormData{
		fileName:    fileName,
		contentType: contentType,
		contents:    contents,
	}
	return requestBuilder
}

// SetBodyContentJSON - set the body content from a JSON structure
func (requestBuilder *RequestBuilder) SetBodyContentJSON(bodyContent interface{}) (*RequestBuilder, error) {
	requestBuilder.Body = new(bytes.Buffer)
	err := json.NewEncoder(requestBuilder.Body.(io.Writer)).Encode(bodyContent)
	return requestBuilder, err
}

// SetBodyContentString - set the body content from a string
func (requestBuilder *RequestBuilder) SetBodyContentString(bodyContent string) (*RequestBuilder, error) {
	requestBuilder.Body = strings.NewReader(bodyContent)
	return requestBuilder, nil
}

// SetBodyContentStream - set the body content from an io.Reader instance
func (requestBuilder *RequestBuilder) SetBodyContentStream(bodyContent io.Reader) (*RequestBuilder, error) {
	requestBuilder.Body = bodyContent
	return requestBuilder, nil
}

// CreateMultipartWriter initializes a new multipart writer
func (requestBuilder *RequestBuilder) createMultipartWriter() *multipart.Writer {
	buff := new(bytes.Buffer)
	requestBuilder.Body = buff
	return multipart.NewWriter(buff)
}

// CreateFormFile is a convenience wrapper around CreatePart. It creates
// a new form-data header with the provided field name and file name and contentType
func createFormFile(formWriter *multipart.Writer, fieldname string, filename string, contentType string) (io.Writer, error) {
	h := make(textproto.MIMEHeader)
	contentDisposition := fmt.Sprintf(`form-data; name="%s";`, fieldname)
	if filename != "" {
		contentDisposition += fmt.Sprintf(` filename="%s"`, filename)
	}

	h.Set(ContentDisposition, contentDisposition)
	if contentType != "" {
		h.Set(ContentType, contentType)
	}

	return formWriter.CreatePart(h)
}

// SetBodyContentForMultipart - sets the body content for a part in a multi-part form
func (requestBuilder *RequestBuilder) SetBodyContentForMultipart(contentType string, content interface{}, writer io.Writer) error {
	var err error
	if stream, ok := content.(io.Reader); ok {
		_, err = io.Copy(writer, stream)
	} else if stream, ok := content.(*io.ReadCloser); ok {
		_, err = io.Copy(writer, *stream)
	} else if IsJSONMimeType(contentType) || IsJSONPatchMimeType(contentType) {
		err = json.NewEncoder(writer).Encode(content)
	} else if str, ok := content.(string); ok {
		writer.Write([]byte(str))
	} else if strPtr, ok := content.(*string); ok {
		writer.Write([]byte(*strPtr))
	} else {
		err = fmt.Errorf("Could not decipher the contents")
	}
	return err
}

// Build the request
func (requestBuilder *RequestBuilder) Build() (*http.Request, error) {
	// URL
	URL, err := url.Parse(requestBuilder.URL.String())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Create multipart form data
	if len(requestBuilder.Form) > 0 {
		// handle both application/x-www-form-urlencoded or multipart/form-data
		contentType := requestBuilder.Header.Get(ContentType)
		if contentType == FormURLEncodedHeader {
			data := url.Values{}
			for fieldName, v := range requestBuilder.Form {
				data.Add(fieldName, v.contents.(string))
			}
			requestBuilder.SetBodyContentString(data.Encode())
		} else {
			formWriter := requestBuilder.createMultipartWriter()
			for fieldName, v := range requestBuilder.Form {
				dataPartWriter, err := createFormFile(formWriter, fieldName, v.fileName, v.contentType)
				if err != nil {
					return nil, err
				}
				if err = requestBuilder.SetBodyContentForMultipart(v.contentType,
					v.contents, dataPartWriter); err != nil {
					return nil, err
				}
			}

			requestBuilder.AddHeader("Content-Type", formWriter.FormDataContentType())
			err = formWriter.Close()
			if err != nil {
				return nil, err
			}
		}
	}

	// Create the request
	req, err := http.NewRequest(requestBuilder.Method, URL.String(), requestBuilder.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Headers
	req.Header = requestBuilder.Header

	// Query
	query := req.URL.Query()
	for k, v := range requestBuilder.Query {
		query.Add(k, v)
	}
	// Encode query
	req.URL.RawQuery = query.Encode()

	return req, nil
}

// SetBodyContent - sets the body content from one of three different sources, based on the content type
func (requestBuilder *RequestBuilder) SetBodyContent(contentType string, jsonContent interface{}, jsonPatchContent interface{},
	nonJSONContent interface{}) (*RequestBuilder, error) {
	if contentType != "" {
		if IsJSONMimeType(contentType) {
			if builder, err := requestBuilder.SetBodyContentJSON(jsonContent); err != nil {
				return builder, err
			}
		} else if IsJSONPatchMimeType(contentType) {
			if builder, err := requestBuilder.SetBodyContentJSON(jsonPatchContent); err != nil {
				return builder, err
			}
		} else {
			// Set the non-JSON body content based on the type of value passed in,
			// which should be a "string", "*string" or an "io.Reader"
			if str, ok := nonJSONContent.(string); ok {
				requestBuilder.SetBodyContentString(str)
			} else if strPtr, ok := nonJSONContent.(*string); ok {
				requestBuilder.SetBodyContentString(*strPtr)
			} else if stream, ok := nonJSONContent.(io.Reader); ok {
				requestBuilder.SetBodyContentStream(stream)
			} else if stream, ok := nonJSONContent.(*io.ReadCloser); ok {
				requestBuilder.SetBodyContentStream(*stream)
			} else {
				return requestBuilder, fmt.Errorf("Invalid type for non-JSON body content: %s",
					reflect.TypeOf(nonJSONContent).String())
			}
		}
	} else {
		return requestBuilder, fmt.Errorf("Content-Type cant be empty")
	}

	return requestBuilder, nil
}
