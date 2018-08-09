package naturalLanguageUnderstandingV1_test

import (
	"go-sdk/naturalLanguageUnderstandingV1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"encoding/base64"
	"net/http/httptest"
	"net/http"
	"strings"
	"fmt"
	"os"
	"encoding/json"
	"github.com/cloudfoundry-community/go-cfenv"
)

var _ = Describe("NaturalLanguageUnderstandingV1", func() {
	Describe("Get credentials from VCAP", func() {
		version := "exampleString"
		username := "hyphenated-user"
		password := "hyphenated-pass"
		VCAPservices := cfenv.Services{
			"conversation" : {
				{
					Name: "natural-language-understanding",
					Tags: []string{},
					Credentials: map[string]interface{}{
						"url": "https://gateway.watsonplatform.net/natural-language-understanding/api",
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
		Context("Successfully - Create NaturalLanguageUnderstandingV1 with VCAP credentials", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
			}))
			It("Succeed to create NaturalLanguageUnderstandingV1", func() {
				defer testServer.Close()

				testService, testServiceErr := naturalLanguageUnderstandingV1.NewNaturalLanguageUnderstandingV1(&naturalLanguageUnderstandingV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				testService.ListModels(naturalLanguageUnderstandingV1.NewListModelsOptions())
			})
		})
	})
	Describe("Analyze(options *AnalyzeOptions)", func() {
		analyzePath := "/v1/analyze"
        version := "exampleString"
        parameters := naturalLanguageUnderstandingV1.Features{}
        analyzeOptions := naturalLanguageUnderstandingV1.NewAnalyzeOptions(parameters)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Analyze text, HTML, or a public webpage", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(analyzePath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(analyzePath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call Analyze", func() {
				defer testServer.Close()

				testService, testServiceErr := naturalLanguageUnderstandingV1.NewNaturalLanguageUnderstandingV1(&naturalLanguageUnderstandingV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.Analyze(analyzeOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteModel(options *DeleteModelOptions)", func() {
		deleteModelPath := "/v1/models/{model_id}"
        version := "exampleString"
        modelID := "exampleString"
        deleteModelOptions := naturalLanguageUnderstandingV1.NewDeleteModelOptions(modelID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        deleteModelPath = strings.Replace(deleteModelPath, "{model_id}", modelID, 1)
		Context("Successfully - Delete model", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(deleteModelPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(deleteModelPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call DeleteModel", func() {
				defer testServer.Close()

				testService, testServiceErr := naturalLanguageUnderstandingV1.NewNaturalLanguageUnderstandingV1(&naturalLanguageUnderstandingV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.DeleteModel(deleteModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("ListModels(options *ListModelsOptions)", func() {
		listModelsPath := "/v1/models"
        version := "exampleString"
        listModelsOptions := naturalLanguageUnderstandingV1.NewListModelsOptions()
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - List models", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(listModelsPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(listModelsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call ListModels", func() {
				defer testServer.Close()

				testService, testServiceErr := naturalLanguageUnderstandingV1.NewNaturalLanguageUnderstandingV1(&naturalLanguageUnderstandingV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.ListModels(listModelsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
})
