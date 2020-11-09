/**
 * (C) Copyright IBM Corp. 2020.
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
 * IBM OpenAPI SDK Code Generator Version: 99-SNAPSHOT-9dacd99b-20201204-091925
 */
 

// Package naturallanguageunderstandingv1 : Operations and models for the NaturalLanguageUnderstandingV1 service
package naturallanguageunderstandingv1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	common "github.com/watson-developer-cloud/go-sdk/common"
	"net/http"
	"reflect"
	"time"
)

// NaturalLanguageUnderstandingV1 : Analyze various features of text content at scale. Provide text, raw HTML, or a
// public URL and IBM Watson Natural Language Understanding will give you results for the features you request. The
// service cleans HTML content before analysis by default, so the results can ignore most advertisements and other
// unwanted content.
//
// You can create [custom
// models](https://cloud.ibm.com/docs/natural-language-understanding?topic=natural-language-understanding-customizing)
// with Watson Knowledge Studio to detect custom entities and relations in Natural Language Understanding.
//
// Version: 1.0
// See: https://cloud.ibm.com/docs/natural-language-understanding
type NaturalLanguageUnderstandingV1 struct {
	Service *core.BaseService

	// Release date of the API version you want to use. Specify dates in YYYY-MM-DD format. The current version is
	// `2020-08-01`.
	Version *string
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.us-south.natural-language-understanding.watson.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "natural-language-understanding"

// NaturalLanguageUnderstandingV1Options : Service options
type NaturalLanguageUnderstandingV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator

	// Release date of the API version you want to use. Specify dates in YYYY-MM-DD format. The current version is
	// `2020-08-01`.
	Version *string `validate:"required"`
}

// NewNaturalLanguageUnderstandingV1 : constructs an instance of NaturalLanguageUnderstandingV1 with passed in options.
func NewNaturalLanguageUnderstandingV1(options *NaturalLanguageUnderstandingV1Options) (service *NaturalLanguageUnderstandingV1, err error) {
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

	service = &NaturalLanguageUnderstandingV1{
		Service: baseService,
		Version: options.Version,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "naturalLanguageUnderstanding" suitable for processing requests.
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) Clone() *NaturalLanguageUnderstandingV1 {
	if core.IsNil(naturalLanguageUnderstanding) {
		return nil
	}
	clone := *naturalLanguageUnderstanding
	clone.Service = naturalLanguageUnderstanding.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) SetServiceURL(url string) error {
	return naturalLanguageUnderstanding.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) GetServiceURL() string {
	return naturalLanguageUnderstanding.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) SetDefaultHeaders(headers http.Header) {
	naturalLanguageUnderstanding.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) SetEnableGzipCompression(enableGzip bool) {
	naturalLanguageUnderstanding.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) GetEnableGzipCompression() bool {
	return naturalLanguageUnderstanding.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	naturalLanguageUnderstanding.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) DisableRetries() {
	naturalLanguageUnderstanding.Service.DisableRetries()
}

// DisableSSLVerification bypasses verification of the server's SSL certificate
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) DisableSSLVerification() {
	naturalLanguageUnderstanding.Service.DisableSSLVerification()
}

// Analyze : Analyze text
// Analyzes text, HTML, or a public webpage for the following features:
// - Categories
// - Concepts
// - Emotion
// - Entities
// - Keywords
// - Metadata
// - Relations
// - Semantic roles
// - Sentiment
// - Syntax
// - Summarization (Experimental)
//
// If a language for the input text is not specified with the `language` parameter, the service [automatically detects
// the
// language](https://cloud.ibm.com/docs/natural-language-understanding?topic=natural-language-understanding-detectable-languages).
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) Analyze(analyzeOptions *AnalyzeOptions) (result *AnalysisResults, response *core.DetailedResponse, err error) {
	return naturalLanguageUnderstanding.AnalyzeWithContext(context.Background(), analyzeOptions)
}

// AnalyzeWithContext is an alternate form of the Analyze method which supports a Context parameter
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) AnalyzeWithContext(ctx context.Context, analyzeOptions *AnalyzeOptions) (result *AnalysisResults, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(analyzeOptions, "analyzeOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(analyzeOptions, "analyzeOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = naturalLanguageUnderstanding.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(naturalLanguageUnderstanding.Service.Options.URL, `/v1/analyze`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range analyzeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("natural-language-understanding", "V1", "Analyze")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*naturalLanguageUnderstanding.Version))

	body := make(map[string]interface{})
	if analyzeOptions.Features != nil {
		body["features"] = analyzeOptions.Features
	}
	if analyzeOptions.Text != nil {
		body["text"] = analyzeOptions.Text
	}
	if analyzeOptions.HTML != nil {
		body["html"] = analyzeOptions.HTML
	}
	if analyzeOptions.URL != nil {
		body["url"] = analyzeOptions.URL
	}
	if analyzeOptions.Clean != nil {
		body["clean"] = analyzeOptions.Clean
	}
	if analyzeOptions.Xpath != nil {
		body["xpath"] = analyzeOptions.Xpath
	}
	if analyzeOptions.FallbackToRaw != nil {
		body["fallback_to_raw"] = analyzeOptions.FallbackToRaw
	}
	if analyzeOptions.ReturnAnalyzedText != nil {
		body["return_analyzed_text"] = analyzeOptions.ReturnAnalyzedText
	}
	if analyzeOptions.Language != nil {
		body["language"] = analyzeOptions.Language
	}
	if analyzeOptions.LimitTextCharacters != nil {
		body["limit_text_characters"] = analyzeOptions.LimitTextCharacters
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
	response, err = naturalLanguageUnderstanding.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAnalysisResults)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ListModels : List models
// Lists Watson Knowledge Studio [custom entities and relations
// models](https://cloud.ibm.com/docs/natural-language-understanding?topic=natural-language-understanding-customizing)
// that are deployed to your Natural Language Understanding service.
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) ListModels(listModelsOptions *ListModelsOptions) (result *ListModelsResults, response *core.DetailedResponse, err error) {
	return naturalLanguageUnderstanding.ListModelsWithContext(context.Background(), listModelsOptions)
}

// ListModelsWithContext is an alternate form of the ListModels method which supports a Context parameter
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) ListModelsWithContext(ctx context.Context, listModelsOptions *ListModelsOptions) (result *ListModelsResults, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listModelsOptions, "listModelsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = naturalLanguageUnderstanding.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(naturalLanguageUnderstanding.Service.Options.URL, `/v1/models`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listModelsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("natural-language-understanding", "V1", "ListModels")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*naturalLanguageUnderstanding.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = naturalLanguageUnderstanding.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListModelsResults)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteModel : Delete model
// Deletes a custom model.
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) DeleteModel(deleteModelOptions *DeleteModelOptions) (result *DeleteModelResults, response *core.DetailedResponse, err error) {
	return naturalLanguageUnderstanding.DeleteModelWithContext(context.Background(), deleteModelOptions)
}

// DeleteModelWithContext is an alternate form of the DeleteModel method which supports a Context parameter
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) DeleteModelWithContext(ctx context.Context, deleteModelOptions *DeleteModelOptions) (result *DeleteModelResults, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = naturalLanguageUnderstanding.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(naturalLanguageUnderstanding.Service.Options.URL, `/v1/models/{model_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("natural-language-understanding", "V1", "DeleteModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*naturalLanguageUnderstanding.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = naturalLanguageUnderstanding.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeleteModelResults)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// AnalysisResults : Results of the analysis, organized by feature.
type AnalysisResults struct {
	// Language used to analyze the text.
	Language *string `json:"language,omitempty"`

	// Text that was used in the analysis.
	AnalyzedText *string `json:"analyzed_text,omitempty"`

	// URL of the webpage that was analyzed.
	RetrievedURL *string `json:"retrieved_url,omitempty"`

	// API usage information for the request.
	Usage *AnalysisResultsUsage `json:"usage,omitempty"`

	// The general concepts referenced or alluded to in the analyzed text.
	Concepts []ConceptsResult `json:"concepts,omitempty"`

	// The entities detected in the analyzed text.
	Entities []EntitiesResult `json:"entities,omitempty"`

	// The keywords from the analyzed text.
	Keywords []KeywordsResult `json:"keywords,omitempty"`

	// The categories that the service assigned to the analyzed text.
	Categories []CategoriesResult `json:"categories,omitempty"`

	// The anger, disgust, fear, joy, or sadness conveyed by the content.
	Emotion *EmotionResult `json:"emotion,omitempty"`

	// Webpage metadata, such as the author and the title of the page.
	Metadata *FeaturesResultsMetadata `json:"metadata,omitempty"`

	// The relationships between entities in the content.
	Relations []RelationsResult `json:"relations,omitempty"`

	// Sentences parsed into `subject`, `action`, and `object` form.
	SemanticRoles []SemanticRolesResult `json:"semantic_roles,omitempty"`

	// The sentiment of the content.
	Sentiment *SentimentResult `json:"sentiment,omitempty"`

	// Tokens and sentences returned from syntax analysis.
	Syntax *SyntaxResult `json:"syntax,omitempty"`
}


// UnmarshalAnalysisResults unmarshals an instance of AnalysisResults from the specified map of raw messages.
func UnmarshalAnalysisResults(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AnalysisResults)
	err = core.UnmarshalPrimitive(m, "language", &obj.Language)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "analyzed_text", &obj.AnalyzedText)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "retrieved_url", &obj.RetrievedURL)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "usage", &obj.Usage, UnmarshalAnalysisResultsUsage)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "concepts", &obj.Concepts, UnmarshalConceptsResult)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "entities", &obj.Entities, UnmarshalEntitiesResult)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "keywords", &obj.Keywords, UnmarshalKeywordsResult)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "categories", &obj.Categories, UnmarshalCategoriesResult)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "emotion", &obj.Emotion, UnmarshalEmotionResult)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalFeaturesResultsMetadata)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "relations", &obj.Relations, UnmarshalRelationsResult)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "semantic_roles", &obj.SemanticRoles, UnmarshalSemanticRolesResult)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "sentiment", &obj.Sentiment, UnmarshalSentimentResult)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "syntax", &obj.Syntax, UnmarshalSyntaxResult)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AnalysisResultsUsage : API usage information for the request.
type AnalysisResultsUsage struct {
	// Number of features used in the API call.
	Features *int64 `json:"features,omitempty"`

	// Number of text characters processed.
	TextCharacters *int64 `json:"text_characters,omitempty"`

	// Number of 10,000-character units processed.
	TextUnits *int64 `json:"text_units,omitempty"`
}


// UnmarshalAnalysisResultsUsage unmarshals an instance of AnalysisResultsUsage from the specified map of raw messages.
func UnmarshalAnalysisResultsUsage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AnalysisResultsUsage)
	err = core.UnmarshalPrimitive(m, "features", &obj.Features)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text_characters", &obj.TextCharacters)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text_units", &obj.TextUnits)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AnalyzeOptions : The Analyze options.
type AnalyzeOptions struct {
	// Specific features to analyze the document for.
	Features *Features `json:"features" validate:"required"`

	// The plain text to analyze. One of the `text`, `html`, or `url` parameters is required.
	Text *string `json:"text,omitempty"`

	// The HTML file to analyze. One of the `text`, `html`, or `url` parameters is required.
	HTML *string `json:"html,omitempty"`

	// The webpage to analyze. One of the `text`, `html`, or `url` parameters is required.
	URL *string `json:"url,omitempty"`

	// Set this to `false` to disable webpage cleaning. For more information about webpage cleaning, see [Analyzing
	// webpages](https://cloud.ibm.com/docs/natural-language-understanding?topic=natural-language-understanding-analyzing-webpages).
	Clean *bool `json:"clean,omitempty"`

	// An [XPath
	// query](https://cloud.ibm.com/docs/natural-language-understanding?topic=natural-language-understanding-analyzing-webpages#xpath)
	// to perform on `html` or `url` input. Results of the query will be appended to the cleaned webpage text before it is
	// analyzed. To analyze only the results of the XPath query, set the `clean` parameter to `false`.
	Xpath *string `json:"xpath,omitempty"`

	// Whether to use raw HTML content if text cleaning fails.
	FallbackToRaw *bool `json:"fallback_to_raw,omitempty"`

	// Whether or not to return the analyzed text.
	ReturnAnalyzedText *bool `json:"return_analyzed_text,omitempty"`

	// ISO 639-1 code that specifies the language of your text. This overrides automatic language detection. Language
	// support differs depending on the features you include in your analysis. For more information, see [Language
	// support](https://cloud.ibm.com/docs/natural-language-understanding?topic=natural-language-understanding-language-support).
	Language *string `json:"language,omitempty"`

	// Sets the maximum number of characters that are processed by the service.
	LimitTextCharacters *int64 `json:"limit_text_characters,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAnalyzeOptions : Instantiate AnalyzeOptions
func (*NaturalLanguageUnderstandingV1) NewAnalyzeOptions(features *Features) *AnalyzeOptions {
	return &AnalyzeOptions{
		Features: features,
	}
}

// SetFeatures : Allow user to set Features
func (options *AnalyzeOptions) SetFeatures(features *Features) *AnalyzeOptions {
	options.Features = features
	return options
}

// SetText : Allow user to set Text
func (options *AnalyzeOptions) SetText(text string) *AnalyzeOptions {
	options.Text = core.StringPtr(text)
	return options
}

// SetHTML : Allow user to set HTML
func (options *AnalyzeOptions) SetHTML(html string) *AnalyzeOptions {
	options.HTML = core.StringPtr(html)
	return options
}

// SetURL : Allow user to set URL
func (options *AnalyzeOptions) SetURL(url string) *AnalyzeOptions {
	options.URL = core.StringPtr(url)
	return options
}

// SetClean : Allow user to set Clean
func (options *AnalyzeOptions) SetClean(clean bool) *AnalyzeOptions {
	options.Clean = core.BoolPtr(clean)
	return options
}

// SetXpath : Allow user to set Xpath
func (options *AnalyzeOptions) SetXpath(xpath string) *AnalyzeOptions {
	options.Xpath = core.StringPtr(xpath)
	return options
}

// SetFallbackToRaw : Allow user to set FallbackToRaw
func (options *AnalyzeOptions) SetFallbackToRaw(fallbackToRaw bool) *AnalyzeOptions {
	options.FallbackToRaw = core.BoolPtr(fallbackToRaw)
	return options
}

// SetReturnAnalyzedText : Allow user to set ReturnAnalyzedText
func (options *AnalyzeOptions) SetReturnAnalyzedText(returnAnalyzedText bool) *AnalyzeOptions {
	options.ReturnAnalyzedText = core.BoolPtr(returnAnalyzedText)
	return options
}

// SetLanguage : Allow user to set Language
func (options *AnalyzeOptions) SetLanguage(language string) *AnalyzeOptions {
	options.Language = core.StringPtr(language)
	return options
}

// SetLimitTextCharacters : Allow user to set LimitTextCharacters
func (options *AnalyzeOptions) SetLimitTextCharacters(limitTextCharacters int64) *AnalyzeOptions {
	options.LimitTextCharacters = core.Int64Ptr(limitTextCharacters)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *AnalyzeOptions) SetHeaders(param map[string]string) *AnalyzeOptions {
	options.Headers = param
	return options
}

// Author : The author of the analyzed content.
type Author struct {
	// Name of the author.
	Name *string `json:"name,omitempty"`
}


// UnmarshalAuthor unmarshals an instance of Author from the specified map of raw messages.
func UnmarshalAuthor(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Author)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CategoriesOptions : Returns a five-level taxonomy of the content. The top three categories are returned.
//
// Supported languages: Arabic, English, French, German, Italian, Japanese, Korean, Portuguese, Spanish.
type CategoriesOptions struct {
	// Set this to `true` to return explanations for each categorization. **This is available only for English
	// categories.**.
	Explanation *bool `json:"explanation,omitempty"`

	// Maximum number of categories to return.
	Limit *int64 `json:"limit,omitempty"`

	// Enter a [custom
	// model](https://cloud.ibm.com/docs/natural-language-understanding?topic=natural-language-understanding-customizing)
	// ID to override the standard categories model.
	//
	// The custom categories experimental feature will be retired on 19 December 2019. On that date, deployed custom
	// categories models will no longer be accessible in Natural Language Understanding. The feature will be removed from
	// Knowledge Studio on an earlier date. Custom categories models will no longer be accessible in Knowledge Studio on 17
	// December 2019.
	Model *string `json:"model,omitempty"`
}


// UnmarshalCategoriesOptions unmarshals an instance of CategoriesOptions from the specified map of raw messages.
func UnmarshalCategoriesOptions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CategoriesOptions)
	err = core.UnmarshalPrimitive(m, "explanation", &obj.Explanation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "model", &obj.Model)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CategoriesRelevantText : Relevant text that contributed to the categorization.
type CategoriesRelevantText struct {
	// Text from the analyzed source that supports the categorization.
	Text *string `json:"text,omitempty"`
}


// UnmarshalCategoriesRelevantText unmarshals an instance of CategoriesRelevantText from the specified map of raw messages.
func UnmarshalCategoriesRelevantText(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CategoriesRelevantText)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CategoriesResult : A categorization of the analyzed text.
type CategoriesResult struct {
	// The path to the category through the 5-level taxonomy hierarchy. For more information about the categories, see
	// [Categories
	// hierarchy](https://cloud.ibm.com/docs/natural-language-understanding?topic=natural-language-understanding-categories#categories-hierarchy).
	Label *string `json:"label,omitempty"`

	// Confidence score for the category classification. Higher values indicate greater confidence.
	Score *float64 `json:"score,omitempty"`

	// Information that helps to explain what contributed to the categories result.
	Explanation *CategoriesResultExplanation `json:"explanation,omitempty"`
}


// UnmarshalCategoriesResult unmarshals an instance of CategoriesResult from the specified map of raw messages.
func UnmarshalCategoriesResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CategoriesResult)
	err = core.UnmarshalPrimitive(m, "label", &obj.Label)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "score", &obj.Score)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "explanation", &obj.Explanation, UnmarshalCategoriesResultExplanation)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CategoriesResultExplanation : Information that helps to explain what contributed to the categories result.
type CategoriesResultExplanation struct {
	// An array of relevant text from the source that contributed to the categorization. The sorted array begins with the
	// phrase that contributed most significantly to the result, followed by phrases that were less and less impactful.
	RelevantText []CategoriesRelevantText `json:"relevant_text,omitempty"`
}


// UnmarshalCategoriesResultExplanation unmarshals an instance of CategoriesResultExplanation from the specified map of raw messages.
func UnmarshalCategoriesResultExplanation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CategoriesResultExplanation)
	err = core.UnmarshalModel(m, "relevant_text", &obj.RelevantText, UnmarshalCategoriesRelevantText)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConceptsOptions : Returns high-level concepts in the content. For example, a research paper about deep learning might return the
// concept, "Artificial Intelligence" although the term is not mentioned.
//
// Supported languages: English, French, German, Italian, Japanese, Korean, Portuguese, Spanish.
type ConceptsOptions struct {
	// Maximum number of concepts to return.
	Limit *int64 `json:"limit,omitempty"`
}


// UnmarshalConceptsOptions unmarshals an instance of ConceptsOptions from the specified map of raw messages.
func UnmarshalConceptsOptions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConceptsOptions)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConceptsResult : The general concepts referenced or alluded to in the analyzed text.
type ConceptsResult struct {
	// Name of the concept.
	Text *string `json:"text,omitempty"`

	// Relevance score between 0 and 1. Higher scores indicate greater relevance.
	Relevance *float64 `json:"relevance,omitempty"`

	// Link to the corresponding DBpedia resource.
	DbpediaResource *string `json:"dbpedia_resource,omitempty"`
}


// UnmarshalConceptsResult unmarshals an instance of ConceptsResult from the specified map of raw messages.
func UnmarshalConceptsResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConceptsResult)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "relevance", &obj.Relevance)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "dbpedia_resource", &obj.DbpediaResource)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteModelOptions : The DeleteModel options.
type DeleteModelOptions struct {
	// Model ID of the model to delete.
	ModelID *string `json:"model_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteModelOptions : Instantiate DeleteModelOptions
func (*NaturalLanguageUnderstandingV1) NewDeleteModelOptions(modelID string) *DeleteModelOptions {
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

// DeleteModelResults : Delete model results.
type DeleteModelResults struct {
	// model_id of the deleted model.
	Deleted *string `json:"deleted,omitempty"`
}


// UnmarshalDeleteModelResults unmarshals an instance of DeleteModelResults from the specified map of raw messages.
func UnmarshalDeleteModelResults(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteModelResults)
	err = core.UnmarshalPrimitive(m, "deleted", &obj.Deleted)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DisambiguationResult : Disambiguation information for the entity.
type DisambiguationResult struct {
	// Common entity name.
	Name *string `json:"name,omitempty"`

	// Link to the corresponding DBpedia resource.
	DbpediaResource *string `json:"dbpedia_resource,omitempty"`

	// Entity subtype information.
	Subtype []string `json:"subtype,omitempty"`
}


// UnmarshalDisambiguationResult unmarshals an instance of DisambiguationResult from the specified map of raw messages.
func UnmarshalDisambiguationResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DisambiguationResult)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "dbpedia_resource", &obj.DbpediaResource)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "subtype", &obj.Subtype)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DocumentEmotionResults : Emotion results for the document as a whole.
type DocumentEmotionResults struct {
	// Emotion results for the document as a whole.
	Emotion *EmotionScores `json:"emotion,omitempty"`
}


// UnmarshalDocumentEmotionResults unmarshals an instance of DocumentEmotionResults from the specified map of raw messages.
func UnmarshalDocumentEmotionResults(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DocumentEmotionResults)
	err = core.UnmarshalModel(m, "emotion", &obj.Emotion, UnmarshalEmotionScores)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DocumentSentimentResults : DocumentSentimentResults struct
type DocumentSentimentResults struct {
	// Indicates whether the sentiment is positive, neutral, or negative.
	Label *string `json:"label,omitempty"`

	// Sentiment score from -1 (negative) to 1 (positive).
	Score *float64 `json:"score,omitempty"`
}


// UnmarshalDocumentSentimentResults unmarshals an instance of DocumentSentimentResults from the specified map of raw messages.
func UnmarshalDocumentSentimentResults(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DocumentSentimentResults)
	err = core.UnmarshalPrimitive(m, "label", &obj.Label)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "score", &obj.Score)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EmotionOptions : Detects anger, disgust, fear, joy, or sadness that is conveyed in the content or by the context around target phrases
// specified in the targets parameter. You can analyze emotion for detected entities with `entities.emotion` and for
// keywords with `keywords.emotion`.
//
// Supported languages: English.
type EmotionOptions struct {
	// Set this to `false` to hide document-level emotion results.
	Document *bool `json:"document,omitempty"`

	// Emotion results will be returned for each target string that is found in the document.
	Targets []string `json:"targets,omitempty"`
}


// UnmarshalEmotionOptions unmarshals an instance of EmotionOptions from the specified map of raw messages.
func UnmarshalEmotionOptions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EmotionOptions)
	err = core.UnmarshalPrimitive(m, "document", &obj.Document)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "targets", &obj.Targets)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EmotionResult : The detected anger, disgust, fear, joy, or sadness that is conveyed by the content. Emotion information can be
// returned for detected entities, keywords, or user-specified target phrases found in the text.
type EmotionResult struct {
	// Emotion results for the document as a whole.
	Document *DocumentEmotionResults `json:"document,omitempty"`

	// Emotion results for specified targets.
	Targets []TargetedEmotionResults `json:"targets,omitempty"`
}


// UnmarshalEmotionResult unmarshals an instance of EmotionResult from the specified map of raw messages.
func UnmarshalEmotionResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EmotionResult)
	err = core.UnmarshalModel(m, "document", &obj.Document, UnmarshalDocumentEmotionResults)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "targets", &obj.Targets, UnmarshalTargetedEmotionResults)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EmotionScores : EmotionScores struct
type EmotionScores struct {
	// Anger score from 0 to 1. A higher score means that the text is more likely to convey anger.
	Anger *float64 `json:"anger,omitempty"`

	// Disgust score from 0 to 1. A higher score means that the text is more likely to convey disgust.
	Disgust *float64 `json:"disgust,omitempty"`

	// Fear score from 0 to 1. A higher score means that the text is more likely to convey fear.
	Fear *float64 `json:"fear,omitempty"`

	// Joy score from 0 to 1. A higher score means that the text is more likely to convey joy.
	Joy *float64 `json:"joy,omitempty"`

	// Sadness score from 0 to 1. A higher score means that the text is more likely to convey sadness.
	Sadness *float64 `json:"sadness,omitempty"`
}


// UnmarshalEmotionScores unmarshals an instance of EmotionScores from the specified map of raw messages.
func UnmarshalEmotionScores(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EmotionScores)
	err = core.UnmarshalPrimitive(m, "anger", &obj.Anger)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "disgust", &obj.Disgust)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "fear", &obj.Fear)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "joy", &obj.Joy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sadness", &obj.Sadness)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EntitiesOptions : Identifies people, cities, organizations, and other entities in the content. For more information, see [Entity types
// and
// subtypes](https://cloud.ibm.com/docs/natural-language-understanding?topic=natural-language-understanding-entity-types).
//
// Supported languages: English, French, German, Italian, Japanese, Korean, Portuguese, Russian, Spanish, Swedish.
// Arabic, Chinese, and Dutch are supported only through custom models.
type EntitiesOptions struct {
	// Maximum number of entities to return.
	Limit *int64 `json:"limit,omitempty"`

	// Set this to `true` to return locations of entity mentions.
	Mentions *bool `json:"mentions,omitempty"`

	// Enter a [custom
	// model](https://cloud.ibm.com/docs/natural-language-understanding?topic=natural-language-understanding-customizing)
	// ID to override the standard entity detection model.
	Model *string `json:"model,omitempty"`

	// Set this to `true` to return sentiment information for detected entities.
	Sentiment *bool `json:"sentiment,omitempty"`

	// Set this to `true` to analyze emotion for detected keywords.
	Emotion *bool `json:"emotion,omitempty"`
}


// UnmarshalEntitiesOptions unmarshals an instance of EntitiesOptions from the specified map of raw messages.
func UnmarshalEntitiesOptions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EntitiesOptions)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "mentions", &obj.Mentions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "model", &obj.Model)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sentiment", &obj.Sentiment)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "emotion", &obj.Emotion)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EntitiesResult : The important people, places, geopolitical entities and other types of entities in your content.
type EntitiesResult struct {
	// Entity type.
	Type *string `json:"type,omitempty"`

	// The name of the entity.
	Text *string `json:"text,omitempty"`

	// Relevance score from 0 to 1. Higher values indicate greater relevance.
	Relevance *float64 `json:"relevance,omitempty"`

	// Confidence in the entity identification from 0 to 1. Higher values indicate higher confidence. In standard entities
	// requests, confidence is returned only for English text. All entities requests that use custom models return the
	// confidence score.
	Confidence *float64 `json:"confidence,omitempty"`

	// Entity mentions and locations.
	Mentions []EntityMention `json:"mentions,omitempty"`

	// How many times the entity was mentioned in the text.
	Count *int64 `json:"count,omitempty"`

	// Emotion analysis results for the entity, enabled with the `emotion` option.
	Emotion *EmotionScores `json:"emotion,omitempty"`

	// Sentiment analysis results for the entity, enabled with the `sentiment` option.
	Sentiment *FeatureSentimentResults `json:"sentiment,omitempty"`

	// Disambiguation information for the entity.
	Disambiguation *DisambiguationResult `json:"disambiguation,omitempty"`
}


// UnmarshalEntitiesResult unmarshals an instance of EntitiesResult from the specified map of raw messages.
func UnmarshalEntitiesResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EntitiesResult)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "relevance", &obj.Relevance)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "confidence", &obj.Confidence)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "mentions", &obj.Mentions, UnmarshalEntityMention)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "count", &obj.Count)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "emotion", &obj.Emotion, UnmarshalEmotionScores)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "sentiment", &obj.Sentiment, UnmarshalFeatureSentimentResults)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "disambiguation", &obj.Disambiguation, UnmarshalDisambiguationResult)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EntityMention : EntityMention struct
type EntityMention struct {
	// Entity mention text.
	Text *string `json:"text,omitempty"`

	// Character offsets indicating the beginning and end of the mention in the analyzed text.
	Location []int64 `json:"location,omitempty"`

	// Confidence in the entity identification from 0 to 1. Higher values indicate higher confidence. In standard entities
	// requests, confidence is returned only for English text. All entities requests that use custom models return the
	// confidence score.
	Confidence *float64 `json:"confidence,omitempty"`
}


// UnmarshalEntityMention unmarshals an instance of EntityMention from the specified map of raw messages.
func UnmarshalEntityMention(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EntityMention)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "location", &obj.Location)
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

// FeatureSentimentResults : FeatureSentimentResults struct
type FeatureSentimentResults struct {
	// Sentiment score from -1 (negative) to 1 (positive).
	Score *float64 `json:"score,omitempty"`
}


// UnmarshalFeatureSentimentResults unmarshals an instance of FeatureSentimentResults from the specified map of raw messages.
func UnmarshalFeatureSentimentResults(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FeatureSentimentResults)
	err = core.UnmarshalPrimitive(m, "score", &obj.Score)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Features : Analysis features and options.
type Features struct {
	// Returns high-level concepts in the content. For example, a research paper about deep learning might return the
	// concept, "Artificial Intelligence" although the term is not mentioned.
	//
	// Supported languages: English, French, German, Italian, Japanese, Korean, Portuguese, Spanish.
	Concepts *ConceptsOptions `json:"concepts,omitempty"`

	// Detects anger, disgust, fear, joy, or sadness that is conveyed in the content or by the context around target
	// phrases specified in the targets parameter. You can analyze emotion for detected entities with `entities.emotion`
	// and for keywords with `keywords.emotion`.
	//
	// Supported languages: English.
	Emotion *EmotionOptions `json:"emotion,omitempty"`

	// Identifies people, cities, organizations, and other entities in the content. For more information, see [Entity types
	// and
	// subtypes](https://cloud.ibm.com/docs/natural-language-understanding?topic=natural-language-understanding-entity-types).
	//
	// Supported languages: English, French, German, Italian, Japanese, Korean, Portuguese, Russian, Spanish, Swedish.
	// Arabic, Chinese, and Dutch are supported only through custom models.
	Entities *EntitiesOptions `json:"entities,omitempty"`

	// Returns important keywords in the content.
	//
	// Supported languages: English, French, German, Italian, Japanese, Korean, Portuguese, Russian, Spanish, Swedish.
	Keywords *KeywordsOptions `json:"keywords,omitempty"`

	// Returns information from the document, including author name, title, RSS/ATOM feeds, prominent page image, and
	// publication date. Supports URL and HTML input types only.
	Metadata interface{} `json:"metadata,omitempty"`

	// Recognizes when two entities are related and identifies the type of relation. For example, an `awardedTo` relation
	// might connect the entities "Nobel Prize" and "Albert Einstein". For more information, see [Relation
	// types](https://cloud.ibm.com/docs/natural-language-understanding?topic=natural-language-understanding-relations).
	//
	// Supported languages: Arabic, English, German, Japanese, Korean, Spanish. Chinese, Dutch, French, Italian, and
	// Portuguese custom models are also supported.
	Relations *RelationsOptions `json:"relations,omitempty"`

	// Parses sentences into subject, action, and object form.
	//
	// Supported languages: English, German, Japanese, Korean, Spanish.
	SemanticRoles *SemanticRolesOptions `json:"semantic_roles,omitempty"`

	// Analyzes the general sentiment of your content or the sentiment toward specific target phrases. You can analyze
	// sentiment for detected entities with `entities.sentiment` and for keywords with `keywords.sentiment`.
	//
	//  Supported languages: Arabic, English, French, German, Italian, Japanese, Korean, Portuguese, Russian, Spanish.
	Sentiment *SentimentOptions `json:"sentiment,omitempty"`

	// Returns a five-level taxonomy of the content. The top three categories are returned.
	//
	// Supported languages: Arabic, English, French, German, Italian, Japanese, Korean, Portuguese, Spanish.
	Categories *CategoriesOptions `json:"categories,omitempty"`

	// Returns tokens and sentences from the input text.
	Syntax *SyntaxOptions `json:"syntax,omitempty"`
}


// UnmarshalFeatures unmarshals an instance of Features from the specified map of raw messages.
func UnmarshalFeatures(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Features)
	err = core.UnmarshalModel(m, "concepts", &obj.Concepts, UnmarshalConceptsOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "emotion", &obj.Emotion, UnmarshalEmotionOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "entities", &obj.Entities, UnmarshalEntitiesOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "keywords", &obj.Keywords, UnmarshalKeywordsOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "metadata", &obj.Metadata)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "relations", &obj.Relations, UnmarshalRelationsOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "semantic_roles", &obj.SemanticRoles, UnmarshalSemanticRolesOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "sentiment", &obj.Sentiment, UnmarshalSentimentOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "categories", &obj.Categories, UnmarshalCategoriesOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "syntax", &obj.Syntax, UnmarshalSyntaxOptions)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FeaturesResultsMetadata : Webpage metadata, such as the author and the title of the page.
type FeaturesResultsMetadata struct {
	// The authors of the document.
	Authors []Author `json:"authors,omitempty"`

	// The publication date in the format ISO 8601.
	PublicationDate *string `json:"publication_date,omitempty"`

	// The title of the document.
	Title *string `json:"title,omitempty"`

	// URL of a prominent image on the webpage.
	Image *string `json:"image,omitempty"`

	// RSS/ATOM feeds found on the webpage.
	Feeds []Feed `json:"feeds,omitempty"`
}


// UnmarshalFeaturesResultsMetadata unmarshals an instance of FeaturesResultsMetadata from the specified map of raw messages.
func UnmarshalFeaturesResultsMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FeaturesResultsMetadata)
	err = core.UnmarshalModel(m, "authors", &obj.Authors, UnmarshalAuthor)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "publication_date", &obj.PublicationDate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "title", &obj.Title)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "image", &obj.Image)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "feeds", &obj.Feeds, UnmarshalFeed)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Feed : RSS or ATOM feed found on the webpage.
type Feed struct {
	// URL of the RSS or ATOM feed.
	Link *string `json:"link,omitempty"`
}


// UnmarshalFeed unmarshals an instance of Feed from the specified map of raw messages.
func UnmarshalFeed(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Feed)
	err = core.UnmarshalPrimitive(m, "link", &obj.Link)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeywordsOptions : Returns important keywords in the content.
//
// Supported languages: English, French, German, Italian, Japanese, Korean, Portuguese, Russian, Spanish, Swedish.
type KeywordsOptions struct {
	// Maximum number of keywords to return.
	Limit *int64 `json:"limit,omitempty"`

	// Set this to `true` to return sentiment information for detected keywords.
	Sentiment *bool `json:"sentiment,omitempty"`

	// Set this to `true` to analyze emotion for detected keywords.
	Emotion *bool `json:"emotion,omitempty"`
}


// UnmarshalKeywordsOptions unmarshals an instance of KeywordsOptions from the specified map of raw messages.
func UnmarshalKeywordsOptions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeywordsOptions)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sentiment", &obj.Sentiment)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "emotion", &obj.Emotion)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeywordsResult : The important keywords in the content, organized by relevance.
type KeywordsResult struct {
	// Number of times the keyword appears in the analyzed text.
	Count *int64 `json:"count,omitempty"`

	// Relevance score from 0 to 1. Higher values indicate greater relevance.
	Relevance *float64 `json:"relevance,omitempty"`

	// The keyword text.
	Text *string `json:"text,omitempty"`

	// Emotion analysis results for the keyword, enabled with the `emotion` option.
	Emotion *EmotionScores `json:"emotion,omitempty"`

	// Sentiment analysis results for the keyword, enabled with the `sentiment` option.
	Sentiment *FeatureSentimentResults `json:"sentiment,omitempty"`
}


// UnmarshalKeywordsResult unmarshals an instance of KeywordsResult from the specified map of raw messages.
func UnmarshalKeywordsResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeywordsResult)
	err = core.UnmarshalPrimitive(m, "count", &obj.Count)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "relevance", &obj.Relevance)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "emotion", &obj.Emotion, UnmarshalEmotionScores)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "sentiment", &obj.Sentiment, UnmarshalFeatureSentimentResults)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListModelsOptions : The ListModels options.
type ListModelsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListModelsOptions : Instantiate ListModelsOptions
func (*NaturalLanguageUnderstandingV1) NewListModelsOptions() *ListModelsOptions {
	return &ListModelsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListModelsOptions) SetHeaders(param map[string]string) *ListModelsOptions {
	options.Headers = param
	return options
}

// ListModelsResults : Custom models that are available for entities and relations.
type ListModelsResults struct {
	// An array of available models.
	Models []Model `json:"models,omitempty"`
}


// UnmarshalListModelsResults unmarshals an instance of ListModelsResults from the specified map of raw messages.
func UnmarshalListModelsResults(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListModelsResults)
	err = core.UnmarshalModel(m, "models", &obj.Models, UnmarshalModel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Model : Model struct
type Model struct {
	// When the status is `available`, the model is ready to use.
	Status *string `json:"status,omitempty"`

	// Unique model ID.
	ModelID *string `json:"model_id,omitempty"`

	// ISO 639-1 code that indicates the language of the model.
	Language *string `json:"language,omitempty"`

	// Model description.
	Description *string `json:"description,omitempty"`

	// ID of the Watson Knowledge Studio workspace that deployed this model to Natural Language Understanding.
	WorkspaceID *string `json:"workspace_id,omitempty"`

	// The model version, if it was manually provided in Watson Knowledge Studio.
	ModelVersion *string `json:"model_version,omitempty"`

	// Deprecated — use `model_version`.
	Version *string `json:"version,omitempty"`

	// The description of the version, if it was manually provided in Watson Knowledge Studio.
	VersionDescription *string `json:"version_description,omitempty"`

	// A dateTime indicating when the model was created.
	Created *strfmt.DateTime `json:"created,omitempty"`
}

// Constants associated with the Model.Status property.
// When the status is `available`, the model is ready to use.
const (
	ModelStatusAvailableConst = "available"
	ModelStatusDeletedConst = "deleted"
	ModelStatusDeployingConst = "deploying"
	ModelStatusErrorConst = "error"
	ModelStatusStartingConst = "starting"
	ModelStatusTrainingConst = "training"
)


// UnmarshalModel unmarshals an instance of Model from the specified map of raw messages.
func UnmarshalModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Model)
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "model_id", &obj.ModelID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "language", &obj.Language)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "workspace_id", &obj.WorkspaceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "model_version", &obj.ModelVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version_description", &obj.VersionDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created", &obj.Created)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RelationArgument : RelationArgument struct
type RelationArgument struct {
	// An array of extracted entities.
	Entities []RelationEntity `json:"entities,omitempty"`

	// Character offsets indicating the beginning and end of the mention in the analyzed text.
	Location []int64 `json:"location,omitempty"`

	// Text that corresponds to the argument.
	Text *string `json:"text,omitempty"`
}


// UnmarshalRelationArgument unmarshals an instance of RelationArgument from the specified map of raw messages.
func UnmarshalRelationArgument(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RelationArgument)
	err = core.UnmarshalModel(m, "entities", &obj.Entities, UnmarshalRelationEntity)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "location", &obj.Location)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RelationEntity : An entity that corresponds with an argument in a relation.
type RelationEntity struct {
	// Text that corresponds to the entity.
	Text *string `json:"text,omitempty"`

	// Entity type.
	Type *string `json:"type,omitempty"`
}


// UnmarshalRelationEntity unmarshals an instance of RelationEntity from the specified map of raw messages.
func UnmarshalRelationEntity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RelationEntity)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RelationsOptions : Recognizes when two entities are related and identifies the type of relation. For example, an `awardedTo` relation
// might connect the entities "Nobel Prize" and "Albert Einstein". For more information, see [Relation
// types](https://cloud.ibm.com/docs/natural-language-understanding?topic=natural-language-understanding-relations).
//
// Supported languages: Arabic, English, German, Japanese, Korean, Spanish. Chinese, Dutch, French, Italian, and
// Portuguese custom models are also supported.
type RelationsOptions struct {
	// Enter a [custom
	// model](https://cloud.ibm.com/docs/natural-language-understanding?topic=natural-language-understanding-customizing)
	// ID to override the default model.
	Model *string `json:"model,omitempty"`
}


// UnmarshalRelationsOptions unmarshals an instance of RelationsOptions from the specified map of raw messages.
func UnmarshalRelationsOptions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RelationsOptions)
	err = core.UnmarshalPrimitive(m, "model", &obj.Model)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RelationsResult : The relations between entities found in the content.
type RelationsResult struct {
	// Confidence score for the relation. Higher values indicate greater confidence.
	Score *float64 `json:"score,omitempty"`

	// The sentence that contains the relation.
	Sentence *string `json:"sentence,omitempty"`

	// The type of the relation.
	Type *string `json:"type,omitempty"`

	// Entity mentions that are involved in the relation.
	Arguments []RelationArgument `json:"arguments,omitempty"`
}


// UnmarshalRelationsResult unmarshals an instance of RelationsResult from the specified map of raw messages.
func UnmarshalRelationsResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RelationsResult)
	err = core.UnmarshalPrimitive(m, "score", &obj.Score)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sentence", &obj.Sentence)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "arguments", &obj.Arguments, UnmarshalRelationArgument)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SemanticRolesEntity : SemanticRolesEntity struct
type SemanticRolesEntity struct {
	// Entity type.
	Type *string `json:"type,omitempty"`

	// The entity text.
	Text *string `json:"text,omitempty"`
}


// UnmarshalSemanticRolesEntity unmarshals an instance of SemanticRolesEntity from the specified map of raw messages.
func UnmarshalSemanticRolesEntity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SemanticRolesEntity)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SemanticRolesKeyword : SemanticRolesKeyword struct
type SemanticRolesKeyword struct {
	// The keyword text.
	Text *string `json:"text,omitempty"`
}


// UnmarshalSemanticRolesKeyword unmarshals an instance of SemanticRolesKeyword from the specified map of raw messages.
func UnmarshalSemanticRolesKeyword(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SemanticRolesKeyword)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SemanticRolesOptions : Parses sentences into subject, action, and object form.
//
// Supported languages: English, German, Japanese, Korean, Spanish.
type SemanticRolesOptions struct {
	// Maximum number of semantic_roles results to return.
	Limit *int64 `json:"limit,omitempty"`

	// Set this to `true` to return keyword information for subjects and objects.
	Keywords *bool `json:"keywords,omitempty"`

	// Set this to `true` to return entity information for subjects and objects.
	Entities *bool `json:"entities,omitempty"`
}


// UnmarshalSemanticRolesOptions unmarshals an instance of SemanticRolesOptions from the specified map of raw messages.
func UnmarshalSemanticRolesOptions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SemanticRolesOptions)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "keywords", &obj.Keywords)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "entities", &obj.Entities)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SemanticRolesResult : The object containing the actions and the objects the actions act upon.
type SemanticRolesResult struct {
	// Sentence from the source that contains the subject, action, and object.
	Sentence *string `json:"sentence,omitempty"`

	// The extracted subject from the sentence.
	Subject *SemanticRolesResultSubject `json:"subject,omitempty"`

	// The extracted action from the sentence.
	Action *SemanticRolesResultAction `json:"action,omitempty"`

	// The extracted object from the sentence.
	Object *SemanticRolesResultObject `json:"object,omitempty"`
}


// UnmarshalSemanticRolesResult unmarshals an instance of SemanticRolesResult from the specified map of raw messages.
func UnmarshalSemanticRolesResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SemanticRolesResult)
	err = core.UnmarshalPrimitive(m, "sentence", &obj.Sentence)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "subject", &obj.Subject, UnmarshalSemanticRolesResultSubject)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "action", &obj.Action, UnmarshalSemanticRolesResultAction)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "object", &obj.Object, UnmarshalSemanticRolesResultObject)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SemanticRolesResultAction : The extracted action from the sentence.
type SemanticRolesResultAction struct {
	// Analyzed text that corresponds to the action.
	Text *string `json:"text,omitempty"`

	// normalized version of the action.
	Normalized *string `json:"normalized,omitempty"`

	Verb *SemanticRolesVerb `json:"verb,omitempty"`
}


// UnmarshalSemanticRolesResultAction unmarshals an instance of SemanticRolesResultAction from the specified map of raw messages.
func UnmarshalSemanticRolesResultAction(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SemanticRolesResultAction)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "normalized", &obj.Normalized)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "verb", &obj.Verb, UnmarshalSemanticRolesVerb)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SemanticRolesResultObject : The extracted object from the sentence.
type SemanticRolesResultObject struct {
	// Object text.
	Text *string `json:"text,omitempty"`

	// An array of extracted keywords.
	Keywords []SemanticRolesKeyword `json:"keywords,omitempty"`
}


// UnmarshalSemanticRolesResultObject unmarshals an instance of SemanticRolesResultObject from the specified map of raw messages.
func UnmarshalSemanticRolesResultObject(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SemanticRolesResultObject)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "keywords", &obj.Keywords, UnmarshalSemanticRolesKeyword)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SemanticRolesResultSubject : The extracted subject from the sentence.
type SemanticRolesResultSubject struct {
	// Text that corresponds to the subject role.
	Text *string `json:"text,omitempty"`

	// An array of extracted entities.
	Entities []SemanticRolesEntity `json:"entities,omitempty"`

	// An array of extracted keywords.
	Keywords []SemanticRolesKeyword `json:"keywords,omitempty"`
}


// UnmarshalSemanticRolesResultSubject unmarshals an instance of SemanticRolesResultSubject from the specified map of raw messages.
func UnmarshalSemanticRolesResultSubject(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SemanticRolesResultSubject)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "entities", &obj.Entities, UnmarshalSemanticRolesEntity)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "keywords", &obj.Keywords, UnmarshalSemanticRolesKeyword)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SemanticRolesVerb : SemanticRolesVerb struct
type SemanticRolesVerb struct {
	// The keyword text.
	Text *string `json:"text,omitempty"`

	// Verb tense.
	Tense *string `json:"tense,omitempty"`
}


// UnmarshalSemanticRolesVerb unmarshals an instance of SemanticRolesVerb from the specified map of raw messages.
func UnmarshalSemanticRolesVerb(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SemanticRolesVerb)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tense", &obj.Tense)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SentenceResult : SentenceResult struct
type SentenceResult struct {
	// The sentence.
	Text *string `json:"text,omitempty"`

	// Character offsets indicating the beginning and end of the sentence in the analyzed text.
	Location []int64 `json:"location,omitempty"`
}


// UnmarshalSentenceResult unmarshals an instance of SentenceResult from the specified map of raw messages.
func UnmarshalSentenceResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SentenceResult)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "location", &obj.Location)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SentimentOptions : Analyzes the general sentiment of your content or the sentiment toward specific target phrases. You can analyze
// sentiment for detected entities with `entities.sentiment` and for keywords with `keywords.sentiment`.
//
//  Supported languages: Arabic, English, French, German, Italian, Japanese, Korean, Portuguese, Russian, Spanish.
type SentimentOptions struct {
	// Set this to `false` to hide document-level sentiment results.
	Document *bool `json:"document,omitempty"`

	// Sentiment results will be returned for each target string that is found in the document.
	Targets []string `json:"targets,omitempty"`
}


// UnmarshalSentimentOptions unmarshals an instance of SentimentOptions from the specified map of raw messages.
func UnmarshalSentimentOptions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SentimentOptions)
	err = core.UnmarshalPrimitive(m, "document", &obj.Document)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "targets", &obj.Targets)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SentimentResult : The sentiment of the content.
type SentimentResult struct {
	// The document level sentiment.
	Document *DocumentSentimentResults `json:"document,omitempty"`

	// The targeted sentiment to analyze.
	Targets []TargetedSentimentResults `json:"targets,omitempty"`
}


// UnmarshalSentimentResult unmarshals an instance of SentimentResult from the specified map of raw messages.
func UnmarshalSentimentResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SentimentResult)
	err = core.UnmarshalModel(m, "document", &obj.Document, UnmarshalDocumentSentimentResults)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "targets", &obj.Targets, UnmarshalTargetedSentimentResults)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SyntaxOptions : Returns tokens and sentences from the input text.
type SyntaxOptions struct {
	// Tokenization options.
	Tokens *SyntaxOptionsTokens `json:"tokens,omitempty"`

	// Set this to `true` to return sentence information.
	Sentences *bool `json:"sentences,omitempty"`
}


// UnmarshalSyntaxOptions unmarshals an instance of SyntaxOptions from the specified map of raw messages.
func UnmarshalSyntaxOptions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SyntaxOptions)
	err = core.UnmarshalModel(m, "tokens", &obj.Tokens, UnmarshalSyntaxOptionsTokens)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sentences", &obj.Sentences)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SyntaxOptionsTokens : Tokenization options.
type SyntaxOptionsTokens struct {
	// Set this to `true` to return the lemma for each token.
	Lemma *bool `json:"lemma,omitempty"`

	// Set this to `true` to return the part of speech for each token.
	PartOfSpeech *bool `json:"part_of_speech,omitempty"`
}


// UnmarshalSyntaxOptionsTokens unmarshals an instance of SyntaxOptionsTokens from the specified map of raw messages.
func UnmarshalSyntaxOptionsTokens(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SyntaxOptionsTokens)
	err = core.UnmarshalPrimitive(m, "lemma", &obj.Lemma)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "part_of_speech", &obj.PartOfSpeech)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SyntaxResult : Tokens and sentences returned from syntax analysis.
type SyntaxResult struct {
	Tokens []TokenResult `json:"tokens,omitempty"`

	Sentences []SentenceResult `json:"sentences,omitempty"`
}


// UnmarshalSyntaxResult unmarshals an instance of SyntaxResult from the specified map of raw messages.
func UnmarshalSyntaxResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SyntaxResult)
	err = core.UnmarshalModel(m, "tokens", &obj.Tokens, UnmarshalTokenResult)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "sentences", &obj.Sentences, UnmarshalSentenceResult)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TargetedEmotionResults : Emotion results for a specified target.
type TargetedEmotionResults struct {
	// Targeted text.
	Text *string `json:"text,omitempty"`

	// The emotion results for the target.
	Emotion *EmotionScores `json:"emotion,omitempty"`
}


// UnmarshalTargetedEmotionResults unmarshals an instance of TargetedEmotionResults from the specified map of raw messages.
func UnmarshalTargetedEmotionResults(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TargetedEmotionResults)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "emotion", &obj.Emotion, UnmarshalEmotionScores)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TargetedSentimentResults : TargetedSentimentResults struct
type TargetedSentimentResults struct {
	// Targeted text.
	Text *string `json:"text,omitempty"`

	// Sentiment score from -1 (negative) to 1 (positive).
	Score *float64 `json:"score,omitempty"`
}


// UnmarshalTargetedSentimentResults unmarshals an instance of TargetedSentimentResults from the specified map of raw messages.
func UnmarshalTargetedSentimentResults(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TargetedSentimentResults)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "score", &obj.Score)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TokenResult : TokenResult struct
type TokenResult struct {
	// The token as it appears in the analyzed text.
	Text *string `json:"text,omitempty"`

	// The part of speech of the token. For more information about the values, see [Universal Dependencies POS
	// tags](https://universaldependencies.org/u/pos/).
	PartOfSpeech *string `json:"part_of_speech,omitempty"`

	// Character offsets indicating the beginning and end of the token in the analyzed text.
	Location []int64 `json:"location,omitempty"`

	// The [lemma](https://wikipedia.org/wiki/Lemma_%28morphology%29) of the token.
	Lemma *string `json:"lemma,omitempty"`
}

// Constants associated with the TokenResult.PartOfSpeech property.
// The part of speech of the token. For more information about the values, see [Universal Dependencies POS
// tags](https://universaldependencies.org/u/pos/).
const (
	TokenResultPartOfSpeechAdjConst = "ADJ"
	TokenResultPartOfSpeechAdpConst = "ADP"
	TokenResultPartOfSpeechAdvConst = "ADV"
	TokenResultPartOfSpeechAuxConst = "AUX"
	TokenResultPartOfSpeechCconjConst = "CCONJ"
	TokenResultPartOfSpeechDetConst = "DET"
	TokenResultPartOfSpeechIntjConst = "INTJ"
	TokenResultPartOfSpeechNounConst = "NOUN"
	TokenResultPartOfSpeechNumConst = "NUM"
	TokenResultPartOfSpeechPartConst = "PART"
	TokenResultPartOfSpeechPronConst = "PRON"
	TokenResultPartOfSpeechPropnConst = "PROPN"
	TokenResultPartOfSpeechPunctConst = "PUNCT"
	TokenResultPartOfSpeechSconjConst = "SCONJ"
	TokenResultPartOfSpeechSymConst = "SYM"
	TokenResultPartOfSpeechVerbConst = "VERB"
	TokenResultPartOfSpeechXConst = "X"
)


// UnmarshalTokenResult unmarshals an instance of TokenResult from the specified map of raw messages.
func UnmarshalTokenResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TokenResult)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "part_of_speech", &obj.PartOfSpeech)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "location", &obj.Location)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "lemma", &obj.Lemma)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
