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

package languagetranslatorv3_test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/languagetranslatorv3"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

var _ = Describe(`LanguageTranslatorV3`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		version := "testString"
		It(`Instantiate service client`, func() {
			languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version: core.StringPtr(version),
			})
			Expect(languageTranslatorService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
			Expect(languageTranslatorService.Service.IsSSLDisabled()).To(BeFalse())
			languageTranslatorService.DisableSSLVerification()
			Expect(languageTranslatorService.Service.IsSSLDisabled()).To(BeTrue())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
				URL: "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(languageTranslatorService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
				URL: "https://languagetranslatorv3/api",
				Version: core.StringPtr(version),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(languageTranslatorService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{})
			Expect(languageTranslatorService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"LANGUAGE_TRANSLATOR_URL": "https://languagetranslatorv3/api",
				"LANGUAGE_TRANSLATOR_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					Version: core.StringPtr(version),
				})
				Expect(languageTranslatorService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := languageTranslatorService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != languageTranslatorService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(languageTranslatorService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(languageTranslatorService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL: "https://testService/api",
					Version: core.StringPtr(version),
				})
				Expect(languageTranslatorService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := languageTranslatorService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != languageTranslatorService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(languageTranslatorService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(languageTranslatorService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					Version: core.StringPtr(version),
				})
				err := languageTranslatorService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := languageTranslatorService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != languageTranslatorService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(languageTranslatorService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(languageTranslatorService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"LANGUAGE_TRANSLATOR_URL": "https://languagetranslatorv3/api",
				"LANGUAGE_TRANSLATOR_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(languageTranslatorService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"LANGUAGE_TRANSLATOR_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
				URL: "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(languageTranslatorService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = languagetranslatorv3.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListLanguages(listLanguagesOptions *ListLanguagesOptions) - Operation response error`, func() {
		version := "testString"
		listLanguagesPath := "/v3/languages"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listLanguagesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListLanguages with error: Operation response processing error`, func() {
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())

				// Construct an instance of the ListLanguagesOptions model
				listLanguagesOptionsModel := new(languagetranslatorv3.ListLanguagesOptions)
				listLanguagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := languageTranslatorService.ListLanguages(listLanguagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				languageTranslatorService.EnableRetries(0, 0)
				result, response, operationErr = languageTranslatorService.ListLanguages(listLanguagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListLanguages(listLanguagesOptions *ListLanguagesOptions)`, func() {
		version := "testString"
		listLanguagesPath := "/v3/languages"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listLanguagesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"languages": [{"language": "Language", "language_name": "LanguageName", "native_language_name": "NativeLanguageName", "country_code": "CountryCode", "words_separated": true, "direction": "Direction", "supported_as_source": false, "supported_as_target": false, "identifiable": true}]}`)
				}))
			})
			It(`Invoke ListLanguages successfully`, func() {
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())
				languageTranslatorService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := languageTranslatorService.ListLanguages(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListLanguagesOptions model
				listLanguagesOptionsModel := new(languagetranslatorv3.ListLanguagesOptions)
				listLanguagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = languageTranslatorService.ListLanguages(listLanguagesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = languageTranslatorService.ListLanguagesWithContext(ctx, listLanguagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				languageTranslatorService.DisableRetries()
				result, response, operationErr = languageTranslatorService.ListLanguages(listLanguagesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = languageTranslatorService.ListLanguagesWithContext(ctx, listLanguagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListLanguages with error: Operation request error`, func() {
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())

				// Construct an instance of the ListLanguagesOptions model
				listLanguagesOptionsModel := new(languagetranslatorv3.ListLanguagesOptions)
				listLanguagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := languageTranslatorService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := languageTranslatorService.ListLanguages(listLanguagesOptionsModel)
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
	Describe(`Service constructor tests`, func() {
		version := "testString"
		It(`Instantiate service client`, func() {
			languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version: core.StringPtr(version),
			})
			Expect(languageTranslatorService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
			Expect(languageTranslatorService.Service.IsSSLDisabled()).To(BeFalse())
			languageTranslatorService.DisableSSLVerification()
			Expect(languageTranslatorService.Service.IsSSLDisabled()).To(BeTrue())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
				URL: "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(languageTranslatorService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
				URL: "https://languagetranslatorv3/api",
				Version: core.StringPtr(version),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(languageTranslatorService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{})
			Expect(languageTranslatorService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"LANGUAGE_TRANSLATOR_URL": "https://languagetranslatorv3/api",
				"LANGUAGE_TRANSLATOR_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					Version: core.StringPtr(version),
				})
				Expect(languageTranslatorService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := languageTranslatorService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != languageTranslatorService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(languageTranslatorService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(languageTranslatorService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL: "https://testService/api",
					Version: core.StringPtr(version),
				})
				Expect(languageTranslatorService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := languageTranslatorService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != languageTranslatorService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(languageTranslatorService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(languageTranslatorService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					Version: core.StringPtr(version),
				})
				err := languageTranslatorService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := languageTranslatorService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != languageTranslatorService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(languageTranslatorService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(languageTranslatorService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"LANGUAGE_TRANSLATOR_URL": "https://languagetranslatorv3/api",
				"LANGUAGE_TRANSLATOR_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(languageTranslatorService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"LANGUAGE_TRANSLATOR_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
				URL: "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(languageTranslatorService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = languagetranslatorv3.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`Translate(translateOptions *TranslateOptions) - Operation response error`, func() {
		version := "testString"
		translatePath := "/v3/translate"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(translatePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke Translate with error: Operation response processing error`, func() {
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())

				// Construct an instance of the TranslateOptions model
				translateOptionsModel := new(languagetranslatorv3.TranslateOptions)
				translateOptionsModel.Text = []string{"testString"}
				translateOptionsModel.ModelID = core.StringPtr("testString")
				translateOptionsModel.Source = core.StringPtr("testString")
				translateOptionsModel.Target = core.StringPtr("testString")
				translateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := languageTranslatorService.Translate(translateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				languageTranslatorService.EnableRetries(0, 0)
				result, response, operationErr = languageTranslatorService.Translate(translateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`Translate(translateOptions *TranslateOptions)`, func() {
		version := "testString"
		translatePath := "/v3/translate"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(translatePath))
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

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"word_count": 9, "character_count": 14, "detected_language": "DetectedLanguage", "detected_language_confidence": 0, "translations": [{"translation": "Translation"}]}`)
				}))
			})
			It(`Invoke Translate successfully`, func() {
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())
				languageTranslatorService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := languageTranslatorService.Translate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TranslateOptions model
				translateOptionsModel := new(languagetranslatorv3.TranslateOptions)
				translateOptionsModel.Text = []string{"testString"}
				translateOptionsModel.ModelID = core.StringPtr("testString")
				translateOptionsModel.Source = core.StringPtr("testString")
				translateOptionsModel.Target = core.StringPtr("testString")
				translateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = languageTranslatorService.Translate(translateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = languageTranslatorService.TranslateWithContext(ctx, translateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				languageTranslatorService.DisableRetries()
				result, response, operationErr = languageTranslatorService.Translate(translateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = languageTranslatorService.TranslateWithContext(ctx, translateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke Translate with error: Operation validation and request error`, func() {
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())

				// Construct an instance of the TranslateOptions model
				translateOptionsModel := new(languagetranslatorv3.TranslateOptions)
				translateOptionsModel.Text = []string{"testString"}
				translateOptionsModel.ModelID = core.StringPtr("testString")
				translateOptionsModel.Source = core.StringPtr("testString")
				translateOptionsModel.Target = core.StringPtr("testString")
				translateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := languageTranslatorService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := languageTranslatorService.Translate(translateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the TranslateOptions model with no property values
				translateOptionsModelNew := new(languagetranslatorv3.TranslateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = languageTranslatorService.Translate(translateOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		version := "testString"
		It(`Instantiate service client`, func() {
			languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version: core.StringPtr(version),
			})
			Expect(languageTranslatorService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
			Expect(languageTranslatorService.Service.IsSSLDisabled()).To(BeFalse())
			languageTranslatorService.DisableSSLVerification()
			Expect(languageTranslatorService.Service.IsSSLDisabled()).To(BeTrue())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
				URL: "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(languageTranslatorService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
				URL: "https://languagetranslatorv3/api",
				Version: core.StringPtr(version),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(languageTranslatorService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{})
			Expect(languageTranslatorService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"LANGUAGE_TRANSLATOR_URL": "https://languagetranslatorv3/api",
				"LANGUAGE_TRANSLATOR_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					Version: core.StringPtr(version),
				})
				Expect(languageTranslatorService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := languageTranslatorService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != languageTranslatorService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(languageTranslatorService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(languageTranslatorService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL: "https://testService/api",
					Version: core.StringPtr(version),
				})
				Expect(languageTranslatorService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := languageTranslatorService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != languageTranslatorService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(languageTranslatorService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(languageTranslatorService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					Version: core.StringPtr(version),
				})
				err := languageTranslatorService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := languageTranslatorService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != languageTranslatorService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(languageTranslatorService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(languageTranslatorService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"LANGUAGE_TRANSLATOR_URL": "https://languagetranslatorv3/api",
				"LANGUAGE_TRANSLATOR_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(languageTranslatorService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"LANGUAGE_TRANSLATOR_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
				URL: "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(languageTranslatorService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = languagetranslatorv3.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListIdentifiableLanguages(listIdentifiableLanguagesOptions *ListIdentifiableLanguagesOptions) - Operation response error`, func() {
		version := "testString"
		listIdentifiableLanguagesPath := "/v3/identifiable_languages"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listIdentifiableLanguagesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListIdentifiableLanguages with error: Operation response processing error`, func() {
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())

				// Construct an instance of the ListIdentifiableLanguagesOptions model
				listIdentifiableLanguagesOptionsModel := new(languagetranslatorv3.ListIdentifiableLanguagesOptions)
				listIdentifiableLanguagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := languageTranslatorService.ListIdentifiableLanguages(listIdentifiableLanguagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				languageTranslatorService.EnableRetries(0, 0)
				result, response, operationErr = languageTranslatorService.ListIdentifiableLanguages(listIdentifiableLanguagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListIdentifiableLanguages(listIdentifiableLanguagesOptions *ListIdentifiableLanguagesOptions)`, func() {
		version := "testString"
		listIdentifiableLanguagesPath := "/v3/identifiable_languages"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listIdentifiableLanguagesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"languages": [{"language": "Language", "name": "Name"}]}`)
				}))
			})
			It(`Invoke ListIdentifiableLanguages successfully`, func() {
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())
				languageTranslatorService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := languageTranslatorService.ListIdentifiableLanguages(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListIdentifiableLanguagesOptions model
				listIdentifiableLanguagesOptionsModel := new(languagetranslatorv3.ListIdentifiableLanguagesOptions)
				listIdentifiableLanguagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = languageTranslatorService.ListIdentifiableLanguages(listIdentifiableLanguagesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = languageTranslatorService.ListIdentifiableLanguagesWithContext(ctx, listIdentifiableLanguagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				languageTranslatorService.DisableRetries()
				result, response, operationErr = languageTranslatorService.ListIdentifiableLanguages(listIdentifiableLanguagesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = languageTranslatorService.ListIdentifiableLanguagesWithContext(ctx, listIdentifiableLanguagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListIdentifiableLanguages with error: Operation request error`, func() {
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())

				// Construct an instance of the ListIdentifiableLanguagesOptions model
				listIdentifiableLanguagesOptionsModel := new(languagetranslatorv3.ListIdentifiableLanguagesOptions)
				listIdentifiableLanguagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := languageTranslatorService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := languageTranslatorService.ListIdentifiableLanguages(listIdentifiableLanguagesOptionsModel)
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
	Describe(`Identify(identifyOptions *IdentifyOptions) - Operation response error`, func() {
		version := "testString"
		identifyPath := "/v3/identify"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(identifyPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke Identify with error: Operation response processing error`, func() {
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())

				// Construct an instance of the IdentifyOptions model
				identifyOptionsModel := new(languagetranslatorv3.IdentifyOptions)
				identifyOptionsModel.Text = core.StringPtr("testString")
				identifyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := languageTranslatorService.Identify(identifyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				languageTranslatorService.EnableRetries(0, 0)
				result, response, operationErr = languageTranslatorService.Identify(identifyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`Identify(identifyOptions *IdentifyOptions)`, func() {
		version := "testString"
		identifyPath := "/v3/identify"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(identifyPath))
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

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"languages": [{"language": "Language", "confidence": 0}]}`)
				}))
			})
			It(`Invoke Identify successfully`, func() {
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())
				languageTranslatorService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := languageTranslatorService.Identify(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the IdentifyOptions model
				identifyOptionsModel := new(languagetranslatorv3.IdentifyOptions)
				identifyOptionsModel.Text = core.StringPtr("testString")
				identifyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = languageTranslatorService.Identify(identifyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = languageTranslatorService.IdentifyWithContext(ctx, identifyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				languageTranslatorService.DisableRetries()
				result, response, operationErr = languageTranslatorService.Identify(identifyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = languageTranslatorService.IdentifyWithContext(ctx, identifyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke Identify with error: Operation validation and request error`, func() {
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())

				// Construct an instance of the IdentifyOptions model
				identifyOptionsModel := new(languagetranslatorv3.IdentifyOptions)
				identifyOptionsModel.Text = core.StringPtr("testString")
				identifyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := languageTranslatorService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := languageTranslatorService.Identify(identifyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the IdentifyOptions model with no property values
				identifyOptionsModelNew := new(languagetranslatorv3.IdentifyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = languageTranslatorService.Identify(identifyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		version := "testString"
		It(`Instantiate service client`, func() {
			languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version: core.StringPtr(version),
			})
			Expect(languageTranslatorService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
			Expect(languageTranslatorService.Service.IsSSLDisabled()).To(BeFalse())
			languageTranslatorService.DisableSSLVerification()
			Expect(languageTranslatorService.Service.IsSSLDisabled()).To(BeTrue())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
				URL: "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(languageTranslatorService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
				URL: "https://languagetranslatorv3/api",
				Version: core.StringPtr(version),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(languageTranslatorService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{})
			Expect(languageTranslatorService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"LANGUAGE_TRANSLATOR_URL": "https://languagetranslatorv3/api",
				"LANGUAGE_TRANSLATOR_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					Version: core.StringPtr(version),
				})
				Expect(languageTranslatorService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := languageTranslatorService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != languageTranslatorService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(languageTranslatorService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(languageTranslatorService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL: "https://testService/api",
					Version: core.StringPtr(version),
				})
				Expect(languageTranslatorService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := languageTranslatorService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != languageTranslatorService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(languageTranslatorService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(languageTranslatorService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					Version: core.StringPtr(version),
				})
				err := languageTranslatorService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := languageTranslatorService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != languageTranslatorService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(languageTranslatorService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(languageTranslatorService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"LANGUAGE_TRANSLATOR_URL": "https://languagetranslatorv3/api",
				"LANGUAGE_TRANSLATOR_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(languageTranslatorService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"LANGUAGE_TRANSLATOR_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
				URL: "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(languageTranslatorService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = languagetranslatorv3.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListModels(listModelsOptions *ListModelsOptions) - Operation response error`, func() {
		version := "testString"
		listModelsPath := "/v3/models"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listModelsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["source"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["target"]).To(Equal([]string{"testString"}))


					// TODO: Add check for default query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListModels with error: Operation response processing error`, func() {
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())

				// Construct an instance of the ListModelsOptions model
				listModelsOptionsModel := new(languagetranslatorv3.ListModelsOptions)
				listModelsOptionsModel.Source = core.StringPtr("testString")
				listModelsOptionsModel.Target = core.StringPtr("testString")
				listModelsOptionsModel.Default = core.BoolPtr(true)
				listModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := languageTranslatorService.ListModels(listModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				languageTranslatorService.EnableRetries(0, 0)
				result, response, operationErr = languageTranslatorService.ListModels(listModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListModels(listModelsOptions *ListModelsOptions)`, func() {
		version := "testString"
		listModelsPath := "/v3/models"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listModelsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["source"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["target"]).To(Equal([]string{"testString"}))


					// TODO: Add check for default query parameter

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"models": [{"model_id": "ModelID", "name": "Name", "source": "Source", "target": "Target", "base_model_id": "BaseModelID", "domain": "Domain", "customizable": true, "default_model": true, "owner": "Owner", "status": "uploading"}]}`)
				}))
			})
			It(`Invoke ListModels successfully`, func() {
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())
				languageTranslatorService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := languageTranslatorService.ListModels(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListModelsOptions model
				listModelsOptionsModel := new(languagetranslatorv3.ListModelsOptions)
				listModelsOptionsModel.Source = core.StringPtr("testString")
				listModelsOptionsModel.Target = core.StringPtr("testString")
				listModelsOptionsModel.Default = core.BoolPtr(true)
				listModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = languageTranslatorService.ListModels(listModelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = languageTranslatorService.ListModelsWithContext(ctx, listModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				languageTranslatorService.DisableRetries()
				result, response, operationErr = languageTranslatorService.ListModels(listModelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = languageTranslatorService.ListModelsWithContext(ctx, listModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListModels with error: Operation request error`, func() {
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())

				// Construct an instance of the ListModelsOptions model
				listModelsOptionsModel := new(languagetranslatorv3.ListModelsOptions)
				listModelsOptionsModel.Source = core.StringPtr("testString")
				listModelsOptionsModel.Target = core.StringPtr("testString")
				listModelsOptionsModel.Default = core.BoolPtr(true)
				listModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := languageTranslatorService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := languageTranslatorService.ListModels(listModelsOptionsModel)
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
	Describe(`CreateModel(createModelOptions *CreateModelOptions) - Operation response error`, func() {
		version := "testString"
		createModelPath := "/v3/models"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createModelPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["base_model_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateModel with error: Operation response processing error`, func() {
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())

				// Construct an instance of the CreateModelOptions model
				createModelOptionsModel := new(languagetranslatorv3.CreateModelOptions)
				createModelOptionsModel.BaseModelID = core.StringPtr("testString")
				createModelOptionsModel.ForcedGlossary = CreateMockReader("This is a mock file.")
				createModelOptionsModel.ParallelCorpus = CreateMockReader("This is a mock file.")
				createModelOptionsModel.Name = core.StringPtr("testString")
				createModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := languageTranslatorService.CreateModel(createModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				languageTranslatorService.EnableRetries(0, 0)
				result, response, operationErr = languageTranslatorService.CreateModel(createModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateModel(createModelOptions *CreateModelOptions)`, func() {
		version := "testString"
		createModelPath := "/v3/models"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createModelPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["base_model_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"model_id": "ModelID", "name": "Name", "source": "Source", "target": "Target", "base_model_id": "BaseModelID", "domain": "Domain", "customizable": true, "default_model": true, "owner": "Owner", "status": "uploading"}`)
				}))
			})
			It(`Invoke CreateModel successfully`, func() {
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())
				languageTranslatorService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := languageTranslatorService.CreateModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateModelOptions model
				createModelOptionsModel := new(languagetranslatorv3.CreateModelOptions)
				createModelOptionsModel.BaseModelID = core.StringPtr("testString")
				createModelOptionsModel.ForcedGlossary = CreateMockReader("This is a mock file.")
				createModelOptionsModel.ParallelCorpus = CreateMockReader("This is a mock file.")
				createModelOptionsModel.Name = core.StringPtr("testString")
				createModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = languageTranslatorService.CreateModel(createModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = languageTranslatorService.CreateModelWithContext(ctx, createModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				languageTranslatorService.DisableRetries()
				result, response, operationErr = languageTranslatorService.CreateModel(createModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = languageTranslatorService.CreateModelWithContext(ctx, createModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CreateModel with error: Param validation error`, func() {
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:  testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())

				// Construct an instance of the CreateModelOptions model
				createModelOptionsModel := new(languagetranslatorv3.CreateModelOptions)
				// Invoke operation with invalid options model (negative test)
				result, response, operationErr := languageTranslatorService.CreateModel(createModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			It(`Invoke CreateModel with error: Operation validation and request error`, func() {
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())

				// Construct an instance of the CreateModelOptions model
				createModelOptionsModel := new(languagetranslatorv3.CreateModelOptions)
				createModelOptionsModel.BaseModelID = core.StringPtr("testString")
				createModelOptionsModel.ForcedGlossary = CreateMockReader("This is a mock file.")
				createModelOptionsModel.ParallelCorpus = CreateMockReader("This is a mock file.")
				createModelOptionsModel.Name = core.StringPtr("testString")
				createModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := languageTranslatorService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := languageTranslatorService.CreateModel(createModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateModelOptions model with no property values
				createModelOptionsModelNew := new(languagetranslatorv3.CreateModelOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = languageTranslatorService.CreateModel(createModelOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteModel(deleteModelOptions *DeleteModelOptions) - Operation response error`, func() {
		version := "testString"
		deleteModelPath := "/v3/models/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteModelPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteModel with error: Operation response processing error`, func() {
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())

				// Construct an instance of the DeleteModelOptions model
				deleteModelOptionsModel := new(languagetranslatorv3.DeleteModelOptions)
				deleteModelOptionsModel.ModelID = core.StringPtr("testString")
				deleteModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := languageTranslatorService.DeleteModel(deleteModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				languageTranslatorService.EnableRetries(0, 0)
				result, response, operationErr = languageTranslatorService.DeleteModel(deleteModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteModel(deleteModelOptions *DeleteModelOptions)`, func() {
		version := "testString"
		deleteModelPath := "/v3/models/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteModelPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status": "Status"}`)
				}))
			})
			It(`Invoke DeleteModel successfully`, func() {
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())
				languageTranslatorService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := languageTranslatorService.DeleteModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteModelOptions model
				deleteModelOptionsModel := new(languagetranslatorv3.DeleteModelOptions)
				deleteModelOptionsModel.ModelID = core.StringPtr("testString")
				deleteModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = languageTranslatorService.DeleteModel(deleteModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = languageTranslatorService.DeleteModelWithContext(ctx, deleteModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				languageTranslatorService.DisableRetries()
				result, response, operationErr = languageTranslatorService.DeleteModel(deleteModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = languageTranslatorService.DeleteModelWithContext(ctx, deleteModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke DeleteModel with error: Operation validation and request error`, func() {
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())

				// Construct an instance of the DeleteModelOptions model
				deleteModelOptionsModel := new(languagetranslatorv3.DeleteModelOptions)
				deleteModelOptionsModel.ModelID = core.StringPtr("testString")
				deleteModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := languageTranslatorService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := languageTranslatorService.DeleteModel(deleteModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteModelOptions model with no property values
				deleteModelOptionsModelNew := new(languagetranslatorv3.DeleteModelOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = languageTranslatorService.DeleteModel(deleteModelOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetModel(getModelOptions *GetModelOptions) - Operation response error`, func() {
		version := "testString"
		getModelPath := "/v3/models/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getModelPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetModel with error: Operation response processing error`, func() {
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())

				// Construct an instance of the GetModelOptions model
				getModelOptionsModel := new(languagetranslatorv3.GetModelOptions)
				getModelOptionsModel.ModelID = core.StringPtr("testString")
				getModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := languageTranslatorService.GetModel(getModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				languageTranslatorService.EnableRetries(0, 0)
				result, response, operationErr = languageTranslatorService.GetModel(getModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetModel(getModelOptions *GetModelOptions)`, func() {
		version := "testString"
		getModelPath := "/v3/models/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getModelPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"model_id": "ModelID", "name": "Name", "source": "Source", "target": "Target", "base_model_id": "BaseModelID", "domain": "Domain", "customizable": true, "default_model": true, "owner": "Owner", "status": "uploading"}`)
				}))
			})
			It(`Invoke GetModel successfully`, func() {
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())
				languageTranslatorService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := languageTranslatorService.GetModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetModelOptions model
				getModelOptionsModel := new(languagetranslatorv3.GetModelOptions)
				getModelOptionsModel.ModelID = core.StringPtr("testString")
				getModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = languageTranslatorService.GetModel(getModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = languageTranslatorService.GetModelWithContext(ctx, getModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				languageTranslatorService.DisableRetries()
				result, response, operationErr = languageTranslatorService.GetModel(getModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = languageTranslatorService.GetModelWithContext(ctx, getModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetModel with error: Operation validation and request error`, func() {
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())

				// Construct an instance of the GetModelOptions model
				getModelOptionsModel := new(languagetranslatorv3.GetModelOptions)
				getModelOptionsModel.ModelID = core.StringPtr("testString")
				getModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := languageTranslatorService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := languageTranslatorService.GetModel(getModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetModelOptions model with no property values
				getModelOptionsModelNew := new(languagetranslatorv3.GetModelOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = languageTranslatorService.GetModel(getModelOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		version := "testString"
		It(`Instantiate service client`, func() {
			languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version: core.StringPtr(version),
			})
			Expect(languageTranslatorService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
			Expect(languageTranslatorService.Service.IsSSLDisabled()).To(BeFalse())
			languageTranslatorService.DisableSSLVerification()
			Expect(languageTranslatorService.Service.IsSSLDisabled()).To(BeTrue())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
				URL: "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(languageTranslatorService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
				URL: "https://languagetranslatorv3/api",
				Version: core.StringPtr(version),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(languageTranslatorService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{})
			Expect(languageTranslatorService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"LANGUAGE_TRANSLATOR_URL": "https://languagetranslatorv3/api",
				"LANGUAGE_TRANSLATOR_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					Version: core.StringPtr(version),
				})
				Expect(languageTranslatorService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := languageTranslatorService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != languageTranslatorService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(languageTranslatorService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(languageTranslatorService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL: "https://testService/api",
					Version: core.StringPtr(version),
				})
				Expect(languageTranslatorService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := languageTranslatorService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != languageTranslatorService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(languageTranslatorService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(languageTranslatorService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					Version: core.StringPtr(version),
				})
				err := languageTranslatorService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := languageTranslatorService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != languageTranslatorService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(languageTranslatorService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(languageTranslatorService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"LANGUAGE_TRANSLATOR_URL": "https://languagetranslatorv3/api",
				"LANGUAGE_TRANSLATOR_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(languageTranslatorService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"LANGUAGE_TRANSLATOR_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
				URL: "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(languageTranslatorService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = languagetranslatorv3.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListDocuments(listDocumentsOptions *ListDocumentsOptions) - Operation response error`, func() {
		version := "testString"
		listDocumentsPath := "/v3/documents"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDocumentsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListDocuments with error: Operation response processing error`, func() {
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())

				// Construct an instance of the ListDocumentsOptions model
				listDocumentsOptionsModel := new(languagetranslatorv3.ListDocumentsOptions)
				listDocumentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := languageTranslatorService.ListDocuments(listDocumentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				languageTranslatorService.EnableRetries(0, 0)
				result, response, operationErr = languageTranslatorService.ListDocuments(listDocumentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListDocuments(listDocumentsOptions *ListDocumentsOptions)`, func() {
		version := "testString"
		listDocumentsPath := "/v3/documents"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDocumentsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"documents": [{"document_id": "DocumentID", "filename": "Filename", "status": "processing", "model_id": "ModelID", "base_model_id": "BaseModelID", "source": "Source", "detected_language_confidence": 0, "target": "Target", "created": "2019-01-01T12:00:00", "completed": "2019-01-01T12:00:00", "word_count": 9, "character_count": 14}]}`)
				}))
			})
			It(`Invoke ListDocuments successfully`, func() {
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())
				languageTranslatorService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := languageTranslatorService.ListDocuments(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListDocumentsOptions model
				listDocumentsOptionsModel := new(languagetranslatorv3.ListDocumentsOptions)
				listDocumentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = languageTranslatorService.ListDocuments(listDocumentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = languageTranslatorService.ListDocumentsWithContext(ctx, listDocumentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				languageTranslatorService.DisableRetries()
				result, response, operationErr = languageTranslatorService.ListDocuments(listDocumentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = languageTranslatorService.ListDocumentsWithContext(ctx, listDocumentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListDocuments with error: Operation request error`, func() {
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())

				// Construct an instance of the ListDocumentsOptions model
				listDocumentsOptionsModel := new(languagetranslatorv3.ListDocumentsOptions)
				listDocumentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := languageTranslatorService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := languageTranslatorService.ListDocuments(listDocumentsOptionsModel)
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
	Describe(`TranslateDocument(translateDocumentOptions *TranslateDocumentOptions) - Operation response error`, func() {
		version := "testString"
		translateDocumentPath := "/v3/documents"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(translateDocumentPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke TranslateDocument with error: Operation response processing error`, func() {
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())

				// Construct an instance of the TranslateDocumentOptions model
				translateDocumentOptionsModel := new(languagetranslatorv3.TranslateDocumentOptions)
				translateDocumentOptionsModel.File = CreateMockReader("This is a mock file.")
				translateDocumentOptionsModel.Filename = core.StringPtr("testString")
				translateDocumentOptionsModel.FileContentType = core.StringPtr("application/powerpoint")
				translateDocumentOptionsModel.ModelID = core.StringPtr("testString")
				translateDocumentOptionsModel.Source = core.StringPtr("testString")
				translateDocumentOptionsModel.Target = core.StringPtr("testString")
				translateDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				translateDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := languageTranslatorService.TranslateDocument(translateDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				languageTranslatorService.EnableRetries(0, 0)
				result, response, operationErr = languageTranslatorService.TranslateDocument(translateDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`TranslateDocument(translateDocumentOptions *TranslateDocumentOptions)`, func() {
		version := "testString"
		translateDocumentPath := "/v3/documents"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(translateDocumentPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"document_id": "DocumentID", "filename": "Filename", "status": "processing", "model_id": "ModelID", "base_model_id": "BaseModelID", "source": "Source", "detected_language_confidence": 0, "target": "Target", "created": "2019-01-01T12:00:00", "completed": "2019-01-01T12:00:00", "word_count": 9, "character_count": 14}`)
				}))
			})
			It(`Invoke TranslateDocument successfully`, func() {
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())
				languageTranslatorService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := languageTranslatorService.TranslateDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TranslateDocumentOptions model
				translateDocumentOptionsModel := new(languagetranslatorv3.TranslateDocumentOptions)
				translateDocumentOptionsModel.File = CreateMockReader("This is a mock file.")
				translateDocumentOptionsModel.Filename = core.StringPtr("testString")
				translateDocumentOptionsModel.FileContentType = core.StringPtr("application/powerpoint")
				translateDocumentOptionsModel.ModelID = core.StringPtr("testString")
				translateDocumentOptionsModel.Source = core.StringPtr("testString")
				translateDocumentOptionsModel.Target = core.StringPtr("testString")
				translateDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				translateDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = languageTranslatorService.TranslateDocument(translateDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = languageTranslatorService.TranslateDocumentWithContext(ctx, translateDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				languageTranslatorService.DisableRetries()
				result, response, operationErr = languageTranslatorService.TranslateDocument(translateDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = languageTranslatorService.TranslateDocumentWithContext(ctx, translateDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke TranslateDocument with error: Operation validation and request error`, func() {
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())

				// Construct an instance of the TranslateDocumentOptions model
				translateDocumentOptionsModel := new(languagetranslatorv3.TranslateDocumentOptions)
				translateDocumentOptionsModel.File = CreateMockReader("This is a mock file.")
				translateDocumentOptionsModel.Filename = core.StringPtr("testString")
				translateDocumentOptionsModel.FileContentType = core.StringPtr("application/powerpoint")
				translateDocumentOptionsModel.ModelID = core.StringPtr("testString")
				translateDocumentOptionsModel.Source = core.StringPtr("testString")
				translateDocumentOptionsModel.Target = core.StringPtr("testString")
				translateDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				translateDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := languageTranslatorService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := languageTranslatorService.TranslateDocument(translateDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the TranslateDocumentOptions model with no property values
				translateDocumentOptionsModelNew := new(languagetranslatorv3.TranslateDocumentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = languageTranslatorService.TranslateDocument(translateDocumentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDocumentStatus(getDocumentStatusOptions *GetDocumentStatusOptions) - Operation response error`, func() {
		version := "testString"
		getDocumentStatusPath := "/v3/documents/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDocumentStatusPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDocumentStatus with error: Operation response processing error`, func() {
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())

				// Construct an instance of the GetDocumentStatusOptions model
				getDocumentStatusOptionsModel := new(languagetranslatorv3.GetDocumentStatusOptions)
				getDocumentStatusOptionsModel.DocumentID = core.StringPtr("testString")
				getDocumentStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := languageTranslatorService.GetDocumentStatus(getDocumentStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				languageTranslatorService.EnableRetries(0, 0)
				result, response, operationErr = languageTranslatorService.GetDocumentStatus(getDocumentStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetDocumentStatus(getDocumentStatusOptions *GetDocumentStatusOptions)`, func() {
		version := "testString"
		getDocumentStatusPath := "/v3/documents/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDocumentStatusPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"document_id": "DocumentID", "filename": "Filename", "status": "processing", "model_id": "ModelID", "base_model_id": "BaseModelID", "source": "Source", "detected_language_confidence": 0, "target": "Target", "created": "2019-01-01T12:00:00", "completed": "2019-01-01T12:00:00", "word_count": 9, "character_count": 14}`)
				}))
			})
			It(`Invoke GetDocumentStatus successfully`, func() {
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())
				languageTranslatorService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := languageTranslatorService.GetDocumentStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDocumentStatusOptions model
				getDocumentStatusOptionsModel := new(languagetranslatorv3.GetDocumentStatusOptions)
				getDocumentStatusOptionsModel.DocumentID = core.StringPtr("testString")
				getDocumentStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = languageTranslatorService.GetDocumentStatus(getDocumentStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = languageTranslatorService.GetDocumentStatusWithContext(ctx, getDocumentStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				languageTranslatorService.DisableRetries()
				result, response, operationErr = languageTranslatorService.GetDocumentStatus(getDocumentStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = languageTranslatorService.GetDocumentStatusWithContext(ctx, getDocumentStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetDocumentStatus with error: Operation validation and request error`, func() {
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())

				// Construct an instance of the GetDocumentStatusOptions model
				getDocumentStatusOptionsModel := new(languagetranslatorv3.GetDocumentStatusOptions)
				getDocumentStatusOptionsModel.DocumentID = core.StringPtr("testString")
				getDocumentStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := languageTranslatorService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := languageTranslatorService.GetDocumentStatus(getDocumentStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDocumentStatusOptions model with no property values
				getDocumentStatusOptionsModelNew := new(languagetranslatorv3.GetDocumentStatusOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = languageTranslatorService.GetDocumentStatus(getDocumentStatusOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteDocument(deleteDocumentOptions *DeleteDocumentOptions)`, func() {
		version := "testString"
		deleteDocumentPath := "/v3/documents/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDocumentPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteDocument successfully`, func() {
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())
				languageTranslatorService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := languageTranslatorService.DeleteDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteDocumentOptions model
				deleteDocumentOptionsModel := new(languagetranslatorv3.DeleteDocumentOptions)
				deleteDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				deleteDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = languageTranslatorService.DeleteDocument(deleteDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				languageTranslatorService.DisableRetries()
				response, operationErr = languageTranslatorService.DeleteDocument(deleteDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteDocument with error: Operation validation and request error`, func() {
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())

				// Construct an instance of the DeleteDocumentOptions model
				deleteDocumentOptionsModel := new(languagetranslatorv3.DeleteDocumentOptions)
				deleteDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				deleteDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := languageTranslatorService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := languageTranslatorService.DeleteDocument(deleteDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteDocumentOptions model with no property values
				deleteDocumentOptionsModelNew := new(languagetranslatorv3.DeleteDocumentOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = languageTranslatorService.DeleteDocument(deleteDocumentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetTranslatedDocument(getTranslatedDocumentOptions *GetTranslatedDocumentOptions)`, func() {
		version := "testString"
		getTranslatedDocumentPath := "/v3/documents/testString/translated_document"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTranslatedDocumentPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "application/powerpoint")))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/powerpoint")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke GetTranslatedDocument successfully`, func() {
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())
				languageTranslatorService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := languageTranslatorService.GetTranslatedDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTranslatedDocumentOptions model
				getTranslatedDocumentOptionsModel := new(languagetranslatorv3.GetTranslatedDocumentOptions)
				getTranslatedDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				getTranslatedDocumentOptionsModel.Accept = core.StringPtr("application/powerpoint")
				getTranslatedDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = languageTranslatorService.GetTranslatedDocument(getTranslatedDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = languageTranslatorService.GetTranslatedDocumentWithContext(ctx, getTranslatedDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				languageTranslatorService.DisableRetries()
				result, response, operationErr = languageTranslatorService.GetTranslatedDocument(getTranslatedDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = languageTranslatorService.GetTranslatedDocumentWithContext(ctx, getTranslatedDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetTranslatedDocument with error: Operation validation and request error`, func() {
				languageTranslatorService, serviceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(languageTranslatorService).ToNot(BeNil())

				// Construct an instance of the GetTranslatedDocumentOptions model
				getTranslatedDocumentOptionsModel := new(languagetranslatorv3.GetTranslatedDocumentOptions)
				getTranslatedDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				getTranslatedDocumentOptionsModel.Accept = core.StringPtr("application/powerpoint")
				getTranslatedDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := languageTranslatorService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := languageTranslatorService.GetTranslatedDocument(getTranslatedDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTranslatedDocumentOptions model with no property values
				getTranslatedDocumentOptionsModelNew := new(languagetranslatorv3.GetTranslatedDocumentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = languageTranslatorService.GetTranslatedDocument(getTranslatedDocumentOptionsModelNew)
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
			languageTranslatorService, _ := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
				URL:           "http://languagetranslatorv3modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Version: core.StringPtr(version),
			})
			It(`Invoke NewCreateModelOptions successfully`, func() {
				// Construct an instance of the CreateModelOptions model
				baseModelID := "testString"
				createModelOptionsModel := languageTranslatorService.NewCreateModelOptions(baseModelID)
				createModelOptionsModel.SetBaseModelID("testString")
				createModelOptionsModel.SetForcedGlossary(CreateMockReader("This is a mock file."))
				createModelOptionsModel.SetParallelCorpus(CreateMockReader("This is a mock file."))
				createModelOptionsModel.SetName("testString")
				createModelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createModelOptionsModel).ToNot(BeNil())
				Expect(createModelOptionsModel.BaseModelID).To(Equal(core.StringPtr("testString")))
				Expect(createModelOptionsModel.ForcedGlossary).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(createModelOptionsModel.ParallelCorpus).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(createModelOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createModelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteDocumentOptions successfully`, func() {
				// Construct an instance of the DeleteDocumentOptions model
				documentID := "testString"
				deleteDocumentOptionsModel := languageTranslatorService.NewDeleteDocumentOptions(documentID)
				deleteDocumentOptionsModel.SetDocumentID("testString")
				deleteDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteDocumentOptionsModel).ToNot(BeNil())
				Expect(deleteDocumentOptionsModel.DocumentID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteModelOptions successfully`, func() {
				// Construct an instance of the DeleteModelOptions model
				modelID := "testString"
				deleteModelOptionsModel := languageTranslatorService.NewDeleteModelOptions(modelID)
				deleteModelOptionsModel.SetModelID("testString")
				deleteModelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteModelOptionsModel).ToNot(BeNil())
				Expect(deleteModelOptionsModel.ModelID).To(Equal(core.StringPtr("testString")))
				Expect(deleteModelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDocumentStatusOptions successfully`, func() {
				// Construct an instance of the GetDocumentStatusOptions model
				documentID := "testString"
				getDocumentStatusOptionsModel := languageTranslatorService.NewGetDocumentStatusOptions(documentID)
				getDocumentStatusOptionsModel.SetDocumentID("testString")
				getDocumentStatusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDocumentStatusOptionsModel).ToNot(BeNil())
				Expect(getDocumentStatusOptionsModel.DocumentID).To(Equal(core.StringPtr("testString")))
				Expect(getDocumentStatusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetModelOptions successfully`, func() {
				// Construct an instance of the GetModelOptions model
				modelID := "testString"
				getModelOptionsModel := languageTranslatorService.NewGetModelOptions(modelID)
				getModelOptionsModel.SetModelID("testString")
				getModelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getModelOptionsModel).ToNot(BeNil())
				Expect(getModelOptionsModel.ModelID).To(Equal(core.StringPtr("testString")))
				Expect(getModelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTranslatedDocumentOptions successfully`, func() {
				// Construct an instance of the GetTranslatedDocumentOptions model
				documentID := "testString"
				getTranslatedDocumentOptionsModel := languageTranslatorService.NewGetTranslatedDocumentOptions(documentID)
				getTranslatedDocumentOptionsModel.SetDocumentID("testString")
				getTranslatedDocumentOptionsModel.SetAccept("application/powerpoint")
				getTranslatedDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTranslatedDocumentOptionsModel).ToNot(BeNil())
				Expect(getTranslatedDocumentOptionsModel.DocumentID).To(Equal(core.StringPtr("testString")))
				Expect(getTranslatedDocumentOptionsModel.Accept).To(Equal(core.StringPtr("application/powerpoint")))
				Expect(getTranslatedDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewIdentifyOptions successfully`, func() {
				// Construct an instance of the IdentifyOptions model
				text := "testString"
				identifyOptionsModel := languageTranslatorService.NewIdentifyOptions(text)
				identifyOptionsModel.SetText("testString")
				identifyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(identifyOptionsModel).ToNot(BeNil())
				Expect(identifyOptionsModel.Text).To(Equal(core.StringPtr("testString")))
				Expect(identifyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListDocumentsOptions successfully`, func() {
				// Construct an instance of the ListDocumentsOptions model
				listDocumentsOptionsModel := languageTranslatorService.NewListDocumentsOptions()
				listDocumentsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listDocumentsOptionsModel).ToNot(BeNil())
				Expect(listDocumentsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListIdentifiableLanguagesOptions successfully`, func() {
				// Construct an instance of the ListIdentifiableLanguagesOptions model
				listIdentifiableLanguagesOptionsModel := languageTranslatorService.NewListIdentifiableLanguagesOptions()
				listIdentifiableLanguagesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listIdentifiableLanguagesOptionsModel).ToNot(BeNil())
				Expect(listIdentifiableLanguagesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListLanguagesOptions successfully`, func() {
				// Construct an instance of the ListLanguagesOptions model
				listLanguagesOptionsModel := languageTranslatorService.NewListLanguagesOptions()
				listLanguagesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listLanguagesOptionsModel).ToNot(BeNil())
				Expect(listLanguagesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListModelsOptions successfully`, func() {
				// Construct an instance of the ListModelsOptions model
				listModelsOptionsModel := languageTranslatorService.NewListModelsOptions()
				listModelsOptionsModel.SetSource("testString")
				listModelsOptionsModel.SetTarget("testString")
				listModelsOptionsModel.SetDefault(true)
				listModelsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listModelsOptionsModel).ToNot(BeNil())
				Expect(listModelsOptionsModel.Source).To(Equal(core.StringPtr("testString")))
				Expect(listModelsOptionsModel.Target).To(Equal(core.StringPtr("testString")))
				Expect(listModelsOptionsModel.Default).To(Equal(core.BoolPtr(true)))
				Expect(listModelsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewTranslateDocumentOptions successfully`, func() {
				// Construct an instance of the TranslateDocumentOptions model
				file := CreateMockReader("This is a mock file.")
				filename := "testString"
				translateDocumentOptionsModel := languageTranslatorService.NewTranslateDocumentOptions(file, filename)
				translateDocumentOptionsModel.SetFile(CreateMockReader("This is a mock file."))
				translateDocumentOptionsModel.SetFilename("testString")
				translateDocumentOptionsModel.SetFileContentType("application/powerpoint")
				translateDocumentOptionsModel.SetModelID("testString")
				translateDocumentOptionsModel.SetSource("testString")
				translateDocumentOptionsModel.SetTarget("testString")
				translateDocumentOptionsModel.SetDocumentID("testString")
				translateDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(translateDocumentOptionsModel).ToNot(BeNil())
				Expect(translateDocumentOptionsModel.File).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(translateDocumentOptionsModel.Filename).To(Equal(core.StringPtr("testString")))
				Expect(translateDocumentOptionsModel.FileContentType).To(Equal(core.StringPtr("application/powerpoint")))
				Expect(translateDocumentOptionsModel.ModelID).To(Equal(core.StringPtr("testString")))
				Expect(translateDocumentOptionsModel.Source).To(Equal(core.StringPtr("testString")))
				Expect(translateDocumentOptionsModel.Target).To(Equal(core.StringPtr("testString")))
				Expect(translateDocumentOptionsModel.DocumentID).To(Equal(core.StringPtr("testString")))
				Expect(translateDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewTranslateOptions successfully`, func() {
				// Construct an instance of the TranslateOptions model
				translateOptionsText := []string{"testString"}
				translateOptionsModel := languageTranslatorService.NewTranslateOptions(translateOptionsText)
				translateOptionsModel.SetText([]string{"testString"})
				translateOptionsModel.SetModelID("testString")
				translateOptionsModel.SetSource("testString")
				translateOptionsModel.SetTarget("testString")
				translateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(translateOptionsModel).ToNot(BeNil())
				Expect(translateOptionsModel.Text).To(Equal([]string{"testString"}))
				Expect(translateOptionsModel.ModelID).To(Equal(core.StringPtr("testString")))
				Expect(translateOptionsModel.Source).To(Equal(core.StringPtr("testString")))
				Expect(translateOptionsModel.Target).To(Equal(core.StringPtr("testString")))
				Expect(translateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
