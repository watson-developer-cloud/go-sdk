package visualrecognitionv3_test

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/visualrecognitionv3"
)

var _ = Describe("VisualRecognitionV3", func() {
	Describe("Classify(classifyOptions *ClassifyOptions)", func() {
		ClassifyPath := "/v3/classify"
		version := "exampleString"
		Context("Successfully - Classify images", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()
				if req.URL.String() != "/v3/classify?version=exampleString" {
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"access_token":"xxxxx"}`)
				} else {
					Expect(req.URL.String()).To(Equal(ClassifyPath + "?version=" + version))
					Expect(req.URL.Path).To(Equal(ClassifyPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Authorization"]).ToNot(BeNil())
					Expect(req.Header["Authorization"][0]).To(Equal("Bearer xxxxx"))
					res.Header().Set("Content-type", "application/json")
					fmt.Fprintf(res, `{"custom_classes":"2"}`)
				}
			}))
			It("Succeed to call Classify", func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv3.
					NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
						URL:       testServer.URL,
						Version:   version,
						IAMURL:    testServer.URL,
						IAMApiKey: "xxxx",
					})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				ClassifyOptions := testService.NewClassifyOptions().SetURL("https://test.com")
				returnValue, returnValueErr := testService.Classify(ClassifyOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetClassifyResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DetectFaces(detectFacesOptions *DetectFacesOptions)", func() {
		DetectFacesPath := "/v3/detect_faces"
		version := "exampleString"
		Context("Successfully - Detect faces in images", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				if req.URL.String() != "/v3/detect_faces?version=exampleString" {
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"access_token":"xxxxx"}`)
				} else {
					Expect(req.URL.String()).To(Equal(DetectFacesPath + "?version=" + version))
					Expect(req.URL.Path).To(Equal(DetectFacesPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Authorization"]).ToNot(BeNil())
					Expect(req.Header["Authorization"][0]).To(Equal("Bearer xxxxx"))
					res.Header().Set("Content-type", "application/json")
					fmt.Fprintf(res, `{"images":[]}`)
				}
			}))
			It("Succeed to call DetectFaces", func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:       testServer.URL,
					Version:   version,
					IAMURL:    testServer.URL,
					IAMApiKey: "xxxx",
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				DetectFacesOptions := testService.NewDetectFacesOptions().
					SetURL("https://test.com")
				returnValue, returnValueErr := testService.DetectFaces(DetectFacesOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetDetectFacesResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateClassifier(createClassifierOptions *CreateClassifierOptions)", func() {
		CreateClassifierPath := "/v3/classifiers"
		version := "exampleString"
		pwd, _ := os.Getwd()
		cars, carsErr := os.Open(pwd + "/../resources/cars.zip")
		if carsErr != nil {
			panic(carsErr)
		}
		Context("Successfully - Create a classifier", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()
				if !strings.Contains(req.URL.String(), "/v3/classifiers") {
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"access_token":"xxxxx"}`)
				} else {
					Expect(req.URL.String()).To(Equal(CreateClassifierPath + "?version=" + version))
					Expect(req.URL.Path).To(Equal(CreateClassifierPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Authorization"]).ToNot(BeNil())
					Expect(req.Header["Authorization"][0]).To(Equal("Bearer xxxxx"))
					res.Header().Set("Content-type", "application/json")
					fmt.Fprintf(res, `{"classifier_id":"xxx", "name": "cars vs trucks"}`)
				}
			}))
			It("Succeed to call CreateClassifier", func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:       testServer.URL,
					Version:   version,
					IAMApiKey: "xxxxx",
					IAMURL:    testServer.URL,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				CreateClassifierOptions := testService.
					NewCreateClassifierOptions("cars vs trucks").
					AddPositiveExamples("cars", cars)
				returnValue, returnValueErr := testService.CreateClassifier(CreateClassifierOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetCreateClassifierResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteClassifier(deleteClassifierOptions *DeleteClassifierOptions)", func() {
		DeleteClassifierPath := "/v3/classifiers/{classifier_id}"
		version := "exampleString"
		ClassifierID := "exampleString"
		DeleteClassifierPath = strings.Replace(DeleteClassifierPath, "{classifier_id}", ClassifierID, 1)
		Context("Successfully - Delete a classifier", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				if !strings.Contains(req.URL.String(), "/v3/classifiers") {
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"access_token":"xxxxx"}`)
				} else {
					Expect(req.URL.String()).To(Equal(DeleteClassifierPath + "?version=" + version))
					Expect(req.URL.Path).To(Equal(DeleteClassifierPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["Authorization"]).ToNot(BeNil())
					Expect(req.Header["Authorization"][0]).To(Equal("Bearer xxxxx"))
					res.WriteHeader(200)
				}
			}))
			It("Succeed to call DeleteClassifier", func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:       testServer.URL,
					Version:   version,
					IAMApiKey: "xxxxx",
					IAMURL:    testServer.URL,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				DeleteClassifierOptions := testService.NewDeleteClassifierOptions(ClassifierID)
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
		GetClassifierPath = strings.Replace(GetClassifierPath, "{classifier_id}", ClassifierID, 1)
		Context("Successfully - Retrieve classifier details", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				if !strings.Contains(req.URL.String(), "/v3/classifiers") {
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"access_token":"xxxxx"}`)
				} else {
					Expect(req.URL.String()).To(Equal(GetClassifierPath + "?version=" + version))
					Expect(req.URL.Path).To(Equal(GetClassifierPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Authorization"]).ToNot(BeNil())
					Expect(req.Header["Authorization"][0]).To(Equal("Bearer xxxxx"))
					res.Header().Set("Content-type", "application/json")
					fmt.Fprintf(res, `{"classifier_id":"xxxx", "name": "cars vs trucks"}`)
				}
			}))
			It("Succeed to call GetClassifier", func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:       testServer.URL,
					Version:   version,
					IAMApiKey: "xxxxx",
					IAMURL:    testServer.URL,
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
		ListClassifiersPath := "/v3/classifiers"
		version := "exampleString"
		Context("Successfully - Retrieve a list of classifiers", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				if !strings.Contains(req.URL.String(), "/v3/classifiers") {
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"access_token":"xxxxx"}`)
				} else {
					Expect(req.URL.String()).To(Equal(ListClassifiersPath + "?version=" + version))
					Expect(req.URL.Path).To(Equal(ListClassifiersPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Authorization"]).ToNot(BeNil())
					Expect(req.Header["Authorization"][0]).To(Equal("Bearer xxxxx"))
					res.Header().Set("Content-type", "application/json")
					fmt.Fprintf(res, `[]`)
				}
			}))
			It("Succeed to call ListClassifiers", func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:       testServer.URL,
					Version:   version,
					IAMApiKey: "xxxxx",
					IAMURL:    testServer.URL,
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
	Describe("UpdateClassifier(updateClassifierOptions *UpdateClassifierOptions)", func() {
		UpdateClassifierPath := "/v3/classifiers/{classifier_id}"
		version := "exampleString"
		ClassifierID := "exampleString"
		pwd, _ := os.Getwd()
		trucks, trucksErr := os.Open(pwd + "/../resources/trucks.zip")
		if trucksErr != nil {
			panic(trucksErr)
		}
		UpdateClassifierPath = strings.Replace(UpdateClassifierPath, "{classifier_id}", ClassifierID, 1)
		Context("Successfully - Update a classifier", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				if !strings.Contains(req.URL.String(), "/v3/classifiers") {
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"access_token":"xxxxx"}`)
				} else {
					Expect(req.URL.String()).To(Equal(UpdateClassifierPath + "?version=" + version))
					Expect(req.URL.Path).To(Equal(UpdateClassifierPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Authorization"]).ToNot(BeNil())
					Expect(req.Header["Authorization"][0]).To(Equal("Bearer xxxxx"))
					res.Header().Set("Content-type", "application/json")
					fmt.Fprintf(res, `{"classifier_id":"xxx"}`)
				}
			}))
			It("Succeed to call UpdateClassifier", func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:       testServer.URL,
					Version:   version,
					IAMApiKey: "xxxxx",
					IAMURL:    testServer.URL,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				UpdateClassifierOptions := testService.NewUpdateClassifierOptions(ClassifierID)
				UpdateClassifierOptions.NegativeExamples = trucks
				returnValue, returnValueErr := testService.UpdateClassifier(UpdateClassifierOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetUpdateClassifierResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetCoreMlModel(getCoreMlModelOptions *GetCoreMlModelOptions)", func() {
		GetCoreMlModelPath := "/v3/classifiers/{classifier_id}/core_ml_model"
		version := "exampleString"
		ClassifierID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		GetCoreMlModelPath = strings.Replace(GetCoreMlModelPath, "{classifier_id}", ClassifierID, 1)
		Context("Successfully - Retrieve a Core ML model of a classifier", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				if req.URL.String() != "/v3/classify?version=exampleString" {
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"access_token":"xxxxx"}`)
				} else {
					Expect(req.URL.String()).To(Equal(GetCoreMlModelPath + "?version=" + version))
					Expect(req.URL.Path).To(Equal(GetCoreMlModelPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Authorization"]).ToNot(BeNil())
					Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
					res.WriteHeader(200)
					pwd, _ := os.Getwd()
					mlmodel, err := os.Open(pwd + "/../resources/CarsvsTrucks.mlmodel")
					if err != nil {
						panic(err)
					}
					bytes, err := ioutil.ReadAll(mlmodel)
					if err != nil {
						panic(err)
					}
					res.Write(bytes)
				}
			}))
			It("Succeed to call GetCoreMlModel", func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:       testServer.URL,
					Version:   version,
					IAMURL:    testServer.URL,
					IAMApiKey: "xxxx",
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				GetCoreMlModelOptions := testService.NewGetCoreMlModelOptions(ClassifierID)
				returnValue, returnValueErr := testService.GetCoreMlModel(GetCoreMlModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteUserData(deleteUserDataOptions *DeleteUserDataOptions)", func() {
		DeleteUserDataPath := "/v3/user_data"
		version := "exampleString"
		CustomerID := "exampleString"
		Context("Successfully - Delete labeled data", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				if !strings.Contains(req.URL.String(), "/v3/user_data") {
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"access_token":"xxxxx"}`)
				} else {
					Expect(req.URL.String()).To(Equal(DeleteUserDataPath + "?customer_id=" + CustomerID + "&version=" + version))
					Expect(req.URL.Path).To(Equal(DeleteUserDataPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["Authorization"]).ToNot(BeNil())
					Expect(req.Header["Authorization"][0]).To(Equal("Bearer xxxxx"))
					res.WriteHeader(200)
				}
			}))
			It("Succeed to call DeleteUserData", func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:       testServer.URL,
					Version:   version,
					IAMApiKey: "xxxxx",
					IAMURL:    testServer.URL,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				DeleteUserDataOptions := testService.NewDeleteUserDataOptions(CustomerID)
				returnValue, returnValueErr := testService.DeleteUserData(DeleteUserDataOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
})
