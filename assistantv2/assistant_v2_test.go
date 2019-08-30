/**
 * (C) Copyright IBM Corp. 2019.
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
	"github.com/watson-developer-cloud/go-sdk/assistantv2"
    "github.com/IBM/go-sdk-core/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"strings"
)

var _ = Describe("AssistantV2", func() {
	Describe("CreateSession(createSessionOptions *CreateSessionOptions)", func() {
		createSessionPath := "/v2/assistants/{assistant_id}/sessions"
		version := "exampleString"
		accessToken := "0ui9876453"
		assistantID := "exampleString"
		createSessionPath = strings.Replace(createSessionPath, "{assistant_id}", assistantID, 1)
		Context("Successfully - Create a session", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createSessionPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"session_id": "fake SessionID"}`)
				res.WriteHeader(201)
			}))
			It("Succeed to call CreateSession", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.CreateSession(nil)
				Expect(returnValueErr).NotTo(BeNil())

				createSessionOptions := testService.NewCreateSessionOptions(assistantID)
				returnValue, returnValueErr = testService.CreateSession(createSessionOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetCreateSessionResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteSession(deleteSessionOptions *DeleteSessionOptions)", func() {
		deleteSessionPath := "/v2/assistants/{assistant_id}/sessions/{session_id}"
		version := "exampleString"
		accessToken := "0ui9876453"
		assistantID := "exampleString"
		sessionID := "exampleString"
		deleteSessionPath = strings.Replace(deleteSessionPath, "{assistant_id}", assistantID, 1)
		deleteSessionPath = strings.Replace(deleteSessionPath, "{session_id}", sessionID, 1)
		Context("Successfully - Delete session", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteSessionPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.WriteHeader(200)
			}))
			It("Succeed to call DeleteSession", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.DeleteSession(nil)
				Expect(returnValueErr).NotTo(BeNil())

				deleteSessionOptions := testService.NewDeleteSessionOptions(assistantID, sessionID)
				returnValue, returnValueErr = testService.DeleteSession(deleteSessionOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("Message(messageOptions *MessageOptions)", func() {
		messagePath := "/v2/assistants/{assistant_id}/sessions/{session_id}/message"
		version := "exampleString"
		accessToken := "0ui9876453"
		assistantID := "exampleString"
		sessionID := "exampleString"
		messagePath = strings.Replace(messagePath, "{assistant_id}", assistantID, 1)
		messagePath = strings.Replace(messagePath, "{session_id}", sessionID, 1)
		Context("Successfully - Send user input to assistant", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(messagePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"output": {}}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call Message", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.Message(nil)
				Expect(returnValueErr).NotTo(BeNil())

				messageOptions := testService.NewMessageOptions(assistantID, sessionID)
				returnValue, returnValueErr = testService.Message(messageOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetMessageResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
})
