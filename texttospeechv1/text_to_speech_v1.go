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
 * IBM OpenAPI SDK Code Generator Version: 3.22.0-937b9a1c-20201211-223043
 */

// Package texttospeechv1 : Operations and models for the TextToSpeechV1 service
package texttospeechv1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	common "github.com/watson-developer-cloud/go-sdk/common"
	"io"
	"net/http"
	"reflect"
	"time"
)

// TextToSpeechV1 : The IBM Watson&trade; Text to Speech service provides APIs that use IBM's speech-synthesis
// capabilities to synthesize text into natural-sounding speech in a variety of languages, dialects, and voices. The
// service supports at least one male or female voice, sometimes both, for each language. The audio is streamed back to
// the client with minimal delay.
//
// For speech synthesis, the service supports a synchronous HTTP Representational State Transfer (REST) interface and a
// WebSocket interface. Both interfaces support plain text and SSML input. SSML is an XML-based markup language that
// provides text annotation for speech-synthesis applications. The WebSocket interface also supports the SSML
// <code>&lt;mark&gt;</code> element and word timings.
//
// The service offers a customization interface that you can use to define sounds-like or phonetic translations for
// words. A sounds-like translation consists of one or more words that, when combined, sound like the word. A phonetic
// translation is based on the SSML phoneme format for representing a word. You can specify a phonetic translation in
// standard International Phonetic Alphabet (IPA) representation or in the proprietary IBM Symbolic Phonetic
// Representation (SPR). The Arabic, Chinese, Dutch, Australian English, and Korean languages support only IPA.
//
// Version: 1.0.0
// See: https://cloud.ibm.com/docs/text-to-speech
type TextToSpeechV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.us-south.text-to-speech.watson.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "text_to_speech"

// TextToSpeechV1Options : Service options
type TextToSpeechV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewTextToSpeechV1 : constructs an instance of TextToSpeechV1 with passed in options.
func NewTextToSpeechV1(options *TextToSpeechV1Options) (service *TextToSpeechV1, err error) {
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

	service = &TextToSpeechV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "textToSpeech" suitable for processing requests.
func (textToSpeech *TextToSpeechV1) Clone() *TextToSpeechV1 {
	if core.IsNil(textToSpeech) {
		return nil
	}
	clone := *textToSpeech
	clone.Service = textToSpeech.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (textToSpeech *TextToSpeechV1) SetServiceURL(url string) error {
	return textToSpeech.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (textToSpeech *TextToSpeechV1) GetServiceURL() string {
	return textToSpeech.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (textToSpeech *TextToSpeechV1) SetDefaultHeaders(headers http.Header) {
	textToSpeech.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (textToSpeech *TextToSpeechV1) SetEnableGzipCompression(enableGzip bool) {
	textToSpeech.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (textToSpeech *TextToSpeechV1) GetEnableGzipCompression() bool {
	return textToSpeech.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (textToSpeech *TextToSpeechV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	textToSpeech.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (textToSpeech *TextToSpeechV1) DisableRetries() {
	textToSpeech.Service.DisableRetries()
}

// DisableSSLVerification bypasses verification of the server's SSL certificate
func (textToSpeech *TextToSpeechV1) DisableSSLVerification() {
	textToSpeech.Service.DisableSSLVerification()
}

// ListVoices : List voices
// Lists all voices available for use with the service. The information includes the name, language, gender, and other
// details about the voice. The ordering of the list of voices can change from call to call; do not rely on an
// alphabetized or static list of voices. To see information about a specific voice, use the **Get a voice** method.
//
// **See also:** [Listing all available
// voices](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-voices#listVoices).
func (textToSpeech *TextToSpeechV1) ListVoices(listVoicesOptions *ListVoicesOptions) (result *Voices, response *core.DetailedResponse, err error) {
	return textToSpeech.ListVoicesWithContext(context.Background(), listVoicesOptions)
}

// ListVoicesWithContext is an alternate form of the ListVoices method which supports a Context parameter
func (textToSpeech *TextToSpeechV1) ListVoicesWithContext(ctx context.Context, listVoicesOptions *ListVoicesOptions) (result *Voices, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listVoicesOptions, "listVoicesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = textToSpeech.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(textToSpeech.Service.Options.URL, `/v1/voices`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listVoicesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "ListVoices")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = textToSpeech.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalVoices)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetVoice : Get a voice
// Gets information about the specified voice. The information includes the name, language, gender, and other details
// about the voice. Specify a customization ID to obtain information for a custom model that is defined for the language
// of the specified voice. To list information about all available voices, use the **List voices** method.
//
// **See also:** [Listing a specific
// voice](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-voices#listVoice).
//
// ### Important voice updates
//
//  The service's voices underwent significant change on 2 December 2020.
// * The Arabic, Chinese, Dutch, Australian English, and Korean voices are now neural instead of concatenative.
// * The `ar-AR_OmarVoice` voice is deprecated. Use `ar-MS_OmarVoice` voice instead.
// * The `ar-AR` language identifier cannot be used to create a custom model. Use the `ar-MS` identifier instead.
// * The standard concatenative voices for the following languages are now deprecated: Brazilian Portuguese, United
// Kingdom and United States English, French, German, Italian, Japanese, and Spanish (all dialects).
// * The features expressive SSML, voice transformation SSML, and use of the `volume` attribute of the `<prosody>`
// element are deprecated and are not supported with any of the service's neural voices.
// * All of the service's voices are now customizable and generally available (GA) for production use.
//
// The deprecated voices and features will continue to function for at least one year but might be removed at a future
// date. You are encouraged to migrate to the equivalent neural voices at your earliest convenience. For more
// information about all voice updates, see the [2 December 2020 service
// update](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-release-notes#December2020) in the release
// notes.
func (textToSpeech *TextToSpeechV1) GetVoice(getVoiceOptions *GetVoiceOptions) (result *Voice, response *core.DetailedResponse, err error) {
	return textToSpeech.GetVoiceWithContext(context.Background(), getVoiceOptions)
}

// GetVoiceWithContext is an alternate form of the GetVoice method which supports a Context parameter
func (textToSpeech *TextToSpeechV1) GetVoiceWithContext(ctx context.Context, getVoiceOptions *GetVoiceOptions) (result *Voice, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getVoiceOptions, "getVoiceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getVoiceOptions, "getVoiceOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"voice": *getVoiceOptions.Voice,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = textToSpeech.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(textToSpeech.Service.Options.URL, `/v1/voices/{voice}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getVoiceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "GetVoice")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getVoiceOptions.CustomizationID != nil {
		builder.AddQuery("customization_id", fmt.Sprint(*getVoiceOptions.CustomizationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = textToSpeech.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalVoice)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// Synthesize : Synthesize audio
// Synthesizes text to audio that is spoken in the specified voice. The service bases its understanding of the language
// for the input text on the specified voice. Use a voice that matches the language of the input text.
//
// The method accepts a maximum of 5 KB of input text in the body of the request, and 8 KB for the URL and headers. The
// 5 KB limit includes any SSML tags that you specify. The service returns the synthesized audio stream as an array of
// bytes.
//
// **See also:** [The HTTP
// interface](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-usingHTTP#usingHTTP).
//
// ### Audio formats (accept types)
//
//  The service can return audio in the following formats (MIME types).
// * Where indicated, you can optionally specify the sampling rate (`rate`) of the audio. You must specify a sampling
// rate for the `audio/l16` and `audio/mulaw` formats. A specified sampling rate must lie in the range of 8 kHz to 192
// kHz. Some formats restrict the sampling rate to certain values, as noted.
// * For the `audio/l16` format, you can optionally specify the endianness (`endianness`) of the audio:
// `endianness=big-endian` or `endianness=little-endian`.
//
// Use the `Accept` header or the `accept` parameter to specify the requested format of the response audio. If you omit
// an audio format altogether, the service returns the audio in Ogg format with the Opus codec
// (`audio/ogg;codecs=opus`). The service always returns single-channel audio.
// * `audio/basic` - The service returns audio with a sampling rate of 8000 Hz.
// * `audio/flac` - You can optionally specify the `rate` of the audio. The default sampling rate is 22,050 Hz.
// * `audio/l16` - You must specify the `rate` of the audio. You can optionally specify the `endianness` of the audio.
// The default endianness is `little-endian`.
// * `audio/mp3` - You can optionally specify the `rate` of the audio. The default sampling rate is 22,050 Hz.
// * `audio/mpeg` - You can optionally specify the `rate` of the audio. The default sampling rate is 22,050 Hz.
// * `audio/mulaw` - You must specify the `rate` of the audio.
// * `audio/ogg` - The service returns the audio in the `vorbis` codec. You can optionally specify the `rate` of the
// audio. The default sampling rate is 22,050 Hz.
// * `audio/ogg;codecs=opus` - You can optionally specify the `rate` of the audio. Only the following values are valid
// sampling rates: `48000`, `24000`, `16000`, `12000`, or `8000`. If you specify a value other than one of these, the
// service returns an error. The default sampling rate is 48,000 Hz.
// * `audio/ogg;codecs=vorbis` - You can optionally specify the `rate` of the audio. The default sampling rate is 22,050
// Hz.
// * `audio/wav` - You can optionally specify the `rate` of the audio. The default sampling rate is 22,050 Hz.
// * `audio/webm` - The service returns the audio in the `opus` codec. The service returns audio with a sampling rate of
// 48,000 Hz.
// * `audio/webm;codecs=opus` - The service returns audio with a sampling rate of 48,000 Hz.
// * `audio/webm;codecs=vorbis` - You can optionally specify the `rate` of the audio. The default sampling rate is
// 22,050 Hz.
//
// For more information about specifying an audio format, including additional details about some of the formats, see
// [Audio formats](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-audioFormats#audioFormats).
//
// ### Important voice updates
//
//  The service's voices underwent significant change on 2 December 2020.
// * The Arabic, Chinese, Dutch, Australian English, and Korean voices are now neural instead of concatenative.
// * The `ar-AR_OmarVoice` voice is deprecated. Use `ar-MS_OmarVoice` voice instead.
// * The `ar-AR` language identifier cannot be used to create a custom model. Use the `ar-MS` identifier instead.
// * The standard concatenative voices for the following languages are now deprecated: Brazilian Portuguese, United
// Kingdom and United States English, French, German, Italian, Japanese, and Spanish (all dialects).
// * The features expressive SSML, voice transformation SSML, and use of the `volume` attribute of the `<prosody>`
// element are deprecated and are not supported with any of the service's neural voices.
// * All of the service's voices are now customizable and generally available (GA) for production use.
//
// The deprecated voices and features will continue to function for at least one year but might be removed at a future
// date. You are encouraged to migrate to the equivalent neural voices at your earliest convenience. For more
// information about all voice updates, see the [2 December 2020 service
// update](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-release-notes#December2020) in the release
// notes.
//
// ### Warning messages
//
//  If a request includes invalid query parameters, the service returns a `Warnings` response header that provides
// messages about the invalid parameters. The warning includes a descriptive message and a list of invalid argument
// strings. For example, a message such as `"Unknown arguments:"` or `"Unknown url query arguments:"` followed by a list
// of the form `"{invalid_arg_1}, {invalid_arg_2}."` The request succeeds despite the warnings.
func (textToSpeech *TextToSpeechV1) Synthesize(synthesizeOptions *SynthesizeOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	return textToSpeech.SynthesizeWithContext(context.Background(), synthesizeOptions)
}

// SynthesizeWithContext is an alternate form of the Synthesize method which supports a Context parameter
func (textToSpeech *TextToSpeechV1) SynthesizeWithContext(ctx context.Context, synthesizeOptions *SynthesizeOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(synthesizeOptions, "synthesizeOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(synthesizeOptions, "synthesizeOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = textToSpeech.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(textToSpeech.Service.Options.URL, `/v1/synthesize`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range synthesizeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "Synthesize")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "audio/basic")
	builder.AddHeader("Content-Type", "application/json")
	if synthesizeOptions.Accept != nil {
		builder.AddHeader("Accept", fmt.Sprint(*synthesizeOptions.Accept))
	}

	if synthesizeOptions.Voice != nil {
		builder.AddQuery("voice", fmt.Sprint(*synthesizeOptions.Voice))
	}
	if synthesizeOptions.CustomizationID != nil {
		builder.AddQuery("customization_id", fmt.Sprint(*synthesizeOptions.CustomizationID))
	}

	body := make(map[string]interface{})
	if synthesizeOptions.Text != nil {
		body["text"] = synthesizeOptions.Text
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = textToSpeech.Service.Request(request, &result)

	return
}

// GetPronunciation : Get pronunciation
// Gets the phonetic pronunciation for the specified word. You can request the pronunciation for a specific format. You
// can also request the pronunciation for a specific voice to see the default translation for the language of that voice
// or for a specific custom model to see the translation for that model.
//
// **See also:** [Querying a word from a
// language](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customWords#cuWordsQueryLanguage).
//
// ### Important voice updates
//
//  The service's voices underwent significant change on 2 December 2020.
// * The Arabic, Chinese, Dutch, Australian English, and Korean voices are now neural instead of concatenative.
// * The `ar-AR_OmarVoice` voice is deprecated. Use `ar-MS_OmarVoice` voice instead.
// * The `ar-AR` language identifier cannot be used to create a custom model. Use the `ar-MS` identifier instead.
// * The standard concatenative voices for the following languages are now deprecated: Brazilian Portuguese, United
// Kingdom and United States English, French, German, Italian, Japanese, and Spanish (all dialects).
// * The features expressive SSML, voice transformation SSML, and use of the `volume` attribute of the `<prosody>`
// element are deprecated and are not supported with any of the service's neural voices.
// * All of the service's voices are now customizable and generally available (GA) for production use.
//
// The deprecated voices and features will continue to function for at least one year but might be removed at a future
// date. You are encouraged to migrate to the equivalent neural voices at your earliest convenience. For more
// information about all voice updates, see the [2 December 2020 service
// update](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-release-notes#December2020) in the release
// notes.
func (textToSpeech *TextToSpeechV1) GetPronunciation(getPronunciationOptions *GetPronunciationOptions) (result *Pronunciation, response *core.DetailedResponse, err error) {
	return textToSpeech.GetPronunciationWithContext(context.Background(), getPronunciationOptions)
}

// GetPronunciationWithContext is an alternate form of the GetPronunciation method which supports a Context parameter
func (textToSpeech *TextToSpeechV1) GetPronunciationWithContext(ctx context.Context, getPronunciationOptions *GetPronunciationOptions) (result *Pronunciation, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getPronunciationOptions, "getPronunciationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getPronunciationOptions, "getPronunciationOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = textToSpeech.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(textToSpeech.Service.Options.URL, `/v1/pronunciation`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getPronunciationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "GetPronunciation")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("text", fmt.Sprint(*getPronunciationOptions.Text))
	if getPronunciationOptions.Voice != nil {
		builder.AddQuery("voice", fmt.Sprint(*getPronunciationOptions.Voice))
	}
	if getPronunciationOptions.Format != nil {
		builder.AddQuery("format", fmt.Sprint(*getPronunciationOptions.Format))
	}
	if getPronunciationOptions.CustomizationID != nil {
		builder.AddQuery("customization_id", fmt.Sprint(*getPronunciationOptions.CustomizationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = textToSpeech.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPronunciation)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CreateCustomModel : Create a custom model
// Creates a new empty custom model. You must specify a name for the new custom model. You can optionally specify the
// language and a description for the new model. The model is owned by the instance of the service whose credentials are
// used to create it.
//
// **See also:** [Creating a custom
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customModels#cuModelsCreate).
//
// ### Important voice updates
//
//  The service's voices underwent significant change on 2 December 2020.
// * The Arabic, Chinese, Dutch, Australian English, and Korean voices are now neural instead of concatenative.
// * The `ar-AR_OmarVoice` voice is deprecated. Use `ar-MS_OmarVoice` voice instead.
// * The `ar-AR` language identifier cannot be used to create a custom model. Use the `ar-MS` identifier instead.
// * The standard concatenative voices for the following languages are now deprecated: Brazilian Portuguese, United
// Kingdom and United States English, French, German, Italian, Japanese, and Spanish (all dialects).
// * The features expressive SSML, voice transformation SSML, and use of the `volume` attribute of the `<prosody>`
// element are deprecated and are not supported with any of the service's neural voices.
// * All of the service's voices are now customizable and generally available (GA) for production use.
//
// The deprecated voices and features will continue to function for at least one year but might be removed at a future
// date. You are encouraged to migrate to the equivalent neural voices at your earliest convenience. For more
// information about all voice updates, see the [2 December 2020 service
// update](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-release-notes#December2020) in the release
// notes.
func (textToSpeech *TextToSpeechV1) CreateCustomModel(createCustomModelOptions *CreateCustomModelOptions) (result *CustomModel, response *core.DetailedResponse, err error) {
	return textToSpeech.CreateCustomModelWithContext(context.Background(), createCustomModelOptions)
}

// CreateCustomModelWithContext is an alternate form of the CreateCustomModel method which supports a Context parameter
func (textToSpeech *TextToSpeechV1) CreateCustomModelWithContext(ctx context.Context, createCustomModelOptions *CreateCustomModelOptions) (result *CustomModel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createCustomModelOptions, "createCustomModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createCustomModelOptions, "createCustomModelOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = textToSpeech.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(textToSpeech.Service.Options.URL, `/v1/customizations`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createCustomModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "CreateCustomModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createCustomModelOptions.Name != nil {
		body["name"] = createCustomModelOptions.Name
	}
	if createCustomModelOptions.Language != nil {
		body["language"] = createCustomModelOptions.Language
	}
	if createCustomModelOptions.Description != nil {
		body["description"] = createCustomModelOptions.Description
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
	response, err = textToSpeech.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCustomModel)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ListCustomModels : List custom models
// Lists metadata such as the name and description for all custom models that are owned by an instance of the service.
// Specify a language to list the custom models for that language only. To see the words in addition to the metadata for
// a specific custom model, use the **List a custom model** method. You must use credentials for the instance of the
// service that owns a model to list information about it.
//
// **See also:** [Querying all custom
// models](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customModels#cuModelsQueryAll).
func (textToSpeech *TextToSpeechV1) ListCustomModels(listCustomModelsOptions *ListCustomModelsOptions) (result *CustomModels, response *core.DetailedResponse, err error) {
	return textToSpeech.ListCustomModelsWithContext(context.Background(), listCustomModelsOptions)
}

// ListCustomModelsWithContext is an alternate form of the ListCustomModels method which supports a Context parameter
func (textToSpeech *TextToSpeechV1) ListCustomModelsWithContext(ctx context.Context, listCustomModelsOptions *ListCustomModelsOptions) (result *CustomModels, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listCustomModelsOptions, "listCustomModelsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = textToSpeech.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(textToSpeech.Service.Options.URL, `/v1/customizations`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listCustomModelsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "ListCustomModels")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listCustomModelsOptions.Language != nil {
		builder.AddQuery("language", fmt.Sprint(*listCustomModelsOptions.Language))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = textToSpeech.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCustomModels)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UpdateCustomModel : Update a custom model
// Updates information for the specified custom model. You can update metadata such as the name and description of the
// model. You can also update the words in the model and their translations. Adding a new translation for a word that
// already exists in a custom model overwrites the word's existing translation. A custom model can contain no more than
// 20,000 entries. You must use credentials for the instance of the service that owns a model to update it.
//
// You can define sounds-like or phonetic translations for words. A sounds-like translation consists of one or more
// words that, when combined, sound like the word. Phonetic translations are based on the SSML phoneme format for
// representing a word. You can specify them in standard International Phonetic Alphabet (IPA) representation
//
//   <code>&lt;phoneme alphabet="ipa" ph="t&#601;m&#712;&#593;to"&gt;&lt;/phoneme&gt;</code>
//
//   or in the proprietary IBM Symbolic Phonetic Representation (SPR)
//
//   <code>&lt;phoneme alphabet="ibm" ph="1gAstroEntxrYFXs"&gt;&lt;/phoneme&gt;</code>
//
// **See also:**
// * [Updating a custom
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customModels#cuModelsUpdate)
// * [Adding words to a Japanese custom
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customWords#cuJapaneseAdd)
// * [Understanding
// customization](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customIntro#customIntro).
func (textToSpeech *TextToSpeechV1) UpdateCustomModel(updateCustomModelOptions *UpdateCustomModelOptions) (response *core.DetailedResponse, err error) {
	return textToSpeech.UpdateCustomModelWithContext(context.Background(), updateCustomModelOptions)
}

// UpdateCustomModelWithContext is an alternate form of the UpdateCustomModel method which supports a Context parameter
func (textToSpeech *TextToSpeechV1) UpdateCustomModelWithContext(ctx context.Context, updateCustomModelOptions *UpdateCustomModelOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateCustomModelOptions, "updateCustomModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateCustomModelOptions, "updateCustomModelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *updateCustomModelOptions.CustomizationID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = textToSpeech.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(textToSpeech.Service.Options.URL, `/v1/customizations/{customization_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateCustomModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "UpdateCustomModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateCustomModelOptions.Name != nil {
		body["name"] = updateCustomModelOptions.Name
	}
	if updateCustomModelOptions.Description != nil {
		body["description"] = updateCustomModelOptions.Description
	}
	if updateCustomModelOptions.Words != nil {
		body["words"] = updateCustomModelOptions.Words
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = textToSpeech.Service.Request(request, nil)

	return
}

// GetCustomModel : Get a custom model
// Gets all information about a specified custom model. In addition to metadata such as the name and description of the
// custom model, the output includes the words and their translations as defined in the model. To see just the metadata
// for a model, use the **List custom models** method.
//
// **See also:** [Querying a custom
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customModels#cuModelsQuery).
func (textToSpeech *TextToSpeechV1) GetCustomModel(getCustomModelOptions *GetCustomModelOptions) (result *CustomModel, response *core.DetailedResponse, err error) {
	return textToSpeech.GetCustomModelWithContext(context.Background(), getCustomModelOptions)
}

// GetCustomModelWithContext is an alternate form of the GetCustomModel method which supports a Context parameter
func (textToSpeech *TextToSpeechV1) GetCustomModelWithContext(ctx context.Context, getCustomModelOptions *GetCustomModelOptions) (result *CustomModel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getCustomModelOptions, "getCustomModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getCustomModelOptions, "getCustomModelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *getCustomModelOptions.CustomizationID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = textToSpeech.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(textToSpeech.Service.Options.URL, `/v1/customizations/{customization_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCustomModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "GetCustomModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = textToSpeech.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCustomModel)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteCustomModel : Delete a custom model
// Deletes the specified custom model. You must use credentials for the instance of the service that owns a model to
// delete it.
//
// **See also:** [Deleting a custom
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customModels#cuModelsDelete).
func (textToSpeech *TextToSpeechV1) DeleteCustomModel(deleteCustomModelOptions *DeleteCustomModelOptions) (response *core.DetailedResponse, err error) {
	return textToSpeech.DeleteCustomModelWithContext(context.Background(), deleteCustomModelOptions)
}

// DeleteCustomModelWithContext is an alternate form of the DeleteCustomModel method which supports a Context parameter
func (textToSpeech *TextToSpeechV1) DeleteCustomModelWithContext(ctx context.Context, deleteCustomModelOptions *DeleteCustomModelOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteCustomModelOptions, "deleteCustomModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteCustomModelOptions, "deleteCustomModelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *deleteCustomModelOptions.CustomizationID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = textToSpeech.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(textToSpeech.Service.Options.URL, `/v1/customizations/{customization_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteCustomModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "DeleteCustomModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = textToSpeech.Service.Request(request, nil)

	return
}

// AddWords : Add custom words
// Adds one or more words and their translations to the specified custom model. Adding a new translation for a word that
// already exists in a custom model overwrites the word's existing translation. A custom model can contain no more than
// 20,000 entries. You must use credentials for the instance of the service that owns a model to add words to it.
//
// You can define sounds-like or phonetic translations for words. A sounds-like translation consists of one or more
// words that, when combined, sound like the word. Phonetic translations are based on the SSML phoneme format for
// representing a word. You can specify them in standard International Phonetic Alphabet (IPA) representation
//
//   <code>&lt;phoneme alphabet="ipa" ph="t&#601;m&#712;&#593;to"&gt;&lt;/phoneme&gt;</code>
//
//   or in the proprietary IBM Symbolic Phonetic Representation (SPR)
//
//   <code>&lt;phoneme alphabet="ibm" ph="1gAstroEntxrYFXs"&gt;&lt;/phoneme&gt;</code>
//
// **See also:**
// * [Adding multiple words to a custom
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customWords#cuWordsAdd)
// * [Adding words to a Japanese custom
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customWords#cuJapaneseAdd)
// * [Understanding
// customization](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customIntro#customIntro).
func (textToSpeech *TextToSpeechV1) AddWords(addWordsOptions *AddWordsOptions) (response *core.DetailedResponse, err error) {
	return textToSpeech.AddWordsWithContext(context.Background(), addWordsOptions)
}

// AddWordsWithContext is an alternate form of the AddWords method which supports a Context parameter
func (textToSpeech *TextToSpeechV1) AddWordsWithContext(ctx context.Context, addWordsOptions *AddWordsOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addWordsOptions, "addWordsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addWordsOptions, "addWordsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *addWordsOptions.CustomizationID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = textToSpeech.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(textToSpeech.Service.Options.URL, `/v1/customizations/{customization_id}/words`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range addWordsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "AddWords")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if addWordsOptions.Words != nil {
		body["words"] = addWordsOptions.Words
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = textToSpeech.Service.Request(request, nil)

	return
}

// ListWords : List custom words
// Lists all of the words and their translations for the specified custom model. The output shows the translations as
// they are defined in the model. You must use credentials for the instance of the service that owns a model to list its
// words.
//
// **See also:** [Querying all words from a custom
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customWords#cuWordsQueryModel).
func (textToSpeech *TextToSpeechV1) ListWords(listWordsOptions *ListWordsOptions) (result *Words, response *core.DetailedResponse, err error) {
	return textToSpeech.ListWordsWithContext(context.Background(), listWordsOptions)
}

// ListWordsWithContext is an alternate form of the ListWords method which supports a Context parameter
func (textToSpeech *TextToSpeechV1) ListWordsWithContext(ctx context.Context, listWordsOptions *ListWordsOptions) (result *Words, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listWordsOptions, "listWordsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listWordsOptions, "listWordsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *listWordsOptions.CustomizationID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = textToSpeech.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(textToSpeech.Service.Options.URL, `/v1/customizations/{customization_id}/words`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listWordsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "ListWords")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = textToSpeech.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalWords)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// AddWord : Add a custom word
// Adds a single word and its translation to the specified custom model. Adding a new translation for a word that
// already exists in a custom model overwrites the word's existing translation. A custom model can contain no more than
// 20,000 entries. You must use credentials for the instance of the service that owns a model to add a word to it.
//
// You can define sounds-like or phonetic translations for words. A sounds-like translation consists of one or more
// words that, when combined, sound like the word. Phonetic translations are based on the SSML phoneme format for
// representing a word. You can specify them in standard International Phonetic Alphabet (IPA) representation
//
//   <code>&lt;phoneme alphabet="ipa" ph="t&#601;m&#712;&#593;to"&gt;&lt;/phoneme&gt;</code>
//
//   or in the proprietary IBM Symbolic Phonetic Representation (SPR)
//
//   <code>&lt;phoneme alphabet="ibm" ph="1gAstroEntxrYFXs"&gt;&lt;/phoneme&gt;</code>
//
// **See also:**
// * [Adding a single word to a custom
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customWords#cuWordAdd)
// * [Adding words to a Japanese custom
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customWords#cuJapaneseAdd)
// * [Understanding
// customization](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customIntro#customIntro).
func (textToSpeech *TextToSpeechV1) AddWord(addWordOptions *AddWordOptions) (response *core.DetailedResponse, err error) {
	return textToSpeech.AddWordWithContext(context.Background(), addWordOptions)
}

// AddWordWithContext is an alternate form of the AddWord method which supports a Context parameter
func (textToSpeech *TextToSpeechV1) AddWordWithContext(ctx context.Context, addWordOptions *AddWordOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addWordOptions, "addWordOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addWordOptions, "addWordOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *addWordOptions.CustomizationID,
		"word":             *addWordOptions.Word,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = textToSpeech.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(textToSpeech.Service.Options.URL, `/v1/customizations/{customization_id}/words/{word}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range addWordOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "AddWord")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if addWordOptions.Translation != nil {
		body["translation"] = addWordOptions.Translation
	}
	if addWordOptions.PartOfSpeech != nil {
		body["part_of_speech"] = addWordOptions.PartOfSpeech
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = textToSpeech.Service.Request(request, nil)

	return
}

// GetWord : Get a custom word
// Gets the translation for a single word from the specified custom model. The output shows the translation as it is
// defined in the model. You must use credentials for the instance of the service that owns a model to list its words.
//
// **See also:** [Querying a single word from a custom
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customWords#cuWordQueryModel).
func (textToSpeech *TextToSpeechV1) GetWord(getWordOptions *GetWordOptions) (result *Translation, response *core.DetailedResponse, err error) {
	return textToSpeech.GetWordWithContext(context.Background(), getWordOptions)
}

// GetWordWithContext is an alternate form of the GetWord method which supports a Context parameter
func (textToSpeech *TextToSpeechV1) GetWordWithContext(ctx context.Context, getWordOptions *GetWordOptions) (result *Translation, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getWordOptions, "getWordOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getWordOptions, "getWordOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *getWordOptions.CustomizationID,
		"word":             *getWordOptions.Word,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = textToSpeech.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(textToSpeech.Service.Options.URL, `/v1/customizations/{customization_id}/words/{word}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getWordOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "GetWord")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = textToSpeech.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTranslation)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteWord : Delete a custom word
// Deletes a single word from the specified custom model. You must use credentials for the instance of the service that
// owns a model to delete its words.
//
// **See also:** [Deleting a word from a custom
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customWords#cuWordDelete).
func (textToSpeech *TextToSpeechV1) DeleteWord(deleteWordOptions *DeleteWordOptions) (response *core.DetailedResponse, err error) {
	return textToSpeech.DeleteWordWithContext(context.Background(), deleteWordOptions)
}

// DeleteWordWithContext is an alternate form of the DeleteWord method which supports a Context parameter
func (textToSpeech *TextToSpeechV1) DeleteWordWithContext(ctx context.Context, deleteWordOptions *DeleteWordOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteWordOptions, "deleteWordOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteWordOptions, "deleteWordOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *deleteWordOptions.CustomizationID,
		"word":             *deleteWordOptions.Word,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = textToSpeech.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(textToSpeech.Service.Options.URL, `/v1/customizations/{customization_id}/words/{word}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteWordOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "DeleteWord")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = textToSpeech.Service.Request(request, nil)

	return
}

// DeleteUserData : Delete labeled data
// Deletes all data that is associated with a specified customer ID. The method deletes all data for the customer ID,
// regardless of the method by which the information was added. The method has no effect if no data is associated with
// the customer ID. You must issue the request with credentials for the same instance of the service that was used to
// associate the customer ID with the data. You associate a customer ID with data by passing the `X-Watson-Metadata`
// header with a request that passes the data.
//
// **Note:** If you delete an instance of the service from the service console, all data associated with that service
// instance is automatically deleted. This includes all custom models and word/translation pairs, and all data related
// to speech synthesis requests.
//
// **See also:** [Information
// security](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-information-security#information-security).
func (textToSpeech *TextToSpeechV1) DeleteUserData(deleteUserDataOptions *DeleteUserDataOptions) (response *core.DetailedResponse, err error) {
	return textToSpeech.DeleteUserDataWithContext(context.Background(), deleteUserDataOptions)
}

// DeleteUserDataWithContext is an alternate form of the DeleteUserData method which supports a Context parameter
func (textToSpeech *TextToSpeechV1) DeleteUserDataWithContext(ctx context.Context, deleteUserDataOptions *DeleteUserDataOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteUserDataOptions, "deleteUserDataOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteUserDataOptions, "deleteUserDataOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = textToSpeech.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(textToSpeech.Service.Options.URL, `/v1/user_data`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteUserDataOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "DeleteUserData")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("customer_id", fmt.Sprint(*deleteUserDataOptions.CustomerID))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = textToSpeech.Service.Request(request, nil)

	return
}

// AddWordOptions : The AddWord options.
type AddWordOptions struct {
	// The customization ID (GUID) of the custom model. You must make the request with credentials for the instance of the
	// service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// The word that is to be added or updated for the custom model.
	Word *string `json:"word" validate:"required,ne="`

	// The phonetic or sounds-like translation for the word. A phonetic translation is based on the SSML format for
	// representing the phonetic string of a word either as an IPA translation or as an IBM SPR translation. The Arabic,
	// Chinese, Dutch, Australian English, and Korean languages support only IPA. A sounds-like is one or more words that,
	// when combined, sound like the word.
	Translation *string `json:"translation" validate:"required"`

	// **Japanese only.** The part of speech for the word. The service uses the value to produce the correct intonation for
	// the word. You can create only a single entry, with or without a single part of speech, for any word; you cannot
	// create multiple entries with different parts of speech for the same word. For more information, see [Working with
	// Japanese entries](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-rules#jaNotes).
	PartOfSpeech *string `json:"part_of_speech,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the AddWordOptions.PartOfSpeech property.
// **Japanese only.** The part of speech for the word. The service uses the value to produce the correct intonation for
// the word. You can create only a single entry, with or without a single part of speech, for any word; you cannot
// create multiple entries with different parts of speech for the same word. For more information, see [Working with
// Japanese entries](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-rules#jaNotes).
const (
	AddWordOptionsPartOfSpeechDosiConst = "Dosi"
	AddWordOptionsPartOfSpeechFukuConst = "Fuku"
	AddWordOptionsPartOfSpeechGobiConst = "Gobi"
	AddWordOptionsPartOfSpeechHokaConst = "Hoka"
	AddWordOptionsPartOfSpeechJodoConst = "Jodo"
	AddWordOptionsPartOfSpeechJosiConst = "Josi"
	AddWordOptionsPartOfSpeechKatoConst = "Kato"
	AddWordOptionsPartOfSpeechKedoConst = "Kedo"
	AddWordOptionsPartOfSpeechKeyoConst = "Keyo"
	AddWordOptionsPartOfSpeechKigoConst = "Kigo"
	AddWordOptionsPartOfSpeechKoyuConst = "Koyu"
	AddWordOptionsPartOfSpeechMesiConst = "Mesi"
	AddWordOptionsPartOfSpeechRetaConst = "Reta"
	AddWordOptionsPartOfSpeechStbiConst = "Stbi"
	AddWordOptionsPartOfSpeechSttoConst = "Stto"
	AddWordOptionsPartOfSpeechStzoConst = "Stzo"
	AddWordOptionsPartOfSpeechSujiConst = "Suji"
)

// NewAddWordOptions : Instantiate AddWordOptions
func (*TextToSpeechV1) NewAddWordOptions(customizationID string, word string, translation string) *AddWordOptions {
	return &AddWordOptions{
		CustomizationID: core.StringPtr(customizationID),
		Word:            core.StringPtr(word),
		Translation:     core.StringPtr(translation),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *AddWordOptions) SetCustomizationID(customizationID string) *AddWordOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetWord : Allow user to set Word
func (options *AddWordOptions) SetWord(word string) *AddWordOptions {
	options.Word = core.StringPtr(word)
	return options
}

// SetTranslation : Allow user to set Translation
func (options *AddWordOptions) SetTranslation(translation string) *AddWordOptions {
	options.Translation = core.StringPtr(translation)
	return options
}

// SetPartOfSpeech : Allow user to set PartOfSpeech
func (options *AddWordOptions) SetPartOfSpeech(partOfSpeech string) *AddWordOptions {
	options.PartOfSpeech = core.StringPtr(partOfSpeech)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *AddWordOptions) SetHeaders(param map[string]string) *AddWordOptions {
	options.Headers = param
	return options
}

// AddWordsOptions : The AddWords options.
type AddWordsOptions struct {
	// The customization ID (GUID) of the custom model. You must make the request with credentials for the instance of the
	// service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// The **Add custom words** method accepts an array of `Word` objects. Each object provides a word that is to be added
	// or updated for the custom model and the word's translation.
	//
	// The **List custom words** method returns an array of `Word` objects. Each object shows a word and its translation
	// from the custom model. The words are listed in alphabetical order, with uppercase letters listed before lowercase
	// letters. The array is empty if the custom model contains no words.
	Words []Word `json:"words" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAddWordsOptions : Instantiate AddWordsOptions
func (*TextToSpeechV1) NewAddWordsOptions(customizationID string, words []Word) *AddWordsOptions {
	return &AddWordsOptions{
		CustomizationID: core.StringPtr(customizationID),
		Words:           words,
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *AddWordsOptions) SetCustomizationID(customizationID string) *AddWordsOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetWords : Allow user to set Words
func (options *AddWordsOptions) SetWords(words []Word) *AddWordsOptions {
	options.Words = words
	return options
}

// SetHeaders : Allow user to set Headers
func (options *AddWordsOptions) SetHeaders(param map[string]string) *AddWordsOptions {
	options.Headers = param
	return options
}

// CreateCustomModelOptions : The CreateCustomModel options.
type CreateCustomModelOptions struct {
	// The name of the new custom model.
	Name *string `json:"name" validate:"required"`

	// The language of the new custom model. You create a custom model for a specific language, not for a specific voice. A
	// custom model can be used with any voice for its specified language. Omit the parameter to use the the default
	// language, `en-US`. **Note:** The `ar-AR` language identifier cannot be used to create a custom model. Use the
	// `ar-MS` identifier instead.
	Language *string `json:"language,omitempty"`

	// A description of the new custom model. Specifying a description is recommended.
	Description *string `json:"description,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateCustomModelOptions.Language property.
// The language of the new custom model. You create a custom model for a specific language, not for a specific voice. A
// custom model can be used with any voice for its specified language. Omit the parameter to use the the default
// language, `en-US`. **Note:** The `ar-AR` language identifier cannot be used to create a custom model. Use the `ar-MS`
// identifier instead.
const (
	CreateCustomModelOptionsLanguageArMsConst = "ar-MS"
	CreateCustomModelOptionsLanguageDeDeConst = "de-DE"
	CreateCustomModelOptionsLanguageEnGbConst = "en-GB"
	CreateCustomModelOptionsLanguageEnUsConst = "en-US"
	CreateCustomModelOptionsLanguageEsEsConst = "es-ES"
	CreateCustomModelOptionsLanguageEsLaConst = "es-LA"
	CreateCustomModelOptionsLanguageEsUsConst = "es-US"
	CreateCustomModelOptionsLanguageFrFrConst = "fr-FR"
	CreateCustomModelOptionsLanguageItItConst = "it-IT"
	CreateCustomModelOptionsLanguageJaJpConst = "ja-JP"
	CreateCustomModelOptionsLanguageKoKrConst = "ko-KR"
	CreateCustomModelOptionsLanguageNlNlConst = "nl-NL"
	CreateCustomModelOptionsLanguagePtBrConst = "pt-BR"
	CreateCustomModelOptionsLanguageZhCnConst = "zh-CN"
)

// NewCreateCustomModelOptions : Instantiate CreateCustomModelOptions
func (*TextToSpeechV1) NewCreateCustomModelOptions(name string) *CreateCustomModelOptions {
	return &CreateCustomModelOptions{
		Name: core.StringPtr(name),
	}
}

// SetName : Allow user to set Name
func (options *CreateCustomModelOptions) SetName(name string) *CreateCustomModelOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetLanguage : Allow user to set Language
func (options *CreateCustomModelOptions) SetLanguage(language string) *CreateCustomModelOptions {
	options.Language = core.StringPtr(language)
	return options
}

// SetDescription : Allow user to set Description
func (options *CreateCustomModelOptions) SetDescription(description string) *CreateCustomModelOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateCustomModelOptions) SetHeaders(param map[string]string) *CreateCustomModelOptions {
	options.Headers = param
	return options
}

// CustomModel : Information about an existing custom model.
type CustomModel struct {
	// The customization ID (GUID) of the custom model. The **Create a custom model** method returns only this field. It
	// does not not return the other fields of this object.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// The name of the custom model.
	Name *string `json:"name,omitempty"`

	// The language identifier of the custom model (for example, `en-US`).
	Language *string `json:"language,omitempty"`

	// The GUID of the credentials for the instance of the service that owns the custom model.
	Owner *string `json:"owner,omitempty"`

	// The date and time in Coordinated Universal Time (UTC) at which the custom model was created. The value is provided
	// in full ISO 8601 format (`YYYY-MM-DDThh:mm:ss.sTZD`).
	Created *string `json:"created,omitempty"`

	// The date and time in Coordinated Universal Time (UTC) at which the custom model was last modified. The `created` and
	// `updated` fields are equal when a model is first added but has yet to be updated. The value is provided in full ISO
	// 8601 format (`YYYY-MM-DDThh:mm:ss.sTZD`).
	LastModified *string `json:"last_modified,omitempty"`

	// The description of the custom model.
	Description *string `json:"description,omitempty"`

	// An array of `Word` objects that lists the words and their translations from the custom model. The words are listed
	// in alphabetical order, with uppercase letters listed before lowercase letters. The array is empty if the custom
	// model contains no words. This field is returned only by the **Get a voice** method and only when you specify the
	// customization ID of a custom model.
	Words []Word `json:"words,omitempty"`
}

// UnmarshalCustomModel unmarshals an instance of CustomModel from the specified map of raw messages.
func UnmarshalCustomModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CustomModel)
	err = core.UnmarshalPrimitive(m, "customization_id", &obj.CustomizationID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "language", &obj.Language)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "owner", &obj.Owner)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created", &obj.Created)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_modified", &obj.LastModified)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "words", &obj.Words, UnmarshalWord)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CustomModels : Information about existing custom models.
type CustomModels struct {
	// An array of `CustomModel` objects that provides information about each available custom model. The array is empty if
	// the requesting credentials own no custom models (if no language is specified) or own no custom models for the
	// specified language.
	Customizations []CustomModel `json:"customizations" validate:"required"`
}

// UnmarshalCustomModels unmarshals an instance of CustomModels from the specified map of raw messages.
func UnmarshalCustomModels(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CustomModels)
	err = core.UnmarshalModel(m, "customizations", &obj.Customizations, UnmarshalCustomModel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteCustomModelOptions : The DeleteCustomModel options.
type DeleteCustomModelOptions struct {
	// The customization ID (GUID) of the custom model. You must make the request with credentials for the instance of the
	// service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteCustomModelOptions : Instantiate DeleteCustomModelOptions
func (*TextToSpeechV1) NewDeleteCustomModelOptions(customizationID string) *DeleteCustomModelOptions {
	return &DeleteCustomModelOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *DeleteCustomModelOptions) SetCustomizationID(customizationID string) *DeleteCustomModelOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteCustomModelOptions) SetHeaders(param map[string]string) *DeleteCustomModelOptions {
	options.Headers = param
	return options
}

// DeleteUserDataOptions : The DeleteUserData options.
type DeleteUserDataOptions struct {
	// The customer ID for which all data is to be deleted.
	CustomerID *string `json:"customer_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteUserDataOptions : Instantiate DeleteUserDataOptions
func (*TextToSpeechV1) NewDeleteUserDataOptions(customerID string) *DeleteUserDataOptions {
	return &DeleteUserDataOptions{
		CustomerID: core.StringPtr(customerID),
	}
}

// SetCustomerID : Allow user to set CustomerID
func (options *DeleteUserDataOptions) SetCustomerID(customerID string) *DeleteUserDataOptions {
	options.CustomerID = core.StringPtr(customerID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteUserDataOptions) SetHeaders(param map[string]string) *DeleteUserDataOptions {
	options.Headers = param
	return options
}

// DeleteWordOptions : The DeleteWord options.
type DeleteWordOptions struct {
	// The customization ID (GUID) of the custom model. You must make the request with credentials for the instance of the
	// service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// The word that is to be deleted from the custom model.
	Word *string `json:"word" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteWordOptions : Instantiate DeleteWordOptions
func (*TextToSpeechV1) NewDeleteWordOptions(customizationID string, word string) *DeleteWordOptions {
	return &DeleteWordOptions{
		CustomizationID: core.StringPtr(customizationID),
		Word:            core.StringPtr(word),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *DeleteWordOptions) SetCustomizationID(customizationID string) *DeleteWordOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetWord : Allow user to set Word
func (options *DeleteWordOptions) SetWord(word string) *DeleteWordOptions {
	options.Word = core.StringPtr(word)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteWordOptions) SetHeaders(param map[string]string) *DeleteWordOptions {
	options.Headers = param
	return options
}

// GetCustomModelOptions : The GetCustomModel options.
type GetCustomModelOptions struct {
	// The customization ID (GUID) of the custom model. You must make the request with credentials for the instance of the
	// service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCustomModelOptions : Instantiate GetCustomModelOptions
func (*TextToSpeechV1) NewGetCustomModelOptions(customizationID string) *GetCustomModelOptions {
	return &GetCustomModelOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *GetCustomModelOptions) SetCustomizationID(customizationID string) *GetCustomModelOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetCustomModelOptions) SetHeaders(param map[string]string) *GetCustomModelOptions {
	options.Headers = param
	return options
}

// GetPronunciationOptions : The GetPronunciation options.
type GetPronunciationOptions struct {
	// The word for which the pronunciation is requested.
	Text *string `json:"text" validate:"required"`

	// A voice that specifies the language in which the pronunciation is to be returned. All voices for the same language
	// (for example, `en-US`) return the same translation. For more information about specifying a voice, see **Important
	// voice updates** in the method description.
	Voice *string `json:"voice,omitempty"`

	// The phoneme format in which to return the pronunciation. The Arabic, Chinese, Dutch, Australian English, and Korean
	// languages support only IPA. Omit the parameter to obtain the pronunciation in the default format.
	Format *string `json:"format,omitempty"`

	// The customization ID (GUID) of a custom model for which the pronunciation is to be returned. The language of a
	// specified custom model must match the language of the specified voice. If the word is not defined in the specified
	// custom model, the service returns the default translation for the custom model's language. You must make the request
	// with credentials for the instance of the service that owns the custom model. Omit the parameter to see the
	// translation for the specified voice with no customization.
	CustomizationID *string `json:"customization_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetPronunciationOptions.Voice property.
// A voice that specifies the language in which the pronunciation is to be returned. All voices for the same language
// (for example, `en-US`) return the same translation. For more information about specifying a voice, see **Important
// voice updates** in the method description.
const (
	GetPronunciationOptionsVoiceArArOmarvoiceConst        = "ar-AR_OmarVoice"
	GetPronunciationOptionsVoiceArMsOmarvoiceConst        = "ar-MS_OmarVoice"
	GetPronunciationOptionsVoiceDeDeBirgitv3voiceConst    = "de-DE_BirgitV3Voice"
	GetPronunciationOptionsVoiceDeDeBirgitvoiceConst      = "de-DE_BirgitVoice"
	GetPronunciationOptionsVoiceDeDeDieterv3voiceConst    = "de-DE_DieterV3Voice"
	GetPronunciationOptionsVoiceDeDeDietervoiceConst      = "de-DE_DieterVoice"
	GetPronunciationOptionsVoiceDeDeErikav3voiceConst     = "de-DE_ErikaV3Voice"
	GetPronunciationOptionsVoiceEnAuCraigvoiceConst       = "en-AU-CraigVoice"
	GetPronunciationOptionsVoiceEnAuMadisonvoiceConst     = "en-AU-MadisonVoice"
	GetPronunciationOptionsVoiceEnGbCharlottev3voiceConst = "en-GB_CharlotteV3Voice"
	GetPronunciationOptionsVoiceEnGbJamesv3voiceConst     = "en-GB_JamesV3Voice"
	GetPronunciationOptionsVoiceEnGbKatev3voiceConst      = "en-GB_KateV3Voice"
	GetPronunciationOptionsVoiceEnGbKatevoiceConst        = "en-GB_KateVoice"
	GetPronunciationOptionsVoiceEnUsAllisonv3voiceConst   = "en-US_AllisonV3Voice"
	GetPronunciationOptionsVoiceEnUsAllisonvoiceConst     = "en-US_AllisonVoice"
	GetPronunciationOptionsVoiceEnUsEmilyv3voiceConst     = "en-US_EmilyV3Voice"
	GetPronunciationOptionsVoiceEnUsHenryv3voiceConst     = "en-US_HenryV3Voice"
	GetPronunciationOptionsVoiceEnUsKevinv3voiceConst     = "en-US_KevinV3Voice"
	GetPronunciationOptionsVoiceEnUsLisav3voiceConst      = "en-US_LisaV3Voice"
	GetPronunciationOptionsVoiceEnUsLisavoiceConst        = "en-US_LisaVoice"
	GetPronunciationOptionsVoiceEnUsMichaelv3voiceConst   = "en-US_MichaelV3Voice"
	GetPronunciationOptionsVoiceEnUsMichaelvoiceConst     = "en-US_MichaelVoice"
	GetPronunciationOptionsVoiceEnUsOliviav3voiceConst    = "en-US_OliviaV3Voice"
	GetPronunciationOptionsVoiceEsEsEnriquev3voiceConst   = "es-ES_EnriqueV3Voice"
	GetPronunciationOptionsVoiceEsEsEnriquevoiceConst     = "es-ES_EnriqueVoice"
	GetPronunciationOptionsVoiceEsEsLaurav3voiceConst     = "es-ES_LauraV3Voice"
	GetPronunciationOptionsVoiceEsEsLauravoiceConst       = "es-ES_LauraVoice"
	GetPronunciationOptionsVoiceEsLaSofiav3voiceConst     = "es-LA_SofiaV3Voice"
	GetPronunciationOptionsVoiceEsLaSofiavoiceConst       = "es-LA_SofiaVoice"
	GetPronunciationOptionsVoiceEsUsSofiav3voiceConst     = "es-US_SofiaV3Voice"
	GetPronunciationOptionsVoiceEsUsSofiavoiceConst       = "es-US_SofiaVoice"
	GetPronunciationOptionsVoiceFrFrNicolasv3voiceConst   = "fr-FR_NicolasV3Voice"
	GetPronunciationOptionsVoiceFrFrReneev3voiceConst     = "fr-FR_ReneeV3Voice"
	GetPronunciationOptionsVoiceFrFrReneevoiceConst       = "fr-FR_ReneeVoice"
	GetPronunciationOptionsVoiceItItFrancescav3voiceConst = "it-IT_FrancescaV3Voice"
	GetPronunciationOptionsVoiceItItFrancescavoiceConst   = "it-IT_FrancescaVoice"
	GetPronunciationOptionsVoiceJaJpEmiv3voiceConst       = "ja-JP_EmiV3Voice"
	GetPronunciationOptionsVoiceJaJpEmivoiceConst         = "ja-JP_EmiVoice"
	GetPronunciationOptionsVoiceKoKrHyunjunvoiceConst     = "ko-KR_HyunjunVoice"
	GetPronunciationOptionsVoiceKoKrSiwoovoiceConst       = "ko-KR_SiWooVoice"
	GetPronunciationOptionsVoiceKoKrYoungmivoiceConst     = "ko-KR_YoungmiVoice"
	GetPronunciationOptionsVoiceKoKrYunavoiceConst        = "ko-KR_YunaVoice"
	GetPronunciationOptionsVoiceNlNlEmmavoiceConst        = "nl-NL_EmmaVoice"
	GetPronunciationOptionsVoiceNlNlLiamvoiceConst        = "nl-NL_LiamVoice"
	GetPronunciationOptionsVoicePtBrIsabelav3voiceConst   = "pt-BR_IsabelaV3Voice"
	GetPronunciationOptionsVoicePtBrIsabelavoiceConst     = "pt-BR_IsabelaVoice"
	GetPronunciationOptionsVoiceZhCnLinavoiceConst        = "zh-CN_LiNaVoice"
	GetPronunciationOptionsVoiceZhCnWangweivoiceConst     = "zh-CN_WangWeiVoice"
	GetPronunciationOptionsVoiceZhCnZhangjingvoiceConst   = "zh-CN_ZhangJingVoice"
)

// Constants associated with the GetPronunciationOptions.Format property.
// The phoneme format in which to return the pronunciation. The Arabic, Chinese, Dutch, Australian English, and Korean
// languages support only IPA. Omit the parameter to obtain the pronunciation in the default format.
const (
	GetPronunciationOptionsFormatIBMConst = "ibm"
	GetPronunciationOptionsFormatIpaConst = "ipa"
)

// NewGetPronunciationOptions : Instantiate GetPronunciationOptions
func (*TextToSpeechV1) NewGetPronunciationOptions(text string) *GetPronunciationOptions {
	return &GetPronunciationOptions{
		Text: core.StringPtr(text),
	}
}

// SetText : Allow user to set Text
func (options *GetPronunciationOptions) SetText(text string) *GetPronunciationOptions {
	options.Text = core.StringPtr(text)
	return options
}

// SetVoice : Allow user to set Voice
func (options *GetPronunciationOptions) SetVoice(voice string) *GetPronunciationOptions {
	options.Voice = core.StringPtr(voice)
	return options
}

// SetFormat : Allow user to set Format
func (options *GetPronunciationOptions) SetFormat(format string) *GetPronunciationOptions {
	options.Format = core.StringPtr(format)
	return options
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *GetPronunciationOptions) SetCustomizationID(customizationID string) *GetPronunciationOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetPronunciationOptions) SetHeaders(param map[string]string) *GetPronunciationOptions {
	options.Headers = param
	return options
}

// GetVoiceOptions : The GetVoice options.
type GetVoiceOptions struct {
	// The voice for which information is to be returned. For more information about specifying a voice, see **Important
	// voice updates** in the method description.
	Voice *string `json:"voice" validate:"required,ne="`

	// The customization ID (GUID) of a custom model for which information is to be returned. You must make the request
	// with credentials for the instance of the service that owns the custom model. Omit the parameter to see information
	// about the specified voice with no customization.
	CustomizationID *string `json:"customization_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetVoiceOptions.Voice property.
// The voice for which information is to be returned. For more information about specifying a voice, see **Important
// voice updates** in the method description.
const (
	GetVoiceOptionsVoiceArArOmarvoiceConst        = "ar-AR_OmarVoice"
	GetVoiceOptionsVoiceArMsOmarvoiceConst        = "ar-MS_OmarVoice"
	GetVoiceOptionsVoiceDeDeBirgitv3voiceConst    = "de-DE_BirgitV3Voice"
	GetVoiceOptionsVoiceDeDeBirgitvoiceConst      = "de-DE_BirgitVoice"
	GetVoiceOptionsVoiceDeDeDieterv3voiceConst    = "de-DE_DieterV3Voice"
	GetVoiceOptionsVoiceDeDeDietervoiceConst      = "de-DE_DieterVoice"
	GetVoiceOptionsVoiceDeDeErikav3voiceConst     = "de-DE_ErikaV3Voice"
	GetVoiceOptionsVoiceEnAuCraigvoiceConst       = "en-AU-CraigVoice"
	GetVoiceOptionsVoiceEnAuMadisonvoiceConst     = "en-AU-MadisonVoice"
	GetVoiceOptionsVoiceEnGbCharlottev3voiceConst = "en-GB_CharlotteV3Voice"
	GetVoiceOptionsVoiceEnGbJamesv3voiceConst     = "en-GB_JamesV3Voice"
	GetVoiceOptionsVoiceEnGbKatev3voiceConst      = "en-GB_KateV3Voice"
	GetVoiceOptionsVoiceEnGbKatevoiceConst        = "en-GB_KateVoice"
	GetVoiceOptionsVoiceEnUsAllisonv3voiceConst   = "en-US_AllisonV3Voice"
	GetVoiceOptionsVoiceEnUsAllisonvoiceConst     = "en-US_AllisonVoice"
	GetVoiceOptionsVoiceEnUsEmilyv3voiceConst     = "en-US_EmilyV3Voice"
	GetVoiceOptionsVoiceEnUsHenryv3voiceConst     = "en-US_HenryV3Voice"
	GetVoiceOptionsVoiceEnUsKevinv3voiceConst     = "en-US_KevinV3Voice"
	GetVoiceOptionsVoiceEnUsLisav3voiceConst      = "en-US_LisaV3Voice"
	GetVoiceOptionsVoiceEnUsLisavoiceConst        = "en-US_LisaVoice"
	GetVoiceOptionsVoiceEnUsMichaelv3voiceConst   = "en-US_MichaelV3Voice"
	GetVoiceOptionsVoiceEnUsMichaelvoiceConst     = "en-US_MichaelVoice"
	GetVoiceOptionsVoiceEnUsOliviav3voiceConst    = "en-US_OliviaV3Voice"
	GetVoiceOptionsVoiceEsEsEnriquev3voiceConst   = "es-ES_EnriqueV3Voice"
	GetVoiceOptionsVoiceEsEsEnriquevoiceConst     = "es-ES_EnriqueVoice"
	GetVoiceOptionsVoiceEsEsLaurav3voiceConst     = "es-ES_LauraV3Voice"
	GetVoiceOptionsVoiceEsEsLauravoiceConst       = "es-ES_LauraVoice"
	GetVoiceOptionsVoiceEsLaSofiav3voiceConst     = "es-LA_SofiaV3Voice"
	GetVoiceOptionsVoiceEsLaSofiavoiceConst       = "es-LA_SofiaVoice"
	GetVoiceOptionsVoiceEsUsSofiav3voiceConst     = "es-US_SofiaV3Voice"
	GetVoiceOptionsVoiceEsUsSofiavoiceConst       = "es-US_SofiaVoice"
	GetVoiceOptionsVoiceFrFrNicolasv3voiceConst   = "fr-FR_NicolasV3Voice"
	GetVoiceOptionsVoiceFrFrReneev3voiceConst     = "fr-FR_ReneeV3Voice"
	GetVoiceOptionsVoiceFrFrReneevoiceConst       = "fr-FR_ReneeVoice"
	GetVoiceOptionsVoiceItItFrancescav3voiceConst = "it-IT_FrancescaV3Voice"
	GetVoiceOptionsVoiceItItFrancescavoiceConst   = "it-IT_FrancescaVoice"
	GetVoiceOptionsVoiceJaJpEmiv3voiceConst       = "ja-JP_EmiV3Voice"
	GetVoiceOptionsVoiceJaJpEmivoiceConst         = "ja-JP_EmiVoice"
	GetVoiceOptionsVoiceKoKrHyunjunvoiceConst     = "ko-KR_HyunjunVoice"
	GetVoiceOptionsVoiceKoKrSiwoovoiceConst       = "ko-KR_SiWooVoice"
	GetVoiceOptionsVoiceKoKrYoungmivoiceConst     = "ko-KR_YoungmiVoice"
	GetVoiceOptionsVoiceKoKrYunavoiceConst        = "ko-KR_YunaVoice"
	GetVoiceOptionsVoiceNlNlEmmavoiceConst        = "nl-NL_EmmaVoice"
	GetVoiceOptionsVoiceNlNlLiamvoiceConst        = "nl-NL_LiamVoice"
	GetVoiceOptionsVoicePtBrIsabelav3voiceConst   = "pt-BR_IsabelaV3Voice"
	GetVoiceOptionsVoicePtBrIsabelavoiceConst     = "pt-BR_IsabelaVoice"
	GetVoiceOptionsVoiceZhCnLinavoiceConst        = "zh-CN_LiNaVoice"
	GetVoiceOptionsVoiceZhCnWangweivoiceConst     = "zh-CN_WangWeiVoice"
	GetVoiceOptionsVoiceZhCnZhangjingvoiceConst   = "zh-CN_ZhangJingVoice"
)

// NewGetVoiceOptions : Instantiate GetVoiceOptions
func (*TextToSpeechV1) NewGetVoiceOptions(voice string) *GetVoiceOptions {
	return &GetVoiceOptions{
		Voice: core.StringPtr(voice),
	}
}

// SetVoice : Allow user to set Voice
func (options *GetVoiceOptions) SetVoice(voice string) *GetVoiceOptions {
	options.Voice = core.StringPtr(voice)
	return options
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *GetVoiceOptions) SetCustomizationID(customizationID string) *GetVoiceOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetVoiceOptions) SetHeaders(param map[string]string) *GetVoiceOptions {
	options.Headers = param
	return options
}

// GetWordOptions : The GetWord options.
type GetWordOptions struct {
	// The customization ID (GUID) of the custom model. You must make the request with credentials for the instance of the
	// service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// The word that is to be queried from the custom model.
	Word *string `json:"word" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetWordOptions : Instantiate GetWordOptions
func (*TextToSpeechV1) NewGetWordOptions(customizationID string, word string) *GetWordOptions {
	return &GetWordOptions{
		CustomizationID: core.StringPtr(customizationID),
		Word:            core.StringPtr(word),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *GetWordOptions) SetCustomizationID(customizationID string) *GetWordOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetWord : Allow user to set Word
func (options *GetWordOptions) SetWord(word string) *GetWordOptions {
	options.Word = core.StringPtr(word)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetWordOptions) SetHeaders(param map[string]string) *GetWordOptions {
	options.Headers = param
	return options
}

// ListCustomModelsOptions : The ListCustomModels options.
type ListCustomModelsOptions struct {
	// The language for which custom models that are owned by the requesting credentials are to be returned. Omit the
	// parameter to see all custom models that are owned by the requester.
	Language *string `json:"language,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListCustomModelsOptions.Language property.
// The language for which custom models that are owned by the requesting credentials are to be returned. Omit the
// parameter to see all custom models that are owned by the requester.
const (
	ListCustomModelsOptionsLanguageArMsConst = "ar-MS"
	ListCustomModelsOptionsLanguageDeDeConst = "de-DE"
	ListCustomModelsOptionsLanguageEnAuConst = "en-AU"
	ListCustomModelsOptionsLanguageEnGbConst = "en-GB"
	ListCustomModelsOptionsLanguageEnUsConst = "en-US"
	ListCustomModelsOptionsLanguageEsEsConst = "es-ES"
	ListCustomModelsOptionsLanguageEsLaConst = "es-LA"
	ListCustomModelsOptionsLanguageEsUsConst = "es-US"
	ListCustomModelsOptionsLanguageFrFrConst = "fr-FR"
	ListCustomModelsOptionsLanguageItItConst = "it-IT"
	ListCustomModelsOptionsLanguageJaJpConst = "ja-JP"
	ListCustomModelsOptionsLanguageKoKrConst = "ko-KR"
	ListCustomModelsOptionsLanguageNlNlConst = "nl-NL"
	ListCustomModelsOptionsLanguagePtBrConst = "pt-BR"
	ListCustomModelsOptionsLanguageZhCnConst = "zh-CN"
)

// NewListCustomModelsOptions : Instantiate ListCustomModelsOptions
func (*TextToSpeechV1) NewListCustomModelsOptions() *ListCustomModelsOptions {
	return &ListCustomModelsOptions{}
}

// SetLanguage : Allow user to set Language
func (options *ListCustomModelsOptions) SetLanguage(language string) *ListCustomModelsOptions {
	options.Language = core.StringPtr(language)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListCustomModelsOptions) SetHeaders(param map[string]string) *ListCustomModelsOptions {
	options.Headers = param
	return options
}

// ListVoicesOptions : The ListVoices options.
type ListVoicesOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListVoicesOptions : Instantiate ListVoicesOptions
func (*TextToSpeechV1) NewListVoicesOptions() *ListVoicesOptions {
	return &ListVoicesOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListVoicesOptions) SetHeaders(param map[string]string) *ListVoicesOptions {
	options.Headers = param
	return options
}

// ListWordsOptions : The ListWords options.
type ListWordsOptions struct {
	// The customization ID (GUID) of the custom model. You must make the request with credentials for the instance of the
	// service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListWordsOptions : Instantiate ListWordsOptions
func (*TextToSpeechV1) NewListWordsOptions(customizationID string) *ListWordsOptions {
	return &ListWordsOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *ListWordsOptions) SetCustomizationID(customizationID string) *ListWordsOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListWordsOptions) SetHeaders(param map[string]string) *ListWordsOptions {
	options.Headers = param
	return options
}

// Pronunciation : The pronunciation of the specified text.
type Pronunciation struct {
	// The pronunciation of the specified text in the requested voice and format. If a custom model is specified, the
	// pronunciation also reflects that custom model.
	Pronunciation *string `json:"pronunciation" validate:"required"`
}

// UnmarshalPronunciation unmarshals an instance of Pronunciation from the specified map of raw messages.
func UnmarshalPronunciation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Pronunciation)
	err = core.UnmarshalPrimitive(m, "pronunciation", &obj.Pronunciation)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SupportedFeatures : Additional service features that are supported with the voice.
type SupportedFeatures struct {
	// If `true`, the voice can be customized; if `false`, the voice cannot be customized. (Same as `customizable`.).
	CustomPronunciation *bool `json:"custom_pronunciation" validate:"required"`

	// If `true`, the voice can be transformed by using the SSML &lt;voice-transformation&gt; element; if `false`, the
	// voice cannot be transformed. The feature was available only for the now-deprecated standard voices. You cannot use
	// the feature with neural voices.
	VoiceTransformation *bool `json:"voice_transformation" validate:"required"`
}

// UnmarshalSupportedFeatures unmarshals an instance of SupportedFeatures from the specified map of raw messages.
func UnmarshalSupportedFeatures(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SupportedFeatures)
	err = core.UnmarshalPrimitive(m, "custom_pronunciation", &obj.CustomPronunciation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "voice_transformation", &obj.VoiceTransformation)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SynthesizeOptions : The Synthesize options.
type SynthesizeOptions struct {
	// The text to synthesize.
	Text *string `json:"text" validate:"required"`

	// The requested format (MIME type) of the audio. You can use the `Accept` header or the `accept` parameter to specify
	// the audio format. For more information about specifying an audio format, see **Audio formats (accept types)** in the
	// method description.
	Accept *string `json:"Accept,omitempty"`

	// The voice to use for synthesis. For more information about specifying a voice, see **Important voice updates** in
	// the method description.
	Voice *string `json:"voice,omitempty"`

	// The customization ID (GUID) of a custom model to use for the synthesis. If a custom model is specified, it works
	// only if it matches the language of the indicated voice. You must make the request with credentials for the instance
	// of the service that owns the custom model. Omit the parameter to use the specified voice with no customization.
	CustomizationID *string `json:"customization_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the SynthesizeOptions.Voice property.
// The voice to use for synthesis. For more information about specifying a voice, see **Important voice updates** in the
// method description.
const (
	SynthesizeOptionsVoiceArArOmarvoiceConst        = "ar-AR_OmarVoice"
	SynthesizeOptionsVoiceArMsOmarvoiceConst        = "ar-MS_OmarVoice"
	SynthesizeOptionsVoiceDeDeBirgitv3voiceConst    = "de-DE_BirgitV3Voice"
	SynthesizeOptionsVoiceDeDeBirgitvoiceConst      = "de-DE_BirgitVoice"
	SynthesizeOptionsVoiceDeDeDieterv3voiceConst    = "de-DE_DieterV3Voice"
	SynthesizeOptionsVoiceDeDeDietervoiceConst      = "de-DE_DieterVoice"
	SynthesizeOptionsVoiceDeDeErikav3voiceConst     = "de-DE_ErikaV3Voice"
	SynthesizeOptionsVoiceEnAuCraigvoiceConst       = "en-AU-CraigVoice"
	SynthesizeOptionsVoiceEnAuMadisonvoiceConst     = "en-AU-MadisonVoice"
	SynthesizeOptionsVoiceEnGbCharlottev3voiceConst = "en-GB_CharlotteV3Voice"
	SynthesizeOptionsVoiceEnGbJamesv3voiceConst     = "en-GB_JamesV3Voice"
	SynthesizeOptionsVoiceEnGbKatev3voiceConst      = "en-GB_KateV3Voice"
	SynthesizeOptionsVoiceEnGbKatevoiceConst        = "en-GB_KateVoice"
	SynthesizeOptionsVoiceEnUsAllisonv3voiceConst   = "en-US_AllisonV3Voice"
	SynthesizeOptionsVoiceEnUsAllisonvoiceConst     = "en-US_AllisonVoice"
	SynthesizeOptionsVoiceEnUsEmilyv3voiceConst     = "en-US_EmilyV3Voice"
	SynthesizeOptionsVoiceEnUsHenryv3voiceConst     = "en-US_HenryV3Voice"
	SynthesizeOptionsVoiceEnUsKevinv3voiceConst     = "en-US_KevinV3Voice"
	SynthesizeOptionsVoiceEnUsLisav3voiceConst      = "en-US_LisaV3Voice"
	SynthesizeOptionsVoiceEnUsLisavoiceConst        = "en-US_LisaVoice"
	SynthesizeOptionsVoiceEnUsMichaelv3voiceConst   = "en-US_MichaelV3Voice"
	SynthesizeOptionsVoiceEnUsMichaelvoiceConst     = "en-US_MichaelVoice"
	SynthesizeOptionsVoiceEnUsOliviav3voiceConst    = "en-US_OliviaV3Voice"
	SynthesizeOptionsVoiceEsEsEnriquev3voiceConst   = "es-ES_EnriqueV3Voice"
	SynthesizeOptionsVoiceEsEsEnriquevoiceConst     = "es-ES_EnriqueVoice"
	SynthesizeOptionsVoiceEsEsLaurav3voiceConst     = "es-ES_LauraV3Voice"
	SynthesizeOptionsVoiceEsEsLauravoiceConst       = "es-ES_LauraVoice"
	SynthesizeOptionsVoiceEsLaSofiav3voiceConst     = "es-LA_SofiaV3Voice"
	SynthesizeOptionsVoiceEsLaSofiavoiceConst       = "es-LA_SofiaVoice"
	SynthesizeOptionsVoiceEsUsSofiav3voiceConst     = "es-US_SofiaV3Voice"
	SynthesizeOptionsVoiceEsUsSofiavoiceConst       = "es-US_SofiaVoice"
	SynthesizeOptionsVoiceFrFrNicolasv3voiceConst   = "fr-FR_NicolasV3Voice"
	SynthesizeOptionsVoiceFrFrReneev3voiceConst     = "fr-FR_ReneeV3Voice"
	SynthesizeOptionsVoiceFrFrReneevoiceConst       = "fr-FR_ReneeVoice"
	SynthesizeOptionsVoiceItItFrancescav3voiceConst = "it-IT_FrancescaV3Voice"
	SynthesizeOptionsVoiceItItFrancescavoiceConst   = "it-IT_FrancescaVoice"
	SynthesizeOptionsVoiceJaJpEmiv3voiceConst       = "ja-JP_EmiV3Voice"
	SynthesizeOptionsVoiceJaJpEmivoiceConst         = "ja-JP_EmiVoice"
	SynthesizeOptionsVoiceKoKrHyunjunvoiceConst     = "ko-KR_HyunjunVoice"
	SynthesizeOptionsVoiceKoKrSiwoovoiceConst       = "ko-KR_SiWooVoice"
	SynthesizeOptionsVoiceKoKrYoungmivoiceConst     = "ko-KR_YoungmiVoice"
	SynthesizeOptionsVoiceKoKrYunavoiceConst        = "ko-KR_YunaVoice"
	SynthesizeOptionsVoiceNlNlEmmavoiceConst        = "nl-NL_EmmaVoice"
	SynthesizeOptionsVoiceNlNlLiamvoiceConst        = "nl-NL_LiamVoice"
	SynthesizeOptionsVoicePtBrIsabelav3voiceConst   = "pt-BR_IsabelaV3Voice"
	SynthesizeOptionsVoicePtBrIsabelavoiceConst     = "pt-BR_IsabelaVoice"
	SynthesizeOptionsVoiceZhCnLinavoiceConst        = "zh-CN_LiNaVoice"
	SynthesizeOptionsVoiceZhCnWangweivoiceConst     = "zh-CN_WangWeiVoice"
	SynthesizeOptionsVoiceZhCnZhangjingvoiceConst   = "zh-CN_ZhangJingVoice"
)

// NewSynthesizeOptions : Instantiate SynthesizeOptions
func (*TextToSpeechV1) NewSynthesizeOptions(text string) *SynthesizeOptions {
	return &SynthesizeOptions{
		Text: core.StringPtr(text),
	}
}

// SetText : Allow user to set Text
func (options *SynthesizeOptions) SetText(text string) *SynthesizeOptions {
	options.Text = core.StringPtr(text)
	return options
}

// SetAccept : Allow user to set Accept
func (options *SynthesizeOptions) SetAccept(accept string) *SynthesizeOptions {
	options.Accept = core.StringPtr(accept)
	return options
}

// SetVoice : Allow user to set Voice
func (options *SynthesizeOptions) SetVoice(voice string) *SynthesizeOptions {
	options.Voice = core.StringPtr(voice)
	return options
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *SynthesizeOptions) SetCustomizationID(customizationID string) *SynthesizeOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *SynthesizeOptions) SetHeaders(param map[string]string) *SynthesizeOptions {
	options.Headers = param
	return options
}

// Translation : Information about the translation for the specified text.
type Translation struct {
	// The phonetic or sounds-like translation for the word. A phonetic translation is based on the SSML format for
	// representing the phonetic string of a word either as an IPA translation or as an IBM SPR translation. The Arabic,
	// Chinese, Dutch, Australian English, and Korean languages support only IPA. A sounds-like is one or more words that,
	// when combined, sound like the word.
	Translation *string `json:"translation" validate:"required"`

	// **Japanese only.** The part of speech for the word. The service uses the value to produce the correct intonation for
	// the word. You can create only a single entry, with or without a single part of speech, for any word; you cannot
	// create multiple entries with different parts of speech for the same word. For more information, see [Working with
	// Japanese entries](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-rules#jaNotes).
	PartOfSpeech *string `json:"part_of_speech,omitempty"`
}

// Constants associated with the Translation.PartOfSpeech property.
// **Japanese only.** The part of speech for the word. The service uses the value to produce the correct intonation for
// the word. You can create only a single entry, with or without a single part of speech, for any word; you cannot
// create multiple entries with different parts of speech for the same word. For more information, see [Working with
// Japanese entries](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-rules#jaNotes).
const (
	TranslationPartOfSpeechDosiConst = "Dosi"
	TranslationPartOfSpeechFukuConst = "Fuku"
	TranslationPartOfSpeechGobiConst = "Gobi"
	TranslationPartOfSpeechHokaConst = "Hoka"
	TranslationPartOfSpeechJodoConst = "Jodo"
	TranslationPartOfSpeechJosiConst = "Josi"
	TranslationPartOfSpeechKatoConst = "Kato"
	TranslationPartOfSpeechKedoConst = "Kedo"
	TranslationPartOfSpeechKeyoConst = "Keyo"
	TranslationPartOfSpeechKigoConst = "Kigo"
	TranslationPartOfSpeechKoyuConst = "Koyu"
	TranslationPartOfSpeechMesiConst = "Mesi"
	TranslationPartOfSpeechRetaConst = "Reta"
	TranslationPartOfSpeechStbiConst = "Stbi"
	TranslationPartOfSpeechSttoConst = "Stto"
	TranslationPartOfSpeechStzoConst = "Stzo"
	TranslationPartOfSpeechSujiConst = "Suji"
)

// NewTranslation : Instantiate Translation (Generic Model Constructor)
func (*TextToSpeechV1) NewTranslation(translation string) (model *Translation, err error) {
	model = &Translation{
		Translation: core.StringPtr(translation),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalTranslation unmarshals an instance of Translation from the specified map of raw messages.
func UnmarshalTranslation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Translation)
	err = core.UnmarshalPrimitive(m, "translation", &obj.Translation)
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

// UpdateCustomModelOptions : The UpdateCustomModel options.
type UpdateCustomModelOptions struct {
	// The customization ID (GUID) of the custom model. You must make the request with credentials for the instance of the
	// service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// A new name for the custom model.
	Name *string `json:"name,omitempty"`

	// A new description for the custom model.
	Description *string `json:"description,omitempty"`

	// An array of `Word` objects that provides the words and their translations that are to be added or updated for the
	// custom model. Pass an empty array to make no additions or updates.
	Words []Word `json:"words,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateCustomModelOptions : Instantiate UpdateCustomModelOptions
func (*TextToSpeechV1) NewUpdateCustomModelOptions(customizationID string) *UpdateCustomModelOptions {
	return &UpdateCustomModelOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *UpdateCustomModelOptions) SetCustomizationID(customizationID string) *UpdateCustomModelOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetName : Allow user to set Name
func (options *UpdateCustomModelOptions) SetName(name string) *UpdateCustomModelOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetDescription : Allow user to set Description
func (options *UpdateCustomModelOptions) SetDescription(description string) *UpdateCustomModelOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetWords : Allow user to set Words
func (options *UpdateCustomModelOptions) SetWords(words []Word) *UpdateCustomModelOptions {
	options.Words = words
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateCustomModelOptions) SetHeaders(param map[string]string) *UpdateCustomModelOptions {
	options.Headers = param
	return options
}

// Voice : Information about an available voice.
type Voice struct {
	// The URI of the voice.
	URL *string `json:"url" validate:"required"`

	// The gender of the voice: `male` or `female`.
	Gender *string `json:"gender" validate:"required"`

	// The name of the voice. Use this as the voice identifier in all requests.
	Name *string `json:"name" validate:"required"`

	// The language and region of the voice (for example, `en-US`).
	Language *string `json:"language" validate:"required"`

	// A textual description of the voice.
	Description *string `json:"description" validate:"required"`

	// If `true`, the voice can be customized; if `false`, the voice cannot be customized. (Same as `custom_pronunciation`;
	// maintained for backward compatibility.).
	Customizable *bool `json:"customizable" validate:"required"`

	// Additional service features that are supported with the voice.
	SupportedFeatures *SupportedFeatures `json:"supported_features" validate:"required"`

	// Returns information about a specified custom model. This field is returned only by the **Get a voice** method and
	// only when you specify the customization ID of a custom model.
	Customization *CustomModel `json:"customization,omitempty"`
}

// UnmarshalVoice unmarshals an instance of Voice from the specified map of raw messages.
func UnmarshalVoice(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Voice)
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "gender", &obj.Gender)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
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
	err = core.UnmarshalPrimitive(m, "customizable", &obj.Customizable)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "supported_features", &obj.SupportedFeatures, UnmarshalSupportedFeatures)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "customization", &obj.Customization, UnmarshalCustomModel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Voices : Information about all available voices.
type Voices struct {
	// A list of available voices.
	Voices []Voice `json:"voices" validate:"required"`
}

// UnmarshalVoices unmarshals an instance of Voices from the specified map of raw messages.
func UnmarshalVoices(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Voices)
	err = core.UnmarshalModel(m, "voices", &obj.Voices, UnmarshalVoice)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Word : Information about a word for the custom model.
type Word struct {
	// The word for the custom model. The maximum length of a word is 49 characters.
	Word *string `json:"word" validate:"required"`

	// The phonetic or sounds-like translation for the word. A phonetic translation is based on the SSML format for
	// representing the phonetic string of a word either as an IPA or IBM SPR translation. The Arabic, Chinese, Dutch,
	// Australian English, and Korean languages support only IPA. A sounds-like translation consists of one or more words
	// that, when combined, sound like the word. The maximum length of a translation is 499 characters.
	Translation *string `json:"translation" validate:"required"`

	// **Japanese only.** The part of speech for the word. The service uses the value to produce the correct intonation for
	// the word. You can create only a single entry, with or without a single part of speech, for any word; you cannot
	// create multiple entries with different parts of speech for the same word. For more information, see [Working with
	// Japanese entries](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-rules#jaNotes).
	PartOfSpeech *string `json:"part_of_speech,omitempty"`
}

// Constants associated with the Word.PartOfSpeech property.
// **Japanese only.** The part of speech for the word. The service uses the value to produce the correct intonation for
// the word. You can create only a single entry, with or without a single part of speech, for any word; you cannot
// create multiple entries with different parts of speech for the same word. For more information, see [Working with
// Japanese entries](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-rules#jaNotes).
const (
	WordPartOfSpeechDosiConst = "Dosi"
	WordPartOfSpeechFukuConst = "Fuku"
	WordPartOfSpeechGobiConst = "Gobi"
	WordPartOfSpeechHokaConst = "Hoka"
	WordPartOfSpeechJodoConst = "Jodo"
	WordPartOfSpeechJosiConst = "Josi"
	WordPartOfSpeechKatoConst = "Kato"
	WordPartOfSpeechKedoConst = "Kedo"
	WordPartOfSpeechKeyoConst = "Keyo"
	WordPartOfSpeechKigoConst = "Kigo"
	WordPartOfSpeechKoyuConst = "Koyu"
	WordPartOfSpeechMesiConst = "Mesi"
	WordPartOfSpeechRetaConst = "Reta"
	WordPartOfSpeechStbiConst = "Stbi"
	WordPartOfSpeechSttoConst = "Stto"
	WordPartOfSpeechStzoConst = "Stzo"
	WordPartOfSpeechSujiConst = "Suji"
)

// NewWord : Instantiate Word (Generic Model Constructor)
func (*TextToSpeechV1) NewWord(word string, translation string) (model *Word, err error) {
	model = &Word{
		Word:        core.StringPtr(word),
		Translation: core.StringPtr(translation),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalWord unmarshals an instance of Word from the specified map of raw messages.
func UnmarshalWord(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Word)
	err = core.UnmarshalPrimitive(m, "word", &obj.Word)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "translation", &obj.Translation)
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

// Words : For the **Add custom words** method, one or more words that are to be added or updated for the custom model and the
// translation for each specified word.
//
// For the **List custom words** method, the words and their translations from the custom model.
type Words struct {
	// The **Add custom words** method accepts an array of `Word` objects. Each object provides a word that is to be added
	// or updated for the custom model and the word's translation.
	//
	// The **List custom words** method returns an array of `Word` objects. Each object shows a word and its translation
	// from the custom model. The words are listed in alphabetical order, with uppercase letters listed before lowercase
	// letters. The array is empty if the custom model contains no words.
	Words []Word `json:"words" validate:"required"`
}

// NewWords : Instantiate Words (Generic Model Constructor)
func (*TextToSpeechV1) NewWords(words []Word) (model *Words, err error) {
	model = &Words{
		Words: words,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalWords unmarshals an instance of Words from the specified map of raw messages.
func UnmarshalWords(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Words)
	err = core.UnmarshalModel(m, "words", &obj.Words, UnmarshalWord)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
