package toneanalyzerv3_test

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/toneanalyzerv3"
)

var _ = Describe("ToneAnalyzerV3", func() {
	Describe("Tone(toneOptions *ToneOptions)", func() {
		tonePath := "/v3/tone"
		version := "exampleString"
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
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"document_tone":"subtle"}`)
			}))
			It("Succeed to call Tone", func() {
				defer testServer.Close()

				testService, testServiceErr := toneanalyzerv3.NewToneAnalyzerV3(&toneanalyzerv3.ToneAnalyzerV3Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.Tone(nil)
				Expect(returnValueErr).NotTo(BeNil())

				toneOptions := testService.
					NewToneOptions().
					SetBody("I am feeling well today").
					SetContentType(toneanalyzerv3.ToneOptions_ContentType_TextPlain)
				returnValue, returnValueErr = testService.Tone(toneOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetToneResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ToneChat(toneChatOptions *ToneChatOptions)", func() {
		toneChatPath := "/v3/tone_chat"
		version := "exampleString"
		utterances := []toneanalyzerv3.Utterance{}
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
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"utterances_tone":[]}`)
			}))
			It("Succeed to call ToneChat", func() {
				defer testServer.Close()

				testService, testServiceErr := toneanalyzerv3.NewToneAnalyzerV3(&toneanalyzerv3.ToneAnalyzerV3Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ToneChat(nil)
				Expect(returnValueErr).NotTo(BeNil())

				toneChatOptions := testService.NewToneChatOptions(utterances)
				returnValue, returnValueErr = testService.ToneChat(toneChatOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetToneChatResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
})
