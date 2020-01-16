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

package assistantv2_test

import (
	"fmt"
	"github.com/IBM/go-sdk-core/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/assistantv2"
	"net/http"
	"net/http/httptest"
	"strings"
)

var _ = Describe(`AssistantV2`, func() {
	Describe(`CreateSession(createSessionOptions *CreateSessionOptions)`, func() {
		createSessionPath := "/v2/assistants/{assistant_id}/sessions"
		version := "exampleString"
		bearerToken := "0ui9876453"
		assistantID := "exampleString"
		createSessionPath = strings.Replace(createSessionPath, "{assistant_id}", assistantID, 1)
		Context(`Successfully - Create a session`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createSessionPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"session_id": "fake_SessionID"}`)
			}))
			It(`Succeed to call CreateSession`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.CreateSession(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				createSessionOptions := testService.NewCreateSessionOptions(assistantID)
				result, response, operationErr = testService.CreateSession(createSessionOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteSession(deleteSessionOptions *DeleteSessionOptions)`, func() {
		deleteSessionPath := "/v2/assistants/{assistant_id}/sessions/{session_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		assistantID := "exampleString"
		sessionID := "exampleString"
		deleteSessionPath = strings.Replace(deleteSessionPath, "{assistant_id}", assistantID, 1)
		deleteSessionPath = strings.Replace(deleteSessionPath, "{session_id}", sessionID, 1)
		Context(`Successfully - Delete session`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteSessionPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(200)
			}))
			It(`Succeed to call DeleteSession`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.DeleteSession(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				deleteSessionOptions := testService.NewDeleteSessionOptions(assistantID, sessionID)
				response, operationErr = testService.DeleteSession(deleteSessionOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`Message(messageOptions *MessageOptions)`, func() {
		messagePath := "/v2/assistants/{assistant_id}/sessions/{session_id}/message"
		version := "exampleString"
		bearerToken := "0ui9876453"
		assistantID := "exampleString"
		sessionID := "exampleString"
		messagePath = strings.Replace(messagePath, "{assistant_id}", assistantID, 1)
		messagePath = strings.Replace(messagePath, "{session_id}", sessionID, 1)
		Context(`Successfully - Send user input to assistant`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(messagePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"output": {}}`)
			}))
			It(`Succeed to call Message`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.Message(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				messageOptions := testService.NewMessageOptions(assistantID, sessionID)
				result, response, operationErr = testService.Message(messageOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("Model constructor tests", func() {
		Context("with a sample service", func() {
			version := "1970-01-01"
			testService, _ := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
				URL:           "http://assistantv2modelgenerator.com",
				Version:       version,
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It("should call NewCaptureGroup successfully", func() {
				group := "exampleString"
				model, err := testService.NewCaptureGroup(group)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It("should call NewRuntimeEntity successfully", func() {
				entity := "exampleString"
				location := []int64{}
				value := "exampleString"
				model, err := testService.NewRuntimeEntity(entity, location, value)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It("should call NewRuntimeIntent successfully", func() {
				intent := "exampleString"
				confidence := float64(1234)
				model, err := testService.NewRuntimeIntent(intent, confidence)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
})
