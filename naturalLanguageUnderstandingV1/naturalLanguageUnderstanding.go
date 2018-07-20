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
package naturalLanguageUnderstandingV1

import (
    "bytes"
    "fmt"
    "strings"
    req "github.com/parnurzeal/gorequest"
    watson "golang-sdk"
)

type NaturalLanguageUnderstandingV1 struct {
	client *watson.Client
}

func NewNaturalLanguageUnderstandingV1(creds watson.Credentials) (*NaturalLanguageUnderstandingV1, error) {
	client, clientErr := watson.NewClient(creds, "natural-language-understanding")

	if clientErr != nil {
		return nil, clientErr
	}

	return &NaturalLanguageUnderstandingV1{ client: client }, nil
}

// Analyze text, HTML, or a public webpage
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) Analyze(body *Parameters) (*watson.WatsonResponse, []error) {
    path := "/v1/analyze"
    creds := naturalLanguageUnderstanding.client.Creds
    useTM := naturalLanguageUnderstanding.client.UseTM
    tokenManager := naturalLanguageUnderstanding.client.TokenManager

    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
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

func GetAnalyzeResult(response *watson.WatsonResponse) *AnalysisResults {
    result, ok := response.Result.(*AnalysisResults)

    if ok {
        return result
    }

    return nil
}

// Delete model
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) DeleteModel(modelID string) (*watson.WatsonResponse, []error) {
    path := "/v1/models/{model_id}"
    creds := naturalLanguageUnderstanding.client.Creds
    useTM := naturalLanguageUnderstanding.client.UseTM
    tokenManager := naturalLanguageUnderstanding.client.TokenManager

    path = strings.Replace(path, "{model_id}", modelID, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
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

func GetDeleteModelResult(response *watson.WatsonResponse) *DeleteModelResults {
    result, ok := response.Result.(*DeleteModelResults)

    if ok {
        return result
    }

    return nil
}

// List models
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) ListModels() (*watson.WatsonResponse, []error) {
    path := "/v1/models"
    creds := naturalLanguageUnderstanding.client.Creds
    useTM := naturalLanguageUnderstanding.client.UseTM
    tokenManager := naturalLanguageUnderstanding.client.TokenManager

    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
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

func GetListModelsResult(response *watson.WatsonResponse) *ListModelsResults {
    result, ok := response.Result.(*ListModelsResults)

    if ok {
        return result
    }

    return nil
}


type AnalysisResults struct {

	// Language used to analyze the text.
	Language string `json:"language,omitempty"`

	// Text that was used in the analysis.
	AnalyzedText string `json:"analyzed_text,omitempty"`

	// URL that was used to retrieve HTML content.
	RetrievedUrl string `json:"retrieved_url,omitempty"`

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

type Author struct {

	// Name of the author.
	Name string `json:"name,omitempty"`
}

type CategoriesOptions struct {
}

type CategoriesResult struct {

	// The path to the category through the taxonomy hierarchy.
	Label string `json:"label,omitempty"`

	// Confidence score for the category classification. Higher values indicate greater confidence.
	Score float64 `json:"score,omitempty"`
}

type ConceptsOptions struct {

	// Maximum number of concepts to return.
	Limit int64 `json:"limit,omitempty"`
}

type ConceptsResult struct {

	// Name of the concept.
	Text string `json:"text,omitempty"`

	// Relevance score between 0 and 1. Higher scores indicate greater relevance.
	Relevance float64 `json:"relevance,omitempty"`

	// Link to the corresponding DBpedia resource.
	DbpediaResource string `json:"dbpedia_resource,omitempty"`
}

type DisambiguationResult struct {

	// Common entity name.
	Name string `json:"name,omitempty"`

	// Link to the corresponding DBpedia resource.
	DbpediaResource string `json:"dbpedia_resource,omitempty"`

	// Entity subtype information.
	Subtype []string `json:"subtype,omitempty"`
}

type DocumentEmotionResults struct {

	// An object containing the emotion results for the document.
	Emotion EmotionScores `json:"emotion,omitempty"`
}

type DocumentSentimentResults struct {

	// Indicates whether the sentiment is positive, neutral, or negative.
	Label string `json:"label,omitempty"`

	// Sentiment score from -1 (negative) to 1 (positive).
	Score float64 `json:"score,omitempty"`
}

type EmotionOptions struct {

	// Set this to false to hide document-level emotion results.
	Document bool `json:"document,omitempty"`

	// Emotion results will be returned for each target string that is found in the document.
	Targets []string `json:"targets,omitempty"`
}

type EmotionResult struct {

	// The returned emotion results across the document.
	Document DocumentEmotionResults `json:"document,omitempty"`

	// The returned emotion results per specified target.
	Targets []TargetedEmotionResults `json:"targets,omitempty"`
}

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

type EntityMention struct {

	// Entity mention text.
	Text string `json:"text,omitempty"`

	// Character offsets indicating the beginning and end of the mention in the analyzed text.
	Location []int64 `json:"location,omitempty"`
}

type FeatureSentimentResults struct {

	// Sentiment score from -1 (negative) to 1 (positive).
	Score float64 `json:"score,omitempty"`
}

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

type Feed struct {

	// URL of the RSS or ATOM feed.
	Link string `json:"link,omitempty"`
}

type inline_response_200 struct {

	// model_id of the deleted model.
	Deleted string `json:"deleted,omitempty"`
}

type KeywordsOptions struct {

	// Maximum number of keywords to return.
	Limit int64 `json:"limit,omitempty"`

	// Set this to true to return sentiment information for detected keywords.
	Sentiment bool `json:"sentiment,omitempty"`

	// Set this to true to analyze emotion for detected keywords.
	Emotion bool `json:"emotion,omitempty"`
}

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

type ListModelsResults struct {

	Models []Model `json:"models,omitempty"`
}

type MetadataOptions struct {
}

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

type Model struct {

	// Shows as available if the model is ready for use.
	Status string `json:"status,omitempty"`

	// Unique model ID.
	ModelId string `json:"model_id,omitempty"`

	// ISO 639-1 code indicating the language of the model.
	Language string `json:"language,omitempty"`

	// Model description.
	Description string `json:"description,omitempty"`
}

type Parameters struct {

	// The plain text to analyze.
	Text string `json:"text,omitempty"`

	// The HTML file to analyze.
	Html string `json:"html,omitempty"`

	// The web page to analyze.
	Url string `json:"url,omitempty"`

	// Specific features to analyze the document for.
	Features Features `json:"features"`

	// Remove website elements, such as links, ads, etc.
	Clean bool `json:"clean,omitempty"`

	// XPath query for targeting nodes in HTML.
	Xpath string `json:"xpath,omitempty"`

	// Whether to use raw HTML content if text cleaning fails.
	FallbackToRaw bool `json:"fallback_to_raw,omitempty"`

	// Whether or not to return the analyzed text.
	ReturnAnalyzedText bool `json:"return_analyzed_text,omitempty"`

	// ISO 639-1 code indicating the language to use in the analysis.
	Language string `json:"language,omitempty"`

	// Sets the maximum number of characters that are processed by the service.
	LimitTextCharacters int64 `json:"limit_text_characters,omitempty"`
}

type RelationArgument struct {

	Entities []RelationEntity `json:"entities,omitempty"`

	// Character offsets indicating the beginning and end of the mention in the analyzed text.
	Location []int64 `json:"location,omitempty"`

	// Text that corresponds to the argument.
	Text string `json:"text,omitempty"`
}

type RelationEntity struct {

	// Text that corresponds to the entity.
	Text string `json:"text,omitempty"`

	// Entity type.
	TypeVar string `json:"type_var,omitempty"`
}

type RelationsOptions struct {

	// Enter a custom model ID to override the default model.
	Model string `json:"model,omitempty"`
}

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

type SemanticRolesAction struct {

	// Analyzed text that corresponds to the action.
	Text string `json:"text,omitempty"`

	// normalized version of the action.
	Normalized string `json:"normalized,omitempty"`

	Verb SemanticRolesVerb `json:"verb,omitempty"`
}

type SemanticRolesEntity struct {

	// Entity type.
	TypeVar string `json:"type_var,omitempty"`

	// The entity text.
	Text string `json:"text,omitempty"`
}

type SemanticRolesKeyword struct {

	// The keyword text.
	Text string `json:"text,omitempty"`
}

type SemanticRolesObject struct {

	// Object text.
	Text string `json:"text,omitempty"`

	Keywords []SemanticRolesKeyword `json:"keywords,omitempty"`
}

type SemanticRolesOptions struct {

	// Maximum number of semantic_roles results to return.
	Limit int64 `json:"limit,omitempty"`

	// Set this to true to return keyword information for subjects and objects.
	Keywords bool `json:"keywords,omitempty"`

	// Set this to true to return entity information for subjects and objects.
	Entities bool `json:"entities,omitempty"`
}

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

type SemanticRolesSubject struct {

	// Text that corresponds to the subject role.
	Text string `json:"text,omitempty"`

	Entities []SemanticRolesEntity `json:"entities,omitempty"`

	Keywords []SemanticRolesKeyword `json:"keywords,omitempty"`
}

type SemanticRolesVerb struct {

	// The keyword text.
	Text string `json:"text,omitempty"`

	// Verb tense.
	Tense string `json:"tense,omitempty"`
}

type SentimentOptions struct {

	// Set this to false to hide document-level sentiment results.
	Document bool `json:"document,omitempty"`

	// Sentiment results will be returned for each target string that is found in the document.
	Targets []string `json:"targets,omitempty"`
}

type SentimentResult struct {

	// The document level sentiment.
	Document DocumentSentimentResults `json:"document,omitempty"`

	// The targeted sentiment to analyze.
	Targets []TargetedSentimentResults `json:"targets,omitempty"`
}

type TargetedEmotionResults struct {

	// Targeted text.
	Text string `json:"text,omitempty"`

	// An object containing the emotion results for the target.
	Emotion EmotionScores `json:"emotion,omitempty"`
}

type TargetedSentimentResults struct {

	// Targeted text.
	Text string `json:"text,omitempty"`

	// Sentiment score from -1 (negative) to 1 (positive).
	Score float64 `json:"score,omitempty"`
}

type Usage struct {

	// Number of features used in the API call.
	Features int64 `json:"features,omitempty"`

	// Number of text characters processed.
	TextCharacters int64 `json:"text_characters,omitempty"`

	// Number of 10,000-character units processed.
	TextUnits int64 `json:"text_units,omitempty"`
}
