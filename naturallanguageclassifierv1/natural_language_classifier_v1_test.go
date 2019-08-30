/**
 * (C) Copyright IBM Corp. 2019.
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
	"github.com/watson-developer-cloud/go-sdk/naturallanguageclassifierv1"
    "github.com/IBM/go-sdk-core/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"strings"
	"os"
)

var _ = Describe("NaturalLanguageClassifierV1", func() {
	Describe("Classify(classifyOptions *ClassifyOptions)", func() {
		classifyPath := "/v1/classifiers/{classifier_id}/classify"
		accessToken := "0ui9876453"
		classifierID := "exampleString"
		text := "exampleString"
		classifyPath = strings.Replace(classifyPath, "{classifier_id}", classifierID, 1)
		Context("Successfully - Classify a phrase", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(classifyPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call Classify", func() {
				defer testServer.Close()

				testService, testServiceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
					URL: testServer.URL,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.Classify(nil)
				Expect(returnValueErr).NotTo(BeNil())

				classifyOptions := testService.NewClassifyOptions(classifierID, text)
				returnValue, returnValueErr = testService.Classify(classifyOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetClassifyResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ClassifyCollection(classifyCollectionOptions *ClassifyCollectionOptions)", func() {
		classifyCollectionPath := "/v1/classifiers/{classifier_id}/classify_collection"
		accessToken := "0ui9876453"
		classifierID := "exampleString"
		collection := []naturallanguageclassifierv1.ClassifyInput{}
		classifyCollectionPath = strings.Replace(classifyCollectionPath, "{classifier_id}", classifierID, 1)
		Context("Successfully - Classify multiple phrases", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(classifyCollectionPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call ClassifyCollection", func() {
				defer testServer.Close()

				testService, testServiceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
					URL: testServer.URL,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ClassifyCollection(nil)
				Expect(returnValueErr).NotTo(BeNil())

				classifyCollectionOptions := testService.NewClassifyCollectionOptions(classifierID, collection)
				returnValue, returnValueErr = testService.ClassifyCollection(classifyCollectionOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetClassifyCollectionResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateClassifier(createClassifierOptions *CreateClassifierOptions)", func() {
		createClassifierPath := "/v1/classifiers"
		accessToken := "0ui9876453"
		pwd, _ := os.Getwd()
		Metadata, metadataErr := os.Open(pwd + "/../resources/weather_training_metadata.json")
		if metadataErr != nil {
			fmt.Println(metadataErr)
		}
		data, dataErr := os.Open(pwd + "/../resources/weather_training_data.csv")
		if dataErr != nil {
			fmt.Println(dataErr)
		}
		Context("Successfully - Create classifier", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createClassifierPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"url": "fake URL", "classifier_id": "fake ClassifierID"}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call CreateClassifier", func() {
				defer testServer.Close()

				testService, testServiceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
					URL: testServer.URL,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.CreateClassifier(nil)
				Expect(returnValueErr).NotTo(BeNil())

				createClassifierOptions := testService.NewCreateClassifierOptions(Metadata, data)
				returnValue, returnValueErr = testService.CreateClassifier(createClassifierOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetCreateClassifierResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListClassifiers(listClassifiersOptions *ListClassifiersOptions)", func() {
		listClassifiersPath := "/v1/classifiers"
		accessToken := "0ui9876453"
		Context("Successfully - List classifiers", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listClassifiersPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"classifiers": []}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call ListClassifiers", func() {
				defer testServer.Close()

				testService, testServiceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
					URL: testServer.URL,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ListClassifiers(nil)
				Expect(returnValueErr).NotTo(BeNil())

				listClassifiersOptions := testService.NewListClassifiersOptions()
				returnValue, returnValueErr = testService.ListClassifiers(listClassifiersOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListClassifiersResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetClassifier(getClassifierOptions *GetClassifierOptions)", func() {
		getClassifierPath := "/v1/classifiers/{classifier_id}"
		accessToken := "0ui9876453"
		classifierID := "exampleString"
		getClassifierPath = strings.Replace(getClassifierPath, "{classifier_id}", classifierID, 1)
		Context("Successfully - Get information about a classifier", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getClassifierPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"url": "fake URL", "classifier_id": "fake ClassifierID"}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call GetClassifier", func() {
				defer testServer.Close()

				testService, testServiceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
					URL: testServer.URL,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.GetClassifier(nil)
				Expect(returnValueErr).NotTo(BeNil())

				getClassifierOptions := testService.NewGetClassifierOptions(classifierID)
				returnValue, returnValueErr = testService.GetClassifier(getClassifierOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetClassifierResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteClassifier(deleteClassifierOptions *DeleteClassifierOptions)", func() {
		deleteClassifierPath := "/v1/classifiers/{classifier_id}"
		accessToken := "0ui9876453"
		classifierID := "exampleString"
		deleteClassifierPath = strings.Replace(deleteClassifierPath, "{classifier_id}", classifierID, 1)
		Context("Successfully - Delete classifier", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteClassifierPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.WriteHeader(200)
			}))
			It("Succeed to call DeleteClassifier", func() {
				defer testServer.Close()

				testService, testServiceErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
					URL: testServer.URL,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.DeleteClassifier(nil)
				Expect(returnValueErr).NotTo(BeNil())

				deleteClassifierOptions := testService.NewDeleteClassifierOptions(classifierID)
				returnValue, returnValueErr = testService.DeleteClassifier(deleteClassifierOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
})
