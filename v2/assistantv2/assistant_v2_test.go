/**
 * (C) Copyright IBM Corp. 2018, 2021.
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

package assistantv2_test

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
	"github.com/watson-developer-cloud/go-sdk/v2/assistantv2"
)

var _ = Describe(`AssistantV2`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		version := "testString"
		It(`Instantiate service client`, func() {
			assistantService, serviceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       core.StringPtr(version),
			})
			Expect(assistantService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
			Expect(assistantService.Service.IsSSLDisabled()).To(BeFalse())
			assistantService.DisableSSLVerification()
			Expect(assistantService.Service.IsSSLDisabled()).To(BeTrue())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			assistantService, serviceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(assistantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			assistantService, serviceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
				URL:     "https://assistantv2/api",
				Version: core.StringPtr(version),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(assistantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			assistantService, serviceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{})
			Expect(assistantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CONVERSATION_URL":       "https://assistantv2/api",
				"CONVERSATION_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				assistantService, serviceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					Version: core.StringPtr(version),
				})
				Expect(assistantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := assistantService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != assistantService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(assistantService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(assistantService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				assistantService, serviceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					URL:     "https://testService/api",
					Version: core.StringPtr(version),
				})
				Expect(assistantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(assistantService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := assistantService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != assistantService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(assistantService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(assistantService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				assistantService, serviceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					Version: core.StringPtr(version),
				})
				err := assistantService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(assistantService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := assistantService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != assistantService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(assistantService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(assistantService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CONVERSATION_URL":       "https://assistantv2/api",
				"CONVERSATION_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			assistantService, serviceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(assistantService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CONVERSATION_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			assistantService, serviceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(assistantService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = assistantv2.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`CreateSession(createSessionOptions *CreateSessionOptions) - Operation response error`, func() {
		version := "testString"
		createSessionPath := "/v2/assistants/testString/sessions"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSessionPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateSession with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the CreateSessionOptions model
				createSessionOptionsModel := new(assistantv2.CreateSessionOptions)
				createSessionOptionsModel.AssistantID = core.StringPtr("testString")
				createSessionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.CreateSession(createSessionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.CreateSession(createSessionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateSession(createSessionOptions *CreateSessionOptions)`, func() {
		version := "testString"
		createSessionPath := "/v2/assistants/testString/sessions"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSessionPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"session_id": "SessionID"}`)
				}))
			})
			It(`Invoke CreateSession successfully with retries`, func() {
				assistantService, serviceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the CreateSessionOptions model
				createSessionOptionsModel := new(assistantv2.CreateSessionOptions)
				createSessionOptionsModel.AssistantID = core.StringPtr("testString")
				createSessionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.CreateSessionWithContext(ctx, createSessionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.CreateSession(createSessionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.CreateSessionWithContext(ctx, createSessionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createSessionPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"session_id": "SessionID"}`)
				}))
			})
			It(`Invoke CreateSession successfully`, func() {
				assistantService, serviceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.CreateSession(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateSessionOptions model
				createSessionOptionsModel := new(assistantv2.CreateSessionOptions)
				createSessionOptionsModel.AssistantID = core.StringPtr("testString")
				createSessionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.CreateSession(createSessionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateSession with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the CreateSessionOptions model
				createSessionOptionsModel := new(assistantv2.CreateSessionOptions)
				createSessionOptionsModel.AssistantID = core.StringPtr("testString")
				createSessionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.CreateSession(createSessionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateSessionOptions model with no property values
				createSessionOptionsModelNew := new(assistantv2.CreateSessionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = assistantService.CreateSession(createSessionOptionsModelNew)
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
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateSession successfully`, func() {
				assistantService, serviceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the CreateSessionOptions model
				createSessionOptionsModel := new(assistantv2.CreateSessionOptions)
				createSessionOptionsModel.AssistantID = core.StringPtr("testString")
				createSessionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.CreateSession(createSessionOptionsModel)
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
	Describe(`DeleteSession(deleteSessionOptions *DeleteSessionOptions)`, func() {
		version := "testString"
		deleteSessionPath := "/v2/assistants/testString/sessions/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteSessionPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteSession successfully`, func() {
				assistantService, serviceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := assistantService.DeleteSession(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteSessionOptions model
				deleteSessionOptionsModel := new(assistantv2.DeleteSessionOptions)
				deleteSessionOptionsModel.AssistantID = core.StringPtr("testString")
				deleteSessionOptionsModel.SessionID = core.StringPtr("testString")
				deleteSessionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = assistantService.DeleteSession(deleteSessionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteSession with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the DeleteSessionOptions model
				deleteSessionOptionsModel := new(assistantv2.DeleteSessionOptions)
				deleteSessionOptionsModel.AssistantID = core.StringPtr("testString")
				deleteSessionOptionsModel.SessionID = core.StringPtr("testString")
				deleteSessionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := assistantService.DeleteSession(deleteSessionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteSessionOptions model with no property values
				deleteSessionOptionsModelNew := new(assistantv2.DeleteSessionOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = assistantService.DeleteSession(deleteSessionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Message(messageOptions *MessageOptions) - Operation response error`, func() {
		version := "testString"
		messagePath := "/v2/assistants/testString/sessions/testString/message"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(messagePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke Message with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the RuntimeIntent model
				runtimeIntentModel := new(assistantv2.RuntimeIntent)
				runtimeIntentModel.Intent = core.StringPtr("testString")
				runtimeIntentModel.Confidence = core.Float64Ptr(float64(72.5))

				// Construct an instance of the CaptureGroup model
				captureGroupModel := new(assistantv2.CaptureGroup)
				captureGroupModel.Group = core.StringPtr("testString")
				captureGroupModel.Location = []int64{int64(38)}

				// Construct an instance of the RuntimeEntityInterpretation model
				runtimeEntityInterpretationModel := new(assistantv2.RuntimeEntityInterpretation)
				runtimeEntityInterpretationModel.CalendarType = core.StringPtr("testString")
				runtimeEntityInterpretationModel.DatetimeLink = core.StringPtr("testString")
				runtimeEntityInterpretationModel.Festival = core.StringPtr("testString")
				runtimeEntityInterpretationModel.Granularity = core.StringPtr("day")
				runtimeEntityInterpretationModel.RangeLink = core.StringPtr("testString")
				runtimeEntityInterpretationModel.RangeModifier = core.StringPtr("testString")
				runtimeEntityInterpretationModel.RelativeDay = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeMonth = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeWeek = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeWeekend = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeYear = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificDay = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificDayOfWeek = core.StringPtr("testString")
				runtimeEntityInterpretationModel.SpecificMonth = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificQuarter = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificYear = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.NumericValue = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.Subtype = core.StringPtr("testString")
				runtimeEntityInterpretationModel.PartOfDay = core.StringPtr("testString")
				runtimeEntityInterpretationModel.RelativeHour = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeMinute = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeSecond = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificHour = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificMinute = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificSecond = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.Timezone = core.StringPtr("testString")

				// Construct an instance of the RuntimeEntityAlternative model
				runtimeEntityAlternativeModel := new(assistantv2.RuntimeEntityAlternative)
				runtimeEntityAlternativeModel.Value = core.StringPtr("testString")
				runtimeEntityAlternativeModel.Confidence = core.Float64Ptr(float64(72.5))

				// Construct an instance of the RuntimeEntityRole model
				runtimeEntityRoleModel := new(assistantv2.RuntimeEntityRole)
				runtimeEntityRoleModel.Type = core.StringPtr("date_from")

				// Construct an instance of the RuntimeEntity model
				runtimeEntityModel := new(assistantv2.RuntimeEntity)
				runtimeEntityModel.Entity = core.StringPtr("testString")
				runtimeEntityModel.Location = []int64{int64(38)}
				runtimeEntityModel.Value = core.StringPtr("testString")
				runtimeEntityModel.Confidence = core.Float64Ptr(float64(72.5))
				runtimeEntityModel.Metadata = make(map[string]interface{})
				runtimeEntityModel.Groups = []assistantv2.CaptureGroup{*captureGroupModel}
				runtimeEntityModel.Interpretation = runtimeEntityInterpretationModel
				runtimeEntityModel.Alternatives = []assistantv2.RuntimeEntityAlternative{*runtimeEntityAlternativeModel}
				runtimeEntityModel.Role = runtimeEntityRoleModel

				// Construct an instance of the MessageInputOptionsSpelling model
				messageInputOptionsSpellingModel := new(assistantv2.MessageInputOptionsSpelling)
				messageInputOptionsSpellingModel.Suggestions = core.BoolPtr(true)
				messageInputOptionsSpellingModel.AutoCorrect = core.BoolPtr(true)

				// Construct an instance of the MessageInputOptions model
				messageInputOptionsModel := new(assistantv2.MessageInputOptions)
				messageInputOptionsModel.Restart = core.BoolPtr(true)
				messageInputOptionsModel.AlternateIntents = core.BoolPtr(true)
				messageInputOptionsModel.Spelling = messageInputOptionsSpellingModel
				messageInputOptionsModel.Debug = core.BoolPtr(true)
				messageInputOptionsModel.ReturnContext = core.BoolPtr(true)
				messageInputOptionsModel.Export = core.BoolPtr(true)

				// Construct an instance of the MessageInput model
				messageInputModel := new(assistantv2.MessageInput)
				messageInputModel.MessageType = core.StringPtr("text")
				messageInputModel.Text = core.StringPtr("testString")
				messageInputModel.Intents = []assistantv2.RuntimeIntent{*runtimeIntentModel}
				messageInputModel.Entities = []assistantv2.RuntimeEntity{*runtimeEntityModel}
				messageInputModel.SuggestionID = core.StringPtr("testString")
				messageInputModel.Options = messageInputOptionsModel

				// Construct an instance of the MessageContextGlobalSystem model
				messageContextGlobalSystemModel := new(assistantv2.MessageContextGlobalSystem)
				messageContextGlobalSystemModel.Timezone = core.StringPtr("testString")
				messageContextGlobalSystemModel.UserID = core.StringPtr("testString")
				messageContextGlobalSystemModel.TurnCount = core.Int64Ptr(int64(38))
				messageContextGlobalSystemModel.Locale = core.StringPtr("en-us")
				messageContextGlobalSystemModel.ReferenceTime = core.StringPtr("testString")

				// Construct an instance of the MessageContextGlobal model
				messageContextGlobalModel := new(assistantv2.MessageContextGlobal)
				messageContextGlobalModel.System = messageContextGlobalSystemModel

				// Construct an instance of the MessageContextSkillSystem model
				messageContextSkillSystemModel := new(assistantv2.MessageContextSkillSystem)
				messageContextSkillSystemModel.State = core.StringPtr("testString")
				messageContextSkillSystemModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the MessageContextSkill model
				messageContextSkillModel := new(assistantv2.MessageContextSkill)
				messageContextSkillModel.UserDefined = make(map[string]interface{})
				messageContextSkillModel.System = messageContextSkillSystemModel

				// Construct an instance of the MessageContext model
				messageContextModel := new(assistantv2.MessageContext)
				messageContextModel.Global = messageContextGlobalModel
				messageContextModel.Skills = make(map[string]assistantv2.MessageContextSkill)
				messageContextModel.Skills["foo"] = *messageContextSkillModel

				// Construct an instance of the MessageOptions model
				messageOptionsModel := new(assistantv2.MessageOptions)
				messageOptionsModel.AssistantID = core.StringPtr("testString")
				messageOptionsModel.SessionID = core.StringPtr("testString")
				messageOptionsModel.Input = messageInputModel
				messageOptionsModel.Context = messageContextModel
				messageOptionsModel.UserID = core.StringPtr("testString")
				messageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.Message(messageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.Message(messageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Message(messageOptions *MessageOptions)`, func() {
		version := "testString"
		messagePath := "/v2/assistants/testString/sessions/testString/message"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(messagePath))
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
					fmt.Fprintf(res, "%s", `{"output": {"generic": [{"response_type": "option", "title": "Title", "description": "Description", "preference": "dropdown", "options": [{"label": "Label", "value": {"input": {"message_type": "text", "text": "Text", "intents": [{"intent": "Intent", "confidence": 10}], "entities": [{"entity": "Entity", "location": [8], "value": "Value", "confidence": 10, "metadata": {"mapKey": "anyValue"}, "groups": [{"group": "Group", "location": [8]}], "interpretation": {"calendar_type": "CalendarType", "datetime_link": "DatetimeLink", "festival": "Festival", "granularity": "day", "range_link": "RangeLink", "range_modifier": "RangeModifier", "relative_day": 11, "relative_month": 13, "relative_week": 12, "relative_weekend": 15, "relative_year": 12, "specific_day": 11, "specific_day_of_week": "SpecificDayOfWeek", "specific_month": 13, "specific_quarter": 15, "specific_year": 12, "numeric_value": 12, "subtype": "Subtype", "part_of_day": "PartOfDay", "relative_hour": 12, "relative_minute": 14, "relative_second": 14, "specific_hour": 12, "specific_minute": 14, "specific_second": 14, "timezone": "Timezone"}, "alternatives": [{"value": "Value", "confidence": 10}], "role": {"type": "date_from"}}], "suggestion_id": "SuggestionID", "options": {"restart": false, "alternate_intents": true, "spelling": {"suggestions": false, "auto_correct": false}, "debug": false, "return_context": false, "export": true}}}}], "channels": [{"channel": "Channel"}]}], "intents": [{"intent": "Intent", "confidence": 10}], "entities": [{"entity": "Entity", "location": [8], "value": "Value", "confidence": 10, "metadata": {"mapKey": "anyValue"}, "groups": [{"group": "Group", "location": [8]}], "interpretation": {"calendar_type": "CalendarType", "datetime_link": "DatetimeLink", "festival": "Festival", "granularity": "day", "range_link": "RangeLink", "range_modifier": "RangeModifier", "relative_day": 11, "relative_month": 13, "relative_week": 12, "relative_weekend": 15, "relative_year": 12, "specific_day": 11, "specific_day_of_week": "SpecificDayOfWeek", "specific_month": 13, "specific_quarter": 15, "specific_year": 12, "numeric_value": 12, "subtype": "Subtype", "part_of_day": "PartOfDay", "relative_hour": 12, "relative_minute": 14, "relative_second": 14, "specific_hour": 12, "specific_minute": 14, "specific_second": 14, "timezone": "Timezone"}, "alternatives": [{"value": "Value", "confidence": 10}], "role": {"type": "date_from"}}], "actions": [{"name": "Name", "type": "client", "parameters": {"mapKey": "anyValue"}, "result_variable": "ResultVariable", "credentials": "Credentials"}], "debug": {"nodes_visited": [{"dialog_node": "DialogNode", "title": "Title", "conditions": "Conditions"}], "log_messages": [{"level": "info", "message": "Message", "code": "Code", "source": {"type": "dialog_node", "dialog_node": "DialogNode"}}], "branch_exited": true, "branch_exited_reason": "completed"}, "user_defined": {"mapKey": "anyValue"}, "spelling": {"text": "Text", "original_text": "OriginalText", "suggested_text": "SuggestedText"}}, "context": {"global": {"system": {"timezone": "Timezone", "user_id": "UserID", "turn_count": 9, "locale": "en-us", "reference_time": "ReferenceTime"}, "session_id": "SessionID"}, "skills": {"mapKey": {"user_defined": {"mapKey": {"anyKey": "anyValue"}}, "system": {"state": "State"}}}}, "user_id": "UserID"}`)
				}))
			})
			It(`Invoke Message successfully with retries`, func() {
				assistantService, serviceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the RuntimeIntent model
				runtimeIntentModel := new(assistantv2.RuntimeIntent)
				runtimeIntentModel.Intent = core.StringPtr("testString")
				runtimeIntentModel.Confidence = core.Float64Ptr(float64(72.5))

				// Construct an instance of the CaptureGroup model
				captureGroupModel := new(assistantv2.CaptureGroup)
				captureGroupModel.Group = core.StringPtr("testString")
				captureGroupModel.Location = []int64{int64(38)}

				// Construct an instance of the RuntimeEntityInterpretation model
				runtimeEntityInterpretationModel := new(assistantv2.RuntimeEntityInterpretation)
				runtimeEntityInterpretationModel.CalendarType = core.StringPtr("testString")
				runtimeEntityInterpretationModel.DatetimeLink = core.StringPtr("testString")
				runtimeEntityInterpretationModel.Festival = core.StringPtr("testString")
				runtimeEntityInterpretationModel.Granularity = core.StringPtr("day")
				runtimeEntityInterpretationModel.RangeLink = core.StringPtr("testString")
				runtimeEntityInterpretationModel.RangeModifier = core.StringPtr("testString")
				runtimeEntityInterpretationModel.RelativeDay = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeMonth = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeWeek = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeWeekend = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeYear = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificDay = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificDayOfWeek = core.StringPtr("testString")
				runtimeEntityInterpretationModel.SpecificMonth = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificQuarter = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificYear = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.NumericValue = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.Subtype = core.StringPtr("testString")
				runtimeEntityInterpretationModel.PartOfDay = core.StringPtr("testString")
				runtimeEntityInterpretationModel.RelativeHour = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeMinute = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeSecond = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificHour = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificMinute = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificSecond = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.Timezone = core.StringPtr("testString")

				// Construct an instance of the RuntimeEntityAlternative model
				runtimeEntityAlternativeModel := new(assistantv2.RuntimeEntityAlternative)
				runtimeEntityAlternativeModel.Value = core.StringPtr("testString")
				runtimeEntityAlternativeModel.Confidence = core.Float64Ptr(float64(72.5))

				// Construct an instance of the RuntimeEntityRole model
				runtimeEntityRoleModel := new(assistantv2.RuntimeEntityRole)
				runtimeEntityRoleModel.Type = core.StringPtr("date_from")

				// Construct an instance of the RuntimeEntity model
				runtimeEntityModel := new(assistantv2.RuntimeEntity)
				runtimeEntityModel.Entity = core.StringPtr("testString")
				runtimeEntityModel.Location = []int64{int64(38)}
				runtimeEntityModel.Value = core.StringPtr("testString")
				runtimeEntityModel.Confidence = core.Float64Ptr(float64(72.5))
				runtimeEntityModel.Metadata = make(map[string]interface{})
				runtimeEntityModel.Groups = []assistantv2.CaptureGroup{*captureGroupModel}
				runtimeEntityModel.Interpretation = runtimeEntityInterpretationModel
				runtimeEntityModel.Alternatives = []assistantv2.RuntimeEntityAlternative{*runtimeEntityAlternativeModel}
				runtimeEntityModel.Role = runtimeEntityRoleModel

				// Construct an instance of the MessageInputOptionsSpelling model
				messageInputOptionsSpellingModel := new(assistantv2.MessageInputOptionsSpelling)
				messageInputOptionsSpellingModel.Suggestions = core.BoolPtr(true)
				messageInputOptionsSpellingModel.AutoCorrect = core.BoolPtr(true)

				// Construct an instance of the MessageInputOptions model
				messageInputOptionsModel := new(assistantv2.MessageInputOptions)
				messageInputOptionsModel.Restart = core.BoolPtr(true)
				messageInputOptionsModel.AlternateIntents = core.BoolPtr(true)
				messageInputOptionsModel.Spelling = messageInputOptionsSpellingModel
				messageInputOptionsModel.Debug = core.BoolPtr(true)
				messageInputOptionsModel.ReturnContext = core.BoolPtr(true)
				messageInputOptionsModel.Export = core.BoolPtr(true)

				// Construct an instance of the MessageInput model
				messageInputModel := new(assistantv2.MessageInput)
				messageInputModel.MessageType = core.StringPtr("text")
				messageInputModel.Text = core.StringPtr("testString")
				messageInputModel.Intents = []assistantv2.RuntimeIntent{*runtimeIntentModel}
				messageInputModel.Entities = []assistantv2.RuntimeEntity{*runtimeEntityModel}
				messageInputModel.SuggestionID = core.StringPtr("testString")
				messageInputModel.Options = messageInputOptionsModel

				// Construct an instance of the MessageContextGlobalSystem model
				messageContextGlobalSystemModel := new(assistantv2.MessageContextGlobalSystem)
				messageContextGlobalSystemModel.Timezone = core.StringPtr("testString")
				messageContextGlobalSystemModel.UserID = core.StringPtr("testString")
				messageContextGlobalSystemModel.TurnCount = core.Int64Ptr(int64(38))
				messageContextGlobalSystemModel.Locale = core.StringPtr("en-us")
				messageContextGlobalSystemModel.ReferenceTime = core.StringPtr("testString")

				// Construct an instance of the MessageContextGlobal model
				messageContextGlobalModel := new(assistantv2.MessageContextGlobal)
				messageContextGlobalModel.System = messageContextGlobalSystemModel

				// Construct an instance of the MessageContextSkillSystem model
				messageContextSkillSystemModel := new(assistantv2.MessageContextSkillSystem)
				messageContextSkillSystemModel.State = core.StringPtr("testString")
				messageContextSkillSystemModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the MessageContextSkill model
				messageContextSkillModel := new(assistantv2.MessageContextSkill)
				messageContextSkillModel.UserDefined = make(map[string]interface{})
				messageContextSkillModel.System = messageContextSkillSystemModel

				// Construct an instance of the MessageContext model
				messageContextModel := new(assistantv2.MessageContext)
				messageContextModel.Global = messageContextGlobalModel
				messageContextModel.Skills = make(map[string]assistantv2.MessageContextSkill)
				messageContextModel.Skills["foo"] = *messageContextSkillModel

				// Construct an instance of the MessageOptions model
				messageOptionsModel := new(assistantv2.MessageOptions)
				messageOptionsModel.AssistantID = core.StringPtr("testString")
				messageOptionsModel.SessionID = core.StringPtr("testString")
				messageOptionsModel.Input = messageInputModel
				messageOptionsModel.Context = messageContextModel
				messageOptionsModel.UserID = core.StringPtr("testString")
				messageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.MessageWithContext(ctx, messageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.Message(messageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.MessageWithContext(ctx, messageOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(messagePath))
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
					fmt.Fprintf(res, "%s", `{"output": {"generic": [{"response_type": "option", "title": "Title", "description": "Description", "preference": "dropdown", "options": [{"label": "Label", "value": {"input": {"message_type": "text", "text": "Text", "intents": [{"intent": "Intent", "confidence": 10}], "entities": [{"entity": "Entity", "location": [8], "value": "Value", "confidence": 10, "metadata": {"mapKey": "anyValue"}, "groups": [{"group": "Group", "location": [8]}], "interpretation": {"calendar_type": "CalendarType", "datetime_link": "DatetimeLink", "festival": "Festival", "granularity": "day", "range_link": "RangeLink", "range_modifier": "RangeModifier", "relative_day": 11, "relative_month": 13, "relative_week": 12, "relative_weekend": 15, "relative_year": 12, "specific_day": 11, "specific_day_of_week": "SpecificDayOfWeek", "specific_month": 13, "specific_quarter": 15, "specific_year": 12, "numeric_value": 12, "subtype": "Subtype", "part_of_day": "PartOfDay", "relative_hour": 12, "relative_minute": 14, "relative_second": 14, "specific_hour": 12, "specific_minute": 14, "specific_second": 14, "timezone": "Timezone"}, "alternatives": [{"value": "Value", "confidence": 10}], "role": {"type": "date_from"}}], "suggestion_id": "SuggestionID", "options": {"restart": false, "alternate_intents": true, "spelling": {"suggestions": false, "auto_correct": false}, "debug": false, "return_context": false, "export": true}}}}], "channels": [{"channel": "Channel"}]}], "intents": [{"intent": "Intent", "confidence": 10}], "entities": [{"entity": "Entity", "location": [8], "value": "Value", "confidence": 10, "metadata": {"mapKey": "anyValue"}, "groups": [{"group": "Group", "location": [8]}], "interpretation": {"calendar_type": "CalendarType", "datetime_link": "DatetimeLink", "festival": "Festival", "granularity": "day", "range_link": "RangeLink", "range_modifier": "RangeModifier", "relative_day": 11, "relative_month": 13, "relative_week": 12, "relative_weekend": 15, "relative_year": 12, "specific_day": 11, "specific_day_of_week": "SpecificDayOfWeek", "specific_month": 13, "specific_quarter": 15, "specific_year": 12, "numeric_value": 12, "subtype": "Subtype", "part_of_day": "PartOfDay", "relative_hour": 12, "relative_minute": 14, "relative_second": 14, "specific_hour": 12, "specific_minute": 14, "specific_second": 14, "timezone": "Timezone"}, "alternatives": [{"value": "Value", "confidence": 10}], "role": {"type": "date_from"}}], "actions": [{"name": "Name", "type": "client", "parameters": {"mapKey": "anyValue"}, "result_variable": "ResultVariable", "credentials": "Credentials"}], "debug": {"nodes_visited": [{"dialog_node": "DialogNode", "title": "Title", "conditions": "Conditions"}], "log_messages": [{"level": "info", "message": "Message", "code": "Code", "source": {"type": "dialog_node", "dialog_node": "DialogNode"}}], "branch_exited": true, "branch_exited_reason": "completed"}, "user_defined": {"mapKey": "anyValue"}, "spelling": {"text": "Text", "original_text": "OriginalText", "suggested_text": "SuggestedText"}}, "context": {"global": {"system": {"timezone": "Timezone", "user_id": "UserID", "turn_count": 9, "locale": "en-us", "reference_time": "ReferenceTime"}, "session_id": "SessionID"}, "skills": {"mapKey": {"user_defined": {"mapKey": {"anyKey": "anyValue"}}, "system": {"state": "State"}}}}, "user_id": "UserID"}`)
				}))
			})
			It(`Invoke Message successfully`, func() {
				assistantService, serviceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.Message(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RuntimeIntent model
				runtimeIntentModel := new(assistantv2.RuntimeIntent)
				runtimeIntentModel.Intent = core.StringPtr("testString")
				runtimeIntentModel.Confidence = core.Float64Ptr(float64(72.5))

				// Construct an instance of the CaptureGroup model
				captureGroupModel := new(assistantv2.CaptureGroup)
				captureGroupModel.Group = core.StringPtr("testString")
				captureGroupModel.Location = []int64{int64(38)}

				// Construct an instance of the RuntimeEntityInterpretation model
				runtimeEntityInterpretationModel := new(assistantv2.RuntimeEntityInterpretation)
				runtimeEntityInterpretationModel.CalendarType = core.StringPtr("testString")
				runtimeEntityInterpretationModel.DatetimeLink = core.StringPtr("testString")
				runtimeEntityInterpretationModel.Festival = core.StringPtr("testString")
				runtimeEntityInterpretationModel.Granularity = core.StringPtr("day")
				runtimeEntityInterpretationModel.RangeLink = core.StringPtr("testString")
				runtimeEntityInterpretationModel.RangeModifier = core.StringPtr("testString")
				runtimeEntityInterpretationModel.RelativeDay = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeMonth = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeWeek = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeWeekend = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeYear = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificDay = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificDayOfWeek = core.StringPtr("testString")
				runtimeEntityInterpretationModel.SpecificMonth = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificQuarter = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificYear = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.NumericValue = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.Subtype = core.StringPtr("testString")
				runtimeEntityInterpretationModel.PartOfDay = core.StringPtr("testString")
				runtimeEntityInterpretationModel.RelativeHour = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeMinute = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeSecond = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificHour = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificMinute = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificSecond = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.Timezone = core.StringPtr("testString")

				// Construct an instance of the RuntimeEntityAlternative model
				runtimeEntityAlternativeModel := new(assistantv2.RuntimeEntityAlternative)
				runtimeEntityAlternativeModel.Value = core.StringPtr("testString")
				runtimeEntityAlternativeModel.Confidence = core.Float64Ptr(float64(72.5))

				// Construct an instance of the RuntimeEntityRole model
				runtimeEntityRoleModel := new(assistantv2.RuntimeEntityRole)
				runtimeEntityRoleModel.Type = core.StringPtr("date_from")

				// Construct an instance of the RuntimeEntity model
				runtimeEntityModel := new(assistantv2.RuntimeEntity)
				runtimeEntityModel.Entity = core.StringPtr("testString")
				runtimeEntityModel.Location = []int64{int64(38)}
				runtimeEntityModel.Value = core.StringPtr("testString")
				runtimeEntityModel.Confidence = core.Float64Ptr(float64(72.5))
				runtimeEntityModel.Metadata = make(map[string]interface{})
				runtimeEntityModel.Groups = []assistantv2.CaptureGroup{*captureGroupModel}
				runtimeEntityModel.Interpretation = runtimeEntityInterpretationModel
				runtimeEntityModel.Alternatives = []assistantv2.RuntimeEntityAlternative{*runtimeEntityAlternativeModel}
				runtimeEntityModel.Role = runtimeEntityRoleModel

				// Construct an instance of the MessageInputOptionsSpelling model
				messageInputOptionsSpellingModel := new(assistantv2.MessageInputOptionsSpelling)
				messageInputOptionsSpellingModel.Suggestions = core.BoolPtr(true)
				messageInputOptionsSpellingModel.AutoCorrect = core.BoolPtr(true)

				// Construct an instance of the MessageInputOptions model
				messageInputOptionsModel := new(assistantv2.MessageInputOptions)
				messageInputOptionsModel.Restart = core.BoolPtr(true)
				messageInputOptionsModel.AlternateIntents = core.BoolPtr(true)
				messageInputOptionsModel.Spelling = messageInputOptionsSpellingModel
				messageInputOptionsModel.Debug = core.BoolPtr(true)
				messageInputOptionsModel.ReturnContext = core.BoolPtr(true)
				messageInputOptionsModel.Export = core.BoolPtr(true)

				// Construct an instance of the MessageInput model
				messageInputModel := new(assistantv2.MessageInput)
				messageInputModel.MessageType = core.StringPtr("text")
				messageInputModel.Text = core.StringPtr("testString")
				messageInputModel.Intents = []assistantv2.RuntimeIntent{*runtimeIntentModel}
				messageInputModel.Entities = []assistantv2.RuntimeEntity{*runtimeEntityModel}
				messageInputModel.SuggestionID = core.StringPtr("testString")
				messageInputModel.Options = messageInputOptionsModel

				// Construct an instance of the MessageContextGlobalSystem model
				messageContextGlobalSystemModel := new(assistantv2.MessageContextGlobalSystem)
				messageContextGlobalSystemModel.Timezone = core.StringPtr("testString")
				messageContextGlobalSystemModel.UserID = core.StringPtr("testString")
				messageContextGlobalSystemModel.TurnCount = core.Int64Ptr(int64(38))
				messageContextGlobalSystemModel.Locale = core.StringPtr("en-us")
				messageContextGlobalSystemModel.ReferenceTime = core.StringPtr("testString")

				// Construct an instance of the MessageContextGlobal model
				messageContextGlobalModel := new(assistantv2.MessageContextGlobal)
				messageContextGlobalModel.System = messageContextGlobalSystemModel

				// Construct an instance of the MessageContextSkillSystem model
				messageContextSkillSystemModel := new(assistantv2.MessageContextSkillSystem)
				messageContextSkillSystemModel.State = core.StringPtr("testString")
				messageContextSkillSystemModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the MessageContextSkill model
				messageContextSkillModel := new(assistantv2.MessageContextSkill)
				messageContextSkillModel.UserDefined = make(map[string]interface{})
				messageContextSkillModel.System = messageContextSkillSystemModel

				// Construct an instance of the MessageContext model
				messageContextModel := new(assistantv2.MessageContext)
				messageContextModel.Global = messageContextGlobalModel
				messageContextModel.Skills = make(map[string]assistantv2.MessageContextSkill)
				messageContextModel.Skills["foo"] = *messageContextSkillModel

				// Construct an instance of the MessageOptions model
				messageOptionsModel := new(assistantv2.MessageOptions)
				messageOptionsModel.AssistantID = core.StringPtr("testString")
				messageOptionsModel.SessionID = core.StringPtr("testString")
				messageOptionsModel.Input = messageInputModel
				messageOptionsModel.Context = messageContextModel
				messageOptionsModel.UserID = core.StringPtr("testString")
				messageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.Message(messageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke Message with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the RuntimeIntent model
				runtimeIntentModel := new(assistantv2.RuntimeIntent)
				runtimeIntentModel.Intent = core.StringPtr("testString")
				runtimeIntentModel.Confidence = core.Float64Ptr(float64(72.5))

				// Construct an instance of the CaptureGroup model
				captureGroupModel := new(assistantv2.CaptureGroup)
				captureGroupModel.Group = core.StringPtr("testString")
				captureGroupModel.Location = []int64{int64(38)}

				// Construct an instance of the RuntimeEntityInterpretation model
				runtimeEntityInterpretationModel := new(assistantv2.RuntimeEntityInterpretation)
				runtimeEntityInterpretationModel.CalendarType = core.StringPtr("testString")
				runtimeEntityInterpretationModel.DatetimeLink = core.StringPtr("testString")
				runtimeEntityInterpretationModel.Festival = core.StringPtr("testString")
				runtimeEntityInterpretationModel.Granularity = core.StringPtr("day")
				runtimeEntityInterpretationModel.RangeLink = core.StringPtr("testString")
				runtimeEntityInterpretationModel.RangeModifier = core.StringPtr("testString")
				runtimeEntityInterpretationModel.RelativeDay = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeMonth = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeWeek = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeWeekend = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeYear = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificDay = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificDayOfWeek = core.StringPtr("testString")
				runtimeEntityInterpretationModel.SpecificMonth = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificQuarter = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificYear = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.NumericValue = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.Subtype = core.StringPtr("testString")
				runtimeEntityInterpretationModel.PartOfDay = core.StringPtr("testString")
				runtimeEntityInterpretationModel.RelativeHour = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeMinute = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeSecond = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificHour = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificMinute = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificSecond = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.Timezone = core.StringPtr("testString")

				// Construct an instance of the RuntimeEntityAlternative model
				runtimeEntityAlternativeModel := new(assistantv2.RuntimeEntityAlternative)
				runtimeEntityAlternativeModel.Value = core.StringPtr("testString")
				runtimeEntityAlternativeModel.Confidence = core.Float64Ptr(float64(72.5))

				// Construct an instance of the RuntimeEntityRole model
				runtimeEntityRoleModel := new(assistantv2.RuntimeEntityRole)
				runtimeEntityRoleModel.Type = core.StringPtr("date_from")

				// Construct an instance of the RuntimeEntity model
				runtimeEntityModel := new(assistantv2.RuntimeEntity)
				runtimeEntityModel.Entity = core.StringPtr("testString")
				runtimeEntityModel.Location = []int64{int64(38)}
				runtimeEntityModel.Value = core.StringPtr("testString")
				runtimeEntityModel.Confidence = core.Float64Ptr(float64(72.5))
				runtimeEntityModel.Metadata = make(map[string]interface{})
				runtimeEntityModel.Groups = []assistantv2.CaptureGroup{*captureGroupModel}
				runtimeEntityModel.Interpretation = runtimeEntityInterpretationModel
				runtimeEntityModel.Alternatives = []assistantv2.RuntimeEntityAlternative{*runtimeEntityAlternativeModel}
				runtimeEntityModel.Role = runtimeEntityRoleModel

				// Construct an instance of the MessageInputOptionsSpelling model
				messageInputOptionsSpellingModel := new(assistantv2.MessageInputOptionsSpelling)
				messageInputOptionsSpellingModel.Suggestions = core.BoolPtr(true)
				messageInputOptionsSpellingModel.AutoCorrect = core.BoolPtr(true)

				// Construct an instance of the MessageInputOptions model
				messageInputOptionsModel := new(assistantv2.MessageInputOptions)
				messageInputOptionsModel.Restart = core.BoolPtr(true)
				messageInputOptionsModel.AlternateIntents = core.BoolPtr(true)
				messageInputOptionsModel.Spelling = messageInputOptionsSpellingModel
				messageInputOptionsModel.Debug = core.BoolPtr(true)
				messageInputOptionsModel.ReturnContext = core.BoolPtr(true)
				messageInputOptionsModel.Export = core.BoolPtr(true)

				// Construct an instance of the MessageInput model
				messageInputModel := new(assistantv2.MessageInput)
				messageInputModel.MessageType = core.StringPtr("text")
				messageInputModel.Text = core.StringPtr("testString")
				messageInputModel.Intents = []assistantv2.RuntimeIntent{*runtimeIntentModel}
				messageInputModel.Entities = []assistantv2.RuntimeEntity{*runtimeEntityModel}
				messageInputModel.SuggestionID = core.StringPtr("testString")
				messageInputModel.Options = messageInputOptionsModel

				// Construct an instance of the MessageContextGlobalSystem model
				messageContextGlobalSystemModel := new(assistantv2.MessageContextGlobalSystem)
				messageContextGlobalSystemModel.Timezone = core.StringPtr("testString")
				messageContextGlobalSystemModel.UserID = core.StringPtr("testString")
				messageContextGlobalSystemModel.TurnCount = core.Int64Ptr(int64(38))
				messageContextGlobalSystemModel.Locale = core.StringPtr("en-us")
				messageContextGlobalSystemModel.ReferenceTime = core.StringPtr("testString")

				// Construct an instance of the MessageContextGlobal model
				messageContextGlobalModel := new(assistantv2.MessageContextGlobal)
				messageContextGlobalModel.System = messageContextGlobalSystemModel

				// Construct an instance of the MessageContextSkillSystem model
				messageContextSkillSystemModel := new(assistantv2.MessageContextSkillSystem)
				messageContextSkillSystemModel.State = core.StringPtr("testString")
				messageContextSkillSystemModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the MessageContextSkill model
				messageContextSkillModel := new(assistantv2.MessageContextSkill)
				messageContextSkillModel.UserDefined = make(map[string]interface{})
				messageContextSkillModel.System = messageContextSkillSystemModel

				// Construct an instance of the MessageContext model
				messageContextModel := new(assistantv2.MessageContext)
				messageContextModel.Global = messageContextGlobalModel
				messageContextModel.Skills = make(map[string]assistantv2.MessageContextSkill)
				messageContextModel.Skills["foo"] = *messageContextSkillModel

				// Construct an instance of the MessageOptions model
				messageOptionsModel := new(assistantv2.MessageOptions)
				messageOptionsModel.AssistantID = core.StringPtr("testString")
				messageOptionsModel.SessionID = core.StringPtr("testString")
				messageOptionsModel.Input = messageInputModel
				messageOptionsModel.Context = messageContextModel
				messageOptionsModel.UserID = core.StringPtr("testString")
				messageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.Message(messageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the MessageOptions model with no property values
				messageOptionsModelNew := new(assistantv2.MessageOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = assistantService.Message(messageOptionsModelNew)
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
			It(`Invoke Message successfully`, func() {
				assistantService, serviceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the RuntimeIntent model
				runtimeIntentModel := new(assistantv2.RuntimeIntent)
				runtimeIntentModel.Intent = core.StringPtr("testString")
				runtimeIntentModel.Confidence = core.Float64Ptr(float64(72.5))

				// Construct an instance of the CaptureGroup model
				captureGroupModel := new(assistantv2.CaptureGroup)
				captureGroupModel.Group = core.StringPtr("testString")
				captureGroupModel.Location = []int64{int64(38)}

				// Construct an instance of the RuntimeEntityInterpretation model
				runtimeEntityInterpretationModel := new(assistantv2.RuntimeEntityInterpretation)
				runtimeEntityInterpretationModel.CalendarType = core.StringPtr("testString")
				runtimeEntityInterpretationModel.DatetimeLink = core.StringPtr("testString")
				runtimeEntityInterpretationModel.Festival = core.StringPtr("testString")
				runtimeEntityInterpretationModel.Granularity = core.StringPtr("day")
				runtimeEntityInterpretationModel.RangeLink = core.StringPtr("testString")
				runtimeEntityInterpretationModel.RangeModifier = core.StringPtr("testString")
				runtimeEntityInterpretationModel.RelativeDay = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeMonth = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeWeek = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeWeekend = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeYear = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificDay = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificDayOfWeek = core.StringPtr("testString")
				runtimeEntityInterpretationModel.SpecificMonth = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificQuarter = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificYear = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.NumericValue = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.Subtype = core.StringPtr("testString")
				runtimeEntityInterpretationModel.PartOfDay = core.StringPtr("testString")
				runtimeEntityInterpretationModel.RelativeHour = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeMinute = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeSecond = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificHour = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificMinute = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificSecond = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.Timezone = core.StringPtr("testString")

				// Construct an instance of the RuntimeEntityAlternative model
				runtimeEntityAlternativeModel := new(assistantv2.RuntimeEntityAlternative)
				runtimeEntityAlternativeModel.Value = core.StringPtr("testString")
				runtimeEntityAlternativeModel.Confidence = core.Float64Ptr(float64(72.5))

				// Construct an instance of the RuntimeEntityRole model
				runtimeEntityRoleModel := new(assistantv2.RuntimeEntityRole)
				runtimeEntityRoleModel.Type = core.StringPtr("date_from")

				// Construct an instance of the RuntimeEntity model
				runtimeEntityModel := new(assistantv2.RuntimeEntity)
				runtimeEntityModel.Entity = core.StringPtr("testString")
				runtimeEntityModel.Location = []int64{int64(38)}
				runtimeEntityModel.Value = core.StringPtr("testString")
				runtimeEntityModel.Confidence = core.Float64Ptr(float64(72.5))
				runtimeEntityModel.Metadata = make(map[string]interface{})
				runtimeEntityModel.Groups = []assistantv2.CaptureGroup{*captureGroupModel}
				runtimeEntityModel.Interpretation = runtimeEntityInterpretationModel
				runtimeEntityModel.Alternatives = []assistantv2.RuntimeEntityAlternative{*runtimeEntityAlternativeModel}
				runtimeEntityModel.Role = runtimeEntityRoleModel

				// Construct an instance of the MessageInputOptionsSpelling model
				messageInputOptionsSpellingModel := new(assistantv2.MessageInputOptionsSpelling)
				messageInputOptionsSpellingModel.Suggestions = core.BoolPtr(true)
				messageInputOptionsSpellingModel.AutoCorrect = core.BoolPtr(true)

				// Construct an instance of the MessageInputOptions model
				messageInputOptionsModel := new(assistantv2.MessageInputOptions)
				messageInputOptionsModel.Restart = core.BoolPtr(true)
				messageInputOptionsModel.AlternateIntents = core.BoolPtr(true)
				messageInputOptionsModel.Spelling = messageInputOptionsSpellingModel
				messageInputOptionsModel.Debug = core.BoolPtr(true)
				messageInputOptionsModel.ReturnContext = core.BoolPtr(true)
				messageInputOptionsModel.Export = core.BoolPtr(true)

				// Construct an instance of the MessageInput model
				messageInputModel := new(assistantv2.MessageInput)
				messageInputModel.MessageType = core.StringPtr("text")
				messageInputModel.Text = core.StringPtr("testString")
				messageInputModel.Intents = []assistantv2.RuntimeIntent{*runtimeIntentModel}
				messageInputModel.Entities = []assistantv2.RuntimeEntity{*runtimeEntityModel}
				messageInputModel.SuggestionID = core.StringPtr("testString")
				messageInputModel.Options = messageInputOptionsModel

				// Construct an instance of the MessageContextGlobalSystem model
				messageContextGlobalSystemModel := new(assistantv2.MessageContextGlobalSystem)
				messageContextGlobalSystemModel.Timezone = core.StringPtr("testString")
				messageContextGlobalSystemModel.UserID = core.StringPtr("testString")
				messageContextGlobalSystemModel.TurnCount = core.Int64Ptr(int64(38))
				messageContextGlobalSystemModel.Locale = core.StringPtr("en-us")
				messageContextGlobalSystemModel.ReferenceTime = core.StringPtr("testString")

				// Construct an instance of the MessageContextGlobal model
				messageContextGlobalModel := new(assistantv2.MessageContextGlobal)
				messageContextGlobalModel.System = messageContextGlobalSystemModel

				// Construct an instance of the MessageContextSkillSystem model
				messageContextSkillSystemModel := new(assistantv2.MessageContextSkillSystem)
				messageContextSkillSystemModel.State = core.StringPtr("testString")
				messageContextSkillSystemModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the MessageContextSkill model
				messageContextSkillModel := new(assistantv2.MessageContextSkill)
				messageContextSkillModel.UserDefined = make(map[string]interface{})
				messageContextSkillModel.System = messageContextSkillSystemModel

				// Construct an instance of the MessageContext model
				messageContextModel := new(assistantv2.MessageContext)
				messageContextModel.Global = messageContextGlobalModel
				messageContextModel.Skills = make(map[string]assistantv2.MessageContextSkill)
				messageContextModel.Skills["foo"] = *messageContextSkillModel

				// Construct an instance of the MessageOptions model
				messageOptionsModel := new(assistantv2.MessageOptions)
				messageOptionsModel.AssistantID = core.StringPtr("testString")
				messageOptionsModel.SessionID = core.StringPtr("testString")
				messageOptionsModel.Input = messageInputModel
				messageOptionsModel.Context = messageContextModel
				messageOptionsModel.UserID = core.StringPtr("testString")
				messageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.Message(messageOptionsModel)
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
	Describe(`MessageStateless(messageStatelessOptions *MessageStatelessOptions) - Operation response error`, func() {
		version := "testString"
		messageStatelessPath := "/v2/assistants/testString/message"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(messageStatelessPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke MessageStateless with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the RuntimeIntent model
				runtimeIntentModel := new(assistantv2.RuntimeIntent)
				runtimeIntentModel.Intent = core.StringPtr("testString")
				runtimeIntentModel.Confidence = core.Float64Ptr(float64(72.5))

				// Construct an instance of the CaptureGroup model
				captureGroupModel := new(assistantv2.CaptureGroup)
				captureGroupModel.Group = core.StringPtr("testString")
				captureGroupModel.Location = []int64{int64(38)}

				// Construct an instance of the RuntimeEntityInterpretation model
				runtimeEntityInterpretationModel := new(assistantv2.RuntimeEntityInterpretation)
				runtimeEntityInterpretationModel.CalendarType = core.StringPtr("testString")
				runtimeEntityInterpretationModel.DatetimeLink = core.StringPtr("testString")
				runtimeEntityInterpretationModel.Festival = core.StringPtr("testString")
				runtimeEntityInterpretationModel.Granularity = core.StringPtr("day")
				runtimeEntityInterpretationModel.RangeLink = core.StringPtr("testString")
				runtimeEntityInterpretationModel.RangeModifier = core.StringPtr("testString")
				runtimeEntityInterpretationModel.RelativeDay = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeMonth = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeWeek = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeWeekend = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeYear = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificDay = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificDayOfWeek = core.StringPtr("testString")
				runtimeEntityInterpretationModel.SpecificMonth = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificQuarter = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificYear = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.NumericValue = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.Subtype = core.StringPtr("testString")
				runtimeEntityInterpretationModel.PartOfDay = core.StringPtr("testString")
				runtimeEntityInterpretationModel.RelativeHour = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeMinute = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeSecond = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificHour = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificMinute = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificSecond = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.Timezone = core.StringPtr("testString")

				// Construct an instance of the RuntimeEntityAlternative model
				runtimeEntityAlternativeModel := new(assistantv2.RuntimeEntityAlternative)
				runtimeEntityAlternativeModel.Value = core.StringPtr("testString")
				runtimeEntityAlternativeModel.Confidence = core.Float64Ptr(float64(72.5))

				// Construct an instance of the RuntimeEntityRole model
				runtimeEntityRoleModel := new(assistantv2.RuntimeEntityRole)
				runtimeEntityRoleModel.Type = core.StringPtr("date_from")

				// Construct an instance of the RuntimeEntity model
				runtimeEntityModel := new(assistantv2.RuntimeEntity)
				runtimeEntityModel.Entity = core.StringPtr("testString")
				runtimeEntityModel.Location = []int64{int64(38)}
				runtimeEntityModel.Value = core.StringPtr("testString")
				runtimeEntityModel.Confidence = core.Float64Ptr(float64(72.5))
				runtimeEntityModel.Metadata = make(map[string]interface{})
				runtimeEntityModel.Groups = []assistantv2.CaptureGroup{*captureGroupModel}
				runtimeEntityModel.Interpretation = runtimeEntityInterpretationModel
				runtimeEntityModel.Alternatives = []assistantv2.RuntimeEntityAlternative{*runtimeEntityAlternativeModel}
				runtimeEntityModel.Role = runtimeEntityRoleModel

				// Construct an instance of the MessageInputOptionsSpelling model
				messageInputOptionsSpellingModel := new(assistantv2.MessageInputOptionsSpelling)
				messageInputOptionsSpellingModel.Suggestions = core.BoolPtr(true)
				messageInputOptionsSpellingModel.AutoCorrect = core.BoolPtr(true)

				// Construct an instance of the MessageInputOptionsStateless model
				messageInputOptionsStatelessModel := new(assistantv2.MessageInputOptionsStateless)
				messageInputOptionsStatelessModel.Restart = core.BoolPtr(true)
				messageInputOptionsStatelessModel.AlternateIntents = core.BoolPtr(true)
				messageInputOptionsStatelessModel.Spelling = messageInputOptionsSpellingModel
				messageInputOptionsStatelessModel.Debug = core.BoolPtr(true)

				// Construct an instance of the MessageInputStateless model
				messageInputStatelessModel := new(assistantv2.MessageInputStateless)
				messageInputStatelessModel.MessageType = core.StringPtr("text")
				messageInputStatelessModel.Text = core.StringPtr("testString")
				messageInputStatelessModel.Intents = []assistantv2.RuntimeIntent{*runtimeIntentModel}
				messageInputStatelessModel.Entities = []assistantv2.RuntimeEntity{*runtimeEntityModel}
				messageInputStatelessModel.SuggestionID = core.StringPtr("testString")
				messageInputStatelessModel.Options = messageInputOptionsStatelessModel

				// Construct an instance of the MessageContextGlobalSystem model
				messageContextGlobalSystemModel := new(assistantv2.MessageContextGlobalSystem)
				messageContextGlobalSystemModel.Timezone = core.StringPtr("testString")
				messageContextGlobalSystemModel.UserID = core.StringPtr("testString")
				messageContextGlobalSystemModel.TurnCount = core.Int64Ptr(int64(38))
				messageContextGlobalSystemModel.Locale = core.StringPtr("en-us")
				messageContextGlobalSystemModel.ReferenceTime = core.StringPtr("testString")

				// Construct an instance of the MessageContextGlobalStateless model
				messageContextGlobalStatelessModel := new(assistantv2.MessageContextGlobalStateless)
				messageContextGlobalStatelessModel.System = messageContextGlobalSystemModel
				messageContextGlobalStatelessModel.SessionID = core.StringPtr("testString")

				// Construct an instance of the MessageContextSkillSystem model
				messageContextSkillSystemModel := new(assistantv2.MessageContextSkillSystem)
				messageContextSkillSystemModel.State = core.StringPtr("testString")
				messageContextSkillSystemModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the MessageContextSkill model
				messageContextSkillModel := new(assistantv2.MessageContextSkill)
				messageContextSkillModel.UserDefined = make(map[string]interface{})
				messageContextSkillModel.System = messageContextSkillSystemModel

				// Construct an instance of the MessageContextStateless model
				messageContextStatelessModel := new(assistantv2.MessageContextStateless)
				messageContextStatelessModel.Global = messageContextGlobalStatelessModel
				messageContextStatelessModel.Skills = make(map[string]assistantv2.MessageContextSkill)
				messageContextStatelessModel.Skills["foo"] = *messageContextSkillModel

				// Construct an instance of the MessageStatelessOptions model
				messageStatelessOptionsModel := new(assistantv2.MessageStatelessOptions)
				messageStatelessOptionsModel.AssistantID = core.StringPtr("testString")
				messageStatelessOptionsModel.Input = messageInputStatelessModel
				messageStatelessOptionsModel.Context = messageContextStatelessModel
				messageStatelessOptionsModel.UserID = core.StringPtr("testString")
				messageStatelessOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.MessageStateless(messageStatelessOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.MessageStateless(messageStatelessOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`MessageStateless(messageStatelessOptions *MessageStatelessOptions)`, func() {
		version := "testString"
		messageStatelessPath := "/v2/assistants/testString/message"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(messageStatelessPath))
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
					fmt.Fprintf(res, "%s", `{"output": {"generic": [{"response_type": "option", "title": "Title", "description": "Description", "preference": "dropdown", "options": [{"label": "Label", "value": {"input": {"message_type": "text", "text": "Text", "intents": [{"intent": "Intent", "confidence": 10}], "entities": [{"entity": "Entity", "location": [8], "value": "Value", "confidence": 10, "metadata": {"mapKey": "anyValue"}, "groups": [{"group": "Group", "location": [8]}], "interpretation": {"calendar_type": "CalendarType", "datetime_link": "DatetimeLink", "festival": "Festival", "granularity": "day", "range_link": "RangeLink", "range_modifier": "RangeModifier", "relative_day": 11, "relative_month": 13, "relative_week": 12, "relative_weekend": 15, "relative_year": 12, "specific_day": 11, "specific_day_of_week": "SpecificDayOfWeek", "specific_month": 13, "specific_quarter": 15, "specific_year": 12, "numeric_value": 12, "subtype": "Subtype", "part_of_day": "PartOfDay", "relative_hour": 12, "relative_minute": 14, "relative_second": 14, "specific_hour": 12, "specific_minute": 14, "specific_second": 14, "timezone": "Timezone"}, "alternatives": [{"value": "Value", "confidence": 10}], "role": {"type": "date_from"}}], "suggestion_id": "SuggestionID", "options": {"restart": false, "alternate_intents": true, "spelling": {"suggestions": false, "auto_correct": false}, "debug": false, "return_context": false, "export": true}}}}], "channels": [{"channel": "Channel"}]}], "intents": [{"intent": "Intent", "confidence": 10}], "entities": [{"entity": "Entity", "location": [8], "value": "Value", "confidence": 10, "metadata": {"mapKey": "anyValue"}, "groups": [{"group": "Group", "location": [8]}], "interpretation": {"calendar_type": "CalendarType", "datetime_link": "DatetimeLink", "festival": "Festival", "granularity": "day", "range_link": "RangeLink", "range_modifier": "RangeModifier", "relative_day": 11, "relative_month": 13, "relative_week": 12, "relative_weekend": 15, "relative_year": 12, "specific_day": 11, "specific_day_of_week": "SpecificDayOfWeek", "specific_month": 13, "specific_quarter": 15, "specific_year": 12, "numeric_value": 12, "subtype": "Subtype", "part_of_day": "PartOfDay", "relative_hour": 12, "relative_minute": 14, "relative_second": 14, "specific_hour": 12, "specific_minute": 14, "specific_second": 14, "timezone": "Timezone"}, "alternatives": [{"value": "Value", "confidence": 10}], "role": {"type": "date_from"}}], "actions": [{"name": "Name", "type": "client", "parameters": {"mapKey": "anyValue"}, "result_variable": "ResultVariable", "credentials": "Credentials"}], "debug": {"nodes_visited": [{"dialog_node": "DialogNode", "title": "Title", "conditions": "Conditions"}], "log_messages": [{"level": "info", "message": "Message", "code": "Code", "source": {"type": "dialog_node", "dialog_node": "DialogNode"}}], "branch_exited": true, "branch_exited_reason": "completed"}, "user_defined": {"mapKey": "anyValue"}, "spelling": {"text": "Text", "original_text": "OriginalText", "suggested_text": "SuggestedText"}}, "context": {"global": {"system": {"timezone": "Timezone", "user_id": "UserID", "turn_count": 9, "locale": "en-us", "reference_time": "ReferenceTime"}, "session_id": "SessionID"}, "skills": {"mapKey": {"user_defined": {"mapKey": {"anyKey": "anyValue"}}, "system": {"state": "State"}}}}, "user_id": "UserID"}`)
				}))
			})
			It(`Invoke MessageStateless successfully with retries`, func() {
				assistantService, serviceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the RuntimeIntent model
				runtimeIntentModel := new(assistantv2.RuntimeIntent)
				runtimeIntentModel.Intent = core.StringPtr("testString")
				runtimeIntentModel.Confidence = core.Float64Ptr(float64(72.5))

				// Construct an instance of the CaptureGroup model
				captureGroupModel := new(assistantv2.CaptureGroup)
				captureGroupModel.Group = core.StringPtr("testString")
				captureGroupModel.Location = []int64{int64(38)}

				// Construct an instance of the RuntimeEntityInterpretation model
				runtimeEntityInterpretationModel := new(assistantv2.RuntimeEntityInterpretation)
				runtimeEntityInterpretationModel.CalendarType = core.StringPtr("testString")
				runtimeEntityInterpretationModel.DatetimeLink = core.StringPtr("testString")
				runtimeEntityInterpretationModel.Festival = core.StringPtr("testString")
				runtimeEntityInterpretationModel.Granularity = core.StringPtr("day")
				runtimeEntityInterpretationModel.RangeLink = core.StringPtr("testString")
				runtimeEntityInterpretationModel.RangeModifier = core.StringPtr("testString")
				runtimeEntityInterpretationModel.RelativeDay = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeMonth = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeWeek = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeWeekend = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeYear = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificDay = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificDayOfWeek = core.StringPtr("testString")
				runtimeEntityInterpretationModel.SpecificMonth = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificQuarter = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificYear = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.NumericValue = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.Subtype = core.StringPtr("testString")
				runtimeEntityInterpretationModel.PartOfDay = core.StringPtr("testString")
				runtimeEntityInterpretationModel.RelativeHour = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeMinute = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeSecond = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificHour = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificMinute = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificSecond = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.Timezone = core.StringPtr("testString")

				// Construct an instance of the RuntimeEntityAlternative model
				runtimeEntityAlternativeModel := new(assistantv2.RuntimeEntityAlternative)
				runtimeEntityAlternativeModel.Value = core.StringPtr("testString")
				runtimeEntityAlternativeModel.Confidence = core.Float64Ptr(float64(72.5))

				// Construct an instance of the RuntimeEntityRole model
				runtimeEntityRoleModel := new(assistantv2.RuntimeEntityRole)
				runtimeEntityRoleModel.Type = core.StringPtr("date_from")

				// Construct an instance of the RuntimeEntity model
				runtimeEntityModel := new(assistantv2.RuntimeEntity)
				runtimeEntityModel.Entity = core.StringPtr("testString")
				runtimeEntityModel.Location = []int64{int64(38)}
				runtimeEntityModel.Value = core.StringPtr("testString")
				runtimeEntityModel.Confidence = core.Float64Ptr(float64(72.5))
				runtimeEntityModel.Metadata = make(map[string]interface{})
				runtimeEntityModel.Groups = []assistantv2.CaptureGroup{*captureGroupModel}
				runtimeEntityModel.Interpretation = runtimeEntityInterpretationModel
				runtimeEntityModel.Alternatives = []assistantv2.RuntimeEntityAlternative{*runtimeEntityAlternativeModel}
				runtimeEntityModel.Role = runtimeEntityRoleModel

				// Construct an instance of the MessageInputOptionsSpelling model
				messageInputOptionsSpellingModel := new(assistantv2.MessageInputOptionsSpelling)
				messageInputOptionsSpellingModel.Suggestions = core.BoolPtr(true)
				messageInputOptionsSpellingModel.AutoCorrect = core.BoolPtr(true)

				// Construct an instance of the MessageInputOptionsStateless model
				messageInputOptionsStatelessModel := new(assistantv2.MessageInputOptionsStateless)
				messageInputOptionsStatelessModel.Restart = core.BoolPtr(true)
				messageInputOptionsStatelessModel.AlternateIntents = core.BoolPtr(true)
				messageInputOptionsStatelessModel.Spelling = messageInputOptionsSpellingModel
				messageInputOptionsStatelessModel.Debug = core.BoolPtr(true)

				// Construct an instance of the MessageInputStateless model
				messageInputStatelessModel := new(assistantv2.MessageInputStateless)
				messageInputStatelessModel.MessageType = core.StringPtr("text")
				messageInputStatelessModel.Text = core.StringPtr("testString")
				messageInputStatelessModel.Intents = []assistantv2.RuntimeIntent{*runtimeIntentModel}
				messageInputStatelessModel.Entities = []assistantv2.RuntimeEntity{*runtimeEntityModel}
				messageInputStatelessModel.SuggestionID = core.StringPtr("testString")
				messageInputStatelessModel.Options = messageInputOptionsStatelessModel

				// Construct an instance of the MessageContextGlobalSystem model
				messageContextGlobalSystemModel := new(assistantv2.MessageContextGlobalSystem)
				messageContextGlobalSystemModel.Timezone = core.StringPtr("testString")
				messageContextGlobalSystemModel.UserID = core.StringPtr("testString")
				messageContextGlobalSystemModel.TurnCount = core.Int64Ptr(int64(38))
				messageContextGlobalSystemModel.Locale = core.StringPtr("en-us")
				messageContextGlobalSystemModel.ReferenceTime = core.StringPtr("testString")

				// Construct an instance of the MessageContextGlobalStateless model
				messageContextGlobalStatelessModel := new(assistantv2.MessageContextGlobalStateless)
				messageContextGlobalStatelessModel.System = messageContextGlobalSystemModel
				messageContextGlobalStatelessModel.SessionID = core.StringPtr("testString")

				// Construct an instance of the MessageContextSkillSystem model
				messageContextSkillSystemModel := new(assistantv2.MessageContextSkillSystem)
				messageContextSkillSystemModel.State = core.StringPtr("testString")
				messageContextSkillSystemModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the MessageContextSkill model
				messageContextSkillModel := new(assistantv2.MessageContextSkill)
				messageContextSkillModel.UserDefined = make(map[string]interface{})
				messageContextSkillModel.System = messageContextSkillSystemModel

				// Construct an instance of the MessageContextStateless model
				messageContextStatelessModel := new(assistantv2.MessageContextStateless)
				messageContextStatelessModel.Global = messageContextGlobalStatelessModel
				messageContextStatelessModel.Skills = make(map[string]assistantv2.MessageContextSkill)
				messageContextStatelessModel.Skills["foo"] = *messageContextSkillModel

				// Construct an instance of the MessageStatelessOptions model
				messageStatelessOptionsModel := new(assistantv2.MessageStatelessOptions)
				messageStatelessOptionsModel.AssistantID = core.StringPtr("testString")
				messageStatelessOptionsModel.Input = messageInputStatelessModel
				messageStatelessOptionsModel.Context = messageContextStatelessModel
				messageStatelessOptionsModel.UserID = core.StringPtr("testString")
				messageStatelessOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.MessageStatelessWithContext(ctx, messageStatelessOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.MessageStateless(messageStatelessOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.MessageStatelessWithContext(ctx, messageStatelessOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(messageStatelessPath))
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
					fmt.Fprintf(res, "%s", `{"output": {"generic": [{"response_type": "option", "title": "Title", "description": "Description", "preference": "dropdown", "options": [{"label": "Label", "value": {"input": {"message_type": "text", "text": "Text", "intents": [{"intent": "Intent", "confidence": 10}], "entities": [{"entity": "Entity", "location": [8], "value": "Value", "confidence": 10, "metadata": {"mapKey": "anyValue"}, "groups": [{"group": "Group", "location": [8]}], "interpretation": {"calendar_type": "CalendarType", "datetime_link": "DatetimeLink", "festival": "Festival", "granularity": "day", "range_link": "RangeLink", "range_modifier": "RangeModifier", "relative_day": 11, "relative_month": 13, "relative_week": 12, "relative_weekend": 15, "relative_year": 12, "specific_day": 11, "specific_day_of_week": "SpecificDayOfWeek", "specific_month": 13, "specific_quarter": 15, "specific_year": 12, "numeric_value": 12, "subtype": "Subtype", "part_of_day": "PartOfDay", "relative_hour": 12, "relative_minute": 14, "relative_second": 14, "specific_hour": 12, "specific_minute": 14, "specific_second": 14, "timezone": "Timezone"}, "alternatives": [{"value": "Value", "confidence": 10}], "role": {"type": "date_from"}}], "suggestion_id": "SuggestionID", "options": {"restart": false, "alternate_intents": true, "spelling": {"suggestions": false, "auto_correct": false}, "debug": false, "return_context": false, "export": true}}}}], "channels": [{"channel": "Channel"}]}], "intents": [{"intent": "Intent", "confidence": 10}], "entities": [{"entity": "Entity", "location": [8], "value": "Value", "confidence": 10, "metadata": {"mapKey": "anyValue"}, "groups": [{"group": "Group", "location": [8]}], "interpretation": {"calendar_type": "CalendarType", "datetime_link": "DatetimeLink", "festival": "Festival", "granularity": "day", "range_link": "RangeLink", "range_modifier": "RangeModifier", "relative_day": 11, "relative_month": 13, "relative_week": 12, "relative_weekend": 15, "relative_year": 12, "specific_day": 11, "specific_day_of_week": "SpecificDayOfWeek", "specific_month": 13, "specific_quarter": 15, "specific_year": 12, "numeric_value": 12, "subtype": "Subtype", "part_of_day": "PartOfDay", "relative_hour": 12, "relative_minute": 14, "relative_second": 14, "specific_hour": 12, "specific_minute": 14, "specific_second": 14, "timezone": "Timezone"}, "alternatives": [{"value": "Value", "confidence": 10}], "role": {"type": "date_from"}}], "actions": [{"name": "Name", "type": "client", "parameters": {"mapKey": "anyValue"}, "result_variable": "ResultVariable", "credentials": "Credentials"}], "debug": {"nodes_visited": [{"dialog_node": "DialogNode", "title": "Title", "conditions": "Conditions"}], "log_messages": [{"level": "info", "message": "Message", "code": "Code", "source": {"type": "dialog_node", "dialog_node": "DialogNode"}}], "branch_exited": true, "branch_exited_reason": "completed"}, "user_defined": {"mapKey": "anyValue"}, "spelling": {"text": "Text", "original_text": "OriginalText", "suggested_text": "SuggestedText"}}, "context": {"global": {"system": {"timezone": "Timezone", "user_id": "UserID", "turn_count": 9, "locale": "en-us", "reference_time": "ReferenceTime"}, "session_id": "SessionID"}, "skills": {"mapKey": {"user_defined": {"mapKey": {"anyKey": "anyValue"}}, "system": {"state": "State"}}}}, "user_id": "UserID"}`)
				}))
			})
			It(`Invoke MessageStateless successfully`, func() {
				assistantService, serviceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.MessageStateless(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RuntimeIntent model
				runtimeIntentModel := new(assistantv2.RuntimeIntent)
				runtimeIntentModel.Intent = core.StringPtr("testString")
				runtimeIntentModel.Confidence = core.Float64Ptr(float64(72.5))

				// Construct an instance of the CaptureGroup model
				captureGroupModel := new(assistantv2.CaptureGroup)
				captureGroupModel.Group = core.StringPtr("testString")
				captureGroupModel.Location = []int64{int64(38)}

				// Construct an instance of the RuntimeEntityInterpretation model
				runtimeEntityInterpretationModel := new(assistantv2.RuntimeEntityInterpretation)
				runtimeEntityInterpretationModel.CalendarType = core.StringPtr("testString")
				runtimeEntityInterpretationModel.DatetimeLink = core.StringPtr("testString")
				runtimeEntityInterpretationModel.Festival = core.StringPtr("testString")
				runtimeEntityInterpretationModel.Granularity = core.StringPtr("day")
				runtimeEntityInterpretationModel.RangeLink = core.StringPtr("testString")
				runtimeEntityInterpretationModel.RangeModifier = core.StringPtr("testString")
				runtimeEntityInterpretationModel.RelativeDay = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeMonth = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeWeek = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeWeekend = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeYear = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificDay = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificDayOfWeek = core.StringPtr("testString")
				runtimeEntityInterpretationModel.SpecificMonth = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificQuarter = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificYear = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.NumericValue = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.Subtype = core.StringPtr("testString")
				runtimeEntityInterpretationModel.PartOfDay = core.StringPtr("testString")
				runtimeEntityInterpretationModel.RelativeHour = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeMinute = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeSecond = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificHour = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificMinute = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificSecond = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.Timezone = core.StringPtr("testString")

				// Construct an instance of the RuntimeEntityAlternative model
				runtimeEntityAlternativeModel := new(assistantv2.RuntimeEntityAlternative)
				runtimeEntityAlternativeModel.Value = core.StringPtr("testString")
				runtimeEntityAlternativeModel.Confidence = core.Float64Ptr(float64(72.5))

				// Construct an instance of the RuntimeEntityRole model
				runtimeEntityRoleModel := new(assistantv2.RuntimeEntityRole)
				runtimeEntityRoleModel.Type = core.StringPtr("date_from")

				// Construct an instance of the RuntimeEntity model
				runtimeEntityModel := new(assistantv2.RuntimeEntity)
				runtimeEntityModel.Entity = core.StringPtr("testString")
				runtimeEntityModel.Location = []int64{int64(38)}
				runtimeEntityModel.Value = core.StringPtr("testString")
				runtimeEntityModel.Confidence = core.Float64Ptr(float64(72.5))
				runtimeEntityModel.Metadata = make(map[string]interface{})
				runtimeEntityModel.Groups = []assistantv2.CaptureGroup{*captureGroupModel}
				runtimeEntityModel.Interpretation = runtimeEntityInterpretationModel
				runtimeEntityModel.Alternatives = []assistantv2.RuntimeEntityAlternative{*runtimeEntityAlternativeModel}
				runtimeEntityModel.Role = runtimeEntityRoleModel

				// Construct an instance of the MessageInputOptionsSpelling model
				messageInputOptionsSpellingModel := new(assistantv2.MessageInputOptionsSpelling)
				messageInputOptionsSpellingModel.Suggestions = core.BoolPtr(true)
				messageInputOptionsSpellingModel.AutoCorrect = core.BoolPtr(true)

				// Construct an instance of the MessageInputOptionsStateless model
				messageInputOptionsStatelessModel := new(assistantv2.MessageInputOptionsStateless)
				messageInputOptionsStatelessModel.Restart = core.BoolPtr(true)
				messageInputOptionsStatelessModel.AlternateIntents = core.BoolPtr(true)
				messageInputOptionsStatelessModel.Spelling = messageInputOptionsSpellingModel
				messageInputOptionsStatelessModel.Debug = core.BoolPtr(true)

				// Construct an instance of the MessageInputStateless model
				messageInputStatelessModel := new(assistantv2.MessageInputStateless)
				messageInputStatelessModel.MessageType = core.StringPtr("text")
				messageInputStatelessModel.Text = core.StringPtr("testString")
				messageInputStatelessModel.Intents = []assistantv2.RuntimeIntent{*runtimeIntentModel}
				messageInputStatelessModel.Entities = []assistantv2.RuntimeEntity{*runtimeEntityModel}
				messageInputStatelessModel.SuggestionID = core.StringPtr("testString")
				messageInputStatelessModel.Options = messageInputOptionsStatelessModel

				// Construct an instance of the MessageContextGlobalSystem model
				messageContextGlobalSystemModel := new(assistantv2.MessageContextGlobalSystem)
				messageContextGlobalSystemModel.Timezone = core.StringPtr("testString")
				messageContextGlobalSystemModel.UserID = core.StringPtr("testString")
				messageContextGlobalSystemModel.TurnCount = core.Int64Ptr(int64(38))
				messageContextGlobalSystemModel.Locale = core.StringPtr("en-us")
				messageContextGlobalSystemModel.ReferenceTime = core.StringPtr("testString")

				// Construct an instance of the MessageContextGlobalStateless model
				messageContextGlobalStatelessModel := new(assistantv2.MessageContextGlobalStateless)
				messageContextGlobalStatelessModel.System = messageContextGlobalSystemModel
				messageContextGlobalStatelessModel.SessionID = core.StringPtr("testString")

				// Construct an instance of the MessageContextSkillSystem model
				messageContextSkillSystemModel := new(assistantv2.MessageContextSkillSystem)
				messageContextSkillSystemModel.State = core.StringPtr("testString")
				messageContextSkillSystemModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the MessageContextSkill model
				messageContextSkillModel := new(assistantv2.MessageContextSkill)
				messageContextSkillModel.UserDefined = make(map[string]interface{})
				messageContextSkillModel.System = messageContextSkillSystemModel

				// Construct an instance of the MessageContextStateless model
				messageContextStatelessModel := new(assistantv2.MessageContextStateless)
				messageContextStatelessModel.Global = messageContextGlobalStatelessModel
				messageContextStatelessModel.Skills = make(map[string]assistantv2.MessageContextSkill)
				messageContextStatelessModel.Skills["foo"] = *messageContextSkillModel

				// Construct an instance of the MessageStatelessOptions model
				messageStatelessOptionsModel := new(assistantv2.MessageStatelessOptions)
				messageStatelessOptionsModel.AssistantID = core.StringPtr("testString")
				messageStatelessOptionsModel.Input = messageInputStatelessModel
				messageStatelessOptionsModel.Context = messageContextStatelessModel
				messageStatelessOptionsModel.UserID = core.StringPtr("testString")
				messageStatelessOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.MessageStateless(messageStatelessOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke MessageStateless with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the RuntimeIntent model
				runtimeIntentModel := new(assistantv2.RuntimeIntent)
				runtimeIntentModel.Intent = core.StringPtr("testString")
				runtimeIntentModel.Confidence = core.Float64Ptr(float64(72.5))

				// Construct an instance of the CaptureGroup model
				captureGroupModel := new(assistantv2.CaptureGroup)
				captureGroupModel.Group = core.StringPtr("testString")
				captureGroupModel.Location = []int64{int64(38)}

				// Construct an instance of the RuntimeEntityInterpretation model
				runtimeEntityInterpretationModel := new(assistantv2.RuntimeEntityInterpretation)
				runtimeEntityInterpretationModel.CalendarType = core.StringPtr("testString")
				runtimeEntityInterpretationModel.DatetimeLink = core.StringPtr("testString")
				runtimeEntityInterpretationModel.Festival = core.StringPtr("testString")
				runtimeEntityInterpretationModel.Granularity = core.StringPtr("day")
				runtimeEntityInterpretationModel.RangeLink = core.StringPtr("testString")
				runtimeEntityInterpretationModel.RangeModifier = core.StringPtr("testString")
				runtimeEntityInterpretationModel.RelativeDay = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeMonth = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeWeek = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeWeekend = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeYear = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificDay = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificDayOfWeek = core.StringPtr("testString")
				runtimeEntityInterpretationModel.SpecificMonth = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificQuarter = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificYear = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.NumericValue = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.Subtype = core.StringPtr("testString")
				runtimeEntityInterpretationModel.PartOfDay = core.StringPtr("testString")
				runtimeEntityInterpretationModel.RelativeHour = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeMinute = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeSecond = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificHour = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificMinute = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificSecond = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.Timezone = core.StringPtr("testString")

				// Construct an instance of the RuntimeEntityAlternative model
				runtimeEntityAlternativeModel := new(assistantv2.RuntimeEntityAlternative)
				runtimeEntityAlternativeModel.Value = core.StringPtr("testString")
				runtimeEntityAlternativeModel.Confidence = core.Float64Ptr(float64(72.5))

				// Construct an instance of the RuntimeEntityRole model
				runtimeEntityRoleModel := new(assistantv2.RuntimeEntityRole)
				runtimeEntityRoleModel.Type = core.StringPtr("date_from")

				// Construct an instance of the RuntimeEntity model
				runtimeEntityModel := new(assistantv2.RuntimeEntity)
				runtimeEntityModel.Entity = core.StringPtr("testString")
				runtimeEntityModel.Location = []int64{int64(38)}
				runtimeEntityModel.Value = core.StringPtr("testString")
				runtimeEntityModel.Confidence = core.Float64Ptr(float64(72.5))
				runtimeEntityModel.Metadata = make(map[string]interface{})
				runtimeEntityModel.Groups = []assistantv2.CaptureGroup{*captureGroupModel}
				runtimeEntityModel.Interpretation = runtimeEntityInterpretationModel
				runtimeEntityModel.Alternatives = []assistantv2.RuntimeEntityAlternative{*runtimeEntityAlternativeModel}
				runtimeEntityModel.Role = runtimeEntityRoleModel

				// Construct an instance of the MessageInputOptionsSpelling model
				messageInputOptionsSpellingModel := new(assistantv2.MessageInputOptionsSpelling)
				messageInputOptionsSpellingModel.Suggestions = core.BoolPtr(true)
				messageInputOptionsSpellingModel.AutoCorrect = core.BoolPtr(true)

				// Construct an instance of the MessageInputOptionsStateless model
				messageInputOptionsStatelessModel := new(assistantv2.MessageInputOptionsStateless)
				messageInputOptionsStatelessModel.Restart = core.BoolPtr(true)
				messageInputOptionsStatelessModel.AlternateIntents = core.BoolPtr(true)
				messageInputOptionsStatelessModel.Spelling = messageInputOptionsSpellingModel
				messageInputOptionsStatelessModel.Debug = core.BoolPtr(true)

				// Construct an instance of the MessageInputStateless model
				messageInputStatelessModel := new(assistantv2.MessageInputStateless)
				messageInputStatelessModel.MessageType = core.StringPtr("text")
				messageInputStatelessModel.Text = core.StringPtr("testString")
				messageInputStatelessModel.Intents = []assistantv2.RuntimeIntent{*runtimeIntentModel}
				messageInputStatelessModel.Entities = []assistantv2.RuntimeEntity{*runtimeEntityModel}
				messageInputStatelessModel.SuggestionID = core.StringPtr("testString")
				messageInputStatelessModel.Options = messageInputOptionsStatelessModel

				// Construct an instance of the MessageContextGlobalSystem model
				messageContextGlobalSystemModel := new(assistantv2.MessageContextGlobalSystem)
				messageContextGlobalSystemModel.Timezone = core.StringPtr("testString")
				messageContextGlobalSystemModel.UserID = core.StringPtr("testString")
				messageContextGlobalSystemModel.TurnCount = core.Int64Ptr(int64(38))
				messageContextGlobalSystemModel.Locale = core.StringPtr("en-us")
				messageContextGlobalSystemModel.ReferenceTime = core.StringPtr("testString")

				// Construct an instance of the MessageContextGlobalStateless model
				messageContextGlobalStatelessModel := new(assistantv2.MessageContextGlobalStateless)
				messageContextGlobalStatelessModel.System = messageContextGlobalSystemModel
				messageContextGlobalStatelessModel.SessionID = core.StringPtr("testString")

				// Construct an instance of the MessageContextSkillSystem model
				messageContextSkillSystemModel := new(assistantv2.MessageContextSkillSystem)
				messageContextSkillSystemModel.State = core.StringPtr("testString")
				messageContextSkillSystemModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the MessageContextSkill model
				messageContextSkillModel := new(assistantv2.MessageContextSkill)
				messageContextSkillModel.UserDefined = make(map[string]interface{})
				messageContextSkillModel.System = messageContextSkillSystemModel

				// Construct an instance of the MessageContextStateless model
				messageContextStatelessModel := new(assistantv2.MessageContextStateless)
				messageContextStatelessModel.Global = messageContextGlobalStatelessModel
				messageContextStatelessModel.Skills = make(map[string]assistantv2.MessageContextSkill)
				messageContextStatelessModel.Skills["foo"] = *messageContextSkillModel

				// Construct an instance of the MessageStatelessOptions model
				messageStatelessOptionsModel := new(assistantv2.MessageStatelessOptions)
				messageStatelessOptionsModel.AssistantID = core.StringPtr("testString")
				messageStatelessOptionsModel.Input = messageInputStatelessModel
				messageStatelessOptionsModel.Context = messageContextStatelessModel
				messageStatelessOptionsModel.UserID = core.StringPtr("testString")
				messageStatelessOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.MessageStateless(messageStatelessOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the MessageStatelessOptions model with no property values
				messageStatelessOptionsModelNew := new(assistantv2.MessageStatelessOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = assistantService.MessageStateless(messageStatelessOptionsModelNew)
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
			It(`Invoke MessageStateless successfully`, func() {
				assistantService, serviceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the RuntimeIntent model
				runtimeIntentModel := new(assistantv2.RuntimeIntent)
				runtimeIntentModel.Intent = core.StringPtr("testString")
				runtimeIntentModel.Confidence = core.Float64Ptr(float64(72.5))

				// Construct an instance of the CaptureGroup model
				captureGroupModel := new(assistantv2.CaptureGroup)
				captureGroupModel.Group = core.StringPtr("testString")
				captureGroupModel.Location = []int64{int64(38)}

				// Construct an instance of the RuntimeEntityInterpretation model
				runtimeEntityInterpretationModel := new(assistantv2.RuntimeEntityInterpretation)
				runtimeEntityInterpretationModel.CalendarType = core.StringPtr("testString")
				runtimeEntityInterpretationModel.DatetimeLink = core.StringPtr("testString")
				runtimeEntityInterpretationModel.Festival = core.StringPtr("testString")
				runtimeEntityInterpretationModel.Granularity = core.StringPtr("day")
				runtimeEntityInterpretationModel.RangeLink = core.StringPtr("testString")
				runtimeEntityInterpretationModel.RangeModifier = core.StringPtr("testString")
				runtimeEntityInterpretationModel.RelativeDay = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeMonth = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeWeek = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeWeekend = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeYear = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificDay = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificDayOfWeek = core.StringPtr("testString")
				runtimeEntityInterpretationModel.SpecificMonth = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificQuarter = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificYear = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.NumericValue = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.Subtype = core.StringPtr("testString")
				runtimeEntityInterpretationModel.PartOfDay = core.StringPtr("testString")
				runtimeEntityInterpretationModel.RelativeHour = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeMinute = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeSecond = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificHour = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificMinute = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificSecond = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.Timezone = core.StringPtr("testString")

				// Construct an instance of the RuntimeEntityAlternative model
				runtimeEntityAlternativeModel := new(assistantv2.RuntimeEntityAlternative)
				runtimeEntityAlternativeModel.Value = core.StringPtr("testString")
				runtimeEntityAlternativeModel.Confidence = core.Float64Ptr(float64(72.5))

				// Construct an instance of the RuntimeEntityRole model
				runtimeEntityRoleModel := new(assistantv2.RuntimeEntityRole)
				runtimeEntityRoleModel.Type = core.StringPtr("date_from")

				// Construct an instance of the RuntimeEntity model
				runtimeEntityModel := new(assistantv2.RuntimeEntity)
				runtimeEntityModel.Entity = core.StringPtr("testString")
				runtimeEntityModel.Location = []int64{int64(38)}
				runtimeEntityModel.Value = core.StringPtr("testString")
				runtimeEntityModel.Confidence = core.Float64Ptr(float64(72.5))
				runtimeEntityModel.Metadata = make(map[string]interface{})
				runtimeEntityModel.Groups = []assistantv2.CaptureGroup{*captureGroupModel}
				runtimeEntityModel.Interpretation = runtimeEntityInterpretationModel
				runtimeEntityModel.Alternatives = []assistantv2.RuntimeEntityAlternative{*runtimeEntityAlternativeModel}
				runtimeEntityModel.Role = runtimeEntityRoleModel

				// Construct an instance of the MessageInputOptionsSpelling model
				messageInputOptionsSpellingModel := new(assistantv2.MessageInputOptionsSpelling)
				messageInputOptionsSpellingModel.Suggestions = core.BoolPtr(true)
				messageInputOptionsSpellingModel.AutoCorrect = core.BoolPtr(true)

				// Construct an instance of the MessageInputOptionsStateless model
				messageInputOptionsStatelessModel := new(assistantv2.MessageInputOptionsStateless)
				messageInputOptionsStatelessModel.Restart = core.BoolPtr(true)
				messageInputOptionsStatelessModel.AlternateIntents = core.BoolPtr(true)
				messageInputOptionsStatelessModel.Spelling = messageInputOptionsSpellingModel
				messageInputOptionsStatelessModel.Debug = core.BoolPtr(true)

				// Construct an instance of the MessageInputStateless model
				messageInputStatelessModel := new(assistantv2.MessageInputStateless)
				messageInputStatelessModel.MessageType = core.StringPtr("text")
				messageInputStatelessModel.Text = core.StringPtr("testString")
				messageInputStatelessModel.Intents = []assistantv2.RuntimeIntent{*runtimeIntentModel}
				messageInputStatelessModel.Entities = []assistantv2.RuntimeEntity{*runtimeEntityModel}
				messageInputStatelessModel.SuggestionID = core.StringPtr("testString")
				messageInputStatelessModel.Options = messageInputOptionsStatelessModel

				// Construct an instance of the MessageContextGlobalSystem model
				messageContextGlobalSystemModel := new(assistantv2.MessageContextGlobalSystem)
				messageContextGlobalSystemModel.Timezone = core.StringPtr("testString")
				messageContextGlobalSystemModel.UserID = core.StringPtr("testString")
				messageContextGlobalSystemModel.TurnCount = core.Int64Ptr(int64(38))
				messageContextGlobalSystemModel.Locale = core.StringPtr("en-us")
				messageContextGlobalSystemModel.ReferenceTime = core.StringPtr("testString")

				// Construct an instance of the MessageContextGlobalStateless model
				messageContextGlobalStatelessModel := new(assistantv2.MessageContextGlobalStateless)
				messageContextGlobalStatelessModel.System = messageContextGlobalSystemModel
				messageContextGlobalStatelessModel.SessionID = core.StringPtr("testString")

				// Construct an instance of the MessageContextSkillSystem model
				messageContextSkillSystemModel := new(assistantv2.MessageContextSkillSystem)
				messageContextSkillSystemModel.State = core.StringPtr("testString")
				messageContextSkillSystemModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the MessageContextSkill model
				messageContextSkillModel := new(assistantv2.MessageContextSkill)
				messageContextSkillModel.UserDefined = make(map[string]interface{})
				messageContextSkillModel.System = messageContextSkillSystemModel

				// Construct an instance of the MessageContextStateless model
				messageContextStatelessModel := new(assistantv2.MessageContextStateless)
				messageContextStatelessModel.Global = messageContextGlobalStatelessModel
				messageContextStatelessModel.Skills = make(map[string]assistantv2.MessageContextSkill)
				messageContextStatelessModel.Skills["foo"] = *messageContextSkillModel

				// Construct an instance of the MessageStatelessOptions model
				messageStatelessOptionsModel := new(assistantv2.MessageStatelessOptions)
				messageStatelessOptionsModel.AssistantID = core.StringPtr("testString")
				messageStatelessOptionsModel.Input = messageInputStatelessModel
				messageStatelessOptionsModel.Context = messageContextStatelessModel
				messageStatelessOptionsModel.UserID = core.StringPtr("testString")
				messageStatelessOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.MessageStateless(messageStatelessOptionsModel)
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
	Describe(`BulkClassify(bulkClassifyOptions *BulkClassifyOptions) - Operation response error`, func() {
		version := "testString"
		bulkClassifyPath := "/v2/skills/testString/workspace/bulk_classify"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(bulkClassifyPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke BulkClassify with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the BulkClassifyUtterance model
				bulkClassifyUtteranceModel := new(assistantv2.BulkClassifyUtterance)
				bulkClassifyUtteranceModel.Text = core.StringPtr("testString")

				// Construct an instance of the BulkClassifyOptions model
				bulkClassifyOptionsModel := new(assistantv2.BulkClassifyOptions)
				bulkClassifyOptionsModel.SkillID = core.StringPtr("testString")
				bulkClassifyOptionsModel.Input = []assistantv2.BulkClassifyUtterance{*bulkClassifyUtteranceModel}
				bulkClassifyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.BulkClassify(bulkClassifyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.BulkClassify(bulkClassifyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`BulkClassify(bulkClassifyOptions *BulkClassifyOptions)`, func() {
		version := "testString"
		bulkClassifyPath := "/v2/skills/testString/workspace/bulk_classify"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(bulkClassifyPath))
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
					fmt.Fprintf(res, "%s", `{"output": [{"input": {"text": "Text"}, "entities": [{"entity": "Entity", "location": [8], "value": "Value", "confidence": 10, "metadata": {"mapKey": "anyValue"}, "groups": [{"group": "Group", "location": [8]}], "interpretation": {"calendar_type": "CalendarType", "datetime_link": "DatetimeLink", "festival": "Festival", "granularity": "day", "range_link": "RangeLink", "range_modifier": "RangeModifier", "relative_day": 11, "relative_month": 13, "relative_week": 12, "relative_weekend": 15, "relative_year": 12, "specific_day": 11, "specific_day_of_week": "SpecificDayOfWeek", "specific_month": 13, "specific_quarter": 15, "specific_year": 12, "numeric_value": 12, "subtype": "Subtype", "part_of_day": "PartOfDay", "relative_hour": 12, "relative_minute": 14, "relative_second": 14, "specific_hour": 12, "specific_minute": 14, "specific_second": 14, "timezone": "Timezone"}, "alternatives": [{"value": "Value", "confidence": 10}], "role": {"type": "date_from"}}], "intents": [{"intent": "Intent", "confidence": 10}]}]}`)
				}))
			})
			It(`Invoke BulkClassify successfully with retries`, func() {
				assistantService, serviceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the BulkClassifyUtterance model
				bulkClassifyUtteranceModel := new(assistantv2.BulkClassifyUtterance)
				bulkClassifyUtteranceModel.Text = core.StringPtr("testString")

				// Construct an instance of the BulkClassifyOptions model
				bulkClassifyOptionsModel := new(assistantv2.BulkClassifyOptions)
				bulkClassifyOptionsModel.SkillID = core.StringPtr("testString")
				bulkClassifyOptionsModel.Input = []assistantv2.BulkClassifyUtterance{*bulkClassifyUtteranceModel}
				bulkClassifyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.BulkClassifyWithContext(ctx, bulkClassifyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.BulkClassify(bulkClassifyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.BulkClassifyWithContext(ctx, bulkClassifyOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(bulkClassifyPath))
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
					fmt.Fprintf(res, "%s", `{"output": [{"input": {"text": "Text"}, "entities": [{"entity": "Entity", "location": [8], "value": "Value", "confidence": 10, "metadata": {"mapKey": "anyValue"}, "groups": [{"group": "Group", "location": [8]}], "interpretation": {"calendar_type": "CalendarType", "datetime_link": "DatetimeLink", "festival": "Festival", "granularity": "day", "range_link": "RangeLink", "range_modifier": "RangeModifier", "relative_day": 11, "relative_month": 13, "relative_week": 12, "relative_weekend": 15, "relative_year": 12, "specific_day": 11, "specific_day_of_week": "SpecificDayOfWeek", "specific_month": 13, "specific_quarter": 15, "specific_year": 12, "numeric_value": 12, "subtype": "Subtype", "part_of_day": "PartOfDay", "relative_hour": 12, "relative_minute": 14, "relative_second": 14, "specific_hour": 12, "specific_minute": 14, "specific_second": 14, "timezone": "Timezone"}, "alternatives": [{"value": "Value", "confidence": 10}], "role": {"type": "date_from"}}], "intents": [{"intent": "Intent", "confidence": 10}]}]}`)
				}))
			})
			It(`Invoke BulkClassify successfully`, func() {
				assistantService, serviceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.BulkClassify(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the BulkClassifyUtterance model
				bulkClassifyUtteranceModel := new(assistantv2.BulkClassifyUtterance)
				bulkClassifyUtteranceModel.Text = core.StringPtr("testString")

				// Construct an instance of the BulkClassifyOptions model
				bulkClassifyOptionsModel := new(assistantv2.BulkClassifyOptions)
				bulkClassifyOptionsModel.SkillID = core.StringPtr("testString")
				bulkClassifyOptionsModel.Input = []assistantv2.BulkClassifyUtterance{*bulkClassifyUtteranceModel}
				bulkClassifyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.BulkClassify(bulkClassifyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke BulkClassify with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the BulkClassifyUtterance model
				bulkClassifyUtteranceModel := new(assistantv2.BulkClassifyUtterance)
				bulkClassifyUtteranceModel.Text = core.StringPtr("testString")

				// Construct an instance of the BulkClassifyOptions model
				bulkClassifyOptionsModel := new(assistantv2.BulkClassifyOptions)
				bulkClassifyOptionsModel.SkillID = core.StringPtr("testString")
				bulkClassifyOptionsModel.Input = []assistantv2.BulkClassifyUtterance{*bulkClassifyUtteranceModel}
				bulkClassifyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.BulkClassify(bulkClassifyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the BulkClassifyOptions model with no property values
				bulkClassifyOptionsModelNew := new(assistantv2.BulkClassifyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = assistantService.BulkClassify(bulkClassifyOptionsModelNew)
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
			It(`Invoke BulkClassify successfully`, func() {
				assistantService, serviceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the BulkClassifyUtterance model
				bulkClassifyUtteranceModel := new(assistantv2.BulkClassifyUtterance)
				bulkClassifyUtteranceModel.Text = core.StringPtr("testString")

				// Construct an instance of the BulkClassifyOptions model
				bulkClassifyOptionsModel := new(assistantv2.BulkClassifyOptions)
				bulkClassifyOptionsModel.SkillID = core.StringPtr("testString")
				bulkClassifyOptionsModel.Input = []assistantv2.BulkClassifyUtterance{*bulkClassifyUtteranceModel}
				bulkClassifyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.BulkClassify(bulkClassifyOptionsModel)
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
	Describe(`ListLogs(listLogsOptions *ListLogsOptions) - Operation response error`, func() {
		version := "testString"
		listLogsPath := "/v2/assistants/testString/logs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listLogsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["filter"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["page_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["cursor"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListLogs with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ListLogsOptions model
				listLogsOptionsModel := new(assistantv2.ListLogsOptions)
				listLogsOptionsModel.AssistantID = core.StringPtr("testString")
				listLogsOptionsModel.Sort = core.StringPtr("testString")
				listLogsOptionsModel.Filter = core.StringPtr("testString")
				listLogsOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listLogsOptionsModel.Cursor = core.StringPtr("testString")
				listLogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.ListLogs(listLogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.ListLogs(listLogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListLogs(listLogsOptions *ListLogsOptions)`, func() {
		version := "testString"
		listLogsPath := "/v2/assistants/testString/logs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listLogsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["filter"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["page_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["cursor"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"logs": [{"log_id": "LogID", "request": {"input": {"message_type": "text", "text": "Text", "intents": [{"intent": "Intent", "confidence": 10}], "entities": [{"entity": "Entity", "location": [8], "value": "Value", "confidence": 10, "metadata": {"mapKey": "anyValue"}, "groups": [{"group": "Group", "location": [8]}], "interpretation": {"calendar_type": "CalendarType", "datetime_link": "DatetimeLink", "festival": "Festival", "granularity": "day", "range_link": "RangeLink", "range_modifier": "RangeModifier", "relative_day": 11, "relative_month": 13, "relative_week": 12, "relative_weekend": 15, "relative_year": 12, "specific_day": 11, "specific_day_of_week": "SpecificDayOfWeek", "specific_month": 13, "specific_quarter": 15, "specific_year": 12, "numeric_value": 12, "subtype": "Subtype", "part_of_day": "PartOfDay", "relative_hour": 12, "relative_minute": 14, "relative_second": 14, "specific_hour": 12, "specific_minute": 14, "specific_second": 14, "timezone": "Timezone"}, "alternatives": [{"value": "Value", "confidence": 10}], "role": {"type": "date_from"}}], "suggestion_id": "SuggestionID", "options": {"restart": false, "alternate_intents": true, "spelling": {"suggestions": false, "auto_correct": false}, "debug": false, "return_context": false, "export": true}}, "context": {"global": {"system": {"timezone": "Timezone", "user_id": "UserID", "turn_count": 9, "locale": "en-us", "reference_time": "ReferenceTime"}, "session_id": "SessionID"}, "skills": {"mapKey": {"user_defined": {"mapKey": {"anyKey": "anyValue"}}, "system": {"state": "State"}}}}, "user_id": "UserID"}, "response": {"output": {"generic": [{"response_type": "option", "title": "Title", "description": "Description", "preference": "dropdown", "options": [{"label": "Label", "value": {"input": {"message_type": "text", "text": "Text", "intents": [{"intent": "Intent", "confidence": 10}], "entities": [{"entity": "Entity", "location": [8], "value": "Value", "confidence": 10, "metadata": {"mapKey": "anyValue"}, "groups": [{"group": "Group", "location": [8]}], "interpretation": {"calendar_type": "CalendarType", "datetime_link": "DatetimeLink", "festival": "Festival", "granularity": "day", "range_link": "RangeLink", "range_modifier": "RangeModifier", "relative_day": 11, "relative_month": 13, "relative_week": 12, "relative_weekend": 15, "relative_year": 12, "specific_day": 11, "specific_day_of_week": "SpecificDayOfWeek", "specific_month": 13, "specific_quarter": 15, "specific_year": 12, "numeric_value": 12, "subtype": "Subtype", "part_of_day": "PartOfDay", "relative_hour": 12, "relative_minute": 14, "relative_second": 14, "specific_hour": 12, "specific_minute": 14, "specific_second": 14, "timezone": "Timezone"}, "alternatives": [{"value": "Value", "confidence": 10}], "role": {"type": "date_from"}}], "suggestion_id": "SuggestionID", "options": {"restart": false, "alternate_intents": true, "spelling": {"suggestions": false, "auto_correct": false}, "debug": false, "return_context": false, "export": true}}}}], "channels": [{"channel": "Channel"}]}], "intents": [{"intent": "Intent", "confidence": 10}], "entities": [{"entity": "Entity", "location": [8], "value": "Value", "confidence": 10, "metadata": {"mapKey": "anyValue"}, "groups": [{"group": "Group", "location": [8]}], "interpretation": {"calendar_type": "CalendarType", "datetime_link": "DatetimeLink", "festival": "Festival", "granularity": "day", "range_link": "RangeLink", "range_modifier": "RangeModifier", "relative_day": 11, "relative_month": 13, "relative_week": 12, "relative_weekend": 15, "relative_year": 12, "specific_day": 11, "specific_day_of_week": "SpecificDayOfWeek", "specific_month": 13, "specific_quarter": 15, "specific_year": 12, "numeric_value": 12, "subtype": "Subtype", "part_of_day": "PartOfDay", "relative_hour": 12, "relative_minute": 14, "relative_second": 14, "specific_hour": 12, "specific_minute": 14, "specific_second": 14, "timezone": "Timezone"}, "alternatives": [{"value": "Value", "confidence": 10}], "role": {"type": "date_from"}}], "actions": [{"name": "Name", "type": "client", "parameters": {"mapKey": "anyValue"}, "result_variable": "ResultVariable", "credentials": "Credentials"}], "debug": {"nodes_visited": [{"dialog_node": "DialogNode", "title": "Title", "conditions": "Conditions"}], "log_messages": [{"level": "info", "message": "Message", "code": "Code", "source": {"type": "dialog_node", "dialog_node": "DialogNode"}}], "branch_exited": true, "branch_exited_reason": "completed"}, "user_defined": {"mapKey": "anyValue"}, "spelling": {"text": "Text", "original_text": "OriginalText", "suggested_text": "SuggestedText"}}, "context": {"global": {"system": {"timezone": "Timezone", "user_id": "UserID", "turn_count": 9, "locale": "en-us", "reference_time": "ReferenceTime"}, "session_id": "SessionID"}, "skills": {"mapKey": {"user_defined": {"mapKey": {"anyKey": "anyValue"}}, "system": {"state": "State"}}}}, "user_id": "UserID"}, "assistant_id": "AssistantID", "session_id": "SessionID", "skill_id": "SkillID", "snapshot": "Snapshot", "request_timestamp": "RequestTimestamp", "response_timestamp": "ResponseTimestamp", "language": "Language", "customer_id": "CustomerID"}], "pagination": {"next_url": "NextURL", "matched": 7, "next_cursor": "NextCursor"}}`)
				}))
			})
			It(`Invoke ListLogs successfully with retries`, func() {
				assistantService, serviceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the ListLogsOptions model
				listLogsOptionsModel := new(assistantv2.ListLogsOptions)
				listLogsOptionsModel.AssistantID = core.StringPtr("testString")
				listLogsOptionsModel.Sort = core.StringPtr("testString")
				listLogsOptionsModel.Filter = core.StringPtr("testString")
				listLogsOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listLogsOptionsModel.Cursor = core.StringPtr("testString")
				listLogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.ListLogsWithContext(ctx, listLogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.ListLogs(listLogsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.ListLogsWithContext(ctx, listLogsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listLogsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["filter"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["page_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["cursor"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"logs": [{"log_id": "LogID", "request": {"input": {"message_type": "text", "text": "Text", "intents": [{"intent": "Intent", "confidence": 10}], "entities": [{"entity": "Entity", "location": [8], "value": "Value", "confidence": 10, "metadata": {"mapKey": "anyValue"}, "groups": [{"group": "Group", "location": [8]}], "interpretation": {"calendar_type": "CalendarType", "datetime_link": "DatetimeLink", "festival": "Festival", "granularity": "day", "range_link": "RangeLink", "range_modifier": "RangeModifier", "relative_day": 11, "relative_month": 13, "relative_week": 12, "relative_weekend": 15, "relative_year": 12, "specific_day": 11, "specific_day_of_week": "SpecificDayOfWeek", "specific_month": 13, "specific_quarter": 15, "specific_year": 12, "numeric_value": 12, "subtype": "Subtype", "part_of_day": "PartOfDay", "relative_hour": 12, "relative_minute": 14, "relative_second": 14, "specific_hour": 12, "specific_minute": 14, "specific_second": 14, "timezone": "Timezone"}, "alternatives": [{"value": "Value", "confidence": 10}], "role": {"type": "date_from"}}], "suggestion_id": "SuggestionID", "options": {"restart": false, "alternate_intents": true, "spelling": {"suggestions": false, "auto_correct": false}, "debug": false, "return_context": false, "export": true}}, "context": {"global": {"system": {"timezone": "Timezone", "user_id": "UserID", "turn_count": 9, "locale": "en-us", "reference_time": "ReferenceTime"}, "session_id": "SessionID"}, "skills": {"mapKey": {"user_defined": {"mapKey": {"anyKey": "anyValue"}}, "system": {"state": "State"}}}}, "user_id": "UserID"}, "response": {"output": {"generic": [{"response_type": "option", "title": "Title", "description": "Description", "preference": "dropdown", "options": [{"label": "Label", "value": {"input": {"message_type": "text", "text": "Text", "intents": [{"intent": "Intent", "confidence": 10}], "entities": [{"entity": "Entity", "location": [8], "value": "Value", "confidence": 10, "metadata": {"mapKey": "anyValue"}, "groups": [{"group": "Group", "location": [8]}], "interpretation": {"calendar_type": "CalendarType", "datetime_link": "DatetimeLink", "festival": "Festival", "granularity": "day", "range_link": "RangeLink", "range_modifier": "RangeModifier", "relative_day": 11, "relative_month": 13, "relative_week": 12, "relative_weekend": 15, "relative_year": 12, "specific_day": 11, "specific_day_of_week": "SpecificDayOfWeek", "specific_month": 13, "specific_quarter": 15, "specific_year": 12, "numeric_value": 12, "subtype": "Subtype", "part_of_day": "PartOfDay", "relative_hour": 12, "relative_minute": 14, "relative_second": 14, "specific_hour": 12, "specific_minute": 14, "specific_second": 14, "timezone": "Timezone"}, "alternatives": [{"value": "Value", "confidence": 10}], "role": {"type": "date_from"}}], "suggestion_id": "SuggestionID", "options": {"restart": false, "alternate_intents": true, "spelling": {"suggestions": false, "auto_correct": false}, "debug": false, "return_context": false, "export": true}}}}], "channels": [{"channel": "Channel"}]}], "intents": [{"intent": "Intent", "confidence": 10}], "entities": [{"entity": "Entity", "location": [8], "value": "Value", "confidence": 10, "metadata": {"mapKey": "anyValue"}, "groups": [{"group": "Group", "location": [8]}], "interpretation": {"calendar_type": "CalendarType", "datetime_link": "DatetimeLink", "festival": "Festival", "granularity": "day", "range_link": "RangeLink", "range_modifier": "RangeModifier", "relative_day": 11, "relative_month": 13, "relative_week": 12, "relative_weekend": 15, "relative_year": 12, "specific_day": 11, "specific_day_of_week": "SpecificDayOfWeek", "specific_month": 13, "specific_quarter": 15, "specific_year": 12, "numeric_value": 12, "subtype": "Subtype", "part_of_day": "PartOfDay", "relative_hour": 12, "relative_minute": 14, "relative_second": 14, "specific_hour": 12, "specific_minute": 14, "specific_second": 14, "timezone": "Timezone"}, "alternatives": [{"value": "Value", "confidence": 10}], "role": {"type": "date_from"}}], "actions": [{"name": "Name", "type": "client", "parameters": {"mapKey": "anyValue"}, "result_variable": "ResultVariable", "credentials": "Credentials"}], "debug": {"nodes_visited": [{"dialog_node": "DialogNode", "title": "Title", "conditions": "Conditions"}], "log_messages": [{"level": "info", "message": "Message", "code": "Code", "source": {"type": "dialog_node", "dialog_node": "DialogNode"}}], "branch_exited": true, "branch_exited_reason": "completed"}, "user_defined": {"mapKey": "anyValue"}, "spelling": {"text": "Text", "original_text": "OriginalText", "suggested_text": "SuggestedText"}}, "context": {"global": {"system": {"timezone": "Timezone", "user_id": "UserID", "turn_count": 9, "locale": "en-us", "reference_time": "ReferenceTime"}, "session_id": "SessionID"}, "skills": {"mapKey": {"user_defined": {"mapKey": {"anyKey": "anyValue"}}, "system": {"state": "State"}}}}, "user_id": "UserID"}, "assistant_id": "AssistantID", "session_id": "SessionID", "skill_id": "SkillID", "snapshot": "Snapshot", "request_timestamp": "RequestTimestamp", "response_timestamp": "ResponseTimestamp", "language": "Language", "customer_id": "CustomerID"}], "pagination": {"next_url": "NextURL", "matched": 7, "next_cursor": "NextCursor"}}`)
				}))
			})
			It(`Invoke ListLogs successfully`, func() {
				assistantService, serviceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.ListLogs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListLogsOptions model
				listLogsOptionsModel := new(assistantv2.ListLogsOptions)
				listLogsOptionsModel.AssistantID = core.StringPtr("testString")
				listLogsOptionsModel.Sort = core.StringPtr("testString")
				listLogsOptionsModel.Filter = core.StringPtr("testString")
				listLogsOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listLogsOptionsModel.Cursor = core.StringPtr("testString")
				listLogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.ListLogs(listLogsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListLogs with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ListLogsOptions model
				listLogsOptionsModel := new(assistantv2.ListLogsOptions)
				listLogsOptionsModel.AssistantID = core.StringPtr("testString")
				listLogsOptionsModel.Sort = core.StringPtr("testString")
				listLogsOptionsModel.Filter = core.StringPtr("testString")
				listLogsOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listLogsOptionsModel.Cursor = core.StringPtr("testString")
				listLogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.ListLogs(listLogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListLogsOptions model with no property values
				listLogsOptionsModelNew := new(assistantv2.ListLogsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = assistantService.ListLogs(listLogsOptionsModelNew)
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
			It(`Invoke ListLogs successfully`, func() {
				assistantService, serviceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ListLogsOptions model
				listLogsOptionsModel := new(assistantv2.ListLogsOptions)
				listLogsOptionsModel.AssistantID = core.StringPtr("testString")
				listLogsOptionsModel.Sort = core.StringPtr("testString")
				listLogsOptionsModel.Filter = core.StringPtr("testString")
				listLogsOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listLogsOptionsModel.Cursor = core.StringPtr("testString")
				listLogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.ListLogs(listLogsOptionsModel)
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
	Describe(`DeleteUserData(deleteUserDataOptions *DeleteUserDataOptions)`, func() {
		version := "testString"
		deleteUserDataPath := "/v2/user_data"
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
				assistantService, serviceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := assistantService.DeleteUserData(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteUserDataOptions model
				deleteUserDataOptionsModel := new(assistantv2.DeleteUserDataOptions)
				deleteUserDataOptionsModel.CustomerID = core.StringPtr("testString")
				deleteUserDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = assistantService.DeleteUserData(deleteUserDataOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteUserData with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the DeleteUserDataOptions model
				deleteUserDataOptionsModel := new(assistantv2.DeleteUserDataOptions)
				deleteUserDataOptionsModel.CustomerID = core.StringPtr("testString")
				deleteUserDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := assistantService.DeleteUserData(deleteUserDataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteUserDataOptions model with no property values
				deleteUserDataOptionsModelNew := new(assistantv2.DeleteUserDataOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = assistantService.DeleteUserData(deleteUserDataOptionsModelNew)
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
			assistantService, _ := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
				URL:           "http://assistantv2modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       core.StringPtr(version),
			})
			It(`Invoke NewBulkClassifyOptions successfully`, func() {
				// Construct an instance of the BulkClassifyUtterance model
				bulkClassifyUtteranceModel := new(assistantv2.BulkClassifyUtterance)
				Expect(bulkClassifyUtteranceModel).ToNot(BeNil())
				bulkClassifyUtteranceModel.Text = core.StringPtr("testString")
				Expect(bulkClassifyUtteranceModel.Text).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the BulkClassifyOptions model
				skillID := "testString"
				bulkClassifyOptionsModel := assistantService.NewBulkClassifyOptions(skillID)
				bulkClassifyOptionsModel.SetSkillID("testString")
				bulkClassifyOptionsModel.SetInput([]assistantv2.BulkClassifyUtterance{*bulkClassifyUtteranceModel})
				bulkClassifyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(bulkClassifyOptionsModel).ToNot(BeNil())
				Expect(bulkClassifyOptionsModel.SkillID).To(Equal(core.StringPtr("testString")))
				Expect(bulkClassifyOptionsModel.Input).To(Equal([]assistantv2.BulkClassifyUtterance{*bulkClassifyUtteranceModel}))
				Expect(bulkClassifyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewBulkClassifyUtterance successfully`, func() {
				text := "testString"
				model, err := assistantService.NewBulkClassifyUtterance(text)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCaptureGroup successfully`, func() {
				group := "testString"
				model, err := assistantService.NewCaptureGroup(group)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCreateSessionOptions successfully`, func() {
				// Construct an instance of the CreateSessionOptions model
				assistantID := "testString"
				createSessionOptionsModel := assistantService.NewCreateSessionOptions(assistantID)
				createSessionOptionsModel.SetAssistantID("testString")
				createSessionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createSessionOptionsModel).ToNot(BeNil())
				Expect(createSessionOptionsModel.AssistantID).To(Equal(core.StringPtr("testString")))
				Expect(createSessionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteSessionOptions successfully`, func() {
				// Construct an instance of the DeleteSessionOptions model
				assistantID := "testString"
				sessionID := "testString"
				deleteSessionOptionsModel := assistantService.NewDeleteSessionOptions(assistantID, sessionID)
				deleteSessionOptionsModel.SetAssistantID("testString")
				deleteSessionOptionsModel.SetSessionID("testString")
				deleteSessionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteSessionOptionsModel).ToNot(BeNil())
				Expect(deleteSessionOptionsModel.AssistantID).To(Equal(core.StringPtr("testString")))
				Expect(deleteSessionOptionsModel.SessionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteSessionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteUserDataOptions successfully`, func() {
				// Construct an instance of the DeleteUserDataOptions model
				customerID := "testString"
				deleteUserDataOptionsModel := assistantService.NewDeleteUserDataOptions(customerID)
				deleteUserDataOptionsModel.SetCustomerID("testString")
				deleteUserDataOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteUserDataOptionsModel).ToNot(BeNil())
				Expect(deleteUserDataOptionsModel.CustomerID).To(Equal(core.StringPtr("testString")))
				Expect(deleteUserDataOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListLogsOptions successfully`, func() {
				// Construct an instance of the ListLogsOptions model
				assistantID := "testString"
				listLogsOptionsModel := assistantService.NewListLogsOptions(assistantID)
				listLogsOptionsModel.SetAssistantID("testString")
				listLogsOptionsModel.SetSort("testString")
				listLogsOptionsModel.SetFilter("testString")
				listLogsOptionsModel.SetPageLimit(int64(38))
				listLogsOptionsModel.SetCursor("testString")
				listLogsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listLogsOptionsModel).ToNot(BeNil())
				Expect(listLogsOptionsModel.AssistantID).To(Equal(core.StringPtr("testString")))
				Expect(listLogsOptionsModel.Sort).To(Equal(core.StringPtr("testString")))
				Expect(listLogsOptionsModel.Filter).To(Equal(core.StringPtr("testString")))
				Expect(listLogsOptionsModel.PageLimit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listLogsOptionsModel.Cursor).To(Equal(core.StringPtr("testString")))
				Expect(listLogsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewMessageOptions successfully`, func() {
				// Construct an instance of the RuntimeIntent model
				runtimeIntentModel := new(assistantv2.RuntimeIntent)
				Expect(runtimeIntentModel).ToNot(BeNil())
				runtimeIntentModel.Intent = core.StringPtr("testString")
				runtimeIntentModel.Confidence = core.Float64Ptr(float64(72.5))
				Expect(runtimeIntentModel.Intent).To(Equal(core.StringPtr("testString")))
				Expect(runtimeIntentModel.Confidence).To(Equal(core.Float64Ptr(float64(72.5))))

				// Construct an instance of the CaptureGroup model
				captureGroupModel := new(assistantv2.CaptureGroup)
				Expect(captureGroupModel).ToNot(BeNil())
				captureGroupModel.Group = core.StringPtr("testString")
				captureGroupModel.Location = []int64{int64(38)}
				Expect(captureGroupModel.Group).To(Equal(core.StringPtr("testString")))
				Expect(captureGroupModel.Location).To(Equal([]int64{int64(38)}))

				// Construct an instance of the RuntimeEntityInterpretation model
				runtimeEntityInterpretationModel := new(assistantv2.RuntimeEntityInterpretation)
				Expect(runtimeEntityInterpretationModel).ToNot(BeNil())
				runtimeEntityInterpretationModel.CalendarType = core.StringPtr("testString")
				runtimeEntityInterpretationModel.DatetimeLink = core.StringPtr("testString")
				runtimeEntityInterpretationModel.Festival = core.StringPtr("testString")
				runtimeEntityInterpretationModel.Granularity = core.StringPtr("day")
				runtimeEntityInterpretationModel.RangeLink = core.StringPtr("testString")
				runtimeEntityInterpretationModel.RangeModifier = core.StringPtr("testString")
				runtimeEntityInterpretationModel.RelativeDay = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeMonth = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeWeek = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeWeekend = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeYear = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificDay = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificDayOfWeek = core.StringPtr("testString")
				runtimeEntityInterpretationModel.SpecificMonth = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificQuarter = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificYear = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.NumericValue = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.Subtype = core.StringPtr("testString")
				runtimeEntityInterpretationModel.PartOfDay = core.StringPtr("testString")
				runtimeEntityInterpretationModel.RelativeHour = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeMinute = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeSecond = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificHour = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificMinute = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificSecond = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.Timezone = core.StringPtr("testString")
				Expect(runtimeEntityInterpretationModel.CalendarType).To(Equal(core.StringPtr("testString")))
				Expect(runtimeEntityInterpretationModel.DatetimeLink).To(Equal(core.StringPtr("testString")))
				Expect(runtimeEntityInterpretationModel.Festival).To(Equal(core.StringPtr("testString")))
				Expect(runtimeEntityInterpretationModel.Granularity).To(Equal(core.StringPtr("day")))
				Expect(runtimeEntityInterpretationModel.RangeLink).To(Equal(core.StringPtr("testString")))
				Expect(runtimeEntityInterpretationModel.RangeModifier).To(Equal(core.StringPtr("testString")))
				Expect(runtimeEntityInterpretationModel.RelativeDay).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(runtimeEntityInterpretationModel.RelativeMonth).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(runtimeEntityInterpretationModel.RelativeWeek).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(runtimeEntityInterpretationModel.RelativeWeekend).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(runtimeEntityInterpretationModel.RelativeYear).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(runtimeEntityInterpretationModel.SpecificDay).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(runtimeEntityInterpretationModel.SpecificDayOfWeek).To(Equal(core.StringPtr("testString")))
				Expect(runtimeEntityInterpretationModel.SpecificMonth).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(runtimeEntityInterpretationModel.SpecificQuarter).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(runtimeEntityInterpretationModel.SpecificYear).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(runtimeEntityInterpretationModel.NumericValue).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(runtimeEntityInterpretationModel.Subtype).To(Equal(core.StringPtr("testString")))
				Expect(runtimeEntityInterpretationModel.PartOfDay).To(Equal(core.StringPtr("testString")))
				Expect(runtimeEntityInterpretationModel.RelativeHour).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(runtimeEntityInterpretationModel.RelativeMinute).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(runtimeEntityInterpretationModel.RelativeSecond).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(runtimeEntityInterpretationModel.SpecificHour).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(runtimeEntityInterpretationModel.SpecificMinute).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(runtimeEntityInterpretationModel.SpecificSecond).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(runtimeEntityInterpretationModel.Timezone).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the RuntimeEntityAlternative model
				runtimeEntityAlternativeModel := new(assistantv2.RuntimeEntityAlternative)
				Expect(runtimeEntityAlternativeModel).ToNot(BeNil())
				runtimeEntityAlternativeModel.Value = core.StringPtr("testString")
				runtimeEntityAlternativeModel.Confidence = core.Float64Ptr(float64(72.5))
				Expect(runtimeEntityAlternativeModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(runtimeEntityAlternativeModel.Confidence).To(Equal(core.Float64Ptr(float64(72.5))))

				// Construct an instance of the RuntimeEntityRole model
				runtimeEntityRoleModel := new(assistantv2.RuntimeEntityRole)
				Expect(runtimeEntityRoleModel).ToNot(BeNil())
				runtimeEntityRoleModel.Type = core.StringPtr("date_from")
				Expect(runtimeEntityRoleModel.Type).To(Equal(core.StringPtr("date_from")))

				// Construct an instance of the RuntimeEntity model
				runtimeEntityModel := new(assistantv2.RuntimeEntity)
				Expect(runtimeEntityModel).ToNot(BeNil())
				runtimeEntityModel.Entity = core.StringPtr("testString")
				runtimeEntityModel.Location = []int64{int64(38)}
				runtimeEntityModel.Value = core.StringPtr("testString")
				runtimeEntityModel.Confidence = core.Float64Ptr(float64(72.5))
				runtimeEntityModel.Metadata = make(map[string]interface{})
				runtimeEntityModel.Groups = []assistantv2.CaptureGroup{*captureGroupModel}
				runtimeEntityModel.Interpretation = runtimeEntityInterpretationModel
				runtimeEntityModel.Alternatives = []assistantv2.RuntimeEntityAlternative{*runtimeEntityAlternativeModel}
				runtimeEntityModel.Role = runtimeEntityRoleModel
				Expect(runtimeEntityModel.Entity).To(Equal(core.StringPtr("testString")))
				Expect(runtimeEntityModel.Location).To(Equal([]int64{int64(38)}))
				Expect(runtimeEntityModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(runtimeEntityModel.Confidence).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(runtimeEntityModel.Metadata).To(Equal(make(map[string]interface{})))
				Expect(runtimeEntityModel.Groups).To(Equal([]assistantv2.CaptureGroup{*captureGroupModel}))
				Expect(runtimeEntityModel.Interpretation).To(Equal(runtimeEntityInterpretationModel))
				Expect(runtimeEntityModel.Alternatives).To(Equal([]assistantv2.RuntimeEntityAlternative{*runtimeEntityAlternativeModel}))
				Expect(runtimeEntityModel.Role).To(Equal(runtimeEntityRoleModel))

				// Construct an instance of the MessageInputOptionsSpelling model
				messageInputOptionsSpellingModel := new(assistantv2.MessageInputOptionsSpelling)
				Expect(messageInputOptionsSpellingModel).ToNot(BeNil())
				messageInputOptionsSpellingModel.Suggestions = core.BoolPtr(true)
				messageInputOptionsSpellingModel.AutoCorrect = core.BoolPtr(true)
				Expect(messageInputOptionsSpellingModel.Suggestions).To(Equal(core.BoolPtr(true)))
				Expect(messageInputOptionsSpellingModel.AutoCorrect).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the MessageInputOptions model
				messageInputOptionsModel := new(assistantv2.MessageInputOptions)
				Expect(messageInputOptionsModel).ToNot(BeNil())
				messageInputOptionsModel.Restart = core.BoolPtr(true)
				messageInputOptionsModel.AlternateIntents = core.BoolPtr(true)
				messageInputOptionsModel.Spelling = messageInputOptionsSpellingModel
				messageInputOptionsModel.Debug = core.BoolPtr(true)
				messageInputOptionsModel.ReturnContext = core.BoolPtr(true)
				messageInputOptionsModel.Export = core.BoolPtr(true)
				Expect(messageInputOptionsModel.Restart).To(Equal(core.BoolPtr(true)))
				Expect(messageInputOptionsModel.AlternateIntents).To(Equal(core.BoolPtr(true)))
				Expect(messageInputOptionsModel.Spelling).To(Equal(messageInputOptionsSpellingModel))
				Expect(messageInputOptionsModel.Debug).To(Equal(core.BoolPtr(true)))
				Expect(messageInputOptionsModel.ReturnContext).To(Equal(core.BoolPtr(true)))
				Expect(messageInputOptionsModel.Export).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the MessageInput model
				messageInputModel := new(assistantv2.MessageInput)
				Expect(messageInputModel).ToNot(BeNil())
				messageInputModel.MessageType = core.StringPtr("text")
				messageInputModel.Text = core.StringPtr("testString")
				messageInputModel.Intents = []assistantv2.RuntimeIntent{*runtimeIntentModel}
				messageInputModel.Entities = []assistantv2.RuntimeEntity{*runtimeEntityModel}
				messageInputModel.SuggestionID = core.StringPtr("testString")
				messageInputModel.Options = messageInputOptionsModel
				Expect(messageInputModel.MessageType).To(Equal(core.StringPtr("text")))
				Expect(messageInputModel.Text).To(Equal(core.StringPtr("testString")))
				Expect(messageInputModel.Intents).To(Equal([]assistantv2.RuntimeIntent{*runtimeIntentModel}))
				Expect(messageInputModel.Entities).To(Equal([]assistantv2.RuntimeEntity{*runtimeEntityModel}))
				Expect(messageInputModel.SuggestionID).To(Equal(core.StringPtr("testString")))
				Expect(messageInputModel.Options).To(Equal(messageInputOptionsModel))

				// Construct an instance of the MessageContextGlobalSystem model
				messageContextGlobalSystemModel := new(assistantv2.MessageContextGlobalSystem)
				Expect(messageContextGlobalSystemModel).ToNot(BeNil())
				messageContextGlobalSystemModel.Timezone = core.StringPtr("testString")
				messageContextGlobalSystemModel.UserID = core.StringPtr("testString")
				messageContextGlobalSystemModel.TurnCount = core.Int64Ptr(int64(38))
				messageContextGlobalSystemModel.Locale = core.StringPtr("en-us")
				messageContextGlobalSystemModel.ReferenceTime = core.StringPtr("testString")
				Expect(messageContextGlobalSystemModel.Timezone).To(Equal(core.StringPtr("testString")))
				Expect(messageContextGlobalSystemModel.UserID).To(Equal(core.StringPtr("testString")))
				Expect(messageContextGlobalSystemModel.TurnCount).To(Equal(core.Int64Ptr(int64(38))))
				Expect(messageContextGlobalSystemModel.Locale).To(Equal(core.StringPtr("en-us")))
				Expect(messageContextGlobalSystemModel.ReferenceTime).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the MessageContextGlobal model
				messageContextGlobalModel := new(assistantv2.MessageContextGlobal)
				Expect(messageContextGlobalModel).ToNot(BeNil())
				messageContextGlobalModel.System = messageContextGlobalSystemModel
				Expect(messageContextGlobalModel.System).To(Equal(messageContextGlobalSystemModel))

				// Construct an instance of the MessageContextSkillSystem model
				messageContextSkillSystemModel := new(assistantv2.MessageContextSkillSystem)
				Expect(messageContextSkillSystemModel).ToNot(BeNil())
				messageContextSkillSystemModel.State = core.StringPtr("testString")
				messageContextSkillSystemModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(messageContextSkillSystemModel.State).To(Equal(core.StringPtr("testString")))
				Expect(messageContextSkillSystemModel.GetProperties()).ToNot(BeEmpty())
				Expect(messageContextSkillSystemModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the MessageContextSkill model
				messageContextSkillModel := new(assistantv2.MessageContextSkill)
				Expect(messageContextSkillModel).ToNot(BeNil())
				messageContextSkillModel.UserDefined = make(map[string]interface{})
				messageContextSkillModel.System = messageContextSkillSystemModel
				Expect(messageContextSkillModel.UserDefined).To(Equal(make(map[string]interface{})))
				Expect(messageContextSkillModel.System).To(Equal(messageContextSkillSystemModel))

				// Construct an instance of the MessageContext model
				messageContextModel := new(assistantv2.MessageContext)
				Expect(messageContextModel).ToNot(BeNil())
				messageContextModel.Global = messageContextGlobalModel
				messageContextModel.Skills = make(map[string]assistantv2.MessageContextSkill)
				messageContextModel.Skills["foo"] = *messageContextSkillModel
				Expect(messageContextModel.Global).To(Equal(messageContextGlobalModel))
				Expect(messageContextModel.Skills["foo"]).To(Equal(*messageContextSkillModel))

				// Construct an instance of the MessageOptions model
				assistantID := "testString"
				sessionID := "testString"
				messageOptionsModel := assistantService.NewMessageOptions(assistantID, sessionID)
				messageOptionsModel.SetAssistantID("testString")
				messageOptionsModel.SetSessionID("testString")
				messageOptionsModel.SetInput(messageInputModel)
				messageOptionsModel.SetContext(messageContextModel)
				messageOptionsModel.SetUserID("testString")
				messageOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(messageOptionsModel).ToNot(BeNil())
				Expect(messageOptionsModel.AssistantID).To(Equal(core.StringPtr("testString")))
				Expect(messageOptionsModel.SessionID).To(Equal(core.StringPtr("testString")))
				Expect(messageOptionsModel.Input).To(Equal(messageInputModel))
				Expect(messageOptionsModel.Context).To(Equal(messageContextModel))
				Expect(messageOptionsModel.UserID).To(Equal(core.StringPtr("testString")))
				Expect(messageOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewMessageStatelessOptions successfully`, func() {
				// Construct an instance of the RuntimeIntent model
				runtimeIntentModel := new(assistantv2.RuntimeIntent)
				Expect(runtimeIntentModel).ToNot(BeNil())
				runtimeIntentModel.Intent = core.StringPtr("testString")
				runtimeIntentModel.Confidence = core.Float64Ptr(float64(72.5))
				Expect(runtimeIntentModel.Intent).To(Equal(core.StringPtr("testString")))
				Expect(runtimeIntentModel.Confidence).To(Equal(core.Float64Ptr(float64(72.5))))

				// Construct an instance of the CaptureGroup model
				captureGroupModel := new(assistantv2.CaptureGroup)
				Expect(captureGroupModel).ToNot(BeNil())
				captureGroupModel.Group = core.StringPtr("testString")
				captureGroupModel.Location = []int64{int64(38)}
				Expect(captureGroupModel.Group).To(Equal(core.StringPtr("testString")))
				Expect(captureGroupModel.Location).To(Equal([]int64{int64(38)}))

				// Construct an instance of the RuntimeEntityInterpretation model
				runtimeEntityInterpretationModel := new(assistantv2.RuntimeEntityInterpretation)
				Expect(runtimeEntityInterpretationModel).ToNot(BeNil())
				runtimeEntityInterpretationModel.CalendarType = core.StringPtr("testString")
				runtimeEntityInterpretationModel.DatetimeLink = core.StringPtr("testString")
				runtimeEntityInterpretationModel.Festival = core.StringPtr("testString")
				runtimeEntityInterpretationModel.Granularity = core.StringPtr("day")
				runtimeEntityInterpretationModel.RangeLink = core.StringPtr("testString")
				runtimeEntityInterpretationModel.RangeModifier = core.StringPtr("testString")
				runtimeEntityInterpretationModel.RelativeDay = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeMonth = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeWeek = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeWeekend = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeYear = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificDay = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificDayOfWeek = core.StringPtr("testString")
				runtimeEntityInterpretationModel.SpecificMonth = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificQuarter = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificYear = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.NumericValue = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.Subtype = core.StringPtr("testString")
				runtimeEntityInterpretationModel.PartOfDay = core.StringPtr("testString")
				runtimeEntityInterpretationModel.RelativeHour = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeMinute = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.RelativeSecond = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificHour = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificMinute = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.SpecificSecond = core.Float64Ptr(float64(72.5))
				runtimeEntityInterpretationModel.Timezone = core.StringPtr("testString")
				Expect(runtimeEntityInterpretationModel.CalendarType).To(Equal(core.StringPtr("testString")))
				Expect(runtimeEntityInterpretationModel.DatetimeLink).To(Equal(core.StringPtr("testString")))
				Expect(runtimeEntityInterpretationModel.Festival).To(Equal(core.StringPtr("testString")))
				Expect(runtimeEntityInterpretationModel.Granularity).To(Equal(core.StringPtr("day")))
				Expect(runtimeEntityInterpretationModel.RangeLink).To(Equal(core.StringPtr("testString")))
				Expect(runtimeEntityInterpretationModel.RangeModifier).To(Equal(core.StringPtr("testString")))
				Expect(runtimeEntityInterpretationModel.RelativeDay).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(runtimeEntityInterpretationModel.RelativeMonth).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(runtimeEntityInterpretationModel.RelativeWeek).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(runtimeEntityInterpretationModel.RelativeWeekend).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(runtimeEntityInterpretationModel.RelativeYear).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(runtimeEntityInterpretationModel.SpecificDay).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(runtimeEntityInterpretationModel.SpecificDayOfWeek).To(Equal(core.StringPtr("testString")))
				Expect(runtimeEntityInterpretationModel.SpecificMonth).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(runtimeEntityInterpretationModel.SpecificQuarter).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(runtimeEntityInterpretationModel.SpecificYear).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(runtimeEntityInterpretationModel.NumericValue).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(runtimeEntityInterpretationModel.Subtype).To(Equal(core.StringPtr("testString")))
				Expect(runtimeEntityInterpretationModel.PartOfDay).To(Equal(core.StringPtr("testString")))
				Expect(runtimeEntityInterpretationModel.RelativeHour).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(runtimeEntityInterpretationModel.RelativeMinute).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(runtimeEntityInterpretationModel.RelativeSecond).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(runtimeEntityInterpretationModel.SpecificHour).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(runtimeEntityInterpretationModel.SpecificMinute).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(runtimeEntityInterpretationModel.SpecificSecond).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(runtimeEntityInterpretationModel.Timezone).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the RuntimeEntityAlternative model
				runtimeEntityAlternativeModel := new(assistantv2.RuntimeEntityAlternative)
				Expect(runtimeEntityAlternativeModel).ToNot(BeNil())
				runtimeEntityAlternativeModel.Value = core.StringPtr("testString")
				runtimeEntityAlternativeModel.Confidence = core.Float64Ptr(float64(72.5))
				Expect(runtimeEntityAlternativeModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(runtimeEntityAlternativeModel.Confidence).To(Equal(core.Float64Ptr(float64(72.5))))

				// Construct an instance of the RuntimeEntityRole model
				runtimeEntityRoleModel := new(assistantv2.RuntimeEntityRole)
				Expect(runtimeEntityRoleModel).ToNot(BeNil())
				runtimeEntityRoleModel.Type = core.StringPtr("date_from")
				Expect(runtimeEntityRoleModel.Type).To(Equal(core.StringPtr("date_from")))

				// Construct an instance of the RuntimeEntity model
				runtimeEntityModel := new(assistantv2.RuntimeEntity)
				Expect(runtimeEntityModel).ToNot(BeNil())
				runtimeEntityModel.Entity = core.StringPtr("testString")
				runtimeEntityModel.Location = []int64{int64(38)}
				runtimeEntityModel.Value = core.StringPtr("testString")
				runtimeEntityModel.Confidence = core.Float64Ptr(float64(72.5))
				runtimeEntityModel.Metadata = make(map[string]interface{})
				runtimeEntityModel.Groups = []assistantv2.CaptureGroup{*captureGroupModel}
				runtimeEntityModel.Interpretation = runtimeEntityInterpretationModel
				runtimeEntityModel.Alternatives = []assistantv2.RuntimeEntityAlternative{*runtimeEntityAlternativeModel}
				runtimeEntityModel.Role = runtimeEntityRoleModel
				Expect(runtimeEntityModel.Entity).To(Equal(core.StringPtr("testString")))
				Expect(runtimeEntityModel.Location).To(Equal([]int64{int64(38)}))
				Expect(runtimeEntityModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(runtimeEntityModel.Confidence).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(runtimeEntityModel.Metadata).To(Equal(make(map[string]interface{})))
				Expect(runtimeEntityModel.Groups).To(Equal([]assistantv2.CaptureGroup{*captureGroupModel}))
				Expect(runtimeEntityModel.Interpretation).To(Equal(runtimeEntityInterpretationModel))
				Expect(runtimeEntityModel.Alternatives).To(Equal([]assistantv2.RuntimeEntityAlternative{*runtimeEntityAlternativeModel}))
				Expect(runtimeEntityModel.Role).To(Equal(runtimeEntityRoleModel))

				// Construct an instance of the MessageInputOptionsSpelling model
				messageInputOptionsSpellingModel := new(assistantv2.MessageInputOptionsSpelling)
				Expect(messageInputOptionsSpellingModel).ToNot(BeNil())
				messageInputOptionsSpellingModel.Suggestions = core.BoolPtr(true)
				messageInputOptionsSpellingModel.AutoCorrect = core.BoolPtr(true)
				Expect(messageInputOptionsSpellingModel.Suggestions).To(Equal(core.BoolPtr(true)))
				Expect(messageInputOptionsSpellingModel.AutoCorrect).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the MessageInputOptionsStateless model
				messageInputOptionsStatelessModel := new(assistantv2.MessageInputOptionsStateless)
				Expect(messageInputOptionsStatelessModel).ToNot(BeNil())
				messageInputOptionsStatelessModel.Restart = core.BoolPtr(true)
				messageInputOptionsStatelessModel.AlternateIntents = core.BoolPtr(true)
				messageInputOptionsStatelessModel.Spelling = messageInputOptionsSpellingModel
				messageInputOptionsStatelessModel.Debug = core.BoolPtr(true)
				Expect(messageInputOptionsStatelessModel.Restart).To(Equal(core.BoolPtr(true)))
				Expect(messageInputOptionsStatelessModel.AlternateIntents).To(Equal(core.BoolPtr(true)))
				Expect(messageInputOptionsStatelessModel.Spelling).To(Equal(messageInputOptionsSpellingModel))
				Expect(messageInputOptionsStatelessModel.Debug).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the MessageInputStateless model
				messageInputStatelessModel := new(assistantv2.MessageInputStateless)
				Expect(messageInputStatelessModel).ToNot(BeNil())
				messageInputStatelessModel.MessageType = core.StringPtr("text")
				messageInputStatelessModel.Text = core.StringPtr("testString")
				messageInputStatelessModel.Intents = []assistantv2.RuntimeIntent{*runtimeIntentModel}
				messageInputStatelessModel.Entities = []assistantv2.RuntimeEntity{*runtimeEntityModel}
				messageInputStatelessModel.SuggestionID = core.StringPtr("testString")
				messageInputStatelessModel.Options = messageInputOptionsStatelessModel
				Expect(messageInputStatelessModel.MessageType).To(Equal(core.StringPtr("text")))
				Expect(messageInputStatelessModel.Text).To(Equal(core.StringPtr("testString")))
				Expect(messageInputStatelessModel.Intents).To(Equal([]assistantv2.RuntimeIntent{*runtimeIntentModel}))
				Expect(messageInputStatelessModel.Entities).To(Equal([]assistantv2.RuntimeEntity{*runtimeEntityModel}))
				Expect(messageInputStatelessModel.SuggestionID).To(Equal(core.StringPtr("testString")))
				Expect(messageInputStatelessModel.Options).To(Equal(messageInputOptionsStatelessModel))

				// Construct an instance of the MessageContextGlobalSystem model
				messageContextGlobalSystemModel := new(assistantv2.MessageContextGlobalSystem)
				Expect(messageContextGlobalSystemModel).ToNot(BeNil())
				messageContextGlobalSystemModel.Timezone = core.StringPtr("testString")
				messageContextGlobalSystemModel.UserID = core.StringPtr("testString")
				messageContextGlobalSystemModel.TurnCount = core.Int64Ptr(int64(38))
				messageContextGlobalSystemModel.Locale = core.StringPtr("en-us")
				messageContextGlobalSystemModel.ReferenceTime = core.StringPtr("testString")
				Expect(messageContextGlobalSystemModel.Timezone).To(Equal(core.StringPtr("testString")))
				Expect(messageContextGlobalSystemModel.UserID).To(Equal(core.StringPtr("testString")))
				Expect(messageContextGlobalSystemModel.TurnCount).To(Equal(core.Int64Ptr(int64(38))))
				Expect(messageContextGlobalSystemModel.Locale).To(Equal(core.StringPtr("en-us")))
				Expect(messageContextGlobalSystemModel.ReferenceTime).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the MessageContextGlobalStateless model
				messageContextGlobalStatelessModel := new(assistantv2.MessageContextGlobalStateless)
				Expect(messageContextGlobalStatelessModel).ToNot(BeNil())
				messageContextGlobalStatelessModel.System = messageContextGlobalSystemModel
				messageContextGlobalStatelessModel.SessionID = core.StringPtr("testString")
				Expect(messageContextGlobalStatelessModel.System).To(Equal(messageContextGlobalSystemModel))
				Expect(messageContextGlobalStatelessModel.SessionID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the MessageContextSkillSystem model
				messageContextSkillSystemModel := new(assistantv2.MessageContextSkillSystem)
				Expect(messageContextSkillSystemModel).ToNot(BeNil())
				messageContextSkillSystemModel.State = core.StringPtr("testString")
				messageContextSkillSystemModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(messageContextSkillSystemModel.State).To(Equal(core.StringPtr("testString")))
				Expect(messageContextSkillSystemModel.GetProperties()).ToNot(BeEmpty())
				Expect(messageContextSkillSystemModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the MessageContextSkill model
				messageContextSkillModel := new(assistantv2.MessageContextSkill)
				Expect(messageContextSkillModel).ToNot(BeNil())
				messageContextSkillModel.UserDefined = make(map[string]interface{})
				messageContextSkillModel.System = messageContextSkillSystemModel
				Expect(messageContextSkillModel.UserDefined).To(Equal(make(map[string]interface{})))
				Expect(messageContextSkillModel.System).To(Equal(messageContextSkillSystemModel))

				// Construct an instance of the MessageContextStateless model
				messageContextStatelessModel := new(assistantv2.MessageContextStateless)
				Expect(messageContextStatelessModel).ToNot(BeNil())
				messageContextStatelessModel.Global = messageContextGlobalStatelessModel
				messageContextStatelessModel.Skills = make(map[string]assistantv2.MessageContextSkill)
				messageContextStatelessModel.Skills["foo"] = *messageContextSkillModel
				Expect(messageContextStatelessModel.Global).To(Equal(messageContextGlobalStatelessModel))
				Expect(messageContextStatelessModel.Skills["foo"]).To(Equal(*messageContextSkillModel))

				// Construct an instance of the MessageStatelessOptions model
				assistantID := "testString"
				messageStatelessOptionsModel := assistantService.NewMessageStatelessOptions(assistantID)
				messageStatelessOptionsModel.SetAssistantID("testString")
				messageStatelessOptionsModel.SetInput(messageInputStatelessModel)
				messageStatelessOptionsModel.SetContext(messageContextStatelessModel)
				messageStatelessOptionsModel.SetUserID("testString")
				messageStatelessOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(messageStatelessOptionsModel).ToNot(BeNil())
				Expect(messageStatelessOptionsModel.AssistantID).To(Equal(core.StringPtr("testString")))
				Expect(messageStatelessOptionsModel.Input).To(Equal(messageInputStatelessModel))
				Expect(messageStatelessOptionsModel.Context).To(Equal(messageContextStatelessModel))
				Expect(messageStatelessOptionsModel.UserID).To(Equal(core.StringPtr("testString")))
				Expect(messageStatelessOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRuntimeEntity successfully`, func() {
				entity := "testString"
				location := []int64{int64(38)}
				value := "testString"
				model, err := assistantService.NewRuntimeEntity(entity, location, value)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRuntimeIntent successfully`, func() {
				intent := "testString"
				confidence := float64(72.5)
				model, err := assistantService.NewRuntimeIntent(intent, confidence)
				Expect(model).ToNot(BeNil())
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
