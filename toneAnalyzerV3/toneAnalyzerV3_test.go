package toneAnalyzerV3_test

import (
	"go-sdk/toneAnalyzerV3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"encoding/base64"
	"net/http/httptest"
	"net/http"
	"fmt"
	"os"
	"encoding/json"
	"github.com/cloudfoundry-community/go-cfenv"
)

var _ = Describe("ToneAnalyzerV3", func() {
	Describe("Get credentials from VCAP", func() {
		version := "exampleString"
		username := "hyphenated-user"
		password := "hyphenated-pass"
		VCAPservices := cfenv.Services{
			"conversation" : {
				{
					Name: "tone_analyzer",
					Tags: []string{},
					Credentials: map[string]interface{}{
						"url": "https://gateway.watsonplatform.net/tone-analyzer/api",
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
		Context("Successfully - Create ToneAnalyzerV3 with VCAP credentials", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
			}))
			It("Succeed to create ToneAnalyzerV3", func() {
				defer testServer.Close()

				testService, testServiceErr := toneAnalyzerV3.NewToneAnalyzerV3(&toneAnalyzerV3.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				testService.Tone(toneAnalyzerV3.NewToneOptionsForPlain("exampleString"))
			})
		})
	})
	Describe("Tone(options *ToneOptions)", func() {
		tonePath := "/v3/tone"
        version := "exampleString"
        toneInput := toneAnalyzerV3.ToneInput{}
        toneOptions := toneAnalyzerV3.NewToneOptionsForToneInput(toneInput)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Analyze general tone", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(tonePath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(tonePath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call Tone", func() {
				defer testServer.Close()

				testService, testServiceErr := toneAnalyzerV3.NewToneAnalyzerV3(&toneAnalyzerV3.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.Tone(toneOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := toneAnalyzerV3.GetToneResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ToneChat(options *ToneChatOptions)", func() {
		toneChatPath := "/v3/tone_chat"
        version := "exampleString"
        utterances := []toneAnalyzerV3.Utterance{}
        toneChatOptions := toneAnalyzerV3.NewToneChatOptions(utterances)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Analyze customer engagement tone", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(toneChatPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(toneChatPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call ToneChat", func() {
				defer testServer.Close()

				testService, testServiceErr := toneAnalyzerV3.NewToneAnalyzerV3(&toneAnalyzerV3.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.ToneChat(toneChatOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := toneAnalyzerV3.GetToneChatResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
})
