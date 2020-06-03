/**
 * (C) Copyright IBM Corp. 2020.
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
	"fmt"
	"github.com/IBM/go-sdk-core/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/personalityinsightsv3"
	"net/http"
	"net/http/httptest"
)

var _ = Describe(`PersonalityInsightsV3`, func() {
	Describe(`Profile(profileOptions *ProfileOptions)`, func() {
		profilePath := "/v3/profile"
		version := "exampleString"
		bearerToken := "0ui9876453"
		Context(`Successfully - Get profile`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(profilePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"processed_language": "fake_ProcessedLanguage", "word_count": 9, "personality": [], "needs": [], "values": [], "warnings": []}`)
			}))
			It(`Succeed to call Profile`, func() {
				defer testServer.Close()

				testService, testServiceErr := personalityinsightsv3.NewPersonalityInsightsV3(&personalityinsightsv3.PersonalityInsightsV3Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.Profile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				profileOptions := testService.NewProfileOptions()
				result, response, operationErr = testService.Profile(profileOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ProfileAsCsv(profileOptions *ProfileOptions)`, func() {
		profileAsCsvPath := "/v3/profile"
		version := "exampleString"
		bearerToken := "0ui9876453"
		Context(`Successfully - Get profile as csv`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(profileAsCsvPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `Contents of response byte-stream...`)
			}))
			It(`Succeed to call ProfileAsCsv`, func() {
				defer testServer.Close()

				testService, testServiceErr := personalityinsightsv3.NewPersonalityInsightsV3(&personalityinsightsv3.PersonalityInsightsV3Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ProfileAsCsv(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				profileAsCsvOptions := testService.NewProfileAsCsvOptions()
				result, response, operationErr = testService.ProfileAsCsv(profileAsCsvOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("Model constructor tests", func() {
		Context("with a sample service", func() {
			version := "1970-01-01"
			testService, _ := personalityinsightsv3.NewPersonalityInsightsV3(&personalityinsightsv3.PersonalityInsightsV3Options{
				URL:           "http://personalityinsightsv3modelgenerator.com",
				Version:       version,
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It("should call NewContent successfully", func() {
				contentItems := []personalityinsightsv3.ContentItem{}
				model, err := testService.NewContent(contentItems)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It("should call NewContentItem successfully", func() {
				content := "exampleString"
				model, err := testService.NewContentItem(content)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
})
