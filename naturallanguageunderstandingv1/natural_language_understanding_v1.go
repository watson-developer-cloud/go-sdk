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

// Package naturallanguageunderstandingv1 : Operations and models for the NaturalLanguageUnderstandingV1 service
package naturallanguageunderstandingv1

import (
	"fmt"
	"github.com/IBM/go-sdk-core/core"
	"github.com/go-openapi/strfmt"
	common "github.com/watson-developer-cloud/go-sdk/common"
)

// NaturalLanguageUnderstandingV1 : Analyze various features of text content at scale. Provide text, raw HTML, or a
// public URL and IBM Watson Natural Language Understanding will give you results for the features you request. The
// service cleans HTML content before analysis by default, so the results can ignore most advertisements and other
// unwanted content.
//
// You can create [custom
// models](https://cloud.ibm.com/docs/services/natural-language-understanding?topic=natural-language-understanding-customizing)
// with Watson Knowledge Studio to detect custom entities, relations, and categories in Natural Language Understanding.
//
// Version: 1.0
// See: https://cloud.ibm.com/docs/services/natural-language-understanding/
type NaturalLanguageUnderstandingV1 struct {
	Service *core.BaseService
	Version string
}

const defaultServiceURL = "https://gateway.watsonplatform.net/natural-language-understanding/api"

// NaturalLanguageUnderstandingV1Options : Service options
type NaturalLanguageUnderstandingV1Options struct {
	URL           string
	Authenticator core.Authenticator
	Version       string
}

// NewNaturalLanguageUnderstandingV1 : Instantiate NaturalLanguageUnderstandingV1
func NewNaturalLanguageUnderstandingV1(options *NaturalLanguageUnderstandingV1Options) (service *NaturalLanguageUnderstandingV1, err error) {
	if options.URL == "" {
		options.URL = defaultServiceURL
	}

	serviceOptions := &core.ServiceOptions{
		URL:           options.URL,
		Authenticator: options.Authenticator,
	}

	if serviceOptions.Authenticator == nil {
		serviceOptions.Authenticator, err = core.GetAuthenticatorFromEnvironment("natural-language-understanding")
		if err != nil {
			return
		}
	}

	baseService, err := core.NewBaseService(serviceOptions, "natural-language-understanding", "Natural Language Understanding")
	if err != nil {
		return
	}

	service = &NaturalLanguageUnderstandingV1{
		Service: baseService,
		Version: options.Version,
	}

	return
}

// SetServiceURL sets the service URL
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) SetServiceURL(url string) error {
	return naturalLanguageUnderstanding.Service.SetServiceURL(url)
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
// - Syntax (Experimental).
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) Analyze(analyzeOptions *AnalyzeOptions) (result *AnalysisResults, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(analyzeOptions, "analyzeOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(analyzeOptions, "analyzeOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/analyze"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(naturalLanguageUnderstanding.Service.Options.URL, pathSegments, pathParameters)
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
	builder.AddQuery("version", naturalLanguageUnderstanding.Version)

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

	response, err = naturalLanguageUnderstanding.Service.Request(request, new(AnalysisResults))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*AnalysisResults)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// ListModels : List models
// Lists Watson Knowledge Studio [custom entities and relations
// models](https://cloud.ibm.com/docs/services/natural-language-understanding?topic=natural-language-understanding-customizing)
// that are deployed to your Natural Language Understanding service.
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) ListModels(listModelsOptions *ListModelsOptions) (result *ListModelsResults, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listModelsOptions, "listModelsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/models"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(naturalLanguageUnderstanding.Service.Options.URL, pathSegments, pathParameters)
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
	builder.AddQuery("version", naturalLanguageUnderstanding.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = naturalLanguageUnderstanding.Service.Request(request, new(ListModelsResults))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*ListModelsResults)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// DeleteModel : Delete model
// Deletes a custom model.
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) DeleteModel(deleteModelOptions *DeleteModelOptions) (result *DeleteModelResults, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteModelOptions, "deleteModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteModelOptions, "deleteModelOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/models"}
	pathParameters := []string{*deleteModelOptions.ModelID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(naturalLanguageUnderstanding.Service.Options.URL, pathSegments, pathParameters)
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
	builder.AddQuery("version", naturalLanguageUnderstanding.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = naturalLanguageUnderstanding.Service.Request(request, new(DeleteModelResults))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*DeleteModelResults)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

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
	Metadata *AnalysisResultsMetadata `json:"metadata,omitempty"`

	// The relationships between entities in the content.
	Relations []RelationsResult `json:"relations,omitempty"`

	// Sentences parsed into `subject`, `action`, and `object` form.
	SemanticRoles []SemanticRolesResult `json:"semantic_roles,omitempty"`

	// The sentiment of the content.
	Sentiment *SentimentResult `json:"sentiment,omitempty"`

	// Tokens and sentences returned from syntax analysis.
	Syntax *SyntaxResult `json:"syntax,omitempty"`
}

// AnalysisResultsMetadata : Webpage metadata, such as the author and the title of the page.
type AnalysisResultsMetadata struct {

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

// AnalysisResultsUsage : API usage information for the request.
type AnalysisResultsUsage struct {

	// Number of features used in the API call.
	Features *int64 `json:"features,omitempty"`

	// Number of text characters processed.
	TextCharacters *int64 `json:"text_characters,omitempty"`

	// Number of 10,000-character units processed.
	TextUnits *int64 `json:"text_units,omitempty"`
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

	// Set this to `false` to disable webpage cleaning. To learn more about webpage cleaning, see the [Analyzing
	// webpages](https://cloud.ibm.com/docs/services/natural-language-understanding?topic=natural-language-understanding-analyzing-webpages)
	// documentation.
	Clean *bool `json:"clean,omitempty"`

	// An [XPath
	// query](https://cloud.ibm.com/docs/services/natural-language-understanding?topic=natural-language-understanding-analyzing-webpages#xpath)
	// to perform on `html` or `url` input. Results of the query will be appended to the cleaned webpage text before it is
	// analyzed. To analyze only the results of the XPath query, set the `clean` parameter to `false`.
	Xpath *string `json:"xpath,omitempty"`

	// Whether to use raw HTML content if text cleaning fails.
	FallbackToRaw *bool `json:"fallback_to_raw,omitempty"`

	// Whether or not to return the analyzed text.
	ReturnAnalyzedText *bool `json:"return_analyzed_text,omitempty"`

	// ISO 639-1 code that specifies the language of your text. This overrides automatic language detection. Language
	// support differs depending on the features you include in your analysis. See [Language
	// support](https://cloud.ibm.com/docs/services/natural-language-understanding?topic=natural-language-understanding-language-support)
	// for more information.
	Language *string `json:"language,omitempty"`

	// Sets the maximum number of characters that are processed by the service.
	LimitTextCharacters *int64 `json:"limit_text_characters,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewAnalyzeOptions : Instantiate AnalyzeOptions
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) NewAnalyzeOptions(features *Features) *AnalyzeOptions {
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
func (options *AnalyzeOptions) SetHTML(HTML string) *AnalyzeOptions {
	options.HTML = core.StringPtr(HTML)
	return options
}

// SetURL : Allow user to set URL
func (options *AnalyzeOptions) SetURL(URL string) *AnalyzeOptions {
	options.URL = core.StringPtr(URL)
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
	// model](https://cloud.ibm.com/docs/services/natural-language-understanding?topic=natural-language-understanding-customizing)
	// ID to override the standard categories model.
	Model *string `json:"model,omitempty"`
}

// CategoriesRelevantText : Relevant text that contributed to the categorization.
type CategoriesRelevantText struct {

	// Text from the analyzed source that supports the categorization.
	Text *string `json:"text,omitempty"`
}

// CategoriesResult : A categorization of the analyzed text.
type CategoriesResult struct {

	// The path to the category through the 5-level taxonomy hierarchy. For the complete list of categories, see the
	// [Categories
	// hierarchy](https://cloud.ibm.com/docs/services/natural-language-understanding?topic=natural-language-understanding-categories#categories-hierarchy)
	// documentation.
	Label *string `json:"label,omitempty"`

	// Confidence score for the category classification. Higher values indicate greater confidence.
	Score *float64 `json:"score,omitempty"`

	// Information that helps to explain what contributed to the categories result.
	Explanation *CategoriesResultExplanation `json:"explanation,omitempty"`
}

// CategoriesResultExplanation : Information that helps to explain what contributed to the categories result.
type CategoriesResultExplanation struct {

	// An array of relevant text from the source that contributed to the categorization. The sorted array begins with the
	// phrase that contributed most significantly to the result, followed by phrases that were less and less impactful.
	RelevantText []CategoriesRelevantText `json:"relevant_text,omitempty"`
}

// ConceptsOptions : Returns high-level concepts in the content. For example, a research paper about deep learning might return the
// concept, "Artificial Intelligence" although the term is not mentioned.
//
// Supported languages: English, French, German, Italian, Japanese, Korean, Portuguese, Spanish.
type ConceptsOptions struct {

	// Maximum number of concepts to return.
	Limit *int64 `json:"limit,omitempty"`
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

// DeleteModelOptions : The DeleteModel options.
type DeleteModelOptions struct {

	// Model ID of the model to delete.
	ModelID *string `json:"model_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteModelOptions : Instantiate DeleteModelOptions
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) NewDeleteModelOptions(modelID string) *DeleteModelOptions {
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

// DisambiguationResult : Disambiguation information for the entity.
type DisambiguationResult struct {

	// Common entity name.
	Name *string `json:"name,omitempty"`

	// Link to the corresponding DBpedia resource.
	DbpediaResource *string `json:"dbpedia_resource,omitempty"`

	// Entity subtype information.
	Subtype []string `json:"subtype,omitempty"`
}

// DocumentEmotionResults : Emotion results for the document as a whole.
type DocumentEmotionResults struct {

	// Emotion results for the document as a whole.
	Emotion *EmotionScores `json:"emotion,omitempty"`
}

// DocumentSentimentResults : DocumentSentimentResults struct
type DocumentSentimentResults struct {

	// Indicates whether the sentiment is positive, neutral, or negative.
	Label *string `json:"label,omitempty"`

	// Sentiment score from -1 (negative) to 1 (positive).
	Score *float64 `json:"score,omitempty"`
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

// EmotionResult : The detected anger, disgust, fear, joy, or sadness that is conveyed by the content. Emotion information can be
// returned for detected entities, keywords, or user-specified target phrases found in the text.
type EmotionResult struct {

	// Emotion results for the document as a whole.
	Document *DocumentEmotionResults `json:"document,omitempty"`

	// Emotion results for specified targets.
	Targets []TargetedEmotionResults `json:"targets,omitempty"`
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

// EntitiesOptions : Identifies people, cities, organizations, and other entities in the content. See [Entity types and
// subtypes](https://cloud.ibm.com/docs/services/natural-language-understanding?topic=natural-language-understanding-entity-types).
//
// Supported languages: English, French, German, Italian, Japanese, Korean, Portuguese, Russian, Spanish, Swedish.
// Arabic, Chinese, and Dutch are supported only through custom models.
type EntitiesOptions struct {

	// Maximum number of entities to return.
	Limit *int64 `json:"limit,omitempty"`

	// Set this to `true` to return locations of entity mentions.
	Mentions *bool `json:"mentions,omitempty"`

	// Enter a [custom
	// model](https://cloud.ibm.com/docs/services/natural-language-understanding?topic=natural-language-understanding-customizing)
	// ID to override the standard entity detection model.
	Model *string `json:"model,omitempty"`

	// Set this to `true` to return sentiment information for detected entities.
	Sentiment *bool `json:"sentiment,omitempty"`

	// Set this to `true` to analyze emotion for detected keywords.
	Emotion *bool `json:"emotion,omitempty"`
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

// FeatureSentimentResults : FeatureSentimentResults struct
type FeatureSentimentResults struct {

	// Sentiment score from -1 (negative) to 1 (positive).
	Score *float64 `json:"score,omitempty"`
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

	// Identifies people, cities, organizations, and other entities in the content. See [Entity types and
	// subtypes](https://cloud.ibm.com/docs/services/natural-language-understanding?topic=natural-language-understanding-entity-types).
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
	Metadata *MetadataOptions `json:"metadata,omitempty"`

	// Recognizes when two entities are related and identifies the type of relation. For example, an `awardedTo` relation
	// might connect the entities "Nobel Prize" and "Albert Einstein". See [Relation
	// types](https://cloud.ibm.com/docs/services/natural-language-understanding?topic=natural-language-understanding-relations).
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

// Feed : RSS or ATOM feed found on the webpage.
type Feed struct {

	// URL of the RSS or ATOM feed.
	Link *string `json:"link,omitempty"`
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

// ListModelsOptions : The ListModels options.
type ListModelsOptions struct {

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListModelsOptions : Instantiate ListModelsOptions
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) NewListModelsOptions() *ListModelsOptions {
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

// MetadataOptions : Returns information from the document, including author name, title, RSS/ATOM feeds, prominent page image, and
// publication date. Supports URL and HTML input types only.
type MetadataOptions struct {
}

// Model : Model struct
type Model struct {

	// When the status is `available`, the model is ready to use.
	Status *string `json:"status,omitempty"`

	// Unique model ID.
	ModelID *string `json:"model_id,omitempty"`

	// ISO 639-1 code indicating the language of the model.
	Language *string `json:"language,omitempty"`

	// Model description.
	Description *string `json:"description,omitempty"`

	// ID of the Watson Knowledge Studio workspace that deployed this model to Natural Language Understanding.
	WorkspaceID *string `json:"workspace_id,omitempty"`

	// The model version, if it was manually provided in Watson Knowledge Studio.
	Version *string `json:"version,omitempty"`

	// The description of the version, if it was manually provided in Watson Knowledge Studio.
	VersionDescription *string `json:"version_description,omitempty"`

	// A dateTime indicating when the model was created.
	Created *strfmt.DateTime `json:"created,omitempty"`
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

// RelationEntity : An entity that corresponds with an argument in a relation.
type RelationEntity struct {

	// Text that corresponds to the entity.
	Text *string `json:"text,omitempty"`

	// Entity type.
	Type *string `json:"type,omitempty"`
}

// RelationsOptions : Recognizes when two entities are related and identifies the type of relation. For example, an `awardedTo` relation
// might connect the entities "Nobel Prize" and "Albert Einstein". See [Relation
// types](https://cloud.ibm.com/docs/services/natural-language-understanding?topic=natural-language-understanding-relations).
//
// Supported languages: Arabic, English, German, Japanese, Korean, Spanish. Chinese, Dutch, French, Italian, and
// Portuguese custom models are also supported.
type RelationsOptions struct {

	// Enter a [custom
	// model](https://cloud.ibm.com/docs/services/natural-language-understanding?topic=natural-language-understanding-customizing)
	// ID to override the default model.
	Model *string `json:"model,omitempty"`
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

// SemanticRolesEntity : SemanticRolesEntity struct
type SemanticRolesEntity struct {

	// Entity type.
	Type *string `json:"type,omitempty"`

	// The entity text.
	Text *string `json:"text,omitempty"`
}

// SemanticRolesKeyword : SemanticRolesKeyword struct
type SemanticRolesKeyword struct {

	// The keyword text.
	Text *string `json:"text,omitempty"`
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

// SemanticRolesResultAction : The extracted action from the sentence.
type SemanticRolesResultAction struct {

	// Analyzed text that corresponds to the action.
	Text *string `json:"text,omitempty"`

	// normalized version of the action.
	Normalized *string `json:"normalized,omitempty"`

	Verb *SemanticRolesVerb `json:"verb,omitempty"`
}

// SemanticRolesResultObject : The extracted object from the sentence.
type SemanticRolesResultObject struct {

	// Object text.
	Text *string `json:"text,omitempty"`

	// An array of extracted keywords.
	Keywords []SemanticRolesKeyword `json:"keywords,omitempty"`
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

// SemanticRolesVerb : SemanticRolesVerb struct
type SemanticRolesVerb struct {

	// The keyword text.
	Text *string `json:"text,omitempty"`

	// Verb tense.
	Tense *string `json:"tense,omitempty"`
}

// SentenceResult : SentenceResult struct
type SentenceResult struct {

	// The sentence.
	Text *string `json:"text,omitempty"`

	// Character offsets indicating the beginning and end of the sentence in the analyzed text.
	Location []int64 `json:"location,omitempty"`
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

// SentimentResult : The sentiment of the content.
type SentimentResult struct {

	// The document level sentiment.
	Document *DocumentSentimentResults `json:"document,omitempty"`

	// The targeted sentiment to analyze.
	Targets []TargetedSentimentResults `json:"targets,omitempty"`
}

// SyntaxOptions : Returns tokens and sentences from the input text.
type SyntaxOptions struct {

	// Tokenization options.
	Tokens *SyntaxOptionsTokens `json:"tokens,omitempty"`

	// Set this to `true` to return sentence information.
	Sentences *bool `json:"sentences,omitempty"`
}

// SyntaxOptionsTokens : Tokenization options.
type SyntaxOptionsTokens struct {

	// Set this to `true` to return the lemma for each token.
	Lemma *bool `json:"lemma,omitempty"`

	// Set this to `true` to return the part of speech for each token.
	PartOfSpeech *bool `json:"part_of_speech,omitempty"`
}

// SyntaxResult : Tokens and sentences returned from syntax analysis.
type SyntaxResult struct {
	Tokens []TokenResult `json:"tokens,omitempty"`

	Sentences []SentenceResult `json:"sentences,omitempty"`
}

// TargetedEmotionResults : Emotion results for a specified target.
type TargetedEmotionResults struct {

	// Targeted text.
	Text *string `json:"text,omitempty"`

	// The emotion results for the target.
	Emotion *EmotionScores `json:"emotion,omitempty"`
}

// TargetedSentimentResults : TargetedSentimentResults struct
type TargetedSentimentResults struct {

	// Targeted text.
	Text *string `json:"text,omitempty"`

	// Sentiment score from -1 (negative) to 1 (positive).
	Score *float64 `json:"score,omitempty"`
}

// TokenResult : TokenResult struct
type TokenResult struct {

	// The token as it appears in the analyzed text.
	Text *string `json:"text,omitempty"`

	// The part of speech of the token. For descriptions of the values, see [Universal Dependencies POS
	// tags](https://universaldependencies.org/u/pos/).
	PartOfSpeech *string `json:"part_of_speech,omitempty"`

	// Character offsets indicating the beginning and end of the token in the analyzed text.
	Location []int64 `json:"location,omitempty"`

	// The [lemma](https://wikipedia.org/wiki/Lemma_%28morphology%29) of the token.
	Lemma *string `json:"lemma,omitempty"`
}

// Constants associated with the TokenResult.PartOfSpeech property.
// The part of speech of the token. For descriptions of the values, see [Universal Dependencies POS
// tags](https://universaldependencies.org/u/pos/).
const (
	TokenResult_PartOfSpeech_Adj   = "ADJ"
	TokenResult_PartOfSpeech_Adp   = "ADP"
	TokenResult_PartOfSpeech_Adv   = "ADV"
	TokenResult_PartOfSpeech_Aux   = "AUX"
	TokenResult_PartOfSpeech_Cconj = "CCONJ"
	TokenResult_PartOfSpeech_Det   = "DET"
	TokenResult_PartOfSpeech_Intj  = "INTJ"
	TokenResult_PartOfSpeech_Noun  = "NOUN"
	TokenResult_PartOfSpeech_Num   = "NUM"
	TokenResult_PartOfSpeech_Part  = "PART"
	TokenResult_PartOfSpeech_Pron  = "PRON"
	TokenResult_PartOfSpeech_Propn = "PROPN"
	TokenResult_PartOfSpeech_Punct = "PUNCT"
	TokenResult_PartOfSpeech_Sconj = "SCONJ"
	TokenResult_PartOfSpeech_Sym   = "SYM"
	TokenResult_PartOfSpeech_Verb  = "VERB"
	TokenResult_PartOfSpeech_X     = "X"
)
