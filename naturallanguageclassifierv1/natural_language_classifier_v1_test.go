/**
 * (C) Copyright IBM Corp. 2018, 2020.
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

package naturallanguageclassifierv1_test

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

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/naturallanguageclassifierv1"
)

var _ = Describe(`NaturalLanguageClassifierV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			naturalLanguageClassifierService, serviceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(naturalLanguageClassifierService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
			Expect(naturalLanguageClassifierService.Service.IsSSLDisabled()).To(BeFalse())
			naturalLanguageClassifierService.DisableSSLVerification()
			Expect(naturalLanguageClassifierService.Service.IsSSLDisabled()).To(BeTrue())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			naturalLanguageClassifierService, serviceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(naturalLanguageClassifierService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			naturalLanguageClassifierService, serviceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
				URL: "https://naturallanguageclassifierv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(naturalLanguageClassifierService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"NATURAL_LANGUAGE_CLASSIFIER_URL":       "https://naturallanguageclassifierv1/api",
				"NATURAL_LANGUAGE_CLASSIFIER_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				naturalLanguageClassifierService, serviceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{})
				Expect(naturalLanguageClassifierService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				naturalLanguageClassifierService, serviceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
					URL: "https://testService/api",
				})
				Expect(naturalLanguageClassifierService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageClassifierService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				naturalLanguageClassifierService, serviceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{})
				err := naturalLanguageClassifierService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(naturalLanguageClassifierService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageClassifierService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"NATURAL_LANGUAGE_CLASSIFIER_URL":       "https://naturallanguageclassifierv1/api",
				"NATURAL_LANGUAGE_CLASSIFIER_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			naturalLanguageClassifierService, serviceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(naturalLanguageClassifierService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"NATURAL_LANGUAGE_CLASSIFIER_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			naturalLanguageClassifierService, serviceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(naturalLanguageClassifierService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Classify(classifyOptions *ClassifyOptions) - Operation response error`, func() {
		classifyPath := "/v1/classifiers/testString/classify"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(classifyPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke Classify with error: Operation response processing error`, func() {
				naturalLanguageClassifierService, serviceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageClassifierService).ToNot(BeNil())

				// Construct an instance of the ClassifyOptions model
				classifyOptionsModel := new(naturallanguageclassifierv1.ClassifyOptions)
				classifyOptionsModel.ClassifierID = core.StringPtr("testString")
				classifyOptionsModel.Text = core.StringPtr("testString")
				classifyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := naturalLanguageClassifierService.Classify(classifyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				naturalLanguageClassifierService.EnableRetries(0, 0)
				result, response, operationErr = naturalLanguageClassifierService.Classify(classifyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`Classify(classifyOptions *ClassifyOptions)`, func() {
		classifyPath := "/v1/classifiers/testString/classify"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(classifyPath))
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

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"classifier_id": "ClassifierID", "url": "URL", "text": "Text", "top_class": "TopClass", "classes": [{"confidence": 10, "class_name": "ClassName"}]}`)
				}))
			})
			It(`Invoke Classify successfully`, func() {
				naturalLanguageClassifierService, serviceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageClassifierService).ToNot(BeNil())
				naturalLanguageClassifierService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := naturalLanguageClassifierService.Classify(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ClassifyOptions model
				classifyOptionsModel := new(naturallanguageclassifierv1.ClassifyOptions)
				classifyOptionsModel.ClassifierID = core.StringPtr("testString")
				classifyOptionsModel.Text = core.StringPtr("testString")
				classifyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = naturalLanguageClassifierService.Classify(classifyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = naturalLanguageClassifierService.ClassifyWithContext(ctx, classifyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				naturalLanguageClassifierService.DisableRetries()
				result, response, operationErr = naturalLanguageClassifierService.Classify(classifyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = naturalLanguageClassifierService.ClassifyWithContext(ctx, classifyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke Classify with error: Operation validation and request error`, func() {
				naturalLanguageClassifierService, serviceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageClassifierService).ToNot(BeNil())

				// Construct an instance of the ClassifyOptions model
				classifyOptionsModel := new(naturallanguageclassifierv1.ClassifyOptions)
				classifyOptionsModel.ClassifierID = core.StringPtr("testString")
				classifyOptionsModel.Text = core.StringPtr("testString")
				classifyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := naturalLanguageClassifierService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := naturalLanguageClassifierService.Classify(classifyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ClassifyOptions model with no property values
				classifyOptionsModelNew := new(naturallanguageclassifierv1.ClassifyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = naturalLanguageClassifierService.Classify(classifyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ClassifyCollection(classifyCollectionOptions *ClassifyCollectionOptions) - Operation response error`, func() {
		classifyCollectionPath := "/v1/classifiers/testString/classify_collection"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(classifyCollectionPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ClassifyCollection with error: Operation response processing error`, func() {
				naturalLanguageClassifierService, serviceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageClassifierService).ToNot(BeNil())

				// Construct an instance of the ClassifyInput model
				classifyInputModel := new(naturallanguageclassifierv1.ClassifyInput)
				classifyInputModel.Text = core.StringPtr("How hot will it be today?")

				// Construct an instance of the ClassifyCollectionOptions model
				classifyCollectionOptionsModel := new(naturallanguageclassifierv1.ClassifyCollectionOptions)
				classifyCollectionOptionsModel.ClassifierID = core.StringPtr("testString")
				classifyCollectionOptionsModel.Collection = []naturallanguageclassifierv1.ClassifyInput{*classifyInputModel}
				classifyCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := naturalLanguageClassifierService.ClassifyCollection(classifyCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				naturalLanguageClassifierService.EnableRetries(0, 0)
				result, response, operationErr = naturalLanguageClassifierService.ClassifyCollection(classifyCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ClassifyCollection(classifyCollectionOptions *ClassifyCollectionOptions)`, func() {
		classifyCollectionPath := "/v1/classifiers/testString/classify_collection"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(classifyCollectionPath))
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

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"classifier_id": "ClassifierID", "url": "URL", "collection": [{"text": "Text", "top_class": "TopClass", "classes": [{"confidence": 10, "class_name": "ClassName"}]}]}`)
				}))
			})
			It(`Invoke ClassifyCollection successfully`, func() {
				naturalLanguageClassifierService, serviceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageClassifierService).ToNot(BeNil())
				naturalLanguageClassifierService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := naturalLanguageClassifierService.ClassifyCollection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ClassifyInput model
				classifyInputModel := new(naturallanguageclassifierv1.ClassifyInput)
				classifyInputModel.Text = core.StringPtr("How hot will it be today?")

				// Construct an instance of the ClassifyCollectionOptions model
				classifyCollectionOptionsModel := new(naturallanguageclassifierv1.ClassifyCollectionOptions)
				classifyCollectionOptionsModel.ClassifierID = core.StringPtr("testString")
				classifyCollectionOptionsModel.Collection = []naturallanguageclassifierv1.ClassifyInput{*classifyInputModel}
				classifyCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = naturalLanguageClassifierService.ClassifyCollection(classifyCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = naturalLanguageClassifierService.ClassifyCollectionWithContext(ctx, classifyCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				naturalLanguageClassifierService.DisableRetries()
				result, response, operationErr = naturalLanguageClassifierService.ClassifyCollection(classifyCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = naturalLanguageClassifierService.ClassifyCollectionWithContext(ctx, classifyCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ClassifyCollection with error: Operation validation and request error`, func() {
				naturalLanguageClassifierService, serviceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageClassifierService).ToNot(BeNil())

				// Construct an instance of the ClassifyInput model
				classifyInputModel := new(naturallanguageclassifierv1.ClassifyInput)
				classifyInputModel.Text = core.StringPtr("How hot will it be today?")

				// Construct an instance of the ClassifyCollectionOptions model
				classifyCollectionOptionsModel := new(naturallanguageclassifierv1.ClassifyCollectionOptions)
				classifyCollectionOptionsModel.ClassifierID = core.StringPtr("testString")
				classifyCollectionOptionsModel.Collection = []naturallanguageclassifierv1.ClassifyInput{*classifyInputModel}
				classifyCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := naturalLanguageClassifierService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := naturalLanguageClassifierService.ClassifyCollection(classifyCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ClassifyCollectionOptions model with no property values
				classifyCollectionOptionsModelNew := new(naturallanguageclassifierv1.ClassifyCollectionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = naturalLanguageClassifierService.ClassifyCollection(classifyCollectionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			naturalLanguageClassifierService, serviceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(naturalLanguageClassifierService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
			Expect(naturalLanguageClassifierService.Service.IsSSLDisabled()).To(BeFalse())
			naturalLanguageClassifierService.DisableSSLVerification()
			Expect(naturalLanguageClassifierService.Service.IsSSLDisabled()).To(BeTrue())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			naturalLanguageClassifierService, serviceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(naturalLanguageClassifierService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			naturalLanguageClassifierService, serviceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
				URL: "https://naturallanguageclassifierv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(naturalLanguageClassifierService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"NATURAL_LANGUAGE_CLASSIFIER_URL":       "https://naturallanguageclassifierv1/api",
				"NATURAL_LANGUAGE_CLASSIFIER_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				naturalLanguageClassifierService, serviceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{})
				Expect(naturalLanguageClassifierService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				naturalLanguageClassifierService, serviceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
					URL: "https://testService/api",
				})
				Expect(naturalLanguageClassifierService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageClassifierService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				naturalLanguageClassifierService, serviceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{})
				err := naturalLanguageClassifierService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(naturalLanguageClassifierService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageClassifierService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"NATURAL_LANGUAGE_CLASSIFIER_URL":       "https://naturallanguageclassifierv1/api",
				"NATURAL_LANGUAGE_CLASSIFIER_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			naturalLanguageClassifierService, serviceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(naturalLanguageClassifierService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"NATURAL_LANGUAGE_CLASSIFIER_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			naturalLanguageClassifierService, serviceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(naturalLanguageClassifierService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`CreateClassifier(createClassifierOptions *CreateClassifierOptions) - Operation response error`, func() {
		createClassifierPath := "/v1/classifiers"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createClassifierPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateClassifier with error: Operation response processing error`, func() {
				naturalLanguageClassifierService, serviceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageClassifierService).ToNot(BeNil())

				// Construct an instance of the CreateClassifierOptions model
				createClassifierOptionsModel := new(naturallanguageclassifierv1.CreateClassifierOptions)
				createClassifierOptionsModel.TrainingMetadata = CreateMockReader("This is a mock file.")
				createClassifierOptionsModel.TrainingData = CreateMockReader("This is a mock file.")
				createClassifierOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := naturalLanguageClassifierService.CreateClassifier(createClassifierOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				naturalLanguageClassifierService.EnableRetries(0, 0)
				result, response, operationErr = naturalLanguageClassifierService.CreateClassifier(createClassifierOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateClassifier(createClassifierOptions *CreateClassifierOptions)`, func() {
		createClassifierPath := "/v1/classifiers"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createClassifierPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "url": "URL", "status": "Non Existent", "classifier_id": "ClassifierID", "created": "2019-01-01T12:00:00", "status_description": "StatusDescription", "language": "Language"}`)
				}))
			})
			It(`Invoke CreateClassifier successfully`, func() {
				naturalLanguageClassifierService, serviceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageClassifierService).ToNot(BeNil())
				naturalLanguageClassifierService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := naturalLanguageClassifierService.CreateClassifier(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateClassifierOptions model
				createClassifierOptionsModel := new(naturallanguageclassifierv1.CreateClassifierOptions)
				createClassifierOptionsModel.TrainingMetadata = CreateMockReader("This is a mock file.")
				createClassifierOptionsModel.TrainingData = CreateMockReader("This is a mock file.")
				createClassifierOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = naturalLanguageClassifierService.CreateClassifier(createClassifierOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = naturalLanguageClassifierService.CreateClassifierWithContext(ctx, createClassifierOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				naturalLanguageClassifierService.DisableRetries()
				result, response, operationErr = naturalLanguageClassifierService.CreateClassifier(createClassifierOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = naturalLanguageClassifierService.CreateClassifierWithContext(ctx, createClassifierOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CreateClassifier with error: Operation validation and request error`, func() {
				naturalLanguageClassifierService, serviceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageClassifierService).ToNot(BeNil())

				// Construct an instance of the CreateClassifierOptions model
				createClassifierOptionsModel := new(naturallanguageclassifierv1.CreateClassifierOptions)
				createClassifierOptionsModel.TrainingMetadata = CreateMockReader("This is a mock file.")
				createClassifierOptionsModel.TrainingData = CreateMockReader("This is a mock file.")
				createClassifierOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := naturalLanguageClassifierService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := naturalLanguageClassifierService.CreateClassifier(createClassifierOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateClassifierOptions model with no property values
				createClassifierOptionsModelNew := new(naturallanguageclassifierv1.CreateClassifierOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = naturalLanguageClassifierService.CreateClassifier(createClassifierOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListClassifiers(listClassifiersOptions *ListClassifiersOptions) - Operation response error`, func() {
		listClassifiersPath := "/v1/classifiers"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listClassifiersPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListClassifiers with error: Operation response processing error`, func() {
				naturalLanguageClassifierService, serviceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageClassifierService).ToNot(BeNil())

				// Construct an instance of the ListClassifiersOptions model
				listClassifiersOptionsModel := new(naturallanguageclassifierv1.ListClassifiersOptions)
				listClassifiersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := naturalLanguageClassifierService.ListClassifiers(listClassifiersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				naturalLanguageClassifierService.EnableRetries(0, 0)
				result, response, operationErr = naturalLanguageClassifierService.ListClassifiers(listClassifiersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListClassifiers(listClassifiersOptions *ListClassifiersOptions)`, func() {
		listClassifiersPath := "/v1/classifiers"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listClassifiersPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"classifiers": [{"name": "Name", "url": "URL", "status": "Non Existent", "classifier_id": "ClassifierID", "created": "2019-01-01T12:00:00", "status_description": "StatusDescription", "language": "Language"}]}`)
				}))
			})
			It(`Invoke ListClassifiers successfully`, func() {
				naturalLanguageClassifierService, serviceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageClassifierService).ToNot(BeNil())
				naturalLanguageClassifierService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := naturalLanguageClassifierService.ListClassifiers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListClassifiersOptions model
				listClassifiersOptionsModel := new(naturallanguageclassifierv1.ListClassifiersOptions)
				listClassifiersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = naturalLanguageClassifierService.ListClassifiers(listClassifiersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = naturalLanguageClassifierService.ListClassifiersWithContext(ctx, listClassifiersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				naturalLanguageClassifierService.DisableRetries()
				result, response, operationErr = naturalLanguageClassifierService.ListClassifiers(listClassifiersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = naturalLanguageClassifierService.ListClassifiersWithContext(ctx, listClassifiersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListClassifiers with error: Operation request error`, func() {
				naturalLanguageClassifierService, serviceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageClassifierService).ToNot(BeNil())

				// Construct an instance of the ListClassifiersOptions model
				listClassifiersOptionsModel := new(naturallanguageclassifierv1.ListClassifiersOptions)
				listClassifiersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := naturalLanguageClassifierService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := naturalLanguageClassifierService.ListClassifiers(listClassifiersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetClassifier(getClassifierOptions *GetClassifierOptions) - Operation response error`, func() {
		getClassifierPath := "/v1/classifiers/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getClassifierPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetClassifier with error: Operation response processing error`, func() {
				naturalLanguageClassifierService, serviceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageClassifierService).ToNot(BeNil())

				// Construct an instance of the GetClassifierOptions model
				getClassifierOptionsModel := new(naturallanguageclassifierv1.GetClassifierOptions)
				getClassifierOptionsModel.ClassifierID = core.StringPtr("testString")
				getClassifierOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := naturalLanguageClassifierService.GetClassifier(getClassifierOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				naturalLanguageClassifierService.EnableRetries(0, 0)
				result, response, operationErr = naturalLanguageClassifierService.GetClassifier(getClassifierOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetClassifier(getClassifierOptions *GetClassifierOptions)`, func() {
		getClassifierPath := "/v1/classifiers/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getClassifierPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "url": "URL", "status": "Non Existent", "classifier_id": "ClassifierID", "created": "2019-01-01T12:00:00", "status_description": "StatusDescription", "language": "Language"}`)
				}))
			})
			It(`Invoke GetClassifier successfully`, func() {
				naturalLanguageClassifierService, serviceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageClassifierService).ToNot(BeNil())
				naturalLanguageClassifierService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := naturalLanguageClassifierService.GetClassifier(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetClassifierOptions model
				getClassifierOptionsModel := new(naturallanguageclassifierv1.GetClassifierOptions)
				getClassifierOptionsModel.ClassifierID = core.StringPtr("testString")
				getClassifierOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = naturalLanguageClassifierService.GetClassifier(getClassifierOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = naturalLanguageClassifierService.GetClassifierWithContext(ctx, getClassifierOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				naturalLanguageClassifierService.DisableRetries()
				result, response, operationErr = naturalLanguageClassifierService.GetClassifier(getClassifierOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = naturalLanguageClassifierService.GetClassifierWithContext(ctx, getClassifierOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetClassifier with error: Operation validation and request error`, func() {
				naturalLanguageClassifierService, serviceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageClassifierService).ToNot(BeNil())

				// Construct an instance of the GetClassifierOptions model
				getClassifierOptionsModel := new(naturallanguageclassifierv1.GetClassifierOptions)
				getClassifierOptionsModel.ClassifierID = core.StringPtr("testString")
				getClassifierOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := naturalLanguageClassifierService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := naturalLanguageClassifierService.GetClassifier(getClassifierOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetClassifierOptions model with no property values
				getClassifierOptionsModelNew := new(naturallanguageclassifierv1.GetClassifierOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = naturalLanguageClassifierService.GetClassifier(getClassifierOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteClassifier(deleteClassifierOptions *DeleteClassifierOptions)`, func() {
		deleteClassifierPath := "/v1/classifiers/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteClassifierPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteClassifier successfully`, func() {
				naturalLanguageClassifierService, serviceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageClassifierService).ToNot(BeNil())
				naturalLanguageClassifierService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := naturalLanguageClassifierService.DeleteClassifier(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteClassifierOptions model
				deleteClassifierOptionsModel := new(naturallanguageclassifierv1.DeleteClassifierOptions)
				deleteClassifierOptionsModel.ClassifierID = core.StringPtr("testString")
				deleteClassifierOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = naturalLanguageClassifierService.DeleteClassifier(deleteClassifierOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				naturalLanguageClassifierService.DisableRetries()
				response, operationErr = naturalLanguageClassifierService.DeleteClassifier(deleteClassifierOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteClassifier with error: Operation validation and request error`, func() {
				naturalLanguageClassifierService, serviceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageClassifierService).ToNot(BeNil())

				// Construct an instance of the DeleteClassifierOptions model
				deleteClassifierOptionsModel := new(naturallanguageclassifierv1.DeleteClassifierOptions)
				deleteClassifierOptionsModel.ClassifierID = core.StringPtr("testString")
				deleteClassifierOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := naturalLanguageClassifierService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := naturalLanguageClassifierService.DeleteClassifier(deleteClassifierOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteClassifierOptions model with no property values
				deleteClassifierOptionsModelNew := new(naturallanguageclassifierv1.DeleteClassifierOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = naturalLanguageClassifierService.DeleteClassifier(deleteClassifierOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			naturalLanguageClassifierService, _ := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
				URL:           "http://naturallanguageclassifierv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewClassifyCollectionOptions successfully`, func() {
				// Construct an instance of the ClassifyInput model
				classifyInputModel := new(naturallanguageclassifierv1.ClassifyInput)
				Expect(classifyInputModel).ToNot(BeNil())
				classifyInputModel.Text = core.StringPtr("How hot will it be today?")
				Expect(classifyInputModel.Text).To(Equal(core.StringPtr("How hot will it be today?")))

				// Construct an instance of the ClassifyCollectionOptions model
				classifierID := "testString"
				classifyCollectionOptionsCollection := []naturallanguageclassifierv1.ClassifyInput{}
				classifyCollectionOptionsModel := naturalLanguageClassifierService.NewClassifyCollectionOptions(classifierID, classifyCollectionOptionsCollection)
				classifyCollectionOptionsModel.SetClassifierID("testString")
				classifyCollectionOptionsModel.SetCollection([]naturallanguageclassifierv1.ClassifyInput{*classifyInputModel})
				classifyCollectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(classifyCollectionOptionsModel).ToNot(BeNil())
				Expect(classifyCollectionOptionsModel.ClassifierID).To(Equal(core.StringPtr("testString")))
				Expect(classifyCollectionOptionsModel.Collection).To(Equal([]naturallanguageclassifierv1.ClassifyInput{*classifyInputModel}))
				Expect(classifyCollectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewClassifyInput successfully`, func() {
				text := "testString"
				model, err := naturalLanguageClassifierService.NewClassifyInput(text)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewClassifyOptions successfully`, func() {
				// Construct an instance of the ClassifyOptions model
				classifierID := "testString"
				classifyOptionsText := "testString"
				classifyOptionsModel := naturalLanguageClassifierService.NewClassifyOptions(classifierID, classifyOptionsText)
				classifyOptionsModel.SetClassifierID("testString")
				classifyOptionsModel.SetText("testString")
				classifyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(classifyOptionsModel).ToNot(BeNil())
				Expect(classifyOptionsModel.ClassifierID).To(Equal(core.StringPtr("testString")))
				Expect(classifyOptionsModel.Text).To(Equal(core.StringPtr("testString")))
				Expect(classifyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateClassifierOptions successfully`, func() {
				// Construct an instance of the CreateClassifierOptions model
				trainingMetadata := CreateMockReader("This is a mock file.")
				trainingData := CreateMockReader("This is a mock file.")
				createClassifierOptionsModel := naturalLanguageClassifierService.NewCreateClassifierOptions(trainingMetadata, trainingData)
				createClassifierOptionsModel.SetTrainingMetadata(CreateMockReader("This is a mock file."))
				createClassifierOptionsModel.SetTrainingData(CreateMockReader("This is a mock file."))
				createClassifierOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createClassifierOptionsModel).ToNot(BeNil())
				Expect(createClassifierOptionsModel.TrainingMetadata).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(createClassifierOptionsModel.TrainingData).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(createClassifierOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteClassifierOptions successfully`, func() {
				// Construct an instance of the DeleteClassifierOptions model
				classifierID := "testString"
				deleteClassifierOptionsModel := naturalLanguageClassifierService.NewDeleteClassifierOptions(classifierID)
				deleteClassifierOptionsModel.SetClassifierID("testString")
				deleteClassifierOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteClassifierOptionsModel).ToNot(BeNil())
				Expect(deleteClassifierOptionsModel.ClassifierID).To(Equal(core.StringPtr("testString")))
				Expect(deleteClassifierOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetClassifierOptions successfully`, func() {
				// Construct an instance of the GetClassifierOptions model
				classifierID := "testString"
				getClassifierOptionsModel := naturalLanguageClassifierService.NewGetClassifierOptions(classifierID)
				getClassifierOptionsModel.SetClassifierID("testString")
				getClassifierOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getClassifierOptionsModel).ToNot(BeNil())
				Expect(getClassifierOptionsModel.ClassifierID).To(Equal(core.StringPtr("testString")))
				Expect(getClassifierOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListClassifiersOptions successfully`, func() {
				// Construct an instance of the ListClassifiersOptions model
				listClassifiersOptionsModel := naturalLanguageClassifierService.NewListClassifiersOptions()
				listClassifiersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listClassifiersOptionsModel).ToNot(BeNil())
				Expect(listClassifiersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			mockDate := CreateMockDate()
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime()
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

func CreateMockDate() *strfmt.Date {
	d := strfmt.Date(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func CreateMockDateTime() *strfmt.DateTime {
	d := strfmt.DateTime(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
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
