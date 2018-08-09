package naturalLanguageClassifierV1_test

import (
	"go-sdk/naturalLanguageClassifierV1"
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

var _ = Describe("NaturalLanguageClassifierV1", func() {
	Describe("Get credentials from VCAP", func() {
		version := "exampleString"
		username := "hyphenated-user"
		password := "hyphenated-pass"
		VCAPservices := cfenv.Services{
			"conversation" : {
				{
					Name: "natural_language_classifier",
					Tags: []string{},
					Credentials: map[string]interface{}{
						"url": "https://gateway.watsonplatform.net/natural-language-classifier/api",
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
		Context("Successfully - Create NaturalLanguageClassifierV1 with VCAP credentials", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
			}))
			It("Succeed to create NaturalLanguageClassifierV1", func() {
				defer testServer.Close()

				testService, testServiceErr := naturalLanguageClassifierV1.NewNaturalLanguageClassifierV1(&naturalLanguageClassifierV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				testService.ListClassifiers(naturalLanguageClassifierV1.NewListClassifiersOptions())
			})
		})
	})
	Describe("Classify(options *ClassifyOptions)", func() {
		classifyPath := "/v1/classifiers/{classifier_id}/classify"
        version := "exampleString"
        classifierID := "exampleString"
        body := "exampleString"
        classifyOptions := naturalLanguageClassifierV1.NewClassifyOptions(classifierID, body)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        classifyPath = strings.Replace(classifyPath, "{classifier_id}", classifierID, 1)
		Context("Successfully - Classify a phrase", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(classifyPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call Classify", func() {
				defer testServer.Close()

				testService, testServiceErr := naturalLanguageClassifierV1.NewNaturalLanguageClassifierV1(&naturalLanguageClassifierV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.Classify(classifyOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("ClassifyCollection(options *ClassifyCollectionOptions)", func() {
		classifyCollectionPath := "/v1/classifiers/{classifier_id}/classify_collection"
        version := "exampleString"
        classifierID := "exampleString"
        body := []naturalLanguageClassifierV1.ClassifyInput{}
        classifyCollectionOptions := naturalLanguageClassifierV1.NewClassifyCollectionOptions(classifierID, body)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        classifyCollectionPath = strings.Replace(classifyCollectionPath, "{classifier_id}", classifierID, 1)
		Context("Successfully - Classify multiple phrases", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(classifyCollectionPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call ClassifyCollection", func() {
				defer testServer.Close()

				testService, testServiceErr := naturalLanguageClassifierV1.NewNaturalLanguageClassifierV1(&naturalLanguageClassifierV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.ClassifyCollection(classifyCollectionOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteClassifier(options *DeleteClassifierOptions)", func() {
		deleteClassifierPath := "/v1/classifiers/{classifier_id}"
        version := "exampleString"
        classifierID := "exampleString"
        deleteClassifierOptions := naturalLanguageClassifierV1.NewDeleteClassifierOptions(classifierID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        deleteClassifierPath = strings.Replace(deleteClassifierPath, "{classifier_id}", classifierID, 1)
		Context("Successfully - Delete classifier", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(deleteClassifierPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call DeleteClassifier", func() {
				defer testServer.Close()

				testService, testServiceErr := naturalLanguageClassifierV1.NewNaturalLanguageClassifierV1(&naturalLanguageClassifierV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.DeleteClassifier(deleteClassifierOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("GetClassifier(options *GetClassifierOptions)", func() {
		getClassifierPath := "/v1/classifiers/{classifier_id}"
        version := "exampleString"
        classifierID := "exampleString"
        getClassifierOptions := naturalLanguageClassifierV1.NewGetClassifierOptions(classifierID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        getClassifierPath = strings.Replace(getClassifierPath, "{classifier_id}", classifierID, 1)
		Context("Successfully - Get information about a classifier", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(getClassifierPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call GetClassifier", func() {
				defer testServer.Close()

				testService, testServiceErr := naturalLanguageClassifierV1.NewNaturalLanguageClassifierV1(&naturalLanguageClassifierV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.GetClassifier(getClassifierOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("ListClassifiers(options *ListClassifiersOptions)", func() {
		listClassifiersPath := "/v1/classifiers"
        version := "exampleString"
        listClassifiersOptions := naturalLanguageClassifierV1.NewListClassifiersOptions()
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - List classifiers", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(listClassifiersPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call ListClassifiers", func() {
				defer testServer.Close()

				testService, testServiceErr := naturalLanguageClassifierV1.NewNaturalLanguageClassifierV1(&naturalLanguageClassifierV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.ListClassifiers(listClassifiersOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
})
