// Package languagetranslatorv3 : Operations and models for the LanguageTranslatorV3 service
package languagetranslatorv3

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
	"fmt"
	"github.com/IBM/go-sdk-core/core"
	common "github.com/watson-developer-cloud/go-sdk/common"
	"os"
)

// LanguageTranslatorV3 : IBM Watson&trade; Language Translator translates text from one language to another. The
// service offers multiple IBM provided translation models that you can customize based on your unique terminology and
// language. Use Language Translator to take news from across the globe and present it in your language, communicate
// with your customers in their own language, and more.
//
// Version: V3
// See: http://www.ibm.com/watson/developercloud/language-translator.html
type LanguageTranslatorV3 struct {
	Service *core.BaseService
}

// LanguageTranslatorV3Options : Service options
type LanguageTranslatorV3Options struct {
	Version        string
	URL            string
	Username       string
	Password       string
	IAMApiKey      string
	IAMAccessToken string
	IAMURL         string
}

// NewLanguageTranslatorV3 : Instantiate LanguageTranslatorV3
func NewLanguageTranslatorV3(options *LanguageTranslatorV3Options) (*LanguageTranslatorV3, error) {
	if options.URL == "" {
		options.URL = "https://gateway.watsonplatform.net/language-translator/api"
	}

	serviceOptions := &core.ServiceOptions{
		URL:            options.URL,
		Version:        options.Version,
		Username:       options.Username,
		Password:       options.Password,
		IAMApiKey:      options.IAMApiKey,
		IAMAccessToken: options.IAMAccessToken,
		IAMURL:         options.IAMURL,
	}
	service, serviceErr := core.NewBaseService(serviceOptions, "language_translator", "Language Translator")
	if serviceErr != nil {
		return nil, serviceErr
	}

	return &LanguageTranslatorV3{Service: service}, nil
}

// Translate : Translate
// Translates the input text from the source language to the target language.
func (languageTranslator *LanguageTranslatorV3) Translate(translateOptions *TranslateOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(translateOptions, "translateOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(translateOptions, "translateOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v3/translate"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(languageTranslator.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range translateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("language_translator", "V3", "Translate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", languageTranslator.Service.Options.Version)

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
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := languageTranslator.Service.Request(request, new(TranslationResult))
	return response, err
}

// GetTranslateResult : Retrieve result of Translate operation
func (languageTranslator *LanguageTranslatorV3) GetTranslateResult(response *core.DetailedResponse) *TranslationResult {
	result, ok := response.Result.(*TranslationResult)
	if ok {
		return result
	}
	return nil
}

// Identify : Identify language
// Identifies the language of the input text.
func (languageTranslator *LanguageTranslatorV3) Identify(identifyOptions *IdentifyOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(identifyOptions, "identifyOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(identifyOptions, "identifyOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v3/identify"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(languageTranslator.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range identifyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("language_translator", "V3", "Identify")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "text/plain")
	builder.AddQuery("version", languageTranslator.Service.Options.Version)

	_, err := builder.SetBodyContent("text/plain", nil, nil, identifyOptions.Text)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := languageTranslator.Service.Request(request, new(IdentifiedLanguages))
	return response, err
}

// GetIdentifyResult : Retrieve result of Identify operation
func (languageTranslator *LanguageTranslatorV3) GetIdentifyResult(response *core.DetailedResponse) *IdentifiedLanguages {
	result, ok := response.Result.(*IdentifiedLanguages)
	if ok {
		return result
	}
	return nil
}

// ListIdentifiableLanguages : List identifiable languages
// Lists the languages that the service can identify. Returns the language code (for example, `en` for English or `es`
// for Spanish) and name of each language.
func (languageTranslator *LanguageTranslatorV3) ListIdentifiableLanguages(listIdentifiableLanguagesOptions *ListIdentifiableLanguagesOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateStruct(listIdentifiableLanguagesOptions, "listIdentifiableLanguagesOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v3/identifiable_languages"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(languageTranslator.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range listIdentifiableLanguagesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("language_translator", "V3", "ListIdentifiableLanguages")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", languageTranslator.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := languageTranslator.Service.Request(request, new(IdentifiableLanguages))
	return response, err
}

// GetListIdentifiableLanguagesResult : Retrieve result of ListIdentifiableLanguages operation
func (languageTranslator *LanguageTranslatorV3) GetListIdentifiableLanguagesResult(response *core.DetailedResponse) *IdentifiableLanguages {
	result, ok := response.Result.(*IdentifiableLanguages)
	if ok {
		return result
	}
	return nil
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
// You can have a <b>maxium of 10 custom models per language pair</b>.
func (languageTranslator *LanguageTranslatorV3) CreateModel(createModelOptions *CreateModelOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(createModelOptions, "createModelOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(createModelOptions, "createModelOptions"); err != nil {
		return nil, err
	}
	if (createModelOptions.ForcedGlossary == nil) && (createModelOptions.ParallelCorpus == nil) {
		return nil, fmt.Errorf("At least one of forcedGlossary or parallelCorpus must be supplied")
	}

	pathSegments := []string{"v3/models"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(languageTranslator.Service.Options.URL, pathSegments, pathParameters)

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
	builder.AddQuery("version", languageTranslator.Service.Options.Version)

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
		return nil, err
	}

	response, err := languageTranslator.Service.Request(request, new(TranslationModel))
	return response, err
}

// GetCreateModelResult : Retrieve result of CreateModel operation
func (languageTranslator *LanguageTranslatorV3) GetCreateModelResult(response *core.DetailedResponse) *TranslationModel {
	result, ok := response.Result.(*TranslationModel)
	if ok {
		return result
	}
	return nil
}

// DeleteModel : Delete model
// Deletes a custom translation model.
func (languageTranslator *LanguageTranslatorV3) DeleteModel(deleteModelOptions *DeleteModelOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(deleteModelOptions, "deleteModelOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(deleteModelOptions, "deleteModelOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v3/models"}
	pathParameters := []string{*deleteModelOptions.ModelID}

	builder := core.NewRequestBuilder(core.DELETE)
	builder.ConstructHTTPURL(languageTranslator.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("language_translator", "V3", "DeleteModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", languageTranslator.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := languageTranslator.Service.Request(request, new(DeleteModelResult))
	return response, err
}

// GetDeleteModelResult : Retrieve result of DeleteModel operation
func (languageTranslator *LanguageTranslatorV3) GetDeleteModelResult(response *core.DetailedResponse) *DeleteModelResult {
	result, ok := response.Result.(*DeleteModelResult)
	if ok {
		return result
	}
	return nil
}

// GetModel : Get model details
// Gets information about a translation model, including training status for custom models. Use this API call to poll
// the status of your customization request. A successfully completed training will have a status of `available`.
func (languageTranslator *LanguageTranslatorV3) GetModel(getModelOptions *GetModelOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(getModelOptions, "getModelOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(getModelOptions, "getModelOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v3/models"}
	pathParameters := []string{*getModelOptions.ModelID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(languageTranslator.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range getModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("language_translator", "V3", "GetModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", languageTranslator.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := languageTranslator.Service.Request(request, new(TranslationModel))
	return response, err
}

// GetGetModelResult : Retrieve result of GetModel operation
func (languageTranslator *LanguageTranslatorV3) GetGetModelResult(response *core.DetailedResponse) *TranslationModel {
	result, ok := response.Result.(*TranslationModel)
	if ok {
		return result
	}
	return nil
}

// ListModels : List models
// Lists available translation models.
func (languageTranslator *LanguageTranslatorV3) ListModels(listModelsOptions *ListModelsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateStruct(listModelsOptions, "listModelsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v3/models"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(languageTranslator.Service.Options.URL, pathSegments, pathParameters)

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
	if listModelsOptions.DefaultModels != nil {
		builder.AddQuery("default", fmt.Sprint(*listModelsOptions.DefaultModels))
	}
	builder.AddQuery("version", languageTranslator.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := languageTranslator.Service.Request(request, new(TranslationModels))
	return response, err
}

// GetListModelsResult : Retrieve result of ListModels operation
func (languageTranslator *LanguageTranslatorV3) GetListModelsResult(response *core.DetailedResponse) *TranslationModels {
	result, ok := response.Result.(*TranslationModels)
	if ok {
		return result
	}
	return nil
}

// CreateModelOptions : The createModel options.
type CreateModelOptions struct {

	// The model ID of the model to use as the base for customization. To see available models, use the `List models`
	// method. Usually all IBM provided models are customizable. In addition, all your models that have been created via
	// parallel corpus customization, can be further customized with a forced glossary.
	BaseModelID *string `json:"base_model_id" validate:"required"`

	// A TMX file with your customizations. The customizations in the file completely overwrite the domain translaton data,
	// including high frequency or high confidence phrase translations. You can upload only one glossary with a file size
	// less than 10 MB per call. A forced glossary should contain single words or short phrases.
	ForcedGlossary *os.File `json:"forced_glossary,omitempty"`

	// A TMX file with parallel sentences for source and target language. You can upload multiple parallel_corpus files in
	// one request. All uploaded parallel_corpus files combined, your parallel corpus must contain at least 5,000 parallel
	// sentences to train successfully.
	ParallelCorpus *os.File `json:"parallel_corpus,omitempty"`

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
func (options *CreateModelOptions) SetForcedGlossary(forcedGlossary *os.File) *CreateModelOptions {
	options.ForcedGlossary = forcedGlossary
	return options
}

// SetParallelCorpus : Allow user to set ParallelCorpus
func (options *CreateModelOptions) SetParallelCorpus(parallelCorpus *os.File) *CreateModelOptions {
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

// DeleteModelOptions : The deleteModel options.
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

// GetModelOptions : The getModel options.
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

// IdentifyOptions : The identify options.
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

// ListIdentifiableLanguagesOptions : The listIdentifiableLanguages options.
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

// ListModelsOptions : The listModels options.
type ListModelsOptions struct {

	// Specify a language code to filter results by source language.
	Source *string `json:"source,omitempty"`

	// Specify a language code to filter results by target language.
	Target *string `json:"target,omitempty"`

	// If the default parameter isn't specified, the service will return all models (default and non-default) for each
	// language pair. To return only default models, set this to `true`. To return only non-default models, set this to
	// `false`. There is exactly one default model per language pair, the IBM provided base model.
	DefaultModels *bool `json:"default,omitempty"`

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

// SetDefaultModels : Allow user to set DefaultModels
func (options *ListModelsOptions) SetDefaultModels(defaultModels bool) *ListModelsOptions {
	options.DefaultModels = core.BoolPtr(defaultModels)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListModelsOptions) SetHeaders(param map[string]string) *ListModelsOptions {
	options.Headers = param
	return options
}

// TranslateOptions : The translate options.
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
	TranslationOutput *string `json:"translation" validate:"required"`
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

	// Number of words in the input text.
	WordCount *int64 `json:"word_count" validate:"required"`

	// Number of characters in the input text.
	CharacterCount *int64 `json:"character_count" validate:"required"`

	// List of translation output in UTF-8, corresponding to the input text entries.
	Translations []Translation `json:"translations" validate:"required"`
}
