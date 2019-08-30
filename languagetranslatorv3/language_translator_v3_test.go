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

package languagetranslatorv3_test

import (
	"fmt"
	"github.com/watson-developer-cloud/go-sdk/languagetranslatorv3"
    "github.com/IBM/go-sdk-core/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"strings"
	"os"
)

var _ = Describe("LanguageTranslatorV3", func() {
	Describe("Translate(translateOptions *TranslateOptions)", func() {
		translatePath := "/v3/translate"
		version := "exampleString"
		accessToken := "0ui9876453"
		text := []string{}
		Context("Successfully - Translate", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(translatePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"word_count": 9, "character_count": 14, "translations": []}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call Translate", func() {
				defer testServer.Close()

				testService, testServiceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.Translate(nil)
				Expect(returnValueErr).NotTo(BeNil())

				translateOptions := testService.NewTranslateOptions(text)
				returnValue, returnValueErr = testService.Translate(translateOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetTranslateResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListIdentifiableLanguages(listIdentifiableLanguagesOptions *ListIdentifiableLanguagesOptions)", func() {
		listIdentifiableLanguagesPath := "/v3/identifiable_languages"
		version := "exampleString"
		accessToken := "0ui9876453"
		Context("Successfully - List identifiable languages", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listIdentifiableLanguagesPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"languages": []}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call ListIdentifiableLanguages", func() {
				defer testServer.Close()

				testService, testServiceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ListIdentifiableLanguages(nil)
				Expect(returnValueErr).NotTo(BeNil())

				listIdentifiableLanguagesOptions := testService.NewListIdentifiableLanguagesOptions()
				returnValue, returnValueErr = testService.ListIdentifiableLanguages(listIdentifiableLanguagesOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListIdentifiableLanguagesResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("Identify(identifyOptions *IdentifyOptions)", func() {
		identifyPath := "/v3/identify"
		version := "exampleString"
		accessToken := "0ui9876453"
		text := "exampleString"
		Context("Successfully - Identify language", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(identifyPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"languages": []}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call Identify", func() {
				defer testServer.Close()

				testService, testServiceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.Identify(nil)
				Expect(returnValueErr).NotTo(BeNil())

				identifyOptions := testService.NewIdentifyOptions(text)
				returnValue, returnValueErr = testService.Identify(identifyOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetIdentifyResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListModels(listModelsOptions *ListModelsOptions)", func() {
		listModelsPath := "/v3/models"
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
				fmt.Fprintf(res, `{"models": []}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call ListModels", func() {
				defer testServer.Close()

				testService, testServiceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
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
	Describe("CreateModel(createModelOptions *CreateModelOptions)", func() {
		createModelPath := "/v3/models"
		version := "exampleString"
		accessToken := "0ui9876453"
		baseModelID := "exampleString"
		Context("Successfully - Create model", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createModelPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				Expect(req.URL.Query()["base_model_id"]).To(Equal([]string{baseModelID}))

				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"model_id": "fake ModelID"}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call CreateModel", func() {
				defer testServer.Close()

				testService, testServiceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.CreateModel(nil)
				Expect(returnValueErr).NotTo(BeNil())

				pwd, _ := os.Getwd()
				file, err := os.Open(pwd + "/../resources/language_translator_model.tmx")
				Expect(err).To(BeNil())
				defer file.Close()

				createModelOptions := testService.NewCreateModelOptions(baseModelID).
					SetForcedGlossary(file)
				returnValue, returnValueErr = testService.CreateModel(createModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetCreateModelResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteModel(deleteModelOptions *DeleteModelOptions)", func() {
		deleteModelPath := "/v3/models/{model_id}"
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
				fmt.Fprintf(res, `{"status": "fake Status"}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call DeleteModel", func() {
				defer testServer.Close()

				testService, testServiceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
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
	Describe("GetModel(getModelOptions *GetModelOptions)", func() {
		getModelPath := "/v3/models/{model_id}"
		version := "exampleString"
		accessToken := "0ui9876453"
		modelID := "exampleString"
		getModelPath = strings.Replace(getModelPath, "{model_id}", modelID, 1)
		Context("Successfully - Get model details", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getModelPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"model_id": "fake ModelID"}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call GetModel", func() {
				defer testServer.Close()

				testService, testServiceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.GetModel(nil)
				Expect(returnValueErr).NotTo(BeNil())

				getModelOptions := testService.NewGetModelOptions(modelID)
				returnValue, returnValueErr = testService.GetModel(getModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetModelResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListDocuments(listDocumentsOptions *ListDocumentsOptions)", func() {
		listDocumentsPath := "/v3/documents"
		version := "exampleString"
		accessToken := "0ui9876453"
		Context("Successfully - List documents", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listDocumentsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"documents": []}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call ListDocuments", func() {
				defer testServer.Close()

				testService, testServiceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ListDocuments(nil)
				Expect(returnValueErr).NotTo(BeNil())

				listDocumentsOptions := testService.NewListDocumentsOptions()
				returnValue, returnValueErr = testService.ListDocuments(listDocumentsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListDocumentsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("TranslateDocument(translateDocumentOptions *TranslateDocumentOptions)", func() {
		translateDocumentPath := "/v3/documents"
		version := "exampleString"
		accessToken := "0ui9876453"
		filename := "exampleString"
		Context("Successfully - Translate document", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(translateDocumentPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"document_id": "fake DocumentID", "filename": "fake Filename", "status": "fake Status", "model_id": "fake ModelID", "source": "fake Source", "target": "fake Target", "created": "2017-05-16T13:56:54.957Z"}`)
				res.WriteHeader(202)
			}))
			It("Succeed to call TranslateDocument", func() {
				defer testServer.Close()

				testService, testServiceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.TranslateDocument(nil)
				Expect(returnValueErr).NotTo(BeNil())

				pwd, _ := os.Getwd()
				file, err := os.Open(pwd + "/../resources/hello_world.txt")
				Expect(err).To(BeNil())
				defer file.Close()

				translateDocumentOptions := testService.NewTranslateDocumentOptions(file, filename)
				returnValue, returnValueErr = testService.TranslateDocument(translateDocumentOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetTranslateDocumentResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetDocumentStatus(getDocumentStatusOptions *GetDocumentStatusOptions)", func() {
		getDocumentStatusPath := "/v3/documents/{document_id}"
		version := "exampleString"
		accessToken := "0ui9876453"
		documentID := "exampleString"
		getDocumentStatusPath = strings.Replace(getDocumentStatusPath, "{document_id}", documentID, 1)
		Context("Successfully - Get document status", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getDocumentStatusPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"document_id": "fake DocumentID", "filename": "fake Filename", "status": "fake Status", "model_id": "fake ModelID", "source": "fake Source", "target": "fake Target", "created": "2017-05-16T13:56:54.957Z"}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call GetDocumentStatus", func() {
				defer testServer.Close()

				testService, testServiceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.GetDocumentStatus(nil)
				Expect(returnValueErr).NotTo(BeNil())

				getDocumentStatusOptions := testService.NewGetDocumentStatusOptions(documentID)
				returnValue, returnValueErr = testService.GetDocumentStatus(getDocumentStatusOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetDocumentStatusResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteDocument(deleteDocumentOptions *DeleteDocumentOptions)", func() {
		deleteDocumentPath := "/v3/documents/{document_id}"
		version := "exampleString"
		accessToken := "0ui9876453"
		documentID := "exampleString"
		deleteDocumentPath = strings.Replace(deleteDocumentPath, "{document_id}", documentID, 1)
		Context("Successfully - Delete document", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteDocumentPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.WriteHeader(204)
			}))
			It("Succeed to call DeleteDocument", func() {
				defer testServer.Close()

				testService, testServiceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.DeleteDocument(nil)
				Expect(returnValueErr).NotTo(BeNil())

				deleteDocumentOptions := testService.NewDeleteDocumentOptions(documentID)
				returnValue, returnValueErr = testService.DeleteDocument(deleteDocumentOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("GetTranslatedDocument(getTranslatedDocumentOptions *GetTranslatedDocumentOptions)", func() {
		getTranslatedDocumentPath := "/v3/documents/{document_id}/translated_document"
		version := "exampleString"
		accessToken := "0ui9876453"
		documentID := "exampleString"
		getTranslatedDocumentPath = strings.Replace(getTranslatedDocumentPath, "{document_id}", documentID, 1)
		Context("Successfully - Get translated document", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getTranslatedDocumentPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.WriteHeader(200)
			}))
			It("Succeed to call GetTranslatedDocument", func() {
				defer testServer.Close()

				testService, testServiceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.GetTranslatedDocument(nil)
				Expect(returnValueErr).NotTo(BeNil())

				getTranslatedDocumentOptions := testService.NewGetTranslatedDocumentOptions(documentID)
				returnValue, returnValueErr = testService.GetTranslatedDocument(getTranslatedDocumentOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
})
