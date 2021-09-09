/**
 * (C) Copyright IBM Corp. 2021.
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

/*
 * IBM OpenAPI SDK Code Generator Version: 3.38.0-07189efd-20210827-205025
 */

// Package languagetranslatorv3 : Operations and models for the LanguageTranslatorV3 service
package languagetranslatorv3

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	common "github.com/watson-developer-cloud/go-sdk/v2/common"
)

// LanguageTranslatorV3 : IBM Watson&trade; Language Translator translates text from one language to another. The
// service offers multiple IBM-provided translation models that you can customize based on your unique terminology and
// language. Use Language Translator to take news from across the globe and present it in your language, communicate
// with your customers in their own language, and more.
//
// API Version: 3.0.0
// See: https://cloud.ibm.com/docs/language-translator
type LanguageTranslatorV3 struct {
	Service *core.BaseService

	// Release date of the version of the API you want to use. Specify dates in YYYY-MM-DD format. The current version is
	// `2018-05-01`.
	Version *string
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.us-south.language-translator.watson.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "language_translator"

// LanguageTranslatorV3Options : Service options
type LanguageTranslatorV3Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator

	// Release date of the version of the API you want to use. Specify dates in YYYY-MM-DD format. The current version is
	// `2018-05-01`.
	Version *string `validate:"required"`
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

	err = core.ValidateStruct(options, "options")
	if err != nil {
		return
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

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "languageTranslator" suitable for processing requests.
func (languageTranslator *LanguageTranslatorV3) Clone() *LanguageTranslatorV3 {
	if core.IsNil(languageTranslator) {
		return nil
	}
	clone := *languageTranslator
	clone.Service = languageTranslator.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (languageTranslator *LanguageTranslatorV3) SetServiceURL(url string) error {
	return languageTranslator.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (languageTranslator *LanguageTranslatorV3) GetServiceURL() string {
	return languageTranslator.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (languageTranslator *LanguageTranslatorV3) SetDefaultHeaders(headers http.Header) {
	languageTranslator.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (languageTranslator *LanguageTranslatorV3) SetEnableGzipCompression(enableGzip bool) {
	languageTranslator.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (languageTranslator *LanguageTranslatorV3) GetEnableGzipCompression() bool {
	return languageTranslator.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (languageTranslator *LanguageTranslatorV3) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	languageTranslator.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (languageTranslator *LanguageTranslatorV3) DisableRetries() {
	languageTranslator.Service.DisableRetries()
}

// DisableSSLVerification bypasses verification of the server's SSL certificate
func (languageTranslator *LanguageTranslatorV3) DisableSSLVerification() {
	languageTranslator.Service.DisableSSLVerification()
}

// ListLanguages : List supported languages
// Lists all supported languages for translation. The method returns an array of supported languages with information
// about each language. Languages are listed in alphabetical order by language code (for example, `af`, `ar`). In
// addition to basic information about each language, the response indicates whether the language is
// `supported_as_source` for translation and `supported_as_target` for translation. It also lists whether the language
// is `identifiable`.
func (languageTranslator *LanguageTranslatorV3) ListLanguages(listLanguagesOptions *ListLanguagesOptions) (result *Languages, response *core.DetailedResponse, err error) {
	return languageTranslator.ListLanguagesWithContext(context.Background(), listLanguagesOptions)
}

// ListLanguagesWithContext is an alternate form of the ListLanguages method which supports a Context parameter
func (languageTranslator *LanguageTranslatorV3) ListLanguagesWithContext(ctx context.Context, listLanguagesOptions *ListLanguagesOptions) (result *Languages, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listLanguagesOptions, "listLanguagesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = languageTranslator.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(languageTranslator.Service.Options.URL, `/v3/languages`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listLanguagesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("language_translator", "V3", "ListLanguages")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*languageTranslator.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = languageTranslator.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalLanguages)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// Translate : Translate
// Translates the input text from the source language to the target language. Specify a model ID that indicates the
// source and target languages, or specify the source and target languages individually. You can omit the source
// language to have the service attempt to detect the language from the input text. If you omit the source language, the
// request must contain sufficient input text for the service to identify the source language.
//
// You can translate a maximum of 50 KB (51,200 bytes) of text with a single request. All input text must be encoded in
// UTF-8 format.
func (languageTranslator *LanguageTranslatorV3) Translate(translateOptions *TranslateOptions) (result *TranslationResult, response *core.DetailedResponse, err error) {
	return languageTranslator.TranslateWithContext(context.Background(), translateOptions)
}

// TranslateWithContext is an alternate form of the Translate method which supports a Context parameter
func (languageTranslator *LanguageTranslatorV3) TranslateWithContext(ctx context.Context, translateOptions *TranslateOptions) (result *TranslationResult, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(translateOptions, "translateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(translateOptions, "translateOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = languageTranslator.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(languageTranslator.Service.Options.URL, `/v3/translate`, nil)
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

	builder.AddQuery("version", fmt.Sprint(*languageTranslator.Version))

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

	var rawResponse map[string]json.RawMessage
	response, err = languageTranslator.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTranslationResult)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListIdentifiableLanguages : List identifiable languages
// Lists the languages that the service can identify. Returns the language code (for example, `en` for English or `es`
// for Spanish) and name of each language.
func (languageTranslator *LanguageTranslatorV3) ListIdentifiableLanguages(listIdentifiableLanguagesOptions *ListIdentifiableLanguagesOptions) (result *IdentifiableLanguages, response *core.DetailedResponse, err error) {
	return languageTranslator.ListIdentifiableLanguagesWithContext(context.Background(), listIdentifiableLanguagesOptions)
}

// ListIdentifiableLanguagesWithContext is an alternate form of the ListIdentifiableLanguages method which supports a Context parameter
func (languageTranslator *LanguageTranslatorV3) ListIdentifiableLanguagesWithContext(ctx context.Context, listIdentifiableLanguagesOptions *ListIdentifiableLanguagesOptions) (result *IdentifiableLanguages, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listIdentifiableLanguagesOptions, "listIdentifiableLanguagesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = languageTranslator.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(languageTranslator.Service.Options.URL, `/v3/identifiable_languages`, nil)
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

	builder.AddQuery("version", fmt.Sprint(*languageTranslator.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = languageTranslator.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalIdentifiableLanguages)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// Identify : Identify language
// Identifies the language of the input text.
func (languageTranslator *LanguageTranslatorV3) Identify(identifyOptions *IdentifyOptions) (result *IdentifiedLanguages, response *core.DetailedResponse, err error) {
	return languageTranslator.IdentifyWithContext(context.Background(), identifyOptions)
}

// IdentifyWithContext is an alternate form of the Identify method which supports a Context parameter
func (languageTranslator *LanguageTranslatorV3) IdentifyWithContext(ctx context.Context, identifyOptions *IdentifyOptions) (result *IdentifiedLanguages, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(identifyOptions, "identifyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(identifyOptions, "identifyOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = languageTranslator.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(languageTranslator.Service.Options.URL, `/v3/identify`, nil)
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

	builder.AddQuery("version", fmt.Sprint(*languageTranslator.Version))

	_, err = builder.SetBodyContent("text/plain", nil, nil, identifyOptions.Text)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = languageTranslator.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalIdentifiedLanguages)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListModels : List models
// Lists available translation models.
func (languageTranslator *LanguageTranslatorV3) ListModels(listModelsOptions *ListModelsOptions) (result *TranslationModels, response *core.DetailedResponse, err error) {
	return languageTranslator.ListModelsWithContext(context.Background(), listModelsOptions)
}

// ListModelsWithContext is an alternate form of the ListModels method which supports a Context parameter
func (languageTranslator *LanguageTranslatorV3) ListModelsWithContext(ctx context.Context, listModelsOptions *ListModelsOptions) (result *TranslationModels, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listModelsOptions, "listModelsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = languageTranslator.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(languageTranslator.Service.Options.URL, `/v3/models`, nil)
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

	builder.AddQuery("version", fmt.Sprint(*languageTranslator.Version))
	if listModelsOptions.Source != nil {
		builder.AddQuery("source", fmt.Sprint(*listModelsOptions.Source))
	}
	if listModelsOptions.Target != nil {
		builder.AddQuery("target", fmt.Sprint(*listModelsOptions.Target))
	}
	if listModelsOptions.Default != nil {
		builder.AddQuery("default", fmt.Sprint(*listModelsOptions.Default))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = languageTranslator.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTranslationModels)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateModel : Create model
// Uploads training files to customize a translation model. You can customize a model with a forced glossary or with a
// parallel corpus:
// * Use a *forced glossary* to force certain terms and phrases to be translated in a specific way. You can upload only
// a single forced glossary file for a model. The size of a forced glossary file for a custom model is limited to 10 MB.
// * Use a *parallel corpus* when you want your custom model to learn from general translation patterns in parallel
// sentences in your samples. What your model learns from a parallel corpus can improve translation results for input
// text that the model has not been trained on. You can upload multiple parallel corpora files with a request. To
// successfully train with parallel corpora, the corpora files must contain a cumulative total of at least 5000 parallel
// sentences. The cumulative size of all uploaded corpus files for a custom model is limited to 250 MB.
//
// Depending on the type of customization and the size of the uploaded files, training time can range from minutes for a
// glossary to several hours for a large parallel corpus. To create a model that is customized with a parallel corpus
// and a forced glossary, customize the model with a parallel corpus first and then customize the resulting model with a
// forced glossary.
//
// You can create a maximum of 10 custom models per language pair. For more information about customizing a translation
// model, including the formatting and character restrictions for data files, see [Customizing your
// model](https://cloud.ibm.com/docs/language-translator?topic=language-translator-customizing).
//
// #### Supported file formats
//
//  You can provide your training data for customization in the following document formats:
// * **TMX** (`.tmx`) - Translation Memory eXchange (TMX) is an XML specification for the exchange of translation
// memories.
// * **XLIFF** (`.xliff`) - XML Localization Interchange File Format (XLIFF) is an XML specification for the exchange of
// translation memories.
// * **CSV** (`.csv`) - Comma-separated values (CSV) file with two columns for aligned sentences and phrases. The first
// row must have two language codes. The first column is for the source language code, and the second column is for the
// target language code.
// * **TSV** (`.tsv` or `.tab`) - Tab-separated values (TSV) file with two columns for aligned sentences and phrases.
// The first row must have two language codes. The first column is for the source language code, and the second column
// is for the target language code.
// * **JSON** (`.json`) - Custom JSON format for specifying aligned sentences and phrases.
// * **Microsoft Excel** (`.xls` or `.xlsx`) - Excel file with the first two columns for aligned sentences and phrases.
// The first row contains the language code.
//
// You must encode all text data in UTF-8 format. For more information, see [Supported document formats for training
// data](https://cloud.ibm.com/docs/language-translator?topic=language-translator-customizing#supported-document-formats-for-training-data).
//
//
// #### Specifying file formats
//
//  You can indicate the format of a file by including the file extension with the file name. Use the file extensions
// shown in **Supported file formats**.
//
// Alternatively, you can omit the file extension and specify one of the following `content-type` specifications for the
// file:
// * **TMX** - `application/x-tmx+xml`
// * **XLIFF** - `application/xliff+xml`
// * **CSV** - `text/csv`
// * **TSV** - `text/tab-separated-values`
// * **JSON** - `application/json`
// * **Microsoft Excel** - `application/vnd.openxmlformats-officedocument.spreadsheetml.sheet`
//
// For example, with `curl`, use the following `content-type` specification to indicate the format of a CSV file named
// **glossary**:
//
// `--form "forced_glossary=@glossary;type=text/csv"`.
func (languageTranslator *LanguageTranslatorV3) CreateModel(createModelOptions *CreateModelOptions) (result *TranslationModel, response *core.DetailedResponse, err error) {
	return languageTranslator.CreateModelWithContext(context.Background(), createModelOptions)
}

// CreateModelWithContext is an alternate form of the CreateModel method which supports a Context parameter
func (languageTranslator *LanguageTranslatorV3) CreateModelWithContext(ctx context.Context, createModelOptions *CreateModelOptions) (result *TranslationModel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createModelOptions, "createModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createModelOptions, "createModelOptions")
	if err != nil {
		return
	}
	if (createModelOptions.ForcedGlossary == nil) && (createModelOptions.ParallelCorpus == nil) {
		err = fmt.Errorf("at least one of forcedGlossary or parallelCorpus must be supplied")
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = languageTranslator.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(languageTranslator.Service.Options.URL, `/v3/models`, nil)
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

	builder.AddQuery("version", fmt.Sprint(*languageTranslator.Version))
	builder.AddQuery("base_model_id", fmt.Sprint(*createModelOptions.BaseModelID))
	if createModelOptions.Name != nil {
		builder.AddQuery("name", fmt.Sprint(*createModelOptions.Name))
	}

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

	var rawResponse map[string]json.RawMessage
	response, err = languageTranslator.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTranslationModel)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteModel : Delete model
// Deletes a custom translation model.
func (languageTranslator *LanguageTranslatorV3) DeleteModel(deleteModelOptions *DeleteModelOptions) (result *DeleteModelResult, response *core.DetailedResponse, err error) {
	return languageTranslator.DeleteModelWithContext(context.Background(), deleteModelOptions)
}

// DeleteModelWithContext is an alternate form of the DeleteModel method which supports a Context parameter
func (languageTranslator *LanguageTranslatorV3) DeleteModelWithContext(ctx context.Context, deleteModelOptions *DeleteModelOptions) (result *DeleteModelResult, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteModelOptions, "deleteModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteModelOptions, "deleteModelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"model_id": *deleteModelOptions.ModelID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = languageTranslator.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(languageTranslator.Service.Options.URL, `/v3/models/{model_id}`, pathParamsMap)
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

	builder.AddQuery("version", fmt.Sprint(*languageTranslator.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = languageTranslator.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeleteModelResult)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetModel : Get model details
// Gets information about a translation model, including training status for custom models. Use this API call to poll
// the status of your customization request. A successfully completed training has a status of `available`.
func (languageTranslator *LanguageTranslatorV3) GetModel(getModelOptions *GetModelOptions) (result *TranslationModel, response *core.DetailedResponse, err error) {
	return languageTranslator.GetModelWithContext(context.Background(), getModelOptions)
}

// GetModelWithContext is an alternate form of the GetModel method which supports a Context parameter
func (languageTranslator *LanguageTranslatorV3) GetModelWithContext(ctx context.Context, getModelOptions *GetModelOptions) (result *TranslationModel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getModelOptions, "getModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getModelOptions, "getModelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"model_id": *getModelOptions.ModelID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = languageTranslator.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(languageTranslator.Service.Options.URL, `/v3/models/{model_id}`, pathParamsMap)
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

	builder.AddQuery("version", fmt.Sprint(*languageTranslator.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = languageTranslator.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTranslationModel)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListDocuments : List documents
// Lists documents that have been submitted for translation.
func (languageTranslator *LanguageTranslatorV3) ListDocuments(listDocumentsOptions *ListDocumentsOptions) (result *DocumentList, response *core.DetailedResponse, err error) {
	return languageTranslator.ListDocumentsWithContext(context.Background(), listDocumentsOptions)
}

// ListDocumentsWithContext is an alternate form of the ListDocuments method which supports a Context parameter
func (languageTranslator *LanguageTranslatorV3) ListDocumentsWithContext(ctx context.Context, listDocumentsOptions *ListDocumentsOptions) (result *DocumentList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listDocumentsOptions, "listDocumentsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = languageTranslator.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(languageTranslator.Service.Options.URL, `/v3/documents`, nil)
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

	builder.AddQuery("version", fmt.Sprint(*languageTranslator.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = languageTranslator.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDocumentList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// TranslateDocument : Translate document
// Submit a document for translation. You can submit the document contents in the `file` parameter, or you can reference
// a previously submitted document by document ID. The maximum file size for document translation is
// * 20 MB for service instances on the Standard, Advanced, and Premium plans
// * 2 MB for service instances on the Lite plan.
func (languageTranslator *LanguageTranslatorV3) TranslateDocument(translateDocumentOptions *TranslateDocumentOptions) (result *DocumentStatus, response *core.DetailedResponse, err error) {
	return languageTranslator.TranslateDocumentWithContext(context.Background(), translateDocumentOptions)
}

// TranslateDocumentWithContext is an alternate form of the TranslateDocument method which supports a Context parameter
func (languageTranslator *LanguageTranslatorV3) TranslateDocumentWithContext(ctx context.Context, translateDocumentOptions *TranslateDocumentOptions) (result *DocumentStatus, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(translateDocumentOptions, "translateDocumentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(translateDocumentOptions, "translateDocumentOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = languageTranslator.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(languageTranslator.Service.Options.URL, `/v3/documents`, nil)
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

	builder.AddQuery("version", fmt.Sprint(*languageTranslator.Version))

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

	var rawResponse map[string]json.RawMessage
	response, err = languageTranslator.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDocumentStatus)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetDocumentStatus : Get document status
// Gets the translation status of a document.
func (languageTranslator *LanguageTranslatorV3) GetDocumentStatus(getDocumentStatusOptions *GetDocumentStatusOptions) (result *DocumentStatus, response *core.DetailedResponse, err error) {
	return languageTranslator.GetDocumentStatusWithContext(context.Background(), getDocumentStatusOptions)
}

// GetDocumentStatusWithContext is an alternate form of the GetDocumentStatus method which supports a Context parameter
func (languageTranslator *LanguageTranslatorV3) GetDocumentStatusWithContext(ctx context.Context, getDocumentStatusOptions *GetDocumentStatusOptions) (result *DocumentStatus, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDocumentStatusOptions, "getDocumentStatusOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getDocumentStatusOptions, "getDocumentStatusOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"document_id": *getDocumentStatusOptions.DocumentID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = languageTranslator.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(languageTranslator.Service.Options.URL, `/v3/documents/{document_id}`, pathParamsMap)
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

	builder.AddQuery("version", fmt.Sprint(*languageTranslator.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = languageTranslator.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDocumentStatus)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteDocument : Delete document
// Deletes a document.
func (languageTranslator *LanguageTranslatorV3) DeleteDocument(deleteDocumentOptions *DeleteDocumentOptions) (response *core.DetailedResponse, err error) {
	return languageTranslator.DeleteDocumentWithContext(context.Background(), deleteDocumentOptions)
}

// DeleteDocumentWithContext is an alternate form of the DeleteDocument method which supports a Context parameter
func (languageTranslator *LanguageTranslatorV3) DeleteDocumentWithContext(ctx context.Context, deleteDocumentOptions *DeleteDocumentOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteDocumentOptions, "deleteDocumentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteDocumentOptions, "deleteDocumentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"document_id": *deleteDocumentOptions.DocumentID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = languageTranslator.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(languageTranslator.Service.Options.URL, `/v3/documents/{document_id}`, pathParamsMap)
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

	builder.AddQuery("version", fmt.Sprint(*languageTranslator.Version))

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
	return languageTranslator.GetTranslatedDocumentWithContext(context.Background(), getTranslatedDocumentOptions)
}

// GetTranslatedDocumentWithContext is an alternate form of the GetTranslatedDocument method which supports a Context parameter
func (languageTranslator *LanguageTranslatorV3) GetTranslatedDocumentWithContext(ctx context.Context, getTranslatedDocumentOptions *GetTranslatedDocumentOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getTranslatedDocumentOptions, "getTranslatedDocumentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getTranslatedDocumentOptions, "getTranslatedDocumentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"document_id": *getTranslatedDocumentOptions.DocumentID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = languageTranslator.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(languageTranslator.Service.Options.URL, `/v3/documents/{document_id}/translated_document`, pathParamsMap)
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

	builder.AddQuery("version", fmt.Sprint(*languageTranslator.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = languageTranslator.Service.Request(request, &result)

	return
}

// CreateModelOptions : The CreateModel options.
type CreateModelOptions struct {
	// The ID of the translation model to use as the base for customization. To see available models and IDs, use the `List
	// models` method. Most models that are provided with the service are customizable. In addition, all models that you
	// create with parallel corpora customization can be further customized with a forced glossary.
	BaseModelID *string `json:"-" validate:"required"`

	// A file with forced glossary terms for the source and target languages. The customizations in the file completely
	// overwrite the domain translation data, including high frequency or high confidence phrase translations.
	//
	// You can upload only one glossary file for a custom model, and the glossary can have a maximum size of 10 MB. A
	// forced glossary must contain single words or short phrases. For more information, see **Supported file formats** in
	// the method description.
	//
	// *With `curl`, use `--form forced_glossary=@{filename}`.*.
	ForcedGlossary io.ReadCloser `json:"-"`

	// A file with parallel sentences for the source and target languages. You can upload multiple parallel corpus files in
	// one request by repeating the parameter. All uploaded parallel corpus files combined must contain at least 5000
	// parallel sentences to train successfully. You can provide a maximum of 500,000 parallel sentences across all
	// corpora.
	//
	// A single entry in a corpus file can contain a maximum of 80 words. All corpora files for a custom model can have a
	// cumulative maximum size of 250 MB. For more information, see **Supported file formats** in the method description.
	//
	// *With `curl`, use `--form parallel_corpus=@{filename}`.*.
	ParallelCorpus io.ReadCloser `json:"-"`

	// An optional model name that you can use to identify the model. Valid characters are letters, numbers, dashes,
	// underscores, spaces, and apostrophes. The maximum length of the name is 32 characters.
	Name *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateModelOptions : Instantiate CreateModelOptions
func (*LanguageTranslatorV3) NewCreateModelOptions(baseModelID string) *CreateModelOptions {
	return &CreateModelOptions{
		BaseModelID: core.StringPtr(baseModelID),
	}
}

// SetBaseModelID : Allow user to set BaseModelID
func (_options *CreateModelOptions) SetBaseModelID(baseModelID string) *CreateModelOptions {
	_options.BaseModelID = core.StringPtr(baseModelID)
	return _options
}

// SetForcedGlossary : Allow user to set ForcedGlossary
func (_options *CreateModelOptions) SetForcedGlossary(forcedGlossary io.ReadCloser) *CreateModelOptions {
	_options.ForcedGlossary = forcedGlossary
	return _options
}

// SetParallelCorpus : Allow user to set ParallelCorpus
func (_options *CreateModelOptions) SetParallelCorpus(parallelCorpus io.ReadCloser) *CreateModelOptions {
	_options.ParallelCorpus = parallelCorpus
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateModelOptions) SetName(name string) *CreateModelOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateModelOptions) SetHeaders(param map[string]string) *CreateModelOptions {
	options.Headers = param
	return options
}

// DeleteDocumentOptions : The DeleteDocument options.
type DeleteDocumentOptions struct {
	// Document ID of the document to delete.
	DocumentID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteDocumentOptions : Instantiate DeleteDocumentOptions
func (*LanguageTranslatorV3) NewDeleteDocumentOptions(documentID string) *DeleteDocumentOptions {
	return &DeleteDocumentOptions{
		DocumentID: core.StringPtr(documentID),
	}
}

// SetDocumentID : Allow user to set DocumentID
func (_options *DeleteDocumentOptions) SetDocumentID(documentID string) *DeleteDocumentOptions {
	_options.DocumentID = core.StringPtr(documentID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteDocumentOptions) SetHeaders(param map[string]string) *DeleteDocumentOptions {
	options.Headers = param
	return options
}

// DeleteModelOptions : The DeleteModel options.
type DeleteModelOptions struct {
	// Model ID of the model to delete.
	ModelID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteModelOptions : Instantiate DeleteModelOptions
func (*LanguageTranslatorV3) NewDeleteModelOptions(modelID string) *DeleteModelOptions {
	return &DeleteModelOptions{
		ModelID: core.StringPtr(modelID),
	}
}

// SetModelID : Allow user to set ModelID
func (_options *DeleteModelOptions) SetModelID(modelID string) *DeleteModelOptions {
	_options.ModelID = core.StringPtr(modelID)
	return _options
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

// UnmarshalDeleteModelResult unmarshals an instance of DeleteModelResult from the specified map of raw messages.
func UnmarshalDeleteModelResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteModelResult)
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DocumentList : DocumentList struct
type DocumentList struct {
	// An array of all previously submitted documents.
	Documents []DocumentStatus `json:"documents" validate:"required"`
}

// UnmarshalDocumentList unmarshals an instance of DocumentList from the specified map of raw messages.
func UnmarshalDocumentList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DocumentList)
	err = core.UnmarshalModel(m, "documents", &obj.Documents, UnmarshalDocumentStatus)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
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

	// A score between 0 and 1 indicating the confidence of source language detection. A higher value indicates greater
	// confidence. This is returned only when the service automatically detects the source language.
	DetectedLanguageConfidence *float64 `json:"detected_language_confidence,omitempty"`

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
	DocumentStatusStatusAvailableConst  = "available"
	DocumentStatusStatusFailedConst     = "failed"
	DocumentStatusStatusProcessingConst = "processing"
)

// UnmarshalDocumentStatus unmarshals an instance of DocumentStatus from the specified map of raw messages.
func UnmarshalDocumentStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DocumentStatus)
	err = core.UnmarshalPrimitive(m, "document_id", &obj.DocumentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "filename", &obj.Filename)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "model_id", &obj.ModelID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "base_model_id", &obj.BaseModelID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source", &obj.Source)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "detected_language_confidence", &obj.DetectedLanguageConfidence)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "target", &obj.Target)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created", &obj.Created)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "completed", &obj.Completed)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "word_count", &obj.WordCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "character_count", &obj.CharacterCount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetDocumentStatusOptions : The GetDocumentStatus options.
type GetDocumentStatusOptions struct {
	// The document ID of the document.
	DocumentID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetDocumentStatusOptions : Instantiate GetDocumentStatusOptions
func (*LanguageTranslatorV3) NewGetDocumentStatusOptions(documentID string) *GetDocumentStatusOptions {
	return &GetDocumentStatusOptions{
		DocumentID: core.StringPtr(documentID),
	}
}

// SetDocumentID : Allow user to set DocumentID
func (_options *GetDocumentStatusOptions) SetDocumentID(documentID string) *GetDocumentStatusOptions {
	_options.DocumentID = core.StringPtr(documentID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDocumentStatusOptions) SetHeaders(param map[string]string) *GetDocumentStatusOptions {
	options.Headers = param
	return options
}

// GetModelOptions : The GetModel options.
type GetModelOptions struct {
	// Model ID of the model to get.
	ModelID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetModelOptions : Instantiate GetModelOptions
func (*LanguageTranslatorV3) NewGetModelOptions(modelID string) *GetModelOptions {
	return &GetModelOptions{
		ModelID: core.StringPtr(modelID),
	}
}

// SetModelID : Allow user to set ModelID
func (_options *GetModelOptions) SetModelID(modelID string) *GetModelOptions {
	_options.ModelID = core.StringPtr(modelID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetModelOptions) SetHeaders(param map[string]string) *GetModelOptions {
	options.Headers = param
	return options
}

// GetTranslatedDocumentOptions : The GetTranslatedDocument options.
type GetTranslatedDocumentOptions struct {
	// The document ID of the document that was submitted for translation.
	DocumentID *string `json:"-" validate:"required,ne="`

	// The type of the response: application/powerpoint, application/mspowerpoint, application/x-rtf, application/json,
	// application/xml, application/vnd.ms-excel, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet,
	// application/vnd.ms-powerpoint, application/vnd.openxmlformats-officedocument.presentationml.presentation,
	// application/msword, application/vnd.openxmlformats-officedocument.wordprocessingml.document,
	// application/vnd.oasis.opendocument.spreadsheet, application/vnd.oasis.opendocument.presentation,
	// application/vnd.oasis.opendocument.text, application/pdf, application/rtf, text/html, text/json, text/plain,
	// text/richtext, text/rtf, or text/xml. A character encoding can be specified by including a `charset` parameter. For
	// example, 'text/html;charset=utf-8'.
	Accept *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetTranslatedDocumentOptions : Instantiate GetTranslatedDocumentOptions
func (*LanguageTranslatorV3) NewGetTranslatedDocumentOptions(documentID string) *GetTranslatedDocumentOptions {
	return &GetTranslatedDocumentOptions{
		DocumentID: core.StringPtr(documentID),
	}
}

// SetDocumentID : Allow user to set DocumentID
func (_options *GetTranslatedDocumentOptions) SetDocumentID(documentID string) *GetTranslatedDocumentOptions {
	_options.DocumentID = core.StringPtr(documentID)
	return _options
}

// SetAccept : Allow user to set Accept
func (_options *GetTranslatedDocumentOptions) SetAccept(accept string) *GetTranslatedDocumentOptions {
	_options.Accept = core.StringPtr(accept)
	return _options
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

// UnmarshalIdentifiableLanguage unmarshals an instance of IdentifiableLanguage from the specified map of raw messages.
func UnmarshalIdentifiableLanguage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(IdentifiableLanguage)
	err = core.UnmarshalPrimitive(m, "language", &obj.Language)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// IdentifiableLanguages : IdentifiableLanguages struct
type IdentifiableLanguages struct {
	// A list of all languages that the service can identify.
	Languages []IdentifiableLanguage `json:"languages" validate:"required"`
}

// UnmarshalIdentifiableLanguages unmarshals an instance of IdentifiableLanguages from the specified map of raw messages.
func UnmarshalIdentifiableLanguages(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(IdentifiableLanguages)
	err = core.UnmarshalModel(m, "languages", &obj.Languages, UnmarshalIdentifiableLanguage)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// IdentifiedLanguage : IdentifiedLanguage struct
type IdentifiedLanguage struct {
	// The language code for an identified language.
	Language *string `json:"language" validate:"required"`

	// The confidence score for the identified language.
	Confidence *float64 `json:"confidence" validate:"required"`
}

// UnmarshalIdentifiedLanguage unmarshals an instance of IdentifiedLanguage from the specified map of raw messages.
func UnmarshalIdentifiedLanguage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(IdentifiedLanguage)
	err = core.UnmarshalPrimitive(m, "language", &obj.Language)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "confidence", &obj.Confidence)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// IdentifiedLanguages : IdentifiedLanguages struct
type IdentifiedLanguages struct {
	// A ranking of identified languages with confidence scores.
	Languages []IdentifiedLanguage `json:"languages" validate:"required"`
}

// UnmarshalIdentifiedLanguages unmarshals an instance of IdentifiedLanguages from the specified map of raw messages.
func UnmarshalIdentifiedLanguages(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(IdentifiedLanguages)
	err = core.UnmarshalModel(m, "languages", &obj.Languages, UnmarshalIdentifiedLanguage)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// IdentifyOptions : The Identify options.
type IdentifyOptions struct {
	// Input text in UTF-8 format.
	Text *string `json:"text" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewIdentifyOptions : Instantiate IdentifyOptions
func (*LanguageTranslatorV3) NewIdentifyOptions(text string) *IdentifyOptions {
	return &IdentifyOptions{
		Text: core.StringPtr(text),
	}
}

// SetText : Allow user to set Text
func (_options *IdentifyOptions) SetText(text string) *IdentifyOptions {
	_options.Text = core.StringPtr(text)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *IdentifyOptions) SetHeaders(param map[string]string) *IdentifyOptions {
	options.Headers = param
	return options
}

// Language : Response payload for languages.
type Language struct {
	// The language code for the language (for example, `af`).
	Language *string `json:"language,omitempty"`

	// The name of the language in English (for example, `Afrikaans`).
	LanguageName *string `json:"language_name,omitempty"`

	// The native name of the language (for example, `Afrikaans`).
	NativeLanguageName *string `json:"native_language_name,omitempty"`

	// The country code for the language (for example, `ZA` for South Africa).
	CountryCode *string `json:"country_code,omitempty"`

	// Indicates whether words of the language are separated by whitespace: `true` if the words are separated; `false`
	// otherwise.
	WordsSeparated *bool `json:"words_separated,omitempty"`

	// Indicates the direction of the language: `right_to_left` or `left_to_right`.
	Direction *string `json:"direction,omitempty"`

	// Indicates whether the language can be used as the source for translation: `true` if the language can be used as the
	// source; `false` otherwise.
	SupportedAsSource *bool `json:"supported_as_source,omitempty"`

	// Indicates whether the language can be used as the target for translation: `true` if the language can be used as the
	// target; `false` otherwise.
	SupportedAsTarget *bool `json:"supported_as_target,omitempty"`

	// Indicates whether the language supports automatic detection: `true` if the language can be detected automatically;
	// `false` otherwise.
	Identifiable *bool `json:"identifiable,omitempty"`
}

// UnmarshalLanguage unmarshals an instance of Language from the specified map of raw messages.
func UnmarshalLanguage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Language)
	err = core.UnmarshalPrimitive(m, "language", &obj.Language)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "language_name", &obj.LanguageName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "native_language_name", &obj.NativeLanguageName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "country_code", &obj.CountryCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "words_separated", &obj.WordsSeparated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "direction", &obj.Direction)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "supported_as_source", &obj.SupportedAsSource)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "supported_as_target", &obj.SupportedAsTarget)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "identifiable", &obj.Identifiable)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Languages : The response type for listing supported languages.
type Languages struct {
	// An array of supported languages with information about each language.
	Languages []Language `json:"languages" validate:"required"`
}

// UnmarshalLanguages unmarshals an instance of Languages from the specified map of raw messages.
func UnmarshalLanguages(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Languages)
	err = core.UnmarshalModel(m, "languages", &obj.Languages, UnmarshalLanguage)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListDocumentsOptions : The ListDocuments options.
type ListDocumentsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListDocumentsOptions : Instantiate ListDocumentsOptions
func (*LanguageTranslatorV3) NewListDocumentsOptions() *ListDocumentsOptions {
	return &ListDocumentsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListDocumentsOptions) SetHeaders(param map[string]string) *ListDocumentsOptions {
	options.Headers = param
	return options
}

// ListIdentifiableLanguagesOptions : The ListIdentifiableLanguages options.
type ListIdentifiableLanguagesOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListIdentifiableLanguagesOptions : Instantiate ListIdentifiableLanguagesOptions
func (*LanguageTranslatorV3) NewListIdentifiableLanguagesOptions() *ListIdentifiableLanguagesOptions {
	return &ListIdentifiableLanguagesOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListIdentifiableLanguagesOptions) SetHeaders(param map[string]string) *ListIdentifiableLanguagesOptions {
	options.Headers = param
	return options
}

// ListLanguagesOptions : The ListLanguages options.
type ListLanguagesOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListLanguagesOptions : Instantiate ListLanguagesOptions
func (*LanguageTranslatorV3) NewListLanguagesOptions() *ListLanguagesOptions {
	return &ListLanguagesOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListLanguagesOptions) SetHeaders(param map[string]string) *ListLanguagesOptions {
	options.Headers = param
	return options
}

// ListModelsOptions : The ListModels options.
type ListModelsOptions struct {
	// Specify a language code to filter results by source language.
	Source *string `json:"-"`

	// Specify a language code to filter results by target language.
	Target *string `json:"-"`

	// If the `default` parameter isn't specified, the service returns all models (default and non-default) for each
	// language pair. To return only default models, set this parameter to `true`. To return only non-default models, set
	// this parameter to `false`. There is exactly one default model, the IBM-provided base model, per language pair.
	Default *bool `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListModelsOptions : Instantiate ListModelsOptions
func (*LanguageTranslatorV3) NewListModelsOptions() *ListModelsOptions {
	return &ListModelsOptions{}
}

// SetSource : Allow user to set Source
func (_options *ListModelsOptions) SetSource(source string) *ListModelsOptions {
	_options.Source = core.StringPtr(source)
	return _options
}

// SetTarget : Allow user to set Target
func (_options *ListModelsOptions) SetTarget(target string) *ListModelsOptions {
	_options.Target = core.StringPtr(target)
	return _options
}

// SetDefault : Allow user to set Default
func (_options *ListModelsOptions) SetDefault(defaultVar bool) *ListModelsOptions {
	_options.Default = core.BoolPtr(defaultVar)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListModelsOptions) SetHeaders(param map[string]string) *ListModelsOptions {
	options.Headers = param
	return options
}

// TranslateDocumentOptions : The TranslateDocument options.
type TranslateDocumentOptions struct {
	// The contents of the source file to translate. The maximum file size for document translation is 20 MB for service
	// instances on the Standard, Advanced, and Premium plans, and 2 MB for service instances on the Lite plan. For more
	// information, see [Supported file formats
	// (Beta)](https://cloud.ibm.com/docs/language-translator?topic=language-translator-document-translator-tutorial#supported-file-formats).
	File io.ReadCloser `json:"-" validate:"required"`

	// The filename for file.
	Filename *string `json:"-" validate:"required"`

	// The content type of file.
	FileContentType *string `json:"-"`

	// The model to use for translation. For example, `en-de` selects the IBM-provided base model for English-to-German
	// translation. A model ID overrides the `source` and `target` parameters and is required if you use a custom model. If
	// no model ID is specified, you must specify at least a target language.
	ModelID *string `json:"-"`

	// Language code that specifies the language of the source document. If omitted, the service derives the source
	// language from the input text. The input must contain sufficient text for the service to identify the language
	// reliably.
	Source *string `json:"-"`

	// Language code that specifies the target language for translation. Required if model ID is not specified.
	Target *string `json:"-"`

	// To use a previously submitted document as the source for a new translation, enter the `document_id` of the document.
	DocumentID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewTranslateDocumentOptions : Instantiate TranslateDocumentOptions
func (*LanguageTranslatorV3) NewTranslateDocumentOptions(file io.ReadCloser, filename string) *TranslateDocumentOptions {
	return &TranslateDocumentOptions{
		File:     file,
		Filename: core.StringPtr(filename),
	}
}

// SetFile : Allow user to set File
func (_options *TranslateDocumentOptions) SetFile(file io.ReadCloser) *TranslateDocumentOptions {
	_options.File = file
	return _options
}

// SetFilename : Allow user to set Filename
func (_options *TranslateDocumentOptions) SetFilename(filename string) *TranslateDocumentOptions {
	_options.Filename = core.StringPtr(filename)
	return _options
}

// SetFileContentType : Allow user to set FileContentType
func (_options *TranslateDocumentOptions) SetFileContentType(fileContentType string) *TranslateDocumentOptions {
	_options.FileContentType = core.StringPtr(fileContentType)
	return _options
}

// SetModelID : Allow user to set ModelID
func (_options *TranslateDocumentOptions) SetModelID(modelID string) *TranslateDocumentOptions {
	_options.ModelID = core.StringPtr(modelID)
	return _options
}

// SetSource : Allow user to set Source
func (_options *TranslateDocumentOptions) SetSource(source string) *TranslateDocumentOptions {
	_options.Source = core.StringPtr(source)
	return _options
}

// SetTarget : Allow user to set Target
func (_options *TranslateDocumentOptions) SetTarget(target string) *TranslateDocumentOptions {
	_options.Target = core.StringPtr(target)
	return _options
}

// SetDocumentID : Allow user to set DocumentID
func (_options *TranslateDocumentOptions) SetDocumentID(documentID string) *TranslateDocumentOptions {
	_options.DocumentID = core.StringPtr(documentID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *TranslateDocumentOptions) SetHeaders(param map[string]string) *TranslateDocumentOptions {
	options.Headers = param
	return options
}

// TranslateOptions : The Translate options.
type TranslateOptions struct {
	// Input text in UTF-8 encoding. Submit a maximum of 50 KB (51,200 bytes) of text with a single request. Multiple
	// elements result in multiple translations in the response.
	Text []string `json:"text" validate:"required"`

	// The model to use for translation. For example, `en-de` selects the IBM-provided base model for English-to-German
	// translation. A model ID overrides the `source` and `target` parameters and is required if you use a custom model. If
	// no model ID is specified, you must specify at least a target language.
	ModelID *string `json:"model_id,omitempty"`

	// Language code that specifies the language of the input text. If omitted, the service derives the source language
	// from the input text. The input must contain sufficient text for the service to identify the language reliably.
	Source *string `json:"source,omitempty"`

	// Language code that specifies the target language for translation. Required if model ID is not specified.
	Target *string `json:"target,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewTranslateOptions : Instantiate TranslateOptions
func (*LanguageTranslatorV3) NewTranslateOptions(text []string) *TranslateOptions {
	return &TranslateOptions{
		Text: text,
	}
}

// SetText : Allow user to set Text
func (_options *TranslateOptions) SetText(text []string) *TranslateOptions {
	_options.Text = text
	return _options
}

// SetModelID : Allow user to set ModelID
func (_options *TranslateOptions) SetModelID(modelID string) *TranslateOptions {
	_options.ModelID = core.StringPtr(modelID)
	return _options
}

// SetSource : Allow user to set Source
func (_options *TranslateOptions) SetSource(source string) *TranslateOptions {
	_options.Source = core.StringPtr(source)
	return _options
}

// SetTarget : Allow user to set Target
func (_options *TranslateOptions) SetTarget(target string) *TranslateOptions {
	_options.Target = core.StringPtr(target)
	return _options
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

// UnmarshalTranslation unmarshals an instance of Translation from the specified map of raw messages.
func UnmarshalTranslation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Translation)
	err = core.UnmarshalPrimitive(m, "translation", &obj.Translation)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
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
	TranslationModelStatusAvailableConst   = "available"
	TranslationModelStatusDeletedConst     = "deleted"
	TranslationModelStatusDispatchingConst = "dispatching"
	TranslationModelStatusErrorConst       = "error"
	TranslationModelStatusPublishingConst  = "publishing"
	TranslationModelStatusQueuedConst      = "queued"
	TranslationModelStatusTrainedConst     = "trained"
	TranslationModelStatusTrainingConst    = "training"
	TranslationModelStatusUploadedConst    = "uploaded"
	TranslationModelStatusUploadingConst   = "uploading"
)

// UnmarshalTranslationModel unmarshals an instance of TranslationModel from the specified map of raw messages.
func UnmarshalTranslationModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TranslationModel)
	err = core.UnmarshalPrimitive(m, "model_id", &obj.ModelID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source", &obj.Source)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "target", &obj.Target)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "base_model_id", &obj.BaseModelID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "domain", &obj.Domain)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "customizable", &obj.Customizable)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "default_model", &obj.DefaultModel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "owner", &obj.Owner)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TranslationModels : The response type for listing existing translation models.
type TranslationModels struct {
	// An array of available models.
	Models []TranslationModel `json:"models" validate:"required"`
}

// UnmarshalTranslationModels unmarshals an instance of TranslationModels from the specified map of raw messages.
func UnmarshalTranslationModels(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TranslationModels)
	err = core.UnmarshalModel(m, "models", &obj.Models, UnmarshalTranslationModel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TranslationResult : TranslationResult struct
type TranslationResult struct {
	// An estimate of the number of words in the input text.
	WordCount *int64 `json:"word_count" validate:"required"`

	// Number of characters in the input text.
	CharacterCount *int64 `json:"character_count" validate:"required"`

	// The language code of the source text if the source language was automatically detected.
	DetectedLanguage *string `json:"detected_language,omitempty"`

	// A score between 0 and 1 indicating the confidence of source language detection. A higher value indicates greater
	// confidence. This is returned only when the service automatically detects the source language.
	DetectedLanguageConfidence *float64 `json:"detected_language_confidence,omitempty"`

	// List of translation output in UTF-8, corresponding to the input text entries.
	Translations []Translation `json:"translations" validate:"required"`
}

// UnmarshalTranslationResult unmarshals an instance of TranslationResult from the specified map of raw messages.
func UnmarshalTranslationResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TranslationResult)
	err = core.UnmarshalPrimitive(m, "word_count", &obj.WordCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "character_count", &obj.CharacterCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "detected_language", &obj.DetectedLanguage)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "detected_language_confidence", &obj.DetectedLanguageConfidence)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "translations", &obj.Translations, UnmarshalTranslation)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
