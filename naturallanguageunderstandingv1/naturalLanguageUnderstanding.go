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
    "bytes"
    "fmt"
    "strings"
    req "github.com/parnurzeal/gorequest"
    watson "golang-sdk"
)

// NaturalLanguageUnderstandingV1 : The NaturalLanguageUnderstandingV1 service
type NaturalLanguageUnderstandingV1 struct {
	client *watson.Client
}

// NewNaturalLanguageUnderstandingV1 : Instantiate NaturalLanguageUnderstandingV1
func NewNaturalLanguageUnderstandingV1(creds watson.Credentials) (*NaturalLanguageUnderstandingV1, error) {
    if creds.ServiceURL == "" {
        creds.ServiceURL = "https://gateway.watsonplatform.net/natural-language-understanding/api"
    }

	client, clientErr := watson.NewClient(creds, "natural-language-understanding")

	if clientErr != nil {
		return nil, clientErr
	}

	return &NaturalLanguageUnderstandingV1{ client: client }, nil
}

// Analyze : Analyze text, HTML, or a public webpage
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) Analyze(options *AnalyzeOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/analyze"
    creds := naturalLanguageUnderstanding.client.Creds
    useTM := naturalLanguageUnderstanding.client.UseTM
    tokenManager := naturalLanguageUnderstanding.client.TokenManager

    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    body := map[string]interface{}{}
    if options.IsTextSet {
        body["text"] = options.Text
    }
    if options.IsHTMLSet {
        body["html"] = options.HTML
    }
    if options.IsURLSet {
        body["url"] = options.URL
    }
    body["features"] = options.Features
    if options.IsCleanSet {
        body["clean"] = options.Clean
    }
    if options.IsXpathSet {
        body["xpath"] = options.Xpath
    }
    if options.IsFallbackToRawSet {
        body["fallback_to_raw"] = options.FallbackToRaw
    }
    if options.IsReturnAnalyzedTextSet {
        body["return_analyzed_text"] = options.ReturnAnalyzedText
    }
    if options.IsLanguageSet {
        body["language"] = options.Language
    }
    if options.IsLimitTextCharactersSet {
        body["limit_text_characters"] = options.LimitTextCharacters
    }
    request.Send(body)

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(AnalysisResults)
    res, _, err := request.EndStruct(&response.Result)

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetAnalyzeResult : Cast result of Analyze operation
func GetAnalyzeResult(response *watson.WatsonResponse) *AnalysisResults {
    result, ok := response.Result.(*AnalysisResults)

    if ok {
        return result
    }

    return nil
}

// DeleteModel : Delete model
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) DeleteModel(options *DeleteModelOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/models/{model_id}"
    creds := naturalLanguageUnderstanding.client.Creds
    useTM := naturalLanguageUnderstanding.client.UseTM
    tokenManager := naturalLanguageUnderstanding.client.TokenManager

    path = strings.Replace(path, "{model_id}", options.ModelID, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(DeleteModelResults)
    res, _, err := request.EndStruct(&response.Result)

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetDeleteModelResult : Cast result of DeleteModel operation
func GetDeleteModelResult(response *watson.WatsonResponse) *DeleteModelResults {
    result, ok := response.Result.(*DeleteModelResults)

    if ok {
        return result
    }

    return nil
}

// ListModels : List models
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) ListModels(options *ListModelsOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/models"
    creds := naturalLanguageUnderstanding.client.Creds
    useTM := naturalLanguageUnderstanding.client.UseTM
    tokenManager := naturalLanguageUnderstanding.client.TokenManager

    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(ListModelsResults)
    res, _, err := request.EndStruct(&response.Result)

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetListModelsResult : Cast result of ListModels operation
func GetListModelsResult(response *watson.WatsonResponse) *ListModelsResults {
    result, ok := response.Result.(*ListModelsResults)

    if ok {
        return result
    }

    return nil
}


// AnalysisResults : Results of the analysis, organized by feature.
type AnalysisResults struct {

	// Language used to analyze the text.
	Language string `json:"language,omitempty"`

	// Text that was used in the analysis.
	AnalyzedText string `json:"analyzed_text,omitempty"`

	// URL that was used to retrieve HTML content.
	RetrievedURL string `json:"retrieved_url,omitempty"`

	// API usage information for the request.
	Usage Usage `json:"usage,omitempty"`

	// The general concepts referenced or alluded to in the specified content.
	Concepts []ConceptsResult `json:"concepts,omitempty"`

	// The important entities in the specified content.
	Entities []EntitiesResult `json:"entities,omitempty"`

	// The important keywords in content organized by relevance.
	Keywords []KeywordsResult `json:"keywords,omitempty"`

	// The hierarchical 5-level taxonomy the content is categorized into.
	Categories []CategoriesResult `json:"categories,omitempty"`

	// The anger, disgust, fear, joy, or sadness conveyed by the content.
	Emotion EmotionResult `json:"emotion,omitempty"`

	// The metadata holds author information, publication date and the title of the text/HTML content.
	Metadata MetadataResult `json:"metadata,omitempty"`

	// The relationships between entities in the content.
	Relations []RelationsResult `json:"relations,omitempty"`

	// The subjects of actions and the objects the actions act upon.
	SemanticRoles []SemanticRolesResult `json:"semantic_roles,omitempty"`

	// The sentiment of the content.
	Sentiment SentimentResult `json:"sentiment,omitempty"`
}

// AnalyzeOptions : The analyze options.
type AnalyzeOptions struct {

	// The plain text to analyze.
	Text string `json:"text,omitempty"`

    // Indicates whether user set optional parameter Text
    IsTextSet bool

	// The HTML file to analyze.
	HTML string `json:"html,omitempty"`

    // Indicates whether user set optional parameter HTML
    IsHTMLSet bool

	// The web page to analyze.
	URL string `json:"url,omitempty"`

    // Indicates whether user set optional parameter URL
    IsURLSet bool

	// Specific features to analyze the document for.
	Features Features `json:"features"`

	// Remove website elements, such as links, ads, etc.
	Clean bool `json:"clean,omitempty"`

    // Indicates whether user set optional parameter Clean
    IsCleanSet bool

	// XPath query for targeting nodes in HTML.
	Xpath string `json:"xpath,omitempty"`

    // Indicates whether user set optional parameter Xpath
    IsXpathSet bool

	// Whether to use raw HTML content if text cleaning fails.
	FallbackToRaw bool `json:"fallback_to_raw,omitempty"`

    // Indicates whether user set optional parameter FallbackToRaw
    IsFallbackToRawSet bool

	// Whether or not to return the analyzed text.
	ReturnAnalyzedText bool `json:"return_analyzed_text,omitempty"`

    // Indicates whether user set optional parameter ReturnAnalyzedText
    IsReturnAnalyzedTextSet bool

	// ISO 639-1 code indicating the language to use in the analysis.
	Language string `json:"language,omitempty"`

    // Indicates whether user set optional parameter Language
    IsLanguageSet bool

	// Sets the maximum number of characters that are processed by the service.
	LimitTextCharacters int64 `json:"limit_text_characters,omitempty"`

    // Indicates whether user set optional parameter LimitTextCharacters
    IsLimitTextCharactersSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewAnalyzeOptions : Instantiate AnalyzeOptions
func NewAnalyzeOptions(features Features) *AnalyzeOptions {
    return &AnalyzeOptions{
        Features: features,
    }
}

// SetText : Allow user to set Text
func (options *AnalyzeOptions) SetText(param string) *AnalyzeOptions {
    options.Text = param
    options.IsTextSet = true
    return options
}

// SetHTML : Allow user to set HTML
func (options *AnalyzeOptions) SetHTML(param string) *AnalyzeOptions {
    options.HTML = param
    options.IsHTMLSet = true
    return options
}

// SetURL : Allow user to set URL
func (options *AnalyzeOptions) SetURL(param string) *AnalyzeOptions {
    options.URL = param
    options.IsURLSet = true
    return options
}

// SetFeatures : Allow user to set Features
func (options *AnalyzeOptions) SetFeatures(param Features) *AnalyzeOptions {
    options.Features = param
    return options
}

// SetClean : Allow user to set Clean
func (options *AnalyzeOptions) SetClean(param bool) *AnalyzeOptions {
    options.Clean = param
    options.IsCleanSet = true
    return options
}

// SetXpath : Allow user to set Xpath
func (options *AnalyzeOptions) SetXpath(param string) *AnalyzeOptions {
    options.Xpath = param
    options.IsXpathSet = true
    return options
}

// SetFallbackToRaw : Allow user to set FallbackToRaw
func (options *AnalyzeOptions) SetFallbackToRaw(param bool) *AnalyzeOptions {
    options.FallbackToRaw = param
    options.IsFallbackToRawSet = true
    return options
}

// SetReturnAnalyzedText : Allow user to set ReturnAnalyzedText
func (options *AnalyzeOptions) SetReturnAnalyzedText(param bool) *AnalyzeOptions {
    options.ReturnAnalyzedText = param
    options.IsReturnAnalyzedTextSet = true
    return options
}

// SetLanguage : Allow user to set Language
func (options *AnalyzeOptions) SetLanguage(param string) *AnalyzeOptions {
    options.Language = param
    options.IsLanguageSet = true
    return options
}

// SetLimitTextCharacters : Allow user to set LimitTextCharacters
func (options *AnalyzeOptions) SetLimitTextCharacters(param int64) *AnalyzeOptions {
    options.LimitTextCharacters = param
    options.IsLimitTextCharactersSet = true
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
	Name string `json:"name,omitempty"`
}

// CategoriesOptions : The hierarchical 5-level taxonomy the content is categorized into.
type CategoriesOptions struct {
}

// CategoriesResult : The hierarchical 5-level taxonomy the content is categorized into.
type CategoriesResult struct {

	// The path to the category through the taxonomy hierarchy.
	Label string `json:"label,omitempty"`

	// Confidence score for the category classification. Higher values indicate greater confidence.
	Score float64 `json:"score,omitempty"`
}

// ConceptsOptions : Whether or not to analyze content for general concepts that are referenced or alluded to.
type ConceptsOptions struct {

	// Maximum number of concepts to return.
	Limit int64 `json:"limit,omitempty"`
}

// ConceptsResult : The general concepts referenced or alluded to in the specified content.
type ConceptsResult struct {

	// Name of the concept.
	Text string `json:"text,omitempty"`

	// Relevance score between 0 and 1. Higher scores indicate greater relevance.
	Relevance float64 `json:"relevance,omitempty"`

	// Link to the corresponding DBpedia resource.
	DbpediaResource string `json:"dbpedia_resource,omitempty"`
}

// DeleteModelOptions : The deleteModel options.
type DeleteModelOptions struct {

	// model_id of the model to delete.
	ModelID string `json:"model_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewDeleteModelOptions : Instantiate DeleteModelOptions
func NewDeleteModelOptions(modelID string) *DeleteModelOptions {
    return &DeleteModelOptions{
        ModelID: modelID,
    }
}

// SetModelID : Allow user to set ModelID
func (options *DeleteModelOptions) SetModelID(param string) *DeleteModelOptions {
    options.ModelID = param
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
	Deleted string `json:"deleted,omitempty"`
}

// DisambiguationResult : Disambiguation information for the entity.
type DisambiguationResult struct {

	// Common entity name.
	Name string `json:"name,omitempty"`

	// Link to the corresponding DBpedia resource.
	DbpediaResource string `json:"dbpedia_resource,omitempty"`

	// Entity subtype information.
	Subtype []string `json:"subtype,omitempty"`
}

// DocumentEmotionResults : An object containing the emotion results of a document.
type DocumentEmotionResults struct {

	// An object containing the emotion results for the document.
	Emotion EmotionScores `json:"emotion,omitempty"`
}

// DocumentSentimentResults : DocumentSentimentResults struct
type DocumentSentimentResults struct {

	// Indicates whether the sentiment is positive, neutral, or negative.
	Label string `json:"label,omitempty"`

	// Sentiment score from -1 (negative) to 1 (positive).
	Score float64 `json:"score,omitempty"`
}

// EmotionOptions : Whether or not to return emotion analysis of the content.
type EmotionOptions struct {

	// Set this to false to hide document-level emotion results.
	Document bool `json:"document,omitempty"`

	// Emotion results will be returned for each target string that is found in the document.
	Targets []string `json:"targets,omitempty"`
}

// EmotionResult : The detected anger, disgust, fear, joy, or sadness that is conveyed by the content. Emotion information can be returned for detected entities, keywords, or user-specified target phrases found in the text.
type EmotionResult struct {

	// The returned emotion results across the document.
	Document DocumentEmotionResults `json:"document,omitempty"`

	// The returned emotion results per specified target.
	Targets []TargetedEmotionResults `json:"targets,omitempty"`
}

// EmotionScores : EmotionScores struct
type EmotionScores struct {

	// Anger score from 0 to 1. A higher score means that the text is more likely to convey anger.
	Anger float64 `json:"anger,omitempty"`

	// Disgust score from 0 to 1. A higher score means that the text is more likely to convey disgust.
	Disgust float64 `json:"disgust,omitempty"`

	// Fear score from 0 to 1. A higher score means that the text is more likely to convey fear.
	Fear float64 `json:"fear,omitempty"`

	// Joy score from 0 to 1. A higher score means that the text is more likely to convey joy.
	Joy float64 `json:"joy,omitempty"`

	// Sadness score from 0 to 1. A higher score means that the text is more likely to convey sadness.
	Sadness float64 `json:"sadness,omitempty"`
}

// EntitiesOptions : Whether or not to return important people, places, geopolitical, and other entities detected in the analyzed content.
type EntitiesOptions struct {

	// Maximum number of entities to return.
	Limit int64 `json:"limit,omitempty"`

	// Set this to true to return locations of entity mentions.
	Mentions bool `json:"mentions,omitempty"`

	// Enter a custom model ID to override the standard entity detection model.
	Model string `json:"model,omitempty"`

	// Set this to true to return sentiment information for detected entities.
	Sentiment bool `json:"sentiment,omitempty"`

	// Set this to true to analyze emotion for detected keywords.
	Emotion bool `json:"emotion,omitempty"`
}

// EntitiesResult : The important people, places, geopolitical entities and other types of entities in your content.
type EntitiesResult struct {

	// Entity type.
	TypeVar string `json:"type_var,omitempty"`

	// The name of the entity.
	Text string `json:"text,omitempty"`

	// Relevance score from 0 to 1. Higher values indicate greater relevance.
	Relevance float64 `json:"relevance,omitempty"`

	// Entity mentions and locations.
	Mentions []EntityMention `json:"mentions,omitempty"`

	// How many times the entity was mentioned in the text.
	Count int64 `json:"count,omitempty"`

	// Emotion analysis results for the entity, enabled with the "emotion" option.
	Emotion EmotionScores `json:"emotion,omitempty"`

	// Sentiment analysis results for the entity, enabled with the "sentiment" option.
	Sentiment FeatureSentimentResults `json:"sentiment,omitempty"`

	// Disambiguation information for the entity.
	Disambiguation DisambiguationResult `json:"disambiguation,omitempty"`
}

// EntityMention : EntityMention struct
type EntityMention struct {

	// Entity mention text.
	Text string `json:"text,omitempty"`

	// Character offsets indicating the beginning and end of the mention in the analyzed text.
	Location []int64 `json:"location,omitempty"`
}

// FeatureSentimentResults : FeatureSentimentResults struct
type FeatureSentimentResults struct {

	// Sentiment score from -1 (negative) to 1 (positive).
	Score float64 `json:"score,omitempty"`
}

// Features : Analysis features and options.
type Features struct {

	// Whether or not to return the concepts that are mentioned in the analyzed text.
	Concepts ConceptsOptions `json:"concepts,omitempty"`

	// Whether or not to extract the emotions implied in the analyzed text.
	Emotion EmotionOptions `json:"emotion,omitempty"`

	// Whether or not to extract detected entity objects from the analyzed text.
	Entities EntitiesOptions `json:"entities,omitempty"`

	// Whether or not to return the keywords in the analyzed text.
	Keywords KeywordsOptions `json:"keywords,omitempty"`

	// Whether or not the author, publication date, and title of the analyzed text should be returned. This parameter is only available for URL and HTML input.
	Metadata MetadataOptions `json:"metadata,omitempty"`

	// Whether or not to return the relationships between detected entities in the analyzed text.
	Relations RelationsOptions `json:"relations,omitempty"`

	// Whether or not to return the subject-action-object relations from the analyzed text.
	SemanticRoles SemanticRolesOptions `json:"semantic_roles,omitempty"`

	// Whether or not to return the overall sentiment of the analyzed text.
	Sentiment SentimentOptions `json:"sentiment,omitempty"`

	// Whether or not to return the high level category the content is categorized as (i.e. news, art).
	Categories CategoriesOptions `json:"categories,omitempty"`
}

// Feed : RSS or ATOM feed found on the webpage.
type Feed struct {

	// URL of the RSS or ATOM feed.
	Link string `json:"link,omitempty"`
}

// KeywordsOptions : An option indicating whether or not important keywords from the analyzed content should be returned.
type KeywordsOptions struct {

	// Maximum number of keywords to return.
	Limit int64 `json:"limit,omitempty"`

	// Set this to true to return sentiment information for detected keywords.
	Sentiment bool `json:"sentiment,omitempty"`

	// Set this to true to analyze emotion for detected keywords.
	Emotion bool `json:"emotion,omitempty"`
}

// KeywordsResult : The most important keywords in the content, organized by relevance.
type KeywordsResult struct {

	// Relevance score from 0 to 1. Higher values indicate greater relevance.
	Relevance float64 `json:"relevance,omitempty"`

	// The keyword text.
	Text string `json:"text,omitempty"`

	// Emotion analysis results for the keyword, enabled with the "emotion" option.
	Emotion EmotionScores `json:"emotion,omitempty"`

	// Sentiment analysis results for the keyword, enabled with the "sentiment" option.
	Sentiment FeatureSentimentResults `json:"sentiment,omitempty"`
}

// ListModelsOptions : The listModels options.
type ListModelsOptions struct {

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewListModelsOptions : Instantiate ListModelsOptions
func NewListModelsOptions() *ListModelsOptions {
    return &ListModelsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListModelsOptions) SetHeaders(param map[string]string) *ListModelsOptions {
    options.Headers = param
    return options
}

// ListModelsResults : Models available for Relations and Entities features.
type ListModelsResults struct {

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
	PublicationDate string `json:"publication_date,omitempty"`

	// The title of the document.
	Title string `json:"title,omitempty"`

	// URL of a prominent image on the webpage.
	Image string `json:"image,omitempty"`

	// RSS/ATOM feeds found on the webpage.
	Feeds []Feed `json:"feeds,omitempty"`
}

// Model : Model struct
type Model struct {

	// Shows as available if the model is ready for use.
	Status string `json:"status,omitempty"`

	// Unique model ID.
	ModelID string `json:"model_id,omitempty"`

	// ISO 639-1 code indicating the language of the model.
	Language string `json:"language,omitempty"`

	// Model description.
	Description string `json:"description,omitempty"`
}

// RelationArgument : RelationArgument struct
type RelationArgument struct {

	Entities []RelationEntity `json:"entities,omitempty"`

	// Character offsets indicating the beginning and end of the mention in the analyzed text.
	Location []int64 `json:"location,omitempty"`

	// Text that corresponds to the argument.
	Text string `json:"text,omitempty"`
}

// RelationEntity : An entity that corresponds with an argument in a relation.
type RelationEntity struct {

	// Text that corresponds to the entity.
	Text string `json:"text,omitempty"`

	// Entity type.
	TypeVar string `json:"type_var,omitempty"`
}

// RelationsOptions : An option specifying if the relationships found between entities in the analyzed content should be returned.
type RelationsOptions struct {

	// Enter a custom model ID to override the default model.
	Model string `json:"model,omitempty"`
}

// RelationsResult : The relations between entities found in the content.
type RelationsResult struct {

	// Confidence score for the relation. Higher values indicate greater confidence.
	Score float64 `json:"score,omitempty"`

	// The sentence that contains the relation.
	Sentence string `json:"sentence,omitempty"`

	// The type of the relation.
	TypeVar string `json:"type_var,omitempty"`

	// The extracted relation objects from the text.
	Arguments []RelationArgument `json:"arguments,omitempty"`
}

// SemanticRolesAction : SemanticRolesAction struct
type SemanticRolesAction struct {

	// Analyzed text that corresponds to the action.
	Text string `json:"text,omitempty"`

	// normalized version of the action.
	Normalized string `json:"normalized,omitempty"`

	Verb SemanticRolesVerb `json:"verb,omitempty"`
}

// SemanticRolesEntity : SemanticRolesEntity struct
type SemanticRolesEntity struct {

	// Entity type.
	TypeVar string `json:"type_var,omitempty"`

	// The entity text.
	Text string `json:"text,omitempty"`
}

// SemanticRolesKeyword : SemanticRolesKeyword struct
type SemanticRolesKeyword struct {

	// The keyword text.
	Text string `json:"text,omitempty"`
}

// SemanticRolesObject : SemanticRolesObject struct
type SemanticRolesObject struct {

	// Object text.
	Text string `json:"text,omitempty"`

	Keywords []SemanticRolesKeyword `json:"keywords,omitempty"`
}

// SemanticRolesOptions : An option specifying whether or not to identify the subjects, actions, and verbs in the analyzed content.
type SemanticRolesOptions struct {

	// Maximum number of semantic_roles results to return.
	Limit int64 `json:"limit,omitempty"`

	// Set this to true to return keyword information for subjects and objects.
	Keywords bool `json:"keywords,omitempty"`

	// Set this to true to return entity information for subjects and objects.
	Entities bool `json:"entities,omitempty"`
}

// SemanticRolesResult : The object containing the actions and the objects the actions act upon.
type SemanticRolesResult struct {

	// Sentence from the source that contains the subject, action, and object.
	Sentence string `json:"sentence,omitempty"`

	// The extracted subject from the sentence.
	Subject SemanticRolesSubject `json:"subject,omitempty"`

	// The extracted action from the sentence.
	Action SemanticRolesAction `json:"action,omitempty"`

	// The extracted object from the sentence.
	Object SemanticRolesObject `json:"object,omitempty"`
}

// SemanticRolesSubject : SemanticRolesSubject struct
type SemanticRolesSubject struct {

	// Text that corresponds to the subject role.
	Text string `json:"text,omitempty"`

	Entities []SemanticRolesEntity `json:"entities,omitempty"`

	Keywords []SemanticRolesKeyword `json:"keywords,omitempty"`
}

// SemanticRolesVerb : SemanticRolesVerb struct
type SemanticRolesVerb struct {

	// The keyword text.
	Text string `json:"text,omitempty"`

	// Verb tense.
	Tense string `json:"tense,omitempty"`
}

// SentimentOptions : An option specifying if sentiment of detected entities, keywords, or phrases should be returned.
type SentimentOptions struct {

	// Set this to false to hide document-level sentiment results.
	Document bool `json:"document,omitempty"`

	// Sentiment results will be returned for each target string that is found in the document.
	Targets []string `json:"targets,omitempty"`
}

// SentimentResult : The sentiment of the content.
type SentimentResult struct {

	// The document level sentiment.
	Document DocumentSentimentResults `json:"document,omitempty"`

	// The targeted sentiment to analyze.
	Targets []TargetedSentimentResults `json:"targets,omitempty"`
}

// TargetedEmotionResults : An object containing the emotion results for the target.
type TargetedEmotionResults struct {

	// Targeted text.
	Text string `json:"text,omitempty"`

	// An object containing the emotion results for the target.
	Emotion EmotionScores `json:"emotion,omitempty"`
}

// TargetedSentimentResults : TargetedSentimentResults struct
type TargetedSentimentResults struct {

	// Targeted text.
	Text string `json:"text,omitempty"`

	// Sentiment score from -1 (negative) to 1 (positive).
	Score float64 `json:"score,omitempty"`
}

// Usage : Usage information.
type Usage struct {

	// Number of features used in the API call.
	Features int64 `json:"features,omitempty"`

	// Number of text characters processed.
	TextCharacters int64 `json:"text_characters,omitempty"`

	// Number of 10,000-character units processed.
	TextUnits int64 `json:"text_units,omitempty"`
}
