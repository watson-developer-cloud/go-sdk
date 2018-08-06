// Package toneAnalyzerV3 : Operations and models for the ToneAnalyzerV3 service
package toneAnalyzerV3
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
    "runtime"
    req "github.com/parnurzeal/gorequest"
    watson "go-sdk"
)

// ToneAnalyzerV3 : The ToneAnalyzerV3 service
type ToneAnalyzerV3 struct {
    client *watson.Client
}

// ServiceCredentials : Service credentials
type ServiceCredentials struct {
    ServiceURL string
    Version string
    Username string
    Password string
    APIkey string
    IAMtoken string
}

// NewToneAnalyzerV3 : Instantiate ToneAnalyzerV3
func NewToneAnalyzerV3(serviceCreds *ServiceCredentials) (*ToneAnalyzerV3, error) {
    if serviceCreds.ServiceURL == "" {
        serviceCreds.ServiceURL = "https://gateway.watsonplatform.net/tone-analyzer/api"
    }

    creds := watson.Credentials(*serviceCreds)
    client, clientErr := watson.NewClient(&creds, "tone_analyzer")

    if clientErr != nil {
        return nil, clientErr
    }

    return &ToneAnalyzerV3{ client: client }, nil
}

// Tone : Analyze general tone
func (toneAnalyzer *ToneAnalyzerV3) Tone(options *ToneOptions) (*watson.WatsonResponse, []error) {
    path := "/v3/tone"
    creds := toneAnalyzer.client.Creds
    useTM := toneAnalyzer.client.UseTM
    tokenManager := toneAnalyzer.client.TokenManager

    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", fmt.Sprint(options.ContentType))
    if options.IsContentLanguageSet {
        request.Set("Content-Language", fmt.Sprint(options.ContentLanguage))
    }
    if options.IsAcceptLanguageSet {
        request.Set("Accept-Language", fmt.Sprint(options.AcceptLanguage))
    }
    request.Query("version=" + creds.Version)
    if options.IsSentencesSet {
        request.Query("sentences=" + fmt.Sprint(options.Sentences))
    }
    if options.IsTonesSet {
        request.Query("tones=" + fmt.Sprint(options.Tones))
    }
    if options.ContentType == "application/json" {
        request.Send(options.ToneInput)
    } else {
        request.Send(options.Body)
    }

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

    response.Result = new(ToneAnalysis)
    res, _, err := request.EndStruct(&response.Result)

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetToneResult : Cast result of Tone operation
func GetToneResult(response *watson.WatsonResponse) *ToneAnalysis {
    result, ok := response.Result.(*ToneAnalysis)

    if ok {
        return result
    }

    return nil
}

// ToneChat : Analyze customer engagement tone
func (toneAnalyzer *ToneAnalyzerV3) ToneChat(options *ToneChatOptions) (*watson.WatsonResponse, []error) {
    path := "/v3/tone_chat"
    creds := toneAnalyzer.client.Creds
    useTM := toneAnalyzer.client.UseTM
    tokenManager := toneAnalyzer.client.TokenManager

    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    if options.IsContentLanguageSet {
        request.Set("Content-Language", fmt.Sprint(options.ContentLanguage))
    }
    if options.IsAcceptLanguageSet {
        request.Set("Accept-Language", fmt.Sprint(options.AcceptLanguage))
    }
    request.Query("version=" + creds.Version)
    body := map[string]interface{}{}
    body["utterances"] = options.Utterances
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

    response.Result = new(UtteranceAnalyses)
    res, _, err := request.EndStruct(&response.Result)

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetToneChatResult : Cast result of ToneChat operation
func GetToneChatResult(response *watson.WatsonResponse) *UtteranceAnalyses {
    result, ok := response.Result.(*UtteranceAnalyses)

    if ok {
        return result
    }

    return nil
}


// DocumentAnalysis : DocumentAnalysis struct
type DocumentAnalysis struct {

	// **`2017-09-21`:** An array of `ToneScore` objects that provides the results of the analysis for each qualifying tone of the document. The array includes results for any tone whose score is at least 0.5. The array is empty if no tone has a score that meets this threshold. **`2016-05-19`:** Not returned.
	Tones []ToneScore `json:"tones,omitempty"`

	// **`2017-09-21`:** Not returned. **`2016-05-19`:** An array of `ToneCategory` objects that provides the results of the tone analysis for the full document of the input content. The service returns results only for the tones specified with the `tones` parameter of the request.
	ToneCategories []ToneCategory `json:"tone_categories,omitempty"`

	// **`2017-09-21`:** A warning message if the overall content exceeds 128 KB or contains more than 1000 sentences. The service analyzes only the first 1000 sentences for document-level analysis and the first 100 sentences for sentence-level analysis. **`2016-05-19`:** Not returned.
	Warning string `json:"warning,omitempty"`
}

// SentenceAnalysis : SentenceAnalysis struct
type SentenceAnalysis struct {

	// The unique identifier of a sentence of the input content. The first sentence has ID 0, and the ID of each subsequent sentence is incremented by one.
	SentenceID int64 `json:"sentence_id"`

	// The text of the input sentence.
	Text string `json:"text"`

	// **`2017-09-21`:** An array of `ToneScore` objects that provides the results of the analysis for each qualifying tone of the sentence. The array includes results for any tone whose score is at least 0.5. The array is empty if no tone has a score that meets this threshold. **`2016-05-19`:** Not returned.
	Tones []ToneScore `json:"tones,omitempty"`

	// **`2017-09-21`:** Not returned. **`2016-05-19`:** An array of `ToneCategory` objects that provides the results of the tone analysis for the sentence. The service returns results only for the tones specified with the `tones` parameter of the request.
	ToneCategories []ToneCategory `json:"tone_categories,omitempty"`

	// **`2017-09-21`:** Not returned. **`2016-05-19`:** The offset of the first character of the sentence in the overall input content.
	InputFrom int64 `json:"input_from,omitempty"`

	// **`2017-09-21`:** Not returned. **`2016-05-19`:** The offset of the last character of the sentence in the overall input content.
	InputTo int64 `json:"input_to,omitempty"`
}

// ToneAnalysis : ToneAnalysis struct
type ToneAnalysis struct {

	// An object of type `DocumentAnalysis` that provides the results of the analysis for the full input document.
	DocumentTone DocumentAnalysis `json:"document_tone"`

	// An array of `SentenceAnalysis` objects that provides the results of the analysis for the individual sentences of the input content. The service returns results only for the first 100 sentences of the input. The field is omitted if the `sentences` parameter of the request is set to `false`.
	SentencesTone []SentenceAnalysis `json:"sentences_tone,omitempty"`
}

// ToneCategory : ToneCategory struct
type ToneCategory struct {

	// An array of `ToneScore` objects that provides the results for the tones of the category.
	Tones []ToneScore `json:"tones"`

	// The unique, non-localized identifier of the category for the results. The service can return results for the following category IDs: `emotion_tone`, `language_tone`, and `social_tone`.
	CategoryID string `json:"category_id"`

	// The user-visible, localized name of the category.
	CategoryName string `json:"category_name"`
}

// ToneChatOptions : The toneChat options.
type ToneChatOptions struct {

	// An array of `Utterance` objects that provides the input content that the service is to analyze.
	Utterances []Utterance `json:"utterances"`

	// The language of the input text for the request: English or French. Regional variants are treated as their parent language; for example, `en-US` is interpreted as `en`. The input content must match the specified language. Do not submit content that contains both languages. You can use different languages for **Content-Language** and **Accept-Language**. * **`2017-09-21`:** Accepts `en` or `fr`. * **`2016-05-19`:** Accepts only `en`.
	ContentLanguage string `json:"Content-Language,omitempty"`

    // Indicates whether user set optional parameter ContentLanguage
    IsContentLanguageSet bool

	// The desired language of the response. For two-character arguments, regional variants are treated as their parent language; for example, `en-US` is interpreted as `en`. You can use different languages for **Content-Language** and **Accept-Language**.
	AcceptLanguage string `json:"Accept-Language,omitempty"`

    // Indicates whether user set optional parameter AcceptLanguage
    IsAcceptLanguageSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewToneChatOptions : Instantiate ToneChatOptions
func NewToneChatOptions(utterances []Utterance) *ToneChatOptions {
    return &ToneChatOptions{
        Utterances: utterances,
    }
}

// SetUtterances : Allow user to set Utterances
func (options *ToneChatOptions) SetUtterances(param []Utterance) *ToneChatOptions {
    options.Utterances = param
    return options
}

// SetContentLanguage : Allow user to set ContentLanguage
func (options *ToneChatOptions) SetContentLanguage(param string) *ToneChatOptions {
    options.ContentLanguage = param
    options.IsContentLanguageSet = true
    return options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (options *ToneChatOptions) SetAcceptLanguage(param string) *ToneChatOptions {
    options.AcceptLanguage = param
    options.IsAcceptLanguageSet = true
    return options
}

// SetHeaders : Allow user to set Headers
func (options *ToneChatOptions) SetHeaders(param map[string]string) *ToneChatOptions {
    options.Headers = param
    return options
}

// ToneChatScore : ToneChatScore struct
type ToneChatScore struct {

	// The score for the tone in the range of 0.5 to 1. A score greater than 0.75 indicates a high likelihood that the tone is perceived in the utterance.
	Score float64 `json:"score"`

	// The unique, non-localized identifier of the tone for the results. The service can return results for the following tone IDs: `sad`, `frustrated`, `satisfied`, `excited`, `polite`, `impolite`, and `sympathetic`. The service returns results only for tones whose scores meet a minimum threshold of 0.5.
	ToneID string `json:"tone_id"`

	// The user-visible, localized name of the tone.
	ToneName string `json:"tone_name"`
}

// ToneInput : ToneInput struct
type ToneInput struct {

	// The input content that the service is to analyze.
	Text string `json:"text"`
}

// ToneOptions : The tone options.
type ToneOptions struct {

	// JSON, plain text, or HTML input that contains the content to be analyzed. For JSON input, provide an object of type `ToneInput`.
	ToneInput ToneInput `json:"tone_input,omitempty"`

    // Indicates whether user set optional parameter ToneInput
    IsToneInputSet bool

	// JSON, plain text, or HTML input that contains the content to be analyzed. For JSON input, provide an object of type `ToneInput`.
	Body string `json:"body,omitempty"`

    // Indicates whether user set optional parameter Body
    IsBodySet bool

	// The type of the input. A character encoding can be specified by including a `charset` parameter. For example, 'text/plain;charset=utf-8'.
	ContentType string `json:"Content-Type"`

	// Indicates whether the service is to return an analysis of each individual sentence in addition to its analysis of the full document. If `true` (the default), the service returns results for each sentence.
	Sentences bool `json:"sentences,omitempty"`

    // Indicates whether user set optional parameter Sentences
    IsSentencesSet bool

	// **`2017-09-21`:** Deprecated. The service continues to accept the parameter for backward-compatibility, but the parameter no longer affects the response. **`2016-05-19`:** A comma-separated list of tones for which the service is to return its analysis of the input; the indicated tones apply both to the full document and to individual sentences of the document. You can specify one or more of the valid values. Omit the parameter to request results for all three tones.
	Tones []string `json:"tones,omitempty"`

    // Indicates whether user set optional parameter Tones
    IsTonesSet bool

	// The language of the input text for the request: English or French. Regional variants are treated as their parent language; for example, `en-US` is interpreted as `en`. The input content must match the specified language. Do not submit content that contains both languages. You can use different languages for **Content-Language** and **Accept-Language**. * **`2017-09-21`:** Accepts `en` or `fr`. * **`2016-05-19`:** Accepts only `en`.
	ContentLanguage string `json:"Content-Language,omitempty"`

    // Indicates whether user set optional parameter ContentLanguage
    IsContentLanguageSet bool

	// The desired language of the response. For two-character arguments, regional variants are treated as their parent language; for example, `en-US` is interpreted as `en`. You can use different languages for **Content-Language** and **Accept-Language**.
	AcceptLanguage string `json:"Accept-Language,omitempty"`

    // Indicates whether user set optional parameter AcceptLanguage
    IsAcceptLanguageSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewToneOptionsForToneInput : Instantiate ToneOptionsForToneInput
func NewToneOptionsForToneInput(toneInput ToneInput) *ToneOptions {
    return &ToneOptions{
        ToneInput: toneInput,
        ContentType: "application/json",
    }
}

// SetToneInput : Allow user to set ToneInput
func (options *ToneOptions) SetToneInput(toneInput ToneInput) *ToneOptions {
    options.ToneInput = toneInput
    options.ContentType = "application/json"
    return options
}

// NewToneOptionsForPlain : Instantiate ToneOptionsForPlain
func NewToneOptionsForPlain(toneInput string) *ToneOptions {
    return &ToneOptions{
        Body: toneInput,
        ContentType: "text/plain",
    }
}

// SetPlain : Allow user to set Plain
func (options *ToneOptions) SetPlain(toneInput string) *ToneOptions {
    options.Body = toneInput
    options.ContentType = "text/plain"
    return options
}

// NewToneOptionsForHTML : Instantiate ToneOptionsForHTML
func NewToneOptionsForHTML(toneInput string) *ToneOptions {
    return &ToneOptions{
        Body: toneInput,
        ContentType: "text/html",
    }
}

// SetHTML : Allow user to set HTML
func (options *ToneOptions) SetHTML(toneInput string) *ToneOptions {
    options.Body = toneInput
    options.ContentType = "text/html"
    return options
}

// SetSentences : Allow user to set Sentences
func (options *ToneOptions) SetSentences(param bool) *ToneOptions {
    options.Sentences = param
    options.IsSentencesSet = true
    return options
}

// SetTones : Allow user to set Tones
func (options *ToneOptions) SetTones(param []string) *ToneOptions {
    options.Tones = param
    options.IsTonesSet = true
    return options
}

// SetContentLanguage : Allow user to set ContentLanguage
func (options *ToneOptions) SetContentLanguage(param string) *ToneOptions {
    options.ContentLanguage = param
    options.IsContentLanguageSet = true
    return options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (options *ToneOptions) SetAcceptLanguage(param string) *ToneOptions {
    options.AcceptLanguage = param
    options.IsAcceptLanguageSet = true
    return options
}

// SetHeaders : Allow user to set Headers
func (options *ToneOptions) SetHeaders(param map[string]string) *ToneOptions {
    options.Headers = param
    return options
}

// ToneScore : ToneScore struct
type ToneScore struct {

	// The score for the tone. * **`2017-09-21`:** The score that is returned lies in the range of 0.5 to 1. A score greater than 0.75 indicates a high likelihood that the tone is perceived in the content. * **`2016-05-19`:** The score that is returned lies in the range of 0 to 1. A score less than 0.5 indicates that the tone is unlikely to be perceived in the content; a score greater than 0.75 indicates a high likelihood that the tone is perceived.
	Score float64 `json:"score"`

	// The unique, non-localized identifier of the tone. * **`2017-09-21`:** The service can return results for the following tone IDs: `anger`, `fear`, `joy`, and `sadness` (emotional tones); `analytical`, `confident`, and `tentative` (language tones). The service returns results only for tones whose scores meet a minimum threshold of 0.5. * **`2016-05-19`:** The service can return results for the following tone IDs of the different categories: for the `emotion` category: `anger`, `disgust`, `fear`, `joy`, and `sadness`; for the `language` category: `analytical`, `confident`, and `tentative`; for the `social` category: `openness_big5`, `conscientiousness_big5`, `extraversion_big5`, `agreeableness_big5`, and `emotional_range_big5`. The service returns scores for all tones of a category, regardless of their values.
	ToneID string `json:"tone_id"`

	// The user-visible, localized name of the tone.
	ToneName string `json:"tone_name"`
}

// Utterance : Utterance struct
type Utterance struct {

	// An utterance contributed by a user in the conversation that is to be analyzed. The utterance can contain multiple sentences.
	Text string `json:"text"`

	// A string that identifies the user who contributed the utterance specified by the `text` parameter.
	User string `json:"user,omitempty"`
}

// UtteranceAnalyses : UtteranceAnalyses struct
type UtteranceAnalyses struct {

	// An array of `UtteranceAnalysis` objects that provides the results for each utterance of the input.
	UtterancesTone []UtteranceAnalysis `json:"utterances_tone"`

	// **`2017-09-21`:** A warning message if the content contains more than 50 utterances. The service analyzes only the first 50 utterances. **`2016-05-19`:** Not returned.
	Warning string `json:"warning,omitempty"`
}

// UtteranceAnalysis : UtteranceAnalysis struct
type UtteranceAnalysis struct {

	// The unique identifier of the utterance. The first utterance has ID 0, and the ID of each subsequent utterance is incremented by one.
	UtteranceID int64 `json:"utterance_id"`

	// The text of the utterance.
	UtteranceText string `json:"utterance_text"`

	// An array of `ToneChatScore` objects that provides results for the most prevalent tones of the utterance. The array includes results for any tone whose score is at least 0.5. The array is empty if no tone has a score that meets this threshold.
	Tones []ToneChatScore `json:"tones"`

	// **`2017-09-21`:** An error message if the utterance contains more than 500 characters. The service does not analyze the utterance. **`2016-05-19`:** Not returned.
	Error string `json:"error,omitempty"`
}
