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

package visualrecognitionv3_test

import (
	"fmt"
	"github.com/IBM/go-sdk-core/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/visualrecognitionv3"
	"net/http"
	"net/http/httptest"
	"strings"
)

var _ = Describe(`VisualRecognitionV3`, func() {
	Describe(`Classify(classifyOptions *ClassifyOptions)`, func() {
		classifyPath := "/v3/classify"
		version := "exampleString"
		bearerToken := "0ui9876453"
		Context(`Successfully - Classify images`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(classifyPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"images": []}`)
			}))
			It(`Succeed to call Classify`, func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.Classify(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				classifyOptions := testService.NewClassifyOptions()
				result, response, operationErr = testService.Classify(classifyOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateClassifier(createClassifierOptions *CreateClassifierOptions)`, func() {
		createClassifierPath := "/v3/classifiers"
		version := "exampleString"
		bearerToken := "0ui9876453"
		name := "exampleString"
		Context(`Successfully - Create a classifier`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createClassifierPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"classifier_id": "fake_ClassifierID", "name": "fake_Name"}`)
			}))
			It(`Succeed to call CreateClassifier`, func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.CreateClassifier(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				createClassifierOptions := testService.NewCreateClassifierOptions(name)
				result, response, operationErr = testService.CreateClassifier(createClassifierOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListClassifiers(listClassifiersOptions *ListClassifiersOptions)`, func() {
		listClassifiersPath := "/v3/classifiers"
		version := "exampleString"
		bearerToken := "0ui9876453"
		Context(`Successfully - Retrieve a list of classifiers`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listClassifiersPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"classifiers": []}`)
			}))
			It(`Succeed to call ListClassifiers`, func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListClassifiers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listClassifiersOptions := testService.NewListClassifiersOptions()
				result, response, operationErr = testService.ListClassifiers(listClassifiersOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetClassifier(getClassifierOptions *GetClassifierOptions)`, func() {
		getClassifierPath := "/v3/classifiers/{classifier_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		classifierID := "exampleString"
		getClassifierPath = strings.Replace(getClassifierPath, "{classifier_id}", classifierID, 1)
		Context(`Successfully - Retrieve classifier details`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getClassifierPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"classifier_id": "fake_ClassifierID", "name": "fake_Name"}`)
			}))
			It(`Succeed to call GetClassifier`, func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetClassifier(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getClassifierOptions := testService.NewGetClassifierOptions(classifierID)
				result, response, operationErr = testService.GetClassifier(getClassifierOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateClassifier(updateClassifierOptions *UpdateClassifierOptions)`, func() {
		updateClassifierPath := "/v3/classifiers/{classifier_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		classifierID := "exampleString"
		updateClassifierPath = strings.Replace(updateClassifierPath, "{classifier_id}", classifierID, 1)
		Context(`Successfully - Update a classifier`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateClassifierPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"classifier_id": "fake_ClassifierID", "name": "fake_Name"}`)
			}))
			It(`Succeed to call UpdateClassifier`, func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.UpdateClassifier(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				updateClassifierOptions := testService.NewUpdateClassifierOptions(classifierID)
				result, response, operationErr = testService.UpdateClassifier(updateClassifierOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteClassifier(deleteClassifierOptions *DeleteClassifierOptions)`, func() {
		deleteClassifierPath := "/v3/classifiers/{classifier_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		classifierID := "exampleString"
		deleteClassifierPath = strings.Replace(deleteClassifierPath, "{classifier_id}", classifierID, 1)
		Context(`Successfully - Delete a classifier`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteClassifierPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(200)
			}))
			It(`Succeed to call DeleteClassifier`, func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.DeleteClassifier(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				deleteClassifierOptions := testService.NewDeleteClassifierOptions(classifierID)
				response, operationErr = testService.DeleteClassifier(deleteClassifierOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetCoreMlModel(getCoreMlModelOptions *GetCoreMlModelOptions)`, func() {
		getCoreMlModelPath := "/v3/classifiers/{classifier_id}/core_ml_model"
		version := "exampleString"
		bearerToken := "0ui9876453"
		classifierID := "exampleString"
		getCoreMlModelPath = strings.Replace(getCoreMlModelPath, "{classifier_id}", classifierID, 1)
		Context(`Successfully - Retrieve a Core ML model of a classifier`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getCoreMlModelPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `Contents of response byte-stream...`)
			}))
			It(`Succeed to call GetCoreMlModel`, func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetCoreMlModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getCoreMlModelOptions := testService.NewGetCoreMlModelOptions(classifierID)
				result, response, operationErr = testService.GetCoreMlModel(getCoreMlModelOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteUserData(deleteUserDataOptions *DeleteUserDataOptions)`, func() {
		deleteUserDataPath := "/v3/user_data"
		version := "exampleString"
		bearerToken := "0ui9876453"
		customerID := "exampleString"
		Context(`Successfully - Delete labeled data`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteUserDataPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				Expect(req.URL.Query()["customer_id"]).To(Equal([]string{customerID}))

				res.WriteHeader(202)
			}))
			It(`Succeed to call DeleteUserData`, func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.DeleteUserData(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				deleteUserDataOptions := testService.NewDeleteUserDataOptions(customerID)
				response, operationErr = testService.DeleteUserData(deleteUserDataOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
})
