/**
 * (C) Copyright IBM Corp. 2018, 2020.
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

// Package toneanalyzerv3 : Operations and models for the ToneAnalyzerV3 service
package toneanalyzerv3

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/IBM/go-sdk-core/v4/core"
	common "github.com/watson-developer-cloud/go-sdk/v2/common"
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
// Version: 3.5.3
// See: https://cloud.ibm.com/docs/tone-analyzer
type ToneAnalyzerV3 struct {
	Service *core.BaseService

	// Release date of the version of the API you want to use. Specify dates in YYYY-MM-DD format. The current version is
	// `2017-09-21`.
	Version *string
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.us-south.tone-analyzer.watson.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "tone_analyzer"

// ToneAnalyzerV3Options : Service options
type ToneAnalyzerV3Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator

	// Release date of the version of the API you want to use. Specify dates in YYYY-MM-DD format. The current version is
	// `2017-09-21`.
	Version *string `validate:"required"`
}

// NewToneAnalyzerV3 : constructs an instance of ToneAnalyzerV3 with passed in options.
func NewToneAnalyzerV3(options *ToneAnalyzerV3Options) (service *ToneAnalyzerV3, err error) {
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

	service = &ToneAnalyzerV3{
		Service: baseService,
		Version: options.Version,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "toneAnalyzer" suitable for processing requests.
func (toneAnalyzer *ToneAnalyzerV3) Clone() *ToneAnalyzerV3 {
	if core.IsNil(toneAnalyzer) {
		return nil
	}
	clone := *toneAnalyzer
	clone.Service = toneAnalyzer.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (toneAnalyzer *ToneAnalyzerV3) SetServiceURL(url string) error {
	return toneAnalyzer.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (toneAnalyzer *ToneAnalyzerV3) GetServiceURL() string {
	return toneAnalyzer.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (toneAnalyzer *ToneAnalyzerV3) SetDefaultHeaders(headers http.Header) {
	toneAnalyzer.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (toneAnalyzer *ToneAnalyzerV3) SetEnableGzipCompression(enableGzip bool) {
	toneAnalyzer.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (toneAnalyzer *ToneAnalyzerV3) GetEnableGzipCompression() bool {
	return toneAnalyzer.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (toneAnalyzer *ToneAnalyzerV3) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	toneAnalyzer.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (toneAnalyzer *ToneAnalyzerV3) DisableRetries() {
	toneAnalyzer.Service.DisableRetries()
}

// DisableSSLVerification bypasses verification of the server's SSL certificate
func (toneAnalyzer *ToneAnalyzerV3) DisableSSLVerification() {
	toneAnalyzer.Service.DisableSSLVerification()
}

// Tone : Analyze general tone
// Use the general-purpose endpoint to analyze the tone of your input content. The service analyzes the content for
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
// endpoint](https://cloud.ibm.com/docs/tone-analyzer?topic=tone-analyzer-utgpe#utgpe).
func (toneAnalyzer *ToneAnalyzerV3) Tone(toneOptions *ToneOptions) (result *ToneAnalysis, response *core.DetailedResponse, err error) {
	return toneAnalyzer.ToneWithContext(context.Background(), toneOptions)
}

// ToneWithContext is an alternate form of the Tone method which supports a Context parameter
func (toneAnalyzer *ToneAnalyzerV3) ToneWithContext(ctx context.Context, toneOptions *ToneOptions) (result *ToneAnalysis, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(toneOptions, "toneOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(toneOptions, "toneOptions")
	if err != nil {
		return
	}

	if toneOptions.ToneInput != nil && toneOptions.ContentType == nil {
		toneOptions.SetContentType("application/json")
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = toneAnalyzer.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(toneAnalyzer.Service.Options.URL, `/v3/tone`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range toneOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("tone_analyzer", "V3", "Tone")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if toneOptions.ContentType != nil {
		builder.AddHeader("Content-Type", fmt.Sprint(*toneOptions.ContentType))
	}
	if toneOptions.ContentLanguage != nil {
		builder.AddHeader("Content-Language", fmt.Sprint(*toneOptions.ContentLanguage))
	}
	if toneOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*toneOptions.AcceptLanguage))
	}

	builder.AddQuery("version", fmt.Sprint(*toneAnalyzer.Version))
	if toneOptions.Sentences != nil {
		builder.AddQuery("sentences", fmt.Sprint(*toneOptions.Sentences))
	}
	if toneOptions.Tones != nil {
		builder.AddQuery("tones", strings.Join(toneOptions.Tones, ","))
	}

	_, err = builder.SetBodyContent(core.StringNilMapper(toneOptions.ContentType), toneOptions.ToneInput, nil, toneOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = toneAnalyzer.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalToneAnalysis)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ToneChat : Analyze customer-engagement tone
// Use the customer-engagement endpoint to analyze the tone of customer service and customer support conversations. For
// each utterance of a conversation, the method reports the most prevalent subset of the following seven tones: sad,
// frustrated, satisfied, excited, polite, impolite, and sympathetic.
//
// If you submit more than 50 utterances, the service returns a warning for the overall content and analyzes only the
// first 50 utterances. If you submit a single utterance that contains more than 500 characters, the service returns an
// error for that utterance and does not analyze the utterance. The request fails if all utterances have more than 500
// characters. Per the JSON specification, the default character encoding for JSON content is effectively always UTF-8.
//
// **See also:** [Using the customer-engagement
// endpoint](https://cloud.ibm.com/docs/tone-analyzer?topic=tone-analyzer-utco#utco).
func (toneAnalyzer *ToneAnalyzerV3) ToneChat(toneChatOptions *ToneChatOptions) (result *UtteranceAnalyses, response *core.DetailedResponse, err error) {
	return toneAnalyzer.ToneChatWithContext(context.Background(), toneChatOptions)
}

// ToneChatWithContext is an alternate form of the ToneChat method which supports a Context parameter
func (toneAnalyzer *ToneAnalyzerV3) ToneChatWithContext(ctx context.Context, toneChatOptions *ToneChatOptions) (result *UtteranceAnalyses, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(toneChatOptions, "toneChatOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(toneChatOptions, "toneChatOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = toneAnalyzer.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(toneAnalyzer.Service.Options.URL, `/v3/tone_chat`, nil)
	if err != nil {
		return
	}

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

	builder.AddQuery("version", fmt.Sprint(*toneAnalyzer.Version))

	body := make(map[string]interface{})
	if toneChatOptions.Utterances != nil {
		body["utterances"] = toneChatOptions.Utterances
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
	response, err = toneAnalyzer.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalUtteranceAnalyses)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DocumentAnalysis : The results of the analysis for the full input content.
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

// UnmarshalDocumentAnalysis unmarshals an instance of DocumentAnalysis from the specified map of raw messages.
func UnmarshalDocumentAnalysis(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DocumentAnalysis)
	err = core.UnmarshalModel(m, "tones", &obj.Tones, UnmarshalToneScore)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "tone_categories", &obj.ToneCategories, UnmarshalToneCategory)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "warning", &obj.Warning)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SentenceAnalysis : The results of the analysis for the individual sentences of the input content.
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

// UnmarshalSentenceAnalysis unmarshals an instance of SentenceAnalysis from the specified map of raw messages.
func UnmarshalSentenceAnalysis(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SentenceAnalysis)
	err = core.UnmarshalPrimitive(m, "sentence_id", &obj.SentenceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "tones", &obj.Tones, UnmarshalToneScore)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "tone_categories", &obj.ToneCategories, UnmarshalToneCategory)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "input_from", &obj.InputFrom)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "input_to", &obj.InputTo)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ToneAnalysis : The tone analysis results for the input from the general-purpose endpoint.
type ToneAnalysis struct {
	// The results of the analysis for the full input content.
	DocumentTone *DocumentAnalysis `json:"document_tone" validate:"required"`

	// An array of `SentenceAnalysis` objects that provides the results of the analysis for the individual sentences of the
	// input content. The service returns results only for the first 100 sentences of the input. The field is omitted if
	// the `sentences` parameter of the request is set to `false`.
	SentencesTone []SentenceAnalysis `json:"sentences_tone,omitempty"`
}

// UnmarshalToneAnalysis unmarshals an instance of ToneAnalysis from the specified map of raw messages.
func UnmarshalToneAnalysis(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ToneAnalysis)
	err = core.UnmarshalModel(m, "document_tone", &obj.DocumentTone, UnmarshalDocumentAnalysis)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "sentences_tone", &obj.SentencesTone, UnmarshalSentenceAnalysis)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ToneCategory : The category for a tone from the input content.
type ToneCategory struct {
	// An array of `ToneScore` objects that provides the results for the tones of the category.
	Tones []ToneScore `json:"tones" validate:"required"`

	// The unique, non-localized identifier of the category for the results. The service can return results for the
	// following category IDs: `emotion_tone`, `language_tone`, and `social_tone`.
	CategoryID *string `json:"category_id" validate:"required"`

	// The user-visible, localized name of the category.
	CategoryName *string `json:"category_name" validate:"required"`
}

// UnmarshalToneCategory unmarshals an instance of ToneCategory from the specified map of raw messages.
func UnmarshalToneCategory(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ToneCategory)
	err = core.UnmarshalModel(m, "tones", &obj.Tones, UnmarshalToneScore)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "category_id", &obj.CategoryID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "category_name", &obj.CategoryName)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ToneChatOptions : The ToneChat options.
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

	// Allows users to set headers on API requests
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
	ToneChatOptionsContentLanguageEnConst = "en"
	ToneChatOptionsContentLanguageFrConst = "fr"
)

// Constants associated with the ToneChatOptions.AcceptLanguage property.
// The desired language of the response. For two-character arguments, regional variants are treated as their parent
// language; for example, `en-US` is interpreted as `en`. You can use different languages for **Content-Language** and
// **Accept-Language**.
const (
	ToneChatOptionsAcceptLanguageArConst   = "ar"
	ToneChatOptionsAcceptLanguageDeConst   = "de"
	ToneChatOptionsAcceptLanguageEnConst   = "en"
	ToneChatOptionsAcceptLanguageEsConst   = "es"
	ToneChatOptionsAcceptLanguageFrConst   = "fr"
	ToneChatOptionsAcceptLanguageItConst   = "it"
	ToneChatOptionsAcceptLanguageJaConst   = "ja"
	ToneChatOptionsAcceptLanguageKoConst   = "ko"
	ToneChatOptionsAcceptLanguagePtBrConst = "pt-br"
	ToneChatOptionsAcceptLanguageZhCnConst = "zh-cn"
	ToneChatOptionsAcceptLanguageZhTwConst = "zh-tw"
)

// NewToneChatOptions : Instantiate ToneChatOptions
func (*ToneAnalyzerV3) NewToneChatOptions(utterances []Utterance) *ToneChatOptions {
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

// ToneChatScore : The score for an utterance from the input content.
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
	ToneChatScoreToneIDExcitedConst     = "excited"
	ToneChatScoreToneIDFrustratedConst  = "frustrated"
	ToneChatScoreToneIDImpoliteConst    = "impolite"
	ToneChatScoreToneIDPoliteConst      = "polite"
	ToneChatScoreToneIDSadConst         = "sad"
	ToneChatScoreToneIDSatisfiedConst   = "satisfied"
	ToneChatScoreToneIDSympatheticConst = "sympathetic"
)

// UnmarshalToneChatScore unmarshals an instance of ToneChatScore from the specified map of raw messages.
func UnmarshalToneChatScore(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ToneChatScore)
	err = core.UnmarshalPrimitive(m, "score", &obj.Score)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tone_id", &obj.ToneID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tone_name", &obj.ToneName)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ToneInput : Input for the general-purpose endpoint.
type ToneInput struct {
	// The input content that the service is to analyze.
	Text *string `json:"text" validate:"required"`
}

// NewToneInput : Instantiate ToneInput (Generic Model Constructor)
func (*ToneAnalyzerV3) NewToneInput(text string) (model *ToneInput, err error) {
	model = &ToneInput{
		Text: core.StringPtr(text),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalToneInput unmarshals an instance of ToneInput from the specified map of raw messages.
func UnmarshalToneInput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ToneInput)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ToneOptions : The Tone options.
type ToneOptions struct {
	// JSON, plain text, or HTML input that contains the content to be analyzed. For JSON input, provide an object of type
	// `ToneInput`.
	ToneInput *ToneInput `json:"tone_input,omitempty"`

	// JSON, plain text, or HTML input that contains the content to be analyzed. For JSON input, provide an object of type
	// `ToneInput`.
	Body *string `json:"body,omitempty"`

	// The type of the input. A character encoding can be specified by including a `charset` parameter. For example,
	// 'text/plain;charset=utf-8'.
	ContentType *string `json:"Content-Type,omitempty"`

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

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ToneOptions.Tone property.
const (
	ToneOptionsToneEmotionConst  = "emotion"
	ToneOptionsToneLanguageConst = "language"
	ToneOptionsToneSocialConst   = "social"
)

// Constants associated with the ToneOptions.ContentLanguage property.
// The language of the input text for the request: English or French. Regional variants are treated as their parent
// language; for example, `en-US` is interpreted as `en`. The input content must match the specified language. Do not
// submit content that contains both languages. You can use different languages for **Content-Language** and
// **Accept-Language**.
// * **`2017-09-21`:** Accepts `en` or `fr`.
// * **`2016-05-19`:** Accepts only `en`.
const (
	ToneOptionsContentLanguageEnConst = "en"
	ToneOptionsContentLanguageFrConst = "fr"
)

// Constants associated with the ToneOptions.AcceptLanguage property.
// The desired language of the response. For two-character arguments, regional variants are treated as their parent
// language; for example, `en-US` is interpreted as `en`. You can use different languages for **Content-Language** and
// **Accept-Language**.
const (
	ToneOptionsAcceptLanguageArConst   = "ar"
	ToneOptionsAcceptLanguageDeConst   = "de"
	ToneOptionsAcceptLanguageEnConst   = "en"
	ToneOptionsAcceptLanguageEsConst   = "es"
	ToneOptionsAcceptLanguageFrConst   = "fr"
	ToneOptionsAcceptLanguageItConst   = "it"
	ToneOptionsAcceptLanguageJaConst   = "ja"
	ToneOptionsAcceptLanguageKoConst   = "ko"
	ToneOptionsAcceptLanguagePtBrConst = "pt-br"
	ToneOptionsAcceptLanguageZhCnConst = "zh-cn"
	ToneOptionsAcceptLanguageZhTwConst = "zh-tw"
)

// NewToneOptions : Instantiate ToneOptions
func (*ToneAnalyzerV3) NewToneOptions() *ToneOptions {
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

// SetContentType : Allow user to set ContentType
func (options *ToneOptions) SetContentType(contentType string) *ToneOptions {
	options.ContentType = core.StringPtr(contentType)
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

// SetHeaders : Allow user to set Headers
func (options *ToneOptions) SetHeaders(param map[string]string) *ToneOptions {
	options.Headers = param
	return options
}

// ToneScore : The score for a tone from the input content.
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

// UnmarshalToneScore unmarshals an instance of ToneScore from the specified map of raw messages.
func UnmarshalToneScore(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ToneScore)
	err = core.UnmarshalPrimitive(m, "score", &obj.Score)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tone_id", &obj.ToneID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tone_name", &obj.ToneName)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Utterance : An utterance for the input of the general-purpose endpoint.
type Utterance struct {
	// An utterance contributed by a user in the conversation that is to be analyzed. The utterance can contain multiple
	// sentences.
	Text *string `json:"text" validate:"required"`

	// A string that identifies the user who contributed the utterance specified by the `text` parameter.
	User *string `json:"user,omitempty"`
}

// NewUtterance : Instantiate Utterance (Generic Model Constructor)
func (*ToneAnalyzerV3) NewUtterance(text string) (model *Utterance, err error) {
	model = &Utterance{
		Text: core.StringPtr(text),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalUtterance unmarshals an instance of Utterance from the specified map of raw messages.
func UnmarshalUtterance(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Utterance)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "user", &obj.User)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UtteranceAnalyses : The results of the analysis for the utterances of the input content.
type UtteranceAnalyses struct {
	// An array of `UtteranceAnalysis` objects that provides the results for each utterance of the input.
	UtterancesTone []UtteranceAnalysis `json:"utterances_tone" validate:"required"`

	// **`2017-09-21`:** A warning message if the content contains more than 50 utterances. The service analyzes only the
	// first 50 utterances. **`2016-05-19`:** Not returned.
	Warning *string `json:"warning,omitempty"`
}

// UnmarshalUtteranceAnalyses unmarshals an instance of UtteranceAnalyses from the specified map of raw messages.
func UnmarshalUtteranceAnalyses(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UtteranceAnalyses)
	err = core.UnmarshalModel(m, "utterances_tone", &obj.UtterancesTone, UnmarshalUtteranceAnalysis)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "warning", &obj.Warning)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UtteranceAnalysis : The results of the analysis for an utterance of the input content.
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

// UnmarshalUtteranceAnalysis unmarshals an instance of UtteranceAnalysis from the specified map of raw messages.
func UnmarshalUtteranceAnalysis(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UtteranceAnalysis)
	err = core.UnmarshalPrimitive(m, "utterance_id", &obj.UtteranceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "utterance_text", &obj.UtteranceText)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "tones", &obj.Tones, UnmarshalToneChatScore)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "error", &obj.Error)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
