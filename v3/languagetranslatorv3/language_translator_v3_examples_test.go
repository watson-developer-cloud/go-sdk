//go:build examples
// +build examples

/**
 * (C) Copyright IBM Corp. 2022.
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
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/v3/languagetranslatorv3"
)

//
// This file provides an example of how to use the Language Translator service.
//
// The following configuration properties are assumed to be defined:
// LANGUAGE_TRANSLATOR_URL=<service base url>
// LANGUAGE_TRANSLATOR_AUTH_TYPE=iam
// LANGUAGE_TRANSLATOR_APIKEY=<IAM apikey>
// LANGUAGE_TRANSLATOR_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
var _ = Describe(`LanguageTranslatorV3 Examples Tests`, func() {

	const externalConfigFile = "../language_translator_v3.env"

	var (
		languageTranslatorService *languagetranslatorv3.LanguageTranslatorV3
		config                    map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping examples...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping examples: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(languagetranslatorv3.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping examples: " + err.Error())
			} else if len(config) == 0 {
				Skip("Unable to load service properties, skipping examples")
			}

			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			languageTranslatorServiceOptions := &languagetranslatorv3.LanguageTranslatorV3Options{
				Version: core.StringPtr("2018-05-01"),
			}

			languageTranslatorService, err = languagetranslatorv3.NewLanguageTranslatorV3(languageTranslatorServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(languageTranslatorService).ToNot(BeNil())
		})
	})

	Describe(`LanguageTranslatorV3 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListLanguages request example`, func() {
			fmt.Println("\nListLanguages() result:")
			// begin-listLanguages

			listLanguagesOptions := languageTranslatorService.NewListLanguagesOptions()

			languages, response, err := languageTranslatorService.ListLanguages(listLanguagesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(languages, "", "  ")
			fmt.Println(string(b))

			// end-listLanguages

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(languages).ToNot(BeNil())

		})
		It(`Translate request example`, func() {
			fmt.Println("\nTranslate() result:")
			// begin-translate

			translateOptions := languageTranslatorService.NewTranslateOptions(
				[]string{"testString"},
			)

			translationResult, response, err := languageTranslatorService.Translate(translateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(translationResult, "", "  ")
			fmt.Println(string(b))

			// end-translate

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(translationResult).ToNot(BeNil())

		})
		It(`ListIdentifiableLanguages request example`, func() {
			fmt.Println("\nListIdentifiableLanguages() result:")
			// begin-listIdentifiableLanguages

			listIdentifiableLanguagesOptions := languageTranslatorService.NewListIdentifiableLanguagesOptions()

			identifiableLanguages, response, err := languageTranslatorService.ListIdentifiableLanguages(listIdentifiableLanguagesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(identifiableLanguages, "", "  ")
			fmt.Println(string(b))

			// end-listIdentifiableLanguages

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(identifiableLanguages).ToNot(BeNil())

		})
		It(`Identify request example`, func() {
			fmt.Println("\nIdentify() result:")
			// begin-identify

			identifyOptions := languageTranslatorService.NewIdentifyOptions(
				"testString",
			)

			identifiedLanguages, response, err := languageTranslatorService.Identify(identifyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(identifiedLanguages, "", "  ")
			fmt.Println(string(b))

			// end-identify

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(identifiedLanguages).ToNot(BeNil())

		})
		It(`ListModels request example`, func() {
			fmt.Println("\nListModels() result:")
			// begin-listModels

			listModelsOptions := languageTranslatorService.NewListModelsOptions()

			translationModels, response, err := languageTranslatorService.ListModels(listModelsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(translationModels, "", "  ")
			fmt.Println(string(b))

			// end-listModels

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(translationModels).ToNot(BeNil())

		})
		It(`CreateModel request example`, func() {
			fmt.Println("\nCreateModel() result:")
			// begin-createModel

			createModelOptions := languageTranslatorService.NewCreateModelOptions(
				"testString",
			)

			translationModel, response, err := languageTranslatorService.CreateModel(createModelOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(translationModel, "", "  ")
			fmt.Println(string(b))

			// end-createModel

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(translationModel).ToNot(BeNil())

		})
		It(`GetModel request example`, func() {
			fmt.Println("\nGetModel() result:")
			// begin-getModel

			getModelOptions := languageTranslatorService.NewGetModelOptions(
				"testString",
			)

			translationModel, response, err := languageTranslatorService.GetModel(getModelOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(translationModel, "", "  ")
			fmt.Println(string(b))

			// end-getModel

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(translationModel).ToNot(BeNil())

		})
		It(`ListDocuments request example`, func() {
			fmt.Println("\nListDocuments() result:")
			// begin-listDocuments

			listDocumentsOptions := languageTranslatorService.NewListDocumentsOptions()

			documentList, response, err := languageTranslatorService.ListDocuments(listDocumentsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(documentList, "", "  ")
			fmt.Println(string(b))

			// end-listDocuments

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(documentList).ToNot(BeNil())

		})
		It(`TranslateDocument request example`, func() {
			fmt.Println("\nTranslateDocument() result:")
			// begin-translateDocument

			translateDocumentOptions := languageTranslatorService.NewTranslateDocumentOptions(
				CreateMockReader("This is a mock file."),
				"testString",
			)

			documentStatus, response, err := languageTranslatorService.TranslateDocument(translateDocumentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(documentStatus, "", "  ")
			fmt.Println(string(b))

			// end-translateDocument

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(documentStatus).ToNot(BeNil())

		})
		It(`GetDocumentStatus request example`, func() {
			fmt.Println("\nGetDocumentStatus() result:")
			// begin-getDocumentStatus

			getDocumentStatusOptions := languageTranslatorService.NewGetDocumentStatusOptions(
				"testString",
			)

			documentStatus, response, err := languageTranslatorService.GetDocumentStatus(getDocumentStatusOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(documentStatus, "", "  ")
			fmt.Println(string(b))

			// end-getDocumentStatus

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(documentStatus).ToNot(BeNil())

		})
		It(`GetTranslatedDocument request example`, func() {
			fmt.Println("\nGetTranslatedDocument() result:")
			// begin-getTranslatedDocument

			getTranslatedDocumentOptions := languageTranslatorService.NewGetTranslatedDocumentOptions(
				"testString",
			)

			result, response, err := languageTranslatorService.GetTranslatedDocument(getTranslatedDocumentOptions)
			if err != nil {
				panic(err)
			}
			if result != nil {
				defer result.Close()
				outFile, err := os.Create("result.out")
				if err != nil {
					panic(err)
				}
				defer outFile.Close()
				_, err = io.Copy(outFile, result)
				if err != nil {
					panic(err)
				}
			}

			// end-getTranslatedDocument

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

		})
		It(`DeleteModel request example`, func() {
			fmt.Println("\nDeleteModel() result:")
			// begin-deleteModel

			deleteModelOptions := languageTranslatorService.NewDeleteModelOptions(
				"testString",
			)

			deleteModelResult, response, err := languageTranslatorService.DeleteModel(deleteModelOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(deleteModelResult, "", "  ")
			fmt.Println(string(b))

			// end-deleteModel

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(deleteModelResult).ToNot(BeNil())

		})
		It(`DeleteDocument request example`, func() {
			// begin-deleteDocument

			deleteDocumentOptions := languageTranslatorService.NewDeleteDocumentOptions(
				"testString",
			)

			response, err := languageTranslatorService.DeleteDocument(deleteDocumentOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteDocument(): %d\n", response.StatusCode)
			}

			// end-deleteDocument

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})
})
