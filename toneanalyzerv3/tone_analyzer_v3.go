// Package toneanalyzerv3 : Operations and models for the ToneAnalyzerV3 service
package toneanalyzerv3

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
	"strings"
)

// ToneAnalyzerV3 : The IBM Watson&trade; Tone Analyzer service uses linguistic analysis to detect emotional and
// language tones in written text. The service can analyze tone at both the document and sentence levels. You can use
// the service to understand how your written communications are perceived and then to improve the tone of your
// communications. Businesses can use the service to learn the tone of their customers' communications and to respond to
// each customer appropriately, or to understand and improve their customer conversations.
//
// **Note:** Request logging is disabled for the Tone Analyzer service. Regardless of whether you set the
// `X-Watson-Learning-Opt-Out` request header, the service does not log or retain data from requests and responses.
//
// Version: V3
// See: http://www.ibm.com/watson/developercloud/tone-analyzer.html
type ToneAnalyzerV3 struct {
	Service *core.BaseService
}

// ToneAnalyzerV3Options : Service options
type ToneAnalyzerV3Options struct {
	Version        string
	URL            string
	Username       string
	Password       string
	IAMApiKey      string
	IAMAccessToken string
	IAMURL         string
}

// NewToneAnalyzerV3 : Instantiate ToneAnalyzerV3
func NewToneAnalyzerV3(options *ToneAnalyzerV3Options) (*ToneAnalyzerV3, error) {
	if options.URL == "" {
		options.URL = "https://gateway.watsonplatform.net/tone-analyzer/api"
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
	service, serviceErr := core.NewBaseService(serviceOptions, "tone_analyzer", "Tone Analyzer")
	if serviceErr != nil {
		return nil, serviceErr
	}

	return &ToneAnalyzerV3{Service: service}, nil
}

// Tone : Analyze general tone
// Use the general purpose endpoint to analyze the tone of your input content. The service analyzes the content for
// emotional and language tones. The method always analyzes the tone of the full document; by default, it also analyzes
// the tone of each individual sentence of the content.
//
// You can submit no more than 128 KB of total input content and no more than 1000 individual sentences in JSON, plain
// text, or HTML format. The service analyzes the first 1000 sentences for document-level analysis and only the first
// 100 sentences for sentence-level analysis.
//
// Per the JSON specification, the default character encoding for JSON content is effectively always UTF-8; per the HTTP
// specification, the default encoding for plain text and HTML is ISO-8859-1 (effectively, the ASCII character set).
// When specifying a content type of plain text or HTML, include the `charset` parameter to indicate the character
// encoding of the input text; for example: `Content-Type: text/plain;charset=utf-8`. For `text/html`, the service
// removes HTML tags and analyzes only the textual content.
//
// **See also:** [Using the general-purpose
// endpoint](https://cloud.ibm.com/docs/services/tone-analyzer/using-tone.html#using-the-general-purpose-endpoint).
func (toneAnalyzer *ToneAnalyzerV3) Tone(toneOptions *ToneOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(toneOptions, "toneOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(toneOptions, "toneOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v3/tone"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(toneAnalyzer.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range toneOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("tone_analyzer", "V3", "Tone")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	if toneOptions.ContentLanguage != nil {
		builder.AddHeader("Content-Language", fmt.Sprint(*toneOptions.ContentLanguage))
	}
	if toneOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*toneOptions.AcceptLanguage))
	}
	if toneOptions.ContentType != nil {
		builder.AddHeader("Content-Type", fmt.Sprint(*toneOptions.ContentType))
	}

	if toneOptions.Sentences != nil {
		builder.AddQuery("sentences", fmt.Sprint(*toneOptions.Sentences))
	}
	if toneOptions.Tones != nil {
		builder.AddQuery("tones", strings.Join(toneOptions.Tones, ","))
	}
	builder.AddQuery("version", toneAnalyzer.Service.Options.Version)

	_, err := builder.SetBodyContent(core.StringNilMapper(toneOptions.ContentType), toneOptions.ToneInput, nil, toneOptions.Body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := toneAnalyzer.Service.Request(request, new(ToneAnalysis))
	return response, err
}

// GetToneResult : Retrieve result of Tone operation
func (toneAnalyzer *ToneAnalyzerV3) GetToneResult(response *core.DetailedResponse) *ToneAnalysis {
	result, ok := response.Result.(*ToneAnalysis)
	if ok {
		return result
	}
	return nil
}

// ToneChat : Analyze customer engagement tone
// Use the customer engagement endpoint to analyze the tone of customer service and customer support conversations. For
// each utterance of a conversation, the method reports the most prevalent subset of the following seven tones: sad,
// frustrated, satisfied, excited, polite, impolite, and sympathetic.
//
// If you submit more than 50 utterances, the service returns a warning for the overall content and analyzes only the
// first 50 utterances. If you submit a single utterance that contains more than 500 characters, the service returns an
// error for that utterance and does not analyze the utterance. The request fails if all utterances have more than 500
// characters. Per the JSON specification, the default character encoding for JSON content is effectively always UTF-8.
//
// **See also:** [Using the customer-engagement
// endpoint](https://cloud.ibm.com/docs/services/tone-analyzer/using-tone-chat.html#using-the-customer-engagement-endpoint).
func (toneAnalyzer *ToneAnalyzerV3) ToneChat(toneChatOptions *ToneChatOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(toneChatOptions, "toneChatOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(toneChatOptions, "toneChatOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v3/tone_chat"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(toneAnalyzer.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range toneChatOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("tone_analyzer", "V3", "ToneChat")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if toneChatOptions.ContentLanguage != nil {
		builder.AddHeader("Content-Language", fmt.Sprint(*toneChatOptions.ContentLanguage))
	}
	if toneChatOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*toneChatOptions.AcceptLanguage))
	}
	builder.AddQuery("version", toneAnalyzer.Service.Options.Version)

	body := make(map[string]interface{})
	if toneChatOptions.Utterances != nil {
		body["utterances"] = toneChatOptions.Utterances
	}
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := toneAnalyzer.Service.Request(request, new(UtteranceAnalyses))
	return response, err
}

// GetToneChatResult : Retrieve result of ToneChat operation
func (toneAnalyzer *ToneAnalyzerV3) GetToneChatResult(response *core.DetailedResponse) *UtteranceAnalyses {
	result, ok := response.Result.(*UtteranceAnalyses)
	if ok {
		return result
	}
	return nil
}

// DocumentAnalysis : An object of type `DocumentAnalysis` that provides the results of the analysis for the full input document.
type DocumentAnalysis struct {

	// **`2017-09-21`:** An array of `ToneScore` objects that provides the results of the analysis for each qualifying tone
	// of the document. The array includes results for any tone whose score is at least 0.5. The array is empty if no tone
	// has a score that meets this threshold. **`2016-05-19`:** Not returned.
	Tones []ToneScore `json:"tones,omitempty"`

	// **`2017-09-21`:** Not returned. **`2016-05-19`:** An array of `ToneCategory` objects that provides the results of
	// the tone analysis for the full document of the input content. The service returns results only for the tones
	// specified with the `tones` parameter of the request.
	ToneCategories []ToneCategory `json:"tone_categories,omitempty"`

	// **`2017-09-21`:** A warning message if the overall content exceeds 128 KB or contains more than 1000 sentences. The
	// service analyzes only the first 1000 sentences for document-level analysis and the first 100 sentences for
	// sentence-level analysis. **`2016-05-19`:** Not returned.
	Warning *string `json:"warning,omitempty"`
}

// SentenceAnalysis : SentenceAnalysis struct
type SentenceAnalysis struct {

	// The unique identifier of a sentence of the input content. The first sentence has ID 0, and the ID of each subsequent
	// sentence is incremented by one.
	SentenceID *int64 `json:"sentence_id" validate:"required"`

	// The text of the input sentence.
	Text *string `json:"text" validate:"required"`

	// **`2017-09-21`:** An array of `ToneScore` objects that provides the results of the analysis for each qualifying tone
	// of the sentence. The array includes results for any tone whose score is at least 0.5. The array is empty if no tone
	// has a score that meets this threshold. **`2016-05-19`:** Not returned.
	Tones []ToneScore `json:"tones,omitempty"`

	// **`2017-09-21`:** Not returned. **`2016-05-19`:** An array of `ToneCategory` objects that provides the results of
	// the tone analysis for the sentence. The service returns results only for the tones specified with the `tones`
	// parameter of the request.
	ToneCategories []ToneCategory `json:"tone_categories,omitempty"`

	// **`2017-09-21`:** Not returned. **`2016-05-19`:** The offset of the first character of the sentence in the overall
	// input content.
	InputFrom *int64 `json:"input_from,omitempty"`

	// **`2017-09-21`:** Not returned. **`2016-05-19`:** The offset of the last character of the sentence in the overall
	// input content.
	InputTo *int64 `json:"input_to,omitempty"`
}

// ToneAnalysis : ToneAnalysis struct
type ToneAnalysis struct {

	// An object of type `DocumentAnalysis` that provides the results of the analysis for the full input document.
	DocumentTone *DocumentAnalysis `json:"document_tone" validate:"required"`

	// An array of `SentenceAnalysis` objects that provides the results of the analysis for the individual sentences of the
	// input content. The service returns results only for the first 100 sentences of the input. The field is omitted if
	// the `sentences` parameter of the request is set to `false`.
	SentencesTone []SentenceAnalysis `json:"sentences_tone,omitempty"`
}

// ToneCategory : ToneCategory struct
type ToneCategory struct {

	// An array of `ToneScore` objects that provides the results for the tones of the category.
	Tones []ToneScore `json:"tones" validate:"required"`

	// The unique, non-localized identifier of the category for the results. The service can return results for the
	// following category IDs: `emotion_tone`, `language_tone`, and `social_tone`.
	CategoryID *string `json:"category_id" validate:"required"`

	// The user-visible, localized name of the category.
	CategoryName *string `json:"category_name" validate:"required"`
}

// ToneChatOptions : The toneChat options.
type ToneChatOptions struct {

	// An array of `Utterance` objects that provides the input content that the service is to analyze.
	Utterances []Utterance `json:"utterances" validate:"required"`

	// The language of the input text for the request: English or French. Regional variants are treated as their parent
	// language; for example, `en-US` is interpreted as `en`. The input content must match the specified language. Do not
	// submit content that contains both languages. You can use different languages for **Content-Language** and
	// **Accept-Language**.
	// * **`2017-09-21`:** Accepts `en` or `fr`.
	// * **`2016-05-19`:** Accepts only `en`.
	ContentLanguage *string `json:"Content-Language,omitempty"`

	// The desired language of the response. For two-character arguments, regional variants are treated as their parent
	// language; for example, `en-US` is interpreted as `en`. You can use different languages for **Content-Language** and
	// **Accept-Language**.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the ToneChatOptions.ContentLanguage property.
// The language of the input text for the request: English or French. Regional variants are treated as their parent
// language; for example, `en-US` is interpreted as `en`. The input content must match the specified language. Do not
// submit content that contains both languages. You can use different languages for **Content-Language** and
// **Accept-Language**.
// * **`2017-09-21`:** Accepts `en` or `fr`.
// * **`2016-05-19`:** Accepts only `en`.
const (
	ToneChatOptions_ContentLanguage_En = "en"
	ToneChatOptions_ContentLanguage_Fr = "fr"
)

// Constants associated with the ToneChatOptions.AcceptLanguage property.
// The desired language of the response. For two-character arguments, regional variants are treated as their parent
// language; for example, `en-US` is interpreted as `en`. You can use different languages for **Content-Language** and
// **Accept-Language**.
const (
	ToneChatOptions_AcceptLanguage_Ar = "ar"
	ToneChatOptions_AcceptLanguage_De = "de"
	ToneChatOptions_AcceptLanguage_En = "en"
	ToneChatOptions_AcceptLanguage_Es = "es"
	ToneChatOptions_AcceptLanguage_Fr = "fr"
	ToneChatOptions_AcceptLanguage_It = "it"
	ToneChatOptions_AcceptLanguage_Ja = "ja"
	ToneChatOptions_AcceptLanguage_Ko = "ko"
	ToneChatOptions_AcceptLanguage_PtBr = "pt-br"
	ToneChatOptions_AcceptLanguage_ZhCn = "zh-cn"
	ToneChatOptions_AcceptLanguage_ZhTw = "zh-tw"
)

// NewToneChatOptions : Instantiate ToneChatOptions
func (toneAnalyzer *ToneAnalyzerV3) NewToneChatOptions(utterances []Utterance) *ToneChatOptions {
	return &ToneChatOptions{
		Utterances: utterances,
	}
}

// SetUtterances : Allow user to set Utterances
func (options *ToneChatOptions) SetUtterances(utterances []Utterance) *ToneChatOptions {
	options.Utterances = utterances
	return options
}

// SetContentLanguage : Allow user to set ContentLanguage
func (options *ToneChatOptions) SetContentLanguage(contentLanguage string) *ToneChatOptions {
	options.ContentLanguage = core.StringPtr(contentLanguage)
	return options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (options *ToneChatOptions) SetAcceptLanguage(acceptLanguage string) *ToneChatOptions {
	options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ToneChatOptions) SetHeaders(param map[string]string) *ToneChatOptions {
	options.Headers = param
	return options
}

// ToneChatScore : ToneChatScore struct
type ToneChatScore struct {

	// The score for the tone in the range of 0.5 to 1. A score greater than 0.75 indicates a high likelihood that the tone
	// is perceived in the utterance.
	Score *float64 `json:"score" validate:"required"`

	// The unique, non-localized identifier of the tone for the results. The service returns results only for tones whose
	// scores meet a minimum threshold of 0.5.
	ToneID *string `json:"tone_id" validate:"required"`

	// The user-visible, localized name of the tone.
	ToneName *string `json:"tone_name" validate:"required"`
}

// Constants associated with the ToneChatScore.ToneID property.
// The unique, non-localized identifier of the tone for the results. The service returns results only for tones whose
// scores meet a minimum threshold of 0.5.
const (
	ToneChatScore_ToneID_Excited = "excited"
	ToneChatScore_ToneID_Frustrated = "frustrated"
	ToneChatScore_ToneID_Impolite = "impolite"
	ToneChatScore_ToneID_Polite = "polite"
	ToneChatScore_ToneID_Sad = "sad"
	ToneChatScore_ToneID_Satisfied = "satisfied"
	ToneChatScore_ToneID_Sympathetic = "sympathetic"
)

// ToneInput : ToneInput struct
type ToneInput struct {

	// The input content that the service is to analyze.
	Text *string `json:"text" validate:"required"`
}

// ToneOptions : The tone options.
type ToneOptions struct {

	// JSON, plain text, or HTML input that contains the content to be analyzed. For JSON input, provide an object of type
	// `ToneInput`.
	ToneInput *ToneInput `json:"tone_input,omitempty"`

	// JSON, plain text, or HTML input that contains the content to be analyzed. For JSON input, provide an object of type
	// `ToneInput`.
	Body *string `json:"body,omitempty"`

	// Indicates whether the service is to return an analysis of each individual sentence in addition to its analysis of
	// the full document. If `true` (the default), the service returns results for each sentence.
	Sentences *bool `json:"sentences,omitempty"`

	// **`2017-09-21`:** Deprecated. The service continues to accept the parameter for backward-compatibility, but the
	// parameter no longer affects the response.
	//
	// **`2016-05-19`:** A comma-separated list of tones for which the service is to return its analysis of the input; the
	// indicated tones apply both to the full document and to individual sentences of the document. You can specify one or
	// more of the valid values. Omit the parameter to request results for all three tones.
	Tones []string `json:"tones,omitempty"`

	// The language of the input text for the request: English or French. Regional variants are treated as their parent
	// language; for example, `en-US` is interpreted as `en`. The input content must match the specified language. Do not
	// submit content that contains both languages. You can use different languages for **Content-Language** and
	// **Accept-Language**.
	// * **`2017-09-21`:** Accepts `en` or `fr`.
	// * **`2016-05-19`:** Accepts only `en`.
	ContentLanguage *string `json:"Content-Language,omitempty"`

	// The desired language of the response. For two-character arguments, regional variants are treated as their parent
	// language; for example, `en-US` is interpreted as `en`. You can use different languages for **Content-Language** and
	// **Accept-Language**.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// The type of the input. A character encoding can be specified by including a `charset` parameter. For example,
	// 'text/plain;charset=utf-8'.
	ContentType *string `json:"Content-Type,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the ToneOptions.Tone property.
const (
	ToneOptions_Tone_Emotion = "emotion"
	ToneOptions_Tone_Language = "language"
	ToneOptions_Tone_Social = "social"
)

// Constants associated with the ToneOptions.ContentLanguage property.
// The language of the input text for the request: English or French. Regional variants are treated as their parent
// language; for example, `en-US` is interpreted as `en`. The input content must match the specified language. Do not
// submit content that contains both languages. You can use different languages for **Content-Language** and
// **Accept-Language**.
// * **`2017-09-21`:** Accepts `en` or `fr`.
// * **`2016-05-19`:** Accepts only `en`.
const (
	ToneOptions_ContentLanguage_En = "en"
	ToneOptions_ContentLanguage_Fr = "fr"
)

// Constants associated with the ToneOptions.AcceptLanguage property.
// The desired language of the response. For two-character arguments, regional variants are treated as their parent
// language; for example, `en-US` is interpreted as `en`. You can use different languages for **Content-Language** and
// **Accept-Language**.
const (
	ToneOptions_AcceptLanguage_Ar = "ar"
	ToneOptions_AcceptLanguage_De = "de"
	ToneOptions_AcceptLanguage_En = "en"
	ToneOptions_AcceptLanguage_Es = "es"
	ToneOptions_AcceptLanguage_Fr = "fr"
	ToneOptions_AcceptLanguage_It = "it"
	ToneOptions_AcceptLanguage_Ja = "ja"
	ToneOptions_AcceptLanguage_Ko = "ko"
	ToneOptions_AcceptLanguage_PtBr = "pt-br"
	ToneOptions_AcceptLanguage_ZhCn = "zh-cn"
	ToneOptions_AcceptLanguage_ZhTw = "zh-tw"
)

// Constants associated with the ToneOptions.ContentType property.
// The type of the input. A character encoding can be specified by including a `charset` parameter. For example,
// 'text/plain;charset=utf-8'.
const (
	ToneOptions_ContentType_ApplicationJSON = "application/json"
	ToneOptions_ContentType_TextHTML = "text/html"
	ToneOptions_ContentType_TextPlain = "text/plain"
)

// NewToneOptions : Instantiate ToneOptions
func (toneAnalyzer *ToneAnalyzerV3) NewToneOptions() *ToneOptions {
	return &ToneOptions{}
}

// SetToneInput : Allow user to set ToneInput
func (options *ToneOptions) SetToneInput(toneInput *ToneInput) *ToneOptions {
	options.ToneInput = toneInput
	return options
}

// SetBody : Allow user to set Body
func (options *ToneOptions) SetBody(body string) *ToneOptions {
	options.Body = core.StringPtr(body)
	return options
}

// SetSentences : Allow user to set Sentences
func (options *ToneOptions) SetSentences(sentences bool) *ToneOptions {
	options.Sentences = core.BoolPtr(sentences)
	return options
}

// SetTones : Allow user to set Tones
func (options *ToneOptions) SetTones(tones []string) *ToneOptions {
	options.Tones = tones
	return options
}

// SetContentLanguage : Allow user to set ContentLanguage
func (options *ToneOptions) SetContentLanguage(contentLanguage string) *ToneOptions {
	options.ContentLanguage = core.StringPtr(contentLanguage)
	return options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (options *ToneOptions) SetAcceptLanguage(acceptLanguage string) *ToneOptions {
	options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return options
}

// SetContentType : Allow user to set ContentType
func (options *ToneOptions) SetContentType(contentType string) *ToneOptions {
	options.ContentType = core.StringPtr(contentType)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ToneOptions) SetHeaders(param map[string]string) *ToneOptions {
	options.Headers = param
	return options
}

// ToneScore : ToneScore struct
type ToneScore struct {

	// The score for the tone.
	// * **`2017-09-21`:** The score that is returned lies in the range of 0.5 to 1. A score greater than 0.75 indicates a
	// high likelihood that the tone is perceived in the content.
	// * **`2016-05-19`:** The score that is returned lies in the range of 0 to 1. A score less than 0.5 indicates that the
	// tone is unlikely to be perceived in the content; a score greater than 0.75 indicates a high likelihood that the tone
	// is perceived.
	Score *float64 `json:"score" validate:"required"`

	// The unique, non-localized identifier of the tone.
	// * **`2017-09-21`:** The service can return results for the following tone IDs: `anger`, `fear`, `joy`, and `sadness`
	// (emotional tones); `analytical`, `confident`, and `tentative` (language tones). The service returns results only for
	// tones whose scores meet a minimum threshold of 0.5.
	// * **`2016-05-19`:** The service can return results for the following tone IDs of the different categories: for the
	// `emotion` category: `anger`, `disgust`, `fear`, `joy`, and `sadness`; for the `language` category: `analytical`,
	// `confident`, and `tentative`; for the `social` category: `openness_big5`, `conscientiousness_big5`,
	// `extraversion_big5`, `agreeableness_big5`, and `emotional_range_big5`. The service returns scores for all tones of a
	// category, regardless of their values.
	ToneID *string `json:"tone_id" validate:"required"`

	// The user-visible, localized name of the tone.
	ToneName *string `json:"tone_name" validate:"required"`
}

// Utterance : Utterance struct
type Utterance struct {

	// An utterance contributed by a user in the conversation that is to be analyzed. The utterance can contain multiple
	// sentences.
	Text *string `json:"text" validate:"required"`

	// A string that identifies the user who contributed the utterance specified by the `text` parameter.
	User *string `json:"user,omitempty"`
}

// UtteranceAnalyses : UtteranceAnalyses struct
type UtteranceAnalyses struct {

	// An array of `UtteranceAnalysis` objects that provides the results for each utterance of the input.
	UtterancesTone []UtteranceAnalysis `json:"utterances_tone" validate:"required"`

	// **`2017-09-21`:** A warning message if the content contains more than 50 utterances. The service analyzes only the
	// first 50 utterances. **`2016-05-19`:** Not returned.
	Warning *string `json:"warning,omitempty"`
}

// UtteranceAnalysis : UtteranceAnalysis struct
type UtteranceAnalysis struct {

	// The unique identifier of the utterance. The first utterance has ID 0, and the ID of each subsequent utterance is
	// incremented by one.
	UtteranceID *int64 `json:"utterance_id" validate:"required"`

	// The text of the utterance.
	UtteranceText *string `json:"utterance_text" validate:"required"`

	// An array of `ToneChatScore` objects that provides results for the most prevalent tones of the utterance. The array
	// includes results for any tone whose score is at least 0.5. The array is empty if no tone has a score that meets this
	// threshold.
	Tones []ToneChatScore `json:"tones" validate:"required"`

	// **`2017-09-21`:** An error message if the utterance contains more than 500 characters. The service does not analyze
	// the utterance. **`2016-05-19`:** Not returned.
	Error *string `json:"error,omitempty"`
}
