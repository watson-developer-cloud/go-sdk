/**
 * (C) Copyright IBM Corp. 2022.
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

package assistantv1_test

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
	"github.com/watson-developer-cloud/go-sdk/v3/assistantv1"
)

var _ = Describe(`AssistantV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		version := "testString"
		It(`Instantiate service client`, func() {
			assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
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
			assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(assistantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
				URL:     "https://assistantv1/api",
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
			assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{})
			Expect(assistantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CONVERSATION_URL":       "https://assistantv1/api",
				"CONVERSATION_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
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
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
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
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
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
				"CONVERSATION_URL":       "https://assistantv1/api",
				"CONVERSATION_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
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
			assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
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
			url, err = assistantv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`Message(messageOptions *MessageOptions) - Operation response error`, func() {
		version := "testString"
		messagePath := "/v1/workspaces/testString/message"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(messagePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for nodes_visited_details query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke Message with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the MessageInput model
				messageInputModel := new(assistantv1.MessageInput)
				messageInputModel.Text = core.StringPtr("testString")
				messageInputModel.SpellingSuggestions = core.BoolPtr(false)
				messageInputModel.SpellingAutoCorrect = core.BoolPtr(false)
				messageInputModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the RuntimeIntent model
				runtimeIntentModel := new(assistantv1.RuntimeIntent)
				runtimeIntentModel.Intent = core.StringPtr("testString")
				runtimeIntentModel.Confidence = core.Float64Ptr(float64(72.5))

				// Construct an instance of the CaptureGroup model
				captureGroupModel := new(assistantv1.CaptureGroup)
				captureGroupModel.Group = core.StringPtr("testString")
				captureGroupModel.Location = []int64{int64(38)}

				// Construct an instance of the RuntimeEntityInterpretation model
				runtimeEntityInterpretationModel := new(assistantv1.RuntimeEntityInterpretation)
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
				runtimeEntityAlternativeModel := new(assistantv1.RuntimeEntityAlternative)
				runtimeEntityAlternativeModel.Value = core.StringPtr("testString")
				runtimeEntityAlternativeModel.Confidence = core.Float64Ptr(float64(72.5))

				// Construct an instance of the RuntimeEntityRole model
				runtimeEntityRoleModel := new(assistantv1.RuntimeEntityRole)
				runtimeEntityRoleModel.Type = core.StringPtr("date_from")

				// Construct an instance of the RuntimeEntity model
				runtimeEntityModel := new(assistantv1.RuntimeEntity)
				runtimeEntityModel.Entity = core.StringPtr("testString")
				runtimeEntityModel.Location = []int64{int64(38)}
				runtimeEntityModel.Value = core.StringPtr("testString")
				runtimeEntityModel.Confidence = core.Float64Ptr(float64(72.5))
				runtimeEntityModel.Groups = []assistantv1.CaptureGroup{*captureGroupModel}
				runtimeEntityModel.Interpretation = runtimeEntityInterpretationModel
				runtimeEntityModel.Alternatives = []assistantv1.RuntimeEntityAlternative{*runtimeEntityAlternativeModel}
				runtimeEntityModel.Role = runtimeEntityRoleModel

				// Construct an instance of the MessageContextMetadata model
				messageContextMetadataModel := new(assistantv1.MessageContextMetadata)
				messageContextMetadataModel.Deployment = core.StringPtr("testString")
				messageContextMetadataModel.UserID = core.StringPtr("testString")

				// Construct an instance of the Context model
				contextModel := new(assistantv1.Context)
				contextModel.ConversationID = core.StringPtr("testString")
				contextModel.System = make(map[string]interface{})
				contextModel.Metadata = messageContextMetadataModel
				contextModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeVisitedDetails model
				dialogNodeVisitedDetailsModel := new(assistantv1.DialogNodeVisitedDetails)
				dialogNodeVisitedDetailsModel.DialogNode = core.StringPtr("testString")
				dialogNodeVisitedDetailsModel.Title = core.StringPtr("testString")
				dialogNodeVisitedDetailsModel.Conditions = core.StringPtr("testString")

				// Construct an instance of the LogMessageSource model
				logMessageSourceModel := new(assistantv1.LogMessageSource)
				logMessageSourceModel.Type = core.StringPtr("dialog_node")
				logMessageSourceModel.DialogNode = core.StringPtr("testString")

				// Construct an instance of the LogMessage model
				logMessageModel := new(assistantv1.LogMessage)
				logMessageModel.Level = core.StringPtr("info")
				logMessageModel.Msg = core.StringPtr("testString")
				logMessageModel.Code = core.StringPtr("testString")
				logMessageModel.Source = logMessageSourceModel

				// Construct an instance of the DialogNodeOutputOptionsElementValue model
				dialogNodeOutputOptionsElementValueModel := new(assistantv1.DialogNodeOutputOptionsElementValue)
				dialogNodeOutputOptionsElementValueModel.Input = messageInputModel
				dialogNodeOutputOptionsElementValueModel.Intents = []assistantv1.RuntimeIntent{*runtimeIntentModel}
				dialogNodeOutputOptionsElementValueModel.Entities = []assistantv1.RuntimeEntity{*runtimeEntityModel}

				// Construct an instance of the DialogNodeOutputOptionsElement model
				dialogNodeOutputOptionsElementModel := new(assistantv1.DialogNodeOutputOptionsElement)
				dialogNodeOutputOptionsElementModel.Label = core.StringPtr("testString")
				dialogNodeOutputOptionsElementModel.Value = dialogNodeOutputOptionsElementValueModel

				// Construct an instance of the ResponseGenericChannel model
				responseGenericChannelModel := new(assistantv1.ResponseGenericChannel)
				responseGenericChannelModel.Channel = core.StringPtr("chat")

				// Construct an instance of the RuntimeResponseGenericRuntimeResponseTypeOption model
				runtimeResponseGenericModel := new(assistantv1.RuntimeResponseGenericRuntimeResponseTypeOption)
				runtimeResponseGenericModel.ResponseType = core.StringPtr("option")
				runtimeResponseGenericModel.Title = core.StringPtr("testString")
				runtimeResponseGenericModel.Description = core.StringPtr("testString")
				runtimeResponseGenericModel.Preference = core.StringPtr("dropdown")
				runtimeResponseGenericModel.Options = []assistantv1.DialogNodeOutputOptionsElement{*dialogNodeOutputOptionsElementModel}
				runtimeResponseGenericModel.Channels = []assistantv1.ResponseGenericChannel{*responseGenericChannelModel}

				// Construct an instance of the OutputData model
				outputDataModel := new(assistantv1.OutputData)
				outputDataModel.NodesVisited = []string{"testString"}
				outputDataModel.NodesVisitedDetails = []assistantv1.DialogNodeVisitedDetails{*dialogNodeVisitedDetailsModel}
				outputDataModel.LogMessages = []assistantv1.LogMessage{*logMessageModel}
				outputDataModel.Generic = []assistantv1.RuntimeResponseGenericIntf{runtimeResponseGenericModel}
				outputDataModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the MessageOptions model
				messageOptionsModel := new(assistantv1.MessageOptions)
				messageOptionsModel.WorkspaceID = core.StringPtr("testString")
				messageOptionsModel.Input = messageInputModel
				messageOptionsModel.Intents = []assistantv1.RuntimeIntent{*runtimeIntentModel}
				messageOptionsModel.Entities = []assistantv1.RuntimeEntity{*runtimeEntityModel}
				messageOptionsModel.AlternateIntents = core.BoolPtr(false)
				messageOptionsModel.Context = contextModel
				messageOptionsModel.Output = outputDataModel
				messageOptionsModel.UserID = core.StringPtr("testString")
				messageOptionsModel.NodesVisitedDetails = core.BoolPtr(false)
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
		messagePath := "/v1/workspaces/testString/message"
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
					// TODO: Add check for nodes_visited_details query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"input": {"text": "Text", "spelling_suggestions": false, "spelling_auto_correct": false, "suggested_text": "SuggestedText", "original_text": "OriginalText"}, "intents": [{"intent": "Intent", "confidence": 10}], "entities": [{"entity": "Entity", "location": [8], "value": "Value", "confidence": 10, "groups": [{"group": "Group", "location": [8]}], "interpretation": {"calendar_type": "CalendarType", "datetime_link": "DatetimeLink", "festival": "Festival", "granularity": "day", "range_link": "RangeLink", "range_modifier": "RangeModifier", "relative_day": 11, "relative_month": 13, "relative_week": 12, "relative_weekend": 15, "relative_year": 12, "specific_day": 11, "specific_day_of_week": "SpecificDayOfWeek", "specific_month": 13, "specific_quarter": 15, "specific_year": 12, "numeric_value": 12, "subtype": "Subtype", "part_of_day": "PartOfDay", "relative_hour": 12, "relative_minute": 14, "relative_second": 14, "specific_hour": 12, "specific_minute": 14, "specific_second": 14, "timezone": "Timezone"}, "alternatives": [{"value": "Value", "confidence": 10}], "role": {"type": "date_from"}}], "alternate_intents": false, "context": {"conversation_id": "ConversationID", "system": {"mapKey": "anyValue"}, "metadata": {"deployment": "Deployment", "user_id": "UserID"}}, "output": {"nodes_visited": ["NodesVisited"], "nodes_visited_details": [{"dialog_node": "DialogNode", "title": "Title", "conditions": "Conditions"}], "log_messages": [{"level": "info", "msg": "Msg", "code": "Code", "source": {"type": "dialog_node", "dialog_node": "DialogNode"}}], "generic": [{"response_type": "option", "title": "Title", "description": "Description", "preference": "dropdown", "options": [{"label": "Label", "value": {"input": {"text": "Text", "spelling_suggestions": false, "spelling_auto_correct": false, "suggested_text": "SuggestedText", "original_text": "OriginalText"}, "intents": [{"intent": "Intent", "confidence": 10}], "entities": [{"entity": "Entity", "location": [8], "value": "Value", "confidence": 10, "groups": [{"group": "Group", "location": [8]}], "interpretation": {"calendar_type": "CalendarType", "datetime_link": "DatetimeLink", "festival": "Festival", "granularity": "day", "range_link": "RangeLink", "range_modifier": "RangeModifier", "relative_day": 11, "relative_month": 13, "relative_week": 12, "relative_weekend": 15, "relative_year": 12, "specific_day": 11, "specific_day_of_week": "SpecificDayOfWeek", "specific_month": 13, "specific_quarter": 15, "specific_year": 12, "numeric_value": 12, "subtype": "Subtype", "part_of_day": "PartOfDay", "relative_hour": 12, "relative_minute": 14, "relative_second": 14, "specific_hour": 12, "specific_minute": 14, "specific_second": 14, "timezone": "Timezone"}, "alternatives": [{"value": "Value", "confidence": 10}], "role": {"type": "date_from"}}]}}], "channels": [{"channel": "chat"}]}]}, "actions": [{"name": "Name", "type": "client", "parameters": {"mapKey": "anyValue"}, "result_variable": "ResultVariable", "credentials": "Credentials"}], "user_id": "UserID"}`)
				}))
			})
			It(`Invoke Message successfully with retries`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the MessageInput model
				messageInputModel := new(assistantv1.MessageInput)
				messageInputModel.Text = core.StringPtr("testString")
				messageInputModel.SpellingSuggestions = core.BoolPtr(false)
				messageInputModel.SpellingAutoCorrect = core.BoolPtr(false)
				messageInputModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the RuntimeIntent model
				runtimeIntentModel := new(assistantv1.RuntimeIntent)
				runtimeIntentModel.Intent = core.StringPtr("testString")
				runtimeIntentModel.Confidence = core.Float64Ptr(float64(72.5))

				// Construct an instance of the CaptureGroup model
				captureGroupModel := new(assistantv1.CaptureGroup)
				captureGroupModel.Group = core.StringPtr("testString")
				captureGroupModel.Location = []int64{int64(38)}

				// Construct an instance of the RuntimeEntityInterpretation model
				runtimeEntityInterpretationModel := new(assistantv1.RuntimeEntityInterpretation)
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
				runtimeEntityAlternativeModel := new(assistantv1.RuntimeEntityAlternative)
				runtimeEntityAlternativeModel.Value = core.StringPtr("testString")
				runtimeEntityAlternativeModel.Confidence = core.Float64Ptr(float64(72.5))

				// Construct an instance of the RuntimeEntityRole model
				runtimeEntityRoleModel := new(assistantv1.RuntimeEntityRole)
				runtimeEntityRoleModel.Type = core.StringPtr("date_from")

				// Construct an instance of the RuntimeEntity model
				runtimeEntityModel := new(assistantv1.RuntimeEntity)
				runtimeEntityModel.Entity = core.StringPtr("testString")
				runtimeEntityModel.Location = []int64{int64(38)}
				runtimeEntityModel.Value = core.StringPtr("testString")
				runtimeEntityModel.Confidence = core.Float64Ptr(float64(72.5))
				runtimeEntityModel.Groups = []assistantv1.CaptureGroup{*captureGroupModel}
				runtimeEntityModel.Interpretation = runtimeEntityInterpretationModel
				runtimeEntityModel.Alternatives = []assistantv1.RuntimeEntityAlternative{*runtimeEntityAlternativeModel}
				runtimeEntityModel.Role = runtimeEntityRoleModel

				// Construct an instance of the MessageContextMetadata model
				messageContextMetadataModel := new(assistantv1.MessageContextMetadata)
				messageContextMetadataModel.Deployment = core.StringPtr("testString")
				messageContextMetadataModel.UserID = core.StringPtr("testString")

				// Construct an instance of the Context model
				contextModel := new(assistantv1.Context)
				contextModel.ConversationID = core.StringPtr("testString")
				contextModel.System = make(map[string]interface{})
				contextModel.Metadata = messageContextMetadataModel
				contextModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeVisitedDetails model
				dialogNodeVisitedDetailsModel := new(assistantv1.DialogNodeVisitedDetails)
				dialogNodeVisitedDetailsModel.DialogNode = core.StringPtr("testString")
				dialogNodeVisitedDetailsModel.Title = core.StringPtr("testString")
				dialogNodeVisitedDetailsModel.Conditions = core.StringPtr("testString")

				// Construct an instance of the LogMessageSource model
				logMessageSourceModel := new(assistantv1.LogMessageSource)
				logMessageSourceModel.Type = core.StringPtr("dialog_node")
				logMessageSourceModel.DialogNode = core.StringPtr("testString")

				// Construct an instance of the LogMessage model
				logMessageModel := new(assistantv1.LogMessage)
				logMessageModel.Level = core.StringPtr("info")
				logMessageModel.Msg = core.StringPtr("testString")
				logMessageModel.Code = core.StringPtr("testString")
				logMessageModel.Source = logMessageSourceModel

				// Construct an instance of the DialogNodeOutputOptionsElementValue model
				dialogNodeOutputOptionsElementValueModel := new(assistantv1.DialogNodeOutputOptionsElementValue)
				dialogNodeOutputOptionsElementValueModel.Input = messageInputModel
				dialogNodeOutputOptionsElementValueModel.Intents = []assistantv1.RuntimeIntent{*runtimeIntentModel}
				dialogNodeOutputOptionsElementValueModel.Entities = []assistantv1.RuntimeEntity{*runtimeEntityModel}

				// Construct an instance of the DialogNodeOutputOptionsElement model
				dialogNodeOutputOptionsElementModel := new(assistantv1.DialogNodeOutputOptionsElement)
				dialogNodeOutputOptionsElementModel.Label = core.StringPtr("testString")
				dialogNodeOutputOptionsElementModel.Value = dialogNodeOutputOptionsElementValueModel

				// Construct an instance of the ResponseGenericChannel model
				responseGenericChannelModel := new(assistantv1.ResponseGenericChannel)
				responseGenericChannelModel.Channel = core.StringPtr("chat")

				// Construct an instance of the RuntimeResponseGenericRuntimeResponseTypeOption model
				runtimeResponseGenericModel := new(assistantv1.RuntimeResponseGenericRuntimeResponseTypeOption)
				runtimeResponseGenericModel.ResponseType = core.StringPtr("option")
				runtimeResponseGenericModel.Title = core.StringPtr("testString")
				runtimeResponseGenericModel.Description = core.StringPtr("testString")
				runtimeResponseGenericModel.Preference = core.StringPtr("dropdown")
				runtimeResponseGenericModel.Options = []assistantv1.DialogNodeOutputOptionsElement{*dialogNodeOutputOptionsElementModel}
				runtimeResponseGenericModel.Channels = []assistantv1.ResponseGenericChannel{*responseGenericChannelModel}

				// Construct an instance of the OutputData model
				outputDataModel := new(assistantv1.OutputData)
				outputDataModel.NodesVisited = []string{"testString"}
				outputDataModel.NodesVisitedDetails = []assistantv1.DialogNodeVisitedDetails{*dialogNodeVisitedDetailsModel}
				outputDataModel.LogMessages = []assistantv1.LogMessage{*logMessageModel}
				outputDataModel.Generic = []assistantv1.RuntimeResponseGenericIntf{runtimeResponseGenericModel}
				outputDataModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the MessageOptions model
				messageOptionsModel := new(assistantv1.MessageOptions)
				messageOptionsModel.WorkspaceID = core.StringPtr("testString")
				messageOptionsModel.Input = messageInputModel
				messageOptionsModel.Intents = []assistantv1.RuntimeIntent{*runtimeIntentModel}
				messageOptionsModel.Entities = []assistantv1.RuntimeEntity{*runtimeEntityModel}
				messageOptionsModel.AlternateIntents = core.BoolPtr(false)
				messageOptionsModel.Context = contextModel
				messageOptionsModel.Output = outputDataModel
				messageOptionsModel.UserID = core.StringPtr("testString")
				messageOptionsModel.NodesVisitedDetails = core.BoolPtr(false)
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
					// TODO: Add check for nodes_visited_details query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"input": {"text": "Text", "spelling_suggestions": false, "spelling_auto_correct": false, "suggested_text": "SuggestedText", "original_text": "OriginalText"}, "intents": [{"intent": "Intent", "confidence": 10}], "entities": [{"entity": "Entity", "location": [8], "value": "Value", "confidence": 10, "groups": [{"group": "Group", "location": [8]}], "interpretation": {"calendar_type": "CalendarType", "datetime_link": "DatetimeLink", "festival": "Festival", "granularity": "day", "range_link": "RangeLink", "range_modifier": "RangeModifier", "relative_day": 11, "relative_month": 13, "relative_week": 12, "relative_weekend": 15, "relative_year": 12, "specific_day": 11, "specific_day_of_week": "SpecificDayOfWeek", "specific_month": 13, "specific_quarter": 15, "specific_year": 12, "numeric_value": 12, "subtype": "Subtype", "part_of_day": "PartOfDay", "relative_hour": 12, "relative_minute": 14, "relative_second": 14, "specific_hour": 12, "specific_minute": 14, "specific_second": 14, "timezone": "Timezone"}, "alternatives": [{"value": "Value", "confidence": 10}], "role": {"type": "date_from"}}], "alternate_intents": false, "context": {"conversation_id": "ConversationID", "system": {"mapKey": "anyValue"}, "metadata": {"deployment": "Deployment", "user_id": "UserID"}}, "output": {"nodes_visited": ["NodesVisited"], "nodes_visited_details": [{"dialog_node": "DialogNode", "title": "Title", "conditions": "Conditions"}], "log_messages": [{"level": "info", "msg": "Msg", "code": "Code", "source": {"type": "dialog_node", "dialog_node": "DialogNode"}}], "generic": [{"response_type": "option", "title": "Title", "description": "Description", "preference": "dropdown", "options": [{"label": "Label", "value": {"input": {"text": "Text", "spelling_suggestions": false, "spelling_auto_correct": false, "suggested_text": "SuggestedText", "original_text": "OriginalText"}, "intents": [{"intent": "Intent", "confidence": 10}], "entities": [{"entity": "Entity", "location": [8], "value": "Value", "confidence": 10, "groups": [{"group": "Group", "location": [8]}], "interpretation": {"calendar_type": "CalendarType", "datetime_link": "DatetimeLink", "festival": "Festival", "granularity": "day", "range_link": "RangeLink", "range_modifier": "RangeModifier", "relative_day": 11, "relative_month": 13, "relative_week": 12, "relative_weekend": 15, "relative_year": 12, "specific_day": 11, "specific_day_of_week": "SpecificDayOfWeek", "specific_month": 13, "specific_quarter": 15, "specific_year": 12, "numeric_value": 12, "subtype": "Subtype", "part_of_day": "PartOfDay", "relative_hour": 12, "relative_minute": 14, "relative_second": 14, "specific_hour": 12, "specific_minute": 14, "specific_second": 14, "timezone": "Timezone"}, "alternatives": [{"value": "Value", "confidence": 10}], "role": {"type": "date_from"}}]}}], "channels": [{"channel": "chat"}]}]}, "actions": [{"name": "Name", "type": "client", "parameters": {"mapKey": "anyValue"}, "result_variable": "ResultVariable", "credentials": "Credentials"}], "user_id": "UserID"}`)
				}))
			})
			It(`Invoke Message successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
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

				// Construct an instance of the MessageInput model
				messageInputModel := new(assistantv1.MessageInput)
				messageInputModel.Text = core.StringPtr("testString")
				messageInputModel.SpellingSuggestions = core.BoolPtr(false)
				messageInputModel.SpellingAutoCorrect = core.BoolPtr(false)
				messageInputModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the RuntimeIntent model
				runtimeIntentModel := new(assistantv1.RuntimeIntent)
				runtimeIntentModel.Intent = core.StringPtr("testString")
				runtimeIntentModel.Confidence = core.Float64Ptr(float64(72.5))

				// Construct an instance of the CaptureGroup model
				captureGroupModel := new(assistantv1.CaptureGroup)
				captureGroupModel.Group = core.StringPtr("testString")
				captureGroupModel.Location = []int64{int64(38)}

				// Construct an instance of the RuntimeEntityInterpretation model
				runtimeEntityInterpretationModel := new(assistantv1.RuntimeEntityInterpretation)
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
				runtimeEntityAlternativeModel := new(assistantv1.RuntimeEntityAlternative)
				runtimeEntityAlternativeModel.Value = core.StringPtr("testString")
				runtimeEntityAlternativeModel.Confidence = core.Float64Ptr(float64(72.5))

				// Construct an instance of the RuntimeEntityRole model
				runtimeEntityRoleModel := new(assistantv1.RuntimeEntityRole)
				runtimeEntityRoleModel.Type = core.StringPtr("date_from")

				// Construct an instance of the RuntimeEntity model
				runtimeEntityModel := new(assistantv1.RuntimeEntity)
				runtimeEntityModel.Entity = core.StringPtr("testString")
				runtimeEntityModel.Location = []int64{int64(38)}
				runtimeEntityModel.Value = core.StringPtr("testString")
				runtimeEntityModel.Confidence = core.Float64Ptr(float64(72.5))
				runtimeEntityModel.Groups = []assistantv1.CaptureGroup{*captureGroupModel}
				runtimeEntityModel.Interpretation = runtimeEntityInterpretationModel
				runtimeEntityModel.Alternatives = []assistantv1.RuntimeEntityAlternative{*runtimeEntityAlternativeModel}
				runtimeEntityModel.Role = runtimeEntityRoleModel

				// Construct an instance of the MessageContextMetadata model
				messageContextMetadataModel := new(assistantv1.MessageContextMetadata)
				messageContextMetadataModel.Deployment = core.StringPtr("testString")
				messageContextMetadataModel.UserID = core.StringPtr("testString")

				// Construct an instance of the Context model
				contextModel := new(assistantv1.Context)
				contextModel.ConversationID = core.StringPtr("testString")
				contextModel.System = make(map[string]interface{})
				contextModel.Metadata = messageContextMetadataModel
				contextModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeVisitedDetails model
				dialogNodeVisitedDetailsModel := new(assistantv1.DialogNodeVisitedDetails)
				dialogNodeVisitedDetailsModel.DialogNode = core.StringPtr("testString")
				dialogNodeVisitedDetailsModel.Title = core.StringPtr("testString")
				dialogNodeVisitedDetailsModel.Conditions = core.StringPtr("testString")

				// Construct an instance of the LogMessageSource model
				logMessageSourceModel := new(assistantv1.LogMessageSource)
				logMessageSourceModel.Type = core.StringPtr("dialog_node")
				logMessageSourceModel.DialogNode = core.StringPtr("testString")

				// Construct an instance of the LogMessage model
				logMessageModel := new(assistantv1.LogMessage)
				logMessageModel.Level = core.StringPtr("info")
				logMessageModel.Msg = core.StringPtr("testString")
				logMessageModel.Code = core.StringPtr("testString")
				logMessageModel.Source = logMessageSourceModel

				// Construct an instance of the DialogNodeOutputOptionsElementValue model
				dialogNodeOutputOptionsElementValueModel := new(assistantv1.DialogNodeOutputOptionsElementValue)
				dialogNodeOutputOptionsElementValueModel.Input = messageInputModel
				dialogNodeOutputOptionsElementValueModel.Intents = []assistantv1.RuntimeIntent{*runtimeIntentModel}
				dialogNodeOutputOptionsElementValueModel.Entities = []assistantv1.RuntimeEntity{*runtimeEntityModel}

				// Construct an instance of the DialogNodeOutputOptionsElement model
				dialogNodeOutputOptionsElementModel := new(assistantv1.DialogNodeOutputOptionsElement)
				dialogNodeOutputOptionsElementModel.Label = core.StringPtr("testString")
				dialogNodeOutputOptionsElementModel.Value = dialogNodeOutputOptionsElementValueModel

				// Construct an instance of the ResponseGenericChannel model
				responseGenericChannelModel := new(assistantv1.ResponseGenericChannel)
				responseGenericChannelModel.Channel = core.StringPtr("chat")

				// Construct an instance of the RuntimeResponseGenericRuntimeResponseTypeOption model
				runtimeResponseGenericModel := new(assistantv1.RuntimeResponseGenericRuntimeResponseTypeOption)
				runtimeResponseGenericModel.ResponseType = core.StringPtr("option")
				runtimeResponseGenericModel.Title = core.StringPtr("testString")
				runtimeResponseGenericModel.Description = core.StringPtr("testString")
				runtimeResponseGenericModel.Preference = core.StringPtr("dropdown")
				runtimeResponseGenericModel.Options = []assistantv1.DialogNodeOutputOptionsElement{*dialogNodeOutputOptionsElementModel}
				runtimeResponseGenericModel.Channels = []assistantv1.ResponseGenericChannel{*responseGenericChannelModel}

				// Construct an instance of the OutputData model
				outputDataModel := new(assistantv1.OutputData)
				outputDataModel.NodesVisited = []string{"testString"}
				outputDataModel.NodesVisitedDetails = []assistantv1.DialogNodeVisitedDetails{*dialogNodeVisitedDetailsModel}
				outputDataModel.LogMessages = []assistantv1.LogMessage{*logMessageModel}
				outputDataModel.Generic = []assistantv1.RuntimeResponseGenericIntf{runtimeResponseGenericModel}
				outputDataModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the MessageOptions model
				messageOptionsModel := new(assistantv1.MessageOptions)
				messageOptionsModel.WorkspaceID = core.StringPtr("testString")
				messageOptionsModel.Input = messageInputModel
				messageOptionsModel.Intents = []assistantv1.RuntimeIntent{*runtimeIntentModel}
				messageOptionsModel.Entities = []assistantv1.RuntimeEntity{*runtimeEntityModel}
				messageOptionsModel.AlternateIntents = core.BoolPtr(false)
				messageOptionsModel.Context = contextModel
				messageOptionsModel.Output = outputDataModel
				messageOptionsModel.UserID = core.StringPtr("testString")
				messageOptionsModel.NodesVisitedDetails = core.BoolPtr(false)
				messageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.Message(messageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke Message with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the MessageInput model
				messageInputModel := new(assistantv1.MessageInput)
				messageInputModel.Text = core.StringPtr("testString")
				messageInputModel.SpellingSuggestions = core.BoolPtr(false)
				messageInputModel.SpellingAutoCorrect = core.BoolPtr(false)
				messageInputModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the RuntimeIntent model
				runtimeIntentModel := new(assistantv1.RuntimeIntent)
				runtimeIntentModel.Intent = core.StringPtr("testString")
				runtimeIntentModel.Confidence = core.Float64Ptr(float64(72.5))

				// Construct an instance of the CaptureGroup model
				captureGroupModel := new(assistantv1.CaptureGroup)
				captureGroupModel.Group = core.StringPtr("testString")
				captureGroupModel.Location = []int64{int64(38)}

				// Construct an instance of the RuntimeEntityInterpretation model
				runtimeEntityInterpretationModel := new(assistantv1.RuntimeEntityInterpretation)
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
				runtimeEntityAlternativeModel := new(assistantv1.RuntimeEntityAlternative)
				runtimeEntityAlternativeModel.Value = core.StringPtr("testString")
				runtimeEntityAlternativeModel.Confidence = core.Float64Ptr(float64(72.5))

				// Construct an instance of the RuntimeEntityRole model
				runtimeEntityRoleModel := new(assistantv1.RuntimeEntityRole)
				runtimeEntityRoleModel.Type = core.StringPtr("date_from")

				// Construct an instance of the RuntimeEntity model
				runtimeEntityModel := new(assistantv1.RuntimeEntity)
				runtimeEntityModel.Entity = core.StringPtr("testString")
				runtimeEntityModel.Location = []int64{int64(38)}
				runtimeEntityModel.Value = core.StringPtr("testString")
				runtimeEntityModel.Confidence = core.Float64Ptr(float64(72.5))
				runtimeEntityModel.Groups = []assistantv1.CaptureGroup{*captureGroupModel}
				runtimeEntityModel.Interpretation = runtimeEntityInterpretationModel
				runtimeEntityModel.Alternatives = []assistantv1.RuntimeEntityAlternative{*runtimeEntityAlternativeModel}
				runtimeEntityModel.Role = runtimeEntityRoleModel

				// Construct an instance of the MessageContextMetadata model
				messageContextMetadataModel := new(assistantv1.MessageContextMetadata)
				messageContextMetadataModel.Deployment = core.StringPtr("testString")
				messageContextMetadataModel.UserID = core.StringPtr("testString")

				// Construct an instance of the Context model
				contextModel := new(assistantv1.Context)
				contextModel.ConversationID = core.StringPtr("testString")
				contextModel.System = make(map[string]interface{})
				contextModel.Metadata = messageContextMetadataModel
				contextModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeVisitedDetails model
				dialogNodeVisitedDetailsModel := new(assistantv1.DialogNodeVisitedDetails)
				dialogNodeVisitedDetailsModel.DialogNode = core.StringPtr("testString")
				dialogNodeVisitedDetailsModel.Title = core.StringPtr("testString")
				dialogNodeVisitedDetailsModel.Conditions = core.StringPtr("testString")

				// Construct an instance of the LogMessageSource model
				logMessageSourceModel := new(assistantv1.LogMessageSource)
				logMessageSourceModel.Type = core.StringPtr("dialog_node")
				logMessageSourceModel.DialogNode = core.StringPtr("testString")

				// Construct an instance of the LogMessage model
				logMessageModel := new(assistantv1.LogMessage)
				logMessageModel.Level = core.StringPtr("info")
				logMessageModel.Msg = core.StringPtr("testString")
				logMessageModel.Code = core.StringPtr("testString")
				logMessageModel.Source = logMessageSourceModel

				// Construct an instance of the DialogNodeOutputOptionsElementValue model
				dialogNodeOutputOptionsElementValueModel := new(assistantv1.DialogNodeOutputOptionsElementValue)
				dialogNodeOutputOptionsElementValueModel.Input = messageInputModel
				dialogNodeOutputOptionsElementValueModel.Intents = []assistantv1.RuntimeIntent{*runtimeIntentModel}
				dialogNodeOutputOptionsElementValueModel.Entities = []assistantv1.RuntimeEntity{*runtimeEntityModel}

				// Construct an instance of the DialogNodeOutputOptionsElement model
				dialogNodeOutputOptionsElementModel := new(assistantv1.DialogNodeOutputOptionsElement)
				dialogNodeOutputOptionsElementModel.Label = core.StringPtr("testString")
				dialogNodeOutputOptionsElementModel.Value = dialogNodeOutputOptionsElementValueModel

				// Construct an instance of the ResponseGenericChannel model
				responseGenericChannelModel := new(assistantv1.ResponseGenericChannel)
				responseGenericChannelModel.Channel = core.StringPtr("chat")

				// Construct an instance of the RuntimeResponseGenericRuntimeResponseTypeOption model
				runtimeResponseGenericModel := new(assistantv1.RuntimeResponseGenericRuntimeResponseTypeOption)
				runtimeResponseGenericModel.ResponseType = core.StringPtr("option")
				runtimeResponseGenericModel.Title = core.StringPtr("testString")
				runtimeResponseGenericModel.Description = core.StringPtr("testString")
				runtimeResponseGenericModel.Preference = core.StringPtr("dropdown")
				runtimeResponseGenericModel.Options = []assistantv1.DialogNodeOutputOptionsElement{*dialogNodeOutputOptionsElementModel}
				runtimeResponseGenericModel.Channels = []assistantv1.ResponseGenericChannel{*responseGenericChannelModel}

				// Construct an instance of the OutputData model
				outputDataModel := new(assistantv1.OutputData)
				outputDataModel.NodesVisited = []string{"testString"}
				outputDataModel.NodesVisitedDetails = []assistantv1.DialogNodeVisitedDetails{*dialogNodeVisitedDetailsModel}
				outputDataModel.LogMessages = []assistantv1.LogMessage{*logMessageModel}
				outputDataModel.Generic = []assistantv1.RuntimeResponseGenericIntf{runtimeResponseGenericModel}
				outputDataModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the MessageOptions model
				messageOptionsModel := new(assistantv1.MessageOptions)
				messageOptionsModel.WorkspaceID = core.StringPtr("testString")
				messageOptionsModel.Input = messageInputModel
				messageOptionsModel.Intents = []assistantv1.RuntimeIntent{*runtimeIntentModel}
				messageOptionsModel.Entities = []assistantv1.RuntimeEntity{*runtimeEntityModel}
				messageOptionsModel.AlternateIntents = core.BoolPtr(false)
				messageOptionsModel.Context = contextModel
				messageOptionsModel.Output = outputDataModel
				messageOptionsModel.UserID = core.StringPtr("testString")
				messageOptionsModel.NodesVisitedDetails = core.BoolPtr(false)
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
				messageOptionsModelNew := new(assistantv1.MessageOptions)
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
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the MessageInput model
				messageInputModel := new(assistantv1.MessageInput)
				messageInputModel.Text = core.StringPtr("testString")
				messageInputModel.SpellingSuggestions = core.BoolPtr(false)
				messageInputModel.SpellingAutoCorrect = core.BoolPtr(false)
				messageInputModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the RuntimeIntent model
				runtimeIntentModel := new(assistantv1.RuntimeIntent)
				runtimeIntentModel.Intent = core.StringPtr("testString")
				runtimeIntentModel.Confidence = core.Float64Ptr(float64(72.5))

				// Construct an instance of the CaptureGroup model
				captureGroupModel := new(assistantv1.CaptureGroup)
				captureGroupModel.Group = core.StringPtr("testString")
				captureGroupModel.Location = []int64{int64(38)}

				// Construct an instance of the RuntimeEntityInterpretation model
				runtimeEntityInterpretationModel := new(assistantv1.RuntimeEntityInterpretation)
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
				runtimeEntityAlternativeModel := new(assistantv1.RuntimeEntityAlternative)
				runtimeEntityAlternativeModel.Value = core.StringPtr("testString")
				runtimeEntityAlternativeModel.Confidence = core.Float64Ptr(float64(72.5))

				// Construct an instance of the RuntimeEntityRole model
				runtimeEntityRoleModel := new(assistantv1.RuntimeEntityRole)
				runtimeEntityRoleModel.Type = core.StringPtr("date_from")

				// Construct an instance of the RuntimeEntity model
				runtimeEntityModel := new(assistantv1.RuntimeEntity)
				runtimeEntityModel.Entity = core.StringPtr("testString")
				runtimeEntityModel.Location = []int64{int64(38)}
				runtimeEntityModel.Value = core.StringPtr("testString")
				runtimeEntityModel.Confidence = core.Float64Ptr(float64(72.5))
				runtimeEntityModel.Groups = []assistantv1.CaptureGroup{*captureGroupModel}
				runtimeEntityModel.Interpretation = runtimeEntityInterpretationModel
				runtimeEntityModel.Alternatives = []assistantv1.RuntimeEntityAlternative{*runtimeEntityAlternativeModel}
				runtimeEntityModel.Role = runtimeEntityRoleModel

				// Construct an instance of the MessageContextMetadata model
				messageContextMetadataModel := new(assistantv1.MessageContextMetadata)
				messageContextMetadataModel.Deployment = core.StringPtr("testString")
				messageContextMetadataModel.UserID = core.StringPtr("testString")

				// Construct an instance of the Context model
				contextModel := new(assistantv1.Context)
				contextModel.ConversationID = core.StringPtr("testString")
				contextModel.System = make(map[string]interface{})
				contextModel.Metadata = messageContextMetadataModel
				contextModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeVisitedDetails model
				dialogNodeVisitedDetailsModel := new(assistantv1.DialogNodeVisitedDetails)
				dialogNodeVisitedDetailsModel.DialogNode = core.StringPtr("testString")
				dialogNodeVisitedDetailsModel.Title = core.StringPtr("testString")
				dialogNodeVisitedDetailsModel.Conditions = core.StringPtr("testString")

				// Construct an instance of the LogMessageSource model
				logMessageSourceModel := new(assistantv1.LogMessageSource)
				logMessageSourceModel.Type = core.StringPtr("dialog_node")
				logMessageSourceModel.DialogNode = core.StringPtr("testString")

				// Construct an instance of the LogMessage model
				logMessageModel := new(assistantv1.LogMessage)
				logMessageModel.Level = core.StringPtr("info")
				logMessageModel.Msg = core.StringPtr("testString")
				logMessageModel.Code = core.StringPtr("testString")
				logMessageModel.Source = logMessageSourceModel

				// Construct an instance of the DialogNodeOutputOptionsElementValue model
				dialogNodeOutputOptionsElementValueModel := new(assistantv1.DialogNodeOutputOptionsElementValue)
				dialogNodeOutputOptionsElementValueModel.Input = messageInputModel
				dialogNodeOutputOptionsElementValueModel.Intents = []assistantv1.RuntimeIntent{*runtimeIntentModel}
				dialogNodeOutputOptionsElementValueModel.Entities = []assistantv1.RuntimeEntity{*runtimeEntityModel}

				// Construct an instance of the DialogNodeOutputOptionsElement model
				dialogNodeOutputOptionsElementModel := new(assistantv1.DialogNodeOutputOptionsElement)
				dialogNodeOutputOptionsElementModel.Label = core.StringPtr("testString")
				dialogNodeOutputOptionsElementModel.Value = dialogNodeOutputOptionsElementValueModel

				// Construct an instance of the ResponseGenericChannel model
				responseGenericChannelModel := new(assistantv1.ResponseGenericChannel)
				responseGenericChannelModel.Channel = core.StringPtr("chat")

				// Construct an instance of the RuntimeResponseGenericRuntimeResponseTypeOption model
				runtimeResponseGenericModel := new(assistantv1.RuntimeResponseGenericRuntimeResponseTypeOption)
				runtimeResponseGenericModel.ResponseType = core.StringPtr("option")
				runtimeResponseGenericModel.Title = core.StringPtr("testString")
				runtimeResponseGenericModel.Description = core.StringPtr("testString")
				runtimeResponseGenericModel.Preference = core.StringPtr("dropdown")
				runtimeResponseGenericModel.Options = []assistantv1.DialogNodeOutputOptionsElement{*dialogNodeOutputOptionsElementModel}
				runtimeResponseGenericModel.Channels = []assistantv1.ResponseGenericChannel{*responseGenericChannelModel}

				// Construct an instance of the OutputData model
				outputDataModel := new(assistantv1.OutputData)
				outputDataModel.NodesVisited = []string{"testString"}
				outputDataModel.NodesVisitedDetails = []assistantv1.DialogNodeVisitedDetails{*dialogNodeVisitedDetailsModel}
				outputDataModel.LogMessages = []assistantv1.LogMessage{*logMessageModel}
				outputDataModel.Generic = []assistantv1.RuntimeResponseGenericIntf{runtimeResponseGenericModel}
				outputDataModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the MessageOptions model
				messageOptionsModel := new(assistantv1.MessageOptions)
				messageOptionsModel.WorkspaceID = core.StringPtr("testString")
				messageOptionsModel.Input = messageInputModel
				messageOptionsModel.Intents = []assistantv1.RuntimeIntent{*runtimeIntentModel}
				messageOptionsModel.Entities = []assistantv1.RuntimeEntity{*runtimeEntityModel}
				messageOptionsModel.AlternateIntents = core.BoolPtr(false)
				messageOptionsModel.Context = contextModel
				messageOptionsModel.Output = outputDataModel
				messageOptionsModel.UserID = core.StringPtr("testString")
				messageOptionsModel.NodesVisitedDetails = core.BoolPtr(false)
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
	Describe(`BulkClassify(bulkClassifyOptions *BulkClassifyOptions) - Operation response error`, func() {
		version := "testString"
		bulkClassifyPath := "/v1/workspaces/testString/bulk_classify"
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
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke BulkClassify with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the BulkClassifyUtterance model
				bulkClassifyUtteranceModel := new(assistantv1.BulkClassifyUtterance)
				bulkClassifyUtteranceModel.Text = core.StringPtr("testString")

				// Construct an instance of the BulkClassifyOptions model
				bulkClassifyOptionsModel := new(assistantv1.BulkClassifyOptions)
				bulkClassifyOptionsModel.WorkspaceID = core.StringPtr("testString")
				bulkClassifyOptionsModel.Input = []assistantv1.BulkClassifyUtterance{*bulkClassifyUtteranceModel}
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
		bulkClassifyPath := "/v1/workspaces/testString/bulk_classify"
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
					fmt.Fprintf(res, "%s", `{"output": [{"input": {"text": "Text"}, "entities": [{"entity": "Entity", "location": [8], "value": "Value", "confidence": 10, "groups": [{"group": "Group", "location": [8]}], "interpretation": {"calendar_type": "CalendarType", "datetime_link": "DatetimeLink", "festival": "Festival", "granularity": "day", "range_link": "RangeLink", "range_modifier": "RangeModifier", "relative_day": 11, "relative_month": 13, "relative_week": 12, "relative_weekend": 15, "relative_year": 12, "specific_day": 11, "specific_day_of_week": "SpecificDayOfWeek", "specific_month": 13, "specific_quarter": 15, "specific_year": 12, "numeric_value": 12, "subtype": "Subtype", "part_of_day": "PartOfDay", "relative_hour": 12, "relative_minute": 14, "relative_second": 14, "specific_hour": 12, "specific_minute": 14, "specific_second": 14, "timezone": "Timezone"}, "alternatives": [{"value": "Value", "confidence": 10}], "role": {"type": "date_from"}}], "intents": [{"intent": "Intent", "confidence": 10}]}]}`)
				}))
			})
			It(`Invoke BulkClassify successfully with retries`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the BulkClassifyUtterance model
				bulkClassifyUtteranceModel := new(assistantv1.BulkClassifyUtterance)
				bulkClassifyUtteranceModel.Text = core.StringPtr("testString")

				// Construct an instance of the BulkClassifyOptions model
				bulkClassifyOptionsModel := new(assistantv1.BulkClassifyOptions)
				bulkClassifyOptionsModel.WorkspaceID = core.StringPtr("testString")
				bulkClassifyOptionsModel.Input = []assistantv1.BulkClassifyUtterance{*bulkClassifyUtteranceModel}
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
					fmt.Fprintf(res, "%s", `{"output": [{"input": {"text": "Text"}, "entities": [{"entity": "Entity", "location": [8], "value": "Value", "confidence": 10, "groups": [{"group": "Group", "location": [8]}], "interpretation": {"calendar_type": "CalendarType", "datetime_link": "DatetimeLink", "festival": "Festival", "granularity": "day", "range_link": "RangeLink", "range_modifier": "RangeModifier", "relative_day": 11, "relative_month": 13, "relative_week": 12, "relative_weekend": 15, "relative_year": 12, "specific_day": 11, "specific_day_of_week": "SpecificDayOfWeek", "specific_month": 13, "specific_quarter": 15, "specific_year": 12, "numeric_value": 12, "subtype": "Subtype", "part_of_day": "PartOfDay", "relative_hour": 12, "relative_minute": 14, "relative_second": 14, "specific_hour": 12, "specific_minute": 14, "specific_second": 14, "timezone": "Timezone"}, "alternatives": [{"value": "Value", "confidence": 10}], "role": {"type": "date_from"}}], "intents": [{"intent": "Intent", "confidence": 10}]}]}`)
				}))
			})
			It(`Invoke BulkClassify successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
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
				bulkClassifyUtteranceModel := new(assistantv1.BulkClassifyUtterance)
				bulkClassifyUtteranceModel.Text = core.StringPtr("testString")

				// Construct an instance of the BulkClassifyOptions model
				bulkClassifyOptionsModel := new(assistantv1.BulkClassifyOptions)
				bulkClassifyOptionsModel.WorkspaceID = core.StringPtr("testString")
				bulkClassifyOptionsModel.Input = []assistantv1.BulkClassifyUtterance{*bulkClassifyUtteranceModel}
				bulkClassifyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.BulkClassify(bulkClassifyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke BulkClassify with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the BulkClassifyUtterance model
				bulkClassifyUtteranceModel := new(assistantv1.BulkClassifyUtterance)
				bulkClassifyUtteranceModel.Text = core.StringPtr("testString")

				// Construct an instance of the BulkClassifyOptions model
				bulkClassifyOptionsModel := new(assistantv1.BulkClassifyOptions)
				bulkClassifyOptionsModel.WorkspaceID = core.StringPtr("testString")
				bulkClassifyOptionsModel.Input = []assistantv1.BulkClassifyUtterance{*bulkClassifyUtteranceModel}
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
				bulkClassifyOptionsModelNew := new(assistantv1.BulkClassifyOptions)
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
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the BulkClassifyUtterance model
				bulkClassifyUtteranceModel := new(assistantv1.BulkClassifyUtterance)
				bulkClassifyUtteranceModel.Text = core.StringPtr("testString")

				// Construct an instance of the BulkClassifyOptions model
				bulkClassifyOptionsModel := new(assistantv1.BulkClassifyOptions)
				bulkClassifyOptionsModel.WorkspaceID = core.StringPtr("testString")
				bulkClassifyOptionsModel.Input = []assistantv1.BulkClassifyUtterance{*bulkClassifyUtteranceModel}
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
	Describe(`ListWorkspaces(listWorkspacesOptions *ListWorkspacesOptions) - Operation response error`, func() {
		version := "testString"
		listWorkspacesPath := "/v1/workspaces"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listWorkspacesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["page_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// TODO: Add check for include_count query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"name"}))
					Expect(req.URL.Query()["cursor"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListWorkspaces with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ListWorkspacesOptions model
				listWorkspacesOptionsModel := new(assistantv1.ListWorkspacesOptions)
				listWorkspacesOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listWorkspacesOptionsModel.IncludeCount = core.BoolPtr(false)
				listWorkspacesOptionsModel.Sort = core.StringPtr("name")
				listWorkspacesOptionsModel.Cursor = core.StringPtr("testString")
				listWorkspacesOptionsModel.IncludeAudit = core.BoolPtr(false)
				listWorkspacesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.ListWorkspaces(listWorkspacesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.ListWorkspaces(listWorkspacesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListWorkspaces(listWorkspacesOptions *ListWorkspacesOptions)`, func() {
		version := "testString"
		listWorkspacesPath := "/v1/workspaces"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listWorkspacesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["page_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// TODO: Add check for include_count query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"name"}))
					Expect(req.URL.Query()["cursor"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"workspaces": [{"name": "Name", "description": "Description", "language": "Language", "workspace_id": "WorkspaceID", "dialog_nodes": [{"dialog_node": "DialogNode", "description": "Description", "conditions": "Conditions", "parent": "Parent", "previous_sibling": "PreviousSibling", "output": {"generic": [{"response_type": "video", "source": "Source", "title": "Title", "description": "Description", "channels": [{"channel": "chat"}], "channel_options": {"anyKey": "anyValue"}, "alt_text": "AltText"}], "integrations": {"mapKey": {"mapKey": "anyValue"}}, "modifiers": {"overwrite": true}}, "context": {"integrations": {"mapKey": {"mapKey": "anyValue"}}}, "metadata": {"mapKey": "anyValue"}, "next_step": {"behavior": "get_user_input", "dialog_node": "DialogNode", "selector": "condition"}, "title": "Title", "type": "standard", "event_name": "focus", "variable": "Variable", "actions": [{"name": "Name", "type": "client", "parameters": {"mapKey": "anyValue"}, "result_variable": "ResultVariable", "credentials": "Credentials"}], "digress_in": "not_available", "digress_out": "allow_returning", "digress_out_slots": "not_allowed", "user_label": "UserLabel", "disambiguation_opt_out": false, "disabled": true, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}], "counterexamples": [{"text": "Text", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "metadata": {"mapKey": "anyValue"}, "learning_opt_out": false, "system_settings": {"tooling": {"store_generic_responses": false}, "disambiguation": {"prompt": "Prompt", "none_of_the_above_prompt": "NoneOfTheAbovePrompt", "enabled": false, "sensitivity": "auto", "randomize": false, "max_suggestions": 1, "suggestion_text_policy": "SuggestionTextPolicy"}, "human_agent_assist": {"mapKey": "anyValue"}, "spelling_suggestions": false, "spelling_auto_correct": false, "system_entities": {"enabled": false}, "off_topic": {"enabled": false}}, "status": "Non Existent", "webhooks": [{"url": "URL", "name": "Name", "headers": [{"name": "Name", "value": "Value"}]}], "intents": [{"intent": "Intent", "description": "Description", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "examples": [{"text": "Text", "mentions": [{"entity": "Entity", "location": [8]}], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}], "entities": [{"entity": "Entity", "description": "Description", "metadata": {"mapKey": "anyValue"}, "fuzzy_match": true, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "values": [{"value": "Value", "metadata": {"mapKey": "anyValue"}, "type": "synonyms", "synonyms": ["Synonym"], "patterns": ["Pattern"], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}]}], "pagination": {"refresh_url": "RefreshURL", "next_url": "NextURL", "total": 5, "matched": 7, "refresh_cursor": "RefreshCursor", "next_cursor": "NextCursor"}}`)
				}))
			})
			It(`Invoke ListWorkspaces successfully with retries`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the ListWorkspacesOptions model
				listWorkspacesOptionsModel := new(assistantv1.ListWorkspacesOptions)
				listWorkspacesOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listWorkspacesOptionsModel.IncludeCount = core.BoolPtr(false)
				listWorkspacesOptionsModel.Sort = core.StringPtr("name")
				listWorkspacesOptionsModel.Cursor = core.StringPtr("testString")
				listWorkspacesOptionsModel.IncludeAudit = core.BoolPtr(false)
				listWorkspacesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.ListWorkspacesWithContext(ctx, listWorkspacesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.ListWorkspaces(listWorkspacesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.ListWorkspacesWithContext(ctx, listWorkspacesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listWorkspacesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["page_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// TODO: Add check for include_count query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"name"}))
					Expect(req.URL.Query()["cursor"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"workspaces": [{"name": "Name", "description": "Description", "language": "Language", "workspace_id": "WorkspaceID", "dialog_nodes": [{"dialog_node": "DialogNode", "description": "Description", "conditions": "Conditions", "parent": "Parent", "previous_sibling": "PreviousSibling", "output": {"generic": [{"response_type": "video", "source": "Source", "title": "Title", "description": "Description", "channels": [{"channel": "chat"}], "channel_options": {"anyKey": "anyValue"}, "alt_text": "AltText"}], "integrations": {"mapKey": {"mapKey": "anyValue"}}, "modifiers": {"overwrite": true}}, "context": {"integrations": {"mapKey": {"mapKey": "anyValue"}}}, "metadata": {"mapKey": "anyValue"}, "next_step": {"behavior": "get_user_input", "dialog_node": "DialogNode", "selector": "condition"}, "title": "Title", "type": "standard", "event_name": "focus", "variable": "Variable", "actions": [{"name": "Name", "type": "client", "parameters": {"mapKey": "anyValue"}, "result_variable": "ResultVariable", "credentials": "Credentials"}], "digress_in": "not_available", "digress_out": "allow_returning", "digress_out_slots": "not_allowed", "user_label": "UserLabel", "disambiguation_opt_out": false, "disabled": true, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}], "counterexamples": [{"text": "Text", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "metadata": {"mapKey": "anyValue"}, "learning_opt_out": false, "system_settings": {"tooling": {"store_generic_responses": false}, "disambiguation": {"prompt": "Prompt", "none_of_the_above_prompt": "NoneOfTheAbovePrompt", "enabled": false, "sensitivity": "auto", "randomize": false, "max_suggestions": 1, "suggestion_text_policy": "SuggestionTextPolicy"}, "human_agent_assist": {"mapKey": "anyValue"}, "spelling_suggestions": false, "spelling_auto_correct": false, "system_entities": {"enabled": false}, "off_topic": {"enabled": false}}, "status": "Non Existent", "webhooks": [{"url": "URL", "name": "Name", "headers": [{"name": "Name", "value": "Value"}]}], "intents": [{"intent": "Intent", "description": "Description", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "examples": [{"text": "Text", "mentions": [{"entity": "Entity", "location": [8]}], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}], "entities": [{"entity": "Entity", "description": "Description", "metadata": {"mapKey": "anyValue"}, "fuzzy_match": true, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "values": [{"value": "Value", "metadata": {"mapKey": "anyValue"}, "type": "synonyms", "synonyms": ["Synonym"], "patterns": ["Pattern"], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}]}], "pagination": {"refresh_url": "RefreshURL", "next_url": "NextURL", "total": 5, "matched": 7, "refresh_cursor": "RefreshCursor", "next_cursor": "NextCursor"}}`)
				}))
			})
			It(`Invoke ListWorkspaces successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.ListWorkspaces(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListWorkspacesOptions model
				listWorkspacesOptionsModel := new(assistantv1.ListWorkspacesOptions)
				listWorkspacesOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listWorkspacesOptionsModel.IncludeCount = core.BoolPtr(false)
				listWorkspacesOptionsModel.Sort = core.StringPtr("name")
				listWorkspacesOptionsModel.Cursor = core.StringPtr("testString")
				listWorkspacesOptionsModel.IncludeAudit = core.BoolPtr(false)
				listWorkspacesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.ListWorkspaces(listWorkspacesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListWorkspaces with error: Operation request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ListWorkspacesOptions model
				listWorkspacesOptionsModel := new(assistantv1.ListWorkspacesOptions)
				listWorkspacesOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listWorkspacesOptionsModel.IncludeCount = core.BoolPtr(false)
				listWorkspacesOptionsModel.Sort = core.StringPtr("name")
				listWorkspacesOptionsModel.Cursor = core.StringPtr("testString")
				listWorkspacesOptionsModel.IncludeAudit = core.BoolPtr(false)
				listWorkspacesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.ListWorkspaces(listWorkspacesOptionsModel)
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
			It(`Invoke ListWorkspaces successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ListWorkspacesOptions model
				listWorkspacesOptionsModel := new(assistantv1.ListWorkspacesOptions)
				listWorkspacesOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listWorkspacesOptionsModel.IncludeCount = core.BoolPtr(false)
				listWorkspacesOptionsModel.Sort = core.StringPtr("name")
				listWorkspacesOptionsModel.Cursor = core.StringPtr("testString")
				listWorkspacesOptionsModel.IncludeAudit = core.BoolPtr(false)
				listWorkspacesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.ListWorkspaces(listWorkspacesOptionsModel)
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
	Describe(`CreateWorkspace(createWorkspaceOptions *CreateWorkspaceOptions) - Operation response error`, func() {
		version := "testString"
		createWorkspacePath := "/v1/workspaces"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createWorkspacePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateWorkspace with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ResponseGenericChannel model
				responseGenericChannelModel := new(assistantv1.ResponseGenericChannel)
				responseGenericChannelModel.Channel = core.StringPtr("chat")

				// Construct an instance of the DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo model
				dialogNodeOutputGenericModel := new(assistantv1.DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo)
				dialogNodeOutputGenericModel.ResponseType = core.StringPtr("video")
				dialogNodeOutputGenericModel.Source = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Title = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Description = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Channels = []assistantv1.ResponseGenericChannel{*responseGenericChannelModel}
				dialogNodeOutputGenericModel.ChannelOptions = map[string]interface{}{"anyKey": "anyValue"}
				dialogNodeOutputGenericModel.AltText = core.StringPtr("testString")

				// Construct an instance of the DialogNodeOutputModifiers model
				dialogNodeOutputModifiersModel := new(assistantv1.DialogNodeOutputModifiers)
				dialogNodeOutputModifiersModel.Overwrite = core.BoolPtr(true)

				// Construct an instance of the DialogNodeOutput model
				dialogNodeOutputModel := new(assistantv1.DialogNodeOutput)
				dialogNodeOutputModel.Generic = []assistantv1.DialogNodeOutputGenericIntf{dialogNodeOutputGenericModel}
				dialogNodeOutputModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeOutputModel.Modifiers = dialogNodeOutputModifiersModel
				dialogNodeOutputModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeContext model
				dialogNodeContextModel := new(assistantv1.DialogNodeContext)
				dialogNodeContextModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeContextModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeNextStep model
				dialogNodeNextStepModel := new(assistantv1.DialogNodeNextStep)
				dialogNodeNextStepModel.Behavior = core.StringPtr("get_user_input")
				dialogNodeNextStepModel.DialogNode = core.StringPtr("testString")
				dialogNodeNextStepModel.Selector = core.StringPtr("condition")

				// Construct an instance of the DialogNodeAction model
				dialogNodeActionModel := new(assistantv1.DialogNodeAction)
				dialogNodeActionModel.Name = core.StringPtr("testString")
				dialogNodeActionModel.Type = core.StringPtr("client")
				dialogNodeActionModel.Parameters = make(map[string]interface{})
				dialogNodeActionModel.ResultVariable = core.StringPtr("testString")
				dialogNodeActionModel.Credentials = core.StringPtr("testString")

				// Construct an instance of the DialogNode model
				dialogNodeModel := new(assistantv1.DialogNode)
				dialogNodeModel.DialogNode = core.StringPtr("testString")
				dialogNodeModel.Description = core.StringPtr("testString")
				dialogNodeModel.Conditions = core.StringPtr("testString")
				dialogNodeModel.Parent = core.StringPtr("testString")
				dialogNodeModel.PreviousSibling = core.StringPtr("testString")
				dialogNodeModel.Output = dialogNodeOutputModel
				dialogNodeModel.Context = dialogNodeContextModel
				dialogNodeModel.Metadata = make(map[string]interface{})
				dialogNodeModel.NextStep = dialogNodeNextStepModel
				dialogNodeModel.Title = core.StringPtr("testString")
				dialogNodeModel.Type = core.StringPtr("standard")
				dialogNodeModel.EventName = core.StringPtr("focus")
				dialogNodeModel.Variable = core.StringPtr("testString")
				dialogNodeModel.Actions = []assistantv1.DialogNodeAction{*dialogNodeActionModel}
				dialogNodeModel.DigressIn = core.StringPtr("not_available")
				dialogNodeModel.DigressOut = core.StringPtr("allow_returning")
				dialogNodeModel.DigressOutSlots = core.StringPtr("not_allowed")
				dialogNodeModel.UserLabel = core.StringPtr("testString")
				dialogNodeModel.DisambiguationOptOut = core.BoolPtr(false)

				// Construct an instance of the Counterexample model
				counterexampleModel := new(assistantv1.Counterexample)
				counterexampleModel.Text = core.StringPtr("testString")

				// Construct an instance of the WorkspaceSystemSettingsTooling model
				workspaceSystemSettingsToolingModel := new(assistantv1.WorkspaceSystemSettingsTooling)
				workspaceSystemSettingsToolingModel.StoreGenericResponses = core.BoolPtr(true)

				// Construct an instance of the WorkspaceSystemSettingsDisambiguation model
				workspaceSystemSettingsDisambiguationModel := new(assistantv1.WorkspaceSystemSettingsDisambiguation)
				workspaceSystemSettingsDisambiguationModel.Prompt = core.StringPtr("testString")
				workspaceSystemSettingsDisambiguationModel.NoneOfTheAbovePrompt = core.StringPtr("testString")
				workspaceSystemSettingsDisambiguationModel.Enabled = core.BoolPtr(false)
				workspaceSystemSettingsDisambiguationModel.Sensitivity = core.StringPtr("auto")
				workspaceSystemSettingsDisambiguationModel.Randomize = core.BoolPtr(true)
				workspaceSystemSettingsDisambiguationModel.MaxSuggestions = core.Int64Ptr(int64(1))
				workspaceSystemSettingsDisambiguationModel.SuggestionTextPolicy = core.StringPtr("testString")

				// Construct an instance of the WorkspaceSystemSettingsSystemEntities model
				workspaceSystemSettingsSystemEntitiesModel := new(assistantv1.WorkspaceSystemSettingsSystemEntities)
				workspaceSystemSettingsSystemEntitiesModel.Enabled = core.BoolPtr(false)

				// Construct an instance of the WorkspaceSystemSettingsOffTopic model
				workspaceSystemSettingsOffTopicModel := new(assistantv1.WorkspaceSystemSettingsOffTopic)
				workspaceSystemSettingsOffTopicModel.Enabled = core.BoolPtr(false)

				// Construct an instance of the WorkspaceSystemSettings model
				workspaceSystemSettingsModel := new(assistantv1.WorkspaceSystemSettings)
				workspaceSystemSettingsModel.Tooling = workspaceSystemSettingsToolingModel
				workspaceSystemSettingsModel.Disambiguation = workspaceSystemSettingsDisambiguationModel
				workspaceSystemSettingsModel.HumanAgentAssist = make(map[string]interface{})
				workspaceSystemSettingsModel.SpellingSuggestions = core.BoolPtr(false)
				workspaceSystemSettingsModel.SpellingAutoCorrect = core.BoolPtr(false)
				workspaceSystemSettingsModel.SystemEntities = workspaceSystemSettingsSystemEntitiesModel
				workspaceSystemSettingsModel.OffTopic = workspaceSystemSettingsOffTopicModel
				workspaceSystemSettingsModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the WebhookHeader model
				webhookHeaderModel := new(assistantv1.WebhookHeader)
				webhookHeaderModel.Name = core.StringPtr("testString")
				webhookHeaderModel.Value = core.StringPtr("testString")

				// Construct an instance of the Webhook model
				webhookModel := new(assistantv1.Webhook)
				webhookModel.URL = core.StringPtr("testString")
				webhookModel.Name = core.StringPtr("testString")
				webhookModel.HeadersVar = []assistantv1.WebhookHeader{*webhookHeaderModel}

				// Construct an instance of the Mention model
				mentionModel := new(assistantv1.Mention)
				mentionModel.Entity = core.StringPtr("testString")
				mentionModel.Location = []int64{int64(38)}

				// Construct an instance of the Example model
				exampleModel := new(assistantv1.Example)
				exampleModel.Text = core.StringPtr("testString")
				exampleModel.Mentions = []assistantv1.Mention{*mentionModel}

				// Construct an instance of the CreateIntent model
				createIntentModel := new(assistantv1.CreateIntent)
				createIntentModel.Intent = core.StringPtr("testString")
				createIntentModel.Description = core.StringPtr("testString")
				createIntentModel.Examples = []assistantv1.Example{*exampleModel}

				// Construct an instance of the CreateValue model
				createValueModel := new(assistantv1.CreateValue)
				createValueModel.Value = core.StringPtr("testString")
				createValueModel.Metadata = make(map[string]interface{})
				createValueModel.Type = core.StringPtr("synonyms")
				createValueModel.Synonyms = []string{"testString"}
				createValueModel.Patterns = []string{"testString"}

				// Construct an instance of the CreateEntity model
				createEntityModel := new(assistantv1.CreateEntity)
				createEntityModel.Entity = core.StringPtr("testString")
				createEntityModel.Description = core.StringPtr("testString")
				createEntityModel.Metadata = make(map[string]interface{})
				createEntityModel.FuzzyMatch = core.BoolPtr(true)
				createEntityModel.Values = []assistantv1.CreateValue{*createValueModel}

				// Construct an instance of the CreateWorkspaceOptions model
				createWorkspaceOptionsModel := new(assistantv1.CreateWorkspaceOptions)
				createWorkspaceOptionsModel.Name = core.StringPtr("testString")
				createWorkspaceOptionsModel.Description = core.StringPtr("testString")
				createWorkspaceOptionsModel.Language = core.StringPtr("testString")
				createWorkspaceOptionsModel.DialogNodes = []assistantv1.DialogNode{*dialogNodeModel}
				createWorkspaceOptionsModel.Counterexamples = []assistantv1.Counterexample{*counterexampleModel}
				createWorkspaceOptionsModel.Metadata = make(map[string]interface{})
				createWorkspaceOptionsModel.LearningOptOut = core.BoolPtr(false)
				createWorkspaceOptionsModel.SystemSettings = workspaceSystemSettingsModel
				createWorkspaceOptionsModel.Webhooks = []assistantv1.Webhook{*webhookModel}
				createWorkspaceOptionsModel.Intents = []assistantv1.CreateIntent{*createIntentModel}
				createWorkspaceOptionsModel.Entities = []assistantv1.CreateEntity{*createEntityModel}
				createWorkspaceOptionsModel.IncludeAudit = core.BoolPtr(false)
				createWorkspaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.CreateWorkspace(createWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.CreateWorkspace(createWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateWorkspace(createWorkspaceOptions *CreateWorkspaceOptions)`, func() {
		version := "testString"
		createWorkspacePath := "/v1/workspaces"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createWorkspacePath))
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
					// TODO: Add check for include_audit query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description", "language": "Language", "workspace_id": "WorkspaceID", "dialog_nodes": [{"dialog_node": "DialogNode", "description": "Description", "conditions": "Conditions", "parent": "Parent", "previous_sibling": "PreviousSibling", "output": {"generic": [{"response_type": "video", "source": "Source", "title": "Title", "description": "Description", "channels": [{"channel": "chat"}], "channel_options": {"anyKey": "anyValue"}, "alt_text": "AltText"}], "integrations": {"mapKey": {"mapKey": "anyValue"}}, "modifiers": {"overwrite": true}}, "context": {"integrations": {"mapKey": {"mapKey": "anyValue"}}}, "metadata": {"mapKey": "anyValue"}, "next_step": {"behavior": "get_user_input", "dialog_node": "DialogNode", "selector": "condition"}, "title": "Title", "type": "standard", "event_name": "focus", "variable": "Variable", "actions": [{"name": "Name", "type": "client", "parameters": {"mapKey": "anyValue"}, "result_variable": "ResultVariable", "credentials": "Credentials"}], "digress_in": "not_available", "digress_out": "allow_returning", "digress_out_slots": "not_allowed", "user_label": "UserLabel", "disambiguation_opt_out": false, "disabled": true, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}], "counterexamples": [{"text": "Text", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "metadata": {"mapKey": "anyValue"}, "learning_opt_out": false, "system_settings": {"tooling": {"store_generic_responses": false}, "disambiguation": {"prompt": "Prompt", "none_of_the_above_prompt": "NoneOfTheAbovePrompt", "enabled": false, "sensitivity": "auto", "randomize": false, "max_suggestions": 1, "suggestion_text_policy": "SuggestionTextPolicy"}, "human_agent_assist": {"mapKey": "anyValue"}, "spelling_suggestions": false, "spelling_auto_correct": false, "system_entities": {"enabled": false}, "off_topic": {"enabled": false}}, "status": "Non Existent", "webhooks": [{"url": "URL", "name": "Name", "headers": [{"name": "Name", "value": "Value"}]}], "intents": [{"intent": "Intent", "description": "Description", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "examples": [{"text": "Text", "mentions": [{"entity": "Entity", "location": [8]}], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}], "entities": [{"entity": "Entity", "description": "Description", "metadata": {"mapKey": "anyValue"}, "fuzzy_match": true, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "values": [{"value": "Value", "metadata": {"mapKey": "anyValue"}, "type": "synonyms", "synonyms": ["Synonym"], "patterns": ["Pattern"], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}]}`)
				}))
			})
			It(`Invoke CreateWorkspace successfully with retries`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the ResponseGenericChannel model
				responseGenericChannelModel := new(assistantv1.ResponseGenericChannel)
				responseGenericChannelModel.Channel = core.StringPtr("chat")

				// Construct an instance of the DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo model
				dialogNodeOutputGenericModel := new(assistantv1.DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo)
				dialogNodeOutputGenericModel.ResponseType = core.StringPtr("video")
				dialogNodeOutputGenericModel.Source = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Title = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Description = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Channels = []assistantv1.ResponseGenericChannel{*responseGenericChannelModel}
				dialogNodeOutputGenericModel.ChannelOptions = map[string]interface{}{"anyKey": "anyValue"}
				dialogNodeOutputGenericModel.AltText = core.StringPtr("testString")

				// Construct an instance of the DialogNodeOutputModifiers model
				dialogNodeOutputModifiersModel := new(assistantv1.DialogNodeOutputModifiers)
				dialogNodeOutputModifiersModel.Overwrite = core.BoolPtr(true)

				// Construct an instance of the DialogNodeOutput model
				dialogNodeOutputModel := new(assistantv1.DialogNodeOutput)
				dialogNodeOutputModel.Generic = []assistantv1.DialogNodeOutputGenericIntf{dialogNodeOutputGenericModel}
				dialogNodeOutputModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeOutputModel.Modifiers = dialogNodeOutputModifiersModel
				dialogNodeOutputModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeContext model
				dialogNodeContextModel := new(assistantv1.DialogNodeContext)
				dialogNodeContextModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeContextModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeNextStep model
				dialogNodeNextStepModel := new(assistantv1.DialogNodeNextStep)
				dialogNodeNextStepModel.Behavior = core.StringPtr("get_user_input")
				dialogNodeNextStepModel.DialogNode = core.StringPtr("testString")
				dialogNodeNextStepModel.Selector = core.StringPtr("condition")

				// Construct an instance of the DialogNodeAction model
				dialogNodeActionModel := new(assistantv1.DialogNodeAction)
				dialogNodeActionModel.Name = core.StringPtr("testString")
				dialogNodeActionModel.Type = core.StringPtr("client")
				dialogNodeActionModel.Parameters = make(map[string]interface{})
				dialogNodeActionModel.ResultVariable = core.StringPtr("testString")
				dialogNodeActionModel.Credentials = core.StringPtr("testString")

				// Construct an instance of the DialogNode model
				dialogNodeModel := new(assistantv1.DialogNode)
				dialogNodeModel.DialogNode = core.StringPtr("testString")
				dialogNodeModel.Description = core.StringPtr("testString")
				dialogNodeModel.Conditions = core.StringPtr("testString")
				dialogNodeModel.Parent = core.StringPtr("testString")
				dialogNodeModel.PreviousSibling = core.StringPtr("testString")
				dialogNodeModel.Output = dialogNodeOutputModel
				dialogNodeModel.Context = dialogNodeContextModel
				dialogNodeModel.Metadata = make(map[string]interface{})
				dialogNodeModel.NextStep = dialogNodeNextStepModel
				dialogNodeModel.Title = core.StringPtr("testString")
				dialogNodeModel.Type = core.StringPtr("standard")
				dialogNodeModel.EventName = core.StringPtr("focus")
				dialogNodeModel.Variable = core.StringPtr("testString")
				dialogNodeModel.Actions = []assistantv1.DialogNodeAction{*dialogNodeActionModel}
				dialogNodeModel.DigressIn = core.StringPtr("not_available")
				dialogNodeModel.DigressOut = core.StringPtr("allow_returning")
				dialogNodeModel.DigressOutSlots = core.StringPtr("not_allowed")
				dialogNodeModel.UserLabel = core.StringPtr("testString")
				dialogNodeModel.DisambiguationOptOut = core.BoolPtr(false)

				// Construct an instance of the Counterexample model
				counterexampleModel := new(assistantv1.Counterexample)
				counterexampleModel.Text = core.StringPtr("testString")

				// Construct an instance of the WorkspaceSystemSettingsTooling model
				workspaceSystemSettingsToolingModel := new(assistantv1.WorkspaceSystemSettingsTooling)
				workspaceSystemSettingsToolingModel.StoreGenericResponses = core.BoolPtr(true)

				// Construct an instance of the WorkspaceSystemSettingsDisambiguation model
				workspaceSystemSettingsDisambiguationModel := new(assistantv1.WorkspaceSystemSettingsDisambiguation)
				workspaceSystemSettingsDisambiguationModel.Prompt = core.StringPtr("testString")
				workspaceSystemSettingsDisambiguationModel.NoneOfTheAbovePrompt = core.StringPtr("testString")
				workspaceSystemSettingsDisambiguationModel.Enabled = core.BoolPtr(false)
				workspaceSystemSettingsDisambiguationModel.Sensitivity = core.StringPtr("auto")
				workspaceSystemSettingsDisambiguationModel.Randomize = core.BoolPtr(true)
				workspaceSystemSettingsDisambiguationModel.MaxSuggestions = core.Int64Ptr(int64(1))
				workspaceSystemSettingsDisambiguationModel.SuggestionTextPolicy = core.StringPtr("testString")

				// Construct an instance of the WorkspaceSystemSettingsSystemEntities model
				workspaceSystemSettingsSystemEntitiesModel := new(assistantv1.WorkspaceSystemSettingsSystemEntities)
				workspaceSystemSettingsSystemEntitiesModel.Enabled = core.BoolPtr(false)

				// Construct an instance of the WorkspaceSystemSettingsOffTopic model
				workspaceSystemSettingsOffTopicModel := new(assistantv1.WorkspaceSystemSettingsOffTopic)
				workspaceSystemSettingsOffTopicModel.Enabled = core.BoolPtr(false)

				// Construct an instance of the WorkspaceSystemSettings model
				workspaceSystemSettingsModel := new(assistantv1.WorkspaceSystemSettings)
				workspaceSystemSettingsModel.Tooling = workspaceSystemSettingsToolingModel
				workspaceSystemSettingsModel.Disambiguation = workspaceSystemSettingsDisambiguationModel
				workspaceSystemSettingsModel.HumanAgentAssist = make(map[string]interface{})
				workspaceSystemSettingsModel.SpellingSuggestions = core.BoolPtr(false)
				workspaceSystemSettingsModel.SpellingAutoCorrect = core.BoolPtr(false)
				workspaceSystemSettingsModel.SystemEntities = workspaceSystemSettingsSystemEntitiesModel
				workspaceSystemSettingsModel.OffTopic = workspaceSystemSettingsOffTopicModel
				workspaceSystemSettingsModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the WebhookHeader model
				webhookHeaderModel := new(assistantv1.WebhookHeader)
				webhookHeaderModel.Name = core.StringPtr("testString")
				webhookHeaderModel.Value = core.StringPtr("testString")

				// Construct an instance of the Webhook model
				webhookModel := new(assistantv1.Webhook)
				webhookModel.URL = core.StringPtr("testString")
				webhookModel.Name = core.StringPtr("testString")
				webhookModel.HeadersVar = []assistantv1.WebhookHeader{*webhookHeaderModel}

				// Construct an instance of the Mention model
				mentionModel := new(assistantv1.Mention)
				mentionModel.Entity = core.StringPtr("testString")
				mentionModel.Location = []int64{int64(38)}

				// Construct an instance of the Example model
				exampleModel := new(assistantv1.Example)
				exampleModel.Text = core.StringPtr("testString")
				exampleModel.Mentions = []assistantv1.Mention{*mentionModel}

				// Construct an instance of the CreateIntent model
				createIntentModel := new(assistantv1.CreateIntent)
				createIntentModel.Intent = core.StringPtr("testString")
				createIntentModel.Description = core.StringPtr("testString")
				createIntentModel.Examples = []assistantv1.Example{*exampleModel}

				// Construct an instance of the CreateValue model
				createValueModel := new(assistantv1.CreateValue)
				createValueModel.Value = core.StringPtr("testString")
				createValueModel.Metadata = make(map[string]interface{})
				createValueModel.Type = core.StringPtr("synonyms")
				createValueModel.Synonyms = []string{"testString"}
				createValueModel.Patterns = []string{"testString"}

				// Construct an instance of the CreateEntity model
				createEntityModel := new(assistantv1.CreateEntity)
				createEntityModel.Entity = core.StringPtr("testString")
				createEntityModel.Description = core.StringPtr("testString")
				createEntityModel.Metadata = make(map[string]interface{})
				createEntityModel.FuzzyMatch = core.BoolPtr(true)
				createEntityModel.Values = []assistantv1.CreateValue{*createValueModel}

				// Construct an instance of the CreateWorkspaceOptions model
				createWorkspaceOptionsModel := new(assistantv1.CreateWorkspaceOptions)
				createWorkspaceOptionsModel.Name = core.StringPtr("testString")
				createWorkspaceOptionsModel.Description = core.StringPtr("testString")
				createWorkspaceOptionsModel.Language = core.StringPtr("testString")
				createWorkspaceOptionsModel.DialogNodes = []assistantv1.DialogNode{*dialogNodeModel}
				createWorkspaceOptionsModel.Counterexamples = []assistantv1.Counterexample{*counterexampleModel}
				createWorkspaceOptionsModel.Metadata = make(map[string]interface{})
				createWorkspaceOptionsModel.LearningOptOut = core.BoolPtr(false)
				createWorkspaceOptionsModel.SystemSettings = workspaceSystemSettingsModel
				createWorkspaceOptionsModel.Webhooks = []assistantv1.Webhook{*webhookModel}
				createWorkspaceOptionsModel.Intents = []assistantv1.CreateIntent{*createIntentModel}
				createWorkspaceOptionsModel.Entities = []assistantv1.CreateEntity{*createEntityModel}
				createWorkspaceOptionsModel.IncludeAudit = core.BoolPtr(false)
				createWorkspaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.CreateWorkspaceWithContext(ctx, createWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.CreateWorkspace(createWorkspaceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.CreateWorkspaceWithContext(ctx, createWorkspaceOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createWorkspacePath))
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
					// TODO: Add check for include_audit query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description", "language": "Language", "workspace_id": "WorkspaceID", "dialog_nodes": [{"dialog_node": "DialogNode", "description": "Description", "conditions": "Conditions", "parent": "Parent", "previous_sibling": "PreviousSibling", "output": {"generic": [{"response_type": "video", "source": "Source", "title": "Title", "description": "Description", "channels": [{"channel": "chat"}], "channel_options": {"anyKey": "anyValue"}, "alt_text": "AltText"}], "integrations": {"mapKey": {"mapKey": "anyValue"}}, "modifiers": {"overwrite": true}}, "context": {"integrations": {"mapKey": {"mapKey": "anyValue"}}}, "metadata": {"mapKey": "anyValue"}, "next_step": {"behavior": "get_user_input", "dialog_node": "DialogNode", "selector": "condition"}, "title": "Title", "type": "standard", "event_name": "focus", "variable": "Variable", "actions": [{"name": "Name", "type": "client", "parameters": {"mapKey": "anyValue"}, "result_variable": "ResultVariable", "credentials": "Credentials"}], "digress_in": "not_available", "digress_out": "allow_returning", "digress_out_slots": "not_allowed", "user_label": "UserLabel", "disambiguation_opt_out": false, "disabled": true, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}], "counterexamples": [{"text": "Text", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "metadata": {"mapKey": "anyValue"}, "learning_opt_out": false, "system_settings": {"tooling": {"store_generic_responses": false}, "disambiguation": {"prompt": "Prompt", "none_of_the_above_prompt": "NoneOfTheAbovePrompt", "enabled": false, "sensitivity": "auto", "randomize": false, "max_suggestions": 1, "suggestion_text_policy": "SuggestionTextPolicy"}, "human_agent_assist": {"mapKey": "anyValue"}, "spelling_suggestions": false, "spelling_auto_correct": false, "system_entities": {"enabled": false}, "off_topic": {"enabled": false}}, "status": "Non Existent", "webhooks": [{"url": "URL", "name": "Name", "headers": [{"name": "Name", "value": "Value"}]}], "intents": [{"intent": "Intent", "description": "Description", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "examples": [{"text": "Text", "mentions": [{"entity": "Entity", "location": [8]}], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}], "entities": [{"entity": "Entity", "description": "Description", "metadata": {"mapKey": "anyValue"}, "fuzzy_match": true, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "values": [{"value": "Value", "metadata": {"mapKey": "anyValue"}, "type": "synonyms", "synonyms": ["Synonym"], "patterns": ["Pattern"], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}]}`)
				}))
			})
			It(`Invoke CreateWorkspace successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.CreateWorkspace(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ResponseGenericChannel model
				responseGenericChannelModel := new(assistantv1.ResponseGenericChannel)
				responseGenericChannelModel.Channel = core.StringPtr("chat")

				// Construct an instance of the DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo model
				dialogNodeOutputGenericModel := new(assistantv1.DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo)
				dialogNodeOutputGenericModel.ResponseType = core.StringPtr("video")
				dialogNodeOutputGenericModel.Source = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Title = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Description = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Channels = []assistantv1.ResponseGenericChannel{*responseGenericChannelModel}
				dialogNodeOutputGenericModel.ChannelOptions = map[string]interface{}{"anyKey": "anyValue"}
				dialogNodeOutputGenericModel.AltText = core.StringPtr("testString")

				// Construct an instance of the DialogNodeOutputModifiers model
				dialogNodeOutputModifiersModel := new(assistantv1.DialogNodeOutputModifiers)
				dialogNodeOutputModifiersModel.Overwrite = core.BoolPtr(true)

				// Construct an instance of the DialogNodeOutput model
				dialogNodeOutputModel := new(assistantv1.DialogNodeOutput)
				dialogNodeOutputModel.Generic = []assistantv1.DialogNodeOutputGenericIntf{dialogNodeOutputGenericModel}
				dialogNodeOutputModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeOutputModel.Modifiers = dialogNodeOutputModifiersModel
				dialogNodeOutputModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeContext model
				dialogNodeContextModel := new(assistantv1.DialogNodeContext)
				dialogNodeContextModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeContextModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeNextStep model
				dialogNodeNextStepModel := new(assistantv1.DialogNodeNextStep)
				dialogNodeNextStepModel.Behavior = core.StringPtr("get_user_input")
				dialogNodeNextStepModel.DialogNode = core.StringPtr("testString")
				dialogNodeNextStepModel.Selector = core.StringPtr("condition")

				// Construct an instance of the DialogNodeAction model
				dialogNodeActionModel := new(assistantv1.DialogNodeAction)
				dialogNodeActionModel.Name = core.StringPtr("testString")
				dialogNodeActionModel.Type = core.StringPtr("client")
				dialogNodeActionModel.Parameters = make(map[string]interface{})
				dialogNodeActionModel.ResultVariable = core.StringPtr("testString")
				dialogNodeActionModel.Credentials = core.StringPtr("testString")

				// Construct an instance of the DialogNode model
				dialogNodeModel := new(assistantv1.DialogNode)
				dialogNodeModel.DialogNode = core.StringPtr("testString")
				dialogNodeModel.Description = core.StringPtr("testString")
				dialogNodeModel.Conditions = core.StringPtr("testString")
				dialogNodeModel.Parent = core.StringPtr("testString")
				dialogNodeModel.PreviousSibling = core.StringPtr("testString")
				dialogNodeModel.Output = dialogNodeOutputModel
				dialogNodeModel.Context = dialogNodeContextModel
				dialogNodeModel.Metadata = make(map[string]interface{})
				dialogNodeModel.NextStep = dialogNodeNextStepModel
				dialogNodeModel.Title = core.StringPtr("testString")
				dialogNodeModel.Type = core.StringPtr("standard")
				dialogNodeModel.EventName = core.StringPtr("focus")
				dialogNodeModel.Variable = core.StringPtr("testString")
				dialogNodeModel.Actions = []assistantv1.DialogNodeAction{*dialogNodeActionModel}
				dialogNodeModel.DigressIn = core.StringPtr("not_available")
				dialogNodeModel.DigressOut = core.StringPtr("allow_returning")
				dialogNodeModel.DigressOutSlots = core.StringPtr("not_allowed")
				dialogNodeModel.UserLabel = core.StringPtr("testString")
				dialogNodeModel.DisambiguationOptOut = core.BoolPtr(false)

				// Construct an instance of the Counterexample model
				counterexampleModel := new(assistantv1.Counterexample)
				counterexampleModel.Text = core.StringPtr("testString")

				// Construct an instance of the WorkspaceSystemSettingsTooling model
				workspaceSystemSettingsToolingModel := new(assistantv1.WorkspaceSystemSettingsTooling)
				workspaceSystemSettingsToolingModel.StoreGenericResponses = core.BoolPtr(true)

				// Construct an instance of the WorkspaceSystemSettingsDisambiguation model
				workspaceSystemSettingsDisambiguationModel := new(assistantv1.WorkspaceSystemSettingsDisambiguation)
				workspaceSystemSettingsDisambiguationModel.Prompt = core.StringPtr("testString")
				workspaceSystemSettingsDisambiguationModel.NoneOfTheAbovePrompt = core.StringPtr("testString")
				workspaceSystemSettingsDisambiguationModel.Enabled = core.BoolPtr(false)
				workspaceSystemSettingsDisambiguationModel.Sensitivity = core.StringPtr("auto")
				workspaceSystemSettingsDisambiguationModel.Randomize = core.BoolPtr(true)
				workspaceSystemSettingsDisambiguationModel.MaxSuggestions = core.Int64Ptr(int64(1))
				workspaceSystemSettingsDisambiguationModel.SuggestionTextPolicy = core.StringPtr("testString")

				// Construct an instance of the WorkspaceSystemSettingsSystemEntities model
				workspaceSystemSettingsSystemEntitiesModel := new(assistantv1.WorkspaceSystemSettingsSystemEntities)
				workspaceSystemSettingsSystemEntitiesModel.Enabled = core.BoolPtr(false)

				// Construct an instance of the WorkspaceSystemSettingsOffTopic model
				workspaceSystemSettingsOffTopicModel := new(assistantv1.WorkspaceSystemSettingsOffTopic)
				workspaceSystemSettingsOffTopicModel.Enabled = core.BoolPtr(false)

				// Construct an instance of the WorkspaceSystemSettings model
				workspaceSystemSettingsModel := new(assistantv1.WorkspaceSystemSettings)
				workspaceSystemSettingsModel.Tooling = workspaceSystemSettingsToolingModel
				workspaceSystemSettingsModel.Disambiguation = workspaceSystemSettingsDisambiguationModel
				workspaceSystemSettingsModel.HumanAgentAssist = make(map[string]interface{})
				workspaceSystemSettingsModel.SpellingSuggestions = core.BoolPtr(false)
				workspaceSystemSettingsModel.SpellingAutoCorrect = core.BoolPtr(false)
				workspaceSystemSettingsModel.SystemEntities = workspaceSystemSettingsSystemEntitiesModel
				workspaceSystemSettingsModel.OffTopic = workspaceSystemSettingsOffTopicModel
				workspaceSystemSettingsModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the WebhookHeader model
				webhookHeaderModel := new(assistantv1.WebhookHeader)
				webhookHeaderModel.Name = core.StringPtr("testString")
				webhookHeaderModel.Value = core.StringPtr("testString")

				// Construct an instance of the Webhook model
				webhookModel := new(assistantv1.Webhook)
				webhookModel.URL = core.StringPtr("testString")
				webhookModel.Name = core.StringPtr("testString")
				webhookModel.HeadersVar = []assistantv1.WebhookHeader{*webhookHeaderModel}

				// Construct an instance of the Mention model
				mentionModel := new(assistantv1.Mention)
				mentionModel.Entity = core.StringPtr("testString")
				mentionModel.Location = []int64{int64(38)}

				// Construct an instance of the Example model
				exampleModel := new(assistantv1.Example)
				exampleModel.Text = core.StringPtr("testString")
				exampleModel.Mentions = []assistantv1.Mention{*mentionModel}

				// Construct an instance of the CreateIntent model
				createIntentModel := new(assistantv1.CreateIntent)
				createIntentModel.Intent = core.StringPtr("testString")
				createIntentModel.Description = core.StringPtr("testString")
				createIntentModel.Examples = []assistantv1.Example{*exampleModel}

				// Construct an instance of the CreateValue model
				createValueModel := new(assistantv1.CreateValue)
				createValueModel.Value = core.StringPtr("testString")
				createValueModel.Metadata = make(map[string]interface{})
				createValueModel.Type = core.StringPtr("synonyms")
				createValueModel.Synonyms = []string{"testString"}
				createValueModel.Patterns = []string{"testString"}

				// Construct an instance of the CreateEntity model
				createEntityModel := new(assistantv1.CreateEntity)
				createEntityModel.Entity = core.StringPtr("testString")
				createEntityModel.Description = core.StringPtr("testString")
				createEntityModel.Metadata = make(map[string]interface{})
				createEntityModel.FuzzyMatch = core.BoolPtr(true)
				createEntityModel.Values = []assistantv1.CreateValue{*createValueModel}

				// Construct an instance of the CreateWorkspaceOptions model
				createWorkspaceOptionsModel := new(assistantv1.CreateWorkspaceOptions)
				createWorkspaceOptionsModel.Name = core.StringPtr("testString")
				createWorkspaceOptionsModel.Description = core.StringPtr("testString")
				createWorkspaceOptionsModel.Language = core.StringPtr("testString")
				createWorkspaceOptionsModel.DialogNodes = []assistantv1.DialogNode{*dialogNodeModel}
				createWorkspaceOptionsModel.Counterexamples = []assistantv1.Counterexample{*counterexampleModel}
				createWorkspaceOptionsModel.Metadata = make(map[string]interface{})
				createWorkspaceOptionsModel.LearningOptOut = core.BoolPtr(false)
				createWorkspaceOptionsModel.SystemSettings = workspaceSystemSettingsModel
				createWorkspaceOptionsModel.Webhooks = []assistantv1.Webhook{*webhookModel}
				createWorkspaceOptionsModel.Intents = []assistantv1.CreateIntent{*createIntentModel}
				createWorkspaceOptionsModel.Entities = []assistantv1.CreateEntity{*createEntityModel}
				createWorkspaceOptionsModel.IncludeAudit = core.BoolPtr(false)
				createWorkspaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.CreateWorkspace(createWorkspaceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateWorkspace with error: Operation request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ResponseGenericChannel model
				responseGenericChannelModel := new(assistantv1.ResponseGenericChannel)
				responseGenericChannelModel.Channel = core.StringPtr("chat")

				// Construct an instance of the DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo model
				dialogNodeOutputGenericModel := new(assistantv1.DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo)
				dialogNodeOutputGenericModel.ResponseType = core.StringPtr("video")
				dialogNodeOutputGenericModel.Source = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Title = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Description = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Channels = []assistantv1.ResponseGenericChannel{*responseGenericChannelModel}
				dialogNodeOutputGenericModel.ChannelOptions = map[string]interface{}{"anyKey": "anyValue"}
				dialogNodeOutputGenericModel.AltText = core.StringPtr("testString")

				// Construct an instance of the DialogNodeOutputModifiers model
				dialogNodeOutputModifiersModel := new(assistantv1.DialogNodeOutputModifiers)
				dialogNodeOutputModifiersModel.Overwrite = core.BoolPtr(true)

				// Construct an instance of the DialogNodeOutput model
				dialogNodeOutputModel := new(assistantv1.DialogNodeOutput)
				dialogNodeOutputModel.Generic = []assistantv1.DialogNodeOutputGenericIntf{dialogNodeOutputGenericModel}
				dialogNodeOutputModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeOutputModel.Modifiers = dialogNodeOutputModifiersModel
				dialogNodeOutputModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeContext model
				dialogNodeContextModel := new(assistantv1.DialogNodeContext)
				dialogNodeContextModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeContextModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeNextStep model
				dialogNodeNextStepModel := new(assistantv1.DialogNodeNextStep)
				dialogNodeNextStepModel.Behavior = core.StringPtr("get_user_input")
				dialogNodeNextStepModel.DialogNode = core.StringPtr("testString")
				dialogNodeNextStepModel.Selector = core.StringPtr("condition")

				// Construct an instance of the DialogNodeAction model
				dialogNodeActionModel := new(assistantv1.DialogNodeAction)
				dialogNodeActionModel.Name = core.StringPtr("testString")
				dialogNodeActionModel.Type = core.StringPtr("client")
				dialogNodeActionModel.Parameters = make(map[string]interface{})
				dialogNodeActionModel.ResultVariable = core.StringPtr("testString")
				dialogNodeActionModel.Credentials = core.StringPtr("testString")

				// Construct an instance of the DialogNode model
				dialogNodeModel := new(assistantv1.DialogNode)
				dialogNodeModel.DialogNode = core.StringPtr("testString")
				dialogNodeModel.Description = core.StringPtr("testString")
				dialogNodeModel.Conditions = core.StringPtr("testString")
				dialogNodeModel.Parent = core.StringPtr("testString")
				dialogNodeModel.PreviousSibling = core.StringPtr("testString")
				dialogNodeModel.Output = dialogNodeOutputModel
				dialogNodeModel.Context = dialogNodeContextModel
				dialogNodeModel.Metadata = make(map[string]interface{})
				dialogNodeModel.NextStep = dialogNodeNextStepModel
				dialogNodeModel.Title = core.StringPtr("testString")
				dialogNodeModel.Type = core.StringPtr("standard")
				dialogNodeModel.EventName = core.StringPtr("focus")
				dialogNodeModel.Variable = core.StringPtr("testString")
				dialogNodeModel.Actions = []assistantv1.DialogNodeAction{*dialogNodeActionModel}
				dialogNodeModel.DigressIn = core.StringPtr("not_available")
				dialogNodeModel.DigressOut = core.StringPtr("allow_returning")
				dialogNodeModel.DigressOutSlots = core.StringPtr("not_allowed")
				dialogNodeModel.UserLabel = core.StringPtr("testString")
				dialogNodeModel.DisambiguationOptOut = core.BoolPtr(false)

				// Construct an instance of the Counterexample model
				counterexampleModel := new(assistantv1.Counterexample)
				counterexampleModel.Text = core.StringPtr("testString")

				// Construct an instance of the WorkspaceSystemSettingsTooling model
				workspaceSystemSettingsToolingModel := new(assistantv1.WorkspaceSystemSettingsTooling)
				workspaceSystemSettingsToolingModel.StoreGenericResponses = core.BoolPtr(true)

				// Construct an instance of the WorkspaceSystemSettingsDisambiguation model
				workspaceSystemSettingsDisambiguationModel := new(assistantv1.WorkspaceSystemSettingsDisambiguation)
				workspaceSystemSettingsDisambiguationModel.Prompt = core.StringPtr("testString")
				workspaceSystemSettingsDisambiguationModel.NoneOfTheAbovePrompt = core.StringPtr("testString")
				workspaceSystemSettingsDisambiguationModel.Enabled = core.BoolPtr(false)
				workspaceSystemSettingsDisambiguationModel.Sensitivity = core.StringPtr("auto")
				workspaceSystemSettingsDisambiguationModel.Randomize = core.BoolPtr(true)
				workspaceSystemSettingsDisambiguationModel.MaxSuggestions = core.Int64Ptr(int64(1))
				workspaceSystemSettingsDisambiguationModel.SuggestionTextPolicy = core.StringPtr("testString")

				// Construct an instance of the WorkspaceSystemSettingsSystemEntities model
				workspaceSystemSettingsSystemEntitiesModel := new(assistantv1.WorkspaceSystemSettingsSystemEntities)
				workspaceSystemSettingsSystemEntitiesModel.Enabled = core.BoolPtr(false)

				// Construct an instance of the WorkspaceSystemSettingsOffTopic model
				workspaceSystemSettingsOffTopicModel := new(assistantv1.WorkspaceSystemSettingsOffTopic)
				workspaceSystemSettingsOffTopicModel.Enabled = core.BoolPtr(false)

				// Construct an instance of the WorkspaceSystemSettings model
				workspaceSystemSettingsModel := new(assistantv1.WorkspaceSystemSettings)
				workspaceSystemSettingsModel.Tooling = workspaceSystemSettingsToolingModel
				workspaceSystemSettingsModel.Disambiguation = workspaceSystemSettingsDisambiguationModel
				workspaceSystemSettingsModel.HumanAgentAssist = make(map[string]interface{})
				workspaceSystemSettingsModel.SpellingSuggestions = core.BoolPtr(false)
				workspaceSystemSettingsModel.SpellingAutoCorrect = core.BoolPtr(false)
				workspaceSystemSettingsModel.SystemEntities = workspaceSystemSettingsSystemEntitiesModel
				workspaceSystemSettingsModel.OffTopic = workspaceSystemSettingsOffTopicModel
				workspaceSystemSettingsModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the WebhookHeader model
				webhookHeaderModel := new(assistantv1.WebhookHeader)
				webhookHeaderModel.Name = core.StringPtr("testString")
				webhookHeaderModel.Value = core.StringPtr("testString")

				// Construct an instance of the Webhook model
				webhookModel := new(assistantv1.Webhook)
				webhookModel.URL = core.StringPtr("testString")
				webhookModel.Name = core.StringPtr("testString")
				webhookModel.HeadersVar = []assistantv1.WebhookHeader{*webhookHeaderModel}

				// Construct an instance of the Mention model
				mentionModel := new(assistantv1.Mention)
				mentionModel.Entity = core.StringPtr("testString")
				mentionModel.Location = []int64{int64(38)}

				// Construct an instance of the Example model
				exampleModel := new(assistantv1.Example)
				exampleModel.Text = core.StringPtr("testString")
				exampleModel.Mentions = []assistantv1.Mention{*mentionModel}

				// Construct an instance of the CreateIntent model
				createIntentModel := new(assistantv1.CreateIntent)
				createIntentModel.Intent = core.StringPtr("testString")
				createIntentModel.Description = core.StringPtr("testString")
				createIntentModel.Examples = []assistantv1.Example{*exampleModel}

				// Construct an instance of the CreateValue model
				createValueModel := new(assistantv1.CreateValue)
				createValueModel.Value = core.StringPtr("testString")
				createValueModel.Metadata = make(map[string]interface{})
				createValueModel.Type = core.StringPtr("synonyms")
				createValueModel.Synonyms = []string{"testString"}
				createValueModel.Patterns = []string{"testString"}

				// Construct an instance of the CreateEntity model
				createEntityModel := new(assistantv1.CreateEntity)
				createEntityModel.Entity = core.StringPtr("testString")
				createEntityModel.Description = core.StringPtr("testString")
				createEntityModel.Metadata = make(map[string]interface{})
				createEntityModel.FuzzyMatch = core.BoolPtr(true)
				createEntityModel.Values = []assistantv1.CreateValue{*createValueModel}

				// Construct an instance of the CreateWorkspaceOptions model
				createWorkspaceOptionsModel := new(assistantv1.CreateWorkspaceOptions)
				createWorkspaceOptionsModel.Name = core.StringPtr("testString")
				createWorkspaceOptionsModel.Description = core.StringPtr("testString")
				createWorkspaceOptionsModel.Language = core.StringPtr("testString")
				createWorkspaceOptionsModel.DialogNodes = []assistantv1.DialogNode{*dialogNodeModel}
				createWorkspaceOptionsModel.Counterexamples = []assistantv1.Counterexample{*counterexampleModel}
				createWorkspaceOptionsModel.Metadata = make(map[string]interface{})
				createWorkspaceOptionsModel.LearningOptOut = core.BoolPtr(false)
				createWorkspaceOptionsModel.SystemSettings = workspaceSystemSettingsModel
				createWorkspaceOptionsModel.Webhooks = []assistantv1.Webhook{*webhookModel}
				createWorkspaceOptionsModel.Intents = []assistantv1.CreateIntent{*createIntentModel}
				createWorkspaceOptionsModel.Entities = []assistantv1.CreateEntity{*createEntityModel}
				createWorkspaceOptionsModel.IncludeAudit = core.BoolPtr(false)
				createWorkspaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.CreateWorkspace(createWorkspaceOptionsModel)
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
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateWorkspace successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ResponseGenericChannel model
				responseGenericChannelModel := new(assistantv1.ResponseGenericChannel)
				responseGenericChannelModel.Channel = core.StringPtr("chat")

				// Construct an instance of the DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo model
				dialogNodeOutputGenericModel := new(assistantv1.DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo)
				dialogNodeOutputGenericModel.ResponseType = core.StringPtr("video")
				dialogNodeOutputGenericModel.Source = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Title = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Description = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Channels = []assistantv1.ResponseGenericChannel{*responseGenericChannelModel}
				dialogNodeOutputGenericModel.ChannelOptions = map[string]interface{}{"anyKey": "anyValue"}
				dialogNodeOutputGenericModel.AltText = core.StringPtr("testString")

				// Construct an instance of the DialogNodeOutputModifiers model
				dialogNodeOutputModifiersModel := new(assistantv1.DialogNodeOutputModifiers)
				dialogNodeOutputModifiersModel.Overwrite = core.BoolPtr(true)

				// Construct an instance of the DialogNodeOutput model
				dialogNodeOutputModel := new(assistantv1.DialogNodeOutput)
				dialogNodeOutputModel.Generic = []assistantv1.DialogNodeOutputGenericIntf{dialogNodeOutputGenericModel}
				dialogNodeOutputModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeOutputModel.Modifiers = dialogNodeOutputModifiersModel
				dialogNodeOutputModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeContext model
				dialogNodeContextModel := new(assistantv1.DialogNodeContext)
				dialogNodeContextModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeContextModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeNextStep model
				dialogNodeNextStepModel := new(assistantv1.DialogNodeNextStep)
				dialogNodeNextStepModel.Behavior = core.StringPtr("get_user_input")
				dialogNodeNextStepModel.DialogNode = core.StringPtr("testString")
				dialogNodeNextStepModel.Selector = core.StringPtr("condition")

				// Construct an instance of the DialogNodeAction model
				dialogNodeActionModel := new(assistantv1.DialogNodeAction)
				dialogNodeActionModel.Name = core.StringPtr("testString")
				dialogNodeActionModel.Type = core.StringPtr("client")
				dialogNodeActionModel.Parameters = make(map[string]interface{})
				dialogNodeActionModel.ResultVariable = core.StringPtr("testString")
				dialogNodeActionModel.Credentials = core.StringPtr("testString")

				// Construct an instance of the DialogNode model
				dialogNodeModel := new(assistantv1.DialogNode)
				dialogNodeModel.DialogNode = core.StringPtr("testString")
				dialogNodeModel.Description = core.StringPtr("testString")
				dialogNodeModel.Conditions = core.StringPtr("testString")
				dialogNodeModel.Parent = core.StringPtr("testString")
				dialogNodeModel.PreviousSibling = core.StringPtr("testString")
				dialogNodeModel.Output = dialogNodeOutputModel
				dialogNodeModel.Context = dialogNodeContextModel
				dialogNodeModel.Metadata = make(map[string]interface{})
				dialogNodeModel.NextStep = dialogNodeNextStepModel
				dialogNodeModel.Title = core.StringPtr("testString")
				dialogNodeModel.Type = core.StringPtr("standard")
				dialogNodeModel.EventName = core.StringPtr("focus")
				dialogNodeModel.Variable = core.StringPtr("testString")
				dialogNodeModel.Actions = []assistantv1.DialogNodeAction{*dialogNodeActionModel}
				dialogNodeModel.DigressIn = core.StringPtr("not_available")
				dialogNodeModel.DigressOut = core.StringPtr("allow_returning")
				dialogNodeModel.DigressOutSlots = core.StringPtr("not_allowed")
				dialogNodeModel.UserLabel = core.StringPtr("testString")
				dialogNodeModel.DisambiguationOptOut = core.BoolPtr(false)

				// Construct an instance of the Counterexample model
				counterexampleModel := new(assistantv1.Counterexample)
				counterexampleModel.Text = core.StringPtr("testString")

				// Construct an instance of the WorkspaceSystemSettingsTooling model
				workspaceSystemSettingsToolingModel := new(assistantv1.WorkspaceSystemSettingsTooling)
				workspaceSystemSettingsToolingModel.StoreGenericResponses = core.BoolPtr(true)

				// Construct an instance of the WorkspaceSystemSettingsDisambiguation model
				workspaceSystemSettingsDisambiguationModel := new(assistantv1.WorkspaceSystemSettingsDisambiguation)
				workspaceSystemSettingsDisambiguationModel.Prompt = core.StringPtr("testString")
				workspaceSystemSettingsDisambiguationModel.NoneOfTheAbovePrompt = core.StringPtr("testString")
				workspaceSystemSettingsDisambiguationModel.Enabled = core.BoolPtr(false)
				workspaceSystemSettingsDisambiguationModel.Sensitivity = core.StringPtr("auto")
				workspaceSystemSettingsDisambiguationModel.Randomize = core.BoolPtr(true)
				workspaceSystemSettingsDisambiguationModel.MaxSuggestions = core.Int64Ptr(int64(1))
				workspaceSystemSettingsDisambiguationModel.SuggestionTextPolicy = core.StringPtr("testString")

				// Construct an instance of the WorkspaceSystemSettingsSystemEntities model
				workspaceSystemSettingsSystemEntitiesModel := new(assistantv1.WorkspaceSystemSettingsSystemEntities)
				workspaceSystemSettingsSystemEntitiesModel.Enabled = core.BoolPtr(false)

				// Construct an instance of the WorkspaceSystemSettingsOffTopic model
				workspaceSystemSettingsOffTopicModel := new(assistantv1.WorkspaceSystemSettingsOffTopic)
				workspaceSystemSettingsOffTopicModel.Enabled = core.BoolPtr(false)

				// Construct an instance of the WorkspaceSystemSettings model
				workspaceSystemSettingsModel := new(assistantv1.WorkspaceSystemSettings)
				workspaceSystemSettingsModel.Tooling = workspaceSystemSettingsToolingModel
				workspaceSystemSettingsModel.Disambiguation = workspaceSystemSettingsDisambiguationModel
				workspaceSystemSettingsModel.HumanAgentAssist = make(map[string]interface{})
				workspaceSystemSettingsModel.SpellingSuggestions = core.BoolPtr(false)
				workspaceSystemSettingsModel.SpellingAutoCorrect = core.BoolPtr(false)
				workspaceSystemSettingsModel.SystemEntities = workspaceSystemSettingsSystemEntitiesModel
				workspaceSystemSettingsModel.OffTopic = workspaceSystemSettingsOffTopicModel
				workspaceSystemSettingsModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the WebhookHeader model
				webhookHeaderModel := new(assistantv1.WebhookHeader)
				webhookHeaderModel.Name = core.StringPtr("testString")
				webhookHeaderModel.Value = core.StringPtr("testString")

				// Construct an instance of the Webhook model
				webhookModel := new(assistantv1.Webhook)
				webhookModel.URL = core.StringPtr("testString")
				webhookModel.Name = core.StringPtr("testString")
				webhookModel.HeadersVar = []assistantv1.WebhookHeader{*webhookHeaderModel}

				// Construct an instance of the Mention model
				mentionModel := new(assistantv1.Mention)
				mentionModel.Entity = core.StringPtr("testString")
				mentionModel.Location = []int64{int64(38)}

				// Construct an instance of the Example model
				exampleModel := new(assistantv1.Example)
				exampleModel.Text = core.StringPtr("testString")
				exampleModel.Mentions = []assistantv1.Mention{*mentionModel}

				// Construct an instance of the CreateIntent model
				createIntentModel := new(assistantv1.CreateIntent)
				createIntentModel.Intent = core.StringPtr("testString")
				createIntentModel.Description = core.StringPtr("testString")
				createIntentModel.Examples = []assistantv1.Example{*exampleModel}

				// Construct an instance of the CreateValue model
				createValueModel := new(assistantv1.CreateValue)
				createValueModel.Value = core.StringPtr("testString")
				createValueModel.Metadata = make(map[string]interface{})
				createValueModel.Type = core.StringPtr("synonyms")
				createValueModel.Synonyms = []string{"testString"}
				createValueModel.Patterns = []string{"testString"}

				// Construct an instance of the CreateEntity model
				createEntityModel := new(assistantv1.CreateEntity)
				createEntityModel.Entity = core.StringPtr("testString")
				createEntityModel.Description = core.StringPtr("testString")
				createEntityModel.Metadata = make(map[string]interface{})
				createEntityModel.FuzzyMatch = core.BoolPtr(true)
				createEntityModel.Values = []assistantv1.CreateValue{*createValueModel}

				// Construct an instance of the CreateWorkspaceOptions model
				createWorkspaceOptionsModel := new(assistantv1.CreateWorkspaceOptions)
				createWorkspaceOptionsModel.Name = core.StringPtr("testString")
				createWorkspaceOptionsModel.Description = core.StringPtr("testString")
				createWorkspaceOptionsModel.Language = core.StringPtr("testString")
				createWorkspaceOptionsModel.DialogNodes = []assistantv1.DialogNode{*dialogNodeModel}
				createWorkspaceOptionsModel.Counterexamples = []assistantv1.Counterexample{*counterexampleModel}
				createWorkspaceOptionsModel.Metadata = make(map[string]interface{})
				createWorkspaceOptionsModel.LearningOptOut = core.BoolPtr(false)
				createWorkspaceOptionsModel.SystemSettings = workspaceSystemSettingsModel
				createWorkspaceOptionsModel.Webhooks = []assistantv1.Webhook{*webhookModel}
				createWorkspaceOptionsModel.Intents = []assistantv1.CreateIntent{*createIntentModel}
				createWorkspaceOptionsModel.Entities = []assistantv1.CreateEntity{*createEntityModel}
				createWorkspaceOptionsModel.IncludeAudit = core.BoolPtr(false)
				createWorkspaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.CreateWorkspace(createWorkspaceOptionsModel)
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
	Describe(`GetWorkspace(getWorkspaceOptions *GetWorkspaceOptions) - Operation response error`, func() {
		version := "testString"
		getWorkspacePath := "/v1/workspaces/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWorkspacePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for export query parameter
					// TODO: Add check for include_audit query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"stable"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetWorkspace with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the GetWorkspaceOptions model
				getWorkspaceOptionsModel := new(assistantv1.GetWorkspaceOptions)
				getWorkspaceOptionsModel.WorkspaceID = core.StringPtr("testString")
				getWorkspaceOptionsModel.Export = core.BoolPtr(false)
				getWorkspaceOptionsModel.IncludeAudit = core.BoolPtr(false)
				getWorkspaceOptionsModel.Sort = core.StringPtr("stable")
				getWorkspaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.GetWorkspace(getWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.GetWorkspace(getWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetWorkspace(getWorkspaceOptions *GetWorkspaceOptions)`, func() {
		version := "testString"
		getWorkspacePath := "/v1/workspaces/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWorkspacePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for export query parameter
					// TODO: Add check for include_audit query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"stable"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description", "language": "Language", "workspace_id": "WorkspaceID", "dialog_nodes": [{"dialog_node": "DialogNode", "description": "Description", "conditions": "Conditions", "parent": "Parent", "previous_sibling": "PreviousSibling", "output": {"generic": [{"response_type": "video", "source": "Source", "title": "Title", "description": "Description", "channels": [{"channel": "chat"}], "channel_options": {"anyKey": "anyValue"}, "alt_text": "AltText"}], "integrations": {"mapKey": {"mapKey": "anyValue"}}, "modifiers": {"overwrite": true}}, "context": {"integrations": {"mapKey": {"mapKey": "anyValue"}}}, "metadata": {"mapKey": "anyValue"}, "next_step": {"behavior": "get_user_input", "dialog_node": "DialogNode", "selector": "condition"}, "title": "Title", "type": "standard", "event_name": "focus", "variable": "Variable", "actions": [{"name": "Name", "type": "client", "parameters": {"mapKey": "anyValue"}, "result_variable": "ResultVariable", "credentials": "Credentials"}], "digress_in": "not_available", "digress_out": "allow_returning", "digress_out_slots": "not_allowed", "user_label": "UserLabel", "disambiguation_opt_out": false, "disabled": true, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}], "counterexamples": [{"text": "Text", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "metadata": {"mapKey": "anyValue"}, "learning_opt_out": false, "system_settings": {"tooling": {"store_generic_responses": false}, "disambiguation": {"prompt": "Prompt", "none_of_the_above_prompt": "NoneOfTheAbovePrompt", "enabled": false, "sensitivity": "auto", "randomize": false, "max_suggestions": 1, "suggestion_text_policy": "SuggestionTextPolicy"}, "human_agent_assist": {"mapKey": "anyValue"}, "spelling_suggestions": false, "spelling_auto_correct": false, "system_entities": {"enabled": false}, "off_topic": {"enabled": false}}, "status": "Non Existent", "webhooks": [{"url": "URL", "name": "Name", "headers": [{"name": "Name", "value": "Value"}]}], "intents": [{"intent": "Intent", "description": "Description", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "examples": [{"text": "Text", "mentions": [{"entity": "Entity", "location": [8]}], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}], "entities": [{"entity": "Entity", "description": "Description", "metadata": {"mapKey": "anyValue"}, "fuzzy_match": true, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "values": [{"value": "Value", "metadata": {"mapKey": "anyValue"}, "type": "synonyms", "synonyms": ["Synonym"], "patterns": ["Pattern"], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}]}`)
				}))
			})
			It(`Invoke GetWorkspace successfully with retries`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the GetWorkspaceOptions model
				getWorkspaceOptionsModel := new(assistantv1.GetWorkspaceOptions)
				getWorkspaceOptionsModel.WorkspaceID = core.StringPtr("testString")
				getWorkspaceOptionsModel.Export = core.BoolPtr(false)
				getWorkspaceOptionsModel.IncludeAudit = core.BoolPtr(false)
				getWorkspaceOptionsModel.Sort = core.StringPtr("stable")
				getWorkspaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.GetWorkspaceWithContext(ctx, getWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.GetWorkspace(getWorkspaceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.GetWorkspaceWithContext(ctx, getWorkspaceOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getWorkspacePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for export query parameter
					// TODO: Add check for include_audit query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"stable"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description", "language": "Language", "workspace_id": "WorkspaceID", "dialog_nodes": [{"dialog_node": "DialogNode", "description": "Description", "conditions": "Conditions", "parent": "Parent", "previous_sibling": "PreviousSibling", "output": {"generic": [{"response_type": "video", "source": "Source", "title": "Title", "description": "Description", "channels": [{"channel": "chat"}], "channel_options": {"anyKey": "anyValue"}, "alt_text": "AltText"}], "integrations": {"mapKey": {"mapKey": "anyValue"}}, "modifiers": {"overwrite": true}}, "context": {"integrations": {"mapKey": {"mapKey": "anyValue"}}}, "metadata": {"mapKey": "anyValue"}, "next_step": {"behavior": "get_user_input", "dialog_node": "DialogNode", "selector": "condition"}, "title": "Title", "type": "standard", "event_name": "focus", "variable": "Variable", "actions": [{"name": "Name", "type": "client", "parameters": {"mapKey": "anyValue"}, "result_variable": "ResultVariable", "credentials": "Credentials"}], "digress_in": "not_available", "digress_out": "allow_returning", "digress_out_slots": "not_allowed", "user_label": "UserLabel", "disambiguation_opt_out": false, "disabled": true, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}], "counterexamples": [{"text": "Text", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "metadata": {"mapKey": "anyValue"}, "learning_opt_out": false, "system_settings": {"tooling": {"store_generic_responses": false}, "disambiguation": {"prompt": "Prompt", "none_of_the_above_prompt": "NoneOfTheAbovePrompt", "enabled": false, "sensitivity": "auto", "randomize": false, "max_suggestions": 1, "suggestion_text_policy": "SuggestionTextPolicy"}, "human_agent_assist": {"mapKey": "anyValue"}, "spelling_suggestions": false, "spelling_auto_correct": false, "system_entities": {"enabled": false}, "off_topic": {"enabled": false}}, "status": "Non Existent", "webhooks": [{"url": "URL", "name": "Name", "headers": [{"name": "Name", "value": "Value"}]}], "intents": [{"intent": "Intent", "description": "Description", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "examples": [{"text": "Text", "mentions": [{"entity": "Entity", "location": [8]}], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}], "entities": [{"entity": "Entity", "description": "Description", "metadata": {"mapKey": "anyValue"}, "fuzzy_match": true, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "values": [{"value": "Value", "metadata": {"mapKey": "anyValue"}, "type": "synonyms", "synonyms": ["Synonym"], "patterns": ["Pattern"], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}]}`)
				}))
			})
			It(`Invoke GetWorkspace successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.GetWorkspace(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetWorkspaceOptions model
				getWorkspaceOptionsModel := new(assistantv1.GetWorkspaceOptions)
				getWorkspaceOptionsModel.WorkspaceID = core.StringPtr("testString")
				getWorkspaceOptionsModel.Export = core.BoolPtr(false)
				getWorkspaceOptionsModel.IncludeAudit = core.BoolPtr(false)
				getWorkspaceOptionsModel.Sort = core.StringPtr("stable")
				getWorkspaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.GetWorkspace(getWorkspaceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetWorkspace with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the GetWorkspaceOptions model
				getWorkspaceOptionsModel := new(assistantv1.GetWorkspaceOptions)
				getWorkspaceOptionsModel.WorkspaceID = core.StringPtr("testString")
				getWorkspaceOptionsModel.Export = core.BoolPtr(false)
				getWorkspaceOptionsModel.IncludeAudit = core.BoolPtr(false)
				getWorkspaceOptionsModel.Sort = core.StringPtr("stable")
				getWorkspaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.GetWorkspace(getWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetWorkspaceOptions model with no property values
				getWorkspaceOptionsModelNew := new(assistantv1.GetWorkspaceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = assistantService.GetWorkspace(getWorkspaceOptionsModelNew)
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
			It(`Invoke GetWorkspace successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the GetWorkspaceOptions model
				getWorkspaceOptionsModel := new(assistantv1.GetWorkspaceOptions)
				getWorkspaceOptionsModel.WorkspaceID = core.StringPtr("testString")
				getWorkspaceOptionsModel.Export = core.BoolPtr(false)
				getWorkspaceOptionsModel.IncludeAudit = core.BoolPtr(false)
				getWorkspaceOptionsModel.Sort = core.StringPtr("stable")
				getWorkspaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.GetWorkspace(getWorkspaceOptionsModel)
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
	Describe(`UpdateWorkspace(updateWorkspaceOptions *UpdateWorkspaceOptions) - Operation response error`, func() {
		version := "testString"
		updateWorkspacePath := "/v1/workspaces/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateWorkspacePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for append query parameter
					// TODO: Add check for include_audit query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateWorkspace with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ResponseGenericChannel model
				responseGenericChannelModel := new(assistantv1.ResponseGenericChannel)
				responseGenericChannelModel.Channel = core.StringPtr("chat")

				// Construct an instance of the DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo model
				dialogNodeOutputGenericModel := new(assistantv1.DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo)
				dialogNodeOutputGenericModel.ResponseType = core.StringPtr("video")
				dialogNodeOutputGenericModel.Source = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Title = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Description = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Channels = []assistantv1.ResponseGenericChannel{*responseGenericChannelModel}
				dialogNodeOutputGenericModel.ChannelOptions = map[string]interface{}{"anyKey": "anyValue"}
				dialogNodeOutputGenericModel.AltText = core.StringPtr("testString")

				// Construct an instance of the DialogNodeOutputModifiers model
				dialogNodeOutputModifiersModel := new(assistantv1.DialogNodeOutputModifiers)
				dialogNodeOutputModifiersModel.Overwrite = core.BoolPtr(true)

				// Construct an instance of the DialogNodeOutput model
				dialogNodeOutputModel := new(assistantv1.DialogNodeOutput)
				dialogNodeOutputModel.Generic = []assistantv1.DialogNodeOutputGenericIntf{dialogNodeOutputGenericModel}
				dialogNodeOutputModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeOutputModel.Modifiers = dialogNodeOutputModifiersModel
				dialogNodeOutputModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeContext model
				dialogNodeContextModel := new(assistantv1.DialogNodeContext)
				dialogNodeContextModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeContextModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeNextStep model
				dialogNodeNextStepModel := new(assistantv1.DialogNodeNextStep)
				dialogNodeNextStepModel.Behavior = core.StringPtr("get_user_input")
				dialogNodeNextStepModel.DialogNode = core.StringPtr("testString")
				dialogNodeNextStepModel.Selector = core.StringPtr("condition")

				// Construct an instance of the DialogNodeAction model
				dialogNodeActionModel := new(assistantv1.DialogNodeAction)
				dialogNodeActionModel.Name = core.StringPtr("testString")
				dialogNodeActionModel.Type = core.StringPtr("client")
				dialogNodeActionModel.Parameters = make(map[string]interface{})
				dialogNodeActionModel.ResultVariable = core.StringPtr("testString")
				dialogNodeActionModel.Credentials = core.StringPtr("testString")

				// Construct an instance of the DialogNode model
				dialogNodeModel := new(assistantv1.DialogNode)
				dialogNodeModel.DialogNode = core.StringPtr("testString")
				dialogNodeModel.Description = core.StringPtr("testString")
				dialogNodeModel.Conditions = core.StringPtr("testString")
				dialogNodeModel.Parent = core.StringPtr("testString")
				dialogNodeModel.PreviousSibling = core.StringPtr("testString")
				dialogNodeModel.Output = dialogNodeOutputModel
				dialogNodeModel.Context = dialogNodeContextModel
				dialogNodeModel.Metadata = make(map[string]interface{})
				dialogNodeModel.NextStep = dialogNodeNextStepModel
				dialogNodeModel.Title = core.StringPtr("testString")
				dialogNodeModel.Type = core.StringPtr("standard")
				dialogNodeModel.EventName = core.StringPtr("focus")
				dialogNodeModel.Variable = core.StringPtr("testString")
				dialogNodeModel.Actions = []assistantv1.DialogNodeAction{*dialogNodeActionModel}
				dialogNodeModel.DigressIn = core.StringPtr("not_available")
				dialogNodeModel.DigressOut = core.StringPtr("allow_returning")
				dialogNodeModel.DigressOutSlots = core.StringPtr("not_allowed")
				dialogNodeModel.UserLabel = core.StringPtr("testString")
				dialogNodeModel.DisambiguationOptOut = core.BoolPtr(false)

				// Construct an instance of the Counterexample model
				counterexampleModel := new(assistantv1.Counterexample)
				counterexampleModel.Text = core.StringPtr("testString")

				// Construct an instance of the WorkspaceSystemSettingsTooling model
				workspaceSystemSettingsToolingModel := new(assistantv1.WorkspaceSystemSettingsTooling)
				workspaceSystemSettingsToolingModel.StoreGenericResponses = core.BoolPtr(true)

				// Construct an instance of the WorkspaceSystemSettingsDisambiguation model
				workspaceSystemSettingsDisambiguationModel := new(assistantv1.WorkspaceSystemSettingsDisambiguation)
				workspaceSystemSettingsDisambiguationModel.Prompt = core.StringPtr("testString")
				workspaceSystemSettingsDisambiguationModel.NoneOfTheAbovePrompt = core.StringPtr("testString")
				workspaceSystemSettingsDisambiguationModel.Enabled = core.BoolPtr(false)
				workspaceSystemSettingsDisambiguationModel.Sensitivity = core.StringPtr("auto")
				workspaceSystemSettingsDisambiguationModel.Randomize = core.BoolPtr(true)
				workspaceSystemSettingsDisambiguationModel.MaxSuggestions = core.Int64Ptr(int64(1))
				workspaceSystemSettingsDisambiguationModel.SuggestionTextPolicy = core.StringPtr("testString")

				// Construct an instance of the WorkspaceSystemSettingsSystemEntities model
				workspaceSystemSettingsSystemEntitiesModel := new(assistantv1.WorkspaceSystemSettingsSystemEntities)
				workspaceSystemSettingsSystemEntitiesModel.Enabled = core.BoolPtr(false)

				// Construct an instance of the WorkspaceSystemSettingsOffTopic model
				workspaceSystemSettingsOffTopicModel := new(assistantv1.WorkspaceSystemSettingsOffTopic)
				workspaceSystemSettingsOffTopicModel.Enabled = core.BoolPtr(false)

				// Construct an instance of the WorkspaceSystemSettings model
				workspaceSystemSettingsModel := new(assistantv1.WorkspaceSystemSettings)
				workspaceSystemSettingsModel.Tooling = workspaceSystemSettingsToolingModel
				workspaceSystemSettingsModel.Disambiguation = workspaceSystemSettingsDisambiguationModel
				workspaceSystemSettingsModel.HumanAgentAssist = make(map[string]interface{})
				workspaceSystemSettingsModel.SpellingSuggestions = core.BoolPtr(false)
				workspaceSystemSettingsModel.SpellingAutoCorrect = core.BoolPtr(false)
				workspaceSystemSettingsModel.SystemEntities = workspaceSystemSettingsSystemEntitiesModel
				workspaceSystemSettingsModel.OffTopic = workspaceSystemSettingsOffTopicModel
				workspaceSystemSettingsModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the WebhookHeader model
				webhookHeaderModel := new(assistantv1.WebhookHeader)
				webhookHeaderModel.Name = core.StringPtr("testString")
				webhookHeaderModel.Value = core.StringPtr("testString")

				// Construct an instance of the Webhook model
				webhookModel := new(assistantv1.Webhook)
				webhookModel.URL = core.StringPtr("testString")
				webhookModel.Name = core.StringPtr("testString")
				webhookModel.HeadersVar = []assistantv1.WebhookHeader{*webhookHeaderModel}

				// Construct an instance of the Mention model
				mentionModel := new(assistantv1.Mention)
				mentionModel.Entity = core.StringPtr("testString")
				mentionModel.Location = []int64{int64(38)}

				// Construct an instance of the Example model
				exampleModel := new(assistantv1.Example)
				exampleModel.Text = core.StringPtr("testString")
				exampleModel.Mentions = []assistantv1.Mention{*mentionModel}

				// Construct an instance of the CreateIntent model
				createIntentModel := new(assistantv1.CreateIntent)
				createIntentModel.Intent = core.StringPtr("testString")
				createIntentModel.Description = core.StringPtr("testString")
				createIntentModel.Examples = []assistantv1.Example{*exampleModel}

				// Construct an instance of the CreateValue model
				createValueModel := new(assistantv1.CreateValue)
				createValueModel.Value = core.StringPtr("testString")
				createValueModel.Metadata = make(map[string]interface{})
				createValueModel.Type = core.StringPtr("synonyms")
				createValueModel.Synonyms = []string{"testString"}
				createValueModel.Patterns = []string{"testString"}

				// Construct an instance of the CreateEntity model
				createEntityModel := new(assistantv1.CreateEntity)
				createEntityModel.Entity = core.StringPtr("testString")
				createEntityModel.Description = core.StringPtr("testString")
				createEntityModel.Metadata = make(map[string]interface{})
				createEntityModel.FuzzyMatch = core.BoolPtr(true)
				createEntityModel.Values = []assistantv1.CreateValue{*createValueModel}

				// Construct an instance of the UpdateWorkspaceOptions model
				updateWorkspaceOptionsModel := new(assistantv1.UpdateWorkspaceOptions)
				updateWorkspaceOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateWorkspaceOptionsModel.Name = core.StringPtr("testString")
				updateWorkspaceOptionsModel.Description = core.StringPtr("testString")
				updateWorkspaceOptionsModel.Language = core.StringPtr("testString")
				updateWorkspaceOptionsModel.DialogNodes = []assistantv1.DialogNode{*dialogNodeModel}
				updateWorkspaceOptionsModel.Counterexamples = []assistantv1.Counterexample{*counterexampleModel}
				updateWorkspaceOptionsModel.Metadata = make(map[string]interface{})
				updateWorkspaceOptionsModel.LearningOptOut = core.BoolPtr(false)
				updateWorkspaceOptionsModel.SystemSettings = workspaceSystemSettingsModel
				updateWorkspaceOptionsModel.Webhooks = []assistantv1.Webhook{*webhookModel}
				updateWorkspaceOptionsModel.Intents = []assistantv1.CreateIntent{*createIntentModel}
				updateWorkspaceOptionsModel.Entities = []assistantv1.CreateEntity{*createEntityModel}
				updateWorkspaceOptionsModel.Append = core.BoolPtr(false)
				updateWorkspaceOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateWorkspaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.UpdateWorkspace(updateWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.UpdateWorkspace(updateWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateWorkspace(updateWorkspaceOptions *UpdateWorkspaceOptions)`, func() {
		version := "testString"
		updateWorkspacePath := "/v1/workspaces/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateWorkspacePath))
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
					// TODO: Add check for append query parameter
					// TODO: Add check for include_audit query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description", "language": "Language", "workspace_id": "WorkspaceID", "dialog_nodes": [{"dialog_node": "DialogNode", "description": "Description", "conditions": "Conditions", "parent": "Parent", "previous_sibling": "PreviousSibling", "output": {"generic": [{"response_type": "video", "source": "Source", "title": "Title", "description": "Description", "channels": [{"channel": "chat"}], "channel_options": {"anyKey": "anyValue"}, "alt_text": "AltText"}], "integrations": {"mapKey": {"mapKey": "anyValue"}}, "modifiers": {"overwrite": true}}, "context": {"integrations": {"mapKey": {"mapKey": "anyValue"}}}, "metadata": {"mapKey": "anyValue"}, "next_step": {"behavior": "get_user_input", "dialog_node": "DialogNode", "selector": "condition"}, "title": "Title", "type": "standard", "event_name": "focus", "variable": "Variable", "actions": [{"name": "Name", "type": "client", "parameters": {"mapKey": "anyValue"}, "result_variable": "ResultVariable", "credentials": "Credentials"}], "digress_in": "not_available", "digress_out": "allow_returning", "digress_out_slots": "not_allowed", "user_label": "UserLabel", "disambiguation_opt_out": false, "disabled": true, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}], "counterexamples": [{"text": "Text", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "metadata": {"mapKey": "anyValue"}, "learning_opt_out": false, "system_settings": {"tooling": {"store_generic_responses": false}, "disambiguation": {"prompt": "Prompt", "none_of_the_above_prompt": "NoneOfTheAbovePrompt", "enabled": false, "sensitivity": "auto", "randomize": false, "max_suggestions": 1, "suggestion_text_policy": "SuggestionTextPolicy"}, "human_agent_assist": {"mapKey": "anyValue"}, "spelling_suggestions": false, "spelling_auto_correct": false, "system_entities": {"enabled": false}, "off_topic": {"enabled": false}}, "status": "Non Existent", "webhooks": [{"url": "URL", "name": "Name", "headers": [{"name": "Name", "value": "Value"}]}], "intents": [{"intent": "Intent", "description": "Description", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "examples": [{"text": "Text", "mentions": [{"entity": "Entity", "location": [8]}], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}], "entities": [{"entity": "Entity", "description": "Description", "metadata": {"mapKey": "anyValue"}, "fuzzy_match": true, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "values": [{"value": "Value", "metadata": {"mapKey": "anyValue"}, "type": "synonyms", "synonyms": ["Synonym"], "patterns": ["Pattern"], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}]}`)
				}))
			})
			It(`Invoke UpdateWorkspace successfully with retries`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the ResponseGenericChannel model
				responseGenericChannelModel := new(assistantv1.ResponseGenericChannel)
				responseGenericChannelModel.Channel = core.StringPtr("chat")

				// Construct an instance of the DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo model
				dialogNodeOutputGenericModel := new(assistantv1.DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo)
				dialogNodeOutputGenericModel.ResponseType = core.StringPtr("video")
				dialogNodeOutputGenericModel.Source = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Title = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Description = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Channels = []assistantv1.ResponseGenericChannel{*responseGenericChannelModel}
				dialogNodeOutputGenericModel.ChannelOptions = map[string]interface{}{"anyKey": "anyValue"}
				dialogNodeOutputGenericModel.AltText = core.StringPtr("testString")

				// Construct an instance of the DialogNodeOutputModifiers model
				dialogNodeOutputModifiersModel := new(assistantv1.DialogNodeOutputModifiers)
				dialogNodeOutputModifiersModel.Overwrite = core.BoolPtr(true)

				// Construct an instance of the DialogNodeOutput model
				dialogNodeOutputModel := new(assistantv1.DialogNodeOutput)
				dialogNodeOutputModel.Generic = []assistantv1.DialogNodeOutputGenericIntf{dialogNodeOutputGenericModel}
				dialogNodeOutputModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeOutputModel.Modifiers = dialogNodeOutputModifiersModel
				dialogNodeOutputModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeContext model
				dialogNodeContextModel := new(assistantv1.DialogNodeContext)
				dialogNodeContextModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeContextModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeNextStep model
				dialogNodeNextStepModel := new(assistantv1.DialogNodeNextStep)
				dialogNodeNextStepModel.Behavior = core.StringPtr("get_user_input")
				dialogNodeNextStepModel.DialogNode = core.StringPtr("testString")
				dialogNodeNextStepModel.Selector = core.StringPtr("condition")

				// Construct an instance of the DialogNodeAction model
				dialogNodeActionModel := new(assistantv1.DialogNodeAction)
				dialogNodeActionModel.Name = core.StringPtr("testString")
				dialogNodeActionModel.Type = core.StringPtr("client")
				dialogNodeActionModel.Parameters = make(map[string]interface{})
				dialogNodeActionModel.ResultVariable = core.StringPtr("testString")
				dialogNodeActionModel.Credentials = core.StringPtr("testString")

				// Construct an instance of the DialogNode model
				dialogNodeModel := new(assistantv1.DialogNode)
				dialogNodeModel.DialogNode = core.StringPtr("testString")
				dialogNodeModel.Description = core.StringPtr("testString")
				dialogNodeModel.Conditions = core.StringPtr("testString")
				dialogNodeModel.Parent = core.StringPtr("testString")
				dialogNodeModel.PreviousSibling = core.StringPtr("testString")
				dialogNodeModel.Output = dialogNodeOutputModel
				dialogNodeModel.Context = dialogNodeContextModel
				dialogNodeModel.Metadata = make(map[string]interface{})
				dialogNodeModel.NextStep = dialogNodeNextStepModel
				dialogNodeModel.Title = core.StringPtr("testString")
				dialogNodeModel.Type = core.StringPtr("standard")
				dialogNodeModel.EventName = core.StringPtr("focus")
				dialogNodeModel.Variable = core.StringPtr("testString")
				dialogNodeModel.Actions = []assistantv1.DialogNodeAction{*dialogNodeActionModel}
				dialogNodeModel.DigressIn = core.StringPtr("not_available")
				dialogNodeModel.DigressOut = core.StringPtr("allow_returning")
				dialogNodeModel.DigressOutSlots = core.StringPtr("not_allowed")
				dialogNodeModel.UserLabel = core.StringPtr("testString")
				dialogNodeModel.DisambiguationOptOut = core.BoolPtr(false)

				// Construct an instance of the Counterexample model
				counterexampleModel := new(assistantv1.Counterexample)
				counterexampleModel.Text = core.StringPtr("testString")

				// Construct an instance of the WorkspaceSystemSettingsTooling model
				workspaceSystemSettingsToolingModel := new(assistantv1.WorkspaceSystemSettingsTooling)
				workspaceSystemSettingsToolingModel.StoreGenericResponses = core.BoolPtr(true)

				// Construct an instance of the WorkspaceSystemSettingsDisambiguation model
				workspaceSystemSettingsDisambiguationModel := new(assistantv1.WorkspaceSystemSettingsDisambiguation)
				workspaceSystemSettingsDisambiguationModel.Prompt = core.StringPtr("testString")
				workspaceSystemSettingsDisambiguationModel.NoneOfTheAbovePrompt = core.StringPtr("testString")
				workspaceSystemSettingsDisambiguationModel.Enabled = core.BoolPtr(false)
				workspaceSystemSettingsDisambiguationModel.Sensitivity = core.StringPtr("auto")
				workspaceSystemSettingsDisambiguationModel.Randomize = core.BoolPtr(true)
				workspaceSystemSettingsDisambiguationModel.MaxSuggestions = core.Int64Ptr(int64(1))
				workspaceSystemSettingsDisambiguationModel.SuggestionTextPolicy = core.StringPtr("testString")

				// Construct an instance of the WorkspaceSystemSettingsSystemEntities model
				workspaceSystemSettingsSystemEntitiesModel := new(assistantv1.WorkspaceSystemSettingsSystemEntities)
				workspaceSystemSettingsSystemEntitiesModel.Enabled = core.BoolPtr(false)

				// Construct an instance of the WorkspaceSystemSettingsOffTopic model
				workspaceSystemSettingsOffTopicModel := new(assistantv1.WorkspaceSystemSettingsOffTopic)
				workspaceSystemSettingsOffTopicModel.Enabled = core.BoolPtr(false)

				// Construct an instance of the WorkspaceSystemSettings model
				workspaceSystemSettingsModel := new(assistantv1.WorkspaceSystemSettings)
				workspaceSystemSettingsModel.Tooling = workspaceSystemSettingsToolingModel
				workspaceSystemSettingsModel.Disambiguation = workspaceSystemSettingsDisambiguationModel
				workspaceSystemSettingsModel.HumanAgentAssist = make(map[string]interface{})
				workspaceSystemSettingsModel.SpellingSuggestions = core.BoolPtr(false)
				workspaceSystemSettingsModel.SpellingAutoCorrect = core.BoolPtr(false)
				workspaceSystemSettingsModel.SystemEntities = workspaceSystemSettingsSystemEntitiesModel
				workspaceSystemSettingsModel.OffTopic = workspaceSystemSettingsOffTopicModel
				workspaceSystemSettingsModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the WebhookHeader model
				webhookHeaderModel := new(assistantv1.WebhookHeader)
				webhookHeaderModel.Name = core.StringPtr("testString")
				webhookHeaderModel.Value = core.StringPtr("testString")

				// Construct an instance of the Webhook model
				webhookModel := new(assistantv1.Webhook)
				webhookModel.URL = core.StringPtr("testString")
				webhookModel.Name = core.StringPtr("testString")
				webhookModel.HeadersVar = []assistantv1.WebhookHeader{*webhookHeaderModel}

				// Construct an instance of the Mention model
				mentionModel := new(assistantv1.Mention)
				mentionModel.Entity = core.StringPtr("testString")
				mentionModel.Location = []int64{int64(38)}

				// Construct an instance of the Example model
				exampleModel := new(assistantv1.Example)
				exampleModel.Text = core.StringPtr("testString")
				exampleModel.Mentions = []assistantv1.Mention{*mentionModel}

				// Construct an instance of the CreateIntent model
				createIntentModel := new(assistantv1.CreateIntent)
				createIntentModel.Intent = core.StringPtr("testString")
				createIntentModel.Description = core.StringPtr("testString")
				createIntentModel.Examples = []assistantv1.Example{*exampleModel}

				// Construct an instance of the CreateValue model
				createValueModel := new(assistantv1.CreateValue)
				createValueModel.Value = core.StringPtr("testString")
				createValueModel.Metadata = make(map[string]interface{})
				createValueModel.Type = core.StringPtr("synonyms")
				createValueModel.Synonyms = []string{"testString"}
				createValueModel.Patterns = []string{"testString"}

				// Construct an instance of the CreateEntity model
				createEntityModel := new(assistantv1.CreateEntity)
				createEntityModel.Entity = core.StringPtr("testString")
				createEntityModel.Description = core.StringPtr("testString")
				createEntityModel.Metadata = make(map[string]interface{})
				createEntityModel.FuzzyMatch = core.BoolPtr(true)
				createEntityModel.Values = []assistantv1.CreateValue{*createValueModel}

				// Construct an instance of the UpdateWorkspaceOptions model
				updateWorkspaceOptionsModel := new(assistantv1.UpdateWorkspaceOptions)
				updateWorkspaceOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateWorkspaceOptionsModel.Name = core.StringPtr("testString")
				updateWorkspaceOptionsModel.Description = core.StringPtr("testString")
				updateWorkspaceOptionsModel.Language = core.StringPtr("testString")
				updateWorkspaceOptionsModel.DialogNodes = []assistantv1.DialogNode{*dialogNodeModel}
				updateWorkspaceOptionsModel.Counterexamples = []assistantv1.Counterexample{*counterexampleModel}
				updateWorkspaceOptionsModel.Metadata = make(map[string]interface{})
				updateWorkspaceOptionsModel.LearningOptOut = core.BoolPtr(false)
				updateWorkspaceOptionsModel.SystemSettings = workspaceSystemSettingsModel
				updateWorkspaceOptionsModel.Webhooks = []assistantv1.Webhook{*webhookModel}
				updateWorkspaceOptionsModel.Intents = []assistantv1.CreateIntent{*createIntentModel}
				updateWorkspaceOptionsModel.Entities = []assistantv1.CreateEntity{*createEntityModel}
				updateWorkspaceOptionsModel.Append = core.BoolPtr(false)
				updateWorkspaceOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateWorkspaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.UpdateWorkspaceWithContext(ctx, updateWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.UpdateWorkspace(updateWorkspaceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.UpdateWorkspaceWithContext(ctx, updateWorkspaceOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateWorkspacePath))
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
					// TODO: Add check for append query parameter
					// TODO: Add check for include_audit query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description", "language": "Language", "workspace_id": "WorkspaceID", "dialog_nodes": [{"dialog_node": "DialogNode", "description": "Description", "conditions": "Conditions", "parent": "Parent", "previous_sibling": "PreviousSibling", "output": {"generic": [{"response_type": "video", "source": "Source", "title": "Title", "description": "Description", "channels": [{"channel": "chat"}], "channel_options": {"anyKey": "anyValue"}, "alt_text": "AltText"}], "integrations": {"mapKey": {"mapKey": "anyValue"}}, "modifiers": {"overwrite": true}}, "context": {"integrations": {"mapKey": {"mapKey": "anyValue"}}}, "metadata": {"mapKey": "anyValue"}, "next_step": {"behavior": "get_user_input", "dialog_node": "DialogNode", "selector": "condition"}, "title": "Title", "type": "standard", "event_name": "focus", "variable": "Variable", "actions": [{"name": "Name", "type": "client", "parameters": {"mapKey": "anyValue"}, "result_variable": "ResultVariable", "credentials": "Credentials"}], "digress_in": "not_available", "digress_out": "allow_returning", "digress_out_slots": "not_allowed", "user_label": "UserLabel", "disambiguation_opt_out": false, "disabled": true, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}], "counterexamples": [{"text": "Text", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "metadata": {"mapKey": "anyValue"}, "learning_opt_out": false, "system_settings": {"tooling": {"store_generic_responses": false}, "disambiguation": {"prompt": "Prompt", "none_of_the_above_prompt": "NoneOfTheAbovePrompt", "enabled": false, "sensitivity": "auto", "randomize": false, "max_suggestions": 1, "suggestion_text_policy": "SuggestionTextPolicy"}, "human_agent_assist": {"mapKey": "anyValue"}, "spelling_suggestions": false, "spelling_auto_correct": false, "system_entities": {"enabled": false}, "off_topic": {"enabled": false}}, "status": "Non Existent", "webhooks": [{"url": "URL", "name": "Name", "headers": [{"name": "Name", "value": "Value"}]}], "intents": [{"intent": "Intent", "description": "Description", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "examples": [{"text": "Text", "mentions": [{"entity": "Entity", "location": [8]}], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}], "entities": [{"entity": "Entity", "description": "Description", "metadata": {"mapKey": "anyValue"}, "fuzzy_match": true, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "values": [{"value": "Value", "metadata": {"mapKey": "anyValue"}, "type": "synonyms", "synonyms": ["Synonym"], "patterns": ["Pattern"], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}]}`)
				}))
			})
			It(`Invoke UpdateWorkspace successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.UpdateWorkspace(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ResponseGenericChannel model
				responseGenericChannelModel := new(assistantv1.ResponseGenericChannel)
				responseGenericChannelModel.Channel = core.StringPtr("chat")

				// Construct an instance of the DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo model
				dialogNodeOutputGenericModel := new(assistantv1.DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo)
				dialogNodeOutputGenericModel.ResponseType = core.StringPtr("video")
				dialogNodeOutputGenericModel.Source = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Title = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Description = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Channels = []assistantv1.ResponseGenericChannel{*responseGenericChannelModel}
				dialogNodeOutputGenericModel.ChannelOptions = map[string]interface{}{"anyKey": "anyValue"}
				dialogNodeOutputGenericModel.AltText = core.StringPtr("testString")

				// Construct an instance of the DialogNodeOutputModifiers model
				dialogNodeOutputModifiersModel := new(assistantv1.DialogNodeOutputModifiers)
				dialogNodeOutputModifiersModel.Overwrite = core.BoolPtr(true)

				// Construct an instance of the DialogNodeOutput model
				dialogNodeOutputModel := new(assistantv1.DialogNodeOutput)
				dialogNodeOutputModel.Generic = []assistantv1.DialogNodeOutputGenericIntf{dialogNodeOutputGenericModel}
				dialogNodeOutputModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeOutputModel.Modifiers = dialogNodeOutputModifiersModel
				dialogNodeOutputModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeContext model
				dialogNodeContextModel := new(assistantv1.DialogNodeContext)
				dialogNodeContextModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeContextModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeNextStep model
				dialogNodeNextStepModel := new(assistantv1.DialogNodeNextStep)
				dialogNodeNextStepModel.Behavior = core.StringPtr("get_user_input")
				dialogNodeNextStepModel.DialogNode = core.StringPtr("testString")
				dialogNodeNextStepModel.Selector = core.StringPtr("condition")

				// Construct an instance of the DialogNodeAction model
				dialogNodeActionModel := new(assistantv1.DialogNodeAction)
				dialogNodeActionModel.Name = core.StringPtr("testString")
				dialogNodeActionModel.Type = core.StringPtr("client")
				dialogNodeActionModel.Parameters = make(map[string]interface{})
				dialogNodeActionModel.ResultVariable = core.StringPtr("testString")
				dialogNodeActionModel.Credentials = core.StringPtr("testString")

				// Construct an instance of the DialogNode model
				dialogNodeModel := new(assistantv1.DialogNode)
				dialogNodeModel.DialogNode = core.StringPtr("testString")
				dialogNodeModel.Description = core.StringPtr("testString")
				dialogNodeModel.Conditions = core.StringPtr("testString")
				dialogNodeModel.Parent = core.StringPtr("testString")
				dialogNodeModel.PreviousSibling = core.StringPtr("testString")
				dialogNodeModel.Output = dialogNodeOutputModel
				dialogNodeModel.Context = dialogNodeContextModel
				dialogNodeModel.Metadata = make(map[string]interface{})
				dialogNodeModel.NextStep = dialogNodeNextStepModel
				dialogNodeModel.Title = core.StringPtr("testString")
				dialogNodeModel.Type = core.StringPtr("standard")
				dialogNodeModel.EventName = core.StringPtr("focus")
				dialogNodeModel.Variable = core.StringPtr("testString")
				dialogNodeModel.Actions = []assistantv1.DialogNodeAction{*dialogNodeActionModel}
				dialogNodeModel.DigressIn = core.StringPtr("not_available")
				dialogNodeModel.DigressOut = core.StringPtr("allow_returning")
				dialogNodeModel.DigressOutSlots = core.StringPtr("not_allowed")
				dialogNodeModel.UserLabel = core.StringPtr("testString")
				dialogNodeModel.DisambiguationOptOut = core.BoolPtr(false)

				// Construct an instance of the Counterexample model
				counterexampleModel := new(assistantv1.Counterexample)
				counterexampleModel.Text = core.StringPtr("testString")

				// Construct an instance of the WorkspaceSystemSettingsTooling model
				workspaceSystemSettingsToolingModel := new(assistantv1.WorkspaceSystemSettingsTooling)
				workspaceSystemSettingsToolingModel.StoreGenericResponses = core.BoolPtr(true)

				// Construct an instance of the WorkspaceSystemSettingsDisambiguation model
				workspaceSystemSettingsDisambiguationModel := new(assistantv1.WorkspaceSystemSettingsDisambiguation)
				workspaceSystemSettingsDisambiguationModel.Prompt = core.StringPtr("testString")
				workspaceSystemSettingsDisambiguationModel.NoneOfTheAbovePrompt = core.StringPtr("testString")
				workspaceSystemSettingsDisambiguationModel.Enabled = core.BoolPtr(false)
				workspaceSystemSettingsDisambiguationModel.Sensitivity = core.StringPtr("auto")
				workspaceSystemSettingsDisambiguationModel.Randomize = core.BoolPtr(true)
				workspaceSystemSettingsDisambiguationModel.MaxSuggestions = core.Int64Ptr(int64(1))
				workspaceSystemSettingsDisambiguationModel.SuggestionTextPolicy = core.StringPtr("testString")

				// Construct an instance of the WorkspaceSystemSettingsSystemEntities model
				workspaceSystemSettingsSystemEntitiesModel := new(assistantv1.WorkspaceSystemSettingsSystemEntities)
				workspaceSystemSettingsSystemEntitiesModel.Enabled = core.BoolPtr(false)

				// Construct an instance of the WorkspaceSystemSettingsOffTopic model
				workspaceSystemSettingsOffTopicModel := new(assistantv1.WorkspaceSystemSettingsOffTopic)
				workspaceSystemSettingsOffTopicModel.Enabled = core.BoolPtr(false)

				// Construct an instance of the WorkspaceSystemSettings model
				workspaceSystemSettingsModel := new(assistantv1.WorkspaceSystemSettings)
				workspaceSystemSettingsModel.Tooling = workspaceSystemSettingsToolingModel
				workspaceSystemSettingsModel.Disambiguation = workspaceSystemSettingsDisambiguationModel
				workspaceSystemSettingsModel.HumanAgentAssist = make(map[string]interface{})
				workspaceSystemSettingsModel.SpellingSuggestions = core.BoolPtr(false)
				workspaceSystemSettingsModel.SpellingAutoCorrect = core.BoolPtr(false)
				workspaceSystemSettingsModel.SystemEntities = workspaceSystemSettingsSystemEntitiesModel
				workspaceSystemSettingsModel.OffTopic = workspaceSystemSettingsOffTopicModel
				workspaceSystemSettingsModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the WebhookHeader model
				webhookHeaderModel := new(assistantv1.WebhookHeader)
				webhookHeaderModel.Name = core.StringPtr("testString")
				webhookHeaderModel.Value = core.StringPtr("testString")

				// Construct an instance of the Webhook model
				webhookModel := new(assistantv1.Webhook)
				webhookModel.URL = core.StringPtr("testString")
				webhookModel.Name = core.StringPtr("testString")
				webhookModel.HeadersVar = []assistantv1.WebhookHeader{*webhookHeaderModel}

				// Construct an instance of the Mention model
				mentionModel := new(assistantv1.Mention)
				mentionModel.Entity = core.StringPtr("testString")
				mentionModel.Location = []int64{int64(38)}

				// Construct an instance of the Example model
				exampleModel := new(assistantv1.Example)
				exampleModel.Text = core.StringPtr("testString")
				exampleModel.Mentions = []assistantv1.Mention{*mentionModel}

				// Construct an instance of the CreateIntent model
				createIntentModel := new(assistantv1.CreateIntent)
				createIntentModel.Intent = core.StringPtr("testString")
				createIntentModel.Description = core.StringPtr("testString")
				createIntentModel.Examples = []assistantv1.Example{*exampleModel}

				// Construct an instance of the CreateValue model
				createValueModel := new(assistantv1.CreateValue)
				createValueModel.Value = core.StringPtr("testString")
				createValueModel.Metadata = make(map[string]interface{})
				createValueModel.Type = core.StringPtr("synonyms")
				createValueModel.Synonyms = []string{"testString"}
				createValueModel.Patterns = []string{"testString"}

				// Construct an instance of the CreateEntity model
				createEntityModel := new(assistantv1.CreateEntity)
				createEntityModel.Entity = core.StringPtr("testString")
				createEntityModel.Description = core.StringPtr("testString")
				createEntityModel.Metadata = make(map[string]interface{})
				createEntityModel.FuzzyMatch = core.BoolPtr(true)
				createEntityModel.Values = []assistantv1.CreateValue{*createValueModel}

				// Construct an instance of the UpdateWorkspaceOptions model
				updateWorkspaceOptionsModel := new(assistantv1.UpdateWorkspaceOptions)
				updateWorkspaceOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateWorkspaceOptionsModel.Name = core.StringPtr("testString")
				updateWorkspaceOptionsModel.Description = core.StringPtr("testString")
				updateWorkspaceOptionsModel.Language = core.StringPtr("testString")
				updateWorkspaceOptionsModel.DialogNodes = []assistantv1.DialogNode{*dialogNodeModel}
				updateWorkspaceOptionsModel.Counterexamples = []assistantv1.Counterexample{*counterexampleModel}
				updateWorkspaceOptionsModel.Metadata = make(map[string]interface{})
				updateWorkspaceOptionsModel.LearningOptOut = core.BoolPtr(false)
				updateWorkspaceOptionsModel.SystemSettings = workspaceSystemSettingsModel
				updateWorkspaceOptionsModel.Webhooks = []assistantv1.Webhook{*webhookModel}
				updateWorkspaceOptionsModel.Intents = []assistantv1.CreateIntent{*createIntentModel}
				updateWorkspaceOptionsModel.Entities = []assistantv1.CreateEntity{*createEntityModel}
				updateWorkspaceOptionsModel.Append = core.BoolPtr(false)
				updateWorkspaceOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateWorkspaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.UpdateWorkspace(updateWorkspaceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateWorkspace with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ResponseGenericChannel model
				responseGenericChannelModel := new(assistantv1.ResponseGenericChannel)
				responseGenericChannelModel.Channel = core.StringPtr("chat")

				// Construct an instance of the DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo model
				dialogNodeOutputGenericModel := new(assistantv1.DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo)
				dialogNodeOutputGenericModel.ResponseType = core.StringPtr("video")
				dialogNodeOutputGenericModel.Source = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Title = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Description = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Channels = []assistantv1.ResponseGenericChannel{*responseGenericChannelModel}
				dialogNodeOutputGenericModel.ChannelOptions = map[string]interface{}{"anyKey": "anyValue"}
				dialogNodeOutputGenericModel.AltText = core.StringPtr("testString")

				// Construct an instance of the DialogNodeOutputModifiers model
				dialogNodeOutputModifiersModel := new(assistantv1.DialogNodeOutputModifiers)
				dialogNodeOutputModifiersModel.Overwrite = core.BoolPtr(true)

				// Construct an instance of the DialogNodeOutput model
				dialogNodeOutputModel := new(assistantv1.DialogNodeOutput)
				dialogNodeOutputModel.Generic = []assistantv1.DialogNodeOutputGenericIntf{dialogNodeOutputGenericModel}
				dialogNodeOutputModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeOutputModel.Modifiers = dialogNodeOutputModifiersModel
				dialogNodeOutputModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeContext model
				dialogNodeContextModel := new(assistantv1.DialogNodeContext)
				dialogNodeContextModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeContextModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeNextStep model
				dialogNodeNextStepModel := new(assistantv1.DialogNodeNextStep)
				dialogNodeNextStepModel.Behavior = core.StringPtr("get_user_input")
				dialogNodeNextStepModel.DialogNode = core.StringPtr("testString")
				dialogNodeNextStepModel.Selector = core.StringPtr("condition")

				// Construct an instance of the DialogNodeAction model
				dialogNodeActionModel := new(assistantv1.DialogNodeAction)
				dialogNodeActionModel.Name = core.StringPtr("testString")
				dialogNodeActionModel.Type = core.StringPtr("client")
				dialogNodeActionModel.Parameters = make(map[string]interface{})
				dialogNodeActionModel.ResultVariable = core.StringPtr("testString")
				dialogNodeActionModel.Credentials = core.StringPtr("testString")

				// Construct an instance of the DialogNode model
				dialogNodeModel := new(assistantv1.DialogNode)
				dialogNodeModel.DialogNode = core.StringPtr("testString")
				dialogNodeModel.Description = core.StringPtr("testString")
				dialogNodeModel.Conditions = core.StringPtr("testString")
				dialogNodeModel.Parent = core.StringPtr("testString")
				dialogNodeModel.PreviousSibling = core.StringPtr("testString")
				dialogNodeModel.Output = dialogNodeOutputModel
				dialogNodeModel.Context = dialogNodeContextModel
				dialogNodeModel.Metadata = make(map[string]interface{})
				dialogNodeModel.NextStep = dialogNodeNextStepModel
				dialogNodeModel.Title = core.StringPtr("testString")
				dialogNodeModel.Type = core.StringPtr("standard")
				dialogNodeModel.EventName = core.StringPtr("focus")
				dialogNodeModel.Variable = core.StringPtr("testString")
				dialogNodeModel.Actions = []assistantv1.DialogNodeAction{*dialogNodeActionModel}
				dialogNodeModel.DigressIn = core.StringPtr("not_available")
				dialogNodeModel.DigressOut = core.StringPtr("allow_returning")
				dialogNodeModel.DigressOutSlots = core.StringPtr("not_allowed")
				dialogNodeModel.UserLabel = core.StringPtr("testString")
				dialogNodeModel.DisambiguationOptOut = core.BoolPtr(false)

				// Construct an instance of the Counterexample model
				counterexampleModel := new(assistantv1.Counterexample)
				counterexampleModel.Text = core.StringPtr("testString")

				// Construct an instance of the WorkspaceSystemSettingsTooling model
				workspaceSystemSettingsToolingModel := new(assistantv1.WorkspaceSystemSettingsTooling)
				workspaceSystemSettingsToolingModel.StoreGenericResponses = core.BoolPtr(true)

				// Construct an instance of the WorkspaceSystemSettingsDisambiguation model
				workspaceSystemSettingsDisambiguationModel := new(assistantv1.WorkspaceSystemSettingsDisambiguation)
				workspaceSystemSettingsDisambiguationModel.Prompt = core.StringPtr("testString")
				workspaceSystemSettingsDisambiguationModel.NoneOfTheAbovePrompt = core.StringPtr("testString")
				workspaceSystemSettingsDisambiguationModel.Enabled = core.BoolPtr(false)
				workspaceSystemSettingsDisambiguationModel.Sensitivity = core.StringPtr("auto")
				workspaceSystemSettingsDisambiguationModel.Randomize = core.BoolPtr(true)
				workspaceSystemSettingsDisambiguationModel.MaxSuggestions = core.Int64Ptr(int64(1))
				workspaceSystemSettingsDisambiguationModel.SuggestionTextPolicy = core.StringPtr("testString")

				// Construct an instance of the WorkspaceSystemSettingsSystemEntities model
				workspaceSystemSettingsSystemEntitiesModel := new(assistantv1.WorkspaceSystemSettingsSystemEntities)
				workspaceSystemSettingsSystemEntitiesModel.Enabled = core.BoolPtr(false)

				// Construct an instance of the WorkspaceSystemSettingsOffTopic model
				workspaceSystemSettingsOffTopicModel := new(assistantv1.WorkspaceSystemSettingsOffTopic)
				workspaceSystemSettingsOffTopicModel.Enabled = core.BoolPtr(false)

				// Construct an instance of the WorkspaceSystemSettings model
				workspaceSystemSettingsModel := new(assistantv1.WorkspaceSystemSettings)
				workspaceSystemSettingsModel.Tooling = workspaceSystemSettingsToolingModel
				workspaceSystemSettingsModel.Disambiguation = workspaceSystemSettingsDisambiguationModel
				workspaceSystemSettingsModel.HumanAgentAssist = make(map[string]interface{})
				workspaceSystemSettingsModel.SpellingSuggestions = core.BoolPtr(false)
				workspaceSystemSettingsModel.SpellingAutoCorrect = core.BoolPtr(false)
				workspaceSystemSettingsModel.SystemEntities = workspaceSystemSettingsSystemEntitiesModel
				workspaceSystemSettingsModel.OffTopic = workspaceSystemSettingsOffTopicModel
				workspaceSystemSettingsModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the WebhookHeader model
				webhookHeaderModel := new(assistantv1.WebhookHeader)
				webhookHeaderModel.Name = core.StringPtr("testString")
				webhookHeaderModel.Value = core.StringPtr("testString")

				// Construct an instance of the Webhook model
				webhookModel := new(assistantv1.Webhook)
				webhookModel.URL = core.StringPtr("testString")
				webhookModel.Name = core.StringPtr("testString")
				webhookModel.HeadersVar = []assistantv1.WebhookHeader{*webhookHeaderModel}

				// Construct an instance of the Mention model
				mentionModel := new(assistantv1.Mention)
				mentionModel.Entity = core.StringPtr("testString")
				mentionModel.Location = []int64{int64(38)}

				// Construct an instance of the Example model
				exampleModel := new(assistantv1.Example)
				exampleModel.Text = core.StringPtr("testString")
				exampleModel.Mentions = []assistantv1.Mention{*mentionModel}

				// Construct an instance of the CreateIntent model
				createIntentModel := new(assistantv1.CreateIntent)
				createIntentModel.Intent = core.StringPtr("testString")
				createIntentModel.Description = core.StringPtr("testString")
				createIntentModel.Examples = []assistantv1.Example{*exampleModel}

				// Construct an instance of the CreateValue model
				createValueModel := new(assistantv1.CreateValue)
				createValueModel.Value = core.StringPtr("testString")
				createValueModel.Metadata = make(map[string]interface{})
				createValueModel.Type = core.StringPtr("synonyms")
				createValueModel.Synonyms = []string{"testString"}
				createValueModel.Patterns = []string{"testString"}

				// Construct an instance of the CreateEntity model
				createEntityModel := new(assistantv1.CreateEntity)
				createEntityModel.Entity = core.StringPtr("testString")
				createEntityModel.Description = core.StringPtr("testString")
				createEntityModel.Metadata = make(map[string]interface{})
				createEntityModel.FuzzyMatch = core.BoolPtr(true)
				createEntityModel.Values = []assistantv1.CreateValue{*createValueModel}

				// Construct an instance of the UpdateWorkspaceOptions model
				updateWorkspaceOptionsModel := new(assistantv1.UpdateWorkspaceOptions)
				updateWorkspaceOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateWorkspaceOptionsModel.Name = core.StringPtr("testString")
				updateWorkspaceOptionsModel.Description = core.StringPtr("testString")
				updateWorkspaceOptionsModel.Language = core.StringPtr("testString")
				updateWorkspaceOptionsModel.DialogNodes = []assistantv1.DialogNode{*dialogNodeModel}
				updateWorkspaceOptionsModel.Counterexamples = []assistantv1.Counterexample{*counterexampleModel}
				updateWorkspaceOptionsModel.Metadata = make(map[string]interface{})
				updateWorkspaceOptionsModel.LearningOptOut = core.BoolPtr(false)
				updateWorkspaceOptionsModel.SystemSettings = workspaceSystemSettingsModel
				updateWorkspaceOptionsModel.Webhooks = []assistantv1.Webhook{*webhookModel}
				updateWorkspaceOptionsModel.Intents = []assistantv1.CreateIntent{*createIntentModel}
				updateWorkspaceOptionsModel.Entities = []assistantv1.CreateEntity{*createEntityModel}
				updateWorkspaceOptionsModel.Append = core.BoolPtr(false)
				updateWorkspaceOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateWorkspaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.UpdateWorkspace(updateWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateWorkspaceOptions model with no property values
				updateWorkspaceOptionsModelNew := new(assistantv1.UpdateWorkspaceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = assistantService.UpdateWorkspace(updateWorkspaceOptionsModelNew)
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
			It(`Invoke UpdateWorkspace successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ResponseGenericChannel model
				responseGenericChannelModel := new(assistantv1.ResponseGenericChannel)
				responseGenericChannelModel.Channel = core.StringPtr("chat")

				// Construct an instance of the DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo model
				dialogNodeOutputGenericModel := new(assistantv1.DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo)
				dialogNodeOutputGenericModel.ResponseType = core.StringPtr("video")
				dialogNodeOutputGenericModel.Source = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Title = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Description = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Channels = []assistantv1.ResponseGenericChannel{*responseGenericChannelModel}
				dialogNodeOutputGenericModel.ChannelOptions = map[string]interface{}{"anyKey": "anyValue"}
				dialogNodeOutputGenericModel.AltText = core.StringPtr("testString")

				// Construct an instance of the DialogNodeOutputModifiers model
				dialogNodeOutputModifiersModel := new(assistantv1.DialogNodeOutputModifiers)
				dialogNodeOutputModifiersModel.Overwrite = core.BoolPtr(true)

				// Construct an instance of the DialogNodeOutput model
				dialogNodeOutputModel := new(assistantv1.DialogNodeOutput)
				dialogNodeOutputModel.Generic = []assistantv1.DialogNodeOutputGenericIntf{dialogNodeOutputGenericModel}
				dialogNodeOutputModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeOutputModel.Modifiers = dialogNodeOutputModifiersModel
				dialogNodeOutputModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeContext model
				dialogNodeContextModel := new(assistantv1.DialogNodeContext)
				dialogNodeContextModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeContextModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeNextStep model
				dialogNodeNextStepModel := new(assistantv1.DialogNodeNextStep)
				dialogNodeNextStepModel.Behavior = core.StringPtr("get_user_input")
				dialogNodeNextStepModel.DialogNode = core.StringPtr("testString")
				dialogNodeNextStepModel.Selector = core.StringPtr("condition")

				// Construct an instance of the DialogNodeAction model
				dialogNodeActionModel := new(assistantv1.DialogNodeAction)
				dialogNodeActionModel.Name = core.StringPtr("testString")
				dialogNodeActionModel.Type = core.StringPtr("client")
				dialogNodeActionModel.Parameters = make(map[string]interface{})
				dialogNodeActionModel.ResultVariable = core.StringPtr("testString")
				dialogNodeActionModel.Credentials = core.StringPtr("testString")

				// Construct an instance of the DialogNode model
				dialogNodeModel := new(assistantv1.DialogNode)
				dialogNodeModel.DialogNode = core.StringPtr("testString")
				dialogNodeModel.Description = core.StringPtr("testString")
				dialogNodeModel.Conditions = core.StringPtr("testString")
				dialogNodeModel.Parent = core.StringPtr("testString")
				dialogNodeModel.PreviousSibling = core.StringPtr("testString")
				dialogNodeModel.Output = dialogNodeOutputModel
				dialogNodeModel.Context = dialogNodeContextModel
				dialogNodeModel.Metadata = make(map[string]interface{})
				dialogNodeModel.NextStep = dialogNodeNextStepModel
				dialogNodeModel.Title = core.StringPtr("testString")
				dialogNodeModel.Type = core.StringPtr("standard")
				dialogNodeModel.EventName = core.StringPtr("focus")
				dialogNodeModel.Variable = core.StringPtr("testString")
				dialogNodeModel.Actions = []assistantv1.DialogNodeAction{*dialogNodeActionModel}
				dialogNodeModel.DigressIn = core.StringPtr("not_available")
				dialogNodeModel.DigressOut = core.StringPtr("allow_returning")
				dialogNodeModel.DigressOutSlots = core.StringPtr("not_allowed")
				dialogNodeModel.UserLabel = core.StringPtr("testString")
				dialogNodeModel.DisambiguationOptOut = core.BoolPtr(false)

				// Construct an instance of the Counterexample model
				counterexampleModel := new(assistantv1.Counterexample)
				counterexampleModel.Text = core.StringPtr("testString")

				// Construct an instance of the WorkspaceSystemSettingsTooling model
				workspaceSystemSettingsToolingModel := new(assistantv1.WorkspaceSystemSettingsTooling)
				workspaceSystemSettingsToolingModel.StoreGenericResponses = core.BoolPtr(true)

				// Construct an instance of the WorkspaceSystemSettingsDisambiguation model
				workspaceSystemSettingsDisambiguationModel := new(assistantv1.WorkspaceSystemSettingsDisambiguation)
				workspaceSystemSettingsDisambiguationModel.Prompt = core.StringPtr("testString")
				workspaceSystemSettingsDisambiguationModel.NoneOfTheAbovePrompt = core.StringPtr("testString")
				workspaceSystemSettingsDisambiguationModel.Enabled = core.BoolPtr(false)
				workspaceSystemSettingsDisambiguationModel.Sensitivity = core.StringPtr("auto")
				workspaceSystemSettingsDisambiguationModel.Randomize = core.BoolPtr(true)
				workspaceSystemSettingsDisambiguationModel.MaxSuggestions = core.Int64Ptr(int64(1))
				workspaceSystemSettingsDisambiguationModel.SuggestionTextPolicy = core.StringPtr("testString")

				// Construct an instance of the WorkspaceSystemSettingsSystemEntities model
				workspaceSystemSettingsSystemEntitiesModel := new(assistantv1.WorkspaceSystemSettingsSystemEntities)
				workspaceSystemSettingsSystemEntitiesModel.Enabled = core.BoolPtr(false)

				// Construct an instance of the WorkspaceSystemSettingsOffTopic model
				workspaceSystemSettingsOffTopicModel := new(assistantv1.WorkspaceSystemSettingsOffTopic)
				workspaceSystemSettingsOffTopicModel.Enabled = core.BoolPtr(false)

				// Construct an instance of the WorkspaceSystemSettings model
				workspaceSystemSettingsModel := new(assistantv1.WorkspaceSystemSettings)
				workspaceSystemSettingsModel.Tooling = workspaceSystemSettingsToolingModel
				workspaceSystemSettingsModel.Disambiguation = workspaceSystemSettingsDisambiguationModel
				workspaceSystemSettingsModel.HumanAgentAssist = make(map[string]interface{})
				workspaceSystemSettingsModel.SpellingSuggestions = core.BoolPtr(false)
				workspaceSystemSettingsModel.SpellingAutoCorrect = core.BoolPtr(false)
				workspaceSystemSettingsModel.SystemEntities = workspaceSystemSettingsSystemEntitiesModel
				workspaceSystemSettingsModel.OffTopic = workspaceSystemSettingsOffTopicModel
				workspaceSystemSettingsModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the WebhookHeader model
				webhookHeaderModel := new(assistantv1.WebhookHeader)
				webhookHeaderModel.Name = core.StringPtr("testString")
				webhookHeaderModel.Value = core.StringPtr("testString")

				// Construct an instance of the Webhook model
				webhookModel := new(assistantv1.Webhook)
				webhookModel.URL = core.StringPtr("testString")
				webhookModel.Name = core.StringPtr("testString")
				webhookModel.HeadersVar = []assistantv1.WebhookHeader{*webhookHeaderModel}

				// Construct an instance of the Mention model
				mentionModel := new(assistantv1.Mention)
				mentionModel.Entity = core.StringPtr("testString")
				mentionModel.Location = []int64{int64(38)}

				// Construct an instance of the Example model
				exampleModel := new(assistantv1.Example)
				exampleModel.Text = core.StringPtr("testString")
				exampleModel.Mentions = []assistantv1.Mention{*mentionModel}

				// Construct an instance of the CreateIntent model
				createIntentModel := new(assistantv1.CreateIntent)
				createIntentModel.Intent = core.StringPtr("testString")
				createIntentModel.Description = core.StringPtr("testString")
				createIntentModel.Examples = []assistantv1.Example{*exampleModel}

				// Construct an instance of the CreateValue model
				createValueModel := new(assistantv1.CreateValue)
				createValueModel.Value = core.StringPtr("testString")
				createValueModel.Metadata = make(map[string]interface{})
				createValueModel.Type = core.StringPtr("synonyms")
				createValueModel.Synonyms = []string{"testString"}
				createValueModel.Patterns = []string{"testString"}

				// Construct an instance of the CreateEntity model
				createEntityModel := new(assistantv1.CreateEntity)
				createEntityModel.Entity = core.StringPtr("testString")
				createEntityModel.Description = core.StringPtr("testString")
				createEntityModel.Metadata = make(map[string]interface{})
				createEntityModel.FuzzyMatch = core.BoolPtr(true)
				createEntityModel.Values = []assistantv1.CreateValue{*createValueModel}

				// Construct an instance of the UpdateWorkspaceOptions model
				updateWorkspaceOptionsModel := new(assistantv1.UpdateWorkspaceOptions)
				updateWorkspaceOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateWorkspaceOptionsModel.Name = core.StringPtr("testString")
				updateWorkspaceOptionsModel.Description = core.StringPtr("testString")
				updateWorkspaceOptionsModel.Language = core.StringPtr("testString")
				updateWorkspaceOptionsModel.DialogNodes = []assistantv1.DialogNode{*dialogNodeModel}
				updateWorkspaceOptionsModel.Counterexamples = []assistantv1.Counterexample{*counterexampleModel}
				updateWorkspaceOptionsModel.Metadata = make(map[string]interface{})
				updateWorkspaceOptionsModel.LearningOptOut = core.BoolPtr(false)
				updateWorkspaceOptionsModel.SystemSettings = workspaceSystemSettingsModel
				updateWorkspaceOptionsModel.Webhooks = []assistantv1.Webhook{*webhookModel}
				updateWorkspaceOptionsModel.Intents = []assistantv1.CreateIntent{*createIntentModel}
				updateWorkspaceOptionsModel.Entities = []assistantv1.CreateEntity{*createEntityModel}
				updateWorkspaceOptionsModel.Append = core.BoolPtr(false)
				updateWorkspaceOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateWorkspaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.UpdateWorkspace(updateWorkspaceOptionsModel)
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
	Describe(`DeleteWorkspace(deleteWorkspaceOptions *DeleteWorkspaceOptions)`, func() {
		version := "testString"
		deleteWorkspacePath := "/v1/workspaces/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteWorkspacePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteWorkspace successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := assistantService.DeleteWorkspace(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteWorkspaceOptions model
				deleteWorkspaceOptionsModel := new(assistantv1.DeleteWorkspaceOptions)
				deleteWorkspaceOptionsModel.WorkspaceID = core.StringPtr("testString")
				deleteWorkspaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = assistantService.DeleteWorkspace(deleteWorkspaceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteWorkspace with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the DeleteWorkspaceOptions model
				deleteWorkspaceOptionsModel := new(assistantv1.DeleteWorkspaceOptions)
				deleteWorkspaceOptionsModel.WorkspaceID = core.StringPtr("testString")
				deleteWorkspaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := assistantService.DeleteWorkspace(deleteWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteWorkspaceOptions model with no property values
				deleteWorkspaceOptionsModelNew := new(assistantv1.DeleteWorkspaceOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = assistantService.DeleteWorkspace(deleteWorkspaceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListIntents(listIntentsOptions *ListIntentsOptions) - Operation response error`, func() {
		version := "testString"
		listIntentsPath := "/v1/workspaces/testString/intents"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listIntentsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for export query parameter
					Expect(req.URL.Query()["page_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// TODO: Add check for include_count query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"intent"}))
					Expect(req.URL.Query()["cursor"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListIntents with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ListIntentsOptions model
				listIntentsOptionsModel := new(assistantv1.ListIntentsOptions)
				listIntentsOptionsModel.WorkspaceID = core.StringPtr("testString")
				listIntentsOptionsModel.Export = core.BoolPtr(false)
				listIntentsOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listIntentsOptionsModel.IncludeCount = core.BoolPtr(false)
				listIntentsOptionsModel.Sort = core.StringPtr("intent")
				listIntentsOptionsModel.Cursor = core.StringPtr("testString")
				listIntentsOptionsModel.IncludeAudit = core.BoolPtr(false)
				listIntentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.ListIntents(listIntentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.ListIntents(listIntentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListIntents(listIntentsOptions *ListIntentsOptions)`, func() {
		version := "testString"
		listIntentsPath := "/v1/workspaces/testString/intents"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listIntentsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for export query parameter
					Expect(req.URL.Query()["page_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// TODO: Add check for include_count query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"intent"}))
					Expect(req.URL.Query()["cursor"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"intents": [{"intent": "Intent", "description": "Description", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "examples": [{"text": "Text", "mentions": [{"entity": "Entity", "location": [8]}], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}], "pagination": {"refresh_url": "RefreshURL", "next_url": "NextURL", "total": 5, "matched": 7, "refresh_cursor": "RefreshCursor", "next_cursor": "NextCursor"}}`)
				}))
			})
			It(`Invoke ListIntents successfully with retries`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the ListIntentsOptions model
				listIntentsOptionsModel := new(assistantv1.ListIntentsOptions)
				listIntentsOptionsModel.WorkspaceID = core.StringPtr("testString")
				listIntentsOptionsModel.Export = core.BoolPtr(false)
				listIntentsOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listIntentsOptionsModel.IncludeCount = core.BoolPtr(false)
				listIntentsOptionsModel.Sort = core.StringPtr("intent")
				listIntentsOptionsModel.Cursor = core.StringPtr("testString")
				listIntentsOptionsModel.IncludeAudit = core.BoolPtr(false)
				listIntentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.ListIntentsWithContext(ctx, listIntentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.ListIntents(listIntentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.ListIntentsWithContext(ctx, listIntentsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listIntentsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for export query parameter
					Expect(req.URL.Query()["page_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// TODO: Add check for include_count query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"intent"}))
					Expect(req.URL.Query()["cursor"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"intents": [{"intent": "Intent", "description": "Description", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "examples": [{"text": "Text", "mentions": [{"entity": "Entity", "location": [8]}], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}], "pagination": {"refresh_url": "RefreshURL", "next_url": "NextURL", "total": 5, "matched": 7, "refresh_cursor": "RefreshCursor", "next_cursor": "NextCursor"}}`)
				}))
			})
			It(`Invoke ListIntents successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.ListIntents(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListIntentsOptions model
				listIntentsOptionsModel := new(assistantv1.ListIntentsOptions)
				listIntentsOptionsModel.WorkspaceID = core.StringPtr("testString")
				listIntentsOptionsModel.Export = core.BoolPtr(false)
				listIntentsOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listIntentsOptionsModel.IncludeCount = core.BoolPtr(false)
				listIntentsOptionsModel.Sort = core.StringPtr("intent")
				listIntentsOptionsModel.Cursor = core.StringPtr("testString")
				listIntentsOptionsModel.IncludeAudit = core.BoolPtr(false)
				listIntentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.ListIntents(listIntentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListIntents with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ListIntentsOptions model
				listIntentsOptionsModel := new(assistantv1.ListIntentsOptions)
				listIntentsOptionsModel.WorkspaceID = core.StringPtr("testString")
				listIntentsOptionsModel.Export = core.BoolPtr(false)
				listIntentsOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listIntentsOptionsModel.IncludeCount = core.BoolPtr(false)
				listIntentsOptionsModel.Sort = core.StringPtr("intent")
				listIntentsOptionsModel.Cursor = core.StringPtr("testString")
				listIntentsOptionsModel.IncludeAudit = core.BoolPtr(false)
				listIntentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.ListIntents(listIntentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListIntentsOptions model with no property values
				listIntentsOptionsModelNew := new(assistantv1.ListIntentsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = assistantService.ListIntents(listIntentsOptionsModelNew)
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
			It(`Invoke ListIntents successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ListIntentsOptions model
				listIntentsOptionsModel := new(assistantv1.ListIntentsOptions)
				listIntentsOptionsModel.WorkspaceID = core.StringPtr("testString")
				listIntentsOptionsModel.Export = core.BoolPtr(false)
				listIntentsOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listIntentsOptionsModel.IncludeCount = core.BoolPtr(false)
				listIntentsOptionsModel.Sort = core.StringPtr("intent")
				listIntentsOptionsModel.Cursor = core.StringPtr("testString")
				listIntentsOptionsModel.IncludeAudit = core.BoolPtr(false)
				listIntentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.ListIntents(listIntentsOptionsModel)
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
	Describe(`CreateIntent(createIntentOptions *CreateIntentOptions) - Operation response error`, func() {
		version := "testString"
		createIntentPath := "/v1/workspaces/testString/intents"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createIntentPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateIntent with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the Mention model
				mentionModel := new(assistantv1.Mention)
				mentionModel.Entity = core.StringPtr("testString")
				mentionModel.Location = []int64{int64(38)}

				// Construct an instance of the Example model
				exampleModel := new(assistantv1.Example)
				exampleModel.Text = core.StringPtr("testString")
				exampleModel.Mentions = []assistantv1.Mention{*mentionModel}

				// Construct an instance of the CreateIntentOptions model
				createIntentOptionsModel := new(assistantv1.CreateIntentOptions)
				createIntentOptionsModel.WorkspaceID = core.StringPtr("testString")
				createIntentOptionsModel.Intent = core.StringPtr("testString")
				createIntentOptionsModel.Description = core.StringPtr("testString")
				createIntentOptionsModel.Examples = []assistantv1.Example{*exampleModel}
				createIntentOptionsModel.IncludeAudit = core.BoolPtr(false)
				createIntentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.CreateIntent(createIntentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.CreateIntent(createIntentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateIntent(createIntentOptions *CreateIntentOptions)`, func() {
		version := "testString"
		createIntentPath := "/v1/workspaces/testString/intents"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createIntentPath))
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
					// TODO: Add check for include_audit query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"intent": "Intent", "description": "Description", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "examples": [{"text": "Text", "mentions": [{"entity": "Entity", "location": [8]}], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke CreateIntent successfully with retries`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the Mention model
				mentionModel := new(assistantv1.Mention)
				mentionModel.Entity = core.StringPtr("testString")
				mentionModel.Location = []int64{int64(38)}

				// Construct an instance of the Example model
				exampleModel := new(assistantv1.Example)
				exampleModel.Text = core.StringPtr("testString")
				exampleModel.Mentions = []assistantv1.Mention{*mentionModel}

				// Construct an instance of the CreateIntentOptions model
				createIntentOptionsModel := new(assistantv1.CreateIntentOptions)
				createIntentOptionsModel.WorkspaceID = core.StringPtr("testString")
				createIntentOptionsModel.Intent = core.StringPtr("testString")
				createIntentOptionsModel.Description = core.StringPtr("testString")
				createIntentOptionsModel.Examples = []assistantv1.Example{*exampleModel}
				createIntentOptionsModel.IncludeAudit = core.BoolPtr(false)
				createIntentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.CreateIntentWithContext(ctx, createIntentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.CreateIntent(createIntentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.CreateIntentWithContext(ctx, createIntentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createIntentPath))
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
					// TODO: Add check for include_audit query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"intent": "Intent", "description": "Description", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "examples": [{"text": "Text", "mentions": [{"entity": "Entity", "location": [8]}], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke CreateIntent successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.CreateIntent(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Mention model
				mentionModel := new(assistantv1.Mention)
				mentionModel.Entity = core.StringPtr("testString")
				mentionModel.Location = []int64{int64(38)}

				// Construct an instance of the Example model
				exampleModel := new(assistantv1.Example)
				exampleModel.Text = core.StringPtr("testString")
				exampleModel.Mentions = []assistantv1.Mention{*mentionModel}

				// Construct an instance of the CreateIntentOptions model
				createIntentOptionsModel := new(assistantv1.CreateIntentOptions)
				createIntentOptionsModel.WorkspaceID = core.StringPtr("testString")
				createIntentOptionsModel.Intent = core.StringPtr("testString")
				createIntentOptionsModel.Description = core.StringPtr("testString")
				createIntentOptionsModel.Examples = []assistantv1.Example{*exampleModel}
				createIntentOptionsModel.IncludeAudit = core.BoolPtr(false)
				createIntentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.CreateIntent(createIntentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateIntent with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the Mention model
				mentionModel := new(assistantv1.Mention)
				mentionModel.Entity = core.StringPtr("testString")
				mentionModel.Location = []int64{int64(38)}

				// Construct an instance of the Example model
				exampleModel := new(assistantv1.Example)
				exampleModel.Text = core.StringPtr("testString")
				exampleModel.Mentions = []assistantv1.Mention{*mentionModel}

				// Construct an instance of the CreateIntentOptions model
				createIntentOptionsModel := new(assistantv1.CreateIntentOptions)
				createIntentOptionsModel.WorkspaceID = core.StringPtr("testString")
				createIntentOptionsModel.Intent = core.StringPtr("testString")
				createIntentOptionsModel.Description = core.StringPtr("testString")
				createIntentOptionsModel.Examples = []assistantv1.Example{*exampleModel}
				createIntentOptionsModel.IncludeAudit = core.BoolPtr(false)
				createIntentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.CreateIntent(createIntentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateIntentOptions model with no property values
				createIntentOptionsModelNew := new(assistantv1.CreateIntentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = assistantService.CreateIntent(createIntentOptionsModelNew)
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
			It(`Invoke CreateIntent successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the Mention model
				mentionModel := new(assistantv1.Mention)
				mentionModel.Entity = core.StringPtr("testString")
				mentionModel.Location = []int64{int64(38)}

				// Construct an instance of the Example model
				exampleModel := new(assistantv1.Example)
				exampleModel.Text = core.StringPtr("testString")
				exampleModel.Mentions = []assistantv1.Mention{*mentionModel}

				// Construct an instance of the CreateIntentOptions model
				createIntentOptionsModel := new(assistantv1.CreateIntentOptions)
				createIntentOptionsModel.WorkspaceID = core.StringPtr("testString")
				createIntentOptionsModel.Intent = core.StringPtr("testString")
				createIntentOptionsModel.Description = core.StringPtr("testString")
				createIntentOptionsModel.Examples = []assistantv1.Example{*exampleModel}
				createIntentOptionsModel.IncludeAudit = core.BoolPtr(false)
				createIntentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.CreateIntent(createIntentOptionsModel)
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
	Describe(`GetIntent(getIntentOptions *GetIntentOptions) - Operation response error`, func() {
		version := "testString"
		getIntentPath := "/v1/workspaces/testString/intents/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getIntentPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for export query parameter
					// TODO: Add check for include_audit query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetIntent with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the GetIntentOptions model
				getIntentOptionsModel := new(assistantv1.GetIntentOptions)
				getIntentOptionsModel.WorkspaceID = core.StringPtr("testString")
				getIntentOptionsModel.Intent = core.StringPtr("testString")
				getIntentOptionsModel.Export = core.BoolPtr(false)
				getIntentOptionsModel.IncludeAudit = core.BoolPtr(false)
				getIntentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.GetIntent(getIntentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.GetIntent(getIntentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetIntent(getIntentOptions *GetIntentOptions)`, func() {
		version := "testString"
		getIntentPath := "/v1/workspaces/testString/intents/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getIntentPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for export query parameter
					// TODO: Add check for include_audit query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"intent": "Intent", "description": "Description", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "examples": [{"text": "Text", "mentions": [{"entity": "Entity", "location": [8]}], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke GetIntent successfully with retries`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the GetIntentOptions model
				getIntentOptionsModel := new(assistantv1.GetIntentOptions)
				getIntentOptionsModel.WorkspaceID = core.StringPtr("testString")
				getIntentOptionsModel.Intent = core.StringPtr("testString")
				getIntentOptionsModel.Export = core.BoolPtr(false)
				getIntentOptionsModel.IncludeAudit = core.BoolPtr(false)
				getIntentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.GetIntentWithContext(ctx, getIntentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.GetIntent(getIntentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.GetIntentWithContext(ctx, getIntentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getIntentPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for export query parameter
					// TODO: Add check for include_audit query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"intent": "Intent", "description": "Description", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "examples": [{"text": "Text", "mentions": [{"entity": "Entity", "location": [8]}], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke GetIntent successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.GetIntent(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetIntentOptions model
				getIntentOptionsModel := new(assistantv1.GetIntentOptions)
				getIntentOptionsModel.WorkspaceID = core.StringPtr("testString")
				getIntentOptionsModel.Intent = core.StringPtr("testString")
				getIntentOptionsModel.Export = core.BoolPtr(false)
				getIntentOptionsModel.IncludeAudit = core.BoolPtr(false)
				getIntentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.GetIntent(getIntentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetIntent with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the GetIntentOptions model
				getIntentOptionsModel := new(assistantv1.GetIntentOptions)
				getIntentOptionsModel.WorkspaceID = core.StringPtr("testString")
				getIntentOptionsModel.Intent = core.StringPtr("testString")
				getIntentOptionsModel.Export = core.BoolPtr(false)
				getIntentOptionsModel.IncludeAudit = core.BoolPtr(false)
				getIntentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.GetIntent(getIntentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetIntentOptions model with no property values
				getIntentOptionsModelNew := new(assistantv1.GetIntentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = assistantService.GetIntent(getIntentOptionsModelNew)
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
			It(`Invoke GetIntent successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the GetIntentOptions model
				getIntentOptionsModel := new(assistantv1.GetIntentOptions)
				getIntentOptionsModel.WorkspaceID = core.StringPtr("testString")
				getIntentOptionsModel.Intent = core.StringPtr("testString")
				getIntentOptionsModel.Export = core.BoolPtr(false)
				getIntentOptionsModel.IncludeAudit = core.BoolPtr(false)
				getIntentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.GetIntent(getIntentOptionsModel)
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
	Describe(`UpdateIntent(updateIntentOptions *UpdateIntentOptions) - Operation response error`, func() {
		version := "testString"
		updateIntentPath := "/v1/workspaces/testString/intents/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateIntentPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for append query parameter
					// TODO: Add check for include_audit query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateIntent with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the Mention model
				mentionModel := new(assistantv1.Mention)
				mentionModel.Entity = core.StringPtr("testString")
				mentionModel.Location = []int64{int64(38)}

				// Construct an instance of the Example model
				exampleModel := new(assistantv1.Example)
				exampleModel.Text = core.StringPtr("testString")
				exampleModel.Mentions = []assistantv1.Mention{*mentionModel}

				// Construct an instance of the UpdateIntentOptions model
				updateIntentOptionsModel := new(assistantv1.UpdateIntentOptions)
				updateIntentOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateIntentOptionsModel.Intent = core.StringPtr("testString")
				updateIntentOptionsModel.NewIntent = core.StringPtr("testString")
				updateIntentOptionsModel.NewDescription = core.StringPtr("testString")
				updateIntentOptionsModel.NewExamples = []assistantv1.Example{*exampleModel}
				updateIntentOptionsModel.Append = core.BoolPtr(false)
				updateIntentOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateIntentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.UpdateIntent(updateIntentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.UpdateIntent(updateIntentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateIntent(updateIntentOptions *UpdateIntentOptions)`, func() {
		version := "testString"
		updateIntentPath := "/v1/workspaces/testString/intents/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateIntentPath))
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
					// TODO: Add check for append query parameter
					// TODO: Add check for include_audit query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"intent": "Intent", "description": "Description", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "examples": [{"text": "Text", "mentions": [{"entity": "Entity", "location": [8]}], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke UpdateIntent successfully with retries`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the Mention model
				mentionModel := new(assistantv1.Mention)
				mentionModel.Entity = core.StringPtr("testString")
				mentionModel.Location = []int64{int64(38)}

				// Construct an instance of the Example model
				exampleModel := new(assistantv1.Example)
				exampleModel.Text = core.StringPtr("testString")
				exampleModel.Mentions = []assistantv1.Mention{*mentionModel}

				// Construct an instance of the UpdateIntentOptions model
				updateIntentOptionsModel := new(assistantv1.UpdateIntentOptions)
				updateIntentOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateIntentOptionsModel.Intent = core.StringPtr("testString")
				updateIntentOptionsModel.NewIntent = core.StringPtr("testString")
				updateIntentOptionsModel.NewDescription = core.StringPtr("testString")
				updateIntentOptionsModel.NewExamples = []assistantv1.Example{*exampleModel}
				updateIntentOptionsModel.Append = core.BoolPtr(false)
				updateIntentOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateIntentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.UpdateIntentWithContext(ctx, updateIntentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.UpdateIntent(updateIntentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.UpdateIntentWithContext(ctx, updateIntentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateIntentPath))
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
					// TODO: Add check for append query parameter
					// TODO: Add check for include_audit query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"intent": "Intent", "description": "Description", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "examples": [{"text": "Text", "mentions": [{"entity": "Entity", "location": [8]}], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke UpdateIntent successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.UpdateIntent(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Mention model
				mentionModel := new(assistantv1.Mention)
				mentionModel.Entity = core.StringPtr("testString")
				mentionModel.Location = []int64{int64(38)}

				// Construct an instance of the Example model
				exampleModel := new(assistantv1.Example)
				exampleModel.Text = core.StringPtr("testString")
				exampleModel.Mentions = []assistantv1.Mention{*mentionModel}

				// Construct an instance of the UpdateIntentOptions model
				updateIntentOptionsModel := new(assistantv1.UpdateIntentOptions)
				updateIntentOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateIntentOptionsModel.Intent = core.StringPtr("testString")
				updateIntentOptionsModel.NewIntent = core.StringPtr("testString")
				updateIntentOptionsModel.NewDescription = core.StringPtr("testString")
				updateIntentOptionsModel.NewExamples = []assistantv1.Example{*exampleModel}
				updateIntentOptionsModel.Append = core.BoolPtr(false)
				updateIntentOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateIntentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.UpdateIntent(updateIntentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateIntent with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the Mention model
				mentionModel := new(assistantv1.Mention)
				mentionModel.Entity = core.StringPtr("testString")
				mentionModel.Location = []int64{int64(38)}

				// Construct an instance of the Example model
				exampleModel := new(assistantv1.Example)
				exampleModel.Text = core.StringPtr("testString")
				exampleModel.Mentions = []assistantv1.Mention{*mentionModel}

				// Construct an instance of the UpdateIntentOptions model
				updateIntentOptionsModel := new(assistantv1.UpdateIntentOptions)
				updateIntentOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateIntentOptionsModel.Intent = core.StringPtr("testString")
				updateIntentOptionsModel.NewIntent = core.StringPtr("testString")
				updateIntentOptionsModel.NewDescription = core.StringPtr("testString")
				updateIntentOptionsModel.NewExamples = []assistantv1.Example{*exampleModel}
				updateIntentOptionsModel.Append = core.BoolPtr(false)
				updateIntentOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateIntentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.UpdateIntent(updateIntentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateIntentOptions model with no property values
				updateIntentOptionsModelNew := new(assistantv1.UpdateIntentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = assistantService.UpdateIntent(updateIntentOptionsModelNew)
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
			It(`Invoke UpdateIntent successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the Mention model
				mentionModel := new(assistantv1.Mention)
				mentionModel.Entity = core.StringPtr("testString")
				mentionModel.Location = []int64{int64(38)}

				// Construct an instance of the Example model
				exampleModel := new(assistantv1.Example)
				exampleModel.Text = core.StringPtr("testString")
				exampleModel.Mentions = []assistantv1.Mention{*mentionModel}

				// Construct an instance of the UpdateIntentOptions model
				updateIntentOptionsModel := new(assistantv1.UpdateIntentOptions)
				updateIntentOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateIntentOptionsModel.Intent = core.StringPtr("testString")
				updateIntentOptionsModel.NewIntent = core.StringPtr("testString")
				updateIntentOptionsModel.NewDescription = core.StringPtr("testString")
				updateIntentOptionsModel.NewExamples = []assistantv1.Example{*exampleModel}
				updateIntentOptionsModel.Append = core.BoolPtr(false)
				updateIntentOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateIntentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.UpdateIntent(updateIntentOptionsModel)
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
	Describe(`DeleteIntent(deleteIntentOptions *DeleteIntentOptions)`, func() {
		version := "testString"
		deleteIntentPath := "/v1/workspaces/testString/intents/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteIntentPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteIntent successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := assistantService.DeleteIntent(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteIntentOptions model
				deleteIntentOptionsModel := new(assistantv1.DeleteIntentOptions)
				deleteIntentOptionsModel.WorkspaceID = core.StringPtr("testString")
				deleteIntentOptionsModel.Intent = core.StringPtr("testString")
				deleteIntentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = assistantService.DeleteIntent(deleteIntentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteIntent with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the DeleteIntentOptions model
				deleteIntentOptionsModel := new(assistantv1.DeleteIntentOptions)
				deleteIntentOptionsModel.WorkspaceID = core.StringPtr("testString")
				deleteIntentOptionsModel.Intent = core.StringPtr("testString")
				deleteIntentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := assistantService.DeleteIntent(deleteIntentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteIntentOptions model with no property values
				deleteIntentOptionsModelNew := new(assistantv1.DeleteIntentOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = assistantService.DeleteIntent(deleteIntentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListExamples(listExamplesOptions *ListExamplesOptions) - Operation response error`, func() {
		version := "testString"
		listExamplesPath := "/v1/workspaces/testString/intents/testString/examples"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listExamplesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["page_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// TODO: Add check for include_count query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"text"}))
					Expect(req.URL.Query()["cursor"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListExamples with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ListExamplesOptions model
				listExamplesOptionsModel := new(assistantv1.ListExamplesOptions)
				listExamplesOptionsModel.WorkspaceID = core.StringPtr("testString")
				listExamplesOptionsModel.Intent = core.StringPtr("testString")
				listExamplesOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listExamplesOptionsModel.IncludeCount = core.BoolPtr(false)
				listExamplesOptionsModel.Sort = core.StringPtr("text")
				listExamplesOptionsModel.Cursor = core.StringPtr("testString")
				listExamplesOptionsModel.IncludeAudit = core.BoolPtr(false)
				listExamplesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.ListExamples(listExamplesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.ListExamples(listExamplesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListExamples(listExamplesOptions *ListExamplesOptions)`, func() {
		version := "testString"
		listExamplesPath := "/v1/workspaces/testString/intents/testString/examples"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listExamplesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["page_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// TODO: Add check for include_count query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"text"}))
					Expect(req.URL.Query()["cursor"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"examples": [{"text": "Text", "mentions": [{"entity": "Entity", "location": [8]}], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}], "pagination": {"refresh_url": "RefreshURL", "next_url": "NextURL", "total": 5, "matched": 7, "refresh_cursor": "RefreshCursor", "next_cursor": "NextCursor"}}`)
				}))
			})
			It(`Invoke ListExamples successfully with retries`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the ListExamplesOptions model
				listExamplesOptionsModel := new(assistantv1.ListExamplesOptions)
				listExamplesOptionsModel.WorkspaceID = core.StringPtr("testString")
				listExamplesOptionsModel.Intent = core.StringPtr("testString")
				listExamplesOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listExamplesOptionsModel.IncludeCount = core.BoolPtr(false)
				listExamplesOptionsModel.Sort = core.StringPtr("text")
				listExamplesOptionsModel.Cursor = core.StringPtr("testString")
				listExamplesOptionsModel.IncludeAudit = core.BoolPtr(false)
				listExamplesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.ListExamplesWithContext(ctx, listExamplesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.ListExamples(listExamplesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.ListExamplesWithContext(ctx, listExamplesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listExamplesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["page_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// TODO: Add check for include_count query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"text"}))
					Expect(req.URL.Query()["cursor"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"examples": [{"text": "Text", "mentions": [{"entity": "Entity", "location": [8]}], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}], "pagination": {"refresh_url": "RefreshURL", "next_url": "NextURL", "total": 5, "matched": 7, "refresh_cursor": "RefreshCursor", "next_cursor": "NextCursor"}}`)
				}))
			})
			It(`Invoke ListExamples successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.ListExamples(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListExamplesOptions model
				listExamplesOptionsModel := new(assistantv1.ListExamplesOptions)
				listExamplesOptionsModel.WorkspaceID = core.StringPtr("testString")
				listExamplesOptionsModel.Intent = core.StringPtr("testString")
				listExamplesOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listExamplesOptionsModel.IncludeCount = core.BoolPtr(false)
				listExamplesOptionsModel.Sort = core.StringPtr("text")
				listExamplesOptionsModel.Cursor = core.StringPtr("testString")
				listExamplesOptionsModel.IncludeAudit = core.BoolPtr(false)
				listExamplesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.ListExamples(listExamplesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListExamples with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ListExamplesOptions model
				listExamplesOptionsModel := new(assistantv1.ListExamplesOptions)
				listExamplesOptionsModel.WorkspaceID = core.StringPtr("testString")
				listExamplesOptionsModel.Intent = core.StringPtr("testString")
				listExamplesOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listExamplesOptionsModel.IncludeCount = core.BoolPtr(false)
				listExamplesOptionsModel.Sort = core.StringPtr("text")
				listExamplesOptionsModel.Cursor = core.StringPtr("testString")
				listExamplesOptionsModel.IncludeAudit = core.BoolPtr(false)
				listExamplesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.ListExamples(listExamplesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListExamplesOptions model with no property values
				listExamplesOptionsModelNew := new(assistantv1.ListExamplesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = assistantService.ListExamples(listExamplesOptionsModelNew)
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
			It(`Invoke ListExamples successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ListExamplesOptions model
				listExamplesOptionsModel := new(assistantv1.ListExamplesOptions)
				listExamplesOptionsModel.WorkspaceID = core.StringPtr("testString")
				listExamplesOptionsModel.Intent = core.StringPtr("testString")
				listExamplesOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listExamplesOptionsModel.IncludeCount = core.BoolPtr(false)
				listExamplesOptionsModel.Sort = core.StringPtr("text")
				listExamplesOptionsModel.Cursor = core.StringPtr("testString")
				listExamplesOptionsModel.IncludeAudit = core.BoolPtr(false)
				listExamplesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.ListExamples(listExamplesOptionsModel)
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
	Describe(`CreateExample(createExampleOptions *CreateExampleOptions) - Operation response error`, func() {
		version := "testString"
		createExamplePath := "/v1/workspaces/testString/intents/testString/examples"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createExamplePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateExample with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the Mention model
				mentionModel := new(assistantv1.Mention)
				mentionModel.Entity = core.StringPtr("testString")
				mentionModel.Location = []int64{int64(38)}

				// Construct an instance of the CreateExampleOptions model
				createExampleOptionsModel := new(assistantv1.CreateExampleOptions)
				createExampleOptionsModel.WorkspaceID = core.StringPtr("testString")
				createExampleOptionsModel.Intent = core.StringPtr("testString")
				createExampleOptionsModel.Text = core.StringPtr("testString")
				createExampleOptionsModel.Mentions = []assistantv1.Mention{*mentionModel}
				createExampleOptionsModel.IncludeAudit = core.BoolPtr(false)
				createExampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.CreateExample(createExampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.CreateExample(createExampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateExample(createExampleOptions *CreateExampleOptions)`, func() {
		version := "testString"
		createExamplePath := "/v1/workspaces/testString/intents/testString/examples"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createExamplePath))
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
					// TODO: Add check for include_audit query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"text": "Text", "mentions": [{"entity": "Entity", "location": [8]}], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke CreateExample successfully with retries`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the Mention model
				mentionModel := new(assistantv1.Mention)
				mentionModel.Entity = core.StringPtr("testString")
				mentionModel.Location = []int64{int64(38)}

				// Construct an instance of the CreateExampleOptions model
				createExampleOptionsModel := new(assistantv1.CreateExampleOptions)
				createExampleOptionsModel.WorkspaceID = core.StringPtr("testString")
				createExampleOptionsModel.Intent = core.StringPtr("testString")
				createExampleOptionsModel.Text = core.StringPtr("testString")
				createExampleOptionsModel.Mentions = []assistantv1.Mention{*mentionModel}
				createExampleOptionsModel.IncludeAudit = core.BoolPtr(false)
				createExampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.CreateExampleWithContext(ctx, createExampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.CreateExample(createExampleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.CreateExampleWithContext(ctx, createExampleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createExamplePath))
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
					// TODO: Add check for include_audit query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"text": "Text", "mentions": [{"entity": "Entity", "location": [8]}], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke CreateExample successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.CreateExample(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Mention model
				mentionModel := new(assistantv1.Mention)
				mentionModel.Entity = core.StringPtr("testString")
				mentionModel.Location = []int64{int64(38)}

				// Construct an instance of the CreateExampleOptions model
				createExampleOptionsModel := new(assistantv1.CreateExampleOptions)
				createExampleOptionsModel.WorkspaceID = core.StringPtr("testString")
				createExampleOptionsModel.Intent = core.StringPtr("testString")
				createExampleOptionsModel.Text = core.StringPtr("testString")
				createExampleOptionsModel.Mentions = []assistantv1.Mention{*mentionModel}
				createExampleOptionsModel.IncludeAudit = core.BoolPtr(false)
				createExampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.CreateExample(createExampleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateExample with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the Mention model
				mentionModel := new(assistantv1.Mention)
				mentionModel.Entity = core.StringPtr("testString")
				mentionModel.Location = []int64{int64(38)}

				// Construct an instance of the CreateExampleOptions model
				createExampleOptionsModel := new(assistantv1.CreateExampleOptions)
				createExampleOptionsModel.WorkspaceID = core.StringPtr("testString")
				createExampleOptionsModel.Intent = core.StringPtr("testString")
				createExampleOptionsModel.Text = core.StringPtr("testString")
				createExampleOptionsModel.Mentions = []assistantv1.Mention{*mentionModel}
				createExampleOptionsModel.IncludeAudit = core.BoolPtr(false)
				createExampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.CreateExample(createExampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateExampleOptions model with no property values
				createExampleOptionsModelNew := new(assistantv1.CreateExampleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = assistantService.CreateExample(createExampleOptionsModelNew)
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
			It(`Invoke CreateExample successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the Mention model
				mentionModel := new(assistantv1.Mention)
				mentionModel.Entity = core.StringPtr("testString")
				mentionModel.Location = []int64{int64(38)}

				// Construct an instance of the CreateExampleOptions model
				createExampleOptionsModel := new(assistantv1.CreateExampleOptions)
				createExampleOptionsModel.WorkspaceID = core.StringPtr("testString")
				createExampleOptionsModel.Intent = core.StringPtr("testString")
				createExampleOptionsModel.Text = core.StringPtr("testString")
				createExampleOptionsModel.Mentions = []assistantv1.Mention{*mentionModel}
				createExampleOptionsModel.IncludeAudit = core.BoolPtr(false)
				createExampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.CreateExample(createExampleOptionsModel)
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
	Describe(`GetExample(getExampleOptions *GetExampleOptions) - Operation response error`, func() {
		version := "testString"
		getExamplePath := "/v1/workspaces/testString/intents/testString/examples/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getExamplePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetExample with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the GetExampleOptions model
				getExampleOptionsModel := new(assistantv1.GetExampleOptions)
				getExampleOptionsModel.WorkspaceID = core.StringPtr("testString")
				getExampleOptionsModel.Intent = core.StringPtr("testString")
				getExampleOptionsModel.Text = core.StringPtr("testString")
				getExampleOptionsModel.IncludeAudit = core.BoolPtr(false)
				getExampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.GetExample(getExampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.GetExample(getExampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetExample(getExampleOptions *GetExampleOptions)`, func() {
		version := "testString"
		getExamplePath := "/v1/workspaces/testString/intents/testString/examples/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getExamplePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"text": "Text", "mentions": [{"entity": "Entity", "location": [8]}], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetExample successfully with retries`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the GetExampleOptions model
				getExampleOptionsModel := new(assistantv1.GetExampleOptions)
				getExampleOptionsModel.WorkspaceID = core.StringPtr("testString")
				getExampleOptionsModel.Intent = core.StringPtr("testString")
				getExampleOptionsModel.Text = core.StringPtr("testString")
				getExampleOptionsModel.IncludeAudit = core.BoolPtr(false)
				getExampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.GetExampleWithContext(ctx, getExampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.GetExample(getExampleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.GetExampleWithContext(ctx, getExampleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getExamplePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"text": "Text", "mentions": [{"entity": "Entity", "location": [8]}], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetExample successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.GetExample(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetExampleOptions model
				getExampleOptionsModel := new(assistantv1.GetExampleOptions)
				getExampleOptionsModel.WorkspaceID = core.StringPtr("testString")
				getExampleOptionsModel.Intent = core.StringPtr("testString")
				getExampleOptionsModel.Text = core.StringPtr("testString")
				getExampleOptionsModel.IncludeAudit = core.BoolPtr(false)
				getExampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.GetExample(getExampleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetExample with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the GetExampleOptions model
				getExampleOptionsModel := new(assistantv1.GetExampleOptions)
				getExampleOptionsModel.WorkspaceID = core.StringPtr("testString")
				getExampleOptionsModel.Intent = core.StringPtr("testString")
				getExampleOptionsModel.Text = core.StringPtr("testString")
				getExampleOptionsModel.IncludeAudit = core.BoolPtr(false)
				getExampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.GetExample(getExampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetExampleOptions model with no property values
				getExampleOptionsModelNew := new(assistantv1.GetExampleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = assistantService.GetExample(getExampleOptionsModelNew)
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
			It(`Invoke GetExample successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the GetExampleOptions model
				getExampleOptionsModel := new(assistantv1.GetExampleOptions)
				getExampleOptionsModel.WorkspaceID = core.StringPtr("testString")
				getExampleOptionsModel.Intent = core.StringPtr("testString")
				getExampleOptionsModel.Text = core.StringPtr("testString")
				getExampleOptionsModel.IncludeAudit = core.BoolPtr(false)
				getExampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.GetExample(getExampleOptionsModel)
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
	Describe(`UpdateExample(updateExampleOptions *UpdateExampleOptions) - Operation response error`, func() {
		version := "testString"
		updateExamplePath := "/v1/workspaces/testString/intents/testString/examples/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateExamplePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateExample with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the Mention model
				mentionModel := new(assistantv1.Mention)
				mentionModel.Entity = core.StringPtr("testString")
				mentionModel.Location = []int64{int64(38)}

				// Construct an instance of the UpdateExampleOptions model
				updateExampleOptionsModel := new(assistantv1.UpdateExampleOptions)
				updateExampleOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateExampleOptionsModel.Intent = core.StringPtr("testString")
				updateExampleOptionsModel.Text = core.StringPtr("testString")
				updateExampleOptionsModel.NewText = core.StringPtr("testString")
				updateExampleOptionsModel.NewMentions = []assistantv1.Mention{*mentionModel}
				updateExampleOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateExampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.UpdateExample(updateExampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.UpdateExample(updateExampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateExample(updateExampleOptions *UpdateExampleOptions)`, func() {
		version := "testString"
		updateExamplePath := "/v1/workspaces/testString/intents/testString/examples/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateExamplePath))
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
					// TODO: Add check for include_audit query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"text": "Text", "mentions": [{"entity": "Entity", "location": [8]}], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke UpdateExample successfully with retries`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the Mention model
				mentionModel := new(assistantv1.Mention)
				mentionModel.Entity = core.StringPtr("testString")
				mentionModel.Location = []int64{int64(38)}

				// Construct an instance of the UpdateExampleOptions model
				updateExampleOptionsModel := new(assistantv1.UpdateExampleOptions)
				updateExampleOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateExampleOptionsModel.Intent = core.StringPtr("testString")
				updateExampleOptionsModel.Text = core.StringPtr("testString")
				updateExampleOptionsModel.NewText = core.StringPtr("testString")
				updateExampleOptionsModel.NewMentions = []assistantv1.Mention{*mentionModel}
				updateExampleOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateExampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.UpdateExampleWithContext(ctx, updateExampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.UpdateExample(updateExampleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.UpdateExampleWithContext(ctx, updateExampleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateExamplePath))
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
					// TODO: Add check for include_audit query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"text": "Text", "mentions": [{"entity": "Entity", "location": [8]}], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke UpdateExample successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.UpdateExample(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Mention model
				mentionModel := new(assistantv1.Mention)
				mentionModel.Entity = core.StringPtr("testString")
				mentionModel.Location = []int64{int64(38)}

				// Construct an instance of the UpdateExampleOptions model
				updateExampleOptionsModel := new(assistantv1.UpdateExampleOptions)
				updateExampleOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateExampleOptionsModel.Intent = core.StringPtr("testString")
				updateExampleOptionsModel.Text = core.StringPtr("testString")
				updateExampleOptionsModel.NewText = core.StringPtr("testString")
				updateExampleOptionsModel.NewMentions = []assistantv1.Mention{*mentionModel}
				updateExampleOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateExampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.UpdateExample(updateExampleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateExample with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the Mention model
				mentionModel := new(assistantv1.Mention)
				mentionModel.Entity = core.StringPtr("testString")
				mentionModel.Location = []int64{int64(38)}

				// Construct an instance of the UpdateExampleOptions model
				updateExampleOptionsModel := new(assistantv1.UpdateExampleOptions)
				updateExampleOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateExampleOptionsModel.Intent = core.StringPtr("testString")
				updateExampleOptionsModel.Text = core.StringPtr("testString")
				updateExampleOptionsModel.NewText = core.StringPtr("testString")
				updateExampleOptionsModel.NewMentions = []assistantv1.Mention{*mentionModel}
				updateExampleOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateExampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.UpdateExample(updateExampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateExampleOptions model with no property values
				updateExampleOptionsModelNew := new(assistantv1.UpdateExampleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = assistantService.UpdateExample(updateExampleOptionsModelNew)
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
			It(`Invoke UpdateExample successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the Mention model
				mentionModel := new(assistantv1.Mention)
				mentionModel.Entity = core.StringPtr("testString")
				mentionModel.Location = []int64{int64(38)}

				// Construct an instance of the UpdateExampleOptions model
				updateExampleOptionsModel := new(assistantv1.UpdateExampleOptions)
				updateExampleOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateExampleOptionsModel.Intent = core.StringPtr("testString")
				updateExampleOptionsModel.Text = core.StringPtr("testString")
				updateExampleOptionsModel.NewText = core.StringPtr("testString")
				updateExampleOptionsModel.NewMentions = []assistantv1.Mention{*mentionModel}
				updateExampleOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateExampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.UpdateExample(updateExampleOptionsModel)
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
	Describe(`DeleteExample(deleteExampleOptions *DeleteExampleOptions)`, func() {
		version := "testString"
		deleteExamplePath := "/v1/workspaces/testString/intents/testString/examples/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteExamplePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteExample successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := assistantService.DeleteExample(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteExampleOptions model
				deleteExampleOptionsModel := new(assistantv1.DeleteExampleOptions)
				deleteExampleOptionsModel.WorkspaceID = core.StringPtr("testString")
				deleteExampleOptionsModel.Intent = core.StringPtr("testString")
				deleteExampleOptionsModel.Text = core.StringPtr("testString")
				deleteExampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = assistantService.DeleteExample(deleteExampleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteExample with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the DeleteExampleOptions model
				deleteExampleOptionsModel := new(assistantv1.DeleteExampleOptions)
				deleteExampleOptionsModel.WorkspaceID = core.StringPtr("testString")
				deleteExampleOptionsModel.Intent = core.StringPtr("testString")
				deleteExampleOptionsModel.Text = core.StringPtr("testString")
				deleteExampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := assistantService.DeleteExample(deleteExampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteExampleOptions model with no property values
				deleteExampleOptionsModelNew := new(assistantv1.DeleteExampleOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = assistantService.DeleteExample(deleteExampleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListCounterexamples(listCounterexamplesOptions *ListCounterexamplesOptions) - Operation response error`, func() {
		version := "testString"
		listCounterexamplesPath := "/v1/workspaces/testString/counterexamples"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCounterexamplesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["page_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// TODO: Add check for include_count query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"text"}))
					Expect(req.URL.Query()["cursor"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListCounterexamples with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ListCounterexamplesOptions model
				listCounterexamplesOptionsModel := new(assistantv1.ListCounterexamplesOptions)
				listCounterexamplesOptionsModel.WorkspaceID = core.StringPtr("testString")
				listCounterexamplesOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listCounterexamplesOptionsModel.IncludeCount = core.BoolPtr(false)
				listCounterexamplesOptionsModel.Sort = core.StringPtr("text")
				listCounterexamplesOptionsModel.Cursor = core.StringPtr("testString")
				listCounterexamplesOptionsModel.IncludeAudit = core.BoolPtr(false)
				listCounterexamplesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.ListCounterexamples(listCounterexamplesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.ListCounterexamples(listCounterexamplesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListCounterexamples(listCounterexamplesOptions *ListCounterexamplesOptions)`, func() {
		version := "testString"
		listCounterexamplesPath := "/v1/workspaces/testString/counterexamples"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCounterexamplesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["page_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// TODO: Add check for include_count query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"text"}))
					Expect(req.URL.Query()["cursor"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"counterexamples": [{"text": "Text", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}], "pagination": {"refresh_url": "RefreshURL", "next_url": "NextURL", "total": 5, "matched": 7, "refresh_cursor": "RefreshCursor", "next_cursor": "NextCursor"}}`)
				}))
			})
			It(`Invoke ListCounterexamples successfully with retries`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the ListCounterexamplesOptions model
				listCounterexamplesOptionsModel := new(assistantv1.ListCounterexamplesOptions)
				listCounterexamplesOptionsModel.WorkspaceID = core.StringPtr("testString")
				listCounterexamplesOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listCounterexamplesOptionsModel.IncludeCount = core.BoolPtr(false)
				listCounterexamplesOptionsModel.Sort = core.StringPtr("text")
				listCounterexamplesOptionsModel.Cursor = core.StringPtr("testString")
				listCounterexamplesOptionsModel.IncludeAudit = core.BoolPtr(false)
				listCounterexamplesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.ListCounterexamplesWithContext(ctx, listCounterexamplesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.ListCounterexamples(listCounterexamplesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.ListCounterexamplesWithContext(ctx, listCounterexamplesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listCounterexamplesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["page_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// TODO: Add check for include_count query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"text"}))
					Expect(req.URL.Query()["cursor"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"counterexamples": [{"text": "Text", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}], "pagination": {"refresh_url": "RefreshURL", "next_url": "NextURL", "total": 5, "matched": 7, "refresh_cursor": "RefreshCursor", "next_cursor": "NextCursor"}}`)
				}))
			})
			It(`Invoke ListCounterexamples successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.ListCounterexamples(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListCounterexamplesOptions model
				listCounterexamplesOptionsModel := new(assistantv1.ListCounterexamplesOptions)
				listCounterexamplesOptionsModel.WorkspaceID = core.StringPtr("testString")
				listCounterexamplesOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listCounterexamplesOptionsModel.IncludeCount = core.BoolPtr(false)
				listCounterexamplesOptionsModel.Sort = core.StringPtr("text")
				listCounterexamplesOptionsModel.Cursor = core.StringPtr("testString")
				listCounterexamplesOptionsModel.IncludeAudit = core.BoolPtr(false)
				listCounterexamplesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.ListCounterexamples(listCounterexamplesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListCounterexamples with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ListCounterexamplesOptions model
				listCounterexamplesOptionsModel := new(assistantv1.ListCounterexamplesOptions)
				listCounterexamplesOptionsModel.WorkspaceID = core.StringPtr("testString")
				listCounterexamplesOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listCounterexamplesOptionsModel.IncludeCount = core.BoolPtr(false)
				listCounterexamplesOptionsModel.Sort = core.StringPtr("text")
				listCounterexamplesOptionsModel.Cursor = core.StringPtr("testString")
				listCounterexamplesOptionsModel.IncludeAudit = core.BoolPtr(false)
				listCounterexamplesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.ListCounterexamples(listCounterexamplesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListCounterexamplesOptions model with no property values
				listCounterexamplesOptionsModelNew := new(assistantv1.ListCounterexamplesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = assistantService.ListCounterexamples(listCounterexamplesOptionsModelNew)
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
			It(`Invoke ListCounterexamples successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ListCounterexamplesOptions model
				listCounterexamplesOptionsModel := new(assistantv1.ListCounterexamplesOptions)
				listCounterexamplesOptionsModel.WorkspaceID = core.StringPtr("testString")
				listCounterexamplesOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listCounterexamplesOptionsModel.IncludeCount = core.BoolPtr(false)
				listCounterexamplesOptionsModel.Sort = core.StringPtr("text")
				listCounterexamplesOptionsModel.Cursor = core.StringPtr("testString")
				listCounterexamplesOptionsModel.IncludeAudit = core.BoolPtr(false)
				listCounterexamplesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.ListCounterexamples(listCounterexamplesOptionsModel)
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
	Describe(`CreateCounterexample(createCounterexampleOptions *CreateCounterexampleOptions) - Operation response error`, func() {
		version := "testString"
		createCounterexamplePath := "/v1/workspaces/testString/counterexamples"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCounterexamplePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateCounterexample with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the CreateCounterexampleOptions model
				createCounterexampleOptionsModel := new(assistantv1.CreateCounterexampleOptions)
				createCounterexampleOptionsModel.WorkspaceID = core.StringPtr("testString")
				createCounterexampleOptionsModel.Text = core.StringPtr("testString")
				createCounterexampleOptionsModel.IncludeAudit = core.BoolPtr(false)
				createCounterexampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.CreateCounterexample(createCounterexampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.CreateCounterexample(createCounterexampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateCounterexample(createCounterexampleOptions *CreateCounterexampleOptions)`, func() {
		version := "testString"
		createCounterexamplePath := "/v1/workspaces/testString/counterexamples"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCounterexamplePath))
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
					// TODO: Add check for include_audit query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"text": "Text", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke CreateCounterexample successfully with retries`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the CreateCounterexampleOptions model
				createCounterexampleOptionsModel := new(assistantv1.CreateCounterexampleOptions)
				createCounterexampleOptionsModel.WorkspaceID = core.StringPtr("testString")
				createCounterexampleOptionsModel.Text = core.StringPtr("testString")
				createCounterexampleOptionsModel.IncludeAudit = core.BoolPtr(false)
				createCounterexampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.CreateCounterexampleWithContext(ctx, createCounterexampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.CreateCounterexample(createCounterexampleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.CreateCounterexampleWithContext(ctx, createCounterexampleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createCounterexamplePath))
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
					// TODO: Add check for include_audit query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"text": "Text", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke CreateCounterexample successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.CreateCounterexample(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateCounterexampleOptions model
				createCounterexampleOptionsModel := new(assistantv1.CreateCounterexampleOptions)
				createCounterexampleOptionsModel.WorkspaceID = core.StringPtr("testString")
				createCounterexampleOptionsModel.Text = core.StringPtr("testString")
				createCounterexampleOptionsModel.IncludeAudit = core.BoolPtr(false)
				createCounterexampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.CreateCounterexample(createCounterexampleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateCounterexample with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the CreateCounterexampleOptions model
				createCounterexampleOptionsModel := new(assistantv1.CreateCounterexampleOptions)
				createCounterexampleOptionsModel.WorkspaceID = core.StringPtr("testString")
				createCounterexampleOptionsModel.Text = core.StringPtr("testString")
				createCounterexampleOptionsModel.IncludeAudit = core.BoolPtr(false)
				createCounterexampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.CreateCounterexample(createCounterexampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateCounterexampleOptions model with no property values
				createCounterexampleOptionsModelNew := new(assistantv1.CreateCounterexampleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = assistantService.CreateCounterexample(createCounterexampleOptionsModelNew)
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
			It(`Invoke CreateCounterexample successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the CreateCounterexampleOptions model
				createCounterexampleOptionsModel := new(assistantv1.CreateCounterexampleOptions)
				createCounterexampleOptionsModel.WorkspaceID = core.StringPtr("testString")
				createCounterexampleOptionsModel.Text = core.StringPtr("testString")
				createCounterexampleOptionsModel.IncludeAudit = core.BoolPtr(false)
				createCounterexampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.CreateCounterexample(createCounterexampleOptionsModel)
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
	Describe(`GetCounterexample(getCounterexampleOptions *GetCounterexampleOptions) - Operation response error`, func() {
		version := "testString"
		getCounterexamplePath := "/v1/workspaces/testString/counterexamples/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCounterexamplePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCounterexample with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the GetCounterexampleOptions model
				getCounterexampleOptionsModel := new(assistantv1.GetCounterexampleOptions)
				getCounterexampleOptionsModel.WorkspaceID = core.StringPtr("testString")
				getCounterexampleOptionsModel.Text = core.StringPtr("testString")
				getCounterexampleOptionsModel.IncludeAudit = core.BoolPtr(false)
				getCounterexampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.GetCounterexample(getCounterexampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.GetCounterexample(getCounterexampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCounterexample(getCounterexampleOptions *GetCounterexampleOptions)`, func() {
		version := "testString"
		getCounterexamplePath := "/v1/workspaces/testString/counterexamples/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCounterexamplePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"text": "Text", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetCounterexample successfully with retries`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the GetCounterexampleOptions model
				getCounterexampleOptionsModel := new(assistantv1.GetCounterexampleOptions)
				getCounterexampleOptionsModel.WorkspaceID = core.StringPtr("testString")
				getCounterexampleOptionsModel.Text = core.StringPtr("testString")
				getCounterexampleOptionsModel.IncludeAudit = core.BoolPtr(false)
				getCounterexampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.GetCounterexampleWithContext(ctx, getCounterexampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.GetCounterexample(getCounterexampleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.GetCounterexampleWithContext(ctx, getCounterexampleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getCounterexamplePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"text": "Text", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetCounterexample successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.GetCounterexample(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCounterexampleOptions model
				getCounterexampleOptionsModel := new(assistantv1.GetCounterexampleOptions)
				getCounterexampleOptionsModel.WorkspaceID = core.StringPtr("testString")
				getCounterexampleOptionsModel.Text = core.StringPtr("testString")
				getCounterexampleOptionsModel.IncludeAudit = core.BoolPtr(false)
				getCounterexampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.GetCounterexample(getCounterexampleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCounterexample with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the GetCounterexampleOptions model
				getCounterexampleOptionsModel := new(assistantv1.GetCounterexampleOptions)
				getCounterexampleOptionsModel.WorkspaceID = core.StringPtr("testString")
				getCounterexampleOptionsModel.Text = core.StringPtr("testString")
				getCounterexampleOptionsModel.IncludeAudit = core.BoolPtr(false)
				getCounterexampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.GetCounterexample(getCounterexampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetCounterexampleOptions model with no property values
				getCounterexampleOptionsModelNew := new(assistantv1.GetCounterexampleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = assistantService.GetCounterexample(getCounterexampleOptionsModelNew)
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
			It(`Invoke GetCounterexample successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the GetCounterexampleOptions model
				getCounterexampleOptionsModel := new(assistantv1.GetCounterexampleOptions)
				getCounterexampleOptionsModel.WorkspaceID = core.StringPtr("testString")
				getCounterexampleOptionsModel.Text = core.StringPtr("testString")
				getCounterexampleOptionsModel.IncludeAudit = core.BoolPtr(false)
				getCounterexampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.GetCounterexample(getCounterexampleOptionsModel)
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
	Describe(`UpdateCounterexample(updateCounterexampleOptions *UpdateCounterexampleOptions) - Operation response error`, func() {
		version := "testString"
		updateCounterexamplePath := "/v1/workspaces/testString/counterexamples/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCounterexamplePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateCounterexample with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the UpdateCounterexampleOptions model
				updateCounterexampleOptionsModel := new(assistantv1.UpdateCounterexampleOptions)
				updateCounterexampleOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateCounterexampleOptionsModel.Text = core.StringPtr("testString")
				updateCounterexampleOptionsModel.NewText = core.StringPtr("testString")
				updateCounterexampleOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateCounterexampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.UpdateCounterexample(updateCounterexampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.UpdateCounterexample(updateCounterexampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateCounterexample(updateCounterexampleOptions *UpdateCounterexampleOptions)`, func() {
		version := "testString"
		updateCounterexamplePath := "/v1/workspaces/testString/counterexamples/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCounterexamplePath))
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
					// TODO: Add check for include_audit query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"text": "Text", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke UpdateCounterexample successfully with retries`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the UpdateCounterexampleOptions model
				updateCounterexampleOptionsModel := new(assistantv1.UpdateCounterexampleOptions)
				updateCounterexampleOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateCounterexampleOptionsModel.Text = core.StringPtr("testString")
				updateCounterexampleOptionsModel.NewText = core.StringPtr("testString")
				updateCounterexampleOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateCounterexampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.UpdateCounterexampleWithContext(ctx, updateCounterexampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.UpdateCounterexample(updateCounterexampleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.UpdateCounterexampleWithContext(ctx, updateCounterexampleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateCounterexamplePath))
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
					// TODO: Add check for include_audit query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"text": "Text", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke UpdateCounterexample successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.UpdateCounterexample(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateCounterexampleOptions model
				updateCounterexampleOptionsModel := new(assistantv1.UpdateCounterexampleOptions)
				updateCounterexampleOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateCounterexampleOptionsModel.Text = core.StringPtr("testString")
				updateCounterexampleOptionsModel.NewText = core.StringPtr("testString")
				updateCounterexampleOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateCounterexampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.UpdateCounterexample(updateCounterexampleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateCounterexample with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the UpdateCounterexampleOptions model
				updateCounterexampleOptionsModel := new(assistantv1.UpdateCounterexampleOptions)
				updateCounterexampleOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateCounterexampleOptionsModel.Text = core.StringPtr("testString")
				updateCounterexampleOptionsModel.NewText = core.StringPtr("testString")
				updateCounterexampleOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateCounterexampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.UpdateCounterexample(updateCounterexampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateCounterexampleOptions model with no property values
				updateCounterexampleOptionsModelNew := new(assistantv1.UpdateCounterexampleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = assistantService.UpdateCounterexample(updateCounterexampleOptionsModelNew)
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
			It(`Invoke UpdateCounterexample successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the UpdateCounterexampleOptions model
				updateCounterexampleOptionsModel := new(assistantv1.UpdateCounterexampleOptions)
				updateCounterexampleOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateCounterexampleOptionsModel.Text = core.StringPtr("testString")
				updateCounterexampleOptionsModel.NewText = core.StringPtr("testString")
				updateCounterexampleOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateCounterexampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.UpdateCounterexample(updateCounterexampleOptionsModel)
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
	Describe(`DeleteCounterexample(deleteCounterexampleOptions *DeleteCounterexampleOptions)`, func() {
		version := "testString"
		deleteCounterexamplePath := "/v1/workspaces/testString/counterexamples/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteCounterexamplePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteCounterexample successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := assistantService.DeleteCounterexample(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteCounterexampleOptions model
				deleteCounterexampleOptionsModel := new(assistantv1.DeleteCounterexampleOptions)
				deleteCounterexampleOptionsModel.WorkspaceID = core.StringPtr("testString")
				deleteCounterexampleOptionsModel.Text = core.StringPtr("testString")
				deleteCounterexampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = assistantService.DeleteCounterexample(deleteCounterexampleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteCounterexample with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the DeleteCounterexampleOptions model
				deleteCounterexampleOptionsModel := new(assistantv1.DeleteCounterexampleOptions)
				deleteCounterexampleOptionsModel.WorkspaceID = core.StringPtr("testString")
				deleteCounterexampleOptionsModel.Text = core.StringPtr("testString")
				deleteCounterexampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := assistantService.DeleteCounterexample(deleteCounterexampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteCounterexampleOptions model with no property values
				deleteCounterexampleOptionsModelNew := new(assistantv1.DeleteCounterexampleOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = assistantService.DeleteCounterexample(deleteCounterexampleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListEntities(listEntitiesOptions *ListEntitiesOptions) - Operation response error`, func() {
		version := "testString"
		listEntitiesPath := "/v1/workspaces/testString/entities"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listEntitiesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for export query parameter
					Expect(req.URL.Query()["page_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// TODO: Add check for include_count query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"entity"}))
					Expect(req.URL.Query()["cursor"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListEntities with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ListEntitiesOptions model
				listEntitiesOptionsModel := new(assistantv1.ListEntitiesOptions)
				listEntitiesOptionsModel.WorkspaceID = core.StringPtr("testString")
				listEntitiesOptionsModel.Export = core.BoolPtr(false)
				listEntitiesOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listEntitiesOptionsModel.IncludeCount = core.BoolPtr(false)
				listEntitiesOptionsModel.Sort = core.StringPtr("entity")
				listEntitiesOptionsModel.Cursor = core.StringPtr("testString")
				listEntitiesOptionsModel.IncludeAudit = core.BoolPtr(false)
				listEntitiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.ListEntities(listEntitiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.ListEntities(listEntitiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListEntities(listEntitiesOptions *ListEntitiesOptions)`, func() {
		version := "testString"
		listEntitiesPath := "/v1/workspaces/testString/entities"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listEntitiesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for export query parameter
					Expect(req.URL.Query()["page_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// TODO: Add check for include_count query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"entity"}))
					Expect(req.URL.Query()["cursor"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"entities": [{"entity": "Entity", "description": "Description", "metadata": {"mapKey": "anyValue"}, "fuzzy_match": true, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "values": [{"value": "Value", "metadata": {"mapKey": "anyValue"}, "type": "synonyms", "synonyms": ["Synonym"], "patterns": ["Pattern"], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}], "pagination": {"refresh_url": "RefreshURL", "next_url": "NextURL", "total": 5, "matched": 7, "refresh_cursor": "RefreshCursor", "next_cursor": "NextCursor"}}`)
				}))
			})
			It(`Invoke ListEntities successfully with retries`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the ListEntitiesOptions model
				listEntitiesOptionsModel := new(assistantv1.ListEntitiesOptions)
				listEntitiesOptionsModel.WorkspaceID = core.StringPtr("testString")
				listEntitiesOptionsModel.Export = core.BoolPtr(false)
				listEntitiesOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listEntitiesOptionsModel.IncludeCount = core.BoolPtr(false)
				listEntitiesOptionsModel.Sort = core.StringPtr("entity")
				listEntitiesOptionsModel.Cursor = core.StringPtr("testString")
				listEntitiesOptionsModel.IncludeAudit = core.BoolPtr(false)
				listEntitiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.ListEntitiesWithContext(ctx, listEntitiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.ListEntities(listEntitiesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.ListEntitiesWithContext(ctx, listEntitiesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listEntitiesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for export query parameter
					Expect(req.URL.Query()["page_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// TODO: Add check for include_count query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"entity"}))
					Expect(req.URL.Query()["cursor"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"entities": [{"entity": "Entity", "description": "Description", "metadata": {"mapKey": "anyValue"}, "fuzzy_match": true, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "values": [{"value": "Value", "metadata": {"mapKey": "anyValue"}, "type": "synonyms", "synonyms": ["Synonym"], "patterns": ["Pattern"], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}], "pagination": {"refresh_url": "RefreshURL", "next_url": "NextURL", "total": 5, "matched": 7, "refresh_cursor": "RefreshCursor", "next_cursor": "NextCursor"}}`)
				}))
			})
			It(`Invoke ListEntities successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.ListEntities(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListEntitiesOptions model
				listEntitiesOptionsModel := new(assistantv1.ListEntitiesOptions)
				listEntitiesOptionsModel.WorkspaceID = core.StringPtr("testString")
				listEntitiesOptionsModel.Export = core.BoolPtr(false)
				listEntitiesOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listEntitiesOptionsModel.IncludeCount = core.BoolPtr(false)
				listEntitiesOptionsModel.Sort = core.StringPtr("entity")
				listEntitiesOptionsModel.Cursor = core.StringPtr("testString")
				listEntitiesOptionsModel.IncludeAudit = core.BoolPtr(false)
				listEntitiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.ListEntities(listEntitiesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListEntities with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ListEntitiesOptions model
				listEntitiesOptionsModel := new(assistantv1.ListEntitiesOptions)
				listEntitiesOptionsModel.WorkspaceID = core.StringPtr("testString")
				listEntitiesOptionsModel.Export = core.BoolPtr(false)
				listEntitiesOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listEntitiesOptionsModel.IncludeCount = core.BoolPtr(false)
				listEntitiesOptionsModel.Sort = core.StringPtr("entity")
				listEntitiesOptionsModel.Cursor = core.StringPtr("testString")
				listEntitiesOptionsModel.IncludeAudit = core.BoolPtr(false)
				listEntitiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.ListEntities(listEntitiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListEntitiesOptions model with no property values
				listEntitiesOptionsModelNew := new(assistantv1.ListEntitiesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = assistantService.ListEntities(listEntitiesOptionsModelNew)
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
			It(`Invoke ListEntities successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ListEntitiesOptions model
				listEntitiesOptionsModel := new(assistantv1.ListEntitiesOptions)
				listEntitiesOptionsModel.WorkspaceID = core.StringPtr("testString")
				listEntitiesOptionsModel.Export = core.BoolPtr(false)
				listEntitiesOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listEntitiesOptionsModel.IncludeCount = core.BoolPtr(false)
				listEntitiesOptionsModel.Sort = core.StringPtr("entity")
				listEntitiesOptionsModel.Cursor = core.StringPtr("testString")
				listEntitiesOptionsModel.IncludeAudit = core.BoolPtr(false)
				listEntitiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.ListEntities(listEntitiesOptionsModel)
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
	Describe(`CreateEntity(createEntityOptions *CreateEntityOptions) - Operation response error`, func() {
		version := "testString"
		createEntityPath := "/v1/workspaces/testString/entities"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createEntityPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateEntity with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the CreateValue model
				createValueModel := new(assistantv1.CreateValue)
				createValueModel.Value = core.StringPtr("testString")
				createValueModel.Metadata = make(map[string]interface{})
				createValueModel.Type = core.StringPtr("synonyms")
				createValueModel.Synonyms = []string{"testString"}
				createValueModel.Patterns = []string{"testString"}

				// Construct an instance of the CreateEntityOptions model
				createEntityOptionsModel := new(assistantv1.CreateEntityOptions)
				createEntityOptionsModel.WorkspaceID = core.StringPtr("testString")
				createEntityOptionsModel.Entity = core.StringPtr("testString")
				createEntityOptionsModel.Description = core.StringPtr("testString")
				createEntityOptionsModel.Metadata = make(map[string]interface{})
				createEntityOptionsModel.FuzzyMatch = core.BoolPtr(true)
				createEntityOptionsModel.Values = []assistantv1.CreateValue{*createValueModel}
				createEntityOptionsModel.IncludeAudit = core.BoolPtr(false)
				createEntityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.CreateEntity(createEntityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.CreateEntity(createEntityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateEntity(createEntityOptions *CreateEntityOptions)`, func() {
		version := "testString"
		createEntityPath := "/v1/workspaces/testString/entities"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createEntityPath))
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
					// TODO: Add check for include_audit query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"entity": "Entity", "description": "Description", "metadata": {"mapKey": "anyValue"}, "fuzzy_match": true, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "values": [{"value": "Value", "metadata": {"mapKey": "anyValue"}, "type": "synonyms", "synonyms": ["Synonym"], "patterns": ["Pattern"], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke CreateEntity successfully with retries`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the CreateValue model
				createValueModel := new(assistantv1.CreateValue)
				createValueModel.Value = core.StringPtr("testString")
				createValueModel.Metadata = make(map[string]interface{})
				createValueModel.Type = core.StringPtr("synonyms")
				createValueModel.Synonyms = []string{"testString"}
				createValueModel.Patterns = []string{"testString"}

				// Construct an instance of the CreateEntityOptions model
				createEntityOptionsModel := new(assistantv1.CreateEntityOptions)
				createEntityOptionsModel.WorkspaceID = core.StringPtr("testString")
				createEntityOptionsModel.Entity = core.StringPtr("testString")
				createEntityOptionsModel.Description = core.StringPtr("testString")
				createEntityOptionsModel.Metadata = make(map[string]interface{})
				createEntityOptionsModel.FuzzyMatch = core.BoolPtr(true)
				createEntityOptionsModel.Values = []assistantv1.CreateValue{*createValueModel}
				createEntityOptionsModel.IncludeAudit = core.BoolPtr(false)
				createEntityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.CreateEntityWithContext(ctx, createEntityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.CreateEntity(createEntityOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.CreateEntityWithContext(ctx, createEntityOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createEntityPath))
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
					// TODO: Add check for include_audit query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"entity": "Entity", "description": "Description", "metadata": {"mapKey": "anyValue"}, "fuzzy_match": true, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "values": [{"value": "Value", "metadata": {"mapKey": "anyValue"}, "type": "synonyms", "synonyms": ["Synonym"], "patterns": ["Pattern"], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke CreateEntity successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.CreateEntity(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateValue model
				createValueModel := new(assistantv1.CreateValue)
				createValueModel.Value = core.StringPtr("testString")
				createValueModel.Metadata = make(map[string]interface{})
				createValueModel.Type = core.StringPtr("synonyms")
				createValueModel.Synonyms = []string{"testString"}
				createValueModel.Patterns = []string{"testString"}

				// Construct an instance of the CreateEntityOptions model
				createEntityOptionsModel := new(assistantv1.CreateEntityOptions)
				createEntityOptionsModel.WorkspaceID = core.StringPtr("testString")
				createEntityOptionsModel.Entity = core.StringPtr("testString")
				createEntityOptionsModel.Description = core.StringPtr("testString")
				createEntityOptionsModel.Metadata = make(map[string]interface{})
				createEntityOptionsModel.FuzzyMatch = core.BoolPtr(true)
				createEntityOptionsModel.Values = []assistantv1.CreateValue{*createValueModel}
				createEntityOptionsModel.IncludeAudit = core.BoolPtr(false)
				createEntityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.CreateEntity(createEntityOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateEntity with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the CreateValue model
				createValueModel := new(assistantv1.CreateValue)
				createValueModel.Value = core.StringPtr("testString")
				createValueModel.Metadata = make(map[string]interface{})
				createValueModel.Type = core.StringPtr("synonyms")
				createValueModel.Synonyms = []string{"testString"}
				createValueModel.Patterns = []string{"testString"}

				// Construct an instance of the CreateEntityOptions model
				createEntityOptionsModel := new(assistantv1.CreateEntityOptions)
				createEntityOptionsModel.WorkspaceID = core.StringPtr("testString")
				createEntityOptionsModel.Entity = core.StringPtr("testString")
				createEntityOptionsModel.Description = core.StringPtr("testString")
				createEntityOptionsModel.Metadata = make(map[string]interface{})
				createEntityOptionsModel.FuzzyMatch = core.BoolPtr(true)
				createEntityOptionsModel.Values = []assistantv1.CreateValue{*createValueModel}
				createEntityOptionsModel.IncludeAudit = core.BoolPtr(false)
				createEntityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.CreateEntity(createEntityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateEntityOptions model with no property values
				createEntityOptionsModelNew := new(assistantv1.CreateEntityOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = assistantService.CreateEntity(createEntityOptionsModelNew)
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
			It(`Invoke CreateEntity successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the CreateValue model
				createValueModel := new(assistantv1.CreateValue)
				createValueModel.Value = core.StringPtr("testString")
				createValueModel.Metadata = make(map[string]interface{})
				createValueModel.Type = core.StringPtr("synonyms")
				createValueModel.Synonyms = []string{"testString"}
				createValueModel.Patterns = []string{"testString"}

				// Construct an instance of the CreateEntityOptions model
				createEntityOptionsModel := new(assistantv1.CreateEntityOptions)
				createEntityOptionsModel.WorkspaceID = core.StringPtr("testString")
				createEntityOptionsModel.Entity = core.StringPtr("testString")
				createEntityOptionsModel.Description = core.StringPtr("testString")
				createEntityOptionsModel.Metadata = make(map[string]interface{})
				createEntityOptionsModel.FuzzyMatch = core.BoolPtr(true)
				createEntityOptionsModel.Values = []assistantv1.CreateValue{*createValueModel}
				createEntityOptionsModel.IncludeAudit = core.BoolPtr(false)
				createEntityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.CreateEntity(createEntityOptionsModel)
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
	Describe(`GetEntity(getEntityOptions *GetEntityOptions) - Operation response error`, func() {
		version := "testString"
		getEntityPath := "/v1/workspaces/testString/entities/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getEntityPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for export query parameter
					// TODO: Add check for include_audit query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetEntity with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the GetEntityOptions model
				getEntityOptionsModel := new(assistantv1.GetEntityOptions)
				getEntityOptionsModel.WorkspaceID = core.StringPtr("testString")
				getEntityOptionsModel.Entity = core.StringPtr("testString")
				getEntityOptionsModel.Export = core.BoolPtr(false)
				getEntityOptionsModel.IncludeAudit = core.BoolPtr(false)
				getEntityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.GetEntity(getEntityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.GetEntity(getEntityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetEntity(getEntityOptions *GetEntityOptions)`, func() {
		version := "testString"
		getEntityPath := "/v1/workspaces/testString/entities/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getEntityPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for export query parameter
					// TODO: Add check for include_audit query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"entity": "Entity", "description": "Description", "metadata": {"mapKey": "anyValue"}, "fuzzy_match": true, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "values": [{"value": "Value", "metadata": {"mapKey": "anyValue"}, "type": "synonyms", "synonyms": ["Synonym"], "patterns": ["Pattern"], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke GetEntity successfully with retries`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the GetEntityOptions model
				getEntityOptionsModel := new(assistantv1.GetEntityOptions)
				getEntityOptionsModel.WorkspaceID = core.StringPtr("testString")
				getEntityOptionsModel.Entity = core.StringPtr("testString")
				getEntityOptionsModel.Export = core.BoolPtr(false)
				getEntityOptionsModel.IncludeAudit = core.BoolPtr(false)
				getEntityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.GetEntityWithContext(ctx, getEntityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.GetEntity(getEntityOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.GetEntityWithContext(ctx, getEntityOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getEntityPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for export query parameter
					// TODO: Add check for include_audit query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"entity": "Entity", "description": "Description", "metadata": {"mapKey": "anyValue"}, "fuzzy_match": true, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "values": [{"value": "Value", "metadata": {"mapKey": "anyValue"}, "type": "synonyms", "synonyms": ["Synonym"], "patterns": ["Pattern"], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke GetEntity successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.GetEntity(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetEntityOptions model
				getEntityOptionsModel := new(assistantv1.GetEntityOptions)
				getEntityOptionsModel.WorkspaceID = core.StringPtr("testString")
				getEntityOptionsModel.Entity = core.StringPtr("testString")
				getEntityOptionsModel.Export = core.BoolPtr(false)
				getEntityOptionsModel.IncludeAudit = core.BoolPtr(false)
				getEntityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.GetEntity(getEntityOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetEntity with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the GetEntityOptions model
				getEntityOptionsModel := new(assistantv1.GetEntityOptions)
				getEntityOptionsModel.WorkspaceID = core.StringPtr("testString")
				getEntityOptionsModel.Entity = core.StringPtr("testString")
				getEntityOptionsModel.Export = core.BoolPtr(false)
				getEntityOptionsModel.IncludeAudit = core.BoolPtr(false)
				getEntityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.GetEntity(getEntityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetEntityOptions model with no property values
				getEntityOptionsModelNew := new(assistantv1.GetEntityOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = assistantService.GetEntity(getEntityOptionsModelNew)
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
			It(`Invoke GetEntity successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the GetEntityOptions model
				getEntityOptionsModel := new(assistantv1.GetEntityOptions)
				getEntityOptionsModel.WorkspaceID = core.StringPtr("testString")
				getEntityOptionsModel.Entity = core.StringPtr("testString")
				getEntityOptionsModel.Export = core.BoolPtr(false)
				getEntityOptionsModel.IncludeAudit = core.BoolPtr(false)
				getEntityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.GetEntity(getEntityOptionsModel)
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
	Describe(`UpdateEntity(updateEntityOptions *UpdateEntityOptions) - Operation response error`, func() {
		version := "testString"
		updateEntityPath := "/v1/workspaces/testString/entities/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateEntityPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for append query parameter
					// TODO: Add check for include_audit query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateEntity with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the CreateValue model
				createValueModel := new(assistantv1.CreateValue)
				createValueModel.Value = core.StringPtr("testString")
				createValueModel.Metadata = make(map[string]interface{})
				createValueModel.Type = core.StringPtr("synonyms")
				createValueModel.Synonyms = []string{"testString"}
				createValueModel.Patterns = []string{"testString"}

				// Construct an instance of the UpdateEntityOptions model
				updateEntityOptionsModel := new(assistantv1.UpdateEntityOptions)
				updateEntityOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateEntityOptionsModel.Entity = core.StringPtr("testString")
				updateEntityOptionsModel.NewEntity = core.StringPtr("testString")
				updateEntityOptionsModel.NewDescription = core.StringPtr("testString")
				updateEntityOptionsModel.NewMetadata = make(map[string]interface{})
				updateEntityOptionsModel.NewFuzzyMatch = core.BoolPtr(true)
				updateEntityOptionsModel.NewValues = []assistantv1.CreateValue{*createValueModel}
				updateEntityOptionsModel.Append = core.BoolPtr(false)
				updateEntityOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateEntityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.UpdateEntity(updateEntityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.UpdateEntity(updateEntityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateEntity(updateEntityOptions *UpdateEntityOptions)`, func() {
		version := "testString"
		updateEntityPath := "/v1/workspaces/testString/entities/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateEntityPath))
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
					// TODO: Add check for append query parameter
					// TODO: Add check for include_audit query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"entity": "Entity", "description": "Description", "metadata": {"mapKey": "anyValue"}, "fuzzy_match": true, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "values": [{"value": "Value", "metadata": {"mapKey": "anyValue"}, "type": "synonyms", "synonyms": ["Synonym"], "patterns": ["Pattern"], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke UpdateEntity successfully with retries`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the CreateValue model
				createValueModel := new(assistantv1.CreateValue)
				createValueModel.Value = core.StringPtr("testString")
				createValueModel.Metadata = make(map[string]interface{})
				createValueModel.Type = core.StringPtr("synonyms")
				createValueModel.Synonyms = []string{"testString"}
				createValueModel.Patterns = []string{"testString"}

				// Construct an instance of the UpdateEntityOptions model
				updateEntityOptionsModel := new(assistantv1.UpdateEntityOptions)
				updateEntityOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateEntityOptionsModel.Entity = core.StringPtr("testString")
				updateEntityOptionsModel.NewEntity = core.StringPtr("testString")
				updateEntityOptionsModel.NewDescription = core.StringPtr("testString")
				updateEntityOptionsModel.NewMetadata = make(map[string]interface{})
				updateEntityOptionsModel.NewFuzzyMatch = core.BoolPtr(true)
				updateEntityOptionsModel.NewValues = []assistantv1.CreateValue{*createValueModel}
				updateEntityOptionsModel.Append = core.BoolPtr(false)
				updateEntityOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateEntityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.UpdateEntityWithContext(ctx, updateEntityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.UpdateEntity(updateEntityOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.UpdateEntityWithContext(ctx, updateEntityOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateEntityPath))
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
					// TODO: Add check for append query parameter
					// TODO: Add check for include_audit query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"entity": "Entity", "description": "Description", "metadata": {"mapKey": "anyValue"}, "fuzzy_match": true, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "values": [{"value": "Value", "metadata": {"mapKey": "anyValue"}, "type": "synonyms", "synonyms": ["Synonym"], "patterns": ["Pattern"], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke UpdateEntity successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.UpdateEntity(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateValue model
				createValueModel := new(assistantv1.CreateValue)
				createValueModel.Value = core.StringPtr("testString")
				createValueModel.Metadata = make(map[string]interface{})
				createValueModel.Type = core.StringPtr("synonyms")
				createValueModel.Synonyms = []string{"testString"}
				createValueModel.Patterns = []string{"testString"}

				// Construct an instance of the UpdateEntityOptions model
				updateEntityOptionsModel := new(assistantv1.UpdateEntityOptions)
				updateEntityOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateEntityOptionsModel.Entity = core.StringPtr("testString")
				updateEntityOptionsModel.NewEntity = core.StringPtr("testString")
				updateEntityOptionsModel.NewDescription = core.StringPtr("testString")
				updateEntityOptionsModel.NewMetadata = make(map[string]interface{})
				updateEntityOptionsModel.NewFuzzyMatch = core.BoolPtr(true)
				updateEntityOptionsModel.NewValues = []assistantv1.CreateValue{*createValueModel}
				updateEntityOptionsModel.Append = core.BoolPtr(false)
				updateEntityOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateEntityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.UpdateEntity(updateEntityOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateEntity with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the CreateValue model
				createValueModel := new(assistantv1.CreateValue)
				createValueModel.Value = core.StringPtr("testString")
				createValueModel.Metadata = make(map[string]interface{})
				createValueModel.Type = core.StringPtr("synonyms")
				createValueModel.Synonyms = []string{"testString"}
				createValueModel.Patterns = []string{"testString"}

				// Construct an instance of the UpdateEntityOptions model
				updateEntityOptionsModel := new(assistantv1.UpdateEntityOptions)
				updateEntityOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateEntityOptionsModel.Entity = core.StringPtr("testString")
				updateEntityOptionsModel.NewEntity = core.StringPtr("testString")
				updateEntityOptionsModel.NewDescription = core.StringPtr("testString")
				updateEntityOptionsModel.NewMetadata = make(map[string]interface{})
				updateEntityOptionsModel.NewFuzzyMatch = core.BoolPtr(true)
				updateEntityOptionsModel.NewValues = []assistantv1.CreateValue{*createValueModel}
				updateEntityOptionsModel.Append = core.BoolPtr(false)
				updateEntityOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateEntityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.UpdateEntity(updateEntityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateEntityOptions model with no property values
				updateEntityOptionsModelNew := new(assistantv1.UpdateEntityOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = assistantService.UpdateEntity(updateEntityOptionsModelNew)
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
			It(`Invoke UpdateEntity successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the CreateValue model
				createValueModel := new(assistantv1.CreateValue)
				createValueModel.Value = core.StringPtr("testString")
				createValueModel.Metadata = make(map[string]interface{})
				createValueModel.Type = core.StringPtr("synonyms")
				createValueModel.Synonyms = []string{"testString"}
				createValueModel.Patterns = []string{"testString"}

				// Construct an instance of the UpdateEntityOptions model
				updateEntityOptionsModel := new(assistantv1.UpdateEntityOptions)
				updateEntityOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateEntityOptionsModel.Entity = core.StringPtr("testString")
				updateEntityOptionsModel.NewEntity = core.StringPtr("testString")
				updateEntityOptionsModel.NewDescription = core.StringPtr("testString")
				updateEntityOptionsModel.NewMetadata = make(map[string]interface{})
				updateEntityOptionsModel.NewFuzzyMatch = core.BoolPtr(true)
				updateEntityOptionsModel.NewValues = []assistantv1.CreateValue{*createValueModel}
				updateEntityOptionsModel.Append = core.BoolPtr(false)
				updateEntityOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateEntityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.UpdateEntity(updateEntityOptionsModel)
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
	Describe(`DeleteEntity(deleteEntityOptions *DeleteEntityOptions)`, func() {
		version := "testString"
		deleteEntityPath := "/v1/workspaces/testString/entities/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteEntityPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteEntity successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := assistantService.DeleteEntity(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteEntityOptions model
				deleteEntityOptionsModel := new(assistantv1.DeleteEntityOptions)
				deleteEntityOptionsModel.WorkspaceID = core.StringPtr("testString")
				deleteEntityOptionsModel.Entity = core.StringPtr("testString")
				deleteEntityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = assistantService.DeleteEntity(deleteEntityOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteEntity with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the DeleteEntityOptions model
				deleteEntityOptionsModel := new(assistantv1.DeleteEntityOptions)
				deleteEntityOptionsModel.WorkspaceID = core.StringPtr("testString")
				deleteEntityOptionsModel.Entity = core.StringPtr("testString")
				deleteEntityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := assistantService.DeleteEntity(deleteEntityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteEntityOptions model with no property values
				deleteEntityOptionsModelNew := new(assistantv1.DeleteEntityOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = assistantService.DeleteEntity(deleteEntityOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListMentions(listMentionsOptions *ListMentionsOptions) - Operation response error`, func() {
		version := "testString"
		listMentionsPath := "/v1/workspaces/testString/entities/testString/mentions"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listMentionsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for export query parameter
					// TODO: Add check for include_audit query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListMentions with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ListMentionsOptions model
				listMentionsOptionsModel := new(assistantv1.ListMentionsOptions)
				listMentionsOptionsModel.WorkspaceID = core.StringPtr("testString")
				listMentionsOptionsModel.Entity = core.StringPtr("testString")
				listMentionsOptionsModel.Export = core.BoolPtr(false)
				listMentionsOptionsModel.IncludeAudit = core.BoolPtr(false)
				listMentionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.ListMentions(listMentionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.ListMentions(listMentionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListMentions(listMentionsOptions *ListMentionsOptions)`, func() {
		version := "testString"
		listMentionsPath := "/v1/workspaces/testString/entities/testString/mentions"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listMentionsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for export query parameter
					// TODO: Add check for include_audit query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"examples": [{"text": "Text", "intent": "Intent", "location": [8]}], "pagination": {"refresh_url": "RefreshURL", "next_url": "NextURL", "total": 5, "matched": 7, "refresh_cursor": "RefreshCursor", "next_cursor": "NextCursor"}}`)
				}))
			})
			It(`Invoke ListMentions successfully with retries`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the ListMentionsOptions model
				listMentionsOptionsModel := new(assistantv1.ListMentionsOptions)
				listMentionsOptionsModel.WorkspaceID = core.StringPtr("testString")
				listMentionsOptionsModel.Entity = core.StringPtr("testString")
				listMentionsOptionsModel.Export = core.BoolPtr(false)
				listMentionsOptionsModel.IncludeAudit = core.BoolPtr(false)
				listMentionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.ListMentionsWithContext(ctx, listMentionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.ListMentions(listMentionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.ListMentionsWithContext(ctx, listMentionsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listMentionsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for export query parameter
					// TODO: Add check for include_audit query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"examples": [{"text": "Text", "intent": "Intent", "location": [8]}], "pagination": {"refresh_url": "RefreshURL", "next_url": "NextURL", "total": 5, "matched": 7, "refresh_cursor": "RefreshCursor", "next_cursor": "NextCursor"}}`)
				}))
			})
			It(`Invoke ListMentions successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.ListMentions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListMentionsOptions model
				listMentionsOptionsModel := new(assistantv1.ListMentionsOptions)
				listMentionsOptionsModel.WorkspaceID = core.StringPtr("testString")
				listMentionsOptionsModel.Entity = core.StringPtr("testString")
				listMentionsOptionsModel.Export = core.BoolPtr(false)
				listMentionsOptionsModel.IncludeAudit = core.BoolPtr(false)
				listMentionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.ListMentions(listMentionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListMentions with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ListMentionsOptions model
				listMentionsOptionsModel := new(assistantv1.ListMentionsOptions)
				listMentionsOptionsModel.WorkspaceID = core.StringPtr("testString")
				listMentionsOptionsModel.Entity = core.StringPtr("testString")
				listMentionsOptionsModel.Export = core.BoolPtr(false)
				listMentionsOptionsModel.IncludeAudit = core.BoolPtr(false)
				listMentionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.ListMentions(listMentionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListMentionsOptions model with no property values
				listMentionsOptionsModelNew := new(assistantv1.ListMentionsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = assistantService.ListMentions(listMentionsOptionsModelNew)
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
			It(`Invoke ListMentions successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ListMentionsOptions model
				listMentionsOptionsModel := new(assistantv1.ListMentionsOptions)
				listMentionsOptionsModel.WorkspaceID = core.StringPtr("testString")
				listMentionsOptionsModel.Entity = core.StringPtr("testString")
				listMentionsOptionsModel.Export = core.BoolPtr(false)
				listMentionsOptionsModel.IncludeAudit = core.BoolPtr(false)
				listMentionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.ListMentions(listMentionsOptionsModel)
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
	Describe(`ListValues(listValuesOptions *ListValuesOptions) - Operation response error`, func() {
		version := "testString"
		listValuesPath := "/v1/workspaces/testString/entities/testString/values"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listValuesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for export query parameter
					Expect(req.URL.Query()["page_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// TODO: Add check for include_count query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"value"}))
					Expect(req.URL.Query()["cursor"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListValues with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ListValuesOptions model
				listValuesOptionsModel := new(assistantv1.ListValuesOptions)
				listValuesOptionsModel.WorkspaceID = core.StringPtr("testString")
				listValuesOptionsModel.Entity = core.StringPtr("testString")
				listValuesOptionsModel.Export = core.BoolPtr(false)
				listValuesOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listValuesOptionsModel.IncludeCount = core.BoolPtr(false)
				listValuesOptionsModel.Sort = core.StringPtr("value")
				listValuesOptionsModel.Cursor = core.StringPtr("testString")
				listValuesOptionsModel.IncludeAudit = core.BoolPtr(false)
				listValuesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.ListValues(listValuesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.ListValues(listValuesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListValues(listValuesOptions *ListValuesOptions)`, func() {
		version := "testString"
		listValuesPath := "/v1/workspaces/testString/entities/testString/values"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listValuesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for export query parameter
					Expect(req.URL.Query()["page_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// TODO: Add check for include_count query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"value"}))
					Expect(req.URL.Query()["cursor"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"values": [{"value": "Value", "metadata": {"mapKey": "anyValue"}, "type": "synonyms", "synonyms": ["Synonym"], "patterns": ["Pattern"], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}], "pagination": {"refresh_url": "RefreshURL", "next_url": "NextURL", "total": 5, "matched": 7, "refresh_cursor": "RefreshCursor", "next_cursor": "NextCursor"}}`)
				}))
			})
			It(`Invoke ListValues successfully with retries`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the ListValuesOptions model
				listValuesOptionsModel := new(assistantv1.ListValuesOptions)
				listValuesOptionsModel.WorkspaceID = core.StringPtr("testString")
				listValuesOptionsModel.Entity = core.StringPtr("testString")
				listValuesOptionsModel.Export = core.BoolPtr(false)
				listValuesOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listValuesOptionsModel.IncludeCount = core.BoolPtr(false)
				listValuesOptionsModel.Sort = core.StringPtr("value")
				listValuesOptionsModel.Cursor = core.StringPtr("testString")
				listValuesOptionsModel.IncludeAudit = core.BoolPtr(false)
				listValuesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.ListValuesWithContext(ctx, listValuesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.ListValues(listValuesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.ListValuesWithContext(ctx, listValuesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listValuesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for export query parameter
					Expect(req.URL.Query()["page_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// TODO: Add check for include_count query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"value"}))
					Expect(req.URL.Query()["cursor"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"values": [{"value": "Value", "metadata": {"mapKey": "anyValue"}, "type": "synonyms", "synonyms": ["Synonym"], "patterns": ["Pattern"], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}], "pagination": {"refresh_url": "RefreshURL", "next_url": "NextURL", "total": 5, "matched": 7, "refresh_cursor": "RefreshCursor", "next_cursor": "NextCursor"}}`)
				}))
			})
			It(`Invoke ListValues successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.ListValues(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListValuesOptions model
				listValuesOptionsModel := new(assistantv1.ListValuesOptions)
				listValuesOptionsModel.WorkspaceID = core.StringPtr("testString")
				listValuesOptionsModel.Entity = core.StringPtr("testString")
				listValuesOptionsModel.Export = core.BoolPtr(false)
				listValuesOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listValuesOptionsModel.IncludeCount = core.BoolPtr(false)
				listValuesOptionsModel.Sort = core.StringPtr("value")
				listValuesOptionsModel.Cursor = core.StringPtr("testString")
				listValuesOptionsModel.IncludeAudit = core.BoolPtr(false)
				listValuesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.ListValues(listValuesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListValues with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ListValuesOptions model
				listValuesOptionsModel := new(assistantv1.ListValuesOptions)
				listValuesOptionsModel.WorkspaceID = core.StringPtr("testString")
				listValuesOptionsModel.Entity = core.StringPtr("testString")
				listValuesOptionsModel.Export = core.BoolPtr(false)
				listValuesOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listValuesOptionsModel.IncludeCount = core.BoolPtr(false)
				listValuesOptionsModel.Sort = core.StringPtr("value")
				listValuesOptionsModel.Cursor = core.StringPtr("testString")
				listValuesOptionsModel.IncludeAudit = core.BoolPtr(false)
				listValuesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.ListValues(listValuesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListValuesOptions model with no property values
				listValuesOptionsModelNew := new(assistantv1.ListValuesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = assistantService.ListValues(listValuesOptionsModelNew)
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
			It(`Invoke ListValues successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ListValuesOptions model
				listValuesOptionsModel := new(assistantv1.ListValuesOptions)
				listValuesOptionsModel.WorkspaceID = core.StringPtr("testString")
				listValuesOptionsModel.Entity = core.StringPtr("testString")
				listValuesOptionsModel.Export = core.BoolPtr(false)
				listValuesOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listValuesOptionsModel.IncludeCount = core.BoolPtr(false)
				listValuesOptionsModel.Sort = core.StringPtr("value")
				listValuesOptionsModel.Cursor = core.StringPtr("testString")
				listValuesOptionsModel.IncludeAudit = core.BoolPtr(false)
				listValuesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.ListValues(listValuesOptionsModel)
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
	Describe(`CreateValue(createValueOptions *CreateValueOptions) - Operation response error`, func() {
		version := "testString"
		createValuePath := "/v1/workspaces/testString/entities/testString/values"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createValuePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateValue with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the CreateValueOptions model
				createValueOptionsModel := new(assistantv1.CreateValueOptions)
				createValueOptionsModel.WorkspaceID = core.StringPtr("testString")
				createValueOptionsModel.Entity = core.StringPtr("testString")
				createValueOptionsModel.Value = core.StringPtr("testString")
				createValueOptionsModel.Metadata = make(map[string]interface{})
				createValueOptionsModel.Type = core.StringPtr("synonyms")
				createValueOptionsModel.Synonyms = []string{"testString"}
				createValueOptionsModel.Patterns = []string{"testString"}
				createValueOptionsModel.IncludeAudit = core.BoolPtr(false)
				createValueOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.CreateValue(createValueOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.CreateValue(createValueOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateValue(createValueOptions *CreateValueOptions)`, func() {
		version := "testString"
		createValuePath := "/v1/workspaces/testString/entities/testString/values"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createValuePath))
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
					// TODO: Add check for include_audit query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"value": "Value", "metadata": {"mapKey": "anyValue"}, "type": "synonyms", "synonyms": ["Synonym"], "patterns": ["Pattern"], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke CreateValue successfully with retries`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the CreateValueOptions model
				createValueOptionsModel := new(assistantv1.CreateValueOptions)
				createValueOptionsModel.WorkspaceID = core.StringPtr("testString")
				createValueOptionsModel.Entity = core.StringPtr("testString")
				createValueOptionsModel.Value = core.StringPtr("testString")
				createValueOptionsModel.Metadata = make(map[string]interface{})
				createValueOptionsModel.Type = core.StringPtr("synonyms")
				createValueOptionsModel.Synonyms = []string{"testString"}
				createValueOptionsModel.Patterns = []string{"testString"}
				createValueOptionsModel.IncludeAudit = core.BoolPtr(false)
				createValueOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.CreateValueWithContext(ctx, createValueOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.CreateValue(createValueOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.CreateValueWithContext(ctx, createValueOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createValuePath))
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
					// TODO: Add check for include_audit query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"value": "Value", "metadata": {"mapKey": "anyValue"}, "type": "synonyms", "synonyms": ["Synonym"], "patterns": ["Pattern"], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke CreateValue successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.CreateValue(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateValueOptions model
				createValueOptionsModel := new(assistantv1.CreateValueOptions)
				createValueOptionsModel.WorkspaceID = core.StringPtr("testString")
				createValueOptionsModel.Entity = core.StringPtr("testString")
				createValueOptionsModel.Value = core.StringPtr("testString")
				createValueOptionsModel.Metadata = make(map[string]interface{})
				createValueOptionsModel.Type = core.StringPtr("synonyms")
				createValueOptionsModel.Synonyms = []string{"testString"}
				createValueOptionsModel.Patterns = []string{"testString"}
				createValueOptionsModel.IncludeAudit = core.BoolPtr(false)
				createValueOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.CreateValue(createValueOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateValue with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the CreateValueOptions model
				createValueOptionsModel := new(assistantv1.CreateValueOptions)
				createValueOptionsModel.WorkspaceID = core.StringPtr("testString")
				createValueOptionsModel.Entity = core.StringPtr("testString")
				createValueOptionsModel.Value = core.StringPtr("testString")
				createValueOptionsModel.Metadata = make(map[string]interface{})
				createValueOptionsModel.Type = core.StringPtr("synonyms")
				createValueOptionsModel.Synonyms = []string{"testString"}
				createValueOptionsModel.Patterns = []string{"testString"}
				createValueOptionsModel.IncludeAudit = core.BoolPtr(false)
				createValueOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.CreateValue(createValueOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateValueOptions model with no property values
				createValueOptionsModelNew := new(assistantv1.CreateValueOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = assistantService.CreateValue(createValueOptionsModelNew)
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
			It(`Invoke CreateValue successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the CreateValueOptions model
				createValueOptionsModel := new(assistantv1.CreateValueOptions)
				createValueOptionsModel.WorkspaceID = core.StringPtr("testString")
				createValueOptionsModel.Entity = core.StringPtr("testString")
				createValueOptionsModel.Value = core.StringPtr("testString")
				createValueOptionsModel.Metadata = make(map[string]interface{})
				createValueOptionsModel.Type = core.StringPtr("synonyms")
				createValueOptionsModel.Synonyms = []string{"testString"}
				createValueOptionsModel.Patterns = []string{"testString"}
				createValueOptionsModel.IncludeAudit = core.BoolPtr(false)
				createValueOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.CreateValue(createValueOptionsModel)
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
	Describe(`GetValue(getValueOptions *GetValueOptions) - Operation response error`, func() {
		version := "testString"
		getValuePath := "/v1/workspaces/testString/entities/testString/values/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getValuePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for export query parameter
					// TODO: Add check for include_audit query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetValue with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the GetValueOptions model
				getValueOptionsModel := new(assistantv1.GetValueOptions)
				getValueOptionsModel.WorkspaceID = core.StringPtr("testString")
				getValueOptionsModel.Entity = core.StringPtr("testString")
				getValueOptionsModel.Value = core.StringPtr("testString")
				getValueOptionsModel.Export = core.BoolPtr(false)
				getValueOptionsModel.IncludeAudit = core.BoolPtr(false)
				getValueOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.GetValue(getValueOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.GetValue(getValueOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetValue(getValueOptions *GetValueOptions)`, func() {
		version := "testString"
		getValuePath := "/v1/workspaces/testString/entities/testString/values/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getValuePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for export query parameter
					// TODO: Add check for include_audit query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"value": "Value", "metadata": {"mapKey": "anyValue"}, "type": "synonyms", "synonyms": ["Synonym"], "patterns": ["Pattern"], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetValue successfully with retries`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the GetValueOptions model
				getValueOptionsModel := new(assistantv1.GetValueOptions)
				getValueOptionsModel.WorkspaceID = core.StringPtr("testString")
				getValueOptionsModel.Entity = core.StringPtr("testString")
				getValueOptionsModel.Value = core.StringPtr("testString")
				getValueOptionsModel.Export = core.BoolPtr(false)
				getValueOptionsModel.IncludeAudit = core.BoolPtr(false)
				getValueOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.GetValueWithContext(ctx, getValueOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.GetValue(getValueOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.GetValueWithContext(ctx, getValueOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getValuePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for export query parameter
					// TODO: Add check for include_audit query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"value": "Value", "metadata": {"mapKey": "anyValue"}, "type": "synonyms", "synonyms": ["Synonym"], "patterns": ["Pattern"], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetValue successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.GetValue(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetValueOptions model
				getValueOptionsModel := new(assistantv1.GetValueOptions)
				getValueOptionsModel.WorkspaceID = core.StringPtr("testString")
				getValueOptionsModel.Entity = core.StringPtr("testString")
				getValueOptionsModel.Value = core.StringPtr("testString")
				getValueOptionsModel.Export = core.BoolPtr(false)
				getValueOptionsModel.IncludeAudit = core.BoolPtr(false)
				getValueOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.GetValue(getValueOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetValue with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the GetValueOptions model
				getValueOptionsModel := new(assistantv1.GetValueOptions)
				getValueOptionsModel.WorkspaceID = core.StringPtr("testString")
				getValueOptionsModel.Entity = core.StringPtr("testString")
				getValueOptionsModel.Value = core.StringPtr("testString")
				getValueOptionsModel.Export = core.BoolPtr(false)
				getValueOptionsModel.IncludeAudit = core.BoolPtr(false)
				getValueOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.GetValue(getValueOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetValueOptions model with no property values
				getValueOptionsModelNew := new(assistantv1.GetValueOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = assistantService.GetValue(getValueOptionsModelNew)
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
			It(`Invoke GetValue successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the GetValueOptions model
				getValueOptionsModel := new(assistantv1.GetValueOptions)
				getValueOptionsModel.WorkspaceID = core.StringPtr("testString")
				getValueOptionsModel.Entity = core.StringPtr("testString")
				getValueOptionsModel.Value = core.StringPtr("testString")
				getValueOptionsModel.Export = core.BoolPtr(false)
				getValueOptionsModel.IncludeAudit = core.BoolPtr(false)
				getValueOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.GetValue(getValueOptionsModel)
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
	Describe(`UpdateValue(updateValueOptions *UpdateValueOptions) - Operation response error`, func() {
		version := "testString"
		updateValuePath := "/v1/workspaces/testString/entities/testString/values/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateValuePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for append query parameter
					// TODO: Add check for include_audit query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateValue with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the UpdateValueOptions model
				updateValueOptionsModel := new(assistantv1.UpdateValueOptions)
				updateValueOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateValueOptionsModel.Entity = core.StringPtr("testString")
				updateValueOptionsModel.Value = core.StringPtr("testString")
				updateValueOptionsModel.NewValue = core.StringPtr("testString")
				updateValueOptionsModel.NewMetadata = make(map[string]interface{})
				updateValueOptionsModel.NewType = core.StringPtr("synonyms")
				updateValueOptionsModel.NewSynonyms = []string{"testString"}
				updateValueOptionsModel.NewPatterns = []string{"testString"}
				updateValueOptionsModel.Append = core.BoolPtr(false)
				updateValueOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateValueOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.UpdateValue(updateValueOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.UpdateValue(updateValueOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateValue(updateValueOptions *UpdateValueOptions)`, func() {
		version := "testString"
		updateValuePath := "/v1/workspaces/testString/entities/testString/values/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateValuePath))
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
					// TODO: Add check for append query parameter
					// TODO: Add check for include_audit query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"value": "Value", "metadata": {"mapKey": "anyValue"}, "type": "synonyms", "synonyms": ["Synonym"], "patterns": ["Pattern"], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke UpdateValue successfully with retries`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the UpdateValueOptions model
				updateValueOptionsModel := new(assistantv1.UpdateValueOptions)
				updateValueOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateValueOptionsModel.Entity = core.StringPtr("testString")
				updateValueOptionsModel.Value = core.StringPtr("testString")
				updateValueOptionsModel.NewValue = core.StringPtr("testString")
				updateValueOptionsModel.NewMetadata = make(map[string]interface{})
				updateValueOptionsModel.NewType = core.StringPtr("synonyms")
				updateValueOptionsModel.NewSynonyms = []string{"testString"}
				updateValueOptionsModel.NewPatterns = []string{"testString"}
				updateValueOptionsModel.Append = core.BoolPtr(false)
				updateValueOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateValueOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.UpdateValueWithContext(ctx, updateValueOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.UpdateValue(updateValueOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.UpdateValueWithContext(ctx, updateValueOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateValuePath))
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
					// TODO: Add check for append query parameter
					// TODO: Add check for include_audit query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"value": "Value", "metadata": {"mapKey": "anyValue"}, "type": "synonyms", "synonyms": ["Synonym"], "patterns": ["Pattern"], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke UpdateValue successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.UpdateValue(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateValueOptions model
				updateValueOptionsModel := new(assistantv1.UpdateValueOptions)
				updateValueOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateValueOptionsModel.Entity = core.StringPtr("testString")
				updateValueOptionsModel.Value = core.StringPtr("testString")
				updateValueOptionsModel.NewValue = core.StringPtr("testString")
				updateValueOptionsModel.NewMetadata = make(map[string]interface{})
				updateValueOptionsModel.NewType = core.StringPtr("synonyms")
				updateValueOptionsModel.NewSynonyms = []string{"testString"}
				updateValueOptionsModel.NewPatterns = []string{"testString"}
				updateValueOptionsModel.Append = core.BoolPtr(false)
				updateValueOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateValueOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.UpdateValue(updateValueOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateValue with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the UpdateValueOptions model
				updateValueOptionsModel := new(assistantv1.UpdateValueOptions)
				updateValueOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateValueOptionsModel.Entity = core.StringPtr("testString")
				updateValueOptionsModel.Value = core.StringPtr("testString")
				updateValueOptionsModel.NewValue = core.StringPtr("testString")
				updateValueOptionsModel.NewMetadata = make(map[string]interface{})
				updateValueOptionsModel.NewType = core.StringPtr("synonyms")
				updateValueOptionsModel.NewSynonyms = []string{"testString"}
				updateValueOptionsModel.NewPatterns = []string{"testString"}
				updateValueOptionsModel.Append = core.BoolPtr(false)
				updateValueOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateValueOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.UpdateValue(updateValueOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateValueOptions model with no property values
				updateValueOptionsModelNew := new(assistantv1.UpdateValueOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = assistantService.UpdateValue(updateValueOptionsModelNew)
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
			It(`Invoke UpdateValue successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the UpdateValueOptions model
				updateValueOptionsModel := new(assistantv1.UpdateValueOptions)
				updateValueOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateValueOptionsModel.Entity = core.StringPtr("testString")
				updateValueOptionsModel.Value = core.StringPtr("testString")
				updateValueOptionsModel.NewValue = core.StringPtr("testString")
				updateValueOptionsModel.NewMetadata = make(map[string]interface{})
				updateValueOptionsModel.NewType = core.StringPtr("synonyms")
				updateValueOptionsModel.NewSynonyms = []string{"testString"}
				updateValueOptionsModel.NewPatterns = []string{"testString"}
				updateValueOptionsModel.Append = core.BoolPtr(false)
				updateValueOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateValueOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.UpdateValue(updateValueOptionsModel)
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
	Describe(`DeleteValue(deleteValueOptions *DeleteValueOptions)`, func() {
		version := "testString"
		deleteValuePath := "/v1/workspaces/testString/entities/testString/values/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteValuePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteValue successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := assistantService.DeleteValue(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteValueOptions model
				deleteValueOptionsModel := new(assistantv1.DeleteValueOptions)
				deleteValueOptionsModel.WorkspaceID = core.StringPtr("testString")
				deleteValueOptionsModel.Entity = core.StringPtr("testString")
				deleteValueOptionsModel.Value = core.StringPtr("testString")
				deleteValueOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = assistantService.DeleteValue(deleteValueOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteValue with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the DeleteValueOptions model
				deleteValueOptionsModel := new(assistantv1.DeleteValueOptions)
				deleteValueOptionsModel.WorkspaceID = core.StringPtr("testString")
				deleteValueOptionsModel.Entity = core.StringPtr("testString")
				deleteValueOptionsModel.Value = core.StringPtr("testString")
				deleteValueOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := assistantService.DeleteValue(deleteValueOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteValueOptions model with no property values
				deleteValueOptionsModelNew := new(assistantv1.DeleteValueOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = assistantService.DeleteValue(deleteValueOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListSynonyms(listSynonymsOptions *ListSynonymsOptions) - Operation response error`, func() {
		version := "testString"
		listSynonymsPath := "/v1/workspaces/testString/entities/testString/values/testString/synonyms"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSynonymsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["page_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// TODO: Add check for include_count query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"synonym"}))
					Expect(req.URL.Query()["cursor"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListSynonyms with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ListSynonymsOptions model
				listSynonymsOptionsModel := new(assistantv1.ListSynonymsOptions)
				listSynonymsOptionsModel.WorkspaceID = core.StringPtr("testString")
				listSynonymsOptionsModel.Entity = core.StringPtr("testString")
				listSynonymsOptionsModel.Value = core.StringPtr("testString")
				listSynonymsOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listSynonymsOptionsModel.IncludeCount = core.BoolPtr(false)
				listSynonymsOptionsModel.Sort = core.StringPtr("synonym")
				listSynonymsOptionsModel.Cursor = core.StringPtr("testString")
				listSynonymsOptionsModel.IncludeAudit = core.BoolPtr(false)
				listSynonymsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.ListSynonyms(listSynonymsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.ListSynonyms(listSynonymsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListSynonyms(listSynonymsOptions *ListSynonymsOptions)`, func() {
		version := "testString"
		listSynonymsPath := "/v1/workspaces/testString/entities/testString/values/testString/synonyms"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSynonymsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["page_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// TODO: Add check for include_count query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"synonym"}))
					Expect(req.URL.Query()["cursor"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"synonyms": [{"synonym": "Synonym", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}], "pagination": {"refresh_url": "RefreshURL", "next_url": "NextURL", "total": 5, "matched": 7, "refresh_cursor": "RefreshCursor", "next_cursor": "NextCursor"}}`)
				}))
			})
			It(`Invoke ListSynonyms successfully with retries`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the ListSynonymsOptions model
				listSynonymsOptionsModel := new(assistantv1.ListSynonymsOptions)
				listSynonymsOptionsModel.WorkspaceID = core.StringPtr("testString")
				listSynonymsOptionsModel.Entity = core.StringPtr("testString")
				listSynonymsOptionsModel.Value = core.StringPtr("testString")
				listSynonymsOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listSynonymsOptionsModel.IncludeCount = core.BoolPtr(false)
				listSynonymsOptionsModel.Sort = core.StringPtr("synonym")
				listSynonymsOptionsModel.Cursor = core.StringPtr("testString")
				listSynonymsOptionsModel.IncludeAudit = core.BoolPtr(false)
				listSynonymsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.ListSynonymsWithContext(ctx, listSynonymsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.ListSynonyms(listSynonymsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.ListSynonymsWithContext(ctx, listSynonymsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listSynonymsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["page_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// TODO: Add check for include_count query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"synonym"}))
					Expect(req.URL.Query()["cursor"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"synonyms": [{"synonym": "Synonym", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}], "pagination": {"refresh_url": "RefreshURL", "next_url": "NextURL", "total": 5, "matched": 7, "refresh_cursor": "RefreshCursor", "next_cursor": "NextCursor"}}`)
				}))
			})
			It(`Invoke ListSynonyms successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.ListSynonyms(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListSynonymsOptions model
				listSynonymsOptionsModel := new(assistantv1.ListSynonymsOptions)
				listSynonymsOptionsModel.WorkspaceID = core.StringPtr("testString")
				listSynonymsOptionsModel.Entity = core.StringPtr("testString")
				listSynonymsOptionsModel.Value = core.StringPtr("testString")
				listSynonymsOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listSynonymsOptionsModel.IncludeCount = core.BoolPtr(false)
				listSynonymsOptionsModel.Sort = core.StringPtr("synonym")
				listSynonymsOptionsModel.Cursor = core.StringPtr("testString")
				listSynonymsOptionsModel.IncludeAudit = core.BoolPtr(false)
				listSynonymsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.ListSynonyms(listSynonymsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListSynonyms with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ListSynonymsOptions model
				listSynonymsOptionsModel := new(assistantv1.ListSynonymsOptions)
				listSynonymsOptionsModel.WorkspaceID = core.StringPtr("testString")
				listSynonymsOptionsModel.Entity = core.StringPtr("testString")
				listSynonymsOptionsModel.Value = core.StringPtr("testString")
				listSynonymsOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listSynonymsOptionsModel.IncludeCount = core.BoolPtr(false)
				listSynonymsOptionsModel.Sort = core.StringPtr("synonym")
				listSynonymsOptionsModel.Cursor = core.StringPtr("testString")
				listSynonymsOptionsModel.IncludeAudit = core.BoolPtr(false)
				listSynonymsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.ListSynonyms(listSynonymsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListSynonymsOptions model with no property values
				listSynonymsOptionsModelNew := new(assistantv1.ListSynonymsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = assistantService.ListSynonyms(listSynonymsOptionsModelNew)
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
			It(`Invoke ListSynonyms successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ListSynonymsOptions model
				listSynonymsOptionsModel := new(assistantv1.ListSynonymsOptions)
				listSynonymsOptionsModel.WorkspaceID = core.StringPtr("testString")
				listSynonymsOptionsModel.Entity = core.StringPtr("testString")
				listSynonymsOptionsModel.Value = core.StringPtr("testString")
				listSynonymsOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listSynonymsOptionsModel.IncludeCount = core.BoolPtr(false)
				listSynonymsOptionsModel.Sort = core.StringPtr("synonym")
				listSynonymsOptionsModel.Cursor = core.StringPtr("testString")
				listSynonymsOptionsModel.IncludeAudit = core.BoolPtr(false)
				listSynonymsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.ListSynonyms(listSynonymsOptionsModel)
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
	Describe(`CreateSynonym(createSynonymOptions *CreateSynonymOptions) - Operation response error`, func() {
		version := "testString"
		createSynonymPath := "/v1/workspaces/testString/entities/testString/values/testString/synonyms"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSynonymPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateSynonym with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the CreateSynonymOptions model
				createSynonymOptionsModel := new(assistantv1.CreateSynonymOptions)
				createSynonymOptionsModel.WorkspaceID = core.StringPtr("testString")
				createSynonymOptionsModel.Entity = core.StringPtr("testString")
				createSynonymOptionsModel.Value = core.StringPtr("testString")
				createSynonymOptionsModel.Synonym = core.StringPtr("testString")
				createSynonymOptionsModel.IncludeAudit = core.BoolPtr(false)
				createSynonymOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.CreateSynonym(createSynonymOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.CreateSynonym(createSynonymOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateSynonym(createSynonymOptions *CreateSynonymOptions)`, func() {
		version := "testString"
		createSynonymPath := "/v1/workspaces/testString/entities/testString/values/testString/synonyms"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSynonymPath))
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
					// TODO: Add check for include_audit query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"synonym": "Synonym", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke CreateSynonym successfully with retries`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the CreateSynonymOptions model
				createSynonymOptionsModel := new(assistantv1.CreateSynonymOptions)
				createSynonymOptionsModel.WorkspaceID = core.StringPtr("testString")
				createSynonymOptionsModel.Entity = core.StringPtr("testString")
				createSynonymOptionsModel.Value = core.StringPtr("testString")
				createSynonymOptionsModel.Synonym = core.StringPtr("testString")
				createSynonymOptionsModel.IncludeAudit = core.BoolPtr(false)
				createSynonymOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.CreateSynonymWithContext(ctx, createSynonymOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.CreateSynonym(createSynonymOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.CreateSynonymWithContext(ctx, createSynonymOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createSynonymPath))
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
					// TODO: Add check for include_audit query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"synonym": "Synonym", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke CreateSynonym successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.CreateSynonym(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateSynonymOptions model
				createSynonymOptionsModel := new(assistantv1.CreateSynonymOptions)
				createSynonymOptionsModel.WorkspaceID = core.StringPtr("testString")
				createSynonymOptionsModel.Entity = core.StringPtr("testString")
				createSynonymOptionsModel.Value = core.StringPtr("testString")
				createSynonymOptionsModel.Synonym = core.StringPtr("testString")
				createSynonymOptionsModel.IncludeAudit = core.BoolPtr(false)
				createSynonymOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.CreateSynonym(createSynonymOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateSynonym with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the CreateSynonymOptions model
				createSynonymOptionsModel := new(assistantv1.CreateSynonymOptions)
				createSynonymOptionsModel.WorkspaceID = core.StringPtr("testString")
				createSynonymOptionsModel.Entity = core.StringPtr("testString")
				createSynonymOptionsModel.Value = core.StringPtr("testString")
				createSynonymOptionsModel.Synonym = core.StringPtr("testString")
				createSynonymOptionsModel.IncludeAudit = core.BoolPtr(false)
				createSynonymOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.CreateSynonym(createSynonymOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateSynonymOptions model with no property values
				createSynonymOptionsModelNew := new(assistantv1.CreateSynonymOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = assistantService.CreateSynonym(createSynonymOptionsModelNew)
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
			It(`Invoke CreateSynonym successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the CreateSynonymOptions model
				createSynonymOptionsModel := new(assistantv1.CreateSynonymOptions)
				createSynonymOptionsModel.WorkspaceID = core.StringPtr("testString")
				createSynonymOptionsModel.Entity = core.StringPtr("testString")
				createSynonymOptionsModel.Value = core.StringPtr("testString")
				createSynonymOptionsModel.Synonym = core.StringPtr("testString")
				createSynonymOptionsModel.IncludeAudit = core.BoolPtr(false)
				createSynonymOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.CreateSynonym(createSynonymOptionsModel)
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
	Describe(`GetSynonym(getSynonymOptions *GetSynonymOptions) - Operation response error`, func() {
		version := "testString"
		getSynonymPath := "/v1/workspaces/testString/entities/testString/values/testString/synonyms/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSynonymPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSynonym with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the GetSynonymOptions model
				getSynonymOptionsModel := new(assistantv1.GetSynonymOptions)
				getSynonymOptionsModel.WorkspaceID = core.StringPtr("testString")
				getSynonymOptionsModel.Entity = core.StringPtr("testString")
				getSynonymOptionsModel.Value = core.StringPtr("testString")
				getSynonymOptionsModel.Synonym = core.StringPtr("testString")
				getSynonymOptionsModel.IncludeAudit = core.BoolPtr(false)
				getSynonymOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.GetSynonym(getSynonymOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.GetSynonym(getSynonymOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSynonym(getSynonymOptions *GetSynonymOptions)`, func() {
		version := "testString"
		getSynonymPath := "/v1/workspaces/testString/entities/testString/values/testString/synonyms/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSynonymPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"synonym": "Synonym", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetSynonym successfully with retries`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the GetSynonymOptions model
				getSynonymOptionsModel := new(assistantv1.GetSynonymOptions)
				getSynonymOptionsModel.WorkspaceID = core.StringPtr("testString")
				getSynonymOptionsModel.Entity = core.StringPtr("testString")
				getSynonymOptionsModel.Value = core.StringPtr("testString")
				getSynonymOptionsModel.Synonym = core.StringPtr("testString")
				getSynonymOptionsModel.IncludeAudit = core.BoolPtr(false)
				getSynonymOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.GetSynonymWithContext(ctx, getSynonymOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.GetSynonym(getSynonymOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.GetSynonymWithContext(ctx, getSynonymOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getSynonymPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"synonym": "Synonym", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetSynonym successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.GetSynonym(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSynonymOptions model
				getSynonymOptionsModel := new(assistantv1.GetSynonymOptions)
				getSynonymOptionsModel.WorkspaceID = core.StringPtr("testString")
				getSynonymOptionsModel.Entity = core.StringPtr("testString")
				getSynonymOptionsModel.Value = core.StringPtr("testString")
				getSynonymOptionsModel.Synonym = core.StringPtr("testString")
				getSynonymOptionsModel.IncludeAudit = core.BoolPtr(false)
				getSynonymOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.GetSynonym(getSynonymOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetSynonym with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the GetSynonymOptions model
				getSynonymOptionsModel := new(assistantv1.GetSynonymOptions)
				getSynonymOptionsModel.WorkspaceID = core.StringPtr("testString")
				getSynonymOptionsModel.Entity = core.StringPtr("testString")
				getSynonymOptionsModel.Value = core.StringPtr("testString")
				getSynonymOptionsModel.Synonym = core.StringPtr("testString")
				getSynonymOptionsModel.IncludeAudit = core.BoolPtr(false)
				getSynonymOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.GetSynonym(getSynonymOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetSynonymOptions model with no property values
				getSynonymOptionsModelNew := new(assistantv1.GetSynonymOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = assistantService.GetSynonym(getSynonymOptionsModelNew)
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
			It(`Invoke GetSynonym successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the GetSynonymOptions model
				getSynonymOptionsModel := new(assistantv1.GetSynonymOptions)
				getSynonymOptionsModel.WorkspaceID = core.StringPtr("testString")
				getSynonymOptionsModel.Entity = core.StringPtr("testString")
				getSynonymOptionsModel.Value = core.StringPtr("testString")
				getSynonymOptionsModel.Synonym = core.StringPtr("testString")
				getSynonymOptionsModel.IncludeAudit = core.BoolPtr(false)
				getSynonymOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.GetSynonym(getSynonymOptionsModel)
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
	Describe(`UpdateSynonym(updateSynonymOptions *UpdateSynonymOptions) - Operation response error`, func() {
		version := "testString"
		updateSynonymPath := "/v1/workspaces/testString/entities/testString/values/testString/synonyms/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSynonymPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateSynonym with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the UpdateSynonymOptions model
				updateSynonymOptionsModel := new(assistantv1.UpdateSynonymOptions)
				updateSynonymOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateSynonymOptionsModel.Entity = core.StringPtr("testString")
				updateSynonymOptionsModel.Value = core.StringPtr("testString")
				updateSynonymOptionsModel.Synonym = core.StringPtr("testString")
				updateSynonymOptionsModel.NewSynonym = core.StringPtr("testString")
				updateSynonymOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateSynonymOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.UpdateSynonym(updateSynonymOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.UpdateSynonym(updateSynonymOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateSynonym(updateSynonymOptions *UpdateSynonymOptions)`, func() {
		version := "testString"
		updateSynonymPath := "/v1/workspaces/testString/entities/testString/values/testString/synonyms/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSynonymPath))
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
					// TODO: Add check for include_audit query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"synonym": "Synonym", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke UpdateSynonym successfully with retries`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the UpdateSynonymOptions model
				updateSynonymOptionsModel := new(assistantv1.UpdateSynonymOptions)
				updateSynonymOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateSynonymOptionsModel.Entity = core.StringPtr("testString")
				updateSynonymOptionsModel.Value = core.StringPtr("testString")
				updateSynonymOptionsModel.Synonym = core.StringPtr("testString")
				updateSynonymOptionsModel.NewSynonym = core.StringPtr("testString")
				updateSynonymOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateSynonymOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.UpdateSynonymWithContext(ctx, updateSynonymOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.UpdateSynonym(updateSynonymOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.UpdateSynonymWithContext(ctx, updateSynonymOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateSynonymPath))
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
					// TODO: Add check for include_audit query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"synonym": "Synonym", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke UpdateSynonym successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.UpdateSynonym(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateSynonymOptions model
				updateSynonymOptionsModel := new(assistantv1.UpdateSynonymOptions)
				updateSynonymOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateSynonymOptionsModel.Entity = core.StringPtr("testString")
				updateSynonymOptionsModel.Value = core.StringPtr("testString")
				updateSynonymOptionsModel.Synonym = core.StringPtr("testString")
				updateSynonymOptionsModel.NewSynonym = core.StringPtr("testString")
				updateSynonymOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateSynonymOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.UpdateSynonym(updateSynonymOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateSynonym with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the UpdateSynonymOptions model
				updateSynonymOptionsModel := new(assistantv1.UpdateSynonymOptions)
				updateSynonymOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateSynonymOptionsModel.Entity = core.StringPtr("testString")
				updateSynonymOptionsModel.Value = core.StringPtr("testString")
				updateSynonymOptionsModel.Synonym = core.StringPtr("testString")
				updateSynonymOptionsModel.NewSynonym = core.StringPtr("testString")
				updateSynonymOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateSynonymOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.UpdateSynonym(updateSynonymOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateSynonymOptions model with no property values
				updateSynonymOptionsModelNew := new(assistantv1.UpdateSynonymOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = assistantService.UpdateSynonym(updateSynonymOptionsModelNew)
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
			It(`Invoke UpdateSynonym successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the UpdateSynonymOptions model
				updateSynonymOptionsModel := new(assistantv1.UpdateSynonymOptions)
				updateSynonymOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateSynonymOptionsModel.Entity = core.StringPtr("testString")
				updateSynonymOptionsModel.Value = core.StringPtr("testString")
				updateSynonymOptionsModel.Synonym = core.StringPtr("testString")
				updateSynonymOptionsModel.NewSynonym = core.StringPtr("testString")
				updateSynonymOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateSynonymOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.UpdateSynonym(updateSynonymOptionsModel)
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
	Describe(`DeleteSynonym(deleteSynonymOptions *DeleteSynonymOptions)`, func() {
		version := "testString"
		deleteSynonymPath := "/v1/workspaces/testString/entities/testString/values/testString/synonyms/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteSynonymPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteSynonym successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := assistantService.DeleteSynonym(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteSynonymOptions model
				deleteSynonymOptionsModel := new(assistantv1.DeleteSynonymOptions)
				deleteSynonymOptionsModel.WorkspaceID = core.StringPtr("testString")
				deleteSynonymOptionsModel.Entity = core.StringPtr("testString")
				deleteSynonymOptionsModel.Value = core.StringPtr("testString")
				deleteSynonymOptionsModel.Synonym = core.StringPtr("testString")
				deleteSynonymOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = assistantService.DeleteSynonym(deleteSynonymOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteSynonym with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the DeleteSynonymOptions model
				deleteSynonymOptionsModel := new(assistantv1.DeleteSynonymOptions)
				deleteSynonymOptionsModel.WorkspaceID = core.StringPtr("testString")
				deleteSynonymOptionsModel.Entity = core.StringPtr("testString")
				deleteSynonymOptionsModel.Value = core.StringPtr("testString")
				deleteSynonymOptionsModel.Synonym = core.StringPtr("testString")
				deleteSynonymOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := assistantService.DeleteSynonym(deleteSynonymOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteSynonymOptions model with no property values
				deleteSynonymOptionsModelNew := new(assistantv1.DeleteSynonymOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = assistantService.DeleteSynonym(deleteSynonymOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDialogNodes(listDialogNodesOptions *ListDialogNodesOptions) - Operation response error`, func() {
		version := "testString"
		listDialogNodesPath := "/v1/workspaces/testString/dialog_nodes"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDialogNodesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["page_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// TODO: Add check for include_count query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"dialog_node"}))
					Expect(req.URL.Query()["cursor"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListDialogNodes with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ListDialogNodesOptions model
				listDialogNodesOptionsModel := new(assistantv1.ListDialogNodesOptions)
				listDialogNodesOptionsModel.WorkspaceID = core.StringPtr("testString")
				listDialogNodesOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listDialogNodesOptionsModel.IncludeCount = core.BoolPtr(false)
				listDialogNodesOptionsModel.Sort = core.StringPtr("dialog_node")
				listDialogNodesOptionsModel.Cursor = core.StringPtr("testString")
				listDialogNodesOptionsModel.IncludeAudit = core.BoolPtr(false)
				listDialogNodesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.ListDialogNodes(listDialogNodesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.ListDialogNodes(listDialogNodesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDialogNodes(listDialogNodesOptions *ListDialogNodesOptions)`, func() {
		version := "testString"
		listDialogNodesPath := "/v1/workspaces/testString/dialog_nodes"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDialogNodesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["page_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// TODO: Add check for include_count query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"dialog_node"}))
					Expect(req.URL.Query()["cursor"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"dialog_nodes": [{"dialog_node": "DialogNode", "description": "Description", "conditions": "Conditions", "parent": "Parent", "previous_sibling": "PreviousSibling", "output": {"generic": [{"response_type": "video", "source": "Source", "title": "Title", "description": "Description", "channels": [{"channel": "chat"}], "channel_options": {"anyKey": "anyValue"}, "alt_text": "AltText"}], "integrations": {"mapKey": {"mapKey": "anyValue"}}, "modifiers": {"overwrite": true}}, "context": {"integrations": {"mapKey": {"mapKey": "anyValue"}}}, "metadata": {"mapKey": "anyValue"}, "next_step": {"behavior": "get_user_input", "dialog_node": "DialogNode", "selector": "condition"}, "title": "Title", "type": "standard", "event_name": "focus", "variable": "Variable", "actions": [{"name": "Name", "type": "client", "parameters": {"mapKey": "anyValue"}, "result_variable": "ResultVariable", "credentials": "Credentials"}], "digress_in": "not_available", "digress_out": "allow_returning", "digress_out_slots": "not_allowed", "user_label": "UserLabel", "disambiguation_opt_out": false, "disabled": true, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}], "pagination": {"refresh_url": "RefreshURL", "next_url": "NextURL", "total": 5, "matched": 7, "refresh_cursor": "RefreshCursor", "next_cursor": "NextCursor"}}`)
				}))
			})
			It(`Invoke ListDialogNodes successfully with retries`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the ListDialogNodesOptions model
				listDialogNodesOptionsModel := new(assistantv1.ListDialogNodesOptions)
				listDialogNodesOptionsModel.WorkspaceID = core.StringPtr("testString")
				listDialogNodesOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listDialogNodesOptionsModel.IncludeCount = core.BoolPtr(false)
				listDialogNodesOptionsModel.Sort = core.StringPtr("dialog_node")
				listDialogNodesOptionsModel.Cursor = core.StringPtr("testString")
				listDialogNodesOptionsModel.IncludeAudit = core.BoolPtr(false)
				listDialogNodesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.ListDialogNodesWithContext(ctx, listDialogNodesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.ListDialogNodes(listDialogNodesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.ListDialogNodesWithContext(ctx, listDialogNodesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listDialogNodesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["page_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// TODO: Add check for include_count query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"dialog_node"}))
					Expect(req.URL.Query()["cursor"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"dialog_nodes": [{"dialog_node": "DialogNode", "description": "Description", "conditions": "Conditions", "parent": "Parent", "previous_sibling": "PreviousSibling", "output": {"generic": [{"response_type": "video", "source": "Source", "title": "Title", "description": "Description", "channels": [{"channel": "chat"}], "channel_options": {"anyKey": "anyValue"}, "alt_text": "AltText"}], "integrations": {"mapKey": {"mapKey": "anyValue"}}, "modifiers": {"overwrite": true}}, "context": {"integrations": {"mapKey": {"mapKey": "anyValue"}}}, "metadata": {"mapKey": "anyValue"}, "next_step": {"behavior": "get_user_input", "dialog_node": "DialogNode", "selector": "condition"}, "title": "Title", "type": "standard", "event_name": "focus", "variable": "Variable", "actions": [{"name": "Name", "type": "client", "parameters": {"mapKey": "anyValue"}, "result_variable": "ResultVariable", "credentials": "Credentials"}], "digress_in": "not_available", "digress_out": "allow_returning", "digress_out_slots": "not_allowed", "user_label": "UserLabel", "disambiguation_opt_out": false, "disabled": true, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}], "pagination": {"refresh_url": "RefreshURL", "next_url": "NextURL", "total": 5, "matched": 7, "refresh_cursor": "RefreshCursor", "next_cursor": "NextCursor"}}`)
				}))
			})
			It(`Invoke ListDialogNodes successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.ListDialogNodes(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListDialogNodesOptions model
				listDialogNodesOptionsModel := new(assistantv1.ListDialogNodesOptions)
				listDialogNodesOptionsModel.WorkspaceID = core.StringPtr("testString")
				listDialogNodesOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listDialogNodesOptionsModel.IncludeCount = core.BoolPtr(false)
				listDialogNodesOptionsModel.Sort = core.StringPtr("dialog_node")
				listDialogNodesOptionsModel.Cursor = core.StringPtr("testString")
				listDialogNodesOptionsModel.IncludeAudit = core.BoolPtr(false)
				listDialogNodesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.ListDialogNodes(listDialogNodesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListDialogNodes with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ListDialogNodesOptions model
				listDialogNodesOptionsModel := new(assistantv1.ListDialogNodesOptions)
				listDialogNodesOptionsModel.WorkspaceID = core.StringPtr("testString")
				listDialogNodesOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listDialogNodesOptionsModel.IncludeCount = core.BoolPtr(false)
				listDialogNodesOptionsModel.Sort = core.StringPtr("dialog_node")
				listDialogNodesOptionsModel.Cursor = core.StringPtr("testString")
				listDialogNodesOptionsModel.IncludeAudit = core.BoolPtr(false)
				listDialogNodesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.ListDialogNodes(listDialogNodesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListDialogNodesOptions model with no property values
				listDialogNodesOptionsModelNew := new(assistantv1.ListDialogNodesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = assistantService.ListDialogNodes(listDialogNodesOptionsModelNew)
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
			It(`Invoke ListDialogNodes successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ListDialogNodesOptions model
				listDialogNodesOptionsModel := new(assistantv1.ListDialogNodesOptions)
				listDialogNodesOptionsModel.WorkspaceID = core.StringPtr("testString")
				listDialogNodesOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listDialogNodesOptionsModel.IncludeCount = core.BoolPtr(false)
				listDialogNodesOptionsModel.Sort = core.StringPtr("dialog_node")
				listDialogNodesOptionsModel.Cursor = core.StringPtr("testString")
				listDialogNodesOptionsModel.IncludeAudit = core.BoolPtr(false)
				listDialogNodesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.ListDialogNodes(listDialogNodesOptionsModel)
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
	Describe(`CreateDialogNode(createDialogNodeOptions *CreateDialogNodeOptions) - Operation response error`, func() {
		version := "testString"
		createDialogNodePath := "/v1/workspaces/testString/dialog_nodes"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDialogNodePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateDialogNode with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ResponseGenericChannel model
				responseGenericChannelModel := new(assistantv1.ResponseGenericChannel)
				responseGenericChannelModel.Channel = core.StringPtr("chat")

				// Construct an instance of the DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo model
				dialogNodeOutputGenericModel := new(assistantv1.DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo)
				dialogNodeOutputGenericModel.ResponseType = core.StringPtr("video")
				dialogNodeOutputGenericModel.Source = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Title = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Description = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Channels = []assistantv1.ResponseGenericChannel{*responseGenericChannelModel}
				dialogNodeOutputGenericModel.ChannelOptions = map[string]interface{}{"anyKey": "anyValue"}
				dialogNodeOutputGenericModel.AltText = core.StringPtr("testString")

				// Construct an instance of the DialogNodeOutputModifiers model
				dialogNodeOutputModifiersModel := new(assistantv1.DialogNodeOutputModifiers)
				dialogNodeOutputModifiersModel.Overwrite = core.BoolPtr(true)

				// Construct an instance of the DialogNodeOutput model
				dialogNodeOutputModel := new(assistantv1.DialogNodeOutput)
				dialogNodeOutputModel.Generic = []assistantv1.DialogNodeOutputGenericIntf{dialogNodeOutputGenericModel}
				dialogNodeOutputModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeOutputModel.Modifiers = dialogNodeOutputModifiersModel
				dialogNodeOutputModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeContext model
				dialogNodeContextModel := new(assistantv1.DialogNodeContext)
				dialogNodeContextModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeContextModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeNextStep model
				dialogNodeNextStepModel := new(assistantv1.DialogNodeNextStep)
				dialogNodeNextStepModel.Behavior = core.StringPtr("get_user_input")
				dialogNodeNextStepModel.DialogNode = core.StringPtr("testString")
				dialogNodeNextStepModel.Selector = core.StringPtr("condition")

				// Construct an instance of the DialogNodeAction model
				dialogNodeActionModel := new(assistantv1.DialogNodeAction)
				dialogNodeActionModel.Name = core.StringPtr("testString")
				dialogNodeActionModel.Type = core.StringPtr("client")
				dialogNodeActionModel.Parameters = make(map[string]interface{})
				dialogNodeActionModel.ResultVariable = core.StringPtr("testString")
				dialogNodeActionModel.Credentials = core.StringPtr("testString")

				// Construct an instance of the CreateDialogNodeOptions model
				createDialogNodeOptionsModel := new(assistantv1.CreateDialogNodeOptions)
				createDialogNodeOptionsModel.WorkspaceID = core.StringPtr("testString")
				createDialogNodeOptionsModel.DialogNode = core.StringPtr("testString")
				createDialogNodeOptionsModel.Description = core.StringPtr("testString")
				createDialogNodeOptionsModel.Conditions = core.StringPtr("testString")
				createDialogNodeOptionsModel.Parent = core.StringPtr("testString")
				createDialogNodeOptionsModel.PreviousSibling = core.StringPtr("testString")
				createDialogNodeOptionsModel.Output = dialogNodeOutputModel
				createDialogNodeOptionsModel.Context = dialogNodeContextModel
				createDialogNodeOptionsModel.Metadata = make(map[string]interface{})
				createDialogNodeOptionsModel.NextStep = dialogNodeNextStepModel
				createDialogNodeOptionsModel.Title = core.StringPtr("testString")
				createDialogNodeOptionsModel.Type = core.StringPtr("standard")
				createDialogNodeOptionsModel.EventName = core.StringPtr("focus")
				createDialogNodeOptionsModel.Variable = core.StringPtr("testString")
				createDialogNodeOptionsModel.Actions = []assistantv1.DialogNodeAction{*dialogNodeActionModel}
				createDialogNodeOptionsModel.DigressIn = core.StringPtr("not_available")
				createDialogNodeOptionsModel.DigressOut = core.StringPtr("allow_returning")
				createDialogNodeOptionsModel.DigressOutSlots = core.StringPtr("not_allowed")
				createDialogNodeOptionsModel.UserLabel = core.StringPtr("testString")
				createDialogNodeOptionsModel.DisambiguationOptOut = core.BoolPtr(false)
				createDialogNodeOptionsModel.IncludeAudit = core.BoolPtr(false)
				createDialogNodeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.CreateDialogNode(createDialogNodeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.CreateDialogNode(createDialogNodeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateDialogNode(createDialogNodeOptions *CreateDialogNodeOptions)`, func() {
		version := "testString"
		createDialogNodePath := "/v1/workspaces/testString/dialog_nodes"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDialogNodePath))
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
					// TODO: Add check for include_audit query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"dialog_node": "DialogNode", "description": "Description", "conditions": "Conditions", "parent": "Parent", "previous_sibling": "PreviousSibling", "output": {"generic": [{"response_type": "video", "source": "Source", "title": "Title", "description": "Description", "channels": [{"channel": "chat"}], "channel_options": {"anyKey": "anyValue"}, "alt_text": "AltText"}], "integrations": {"mapKey": {"mapKey": "anyValue"}}, "modifiers": {"overwrite": true}}, "context": {"integrations": {"mapKey": {"mapKey": "anyValue"}}}, "metadata": {"mapKey": "anyValue"}, "next_step": {"behavior": "get_user_input", "dialog_node": "DialogNode", "selector": "condition"}, "title": "Title", "type": "standard", "event_name": "focus", "variable": "Variable", "actions": [{"name": "Name", "type": "client", "parameters": {"mapKey": "anyValue"}, "result_variable": "ResultVariable", "credentials": "Credentials"}], "digress_in": "not_available", "digress_out": "allow_returning", "digress_out_slots": "not_allowed", "user_label": "UserLabel", "disambiguation_opt_out": false, "disabled": true, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke CreateDialogNode successfully with retries`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the ResponseGenericChannel model
				responseGenericChannelModel := new(assistantv1.ResponseGenericChannel)
				responseGenericChannelModel.Channel = core.StringPtr("chat")

				// Construct an instance of the DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo model
				dialogNodeOutputGenericModel := new(assistantv1.DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo)
				dialogNodeOutputGenericModel.ResponseType = core.StringPtr("video")
				dialogNodeOutputGenericModel.Source = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Title = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Description = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Channels = []assistantv1.ResponseGenericChannel{*responseGenericChannelModel}
				dialogNodeOutputGenericModel.ChannelOptions = map[string]interface{}{"anyKey": "anyValue"}
				dialogNodeOutputGenericModel.AltText = core.StringPtr("testString")

				// Construct an instance of the DialogNodeOutputModifiers model
				dialogNodeOutputModifiersModel := new(assistantv1.DialogNodeOutputModifiers)
				dialogNodeOutputModifiersModel.Overwrite = core.BoolPtr(true)

				// Construct an instance of the DialogNodeOutput model
				dialogNodeOutputModel := new(assistantv1.DialogNodeOutput)
				dialogNodeOutputModel.Generic = []assistantv1.DialogNodeOutputGenericIntf{dialogNodeOutputGenericModel}
				dialogNodeOutputModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeOutputModel.Modifiers = dialogNodeOutputModifiersModel
				dialogNodeOutputModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeContext model
				dialogNodeContextModel := new(assistantv1.DialogNodeContext)
				dialogNodeContextModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeContextModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeNextStep model
				dialogNodeNextStepModel := new(assistantv1.DialogNodeNextStep)
				dialogNodeNextStepModel.Behavior = core.StringPtr("get_user_input")
				dialogNodeNextStepModel.DialogNode = core.StringPtr("testString")
				dialogNodeNextStepModel.Selector = core.StringPtr("condition")

				// Construct an instance of the DialogNodeAction model
				dialogNodeActionModel := new(assistantv1.DialogNodeAction)
				dialogNodeActionModel.Name = core.StringPtr("testString")
				dialogNodeActionModel.Type = core.StringPtr("client")
				dialogNodeActionModel.Parameters = make(map[string]interface{})
				dialogNodeActionModel.ResultVariable = core.StringPtr("testString")
				dialogNodeActionModel.Credentials = core.StringPtr("testString")

				// Construct an instance of the CreateDialogNodeOptions model
				createDialogNodeOptionsModel := new(assistantv1.CreateDialogNodeOptions)
				createDialogNodeOptionsModel.WorkspaceID = core.StringPtr("testString")
				createDialogNodeOptionsModel.DialogNode = core.StringPtr("testString")
				createDialogNodeOptionsModel.Description = core.StringPtr("testString")
				createDialogNodeOptionsModel.Conditions = core.StringPtr("testString")
				createDialogNodeOptionsModel.Parent = core.StringPtr("testString")
				createDialogNodeOptionsModel.PreviousSibling = core.StringPtr("testString")
				createDialogNodeOptionsModel.Output = dialogNodeOutputModel
				createDialogNodeOptionsModel.Context = dialogNodeContextModel
				createDialogNodeOptionsModel.Metadata = make(map[string]interface{})
				createDialogNodeOptionsModel.NextStep = dialogNodeNextStepModel
				createDialogNodeOptionsModel.Title = core.StringPtr("testString")
				createDialogNodeOptionsModel.Type = core.StringPtr("standard")
				createDialogNodeOptionsModel.EventName = core.StringPtr("focus")
				createDialogNodeOptionsModel.Variable = core.StringPtr("testString")
				createDialogNodeOptionsModel.Actions = []assistantv1.DialogNodeAction{*dialogNodeActionModel}
				createDialogNodeOptionsModel.DigressIn = core.StringPtr("not_available")
				createDialogNodeOptionsModel.DigressOut = core.StringPtr("allow_returning")
				createDialogNodeOptionsModel.DigressOutSlots = core.StringPtr("not_allowed")
				createDialogNodeOptionsModel.UserLabel = core.StringPtr("testString")
				createDialogNodeOptionsModel.DisambiguationOptOut = core.BoolPtr(false)
				createDialogNodeOptionsModel.IncludeAudit = core.BoolPtr(false)
				createDialogNodeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.CreateDialogNodeWithContext(ctx, createDialogNodeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.CreateDialogNode(createDialogNodeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.CreateDialogNodeWithContext(ctx, createDialogNodeOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createDialogNodePath))
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
					// TODO: Add check for include_audit query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"dialog_node": "DialogNode", "description": "Description", "conditions": "Conditions", "parent": "Parent", "previous_sibling": "PreviousSibling", "output": {"generic": [{"response_type": "video", "source": "Source", "title": "Title", "description": "Description", "channels": [{"channel": "chat"}], "channel_options": {"anyKey": "anyValue"}, "alt_text": "AltText"}], "integrations": {"mapKey": {"mapKey": "anyValue"}}, "modifiers": {"overwrite": true}}, "context": {"integrations": {"mapKey": {"mapKey": "anyValue"}}}, "metadata": {"mapKey": "anyValue"}, "next_step": {"behavior": "get_user_input", "dialog_node": "DialogNode", "selector": "condition"}, "title": "Title", "type": "standard", "event_name": "focus", "variable": "Variable", "actions": [{"name": "Name", "type": "client", "parameters": {"mapKey": "anyValue"}, "result_variable": "ResultVariable", "credentials": "Credentials"}], "digress_in": "not_available", "digress_out": "allow_returning", "digress_out_slots": "not_allowed", "user_label": "UserLabel", "disambiguation_opt_out": false, "disabled": true, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke CreateDialogNode successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.CreateDialogNode(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ResponseGenericChannel model
				responseGenericChannelModel := new(assistantv1.ResponseGenericChannel)
				responseGenericChannelModel.Channel = core.StringPtr("chat")

				// Construct an instance of the DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo model
				dialogNodeOutputGenericModel := new(assistantv1.DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo)
				dialogNodeOutputGenericModel.ResponseType = core.StringPtr("video")
				dialogNodeOutputGenericModel.Source = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Title = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Description = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Channels = []assistantv1.ResponseGenericChannel{*responseGenericChannelModel}
				dialogNodeOutputGenericModel.ChannelOptions = map[string]interface{}{"anyKey": "anyValue"}
				dialogNodeOutputGenericModel.AltText = core.StringPtr("testString")

				// Construct an instance of the DialogNodeOutputModifiers model
				dialogNodeOutputModifiersModel := new(assistantv1.DialogNodeOutputModifiers)
				dialogNodeOutputModifiersModel.Overwrite = core.BoolPtr(true)

				// Construct an instance of the DialogNodeOutput model
				dialogNodeOutputModel := new(assistantv1.DialogNodeOutput)
				dialogNodeOutputModel.Generic = []assistantv1.DialogNodeOutputGenericIntf{dialogNodeOutputGenericModel}
				dialogNodeOutputModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeOutputModel.Modifiers = dialogNodeOutputModifiersModel
				dialogNodeOutputModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeContext model
				dialogNodeContextModel := new(assistantv1.DialogNodeContext)
				dialogNodeContextModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeContextModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeNextStep model
				dialogNodeNextStepModel := new(assistantv1.DialogNodeNextStep)
				dialogNodeNextStepModel.Behavior = core.StringPtr("get_user_input")
				dialogNodeNextStepModel.DialogNode = core.StringPtr("testString")
				dialogNodeNextStepModel.Selector = core.StringPtr("condition")

				// Construct an instance of the DialogNodeAction model
				dialogNodeActionModel := new(assistantv1.DialogNodeAction)
				dialogNodeActionModel.Name = core.StringPtr("testString")
				dialogNodeActionModel.Type = core.StringPtr("client")
				dialogNodeActionModel.Parameters = make(map[string]interface{})
				dialogNodeActionModel.ResultVariable = core.StringPtr("testString")
				dialogNodeActionModel.Credentials = core.StringPtr("testString")

				// Construct an instance of the CreateDialogNodeOptions model
				createDialogNodeOptionsModel := new(assistantv1.CreateDialogNodeOptions)
				createDialogNodeOptionsModel.WorkspaceID = core.StringPtr("testString")
				createDialogNodeOptionsModel.DialogNode = core.StringPtr("testString")
				createDialogNodeOptionsModel.Description = core.StringPtr("testString")
				createDialogNodeOptionsModel.Conditions = core.StringPtr("testString")
				createDialogNodeOptionsModel.Parent = core.StringPtr("testString")
				createDialogNodeOptionsModel.PreviousSibling = core.StringPtr("testString")
				createDialogNodeOptionsModel.Output = dialogNodeOutputModel
				createDialogNodeOptionsModel.Context = dialogNodeContextModel
				createDialogNodeOptionsModel.Metadata = make(map[string]interface{})
				createDialogNodeOptionsModel.NextStep = dialogNodeNextStepModel
				createDialogNodeOptionsModel.Title = core.StringPtr("testString")
				createDialogNodeOptionsModel.Type = core.StringPtr("standard")
				createDialogNodeOptionsModel.EventName = core.StringPtr("focus")
				createDialogNodeOptionsModel.Variable = core.StringPtr("testString")
				createDialogNodeOptionsModel.Actions = []assistantv1.DialogNodeAction{*dialogNodeActionModel}
				createDialogNodeOptionsModel.DigressIn = core.StringPtr("not_available")
				createDialogNodeOptionsModel.DigressOut = core.StringPtr("allow_returning")
				createDialogNodeOptionsModel.DigressOutSlots = core.StringPtr("not_allowed")
				createDialogNodeOptionsModel.UserLabel = core.StringPtr("testString")
				createDialogNodeOptionsModel.DisambiguationOptOut = core.BoolPtr(false)
				createDialogNodeOptionsModel.IncludeAudit = core.BoolPtr(false)
				createDialogNodeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.CreateDialogNode(createDialogNodeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateDialogNode with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ResponseGenericChannel model
				responseGenericChannelModel := new(assistantv1.ResponseGenericChannel)
				responseGenericChannelModel.Channel = core.StringPtr("chat")

				// Construct an instance of the DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo model
				dialogNodeOutputGenericModel := new(assistantv1.DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo)
				dialogNodeOutputGenericModel.ResponseType = core.StringPtr("video")
				dialogNodeOutputGenericModel.Source = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Title = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Description = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Channels = []assistantv1.ResponseGenericChannel{*responseGenericChannelModel}
				dialogNodeOutputGenericModel.ChannelOptions = map[string]interface{}{"anyKey": "anyValue"}
				dialogNodeOutputGenericModel.AltText = core.StringPtr("testString")

				// Construct an instance of the DialogNodeOutputModifiers model
				dialogNodeOutputModifiersModel := new(assistantv1.DialogNodeOutputModifiers)
				dialogNodeOutputModifiersModel.Overwrite = core.BoolPtr(true)

				// Construct an instance of the DialogNodeOutput model
				dialogNodeOutputModel := new(assistantv1.DialogNodeOutput)
				dialogNodeOutputModel.Generic = []assistantv1.DialogNodeOutputGenericIntf{dialogNodeOutputGenericModel}
				dialogNodeOutputModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeOutputModel.Modifiers = dialogNodeOutputModifiersModel
				dialogNodeOutputModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeContext model
				dialogNodeContextModel := new(assistantv1.DialogNodeContext)
				dialogNodeContextModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeContextModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeNextStep model
				dialogNodeNextStepModel := new(assistantv1.DialogNodeNextStep)
				dialogNodeNextStepModel.Behavior = core.StringPtr("get_user_input")
				dialogNodeNextStepModel.DialogNode = core.StringPtr("testString")
				dialogNodeNextStepModel.Selector = core.StringPtr("condition")

				// Construct an instance of the DialogNodeAction model
				dialogNodeActionModel := new(assistantv1.DialogNodeAction)
				dialogNodeActionModel.Name = core.StringPtr("testString")
				dialogNodeActionModel.Type = core.StringPtr("client")
				dialogNodeActionModel.Parameters = make(map[string]interface{})
				dialogNodeActionModel.ResultVariable = core.StringPtr("testString")
				dialogNodeActionModel.Credentials = core.StringPtr("testString")

				// Construct an instance of the CreateDialogNodeOptions model
				createDialogNodeOptionsModel := new(assistantv1.CreateDialogNodeOptions)
				createDialogNodeOptionsModel.WorkspaceID = core.StringPtr("testString")
				createDialogNodeOptionsModel.DialogNode = core.StringPtr("testString")
				createDialogNodeOptionsModel.Description = core.StringPtr("testString")
				createDialogNodeOptionsModel.Conditions = core.StringPtr("testString")
				createDialogNodeOptionsModel.Parent = core.StringPtr("testString")
				createDialogNodeOptionsModel.PreviousSibling = core.StringPtr("testString")
				createDialogNodeOptionsModel.Output = dialogNodeOutputModel
				createDialogNodeOptionsModel.Context = dialogNodeContextModel
				createDialogNodeOptionsModel.Metadata = make(map[string]interface{})
				createDialogNodeOptionsModel.NextStep = dialogNodeNextStepModel
				createDialogNodeOptionsModel.Title = core.StringPtr("testString")
				createDialogNodeOptionsModel.Type = core.StringPtr("standard")
				createDialogNodeOptionsModel.EventName = core.StringPtr("focus")
				createDialogNodeOptionsModel.Variable = core.StringPtr("testString")
				createDialogNodeOptionsModel.Actions = []assistantv1.DialogNodeAction{*dialogNodeActionModel}
				createDialogNodeOptionsModel.DigressIn = core.StringPtr("not_available")
				createDialogNodeOptionsModel.DigressOut = core.StringPtr("allow_returning")
				createDialogNodeOptionsModel.DigressOutSlots = core.StringPtr("not_allowed")
				createDialogNodeOptionsModel.UserLabel = core.StringPtr("testString")
				createDialogNodeOptionsModel.DisambiguationOptOut = core.BoolPtr(false)
				createDialogNodeOptionsModel.IncludeAudit = core.BoolPtr(false)
				createDialogNodeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.CreateDialogNode(createDialogNodeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateDialogNodeOptions model with no property values
				createDialogNodeOptionsModelNew := new(assistantv1.CreateDialogNodeOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = assistantService.CreateDialogNode(createDialogNodeOptionsModelNew)
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
			It(`Invoke CreateDialogNode successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ResponseGenericChannel model
				responseGenericChannelModel := new(assistantv1.ResponseGenericChannel)
				responseGenericChannelModel.Channel = core.StringPtr("chat")

				// Construct an instance of the DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo model
				dialogNodeOutputGenericModel := new(assistantv1.DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo)
				dialogNodeOutputGenericModel.ResponseType = core.StringPtr("video")
				dialogNodeOutputGenericModel.Source = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Title = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Description = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Channels = []assistantv1.ResponseGenericChannel{*responseGenericChannelModel}
				dialogNodeOutputGenericModel.ChannelOptions = map[string]interface{}{"anyKey": "anyValue"}
				dialogNodeOutputGenericModel.AltText = core.StringPtr("testString")

				// Construct an instance of the DialogNodeOutputModifiers model
				dialogNodeOutputModifiersModel := new(assistantv1.DialogNodeOutputModifiers)
				dialogNodeOutputModifiersModel.Overwrite = core.BoolPtr(true)

				// Construct an instance of the DialogNodeOutput model
				dialogNodeOutputModel := new(assistantv1.DialogNodeOutput)
				dialogNodeOutputModel.Generic = []assistantv1.DialogNodeOutputGenericIntf{dialogNodeOutputGenericModel}
				dialogNodeOutputModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeOutputModel.Modifiers = dialogNodeOutputModifiersModel
				dialogNodeOutputModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeContext model
				dialogNodeContextModel := new(assistantv1.DialogNodeContext)
				dialogNodeContextModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeContextModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeNextStep model
				dialogNodeNextStepModel := new(assistantv1.DialogNodeNextStep)
				dialogNodeNextStepModel.Behavior = core.StringPtr("get_user_input")
				dialogNodeNextStepModel.DialogNode = core.StringPtr("testString")
				dialogNodeNextStepModel.Selector = core.StringPtr("condition")

				// Construct an instance of the DialogNodeAction model
				dialogNodeActionModel := new(assistantv1.DialogNodeAction)
				dialogNodeActionModel.Name = core.StringPtr("testString")
				dialogNodeActionModel.Type = core.StringPtr("client")
				dialogNodeActionModel.Parameters = make(map[string]interface{})
				dialogNodeActionModel.ResultVariable = core.StringPtr("testString")
				dialogNodeActionModel.Credentials = core.StringPtr("testString")

				// Construct an instance of the CreateDialogNodeOptions model
				createDialogNodeOptionsModel := new(assistantv1.CreateDialogNodeOptions)
				createDialogNodeOptionsModel.WorkspaceID = core.StringPtr("testString")
				createDialogNodeOptionsModel.DialogNode = core.StringPtr("testString")
				createDialogNodeOptionsModel.Description = core.StringPtr("testString")
				createDialogNodeOptionsModel.Conditions = core.StringPtr("testString")
				createDialogNodeOptionsModel.Parent = core.StringPtr("testString")
				createDialogNodeOptionsModel.PreviousSibling = core.StringPtr("testString")
				createDialogNodeOptionsModel.Output = dialogNodeOutputModel
				createDialogNodeOptionsModel.Context = dialogNodeContextModel
				createDialogNodeOptionsModel.Metadata = make(map[string]interface{})
				createDialogNodeOptionsModel.NextStep = dialogNodeNextStepModel
				createDialogNodeOptionsModel.Title = core.StringPtr("testString")
				createDialogNodeOptionsModel.Type = core.StringPtr("standard")
				createDialogNodeOptionsModel.EventName = core.StringPtr("focus")
				createDialogNodeOptionsModel.Variable = core.StringPtr("testString")
				createDialogNodeOptionsModel.Actions = []assistantv1.DialogNodeAction{*dialogNodeActionModel}
				createDialogNodeOptionsModel.DigressIn = core.StringPtr("not_available")
				createDialogNodeOptionsModel.DigressOut = core.StringPtr("allow_returning")
				createDialogNodeOptionsModel.DigressOutSlots = core.StringPtr("not_allowed")
				createDialogNodeOptionsModel.UserLabel = core.StringPtr("testString")
				createDialogNodeOptionsModel.DisambiguationOptOut = core.BoolPtr(false)
				createDialogNodeOptionsModel.IncludeAudit = core.BoolPtr(false)
				createDialogNodeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.CreateDialogNode(createDialogNodeOptionsModel)
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
	Describe(`GetDialogNode(getDialogNodeOptions *GetDialogNodeOptions) - Operation response error`, func() {
		version := "testString"
		getDialogNodePath := "/v1/workspaces/testString/dialog_nodes/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDialogNodePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDialogNode with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the GetDialogNodeOptions model
				getDialogNodeOptionsModel := new(assistantv1.GetDialogNodeOptions)
				getDialogNodeOptionsModel.WorkspaceID = core.StringPtr("testString")
				getDialogNodeOptionsModel.DialogNode = core.StringPtr("testString")
				getDialogNodeOptionsModel.IncludeAudit = core.BoolPtr(false)
				getDialogNodeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.GetDialogNode(getDialogNodeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.GetDialogNode(getDialogNodeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDialogNode(getDialogNodeOptions *GetDialogNodeOptions)`, func() {
		version := "testString"
		getDialogNodePath := "/v1/workspaces/testString/dialog_nodes/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDialogNodePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"dialog_node": "DialogNode", "description": "Description", "conditions": "Conditions", "parent": "Parent", "previous_sibling": "PreviousSibling", "output": {"generic": [{"response_type": "video", "source": "Source", "title": "Title", "description": "Description", "channels": [{"channel": "chat"}], "channel_options": {"anyKey": "anyValue"}, "alt_text": "AltText"}], "integrations": {"mapKey": {"mapKey": "anyValue"}}, "modifiers": {"overwrite": true}}, "context": {"integrations": {"mapKey": {"mapKey": "anyValue"}}}, "metadata": {"mapKey": "anyValue"}, "next_step": {"behavior": "get_user_input", "dialog_node": "DialogNode", "selector": "condition"}, "title": "Title", "type": "standard", "event_name": "focus", "variable": "Variable", "actions": [{"name": "Name", "type": "client", "parameters": {"mapKey": "anyValue"}, "result_variable": "ResultVariable", "credentials": "Credentials"}], "digress_in": "not_available", "digress_out": "allow_returning", "digress_out_slots": "not_allowed", "user_label": "UserLabel", "disambiguation_opt_out": false, "disabled": true, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetDialogNode successfully with retries`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the GetDialogNodeOptions model
				getDialogNodeOptionsModel := new(assistantv1.GetDialogNodeOptions)
				getDialogNodeOptionsModel.WorkspaceID = core.StringPtr("testString")
				getDialogNodeOptionsModel.DialogNode = core.StringPtr("testString")
				getDialogNodeOptionsModel.IncludeAudit = core.BoolPtr(false)
				getDialogNodeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.GetDialogNodeWithContext(ctx, getDialogNodeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.GetDialogNode(getDialogNodeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.GetDialogNodeWithContext(ctx, getDialogNodeOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getDialogNodePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"dialog_node": "DialogNode", "description": "Description", "conditions": "Conditions", "parent": "Parent", "previous_sibling": "PreviousSibling", "output": {"generic": [{"response_type": "video", "source": "Source", "title": "Title", "description": "Description", "channels": [{"channel": "chat"}], "channel_options": {"anyKey": "anyValue"}, "alt_text": "AltText"}], "integrations": {"mapKey": {"mapKey": "anyValue"}}, "modifiers": {"overwrite": true}}, "context": {"integrations": {"mapKey": {"mapKey": "anyValue"}}}, "metadata": {"mapKey": "anyValue"}, "next_step": {"behavior": "get_user_input", "dialog_node": "DialogNode", "selector": "condition"}, "title": "Title", "type": "standard", "event_name": "focus", "variable": "Variable", "actions": [{"name": "Name", "type": "client", "parameters": {"mapKey": "anyValue"}, "result_variable": "ResultVariable", "credentials": "Credentials"}], "digress_in": "not_available", "digress_out": "allow_returning", "digress_out_slots": "not_allowed", "user_label": "UserLabel", "disambiguation_opt_out": false, "disabled": true, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetDialogNode successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.GetDialogNode(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDialogNodeOptions model
				getDialogNodeOptionsModel := new(assistantv1.GetDialogNodeOptions)
				getDialogNodeOptionsModel.WorkspaceID = core.StringPtr("testString")
				getDialogNodeOptionsModel.DialogNode = core.StringPtr("testString")
				getDialogNodeOptionsModel.IncludeAudit = core.BoolPtr(false)
				getDialogNodeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.GetDialogNode(getDialogNodeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDialogNode with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the GetDialogNodeOptions model
				getDialogNodeOptionsModel := new(assistantv1.GetDialogNodeOptions)
				getDialogNodeOptionsModel.WorkspaceID = core.StringPtr("testString")
				getDialogNodeOptionsModel.DialogNode = core.StringPtr("testString")
				getDialogNodeOptionsModel.IncludeAudit = core.BoolPtr(false)
				getDialogNodeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.GetDialogNode(getDialogNodeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDialogNodeOptions model with no property values
				getDialogNodeOptionsModelNew := new(assistantv1.GetDialogNodeOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = assistantService.GetDialogNode(getDialogNodeOptionsModelNew)
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
			It(`Invoke GetDialogNode successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the GetDialogNodeOptions model
				getDialogNodeOptionsModel := new(assistantv1.GetDialogNodeOptions)
				getDialogNodeOptionsModel.WorkspaceID = core.StringPtr("testString")
				getDialogNodeOptionsModel.DialogNode = core.StringPtr("testString")
				getDialogNodeOptionsModel.IncludeAudit = core.BoolPtr(false)
				getDialogNodeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.GetDialogNode(getDialogNodeOptionsModel)
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
	Describe(`UpdateDialogNode(updateDialogNodeOptions *UpdateDialogNodeOptions) - Operation response error`, func() {
		version := "testString"
		updateDialogNodePath := "/v1/workspaces/testString/dialog_nodes/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDialogNodePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_audit query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateDialogNode with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ResponseGenericChannel model
				responseGenericChannelModel := new(assistantv1.ResponseGenericChannel)
				responseGenericChannelModel.Channel = core.StringPtr("chat")

				// Construct an instance of the DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo model
				dialogNodeOutputGenericModel := new(assistantv1.DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo)
				dialogNodeOutputGenericModel.ResponseType = core.StringPtr("video")
				dialogNodeOutputGenericModel.Source = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Title = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Description = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Channels = []assistantv1.ResponseGenericChannel{*responseGenericChannelModel}
				dialogNodeOutputGenericModel.ChannelOptions = map[string]interface{}{"anyKey": "anyValue"}
				dialogNodeOutputGenericModel.AltText = core.StringPtr("testString")

				// Construct an instance of the DialogNodeOutputModifiers model
				dialogNodeOutputModifiersModel := new(assistantv1.DialogNodeOutputModifiers)
				dialogNodeOutputModifiersModel.Overwrite = core.BoolPtr(true)

				// Construct an instance of the DialogNodeOutput model
				dialogNodeOutputModel := new(assistantv1.DialogNodeOutput)
				dialogNodeOutputModel.Generic = []assistantv1.DialogNodeOutputGenericIntf{dialogNodeOutputGenericModel}
				dialogNodeOutputModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeOutputModel.Modifiers = dialogNodeOutputModifiersModel
				dialogNodeOutputModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeContext model
				dialogNodeContextModel := new(assistantv1.DialogNodeContext)
				dialogNodeContextModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeContextModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeNextStep model
				dialogNodeNextStepModel := new(assistantv1.DialogNodeNextStep)
				dialogNodeNextStepModel.Behavior = core.StringPtr("get_user_input")
				dialogNodeNextStepModel.DialogNode = core.StringPtr("testString")
				dialogNodeNextStepModel.Selector = core.StringPtr("condition")

				// Construct an instance of the DialogNodeAction model
				dialogNodeActionModel := new(assistantv1.DialogNodeAction)
				dialogNodeActionModel.Name = core.StringPtr("testString")
				dialogNodeActionModel.Type = core.StringPtr("client")
				dialogNodeActionModel.Parameters = make(map[string]interface{})
				dialogNodeActionModel.ResultVariable = core.StringPtr("testString")
				dialogNodeActionModel.Credentials = core.StringPtr("testString")

				// Construct an instance of the UpdateDialogNodeOptions model
				updateDialogNodeOptionsModel := new(assistantv1.UpdateDialogNodeOptions)
				updateDialogNodeOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateDialogNodeOptionsModel.DialogNode = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewDialogNode = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewDescription = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewConditions = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewParent = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewPreviousSibling = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewOutput = dialogNodeOutputModel
				updateDialogNodeOptionsModel.NewContext = dialogNodeContextModel
				updateDialogNodeOptionsModel.NewMetadata = make(map[string]interface{})
				updateDialogNodeOptionsModel.NewNextStep = dialogNodeNextStepModel
				updateDialogNodeOptionsModel.NewTitle = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewType = core.StringPtr("standard")
				updateDialogNodeOptionsModel.NewEventName = core.StringPtr("focus")
				updateDialogNodeOptionsModel.NewVariable = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewActions = []assistantv1.DialogNodeAction{*dialogNodeActionModel}
				updateDialogNodeOptionsModel.NewDigressIn = core.StringPtr("not_available")
				updateDialogNodeOptionsModel.NewDigressOut = core.StringPtr("allow_returning")
				updateDialogNodeOptionsModel.NewDigressOutSlots = core.StringPtr("not_allowed")
				updateDialogNodeOptionsModel.NewUserLabel = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewDisambiguationOptOut = core.BoolPtr(false)
				updateDialogNodeOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateDialogNodeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.UpdateDialogNode(updateDialogNodeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.UpdateDialogNode(updateDialogNodeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateDialogNode(updateDialogNodeOptions *UpdateDialogNodeOptions)`, func() {
		version := "testString"
		updateDialogNodePath := "/v1/workspaces/testString/dialog_nodes/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDialogNodePath))
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
					// TODO: Add check for include_audit query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"dialog_node": "DialogNode", "description": "Description", "conditions": "Conditions", "parent": "Parent", "previous_sibling": "PreviousSibling", "output": {"generic": [{"response_type": "video", "source": "Source", "title": "Title", "description": "Description", "channels": [{"channel": "chat"}], "channel_options": {"anyKey": "anyValue"}, "alt_text": "AltText"}], "integrations": {"mapKey": {"mapKey": "anyValue"}}, "modifiers": {"overwrite": true}}, "context": {"integrations": {"mapKey": {"mapKey": "anyValue"}}}, "metadata": {"mapKey": "anyValue"}, "next_step": {"behavior": "get_user_input", "dialog_node": "DialogNode", "selector": "condition"}, "title": "Title", "type": "standard", "event_name": "focus", "variable": "Variable", "actions": [{"name": "Name", "type": "client", "parameters": {"mapKey": "anyValue"}, "result_variable": "ResultVariable", "credentials": "Credentials"}], "digress_in": "not_available", "digress_out": "allow_returning", "digress_out_slots": "not_allowed", "user_label": "UserLabel", "disambiguation_opt_out": false, "disabled": true, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke UpdateDialogNode successfully with retries`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the ResponseGenericChannel model
				responseGenericChannelModel := new(assistantv1.ResponseGenericChannel)
				responseGenericChannelModel.Channel = core.StringPtr("chat")

				// Construct an instance of the DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo model
				dialogNodeOutputGenericModel := new(assistantv1.DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo)
				dialogNodeOutputGenericModel.ResponseType = core.StringPtr("video")
				dialogNodeOutputGenericModel.Source = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Title = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Description = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Channels = []assistantv1.ResponseGenericChannel{*responseGenericChannelModel}
				dialogNodeOutputGenericModel.ChannelOptions = map[string]interface{}{"anyKey": "anyValue"}
				dialogNodeOutputGenericModel.AltText = core.StringPtr("testString")

				// Construct an instance of the DialogNodeOutputModifiers model
				dialogNodeOutputModifiersModel := new(assistantv1.DialogNodeOutputModifiers)
				dialogNodeOutputModifiersModel.Overwrite = core.BoolPtr(true)

				// Construct an instance of the DialogNodeOutput model
				dialogNodeOutputModel := new(assistantv1.DialogNodeOutput)
				dialogNodeOutputModel.Generic = []assistantv1.DialogNodeOutputGenericIntf{dialogNodeOutputGenericModel}
				dialogNodeOutputModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeOutputModel.Modifiers = dialogNodeOutputModifiersModel
				dialogNodeOutputModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeContext model
				dialogNodeContextModel := new(assistantv1.DialogNodeContext)
				dialogNodeContextModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeContextModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeNextStep model
				dialogNodeNextStepModel := new(assistantv1.DialogNodeNextStep)
				dialogNodeNextStepModel.Behavior = core.StringPtr("get_user_input")
				dialogNodeNextStepModel.DialogNode = core.StringPtr("testString")
				dialogNodeNextStepModel.Selector = core.StringPtr("condition")

				// Construct an instance of the DialogNodeAction model
				dialogNodeActionModel := new(assistantv1.DialogNodeAction)
				dialogNodeActionModel.Name = core.StringPtr("testString")
				dialogNodeActionModel.Type = core.StringPtr("client")
				dialogNodeActionModel.Parameters = make(map[string]interface{})
				dialogNodeActionModel.ResultVariable = core.StringPtr("testString")
				dialogNodeActionModel.Credentials = core.StringPtr("testString")

				// Construct an instance of the UpdateDialogNodeOptions model
				updateDialogNodeOptionsModel := new(assistantv1.UpdateDialogNodeOptions)
				updateDialogNodeOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateDialogNodeOptionsModel.DialogNode = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewDialogNode = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewDescription = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewConditions = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewParent = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewPreviousSibling = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewOutput = dialogNodeOutputModel
				updateDialogNodeOptionsModel.NewContext = dialogNodeContextModel
				updateDialogNodeOptionsModel.NewMetadata = make(map[string]interface{})
				updateDialogNodeOptionsModel.NewNextStep = dialogNodeNextStepModel
				updateDialogNodeOptionsModel.NewTitle = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewType = core.StringPtr("standard")
				updateDialogNodeOptionsModel.NewEventName = core.StringPtr("focus")
				updateDialogNodeOptionsModel.NewVariable = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewActions = []assistantv1.DialogNodeAction{*dialogNodeActionModel}
				updateDialogNodeOptionsModel.NewDigressIn = core.StringPtr("not_available")
				updateDialogNodeOptionsModel.NewDigressOut = core.StringPtr("allow_returning")
				updateDialogNodeOptionsModel.NewDigressOutSlots = core.StringPtr("not_allowed")
				updateDialogNodeOptionsModel.NewUserLabel = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewDisambiguationOptOut = core.BoolPtr(false)
				updateDialogNodeOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateDialogNodeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.UpdateDialogNodeWithContext(ctx, updateDialogNodeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.UpdateDialogNode(updateDialogNodeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.UpdateDialogNodeWithContext(ctx, updateDialogNodeOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateDialogNodePath))
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
					// TODO: Add check for include_audit query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"dialog_node": "DialogNode", "description": "Description", "conditions": "Conditions", "parent": "Parent", "previous_sibling": "PreviousSibling", "output": {"generic": [{"response_type": "video", "source": "Source", "title": "Title", "description": "Description", "channels": [{"channel": "chat"}], "channel_options": {"anyKey": "anyValue"}, "alt_text": "AltText"}], "integrations": {"mapKey": {"mapKey": "anyValue"}}, "modifiers": {"overwrite": true}}, "context": {"integrations": {"mapKey": {"mapKey": "anyValue"}}}, "metadata": {"mapKey": "anyValue"}, "next_step": {"behavior": "get_user_input", "dialog_node": "DialogNode", "selector": "condition"}, "title": "Title", "type": "standard", "event_name": "focus", "variable": "Variable", "actions": [{"name": "Name", "type": "client", "parameters": {"mapKey": "anyValue"}, "result_variable": "ResultVariable", "credentials": "Credentials"}], "digress_in": "not_available", "digress_out": "allow_returning", "digress_out_slots": "not_allowed", "user_label": "UserLabel", "disambiguation_opt_out": false, "disabled": true, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke UpdateDialogNode successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.UpdateDialogNode(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ResponseGenericChannel model
				responseGenericChannelModel := new(assistantv1.ResponseGenericChannel)
				responseGenericChannelModel.Channel = core.StringPtr("chat")

				// Construct an instance of the DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo model
				dialogNodeOutputGenericModel := new(assistantv1.DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo)
				dialogNodeOutputGenericModel.ResponseType = core.StringPtr("video")
				dialogNodeOutputGenericModel.Source = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Title = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Description = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Channels = []assistantv1.ResponseGenericChannel{*responseGenericChannelModel}
				dialogNodeOutputGenericModel.ChannelOptions = map[string]interface{}{"anyKey": "anyValue"}
				dialogNodeOutputGenericModel.AltText = core.StringPtr("testString")

				// Construct an instance of the DialogNodeOutputModifiers model
				dialogNodeOutputModifiersModel := new(assistantv1.DialogNodeOutputModifiers)
				dialogNodeOutputModifiersModel.Overwrite = core.BoolPtr(true)

				// Construct an instance of the DialogNodeOutput model
				dialogNodeOutputModel := new(assistantv1.DialogNodeOutput)
				dialogNodeOutputModel.Generic = []assistantv1.DialogNodeOutputGenericIntf{dialogNodeOutputGenericModel}
				dialogNodeOutputModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeOutputModel.Modifiers = dialogNodeOutputModifiersModel
				dialogNodeOutputModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeContext model
				dialogNodeContextModel := new(assistantv1.DialogNodeContext)
				dialogNodeContextModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeContextModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeNextStep model
				dialogNodeNextStepModel := new(assistantv1.DialogNodeNextStep)
				dialogNodeNextStepModel.Behavior = core.StringPtr("get_user_input")
				dialogNodeNextStepModel.DialogNode = core.StringPtr("testString")
				dialogNodeNextStepModel.Selector = core.StringPtr("condition")

				// Construct an instance of the DialogNodeAction model
				dialogNodeActionModel := new(assistantv1.DialogNodeAction)
				dialogNodeActionModel.Name = core.StringPtr("testString")
				dialogNodeActionModel.Type = core.StringPtr("client")
				dialogNodeActionModel.Parameters = make(map[string]interface{})
				dialogNodeActionModel.ResultVariable = core.StringPtr("testString")
				dialogNodeActionModel.Credentials = core.StringPtr("testString")

				// Construct an instance of the UpdateDialogNodeOptions model
				updateDialogNodeOptionsModel := new(assistantv1.UpdateDialogNodeOptions)
				updateDialogNodeOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateDialogNodeOptionsModel.DialogNode = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewDialogNode = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewDescription = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewConditions = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewParent = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewPreviousSibling = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewOutput = dialogNodeOutputModel
				updateDialogNodeOptionsModel.NewContext = dialogNodeContextModel
				updateDialogNodeOptionsModel.NewMetadata = make(map[string]interface{})
				updateDialogNodeOptionsModel.NewNextStep = dialogNodeNextStepModel
				updateDialogNodeOptionsModel.NewTitle = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewType = core.StringPtr("standard")
				updateDialogNodeOptionsModel.NewEventName = core.StringPtr("focus")
				updateDialogNodeOptionsModel.NewVariable = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewActions = []assistantv1.DialogNodeAction{*dialogNodeActionModel}
				updateDialogNodeOptionsModel.NewDigressIn = core.StringPtr("not_available")
				updateDialogNodeOptionsModel.NewDigressOut = core.StringPtr("allow_returning")
				updateDialogNodeOptionsModel.NewDigressOutSlots = core.StringPtr("not_allowed")
				updateDialogNodeOptionsModel.NewUserLabel = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewDisambiguationOptOut = core.BoolPtr(false)
				updateDialogNodeOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateDialogNodeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.UpdateDialogNode(updateDialogNodeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateDialogNode with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ResponseGenericChannel model
				responseGenericChannelModel := new(assistantv1.ResponseGenericChannel)
				responseGenericChannelModel.Channel = core.StringPtr("chat")

				// Construct an instance of the DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo model
				dialogNodeOutputGenericModel := new(assistantv1.DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo)
				dialogNodeOutputGenericModel.ResponseType = core.StringPtr("video")
				dialogNodeOutputGenericModel.Source = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Title = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Description = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Channels = []assistantv1.ResponseGenericChannel{*responseGenericChannelModel}
				dialogNodeOutputGenericModel.ChannelOptions = map[string]interface{}{"anyKey": "anyValue"}
				dialogNodeOutputGenericModel.AltText = core.StringPtr("testString")

				// Construct an instance of the DialogNodeOutputModifiers model
				dialogNodeOutputModifiersModel := new(assistantv1.DialogNodeOutputModifiers)
				dialogNodeOutputModifiersModel.Overwrite = core.BoolPtr(true)

				// Construct an instance of the DialogNodeOutput model
				dialogNodeOutputModel := new(assistantv1.DialogNodeOutput)
				dialogNodeOutputModel.Generic = []assistantv1.DialogNodeOutputGenericIntf{dialogNodeOutputGenericModel}
				dialogNodeOutputModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeOutputModel.Modifiers = dialogNodeOutputModifiersModel
				dialogNodeOutputModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeContext model
				dialogNodeContextModel := new(assistantv1.DialogNodeContext)
				dialogNodeContextModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeContextModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeNextStep model
				dialogNodeNextStepModel := new(assistantv1.DialogNodeNextStep)
				dialogNodeNextStepModel.Behavior = core.StringPtr("get_user_input")
				dialogNodeNextStepModel.DialogNode = core.StringPtr("testString")
				dialogNodeNextStepModel.Selector = core.StringPtr("condition")

				// Construct an instance of the DialogNodeAction model
				dialogNodeActionModel := new(assistantv1.DialogNodeAction)
				dialogNodeActionModel.Name = core.StringPtr("testString")
				dialogNodeActionModel.Type = core.StringPtr("client")
				dialogNodeActionModel.Parameters = make(map[string]interface{})
				dialogNodeActionModel.ResultVariable = core.StringPtr("testString")
				dialogNodeActionModel.Credentials = core.StringPtr("testString")

				// Construct an instance of the UpdateDialogNodeOptions model
				updateDialogNodeOptionsModel := new(assistantv1.UpdateDialogNodeOptions)
				updateDialogNodeOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateDialogNodeOptionsModel.DialogNode = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewDialogNode = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewDescription = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewConditions = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewParent = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewPreviousSibling = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewOutput = dialogNodeOutputModel
				updateDialogNodeOptionsModel.NewContext = dialogNodeContextModel
				updateDialogNodeOptionsModel.NewMetadata = make(map[string]interface{})
				updateDialogNodeOptionsModel.NewNextStep = dialogNodeNextStepModel
				updateDialogNodeOptionsModel.NewTitle = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewType = core.StringPtr("standard")
				updateDialogNodeOptionsModel.NewEventName = core.StringPtr("focus")
				updateDialogNodeOptionsModel.NewVariable = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewActions = []assistantv1.DialogNodeAction{*dialogNodeActionModel}
				updateDialogNodeOptionsModel.NewDigressIn = core.StringPtr("not_available")
				updateDialogNodeOptionsModel.NewDigressOut = core.StringPtr("allow_returning")
				updateDialogNodeOptionsModel.NewDigressOutSlots = core.StringPtr("not_allowed")
				updateDialogNodeOptionsModel.NewUserLabel = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewDisambiguationOptOut = core.BoolPtr(false)
				updateDialogNodeOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateDialogNodeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.UpdateDialogNode(updateDialogNodeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateDialogNodeOptions model with no property values
				updateDialogNodeOptionsModelNew := new(assistantv1.UpdateDialogNodeOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = assistantService.UpdateDialogNode(updateDialogNodeOptionsModelNew)
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
			It(`Invoke UpdateDialogNode successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ResponseGenericChannel model
				responseGenericChannelModel := new(assistantv1.ResponseGenericChannel)
				responseGenericChannelModel.Channel = core.StringPtr("chat")

				// Construct an instance of the DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo model
				dialogNodeOutputGenericModel := new(assistantv1.DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo)
				dialogNodeOutputGenericModel.ResponseType = core.StringPtr("video")
				dialogNodeOutputGenericModel.Source = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Title = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Description = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Channels = []assistantv1.ResponseGenericChannel{*responseGenericChannelModel}
				dialogNodeOutputGenericModel.ChannelOptions = map[string]interface{}{"anyKey": "anyValue"}
				dialogNodeOutputGenericModel.AltText = core.StringPtr("testString")

				// Construct an instance of the DialogNodeOutputModifiers model
				dialogNodeOutputModifiersModel := new(assistantv1.DialogNodeOutputModifiers)
				dialogNodeOutputModifiersModel.Overwrite = core.BoolPtr(true)

				// Construct an instance of the DialogNodeOutput model
				dialogNodeOutputModel := new(assistantv1.DialogNodeOutput)
				dialogNodeOutputModel.Generic = []assistantv1.DialogNodeOutputGenericIntf{dialogNodeOutputGenericModel}
				dialogNodeOutputModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeOutputModel.Modifiers = dialogNodeOutputModifiersModel
				dialogNodeOutputModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeContext model
				dialogNodeContextModel := new(assistantv1.DialogNodeContext)
				dialogNodeContextModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeContextModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the DialogNodeNextStep model
				dialogNodeNextStepModel := new(assistantv1.DialogNodeNextStep)
				dialogNodeNextStepModel.Behavior = core.StringPtr("get_user_input")
				dialogNodeNextStepModel.DialogNode = core.StringPtr("testString")
				dialogNodeNextStepModel.Selector = core.StringPtr("condition")

				// Construct an instance of the DialogNodeAction model
				dialogNodeActionModel := new(assistantv1.DialogNodeAction)
				dialogNodeActionModel.Name = core.StringPtr("testString")
				dialogNodeActionModel.Type = core.StringPtr("client")
				dialogNodeActionModel.Parameters = make(map[string]interface{})
				dialogNodeActionModel.ResultVariable = core.StringPtr("testString")
				dialogNodeActionModel.Credentials = core.StringPtr("testString")

				// Construct an instance of the UpdateDialogNodeOptions model
				updateDialogNodeOptionsModel := new(assistantv1.UpdateDialogNodeOptions)
				updateDialogNodeOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateDialogNodeOptionsModel.DialogNode = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewDialogNode = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewDescription = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewConditions = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewParent = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewPreviousSibling = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewOutput = dialogNodeOutputModel
				updateDialogNodeOptionsModel.NewContext = dialogNodeContextModel
				updateDialogNodeOptionsModel.NewMetadata = make(map[string]interface{})
				updateDialogNodeOptionsModel.NewNextStep = dialogNodeNextStepModel
				updateDialogNodeOptionsModel.NewTitle = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewType = core.StringPtr("standard")
				updateDialogNodeOptionsModel.NewEventName = core.StringPtr("focus")
				updateDialogNodeOptionsModel.NewVariable = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewActions = []assistantv1.DialogNodeAction{*dialogNodeActionModel}
				updateDialogNodeOptionsModel.NewDigressIn = core.StringPtr("not_available")
				updateDialogNodeOptionsModel.NewDigressOut = core.StringPtr("allow_returning")
				updateDialogNodeOptionsModel.NewDigressOutSlots = core.StringPtr("not_allowed")
				updateDialogNodeOptionsModel.NewUserLabel = core.StringPtr("testString")
				updateDialogNodeOptionsModel.NewDisambiguationOptOut = core.BoolPtr(false)
				updateDialogNodeOptionsModel.IncludeAudit = core.BoolPtr(false)
				updateDialogNodeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.UpdateDialogNode(updateDialogNodeOptionsModel)
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
	Describe(`DeleteDialogNode(deleteDialogNodeOptions *DeleteDialogNodeOptions)`, func() {
		version := "testString"
		deleteDialogNodePath := "/v1/workspaces/testString/dialog_nodes/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDialogNodePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteDialogNode successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := assistantService.DeleteDialogNode(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteDialogNodeOptions model
				deleteDialogNodeOptionsModel := new(assistantv1.DeleteDialogNodeOptions)
				deleteDialogNodeOptionsModel.WorkspaceID = core.StringPtr("testString")
				deleteDialogNodeOptionsModel.DialogNode = core.StringPtr("testString")
				deleteDialogNodeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = assistantService.DeleteDialogNode(deleteDialogNodeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteDialogNode with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the DeleteDialogNodeOptions model
				deleteDialogNodeOptionsModel := new(assistantv1.DeleteDialogNodeOptions)
				deleteDialogNodeOptionsModel.WorkspaceID = core.StringPtr("testString")
				deleteDialogNodeOptionsModel.DialogNode = core.StringPtr("testString")
				deleteDialogNodeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := assistantService.DeleteDialogNode(deleteDialogNodeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteDialogNodeOptions model with no property values
				deleteDialogNodeOptionsModelNew := new(assistantv1.DeleteDialogNodeOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = assistantService.DeleteDialogNode(deleteDialogNodeOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListLogs(listLogsOptions *ListLogsOptions) - Operation response error`, func() {
		version := "testString"
		listLogsPath := "/v1/workspaces/testString/logs"
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
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListLogs with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ListLogsOptions model
				listLogsOptionsModel := new(assistantv1.ListLogsOptions)
				listLogsOptionsModel.WorkspaceID = core.StringPtr("testString")
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
		listLogsPath := "/v1/workspaces/testString/logs"
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
					fmt.Fprintf(res, "%s", `{"logs": [{"request": {"input": {"text": "Text", "spelling_suggestions": false, "spelling_auto_correct": false, "suggested_text": "SuggestedText", "original_text": "OriginalText"}, "intents": [{"intent": "Intent", "confidence": 10}], "entities": [{"entity": "Entity", "location": [8], "value": "Value", "confidence": 10, "groups": [{"group": "Group", "location": [8]}], "interpretation": {"calendar_type": "CalendarType", "datetime_link": "DatetimeLink", "festival": "Festival", "granularity": "day", "range_link": "RangeLink", "range_modifier": "RangeModifier", "relative_day": 11, "relative_month": 13, "relative_week": 12, "relative_weekend": 15, "relative_year": 12, "specific_day": 11, "specific_day_of_week": "SpecificDayOfWeek", "specific_month": 13, "specific_quarter": 15, "specific_year": 12, "numeric_value": 12, "subtype": "Subtype", "part_of_day": "PartOfDay", "relative_hour": 12, "relative_minute": 14, "relative_second": 14, "specific_hour": 12, "specific_minute": 14, "specific_second": 14, "timezone": "Timezone"}, "alternatives": [{"value": "Value", "confidence": 10}], "role": {"type": "date_from"}}], "alternate_intents": false, "context": {"conversation_id": "ConversationID", "system": {"mapKey": "anyValue"}, "metadata": {"deployment": "Deployment", "user_id": "UserID"}}, "output": {"nodes_visited": ["NodesVisited"], "nodes_visited_details": [{"dialog_node": "DialogNode", "title": "Title", "conditions": "Conditions"}], "log_messages": [{"level": "info", "msg": "Msg", "code": "Code", "source": {"type": "dialog_node", "dialog_node": "DialogNode"}}], "generic": [{"response_type": "option", "title": "Title", "description": "Description", "preference": "dropdown", "options": [{"label": "Label", "value": {"input": {"text": "Text", "spelling_suggestions": false, "spelling_auto_correct": false, "suggested_text": "SuggestedText", "original_text": "OriginalText"}, "intents": [{"intent": "Intent", "confidence": 10}], "entities": [{"entity": "Entity", "location": [8], "value": "Value", "confidence": 10, "groups": [{"group": "Group", "location": [8]}], "interpretation": {"calendar_type": "CalendarType", "datetime_link": "DatetimeLink", "festival": "Festival", "granularity": "day", "range_link": "RangeLink", "range_modifier": "RangeModifier", "relative_day": 11, "relative_month": 13, "relative_week": 12, "relative_weekend": 15, "relative_year": 12, "specific_day": 11, "specific_day_of_week": "SpecificDayOfWeek", "specific_month": 13, "specific_quarter": 15, "specific_year": 12, "numeric_value": 12, "subtype": "Subtype", "part_of_day": "PartOfDay", "relative_hour": 12, "relative_minute": 14, "relative_second": 14, "specific_hour": 12, "specific_minute": 14, "specific_second": 14, "timezone": "Timezone"}, "alternatives": [{"value": "Value", "confidence": 10}], "role": {"type": "date_from"}}]}}], "channels": [{"channel": "chat"}]}]}, "actions": [{"name": "Name", "type": "client", "parameters": {"mapKey": "anyValue"}, "result_variable": "ResultVariable", "credentials": "Credentials"}], "user_id": "UserID"}, "response": {"input": {"text": "Text", "spelling_suggestions": false, "spelling_auto_correct": false, "suggested_text": "SuggestedText", "original_text": "OriginalText"}, "intents": [{"intent": "Intent", "confidence": 10}], "entities": [{"entity": "Entity", "location": [8], "value": "Value", "confidence": 10, "groups": [{"group": "Group", "location": [8]}], "interpretation": {"calendar_type": "CalendarType", "datetime_link": "DatetimeLink", "festival": "Festival", "granularity": "day", "range_link": "RangeLink", "range_modifier": "RangeModifier", "relative_day": 11, "relative_month": 13, "relative_week": 12, "relative_weekend": 15, "relative_year": 12, "specific_day": 11, "specific_day_of_week": "SpecificDayOfWeek", "specific_month": 13, "specific_quarter": 15, "specific_year": 12, "numeric_value": 12, "subtype": "Subtype", "part_of_day": "PartOfDay", "relative_hour": 12, "relative_minute": 14, "relative_second": 14, "specific_hour": 12, "specific_minute": 14, "specific_second": 14, "timezone": "Timezone"}, "alternatives": [{"value": "Value", "confidence": 10}], "role": {"type": "date_from"}}], "alternate_intents": false, "context": {"conversation_id": "ConversationID", "system": {"mapKey": "anyValue"}, "metadata": {"deployment": "Deployment", "user_id": "UserID"}}, "output": {"nodes_visited": ["NodesVisited"], "nodes_visited_details": [{"dialog_node": "DialogNode", "title": "Title", "conditions": "Conditions"}], "log_messages": [{"level": "info", "msg": "Msg", "code": "Code", "source": {"type": "dialog_node", "dialog_node": "DialogNode"}}], "generic": [{"response_type": "option", "title": "Title", "description": "Description", "preference": "dropdown", "options": [{"label": "Label", "value": {"input": {"text": "Text", "spelling_suggestions": false, "spelling_auto_correct": false, "suggested_text": "SuggestedText", "original_text": "OriginalText"}, "intents": [{"intent": "Intent", "confidence": 10}], "entities": [{"entity": "Entity", "location": [8], "value": "Value", "confidence": 10, "groups": [{"group": "Group", "location": [8]}], "interpretation": {"calendar_type": "CalendarType", "datetime_link": "DatetimeLink", "festival": "Festival", "granularity": "day", "range_link": "RangeLink", "range_modifier": "RangeModifier", "relative_day": 11, "relative_month": 13, "relative_week": 12, "relative_weekend": 15, "relative_year": 12, "specific_day": 11, "specific_day_of_week": "SpecificDayOfWeek", "specific_month": 13, "specific_quarter": 15, "specific_year": 12, "numeric_value": 12, "subtype": "Subtype", "part_of_day": "PartOfDay", "relative_hour": 12, "relative_minute": 14, "relative_second": 14, "specific_hour": 12, "specific_minute": 14, "specific_second": 14, "timezone": "Timezone"}, "alternatives": [{"value": "Value", "confidence": 10}], "role": {"type": "date_from"}}]}}], "channels": [{"channel": "chat"}]}]}, "actions": [{"name": "Name", "type": "client", "parameters": {"mapKey": "anyValue"}, "result_variable": "ResultVariable", "credentials": "Credentials"}], "user_id": "UserID"}, "log_id": "LogID", "request_timestamp": "RequestTimestamp", "response_timestamp": "ResponseTimestamp", "workspace_id": "WorkspaceID", "language": "Language"}], "pagination": {"next_url": "NextURL", "matched": 7, "next_cursor": "NextCursor"}}`)
				}))
			})
			It(`Invoke ListLogs successfully with retries`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the ListLogsOptions model
				listLogsOptionsModel := new(assistantv1.ListLogsOptions)
				listLogsOptionsModel.WorkspaceID = core.StringPtr("testString")
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
					fmt.Fprintf(res, "%s", `{"logs": [{"request": {"input": {"text": "Text", "spelling_suggestions": false, "spelling_auto_correct": false, "suggested_text": "SuggestedText", "original_text": "OriginalText"}, "intents": [{"intent": "Intent", "confidence": 10}], "entities": [{"entity": "Entity", "location": [8], "value": "Value", "confidence": 10, "groups": [{"group": "Group", "location": [8]}], "interpretation": {"calendar_type": "CalendarType", "datetime_link": "DatetimeLink", "festival": "Festival", "granularity": "day", "range_link": "RangeLink", "range_modifier": "RangeModifier", "relative_day": 11, "relative_month": 13, "relative_week": 12, "relative_weekend": 15, "relative_year": 12, "specific_day": 11, "specific_day_of_week": "SpecificDayOfWeek", "specific_month": 13, "specific_quarter": 15, "specific_year": 12, "numeric_value": 12, "subtype": "Subtype", "part_of_day": "PartOfDay", "relative_hour": 12, "relative_minute": 14, "relative_second": 14, "specific_hour": 12, "specific_minute": 14, "specific_second": 14, "timezone": "Timezone"}, "alternatives": [{"value": "Value", "confidence": 10}], "role": {"type": "date_from"}}], "alternate_intents": false, "context": {"conversation_id": "ConversationID", "system": {"mapKey": "anyValue"}, "metadata": {"deployment": "Deployment", "user_id": "UserID"}}, "output": {"nodes_visited": ["NodesVisited"], "nodes_visited_details": [{"dialog_node": "DialogNode", "title": "Title", "conditions": "Conditions"}], "log_messages": [{"level": "info", "msg": "Msg", "code": "Code", "source": {"type": "dialog_node", "dialog_node": "DialogNode"}}], "generic": [{"response_type": "option", "title": "Title", "description": "Description", "preference": "dropdown", "options": [{"label": "Label", "value": {"input": {"text": "Text", "spelling_suggestions": false, "spelling_auto_correct": false, "suggested_text": "SuggestedText", "original_text": "OriginalText"}, "intents": [{"intent": "Intent", "confidence": 10}], "entities": [{"entity": "Entity", "location": [8], "value": "Value", "confidence": 10, "groups": [{"group": "Group", "location": [8]}], "interpretation": {"calendar_type": "CalendarType", "datetime_link": "DatetimeLink", "festival": "Festival", "granularity": "day", "range_link": "RangeLink", "range_modifier": "RangeModifier", "relative_day": 11, "relative_month": 13, "relative_week": 12, "relative_weekend": 15, "relative_year": 12, "specific_day": 11, "specific_day_of_week": "SpecificDayOfWeek", "specific_month": 13, "specific_quarter": 15, "specific_year": 12, "numeric_value": 12, "subtype": "Subtype", "part_of_day": "PartOfDay", "relative_hour": 12, "relative_minute": 14, "relative_second": 14, "specific_hour": 12, "specific_minute": 14, "specific_second": 14, "timezone": "Timezone"}, "alternatives": [{"value": "Value", "confidence": 10}], "role": {"type": "date_from"}}]}}], "channels": [{"channel": "chat"}]}]}, "actions": [{"name": "Name", "type": "client", "parameters": {"mapKey": "anyValue"}, "result_variable": "ResultVariable", "credentials": "Credentials"}], "user_id": "UserID"}, "response": {"input": {"text": "Text", "spelling_suggestions": false, "spelling_auto_correct": false, "suggested_text": "SuggestedText", "original_text": "OriginalText"}, "intents": [{"intent": "Intent", "confidence": 10}], "entities": [{"entity": "Entity", "location": [8], "value": "Value", "confidence": 10, "groups": [{"group": "Group", "location": [8]}], "interpretation": {"calendar_type": "CalendarType", "datetime_link": "DatetimeLink", "festival": "Festival", "granularity": "day", "range_link": "RangeLink", "range_modifier": "RangeModifier", "relative_day": 11, "relative_month": 13, "relative_week": 12, "relative_weekend": 15, "relative_year": 12, "specific_day": 11, "specific_day_of_week": "SpecificDayOfWeek", "specific_month": 13, "specific_quarter": 15, "specific_year": 12, "numeric_value": 12, "subtype": "Subtype", "part_of_day": "PartOfDay", "relative_hour": 12, "relative_minute": 14, "relative_second": 14, "specific_hour": 12, "specific_minute": 14, "specific_second": 14, "timezone": "Timezone"}, "alternatives": [{"value": "Value", "confidence": 10}], "role": {"type": "date_from"}}], "alternate_intents": false, "context": {"conversation_id": "ConversationID", "system": {"mapKey": "anyValue"}, "metadata": {"deployment": "Deployment", "user_id": "UserID"}}, "output": {"nodes_visited": ["NodesVisited"], "nodes_visited_details": [{"dialog_node": "DialogNode", "title": "Title", "conditions": "Conditions"}], "log_messages": [{"level": "info", "msg": "Msg", "code": "Code", "source": {"type": "dialog_node", "dialog_node": "DialogNode"}}], "generic": [{"response_type": "option", "title": "Title", "description": "Description", "preference": "dropdown", "options": [{"label": "Label", "value": {"input": {"text": "Text", "spelling_suggestions": false, "spelling_auto_correct": false, "suggested_text": "SuggestedText", "original_text": "OriginalText"}, "intents": [{"intent": "Intent", "confidence": 10}], "entities": [{"entity": "Entity", "location": [8], "value": "Value", "confidence": 10, "groups": [{"group": "Group", "location": [8]}], "interpretation": {"calendar_type": "CalendarType", "datetime_link": "DatetimeLink", "festival": "Festival", "granularity": "day", "range_link": "RangeLink", "range_modifier": "RangeModifier", "relative_day": 11, "relative_month": 13, "relative_week": 12, "relative_weekend": 15, "relative_year": 12, "specific_day": 11, "specific_day_of_week": "SpecificDayOfWeek", "specific_month": 13, "specific_quarter": 15, "specific_year": 12, "numeric_value": 12, "subtype": "Subtype", "part_of_day": "PartOfDay", "relative_hour": 12, "relative_minute": 14, "relative_second": 14, "specific_hour": 12, "specific_minute": 14, "specific_second": 14, "timezone": "Timezone"}, "alternatives": [{"value": "Value", "confidence": 10}], "role": {"type": "date_from"}}]}}], "channels": [{"channel": "chat"}]}]}, "actions": [{"name": "Name", "type": "client", "parameters": {"mapKey": "anyValue"}, "result_variable": "ResultVariable", "credentials": "Credentials"}], "user_id": "UserID"}, "log_id": "LogID", "request_timestamp": "RequestTimestamp", "response_timestamp": "ResponseTimestamp", "workspace_id": "WorkspaceID", "language": "Language"}], "pagination": {"next_url": "NextURL", "matched": 7, "next_cursor": "NextCursor"}}`)
				}))
			})
			It(`Invoke ListLogs successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
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
				listLogsOptionsModel := new(assistantv1.ListLogsOptions)
				listLogsOptionsModel.WorkspaceID = core.StringPtr("testString")
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
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ListLogsOptions model
				listLogsOptionsModel := new(assistantv1.ListLogsOptions)
				listLogsOptionsModel.WorkspaceID = core.StringPtr("testString")
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
				listLogsOptionsModelNew := new(assistantv1.ListLogsOptions)
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
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ListLogsOptions model
				listLogsOptionsModel := new(assistantv1.ListLogsOptions)
				listLogsOptionsModel.WorkspaceID = core.StringPtr("testString")
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
	Describe(`ListAllLogs(listAllLogsOptions *ListAllLogsOptions) - Operation response error`, func() {
		version := "testString"
		listAllLogsPath := "/v1/logs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAllLogsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["filter"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["page_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["cursor"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAllLogs with error: Operation response processing error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ListAllLogsOptions model
				listAllLogsOptionsModel := new(assistantv1.ListAllLogsOptions)
				listAllLogsOptionsModel.Filter = core.StringPtr("testString")
				listAllLogsOptionsModel.Sort = core.StringPtr("testString")
				listAllLogsOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listAllLogsOptionsModel.Cursor = core.StringPtr("testString")
				listAllLogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := assistantService.ListAllLogs(listAllLogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				assistantService.EnableRetries(0, 0)
				result, response, operationErr = assistantService.ListAllLogs(listAllLogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListAllLogs(listAllLogsOptions *ListAllLogsOptions)`, func() {
		version := "testString"
		listAllLogsPath := "/v1/logs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAllLogsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["filter"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["page_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["cursor"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"logs": [{"request": {"input": {"text": "Text", "spelling_suggestions": false, "spelling_auto_correct": false, "suggested_text": "SuggestedText", "original_text": "OriginalText"}, "intents": [{"intent": "Intent", "confidence": 10}], "entities": [{"entity": "Entity", "location": [8], "value": "Value", "confidence": 10, "groups": [{"group": "Group", "location": [8]}], "interpretation": {"calendar_type": "CalendarType", "datetime_link": "DatetimeLink", "festival": "Festival", "granularity": "day", "range_link": "RangeLink", "range_modifier": "RangeModifier", "relative_day": 11, "relative_month": 13, "relative_week": 12, "relative_weekend": 15, "relative_year": 12, "specific_day": 11, "specific_day_of_week": "SpecificDayOfWeek", "specific_month": 13, "specific_quarter": 15, "specific_year": 12, "numeric_value": 12, "subtype": "Subtype", "part_of_day": "PartOfDay", "relative_hour": 12, "relative_minute": 14, "relative_second": 14, "specific_hour": 12, "specific_minute": 14, "specific_second": 14, "timezone": "Timezone"}, "alternatives": [{"value": "Value", "confidence": 10}], "role": {"type": "date_from"}}], "alternate_intents": false, "context": {"conversation_id": "ConversationID", "system": {"mapKey": "anyValue"}, "metadata": {"deployment": "Deployment", "user_id": "UserID"}}, "output": {"nodes_visited": ["NodesVisited"], "nodes_visited_details": [{"dialog_node": "DialogNode", "title": "Title", "conditions": "Conditions"}], "log_messages": [{"level": "info", "msg": "Msg", "code": "Code", "source": {"type": "dialog_node", "dialog_node": "DialogNode"}}], "generic": [{"response_type": "option", "title": "Title", "description": "Description", "preference": "dropdown", "options": [{"label": "Label", "value": {"input": {"text": "Text", "spelling_suggestions": false, "spelling_auto_correct": false, "suggested_text": "SuggestedText", "original_text": "OriginalText"}, "intents": [{"intent": "Intent", "confidence": 10}], "entities": [{"entity": "Entity", "location": [8], "value": "Value", "confidence": 10, "groups": [{"group": "Group", "location": [8]}], "interpretation": {"calendar_type": "CalendarType", "datetime_link": "DatetimeLink", "festival": "Festival", "granularity": "day", "range_link": "RangeLink", "range_modifier": "RangeModifier", "relative_day": 11, "relative_month": 13, "relative_week": 12, "relative_weekend": 15, "relative_year": 12, "specific_day": 11, "specific_day_of_week": "SpecificDayOfWeek", "specific_month": 13, "specific_quarter": 15, "specific_year": 12, "numeric_value": 12, "subtype": "Subtype", "part_of_day": "PartOfDay", "relative_hour": 12, "relative_minute": 14, "relative_second": 14, "specific_hour": 12, "specific_minute": 14, "specific_second": 14, "timezone": "Timezone"}, "alternatives": [{"value": "Value", "confidence": 10}], "role": {"type": "date_from"}}]}}], "channels": [{"channel": "chat"}]}]}, "actions": [{"name": "Name", "type": "client", "parameters": {"mapKey": "anyValue"}, "result_variable": "ResultVariable", "credentials": "Credentials"}], "user_id": "UserID"}, "response": {"input": {"text": "Text", "spelling_suggestions": false, "spelling_auto_correct": false, "suggested_text": "SuggestedText", "original_text": "OriginalText"}, "intents": [{"intent": "Intent", "confidence": 10}], "entities": [{"entity": "Entity", "location": [8], "value": "Value", "confidence": 10, "groups": [{"group": "Group", "location": [8]}], "interpretation": {"calendar_type": "CalendarType", "datetime_link": "DatetimeLink", "festival": "Festival", "granularity": "day", "range_link": "RangeLink", "range_modifier": "RangeModifier", "relative_day": 11, "relative_month": 13, "relative_week": 12, "relative_weekend": 15, "relative_year": 12, "specific_day": 11, "specific_day_of_week": "SpecificDayOfWeek", "specific_month": 13, "specific_quarter": 15, "specific_year": 12, "numeric_value": 12, "subtype": "Subtype", "part_of_day": "PartOfDay", "relative_hour": 12, "relative_minute": 14, "relative_second": 14, "specific_hour": 12, "specific_minute": 14, "specific_second": 14, "timezone": "Timezone"}, "alternatives": [{"value": "Value", "confidence": 10}], "role": {"type": "date_from"}}], "alternate_intents": false, "context": {"conversation_id": "ConversationID", "system": {"mapKey": "anyValue"}, "metadata": {"deployment": "Deployment", "user_id": "UserID"}}, "output": {"nodes_visited": ["NodesVisited"], "nodes_visited_details": [{"dialog_node": "DialogNode", "title": "Title", "conditions": "Conditions"}], "log_messages": [{"level": "info", "msg": "Msg", "code": "Code", "source": {"type": "dialog_node", "dialog_node": "DialogNode"}}], "generic": [{"response_type": "option", "title": "Title", "description": "Description", "preference": "dropdown", "options": [{"label": "Label", "value": {"input": {"text": "Text", "spelling_suggestions": false, "spelling_auto_correct": false, "suggested_text": "SuggestedText", "original_text": "OriginalText"}, "intents": [{"intent": "Intent", "confidence": 10}], "entities": [{"entity": "Entity", "location": [8], "value": "Value", "confidence": 10, "groups": [{"group": "Group", "location": [8]}], "interpretation": {"calendar_type": "CalendarType", "datetime_link": "DatetimeLink", "festival": "Festival", "granularity": "day", "range_link": "RangeLink", "range_modifier": "RangeModifier", "relative_day": 11, "relative_month": 13, "relative_week": 12, "relative_weekend": 15, "relative_year": 12, "specific_day": 11, "specific_day_of_week": "SpecificDayOfWeek", "specific_month": 13, "specific_quarter": 15, "specific_year": 12, "numeric_value": 12, "subtype": "Subtype", "part_of_day": "PartOfDay", "relative_hour": 12, "relative_minute": 14, "relative_second": 14, "specific_hour": 12, "specific_minute": 14, "specific_second": 14, "timezone": "Timezone"}, "alternatives": [{"value": "Value", "confidence": 10}], "role": {"type": "date_from"}}]}}], "channels": [{"channel": "chat"}]}]}, "actions": [{"name": "Name", "type": "client", "parameters": {"mapKey": "anyValue"}, "result_variable": "ResultVariable", "credentials": "Credentials"}], "user_id": "UserID"}, "log_id": "LogID", "request_timestamp": "RequestTimestamp", "response_timestamp": "ResponseTimestamp", "workspace_id": "WorkspaceID", "language": "Language"}], "pagination": {"next_url": "NextURL", "matched": 7, "next_cursor": "NextCursor"}}`)
				}))
			})
			It(`Invoke ListAllLogs successfully with retries`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())
				assistantService.EnableRetries(0, 0)

				// Construct an instance of the ListAllLogsOptions model
				listAllLogsOptionsModel := new(assistantv1.ListAllLogsOptions)
				listAllLogsOptionsModel.Filter = core.StringPtr("testString")
				listAllLogsOptionsModel.Sort = core.StringPtr("testString")
				listAllLogsOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listAllLogsOptionsModel.Cursor = core.StringPtr("testString")
				listAllLogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := assistantService.ListAllLogsWithContext(ctx, listAllLogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				assistantService.DisableRetries()
				result, response, operationErr := assistantService.ListAllLogs(listAllLogsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = assistantService.ListAllLogsWithContext(ctx, listAllLogsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listAllLogsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["filter"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["page_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["cursor"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"logs": [{"request": {"input": {"text": "Text", "spelling_suggestions": false, "spelling_auto_correct": false, "suggested_text": "SuggestedText", "original_text": "OriginalText"}, "intents": [{"intent": "Intent", "confidence": 10}], "entities": [{"entity": "Entity", "location": [8], "value": "Value", "confidence": 10, "groups": [{"group": "Group", "location": [8]}], "interpretation": {"calendar_type": "CalendarType", "datetime_link": "DatetimeLink", "festival": "Festival", "granularity": "day", "range_link": "RangeLink", "range_modifier": "RangeModifier", "relative_day": 11, "relative_month": 13, "relative_week": 12, "relative_weekend": 15, "relative_year": 12, "specific_day": 11, "specific_day_of_week": "SpecificDayOfWeek", "specific_month": 13, "specific_quarter": 15, "specific_year": 12, "numeric_value": 12, "subtype": "Subtype", "part_of_day": "PartOfDay", "relative_hour": 12, "relative_minute": 14, "relative_second": 14, "specific_hour": 12, "specific_minute": 14, "specific_second": 14, "timezone": "Timezone"}, "alternatives": [{"value": "Value", "confidence": 10}], "role": {"type": "date_from"}}], "alternate_intents": false, "context": {"conversation_id": "ConversationID", "system": {"mapKey": "anyValue"}, "metadata": {"deployment": "Deployment", "user_id": "UserID"}}, "output": {"nodes_visited": ["NodesVisited"], "nodes_visited_details": [{"dialog_node": "DialogNode", "title": "Title", "conditions": "Conditions"}], "log_messages": [{"level": "info", "msg": "Msg", "code": "Code", "source": {"type": "dialog_node", "dialog_node": "DialogNode"}}], "generic": [{"response_type": "option", "title": "Title", "description": "Description", "preference": "dropdown", "options": [{"label": "Label", "value": {"input": {"text": "Text", "spelling_suggestions": false, "spelling_auto_correct": false, "suggested_text": "SuggestedText", "original_text": "OriginalText"}, "intents": [{"intent": "Intent", "confidence": 10}], "entities": [{"entity": "Entity", "location": [8], "value": "Value", "confidence": 10, "groups": [{"group": "Group", "location": [8]}], "interpretation": {"calendar_type": "CalendarType", "datetime_link": "DatetimeLink", "festival": "Festival", "granularity": "day", "range_link": "RangeLink", "range_modifier": "RangeModifier", "relative_day": 11, "relative_month": 13, "relative_week": 12, "relative_weekend": 15, "relative_year": 12, "specific_day": 11, "specific_day_of_week": "SpecificDayOfWeek", "specific_month": 13, "specific_quarter": 15, "specific_year": 12, "numeric_value": 12, "subtype": "Subtype", "part_of_day": "PartOfDay", "relative_hour": 12, "relative_minute": 14, "relative_second": 14, "specific_hour": 12, "specific_minute": 14, "specific_second": 14, "timezone": "Timezone"}, "alternatives": [{"value": "Value", "confidence": 10}], "role": {"type": "date_from"}}]}}], "channels": [{"channel": "chat"}]}]}, "actions": [{"name": "Name", "type": "client", "parameters": {"mapKey": "anyValue"}, "result_variable": "ResultVariable", "credentials": "Credentials"}], "user_id": "UserID"}, "response": {"input": {"text": "Text", "spelling_suggestions": false, "spelling_auto_correct": false, "suggested_text": "SuggestedText", "original_text": "OriginalText"}, "intents": [{"intent": "Intent", "confidence": 10}], "entities": [{"entity": "Entity", "location": [8], "value": "Value", "confidence": 10, "groups": [{"group": "Group", "location": [8]}], "interpretation": {"calendar_type": "CalendarType", "datetime_link": "DatetimeLink", "festival": "Festival", "granularity": "day", "range_link": "RangeLink", "range_modifier": "RangeModifier", "relative_day": 11, "relative_month": 13, "relative_week": 12, "relative_weekend": 15, "relative_year": 12, "specific_day": 11, "specific_day_of_week": "SpecificDayOfWeek", "specific_month": 13, "specific_quarter": 15, "specific_year": 12, "numeric_value": 12, "subtype": "Subtype", "part_of_day": "PartOfDay", "relative_hour": 12, "relative_minute": 14, "relative_second": 14, "specific_hour": 12, "specific_minute": 14, "specific_second": 14, "timezone": "Timezone"}, "alternatives": [{"value": "Value", "confidence": 10}], "role": {"type": "date_from"}}], "alternate_intents": false, "context": {"conversation_id": "ConversationID", "system": {"mapKey": "anyValue"}, "metadata": {"deployment": "Deployment", "user_id": "UserID"}}, "output": {"nodes_visited": ["NodesVisited"], "nodes_visited_details": [{"dialog_node": "DialogNode", "title": "Title", "conditions": "Conditions"}], "log_messages": [{"level": "info", "msg": "Msg", "code": "Code", "source": {"type": "dialog_node", "dialog_node": "DialogNode"}}], "generic": [{"response_type": "option", "title": "Title", "description": "Description", "preference": "dropdown", "options": [{"label": "Label", "value": {"input": {"text": "Text", "spelling_suggestions": false, "spelling_auto_correct": false, "suggested_text": "SuggestedText", "original_text": "OriginalText"}, "intents": [{"intent": "Intent", "confidence": 10}], "entities": [{"entity": "Entity", "location": [8], "value": "Value", "confidence": 10, "groups": [{"group": "Group", "location": [8]}], "interpretation": {"calendar_type": "CalendarType", "datetime_link": "DatetimeLink", "festival": "Festival", "granularity": "day", "range_link": "RangeLink", "range_modifier": "RangeModifier", "relative_day": 11, "relative_month": 13, "relative_week": 12, "relative_weekend": 15, "relative_year": 12, "specific_day": 11, "specific_day_of_week": "SpecificDayOfWeek", "specific_month": 13, "specific_quarter": 15, "specific_year": 12, "numeric_value": 12, "subtype": "Subtype", "part_of_day": "PartOfDay", "relative_hour": 12, "relative_minute": 14, "relative_second": 14, "specific_hour": 12, "specific_minute": 14, "specific_second": 14, "timezone": "Timezone"}, "alternatives": [{"value": "Value", "confidence": 10}], "role": {"type": "date_from"}}]}}], "channels": [{"channel": "chat"}]}]}, "actions": [{"name": "Name", "type": "client", "parameters": {"mapKey": "anyValue"}, "result_variable": "ResultVariable", "credentials": "Credentials"}], "user_id": "UserID"}, "log_id": "LogID", "request_timestamp": "RequestTimestamp", "response_timestamp": "ResponseTimestamp", "workspace_id": "WorkspaceID", "language": "Language"}], "pagination": {"next_url": "NextURL", "matched": 7, "next_cursor": "NextCursor"}}`)
				}))
			})
			It(`Invoke ListAllLogs successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := assistantService.ListAllLogs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAllLogsOptions model
				listAllLogsOptionsModel := new(assistantv1.ListAllLogsOptions)
				listAllLogsOptionsModel.Filter = core.StringPtr("testString")
				listAllLogsOptionsModel.Sort = core.StringPtr("testString")
				listAllLogsOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listAllLogsOptionsModel.Cursor = core.StringPtr("testString")
				listAllLogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = assistantService.ListAllLogs(listAllLogsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListAllLogs with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ListAllLogsOptions model
				listAllLogsOptionsModel := new(assistantv1.ListAllLogsOptions)
				listAllLogsOptionsModel.Filter = core.StringPtr("testString")
				listAllLogsOptionsModel.Sort = core.StringPtr("testString")
				listAllLogsOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listAllLogsOptionsModel.Cursor = core.StringPtr("testString")
				listAllLogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := assistantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := assistantService.ListAllLogs(listAllLogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListAllLogsOptions model with no property values
				listAllLogsOptionsModelNew := new(assistantv1.ListAllLogsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = assistantService.ListAllLogs(listAllLogsOptionsModelNew)
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
			It(`Invoke ListAllLogs successfully`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the ListAllLogsOptions model
				listAllLogsOptionsModel := new(assistantv1.ListAllLogsOptions)
				listAllLogsOptionsModel.Filter = core.StringPtr("testString")
				listAllLogsOptionsModel.Sort = core.StringPtr("testString")
				listAllLogsOptionsModel.PageLimit = core.Int64Ptr(int64(38))
				listAllLogsOptionsModel.Cursor = core.StringPtr("testString")
				listAllLogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := assistantService.ListAllLogs(listAllLogsOptionsModel)
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
		deleteUserDataPath := "/v1/user_data"
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
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
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
				deleteUserDataOptionsModel := new(assistantv1.DeleteUserDataOptions)
				deleteUserDataOptionsModel.CustomerID = core.StringPtr("testString")
				deleteUserDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = assistantService.DeleteUserData(deleteUserDataOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteUserData with error: Operation validation and request error`, func() {
				assistantService, serviceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(assistantService).ToNot(BeNil())

				// Construct an instance of the DeleteUserDataOptions model
				deleteUserDataOptionsModel := new(assistantv1.DeleteUserDataOptions)
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
				deleteUserDataOptionsModelNew := new(assistantv1.DeleteUserDataOptions)
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
			assistantService, _ := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
				URL:           "http://assistantv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       core.StringPtr(version),
			})
			It(`Invoke NewBulkClassifyOptions successfully`, func() {
				// Construct an instance of the BulkClassifyUtterance model
				bulkClassifyUtteranceModel := new(assistantv1.BulkClassifyUtterance)
				Expect(bulkClassifyUtteranceModel).ToNot(BeNil())
				bulkClassifyUtteranceModel.Text = core.StringPtr("testString")
				Expect(bulkClassifyUtteranceModel.Text).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the BulkClassifyOptions model
				workspaceID := "testString"
				bulkClassifyOptionsModel := assistantService.NewBulkClassifyOptions(workspaceID)
				bulkClassifyOptionsModel.SetWorkspaceID("testString")
				bulkClassifyOptionsModel.SetInput([]assistantv1.BulkClassifyUtterance{*bulkClassifyUtteranceModel})
				bulkClassifyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(bulkClassifyOptionsModel).ToNot(BeNil())
				Expect(bulkClassifyOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(bulkClassifyOptionsModel.Input).To(Equal([]assistantv1.BulkClassifyUtterance{*bulkClassifyUtteranceModel}))
				Expect(bulkClassifyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewBulkClassifyUtterance successfully`, func() {
				text := "testString"
				_model, err := assistantService.NewBulkClassifyUtterance(text)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCaptureGroup successfully`, func() {
				group := "testString"
				_model, err := assistantService.NewCaptureGroup(group)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewChannelTransferInfo successfully`, func() {
				var target *assistantv1.ChannelTransferTarget = nil
				_, err := assistantService.NewChannelTransferInfo(target)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewCounterexample successfully`, func() {
				text := "testString"
				_model, err := assistantService.NewCounterexample(text)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCreateCounterexampleOptions successfully`, func() {
				// Construct an instance of the CreateCounterexampleOptions model
				workspaceID := "testString"
				createCounterexampleOptionsText := "testString"
				createCounterexampleOptionsModel := assistantService.NewCreateCounterexampleOptions(workspaceID, createCounterexampleOptionsText)
				createCounterexampleOptionsModel.SetWorkspaceID("testString")
				createCounterexampleOptionsModel.SetText("testString")
				createCounterexampleOptionsModel.SetIncludeAudit(false)
				createCounterexampleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createCounterexampleOptionsModel).ToNot(BeNil())
				Expect(createCounterexampleOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(createCounterexampleOptionsModel.Text).To(Equal(core.StringPtr("testString")))
				Expect(createCounterexampleOptionsModel.IncludeAudit).To(Equal(core.BoolPtr(false)))
				Expect(createCounterexampleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateDialogNodeOptions successfully`, func() {
				// Construct an instance of the ResponseGenericChannel model
				responseGenericChannelModel := new(assistantv1.ResponseGenericChannel)
				Expect(responseGenericChannelModel).ToNot(BeNil())
				responseGenericChannelModel.Channel = core.StringPtr("chat")
				Expect(responseGenericChannelModel.Channel).To(Equal(core.StringPtr("chat")))

				// Construct an instance of the DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo model
				dialogNodeOutputGenericModel := new(assistantv1.DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo)
				Expect(dialogNodeOutputGenericModel).ToNot(BeNil())
				dialogNodeOutputGenericModel.ResponseType = core.StringPtr("video")
				dialogNodeOutputGenericModel.Source = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Title = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Description = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Channels = []assistantv1.ResponseGenericChannel{*responseGenericChannelModel}
				dialogNodeOutputGenericModel.ChannelOptions = map[string]interface{}{"anyKey": "anyValue"}
				dialogNodeOutputGenericModel.AltText = core.StringPtr("testString")
				Expect(dialogNodeOutputGenericModel.ResponseType).To(Equal(core.StringPtr("video")))
				Expect(dialogNodeOutputGenericModel.Source).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeOutputGenericModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeOutputGenericModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeOutputGenericModel.Channels).To(Equal([]assistantv1.ResponseGenericChannel{*responseGenericChannelModel}))
				Expect(dialogNodeOutputGenericModel.ChannelOptions).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(dialogNodeOutputGenericModel.AltText).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the DialogNodeOutputModifiers model
				dialogNodeOutputModifiersModel := new(assistantv1.DialogNodeOutputModifiers)
				Expect(dialogNodeOutputModifiersModel).ToNot(BeNil())
				dialogNodeOutputModifiersModel.Overwrite = core.BoolPtr(true)
				Expect(dialogNodeOutputModifiersModel.Overwrite).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the DialogNodeOutput model
				dialogNodeOutputModel := new(assistantv1.DialogNodeOutput)
				Expect(dialogNodeOutputModel).ToNot(BeNil())
				dialogNodeOutputModel.Generic = []assistantv1.DialogNodeOutputGenericIntf{dialogNodeOutputGenericModel}
				dialogNodeOutputModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeOutputModel.Modifiers = dialogNodeOutputModifiersModel
				dialogNodeOutputModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(dialogNodeOutputModel.Generic).To(Equal([]assistantv1.DialogNodeOutputGenericIntf{dialogNodeOutputGenericModel}))
				Expect(dialogNodeOutputModel.Integrations).To(Equal(make(map[string]map[string]interface{})))
				Expect(dialogNodeOutputModel.Modifiers).To(Equal(dialogNodeOutputModifiersModel))
				Expect(dialogNodeOutputModel.GetProperties()).ToNot(BeEmpty())
				Expect(dialogNodeOutputModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))

				dialogNodeOutputModel.SetProperties(nil)
				Expect(dialogNodeOutputModel.GetProperties()).To(BeEmpty())

				dialogNodeOutputModelExpectedMap := make(map[string]interface{})
				dialogNodeOutputModelExpectedMap["foo"] = core.StringPtr("testString")
				dialogNodeOutputModel.SetProperties(dialogNodeOutputModelExpectedMap)
				dialogNodeOutputModelActualMap := dialogNodeOutputModel.GetProperties()
				Expect(dialogNodeOutputModelActualMap).To(Equal(dialogNodeOutputModelExpectedMap))

				// Construct an instance of the DialogNodeContext model
				dialogNodeContextModel := new(assistantv1.DialogNodeContext)
				Expect(dialogNodeContextModel).ToNot(BeNil())
				dialogNodeContextModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeContextModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(dialogNodeContextModel.Integrations).To(Equal(make(map[string]map[string]interface{})))
				Expect(dialogNodeContextModel.GetProperties()).ToNot(BeEmpty())
				Expect(dialogNodeContextModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))

				dialogNodeContextModel.SetProperties(nil)
				Expect(dialogNodeContextModel.GetProperties()).To(BeEmpty())

				dialogNodeContextModelExpectedMap := make(map[string]interface{})
				dialogNodeContextModelExpectedMap["foo"] = core.StringPtr("testString")
				dialogNodeContextModel.SetProperties(dialogNodeContextModelExpectedMap)
				dialogNodeContextModelActualMap := dialogNodeContextModel.GetProperties()
				Expect(dialogNodeContextModelActualMap).To(Equal(dialogNodeContextModelExpectedMap))

				// Construct an instance of the DialogNodeNextStep model
				dialogNodeNextStepModel := new(assistantv1.DialogNodeNextStep)
				Expect(dialogNodeNextStepModel).ToNot(BeNil())
				dialogNodeNextStepModel.Behavior = core.StringPtr("get_user_input")
				dialogNodeNextStepModel.DialogNode = core.StringPtr("testString")
				dialogNodeNextStepModel.Selector = core.StringPtr("condition")
				Expect(dialogNodeNextStepModel.Behavior).To(Equal(core.StringPtr("get_user_input")))
				Expect(dialogNodeNextStepModel.DialogNode).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeNextStepModel.Selector).To(Equal(core.StringPtr("condition")))

				// Construct an instance of the DialogNodeAction model
				dialogNodeActionModel := new(assistantv1.DialogNodeAction)
				Expect(dialogNodeActionModel).ToNot(BeNil())
				dialogNodeActionModel.Name = core.StringPtr("testString")
				dialogNodeActionModel.Type = core.StringPtr("client")
				dialogNodeActionModel.Parameters = make(map[string]interface{})
				dialogNodeActionModel.ResultVariable = core.StringPtr("testString")
				dialogNodeActionModel.Credentials = core.StringPtr("testString")
				Expect(dialogNodeActionModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeActionModel.Type).To(Equal(core.StringPtr("client")))
				Expect(dialogNodeActionModel.Parameters).To(Equal(make(map[string]interface{})))
				Expect(dialogNodeActionModel.ResultVariable).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeActionModel.Credentials).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CreateDialogNodeOptions model
				workspaceID := "testString"
				createDialogNodeOptionsDialogNode := "testString"
				createDialogNodeOptionsModel := assistantService.NewCreateDialogNodeOptions(workspaceID, createDialogNodeOptionsDialogNode)
				createDialogNodeOptionsModel.SetWorkspaceID("testString")
				createDialogNodeOptionsModel.SetDialogNode("testString")
				createDialogNodeOptionsModel.SetDescription("testString")
				createDialogNodeOptionsModel.SetConditions("testString")
				createDialogNodeOptionsModel.SetParent("testString")
				createDialogNodeOptionsModel.SetPreviousSibling("testString")
				createDialogNodeOptionsModel.SetOutput(dialogNodeOutputModel)
				createDialogNodeOptionsModel.SetContext(dialogNodeContextModel)
				createDialogNodeOptionsModel.SetMetadata(make(map[string]interface{}))
				createDialogNodeOptionsModel.SetNextStep(dialogNodeNextStepModel)
				createDialogNodeOptionsModel.SetTitle("testString")
				createDialogNodeOptionsModel.SetType("standard")
				createDialogNodeOptionsModel.SetEventName("focus")
				createDialogNodeOptionsModel.SetVariable("testString")
				createDialogNodeOptionsModel.SetActions([]assistantv1.DialogNodeAction{*dialogNodeActionModel})
				createDialogNodeOptionsModel.SetDigressIn("not_available")
				createDialogNodeOptionsModel.SetDigressOut("allow_returning")
				createDialogNodeOptionsModel.SetDigressOutSlots("not_allowed")
				createDialogNodeOptionsModel.SetUserLabel("testString")
				createDialogNodeOptionsModel.SetDisambiguationOptOut(false)
				createDialogNodeOptionsModel.SetIncludeAudit(false)
				createDialogNodeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createDialogNodeOptionsModel).ToNot(BeNil())
				Expect(createDialogNodeOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(createDialogNodeOptionsModel.DialogNode).To(Equal(core.StringPtr("testString")))
				Expect(createDialogNodeOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createDialogNodeOptionsModel.Conditions).To(Equal(core.StringPtr("testString")))
				Expect(createDialogNodeOptionsModel.Parent).To(Equal(core.StringPtr("testString")))
				Expect(createDialogNodeOptionsModel.PreviousSibling).To(Equal(core.StringPtr("testString")))
				Expect(createDialogNodeOptionsModel.Output).To(Equal(dialogNodeOutputModel))
				Expect(createDialogNodeOptionsModel.Context).To(Equal(dialogNodeContextModel))
				Expect(createDialogNodeOptionsModel.Metadata).To(Equal(make(map[string]interface{})))
				Expect(createDialogNodeOptionsModel.NextStep).To(Equal(dialogNodeNextStepModel))
				Expect(createDialogNodeOptionsModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(createDialogNodeOptionsModel.Type).To(Equal(core.StringPtr("standard")))
				Expect(createDialogNodeOptionsModel.EventName).To(Equal(core.StringPtr("focus")))
				Expect(createDialogNodeOptionsModel.Variable).To(Equal(core.StringPtr("testString")))
				Expect(createDialogNodeOptionsModel.Actions).To(Equal([]assistantv1.DialogNodeAction{*dialogNodeActionModel}))
				Expect(createDialogNodeOptionsModel.DigressIn).To(Equal(core.StringPtr("not_available")))
				Expect(createDialogNodeOptionsModel.DigressOut).To(Equal(core.StringPtr("allow_returning")))
				Expect(createDialogNodeOptionsModel.DigressOutSlots).To(Equal(core.StringPtr("not_allowed")))
				Expect(createDialogNodeOptionsModel.UserLabel).To(Equal(core.StringPtr("testString")))
				Expect(createDialogNodeOptionsModel.DisambiguationOptOut).To(Equal(core.BoolPtr(false)))
				Expect(createDialogNodeOptionsModel.IncludeAudit).To(Equal(core.BoolPtr(false)))
				Expect(createDialogNodeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateEntity successfully`, func() {
				entity := "testString"
				_model, err := assistantService.NewCreateEntity(entity)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCreateEntityOptions successfully`, func() {
				// Construct an instance of the CreateValue model
				createValueModel := new(assistantv1.CreateValue)
				Expect(createValueModel).ToNot(BeNil())
				createValueModel.Value = core.StringPtr("testString")
				createValueModel.Metadata = make(map[string]interface{})
				createValueModel.Type = core.StringPtr("synonyms")
				createValueModel.Synonyms = []string{"testString"}
				createValueModel.Patterns = []string{"testString"}
				Expect(createValueModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(createValueModel.Metadata).To(Equal(make(map[string]interface{})))
				Expect(createValueModel.Type).To(Equal(core.StringPtr("synonyms")))
				Expect(createValueModel.Synonyms).To(Equal([]string{"testString"}))
				Expect(createValueModel.Patterns).To(Equal([]string{"testString"}))

				// Construct an instance of the CreateEntityOptions model
				workspaceID := "testString"
				createEntityOptionsEntity := "testString"
				createEntityOptionsModel := assistantService.NewCreateEntityOptions(workspaceID, createEntityOptionsEntity)
				createEntityOptionsModel.SetWorkspaceID("testString")
				createEntityOptionsModel.SetEntity("testString")
				createEntityOptionsModel.SetDescription("testString")
				createEntityOptionsModel.SetMetadata(make(map[string]interface{}))
				createEntityOptionsModel.SetFuzzyMatch(true)
				createEntityOptionsModel.SetValues([]assistantv1.CreateValue{*createValueModel})
				createEntityOptionsModel.SetIncludeAudit(false)
				createEntityOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createEntityOptionsModel).ToNot(BeNil())
				Expect(createEntityOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(createEntityOptionsModel.Entity).To(Equal(core.StringPtr("testString")))
				Expect(createEntityOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createEntityOptionsModel.Metadata).To(Equal(make(map[string]interface{})))
				Expect(createEntityOptionsModel.FuzzyMatch).To(Equal(core.BoolPtr(true)))
				Expect(createEntityOptionsModel.Values).To(Equal([]assistantv1.CreateValue{*createValueModel}))
				Expect(createEntityOptionsModel.IncludeAudit).To(Equal(core.BoolPtr(false)))
				Expect(createEntityOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateExampleOptions successfully`, func() {
				// Construct an instance of the Mention model
				mentionModel := new(assistantv1.Mention)
				Expect(mentionModel).ToNot(BeNil())
				mentionModel.Entity = core.StringPtr("testString")
				mentionModel.Location = []int64{int64(38)}
				Expect(mentionModel.Entity).To(Equal(core.StringPtr("testString")))
				Expect(mentionModel.Location).To(Equal([]int64{int64(38)}))

				// Construct an instance of the CreateExampleOptions model
				workspaceID := "testString"
				intent := "testString"
				createExampleOptionsText := "testString"
				createExampleOptionsModel := assistantService.NewCreateExampleOptions(workspaceID, intent, createExampleOptionsText)
				createExampleOptionsModel.SetWorkspaceID("testString")
				createExampleOptionsModel.SetIntent("testString")
				createExampleOptionsModel.SetText("testString")
				createExampleOptionsModel.SetMentions([]assistantv1.Mention{*mentionModel})
				createExampleOptionsModel.SetIncludeAudit(false)
				createExampleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createExampleOptionsModel).ToNot(BeNil())
				Expect(createExampleOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(createExampleOptionsModel.Intent).To(Equal(core.StringPtr("testString")))
				Expect(createExampleOptionsModel.Text).To(Equal(core.StringPtr("testString")))
				Expect(createExampleOptionsModel.Mentions).To(Equal([]assistantv1.Mention{*mentionModel}))
				Expect(createExampleOptionsModel.IncludeAudit).To(Equal(core.BoolPtr(false)))
				Expect(createExampleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateIntent successfully`, func() {
				intent := "testString"
				_model, err := assistantService.NewCreateIntent(intent)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCreateIntentOptions successfully`, func() {
				// Construct an instance of the Mention model
				mentionModel := new(assistantv1.Mention)
				Expect(mentionModel).ToNot(BeNil())
				mentionModel.Entity = core.StringPtr("testString")
				mentionModel.Location = []int64{int64(38)}
				Expect(mentionModel.Entity).To(Equal(core.StringPtr("testString")))
				Expect(mentionModel.Location).To(Equal([]int64{int64(38)}))

				// Construct an instance of the Example model
				exampleModel := new(assistantv1.Example)
				Expect(exampleModel).ToNot(BeNil())
				exampleModel.Text = core.StringPtr("testString")
				exampleModel.Mentions = []assistantv1.Mention{*mentionModel}
				Expect(exampleModel.Text).To(Equal(core.StringPtr("testString")))
				Expect(exampleModel.Mentions).To(Equal([]assistantv1.Mention{*mentionModel}))

				// Construct an instance of the CreateIntentOptions model
				workspaceID := "testString"
				createIntentOptionsIntent := "testString"
				createIntentOptionsModel := assistantService.NewCreateIntentOptions(workspaceID, createIntentOptionsIntent)
				createIntentOptionsModel.SetWorkspaceID("testString")
				createIntentOptionsModel.SetIntent("testString")
				createIntentOptionsModel.SetDescription("testString")
				createIntentOptionsModel.SetExamples([]assistantv1.Example{*exampleModel})
				createIntentOptionsModel.SetIncludeAudit(false)
				createIntentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createIntentOptionsModel).ToNot(BeNil())
				Expect(createIntentOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(createIntentOptionsModel.Intent).To(Equal(core.StringPtr("testString")))
				Expect(createIntentOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createIntentOptionsModel.Examples).To(Equal([]assistantv1.Example{*exampleModel}))
				Expect(createIntentOptionsModel.IncludeAudit).To(Equal(core.BoolPtr(false)))
				Expect(createIntentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateSynonymOptions successfully`, func() {
				// Construct an instance of the CreateSynonymOptions model
				workspaceID := "testString"
				entity := "testString"
				value := "testString"
				createSynonymOptionsSynonym := "testString"
				createSynonymOptionsModel := assistantService.NewCreateSynonymOptions(workspaceID, entity, value, createSynonymOptionsSynonym)
				createSynonymOptionsModel.SetWorkspaceID("testString")
				createSynonymOptionsModel.SetEntity("testString")
				createSynonymOptionsModel.SetValue("testString")
				createSynonymOptionsModel.SetSynonym("testString")
				createSynonymOptionsModel.SetIncludeAudit(false)
				createSynonymOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createSynonymOptionsModel).ToNot(BeNil())
				Expect(createSynonymOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(createSynonymOptionsModel.Entity).To(Equal(core.StringPtr("testString")))
				Expect(createSynonymOptionsModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(createSynonymOptionsModel.Synonym).To(Equal(core.StringPtr("testString")))
				Expect(createSynonymOptionsModel.IncludeAudit).To(Equal(core.BoolPtr(false)))
				Expect(createSynonymOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateValue successfully`, func() {
				value := "testString"
				_model, err := assistantService.NewCreateValue(value)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCreateValueOptions successfully`, func() {
				// Construct an instance of the CreateValueOptions model
				workspaceID := "testString"
				entity := "testString"
				createValueOptionsValue := "testString"
				createValueOptionsModel := assistantService.NewCreateValueOptions(workspaceID, entity, createValueOptionsValue)
				createValueOptionsModel.SetWorkspaceID("testString")
				createValueOptionsModel.SetEntity("testString")
				createValueOptionsModel.SetValue("testString")
				createValueOptionsModel.SetMetadata(make(map[string]interface{}))
				createValueOptionsModel.SetType("synonyms")
				createValueOptionsModel.SetSynonyms([]string{"testString"})
				createValueOptionsModel.SetPatterns([]string{"testString"})
				createValueOptionsModel.SetIncludeAudit(false)
				createValueOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createValueOptionsModel).ToNot(BeNil())
				Expect(createValueOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(createValueOptionsModel.Entity).To(Equal(core.StringPtr("testString")))
				Expect(createValueOptionsModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(createValueOptionsModel.Metadata).To(Equal(make(map[string]interface{})))
				Expect(createValueOptionsModel.Type).To(Equal(core.StringPtr("synonyms")))
				Expect(createValueOptionsModel.Synonyms).To(Equal([]string{"testString"}))
				Expect(createValueOptionsModel.Patterns).To(Equal([]string{"testString"}))
				Expect(createValueOptionsModel.IncludeAudit).To(Equal(core.BoolPtr(false)))
				Expect(createValueOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateWorkspaceOptions successfully`, func() {
				// Construct an instance of the ResponseGenericChannel model
				responseGenericChannelModel := new(assistantv1.ResponseGenericChannel)
				Expect(responseGenericChannelModel).ToNot(BeNil())
				responseGenericChannelModel.Channel = core.StringPtr("chat")
				Expect(responseGenericChannelModel.Channel).To(Equal(core.StringPtr("chat")))

				// Construct an instance of the DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo model
				dialogNodeOutputGenericModel := new(assistantv1.DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo)
				Expect(dialogNodeOutputGenericModel).ToNot(BeNil())
				dialogNodeOutputGenericModel.ResponseType = core.StringPtr("video")
				dialogNodeOutputGenericModel.Source = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Title = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Description = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Channels = []assistantv1.ResponseGenericChannel{*responseGenericChannelModel}
				dialogNodeOutputGenericModel.ChannelOptions = map[string]interface{}{"anyKey": "anyValue"}
				dialogNodeOutputGenericModel.AltText = core.StringPtr("testString")
				Expect(dialogNodeOutputGenericModel.ResponseType).To(Equal(core.StringPtr("video")))
				Expect(dialogNodeOutputGenericModel.Source).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeOutputGenericModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeOutputGenericModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeOutputGenericModel.Channels).To(Equal([]assistantv1.ResponseGenericChannel{*responseGenericChannelModel}))
				Expect(dialogNodeOutputGenericModel.ChannelOptions).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(dialogNodeOutputGenericModel.AltText).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the DialogNodeOutputModifiers model
				dialogNodeOutputModifiersModel := new(assistantv1.DialogNodeOutputModifiers)
				Expect(dialogNodeOutputModifiersModel).ToNot(BeNil())
				dialogNodeOutputModifiersModel.Overwrite = core.BoolPtr(true)
				Expect(dialogNodeOutputModifiersModel.Overwrite).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the DialogNodeOutput model
				dialogNodeOutputModel := new(assistantv1.DialogNodeOutput)
				Expect(dialogNodeOutputModel).ToNot(BeNil())
				dialogNodeOutputModel.Generic = []assistantv1.DialogNodeOutputGenericIntf{dialogNodeOutputGenericModel}
				dialogNodeOutputModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeOutputModel.Modifiers = dialogNodeOutputModifiersModel
				dialogNodeOutputModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(dialogNodeOutputModel.Generic).To(Equal([]assistantv1.DialogNodeOutputGenericIntf{dialogNodeOutputGenericModel}))
				Expect(dialogNodeOutputModel.Integrations).To(Equal(make(map[string]map[string]interface{})))
				Expect(dialogNodeOutputModel.Modifiers).To(Equal(dialogNodeOutputModifiersModel))
				Expect(dialogNodeOutputModel.GetProperties()).ToNot(BeEmpty())
				Expect(dialogNodeOutputModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))

				dialogNodeOutputModel.SetProperties(nil)
				Expect(dialogNodeOutputModel.GetProperties()).To(BeEmpty())

				dialogNodeOutputModelExpectedMap := make(map[string]interface{})
				dialogNodeOutputModelExpectedMap["foo"] = core.StringPtr("testString")
				dialogNodeOutputModel.SetProperties(dialogNodeOutputModelExpectedMap)
				dialogNodeOutputModelActualMap := dialogNodeOutputModel.GetProperties()
				Expect(dialogNodeOutputModelActualMap).To(Equal(dialogNodeOutputModelExpectedMap))

				// Construct an instance of the DialogNodeContext model
				dialogNodeContextModel := new(assistantv1.DialogNodeContext)
				Expect(dialogNodeContextModel).ToNot(BeNil())
				dialogNodeContextModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeContextModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(dialogNodeContextModel.Integrations).To(Equal(make(map[string]map[string]interface{})))
				Expect(dialogNodeContextModel.GetProperties()).ToNot(BeEmpty())
				Expect(dialogNodeContextModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))

				dialogNodeContextModel.SetProperties(nil)
				Expect(dialogNodeContextModel.GetProperties()).To(BeEmpty())

				dialogNodeContextModelExpectedMap := make(map[string]interface{})
				dialogNodeContextModelExpectedMap["foo"] = core.StringPtr("testString")
				dialogNodeContextModel.SetProperties(dialogNodeContextModelExpectedMap)
				dialogNodeContextModelActualMap := dialogNodeContextModel.GetProperties()
				Expect(dialogNodeContextModelActualMap).To(Equal(dialogNodeContextModelExpectedMap))

				// Construct an instance of the DialogNodeNextStep model
				dialogNodeNextStepModel := new(assistantv1.DialogNodeNextStep)
				Expect(dialogNodeNextStepModel).ToNot(BeNil())
				dialogNodeNextStepModel.Behavior = core.StringPtr("get_user_input")
				dialogNodeNextStepModel.DialogNode = core.StringPtr("testString")
				dialogNodeNextStepModel.Selector = core.StringPtr("condition")
				Expect(dialogNodeNextStepModel.Behavior).To(Equal(core.StringPtr("get_user_input")))
				Expect(dialogNodeNextStepModel.DialogNode).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeNextStepModel.Selector).To(Equal(core.StringPtr("condition")))

				// Construct an instance of the DialogNodeAction model
				dialogNodeActionModel := new(assistantv1.DialogNodeAction)
				Expect(dialogNodeActionModel).ToNot(BeNil())
				dialogNodeActionModel.Name = core.StringPtr("testString")
				dialogNodeActionModel.Type = core.StringPtr("client")
				dialogNodeActionModel.Parameters = make(map[string]interface{})
				dialogNodeActionModel.ResultVariable = core.StringPtr("testString")
				dialogNodeActionModel.Credentials = core.StringPtr("testString")
				Expect(dialogNodeActionModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeActionModel.Type).To(Equal(core.StringPtr("client")))
				Expect(dialogNodeActionModel.Parameters).To(Equal(make(map[string]interface{})))
				Expect(dialogNodeActionModel.ResultVariable).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeActionModel.Credentials).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the DialogNode model
				dialogNodeModel := new(assistantv1.DialogNode)
				Expect(dialogNodeModel).ToNot(BeNil())
				dialogNodeModel.DialogNode = core.StringPtr("testString")
				dialogNodeModel.Description = core.StringPtr("testString")
				dialogNodeModel.Conditions = core.StringPtr("testString")
				dialogNodeModel.Parent = core.StringPtr("testString")
				dialogNodeModel.PreviousSibling = core.StringPtr("testString")
				dialogNodeModel.Output = dialogNodeOutputModel
				dialogNodeModel.Context = dialogNodeContextModel
				dialogNodeModel.Metadata = make(map[string]interface{})
				dialogNodeModel.NextStep = dialogNodeNextStepModel
				dialogNodeModel.Title = core.StringPtr("testString")
				dialogNodeModel.Type = core.StringPtr("standard")
				dialogNodeModel.EventName = core.StringPtr("focus")
				dialogNodeModel.Variable = core.StringPtr("testString")
				dialogNodeModel.Actions = []assistantv1.DialogNodeAction{*dialogNodeActionModel}
				dialogNodeModel.DigressIn = core.StringPtr("not_available")
				dialogNodeModel.DigressOut = core.StringPtr("allow_returning")
				dialogNodeModel.DigressOutSlots = core.StringPtr("not_allowed")
				dialogNodeModel.UserLabel = core.StringPtr("testString")
				dialogNodeModel.DisambiguationOptOut = core.BoolPtr(false)
				Expect(dialogNodeModel.DialogNode).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeModel.Conditions).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeModel.Parent).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeModel.PreviousSibling).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeModel.Output).To(Equal(dialogNodeOutputModel))
				Expect(dialogNodeModel.Context).To(Equal(dialogNodeContextModel))
				Expect(dialogNodeModel.Metadata).To(Equal(make(map[string]interface{})))
				Expect(dialogNodeModel.NextStep).To(Equal(dialogNodeNextStepModel))
				Expect(dialogNodeModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeModel.Type).To(Equal(core.StringPtr("standard")))
				Expect(dialogNodeModel.EventName).To(Equal(core.StringPtr("focus")))
				Expect(dialogNodeModel.Variable).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeModel.Actions).To(Equal([]assistantv1.DialogNodeAction{*dialogNodeActionModel}))
				Expect(dialogNodeModel.DigressIn).To(Equal(core.StringPtr("not_available")))
				Expect(dialogNodeModel.DigressOut).To(Equal(core.StringPtr("allow_returning")))
				Expect(dialogNodeModel.DigressOutSlots).To(Equal(core.StringPtr("not_allowed")))
				Expect(dialogNodeModel.UserLabel).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeModel.DisambiguationOptOut).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the Counterexample model
				counterexampleModel := new(assistantv1.Counterexample)
				Expect(counterexampleModel).ToNot(BeNil())
				counterexampleModel.Text = core.StringPtr("testString")
				Expect(counterexampleModel.Text).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the WorkspaceSystemSettingsTooling model
				workspaceSystemSettingsToolingModel := new(assistantv1.WorkspaceSystemSettingsTooling)
				Expect(workspaceSystemSettingsToolingModel).ToNot(BeNil())
				workspaceSystemSettingsToolingModel.StoreGenericResponses = core.BoolPtr(true)
				Expect(workspaceSystemSettingsToolingModel.StoreGenericResponses).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the WorkspaceSystemSettingsDisambiguation model
				workspaceSystemSettingsDisambiguationModel := new(assistantv1.WorkspaceSystemSettingsDisambiguation)
				Expect(workspaceSystemSettingsDisambiguationModel).ToNot(BeNil())
				workspaceSystemSettingsDisambiguationModel.Prompt = core.StringPtr("testString")
				workspaceSystemSettingsDisambiguationModel.NoneOfTheAbovePrompt = core.StringPtr("testString")
				workspaceSystemSettingsDisambiguationModel.Enabled = core.BoolPtr(false)
				workspaceSystemSettingsDisambiguationModel.Sensitivity = core.StringPtr("auto")
				workspaceSystemSettingsDisambiguationModel.Randomize = core.BoolPtr(true)
				workspaceSystemSettingsDisambiguationModel.MaxSuggestions = core.Int64Ptr(int64(1))
				workspaceSystemSettingsDisambiguationModel.SuggestionTextPolicy = core.StringPtr("testString")
				Expect(workspaceSystemSettingsDisambiguationModel.Prompt).To(Equal(core.StringPtr("testString")))
				Expect(workspaceSystemSettingsDisambiguationModel.NoneOfTheAbovePrompt).To(Equal(core.StringPtr("testString")))
				Expect(workspaceSystemSettingsDisambiguationModel.Enabled).To(Equal(core.BoolPtr(false)))
				Expect(workspaceSystemSettingsDisambiguationModel.Sensitivity).To(Equal(core.StringPtr("auto")))
				Expect(workspaceSystemSettingsDisambiguationModel.Randomize).To(Equal(core.BoolPtr(true)))
				Expect(workspaceSystemSettingsDisambiguationModel.MaxSuggestions).To(Equal(core.Int64Ptr(int64(1))))
				Expect(workspaceSystemSettingsDisambiguationModel.SuggestionTextPolicy).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the WorkspaceSystemSettingsSystemEntities model
				workspaceSystemSettingsSystemEntitiesModel := new(assistantv1.WorkspaceSystemSettingsSystemEntities)
				Expect(workspaceSystemSettingsSystemEntitiesModel).ToNot(BeNil())
				workspaceSystemSettingsSystemEntitiesModel.Enabled = core.BoolPtr(false)
				Expect(workspaceSystemSettingsSystemEntitiesModel.Enabled).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the WorkspaceSystemSettingsOffTopic model
				workspaceSystemSettingsOffTopicModel := new(assistantv1.WorkspaceSystemSettingsOffTopic)
				Expect(workspaceSystemSettingsOffTopicModel).ToNot(BeNil())
				workspaceSystemSettingsOffTopicModel.Enabled = core.BoolPtr(false)
				Expect(workspaceSystemSettingsOffTopicModel.Enabled).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the WorkspaceSystemSettings model
				workspaceSystemSettingsModel := new(assistantv1.WorkspaceSystemSettings)
				Expect(workspaceSystemSettingsModel).ToNot(BeNil())
				workspaceSystemSettingsModel.Tooling = workspaceSystemSettingsToolingModel
				workspaceSystemSettingsModel.Disambiguation = workspaceSystemSettingsDisambiguationModel
				workspaceSystemSettingsModel.HumanAgentAssist = make(map[string]interface{})
				workspaceSystemSettingsModel.SpellingSuggestions = core.BoolPtr(false)
				workspaceSystemSettingsModel.SpellingAutoCorrect = core.BoolPtr(false)
				workspaceSystemSettingsModel.SystemEntities = workspaceSystemSettingsSystemEntitiesModel
				workspaceSystemSettingsModel.OffTopic = workspaceSystemSettingsOffTopicModel
				workspaceSystemSettingsModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(workspaceSystemSettingsModel.Tooling).To(Equal(workspaceSystemSettingsToolingModel))
				Expect(workspaceSystemSettingsModel.Disambiguation).To(Equal(workspaceSystemSettingsDisambiguationModel))
				Expect(workspaceSystemSettingsModel.HumanAgentAssist).To(Equal(make(map[string]interface{})))
				Expect(workspaceSystemSettingsModel.SpellingSuggestions).To(Equal(core.BoolPtr(false)))
				Expect(workspaceSystemSettingsModel.SpellingAutoCorrect).To(Equal(core.BoolPtr(false)))
				Expect(workspaceSystemSettingsModel.SystemEntities).To(Equal(workspaceSystemSettingsSystemEntitiesModel))
				Expect(workspaceSystemSettingsModel.OffTopic).To(Equal(workspaceSystemSettingsOffTopicModel))
				Expect(workspaceSystemSettingsModel.GetProperties()).ToNot(BeEmpty())
				Expect(workspaceSystemSettingsModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))

				workspaceSystemSettingsModel.SetProperties(nil)
				Expect(workspaceSystemSettingsModel.GetProperties()).To(BeEmpty())

				workspaceSystemSettingsModelExpectedMap := make(map[string]interface{})
				workspaceSystemSettingsModelExpectedMap["foo"] = core.StringPtr("testString")
				workspaceSystemSettingsModel.SetProperties(workspaceSystemSettingsModelExpectedMap)
				workspaceSystemSettingsModelActualMap := workspaceSystemSettingsModel.GetProperties()
				Expect(workspaceSystemSettingsModelActualMap).To(Equal(workspaceSystemSettingsModelExpectedMap))

				// Construct an instance of the WebhookHeader model
				webhookHeaderModel := new(assistantv1.WebhookHeader)
				Expect(webhookHeaderModel).ToNot(BeNil())
				webhookHeaderModel.Name = core.StringPtr("testString")
				webhookHeaderModel.Value = core.StringPtr("testString")
				Expect(webhookHeaderModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(webhookHeaderModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Webhook model
				webhookModel := new(assistantv1.Webhook)
				Expect(webhookModel).ToNot(BeNil())
				webhookModel.URL = core.StringPtr("testString")
				webhookModel.Name = core.StringPtr("testString")
				webhookModel.HeadersVar = []assistantv1.WebhookHeader{*webhookHeaderModel}
				Expect(webhookModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(webhookModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(webhookModel.HeadersVar).To(Equal([]assistantv1.WebhookHeader{*webhookHeaderModel}))

				// Construct an instance of the Mention model
				mentionModel := new(assistantv1.Mention)
				Expect(mentionModel).ToNot(BeNil())
				mentionModel.Entity = core.StringPtr("testString")
				mentionModel.Location = []int64{int64(38)}
				Expect(mentionModel.Entity).To(Equal(core.StringPtr("testString")))
				Expect(mentionModel.Location).To(Equal([]int64{int64(38)}))

				// Construct an instance of the Example model
				exampleModel := new(assistantv1.Example)
				Expect(exampleModel).ToNot(BeNil())
				exampleModel.Text = core.StringPtr("testString")
				exampleModel.Mentions = []assistantv1.Mention{*mentionModel}
				Expect(exampleModel.Text).To(Equal(core.StringPtr("testString")))
				Expect(exampleModel.Mentions).To(Equal([]assistantv1.Mention{*mentionModel}))

				// Construct an instance of the CreateIntent model
				createIntentModel := new(assistantv1.CreateIntent)
				Expect(createIntentModel).ToNot(BeNil())
				createIntentModel.Intent = core.StringPtr("testString")
				createIntentModel.Description = core.StringPtr("testString")
				createIntentModel.Examples = []assistantv1.Example{*exampleModel}
				Expect(createIntentModel.Intent).To(Equal(core.StringPtr("testString")))
				Expect(createIntentModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createIntentModel.Examples).To(Equal([]assistantv1.Example{*exampleModel}))

				// Construct an instance of the CreateValue model
				createValueModel := new(assistantv1.CreateValue)
				Expect(createValueModel).ToNot(BeNil())
				createValueModel.Value = core.StringPtr("testString")
				createValueModel.Metadata = make(map[string]interface{})
				createValueModel.Type = core.StringPtr("synonyms")
				createValueModel.Synonyms = []string{"testString"}
				createValueModel.Patterns = []string{"testString"}
				Expect(createValueModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(createValueModel.Metadata).To(Equal(make(map[string]interface{})))
				Expect(createValueModel.Type).To(Equal(core.StringPtr("synonyms")))
				Expect(createValueModel.Synonyms).To(Equal([]string{"testString"}))
				Expect(createValueModel.Patterns).To(Equal([]string{"testString"}))

				// Construct an instance of the CreateEntity model
				createEntityModel := new(assistantv1.CreateEntity)
				Expect(createEntityModel).ToNot(BeNil())
				createEntityModel.Entity = core.StringPtr("testString")
				createEntityModel.Description = core.StringPtr("testString")
				createEntityModel.Metadata = make(map[string]interface{})
				createEntityModel.FuzzyMatch = core.BoolPtr(true)
				createEntityModel.Values = []assistantv1.CreateValue{*createValueModel}
				Expect(createEntityModel.Entity).To(Equal(core.StringPtr("testString")))
				Expect(createEntityModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createEntityModel.Metadata).To(Equal(make(map[string]interface{})))
				Expect(createEntityModel.FuzzyMatch).To(Equal(core.BoolPtr(true)))
				Expect(createEntityModel.Values).To(Equal([]assistantv1.CreateValue{*createValueModel}))

				// Construct an instance of the CreateWorkspaceOptions model
				createWorkspaceOptionsModel := assistantService.NewCreateWorkspaceOptions()
				createWorkspaceOptionsModel.SetName("testString")
				createWorkspaceOptionsModel.SetDescription("testString")
				createWorkspaceOptionsModel.SetLanguage("testString")
				createWorkspaceOptionsModel.SetDialogNodes([]assistantv1.DialogNode{*dialogNodeModel})
				createWorkspaceOptionsModel.SetCounterexamples([]assistantv1.Counterexample{*counterexampleModel})
				createWorkspaceOptionsModel.SetMetadata(make(map[string]interface{}))
				createWorkspaceOptionsModel.SetLearningOptOut(false)
				createWorkspaceOptionsModel.SetSystemSettings(workspaceSystemSettingsModel)
				createWorkspaceOptionsModel.SetWebhooks([]assistantv1.Webhook{*webhookModel})
				createWorkspaceOptionsModel.SetIntents([]assistantv1.CreateIntent{*createIntentModel})
				createWorkspaceOptionsModel.SetEntities([]assistantv1.CreateEntity{*createEntityModel})
				createWorkspaceOptionsModel.SetIncludeAudit(false)
				createWorkspaceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createWorkspaceOptionsModel).ToNot(BeNil())
				Expect(createWorkspaceOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createWorkspaceOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createWorkspaceOptionsModel.Language).To(Equal(core.StringPtr("testString")))
				Expect(createWorkspaceOptionsModel.DialogNodes).To(Equal([]assistantv1.DialogNode{*dialogNodeModel}))
				Expect(createWorkspaceOptionsModel.Counterexamples).To(Equal([]assistantv1.Counterexample{*counterexampleModel}))
				Expect(createWorkspaceOptionsModel.Metadata).To(Equal(make(map[string]interface{})))
				Expect(createWorkspaceOptionsModel.LearningOptOut).To(Equal(core.BoolPtr(false)))
				Expect(createWorkspaceOptionsModel.SystemSettings).To(Equal(workspaceSystemSettingsModel))
				Expect(createWorkspaceOptionsModel.Webhooks).To(Equal([]assistantv1.Webhook{*webhookModel}))
				Expect(createWorkspaceOptionsModel.Intents).To(Equal([]assistantv1.CreateIntent{*createIntentModel}))
				Expect(createWorkspaceOptionsModel.Entities).To(Equal([]assistantv1.CreateEntity{*createEntityModel}))
				Expect(createWorkspaceOptionsModel.IncludeAudit).To(Equal(core.BoolPtr(false)))
				Expect(createWorkspaceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteCounterexampleOptions successfully`, func() {
				// Construct an instance of the DeleteCounterexampleOptions model
				workspaceID := "testString"
				text := "testString"
				deleteCounterexampleOptionsModel := assistantService.NewDeleteCounterexampleOptions(workspaceID, text)
				deleteCounterexampleOptionsModel.SetWorkspaceID("testString")
				deleteCounterexampleOptionsModel.SetText("testString")
				deleteCounterexampleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteCounterexampleOptionsModel).ToNot(BeNil())
				Expect(deleteCounterexampleOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCounterexampleOptionsModel.Text).To(Equal(core.StringPtr("testString")))
				Expect(deleteCounterexampleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteDialogNodeOptions successfully`, func() {
				// Construct an instance of the DeleteDialogNodeOptions model
				workspaceID := "testString"
				dialogNode := "testString"
				deleteDialogNodeOptionsModel := assistantService.NewDeleteDialogNodeOptions(workspaceID, dialogNode)
				deleteDialogNodeOptionsModel.SetWorkspaceID("testString")
				deleteDialogNodeOptionsModel.SetDialogNode("testString")
				deleteDialogNodeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteDialogNodeOptionsModel).ToNot(BeNil())
				Expect(deleteDialogNodeOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDialogNodeOptionsModel.DialogNode).To(Equal(core.StringPtr("testString")))
				Expect(deleteDialogNodeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteEntityOptions successfully`, func() {
				// Construct an instance of the DeleteEntityOptions model
				workspaceID := "testString"
				entity := "testString"
				deleteEntityOptionsModel := assistantService.NewDeleteEntityOptions(workspaceID, entity)
				deleteEntityOptionsModel.SetWorkspaceID("testString")
				deleteEntityOptionsModel.SetEntity("testString")
				deleteEntityOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteEntityOptionsModel).ToNot(BeNil())
				Expect(deleteEntityOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteEntityOptionsModel.Entity).To(Equal(core.StringPtr("testString")))
				Expect(deleteEntityOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteExampleOptions successfully`, func() {
				// Construct an instance of the DeleteExampleOptions model
				workspaceID := "testString"
				intent := "testString"
				text := "testString"
				deleteExampleOptionsModel := assistantService.NewDeleteExampleOptions(workspaceID, intent, text)
				deleteExampleOptionsModel.SetWorkspaceID("testString")
				deleteExampleOptionsModel.SetIntent("testString")
				deleteExampleOptionsModel.SetText("testString")
				deleteExampleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteExampleOptionsModel).ToNot(BeNil())
				Expect(deleteExampleOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteExampleOptionsModel.Intent).To(Equal(core.StringPtr("testString")))
				Expect(deleteExampleOptionsModel.Text).To(Equal(core.StringPtr("testString")))
				Expect(deleteExampleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteIntentOptions successfully`, func() {
				// Construct an instance of the DeleteIntentOptions model
				workspaceID := "testString"
				intent := "testString"
				deleteIntentOptionsModel := assistantService.NewDeleteIntentOptions(workspaceID, intent)
				deleteIntentOptionsModel.SetWorkspaceID("testString")
				deleteIntentOptionsModel.SetIntent("testString")
				deleteIntentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteIntentOptionsModel).ToNot(BeNil())
				Expect(deleteIntentOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteIntentOptionsModel.Intent).To(Equal(core.StringPtr("testString")))
				Expect(deleteIntentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteSynonymOptions successfully`, func() {
				// Construct an instance of the DeleteSynonymOptions model
				workspaceID := "testString"
				entity := "testString"
				value := "testString"
				synonym := "testString"
				deleteSynonymOptionsModel := assistantService.NewDeleteSynonymOptions(workspaceID, entity, value, synonym)
				deleteSynonymOptionsModel.SetWorkspaceID("testString")
				deleteSynonymOptionsModel.SetEntity("testString")
				deleteSynonymOptionsModel.SetValue("testString")
				deleteSynonymOptionsModel.SetSynonym("testString")
				deleteSynonymOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteSynonymOptionsModel).ToNot(BeNil())
				Expect(deleteSynonymOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteSynonymOptionsModel.Entity).To(Equal(core.StringPtr("testString")))
				Expect(deleteSynonymOptionsModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(deleteSynonymOptionsModel.Synonym).To(Equal(core.StringPtr("testString")))
				Expect(deleteSynonymOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			It(`Invoke NewDeleteValueOptions successfully`, func() {
				// Construct an instance of the DeleteValueOptions model
				workspaceID := "testString"
				entity := "testString"
				value := "testString"
				deleteValueOptionsModel := assistantService.NewDeleteValueOptions(workspaceID, entity, value)
				deleteValueOptionsModel.SetWorkspaceID("testString")
				deleteValueOptionsModel.SetEntity("testString")
				deleteValueOptionsModel.SetValue("testString")
				deleteValueOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteValueOptionsModel).ToNot(BeNil())
				Expect(deleteValueOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteValueOptionsModel.Entity).To(Equal(core.StringPtr("testString")))
				Expect(deleteValueOptionsModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(deleteValueOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteWorkspaceOptions successfully`, func() {
				// Construct an instance of the DeleteWorkspaceOptions model
				workspaceID := "testString"
				deleteWorkspaceOptionsModel := assistantService.NewDeleteWorkspaceOptions(workspaceID)
				deleteWorkspaceOptionsModel.SetWorkspaceID("testString")
				deleteWorkspaceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteWorkspaceOptionsModel).ToNot(BeNil())
				Expect(deleteWorkspaceOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteWorkspaceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDialogNode successfully`, func() {
				dialogNode := "testString"
				_model, err := assistantService.NewDialogNode(dialogNode)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDialogNodeAction successfully`, func() {
				name := "testString"
				resultVariable := "testString"
				_model, err := assistantService.NewDialogNodeAction(name, resultVariable)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDialogNodeNextStep successfully`, func() {
				behavior := "get_user_input"
				_model, err := assistantService.NewDialogNodeNextStep(behavior)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDialogNodeOutputOptionsElement successfully`, func() {
				label := "testString"
				var value *assistantv1.DialogNodeOutputOptionsElementValue = nil
				_, err := assistantService.NewDialogNodeOutputOptionsElement(label, value)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewDialogSuggestion successfully`, func() {
				label := "testString"
				var value *assistantv1.DialogSuggestionValue = nil
				_, err := assistantService.NewDialogSuggestion(label, value)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewExample successfully`, func() {
				text := "testString"
				_model, err := assistantService.NewExample(text)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewGetCounterexampleOptions successfully`, func() {
				// Construct an instance of the GetCounterexampleOptions model
				workspaceID := "testString"
				text := "testString"
				getCounterexampleOptionsModel := assistantService.NewGetCounterexampleOptions(workspaceID, text)
				getCounterexampleOptionsModel.SetWorkspaceID("testString")
				getCounterexampleOptionsModel.SetText("testString")
				getCounterexampleOptionsModel.SetIncludeAudit(false)
				getCounterexampleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCounterexampleOptionsModel).ToNot(BeNil())
				Expect(getCounterexampleOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(getCounterexampleOptionsModel.Text).To(Equal(core.StringPtr("testString")))
				Expect(getCounterexampleOptionsModel.IncludeAudit).To(Equal(core.BoolPtr(false)))
				Expect(getCounterexampleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDialogNodeOptions successfully`, func() {
				// Construct an instance of the GetDialogNodeOptions model
				workspaceID := "testString"
				dialogNode := "testString"
				getDialogNodeOptionsModel := assistantService.NewGetDialogNodeOptions(workspaceID, dialogNode)
				getDialogNodeOptionsModel.SetWorkspaceID("testString")
				getDialogNodeOptionsModel.SetDialogNode("testString")
				getDialogNodeOptionsModel.SetIncludeAudit(false)
				getDialogNodeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDialogNodeOptionsModel).ToNot(BeNil())
				Expect(getDialogNodeOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(getDialogNodeOptionsModel.DialogNode).To(Equal(core.StringPtr("testString")))
				Expect(getDialogNodeOptionsModel.IncludeAudit).To(Equal(core.BoolPtr(false)))
				Expect(getDialogNodeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetEntityOptions successfully`, func() {
				// Construct an instance of the GetEntityOptions model
				workspaceID := "testString"
				entity := "testString"
				getEntityOptionsModel := assistantService.NewGetEntityOptions(workspaceID, entity)
				getEntityOptionsModel.SetWorkspaceID("testString")
				getEntityOptionsModel.SetEntity("testString")
				getEntityOptionsModel.SetExport(false)
				getEntityOptionsModel.SetIncludeAudit(false)
				getEntityOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getEntityOptionsModel).ToNot(BeNil())
				Expect(getEntityOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(getEntityOptionsModel.Entity).To(Equal(core.StringPtr("testString")))
				Expect(getEntityOptionsModel.Export).To(Equal(core.BoolPtr(false)))
				Expect(getEntityOptionsModel.IncludeAudit).To(Equal(core.BoolPtr(false)))
				Expect(getEntityOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetExampleOptions successfully`, func() {
				// Construct an instance of the GetExampleOptions model
				workspaceID := "testString"
				intent := "testString"
				text := "testString"
				getExampleOptionsModel := assistantService.NewGetExampleOptions(workspaceID, intent, text)
				getExampleOptionsModel.SetWorkspaceID("testString")
				getExampleOptionsModel.SetIntent("testString")
				getExampleOptionsModel.SetText("testString")
				getExampleOptionsModel.SetIncludeAudit(false)
				getExampleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getExampleOptionsModel).ToNot(BeNil())
				Expect(getExampleOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(getExampleOptionsModel.Intent).To(Equal(core.StringPtr("testString")))
				Expect(getExampleOptionsModel.Text).To(Equal(core.StringPtr("testString")))
				Expect(getExampleOptionsModel.IncludeAudit).To(Equal(core.BoolPtr(false)))
				Expect(getExampleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetIntentOptions successfully`, func() {
				// Construct an instance of the GetIntentOptions model
				workspaceID := "testString"
				intent := "testString"
				getIntentOptionsModel := assistantService.NewGetIntentOptions(workspaceID, intent)
				getIntentOptionsModel.SetWorkspaceID("testString")
				getIntentOptionsModel.SetIntent("testString")
				getIntentOptionsModel.SetExport(false)
				getIntentOptionsModel.SetIncludeAudit(false)
				getIntentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getIntentOptionsModel).ToNot(BeNil())
				Expect(getIntentOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(getIntentOptionsModel.Intent).To(Equal(core.StringPtr("testString")))
				Expect(getIntentOptionsModel.Export).To(Equal(core.BoolPtr(false)))
				Expect(getIntentOptionsModel.IncludeAudit).To(Equal(core.BoolPtr(false)))
				Expect(getIntentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSynonymOptions successfully`, func() {
				// Construct an instance of the GetSynonymOptions model
				workspaceID := "testString"
				entity := "testString"
				value := "testString"
				synonym := "testString"
				getSynonymOptionsModel := assistantService.NewGetSynonymOptions(workspaceID, entity, value, synonym)
				getSynonymOptionsModel.SetWorkspaceID("testString")
				getSynonymOptionsModel.SetEntity("testString")
				getSynonymOptionsModel.SetValue("testString")
				getSynonymOptionsModel.SetSynonym("testString")
				getSynonymOptionsModel.SetIncludeAudit(false)
				getSynonymOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSynonymOptionsModel).ToNot(BeNil())
				Expect(getSynonymOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(getSynonymOptionsModel.Entity).To(Equal(core.StringPtr("testString")))
				Expect(getSynonymOptionsModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(getSynonymOptionsModel.Synonym).To(Equal(core.StringPtr("testString")))
				Expect(getSynonymOptionsModel.IncludeAudit).To(Equal(core.BoolPtr(false)))
				Expect(getSynonymOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetValueOptions successfully`, func() {
				// Construct an instance of the GetValueOptions model
				workspaceID := "testString"
				entity := "testString"
				value := "testString"
				getValueOptionsModel := assistantService.NewGetValueOptions(workspaceID, entity, value)
				getValueOptionsModel.SetWorkspaceID("testString")
				getValueOptionsModel.SetEntity("testString")
				getValueOptionsModel.SetValue("testString")
				getValueOptionsModel.SetExport(false)
				getValueOptionsModel.SetIncludeAudit(false)
				getValueOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getValueOptionsModel).ToNot(BeNil())
				Expect(getValueOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(getValueOptionsModel.Entity).To(Equal(core.StringPtr("testString")))
				Expect(getValueOptionsModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(getValueOptionsModel.Export).To(Equal(core.BoolPtr(false)))
				Expect(getValueOptionsModel.IncludeAudit).To(Equal(core.BoolPtr(false)))
				Expect(getValueOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetWorkspaceOptions successfully`, func() {
				// Construct an instance of the GetWorkspaceOptions model
				workspaceID := "testString"
				getWorkspaceOptionsModel := assistantService.NewGetWorkspaceOptions(workspaceID)
				getWorkspaceOptionsModel.SetWorkspaceID("testString")
				getWorkspaceOptionsModel.SetExport(false)
				getWorkspaceOptionsModel.SetIncludeAudit(false)
				getWorkspaceOptionsModel.SetSort("stable")
				getWorkspaceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getWorkspaceOptionsModel).ToNot(BeNil())
				Expect(getWorkspaceOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(getWorkspaceOptionsModel.Export).To(Equal(core.BoolPtr(false)))
				Expect(getWorkspaceOptionsModel.IncludeAudit).To(Equal(core.BoolPtr(false)))
				Expect(getWorkspaceOptionsModel.Sort).To(Equal(core.StringPtr("stable")))
				Expect(getWorkspaceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAllLogsOptions successfully`, func() {
				// Construct an instance of the ListAllLogsOptions model
				filter := "testString"
				listAllLogsOptionsModel := assistantService.NewListAllLogsOptions(filter)
				listAllLogsOptionsModel.SetFilter("testString")
				listAllLogsOptionsModel.SetSort("testString")
				listAllLogsOptionsModel.SetPageLimit(int64(38))
				listAllLogsOptionsModel.SetCursor("testString")
				listAllLogsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAllLogsOptionsModel).ToNot(BeNil())
				Expect(listAllLogsOptionsModel.Filter).To(Equal(core.StringPtr("testString")))
				Expect(listAllLogsOptionsModel.Sort).To(Equal(core.StringPtr("testString")))
				Expect(listAllLogsOptionsModel.PageLimit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listAllLogsOptionsModel.Cursor).To(Equal(core.StringPtr("testString")))
				Expect(listAllLogsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListCounterexamplesOptions successfully`, func() {
				// Construct an instance of the ListCounterexamplesOptions model
				workspaceID := "testString"
				listCounterexamplesOptionsModel := assistantService.NewListCounterexamplesOptions(workspaceID)
				listCounterexamplesOptionsModel.SetWorkspaceID("testString")
				listCounterexamplesOptionsModel.SetPageLimit(int64(38))
				listCounterexamplesOptionsModel.SetIncludeCount(false)
				listCounterexamplesOptionsModel.SetSort("text")
				listCounterexamplesOptionsModel.SetCursor("testString")
				listCounterexamplesOptionsModel.SetIncludeAudit(false)
				listCounterexamplesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listCounterexamplesOptionsModel).ToNot(BeNil())
				Expect(listCounterexamplesOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(listCounterexamplesOptionsModel.PageLimit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listCounterexamplesOptionsModel.IncludeCount).To(Equal(core.BoolPtr(false)))
				Expect(listCounterexamplesOptionsModel.Sort).To(Equal(core.StringPtr("text")))
				Expect(listCounterexamplesOptionsModel.Cursor).To(Equal(core.StringPtr("testString")))
				Expect(listCounterexamplesOptionsModel.IncludeAudit).To(Equal(core.BoolPtr(false)))
				Expect(listCounterexamplesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListDialogNodesOptions successfully`, func() {
				// Construct an instance of the ListDialogNodesOptions model
				workspaceID := "testString"
				listDialogNodesOptionsModel := assistantService.NewListDialogNodesOptions(workspaceID)
				listDialogNodesOptionsModel.SetWorkspaceID("testString")
				listDialogNodesOptionsModel.SetPageLimit(int64(38))
				listDialogNodesOptionsModel.SetIncludeCount(false)
				listDialogNodesOptionsModel.SetSort("dialog_node")
				listDialogNodesOptionsModel.SetCursor("testString")
				listDialogNodesOptionsModel.SetIncludeAudit(false)
				listDialogNodesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listDialogNodesOptionsModel).ToNot(BeNil())
				Expect(listDialogNodesOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(listDialogNodesOptionsModel.PageLimit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listDialogNodesOptionsModel.IncludeCount).To(Equal(core.BoolPtr(false)))
				Expect(listDialogNodesOptionsModel.Sort).To(Equal(core.StringPtr("dialog_node")))
				Expect(listDialogNodesOptionsModel.Cursor).To(Equal(core.StringPtr("testString")))
				Expect(listDialogNodesOptionsModel.IncludeAudit).To(Equal(core.BoolPtr(false)))
				Expect(listDialogNodesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListEntitiesOptions successfully`, func() {
				// Construct an instance of the ListEntitiesOptions model
				workspaceID := "testString"
				listEntitiesOptionsModel := assistantService.NewListEntitiesOptions(workspaceID)
				listEntitiesOptionsModel.SetWorkspaceID("testString")
				listEntitiesOptionsModel.SetExport(false)
				listEntitiesOptionsModel.SetPageLimit(int64(38))
				listEntitiesOptionsModel.SetIncludeCount(false)
				listEntitiesOptionsModel.SetSort("entity")
				listEntitiesOptionsModel.SetCursor("testString")
				listEntitiesOptionsModel.SetIncludeAudit(false)
				listEntitiesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listEntitiesOptionsModel).ToNot(BeNil())
				Expect(listEntitiesOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(listEntitiesOptionsModel.Export).To(Equal(core.BoolPtr(false)))
				Expect(listEntitiesOptionsModel.PageLimit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listEntitiesOptionsModel.IncludeCount).To(Equal(core.BoolPtr(false)))
				Expect(listEntitiesOptionsModel.Sort).To(Equal(core.StringPtr("entity")))
				Expect(listEntitiesOptionsModel.Cursor).To(Equal(core.StringPtr("testString")))
				Expect(listEntitiesOptionsModel.IncludeAudit).To(Equal(core.BoolPtr(false)))
				Expect(listEntitiesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListExamplesOptions successfully`, func() {
				// Construct an instance of the ListExamplesOptions model
				workspaceID := "testString"
				intent := "testString"
				listExamplesOptionsModel := assistantService.NewListExamplesOptions(workspaceID, intent)
				listExamplesOptionsModel.SetWorkspaceID("testString")
				listExamplesOptionsModel.SetIntent("testString")
				listExamplesOptionsModel.SetPageLimit(int64(38))
				listExamplesOptionsModel.SetIncludeCount(false)
				listExamplesOptionsModel.SetSort("text")
				listExamplesOptionsModel.SetCursor("testString")
				listExamplesOptionsModel.SetIncludeAudit(false)
				listExamplesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listExamplesOptionsModel).ToNot(BeNil())
				Expect(listExamplesOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(listExamplesOptionsModel.Intent).To(Equal(core.StringPtr("testString")))
				Expect(listExamplesOptionsModel.PageLimit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listExamplesOptionsModel.IncludeCount).To(Equal(core.BoolPtr(false)))
				Expect(listExamplesOptionsModel.Sort).To(Equal(core.StringPtr("text")))
				Expect(listExamplesOptionsModel.Cursor).To(Equal(core.StringPtr("testString")))
				Expect(listExamplesOptionsModel.IncludeAudit).To(Equal(core.BoolPtr(false)))
				Expect(listExamplesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListIntentsOptions successfully`, func() {
				// Construct an instance of the ListIntentsOptions model
				workspaceID := "testString"
				listIntentsOptionsModel := assistantService.NewListIntentsOptions(workspaceID)
				listIntentsOptionsModel.SetWorkspaceID("testString")
				listIntentsOptionsModel.SetExport(false)
				listIntentsOptionsModel.SetPageLimit(int64(38))
				listIntentsOptionsModel.SetIncludeCount(false)
				listIntentsOptionsModel.SetSort("intent")
				listIntentsOptionsModel.SetCursor("testString")
				listIntentsOptionsModel.SetIncludeAudit(false)
				listIntentsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listIntentsOptionsModel).ToNot(BeNil())
				Expect(listIntentsOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(listIntentsOptionsModel.Export).To(Equal(core.BoolPtr(false)))
				Expect(listIntentsOptionsModel.PageLimit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listIntentsOptionsModel.IncludeCount).To(Equal(core.BoolPtr(false)))
				Expect(listIntentsOptionsModel.Sort).To(Equal(core.StringPtr("intent")))
				Expect(listIntentsOptionsModel.Cursor).To(Equal(core.StringPtr("testString")))
				Expect(listIntentsOptionsModel.IncludeAudit).To(Equal(core.BoolPtr(false)))
				Expect(listIntentsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListLogsOptions successfully`, func() {
				// Construct an instance of the ListLogsOptions model
				workspaceID := "testString"
				listLogsOptionsModel := assistantService.NewListLogsOptions(workspaceID)
				listLogsOptionsModel.SetWorkspaceID("testString")
				listLogsOptionsModel.SetSort("testString")
				listLogsOptionsModel.SetFilter("testString")
				listLogsOptionsModel.SetPageLimit(int64(38))
				listLogsOptionsModel.SetCursor("testString")
				listLogsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listLogsOptionsModel).ToNot(BeNil())
				Expect(listLogsOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(listLogsOptionsModel.Sort).To(Equal(core.StringPtr("testString")))
				Expect(listLogsOptionsModel.Filter).To(Equal(core.StringPtr("testString")))
				Expect(listLogsOptionsModel.PageLimit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listLogsOptionsModel.Cursor).To(Equal(core.StringPtr("testString")))
				Expect(listLogsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListMentionsOptions successfully`, func() {
				// Construct an instance of the ListMentionsOptions model
				workspaceID := "testString"
				entity := "testString"
				listMentionsOptionsModel := assistantService.NewListMentionsOptions(workspaceID, entity)
				listMentionsOptionsModel.SetWorkspaceID("testString")
				listMentionsOptionsModel.SetEntity("testString")
				listMentionsOptionsModel.SetExport(false)
				listMentionsOptionsModel.SetIncludeAudit(false)
				listMentionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listMentionsOptionsModel).ToNot(BeNil())
				Expect(listMentionsOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(listMentionsOptionsModel.Entity).To(Equal(core.StringPtr("testString")))
				Expect(listMentionsOptionsModel.Export).To(Equal(core.BoolPtr(false)))
				Expect(listMentionsOptionsModel.IncludeAudit).To(Equal(core.BoolPtr(false)))
				Expect(listMentionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListSynonymsOptions successfully`, func() {
				// Construct an instance of the ListSynonymsOptions model
				workspaceID := "testString"
				entity := "testString"
				value := "testString"
				listSynonymsOptionsModel := assistantService.NewListSynonymsOptions(workspaceID, entity, value)
				listSynonymsOptionsModel.SetWorkspaceID("testString")
				listSynonymsOptionsModel.SetEntity("testString")
				listSynonymsOptionsModel.SetValue("testString")
				listSynonymsOptionsModel.SetPageLimit(int64(38))
				listSynonymsOptionsModel.SetIncludeCount(false)
				listSynonymsOptionsModel.SetSort("synonym")
				listSynonymsOptionsModel.SetCursor("testString")
				listSynonymsOptionsModel.SetIncludeAudit(false)
				listSynonymsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listSynonymsOptionsModel).ToNot(BeNil())
				Expect(listSynonymsOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(listSynonymsOptionsModel.Entity).To(Equal(core.StringPtr("testString")))
				Expect(listSynonymsOptionsModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(listSynonymsOptionsModel.PageLimit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listSynonymsOptionsModel.IncludeCount).To(Equal(core.BoolPtr(false)))
				Expect(listSynonymsOptionsModel.Sort).To(Equal(core.StringPtr("synonym")))
				Expect(listSynonymsOptionsModel.Cursor).To(Equal(core.StringPtr("testString")))
				Expect(listSynonymsOptionsModel.IncludeAudit).To(Equal(core.BoolPtr(false)))
				Expect(listSynonymsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListValuesOptions successfully`, func() {
				// Construct an instance of the ListValuesOptions model
				workspaceID := "testString"
				entity := "testString"
				listValuesOptionsModel := assistantService.NewListValuesOptions(workspaceID, entity)
				listValuesOptionsModel.SetWorkspaceID("testString")
				listValuesOptionsModel.SetEntity("testString")
				listValuesOptionsModel.SetExport(false)
				listValuesOptionsModel.SetPageLimit(int64(38))
				listValuesOptionsModel.SetIncludeCount(false)
				listValuesOptionsModel.SetSort("value")
				listValuesOptionsModel.SetCursor("testString")
				listValuesOptionsModel.SetIncludeAudit(false)
				listValuesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listValuesOptionsModel).ToNot(BeNil())
				Expect(listValuesOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(listValuesOptionsModel.Entity).To(Equal(core.StringPtr("testString")))
				Expect(listValuesOptionsModel.Export).To(Equal(core.BoolPtr(false)))
				Expect(listValuesOptionsModel.PageLimit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listValuesOptionsModel.IncludeCount).To(Equal(core.BoolPtr(false)))
				Expect(listValuesOptionsModel.Sort).To(Equal(core.StringPtr("value")))
				Expect(listValuesOptionsModel.Cursor).To(Equal(core.StringPtr("testString")))
				Expect(listValuesOptionsModel.IncludeAudit).To(Equal(core.BoolPtr(false)))
				Expect(listValuesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListWorkspacesOptions successfully`, func() {
				// Construct an instance of the ListWorkspacesOptions model
				listWorkspacesOptionsModel := assistantService.NewListWorkspacesOptions()
				listWorkspacesOptionsModel.SetPageLimit(int64(38))
				listWorkspacesOptionsModel.SetIncludeCount(false)
				listWorkspacesOptionsModel.SetSort("name")
				listWorkspacesOptionsModel.SetCursor("testString")
				listWorkspacesOptionsModel.SetIncludeAudit(false)
				listWorkspacesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listWorkspacesOptionsModel).ToNot(BeNil())
				Expect(listWorkspacesOptionsModel.PageLimit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listWorkspacesOptionsModel.IncludeCount).To(Equal(core.BoolPtr(false)))
				Expect(listWorkspacesOptionsModel.Sort).To(Equal(core.StringPtr("name")))
				Expect(listWorkspacesOptionsModel.Cursor).To(Equal(core.StringPtr("testString")))
				Expect(listWorkspacesOptionsModel.IncludeAudit).To(Equal(core.BoolPtr(false)))
				Expect(listWorkspacesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewLogMessage successfully`, func() {
				level := "info"
				msg := "testString"
				code := "testString"
				_model, err := assistantService.NewLogMessage(level, msg, code)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewMention successfully`, func() {
				entity := "testString"
				location := []int64{int64(38)}
				_model, err := assistantService.NewMention(entity, location)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewMessageOptions successfully`, func() {
				// Construct an instance of the MessageInput model
				messageInputModel := new(assistantv1.MessageInput)
				Expect(messageInputModel).ToNot(BeNil())
				messageInputModel.Text = core.StringPtr("testString")
				messageInputModel.SpellingSuggestions = core.BoolPtr(false)
				messageInputModel.SpellingAutoCorrect = core.BoolPtr(false)
				messageInputModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(messageInputModel.Text).To(Equal(core.StringPtr("testString")))
				Expect(messageInputModel.SpellingSuggestions).To(Equal(core.BoolPtr(false)))
				Expect(messageInputModel.SpellingAutoCorrect).To(Equal(core.BoolPtr(false)))
				Expect(messageInputModel.GetProperties()).ToNot(BeEmpty())
				Expect(messageInputModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))

				messageInputModel.SetProperties(nil)
				Expect(messageInputModel.GetProperties()).To(BeEmpty())

				messageInputModelExpectedMap := make(map[string]interface{})
				messageInputModelExpectedMap["foo"] = core.StringPtr("testString")
				messageInputModel.SetProperties(messageInputModelExpectedMap)
				messageInputModelActualMap := messageInputModel.GetProperties()
				Expect(messageInputModelActualMap).To(Equal(messageInputModelExpectedMap))

				// Construct an instance of the RuntimeIntent model
				runtimeIntentModel := new(assistantv1.RuntimeIntent)
				Expect(runtimeIntentModel).ToNot(BeNil())
				runtimeIntentModel.Intent = core.StringPtr("testString")
				runtimeIntentModel.Confidence = core.Float64Ptr(float64(72.5))
				Expect(runtimeIntentModel.Intent).To(Equal(core.StringPtr("testString")))
				Expect(runtimeIntentModel.Confidence).To(Equal(core.Float64Ptr(float64(72.5))))

				// Construct an instance of the CaptureGroup model
				captureGroupModel := new(assistantv1.CaptureGroup)
				Expect(captureGroupModel).ToNot(BeNil())
				captureGroupModel.Group = core.StringPtr("testString")
				captureGroupModel.Location = []int64{int64(38)}
				Expect(captureGroupModel.Group).To(Equal(core.StringPtr("testString")))
				Expect(captureGroupModel.Location).To(Equal([]int64{int64(38)}))

				// Construct an instance of the RuntimeEntityInterpretation model
				runtimeEntityInterpretationModel := new(assistantv1.RuntimeEntityInterpretation)
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
				runtimeEntityAlternativeModel := new(assistantv1.RuntimeEntityAlternative)
				Expect(runtimeEntityAlternativeModel).ToNot(BeNil())
				runtimeEntityAlternativeModel.Value = core.StringPtr("testString")
				runtimeEntityAlternativeModel.Confidence = core.Float64Ptr(float64(72.5))
				Expect(runtimeEntityAlternativeModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(runtimeEntityAlternativeModel.Confidence).To(Equal(core.Float64Ptr(float64(72.5))))

				// Construct an instance of the RuntimeEntityRole model
				runtimeEntityRoleModel := new(assistantv1.RuntimeEntityRole)
				Expect(runtimeEntityRoleModel).ToNot(BeNil())
				runtimeEntityRoleModel.Type = core.StringPtr("date_from")
				Expect(runtimeEntityRoleModel.Type).To(Equal(core.StringPtr("date_from")))

				// Construct an instance of the RuntimeEntity model
				runtimeEntityModel := new(assistantv1.RuntimeEntity)
				Expect(runtimeEntityModel).ToNot(BeNil())
				runtimeEntityModel.Entity = core.StringPtr("testString")
				runtimeEntityModel.Location = []int64{int64(38)}
				runtimeEntityModel.Value = core.StringPtr("testString")
				runtimeEntityModel.Confidence = core.Float64Ptr(float64(72.5))
				runtimeEntityModel.Groups = []assistantv1.CaptureGroup{*captureGroupModel}
				runtimeEntityModel.Interpretation = runtimeEntityInterpretationModel
				runtimeEntityModel.Alternatives = []assistantv1.RuntimeEntityAlternative{*runtimeEntityAlternativeModel}
				runtimeEntityModel.Role = runtimeEntityRoleModel
				Expect(runtimeEntityModel.Entity).To(Equal(core.StringPtr("testString")))
				Expect(runtimeEntityModel.Location).To(Equal([]int64{int64(38)}))
				Expect(runtimeEntityModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(runtimeEntityModel.Confidence).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(runtimeEntityModel.Groups).To(Equal([]assistantv1.CaptureGroup{*captureGroupModel}))
				Expect(runtimeEntityModel.Interpretation).To(Equal(runtimeEntityInterpretationModel))
				Expect(runtimeEntityModel.Alternatives).To(Equal([]assistantv1.RuntimeEntityAlternative{*runtimeEntityAlternativeModel}))
				Expect(runtimeEntityModel.Role).To(Equal(runtimeEntityRoleModel))

				// Construct an instance of the MessageContextMetadata model
				messageContextMetadataModel := new(assistantv1.MessageContextMetadata)
				Expect(messageContextMetadataModel).ToNot(BeNil())
				messageContextMetadataModel.Deployment = core.StringPtr("testString")
				messageContextMetadataModel.UserID = core.StringPtr("testString")
				Expect(messageContextMetadataModel.Deployment).To(Equal(core.StringPtr("testString")))
				Expect(messageContextMetadataModel.UserID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Context model
				contextModel := new(assistantv1.Context)
				Expect(contextModel).ToNot(BeNil())
				contextModel.ConversationID = core.StringPtr("testString")
				contextModel.System = make(map[string]interface{})
				contextModel.Metadata = messageContextMetadataModel
				contextModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(contextModel.ConversationID).To(Equal(core.StringPtr("testString")))
				Expect(contextModel.System).To(Equal(make(map[string]interface{})))
				Expect(contextModel.Metadata).To(Equal(messageContextMetadataModel))
				Expect(contextModel.GetProperties()).ToNot(BeEmpty())
				Expect(contextModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))

				contextModel.SetProperties(nil)
				Expect(contextModel.GetProperties()).To(BeEmpty())

				contextModelExpectedMap := make(map[string]interface{})
				contextModelExpectedMap["foo"] = core.StringPtr("testString")
				contextModel.SetProperties(contextModelExpectedMap)
				contextModelActualMap := contextModel.GetProperties()
				Expect(contextModelActualMap).To(Equal(contextModelExpectedMap))

				// Construct an instance of the DialogNodeVisitedDetails model
				dialogNodeVisitedDetailsModel := new(assistantv1.DialogNodeVisitedDetails)
				Expect(dialogNodeVisitedDetailsModel).ToNot(BeNil())
				dialogNodeVisitedDetailsModel.DialogNode = core.StringPtr("testString")
				dialogNodeVisitedDetailsModel.Title = core.StringPtr("testString")
				dialogNodeVisitedDetailsModel.Conditions = core.StringPtr("testString")
				Expect(dialogNodeVisitedDetailsModel.DialogNode).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeVisitedDetailsModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeVisitedDetailsModel.Conditions).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the LogMessageSource model
				logMessageSourceModel := new(assistantv1.LogMessageSource)
				Expect(logMessageSourceModel).ToNot(BeNil())
				logMessageSourceModel.Type = core.StringPtr("dialog_node")
				logMessageSourceModel.DialogNode = core.StringPtr("testString")
				Expect(logMessageSourceModel.Type).To(Equal(core.StringPtr("dialog_node")))
				Expect(logMessageSourceModel.DialogNode).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the LogMessage model
				logMessageModel := new(assistantv1.LogMessage)
				Expect(logMessageModel).ToNot(BeNil())
				logMessageModel.Level = core.StringPtr("info")
				logMessageModel.Msg = core.StringPtr("testString")
				logMessageModel.Code = core.StringPtr("testString")
				logMessageModel.Source = logMessageSourceModel
				Expect(logMessageModel.Level).To(Equal(core.StringPtr("info")))
				Expect(logMessageModel.Msg).To(Equal(core.StringPtr("testString")))
				Expect(logMessageModel.Code).To(Equal(core.StringPtr("testString")))
				Expect(logMessageModel.Source).To(Equal(logMessageSourceModel))

				// Construct an instance of the DialogNodeOutputOptionsElementValue model
				dialogNodeOutputOptionsElementValueModel := new(assistantv1.DialogNodeOutputOptionsElementValue)
				Expect(dialogNodeOutputOptionsElementValueModel).ToNot(BeNil())
				dialogNodeOutputOptionsElementValueModel.Input = messageInputModel
				dialogNodeOutputOptionsElementValueModel.Intents = []assistantv1.RuntimeIntent{*runtimeIntentModel}
				dialogNodeOutputOptionsElementValueModel.Entities = []assistantv1.RuntimeEntity{*runtimeEntityModel}
				Expect(dialogNodeOutputOptionsElementValueModel.Input).To(Equal(messageInputModel))
				Expect(dialogNodeOutputOptionsElementValueModel.Intents).To(Equal([]assistantv1.RuntimeIntent{*runtimeIntentModel}))
				Expect(dialogNodeOutputOptionsElementValueModel.Entities).To(Equal([]assistantv1.RuntimeEntity{*runtimeEntityModel}))

				// Construct an instance of the DialogNodeOutputOptionsElement model
				dialogNodeOutputOptionsElementModel := new(assistantv1.DialogNodeOutputOptionsElement)
				Expect(dialogNodeOutputOptionsElementModel).ToNot(BeNil())
				dialogNodeOutputOptionsElementModel.Label = core.StringPtr("testString")
				dialogNodeOutputOptionsElementModel.Value = dialogNodeOutputOptionsElementValueModel
				Expect(dialogNodeOutputOptionsElementModel.Label).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeOutputOptionsElementModel.Value).To(Equal(dialogNodeOutputOptionsElementValueModel))

				// Construct an instance of the ResponseGenericChannel model
				responseGenericChannelModel := new(assistantv1.ResponseGenericChannel)
				Expect(responseGenericChannelModel).ToNot(BeNil())
				responseGenericChannelModel.Channel = core.StringPtr("chat")
				Expect(responseGenericChannelModel.Channel).To(Equal(core.StringPtr("chat")))

				// Construct an instance of the RuntimeResponseGenericRuntimeResponseTypeOption model
				runtimeResponseGenericModel := new(assistantv1.RuntimeResponseGenericRuntimeResponseTypeOption)
				Expect(runtimeResponseGenericModel).ToNot(BeNil())
				runtimeResponseGenericModel.ResponseType = core.StringPtr("option")
				runtimeResponseGenericModel.Title = core.StringPtr("testString")
				runtimeResponseGenericModel.Description = core.StringPtr("testString")
				runtimeResponseGenericModel.Preference = core.StringPtr("dropdown")
				runtimeResponseGenericModel.Options = []assistantv1.DialogNodeOutputOptionsElement{*dialogNodeOutputOptionsElementModel}
				runtimeResponseGenericModel.Channels = []assistantv1.ResponseGenericChannel{*responseGenericChannelModel}
				Expect(runtimeResponseGenericModel.ResponseType).To(Equal(core.StringPtr("option")))
				Expect(runtimeResponseGenericModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(runtimeResponseGenericModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(runtimeResponseGenericModel.Preference).To(Equal(core.StringPtr("dropdown")))
				Expect(runtimeResponseGenericModel.Options).To(Equal([]assistantv1.DialogNodeOutputOptionsElement{*dialogNodeOutputOptionsElementModel}))
				Expect(runtimeResponseGenericModel.Channels).To(Equal([]assistantv1.ResponseGenericChannel{*responseGenericChannelModel}))

				// Construct an instance of the OutputData model
				outputDataModel := new(assistantv1.OutputData)
				Expect(outputDataModel).ToNot(BeNil())
				outputDataModel.NodesVisited = []string{"testString"}
				outputDataModel.NodesVisitedDetails = []assistantv1.DialogNodeVisitedDetails{*dialogNodeVisitedDetailsModel}
				outputDataModel.LogMessages = []assistantv1.LogMessage{*logMessageModel}
				outputDataModel.Generic = []assistantv1.RuntimeResponseGenericIntf{runtimeResponseGenericModel}
				outputDataModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(outputDataModel.NodesVisited).To(Equal([]string{"testString"}))
				Expect(outputDataModel.NodesVisitedDetails).To(Equal([]assistantv1.DialogNodeVisitedDetails{*dialogNodeVisitedDetailsModel}))
				Expect(outputDataModel.LogMessages).To(Equal([]assistantv1.LogMessage{*logMessageModel}))
				Expect(outputDataModel.Generic).To(Equal([]assistantv1.RuntimeResponseGenericIntf{runtimeResponseGenericModel}))
				Expect(outputDataModel.GetProperties()).ToNot(BeEmpty())
				Expect(outputDataModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))

				outputDataModel.SetProperties(nil)
				Expect(outputDataModel.GetProperties()).To(BeEmpty())

				outputDataModelExpectedMap := make(map[string]interface{})
				outputDataModelExpectedMap["foo"] = core.StringPtr("testString")
				outputDataModel.SetProperties(outputDataModelExpectedMap)
				outputDataModelActualMap := outputDataModel.GetProperties()
				Expect(outputDataModelActualMap).To(Equal(outputDataModelExpectedMap))

				// Construct an instance of the MessageOptions model
				workspaceID := "testString"
				messageOptionsModel := assistantService.NewMessageOptions(workspaceID)
				messageOptionsModel.SetWorkspaceID("testString")
				messageOptionsModel.SetInput(messageInputModel)
				messageOptionsModel.SetIntents([]assistantv1.RuntimeIntent{*runtimeIntentModel})
				messageOptionsModel.SetEntities([]assistantv1.RuntimeEntity{*runtimeEntityModel})
				messageOptionsModel.SetAlternateIntents(false)
				messageOptionsModel.SetContext(contextModel)
				messageOptionsModel.SetOutput(outputDataModel)
				messageOptionsModel.SetUserID("testString")
				messageOptionsModel.SetNodesVisitedDetails(false)
				messageOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(messageOptionsModel).ToNot(BeNil())
				Expect(messageOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(messageOptionsModel.Input).To(Equal(messageInputModel))
				Expect(messageOptionsModel.Intents).To(Equal([]assistantv1.RuntimeIntent{*runtimeIntentModel}))
				Expect(messageOptionsModel.Entities).To(Equal([]assistantv1.RuntimeEntity{*runtimeEntityModel}))
				Expect(messageOptionsModel.AlternateIntents).To(Equal(core.BoolPtr(false)))
				Expect(messageOptionsModel.Context).To(Equal(contextModel))
				Expect(messageOptionsModel.Output).To(Equal(outputDataModel))
				Expect(messageOptionsModel.UserID).To(Equal(core.StringPtr("testString")))
				Expect(messageOptionsModel.NodesVisitedDetails).To(Equal(core.BoolPtr(false)))
				Expect(messageOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewOutputData successfully`, func() {
				logMessages := []assistantv1.LogMessage{}
				_model, err := assistantService.NewOutputData(logMessages)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRuntimeEntity successfully`, func() {
				entity := "testString"
				value := "testString"
				_model, err := assistantService.NewRuntimeEntity(entity, value)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRuntimeIntent successfully`, func() {
				intent := "testString"
				confidence := float64(72.5)
				_model, err := assistantService.NewRuntimeIntent(intent, confidence)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSynonym successfully`, func() {
				synonym := "testString"
				_model, err := assistantService.NewSynonym(synonym)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdateCounterexampleOptions successfully`, func() {
				// Construct an instance of the UpdateCounterexampleOptions model
				workspaceID := "testString"
				text := "testString"
				updateCounterexampleOptionsModel := assistantService.NewUpdateCounterexampleOptions(workspaceID, text)
				updateCounterexampleOptionsModel.SetWorkspaceID("testString")
				updateCounterexampleOptionsModel.SetText("testString")
				updateCounterexampleOptionsModel.SetNewText("testString")
				updateCounterexampleOptionsModel.SetIncludeAudit(false)
				updateCounterexampleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateCounterexampleOptionsModel).ToNot(BeNil())
				Expect(updateCounterexampleOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(updateCounterexampleOptionsModel.Text).To(Equal(core.StringPtr("testString")))
				Expect(updateCounterexampleOptionsModel.NewText).To(Equal(core.StringPtr("testString")))
				Expect(updateCounterexampleOptionsModel.IncludeAudit).To(Equal(core.BoolPtr(false)))
				Expect(updateCounterexampleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateDialogNodeOptions successfully`, func() {
				// Construct an instance of the ResponseGenericChannel model
				responseGenericChannelModel := new(assistantv1.ResponseGenericChannel)
				Expect(responseGenericChannelModel).ToNot(BeNil())
				responseGenericChannelModel.Channel = core.StringPtr("chat")
				Expect(responseGenericChannelModel.Channel).To(Equal(core.StringPtr("chat")))

				// Construct an instance of the DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo model
				dialogNodeOutputGenericModel := new(assistantv1.DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo)
				Expect(dialogNodeOutputGenericModel).ToNot(BeNil())
				dialogNodeOutputGenericModel.ResponseType = core.StringPtr("video")
				dialogNodeOutputGenericModel.Source = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Title = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Description = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Channels = []assistantv1.ResponseGenericChannel{*responseGenericChannelModel}
				dialogNodeOutputGenericModel.ChannelOptions = map[string]interface{}{"anyKey": "anyValue"}
				dialogNodeOutputGenericModel.AltText = core.StringPtr("testString")
				Expect(dialogNodeOutputGenericModel.ResponseType).To(Equal(core.StringPtr("video")))
				Expect(dialogNodeOutputGenericModel.Source).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeOutputGenericModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeOutputGenericModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeOutputGenericModel.Channels).To(Equal([]assistantv1.ResponseGenericChannel{*responseGenericChannelModel}))
				Expect(dialogNodeOutputGenericModel.ChannelOptions).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(dialogNodeOutputGenericModel.AltText).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the DialogNodeOutputModifiers model
				dialogNodeOutputModifiersModel := new(assistantv1.DialogNodeOutputModifiers)
				Expect(dialogNodeOutputModifiersModel).ToNot(BeNil())
				dialogNodeOutputModifiersModel.Overwrite = core.BoolPtr(true)
				Expect(dialogNodeOutputModifiersModel.Overwrite).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the DialogNodeOutput model
				dialogNodeOutputModel := new(assistantv1.DialogNodeOutput)
				Expect(dialogNodeOutputModel).ToNot(BeNil())
				dialogNodeOutputModel.Generic = []assistantv1.DialogNodeOutputGenericIntf{dialogNodeOutputGenericModel}
				dialogNodeOutputModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeOutputModel.Modifiers = dialogNodeOutputModifiersModel
				dialogNodeOutputModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(dialogNodeOutputModel.Generic).To(Equal([]assistantv1.DialogNodeOutputGenericIntf{dialogNodeOutputGenericModel}))
				Expect(dialogNodeOutputModel.Integrations).To(Equal(make(map[string]map[string]interface{})))
				Expect(dialogNodeOutputModel.Modifiers).To(Equal(dialogNodeOutputModifiersModel))
				Expect(dialogNodeOutputModel.GetProperties()).ToNot(BeEmpty())
				Expect(dialogNodeOutputModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))

				dialogNodeOutputModel.SetProperties(nil)
				Expect(dialogNodeOutputModel.GetProperties()).To(BeEmpty())

				dialogNodeOutputModelExpectedMap := make(map[string]interface{})
				dialogNodeOutputModelExpectedMap["foo"] = core.StringPtr("testString")
				dialogNodeOutputModel.SetProperties(dialogNodeOutputModelExpectedMap)
				dialogNodeOutputModelActualMap := dialogNodeOutputModel.GetProperties()
				Expect(dialogNodeOutputModelActualMap).To(Equal(dialogNodeOutputModelExpectedMap))

				// Construct an instance of the DialogNodeContext model
				dialogNodeContextModel := new(assistantv1.DialogNodeContext)
				Expect(dialogNodeContextModel).ToNot(BeNil())
				dialogNodeContextModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeContextModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(dialogNodeContextModel.Integrations).To(Equal(make(map[string]map[string]interface{})))
				Expect(dialogNodeContextModel.GetProperties()).ToNot(BeEmpty())
				Expect(dialogNodeContextModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))

				dialogNodeContextModel.SetProperties(nil)
				Expect(dialogNodeContextModel.GetProperties()).To(BeEmpty())

				dialogNodeContextModelExpectedMap := make(map[string]interface{})
				dialogNodeContextModelExpectedMap["foo"] = core.StringPtr("testString")
				dialogNodeContextModel.SetProperties(dialogNodeContextModelExpectedMap)
				dialogNodeContextModelActualMap := dialogNodeContextModel.GetProperties()
				Expect(dialogNodeContextModelActualMap).To(Equal(dialogNodeContextModelExpectedMap))

				// Construct an instance of the DialogNodeNextStep model
				dialogNodeNextStepModel := new(assistantv1.DialogNodeNextStep)
				Expect(dialogNodeNextStepModel).ToNot(BeNil())
				dialogNodeNextStepModel.Behavior = core.StringPtr("get_user_input")
				dialogNodeNextStepModel.DialogNode = core.StringPtr("testString")
				dialogNodeNextStepModel.Selector = core.StringPtr("condition")
				Expect(dialogNodeNextStepModel.Behavior).To(Equal(core.StringPtr("get_user_input")))
				Expect(dialogNodeNextStepModel.DialogNode).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeNextStepModel.Selector).To(Equal(core.StringPtr("condition")))

				// Construct an instance of the DialogNodeAction model
				dialogNodeActionModel := new(assistantv1.DialogNodeAction)
				Expect(dialogNodeActionModel).ToNot(BeNil())
				dialogNodeActionModel.Name = core.StringPtr("testString")
				dialogNodeActionModel.Type = core.StringPtr("client")
				dialogNodeActionModel.Parameters = make(map[string]interface{})
				dialogNodeActionModel.ResultVariable = core.StringPtr("testString")
				dialogNodeActionModel.Credentials = core.StringPtr("testString")
				Expect(dialogNodeActionModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeActionModel.Type).To(Equal(core.StringPtr("client")))
				Expect(dialogNodeActionModel.Parameters).To(Equal(make(map[string]interface{})))
				Expect(dialogNodeActionModel.ResultVariable).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeActionModel.Credentials).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UpdateDialogNodeOptions model
				workspaceID := "testString"
				dialogNode := "testString"
				updateDialogNodeOptionsModel := assistantService.NewUpdateDialogNodeOptions(workspaceID, dialogNode)
				updateDialogNodeOptionsModel.SetWorkspaceID("testString")
				updateDialogNodeOptionsModel.SetDialogNode("testString")
				updateDialogNodeOptionsModel.SetNewDialogNode("testString")
				updateDialogNodeOptionsModel.SetNewDescription("testString")
				updateDialogNodeOptionsModel.SetNewConditions("testString")
				updateDialogNodeOptionsModel.SetNewParent("testString")
				updateDialogNodeOptionsModel.SetNewPreviousSibling("testString")
				updateDialogNodeOptionsModel.SetNewOutput(dialogNodeOutputModel)
				updateDialogNodeOptionsModel.SetNewContext(dialogNodeContextModel)
				updateDialogNodeOptionsModel.SetNewMetadata(make(map[string]interface{}))
				updateDialogNodeOptionsModel.SetNewNextStep(dialogNodeNextStepModel)
				updateDialogNodeOptionsModel.SetNewTitle("testString")
				updateDialogNodeOptionsModel.SetNewType("standard")
				updateDialogNodeOptionsModel.SetNewEventName("focus")
				updateDialogNodeOptionsModel.SetNewVariable("testString")
				updateDialogNodeOptionsModel.SetNewActions([]assistantv1.DialogNodeAction{*dialogNodeActionModel})
				updateDialogNodeOptionsModel.SetNewDigressIn("not_available")
				updateDialogNodeOptionsModel.SetNewDigressOut("allow_returning")
				updateDialogNodeOptionsModel.SetNewDigressOutSlots("not_allowed")
				updateDialogNodeOptionsModel.SetNewUserLabel("testString")
				updateDialogNodeOptionsModel.SetNewDisambiguationOptOut(false)
				updateDialogNodeOptionsModel.SetIncludeAudit(false)
				updateDialogNodeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateDialogNodeOptionsModel).ToNot(BeNil())
				Expect(updateDialogNodeOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(updateDialogNodeOptionsModel.DialogNode).To(Equal(core.StringPtr("testString")))
				Expect(updateDialogNodeOptionsModel.NewDialogNode).To(Equal(core.StringPtr("testString")))
				Expect(updateDialogNodeOptionsModel.NewDescription).To(Equal(core.StringPtr("testString")))
				Expect(updateDialogNodeOptionsModel.NewConditions).To(Equal(core.StringPtr("testString")))
				Expect(updateDialogNodeOptionsModel.NewParent).To(Equal(core.StringPtr("testString")))
				Expect(updateDialogNodeOptionsModel.NewPreviousSibling).To(Equal(core.StringPtr("testString")))
				Expect(updateDialogNodeOptionsModel.NewOutput).To(Equal(dialogNodeOutputModel))
				Expect(updateDialogNodeOptionsModel.NewContext).To(Equal(dialogNodeContextModel))
				Expect(updateDialogNodeOptionsModel.NewMetadata).To(Equal(make(map[string]interface{})))
				Expect(updateDialogNodeOptionsModel.NewNextStep).To(Equal(dialogNodeNextStepModel))
				Expect(updateDialogNodeOptionsModel.NewTitle).To(Equal(core.StringPtr("testString")))
				Expect(updateDialogNodeOptionsModel.NewType).To(Equal(core.StringPtr("standard")))
				Expect(updateDialogNodeOptionsModel.NewEventName).To(Equal(core.StringPtr("focus")))
				Expect(updateDialogNodeOptionsModel.NewVariable).To(Equal(core.StringPtr("testString")))
				Expect(updateDialogNodeOptionsModel.NewActions).To(Equal([]assistantv1.DialogNodeAction{*dialogNodeActionModel}))
				Expect(updateDialogNodeOptionsModel.NewDigressIn).To(Equal(core.StringPtr("not_available")))
				Expect(updateDialogNodeOptionsModel.NewDigressOut).To(Equal(core.StringPtr("allow_returning")))
				Expect(updateDialogNodeOptionsModel.NewDigressOutSlots).To(Equal(core.StringPtr("not_allowed")))
				Expect(updateDialogNodeOptionsModel.NewUserLabel).To(Equal(core.StringPtr("testString")))
				Expect(updateDialogNodeOptionsModel.NewDisambiguationOptOut).To(Equal(core.BoolPtr(false)))
				Expect(updateDialogNodeOptionsModel.IncludeAudit).To(Equal(core.BoolPtr(false)))
				Expect(updateDialogNodeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateEntityOptions successfully`, func() {
				// Construct an instance of the CreateValue model
				createValueModel := new(assistantv1.CreateValue)
				Expect(createValueModel).ToNot(BeNil())
				createValueModel.Value = core.StringPtr("testString")
				createValueModel.Metadata = make(map[string]interface{})
				createValueModel.Type = core.StringPtr("synonyms")
				createValueModel.Synonyms = []string{"testString"}
				createValueModel.Patterns = []string{"testString"}
				Expect(createValueModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(createValueModel.Metadata).To(Equal(make(map[string]interface{})))
				Expect(createValueModel.Type).To(Equal(core.StringPtr("synonyms")))
				Expect(createValueModel.Synonyms).To(Equal([]string{"testString"}))
				Expect(createValueModel.Patterns).To(Equal([]string{"testString"}))

				// Construct an instance of the UpdateEntityOptions model
				workspaceID := "testString"
				entity := "testString"
				updateEntityOptionsModel := assistantService.NewUpdateEntityOptions(workspaceID, entity)
				updateEntityOptionsModel.SetWorkspaceID("testString")
				updateEntityOptionsModel.SetEntity("testString")
				updateEntityOptionsModel.SetNewEntity("testString")
				updateEntityOptionsModel.SetNewDescription("testString")
				updateEntityOptionsModel.SetNewMetadata(make(map[string]interface{}))
				updateEntityOptionsModel.SetNewFuzzyMatch(true)
				updateEntityOptionsModel.SetNewValues([]assistantv1.CreateValue{*createValueModel})
				updateEntityOptionsModel.SetAppend(false)
				updateEntityOptionsModel.SetIncludeAudit(false)
				updateEntityOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateEntityOptionsModel).ToNot(BeNil())
				Expect(updateEntityOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(updateEntityOptionsModel.Entity).To(Equal(core.StringPtr("testString")))
				Expect(updateEntityOptionsModel.NewEntity).To(Equal(core.StringPtr("testString")))
				Expect(updateEntityOptionsModel.NewDescription).To(Equal(core.StringPtr("testString")))
				Expect(updateEntityOptionsModel.NewMetadata).To(Equal(make(map[string]interface{})))
				Expect(updateEntityOptionsModel.NewFuzzyMatch).To(Equal(core.BoolPtr(true)))
				Expect(updateEntityOptionsModel.NewValues).To(Equal([]assistantv1.CreateValue{*createValueModel}))
				Expect(updateEntityOptionsModel.Append).To(Equal(core.BoolPtr(false)))
				Expect(updateEntityOptionsModel.IncludeAudit).To(Equal(core.BoolPtr(false)))
				Expect(updateEntityOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateExampleOptions successfully`, func() {
				// Construct an instance of the Mention model
				mentionModel := new(assistantv1.Mention)
				Expect(mentionModel).ToNot(BeNil())
				mentionModel.Entity = core.StringPtr("testString")
				mentionModel.Location = []int64{int64(38)}
				Expect(mentionModel.Entity).To(Equal(core.StringPtr("testString")))
				Expect(mentionModel.Location).To(Equal([]int64{int64(38)}))

				// Construct an instance of the UpdateExampleOptions model
				workspaceID := "testString"
				intent := "testString"
				text := "testString"
				updateExampleOptionsModel := assistantService.NewUpdateExampleOptions(workspaceID, intent, text)
				updateExampleOptionsModel.SetWorkspaceID("testString")
				updateExampleOptionsModel.SetIntent("testString")
				updateExampleOptionsModel.SetText("testString")
				updateExampleOptionsModel.SetNewText("testString")
				updateExampleOptionsModel.SetNewMentions([]assistantv1.Mention{*mentionModel})
				updateExampleOptionsModel.SetIncludeAudit(false)
				updateExampleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateExampleOptionsModel).ToNot(BeNil())
				Expect(updateExampleOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(updateExampleOptionsModel.Intent).To(Equal(core.StringPtr("testString")))
				Expect(updateExampleOptionsModel.Text).To(Equal(core.StringPtr("testString")))
				Expect(updateExampleOptionsModel.NewText).To(Equal(core.StringPtr("testString")))
				Expect(updateExampleOptionsModel.NewMentions).To(Equal([]assistantv1.Mention{*mentionModel}))
				Expect(updateExampleOptionsModel.IncludeAudit).To(Equal(core.BoolPtr(false)))
				Expect(updateExampleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateIntentOptions successfully`, func() {
				// Construct an instance of the Mention model
				mentionModel := new(assistantv1.Mention)
				Expect(mentionModel).ToNot(BeNil())
				mentionModel.Entity = core.StringPtr("testString")
				mentionModel.Location = []int64{int64(38)}
				Expect(mentionModel.Entity).To(Equal(core.StringPtr("testString")))
				Expect(mentionModel.Location).To(Equal([]int64{int64(38)}))

				// Construct an instance of the Example model
				exampleModel := new(assistantv1.Example)
				Expect(exampleModel).ToNot(BeNil())
				exampleModel.Text = core.StringPtr("testString")
				exampleModel.Mentions = []assistantv1.Mention{*mentionModel}
				Expect(exampleModel.Text).To(Equal(core.StringPtr("testString")))
				Expect(exampleModel.Mentions).To(Equal([]assistantv1.Mention{*mentionModel}))

				// Construct an instance of the UpdateIntentOptions model
				workspaceID := "testString"
				intent := "testString"
				updateIntentOptionsModel := assistantService.NewUpdateIntentOptions(workspaceID, intent)
				updateIntentOptionsModel.SetWorkspaceID("testString")
				updateIntentOptionsModel.SetIntent("testString")
				updateIntentOptionsModel.SetNewIntent("testString")
				updateIntentOptionsModel.SetNewDescription("testString")
				updateIntentOptionsModel.SetNewExamples([]assistantv1.Example{*exampleModel})
				updateIntentOptionsModel.SetAppend(false)
				updateIntentOptionsModel.SetIncludeAudit(false)
				updateIntentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateIntentOptionsModel).ToNot(BeNil())
				Expect(updateIntentOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(updateIntentOptionsModel.Intent).To(Equal(core.StringPtr("testString")))
				Expect(updateIntentOptionsModel.NewIntent).To(Equal(core.StringPtr("testString")))
				Expect(updateIntentOptionsModel.NewDescription).To(Equal(core.StringPtr("testString")))
				Expect(updateIntentOptionsModel.NewExamples).To(Equal([]assistantv1.Example{*exampleModel}))
				Expect(updateIntentOptionsModel.Append).To(Equal(core.BoolPtr(false)))
				Expect(updateIntentOptionsModel.IncludeAudit).To(Equal(core.BoolPtr(false)))
				Expect(updateIntentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateSynonymOptions successfully`, func() {
				// Construct an instance of the UpdateSynonymOptions model
				workspaceID := "testString"
				entity := "testString"
				value := "testString"
				synonym := "testString"
				updateSynonymOptionsModel := assistantService.NewUpdateSynonymOptions(workspaceID, entity, value, synonym)
				updateSynonymOptionsModel.SetWorkspaceID("testString")
				updateSynonymOptionsModel.SetEntity("testString")
				updateSynonymOptionsModel.SetValue("testString")
				updateSynonymOptionsModel.SetSynonym("testString")
				updateSynonymOptionsModel.SetNewSynonym("testString")
				updateSynonymOptionsModel.SetIncludeAudit(false)
				updateSynonymOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateSynonymOptionsModel).ToNot(BeNil())
				Expect(updateSynonymOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(updateSynonymOptionsModel.Entity).To(Equal(core.StringPtr("testString")))
				Expect(updateSynonymOptionsModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(updateSynonymOptionsModel.Synonym).To(Equal(core.StringPtr("testString")))
				Expect(updateSynonymOptionsModel.NewSynonym).To(Equal(core.StringPtr("testString")))
				Expect(updateSynonymOptionsModel.IncludeAudit).To(Equal(core.BoolPtr(false)))
				Expect(updateSynonymOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateValueOptions successfully`, func() {
				// Construct an instance of the UpdateValueOptions model
				workspaceID := "testString"
				entity := "testString"
				value := "testString"
				updateValueOptionsModel := assistantService.NewUpdateValueOptions(workspaceID, entity, value)
				updateValueOptionsModel.SetWorkspaceID("testString")
				updateValueOptionsModel.SetEntity("testString")
				updateValueOptionsModel.SetValue("testString")
				updateValueOptionsModel.SetNewValue("testString")
				updateValueOptionsModel.SetNewMetadata(make(map[string]interface{}))
				updateValueOptionsModel.SetNewType("synonyms")
				updateValueOptionsModel.SetNewSynonyms([]string{"testString"})
				updateValueOptionsModel.SetNewPatterns([]string{"testString"})
				updateValueOptionsModel.SetAppend(false)
				updateValueOptionsModel.SetIncludeAudit(false)
				updateValueOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateValueOptionsModel).ToNot(BeNil())
				Expect(updateValueOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(updateValueOptionsModel.Entity).To(Equal(core.StringPtr("testString")))
				Expect(updateValueOptionsModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(updateValueOptionsModel.NewValue).To(Equal(core.StringPtr("testString")))
				Expect(updateValueOptionsModel.NewMetadata).To(Equal(make(map[string]interface{})))
				Expect(updateValueOptionsModel.NewType).To(Equal(core.StringPtr("synonyms")))
				Expect(updateValueOptionsModel.NewSynonyms).To(Equal([]string{"testString"}))
				Expect(updateValueOptionsModel.NewPatterns).To(Equal([]string{"testString"}))
				Expect(updateValueOptionsModel.Append).To(Equal(core.BoolPtr(false)))
				Expect(updateValueOptionsModel.IncludeAudit).To(Equal(core.BoolPtr(false)))
				Expect(updateValueOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateWorkspaceOptions successfully`, func() {
				// Construct an instance of the ResponseGenericChannel model
				responseGenericChannelModel := new(assistantv1.ResponseGenericChannel)
				Expect(responseGenericChannelModel).ToNot(BeNil())
				responseGenericChannelModel.Channel = core.StringPtr("chat")
				Expect(responseGenericChannelModel.Channel).To(Equal(core.StringPtr("chat")))

				// Construct an instance of the DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo model
				dialogNodeOutputGenericModel := new(assistantv1.DialogNodeOutputGenericDialogNodeOutputResponseTypeVideo)
				Expect(dialogNodeOutputGenericModel).ToNot(BeNil())
				dialogNodeOutputGenericModel.ResponseType = core.StringPtr("video")
				dialogNodeOutputGenericModel.Source = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Title = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Description = core.StringPtr("testString")
				dialogNodeOutputGenericModel.Channels = []assistantv1.ResponseGenericChannel{*responseGenericChannelModel}
				dialogNodeOutputGenericModel.ChannelOptions = map[string]interface{}{"anyKey": "anyValue"}
				dialogNodeOutputGenericModel.AltText = core.StringPtr("testString")
				Expect(dialogNodeOutputGenericModel.ResponseType).To(Equal(core.StringPtr("video")))
				Expect(dialogNodeOutputGenericModel.Source).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeOutputGenericModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeOutputGenericModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeOutputGenericModel.Channels).To(Equal([]assistantv1.ResponseGenericChannel{*responseGenericChannelModel}))
				Expect(dialogNodeOutputGenericModel.ChannelOptions).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(dialogNodeOutputGenericModel.AltText).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the DialogNodeOutputModifiers model
				dialogNodeOutputModifiersModel := new(assistantv1.DialogNodeOutputModifiers)
				Expect(dialogNodeOutputModifiersModel).ToNot(BeNil())
				dialogNodeOutputModifiersModel.Overwrite = core.BoolPtr(true)
				Expect(dialogNodeOutputModifiersModel.Overwrite).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the DialogNodeOutput model
				dialogNodeOutputModel := new(assistantv1.DialogNodeOutput)
				Expect(dialogNodeOutputModel).ToNot(BeNil())
				dialogNodeOutputModel.Generic = []assistantv1.DialogNodeOutputGenericIntf{dialogNodeOutputGenericModel}
				dialogNodeOutputModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeOutputModel.Modifiers = dialogNodeOutputModifiersModel
				dialogNodeOutputModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(dialogNodeOutputModel.Generic).To(Equal([]assistantv1.DialogNodeOutputGenericIntf{dialogNodeOutputGenericModel}))
				Expect(dialogNodeOutputModel.Integrations).To(Equal(make(map[string]map[string]interface{})))
				Expect(dialogNodeOutputModel.Modifiers).To(Equal(dialogNodeOutputModifiersModel))
				Expect(dialogNodeOutputModel.GetProperties()).ToNot(BeEmpty())
				Expect(dialogNodeOutputModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))

				dialogNodeOutputModel.SetProperties(nil)
				Expect(dialogNodeOutputModel.GetProperties()).To(BeEmpty())

				dialogNodeOutputModelExpectedMap := make(map[string]interface{})
				dialogNodeOutputModelExpectedMap["foo"] = core.StringPtr("testString")
				dialogNodeOutputModel.SetProperties(dialogNodeOutputModelExpectedMap)
				dialogNodeOutputModelActualMap := dialogNodeOutputModel.GetProperties()
				Expect(dialogNodeOutputModelActualMap).To(Equal(dialogNodeOutputModelExpectedMap))

				// Construct an instance of the DialogNodeContext model
				dialogNodeContextModel := new(assistantv1.DialogNodeContext)
				Expect(dialogNodeContextModel).ToNot(BeNil())
				dialogNodeContextModel.Integrations = make(map[string]map[string]interface{})
				dialogNodeContextModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(dialogNodeContextModel.Integrations).To(Equal(make(map[string]map[string]interface{})))
				Expect(dialogNodeContextModel.GetProperties()).ToNot(BeEmpty())
				Expect(dialogNodeContextModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))

				dialogNodeContextModel.SetProperties(nil)
				Expect(dialogNodeContextModel.GetProperties()).To(BeEmpty())

				dialogNodeContextModelExpectedMap := make(map[string]interface{})
				dialogNodeContextModelExpectedMap["foo"] = core.StringPtr("testString")
				dialogNodeContextModel.SetProperties(dialogNodeContextModelExpectedMap)
				dialogNodeContextModelActualMap := dialogNodeContextModel.GetProperties()
				Expect(dialogNodeContextModelActualMap).To(Equal(dialogNodeContextModelExpectedMap))

				// Construct an instance of the DialogNodeNextStep model
				dialogNodeNextStepModel := new(assistantv1.DialogNodeNextStep)
				Expect(dialogNodeNextStepModel).ToNot(BeNil())
				dialogNodeNextStepModel.Behavior = core.StringPtr("get_user_input")
				dialogNodeNextStepModel.DialogNode = core.StringPtr("testString")
				dialogNodeNextStepModel.Selector = core.StringPtr("condition")
				Expect(dialogNodeNextStepModel.Behavior).To(Equal(core.StringPtr("get_user_input")))
				Expect(dialogNodeNextStepModel.DialogNode).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeNextStepModel.Selector).To(Equal(core.StringPtr("condition")))

				// Construct an instance of the DialogNodeAction model
				dialogNodeActionModel := new(assistantv1.DialogNodeAction)
				Expect(dialogNodeActionModel).ToNot(BeNil())
				dialogNodeActionModel.Name = core.StringPtr("testString")
				dialogNodeActionModel.Type = core.StringPtr("client")
				dialogNodeActionModel.Parameters = make(map[string]interface{})
				dialogNodeActionModel.ResultVariable = core.StringPtr("testString")
				dialogNodeActionModel.Credentials = core.StringPtr("testString")
				Expect(dialogNodeActionModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeActionModel.Type).To(Equal(core.StringPtr("client")))
				Expect(dialogNodeActionModel.Parameters).To(Equal(make(map[string]interface{})))
				Expect(dialogNodeActionModel.ResultVariable).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeActionModel.Credentials).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the DialogNode model
				dialogNodeModel := new(assistantv1.DialogNode)
				Expect(dialogNodeModel).ToNot(BeNil())
				dialogNodeModel.DialogNode = core.StringPtr("testString")
				dialogNodeModel.Description = core.StringPtr("testString")
				dialogNodeModel.Conditions = core.StringPtr("testString")
				dialogNodeModel.Parent = core.StringPtr("testString")
				dialogNodeModel.PreviousSibling = core.StringPtr("testString")
				dialogNodeModel.Output = dialogNodeOutputModel
				dialogNodeModel.Context = dialogNodeContextModel
				dialogNodeModel.Metadata = make(map[string]interface{})
				dialogNodeModel.NextStep = dialogNodeNextStepModel
				dialogNodeModel.Title = core.StringPtr("testString")
				dialogNodeModel.Type = core.StringPtr("standard")
				dialogNodeModel.EventName = core.StringPtr("focus")
				dialogNodeModel.Variable = core.StringPtr("testString")
				dialogNodeModel.Actions = []assistantv1.DialogNodeAction{*dialogNodeActionModel}
				dialogNodeModel.DigressIn = core.StringPtr("not_available")
				dialogNodeModel.DigressOut = core.StringPtr("allow_returning")
				dialogNodeModel.DigressOutSlots = core.StringPtr("not_allowed")
				dialogNodeModel.UserLabel = core.StringPtr("testString")
				dialogNodeModel.DisambiguationOptOut = core.BoolPtr(false)
				Expect(dialogNodeModel.DialogNode).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeModel.Conditions).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeModel.Parent).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeModel.PreviousSibling).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeModel.Output).To(Equal(dialogNodeOutputModel))
				Expect(dialogNodeModel.Context).To(Equal(dialogNodeContextModel))
				Expect(dialogNodeModel.Metadata).To(Equal(make(map[string]interface{})))
				Expect(dialogNodeModel.NextStep).To(Equal(dialogNodeNextStepModel))
				Expect(dialogNodeModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeModel.Type).To(Equal(core.StringPtr("standard")))
				Expect(dialogNodeModel.EventName).To(Equal(core.StringPtr("focus")))
				Expect(dialogNodeModel.Variable).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeModel.Actions).To(Equal([]assistantv1.DialogNodeAction{*dialogNodeActionModel}))
				Expect(dialogNodeModel.DigressIn).To(Equal(core.StringPtr("not_available")))
				Expect(dialogNodeModel.DigressOut).To(Equal(core.StringPtr("allow_returning")))
				Expect(dialogNodeModel.DigressOutSlots).To(Equal(core.StringPtr("not_allowed")))
				Expect(dialogNodeModel.UserLabel).To(Equal(core.StringPtr("testString")))
				Expect(dialogNodeModel.DisambiguationOptOut).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the Counterexample model
				counterexampleModel := new(assistantv1.Counterexample)
				Expect(counterexampleModel).ToNot(BeNil())
				counterexampleModel.Text = core.StringPtr("testString")
				Expect(counterexampleModel.Text).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the WorkspaceSystemSettingsTooling model
				workspaceSystemSettingsToolingModel := new(assistantv1.WorkspaceSystemSettingsTooling)
				Expect(workspaceSystemSettingsToolingModel).ToNot(BeNil())
				workspaceSystemSettingsToolingModel.StoreGenericResponses = core.BoolPtr(true)
				Expect(workspaceSystemSettingsToolingModel.StoreGenericResponses).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the WorkspaceSystemSettingsDisambiguation model
				workspaceSystemSettingsDisambiguationModel := new(assistantv1.WorkspaceSystemSettingsDisambiguation)
				Expect(workspaceSystemSettingsDisambiguationModel).ToNot(BeNil())
				workspaceSystemSettingsDisambiguationModel.Prompt = core.StringPtr("testString")
				workspaceSystemSettingsDisambiguationModel.NoneOfTheAbovePrompt = core.StringPtr("testString")
				workspaceSystemSettingsDisambiguationModel.Enabled = core.BoolPtr(false)
				workspaceSystemSettingsDisambiguationModel.Sensitivity = core.StringPtr("auto")
				workspaceSystemSettingsDisambiguationModel.Randomize = core.BoolPtr(true)
				workspaceSystemSettingsDisambiguationModel.MaxSuggestions = core.Int64Ptr(int64(1))
				workspaceSystemSettingsDisambiguationModel.SuggestionTextPolicy = core.StringPtr("testString")
				Expect(workspaceSystemSettingsDisambiguationModel.Prompt).To(Equal(core.StringPtr("testString")))
				Expect(workspaceSystemSettingsDisambiguationModel.NoneOfTheAbovePrompt).To(Equal(core.StringPtr("testString")))
				Expect(workspaceSystemSettingsDisambiguationModel.Enabled).To(Equal(core.BoolPtr(false)))
				Expect(workspaceSystemSettingsDisambiguationModel.Sensitivity).To(Equal(core.StringPtr("auto")))
				Expect(workspaceSystemSettingsDisambiguationModel.Randomize).To(Equal(core.BoolPtr(true)))
				Expect(workspaceSystemSettingsDisambiguationModel.MaxSuggestions).To(Equal(core.Int64Ptr(int64(1))))
				Expect(workspaceSystemSettingsDisambiguationModel.SuggestionTextPolicy).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the WorkspaceSystemSettingsSystemEntities model
				workspaceSystemSettingsSystemEntitiesModel := new(assistantv1.WorkspaceSystemSettingsSystemEntities)
				Expect(workspaceSystemSettingsSystemEntitiesModel).ToNot(BeNil())
				workspaceSystemSettingsSystemEntitiesModel.Enabled = core.BoolPtr(false)
				Expect(workspaceSystemSettingsSystemEntitiesModel.Enabled).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the WorkspaceSystemSettingsOffTopic model
				workspaceSystemSettingsOffTopicModel := new(assistantv1.WorkspaceSystemSettingsOffTopic)
				Expect(workspaceSystemSettingsOffTopicModel).ToNot(BeNil())
				workspaceSystemSettingsOffTopicModel.Enabled = core.BoolPtr(false)
				Expect(workspaceSystemSettingsOffTopicModel.Enabled).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the WorkspaceSystemSettings model
				workspaceSystemSettingsModel := new(assistantv1.WorkspaceSystemSettings)
				Expect(workspaceSystemSettingsModel).ToNot(BeNil())
				workspaceSystemSettingsModel.Tooling = workspaceSystemSettingsToolingModel
				workspaceSystemSettingsModel.Disambiguation = workspaceSystemSettingsDisambiguationModel
				workspaceSystemSettingsModel.HumanAgentAssist = make(map[string]interface{})
				workspaceSystemSettingsModel.SpellingSuggestions = core.BoolPtr(false)
				workspaceSystemSettingsModel.SpellingAutoCorrect = core.BoolPtr(false)
				workspaceSystemSettingsModel.SystemEntities = workspaceSystemSettingsSystemEntitiesModel
				workspaceSystemSettingsModel.OffTopic = workspaceSystemSettingsOffTopicModel
				workspaceSystemSettingsModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(workspaceSystemSettingsModel.Tooling).To(Equal(workspaceSystemSettingsToolingModel))
				Expect(workspaceSystemSettingsModel.Disambiguation).To(Equal(workspaceSystemSettingsDisambiguationModel))
				Expect(workspaceSystemSettingsModel.HumanAgentAssist).To(Equal(make(map[string]interface{})))
				Expect(workspaceSystemSettingsModel.SpellingSuggestions).To(Equal(core.BoolPtr(false)))
				Expect(workspaceSystemSettingsModel.SpellingAutoCorrect).To(Equal(core.BoolPtr(false)))
				Expect(workspaceSystemSettingsModel.SystemEntities).To(Equal(workspaceSystemSettingsSystemEntitiesModel))
				Expect(workspaceSystemSettingsModel.OffTopic).To(Equal(workspaceSystemSettingsOffTopicModel))
				Expect(workspaceSystemSettingsModel.GetProperties()).ToNot(BeEmpty())
				Expect(workspaceSystemSettingsModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))

				workspaceSystemSettingsModel.SetProperties(nil)
				Expect(workspaceSystemSettingsModel.GetProperties()).To(BeEmpty())

				workspaceSystemSettingsModelExpectedMap := make(map[string]interface{})
				workspaceSystemSettingsModelExpectedMap["foo"] = core.StringPtr("testString")
				workspaceSystemSettingsModel.SetProperties(workspaceSystemSettingsModelExpectedMap)
				workspaceSystemSettingsModelActualMap := workspaceSystemSettingsModel.GetProperties()
				Expect(workspaceSystemSettingsModelActualMap).To(Equal(workspaceSystemSettingsModelExpectedMap))

				// Construct an instance of the WebhookHeader model
				webhookHeaderModel := new(assistantv1.WebhookHeader)
				Expect(webhookHeaderModel).ToNot(BeNil())
				webhookHeaderModel.Name = core.StringPtr("testString")
				webhookHeaderModel.Value = core.StringPtr("testString")
				Expect(webhookHeaderModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(webhookHeaderModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Webhook model
				webhookModel := new(assistantv1.Webhook)
				Expect(webhookModel).ToNot(BeNil())
				webhookModel.URL = core.StringPtr("testString")
				webhookModel.Name = core.StringPtr("testString")
				webhookModel.HeadersVar = []assistantv1.WebhookHeader{*webhookHeaderModel}
				Expect(webhookModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(webhookModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(webhookModel.HeadersVar).To(Equal([]assistantv1.WebhookHeader{*webhookHeaderModel}))

				// Construct an instance of the Mention model
				mentionModel := new(assistantv1.Mention)
				Expect(mentionModel).ToNot(BeNil())
				mentionModel.Entity = core.StringPtr("testString")
				mentionModel.Location = []int64{int64(38)}
				Expect(mentionModel.Entity).To(Equal(core.StringPtr("testString")))
				Expect(mentionModel.Location).To(Equal([]int64{int64(38)}))

				// Construct an instance of the Example model
				exampleModel := new(assistantv1.Example)
				Expect(exampleModel).ToNot(BeNil())
				exampleModel.Text = core.StringPtr("testString")
				exampleModel.Mentions = []assistantv1.Mention{*mentionModel}
				Expect(exampleModel.Text).To(Equal(core.StringPtr("testString")))
				Expect(exampleModel.Mentions).To(Equal([]assistantv1.Mention{*mentionModel}))

				// Construct an instance of the CreateIntent model
				createIntentModel := new(assistantv1.CreateIntent)
				Expect(createIntentModel).ToNot(BeNil())
				createIntentModel.Intent = core.StringPtr("testString")
				createIntentModel.Description = core.StringPtr("testString")
				createIntentModel.Examples = []assistantv1.Example{*exampleModel}
				Expect(createIntentModel.Intent).To(Equal(core.StringPtr("testString")))
				Expect(createIntentModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createIntentModel.Examples).To(Equal([]assistantv1.Example{*exampleModel}))

				// Construct an instance of the CreateValue model
				createValueModel := new(assistantv1.CreateValue)
				Expect(createValueModel).ToNot(BeNil())
				createValueModel.Value = core.StringPtr("testString")
				createValueModel.Metadata = make(map[string]interface{})
				createValueModel.Type = core.StringPtr("synonyms")
				createValueModel.Synonyms = []string{"testString"}
				createValueModel.Patterns = []string{"testString"}
				Expect(createValueModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(createValueModel.Metadata).To(Equal(make(map[string]interface{})))
				Expect(createValueModel.Type).To(Equal(core.StringPtr("synonyms")))
				Expect(createValueModel.Synonyms).To(Equal([]string{"testString"}))
				Expect(createValueModel.Patterns).To(Equal([]string{"testString"}))

				// Construct an instance of the CreateEntity model
				createEntityModel := new(assistantv1.CreateEntity)
				Expect(createEntityModel).ToNot(BeNil())
				createEntityModel.Entity = core.StringPtr("testString")
				createEntityModel.Description = core.StringPtr("testString")
				createEntityModel.Metadata = make(map[string]interface{})
				createEntityModel.FuzzyMatch = core.BoolPtr(true)
				createEntityModel.Values = []assistantv1.CreateValue{*createValueModel}
				Expect(createEntityModel.Entity).To(Equal(core.StringPtr("testString")))
				Expect(createEntityModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createEntityModel.Metadata).To(Equal(make(map[string]interface{})))
				Expect(createEntityModel.FuzzyMatch).To(Equal(core.BoolPtr(true)))
				Expect(createEntityModel.Values).To(Equal([]assistantv1.CreateValue{*createValueModel}))

				// Construct an instance of the UpdateWorkspaceOptions model
				workspaceID := "testString"
				updateWorkspaceOptionsModel := assistantService.NewUpdateWorkspaceOptions(workspaceID)
				updateWorkspaceOptionsModel.SetWorkspaceID("testString")
				updateWorkspaceOptionsModel.SetName("testString")
				updateWorkspaceOptionsModel.SetDescription("testString")
				updateWorkspaceOptionsModel.SetLanguage("testString")
				updateWorkspaceOptionsModel.SetDialogNodes([]assistantv1.DialogNode{*dialogNodeModel})
				updateWorkspaceOptionsModel.SetCounterexamples([]assistantv1.Counterexample{*counterexampleModel})
				updateWorkspaceOptionsModel.SetMetadata(make(map[string]interface{}))
				updateWorkspaceOptionsModel.SetLearningOptOut(false)
				updateWorkspaceOptionsModel.SetSystemSettings(workspaceSystemSettingsModel)
				updateWorkspaceOptionsModel.SetWebhooks([]assistantv1.Webhook{*webhookModel})
				updateWorkspaceOptionsModel.SetIntents([]assistantv1.CreateIntent{*createIntentModel})
				updateWorkspaceOptionsModel.SetEntities([]assistantv1.CreateEntity{*createEntityModel})
				updateWorkspaceOptionsModel.SetAppend(false)
				updateWorkspaceOptionsModel.SetIncludeAudit(false)
				updateWorkspaceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateWorkspaceOptionsModel).ToNot(BeNil())
				Expect(updateWorkspaceOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(updateWorkspaceOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateWorkspaceOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateWorkspaceOptionsModel.Language).To(Equal(core.StringPtr("testString")))
				Expect(updateWorkspaceOptionsModel.DialogNodes).To(Equal([]assistantv1.DialogNode{*dialogNodeModel}))
				Expect(updateWorkspaceOptionsModel.Counterexamples).To(Equal([]assistantv1.Counterexample{*counterexampleModel}))
				Expect(updateWorkspaceOptionsModel.Metadata).To(Equal(make(map[string]interface{})))
				Expect(updateWorkspaceOptionsModel.LearningOptOut).To(Equal(core.BoolPtr(false)))
				Expect(updateWorkspaceOptionsModel.SystemSettings).To(Equal(workspaceSystemSettingsModel))
				Expect(updateWorkspaceOptionsModel.Webhooks).To(Equal([]assistantv1.Webhook{*webhookModel}))
				Expect(updateWorkspaceOptionsModel.Intents).To(Equal([]assistantv1.CreateIntent{*createIntentModel}))
				Expect(updateWorkspaceOptionsModel.Entities).To(Equal([]assistantv1.CreateEntity{*createEntityModel}))
				Expect(updateWorkspaceOptionsModel.Append).To(Equal(core.BoolPtr(false)))
				Expect(updateWorkspaceOptionsModel.IncludeAudit).To(Equal(core.BoolPtr(false)))
				Expect(updateWorkspaceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewWebhook successfully`, func() {
				url := "testString"
				name := "testString"
				_model, err := assistantService.NewWebhook(url, name)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewWebhookHeader successfully`, func() {
				name := "testString"
				value := "testString"
				_model, err := assistantService.NewWebhookHeader(name, value)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDialogNodeOutputGenericDialogNodeOutputResponseTypeAudio successfully`, func() {
				responseType := "audio"
				source := "testString"
				_model, err := assistantService.NewDialogNodeOutputGenericDialogNodeOutputResponseTypeAudio(responseType, source)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDialogNodeOutputGenericDialogNodeOutputResponseTypeChannelTransfer successfully`, func() {
				responseType := "channel_transfer"
				messageToUser := "testString"
				var transferInfo *assistantv1.ChannelTransferInfo = nil
				_, err := assistantService.NewDialogNodeOutputGenericDialogNodeOutputResponseTypeChannelTransfer(responseType, messageToUser, transferInfo)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewDialogNodeOutputGenericDialogNodeOutputResponseTypeConnectToAgent successfully`, func() {
				responseType := "connect_to_agent"
				_model, err := assistantService.NewDialogNodeOutputGenericDialogNodeOutputResponseTypeConnectToAgent(responseType)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDialogNodeOutputGenericDialogNodeOutputResponseTypeIframe successfully`, func() {
				responseType := "iframe"
				source := "testString"
				_model, err := assistantService.NewDialogNodeOutputGenericDialogNodeOutputResponseTypeIframe(responseType, source)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDialogNodeOutputGenericDialogNodeOutputResponseTypeImage successfully`, func() {
				responseType := "image"
				source := "testString"
				_model, err := assistantService.NewDialogNodeOutputGenericDialogNodeOutputResponseTypeImage(responseType, source)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDialogNodeOutputGenericDialogNodeOutputResponseTypeOption successfully`, func() {
				responseType := "option"
				title := "testString"
				options := []assistantv1.DialogNodeOutputOptionsElement{}
				_model, err := assistantService.NewDialogNodeOutputGenericDialogNodeOutputResponseTypeOption(responseType, title, options)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDialogNodeOutputGenericDialogNodeOutputResponseTypePause successfully`, func() {
				responseType := "pause"
				time := int64(38)
				_model, err := assistantService.NewDialogNodeOutputGenericDialogNodeOutputResponseTypePause(responseType, time)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDialogNodeOutputGenericDialogNodeOutputResponseTypeSearchSkill successfully`, func() {
				responseType := "search_skill"
				query := "testString"
				queryType := "natural_language"
				_model, err := assistantService.NewDialogNodeOutputGenericDialogNodeOutputResponseTypeSearchSkill(responseType, query, queryType)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDialogNodeOutputGenericDialogNodeOutputResponseTypeText successfully`, func() {
				responseType := "text"
				values := []assistantv1.DialogNodeOutputTextValuesElement{}
				_model, err := assistantService.NewDialogNodeOutputGenericDialogNodeOutputResponseTypeText(responseType, values)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDialogNodeOutputGenericDialogNodeOutputResponseTypeUserDefined successfully`, func() {
				responseType := "user_defined"
				userDefined := make(map[string]interface{})
				_model, err := assistantService.NewDialogNodeOutputGenericDialogNodeOutputResponseTypeUserDefined(responseType, userDefined)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDialogNodeOutputGenericDialogNodeOutputResponseTypeVideo successfully`, func() {
				responseType := "video"
				source := "testString"
				_model, err := assistantService.NewDialogNodeOutputGenericDialogNodeOutputResponseTypeVideo(responseType, source)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRuntimeResponseGenericRuntimeResponseTypeAudio successfully`, func() {
				responseType := "audio"
				source := "testString"
				_model, err := assistantService.NewRuntimeResponseGenericRuntimeResponseTypeAudio(responseType, source)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRuntimeResponseGenericRuntimeResponseTypeChannelTransfer successfully`, func() {
				responseType := "channel_transfer"
				messageToUser := "testString"
				var transferInfo *assistantv1.ChannelTransferInfo = nil
				_, err := assistantService.NewRuntimeResponseGenericRuntimeResponseTypeChannelTransfer(responseType, messageToUser, transferInfo)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewRuntimeResponseGenericRuntimeResponseTypeConnectToAgent successfully`, func() {
				responseType := "connect_to_agent"
				_model, err := assistantService.NewRuntimeResponseGenericRuntimeResponseTypeConnectToAgent(responseType)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRuntimeResponseGenericRuntimeResponseTypeIframe successfully`, func() {
				responseType := "iframe"
				source := "testString"
				_model, err := assistantService.NewRuntimeResponseGenericRuntimeResponseTypeIframe(responseType, source)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRuntimeResponseGenericRuntimeResponseTypeImage successfully`, func() {
				responseType := "image"
				source := "testString"
				_model, err := assistantService.NewRuntimeResponseGenericRuntimeResponseTypeImage(responseType, source)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRuntimeResponseGenericRuntimeResponseTypeOption successfully`, func() {
				responseType := "option"
				title := "testString"
				options := []assistantv1.DialogNodeOutputOptionsElement{}
				_model, err := assistantService.NewRuntimeResponseGenericRuntimeResponseTypeOption(responseType, title, options)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRuntimeResponseGenericRuntimeResponseTypePause successfully`, func() {
				responseType := "pause"
				time := int64(38)
				_model, err := assistantService.NewRuntimeResponseGenericRuntimeResponseTypePause(responseType, time)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRuntimeResponseGenericRuntimeResponseTypeSuggestion successfully`, func() {
				responseType := "suggestion"
				title := "testString"
				suggestions := []assistantv1.DialogSuggestion{}
				_model, err := assistantService.NewRuntimeResponseGenericRuntimeResponseTypeSuggestion(responseType, title, suggestions)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRuntimeResponseGenericRuntimeResponseTypeText successfully`, func() {
				responseType := "text"
				text := "testString"
				_model, err := assistantService.NewRuntimeResponseGenericRuntimeResponseTypeText(responseType, text)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRuntimeResponseGenericRuntimeResponseTypeUserDefined successfully`, func() {
				responseType := "user_defined"
				userDefined := make(map[string]interface{})
				_model, err := assistantService.NewRuntimeResponseGenericRuntimeResponseTypeUserDefined(responseType, userDefined)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRuntimeResponseGenericRuntimeResponseTypeVideo successfully`, func() {
				responseType := "video"
				source := "testString"
				_model, err := assistantService.NewRuntimeResponseGenericRuntimeResponseTypeVideo(responseType, source)
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
