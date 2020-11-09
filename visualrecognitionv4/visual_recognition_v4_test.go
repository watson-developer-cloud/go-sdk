/**
 * (C) Copyright IBM Corp. 2020.
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

package visualrecognitionv4_test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/visualrecognitionv4"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

var _ = Describe(`VisualRecognitionV4`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		version := "testString"
		It(`Instantiate service client`, func() {
			visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
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
			visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
				URL: "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(visualRecognitionService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
				URL: "https://visualrecognitionv4/api",
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
			visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{})
			Expect(visualRecognitionService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"WATSON_VISION_COMBINED_URL": "https://visualrecognitionv4/api",
				"WATSON_VISION_COMBINED_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					Version: core.StringPtr(version),
				})
				Expect(visualRecognitionService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := visualRecognitionService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != visualRecognitionService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(visualRecognitionService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(visualRecognitionService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL: "https://testService/api",
					Version: core.StringPtr(version),
				})
				Expect(visualRecognitionService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := visualRecognitionService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != visualRecognitionService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(visualRecognitionService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(visualRecognitionService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					Version: core.StringPtr(version),
				})
				err := visualRecognitionService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := visualRecognitionService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != visualRecognitionService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(visualRecognitionService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(visualRecognitionService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"WATSON_VISION_COMBINED_URL": "https://visualrecognitionv4/api",
				"WATSON_VISION_COMBINED_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
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
			visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
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
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = visualrecognitionv4.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`Analyze(analyzeOptions *AnalyzeOptions) - Operation response error`, func() {
		version := "testString"
		analyzePath := "/v4/analyze"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(analyzePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke Analyze with error: Operation response processing error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the AnalyzeOptions model
				analyzeOptionsModel := new(visualrecognitionv4.AnalyzeOptions)
				analyzeOptionsModel.CollectionIds = []string{"testString"}
				analyzeOptionsModel.Features = []string{"objects"}
				analyzeOptionsModel.ImagesFile = []visualrecognitionv4.FileWithMetadata{visualrecognitionv4.FileWithMetadata{Data: CreateMockReader("This is a mock file."), Filename: core.StringPtr("mockfilename.txt"), }}
				analyzeOptionsModel.ImageURL = []string{"testString"}
				analyzeOptionsModel.Threshold = core.Float32Ptr(float32(0.15))
				analyzeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := visualRecognitionService.Analyze(analyzeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				visualRecognitionService.EnableRetries(0, 0)
				result, response, operationErr = visualRecognitionService.Analyze(analyzeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`Analyze(analyzeOptions *AnalyzeOptions)`, func() {
		version := "testString"
		analyzePath := "/v4/analyze"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(analyzePath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"images": [{"source": {"type": "file", "filename": "Filename", "archive_filename": "ArchiveFilename", "source_url": "SourceURL", "resolved_url": "ResolvedURL"}, "dimensions": {"height": 6, "width": 5}, "objects": {"collections": [{"collection_id": "CollectionID", "objects": [{"object": "Object", "location": {"top": 3, "left": 4, "width": 5, "height": 6}, "score": 5}]}]}, "errors": [{"code": "invalid_field", "message": "Message", "more_info": "MoreInfo", "target": {"type": "field", "name": "Name"}}]}], "warnings": [{"code": "invalid_field", "message": "Message", "more_info": "MoreInfo"}], "trace": "Trace"}`)
				}))
			})
			It(`Invoke Analyze successfully`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())
				visualRecognitionService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := visualRecognitionService.Analyze(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AnalyzeOptions model
				analyzeOptionsModel := new(visualrecognitionv4.AnalyzeOptions)
				analyzeOptionsModel.CollectionIds = []string{"testString"}
				analyzeOptionsModel.Features = []string{"objects"}
				analyzeOptionsModel.ImagesFile = []visualrecognitionv4.FileWithMetadata{visualrecognitionv4.FileWithMetadata{Data: CreateMockReader("This is a mock file."), Filename: core.StringPtr("mockfilename.txt"), }}
				analyzeOptionsModel.ImageURL = []string{"testString"}
				analyzeOptionsModel.Threshold = core.Float32Ptr(float32(0.15))
				analyzeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = visualRecognitionService.Analyze(analyzeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.AnalyzeWithContext(ctx, analyzeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				visualRecognitionService.DisableRetries()
				result, response, operationErr = visualRecognitionService.Analyze(analyzeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.AnalyzeWithContext(ctx, analyzeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke Analyze with error: Operation validation and request error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the AnalyzeOptions model
				analyzeOptionsModel := new(visualrecognitionv4.AnalyzeOptions)
				analyzeOptionsModel.CollectionIds = []string{"testString"}
				analyzeOptionsModel.Features = []string{"objects"}
				analyzeOptionsModel.ImagesFile = []visualrecognitionv4.FileWithMetadata{visualrecognitionv4.FileWithMetadata{Data: CreateMockReader("This is a mock file."), Filename: core.StringPtr("mockfilename.txt"), }}
				analyzeOptionsModel.ImageURL = []string{"testString"}
				analyzeOptionsModel.Threshold = core.Float32Ptr(float32(0.15))
				analyzeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := visualRecognitionService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := visualRecognitionService.Analyze(analyzeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the AnalyzeOptions model with no property values
				analyzeOptionsModelNew := new(visualrecognitionv4.AnalyzeOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = visualRecognitionService.Analyze(analyzeOptionsModelNew)
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
			visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
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
			visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
				URL: "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(visualRecognitionService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
				URL: "https://visualrecognitionv4/api",
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
			visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{})
			Expect(visualRecognitionService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"WATSON_VISION_COMBINED_URL": "https://visualrecognitionv4/api",
				"WATSON_VISION_COMBINED_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					Version: core.StringPtr(version),
				})
				Expect(visualRecognitionService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := visualRecognitionService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != visualRecognitionService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(visualRecognitionService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(visualRecognitionService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL: "https://testService/api",
					Version: core.StringPtr(version),
				})
				Expect(visualRecognitionService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := visualRecognitionService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != visualRecognitionService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(visualRecognitionService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(visualRecognitionService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					Version: core.StringPtr(version),
				})
				err := visualRecognitionService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := visualRecognitionService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != visualRecognitionService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(visualRecognitionService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(visualRecognitionService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"WATSON_VISION_COMBINED_URL": "https://visualrecognitionv4/api",
				"WATSON_VISION_COMBINED_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
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
			visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
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
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = visualrecognitionv4.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`CreateCollection(createCollectionOptions *CreateCollectionOptions) - Operation response error`, func() {
		version := "testString"
		createCollectionPath := "/v4/collections"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCollectionPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateCollection with error: Operation response processing error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the ObjectTrainingStatus model
				objectTrainingStatusModel := new(visualrecognitionv4.ObjectTrainingStatus)
				objectTrainingStatusModel.Ready = core.BoolPtr(true)
				objectTrainingStatusModel.InProgress = core.BoolPtr(true)
				objectTrainingStatusModel.DataChanged = core.BoolPtr(true)
				objectTrainingStatusModel.LatestFailed = core.BoolPtr(true)
				objectTrainingStatusModel.RscnnReady = core.BoolPtr(true)
				objectTrainingStatusModel.Description = core.StringPtr("testString")

				// Construct an instance of the TrainingStatus model
				trainingStatusModel := new(visualrecognitionv4.TrainingStatus)
				trainingStatusModel.Objects = objectTrainingStatusModel

				// Construct an instance of the CreateCollectionOptions model
				createCollectionOptionsModel := new(visualrecognitionv4.CreateCollectionOptions)
				createCollectionOptionsModel.Name = core.StringPtr("testString")
				createCollectionOptionsModel.Description = core.StringPtr("testString")
				createCollectionOptionsModel.TrainingStatus = trainingStatusModel
				createCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := visualRecognitionService.CreateCollection(createCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				visualRecognitionService.EnableRetries(0, 0)
				result, response, operationErr = visualRecognitionService.CreateCollection(createCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateCollection(createCollectionOptions *CreateCollectionOptions)`, func() {
		version := "testString"
		createCollectionPath := "/v4/collections"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCollectionPath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"collection_id": "CollectionID", "name": "Name", "description": "Description", "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "image_count": 10, "training_status": {"objects": {"ready": false, "in_progress": true, "data_changed": false, "latest_failed": true, "rscnn_ready": true, "description": "Description"}}}`)
				}))
			})
			It(`Invoke CreateCollection successfully`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())
				visualRecognitionService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := visualRecognitionService.CreateCollection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ObjectTrainingStatus model
				objectTrainingStatusModel := new(visualrecognitionv4.ObjectTrainingStatus)
				objectTrainingStatusModel.Ready = core.BoolPtr(true)
				objectTrainingStatusModel.InProgress = core.BoolPtr(true)
				objectTrainingStatusModel.DataChanged = core.BoolPtr(true)
				objectTrainingStatusModel.LatestFailed = core.BoolPtr(true)
				objectTrainingStatusModel.RscnnReady = core.BoolPtr(true)
				objectTrainingStatusModel.Description = core.StringPtr("testString")

				// Construct an instance of the TrainingStatus model
				trainingStatusModel := new(visualrecognitionv4.TrainingStatus)
				trainingStatusModel.Objects = objectTrainingStatusModel

				// Construct an instance of the CreateCollectionOptions model
				createCollectionOptionsModel := new(visualrecognitionv4.CreateCollectionOptions)
				createCollectionOptionsModel.Name = core.StringPtr("testString")
				createCollectionOptionsModel.Description = core.StringPtr("testString")
				createCollectionOptionsModel.TrainingStatus = trainingStatusModel
				createCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = visualRecognitionService.CreateCollection(createCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.CreateCollectionWithContext(ctx, createCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				visualRecognitionService.DisableRetries()
				result, response, operationErr = visualRecognitionService.CreateCollection(createCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.CreateCollectionWithContext(ctx, createCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CreateCollection with error: Operation request error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the ObjectTrainingStatus model
				objectTrainingStatusModel := new(visualrecognitionv4.ObjectTrainingStatus)
				objectTrainingStatusModel.Ready = core.BoolPtr(true)
				objectTrainingStatusModel.InProgress = core.BoolPtr(true)
				objectTrainingStatusModel.DataChanged = core.BoolPtr(true)
				objectTrainingStatusModel.LatestFailed = core.BoolPtr(true)
				objectTrainingStatusModel.RscnnReady = core.BoolPtr(true)
				objectTrainingStatusModel.Description = core.StringPtr("testString")

				// Construct an instance of the TrainingStatus model
				trainingStatusModel := new(visualrecognitionv4.TrainingStatus)
				trainingStatusModel.Objects = objectTrainingStatusModel

				// Construct an instance of the CreateCollectionOptions model
				createCollectionOptionsModel := new(visualrecognitionv4.CreateCollectionOptions)
				createCollectionOptionsModel.Name = core.StringPtr("testString")
				createCollectionOptionsModel.Description = core.StringPtr("testString")
				createCollectionOptionsModel.TrainingStatus = trainingStatusModel
				createCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := visualRecognitionService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := visualRecognitionService.CreateCollection(createCollectionOptionsModel)
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
	Describe(`ListCollections(listCollectionsOptions *ListCollectionsOptions) - Operation response error`, func() {
		version := "testString"
		listCollectionsPath := "/v4/collections"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCollectionsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListCollections with error: Operation response processing error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the ListCollectionsOptions model
				listCollectionsOptionsModel := new(visualrecognitionv4.ListCollectionsOptions)
				listCollectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := visualRecognitionService.ListCollections(listCollectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				visualRecognitionService.EnableRetries(0, 0)
				result, response, operationErr = visualRecognitionService.ListCollections(listCollectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListCollections(listCollectionsOptions *ListCollectionsOptions)`, func() {
		version := "testString"
		listCollectionsPath := "/v4/collections"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCollectionsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"collections": [{"collection_id": "CollectionID", "name": "Name", "description": "Description", "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "image_count": 10, "training_status": {"objects": {"ready": false, "in_progress": true, "data_changed": false, "latest_failed": true, "rscnn_ready": true, "description": "Description"}}}]}`)
				}))
			})
			It(`Invoke ListCollections successfully`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())
				visualRecognitionService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := visualRecognitionService.ListCollections(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListCollectionsOptions model
				listCollectionsOptionsModel := new(visualrecognitionv4.ListCollectionsOptions)
				listCollectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = visualRecognitionService.ListCollections(listCollectionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.ListCollectionsWithContext(ctx, listCollectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				visualRecognitionService.DisableRetries()
				result, response, operationErr = visualRecognitionService.ListCollections(listCollectionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.ListCollectionsWithContext(ctx, listCollectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListCollections with error: Operation request error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the ListCollectionsOptions model
				listCollectionsOptionsModel := new(visualrecognitionv4.ListCollectionsOptions)
				listCollectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := visualRecognitionService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := visualRecognitionService.ListCollections(listCollectionsOptionsModel)
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
	Describe(`GetCollection(getCollectionOptions *GetCollectionOptions) - Operation response error`, func() {
		version := "testString"
		getCollectionPath := "/v4/collections/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCollectionPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCollection with error: Operation response processing error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the GetCollectionOptions model
				getCollectionOptionsModel := new(visualrecognitionv4.GetCollectionOptions)
				getCollectionOptionsModel.CollectionID = core.StringPtr("testString")
				getCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := visualRecognitionService.GetCollection(getCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				visualRecognitionService.EnableRetries(0, 0)
				result, response, operationErr = visualRecognitionService.GetCollection(getCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetCollection(getCollectionOptions *GetCollectionOptions)`, func() {
		version := "testString"
		getCollectionPath := "/v4/collections/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCollectionPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"collection_id": "CollectionID", "name": "Name", "description": "Description", "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "image_count": 10, "training_status": {"objects": {"ready": false, "in_progress": true, "data_changed": false, "latest_failed": true, "rscnn_ready": true, "description": "Description"}}}`)
				}))
			})
			It(`Invoke GetCollection successfully`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())
				visualRecognitionService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := visualRecognitionService.GetCollection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCollectionOptions model
				getCollectionOptionsModel := new(visualrecognitionv4.GetCollectionOptions)
				getCollectionOptionsModel.CollectionID = core.StringPtr("testString")
				getCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = visualRecognitionService.GetCollection(getCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.GetCollectionWithContext(ctx, getCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				visualRecognitionService.DisableRetries()
				result, response, operationErr = visualRecognitionService.GetCollection(getCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.GetCollectionWithContext(ctx, getCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetCollection with error: Operation validation and request error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the GetCollectionOptions model
				getCollectionOptionsModel := new(visualrecognitionv4.GetCollectionOptions)
				getCollectionOptionsModel.CollectionID = core.StringPtr("testString")
				getCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := visualRecognitionService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := visualRecognitionService.GetCollection(getCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetCollectionOptions model with no property values
				getCollectionOptionsModelNew := new(visualrecognitionv4.GetCollectionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = visualRecognitionService.GetCollection(getCollectionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateCollection(updateCollectionOptions *UpdateCollectionOptions) - Operation response error`, func() {
		version := "testString"
		updateCollectionPath := "/v4/collections/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCollectionPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateCollection with error: Operation response processing error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the ObjectTrainingStatus model
				objectTrainingStatusModel := new(visualrecognitionv4.ObjectTrainingStatus)
				objectTrainingStatusModel.Ready = core.BoolPtr(true)
				objectTrainingStatusModel.InProgress = core.BoolPtr(true)
				objectTrainingStatusModel.DataChanged = core.BoolPtr(true)
				objectTrainingStatusModel.LatestFailed = core.BoolPtr(true)
				objectTrainingStatusModel.RscnnReady = core.BoolPtr(true)
				objectTrainingStatusModel.Description = core.StringPtr("testString")

				// Construct an instance of the TrainingStatus model
				trainingStatusModel := new(visualrecognitionv4.TrainingStatus)
				trainingStatusModel.Objects = objectTrainingStatusModel

				// Construct an instance of the UpdateCollectionOptions model
				updateCollectionOptionsModel := new(visualrecognitionv4.UpdateCollectionOptions)
				updateCollectionOptionsModel.CollectionID = core.StringPtr("testString")
				updateCollectionOptionsModel.Name = core.StringPtr("testString")
				updateCollectionOptionsModel.Description = core.StringPtr("testString")
				updateCollectionOptionsModel.TrainingStatus = trainingStatusModel
				updateCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := visualRecognitionService.UpdateCollection(updateCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				visualRecognitionService.EnableRetries(0, 0)
				result, response, operationErr = visualRecognitionService.UpdateCollection(updateCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateCollection(updateCollectionOptions *UpdateCollectionOptions)`, func() {
		version := "testString"
		updateCollectionPath := "/v4/collections/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCollectionPath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"collection_id": "CollectionID", "name": "Name", "description": "Description", "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "image_count": 10, "training_status": {"objects": {"ready": false, "in_progress": true, "data_changed": false, "latest_failed": true, "rscnn_ready": true, "description": "Description"}}}`)
				}))
			})
			It(`Invoke UpdateCollection successfully`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())
				visualRecognitionService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := visualRecognitionService.UpdateCollection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ObjectTrainingStatus model
				objectTrainingStatusModel := new(visualrecognitionv4.ObjectTrainingStatus)
				objectTrainingStatusModel.Ready = core.BoolPtr(true)
				objectTrainingStatusModel.InProgress = core.BoolPtr(true)
				objectTrainingStatusModel.DataChanged = core.BoolPtr(true)
				objectTrainingStatusModel.LatestFailed = core.BoolPtr(true)
				objectTrainingStatusModel.RscnnReady = core.BoolPtr(true)
				objectTrainingStatusModel.Description = core.StringPtr("testString")

				// Construct an instance of the TrainingStatus model
				trainingStatusModel := new(visualrecognitionv4.TrainingStatus)
				trainingStatusModel.Objects = objectTrainingStatusModel

				// Construct an instance of the UpdateCollectionOptions model
				updateCollectionOptionsModel := new(visualrecognitionv4.UpdateCollectionOptions)
				updateCollectionOptionsModel.CollectionID = core.StringPtr("testString")
				updateCollectionOptionsModel.Name = core.StringPtr("testString")
				updateCollectionOptionsModel.Description = core.StringPtr("testString")
				updateCollectionOptionsModel.TrainingStatus = trainingStatusModel
				updateCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = visualRecognitionService.UpdateCollection(updateCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.UpdateCollectionWithContext(ctx, updateCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				visualRecognitionService.DisableRetries()
				result, response, operationErr = visualRecognitionService.UpdateCollection(updateCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.UpdateCollectionWithContext(ctx, updateCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateCollection with error: Operation validation and request error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the ObjectTrainingStatus model
				objectTrainingStatusModel := new(visualrecognitionv4.ObjectTrainingStatus)
				objectTrainingStatusModel.Ready = core.BoolPtr(true)
				objectTrainingStatusModel.InProgress = core.BoolPtr(true)
				objectTrainingStatusModel.DataChanged = core.BoolPtr(true)
				objectTrainingStatusModel.LatestFailed = core.BoolPtr(true)
				objectTrainingStatusModel.RscnnReady = core.BoolPtr(true)
				objectTrainingStatusModel.Description = core.StringPtr("testString")

				// Construct an instance of the TrainingStatus model
				trainingStatusModel := new(visualrecognitionv4.TrainingStatus)
				trainingStatusModel.Objects = objectTrainingStatusModel

				// Construct an instance of the UpdateCollectionOptions model
				updateCollectionOptionsModel := new(visualrecognitionv4.UpdateCollectionOptions)
				updateCollectionOptionsModel.CollectionID = core.StringPtr("testString")
				updateCollectionOptionsModel.Name = core.StringPtr("testString")
				updateCollectionOptionsModel.Description = core.StringPtr("testString")
				updateCollectionOptionsModel.TrainingStatus = trainingStatusModel
				updateCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := visualRecognitionService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := visualRecognitionService.UpdateCollection(updateCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateCollectionOptions model with no property values
				updateCollectionOptionsModelNew := new(visualrecognitionv4.UpdateCollectionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = visualRecognitionService.UpdateCollection(updateCollectionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteCollection(deleteCollectionOptions *DeleteCollectionOptions)`, func() {
		version := "testString"
		deleteCollectionPath := "/v4/collections/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteCollectionPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteCollection successfully`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())
				visualRecognitionService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := visualRecognitionService.DeleteCollection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteCollectionOptions model
				deleteCollectionOptionsModel := new(visualrecognitionv4.DeleteCollectionOptions)
				deleteCollectionOptionsModel.CollectionID = core.StringPtr("testString")
				deleteCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = visualRecognitionService.DeleteCollection(deleteCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				visualRecognitionService.DisableRetries()
				response, operationErr = visualRecognitionService.DeleteCollection(deleteCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteCollection with error: Operation validation and request error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the DeleteCollectionOptions model
				deleteCollectionOptionsModel := new(visualrecognitionv4.DeleteCollectionOptions)
				deleteCollectionOptionsModel.CollectionID = core.StringPtr("testString")
				deleteCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := visualRecognitionService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := visualRecognitionService.DeleteCollection(deleteCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteCollectionOptions model with no property values
				deleteCollectionOptionsModelNew := new(visualrecognitionv4.DeleteCollectionOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = visualRecognitionService.DeleteCollection(deleteCollectionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetModelFile(getModelFileOptions *GetModelFileOptions)`, func() {
		version := "testString"
		getModelFilePath := "/v4/collections/testString/model"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getModelFilePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["feature"]).To(Equal([]string{"objects"}))

					Expect(req.URL.Query()["model_format"]).To(Equal([]string{"rscnn"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/octet-stream")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke GetModelFile successfully`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())
				visualRecognitionService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := visualRecognitionService.GetModelFile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetModelFileOptions model
				getModelFileOptionsModel := new(visualrecognitionv4.GetModelFileOptions)
				getModelFileOptionsModel.CollectionID = core.StringPtr("testString")
				getModelFileOptionsModel.Feature = core.StringPtr("objects")
				getModelFileOptionsModel.ModelFormat = core.StringPtr("rscnn")
				getModelFileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = visualRecognitionService.GetModelFile(getModelFileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.GetModelFileWithContext(ctx, getModelFileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				visualRecognitionService.DisableRetries()
				result, response, operationErr = visualRecognitionService.GetModelFile(getModelFileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.GetModelFileWithContext(ctx, getModelFileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetModelFile with error: Operation validation and request error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the GetModelFileOptions model
				getModelFileOptionsModel := new(visualrecognitionv4.GetModelFileOptions)
				getModelFileOptionsModel.CollectionID = core.StringPtr("testString")
				getModelFileOptionsModel.Feature = core.StringPtr("objects")
				getModelFileOptionsModel.ModelFormat = core.StringPtr("rscnn")
				getModelFileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := visualRecognitionService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := visualRecognitionService.GetModelFile(getModelFileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetModelFileOptions model with no property values
				getModelFileOptionsModelNew := new(visualrecognitionv4.GetModelFileOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = visualRecognitionService.GetModelFile(getModelFileOptionsModelNew)
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
			visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
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
			visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
				URL: "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(visualRecognitionService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
				URL: "https://visualrecognitionv4/api",
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
			visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{})
			Expect(visualRecognitionService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"WATSON_VISION_COMBINED_URL": "https://visualrecognitionv4/api",
				"WATSON_VISION_COMBINED_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					Version: core.StringPtr(version),
				})
				Expect(visualRecognitionService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := visualRecognitionService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != visualRecognitionService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(visualRecognitionService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(visualRecognitionService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL: "https://testService/api",
					Version: core.StringPtr(version),
				})
				Expect(visualRecognitionService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := visualRecognitionService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != visualRecognitionService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(visualRecognitionService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(visualRecognitionService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					Version: core.StringPtr(version),
				})
				err := visualRecognitionService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := visualRecognitionService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != visualRecognitionService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(visualRecognitionService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(visualRecognitionService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"WATSON_VISION_COMBINED_URL": "https://visualrecognitionv4/api",
				"WATSON_VISION_COMBINED_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
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
			visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
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
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = visualrecognitionv4.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`AddImages(addImagesOptions *AddImagesOptions) - Operation response error`, func() {
		version := "testString"
		addImagesPath := "/v4/collections/testString/images"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addImagesPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke AddImages with error: Operation response processing error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the AddImagesOptions model
				addImagesOptionsModel := new(visualrecognitionv4.AddImagesOptions)
				addImagesOptionsModel.CollectionID = core.StringPtr("testString")
				addImagesOptionsModel.ImagesFile = []visualrecognitionv4.FileWithMetadata{visualrecognitionv4.FileWithMetadata{Data: CreateMockReader("This is a mock file."), Filename: core.StringPtr("mockfilename.txt"), }}
				addImagesOptionsModel.ImageURL = []string{"testString"}
				addImagesOptionsModel.TrainingData = core.StringPtr(`{"objects":[{"object":"2018-Fit","location":{"left":33,"top":8,"width":760,"height":419}}]}`)
				addImagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := visualRecognitionService.AddImages(addImagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				visualRecognitionService.EnableRetries(0, 0)
				result, response, operationErr = visualRecognitionService.AddImages(addImagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`AddImages(addImagesOptions *AddImagesOptions)`, func() {
		version := "testString"
		addImagesPath := "/v4/collections/testString/images"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addImagesPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"images": [{"image_id": "ImageID", "updated": "2019-01-01T12:00:00", "created": "2019-01-01T12:00:00", "source": {"type": "file", "filename": "Filename", "archive_filename": "ArchiveFilename", "source_url": "SourceURL", "resolved_url": "ResolvedURL"}, "dimensions": {"height": 6, "width": 5}, "errors": [{"code": "invalid_field", "message": "Message", "more_info": "MoreInfo", "target": {"type": "field", "name": "Name"}}], "training_data": {"objects": [{"object": "Object", "location": {"top": 3, "left": 4, "width": 5, "height": 6}}]}}], "warnings": [{"code": "invalid_field", "message": "Message", "more_info": "MoreInfo"}], "trace": "Trace"}`)
				}))
			})
			It(`Invoke AddImages successfully`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())
				visualRecognitionService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := visualRecognitionService.AddImages(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AddImagesOptions model
				addImagesOptionsModel := new(visualrecognitionv4.AddImagesOptions)
				addImagesOptionsModel.CollectionID = core.StringPtr("testString")
				addImagesOptionsModel.ImagesFile = []visualrecognitionv4.FileWithMetadata{visualrecognitionv4.FileWithMetadata{Data: CreateMockReader("This is a mock file."), Filename: core.StringPtr("mockfilename.txt"), }}
				addImagesOptionsModel.ImageURL = []string{"testString"}
				addImagesOptionsModel.TrainingData = core.StringPtr(`{"objects":[{"object":"2018-Fit","location":{"left":33,"top":8,"width":760,"height":419}}]}`)
				addImagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = visualRecognitionService.AddImages(addImagesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.AddImagesWithContext(ctx, addImagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				visualRecognitionService.DisableRetries()
				result, response, operationErr = visualRecognitionService.AddImages(addImagesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.AddImagesWithContext(ctx, addImagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke AddImages with error: Param validation error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:  testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the AddImagesOptions model
				addImagesOptionsModel := new(visualrecognitionv4.AddImagesOptions)
				// Invoke operation with invalid options model (negative test)
				result, response, operationErr := visualRecognitionService.AddImages(addImagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			It(`Invoke AddImages with error: Operation validation and request error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the AddImagesOptions model
				addImagesOptionsModel := new(visualrecognitionv4.AddImagesOptions)
				addImagesOptionsModel.CollectionID = core.StringPtr("testString")
				addImagesOptionsModel.ImagesFile = []visualrecognitionv4.FileWithMetadata{visualrecognitionv4.FileWithMetadata{Data: CreateMockReader("This is a mock file."), Filename: core.StringPtr("mockfilename.txt"), }}
				addImagesOptionsModel.ImageURL = []string{"testString"}
				addImagesOptionsModel.TrainingData = core.StringPtr(`{"objects":[{"object":"2018-Fit","location":{"left":33,"top":8,"width":760,"height":419}}]}`)
				addImagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := visualRecognitionService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := visualRecognitionService.AddImages(addImagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the AddImagesOptions model with no property values
				addImagesOptionsModelNew := new(visualrecognitionv4.AddImagesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = visualRecognitionService.AddImages(addImagesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListImages(listImagesOptions *ListImagesOptions) - Operation response error`, func() {
		version := "testString"
		listImagesPath := "/v4/collections/testString/images"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listImagesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListImages with error: Operation response processing error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the ListImagesOptions model
				listImagesOptionsModel := new(visualrecognitionv4.ListImagesOptions)
				listImagesOptionsModel.CollectionID = core.StringPtr("testString")
				listImagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := visualRecognitionService.ListImages(listImagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				visualRecognitionService.EnableRetries(0, 0)
				result, response, operationErr = visualRecognitionService.ListImages(listImagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListImages(listImagesOptions *ListImagesOptions)`, func() {
		version := "testString"
		listImagesPath := "/v4/collections/testString/images"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listImagesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"images": [{"image_id": "ImageID", "updated": "2019-01-01T12:00:00"}]}`)
				}))
			})
			It(`Invoke ListImages successfully`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())
				visualRecognitionService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := visualRecognitionService.ListImages(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListImagesOptions model
				listImagesOptionsModel := new(visualrecognitionv4.ListImagesOptions)
				listImagesOptionsModel.CollectionID = core.StringPtr("testString")
				listImagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = visualRecognitionService.ListImages(listImagesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.ListImagesWithContext(ctx, listImagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				visualRecognitionService.DisableRetries()
				result, response, operationErr = visualRecognitionService.ListImages(listImagesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.ListImagesWithContext(ctx, listImagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListImages with error: Operation validation and request error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the ListImagesOptions model
				listImagesOptionsModel := new(visualrecognitionv4.ListImagesOptions)
				listImagesOptionsModel.CollectionID = core.StringPtr("testString")
				listImagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := visualRecognitionService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := visualRecognitionService.ListImages(listImagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListImagesOptions model with no property values
				listImagesOptionsModelNew := new(visualrecognitionv4.ListImagesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = visualRecognitionService.ListImages(listImagesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetImageDetails(getImageDetailsOptions *GetImageDetailsOptions) - Operation response error`, func() {
		version := "testString"
		getImageDetailsPath := "/v4/collections/testString/images/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getImageDetailsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetImageDetails with error: Operation response processing error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the GetImageDetailsOptions model
				getImageDetailsOptionsModel := new(visualrecognitionv4.GetImageDetailsOptions)
				getImageDetailsOptionsModel.CollectionID = core.StringPtr("testString")
				getImageDetailsOptionsModel.ImageID = core.StringPtr("testString")
				getImageDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := visualRecognitionService.GetImageDetails(getImageDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				visualRecognitionService.EnableRetries(0, 0)
				result, response, operationErr = visualRecognitionService.GetImageDetails(getImageDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetImageDetails(getImageDetailsOptions *GetImageDetailsOptions)`, func() {
		version := "testString"
		getImageDetailsPath := "/v4/collections/testString/images/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getImageDetailsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"image_id": "ImageID", "updated": "2019-01-01T12:00:00", "created": "2019-01-01T12:00:00", "source": {"type": "file", "filename": "Filename", "archive_filename": "ArchiveFilename", "source_url": "SourceURL", "resolved_url": "ResolvedURL"}, "dimensions": {"height": 6, "width": 5}, "errors": [{"code": "invalid_field", "message": "Message", "more_info": "MoreInfo", "target": {"type": "field", "name": "Name"}}], "training_data": {"objects": [{"object": "Object", "location": {"top": 3, "left": 4, "width": 5, "height": 6}}]}}`)
				}))
			})
			It(`Invoke GetImageDetails successfully`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())
				visualRecognitionService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := visualRecognitionService.GetImageDetails(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetImageDetailsOptions model
				getImageDetailsOptionsModel := new(visualrecognitionv4.GetImageDetailsOptions)
				getImageDetailsOptionsModel.CollectionID = core.StringPtr("testString")
				getImageDetailsOptionsModel.ImageID = core.StringPtr("testString")
				getImageDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = visualRecognitionService.GetImageDetails(getImageDetailsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.GetImageDetailsWithContext(ctx, getImageDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				visualRecognitionService.DisableRetries()
				result, response, operationErr = visualRecognitionService.GetImageDetails(getImageDetailsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.GetImageDetailsWithContext(ctx, getImageDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetImageDetails with error: Operation validation and request error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the GetImageDetailsOptions model
				getImageDetailsOptionsModel := new(visualrecognitionv4.GetImageDetailsOptions)
				getImageDetailsOptionsModel.CollectionID = core.StringPtr("testString")
				getImageDetailsOptionsModel.ImageID = core.StringPtr("testString")
				getImageDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := visualRecognitionService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := visualRecognitionService.GetImageDetails(getImageDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetImageDetailsOptions model with no property values
				getImageDetailsOptionsModelNew := new(visualrecognitionv4.GetImageDetailsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = visualRecognitionService.GetImageDetails(getImageDetailsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteImage(deleteImageOptions *DeleteImageOptions)`, func() {
		version := "testString"
		deleteImagePath := "/v4/collections/testString/images/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteImagePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteImage successfully`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())
				visualRecognitionService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := visualRecognitionService.DeleteImage(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteImageOptions model
				deleteImageOptionsModel := new(visualrecognitionv4.DeleteImageOptions)
				deleteImageOptionsModel.CollectionID = core.StringPtr("testString")
				deleteImageOptionsModel.ImageID = core.StringPtr("testString")
				deleteImageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = visualRecognitionService.DeleteImage(deleteImageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				visualRecognitionService.DisableRetries()
				response, operationErr = visualRecognitionService.DeleteImage(deleteImageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteImage with error: Operation validation and request error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the DeleteImageOptions model
				deleteImageOptionsModel := new(visualrecognitionv4.DeleteImageOptions)
				deleteImageOptionsModel.CollectionID = core.StringPtr("testString")
				deleteImageOptionsModel.ImageID = core.StringPtr("testString")
				deleteImageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := visualRecognitionService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := visualRecognitionService.DeleteImage(deleteImageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteImageOptions model with no property values
				deleteImageOptionsModelNew := new(visualrecognitionv4.DeleteImageOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = visualRecognitionService.DeleteImage(deleteImageOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetJpegImage(getJpegImageOptions *GetJpegImageOptions)`, func() {
		version := "testString"
		getJpegImagePath := "/v4/collections/testString/images/testString/jpeg"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getJpegImagePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["size"]).To(Equal([]string{"full"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "image/jpeg")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke GetJpegImage successfully`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())
				visualRecognitionService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := visualRecognitionService.GetJpegImage(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetJpegImageOptions model
				getJpegImageOptionsModel := new(visualrecognitionv4.GetJpegImageOptions)
				getJpegImageOptionsModel.CollectionID = core.StringPtr("testString")
				getJpegImageOptionsModel.ImageID = core.StringPtr("testString")
				getJpegImageOptionsModel.Size = core.StringPtr("full")
				getJpegImageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = visualRecognitionService.GetJpegImage(getJpegImageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.GetJpegImageWithContext(ctx, getJpegImageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				visualRecognitionService.DisableRetries()
				result, response, operationErr = visualRecognitionService.GetJpegImage(getJpegImageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.GetJpegImageWithContext(ctx, getJpegImageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetJpegImage with error: Operation validation and request error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the GetJpegImageOptions model
				getJpegImageOptionsModel := new(visualrecognitionv4.GetJpegImageOptions)
				getJpegImageOptionsModel.CollectionID = core.StringPtr("testString")
				getJpegImageOptionsModel.ImageID = core.StringPtr("testString")
				getJpegImageOptionsModel.Size = core.StringPtr("full")
				getJpegImageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := visualRecognitionService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := visualRecognitionService.GetJpegImage(getJpegImageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetJpegImageOptions model with no property values
				getJpegImageOptionsModelNew := new(visualrecognitionv4.GetJpegImageOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = visualRecognitionService.GetJpegImage(getJpegImageOptionsModelNew)
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
			visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
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
			visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
				URL: "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(visualRecognitionService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
				URL: "https://visualrecognitionv4/api",
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
			visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{})
			Expect(visualRecognitionService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"WATSON_VISION_COMBINED_URL": "https://visualrecognitionv4/api",
				"WATSON_VISION_COMBINED_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					Version: core.StringPtr(version),
				})
				Expect(visualRecognitionService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := visualRecognitionService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != visualRecognitionService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(visualRecognitionService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(visualRecognitionService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL: "https://testService/api",
					Version: core.StringPtr(version),
				})
				Expect(visualRecognitionService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := visualRecognitionService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != visualRecognitionService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(visualRecognitionService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(visualRecognitionService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					Version: core.StringPtr(version),
				})
				err := visualRecognitionService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := visualRecognitionService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != visualRecognitionService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(visualRecognitionService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(visualRecognitionService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"WATSON_VISION_COMBINED_URL": "https://visualrecognitionv4/api",
				"WATSON_VISION_COMBINED_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
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
			visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
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
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = visualrecognitionv4.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListObjectMetadata(listObjectMetadataOptions *ListObjectMetadataOptions) - Operation response error`, func() {
		version := "testString"
		listObjectMetadataPath := "/v4/collections/testString/objects"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listObjectMetadataPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListObjectMetadata with error: Operation response processing error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the ListObjectMetadataOptions model
				listObjectMetadataOptionsModel := new(visualrecognitionv4.ListObjectMetadataOptions)
				listObjectMetadataOptionsModel.CollectionID = core.StringPtr("testString")
				listObjectMetadataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := visualRecognitionService.ListObjectMetadata(listObjectMetadataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				visualRecognitionService.EnableRetries(0, 0)
				result, response, operationErr = visualRecognitionService.ListObjectMetadata(listObjectMetadataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListObjectMetadata(listObjectMetadataOptions *ListObjectMetadataOptions)`, func() {
		version := "testString"
		listObjectMetadataPath := "/v4/collections/testString/objects"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listObjectMetadataPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"object_count": 11, "objects": [{"object": "Object", "count": 5}]}`)
				}))
			})
			It(`Invoke ListObjectMetadata successfully`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())
				visualRecognitionService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := visualRecognitionService.ListObjectMetadata(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListObjectMetadataOptions model
				listObjectMetadataOptionsModel := new(visualrecognitionv4.ListObjectMetadataOptions)
				listObjectMetadataOptionsModel.CollectionID = core.StringPtr("testString")
				listObjectMetadataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = visualRecognitionService.ListObjectMetadata(listObjectMetadataOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.ListObjectMetadataWithContext(ctx, listObjectMetadataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				visualRecognitionService.DisableRetries()
				result, response, operationErr = visualRecognitionService.ListObjectMetadata(listObjectMetadataOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.ListObjectMetadataWithContext(ctx, listObjectMetadataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListObjectMetadata with error: Operation validation and request error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the ListObjectMetadataOptions model
				listObjectMetadataOptionsModel := new(visualrecognitionv4.ListObjectMetadataOptions)
				listObjectMetadataOptionsModel.CollectionID = core.StringPtr("testString")
				listObjectMetadataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := visualRecognitionService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := visualRecognitionService.ListObjectMetadata(listObjectMetadataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListObjectMetadataOptions model with no property values
				listObjectMetadataOptionsModelNew := new(visualrecognitionv4.ListObjectMetadataOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = visualRecognitionService.ListObjectMetadata(listObjectMetadataOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateObjectMetadata(updateObjectMetadataOptions *UpdateObjectMetadataOptions) - Operation response error`, func() {
		version := "testString"
		updateObjectMetadataPath := "/v4/collections/testString/objects/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateObjectMetadataPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateObjectMetadata with error: Operation response processing error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the UpdateObjectMetadataOptions model
				updateObjectMetadataOptionsModel := new(visualrecognitionv4.UpdateObjectMetadataOptions)
				updateObjectMetadataOptionsModel.CollectionID = core.StringPtr("testString")
				updateObjectMetadataOptionsModel.Object = core.StringPtr("testString")
				updateObjectMetadataOptionsModel.NewObject = core.StringPtr("testString")
				updateObjectMetadataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := visualRecognitionService.UpdateObjectMetadata(updateObjectMetadataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				visualRecognitionService.EnableRetries(0, 0)
				result, response, operationErr = visualRecognitionService.UpdateObjectMetadata(updateObjectMetadataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateObjectMetadata(updateObjectMetadataOptions *UpdateObjectMetadataOptions)`, func() {
		version := "testString"
		updateObjectMetadataPath := "/v4/collections/testString/objects/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateObjectMetadataPath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"object": "Object", "count": 5}`)
				}))
			})
			It(`Invoke UpdateObjectMetadata successfully`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())
				visualRecognitionService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := visualRecognitionService.UpdateObjectMetadata(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateObjectMetadataOptions model
				updateObjectMetadataOptionsModel := new(visualrecognitionv4.UpdateObjectMetadataOptions)
				updateObjectMetadataOptionsModel.CollectionID = core.StringPtr("testString")
				updateObjectMetadataOptionsModel.Object = core.StringPtr("testString")
				updateObjectMetadataOptionsModel.NewObject = core.StringPtr("testString")
				updateObjectMetadataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = visualRecognitionService.UpdateObjectMetadata(updateObjectMetadataOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.UpdateObjectMetadataWithContext(ctx, updateObjectMetadataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				visualRecognitionService.DisableRetries()
				result, response, operationErr = visualRecognitionService.UpdateObjectMetadata(updateObjectMetadataOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.UpdateObjectMetadataWithContext(ctx, updateObjectMetadataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateObjectMetadata with error: Operation validation and request error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the UpdateObjectMetadataOptions model
				updateObjectMetadataOptionsModel := new(visualrecognitionv4.UpdateObjectMetadataOptions)
				updateObjectMetadataOptionsModel.CollectionID = core.StringPtr("testString")
				updateObjectMetadataOptionsModel.Object = core.StringPtr("testString")
				updateObjectMetadataOptionsModel.NewObject = core.StringPtr("testString")
				updateObjectMetadataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := visualRecognitionService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := visualRecognitionService.UpdateObjectMetadata(updateObjectMetadataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateObjectMetadataOptions model with no property values
				updateObjectMetadataOptionsModelNew := new(visualrecognitionv4.UpdateObjectMetadataOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = visualRecognitionService.UpdateObjectMetadata(updateObjectMetadataOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetObjectMetadata(getObjectMetadataOptions *GetObjectMetadataOptions) - Operation response error`, func() {
		version := "testString"
		getObjectMetadataPath := "/v4/collections/testString/objects/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getObjectMetadataPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetObjectMetadata with error: Operation response processing error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the GetObjectMetadataOptions model
				getObjectMetadataOptionsModel := new(visualrecognitionv4.GetObjectMetadataOptions)
				getObjectMetadataOptionsModel.CollectionID = core.StringPtr("testString")
				getObjectMetadataOptionsModel.Object = core.StringPtr("testString")
				getObjectMetadataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := visualRecognitionService.GetObjectMetadata(getObjectMetadataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				visualRecognitionService.EnableRetries(0, 0)
				result, response, operationErr = visualRecognitionService.GetObjectMetadata(getObjectMetadataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetObjectMetadata(getObjectMetadataOptions *GetObjectMetadataOptions)`, func() {
		version := "testString"
		getObjectMetadataPath := "/v4/collections/testString/objects/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getObjectMetadataPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"object": "Object", "count": 5}`)
				}))
			})
			It(`Invoke GetObjectMetadata successfully`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())
				visualRecognitionService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := visualRecognitionService.GetObjectMetadata(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetObjectMetadataOptions model
				getObjectMetadataOptionsModel := new(visualrecognitionv4.GetObjectMetadataOptions)
				getObjectMetadataOptionsModel.CollectionID = core.StringPtr("testString")
				getObjectMetadataOptionsModel.Object = core.StringPtr("testString")
				getObjectMetadataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = visualRecognitionService.GetObjectMetadata(getObjectMetadataOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.GetObjectMetadataWithContext(ctx, getObjectMetadataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				visualRecognitionService.DisableRetries()
				result, response, operationErr = visualRecognitionService.GetObjectMetadata(getObjectMetadataOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.GetObjectMetadataWithContext(ctx, getObjectMetadataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetObjectMetadata with error: Operation validation and request error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the GetObjectMetadataOptions model
				getObjectMetadataOptionsModel := new(visualrecognitionv4.GetObjectMetadataOptions)
				getObjectMetadataOptionsModel.CollectionID = core.StringPtr("testString")
				getObjectMetadataOptionsModel.Object = core.StringPtr("testString")
				getObjectMetadataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := visualRecognitionService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := visualRecognitionService.GetObjectMetadata(getObjectMetadataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetObjectMetadataOptions model with no property values
				getObjectMetadataOptionsModelNew := new(visualrecognitionv4.GetObjectMetadataOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = visualRecognitionService.GetObjectMetadata(getObjectMetadataOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteObject(deleteObjectOptions *DeleteObjectOptions)`, func() {
		version := "testString"
		deleteObjectPath := "/v4/collections/testString/objects/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteObjectPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteObject successfully`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())
				visualRecognitionService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := visualRecognitionService.DeleteObject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteObjectOptions model
				deleteObjectOptionsModel := new(visualrecognitionv4.DeleteObjectOptions)
				deleteObjectOptionsModel.CollectionID = core.StringPtr("testString")
				deleteObjectOptionsModel.Object = core.StringPtr("testString")
				deleteObjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = visualRecognitionService.DeleteObject(deleteObjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				visualRecognitionService.DisableRetries()
				response, operationErr = visualRecognitionService.DeleteObject(deleteObjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteObject with error: Operation validation and request error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the DeleteObjectOptions model
				deleteObjectOptionsModel := new(visualrecognitionv4.DeleteObjectOptions)
				deleteObjectOptionsModel.CollectionID = core.StringPtr("testString")
				deleteObjectOptionsModel.Object = core.StringPtr("testString")
				deleteObjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := visualRecognitionService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := visualRecognitionService.DeleteObject(deleteObjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteObjectOptions model with no property values
				deleteObjectOptionsModelNew := new(visualrecognitionv4.DeleteObjectOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = visualRecognitionService.DeleteObject(deleteObjectOptionsModelNew)
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
			visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
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
			visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
				URL: "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(visualRecognitionService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
				URL: "https://visualrecognitionv4/api",
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
			visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{})
			Expect(visualRecognitionService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"WATSON_VISION_COMBINED_URL": "https://visualrecognitionv4/api",
				"WATSON_VISION_COMBINED_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					Version: core.StringPtr(version),
				})
				Expect(visualRecognitionService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := visualRecognitionService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != visualRecognitionService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(visualRecognitionService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(visualRecognitionService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL: "https://testService/api",
					Version: core.StringPtr(version),
				})
				Expect(visualRecognitionService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := visualRecognitionService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != visualRecognitionService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(visualRecognitionService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(visualRecognitionService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					Version: core.StringPtr(version),
				})
				err := visualRecognitionService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := visualRecognitionService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != visualRecognitionService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(visualRecognitionService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(visualRecognitionService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"WATSON_VISION_COMBINED_URL": "https://visualrecognitionv4/api",
				"WATSON_VISION_COMBINED_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
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
			visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
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
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = visualrecognitionv4.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`Train(trainOptions *TrainOptions) - Operation response error`, func() {
		version := "testString"
		trainPath := "/v4/collections/testString/train"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(trainPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke Train with error: Operation response processing error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the TrainOptions model
				trainOptionsModel := new(visualrecognitionv4.TrainOptions)
				trainOptionsModel.CollectionID = core.StringPtr("testString")
				trainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := visualRecognitionService.Train(trainOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				visualRecognitionService.EnableRetries(0, 0)
				result, response, operationErr = visualRecognitionService.Train(trainOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`Train(trainOptions *TrainOptions)`, func() {
		version := "testString"
		trainPath := "/v4/collections/testString/train"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(trainPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"collection_id": "CollectionID", "name": "Name", "description": "Description", "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "image_count": 10, "training_status": {"objects": {"ready": false, "in_progress": true, "data_changed": false, "latest_failed": true, "rscnn_ready": true, "description": "Description"}}}`)
				}))
			})
			It(`Invoke Train successfully`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())
				visualRecognitionService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := visualRecognitionService.Train(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TrainOptions model
				trainOptionsModel := new(visualrecognitionv4.TrainOptions)
				trainOptionsModel.CollectionID = core.StringPtr("testString")
				trainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = visualRecognitionService.Train(trainOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.TrainWithContext(ctx, trainOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				visualRecognitionService.DisableRetries()
				result, response, operationErr = visualRecognitionService.Train(trainOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.TrainWithContext(ctx, trainOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke Train with error: Operation validation and request error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the TrainOptions model
				trainOptionsModel := new(visualrecognitionv4.TrainOptions)
				trainOptionsModel.CollectionID = core.StringPtr("testString")
				trainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := visualRecognitionService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := visualRecognitionService.Train(trainOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the TrainOptions model with no property values
				trainOptionsModelNew := new(visualrecognitionv4.TrainOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = visualRecognitionService.Train(trainOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AddImageTrainingData(addImageTrainingDataOptions *AddImageTrainingDataOptions) - Operation response error`, func() {
		version := "testString"
		addImageTrainingDataPath := "/v4/collections/testString/images/testString/training_data"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addImageTrainingDataPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke AddImageTrainingData with error: Operation response processing error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the Location model
				locationModel := new(visualrecognitionv4.Location)
				locationModel.Top = core.Int64Ptr(int64(38))
				locationModel.Left = core.Int64Ptr(int64(38))
				locationModel.Width = core.Int64Ptr(int64(38))
				locationModel.Height = core.Int64Ptr(int64(38))

				// Construct an instance of the TrainingDataObject model
				trainingDataObjectModel := new(visualrecognitionv4.TrainingDataObject)
				trainingDataObjectModel.Object = core.StringPtr("testString")
				trainingDataObjectModel.Location = locationModel

				// Construct an instance of the AddImageTrainingDataOptions model
				addImageTrainingDataOptionsModel := new(visualrecognitionv4.AddImageTrainingDataOptions)
				addImageTrainingDataOptionsModel.CollectionID = core.StringPtr("testString")
				addImageTrainingDataOptionsModel.ImageID = core.StringPtr("testString")
				addImageTrainingDataOptionsModel.Objects = []visualrecognitionv4.TrainingDataObject{*trainingDataObjectModel}
				addImageTrainingDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := visualRecognitionService.AddImageTrainingData(addImageTrainingDataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				visualRecognitionService.EnableRetries(0, 0)
				result, response, operationErr = visualRecognitionService.AddImageTrainingData(addImageTrainingDataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`AddImageTrainingData(addImageTrainingDataOptions *AddImageTrainingDataOptions)`, func() {
		version := "testString"
		addImageTrainingDataPath := "/v4/collections/testString/images/testString/training_data"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addImageTrainingDataPath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"objects": [{"object": "Object", "location": {"top": 3, "left": 4, "width": 5, "height": 6}}]}`)
				}))
			})
			It(`Invoke AddImageTrainingData successfully`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())
				visualRecognitionService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := visualRecognitionService.AddImageTrainingData(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Location model
				locationModel := new(visualrecognitionv4.Location)
				locationModel.Top = core.Int64Ptr(int64(38))
				locationModel.Left = core.Int64Ptr(int64(38))
				locationModel.Width = core.Int64Ptr(int64(38))
				locationModel.Height = core.Int64Ptr(int64(38))

				// Construct an instance of the TrainingDataObject model
				trainingDataObjectModel := new(visualrecognitionv4.TrainingDataObject)
				trainingDataObjectModel.Object = core.StringPtr("testString")
				trainingDataObjectModel.Location = locationModel

				// Construct an instance of the AddImageTrainingDataOptions model
				addImageTrainingDataOptionsModel := new(visualrecognitionv4.AddImageTrainingDataOptions)
				addImageTrainingDataOptionsModel.CollectionID = core.StringPtr("testString")
				addImageTrainingDataOptionsModel.ImageID = core.StringPtr("testString")
				addImageTrainingDataOptionsModel.Objects = []visualrecognitionv4.TrainingDataObject{*trainingDataObjectModel}
				addImageTrainingDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = visualRecognitionService.AddImageTrainingData(addImageTrainingDataOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.AddImageTrainingDataWithContext(ctx, addImageTrainingDataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				visualRecognitionService.DisableRetries()
				result, response, operationErr = visualRecognitionService.AddImageTrainingData(addImageTrainingDataOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.AddImageTrainingDataWithContext(ctx, addImageTrainingDataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke AddImageTrainingData with error: Operation validation and request error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the Location model
				locationModel := new(visualrecognitionv4.Location)
				locationModel.Top = core.Int64Ptr(int64(38))
				locationModel.Left = core.Int64Ptr(int64(38))
				locationModel.Width = core.Int64Ptr(int64(38))
				locationModel.Height = core.Int64Ptr(int64(38))

				// Construct an instance of the TrainingDataObject model
				trainingDataObjectModel := new(visualrecognitionv4.TrainingDataObject)
				trainingDataObjectModel.Object = core.StringPtr("testString")
				trainingDataObjectModel.Location = locationModel

				// Construct an instance of the AddImageTrainingDataOptions model
				addImageTrainingDataOptionsModel := new(visualrecognitionv4.AddImageTrainingDataOptions)
				addImageTrainingDataOptionsModel.CollectionID = core.StringPtr("testString")
				addImageTrainingDataOptionsModel.ImageID = core.StringPtr("testString")
				addImageTrainingDataOptionsModel.Objects = []visualrecognitionv4.TrainingDataObject{*trainingDataObjectModel}
				addImageTrainingDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := visualRecognitionService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := visualRecognitionService.AddImageTrainingData(addImageTrainingDataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the AddImageTrainingDataOptions model with no property values
				addImageTrainingDataOptionsModelNew := new(visualrecognitionv4.AddImageTrainingDataOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = visualRecognitionService.AddImageTrainingData(addImageTrainingDataOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTrainingUsage(getTrainingUsageOptions *GetTrainingUsageOptions) - Operation response error`, func() {
		version := "testString"
		getTrainingUsagePath := "/v4/training_usage"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTrainingUsagePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))


					// TODO: Add check for start_time query parameter


					// TODO: Add check for end_time query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTrainingUsage with error: Operation response processing error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the GetTrainingUsageOptions model
				getTrainingUsageOptionsModel := new(visualrecognitionv4.GetTrainingUsageOptions)
				getTrainingUsageOptionsModel.StartTime = CreateMockDate()
				getTrainingUsageOptionsModel.EndTime = CreateMockDate()
				getTrainingUsageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := visualRecognitionService.GetTrainingUsage(getTrainingUsageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				visualRecognitionService.EnableRetries(0, 0)
				result, response, operationErr = visualRecognitionService.GetTrainingUsage(getTrainingUsageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetTrainingUsage(getTrainingUsageOptions *GetTrainingUsageOptions)`, func() {
		version := "testString"
		getTrainingUsagePath := "/v4/training_usage"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTrainingUsagePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))


					// TODO: Add check for start_time query parameter


					// TODO: Add check for end_time query parameter

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"start_time": "2019-01-01T12:00:00", "end_time": "2019-01-01T12:00:00", "completed_events": 15, "trained_images": 13, "events": [{"type": "objects", "collection_id": "CollectionID", "completion_time": "2019-01-01T12:00:00", "status": "failed", "image_count": 10}]}`)
				}))
			})
			It(`Invoke GetTrainingUsage successfully`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())
				visualRecognitionService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := visualRecognitionService.GetTrainingUsage(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTrainingUsageOptions model
				getTrainingUsageOptionsModel := new(visualrecognitionv4.GetTrainingUsageOptions)
				getTrainingUsageOptionsModel.StartTime = CreateMockDate()
				getTrainingUsageOptionsModel.EndTime = CreateMockDate()
				getTrainingUsageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = visualRecognitionService.GetTrainingUsage(getTrainingUsageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.GetTrainingUsageWithContext(ctx, getTrainingUsageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				visualRecognitionService.DisableRetries()
				result, response, operationErr = visualRecognitionService.GetTrainingUsage(getTrainingUsageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = visualRecognitionService.GetTrainingUsageWithContext(ctx, getTrainingUsageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetTrainingUsage with error: Operation request error`, func() {
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the GetTrainingUsageOptions model
				getTrainingUsageOptionsModel := new(visualrecognitionv4.GetTrainingUsageOptions)
				getTrainingUsageOptionsModel.StartTime = CreateMockDate()
				getTrainingUsageOptionsModel.EndTime = CreateMockDate()
				getTrainingUsageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := visualRecognitionService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := visualRecognitionService.GetTrainingUsage(getTrainingUsageOptionsModel)
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
			visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
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
			visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
				URL: "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(visualRecognitionService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
				URL: "https://visualrecognitionv4/api",
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
			visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{})
			Expect(visualRecognitionService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"WATSON_VISION_COMBINED_URL": "https://visualrecognitionv4/api",
				"WATSON_VISION_COMBINED_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					Version: core.StringPtr(version),
				})
				Expect(visualRecognitionService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := visualRecognitionService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != visualRecognitionService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(visualRecognitionService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(visualRecognitionService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL: "https://testService/api",
					Version: core.StringPtr(version),
				})
				Expect(visualRecognitionService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := visualRecognitionService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != visualRecognitionService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(visualRecognitionService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(visualRecognitionService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					Version: core.StringPtr(version),
				})
				err := visualRecognitionService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := visualRecognitionService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != visualRecognitionService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(visualRecognitionService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(visualRecognitionService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"WATSON_VISION_COMBINED_URL": "https://visualrecognitionv4/api",
				"WATSON_VISION_COMBINED_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
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
			visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
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
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = visualrecognitionv4.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})

	Describe(`DeleteUserData(deleteUserDataOptions *DeleteUserDataOptions)`, func() {
		version := "testString"
		deleteUserDataPath := "/v4/user_data"
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
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
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
				deleteUserDataOptionsModel := new(visualrecognitionv4.DeleteUserDataOptions)
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
				visualRecognitionService, serviceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(visualRecognitionService).ToNot(BeNil())

				// Construct an instance of the DeleteUserDataOptions model
				deleteUserDataOptionsModel := new(visualrecognitionv4.DeleteUserDataOptions)
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
				deleteUserDataOptionsModelNew := new(visualrecognitionv4.DeleteUserDataOptions)
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
			visualRecognitionService, _ := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
				URL:           "http://visualrecognitionv4modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Version: core.StringPtr(version),
			})
			It(`Invoke NewAddImageTrainingDataOptions successfully`, func() {
				// Construct an instance of the Location model
				locationModel := new(visualrecognitionv4.Location)
				Expect(locationModel).ToNot(BeNil())
				locationModel.Top = core.Int64Ptr(int64(38))
				locationModel.Left = core.Int64Ptr(int64(38))
				locationModel.Width = core.Int64Ptr(int64(38))
				locationModel.Height = core.Int64Ptr(int64(38))
				Expect(locationModel.Top).To(Equal(core.Int64Ptr(int64(38))))
				Expect(locationModel.Left).To(Equal(core.Int64Ptr(int64(38))))
				Expect(locationModel.Width).To(Equal(core.Int64Ptr(int64(38))))
				Expect(locationModel.Height).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the TrainingDataObject model
				trainingDataObjectModel := new(visualrecognitionv4.TrainingDataObject)
				Expect(trainingDataObjectModel).ToNot(BeNil())
				trainingDataObjectModel.Object = core.StringPtr("testString")
				trainingDataObjectModel.Location = locationModel
				Expect(trainingDataObjectModel.Object).To(Equal(core.StringPtr("testString")))
				Expect(trainingDataObjectModel.Location).To(Equal(locationModel))

				// Construct an instance of the AddImageTrainingDataOptions model
				collectionID := "testString"
				imageID := "testString"
				addImageTrainingDataOptionsModel := visualRecognitionService.NewAddImageTrainingDataOptions(collectionID, imageID)
				addImageTrainingDataOptionsModel.SetCollectionID("testString")
				addImageTrainingDataOptionsModel.SetImageID("testString")
				addImageTrainingDataOptionsModel.SetObjects([]visualrecognitionv4.TrainingDataObject{*trainingDataObjectModel})
				addImageTrainingDataOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addImageTrainingDataOptionsModel).ToNot(BeNil())
				Expect(addImageTrainingDataOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(addImageTrainingDataOptionsModel.ImageID).To(Equal(core.StringPtr("testString")))
				Expect(addImageTrainingDataOptionsModel.Objects).To(Equal([]visualrecognitionv4.TrainingDataObject{*trainingDataObjectModel}))
				Expect(addImageTrainingDataOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewAddImagesOptions successfully`, func() {
				// Construct an instance of the FileWithMetadata model
				fileWithMetadataModel := new(visualrecognitionv4.FileWithMetadata)
				Expect(fileWithMetadataModel).ToNot(BeNil())
				fileWithMetadataModel.Data = CreateMockReader("This is a mock file.")
				fileWithMetadataModel.Filename = core.StringPtr("testString")
				fileWithMetadataModel.ContentType = core.StringPtr("testString")
				Expect(fileWithMetadataModel.Data).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(fileWithMetadataModel.Filename).To(Equal(core.StringPtr("testString")))
				Expect(fileWithMetadataModel.ContentType).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the AddImagesOptions model
				collectionID := "testString"
				addImagesOptionsModel := visualRecognitionService.NewAddImagesOptions(collectionID)
				addImagesOptionsModel.SetCollectionID("testString")
				addImagesOptionsModel.SetImagesFile([]visualrecognitionv4.FileWithMetadata{visualrecognitionv4.FileWithMetadata{Data: CreateMockReader("This is a mock file."), Filename: core.StringPtr("mockfilename.txt"), }})
				addImagesOptionsModel.SetImageURL([]string{"testString"})
				addImagesOptionsModel.SetTrainingData(`{"objects":[{"object":"2018-Fit","location":{"left":33,"top":8,"width":760,"height":419}}]}`)
				addImagesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addImagesOptionsModel).ToNot(BeNil())
				Expect(addImagesOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(addImagesOptionsModel.ImagesFile).To(Equal([]visualrecognitionv4.FileWithMetadata{visualrecognitionv4.FileWithMetadata{Data: CreateMockReader("This is a mock file."), Filename: core.StringPtr("mockfilename.txt"), }}))
				Expect(addImagesOptionsModel.ImageURL).To(Equal([]string{"testString"}))
				Expect(addImagesOptionsModel.TrainingData).To(Equal(core.StringPtr(`{"objects":[{"object":"2018-Fit","location":{"left":33,"top":8,"width":760,"height":419}}]}`)))
				Expect(addImagesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewAnalyzeOptions successfully`, func() {
				// Construct an instance of the FileWithMetadata model
				fileWithMetadataModel := new(visualrecognitionv4.FileWithMetadata)
				Expect(fileWithMetadataModel).ToNot(BeNil())
				fileWithMetadataModel.Data = CreateMockReader("This is a mock file.")
				fileWithMetadataModel.Filename = core.StringPtr("testString")
				fileWithMetadataModel.ContentType = core.StringPtr("testString")
				Expect(fileWithMetadataModel.Data).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(fileWithMetadataModel.Filename).To(Equal(core.StringPtr("testString")))
				Expect(fileWithMetadataModel.ContentType).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the AnalyzeOptions model
				collectionIds := []string{"testString"}
				features := []string{"objects"}
				analyzeOptionsModel := visualRecognitionService.NewAnalyzeOptions(collectionIds, features)
				analyzeOptionsModel.SetCollectionIds([]string{"testString"})
				analyzeOptionsModel.SetFeatures([]string{"objects"})
				analyzeOptionsModel.SetImagesFile([]visualrecognitionv4.FileWithMetadata{visualrecognitionv4.FileWithMetadata{Data: CreateMockReader("This is a mock file."), Filename: core.StringPtr("mockfilename.txt"), }})
				analyzeOptionsModel.SetImageURL([]string{"testString"})
				analyzeOptionsModel.SetThreshold(float32(0.15))
				analyzeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(analyzeOptionsModel).ToNot(BeNil())
				Expect(analyzeOptionsModel.CollectionIds).To(Equal([]string{"testString"}))
				Expect(analyzeOptionsModel.Features).To(Equal([]string{"objects"}))
				Expect(analyzeOptionsModel.ImagesFile).To(Equal([]visualrecognitionv4.FileWithMetadata{visualrecognitionv4.FileWithMetadata{Data: CreateMockReader("This is a mock file."), Filename: core.StringPtr("mockfilename.txt"), }}))
				Expect(analyzeOptionsModel.ImageURL).To(Equal([]string{"testString"}))
				Expect(analyzeOptionsModel.Threshold).To(Equal(core.Float32Ptr(float32(0.15))))
				Expect(analyzeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateCollectionOptions successfully`, func() {
				// Construct an instance of the ObjectTrainingStatus model
				objectTrainingStatusModel := new(visualrecognitionv4.ObjectTrainingStatus)
				Expect(objectTrainingStatusModel).ToNot(BeNil())
				objectTrainingStatusModel.Ready = core.BoolPtr(true)
				objectTrainingStatusModel.InProgress = core.BoolPtr(true)
				objectTrainingStatusModel.DataChanged = core.BoolPtr(true)
				objectTrainingStatusModel.LatestFailed = core.BoolPtr(true)
				objectTrainingStatusModel.RscnnReady = core.BoolPtr(true)
				objectTrainingStatusModel.Description = core.StringPtr("testString")
				Expect(objectTrainingStatusModel.Ready).To(Equal(core.BoolPtr(true)))
				Expect(objectTrainingStatusModel.InProgress).To(Equal(core.BoolPtr(true)))
				Expect(objectTrainingStatusModel.DataChanged).To(Equal(core.BoolPtr(true)))
				Expect(objectTrainingStatusModel.LatestFailed).To(Equal(core.BoolPtr(true)))
				Expect(objectTrainingStatusModel.RscnnReady).To(Equal(core.BoolPtr(true)))
				Expect(objectTrainingStatusModel.Description).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the TrainingStatus model
				trainingStatusModel := new(visualrecognitionv4.TrainingStatus)
				Expect(trainingStatusModel).ToNot(BeNil())
				trainingStatusModel.Objects = objectTrainingStatusModel
				Expect(trainingStatusModel.Objects).To(Equal(objectTrainingStatusModel))

				// Construct an instance of the CreateCollectionOptions model
				createCollectionOptionsModel := visualRecognitionService.NewCreateCollectionOptions()
				createCollectionOptionsModel.SetName("testString")
				createCollectionOptionsModel.SetDescription("testString")
				createCollectionOptionsModel.SetTrainingStatus(trainingStatusModel)
				createCollectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createCollectionOptionsModel).ToNot(BeNil())
				Expect(createCollectionOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createCollectionOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createCollectionOptionsModel.TrainingStatus).To(Equal(trainingStatusModel))
				Expect(createCollectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteCollectionOptions successfully`, func() {
				// Construct an instance of the DeleteCollectionOptions model
				collectionID := "testString"
				deleteCollectionOptionsModel := visualRecognitionService.NewDeleteCollectionOptions(collectionID)
				deleteCollectionOptionsModel.SetCollectionID("testString")
				deleteCollectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteCollectionOptionsModel).ToNot(BeNil())
				Expect(deleteCollectionOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCollectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteImageOptions successfully`, func() {
				// Construct an instance of the DeleteImageOptions model
				collectionID := "testString"
				imageID := "testString"
				deleteImageOptionsModel := visualRecognitionService.NewDeleteImageOptions(collectionID, imageID)
				deleteImageOptionsModel.SetCollectionID("testString")
				deleteImageOptionsModel.SetImageID("testString")
				deleteImageOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteImageOptionsModel).ToNot(BeNil())
				Expect(deleteImageOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteImageOptionsModel.ImageID).To(Equal(core.StringPtr("testString")))
				Expect(deleteImageOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteObjectOptions successfully`, func() {
				// Construct an instance of the DeleteObjectOptions model
				collectionID := "testString"
				object := "testString"
				deleteObjectOptionsModel := visualRecognitionService.NewDeleteObjectOptions(collectionID, object)
				deleteObjectOptionsModel.SetCollectionID("testString")
				deleteObjectOptionsModel.SetObject("testString")
				deleteObjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteObjectOptionsModel).ToNot(BeNil())
				Expect(deleteObjectOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteObjectOptionsModel.Object).To(Equal(core.StringPtr("testString")))
				Expect(deleteObjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			It(`Invoke NewFileWithMetadata successfully`, func() {
				data := CreateMockReader("This is a mock file.")
				model, err := visualRecognitionService.NewFileWithMetadata(data)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewGetCollectionOptions successfully`, func() {
				// Construct an instance of the GetCollectionOptions model
				collectionID := "testString"
				getCollectionOptionsModel := visualRecognitionService.NewGetCollectionOptions(collectionID)
				getCollectionOptionsModel.SetCollectionID("testString")
				getCollectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCollectionOptionsModel).ToNot(BeNil())
				Expect(getCollectionOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(getCollectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetImageDetailsOptions successfully`, func() {
				// Construct an instance of the GetImageDetailsOptions model
				collectionID := "testString"
				imageID := "testString"
				getImageDetailsOptionsModel := visualRecognitionService.NewGetImageDetailsOptions(collectionID, imageID)
				getImageDetailsOptionsModel.SetCollectionID("testString")
				getImageDetailsOptionsModel.SetImageID("testString")
				getImageDetailsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getImageDetailsOptionsModel).ToNot(BeNil())
				Expect(getImageDetailsOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(getImageDetailsOptionsModel.ImageID).To(Equal(core.StringPtr("testString")))
				Expect(getImageDetailsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetJpegImageOptions successfully`, func() {
				// Construct an instance of the GetJpegImageOptions model
				collectionID := "testString"
				imageID := "testString"
				getJpegImageOptionsModel := visualRecognitionService.NewGetJpegImageOptions(collectionID, imageID)
				getJpegImageOptionsModel.SetCollectionID("testString")
				getJpegImageOptionsModel.SetImageID("testString")
				getJpegImageOptionsModel.SetSize("full")
				getJpegImageOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getJpegImageOptionsModel).ToNot(BeNil())
				Expect(getJpegImageOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(getJpegImageOptionsModel.ImageID).To(Equal(core.StringPtr("testString")))
				Expect(getJpegImageOptionsModel.Size).To(Equal(core.StringPtr("full")))
				Expect(getJpegImageOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetModelFileOptions successfully`, func() {
				// Construct an instance of the GetModelFileOptions model
				collectionID := "testString"
				feature := "objects"
				modelFormat := "rscnn"
				getModelFileOptionsModel := visualRecognitionService.NewGetModelFileOptions(collectionID, feature, modelFormat)
				getModelFileOptionsModel.SetCollectionID("testString")
				getModelFileOptionsModel.SetFeature("objects")
				getModelFileOptionsModel.SetModelFormat("rscnn")
				getModelFileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getModelFileOptionsModel).ToNot(BeNil())
				Expect(getModelFileOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(getModelFileOptionsModel.Feature).To(Equal(core.StringPtr("objects")))
				Expect(getModelFileOptionsModel.ModelFormat).To(Equal(core.StringPtr("rscnn")))
				Expect(getModelFileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetObjectMetadataOptions successfully`, func() {
				// Construct an instance of the GetObjectMetadataOptions model
				collectionID := "testString"
				object := "testString"
				getObjectMetadataOptionsModel := visualRecognitionService.NewGetObjectMetadataOptions(collectionID, object)
				getObjectMetadataOptionsModel.SetCollectionID("testString")
				getObjectMetadataOptionsModel.SetObject("testString")
				getObjectMetadataOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getObjectMetadataOptionsModel).ToNot(BeNil())
				Expect(getObjectMetadataOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(getObjectMetadataOptionsModel.Object).To(Equal(core.StringPtr("testString")))
				Expect(getObjectMetadataOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTrainingUsageOptions successfully`, func() {
				// Construct an instance of the GetTrainingUsageOptions model
				getTrainingUsageOptionsModel := visualRecognitionService.NewGetTrainingUsageOptions()
				getTrainingUsageOptionsModel.SetStartTime(CreateMockDate())
				getTrainingUsageOptionsModel.SetEndTime(CreateMockDate())
				getTrainingUsageOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTrainingUsageOptionsModel).ToNot(BeNil())
				Expect(getTrainingUsageOptionsModel.StartTime).To(Equal(CreateMockDate()))
				Expect(getTrainingUsageOptionsModel.EndTime).To(Equal(CreateMockDate()))
				Expect(getTrainingUsageOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListCollectionsOptions successfully`, func() {
				// Construct an instance of the ListCollectionsOptions model
				listCollectionsOptionsModel := visualRecognitionService.NewListCollectionsOptions()
				listCollectionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listCollectionsOptionsModel).ToNot(BeNil())
				Expect(listCollectionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListImagesOptions successfully`, func() {
				// Construct an instance of the ListImagesOptions model
				collectionID := "testString"
				listImagesOptionsModel := visualRecognitionService.NewListImagesOptions(collectionID)
				listImagesOptionsModel.SetCollectionID("testString")
				listImagesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listImagesOptionsModel).ToNot(BeNil())
				Expect(listImagesOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(listImagesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListObjectMetadataOptions successfully`, func() {
				// Construct an instance of the ListObjectMetadataOptions model
				collectionID := "testString"
				listObjectMetadataOptionsModel := visualRecognitionService.NewListObjectMetadataOptions(collectionID)
				listObjectMetadataOptionsModel.SetCollectionID("testString")
				listObjectMetadataOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listObjectMetadataOptionsModel).ToNot(BeNil())
				Expect(listObjectMetadataOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(listObjectMetadataOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewLocation successfully`, func() {
				top := int64(38)
				left := int64(38)
				width := int64(38)
				height := int64(38)
				model, err := visualRecognitionService.NewLocation(top, left, width, height)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewObjectTrainingStatus successfully`, func() {
				ready := true
				inProgress := true
				dataChanged := true
				latestFailed := true
				rscnnReady := true
				description := "testString"
				model, err := visualRecognitionService.NewObjectTrainingStatus(ready, inProgress, dataChanged, latestFailed, rscnnReady, description)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewTrainOptions successfully`, func() {
				// Construct an instance of the TrainOptions model
				collectionID := "testString"
				trainOptionsModel := visualRecognitionService.NewTrainOptions(collectionID)
				trainOptionsModel.SetCollectionID("testString")
				trainOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(trainOptionsModel).ToNot(BeNil())
				Expect(trainOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(trainOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewTrainingStatus successfully`, func() {
				var objects *visualrecognitionv4.ObjectTrainingStatus = nil
				_, err := visualRecognitionService.NewTrainingStatus(objects)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewUpdateCollectionOptions successfully`, func() {
				// Construct an instance of the ObjectTrainingStatus model
				objectTrainingStatusModel := new(visualrecognitionv4.ObjectTrainingStatus)
				Expect(objectTrainingStatusModel).ToNot(BeNil())
				objectTrainingStatusModel.Ready = core.BoolPtr(true)
				objectTrainingStatusModel.InProgress = core.BoolPtr(true)
				objectTrainingStatusModel.DataChanged = core.BoolPtr(true)
				objectTrainingStatusModel.LatestFailed = core.BoolPtr(true)
				objectTrainingStatusModel.RscnnReady = core.BoolPtr(true)
				objectTrainingStatusModel.Description = core.StringPtr("testString")
				Expect(objectTrainingStatusModel.Ready).To(Equal(core.BoolPtr(true)))
				Expect(objectTrainingStatusModel.InProgress).To(Equal(core.BoolPtr(true)))
				Expect(objectTrainingStatusModel.DataChanged).To(Equal(core.BoolPtr(true)))
				Expect(objectTrainingStatusModel.LatestFailed).To(Equal(core.BoolPtr(true)))
				Expect(objectTrainingStatusModel.RscnnReady).To(Equal(core.BoolPtr(true)))
				Expect(objectTrainingStatusModel.Description).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the TrainingStatus model
				trainingStatusModel := new(visualrecognitionv4.TrainingStatus)
				Expect(trainingStatusModel).ToNot(BeNil())
				trainingStatusModel.Objects = objectTrainingStatusModel
				Expect(trainingStatusModel.Objects).To(Equal(objectTrainingStatusModel))

				// Construct an instance of the UpdateCollectionOptions model
				collectionID := "testString"
				updateCollectionOptionsModel := visualRecognitionService.NewUpdateCollectionOptions(collectionID)
				updateCollectionOptionsModel.SetCollectionID("testString")
				updateCollectionOptionsModel.SetName("testString")
				updateCollectionOptionsModel.SetDescription("testString")
				updateCollectionOptionsModel.SetTrainingStatus(trainingStatusModel)
				updateCollectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateCollectionOptionsModel).ToNot(BeNil())
				Expect(updateCollectionOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(updateCollectionOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateCollectionOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateCollectionOptionsModel.TrainingStatus).To(Equal(trainingStatusModel))
				Expect(updateCollectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateObjectMetadata successfully`, func() {
				object := "testString"
				model, err := visualRecognitionService.NewUpdateObjectMetadata(object)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdateObjectMetadataOptions successfully`, func() {
				// Construct an instance of the UpdateObjectMetadataOptions model
				collectionID := "testString"
				object := "testString"
				updateObjectMetadataOptionsNewObject := "testString"
				updateObjectMetadataOptionsModel := visualRecognitionService.NewUpdateObjectMetadataOptions(collectionID, object, updateObjectMetadataOptionsNewObject)
				updateObjectMetadataOptionsModel.SetCollectionID("testString")
				updateObjectMetadataOptionsModel.SetObject("testString")
				updateObjectMetadataOptionsModel.SetNewObject("testString")
				updateObjectMetadataOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateObjectMetadataOptionsModel).ToNot(BeNil())
				Expect(updateObjectMetadataOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(updateObjectMetadataOptionsModel.Object).To(Equal(core.StringPtr("testString")))
				Expect(updateObjectMetadataOptionsModel.NewObject).To(Equal(core.StringPtr("testString")))
				Expect(updateObjectMetadataOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
