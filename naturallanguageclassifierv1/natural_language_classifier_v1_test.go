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

package naturallanguageclassifierv1_test

import (
	"fmt"
	"github.com/IBM/go-sdk-core/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/naturallanguageclassifierv1"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
)

var _ = Describe(`NaturalLanguageClassifierV1`, func() {
	Describe(`Classify(classifyOptions *ClassifyOptions)`, func() {
		classifyPath := "/v1/classifiers/{classifier_id}/classify"
		bearerToken := "0ui9876453"
		classifierID := "exampleString"
		text := "exampleString"
		classifyPath = strings.Replace(classifyPath, "{classifier_id}", classifierID, 1)
		Context(`Successfully - Classify a phrase`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(classifyPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call Classify`, func() {
				defer testServer.Close()

				testService, testServiceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
					URL: testServer.URL,
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

				classifyOptions := testService.NewClassifyOptions(classifierID, text)
				result, response, operationErr = testService.Classify(classifyOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ClassifyCollection(classifyCollectionOptions *ClassifyCollectionOptions)`, func() {
		classifyCollectionPath := "/v1/classifiers/{classifier_id}/classify_collection"
		bearerToken := "0ui9876453"
		classifierID := "exampleString"
		collection := []naturallanguageclassifierv1.ClassifyInput{}
		classifyCollectionPath = strings.Replace(classifyCollectionPath, "{classifier_id}", classifierID, 1)
		Context(`Successfully - Classify multiple phrases`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(classifyCollectionPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call ClassifyCollection`, func() {
				defer testServer.Close()

				testService, testServiceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ClassifyCollection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				classifyCollectionOptions := testService.NewClassifyCollectionOptions(classifierID, collection)
				result, response, operationErr = testService.ClassifyCollection(classifyCollectionOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateClassifier(createClassifierOptions *CreateClassifierOptions)`, func() {
		createClassifierPath := "/v1/classifiers"
		bearerToken := "0ui9876453"
		Context(`Successfully - Create classifier`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createClassifierPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"url": "fake_URL", "classifier_id": "fake_ClassifierID"}`)
			}))
			It(`Succeed to call CreateClassifier`, func() {
				defer testServer.Close()

				testService, testServiceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
					URL: testServer.URL,
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

				pwd, _ := os.Getwd()
				metadata, metadataErr := os.Open(pwd + "/../resources/weather_training_metadata.json")
				if metadataErr != nil {
					fmt.Println(metadataErr)
				}
				data, dataErr := os.Open(pwd + "/../resources/weather_training_data.csv")
				if dataErr != nil {
					fmt.Println(dataErr)
				}
				createClassifierOptions := testService.NewCreateClassifierOptions(metadata, data)
				result, response, operationErr = testService.CreateClassifier(createClassifierOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListClassifiers(listClassifiersOptions *ListClassifiersOptions)`, func() {
		listClassifiersPath := "/v1/classifiers"
		bearerToken := "0ui9876453"
		Context(`Successfully - List classifiers`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listClassifiersPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"classifiers": []}`)
			}))
			It(`Succeed to call ListClassifiers`, func() {
				defer testServer.Close()

				testService, testServiceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
					URL: testServer.URL,
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
		getClassifierPath := "/v1/classifiers/{classifier_id}"
		bearerToken := "0ui9876453"
		classifierID := "exampleString"
		getClassifierPath = strings.Replace(getClassifierPath, "{classifier_id}", classifierID, 1)
		Context(`Successfully - Get information about a classifier`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getClassifierPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"url": "fake_URL", "classifier_id": "fake_ClassifierID"}`)
			}))
			It(`Succeed to call GetClassifier`, func() {
				defer testServer.Close()

				testService, testServiceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
					URL: testServer.URL,
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
	Describe(`DeleteClassifier(deleteClassifierOptions *DeleteClassifierOptions)`, func() {
		deleteClassifierPath := "/v1/classifiers/{classifier_id}"
		bearerToken := "0ui9876453"
		classifierID := "exampleString"
		deleteClassifierPath = strings.Replace(deleteClassifierPath, "{classifier_id}", classifierID, 1)
		Context(`Successfully - Delete classifier`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteClassifierPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(200)
			}))
			It(`Succeed to call DeleteClassifier`, func() {
				defer testServer.Close()

				testService, testServiceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
					URL: testServer.URL,
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
	Describe("Model constructor tests", func() {
		Context("with a sample service", func() {
			testService, _ := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
				URL:           "http://naturallanguageclassifierv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It("should call NewClassifyInput successfully", func() {
				text := "exampleString"
				model, err := testService.NewClassifyInput(text)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
})
