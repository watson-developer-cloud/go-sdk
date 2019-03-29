package languagetranslatorv3_test

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	languagetranslatorv3 "github.com/watson-developer-cloud/go-sdk/languagetranslatorv3"
)

var _ = Describe("LanguageTranslatorV3", func() {
	Describe("Translate(options *TranslateOptions)", func() {
		translatePath := "/v3/translate"
		version := "exampleString"
		text := []string{}
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Translate", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(translatePath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(translatePath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"word_count":28}`)
			}))
			It("Succeed to call Translate", func() {
				defer testServer.Close()

				testService, testServiceErr := languagetranslatorv3.
					NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
						URL:      testServer.URL,
						Version:  version,
						Username: username,
						Password: password,
					})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				translateOptions := testService.NewTranslateOptions(text)
				returnValue, returnValueErr := testService.Translate(translateOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetTranslateResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("Identify(options *IdentifyOptions)", func() {
		identifyPath := "/v3/identify"
		version := "exampleString"
		text := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Identify language", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(identifyPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(identifyPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `[]`)
			}))
			It("Succeed to call Identify", func() {
				defer testServer.Close()

				testService, testServiceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				identifyOptions := testService.NewIdentifyOptions(text)
				returnValue, returnValueErr := testService.Identify(identifyOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetIdentifyResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListIdentifiableLanguages(options *ListIdentifiableLanguagesOptions)", func() {
		listIdentifiableLanguagesPath := "/v3/identifiable_languages"
		version := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - List identifiable languages", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(listIdentifiableLanguagesPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(listIdentifiableLanguagesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `[{"language":"en", "name":"english"}]`)
			}))
			It("Succeed to call ListIdentifiableLanguages", func() {
				defer testServer.Close()

				testService, testServiceErr := languagetranslatorv3.
					NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
						URL:      testServer.URL,
						Version:  version,
						Username: username,
						Password: password,
					})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				listIdentifiableLanguagesOptions := testService.NewListIdentifiableLanguagesOptions()
				returnValue, returnValueErr := testService.ListIdentifiableLanguages(listIdentifiableLanguagesOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListIdentifiableLanguagesResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateModel(options *CreateModelOptions)", func() {
		createModelPath := "/v3/models"
		version := "exampleString"
		baseModelID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Create model", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(createModelPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"model_id":"xxx"}`)
			}))
			It("Succeed to call CreateModel", func() {
				defer testServer.Close()

				testService, testServiceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				pwd, _ := os.Getwd()
				file, err := os.Open(pwd + "/../resources/language_translator_model.tmx")
				Expect(err).To(BeNil())
				defer file.Close()

				createModelOptions := testService.NewCreateModelOptions(baseModelID).
					SetForcedGlossary(file)
				returnValue, returnValueErr := testService.CreateModel(createModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetCreateModelResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteModel(options *DeleteModelOptions)", func() {
		deleteModelPath := "/v3/models/{model_id}"
		version := "exampleString"
		modelID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		deleteModelPath = strings.Replace(deleteModelPath, "{model_id}", modelID, 1)
		Context("Successfully - Delete model", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(deleteModelPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(deleteModelPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"status":"success"}`)
			}))
			It("Succeed to call DeleteModel", func() {
				defer testServer.Close()

				testService, testServiceErr := languagetranslatorv3.
					NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
						URL:      testServer.URL,
						Version:  version,
						Username: username,
						Password: password,
					})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				deleteModelOptions := testService.NewDeleteModelOptions(modelID)
				returnValue, returnValueErr := testService.DeleteModel(deleteModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetDeleteModelResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetModel(options *GetModelOptions)", func() {
		getModelPath := "/v3/models/{model_id}"
		version := "exampleString"
		modelID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		getModelPath = strings.Replace(getModelPath, "{model_id}", modelID, 1)
		Context("Successfully - Get model details", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(getModelPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(getModelPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"model_id":"xxx"}`)
			}))
			It("Succeed to call GetModel", func() {
				defer testServer.Close()

				testService, testServiceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				getModelOptions := testService.NewGetModelOptions(modelID)
				returnValue, returnValueErr := testService.GetModel(getModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetModelResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListModels(options *ListModelsOptions)", func() {
		listModelsPath := "/v3/models"
		version := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - List models", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(listModelsPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(listModelsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `[]`)
			}))
			It("Succeed to call ListModels", func() {
				defer testServer.Close()

				testService, testServiceErr := languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				listModelsOptions := testService.NewListModelsOptions()
				returnValue, returnValueErr := testService.ListModels(listModelsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListModelsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
})
