package visualRecognitionV3_test

import (
	"go-sdk/visualRecognitionV3"
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

var _ = Describe("VisualRecognitionV3", func() {
	Describe("Get credentials from VCAP", func() {
		version := "exampleString"
		username := "hyphenated-user"
		password := "hyphenated-pass"
		VCAPservices := cfenv.Services{
			"conversation" : {
				{
					Name: "watson_vision_combined",
					Tags: []string{},
					Credentials: map[string]interface{}{
						"url": "https://gateway.watsonplatform.net/visual-recognition/api",
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
		Context("Successfully - Create VisualRecognitionV3 with VCAP credentials", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
			}))
			It("Succeed to create VisualRecognitionV3", func() {
				defer testServer.Close()

				testService, testServiceErr := visualRecognitionV3.NewVisualRecognitionV3(&visualRecognitionV3.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				testService.ListClassifiers(visualRecognitionV3.NewListClassifiersOptions())
			})
		})
	})
	Describe("Classify(options *ClassifyOptions)", func() {
		classifyPath := "/v3/classify"
        version := "exampleString"
        classifyOptions := visualRecognitionV3.NewClassifyOptions()
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Classify images", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(classifyPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(classifyPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call Classify", func() {
				defer testServer.Close()

				testService, testServiceErr := visualRecognitionV3.NewVisualRecognitionV3(&visualRecognitionV3.ServiceCredentials{
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

                result := visualRecognitionV3.GetClassifyResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DetectFaces(options *DetectFacesOptions)", func() {
		detectFacesPath := "/v3/detect_faces"
        version := "exampleString"
        detectFacesOptions := visualRecognitionV3.NewDetectFacesOptions()
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Detect faces in images", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(detectFacesPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(detectFacesPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call DetectFaces", func() {
				defer testServer.Close()

				testService, testServiceErr := visualRecognitionV3.NewVisualRecognitionV3(&visualRecognitionV3.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.DetectFaces(detectFacesOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := visualRecognitionV3.GetDetectFacesResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteClassifier(options *DeleteClassifierOptions)", func() {
		deleteClassifierPath := "/v3/classifiers/{classifier_id}"
        version := "exampleString"
        classifierID := "exampleString"
        deleteClassifierOptions := visualRecognitionV3.NewDeleteClassifierOptions(classifierID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        deleteClassifierPath = strings.Replace(deleteClassifierPath, "{classifier_id}", classifierID, 1)
		Context("Successfully - Delete a classifier", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(deleteClassifierPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(deleteClassifierPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call DeleteClassifier", func() {
				defer testServer.Close()

				testService, testServiceErr := visualRecognitionV3.NewVisualRecognitionV3(&visualRecognitionV3.ServiceCredentials{
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
		getClassifierPath := "/v3/classifiers/{classifier_id}"
        version := "exampleString"
        classifierID := "exampleString"
        getClassifierOptions := visualRecognitionV3.NewGetClassifierOptions(classifierID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        getClassifierPath = strings.Replace(getClassifierPath, "{classifier_id}", classifierID, 1)
		Context("Successfully - Retrieve classifier details", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(getClassifierPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(getClassifierPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call GetClassifier", func() {
				defer testServer.Close()

				testService, testServiceErr := visualRecognitionV3.NewVisualRecognitionV3(&visualRecognitionV3.ServiceCredentials{
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

                result := visualRecognitionV3.GetGetClassifierResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListClassifiers(options *ListClassifiersOptions)", func() {
		listClassifiersPath := "/v3/classifiers"
        version := "exampleString"
        listClassifiersOptions := visualRecognitionV3.NewListClassifiersOptions()
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Retrieve a list of classifiers", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(listClassifiersPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(listClassifiersPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call ListClassifiers", func() {
				defer testServer.Close()

				testService, testServiceErr := visualRecognitionV3.NewVisualRecognitionV3(&visualRecognitionV3.ServiceCredentials{
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

                result := visualRecognitionV3.GetListClassifiersResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("UpdateClassifier(options *UpdateClassifierOptions)", func() {
		updateClassifierPath := "/v3/classifiers/{classifier_id}"
        version := "exampleString"
        classifierID := "exampleString"
        updateClassifierOptions := visualRecognitionV3.NewUpdateClassifierOptions(classifierID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        updateClassifierPath = strings.Replace(updateClassifierPath, "{classifier_id}", classifierID, 1)
		Context("Successfully - Update a classifier", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(updateClassifierPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(updateClassifierPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call UpdateClassifier", func() {
				defer testServer.Close()

				testService, testServiceErr := visualRecognitionV3.NewVisualRecognitionV3(&visualRecognitionV3.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.UpdateClassifier(updateClassifierOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := visualRecognitionV3.GetUpdateClassifierResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetCoreMlModel(options *GetCoreMlModelOptions)", func() {
		getCoreMlModelPath := "/v3/classifiers/{classifier_id}/core_ml_model"
        version := "exampleString"
        classifierID := "exampleString"
        getCoreMlModelOptions := visualRecognitionV3.NewGetCoreMlModelOptions(classifierID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        getCoreMlModelPath = strings.Replace(getCoreMlModelPath, "{classifier_id}", classifierID, 1)
		Context("Successfully - Retrieve a Core ML model of a classifier", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(getCoreMlModelPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(getCoreMlModelPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call GetCoreMlModel", func() {
				defer testServer.Close()

				testService, testServiceErr := visualRecognitionV3.NewVisualRecognitionV3(&visualRecognitionV3.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.GetCoreMlModel(getCoreMlModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := visualRecognitionV3.GetGetCoreMlModelResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteUserData(options *DeleteUserDataOptions)", func() {
		deleteUserDataPath := "/v3/user_data"
        version := "exampleString"
        customerID := "exampleString"
        deleteUserDataOptions := visualRecognitionV3.NewDeleteUserDataOptions(customerID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Delete labeled data", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(deleteUserDataPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call DeleteUserData", func() {
				defer testServer.Close()

				testService, testServiceErr := visualRecognitionV3.NewVisualRecognitionV3(&visualRecognitionV3.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.DeleteUserData(deleteUserDataOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
})
