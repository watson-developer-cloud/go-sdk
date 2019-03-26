package naturallanguageclassifierv1_test

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/naturallanguageclassifierv1"
)

var _ = Describe("NaturalLanguageClassifierV1", func() {
	Describe("ClassifyCollection(classifyCollectionOptions *ClassifyCollectionOptions)", func() {
		ClassifyCollectionPath := "/v1/classifiers/{classifier_id}/classify_collection"
		ClassifierID := "exampleString"
		Collection := []naturallanguageclassifierv1.ClassifyInput{}
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		ClassifyCollectionPath = strings.Replace(ClassifyCollectionPath, "{classifier_id}", ClassifierID, 1)
		Context("Successfully - Classify multiple phrases", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(ClassifyCollectionPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"classifier_id":"xxx"}`)
			}))
			It("Succeed to call ClassifyCollection", func() {
				defer testServer.Close()

				testService, testServiceErr := naturallanguageclassifierv1.
					NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
						URL:      testServer.URL,
						Username: username,
						Password: password,
					})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				classifyCollectionOptions := testService.NewClassifyCollectionOptions(ClassifierID, Collection)
				returnValue, returnValueErr := testService.ClassifyCollection(classifyCollectionOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetClassifyCollectionResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateClassifier(createClassifierOptions *CreateClassifierOptions)", func() {
		CreateClassifierPath := "/v1/classifiers"
		pwd, _ := os.Getwd()
		Metadata, metadataErr := os.Open(pwd + "/../resources/weather_training_metadata.json")
		if metadataErr != nil {
			fmt.Println(metadataErr)
		}
		data, dataErr := os.Open(pwd + "/../resources/weather_training_data.csv")
		if dataErr != nil {
			fmt.Println(dataErr)
		}
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Create classifier", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(CreateClassifierPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
			}))
			It("Succeed to call CreateClassifier", func() {
				defer testServer.Close()

				testService, testServiceErr := naturallanguageclassifierv1.
					NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
						URL:      testServer.URL,
						Username: username,
						Password: password,
					})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				createClassifierOptions := testService.NewCreateClassifierOptions(Metadata, data)
				returnValue, returnValueErr := testService.CreateClassifier(createClassifierOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetCreateClassifierResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteClassifier(deleteClassifierOptions *DeleteClassifierOptions)", func() {
		DeleteClassifierPath := "/v1/classifiers/{classifier_id}"
		ClassifierID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		DeleteClassifierPath = strings.Replace(DeleteClassifierPath, "{classifier_id}", ClassifierID, 1)
		Context("Successfully - Delete classifier", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(DeleteClassifierPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call DeleteClassifier", func() {
				defer testServer.Close()

				testService, testServiceErr := naturallanguageclassifierv1.
					NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
						URL:      testServer.URL,
						Username: username,
						Password: password,
					})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				deleteClassifierOptions := testService.NewDeleteClassifierOptions(ClassifierID)
				returnValue, returnValueErr := testService.DeleteClassifier(deleteClassifierOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("GetClassifier(getClassifierOptions *GetClassifierOptions)", func() {
		GetClassifierPath := "/v1/classifiers/{classifier_id}"
		ClassifierID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		GetClassifierPath = strings.Replace(GetClassifierPath, "{classifier_id}", ClassifierID, 1)
		Context("Successfully - Get information about a classifier", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(GetClassifierPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"name":"xxx"}`)
			}))
			It("Succeed to call GetClassifier", func() {
				defer testServer.Close()

				testService, testServiceErr := naturallanguageclassifierv1.
					NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
						URL:      testServer.URL,
						Username: username,
						Password: password,
					})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				GetClassifierOptions := testService.NewGetClassifierOptions(ClassifierID)
				returnValue, returnValueErr := testService.GetClassifier(GetClassifierOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetClassifierResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListClassifiers(listClassifiersOptions *ListClassifiersOptions)", func() {
		ListClassifiersPath := "/v1/classifiers"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - List classifiers", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(ListClassifiersPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `[]`)
			}))
			It("Succeed to call ListClassifiers", func() {
				defer testServer.Close()

				testService, testServiceErr := naturallanguageclassifierv1.
					NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
						URL:      testServer.URL,
						Username: username,
						Password: password,
					})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				ListClassifiersOptions := testService.NewListClassifiersOptions()
				returnValue, returnValueErr := testService.ListClassifiers(ListClassifiersOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListClassifiersResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
})
