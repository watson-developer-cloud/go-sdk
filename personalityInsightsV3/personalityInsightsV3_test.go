package personalityInsightsV3_test

import (
	"go-sdk/personalityInsightsV3"
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

var _ = Describe("PersonalityInsightsV3", func() {
	Describe("Get credentials from VCAP", func() {
		version := "exampleString"
		username := "hyphenated-user"
		password := "hyphenated-pass"
		VCAPservices := cfenv.Services{
			"conversation" : {
				{
					Name: "personality_insights",
					Tags: []string{},
					Credentials: map[string]interface{}{
						"url": "https://gateway.watsonplatform.net/personality-insights/api",
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
		Context("Successfully - Create PersonalityInsightsV3 with VCAP credentials", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
			}))
			It("Succeed to create PersonalityInsightsV3", func() {
				defer testServer.Close()

				testService, testServiceErr := personalityInsightsV3.NewPersonalityInsightsV3(&personalityInsightsV3.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				testService.Profile(personalityInsightsV3.NewProfileOptionsForPlain("exampleString"))
			})
		})
	})
	Describe("Profile(options *ProfileOptions)", func() {
		profilePath := "/v3/profile"
        version := "exampleString"
        contentType := "exampleString"
        profileOptions := personalityInsightsV3.NewProfileOptionsForPlain(contentType)
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
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call Profile", func() {
				defer testServer.Close()

				testService, testServiceErr := personalityInsightsV3.NewPersonalityInsightsV3(&personalityInsightsV3.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.Profile(profileOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("ProfileAsCsv(options *ProfileOptions)", func() {
		profileAsCsvPath := "/v3/profile"
        version := "exampleString"
        content := personalityInsightsV3.Content{}
        profileAsCsvOptions := personalityInsightsV3.NewProfileOptionsForContent(content)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Get profile as csv", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(profileAsCsvPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(profileAsCsvPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call ProfileAsCsv", func() {
				defer testServer.Close()

				testService, testServiceErr := personalityInsightsV3.NewPersonalityInsightsV3(&personalityInsightsV3.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.ProfileAsCsv(profileAsCsvOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
})
