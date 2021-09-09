/**
 * (C) Copyright IBM Corp. 2021.
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

package comparecomplyv1_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/v2/comparecomplyv1"
)

var _ = Describe(`CompareComplyV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		version := "testString"
		It(`Instantiate service client`, func() {
			compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       core.StringPtr(version),
			})
			Expect(compareComplyService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
			Expect(compareComplyService.Service.IsSSLDisabled()).To(BeFalse())
			compareComplyService.DisableSSLVerification()
			Expect(compareComplyService.Service.IsSSLDisabled()).To(BeTrue())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(compareComplyService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
				URL:     "https://comparecomplyv1/api",
				Version: core.StringPtr(version),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(compareComplyService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{})
			Expect(compareComplyService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"COMPARE_COMPLY_URL":       "https://comparecomplyv1/api",
				"COMPARE_COMPLY_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					Version: core.StringPtr(version),
				})
				Expect(compareComplyService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := compareComplyService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != compareComplyService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(compareComplyService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(compareComplyService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:     "https://testService/api",
					Version: core.StringPtr(version),
				})
				Expect(compareComplyService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := compareComplyService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != compareComplyService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(compareComplyService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(compareComplyService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					Version: core.StringPtr(version),
				})
				err := compareComplyService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := compareComplyService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != compareComplyService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(compareComplyService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(compareComplyService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"COMPARE_COMPLY_URL":       "https://comparecomplyv1/api",
				"COMPARE_COMPLY_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(compareComplyService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"COMPARE_COMPLY_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(compareComplyService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = comparecomplyv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ConvertToHTML(convertToHTMLOptions *ConvertToHTMLOptions) - Operation response error`, func() {
		version := "testString"
		convertToHTMLPath := "/v1/html_conversion"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(convertToHTMLPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["model"]).To(Equal([]string{"contracts"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ConvertToHTML with error: Operation response processing error`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Construct an instance of the ConvertToHTMLOptions model
				convertToHTMLOptionsModel := new(comparecomplyv1.ConvertToHTMLOptions)
				convertToHTMLOptionsModel.File = CreateMockReader("This is a mock file.")
				convertToHTMLOptionsModel.FileContentType = core.StringPtr("application/pdf")
				convertToHTMLOptionsModel.Model = core.StringPtr("contracts")
				convertToHTMLOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := compareComplyService.ConvertToHTML(convertToHTMLOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				compareComplyService.EnableRetries(0, 0)
				result, response, operationErr = compareComplyService.ConvertToHTML(convertToHTMLOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ConvertToHTML(convertToHTMLOptions *ConvertToHTMLOptions)`, func() {
		version := "testString"
		convertToHTMLPath := "/v1/html_conversion"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(convertToHTMLPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["model"]).To(Equal([]string{"contracts"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"num_pages": "NumPages", "author": "Author", "publication_date": "PublicationDate", "title": "Title", "html": "HTML"}`)
				}))
			})
			It(`Invoke ConvertToHTML successfully with retries`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())
				compareComplyService.EnableRetries(0, 0)

				// Construct an instance of the ConvertToHTMLOptions model
				convertToHTMLOptionsModel := new(comparecomplyv1.ConvertToHTMLOptions)
				convertToHTMLOptionsModel.File = CreateMockReader("This is a mock file.")
				convertToHTMLOptionsModel.FileContentType = core.StringPtr("application/pdf")
				convertToHTMLOptionsModel.Model = core.StringPtr("contracts")
				convertToHTMLOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := compareComplyService.ConvertToHTMLWithContext(ctx, convertToHTMLOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				compareComplyService.DisableRetries()
				result, response, operationErr := compareComplyService.ConvertToHTML(convertToHTMLOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = compareComplyService.ConvertToHTMLWithContext(ctx, convertToHTMLOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(convertToHTMLPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["model"]).To(Equal([]string{"contracts"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"num_pages": "NumPages", "author": "Author", "publication_date": "PublicationDate", "title": "Title", "html": "HTML"}`)
				}))
			})
			It(`Invoke ConvertToHTML successfully`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := compareComplyService.ConvertToHTML(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ConvertToHTMLOptions model
				convertToHTMLOptionsModel := new(comparecomplyv1.ConvertToHTMLOptions)
				convertToHTMLOptionsModel.File = CreateMockReader("This is a mock file.")
				convertToHTMLOptionsModel.FileContentType = core.StringPtr("application/pdf")
				convertToHTMLOptionsModel.Model = core.StringPtr("contracts")
				convertToHTMLOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = compareComplyService.ConvertToHTML(convertToHTMLOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ConvertToHTML with error: Operation validation and request error`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Construct an instance of the ConvertToHTMLOptions model
				convertToHTMLOptionsModel := new(comparecomplyv1.ConvertToHTMLOptions)
				convertToHTMLOptionsModel.File = CreateMockReader("This is a mock file.")
				convertToHTMLOptionsModel.FileContentType = core.StringPtr("application/pdf")
				convertToHTMLOptionsModel.Model = core.StringPtr("contracts")
				convertToHTMLOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := compareComplyService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := compareComplyService.ConvertToHTML(convertToHTMLOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ConvertToHTMLOptions model with no property values
				convertToHTMLOptionsModelNew := new(comparecomplyv1.ConvertToHTMLOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = compareComplyService.ConvertToHTML(convertToHTMLOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ConvertToHTML successfully`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Construct an instance of the ConvertToHTMLOptions model
				convertToHTMLOptionsModel := new(comparecomplyv1.ConvertToHTMLOptions)
				convertToHTMLOptionsModel.File = CreateMockReader("This is a mock file.")
				convertToHTMLOptionsModel.FileContentType = core.StringPtr("application/pdf")
				convertToHTMLOptionsModel.Model = core.StringPtr("contracts")
				convertToHTMLOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := compareComplyService.ConvertToHTML(convertToHTMLOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ClassifyElements(classifyElementsOptions *ClassifyElementsOptions) - Operation response error`, func() {
		version := "testString"
		classifyElementsPath := "/v1/element_classification"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(classifyElementsPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["model"]).To(Equal([]string{"contracts"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ClassifyElements with error: Operation response processing error`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Construct an instance of the ClassifyElementsOptions model
				classifyElementsOptionsModel := new(comparecomplyv1.ClassifyElementsOptions)
				classifyElementsOptionsModel.File = CreateMockReader("This is a mock file.")
				classifyElementsOptionsModel.FileContentType = core.StringPtr("application/pdf")
				classifyElementsOptionsModel.Model = core.StringPtr("contracts")
				classifyElementsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := compareComplyService.ClassifyElements(classifyElementsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				compareComplyService.EnableRetries(0, 0)
				result, response, operationErr = compareComplyService.ClassifyElements(classifyElementsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ClassifyElements(classifyElementsOptions *ClassifyElementsOptions)`, func() {
		version := "testString"
		classifyElementsPath := "/v1/element_classification"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(classifyElementsPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["model"]).To(Equal([]string{"contracts"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"document": {"title": "Title", "html": "HTML", "hash": "Hash", "label": "Label"}, "model_id": "ModelID", "model_version": "ModelVersion", "elements": [{"location": {"begin": 5, "end": 3}, "text": "Text", "types": [{"label": {"nature": "Nature", "party": "Party"}, "provenance_ids": ["ProvenanceIds"], "modification": "added"}], "categories": [{"label": "Amendments", "provenance_ids": ["ProvenanceIds"], "modification": "added"}], "attributes": [{"type": "Currency", "text": "Text", "location": {"begin": 5, "end": 3}}]}], "effective_dates": [{"confidence_level": "High", "text": "Text", "text_normalized": "TextNormalized", "provenance_ids": ["ProvenanceIds"], "location": {"begin": 5, "end": 3}}], "contract_amounts": [{"confidence_level": "High", "text": "Text", "text_normalized": "TextNormalized", "interpretation": {"value": "Value", "numeric_value": 12, "unit": "Unit"}, "provenance_ids": ["ProvenanceIds"], "location": {"begin": 5, "end": 3}}], "termination_dates": [{"confidence_level": "High", "text": "Text", "text_normalized": "TextNormalized", "provenance_ids": ["ProvenanceIds"], "location": {"begin": 5, "end": 3}}], "contract_types": [{"confidence_level": "High", "text": "Text", "provenance_ids": ["ProvenanceIds"], "location": {"begin": 5, "end": 3}}], "contract_terms": [{"confidence_level": "High", "text": "Text", "text_normalized": "TextNormalized", "interpretation": {"value": "Value", "numeric_value": 12, "unit": "Unit"}, "provenance_ids": ["ProvenanceIds"], "location": {"begin": 5, "end": 3}}], "payment_terms": [{"confidence_level": "High", "text": "Text", "text_normalized": "TextNormalized", "interpretation": {"value": "Value", "numeric_value": 12, "unit": "Unit"}, "provenance_ids": ["ProvenanceIds"], "location": {"begin": 5, "end": 3}}], "contract_currencies": [{"confidence_level": "High", "text": "Text", "text_normalized": "TextNormalized", "provenance_ids": ["ProvenanceIds"], "location": {"begin": 5, "end": 3}}], "tables": [{"location": {"begin": 5, "end": 3}, "text": "Text", "section_title": {"text": "Text", "location": {"begin": 5, "end": 3}}, "title": {"location": {"begin": 5, "end": 3}, "text": "Text"}, "table_headers": [{"cell_id": "CellID", "location": {"anyKey": "anyValue"}, "text": "Text", "row_index_begin": 13, "row_index_end": 11, "column_index_begin": 16, "column_index_end": 14}], "row_headers": [{"cell_id": "CellID", "location": {"begin": 5, "end": 3}, "text": "Text", "text_normalized": "TextNormalized", "row_index_begin": 13, "row_index_end": 11, "column_index_begin": 16, "column_index_end": 14}], "column_headers": [{"cell_id": "CellID", "location": {"anyKey": "anyValue"}, "text": "Text", "text_normalized": "TextNormalized", "row_index_begin": 13, "row_index_end": 11, "column_index_begin": 16, "column_index_end": 14}], "body_cells": [{"cell_id": "CellID", "location": {"begin": 5, "end": 3}, "text": "Text", "row_index_begin": 13, "row_index_end": 11, "column_index_begin": 16, "column_index_end": 14, "row_header_ids": ["RowHeaderIds"], "row_header_texts": ["RowHeaderTexts"], "row_header_texts_normalized": ["RowHeaderTextsNormalized"], "column_header_ids": ["ColumnHeaderIds"], "column_header_texts": ["ColumnHeaderTexts"], "column_header_texts_normalized": ["ColumnHeaderTextsNormalized"], "attributes": [{"type": "Currency", "text": "Text", "location": {"begin": 5, "end": 3}}]}], "contexts": [{"text": "Text", "location": {"begin": 5, "end": 3}}], "key_value_pairs": [{"key": {"cell_id": "CellID", "location": {"begin": 5, "end": 3}, "text": "Text"}, "value": [{"cell_id": "CellID", "location": {"begin": 5, "end": 3}, "text": "Text"}]}]}], "document_structure": {"section_titles": [{"text": "Text", "location": {"begin": 5, "end": 3}, "level": 5, "element_locations": [{"begin": 5, "end": 3}]}], "leading_sentences": [{"text": "Text", "location": {"begin": 5, "end": 3}, "element_locations": [{"begin": 5, "end": 3}]}], "paragraphs": [{"location": {"begin": 5, "end": 3}}]}, "parties": [{"party": "Party", "role": "Role", "importance": "Primary", "addresses": [{"text": "Text", "location": {"begin": 5, "end": 3}}], "contacts": [{"name": "Name", "role": "Role"}], "mentions": [{"text": "Text", "location": {"begin": 5, "end": 3}}]}]}`)
				}))
			})
			It(`Invoke ClassifyElements successfully with retries`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())
				compareComplyService.EnableRetries(0, 0)

				// Construct an instance of the ClassifyElementsOptions model
				classifyElementsOptionsModel := new(comparecomplyv1.ClassifyElementsOptions)
				classifyElementsOptionsModel.File = CreateMockReader("This is a mock file.")
				classifyElementsOptionsModel.FileContentType = core.StringPtr("application/pdf")
				classifyElementsOptionsModel.Model = core.StringPtr("contracts")
				classifyElementsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := compareComplyService.ClassifyElementsWithContext(ctx, classifyElementsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				compareComplyService.DisableRetries()
				result, response, operationErr := compareComplyService.ClassifyElements(classifyElementsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = compareComplyService.ClassifyElementsWithContext(ctx, classifyElementsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(classifyElementsPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["model"]).To(Equal([]string{"contracts"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"document": {"title": "Title", "html": "HTML", "hash": "Hash", "label": "Label"}, "model_id": "ModelID", "model_version": "ModelVersion", "elements": [{"location": {"begin": 5, "end": 3}, "text": "Text", "types": [{"label": {"nature": "Nature", "party": "Party"}, "provenance_ids": ["ProvenanceIds"], "modification": "added"}], "categories": [{"label": "Amendments", "provenance_ids": ["ProvenanceIds"], "modification": "added"}], "attributes": [{"type": "Currency", "text": "Text", "location": {"begin": 5, "end": 3}}]}], "effective_dates": [{"confidence_level": "High", "text": "Text", "text_normalized": "TextNormalized", "provenance_ids": ["ProvenanceIds"], "location": {"begin": 5, "end": 3}}], "contract_amounts": [{"confidence_level": "High", "text": "Text", "text_normalized": "TextNormalized", "interpretation": {"value": "Value", "numeric_value": 12, "unit": "Unit"}, "provenance_ids": ["ProvenanceIds"], "location": {"begin": 5, "end": 3}}], "termination_dates": [{"confidence_level": "High", "text": "Text", "text_normalized": "TextNormalized", "provenance_ids": ["ProvenanceIds"], "location": {"begin": 5, "end": 3}}], "contract_types": [{"confidence_level": "High", "text": "Text", "provenance_ids": ["ProvenanceIds"], "location": {"begin": 5, "end": 3}}], "contract_terms": [{"confidence_level": "High", "text": "Text", "text_normalized": "TextNormalized", "interpretation": {"value": "Value", "numeric_value": 12, "unit": "Unit"}, "provenance_ids": ["ProvenanceIds"], "location": {"begin": 5, "end": 3}}], "payment_terms": [{"confidence_level": "High", "text": "Text", "text_normalized": "TextNormalized", "interpretation": {"value": "Value", "numeric_value": 12, "unit": "Unit"}, "provenance_ids": ["ProvenanceIds"], "location": {"begin": 5, "end": 3}}], "contract_currencies": [{"confidence_level": "High", "text": "Text", "text_normalized": "TextNormalized", "provenance_ids": ["ProvenanceIds"], "location": {"begin": 5, "end": 3}}], "tables": [{"location": {"begin": 5, "end": 3}, "text": "Text", "section_title": {"text": "Text", "location": {"begin": 5, "end": 3}}, "title": {"location": {"begin": 5, "end": 3}, "text": "Text"}, "table_headers": [{"cell_id": "CellID", "location": {"anyKey": "anyValue"}, "text": "Text", "row_index_begin": 13, "row_index_end": 11, "column_index_begin": 16, "column_index_end": 14}], "row_headers": [{"cell_id": "CellID", "location": {"begin": 5, "end": 3}, "text": "Text", "text_normalized": "TextNormalized", "row_index_begin": 13, "row_index_end": 11, "column_index_begin": 16, "column_index_end": 14}], "column_headers": [{"cell_id": "CellID", "location": {"anyKey": "anyValue"}, "text": "Text", "text_normalized": "TextNormalized", "row_index_begin": 13, "row_index_end": 11, "column_index_begin": 16, "column_index_end": 14}], "body_cells": [{"cell_id": "CellID", "location": {"begin": 5, "end": 3}, "text": "Text", "row_index_begin": 13, "row_index_end": 11, "column_index_begin": 16, "column_index_end": 14, "row_header_ids": ["RowHeaderIds"], "row_header_texts": ["RowHeaderTexts"], "row_header_texts_normalized": ["RowHeaderTextsNormalized"], "column_header_ids": ["ColumnHeaderIds"], "column_header_texts": ["ColumnHeaderTexts"], "column_header_texts_normalized": ["ColumnHeaderTextsNormalized"], "attributes": [{"type": "Currency", "text": "Text", "location": {"begin": 5, "end": 3}}]}], "contexts": [{"text": "Text", "location": {"begin": 5, "end": 3}}], "key_value_pairs": [{"key": {"cell_id": "CellID", "location": {"begin": 5, "end": 3}, "text": "Text"}, "value": [{"cell_id": "CellID", "location": {"begin": 5, "end": 3}, "text": "Text"}]}]}], "document_structure": {"section_titles": [{"text": "Text", "location": {"begin": 5, "end": 3}, "level": 5, "element_locations": [{"begin": 5, "end": 3}]}], "leading_sentences": [{"text": "Text", "location": {"begin": 5, "end": 3}, "element_locations": [{"begin": 5, "end": 3}]}], "paragraphs": [{"location": {"begin": 5, "end": 3}}]}, "parties": [{"party": "Party", "role": "Role", "importance": "Primary", "addresses": [{"text": "Text", "location": {"begin": 5, "end": 3}}], "contacts": [{"name": "Name", "role": "Role"}], "mentions": [{"text": "Text", "location": {"begin": 5, "end": 3}}]}]}`)
				}))
			})
			It(`Invoke ClassifyElements successfully`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := compareComplyService.ClassifyElements(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ClassifyElementsOptions model
				classifyElementsOptionsModel := new(comparecomplyv1.ClassifyElementsOptions)
				classifyElementsOptionsModel.File = CreateMockReader("This is a mock file.")
				classifyElementsOptionsModel.FileContentType = core.StringPtr("application/pdf")
				classifyElementsOptionsModel.Model = core.StringPtr("contracts")
				classifyElementsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = compareComplyService.ClassifyElements(classifyElementsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ClassifyElements with error: Operation validation and request error`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Construct an instance of the ClassifyElementsOptions model
				classifyElementsOptionsModel := new(comparecomplyv1.ClassifyElementsOptions)
				classifyElementsOptionsModel.File = CreateMockReader("This is a mock file.")
				classifyElementsOptionsModel.FileContentType = core.StringPtr("application/pdf")
				classifyElementsOptionsModel.Model = core.StringPtr("contracts")
				classifyElementsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := compareComplyService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := compareComplyService.ClassifyElements(classifyElementsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ClassifyElementsOptions model with no property values
				classifyElementsOptionsModelNew := new(comparecomplyv1.ClassifyElementsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = compareComplyService.ClassifyElements(classifyElementsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ClassifyElements successfully`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Construct an instance of the ClassifyElementsOptions model
				classifyElementsOptionsModel := new(comparecomplyv1.ClassifyElementsOptions)
				classifyElementsOptionsModel.File = CreateMockReader("This is a mock file.")
				classifyElementsOptionsModel.FileContentType = core.StringPtr("application/pdf")
				classifyElementsOptionsModel.Model = core.StringPtr("contracts")
				classifyElementsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := compareComplyService.ClassifyElements(classifyElementsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ExtractTables(extractTablesOptions *ExtractTablesOptions) - Operation response error`, func() {
		version := "testString"
		extractTablesPath := "/v1/tables"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(extractTablesPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["model"]).To(Equal([]string{"contracts"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ExtractTables with error: Operation response processing error`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Construct an instance of the ExtractTablesOptions model
				extractTablesOptionsModel := new(comparecomplyv1.ExtractTablesOptions)
				extractTablesOptionsModel.File = CreateMockReader("This is a mock file.")
				extractTablesOptionsModel.FileContentType = core.StringPtr("application/pdf")
				extractTablesOptionsModel.Model = core.StringPtr("contracts")
				extractTablesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := compareComplyService.ExtractTables(extractTablesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				compareComplyService.EnableRetries(0, 0)
				result, response, operationErr = compareComplyService.ExtractTables(extractTablesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ExtractTables(extractTablesOptions *ExtractTablesOptions)`, func() {
		version := "testString"
		extractTablesPath := "/v1/tables"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(extractTablesPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["model"]).To(Equal([]string{"contracts"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"document": {"html": "HTML", "title": "Title", "hash": "Hash"}, "model_id": "ModelID", "model_version": "ModelVersion", "tables": [{"location": {"begin": 5, "end": 3}, "text": "Text", "section_title": {"text": "Text", "location": {"begin": 5, "end": 3}}, "title": {"location": {"begin": 5, "end": 3}, "text": "Text"}, "table_headers": [{"cell_id": "CellID", "location": {"anyKey": "anyValue"}, "text": "Text", "row_index_begin": 13, "row_index_end": 11, "column_index_begin": 16, "column_index_end": 14}], "row_headers": [{"cell_id": "CellID", "location": {"begin": 5, "end": 3}, "text": "Text", "text_normalized": "TextNormalized", "row_index_begin": 13, "row_index_end": 11, "column_index_begin": 16, "column_index_end": 14}], "column_headers": [{"cell_id": "CellID", "location": {"anyKey": "anyValue"}, "text": "Text", "text_normalized": "TextNormalized", "row_index_begin": 13, "row_index_end": 11, "column_index_begin": 16, "column_index_end": 14}], "body_cells": [{"cell_id": "CellID", "location": {"begin": 5, "end": 3}, "text": "Text", "row_index_begin": 13, "row_index_end": 11, "column_index_begin": 16, "column_index_end": 14, "row_header_ids": ["RowHeaderIds"], "row_header_texts": ["RowHeaderTexts"], "row_header_texts_normalized": ["RowHeaderTextsNormalized"], "column_header_ids": ["ColumnHeaderIds"], "column_header_texts": ["ColumnHeaderTexts"], "column_header_texts_normalized": ["ColumnHeaderTextsNormalized"], "attributes": [{"type": "Currency", "text": "Text", "location": {"begin": 5, "end": 3}}]}], "contexts": [{"text": "Text", "location": {"begin": 5, "end": 3}}], "key_value_pairs": [{"key": {"cell_id": "CellID", "location": {"begin": 5, "end": 3}, "text": "Text"}, "value": [{"cell_id": "CellID", "location": {"begin": 5, "end": 3}, "text": "Text"}]}]}]}`)
				}))
			})
			It(`Invoke ExtractTables successfully with retries`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())
				compareComplyService.EnableRetries(0, 0)

				// Construct an instance of the ExtractTablesOptions model
				extractTablesOptionsModel := new(comparecomplyv1.ExtractTablesOptions)
				extractTablesOptionsModel.File = CreateMockReader("This is a mock file.")
				extractTablesOptionsModel.FileContentType = core.StringPtr("application/pdf")
				extractTablesOptionsModel.Model = core.StringPtr("contracts")
				extractTablesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := compareComplyService.ExtractTablesWithContext(ctx, extractTablesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				compareComplyService.DisableRetries()
				result, response, operationErr := compareComplyService.ExtractTables(extractTablesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = compareComplyService.ExtractTablesWithContext(ctx, extractTablesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(extractTablesPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["model"]).To(Equal([]string{"contracts"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"document": {"html": "HTML", "title": "Title", "hash": "Hash"}, "model_id": "ModelID", "model_version": "ModelVersion", "tables": [{"location": {"begin": 5, "end": 3}, "text": "Text", "section_title": {"text": "Text", "location": {"begin": 5, "end": 3}}, "title": {"location": {"begin": 5, "end": 3}, "text": "Text"}, "table_headers": [{"cell_id": "CellID", "location": {"anyKey": "anyValue"}, "text": "Text", "row_index_begin": 13, "row_index_end": 11, "column_index_begin": 16, "column_index_end": 14}], "row_headers": [{"cell_id": "CellID", "location": {"begin": 5, "end": 3}, "text": "Text", "text_normalized": "TextNormalized", "row_index_begin": 13, "row_index_end": 11, "column_index_begin": 16, "column_index_end": 14}], "column_headers": [{"cell_id": "CellID", "location": {"anyKey": "anyValue"}, "text": "Text", "text_normalized": "TextNormalized", "row_index_begin": 13, "row_index_end": 11, "column_index_begin": 16, "column_index_end": 14}], "body_cells": [{"cell_id": "CellID", "location": {"begin": 5, "end": 3}, "text": "Text", "row_index_begin": 13, "row_index_end": 11, "column_index_begin": 16, "column_index_end": 14, "row_header_ids": ["RowHeaderIds"], "row_header_texts": ["RowHeaderTexts"], "row_header_texts_normalized": ["RowHeaderTextsNormalized"], "column_header_ids": ["ColumnHeaderIds"], "column_header_texts": ["ColumnHeaderTexts"], "column_header_texts_normalized": ["ColumnHeaderTextsNormalized"], "attributes": [{"type": "Currency", "text": "Text", "location": {"begin": 5, "end": 3}}]}], "contexts": [{"text": "Text", "location": {"begin": 5, "end": 3}}], "key_value_pairs": [{"key": {"cell_id": "CellID", "location": {"begin": 5, "end": 3}, "text": "Text"}, "value": [{"cell_id": "CellID", "location": {"begin": 5, "end": 3}, "text": "Text"}]}]}]}`)
				}))
			})
			It(`Invoke ExtractTables successfully`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := compareComplyService.ExtractTables(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ExtractTablesOptions model
				extractTablesOptionsModel := new(comparecomplyv1.ExtractTablesOptions)
				extractTablesOptionsModel.File = CreateMockReader("This is a mock file.")
				extractTablesOptionsModel.FileContentType = core.StringPtr("application/pdf")
				extractTablesOptionsModel.Model = core.StringPtr("contracts")
				extractTablesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = compareComplyService.ExtractTables(extractTablesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ExtractTables with error: Operation validation and request error`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Construct an instance of the ExtractTablesOptions model
				extractTablesOptionsModel := new(comparecomplyv1.ExtractTablesOptions)
				extractTablesOptionsModel.File = CreateMockReader("This is a mock file.")
				extractTablesOptionsModel.FileContentType = core.StringPtr("application/pdf")
				extractTablesOptionsModel.Model = core.StringPtr("contracts")
				extractTablesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := compareComplyService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := compareComplyService.ExtractTables(extractTablesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ExtractTablesOptions model with no property values
				extractTablesOptionsModelNew := new(comparecomplyv1.ExtractTablesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = compareComplyService.ExtractTables(extractTablesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ExtractTables successfully`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Construct an instance of the ExtractTablesOptions model
				extractTablesOptionsModel := new(comparecomplyv1.ExtractTablesOptions)
				extractTablesOptionsModel.File = CreateMockReader("This is a mock file.")
				extractTablesOptionsModel.FileContentType = core.StringPtr("application/pdf")
				extractTablesOptionsModel.Model = core.StringPtr("contracts")
				extractTablesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := compareComplyService.ExtractTables(extractTablesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CompareDocuments(compareDocumentsOptions *CompareDocumentsOptions) - Operation response error`, func() {
		version := "testString"
		compareDocumentsPath := "/v1/comparison"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(compareDocumentsPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["file_1_label"]).To(Equal([]string{"file_1"}))
					Expect(req.URL.Query()["file_2_label"]).To(Equal([]string{"file_2"}))
					Expect(req.URL.Query()["model"]).To(Equal([]string{"contracts"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CompareDocuments with error: Operation response processing error`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Construct an instance of the CompareDocumentsOptions model
				compareDocumentsOptionsModel := new(comparecomplyv1.CompareDocumentsOptions)
				compareDocumentsOptionsModel.File1 = CreateMockReader("This is a mock file.")
				compareDocumentsOptionsModel.File2 = CreateMockReader("This is a mock file.")
				compareDocumentsOptionsModel.File1ContentType = core.StringPtr("application/pdf")
				compareDocumentsOptionsModel.File2ContentType = core.StringPtr("application/pdf")
				compareDocumentsOptionsModel.File1Label = core.StringPtr("file_1")
				compareDocumentsOptionsModel.File2Label = core.StringPtr("file_2")
				compareDocumentsOptionsModel.Model = core.StringPtr("contracts")
				compareDocumentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := compareComplyService.CompareDocuments(compareDocumentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				compareComplyService.EnableRetries(0, 0)
				result, response, operationErr = compareComplyService.CompareDocuments(compareDocumentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CompareDocuments(compareDocumentsOptions *CompareDocumentsOptions)`, func() {
		version := "testString"
		compareDocumentsPath := "/v1/comparison"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(compareDocumentsPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["file_1_label"]).To(Equal([]string{"file_1"}))
					Expect(req.URL.Query()["file_2_label"]).To(Equal([]string{"file_2"}))
					Expect(req.URL.Query()["model"]).To(Equal([]string{"contracts"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"model_id": "ModelID", "model_version": "ModelVersion", "documents": [{"title": "Title", "html": "HTML", "hash": "Hash", "label": "Label"}], "aligned_elements": [{"element_pair": [{"document_label": "DocumentLabel", "text": "Text", "location": {"begin": 5, "end": 3}, "types": [{"label": {"nature": "Nature", "party": "Party"}}], "categories": [{"label": "Amendments"}], "attributes": [{"type": "Currency", "text": "Text", "location": {"begin": 5, "end": 3}}]}], "identical_text": false, "provenance_ids": ["ProvenanceIds"], "significant_elements": false}], "unaligned_elements": [{"document_label": "DocumentLabel", "location": {"begin": 5, "end": 3}, "text": "Text", "types": [{"label": {"nature": "Nature", "party": "Party"}}], "categories": [{"label": "Amendments"}], "attributes": [{"type": "Currency", "text": "Text", "location": {"begin": 5, "end": 3}}]}]}`)
				}))
			})
			It(`Invoke CompareDocuments successfully with retries`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())
				compareComplyService.EnableRetries(0, 0)

				// Construct an instance of the CompareDocumentsOptions model
				compareDocumentsOptionsModel := new(comparecomplyv1.CompareDocumentsOptions)
				compareDocumentsOptionsModel.File1 = CreateMockReader("This is a mock file.")
				compareDocumentsOptionsModel.File2 = CreateMockReader("This is a mock file.")
				compareDocumentsOptionsModel.File1ContentType = core.StringPtr("application/pdf")
				compareDocumentsOptionsModel.File2ContentType = core.StringPtr("application/pdf")
				compareDocumentsOptionsModel.File1Label = core.StringPtr("file_1")
				compareDocumentsOptionsModel.File2Label = core.StringPtr("file_2")
				compareDocumentsOptionsModel.Model = core.StringPtr("contracts")
				compareDocumentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := compareComplyService.CompareDocumentsWithContext(ctx, compareDocumentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				compareComplyService.DisableRetries()
				result, response, operationErr := compareComplyService.CompareDocuments(compareDocumentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = compareComplyService.CompareDocumentsWithContext(ctx, compareDocumentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(compareDocumentsPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["file_1_label"]).To(Equal([]string{"file_1"}))
					Expect(req.URL.Query()["file_2_label"]).To(Equal([]string{"file_2"}))
					Expect(req.URL.Query()["model"]).To(Equal([]string{"contracts"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"model_id": "ModelID", "model_version": "ModelVersion", "documents": [{"title": "Title", "html": "HTML", "hash": "Hash", "label": "Label"}], "aligned_elements": [{"element_pair": [{"document_label": "DocumentLabel", "text": "Text", "location": {"begin": 5, "end": 3}, "types": [{"label": {"nature": "Nature", "party": "Party"}}], "categories": [{"label": "Amendments"}], "attributes": [{"type": "Currency", "text": "Text", "location": {"begin": 5, "end": 3}}]}], "identical_text": false, "provenance_ids": ["ProvenanceIds"], "significant_elements": false}], "unaligned_elements": [{"document_label": "DocumentLabel", "location": {"begin": 5, "end": 3}, "text": "Text", "types": [{"label": {"nature": "Nature", "party": "Party"}}], "categories": [{"label": "Amendments"}], "attributes": [{"type": "Currency", "text": "Text", "location": {"begin": 5, "end": 3}}]}]}`)
				}))
			})
			It(`Invoke CompareDocuments successfully`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := compareComplyService.CompareDocuments(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CompareDocumentsOptions model
				compareDocumentsOptionsModel := new(comparecomplyv1.CompareDocumentsOptions)
				compareDocumentsOptionsModel.File1 = CreateMockReader("This is a mock file.")
				compareDocumentsOptionsModel.File2 = CreateMockReader("This is a mock file.")
				compareDocumentsOptionsModel.File1ContentType = core.StringPtr("application/pdf")
				compareDocumentsOptionsModel.File2ContentType = core.StringPtr("application/pdf")
				compareDocumentsOptionsModel.File1Label = core.StringPtr("file_1")
				compareDocumentsOptionsModel.File2Label = core.StringPtr("file_2")
				compareDocumentsOptionsModel.Model = core.StringPtr("contracts")
				compareDocumentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = compareComplyService.CompareDocuments(compareDocumentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CompareDocuments with error: Operation validation and request error`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Construct an instance of the CompareDocumentsOptions model
				compareDocumentsOptionsModel := new(comparecomplyv1.CompareDocumentsOptions)
				compareDocumentsOptionsModel.File1 = CreateMockReader("This is a mock file.")
				compareDocumentsOptionsModel.File2 = CreateMockReader("This is a mock file.")
				compareDocumentsOptionsModel.File1ContentType = core.StringPtr("application/pdf")
				compareDocumentsOptionsModel.File2ContentType = core.StringPtr("application/pdf")
				compareDocumentsOptionsModel.File1Label = core.StringPtr("file_1")
				compareDocumentsOptionsModel.File2Label = core.StringPtr("file_2")
				compareDocumentsOptionsModel.Model = core.StringPtr("contracts")
				compareDocumentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := compareComplyService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := compareComplyService.CompareDocuments(compareDocumentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CompareDocumentsOptions model with no property values
				compareDocumentsOptionsModelNew := new(comparecomplyv1.CompareDocumentsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = compareComplyService.CompareDocuments(compareDocumentsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke CompareDocuments successfully`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Construct an instance of the CompareDocumentsOptions model
				compareDocumentsOptionsModel := new(comparecomplyv1.CompareDocumentsOptions)
				compareDocumentsOptionsModel.File1 = CreateMockReader("This is a mock file.")
				compareDocumentsOptionsModel.File2 = CreateMockReader("This is a mock file.")
				compareDocumentsOptionsModel.File1ContentType = core.StringPtr("application/pdf")
				compareDocumentsOptionsModel.File2ContentType = core.StringPtr("application/pdf")
				compareDocumentsOptionsModel.File1Label = core.StringPtr("file_1")
				compareDocumentsOptionsModel.File2Label = core.StringPtr("file_2")
				compareDocumentsOptionsModel.Model = core.StringPtr("contracts")
				compareDocumentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := compareComplyService.CompareDocuments(compareDocumentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AddFeedback(addFeedbackOptions *AddFeedbackOptions) - Operation response error`, func() {
		version := "testString"
		addFeedbackPath := "/v1/feedback"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addFeedbackPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke AddFeedback with error: Operation response processing error`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Construct an instance of the ShortDoc model
				shortDocModel := new(comparecomplyv1.ShortDoc)
				shortDocModel.Title = core.StringPtr("testString")
				shortDocModel.Hash = core.StringPtr("testString")

				// Construct an instance of the Location model
				locationModel := new(comparecomplyv1.Location)
				locationModel.Begin = core.Int64Ptr(int64(26))
				locationModel.End = core.Int64Ptr(int64(26))

				// Construct an instance of the Label model
				labelModel := new(comparecomplyv1.Label)
				labelModel.Nature = core.StringPtr("testString")
				labelModel.Party = core.StringPtr("testString")

				// Construct an instance of the TypeLabel model
				typeLabelModel := new(comparecomplyv1.TypeLabel)
				typeLabelModel.Label = labelModel
				typeLabelModel.ProvenanceIds = []string{"testString"}
				typeLabelModel.Modification = core.StringPtr("added")

				// Construct an instance of the Category model
				categoryModel := new(comparecomplyv1.Category)
				categoryModel.Label = core.StringPtr("Amendments")
				categoryModel.ProvenanceIds = []string{"testString"}
				categoryModel.Modification = core.StringPtr("added")

				// Construct an instance of the OriginalLabelsIn model
				originalLabelsInModel := new(comparecomplyv1.OriginalLabelsIn)
				originalLabelsInModel.Types = []comparecomplyv1.TypeLabel{*typeLabelModel}
				originalLabelsInModel.Categories = []comparecomplyv1.Category{*categoryModel}

				// Construct an instance of the UpdatedLabelsIn model
				updatedLabelsInModel := new(comparecomplyv1.UpdatedLabelsIn)
				updatedLabelsInModel.Types = []comparecomplyv1.TypeLabel{*typeLabelModel}
				updatedLabelsInModel.Categories = []comparecomplyv1.Category{*categoryModel}

				// Construct an instance of the FeedbackDataInput model
				feedbackDataInputModel := new(comparecomplyv1.FeedbackDataInput)
				feedbackDataInputModel.FeedbackType = core.StringPtr("testString")
				feedbackDataInputModel.Document = shortDocModel
				feedbackDataInputModel.ModelID = core.StringPtr("testString")
				feedbackDataInputModel.ModelVersion = core.StringPtr("testString")
				feedbackDataInputModel.Location = locationModel
				feedbackDataInputModel.Text = core.StringPtr("testString")
				feedbackDataInputModel.OriginalLabels = originalLabelsInModel
				feedbackDataInputModel.UpdatedLabels = updatedLabelsInModel

				// Construct an instance of the AddFeedbackOptions model
				addFeedbackOptionsModel := new(comparecomplyv1.AddFeedbackOptions)
				addFeedbackOptionsModel.FeedbackData = feedbackDataInputModel
				addFeedbackOptionsModel.UserID = core.StringPtr("testString")
				addFeedbackOptionsModel.Comment = core.StringPtr("testString")
				addFeedbackOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := compareComplyService.AddFeedback(addFeedbackOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				compareComplyService.EnableRetries(0, 0)
				result, response, operationErr = compareComplyService.AddFeedback(addFeedbackOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AddFeedback(addFeedbackOptions *AddFeedbackOptions)`, func() {
		version := "testString"
		addFeedbackPath := "/v1/feedback"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addFeedbackPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"feedback_id": "FeedbackID", "user_id": "UserID", "comment": "Comment", "created": "2019-01-01T12:00:00.000Z", "feedback_data": {"feedback_type": "FeedbackType", "document": {"title": "Title", "hash": "Hash"}, "model_id": "ModelID", "model_version": "ModelVersion", "location": {"begin": 5, "end": 3}, "text": "Text", "original_labels": {"types": [{"label": {"nature": "Nature", "party": "Party"}, "provenance_ids": ["ProvenanceIds"], "modification": "added"}], "categories": [{"label": "Amendments", "provenance_ids": ["ProvenanceIds"], "modification": "added"}]}, "updated_labels": {"types": [{"label": {"nature": "Nature", "party": "Party"}, "provenance_ids": ["ProvenanceIds"], "modification": "added"}], "categories": [{"label": "Amendments", "provenance_ids": ["ProvenanceIds"], "modification": "added"}]}, "pagination": {"refresh_cursor": "RefreshCursor", "next_cursor": "NextCursor", "refresh_url": "RefreshURL", "next_url": "NextURL", "total": 5}}}`)
				}))
			})
			It(`Invoke AddFeedback successfully with retries`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())
				compareComplyService.EnableRetries(0, 0)

				// Construct an instance of the ShortDoc model
				shortDocModel := new(comparecomplyv1.ShortDoc)
				shortDocModel.Title = core.StringPtr("testString")
				shortDocModel.Hash = core.StringPtr("testString")

				// Construct an instance of the Location model
				locationModel := new(comparecomplyv1.Location)
				locationModel.Begin = core.Int64Ptr(int64(26))
				locationModel.End = core.Int64Ptr(int64(26))

				// Construct an instance of the Label model
				labelModel := new(comparecomplyv1.Label)
				labelModel.Nature = core.StringPtr("testString")
				labelModel.Party = core.StringPtr("testString")

				// Construct an instance of the TypeLabel model
				typeLabelModel := new(comparecomplyv1.TypeLabel)
				typeLabelModel.Label = labelModel
				typeLabelModel.ProvenanceIds = []string{"testString"}
				typeLabelModel.Modification = core.StringPtr("added")

				// Construct an instance of the Category model
				categoryModel := new(comparecomplyv1.Category)
				categoryModel.Label = core.StringPtr("Amendments")
				categoryModel.ProvenanceIds = []string{"testString"}
				categoryModel.Modification = core.StringPtr("added")

				// Construct an instance of the OriginalLabelsIn model
				originalLabelsInModel := new(comparecomplyv1.OriginalLabelsIn)
				originalLabelsInModel.Types = []comparecomplyv1.TypeLabel{*typeLabelModel}
				originalLabelsInModel.Categories = []comparecomplyv1.Category{*categoryModel}

				// Construct an instance of the UpdatedLabelsIn model
				updatedLabelsInModel := new(comparecomplyv1.UpdatedLabelsIn)
				updatedLabelsInModel.Types = []comparecomplyv1.TypeLabel{*typeLabelModel}
				updatedLabelsInModel.Categories = []comparecomplyv1.Category{*categoryModel}

				// Construct an instance of the FeedbackDataInput model
				feedbackDataInputModel := new(comparecomplyv1.FeedbackDataInput)
				feedbackDataInputModel.FeedbackType = core.StringPtr("testString")
				feedbackDataInputModel.Document = shortDocModel
				feedbackDataInputModel.ModelID = core.StringPtr("testString")
				feedbackDataInputModel.ModelVersion = core.StringPtr("testString")
				feedbackDataInputModel.Location = locationModel
				feedbackDataInputModel.Text = core.StringPtr("testString")
				feedbackDataInputModel.OriginalLabels = originalLabelsInModel
				feedbackDataInputModel.UpdatedLabels = updatedLabelsInModel

				// Construct an instance of the AddFeedbackOptions model
				addFeedbackOptionsModel := new(comparecomplyv1.AddFeedbackOptions)
				addFeedbackOptionsModel.FeedbackData = feedbackDataInputModel
				addFeedbackOptionsModel.UserID = core.StringPtr("testString")
				addFeedbackOptionsModel.Comment = core.StringPtr("testString")
				addFeedbackOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := compareComplyService.AddFeedbackWithContext(ctx, addFeedbackOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				compareComplyService.DisableRetries()
				result, response, operationErr := compareComplyService.AddFeedback(addFeedbackOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = compareComplyService.AddFeedbackWithContext(ctx, addFeedbackOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addFeedbackPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"feedback_id": "FeedbackID", "user_id": "UserID", "comment": "Comment", "created": "2019-01-01T12:00:00.000Z", "feedback_data": {"feedback_type": "FeedbackType", "document": {"title": "Title", "hash": "Hash"}, "model_id": "ModelID", "model_version": "ModelVersion", "location": {"begin": 5, "end": 3}, "text": "Text", "original_labels": {"types": [{"label": {"nature": "Nature", "party": "Party"}, "provenance_ids": ["ProvenanceIds"], "modification": "added"}], "categories": [{"label": "Amendments", "provenance_ids": ["ProvenanceIds"], "modification": "added"}]}, "updated_labels": {"types": [{"label": {"nature": "Nature", "party": "Party"}, "provenance_ids": ["ProvenanceIds"], "modification": "added"}], "categories": [{"label": "Amendments", "provenance_ids": ["ProvenanceIds"], "modification": "added"}]}, "pagination": {"refresh_cursor": "RefreshCursor", "next_cursor": "NextCursor", "refresh_url": "RefreshURL", "next_url": "NextURL", "total": 5}}}`)
				}))
			})
			It(`Invoke AddFeedback successfully`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := compareComplyService.AddFeedback(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ShortDoc model
				shortDocModel := new(comparecomplyv1.ShortDoc)
				shortDocModel.Title = core.StringPtr("testString")
				shortDocModel.Hash = core.StringPtr("testString")

				// Construct an instance of the Location model
				locationModel := new(comparecomplyv1.Location)
				locationModel.Begin = core.Int64Ptr(int64(26))
				locationModel.End = core.Int64Ptr(int64(26))

				// Construct an instance of the Label model
				labelModel := new(comparecomplyv1.Label)
				labelModel.Nature = core.StringPtr("testString")
				labelModel.Party = core.StringPtr("testString")

				// Construct an instance of the TypeLabel model
				typeLabelModel := new(comparecomplyv1.TypeLabel)
				typeLabelModel.Label = labelModel
				typeLabelModel.ProvenanceIds = []string{"testString"}
				typeLabelModel.Modification = core.StringPtr("added")

				// Construct an instance of the Category model
				categoryModel := new(comparecomplyv1.Category)
				categoryModel.Label = core.StringPtr("Amendments")
				categoryModel.ProvenanceIds = []string{"testString"}
				categoryModel.Modification = core.StringPtr("added")

				// Construct an instance of the OriginalLabelsIn model
				originalLabelsInModel := new(comparecomplyv1.OriginalLabelsIn)
				originalLabelsInModel.Types = []comparecomplyv1.TypeLabel{*typeLabelModel}
				originalLabelsInModel.Categories = []comparecomplyv1.Category{*categoryModel}

				// Construct an instance of the UpdatedLabelsIn model
				updatedLabelsInModel := new(comparecomplyv1.UpdatedLabelsIn)
				updatedLabelsInModel.Types = []comparecomplyv1.TypeLabel{*typeLabelModel}
				updatedLabelsInModel.Categories = []comparecomplyv1.Category{*categoryModel}

				// Construct an instance of the FeedbackDataInput model
				feedbackDataInputModel := new(comparecomplyv1.FeedbackDataInput)
				feedbackDataInputModel.FeedbackType = core.StringPtr("testString")
				feedbackDataInputModel.Document = shortDocModel
				feedbackDataInputModel.ModelID = core.StringPtr("testString")
				feedbackDataInputModel.ModelVersion = core.StringPtr("testString")
				feedbackDataInputModel.Location = locationModel
				feedbackDataInputModel.Text = core.StringPtr("testString")
				feedbackDataInputModel.OriginalLabels = originalLabelsInModel
				feedbackDataInputModel.UpdatedLabels = updatedLabelsInModel

				// Construct an instance of the AddFeedbackOptions model
				addFeedbackOptionsModel := new(comparecomplyv1.AddFeedbackOptions)
				addFeedbackOptionsModel.FeedbackData = feedbackDataInputModel
				addFeedbackOptionsModel.UserID = core.StringPtr("testString")
				addFeedbackOptionsModel.Comment = core.StringPtr("testString")
				addFeedbackOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = compareComplyService.AddFeedback(addFeedbackOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke AddFeedback with error: Operation validation and request error`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Construct an instance of the ShortDoc model
				shortDocModel := new(comparecomplyv1.ShortDoc)
				shortDocModel.Title = core.StringPtr("testString")
				shortDocModel.Hash = core.StringPtr("testString")

				// Construct an instance of the Location model
				locationModel := new(comparecomplyv1.Location)
				locationModel.Begin = core.Int64Ptr(int64(26))
				locationModel.End = core.Int64Ptr(int64(26))

				// Construct an instance of the Label model
				labelModel := new(comparecomplyv1.Label)
				labelModel.Nature = core.StringPtr("testString")
				labelModel.Party = core.StringPtr("testString")

				// Construct an instance of the TypeLabel model
				typeLabelModel := new(comparecomplyv1.TypeLabel)
				typeLabelModel.Label = labelModel
				typeLabelModel.ProvenanceIds = []string{"testString"}
				typeLabelModel.Modification = core.StringPtr("added")

				// Construct an instance of the Category model
				categoryModel := new(comparecomplyv1.Category)
				categoryModel.Label = core.StringPtr("Amendments")
				categoryModel.ProvenanceIds = []string{"testString"}
				categoryModel.Modification = core.StringPtr("added")

				// Construct an instance of the OriginalLabelsIn model
				originalLabelsInModel := new(comparecomplyv1.OriginalLabelsIn)
				originalLabelsInModel.Types = []comparecomplyv1.TypeLabel{*typeLabelModel}
				originalLabelsInModel.Categories = []comparecomplyv1.Category{*categoryModel}

				// Construct an instance of the UpdatedLabelsIn model
				updatedLabelsInModel := new(comparecomplyv1.UpdatedLabelsIn)
				updatedLabelsInModel.Types = []comparecomplyv1.TypeLabel{*typeLabelModel}
				updatedLabelsInModel.Categories = []comparecomplyv1.Category{*categoryModel}

				// Construct an instance of the FeedbackDataInput model
				feedbackDataInputModel := new(comparecomplyv1.FeedbackDataInput)
				feedbackDataInputModel.FeedbackType = core.StringPtr("testString")
				feedbackDataInputModel.Document = shortDocModel
				feedbackDataInputModel.ModelID = core.StringPtr("testString")
				feedbackDataInputModel.ModelVersion = core.StringPtr("testString")
				feedbackDataInputModel.Location = locationModel
				feedbackDataInputModel.Text = core.StringPtr("testString")
				feedbackDataInputModel.OriginalLabels = originalLabelsInModel
				feedbackDataInputModel.UpdatedLabels = updatedLabelsInModel

				// Construct an instance of the AddFeedbackOptions model
				addFeedbackOptionsModel := new(comparecomplyv1.AddFeedbackOptions)
				addFeedbackOptionsModel.FeedbackData = feedbackDataInputModel
				addFeedbackOptionsModel.UserID = core.StringPtr("testString")
				addFeedbackOptionsModel.Comment = core.StringPtr("testString")
				addFeedbackOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := compareComplyService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := compareComplyService.AddFeedback(addFeedbackOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the AddFeedbackOptions model with no property values
				addFeedbackOptionsModelNew := new(comparecomplyv1.AddFeedbackOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = compareComplyService.AddFeedback(addFeedbackOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke AddFeedback successfully`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Construct an instance of the ShortDoc model
				shortDocModel := new(comparecomplyv1.ShortDoc)
				shortDocModel.Title = core.StringPtr("testString")
				shortDocModel.Hash = core.StringPtr("testString")

				// Construct an instance of the Location model
				locationModel := new(comparecomplyv1.Location)
				locationModel.Begin = core.Int64Ptr(int64(26))
				locationModel.End = core.Int64Ptr(int64(26))

				// Construct an instance of the Label model
				labelModel := new(comparecomplyv1.Label)
				labelModel.Nature = core.StringPtr("testString")
				labelModel.Party = core.StringPtr("testString")

				// Construct an instance of the TypeLabel model
				typeLabelModel := new(comparecomplyv1.TypeLabel)
				typeLabelModel.Label = labelModel
				typeLabelModel.ProvenanceIds = []string{"testString"}
				typeLabelModel.Modification = core.StringPtr("added")

				// Construct an instance of the Category model
				categoryModel := new(comparecomplyv1.Category)
				categoryModel.Label = core.StringPtr("Amendments")
				categoryModel.ProvenanceIds = []string{"testString"}
				categoryModel.Modification = core.StringPtr("added")

				// Construct an instance of the OriginalLabelsIn model
				originalLabelsInModel := new(comparecomplyv1.OriginalLabelsIn)
				originalLabelsInModel.Types = []comparecomplyv1.TypeLabel{*typeLabelModel}
				originalLabelsInModel.Categories = []comparecomplyv1.Category{*categoryModel}

				// Construct an instance of the UpdatedLabelsIn model
				updatedLabelsInModel := new(comparecomplyv1.UpdatedLabelsIn)
				updatedLabelsInModel.Types = []comparecomplyv1.TypeLabel{*typeLabelModel}
				updatedLabelsInModel.Categories = []comparecomplyv1.Category{*categoryModel}

				// Construct an instance of the FeedbackDataInput model
				feedbackDataInputModel := new(comparecomplyv1.FeedbackDataInput)
				feedbackDataInputModel.FeedbackType = core.StringPtr("testString")
				feedbackDataInputModel.Document = shortDocModel
				feedbackDataInputModel.ModelID = core.StringPtr("testString")
				feedbackDataInputModel.ModelVersion = core.StringPtr("testString")
				feedbackDataInputModel.Location = locationModel
				feedbackDataInputModel.Text = core.StringPtr("testString")
				feedbackDataInputModel.OriginalLabels = originalLabelsInModel
				feedbackDataInputModel.UpdatedLabels = updatedLabelsInModel

				// Construct an instance of the AddFeedbackOptions model
				addFeedbackOptionsModel := new(comparecomplyv1.AddFeedbackOptions)
				addFeedbackOptionsModel.FeedbackData = feedbackDataInputModel
				addFeedbackOptionsModel.UserID = core.StringPtr("testString")
				addFeedbackOptionsModel.Comment = core.StringPtr("testString")
				addFeedbackOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := compareComplyService.AddFeedback(addFeedbackOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListFeedback(listFeedbackOptions *ListFeedbackOptions) - Operation response error`, func() {
		version := "testString"
		listFeedbackPath := "/v1/feedback"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listFeedbackPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["feedback_type"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["document_title"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["model_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["model_version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["category_removed"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["category_added"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["category_not_changed"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["type_removed"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["type_added"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["type_not_changed"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["page_limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["cursor"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_total query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListFeedback with error: Operation response processing error`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Construct an instance of the ListFeedbackOptions model
				listFeedbackOptionsModel := new(comparecomplyv1.ListFeedbackOptions)
				listFeedbackOptionsModel.FeedbackType = core.StringPtr("testString")
				listFeedbackOptionsModel.DocumentTitle = core.StringPtr("testString")
				listFeedbackOptionsModel.ModelID = core.StringPtr("testString")
				listFeedbackOptionsModel.ModelVersion = core.StringPtr("testString")
				listFeedbackOptionsModel.CategoryRemoved = core.StringPtr("testString")
				listFeedbackOptionsModel.CategoryAdded = core.StringPtr("testString")
				listFeedbackOptionsModel.CategoryNotChanged = core.StringPtr("testString")
				listFeedbackOptionsModel.TypeRemoved = core.StringPtr("testString")
				listFeedbackOptionsModel.TypeAdded = core.StringPtr("testString")
				listFeedbackOptionsModel.TypeNotChanged = core.StringPtr("testString")
				listFeedbackOptionsModel.PageLimit = core.Int64Ptr(int64(100))
				listFeedbackOptionsModel.Cursor = core.StringPtr("testString")
				listFeedbackOptionsModel.Sort = core.StringPtr("testString")
				listFeedbackOptionsModel.IncludeTotal = core.BoolPtr(true)
				listFeedbackOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := compareComplyService.ListFeedback(listFeedbackOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				compareComplyService.EnableRetries(0, 0)
				result, response, operationErr = compareComplyService.ListFeedback(listFeedbackOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListFeedback(listFeedbackOptions *ListFeedbackOptions)`, func() {
		version := "testString"
		listFeedbackPath := "/v1/feedback"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listFeedbackPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["feedback_type"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["document_title"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["model_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["model_version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["category_removed"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["category_added"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["category_not_changed"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["type_removed"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["type_added"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["type_not_changed"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["page_limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["cursor"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_total query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"feedback": [{"feedback_id": "FeedbackID", "created": "2019-01-01T12:00:00.000Z", "comment": "Comment", "feedback_data": {"feedback_type": "FeedbackType", "document": {"title": "Title", "hash": "Hash"}, "model_id": "ModelID", "model_version": "ModelVersion", "location": {"begin": 5, "end": 3}, "text": "Text", "original_labels": {"types": [{"label": {"nature": "Nature", "party": "Party"}, "provenance_ids": ["ProvenanceIds"], "modification": "added"}], "categories": [{"label": "Amendments", "provenance_ids": ["ProvenanceIds"], "modification": "added"}]}, "updated_labels": {"types": [{"label": {"nature": "Nature", "party": "Party"}, "provenance_ids": ["ProvenanceIds"], "modification": "added"}], "categories": [{"label": "Amendments", "provenance_ids": ["ProvenanceIds"], "modification": "added"}]}, "pagination": {"refresh_cursor": "RefreshCursor", "next_cursor": "NextCursor", "refresh_url": "RefreshURL", "next_url": "NextURL", "total": 5}}}]}`)
				}))
			})
			It(`Invoke ListFeedback successfully with retries`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())
				compareComplyService.EnableRetries(0, 0)

				// Construct an instance of the ListFeedbackOptions model
				listFeedbackOptionsModel := new(comparecomplyv1.ListFeedbackOptions)
				listFeedbackOptionsModel.FeedbackType = core.StringPtr("testString")
				listFeedbackOptionsModel.DocumentTitle = core.StringPtr("testString")
				listFeedbackOptionsModel.ModelID = core.StringPtr("testString")
				listFeedbackOptionsModel.ModelVersion = core.StringPtr("testString")
				listFeedbackOptionsModel.CategoryRemoved = core.StringPtr("testString")
				listFeedbackOptionsModel.CategoryAdded = core.StringPtr("testString")
				listFeedbackOptionsModel.CategoryNotChanged = core.StringPtr("testString")
				listFeedbackOptionsModel.TypeRemoved = core.StringPtr("testString")
				listFeedbackOptionsModel.TypeAdded = core.StringPtr("testString")
				listFeedbackOptionsModel.TypeNotChanged = core.StringPtr("testString")
				listFeedbackOptionsModel.PageLimit = core.Int64Ptr(int64(100))
				listFeedbackOptionsModel.Cursor = core.StringPtr("testString")
				listFeedbackOptionsModel.Sort = core.StringPtr("testString")
				listFeedbackOptionsModel.IncludeTotal = core.BoolPtr(true)
				listFeedbackOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := compareComplyService.ListFeedbackWithContext(ctx, listFeedbackOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				compareComplyService.DisableRetries()
				result, response, operationErr := compareComplyService.ListFeedback(listFeedbackOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = compareComplyService.ListFeedbackWithContext(ctx, listFeedbackOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listFeedbackPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["feedback_type"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["document_title"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["model_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["model_version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["category_removed"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["category_added"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["category_not_changed"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["type_removed"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["type_added"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["type_not_changed"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["page_limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["cursor"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_total query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"feedback": [{"feedback_id": "FeedbackID", "created": "2019-01-01T12:00:00.000Z", "comment": "Comment", "feedback_data": {"feedback_type": "FeedbackType", "document": {"title": "Title", "hash": "Hash"}, "model_id": "ModelID", "model_version": "ModelVersion", "location": {"begin": 5, "end": 3}, "text": "Text", "original_labels": {"types": [{"label": {"nature": "Nature", "party": "Party"}, "provenance_ids": ["ProvenanceIds"], "modification": "added"}], "categories": [{"label": "Amendments", "provenance_ids": ["ProvenanceIds"], "modification": "added"}]}, "updated_labels": {"types": [{"label": {"nature": "Nature", "party": "Party"}, "provenance_ids": ["ProvenanceIds"], "modification": "added"}], "categories": [{"label": "Amendments", "provenance_ids": ["ProvenanceIds"], "modification": "added"}]}, "pagination": {"refresh_cursor": "RefreshCursor", "next_cursor": "NextCursor", "refresh_url": "RefreshURL", "next_url": "NextURL", "total": 5}}}]}`)
				}))
			})
			It(`Invoke ListFeedback successfully`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := compareComplyService.ListFeedback(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListFeedbackOptions model
				listFeedbackOptionsModel := new(comparecomplyv1.ListFeedbackOptions)
				listFeedbackOptionsModel.FeedbackType = core.StringPtr("testString")
				listFeedbackOptionsModel.DocumentTitle = core.StringPtr("testString")
				listFeedbackOptionsModel.ModelID = core.StringPtr("testString")
				listFeedbackOptionsModel.ModelVersion = core.StringPtr("testString")
				listFeedbackOptionsModel.CategoryRemoved = core.StringPtr("testString")
				listFeedbackOptionsModel.CategoryAdded = core.StringPtr("testString")
				listFeedbackOptionsModel.CategoryNotChanged = core.StringPtr("testString")
				listFeedbackOptionsModel.TypeRemoved = core.StringPtr("testString")
				listFeedbackOptionsModel.TypeAdded = core.StringPtr("testString")
				listFeedbackOptionsModel.TypeNotChanged = core.StringPtr("testString")
				listFeedbackOptionsModel.PageLimit = core.Int64Ptr(int64(100))
				listFeedbackOptionsModel.Cursor = core.StringPtr("testString")
				listFeedbackOptionsModel.Sort = core.StringPtr("testString")
				listFeedbackOptionsModel.IncludeTotal = core.BoolPtr(true)
				listFeedbackOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = compareComplyService.ListFeedback(listFeedbackOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListFeedback with error: Operation request error`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Construct an instance of the ListFeedbackOptions model
				listFeedbackOptionsModel := new(comparecomplyv1.ListFeedbackOptions)
				listFeedbackOptionsModel.FeedbackType = core.StringPtr("testString")
				listFeedbackOptionsModel.DocumentTitle = core.StringPtr("testString")
				listFeedbackOptionsModel.ModelID = core.StringPtr("testString")
				listFeedbackOptionsModel.ModelVersion = core.StringPtr("testString")
				listFeedbackOptionsModel.CategoryRemoved = core.StringPtr("testString")
				listFeedbackOptionsModel.CategoryAdded = core.StringPtr("testString")
				listFeedbackOptionsModel.CategoryNotChanged = core.StringPtr("testString")
				listFeedbackOptionsModel.TypeRemoved = core.StringPtr("testString")
				listFeedbackOptionsModel.TypeAdded = core.StringPtr("testString")
				listFeedbackOptionsModel.TypeNotChanged = core.StringPtr("testString")
				listFeedbackOptionsModel.PageLimit = core.Int64Ptr(int64(100))
				listFeedbackOptionsModel.Cursor = core.StringPtr("testString")
				listFeedbackOptionsModel.Sort = core.StringPtr("testString")
				listFeedbackOptionsModel.IncludeTotal = core.BoolPtr(true)
				listFeedbackOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := compareComplyService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := compareComplyService.ListFeedback(listFeedbackOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListFeedback successfully`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Construct an instance of the ListFeedbackOptions model
				listFeedbackOptionsModel := new(comparecomplyv1.ListFeedbackOptions)
				listFeedbackOptionsModel.FeedbackType = core.StringPtr("testString")
				listFeedbackOptionsModel.DocumentTitle = core.StringPtr("testString")
				listFeedbackOptionsModel.ModelID = core.StringPtr("testString")
				listFeedbackOptionsModel.ModelVersion = core.StringPtr("testString")
				listFeedbackOptionsModel.CategoryRemoved = core.StringPtr("testString")
				listFeedbackOptionsModel.CategoryAdded = core.StringPtr("testString")
				listFeedbackOptionsModel.CategoryNotChanged = core.StringPtr("testString")
				listFeedbackOptionsModel.TypeRemoved = core.StringPtr("testString")
				listFeedbackOptionsModel.TypeAdded = core.StringPtr("testString")
				listFeedbackOptionsModel.TypeNotChanged = core.StringPtr("testString")
				listFeedbackOptionsModel.PageLimit = core.Int64Ptr(int64(100))
				listFeedbackOptionsModel.Cursor = core.StringPtr("testString")
				listFeedbackOptionsModel.Sort = core.StringPtr("testString")
				listFeedbackOptionsModel.IncludeTotal = core.BoolPtr(true)
				listFeedbackOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := compareComplyService.ListFeedback(listFeedbackOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetFeedback(getFeedbackOptions *GetFeedbackOptions) - Operation response error`, func() {
		version := "testString"
		getFeedbackPath := "/v1/feedback/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getFeedbackPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["model"]).To(Equal([]string{"contracts"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetFeedback with error: Operation response processing error`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Construct an instance of the GetFeedbackOptions model
				getFeedbackOptionsModel := new(comparecomplyv1.GetFeedbackOptions)
				getFeedbackOptionsModel.FeedbackID = core.StringPtr("testString")
				getFeedbackOptionsModel.Model = core.StringPtr("contracts")
				getFeedbackOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := compareComplyService.GetFeedback(getFeedbackOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				compareComplyService.EnableRetries(0, 0)
				result, response, operationErr = compareComplyService.GetFeedback(getFeedbackOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetFeedback(getFeedbackOptions *GetFeedbackOptions)`, func() {
		version := "testString"
		getFeedbackPath := "/v1/feedback/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getFeedbackPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["model"]).To(Equal([]string{"contracts"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"feedback_id": "FeedbackID", "created": "2019-01-01T12:00:00.000Z", "comment": "Comment", "feedback_data": {"feedback_type": "FeedbackType", "document": {"title": "Title", "hash": "Hash"}, "model_id": "ModelID", "model_version": "ModelVersion", "location": {"begin": 5, "end": 3}, "text": "Text", "original_labels": {"types": [{"label": {"nature": "Nature", "party": "Party"}, "provenance_ids": ["ProvenanceIds"], "modification": "added"}], "categories": [{"label": "Amendments", "provenance_ids": ["ProvenanceIds"], "modification": "added"}]}, "updated_labels": {"types": [{"label": {"nature": "Nature", "party": "Party"}, "provenance_ids": ["ProvenanceIds"], "modification": "added"}], "categories": [{"label": "Amendments", "provenance_ids": ["ProvenanceIds"], "modification": "added"}]}, "pagination": {"refresh_cursor": "RefreshCursor", "next_cursor": "NextCursor", "refresh_url": "RefreshURL", "next_url": "NextURL", "total": 5}}}`)
				}))
			})
			It(`Invoke GetFeedback successfully with retries`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())
				compareComplyService.EnableRetries(0, 0)

				// Construct an instance of the GetFeedbackOptions model
				getFeedbackOptionsModel := new(comparecomplyv1.GetFeedbackOptions)
				getFeedbackOptionsModel.FeedbackID = core.StringPtr("testString")
				getFeedbackOptionsModel.Model = core.StringPtr("contracts")
				getFeedbackOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := compareComplyService.GetFeedbackWithContext(ctx, getFeedbackOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				compareComplyService.DisableRetries()
				result, response, operationErr := compareComplyService.GetFeedback(getFeedbackOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = compareComplyService.GetFeedbackWithContext(ctx, getFeedbackOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getFeedbackPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["model"]).To(Equal([]string{"contracts"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"feedback_id": "FeedbackID", "created": "2019-01-01T12:00:00.000Z", "comment": "Comment", "feedback_data": {"feedback_type": "FeedbackType", "document": {"title": "Title", "hash": "Hash"}, "model_id": "ModelID", "model_version": "ModelVersion", "location": {"begin": 5, "end": 3}, "text": "Text", "original_labels": {"types": [{"label": {"nature": "Nature", "party": "Party"}, "provenance_ids": ["ProvenanceIds"], "modification": "added"}], "categories": [{"label": "Amendments", "provenance_ids": ["ProvenanceIds"], "modification": "added"}]}, "updated_labels": {"types": [{"label": {"nature": "Nature", "party": "Party"}, "provenance_ids": ["ProvenanceIds"], "modification": "added"}], "categories": [{"label": "Amendments", "provenance_ids": ["ProvenanceIds"], "modification": "added"}]}, "pagination": {"refresh_cursor": "RefreshCursor", "next_cursor": "NextCursor", "refresh_url": "RefreshURL", "next_url": "NextURL", "total": 5}}}`)
				}))
			})
			It(`Invoke GetFeedback successfully`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := compareComplyService.GetFeedback(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetFeedbackOptions model
				getFeedbackOptionsModel := new(comparecomplyv1.GetFeedbackOptions)
				getFeedbackOptionsModel.FeedbackID = core.StringPtr("testString")
				getFeedbackOptionsModel.Model = core.StringPtr("contracts")
				getFeedbackOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = compareComplyService.GetFeedback(getFeedbackOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetFeedback with error: Operation validation and request error`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Construct an instance of the GetFeedbackOptions model
				getFeedbackOptionsModel := new(comparecomplyv1.GetFeedbackOptions)
				getFeedbackOptionsModel.FeedbackID = core.StringPtr("testString")
				getFeedbackOptionsModel.Model = core.StringPtr("contracts")
				getFeedbackOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := compareComplyService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := compareComplyService.GetFeedback(getFeedbackOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetFeedbackOptions model with no property values
				getFeedbackOptionsModelNew := new(comparecomplyv1.GetFeedbackOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = compareComplyService.GetFeedback(getFeedbackOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetFeedback successfully`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Construct an instance of the GetFeedbackOptions model
				getFeedbackOptionsModel := new(comparecomplyv1.GetFeedbackOptions)
				getFeedbackOptionsModel.FeedbackID = core.StringPtr("testString")
				getFeedbackOptionsModel.Model = core.StringPtr("contracts")
				getFeedbackOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := compareComplyService.GetFeedback(getFeedbackOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteFeedback(deleteFeedbackOptions *DeleteFeedbackOptions) - Operation response error`, func() {
		version := "testString"
		deleteFeedbackPath := "/v1/feedback/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteFeedbackPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["model"]).To(Equal([]string{"contracts"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteFeedback with error: Operation response processing error`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Construct an instance of the DeleteFeedbackOptions model
				deleteFeedbackOptionsModel := new(comparecomplyv1.DeleteFeedbackOptions)
				deleteFeedbackOptionsModel.FeedbackID = core.StringPtr("testString")
				deleteFeedbackOptionsModel.Model = core.StringPtr("contracts")
				deleteFeedbackOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := compareComplyService.DeleteFeedback(deleteFeedbackOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				compareComplyService.EnableRetries(0, 0)
				result, response, operationErr = compareComplyService.DeleteFeedback(deleteFeedbackOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteFeedback(deleteFeedbackOptions *DeleteFeedbackOptions)`, func() {
		version := "testString"
		deleteFeedbackPath := "/v1/feedback/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteFeedbackPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["model"]).To(Equal([]string{"contracts"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status": 6, "message": "Message"}`)
				}))
			})
			It(`Invoke DeleteFeedback successfully with retries`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())
				compareComplyService.EnableRetries(0, 0)

				// Construct an instance of the DeleteFeedbackOptions model
				deleteFeedbackOptionsModel := new(comparecomplyv1.DeleteFeedbackOptions)
				deleteFeedbackOptionsModel.FeedbackID = core.StringPtr("testString")
				deleteFeedbackOptionsModel.Model = core.StringPtr("contracts")
				deleteFeedbackOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := compareComplyService.DeleteFeedbackWithContext(ctx, deleteFeedbackOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				compareComplyService.DisableRetries()
				result, response, operationErr := compareComplyService.DeleteFeedback(deleteFeedbackOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = compareComplyService.DeleteFeedbackWithContext(ctx, deleteFeedbackOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteFeedbackPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["model"]).To(Equal([]string{"contracts"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status": 6, "message": "Message"}`)
				}))
			})
			It(`Invoke DeleteFeedback successfully`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := compareComplyService.DeleteFeedback(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteFeedbackOptions model
				deleteFeedbackOptionsModel := new(comparecomplyv1.DeleteFeedbackOptions)
				deleteFeedbackOptionsModel.FeedbackID = core.StringPtr("testString")
				deleteFeedbackOptionsModel.Model = core.StringPtr("contracts")
				deleteFeedbackOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = compareComplyService.DeleteFeedback(deleteFeedbackOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteFeedback with error: Operation validation and request error`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Construct an instance of the DeleteFeedbackOptions model
				deleteFeedbackOptionsModel := new(comparecomplyv1.DeleteFeedbackOptions)
				deleteFeedbackOptionsModel.FeedbackID = core.StringPtr("testString")
				deleteFeedbackOptionsModel.Model = core.StringPtr("contracts")
				deleteFeedbackOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := compareComplyService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := compareComplyService.DeleteFeedback(deleteFeedbackOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteFeedbackOptions model with no property values
				deleteFeedbackOptionsModelNew := new(comparecomplyv1.DeleteFeedbackOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = compareComplyService.DeleteFeedback(deleteFeedbackOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteFeedback successfully`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Construct an instance of the DeleteFeedbackOptions model
				deleteFeedbackOptionsModel := new(comparecomplyv1.DeleteFeedbackOptions)
				deleteFeedbackOptionsModel.FeedbackID = core.StringPtr("testString")
				deleteFeedbackOptionsModel.Model = core.StringPtr("contracts")
				deleteFeedbackOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := compareComplyService.DeleteFeedback(deleteFeedbackOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateBatch(createBatchOptions *CreateBatchOptions) - Operation response error`, func() {
		version := "testString"
		createBatchPath := "/v1/batches"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createBatchPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["function"]).To(Equal([]string{"html_conversion"}))
					Expect(req.URL.Query()["model"]).To(Equal([]string{"contracts"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateBatch with error: Operation response processing error`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Construct an instance of the CreateBatchOptions model
				createBatchOptionsModel := new(comparecomplyv1.CreateBatchOptions)
				createBatchOptionsModel.Function = core.StringPtr("html_conversion")
				createBatchOptionsModel.InputCredentialsFile = CreateMockReader("This is a mock file.")
				createBatchOptionsModel.InputBucketLocation = core.StringPtr("testString")
				createBatchOptionsModel.InputBucketName = core.StringPtr("testString")
				createBatchOptionsModel.OutputCredentialsFile = CreateMockReader("This is a mock file.")
				createBatchOptionsModel.OutputBucketLocation = core.StringPtr("testString")
				createBatchOptionsModel.OutputBucketName = core.StringPtr("testString")
				createBatchOptionsModel.Model = core.StringPtr("contracts")
				createBatchOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := compareComplyService.CreateBatch(createBatchOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				compareComplyService.EnableRetries(0, 0)
				result, response, operationErr = compareComplyService.CreateBatch(createBatchOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateBatch(createBatchOptions *CreateBatchOptions)`, func() {
		version := "testString"
		createBatchPath := "/v1/batches"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createBatchPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["function"]).To(Equal([]string{"html_conversion"}))
					Expect(req.URL.Query()["model"]).To(Equal([]string{"contracts"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"function": "element_classification", "input_bucket_location": "InputBucketLocation", "input_bucket_name": "InputBucketName", "output_bucket_location": "OutputBucketLocation", "output_bucket_name": "OutputBucketName", "batch_id": "BatchID", "document_counts": {"total": 5, "pending": 7, "successful": 10, "failed": 6}, "status": "Status", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke CreateBatch successfully with retries`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())
				compareComplyService.EnableRetries(0, 0)

				// Construct an instance of the CreateBatchOptions model
				createBatchOptionsModel := new(comparecomplyv1.CreateBatchOptions)
				createBatchOptionsModel.Function = core.StringPtr("html_conversion")
				createBatchOptionsModel.InputCredentialsFile = CreateMockReader("This is a mock file.")
				createBatchOptionsModel.InputBucketLocation = core.StringPtr("testString")
				createBatchOptionsModel.InputBucketName = core.StringPtr("testString")
				createBatchOptionsModel.OutputCredentialsFile = CreateMockReader("This is a mock file.")
				createBatchOptionsModel.OutputBucketLocation = core.StringPtr("testString")
				createBatchOptionsModel.OutputBucketName = core.StringPtr("testString")
				createBatchOptionsModel.Model = core.StringPtr("contracts")
				createBatchOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := compareComplyService.CreateBatchWithContext(ctx, createBatchOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				compareComplyService.DisableRetries()
				result, response, operationErr := compareComplyService.CreateBatch(createBatchOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = compareComplyService.CreateBatchWithContext(ctx, createBatchOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createBatchPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["function"]).To(Equal([]string{"html_conversion"}))
					Expect(req.URL.Query()["model"]).To(Equal([]string{"contracts"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"function": "element_classification", "input_bucket_location": "InputBucketLocation", "input_bucket_name": "InputBucketName", "output_bucket_location": "OutputBucketLocation", "output_bucket_name": "OutputBucketName", "batch_id": "BatchID", "document_counts": {"total": 5, "pending": 7, "successful": 10, "failed": 6}, "status": "Status", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke CreateBatch successfully`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := compareComplyService.CreateBatch(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateBatchOptions model
				createBatchOptionsModel := new(comparecomplyv1.CreateBatchOptions)
				createBatchOptionsModel.Function = core.StringPtr("html_conversion")
				createBatchOptionsModel.InputCredentialsFile = CreateMockReader("This is a mock file.")
				createBatchOptionsModel.InputBucketLocation = core.StringPtr("testString")
				createBatchOptionsModel.InputBucketName = core.StringPtr("testString")
				createBatchOptionsModel.OutputCredentialsFile = CreateMockReader("This is a mock file.")
				createBatchOptionsModel.OutputBucketLocation = core.StringPtr("testString")
				createBatchOptionsModel.OutputBucketName = core.StringPtr("testString")
				createBatchOptionsModel.Model = core.StringPtr("contracts")
				createBatchOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = compareComplyService.CreateBatch(createBatchOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateBatch with error: Operation validation and request error`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Construct an instance of the CreateBatchOptions model
				createBatchOptionsModel := new(comparecomplyv1.CreateBatchOptions)
				createBatchOptionsModel.Function = core.StringPtr("html_conversion")
				createBatchOptionsModel.InputCredentialsFile = CreateMockReader("This is a mock file.")
				createBatchOptionsModel.InputBucketLocation = core.StringPtr("testString")
				createBatchOptionsModel.InputBucketName = core.StringPtr("testString")
				createBatchOptionsModel.OutputCredentialsFile = CreateMockReader("This is a mock file.")
				createBatchOptionsModel.OutputBucketLocation = core.StringPtr("testString")
				createBatchOptionsModel.OutputBucketName = core.StringPtr("testString")
				createBatchOptionsModel.Model = core.StringPtr("contracts")
				createBatchOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := compareComplyService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := compareComplyService.CreateBatch(createBatchOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateBatchOptions model with no property values
				createBatchOptionsModelNew := new(comparecomplyv1.CreateBatchOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = compareComplyService.CreateBatch(createBatchOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke CreateBatch successfully`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Construct an instance of the CreateBatchOptions model
				createBatchOptionsModel := new(comparecomplyv1.CreateBatchOptions)
				createBatchOptionsModel.Function = core.StringPtr("html_conversion")
				createBatchOptionsModel.InputCredentialsFile = CreateMockReader("This is a mock file.")
				createBatchOptionsModel.InputBucketLocation = core.StringPtr("testString")
				createBatchOptionsModel.InputBucketName = core.StringPtr("testString")
				createBatchOptionsModel.OutputCredentialsFile = CreateMockReader("This is a mock file.")
				createBatchOptionsModel.OutputBucketLocation = core.StringPtr("testString")
				createBatchOptionsModel.OutputBucketName = core.StringPtr("testString")
				createBatchOptionsModel.Model = core.StringPtr("contracts")
				createBatchOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := compareComplyService.CreateBatch(createBatchOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListBatches(listBatchesOptions *ListBatchesOptions) - Operation response error`, func() {
		version := "testString"
		listBatchesPath := "/v1/batches"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listBatchesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListBatches with error: Operation response processing error`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Construct an instance of the ListBatchesOptions model
				listBatchesOptionsModel := new(comparecomplyv1.ListBatchesOptions)
				listBatchesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := compareComplyService.ListBatches(listBatchesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				compareComplyService.EnableRetries(0, 0)
				result, response, operationErr = compareComplyService.ListBatches(listBatchesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListBatches(listBatchesOptions *ListBatchesOptions)`, func() {
		version := "testString"
		listBatchesPath := "/v1/batches"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listBatchesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"batches": [{"function": "element_classification", "input_bucket_location": "InputBucketLocation", "input_bucket_name": "InputBucketName", "output_bucket_location": "OutputBucketLocation", "output_bucket_name": "OutputBucketName", "batch_id": "BatchID", "document_counts": {"total": 5, "pending": 7, "successful": 10, "failed": 6}, "status": "Status", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke ListBatches successfully with retries`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())
				compareComplyService.EnableRetries(0, 0)

				// Construct an instance of the ListBatchesOptions model
				listBatchesOptionsModel := new(comparecomplyv1.ListBatchesOptions)
				listBatchesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := compareComplyService.ListBatchesWithContext(ctx, listBatchesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				compareComplyService.DisableRetries()
				result, response, operationErr := compareComplyService.ListBatches(listBatchesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = compareComplyService.ListBatchesWithContext(ctx, listBatchesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listBatchesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"batches": [{"function": "element_classification", "input_bucket_location": "InputBucketLocation", "input_bucket_name": "InputBucketName", "output_bucket_location": "OutputBucketLocation", "output_bucket_name": "OutputBucketName", "batch_id": "BatchID", "document_counts": {"total": 5, "pending": 7, "successful": 10, "failed": 6}, "status": "Status", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke ListBatches successfully`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := compareComplyService.ListBatches(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListBatchesOptions model
				listBatchesOptionsModel := new(comparecomplyv1.ListBatchesOptions)
				listBatchesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = compareComplyService.ListBatches(listBatchesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListBatches with error: Operation request error`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Construct an instance of the ListBatchesOptions model
				listBatchesOptionsModel := new(comparecomplyv1.ListBatchesOptions)
				listBatchesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := compareComplyService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := compareComplyService.ListBatches(listBatchesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListBatches successfully`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Construct an instance of the ListBatchesOptions model
				listBatchesOptionsModel := new(comparecomplyv1.ListBatchesOptions)
				listBatchesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := compareComplyService.ListBatches(listBatchesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetBatch(getBatchOptions *GetBatchOptions) - Operation response error`, func() {
		version := "testString"
		getBatchPath := "/v1/batches/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBatchPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetBatch with error: Operation response processing error`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Construct an instance of the GetBatchOptions model
				getBatchOptionsModel := new(comparecomplyv1.GetBatchOptions)
				getBatchOptionsModel.BatchID = core.StringPtr("testString")
				getBatchOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := compareComplyService.GetBatch(getBatchOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				compareComplyService.EnableRetries(0, 0)
				result, response, operationErr = compareComplyService.GetBatch(getBatchOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetBatch(getBatchOptions *GetBatchOptions)`, func() {
		version := "testString"
		getBatchPath := "/v1/batches/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBatchPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"function": "element_classification", "input_bucket_location": "InputBucketLocation", "input_bucket_name": "InputBucketName", "output_bucket_location": "OutputBucketLocation", "output_bucket_name": "OutputBucketName", "batch_id": "BatchID", "document_counts": {"total": 5, "pending": 7, "successful": 10, "failed": 6}, "status": "Status", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetBatch successfully with retries`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())
				compareComplyService.EnableRetries(0, 0)

				// Construct an instance of the GetBatchOptions model
				getBatchOptionsModel := new(comparecomplyv1.GetBatchOptions)
				getBatchOptionsModel.BatchID = core.StringPtr("testString")
				getBatchOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := compareComplyService.GetBatchWithContext(ctx, getBatchOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				compareComplyService.DisableRetries()
				result, response, operationErr := compareComplyService.GetBatch(getBatchOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = compareComplyService.GetBatchWithContext(ctx, getBatchOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBatchPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"function": "element_classification", "input_bucket_location": "InputBucketLocation", "input_bucket_name": "InputBucketName", "output_bucket_location": "OutputBucketLocation", "output_bucket_name": "OutputBucketName", "batch_id": "BatchID", "document_counts": {"total": 5, "pending": 7, "successful": 10, "failed": 6}, "status": "Status", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetBatch successfully`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := compareComplyService.GetBatch(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetBatchOptions model
				getBatchOptionsModel := new(comparecomplyv1.GetBatchOptions)
				getBatchOptionsModel.BatchID = core.StringPtr("testString")
				getBatchOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = compareComplyService.GetBatch(getBatchOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetBatch with error: Operation validation and request error`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Construct an instance of the GetBatchOptions model
				getBatchOptionsModel := new(comparecomplyv1.GetBatchOptions)
				getBatchOptionsModel.BatchID = core.StringPtr("testString")
				getBatchOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := compareComplyService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := compareComplyService.GetBatch(getBatchOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetBatchOptions model with no property values
				getBatchOptionsModelNew := new(comparecomplyv1.GetBatchOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = compareComplyService.GetBatch(getBatchOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetBatch successfully`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Construct an instance of the GetBatchOptions model
				getBatchOptionsModel := new(comparecomplyv1.GetBatchOptions)
				getBatchOptionsModel.BatchID = core.StringPtr("testString")
				getBatchOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := compareComplyService.GetBatch(getBatchOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateBatch(updateBatchOptions *UpdateBatchOptions) - Operation response error`, func() {
		version := "testString"
		updateBatchPath := "/v1/batches/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateBatchPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["action"]).To(Equal([]string{"rescan"}))
					Expect(req.URL.Query()["model"]).To(Equal([]string{"contracts"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateBatch with error: Operation response processing error`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Construct an instance of the UpdateBatchOptions model
				updateBatchOptionsModel := new(comparecomplyv1.UpdateBatchOptions)
				updateBatchOptionsModel.BatchID = core.StringPtr("testString")
				updateBatchOptionsModel.Action = core.StringPtr("rescan")
				updateBatchOptionsModel.Model = core.StringPtr("contracts")
				updateBatchOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := compareComplyService.UpdateBatch(updateBatchOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				compareComplyService.EnableRetries(0, 0)
				result, response, operationErr = compareComplyService.UpdateBatch(updateBatchOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateBatch(updateBatchOptions *UpdateBatchOptions)`, func() {
		version := "testString"
		updateBatchPath := "/v1/batches/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateBatchPath))
					Expect(req.Method).To(Equal("PUT"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["action"]).To(Equal([]string{"rescan"}))
					Expect(req.URL.Query()["model"]).To(Equal([]string{"contracts"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"function": "element_classification", "input_bucket_location": "InputBucketLocation", "input_bucket_name": "InputBucketName", "output_bucket_location": "OutputBucketLocation", "output_bucket_name": "OutputBucketName", "batch_id": "BatchID", "document_counts": {"total": 5, "pending": 7, "successful": 10, "failed": 6}, "status": "Status", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke UpdateBatch successfully with retries`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())
				compareComplyService.EnableRetries(0, 0)

				// Construct an instance of the UpdateBatchOptions model
				updateBatchOptionsModel := new(comparecomplyv1.UpdateBatchOptions)
				updateBatchOptionsModel.BatchID = core.StringPtr("testString")
				updateBatchOptionsModel.Action = core.StringPtr("rescan")
				updateBatchOptionsModel.Model = core.StringPtr("contracts")
				updateBatchOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := compareComplyService.UpdateBatchWithContext(ctx, updateBatchOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				compareComplyService.DisableRetries()
				result, response, operationErr := compareComplyService.UpdateBatch(updateBatchOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = compareComplyService.UpdateBatchWithContext(ctx, updateBatchOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateBatchPath))
					Expect(req.Method).To(Equal("PUT"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["action"]).To(Equal([]string{"rescan"}))
					Expect(req.URL.Query()["model"]).To(Equal([]string{"contracts"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"function": "element_classification", "input_bucket_location": "InputBucketLocation", "input_bucket_name": "InputBucketName", "output_bucket_location": "OutputBucketLocation", "output_bucket_name": "OutputBucketName", "batch_id": "BatchID", "document_counts": {"total": 5, "pending": 7, "successful": 10, "failed": 6}, "status": "Status", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke UpdateBatch successfully`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := compareComplyService.UpdateBatch(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateBatchOptions model
				updateBatchOptionsModel := new(comparecomplyv1.UpdateBatchOptions)
				updateBatchOptionsModel.BatchID = core.StringPtr("testString")
				updateBatchOptionsModel.Action = core.StringPtr("rescan")
				updateBatchOptionsModel.Model = core.StringPtr("contracts")
				updateBatchOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = compareComplyService.UpdateBatch(updateBatchOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateBatch with error: Operation validation and request error`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Construct an instance of the UpdateBatchOptions model
				updateBatchOptionsModel := new(comparecomplyv1.UpdateBatchOptions)
				updateBatchOptionsModel.BatchID = core.StringPtr("testString")
				updateBatchOptionsModel.Action = core.StringPtr("rescan")
				updateBatchOptionsModel.Model = core.StringPtr("contracts")
				updateBatchOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := compareComplyService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := compareComplyService.UpdateBatch(updateBatchOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateBatchOptions model with no property values
				updateBatchOptionsModelNew := new(comparecomplyv1.UpdateBatchOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = compareComplyService.UpdateBatch(updateBatchOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateBatch successfully`, func() {
				compareComplyService, serviceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(compareComplyService).ToNot(BeNil())

				// Construct an instance of the UpdateBatchOptions model
				updateBatchOptionsModel := new(comparecomplyv1.UpdateBatchOptions)
				updateBatchOptionsModel.BatchID = core.StringPtr("testString")
				updateBatchOptionsModel.Action = core.StringPtr("rescan")
				updateBatchOptionsModel.Model = core.StringPtr("contracts")
				updateBatchOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := compareComplyService.UpdateBatch(updateBatchOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			version := "testString"
			compareComplyService, _ := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
				URL:           "http://comparecomplyv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       core.StringPtr(version),
			})
			It(`Invoke NewAddFeedbackOptions successfully`, func() {
				// Construct an instance of the ShortDoc model
				shortDocModel := new(comparecomplyv1.ShortDoc)
				Expect(shortDocModel).ToNot(BeNil())
				shortDocModel.Title = core.StringPtr("testString")
				shortDocModel.Hash = core.StringPtr("testString")
				Expect(shortDocModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(shortDocModel.Hash).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Location model
				locationModel := new(comparecomplyv1.Location)
				Expect(locationModel).ToNot(BeNil())
				locationModel.Begin = core.Int64Ptr(int64(26))
				locationModel.End = core.Int64Ptr(int64(26))
				Expect(locationModel.Begin).To(Equal(core.Int64Ptr(int64(26))))
				Expect(locationModel.End).To(Equal(core.Int64Ptr(int64(26))))

				// Construct an instance of the Label model
				labelModel := new(comparecomplyv1.Label)
				Expect(labelModel).ToNot(BeNil())
				labelModel.Nature = core.StringPtr("testString")
				labelModel.Party = core.StringPtr("testString")
				Expect(labelModel.Nature).To(Equal(core.StringPtr("testString")))
				Expect(labelModel.Party).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the TypeLabel model
				typeLabelModel := new(comparecomplyv1.TypeLabel)
				Expect(typeLabelModel).ToNot(BeNil())
				typeLabelModel.Label = labelModel
				typeLabelModel.ProvenanceIds = []string{"testString"}
				typeLabelModel.Modification = core.StringPtr("added")
				Expect(typeLabelModel.Label).To(Equal(labelModel))
				Expect(typeLabelModel.ProvenanceIds).To(Equal([]string{"testString"}))
				Expect(typeLabelModel.Modification).To(Equal(core.StringPtr("added")))

				// Construct an instance of the Category model
				categoryModel := new(comparecomplyv1.Category)
				Expect(categoryModel).ToNot(BeNil())
				categoryModel.Label = core.StringPtr("Amendments")
				categoryModel.ProvenanceIds = []string{"testString"}
				categoryModel.Modification = core.StringPtr("added")
				Expect(categoryModel.Label).To(Equal(core.StringPtr("Amendments")))
				Expect(categoryModel.ProvenanceIds).To(Equal([]string{"testString"}))
				Expect(categoryModel.Modification).To(Equal(core.StringPtr("added")))

				// Construct an instance of the OriginalLabelsIn model
				originalLabelsInModel := new(comparecomplyv1.OriginalLabelsIn)
				Expect(originalLabelsInModel).ToNot(BeNil())
				originalLabelsInModel.Types = []comparecomplyv1.TypeLabel{*typeLabelModel}
				originalLabelsInModel.Categories = []comparecomplyv1.Category{*categoryModel}
				Expect(originalLabelsInModel.Types).To(Equal([]comparecomplyv1.TypeLabel{*typeLabelModel}))
				Expect(originalLabelsInModel.Categories).To(Equal([]comparecomplyv1.Category{*categoryModel}))

				// Construct an instance of the UpdatedLabelsIn model
				updatedLabelsInModel := new(comparecomplyv1.UpdatedLabelsIn)
				Expect(updatedLabelsInModel).ToNot(BeNil())
				updatedLabelsInModel.Types = []comparecomplyv1.TypeLabel{*typeLabelModel}
				updatedLabelsInModel.Categories = []comparecomplyv1.Category{*categoryModel}
				Expect(updatedLabelsInModel.Types).To(Equal([]comparecomplyv1.TypeLabel{*typeLabelModel}))
				Expect(updatedLabelsInModel.Categories).To(Equal([]comparecomplyv1.Category{*categoryModel}))

				// Construct an instance of the FeedbackDataInput model
				feedbackDataInputModel := new(comparecomplyv1.FeedbackDataInput)
				Expect(feedbackDataInputModel).ToNot(BeNil())
				feedbackDataInputModel.FeedbackType = core.StringPtr("testString")
				feedbackDataInputModel.Document = shortDocModel
				feedbackDataInputModel.ModelID = core.StringPtr("testString")
				feedbackDataInputModel.ModelVersion = core.StringPtr("testString")
				feedbackDataInputModel.Location = locationModel
				feedbackDataInputModel.Text = core.StringPtr("testString")
				feedbackDataInputModel.OriginalLabels = originalLabelsInModel
				feedbackDataInputModel.UpdatedLabels = updatedLabelsInModel
				Expect(feedbackDataInputModel.FeedbackType).To(Equal(core.StringPtr("testString")))
				Expect(feedbackDataInputModel.Document).To(Equal(shortDocModel))
				Expect(feedbackDataInputModel.ModelID).To(Equal(core.StringPtr("testString")))
				Expect(feedbackDataInputModel.ModelVersion).To(Equal(core.StringPtr("testString")))
				Expect(feedbackDataInputModel.Location).To(Equal(locationModel))
				Expect(feedbackDataInputModel.Text).To(Equal(core.StringPtr("testString")))
				Expect(feedbackDataInputModel.OriginalLabels).To(Equal(originalLabelsInModel))
				Expect(feedbackDataInputModel.UpdatedLabels).To(Equal(updatedLabelsInModel))

				// Construct an instance of the AddFeedbackOptions model
				var addFeedbackOptionsFeedbackData *comparecomplyv1.FeedbackDataInput = nil
				addFeedbackOptionsModel := compareComplyService.NewAddFeedbackOptions(addFeedbackOptionsFeedbackData)
				addFeedbackOptionsModel.SetFeedbackData(feedbackDataInputModel)
				addFeedbackOptionsModel.SetUserID("testString")
				addFeedbackOptionsModel.SetComment("testString")
				addFeedbackOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addFeedbackOptionsModel).ToNot(BeNil())
				Expect(addFeedbackOptionsModel.FeedbackData).To(Equal(feedbackDataInputModel))
				Expect(addFeedbackOptionsModel.UserID).To(Equal(core.StringPtr("testString")))
				Expect(addFeedbackOptionsModel.Comment).To(Equal(core.StringPtr("testString")))
				Expect(addFeedbackOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewClassifyElementsOptions successfully`, func() {
				// Construct an instance of the ClassifyElementsOptions model
				file := CreateMockReader("This is a mock file.")
				classifyElementsOptionsModel := compareComplyService.NewClassifyElementsOptions(file)
				classifyElementsOptionsModel.SetFile(CreateMockReader("This is a mock file."))
				classifyElementsOptionsModel.SetFileContentType("application/pdf")
				classifyElementsOptionsModel.SetModel("contracts")
				classifyElementsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(classifyElementsOptionsModel).ToNot(BeNil())
				Expect(classifyElementsOptionsModel.File).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(classifyElementsOptionsModel.FileContentType).To(Equal(core.StringPtr("application/pdf")))
				Expect(classifyElementsOptionsModel.Model).To(Equal(core.StringPtr("contracts")))
				Expect(classifyElementsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCompareDocumentsOptions successfully`, func() {
				// Construct an instance of the CompareDocumentsOptions model
				file1 := CreateMockReader("This is a mock file.")
				file2 := CreateMockReader("This is a mock file.")
				compareDocumentsOptionsModel := compareComplyService.NewCompareDocumentsOptions(file1, file2)
				compareDocumentsOptionsModel.SetFile1(CreateMockReader("This is a mock file."))
				compareDocumentsOptionsModel.SetFile2(CreateMockReader("This is a mock file."))
				compareDocumentsOptionsModel.SetFile1ContentType("application/pdf")
				compareDocumentsOptionsModel.SetFile2ContentType("application/pdf")
				compareDocumentsOptionsModel.SetFile1Label("file_1")
				compareDocumentsOptionsModel.SetFile2Label("file_2")
				compareDocumentsOptionsModel.SetModel("contracts")
				compareDocumentsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(compareDocumentsOptionsModel).ToNot(BeNil())
				Expect(compareDocumentsOptionsModel.File1).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(compareDocumentsOptionsModel.File2).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(compareDocumentsOptionsModel.File1ContentType).To(Equal(core.StringPtr("application/pdf")))
				Expect(compareDocumentsOptionsModel.File2ContentType).To(Equal(core.StringPtr("application/pdf")))
				Expect(compareDocumentsOptionsModel.File1Label).To(Equal(core.StringPtr("file_1")))
				Expect(compareDocumentsOptionsModel.File2Label).To(Equal(core.StringPtr("file_2")))
				Expect(compareDocumentsOptionsModel.Model).To(Equal(core.StringPtr("contracts")))
				Expect(compareDocumentsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewConvertToHTMLOptions successfully`, func() {
				// Construct an instance of the ConvertToHTMLOptions model
				file := CreateMockReader("This is a mock file.")
				convertToHTMLOptionsModel := compareComplyService.NewConvertToHTMLOptions(file)
				convertToHTMLOptionsModel.SetFile(CreateMockReader("This is a mock file."))
				convertToHTMLOptionsModel.SetFileContentType("application/pdf")
				convertToHTMLOptionsModel.SetModel("contracts")
				convertToHTMLOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(convertToHTMLOptionsModel).ToNot(BeNil())
				Expect(convertToHTMLOptionsModel.File).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(convertToHTMLOptionsModel.FileContentType).To(Equal(core.StringPtr("application/pdf")))
				Expect(convertToHTMLOptionsModel.Model).To(Equal(core.StringPtr("contracts")))
				Expect(convertToHTMLOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateBatchOptions successfully`, func() {
				// Construct an instance of the CreateBatchOptions model
				function := "html_conversion"
				inputCredentialsFile := CreateMockReader("This is a mock file.")
				inputBucketLocation := "testString"
				inputBucketName := "testString"
				outputCredentialsFile := CreateMockReader("This is a mock file.")
				outputBucketLocation := "testString"
				outputBucketName := "testString"
				createBatchOptionsModel := compareComplyService.NewCreateBatchOptions(function, inputCredentialsFile, inputBucketLocation, inputBucketName, outputCredentialsFile, outputBucketLocation, outputBucketName)
				createBatchOptionsModel.SetFunction("html_conversion")
				createBatchOptionsModel.SetInputCredentialsFile(CreateMockReader("This is a mock file."))
				createBatchOptionsModel.SetInputBucketLocation("testString")
				createBatchOptionsModel.SetInputBucketName("testString")
				createBatchOptionsModel.SetOutputCredentialsFile(CreateMockReader("This is a mock file."))
				createBatchOptionsModel.SetOutputBucketLocation("testString")
				createBatchOptionsModel.SetOutputBucketName("testString")
				createBatchOptionsModel.SetModel("contracts")
				createBatchOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createBatchOptionsModel).ToNot(BeNil())
				Expect(createBatchOptionsModel.Function).To(Equal(core.StringPtr("html_conversion")))
				Expect(createBatchOptionsModel.InputCredentialsFile).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(createBatchOptionsModel.InputBucketLocation).To(Equal(core.StringPtr("testString")))
				Expect(createBatchOptionsModel.InputBucketName).To(Equal(core.StringPtr("testString")))
				Expect(createBatchOptionsModel.OutputCredentialsFile).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(createBatchOptionsModel.OutputBucketLocation).To(Equal(core.StringPtr("testString")))
				Expect(createBatchOptionsModel.OutputBucketName).To(Equal(core.StringPtr("testString")))
				Expect(createBatchOptionsModel.Model).To(Equal(core.StringPtr("contracts")))
				Expect(createBatchOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteFeedbackOptions successfully`, func() {
				// Construct an instance of the DeleteFeedbackOptions model
				feedbackID := "testString"
				deleteFeedbackOptionsModel := compareComplyService.NewDeleteFeedbackOptions(feedbackID)
				deleteFeedbackOptionsModel.SetFeedbackID("testString")
				deleteFeedbackOptionsModel.SetModel("contracts")
				deleteFeedbackOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteFeedbackOptionsModel).ToNot(BeNil())
				Expect(deleteFeedbackOptionsModel.FeedbackID).To(Equal(core.StringPtr("testString")))
				Expect(deleteFeedbackOptionsModel.Model).To(Equal(core.StringPtr("contracts")))
				Expect(deleteFeedbackOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewExtractTablesOptions successfully`, func() {
				// Construct an instance of the ExtractTablesOptions model
				file := CreateMockReader("This is a mock file.")
				extractTablesOptionsModel := compareComplyService.NewExtractTablesOptions(file)
				extractTablesOptionsModel.SetFile(CreateMockReader("This is a mock file."))
				extractTablesOptionsModel.SetFileContentType("application/pdf")
				extractTablesOptionsModel.SetModel("contracts")
				extractTablesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(extractTablesOptionsModel).ToNot(BeNil())
				Expect(extractTablesOptionsModel.File).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(extractTablesOptionsModel.FileContentType).To(Equal(core.StringPtr("application/pdf")))
				Expect(extractTablesOptionsModel.Model).To(Equal(core.StringPtr("contracts")))
				Expect(extractTablesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewFeedbackDataInput successfully`, func() {
				feedbackType := "testString"
				var location *comparecomplyv1.Location = nil
				text := "testString"
				var originalLabels *comparecomplyv1.OriginalLabelsIn = nil
				var updatedLabels *comparecomplyv1.UpdatedLabelsIn = nil
				_, err := compareComplyService.NewFeedbackDataInput(feedbackType, location, text, originalLabels, updatedLabels)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewGetBatchOptions successfully`, func() {
				// Construct an instance of the GetBatchOptions model
				batchID := "testString"
				getBatchOptionsModel := compareComplyService.NewGetBatchOptions(batchID)
				getBatchOptionsModel.SetBatchID("testString")
				getBatchOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getBatchOptionsModel).ToNot(BeNil())
				Expect(getBatchOptionsModel.BatchID).To(Equal(core.StringPtr("testString")))
				Expect(getBatchOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetFeedbackOptions successfully`, func() {
				// Construct an instance of the GetFeedbackOptions model
				feedbackID := "testString"
				getFeedbackOptionsModel := compareComplyService.NewGetFeedbackOptions(feedbackID)
				getFeedbackOptionsModel.SetFeedbackID("testString")
				getFeedbackOptionsModel.SetModel("contracts")
				getFeedbackOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getFeedbackOptionsModel).ToNot(BeNil())
				Expect(getFeedbackOptionsModel.FeedbackID).To(Equal(core.StringPtr("testString")))
				Expect(getFeedbackOptionsModel.Model).To(Equal(core.StringPtr("contracts")))
				Expect(getFeedbackOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewLabel successfully`, func() {
				nature := "testString"
				party := "testString"
				_model, err := compareComplyService.NewLabel(nature, party)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewListBatchesOptions successfully`, func() {
				// Construct an instance of the ListBatchesOptions model
				listBatchesOptionsModel := compareComplyService.NewListBatchesOptions()
				listBatchesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listBatchesOptionsModel).ToNot(BeNil())
				Expect(listBatchesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListFeedbackOptions successfully`, func() {
				// Construct an instance of the ListFeedbackOptions model
				listFeedbackOptionsModel := compareComplyService.NewListFeedbackOptions()
				listFeedbackOptionsModel.SetFeedbackType("testString")
				listFeedbackOptionsModel.SetDocumentTitle("testString")
				listFeedbackOptionsModel.SetModelID("testString")
				listFeedbackOptionsModel.SetModelVersion("testString")
				listFeedbackOptionsModel.SetCategoryRemoved("testString")
				listFeedbackOptionsModel.SetCategoryAdded("testString")
				listFeedbackOptionsModel.SetCategoryNotChanged("testString")
				listFeedbackOptionsModel.SetTypeRemoved("testString")
				listFeedbackOptionsModel.SetTypeAdded("testString")
				listFeedbackOptionsModel.SetTypeNotChanged("testString")
				listFeedbackOptionsModel.SetPageLimit(int64(100))
				listFeedbackOptionsModel.SetCursor("testString")
				listFeedbackOptionsModel.SetSort("testString")
				listFeedbackOptionsModel.SetIncludeTotal(true)
				listFeedbackOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listFeedbackOptionsModel).ToNot(BeNil())
				Expect(listFeedbackOptionsModel.FeedbackType).To(Equal(core.StringPtr("testString")))
				Expect(listFeedbackOptionsModel.DocumentTitle).To(Equal(core.StringPtr("testString")))
				Expect(listFeedbackOptionsModel.ModelID).To(Equal(core.StringPtr("testString")))
				Expect(listFeedbackOptionsModel.ModelVersion).To(Equal(core.StringPtr("testString")))
				Expect(listFeedbackOptionsModel.CategoryRemoved).To(Equal(core.StringPtr("testString")))
				Expect(listFeedbackOptionsModel.CategoryAdded).To(Equal(core.StringPtr("testString")))
				Expect(listFeedbackOptionsModel.CategoryNotChanged).To(Equal(core.StringPtr("testString")))
				Expect(listFeedbackOptionsModel.TypeRemoved).To(Equal(core.StringPtr("testString")))
				Expect(listFeedbackOptionsModel.TypeAdded).To(Equal(core.StringPtr("testString")))
				Expect(listFeedbackOptionsModel.TypeNotChanged).To(Equal(core.StringPtr("testString")))
				Expect(listFeedbackOptionsModel.PageLimit).To(Equal(core.Int64Ptr(int64(100))))
				Expect(listFeedbackOptionsModel.Cursor).To(Equal(core.StringPtr("testString")))
				Expect(listFeedbackOptionsModel.Sort).To(Equal(core.StringPtr("testString")))
				Expect(listFeedbackOptionsModel.IncludeTotal).To(Equal(core.BoolPtr(true)))
				Expect(listFeedbackOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewLocation successfully`, func() {
				begin := int64(26)
				end := int64(26)
				_model, err := compareComplyService.NewLocation(begin, end)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewOriginalLabelsIn successfully`, func() {
				types := []comparecomplyv1.TypeLabel{}
				categories := []comparecomplyv1.Category{}
				_model, err := compareComplyService.NewOriginalLabelsIn(types, categories)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdateBatchOptions successfully`, func() {
				// Construct an instance of the UpdateBatchOptions model
				batchID := "testString"
				action := "rescan"
				updateBatchOptionsModel := compareComplyService.NewUpdateBatchOptions(batchID, action)
				updateBatchOptionsModel.SetBatchID("testString")
				updateBatchOptionsModel.SetAction("rescan")
				updateBatchOptionsModel.SetModel("contracts")
				updateBatchOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateBatchOptionsModel).ToNot(BeNil())
				Expect(updateBatchOptionsModel.BatchID).To(Equal(core.StringPtr("testString")))
				Expect(updateBatchOptionsModel.Action).To(Equal(core.StringPtr("rescan")))
				Expect(updateBatchOptionsModel.Model).To(Equal(core.StringPtr("contracts")))
				Expect(updateBatchOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdatedLabelsIn successfully`, func() {
				types := []comparecomplyv1.TypeLabel{}
				categories := []comparecomplyv1.Category{}
				_model, err := compareComplyService.NewUpdatedLabelsIn(types, categories)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
