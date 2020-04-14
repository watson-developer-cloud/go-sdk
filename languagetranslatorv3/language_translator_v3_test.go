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

package languagetranslatorv3_test

import (
	"fmt"
	"github.com/IBM/go-sdk-core/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/languagetranslatorv3"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
)

var _ = Describe(`LanguageTranslatorV3`, func() {
	Describe(`Translate(translateOptions *TranslateOptions)`, func() {
		translatePath := "/v3/translate"
		version := "exampleString"
		bearerToken := "0ui9876453"
		text := []string{}
		Context(`Successfully - Translate`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(translatePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"word_count": 9, "character_count": 14, "translations": []}`)
			}))
			It(`Succeed to call Translate`, func() {
				defer testServer.Close()

				testService, testServiceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.Translate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				translateOptions := testService.NewTranslateOptions(text)
				result, response, operationErr = testService.Translate(translateOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListIdentifiableLanguages(listIdentifiableLanguagesOptions *ListIdentifiableLanguagesOptions)`, func() {
		listIdentifiableLanguagesPath := "/v3/identifiable_languages"
		version := "exampleString"
		bearerToken := "0ui9876453"
		Context(`Successfully - List identifiable languages`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listIdentifiableLanguagesPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"languages": []}`)
			}))
			It(`Succeed to call ListIdentifiableLanguages`, func() {
				defer testServer.Close()

				testService, testServiceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListIdentifiableLanguages(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listIdentifiableLanguagesOptions := testService.NewListIdentifiableLanguagesOptions()
				result, response, operationErr = testService.ListIdentifiableLanguages(listIdentifiableLanguagesOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`Identify(identifyOptions *IdentifyOptions)`, func() {
		identifyPath := "/v3/identify"
		version := "exampleString"
		bearerToken := "0ui9876453"
		text := "exampleString"
		Context(`Successfully - Identify language`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(identifyPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"languages": []}`)
			}))
			It(`Succeed to call Identify`, func() {
				defer testServer.Close()

				testService, testServiceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.Identify(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				identifyOptions := testService.NewIdentifyOptions(text)
				result, response, operationErr = testService.Identify(identifyOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListModels(listModelsOptions *ListModelsOptions)`, func() {
		listModelsPath := "/v3/models"
		version := "exampleString"
		bearerToken := "0ui9876453"
		Context(`Successfully - List models`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listModelsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"models": []}`)
			}))
			It(`Succeed to call ListModels`, func() {
				defer testServer.Close()

				testService, testServiceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListModels(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listModelsOptions := testService.NewListModelsOptions()
				result, response, operationErr = testService.ListModels(listModelsOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateModel(createModelOptions *CreateModelOptions)`, func() {
		createModelPath := "/v3/models"
		version := "exampleString"
		bearerToken := "0ui9876453"
		baseModelID := "exampleString"
		Context(`Successfully - Create model`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createModelPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				Expect(req.URL.Query()["base_model_id"]).To(Equal([]string{baseModelID}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"model_id": "fake_ModelID"}`)
			}))
			It(`Succeed to call CreateModel`, func() {
				defer testServer.Close()

				testService, testServiceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.CreateModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				file, err := os.Open("../resources/language_translator_model.tmx")
				Expect(err).To(BeNil())
				defer file.Close()

				createModelOptions := testService.NewCreateModelOptions(baseModelID).
					SetForcedGlossary(file)
				result, response, operationErr = testService.CreateModel(createModelOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteModel(deleteModelOptions *DeleteModelOptions)`, func() {
		deleteModelPath := "/v3/models/{model_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		modelID := "exampleString"
		deleteModelPath = strings.Replace(deleteModelPath, "{model_id}", modelID, 1)
		Context(`Successfully - Delete model`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteModelPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"status": "fake_Status"}`)
			}))
			It(`Succeed to call DeleteModel`, func() {
				defer testServer.Close()

				testService, testServiceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.DeleteModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				deleteModelOptions := testService.NewDeleteModelOptions(modelID)
				result, response, operationErr = testService.DeleteModel(deleteModelOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetModel(getModelOptions *GetModelOptions)`, func() {
		getModelPath := "/v3/models/{model_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		modelID := "exampleString"
		getModelPath = strings.Replace(getModelPath, "{model_id}", modelID, 1)
		Context(`Successfully - Get model details`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getModelPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"model_id": "fake_ModelID"}`)
			}))
			It(`Succeed to call GetModel`, func() {
				defer testServer.Close()

				testService, testServiceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getModelOptions := testService.NewGetModelOptions(modelID)
				result, response, operationErr = testService.GetModel(getModelOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListDocuments(listDocumentsOptions *ListDocumentsOptions)`, func() {
		listDocumentsPath := "/v3/documents"
		version := "exampleString"
		bearerToken := "0ui9876453"
		Context(`Successfully - List documents`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listDocumentsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"documents": []}`)
			}))
			It(`Succeed to call ListDocuments`, func() {
				defer testServer.Close()

				testService, testServiceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListDocuments(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listDocumentsOptions := testService.NewListDocumentsOptions()
				result, response, operationErr = testService.ListDocuments(listDocumentsOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`TranslateDocument(translateDocumentOptions *TranslateDocumentOptions)`, func() {
		translateDocumentPath := "/v3/documents"
		version := "exampleString"
		bearerToken := "0ui9876453"
		filename := "exampleString"
		Context(`Successfully - Translate document`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(translateDocumentPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(202)
				fmt.Fprintf(res, `{"document_id": "fake_DocumentID", "filename": "fake_Filename", "status": "fake_Status", "model_id": "fake_ModelID", "source": "fake_Source", "target": "fake_Target", "created": "2017-05-16T13:56:54.957Z"}`)
			}))
			It(`Succeed to call TranslateDocument`, func() {
				defer testServer.Close()

				testService, testServiceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.TranslateDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				file, err := os.Open("../resources/hello_world.txt")
				Expect(err).To(BeNil())
				defer file.Close()

				translateDocumentOptions := testService.NewTranslateDocumentOptions(file, filename)
				result, response, operationErr = testService.TranslateDocument(translateDocumentOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetDocumentStatus(getDocumentStatusOptions *GetDocumentStatusOptions)`, func() {
		getDocumentStatusPath := "/v3/documents/{document_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		documentID := "exampleString"
		getDocumentStatusPath = strings.Replace(getDocumentStatusPath, "{document_id}", documentID, 1)
		Context(`Successfully - Get document status`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getDocumentStatusPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"document_id": "fake_DocumentID", "filename": "fake_Filename", "status": "fake_Status", "model_id": "fake_ModelID", "source": "fake_Source", "target": "fake_Target", "created": "2017-05-16T13:56:54.957Z"}`)
			}))
			It(`Succeed to call GetDocumentStatus`, func() {
				defer testServer.Close()

				testService, testServiceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetDocumentStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getDocumentStatusOptions := testService.NewGetDocumentStatusOptions(documentID)
				result, response, operationErr = testService.GetDocumentStatus(getDocumentStatusOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteDocument(deleteDocumentOptions *DeleteDocumentOptions)`, func() {
		deleteDocumentPath := "/v3/documents/{document_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		documentID := "exampleString"
		deleteDocumentPath = strings.Replace(deleteDocumentPath, "{document_id}", documentID, 1)
		Context(`Successfully - Delete document`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteDocumentPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(204)
			}))
			It(`Succeed to call DeleteDocument`, func() {
				defer testServer.Close()

				testService, testServiceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.DeleteDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				deleteDocumentOptions := testService.NewDeleteDocumentOptions(documentID)
				response, operationErr = testService.DeleteDocument(deleteDocumentOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetTranslatedDocument(getTranslatedDocumentOptions *GetTranslatedDocumentOptions)`, func() {
		getTranslatedDocumentPath := "/v3/documents/{document_id}/translated_document"
		version := "exampleString"
		bearerToken := "0ui9876453"
		documentID := "exampleString"
		getTranslatedDocumentPath = strings.Replace(getTranslatedDocumentPath, "{document_id}", documentID, 1)
		Context(`Successfully - Get translated document`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getTranslatedDocumentPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `Contents of response byte-stream...`)
			}))
			It(`Succeed to call GetTranslatedDocument`, func() {
				defer testServer.Close()

				testService, testServiceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetTranslatedDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getTranslatedDocumentOptions := testService.NewGetTranslatedDocumentOptions(documentID)
				result, response, operationErr = testService.GetTranslatedDocument(getTranslatedDocumentOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
})
