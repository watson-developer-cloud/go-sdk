// +build integration

package languagetranslatorv3_test

/**
 * (C) Copyright IBM Corp. 2018, 2021.
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

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/watson-developer-cloud/go-sdk/v2/languagetranslatorv3"
)

const skipMessage = "External configuration could not be loaded, skipping..."

var configLoaded bool
var configFile = "../../.env"

var service *languagetranslatorv3.LanguageTranslatorV3

func shouldSkipTest(t *testing.T) {
	if !configLoaded {
		t.Skip(skipMessage)
	}
}

func TestLoadConfig(t *testing.T) {
	err := godotenv.Load(configFile)
	if err != nil {
		t.Skip(skipMessage)
	} else {
		configLoaded = true
	}
}

func TestConstructService(t *testing.T) {
	shouldSkipTest(t)

	var err error

	service, err = languagetranslatorv3.NewLanguageTranslatorV3(
		&languagetranslatorv3.LanguageTranslatorV3Options{
			Version: core.StringPtr("2020-04-01"),
		})
	assert.Nil(t, err)
	assert.NotNil(t, service)

	if err == nil {
		customHeaders := http.Header{}
		customHeaders.Add("X-Watson-Learning-Opt-Out", "1")
		customHeaders.Add("X-Watson-Test", "1")
		service.Service.SetDefaultHeaders(customHeaders)
	}
}

func TestModels(t *testing.T) {
	shouldSkipTest(t)

	// List models
	listModels, response, responseErr := service.ListModels(
		&languagetranslatorv3.ListModelsOptions{},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, response)
	assert.NotNil(t, listModels)

	// Create model
	glossary, glossaryErr := os.Open("../resources/glossary.tmx")
	assert.Nil(t, glossaryErr)

	createModel, _, responseErr := service.CreateModel(
		&languagetranslatorv3.CreateModelOptions{
			BaseModelID:    core.StringPtr("en-fr"),
			Name:           core.StringPtr("custom-en-fr"),
			ForcedGlossary: glossary,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, createModel)

	// Get model
	getModel, _, responseErr := service.GetModel(
		&languagetranslatorv3.GetModelOptions{
			ModelID: createModel.ModelID,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, getModel)

	// Delete model
	_, _, responseErr = service.DeleteModel(
		&languagetranslatorv3.DeleteModelOptions{
			ModelID: createModel.ModelID,
		},
	)
	assert.Nil(t, responseErr)
}

func TestTranslate(t *testing.T) {
	shouldSkipTest(t)

	translate, _, responseErr := service.Translate(
		&languagetranslatorv3.TranslateOptions{
			Text:    []string{"Hello"},
			ModelID: core.StringPtr("en-es"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, translate)
}

func TestIdentifiableLanguage(t *testing.T) {
	shouldSkipTest(t)

	identifiableLanguage, _, responseErr := service.ListIdentifiableLanguages(
		&languagetranslatorv3.ListIdentifiableLanguagesOptions{},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, identifiableLanguage)
}

func TestIdentify(t *testing.T) {
	shouldSkipTest(t)

	identify, _, responseErr := service.Identify(
		&languagetranslatorv3.IdentifyOptions{
			Text: core.StringPtr("Language translator translates text from one language to another"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, identify)
}

func TestDocumentTranslation(t *testing.T) {
	shouldSkipTest(t)

	// List documents
	listDocuments, _, responseErr := service.ListDocuments(
		&languagetranslatorv3.ListDocumentsOptions{},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, listDocuments)

	// translate document
	pwd, _ := os.Getwd()
	document, documentErr := os.Open(pwd + "/../resources/hello_world.txt")
	assert.Nil(t, documentErr)

	translateDocument, _, responseErr := service.TranslateDocument(
		&languagetranslatorv3.TranslateDocumentOptions{
			File:            document,
			Filename:        core.StringPtr("hello_world"),
			FileContentType: core.StringPtr("text/plain"),
			ModelID:         core.StringPtr("en-es"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, translateDocument)

	// Document status
	documentStatus, _, responseErr := service.GetDocumentStatus(
		&languagetranslatorv3.GetDocumentStatusOptions{
			DocumentID: translateDocument.DocumentID,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, documentStatus)

	// Delete document
	_, responseErr = service.DeleteDocument(
		&languagetranslatorv3.DeleteDocumentOptions{
			DocumentID: translateDocument.DocumentID,
		},
	)
	assert.Nil(t, responseErr)
}

func TestListLanguages(t *testing.T) {
	shouldSkipTest(t)

	listLanguages, _, listErr := service.ListLanguages(
		&languagetranslatorv3.ListLanguagesOptions{},
	)

	assert.NotNil(t, listLanguages)
	assert.Nil(t, listErr)
}
