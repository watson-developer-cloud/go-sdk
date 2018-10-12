package assistantv1_test

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	"github.com/cloudfoundry-community/go-cfenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/assistantv1"
)

var _ = Describe("AssistantV1", func() {
	Describe("Get credentials from VCAP", func() {
		version := "exampleString"
		username := "hyphenated-user"
		password := "hyphenated-pass"
		VCAPservices := cfenv.Services{
			"conversation": {
				{
					Name: "conversation",
					Tags: []string{},
					Credentials: map[string]interface{}{
						"url":      "https://gateway.watsonplatform.net/assistant/api",
						"username": username,
						"password": password,
					},
				},
			},
		}
		VCAPbytes, _ := json.Marshal(cfenv.App{})
		os.Setenv("VCAP_APPLICATION", string(VCAPbytes))
		VCAPbytes, _ = json.Marshal(VCAPservices)
		os.Setenv("VCAP_SERVICES", string(VCAPbytes))
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Create AssistantV1 with VCAP credentials", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
			}))
			It("Succeed to create AssistantV1", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:     testServer.URL,
					Version: version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				testService.ListWorkspaces(testService.NewListWorkspacesOptions())
			})
		})
	})
	Describe("Message(messageOptions *MessageOptions)", func() {
		messagePath := "/v1/workspaces/{workspace_id}/message"
		version := "exampleString"
		workspaceID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(messagePath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - Get response to user input", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"input":"xxx"}`)
			}))
			It("Succeed to call Message", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("CreateWorkspace(createWorkspaceOptions *CreateWorkspaceOptions)", func() {
		createWorkspacePath := "/v1/workspaces"
		version := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Create workspace", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(createWorkspacePath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(createWorkspacePath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"name":"xxx"}`)
			}))
			It("Succeed to call CreateWorkspace", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("DeleteWorkspace(deleteWorkspaceOptions *DeleteWorkspaceOptions)", func() {
		deleteWorkspacePath := "/v1/workspaces/{workspace_id}"
		version := "exampleString"
		workspaceID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(deleteWorkspacePath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - Delete workspace", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
			}))
			It("Succeed to call DeleteWorkspace", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("GetWorkspace(getWorkspaceOptions *GetWorkspaceOptions)", func() {
		getWorkspacePath := "/v1/workspaces/{workspace_id}"
		version := "exampleString"
		workspaceID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(getWorkspacePath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - Get information about a workspace", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"name":"xxx"}`)
			}))
			It("Succeed to call GetWorkspace", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("ListWorkspaces(listWorkspacesOptions *ListWorkspacesOptions)", func() {
		listWorkspacesPath := "/v1/workspaces"
		version := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - List workspaces", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(listWorkspacesPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(listWorkspacesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"workspaces":[]}`)
			}))
			It("Succeed to call ListWorkspaces", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("UpdateWorkspace(updateWorkspaceOptions *UpdateWorkspaceOptions)", func() {
		updateWorkspacePath := "/v1/workspaces/{workspace_id}"
		version := "exampleString"
		workspaceID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(updateWorkspacePath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - Update workspace", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"name":"xxx"}`)
			}))
			It("Succeed to call UpdateWorkspace", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("CreateIntent(createIntentOptions *CreateIntentOptions)", func() {
		createIntentPath := "/v1/workspaces/{workspace_id}/intents"
		version := "exampleString"
		workspaceID := "exampleString"
		intent := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(createIntentPath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - Create intent", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"intent":"xxx"}`)
			}))
			It("Succeed to call CreateIntent", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("DeleteIntent(deleteIntentOptions *DeleteIntentOptions)", func() {
		deleteIntentPath := "/v1/workspaces/{workspace_id}/intents/{intent}"
		version := "exampleString"
		workspaceID := "exampleString"
		intent := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(deleteIntentPath, "{workspace_id}", workspaceID, 1)
		Path = strings.Replace(Path, "{intent}", intent, 1)
		Context("Successfully - Delete intent", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
			}))
			It("Succeed to call DeleteIntent", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("GetIntent(getIntentOptions *GetIntentOptions)", func() {
		getIntentPath := "/v1/workspaces/{workspace_id}/intents/{intent}"
		version := "exampleString"
		workspaceID := "exampleString"
		intent := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(getIntentPath, "{workspace_id}", workspaceID, 1)
		Path = strings.Replace(Path, "{intent}", intent, 1)
		Context("Successfully - Get intent", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"intent":"xxx"}`)
			}))
			It("Succeed to call GetIntent", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("ListIntents(listIntentsOptions *ListIntentsOptions)", func() {
		listIntentsPath := "/v1/workspaces/{workspace_id}/intents"
		version := "exampleString"
		workspaceID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(listIntentsPath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - List intents", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"intents":[]}`)
			}))
			It("Succeed to call ListIntents", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("UpdateIntent(updateIntentOptions *UpdateIntentOptions)", func() {
		updateIntentPath := "/v1/workspaces/{workspace_id}/intents/{intent}"
		version := "exampleString"
		workspaceID := "exampleString"
		intent := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(updateIntentPath, "{workspace_id}", workspaceID, 1)
		Path = strings.Replace(Path, "{intent}", intent, 1)
		Context("Successfully - Update intent", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"intent":"xxx"}`)
			}))
			It("Succeed to call UpdateIntent", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("CreateExample(createExampleOptions *CreateExampleOptions)", func() {
		createExamplePath := "/v1/workspaces/{workspace_id}/intents/{intent}/examples"
		version := "exampleString"
		workspaceID := "exampleString"
		intent := "exampleString"
		text := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(createExamplePath, "{workspace_id}", workspaceID, 1)
		Path = strings.Replace(Path, "{intent}", intent, 1)
		Context("Successfully - Create user input example", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"text":"xxx"}`)
			}))
			It("Succeed to call CreateExample", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("DeleteExample(deleteExampleOptions *DeleteExampleOptions)", func() {
		deleteExamplePath := "/v1/workspaces/{workspace_id}/intents/{intent}/examples/{text}"
		version := "exampleString"
		workspaceID := "exampleString"
		intent := "exampleString"
		text := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(deleteExamplePath, "{workspace_id}", workspaceID, 1)
		Path = strings.Replace(Path, "{intent}", intent, 1)
		Path = strings.Replace(Path, "{text}", text, 1)
		Context("Successfully - Delete user input example", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
			}))
			It("Succeed to call DeleteExample", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("GetExample(getExampleOptions *GetExampleOptions)", func() {
		getExamplePath := "/v1/workspaces/{workspace_id}/intents/{intent}/examples/{text}"
		version := "exampleString"
		workspaceID := "exampleString"
		intent := "exampleString"
		text := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(getExamplePath, "{workspace_id}", workspaceID, 1)
		Path = strings.Replace(Path, "{intent}", intent, 1)
		Path = strings.Replace(Path, "{text}", text, 1)
		Context("Successfully - Get user input example", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"text":"xxx"}`)
			}))
			It("Succeed to call GetExample", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("ListExamples(listExamplesOptions *ListExamplesOptions)", func() {
		listExamplesPath := "/v1/workspaces/{workspace_id}/intents/{intent}/examples"
		version := "exampleString"
		workspaceID := "exampleString"
		intent := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(listExamplesPath, "{workspace_id}", workspaceID, 1)
		Path = strings.Replace(Path, "{intent}", intent, 1)
		Context("Successfully - List user input examples", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"examples":[]}`)
			}))
			It("Succeed to call ListExamples", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("UpdateExample(updateExampleOptions *UpdateExampleOptions)", func() {
		updateExamplePath := "/v1/workspaces/{workspace_id}/intents/{intent}/examples/{text}"
		version := "exampleString"
		workspaceID := "exampleString"
		intent := "exampleString"
		text := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(updateExamplePath, "{workspace_id}", workspaceID, 1)
		Path = strings.Replace(Path, "{intent}", intent, 1)
		Path = strings.Replace(Path, "{text}", text, 1)
		Context("Successfully - Update user input example", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"text":"xxx"}`)
			}))
			It("Succeed to call UpdateExample", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("CreateCounterexample(createCounterexampleOptions *CreateCounterexampleOptions)", func() {
		createCounterexamplePath := "/v1/workspaces/{workspace_id}/counterexamples"
		version := "exampleString"
		workspaceID := "exampleString"
		text := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(createCounterexamplePath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - Create counterexample", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"text":"xxx"}`)
			}))
			It("Succeed to call CreateCounterexample", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("DeleteCounterexample(deleteCounterexampleOptions *DeleteCounterexampleOptions)", func() {
		deleteCounterexamplePath := "/v1/workspaces/{workspace_id}/counterexamples/{text}"
		version := "exampleString"
		workspaceID := "exampleString"
		text := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(deleteCounterexamplePath, "{workspace_id}", workspaceID, 1)
		Path = strings.Replace(Path, "{text}", text, 1)
		Context("Successfully - Delete counterexample", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
			}))
			It("Succeed to call DeleteCounterexample", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("GetCounterexample(getCounterexampleOptions *GetCounterexampleOptions)", func() {
		getCounterexamplePath := "/v1/workspaces/{workspace_id}/counterexamples/{text}"
		version := "exampleString"
		workspaceID := "exampleString"
		text := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(getCounterexamplePath, "{workspace_id}", workspaceID, 1)
		Path = strings.Replace(Path, "{text}", text, 1)
		Context("Successfully - Get counterexample", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"text":"xxx"}`)
			}))
			It("Succeed to call GetCounterexample", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("ListCounterexamples(listCounterexamplesOptions *ListCounterexamplesOptions)", func() {
		listCounterexamplesPath := "/v1/workspaces/{workspace_id}/counterexamples"
		version := "exampleString"
		workspaceID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(listCounterexamplesPath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - List counterexamples", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"counterexamples":[]}`)
			}))
			It("Succeed to call ListCounterexamples", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("UpdateCounterexample(updateCounterexampleOptions *UpdateCounterexampleOptions)", func() {
		updateCounterexamplePath := "/v1/workspaces/{workspace_id}/counterexamples/{text}"
		version := "exampleString"
		workspaceID := "exampleString"
		text := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(updateCounterexamplePath, "{workspace_id}", workspaceID, 1)
		Path = strings.Replace(Path, "{text}", text, 1)
		Context("Successfully - Update counterexample", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"text":"xxx"}`)
			}))
			It("Succeed to call UpdateCounterexample", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("CreateEntity(createEntityOptions *CreateEntityOptions)", func() {
		createEntityPath := "/v1/workspaces/{workspace_id}/entities"
		version := "exampleString"
		workspaceID := "exampleString"
		entity := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(createEntityPath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - Create entity", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"entity":"xxx"}`)
			}))
			It("Succeed to call CreateEntity", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("DeleteEntity(deleteEntityOptions *DeleteEntityOptions)", func() {
		deleteEntityPath := "/v1/workspaces/{workspace_id}/entities/{entity}"
		version := "exampleString"
		workspaceID := "exampleString"
		entity := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(deleteEntityPath, "{workspace_id}", workspaceID, 1)
		Path = strings.Replace(Path, "{entity}", entity, 1)
		Context("Successfully - Delete entity", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
			}))
			It("Succeed to call DeleteEntity", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("GetEntity(getEntityOptions *GetEntityOptions)", func() {
		getEntityPath := "/v1/workspaces/{workspace_id}/entities/{entity}"
		version := "exampleString"
		workspaceID := "exampleString"
		entity := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(getEntityPath, "{workspace_id}", workspaceID, 1)
		Path = strings.Replace(Path, "{entity}", entity, 1)
		Context("Successfully - Get entity", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"entity":"xxx"}`)
			}))
			It("Succeed to call GetEntity", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("ListEntities(listEntitiesOptions *ListEntitiesOptions)", func() {
		listEntitiesPath := "/v1/workspaces/{workspace_id}/entities"
		version := "exampleString"
		workspaceID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(listEntitiesPath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - List entities", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"entities":"xxx"}`)
			}))
			It("Succeed to call ListEntities", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("UpdateEntity(updateEntityOptions *UpdateEntityOptions)", func() {
		updateEntityPath := "/v1/workspaces/{workspace_id}/entities/{entity}"
		version := "exampleString"
		workspaceID := "exampleString"
		entity := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(updateEntityPath, "{workspace_id}", workspaceID, 1)
		Path = strings.Replace(Path, "{entity}", entity, 1)
		Context("Successfully - Update entity", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"entity":"xxx"}`)
			}))
			It("Succeed to call UpdateEntity", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("ListMentions(listMentionsOptions *ListMentionsOptions)", func() {
		listMentionsPath := "/v1/workspaces/{workspace_id}/entities/{entity}/mentions"
		version := "exampleString"
		workspaceID := "exampleString"
		entity := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(listMentionsPath, "{workspace_id}", workspaceID, 1)
		Path = strings.Replace(Path, "{entity}", entity, 1)
		Context("Successfully - List entity mentions", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call ListMentions", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("CreateValue(createValueOptions *CreateValueOptions)", func() {
		createValuePath := "/v1/workspaces/{workspace_id}/entities/{entity}/values"
		version := "exampleString"
		workspaceID := "exampleString"
		entity := "exampleString"
		value := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(createValuePath, "{workspace_id}", workspaceID, 1)
		Path = strings.Replace(Path, "{entity}", entity, 1)
		Context("Successfully - Add entity value", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"value":"xxx"}`)
			}))
			It("Succeed to call CreateValue", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("DeleteValue(deleteValueOptions *DeleteValueOptions)", func() {
		deleteValuePath := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}"
		version := "exampleString"
		workspaceID := "exampleString"
		entity := "exampleString"
		value := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(deleteValuePath, "{workspace_id}", workspaceID, 1)
		Path = strings.Replace(Path, "{entity}", entity, 1)
		Path = strings.Replace(Path, "{value}", value, 1)
		Context("Successfully - Delete entity value", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
			}))
			It("Succeed to call DeleteValue", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("GetValue(getValueOptions *GetValueOptions)", func() {
		getValuePath := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}"
		version := "exampleString"
		workspaceID := "exampleString"
		entity := "exampleString"
		value := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(getValuePath, "{workspace_id}", workspaceID, 1)
		Path = strings.Replace(Path, "{entity}", entity, 1)
		Path = strings.Replace(Path, "{value}", value, 1)
		Context("Successfully - Get entity value", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"value":"xxx"}`)
			}))
			It("Succeed to call GetValue", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("ListValues(listValuesOptions *ListValuesOptions)", func() {
		listValuesPath := "/v1/workspaces/{workspace_id}/entities/{entity}/values"
		version := "exampleString"
		workspaceID := "exampleString"
		entity := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(listValuesPath, "{workspace_id}", workspaceID, 1)
		Path = strings.Replace(Path, "{entity}", entity, 1)
		Context("Successfully - List entity values", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"values":[]}`)
			}))
			It("Succeed to call ListValues", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("UpdateValue(updateValueOptions *UpdateValueOptions)", func() {
		updateValuePath := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}"
		version := "exampleString"
		workspaceID := "exampleString"
		entity := "exampleString"
		value := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(updateValuePath, "{workspace_id}", workspaceID, 1)
		Path = strings.Replace(Path, "{entity}", entity, 1)
		Path = strings.Replace(Path, "{value}", value, 1)
		Context("Successfully - Update entity value", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"value":"xxx"}`)
			}))
			It("Succeed to call UpdateValue", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("CreateSynonym(createSynonymOptions *CreateSynonymOptions)", func() {
		createSynonymPath := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}/synonyms"
		version := "exampleString"
		workspaceID := "exampleString"
		entity := "exampleString"
		value := "exampleString"
		synonym := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(createSynonymPath, "{workspace_id}", workspaceID, 1)
		Path = strings.Replace(Path, "{entity}", entity, 1)
		Path = strings.Replace(Path, "{value}", value, 1)
		Context("Successfully - Add entity value synonym", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"synonym":"xxx"}`)
			}))
			It("Succeed to call CreateSynonym", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("DeleteSynonym(deleteSynonymOptions *DeleteSynonymOptions)", func() {
		deleteSynonymPath := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}/synonyms/{synonym}"
		version := "exampleString"
		workspaceID := "exampleString"
		entity := "exampleString"
		value := "exampleString"
		synonym := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(deleteSynonymPath, "{workspace_id}", workspaceID, 1)
		Path = strings.Replace(Path, "{entity}", entity, 1)
		Path = strings.Replace(Path, "{value}", value, 1)
		Path = strings.Replace(Path, "{synonym}", synonym, 1)
		Context("Successfully - Delete entity value synonym", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
			}))
			It("Succeed to call DeleteSynonym", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("GetSynonym(getSynonymOptions *GetSynonymOptions)", func() {
		getSynonymPath := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}/synonyms/{synonym}"
		version := "exampleString"
		workspaceID := "exampleString"
		entity := "exampleString"
		value := "exampleString"
		synonym := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(getSynonymPath, "{workspace_id}", workspaceID, 1)
		Path = strings.Replace(Path, "{entity}", entity, 1)
		Path = strings.Replace(Path, "{value}", value, 1)
		Path = strings.Replace(Path, "{synonym}", synonym, 1)
		Context("Successfully - Get entity value synonym", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"synonym":"xxx"}`)
			}))
			It("Succeed to call GetSynonym", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("ListSynonyms(listSynonymsOptions *ListSynonymsOptions)", func() {
		listSynonymsPath := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}/synonyms"
		version := "exampleString"
		workspaceID := "exampleString"
		entity := "exampleString"
		value := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(listSynonymsPath, "{workspace_id}", workspaceID, 1)
		Path = strings.Replace(Path, "{entity}", entity, 1)
		Path = strings.Replace(Path, "{value}", value, 1)
		Context("Successfully - List entity value synonyms", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"synonyms":[]}`)
			}))
			It("Succeed to call ListSynonyms", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("UpdateSynonym(updateSynonymOptions *UpdateSynonymOptions)", func() {
		updateSynonymPath := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}/synonyms/{synonym}"
		version := "exampleString"
		workspaceID := "exampleString"
		entity := "exampleString"
		value := "exampleString"
		synonym := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(updateSynonymPath, "{workspace_id}", workspaceID, 1)
		Path = strings.Replace(Path, "{entity}", entity, 1)
		Path = strings.Replace(Path, "{value}", value, 1)
		Path = strings.Replace(Path, "{synonym}", synonym, 1)
		Context("Successfully - Update entity value synonym", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"synonym":"xxx"}`)
			}))
			It("Succeed to call UpdateSynonym", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("CreateDialogNode(createDialogNodeOptions *CreateDialogNodeOptions)", func() {
		createDialogNodePath := "/v1/workspaces/{workspace_id}/dialog_nodes"
		version := "exampleString"
		workspaceID := "exampleString"
		dialogNode := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(createDialogNodePath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - Create dialog node", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"dialog_node":"xxx"}`)
			}))
			It("Succeed to call CreateDialogNode", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("DeleteDialogNode(deleteDialogNodeOptions *DeleteDialogNodeOptions)", func() {
		deleteDialogNodePath := "/v1/workspaces/{workspace_id}/dialog_nodes/{dialog_node}"
		version := "exampleString"
		workspaceID := "exampleString"
		dialogNode := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(deleteDialogNodePath, "{workspace_id}", workspaceID, 1)
		Path = strings.Replace(Path, "{dialog_node}", dialogNode, 1)
		Context("Successfully - Delete dialog node", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
			}))
			It("Succeed to call DeleteDialogNode", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("GetDialogNode(getDialogNodeOptions *GetDialogNodeOptions)", func() {
		getDialogNodePath := "/v1/workspaces/{workspace_id}/dialog_nodes/{dialog_node}"
		version := "exampleString"
		workspaceID := "exampleString"
		dialogNode := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(getDialogNodePath, "{workspace_id}", workspaceID, 1)
		Path = strings.Replace(Path, "{dialog_node}", dialogNode, 1)
		Context("Successfully - Get dialog node", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"dialog_node":"xxx"}`)
			}))
			It("Succeed to call GetDialogNode", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("ListDialogNodes(listDialogNodesOptions *ListDialogNodesOptions)", func() {
		listDialogNodesPath := "/v1/workspaces/{workspace_id}/dialog_nodes"
		version := "exampleString"
		workspaceID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(listDialogNodesPath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - List dialog nodes", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"dialog_nodes":[]}`)
			}))
			It("Succeed to call ListDialogNodes", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("UpdateDialogNode(updateDialogNodeOptions *UpdateDialogNodeOptions)", func() {
		updateDialogNodePath := "/v1/workspaces/{workspace_id}/dialog_nodes/{dialog_node}"
		version := "exampleString"
		workspaceID := "exampleString"
		dialogNode := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(updateDialogNodePath, "{workspace_id}", workspaceID, 1)
		Path = strings.Replace(Path, "{dialog_node}", dialogNode, 1)
		Context("Successfully - Update dialog node", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"dialog_node":"xxx"}`)
			}))
			It("Succeed to call UpdateDialogNode", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("ListAllLogs(listAllLogsOptions *ListAllLogsOptions)", func() {
		listAllLogsPath := "/v1/logs"
		version := "exampleString"
		filter := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - List log events in all workspaces", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(listAllLogsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"logs":[]}`)
			}))
			It("Succeed to call ListAllLogs", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("ListLogs(listLogsOptions *ListLogsOptions)", func() {
		listLogsPath := "/v1/workspaces/{workspace_id}/logs"
		version := "exampleString"
		workspaceID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(listLogsPath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - List log events in a workspace", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"logs":[]}`)
			}))
			It("Succeed to call ListLogs", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("DeleteUserData(deleteUserDataOptions *DeleteUserDataOptions)", func() {
		deleteUserDataPath := "/v1/user_data"
		version := "exampleString"
		customerID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Delete labeled data", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(deleteUserDataPath + "?customer_id=" + customerID + "&version=" + version))
				Expect(req.URL.Path).To(Equal(deleteUserDataPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
			}))
			It("Succeed to call DeleteUserData", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
