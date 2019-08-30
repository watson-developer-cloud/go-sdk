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

package naturallanguageunderstandingv1_test

import (
	"fmt"
	"github.com/watson-developer-cloud/go-sdk/naturallanguageunderstandingv1"
    "github.com/IBM/go-sdk-core/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"strings"
)

var _ = Describe("NaturalLanguageUnderstandingV1", func() {
	Describe("Analyze(analyzeOptions *AnalyzeOptions)", func() {
		analyzePath := "/v1/analyze"
		version := "exampleString"
		accessToken := "0ui9876453"
		features := new(naturallanguageunderstandingv1.Features)
		Context("Successfully - Analyze text", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(analyzePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call Analyze", func() {
				defer testServer.Close()

				testService, testServiceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.Analyze(nil)
				Expect(returnValueErr).NotTo(BeNil())

				analyzeOptions := testService.NewAnalyzeOptions(features)
				returnValue, returnValueErr = testService.Analyze(analyzeOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetAnalyzeResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListModels(listModelsOptions *ListModelsOptions)", func() {
		listModelsPath := "/v1/models"
		version := "exampleString"
		accessToken := "0ui9876453"
		Context("Successfully - List models", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listModelsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call ListModels", func() {
				defer testServer.Close()

				testService, testServiceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ListModels(nil)
				Expect(returnValueErr).NotTo(BeNil())

				listModelsOptions := testService.NewListModelsOptions()
				returnValue, returnValueErr = testService.ListModels(listModelsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListModelsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteModel(deleteModelOptions *DeleteModelOptions)", func() {
		deleteModelPath := "/v1/models/{model_id}"
		version := "exampleString"
		accessToken := "0ui9876453"
		modelID := "exampleString"
		deleteModelPath = strings.Replace(deleteModelPath, "{model_id}", modelID, 1)
		Context("Successfully - Delete model", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteModelPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call DeleteModel", func() {
				defer testServer.Close()

				testService, testServiceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.DeleteModel(nil)
				Expect(returnValueErr).NotTo(BeNil())

				deleteModelOptions := testService.NewDeleteModelOptions(modelID)
				returnValue, returnValueErr = testService.DeleteModel(deleteModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetDeleteModelResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
})
