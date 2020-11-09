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

package visualrecognitionv3_test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/visualrecognitionv3"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

var _ = Describe(`VisualRecognitionV3`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		version := "testString"
		It(`Instantiate service client`, func() {
			visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version: core.StringPtr(version),
			})
			Expect(visualRecognitionService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
			Expect(visualRecognitionService.Service.IsSSLDisabled()).To(BeFalse())
			visualRecognitionService.DisableSSLVerification()
			Expect(visualRecognitionService.Service.IsSSLDisabled()).To(BeTrue())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
				URL: "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(visualRecognitionService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
				URL: "https://visualrecognitionv3/api",
				Version: core.StringPtr(version),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(visualRecognitionService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{})
			Expect(visualRecognitionService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"WATSON_VISION_COMBINED_URL": "https://visualrecognitionv3/api",
				"WATSON_VISION_COMBINED_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					Version: core.StringPtr(version),
				})
				Expect(visualRecognitionService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL: "https://testService/api",
					Version: core.StringPtr(version),
				})
				Expect(visualRecognitionService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					Version: core.StringPtr(version),
				})
				err := visualRecognitionService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"WATSON_VISION_COMBINED_URL": "https://visualrecognitionv3/api",
				"WATSON_VISION_COMBINED_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(visualRecognitionService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"WATSON_VISION_COMBINED_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
				URL: "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(visualRecognitionService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Classify(classifyOptions *ClassifyOptions) - Operation response error`, func() {
		version := "testString"
		classifyPath := "/v3/classify"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(classifyPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "en")))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke Classify with error: Operation response processing error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the ClassifyOptions model
				classifyOptionsModel := new(visualrecognitionv3.ClassifyOptions)
				classifyOptionsModel.ImagesFile = CreateMockReader("This is a mock file.")
				classifyOptionsModel.ImagesFilename = core.StringPtr("testString")
				classifyOptionsModel.ImagesFileContentType = core.StringPtr("testString")
				classifyOptionsModel.URL = core.StringPtr("testString")
				classifyOptionsModel.Threshold = core.Float32Ptr(float32(36.0))
				classifyOptionsModel.Owners = []string{"testString"}
				classifyOptionsModel.ClassifierIds = []string{"testString"}
				classifyOptionsModel.AcceptLanguage = core.StringPtr("en")
				classifyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := visualRecognitionService.Classify(classifyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				visualRecognitionService.EnableRetries(0, 0)
				result, response, operationErr = visualRecognitionService.Classify(classifyOptionsModel)
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
		version := "testString"
		classifyPath := "/v3/classify"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(classifyPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "en")))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"custom_classes": 13, "images_processed": 15, "images": [{"source_url": "SourceURL", "resolved_url": "ResolvedURL", "image": "Image", "error": {"code": 4, "description": "Description", "error_id": "ErrorID"}, "classifiers": [{"name": "Name", "classifier_id": "ClassifierID", "classes": [{"class": "Class", "score": 0, "type_hierarchy": "TypeHierarchy"}]}]}], "warnings": [{"warning_id": "WarningID", "description": "Description"}]}`)
				}))
			})
			It(`Invoke Classify successfully`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())
				visualRecognitionService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := visualRecognitionService.Classify(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ClassifyOptions model
				classifyOptionsModel := new(visualrecognitionv3.ClassifyOptions)
				classifyOptionsModel.ImagesFile = CreateMockReader("This is a mock file.")
				classifyOptionsModel.ImagesFilename = core.StringPtr("testString")
				classifyOptionsModel.ImagesFileContentType = core.StringPtr("testString")
				classifyOptionsModel.URL = core.StringPtr("testString")
				classifyOptionsModel.Threshold = core.Float32Ptr(float32(36.0))
				classifyOptionsModel.Owners = []string{"testString"}
				classifyOptionsModel.ClassifierIds = []string{"testString"}
				classifyOptionsModel.AcceptLanguage = core.StringPtr("en")
				classifyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = visualRecognitionService.Classify(classifyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.ClassifyWithContext(ctx, classifyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				visualRecognitionService.DisableRetries()
				result, response, operationErr = visualRecognitionService.Classify(classifyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.ClassifyWithContext(ctx, classifyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke Classify with error: Param validation error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:  testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the ClassifyOptions model
				classifyOptionsModel := new(visualrecognitionv3.ClassifyOptions)
				// Invoke operation with invalid options model (negative test)
				result, response, operationErr := visualRecognitionService.Classify(classifyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			It(`Invoke Classify with error: Operation request error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the ClassifyOptions model
				classifyOptionsModel := new(visualrecognitionv3.ClassifyOptions)
				classifyOptionsModel.ImagesFile = CreateMockReader("This is a mock file.")
				classifyOptionsModel.ImagesFilename = core.StringPtr("testString")
				classifyOptionsModel.ImagesFileContentType = core.StringPtr("testString")
				classifyOptionsModel.URL = core.StringPtr("testString")
				classifyOptionsModel.Threshold = core.Float32Ptr(float32(36.0))
				classifyOptionsModel.Owners = []string{"testString"}
				classifyOptionsModel.ClassifierIds = []string{"testString"}
				classifyOptionsModel.AcceptLanguage = core.StringPtr("en")
				classifyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := visualRecognitionService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := visualRecognitionService.Classify(classifyOptionsModel)
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
	Describe(`Service constructor tests`, func() {
		version := "testString"
		It(`Instantiate service client`, func() {
			visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version: core.StringPtr(version),
			})
			Expect(visualRecognitionService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
			Expect(visualRecognitionService.Service.IsSSLDisabled()).To(BeFalse())
			visualRecognitionService.DisableSSLVerification()
			Expect(visualRecognitionService.Service.IsSSLDisabled()).To(BeTrue())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
				URL: "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(visualRecognitionService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
				URL: "https://visualrecognitionv3/api",
				Version: core.StringPtr(version),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(visualRecognitionService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{})
			Expect(visualRecognitionService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"WATSON_VISION_COMBINED_URL": "https://visualrecognitionv3/api",
				"WATSON_VISION_COMBINED_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					Version: core.StringPtr(version),
				})
				Expect(visualRecognitionService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL: "https://testService/api",
					Version: core.StringPtr(version),
				})
				Expect(visualRecognitionService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					Version: core.StringPtr(version),
				})
				err := visualRecognitionService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"WATSON_VISION_COMBINED_URL": "https://visualrecognitionv3/api",
				"WATSON_VISION_COMBINED_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(visualRecognitionService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"WATSON_VISION_COMBINED_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
				URL: "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(visualRecognitionService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`CreateClassifier(createClassifierOptions *CreateClassifierOptions) - Operation response error`, func() {
		version := "testString"
		createClassifierPath := "/v3/classifiers"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createClassifierPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateClassifier with error: Operation response processing error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the CreateClassifierOptions model
				createClassifierOptionsModel := new(visualrecognitionv3.CreateClassifierOptions)
				createClassifierOptionsModel.Name = core.StringPtr("testString")
				createClassifierOptionsModel.PositiveExamples = map[string]io.ReadCloser { "key1": CreateMockReader("This is a mock file.") }
				createClassifierOptionsModel.NegativeExamples = CreateMockReader("This is a mock file.")
				createClassifierOptionsModel.NegativeExamplesFilename = core.StringPtr("testString")
				createClassifierOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := visualRecognitionService.CreateClassifier(createClassifierOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				visualRecognitionService.EnableRetries(0, 0)
				result, response, operationErr = visualRecognitionService.CreateClassifier(createClassifierOptionsModel)
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
		version := "testString"
		createClassifierPath := "/v3/classifiers"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createClassifierPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"classifier_id": "ClassifierID", "name": "Name", "owner": "Owner", "status": "ready", "core_ml_enabled": false, "explanation": "Explanation", "created": "2019-01-01T12:00:00", "classes": [{"class": "Class"}], "retrained": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke CreateClassifier successfully`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())
				visualRecognitionService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := visualRecognitionService.CreateClassifier(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateClassifierOptions model
				createClassifierOptionsModel := new(visualrecognitionv3.CreateClassifierOptions)
				createClassifierOptionsModel.Name = core.StringPtr("testString")
				createClassifierOptionsModel.PositiveExamples = map[string]io.ReadCloser { "key1": CreateMockReader("This is a mock file.") }
				createClassifierOptionsModel.NegativeExamples = CreateMockReader("This is a mock file.")
				createClassifierOptionsModel.NegativeExamplesFilename = core.StringPtr("testString")
				createClassifierOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = visualRecognitionService.CreateClassifier(createClassifierOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.CreateClassifierWithContext(ctx, createClassifierOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				visualRecognitionService.DisableRetries()
				result, response, operationErr = visualRecognitionService.CreateClassifier(createClassifierOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.CreateClassifierWithContext(ctx, createClassifierOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CreateClassifier with error: Operation validation and request error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the CreateClassifierOptions model
				createClassifierOptionsModel := new(visualrecognitionv3.CreateClassifierOptions)
				createClassifierOptionsModel.Name = core.StringPtr("testString")
				createClassifierOptionsModel.PositiveExamples = map[string]io.ReadCloser { "key1": CreateMockReader("This is a mock file.") }
				createClassifierOptionsModel.NegativeExamples = CreateMockReader("This is a mock file.")
				createClassifierOptionsModel.NegativeExamplesFilename = core.StringPtr("testString")
				createClassifierOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := visualRecognitionService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := visualRecognitionService.CreateClassifier(createClassifierOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateClassifierOptions model with no property values
				createClassifierOptionsModelNew := new(visualrecognitionv3.CreateClassifierOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = visualRecognitionService.CreateClassifier(createClassifierOptionsModelNew)
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
		version := "testString"
		listClassifiersPath := "/v3/classifiers"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listClassifiersPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))


					// TODO: Add check for verbose query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListClassifiers with error: Operation response processing error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the ListClassifiersOptions model
				listClassifiersOptionsModel := new(visualrecognitionv3.ListClassifiersOptions)
				listClassifiersOptionsModel.Verbose = core.BoolPtr(true)
				listClassifiersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := visualRecognitionService.ListClassifiers(listClassifiersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				visualRecognitionService.EnableRetries(0, 0)
				result, response, operationErr = visualRecognitionService.ListClassifiers(listClassifiersOptionsModel)
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
		version := "testString"
		listClassifiersPath := "/v3/classifiers"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listClassifiersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))


					// TODO: Add check for verbose query parameter

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"classifiers": [{"classifier_id": "ClassifierID", "name": "Name", "owner": "Owner", "status": "ready", "core_ml_enabled": false, "explanation": "Explanation", "created": "2019-01-01T12:00:00", "classes": [{"class": "Class"}], "retrained": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00"}]}`)
				}))
			})
			It(`Invoke ListClassifiers successfully`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())
				visualRecognitionService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := visualRecognitionService.ListClassifiers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListClassifiersOptions model
				listClassifiersOptionsModel := new(visualrecognitionv3.ListClassifiersOptions)
				listClassifiersOptionsModel.Verbose = core.BoolPtr(true)
				listClassifiersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = visualRecognitionService.ListClassifiers(listClassifiersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.ListClassifiersWithContext(ctx, listClassifiersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				visualRecognitionService.DisableRetries()
				result, response, operationErr = visualRecognitionService.ListClassifiers(listClassifiersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.ListClassifiersWithContext(ctx, listClassifiersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListClassifiers with error: Operation request error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the ListClassifiersOptions model
				listClassifiersOptionsModel := new(visualrecognitionv3.ListClassifiersOptions)
				listClassifiersOptionsModel.Verbose = core.BoolPtr(true)
				listClassifiersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := visualRecognitionService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := visualRecognitionService.ListClassifiers(listClassifiersOptionsModel)
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
		version := "testString"
		getClassifierPath := "/v3/classifiers/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getClassifierPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetClassifier with error: Operation response processing error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the GetClassifierOptions model
				getClassifierOptionsModel := new(visualrecognitionv3.GetClassifierOptions)
				getClassifierOptionsModel.ClassifierID = core.StringPtr("testString")
				getClassifierOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := visualRecognitionService.GetClassifier(getClassifierOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				visualRecognitionService.EnableRetries(0, 0)
				result, response, operationErr = visualRecognitionService.GetClassifier(getClassifierOptionsModel)
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
		version := "testString"
		getClassifierPath := "/v3/classifiers/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getClassifierPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"classifier_id": "ClassifierID", "name": "Name", "owner": "Owner", "status": "ready", "core_ml_enabled": false, "explanation": "Explanation", "created": "2019-01-01T12:00:00", "classes": [{"class": "Class"}], "retrained": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke GetClassifier successfully`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())
				visualRecognitionService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := visualRecognitionService.GetClassifier(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetClassifierOptions model
				getClassifierOptionsModel := new(visualrecognitionv3.GetClassifierOptions)
				getClassifierOptionsModel.ClassifierID = core.StringPtr("testString")
				getClassifierOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = visualRecognitionService.GetClassifier(getClassifierOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.GetClassifierWithContext(ctx, getClassifierOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				visualRecognitionService.DisableRetries()
				result, response, operationErr = visualRecognitionService.GetClassifier(getClassifierOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.GetClassifierWithContext(ctx, getClassifierOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetClassifier with error: Operation validation and request error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the GetClassifierOptions model
				getClassifierOptionsModel := new(visualrecognitionv3.GetClassifierOptions)
				getClassifierOptionsModel.ClassifierID = core.StringPtr("testString")
				getClassifierOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := visualRecognitionService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := visualRecognitionService.GetClassifier(getClassifierOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetClassifierOptions model with no property values
				getClassifierOptionsModelNew := new(visualrecognitionv3.GetClassifierOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = visualRecognitionService.GetClassifier(getClassifierOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateClassifier(updateClassifierOptions *UpdateClassifierOptions) - Operation response error`, func() {
		version := "testString"
		updateClassifierPath := "/v3/classifiers/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateClassifierPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateClassifier with error: Operation response processing error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the UpdateClassifierOptions model
				updateClassifierOptionsModel := new(visualrecognitionv3.UpdateClassifierOptions)
				updateClassifierOptionsModel.ClassifierID = core.StringPtr("testString")
				updateClassifierOptionsModel.PositiveExamples = map[string]io.ReadCloser { "key1": CreateMockReader("This is a mock file.") }
				updateClassifierOptionsModel.NegativeExamples = CreateMockReader("This is a mock file.")
				updateClassifierOptionsModel.NegativeExamplesFilename = core.StringPtr("testString")
				updateClassifierOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := visualRecognitionService.UpdateClassifier(updateClassifierOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				visualRecognitionService.EnableRetries(0, 0)
				result, response, operationErr = visualRecognitionService.UpdateClassifier(updateClassifierOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateClassifier(updateClassifierOptions *UpdateClassifierOptions)`, func() {
		version := "testString"
		updateClassifierPath := "/v3/classifiers/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateClassifierPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"classifier_id": "ClassifierID", "name": "Name", "owner": "Owner", "status": "ready", "core_ml_enabled": false, "explanation": "Explanation", "created": "2019-01-01T12:00:00", "classes": [{"class": "Class"}], "retrained": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke UpdateClassifier successfully`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())
				visualRecognitionService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := visualRecognitionService.UpdateClassifier(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateClassifierOptions model
				updateClassifierOptionsModel := new(visualrecognitionv3.UpdateClassifierOptions)
				updateClassifierOptionsModel.ClassifierID = core.StringPtr("testString")
				updateClassifierOptionsModel.PositiveExamples = map[string]io.ReadCloser { "key1": CreateMockReader("This is a mock file.") }
				updateClassifierOptionsModel.NegativeExamples = CreateMockReader("This is a mock file.")
				updateClassifierOptionsModel.NegativeExamplesFilename = core.StringPtr("testString")
				updateClassifierOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = visualRecognitionService.UpdateClassifier(updateClassifierOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.UpdateClassifierWithContext(ctx, updateClassifierOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				visualRecognitionService.DisableRetries()
				result, response, operationErr = visualRecognitionService.UpdateClassifier(updateClassifierOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.UpdateClassifierWithContext(ctx, updateClassifierOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateClassifier with error: Param validation error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:  testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the UpdateClassifierOptions model
				updateClassifierOptionsModel := new(visualrecognitionv3.UpdateClassifierOptions)
				// Invoke operation with invalid options model (negative test)
				result, response, operationErr := visualRecognitionService.UpdateClassifier(updateClassifierOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			It(`Invoke UpdateClassifier with error: Operation validation and request error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the UpdateClassifierOptions model
				updateClassifierOptionsModel := new(visualrecognitionv3.UpdateClassifierOptions)
				updateClassifierOptionsModel.ClassifierID = core.StringPtr("testString")
				updateClassifierOptionsModel.PositiveExamples = map[string]io.ReadCloser { "key1": CreateMockReader("This is a mock file.") }
				updateClassifierOptionsModel.NegativeExamples = CreateMockReader("This is a mock file.")
				updateClassifierOptionsModel.NegativeExamplesFilename = core.StringPtr("testString")
				updateClassifierOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := visualRecognitionService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := visualRecognitionService.UpdateClassifier(updateClassifierOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateClassifierOptions model with no property values
				updateClassifierOptionsModelNew := new(visualrecognitionv3.UpdateClassifierOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = visualRecognitionService.UpdateClassifier(updateClassifierOptionsModelNew)
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
		version := "testString"
		deleteClassifierPath := "/v3/classifiers/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteClassifierPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteClassifier successfully`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())
				visualRecognitionService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := visualRecognitionService.DeleteClassifier(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteClassifierOptions model
				deleteClassifierOptionsModel := new(visualrecognitionv3.DeleteClassifierOptions)
				deleteClassifierOptionsModel.ClassifierID = core.StringPtr("testString")
				deleteClassifierOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = visualRecognitionService.DeleteClassifier(deleteClassifierOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				visualRecognitionService.DisableRetries()
				response, operationErr = visualRecognitionService.DeleteClassifier(deleteClassifierOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteClassifier with error: Operation validation and request error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the DeleteClassifierOptions model
				deleteClassifierOptionsModel := new(visualrecognitionv3.DeleteClassifierOptions)
				deleteClassifierOptionsModel.ClassifierID = core.StringPtr("testString")
				deleteClassifierOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := visualRecognitionService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := visualRecognitionService.DeleteClassifier(deleteClassifierOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteClassifierOptions model with no property values
				deleteClassifierOptionsModelNew := new(visualrecognitionv3.DeleteClassifierOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = visualRecognitionService.DeleteClassifier(deleteClassifierOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		version := "testString"
		It(`Instantiate service client`, func() {
			visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version: core.StringPtr(version),
			})
			Expect(visualRecognitionService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
			Expect(visualRecognitionService.Service.IsSSLDisabled()).To(BeFalse())
			visualRecognitionService.DisableSSLVerification()
			Expect(visualRecognitionService.Service.IsSSLDisabled()).To(BeTrue())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
				URL: "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(visualRecognitionService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
				URL: "https://visualrecognitionv3/api",
				Version: core.StringPtr(version),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(visualRecognitionService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{})
			Expect(visualRecognitionService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"WATSON_VISION_COMBINED_URL": "https://visualrecognitionv3/api",
				"WATSON_VISION_COMBINED_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					Version: core.StringPtr(version),
				})
				Expect(visualRecognitionService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL: "https://testService/api",
					Version: core.StringPtr(version),
				})
				Expect(visualRecognitionService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					Version: core.StringPtr(version),
				})
				err := visualRecognitionService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"WATSON_VISION_COMBINED_URL": "https://visualrecognitionv3/api",
				"WATSON_VISION_COMBINED_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(visualRecognitionService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"WATSON_VISION_COMBINED_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
				URL: "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(visualRecognitionService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})

	Describe(`GetCoreMlModel(getCoreMlModelOptions *GetCoreMlModelOptions)`, func() {
		version := "testString"
		getCoreMlModelPath := "/v3/classifiers/testString/core_ml_model"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCoreMlModelPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/octet-stream")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke GetCoreMlModel successfully`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())
				visualRecognitionService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := visualRecognitionService.GetCoreMlModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCoreMlModelOptions model
				getCoreMlModelOptionsModel := new(visualrecognitionv3.GetCoreMlModelOptions)
				getCoreMlModelOptionsModel.ClassifierID = core.StringPtr("testString")
				getCoreMlModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = visualRecognitionService.GetCoreMlModel(getCoreMlModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.GetCoreMlModelWithContext(ctx, getCoreMlModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				visualRecognitionService.DisableRetries()
				result, response, operationErr = visualRecognitionService.GetCoreMlModel(getCoreMlModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.GetCoreMlModelWithContext(ctx, getCoreMlModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetCoreMlModel with error: Operation validation and request error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the GetCoreMlModelOptions model
				getCoreMlModelOptionsModel := new(visualrecognitionv3.GetCoreMlModelOptions)
				getCoreMlModelOptionsModel.ClassifierID = core.StringPtr("testString")
				getCoreMlModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := visualRecognitionService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := visualRecognitionService.GetCoreMlModel(getCoreMlModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetCoreMlModelOptions model with no property values
				getCoreMlModelOptionsModelNew := new(visualrecognitionv3.GetCoreMlModelOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = visualRecognitionService.GetCoreMlModel(getCoreMlModelOptionsModelNew)
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
		version := "testString"
		It(`Instantiate service client`, func() {
			visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version: core.StringPtr(version),
			})
			Expect(visualRecognitionService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
			Expect(visualRecognitionService.Service.IsSSLDisabled()).To(BeFalse())
			visualRecognitionService.DisableSSLVerification()
			Expect(visualRecognitionService.Service.IsSSLDisabled()).To(BeTrue())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
				URL: "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(visualRecognitionService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
				URL: "https://visualrecognitionv3/api",
				Version: core.StringPtr(version),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(visualRecognitionService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{})
			Expect(visualRecognitionService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"WATSON_VISION_COMBINED_URL": "https://visualrecognitionv3/api",
				"WATSON_VISION_COMBINED_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					Version: core.StringPtr(version),
				})
				Expect(visualRecognitionService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL: "https://testService/api",
					Version: core.StringPtr(version),
				})
				Expect(visualRecognitionService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					Version: core.StringPtr(version),
				})
				err := visualRecognitionService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"WATSON_VISION_COMBINED_URL": "https://visualrecognitionv3/api",
				"WATSON_VISION_COMBINED_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(visualRecognitionService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"WATSON_VISION_COMBINED_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
				URL: "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(visualRecognitionService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})

	Describe(`DeleteUserData(deleteUserDataOptions *DeleteUserDataOptions)`, func() {
		version := "testString"
		deleteUserDataPath := "/v3/user_data"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteUserDataPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["customer_id"]).To(Equal([]string{"testString"}))

					res.WriteHeader(202)
				}))
			})
			It(`Invoke DeleteUserData successfully`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())
				visualRecognitionService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := visualRecognitionService.DeleteUserData(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteUserDataOptions model
				deleteUserDataOptionsModel := new(visualrecognitionv3.DeleteUserDataOptions)
				deleteUserDataOptionsModel.CustomerID = core.StringPtr("testString")
				deleteUserDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = visualRecognitionService.DeleteUserData(deleteUserDataOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				visualRecognitionService.DisableRetries()
				response, operationErr = visualRecognitionService.DeleteUserData(deleteUserDataOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteUserData with error: Operation validation and request error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the DeleteUserDataOptions model
				deleteUserDataOptionsModel := new(visualrecognitionv3.DeleteUserDataOptions)
				deleteUserDataOptionsModel.CustomerID = core.StringPtr("testString")
				deleteUserDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := visualRecognitionService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := visualRecognitionService.DeleteUserData(deleteUserDataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteUserDataOptions model with no property values
				deleteUserDataOptionsModelNew := new(visualrecognitionv3.DeleteUserDataOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = visualRecognitionService.DeleteUserData(deleteUserDataOptionsModelNew)
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
			version := "testString"
			visualRecognitionService, _ := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
				URL:           "http://visualrecognitionv3modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Version: core.StringPtr(version),
			})
			It(`Invoke NewClassifyOptions successfully`, func() {
				// Construct an instance of the ClassifyOptions model
				classifyOptionsModel := visualRecognitionService.NewClassifyOptions()
				classifyOptionsModel.SetImagesFile(CreateMockReader("This is a mock file."))
				classifyOptionsModel.SetImagesFilename("testString")
				classifyOptionsModel.SetImagesFileContentType("testString")
				classifyOptionsModel.SetURL("testString")
				classifyOptionsModel.SetThreshold(float32(36.0))
				classifyOptionsModel.SetOwners([]string{"testString"})
				classifyOptionsModel.SetClassifierIds([]string{"testString"})
				classifyOptionsModel.SetAcceptLanguage("en")
				classifyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(classifyOptionsModel).ToNot(BeNil())
				Expect(classifyOptionsModel.ImagesFile).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(classifyOptionsModel.ImagesFilename).To(Equal(core.StringPtr("testString")))
				Expect(classifyOptionsModel.ImagesFileContentType).To(Equal(core.StringPtr("testString")))
				Expect(classifyOptionsModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(classifyOptionsModel.Threshold).To(Equal(core.Float32Ptr(float32(36.0))))
				Expect(classifyOptionsModel.Owners).To(Equal([]string{"testString"}))
				Expect(classifyOptionsModel.ClassifierIds).To(Equal([]string{"testString"}))
				Expect(classifyOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("en")))
				Expect(classifyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateClassifierOptions successfully`, func() {
				// Construct an instance of the CreateClassifierOptions model
				name := "testString"
				createClassifierOptionsModel := visualRecognitionService.NewCreateClassifierOptions(name)
				createClassifierOptionsModel.SetName("testString")
				createClassifierOptionsModel.AddPositiveExamples("foo", CreateMockReader("This is a mock file."))
				createClassifierOptionsModel.SetNegativeExamples(CreateMockReader("This is a mock file."))
				createClassifierOptionsModel.SetNegativeExamplesFilename("testString")
				createClassifierOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createClassifierOptionsModel).ToNot(BeNil())
				Expect(createClassifierOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createClassifierOptionsModel.PositiveExamples["foo"]).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(createClassifierOptionsModel.NegativeExamples).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(createClassifierOptionsModel.NegativeExamplesFilename).To(Equal(core.StringPtr("testString")))
				Expect(createClassifierOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteClassifierOptions successfully`, func() {
				// Construct an instance of the DeleteClassifierOptions model
				classifierID := "testString"
				deleteClassifierOptionsModel := visualRecognitionService.NewDeleteClassifierOptions(classifierID)
				deleteClassifierOptionsModel.SetClassifierID("testString")
				deleteClassifierOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteClassifierOptionsModel).ToNot(BeNil())
				Expect(deleteClassifierOptionsModel.ClassifierID).To(Equal(core.StringPtr("testString")))
				Expect(deleteClassifierOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteUserDataOptions successfully`, func() {
				// Construct an instance of the DeleteUserDataOptions model
				customerID := "testString"
				deleteUserDataOptionsModel := visualRecognitionService.NewDeleteUserDataOptions(customerID)
				deleteUserDataOptionsModel.SetCustomerID("testString")
				deleteUserDataOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteUserDataOptionsModel).ToNot(BeNil())
				Expect(deleteUserDataOptionsModel.CustomerID).To(Equal(core.StringPtr("testString")))
				Expect(deleteUserDataOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetClassifierOptions successfully`, func() {
				// Construct an instance of the GetClassifierOptions model
				classifierID := "testString"
				getClassifierOptionsModel := visualRecognitionService.NewGetClassifierOptions(classifierID)
				getClassifierOptionsModel.SetClassifierID("testString")
				getClassifierOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getClassifierOptionsModel).ToNot(BeNil())
				Expect(getClassifierOptionsModel.ClassifierID).To(Equal(core.StringPtr("testString")))
				Expect(getClassifierOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCoreMlModelOptions successfully`, func() {
				// Construct an instance of the GetCoreMlModelOptions model
				classifierID := "testString"
				getCoreMlModelOptionsModel := visualRecognitionService.NewGetCoreMlModelOptions(classifierID)
				getCoreMlModelOptionsModel.SetClassifierID("testString")
				getCoreMlModelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCoreMlModelOptionsModel).ToNot(BeNil())
				Expect(getCoreMlModelOptionsModel.ClassifierID).To(Equal(core.StringPtr("testString")))
				Expect(getCoreMlModelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListClassifiersOptions successfully`, func() {
				// Construct an instance of the ListClassifiersOptions model
				listClassifiersOptionsModel := visualRecognitionService.NewListClassifiersOptions()
				listClassifiersOptionsModel.SetVerbose(true)
				listClassifiersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listClassifiersOptionsModel).ToNot(BeNil())
				Expect(listClassifiersOptionsModel.Verbose).To(Equal(core.BoolPtr(true)))
				Expect(listClassifiersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateClassifierOptions successfully`, func() {
				// Construct an instance of the UpdateClassifierOptions model
				classifierID := "testString"
				updateClassifierOptionsModel := visualRecognitionService.NewUpdateClassifierOptions(classifierID)
				updateClassifierOptionsModel.SetClassifierID("testString")
				updateClassifierOptionsModel.AddPositiveExamples("foo", CreateMockReader("This is a mock file."))
				updateClassifierOptionsModel.SetNegativeExamples(CreateMockReader("This is a mock file."))
				updateClassifierOptionsModel.SetNegativeExamplesFilename("testString")
				updateClassifierOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateClassifierOptionsModel).ToNot(BeNil())
				Expect(updateClassifierOptionsModel.ClassifierID).To(Equal(core.StringPtr("testString")))
				Expect(updateClassifierOptionsModel.PositiveExamples["foo"]).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(updateClassifierOptionsModel.NegativeExamples).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(updateClassifierOptionsModel.NegativeExamplesFilename).To(Equal(core.StringPtr("testString")))
				Expect(updateClassifierOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
