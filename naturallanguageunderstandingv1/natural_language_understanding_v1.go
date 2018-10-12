// Package naturallanguageunderstandingv1 : Operations and models for the NaturalLanguageUnderstandingV1 service
package naturallanguageunderstandingv1

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
	core "github.com/watson-developer-cloud/go-sdk/core"
)

// NaturalLanguageUnderstandingV1 : Analyze various features of text content at scale. Provide text, raw HTML, or a
// public URL and IBM Watson Natural Language Understanding will give you results for the features you request. The
// service cleans HTML content before analysis by default, so the results can ignore most advertisements and other
// unwanted content.
//
// You can create [custom models](/docs/services/natural-language-understanding/customizing.html) with Watson Knowledge
// Studio to detect custom entities and relations in Natural Language Understanding.
//
// Version: V1
// See: http://www.ibm.com/watson/developercloud/natural-language-understanding.html
type NaturalLanguageUnderstandingV1 struct {
	Service *core.WatsonService
}

// NaturalLanguageUnderstandingV1Options : Service options
type NaturalLanguageUnderstandingV1Options struct {
	Version        string
	URL            string
	Username       string
	Password       string
	IAMApiKey      string
	IAMAccessToken string
	IAMURL         string
}

// NewNaturalLanguageUnderstandingV1 : Instantiate NaturalLanguageUnderstandingV1
func NewNaturalLanguageUnderstandingV1(options *NaturalLanguageUnderstandingV1Options) (*NaturalLanguageUnderstandingV1, error) {
	if options.URL == "" {
		options.URL = "https://gateway.watsonplatform.net/natural-language-understanding/api"
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
	service, serviceErr := core.NewWatsonService(serviceOptions, "natural-language-understanding")
	if serviceErr != nil {
		return nil, serviceErr
	}

	return &NaturalLanguageUnderstandingV1{Service: service}, nil
}

// Analyze : Analyze text, HTML, or a public webpage
// Analyzes text, HTML, or a public webpage with one or more text analysis features.
//
// ### Concepts
// Identify general concepts that are referenced or alluded to in your content. Concepts that are detected typically
// have an associated link to a DBpedia resource.
//
// ### Emotion
// Detect anger, disgust, fear, joy, or sadness that is conveyed by your content. Emotion information can be returned
// for detected entities, keywords, or user-specified target phrases found in the text.
//
// ### Entities
// Detect important people, places, geopolitical entities and other types of entities in your content. Entity detection
// recognizes consecutive coreferences of each entity. For example, analysis of the following text would count \"Barack
// Obama\" and \"He\" as the same entity:
//
// \"Barack Obama was the 44th President of the United States. He took office in January 2009.\"
//
// ### Keywords
// Determine the most important keywords in your content. Keyword phrases are organized by relevance in the results.
//
// ### Metadata
// Get author information, publication date, and the title of your text/HTML content.
//
// ### Relations
// Recognize when two entities are related, and identify the type of relation.  For example, you can identify an
// \"awardedTo\" relation between an award and its recipient.
//
// ### Semantic Roles
// Parse sentences into subject-action-object form, and identify entities and keywords that are subjects or objects of
// an action.
//
// ### Sentiment
// Determine whether your content conveys postive or negative sentiment. Sentiment information can be returned for
// detected entities, keywords, or user-specified target phrases found in the text.
//
// ### Categories
// Categorize your content into a hierarchical 5-level taxonomy. For example, \"Leonardo DiCaprio won an Oscar\" returns
// \"/art and entertainment/movies and tv/movies\" as the most confident classification.
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) Analyze(analyzeOptions *AnalyzeOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(analyzeOptions, "analyzeOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(analyzeOptions, "analyzeOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/analyze"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(naturalLanguageUnderstanding.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range analyzeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", naturalLanguageUnderstanding.Service.Options.Version)

	body := make(map[string]interface{})
	if analyzeOptions.Text != nil {
		body["text"] = analyzeOptions.Text
	}
	if analyzeOptions.HTML != nil {
		body["html"] = analyzeOptions.HTML
	}
	if analyzeOptions.URL != nil {
		body["url"] = analyzeOptions.URL
	}
	if analyzeOptions.Features != nil {
		body["features"] = analyzeOptions.Features
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
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := naturalLanguageUnderstanding.Service.Request(request, new(AnalysisResults))
	return response, err
}

// GetAnalyzeResult : Retrieve result of Analyze operation
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) GetAnalyzeResult(response *core.DetailedResponse) *AnalysisResults {
	result, ok := response.Result.(*AnalysisResults)
	if ok {
		return result
	}
	return nil
}

// DeleteModel : Delete model
// Deletes a custom model.
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) DeleteModel(deleteModelOptions *DeleteModelOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(deleteModelOptions, "deleteModelOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(deleteModelOptions, "deleteModelOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/models"}
	pathParameters := []string{*deleteModelOptions.ModelID}

	builder := core.NewRequestBuilder(core.DELETE)
	builder.ConstructHTTPURL(naturalLanguageUnderstanding.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", naturalLanguageUnderstanding.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := naturalLanguageUnderstanding.Service.Request(request, new(DeleteModelResults))
	return response, err
}

// GetDeleteModelResult : Retrieve result of DeleteModel operation
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) GetDeleteModelResult(response *core.DetailedResponse) *DeleteModelResults {
	result, ok := response.Result.(*DeleteModelResults)
	if ok {
		return result
	}
	return nil
}

// ListModels : List models
// Lists available models for Relations and Entities features, including Watson Knowledge Studio custom models that you
// have created and linked to your Natural Language Understanding service.
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) ListModels(listModelsOptions *ListModelsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateStruct(listModelsOptions, "listModelsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/models"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(naturalLanguageUnderstanding.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range listModelsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", naturalLanguageUnderstanding.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := naturalLanguageUnderstanding.Service.Request(request, new(ListModelsResults))
	return response, err
}

// GetListModelsResult : Retrieve result of ListModels operation
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) GetListModelsResult(response *core.DetailedResponse) *ListModelsResults {
	result, ok := response.Result.(*ListModelsResults)
	if ok {
		return result
	}
	return nil
}

// AnalysisResults : Results of the analysis, organized by feature.
type AnalysisResults struct {

	// Language used to analyze the text.
	Language *string `json:"language,omitempty"`

	// Text that was used in the analysis.
	AnalyzedText *string `json:"analyzed_text,omitempty"`

	// URL that was used to retrieve HTML content.
	RetrievedURL *string `json:"retrieved_url,omitempty"`

	// API usage information for the request.
	Usage *Usage `json:"usage,omitempty"`

	// The general concepts referenced or alluded to in the specified content.
	Concepts []ConceptsResult `json:"concepts,omitempty"`

	// The important entities in the specified content.
	Entities []EntitiesResult `json:"entities,omitempty"`

	// The important keywords in content organized by relevance.
	Keywords []KeywordsResult `json:"keywords,omitempty"`

	// The hierarchical 5-level taxonomy the content is categorized into.
	Categories []CategoriesResult `json:"categories,omitempty"`

	// The anger, disgust, fear, joy, or sadness conveyed by the content.
	Emotion *EmotionResult `json:"emotion,omitempty"`

	// The metadata holds author information, publication date and the title of the text/HTML content.
	Metadata *MetadataResult `json:"metadata,omitempty"`

	// The relationships between entities in the content.
	Relations []RelationsResult `json:"relations,omitempty"`

	// The subjects of actions and the objects the actions act upon.
	SemanticRoles []SemanticRolesResult `json:"semantic_roles,omitempty"`

	// The sentiment of the content.
	Sentiment *SentimentResult `json:"sentiment,omitempty"`
}

// AnalyzeOptions : The analyze options.
type AnalyzeOptions struct {

	// The plain text to analyze. One of the `text`, `html`, or `url` parameters is required.
	Text *string `json:"text,omitempty"`

	// The HTML file to analyze. One of the `text`, `html`, or `url` parameters is required.
	HTML *string `json:"html,omitempty"`

	// The web page to analyze. One of the `text`, `html`, or `url` parameters is required.
	URL *string `json:"url,omitempty"`

	// Specific features to analyze the document for.
	Features *Features `json:"features" validate:"required"`

	// Remove website elements, such as links, ads, etc.
	Clean *bool `json:"clean,omitempty"`

	// An [XPath query](https://www.w3.org/TR/xpath/) to perform on `html` or `url` input. Results of the query will be
	// appended to the cleaned webpage text before it is analyzed. To analyze only the results of the XPath query, set the
	// `clean` parameter to `false`.
	Xpath *string `json:"xpath,omitempty"`

	// Whether to use raw HTML content if text cleaning fails.
	FallbackToRaw *bool `json:"fallback_to_raw,omitempty"`

	// Whether or not to return the analyzed text.
	ReturnAnalyzedText *bool `json:"return_analyzed_text,omitempty"`

	// ISO 639-1 code that specifies the language of your text. This overrides automatic language detection. Language
	// support differs depending on the features you include in your analysis. See [Language
	// support](https://www.bluemix.net/docs/services/natural-language-understanding/language-support.html) for more
	// information.
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

// SetFeatures : Allow user to set Features
func (options *AnalyzeOptions) SetFeatures(features *Features) *AnalyzeOptions {
	options.Features = features
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

// CategoriesOptions : The hierarchical 5-level taxonomy the content is categorized into.
type CategoriesOptions struct {
}

// CategoriesResult : The hierarchical 5-level taxonomy the content is categorized into.
type CategoriesResult struct {

	// The path to the category through the taxonomy hierarchy.
	Label *string `json:"label,omitempty"`

	// Confidence score for the category classification. Higher values indicate greater confidence.
	Score *float64 `json:"score,omitempty"`
}

// ConceptsOptions : Whether or not to analyze content for general concepts that are referenced or alluded to.
type ConceptsOptions struct {

	// Maximum number of concepts to return.
	Limit *int64 `json:"limit,omitempty"`
}

// ConceptsResult : The general concepts referenced or alluded to in the specified content.
type ConceptsResult struct {

	// Name of the concept.
	Text *string `json:"text,omitempty"`

	// Relevance score between 0 and 1. Higher scores indicate greater relevance.
	Relevance *float64 `json:"relevance,omitempty"`

	// Link to the corresponding DBpedia resource.
	DbpediaResource *string `json:"dbpedia_resource,omitempty"`
}

// DeleteModelOptions : The deleteModel options.
type DeleteModelOptions struct {

	// model_id of the model to delete.
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

// DocumentEmotionResults : An object containing the emotion results of a document.
type DocumentEmotionResults struct {

	// An object containing the emotion results for the document.
	Emotion *EmotionScores `json:"emotion,omitempty"`
}

// DocumentSentimentResults : DocumentSentimentResults struct
type DocumentSentimentResults struct {

	// Indicates whether the sentiment is positive, neutral, or negative.
	Label *string `json:"label,omitempty"`

	// Sentiment score from -1 (negative) to 1 (positive).
	Score *float64 `json:"score,omitempty"`
}

// EmotionOptions : Whether or not to return emotion analysis of the content.
type EmotionOptions struct {

	// Set this to `false` to hide document-level emotion results.
	Document *bool `json:"document,omitempty"`

	// Emotion results will be returned for each target string that is found in the document.
	Targets []string `json:"targets,omitempty"`
}

// EmotionResult : The detected anger, disgust, fear, joy, or sadness that is conveyed by the content. Emotion information can be
// returned for detected entities, keywords, or user-specified target phrases found in the text.
type EmotionResult struct {

	// The returned emotion results across the document.
	Document *DocumentEmotionResults `json:"document,omitempty"`

	// The returned emotion results per specified target.
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

// EntitiesOptions : Whether or not to return important people, places, geopolitical, and other entities detected in the analyzed content.
type EntitiesOptions struct {

	// Maximum number of entities to return.
	Limit *int64 `json:"limit,omitempty"`

	// Set this to `true` to return locations of entity mentions.
	Mentions *bool `json:"mentions,omitempty"`

	// Enter a [custom model](https://www.bluemix.net/docs/services/natural-language-understanding/customizing.html) ID to
	// override the standard entity detection model.
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

	// Entity mentions and locations.
	Mentions []EntityMention `json:"mentions,omitempty"`

	// How many times the entity was mentioned in the text.
	Count *int64 `json:"count,omitempty"`

	// Emotion analysis results for the entity, enabled with the "emotion" option.
	Emotion *EmotionScores `json:"emotion,omitempty"`

	// Sentiment analysis results for the entity, enabled with the "sentiment" option.
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
}

// FeatureSentimentResults : FeatureSentimentResults struct
type FeatureSentimentResults struct {

	// Sentiment score from -1 (negative) to 1 (positive).
	Score *float64 `json:"score,omitempty"`
}

// Features : Analysis features and options.
type Features struct {

	// Whether or not to return the concepts that are mentioned in the analyzed text.
	Concepts *ConceptsOptions `json:"concepts,omitempty"`

	// Whether or not to extract the emotions implied in the analyzed text.
	Emotion *EmotionOptions `json:"emotion,omitempty"`

	// Whether or not to extract detected entity objects from the analyzed text.
	Entities *EntitiesOptions `json:"entities,omitempty"`

	// Whether or not to return the keywords in the analyzed text.
	Keywords *KeywordsOptions `json:"keywords,omitempty"`

	// Whether or not the author, publication date, and title of the analyzed text should be returned. This parameter is
	// only available for URL and HTML input.
	Metadata *MetadataOptions `json:"metadata,omitempty"`

	// Whether or not to return the relationships between detected entities in the analyzed text.
	Relations *RelationsOptions `json:"relations,omitempty"`

	// Whether or not to return the subject-action-object relations from the analyzed text.
	SemanticRoles *SemanticRolesOptions `json:"semantic_roles,omitempty"`

	// Whether or not to return the overall sentiment of the analyzed text.
	Sentiment *SentimentOptions `json:"sentiment,omitempty"`

	// Whether or not to return the high level category the content is categorized as (i.e. news, art).
	Categories *CategoriesOptions `json:"categories,omitempty"`
}

// Feed : RSS or ATOM feed found on the webpage.
type Feed struct {

	// URL of the RSS or ATOM feed.
	Link *string `json:"link,omitempty"`
}

// KeywordsOptions : An option indicating whether or not important keywords from the analyzed content should be returned.
type KeywordsOptions struct {

	// Maximum number of keywords to return.
	Limit *int64 `json:"limit,omitempty"`

	// Set this to `true` to return sentiment information for detected keywords.
	Sentiment *bool `json:"sentiment,omitempty"`

	// Set this to `true` to analyze emotion for detected keywords.
	Emotion *bool `json:"emotion,omitempty"`
}

// KeywordsResult : The most important keywords in the content, organized by relevance.
type KeywordsResult struct {

	// Relevance score from 0 to 1. Higher values indicate greater relevance.
	Relevance *float64 `json:"relevance,omitempty"`

	// The keyword text.
	Text *string `json:"text,omitempty"`

	// Emotion analysis results for the keyword, enabled with the "emotion" option.
	Emotion *EmotionScores `json:"emotion,omitempty"`

	// Sentiment analysis results for the keyword, enabled with the "sentiment" option.
	Sentiment *FeatureSentimentResults `json:"sentiment,omitempty"`
}

// ListModelsOptions : The listModels options.
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

// ListModelsResults : Models available for Relations and Entities features.
type ListModelsResults struct {

	// An array of available models.
	Models []Model `json:"models,omitempty"`
}

// MetadataOptions : The Authors, Publication Date, and Title of the document. Supports URL and HTML input types.
type MetadataOptions struct {
}

// MetadataResult : The Authors, Publication Date, and Title of the document. Supports URL and HTML input types.
type MetadataResult struct {

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

// Model : Model struct
type Model struct {

	// Shows as available if the model is ready for use.
	Status *string `json:"status,omitempty"`

	// Unique model ID.
	ModelID *string `json:"model_id,omitempty"`

	// ISO 639-1 code indicating the language of the model.
	Language *string `json:"language,omitempty"`

	// Model description.
	Description *string `json:"description,omitempty"`
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

// RelationsOptions : An option specifying if the relationships found between entities in the analyzed content should be returned.
type RelationsOptions struct {

	// Enter a [custom model](https://www.bluemix.net/docs/services/natural-language-understanding/customizing.html) ID to
	// override the default model.
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

	// The extracted relation objects from the text.
	Arguments []RelationArgument `json:"arguments,omitempty"`
}

// SemanticRolesAction : SemanticRolesAction struct
type SemanticRolesAction struct {

	// Analyzed text that corresponds to the action.
	Text *string `json:"text,omitempty"`

	// normalized version of the action.
	Normalized *string `json:"normalized,omitempty"`

	Verb *SemanticRolesVerb `json:"verb,omitempty"`
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

// SemanticRolesObject : SemanticRolesObject struct
type SemanticRolesObject struct {

	// Object text.
	Text *string `json:"text,omitempty"`

	// An array of extracted keywords.
	Keywords []SemanticRolesKeyword `json:"keywords,omitempty"`
}

// SemanticRolesOptions : An option specifying whether or not to identify the subjects, actions, and verbs in the analyzed content.
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
	Subject *SemanticRolesSubject `json:"subject,omitempty"`

	// The extracted action from the sentence.
	Action *SemanticRolesAction `json:"action,omitempty"`

	// The extracted object from the sentence.
	Object *SemanticRolesObject `json:"object,omitempty"`
}

// SemanticRolesSubject : SemanticRolesSubject struct
type SemanticRolesSubject struct {

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

// SentimentOptions : An option specifying if sentiment of detected entities, keywords, or phrases should be returned.
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

// TargetedEmotionResults : An object containing the emotion results for the target.
type TargetedEmotionResults struct {

	// Targeted text.
	Text *string `json:"text,omitempty"`

	// An object containing the emotion results for the target.
	Emotion *EmotionScores `json:"emotion,omitempty"`
}

// TargetedSentimentResults : TargetedSentimentResults struct
type TargetedSentimentResults struct {

	// Targeted text.
	Text *string `json:"text,omitempty"`

	// Sentiment score from -1 (negative) to 1 (positive).
	Score *float64 `json:"score,omitempty"`
}

// Usage : Usage information.
type Usage struct {

	// Number of features used in the API call.
	Features *int64 `json:"features,omitempty"`

	// Number of text characters processed.
	TextCharacters *int64 `json:"text_characters,omitempty"`

	// Number of 10,000-character units processed.
	TextUnits *int64 `json:"text_units,omitempty"`
}
