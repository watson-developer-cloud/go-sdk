/**
 * (C) Copyright IBM Corp. 2018, 2021.
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

package personalityinsightsv3_test

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

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/v2/personalityinsightsv3"
)

var _ = Describe(`PersonalityInsightsV3`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		version := "testString"
		It(`Instantiate service client`, func() {
			personalityInsightsService, serviceErr := personalityinsightsv3.NewPersonalityInsightsV3(&personalityinsightsv3.PersonalityInsightsV3Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       core.StringPtr(version),
			})
			Expect(personalityInsightsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
			Expect(personalityInsightsService.Service.IsSSLDisabled()).To(BeFalse())
			personalityInsightsService.DisableSSLVerification()
			Expect(personalityInsightsService.Service.IsSSLDisabled()).To(BeTrue())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			personalityInsightsService, serviceErr := personalityinsightsv3.NewPersonalityInsightsV3(&personalityinsightsv3.PersonalityInsightsV3Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(personalityInsightsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			personalityInsightsService, serviceErr := personalityinsightsv3.NewPersonalityInsightsV3(&personalityinsightsv3.PersonalityInsightsV3Options{
				URL:     "https://personalityinsightsv3/api",
				Version: core.StringPtr(version),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(personalityInsightsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			personalityInsightsService, serviceErr := personalityinsightsv3.NewPersonalityInsightsV3(&personalityinsightsv3.PersonalityInsightsV3Options{})
			Expect(personalityInsightsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"PERSONALITY_INSIGHTS_URL":       "https://personalityinsightsv3/api",
				"PERSONALITY_INSIGHTS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				personalityInsightsService, serviceErr := personalityinsightsv3.NewPersonalityInsightsV3(&personalityinsightsv3.PersonalityInsightsV3Options{
					Version: core.StringPtr(version),
				})
				Expect(personalityInsightsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := personalityInsightsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != personalityInsightsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(personalityInsightsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(personalityInsightsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				personalityInsightsService, serviceErr := personalityinsightsv3.NewPersonalityInsightsV3(&personalityinsightsv3.PersonalityInsightsV3Options{
					URL:     "https://testService/api",
					Version: core.StringPtr(version),
				})
				Expect(personalityInsightsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(personalityInsightsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := personalityInsightsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != personalityInsightsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(personalityInsightsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(personalityInsightsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				personalityInsightsService, serviceErr := personalityinsightsv3.NewPersonalityInsightsV3(&personalityinsightsv3.PersonalityInsightsV3Options{
					Version: core.StringPtr(version),
				})
				err := personalityInsightsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(personalityInsightsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(personalityInsightsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := personalityInsightsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != personalityInsightsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(personalityInsightsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(personalityInsightsService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"PERSONALITY_INSIGHTS_URL":       "https://personalityinsightsv3/api",
				"PERSONALITY_INSIGHTS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			personalityInsightsService, serviceErr := personalityinsightsv3.NewPersonalityInsightsV3(&personalityinsightsv3.PersonalityInsightsV3Options{
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(personalityInsightsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"PERSONALITY_INSIGHTS_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			personalityInsightsService, serviceErr := personalityinsightsv3.NewPersonalityInsightsV3(&personalityinsightsv3.PersonalityInsightsV3Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(personalityInsightsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = personalityinsightsv3.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`Profile(profileOptions *ProfileOptions) - Operation response error`, func() {
		version := "testString"
		profilePath := "/v3/profile"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(profilePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Content-Type"]).ToNot(BeNil())
					Expect(req.Header["Content-Type"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					Expect(req.Header["Content-Language"]).ToNot(BeNil())
					Expect(req.Header["Content-Language"][0]).To(Equal(fmt.Sprintf("%v", "ar")))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "ar")))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for raw_scores query parameter
					// TODO: Add check for csv_headers query parameter
					// TODO: Add check for consumption_preferences query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke Profile with error: Operation response processing error`, func() {
				personalityInsightsService, serviceErr := personalityinsightsv3.NewPersonalityInsightsV3(&personalityinsightsv3.PersonalityInsightsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(personalityInsightsService).ToNot(BeNil())

				// Construct an instance of the ContentItem model
				contentItemModel := new(personalityinsightsv3.ContentItem)
				contentItemModel.Content = core.StringPtr("testString")
				contentItemModel.ID = core.StringPtr("testString")
				contentItemModel.Created = core.Int64Ptr(int64(26))
				contentItemModel.Updated = core.Int64Ptr(int64(26))
				contentItemModel.Contenttype = core.StringPtr("text/plain")
				contentItemModel.Language = core.StringPtr("ar")
				contentItemModel.Parentid = core.StringPtr("testString")
				contentItemModel.Reply = core.BoolPtr(true)
				contentItemModel.Forward = core.BoolPtr(true)

				// Construct an instance of the Content model
				contentModel := new(personalityinsightsv3.Content)
				contentModel.ContentItems = []personalityinsightsv3.ContentItem{*contentItemModel}

				// Construct an instance of the ProfileOptions model
				profileOptionsModel := new(personalityinsightsv3.ProfileOptions)
				profileOptionsModel.Content = contentModel
				profileOptionsModel.ContentType = core.StringPtr("application/json")
				profileOptionsModel.ContentLanguage = core.StringPtr("ar")
				profileOptionsModel.AcceptLanguage = core.StringPtr("ar")
				profileOptionsModel.RawScores = core.BoolPtr(true)
				profileOptionsModel.CSVHeaders = core.BoolPtr(true)
				profileOptionsModel.ConsumptionPreferences = core.BoolPtr(true)
				profileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := personalityInsightsService.Profile(profileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				personalityInsightsService.EnableRetries(0, 0)
				result, response, operationErr = personalityInsightsService.Profile(profileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Profile(profileOptions *ProfileOptions)`, func() {
		version := "testString"
		profilePath := "/v3/profile"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(profilePath))
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
					Expect(req.Header["Content-Language"][0]).To(Equal(fmt.Sprintf("%v", "ar")))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "ar")))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for raw_scores query parameter
					// TODO: Add check for csv_headers query parameter
					// TODO: Add check for consumption_preferences query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"processed_language": "ar", "word_count": 9, "word_count_message": "WordCountMessage", "personality": [{"trait_id": "TraitID", "name": "Name", "category": "personality", "percentile": 10, "raw_score": 8, "significant": false}], "needs": [{"trait_id": "TraitID", "name": "Name", "category": "personality", "percentile": 10, "raw_score": 8, "significant": false}], "values": [{"trait_id": "TraitID", "name": "Name", "category": "personality", "percentile": 10, "raw_score": 8, "significant": false}], "behavior": [{"trait_id": "TraitID", "name": "Name", "category": "Category", "percentage": 10}], "consumption_preferences": [{"consumption_preference_category_id": "ConsumptionPreferenceCategoryID", "name": "Name", "consumption_preferences": [{"consumption_preference_id": "ConsumptionPreferenceID", "name": "Name", "score": 0.0}]}], "warnings": [{"warning_id": "WORD_COUNT_MESSAGE", "message": "Message"}]}`)
				}))
			})
			It(`Invoke Profile successfully with retries`, func() {
				personalityInsightsService, serviceErr := personalityinsightsv3.NewPersonalityInsightsV3(&personalityinsightsv3.PersonalityInsightsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(personalityInsightsService).ToNot(BeNil())
				personalityInsightsService.EnableRetries(0, 0)

				// Construct an instance of the ContentItem model
				contentItemModel := new(personalityinsightsv3.ContentItem)
				contentItemModel.Content = core.StringPtr("testString")
				contentItemModel.ID = core.StringPtr("testString")
				contentItemModel.Created = core.Int64Ptr(int64(26))
				contentItemModel.Updated = core.Int64Ptr(int64(26))
				contentItemModel.Contenttype = core.StringPtr("text/plain")
				contentItemModel.Language = core.StringPtr("ar")
				contentItemModel.Parentid = core.StringPtr("testString")
				contentItemModel.Reply = core.BoolPtr(true)
				contentItemModel.Forward = core.BoolPtr(true)

				// Construct an instance of the Content model
				contentModel := new(personalityinsightsv3.Content)
				contentModel.ContentItems = []personalityinsightsv3.ContentItem{*contentItemModel}

				// Construct an instance of the ProfileOptions model
				profileOptionsModel := new(personalityinsightsv3.ProfileOptions)
				profileOptionsModel.Content = contentModel
				profileOptionsModel.ContentType = core.StringPtr("application/json")
				profileOptionsModel.ContentLanguage = core.StringPtr("ar")
				profileOptionsModel.AcceptLanguage = core.StringPtr("ar")
				profileOptionsModel.RawScores = core.BoolPtr(true)
				profileOptionsModel.CSVHeaders = core.BoolPtr(true)
				profileOptionsModel.ConsumptionPreferences = core.BoolPtr(true)
				profileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := personalityInsightsService.ProfileWithContext(ctx, profileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				personalityInsightsService.DisableRetries()
				result, response, operationErr := personalityInsightsService.Profile(profileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = personalityInsightsService.ProfileWithContext(ctx, profileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(profilePath))
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
					Expect(req.Header["Content-Language"][0]).To(Equal(fmt.Sprintf("%v", "ar")))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "ar")))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for raw_scores query parameter
					// TODO: Add check for csv_headers query parameter
					// TODO: Add check for consumption_preferences query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"processed_language": "ar", "word_count": 9, "word_count_message": "WordCountMessage", "personality": [{"trait_id": "TraitID", "name": "Name", "category": "personality", "percentile": 10, "raw_score": 8, "significant": false}], "needs": [{"trait_id": "TraitID", "name": "Name", "category": "personality", "percentile": 10, "raw_score": 8, "significant": false}], "values": [{"trait_id": "TraitID", "name": "Name", "category": "personality", "percentile": 10, "raw_score": 8, "significant": false}], "behavior": [{"trait_id": "TraitID", "name": "Name", "category": "Category", "percentage": 10}], "consumption_preferences": [{"consumption_preference_category_id": "ConsumptionPreferenceCategoryID", "name": "Name", "consumption_preferences": [{"consumption_preference_id": "ConsumptionPreferenceID", "name": "Name", "score": 0.0}]}], "warnings": [{"warning_id": "WORD_COUNT_MESSAGE", "message": "Message"}]}`)
				}))
			})
			It(`Invoke Profile successfully`, func() {
				personalityInsightsService, serviceErr := personalityinsightsv3.NewPersonalityInsightsV3(&personalityinsightsv3.PersonalityInsightsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(personalityInsightsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := personalityInsightsService.Profile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ContentItem model
				contentItemModel := new(personalityinsightsv3.ContentItem)
				contentItemModel.Content = core.StringPtr("testString")
				contentItemModel.ID = core.StringPtr("testString")
				contentItemModel.Created = core.Int64Ptr(int64(26))
				contentItemModel.Updated = core.Int64Ptr(int64(26))
				contentItemModel.Contenttype = core.StringPtr("text/plain")
				contentItemModel.Language = core.StringPtr("ar")
				contentItemModel.Parentid = core.StringPtr("testString")
				contentItemModel.Reply = core.BoolPtr(true)
				contentItemModel.Forward = core.BoolPtr(true)

				// Construct an instance of the Content model
				contentModel := new(personalityinsightsv3.Content)
				contentModel.ContentItems = []personalityinsightsv3.ContentItem{*contentItemModel}

				// Construct an instance of the ProfileOptions model
				profileOptionsModel := new(personalityinsightsv3.ProfileOptions)
				profileOptionsModel.Content = contentModel
				profileOptionsModel.ContentType = core.StringPtr("application/json")
				profileOptionsModel.ContentLanguage = core.StringPtr("ar")
				profileOptionsModel.AcceptLanguage = core.StringPtr("ar")
				profileOptionsModel.RawScores = core.BoolPtr(true)
				profileOptionsModel.CSVHeaders = core.BoolPtr(true)
				profileOptionsModel.ConsumptionPreferences = core.BoolPtr(true)
				profileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = personalityInsightsService.Profile(profileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke Profile with error: Operation request error`, func() {
				personalityInsightsService, serviceErr := personalityinsightsv3.NewPersonalityInsightsV3(&personalityinsightsv3.PersonalityInsightsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(personalityInsightsService).ToNot(BeNil())

				// Construct an instance of the ContentItem model
				contentItemModel := new(personalityinsightsv3.ContentItem)
				contentItemModel.Content = core.StringPtr("testString")
				contentItemModel.ID = core.StringPtr("testString")
				contentItemModel.Created = core.Int64Ptr(int64(26))
				contentItemModel.Updated = core.Int64Ptr(int64(26))
				contentItemModel.Contenttype = core.StringPtr("text/plain")
				contentItemModel.Language = core.StringPtr("ar")
				contentItemModel.Parentid = core.StringPtr("testString")
				contentItemModel.Reply = core.BoolPtr(true)
				contentItemModel.Forward = core.BoolPtr(true)

				// Construct an instance of the Content model
				contentModel := new(personalityinsightsv3.Content)
				contentModel.ContentItems = []personalityinsightsv3.ContentItem{*contentItemModel}

				// Construct an instance of the ProfileOptions model
				profileOptionsModel := new(personalityinsightsv3.ProfileOptions)
				profileOptionsModel.Content = contentModel
				profileOptionsModel.ContentType = core.StringPtr("application/json")
				profileOptionsModel.ContentLanguage = core.StringPtr("ar")
				profileOptionsModel.AcceptLanguage = core.StringPtr("ar")
				profileOptionsModel.RawScores = core.BoolPtr(true)
				profileOptionsModel.CSVHeaders = core.BoolPtr(true)
				profileOptionsModel.ConsumptionPreferences = core.BoolPtr(true)
				profileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := personalityInsightsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := personalityInsightsService.Profile(profileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke Profile successfully`, func() {
				personalityInsightsService, serviceErr := personalityinsightsv3.NewPersonalityInsightsV3(&personalityinsightsv3.PersonalityInsightsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(personalityInsightsService).ToNot(BeNil())

				// Construct an instance of the ContentItem model
				contentItemModel := new(personalityinsightsv3.ContentItem)
				contentItemModel.Content = core.StringPtr("testString")
				contentItemModel.ID = core.StringPtr("testString")
				contentItemModel.Created = core.Int64Ptr(int64(26))
				contentItemModel.Updated = core.Int64Ptr(int64(26))
				contentItemModel.Contenttype = core.StringPtr("text/plain")
				contentItemModel.Language = core.StringPtr("ar")
				contentItemModel.Parentid = core.StringPtr("testString")
				contentItemModel.Reply = core.BoolPtr(true)
				contentItemModel.Forward = core.BoolPtr(true)

				// Construct an instance of the Content model
				contentModel := new(personalityinsightsv3.Content)
				contentModel.ContentItems = []personalityinsightsv3.ContentItem{*contentItemModel}

				// Construct an instance of the ProfileOptions model
				profileOptionsModel := new(personalityinsightsv3.ProfileOptions)
				profileOptionsModel.Content = contentModel
				profileOptionsModel.ContentType = core.StringPtr("application/json")
				profileOptionsModel.ContentLanguage = core.StringPtr("ar")
				profileOptionsModel.AcceptLanguage = core.StringPtr("ar")
				profileOptionsModel.RawScores = core.BoolPtr(true)
				profileOptionsModel.CSVHeaders = core.BoolPtr(true)
				profileOptionsModel.ConsumptionPreferences = core.BoolPtr(true)
				profileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := personalityInsightsService.Profile(profileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ProfileAsCSV(profileOptions *ProfileOptions)`, func() {
		version := "testString"
		profileAsCSVPath := "/v3/profile"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(profileAsCSVPath))
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
					Expect(req.Header["Content-Language"][0]).To(Equal(fmt.Sprintf("%v", "ar")))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "ar")))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for raw_scores query parameter
					// TODO: Add check for csv_headers query parameter
					// TODO: Add check for consumption_preferences query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "text/csv")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `"OperationResponse"`)
				}))
			})
			It(`Invoke ProfileAsCSV successfully with retries`, func() {
				personalityInsightsService, serviceErr := personalityinsightsv3.NewPersonalityInsightsV3(&personalityinsightsv3.PersonalityInsightsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(personalityInsightsService).ToNot(BeNil())
				personalityInsightsService.EnableRetries(0, 0)

				// Construct an instance of the ContentItem model
				contentItemModel := new(personalityinsightsv3.ContentItem)
				contentItemModel.Content = core.StringPtr("testString")
				contentItemModel.ID = core.StringPtr("testString")
				contentItemModel.Created = core.Int64Ptr(int64(26))
				contentItemModel.Updated = core.Int64Ptr(int64(26))
				contentItemModel.Contenttype = core.StringPtr("text/plain")
				contentItemModel.Language = core.StringPtr("ar")
				contentItemModel.Parentid = core.StringPtr("testString")
				contentItemModel.Reply = core.BoolPtr(true)
				contentItemModel.Forward = core.BoolPtr(true)

				// Construct an instance of the Content model
				contentModel := new(personalityinsightsv3.Content)
				contentModel.ContentItems = []personalityinsightsv3.ContentItem{*contentItemModel}

				// Construct an instance of the ProfileOptions model
				profileOptionsModel := new(personalityinsightsv3.ProfileOptions)
				profileOptionsModel.Content = contentModel
				profileOptionsModel.ContentType = core.StringPtr("application/json")
				profileOptionsModel.ContentLanguage = core.StringPtr("ar")
				profileOptionsModel.AcceptLanguage = core.StringPtr("ar")
				profileOptionsModel.RawScores = core.BoolPtr(true)
				profileOptionsModel.CSVHeaders = core.BoolPtr(true)
				profileOptionsModel.ConsumptionPreferences = core.BoolPtr(true)
				profileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := personalityInsightsService.ProfileAsCSVWithContext(ctx, profileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				personalityInsightsService.DisableRetries()
				result, response, operationErr := personalityInsightsService.ProfileAsCSV(profileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = personalityInsightsService.ProfileAsCSVWithContext(ctx, profileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(profileAsCSVPath))
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
					Expect(req.Header["Content-Language"][0]).To(Equal(fmt.Sprintf("%v", "ar")))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "ar")))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for raw_scores query parameter
					// TODO: Add check for csv_headers query parameter
					// TODO: Add check for consumption_preferences query parameter
					// Set mock response
					res.Header().Set("Content-type", "text/csv")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `"OperationResponse"`)
				}))
			})
			It(`Invoke ProfileAsCSV successfully`, func() {
				personalityInsightsService, serviceErr := personalityinsightsv3.NewPersonalityInsightsV3(&personalityinsightsv3.PersonalityInsightsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(personalityInsightsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := personalityInsightsService.ProfileAsCSV(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ContentItem model
				contentItemModel := new(personalityinsightsv3.ContentItem)
				contentItemModel.Content = core.StringPtr("testString")
				contentItemModel.ID = core.StringPtr("testString")
				contentItemModel.Created = core.Int64Ptr(int64(26))
				contentItemModel.Updated = core.Int64Ptr(int64(26))
				contentItemModel.Contenttype = core.StringPtr("text/plain")
				contentItemModel.Language = core.StringPtr("ar")
				contentItemModel.Parentid = core.StringPtr("testString")
				contentItemModel.Reply = core.BoolPtr(true)
				contentItemModel.Forward = core.BoolPtr(true)

				// Construct an instance of the Content model
				contentModel := new(personalityinsightsv3.Content)
				contentModel.ContentItems = []personalityinsightsv3.ContentItem{*contentItemModel}

				// Construct an instance of the ProfileOptions model
				profileOptionsModel := new(personalityinsightsv3.ProfileOptions)
				profileOptionsModel.Content = contentModel
				profileOptionsModel.ContentType = core.StringPtr("application/json")
				profileOptionsModel.ContentLanguage = core.StringPtr("ar")
				profileOptionsModel.AcceptLanguage = core.StringPtr("ar")
				profileOptionsModel.RawScores = core.BoolPtr(true)
				profileOptionsModel.CSVHeaders = core.BoolPtr(true)
				profileOptionsModel.ConsumptionPreferences = core.BoolPtr(true)
				profileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = personalityInsightsService.ProfileAsCSV(profileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ProfileAsCSV with error: Operation request error`, func() {
				personalityInsightsService, serviceErr := personalityinsightsv3.NewPersonalityInsightsV3(&personalityinsightsv3.PersonalityInsightsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(personalityInsightsService).ToNot(BeNil())

				// Construct an instance of the ContentItem model
				contentItemModel := new(personalityinsightsv3.ContentItem)
				contentItemModel.Content = core.StringPtr("testString")
				contentItemModel.ID = core.StringPtr("testString")
				contentItemModel.Created = core.Int64Ptr(int64(26))
				contentItemModel.Updated = core.Int64Ptr(int64(26))
				contentItemModel.Contenttype = core.StringPtr("text/plain")
				contentItemModel.Language = core.StringPtr("ar")
				contentItemModel.Parentid = core.StringPtr("testString")
				contentItemModel.Reply = core.BoolPtr(true)
				contentItemModel.Forward = core.BoolPtr(true)

				// Construct an instance of the Content model
				contentModel := new(personalityinsightsv3.Content)
				contentModel.ContentItems = []personalityinsightsv3.ContentItem{*contentItemModel}

				// Construct an instance of the ProfileOptions model
				profileOptionsModel := new(personalityinsightsv3.ProfileOptions)
				profileOptionsModel.Content = contentModel
				profileOptionsModel.ContentType = core.StringPtr("application/json")
				profileOptionsModel.ContentLanguage = core.StringPtr("ar")
				profileOptionsModel.AcceptLanguage = core.StringPtr("ar")
				profileOptionsModel.RawScores = core.BoolPtr(true)
				profileOptionsModel.CSVHeaders = core.BoolPtr(true)
				profileOptionsModel.ConsumptionPreferences = core.BoolPtr(true)
				profileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := personalityInsightsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := personalityInsightsService.ProfileAsCSV(profileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ProfileAsCSV successfully`, func() {
				personalityInsightsService, serviceErr := personalityinsightsv3.NewPersonalityInsightsV3(&personalityinsightsv3.PersonalityInsightsV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(personalityInsightsService).ToNot(BeNil())

				// Construct an instance of the ContentItem model
				contentItemModel := new(personalityinsightsv3.ContentItem)
				contentItemModel.Content = core.StringPtr("testString")
				contentItemModel.ID = core.StringPtr("testString")
				contentItemModel.Created = core.Int64Ptr(int64(26))
				contentItemModel.Updated = core.Int64Ptr(int64(26))
				contentItemModel.Contenttype = core.StringPtr("text/plain")
				contentItemModel.Language = core.StringPtr("ar")
				contentItemModel.Parentid = core.StringPtr("testString")
				contentItemModel.Reply = core.BoolPtr(true)
				contentItemModel.Forward = core.BoolPtr(true)

				// Construct an instance of the Content model
				contentModel := new(personalityinsightsv3.Content)
				contentModel.ContentItems = []personalityinsightsv3.ContentItem{*contentItemModel}

				// Construct an instance of the ProfileOptions model
				profileOptionsModel := new(personalityinsightsv3.ProfileOptions)
				profileOptionsModel.Content = contentModel
				profileOptionsModel.ContentType = core.StringPtr("application/json")
				profileOptionsModel.ContentLanguage = core.StringPtr("ar")
				profileOptionsModel.AcceptLanguage = core.StringPtr("ar")
				profileOptionsModel.RawScores = core.BoolPtr(true)
				profileOptionsModel.CSVHeaders = core.BoolPtr(true)
				profileOptionsModel.ConsumptionPreferences = core.BoolPtr(true)
				profileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := personalityInsightsService.ProfileAsCSV(profileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify empty byte buffer.
				Expect(result).ToNot(BeNil())
				buffer, operationErr := ioutil.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(len(buffer)).To(Equal(0))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			version := "testString"
			personalityInsightsService, _ := personalityinsightsv3.NewPersonalityInsightsV3(&personalityinsightsv3.PersonalityInsightsV3Options{
				URL:           "http://personalityinsightsv3modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       core.StringPtr(version),
			})
			It(`Invoke NewContent successfully`, func() {
				contentItems := []personalityinsightsv3.ContentItem{}
				model, err := personalityInsightsService.NewContent(contentItems)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewContentItem successfully`, func() {
				content := "testString"
				model, err := personalityInsightsService.NewContentItem(content)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewProfileOptions successfully`, func() {
				// Construct an instance of the ContentItem model
				contentItemModel := new(personalityinsightsv3.ContentItem)
				Expect(contentItemModel).ToNot(BeNil())
				contentItemModel.Content = core.StringPtr("testString")
				contentItemModel.ID = core.StringPtr("testString")
				contentItemModel.Created = core.Int64Ptr(int64(26))
				contentItemModel.Updated = core.Int64Ptr(int64(26))
				contentItemModel.Contenttype = core.StringPtr("text/plain")
				contentItemModel.Language = core.StringPtr("ar")
				contentItemModel.Parentid = core.StringPtr("testString")
				contentItemModel.Reply = core.BoolPtr(true)
				contentItemModel.Forward = core.BoolPtr(true)
				Expect(contentItemModel.Content).To(Equal(core.StringPtr("testString")))
				Expect(contentItemModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(contentItemModel.Created).To(Equal(core.Int64Ptr(int64(26))))
				Expect(contentItemModel.Updated).To(Equal(core.Int64Ptr(int64(26))))
				Expect(contentItemModel.Contenttype).To(Equal(core.StringPtr("text/plain")))
				Expect(contentItemModel.Language).To(Equal(core.StringPtr("ar")))
				Expect(contentItemModel.Parentid).To(Equal(core.StringPtr("testString")))
				Expect(contentItemModel.Reply).To(Equal(core.BoolPtr(true)))
				Expect(contentItemModel.Forward).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the Content model
				contentModel := new(personalityinsightsv3.Content)
				Expect(contentModel).ToNot(BeNil())
				contentModel.ContentItems = []personalityinsightsv3.ContentItem{*contentItemModel}
				Expect(contentModel.ContentItems).To(Equal([]personalityinsightsv3.ContentItem{*contentItemModel}))

				// Construct an instance of the ProfileOptions model
				profileOptionsModel := personalityInsightsService.NewProfileOptions()
				profileOptionsModel.SetContent(contentModel)
				profileOptionsModel.SetBody("testString")
				profileOptionsModel.SetContentType("application/json")
				profileOptionsModel.SetContentLanguage("ar")
				profileOptionsModel.SetAcceptLanguage("ar")
				profileOptionsModel.SetRawScores(true)
				profileOptionsModel.SetCSVHeaders(true)
				profileOptionsModel.SetConsumptionPreferences(true)
				profileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(profileOptionsModel).ToNot(BeNil())
				Expect(profileOptionsModel.Content).To(Equal(contentModel))
				Expect(profileOptionsModel.Body).To(Equal(core.StringPtr("testString")))
				Expect(profileOptionsModel.ContentType).To(Equal(core.StringPtr("application/json")))
				Expect(profileOptionsModel.ContentLanguage).To(Equal(core.StringPtr("ar")))
				Expect(profileOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("ar")))
				Expect(profileOptionsModel.RawScores).To(Equal(core.BoolPtr(true)))
				Expect(profileOptionsModel.CSVHeaders).To(Equal(core.BoolPtr(true)))
				Expect(profileOptionsModel.ConsumptionPreferences).To(Equal(core.BoolPtr(true)))
				Expect(profileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
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

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
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
