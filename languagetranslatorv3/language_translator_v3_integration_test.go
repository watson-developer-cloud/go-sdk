// +build integration

package languagetranslatorv3_test

/**
 * Copyright 2018 IBM All Rights Reserved.
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

import (
	"net/http"
	"os"
	"testing"

	"github.com/IBM/go-sdk-core/core"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/watson-developer-cloud/go-sdk/languagetranslatorv3"
)

var service *languagetranslatorv3.LanguageTranslatorV3
var serviceErr error

func init() {
	err := godotenv.Load("../.env")

	if err == nil {
		service, serviceErr = languagetranslatorv3.
			NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
				URL:      os.Getenv("LANGUAGE_TRANSLATOR_URL"),
				Version:  "2019-06-03",
				Authenticator: &core.BasicAuthenticator{
						Username: os.Getenv("LANGUAGE_TRANSLATOR_USERNAME"),
						Password: os.Getenv("LANGUAGE_TRANSLATOR_PASSWORD"),
                },
			})

		if serviceErr == nil {
			customHeaders := http.Header{}
			customHeaders.Add("X-Watson-Learning-Opt-Out", "1")
			customHeaders.Add("X-Watson-Test", "1")
			service.Service.SetDefaultHeaders(customHeaders)
		}
	}
}

func shouldSkipTest(t *testing.T) {
	if service == nil {
		t.Skip("Skipping test as service credentials are missing")
	}
}

func TestModels(t *testing.T) {
	shouldSkipTest(t)

	// List models
	response, responseErr := service.ListModels(
		&languagetranslatorv3.ListModelsOptions{},
	)
	assert.Nil(t, responseErr)

	listModels := service.GetListModelsResult(response)
	assert.NotNil(t, listModels)

	// Create model
	pwd, _ := os.Getwd()
	glossary, glossaryErr := os.Open(pwd + "/../resources/glossary.tmx")
	assert.Nil(t, glossaryErr)

	response, responseErr = service.CreateModel(
		&languagetranslatorv3.CreateModelOptions{
			BaseModelID:    core.StringPtr("en-es"),
			Name:           core.StringPtr("custom-en-es"),
			ForcedGlossary: glossary,
		},
	)
	assert.Nil(t, responseErr)

	createModel := service.GetCreateModelResult(response)
	assert.NotNil(t, createModel)

	// Get model
	response, responseErr = service.GetModel(
		&languagetranslatorv3.GetModelOptions{
			ModelID: createModel.ModelID,
		},
	)
	assert.Nil(t, responseErr)

	getModel := service.GetGetModelResult(response)
	assert.NotNil(t, getModel)

	// Delete model
	response, responseErr = service.DeleteModel(
		&languagetranslatorv3.DeleteModelOptions{
			ModelID: createModel.ModelID,
		},
	)
	assert.Nil(t, responseErr)
}

func TestTranslate(t *testing.T) {
	shouldSkipTest(t)

	response, responseErr := service.Translate(
		&languagetranslatorv3.TranslateOptions{
			Text:    []string{"Hello"},
			ModelID: core.StringPtr("en-es"),
		},
	)
	assert.Nil(t, responseErr)

	translate := service.GetTranslateResult(response)
	assert.NotNil(t, translate)
}

func TestIdentifiableLanguage(t *testing.T) {
	shouldSkipTest(t)

	response, responseErr := service.ListIdentifiableLanguages(
		&languagetranslatorv3.ListIdentifiableLanguagesOptions{},
	)
	assert.Nil(t, responseErr)

	identifiableLanguage := service.GetListIdentifiableLanguagesResult(response)
	assert.NotNil(t, identifiableLanguage)
}

func TestIdentify(t *testing.T) {
	shouldSkipTest(t)

	response, responseErr := service.Identify(
		&languagetranslatorv3.IdentifyOptions{
			Text: core.StringPtr("Language translator translates text from one language to another"),
		},
	)
	assert.Nil(t, responseErr)

	identify := service.GetIdentifyResult(response)
	assert.NotNil(t, identify)
}

func TestDocumentTranslation(t *testing.T) {
	shouldSkipTest(t)

	// List documents
	response, responseErr := service.ListDocuments(
		&languagetranslatorv3.ListDocumentsOptions{},
	)
	assert.Nil(t, responseErr)

	listDocuments := service.GetListDocumentsResult(response)
	assert.NotNil(t, listDocuments)

	// translate document
	pwd, _ := os.Getwd()
	document, documentErr := os.Open(pwd + "/../resources/hello_world.txt")
	assert.Nil(t, documentErr)

	response, responseErr = service.TranslateDocument(
		&languagetranslatorv3.TranslateDocumentOptions{
			File:            document,
			Filename:        core.StringPtr("hello_world"),
			FileContentType: core.StringPtr("text/plain"),
			ModelID:         core.StringPtr("en-es"),
		},
	)
	assert.Nil(t, responseErr)

	translateDocument := service.GetTranslateDocumentResult(response)
	assert.NotNil(t, translateDocument)

	// Document status
	response, responseErr = service.GetDocumentStatus(
		&languagetranslatorv3.GetDocumentStatusOptions{
			DocumentID: translateDocument.DocumentID,
		},
	)
	assert.Nil(t, responseErr)

	documentStatus := service.GetGetDocumentStatusResult(response)
	assert.NotNil(t, documentStatus)

	// Delete document
	response, responseErr = service.DeleteDocument(
		&languagetranslatorv3.DeleteDocumentOptions{
			DocumentID: translateDocument.DocumentID,
		},
	)
	assert.Nil(t, responseErr)
}
