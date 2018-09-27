package toneanalyzerv3_test

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/ibm-watson/go-sdk/toneanalyzerv3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
)

var _ = Describe("ToneAnalyzerV3", func() {
	Describe("Get credentials from VCAP", func() {
		version := "exampleString"
		username := "hyphenated-user"
		password := "hyphenated-pass"
		VCAPservices := cfenv.Services{
			"conversation": {
				{
					Name: "tone_analyzer",
					Tags: []string{},
					Credentials: map[string]interface{}{
						"url":      "https://gateway.watsonplatform.net/tone-analyzer/api",
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

				testService, testServiceErr := toneanalyzerv3.NewToneAnalyzerV3(&toneanalyzerv3.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				testService.(testService.NewOptions())
			})
		})
	})
	Describe("Tone(toneOptions *ToneOptions)", func() {
		TonePath := "/v3/tone"
		version := "exampleString"
		ContentType := "exampleString"
		ToneOptions := testService.NewToneOptions(ContentType)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Analyze general tone", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(TonePath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(TonePath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call Tone", func() {
				defer testServer.Close()

				testService, testServiceErr := toneanalyzerv3.NewToneAnalyzerV3(&toneanalyzerv3.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.Tone(ToneOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := toneanalyzerv3.GetToneResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ToneChat(toneChatOptions *ToneChatOptions)", func() {
		ToneChatPath := "/v3/tone_chat"
		version := "exampleString"
		Utterances := []toneanalyzerv3.Utterance{}
		ToneChatOptions := testService.NewToneChatOptions(Utterances)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Analyze customer engagement tone", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(ToneChatPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(ToneChatPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call ToneChat", func() {
				defer testServer.Close()

				testService, testServiceErr := toneanalyzerv3.NewToneAnalyzerV3(&toneanalyzerv3.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.ToneChat(ToneChatOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := toneanalyzerv3.GetToneChatResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
})
