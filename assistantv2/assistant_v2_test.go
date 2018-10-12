package assistantv2_test

import (
	"encoding/base64"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/assistantv2"
	"net/http"
	"net/http/httptest"
	"strings"
)

var _ = Describe("AssistantV2", func() {
	Describe("CreateSession(createSessionOptions *CreateSessionOptions)", func() {
		createSessionPath := "/v2/assistants/{assistant_id}/sessions"
		version := "exampleString"
		assistantID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(createSessionPath, "{assistant_id}", assistantID, 1)
		Context("Successfully - Create a session", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"session_id":"xxx"}`)
			}))
			It("Succeed to call CreateSession", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
		assistantID := "exampleString"
		sessionID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(deleteSessionPath, "{assistant_id}", assistantID, 1)
		Path = strings.Replace(Path, "{session_id}", sessionID, 1)
		Context("Successfully - Delete session", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call DeleteSession", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
		assistantID := "exampleString"
		sessionID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(messagePath, "{assistant_id}", assistantID, 1)
		Path = strings.Replace(Path, "{session_id}", sessionID, 1)
		Context("Successfully - Send user input to assistant", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"output": {"generic": []}}`)
			}))
			It("Succeed to call Message", func() {
				defer testServer.Close()

				testService, testServiceErr := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
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
