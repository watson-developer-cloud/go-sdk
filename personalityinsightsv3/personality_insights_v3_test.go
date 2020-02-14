package personalityinsightsv3_test

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

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	core "github.com/IBM/go-sdk-core/v3/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/personalityinsightsv3"
)

var _ = Describe("PersonalityInsightsV3", func() {
	Describe("Profile(profileOptions *ProfileOptions)", func() {
		profilePath := "/v3/profile"
		version := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Get profile", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(profilePath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(profilePath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprint(res, serializeJSON(createProfileResult()))
			}))
			It("Succeed to call Profile", func() {
				defer testServer.Close()

				testService, testServiceErr := personalityinsightsv3.NewPersonalityInsightsV3(&personalityinsightsv3.PersonalityInsightsV3Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BasicAuthenticator{
						Username: username,
						Password: password,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				_, _, returnValueErr := testService.Profile(nil)
				Expect(returnValueErr).NotTo(BeNil())

				// First test with invalid (incomplete) input
				content := new(personalityinsightsv3.Content)
				profileOptions := testService.
					NewProfileOptions().
					SetContent(content).
					SetContentType("text/plain")
				_, returnValue, returnValueErr := testService.Profile(profileOptions)
				Expect(returnValueErr).ToNot(BeNil())
				Expect(returnValue).To(BeNil())

				// Next, initialize content properly and retest
				contentItem := new(personalityinsightsv3.ContentItem)
				theContent := "theContent"
				contentItem.Content = &theContent
				content.ContentItems = []personalityinsightsv3.ContentItem{*contentItem}
				profileOptions = testService.
					NewProfileOptions().
					SetContent(content).
					SetContentType("application/json")
				result, returnValue, returnValueErr := testService.Profile(profileOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(createProfileResult()))
			})
		})
	})
	Describe("ProfileAsCsv(profileOptions *ProfileOptions)", func() {
		profileAsCsvPath := "/v3/profile"
		version := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		expectedCsvString := "field1,field2,field3,field4"
		Context("Successfully - Get profile as csv", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(profileAsCsvPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(profileAsCsvPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "text/plain")
				res.WriteHeader(201)
				fmt.Fprint(res, expectedCsvString)
			}))
			It("Succeed to call ProfileAsCsv", func() {
				defer testServer.Close()

				testService, testServiceErr := personalityinsightsv3.NewPersonalityInsightsV3(&personalityinsightsv3.PersonalityInsightsV3Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BasicAuthenticator{
						Username: username,
						Password: password,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				_, _, returnValueErr := testService.ProfileAsCsv(nil)
				Expect(returnValueErr).NotTo(BeNil())

				profileAsCsvOptions := testService.
					NewProfileOptions().
					SetBody("html").
					SetContentType("text/html")
				result, returnValue, returnValueErr := testService.ProfileAsCsv(profileAsCsvOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				result.Close()
			})
		})
	})
})

func serializeJSON(object interface{}) string {
	buff := new(bytes.Buffer)
	if err := json.NewEncoder(buff).Encode(object); err != nil {
		return "SERIALIZATION ERROR!!!"
	}
	return buff.String()
}

func createProfileResult() *personalityinsightsv3.Profile {
	personalityTrait := &personalityinsightsv3.Trait{
		TraitID:  core.StringPtr("personality"),
		Name:     core.StringPtr("trait1"),
		Category: core.StringPtr("category1"),
	}
	needsTrait := &personalityinsightsv3.Trait{
		TraitID:  core.StringPtr("needs"),
		Name:     core.StringPtr("trait2"),
		Category: core.StringPtr("category2"),
	}
	valuesTrait := &personalityinsightsv3.Trait{
		TraitID:  core.StringPtr("values"),
		Name:     core.StringPtr("trait3"),
		Category: core.StringPtr("category3"),
	}
	warning := &personalityinsightsv3.Warning{
		WarningID: core.StringPtr("warning1"),
		Message:   core.StringPtr("You messed up!"),
	}
	return &personalityinsightsv3.Profile{
		ProcessedLanguage: core.StringPtr("English"),
		WordCount:         core.Int64Ptr(38),
		Personality:       []personalityinsightsv3.Trait{*personalityTrait},
		Needs:             []personalityinsightsv3.Trait{*needsTrait},
		Values:            []personalityinsightsv3.Trait{*valuesTrait},
		Warnings:          []personalityinsightsv3.Warning{*warning},
	}
}
