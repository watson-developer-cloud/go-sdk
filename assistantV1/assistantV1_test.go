package assistantV1_test

import (
	"go-sdk/assistantV1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"encoding/base64"
	"net/http/httptest"
	"net/http"
	"strings"
	"fmt"
	"os"
	"encoding/json"
	"github.com/cloudfoundry-community/go-cfenv"
)

var _ = Describe("AssistantV1", func() {
	Describe("Get credentials from VCAP", func() {
		version := "exampleString"
		username := "hyphenated-user"
		password := "hyphenated-pass"
		VCAPservices := cfenv.Services{
			"conversation" : {
				{
					Name: "conversation",
					Tags: []string{},
					Credentials: map[string]interface{}{
						"url": "https://gateway.watsonplatform.net/assistant/api",
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

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				testService.ListWorkspaces(assistantV1.NewListWorkspacesOptions())
			})
		})
	})
	Describe("Message(options *MessageOptions)", func() {
		messagePath := "/v1/workspaces/{workspace_id}/message"
        version := "exampleString"
        workspaceID := "exampleString"
        messageOptions := assistantV1.NewMessageOptions(workspaceID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        messagePath = strings.Replace(messagePath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - Get response to user input", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(messagePath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(messagePath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call Message", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.Message(messageOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := assistantV1.GetMessageResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateWorkspace(options *CreateWorkspaceOptions)", func() {
		createWorkspacePath := "/v1/workspaces"
        version := "exampleString"
        createWorkspaceOptions := assistantV1.NewCreateWorkspaceOptions()
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
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call CreateWorkspace", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.CreateWorkspace(createWorkspaceOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := assistantV1.GetCreateWorkspaceResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteWorkspace(options *DeleteWorkspaceOptions)", func() {
		deleteWorkspacePath := "/v1/workspaces/{workspace_id}"
        version := "exampleString"
        workspaceID := "exampleString"
        deleteWorkspaceOptions := assistantV1.NewDeleteWorkspaceOptions(workspaceID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        deleteWorkspacePath = strings.Replace(deleteWorkspacePath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - Delete workspace", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(deleteWorkspacePath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(deleteWorkspacePath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call DeleteWorkspace", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.DeleteWorkspace(deleteWorkspaceOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("GetWorkspace(options *GetWorkspaceOptions)", func() {
		getWorkspacePath := "/v1/workspaces/{workspace_id}"
        version := "exampleString"
        workspaceID := "exampleString"
        getWorkspaceOptions := assistantV1.NewGetWorkspaceOptions(workspaceID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        getWorkspacePath = strings.Replace(getWorkspacePath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - Get information about a workspace", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(getWorkspacePath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(getWorkspacePath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call GetWorkspace", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.GetWorkspace(getWorkspaceOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := assistantV1.GetGetWorkspaceResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListWorkspaces(options *ListWorkspacesOptions)", func() {
		listWorkspacesPath := "/v1/workspaces"
        version := "exampleString"
        listWorkspacesOptions := assistantV1.NewListWorkspacesOptions()
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
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call ListWorkspaces", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.ListWorkspaces(listWorkspacesOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := assistantV1.GetListWorkspacesResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("UpdateWorkspace(options *UpdateWorkspaceOptions)", func() {
		updateWorkspacePath := "/v1/workspaces/{workspace_id}"
        version := "exampleString"
        workspaceID := "exampleString"
        updateWorkspaceOptions := assistantV1.NewUpdateWorkspaceOptions(workspaceID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        updateWorkspacePath = strings.Replace(updateWorkspacePath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - Update workspace", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(updateWorkspacePath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(updateWorkspacePath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call UpdateWorkspace", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.UpdateWorkspace(updateWorkspaceOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := assistantV1.GetUpdateWorkspaceResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateIntent(options *CreateIntentOptions)", func() {
		createIntentPath := "/v1/workspaces/{workspace_id}/intents"
        version := "exampleString"
        workspaceID := "exampleString"
        intent := "exampleString"
        createIntentOptions := assistantV1.NewCreateIntentOptions(workspaceID, intent)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        createIntentPath = strings.Replace(createIntentPath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - Create intent", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(createIntentPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(createIntentPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call CreateIntent", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.CreateIntent(createIntentOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := assistantV1.GetCreateIntentResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteIntent(options *DeleteIntentOptions)", func() {
		deleteIntentPath := "/v1/workspaces/{workspace_id}/intents/{intent}"
        version := "exampleString"
        workspaceID := "exampleString"
        intent := "exampleString"
        deleteIntentOptions := assistantV1.NewDeleteIntentOptions(workspaceID, intent)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        deleteIntentPath = strings.Replace(deleteIntentPath, "{workspace_id}", workspaceID, 1)
        deleteIntentPath = strings.Replace(deleteIntentPath, "{intent}", intent, 1)
		Context("Successfully - Delete intent", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(deleteIntentPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(deleteIntentPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call DeleteIntent", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.DeleteIntent(deleteIntentOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("GetIntent(options *GetIntentOptions)", func() {
		getIntentPath := "/v1/workspaces/{workspace_id}/intents/{intent}"
        version := "exampleString"
        workspaceID := "exampleString"
        intent := "exampleString"
        getIntentOptions := assistantV1.NewGetIntentOptions(workspaceID, intent)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        getIntentPath = strings.Replace(getIntentPath, "{workspace_id}", workspaceID, 1)
        getIntentPath = strings.Replace(getIntentPath, "{intent}", intent, 1)
		Context("Successfully - Get intent", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(getIntentPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(getIntentPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call GetIntent", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.GetIntent(getIntentOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := assistantV1.GetGetIntentResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListIntents(options *ListIntentsOptions)", func() {
		listIntentsPath := "/v1/workspaces/{workspace_id}/intents"
        version := "exampleString"
        workspaceID := "exampleString"
        listIntentsOptions := assistantV1.NewListIntentsOptions(workspaceID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        listIntentsPath = strings.Replace(listIntentsPath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - List intents", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(listIntentsPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(listIntentsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call ListIntents", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.ListIntents(listIntentsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := assistantV1.GetListIntentsResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("UpdateIntent(options *UpdateIntentOptions)", func() {
		updateIntentPath := "/v1/workspaces/{workspace_id}/intents/{intent}"
        version := "exampleString"
        workspaceID := "exampleString"
        intent := "exampleString"
        updateIntentOptions := assistantV1.NewUpdateIntentOptions(workspaceID, intent)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        updateIntentPath = strings.Replace(updateIntentPath, "{workspace_id}", workspaceID, 1)
        updateIntentPath = strings.Replace(updateIntentPath, "{intent}", intent, 1)
		Context("Successfully - Update intent", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(updateIntentPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(updateIntentPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call UpdateIntent", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.UpdateIntent(updateIntentOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := assistantV1.GetUpdateIntentResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateExample(options *CreateExampleOptions)", func() {
		createExamplePath := "/v1/workspaces/{workspace_id}/intents/{intent}/examples"
        version := "exampleString"
        workspaceID := "exampleString"
        intent := "exampleString"
        text := "exampleString"
        createExampleOptions := assistantV1.NewCreateExampleOptions(workspaceID, intent, text)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        createExamplePath = strings.Replace(createExamplePath, "{workspace_id}", workspaceID, 1)
        createExamplePath = strings.Replace(createExamplePath, "{intent}", intent, 1)
		Context("Successfully - Create user input example", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(createExamplePath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(createExamplePath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call CreateExample", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.CreateExample(createExampleOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := assistantV1.GetCreateExampleResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteExample(options *DeleteExampleOptions)", func() {
		deleteExamplePath := "/v1/workspaces/{workspace_id}/intents/{intent}/examples/{text}"
        version := "exampleString"
        workspaceID := "exampleString"
        intent := "exampleString"
        text := "exampleString"
        deleteExampleOptions := assistantV1.NewDeleteExampleOptions(workspaceID, intent, text)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        deleteExamplePath = strings.Replace(deleteExamplePath, "{workspace_id}", workspaceID, 1)
        deleteExamplePath = strings.Replace(deleteExamplePath, "{intent}", intent, 1)
        deleteExamplePath = strings.Replace(deleteExamplePath, "{text}", text, 1)
		Context("Successfully - Delete user input example", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(deleteExamplePath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(deleteExamplePath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call DeleteExample", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.DeleteExample(deleteExampleOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("GetExample(options *GetExampleOptions)", func() {
		getExamplePath := "/v1/workspaces/{workspace_id}/intents/{intent}/examples/{text}"
        version := "exampleString"
        workspaceID := "exampleString"
        intent := "exampleString"
        text := "exampleString"
        getExampleOptions := assistantV1.NewGetExampleOptions(workspaceID, intent, text)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        getExamplePath = strings.Replace(getExamplePath, "{workspace_id}", workspaceID, 1)
        getExamplePath = strings.Replace(getExamplePath, "{intent}", intent, 1)
        getExamplePath = strings.Replace(getExamplePath, "{text}", text, 1)
		Context("Successfully - Get user input example", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(getExamplePath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(getExamplePath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call GetExample", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.GetExample(getExampleOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := assistantV1.GetGetExampleResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListExamples(options *ListExamplesOptions)", func() {
		listExamplesPath := "/v1/workspaces/{workspace_id}/intents/{intent}/examples"
        version := "exampleString"
        workspaceID := "exampleString"
        intent := "exampleString"
        listExamplesOptions := assistantV1.NewListExamplesOptions(workspaceID, intent)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        listExamplesPath = strings.Replace(listExamplesPath, "{workspace_id}", workspaceID, 1)
        listExamplesPath = strings.Replace(listExamplesPath, "{intent}", intent, 1)
		Context("Successfully - List user input examples", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(listExamplesPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(listExamplesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call ListExamples", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.ListExamples(listExamplesOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := assistantV1.GetListExamplesResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("UpdateExample(options *UpdateExampleOptions)", func() {
		updateExamplePath := "/v1/workspaces/{workspace_id}/intents/{intent}/examples/{text}"
        version := "exampleString"
        workspaceID := "exampleString"
        intent := "exampleString"
        text := "exampleString"
        updateExampleOptions := assistantV1.NewUpdateExampleOptions(workspaceID, intent, text)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        updateExamplePath = strings.Replace(updateExamplePath, "{workspace_id}", workspaceID, 1)
        updateExamplePath = strings.Replace(updateExamplePath, "{intent}", intent, 1)
        updateExamplePath = strings.Replace(updateExamplePath, "{text}", text, 1)
		Context("Successfully - Update user input example", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(updateExamplePath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(updateExamplePath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call UpdateExample", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.UpdateExample(updateExampleOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := assistantV1.GetUpdateExampleResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateCounterexample(options *CreateCounterexampleOptions)", func() {
		createCounterexamplePath := "/v1/workspaces/{workspace_id}/counterexamples"
        version := "exampleString"
        workspaceID := "exampleString"
        text := "exampleString"
        createCounterexampleOptions := assistantV1.NewCreateCounterexampleOptions(workspaceID, text)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        createCounterexamplePath = strings.Replace(createCounterexamplePath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - Create counterexample", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(createCounterexamplePath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(createCounterexamplePath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call CreateCounterexample", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.CreateCounterexample(createCounterexampleOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := assistantV1.GetCreateCounterexampleResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteCounterexample(options *DeleteCounterexampleOptions)", func() {
		deleteCounterexamplePath := "/v1/workspaces/{workspace_id}/counterexamples/{text}"
        version := "exampleString"
        workspaceID := "exampleString"
        text := "exampleString"
        deleteCounterexampleOptions := assistantV1.NewDeleteCounterexampleOptions(workspaceID, text)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        deleteCounterexamplePath = strings.Replace(deleteCounterexamplePath, "{workspace_id}", workspaceID, 1)
        deleteCounterexamplePath = strings.Replace(deleteCounterexamplePath, "{text}", text, 1)
		Context("Successfully - Delete counterexample", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(deleteCounterexamplePath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(deleteCounterexamplePath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call DeleteCounterexample", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.DeleteCounterexample(deleteCounterexampleOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("GetCounterexample(options *GetCounterexampleOptions)", func() {
		getCounterexamplePath := "/v1/workspaces/{workspace_id}/counterexamples/{text}"
        version := "exampleString"
        workspaceID := "exampleString"
        text := "exampleString"
        getCounterexampleOptions := assistantV1.NewGetCounterexampleOptions(workspaceID, text)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        getCounterexamplePath = strings.Replace(getCounterexamplePath, "{workspace_id}", workspaceID, 1)
        getCounterexamplePath = strings.Replace(getCounterexamplePath, "{text}", text, 1)
		Context("Successfully - Get counterexample", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(getCounterexamplePath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(getCounterexamplePath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call GetCounterexample", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.GetCounterexample(getCounterexampleOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := assistantV1.GetGetCounterexampleResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListCounterexamples(options *ListCounterexamplesOptions)", func() {
		listCounterexamplesPath := "/v1/workspaces/{workspace_id}/counterexamples"
        version := "exampleString"
        workspaceID := "exampleString"
        listCounterexamplesOptions := assistantV1.NewListCounterexamplesOptions(workspaceID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        listCounterexamplesPath = strings.Replace(listCounterexamplesPath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - List counterexamples", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(listCounterexamplesPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(listCounterexamplesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call ListCounterexamples", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.ListCounterexamples(listCounterexamplesOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := assistantV1.GetListCounterexamplesResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("UpdateCounterexample(options *UpdateCounterexampleOptions)", func() {
		updateCounterexamplePath := "/v1/workspaces/{workspace_id}/counterexamples/{text}"
        version := "exampleString"
        workspaceID := "exampleString"
        text := "exampleString"
        updateCounterexampleOptions := assistantV1.NewUpdateCounterexampleOptions(workspaceID, text)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        updateCounterexamplePath = strings.Replace(updateCounterexamplePath, "{workspace_id}", workspaceID, 1)
        updateCounterexamplePath = strings.Replace(updateCounterexamplePath, "{text}", text, 1)
		Context("Successfully - Update counterexample", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(updateCounterexamplePath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(updateCounterexamplePath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call UpdateCounterexample", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.UpdateCounterexample(updateCounterexampleOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := assistantV1.GetUpdateCounterexampleResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateEntity(options *CreateEntityOptions)", func() {
		createEntityPath := "/v1/workspaces/{workspace_id}/entities"
        version := "exampleString"
        workspaceID := "exampleString"
        entity := "exampleString"
        createEntityOptions := assistantV1.NewCreateEntityOptions(workspaceID, entity)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        createEntityPath = strings.Replace(createEntityPath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - Create entity", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(createEntityPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(createEntityPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call CreateEntity", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.CreateEntity(createEntityOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := assistantV1.GetCreateEntityResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteEntity(options *DeleteEntityOptions)", func() {
		deleteEntityPath := "/v1/workspaces/{workspace_id}/entities/{entity}"
        version := "exampleString"
        workspaceID := "exampleString"
        entity := "exampleString"
        deleteEntityOptions := assistantV1.NewDeleteEntityOptions(workspaceID, entity)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        deleteEntityPath = strings.Replace(deleteEntityPath, "{workspace_id}", workspaceID, 1)
        deleteEntityPath = strings.Replace(deleteEntityPath, "{entity}", entity, 1)
		Context("Successfully - Delete entity", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(deleteEntityPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(deleteEntityPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call DeleteEntity", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.DeleteEntity(deleteEntityOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("GetEntity(options *GetEntityOptions)", func() {
		getEntityPath := "/v1/workspaces/{workspace_id}/entities/{entity}"
        version := "exampleString"
        workspaceID := "exampleString"
        entity := "exampleString"
        getEntityOptions := assistantV1.NewGetEntityOptions(workspaceID, entity)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        getEntityPath = strings.Replace(getEntityPath, "{workspace_id}", workspaceID, 1)
        getEntityPath = strings.Replace(getEntityPath, "{entity}", entity, 1)
		Context("Successfully - Get entity", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(getEntityPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(getEntityPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call GetEntity", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.GetEntity(getEntityOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := assistantV1.GetGetEntityResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListEntities(options *ListEntitiesOptions)", func() {
		listEntitiesPath := "/v1/workspaces/{workspace_id}/entities"
        version := "exampleString"
        workspaceID := "exampleString"
        listEntitiesOptions := assistantV1.NewListEntitiesOptions(workspaceID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        listEntitiesPath = strings.Replace(listEntitiesPath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - List entities", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(listEntitiesPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(listEntitiesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call ListEntities", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.ListEntities(listEntitiesOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := assistantV1.GetListEntitiesResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("UpdateEntity(options *UpdateEntityOptions)", func() {
		updateEntityPath := "/v1/workspaces/{workspace_id}/entities/{entity}"
        version := "exampleString"
        workspaceID := "exampleString"
        entity := "exampleString"
        updateEntityOptions := assistantV1.NewUpdateEntityOptions(workspaceID, entity)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        updateEntityPath = strings.Replace(updateEntityPath, "{workspace_id}", workspaceID, 1)
        updateEntityPath = strings.Replace(updateEntityPath, "{entity}", entity, 1)
		Context("Successfully - Update entity", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(updateEntityPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(updateEntityPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call UpdateEntity", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.UpdateEntity(updateEntityOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := assistantV1.GetUpdateEntityResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListEntityMentions(options *ListEntityMentionsOptions)", func() {
		listEntityMentionsPath := "/v1/workspaces/{workspace_id}/entities/{entity}/mentions"
        version := "exampleString"
        workspaceID := "exampleString"
        entity := "exampleString"
        listEntityMentionsOptions := assistantV1.NewListEntityMentionsOptions(workspaceID, entity)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        listEntityMentionsPath = strings.Replace(listEntityMentionsPath, "{workspace_id}", workspaceID, 1)
        listEntityMentionsPath = strings.Replace(listEntityMentionsPath, "{entity}", entity, 1)
		Context("Successfully - List entity mentions", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(listEntityMentionsPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(listEntityMentionsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call ListEntityMentions", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.ListEntityMentions(listEntityMentionsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := assistantV1.GetListEntityMentionsResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateValue(options *CreateValueOptions)", func() {
		createValuePath := "/v1/workspaces/{workspace_id}/entities/{entity}/values"
        version := "exampleString"
        workspaceID := "exampleString"
        entity := "exampleString"
        value := "exampleString"
        createValueOptions := assistantV1.NewCreateValueOptions(workspaceID, entity, value)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        createValuePath = strings.Replace(createValuePath, "{workspace_id}", workspaceID, 1)
        createValuePath = strings.Replace(createValuePath, "{entity}", entity, 1)
		Context("Successfully - Add entity value", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(createValuePath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(createValuePath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call CreateValue", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.CreateValue(createValueOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := assistantV1.GetCreateValueResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteValue(options *DeleteValueOptions)", func() {
		deleteValuePath := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}"
        version := "exampleString"
        workspaceID := "exampleString"
        entity := "exampleString"
        value := "exampleString"
        deleteValueOptions := assistantV1.NewDeleteValueOptions(workspaceID, entity, value)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        deleteValuePath = strings.Replace(deleteValuePath, "{workspace_id}", workspaceID, 1)
        deleteValuePath = strings.Replace(deleteValuePath, "{entity}", entity, 1)
        deleteValuePath = strings.Replace(deleteValuePath, "{value}", value, 1)
		Context("Successfully - Delete entity value", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(deleteValuePath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(deleteValuePath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call DeleteValue", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.DeleteValue(deleteValueOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("GetValue(options *GetValueOptions)", func() {
		getValuePath := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}"
        version := "exampleString"
        workspaceID := "exampleString"
        entity := "exampleString"
        value := "exampleString"
        getValueOptions := assistantV1.NewGetValueOptions(workspaceID, entity, value)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        getValuePath = strings.Replace(getValuePath, "{workspace_id}", workspaceID, 1)
        getValuePath = strings.Replace(getValuePath, "{entity}", entity, 1)
        getValuePath = strings.Replace(getValuePath, "{value}", value, 1)
		Context("Successfully - Get entity value", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(getValuePath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(getValuePath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call GetValue", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.GetValue(getValueOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := assistantV1.GetGetValueResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListValues(options *ListValuesOptions)", func() {
		listValuesPath := "/v1/workspaces/{workspace_id}/entities/{entity}/values"
        version := "exampleString"
        workspaceID := "exampleString"
        entity := "exampleString"
        listValuesOptions := assistantV1.NewListValuesOptions(workspaceID, entity)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        listValuesPath = strings.Replace(listValuesPath, "{workspace_id}", workspaceID, 1)
        listValuesPath = strings.Replace(listValuesPath, "{entity}", entity, 1)
		Context("Successfully - List entity values", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(listValuesPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(listValuesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call ListValues", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.ListValues(listValuesOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := assistantV1.GetListValuesResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("UpdateValue(options *UpdateValueOptions)", func() {
		updateValuePath := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}"
        version := "exampleString"
        workspaceID := "exampleString"
        entity := "exampleString"
        value := "exampleString"
        updateValueOptions := assistantV1.NewUpdateValueOptions(workspaceID, entity, value)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        updateValuePath = strings.Replace(updateValuePath, "{workspace_id}", workspaceID, 1)
        updateValuePath = strings.Replace(updateValuePath, "{entity}", entity, 1)
        updateValuePath = strings.Replace(updateValuePath, "{value}", value, 1)
		Context("Successfully - Update entity value", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(updateValuePath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(updateValuePath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call UpdateValue", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.UpdateValue(updateValueOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := assistantV1.GetUpdateValueResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateSynonym(options *CreateSynonymOptions)", func() {
		createSynonymPath := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}/synonyms"
        version := "exampleString"
        workspaceID := "exampleString"
        entity := "exampleString"
        value := "exampleString"
        synonym := "exampleString"
        createSynonymOptions := assistantV1.NewCreateSynonymOptions(workspaceID, entity, value, synonym)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        createSynonymPath = strings.Replace(createSynonymPath, "{workspace_id}", workspaceID, 1)
        createSynonymPath = strings.Replace(createSynonymPath, "{entity}", entity, 1)
        createSynonymPath = strings.Replace(createSynonymPath, "{value}", value, 1)
		Context("Successfully - Add entity value synonym", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(createSynonymPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(createSynonymPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call CreateSynonym", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.CreateSynonym(createSynonymOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := assistantV1.GetCreateSynonymResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteSynonym(options *DeleteSynonymOptions)", func() {
		deleteSynonymPath := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}/synonyms/{synonym}"
        version := "exampleString"
        workspaceID := "exampleString"
        entity := "exampleString"
        value := "exampleString"
        synonym := "exampleString"
        deleteSynonymOptions := assistantV1.NewDeleteSynonymOptions(workspaceID, entity, value, synonym)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        deleteSynonymPath = strings.Replace(deleteSynonymPath, "{workspace_id}", workspaceID, 1)
        deleteSynonymPath = strings.Replace(deleteSynonymPath, "{entity}", entity, 1)
        deleteSynonymPath = strings.Replace(deleteSynonymPath, "{value}", value, 1)
        deleteSynonymPath = strings.Replace(deleteSynonymPath, "{synonym}", synonym, 1)
		Context("Successfully - Delete entity value synonym", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(deleteSynonymPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(deleteSynonymPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call DeleteSynonym", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.DeleteSynonym(deleteSynonymOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("GetSynonym(options *GetSynonymOptions)", func() {
		getSynonymPath := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}/synonyms/{synonym}"
        version := "exampleString"
        workspaceID := "exampleString"
        entity := "exampleString"
        value := "exampleString"
        synonym := "exampleString"
        getSynonymOptions := assistantV1.NewGetSynonymOptions(workspaceID, entity, value, synonym)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        getSynonymPath = strings.Replace(getSynonymPath, "{workspace_id}", workspaceID, 1)
        getSynonymPath = strings.Replace(getSynonymPath, "{entity}", entity, 1)
        getSynonymPath = strings.Replace(getSynonymPath, "{value}", value, 1)
        getSynonymPath = strings.Replace(getSynonymPath, "{synonym}", synonym, 1)
		Context("Successfully - Get entity value synonym", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(getSynonymPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(getSynonymPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call GetSynonym", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.GetSynonym(getSynonymOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := assistantV1.GetGetSynonymResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListSynonyms(options *ListSynonymsOptions)", func() {
		listSynonymsPath := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}/synonyms"
        version := "exampleString"
        workspaceID := "exampleString"
        entity := "exampleString"
        value := "exampleString"
        listSynonymsOptions := assistantV1.NewListSynonymsOptions(workspaceID, entity, value)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        listSynonymsPath = strings.Replace(listSynonymsPath, "{workspace_id}", workspaceID, 1)
        listSynonymsPath = strings.Replace(listSynonymsPath, "{entity}", entity, 1)
        listSynonymsPath = strings.Replace(listSynonymsPath, "{value}", value, 1)
		Context("Successfully - List entity value synonyms", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(listSynonymsPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(listSynonymsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call ListSynonyms", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.ListSynonyms(listSynonymsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := assistantV1.GetListSynonymsResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("UpdateSynonym(options *UpdateSynonymOptions)", func() {
		updateSynonymPath := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}/synonyms/{synonym}"
        version := "exampleString"
        workspaceID := "exampleString"
        entity := "exampleString"
        value := "exampleString"
        synonym := "exampleString"
        updateSynonymOptions := assistantV1.NewUpdateSynonymOptions(workspaceID, entity, value, synonym)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        updateSynonymPath = strings.Replace(updateSynonymPath, "{workspace_id}", workspaceID, 1)
        updateSynonymPath = strings.Replace(updateSynonymPath, "{entity}", entity, 1)
        updateSynonymPath = strings.Replace(updateSynonymPath, "{value}", value, 1)
        updateSynonymPath = strings.Replace(updateSynonymPath, "{synonym}", synonym, 1)
		Context("Successfully - Update entity value synonym", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(updateSynonymPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(updateSynonymPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call UpdateSynonym", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.UpdateSynonym(updateSynonymOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := assistantV1.GetUpdateSynonymResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateDialogNode(options *CreateDialogNodeOptions)", func() {
		createDialogNodePath := "/v1/workspaces/{workspace_id}/dialog_nodes"
        version := "exampleString"
        workspaceID := "exampleString"
        dialogNode := "exampleString"
        createDialogNodeOptions := assistantV1.NewCreateDialogNodeOptions(workspaceID, dialogNode)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        createDialogNodePath = strings.Replace(createDialogNodePath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - Create dialog node", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(createDialogNodePath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(createDialogNodePath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call CreateDialogNode", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.CreateDialogNode(createDialogNodeOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := assistantV1.GetCreateDialogNodeResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteDialogNode(options *DeleteDialogNodeOptions)", func() {
		deleteDialogNodePath := "/v1/workspaces/{workspace_id}/dialog_nodes/{dialog_node}"
        version := "exampleString"
        workspaceID := "exampleString"
        dialogNode := "exampleString"
        deleteDialogNodeOptions := assistantV1.NewDeleteDialogNodeOptions(workspaceID, dialogNode)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        deleteDialogNodePath = strings.Replace(deleteDialogNodePath, "{workspace_id}", workspaceID, 1)
        deleteDialogNodePath = strings.Replace(deleteDialogNodePath, "{dialog_node}", dialogNode, 1)
		Context("Successfully - Delete dialog node", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(deleteDialogNodePath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(deleteDialogNodePath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call DeleteDialogNode", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.DeleteDialogNode(deleteDialogNodeOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("GetDialogNode(options *GetDialogNodeOptions)", func() {
		getDialogNodePath := "/v1/workspaces/{workspace_id}/dialog_nodes/{dialog_node}"
        version := "exampleString"
        workspaceID := "exampleString"
        dialogNode := "exampleString"
        getDialogNodeOptions := assistantV1.NewGetDialogNodeOptions(workspaceID, dialogNode)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        getDialogNodePath = strings.Replace(getDialogNodePath, "{workspace_id}", workspaceID, 1)
        getDialogNodePath = strings.Replace(getDialogNodePath, "{dialog_node}", dialogNode, 1)
		Context("Successfully - Get dialog node", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(getDialogNodePath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(getDialogNodePath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call GetDialogNode", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.GetDialogNode(getDialogNodeOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := assistantV1.GetGetDialogNodeResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListDialogNodes(options *ListDialogNodesOptions)", func() {
		listDialogNodesPath := "/v1/workspaces/{workspace_id}/dialog_nodes"
        version := "exampleString"
        workspaceID := "exampleString"
        listDialogNodesOptions := assistantV1.NewListDialogNodesOptions(workspaceID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        listDialogNodesPath = strings.Replace(listDialogNodesPath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - List dialog nodes", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(listDialogNodesPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(listDialogNodesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call ListDialogNodes", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.ListDialogNodes(listDialogNodesOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := assistantV1.GetListDialogNodesResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("UpdateDialogNode(options *UpdateDialogNodeOptions)", func() {
		updateDialogNodePath := "/v1/workspaces/{workspace_id}/dialog_nodes/{dialog_node}"
        version := "exampleString"
        workspaceID := "exampleString"
        dialogNode := "exampleString"
        updateDialogNodeOptions := assistantV1.NewUpdateDialogNodeOptions(workspaceID, dialogNode)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        updateDialogNodePath = strings.Replace(updateDialogNodePath, "{workspace_id}", workspaceID, 1)
        updateDialogNodePath = strings.Replace(updateDialogNodePath, "{dialog_node}", dialogNode, 1)
		Context("Successfully - Update dialog node", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(updateDialogNodePath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(updateDialogNodePath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call UpdateDialogNode", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.UpdateDialogNode(updateDialogNodeOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := assistantV1.GetUpdateDialogNodeResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListAllLogs(options *ListAllLogsOptions)", func() {
		listAllLogsPath := "/v1/logs"
        version := "exampleString"
        filter := "exampleString"
        listAllLogsOptions := assistantV1.NewListAllLogsOptions(filter)
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
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call ListAllLogs", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.ListAllLogs(listAllLogsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := assistantV1.GetListAllLogsResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListLogs(options *ListLogsOptions)", func() {
		listLogsPath := "/v1/workspaces/{workspace_id}/logs"
        version := "exampleString"
        workspaceID := "exampleString"
        listLogsOptions := assistantV1.NewListLogsOptions(workspaceID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        listLogsPath = strings.Replace(listLogsPath, "{workspace_id}", workspaceID, 1)
		Context("Successfully - List log events in a workspace", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(listLogsPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(listLogsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call ListLogs", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.ListLogs(listLogsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := assistantV1.GetListLogsResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteUserData(options *DeleteUserDataOptions)", func() {
		deleteUserDataPath := "/v1/user_data"
        version := "exampleString"
        customerID := "exampleString"
        deleteUserDataOptions := assistantV1.NewDeleteUserDataOptions(customerID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Delete labeled data", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(deleteUserDataPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call DeleteUserData", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantV1.NewAssistantV1(&assistantV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.DeleteUserData(deleteUserDataOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
})
