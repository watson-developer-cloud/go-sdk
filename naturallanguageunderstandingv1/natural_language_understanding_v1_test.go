package naturallanguageunderstandingv1_test

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/naturallanguageunderstandingv1"
)

var _ = Describe("NaturalLanguageUnderstandingV1", func() {
	Describe("Analyze(analyzeOptions *AnalyzeOptions)", func() {
		AnalyzePath := "/v1/analyze"
		version := "exampleString"
		Features := &naturallanguageunderstandingv1.Features{}
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Analyze text, HTML, or a public webpage", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(AnalyzePath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(AnalyzePath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"language":"xxx"}`)
			}))
			It("Succeed to call Analyze", func() {
				defer testServer.Close()

				testService, testServiceErr := naturallanguageunderstandingv1.
					NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
						URL:      testServer.URL,
						Version:  version,
						Username: username,
						Password: password,
					})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				AnalyzeOptions := testService.NewAnalyzeOptions(Features)
				returnValue, returnValueErr := testService.Analyze(AnalyzeOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetAnalyzeResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteModel(deleteModelOptions *DeleteModelOptions)", func() {
		DeleteModelPath := "/v1/models/{model_id}"
		version := "exampleString"
		ModelID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		DeleteModelPath = strings.Replace(DeleteModelPath, "{model_id}", ModelID, 1)
		Context("Successfully - Delete model", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(DeleteModelPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(DeleteModelPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"deleted":"yes"}`)
			}))
			It("Succeed to call DeleteModel", func() {
				defer testServer.Close()

				testService, testServiceErr := naturallanguageunderstandingv1.
					NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
						URL:      testServer.URL,
						Version:  version,
						Username: username,
						Password: password,
					})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				DeleteModelOptions := testService.NewDeleteModelOptions(ModelID)
				returnValue, returnValueErr := testService.DeleteModel(DeleteModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetDeleteModelResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListModels(listModelsOptions *ListModelsOptions)", func() {
		ListModelsPath := "/v1/models"
		version := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - List models", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(ListModelsPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(ListModelsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `[]`)
			}))
			It("Succeed to call ListModels", func() {
				defer testServer.Close()

				testService, testServiceErr := naturallanguageunderstandingv1.
					NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
						URL:      testServer.URL,
						Version:  version,
						Username: username,
						Password: password,
					})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				ListModelsOptions := testService.NewListModelsOptions()
				returnValue, returnValueErr := testService.ListModels(ListModelsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListModelsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
})
