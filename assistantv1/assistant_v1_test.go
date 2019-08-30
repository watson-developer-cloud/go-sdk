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

package assistantv1_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/IBM/go-sdk-core/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/assistantv1"
)

var _ = Describe("AssistantV1", func() {
	Describe("Message(messageOptions *MessageOptions)", func() {
		messagePath := "/v1/workspaces/{workspace_id}/message"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		messagePath = strings.Replace(messagePath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - Get response to user input", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(messagePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"input": {}, "intents": [], "entities": [], "context": {}, "output": {"log_messages": [], "text": []}}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call Message", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
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

				messageOptions := testService.NewMessageOptions(workspaceID)
				returnValue, returnValueErr = testService.Message(messageOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetMessageResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListWorkspaces(listWorkspacesOptions *ListWorkspacesOptions)", func() {
		listWorkspacesPath := "/v1/workspaces"
		version := "exampleString"
		accessToken := "0ui9876453"
		Context("Successfully - List workspaces", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listWorkspacesPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"workspaces": [], "pagination": {"refresh_url": "fake RefreshURL"}}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call ListWorkspaces", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ListWorkspaces(nil)
				Expect(returnValueErr).NotTo(BeNil())

				listWorkspacesOptions := testService.NewListWorkspacesOptions()
				returnValue, returnValueErr = testService.ListWorkspaces(listWorkspacesOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListWorkspacesResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateWorkspace(createWorkspaceOptions *CreateWorkspaceOptions)", func() {
		createWorkspacePath := "/v1/workspaces"
		version := "exampleString"
		accessToken := "0ui9876453"
		Context("Successfully - Create workspace", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createWorkspacePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"name": "fake Name", "language": "fake Language", "learning_opt_out": true, "workspace_id": "fake WorkspaceID"}`)
				res.WriteHeader(201)
			}))
			It("Succeed to call CreateWorkspace", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.CreateWorkspace(nil)
				Expect(returnValueErr).NotTo(BeNil())

				createWorkspaceOptions := testService.NewCreateWorkspaceOptions()
				returnValue, returnValueErr = testService.CreateWorkspace(createWorkspaceOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetCreateWorkspaceResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetWorkspace(getWorkspaceOptions *GetWorkspaceOptions)", func() {
		getWorkspacePath := "/v1/workspaces/{workspace_id}"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		getWorkspacePath = strings.Replace(getWorkspacePath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - Get information about a workspace", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getWorkspacePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"name": "fake Name", "language": "fake Language", "learning_opt_out": true, "workspace_id": "fake WorkspaceID"}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call GetWorkspace", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.GetWorkspace(nil)
				Expect(returnValueErr).NotTo(BeNil())

				getWorkspaceOptions := testService.NewGetWorkspaceOptions(workspaceID)
				returnValue, returnValueErr = testService.GetWorkspace(getWorkspaceOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetWorkspaceResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("UpdateWorkspace(updateWorkspaceOptions *UpdateWorkspaceOptions)", func() {
		updateWorkspacePath := "/v1/workspaces/{workspace_id}"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		updateWorkspacePath = strings.Replace(updateWorkspacePath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - Update workspace", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateWorkspacePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"name": "fake Name", "language": "fake Language", "learning_opt_out": true, "workspace_id": "fake WorkspaceID"}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call UpdateWorkspace", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.UpdateWorkspace(nil)
				Expect(returnValueErr).NotTo(BeNil())

				updateWorkspaceOptions := testService.NewUpdateWorkspaceOptions(workspaceID)
				returnValue, returnValueErr = testService.UpdateWorkspace(updateWorkspaceOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetUpdateWorkspaceResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteWorkspace(deleteWorkspaceOptions *DeleteWorkspaceOptions)", func() {
		deleteWorkspacePath := "/v1/workspaces/{workspace_id}"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		deleteWorkspacePath = strings.Replace(deleteWorkspacePath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - Delete workspace", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteWorkspacePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.WriteHeader(200)
			}))
			It("Succeed to call DeleteWorkspace", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.DeleteWorkspace(nil)
				Expect(returnValueErr).NotTo(BeNil())

				deleteWorkspaceOptions := testService.NewDeleteWorkspaceOptions(workspaceID)
				returnValue, returnValueErr = testService.DeleteWorkspace(deleteWorkspaceOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("ListIntents(listIntentsOptions *ListIntentsOptions)", func() {
		listIntentsPath := "/v1/workspaces/{workspace_id}/intents"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		listIntentsPath = strings.Replace(listIntentsPath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - List intents", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listIntentsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"intents": [], "pagination": {"refresh_url": "fake RefreshURL"}}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call ListIntents", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ListIntents(nil)
				Expect(returnValueErr).NotTo(BeNil())

				listIntentsOptions := testService.NewListIntentsOptions(workspaceID)
				returnValue, returnValueErr = testService.ListIntents(listIntentsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListIntentsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateIntent(createIntentOptions *CreateIntentOptions)", func() {
		createIntentPath := "/v1/workspaces/{workspace_id}/intents"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		intent := "exampleString"
		createIntentPath = strings.Replace(createIntentPath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - Create intent", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createIntentPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"intent": "fake Intent"}`)
				res.WriteHeader(201)
			}))
			It("Succeed to call CreateIntent", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.CreateIntent(nil)
				Expect(returnValueErr).NotTo(BeNil())

				createIntentOptions := testService.NewCreateIntentOptions(workspaceID, intent)
				returnValue, returnValueErr = testService.CreateIntent(createIntentOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetCreateIntentResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetIntent(getIntentOptions *GetIntentOptions)", func() {
		getIntentPath := "/v1/workspaces/{workspace_id}/intents/{intent}"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		intent := "exampleString"
		getIntentPath = strings.Replace(getIntentPath, "{workspace_id}", workspaceID, 1)
		getIntentPath = strings.Replace(getIntentPath, "{intent}", intent, 1)
		Context("Successfully - Get intent", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getIntentPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"intent": "fake Intent"}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call GetIntent", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.GetIntent(nil)
				Expect(returnValueErr).NotTo(BeNil())

				getIntentOptions := testService.NewGetIntentOptions(workspaceID, intent)
				returnValue, returnValueErr = testService.GetIntent(getIntentOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetIntentResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("UpdateIntent(updateIntentOptions *UpdateIntentOptions)", func() {
		updateIntentPath := "/v1/workspaces/{workspace_id}/intents/{intent}"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		intent := "exampleString"
		updateIntentPath = strings.Replace(updateIntentPath, "{workspace_id}", workspaceID, 1)
		updateIntentPath = strings.Replace(updateIntentPath, "{intent}", intent, 1)
		Context("Successfully - Update intent", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateIntentPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"intent": "fake Intent"}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call UpdateIntent", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.UpdateIntent(nil)
				Expect(returnValueErr).NotTo(BeNil())

				updateIntentOptions := testService.NewUpdateIntentOptions(workspaceID, intent)
				returnValue, returnValueErr = testService.UpdateIntent(updateIntentOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetUpdateIntentResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteIntent(deleteIntentOptions *DeleteIntentOptions)", func() {
		deleteIntentPath := "/v1/workspaces/{workspace_id}/intents/{intent}"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		intent := "exampleString"
		deleteIntentPath = strings.Replace(deleteIntentPath, "{workspace_id}", workspaceID, 1)
		deleteIntentPath = strings.Replace(deleteIntentPath, "{intent}", intent, 1)
		Context("Successfully - Delete intent", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteIntentPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.WriteHeader(200)
			}))
			It("Succeed to call DeleteIntent", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.DeleteIntent(nil)
				Expect(returnValueErr).NotTo(BeNil())

				deleteIntentOptions := testService.NewDeleteIntentOptions(workspaceID, intent)
				returnValue, returnValueErr = testService.DeleteIntent(deleteIntentOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("ListExamples(listExamplesOptions *ListExamplesOptions)", func() {
		listExamplesPath := "/v1/workspaces/{workspace_id}/intents/{intent}/examples"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		intent := "exampleString"
		listExamplesPath = strings.Replace(listExamplesPath, "{workspace_id}", workspaceID, 1)
		listExamplesPath = strings.Replace(listExamplesPath, "{intent}", intent, 1)
		Context("Successfully - List user input examples", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listExamplesPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"examples": [], "pagination": {"refresh_url": "fake RefreshURL"}}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call ListExamples", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ListExamples(nil)
				Expect(returnValueErr).NotTo(BeNil())

				listExamplesOptions := testService.NewListExamplesOptions(workspaceID, intent)
				returnValue, returnValueErr = testService.ListExamples(listExamplesOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListExamplesResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateExample(createExampleOptions *CreateExampleOptions)", func() {
		createExamplePath := "/v1/workspaces/{workspace_id}/intents/{intent}/examples"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		intent := "exampleString"
		text := "exampleString"
		createExamplePath = strings.Replace(createExamplePath, "{workspace_id}", workspaceID, 1)
		createExamplePath = strings.Replace(createExamplePath, "{intent}", intent, 1)
		Context("Successfully - Create user input example", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createExamplePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"text": "fake Text"}`)
				res.WriteHeader(201)
			}))
			It("Succeed to call CreateExample", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.CreateExample(nil)
				Expect(returnValueErr).NotTo(BeNil())

				createExampleOptions := testService.NewCreateExampleOptions(workspaceID, intent, text)
				returnValue, returnValueErr = testService.CreateExample(createExampleOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetCreateExampleResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetExample(getExampleOptions *GetExampleOptions)", func() {
		getExamplePath := "/v1/workspaces/{workspace_id}/intents/{intent}/examples/{text}"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		intent := "exampleString"
		text := "exampleString"
		getExamplePath = strings.Replace(getExamplePath, "{workspace_id}", workspaceID, 1)
		getExamplePath = strings.Replace(getExamplePath, "{intent}", intent, 1)
		getExamplePath = strings.Replace(getExamplePath, "{text}", text, 1)
		Context("Successfully - Get user input example", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getExamplePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"text": "fake Text"}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call GetExample", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.GetExample(nil)
				Expect(returnValueErr).NotTo(BeNil())

				getExampleOptions := testService.NewGetExampleOptions(workspaceID, intent, text)
				returnValue, returnValueErr = testService.GetExample(getExampleOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetExampleResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("UpdateExample(updateExampleOptions *UpdateExampleOptions)", func() {
		updateExamplePath := "/v1/workspaces/{workspace_id}/intents/{intent}/examples/{text}"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		intent := "exampleString"
		text := "exampleString"
		updateExamplePath = strings.Replace(updateExamplePath, "{workspace_id}", workspaceID, 1)
		updateExamplePath = strings.Replace(updateExamplePath, "{intent}", intent, 1)
		updateExamplePath = strings.Replace(updateExamplePath, "{text}", text, 1)
		Context("Successfully - Update user input example", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateExamplePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"text": "fake Text"}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call UpdateExample", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.UpdateExample(nil)
				Expect(returnValueErr).NotTo(BeNil())

				updateExampleOptions := testService.NewUpdateExampleOptions(workspaceID, intent, text)
				returnValue, returnValueErr = testService.UpdateExample(updateExampleOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetUpdateExampleResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteExample(deleteExampleOptions *DeleteExampleOptions)", func() {
		deleteExamplePath := "/v1/workspaces/{workspace_id}/intents/{intent}/examples/{text}"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		intent := "exampleString"
		text := "exampleString"
		deleteExamplePath = strings.Replace(deleteExamplePath, "{workspace_id}", workspaceID, 1)
		deleteExamplePath = strings.Replace(deleteExamplePath, "{intent}", intent, 1)
		deleteExamplePath = strings.Replace(deleteExamplePath, "{text}", text, 1)
		Context("Successfully - Delete user input example", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteExamplePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.WriteHeader(200)
			}))
			It("Succeed to call DeleteExample", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.DeleteExample(nil)
				Expect(returnValueErr).NotTo(BeNil())

				deleteExampleOptions := testService.NewDeleteExampleOptions(workspaceID, intent, text)
				returnValue, returnValueErr = testService.DeleteExample(deleteExampleOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("ListCounterexamples(listCounterexamplesOptions *ListCounterexamplesOptions)", func() {
		listCounterexamplesPath := "/v1/workspaces/{workspace_id}/counterexamples"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		listCounterexamplesPath = strings.Replace(listCounterexamplesPath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - List counterexamples", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listCounterexamplesPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"counterexamples": [], "pagination": {"refresh_url": "fake RefreshURL"}}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call ListCounterexamples", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ListCounterexamples(nil)
				Expect(returnValueErr).NotTo(BeNil())

				listCounterexamplesOptions := testService.NewListCounterexamplesOptions(workspaceID)
				returnValue, returnValueErr = testService.ListCounterexamples(listCounterexamplesOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListCounterexamplesResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateCounterexample(createCounterexampleOptions *CreateCounterexampleOptions)", func() {
		createCounterexamplePath := "/v1/workspaces/{workspace_id}/counterexamples"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		text := "exampleString"
		createCounterexamplePath = strings.Replace(createCounterexamplePath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - Create counterexample", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createCounterexamplePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"text": "fake Text"}`)
				res.WriteHeader(201)
			}))
			It("Succeed to call CreateCounterexample", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.CreateCounterexample(nil)
				Expect(returnValueErr).NotTo(BeNil())

				createCounterexampleOptions := testService.NewCreateCounterexampleOptions(workspaceID, text)
				returnValue, returnValueErr = testService.CreateCounterexample(createCounterexampleOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetCreateCounterexampleResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetCounterexample(getCounterexampleOptions *GetCounterexampleOptions)", func() {
		getCounterexamplePath := "/v1/workspaces/{workspace_id}/counterexamples/{text}"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		text := "exampleString"
		getCounterexamplePath = strings.Replace(getCounterexamplePath, "{workspace_id}", workspaceID, 1)
		getCounterexamplePath = strings.Replace(getCounterexamplePath, "{text}", text, 1)
		Context("Successfully - Get counterexample", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getCounterexamplePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"text": "fake Text"}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call GetCounterexample", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.GetCounterexample(nil)
				Expect(returnValueErr).NotTo(BeNil())

				getCounterexampleOptions := testService.NewGetCounterexampleOptions(workspaceID, text)
				returnValue, returnValueErr = testService.GetCounterexample(getCounterexampleOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetCounterexampleResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("UpdateCounterexample(updateCounterexampleOptions *UpdateCounterexampleOptions)", func() {
		updateCounterexamplePath := "/v1/workspaces/{workspace_id}/counterexamples/{text}"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		text := "exampleString"
		updateCounterexamplePath = strings.Replace(updateCounterexamplePath, "{workspace_id}", workspaceID, 1)
		updateCounterexamplePath = strings.Replace(updateCounterexamplePath, "{text}", text, 1)
		Context("Successfully - Update counterexample", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateCounterexamplePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"text": "fake Text"}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call UpdateCounterexample", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.UpdateCounterexample(nil)
				Expect(returnValueErr).NotTo(BeNil())

				updateCounterexampleOptions := testService.NewUpdateCounterexampleOptions(workspaceID, text)
				returnValue, returnValueErr = testService.UpdateCounterexample(updateCounterexampleOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetUpdateCounterexampleResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteCounterexample(deleteCounterexampleOptions *DeleteCounterexampleOptions)", func() {
		deleteCounterexamplePath := "/v1/workspaces/{workspace_id}/counterexamples/{text}"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		text := "exampleString"
		deleteCounterexamplePath = strings.Replace(deleteCounterexamplePath, "{workspace_id}", workspaceID, 1)
		deleteCounterexamplePath = strings.Replace(deleteCounterexamplePath, "{text}", text, 1)
		Context("Successfully - Delete counterexample", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteCounterexamplePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.WriteHeader(200)
			}))
			It("Succeed to call DeleteCounterexample", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.DeleteCounterexample(nil)
				Expect(returnValueErr).NotTo(BeNil())

				deleteCounterexampleOptions := testService.NewDeleteCounterexampleOptions(workspaceID, text)
				returnValue, returnValueErr = testService.DeleteCounterexample(deleteCounterexampleOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("ListEntities(listEntitiesOptions *ListEntitiesOptions)", func() {
		listEntitiesPath := "/v1/workspaces/{workspace_id}/entities"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		listEntitiesPath = strings.Replace(listEntitiesPath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - List entities", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listEntitiesPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"entities": [], "pagination": {"refresh_url": "fake RefreshURL"}}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call ListEntities", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ListEntities(nil)
				Expect(returnValueErr).NotTo(BeNil())

				listEntitiesOptions := testService.NewListEntitiesOptions(workspaceID)
				returnValue, returnValueErr = testService.ListEntities(listEntitiesOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListEntitiesResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateEntity(createEntityOptions *CreateEntityOptions)", func() {
		createEntityPath := "/v1/workspaces/{workspace_id}/entities"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		entity := "exampleString"
		createEntityPath = strings.Replace(createEntityPath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - Create entity", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createEntityPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"entity": "fake Entity"}`)
				res.WriteHeader(201)
			}))
			It("Succeed to call CreateEntity", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.CreateEntity(nil)
				Expect(returnValueErr).NotTo(BeNil())

				createEntityOptions := testService.NewCreateEntityOptions(workspaceID, entity)
				returnValue, returnValueErr = testService.CreateEntity(createEntityOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetCreateEntityResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetEntity(getEntityOptions *GetEntityOptions)", func() {
		getEntityPath := "/v1/workspaces/{workspace_id}/entities/{entity}"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		entity := "exampleString"
		getEntityPath = strings.Replace(getEntityPath, "{workspace_id}", workspaceID, 1)
		getEntityPath = strings.Replace(getEntityPath, "{entity}", entity, 1)
		Context("Successfully - Get entity", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getEntityPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"entity": "fake Entity"}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call GetEntity", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.GetEntity(nil)
				Expect(returnValueErr).NotTo(BeNil())

				getEntityOptions := testService.NewGetEntityOptions(workspaceID, entity)
				returnValue, returnValueErr = testService.GetEntity(getEntityOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetEntityResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("UpdateEntity(updateEntityOptions *UpdateEntityOptions)", func() {
		updateEntityPath := "/v1/workspaces/{workspace_id}/entities/{entity}"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		entity := "exampleString"
		updateEntityPath = strings.Replace(updateEntityPath, "{workspace_id}", workspaceID, 1)
		updateEntityPath = strings.Replace(updateEntityPath, "{entity}", entity, 1)
		Context("Successfully - Update entity", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateEntityPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"entity": "fake Entity"}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call UpdateEntity", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.UpdateEntity(nil)
				Expect(returnValueErr).NotTo(BeNil())

				updateEntityOptions := testService.NewUpdateEntityOptions(workspaceID, entity)
				returnValue, returnValueErr = testService.UpdateEntity(updateEntityOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetUpdateEntityResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteEntity(deleteEntityOptions *DeleteEntityOptions)", func() {
		deleteEntityPath := "/v1/workspaces/{workspace_id}/entities/{entity}"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		entity := "exampleString"
		deleteEntityPath = strings.Replace(deleteEntityPath, "{workspace_id}", workspaceID, 1)
		deleteEntityPath = strings.Replace(deleteEntityPath, "{entity}", entity, 1)
		Context("Successfully - Delete entity", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteEntityPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.WriteHeader(200)
			}))
			It("Succeed to call DeleteEntity", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.DeleteEntity(nil)
				Expect(returnValueErr).NotTo(BeNil())

				deleteEntityOptions := testService.NewDeleteEntityOptions(workspaceID, entity)
				returnValue, returnValueErr = testService.DeleteEntity(deleteEntityOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("ListMentions(listMentionsOptions *ListMentionsOptions)", func() {
		listMentionsPath := "/v1/workspaces/{workspace_id}/entities/{entity}/mentions"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		entity := "exampleString"
		listMentionsPath = strings.Replace(listMentionsPath, "{workspace_id}", workspaceID, 1)
		listMentionsPath = strings.Replace(listMentionsPath, "{entity}", entity, 1)
		Context("Successfully - List entity mentions", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listMentionsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"examples": [], "pagination": {"refresh_url": "fake RefreshURL"}}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call ListMentions", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ListMentions(nil)
				Expect(returnValueErr).NotTo(BeNil())

				listMentionsOptions := testService.NewListMentionsOptions(workspaceID, entity)
				returnValue, returnValueErr = testService.ListMentions(listMentionsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListMentionsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListValues(listValuesOptions *ListValuesOptions)", func() {
		listValuesPath := "/v1/workspaces/{workspace_id}/entities/{entity}/values"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		entity := "exampleString"
		listValuesPath = strings.Replace(listValuesPath, "{workspace_id}", workspaceID, 1)
		listValuesPath = strings.Replace(listValuesPath, "{entity}", entity, 1)
		Context("Successfully - List entity values", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listValuesPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"values": [], "pagination": {"refresh_url": "fake RefreshURL"}}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call ListValues", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ListValues(nil)
				Expect(returnValueErr).NotTo(BeNil())

				listValuesOptions := testService.NewListValuesOptions(workspaceID, entity)
				returnValue, returnValueErr = testService.ListValues(listValuesOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListValuesResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateValue(createValueOptions *CreateValueOptions)", func() {
		createValuePath := "/v1/workspaces/{workspace_id}/entities/{entity}/values"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		entity := "exampleString"
		value := "exampleString"
		createValuePath = strings.Replace(createValuePath, "{workspace_id}", workspaceID, 1)
		createValuePath = strings.Replace(createValuePath, "{entity}", entity, 1)
		Context("Successfully - Create entity value", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createValuePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"value": "fake Value", "type": "fake Type"}`)
				res.WriteHeader(201)
			}))
			It("Succeed to call CreateValue", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.CreateValue(nil)
				Expect(returnValueErr).NotTo(BeNil())

				createValueOptions := testService.NewCreateValueOptions(workspaceID, entity, value)
				returnValue, returnValueErr = testService.CreateValue(createValueOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetCreateValueResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetValue(getValueOptions *GetValueOptions)", func() {
		getValuePath := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		entity := "exampleString"
		value := "exampleString"
		getValuePath = strings.Replace(getValuePath, "{workspace_id}", workspaceID, 1)
		getValuePath = strings.Replace(getValuePath, "{entity}", entity, 1)
		getValuePath = strings.Replace(getValuePath, "{value}", value, 1)
		Context("Successfully - Get entity value", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getValuePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"value": "fake Value", "type": "fake Type"}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call GetValue", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.GetValue(nil)
				Expect(returnValueErr).NotTo(BeNil())

				getValueOptions := testService.NewGetValueOptions(workspaceID, entity, value)
				returnValue, returnValueErr = testService.GetValue(getValueOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetValueResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("UpdateValue(updateValueOptions *UpdateValueOptions)", func() {
		updateValuePath := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		entity := "exampleString"
		value := "exampleString"
		updateValuePath = strings.Replace(updateValuePath, "{workspace_id}", workspaceID, 1)
		updateValuePath = strings.Replace(updateValuePath, "{entity}", entity, 1)
		updateValuePath = strings.Replace(updateValuePath, "{value}", value, 1)
		Context("Successfully - Update entity value", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateValuePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"value": "fake Value", "type": "fake Type"}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call UpdateValue", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.UpdateValue(nil)
				Expect(returnValueErr).NotTo(BeNil())

				updateValueOptions := testService.NewUpdateValueOptions(workspaceID, entity, value)
				returnValue, returnValueErr = testService.UpdateValue(updateValueOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetUpdateValueResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteValue(deleteValueOptions *DeleteValueOptions)", func() {
		deleteValuePath := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		entity := "exampleString"
		value := "exampleString"
		deleteValuePath = strings.Replace(deleteValuePath, "{workspace_id}", workspaceID, 1)
		deleteValuePath = strings.Replace(deleteValuePath, "{entity}", entity, 1)
		deleteValuePath = strings.Replace(deleteValuePath, "{value}", value, 1)
		Context("Successfully - Delete entity value", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteValuePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.WriteHeader(200)
			}))
			It("Succeed to call DeleteValue", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.DeleteValue(nil)
				Expect(returnValueErr).NotTo(BeNil())

				deleteValueOptions := testService.NewDeleteValueOptions(workspaceID, entity, value)
				returnValue, returnValueErr = testService.DeleteValue(deleteValueOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("ListSynonyms(listSynonymsOptions *ListSynonymsOptions)", func() {
		listSynonymsPath := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}/synonyms"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		entity := "exampleString"
		value := "exampleString"
		listSynonymsPath = strings.Replace(listSynonymsPath, "{workspace_id}", workspaceID, 1)
		listSynonymsPath = strings.Replace(listSynonymsPath, "{entity}", entity, 1)
		listSynonymsPath = strings.Replace(listSynonymsPath, "{value}", value, 1)
		Context("Successfully - List entity value synonyms", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listSynonymsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"synonyms": [], "pagination": {"refresh_url": "fake RefreshURL"}}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call ListSynonyms", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ListSynonyms(nil)
				Expect(returnValueErr).NotTo(BeNil())

				listSynonymsOptions := testService.NewListSynonymsOptions(workspaceID, entity, value)
				returnValue, returnValueErr = testService.ListSynonyms(listSynonymsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListSynonymsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateSynonym(createSynonymOptions *CreateSynonymOptions)", func() {
		createSynonymPath := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}/synonyms"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		entity := "exampleString"
		value := "exampleString"
		synonym := "exampleString"
		createSynonymPath = strings.Replace(createSynonymPath, "{workspace_id}", workspaceID, 1)
		createSynonymPath = strings.Replace(createSynonymPath, "{entity}", entity, 1)
		createSynonymPath = strings.Replace(createSynonymPath, "{value}", value, 1)
		Context("Successfully - Create entity value synonym", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createSynonymPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"synonym": "fake Synonym"}`)
				res.WriteHeader(201)
			}))
			It("Succeed to call CreateSynonym", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.CreateSynonym(nil)
				Expect(returnValueErr).NotTo(BeNil())

				createSynonymOptions := testService.NewCreateSynonymOptions(workspaceID, entity, value, synonym)
				returnValue, returnValueErr = testService.CreateSynonym(createSynonymOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetCreateSynonymResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetSynonym(getSynonymOptions *GetSynonymOptions)", func() {
		getSynonymPath := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}/synonyms/{synonym}"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		entity := "exampleString"
		value := "exampleString"
		synonym := "exampleString"
		getSynonymPath = strings.Replace(getSynonymPath, "{workspace_id}", workspaceID, 1)
		getSynonymPath = strings.Replace(getSynonymPath, "{entity}", entity, 1)
		getSynonymPath = strings.Replace(getSynonymPath, "{value}", value, 1)
		getSynonymPath = strings.Replace(getSynonymPath, "{synonym}", synonym, 1)
		Context("Successfully - Get entity value synonym", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getSynonymPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"synonym": "fake Synonym"}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call GetSynonym", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.GetSynonym(nil)
				Expect(returnValueErr).NotTo(BeNil())

				getSynonymOptions := testService.NewGetSynonymOptions(workspaceID, entity, value, synonym)
				returnValue, returnValueErr = testService.GetSynonym(getSynonymOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetSynonymResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("UpdateSynonym(updateSynonymOptions *UpdateSynonymOptions)", func() {
		updateSynonymPath := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}/synonyms/{synonym}"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		entity := "exampleString"
		value := "exampleString"
		synonym := "exampleString"
		updateSynonymPath = strings.Replace(updateSynonymPath, "{workspace_id}", workspaceID, 1)
		updateSynonymPath = strings.Replace(updateSynonymPath, "{entity}", entity, 1)
		updateSynonymPath = strings.Replace(updateSynonymPath, "{value}", value, 1)
		updateSynonymPath = strings.Replace(updateSynonymPath, "{synonym}", synonym, 1)
		Context("Successfully - Update entity value synonym", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateSynonymPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"synonym": "fake Synonym"}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call UpdateSynonym", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.UpdateSynonym(nil)
				Expect(returnValueErr).NotTo(BeNil())

				updateSynonymOptions := testService.NewUpdateSynonymOptions(workspaceID, entity, value, synonym)
				returnValue, returnValueErr = testService.UpdateSynonym(updateSynonymOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetUpdateSynonymResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteSynonym(deleteSynonymOptions *DeleteSynonymOptions)", func() {
		deleteSynonymPath := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}/synonyms/{synonym}"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		entity := "exampleString"
		value := "exampleString"
		synonym := "exampleString"
		deleteSynonymPath = strings.Replace(deleteSynonymPath, "{workspace_id}", workspaceID, 1)
		deleteSynonymPath = strings.Replace(deleteSynonymPath, "{entity}", entity, 1)
		deleteSynonymPath = strings.Replace(deleteSynonymPath, "{value}", value, 1)
		deleteSynonymPath = strings.Replace(deleteSynonymPath, "{synonym}", synonym, 1)
		Context("Successfully - Delete entity value synonym", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteSynonymPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.WriteHeader(200)
			}))
			It("Succeed to call DeleteSynonym", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.DeleteSynonym(nil)
				Expect(returnValueErr).NotTo(BeNil())

				deleteSynonymOptions := testService.NewDeleteSynonymOptions(workspaceID, entity, value, synonym)
				returnValue, returnValueErr = testService.DeleteSynonym(deleteSynonymOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("ListDialogNodes(listDialogNodesOptions *ListDialogNodesOptions)", func() {
		listDialogNodesPath := "/v1/workspaces/{workspace_id}/dialog_nodes"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		listDialogNodesPath = strings.Replace(listDialogNodesPath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - List dialog nodes", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listDialogNodesPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"dialog_nodes": [], "pagination": {"refresh_url": "fake RefreshURL"}}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call ListDialogNodes", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ListDialogNodes(nil)
				Expect(returnValueErr).NotTo(BeNil())

				listDialogNodesOptions := testService.NewListDialogNodesOptions(workspaceID)
				returnValue, returnValueErr = testService.ListDialogNodes(listDialogNodesOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListDialogNodesResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateDialogNode(createDialogNodeOptions *CreateDialogNodeOptions)", func() {
		createDialogNodePath := "/v1/workspaces/{workspace_id}/dialog_nodes"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		dialogNode := "exampleString"
		createDialogNodePath = strings.Replace(createDialogNodePath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - Create dialog node", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createDialogNodePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"dialog_node": "fake DialogNode"}`)
				res.WriteHeader(201)
			}))
			It("Succeed to call CreateDialogNode", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.CreateDialogNode(nil)
				Expect(returnValueErr).NotTo(BeNil())

				createDialogNodeOptions := testService.NewCreateDialogNodeOptions(workspaceID, dialogNode)
				returnValue, returnValueErr = testService.CreateDialogNode(createDialogNodeOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetCreateDialogNodeResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetDialogNode(getDialogNodeOptions *GetDialogNodeOptions)", func() {
		getDialogNodePath := "/v1/workspaces/{workspace_id}/dialog_nodes/{dialog_node}"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		dialogNode := "exampleString"
		getDialogNodePath = strings.Replace(getDialogNodePath, "{workspace_id}", workspaceID, 1)
		getDialogNodePath = strings.Replace(getDialogNodePath, "{dialog_node}", dialogNode, 1)
		Context("Successfully - Get dialog node", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getDialogNodePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"dialog_node": "fake DialogNode"}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call GetDialogNode", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.GetDialogNode(nil)
				Expect(returnValueErr).NotTo(BeNil())

				getDialogNodeOptions := testService.NewGetDialogNodeOptions(workspaceID, dialogNode)
				returnValue, returnValueErr = testService.GetDialogNode(getDialogNodeOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetDialogNodeResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("UpdateDialogNode(updateDialogNodeOptions *UpdateDialogNodeOptions)", func() {
		updateDialogNodePath := "/v1/workspaces/{workspace_id}/dialog_nodes/{dialog_node}"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		dialogNode := "exampleString"
		updateDialogNodePath = strings.Replace(updateDialogNodePath, "{workspace_id}", workspaceID, 1)
		updateDialogNodePath = strings.Replace(updateDialogNodePath, "{dialog_node}", dialogNode, 1)
		Context("Successfully - Update dialog node", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateDialogNodePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"dialog_node": "fake DialogNode"}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call UpdateDialogNode", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.UpdateDialogNode(nil)
				Expect(returnValueErr).NotTo(BeNil())

				updateDialogNodeOptions := testService.NewUpdateDialogNodeOptions(workspaceID, dialogNode)
				returnValue, returnValueErr = testService.UpdateDialogNode(updateDialogNodeOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetUpdateDialogNodeResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteDialogNode(deleteDialogNodeOptions *DeleteDialogNodeOptions)", func() {
		deleteDialogNodePath := "/v1/workspaces/{workspace_id}/dialog_nodes/{dialog_node}"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		dialogNode := "exampleString"
		deleteDialogNodePath = strings.Replace(deleteDialogNodePath, "{workspace_id}", workspaceID, 1)
		deleteDialogNodePath = strings.Replace(deleteDialogNodePath, "{dialog_node}", dialogNode, 1)
		Context("Successfully - Delete dialog node", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteDialogNodePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.WriteHeader(200)
			}))
			It("Succeed to call DeleteDialogNode", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.DeleteDialogNode(nil)
				Expect(returnValueErr).NotTo(BeNil())

				deleteDialogNodeOptions := testService.NewDeleteDialogNodeOptions(workspaceID, dialogNode)
				returnValue, returnValueErr = testService.DeleteDialogNode(deleteDialogNodeOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("ListLogs(listLogsOptions *ListLogsOptions)", func() {
		listLogsPath := "/v1/workspaces/{workspace_id}/logs"
		version := "exampleString"
		accessToken := "0ui9876453"
		workspaceID := "exampleString"
		listLogsPath = strings.Replace(listLogsPath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - List log events in a workspace", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listLogsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"logs": [], "pagination": {}}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call ListLogs", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ListLogs(nil)
				Expect(returnValueErr).NotTo(BeNil())

				listLogsOptions := testService.NewListLogsOptions(workspaceID)
				returnValue, returnValueErr = testService.ListLogs(listLogsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListLogsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListAllLogs(listAllLogsOptions *ListAllLogsOptions)", func() {
		listAllLogsPath := "/v1/logs"
		version := "exampleString"
		accessToken := "0ui9876453"
		filter := "exampleString"
		Context("Successfully - List log events in all workspaces", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listAllLogsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				Expect(req.URL.Query()["filter"]).To(Equal([]string{filter}))

				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"logs": [], "pagination": {}}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call ListAllLogs", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ListAllLogs(nil)
				Expect(returnValueErr).NotTo(BeNil())

				listAllLogsOptions := testService.NewListAllLogsOptions(filter)
				returnValue, returnValueErr = testService.ListAllLogs(listAllLogsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListAllLogsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteUserData(deleteUserDataOptions *DeleteUserDataOptions)", func() {
		deleteUserDataPath := "/v1/user_data"
		version := "exampleString"
		accessToken := "0ui9876453"
		customerID := "exampleString"
		Context("Successfully - Delete labeled data", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteUserDataPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				Expect(req.URL.Query()["customer_id"]).To(Equal([]string{customerID}))

				res.WriteHeader(202)
			}))
			It("Succeed to call DeleteUserData", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: accessToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.DeleteUserData(nil)
				Expect(returnValueErr).NotTo(BeNil())

				deleteUserDataOptions := testService.NewDeleteUserDataOptions(customerID)
				returnValue, returnValueErr = testService.DeleteUserData(deleteUserDataOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
})
