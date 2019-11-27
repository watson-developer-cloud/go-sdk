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
	"github.com/IBM/go-sdk-core/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/assistantv1"
	"net/http"
	"net/http/httptest"
	"strings"
)

var _ = Describe(`AssistantV1`, func() {
	Describe(`Message(messageOptions *MessageOptions)`, func() {
		messagePath := "/v1/workspaces/{workspace_id}/message"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		messagePath = strings.Replace(messagePath, "{workspace_id}", workspaceID, 1)
		Context(`Successfully - Get response to user input`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(messagePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"input": {}, "intents": [], "entities": [], "context": {}, "output": {"log_messages": [], "text": []}}`)
				res.WriteHeader(200)
			}))
			It(`Succeed to call Message`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
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

				messageOptions := testService.NewMessageOptions(workspaceID)
				result, response, operationErr = testService.Message(messageOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListWorkspaces(listWorkspacesOptions *ListWorkspacesOptions)`, func() {
		listWorkspacesPath := "/v1/workspaces"
		version := "exampleString"
		bearerToken := "0ui9876453"
		Context(`Successfully - List workspaces`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listWorkspacesPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"workspaces": [], "pagination": {"refresh_url": "fake RefreshURL"}}`)
				res.WriteHeader(200)
			}))
			It(`Succeed to call ListWorkspaces`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListWorkspaces(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listWorkspacesOptions := testService.NewListWorkspacesOptions()
				result, response, operationErr = testService.ListWorkspaces(listWorkspacesOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateWorkspace(createWorkspaceOptions *CreateWorkspaceOptions)`, func() {
		createWorkspacePath := "/v1/workspaces"
		version := "exampleString"
		bearerToken := "0ui9876453"
		Context(`Successfully - Create workspace`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createWorkspacePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"name": "fake Name", "language": "fake Language", "learning_opt_out": true, "workspace_id": "fake WorkspaceID"}`)
				res.WriteHeader(201)
			}))
			It(`Succeed to call CreateWorkspace`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.CreateWorkspace(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				createWorkspaceOptions := testService.NewCreateWorkspaceOptions()
				result, response, operationErr = testService.CreateWorkspace(createWorkspaceOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetWorkspace(getWorkspaceOptions *GetWorkspaceOptions)`, func() {
		getWorkspacePath := "/v1/workspaces/{workspace_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		getWorkspacePath = strings.Replace(getWorkspacePath, "{workspace_id}", workspaceID, 1)
		Context(`Successfully - Get information about a workspace`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getWorkspacePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"name": "fake Name", "language": "fake Language", "learning_opt_out": true, "workspace_id": "fake WorkspaceID"}`)
				res.WriteHeader(200)
			}))
			It(`Succeed to call GetWorkspace`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetWorkspace(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getWorkspaceOptions := testService.NewGetWorkspaceOptions(workspaceID)
				result, response, operationErr = testService.GetWorkspace(getWorkspaceOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateWorkspace(updateWorkspaceOptions *UpdateWorkspaceOptions)`, func() {
		updateWorkspacePath := "/v1/workspaces/{workspace_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		updateWorkspacePath = strings.Replace(updateWorkspacePath, "{workspace_id}", workspaceID, 1)
		Context(`Successfully - Update workspace`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateWorkspacePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"name": "fake Name", "language": "fake Language", "learning_opt_out": true, "workspace_id": "fake WorkspaceID"}`)
				res.WriteHeader(200)
			}))
			It(`Succeed to call UpdateWorkspace`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.UpdateWorkspace(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				updateWorkspaceOptions := testService.NewUpdateWorkspaceOptions(workspaceID)
				result, response, operationErr = testService.UpdateWorkspace(updateWorkspaceOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteWorkspace(deleteWorkspaceOptions *DeleteWorkspaceOptions)`, func() {
		deleteWorkspacePath := "/v1/workspaces/{workspace_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		deleteWorkspacePath = strings.Replace(deleteWorkspacePath, "{workspace_id}", workspaceID, 1)
		Context(`Successfully - Delete workspace`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteWorkspacePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(200)
			}))
			It(`Succeed to call DeleteWorkspace`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.DeleteWorkspace(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				deleteWorkspaceOptions := testService.NewDeleteWorkspaceOptions(workspaceID)
				response, operationErr = testService.DeleteWorkspace(deleteWorkspaceOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`ListIntents(listIntentsOptions *ListIntentsOptions)`, func() {
		listIntentsPath := "/v1/workspaces/{workspace_id}/intents"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		listIntentsPath = strings.Replace(listIntentsPath, "{workspace_id}", workspaceID, 1)
		Context(`Successfully - List intents`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listIntentsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"intents": [], "pagination": {"refresh_url": "fake RefreshURL"}}`)
				res.WriteHeader(200)
			}))
			It(`Succeed to call ListIntents`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListIntents(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listIntentsOptions := testService.NewListIntentsOptions(workspaceID)
				result, response, operationErr = testService.ListIntents(listIntentsOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateIntent(createIntentOptions *CreateIntentOptions)`, func() {
		createIntentPath := "/v1/workspaces/{workspace_id}/intents"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		intent := "exampleString"
		createIntentPath = strings.Replace(createIntentPath, "{workspace_id}", workspaceID, 1)
		Context(`Successfully - Create intent`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createIntentPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"intent": "fake Intent"}`)
				res.WriteHeader(201)
			}))
			It(`Succeed to call CreateIntent`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.CreateIntent(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				createIntentOptions := testService.NewCreateIntentOptions(workspaceID, intent)
				result, response, operationErr = testService.CreateIntent(createIntentOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetIntent(getIntentOptions *GetIntentOptions)`, func() {
		getIntentPath := "/v1/workspaces/{workspace_id}/intents/{intent}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		intent := "exampleString"
		getIntentPath = strings.Replace(getIntentPath, "{workspace_id}", workspaceID, 1)
		getIntentPath = strings.Replace(getIntentPath, "{intent}", intent, 1)
		Context(`Successfully - Get intent`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getIntentPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"intent": "fake Intent"}`)
				res.WriteHeader(200)
			}))
			It(`Succeed to call GetIntent`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetIntent(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getIntentOptions := testService.NewGetIntentOptions(workspaceID, intent)
				result, response, operationErr = testService.GetIntent(getIntentOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateIntent(updateIntentOptions *UpdateIntentOptions)`, func() {
		updateIntentPath := "/v1/workspaces/{workspace_id}/intents/{intent}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		intent := "exampleString"
		updateIntentPath = strings.Replace(updateIntentPath, "{workspace_id}", workspaceID, 1)
		updateIntentPath = strings.Replace(updateIntentPath, "{intent}", intent, 1)
		Context(`Successfully - Update intent`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateIntentPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"intent": "fake Intent"}`)
				res.WriteHeader(200)
			}))
			It(`Succeed to call UpdateIntent`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.UpdateIntent(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				updateIntentOptions := testService.NewUpdateIntentOptions(workspaceID, intent)
				result, response, operationErr = testService.UpdateIntent(updateIntentOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteIntent(deleteIntentOptions *DeleteIntentOptions)`, func() {
		deleteIntentPath := "/v1/workspaces/{workspace_id}/intents/{intent}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		intent := "exampleString"
		deleteIntentPath = strings.Replace(deleteIntentPath, "{workspace_id}", workspaceID, 1)
		deleteIntentPath = strings.Replace(deleteIntentPath, "{intent}", intent, 1)
		Context(`Successfully - Delete intent`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteIntentPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(200)
			}))
			It(`Succeed to call DeleteIntent`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.DeleteIntent(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				deleteIntentOptions := testService.NewDeleteIntentOptions(workspaceID, intent)
				response, operationErr = testService.DeleteIntent(deleteIntentOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`ListExamples(listExamplesOptions *ListExamplesOptions)`, func() {
		listExamplesPath := "/v1/workspaces/{workspace_id}/intents/{intent}/examples"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		intent := "exampleString"
		listExamplesPath = strings.Replace(listExamplesPath, "{workspace_id}", workspaceID, 1)
		listExamplesPath = strings.Replace(listExamplesPath, "{intent}", intent, 1)
		Context(`Successfully - List user input examples`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listExamplesPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"examples": [], "pagination": {"refresh_url": "fake RefreshURL"}}`)
				res.WriteHeader(200)
			}))
			It(`Succeed to call ListExamples`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListExamples(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listExamplesOptions := testService.NewListExamplesOptions(workspaceID, intent)
				result, response, operationErr = testService.ListExamples(listExamplesOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateExample(createExampleOptions *CreateExampleOptions)`, func() {
		createExamplePath := "/v1/workspaces/{workspace_id}/intents/{intent}/examples"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		intent := "exampleString"
		text := "exampleString"
		createExamplePath = strings.Replace(createExamplePath, "{workspace_id}", workspaceID, 1)
		createExamplePath = strings.Replace(createExamplePath, "{intent}", intent, 1)
		Context(`Successfully - Create user input example`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createExamplePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"text": "fake Text"}`)
				res.WriteHeader(201)
			}))
			It(`Succeed to call CreateExample`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.CreateExample(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				createExampleOptions := testService.NewCreateExampleOptions(workspaceID, intent, text)
				result, response, operationErr = testService.CreateExample(createExampleOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetExample(getExampleOptions *GetExampleOptions)`, func() {
		getExamplePath := "/v1/workspaces/{workspace_id}/intents/{intent}/examples/{text}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		intent := "exampleString"
		text := "exampleString"
		getExamplePath = strings.Replace(getExamplePath, "{workspace_id}", workspaceID, 1)
		getExamplePath = strings.Replace(getExamplePath, "{intent}", intent, 1)
		getExamplePath = strings.Replace(getExamplePath, "{text}", text, 1)
		Context(`Successfully - Get user input example`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getExamplePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"text": "fake Text"}`)
				res.WriteHeader(200)
			}))
			It(`Succeed to call GetExample`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetExample(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getExampleOptions := testService.NewGetExampleOptions(workspaceID, intent, text)
				result, response, operationErr = testService.GetExample(getExampleOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateExample(updateExampleOptions *UpdateExampleOptions)`, func() {
		updateExamplePath := "/v1/workspaces/{workspace_id}/intents/{intent}/examples/{text}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		intent := "exampleString"
		text := "exampleString"
		updateExamplePath = strings.Replace(updateExamplePath, "{workspace_id}", workspaceID, 1)
		updateExamplePath = strings.Replace(updateExamplePath, "{intent}", intent, 1)
		updateExamplePath = strings.Replace(updateExamplePath, "{text}", text, 1)
		Context(`Successfully - Update user input example`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateExamplePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"text": "fake Text"}`)
				res.WriteHeader(200)
			}))
			It(`Succeed to call UpdateExample`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.UpdateExample(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				updateExampleOptions := testService.NewUpdateExampleOptions(workspaceID, intent, text)
				result, response, operationErr = testService.UpdateExample(updateExampleOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteExample(deleteExampleOptions *DeleteExampleOptions)`, func() {
		deleteExamplePath := "/v1/workspaces/{workspace_id}/intents/{intent}/examples/{text}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		intent := "exampleString"
		text := "exampleString"
		deleteExamplePath = strings.Replace(deleteExamplePath, "{workspace_id}", workspaceID, 1)
		deleteExamplePath = strings.Replace(deleteExamplePath, "{intent}", intent, 1)
		deleteExamplePath = strings.Replace(deleteExamplePath, "{text}", text, 1)
		Context(`Successfully - Delete user input example`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteExamplePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(200)
			}))
			It(`Succeed to call DeleteExample`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.DeleteExample(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				deleteExampleOptions := testService.NewDeleteExampleOptions(workspaceID, intent, text)
				response, operationErr = testService.DeleteExample(deleteExampleOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`ListCounterexamples(listCounterexamplesOptions *ListCounterexamplesOptions)`, func() {
		listCounterexamplesPath := "/v1/workspaces/{workspace_id}/counterexamples"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		listCounterexamplesPath = strings.Replace(listCounterexamplesPath, "{workspace_id}", workspaceID, 1)
		Context(`Successfully - List counterexamples`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listCounterexamplesPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"counterexamples": [], "pagination": {"refresh_url": "fake RefreshURL"}}`)
				res.WriteHeader(200)
			}))
			It(`Succeed to call ListCounterexamples`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListCounterexamples(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listCounterexamplesOptions := testService.NewListCounterexamplesOptions(workspaceID)
				result, response, operationErr = testService.ListCounterexamples(listCounterexamplesOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateCounterexample(createCounterexampleOptions *CreateCounterexampleOptions)`, func() {
		createCounterexamplePath := "/v1/workspaces/{workspace_id}/counterexamples"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		text := "exampleString"
		createCounterexamplePath = strings.Replace(createCounterexamplePath, "{workspace_id}", workspaceID, 1)
		Context(`Successfully - Create counterexample`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createCounterexamplePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"text": "fake Text"}`)
				res.WriteHeader(201)
			}))
			It(`Succeed to call CreateCounterexample`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.CreateCounterexample(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				createCounterexampleOptions := testService.NewCreateCounterexampleOptions(workspaceID, text)
				result, response, operationErr = testService.CreateCounterexample(createCounterexampleOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetCounterexample(getCounterexampleOptions *GetCounterexampleOptions)`, func() {
		getCounterexamplePath := "/v1/workspaces/{workspace_id}/counterexamples/{text}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		text := "exampleString"
		getCounterexamplePath = strings.Replace(getCounterexamplePath, "{workspace_id}", workspaceID, 1)
		getCounterexamplePath = strings.Replace(getCounterexamplePath, "{text}", text, 1)
		Context(`Successfully - Get counterexample`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getCounterexamplePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"text": "fake Text"}`)
				res.WriteHeader(200)
			}))
			It(`Succeed to call GetCounterexample`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetCounterexample(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getCounterexampleOptions := testService.NewGetCounterexampleOptions(workspaceID, text)
				result, response, operationErr = testService.GetCounterexample(getCounterexampleOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateCounterexample(updateCounterexampleOptions *UpdateCounterexampleOptions)`, func() {
		updateCounterexamplePath := "/v1/workspaces/{workspace_id}/counterexamples/{text}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		text := "exampleString"
		updateCounterexamplePath = strings.Replace(updateCounterexamplePath, "{workspace_id}", workspaceID, 1)
		updateCounterexamplePath = strings.Replace(updateCounterexamplePath, "{text}", text, 1)
		Context(`Successfully - Update counterexample`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateCounterexamplePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"text": "fake Text"}`)
				res.WriteHeader(200)
			}))
			It(`Succeed to call UpdateCounterexample`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.UpdateCounterexample(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				updateCounterexampleOptions := testService.NewUpdateCounterexampleOptions(workspaceID, text)
				result, response, operationErr = testService.UpdateCounterexample(updateCounterexampleOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteCounterexample(deleteCounterexampleOptions *DeleteCounterexampleOptions)`, func() {
		deleteCounterexamplePath := "/v1/workspaces/{workspace_id}/counterexamples/{text}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		text := "exampleString"
		deleteCounterexamplePath = strings.Replace(deleteCounterexamplePath, "{workspace_id}", workspaceID, 1)
		deleteCounterexamplePath = strings.Replace(deleteCounterexamplePath, "{text}", text, 1)
		Context(`Successfully - Delete counterexample`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteCounterexamplePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(200)
			}))
			It(`Succeed to call DeleteCounterexample`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.DeleteCounterexample(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				deleteCounterexampleOptions := testService.NewDeleteCounterexampleOptions(workspaceID, text)
				response, operationErr = testService.DeleteCounterexample(deleteCounterexampleOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`ListEntities(listEntitiesOptions *ListEntitiesOptions)`, func() {
		listEntitiesPath := "/v1/workspaces/{workspace_id}/entities"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		listEntitiesPath = strings.Replace(listEntitiesPath, "{workspace_id}", workspaceID, 1)
		Context(`Successfully - List entities`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listEntitiesPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"entities": [], "pagination": {"refresh_url": "fake RefreshURL"}}`)
				res.WriteHeader(200)
			}))
			It(`Succeed to call ListEntities`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListEntities(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listEntitiesOptions := testService.NewListEntitiesOptions(workspaceID)
				result, response, operationErr = testService.ListEntities(listEntitiesOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateEntity(createEntityOptions *CreateEntityOptions)`, func() {
		createEntityPath := "/v1/workspaces/{workspace_id}/entities"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		entity := "exampleString"
		createEntityPath = strings.Replace(createEntityPath, "{workspace_id}", workspaceID, 1)
		Context(`Successfully - Create entity`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createEntityPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"entity": "fake Entity"}`)
				res.WriteHeader(201)
			}))
			It(`Succeed to call CreateEntity`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.CreateEntity(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				createEntityOptions := testService.NewCreateEntityOptions(workspaceID, entity)
				result, response, operationErr = testService.CreateEntity(createEntityOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetEntity(getEntityOptions *GetEntityOptions)`, func() {
		getEntityPath := "/v1/workspaces/{workspace_id}/entities/{entity}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		entity := "exampleString"
		getEntityPath = strings.Replace(getEntityPath, "{workspace_id}", workspaceID, 1)
		getEntityPath = strings.Replace(getEntityPath, "{entity}", entity, 1)
		Context(`Successfully - Get entity`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getEntityPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"entity": "fake Entity"}`)
				res.WriteHeader(200)
			}))
			It(`Succeed to call GetEntity`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetEntity(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getEntityOptions := testService.NewGetEntityOptions(workspaceID, entity)
				result, response, operationErr = testService.GetEntity(getEntityOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateEntity(updateEntityOptions *UpdateEntityOptions)`, func() {
		updateEntityPath := "/v1/workspaces/{workspace_id}/entities/{entity}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		entity := "exampleString"
		updateEntityPath = strings.Replace(updateEntityPath, "{workspace_id}", workspaceID, 1)
		updateEntityPath = strings.Replace(updateEntityPath, "{entity}", entity, 1)
		Context(`Successfully - Update entity`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateEntityPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"entity": "fake Entity"}`)
				res.WriteHeader(200)
			}))
			It(`Succeed to call UpdateEntity`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.UpdateEntity(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				updateEntityOptions := testService.NewUpdateEntityOptions(workspaceID, entity)
				result, response, operationErr = testService.UpdateEntity(updateEntityOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteEntity(deleteEntityOptions *DeleteEntityOptions)`, func() {
		deleteEntityPath := "/v1/workspaces/{workspace_id}/entities/{entity}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		entity := "exampleString"
		deleteEntityPath = strings.Replace(deleteEntityPath, "{workspace_id}", workspaceID, 1)
		deleteEntityPath = strings.Replace(deleteEntityPath, "{entity}", entity, 1)
		Context(`Successfully - Delete entity`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteEntityPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(200)
			}))
			It(`Succeed to call DeleteEntity`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.DeleteEntity(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				deleteEntityOptions := testService.NewDeleteEntityOptions(workspaceID, entity)
				response, operationErr = testService.DeleteEntity(deleteEntityOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`ListMentions(listMentionsOptions *ListMentionsOptions)`, func() {
		listMentionsPath := "/v1/workspaces/{workspace_id}/entities/{entity}/mentions"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		entity := "exampleString"
		listMentionsPath = strings.Replace(listMentionsPath, "{workspace_id}", workspaceID, 1)
		listMentionsPath = strings.Replace(listMentionsPath, "{entity}", entity, 1)
		Context(`Successfully - List entity mentions`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listMentionsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"examples": [], "pagination": {"refresh_url": "fake RefreshURL"}}`)
				res.WriteHeader(200)
			}))
			It(`Succeed to call ListMentions`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListMentions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listMentionsOptions := testService.NewListMentionsOptions(workspaceID, entity)
				result, response, operationErr = testService.ListMentions(listMentionsOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListValues(listValuesOptions *ListValuesOptions)`, func() {
		listValuesPath := "/v1/workspaces/{workspace_id}/entities/{entity}/values"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		entity := "exampleString"
		listValuesPath = strings.Replace(listValuesPath, "{workspace_id}", workspaceID, 1)
		listValuesPath = strings.Replace(listValuesPath, "{entity}", entity, 1)
		Context(`Successfully - List entity values`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listValuesPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"values": [], "pagination": {"refresh_url": "fake RefreshURL"}}`)
				res.WriteHeader(200)
			}))
			It(`Succeed to call ListValues`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListValues(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listValuesOptions := testService.NewListValuesOptions(workspaceID, entity)
				result, response, operationErr = testService.ListValues(listValuesOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateValue(createValueOptions *CreateValueOptions)`, func() {
		createValuePath := "/v1/workspaces/{workspace_id}/entities/{entity}/values"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		entity := "exampleString"
		value := "exampleString"
		createValuePath = strings.Replace(createValuePath, "{workspace_id}", workspaceID, 1)
		createValuePath = strings.Replace(createValuePath, "{entity}", entity, 1)
		Context(`Successfully - Create entity value`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createValuePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"value": "fake Value", "type": "fake Type"}`)
				res.WriteHeader(201)
			}))
			It(`Succeed to call CreateValue`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.CreateValue(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				createValueOptions := testService.NewCreateValueOptions(workspaceID, entity, value)
				result, response, operationErr = testService.CreateValue(createValueOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetValue(getValueOptions *GetValueOptions)`, func() {
		getValuePath := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		entity := "exampleString"
		value := "exampleString"
		getValuePath = strings.Replace(getValuePath, "{workspace_id}", workspaceID, 1)
		getValuePath = strings.Replace(getValuePath, "{entity}", entity, 1)
		getValuePath = strings.Replace(getValuePath, "{value}", value, 1)
		Context(`Successfully - Get entity value`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getValuePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"value": "fake Value", "type": "fake Type"}`)
				res.WriteHeader(200)
			}))
			It(`Succeed to call GetValue`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetValue(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getValueOptions := testService.NewGetValueOptions(workspaceID, entity, value)
				result, response, operationErr = testService.GetValue(getValueOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateValue(updateValueOptions *UpdateValueOptions)`, func() {
		updateValuePath := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		entity := "exampleString"
		value := "exampleString"
		updateValuePath = strings.Replace(updateValuePath, "{workspace_id}", workspaceID, 1)
		updateValuePath = strings.Replace(updateValuePath, "{entity}", entity, 1)
		updateValuePath = strings.Replace(updateValuePath, "{value}", value, 1)
		Context(`Successfully - Update entity value`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateValuePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"value": "fake Value", "type": "fake Type"}`)
				res.WriteHeader(200)
			}))
			It(`Succeed to call UpdateValue`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.UpdateValue(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				updateValueOptions := testService.NewUpdateValueOptions(workspaceID, entity, value)
				result, response, operationErr = testService.UpdateValue(updateValueOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteValue(deleteValueOptions *DeleteValueOptions)`, func() {
		deleteValuePath := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		entity := "exampleString"
		value := "exampleString"
		deleteValuePath = strings.Replace(deleteValuePath, "{workspace_id}", workspaceID, 1)
		deleteValuePath = strings.Replace(deleteValuePath, "{entity}", entity, 1)
		deleteValuePath = strings.Replace(deleteValuePath, "{value}", value, 1)
		Context(`Successfully - Delete entity value`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteValuePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(200)
			}))
			It(`Succeed to call DeleteValue`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.DeleteValue(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				deleteValueOptions := testService.NewDeleteValueOptions(workspaceID, entity, value)
				response, operationErr = testService.DeleteValue(deleteValueOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`ListSynonyms(listSynonymsOptions *ListSynonymsOptions)`, func() {
		listSynonymsPath := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}/synonyms"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		entity := "exampleString"
		value := "exampleString"
		listSynonymsPath = strings.Replace(listSynonymsPath, "{workspace_id}", workspaceID, 1)
		listSynonymsPath = strings.Replace(listSynonymsPath, "{entity}", entity, 1)
		listSynonymsPath = strings.Replace(listSynonymsPath, "{value}", value, 1)
		Context(`Successfully - List entity value synonyms`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listSynonymsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"synonyms": [], "pagination": {"refresh_url": "fake RefreshURL"}}`)
				res.WriteHeader(200)
			}))
			It(`Succeed to call ListSynonyms`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListSynonyms(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listSynonymsOptions := testService.NewListSynonymsOptions(workspaceID, entity, value)
				result, response, operationErr = testService.ListSynonyms(listSynonymsOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateSynonym(createSynonymOptions *CreateSynonymOptions)`, func() {
		createSynonymPath := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}/synonyms"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		entity := "exampleString"
		value := "exampleString"
		synonym := "exampleString"
		createSynonymPath = strings.Replace(createSynonymPath, "{workspace_id}", workspaceID, 1)
		createSynonymPath = strings.Replace(createSynonymPath, "{entity}", entity, 1)
		createSynonymPath = strings.Replace(createSynonymPath, "{value}", value, 1)
		Context(`Successfully - Create entity value synonym`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createSynonymPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"synonym": "fake Synonym"}`)
				res.WriteHeader(201)
			}))
			It(`Succeed to call CreateSynonym`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.CreateSynonym(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				createSynonymOptions := testService.NewCreateSynonymOptions(workspaceID, entity, value, synonym)
				result, response, operationErr = testService.CreateSynonym(createSynonymOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetSynonym(getSynonymOptions *GetSynonymOptions)`, func() {
		getSynonymPath := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}/synonyms/{synonym}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		entity := "exampleString"
		value := "exampleString"
		synonym := "exampleString"
		getSynonymPath = strings.Replace(getSynonymPath, "{workspace_id}", workspaceID, 1)
		getSynonymPath = strings.Replace(getSynonymPath, "{entity}", entity, 1)
		getSynonymPath = strings.Replace(getSynonymPath, "{value}", value, 1)
		getSynonymPath = strings.Replace(getSynonymPath, "{synonym}", synonym, 1)
		Context(`Successfully - Get entity value synonym`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getSynonymPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"synonym": "fake Synonym"}`)
				res.WriteHeader(200)
			}))
			It(`Succeed to call GetSynonym`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetSynonym(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getSynonymOptions := testService.NewGetSynonymOptions(workspaceID, entity, value, synonym)
				result, response, operationErr = testService.GetSynonym(getSynonymOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateSynonym(updateSynonymOptions *UpdateSynonymOptions)`, func() {
		updateSynonymPath := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}/synonyms/{synonym}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		entity := "exampleString"
		value := "exampleString"
		synonym := "exampleString"
		updateSynonymPath = strings.Replace(updateSynonymPath, "{workspace_id}", workspaceID, 1)
		updateSynonymPath = strings.Replace(updateSynonymPath, "{entity}", entity, 1)
		updateSynonymPath = strings.Replace(updateSynonymPath, "{value}", value, 1)
		updateSynonymPath = strings.Replace(updateSynonymPath, "{synonym}", synonym, 1)
		Context(`Successfully - Update entity value synonym`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateSynonymPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"synonym": "fake Synonym"}`)
				res.WriteHeader(200)
			}))
			It(`Succeed to call UpdateSynonym`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.UpdateSynonym(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				updateSynonymOptions := testService.NewUpdateSynonymOptions(workspaceID, entity, value, synonym)
				result, response, operationErr = testService.UpdateSynonym(updateSynonymOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteSynonym(deleteSynonymOptions *DeleteSynonymOptions)`, func() {
		deleteSynonymPath := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}/synonyms/{synonym}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		entity := "exampleString"
		value := "exampleString"
		synonym := "exampleString"
		deleteSynonymPath = strings.Replace(deleteSynonymPath, "{workspace_id}", workspaceID, 1)
		deleteSynonymPath = strings.Replace(deleteSynonymPath, "{entity}", entity, 1)
		deleteSynonymPath = strings.Replace(deleteSynonymPath, "{value}", value, 1)
		deleteSynonymPath = strings.Replace(deleteSynonymPath, "{synonym}", synonym, 1)
		Context(`Successfully - Delete entity value synonym`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteSynonymPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(200)
			}))
			It(`Succeed to call DeleteSynonym`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.DeleteSynonym(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				deleteSynonymOptions := testService.NewDeleteSynonymOptions(workspaceID, entity, value, synonym)
				response, operationErr = testService.DeleteSynonym(deleteSynonymOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`ListDialogNodes(listDialogNodesOptions *ListDialogNodesOptions)`, func() {
		listDialogNodesPath := "/v1/workspaces/{workspace_id}/dialog_nodes"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		listDialogNodesPath = strings.Replace(listDialogNodesPath, "{workspace_id}", workspaceID, 1)
		Context(`Successfully - List dialog nodes`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listDialogNodesPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"dialog_nodes": [], "pagination": {"refresh_url": "fake RefreshURL"}}`)
				res.WriteHeader(200)
			}))
			It(`Succeed to call ListDialogNodes`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListDialogNodes(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listDialogNodesOptions := testService.NewListDialogNodesOptions(workspaceID)
				result, response, operationErr = testService.ListDialogNodes(listDialogNodesOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateDialogNode(createDialogNodeOptions *CreateDialogNodeOptions)`, func() {
		createDialogNodePath := "/v1/workspaces/{workspace_id}/dialog_nodes"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		dialogNode := "exampleString"
		createDialogNodePath = strings.Replace(createDialogNodePath, "{workspace_id}", workspaceID, 1)
		Context(`Successfully - Create dialog node`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createDialogNodePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"dialog_node": "fake DialogNode"}`)
				res.WriteHeader(201)
			}))
			It(`Succeed to call CreateDialogNode`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.CreateDialogNode(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				createDialogNodeOptions := testService.NewCreateDialogNodeOptions(workspaceID, dialogNode)
				result, response, operationErr = testService.CreateDialogNode(createDialogNodeOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetDialogNode(getDialogNodeOptions *GetDialogNodeOptions)`, func() {
		getDialogNodePath := "/v1/workspaces/{workspace_id}/dialog_nodes/{dialog_node}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		dialogNode := "exampleString"
		getDialogNodePath = strings.Replace(getDialogNodePath, "{workspace_id}", workspaceID, 1)
		getDialogNodePath = strings.Replace(getDialogNodePath, "{dialog_node}", dialogNode, 1)
		Context(`Successfully - Get dialog node`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getDialogNodePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"dialog_node": "fake DialogNode"}`)
				res.WriteHeader(200)
			}))
			It(`Succeed to call GetDialogNode`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetDialogNode(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getDialogNodeOptions := testService.NewGetDialogNodeOptions(workspaceID, dialogNode)
				result, response, operationErr = testService.GetDialogNode(getDialogNodeOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateDialogNode(updateDialogNodeOptions *UpdateDialogNodeOptions)`, func() {
		updateDialogNodePath := "/v1/workspaces/{workspace_id}/dialog_nodes/{dialog_node}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		dialogNode := "exampleString"
		updateDialogNodePath = strings.Replace(updateDialogNodePath, "{workspace_id}", workspaceID, 1)
		updateDialogNodePath = strings.Replace(updateDialogNodePath, "{dialog_node}", dialogNode, 1)
		Context(`Successfully - Update dialog node`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateDialogNodePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"dialog_node": "fake DialogNode"}`)
				res.WriteHeader(200)
			}))
			It(`Succeed to call UpdateDialogNode`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.UpdateDialogNode(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				updateDialogNodeOptions := testService.NewUpdateDialogNodeOptions(workspaceID, dialogNode)
				result, response, operationErr = testService.UpdateDialogNode(updateDialogNodeOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteDialogNode(deleteDialogNodeOptions *DeleteDialogNodeOptions)`, func() {
		deleteDialogNodePath := "/v1/workspaces/{workspace_id}/dialog_nodes/{dialog_node}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		dialogNode := "exampleString"
		deleteDialogNodePath = strings.Replace(deleteDialogNodePath, "{workspace_id}", workspaceID, 1)
		deleteDialogNodePath = strings.Replace(deleteDialogNodePath, "{dialog_node}", dialogNode, 1)
		Context(`Successfully - Delete dialog node`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteDialogNodePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(200)
			}))
			It(`Succeed to call DeleteDialogNode`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.DeleteDialogNode(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				deleteDialogNodeOptions := testService.NewDeleteDialogNodeOptions(workspaceID, dialogNode)
				response, operationErr = testService.DeleteDialogNode(deleteDialogNodeOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`ListLogs(listLogsOptions *ListLogsOptions)`, func() {
		listLogsPath := "/v1/workspaces/{workspace_id}/logs"
		version := "exampleString"
		bearerToken := "0ui9876453"
		workspaceID := "exampleString"
		listLogsPath = strings.Replace(listLogsPath, "{workspace_id}", workspaceID, 1)
		Context(`Successfully - List log events in a workspace`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listLogsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"logs": [], "pagination": {}}`)
				res.WriteHeader(200)
			}))
			It(`Succeed to call ListLogs`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListLogs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listLogsOptions := testService.NewListLogsOptions(workspaceID)
				result, response, operationErr = testService.ListLogs(listLogsOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListAllLogs(listAllLogsOptions *ListAllLogsOptions)`, func() {
		listAllLogsPath := "/v1/logs"
		version := "exampleString"
		bearerToken := "0ui9876453"
		filter := "exampleString"
		Context(`Successfully - List log events in all workspaces`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listAllLogsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				Expect(req.URL.Query()["filter"]).To(Equal([]string{filter}))

				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"logs": [], "pagination": {}}`)
				res.WriteHeader(200)
			}))
			It(`Succeed to call ListAllLogs`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListAllLogs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listAllLogsOptions := testService.NewListAllLogsOptions(filter)
				result, response, operationErr = testService.ListAllLogs(listAllLogsOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteUserData(deleteUserDataOptions *DeleteUserDataOptions)`, func() {
		deleteUserDataPath := "/v1/user_data"
		version := "exampleString"
		bearerToken := "0ui9876453"
		customerID := "exampleString"
		Context(`Successfully - Delete labeled data`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteUserDataPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				Expect(req.URL.Query()["customer_id"]).To(Equal([]string{customerID}))

				res.WriteHeader(202)
			}))
			It(`Succeed to call DeleteUserData`, func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.DeleteUserData(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				deleteUserDataOptions := testService.NewDeleteUserDataOptions(customerID)
				response, operationErr = testService.DeleteUserData(deleteUserDataOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
})
