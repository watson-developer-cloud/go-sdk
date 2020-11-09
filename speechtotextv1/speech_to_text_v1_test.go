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

package speechtotextv1_test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/speechtotextv1"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

var _ = Describe(`SpeechToTextV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(speechToTextService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
			Expect(speechToTextService.Service.IsSSLDisabled()).To(BeFalse())
			speechToTextService.DisableSSLVerification()
			Expect(speechToTextService.Service.IsSSLDisabled()).To(BeTrue())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(speechToTextService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				URL: "https://speechtotextv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(speechToTextService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SPEECH_TO_TEXT_URL": "https://speechtotextv1/api",
				"SPEECH_TO_TEXT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				})
				Expect(speechToTextService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := speechToTextService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != speechToTextService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(speechToTextService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(speechToTextService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: "https://testService/api",
				})
				Expect(speechToTextService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := speechToTextService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != speechToTextService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(speechToTextService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(speechToTextService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				})
				err := speechToTextService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := speechToTextService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != speechToTextService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(speechToTextService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(speechToTextService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SPEECH_TO_TEXT_URL": "https://speechtotextv1/api",
				"SPEECH_TO_TEXT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(speechToTextService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SPEECH_TO_TEXT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(speechToTextService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = speechtotextv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListModels(listModelsOptions *ListModelsOptions) - Operation response error`, func() {
		listModelsPath := "/v1/models"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listModelsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListModels with error: Operation response processing error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the ListModelsOptions model
				listModelsOptionsModel := new(speechtotextv1.ListModelsOptions)
				listModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := speechToTextService.ListModels(listModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				speechToTextService.EnableRetries(0, 0)
				result, response, operationErr = speechToTextService.ListModels(listModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListModels(listModelsOptions *ListModelsOptions)`, func() {
		listModelsPath := "/v1/models"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listModelsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"models": [{"name": "Name", "language": "Language", "rate": 4, "url": "URL", "supported_features": {"custom_language_model": false, "speaker_labels": false}, "description": "Description"}]}`)
				}))
			})
			It(`Invoke ListModels successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := speechToTextService.ListModels(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListModelsOptions model
				listModelsOptionsModel := new(speechtotextv1.ListModelsOptions)
				listModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = speechToTextService.ListModels(listModelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.ListModelsWithContext(ctx, listModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				speechToTextService.DisableRetries()
				result, response, operationErr = speechToTextService.ListModels(listModelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.ListModelsWithContext(ctx, listModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListModels with error: Operation request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the ListModelsOptions model
				listModelsOptionsModel := new(speechtotextv1.ListModelsOptions)
				listModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := speechToTextService.ListModels(listModelsOptionsModel)
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
	Describe(`GetModel(getModelOptions *GetModelOptions) - Operation response error`, func() {
		getModelPath := "/v1/models/ar-AR_BroadbandModel"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getModelPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetModel with error: Operation response processing error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the GetModelOptions model
				getModelOptionsModel := new(speechtotextv1.GetModelOptions)
				getModelOptionsModel.ModelID = core.StringPtr("ar-AR_BroadbandModel")
				getModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := speechToTextService.GetModel(getModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				speechToTextService.EnableRetries(0, 0)
				result, response, operationErr = speechToTextService.GetModel(getModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetModel(getModelOptions *GetModelOptions)`, func() {
		getModelPath := "/v1/models/ar-AR_BroadbandModel"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getModelPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "language": "Language", "rate": 4, "url": "URL", "supported_features": {"custom_language_model": false, "speaker_labels": false}, "description": "Description"}`)
				}))
			})
			It(`Invoke GetModel successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := speechToTextService.GetModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetModelOptions model
				getModelOptionsModel := new(speechtotextv1.GetModelOptions)
				getModelOptionsModel.ModelID = core.StringPtr("ar-AR_BroadbandModel")
				getModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = speechToTextService.GetModel(getModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.GetModelWithContext(ctx, getModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				speechToTextService.DisableRetries()
				result, response, operationErr = speechToTextService.GetModel(getModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.GetModelWithContext(ctx, getModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetModel with error: Operation validation and request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the GetModelOptions model
				getModelOptionsModel := new(speechtotextv1.GetModelOptions)
				getModelOptionsModel.ModelID = core.StringPtr("ar-AR_BroadbandModel")
				getModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := speechToTextService.GetModel(getModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetModelOptions model with no property values
				getModelOptionsModelNew := new(speechtotextv1.GetModelOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = speechToTextService.GetModel(getModelOptionsModelNew)
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
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(speechToTextService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
			Expect(speechToTextService.Service.IsSSLDisabled()).To(BeFalse())
			speechToTextService.DisableSSLVerification()
			Expect(speechToTextService.Service.IsSSLDisabled()).To(BeTrue())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(speechToTextService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				URL: "https://speechtotextv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(speechToTextService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SPEECH_TO_TEXT_URL": "https://speechtotextv1/api",
				"SPEECH_TO_TEXT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				})
				Expect(speechToTextService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := speechToTextService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != speechToTextService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(speechToTextService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(speechToTextService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: "https://testService/api",
				})
				Expect(speechToTextService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := speechToTextService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != speechToTextService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(speechToTextService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(speechToTextService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				})
				err := speechToTextService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := speechToTextService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != speechToTextService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(speechToTextService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(speechToTextService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SPEECH_TO_TEXT_URL": "https://speechtotextv1/api",
				"SPEECH_TO_TEXT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(speechToTextService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SPEECH_TO_TEXT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(speechToTextService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = speechtotextv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`Recognize(recognizeOptions *RecognizeOptions) - Operation response error`, func() {
		recognizePath := "/v1/recognize"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(recognizePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Content-Type"]).ToNot(BeNil())
					Expect(req.Header["Content-Type"][0]).To(Equal(fmt.Sprintf("%v", "application/octet-stream")))
					Expect(req.URL.Query()["model"]).To(Equal([]string{"ar-AR_BroadbandModel"}))

					Expect(req.URL.Query()["language_customization_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["acoustic_customization_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["base_model_version"]).To(Equal([]string{"testString"}))


					// TODO: Add check for customization_weight query parameter

					Expect(req.URL.Query()["inactivity_timeout"]).To(Equal([]string{fmt.Sprint(int64(38))}))


					// TODO: Add check for keywords_threshold query parameter

					Expect(req.URL.Query()["max_alternatives"]).To(Equal([]string{fmt.Sprint(int64(38))}))


					// TODO: Add check for word_alternatives_threshold query parameter


					// TODO: Add check for word_confidence query parameter


					// TODO: Add check for timestamps query parameter


					// TODO: Add check for profanity_filter query parameter


					// TODO: Add check for smart_formatting query parameter


					// TODO: Add check for speaker_labels query parameter

					Expect(req.URL.Query()["customization_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["grammar_name"]).To(Equal([]string{"testString"}))


					// TODO: Add check for redaction query parameter


					// TODO: Add check for audio_metrics query parameter


					// TODO: Add check for end_of_phrase_silence_time query parameter


					// TODO: Add check for split_transcript_at_phrase_end query parameter


					// TODO: Add check for speech_detector_sensitivity query parameter


					// TODO: Add check for background_audio_suppression query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke Recognize with error: Operation response processing error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the RecognizeOptions model
				recognizeOptionsModel := new(speechtotextv1.RecognizeOptions)
				recognizeOptionsModel.Audio = CreateMockReader("This is a mock file.")
				recognizeOptionsModel.ContentType = core.StringPtr("application/octet-stream")
				recognizeOptionsModel.Model = core.StringPtr("ar-AR_BroadbandModel")
				recognizeOptionsModel.LanguageCustomizationID = core.StringPtr("testString")
				recognizeOptionsModel.AcousticCustomizationID = core.StringPtr("testString")
				recognizeOptionsModel.BaseModelVersion = core.StringPtr("testString")
				recognizeOptionsModel.CustomizationWeight = core.Float64Ptr(float64(72.5))
				recognizeOptionsModel.InactivityTimeout = core.Int64Ptr(int64(38))
				recognizeOptionsModel.Keywords = []string{"testString"}
				recognizeOptionsModel.KeywordsThreshold = core.Float32Ptr(float32(36.0))
				recognizeOptionsModel.MaxAlternatives = core.Int64Ptr(int64(38))
				recognizeOptionsModel.WordAlternativesThreshold = core.Float32Ptr(float32(36.0))
				recognizeOptionsModel.WordConfidence = core.BoolPtr(true)
				recognizeOptionsModel.Timestamps = core.BoolPtr(true)
				recognizeOptionsModel.ProfanityFilter = core.BoolPtr(true)
				recognizeOptionsModel.SmartFormatting = core.BoolPtr(true)
				recognizeOptionsModel.SpeakerLabels = core.BoolPtr(true)
				recognizeOptionsModel.CustomizationID = core.StringPtr("testString")
				recognizeOptionsModel.GrammarName = core.StringPtr("testString")
				recognizeOptionsModel.Redaction = core.BoolPtr(true)
				recognizeOptionsModel.AudioMetrics = core.BoolPtr(true)
				recognizeOptionsModel.EndOfPhraseSilenceTime = core.Float64Ptr(float64(72.5))
				recognizeOptionsModel.SplitTranscriptAtPhraseEnd = core.BoolPtr(true)
				recognizeOptionsModel.SpeechDetectorSensitivity = core.Float32Ptr(float32(36.0))
				recognizeOptionsModel.BackgroundAudioSuppression = core.Float32Ptr(float32(36.0))
				recognizeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := speechToTextService.Recognize(recognizeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				speechToTextService.EnableRetries(0, 0)
				result, response, operationErr = speechToTextService.Recognize(recognizeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`Recognize(recognizeOptions *RecognizeOptions)`, func() {
		recognizePath := "/v1/recognize"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(recognizePath))
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

					Expect(req.Header["Content-Type"]).ToNot(BeNil())
					Expect(req.Header["Content-Type"][0]).To(Equal(fmt.Sprintf("%v", "application/octet-stream")))
					Expect(req.URL.Query()["model"]).To(Equal([]string{"ar-AR_BroadbandModel"}))

					Expect(req.URL.Query()["language_customization_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["acoustic_customization_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["base_model_version"]).To(Equal([]string{"testString"}))


					// TODO: Add check for customization_weight query parameter

					Expect(req.URL.Query()["inactivity_timeout"]).To(Equal([]string{fmt.Sprint(int64(38))}))


					// TODO: Add check for keywords_threshold query parameter

					Expect(req.URL.Query()["max_alternatives"]).To(Equal([]string{fmt.Sprint(int64(38))}))


					// TODO: Add check for word_alternatives_threshold query parameter


					// TODO: Add check for word_confidence query parameter


					// TODO: Add check for timestamps query parameter


					// TODO: Add check for profanity_filter query parameter


					// TODO: Add check for smart_formatting query parameter


					// TODO: Add check for speaker_labels query parameter

					Expect(req.URL.Query()["customization_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["grammar_name"]).To(Equal([]string{"testString"}))


					// TODO: Add check for redaction query parameter


					// TODO: Add check for audio_metrics query parameter


					// TODO: Add check for end_of_phrase_silence_time query parameter


					// TODO: Add check for split_transcript_at_phrase_end query parameter


					// TODO: Add check for speech_detector_sensitivity query parameter


					// TODO: Add check for background_audio_suppression query parameter

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"results": [{"final": false, "alternatives": [{"transcript": "Transcript", "confidence": 0, "timestamps": ["Timestamps"], "word_confidence": ["WordConfidence"]}], "keywords_result": {"mapKey": [{"normalized_text": "NormalizedText", "start_time": 9, "end_time": 7, "confidence": 0}]}, "word_alternatives": [{"start_time": 9, "end_time": 7, "alternatives": [{"confidence": 0, "word": "Word"}]}], "end_of_utterance": "end_of_data"}], "result_index": 11, "speaker_labels": [{"from": 4, "to": 2, "speaker": 7, "confidence": 10, "final": false}], "processing_metrics": {"processed_audio": {"received": 8, "seen_by_engine": 12, "transcription": 13, "speaker_labels": 13}, "wall_clock_since_first_byte_received": 31, "periodic": true}, "audio_metrics": {"sampling_interval": 16, "accumulated": {"final": false, "end_time": 7, "signal_to_noise_ratio": 18, "speech_ratio": 11, "high_frequency_loss": 17, "direct_current_offset": [{"begin": 5, "end": 3, "count": 5}], "clipping_rate": [{"begin": 5, "end": 3, "count": 5}], "speech_level": [{"begin": 5, "end": 3, "count": 5}], "non_speech_level": [{"begin": 5, "end": 3, "count": 5}]}}, "warnings": ["Warnings"]}`)
				}))
			})
			It(`Invoke Recognize successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := speechToTextService.Recognize(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RecognizeOptions model
				recognizeOptionsModel := new(speechtotextv1.RecognizeOptions)
				recognizeOptionsModel.Audio = CreateMockReader("This is a mock file.")
				recognizeOptionsModel.ContentType = core.StringPtr("application/octet-stream")
				recognizeOptionsModel.Model = core.StringPtr("ar-AR_BroadbandModel")
				recognizeOptionsModel.LanguageCustomizationID = core.StringPtr("testString")
				recognizeOptionsModel.AcousticCustomizationID = core.StringPtr("testString")
				recognizeOptionsModel.BaseModelVersion = core.StringPtr("testString")
				recognizeOptionsModel.CustomizationWeight = core.Float64Ptr(float64(72.5))
				recognizeOptionsModel.InactivityTimeout = core.Int64Ptr(int64(38))
				recognizeOptionsModel.Keywords = []string{"testString"}
				recognizeOptionsModel.KeywordsThreshold = core.Float32Ptr(float32(36.0))
				recognizeOptionsModel.MaxAlternatives = core.Int64Ptr(int64(38))
				recognizeOptionsModel.WordAlternativesThreshold = core.Float32Ptr(float32(36.0))
				recognizeOptionsModel.WordConfidence = core.BoolPtr(true)
				recognizeOptionsModel.Timestamps = core.BoolPtr(true)
				recognizeOptionsModel.ProfanityFilter = core.BoolPtr(true)
				recognizeOptionsModel.SmartFormatting = core.BoolPtr(true)
				recognizeOptionsModel.SpeakerLabels = core.BoolPtr(true)
				recognizeOptionsModel.CustomizationID = core.StringPtr("testString")
				recognizeOptionsModel.GrammarName = core.StringPtr("testString")
				recognizeOptionsModel.Redaction = core.BoolPtr(true)
				recognizeOptionsModel.AudioMetrics = core.BoolPtr(true)
				recognizeOptionsModel.EndOfPhraseSilenceTime = core.Float64Ptr(float64(72.5))
				recognizeOptionsModel.SplitTranscriptAtPhraseEnd = core.BoolPtr(true)
				recognizeOptionsModel.SpeechDetectorSensitivity = core.Float32Ptr(float32(36.0))
				recognizeOptionsModel.BackgroundAudioSuppression = core.Float32Ptr(float32(36.0))
				recognizeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = speechToTextService.Recognize(recognizeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.RecognizeWithContext(ctx, recognizeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				speechToTextService.DisableRetries()
				result, response, operationErr = speechToTextService.Recognize(recognizeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.RecognizeWithContext(ctx, recognizeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke Recognize with error: Operation validation and request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the RecognizeOptions model
				recognizeOptionsModel := new(speechtotextv1.RecognizeOptions)
				recognizeOptionsModel.Audio = CreateMockReader("This is a mock file.")
				recognizeOptionsModel.ContentType = core.StringPtr("application/octet-stream")
				recognizeOptionsModel.Model = core.StringPtr("ar-AR_BroadbandModel")
				recognizeOptionsModel.LanguageCustomizationID = core.StringPtr("testString")
				recognizeOptionsModel.AcousticCustomizationID = core.StringPtr("testString")
				recognizeOptionsModel.BaseModelVersion = core.StringPtr("testString")
				recognizeOptionsModel.CustomizationWeight = core.Float64Ptr(float64(72.5))
				recognizeOptionsModel.InactivityTimeout = core.Int64Ptr(int64(38))
				recognizeOptionsModel.Keywords = []string{"testString"}
				recognizeOptionsModel.KeywordsThreshold = core.Float32Ptr(float32(36.0))
				recognizeOptionsModel.MaxAlternatives = core.Int64Ptr(int64(38))
				recognizeOptionsModel.WordAlternativesThreshold = core.Float32Ptr(float32(36.0))
				recognizeOptionsModel.WordConfidence = core.BoolPtr(true)
				recognizeOptionsModel.Timestamps = core.BoolPtr(true)
				recognizeOptionsModel.ProfanityFilter = core.BoolPtr(true)
				recognizeOptionsModel.SmartFormatting = core.BoolPtr(true)
				recognizeOptionsModel.SpeakerLabels = core.BoolPtr(true)
				recognizeOptionsModel.CustomizationID = core.StringPtr("testString")
				recognizeOptionsModel.GrammarName = core.StringPtr("testString")
				recognizeOptionsModel.Redaction = core.BoolPtr(true)
				recognizeOptionsModel.AudioMetrics = core.BoolPtr(true)
				recognizeOptionsModel.EndOfPhraseSilenceTime = core.Float64Ptr(float64(72.5))
				recognizeOptionsModel.SplitTranscriptAtPhraseEnd = core.BoolPtr(true)
				recognizeOptionsModel.SpeechDetectorSensitivity = core.Float32Ptr(float32(36.0))
				recognizeOptionsModel.BackgroundAudioSuppression = core.Float32Ptr(float32(36.0))
				recognizeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := speechToTextService.Recognize(recognizeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the RecognizeOptions model with no property values
				recognizeOptionsModelNew := new(speechtotextv1.RecognizeOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = speechToTextService.Recognize(recognizeOptionsModelNew)
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
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(speechToTextService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
			Expect(speechToTextService.Service.IsSSLDisabled()).To(BeFalse())
			speechToTextService.DisableSSLVerification()
			Expect(speechToTextService.Service.IsSSLDisabled()).To(BeTrue())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(speechToTextService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				URL: "https://speechtotextv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(speechToTextService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SPEECH_TO_TEXT_URL": "https://speechtotextv1/api",
				"SPEECH_TO_TEXT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				})
				Expect(speechToTextService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := speechToTextService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != speechToTextService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(speechToTextService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(speechToTextService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: "https://testService/api",
				})
				Expect(speechToTextService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := speechToTextService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != speechToTextService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(speechToTextService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(speechToTextService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				})
				err := speechToTextService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := speechToTextService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != speechToTextService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(speechToTextService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(speechToTextService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SPEECH_TO_TEXT_URL": "https://speechtotextv1/api",
				"SPEECH_TO_TEXT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(speechToTextService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SPEECH_TO_TEXT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(speechToTextService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = speechtotextv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`RegisterCallback(registerCallbackOptions *RegisterCallbackOptions) - Operation response error`, func() {
		registerCallbackPath := "/v1/register_callback"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(registerCallbackPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["callback_url"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["user_secret"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke RegisterCallback with error: Operation response processing error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the RegisterCallbackOptions model
				registerCallbackOptionsModel := new(speechtotextv1.RegisterCallbackOptions)
				registerCallbackOptionsModel.CallbackURL = core.StringPtr("testString")
				registerCallbackOptionsModel.UserSecret = core.StringPtr("testString")
				registerCallbackOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := speechToTextService.RegisterCallback(registerCallbackOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				speechToTextService.EnableRetries(0, 0)
				result, response, operationErr = speechToTextService.RegisterCallback(registerCallbackOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`RegisterCallback(registerCallbackOptions *RegisterCallbackOptions)`, func() {
		registerCallbackPath := "/v1/register_callback"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(registerCallbackPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["callback_url"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["user_secret"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status": "created", "url": "URL"}`)
				}))
			})
			It(`Invoke RegisterCallback successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := speechToTextService.RegisterCallback(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RegisterCallbackOptions model
				registerCallbackOptionsModel := new(speechtotextv1.RegisterCallbackOptions)
				registerCallbackOptionsModel.CallbackURL = core.StringPtr("testString")
				registerCallbackOptionsModel.UserSecret = core.StringPtr("testString")
				registerCallbackOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = speechToTextService.RegisterCallback(registerCallbackOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.RegisterCallbackWithContext(ctx, registerCallbackOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				speechToTextService.DisableRetries()
				result, response, operationErr = speechToTextService.RegisterCallback(registerCallbackOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.RegisterCallbackWithContext(ctx, registerCallbackOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke RegisterCallback with error: Operation validation and request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the RegisterCallbackOptions model
				registerCallbackOptionsModel := new(speechtotextv1.RegisterCallbackOptions)
				registerCallbackOptionsModel.CallbackURL = core.StringPtr("testString")
				registerCallbackOptionsModel.UserSecret = core.StringPtr("testString")
				registerCallbackOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := speechToTextService.RegisterCallback(registerCallbackOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the RegisterCallbackOptions model with no property values
				registerCallbackOptionsModelNew := new(speechtotextv1.RegisterCallbackOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = speechToTextService.RegisterCallback(registerCallbackOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UnregisterCallback(unregisterCallbackOptions *UnregisterCallbackOptions)`, func() {
		unregisterCallbackPath := "/v1/unregister_callback"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(unregisterCallbackPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["callback_url"]).To(Equal([]string{"testString"}))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke UnregisterCallback successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := speechToTextService.UnregisterCallback(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UnregisterCallbackOptions model
				unregisterCallbackOptionsModel := new(speechtotextv1.UnregisterCallbackOptions)
				unregisterCallbackOptionsModel.CallbackURL = core.StringPtr("testString")
				unregisterCallbackOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = speechToTextService.UnregisterCallback(unregisterCallbackOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				speechToTextService.DisableRetries()
				response, operationErr = speechToTextService.UnregisterCallback(unregisterCallbackOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UnregisterCallback with error: Operation validation and request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the UnregisterCallbackOptions model
				unregisterCallbackOptionsModel := new(speechtotextv1.UnregisterCallbackOptions)
				unregisterCallbackOptionsModel.CallbackURL = core.StringPtr("testString")
				unregisterCallbackOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := speechToTextService.UnregisterCallback(unregisterCallbackOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UnregisterCallbackOptions model with no property values
				unregisterCallbackOptionsModelNew := new(speechtotextv1.UnregisterCallbackOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = speechToTextService.UnregisterCallback(unregisterCallbackOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateJob(createJobOptions *CreateJobOptions) - Operation response error`, func() {
		createJobPath := "/v1/recognitions"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createJobPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Content-Type"]).ToNot(BeNil())
					Expect(req.Header["Content-Type"][0]).To(Equal(fmt.Sprintf("%v", "application/octet-stream")))
					Expect(req.URL.Query()["model"]).To(Equal([]string{"ar-AR_BroadbandModel"}))

					Expect(req.URL.Query()["callback_url"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["events"]).To(Equal([]string{"recognitions.started"}))

					Expect(req.URL.Query()["user_token"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["results_ttl"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["language_customization_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["acoustic_customization_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["base_model_version"]).To(Equal([]string{"testString"}))


					// TODO: Add check for customization_weight query parameter

					Expect(req.URL.Query()["inactivity_timeout"]).To(Equal([]string{fmt.Sprint(int64(38))}))


					// TODO: Add check for keywords_threshold query parameter

					Expect(req.URL.Query()["max_alternatives"]).To(Equal([]string{fmt.Sprint(int64(38))}))


					// TODO: Add check for word_alternatives_threshold query parameter


					// TODO: Add check for word_confidence query parameter


					// TODO: Add check for timestamps query parameter


					// TODO: Add check for profanity_filter query parameter


					// TODO: Add check for smart_formatting query parameter


					// TODO: Add check for speaker_labels query parameter

					Expect(req.URL.Query()["customization_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["grammar_name"]).To(Equal([]string{"testString"}))


					// TODO: Add check for redaction query parameter


					// TODO: Add check for processing_metrics query parameter


					// TODO: Add check for processing_metrics_interval query parameter


					// TODO: Add check for audio_metrics query parameter


					// TODO: Add check for end_of_phrase_silence_time query parameter


					// TODO: Add check for split_transcript_at_phrase_end query parameter


					// TODO: Add check for speech_detector_sensitivity query parameter


					// TODO: Add check for background_audio_suppression query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateJob with error: Operation response processing error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the CreateJobOptions model
				createJobOptionsModel := new(speechtotextv1.CreateJobOptions)
				createJobOptionsModel.Audio = CreateMockReader("This is a mock file.")
				createJobOptionsModel.ContentType = core.StringPtr("application/octet-stream")
				createJobOptionsModel.Model = core.StringPtr("ar-AR_BroadbandModel")
				createJobOptionsModel.CallbackURL = core.StringPtr("testString")
				createJobOptionsModel.Events = core.StringPtr("recognitions.started")
				createJobOptionsModel.UserToken = core.StringPtr("testString")
				createJobOptionsModel.ResultsTTL = core.Int64Ptr(int64(38))
				createJobOptionsModel.LanguageCustomizationID = core.StringPtr("testString")
				createJobOptionsModel.AcousticCustomizationID = core.StringPtr("testString")
				createJobOptionsModel.BaseModelVersion = core.StringPtr("testString")
				createJobOptionsModel.CustomizationWeight = core.Float64Ptr(float64(72.5))
				createJobOptionsModel.InactivityTimeout = core.Int64Ptr(int64(38))
				createJobOptionsModel.Keywords = []string{"testString"}
				createJobOptionsModel.KeywordsThreshold = core.Float32Ptr(float32(36.0))
				createJobOptionsModel.MaxAlternatives = core.Int64Ptr(int64(38))
				createJobOptionsModel.WordAlternativesThreshold = core.Float32Ptr(float32(36.0))
				createJobOptionsModel.WordConfidence = core.BoolPtr(true)
				createJobOptionsModel.Timestamps = core.BoolPtr(true)
				createJobOptionsModel.ProfanityFilter = core.BoolPtr(true)
				createJobOptionsModel.SmartFormatting = core.BoolPtr(true)
				createJobOptionsModel.SpeakerLabels = core.BoolPtr(true)
				createJobOptionsModel.CustomizationID = core.StringPtr("testString")
				createJobOptionsModel.GrammarName = core.StringPtr("testString")
				createJobOptionsModel.Redaction = core.BoolPtr(true)
				createJobOptionsModel.ProcessingMetrics = core.BoolPtr(true)
				createJobOptionsModel.ProcessingMetricsInterval = core.Float32Ptr(float32(36.0))
				createJobOptionsModel.AudioMetrics = core.BoolPtr(true)
				createJobOptionsModel.EndOfPhraseSilenceTime = core.Float64Ptr(float64(72.5))
				createJobOptionsModel.SplitTranscriptAtPhraseEnd = core.BoolPtr(true)
				createJobOptionsModel.SpeechDetectorSensitivity = core.Float32Ptr(float32(36.0))
				createJobOptionsModel.BackgroundAudioSuppression = core.Float32Ptr(float32(36.0))
				createJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := speechToTextService.CreateJob(createJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				speechToTextService.EnableRetries(0, 0)
				result, response, operationErr = speechToTextService.CreateJob(createJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateJob(createJobOptions *CreateJobOptions)`, func() {
		createJobPath := "/v1/recognitions"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createJobPath))
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

					Expect(req.Header["Content-Type"]).ToNot(BeNil())
					Expect(req.Header["Content-Type"][0]).To(Equal(fmt.Sprintf("%v", "application/octet-stream")))
					Expect(req.URL.Query()["model"]).To(Equal([]string{"ar-AR_BroadbandModel"}))

					Expect(req.URL.Query()["callback_url"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["events"]).To(Equal([]string{"recognitions.started"}))

					Expect(req.URL.Query()["user_token"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["results_ttl"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["language_customization_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["acoustic_customization_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["base_model_version"]).To(Equal([]string{"testString"}))


					// TODO: Add check for customization_weight query parameter

					Expect(req.URL.Query()["inactivity_timeout"]).To(Equal([]string{fmt.Sprint(int64(38))}))


					// TODO: Add check for keywords_threshold query parameter

					Expect(req.URL.Query()["max_alternatives"]).To(Equal([]string{fmt.Sprint(int64(38))}))


					// TODO: Add check for word_alternatives_threshold query parameter


					// TODO: Add check for word_confidence query parameter


					// TODO: Add check for timestamps query parameter


					// TODO: Add check for profanity_filter query parameter


					// TODO: Add check for smart_formatting query parameter


					// TODO: Add check for speaker_labels query parameter

					Expect(req.URL.Query()["customization_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["grammar_name"]).To(Equal([]string{"testString"}))


					// TODO: Add check for redaction query parameter


					// TODO: Add check for processing_metrics query parameter


					// TODO: Add check for processing_metrics_interval query parameter


					// TODO: Add check for audio_metrics query parameter


					// TODO: Add check for end_of_phrase_silence_time query parameter


					// TODO: Add check for split_transcript_at_phrase_end query parameter


					// TODO: Add check for speech_detector_sensitivity query parameter


					// TODO: Add check for background_audio_suppression query parameter

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "status": "waiting", "created": "Created", "updated": "Updated", "url": "URL", "user_token": "UserToken", "results": [{"results": [{"final": false, "alternatives": [{"transcript": "Transcript", "confidence": 0, "timestamps": ["Timestamps"], "word_confidence": ["WordConfidence"]}], "keywords_result": {"mapKey": [{"normalized_text": "NormalizedText", "start_time": 9, "end_time": 7, "confidence": 0}]}, "word_alternatives": [{"start_time": 9, "end_time": 7, "alternatives": [{"confidence": 0, "word": "Word"}]}], "end_of_utterance": "end_of_data"}], "result_index": 11, "speaker_labels": [{"from": 4, "to": 2, "speaker": 7, "confidence": 10, "final": false}], "processing_metrics": {"processed_audio": {"received": 8, "seen_by_engine": 12, "transcription": 13, "speaker_labels": 13}, "wall_clock_since_first_byte_received": 31, "periodic": true}, "audio_metrics": {"sampling_interval": 16, "accumulated": {"final": false, "end_time": 7, "signal_to_noise_ratio": 18, "speech_ratio": 11, "high_frequency_loss": 17, "direct_current_offset": [{"begin": 5, "end": 3, "count": 5}], "clipping_rate": [{"begin": 5, "end": 3, "count": 5}], "speech_level": [{"begin": 5, "end": 3, "count": 5}], "non_speech_level": [{"begin": 5, "end": 3, "count": 5}]}}, "warnings": ["Warnings"]}], "warnings": ["Warnings"]}`)
				}))
			})
			It(`Invoke CreateJob successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := speechToTextService.CreateJob(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateJobOptions model
				createJobOptionsModel := new(speechtotextv1.CreateJobOptions)
				createJobOptionsModel.Audio = CreateMockReader("This is a mock file.")
				createJobOptionsModel.ContentType = core.StringPtr("application/octet-stream")
				createJobOptionsModel.Model = core.StringPtr("ar-AR_BroadbandModel")
				createJobOptionsModel.CallbackURL = core.StringPtr("testString")
				createJobOptionsModel.Events = core.StringPtr("recognitions.started")
				createJobOptionsModel.UserToken = core.StringPtr("testString")
				createJobOptionsModel.ResultsTTL = core.Int64Ptr(int64(38))
				createJobOptionsModel.LanguageCustomizationID = core.StringPtr("testString")
				createJobOptionsModel.AcousticCustomizationID = core.StringPtr("testString")
				createJobOptionsModel.BaseModelVersion = core.StringPtr("testString")
				createJobOptionsModel.CustomizationWeight = core.Float64Ptr(float64(72.5))
				createJobOptionsModel.InactivityTimeout = core.Int64Ptr(int64(38))
				createJobOptionsModel.Keywords = []string{"testString"}
				createJobOptionsModel.KeywordsThreshold = core.Float32Ptr(float32(36.0))
				createJobOptionsModel.MaxAlternatives = core.Int64Ptr(int64(38))
				createJobOptionsModel.WordAlternativesThreshold = core.Float32Ptr(float32(36.0))
				createJobOptionsModel.WordConfidence = core.BoolPtr(true)
				createJobOptionsModel.Timestamps = core.BoolPtr(true)
				createJobOptionsModel.ProfanityFilter = core.BoolPtr(true)
				createJobOptionsModel.SmartFormatting = core.BoolPtr(true)
				createJobOptionsModel.SpeakerLabels = core.BoolPtr(true)
				createJobOptionsModel.CustomizationID = core.StringPtr("testString")
				createJobOptionsModel.GrammarName = core.StringPtr("testString")
				createJobOptionsModel.Redaction = core.BoolPtr(true)
				createJobOptionsModel.ProcessingMetrics = core.BoolPtr(true)
				createJobOptionsModel.ProcessingMetricsInterval = core.Float32Ptr(float32(36.0))
				createJobOptionsModel.AudioMetrics = core.BoolPtr(true)
				createJobOptionsModel.EndOfPhraseSilenceTime = core.Float64Ptr(float64(72.5))
				createJobOptionsModel.SplitTranscriptAtPhraseEnd = core.BoolPtr(true)
				createJobOptionsModel.SpeechDetectorSensitivity = core.Float32Ptr(float32(36.0))
				createJobOptionsModel.BackgroundAudioSuppression = core.Float32Ptr(float32(36.0))
				createJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = speechToTextService.CreateJob(createJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.CreateJobWithContext(ctx, createJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				speechToTextService.DisableRetries()
				result, response, operationErr = speechToTextService.CreateJob(createJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.CreateJobWithContext(ctx, createJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CreateJob with error: Operation validation and request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the CreateJobOptions model
				createJobOptionsModel := new(speechtotextv1.CreateJobOptions)
				createJobOptionsModel.Audio = CreateMockReader("This is a mock file.")
				createJobOptionsModel.ContentType = core.StringPtr("application/octet-stream")
				createJobOptionsModel.Model = core.StringPtr("ar-AR_BroadbandModel")
				createJobOptionsModel.CallbackURL = core.StringPtr("testString")
				createJobOptionsModel.Events = core.StringPtr("recognitions.started")
				createJobOptionsModel.UserToken = core.StringPtr("testString")
				createJobOptionsModel.ResultsTTL = core.Int64Ptr(int64(38))
				createJobOptionsModel.LanguageCustomizationID = core.StringPtr("testString")
				createJobOptionsModel.AcousticCustomizationID = core.StringPtr("testString")
				createJobOptionsModel.BaseModelVersion = core.StringPtr("testString")
				createJobOptionsModel.CustomizationWeight = core.Float64Ptr(float64(72.5))
				createJobOptionsModel.InactivityTimeout = core.Int64Ptr(int64(38))
				createJobOptionsModel.Keywords = []string{"testString"}
				createJobOptionsModel.KeywordsThreshold = core.Float32Ptr(float32(36.0))
				createJobOptionsModel.MaxAlternatives = core.Int64Ptr(int64(38))
				createJobOptionsModel.WordAlternativesThreshold = core.Float32Ptr(float32(36.0))
				createJobOptionsModel.WordConfidence = core.BoolPtr(true)
				createJobOptionsModel.Timestamps = core.BoolPtr(true)
				createJobOptionsModel.ProfanityFilter = core.BoolPtr(true)
				createJobOptionsModel.SmartFormatting = core.BoolPtr(true)
				createJobOptionsModel.SpeakerLabels = core.BoolPtr(true)
				createJobOptionsModel.CustomizationID = core.StringPtr("testString")
				createJobOptionsModel.GrammarName = core.StringPtr("testString")
				createJobOptionsModel.Redaction = core.BoolPtr(true)
				createJobOptionsModel.ProcessingMetrics = core.BoolPtr(true)
				createJobOptionsModel.ProcessingMetricsInterval = core.Float32Ptr(float32(36.0))
				createJobOptionsModel.AudioMetrics = core.BoolPtr(true)
				createJobOptionsModel.EndOfPhraseSilenceTime = core.Float64Ptr(float64(72.5))
				createJobOptionsModel.SplitTranscriptAtPhraseEnd = core.BoolPtr(true)
				createJobOptionsModel.SpeechDetectorSensitivity = core.Float32Ptr(float32(36.0))
				createJobOptionsModel.BackgroundAudioSuppression = core.Float32Ptr(float32(36.0))
				createJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := speechToTextService.CreateJob(createJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateJobOptions model with no property values
				createJobOptionsModelNew := new(speechtotextv1.CreateJobOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = speechToTextService.CreateJob(createJobOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CheckJobs(checkJobsOptions *CheckJobsOptions) - Operation response error`, func() {
		checkJobsPath := "/v1/recognitions"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(checkJobsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CheckJobs with error: Operation response processing error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the CheckJobsOptions model
				checkJobsOptionsModel := new(speechtotextv1.CheckJobsOptions)
				checkJobsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := speechToTextService.CheckJobs(checkJobsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				speechToTextService.EnableRetries(0, 0)
				result, response, operationErr = speechToTextService.CheckJobs(checkJobsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CheckJobs(checkJobsOptions *CheckJobsOptions)`, func() {
		checkJobsPath := "/v1/recognitions"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(checkJobsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"recognitions": [{"id": "ID", "status": "waiting", "created": "Created", "updated": "Updated", "url": "URL", "user_token": "UserToken", "results": [{"results": [{"final": false, "alternatives": [{"transcript": "Transcript", "confidence": 0, "timestamps": ["Timestamps"], "word_confidence": ["WordConfidence"]}], "keywords_result": {"mapKey": [{"normalized_text": "NormalizedText", "start_time": 9, "end_time": 7, "confidence": 0}]}, "word_alternatives": [{"start_time": 9, "end_time": 7, "alternatives": [{"confidence": 0, "word": "Word"}]}], "end_of_utterance": "end_of_data"}], "result_index": 11, "speaker_labels": [{"from": 4, "to": 2, "speaker": 7, "confidence": 10, "final": false}], "processing_metrics": {"processed_audio": {"received": 8, "seen_by_engine": 12, "transcription": 13, "speaker_labels": 13}, "wall_clock_since_first_byte_received": 31, "periodic": true}, "audio_metrics": {"sampling_interval": 16, "accumulated": {"final": false, "end_time": 7, "signal_to_noise_ratio": 18, "speech_ratio": 11, "high_frequency_loss": 17, "direct_current_offset": [{"begin": 5, "end": 3, "count": 5}], "clipping_rate": [{"begin": 5, "end": 3, "count": 5}], "speech_level": [{"begin": 5, "end": 3, "count": 5}], "non_speech_level": [{"begin": 5, "end": 3, "count": 5}]}}, "warnings": ["Warnings"]}], "warnings": ["Warnings"]}]}`)
				}))
			})
			It(`Invoke CheckJobs successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := speechToTextService.CheckJobs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CheckJobsOptions model
				checkJobsOptionsModel := new(speechtotextv1.CheckJobsOptions)
				checkJobsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = speechToTextService.CheckJobs(checkJobsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.CheckJobsWithContext(ctx, checkJobsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				speechToTextService.DisableRetries()
				result, response, operationErr = speechToTextService.CheckJobs(checkJobsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.CheckJobsWithContext(ctx, checkJobsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CheckJobs with error: Operation request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the CheckJobsOptions model
				checkJobsOptionsModel := new(speechtotextv1.CheckJobsOptions)
				checkJobsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := speechToTextService.CheckJobs(checkJobsOptionsModel)
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
	Describe(`CheckJob(checkJobOptions *CheckJobOptions) - Operation response error`, func() {
		checkJobPath := "/v1/recognitions/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(checkJobPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CheckJob with error: Operation response processing error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the CheckJobOptions model
				checkJobOptionsModel := new(speechtotextv1.CheckJobOptions)
				checkJobOptionsModel.ID = core.StringPtr("testString")
				checkJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := speechToTextService.CheckJob(checkJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				speechToTextService.EnableRetries(0, 0)
				result, response, operationErr = speechToTextService.CheckJob(checkJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CheckJob(checkJobOptions *CheckJobOptions)`, func() {
		checkJobPath := "/v1/recognitions/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(checkJobPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "status": "waiting", "created": "Created", "updated": "Updated", "url": "URL", "user_token": "UserToken", "results": [{"results": [{"final": false, "alternatives": [{"transcript": "Transcript", "confidence": 0, "timestamps": ["Timestamps"], "word_confidence": ["WordConfidence"]}], "keywords_result": {"mapKey": [{"normalized_text": "NormalizedText", "start_time": 9, "end_time": 7, "confidence": 0}]}, "word_alternatives": [{"start_time": 9, "end_time": 7, "alternatives": [{"confidence": 0, "word": "Word"}]}], "end_of_utterance": "end_of_data"}], "result_index": 11, "speaker_labels": [{"from": 4, "to": 2, "speaker": 7, "confidence": 10, "final": false}], "processing_metrics": {"processed_audio": {"received": 8, "seen_by_engine": 12, "transcription": 13, "speaker_labels": 13}, "wall_clock_since_first_byte_received": 31, "periodic": true}, "audio_metrics": {"sampling_interval": 16, "accumulated": {"final": false, "end_time": 7, "signal_to_noise_ratio": 18, "speech_ratio": 11, "high_frequency_loss": 17, "direct_current_offset": [{"begin": 5, "end": 3, "count": 5}], "clipping_rate": [{"begin": 5, "end": 3, "count": 5}], "speech_level": [{"begin": 5, "end": 3, "count": 5}], "non_speech_level": [{"begin": 5, "end": 3, "count": 5}]}}, "warnings": ["Warnings"]}], "warnings": ["Warnings"]}`)
				}))
			})
			It(`Invoke CheckJob successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := speechToTextService.CheckJob(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CheckJobOptions model
				checkJobOptionsModel := new(speechtotextv1.CheckJobOptions)
				checkJobOptionsModel.ID = core.StringPtr("testString")
				checkJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = speechToTextService.CheckJob(checkJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.CheckJobWithContext(ctx, checkJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				speechToTextService.DisableRetries()
				result, response, operationErr = speechToTextService.CheckJob(checkJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.CheckJobWithContext(ctx, checkJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CheckJob with error: Operation validation and request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the CheckJobOptions model
				checkJobOptionsModel := new(speechtotextv1.CheckJobOptions)
				checkJobOptionsModel.ID = core.StringPtr("testString")
				checkJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := speechToTextService.CheckJob(checkJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CheckJobOptions model with no property values
				checkJobOptionsModelNew := new(speechtotextv1.CheckJobOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = speechToTextService.CheckJob(checkJobOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteJob(deleteJobOptions *DeleteJobOptions)`, func() {
		deleteJobPath := "/v1/recognitions/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteJobPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteJob successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := speechToTextService.DeleteJob(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteJobOptions model
				deleteJobOptionsModel := new(speechtotextv1.DeleteJobOptions)
				deleteJobOptionsModel.ID = core.StringPtr("testString")
				deleteJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = speechToTextService.DeleteJob(deleteJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				speechToTextService.DisableRetries()
				response, operationErr = speechToTextService.DeleteJob(deleteJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteJob with error: Operation validation and request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the DeleteJobOptions model
				deleteJobOptionsModel := new(speechtotextv1.DeleteJobOptions)
				deleteJobOptionsModel.ID = core.StringPtr("testString")
				deleteJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := speechToTextService.DeleteJob(deleteJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteJobOptions model with no property values
				deleteJobOptionsModelNew := new(speechtotextv1.DeleteJobOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = speechToTextService.DeleteJob(deleteJobOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(speechToTextService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
			Expect(speechToTextService.Service.IsSSLDisabled()).To(BeFalse())
			speechToTextService.DisableSSLVerification()
			Expect(speechToTextService.Service.IsSSLDisabled()).To(BeTrue())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(speechToTextService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				URL: "https://speechtotextv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(speechToTextService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SPEECH_TO_TEXT_URL": "https://speechtotextv1/api",
				"SPEECH_TO_TEXT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				})
				Expect(speechToTextService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := speechToTextService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != speechToTextService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(speechToTextService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(speechToTextService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: "https://testService/api",
				})
				Expect(speechToTextService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := speechToTextService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != speechToTextService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(speechToTextService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(speechToTextService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				})
				err := speechToTextService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := speechToTextService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != speechToTextService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(speechToTextService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(speechToTextService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SPEECH_TO_TEXT_URL": "https://speechtotextv1/api",
				"SPEECH_TO_TEXT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(speechToTextService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SPEECH_TO_TEXT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(speechToTextService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = speechtotextv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`CreateLanguageModel(createLanguageModelOptions *CreateLanguageModelOptions) - Operation response error`, func() {
		createLanguageModelPath := "/v1/customizations"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createLanguageModelPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateLanguageModel with error: Operation response processing error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the CreateLanguageModelOptions model
				createLanguageModelOptionsModel := new(speechtotextv1.CreateLanguageModelOptions)
				createLanguageModelOptionsModel.Name = core.StringPtr("testString")
				createLanguageModelOptionsModel.BaseModelName = core.StringPtr("de-DE_BroadbandModel")
				createLanguageModelOptionsModel.Dialect = core.StringPtr("testString")
				createLanguageModelOptionsModel.Description = core.StringPtr("testString")
				createLanguageModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := speechToTextService.CreateLanguageModel(createLanguageModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				speechToTextService.EnableRetries(0, 0)
				result, response, operationErr = speechToTextService.CreateLanguageModel(createLanguageModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateLanguageModel(createLanguageModelOptions *CreateLanguageModelOptions)`, func() {
		createLanguageModelPath := "/v1/customizations"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createLanguageModelPath))
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
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"customization_id": "CustomizationID", "created": "Created", "updated": "Updated", "language": "Language", "dialect": "Dialect", "versions": ["Versions"], "owner": "Owner", "name": "Name", "description": "Description", "base_model_name": "BaseModelName", "status": "pending", "progress": 8, "error": "Error", "warnings": "Warnings"}`)
				}))
			})
			It(`Invoke CreateLanguageModel successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := speechToTextService.CreateLanguageModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateLanguageModelOptions model
				createLanguageModelOptionsModel := new(speechtotextv1.CreateLanguageModelOptions)
				createLanguageModelOptionsModel.Name = core.StringPtr("testString")
				createLanguageModelOptionsModel.BaseModelName = core.StringPtr("de-DE_BroadbandModel")
				createLanguageModelOptionsModel.Dialect = core.StringPtr("testString")
				createLanguageModelOptionsModel.Description = core.StringPtr("testString")
				createLanguageModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = speechToTextService.CreateLanguageModel(createLanguageModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.CreateLanguageModelWithContext(ctx, createLanguageModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				speechToTextService.DisableRetries()
				result, response, operationErr = speechToTextService.CreateLanguageModel(createLanguageModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.CreateLanguageModelWithContext(ctx, createLanguageModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CreateLanguageModel with error: Operation validation and request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the CreateLanguageModelOptions model
				createLanguageModelOptionsModel := new(speechtotextv1.CreateLanguageModelOptions)
				createLanguageModelOptionsModel.Name = core.StringPtr("testString")
				createLanguageModelOptionsModel.BaseModelName = core.StringPtr("de-DE_BroadbandModel")
				createLanguageModelOptionsModel.Dialect = core.StringPtr("testString")
				createLanguageModelOptionsModel.Description = core.StringPtr("testString")
				createLanguageModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := speechToTextService.CreateLanguageModel(createLanguageModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateLanguageModelOptions model with no property values
				createLanguageModelOptionsModelNew := new(speechtotextv1.CreateLanguageModelOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = speechToTextService.CreateLanguageModel(createLanguageModelOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListLanguageModels(listLanguageModelsOptions *ListLanguageModelsOptions) - Operation response error`, func() {
		listLanguageModelsPath := "/v1/customizations"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listLanguageModelsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["language"]).To(Equal([]string{"ar-AR"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListLanguageModels with error: Operation response processing error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the ListLanguageModelsOptions model
				listLanguageModelsOptionsModel := new(speechtotextv1.ListLanguageModelsOptions)
				listLanguageModelsOptionsModel.Language = core.StringPtr("ar-AR")
				listLanguageModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := speechToTextService.ListLanguageModels(listLanguageModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				speechToTextService.EnableRetries(0, 0)
				result, response, operationErr = speechToTextService.ListLanguageModels(listLanguageModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListLanguageModels(listLanguageModelsOptions *ListLanguageModelsOptions)`, func() {
		listLanguageModelsPath := "/v1/customizations"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listLanguageModelsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["language"]).To(Equal([]string{"ar-AR"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"customizations": [{"customization_id": "CustomizationID", "created": "Created", "updated": "Updated", "language": "Language", "dialect": "Dialect", "versions": ["Versions"], "owner": "Owner", "name": "Name", "description": "Description", "base_model_name": "BaseModelName", "status": "pending", "progress": 8, "error": "Error", "warnings": "Warnings"}]}`)
				}))
			})
			It(`Invoke ListLanguageModels successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := speechToTextService.ListLanguageModels(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListLanguageModelsOptions model
				listLanguageModelsOptionsModel := new(speechtotextv1.ListLanguageModelsOptions)
				listLanguageModelsOptionsModel.Language = core.StringPtr("ar-AR")
				listLanguageModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = speechToTextService.ListLanguageModels(listLanguageModelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.ListLanguageModelsWithContext(ctx, listLanguageModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				speechToTextService.DisableRetries()
				result, response, operationErr = speechToTextService.ListLanguageModels(listLanguageModelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.ListLanguageModelsWithContext(ctx, listLanguageModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListLanguageModels with error: Operation request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the ListLanguageModelsOptions model
				listLanguageModelsOptionsModel := new(speechtotextv1.ListLanguageModelsOptions)
				listLanguageModelsOptionsModel.Language = core.StringPtr("ar-AR")
				listLanguageModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := speechToTextService.ListLanguageModels(listLanguageModelsOptionsModel)
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
	Describe(`GetLanguageModel(getLanguageModelOptions *GetLanguageModelOptions) - Operation response error`, func() {
		getLanguageModelPath := "/v1/customizations/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLanguageModelPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLanguageModel with error: Operation response processing error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the GetLanguageModelOptions model
				getLanguageModelOptionsModel := new(speechtotextv1.GetLanguageModelOptions)
				getLanguageModelOptionsModel.CustomizationID = core.StringPtr("testString")
				getLanguageModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := speechToTextService.GetLanguageModel(getLanguageModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				speechToTextService.EnableRetries(0, 0)
				result, response, operationErr = speechToTextService.GetLanguageModel(getLanguageModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetLanguageModel(getLanguageModelOptions *GetLanguageModelOptions)`, func() {
		getLanguageModelPath := "/v1/customizations/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLanguageModelPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"customization_id": "CustomizationID", "created": "Created", "updated": "Updated", "language": "Language", "dialect": "Dialect", "versions": ["Versions"], "owner": "Owner", "name": "Name", "description": "Description", "base_model_name": "BaseModelName", "status": "pending", "progress": 8, "error": "Error", "warnings": "Warnings"}`)
				}))
			})
			It(`Invoke GetLanguageModel successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := speechToTextService.GetLanguageModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLanguageModelOptions model
				getLanguageModelOptionsModel := new(speechtotextv1.GetLanguageModelOptions)
				getLanguageModelOptionsModel.CustomizationID = core.StringPtr("testString")
				getLanguageModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = speechToTextService.GetLanguageModel(getLanguageModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.GetLanguageModelWithContext(ctx, getLanguageModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				speechToTextService.DisableRetries()
				result, response, operationErr = speechToTextService.GetLanguageModel(getLanguageModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.GetLanguageModelWithContext(ctx, getLanguageModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetLanguageModel with error: Operation validation and request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the GetLanguageModelOptions model
				getLanguageModelOptionsModel := new(speechtotextv1.GetLanguageModelOptions)
				getLanguageModelOptionsModel.CustomizationID = core.StringPtr("testString")
				getLanguageModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := speechToTextService.GetLanguageModel(getLanguageModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetLanguageModelOptions model with no property values
				getLanguageModelOptionsModelNew := new(speechtotextv1.GetLanguageModelOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = speechToTextService.GetLanguageModel(getLanguageModelOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteLanguageModel(deleteLanguageModelOptions *DeleteLanguageModelOptions)`, func() {
		deleteLanguageModelPath := "/v1/customizations/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteLanguageModelPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteLanguageModel successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := speechToTextService.DeleteLanguageModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteLanguageModelOptions model
				deleteLanguageModelOptionsModel := new(speechtotextv1.DeleteLanguageModelOptions)
				deleteLanguageModelOptionsModel.CustomizationID = core.StringPtr("testString")
				deleteLanguageModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = speechToTextService.DeleteLanguageModel(deleteLanguageModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				speechToTextService.DisableRetries()
				response, operationErr = speechToTextService.DeleteLanguageModel(deleteLanguageModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteLanguageModel with error: Operation validation and request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the DeleteLanguageModelOptions model
				deleteLanguageModelOptionsModel := new(speechtotextv1.DeleteLanguageModelOptions)
				deleteLanguageModelOptionsModel.CustomizationID = core.StringPtr("testString")
				deleteLanguageModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := speechToTextService.DeleteLanguageModel(deleteLanguageModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteLanguageModelOptions model with no property values
				deleteLanguageModelOptionsModelNew := new(speechtotextv1.DeleteLanguageModelOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = speechToTextService.DeleteLanguageModel(deleteLanguageModelOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`TrainLanguageModel(trainLanguageModelOptions *TrainLanguageModelOptions) - Operation response error`, func() {
		trainLanguageModelPath := "/v1/customizations/testString/train"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(trainLanguageModelPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["word_type_to_add"]).To(Equal([]string{"all"}))


					// TODO: Add check for customization_weight query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke TrainLanguageModel with error: Operation response processing error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the TrainLanguageModelOptions model
				trainLanguageModelOptionsModel := new(speechtotextv1.TrainLanguageModelOptions)
				trainLanguageModelOptionsModel.CustomizationID = core.StringPtr("testString")
				trainLanguageModelOptionsModel.WordTypeToAdd = core.StringPtr("all")
				trainLanguageModelOptionsModel.CustomizationWeight = core.Float64Ptr(float64(72.5))
				trainLanguageModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := speechToTextService.TrainLanguageModel(trainLanguageModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				speechToTextService.EnableRetries(0, 0)
				result, response, operationErr = speechToTextService.TrainLanguageModel(trainLanguageModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`TrainLanguageModel(trainLanguageModelOptions *TrainLanguageModelOptions)`, func() {
		trainLanguageModelPath := "/v1/customizations/testString/train"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(trainLanguageModelPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["word_type_to_add"]).To(Equal([]string{"all"}))


					// TODO: Add check for customization_weight query parameter

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"warnings": [{"code": "invalid_audio_files", "message": "Message"}]}`)
				}))
			})
			It(`Invoke TrainLanguageModel successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := speechToTextService.TrainLanguageModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TrainLanguageModelOptions model
				trainLanguageModelOptionsModel := new(speechtotextv1.TrainLanguageModelOptions)
				trainLanguageModelOptionsModel.CustomizationID = core.StringPtr("testString")
				trainLanguageModelOptionsModel.WordTypeToAdd = core.StringPtr("all")
				trainLanguageModelOptionsModel.CustomizationWeight = core.Float64Ptr(float64(72.5))
				trainLanguageModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = speechToTextService.TrainLanguageModel(trainLanguageModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.TrainLanguageModelWithContext(ctx, trainLanguageModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				speechToTextService.DisableRetries()
				result, response, operationErr = speechToTextService.TrainLanguageModel(trainLanguageModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.TrainLanguageModelWithContext(ctx, trainLanguageModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke TrainLanguageModel with error: Operation validation and request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the TrainLanguageModelOptions model
				trainLanguageModelOptionsModel := new(speechtotextv1.TrainLanguageModelOptions)
				trainLanguageModelOptionsModel.CustomizationID = core.StringPtr("testString")
				trainLanguageModelOptionsModel.WordTypeToAdd = core.StringPtr("all")
				trainLanguageModelOptionsModel.CustomizationWeight = core.Float64Ptr(float64(72.5))
				trainLanguageModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := speechToTextService.TrainLanguageModel(trainLanguageModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the TrainLanguageModelOptions model with no property values
				trainLanguageModelOptionsModelNew := new(speechtotextv1.TrainLanguageModelOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = speechToTextService.TrainLanguageModel(trainLanguageModelOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ResetLanguageModel(resetLanguageModelOptions *ResetLanguageModelOptions)`, func() {
		resetLanguageModelPath := "/v1/customizations/testString/reset"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(resetLanguageModelPath))
					Expect(req.Method).To(Equal("POST"))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke ResetLanguageModel successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := speechToTextService.ResetLanguageModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the ResetLanguageModelOptions model
				resetLanguageModelOptionsModel := new(speechtotextv1.ResetLanguageModelOptions)
				resetLanguageModelOptionsModel.CustomizationID = core.StringPtr("testString")
				resetLanguageModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = speechToTextService.ResetLanguageModel(resetLanguageModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				speechToTextService.DisableRetries()
				response, operationErr = speechToTextService.ResetLanguageModel(resetLanguageModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke ResetLanguageModel with error: Operation validation and request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the ResetLanguageModelOptions model
				resetLanguageModelOptionsModel := new(speechtotextv1.ResetLanguageModelOptions)
				resetLanguageModelOptionsModel.CustomizationID = core.StringPtr("testString")
				resetLanguageModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := speechToTextService.ResetLanguageModel(resetLanguageModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the ResetLanguageModelOptions model with no property values
				resetLanguageModelOptionsModelNew := new(speechtotextv1.ResetLanguageModelOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = speechToTextService.ResetLanguageModel(resetLanguageModelOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpgradeLanguageModel(upgradeLanguageModelOptions *UpgradeLanguageModelOptions)`, func() {
		upgradeLanguageModelPath := "/v1/customizations/testString/upgrade_model"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(upgradeLanguageModelPath))
					Expect(req.Method).To(Equal("POST"))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpgradeLanguageModel successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := speechToTextService.UpgradeLanguageModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UpgradeLanguageModelOptions model
				upgradeLanguageModelOptionsModel := new(speechtotextv1.UpgradeLanguageModelOptions)
				upgradeLanguageModelOptionsModel.CustomizationID = core.StringPtr("testString")
				upgradeLanguageModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = speechToTextService.UpgradeLanguageModel(upgradeLanguageModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				speechToTextService.DisableRetries()
				response, operationErr = speechToTextService.UpgradeLanguageModel(upgradeLanguageModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UpgradeLanguageModel with error: Operation validation and request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the UpgradeLanguageModelOptions model
				upgradeLanguageModelOptionsModel := new(speechtotextv1.UpgradeLanguageModelOptions)
				upgradeLanguageModelOptionsModel.CustomizationID = core.StringPtr("testString")
				upgradeLanguageModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := speechToTextService.UpgradeLanguageModel(upgradeLanguageModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UpgradeLanguageModelOptions model with no property values
				upgradeLanguageModelOptionsModelNew := new(speechtotextv1.UpgradeLanguageModelOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = speechToTextService.UpgradeLanguageModel(upgradeLanguageModelOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(speechToTextService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
			Expect(speechToTextService.Service.IsSSLDisabled()).To(BeFalse())
			speechToTextService.DisableSSLVerification()
			Expect(speechToTextService.Service.IsSSLDisabled()).To(BeTrue())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(speechToTextService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				URL: "https://speechtotextv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(speechToTextService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SPEECH_TO_TEXT_URL": "https://speechtotextv1/api",
				"SPEECH_TO_TEXT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				})
				Expect(speechToTextService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := speechToTextService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != speechToTextService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(speechToTextService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(speechToTextService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: "https://testService/api",
				})
				Expect(speechToTextService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := speechToTextService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != speechToTextService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(speechToTextService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(speechToTextService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				})
				err := speechToTextService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := speechToTextService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != speechToTextService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(speechToTextService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(speechToTextService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SPEECH_TO_TEXT_URL": "https://speechtotextv1/api",
				"SPEECH_TO_TEXT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(speechToTextService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SPEECH_TO_TEXT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(speechToTextService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = speechtotextv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListCorpora(listCorporaOptions *ListCorporaOptions) - Operation response error`, func() {
		listCorporaPath := "/v1/customizations/testString/corpora"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCorporaPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListCorpora with error: Operation response processing error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the ListCorporaOptions model
				listCorporaOptionsModel := new(speechtotextv1.ListCorporaOptions)
				listCorporaOptionsModel.CustomizationID = core.StringPtr("testString")
				listCorporaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := speechToTextService.ListCorpora(listCorporaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				speechToTextService.EnableRetries(0, 0)
				result, response, operationErr = speechToTextService.ListCorpora(listCorporaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListCorpora(listCorporaOptions *ListCorporaOptions)`, func() {
		listCorporaPath := "/v1/customizations/testString/corpora"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCorporaPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"corpora": [{"name": "Name", "total_words": 10, "out_of_vocabulary_words": 20, "status": "analyzed", "error": "Error"}]}`)
				}))
			})
			It(`Invoke ListCorpora successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := speechToTextService.ListCorpora(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListCorporaOptions model
				listCorporaOptionsModel := new(speechtotextv1.ListCorporaOptions)
				listCorporaOptionsModel.CustomizationID = core.StringPtr("testString")
				listCorporaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = speechToTextService.ListCorpora(listCorporaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.ListCorporaWithContext(ctx, listCorporaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				speechToTextService.DisableRetries()
				result, response, operationErr = speechToTextService.ListCorpora(listCorporaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.ListCorporaWithContext(ctx, listCorporaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListCorpora with error: Operation validation and request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the ListCorporaOptions model
				listCorporaOptionsModel := new(speechtotextv1.ListCorporaOptions)
				listCorporaOptionsModel.CustomizationID = core.StringPtr("testString")
				listCorporaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := speechToTextService.ListCorpora(listCorporaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListCorporaOptions model with no property values
				listCorporaOptionsModelNew := new(speechtotextv1.ListCorporaOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = speechToTextService.ListCorpora(listCorporaOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`AddCorpus(addCorpusOptions *AddCorpusOptions)`, func() {
		addCorpusPath := "/v1/customizations/testString/corpora/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addCorpusPath))
					Expect(req.Method).To(Equal("POST"))


					// TODO: Add check for allow_overwrite query parameter

					res.WriteHeader(201)
				}))
			})
			It(`Invoke AddCorpus successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := speechToTextService.AddCorpus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the AddCorpusOptions model
				addCorpusOptionsModel := new(speechtotextv1.AddCorpusOptions)
				addCorpusOptionsModel.CustomizationID = core.StringPtr("testString")
				addCorpusOptionsModel.CorpusName = core.StringPtr("testString")
				addCorpusOptionsModel.CorpusFile = CreateMockReader("This is a mock file.")
				addCorpusOptionsModel.AllowOverwrite = core.BoolPtr(true)
				addCorpusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = speechToTextService.AddCorpus(addCorpusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				speechToTextService.DisableRetries()
				response, operationErr = speechToTextService.AddCorpus(addCorpusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke AddCorpus with error: Operation validation and request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the AddCorpusOptions model
				addCorpusOptionsModel := new(speechtotextv1.AddCorpusOptions)
				addCorpusOptionsModel.CustomizationID = core.StringPtr("testString")
				addCorpusOptionsModel.CorpusName = core.StringPtr("testString")
				addCorpusOptionsModel.CorpusFile = CreateMockReader("This is a mock file.")
				addCorpusOptionsModel.AllowOverwrite = core.BoolPtr(true)
				addCorpusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := speechToTextService.AddCorpus(addCorpusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the AddCorpusOptions model with no property values
				addCorpusOptionsModelNew := new(speechtotextv1.AddCorpusOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = speechToTextService.AddCorpus(addCorpusOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCorpus(getCorpusOptions *GetCorpusOptions) - Operation response error`, func() {
		getCorpusPath := "/v1/customizations/testString/corpora/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCorpusPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCorpus with error: Operation response processing error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the GetCorpusOptions model
				getCorpusOptionsModel := new(speechtotextv1.GetCorpusOptions)
				getCorpusOptionsModel.CustomizationID = core.StringPtr("testString")
				getCorpusOptionsModel.CorpusName = core.StringPtr("testString")
				getCorpusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := speechToTextService.GetCorpus(getCorpusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				speechToTextService.EnableRetries(0, 0)
				result, response, operationErr = speechToTextService.GetCorpus(getCorpusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetCorpus(getCorpusOptions *GetCorpusOptions)`, func() {
		getCorpusPath := "/v1/customizations/testString/corpora/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCorpusPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "total_words": 10, "out_of_vocabulary_words": 20, "status": "analyzed", "error": "Error"}`)
				}))
			})
			It(`Invoke GetCorpus successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := speechToTextService.GetCorpus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCorpusOptions model
				getCorpusOptionsModel := new(speechtotextv1.GetCorpusOptions)
				getCorpusOptionsModel.CustomizationID = core.StringPtr("testString")
				getCorpusOptionsModel.CorpusName = core.StringPtr("testString")
				getCorpusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = speechToTextService.GetCorpus(getCorpusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.GetCorpusWithContext(ctx, getCorpusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				speechToTextService.DisableRetries()
				result, response, operationErr = speechToTextService.GetCorpus(getCorpusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.GetCorpusWithContext(ctx, getCorpusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetCorpus with error: Operation validation and request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the GetCorpusOptions model
				getCorpusOptionsModel := new(speechtotextv1.GetCorpusOptions)
				getCorpusOptionsModel.CustomizationID = core.StringPtr("testString")
				getCorpusOptionsModel.CorpusName = core.StringPtr("testString")
				getCorpusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := speechToTextService.GetCorpus(getCorpusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetCorpusOptions model with no property values
				getCorpusOptionsModelNew := new(speechtotextv1.GetCorpusOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = speechToTextService.GetCorpus(getCorpusOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteCorpus(deleteCorpusOptions *DeleteCorpusOptions)`, func() {
		deleteCorpusPath := "/v1/customizations/testString/corpora/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteCorpusPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteCorpus successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := speechToTextService.DeleteCorpus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteCorpusOptions model
				deleteCorpusOptionsModel := new(speechtotextv1.DeleteCorpusOptions)
				deleteCorpusOptionsModel.CustomizationID = core.StringPtr("testString")
				deleteCorpusOptionsModel.CorpusName = core.StringPtr("testString")
				deleteCorpusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = speechToTextService.DeleteCorpus(deleteCorpusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				speechToTextService.DisableRetries()
				response, operationErr = speechToTextService.DeleteCorpus(deleteCorpusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteCorpus with error: Operation validation and request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the DeleteCorpusOptions model
				deleteCorpusOptionsModel := new(speechtotextv1.DeleteCorpusOptions)
				deleteCorpusOptionsModel.CustomizationID = core.StringPtr("testString")
				deleteCorpusOptionsModel.CorpusName = core.StringPtr("testString")
				deleteCorpusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := speechToTextService.DeleteCorpus(deleteCorpusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteCorpusOptions model with no property values
				deleteCorpusOptionsModelNew := new(speechtotextv1.DeleteCorpusOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = speechToTextService.DeleteCorpus(deleteCorpusOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(speechToTextService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
			Expect(speechToTextService.Service.IsSSLDisabled()).To(BeFalse())
			speechToTextService.DisableSSLVerification()
			Expect(speechToTextService.Service.IsSSLDisabled()).To(BeTrue())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(speechToTextService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				URL: "https://speechtotextv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(speechToTextService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SPEECH_TO_TEXT_URL": "https://speechtotextv1/api",
				"SPEECH_TO_TEXT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				})
				Expect(speechToTextService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := speechToTextService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != speechToTextService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(speechToTextService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(speechToTextService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: "https://testService/api",
				})
				Expect(speechToTextService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := speechToTextService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != speechToTextService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(speechToTextService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(speechToTextService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				})
				err := speechToTextService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := speechToTextService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != speechToTextService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(speechToTextService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(speechToTextService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SPEECH_TO_TEXT_URL": "https://speechtotextv1/api",
				"SPEECH_TO_TEXT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(speechToTextService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SPEECH_TO_TEXT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(speechToTextService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = speechtotextv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListWords(listWordsOptions *ListWordsOptions) - Operation response error`, func() {
		listWordsPath := "/v1/customizations/testString/words"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listWordsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["word_type"]).To(Equal([]string{"all"}))

					Expect(req.URL.Query()["sort"]).To(Equal([]string{"alphabetical"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListWords with error: Operation response processing error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the ListWordsOptions model
				listWordsOptionsModel := new(speechtotextv1.ListWordsOptions)
				listWordsOptionsModel.CustomizationID = core.StringPtr("testString")
				listWordsOptionsModel.WordType = core.StringPtr("all")
				listWordsOptionsModel.Sort = core.StringPtr("alphabetical")
				listWordsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := speechToTextService.ListWords(listWordsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				speechToTextService.EnableRetries(0, 0)
				result, response, operationErr = speechToTextService.ListWords(listWordsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListWords(listWordsOptions *ListWordsOptions)`, func() {
		listWordsPath := "/v1/customizations/testString/words"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listWordsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["word_type"]).To(Equal([]string{"all"}))

					Expect(req.URL.Query()["sort"]).To(Equal([]string{"alphabetical"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"words": [{"word": "Word", "sounds_like": ["SoundsLike"], "display_as": "DisplayAs", "count": 5, "source": ["Source"], "error": [{"element": "Element"}]}]}`)
				}))
			})
			It(`Invoke ListWords successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := speechToTextService.ListWords(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListWordsOptions model
				listWordsOptionsModel := new(speechtotextv1.ListWordsOptions)
				listWordsOptionsModel.CustomizationID = core.StringPtr("testString")
				listWordsOptionsModel.WordType = core.StringPtr("all")
				listWordsOptionsModel.Sort = core.StringPtr("alphabetical")
				listWordsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = speechToTextService.ListWords(listWordsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.ListWordsWithContext(ctx, listWordsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				speechToTextService.DisableRetries()
				result, response, operationErr = speechToTextService.ListWords(listWordsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.ListWordsWithContext(ctx, listWordsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListWords with error: Operation validation and request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the ListWordsOptions model
				listWordsOptionsModel := new(speechtotextv1.ListWordsOptions)
				listWordsOptionsModel.CustomizationID = core.StringPtr("testString")
				listWordsOptionsModel.WordType = core.StringPtr("all")
				listWordsOptionsModel.Sort = core.StringPtr("alphabetical")
				listWordsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := speechToTextService.ListWords(listWordsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListWordsOptions model with no property values
				listWordsOptionsModelNew := new(speechtotextv1.ListWordsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = speechToTextService.ListWords(listWordsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`AddWords(addWordsOptions *AddWordsOptions)`, func() {
		addWordsPath := "/v1/customizations/testString/words"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addWordsPath))
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

					res.WriteHeader(201)
				}))
			})
			It(`Invoke AddWords successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := speechToTextService.AddWords(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the CustomWord model
				customWordModel := new(speechtotextv1.CustomWord)
				customWordModel.Word = core.StringPtr("testString")
				customWordModel.SoundsLike = []string{"testString"}
				customWordModel.DisplayAs = core.StringPtr("testString")

				// Construct an instance of the AddWordsOptions model
				addWordsOptionsModel := new(speechtotextv1.AddWordsOptions)
				addWordsOptionsModel.CustomizationID = core.StringPtr("testString")
				addWordsOptionsModel.Words = []speechtotextv1.CustomWord{*customWordModel}
				addWordsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = speechToTextService.AddWords(addWordsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				speechToTextService.DisableRetries()
				response, operationErr = speechToTextService.AddWords(addWordsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke AddWords with error: Operation validation and request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the CustomWord model
				customWordModel := new(speechtotextv1.CustomWord)
				customWordModel.Word = core.StringPtr("testString")
				customWordModel.SoundsLike = []string{"testString"}
				customWordModel.DisplayAs = core.StringPtr("testString")

				// Construct an instance of the AddWordsOptions model
				addWordsOptionsModel := new(speechtotextv1.AddWordsOptions)
				addWordsOptionsModel.CustomizationID = core.StringPtr("testString")
				addWordsOptionsModel.Words = []speechtotextv1.CustomWord{*customWordModel}
				addWordsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := speechToTextService.AddWords(addWordsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the AddWordsOptions model with no property values
				addWordsOptionsModelNew := new(speechtotextv1.AddWordsOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = speechToTextService.AddWords(addWordsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`AddWord(addWordOptions *AddWordOptions)`, func() {
		addWordPath := "/v1/customizations/testString/words/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addWordPath))
					Expect(req.Method).To(Equal("PUT"))

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

					res.WriteHeader(201)
				}))
			})
			It(`Invoke AddWord successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := speechToTextService.AddWord(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the AddWordOptions model
				addWordOptionsModel := new(speechtotextv1.AddWordOptions)
				addWordOptionsModel.CustomizationID = core.StringPtr("testString")
				addWordOptionsModel.WordName = core.StringPtr("testString")
				addWordOptionsModel.Word = core.StringPtr("testString")
				addWordOptionsModel.SoundsLike = []string{"testString"}
				addWordOptionsModel.DisplayAs = core.StringPtr("testString")
				addWordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = speechToTextService.AddWord(addWordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				speechToTextService.DisableRetries()
				response, operationErr = speechToTextService.AddWord(addWordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke AddWord with error: Operation validation and request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the AddWordOptions model
				addWordOptionsModel := new(speechtotextv1.AddWordOptions)
				addWordOptionsModel.CustomizationID = core.StringPtr("testString")
				addWordOptionsModel.WordName = core.StringPtr("testString")
				addWordOptionsModel.Word = core.StringPtr("testString")
				addWordOptionsModel.SoundsLike = []string{"testString"}
				addWordOptionsModel.DisplayAs = core.StringPtr("testString")
				addWordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := speechToTextService.AddWord(addWordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the AddWordOptions model with no property values
				addWordOptionsModelNew := new(speechtotextv1.AddWordOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = speechToTextService.AddWord(addWordOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetWord(getWordOptions *GetWordOptions) - Operation response error`, func() {
		getWordPath := "/v1/customizations/testString/words/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWordPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetWord with error: Operation response processing error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the GetWordOptions model
				getWordOptionsModel := new(speechtotextv1.GetWordOptions)
				getWordOptionsModel.CustomizationID = core.StringPtr("testString")
				getWordOptionsModel.WordName = core.StringPtr("testString")
				getWordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := speechToTextService.GetWord(getWordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				speechToTextService.EnableRetries(0, 0)
				result, response, operationErr = speechToTextService.GetWord(getWordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetWord(getWordOptions *GetWordOptions)`, func() {
		getWordPath := "/v1/customizations/testString/words/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWordPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"word": "Word", "sounds_like": ["SoundsLike"], "display_as": "DisplayAs", "count": 5, "source": ["Source"], "error": [{"element": "Element"}]}`)
				}))
			})
			It(`Invoke GetWord successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := speechToTextService.GetWord(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetWordOptions model
				getWordOptionsModel := new(speechtotextv1.GetWordOptions)
				getWordOptionsModel.CustomizationID = core.StringPtr("testString")
				getWordOptionsModel.WordName = core.StringPtr("testString")
				getWordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = speechToTextService.GetWord(getWordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.GetWordWithContext(ctx, getWordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				speechToTextService.DisableRetries()
				result, response, operationErr = speechToTextService.GetWord(getWordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.GetWordWithContext(ctx, getWordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetWord with error: Operation validation and request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the GetWordOptions model
				getWordOptionsModel := new(speechtotextv1.GetWordOptions)
				getWordOptionsModel.CustomizationID = core.StringPtr("testString")
				getWordOptionsModel.WordName = core.StringPtr("testString")
				getWordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := speechToTextService.GetWord(getWordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetWordOptions model with no property values
				getWordOptionsModelNew := new(speechtotextv1.GetWordOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = speechToTextService.GetWord(getWordOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteWord(deleteWordOptions *DeleteWordOptions)`, func() {
		deleteWordPath := "/v1/customizations/testString/words/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteWordPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteWord successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := speechToTextService.DeleteWord(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteWordOptions model
				deleteWordOptionsModel := new(speechtotextv1.DeleteWordOptions)
				deleteWordOptionsModel.CustomizationID = core.StringPtr("testString")
				deleteWordOptionsModel.WordName = core.StringPtr("testString")
				deleteWordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = speechToTextService.DeleteWord(deleteWordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				speechToTextService.DisableRetries()
				response, operationErr = speechToTextService.DeleteWord(deleteWordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteWord with error: Operation validation and request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the DeleteWordOptions model
				deleteWordOptionsModel := new(speechtotextv1.DeleteWordOptions)
				deleteWordOptionsModel.CustomizationID = core.StringPtr("testString")
				deleteWordOptionsModel.WordName = core.StringPtr("testString")
				deleteWordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := speechToTextService.DeleteWord(deleteWordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteWordOptions model with no property values
				deleteWordOptionsModelNew := new(speechtotextv1.DeleteWordOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = speechToTextService.DeleteWord(deleteWordOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(speechToTextService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
			Expect(speechToTextService.Service.IsSSLDisabled()).To(BeFalse())
			speechToTextService.DisableSSLVerification()
			Expect(speechToTextService.Service.IsSSLDisabled()).To(BeTrue())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(speechToTextService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				URL: "https://speechtotextv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(speechToTextService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SPEECH_TO_TEXT_URL": "https://speechtotextv1/api",
				"SPEECH_TO_TEXT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				})
				Expect(speechToTextService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := speechToTextService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != speechToTextService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(speechToTextService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(speechToTextService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: "https://testService/api",
				})
				Expect(speechToTextService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := speechToTextService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != speechToTextService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(speechToTextService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(speechToTextService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				})
				err := speechToTextService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := speechToTextService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != speechToTextService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(speechToTextService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(speechToTextService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SPEECH_TO_TEXT_URL": "https://speechtotextv1/api",
				"SPEECH_TO_TEXT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(speechToTextService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SPEECH_TO_TEXT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(speechToTextService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = speechtotextv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListGrammars(listGrammarsOptions *ListGrammarsOptions) - Operation response error`, func() {
		listGrammarsPath := "/v1/customizations/testString/grammars"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listGrammarsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListGrammars with error: Operation response processing error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the ListGrammarsOptions model
				listGrammarsOptionsModel := new(speechtotextv1.ListGrammarsOptions)
				listGrammarsOptionsModel.CustomizationID = core.StringPtr("testString")
				listGrammarsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := speechToTextService.ListGrammars(listGrammarsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				speechToTextService.EnableRetries(0, 0)
				result, response, operationErr = speechToTextService.ListGrammars(listGrammarsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListGrammars(listGrammarsOptions *ListGrammarsOptions)`, func() {
		listGrammarsPath := "/v1/customizations/testString/grammars"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listGrammarsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"grammars": [{"name": "Name", "out_of_vocabulary_words": 20, "status": "analyzed", "error": "Error"}]}`)
				}))
			})
			It(`Invoke ListGrammars successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := speechToTextService.ListGrammars(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListGrammarsOptions model
				listGrammarsOptionsModel := new(speechtotextv1.ListGrammarsOptions)
				listGrammarsOptionsModel.CustomizationID = core.StringPtr("testString")
				listGrammarsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = speechToTextService.ListGrammars(listGrammarsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.ListGrammarsWithContext(ctx, listGrammarsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				speechToTextService.DisableRetries()
				result, response, operationErr = speechToTextService.ListGrammars(listGrammarsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.ListGrammarsWithContext(ctx, listGrammarsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListGrammars with error: Operation validation and request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the ListGrammarsOptions model
				listGrammarsOptionsModel := new(speechtotextv1.ListGrammarsOptions)
				listGrammarsOptionsModel.CustomizationID = core.StringPtr("testString")
				listGrammarsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := speechToTextService.ListGrammars(listGrammarsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListGrammarsOptions model with no property values
				listGrammarsOptionsModelNew := new(speechtotextv1.ListGrammarsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = speechToTextService.ListGrammars(listGrammarsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`AddGrammar(addGrammarOptions *AddGrammarOptions)`, func() {
		addGrammarPath := "/v1/customizations/testString/grammars/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addGrammarPath))
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

					Expect(req.Header["Content-Type"]).ToNot(BeNil())
					Expect(req.Header["Content-Type"][0]).To(Equal(fmt.Sprintf("%v", "application/srgs")))

					// TODO: Add check for allow_overwrite query parameter

					res.WriteHeader(201)
				}))
			})
			It(`Invoke AddGrammar successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := speechToTextService.AddGrammar(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the AddGrammarOptions model
				addGrammarOptionsModel := new(speechtotextv1.AddGrammarOptions)
				addGrammarOptionsModel.CustomizationID = core.StringPtr("testString")
				addGrammarOptionsModel.GrammarName = core.StringPtr("testString")
				addGrammarOptionsModel.GrammarFile = CreateMockReader("This is a mock file.")
				addGrammarOptionsModel.ContentType = core.StringPtr("application/srgs")
				addGrammarOptionsModel.AllowOverwrite = core.BoolPtr(true)
				addGrammarOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = speechToTextService.AddGrammar(addGrammarOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				speechToTextService.DisableRetries()
				response, operationErr = speechToTextService.AddGrammar(addGrammarOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke AddGrammar with error: Operation validation and request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the AddGrammarOptions model
				addGrammarOptionsModel := new(speechtotextv1.AddGrammarOptions)
				addGrammarOptionsModel.CustomizationID = core.StringPtr("testString")
				addGrammarOptionsModel.GrammarName = core.StringPtr("testString")
				addGrammarOptionsModel.GrammarFile = CreateMockReader("This is a mock file.")
				addGrammarOptionsModel.ContentType = core.StringPtr("application/srgs")
				addGrammarOptionsModel.AllowOverwrite = core.BoolPtr(true)
				addGrammarOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := speechToTextService.AddGrammar(addGrammarOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the AddGrammarOptions model with no property values
				addGrammarOptionsModelNew := new(speechtotextv1.AddGrammarOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = speechToTextService.AddGrammar(addGrammarOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetGrammar(getGrammarOptions *GetGrammarOptions) - Operation response error`, func() {
		getGrammarPath := "/v1/customizations/testString/grammars/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGrammarPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetGrammar with error: Operation response processing error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the GetGrammarOptions model
				getGrammarOptionsModel := new(speechtotextv1.GetGrammarOptions)
				getGrammarOptionsModel.CustomizationID = core.StringPtr("testString")
				getGrammarOptionsModel.GrammarName = core.StringPtr("testString")
				getGrammarOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := speechToTextService.GetGrammar(getGrammarOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				speechToTextService.EnableRetries(0, 0)
				result, response, operationErr = speechToTextService.GetGrammar(getGrammarOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetGrammar(getGrammarOptions *GetGrammarOptions)`, func() {
		getGrammarPath := "/v1/customizations/testString/grammars/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGrammarPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "out_of_vocabulary_words": 20, "status": "analyzed", "error": "Error"}`)
				}))
			})
			It(`Invoke GetGrammar successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := speechToTextService.GetGrammar(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetGrammarOptions model
				getGrammarOptionsModel := new(speechtotextv1.GetGrammarOptions)
				getGrammarOptionsModel.CustomizationID = core.StringPtr("testString")
				getGrammarOptionsModel.GrammarName = core.StringPtr("testString")
				getGrammarOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = speechToTextService.GetGrammar(getGrammarOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.GetGrammarWithContext(ctx, getGrammarOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				speechToTextService.DisableRetries()
				result, response, operationErr = speechToTextService.GetGrammar(getGrammarOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.GetGrammarWithContext(ctx, getGrammarOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetGrammar with error: Operation validation and request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the GetGrammarOptions model
				getGrammarOptionsModel := new(speechtotextv1.GetGrammarOptions)
				getGrammarOptionsModel.CustomizationID = core.StringPtr("testString")
				getGrammarOptionsModel.GrammarName = core.StringPtr("testString")
				getGrammarOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := speechToTextService.GetGrammar(getGrammarOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetGrammarOptions model with no property values
				getGrammarOptionsModelNew := new(speechtotextv1.GetGrammarOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = speechToTextService.GetGrammar(getGrammarOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteGrammar(deleteGrammarOptions *DeleteGrammarOptions)`, func() {
		deleteGrammarPath := "/v1/customizations/testString/grammars/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteGrammarPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteGrammar successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := speechToTextService.DeleteGrammar(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteGrammarOptions model
				deleteGrammarOptionsModel := new(speechtotextv1.DeleteGrammarOptions)
				deleteGrammarOptionsModel.CustomizationID = core.StringPtr("testString")
				deleteGrammarOptionsModel.GrammarName = core.StringPtr("testString")
				deleteGrammarOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = speechToTextService.DeleteGrammar(deleteGrammarOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				speechToTextService.DisableRetries()
				response, operationErr = speechToTextService.DeleteGrammar(deleteGrammarOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteGrammar with error: Operation validation and request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the DeleteGrammarOptions model
				deleteGrammarOptionsModel := new(speechtotextv1.DeleteGrammarOptions)
				deleteGrammarOptionsModel.CustomizationID = core.StringPtr("testString")
				deleteGrammarOptionsModel.GrammarName = core.StringPtr("testString")
				deleteGrammarOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := speechToTextService.DeleteGrammar(deleteGrammarOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteGrammarOptions model with no property values
				deleteGrammarOptionsModelNew := new(speechtotextv1.DeleteGrammarOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = speechToTextService.DeleteGrammar(deleteGrammarOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(speechToTextService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
			Expect(speechToTextService.Service.IsSSLDisabled()).To(BeFalse())
			speechToTextService.DisableSSLVerification()
			Expect(speechToTextService.Service.IsSSLDisabled()).To(BeTrue())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(speechToTextService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				URL: "https://speechtotextv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(speechToTextService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SPEECH_TO_TEXT_URL": "https://speechtotextv1/api",
				"SPEECH_TO_TEXT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				})
				Expect(speechToTextService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := speechToTextService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != speechToTextService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(speechToTextService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(speechToTextService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: "https://testService/api",
				})
				Expect(speechToTextService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := speechToTextService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != speechToTextService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(speechToTextService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(speechToTextService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				})
				err := speechToTextService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := speechToTextService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != speechToTextService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(speechToTextService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(speechToTextService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SPEECH_TO_TEXT_URL": "https://speechtotextv1/api",
				"SPEECH_TO_TEXT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(speechToTextService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SPEECH_TO_TEXT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(speechToTextService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = speechtotextv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`CreateAcousticModel(createAcousticModelOptions *CreateAcousticModelOptions) - Operation response error`, func() {
		createAcousticModelPath := "/v1/acoustic_customizations"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAcousticModelPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateAcousticModel with error: Operation response processing error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the CreateAcousticModelOptions model
				createAcousticModelOptionsModel := new(speechtotextv1.CreateAcousticModelOptions)
				createAcousticModelOptionsModel.Name = core.StringPtr("testString")
				createAcousticModelOptionsModel.BaseModelName = core.StringPtr("ar-AR_BroadbandModel")
				createAcousticModelOptionsModel.Description = core.StringPtr("testString")
				createAcousticModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := speechToTextService.CreateAcousticModel(createAcousticModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				speechToTextService.EnableRetries(0, 0)
				result, response, operationErr = speechToTextService.CreateAcousticModel(createAcousticModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateAcousticModel(createAcousticModelOptions *CreateAcousticModelOptions)`, func() {
		createAcousticModelPath := "/v1/acoustic_customizations"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAcousticModelPath))
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
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"customization_id": "CustomizationID", "created": "Created", "updated": "Updated", "language": "Language", "versions": ["Versions"], "owner": "Owner", "name": "Name", "description": "Description", "base_model_name": "BaseModelName", "status": "pending", "progress": 8, "warnings": "Warnings"}`)
				}))
			})
			It(`Invoke CreateAcousticModel successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := speechToTextService.CreateAcousticModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateAcousticModelOptions model
				createAcousticModelOptionsModel := new(speechtotextv1.CreateAcousticModelOptions)
				createAcousticModelOptionsModel.Name = core.StringPtr("testString")
				createAcousticModelOptionsModel.BaseModelName = core.StringPtr("ar-AR_BroadbandModel")
				createAcousticModelOptionsModel.Description = core.StringPtr("testString")
				createAcousticModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = speechToTextService.CreateAcousticModel(createAcousticModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.CreateAcousticModelWithContext(ctx, createAcousticModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				speechToTextService.DisableRetries()
				result, response, operationErr = speechToTextService.CreateAcousticModel(createAcousticModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.CreateAcousticModelWithContext(ctx, createAcousticModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CreateAcousticModel with error: Operation validation and request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the CreateAcousticModelOptions model
				createAcousticModelOptionsModel := new(speechtotextv1.CreateAcousticModelOptions)
				createAcousticModelOptionsModel.Name = core.StringPtr("testString")
				createAcousticModelOptionsModel.BaseModelName = core.StringPtr("ar-AR_BroadbandModel")
				createAcousticModelOptionsModel.Description = core.StringPtr("testString")
				createAcousticModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := speechToTextService.CreateAcousticModel(createAcousticModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateAcousticModelOptions model with no property values
				createAcousticModelOptionsModelNew := new(speechtotextv1.CreateAcousticModelOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = speechToTextService.CreateAcousticModel(createAcousticModelOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListAcousticModels(listAcousticModelsOptions *ListAcousticModelsOptions) - Operation response error`, func() {
		listAcousticModelsPath := "/v1/acoustic_customizations"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAcousticModelsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["language"]).To(Equal([]string{"ar-AR"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAcousticModels with error: Operation response processing error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the ListAcousticModelsOptions model
				listAcousticModelsOptionsModel := new(speechtotextv1.ListAcousticModelsOptions)
				listAcousticModelsOptionsModel.Language = core.StringPtr("ar-AR")
				listAcousticModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := speechToTextService.ListAcousticModels(listAcousticModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				speechToTextService.EnableRetries(0, 0)
				result, response, operationErr = speechToTextService.ListAcousticModels(listAcousticModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListAcousticModels(listAcousticModelsOptions *ListAcousticModelsOptions)`, func() {
		listAcousticModelsPath := "/v1/acoustic_customizations"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAcousticModelsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["language"]).To(Equal([]string{"ar-AR"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"customizations": [{"customization_id": "CustomizationID", "created": "Created", "updated": "Updated", "language": "Language", "versions": ["Versions"], "owner": "Owner", "name": "Name", "description": "Description", "base_model_name": "BaseModelName", "status": "pending", "progress": 8, "warnings": "Warnings"}]}`)
				}))
			})
			It(`Invoke ListAcousticModels successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := speechToTextService.ListAcousticModels(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAcousticModelsOptions model
				listAcousticModelsOptionsModel := new(speechtotextv1.ListAcousticModelsOptions)
				listAcousticModelsOptionsModel.Language = core.StringPtr("ar-AR")
				listAcousticModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = speechToTextService.ListAcousticModels(listAcousticModelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.ListAcousticModelsWithContext(ctx, listAcousticModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				speechToTextService.DisableRetries()
				result, response, operationErr = speechToTextService.ListAcousticModels(listAcousticModelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.ListAcousticModelsWithContext(ctx, listAcousticModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListAcousticModels with error: Operation request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the ListAcousticModelsOptions model
				listAcousticModelsOptionsModel := new(speechtotextv1.ListAcousticModelsOptions)
				listAcousticModelsOptionsModel.Language = core.StringPtr("ar-AR")
				listAcousticModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := speechToTextService.ListAcousticModels(listAcousticModelsOptionsModel)
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
	Describe(`GetAcousticModel(getAcousticModelOptions *GetAcousticModelOptions) - Operation response error`, func() {
		getAcousticModelPath := "/v1/acoustic_customizations/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAcousticModelPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAcousticModel with error: Operation response processing error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the GetAcousticModelOptions model
				getAcousticModelOptionsModel := new(speechtotextv1.GetAcousticModelOptions)
				getAcousticModelOptionsModel.CustomizationID = core.StringPtr("testString")
				getAcousticModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := speechToTextService.GetAcousticModel(getAcousticModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				speechToTextService.EnableRetries(0, 0)
				result, response, operationErr = speechToTextService.GetAcousticModel(getAcousticModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetAcousticModel(getAcousticModelOptions *GetAcousticModelOptions)`, func() {
		getAcousticModelPath := "/v1/acoustic_customizations/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAcousticModelPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"customization_id": "CustomizationID", "created": "Created", "updated": "Updated", "language": "Language", "versions": ["Versions"], "owner": "Owner", "name": "Name", "description": "Description", "base_model_name": "BaseModelName", "status": "pending", "progress": 8, "warnings": "Warnings"}`)
				}))
			})
			It(`Invoke GetAcousticModel successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := speechToTextService.GetAcousticModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAcousticModelOptions model
				getAcousticModelOptionsModel := new(speechtotextv1.GetAcousticModelOptions)
				getAcousticModelOptionsModel.CustomizationID = core.StringPtr("testString")
				getAcousticModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = speechToTextService.GetAcousticModel(getAcousticModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.GetAcousticModelWithContext(ctx, getAcousticModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				speechToTextService.DisableRetries()
				result, response, operationErr = speechToTextService.GetAcousticModel(getAcousticModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.GetAcousticModelWithContext(ctx, getAcousticModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetAcousticModel with error: Operation validation and request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the GetAcousticModelOptions model
				getAcousticModelOptionsModel := new(speechtotextv1.GetAcousticModelOptions)
				getAcousticModelOptionsModel.CustomizationID = core.StringPtr("testString")
				getAcousticModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := speechToTextService.GetAcousticModel(getAcousticModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAcousticModelOptions model with no property values
				getAcousticModelOptionsModelNew := new(speechtotextv1.GetAcousticModelOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = speechToTextService.GetAcousticModel(getAcousticModelOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteAcousticModel(deleteAcousticModelOptions *DeleteAcousticModelOptions)`, func() {
		deleteAcousticModelPath := "/v1/acoustic_customizations/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAcousticModelPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteAcousticModel successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := speechToTextService.DeleteAcousticModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteAcousticModelOptions model
				deleteAcousticModelOptionsModel := new(speechtotextv1.DeleteAcousticModelOptions)
				deleteAcousticModelOptionsModel.CustomizationID = core.StringPtr("testString")
				deleteAcousticModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = speechToTextService.DeleteAcousticModel(deleteAcousticModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				speechToTextService.DisableRetries()
				response, operationErr = speechToTextService.DeleteAcousticModel(deleteAcousticModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteAcousticModel with error: Operation validation and request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the DeleteAcousticModelOptions model
				deleteAcousticModelOptionsModel := new(speechtotextv1.DeleteAcousticModelOptions)
				deleteAcousticModelOptionsModel.CustomizationID = core.StringPtr("testString")
				deleteAcousticModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := speechToTextService.DeleteAcousticModel(deleteAcousticModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteAcousticModelOptions model with no property values
				deleteAcousticModelOptionsModelNew := new(speechtotextv1.DeleteAcousticModelOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = speechToTextService.DeleteAcousticModel(deleteAcousticModelOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`TrainAcousticModel(trainAcousticModelOptions *TrainAcousticModelOptions) - Operation response error`, func() {
		trainAcousticModelPath := "/v1/acoustic_customizations/testString/train"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(trainAcousticModelPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["custom_language_model_id"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke TrainAcousticModel with error: Operation response processing error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the TrainAcousticModelOptions model
				trainAcousticModelOptionsModel := new(speechtotextv1.TrainAcousticModelOptions)
				trainAcousticModelOptionsModel.CustomizationID = core.StringPtr("testString")
				trainAcousticModelOptionsModel.CustomLanguageModelID = core.StringPtr("testString")
				trainAcousticModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := speechToTextService.TrainAcousticModel(trainAcousticModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				speechToTextService.EnableRetries(0, 0)
				result, response, operationErr = speechToTextService.TrainAcousticModel(trainAcousticModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`TrainAcousticModel(trainAcousticModelOptions *TrainAcousticModelOptions)`, func() {
		trainAcousticModelPath := "/v1/acoustic_customizations/testString/train"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(trainAcousticModelPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["custom_language_model_id"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"warnings": [{"code": "invalid_audio_files", "message": "Message"}]}`)
				}))
			})
			It(`Invoke TrainAcousticModel successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := speechToTextService.TrainAcousticModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TrainAcousticModelOptions model
				trainAcousticModelOptionsModel := new(speechtotextv1.TrainAcousticModelOptions)
				trainAcousticModelOptionsModel.CustomizationID = core.StringPtr("testString")
				trainAcousticModelOptionsModel.CustomLanguageModelID = core.StringPtr("testString")
				trainAcousticModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = speechToTextService.TrainAcousticModel(trainAcousticModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.TrainAcousticModelWithContext(ctx, trainAcousticModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				speechToTextService.DisableRetries()
				result, response, operationErr = speechToTextService.TrainAcousticModel(trainAcousticModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.TrainAcousticModelWithContext(ctx, trainAcousticModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke TrainAcousticModel with error: Operation validation and request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the TrainAcousticModelOptions model
				trainAcousticModelOptionsModel := new(speechtotextv1.TrainAcousticModelOptions)
				trainAcousticModelOptionsModel.CustomizationID = core.StringPtr("testString")
				trainAcousticModelOptionsModel.CustomLanguageModelID = core.StringPtr("testString")
				trainAcousticModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := speechToTextService.TrainAcousticModel(trainAcousticModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the TrainAcousticModelOptions model with no property values
				trainAcousticModelOptionsModelNew := new(speechtotextv1.TrainAcousticModelOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = speechToTextService.TrainAcousticModel(trainAcousticModelOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ResetAcousticModel(resetAcousticModelOptions *ResetAcousticModelOptions)`, func() {
		resetAcousticModelPath := "/v1/acoustic_customizations/testString/reset"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(resetAcousticModelPath))
					Expect(req.Method).To(Equal("POST"))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke ResetAcousticModel successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := speechToTextService.ResetAcousticModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the ResetAcousticModelOptions model
				resetAcousticModelOptionsModel := new(speechtotextv1.ResetAcousticModelOptions)
				resetAcousticModelOptionsModel.CustomizationID = core.StringPtr("testString")
				resetAcousticModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = speechToTextService.ResetAcousticModel(resetAcousticModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				speechToTextService.DisableRetries()
				response, operationErr = speechToTextService.ResetAcousticModel(resetAcousticModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke ResetAcousticModel with error: Operation validation and request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the ResetAcousticModelOptions model
				resetAcousticModelOptionsModel := new(speechtotextv1.ResetAcousticModelOptions)
				resetAcousticModelOptionsModel.CustomizationID = core.StringPtr("testString")
				resetAcousticModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := speechToTextService.ResetAcousticModel(resetAcousticModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the ResetAcousticModelOptions model with no property values
				resetAcousticModelOptionsModelNew := new(speechtotextv1.ResetAcousticModelOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = speechToTextService.ResetAcousticModel(resetAcousticModelOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpgradeAcousticModel(upgradeAcousticModelOptions *UpgradeAcousticModelOptions)`, func() {
		upgradeAcousticModelPath := "/v1/acoustic_customizations/testString/upgrade_model"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(upgradeAcousticModelPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["custom_language_model_id"]).To(Equal([]string{"testString"}))


					// TODO: Add check for force query parameter

					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpgradeAcousticModel successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := speechToTextService.UpgradeAcousticModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UpgradeAcousticModelOptions model
				upgradeAcousticModelOptionsModel := new(speechtotextv1.UpgradeAcousticModelOptions)
				upgradeAcousticModelOptionsModel.CustomizationID = core.StringPtr("testString")
				upgradeAcousticModelOptionsModel.CustomLanguageModelID = core.StringPtr("testString")
				upgradeAcousticModelOptionsModel.Force = core.BoolPtr(true)
				upgradeAcousticModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = speechToTextService.UpgradeAcousticModel(upgradeAcousticModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				speechToTextService.DisableRetries()
				response, operationErr = speechToTextService.UpgradeAcousticModel(upgradeAcousticModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UpgradeAcousticModel with error: Operation validation and request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the UpgradeAcousticModelOptions model
				upgradeAcousticModelOptionsModel := new(speechtotextv1.UpgradeAcousticModelOptions)
				upgradeAcousticModelOptionsModel.CustomizationID = core.StringPtr("testString")
				upgradeAcousticModelOptionsModel.CustomLanguageModelID = core.StringPtr("testString")
				upgradeAcousticModelOptionsModel.Force = core.BoolPtr(true)
				upgradeAcousticModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := speechToTextService.UpgradeAcousticModel(upgradeAcousticModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UpgradeAcousticModelOptions model with no property values
				upgradeAcousticModelOptionsModelNew := new(speechtotextv1.UpgradeAcousticModelOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = speechToTextService.UpgradeAcousticModel(upgradeAcousticModelOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(speechToTextService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
			Expect(speechToTextService.Service.IsSSLDisabled()).To(BeFalse())
			speechToTextService.DisableSSLVerification()
			Expect(speechToTextService.Service.IsSSLDisabled()).To(BeTrue())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(speechToTextService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				URL: "https://speechtotextv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(speechToTextService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SPEECH_TO_TEXT_URL": "https://speechtotextv1/api",
				"SPEECH_TO_TEXT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				})
				Expect(speechToTextService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := speechToTextService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != speechToTextService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(speechToTextService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(speechToTextService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: "https://testService/api",
				})
				Expect(speechToTextService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := speechToTextService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != speechToTextService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(speechToTextService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(speechToTextService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				})
				err := speechToTextService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := speechToTextService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != speechToTextService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(speechToTextService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(speechToTextService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SPEECH_TO_TEXT_URL": "https://speechtotextv1/api",
				"SPEECH_TO_TEXT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(speechToTextService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SPEECH_TO_TEXT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(speechToTextService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = speechtotextv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListAudio(listAudioOptions *ListAudioOptions) - Operation response error`, func() {
		listAudioPath := "/v1/acoustic_customizations/testString/audio"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAudioPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAudio with error: Operation response processing error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the ListAudioOptions model
				listAudioOptionsModel := new(speechtotextv1.ListAudioOptions)
				listAudioOptionsModel.CustomizationID = core.StringPtr("testString")
				listAudioOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := speechToTextService.ListAudio(listAudioOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				speechToTextService.EnableRetries(0, 0)
				result, response, operationErr = speechToTextService.ListAudio(listAudioOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListAudio(listAudioOptions *ListAudioOptions)`, func() {
		listAudioPath := "/v1/acoustic_customizations/testString/audio"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAudioPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_minutes_of_audio": 19, "audio": [{"duration": 8, "name": "Name", "details": {"type": "audio", "codec": "Codec", "frequency": 9, "compression": "zip"}, "status": "ok"}]}`)
				}))
			})
			It(`Invoke ListAudio successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := speechToTextService.ListAudio(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAudioOptions model
				listAudioOptionsModel := new(speechtotextv1.ListAudioOptions)
				listAudioOptionsModel.CustomizationID = core.StringPtr("testString")
				listAudioOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = speechToTextService.ListAudio(listAudioOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.ListAudioWithContext(ctx, listAudioOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				speechToTextService.DisableRetries()
				result, response, operationErr = speechToTextService.ListAudio(listAudioOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.ListAudioWithContext(ctx, listAudioOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListAudio with error: Operation validation and request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the ListAudioOptions model
				listAudioOptionsModel := new(speechtotextv1.ListAudioOptions)
				listAudioOptionsModel.CustomizationID = core.StringPtr("testString")
				listAudioOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := speechToTextService.ListAudio(listAudioOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListAudioOptions model with no property values
				listAudioOptionsModelNew := new(speechtotextv1.ListAudioOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = speechToTextService.ListAudio(listAudioOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`AddAudio(addAudioOptions *AddAudioOptions)`, func() {
		addAudioPath := "/v1/acoustic_customizations/testString/audio/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addAudioPath))
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

					Expect(req.Header["Content-Type"]).ToNot(BeNil())
					Expect(req.Header["Content-Type"][0]).To(Equal(fmt.Sprintf("%v", "application/zip")))
					Expect(req.Header["Contained-Content-Type"]).ToNot(BeNil())
					Expect(req.Header["Contained-Content-Type"][0]).To(Equal(fmt.Sprintf("%v", "audio/alaw")))

					// TODO: Add check for allow_overwrite query parameter

					res.WriteHeader(201)
				}))
			})
			It(`Invoke AddAudio successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := speechToTextService.AddAudio(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the AddAudioOptions model
				addAudioOptionsModel := new(speechtotextv1.AddAudioOptions)
				addAudioOptionsModel.CustomizationID = core.StringPtr("testString")
				addAudioOptionsModel.AudioName = core.StringPtr("testString")
				addAudioOptionsModel.AudioResource = CreateMockReader("This is a mock file.")
				addAudioOptionsModel.ContentType = core.StringPtr("application/zip")
				addAudioOptionsModel.ContainedContentType = core.StringPtr("audio/alaw")
				addAudioOptionsModel.AllowOverwrite = core.BoolPtr(true)
				addAudioOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = speechToTextService.AddAudio(addAudioOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				speechToTextService.DisableRetries()
				response, operationErr = speechToTextService.AddAudio(addAudioOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke AddAudio with error: Operation validation and request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the AddAudioOptions model
				addAudioOptionsModel := new(speechtotextv1.AddAudioOptions)
				addAudioOptionsModel.CustomizationID = core.StringPtr("testString")
				addAudioOptionsModel.AudioName = core.StringPtr("testString")
				addAudioOptionsModel.AudioResource = CreateMockReader("This is a mock file.")
				addAudioOptionsModel.ContentType = core.StringPtr("application/zip")
				addAudioOptionsModel.ContainedContentType = core.StringPtr("audio/alaw")
				addAudioOptionsModel.AllowOverwrite = core.BoolPtr(true)
				addAudioOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := speechToTextService.AddAudio(addAudioOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the AddAudioOptions model with no property values
				addAudioOptionsModelNew := new(speechtotextv1.AddAudioOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = speechToTextService.AddAudio(addAudioOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAudio(getAudioOptions *GetAudioOptions) - Operation response error`, func() {
		getAudioPath := "/v1/acoustic_customizations/testString/audio/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAudioPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAudio with error: Operation response processing error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the GetAudioOptions model
				getAudioOptionsModel := new(speechtotextv1.GetAudioOptions)
				getAudioOptionsModel.CustomizationID = core.StringPtr("testString")
				getAudioOptionsModel.AudioName = core.StringPtr("testString")
				getAudioOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := speechToTextService.GetAudio(getAudioOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				speechToTextService.EnableRetries(0, 0)
				result, response, operationErr = speechToTextService.GetAudio(getAudioOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetAudio(getAudioOptions *GetAudioOptions)`, func() {
		getAudioPath := "/v1/acoustic_customizations/testString/audio/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAudioPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"duration": 8, "name": "Name", "details": {"type": "audio", "codec": "Codec", "frequency": 9, "compression": "zip"}, "status": "ok", "container": {"duration": 8, "name": "Name", "details": {"type": "audio", "codec": "Codec", "frequency": 9, "compression": "zip"}, "status": "ok"}, "audio": [{"duration": 8, "name": "Name", "details": {"type": "audio", "codec": "Codec", "frequency": 9, "compression": "zip"}, "status": "ok"}]}`)
				}))
			})
			It(`Invoke GetAudio successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := speechToTextService.GetAudio(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAudioOptions model
				getAudioOptionsModel := new(speechtotextv1.GetAudioOptions)
				getAudioOptionsModel.CustomizationID = core.StringPtr("testString")
				getAudioOptionsModel.AudioName = core.StringPtr("testString")
				getAudioOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = speechToTextService.GetAudio(getAudioOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.GetAudioWithContext(ctx, getAudioOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				speechToTextService.DisableRetries()
				result, response, operationErr = speechToTextService.GetAudio(getAudioOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = speechToTextService.GetAudioWithContext(ctx, getAudioOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetAudio with error: Operation validation and request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the GetAudioOptions model
				getAudioOptionsModel := new(speechtotextv1.GetAudioOptions)
				getAudioOptionsModel.CustomizationID = core.StringPtr("testString")
				getAudioOptionsModel.AudioName = core.StringPtr("testString")
				getAudioOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := speechToTextService.GetAudio(getAudioOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAudioOptions model with no property values
				getAudioOptionsModelNew := new(speechtotextv1.GetAudioOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = speechToTextService.GetAudio(getAudioOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteAudio(deleteAudioOptions *DeleteAudioOptions)`, func() {
		deleteAudioPath := "/v1/acoustic_customizations/testString/audio/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAudioPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteAudio successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := speechToTextService.DeleteAudio(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteAudioOptions model
				deleteAudioOptionsModel := new(speechtotextv1.DeleteAudioOptions)
				deleteAudioOptionsModel.CustomizationID = core.StringPtr("testString")
				deleteAudioOptionsModel.AudioName = core.StringPtr("testString")
				deleteAudioOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = speechToTextService.DeleteAudio(deleteAudioOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				speechToTextService.DisableRetries()
				response, operationErr = speechToTextService.DeleteAudio(deleteAudioOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteAudio with error: Operation validation and request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the DeleteAudioOptions model
				deleteAudioOptionsModel := new(speechtotextv1.DeleteAudioOptions)
				deleteAudioOptionsModel.CustomizationID = core.StringPtr("testString")
				deleteAudioOptionsModel.AudioName = core.StringPtr("testString")
				deleteAudioOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := speechToTextService.DeleteAudio(deleteAudioOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteAudioOptions model with no property values
				deleteAudioOptionsModelNew := new(speechtotextv1.DeleteAudioOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = speechToTextService.DeleteAudio(deleteAudioOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(speechToTextService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
			Expect(speechToTextService.Service.IsSSLDisabled()).To(BeFalse())
			speechToTextService.DisableSSLVerification()
			Expect(speechToTextService.Service.IsSSLDisabled()).To(BeTrue())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(speechToTextService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				URL: "https://speechtotextv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(speechToTextService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SPEECH_TO_TEXT_URL": "https://speechtotextv1/api",
				"SPEECH_TO_TEXT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				})
				Expect(speechToTextService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := speechToTextService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != speechToTextService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(speechToTextService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(speechToTextService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: "https://testService/api",
				})
				Expect(speechToTextService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := speechToTextService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != speechToTextService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(speechToTextService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(speechToTextService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				})
				err := speechToTextService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := speechToTextService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != speechToTextService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(speechToTextService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(speechToTextService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SPEECH_TO_TEXT_URL": "https://speechtotextv1/api",
				"SPEECH_TO_TEXT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(speechToTextService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SPEECH_TO_TEXT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(speechToTextService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = speechtotextv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})

	Describe(`DeleteUserData(deleteUserDataOptions *DeleteUserDataOptions)`, func() {
		deleteUserDataPath := "/v1/user_data"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteUserDataPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["customer_id"]).To(Equal([]string{"testString"}))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteUserData successfully`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())
				speechToTextService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := speechToTextService.DeleteUserData(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteUserDataOptions model
				deleteUserDataOptionsModel := new(speechtotextv1.DeleteUserDataOptions)
				deleteUserDataOptionsModel.CustomerID = core.StringPtr("testString")
				deleteUserDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = speechToTextService.DeleteUserData(deleteUserDataOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				speechToTextService.DisableRetries()
				response, operationErr = speechToTextService.DeleteUserData(deleteUserDataOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteUserData with error: Operation validation and request error`, func() {
				speechToTextService, serviceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(speechToTextService).ToNot(BeNil())

				// Construct an instance of the DeleteUserDataOptions model
				deleteUserDataOptionsModel := new(speechtotextv1.DeleteUserDataOptions)
				deleteUserDataOptionsModel.CustomerID = core.StringPtr("testString")
				deleteUserDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := speechToTextService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := speechToTextService.DeleteUserData(deleteUserDataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteUserDataOptions model with no property values
				deleteUserDataOptionsModelNew := new(speechtotextv1.DeleteUserDataOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = speechToTextService.DeleteUserData(deleteUserDataOptionsModelNew)
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
			speechToTextService, _ := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				URL:           "http://speechtotextv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewAddAudioOptions successfully`, func() {
				// Construct an instance of the AddAudioOptions model
				customizationID := "testString"
				audioName := "testString"
				audioResource := CreateMockReader("This is a mock file.")
				addAudioOptionsModel := speechToTextService.NewAddAudioOptions(customizationID, audioName, audioResource)
				addAudioOptionsModel.SetCustomizationID("testString")
				addAudioOptionsModel.SetAudioName("testString")
				addAudioOptionsModel.SetAudioResource(CreateMockReader("This is a mock file."))
				addAudioOptionsModel.SetContentType("application/zip")
				addAudioOptionsModel.SetContainedContentType("audio/alaw")
				addAudioOptionsModel.SetAllowOverwrite(true)
				addAudioOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addAudioOptionsModel).ToNot(BeNil())
				Expect(addAudioOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(addAudioOptionsModel.AudioName).To(Equal(core.StringPtr("testString")))
				Expect(addAudioOptionsModel.AudioResource).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(addAudioOptionsModel.ContentType).To(Equal(core.StringPtr("application/zip")))
				Expect(addAudioOptionsModel.ContainedContentType).To(Equal(core.StringPtr("audio/alaw")))
				Expect(addAudioOptionsModel.AllowOverwrite).To(Equal(core.BoolPtr(true)))
				Expect(addAudioOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewAddCorpusOptions successfully`, func() {
				// Construct an instance of the AddCorpusOptions model
				customizationID := "testString"
				corpusName := "testString"
				corpusFile := CreateMockReader("This is a mock file.")
				addCorpusOptionsModel := speechToTextService.NewAddCorpusOptions(customizationID, corpusName, corpusFile)
				addCorpusOptionsModel.SetCustomizationID("testString")
				addCorpusOptionsModel.SetCorpusName("testString")
				addCorpusOptionsModel.SetCorpusFile(CreateMockReader("This is a mock file."))
				addCorpusOptionsModel.SetAllowOverwrite(true)
				addCorpusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addCorpusOptionsModel).ToNot(BeNil())
				Expect(addCorpusOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(addCorpusOptionsModel.CorpusName).To(Equal(core.StringPtr("testString")))
				Expect(addCorpusOptionsModel.CorpusFile).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(addCorpusOptionsModel.AllowOverwrite).To(Equal(core.BoolPtr(true)))
				Expect(addCorpusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewAddGrammarOptions successfully`, func() {
				// Construct an instance of the AddGrammarOptions model
				customizationID := "testString"
				grammarName := "testString"
				grammarFile := CreateMockReader("This is a mock file.")
				contentType := "application/srgs"
				addGrammarOptionsModel := speechToTextService.NewAddGrammarOptions(customizationID, grammarName, grammarFile, contentType)
				addGrammarOptionsModel.SetCustomizationID("testString")
				addGrammarOptionsModel.SetGrammarName("testString")
				addGrammarOptionsModel.SetGrammarFile(CreateMockReader("This is a mock file."))
				addGrammarOptionsModel.SetContentType("application/srgs")
				addGrammarOptionsModel.SetAllowOverwrite(true)
				addGrammarOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addGrammarOptionsModel).ToNot(BeNil())
				Expect(addGrammarOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(addGrammarOptionsModel.GrammarName).To(Equal(core.StringPtr("testString")))
				Expect(addGrammarOptionsModel.GrammarFile).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(addGrammarOptionsModel.ContentType).To(Equal(core.StringPtr("application/srgs")))
				Expect(addGrammarOptionsModel.AllowOverwrite).To(Equal(core.BoolPtr(true)))
				Expect(addGrammarOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewAddWordOptions successfully`, func() {
				// Construct an instance of the AddWordOptions model
				customizationID := "testString"
				wordName := "testString"
				addWordOptionsModel := speechToTextService.NewAddWordOptions(customizationID, wordName)
				addWordOptionsModel.SetCustomizationID("testString")
				addWordOptionsModel.SetWordName("testString")
				addWordOptionsModel.SetWord("testString")
				addWordOptionsModel.SetSoundsLike([]string{"testString"})
				addWordOptionsModel.SetDisplayAs("testString")
				addWordOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addWordOptionsModel).ToNot(BeNil())
				Expect(addWordOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(addWordOptionsModel.WordName).To(Equal(core.StringPtr("testString")))
				Expect(addWordOptionsModel.Word).To(Equal(core.StringPtr("testString")))
				Expect(addWordOptionsModel.SoundsLike).To(Equal([]string{"testString"}))
				Expect(addWordOptionsModel.DisplayAs).To(Equal(core.StringPtr("testString")))
				Expect(addWordOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewAddWordsOptions successfully`, func() {
				// Construct an instance of the CustomWord model
				customWordModel := new(speechtotextv1.CustomWord)
				Expect(customWordModel).ToNot(BeNil())
				customWordModel.Word = core.StringPtr("testString")
				customWordModel.SoundsLike = []string{"testString"}
				customWordModel.DisplayAs = core.StringPtr("testString")
				Expect(customWordModel.Word).To(Equal(core.StringPtr("testString")))
				Expect(customWordModel.SoundsLike).To(Equal([]string{"testString"}))
				Expect(customWordModel.DisplayAs).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the AddWordsOptions model
				customizationID := "testString"
				addWordsOptionsWords := []speechtotextv1.CustomWord{}
				addWordsOptionsModel := speechToTextService.NewAddWordsOptions(customizationID, addWordsOptionsWords)
				addWordsOptionsModel.SetCustomizationID("testString")
				addWordsOptionsModel.SetWords([]speechtotextv1.CustomWord{*customWordModel})
				addWordsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addWordsOptionsModel).ToNot(BeNil())
				Expect(addWordsOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(addWordsOptionsModel.Words).To(Equal([]speechtotextv1.CustomWord{*customWordModel}))
				Expect(addWordsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCheckJobOptions successfully`, func() {
				// Construct an instance of the CheckJobOptions model
				id := "testString"
				checkJobOptionsModel := speechToTextService.NewCheckJobOptions(id)
				checkJobOptionsModel.SetID("testString")
				checkJobOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(checkJobOptionsModel).ToNot(BeNil())
				Expect(checkJobOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(checkJobOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCheckJobsOptions successfully`, func() {
				// Construct an instance of the CheckJobsOptions model
				checkJobsOptionsModel := speechToTextService.NewCheckJobsOptions()
				checkJobsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(checkJobsOptionsModel).ToNot(BeNil())
				Expect(checkJobsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateAcousticModelOptions successfully`, func() {
				// Construct an instance of the CreateAcousticModelOptions model
				createAcousticModelOptionsName := "testString"
				createAcousticModelOptionsBaseModelName := "ar-AR_BroadbandModel"
				createAcousticModelOptionsModel := speechToTextService.NewCreateAcousticModelOptions(createAcousticModelOptionsName, createAcousticModelOptionsBaseModelName)
				createAcousticModelOptionsModel.SetName("testString")
				createAcousticModelOptionsModel.SetBaseModelName("ar-AR_BroadbandModel")
				createAcousticModelOptionsModel.SetDescription("testString")
				createAcousticModelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createAcousticModelOptionsModel).ToNot(BeNil())
				Expect(createAcousticModelOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createAcousticModelOptionsModel.BaseModelName).To(Equal(core.StringPtr("ar-AR_BroadbandModel")))
				Expect(createAcousticModelOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createAcousticModelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateJobOptions successfully`, func() {
				// Construct an instance of the CreateJobOptions model
				audio := CreateMockReader("This is a mock file.")
				createJobOptionsModel := speechToTextService.NewCreateJobOptions(audio)
				createJobOptionsModel.SetAudio(CreateMockReader("This is a mock file."))
				createJobOptionsModel.SetContentType("application/octet-stream")
				createJobOptionsModel.SetModel("ar-AR_BroadbandModel")
				createJobOptionsModel.SetCallbackURL("testString")
				createJobOptionsModel.SetEvents("recognitions.started")
				createJobOptionsModel.SetUserToken("testString")
				createJobOptionsModel.SetResultsTTL(int64(38))
				createJobOptionsModel.SetLanguageCustomizationID("testString")
				createJobOptionsModel.SetAcousticCustomizationID("testString")
				createJobOptionsModel.SetBaseModelVersion("testString")
				createJobOptionsModel.SetCustomizationWeight(float64(72.5))
				createJobOptionsModel.SetInactivityTimeout(int64(38))
				createJobOptionsModel.SetKeywords([]string{"testString"})
				createJobOptionsModel.SetKeywordsThreshold(float32(36.0))
				createJobOptionsModel.SetMaxAlternatives(int64(38))
				createJobOptionsModel.SetWordAlternativesThreshold(float32(36.0))
				createJobOptionsModel.SetWordConfidence(true)
				createJobOptionsModel.SetTimestamps(true)
				createJobOptionsModel.SetProfanityFilter(true)
				createJobOptionsModel.SetSmartFormatting(true)
				createJobOptionsModel.SetSpeakerLabels(true)
				createJobOptionsModel.SetCustomizationID("testString")
				createJobOptionsModel.SetGrammarName("testString")
				createJobOptionsModel.SetRedaction(true)
				createJobOptionsModel.SetProcessingMetrics(true)
				createJobOptionsModel.SetProcessingMetricsInterval(float32(36.0))
				createJobOptionsModel.SetAudioMetrics(true)
				createJobOptionsModel.SetEndOfPhraseSilenceTime(float64(72.5))
				createJobOptionsModel.SetSplitTranscriptAtPhraseEnd(true)
				createJobOptionsModel.SetSpeechDetectorSensitivity(float32(36.0))
				createJobOptionsModel.SetBackgroundAudioSuppression(float32(36.0))
				createJobOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createJobOptionsModel).ToNot(BeNil())
				Expect(createJobOptionsModel.Audio).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(createJobOptionsModel.ContentType).To(Equal(core.StringPtr("application/octet-stream")))
				Expect(createJobOptionsModel.Model).To(Equal(core.StringPtr("ar-AR_BroadbandModel")))
				Expect(createJobOptionsModel.CallbackURL).To(Equal(core.StringPtr("testString")))
				Expect(createJobOptionsModel.Events).To(Equal(core.StringPtr("recognitions.started")))
				Expect(createJobOptionsModel.UserToken).To(Equal(core.StringPtr("testString")))
				Expect(createJobOptionsModel.ResultsTTL).To(Equal(core.Int64Ptr(int64(38))))
				Expect(createJobOptionsModel.LanguageCustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(createJobOptionsModel.AcousticCustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(createJobOptionsModel.BaseModelVersion).To(Equal(core.StringPtr("testString")))
				Expect(createJobOptionsModel.CustomizationWeight).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(createJobOptionsModel.InactivityTimeout).To(Equal(core.Int64Ptr(int64(38))))
				Expect(createJobOptionsModel.Keywords).To(Equal([]string{"testString"}))
				Expect(createJobOptionsModel.KeywordsThreshold).To(Equal(core.Float32Ptr(float32(36.0))))
				Expect(createJobOptionsModel.MaxAlternatives).To(Equal(core.Int64Ptr(int64(38))))
				Expect(createJobOptionsModel.WordAlternativesThreshold).To(Equal(core.Float32Ptr(float32(36.0))))
				Expect(createJobOptionsModel.WordConfidence).To(Equal(core.BoolPtr(true)))
				Expect(createJobOptionsModel.Timestamps).To(Equal(core.BoolPtr(true)))
				Expect(createJobOptionsModel.ProfanityFilter).To(Equal(core.BoolPtr(true)))
				Expect(createJobOptionsModel.SmartFormatting).To(Equal(core.BoolPtr(true)))
				Expect(createJobOptionsModel.SpeakerLabels).To(Equal(core.BoolPtr(true)))
				Expect(createJobOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(createJobOptionsModel.GrammarName).To(Equal(core.StringPtr("testString")))
				Expect(createJobOptionsModel.Redaction).To(Equal(core.BoolPtr(true)))
				Expect(createJobOptionsModel.ProcessingMetrics).To(Equal(core.BoolPtr(true)))
				Expect(createJobOptionsModel.ProcessingMetricsInterval).To(Equal(core.Float32Ptr(float32(36.0))))
				Expect(createJobOptionsModel.AudioMetrics).To(Equal(core.BoolPtr(true)))
				Expect(createJobOptionsModel.EndOfPhraseSilenceTime).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(createJobOptionsModel.SplitTranscriptAtPhraseEnd).To(Equal(core.BoolPtr(true)))
				Expect(createJobOptionsModel.SpeechDetectorSensitivity).To(Equal(core.Float32Ptr(float32(36.0))))
				Expect(createJobOptionsModel.BackgroundAudioSuppression).To(Equal(core.Float32Ptr(float32(36.0))))
				Expect(createJobOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateLanguageModelOptions successfully`, func() {
				// Construct an instance of the CreateLanguageModelOptions model
				createLanguageModelOptionsName := "testString"
				createLanguageModelOptionsBaseModelName := "de-DE_BroadbandModel"
				createLanguageModelOptionsModel := speechToTextService.NewCreateLanguageModelOptions(createLanguageModelOptionsName, createLanguageModelOptionsBaseModelName)
				createLanguageModelOptionsModel.SetName("testString")
				createLanguageModelOptionsModel.SetBaseModelName("de-DE_BroadbandModel")
				createLanguageModelOptionsModel.SetDialect("testString")
				createLanguageModelOptionsModel.SetDescription("testString")
				createLanguageModelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createLanguageModelOptionsModel).ToNot(BeNil())
				Expect(createLanguageModelOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createLanguageModelOptionsModel.BaseModelName).To(Equal(core.StringPtr("de-DE_BroadbandModel")))
				Expect(createLanguageModelOptionsModel.Dialect).To(Equal(core.StringPtr("testString")))
				Expect(createLanguageModelOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createLanguageModelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteAcousticModelOptions successfully`, func() {
				// Construct an instance of the DeleteAcousticModelOptions model
				customizationID := "testString"
				deleteAcousticModelOptionsModel := speechToTextService.NewDeleteAcousticModelOptions(customizationID)
				deleteAcousticModelOptionsModel.SetCustomizationID("testString")
				deleteAcousticModelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteAcousticModelOptionsModel).ToNot(BeNil())
				Expect(deleteAcousticModelOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteAcousticModelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteAudioOptions successfully`, func() {
				// Construct an instance of the DeleteAudioOptions model
				customizationID := "testString"
				audioName := "testString"
				deleteAudioOptionsModel := speechToTextService.NewDeleteAudioOptions(customizationID, audioName)
				deleteAudioOptionsModel.SetCustomizationID("testString")
				deleteAudioOptionsModel.SetAudioName("testString")
				deleteAudioOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteAudioOptionsModel).ToNot(BeNil())
				Expect(deleteAudioOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteAudioOptionsModel.AudioName).To(Equal(core.StringPtr("testString")))
				Expect(deleteAudioOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteCorpusOptions successfully`, func() {
				// Construct an instance of the DeleteCorpusOptions model
				customizationID := "testString"
				corpusName := "testString"
				deleteCorpusOptionsModel := speechToTextService.NewDeleteCorpusOptions(customizationID, corpusName)
				deleteCorpusOptionsModel.SetCustomizationID("testString")
				deleteCorpusOptionsModel.SetCorpusName("testString")
				deleteCorpusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteCorpusOptionsModel).ToNot(BeNil())
				Expect(deleteCorpusOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCorpusOptionsModel.CorpusName).To(Equal(core.StringPtr("testString")))
				Expect(deleteCorpusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteGrammarOptions successfully`, func() {
				// Construct an instance of the DeleteGrammarOptions model
				customizationID := "testString"
				grammarName := "testString"
				deleteGrammarOptionsModel := speechToTextService.NewDeleteGrammarOptions(customizationID, grammarName)
				deleteGrammarOptionsModel.SetCustomizationID("testString")
				deleteGrammarOptionsModel.SetGrammarName("testString")
				deleteGrammarOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteGrammarOptionsModel).ToNot(BeNil())
				Expect(deleteGrammarOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteGrammarOptionsModel.GrammarName).To(Equal(core.StringPtr("testString")))
				Expect(deleteGrammarOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteJobOptions successfully`, func() {
				// Construct an instance of the DeleteJobOptions model
				id := "testString"
				deleteJobOptionsModel := speechToTextService.NewDeleteJobOptions(id)
				deleteJobOptionsModel.SetID("testString")
				deleteJobOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteJobOptionsModel).ToNot(BeNil())
				Expect(deleteJobOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteJobOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteLanguageModelOptions successfully`, func() {
				// Construct an instance of the DeleteLanguageModelOptions model
				customizationID := "testString"
				deleteLanguageModelOptionsModel := speechToTextService.NewDeleteLanguageModelOptions(customizationID)
				deleteLanguageModelOptionsModel.SetCustomizationID("testString")
				deleteLanguageModelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteLanguageModelOptionsModel).ToNot(BeNil())
				Expect(deleteLanguageModelOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteLanguageModelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteUserDataOptions successfully`, func() {
				// Construct an instance of the DeleteUserDataOptions model
				customerID := "testString"
				deleteUserDataOptionsModel := speechToTextService.NewDeleteUserDataOptions(customerID)
				deleteUserDataOptionsModel.SetCustomerID("testString")
				deleteUserDataOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteUserDataOptionsModel).ToNot(BeNil())
				Expect(deleteUserDataOptionsModel.CustomerID).To(Equal(core.StringPtr("testString")))
				Expect(deleteUserDataOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteWordOptions successfully`, func() {
				// Construct an instance of the DeleteWordOptions model
				customizationID := "testString"
				wordName := "testString"
				deleteWordOptionsModel := speechToTextService.NewDeleteWordOptions(customizationID, wordName)
				deleteWordOptionsModel.SetCustomizationID("testString")
				deleteWordOptionsModel.SetWordName("testString")
				deleteWordOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteWordOptionsModel).ToNot(BeNil())
				Expect(deleteWordOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteWordOptionsModel.WordName).To(Equal(core.StringPtr("testString")))
				Expect(deleteWordOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAcousticModelOptions successfully`, func() {
				// Construct an instance of the GetAcousticModelOptions model
				customizationID := "testString"
				getAcousticModelOptionsModel := speechToTextService.NewGetAcousticModelOptions(customizationID)
				getAcousticModelOptionsModel.SetCustomizationID("testString")
				getAcousticModelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAcousticModelOptionsModel).ToNot(BeNil())
				Expect(getAcousticModelOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(getAcousticModelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAudioOptions successfully`, func() {
				// Construct an instance of the GetAudioOptions model
				customizationID := "testString"
				audioName := "testString"
				getAudioOptionsModel := speechToTextService.NewGetAudioOptions(customizationID, audioName)
				getAudioOptionsModel.SetCustomizationID("testString")
				getAudioOptionsModel.SetAudioName("testString")
				getAudioOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAudioOptionsModel).ToNot(BeNil())
				Expect(getAudioOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(getAudioOptionsModel.AudioName).To(Equal(core.StringPtr("testString")))
				Expect(getAudioOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCorpusOptions successfully`, func() {
				// Construct an instance of the GetCorpusOptions model
				customizationID := "testString"
				corpusName := "testString"
				getCorpusOptionsModel := speechToTextService.NewGetCorpusOptions(customizationID, corpusName)
				getCorpusOptionsModel.SetCustomizationID("testString")
				getCorpusOptionsModel.SetCorpusName("testString")
				getCorpusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCorpusOptionsModel).ToNot(BeNil())
				Expect(getCorpusOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(getCorpusOptionsModel.CorpusName).To(Equal(core.StringPtr("testString")))
				Expect(getCorpusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetGrammarOptions successfully`, func() {
				// Construct an instance of the GetGrammarOptions model
				customizationID := "testString"
				grammarName := "testString"
				getGrammarOptionsModel := speechToTextService.NewGetGrammarOptions(customizationID, grammarName)
				getGrammarOptionsModel.SetCustomizationID("testString")
				getGrammarOptionsModel.SetGrammarName("testString")
				getGrammarOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getGrammarOptionsModel).ToNot(BeNil())
				Expect(getGrammarOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(getGrammarOptionsModel.GrammarName).To(Equal(core.StringPtr("testString")))
				Expect(getGrammarOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetLanguageModelOptions successfully`, func() {
				// Construct an instance of the GetLanguageModelOptions model
				customizationID := "testString"
				getLanguageModelOptionsModel := speechToTextService.NewGetLanguageModelOptions(customizationID)
				getLanguageModelOptionsModel.SetCustomizationID("testString")
				getLanguageModelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLanguageModelOptionsModel).ToNot(BeNil())
				Expect(getLanguageModelOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(getLanguageModelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetModelOptions successfully`, func() {
				// Construct an instance of the GetModelOptions model
				modelID := "ar-AR_BroadbandModel"
				getModelOptionsModel := speechToTextService.NewGetModelOptions(modelID)
				getModelOptionsModel.SetModelID("ar-AR_BroadbandModel")
				getModelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getModelOptionsModel).ToNot(BeNil())
				Expect(getModelOptionsModel.ModelID).To(Equal(core.StringPtr("ar-AR_BroadbandModel")))
				Expect(getModelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetWordOptions successfully`, func() {
				// Construct an instance of the GetWordOptions model
				customizationID := "testString"
				wordName := "testString"
				getWordOptionsModel := speechToTextService.NewGetWordOptions(customizationID, wordName)
				getWordOptionsModel.SetCustomizationID("testString")
				getWordOptionsModel.SetWordName("testString")
				getWordOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getWordOptionsModel).ToNot(BeNil())
				Expect(getWordOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(getWordOptionsModel.WordName).To(Equal(core.StringPtr("testString")))
				Expect(getWordOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAcousticModelsOptions successfully`, func() {
				// Construct an instance of the ListAcousticModelsOptions model
				listAcousticModelsOptionsModel := speechToTextService.NewListAcousticModelsOptions()
				listAcousticModelsOptionsModel.SetLanguage("ar-AR")
				listAcousticModelsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAcousticModelsOptionsModel).ToNot(BeNil())
				Expect(listAcousticModelsOptionsModel.Language).To(Equal(core.StringPtr("ar-AR")))
				Expect(listAcousticModelsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAudioOptions successfully`, func() {
				// Construct an instance of the ListAudioOptions model
				customizationID := "testString"
				listAudioOptionsModel := speechToTextService.NewListAudioOptions(customizationID)
				listAudioOptionsModel.SetCustomizationID("testString")
				listAudioOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAudioOptionsModel).ToNot(BeNil())
				Expect(listAudioOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(listAudioOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListCorporaOptions successfully`, func() {
				// Construct an instance of the ListCorporaOptions model
				customizationID := "testString"
				listCorporaOptionsModel := speechToTextService.NewListCorporaOptions(customizationID)
				listCorporaOptionsModel.SetCustomizationID("testString")
				listCorporaOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listCorporaOptionsModel).ToNot(BeNil())
				Expect(listCorporaOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(listCorporaOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListGrammarsOptions successfully`, func() {
				// Construct an instance of the ListGrammarsOptions model
				customizationID := "testString"
				listGrammarsOptionsModel := speechToTextService.NewListGrammarsOptions(customizationID)
				listGrammarsOptionsModel.SetCustomizationID("testString")
				listGrammarsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listGrammarsOptionsModel).ToNot(BeNil())
				Expect(listGrammarsOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(listGrammarsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListLanguageModelsOptions successfully`, func() {
				// Construct an instance of the ListLanguageModelsOptions model
				listLanguageModelsOptionsModel := speechToTextService.NewListLanguageModelsOptions()
				listLanguageModelsOptionsModel.SetLanguage("ar-AR")
				listLanguageModelsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listLanguageModelsOptionsModel).ToNot(BeNil())
				Expect(listLanguageModelsOptionsModel.Language).To(Equal(core.StringPtr("ar-AR")))
				Expect(listLanguageModelsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListModelsOptions successfully`, func() {
				// Construct an instance of the ListModelsOptions model
				listModelsOptionsModel := speechToTextService.NewListModelsOptions()
				listModelsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listModelsOptionsModel).ToNot(BeNil())
				Expect(listModelsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListWordsOptions successfully`, func() {
				// Construct an instance of the ListWordsOptions model
				customizationID := "testString"
				listWordsOptionsModel := speechToTextService.NewListWordsOptions(customizationID)
				listWordsOptionsModel.SetCustomizationID("testString")
				listWordsOptionsModel.SetWordType("all")
				listWordsOptionsModel.SetSort("alphabetical")
				listWordsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listWordsOptionsModel).ToNot(BeNil())
				Expect(listWordsOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(listWordsOptionsModel.WordType).To(Equal(core.StringPtr("all")))
				Expect(listWordsOptionsModel.Sort).To(Equal(core.StringPtr("alphabetical")))
				Expect(listWordsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRecognizeOptions successfully`, func() {
				// Construct an instance of the RecognizeOptions model
				audio := CreateMockReader("This is a mock file.")
				recognizeOptionsModel := speechToTextService.NewRecognizeOptions(audio)
				recognizeOptionsModel.SetAudio(CreateMockReader("This is a mock file."))
				recognizeOptionsModel.SetContentType("application/octet-stream")
				recognizeOptionsModel.SetModel("ar-AR_BroadbandModel")
				recognizeOptionsModel.SetLanguageCustomizationID("testString")
				recognizeOptionsModel.SetAcousticCustomizationID("testString")
				recognizeOptionsModel.SetBaseModelVersion("testString")
				recognizeOptionsModel.SetCustomizationWeight(float64(72.5))
				recognizeOptionsModel.SetInactivityTimeout(int64(38))
				recognizeOptionsModel.SetKeywords([]string{"testString"})
				recognizeOptionsModel.SetKeywordsThreshold(float32(36.0))
				recognizeOptionsModel.SetMaxAlternatives(int64(38))
				recognizeOptionsModel.SetWordAlternativesThreshold(float32(36.0))
				recognizeOptionsModel.SetWordConfidence(true)
				recognizeOptionsModel.SetTimestamps(true)
				recognizeOptionsModel.SetProfanityFilter(true)
				recognizeOptionsModel.SetSmartFormatting(true)
				recognizeOptionsModel.SetSpeakerLabels(true)
				recognizeOptionsModel.SetCustomizationID("testString")
				recognizeOptionsModel.SetGrammarName("testString")
				recognizeOptionsModel.SetRedaction(true)
				recognizeOptionsModel.SetAudioMetrics(true)
				recognizeOptionsModel.SetEndOfPhraseSilenceTime(float64(72.5))
				recognizeOptionsModel.SetSplitTranscriptAtPhraseEnd(true)
				recognizeOptionsModel.SetSpeechDetectorSensitivity(float32(36.0))
				recognizeOptionsModel.SetBackgroundAudioSuppression(float32(36.0))
				recognizeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(recognizeOptionsModel).ToNot(BeNil())
				Expect(recognizeOptionsModel.Audio).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(recognizeOptionsModel.ContentType).To(Equal(core.StringPtr("application/octet-stream")))
				Expect(recognizeOptionsModel.Model).To(Equal(core.StringPtr("ar-AR_BroadbandModel")))
				Expect(recognizeOptionsModel.LanguageCustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(recognizeOptionsModel.AcousticCustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(recognizeOptionsModel.BaseModelVersion).To(Equal(core.StringPtr("testString")))
				Expect(recognizeOptionsModel.CustomizationWeight).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(recognizeOptionsModel.InactivityTimeout).To(Equal(core.Int64Ptr(int64(38))))
				Expect(recognizeOptionsModel.Keywords).To(Equal([]string{"testString"}))
				Expect(recognizeOptionsModel.KeywordsThreshold).To(Equal(core.Float32Ptr(float32(36.0))))
				Expect(recognizeOptionsModel.MaxAlternatives).To(Equal(core.Int64Ptr(int64(38))))
				Expect(recognizeOptionsModel.WordAlternativesThreshold).To(Equal(core.Float32Ptr(float32(36.0))))
				Expect(recognizeOptionsModel.WordConfidence).To(Equal(core.BoolPtr(true)))
				Expect(recognizeOptionsModel.Timestamps).To(Equal(core.BoolPtr(true)))
				Expect(recognizeOptionsModel.ProfanityFilter).To(Equal(core.BoolPtr(true)))
				Expect(recognizeOptionsModel.SmartFormatting).To(Equal(core.BoolPtr(true)))
				Expect(recognizeOptionsModel.SpeakerLabels).To(Equal(core.BoolPtr(true)))
				Expect(recognizeOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(recognizeOptionsModel.GrammarName).To(Equal(core.StringPtr("testString")))
				Expect(recognizeOptionsModel.Redaction).To(Equal(core.BoolPtr(true)))
				Expect(recognizeOptionsModel.AudioMetrics).To(Equal(core.BoolPtr(true)))
				Expect(recognizeOptionsModel.EndOfPhraseSilenceTime).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(recognizeOptionsModel.SplitTranscriptAtPhraseEnd).To(Equal(core.BoolPtr(true)))
				Expect(recognizeOptionsModel.SpeechDetectorSensitivity).To(Equal(core.Float32Ptr(float32(36.0))))
				Expect(recognizeOptionsModel.BackgroundAudioSuppression).To(Equal(core.Float32Ptr(float32(36.0))))
				Expect(recognizeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRegisterCallbackOptions successfully`, func() {
				// Construct an instance of the RegisterCallbackOptions model
				callbackURL := "testString"
				registerCallbackOptionsModel := speechToTextService.NewRegisterCallbackOptions(callbackURL)
				registerCallbackOptionsModel.SetCallbackURL("testString")
				registerCallbackOptionsModel.SetUserSecret("testString")
				registerCallbackOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(registerCallbackOptionsModel).ToNot(BeNil())
				Expect(registerCallbackOptionsModel.CallbackURL).To(Equal(core.StringPtr("testString")))
				Expect(registerCallbackOptionsModel.UserSecret).To(Equal(core.StringPtr("testString")))
				Expect(registerCallbackOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewResetAcousticModelOptions successfully`, func() {
				// Construct an instance of the ResetAcousticModelOptions model
				customizationID := "testString"
				resetAcousticModelOptionsModel := speechToTextService.NewResetAcousticModelOptions(customizationID)
				resetAcousticModelOptionsModel.SetCustomizationID("testString")
				resetAcousticModelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(resetAcousticModelOptionsModel).ToNot(BeNil())
				Expect(resetAcousticModelOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(resetAcousticModelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewResetLanguageModelOptions successfully`, func() {
				// Construct an instance of the ResetLanguageModelOptions model
				customizationID := "testString"
				resetLanguageModelOptionsModel := speechToTextService.NewResetLanguageModelOptions(customizationID)
				resetLanguageModelOptionsModel.SetCustomizationID("testString")
				resetLanguageModelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(resetLanguageModelOptionsModel).ToNot(BeNil())
				Expect(resetLanguageModelOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(resetLanguageModelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewTrainAcousticModelOptions successfully`, func() {
				// Construct an instance of the TrainAcousticModelOptions model
				customizationID := "testString"
				trainAcousticModelOptionsModel := speechToTextService.NewTrainAcousticModelOptions(customizationID)
				trainAcousticModelOptionsModel.SetCustomizationID("testString")
				trainAcousticModelOptionsModel.SetCustomLanguageModelID("testString")
				trainAcousticModelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(trainAcousticModelOptionsModel).ToNot(BeNil())
				Expect(trainAcousticModelOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(trainAcousticModelOptionsModel.CustomLanguageModelID).To(Equal(core.StringPtr("testString")))
				Expect(trainAcousticModelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewTrainLanguageModelOptions successfully`, func() {
				// Construct an instance of the TrainLanguageModelOptions model
				customizationID := "testString"
				trainLanguageModelOptionsModel := speechToTextService.NewTrainLanguageModelOptions(customizationID)
				trainLanguageModelOptionsModel.SetCustomizationID("testString")
				trainLanguageModelOptionsModel.SetWordTypeToAdd("all")
				trainLanguageModelOptionsModel.SetCustomizationWeight(float64(72.5))
				trainLanguageModelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(trainLanguageModelOptionsModel).ToNot(BeNil())
				Expect(trainLanguageModelOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(trainLanguageModelOptionsModel.WordTypeToAdd).To(Equal(core.StringPtr("all")))
				Expect(trainLanguageModelOptionsModel.CustomizationWeight).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(trainLanguageModelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUnregisterCallbackOptions successfully`, func() {
				// Construct an instance of the UnregisterCallbackOptions model
				callbackURL := "testString"
				unregisterCallbackOptionsModel := speechToTextService.NewUnregisterCallbackOptions(callbackURL)
				unregisterCallbackOptionsModel.SetCallbackURL("testString")
				unregisterCallbackOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(unregisterCallbackOptionsModel).ToNot(BeNil())
				Expect(unregisterCallbackOptionsModel.CallbackURL).To(Equal(core.StringPtr("testString")))
				Expect(unregisterCallbackOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpgradeAcousticModelOptions successfully`, func() {
				// Construct an instance of the UpgradeAcousticModelOptions model
				customizationID := "testString"
				upgradeAcousticModelOptionsModel := speechToTextService.NewUpgradeAcousticModelOptions(customizationID)
				upgradeAcousticModelOptionsModel.SetCustomizationID("testString")
				upgradeAcousticModelOptionsModel.SetCustomLanguageModelID("testString")
				upgradeAcousticModelOptionsModel.SetForce(true)
				upgradeAcousticModelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(upgradeAcousticModelOptionsModel).ToNot(BeNil())
				Expect(upgradeAcousticModelOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(upgradeAcousticModelOptionsModel.CustomLanguageModelID).To(Equal(core.StringPtr("testString")))
				Expect(upgradeAcousticModelOptionsModel.Force).To(Equal(core.BoolPtr(true)))
				Expect(upgradeAcousticModelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpgradeLanguageModelOptions successfully`, func() {
				// Construct an instance of the UpgradeLanguageModelOptions model
				customizationID := "testString"
				upgradeLanguageModelOptionsModel := speechToTextService.NewUpgradeLanguageModelOptions(customizationID)
				upgradeLanguageModelOptionsModel.SetCustomizationID("testString")
				upgradeLanguageModelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(upgradeLanguageModelOptionsModel).ToNot(BeNil())
				Expect(upgradeLanguageModelOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(upgradeLanguageModelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
