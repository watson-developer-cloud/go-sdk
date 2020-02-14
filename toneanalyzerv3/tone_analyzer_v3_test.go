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

package toneanalyzerv3_test

import (
	"fmt"
	"github.com/IBM/go-sdk-core/v3/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/toneanalyzerv3"
	"net/http"
	"net/http/httptest"
)

var _ = Describe(`ToneAnalyzerV3`, func() {
	Describe(`Tone(toneOptions *ToneOptions)`, func() {
		tonePath := "/v3/tone"
		version := "exampleString"
		bearerToken := "0ui9876453"
		Context(`Successfully - Analyze general tone`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(tonePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"document_tone": {}}`)
			}))
			It(`Succeed to call Tone`, func() {
				defer testServer.Close()

				testService, testServiceErr := toneanalyzerv3.NewToneAnalyzerV3(&toneanalyzerv3.ToneAnalyzerV3Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.Tone(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				toneOptions := testService.NewToneOptions()
				result, response, operationErr = testService.Tone(toneOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ToneChat(toneChatOptions *ToneChatOptions)`, func() {
		toneChatPath := "/v3/tone_chat"
		version := "exampleString"
		bearerToken := "0ui9876453"
		utterances := []toneanalyzerv3.Utterance{}
		Context(`Successfully - Analyze customer-engagement tone`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(toneChatPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"utterances_tone": []}`)
			}))
			It(`Succeed to call ToneChat`, func() {
				defer testServer.Close()

				testService, testServiceErr := toneanalyzerv3.NewToneAnalyzerV3(&toneanalyzerv3.ToneAnalyzerV3Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ToneChat(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				toneChatOptions := testService.NewToneChatOptions(utterances)
				result, response, operationErr = testService.ToneChat(toneChatOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("Model constructor tests", func() {
		Context("with a sample service", func() {
			version := "1970-01-01"
			testService, _ := toneanalyzerv3.NewToneAnalyzerV3(&toneanalyzerv3.ToneAnalyzerV3Options{
				URL:           "http://toneanalyzerv3modelgenerator.com",
				Version:       version,
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It("should call NewToneInput successfully", func() {
				text := "exampleString"
				model, err := testService.NewToneInput(text)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It("should call NewUtterance successfully", func() {
				text := "exampleString"
				model, err := testService.NewUtterance(text)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
})
