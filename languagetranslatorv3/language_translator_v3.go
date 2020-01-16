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

// Package languagetranslatorv3 : Operations and models for the LanguageTranslatorV3 service
package languagetranslatorv3

import (
	"fmt"
	"github.com/IBM/go-sdk-core/core"
	"github.com/go-openapi/strfmt"
	common "github.com/watson-developer-cloud/go-sdk/common"
	"io"
)

// LanguageTranslatorV3 : IBM Watson&trade; Language Translator translates text from one language to another. The
// service offers multiple IBM provided translation models that you can customize based on your unique terminology and
// language. Use Language Translator to take news from across the globe and present it in your language, communicate
// with your customers in their own language, and more.
//
// Version: 3.0.0
// See: https://cloud.ibm.com/docs/services/language-translator/
type LanguageTranslatorV3 struct {
	Service *core.BaseService
	Version string
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://gateway.watsonplatform.net/language-translator/api"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "language_translator"

// LanguageTranslatorV3Options : Service options
type LanguageTranslatorV3Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
	Version       string
}

// NewLanguageTranslatorV3 : constructs an instance of LanguageTranslatorV3 with passed in options.
func NewLanguageTranslatorV3(options *LanguageTranslatorV3Options) (service *LanguageTranslatorV3, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	if serviceOptions.Authenticator == nil {
		serviceOptions.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	err = baseService.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &LanguageTranslatorV3{
		Service: baseService,
		Version: options.Version,
	}

	return
}

// SetServiceURL sets the service URL
func (languageTranslator *LanguageTranslatorV3) SetServiceURL(url string) error {
	return languageTranslator.Service.SetServiceURL(url)
}

// DisableSSLVerification bypasses verification of the server's SSL certificate
func (languageTranslator *LanguageTranslatorV3) DisableSSLVerification() {
	languageTranslator.Service.DisableSSLVerification()
}

// Translate : Translate
// Translates the input text from the source language to the target language.
func (languageTranslator *LanguageTranslatorV3) Translate(translateOptions *TranslateOptions) (result *TranslationResult, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(translateOptions, "translateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(translateOptions, "translateOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v3/translate"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(languageTranslator.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range translateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("language_translator", "V3", "Translate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", languageTranslator.Version)

	body := make(map[string]interface{})
	if translateOptions.Text != nil {
		body["text"] = translateOptions.Text
	}
	if translateOptions.ModelID != nil {
		body["model_id"] = translateOptions.ModelID
	}
	if translateOptions.Source != nil {
		body["source"] = translateOptions.Source
	}
	if translateOptions.Target != nil {
		body["target"] = translateOptions.Target
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = languageTranslator.Service.Request(request, new(TranslationResult))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*TranslationResult)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// ListIdentifiableLanguages : List identifiable languages
// Lists the languages that the service can identify. Returns the language code (for example, `en` for English or `es`
// for Spanish) and name of each language.
func (languageTranslator *LanguageTranslatorV3) ListIdentifiableLanguages(listIdentifiableLanguagesOptions *ListIdentifiableLanguagesOptions) (result *IdentifiableLanguages, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listIdentifiableLanguagesOptions, "listIdentifiableLanguagesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v3/identifiable_languages"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(languageTranslator.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listIdentifiableLanguagesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("language_translator", "V3", "ListIdentifiableLanguages")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", languageTranslator.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = languageTranslator.Service.Request(request, new(IdentifiableLanguages))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*IdentifiableLanguages)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// Identify : Identify language
// Identifies the language of the input text.
func (languageTranslator *LanguageTranslatorV3) Identify(identifyOptions *IdentifyOptions) (result *IdentifiedLanguages, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(identifyOptions, "identifyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(identifyOptions, "identifyOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v3/identify"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(languageTranslator.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range identifyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("language_translator", "V3", "Identify")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "text/plain")
	builder.AddQuery("version", languageTranslator.Version)

	_, err = builder.SetBodyContent("text/plain", nil, nil, identifyOptions.Text)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = languageTranslator.Service.Request(request, new(IdentifiedLanguages))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*IdentifiedLanguages)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// ListModels : List models
// Lists available translation models.
func (languageTranslator *LanguageTranslatorV3) ListModels(listModelsOptions *ListModelsOptions) (result *TranslationModels, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listModelsOptions, "listModelsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v3/models"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(languageTranslator.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listModelsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("language_translator", "V3", "ListModels")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")

	if listModelsOptions.Source != nil {
		builder.AddQuery("source", fmt.Sprint(*listModelsOptions.Source))
	}
	if listModelsOptions.Target != nil {
		builder.AddQuery("target", fmt.Sprint(*listModelsOptions.Target))
	}
	if listModelsOptions.Default != nil {
		builder.AddQuery("default", fmt.Sprint(*listModelsOptions.Default))
	}
	builder.AddQuery("version", languageTranslator.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = languageTranslator.Service.Request(request, new(TranslationModels))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*TranslationModels)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// CreateModel : Create model
// Uploads Translation Memory eXchange (TMX) files to customize a translation model.
//
// You can either customize a model with a forced glossary or with a corpus that contains parallel sentences. To create
// a model that is customized with a parallel corpus <b>and</b> a forced glossary, proceed in two steps: customize with
// a parallel corpus first and then customize the resulting model with a glossary. Depending on the type of
// customization and the size of the uploaded corpora, training can range from minutes for a glossary to several hours
// for a large parallel corpus. You can upload a single forced glossary file and this file must be less than <b>10
// MB</b>. You can upload multiple parallel corpora tmx files. The cumulative file size of all uploaded files is limited
// to <b>250 MB</b>. To successfully train with a parallel corpus you must have at least <b>5,000 parallel sentences</b>
// in your corpus.
//
// You can have a <b>maximum of 10 custom models per language pair</b>.
func (languageTranslator *LanguageTranslatorV3) CreateModel(createModelOptions *CreateModelOptions) (result *TranslationModel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createModelOptions, "createModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createModelOptions, "createModelOptions")
	if err != nil {
		return
	}
	if (createModelOptions.ForcedGlossary == nil) && (createModelOptions.ParallelCorpus == nil) {
		err = fmt.Errorf("At least one of forcedGlossary or parallelCorpus must be supplied")
		return
	}

	pathSegments := []string{"v3/models"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(languageTranslator.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("language_translator", "V3", "CreateModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("base_model_id", fmt.Sprint(*createModelOptions.BaseModelID))
	if createModelOptions.Name != nil {
		builder.AddQuery("name", fmt.Sprint(*createModelOptions.Name))
	}
	builder.AddQuery("version", languageTranslator.Version)

	if createModelOptions.ForcedGlossary != nil {
		builder.AddFormData("forced_glossary", "filename",
			"application/octet-stream", createModelOptions.ForcedGlossary)
	}
	if createModelOptions.ParallelCorpus != nil {
		builder.AddFormData("parallel_corpus", "filename",
			"application/octet-stream", createModelOptions.ParallelCorpus)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = languageTranslator.Service.Request(request, new(TranslationModel))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*TranslationModel)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// DeleteModel : Delete model
// Deletes a custom translation model.
func (languageTranslator *LanguageTranslatorV3) DeleteModel(deleteModelOptions *DeleteModelOptions) (result *DeleteModelResult, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteModelOptions, "deleteModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteModelOptions, "deleteModelOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v3/models"}
	pathParameters := []string{*deleteModelOptions.ModelID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(languageTranslator.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("language_translator", "V3", "DeleteModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", languageTranslator.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = languageTranslator.Service.Request(request, new(DeleteModelResult))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*DeleteModelResult)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// GetModel : Get model details
// Gets information about a translation model, including training status for custom models. Use this API call to poll
// the status of your customization request. A successfully completed training will have a status of `available`.
func (languageTranslator *LanguageTranslatorV3) GetModel(getModelOptions *GetModelOptions) (result *TranslationModel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getModelOptions, "getModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getModelOptions, "getModelOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v3/models"}
	pathParameters := []string{*getModelOptions.ModelID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(languageTranslator.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("language_translator", "V3", "GetModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", languageTranslator.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = languageTranslator.Service.Request(request, new(TranslationModel))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*TranslationModel)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// ListDocuments : List documents
// Lists documents that have been submitted for translation.
func (languageTranslator *LanguageTranslatorV3) ListDocuments(listDocumentsOptions *ListDocumentsOptions) (result *DocumentList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listDocumentsOptions, "listDocumentsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v3/documents"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(languageTranslator.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listDocumentsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("language_translator", "V3", "ListDocuments")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", languageTranslator.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = languageTranslator.Service.Request(request, new(DocumentList))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*DocumentList)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// TranslateDocument : Translate document
// Submit a document for translation. You can submit the document contents in the `file` parameter, or you can reference
// a previously submitted document by document ID.
func (languageTranslator *LanguageTranslatorV3) TranslateDocument(translateDocumentOptions *TranslateDocumentOptions) (result *DocumentStatus, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(translateDocumentOptions, "translateDocumentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(translateDocumentOptions, "translateDocumentOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v3/documents"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(languageTranslator.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range translateDocumentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("language_translator", "V3", "TranslateDocument")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", languageTranslator.Version)

	builder.AddFormData("file", core.StringNilMapper(translateDocumentOptions.Filename),
		core.StringNilMapper(translateDocumentOptions.FileContentType), translateDocumentOptions.File)
	if translateDocumentOptions.ModelID != nil {
		builder.AddFormData("model_id", "", "", fmt.Sprint(*translateDocumentOptions.ModelID))
	}
	if translateDocumentOptions.Source != nil {
		builder.AddFormData("source", "", "", fmt.Sprint(*translateDocumentOptions.Source))
	}
	if translateDocumentOptions.Target != nil {
		builder.AddFormData("target", "", "", fmt.Sprint(*translateDocumentOptions.Target))
	}
	if translateDocumentOptions.DocumentID != nil {
		builder.AddFormData("document_id", "", "", fmt.Sprint(*translateDocumentOptions.DocumentID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = languageTranslator.Service.Request(request, new(DocumentStatus))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*DocumentStatus)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// GetDocumentStatus : Get document status
// Gets the translation status of a document.
func (languageTranslator *LanguageTranslatorV3) GetDocumentStatus(getDocumentStatusOptions *GetDocumentStatusOptions) (result *DocumentStatus, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDocumentStatusOptions, "getDocumentStatusOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getDocumentStatusOptions, "getDocumentStatusOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v3/documents"}
	pathParameters := []string{*getDocumentStatusOptions.DocumentID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(languageTranslator.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getDocumentStatusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("language_translator", "V3", "GetDocumentStatus")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", languageTranslator.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = languageTranslator.Service.Request(request, new(DocumentStatus))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*DocumentStatus)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// DeleteDocument : Delete document
// Deletes a document.
func (languageTranslator *LanguageTranslatorV3) DeleteDocument(deleteDocumentOptions *DeleteDocumentOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteDocumentOptions, "deleteDocumentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteDocumentOptions, "deleteDocumentOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v3/documents"}
	pathParameters := []string{*deleteDocumentOptions.DocumentID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(languageTranslator.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteDocumentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("language_translator", "V3", "DeleteDocument")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("version", languageTranslator.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = languageTranslator.Service.Request(request, nil)

	return
}

// GetTranslatedDocument : Get translated document
// Gets the translated document associated with the given document ID.
func (languageTranslator *LanguageTranslatorV3) GetTranslatedDocument(getTranslatedDocumentOptions *GetTranslatedDocumentOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getTranslatedDocumentOptions, "getTranslatedDocumentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getTranslatedDocumentOptions, "getTranslatedDocumentOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v3/documents", "translated_document"}
	pathParameters := []string{*getTranslatedDocumentOptions.DocumentID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(languageTranslator.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getTranslatedDocumentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("language_translator", "V3", "GetTranslatedDocument")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/powerpoint")
	if getTranslatedDocumentOptions.Accept != nil {
		builder.AddHeader("Accept", fmt.Sprint(*getTranslatedDocumentOptions.Accept))
	}
	builder.AddQuery("version", languageTranslator.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = languageTranslator.Service.Request(request, new(io.ReadCloser))
	if err == nil {
		var ok bool
		result, ok = response.Result.(io.ReadCloser)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// CreateModelOptions : The CreateModel options.
type CreateModelOptions struct {

	// The model ID of the model to use as the base for customization. To see available models, use the `List models`
	// method. Usually all IBM provided models are customizable. In addition, all your models that have been created via
	// parallel corpus customization, can be further customized with a forced glossary.
	BaseModelID *string `json:"base_model_id" validate:"required"`

	// A TMX file with your customizations. The customizations in the file completely overwrite the domain translaton data,
	// including high frequency or high confidence phrase translations. You can upload only one glossary with a file size
	// less than 10 MB per call. A forced glossary should contain single words or short phrases.
	ForcedGlossary io.ReadCloser `json:"forced_glossary,omitempty"`

	// A TMX file with parallel sentences for source and target language. You can upload multiple parallel_corpus files in
	// one request. All uploaded parallel_corpus files combined, your parallel corpus must contain at least 5,000 parallel
	// sentences to train successfully.
	ParallelCorpus io.ReadCloser `json:"parallel_corpus,omitempty"`

	// An optional model name that you can use to identify the model. Valid characters are letters, numbers, dashes,
	// underscores, spaces and apostrophes. The maximum length is 32 characters.
	Name *string `json:"name,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewCreateModelOptions : Instantiate CreateModelOptions
func (languageTranslator *LanguageTranslatorV3) NewCreateModelOptions(baseModelID string) *CreateModelOptions {
	return &CreateModelOptions{
		BaseModelID: core.StringPtr(baseModelID),
	}
}

// SetBaseModelID : Allow user to set BaseModelID
func (options *CreateModelOptions) SetBaseModelID(baseModelID string) *CreateModelOptions {
	options.BaseModelID = core.StringPtr(baseModelID)
	return options
}

// SetForcedGlossary : Allow user to set ForcedGlossary
func (options *CreateModelOptions) SetForcedGlossary(forcedGlossary io.ReadCloser) *CreateModelOptions {
	options.ForcedGlossary = forcedGlossary
	return options
}

// SetParallelCorpus : Allow user to set ParallelCorpus
func (options *CreateModelOptions) SetParallelCorpus(parallelCorpus io.ReadCloser) *CreateModelOptions {
	options.ParallelCorpus = parallelCorpus
	return options
}

// SetName : Allow user to set Name
func (options *CreateModelOptions) SetName(name string) *CreateModelOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateModelOptions) SetHeaders(param map[string]string) *CreateModelOptions {
	options.Headers = param
	return options
}

// DeleteDocumentOptions : The DeleteDocument options.
type DeleteDocumentOptions struct {

	// Document ID of the document to delete.
	DocumentID *string `json:"document_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteDocumentOptions : Instantiate DeleteDocumentOptions
func (languageTranslator *LanguageTranslatorV3) NewDeleteDocumentOptions(documentID string) *DeleteDocumentOptions {
	return &DeleteDocumentOptions{
		DocumentID: core.StringPtr(documentID),
	}
}

// SetDocumentID : Allow user to set DocumentID
func (options *DeleteDocumentOptions) SetDocumentID(documentID string) *DeleteDocumentOptions {
	options.DocumentID = core.StringPtr(documentID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteDocumentOptions) SetHeaders(param map[string]string) *DeleteDocumentOptions {
	options.Headers = param
	return options
}

// DeleteModelOptions : The DeleteModel options.
type DeleteModelOptions struct {

	// Model ID of the model to delete.
	ModelID *string `json:"model_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteModelOptions : Instantiate DeleteModelOptions
func (languageTranslator *LanguageTranslatorV3) NewDeleteModelOptions(modelID string) *DeleteModelOptions {
	return &DeleteModelOptions{
		ModelID: core.StringPtr(modelID),
	}
}

// SetModelID : Allow user to set ModelID
func (options *DeleteModelOptions) SetModelID(modelID string) *DeleteModelOptions {
	options.ModelID = core.StringPtr(modelID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteModelOptions) SetHeaders(param map[string]string) *DeleteModelOptions {
	options.Headers = param
	return options
}

// DeleteModelResult : DeleteModelResult struct
type DeleteModelResult struct {

	// "OK" indicates that the model was successfully deleted.
	Status *string `json:"status" validate:"required"`
}

// DocumentList : DocumentList struct
type DocumentList struct {

	// An array of all previously submitted documents.
	Documents []DocumentStatus `json:"documents" validate:"required"`
}

// DocumentStatus : Document information, including translation status.
type DocumentStatus struct {

	// System generated ID identifying a document being translated using one specific translation model.
	DocumentID *string `json:"document_id" validate:"required"`

	// filename from the submission (if it was missing in the multipart-form, 'noname.<ext matching content type>' is used.
	Filename *string `json:"filename" validate:"required"`

	// The status of the translation job associated with a submitted document.
	Status *string `json:"status" validate:"required"`

	// A globally unique string that identifies the underlying model that is used for translation.
	ModelID *string `json:"model_id" validate:"required"`

	// Model ID of the base model that was used to customize the model. If the model is not a custom model, this will be
	// absent or an empty string.
	BaseModelID *string `json:"base_model_id,omitempty"`

	// Translation source language code.
	Source *string `json:"source" validate:"required"`

	// Translation target language code.
	Target *string `json:"target" validate:"required"`

	// The time when the document was submitted.
	Created *strfmt.DateTime `json:"created" validate:"required"`

	// The time when the translation completed.
	Completed *strfmt.DateTime `json:"completed,omitempty"`

	// An estimate of the number of words in the source document. Returned only if `status` is `available`.
	WordCount *int64 `json:"word_count,omitempty"`

	// The number of characters in the source document, present only if status=available.
	CharacterCount *int64 `json:"character_count,omitempty"`
}

// Constants associated with the DocumentStatus.Status property.
// The status of the translation job associated with a submitted document.
const (
	DocumentStatus_Status_Available  = "available"
	DocumentStatus_Status_Failed     = "failed"
	DocumentStatus_Status_Processing = "processing"
)

// GetDocumentStatusOptions : The GetDocumentStatus options.
type GetDocumentStatusOptions struct {

	// The document ID of the document.
	DocumentID *string `json:"document_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetDocumentStatusOptions : Instantiate GetDocumentStatusOptions
func (languageTranslator *LanguageTranslatorV3) NewGetDocumentStatusOptions(documentID string) *GetDocumentStatusOptions {
	return &GetDocumentStatusOptions{
		DocumentID: core.StringPtr(documentID),
	}
}

// SetDocumentID : Allow user to set DocumentID
func (options *GetDocumentStatusOptions) SetDocumentID(documentID string) *GetDocumentStatusOptions {
	options.DocumentID = core.StringPtr(documentID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetDocumentStatusOptions) SetHeaders(param map[string]string) *GetDocumentStatusOptions {
	options.Headers = param
	return options
}

// GetModelOptions : The GetModel options.
type GetModelOptions struct {

	// Model ID of the model to get.
	ModelID *string `json:"model_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetModelOptions : Instantiate GetModelOptions
func (languageTranslator *LanguageTranslatorV3) NewGetModelOptions(modelID string) *GetModelOptions {
	return &GetModelOptions{
		ModelID: core.StringPtr(modelID),
	}
}

// SetModelID : Allow user to set ModelID
func (options *GetModelOptions) SetModelID(modelID string) *GetModelOptions {
	options.ModelID = core.StringPtr(modelID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetModelOptions) SetHeaders(param map[string]string) *GetModelOptions {
	options.Headers = param
	return options
}

// GetTranslatedDocumentOptions : The GetTranslatedDocument options.
type GetTranslatedDocumentOptions struct {

	// The document ID of the document that was submitted for translation.
	DocumentID *string `json:"document_id" validate:"required"`

	// The type of the response: application/powerpoint, application/mspowerpoint, application/x-rtf, application/json,
	// application/xml, application/vnd.ms-excel, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet,
	// application/vnd.ms-powerpoint, application/vnd.openxmlformats-officedocument.presentationml.presentation,
	// application/msword, application/vnd.openxmlformats-officedocument.wordprocessingml.document,
	// application/vnd.oasis.opendocument.spreadsheet, application/vnd.oasis.opendocument.presentation,
	// application/vnd.oasis.opendocument.text, application/pdf, application/rtf, text/html, text/json, text/plain,
	// text/richtext, text/rtf, or text/xml. A character encoding can be specified by including a `charset` parameter. For
	// example, 'text/html;charset=utf-8'.
	Accept *string `json:"Accept,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetTranslatedDocumentOptions : Instantiate GetTranslatedDocumentOptions
func (languageTranslator *LanguageTranslatorV3) NewGetTranslatedDocumentOptions(documentID string) *GetTranslatedDocumentOptions {
	return &GetTranslatedDocumentOptions{
		DocumentID: core.StringPtr(documentID),
	}
}

// SetDocumentID : Allow user to set DocumentID
func (options *GetTranslatedDocumentOptions) SetDocumentID(documentID string) *GetTranslatedDocumentOptions {
	options.DocumentID = core.StringPtr(documentID)
	return options
}

// SetAccept : Allow user to set Accept
func (options *GetTranslatedDocumentOptions) SetAccept(accept string) *GetTranslatedDocumentOptions {
	options.Accept = core.StringPtr(accept)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetTranslatedDocumentOptions) SetHeaders(param map[string]string) *GetTranslatedDocumentOptions {
	options.Headers = param
	return options
}

// IdentifiableLanguage : IdentifiableLanguage struct
type IdentifiableLanguage struct {

	// The language code for an identifiable language.
	Language *string `json:"language" validate:"required"`

	// The name of the identifiable language.
	Name *string `json:"name" validate:"required"`
}

// IdentifiableLanguages : IdentifiableLanguages struct
type IdentifiableLanguages struct {

	// A list of all languages that the service can identify.
	Languages []IdentifiableLanguage `json:"languages" validate:"required"`
}

// IdentifiedLanguage : IdentifiedLanguage struct
type IdentifiedLanguage struct {

	// The language code for an identified language.
	Language *string `json:"language" validate:"required"`

	// The confidence score for the identified language.
	Confidence *float64 `json:"confidence" validate:"required"`
}

// IdentifiedLanguages : IdentifiedLanguages struct
type IdentifiedLanguages struct {

	// A ranking of identified languages with confidence scores.
	Languages []IdentifiedLanguage `json:"languages" validate:"required"`
}

// IdentifyOptions : The Identify options.
type IdentifyOptions struct {

	// Input text in UTF-8 format.
	Text *string `json:"text" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewIdentifyOptions : Instantiate IdentifyOptions
func (languageTranslator *LanguageTranslatorV3) NewIdentifyOptions(text string) *IdentifyOptions {
	return &IdentifyOptions{
		Text: core.StringPtr(text),
	}
}

// SetText : Allow user to set Text
func (options *IdentifyOptions) SetText(text string) *IdentifyOptions {
	options.Text = core.StringPtr(text)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *IdentifyOptions) SetHeaders(param map[string]string) *IdentifyOptions {
	options.Headers = param
	return options
}

// ListDocumentsOptions : The ListDocuments options.
type ListDocumentsOptions struct {

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListDocumentsOptions : Instantiate ListDocumentsOptions
func (languageTranslator *LanguageTranslatorV3) NewListDocumentsOptions() *ListDocumentsOptions {
	return &ListDocumentsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListDocumentsOptions) SetHeaders(param map[string]string) *ListDocumentsOptions {
	options.Headers = param
	return options
}

// ListIdentifiableLanguagesOptions : The ListIdentifiableLanguages options.
type ListIdentifiableLanguagesOptions struct {

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListIdentifiableLanguagesOptions : Instantiate ListIdentifiableLanguagesOptions
func (languageTranslator *LanguageTranslatorV3) NewListIdentifiableLanguagesOptions() *ListIdentifiableLanguagesOptions {
	return &ListIdentifiableLanguagesOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListIdentifiableLanguagesOptions) SetHeaders(param map[string]string) *ListIdentifiableLanguagesOptions {
	options.Headers = param
	return options
}

// ListModelsOptions : The ListModels options.
type ListModelsOptions struct {

	// Specify a language code to filter results by source language.
	Source *string `json:"source,omitempty"`

	// Specify a language code to filter results by target language.
	Target *string `json:"target,omitempty"`

	// If the default parameter isn't specified, the service will return all models (default and non-default) for each
	// language pair. To return only default models, set this to `true`. To return only non-default models, set this to
	// `false`. There is exactly one default model per language pair, the IBM provided base model.
	Default *bool `json:"default,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListModelsOptions : Instantiate ListModelsOptions
func (languageTranslator *LanguageTranslatorV3) NewListModelsOptions() *ListModelsOptions {
	return &ListModelsOptions{}
}

// SetSource : Allow user to set Source
func (options *ListModelsOptions) SetSource(source string) *ListModelsOptions {
	options.Source = core.StringPtr(source)
	return options
}

// SetTarget : Allow user to set Target
func (options *ListModelsOptions) SetTarget(target string) *ListModelsOptions {
	options.Target = core.StringPtr(target)
	return options
}

// SetDefault : Allow user to set Default
func (options *ListModelsOptions) SetDefault(defaultVar bool) *ListModelsOptions {
	options.Default = core.BoolPtr(defaultVar)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListModelsOptions) SetHeaders(param map[string]string) *ListModelsOptions {
	options.Headers = param
	return options
}

// TranslateDocumentOptions : The TranslateDocument options.
type TranslateDocumentOptions struct {

	// The contents of the source file to translate.
	//
	// [Supported file
	// types](https://cloud.ibm.com/docs/services/language-translator?topic=language-translator-document-translator-tutorial#supported-file-formats)
	//
	// Maximum file size: **20 MB**.
	File io.ReadCloser `json:"file" validate:"required"`

	// The filename for file.
	Filename *string `json:"filename" validate:"required"`

	// The content type of file.
	FileContentType *string `json:"file_content_type,omitempty"`

	// The model to use for translation. `model_id` or both `source` and `target` are required.
	ModelID *string `json:"model_id,omitempty"`

	// Language code that specifies the language of the source document.
	Source *string `json:"source,omitempty"`

	// Language code that specifies the target language for translation.
	Target *string `json:"target,omitempty"`

	// To use a previously submitted document as the source for a new translation, enter the `document_id` of the document.
	DocumentID *string `json:"document_id,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewTranslateDocumentOptions : Instantiate TranslateDocumentOptions
func (languageTranslator *LanguageTranslatorV3) NewTranslateDocumentOptions(file io.ReadCloser, filename string) *TranslateDocumentOptions {
	return &TranslateDocumentOptions{
		File:     file,
		Filename: core.StringPtr(filename),
	}
}

// SetFile : Allow user to set File
func (options *TranslateDocumentOptions) SetFile(file io.ReadCloser) *TranslateDocumentOptions {
	options.File = file
	return options
}

// SetFilename : Allow user to set Filename
func (options *TranslateDocumentOptions) SetFilename(filename string) *TranslateDocumentOptions {
	options.Filename = core.StringPtr(filename)
	return options
}

// SetFileContentType : Allow user to set FileContentType
func (options *TranslateDocumentOptions) SetFileContentType(fileContentType string) *TranslateDocumentOptions {
	options.FileContentType = core.StringPtr(fileContentType)
	return options
}

// SetModelID : Allow user to set ModelID
func (options *TranslateDocumentOptions) SetModelID(modelID string) *TranslateDocumentOptions {
	options.ModelID = core.StringPtr(modelID)
	return options
}

// SetSource : Allow user to set Source
func (options *TranslateDocumentOptions) SetSource(source string) *TranslateDocumentOptions {
	options.Source = core.StringPtr(source)
	return options
}

// SetTarget : Allow user to set Target
func (options *TranslateDocumentOptions) SetTarget(target string) *TranslateDocumentOptions {
	options.Target = core.StringPtr(target)
	return options
}

// SetDocumentID : Allow user to set DocumentID
func (options *TranslateDocumentOptions) SetDocumentID(documentID string) *TranslateDocumentOptions {
	options.DocumentID = core.StringPtr(documentID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *TranslateDocumentOptions) SetHeaders(param map[string]string) *TranslateDocumentOptions {
	options.Headers = param
	return options
}

// TranslateOptions : The Translate options.
type TranslateOptions struct {

	// Input text in UTF-8 encoding. Multiple entries will result in multiple translations in the response.
	Text []string `json:"text" validate:"required"`

	// A globally unique string that identifies the underlying model that is used for translation.
	ModelID *string `json:"model_id,omitempty"`

	// Translation source language code.
	Source *string `json:"source,omitempty"`

	// Translation target language code.
	Target *string `json:"target,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewTranslateOptions : Instantiate TranslateOptions
func (languageTranslator *LanguageTranslatorV3) NewTranslateOptions(text []string) *TranslateOptions {
	return &TranslateOptions{
		Text: text,
	}
}

// SetText : Allow user to set Text
func (options *TranslateOptions) SetText(text []string) *TranslateOptions {
	options.Text = text
	return options
}

// SetModelID : Allow user to set ModelID
func (options *TranslateOptions) SetModelID(modelID string) *TranslateOptions {
	options.ModelID = core.StringPtr(modelID)
	return options
}

// SetSource : Allow user to set Source
func (options *TranslateOptions) SetSource(source string) *TranslateOptions {
	options.Source = core.StringPtr(source)
	return options
}

// SetTarget : Allow user to set Target
func (options *TranslateOptions) SetTarget(target string) *TranslateOptions {
	options.Target = core.StringPtr(target)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *TranslateOptions) SetHeaders(param map[string]string) *TranslateOptions {
	options.Headers = param
	return options
}

// Translation : Translation struct
type Translation struct {

	// Translation output in UTF-8.
	Translation *string `json:"translation" validate:"required"`
}

// TranslationModel : Response payload for models.
type TranslationModel struct {

	// A globally unique string that identifies the underlying model that is used for translation.
	ModelID *string `json:"model_id" validate:"required"`

	// Optional name that can be specified when the model is created.
	Name *string `json:"name,omitempty"`

	// Translation source language code.
	Source *string `json:"source,omitempty"`

	// Translation target language code.
	Target *string `json:"target,omitempty"`

	// Model ID of the base model that was used to customize the model. If the model is not a custom model, this will be an
	// empty string.
	BaseModelID *string `json:"base_model_id,omitempty"`

	// The domain of the translation model.
	Domain *string `json:"domain,omitempty"`

	// Whether this model can be used as a base for customization. Customized models are not further customizable, and some
	// base models are not customizable.
	Customizable *bool `json:"customizable,omitempty"`

	// Whether or not the model is a default model. A default model is the model for a given language pair that will be
	// used when that language pair is specified in the source and target parameters.
	DefaultModel *bool `json:"default_model,omitempty"`

	// Either an empty string, indicating the model is not a custom model, or the ID of the service instance that created
	// the model.
	Owner *string `json:"owner,omitempty"`

	// Availability of a model.
	Status *string `json:"status,omitempty"`
}

// Constants associated with the TranslationModel.Status property.
// Availability of a model.
const (
	TranslationModel_Status_Available   = "available"
	TranslationModel_Status_Deleted     = "deleted"
	TranslationModel_Status_Dispatching = "dispatching"
	TranslationModel_Status_Error       = "error"
	TranslationModel_Status_Publishing  = "publishing"
	TranslationModel_Status_Queued      = "queued"
	TranslationModel_Status_Trained     = "trained"
	TranslationModel_Status_Training    = "training"
	TranslationModel_Status_Uploaded    = "uploaded"
	TranslationModel_Status_Uploading   = "uploading"
)

// TranslationModels : The response type for listing existing translation models.
type TranslationModels struct {

	// An array of available models.
	Models []TranslationModel `json:"models" validate:"required"`
}

// TranslationResult : TranslationResult struct
type TranslationResult struct {

	// An estimate of the number of words in the input text.
	WordCount *int64 `json:"word_count" validate:"required"`

	// Number of characters in the input text.
	CharacterCount *int64 `json:"character_count" validate:"required"`

	// List of translation output in UTF-8, corresponding to the input text entries.
	Translations []Translation `json:"translations" validate:"required"`
}
