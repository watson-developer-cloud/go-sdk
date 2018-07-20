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
package toneAnalyzerV3

import (
    "bytes"
    "fmt"
    req "github.com/parnurzeal/gorequest"
    watson "golang-sdk"
)

type ToneAnalyzerV3 struct {
	client *watson.Client
}

func NewToneAnalyzerV3(creds watson.Credentials) (*ToneAnalyzerV3, error) {
	client, clientErr := watson.NewClient(creds, "tone_analyzer")

	if clientErr != nil {
		return nil, clientErr
	}

	return &ToneAnalyzerV3{ client: client }, nil
}

// Analyze general tone
func (toneAnalyzer *ToneAnalyzerV3) Tone(body *ToneInput, contentType string, sentences bool, tones []string, contentLanguage string, acceptLanguage string) (*watson.WatsonResponse, []error) {
    path := "/v3/tone"
    creds := toneAnalyzer.client.Creds
    useTM := toneAnalyzer.client.UseTM
    tokenManager := toneAnalyzer.client.TokenManager

    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", fmt.Sprint(contentType))
    request.Set("Content-Language", fmt.Sprint(contentLanguage))
    request.Set("Accept-Language", fmt.Sprint(acceptLanguage))
    request.Query("version=" + creds.Version)
    request.Query("sentences=" + fmt.Sprint(sentences))
    request.Query("tones=" + fmt.Sprint(tones))
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

    response.Result = new(ToneAnalysis)
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

func GetToneResult(response *watson.WatsonResponse) *ToneAnalysis {
    result, ok := response.Result.(*ToneAnalysis)

    if ok {
        return result
    }

    return nil
}

// Analyze customer engagement tone
func (toneAnalyzer *ToneAnalyzerV3) ToneChat(body *ToneChatInput, contentLanguage string, acceptLanguage string) (*watson.WatsonResponse, []error) {
    path := "/v3/tone_chat"
    creds := toneAnalyzer.client.Creds
    useTM := toneAnalyzer.client.UseTM
    tokenManager := toneAnalyzer.client.TokenManager

    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Language", fmt.Sprint(contentLanguage))
    request.Set("Accept-Language", fmt.Sprint(acceptLanguage))
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

    response.Result = new(UtteranceAnalyses)
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

func GetToneChatResult(response *watson.WatsonResponse) *UtteranceAnalyses {
    result, ok := response.Result.(*UtteranceAnalyses)

    if ok {
        return result
    }

    return nil
}


type DocumentAnalysis struct {

	// **`2017-09-21`:** An array of `ToneScore` objects that provides the results of the analysis for each qualifying tone of the document. The array includes results for any tone whose score is at least 0.5. The array is empty if no tone has a score that meets this threshold. **`2016-05-19`:** Not returned.
	Tones []ToneScore `json:"tones,omitempty"`

	// **`2017-09-21`:** Not returned. **`2016-05-19`:** An array of `ToneCategory` objects that provides the results of the tone analysis for the full document of the input content. The service returns results only for the tones specified with the `tones` parameter of the request.
	ToneCategories []ToneCategory `json:"tone_categories,omitempty"`

	// **`2017-09-21`:** A warning message if the overall content exceeds 128 KB or contains more than 1000 sentences. The service analyzes only the first 1000 sentences for document-level analysis and the first 100 sentences for sentence-level analysis. **`2016-05-19`:** Not returned.
	Warning string `json:"warning,omitempty"`
}

type SentenceAnalysis struct {

	// The unique identifier of a sentence of the input content. The first sentence has ID 0, and the ID of each subsequent sentence is incremented by one.
	SentenceId int64 `json:"sentence_id"`

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

type ToneAnalysis struct {

	// An object of type `DocumentAnalysis` that provides the results of the analysis for the full input document.
	DocumentTone DocumentAnalysis `json:"document_tone"`

	// An array of `SentenceAnalysis` objects that provides the results of the analysis for the individual sentences of the input content. The service returns results only for the first 100 sentences of the input. The field is omitted if the `sentences` parameter of the request is set to `false`.
	SentencesTone []SentenceAnalysis `json:"sentences_tone,omitempty"`
}

type ToneCategory struct {

	// An array of `ToneScore` objects that provides the results for the tones of the category.
	Tones []ToneScore `json:"tones"`

	// The unique, non-localized identifier of the category for the results. The service can return results for the following category IDs: `emotion_tone`, `language_tone`, and `social_tone`.
	CategoryId string `json:"category_id"`

	// The user-visible, localized name of the category.
	CategoryName string `json:"category_name"`
}

type ToneChatInput struct {

	// An array of `Utterance` objects that provides the input content that the service is to analyze.
	Utterances []Utterance `json:"utterances"`
}

type ToneChatScore struct {

	// The score for the tone in the range of 0.5 to 1. A score greater than 0.75 indicates a high likelihood that the tone is perceived in the utterance.
	Score float64 `json:"score"`

	// The unique, non-localized identifier of the tone for the results. The service can return results for the following tone IDs: `sad`, `frustrated`, `satisfied`, `excited`, `polite`, `impolite`, and `sympathetic`. The service returns results only for tones whose scores meet a minimum threshold of 0.5.
	ToneId string `json:"tone_id"`

	// The user-visible, localized name of the tone.
	ToneName string `json:"tone_name"`
}

type ToneInput struct {

	// The input content that the service is to analyze.
	Text string `json:"text"`
}

type ToneScore struct {

	// The score for the tone. * **`2017-09-21`:** The score that is returned lies in the range of 0.5 to 1. A score greater than 0.75 indicates a high likelihood that the tone is perceived in the content. * **`2016-05-19`:** The score that is returned lies in the range of 0 to 1. A score less than 0.5 indicates that the tone is unlikely to be perceived in the content; a score greater than 0.75 indicates a high likelihood that the tone is perceived.
	Score float64 `json:"score"`

	// The unique, non-localized identifier of the tone. * **`2017-09-21`:** The service can return results for the following tone IDs: `anger`, `fear`, `joy`, and `sadness` (emotional tones); `analytical`, `confident`, and `tentative` (language tones). The service returns results only for tones whose scores meet a minimum threshold of 0.5. * **`2016-05-19`:** The service can return results for the following tone IDs of the different categories: for the `emotion` category: `anger`, `disgust`, `fear`, `joy`, and `sadness`; for the `language` category: `analytical`, `confident`, and `tentative`; for the `social` category: `openness_big5`, `conscientiousness_big5`, `extraversion_big5`, `agreeableness_big5`, and `emotional_range_big5`. The service returns scores for all tones of a category, regardless of their values.
	ToneId string `json:"tone_id"`

	// The user-visible, localized name of the tone.
	ToneName string `json:"tone_name"`
}

type Utterance struct {

	// An utterance contributed by a user in the conversation that is to be analyzed. The utterance can contain multiple sentences.
	Text string `json:"text"`

	// A string that identifies the user who contributed the utterance specified by the `text` parameter.
	User string `json:"user,omitempty"`
}

type UtteranceAnalyses struct {

	// An array of `UtteranceAnalysis` objects that provides the results for each utterance of the input.
	UtterancesTone []UtteranceAnalysis `json:"utterances_tone"`

	// **`2017-09-21`:** A warning message if the content contains more than 50 utterances. The service analyzes only the first 50 utterances. **`2016-05-19`:** Not returned.
	Warning string `json:"warning,omitempty"`
}

type UtteranceAnalysis struct {

	// The unique identifier of the utterance. The first utterance has ID 0, and the ID of each subsequent utterance is incremented by one.
	UtteranceId int64 `json:"utterance_id"`

	// The text of the utterance.
	UtteranceText string `json:"utterance_text"`

	// An array of `ToneChatScore` objects that provides results for the most prevalent tones of the utterance. The array includes results for any tone whose score is at least 0.5. The array is empty if no tone has a score that meets this threshold.
	Tones []ToneChatScore `json:"tones"`

	// **`2017-09-21`:** An error message if the utterance contains more than 500 characters. The service does not analyze the utterance. **`2016-05-19`:** Not returned.
	Error string `json:"error,omitempty"`
}
