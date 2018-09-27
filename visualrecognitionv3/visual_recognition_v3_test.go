package visualrecognitionv3_test

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/ibm-watson/go-sdk/visualrecognitionv3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
)

var _ = Describe("VisualRecognitionV3", func() {
	Describe("Get credentials from VCAP", func() {
		version := "exampleString"
		username := "hyphenated-user"
		password := "hyphenated-pass"
		VCAPservices := cfenv.Services{
			"conversation": {
				{
					Name: "watson_vision_combined",
					Tags: []string{},
					Credentials: map[string]interface{}{
						"url":      "https://gateway.watsonplatform.net/visual-recognition/api",
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

				testService, testServiceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				testService.ListClassifiers(testService.NewListClassifiersOptions())
			})
		})
	})
	Describe("Classify(classifyOptions *ClassifyOptions)", func() {
		ClassifyPath := "/v3/classify"
		version := "exampleString"
		ClassifyOptions := testService.NewClassifyOptions()
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Classify images", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(ClassifyPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(ClassifyPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call Classify", func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.Classify(ClassifyOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := visualrecognitionv3.GetClassifyResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DetectFaces(detectFacesOptions *DetectFacesOptions)", func() {
		DetectFacesPath := "/v3/detect_faces"
		version := "exampleString"
		DetectFacesOptions := testService.NewDetectFacesOptions()
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Detect faces in images", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(DetectFacesPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(DetectFacesPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call DetectFaces", func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.DetectFaces(DetectFacesOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := visualrecognitionv3.GetDetectFacesResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateClassifier(createClassifierOptions *CreateClassifierOptions)", func() {
		CreateClassifierPath := "/v3/classifiers"
		version := "exampleString"
		Name := "exampleString"
		ClassnamePositiveExamples := new(os.File)
		CreateClassifierOptions := testService.NewCreateClassifierOptions(Name, *ClassnamePositiveExamples)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Create a classifier", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(CreateClassifierPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(CreateClassifierPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call CreateClassifier", func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.CreateClassifier(CreateClassifierOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := visualrecognitionv3.GetCreateClassifierResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteClassifier(deleteClassifierOptions *DeleteClassifierOptions)", func() {
		DeleteClassifierPath := "/v3/classifiers/{classifier_id}"
		version := "exampleString"
		ClassifierID := "exampleString"
		DeleteClassifierOptions := testService.NewDeleteClassifierOptions(ClassifierID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		DeleteClassifierPath = strings.Replace(DeleteClassifierPath, "{classifier_id}", ClassifierID, 1)
		Context("Successfully - Delete a classifier", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(DeleteClassifierPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(DeleteClassifierPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call DeleteClassifier", func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.DeleteClassifier(DeleteClassifierOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("GetClassifier(getClassifierOptions *GetClassifierOptions)", func() {
		GetClassifierPath := "/v3/classifiers/{classifier_id}"
		version := "exampleString"
		ClassifierID := "exampleString"
		GetClassifierOptions := testService.NewGetClassifierOptions(ClassifierID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		GetClassifierPath = strings.Replace(GetClassifierPath, "{classifier_id}", ClassifierID, 1)
		Context("Successfully - Retrieve classifier details", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(GetClassifierPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(GetClassifierPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call GetClassifier", func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.GetClassifier(GetClassifierOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := visualrecognitionv3.GetGetClassifierResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListClassifiers(listClassifiersOptions *ListClassifiersOptions)", func() {
		ListClassifiersPath := "/v3/classifiers"
		version := "exampleString"
		ListClassifiersOptions := testService.NewListClassifiersOptions()
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Retrieve a list of classifiers", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(ListClassifiersPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(ListClassifiersPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call ListClassifiers", func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.ListClassifiers(ListClassifiersOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := visualrecognitionv3.GetListClassifiersResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("UpdateClassifier(updateClassifierOptions *UpdateClassifierOptions)", func() {
		UpdateClassifierPath := "/v3/classifiers/{classifier_id}"
		version := "exampleString"
		ClassifierID := "exampleString"
		UpdateClassifierOptions := testService.NewUpdateClassifierOptions(ClassifierID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		UpdateClassifierPath = strings.Replace(UpdateClassifierPath, "{classifier_id}", ClassifierID, 1)
		Context("Successfully - Update a classifier", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(UpdateClassifierPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(UpdateClassifierPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call UpdateClassifier", func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.UpdateClassifier(UpdateClassifierOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := visualrecognitionv3.GetUpdateClassifierResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetCoreMlModel(getCoreMlModelOptions *GetCoreMlModelOptions)", func() {
		GetCoreMlModelPath := "/v3/classifiers/{classifier_id}/core_ml_model"
		version := "exampleString"
		ClassifierID := "exampleString"
		GetCoreMlModelOptions := testService.NewGetCoreMlModelOptions(ClassifierID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		GetCoreMlModelPath = strings.Replace(GetCoreMlModelPath, "{classifier_id}", ClassifierID, 1)
		Context("Successfully - Retrieve a Core ML model of a classifier", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(GetCoreMlModelPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(GetCoreMlModelPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call GetCoreMlModel", func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.GetCoreMlModel(GetCoreMlModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := visualrecognitionv3.GetGetCoreMlModelResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteUserData(deleteUserDataOptions *DeleteUserDataOptions)", func() {
		DeleteUserDataPath := "/v3/user_data"
		version := "exampleString"
		CustomerID := "exampleString"
		DeleteUserDataOptions := testService.NewDeleteUserDataOptions(CustomerID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Delete labeled data", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(DeleteUserDataPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(DeleteUserDataPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call DeleteUserData", func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.DeleteUserData(DeleteUserDataOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
})
