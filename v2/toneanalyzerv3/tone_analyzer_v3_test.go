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
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/v2/toneanalyzerv3"
)

var _ = Describe(`ToneAnalyzerV3`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		version := "testString"
		It(`Instantiate service client`, func() {
			toneAnalyzerService, serviceErr := toneanalyzerv3.NewToneAnalyzerV3(&toneanalyzerv3.ToneAnalyzerV3Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       core.StringPtr(version),
			})
			Expect(toneAnalyzerService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
			Expect(toneAnalyzerService.Service.IsSSLDisabled()).To(BeFalse())
			toneAnalyzerService.DisableSSLVerification()
			Expect(toneAnalyzerService.Service.IsSSLDisabled()).To(BeTrue())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			toneAnalyzerService, serviceErr := toneanalyzerv3.NewToneAnalyzerV3(&toneanalyzerv3.ToneAnalyzerV3Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(toneAnalyzerService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			toneAnalyzerService, serviceErr := toneanalyzerv3.NewToneAnalyzerV3(&toneanalyzerv3.ToneAnalyzerV3Options{
				URL:     "https://toneanalyzerv3/api",
				Version: core.StringPtr(version),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(toneAnalyzerService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			toneAnalyzerService, serviceErr := toneanalyzerv3.NewToneAnalyzerV3(&toneanalyzerv3.ToneAnalyzerV3Options{})
			Expect(toneAnalyzerService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"TONE_ANALYZER_URL":       "https://toneanalyzerv3/api",
				"TONE_ANALYZER_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				toneAnalyzerService, serviceErr := toneanalyzerv3.NewToneAnalyzerV3(&toneanalyzerv3.ToneAnalyzerV3Options{
					Version: core.StringPtr(version),
				})
				Expect(toneAnalyzerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := toneAnalyzerService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != toneAnalyzerService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(toneAnalyzerService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(toneAnalyzerService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				toneAnalyzerService, serviceErr := toneanalyzerv3.NewToneAnalyzerV3(&toneanalyzerv3.ToneAnalyzerV3Options{
					URL:     "https://testService/api",
					Version: core.StringPtr(version),
				})
				Expect(toneAnalyzerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(toneAnalyzerService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := toneAnalyzerService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != toneAnalyzerService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(toneAnalyzerService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(toneAnalyzerService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				toneAnalyzerService, serviceErr := toneanalyzerv3.NewToneAnalyzerV3(&toneanalyzerv3.ToneAnalyzerV3Options{
					Version: core.StringPtr(version),
				})
				err := toneAnalyzerService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(toneAnalyzerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(toneAnalyzerService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := toneAnalyzerService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != toneAnalyzerService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(toneAnalyzerService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(toneAnalyzerService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"TONE_ANALYZER_URL":       "https://toneanalyzerv3/api",
				"TONE_ANALYZER_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			toneAnalyzerService, serviceErr := toneanalyzerv3.NewToneAnalyzerV3(&toneanalyzerv3.ToneAnalyzerV3Options{
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(toneAnalyzerService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"TONE_ANALYZER_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			toneAnalyzerService, serviceErr := toneanalyzerv3.NewToneAnalyzerV3(&toneanalyzerv3.ToneAnalyzerV3Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(toneAnalyzerService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = toneanalyzerv3.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`Tone(toneOptions *ToneOptions) - Operation response error`, func() {
		version := "testString"
		tonePath := "/v3/tone"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(tonePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Content-Type"]).ToNot(BeNil())
					Expect(req.Header["Content-Type"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					Expect(req.Header["Content-Language"]).ToNot(BeNil())
					Expect(req.Header["Content-Language"][0]).To(Equal(fmt.Sprintf("%v", "en")))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "ar")))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// TODO: Add check for sentences query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke Tone with error: Operation response processing error`, func() {
				toneAnalyzerService, serviceErr := toneanalyzerv3.NewToneAnalyzerV3(&toneanalyzerv3.ToneAnalyzerV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(toneAnalyzerService).ToNot(BeNil())

				// Construct an instance of the ToneInput model
				toneInputModel := new(toneanalyzerv3.ToneInput)
				toneInputModel.Text = core.StringPtr("testString")

				// Construct an instance of the ToneOptions model
				toneOptionsModel := new(toneanalyzerv3.ToneOptions)
				toneOptionsModel.ToneInput = toneInputModel
				toneOptionsModel.ContentType = core.StringPtr("application/json")
				toneOptionsModel.Sentences = core.BoolPtr(true)
				toneOptionsModel.Tones = []string{"emotion"}
				toneOptionsModel.ContentLanguage = core.StringPtr("en")
				toneOptionsModel.AcceptLanguage = core.StringPtr("ar")
				toneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := toneAnalyzerService.Tone(toneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				toneAnalyzerService.EnableRetries(0, 0)
				result, response, operationErr = toneAnalyzerService.Tone(toneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`Tone(toneOptions *ToneOptions)`, func() {
		version := "testString"
		tonePath := "/v3/tone"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(tonePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Content-Type"]).ToNot(BeNil())
					Expect(req.Header["Content-Type"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					Expect(req.Header["Content-Language"]).ToNot(BeNil())
					Expect(req.Header["Content-Language"][0]).To(Equal(fmt.Sprintf("%v", "en")))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "ar")))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// TODO: Add check for sentences query parameter

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"document_tone": {"tones": [{"score": 5, "tone_id": "ToneID", "tone_name": "ToneName"}], "tone_categories": [{"tones": [{"score": 5, "tone_id": "ToneID", "tone_name": "ToneName"}], "category_id": "CategoryID", "category_name": "CategoryName"}], "warning": "Warning"}, "sentences_tone": [{"sentence_id": 10, "text": "Text", "tones": [{"score": 5, "tone_id": "ToneID", "tone_name": "ToneName"}], "tone_categories": [{"tones": [{"score": 5, "tone_id": "ToneID", "tone_name": "ToneName"}], "category_id": "CategoryID", "category_name": "CategoryName"}], "input_from": 9, "input_to": 7}]}`)
				}))
			})
			It(`Invoke Tone successfully`, func() {
				toneAnalyzerService, serviceErr := toneanalyzerv3.NewToneAnalyzerV3(&toneanalyzerv3.ToneAnalyzerV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(toneAnalyzerService).ToNot(BeNil())
				toneAnalyzerService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := toneAnalyzerService.Tone(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ToneInput model
				toneInputModel := new(toneanalyzerv3.ToneInput)
				toneInputModel.Text = core.StringPtr("testString")

				// Construct an instance of the ToneOptions model
				toneOptionsModel := new(toneanalyzerv3.ToneOptions)
				toneOptionsModel.ToneInput = toneInputModel
				toneOptionsModel.ContentType = core.StringPtr("application/json")
				toneOptionsModel.Sentences = core.BoolPtr(true)
				toneOptionsModel.Tones = []string{"emotion"}
				toneOptionsModel.ContentLanguage = core.StringPtr("en")
				toneOptionsModel.AcceptLanguage = core.StringPtr("ar")
				toneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = toneAnalyzerService.Tone(toneOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = toneAnalyzerService.ToneWithContext(ctx, toneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				toneAnalyzerService.DisableRetries()
				result, response, operationErr = toneAnalyzerService.Tone(toneOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = toneAnalyzerService.ToneWithContext(ctx, toneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke Tone with error: Operation request error`, func() {
				toneAnalyzerService, serviceErr := toneanalyzerv3.NewToneAnalyzerV3(&toneanalyzerv3.ToneAnalyzerV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(toneAnalyzerService).ToNot(BeNil())

				// Construct an instance of the ToneInput model
				toneInputModel := new(toneanalyzerv3.ToneInput)
				toneInputModel.Text = core.StringPtr("testString")

				// Construct an instance of the ToneOptions model
				toneOptionsModel := new(toneanalyzerv3.ToneOptions)
				toneOptionsModel.ToneInput = toneInputModel
				toneOptionsModel.ContentType = core.StringPtr("application/json")
				toneOptionsModel.Sentences = core.BoolPtr(true)
				toneOptionsModel.Tones = []string{"emotion"}
				toneOptionsModel.ContentLanguage = core.StringPtr("en")
				toneOptionsModel.AcceptLanguage = core.StringPtr("ar")
				toneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := toneAnalyzerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := toneAnalyzerService.Tone(toneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ToneChat(toneChatOptions *ToneChatOptions) - Operation response error`, func() {
		version := "testString"
		toneChatPath := "/v3/tone_chat"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(toneChatPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Content-Language"]).ToNot(BeNil())
					Expect(req.Header["Content-Language"][0]).To(Equal(fmt.Sprintf("%v", "en")))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "ar")))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ToneChat with error: Operation response processing error`, func() {
				toneAnalyzerService, serviceErr := toneanalyzerv3.NewToneAnalyzerV3(&toneanalyzerv3.ToneAnalyzerV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(toneAnalyzerService).ToNot(BeNil())

				// Construct an instance of the Utterance model
				utteranceModel := new(toneanalyzerv3.Utterance)
				utteranceModel.Text = core.StringPtr("testString")
				utteranceModel.User = core.StringPtr("testString")

				// Construct an instance of the ToneChatOptions model
				toneChatOptionsModel := new(toneanalyzerv3.ToneChatOptions)
				toneChatOptionsModel.Utterances = []toneanalyzerv3.Utterance{*utteranceModel}
				toneChatOptionsModel.ContentLanguage = core.StringPtr("en")
				toneChatOptionsModel.AcceptLanguage = core.StringPtr("ar")
				toneChatOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := toneAnalyzerService.ToneChat(toneChatOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				toneAnalyzerService.EnableRetries(0, 0)
				result, response, operationErr = toneAnalyzerService.ToneChat(toneChatOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ToneChat(toneChatOptions *ToneChatOptions)`, func() {
		version := "testString"
		toneChatPath := "/v3/tone_chat"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(toneChatPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Content-Language"]).ToNot(BeNil())
					Expect(req.Header["Content-Language"][0]).To(Equal(fmt.Sprintf("%v", "en")))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "ar")))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"utterances_tone": [{"utterance_id": 11, "utterance_text": "UtteranceText", "tones": [{"score": 5, "tone_id": "excited", "tone_name": "ToneName"}], "error": "Error"}], "warning": "Warning"}`)
				}))
			})
			It(`Invoke ToneChat successfully`, func() {
				toneAnalyzerService, serviceErr := toneanalyzerv3.NewToneAnalyzerV3(&toneanalyzerv3.ToneAnalyzerV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(toneAnalyzerService).ToNot(BeNil())
				toneAnalyzerService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := toneAnalyzerService.ToneChat(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Utterance model
				utteranceModel := new(toneanalyzerv3.Utterance)
				utteranceModel.Text = core.StringPtr("testString")
				utteranceModel.User = core.StringPtr("testString")

				// Construct an instance of the ToneChatOptions model
				toneChatOptionsModel := new(toneanalyzerv3.ToneChatOptions)
				toneChatOptionsModel.Utterances = []toneanalyzerv3.Utterance{*utteranceModel}
				toneChatOptionsModel.ContentLanguage = core.StringPtr("en")
				toneChatOptionsModel.AcceptLanguage = core.StringPtr("ar")
				toneChatOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = toneAnalyzerService.ToneChat(toneChatOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = toneAnalyzerService.ToneChatWithContext(ctx, toneChatOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				toneAnalyzerService.DisableRetries()
				result, response, operationErr = toneAnalyzerService.ToneChat(toneChatOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = toneAnalyzerService.ToneChatWithContext(ctx, toneChatOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ToneChat with error: Operation validation and request error`, func() {
				toneAnalyzerService, serviceErr := toneanalyzerv3.NewToneAnalyzerV3(&toneanalyzerv3.ToneAnalyzerV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(toneAnalyzerService).ToNot(BeNil())

				// Construct an instance of the Utterance model
				utteranceModel := new(toneanalyzerv3.Utterance)
				utteranceModel.Text = core.StringPtr("testString")
				utteranceModel.User = core.StringPtr("testString")

				// Construct an instance of the ToneChatOptions model
				toneChatOptionsModel := new(toneanalyzerv3.ToneChatOptions)
				toneChatOptionsModel.Utterances = []toneanalyzerv3.Utterance{*utteranceModel}
				toneChatOptionsModel.ContentLanguage = core.StringPtr("en")
				toneChatOptionsModel.AcceptLanguage = core.StringPtr("ar")
				toneChatOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := toneAnalyzerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := toneAnalyzerService.ToneChat(toneChatOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ToneChatOptions model with no property values
				toneChatOptionsModelNew := new(toneanalyzerv3.ToneChatOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = toneAnalyzerService.ToneChat(toneChatOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			version := "testString"
			toneAnalyzerService, _ := toneanalyzerv3.NewToneAnalyzerV3(&toneanalyzerv3.ToneAnalyzerV3Options{
				URL:           "http://toneanalyzerv3modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       core.StringPtr(version),
			})
			It(`Invoke NewToneChatOptions successfully`, func() {
				// Construct an instance of the Utterance model
				utteranceModel := new(toneanalyzerv3.Utterance)
				Expect(utteranceModel).ToNot(BeNil())
				utteranceModel.Text = core.StringPtr("testString")
				utteranceModel.User = core.StringPtr("testString")
				Expect(utteranceModel.Text).To(Equal(core.StringPtr("testString")))
				Expect(utteranceModel.User).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ToneChatOptions model
				toneChatOptionsUtterances := []toneanalyzerv3.Utterance{}
				toneChatOptionsModel := toneAnalyzerService.NewToneChatOptions(toneChatOptionsUtterances)
				toneChatOptionsModel.SetUtterances([]toneanalyzerv3.Utterance{*utteranceModel})
				toneChatOptionsModel.SetContentLanguage("en")
				toneChatOptionsModel.SetAcceptLanguage("ar")
				toneChatOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(toneChatOptionsModel).ToNot(BeNil())
				Expect(toneChatOptionsModel.Utterances).To(Equal([]toneanalyzerv3.Utterance{*utteranceModel}))
				Expect(toneChatOptionsModel.ContentLanguage).To(Equal(core.StringPtr("en")))
				Expect(toneChatOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("ar")))
				Expect(toneChatOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewToneInput successfully`, func() {
				text := "testString"
				model, err := toneAnalyzerService.NewToneInput(text)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewToneOptions successfully`, func() {
				// Construct an instance of the ToneInput model
				toneInputModel := new(toneanalyzerv3.ToneInput)
				Expect(toneInputModel).ToNot(BeNil())
				toneInputModel.Text = core.StringPtr("testString")
				Expect(toneInputModel.Text).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ToneOptions model
				toneOptionsModel := toneAnalyzerService.NewToneOptions()
				toneOptionsModel.SetToneInput(toneInputModel)
				toneOptionsModel.SetBody("testString")
				toneOptionsModel.SetContentType("application/json")
				toneOptionsModel.SetSentences(true)
				toneOptionsModel.SetTones([]string{"emotion"})
				toneOptionsModel.SetContentLanguage("en")
				toneOptionsModel.SetAcceptLanguage("ar")
				toneOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(toneOptionsModel).ToNot(BeNil())
				Expect(toneOptionsModel.ToneInput).To(Equal(toneInputModel))
				Expect(toneOptionsModel.Body).To(Equal(core.StringPtr("testString")))
				Expect(toneOptionsModel.ContentType).To(Equal(core.StringPtr("application/json")))
				Expect(toneOptionsModel.Sentences).To(Equal(core.BoolPtr(true)))
				Expect(toneOptionsModel.Tones).To(Equal([]string{"emotion"}))
				Expect(toneOptionsModel.ContentLanguage).To(Equal(core.StringPtr("en")))
				Expect(toneOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("ar")))
				Expect(toneOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUtterance successfully`, func() {
				text := "testString"
				model, err := toneAnalyzerService.NewUtterance(text)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate()
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime()
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate() *strfmt.Date {
	d := strfmt.Date(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func CreateMockDateTime() *strfmt.DateTime {
	d := strfmt.DateTime(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
