/**
 * (C) Copyright IBM Corp. 2022.
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
 * IBM OpenAPI SDK Code Generator Version: 3.46.0-a4e29da0-20220224-210428
 */

// Package speechtotextv1 : Operations and models for the SpeechToTextV1 service
package speechtotextv1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/watson-developer-cloud/go-sdk/v3/common"
)

// SpeechToTextV1 : The IBM Watson&trade; Speech to Text service provides APIs that use IBM's speech-recognition
// capabilities to produce transcripts of spoken audio.  The service can transcribe speech from various languages and
// audio formats. In addition to basic transcription, the service can produce detailed information about many different
// aspects of the audio. It returns all JSON response content in the UTF-8 character set.
//
// The service supports two types of models: previous-generation models that include the terms `Broadband` and
// `Narrowband` in their names, and next-generation models that include the terms `Multimedia` and `Telephony` in their
// names. Broadband and multimedia models have minimum sampling rates of 16 kHz. Narrowband and telephony models have
// minimum sampling rates of 8 kHz. The next-generation models offer high throughput and greater transcription accuracy.
//
// Effective 15 March 2022, previous-generation models for all languages other than Arabic and Japanese are deprecated.
// The deprecated models remain available until 15 September 2022, when they will be removed from the service and the
// documentation. You must migrate to the equivalent next-generation model by the end of service date. For more
// information, see [Migrating to next-generation
// models](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-models-migrate).{: deprecated}
//
// For speech recognition, the service supports synchronous and asynchronous HTTP Representational State Transfer (REST)
// interfaces. It also supports a WebSocket interface that provides a full-duplex, low-latency communication channel:
// Clients send requests and audio to the service and receive results over a single connection asynchronously.
//
// The service also offers two customization interfaces. Use language model customization to expand the vocabulary of a
// base model with domain-specific terminology. Use acoustic model customization to adapt a base model for the acoustic
// characteristics of your audio. For language model customization, the service also supports grammars. A grammar is a
// formal language specification that lets you restrict the phrases that the service can recognize.
//
// Language model customization and grammars are available for most previous- and next-generation models. Acoustic model
// customization is available for all previous-generation models.
//
// API Version: 1.0.0
// See: https://cloud.ibm.com/docs/speech-to-text
type SpeechToTextV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.us-south.speech-to-text.watson.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "speech_to_text"

// SpeechToTextV1Options : Service options
type SpeechToTextV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewSpeechToTextV1 : constructs an instance of SpeechToTextV1 with passed in options.
func NewSpeechToTextV1(options *SpeechToTextV1Options) (service *SpeechToTextV1, err error) {
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

	service = &SpeechToTextV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "speechToText" suitable for processing requests.
func (speechToText *SpeechToTextV1) Clone() *SpeechToTextV1 {
	if core.IsNil(speechToText) {
		return nil
	}
	clone := *speechToText
	clone.Service = speechToText.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (speechToText *SpeechToTextV1) SetServiceURL(url string) error {
	return speechToText.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (speechToText *SpeechToTextV1) GetServiceURL() string {
	return speechToText.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (speechToText *SpeechToTextV1) SetDefaultHeaders(headers http.Header) {
	speechToText.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (speechToText *SpeechToTextV1) SetEnableGzipCompression(enableGzip bool) {
	speechToText.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (speechToText *SpeechToTextV1) GetEnableGzipCompression() bool {
	return speechToText.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (speechToText *SpeechToTextV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	speechToText.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (speechToText *SpeechToTextV1) DisableRetries() {
	speechToText.Service.DisableRetries()
}

// DisableSSLVerification bypasses verification of the server's SSL certificate
func (speechToText *SpeechToTextV1) DisableSSLVerification() {
	speechToText.Service.DisableSSLVerification()
}

// ListModels : List models
// Lists all language models that are available for use with the service. The information includes the name of the model
// and its minimum sampling rate in Hertz, among other things. The ordering of the list of models can change from call
// to call; do not rely on an alphabetized or static list of models.
//
// **See also:** [Listing all
// models](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-models-list#models-list-all).
func (speechToText *SpeechToTextV1) ListModels(listModelsOptions *ListModelsOptions) (result *SpeechModels, response *core.DetailedResponse, err error) {
	return speechToText.ListModelsWithContext(context.Background(), listModelsOptions)
}

// ListModelsWithContext is an alternate form of the ListModels method which supports a Context parameter
func (speechToText *SpeechToTextV1) ListModelsWithContext(ctx context.Context, listModelsOptions *ListModelsOptions) (result *SpeechModels, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listModelsOptions, "listModelsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/models`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listModelsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "ListModels")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = speechToText.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSpeechModels)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetModel : Get a model
// Gets information for a single specified language model that is available for use with the service. The information
// includes the name of the model and its minimum sampling rate in Hertz, among other things.
//
// **See also:** [Listing a specific
// model](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-models-list#models-list-specific).
func (speechToText *SpeechToTextV1) GetModel(getModelOptions *GetModelOptions) (result *SpeechModel, response *core.DetailedResponse, err error) {
	return speechToText.GetModelWithContext(context.Background(), getModelOptions)
}

// GetModelWithContext is an alternate form of the GetModel method which supports a Context parameter
func (speechToText *SpeechToTextV1) GetModelWithContext(ctx context.Context, getModelOptions *GetModelOptions) (result *SpeechModel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getModelOptions, "getModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getModelOptions, "getModelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"model_id": *getModelOptions.ModelID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/models/{model_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "GetModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = speechToText.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSpeechModel)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// Recognize : Recognize audio
// Sends audio and returns transcription results for a recognition request. You can pass a maximum of 100 MB and a
// minimum of 100 bytes of audio with a request. The service automatically detects the endianness of the incoming audio
// and, for audio that includes multiple channels, downmixes the audio to one-channel mono during transcoding. The
// method returns only final results; to enable interim results, use the WebSocket API. (With the `curl` command, use
// the `--data-binary` option to upload the file for the request.)
//
// **See also:** [Making a basic HTTP
// request](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-http#HTTP-basic).
//
// ### Streaming mode
//
//	For requests to transcribe live audio as it becomes available, you must set the `Transfer-Encoding` header to
//
// `chunked` to use streaming mode. In streaming mode, the service closes the connection (status code 408) if it does
// not receive at least 15 seconds of audio (including silence) in any 30-second period. The service also closes the
// connection (status code 400) if it detects no speech for `inactivity_timeout` seconds of streaming audio; use the
// `inactivity_timeout` parameter to change the default of 30 seconds.
//
// **See also:**
// * [Audio transmission](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-input#transmission)
// * [Timeouts](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-input#timeouts)
//
// ### Audio formats (content types)
//
//	The service accepts audio in the following formats (MIME types).
//
// * For formats that are labeled **Required**, you must use the `Content-Type` header with the request to specify the
// format of the audio.
// * For all other formats, you can omit the `Content-Type` header or specify `application/octet-stream` with the header
// to have the service automatically detect the format of the audio. (With the `curl` command, you can specify either
// `"Content-Type:"` or `"Content-Type: application/octet-stream"`.)
//
// Where indicated, the format that you specify must include the sampling rate and can optionally include the number of
// channels and the endianness of the audio.
// * `audio/alaw` (**Required.** Specify the sampling rate (`rate`) of the audio.)
// * `audio/basic` (**Required.** Use only with narrowband models.)
// * `audio/flac`
// * `audio/g729` (Use only with narrowband models.)
// * `audio/l16` (**Required.** Specify the sampling rate (`rate`) and optionally the number of channels (`channels`)
// and endianness (`endianness`) of the audio.)
// * `audio/mp3`
// * `audio/mpeg`
// * `audio/mulaw` (**Required.** Specify the sampling rate (`rate`) of the audio.)
// * `audio/ogg` (The service automatically detects the codec of the input audio.)
// * `audio/ogg;codecs=opus`
// * `audio/ogg;codecs=vorbis`
// * `audio/wav` (Provide audio with a maximum of nine channels.)
// * `audio/webm` (The service automatically detects the codec of the input audio.)
// * `audio/webm;codecs=opus`
// * `audio/webm;codecs=vorbis`
//
// The sampling rate of the audio must match the sampling rate of the model for the recognition request: for broadband
// models, at least 16 kHz; for narrowband models, at least 8 kHz. If the sampling rate of the audio is higher than the
// minimum required rate, the service down-samples the audio to the appropriate rate. If the sampling rate of the audio
// is lower than the minimum required rate, the request fails.
//
//	**See also:** [Supported audio
//
// formats](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-audio-formats).
//
// ### Next-generation models
//
//	The service supports next-generation `Multimedia` (16 kHz) and `Telephony` (8 kHz) models for many languages.
//
// Next-generation models have higher throughput than the service's previous generation of `Broadband` and `Narrowband`
// models. When you use next-generation models, the service can return transcriptions more quickly and also provide
// noticeably better transcription accuracy.
//
// You specify a next-generation model by using the `model` query parameter, as you do a previous-generation model. Many
// next-generation models also support the `low_latency` parameter, which is not available with previous-generation
// models. Next-generation models do not support all of the parameters that are available for use with
// previous-generation models.
//
// **Important:** Effective 15 March 2022, previous-generation models for all languages other than Arabic and Japanese
// are deprecated. The deprecated models remain available until 15 September 2022, when they will be removed from the
// service and the documentation. You must migrate to the equivalent next-generation model by the end of service date.
// For more information, see [Migrating to next-generation
// models](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-models-migrate).
//
// **See also:**
// * [Next-generation languages and models](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-models-ng)
// * [Supported features for next-generation
// models](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-models-ng#models-ng-features)
//
// ### Multipart speech recognition
//
//	**Note:** The asynchronous HTTP interface, WebSocket interface, and Watson SDKs do not support multipart speech
//
// recognition.
//
// The HTTP `POST` method of the service also supports multipart speech recognition. With multipart requests, you pass
// all audio data as multipart form data. You specify some parameters as request headers and query parameters, but you
// pass JSON metadata as form data to control most aspects of the transcription. You can use multipart recognition to
// pass multiple audio files with a single request.
//
// Use the multipart approach with browsers for which JavaScript is disabled or when the parameters used with the
// request are greater than the 8 KB limit imposed by most HTTP servers and proxies. You can encounter this limit, for
// example, if you want to spot a very large number of keywords.
//
// **See also:** [Making a multipart HTTP
// request](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-http#HTTP-multi).
func (speechToText *SpeechToTextV1) Recognize(recognizeOptions *RecognizeOptions) (result *SpeechRecognitionResults, response *core.DetailedResponse, err error) {
	return speechToText.RecognizeWithContext(context.Background(), recognizeOptions)
}

// RecognizeWithContext is an alternate form of the Recognize method which supports a Context parameter
func (speechToText *SpeechToTextV1) RecognizeWithContext(ctx context.Context, recognizeOptions *RecognizeOptions) (result *SpeechRecognitionResults, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(recognizeOptions, "recognizeOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(recognizeOptions, "recognizeOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/recognize`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range recognizeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "Recognize")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if recognizeOptions.ContentType != nil {
		builder.AddHeader("Content-Type", fmt.Sprint(*recognizeOptions.ContentType))
	}

	if recognizeOptions.Model != nil {
		builder.AddQuery("model", fmt.Sprint(*recognizeOptions.Model))
	}
	if recognizeOptions.LanguageCustomizationID != nil {
		builder.AddQuery("language_customization_id", fmt.Sprint(*recognizeOptions.LanguageCustomizationID))
	}
	if recognizeOptions.AcousticCustomizationID != nil {
		builder.AddQuery("acoustic_customization_id", fmt.Sprint(*recognizeOptions.AcousticCustomizationID))
	}
	if recognizeOptions.BaseModelVersion != nil {
		builder.AddQuery("base_model_version", fmt.Sprint(*recognizeOptions.BaseModelVersion))
	}
	if recognizeOptions.CustomizationWeight != nil {
		builder.AddQuery("customization_weight", fmt.Sprint(*recognizeOptions.CustomizationWeight))
	}
	if recognizeOptions.InactivityTimeout != nil {
		builder.AddQuery("inactivity_timeout", fmt.Sprint(*recognizeOptions.InactivityTimeout))
	}
	if recognizeOptions.Keywords != nil {
		builder.AddQuery("keywords", strings.Join(recognizeOptions.Keywords, ","))
	}
	if recognizeOptions.KeywordsThreshold != nil {
		builder.AddQuery("keywords_threshold", fmt.Sprint(*recognizeOptions.KeywordsThreshold))
	}
	if recognizeOptions.MaxAlternatives != nil {
		builder.AddQuery("max_alternatives", fmt.Sprint(*recognizeOptions.MaxAlternatives))
	}
	if recognizeOptions.WordAlternativesThreshold != nil {
		builder.AddQuery("word_alternatives_threshold", fmt.Sprint(*recognizeOptions.WordAlternativesThreshold))
	}
	if recognizeOptions.WordConfidence != nil {
		builder.AddQuery("word_confidence", fmt.Sprint(*recognizeOptions.WordConfidence))
	}
	if recognizeOptions.Timestamps != nil {
		builder.AddQuery("timestamps", fmt.Sprint(*recognizeOptions.Timestamps))
	}
	if recognizeOptions.ProfanityFilter != nil {
		builder.AddQuery("profanity_filter", fmt.Sprint(*recognizeOptions.ProfanityFilter))
	}
	if recognizeOptions.SmartFormatting != nil {
		builder.AddQuery("smart_formatting", fmt.Sprint(*recognizeOptions.SmartFormatting))
	}
	if recognizeOptions.SpeakerLabels != nil {
		builder.AddQuery("speaker_labels", fmt.Sprint(*recognizeOptions.SpeakerLabels))
	}
	if recognizeOptions.CustomizationID != nil {
		builder.AddQuery("customization_id", fmt.Sprint(*recognizeOptions.CustomizationID))
	}
	if recognizeOptions.GrammarName != nil {
		builder.AddQuery("grammar_name", fmt.Sprint(*recognizeOptions.GrammarName))
	}
	if recognizeOptions.Redaction != nil {
		builder.AddQuery("redaction", fmt.Sprint(*recognizeOptions.Redaction))
	}
	if recognizeOptions.AudioMetrics != nil {
		builder.AddQuery("audio_metrics", fmt.Sprint(*recognizeOptions.AudioMetrics))
	}
	if recognizeOptions.EndOfPhraseSilenceTime != nil {
		builder.AddQuery("end_of_phrase_silence_time", fmt.Sprint(*recognizeOptions.EndOfPhraseSilenceTime))
	}
	if recognizeOptions.SplitTranscriptAtPhraseEnd != nil {
		builder.AddQuery("split_transcript_at_phrase_end", fmt.Sprint(*recognizeOptions.SplitTranscriptAtPhraseEnd))
	}
	if recognizeOptions.SpeechDetectorSensitivity != nil {
		builder.AddQuery("speech_detector_sensitivity", fmt.Sprint(*recognizeOptions.SpeechDetectorSensitivity))
	}
	if recognizeOptions.BackgroundAudioSuppression != nil {
		builder.AddQuery("background_audio_suppression", fmt.Sprint(*recognizeOptions.BackgroundAudioSuppression))
	}
	if recognizeOptions.CharacterInsertionBias != nil {
		builder.AddQuery("character_insertion_bias", fmt.Sprint(*recognizeOptions.CharacterInsertionBias))
	}
	if recognizeOptions.LowLatency != nil {
		builder.AddQuery("low_latency", fmt.Sprint(*recognizeOptions.LowLatency))
	}

	_, err = builder.SetBodyContent(core.StringNilMapper(recognizeOptions.ContentType), nil, nil, recognizeOptions.Audio)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = speechToText.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSpeechRecognitionResults)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// RegisterCallback : Register a callback
// Registers a callback URL with the service for use with subsequent asynchronous recognition requests. The service
// attempts to register, or allowlist, the callback URL if it is not already registered by sending a `GET` request to
// the callback URL. The service passes a random alphanumeric challenge string via the `challenge_string` parameter of
// the request. The request includes an `Accept` header that specifies `text/plain` as the required response type.
//
// To be registered successfully, the callback URL must respond to the `GET` request from the service. The response must
// send status code 200 and must include the challenge string in its body. Set the `Content-Type` response header to
// `text/plain`. Upon receiving this response, the service responds to the original registration request with response
// code 201.
//
// The service sends only a single `GET` request to the callback URL. If the service does not receive a reply with a
// response code of 200 and a body that echoes the challenge string sent by the service within five seconds, it does not
// allowlist the URL; it instead sends status code 400 in response to the request to register a callback. If the
// requested callback URL is already allowlisted, the service responds to the initial registration request with response
// code 200.
//
// If you specify a user secret with the request, the service uses it as a key to calculate an HMAC-SHA1 signature of
// the challenge string in its response to the `POST` request. It sends this signature in the `X-Callback-Signature`
// header of its `GET` request to the URL during registration. It also uses the secret to calculate a signature over the
// payload of every callback notification that uses the URL. The signature provides authentication and data integrity
// for HTTP communications.
//
// After you successfully register a callback URL, you can use it with an indefinite number of recognition requests. You
// can register a maximum of 20 callback URLS in a one-hour span of time.
//
// **See also:** [Registering a callback
// URL](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-async#register).
func (speechToText *SpeechToTextV1) RegisterCallback(registerCallbackOptions *RegisterCallbackOptions) (result *RegisterStatus, response *core.DetailedResponse, err error) {
	return speechToText.RegisterCallbackWithContext(context.Background(), registerCallbackOptions)
}

// RegisterCallbackWithContext is an alternate form of the RegisterCallback method which supports a Context parameter
func (speechToText *SpeechToTextV1) RegisterCallbackWithContext(ctx context.Context, registerCallbackOptions *RegisterCallbackOptions) (result *RegisterStatus, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(registerCallbackOptions, "registerCallbackOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(registerCallbackOptions, "registerCallbackOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/register_callback`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range registerCallbackOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "RegisterCallback")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("callback_url", fmt.Sprint(*registerCallbackOptions.CallbackURL))
	if registerCallbackOptions.UserSecret != nil {
		builder.AddQuery("user_secret", fmt.Sprint(*registerCallbackOptions.UserSecret))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = speechToText.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRegisterStatus)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UnregisterCallback : Unregister a callback
// Unregisters a callback URL that was previously allowlisted with a [Register a callback](#registercallback) request
// for use with the asynchronous interface. Once unregistered, the URL can no longer be used with asynchronous
// recognition requests.
//
// **See also:** [Unregistering a callback
// URL](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-async#unregister).
func (speechToText *SpeechToTextV1) UnregisterCallback(unregisterCallbackOptions *UnregisterCallbackOptions) (response *core.DetailedResponse, err error) {
	return speechToText.UnregisterCallbackWithContext(context.Background(), unregisterCallbackOptions)
}

// UnregisterCallbackWithContext is an alternate form of the UnregisterCallback method which supports a Context parameter
func (speechToText *SpeechToTextV1) UnregisterCallbackWithContext(ctx context.Context, unregisterCallbackOptions *UnregisterCallbackOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(unregisterCallbackOptions, "unregisterCallbackOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(unregisterCallbackOptions, "unregisterCallbackOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/unregister_callback`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range unregisterCallbackOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "UnregisterCallback")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("callback_url", fmt.Sprint(*unregisterCallbackOptions.CallbackURL))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = speechToText.Service.Request(request, nil)

	return
}

// CreateJob : Create a job
// Creates a job for a new asynchronous recognition request. The job is owned by the instance of the service whose
// credentials are used to create it. How you learn the status and results of a job depends on the parameters you
// include with the job creation request:
// * By callback notification: Include the `callback_url` parameter to specify a URL to which the service is to send
// callback notifications when the status of the job changes. Optionally, you can also include the `events` and
// `user_token` parameters to subscribe to specific events and to specify a string that is to be included with each
// notification for the job.
// * By polling the service: Omit the `callback_url`, `events`, and `user_token` parameters. You must then use the
// [Check jobs](#checkjobs) or [Check a job](#checkjob) methods to check the status of the job, using the latter to
// retrieve the results when the job is complete.
//
// The two approaches are not mutually exclusive. You can poll the service for job status or obtain results from the
// service manually even if you include a callback URL. In both cases, you can include the `results_ttl` parameter to
// specify how long the results are to remain available after the job is complete. Using the HTTPS [Check a
// job](#checkjob) method to retrieve results is more secure than receiving them via callback notification over HTTP
// because it provides confidentiality in addition to authentication and data integrity.
//
// The method supports the same basic parameters as other HTTP and WebSocket recognition requests. It also supports the
// following parameters specific to the asynchronous interface:
// * `callback_url`
// * `events`
// * `user_token`
// * `results_ttl`
//
// You can pass a maximum of 1 GB and a minimum of 100 bytes of audio with a request. The service automatically detects
// the endianness of the incoming audio and, for audio that includes multiple channels, downmixes the audio to
// one-channel mono during transcoding. The method returns only final results; to enable interim results, use the
// WebSocket API. (With the `curl` command, use the `--data-binary` option to upload the file for the request.)
//
// **See also:** [Creating a job](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-async#create).
//
// ### Streaming mode
//
//	For requests to transcribe live audio as it becomes available, you must set the `Transfer-Encoding` header to
//
// `chunked` to use streaming mode. In streaming mode, the service closes the connection (status code 408) if it does
// not receive at least 15 seconds of audio (including silence) in any 30-second period. The service also closes the
// connection (status code 400) if it detects no speech for `inactivity_timeout` seconds of streaming audio; use the
// `inactivity_timeout` parameter to change the default of 30 seconds.
//
// **See also:**
// * [Audio transmission](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-input#transmission)
// * [Timeouts](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-input#timeouts)
//
// ### Audio formats (content types)
//
//	The service accepts audio in the following formats (MIME types).
//
// * For formats that are labeled **Required**, you must use the `Content-Type` header with the request to specify the
// format of the audio.
// * For all other formats, you can omit the `Content-Type` header or specify `application/octet-stream` with the header
// to have the service automatically detect the format of the audio. (With the `curl` command, you can specify either
// `"Content-Type:"` or `"Content-Type: application/octet-stream"`.)
//
// Where indicated, the format that you specify must include the sampling rate and can optionally include the number of
// channels and the endianness of the audio.
// * `audio/alaw` (**Required.** Specify the sampling rate (`rate`) of the audio.)
// * `audio/basic` (**Required.** Use only with narrowband models.)
// * `audio/flac`
// * `audio/g729` (Use only with narrowband models.)
// * `audio/l16` (**Required.** Specify the sampling rate (`rate`) and optionally the number of channels (`channels`)
// and endianness (`endianness`) of the audio.)
// * `audio/mp3`
// * `audio/mpeg`
// * `audio/mulaw` (**Required.** Specify the sampling rate (`rate`) of the audio.)
// * `audio/ogg` (The service automatically detects the codec of the input audio.)
// * `audio/ogg;codecs=opus`
// * `audio/ogg;codecs=vorbis`
// * `audio/wav` (Provide audio with a maximum of nine channels.)
// * `audio/webm` (The service automatically detects the codec of the input audio.)
// * `audio/webm;codecs=opus`
// * `audio/webm;codecs=vorbis`
//
// The sampling rate of the audio must match the sampling rate of the model for the recognition request: for broadband
// models, at least 16 kHz; for narrowband models, at least 8 kHz. If the sampling rate of the audio is higher than the
// minimum required rate, the service down-samples the audio to the appropriate rate. If the sampling rate of the audio
// is lower than the minimum required rate, the request fails.
//
//	**See also:** [Supported audio
//
// formats](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-audio-formats).
//
// ### Next-generation models
//
//	The service supports next-generation `Multimedia` (16 kHz) and `Telephony` (8 kHz) models for many languages.
//
// Next-generation models have higher throughput than the service's previous generation of `Broadband` and `Narrowband`
// models. When you use next-generation models, the service can return transcriptions more quickly and also provide
// noticeably better transcription accuracy.
//
// You specify a next-generation model by using the `model` query parameter, as you do a previous-generation model. Many
// next-generation models also support the `low_latency` parameter, which is not available with previous-generation
// models. Next-generation models do not support all of the parameters that are available for use with
// previous-generation models.
//
// **Important:** Effective 15 March 2022, previous-generation models for all languages other than Arabic and Japanese
// are deprecated. The deprecated models remain available until 15 September 2022, when they will be removed from the
// service and the documentation. You must migrate to the equivalent next-generation model by the end of service date.
// For more information, see  [Migrating to next-generation
// models](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-models-migrate).
//
// **See also:**
// * [Next-generation languages and models](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-models-ng)
// * [Supported features for next-generation
// models](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-models-ng#models-ng-features).
func (speechToText *SpeechToTextV1) CreateJob(createJobOptions *CreateJobOptions) (result *RecognitionJob, response *core.DetailedResponse, err error) {
	return speechToText.CreateJobWithContext(context.Background(), createJobOptions)
}

// CreateJobWithContext is an alternate form of the CreateJob method which supports a Context parameter
func (speechToText *SpeechToTextV1) CreateJobWithContext(ctx context.Context, createJobOptions *CreateJobOptions) (result *RecognitionJob, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createJobOptions, "createJobOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createJobOptions, "createJobOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/recognitions`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createJobOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "CreateJob")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if createJobOptions.ContentType != nil {
		builder.AddHeader("Content-Type", fmt.Sprint(*createJobOptions.ContentType))
	}

	if createJobOptions.Model != nil {
		builder.AddQuery("model", fmt.Sprint(*createJobOptions.Model))
	}
	if createJobOptions.CallbackURL != nil {
		builder.AddQuery("callback_url", fmt.Sprint(*createJobOptions.CallbackURL))
	}
	if createJobOptions.Events != nil {
		builder.AddQuery("events", fmt.Sprint(*createJobOptions.Events))
	}
	if createJobOptions.UserToken != nil {
		builder.AddQuery("user_token", fmt.Sprint(*createJobOptions.UserToken))
	}
	if createJobOptions.ResultsTTL != nil {
		builder.AddQuery("results_ttl", fmt.Sprint(*createJobOptions.ResultsTTL))
	}
	if createJobOptions.LanguageCustomizationID != nil {
		builder.AddQuery("language_customization_id", fmt.Sprint(*createJobOptions.LanguageCustomizationID))
	}
	if createJobOptions.AcousticCustomizationID != nil {
		builder.AddQuery("acoustic_customization_id", fmt.Sprint(*createJobOptions.AcousticCustomizationID))
	}
	if createJobOptions.BaseModelVersion != nil {
		builder.AddQuery("base_model_version", fmt.Sprint(*createJobOptions.BaseModelVersion))
	}
	if createJobOptions.CustomizationWeight != nil {
		builder.AddQuery("customization_weight", fmt.Sprint(*createJobOptions.CustomizationWeight))
	}
	if createJobOptions.InactivityTimeout != nil {
		builder.AddQuery("inactivity_timeout", fmt.Sprint(*createJobOptions.InactivityTimeout))
	}
	if createJobOptions.Keywords != nil {
		builder.AddQuery("keywords", strings.Join(createJobOptions.Keywords, ","))
	}
	if createJobOptions.KeywordsThreshold != nil {
		builder.AddQuery("keywords_threshold", fmt.Sprint(*createJobOptions.KeywordsThreshold))
	}
	if createJobOptions.MaxAlternatives != nil {
		builder.AddQuery("max_alternatives", fmt.Sprint(*createJobOptions.MaxAlternatives))
	}
	if createJobOptions.WordAlternativesThreshold != nil {
		builder.AddQuery("word_alternatives_threshold", fmt.Sprint(*createJobOptions.WordAlternativesThreshold))
	}
	if createJobOptions.WordConfidence != nil {
		builder.AddQuery("word_confidence", fmt.Sprint(*createJobOptions.WordConfidence))
	}
	if createJobOptions.Timestamps != nil {
		builder.AddQuery("timestamps", fmt.Sprint(*createJobOptions.Timestamps))
	}
	if createJobOptions.ProfanityFilter != nil {
		builder.AddQuery("profanity_filter", fmt.Sprint(*createJobOptions.ProfanityFilter))
	}
	if createJobOptions.SmartFormatting != nil {
		builder.AddQuery("smart_formatting", fmt.Sprint(*createJobOptions.SmartFormatting))
	}
	if createJobOptions.SpeakerLabels != nil {
		builder.AddQuery("speaker_labels", fmt.Sprint(*createJobOptions.SpeakerLabels))
	}
	if createJobOptions.CustomizationID != nil {
		builder.AddQuery("customization_id", fmt.Sprint(*createJobOptions.CustomizationID))
	}
	if createJobOptions.GrammarName != nil {
		builder.AddQuery("grammar_name", fmt.Sprint(*createJobOptions.GrammarName))
	}
	if createJobOptions.Redaction != nil {
		builder.AddQuery("redaction", fmt.Sprint(*createJobOptions.Redaction))
	}
	if createJobOptions.ProcessingMetrics != nil {
		builder.AddQuery("processing_metrics", fmt.Sprint(*createJobOptions.ProcessingMetrics))
	}
	if createJobOptions.ProcessingMetricsInterval != nil {
		builder.AddQuery("processing_metrics_interval", fmt.Sprint(*createJobOptions.ProcessingMetricsInterval))
	}
	if createJobOptions.AudioMetrics != nil {
		builder.AddQuery("audio_metrics", fmt.Sprint(*createJobOptions.AudioMetrics))
	}
	if createJobOptions.EndOfPhraseSilenceTime != nil {
		builder.AddQuery("end_of_phrase_silence_time", fmt.Sprint(*createJobOptions.EndOfPhraseSilenceTime))
	}
	if createJobOptions.SplitTranscriptAtPhraseEnd != nil {
		builder.AddQuery("split_transcript_at_phrase_end", fmt.Sprint(*createJobOptions.SplitTranscriptAtPhraseEnd))
	}
	if createJobOptions.SpeechDetectorSensitivity != nil {
		builder.AddQuery("speech_detector_sensitivity", fmt.Sprint(*createJobOptions.SpeechDetectorSensitivity))
	}
	if createJobOptions.BackgroundAudioSuppression != nil {
		builder.AddQuery("background_audio_suppression", fmt.Sprint(*createJobOptions.BackgroundAudioSuppression))
	}
	if createJobOptions.LowLatency != nil {
		builder.AddQuery("low_latency", fmt.Sprint(*createJobOptions.LowLatency))
	}

	_, err = builder.SetBodyContent(core.StringNilMapper(createJobOptions.ContentType), nil, nil, createJobOptions.Audio)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = speechToText.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRecognitionJob)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CheckJobs : Check jobs
// Returns the ID and status of the latest 100 outstanding jobs associated with the credentials with which it is called.
// The method also returns the creation and update times of each job, and, if a job was created with a callback URL and
// a user token, the user token for the job. To obtain the results for a job whose status is `completed` or not one of
// the latest 100 outstanding jobs, use the [Check a job[(#checkjob) method. A job and its results remain available
// until you delete them with the [Delete a job](#deletejob) method or until the job's time to live expires, whichever
// comes first.
//
// **See also:** [Checking the status of the latest
// jobs](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-async#jobs).
func (speechToText *SpeechToTextV1) CheckJobs(checkJobsOptions *CheckJobsOptions) (result *RecognitionJobs, response *core.DetailedResponse, err error) {
	return speechToText.CheckJobsWithContext(context.Background(), checkJobsOptions)
}

// CheckJobsWithContext is an alternate form of the CheckJobs method which supports a Context parameter
func (speechToText *SpeechToTextV1) CheckJobsWithContext(ctx context.Context, checkJobsOptions *CheckJobsOptions) (result *RecognitionJobs, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(checkJobsOptions, "checkJobsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/recognitions`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range checkJobsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "CheckJobs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = speechToText.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRecognitionJobs)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CheckJob : Check a job
// Returns information about the specified job. The response always includes the status of the job and its creation and
// update times. If the status is `completed`, the response includes the results of the recognition request. You must
// use credentials for the instance of the service that owns a job to list information about it.
//
// You can use the method to retrieve the results of any job, regardless of whether it was submitted with a callback URL
// and the `recognitions.completed_with_results` event, and you can retrieve the results multiple times for as long as
// they remain available. Use the [Check jobs](#checkjobs) method to request information about the most recent jobs
// associated with the calling credentials.
//
// **See also:** [Checking the status and retrieving the results of a
// job](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-async#job).
func (speechToText *SpeechToTextV1) CheckJob(checkJobOptions *CheckJobOptions) (result *RecognitionJob, response *core.DetailedResponse, err error) {
	return speechToText.CheckJobWithContext(context.Background(), checkJobOptions)
}

// CheckJobWithContext is an alternate form of the CheckJob method which supports a Context parameter
func (speechToText *SpeechToTextV1) CheckJobWithContext(ctx context.Context, checkJobOptions *CheckJobOptions) (result *RecognitionJob, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(checkJobOptions, "checkJobOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(checkJobOptions, "checkJobOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *checkJobOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/recognitions/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range checkJobOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "CheckJob")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = speechToText.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRecognitionJob)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteJob : Delete a job
// Deletes the specified job. You cannot delete a job that the service is actively processing. Once you delete a job,
// its results are no longer available. The service automatically deletes a job and its results when the time to live
// for the results expires. You must use credentials for the instance of the service that owns a job to delete it.
//
// **See also:** [Deleting a job](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-async#delete-async).
func (speechToText *SpeechToTextV1) DeleteJob(deleteJobOptions *DeleteJobOptions) (response *core.DetailedResponse, err error) {
	return speechToText.DeleteJobWithContext(context.Background(), deleteJobOptions)
}

// DeleteJobWithContext is an alternate form of the DeleteJob method which supports a Context parameter
func (speechToText *SpeechToTextV1) DeleteJobWithContext(ctx context.Context, deleteJobOptions *DeleteJobOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteJobOptions, "deleteJobOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteJobOptions, "deleteJobOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *deleteJobOptions.ID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/recognitions/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteJobOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "DeleteJob")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = speechToText.Service.Request(request, nil)

	return
}

// CreateLanguageModel : Create a custom language model
// Creates a new custom language model for a specified base model. The custom language model can be used only with the
// base model for which it is created. The model is owned by the instance of the service whose credentials are used to
// create it.
//
// You can create a maximum of 1024 custom language models per owning credentials. The service returns an error if you
// attempt to create more than 1024 models. You do not lose any models, but you cannot create any more until your model
// count is below the limit.
//
// **Important:** Effective 15 March 2022, previous-generation models for all languages other than Arabic and Japanese
// are deprecated. The deprecated models remain available until 15 September 2022, when they will be removed from the
// service and the documentation. You must migrate to the equivalent next-generation model by the end of service date.
// For more information, see [Migrating to next-generation
// models](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-models-migrate).
//
// **See also:**
// * [Create a custom language
// model](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-languageCreate#createModel-language)
// * [Language support for
// customization](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-custom-support).
func (speechToText *SpeechToTextV1) CreateLanguageModel(createLanguageModelOptions *CreateLanguageModelOptions) (result *LanguageModel, response *core.DetailedResponse, err error) {
	return speechToText.CreateLanguageModelWithContext(context.Background(), createLanguageModelOptions)
}

// CreateLanguageModelWithContext is an alternate form of the CreateLanguageModel method which supports a Context parameter
func (speechToText *SpeechToTextV1) CreateLanguageModelWithContext(ctx context.Context, createLanguageModelOptions *CreateLanguageModelOptions) (result *LanguageModel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createLanguageModelOptions, "createLanguageModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createLanguageModelOptions, "createLanguageModelOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/customizations`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createLanguageModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "CreateLanguageModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createLanguageModelOptions.Name != nil {
		body["name"] = createLanguageModelOptions.Name
	}
	if createLanguageModelOptions.BaseModelName != nil {
		body["base_model_name"] = createLanguageModelOptions.BaseModelName
	}
	if createLanguageModelOptions.Dialect != nil {
		body["dialect"] = createLanguageModelOptions.Dialect
	}
	if createLanguageModelOptions.Description != nil {
		body["description"] = createLanguageModelOptions.Description
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
	response, err = speechToText.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalLanguageModel)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListLanguageModels : List custom language models
// Lists information about all custom language models that are owned by an instance of the service. Use the `language`
// parameter to see all custom language models for the specified language. Omit the parameter to see all custom language
// models for all languages. You must use credentials for the instance of the service that owns a model to list
// information about it.
//
// **See also:**
// * [Listing custom language
// models](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-manageLanguageModels#listModels-language)
// * [Language support for
// customization](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-custom-support).
func (speechToText *SpeechToTextV1) ListLanguageModels(listLanguageModelsOptions *ListLanguageModelsOptions) (result *LanguageModels, response *core.DetailedResponse, err error) {
	return speechToText.ListLanguageModelsWithContext(context.Background(), listLanguageModelsOptions)
}

// ListLanguageModelsWithContext is an alternate form of the ListLanguageModels method which supports a Context parameter
func (speechToText *SpeechToTextV1) ListLanguageModelsWithContext(ctx context.Context, listLanguageModelsOptions *ListLanguageModelsOptions) (result *LanguageModels, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listLanguageModelsOptions, "listLanguageModelsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/customizations`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listLanguageModelsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "ListLanguageModels")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listLanguageModelsOptions.Language != nil {
		builder.AddQuery("language", fmt.Sprint(*listLanguageModelsOptions.Language))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = speechToText.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalLanguageModels)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetLanguageModel : Get a custom language model
// Gets information about a specified custom language model. You must use credentials for the instance of the service
// that owns a model to list information about it.
//
// **See also:**
// * [Listing custom language
// models](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-manageLanguageModels#listModels-language)
// * [Language support for
// customization](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-custom-support).
func (speechToText *SpeechToTextV1) GetLanguageModel(getLanguageModelOptions *GetLanguageModelOptions) (result *LanguageModel, response *core.DetailedResponse, err error) {
	return speechToText.GetLanguageModelWithContext(context.Background(), getLanguageModelOptions)
}

// GetLanguageModelWithContext is an alternate form of the GetLanguageModel method which supports a Context parameter
func (speechToText *SpeechToTextV1) GetLanguageModelWithContext(ctx context.Context, getLanguageModelOptions *GetLanguageModelOptions) (result *LanguageModel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getLanguageModelOptions, "getLanguageModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getLanguageModelOptions, "getLanguageModelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *getLanguageModelOptions.CustomizationID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/customizations/{customization_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getLanguageModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "GetLanguageModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = speechToText.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalLanguageModel)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteLanguageModel : Delete a custom language model
// Deletes an existing custom language model. The custom model cannot be deleted if another request, such as adding a
// corpus or grammar to the model, is currently being processed. You must use credentials for the instance of the
// service that owns a model to delete it.
//
// **See also:**
// * [Deleting a custom language
// model](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-manageLanguageModels#deleteModel-language)
// * [Language support for
// customization](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-custom-support).
func (speechToText *SpeechToTextV1) DeleteLanguageModel(deleteLanguageModelOptions *DeleteLanguageModelOptions) (response *core.DetailedResponse, err error) {
	return speechToText.DeleteLanguageModelWithContext(context.Background(), deleteLanguageModelOptions)
}

// DeleteLanguageModelWithContext is an alternate form of the DeleteLanguageModel method which supports a Context parameter
func (speechToText *SpeechToTextV1) DeleteLanguageModelWithContext(ctx context.Context, deleteLanguageModelOptions *DeleteLanguageModelOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteLanguageModelOptions, "deleteLanguageModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteLanguageModelOptions, "deleteLanguageModelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *deleteLanguageModelOptions.CustomizationID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/customizations/{customization_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteLanguageModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "DeleteLanguageModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = speechToText.Service.Request(request, nil)

	return
}

// TrainLanguageModel : Train a custom language model
// Initiates the training of a custom language model with new resources such as corpora, grammars, and custom words.
// After adding, modifying, or deleting resources for a custom language model, use this method to begin the actual
// training of the model on the latest data. You can specify whether the custom language model is to be trained with all
// words from its words resource or only with words that were added or modified by the user directly. You must use
// credentials for the instance of the service that owns a model to train it.
//
// The training method is asynchronous. It can take on the order of minutes to complete depending on the amount of data
// on which the service is being trained and the current load on the service. The method returns an HTTP 200 response
// code to indicate that the training process has begun.
//
// You can monitor the status of the training by using the [Get a custom language model](#getlanguagemodel) method to
// poll the model's status. Use a loop to check the status every 10 seconds. The method returns a `LanguageModel` object
// that includes `status` and `progress` fields. A status of `available` means that the custom model is trained and
// ready to use. The service cannot accept subsequent training requests or requests to add new resources until the
// existing request completes.
//
// **See also:**
// * [Train the custom language
// model](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-languageCreate#trainModel-language)
// * [Language support for customization](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-custom-support)
//
// ### Training failures
//
//	Training can fail to start for the following reasons:
//
// * The service is currently handling another request for the custom model, such as another training request or a
// request to add a corpus or grammar to the model.
// * No training data have been added to the custom model.
// * The custom model contains one or more invalid corpora, grammars, or words (for example, a custom word has an
// invalid sounds-like pronunciation). You can correct the invalid resources or set the `strict` parameter to `false` to
// exclude the invalid resources from the training. The model must contain at least one valid resource for training to
// succeed.
func (speechToText *SpeechToTextV1) TrainLanguageModel(trainLanguageModelOptions *TrainLanguageModelOptions) (result *TrainingResponse, response *core.DetailedResponse, err error) {
	return speechToText.TrainLanguageModelWithContext(context.Background(), trainLanguageModelOptions)
}

// TrainLanguageModelWithContext is an alternate form of the TrainLanguageModel method which supports a Context parameter
func (speechToText *SpeechToTextV1) TrainLanguageModelWithContext(ctx context.Context, trainLanguageModelOptions *TrainLanguageModelOptions) (result *TrainingResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(trainLanguageModelOptions, "trainLanguageModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(trainLanguageModelOptions, "trainLanguageModelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *trainLanguageModelOptions.CustomizationID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/customizations/{customization_id}/train`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range trainLanguageModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "TrainLanguageModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if trainLanguageModelOptions.WordTypeToAdd != nil {
		builder.AddQuery("word_type_to_add", fmt.Sprint(*trainLanguageModelOptions.WordTypeToAdd))
	}
	if trainLanguageModelOptions.CustomizationWeight != nil {
		builder.AddQuery("customization_weight", fmt.Sprint(*trainLanguageModelOptions.CustomizationWeight))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = speechToText.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTrainingResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ResetLanguageModel : Reset a custom language model
// Resets a custom language model by removing all corpora, grammars, and words from the model. Resetting a custom
// language model initializes the model to its state when it was first created. Metadata such as the name and language
// of the model are preserved, but the model's words resource is removed and must be re-created. You must use
// credentials for the instance of the service that owns a model to reset it.
//
// **See also:**
// * [Resetting a custom language
// model](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-manageLanguageModels#resetModel-language)
// * [Language support for
// customization](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-custom-support).
func (speechToText *SpeechToTextV1) ResetLanguageModel(resetLanguageModelOptions *ResetLanguageModelOptions) (response *core.DetailedResponse, err error) {
	return speechToText.ResetLanguageModelWithContext(context.Background(), resetLanguageModelOptions)
}

// ResetLanguageModelWithContext is an alternate form of the ResetLanguageModel method which supports a Context parameter
func (speechToText *SpeechToTextV1) ResetLanguageModelWithContext(ctx context.Context, resetLanguageModelOptions *ResetLanguageModelOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(resetLanguageModelOptions, "resetLanguageModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(resetLanguageModelOptions, "resetLanguageModelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *resetLanguageModelOptions.CustomizationID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/customizations/{customization_id}/reset`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range resetLanguageModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "ResetLanguageModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = speechToText.Service.Request(request, nil)

	return
}

// UpgradeLanguageModel : Upgrade a custom language model
// Initiates the upgrade of a custom language model to the latest version of its base language model. The upgrade method
// is asynchronous. It can take on the order of minutes to complete depending on the amount of data in the custom model
// and the current load on the service. A custom model must be in the `ready` or `available` state to be upgraded. You
// must use credentials for the instance of the service that owns a model to upgrade it.
//
// The method returns an HTTP 200 response code to indicate that the upgrade process has begun successfully. You can
// monitor the status of the upgrade by using the [Get a custom language model](#getlanguagemodel) method to poll the
// model's status. The method returns a `LanguageModel` object that includes `status` and `progress` fields. Use a loop
// to check the status every 10 seconds.
//
// While it is being upgraded, the custom model has the status `upgrading`. When the upgrade is complete, the model
// resumes the status that it had prior to upgrade. The service cannot accept subsequent requests for the model until
// the upgrade completes.
//
// **See also:**
// * [Upgrading a custom language
// model](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-custom-upgrade#custom-upgrade-language)
// * [Language support for
// customization](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-custom-support).
func (speechToText *SpeechToTextV1) UpgradeLanguageModel(upgradeLanguageModelOptions *UpgradeLanguageModelOptions) (response *core.DetailedResponse, err error) {
	return speechToText.UpgradeLanguageModelWithContext(context.Background(), upgradeLanguageModelOptions)
}

// UpgradeLanguageModelWithContext is an alternate form of the UpgradeLanguageModel method which supports a Context parameter
func (speechToText *SpeechToTextV1) UpgradeLanguageModelWithContext(ctx context.Context, upgradeLanguageModelOptions *UpgradeLanguageModelOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(upgradeLanguageModelOptions, "upgradeLanguageModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(upgradeLanguageModelOptions, "upgradeLanguageModelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *upgradeLanguageModelOptions.CustomizationID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/customizations/{customization_id}/upgrade_model`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range upgradeLanguageModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "UpgradeLanguageModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = speechToText.Service.Request(request, nil)

	return
}

// ListCorpora : List corpora
// Lists information about all corpora from a custom language model. The information includes the name, status, and
// total number of words for each corpus. _For custom models that are based on previous-generation models_, it also
// includes the number of out-of-vocabulary (OOV) words from the corpus. You must use credentials for the instance of
// the service that owns a model to list its corpora.
//
// **See also:** [Listing corpora for a custom language
// model](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-manageCorpora#listCorpora).
func (speechToText *SpeechToTextV1) ListCorpora(listCorporaOptions *ListCorporaOptions) (result *Corpora, response *core.DetailedResponse, err error) {
	return speechToText.ListCorporaWithContext(context.Background(), listCorporaOptions)
}

// ListCorporaWithContext is an alternate form of the ListCorpora method which supports a Context parameter
func (speechToText *SpeechToTextV1) ListCorporaWithContext(ctx context.Context, listCorporaOptions *ListCorporaOptions) (result *Corpora, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listCorporaOptions, "listCorporaOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listCorporaOptions, "listCorporaOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *listCorporaOptions.CustomizationID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/customizations/{customization_id}/corpora`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listCorporaOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "ListCorpora")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = speechToText.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCorpora)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// AddCorpus : Add a corpus
// Adds a single corpus text file of new training data to a custom language model. Use multiple requests to submit
// multiple corpus text files. You must use credentials for the instance of the service that owns a model to add a
// corpus to it. Adding a corpus does not affect the custom language model until you train the model for the new data by
// using the [Train a custom language model](#trainlanguagemodel) method.
//
// Submit a plain text file that contains sample sentences from the domain of interest to enable the service to parse
// the words in context. The more sentences you add that represent the context in which speakers use words from the
// domain, the better the service's recognition accuracy.
//
// The call returns an HTTP 201 response code if the corpus is valid. The service then asynchronously processes and
// automatically extracts data from the contents of the corpus. This operation can take on the order of minutes to
// complete depending on the current load on the service, the total number of words in the corpus, and, _for custom
// models that are based on previous-generation models_, the number of new (out-of-vocabulary) words in the corpus. You
// cannot submit requests to add additional resources to the custom model or to train the model until the service's
// analysis of the corpus for the current request completes. Use the [Get a corpus](#getcorpus) method to check the
// status of the analysis.
//
// _For custom models that are based on previous-generation models_, the service auto-populates the model's words
// resource with words from the corpus that are not found in its base vocabulary. These words are referred to as
// out-of-vocabulary (OOV) words. After adding a corpus, you must validate the words resource to ensure that each OOV
// word's definition is complete and valid. You can use the [List custom words](#listwords) method to examine the words
// resource. You can use other words method to eliminate typos and modify how words are pronounced as needed.
//
// To add a corpus file that has the same name as an existing corpus, set the `allow_overwrite` parameter to `true`;
// otherwise, the request fails. Overwriting an existing corpus causes the service to process the corpus text file and
// extract its data anew. _For a custom model that is based on a previous-generation model_, the service first removes
// any OOV words that are associated with the existing corpus from the model's words resource unless they were also
// added by another corpus or grammar, or they have been modified in some way with the [Add custom words](#addwords) or
// [Add a custom word](#addword) method.
//
// The service limits the overall amount of data that you can add to a custom model to a maximum of 10 million total
// words from all sources combined. _For a custom model that is based on a previous-generation model_, you can add no
// more than 90 thousand custom (OOV) words to a model. This includes words that the service extracts from corpora and
// grammars, and words that you add directly.
//
// **See also:**
// * [Add a corpus to the custom language
// model](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-languageCreate#addCorpus)
// * [Working with corpora for previous-generation
// models](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-corporaWords#workingCorpora)
// * [Working with corpora for next-generation
// models](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-corporaWords-ng#workingCorpora-ng)
// * [Validating a words resource for previous-generation
// models](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-corporaWords#validateModel)
// * [Validating a words resource for next-generation
// models](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-corporaWords-ng#validateModel-ng).
func (speechToText *SpeechToTextV1) AddCorpus(addCorpusOptions *AddCorpusOptions) (response *core.DetailedResponse, err error) {
	return speechToText.AddCorpusWithContext(context.Background(), addCorpusOptions)
}

// AddCorpusWithContext is an alternate form of the AddCorpus method which supports a Context parameter
func (speechToText *SpeechToTextV1) AddCorpusWithContext(ctx context.Context, addCorpusOptions *AddCorpusOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addCorpusOptions, "addCorpusOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addCorpusOptions, "addCorpusOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *addCorpusOptions.CustomizationID,
		"corpus_name":      *addCorpusOptions.CorpusName,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/customizations/{customization_id}/corpora/{corpus_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range addCorpusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "AddCorpus")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if addCorpusOptions.AllowOverwrite != nil {
		builder.AddQuery("allow_overwrite", fmt.Sprint(*addCorpusOptions.AllowOverwrite))
	}

	builder.AddFormData("corpus_file", "filename",
		"text/plain", addCorpusOptions.CorpusFile)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = speechToText.Service.Request(request, nil)

	return
}

// GetCorpus : Get a corpus
// Gets information about a corpus from a custom language model. The information includes the name, status, and total
// number of words for the corpus. _For custom models that are based on previous-generation models_, it also includes
// the number of out-of-vocabulary (OOV) words from the corpus. You must use credentials for the instance of the service
// that owns a model to list its corpora.
//
// **See also:** [Listing corpora for a custom language
// model](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-manageCorpora#listCorpora).
func (speechToText *SpeechToTextV1) GetCorpus(getCorpusOptions *GetCorpusOptions) (result *Corpus, response *core.DetailedResponse, err error) {
	return speechToText.GetCorpusWithContext(context.Background(), getCorpusOptions)
}

// GetCorpusWithContext is an alternate form of the GetCorpus method which supports a Context parameter
func (speechToText *SpeechToTextV1) GetCorpusWithContext(ctx context.Context, getCorpusOptions *GetCorpusOptions) (result *Corpus, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getCorpusOptions, "getCorpusOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getCorpusOptions, "getCorpusOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *getCorpusOptions.CustomizationID,
		"corpus_name":      *getCorpusOptions.CorpusName,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/customizations/{customization_id}/corpora/{corpus_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCorpusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "GetCorpus")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = speechToText.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCorpus)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteCorpus : Delete a corpus
// Deletes an existing corpus from a custom language model. Removing a corpus does not affect the custom model until you
// train the model with the [Train a custom language model](#trainlanguagemodel) method. You must use credentials for
// the instance of the service that owns a model to delete its corpora.
//
// _For custom models that are based on previous-generation models_, the service removes any out-of-vocabulary (OOV)
// words that are associated with the corpus from the custom model's words resource unless they were also added by
// another corpus or grammar, or they were modified in some way with the [Add custom words](#addwords) or [Add a custom
// word](#addword) method.
//
// **See also:** [Deleting a corpus from a custom language
// model](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-manageCorpora#deleteCorpus).
func (speechToText *SpeechToTextV1) DeleteCorpus(deleteCorpusOptions *DeleteCorpusOptions) (response *core.DetailedResponse, err error) {
	return speechToText.DeleteCorpusWithContext(context.Background(), deleteCorpusOptions)
}

// DeleteCorpusWithContext is an alternate form of the DeleteCorpus method which supports a Context parameter
func (speechToText *SpeechToTextV1) DeleteCorpusWithContext(ctx context.Context, deleteCorpusOptions *DeleteCorpusOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteCorpusOptions, "deleteCorpusOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteCorpusOptions, "deleteCorpusOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *deleteCorpusOptions.CustomizationID,
		"corpus_name":      *deleteCorpusOptions.CorpusName,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/customizations/{customization_id}/corpora/{corpus_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteCorpusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "DeleteCorpus")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = speechToText.Service.Request(request, nil)

	return
}

// ListWords : List custom words
// Lists information about custom words from a custom language model. You can list all words from the custom model's
// words resource, only custom words that were added or modified by the user, or, _for a custom model that is based on a
// previous-generation model_, only out-of-vocabulary (OOV) words that were extracted from corpora or are recognized by
// grammars. You can also indicate the order in which the service is to return words; by default, the service lists
// words in ascending alphabetical order. You must use credentials for the instance of the service that owns a model to
// list information about its words.
//
// **See also:** [Listing words from a custom language
// model](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-manageWords#listWords).
func (speechToText *SpeechToTextV1) ListWords(listWordsOptions *ListWordsOptions) (result *Words, response *core.DetailedResponse, err error) {
	return speechToText.ListWordsWithContext(context.Background(), listWordsOptions)
}

// ListWordsWithContext is an alternate form of the ListWords method which supports a Context parameter
func (speechToText *SpeechToTextV1) ListWordsWithContext(ctx context.Context, listWordsOptions *ListWordsOptions) (result *Words, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/customizations/{customization_id}/words`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listWordsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "ListWords")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listWordsOptions.WordType != nil {
		builder.AddQuery("word_type", fmt.Sprint(*listWordsOptions.WordType))
	}
	if listWordsOptions.Sort != nil {
		builder.AddQuery("sort", fmt.Sprint(*listWordsOptions.Sort))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = speechToText.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalWords)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// AddWords : Add custom words
// Adds one or more custom words to a custom language model. You can use this method to add words or to modify existing
// words in a custom model's words resource. _For custom models that are based on previous-generation models_, the
// service populates the words resource for a custom model with out-of-vocabulary (OOV) words from each corpus or
// grammar that is added to the model. You can use this method to modify OOV words in the model's words resource.
//
// _For a custom model that is based on a previous-generation model_, the words resource for a model can contain a
// maximum of 90 thousand custom (OOV) words. This includes words that the service extracts from corpora and grammars
// and words that you add directly.
//
// You must use credentials for the instance of the service that owns a model to add or modify custom words for the
// model. Adding or modifying custom words does not affect the custom model until you train the model for the new data
// by using the [Train a custom language model](#trainlanguagemodel) method.
//
// You add custom words by providing a `CustomWords` object, which is an array of `CustomWord` objects, one per word.
// Use the object's `word` parameter to identify the word that is to be added. You can also provide one or both of the
// optional `display_as` or `sounds_like` fields for each word.
// * The `display_as` field provides a different way of spelling the word in a transcript. Use the parameter when you
// want the word to appear different from its usual representation or from its spelling in training data. For example,
// you might indicate that the word `IBM` is to be displayed as `IBM&trade;`.
// * The `sounds_like` field, _which can be used only with a custom model that is based on a previous-generation model_,
// provides an array of one or more pronunciations for the word. Use the parameter to specify how the word can be
// pronounced by users. Use the parameter for words that are difficult to pronounce, foreign words, acronyms, and so on.
// For example, you might specify that the word `IEEE` can sound like `i triple e`. You can specify a maximum of five
// sounds-like pronunciations for a word. If you omit the `sounds_like` field, the service attempts to set the field to
// its pronunciation of the word. It cannot generate a pronunciation for all words, so you must review the word's
// definition to ensure that it is complete and valid.
//
// If you add a custom word that already exists in the words resource for the custom model, the new definition
// overwrites the existing data for the word. If the service encounters an error with the input data, it returns a
// failure code and does not add any of the words to the words resource.
//
// The call returns an HTTP 201 response code if the input data is valid. It then asynchronously processes the words to
// add them to the model's words resource. The time that it takes for the analysis to complete depends on the number of
// new words that you add but is generally faster than adding a corpus or grammar.
//
// You can monitor the status of the request by using the [Get a custom language model](#getlanguagemodel) method to
// poll the model's status. Use a loop to check the status every 10 seconds. The method returns a `Customization` object
// that includes a `status` field. A status of `ready` means that the words have been added to the custom model. The
// service cannot accept requests to add new data or to train the model until the existing request completes.
//
// You can use the [List custom words](#listwords) or [Get a custom word](#getword) method to review the words that you
// add. Words with an invalid `sounds_like` field include an `error` field that describes the problem. You can use other
// words-related methods to correct errors, eliminate typos, and modify how words are pronounced as needed.
//
// **See also:**
// * [Add words to the custom language
// model](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-languageCreate#addWords)
// * [Working with custom words for previous-generation
// models](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-corporaWords#workingWords)
// * [Working with custom words for next-generation
// models](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-corporaWords-ng#workingWords-ng)
// * [Validating a words resource for previous-generation
// models](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-corporaWords#validateModel)
// * [Validating a words resource for next-generation
// models](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-corporaWords-ng#validateModel-ng).
func (speechToText *SpeechToTextV1) AddWords(addWordsOptions *AddWordsOptions) (response *core.DetailedResponse, err error) {
	return speechToText.AddWordsWithContext(context.Background(), addWordsOptions)
}

// AddWordsWithContext is an alternate form of the AddWords method which supports a Context parameter
func (speechToText *SpeechToTextV1) AddWordsWithContext(ctx context.Context, addWordsOptions *AddWordsOptions) (response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/customizations/{customization_id}/words`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range addWordsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "AddWords")
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

	response, err = speechToText.Service.Request(request, nil)

	return
}

// AddWord : Add a custom word
// Adds a custom word to a custom language model. You can use this method to add a word or to modify an existing word in
// the words resource. _For custom models that are based on previous-generation models_, the service populates the words
// resource for a custom model with out-of-vocabulary (OOV) words from each corpus or grammar that is added to the
// model. You can use this method to modify OOV words in the model's words resource.
//
// _For a custom model that is based on a previous-generation models_, the words resource for a model can contain a
// maximum of 90 thousand custom (OOV) words. This includes words that the service extracts from corpora and grammars
// and words that you add directly.
//
// You must use credentials for the instance of the service that owns a model to add or modify a custom word for the
// model. Adding or modifying a custom word does not affect the custom model until you train the model for the new data
// by using the [Train a custom language model](#trainlanguagemodel) method.
//
// Use the `word_name` parameter to specify the custom word that is to be added or modified. Use the `CustomWord` object
// to provide one or both of the optional `display_as` or `sounds_like` fields for the word.
// * The `display_as` field provides a different way of spelling the word in a transcript. Use the parameter when you
// want the word to appear different from its usual representation or from its spelling in training data. For example,
// you might indicate that the word `IBM` is to be displayed as `IBM&trade;`.
// * The `sounds_like` field, _which can be used only with a custom model that is based on a previous-generation model_,
// provides an array of one or more pronunciations for the word. Use the parameter to specify how the word can be
// pronounced by users. Use the parameter for words that are difficult to pronounce, foreign words, acronyms, and so on.
// For example, you might specify that the word `IEEE` can sound like `i triple e`. You can specify a maximum of five
// sounds-like pronunciations for a word. If you omit the `sounds_like` field, the service attempts to set the field to
// its pronunciation of the word. It cannot generate a pronunciation for all words, so you must review the word's
// definition to ensure that it is complete and valid.
//
// If you add a custom word that already exists in the words resource for the custom model, the new definition
// overwrites the existing data for the word. If the service encounters an error, it does not add the word to the words
// resource. Use the [Get a custom word](#getword) method to review the word that you add.
//
// **See also:**
// * [Add words to the custom language
// model](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-languageCreate#addWords)
// * [Working with custom words for previous-generation
// models](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-corporaWords#workingWords)
// * [Working with custom words for next-generation
// models](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-corporaWords-ng#workingWords-ng)
// * [Validating a words resource for previous-generation
// models](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-corporaWords#validateModel)
// * [Validating a words resource for next-generation
// models](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-corporaWords-ng#validateModel-ng).
func (speechToText *SpeechToTextV1) AddWord(addWordOptions *AddWordOptions) (response *core.DetailedResponse, err error) {
	return speechToText.AddWordWithContext(context.Background(), addWordOptions)
}

// AddWordWithContext is an alternate form of the AddWord method which supports a Context parameter
func (speechToText *SpeechToTextV1) AddWordWithContext(ctx context.Context, addWordOptions *AddWordOptions) (response *core.DetailedResponse, err error) {
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
		"word_name":        *addWordOptions.WordName,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/customizations/{customization_id}/words/{word_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range addWordOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "AddWord")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if addWordOptions.Word != nil {
		body["word"] = addWordOptions.Word
	}
	if addWordOptions.SoundsLike != nil {
		body["sounds_like"] = addWordOptions.SoundsLike
	}
	if addWordOptions.DisplayAs != nil {
		body["display_as"] = addWordOptions.DisplayAs
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = speechToText.Service.Request(request, nil)

	return
}

// GetWord : Get a custom word
// Gets information about a custom word from a custom language model. You must use credentials for the instance of the
// service that owns a model to list information about its words.
//
// **See also:** [Listing words from a custom language
// model](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-manageWords#listWords).
func (speechToText *SpeechToTextV1) GetWord(getWordOptions *GetWordOptions) (result *Word, response *core.DetailedResponse, err error) {
	return speechToText.GetWordWithContext(context.Background(), getWordOptions)
}

// GetWordWithContext is an alternate form of the GetWord method which supports a Context parameter
func (speechToText *SpeechToTextV1) GetWordWithContext(ctx context.Context, getWordOptions *GetWordOptions) (result *Word, response *core.DetailedResponse, err error) {
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
		"word_name":        *getWordOptions.WordName,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/customizations/{customization_id}/words/{word_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getWordOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "GetWord")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = speechToText.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalWord)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteWord : Delete a custom word
// Deletes a custom word from a custom language model. You can remove any word that you added to the custom model's
// words resource via any means. However, if the word also exists in the service's base vocabulary, the service removes
// the word only from the words resource; the word remains in the base vocabulary. Removing a custom word does not
// affect the custom model until you train the model with the [Train a custom language model](#trainlanguagemodel)
// method. You must use credentials for the instance of the service that owns a model to delete its words.
//
// **See also:** [Deleting a word from a custom language
// model](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-manageWords#deleteWord).
func (speechToText *SpeechToTextV1) DeleteWord(deleteWordOptions *DeleteWordOptions) (response *core.DetailedResponse, err error) {
	return speechToText.DeleteWordWithContext(context.Background(), deleteWordOptions)
}

// DeleteWordWithContext is an alternate form of the DeleteWord method which supports a Context parameter
func (speechToText *SpeechToTextV1) DeleteWordWithContext(ctx context.Context, deleteWordOptions *DeleteWordOptions) (response *core.DetailedResponse, err error) {
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
		"word_name":        *deleteWordOptions.WordName,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/customizations/{customization_id}/words/{word_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteWordOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "DeleteWord")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = speechToText.Service.Request(request, nil)

	return
}

// ListGrammars : List grammars
// Lists information about all grammars from a custom language model. For each grammar, the information includes the
// name, status, and (for grammars that are based on previous-generation models) the total number of out-of-vocabulary
// (OOV) words. You must use credentials for the instance of the service that owns a model to list its grammars.
//
// **See also:**
// * [Listing grammars from a custom language
// model](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-manageGrammars#listGrammars)
// * [Language support for
// customization](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-custom-support).
func (speechToText *SpeechToTextV1) ListGrammars(listGrammarsOptions *ListGrammarsOptions) (result *Grammars, response *core.DetailedResponse, err error) {
	return speechToText.ListGrammarsWithContext(context.Background(), listGrammarsOptions)
}

// ListGrammarsWithContext is an alternate form of the ListGrammars method which supports a Context parameter
func (speechToText *SpeechToTextV1) ListGrammarsWithContext(ctx context.Context, listGrammarsOptions *ListGrammarsOptions) (result *Grammars, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listGrammarsOptions, "listGrammarsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listGrammarsOptions, "listGrammarsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *listGrammarsOptions.CustomizationID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/customizations/{customization_id}/grammars`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listGrammarsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "ListGrammars")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = speechToText.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGrammars)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// AddGrammar : Add a grammar
// Adds a single grammar file to a custom language model. Submit a plain text file in UTF-8 format that defines the
// grammar. Use multiple requests to submit multiple grammar files. You must use credentials for the instance of the
// service that owns a model to add a grammar to it. Adding a grammar does not affect the custom language model until
// you train the model for the new data by using the [Train a custom language model](#trainlanguagemodel) method.
//
// The call returns an HTTP 201 response code if the grammar is valid. The service then asynchronously processes the
// contents of the grammar and automatically extracts new words that it finds. This operation can take a few seconds or
// minutes to complete depending on the size and complexity of the grammar, as well as the current load on the service.
// You cannot submit requests to add additional resources to the custom model or to train the model until the service's
// analysis of the grammar for the current request completes. Use the [Get a grammar](#getgrammar) method to check the
// status of the analysis.
//
// _For grammars that are based on previous-generation models,_ the service populates the model's words resource with
// any word that is recognized by the grammar that is not found in the model's base vocabulary. These are referred to as
// out-of-vocabulary (OOV) words. You can use the [List custom words](#listwords) method to examine the words resource
// and use other words-related methods to eliminate typos and modify how words are pronounced as needed. _For grammars
// that are based on next-generation models,_ the service extracts no OOV words from the grammars.
//
// To add a grammar that has the same name as an existing grammar, set the `allow_overwrite` parameter to `true`;
// otherwise, the request fails. Overwriting an existing grammar causes the service to process the grammar file and
// extract OOV words anew. Before doing so, it removes any OOV words associated with the existing grammar from the
// model's words resource unless they were also added by another resource or they have been modified in some way with
// the [Add custom words](#addwords) or [Add a custom word](#addword) method.
//
// _For grammars that are based on previous-generation models,_ the service limits the overall amount of data that you
// can add to a custom model to a maximum of 10 million total words from all sources combined. Also, you can add no more
// than 90 thousand OOV words to a model. This includes words that the service extracts from corpora and grammars and
// words that you add directly.
//
// **See also:**
// * [Understanding
// grammars](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-grammarUnderstand#grammarUnderstand)
// * [Add a grammar to the custom language
// model](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-grammarAdd#addGrammar)
// * [Language support for
// customization](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-custom-support).
func (speechToText *SpeechToTextV1) AddGrammar(addGrammarOptions *AddGrammarOptions) (response *core.DetailedResponse, err error) {
	return speechToText.AddGrammarWithContext(context.Background(), addGrammarOptions)
}

// AddGrammarWithContext is an alternate form of the AddGrammar method which supports a Context parameter
func (speechToText *SpeechToTextV1) AddGrammarWithContext(ctx context.Context, addGrammarOptions *AddGrammarOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addGrammarOptions, "addGrammarOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addGrammarOptions, "addGrammarOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *addGrammarOptions.CustomizationID,
		"grammar_name":     *addGrammarOptions.GrammarName,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/customizations/{customization_id}/grammars/{grammar_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range addGrammarOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "AddGrammar")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if addGrammarOptions.ContentType != nil {
		builder.AddHeader("Content-Type", fmt.Sprint(*addGrammarOptions.ContentType))
	}

	if addGrammarOptions.AllowOverwrite != nil {
		builder.AddQuery("allow_overwrite", fmt.Sprint(*addGrammarOptions.AllowOverwrite))
	}

	_, err = builder.SetBodyContent(core.StringNilMapper(addGrammarOptions.ContentType), nil, nil, addGrammarOptions.GrammarFile)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = speechToText.Service.Request(request, nil)

	return
}

// GetGrammar : Get a grammar
// Gets information about a grammar from a custom language model. For each grammar, the information includes the name,
// status, and (for grammars that are based on previous-generation models) the total number of out-of-vocabulary (OOV)
// words. You must use credentials for the instance of the service that owns a model to list its grammars.
//
// **See also:**
// * [Listing grammars from a custom language
// model](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-manageGrammars#listGrammars)
// * [Language support for
// customization](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-custom-support).
func (speechToText *SpeechToTextV1) GetGrammar(getGrammarOptions *GetGrammarOptions) (result *Grammar, response *core.DetailedResponse, err error) {
	return speechToText.GetGrammarWithContext(context.Background(), getGrammarOptions)
}

// GetGrammarWithContext is an alternate form of the GetGrammar method which supports a Context parameter
func (speechToText *SpeechToTextV1) GetGrammarWithContext(ctx context.Context, getGrammarOptions *GetGrammarOptions) (result *Grammar, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getGrammarOptions, "getGrammarOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getGrammarOptions, "getGrammarOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *getGrammarOptions.CustomizationID,
		"grammar_name":     *getGrammarOptions.GrammarName,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/customizations/{customization_id}/grammars/{grammar_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getGrammarOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "GetGrammar")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = speechToText.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGrammar)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteGrammar : Delete a grammar
// Deletes an existing grammar from a custom language model. _For grammars that are based on previous-generation
// models,_ the service removes any out-of-vocabulary (OOV) words associated with the grammar from the custom model's
// words resource unless they were also added by another resource or they were modified in some way with the [Add custom
// words](#addwords) or [Add a custom word](#addword) method. Removing a grammar does not affect the custom model until
// you train the model with the [Train a custom language model](#trainlanguagemodel) method. You must use credentials
// for the instance of the service that owns a model to delete its grammar.
//
// **See also:**
// * [Deleting a grammar from a custom language
// model](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-manageGrammars#deleteGrammar)
// * [Language support for
// customization](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-custom-support).
func (speechToText *SpeechToTextV1) DeleteGrammar(deleteGrammarOptions *DeleteGrammarOptions) (response *core.DetailedResponse, err error) {
	return speechToText.DeleteGrammarWithContext(context.Background(), deleteGrammarOptions)
}

// DeleteGrammarWithContext is an alternate form of the DeleteGrammar method which supports a Context parameter
func (speechToText *SpeechToTextV1) DeleteGrammarWithContext(ctx context.Context, deleteGrammarOptions *DeleteGrammarOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteGrammarOptions, "deleteGrammarOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteGrammarOptions, "deleteGrammarOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *deleteGrammarOptions.CustomizationID,
		"grammar_name":     *deleteGrammarOptions.GrammarName,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/customizations/{customization_id}/grammars/{grammar_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteGrammarOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "DeleteGrammar")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = speechToText.Service.Request(request, nil)

	return
}

// CreateAcousticModel : Create a custom acoustic model
// Creates a new custom acoustic model for a specified base model. The custom acoustic model can be used only with the
// base model for which it is created. The model is owned by the instance of the service whose credentials are used to
// create it.
//
// You can create a maximum of 1024 custom acoustic models per owning credentials. The service returns an error if you
// attempt to create more than 1024 models. You do not lose any models, but you cannot create any more until your model
// count is below the limit.
//
// **Note:** Acoustic model customization is supported only for use with previous-generation models. It is not supported
// for next-generation models.
//
// **Important:** Effective 15 March 2022, previous-generation models for all languages other than Arabic and Japanese
// are deprecated. The deprecated models remain available until 15 September 2022, when they will be removed from the
// service and the documentation. You must migrate to the equivalent next-generation model by the end of service date.
// For more information, see [Migrating to next-generation
// models](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-models-migrate).
//
// **See also:** [Create a custom acoustic
// model](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-acoustic#createModel-acoustic).
func (speechToText *SpeechToTextV1) CreateAcousticModel(createAcousticModelOptions *CreateAcousticModelOptions) (result *AcousticModel, response *core.DetailedResponse, err error) {
	return speechToText.CreateAcousticModelWithContext(context.Background(), createAcousticModelOptions)
}

// CreateAcousticModelWithContext is an alternate form of the CreateAcousticModel method which supports a Context parameter
func (speechToText *SpeechToTextV1) CreateAcousticModelWithContext(ctx context.Context, createAcousticModelOptions *CreateAcousticModelOptions) (result *AcousticModel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createAcousticModelOptions, "createAcousticModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createAcousticModelOptions, "createAcousticModelOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/acoustic_customizations`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createAcousticModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "CreateAcousticModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createAcousticModelOptions.Name != nil {
		body["name"] = createAcousticModelOptions.Name
	}
	if createAcousticModelOptions.BaseModelName != nil {
		body["base_model_name"] = createAcousticModelOptions.BaseModelName
	}
	if createAcousticModelOptions.Description != nil {
		body["description"] = createAcousticModelOptions.Description
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
	response, err = speechToText.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAcousticModel)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListAcousticModels : List custom acoustic models
// Lists information about all custom acoustic models that are owned by an instance of the service. Use the `language`
// parameter to see all custom acoustic models for the specified language. Omit the parameter to see all custom acoustic
// models for all languages. You must use credentials for the instance of the service that owns a model to list
// information about it.
//
// **Note:** Acoustic model customization is supported only for use with previous-generation models. It is not supported
// for next-generation models.
//
// **See also:** [Listing custom acoustic
// models](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-manageAcousticModels#listModels-acoustic).
func (speechToText *SpeechToTextV1) ListAcousticModels(listAcousticModelsOptions *ListAcousticModelsOptions) (result *AcousticModels, response *core.DetailedResponse, err error) {
	return speechToText.ListAcousticModelsWithContext(context.Background(), listAcousticModelsOptions)
}

// ListAcousticModelsWithContext is an alternate form of the ListAcousticModels method which supports a Context parameter
func (speechToText *SpeechToTextV1) ListAcousticModelsWithContext(ctx context.Context, listAcousticModelsOptions *ListAcousticModelsOptions) (result *AcousticModels, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listAcousticModelsOptions, "listAcousticModelsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/acoustic_customizations`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listAcousticModelsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "ListAcousticModels")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listAcousticModelsOptions.Language != nil {
		builder.AddQuery("language", fmt.Sprint(*listAcousticModelsOptions.Language))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = speechToText.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAcousticModels)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetAcousticModel : Get a custom acoustic model
// Gets information about a specified custom acoustic model. You must use credentials for the instance of the service
// that owns a model to list information about it.
//
// **Note:** Acoustic model customization is supported only for use with previous-generation models. It is not supported
// for next-generation models.
//
// **See also:** [Listing custom acoustic
// models](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-manageAcousticModels#listModels-acoustic).
func (speechToText *SpeechToTextV1) GetAcousticModel(getAcousticModelOptions *GetAcousticModelOptions) (result *AcousticModel, response *core.DetailedResponse, err error) {
	return speechToText.GetAcousticModelWithContext(context.Background(), getAcousticModelOptions)
}

// GetAcousticModelWithContext is an alternate form of the GetAcousticModel method which supports a Context parameter
func (speechToText *SpeechToTextV1) GetAcousticModelWithContext(ctx context.Context, getAcousticModelOptions *GetAcousticModelOptions) (result *AcousticModel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAcousticModelOptions, "getAcousticModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAcousticModelOptions, "getAcousticModelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *getAcousticModelOptions.CustomizationID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/acoustic_customizations/{customization_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAcousticModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "GetAcousticModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = speechToText.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAcousticModel)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteAcousticModel : Delete a custom acoustic model
// Deletes an existing custom acoustic model. The custom model cannot be deleted if another request, such as adding an
// audio resource to the model, is currently being processed. You must use credentials for the instance of the service
// that owns a model to delete it.
//
// **Note:** Acoustic model customization is supported only for use with previous-generation models. It is not supported
// for next-generation models.
//
// **See also:** [Deleting a custom acoustic
// model](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-manageAcousticModels#deleteModel-acoustic).
func (speechToText *SpeechToTextV1) DeleteAcousticModel(deleteAcousticModelOptions *DeleteAcousticModelOptions) (response *core.DetailedResponse, err error) {
	return speechToText.DeleteAcousticModelWithContext(context.Background(), deleteAcousticModelOptions)
}

// DeleteAcousticModelWithContext is an alternate form of the DeleteAcousticModel method which supports a Context parameter
func (speechToText *SpeechToTextV1) DeleteAcousticModelWithContext(ctx context.Context, deleteAcousticModelOptions *DeleteAcousticModelOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteAcousticModelOptions, "deleteAcousticModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteAcousticModelOptions, "deleteAcousticModelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *deleteAcousticModelOptions.CustomizationID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/acoustic_customizations/{customization_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteAcousticModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "DeleteAcousticModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = speechToText.Service.Request(request, nil)

	return
}

// TrainAcousticModel : Train a custom acoustic model
// Initiates the training of a custom acoustic model with new or changed audio resources. After adding or deleting audio
// resources for a custom acoustic model, use this method to begin the actual training of the model on the latest audio
// data. The custom acoustic model does not reflect its changed data until you train it. You must use credentials for
// the instance of the service that owns a model to train it.
//
// The training method is asynchronous. Training time depends on the cumulative amount of audio data that the custom
// acoustic model contains and the current load on the service. When you train or retrain a model, the service uses all
// of the model's audio data in the training. Training a custom acoustic model takes approximately as long as the length
// of its cumulative audio data. For example, it takes approximately 2 hours to train a model that contains a total of 2
// hours of audio. The method returns an HTTP 200 response code to indicate that the training process has begun.
//
// You can monitor the status of the training by using the [Get a custom acoustic model](#getacousticmodel) method to
// poll the model's status. Use a loop to check the status once a minute. The method returns an `AcousticModel` object
// that includes `status` and `progress` fields. A status of `available` indicates that the custom model is trained and
// ready to use. The service cannot train a model while it is handling another request for the model. The service cannot
// accept subsequent training requests, or requests to add new audio resources, until the existing training request
// completes.
//
// You can use the optional `custom_language_model_id` parameter to specify the GUID of a separately created custom
// language model that is to be used during training. Train with a custom language model if you have verbatim
// transcriptions of the audio files that you have added to the custom model or you have either corpora (text files) or
// a list of words that are relevant to the contents of the audio files. For training to succeed, both of the custom
// models must be based on the same version of the same base model, and the custom language model must be fully trained
// and available.
//
// **Note:** Acoustic model customization is supported only for use with previous-generation models. It is not supported
// for next-generation models.
//
// **See also:**
// * [Train the custom acoustic
// model](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-acoustic#trainModel-acoustic)
// * [Using custom acoustic and custom language models
// together](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-useBoth#useBoth)
//
// ### Training failures
//
//	Training can fail to start for the following reasons:
//
// * The service is currently handling another request for the custom model, such as another training request or a
// request to add audio resources to the model.
// * The custom model contains less than 10 minutes or more than 200 hours of audio data.
// * You passed a custom language model with the `custom_language_model_id` query parameter that is not in the available
// state. A custom language model must be fully trained and available to be used to train a custom acoustic model.
// * You passed an incompatible custom language model with the `custom_language_model_id` query parameter. Both custom
// models must be based on the same version of the same base model.
// * The custom model contains one or more invalid audio resources. You can correct the invalid audio resources or set
// the `strict` parameter to `false` to exclude the invalid resources from the training. The model must contain at least
// one valid resource for training to succeed.
func (speechToText *SpeechToTextV1) TrainAcousticModel(trainAcousticModelOptions *TrainAcousticModelOptions) (result *TrainingResponse, response *core.DetailedResponse, err error) {
	return speechToText.TrainAcousticModelWithContext(context.Background(), trainAcousticModelOptions)
}

// TrainAcousticModelWithContext is an alternate form of the TrainAcousticModel method which supports a Context parameter
func (speechToText *SpeechToTextV1) TrainAcousticModelWithContext(ctx context.Context, trainAcousticModelOptions *TrainAcousticModelOptions) (result *TrainingResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(trainAcousticModelOptions, "trainAcousticModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(trainAcousticModelOptions, "trainAcousticModelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *trainAcousticModelOptions.CustomizationID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/acoustic_customizations/{customization_id}/train`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range trainAcousticModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "TrainAcousticModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if trainAcousticModelOptions.CustomLanguageModelID != nil {
		builder.AddQuery("custom_language_model_id", fmt.Sprint(*trainAcousticModelOptions.CustomLanguageModelID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = speechToText.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTrainingResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ResetAcousticModel : Reset a custom acoustic model
// Resets a custom acoustic model by removing all audio resources from the model. Resetting a custom acoustic model
// initializes the model to its state when it was first created. Metadata such as the name and language of the model are
// preserved, but the model's audio resources are removed and must be re-created. The service cannot reset a model while
// it is handling another request for the model. The service cannot accept subsequent requests for the model until the
// existing reset request completes. You must use credentials for the instance of the service that owns a model to reset
// it.
//
// **Note:** Acoustic model customization is supported only for use with previous-generation models. It is not supported
// for next-generation models.
//
// **See also:** [Resetting a custom acoustic
// model](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-manageAcousticModels#resetModel-acoustic).
func (speechToText *SpeechToTextV1) ResetAcousticModel(resetAcousticModelOptions *ResetAcousticModelOptions) (response *core.DetailedResponse, err error) {
	return speechToText.ResetAcousticModelWithContext(context.Background(), resetAcousticModelOptions)
}

// ResetAcousticModelWithContext is an alternate form of the ResetAcousticModel method which supports a Context parameter
func (speechToText *SpeechToTextV1) ResetAcousticModelWithContext(ctx context.Context, resetAcousticModelOptions *ResetAcousticModelOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(resetAcousticModelOptions, "resetAcousticModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(resetAcousticModelOptions, "resetAcousticModelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *resetAcousticModelOptions.CustomizationID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/acoustic_customizations/{customization_id}/reset`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range resetAcousticModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "ResetAcousticModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = speechToText.Service.Request(request, nil)

	return
}

// UpgradeAcousticModel : Upgrade a custom acoustic model
// Initiates the upgrade of a custom acoustic model to the latest version of its base language model. The upgrade method
// is asynchronous. It can take on the order of minutes or hours to complete depending on the amount of data in the
// custom model and the current load on the service; typically, upgrade takes approximately twice the length of the
// total audio contained in the custom model. A custom model must be in the `ready` or `available` state to be upgraded.
// You must use credentials for the instance of the service that owns a model to upgrade it.
//
// The method returns an HTTP 200 response code to indicate that the upgrade process has begun successfully. You can
// monitor the status of the upgrade by using the [Get a custom acoustic model](#getacousticmodel) method to poll the
// model's status. The method returns an `AcousticModel` object that includes `status` and `progress` fields. Use a loop
// to check the status once a minute.
//
// While it is being upgraded, the custom model has the status `upgrading`. When the upgrade is complete, the model
// resumes the status that it had prior to upgrade. The service cannot upgrade a model while it is handling another
// request for the model. The service cannot accept subsequent requests for the model until the existing upgrade request
// completes.
//
// If the custom acoustic model was trained with a separately created custom language model, you must use the
// `custom_language_model_id` parameter to specify the GUID of that custom language model. The custom language model
// must be upgraded before the custom acoustic model can be upgraded. Omit the parameter if the custom acoustic model
// was not trained with a custom language model.
//
// **Note:** Acoustic model customization is supported only for use with previous-generation models. It is not supported
// for next-generation models.
//
// **See also:** [Upgrading a custom acoustic
// model](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-custom-upgrade#custom-upgrade-acoustic).
func (speechToText *SpeechToTextV1) UpgradeAcousticModel(upgradeAcousticModelOptions *UpgradeAcousticModelOptions) (response *core.DetailedResponse, err error) {
	return speechToText.UpgradeAcousticModelWithContext(context.Background(), upgradeAcousticModelOptions)
}

// UpgradeAcousticModelWithContext is an alternate form of the UpgradeAcousticModel method which supports a Context parameter
func (speechToText *SpeechToTextV1) UpgradeAcousticModelWithContext(ctx context.Context, upgradeAcousticModelOptions *UpgradeAcousticModelOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(upgradeAcousticModelOptions, "upgradeAcousticModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(upgradeAcousticModelOptions, "upgradeAcousticModelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *upgradeAcousticModelOptions.CustomizationID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/acoustic_customizations/{customization_id}/upgrade_model`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range upgradeAcousticModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "UpgradeAcousticModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if upgradeAcousticModelOptions.CustomLanguageModelID != nil {
		builder.AddQuery("custom_language_model_id", fmt.Sprint(*upgradeAcousticModelOptions.CustomLanguageModelID))
	}
	if upgradeAcousticModelOptions.Force != nil {
		builder.AddQuery("force", fmt.Sprint(*upgradeAcousticModelOptions.Force))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = speechToText.Service.Request(request, nil)

	return
}

// ListAudio : List audio resources
// Lists information about all audio resources from a custom acoustic model. The information includes the name of the
// resource and information about its audio data, such as its duration. It also includes the status of the audio
// resource, which is important for checking the service's analysis of the resource in response to a request to add it
// to the custom acoustic model. You must use credentials for the instance of the service that owns a model to list its
// audio resources.
//
// **Note:** Acoustic model customization is supported only for use with previous-generation models. It is not supported
// for next-generation models.
//
// **See also:** [Listing audio resources for a custom acoustic
// model](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-manageAudio#listAudio).
func (speechToText *SpeechToTextV1) ListAudio(listAudioOptions *ListAudioOptions) (result *AudioResources, response *core.DetailedResponse, err error) {
	return speechToText.ListAudioWithContext(context.Background(), listAudioOptions)
}

// ListAudioWithContext is an alternate form of the ListAudio method which supports a Context parameter
func (speechToText *SpeechToTextV1) ListAudioWithContext(ctx context.Context, listAudioOptions *ListAudioOptions) (result *AudioResources, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listAudioOptions, "listAudioOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listAudioOptions, "listAudioOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *listAudioOptions.CustomizationID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/acoustic_customizations/{customization_id}/audio`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listAudioOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "ListAudio")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = speechToText.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAudioResources)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// AddAudio : Add an audio resource
// Adds an audio resource to a custom acoustic model. Add audio content that reflects the acoustic characteristics of
// the audio that you plan to transcribe. You must use credentials for the instance of the service that owns a model to
// add an audio resource to it. Adding audio data does not affect the custom acoustic model until you train the model
// for the new data by using the [Train a custom acoustic model](#trainacousticmodel) method.
//
// You can add individual audio files or an archive file that contains multiple audio files. Adding multiple audio files
// via a single archive file is significantly more efficient than adding each file individually. You can add audio
// resources in any format that the service supports for speech recognition.
//
// You can use this method to add any number of audio resources to a custom model by calling the method once for each
// audio or archive file. You can add multiple different audio resources at the same time. You must add a minimum of 10
// minutes and a maximum of 200 hours of audio that includes speech, not just silence, to a custom acoustic model before
// you can train it. No audio resource, audio- or archive-type, can be larger than 100 MB. To add an audio resource that
// has the same name as an existing audio resource, set the `allow_overwrite` parameter to `true`; otherwise, the
// request fails.
//
// The method is asynchronous. It can take several seconds or minutes to complete depending on the duration of the audio
// and, in the case of an archive file, the total number of audio files being processed. The service returns a 201
// response code if the audio is valid. It then asynchronously analyzes the contents of the audio file or files and
// automatically extracts information about the audio such as its length, sampling rate, and encoding. You cannot submit
// requests to train or upgrade the model until the service's analysis of all audio resources for current requests
// completes.
//
// To determine the status of the service's analysis of the audio, use the [Get an audio resource](#getaudio) method to
// poll the status of the audio. The method accepts the customization ID of the custom model and the name of the audio
// resource, and it returns the status of the resource. Use a loop to check the status of the audio every few seconds
// until it becomes `ok`.
//
// **Note:** Acoustic model customization is supported only for use with previous-generation models. It is not supported
// for next-generation models.
//
// **See also:** [Add audio to the custom acoustic
// model](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-acoustic#addAudio).
//
// ### Content types for audio-type resources
//
//	You can add an individual audio file in any format that the service supports for speech recognition. For an
//
// audio-type resource, use the `Content-Type` parameter to specify the audio format (MIME type) of the audio file,
// including specifying the sampling rate, channels, and endianness where indicated.
// * `audio/alaw` (Specify the sampling rate (`rate`) of the audio.)
// * `audio/basic` (Use only with narrowband models.)
// * `audio/flac`
// * `audio/g729` (Use only with narrowband models.)
// * `audio/l16` (Specify the sampling rate (`rate`) and optionally the number of channels (`channels`) and endianness
// (`endianness`) of the audio.)
// * `audio/mp3`
// * `audio/mpeg`
// * `audio/mulaw` (Specify the sampling rate (`rate`) of the audio.)
// * `audio/ogg` (The service automatically detects the codec of the input audio.)
// * `audio/ogg;codecs=opus`
// * `audio/ogg;codecs=vorbis`
// * `audio/wav` (Provide audio with a maximum of nine channels.)
// * `audio/webm` (The service automatically detects the codec of the input audio.)
// * `audio/webm;codecs=opus`
// * `audio/webm;codecs=vorbis`
//
// The sampling rate of an audio file must match the sampling rate of the base model for the custom model: for broadband
// models, at least 16 kHz; for narrowband models, at least 8 kHz. If the sampling rate of the audio is higher than the
// minimum required rate, the service down-samples the audio to the appropriate rate. If the sampling rate of the audio
// is lower than the minimum required rate, the service labels the audio file as `invalid`.
//
//	**See also:** [Supported audio
//
// formats](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-audio-formats).
//
// ### Content types for archive-type resources
//
//	You can add an archive file (**.zip** or **.tar.gz** file) that contains audio files in any format that the service
//
// supports for speech recognition. For an archive-type resource, use the `Content-Type` parameter to specify the media
// type of the archive file:
// * `application/zip` for a **.zip** file
// * `application/gzip` for a **.tar.gz** file.
//
// When you add an archive-type resource, the `Contained-Content-Type` header is optional depending on the format of the
// files that you are adding:
// * For audio files of type `audio/alaw`, `audio/basic`, `audio/l16`, or `audio/mulaw`, you must use the
// `Contained-Content-Type` header to specify the format of the contained audio files. Include the `rate`, `channels`,
// and `endianness` parameters where necessary. In this case, all audio files contained in the archive file must have
// the same audio format.
// * For audio files of all other types, you can omit the `Contained-Content-Type` header. In this case, the audio files
// contained in the archive file can have any of the formats not listed in the previous bullet. The audio files do not
// need to have the same format.
//
// Do not use the `Contained-Content-Type` header when adding an audio-type resource.
//
// ### Naming restrictions for embedded audio files
//
//	The name of an audio file that is contained in an archive-type resource can include a maximum of 128 characters.
//
// This includes the file extension and all elements of the name (for example, slashes).
func (speechToText *SpeechToTextV1) AddAudio(addAudioOptions *AddAudioOptions) (response *core.DetailedResponse, err error) {
	return speechToText.AddAudioWithContext(context.Background(), addAudioOptions)
}

// AddAudioWithContext is an alternate form of the AddAudio method which supports a Context parameter
func (speechToText *SpeechToTextV1) AddAudioWithContext(ctx context.Context, addAudioOptions *AddAudioOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addAudioOptions, "addAudioOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addAudioOptions, "addAudioOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *addAudioOptions.CustomizationID,
		"audio_name":       *addAudioOptions.AudioName,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/acoustic_customizations/{customization_id}/audio/{audio_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range addAudioOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "AddAudio")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if addAudioOptions.ContentType != nil {
		builder.AddHeader("Content-Type", fmt.Sprint(*addAudioOptions.ContentType))
	}
	if addAudioOptions.ContainedContentType != nil {
		builder.AddHeader("Contained-Content-Type", fmt.Sprint(*addAudioOptions.ContainedContentType))
	}

	if addAudioOptions.AllowOverwrite != nil {
		builder.AddQuery("allow_overwrite", fmt.Sprint(*addAudioOptions.AllowOverwrite))
	}

	_, err = builder.SetBodyContent(core.StringNilMapper(addAudioOptions.ContentType), nil, nil, addAudioOptions.AudioResource)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = speechToText.Service.Request(request, nil)

	return
}

// GetAudio : Get an audio resource
// Gets information about an audio resource from a custom acoustic model. The method returns an `AudioListing` object
// whose fields depend on the type of audio resource that you specify with the method's `audio_name` parameter:
// * _For an audio-type resource_, the object's fields match those of an `AudioResource` object: `duration`, `name`,
// `details`, and `status`.
// * _For an archive-type resource_, the object includes a `container` field whose fields match those of an
// `AudioResource` object. It also includes an `audio` field, which contains an array of `AudioResource` objects that
// provides information about the audio files that are contained in the archive.
//
// The information includes the status of the specified audio resource. The status is important for checking the
// service's analysis of a resource that you add to the custom model.
// * _For an audio-type resource_, the `status` field is located in the `AudioListing` object.
// * _For an archive-type resource_, the `status` field is located in the `AudioResource` object that is returned in the
// `container` field.
//
// You must use credentials for the instance of the service that owns a model to list its audio resources.
//
// **Note:** Acoustic model customization is supported only for use with previous-generation models. It is not supported
// for next-generation models.
//
// **See also:** [Listing audio resources for a custom acoustic
// model](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-manageAudio#listAudio).
func (speechToText *SpeechToTextV1) GetAudio(getAudioOptions *GetAudioOptions) (result *AudioListing, response *core.DetailedResponse, err error) {
	return speechToText.GetAudioWithContext(context.Background(), getAudioOptions)
}

// GetAudioWithContext is an alternate form of the GetAudio method which supports a Context parameter
func (speechToText *SpeechToTextV1) GetAudioWithContext(ctx context.Context, getAudioOptions *GetAudioOptions) (result *AudioListing, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAudioOptions, "getAudioOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAudioOptions, "getAudioOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *getAudioOptions.CustomizationID,
		"audio_name":       *getAudioOptions.AudioName,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/acoustic_customizations/{customization_id}/audio/{audio_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAudioOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "GetAudio")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = speechToText.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAudioListing)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteAudio : Delete an audio resource
// Deletes an existing audio resource from a custom acoustic model. Deleting an archive-type audio resource removes the
// entire archive of files. The service does not allow deletion of individual files from an archive resource.
//
// Removing an audio resource does not affect the custom model until you train the model on its updated data by using
// the [Train a custom acoustic model](#trainacousticmodel) method. You can delete an existing audio resource from a
// model while a different resource is being added to the model. You must use credentials for the instance of the
// service that owns a model to delete its audio resources.
//
// **Note:** Acoustic model customization is supported only for use with previous-generation models. It is not supported
// for next-generation models.
//
// **See also:** [Deleting an audio resource from a custom acoustic
// model](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-manageAudio#deleteAudio).
func (speechToText *SpeechToTextV1) DeleteAudio(deleteAudioOptions *DeleteAudioOptions) (response *core.DetailedResponse, err error) {
	return speechToText.DeleteAudioWithContext(context.Background(), deleteAudioOptions)
}

// DeleteAudioWithContext is an alternate form of the DeleteAudio method which supports a Context parameter
func (speechToText *SpeechToTextV1) DeleteAudioWithContext(ctx context.Context, deleteAudioOptions *DeleteAudioOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteAudioOptions, "deleteAudioOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteAudioOptions, "deleteAudioOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *deleteAudioOptions.CustomizationID,
		"audio_name":       *deleteAudioOptions.AudioName,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/acoustic_customizations/{customization_id}/audio/{audio_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteAudioOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "DeleteAudio")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = speechToText.Service.Request(request, nil)

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
// instance is automatically deleted. This includes all custom language models, corpora, grammars, and words; all custom
// acoustic models and audio resources; all registered endpoints for the asynchronous HTTP interface; and all data
// related to speech recognition requests.
//
// **See also:** [Information
// security](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-information-security#information-security).
func (speechToText *SpeechToTextV1) DeleteUserData(deleteUserDataOptions *DeleteUserDataOptions) (response *core.DetailedResponse, err error) {
	return speechToText.DeleteUserDataWithContext(context.Background(), deleteUserDataOptions)
}

// DeleteUserDataWithContext is an alternate form of the DeleteUserData method which supports a Context parameter
func (speechToText *SpeechToTextV1) DeleteUserDataWithContext(ctx context.Context, deleteUserDataOptions *DeleteUserDataOptions) (response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = speechToText.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(speechToText.Service.Options.URL, `/v1/user_data`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteUserDataOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "DeleteUserData")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("customer_id", fmt.Sprint(*deleteUserDataOptions.CustomerID))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = speechToText.Service.Request(request, nil)

	return
}

// AcousticModel : Information about an existing custom acoustic model.
type AcousticModel struct {
	// The customization ID (GUID) of the custom acoustic model. The [Create a custom acoustic model](#createacousticmodel)
	// method returns only this field of the object; it does not return the other fields.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// The date and time in Coordinated Universal Time (UTC) at which the custom acoustic model was created. The value is
	// provided in full ISO 8601 format (`YYYY-MM-DDThh:mm:ss.sTZD`).
	Created *string `json:"created,omitempty"`

	// The date and time in Coordinated Universal Time (UTC) at which the custom acoustic model was last modified. The
	// `created` and `updated` fields are equal when an acoustic model is first added but has yet to be updated. The value
	// is provided in full ISO 8601 format (YYYY-MM-DDThh:mm:ss.sTZD).
	Updated *string `json:"updated,omitempty"`

	// The language identifier of the custom acoustic model (for example, `en-US`).
	Language *string `json:"language,omitempty"`

	// A list of the available versions of the custom acoustic model. Each element of the array indicates a version of the
	// base model with which the custom model can be used. Multiple versions exist only if the custom model has been
	// upgraded to a new version of its base model. Otherwise, only a single version is shown.
	Versions []string `json:"versions,omitempty"`

	// The GUID of the credentials for the instance of the service that owns the custom acoustic model.
	Owner *string `json:"owner,omitempty"`

	// The name of the custom acoustic model.
	Name *string `json:"name,omitempty"`

	// The description of the custom acoustic model.
	Description *string `json:"description,omitempty"`

	// The name of the language model for which the custom acoustic model was created.
	BaseModelName *string `json:"base_model_name,omitempty"`

	// The current status of the custom acoustic model:
	// * `pending`: The model was created but is waiting either for valid training data to be added or for the service to
	// finish analyzing added data.
	// * `ready`: The model contains valid data and is ready to be trained. If the model contains a mix of valid and
	// invalid resources, you need to set the `strict` parameter to `false` for the training to proceed.
	// * `training`: The model is currently being trained.
	// * `available`: The model is trained and ready to use.
	// * `upgrading`: The model is currently being upgraded.
	// * `failed`: Training of the model failed.
	Status *string `json:"status,omitempty"`

	// A percentage that indicates the progress of the custom acoustic model's current training. A value of `100` means
	// that the model is fully trained. **Note:** The `progress` field does not currently reflect the progress of the
	// training. The field changes from `0` to `100` when training is complete.
	Progress *int64 `json:"progress,omitempty"`

	// If the request included unknown parameters, the following message: `Unexpected query parameter(s) ['parameters']
	// detected`, where `parameters` is a list that includes a quoted string for each unknown parameter.
	Warnings *string `json:"warnings,omitempty"`
}

// Constants associated with the AcousticModel.Status property.
// The current status of the custom acoustic model:
// * `pending`: The model was created but is waiting either for valid training data to be added or for the service to
// finish analyzing added data.
// * `ready`: The model contains valid data and is ready to be trained. If the model contains a mix of valid and invalid
// resources, you need to set the `strict` parameter to `false` for the training to proceed.
// * `training`: The model is currently being trained.
// * `available`: The model is trained and ready to use.
// * `upgrading`: The model is currently being upgraded.
// * `failed`: Training of the model failed.
const (
	AcousticModelStatusAvailableConst = "available"
	AcousticModelStatusFailedConst    = "failed"
	AcousticModelStatusPendingConst   = "pending"
	AcousticModelStatusReadyConst     = "ready"
	AcousticModelStatusTrainingConst  = "training"
	AcousticModelStatusUpgradingConst = "upgrading"
)

// UnmarshalAcousticModel unmarshals an instance of AcousticModel from the specified map of raw messages.
func UnmarshalAcousticModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AcousticModel)
	err = core.UnmarshalPrimitive(m, "customization_id", &obj.CustomizationID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created", &obj.Created)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated", &obj.Updated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "language", &obj.Language)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "versions", &obj.Versions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "owner", &obj.Owner)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "base_model_name", &obj.BaseModelName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "progress", &obj.Progress)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "warnings", &obj.Warnings)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AcousticModels : Information about existing custom acoustic models.
type AcousticModels struct {
	// An array of `AcousticModel` objects that provides information about each available custom acoustic model. The array
	// is empty if the requesting credentials own no custom acoustic models (if no language is specified) or own no custom
	// acoustic models for the specified language.
	Customizations []AcousticModel `json:"customizations" validate:"required"`
}

// UnmarshalAcousticModels unmarshals an instance of AcousticModels from the specified map of raw messages.
func UnmarshalAcousticModels(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AcousticModels)
	err = core.UnmarshalModel(m, "customizations", &obj.Customizations, UnmarshalAcousticModel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AddAudioOptions : The AddAudio options.
type AddAudioOptions struct {
	// The customization ID (GUID) of the custom acoustic model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// The name of the new audio resource for the custom acoustic model. Use a localized name that matches the language of
	// the custom model and reflects the contents of the resource.
	// * Include a maximum of 128 characters in the name.
	// * Do not use characters that need to be URL-encoded. For example, do not use spaces, slashes, backslashes, colons,
	// ampersands, double quotes, plus signs, equals signs, questions marks, and so on in the name. (The service does not
	// prevent the use of these characters. But because they must be URL-encoded wherever used, their use is strongly
	// discouraged.)
	// * Do not use the name of an audio resource that has already been added to the custom model.
	AudioName *string `json:"audio_name" validate:"required,ne="`

	// The audio resource that is to be added to the custom acoustic model, an individual audio file or an archive file.
	//
	// With the `curl` command, use the `--data-binary` option to upload the file for the request.
	AudioResource io.ReadCloser `json:"audio_resource" validate:"required"`

	// For an audio-type resource, the format (MIME type) of the audio. For more information, see **Content types for
	// audio-type resources** in the method description.
	//
	// For an archive-type resource, the media type of the archive file. For more information, see **Content types for
	// archive-type resources** in the method description.
	ContentType *string `json:"Content-Type,omitempty"`

	// _For an archive-type resource_, specify the format of the audio files that are contained in the archive file if they
	// are of type `audio/alaw`, `audio/basic`, `audio/l16`, or `audio/mulaw`. Include the `rate`, `channels`, and
	// `endianness` parameters where necessary. In this case, all audio files that are contained in the archive file must
	// be of the indicated type.
	//
	// For all other audio formats, you can omit the header. In this case, the audio files can be of multiple types as long
	// as they are not of the types listed in the previous paragraph.
	//
	// The parameter accepts all of the audio formats that are supported for use with speech recognition. For more
	// information, see **Content types for audio-type resources** in the method description.
	//
	// _For an audio-type resource_, omit the header.
	ContainedContentType *string `json:"Contained-Content-Type,omitempty"`

	// If `true`, the specified audio resource overwrites an existing audio resource with the same name. If `false`, the
	// request fails if an audio resource with the same name already exists. The parameter has no effect if an audio
	// resource with the same name does not already exist.
	AllowOverwrite *bool `json:"allow_overwrite,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the AddAudioOptions.ContainedContentType property.
// _For an archive-type resource_, specify the format of the audio files that are contained in the archive file if they
// are of type `audio/alaw`, `audio/basic`, `audio/l16`, or `audio/mulaw`. Include the `rate`, `channels`, and
// `endianness` parameters where necessary. In this case, all audio files that are contained in the archive file must be
// of the indicated type.
//
// For all other audio formats, you can omit the header. In this case, the audio files can be of multiple types as long
// as they are not of the types listed in the previous paragraph.
//
// The parameter accepts all of the audio formats that are supported for use with speech recognition. For more
// information, see **Content types for audio-type resources** in the method description.
//
// _For an audio-type resource_, omit the header.
const (
	AddAudioOptionsContainedContentTypeAudioAlawConst             = "audio/alaw"
	AddAudioOptionsContainedContentTypeAudioBasicConst            = "audio/basic"
	AddAudioOptionsContainedContentTypeAudioFlacConst             = "audio/flac"
	AddAudioOptionsContainedContentTypeAudioG729Const             = "audio/g729"
	AddAudioOptionsContainedContentTypeAudioL16Const              = "audio/l16"
	AddAudioOptionsContainedContentTypeAudioMp3Const              = "audio/mp3"
	AddAudioOptionsContainedContentTypeAudioMpegConst             = "audio/mpeg"
	AddAudioOptionsContainedContentTypeAudioMulawConst            = "audio/mulaw"
	AddAudioOptionsContainedContentTypeAudioOggConst              = "audio/ogg"
	AddAudioOptionsContainedContentTypeAudioOggCodecsOpusConst    = "audio/ogg;codecs=opus"
	AddAudioOptionsContainedContentTypeAudioOggCodecsVorbisConst  = "audio/ogg;codecs=vorbis"
	AddAudioOptionsContainedContentTypeAudioWavConst              = "audio/wav"
	AddAudioOptionsContainedContentTypeAudioWebmConst             = "audio/webm"
	AddAudioOptionsContainedContentTypeAudioWebmCodecsOpusConst   = "audio/webm;codecs=opus"
	AddAudioOptionsContainedContentTypeAudioWebmCodecsVorbisConst = "audio/webm;codecs=vorbis"
)

// NewAddAudioOptions : Instantiate AddAudioOptions
func (*SpeechToTextV1) NewAddAudioOptions(customizationID string, audioName string, audioResource io.ReadCloser) *AddAudioOptions {
	return &AddAudioOptions{
		CustomizationID: core.StringPtr(customizationID),
		AudioName:       core.StringPtr(audioName),
		AudioResource:   audioResource,
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *AddAudioOptions) SetCustomizationID(customizationID string) *AddAudioOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetAudioName : Allow user to set AudioName
func (_options *AddAudioOptions) SetAudioName(audioName string) *AddAudioOptions {
	_options.AudioName = core.StringPtr(audioName)
	return _options
}

// SetAudioResource : Allow user to set AudioResource
func (_options *AddAudioOptions) SetAudioResource(audioResource io.ReadCloser) *AddAudioOptions {
	_options.AudioResource = audioResource
	return _options
}

// SetContentType : Allow user to set ContentType
func (_options *AddAudioOptions) SetContentType(contentType string) *AddAudioOptions {
	_options.ContentType = core.StringPtr(contentType)
	return _options
}

// SetContainedContentType : Allow user to set ContainedContentType
func (_options *AddAudioOptions) SetContainedContentType(containedContentType string) *AddAudioOptions {
	_options.ContainedContentType = core.StringPtr(containedContentType)
	return _options
}

// SetAllowOverwrite : Allow user to set AllowOverwrite
func (_options *AddAudioOptions) SetAllowOverwrite(allowOverwrite bool) *AddAudioOptions {
	_options.AllowOverwrite = core.BoolPtr(allowOverwrite)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *AddAudioOptions) SetHeaders(param map[string]string) *AddAudioOptions {
	options.Headers = param
	return options
}

// AddCorpusOptions : The AddCorpus options.
type AddCorpusOptions struct {
	// The customization ID (GUID) of the custom language model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// The name of the new corpus for the custom language model. Use a localized name that matches the language of the
	// custom model and reflects the contents of the corpus.
	// * Include a maximum of 128 characters in the name.
	// * Do not use characters that need to be URL-encoded. For example, do not use spaces, slashes, backslashes, colons,
	// ampersands, double quotes, plus signs, equals signs, questions marks, and so on in the name. (The service does not
	// prevent the use of these characters. But because they must be URL-encoded wherever used, their use is strongly
	// discouraged.)
	// * Do not use the name of an existing corpus or grammar that is already defined for the custom model.
	// * Do not use the name `user`, which is reserved by the service to denote custom words that are added or modified by
	// the user.
	// * Do not use the name `base_lm` or `default_lm`. Both names are reserved for future use by the service.
	CorpusName *string `json:"corpus_name" validate:"required,ne="`

	// A plain text file that contains the training data for the corpus. Encode the file in UTF-8 if it contains non-ASCII
	// characters; the service assumes UTF-8 encoding if it encounters non-ASCII characters.
	//
	// Make sure that you know the character encoding of the file. You must use that same encoding when working with the
	// words in the custom language model. For more information, see [Character encoding for custom
	// words](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-manageWords#charEncoding).
	//
	// With the `curl` command, use the `--data-binary` option to upload the file for the request.
	CorpusFile io.ReadCloser `json:"corpus_file" validate:"required"`

	// If `true`, the specified corpus overwrites an existing corpus with the same name. If `false`, the request fails if a
	// corpus with the same name already exists. The parameter has no effect if a corpus with the same name does not
	// already exist.
	AllowOverwrite *bool `json:"allow_overwrite,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAddCorpusOptions : Instantiate AddCorpusOptions
func (*SpeechToTextV1) NewAddCorpusOptions(customizationID string, corpusName string, corpusFile io.ReadCloser) *AddCorpusOptions {
	return &AddCorpusOptions{
		CustomizationID: core.StringPtr(customizationID),
		CorpusName:      core.StringPtr(corpusName),
		CorpusFile:      corpusFile,
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *AddCorpusOptions) SetCustomizationID(customizationID string) *AddCorpusOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetCorpusName : Allow user to set CorpusName
func (_options *AddCorpusOptions) SetCorpusName(corpusName string) *AddCorpusOptions {
	_options.CorpusName = core.StringPtr(corpusName)
	return _options
}

// SetCorpusFile : Allow user to set CorpusFile
func (_options *AddCorpusOptions) SetCorpusFile(corpusFile io.ReadCloser) *AddCorpusOptions {
	_options.CorpusFile = corpusFile
	return _options
}

// SetAllowOverwrite : Allow user to set AllowOverwrite
func (_options *AddCorpusOptions) SetAllowOverwrite(allowOverwrite bool) *AddCorpusOptions {
	_options.AllowOverwrite = core.BoolPtr(allowOverwrite)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *AddCorpusOptions) SetHeaders(param map[string]string) *AddCorpusOptions {
	options.Headers = param
	return options
}

// AddGrammarOptions : The AddGrammar options.
type AddGrammarOptions struct {
	// The customization ID (GUID) of the custom language model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// The name of the new grammar for the custom language model. Use a localized name that matches the language of the
	// custom model and reflects the contents of the grammar.
	// * Include a maximum of 128 characters in the name.
	// * Do not use characters that need to be URL-encoded. For example, do not use spaces, slashes, backslashes, colons,
	// ampersands, double quotes, plus signs, equals signs, questions marks, and so on in the name. (The service does not
	// prevent the use of these characters. But because they must be URL-encoded wherever used, their use is strongly
	// discouraged.)
	// * Do not use the name of an existing grammar or corpus that is already defined for the custom model.
	// * Do not use the name `user`, which is reserved by the service to denote custom words that are added or modified by
	// the user.
	// * Do not use the name `base_lm` or `default_lm`. Both names are reserved for future use by the service.
	GrammarName *string `json:"grammar_name" validate:"required,ne="`

	// A plain text file that contains the grammar in the format specified by the `Content-Type` header. Encode the file in
	// UTF-8 (ASCII is a subset of UTF-8). Using any other encoding can lead to issues when compiling the grammar or to
	// unexpected results in decoding. The service ignores an encoding that is specified in the header of the grammar.
	//
	// With the `curl` command, use the `--data-binary` option to upload the file for the request.
	GrammarFile io.ReadCloser `json:"grammar_file" validate:"required"`

	// The format (MIME type) of the grammar file:
	// * `application/srgs` for Augmented Backus-Naur Form (ABNF), which uses a plain-text representation that is similar
	// to traditional BNF grammars.
	// * `application/srgs+xml` for XML Form, which uses XML elements to represent the grammar.
	ContentType *string `json:"Content-Type" validate:"required"`

	// If `true`, the specified grammar overwrites an existing grammar with the same name. If `false`, the request fails if
	// a grammar with the same name already exists. The parameter has no effect if a grammar with the same name does not
	// already exist.
	AllowOverwrite *bool `json:"allow_overwrite,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAddGrammarOptions : Instantiate AddGrammarOptions
func (*SpeechToTextV1) NewAddGrammarOptions(customizationID string, grammarName string, grammarFile io.ReadCloser, contentType string) *AddGrammarOptions {
	return &AddGrammarOptions{
		CustomizationID: core.StringPtr(customizationID),
		GrammarName:     core.StringPtr(grammarName),
		GrammarFile:     grammarFile,
		ContentType:     core.StringPtr(contentType),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *AddGrammarOptions) SetCustomizationID(customizationID string) *AddGrammarOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetGrammarName : Allow user to set GrammarName
func (_options *AddGrammarOptions) SetGrammarName(grammarName string) *AddGrammarOptions {
	_options.GrammarName = core.StringPtr(grammarName)
	return _options
}

// SetGrammarFile : Allow user to set GrammarFile
func (_options *AddGrammarOptions) SetGrammarFile(grammarFile io.ReadCloser) *AddGrammarOptions {
	_options.GrammarFile = grammarFile
	return _options
}

// SetContentType : Allow user to set ContentType
func (_options *AddGrammarOptions) SetContentType(contentType string) *AddGrammarOptions {
	_options.ContentType = core.StringPtr(contentType)
	return _options
}

// SetAllowOverwrite : Allow user to set AllowOverwrite
func (_options *AddGrammarOptions) SetAllowOverwrite(allowOverwrite bool) *AddGrammarOptions {
	_options.AllowOverwrite = core.BoolPtr(allowOverwrite)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *AddGrammarOptions) SetHeaders(param map[string]string) *AddGrammarOptions {
	options.Headers = param
	return options
}

// AddWordOptions : The AddWord options.
type AddWordOptions struct {
	// The customization ID (GUID) of the custom language model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// The custom word that is to be added to or updated in the custom language model. Do not include spaces in the word.
	// Use a `-` (dash) or `_` (underscore) to connect the tokens of compound words. URL-encode the word if it includes
	// non-ASCII characters. For more information, see [Character
	// encoding](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-corporaWords#charEncoding).
	WordName *string `json:"word_name" validate:"required,ne="`

	// For the [Add custom words](#addwords) method, you must specify the custom word that is to be added to or updated in
	// the custom model. Do not include spaces in the word. Use a `-` (dash) or `_` (underscore) to connect the tokens of
	// compound words.
	//
	// Omit this parameter for the [Add a custom word](#addword) method.
	Word *string `json:"word,omitempty"`

	// _For a custom model that is based on a previous-generation model_, an array of sounds-like pronunciations for the
	// custom word. Specify how words that are difficult to pronounce, foreign words, acronyms, and so on can be pronounced
	// by users.
	// * For a word that is not in the service's base vocabulary, omit the parameter to have the service automatically
	// generate a sounds-like pronunciation for the word.
	// * For a word that is in the service's base vocabulary, use the parameter to specify additional pronunciations for
	// the word. You cannot override the default pronunciation of a word; pronunciations you add augment the pronunciation
	// from the base vocabulary.
	//
	// A word can have at most five sounds-like pronunciations. A pronunciation can include at most 40 characters not
	// including spaces.
	//
	// _For a custom model that is based on a next-generation model_, omit this field. Custom models based on
	// next-generation models do not support the `sounds_like` field. The service ignores the field.
	SoundsLike []string `json:"sounds_like,omitempty"`

	// An alternative spelling for the custom word when it appears in a transcript. Use the parameter when you want the
	// word to have a spelling that is different from its usual representation or from its spelling in corpora training
	// data.
	DisplayAs *string `json:"display_as,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAddWordOptions : Instantiate AddWordOptions
func (*SpeechToTextV1) NewAddWordOptions(customizationID string, wordName string) *AddWordOptions {
	return &AddWordOptions{
		CustomizationID: core.StringPtr(customizationID),
		WordName:        core.StringPtr(wordName),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *AddWordOptions) SetCustomizationID(customizationID string) *AddWordOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetWordName : Allow user to set WordName
func (_options *AddWordOptions) SetWordName(wordName string) *AddWordOptions {
	_options.WordName = core.StringPtr(wordName)
	return _options
}

// SetWord : Allow user to set Word
func (_options *AddWordOptions) SetWord(word string) *AddWordOptions {
	_options.Word = core.StringPtr(word)
	return _options
}

// SetSoundsLike : Allow user to set SoundsLike
func (_options *AddWordOptions) SetSoundsLike(soundsLike []string) *AddWordOptions {
	_options.SoundsLike = soundsLike
	return _options
}

// SetDisplayAs : Allow user to set DisplayAs
func (_options *AddWordOptions) SetDisplayAs(displayAs string) *AddWordOptions {
	_options.DisplayAs = core.StringPtr(displayAs)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *AddWordOptions) SetHeaders(param map[string]string) *AddWordOptions {
	options.Headers = param
	return options
}

// AddWordsOptions : The AddWords options.
type AddWordsOptions struct {
	// The customization ID (GUID) of the custom language model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// An array of `CustomWord` objects that provides information about each custom word that is to be added to or updated
	// in the custom language model.
	Words []CustomWord `json:"words" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAddWordsOptions : Instantiate AddWordsOptions
func (*SpeechToTextV1) NewAddWordsOptions(customizationID string, words []CustomWord) *AddWordsOptions {
	return &AddWordsOptions{
		CustomizationID: core.StringPtr(customizationID),
		Words:           words,
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *AddWordsOptions) SetCustomizationID(customizationID string) *AddWordsOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetWords : Allow user to set Words
func (_options *AddWordsOptions) SetWords(words []CustomWord) *AddWordsOptions {
	_options.Words = words
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *AddWordsOptions) SetHeaders(param map[string]string) *AddWordsOptions {
	options.Headers = param
	return options
}

// AudioDetails : Information about an audio resource from a custom acoustic model.
type AudioDetails struct {
	// The type of the audio resource:
	// * `audio` for an individual audio file
	// * `archive` for an archive (**.zip** or **.tar.gz**) file that contains audio files
	// * `undetermined` for a resource that the service cannot validate (for example, if the user mistakenly passes a file
	// that does not contain audio, such as a JPEG file).
	Type *string `json:"type,omitempty"`

	// _For an audio-type resource_, the codec in which the audio is encoded. Omitted for an archive-type resource.
	Codec *string `json:"codec,omitempty"`

	// _For an audio-type resource_, the sampling rate of the audio in Hertz (samples per second). Omitted for an
	// archive-type resource.
	Frequency *int64 `json:"frequency,omitempty"`

	// _For an archive-type resource_, the format of the compressed archive:
	// * `zip` for a **.zip** file
	// * `gzip` for a **.tar.gz** file
	//
	// Omitted for an audio-type resource.
	Compression *string `json:"compression,omitempty"`
}

// Constants associated with the AudioDetails.Type property.
// The type of the audio resource:
// * `audio` for an individual audio file
// * `archive` for an archive (**.zip** or **.tar.gz**) file that contains audio files
// * `undetermined` for a resource that the service cannot validate (for example, if the user mistakenly passes a file
// that does not contain audio, such as a JPEG file).
const (
	AudioDetailsTypeArchiveConst      = "archive"
	AudioDetailsTypeAudioConst        = "audio"
	AudioDetailsTypeUndeterminedConst = "undetermined"
)

// Constants associated with the AudioDetails.Compression property.
// _For an archive-type resource_, the format of the compressed archive:
// * `zip` for a **.zip** file
// * `gzip` for a **.tar.gz** file
//
// Omitted for an audio-type resource.
const (
	AudioDetailsCompressionGzipConst = "gzip"
	AudioDetailsCompressionZipConst  = "zip"
)

// UnmarshalAudioDetails unmarshals an instance of AudioDetails from the specified map of raw messages.
func UnmarshalAudioDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AudioDetails)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "codec", &obj.Codec)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "frequency", &obj.Frequency)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "compression", &obj.Compression)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AudioListing : Information about an audio resource from a custom acoustic model.
type AudioListing struct {
	// _For an audio-type resource_, the total seconds of audio in the resource. Omitted for an archive-type resource.
	Duration *int64 `json:"duration,omitempty"`

	// _For an audio-type resource_, the user-specified name of the resource. Omitted for an archive-type resource.
	Name *string `json:"name,omitempty"`

	// _For an audio-type resource_, an `AudioDetails` object that provides detailed information about the resource. The
	// object is empty until the service finishes processing the audio. Omitted for an archive-type resource.
	Details *AudioDetails `json:"details,omitempty"`

	// _For an audio-type resource_, the status of the resource:
	// * `ok`: The service successfully analyzed the audio data. The data can be used to train the custom model.
	// * `being_processed`: The service is still analyzing the audio data. The service cannot accept requests to add new
	// audio resources or to train the custom model until its analysis is complete.
	// * `invalid`: The audio data is not valid for training the custom model (possibly because it has the wrong format or
	// sampling rate, or because it is corrupted).
	//
	// Omitted for an archive-type resource.
	Status *string `json:"status,omitempty"`

	// _For an archive-type resource_, an object of type `AudioResource` that provides information about the resource.
	// Omitted for an audio-type resource.
	Container *AudioResource `json:"container,omitempty"`

	// _For an archive-type resource_, an array of `AudioResource` objects that provides information about the audio-type
	// resources that are contained in the resource. Omitted for an audio-type resource.
	Audio []AudioResource `json:"audio,omitempty"`
}

// Constants associated with the AudioListing.Status property.
// _For an audio-type resource_, the status of the resource:
// * `ok`: The service successfully analyzed the audio data. The data can be used to train the custom model.
// * `being_processed`: The service is still analyzing the audio data. The service cannot accept requests to add new
// audio resources or to train the custom model until its analysis is complete.
// * `invalid`: The audio data is not valid for training the custom model (possibly because it has the wrong format or
// sampling rate, or because it is corrupted).
//
// Omitted for an archive-type resource.
const (
	AudioListingStatusBeingProcessedConst = "being_processed"
	AudioListingStatusInvalidConst        = "invalid"
	AudioListingStatusOkConst             = "ok"
)

// UnmarshalAudioListing unmarshals an instance of AudioListing from the specified map of raw messages.
func UnmarshalAudioListing(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AudioListing)
	err = core.UnmarshalPrimitive(m, "duration", &obj.Duration)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "details", &obj.Details, UnmarshalAudioDetails)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "container", &obj.Container, UnmarshalAudioResource)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "audio", &obj.Audio, UnmarshalAudioResource)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AudioMetrics : If audio metrics are requested, information about the signal characteristics of the input audio.
type AudioMetrics struct {
	// The interval in seconds (typically 0.1 seconds) at which the service calculated the audio metrics. In other words,
	// how often the service calculated the metrics. A single unit in each histogram (see the `AudioMetricsHistogramBin`
	// object) is calculated based on a `sampling_interval` length of audio.
	SamplingInterval *float32 `json:"sampling_interval" validate:"required"`

	// Detailed information about the signal characteristics of the input audio.
	Accumulated *AudioMetricsDetails `json:"accumulated" validate:"required"`
}

// UnmarshalAudioMetrics unmarshals an instance of AudioMetrics from the specified map of raw messages.
func UnmarshalAudioMetrics(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AudioMetrics)
	err = core.UnmarshalPrimitive(m, "sampling_interval", &obj.SamplingInterval)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "accumulated", &obj.Accumulated, UnmarshalAudioMetricsDetails)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AudioMetricsDetails : Detailed information about the signal characteristics of the input audio.
type AudioMetricsDetails struct {
	// If `true`, indicates the end of the audio stream, meaning that transcription is complete. Currently, the field is
	// always `true`. The service returns metrics just once per audio stream. The results provide aggregated audio metrics
	// that pertain to the complete audio stream.
	Final *bool `json:"final" validate:"required"`

	// The end time in seconds of the block of audio to which the metrics apply.
	EndTime *float32 `json:"end_time" validate:"required"`

	// The signal-to-noise ratio (SNR) for the audio signal. The value indicates the ratio of speech to noise in the audio.
	// A valid value lies in the range of 0 to 100 decibels (dB). The service omits the field if it cannot compute the SNR
	// for the audio.
	SignalToNoiseRatio *float32 `json:"signal_to_noise_ratio,omitempty"`

	// The ratio of speech to non-speech segments in the audio signal. The value lies in the range of 0.0 to 1.0.
	SpeechRatio *float32 `json:"speech_ratio" validate:"required"`

	// The probability that the audio signal is missing the upper half of its frequency content.
	// * A value close to 1.0 typically indicates artificially up-sampled audio, which negatively impacts the accuracy of
	// the transcription results.
	// * A value at or near 0.0 indicates that the audio signal is good and has a full spectrum.
	// * A value around 0.5 means that detection of the frequency content is unreliable or not available.
	HighFrequencyLoss *float32 `json:"high_frequency_loss" validate:"required"`

	// An array of `AudioMetricsHistogramBin` objects that defines a histogram of the cumulative direct current (DC)
	// component of the audio signal.
	DirectCurrentOffset []AudioMetricsHistogramBin `json:"direct_current_offset" validate:"required"`

	// An array of `AudioMetricsHistogramBin` objects that defines a histogram of the clipping rate for the audio segments.
	// The clipping rate is defined as the fraction of samples in the segment that reach the maximum or minimum value that
	// is offered by the audio quantization range. The service auto-detects either a 16-bit Pulse-Code Modulation(PCM)
	// audio range (-32768 to +32767) or a unit range (-1.0 to +1.0). The clipping rate is between 0.0 and 1.0, with higher
	// values indicating possible degradation of speech recognition.
	ClippingRate []AudioMetricsHistogramBin `json:"clipping_rate" validate:"required"`

	// An array of `AudioMetricsHistogramBin` objects that defines a histogram of the signal level in segments of the audio
	// that contain speech. The signal level is computed as the Root-Mean-Square (RMS) value in a decibel (dB) scale
	// normalized to the range 0.0 (minimum level) to 1.0 (maximum level).
	SpeechLevel []AudioMetricsHistogramBin `json:"speech_level" validate:"required"`

	// An array of `AudioMetricsHistogramBin` objects that defines a histogram of the signal level in segments of the audio
	// that do not contain speech. The signal level is computed as the Root-Mean-Square (RMS) value in a decibel (dB) scale
	// normalized to the range 0.0 (minimum level) to 1.0 (maximum level).
	NonSpeechLevel []AudioMetricsHistogramBin `json:"non_speech_level" validate:"required"`
}

// UnmarshalAudioMetricsDetails unmarshals an instance of AudioMetricsDetails from the specified map of raw messages.
func UnmarshalAudioMetricsDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AudioMetricsDetails)
	err = core.UnmarshalPrimitive(m, "final", &obj.Final)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "end_time", &obj.EndTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "signal_to_noise_ratio", &obj.SignalToNoiseRatio)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "speech_ratio", &obj.SpeechRatio)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "high_frequency_loss", &obj.HighFrequencyLoss)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "direct_current_offset", &obj.DirectCurrentOffset, UnmarshalAudioMetricsHistogramBin)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "clipping_rate", &obj.ClippingRate, UnmarshalAudioMetricsHistogramBin)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "speech_level", &obj.SpeechLevel, UnmarshalAudioMetricsHistogramBin)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "non_speech_level", &obj.NonSpeechLevel, UnmarshalAudioMetricsHistogramBin)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AudioMetricsHistogramBin : A bin with defined boundaries that indicates the number of values in a range of signal characteristics for a
// histogram. The first and last bins of a histogram are the boundary bins. They cover the intervals between negative
// infinity and the first boundary, and between the last boundary and positive infinity, respectively.
type AudioMetricsHistogramBin struct {
	// The lower boundary of the bin in the histogram.
	Begin *float32 `json:"begin" validate:"required"`

	// The upper boundary of the bin in the histogram.
	End *float32 `json:"end" validate:"required"`

	// The number of values in the bin of the histogram.
	Count *int64 `json:"count" validate:"required"`
}

// UnmarshalAudioMetricsHistogramBin unmarshals an instance of AudioMetricsHistogramBin from the specified map of raw messages.
func UnmarshalAudioMetricsHistogramBin(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AudioMetricsHistogramBin)
	err = core.UnmarshalPrimitive(m, "begin", &obj.Begin)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "end", &obj.End)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "count", &obj.Count)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AudioResource : Information about an audio resource from a custom acoustic model.
type AudioResource struct {
	// The total seconds of audio in the audio resource.
	Duration *int64 `json:"duration" validate:"required"`

	// _For an archive-type resource_, the user-specified name of the resource.
	//
	// _For an audio-type resource_, the user-specified name of the resource or the name of the audio file that the user
	// added for the resource. The value depends on the method that is called.
	Name *string `json:"name" validate:"required"`

	// An `AudioDetails` object that provides detailed information about the audio resource. The object is empty until the
	// service finishes processing the audio.
	Details *AudioDetails `json:"details" validate:"required"`

	// The status of the audio resource:
	// * `ok`: The service successfully analyzed the audio data. The data can be used to train the custom model.
	// * `being_processed`: The service is still analyzing the audio data. The service cannot accept requests to add new
	// audio resources or to train the custom model until its analysis is complete.
	// * `invalid`: The audio data is not valid for training the custom model (possibly because it has the wrong format or
	// sampling rate, or because it is corrupted). For an archive file, the entire archive is invalid if any of its audio
	// files are invalid.
	Status *string `json:"status" validate:"required"`
}

// Constants associated with the AudioResource.Status property.
// The status of the audio resource:
// * `ok`: The service successfully analyzed the audio data. The data can be used to train the custom model.
// * `being_processed`: The service is still analyzing the audio data. The service cannot accept requests to add new
// audio resources or to train the custom model until its analysis is complete.
// * `invalid`: The audio data is not valid for training the custom model (possibly because it has the wrong format or
// sampling rate, or because it is corrupted). For an archive file, the entire archive is invalid if any of its audio
// files are invalid.
const (
	AudioResourceStatusBeingProcessedConst = "being_processed"
	AudioResourceStatusInvalidConst        = "invalid"
	AudioResourceStatusOkConst             = "ok"
)

// UnmarshalAudioResource unmarshals an instance of AudioResource from the specified map of raw messages.
func UnmarshalAudioResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AudioResource)
	err = core.UnmarshalPrimitive(m, "duration", &obj.Duration)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "details", &obj.Details, UnmarshalAudioDetails)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AudioResources : Information about the audio resources from a custom acoustic model.
type AudioResources struct {
	// The total minutes of accumulated audio summed over all of the valid audio resources for the custom acoustic model.
	// You can use this value to determine whether the custom model has too little or too much audio to begin training.
	TotalMinutesOfAudio *float64 `json:"total_minutes_of_audio" validate:"required"`

	// An array of `AudioResource` objects that provides information about the audio resources of the custom acoustic
	// model. The array is empty if the custom model has no audio resources.
	Audio []AudioResource `json:"audio" validate:"required"`
}

// UnmarshalAudioResources unmarshals an instance of AudioResources from the specified map of raw messages.
func UnmarshalAudioResources(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AudioResources)
	err = core.UnmarshalPrimitive(m, "total_minutes_of_audio", &obj.TotalMinutesOfAudio)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "audio", &obj.Audio, UnmarshalAudioResource)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CheckJobOptions : The CheckJob options.
type CheckJobOptions struct {
	// The identifier of the asynchronous job that is to be used for the request. You must make the request with
	// credentials for the instance of the service that owns the job.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCheckJobOptions : Instantiate CheckJobOptions
func (*SpeechToTextV1) NewCheckJobOptions(id string) *CheckJobOptions {
	return &CheckJobOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *CheckJobOptions) SetID(id string) *CheckJobOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CheckJobOptions) SetHeaders(param map[string]string) *CheckJobOptions {
	options.Headers = param
	return options
}

// CheckJobsOptions : The CheckJobs options.
type CheckJobsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCheckJobsOptions : Instantiate CheckJobsOptions
func (*SpeechToTextV1) NewCheckJobsOptions() *CheckJobsOptions {
	return &CheckJobsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *CheckJobsOptions) SetHeaders(param map[string]string) *CheckJobsOptions {
	options.Headers = param
	return options
}

// Corpora : Information about the corpora from a custom language model.
type Corpora struct {
	// An array of `Corpus` objects that provides information about the corpora for the custom model. The array is empty if
	// the custom model has no corpora.
	Corpora []Corpus `json:"corpora" validate:"required"`
}

// UnmarshalCorpora unmarshals an instance of Corpora from the specified map of raw messages.
func UnmarshalCorpora(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Corpora)
	err = core.UnmarshalModel(m, "corpora", &obj.Corpora, UnmarshalCorpus)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Corpus : Information about a corpus from a custom language model.
type Corpus struct {
	// The name of the corpus.
	Name *string `json:"name" validate:"required"`

	// The total number of words in the corpus. The value is `0` while the corpus is being processed.
	TotalWords *int64 `json:"total_words" validate:"required"`

	// _For custom models that are based on previous-generation models_, the number of OOV words extracted from the corpus.
	// The value is `0` while the corpus is being processed.
	//
	// _For custom models that are based on next-generation models_, no OOV words are extracted from corpora, so the value
	// is always `0`.
	OutOfVocabularyWords *int64 `json:"out_of_vocabulary_words" validate:"required"`

	// The status of the corpus:
	// * `analyzed`: The service successfully analyzed the corpus. The custom model can be trained with data from the
	// corpus.
	// * `being_processed`: The service is still analyzing the corpus. The service cannot accept requests to add new
	// resources or to train the custom model.
	// * `undetermined`: The service encountered an error while processing the corpus. The `error` field describes the
	// failure.
	Status *string `json:"status" validate:"required"`

	// If the status of the corpus is `undetermined`, the following message: `Analysis of corpus 'name' failed. Please try
	// adding the corpus again by setting the 'allow_overwrite' flag to 'true'`.
	Error *string `json:"error,omitempty"`
}

// Constants associated with the Corpus.Status property.
// The status of the corpus:
// * `analyzed`: The service successfully analyzed the corpus. The custom model can be trained with data from the
// corpus.
// * `being_processed`: The service is still analyzing the corpus. The service cannot accept requests to add new
// resources or to train the custom model.
// * `undetermined`: The service encountered an error while processing the corpus. The `error` field describes the
// failure.
const (
	CorpusStatusAnalyzedConst       = "analyzed"
	CorpusStatusBeingProcessedConst = "being_processed"
	CorpusStatusUndeterminedConst   = "undetermined"
)

// UnmarshalCorpus unmarshals an instance of Corpus from the specified map of raw messages.
func UnmarshalCorpus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Corpus)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_words", &obj.TotalWords)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "out_of_vocabulary_words", &obj.OutOfVocabularyWords)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
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

// CreateAcousticModelOptions : The CreateAcousticModel options.
type CreateAcousticModelOptions struct {
	// A user-defined name for the new custom acoustic model. Use a name that is unique among all custom acoustic models
	// that you own. Use a localized name that matches the language of the custom model. Use a name that describes the
	// acoustic environment of the custom model, such as `Mobile custom model` or `Noisy car custom model`.
	Name *string `json:"name" validate:"required"`

	// The name of the base language model that is to be customized by the new custom acoustic model. The new custom model
	// can be used only with the base model that it customizes. (**Note:** The model `ar-AR_BroadbandModel` is deprecated;
	// use `ar-MS_BroadbandModel` instead.)
	//
	// To determine whether a base model supports acoustic model customization, refer to [Language support for
	// customization](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-custom-support).
	BaseModelName *string `json:"base_model_name" validate:"required"`

	// A description of the new custom acoustic model. Use a localized description that matches the language of the custom
	// model.
	Description *string `json:"description,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateAcousticModelOptions.BaseModelName property.
// The name of the base language model that is to be customized by the new custom acoustic model. The new custom model
// can be used only with the base model that it customizes. (**Note:** The model `ar-AR_BroadbandModel` is deprecated;
// use `ar-MS_BroadbandModel` instead.)
//
// To determine whether a base model supports acoustic model customization, refer to [Language support for
// customization](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-custom-support).
const (
	CreateAcousticModelOptionsBaseModelNameArArBroadbandmodelConst           = "ar-AR_BroadbandModel"
	CreateAcousticModelOptionsBaseModelNameArMsBroadbandmodelConst           = "ar-MS_BroadbandModel"
	CreateAcousticModelOptionsBaseModelNameDeDeBroadbandmodelConst           = "de-DE_BroadbandModel"
	CreateAcousticModelOptionsBaseModelNameDeDeNarrowbandmodelConst          = "de-DE_NarrowbandModel"
	CreateAcousticModelOptionsBaseModelNameEnAuBroadbandmodelConst           = "en-AU_BroadbandModel"
	CreateAcousticModelOptionsBaseModelNameEnAuNarrowbandmodelConst          = "en-AU_NarrowbandModel"
	CreateAcousticModelOptionsBaseModelNameEnGbBroadbandmodelConst           = "en-GB_BroadbandModel"
	CreateAcousticModelOptionsBaseModelNameEnGbNarrowbandmodelConst          = "en-GB_NarrowbandModel"
	CreateAcousticModelOptionsBaseModelNameEnUsBroadbandmodelConst           = "en-US_BroadbandModel"
	CreateAcousticModelOptionsBaseModelNameEnUsNarrowbandmodelConst          = "en-US_NarrowbandModel"
	CreateAcousticModelOptionsBaseModelNameEnUsShortformNarrowbandmodelConst = "en-US_ShortForm_NarrowbandModel"
	CreateAcousticModelOptionsBaseModelNameEsArBroadbandmodelConst           = "es-AR_BroadbandModel"
	CreateAcousticModelOptionsBaseModelNameEsArNarrowbandmodelConst          = "es-AR_NarrowbandModel"
	CreateAcousticModelOptionsBaseModelNameEsClBroadbandmodelConst           = "es-CL_BroadbandModel"
	CreateAcousticModelOptionsBaseModelNameEsClNarrowbandmodelConst          = "es-CL_NarrowbandModel"
	CreateAcousticModelOptionsBaseModelNameEsCoBroadbandmodelConst           = "es-CO_BroadbandModel"
	CreateAcousticModelOptionsBaseModelNameEsCoNarrowbandmodelConst          = "es-CO_NarrowbandModel"
	CreateAcousticModelOptionsBaseModelNameEsEsBroadbandmodelConst           = "es-ES_BroadbandModel"
	CreateAcousticModelOptionsBaseModelNameEsEsNarrowbandmodelConst          = "es-ES_NarrowbandModel"
	CreateAcousticModelOptionsBaseModelNameEsMxBroadbandmodelConst           = "es-MX_BroadbandModel"
	CreateAcousticModelOptionsBaseModelNameEsMxNarrowbandmodelConst          = "es-MX_NarrowbandModel"
	CreateAcousticModelOptionsBaseModelNameEsPeBroadbandmodelConst           = "es-PE_BroadbandModel"
	CreateAcousticModelOptionsBaseModelNameEsPeNarrowbandmodelConst          = "es-PE_NarrowbandModel"
	CreateAcousticModelOptionsBaseModelNameFrCaBroadbandmodelConst           = "fr-CA_BroadbandModel"
	CreateAcousticModelOptionsBaseModelNameFrCaNarrowbandmodelConst          = "fr-CA_NarrowbandModel"
	CreateAcousticModelOptionsBaseModelNameFrFrBroadbandmodelConst           = "fr-FR_BroadbandModel"
	CreateAcousticModelOptionsBaseModelNameFrFrNarrowbandmodelConst          = "fr-FR_NarrowbandModel"
	CreateAcousticModelOptionsBaseModelNameItItBroadbandmodelConst           = "it-IT_BroadbandModel"
	CreateAcousticModelOptionsBaseModelNameItItNarrowbandmodelConst          = "it-IT_NarrowbandModel"
	CreateAcousticModelOptionsBaseModelNameJaJpBroadbandmodelConst           = "ja-JP_BroadbandModel"
	CreateAcousticModelOptionsBaseModelNameJaJpNarrowbandmodelConst          = "ja-JP_NarrowbandModel"
	CreateAcousticModelOptionsBaseModelNameKoKrBroadbandmodelConst           = "ko-KR_BroadbandModel"
	CreateAcousticModelOptionsBaseModelNameKoKrNarrowbandmodelConst          = "ko-KR_NarrowbandModel"
	CreateAcousticModelOptionsBaseModelNameNlNlBroadbandmodelConst           = "nl-NL_BroadbandModel"
	CreateAcousticModelOptionsBaseModelNameNlNlNarrowbandmodelConst          = "nl-NL_NarrowbandModel"
	CreateAcousticModelOptionsBaseModelNamePtBrBroadbandmodelConst           = "pt-BR_BroadbandModel"
	CreateAcousticModelOptionsBaseModelNamePtBrNarrowbandmodelConst          = "pt-BR_NarrowbandModel"
	CreateAcousticModelOptionsBaseModelNameZhCnBroadbandmodelConst           = "zh-CN_BroadbandModel"
	CreateAcousticModelOptionsBaseModelNameZhCnNarrowbandmodelConst          = "zh-CN_NarrowbandModel"
)

// NewCreateAcousticModelOptions : Instantiate CreateAcousticModelOptions
func (*SpeechToTextV1) NewCreateAcousticModelOptions(name string, baseModelName string) *CreateAcousticModelOptions {
	return &CreateAcousticModelOptions{
		Name:          core.StringPtr(name),
		BaseModelName: core.StringPtr(baseModelName),
	}
}

// SetName : Allow user to set Name
func (_options *CreateAcousticModelOptions) SetName(name string) *CreateAcousticModelOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetBaseModelName : Allow user to set BaseModelName
func (_options *CreateAcousticModelOptions) SetBaseModelName(baseModelName string) *CreateAcousticModelOptions {
	_options.BaseModelName = core.StringPtr(baseModelName)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateAcousticModelOptions) SetDescription(description string) *CreateAcousticModelOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateAcousticModelOptions) SetHeaders(param map[string]string) *CreateAcousticModelOptions {
	options.Headers = param
	return options
}

// CreateJobOptions : The CreateJob options.
type CreateJobOptions struct {
	// The audio to transcribe.
	Audio io.ReadCloser `json:"audio" validate:"required"`

	// The format (MIME type) of the audio. For more information about specifying an audio format, see **Audio formats
	// (content types)** in the method description.
	ContentType *string `json:"Content-Type,omitempty"`

	// The identifier of the model that is to be used for the recognition request. (**Note:** The model
	// `ar-AR_BroadbandModel` is deprecated; use `ar-MS_BroadbandModel` instead.) See [Using a model for speech
	// recognition](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-models-use).
	Model *string `json:"model,omitempty"`

	// A URL to which callback notifications are to be sent. The URL must already be successfully allowlisted by using the
	// [Register a callback](#registercallback) method. You can include the same callback URL with any number of job
	// creation requests. Omit the parameter to poll the service for job completion and results.
	//
	// Use the `user_token` parameter to specify a unique user-specified string with each job to differentiate the callback
	// notifications for the jobs.
	CallbackURL *string `json:"callback_url,omitempty"`

	// If the job includes a callback URL, a comma-separated list of notification events to which to subscribe. Valid
	// events are
	// * `recognitions.started` generates a callback notification when the service begins to process the job.
	// * `recognitions.completed` generates a callback notification when the job is complete. You must use the [Check a
	// job](#checkjob) method to retrieve the results before they time out or are deleted.
	// * `recognitions.completed_with_results` generates a callback notification when the job is complete. The notification
	// includes the results of the request.
	// * `recognitions.failed` generates a callback notification if the service experiences an error while processing the
	// job.
	//
	// The `recognitions.completed` and `recognitions.completed_with_results` events are incompatible. You can specify only
	// of the two events.
	//
	// If the job includes a callback URL, omit the parameter to subscribe to the default events: `recognitions.started`,
	// `recognitions.completed`, and `recognitions.failed`. If the job does not include a callback URL, omit the parameter.
	Events *string `json:"events,omitempty"`

	// If the job includes a callback URL, a user-specified string that the service is to include with each callback
	// notification for the job; the token allows the user to maintain an internal mapping between jobs and notification
	// events. If the job does not include a callback URL, omit the parameter.
	UserToken *string `json:"user_token,omitempty"`

	// The number of minutes for which the results are to be available after the job has finished. If not delivered via a
	// callback, the results must be retrieved within this time. Omit the parameter to use a time to live of one week. The
	// parameter is valid with or without a callback URL.
	ResultsTTL *int64 `json:"results_ttl,omitempty"`

	// The customization ID (GUID) of a custom language model that is to be used with the recognition request. The base
	// model of the specified custom language model must match the model specified with the `model` parameter. You must
	// make the request with credentials for the instance of the service that owns the custom model. By default, no custom
	// language model is used. See [Using a custom language model for speech
	// recognition](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-languageUse).
	//
	// **Note:** Use this parameter instead of the deprecated `customization_id` parameter.
	LanguageCustomizationID *string `json:"language_customization_id,omitempty"`

	// The customization ID (GUID) of a custom acoustic model that is to be used with the recognition request. The base
	// model of the specified custom acoustic model must match the model specified with the `model` parameter. You must
	// make the request with credentials for the instance of the service that owns the custom model. By default, no custom
	// acoustic model is used. See [Using a custom acoustic model for speech
	// recognition](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-acousticUse).
	AcousticCustomizationID *string `json:"acoustic_customization_id,omitempty"`

	// The version of the specified base model that is to be used with the recognition request. Multiple versions of a base
	// model can exist when a model is updated for internal improvements. The parameter is intended primarily for use with
	// custom models that have been upgraded for a new base model. The default value depends on whether the parameter is
	// used with or without a custom model. See [Making speech recognition requests with upgraded custom
	// models](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-custom-upgrade-use#custom-upgrade-use-recognition).
	BaseModelVersion *string `json:"base_model_version,omitempty"`

	// If you specify the customization ID (GUID) of a custom language model with the recognition request, the
	// customization weight tells the service how much weight to give to words from the custom language model compared to
	// those from the base model for the current request.
	//
	// Specify a value between 0.0 and 1.0. Unless a different customization weight was specified for the custom model when
	// it was trained, the default value is 0.3. A customization weight that you specify overrides a weight that was
	// specified when the custom model was trained.
	//
	// The default value yields the best performance in general. Assign a higher value if your audio makes frequent use of
	// OOV words from the custom model. Use caution when setting the weight: a higher value can improve the accuracy of
	// phrases from the custom model's domain, but it can negatively affect performance on non-domain phrases.
	//
	// See [Using customization weight](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-languageUse#weight).
	CustomizationWeight *float64 `json:"customization_weight,omitempty"`

	// The time in seconds after which, if only silence (no speech) is detected in streaming audio, the connection is
	// closed with a 400 error. The parameter is useful for stopping audio submission from a live microphone when a user
	// simply walks away. Use `-1` for infinity. See [Inactivity
	// timeout](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-input#timeouts-inactivity).
	InactivityTimeout *int64 `json:"inactivity_timeout,omitempty"`

	// An array of keyword strings to spot in the audio. Each keyword string can include one or more string tokens.
	// Keywords are spotted only in the final results, not in interim hypotheses. If you specify any keywords, you must
	// also specify a keywords threshold. Omit the parameter or specify an empty array if you do not need to spot keywords.
	//
	//
	// You can spot a maximum of 1000 keywords with a single request. A single keyword can have a maximum length of 1024
	// characters, though the maximum effective length for double-byte languages might be shorter. Keywords are
	// case-insensitive.
	//
	// See [Keyword spotting](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-spotting#keyword-spotting).
	Keywords []string `json:"keywords,omitempty"`

	// A confidence value that is the lower bound for spotting a keyword. A word is considered to match a keyword if its
	// confidence is greater than or equal to the threshold. Specify a probability between 0.0 and 1.0. If you specify a
	// threshold, you must also specify one or more keywords. The service performs no keyword spotting if you omit either
	// parameter. See [Keyword
	// spotting](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-spotting#keyword-spotting).
	KeywordsThreshold *float32 `json:"keywords_threshold,omitempty"`

	// The maximum number of alternative transcripts that the service is to return. By default, the service returns a
	// single transcript. If you specify a value of `0`, the service uses the default value, `1`. See [Maximum
	// alternatives](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-metadata#max-alternatives).
	MaxAlternatives *int64 `json:"max_alternatives,omitempty"`

	// A confidence value that is the lower bound for identifying a hypothesis as a possible word alternative (also known
	// as "Confusion Networks"). An alternative word is considered if its confidence is greater than or equal to the
	// threshold. Specify a probability between 0.0 and 1.0. By default, the service computes no alternative words. See
	// [Word alternatives](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-spotting#word-alternatives).
	WordAlternativesThreshold *float32 `json:"word_alternatives_threshold,omitempty"`

	// If `true`, the service returns a confidence measure in the range of 0.0 to 1.0 for each word. By default, the
	// service returns no word confidence scores. See [Word
	// confidence](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-metadata#word-confidence).
	WordConfidence *bool `json:"word_confidence,omitempty"`

	// If `true`, the service returns time alignment for each word. By default, no timestamps are returned. See [Word
	// timestamps](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-metadata#word-timestamps).
	Timestamps *bool `json:"timestamps,omitempty"`

	// If `true`, the service filters profanity from all output except for keyword results by replacing inappropriate words
	// with a series of asterisks. Set the parameter to `false` to return results with no censoring.
	//
	// **Note:** The parameter can be used with US English and Japanese transcription only. See [Profanity
	// filtering](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-formatting#profanity-filtering).
	ProfanityFilter *bool `json:"profanity_filter,omitempty"`

	// If `true`, the service converts dates, times, series of digits and numbers, phone numbers, currency values, and
	// internet addresses into more readable, conventional representations in the final transcript of a recognition
	// request. For US English, the service also converts certain keyword strings to punctuation symbols. By default, the
	// service performs no smart formatting.
	//
	// **Note:** The parameter can be used with US English, Japanese, and Spanish (all dialects) transcription only.
	//
	// See [Smart formatting](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-formatting#smart-formatting).
	SmartFormatting *bool `json:"smart_formatting,omitempty"`

	// If `true`, the response includes labels that identify which words were spoken by which participants in a
	// multi-person exchange. By default, the service returns no speaker labels. Setting `speaker_labels` to `true` forces
	// the `timestamps` parameter to be `true`, regardless of whether you specify `false` for the parameter.
	// * _For previous-generation models,_ the parameter can be used with Australian English, US English, German, Japanese,
	// Korean, and Spanish (both broadband and narrowband models) and UK English (narrowband model) transcription only.
	// * _For next-generation models,_ the parameter can be used with Czech, English (Australian, Indian, UK, and US),
	// German, Japanese, Korean, and Spanish transcription only.
	//
	// See [Speaker labels](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-speaker-labels).
	SpeakerLabels *bool `json:"speaker_labels,omitempty"`

	// **Deprecated.** Use the `language_customization_id` parameter to specify the customization ID (GUID) of a custom
	// language model that is to be used with the recognition request. Do not specify both parameters with a request.
	CustomizationID *string `json:"customization_id,omitempty"`

	// The name of a grammar that is to be used with the recognition request. If you specify a grammar, you must also use
	// the `language_customization_id` parameter to specify the name of the custom language model for which the grammar is
	// defined. The service recognizes only strings that are recognized by the specified grammar; it does not recognize
	// other custom words from the model's words resource.
	//
	// See [Using a grammar for speech
	// recognition](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-grammarUse).
	GrammarName *string `json:"grammar_name,omitempty"`

	// If `true`, the service redacts, or masks, numeric data from final transcripts. The feature redacts any number that
	// has three or more consecutive digits by replacing each digit with an `X` character. It is intended to redact
	// sensitive numeric data, such as credit card numbers. By default, the service performs no redaction.
	//
	// When you enable redaction, the service automatically enables smart formatting, regardless of whether you explicitly
	// disable that feature. To ensure maximum security, the service also disables keyword spotting (ignores the `keywords`
	// and `keywords_threshold` parameters) and returns only a single final transcript (forces the `max_alternatives`
	// parameter to be `1`).
	//
	// **Note:** The parameter can be used with US English, Japanese, and Korean transcription only.
	//
	// See [Numeric
	// redaction](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-formatting#numeric-redaction).
	Redaction *bool `json:"redaction,omitempty"`

	// If `true`, requests processing metrics about the service's transcription of the input audio. The service returns
	// processing metrics at the interval specified by the `processing_metrics_interval` parameter. It also returns
	// processing metrics for transcription events, for example, for final and interim results. By default, the service
	// returns no processing metrics.
	//
	// See [Processing metrics](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-metrics#processing-metrics).
	ProcessingMetrics *bool `json:"processing_metrics,omitempty"`

	// Specifies the interval in real wall-clock seconds at which the service is to return processing metrics. The
	// parameter is ignored unless the `processing_metrics` parameter is set to `true`.
	//
	// The parameter accepts a minimum value of 0.1 seconds. The level of precision is not restricted, so you can specify
	// values such as 0.25 and 0.125.
	//
	// The service does not impose a maximum value. If you want to receive processing metrics only for transcription events
	// instead of at periodic intervals, set the value to a large number. If the value is larger than the duration of the
	// audio, the service returns processing metrics only for transcription events.
	//
	// See [Processing metrics](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-metrics#processing-metrics).
	ProcessingMetricsInterval *float32 `json:"processing_metrics_interval,omitempty"`

	// If `true`, requests detailed information about the signal characteristics of the input audio. The service returns
	// audio metrics with the final transcription results. By default, the service returns no audio metrics.
	//
	// See [Audio metrics](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-metrics#audio-metrics).
	AudioMetrics *bool `json:"audio_metrics,omitempty"`

	// If `true`, specifies the duration of the pause interval at which the service splits a transcript into multiple final
	// results. If the service detects pauses or extended silence before it reaches the end of the audio stream, its
	// response can include multiple final results. Silence indicates a point at which the speaker pauses between spoken
	// words or phrases.
	//
	// Specify a value for the pause interval in the range of 0.0 to 120.0.
	// * A value greater than 0 specifies the interval that the service is to use for speech recognition.
	// * A value of 0 indicates that the service is to use the default interval. It is equivalent to omitting the
	// parameter.
	//
	// The default pause interval for most languages is 0.8 seconds; the default for Chinese is 0.6 seconds.
	//
	// See [End of phrase silence
	// time](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-parsing#silence-time).
	EndOfPhraseSilenceTime *float64 `json:"end_of_phrase_silence_time,omitempty"`

	// If `true`, directs the service to split the transcript into multiple final results based on semantic features of the
	// input, for example, at the conclusion of meaningful phrases such as sentences. The service bases its understanding
	// of semantic features on the base language model that you use with a request. Custom language models and grammars can
	// also influence how and where the service splits a transcript.
	//
	// By default, the service splits transcripts based solely on the pause interval. If the parameters are used together
	// on the same request, `end_of_phrase_silence_time` has precedence over `split_transcript_at_phrase_end`.
	//
	// See [Split transcript at phrase
	// end](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-parsing#split-transcript).
	SplitTranscriptAtPhraseEnd *bool `json:"split_transcript_at_phrase_end,omitempty"`

	// The sensitivity of speech activity detection that the service is to perform. Use the parameter to suppress word
	// insertions from music, coughing, and other non-speech events. The service biases the audio it passes for speech
	// recognition by evaluating the input audio against prior models of speech and non-speech activity.
	//
	// Specify a value between 0.0 and 1.0:
	// * 0.0 suppresses all audio (no speech is transcribed).
	// * 0.5 (the default) provides a reasonable compromise for the level of sensitivity.
	// * 1.0 suppresses no audio (speech detection sensitivity is disabled).
	//
	// The values increase on a monotonic curve.
	//
	// The parameter is supported with all next-generation models and with most previous-generation models. See [Speech
	// detector
	// sensitivity](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-detection#detection-parameters-sensitivity)
	// and [Language model
	// support](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-detection#detection-support).
	SpeechDetectorSensitivity *float32 `json:"speech_detector_sensitivity,omitempty"`

	// The level to which the service is to suppress background audio based on its volume to prevent it from being
	// transcribed as speech. Use the parameter to suppress side conversations or background noise.
	//
	// Specify a value in the range of 0.0 to 1.0:
	// * 0.0 (the default) provides no suppression (background audio suppression is disabled).
	// * 0.5 provides a reasonable level of audio suppression for general usage.
	// * 1.0 suppresses all audio (no audio is transcribed).
	//
	// The values increase on a monotonic curve.
	//
	// The parameter is supported with all next-generation models and with most previous-generation models. See [Background
	// audio
	// suppression](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-detection#detection-parameters-suppression)
	// and [Language model
	// support](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-detection#detection-support).
	BackgroundAudioSuppression *float32 `json:"background_audio_suppression,omitempty"`

	// The character_insertion_bias parameter controls the service's bias for competing strings of different lengths
	// during speech recognition. With next-generation models, the service parses audio character by character.
	// As it does, it establishes hypotheses of previous character strings to help determine viable next characters.
	// During this process, it collects candidate strings of different lengths.
	//
	// By default, each model uses a default character_insertion_bias of 0.0.
	// This value is optimized to produce the best balance between hypotheses with different numbers of characters.
	// The default is typically adequate for most speech recognition.
	// However, certain use cases might benefit from favoring hypotheses with shorter or longer strings of characters.
	// In such cases, specifying a change from the default can improve speech recognition.
	//
	// You can use the character_insertion_bias parameter to indicate that the service is to favor shorter or longer
	//  strings as it considers subsequent characters for its hypotheses.
	// The value you provide depends on the characteristics of your audio.
	// The range of acceptable values is from -1.0 to 1.0:
	//
	// Negative values cause the service to prefer hypotheses with shorter strings of characters.
	// Positive values cause the service to prefer hypotheses with longer strings of characters.
	// As your value approaches -1.0 or 1.0, the impact of the parameter becomes more pronounced.
	// To determine the most effective value for your scenario, start by setting the value of the parameter
	// to a small increment, such as -0.1, -0.05, 0.05, or 0.1, and assess how the value impacts the transcription results.
	//
	// The parameter is not available for previous-generation models.
	CharacterInsertionBias *float32 `json:"character_insertion_bias,omitempty"`

	// If `true` for next-generation `Multimedia` and `Telephony` models that support low latency, directs the service to
	// produce results even more quickly than it usually does. Next-generation models produce transcription results faster
	// than previous-generation models. The `low_latency` parameter causes the models to produce results even more quickly,
	// though the results might be less accurate when the parameter is used.
	//
	// The parameter is not available for previous-generation `Broadband` and `Narrowband` models. It is available only for
	// some next-generation models. For a list of next-generation models that support low latency, see [Supported
	// next-generation language
	// models](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-models-ng#models-ng-supported).
	// * For more information about the `low_latency` parameter, see [Low
	// latency](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-interim#low-latency).
	LowLatency *bool `json:"low_latency,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateJobOptions.Model property.
// The identifier of the model that is to be used for the recognition request. (**Note:** The model
// `ar-AR_BroadbandModel` is deprecated; use `ar-MS_BroadbandModel` instead.) See [Using a model for speech
// recognition](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-models-use).
const (
	CreateJobOptionsModelArArBroadbandmodelConst           = "ar-AR_BroadbandModel"
	CreateJobOptionsModelArMsBroadbandmodelConst           = "ar-MS_BroadbandModel"
	CreateJobOptionsModelArMsTelephonyConst                = "ar-MS_Telephony"
	CreateJobOptionsModelCsCzTelephonyConst                = "cs-CZ_Telephony"
	CreateJobOptionsModelDeDeBroadbandmodelConst           = "de-DE_BroadbandModel"
	CreateJobOptionsModelDeDeMultimediaConst               = "de-DE_Multimedia"
	CreateJobOptionsModelDeDeNarrowbandmodelConst          = "de-DE_NarrowbandModel"
	CreateJobOptionsModelDeDeTelephonyConst                = "de-DE_Telephony"
	CreateJobOptionsModelEnAuBroadbandmodelConst           = "en-AU_BroadbandModel"
	CreateJobOptionsModelEnAuMultimediaConst               = "en-AU_Multimedia"
	CreateJobOptionsModelEnAuNarrowbandmodelConst          = "en-AU_NarrowbandModel"
	CreateJobOptionsModelEnAuTelephonyConst                = "en-AU_Telephony"
	CreateJobOptionsModelEnGbBroadbandmodelConst           = "en-GB_BroadbandModel"
	CreateJobOptionsModelEnGbMultimediaConst               = "en-GB_Multimedia"
	CreateJobOptionsModelEnGbNarrowbandmodelConst          = "en-GB_NarrowbandModel"
	CreateJobOptionsModelEnGbTelephonyConst                = "en-GB_Telephony"
	CreateJobOptionsModelEnInTelephonyConst                = "en-IN_Telephony"
	CreateJobOptionsModelEnUsBroadbandmodelConst           = "en-US_BroadbandModel"
	CreateJobOptionsModelEnUsMultimediaConst               = "en-US_Multimedia"
	CreateJobOptionsModelEnUsNarrowbandmodelConst          = "en-US_NarrowbandModel"
	CreateJobOptionsModelEnUsShortformNarrowbandmodelConst = "en-US_ShortForm_NarrowbandModel"
	CreateJobOptionsModelEnUsTelephonyConst                = "en-US_Telephony"
	CreateJobOptionsModelEnWwMedicalTelephonyConst         = "en-WW_Medical_Telephony"
	CreateJobOptionsModelEsArBroadbandmodelConst           = "es-AR_BroadbandModel"
	CreateJobOptionsModelEsArNarrowbandmodelConst          = "es-AR_NarrowbandModel"
	CreateJobOptionsModelEsClBroadbandmodelConst           = "es-CL_BroadbandModel"
	CreateJobOptionsModelEsClNarrowbandmodelConst          = "es-CL_NarrowbandModel"
	CreateJobOptionsModelEsCoBroadbandmodelConst           = "es-CO_BroadbandModel"
	CreateJobOptionsModelEsCoNarrowbandmodelConst          = "es-CO_NarrowbandModel"
	CreateJobOptionsModelEsEsBroadbandmodelConst           = "es-ES_BroadbandModel"
	CreateJobOptionsModelEsEsMultimediaConst               = "es-ES_Multimedia"
	CreateJobOptionsModelEsEsNarrowbandmodelConst          = "es-ES_NarrowbandModel"
	CreateJobOptionsModelEsEsTelephonyConst                = "es-ES_Telephony"
	CreateJobOptionsModelEsLaTelephonyConst                = "es-LA_Telephony"
	CreateJobOptionsModelEsMxBroadbandmodelConst           = "es-MX_BroadbandModel"
	CreateJobOptionsModelEsMxNarrowbandmodelConst          = "es-MX_NarrowbandModel"
	CreateJobOptionsModelEsPeBroadbandmodelConst           = "es-PE_BroadbandModel"
	CreateJobOptionsModelEsPeNarrowbandmodelConst          = "es-PE_NarrowbandModel"
	CreateJobOptionsModelFrCaBroadbandmodelConst           = "fr-CA_BroadbandModel"
	CreateJobOptionsModelFrCaNarrowbandmodelConst          = "fr-CA_NarrowbandModel"
	CreateJobOptionsModelFrCaTelephonyConst                = "fr-CA_Telephony"
	CreateJobOptionsModelFrFrBroadbandmodelConst           = "fr-FR_BroadbandModel"
	CreateJobOptionsModelFrFrMultimediaConst               = "fr-FR_Multimedia"
	CreateJobOptionsModelFrFrNarrowbandmodelConst          = "fr-FR_NarrowbandModel"
	CreateJobOptionsModelFrFrTelephonyConst                = "fr-FR_Telephony"
	CreateJobOptionsModelHiInTelephonyConst                = "hi-IN_Telephony"
	CreateJobOptionsModelItItBroadbandmodelConst           = "it-IT_BroadbandModel"
	CreateJobOptionsModelItItNarrowbandmodelConst          = "it-IT_NarrowbandModel"
	CreateJobOptionsModelItItTelephonyConst                = "it-IT_Telephony"
	CreateJobOptionsModelJaJpBroadbandmodelConst           = "ja-JP_BroadbandModel"
	CreateJobOptionsModelJaJpMultimediaConst               = "ja-JP_Multimedia"
	CreateJobOptionsModelJaJpNarrowbandmodelConst          = "ja-JP_NarrowbandModel"
	CreateJobOptionsModelKoKrBroadbandmodelConst           = "ko-KR_BroadbandModel"
	CreateJobOptionsModelKoKrMultimediaConst               = "ko-KR_Multimedia"
	CreateJobOptionsModelKoKrNarrowbandmodelConst          = "ko-KR_NarrowbandModel"
	CreateJobOptionsModelKoKrTelephonyConst                = "ko-KR_Telephony"
	CreateJobOptionsModelNlBeTelephonyConst                = "nl-BE_Telephony"
	CreateJobOptionsModelNlNlBroadbandmodelConst           = "nl-NL_BroadbandModel"
	CreateJobOptionsModelNlNlNarrowbandmodelConst          = "nl-NL_NarrowbandModel"
	CreateJobOptionsModelNlNlTelephonyConst                = "nl-NL_Telephony"
	CreateJobOptionsModelPtBrBroadbandmodelConst           = "pt-BR_BroadbandModel"
	CreateJobOptionsModelPtBrNarrowbandmodelConst          = "pt-BR_NarrowbandModel"
	CreateJobOptionsModelPtBrTelephonyConst                = "pt-BR_Telephony"
	CreateJobOptionsModelZhCnBroadbandmodelConst           = "zh-CN_BroadbandModel"
	CreateJobOptionsModelZhCnNarrowbandmodelConst          = "zh-CN_NarrowbandModel"
	CreateJobOptionsModelZhCnTelephonyConst                = "zh-CN_Telephony"
)

// Constants associated with the CreateJobOptions.Events property.
// If the job includes a callback URL, a comma-separated list of notification events to which to subscribe. Valid events
// are
// * `recognitions.started` generates a callback notification when the service begins to process the job.
// * `recognitions.completed` generates a callback notification when the job is complete. You must use the [Check a
// job](#checkjob) method to retrieve the results before they time out or are deleted.
// * `recognitions.completed_with_results` generates a callback notification when the job is complete. The notification
// includes the results of the request.
// * `recognitions.failed` generates a callback notification if the service experiences an error while processing the
// job.
//
// The `recognitions.completed` and `recognitions.completed_with_results` events are incompatible. You can specify only
// of the two events.
//
// If the job includes a callback URL, omit the parameter to subscribe to the default events: `recognitions.started`,
// `recognitions.completed`, and `recognitions.failed`. If the job does not include a callback URL, omit the parameter.
const (
	CreateJobOptionsEventsRecognitionsCompletedConst            = "recognitions.completed"
	CreateJobOptionsEventsRecognitionsCompletedWithResultsConst = "recognitions.completed_with_results"
	CreateJobOptionsEventsRecognitionsFailedConst               = "recognitions.failed"
	CreateJobOptionsEventsRecognitionsStartedConst              = "recognitions.started"
)

// NewCreateJobOptions : Instantiate CreateJobOptions
func (*SpeechToTextV1) NewCreateJobOptions(audio io.ReadCloser) *CreateJobOptions {
	return &CreateJobOptions{
		Audio: audio,
	}
}

// SetAudio : Allow user to set Audio
func (_options *CreateJobOptions) SetAudio(audio io.ReadCloser) *CreateJobOptions {
	_options.Audio = audio
	return _options
}

// SetContentType : Allow user to set ContentType
func (_options *CreateJobOptions) SetContentType(contentType string) *CreateJobOptions {
	_options.ContentType = core.StringPtr(contentType)
	return _options
}

// SetModel : Allow user to set Model
func (_options *CreateJobOptions) SetModel(model string) *CreateJobOptions {
	_options.Model = core.StringPtr(model)
	return _options
}

// SetCallbackURL : Allow user to set CallbackURL
func (_options *CreateJobOptions) SetCallbackURL(callbackURL string) *CreateJobOptions {
	_options.CallbackURL = core.StringPtr(callbackURL)
	return _options
}

// SetEvents : Allow user to set Events
func (_options *CreateJobOptions) SetEvents(events string) *CreateJobOptions {
	_options.Events = core.StringPtr(events)
	return _options
}

// SetUserToken : Allow user to set UserToken
func (_options *CreateJobOptions) SetUserToken(userToken string) *CreateJobOptions {
	_options.UserToken = core.StringPtr(userToken)
	return _options
}

// SetResultsTTL : Allow user to set ResultsTTL
func (_options *CreateJobOptions) SetResultsTTL(resultsTTL int64) *CreateJobOptions {
	_options.ResultsTTL = core.Int64Ptr(resultsTTL)
	return _options
}

// SetLanguageCustomizationID : Allow user to set LanguageCustomizationID
func (_options *CreateJobOptions) SetLanguageCustomizationID(languageCustomizationID string) *CreateJobOptions {
	_options.LanguageCustomizationID = core.StringPtr(languageCustomizationID)
	return _options
}

// SetAcousticCustomizationID : Allow user to set AcousticCustomizationID
func (_options *CreateJobOptions) SetAcousticCustomizationID(acousticCustomizationID string) *CreateJobOptions {
	_options.AcousticCustomizationID = core.StringPtr(acousticCustomizationID)
	return _options
}

// SetBaseModelVersion : Allow user to set BaseModelVersion
func (_options *CreateJobOptions) SetBaseModelVersion(baseModelVersion string) *CreateJobOptions {
	_options.BaseModelVersion = core.StringPtr(baseModelVersion)
	return _options
}

// SetCustomizationWeight : Allow user to set CustomizationWeight
func (_options *CreateJobOptions) SetCustomizationWeight(customizationWeight float64) *CreateJobOptions {
	_options.CustomizationWeight = core.Float64Ptr(customizationWeight)
	return _options
}

// SetInactivityTimeout : Allow user to set InactivityTimeout
func (_options *CreateJobOptions) SetInactivityTimeout(inactivityTimeout int64) *CreateJobOptions {
	_options.InactivityTimeout = core.Int64Ptr(inactivityTimeout)
	return _options
}

// SetKeywords : Allow user to set Keywords
func (_options *CreateJobOptions) SetKeywords(keywords []string) *CreateJobOptions {
	_options.Keywords = keywords
	return _options
}

// SetKeywordsThreshold : Allow user to set KeywordsThreshold
func (_options *CreateJobOptions) SetKeywordsThreshold(keywordsThreshold float32) *CreateJobOptions {
	_options.KeywordsThreshold = core.Float32Ptr(keywordsThreshold)
	return _options
}

// SetMaxAlternatives : Allow user to set MaxAlternatives
func (_options *CreateJobOptions) SetMaxAlternatives(maxAlternatives int64) *CreateJobOptions {
	_options.MaxAlternatives = core.Int64Ptr(maxAlternatives)
	return _options
}

// SetWordAlternativesThreshold : Allow user to set WordAlternativesThreshold
func (_options *CreateJobOptions) SetWordAlternativesThreshold(wordAlternativesThreshold float32) *CreateJobOptions {
	_options.WordAlternativesThreshold = core.Float32Ptr(wordAlternativesThreshold)
	return _options
}

// SetWordConfidence : Allow user to set WordConfidence
func (_options *CreateJobOptions) SetWordConfidence(wordConfidence bool) *CreateJobOptions {
	_options.WordConfidence = core.BoolPtr(wordConfidence)
	return _options
}

// SetTimestamps : Allow user to set Timestamps
func (_options *CreateJobOptions) SetTimestamps(timestamps bool) *CreateJobOptions {
	_options.Timestamps = core.BoolPtr(timestamps)
	return _options
}

// SetProfanityFilter : Allow user to set ProfanityFilter
func (_options *CreateJobOptions) SetProfanityFilter(profanityFilter bool) *CreateJobOptions {
	_options.ProfanityFilter = core.BoolPtr(profanityFilter)
	return _options
}

// SetSmartFormatting : Allow user to set SmartFormatting
func (_options *CreateJobOptions) SetSmartFormatting(smartFormatting bool) *CreateJobOptions {
	_options.SmartFormatting = core.BoolPtr(smartFormatting)
	return _options
}

// SetSpeakerLabels : Allow user to set SpeakerLabels
func (_options *CreateJobOptions) SetSpeakerLabels(speakerLabels bool) *CreateJobOptions {
	_options.SpeakerLabels = core.BoolPtr(speakerLabels)
	return _options
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *CreateJobOptions) SetCustomizationID(customizationID string) *CreateJobOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetGrammarName : Allow user to set GrammarName
func (_options *CreateJobOptions) SetGrammarName(grammarName string) *CreateJobOptions {
	_options.GrammarName = core.StringPtr(grammarName)
	return _options
}

// SetRedaction : Allow user to set Redaction
func (_options *CreateJobOptions) SetRedaction(redaction bool) *CreateJobOptions {
	_options.Redaction = core.BoolPtr(redaction)
	return _options
}

// SetProcessingMetrics : Allow user to set ProcessingMetrics
func (_options *CreateJobOptions) SetProcessingMetrics(processingMetrics bool) *CreateJobOptions {
	_options.ProcessingMetrics = core.BoolPtr(processingMetrics)
	return _options
}

// SetProcessingMetricsInterval : Allow user to set ProcessingMetricsInterval
func (_options *CreateJobOptions) SetProcessingMetricsInterval(processingMetricsInterval float32) *CreateJobOptions {
	_options.ProcessingMetricsInterval = core.Float32Ptr(processingMetricsInterval)
	return _options
}

// SetAudioMetrics : Allow user to set AudioMetrics
func (_options *CreateJobOptions) SetAudioMetrics(audioMetrics bool) *CreateJobOptions {
	_options.AudioMetrics = core.BoolPtr(audioMetrics)
	return _options
}

// SetEndOfPhraseSilenceTime : Allow user to set EndOfPhraseSilenceTime
func (_options *CreateJobOptions) SetEndOfPhraseSilenceTime(endOfPhraseSilenceTime float64) *CreateJobOptions {
	_options.EndOfPhraseSilenceTime = core.Float64Ptr(endOfPhraseSilenceTime)
	return _options
}

// SetSplitTranscriptAtPhraseEnd : Allow user to set SplitTranscriptAtPhraseEnd
func (_options *CreateJobOptions) SetSplitTranscriptAtPhraseEnd(splitTranscriptAtPhraseEnd bool) *CreateJobOptions {
	_options.SplitTranscriptAtPhraseEnd = core.BoolPtr(splitTranscriptAtPhraseEnd)
	return _options
}

// SetSpeechDetectorSensitivity : Allow user to set SpeechDetectorSensitivity
func (_options *CreateJobOptions) SetSpeechDetectorSensitivity(speechDetectorSensitivity float32) *CreateJobOptions {
	_options.SpeechDetectorSensitivity = core.Float32Ptr(speechDetectorSensitivity)
	return _options
}

// SetBackgroundAudioSuppression : Allow user to set BackgroundAudioSuppression
func (_options *CreateJobOptions) SetBackgroundAudioSuppression(backgroundAudioSuppression float32) *CreateJobOptions {
	_options.BackgroundAudioSuppression = core.Float32Ptr(backgroundAudioSuppression)
	return _options
}

// SetLowLatency : Allow user to set LowLatency
func (_options *CreateJobOptions) SetLowLatency(lowLatency bool) *CreateJobOptions {
	_options.LowLatency = core.BoolPtr(lowLatency)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateJobOptions) SetHeaders(param map[string]string) *CreateJobOptions {
	options.Headers = param
	return options
}

// CreateLanguageModelOptions : The CreateLanguageModel options.
type CreateLanguageModelOptions struct {
	// A user-defined name for the new custom language model. Use a name that is unique among all custom language models
	// that you own. Use a localized name that matches the language of the custom model. Use a name that describes the
	// domain of the custom model, such as `Medical custom model` or `Legal custom model`.
	Name *string `json:"name" validate:"required"`

	// The name of the base language model that is to be customized by the new custom language model. The new custom model
	// can be used only with the base model that it customizes.
	//
	// To determine whether a base model supports language model customization, use the [Get a model](#getmodel) method and
	// check that the attribute `custom_language_model` is set to `true`. You can also refer to [Language support for
	// customization](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-custom-support).
	BaseModelName *string `json:"base_model_name" validate:"required"`

	// The dialect of the specified language that is to be used with the custom language model. _For all languages, it is
	// always safe to omit this field._ The service automatically uses the language identifier from the name of the base
	// model. For example, the service automatically uses `en-US` for all US English models.
	//
	// If you specify the `dialect` for a new custom model, follow these guidelines. _For non-Spanish previous-generation
	// models and for next-generation models,_ you must specify a value that matches the five-character language identifier
	// from the name of the base model. _For Spanish previous-generation models,_ you must specify one of the following
	// values:
	// * `es-ES` for Castilian Spanish (`es-ES` models)
	// * `es-LA` for Latin American Spanish (`es-AR`, `es-CL`, `es-CO`, and `es-PE` models)
	// * `es-US` for Mexican (North American) Spanish (`es-MX` models)
	//
	// All values that you pass for the `dialect` field are case-insensitive.
	Dialect *string `json:"dialect,omitempty"`

	// A description of the new custom language model. Use a localized description that matches the language of the custom
	// model.
	Description *string `json:"description,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateLanguageModelOptions.BaseModelName property.
// The name of the base language model that is to be customized by the new custom language model. The new custom model
// can be used only with the base model that it customizes.
//
// To determine whether a base model supports language model customization, use the [Get a model](#getmodel) method and
// check that the attribute `custom_language_model` is set to `true`. You can also refer to [Language support for
// customization](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-custom-support).
const (
	CreateLanguageModelOptionsBaseModelNameArMsTelephonyConst                = "ar-MS_Telephony"
	CreateLanguageModelOptionsBaseModelNameCsCzTelephonyConst                = "cs-CZ_Telephony"
	CreateLanguageModelOptionsBaseModelNameDeDeBroadbandmodelConst           = "de-DE_BroadbandModel"
	CreateLanguageModelOptionsBaseModelNameDeDeMultimediaConst               = "de-DE_Multimedia"
	CreateLanguageModelOptionsBaseModelNameDeDeNarrowbandmodelConst          = "de-DE_NarrowbandModel"
	CreateLanguageModelOptionsBaseModelNameDeDeTelephonyConst                = "de-DE_Telephony"
	CreateLanguageModelOptionsBaseModelNameEnAuBroadbandmodelConst           = "en-AU_BroadbandModel"
	CreateLanguageModelOptionsBaseModelNameEnAuMultimediaConst               = "en-AU_Multimedia"
	CreateLanguageModelOptionsBaseModelNameEnAuNarrowbandmodelConst          = "en-AU_NarrowbandModel"
	CreateLanguageModelOptionsBaseModelNameEnAuTelephonyConst                = "en-AU_Telephony"
	CreateLanguageModelOptionsBaseModelNameEnGbBroadbandmodelConst           = "en-GB_BroadbandModel"
	CreateLanguageModelOptionsBaseModelNameEnGbMultimediaConst               = "en-GB_Multimedia"
	CreateLanguageModelOptionsBaseModelNameEnGbNarrowbandmodelConst          = "en-GB_NarrowbandModel"
	CreateLanguageModelOptionsBaseModelNameEnGbTelephonyConst                = "en-GB_Telephony"
	CreateLanguageModelOptionsBaseModelNameEnInTelephonyConst                = "en-IN_Telephony"
	CreateLanguageModelOptionsBaseModelNameEnUsBroadbandmodelConst           = "en-US_BroadbandModel"
	CreateLanguageModelOptionsBaseModelNameEnUsMultimediaConst               = "en-US_Multimedia"
	CreateLanguageModelOptionsBaseModelNameEnUsNarrowbandmodelConst          = "en-US_NarrowbandModel"
	CreateLanguageModelOptionsBaseModelNameEnUsShortformNarrowbandmodelConst = "en-US_ShortForm_NarrowbandModel"
	CreateLanguageModelOptionsBaseModelNameEnUsTelephonyConst                = "en-US_Telephony"
	CreateLanguageModelOptionsBaseModelNameEnWwMedicalTelephonyConst         = "en-WW_Medical_Telephony"
	CreateLanguageModelOptionsBaseModelNameEsArBroadbandmodelConst           = "es-AR_BroadbandModel"
	CreateLanguageModelOptionsBaseModelNameEsArNarrowbandmodelConst          = "es-AR_NarrowbandModel"
	CreateLanguageModelOptionsBaseModelNameEsClBroadbandmodelConst           = "es-CL_BroadbandModel"
	CreateLanguageModelOptionsBaseModelNameEsClNarrowbandmodelConst          = "es-CL_NarrowbandModel"
	CreateLanguageModelOptionsBaseModelNameEsCoBroadbandmodelConst           = "es-CO_BroadbandModel"
	CreateLanguageModelOptionsBaseModelNameEsCoNarrowbandmodelConst          = "es-CO_NarrowbandModel"
	CreateLanguageModelOptionsBaseModelNameEsEsBroadbandmodelConst           = "es-ES_BroadbandModel"
	CreateLanguageModelOptionsBaseModelNameEsEsMultimediaConst               = "es-ES_Multimedia"
	CreateLanguageModelOptionsBaseModelNameEsEsNarrowbandmodelConst          = "es-ES_NarrowbandModel"
	CreateLanguageModelOptionsBaseModelNameEsEsTelephonyConst                = "es-ES_Telephony"
	CreateLanguageModelOptionsBaseModelNameEsLaTelephonyConst                = "es-LA_Telephony"
	CreateLanguageModelOptionsBaseModelNameEsMxBroadbandmodelConst           = "es-MX_BroadbandModel"
	CreateLanguageModelOptionsBaseModelNameEsMxNarrowbandmodelConst          = "es-MX_NarrowbandModel"
	CreateLanguageModelOptionsBaseModelNameEsPeBroadbandmodelConst           = "es-PE_BroadbandModel"
	CreateLanguageModelOptionsBaseModelNameEsPeNarrowbandmodelConst          = "es-PE_NarrowbandModel"
	CreateLanguageModelOptionsBaseModelNameFrCaBroadbandmodelConst           = "fr-CA_BroadbandModel"
	CreateLanguageModelOptionsBaseModelNameFrCaNarrowbandmodelConst          = "fr-CA_NarrowbandModel"
	CreateLanguageModelOptionsBaseModelNameFrCaTelephonyConst                = "fr-CA_Telephony"
	CreateLanguageModelOptionsBaseModelNameFrFrBroadbandmodelConst           = "fr-FR_BroadbandModel"
	CreateLanguageModelOptionsBaseModelNameFrFrMultimediaConst               = "fr-FR_Multimedia"
	CreateLanguageModelOptionsBaseModelNameFrFrNarrowbandmodelConst          = "fr-FR_NarrowbandModel"
	CreateLanguageModelOptionsBaseModelNameFrFrTelephonyConst                = "fr-FR_Telephony"
	CreateLanguageModelOptionsBaseModelNameHiInTelephonyConst                = "hi-IN_Telephony"
	CreateLanguageModelOptionsBaseModelNameItItBroadbandmodelConst           = "it-IT_BroadbandModel"
	CreateLanguageModelOptionsBaseModelNameItItNarrowbandmodelConst          = "it-IT_NarrowbandModel"
	CreateLanguageModelOptionsBaseModelNameItItTelephonyConst                = "it-IT_Telephony"
	CreateLanguageModelOptionsBaseModelNameJaJpBroadbandmodelConst           = "ja-JP_BroadbandModel"
	CreateLanguageModelOptionsBaseModelNameJaJpMultimediaConst               = "ja-JP_Multimedia"
	CreateLanguageModelOptionsBaseModelNameJaJpNarrowbandmodelConst          = "ja-JP_NarrowbandModel"
	CreateLanguageModelOptionsBaseModelNameKoKrBroadbandmodelConst           = "ko-KR_BroadbandModel"
	CreateLanguageModelOptionsBaseModelNameKoKrMultimediaConst               = "ko-KR_Multimedia"
	CreateLanguageModelOptionsBaseModelNameKoKrNarrowbandmodelConst          = "ko-KR_NarrowbandModel"
	CreateLanguageModelOptionsBaseModelNameKoKrTelephonyConst                = "ko-KR_Telephony"
	CreateLanguageModelOptionsBaseModelNameNlBeTelephonyConst                = "nl-BE_Telephony"
	CreateLanguageModelOptionsBaseModelNameNlNlBroadbandmodelConst           = "nl-NL_BroadbandModel"
	CreateLanguageModelOptionsBaseModelNameNlNlNarrowbandmodelConst          = "nl-NL_NarrowbandModel"
	CreateLanguageModelOptionsBaseModelNameNlNlTelephonyConst                = "nl-NL_Telephony"
	CreateLanguageModelOptionsBaseModelNamePtBrBroadbandmodelConst           = "pt-BR_BroadbandModel"
	CreateLanguageModelOptionsBaseModelNamePtBrNarrowbandmodelConst          = "pt-BR_NarrowbandModel"
	CreateLanguageModelOptionsBaseModelNamePtBrTelephonyConst                = "pt-BR_Telephony"
	CreateLanguageModelOptionsBaseModelNameZhCnTelephonyConst                = "zh-CN_Telephony"
)

// NewCreateLanguageModelOptions : Instantiate CreateLanguageModelOptions
func (*SpeechToTextV1) NewCreateLanguageModelOptions(name string, baseModelName string) *CreateLanguageModelOptions {
	return &CreateLanguageModelOptions{
		Name:          core.StringPtr(name),
		BaseModelName: core.StringPtr(baseModelName),
	}
}

// SetName : Allow user to set Name
func (_options *CreateLanguageModelOptions) SetName(name string) *CreateLanguageModelOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetBaseModelName : Allow user to set BaseModelName
func (_options *CreateLanguageModelOptions) SetBaseModelName(baseModelName string) *CreateLanguageModelOptions {
	_options.BaseModelName = core.StringPtr(baseModelName)
	return _options
}

// SetDialect : Allow user to set Dialect
func (_options *CreateLanguageModelOptions) SetDialect(dialect string) *CreateLanguageModelOptions {
	_options.Dialect = core.StringPtr(dialect)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateLanguageModelOptions) SetDescription(description string) *CreateLanguageModelOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateLanguageModelOptions) SetHeaders(param map[string]string) *CreateLanguageModelOptions {
	options.Headers = param
	return options
}

// CustomWord : Information about a word that is to be added to a custom language model.
type CustomWord struct {
	// For the [Add custom words](#addwords) method, you must specify the custom word that is to be added to or updated in
	// the custom model. Do not include spaces in the word. Use a `-` (dash) or `_` (underscore) to connect the tokens of
	// compound words.
	//
	// Omit this parameter for the [Add a custom word](#addword) method.
	Word *string `json:"word,omitempty"`

	// _For a custom model that is based on a previous-generation model_, an array of sounds-like pronunciations for the
	// custom word. Specify how words that are difficult to pronounce, foreign words, acronyms, and so on can be pronounced
	// by users.
	// * For a word that is not in the service's base vocabulary, omit the parameter to have the service automatically
	// generate a sounds-like pronunciation for the word.
	// * For a word that is in the service's base vocabulary, use the parameter to specify additional pronunciations for
	// the word. You cannot override the default pronunciation of a word; pronunciations you add augment the pronunciation
	// from the base vocabulary.
	//
	// A word can have at most five sounds-like pronunciations. A pronunciation can include at most 40 characters not
	// including spaces.
	//
	// _For a custom model that is based on a next-generation model_, omit this field. Custom models based on
	// next-generation models do not support the `sounds_like` field. The service ignores the field.
	SoundsLike []string `json:"sounds_like,omitempty"`

	// An alternative spelling for the custom word when it appears in a transcript. Use the parameter when you want the
	// word to have a spelling that is different from its usual representation or from its spelling in corpora training
	// data.
	DisplayAs *string `json:"display_as,omitempty"`
}

// UnmarshalCustomWord unmarshals an instance of CustomWord from the specified map of raw messages.
func UnmarshalCustomWord(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CustomWord)
	err = core.UnmarshalPrimitive(m, "word", &obj.Word)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sounds_like", &obj.SoundsLike)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "display_as", &obj.DisplayAs)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteAcousticModelOptions : The DeleteAcousticModel options.
type DeleteAcousticModelOptions struct {
	// The customization ID (GUID) of the custom acoustic model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteAcousticModelOptions : Instantiate DeleteAcousticModelOptions
func (*SpeechToTextV1) NewDeleteAcousticModelOptions(customizationID string) *DeleteAcousticModelOptions {
	return &DeleteAcousticModelOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *DeleteAcousticModelOptions) SetCustomizationID(customizationID string) *DeleteAcousticModelOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteAcousticModelOptions) SetHeaders(param map[string]string) *DeleteAcousticModelOptions {
	options.Headers = param
	return options
}

// DeleteAudioOptions : The DeleteAudio options.
type DeleteAudioOptions struct {
	// The customization ID (GUID) of the custom acoustic model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// The name of the audio resource for the custom acoustic model.
	AudioName *string `json:"audio_name" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteAudioOptions : Instantiate DeleteAudioOptions
func (*SpeechToTextV1) NewDeleteAudioOptions(customizationID string, audioName string) *DeleteAudioOptions {
	return &DeleteAudioOptions{
		CustomizationID: core.StringPtr(customizationID),
		AudioName:       core.StringPtr(audioName),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *DeleteAudioOptions) SetCustomizationID(customizationID string) *DeleteAudioOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetAudioName : Allow user to set AudioName
func (_options *DeleteAudioOptions) SetAudioName(audioName string) *DeleteAudioOptions {
	_options.AudioName = core.StringPtr(audioName)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteAudioOptions) SetHeaders(param map[string]string) *DeleteAudioOptions {
	options.Headers = param
	return options
}

// DeleteCorpusOptions : The DeleteCorpus options.
type DeleteCorpusOptions struct {
	// The customization ID (GUID) of the custom language model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// The name of the corpus for the custom language model.
	CorpusName *string `json:"corpus_name" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteCorpusOptions : Instantiate DeleteCorpusOptions
func (*SpeechToTextV1) NewDeleteCorpusOptions(customizationID string, corpusName string) *DeleteCorpusOptions {
	return &DeleteCorpusOptions{
		CustomizationID: core.StringPtr(customizationID),
		CorpusName:      core.StringPtr(corpusName),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *DeleteCorpusOptions) SetCustomizationID(customizationID string) *DeleteCorpusOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetCorpusName : Allow user to set CorpusName
func (_options *DeleteCorpusOptions) SetCorpusName(corpusName string) *DeleteCorpusOptions {
	_options.CorpusName = core.StringPtr(corpusName)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteCorpusOptions) SetHeaders(param map[string]string) *DeleteCorpusOptions {
	options.Headers = param
	return options
}

// DeleteGrammarOptions : The DeleteGrammar options.
type DeleteGrammarOptions struct {
	// The customization ID (GUID) of the custom language model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// The name of the grammar for the custom language model.
	GrammarName *string `json:"grammar_name" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteGrammarOptions : Instantiate DeleteGrammarOptions
func (*SpeechToTextV1) NewDeleteGrammarOptions(customizationID string, grammarName string) *DeleteGrammarOptions {
	return &DeleteGrammarOptions{
		CustomizationID: core.StringPtr(customizationID),
		GrammarName:     core.StringPtr(grammarName),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *DeleteGrammarOptions) SetCustomizationID(customizationID string) *DeleteGrammarOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetGrammarName : Allow user to set GrammarName
func (_options *DeleteGrammarOptions) SetGrammarName(grammarName string) *DeleteGrammarOptions {
	_options.GrammarName = core.StringPtr(grammarName)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteGrammarOptions) SetHeaders(param map[string]string) *DeleteGrammarOptions {
	options.Headers = param
	return options
}

// DeleteJobOptions : The DeleteJob options.
type DeleteJobOptions struct {
	// The identifier of the asynchronous job that is to be used for the request. You must make the request with
	// credentials for the instance of the service that owns the job.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteJobOptions : Instantiate DeleteJobOptions
func (*SpeechToTextV1) NewDeleteJobOptions(id string) *DeleteJobOptions {
	return &DeleteJobOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *DeleteJobOptions) SetID(id string) *DeleteJobOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteJobOptions) SetHeaders(param map[string]string) *DeleteJobOptions {
	options.Headers = param
	return options
}

// DeleteLanguageModelOptions : The DeleteLanguageModel options.
type DeleteLanguageModelOptions struct {
	// The customization ID (GUID) of the custom language model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteLanguageModelOptions : Instantiate DeleteLanguageModelOptions
func (*SpeechToTextV1) NewDeleteLanguageModelOptions(customizationID string) *DeleteLanguageModelOptions {
	return &DeleteLanguageModelOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *DeleteLanguageModelOptions) SetCustomizationID(customizationID string) *DeleteLanguageModelOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteLanguageModelOptions) SetHeaders(param map[string]string) *DeleteLanguageModelOptions {
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
func (*SpeechToTextV1) NewDeleteUserDataOptions(customerID string) *DeleteUserDataOptions {
	return &DeleteUserDataOptions{
		CustomerID: core.StringPtr(customerID),
	}
}

// SetCustomerID : Allow user to set CustomerID
func (_options *DeleteUserDataOptions) SetCustomerID(customerID string) *DeleteUserDataOptions {
	_options.CustomerID = core.StringPtr(customerID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteUserDataOptions) SetHeaders(param map[string]string) *DeleteUserDataOptions {
	options.Headers = param
	return options
}

// DeleteWordOptions : The DeleteWord options.
type DeleteWordOptions struct {
	// The customization ID (GUID) of the custom language model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// The custom word that is to be deleted from the custom language model. URL-encode the word if it includes non-ASCII
	// characters. For more information, see [Character
	// encoding](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-corporaWords#charEncoding).
	WordName *string `json:"word_name" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteWordOptions : Instantiate DeleteWordOptions
func (*SpeechToTextV1) NewDeleteWordOptions(customizationID string, wordName string) *DeleteWordOptions {
	return &DeleteWordOptions{
		CustomizationID: core.StringPtr(customizationID),
		WordName:        core.StringPtr(wordName),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *DeleteWordOptions) SetCustomizationID(customizationID string) *DeleteWordOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetWordName : Allow user to set WordName
func (_options *DeleteWordOptions) SetWordName(wordName string) *DeleteWordOptions {
	_options.WordName = core.StringPtr(wordName)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteWordOptions) SetHeaders(param map[string]string) *DeleteWordOptions {
	options.Headers = param
	return options
}

// GetAcousticModelOptions : The GetAcousticModel options.
type GetAcousticModelOptions struct {
	// The customization ID (GUID) of the custom acoustic model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAcousticModelOptions : Instantiate GetAcousticModelOptions
func (*SpeechToTextV1) NewGetAcousticModelOptions(customizationID string) *GetAcousticModelOptions {
	return &GetAcousticModelOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *GetAcousticModelOptions) SetCustomizationID(customizationID string) *GetAcousticModelOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetAcousticModelOptions) SetHeaders(param map[string]string) *GetAcousticModelOptions {
	options.Headers = param
	return options
}

// GetAudioOptions : The GetAudio options.
type GetAudioOptions struct {
	// The customization ID (GUID) of the custom acoustic model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// The name of the audio resource for the custom acoustic model.
	AudioName *string `json:"audio_name" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAudioOptions : Instantiate GetAudioOptions
func (*SpeechToTextV1) NewGetAudioOptions(customizationID string, audioName string) *GetAudioOptions {
	return &GetAudioOptions{
		CustomizationID: core.StringPtr(customizationID),
		AudioName:       core.StringPtr(audioName),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *GetAudioOptions) SetCustomizationID(customizationID string) *GetAudioOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetAudioName : Allow user to set AudioName
func (_options *GetAudioOptions) SetAudioName(audioName string) *GetAudioOptions {
	_options.AudioName = core.StringPtr(audioName)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetAudioOptions) SetHeaders(param map[string]string) *GetAudioOptions {
	options.Headers = param
	return options
}

// GetCorpusOptions : The GetCorpus options.
type GetCorpusOptions struct {
	// The customization ID (GUID) of the custom language model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// The name of the corpus for the custom language model.
	CorpusName *string `json:"corpus_name" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCorpusOptions : Instantiate GetCorpusOptions
func (*SpeechToTextV1) NewGetCorpusOptions(customizationID string, corpusName string) *GetCorpusOptions {
	return &GetCorpusOptions{
		CustomizationID: core.StringPtr(customizationID),
		CorpusName:      core.StringPtr(corpusName),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *GetCorpusOptions) SetCustomizationID(customizationID string) *GetCorpusOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetCorpusName : Allow user to set CorpusName
func (_options *GetCorpusOptions) SetCorpusName(corpusName string) *GetCorpusOptions {
	_options.CorpusName = core.StringPtr(corpusName)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetCorpusOptions) SetHeaders(param map[string]string) *GetCorpusOptions {
	options.Headers = param
	return options
}

// GetGrammarOptions : The GetGrammar options.
type GetGrammarOptions struct {
	// The customization ID (GUID) of the custom language model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// The name of the grammar for the custom language model.
	GrammarName *string `json:"grammar_name" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetGrammarOptions : Instantiate GetGrammarOptions
func (*SpeechToTextV1) NewGetGrammarOptions(customizationID string, grammarName string) *GetGrammarOptions {
	return &GetGrammarOptions{
		CustomizationID: core.StringPtr(customizationID),
		GrammarName:     core.StringPtr(grammarName),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *GetGrammarOptions) SetCustomizationID(customizationID string) *GetGrammarOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetGrammarName : Allow user to set GrammarName
func (_options *GetGrammarOptions) SetGrammarName(grammarName string) *GetGrammarOptions {
	_options.GrammarName = core.StringPtr(grammarName)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetGrammarOptions) SetHeaders(param map[string]string) *GetGrammarOptions {
	options.Headers = param
	return options
}

// GetLanguageModelOptions : The GetLanguageModel options.
type GetLanguageModelOptions struct {
	// The customization ID (GUID) of the custom language model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetLanguageModelOptions : Instantiate GetLanguageModelOptions
func (*SpeechToTextV1) NewGetLanguageModelOptions(customizationID string) *GetLanguageModelOptions {
	return &GetLanguageModelOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *GetLanguageModelOptions) SetCustomizationID(customizationID string) *GetLanguageModelOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetLanguageModelOptions) SetHeaders(param map[string]string) *GetLanguageModelOptions {
	options.Headers = param
	return options
}

// GetModelOptions : The GetModel options.
type GetModelOptions struct {
	// The identifier of the model in the form of its name from the output of the [List models](#listmodels) method.
	// (**Note:** The model `ar-AR_BroadbandModel` is deprecated; use `ar-MS_BroadbandModel` instead.).
	ModelID *string `json:"model_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetModelOptions.ModelID property.
// The identifier of the model in the form of its name from the output of the [List models](#listmodels) method.
// (**Note:** The model `ar-AR_BroadbandModel` is deprecated; use `ar-MS_BroadbandModel` instead.).
const (
	GetModelOptionsModelIDArArBroadbandmodelConst           = "ar-AR_BroadbandModel"
	GetModelOptionsModelIDArMsBroadbandmodelConst           = "ar-MS_BroadbandModel"
	GetModelOptionsModelIDArMsTelephonyConst                = "ar-MS_Telephony"
	GetModelOptionsModelIDCsCzTelephonyConst                = "cs-CZ_Telephony"
	GetModelOptionsModelIDDeDeBroadbandmodelConst           = "de-DE_BroadbandModel"
	GetModelOptionsModelIDDeDeMultimediaConst               = "de-DE_Multimedia"
	GetModelOptionsModelIDDeDeNarrowbandmodelConst          = "de-DE_NarrowbandModel"
	GetModelOptionsModelIDDeDeTelephonyConst                = "de-DE_Telephony"
	GetModelOptionsModelIDEnAuBroadbandmodelConst           = "en-AU_BroadbandModel"
	GetModelOptionsModelIDEnAuMultimediaConst               = "en-AU_Multimedia"
	GetModelOptionsModelIDEnAuNarrowbandmodelConst          = "en-AU_NarrowbandModel"
	GetModelOptionsModelIDEnAuTelephonyConst                = "en-AU_Telephony"
	GetModelOptionsModelIDEnGbBroadbandmodelConst           = "en-GB_BroadbandModel"
	GetModelOptionsModelIDEnGbMultimediaConst               = "en-GB_Multimedia"
	GetModelOptionsModelIDEnGbNarrowbandmodelConst          = "en-GB_NarrowbandModel"
	GetModelOptionsModelIDEnGbTelephonyConst                = "en-GB_Telephony"
	GetModelOptionsModelIDEnInTelephonyConst                = "en-IN_Telephony"
	GetModelOptionsModelIDEnUsBroadbandmodelConst           = "en-US_BroadbandModel"
	GetModelOptionsModelIDEnUsMultimediaConst               = "en-US_Multimedia"
	GetModelOptionsModelIDEnUsNarrowbandmodelConst          = "en-US_NarrowbandModel"
	GetModelOptionsModelIDEnUsShortformNarrowbandmodelConst = "en-US_ShortForm_NarrowbandModel"
	GetModelOptionsModelIDEnUsTelephonyConst                = "en-US_Telephony"
	GetModelOptionsModelIDEnWwMedicalTelephonyConst         = "en-WW_Medical_Telephony"
	GetModelOptionsModelIDEsArBroadbandmodelConst           = "es-AR_BroadbandModel"
	GetModelOptionsModelIDEsArNarrowbandmodelConst          = "es-AR_NarrowbandModel"
	GetModelOptionsModelIDEsClBroadbandmodelConst           = "es-CL_BroadbandModel"
	GetModelOptionsModelIDEsClNarrowbandmodelConst          = "es-CL_NarrowbandModel"
	GetModelOptionsModelIDEsCoBroadbandmodelConst           = "es-CO_BroadbandModel"
	GetModelOptionsModelIDEsCoNarrowbandmodelConst          = "es-CO_NarrowbandModel"
	GetModelOptionsModelIDEsEsBroadbandmodelConst           = "es-ES_BroadbandModel"
	GetModelOptionsModelIDEsEsMultimediaConst               = "es-ES_Multimedia"
	GetModelOptionsModelIDEsEsNarrowbandmodelConst          = "es-ES_NarrowbandModel"
	GetModelOptionsModelIDEsEsTelephonyConst                = "es-ES_Telephony"
	GetModelOptionsModelIDEsLaTelephonyConst                = "es-LA_Telephony"
	GetModelOptionsModelIDEsMxBroadbandmodelConst           = "es-MX_BroadbandModel"
	GetModelOptionsModelIDEsMxNarrowbandmodelConst          = "es-MX_NarrowbandModel"
	GetModelOptionsModelIDEsPeBroadbandmodelConst           = "es-PE_BroadbandModel"
	GetModelOptionsModelIDEsPeNarrowbandmodelConst          = "es-PE_NarrowbandModel"
	GetModelOptionsModelIDFrCaBroadbandmodelConst           = "fr-CA_BroadbandModel"
	GetModelOptionsModelIDFrCaNarrowbandmodelConst          = "fr-CA_NarrowbandModel"
	GetModelOptionsModelIDFrCaTelephonyConst                = "fr-CA_Telephony"
	GetModelOptionsModelIDFrFrBroadbandmodelConst           = "fr-FR_BroadbandModel"
	GetModelOptionsModelIDFrFrMultimediaConst               = "fr-FR_Multimedia"
	GetModelOptionsModelIDFrFrNarrowbandmodelConst          = "fr-FR_NarrowbandModel"
	GetModelOptionsModelIDFrFrTelephonyConst                = "fr-FR_Telephony"
	GetModelOptionsModelIDHiInTelephonyConst                = "hi-IN_Telephony"
	GetModelOptionsModelIDItItBroadbandmodelConst           = "it-IT_BroadbandModel"
	GetModelOptionsModelIDItItNarrowbandmodelConst          = "it-IT_NarrowbandModel"
	GetModelOptionsModelIDItItTelephonyConst                = "it-IT_Telephony"
	GetModelOptionsModelIDJaJpBroadbandmodelConst           = "ja-JP_BroadbandModel"
	GetModelOptionsModelIDJaJpMultimediaConst               = "ja-JP_Multimedia"
	GetModelOptionsModelIDJaJpNarrowbandmodelConst          = "ja-JP_NarrowbandModel"
	GetModelOptionsModelIDKoKrBroadbandmodelConst           = "ko-KR_BroadbandModel"
	GetModelOptionsModelIDKoKrMultimediaConst               = "ko-KR_Multimedia"
	GetModelOptionsModelIDKoKrNarrowbandmodelConst          = "ko-KR_NarrowbandModel"
	GetModelOptionsModelIDKoKrTelephonyConst                = "ko-KR_Telephony"
	GetModelOptionsModelIDNlBeTelephonyConst                = "nl-BE_Telephony"
	GetModelOptionsModelIDNlNlBroadbandmodelConst           = "nl-NL_BroadbandModel"
	GetModelOptionsModelIDNlNlNarrowbandmodelConst          = "nl-NL_NarrowbandModel"
	GetModelOptionsModelIDNlNlTelephonyConst                = "nl-NL_Telephony"
	GetModelOptionsModelIDPtBrBroadbandmodelConst           = "pt-BR_BroadbandModel"
	GetModelOptionsModelIDPtBrNarrowbandmodelConst          = "pt-BR_NarrowbandModel"
	GetModelOptionsModelIDPtBrTelephonyConst                = "pt-BR_Telephony"
	GetModelOptionsModelIDZhCnBroadbandmodelConst           = "zh-CN_BroadbandModel"
	GetModelOptionsModelIDZhCnNarrowbandmodelConst          = "zh-CN_NarrowbandModel"
	GetModelOptionsModelIDZhCnTelephonyConst                = "zh-CN_Telephony"
)

// NewGetModelOptions : Instantiate GetModelOptions
func (*SpeechToTextV1) NewGetModelOptions(modelID string) *GetModelOptions {
	return &GetModelOptions{
		ModelID: core.StringPtr(modelID),
	}
}

// SetModelID : Allow user to set ModelID
func (_options *GetModelOptions) SetModelID(modelID string) *GetModelOptions {
	_options.ModelID = core.StringPtr(modelID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetModelOptions) SetHeaders(param map[string]string) *GetModelOptions {
	options.Headers = param
	return options
}

// GetWordOptions : The GetWord options.
type GetWordOptions struct {
	// The customization ID (GUID) of the custom language model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// The custom word that is to be read from the custom language model. URL-encode the word if it includes non-ASCII
	// characters. For more information, see [Character
	// encoding](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-corporaWords#charEncoding).
	WordName *string `json:"word_name" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetWordOptions : Instantiate GetWordOptions
func (*SpeechToTextV1) NewGetWordOptions(customizationID string, wordName string) *GetWordOptions {
	return &GetWordOptions{
		CustomizationID: core.StringPtr(customizationID),
		WordName:        core.StringPtr(wordName),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *GetWordOptions) SetCustomizationID(customizationID string) *GetWordOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetWordName : Allow user to set WordName
func (_options *GetWordOptions) SetWordName(wordName string) *GetWordOptions {
	_options.WordName = core.StringPtr(wordName)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetWordOptions) SetHeaders(param map[string]string) *GetWordOptions {
	options.Headers = param
	return options
}

// Grammar : Information about a grammar from a custom language model.
type Grammar struct {
	// The name of the grammar.
	Name *string `json:"name" validate:"required"`

	// _For custom models that are based on previous-generation models_, the number of OOV words extracted from the
	// grammar. The value is `0` while the grammar is being processed.
	//
	// _For custom models that are based on next-generation models_, no OOV words are extracted from grammars, so the value
	// is always `0`.
	OutOfVocabularyWords *int64 `json:"out_of_vocabulary_words" validate:"required"`

	// The status of the grammar:
	// * `analyzed`: The service successfully analyzed the grammar. The custom model can be trained with data from the
	// grammar.
	// * `being_processed`: The service is still analyzing the grammar. The service cannot accept requests to add new
	// resources or to train the custom model.
	// * `undetermined`: The service encountered an error while processing the grammar. The `error` field describes the
	// failure.
	Status *string `json:"status" validate:"required"`

	// If the status of the grammar is `undetermined`, the following message: `Analysis of grammar '{grammar_name}' failed.
	// Please try fixing the error or adding the grammar again by setting the 'allow_overwrite' flag to 'true'.`.
	Error *string `json:"error,omitempty"`
}

// Constants associated with the Grammar.Status property.
// The status of the grammar:
// * `analyzed`: The service successfully analyzed the grammar. The custom model can be trained with data from the
// grammar.
// * `being_processed`: The service is still analyzing the grammar. The service cannot accept requests to add new
// resources or to train the custom model.
// * `undetermined`: The service encountered an error while processing the grammar. The `error` field describes the
// failure.
const (
	GrammarStatusAnalyzedConst       = "analyzed"
	GrammarStatusBeingProcessedConst = "being_processed"
	GrammarStatusUndeterminedConst   = "undetermined"
)

// UnmarshalGrammar unmarshals an instance of Grammar from the specified map of raw messages.
func UnmarshalGrammar(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Grammar)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "out_of_vocabulary_words", &obj.OutOfVocabularyWords)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
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

// Grammars : Information about the grammars from a custom language model.
type Grammars struct {
	// An array of `Grammar` objects that provides information about the grammars for the custom model. The array is empty
	// if the custom model has no grammars.
	Grammars []Grammar `json:"grammars" validate:"required"`
}

// UnmarshalGrammars unmarshals an instance of Grammars from the specified map of raw messages.
func UnmarshalGrammars(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Grammars)
	err = core.UnmarshalModel(m, "grammars", &obj.Grammars, UnmarshalGrammar)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeywordResult : Information about a match for a keyword from speech recognition results.
type KeywordResult struct {
	// A specified keyword normalized to the spoken phrase that matched in the audio input.
	NormalizedText *string `json:"normalized_text" validate:"required"`

	// The start time in seconds of the keyword match.
	StartTime *float64 `json:"start_time" validate:"required"`

	// The end time in seconds of the keyword match.
	EndTime *float64 `json:"end_time" validate:"required"`

	// A confidence score for the keyword match in the range of 0.0 to 1.0.
	Confidence *float64 `json:"confidence" validate:"required"`
}

// UnmarshalKeywordResult unmarshals an instance of KeywordResult from the specified map of raw messages.
func UnmarshalKeywordResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeywordResult)
	err = core.UnmarshalPrimitive(m, "normalized_text", &obj.NormalizedText)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "start_time", &obj.StartTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "end_time", &obj.EndTime)
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

// LanguageModel : Information about an existing custom language model.
type LanguageModel struct {
	// The customization ID (GUID) of the custom language model. The [Create a custom language model](#createlanguagemodel)
	// method returns only this field of the object; it does not return the other fields.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// The date and time in Coordinated Universal Time (UTC) at which the custom language model was created. The value is
	// provided in full ISO 8601 format (`YYYY-MM-DDThh:mm:ss.sTZD`).
	Created *string `json:"created,omitempty"`

	// The date and time in Coordinated Universal Time (UTC) at which the custom language model was last modified. The
	// `created` and `updated` fields are equal when a language model is first added but has yet to be updated. The value
	// is provided in full ISO 8601 format (YYYY-MM-DDThh:mm:ss.sTZD).
	Updated *string `json:"updated,omitempty"`

	// The language identifier of the custom language model (for example, `en-US`). The value matches the five-character
	// language identifier from the name of the base model for the custom model. This value might be different from the
	// value of the `dialect` field.
	Language *string `json:"language,omitempty"`

	// The dialect of the language for the custom language model. _For custom models that are based on non-Spanish
	// previous-generation models and on next-generation models,_ the field matches the language of the base model; for
	// example, `en-US` for one of the US English models. _For custom models that are based on Spanish previous-generation
	// models,_ the field indicates the dialect with which the model was created. The value can match the name of the base
	// model or, if it was specified by the user, can be one of the following:
	// * `es-ES` for Castilian Spanish (`es-ES` models)
	// * `es-LA` for Latin American Spanish (`es-AR`, `es-CL`, `es-CO`, and `es-PE` models)
	// * `es-US` for Mexican (North American) Spanish (`es-MX` models)
	//
	// Dialect values are case-insensitive.
	Dialect *string `json:"dialect,omitempty"`

	// A list of the available versions of the custom language model. Each element of the array indicates a version of the
	// base model with which the custom model can be used. Multiple versions exist only if the custom model has been
	// upgraded to a new version of its base model. Otherwise, only a single version is shown.
	Versions []string `json:"versions,omitempty"`

	// The GUID of the credentials for the instance of the service that owns the custom language model.
	Owner *string `json:"owner,omitempty"`

	// The name of the custom language model.
	Name *string `json:"name,omitempty"`

	// The description of the custom language model.
	Description *string `json:"description,omitempty"`

	// The name of the language model for which the custom language model was created.
	BaseModelName *string `json:"base_model_name,omitempty"`

	// The current status of the custom language model:
	// * `pending`: The model was created but is waiting either for valid training data to be added or for the service to
	// finish analyzing added data.
	// * `ready`: The model contains valid data and is ready to be trained. If the model contains a mix of valid and
	// invalid resources, you need to set the `strict` parameter to `false` for the training to proceed.
	// * `training`: The model is currently being trained.
	// * `available`: The model is trained and ready to use.
	// * `upgrading`: The model is currently being upgraded.
	// * `failed`: Training of the model failed.
	Status *string `json:"status,omitempty"`

	// A percentage that indicates the progress of the custom language model's current training. A value of `100` means
	// that the model is fully trained. **Note:** The `progress` field does not currently reflect the progress of the
	// training. The field changes from `0` to `100` when training is complete.
	Progress *int64 `json:"progress,omitempty"`

	// If an error occurred while adding a grammar file to the custom language model, a message that describes an `Internal
	// Server Error` and includes the string `Cannot compile grammar`. The status of the custom model is not affected by
	// the error, but the grammar cannot be used with the model.
	Error *string `json:"error,omitempty"`

	// If the request included unknown parameters, the following message: `Unexpected query parameter(s) ['parameters']
	// detected`, where `parameters` is a list that includes a quoted string for each unknown parameter.
	Warnings *string `json:"warnings,omitempty"`
}

// Constants associated with the LanguageModel.Status property.
// The current status of the custom language model:
// * `pending`: The model was created but is waiting either for valid training data to be added or for the service to
// finish analyzing added data.
// * `ready`: The model contains valid data and is ready to be trained. If the model contains a mix of valid and invalid
// resources, you need to set the `strict` parameter to `false` for the training to proceed.
// * `training`: The model is currently being trained.
// * `available`: The model is trained and ready to use.
// * `upgrading`: The model is currently being upgraded.
// * `failed`: Training of the model failed.
const (
	LanguageModelStatusAvailableConst = "available"
	LanguageModelStatusFailedConst    = "failed"
	LanguageModelStatusPendingConst   = "pending"
	LanguageModelStatusReadyConst     = "ready"
	LanguageModelStatusTrainingConst  = "training"
	LanguageModelStatusUpgradingConst = "upgrading"
)

// UnmarshalLanguageModel unmarshals an instance of LanguageModel from the specified map of raw messages.
func UnmarshalLanguageModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LanguageModel)
	err = core.UnmarshalPrimitive(m, "customization_id", &obj.CustomizationID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created", &obj.Created)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated", &obj.Updated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "language", &obj.Language)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "dialect", &obj.Dialect)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "versions", &obj.Versions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "owner", &obj.Owner)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "base_model_name", &obj.BaseModelName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "progress", &obj.Progress)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "error", &obj.Error)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "warnings", &obj.Warnings)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LanguageModels : Information about existing custom language models.
type LanguageModels struct {
	// An array of `LanguageModel` objects that provides information about each available custom language model. The array
	// is empty if the requesting credentials own no custom language models (if no language is specified) or own no custom
	// language models for the specified language.
	Customizations []LanguageModel `json:"customizations" validate:"required"`
}

// UnmarshalLanguageModels unmarshals an instance of LanguageModels from the specified map of raw messages.
func UnmarshalLanguageModels(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LanguageModels)
	err = core.UnmarshalModel(m, "customizations", &obj.Customizations, UnmarshalLanguageModel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListAcousticModelsOptions : The ListAcousticModels options.
type ListAcousticModelsOptions struct {
	// The identifier of the language for which custom language or custom acoustic models are to be returned. Specify the
	// five-character language identifier; for example, specify `en-US` to see all custom language or custom acoustic
	// models that are based on US English models. Omit the parameter to see all custom language or custom acoustic models
	// that are owned by the requesting credentials. (**Note:** The identifier `ar-AR` is deprecated; use `ar-MS` instead.)
	//
	//
	// To determine the languages for which customization is available, see [Language support for
	// customization](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-custom-support).
	Language *string `json:"language,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListAcousticModelsOptions.Language property.
// The identifier of the language for which custom language or custom acoustic models are to be returned. Specify the
// five-character language identifier; for example, specify `en-US` to see all custom language or custom acoustic models
// that are based on US English models. Omit the parameter to see all custom language or custom acoustic models that are
// owned by the requesting credentials. (**Note:** The identifier `ar-AR` is deprecated; use `ar-MS` instead.)
//
// To determine the languages for which customization is available, see [Language support for
// customization](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-custom-support).
const (
	ListAcousticModelsOptionsLanguageArArConst = "ar-AR"
	ListAcousticModelsOptionsLanguageArMsConst = "ar-MS"
	ListAcousticModelsOptionsLanguageCsCzConst = "cs-CZ"
	ListAcousticModelsOptionsLanguageDeDeConst = "de-DE"
	ListAcousticModelsOptionsLanguageEnAuConst = "en-AU"
	ListAcousticModelsOptionsLanguageEnGbConst = "en-GB"
	ListAcousticModelsOptionsLanguageEnInConst = "en-IN"
	ListAcousticModelsOptionsLanguageEnUsConst = "en-US"
	ListAcousticModelsOptionsLanguageEnWwConst = "en-WW"
	ListAcousticModelsOptionsLanguageEsArConst = "es-AR"
	ListAcousticModelsOptionsLanguageEsClConst = "es-CL"
	ListAcousticModelsOptionsLanguageEsCoConst = "es-CO"
	ListAcousticModelsOptionsLanguageEsEsConst = "es-ES"
	ListAcousticModelsOptionsLanguageEsLaConst = "es-LA"
	ListAcousticModelsOptionsLanguageEsMxConst = "es-MX"
	ListAcousticModelsOptionsLanguageEsPeConst = "es-PE"
	ListAcousticModelsOptionsLanguageFrCaConst = "fr-CA"
	ListAcousticModelsOptionsLanguageFrFrConst = "fr-FR"
	ListAcousticModelsOptionsLanguageHiInConst = "hi-IN"
	ListAcousticModelsOptionsLanguageItItConst = "it-IT"
	ListAcousticModelsOptionsLanguageJaJpConst = "ja-JP"
	ListAcousticModelsOptionsLanguageKoKrConst = "ko-KR"
	ListAcousticModelsOptionsLanguageNlBeConst = "nl-BE"
	ListAcousticModelsOptionsLanguageNlNlConst = "nl-NL"
	ListAcousticModelsOptionsLanguagePtBrConst = "pt-BR"
	ListAcousticModelsOptionsLanguageZhCnConst = "zh-CN"
)

// NewListAcousticModelsOptions : Instantiate ListAcousticModelsOptions
func (*SpeechToTextV1) NewListAcousticModelsOptions() *ListAcousticModelsOptions {
	return &ListAcousticModelsOptions{}
}

// SetLanguage : Allow user to set Language
func (_options *ListAcousticModelsOptions) SetLanguage(language string) *ListAcousticModelsOptions {
	_options.Language = core.StringPtr(language)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListAcousticModelsOptions) SetHeaders(param map[string]string) *ListAcousticModelsOptions {
	options.Headers = param
	return options
}

// ListAudioOptions : The ListAudio options.
type ListAudioOptions struct {
	// The customization ID (GUID) of the custom acoustic model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListAudioOptions : Instantiate ListAudioOptions
func (*SpeechToTextV1) NewListAudioOptions(customizationID string) *ListAudioOptions {
	return &ListAudioOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *ListAudioOptions) SetCustomizationID(customizationID string) *ListAudioOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListAudioOptions) SetHeaders(param map[string]string) *ListAudioOptions {
	options.Headers = param
	return options
}

// ListCorporaOptions : The ListCorpora options.
type ListCorporaOptions struct {
	// The customization ID (GUID) of the custom language model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListCorporaOptions : Instantiate ListCorporaOptions
func (*SpeechToTextV1) NewListCorporaOptions(customizationID string) *ListCorporaOptions {
	return &ListCorporaOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *ListCorporaOptions) SetCustomizationID(customizationID string) *ListCorporaOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListCorporaOptions) SetHeaders(param map[string]string) *ListCorporaOptions {
	options.Headers = param
	return options
}

// ListGrammarsOptions : The ListGrammars options.
type ListGrammarsOptions struct {
	// The customization ID (GUID) of the custom language model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListGrammarsOptions : Instantiate ListGrammarsOptions
func (*SpeechToTextV1) NewListGrammarsOptions(customizationID string) *ListGrammarsOptions {
	return &ListGrammarsOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *ListGrammarsOptions) SetCustomizationID(customizationID string) *ListGrammarsOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListGrammarsOptions) SetHeaders(param map[string]string) *ListGrammarsOptions {
	options.Headers = param
	return options
}

// ListLanguageModelsOptions : The ListLanguageModels options.
type ListLanguageModelsOptions struct {
	// The identifier of the language for which custom language or custom acoustic models are to be returned. Specify the
	// five-character language identifier; for example, specify `en-US` to see all custom language or custom acoustic
	// models that are based on US English models. Omit the parameter to see all custom language or custom acoustic models
	// that are owned by the requesting credentials. (**Note:** The identifier `ar-AR` is deprecated; use `ar-MS` instead.)
	//
	//
	// To determine the languages for which customization is available, see [Language support for
	// customization](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-custom-support).
	Language *string `json:"language,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListLanguageModelsOptions.Language property.
// The identifier of the language for which custom language or custom acoustic models are to be returned. Specify the
// five-character language identifier; for example, specify `en-US` to see all custom language or custom acoustic models
// that are based on US English models. Omit the parameter to see all custom language or custom acoustic models that are
// owned by the requesting credentials. (**Note:** The identifier `ar-AR` is deprecated; use `ar-MS` instead.)
//
// To determine the languages for which customization is available, see [Language support for
// customization](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-custom-support).
const (
	ListLanguageModelsOptionsLanguageArArConst = "ar-AR"
	ListLanguageModelsOptionsLanguageArMsConst = "ar-MS"
	ListLanguageModelsOptionsLanguageCsCzConst = "cs-CZ"
	ListLanguageModelsOptionsLanguageDeDeConst = "de-DE"
	ListLanguageModelsOptionsLanguageEnAuConst = "en-AU"
	ListLanguageModelsOptionsLanguageEnGbConst = "en-GB"
	ListLanguageModelsOptionsLanguageEnInConst = "en-IN"
	ListLanguageModelsOptionsLanguageEnUsConst = "en-US"
	ListLanguageModelsOptionsLanguageEnWwConst = "en-WW"
	ListLanguageModelsOptionsLanguageEsArConst = "es-AR"
	ListLanguageModelsOptionsLanguageEsClConst = "es-CL"
	ListLanguageModelsOptionsLanguageEsCoConst = "es-CO"
	ListLanguageModelsOptionsLanguageEsEsConst = "es-ES"
	ListLanguageModelsOptionsLanguageEsLaConst = "es-LA"
	ListLanguageModelsOptionsLanguageEsMxConst = "es-MX"
	ListLanguageModelsOptionsLanguageEsPeConst = "es-PE"
	ListLanguageModelsOptionsLanguageFrCaConst = "fr-CA"
	ListLanguageModelsOptionsLanguageFrFrConst = "fr-FR"
	ListLanguageModelsOptionsLanguageHiInConst = "hi-IN"
	ListLanguageModelsOptionsLanguageItItConst = "it-IT"
	ListLanguageModelsOptionsLanguageJaJpConst = "ja-JP"
	ListLanguageModelsOptionsLanguageKoKrConst = "ko-KR"
	ListLanguageModelsOptionsLanguageNlBeConst = "nl-BE"
	ListLanguageModelsOptionsLanguageNlNlConst = "nl-NL"
	ListLanguageModelsOptionsLanguagePtBrConst = "pt-BR"
	ListLanguageModelsOptionsLanguageZhCnConst = "zh-CN"
)

// NewListLanguageModelsOptions : Instantiate ListLanguageModelsOptions
func (*SpeechToTextV1) NewListLanguageModelsOptions() *ListLanguageModelsOptions {
	return &ListLanguageModelsOptions{}
}

// SetLanguage : Allow user to set Language
func (_options *ListLanguageModelsOptions) SetLanguage(language string) *ListLanguageModelsOptions {
	_options.Language = core.StringPtr(language)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListLanguageModelsOptions) SetHeaders(param map[string]string) *ListLanguageModelsOptions {
	options.Headers = param
	return options
}

// ListModelsOptions : The ListModels options.
type ListModelsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListModelsOptions : Instantiate ListModelsOptions
func (*SpeechToTextV1) NewListModelsOptions() *ListModelsOptions {
	return &ListModelsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListModelsOptions) SetHeaders(param map[string]string) *ListModelsOptions {
	options.Headers = param
	return options
}

// ListWordsOptions : The ListWords options.
type ListWordsOptions struct {
	// The customization ID (GUID) of the custom language model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// The type of words to be listed from the custom language model's words resource:
	// * `all` (the default) shows all words.
	// * `user` shows only custom words that were added or modified by the user directly.
	// * `corpora` shows only OOV that were extracted from corpora.
	// * `grammars` shows only OOV words that are recognized by grammars.
	//
	// _For a custom model that is based on a next-generation model_, only `all` and `user` apply. Both options return the
	// same results. Words from other sources are not added to custom models that are based on next-generation models.
	WordType *string `json:"word_type,omitempty"`

	// Indicates the order in which the words are to be listed, `alphabetical` or by `count`. You can prepend an optional
	// `+` or `-` to an argument to indicate whether the results are to be sorted in ascending or descending order. By
	// default, words are sorted in ascending alphabetical order. For alphabetical ordering, the lexicographical precedence
	// is numeric values, uppercase letters, and lowercase letters. For count ordering, values with the same count are
	// ordered alphabetically. With the `curl` command, URL-encode the `+` symbol as `%2B`.
	Sort *string `json:"sort,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListWordsOptions.WordType property.
// The type of words to be listed from the custom language model's words resource:
// * `all` (the default) shows all words.
// * `user` shows only custom words that were added or modified by the user directly.
// * `corpora` shows only OOV that were extracted from corpora.
// * `grammars` shows only OOV words that are recognized by grammars.
//
// _For a custom model that is based on a next-generation model_, only `all` and `user` apply. Both options return the
// same results. Words from other sources are not added to custom models that are based on next-generation models.
const (
	ListWordsOptionsWordTypeAllConst      = "all"
	ListWordsOptionsWordTypeCorporaConst  = "corpora"
	ListWordsOptionsWordTypeGrammarsConst = "grammars"
	ListWordsOptionsWordTypeUserConst     = "user"
)

// Constants associated with the ListWordsOptions.Sort property.
// Indicates the order in which the words are to be listed, `alphabetical` or by `count`. You can prepend an optional
// `+` or `-` to an argument to indicate whether the results are to be sorted in ascending or descending order. By
// default, words are sorted in ascending alphabetical order. For alphabetical ordering, the lexicographical precedence
// is numeric values, uppercase letters, and lowercase letters. For count ordering, values with the same count are
// ordered alphabetically. With the `curl` command, URL-encode the `+` symbol as `%2B`.
const (
	ListWordsOptionsSortAlphabeticalConst = "alphabetical"
	ListWordsOptionsSortCountConst        = "count"
)

// NewListWordsOptions : Instantiate ListWordsOptions
func (*SpeechToTextV1) NewListWordsOptions(customizationID string) *ListWordsOptions {
	return &ListWordsOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *ListWordsOptions) SetCustomizationID(customizationID string) *ListWordsOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetWordType : Allow user to set WordType
func (_options *ListWordsOptions) SetWordType(wordType string) *ListWordsOptions {
	_options.WordType = core.StringPtr(wordType)
	return _options
}

// SetSort : Allow user to set Sort
func (_options *ListWordsOptions) SetSort(sort string) *ListWordsOptions {
	_options.Sort = core.StringPtr(sort)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListWordsOptions) SetHeaders(param map[string]string) *ListWordsOptions {
	options.Headers = param
	return options
}

// ProcessedAudio : Detailed timing information about the service's processing of the input audio.
type ProcessedAudio struct {
	// The seconds of audio that the service has received as of this response. The value of the field is greater than the
	// values of the `transcription` and `speaker_labels` fields during speech recognition processing, since the service
	// first has to receive the audio before it can begin to process it. The final value can also be greater than the value
	// of the `transcription` and `speaker_labels` fields by a fractional number of seconds.
	Received *float32 `json:"received" validate:"required"`

	// The seconds of audio that the service has passed to its speech-processing engine as of this response. The value of
	// the field is greater than the values of the `transcription` and `speaker_labels` fields during speech recognition
	// processing. The `received` and `seen_by_engine` fields have identical values when the service has finished
	// processing all audio. This final value can be greater than the value of the `transcription` and `speaker_labels`
	// fields by a fractional number of seconds.
	SeenByEngine *float32 `json:"seen_by_engine" validate:"required"`

	// The seconds of audio that the service has processed for speech recognition as of this response.
	Transcription *float32 `json:"transcription" validate:"required"`

	// If speaker labels are requested, the seconds of audio that the service has processed to determine speaker labels as
	// of this response. This value often trails the value of the `transcription` field during speech recognition
	// processing. The `transcription` and `speaker_labels` fields have identical values when the service has finished
	// processing all audio.
	SpeakerLabels *float32 `json:"speaker_labels,omitempty"`
}

// UnmarshalProcessedAudio unmarshals an instance of ProcessedAudio from the specified map of raw messages.
func UnmarshalProcessedAudio(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProcessedAudio)
	err = core.UnmarshalPrimitive(m, "received", &obj.Received)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "seen_by_engine", &obj.SeenByEngine)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "transcription", &obj.Transcription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "speaker_labels", &obj.SpeakerLabels)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProcessingMetrics : If processing metrics are requested, information about the service's processing of the input audio. Processing
// metrics are not available with the synchronous [Recognize audio](#recognize) method.
type ProcessingMetrics struct {
	// Detailed timing information about the service's processing of the input audio.
	ProcessedAudio *ProcessedAudio `json:"processed_audio" validate:"required"`

	// The amount of real time in seconds that has passed since the service received the first byte of input audio. Values
	// in this field are generally multiples of the specified metrics interval, with two differences:
	// * Values might not reflect exact intervals (for instance, 0.25, 0.5, and so on). Actual values might be 0.27, 0.52,
	// and so on, depending on when the service receives and processes audio.
	// * The service also returns values for transcription events if you set the `interim_results` parameter to `true`. The
	// service returns both processing metrics and transcription results when such events occur.
	WallClockSinceFirstByteReceived *float32 `json:"wall_clock_since_first_byte_received" validate:"required"`

	// An indication of whether the metrics apply to a periodic interval or a transcription event:
	// * `true` means that the response was triggered by a specified processing interval. The information contains
	// processing metrics only.
	// * `false` means that the response was triggered by a transcription event. The information contains processing
	// metrics plus transcription results.
	//
	// Use the field to identify why the service generated the response and to filter different results if necessary.
	Periodic *bool `json:"periodic" validate:"required"`
}

// UnmarshalProcessingMetrics unmarshals an instance of ProcessingMetrics from the specified map of raw messages.
func UnmarshalProcessingMetrics(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProcessingMetrics)
	err = core.UnmarshalModel(m, "processed_audio", &obj.ProcessedAudio, UnmarshalProcessedAudio)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "wall_clock_since_first_byte_received", &obj.WallClockSinceFirstByteReceived)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "periodic", &obj.Periodic)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RecognitionJob : Information about a current asynchronous speech recognition job.
type RecognitionJob struct {
	// The ID of the asynchronous job.
	ID *string `json:"id" validate:"required"`

	// The current status of the job:
	// * `waiting`: The service is preparing the job for processing. The service returns this status when the job is
	// initially created or when it is waiting for capacity to process the job. The job remains in this state until the
	// service has the capacity to begin processing it.
	// * `processing`: The service is actively processing the job.
	// * `completed`: The service has finished processing the job. If the job specified a callback URL and the event
	// `recognitions.completed_with_results`, the service sent the results with the callback notification. Otherwise, you
	// must retrieve the results by checking the individual job.
	// * `failed`: The job failed.
	Status *string `json:"status" validate:"required"`

	// The date and time in Coordinated Universal Time (UTC) at which the job was created. The value is provided in full
	// ISO 8601 format (`YYYY-MM-DDThh:mm:ss.sTZD`).
	Created *string `json:"created" validate:"required"`

	// The date and time in Coordinated Universal Time (UTC) at which the job was last updated by the service. The value is
	// provided in full ISO 8601 format (`YYYY-MM-DDThh:mm:ss.sTZD`). This field is returned only by the [Check
	// jobs](#checkjobs) and [Check a job[(#checkjob) methods.
	Updated *string `json:"updated,omitempty"`

	// The URL to use to request information about the job with the [Check a job](#checkjob) method. This field is returned
	// only by the [Create a job](#createjob) method.
	URL *string `json:"url,omitempty"`

	// The user token associated with a job that was created with a callback URL and a user token. This field can be
	// returned only by the [Check jobs](#checkjobs) method.
	UserToken *string `json:"user_token,omitempty"`

	// If the status is `completed`, the results of the recognition request as an array that includes a single instance of
	// a `SpeechRecognitionResults` object. This field is returned only by the [Check a job](#checkjob) method.
	Results []SpeechRecognitionResults `json:"results,omitempty"`

	// An array of warning messages about invalid parameters included with the request. Each warning includes a descriptive
	// message and a list of invalid argument strings, for example, `"unexpected query parameter 'user_token', query
	// parameter 'callback_url' was not specified"`. The request succeeds despite the warnings. This field can be returned
	// only by the [Create a job](#createjob) method.
	Warnings []string `json:"warnings,omitempty"`
}

// Constants associated with the RecognitionJob.Status property.
// The current status of the job:
// * `waiting`: The service is preparing the job for processing. The service returns this status when the job is
// initially created or when it is waiting for capacity to process the job. The job remains in this state until the
// service has the capacity to begin processing it.
// * `processing`: The service is actively processing the job.
// * `completed`: The service has finished processing the job. If the job specified a callback URL and the event
// `recognitions.completed_with_results`, the service sent the results with the callback notification. Otherwise, you
// must retrieve the results by checking the individual job.
// * `failed`: The job failed.
const (
	RecognitionJobStatusCompletedConst  = "completed"
	RecognitionJobStatusFailedConst     = "failed"
	RecognitionJobStatusProcessingConst = "processing"
	RecognitionJobStatusWaitingConst    = "waiting"
)

// UnmarshalRecognitionJob unmarshals an instance of RecognitionJob from the specified map of raw messages.
func UnmarshalRecognitionJob(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RecognitionJob)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created", &obj.Created)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated", &obj.Updated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "user_token", &obj.UserToken)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "results", &obj.Results, UnmarshalSpeechRecognitionResults)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "warnings", &obj.Warnings)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RecognitionJobs : Information about current asynchronous speech recognition jobs.
type RecognitionJobs struct {
	// An array of `RecognitionJob` objects that provides the status for each of the user's current jobs. The array is
	// empty if the user has no current jobs.
	Recognitions []RecognitionJob `json:"recognitions" validate:"required"`
}

// UnmarshalRecognitionJobs unmarshals an instance of RecognitionJobs from the specified map of raw messages.
func UnmarshalRecognitionJobs(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RecognitionJobs)
	err = core.UnmarshalModel(m, "recognitions", &obj.Recognitions, UnmarshalRecognitionJob)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RecognizeOptions : The Recognize options.
type RecognizeOptions struct {
	// The audio to transcribe.
	Audio io.ReadCloser `json:"audio" validate:"required"`

	// The format (MIME type) of the audio. For more information about specifying an audio format, see **Audio formats
	// (content types)** in the method description.
	ContentType *string `json:"content-type,omitempty"`

	// The identifier of the model that is to be used for the recognition request. (**Note:** The model
	// `ar-AR_BroadbandModel` is deprecated; use `ar-MS_BroadbandModel` instead.) See [Using a model for speech
	// recognition](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-models-use).
	Model *string `json:"model,omitempty"`

	// The customization ID (GUID) of a custom language model that is to be used with the recognition request. The base
	// model of the specified custom language model must match the model specified with the `model` parameter. You must
	// make the request with credentials for the instance of the service that owns the custom model. By default, no custom
	// language model is used. See [Using a custom language model for speech
	// recognition](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-languageUse).
	//
	// **Note:** Use this parameter instead of the deprecated `customization_id` parameter.
	LanguageCustomizationID *string `json:"language_customization_id,omitempty"`

	// The customization ID (GUID) of a custom acoustic model that is to be used with the recognition request. The base
	// model of the specified custom acoustic model must match the model specified with the `model` parameter. You must
	// make the request with credentials for the instance of the service that owns the custom model. By default, no custom
	// acoustic model is used. See [Using a custom acoustic model for speech
	// recognition](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-acousticUse).
	AcousticCustomizationID *string `json:"acoustic_customization_id,omitempty"`

	// The version of the specified base model that is to be used with the recognition request. Multiple versions of a base
	// model can exist when a model is updated for internal improvements. The parameter is intended primarily for use with
	// custom models that have been upgraded for a new base model. The default value depends on whether the parameter is
	// used with or without a custom model. See [Making speech recognition requests with upgraded custom
	// models](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-custom-upgrade-use#custom-upgrade-use-recognition).
	BaseModelVersion *string `json:"base_model_version,omitempty"`

	// If you specify the customization ID (GUID) of a custom language model with the recognition request, the
	// customization weight tells the service how much weight to give to words from the custom language model compared to
	// those from the base model for the current request.
	//
	// Specify a value between 0.0 and 1.0. Unless a different customization weight was specified for the custom model when
	// it was trained, the default value is 0.3. A customization weight that you specify overrides a weight that was
	// specified when the custom model was trained.
	//
	// The default value yields the best performance in general. Assign a higher value if your audio makes frequent use of
	// OOV words from the custom model. Use caution when setting the weight: a higher value can improve the accuracy of
	// phrases from the custom model's domain, but it can negatively affect performance on non-domain phrases.
	//
	// See [Using customization weight](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-languageUse#weight).
	CustomizationWeight *float64 `json:"customization_weight,omitempty"`

	// The time in seconds after which, if only silence (no speech) is detected in streaming audio, the connection is
	// closed with a 400 error. The parameter is useful for stopping audio submission from a live microphone when a user
	// simply walks away. Use `-1` for infinity. See [Inactivity
	// timeout](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-input#timeouts-inactivity).
	InactivityTimeout *int64 `json:"inactivity_timeout,omitempty"`

	// An array of keyword strings to spot in the audio. Each keyword string can include one or more string tokens.
	// Keywords are spotted only in the final results, not in interim hypotheses. If you specify any keywords, you must
	// also specify a keywords threshold. Omit the parameter or specify an empty array if you do not need to spot keywords.
	//
	//
	// You can spot a maximum of 1000 keywords with a single request. A single keyword can have a maximum length of 1024
	// characters, though the maximum effective length for double-byte languages might be shorter. Keywords are
	// case-insensitive.
	//
	// See [Keyword spotting](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-spotting#keyword-spotting).
	Keywords []string `json:"keywords,omitempty"`

	// A confidence value that is the lower bound for spotting a keyword. A word is considered to match a keyword if its
	// confidence is greater than or equal to the threshold. Specify a probability between 0.0 and 1.0. If you specify a
	// threshold, you must also specify one or more keywords. The service performs no keyword spotting if you omit either
	// parameter. See [Keyword
	// spotting](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-spotting#keyword-spotting).
	KeywordsThreshold *float32 `json:"keywords_threshold,omitempty"`

	// The maximum number of alternative transcripts that the service is to return. By default, the service returns a
	// single transcript. If you specify a value of `0`, the service uses the default value, `1`. See [Maximum
	// alternatives](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-metadata#max-alternatives).
	MaxAlternatives *int64 `json:"max_alternatives,omitempty"`

	// A confidence value that is the lower bound for identifying a hypothesis as a possible word alternative (also known
	// as "Confusion Networks"). An alternative word is considered if its confidence is greater than or equal to the
	// threshold. Specify a probability between 0.0 and 1.0. By default, the service computes no alternative words. See
	// [Word alternatives](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-spotting#word-alternatives).
	WordAlternativesThreshold *float32 `json:"word_alternatives_threshold,omitempty"`

	// If `true`, the service returns a confidence measure in the range of 0.0 to 1.0 for each word. By default, the
	// service returns no word confidence scores. See [Word
	// confidence](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-metadata#word-confidence).
	WordConfidence *bool `json:"word_confidence,omitempty"`

	// If `true`, the service returns time alignment for each word. By default, no timestamps are returned. See [Word
	// timestamps](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-metadata#word-timestamps).
	Timestamps *bool `json:"timestamps,omitempty"`

	// If `true`, the service filters profanity from all output except for keyword results by replacing inappropriate words
	// with a series of asterisks. Set the parameter to `false` to return results with no censoring.
	//
	// **Note:** The parameter can be used with US English and Japanese transcription only. See [Profanity
	// filtering](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-formatting#profanity-filtering).
	ProfanityFilter *bool `json:"profanity_filter,omitempty"`

	// If `true`, the service converts dates, times, series of digits and numbers, phone numbers, currency values, and
	// internet addresses into more readable, conventional representations in the final transcript of a recognition
	// request. For US English, the service also converts certain keyword strings to punctuation symbols. By default, the
	// service performs no smart formatting.
	//
	// **Note:** The parameter can be used with US English, Japanese, and Spanish (all dialects) transcription only.
	//
	// See [Smart formatting](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-formatting#smart-formatting).
	SmartFormatting *bool `json:"smart_formatting,omitempty"`

	// If `true`, the response includes labels that identify which words were spoken by which participants in a
	// multi-person exchange. By default, the service returns no speaker labels. Setting `speaker_labels` to `true` forces
	// the `timestamps` parameter to be `true`, regardless of whether you specify `false` for the parameter.
	// * _For previous-generation models,_ the parameter can be used with Australian English, US English, German, Japanese,
	// Korean, and Spanish (both broadband and narrowband models) and UK English (narrowband model) transcription only.
	// * _For next-generation models,_ the parameter can be used with Czech, English (Australian, Indian, UK, and US),
	// German, Japanese, Korean, and Spanish transcription only.
	//
	// See [Speaker labels](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-speaker-labels).
	SpeakerLabels *bool `json:"speaker_labels,omitempty"`

	// **Deprecated.** Use the `language_customization_id` parameter to specify the customization ID (GUID) of a custom
	// language model that is to be used with the recognition request. Do not specify both parameters with a request.
	CustomizationID *string `json:"customization_id,omitempty"`

	// The name of a grammar that is to be used with the recognition request. If you specify a grammar, you must also use
	// the `language_customization_id` parameter to specify the name of the custom language model for which the grammar is
	// defined. The service recognizes only strings that are recognized by the specified grammar; it does not recognize
	// other custom words from the model's words resource.
	//
	// See [Using a grammar for speech
	// recognition](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-grammarUse).
	GrammarName *string `json:"grammar_name,omitempty"`

	// If `true`, the service redacts, or masks, numeric data from final transcripts. The feature redacts any number that
	// has three or more consecutive digits by replacing each digit with an `X` character. It is intended to redact
	// sensitive numeric data, such as credit card numbers. By default, the service performs no redaction.
	//
	// When you enable redaction, the service automatically enables smart formatting, regardless of whether you explicitly
	// disable that feature. To ensure maximum security, the service also disables keyword spotting (ignores the `keywords`
	// and `keywords_threshold` parameters) and returns only a single final transcript (forces the `max_alternatives`
	// parameter to be `1`).
	//
	// **Note:** The parameter can be used with US English, Japanese, and Korean transcription only.
	//
	// See [Numeric
	// redaction](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-formatting#numeric-redaction).
	Redaction *bool `json:"redaction,omitempty"`

	// If `true`, requests detailed information about the signal characteristics of the input audio. The service returns
	// audio metrics with the final transcription results. By default, the service returns no audio metrics.
	//
	// See [Audio metrics](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-metrics#audio-metrics).
	AudioMetrics *bool `json:"audio_metrics,omitempty"`

	// If `true`, specifies the duration of the pause interval at which the service splits a transcript into multiple final
	// results. If the service detects pauses or extended silence before it reaches the end of the audio stream, its
	// response can include multiple final results. Silence indicates a point at which the speaker pauses between spoken
	// words or phrases.
	//
	// Specify a value for the pause interval in the range of 0.0 to 120.0.
	// * A value greater than 0 specifies the interval that the service is to use for speech recognition.
	// * A value of 0 indicates that the service is to use the default interval. It is equivalent to omitting the
	// parameter.
	//
	// The default pause interval for most languages is 0.8 seconds; the default for Chinese is 0.6 seconds.
	//
	// See [End of phrase silence
	// time](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-parsing#silence-time).
	EndOfPhraseSilenceTime *float64 `json:"end_of_phrase_silence_time,omitempty"`

	// If `true`, directs the service to split the transcript into multiple final results based on semantic features of the
	// input, for example, at the conclusion of meaningful phrases such as sentences. The service bases its understanding
	// of semantic features on the base language model that you use with a request. Custom language models and grammars can
	// also influence how and where the service splits a transcript.
	//
	// By default, the service splits transcripts based solely on the pause interval. If the parameters are used together
	// on the same request, `end_of_phrase_silence_time` has precedence over `split_transcript_at_phrase_end`.
	//
	// See [Split transcript at phrase
	// end](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-parsing#split-transcript).
	SplitTranscriptAtPhraseEnd *bool `json:"split_transcript_at_phrase_end,omitempty"`

	// The sensitivity of speech activity detection that the service is to perform. Use the parameter to suppress word
	// insertions from music, coughing, and other non-speech events. The service biases the audio it passes for speech
	// recognition by evaluating the input audio against prior models of speech and non-speech activity.
	//
	// Specify a value between 0.0 and 1.0:
	// * 0.0 suppresses all audio (no speech is transcribed).
	// * 0.5 (the default) provides a reasonable compromise for the level of sensitivity.
	// * 1.0 suppresses no audio (speech detection sensitivity is disabled).
	//
	// The values increase on a monotonic curve.
	//
	// The parameter is supported with all next-generation models and with most previous-generation models. See [Speech
	// detector
	// sensitivity](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-detection#detection-parameters-sensitivity)
	// and [Language model
	// support](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-detection#detection-support).
	SpeechDetectorSensitivity *float32 `json:"speech_detector_sensitivity,omitempty"`

	// The level to which the service is to suppress background audio based on its volume to prevent it from being
	// transcribed as speech. Use the parameter to suppress side conversations or background noise.
	//
	// Specify a value in the range of 0.0 to 1.0:
	// * 0.0 (the default) provides no suppression (background audio suppression is disabled).
	// * 0.5 provides a reasonable level of audio suppression for general usage.
	// * 1.0 suppresses all audio (no audio is transcribed).
	//
	// The values increase on a monotonic curve.
	//
	// The parameter is supported with all next-generation models and with most previous-generation models. See [Background
	// audio
	// suppression](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-detection#detection-parameters-suppression)
	// and [Language model
	// support](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-detection#detection-support).
	BackgroundAudioSuppression *float32 `json:"background_audio_suppression,omitempty"`

	// The character_insertion_bias parameter controls the service's bias for competing strings of different lengths
	// during speech recognition. With next-generation models, the service parses audio character by character.
	// As it does, it establishes hypotheses of previous character strings to help determine viable next characters.
	// During this process, it collects candidate strings of different lengths.
	//
	// By default, each model uses a default character_insertion_bias of 0.0.
	// This value is optimized to produce the best balance between hypotheses with different numbers of characters.
	// The default is typically adequate for most speech recognition.
	// However, certain use cases might benefit from favoring hypotheses with shorter or longer strings of characters.
	// In such cases, specifying a change from the default can improve speech recognition.
	//
	// You can use the character_insertion_bias parameter to indicate that the service is to favor shorter or longer
	//  strings as it considers subsequent characters for its hypotheses.
	// The value you provide depends on the characteristics of your audio.
	// The range of acceptable values is from -1.0 to 1.0:
	//
	// Negative values cause the service to prefer hypotheses with shorter strings of characters.
	// Positive values cause the service to prefer hypotheses with longer strings of characters.
	// As your value approaches -1.0 or 1.0, the impact of the parameter becomes more pronounced.
	// To determine the most effective value for your scenario, start by setting the value of the parameter
	// to a small increment, such as -0.1, -0.05, 0.05, or 0.1, and assess how the value impacts the transcription results.
	//
	// The parameter is not available for previous-generation models.
	CharacterInsertionBias *float32 `json:"character_insertion_bias,omitempty"`

	// If `true` for next-generation `Multimedia` and `Telephony` models that support low latency, directs the service to
	// produce results even more quickly than it usually does. Next-generation models produce transcription results faster
	// than previous-generation models. The `low_latency` parameter causes the models to produce results even more quickly,
	// though the results might be less accurate when the parameter is used.
	//
	// The parameter is not available for previous-generation `Broadband` and `Narrowband` models. It is available only for
	// some next-generation models. For a list of next-generation models that support low latency, see [Supported
	// next-generation language
	// models](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-models-ng#models-ng-supported).
	// * For more information about the `low_latency` parameter, see [Low
	// latency](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-interim#low-latency).
	LowLatency *bool `json:"low_latency,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the RecognizeOptions.Model property.
// The identifier of the model that is to be used for the recognition request. (**Note:** The model
// `ar-AR_BroadbandModel` is deprecated; use `ar-MS_BroadbandModel` instead.) See [Using a model for speech
// recognition](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-models-use).
const (
	RecognizeOptionsModelArArBroadbandmodelConst           = "ar-AR_BroadbandModel"
	RecognizeOptionsModelArMsBroadbandmodelConst           = "ar-MS_BroadbandModel"
	RecognizeOptionsModelArMsTelephonyConst                = "ar-MS_Telephony"
	RecognizeOptionsModelCsCzTelephonyConst                = "cs-CZ_Telephony"
	RecognizeOptionsModelDeDeBroadbandmodelConst           = "de-DE_BroadbandModel"
	RecognizeOptionsModelDeDeMultimediaConst               = "de-DE_Multimedia"
	RecognizeOptionsModelDeDeNarrowbandmodelConst          = "de-DE_NarrowbandModel"
	RecognizeOptionsModelDeDeTelephonyConst                = "de-DE_Telephony"
	RecognizeOptionsModelEnAuBroadbandmodelConst           = "en-AU_BroadbandModel"
	RecognizeOptionsModelEnAuMultimediaConst               = "en-AU_Multimedia"
	RecognizeOptionsModelEnAuNarrowbandmodelConst          = "en-AU_NarrowbandModel"
	RecognizeOptionsModelEnAuTelephonyConst                = "en-AU_Telephony"
	RecognizeOptionsModelEnGbBroadbandmodelConst           = "en-GB_BroadbandModel"
	RecognizeOptionsModelEnGbMultimediaConst               = "en-GB_Multimedia"
	RecognizeOptionsModelEnGbNarrowbandmodelConst          = "en-GB_NarrowbandModel"
	RecognizeOptionsModelEnGbTelephonyConst                = "en-GB_Telephony"
	RecognizeOptionsModelEnInTelephonyConst                = "en-IN_Telephony"
	RecognizeOptionsModelEnUsBroadbandmodelConst           = "en-US_BroadbandModel"
	RecognizeOptionsModelEnUsMultimediaConst               = "en-US_Multimedia"
	RecognizeOptionsModelEnUsNarrowbandmodelConst          = "en-US_NarrowbandModel"
	RecognizeOptionsModelEnUsShortformNarrowbandmodelConst = "en-US_ShortForm_NarrowbandModel"
	RecognizeOptionsModelEnUsTelephonyConst                = "en-US_Telephony"
	RecognizeOptionsModelEnWwMedicalTelephonyConst         = "en-WW_Medical_Telephony"
	RecognizeOptionsModelEsArBroadbandmodelConst           = "es-AR_BroadbandModel"
	RecognizeOptionsModelEsArNarrowbandmodelConst          = "es-AR_NarrowbandModel"
	RecognizeOptionsModelEsClBroadbandmodelConst           = "es-CL_BroadbandModel"
	RecognizeOptionsModelEsClNarrowbandmodelConst          = "es-CL_NarrowbandModel"
	RecognizeOptionsModelEsCoBroadbandmodelConst           = "es-CO_BroadbandModel"
	RecognizeOptionsModelEsCoNarrowbandmodelConst          = "es-CO_NarrowbandModel"
	RecognizeOptionsModelEsEsBroadbandmodelConst           = "es-ES_BroadbandModel"
	RecognizeOptionsModelEsEsMultimediaConst               = "es-ES_Multimedia"
	RecognizeOptionsModelEsEsNarrowbandmodelConst          = "es-ES_NarrowbandModel"
	RecognizeOptionsModelEsEsTelephonyConst                = "es-ES_Telephony"
	RecognizeOptionsModelEsLaTelephonyConst                = "es-LA_Telephony"
	RecognizeOptionsModelEsMxBroadbandmodelConst           = "es-MX_BroadbandModel"
	RecognizeOptionsModelEsMxNarrowbandmodelConst          = "es-MX_NarrowbandModel"
	RecognizeOptionsModelEsPeBroadbandmodelConst           = "es-PE_BroadbandModel"
	RecognizeOptionsModelEsPeNarrowbandmodelConst          = "es-PE_NarrowbandModel"
	RecognizeOptionsModelFrCaBroadbandmodelConst           = "fr-CA_BroadbandModel"
	RecognizeOptionsModelFrCaNarrowbandmodelConst          = "fr-CA_NarrowbandModel"
	RecognizeOptionsModelFrCaTelephonyConst                = "fr-CA_Telephony"
	RecognizeOptionsModelFrFrBroadbandmodelConst           = "fr-FR_BroadbandModel"
	RecognizeOptionsModelFrFrMultimediaConst               = "fr-FR_Multimedia"
	RecognizeOptionsModelFrFrNarrowbandmodelConst          = "fr-FR_NarrowbandModel"
	RecognizeOptionsModelFrFrTelephonyConst                = "fr-FR_Telephony"
	RecognizeOptionsModelHiInTelephonyConst                = "hi-IN_Telephony"
	RecognizeOptionsModelItItBroadbandmodelConst           = "it-IT_BroadbandModel"
	RecognizeOptionsModelItItNarrowbandmodelConst          = "it-IT_NarrowbandModel"
	RecognizeOptionsModelItItTelephonyConst                = "it-IT_Telephony"
	RecognizeOptionsModelJaJpBroadbandmodelConst           = "ja-JP_BroadbandModel"
	RecognizeOptionsModelJaJpMultimediaConst               = "ja-JP_Multimedia"
	RecognizeOptionsModelJaJpNarrowbandmodelConst          = "ja-JP_NarrowbandModel"
	RecognizeOptionsModelKoKrBroadbandmodelConst           = "ko-KR_BroadbandModel"
	RecognizeOptionsModelKoKrMultimediaConst               = "ko-KR_Multimedia"
	RecognizeOptionsModelKoKrNarrowbandmodelConst          = "ko-KR_NarrowbandModel"
	RecognizeOptionsModelKoKrTelephonyConst                = "ko-KR_Telephony"
	RecognizeOptionsModelNlBeTelephonyConst                = "nl-BE_Telephony"
	RecognizeOptionsModelNlNlBroadbandmodelConst           = "nl-NL_BroadbandModel"
	RecognizeOptionsModelNlNlNarrowbandmodelConst          = "nl-NL_NarrowbandModel"
	RecognizeOptionsModelNlNlTelephonyConst                = "nl-NL_Telephony"
	RecognizeOptionsModelPtBrBroadbandmodelConst           = "pt-BR_BroadbandModel"
	RecognizeOptionsModelPtBrNarrowbandmodelConst          = "pt-BR_NarrowbandModel"
	RecognizeOptionsModelPtBrTelephonyConst                = "pt-BR_Telephony"
	RecognizeOptionsModelZhCnBroadbandmodelConst           = "zh-CN_BroadbandModel"
	RecognizeOptionsModelZhCnNarrowbandmodelConst          = "zh-CN_NarrowbandModel"
	RecognizeOptionsModelZhCnTelephonyConst                = "zh-CN_Telephony"
)

// NewRecognizeOptions : Instantiate RecognizeOptions
func (*SpeechToTextV1) NewRecognizeOptions(audio io.ReadCloser) *RecognizeOptions {
	return &RecognizeOptions{
		Audio: audio,
	}
}

// SetAudio : Allow user to set Audio
func (_options *RecognizeOptions) SetAudio(audio io.ReadCloser) *RecognizeOptions {
	_options.Audio = audio
	return _options
}

// SetContentType : Allow user to set ContentType
func (_options *RecognizeOptions) SetContentType(contentType string) *RecognizeOptions {
	_options.ContentType = core.StringPtr(contentType)
	return _options
}

// SetModel : Allow user to set Model
func (_options *RecognizeOptions) SetModel(model string) *RecognizeOptions {
	_options.Model = core.StringPtr(model)
	return _options
}

// SetLanguageCustomizationID : Allow user to set LanguageCustomizationID
func (_options *RecognizeOptions) SetLanguageCustomizationID(languageCustomizationID string) *RecognizeOptions {
	_options.LanguageCustomizationID = core.StringPtr(languageCustomizationID)
	return _options
}

// SetAcousticCustomizationID : Allow user to set AcousticCustomizationID
func (_options *RecognizeOptions) SetAcousticCustomizationID(acousticCustomizationID string) *RecognizeOptions {
	_options.AcousticCustomizationID = core.StringPtr(acousticCustomizationID)
	return _options
}

// SetBaseModelVersion : Allow user to set BaseModelVersion
func (_options *RecognizeOptions) SetBaseModelVersion(baseModelVersion string) *RecognizeOptions {
	_options.BaseModelVersion = core.StringPtr(baseModelVersion)
	return _options
}

// SetCustomizationWeight : Allow user to set CustomizationWeight
func (_options *RecognizeOptions) SetCustomizationWeight(customizationWeight float64) *RecognizeOptions {
	_options.CustomizationWeight = core.Float64Ptr(customizationWeight)
	return _options
}

// SetInactivityTimeout : Allow user to set InactivityTimeout
func (_options *RecognizeOptions) SetInactivityTimeout(inactivityTimeout int64) *RecognizeOptions {
	_options.InactivityTimeout = core.Int64Ptr(inactivityTimeout)
	return _options
}

// SetKeywords : Allow user to set Keywords
func (_options *RecognizeOptions) SetKeywords(keywords []string) *RecognizeOptions {
	_options.Keywords = keywords
	return _options
}

// SetKeywordsThreshold : Allow user to set KeywordsThreshold
func (_options *RecognizeOptions) SetKeywordsThreshold(keywordsThreshold float32) *RecognizeOptions {
	_options.KeywordsThreshold = core.Float32Ptr(keywordsThreshold)
	return _options
}

// SetMaxAlternatives : Allow user to set MaxAlternatives
func (_options *RecognizeOptions) SetMaxAlternatives(maxAlternatives int64) *RecognizeOptions {
	_options.MaxAlternatives = core.Int64Ptr(maxAlternatives)
	return _options
}

// SetWordAlternativesThreshold : Allow user to set WordAlternativesThreshold
func (_options *RecognizeOptions) SetWordAlternativesThreshold(wordAlternativesThreshold float32) *RecognizeOptions {
	_options.WordAlternativesThreshold = core.Float32Ptr(wordAlternativesThreshold)
	return _options
}

// SetWordConfidence : Allow user to set WordConfidence
func (_options *RecognizeOptions) SetWordConfidence(wordConfidence bool) *RecognizeOptions {
	_options.WordConfidence = core.BoolPtr(wordConfidence)
	return _options
}

// SetTimestamps : Allow user to set Timestamps
func (_options *RecognizeOptions) SetTimestamps(timestamps bool) *RecognizeOptions {
	_options.Timestamps = core.BoolPtr(timestamps)
	return _options
}

// SetProfanityFilter : Allow user to set ProfanityFilter
func (_options *RecognizeOptions) SetProfanityFilter(profanityFilter bool) *RecognizeOptions {
	_options.ProfanityFilter = core.BoolPtr(profanityFilter)
	return _options
}

// SetSmartFormatting : Allow user to set SmartFormatting
func (_options *RecognizeOptions) SetSmartFormatting(smartFormatting bool) *RecognizeOptions {
	_options.SmartFormatting = core.BoolPtr(smartFormatting)
	return _options
}

// SetSpeakerLabels : Allow user to set SpeakerLabels
func (_options *RecognizeOptions) SetSpeakerLabels(speakerLabels bool) *RecognizeOptions {
	_options.SpeakerLabels = core.BoolPtr(speakerLabels)
	return _options
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *RecognizeOptions) SetCustomizationID(customizationID string) *RecognizeOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetGrammarName : Allow user to set GrammarName
func (_options *RecognizeOptions) SetGrammarName(grammarName string) *RecognizeOptions {
	_options.GrammarName = core.StringPtr(grammarName)
	return _options
}

// SetRedaction : Allow user to set Redaction
func (_options *RecognizeOptions) SetRedaction(redaction bool) *RecognizeOptions {
	_options.Redaction = core.BoolPtr(redaction)
	return _options
}

// SetAudioMetrics : Allow user to set AudioMetrics
func (_options *RecognizeOptions) SetAudioMetrics(audioMetrics bool) *RecognizeOptions {
	_options.AudioMetrics = core.BoolPtr(audioMetrics)
	return _options
}

// SetEndOfPhraseSilenceTime : Allow user to set EndOfPhraseSilenceTime
func (_options *RecognizeOptions) SetEndOfPhraseSilenceTime(endOfPhraseSilenceTime float64) *RecognizeOptions {
	_options.EndOfPhraseSilenceTime = core.Float64Ptr(endOfPhraseSilenceTime)
	return _options
}

// SetSplitTranscriptAtPhraseEnd : Allow user to set SplitTranscriptAtPhraseEnd
func (_options *RecognizeOptions) SetSplitTranscriptAtPhraseEnd(splitTranscriptAtPhraseEnd bool) *RecognizeOptions {
	_options.SplitTranscriptAtPhraseEnd = core.BoolPtr(splitTranscriptAtPhraseEnd)
	return _options
}

// SetSpeechDetectorSensitivity : Allow user to set SpeechDetectorSensitivity
func (_options *RecognizeOptions) SetSpeechDetectorSensitivity(speechDetectorSensitivity float32) *RecognizeOptions {
	_options.SpeechDetectorSensitivity = core.Float32Ptr(speechDetectorSensitivity)
	return _options
}

// SetBackgroundAudioSuppression : Allow user to set BackgroundAudioSuppression
func (_options *RecognizeOptions) SetBackgroundAudioSuppression(backgroundAudioSuppression float32) *RecognizeOptions {
	_options.BackgroundAudioSuppression = core.Float32Ptr(backgroundAudioSuppression)
	return _options
}

// SetCharacterInsertionBias : Allow user to set CharacterInsertionBias
func (_options *RecognizeOptions) SetCharacterInsertionBias(characterInsertionBias float32) *RecognizeOptions {
	_options.CharacterInsertionBias = core.Float32Ptr(characterInsertionBias)
	return _options
}

// SetLowLatency : Allow user to set LowLatency
func (_options *RecognizeOptions) SetLowLatency(lowLatency bool) *RecognizeOptions {
	_options.LowLatency = core.BoolPtr(lowLatency)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RecognizeOptions) SetHeaders(param map[string]string) *RecognizeOptions {
	options.Headers = param
	return options
}

// RegisterCallbackOptions : The RegisterCallback options.
type RegisterCallbackOptions struct {
	// An HTTP or HTTPS URL to which callback notifications are to be sent. To be allowlisted, the URL must successfully
	// echo the challenge string during URL verification. During verification, the client can also check the signature that
	// the service sends in the `X-Callback-Signature` header to verify the origin of the request.
	CallbackURL *string `json:"callback_url" validate:"required"`

	// A user-specified string that the service uses to generate the HMAC-SHA1 signature that it sends via the
	// `X-Callback-Signature` header. The service includes the header during URL verification and with every notification
	// sent to the callback URL. It calculates the signature over the payload of the notification. If you omit the
	// parameter, the service does not send the header.
	UserSecret *string `json:"user_secret,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRegisterCallbackOptions : Instantiate RegisterCallbackOptions
func (*SpeechToTextV1) NewRegisterCallbackOptions(callbackURL string) *RegisterCallbackOptions {
	return &RegisterCallbackOptions{
		CallbackURL: core.StringPtr(callbackURL),
	}
}

// SetCallbackURL : Allow user to set CallbackURL
func (_options *RegisterCallbackOptions) SetCallbackURL(callbackURL string) *RegisterCallbackOptions {
	_options.CallbackURL = core.StringPtr(callbackURL)
	return _options
}

// SetUserSecret : Allow user to set UserSecret
func (_options *RegisterCallbackOptions) SetUserSecret(userSecret string) *RegisterCallbackOptions {
	_options.UserSecret = core.StringPtr(userSecret)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RegisterCallbackOptions) SetHeaders(param map[string]string) *RegisterCallbackOptions {
	options.Headers = param
	return options
}

// RegisterStatus : Information about a request to register a callback for asynchronous speech recognition.
type RegisterStatus struct {
	// The current status of the job:
	// * `created`: The service successfully allowlisted the callback URL as a result of the call.
	// * `already created`: The URL was already allowlisted.
	Status *string `json:"status" validate:"required"`

	// The callback URL that is successfully registered.
	URL *string `json:"url" validate:"required"`
}

// Constants associated with the RegisterStatus.Status property.
// The current status of the job:
// * `created`: The service successfully allowlisted the callback URL as a result of the call.
// * `already created`: The URL was already allowlisted.
const (
	RegisterStatusStatusAlreadyCreatedConst = "already created"
	RegisterStatusStatusCreatedConst        = "created"
)

// UnmarshalRegisterStatus unmarshals an instance of RegisterStatus from the specified map of raw messages.
func UnmarshalRegisterStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RegisterStatus)
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResetAcousticModelOptions : The ResetAcousticModel options.
type ResetAcousticModelOptions struct {
	// The customization ID (GUID) of the custom acoustic model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewResetAcousticModelOptions : Instantiate ResetAcousticModelOptions
func (*SpeechToTextV1) NewResetAcousticModelOptions(customizationID string) *ResetAcousticModelOptions {
	return &ResetAcousticModelOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *ResetAcousticModelOptions) SetCustomizationID(customizationID string) *ResetAcousticModelOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ResetAcousticModelOptions) SetHeaders(param map[string]string) *ResetAcousticModelOptions {
	options.Headers = param
	return options
}

// ResetLanguageModelOptions : The ResetLanguageModel options.
type ResetLanguageModelOptions struct {
	// The customization ID (GUID) of the custom language model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewResetLanguageModelOptions : Instantiate ResetLanguageModelOptions
func (*SpeechToTextV1) NewResetLanguageModelOptions(customizationID string) *ResetLanguageModelOptions {
	return &ResetLanguageModelOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *ResetLanguageModelOptions) SetCustomizationID(customizationID string) *ResetLanguageModelOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ResetLanguageModelOptions) SetHeaders(param map[string]string) *ResetLanguageModelOptions {
	options.Headers = param
	return options
}

// SpeakerLabelsResult : Information about the speakers from speech recognition results.
type SpeakerLabelsResult struct {
	// The start time of a word from the transcript. The value matches the start time of a word from the `timestamps`
	// array.
	From *float32 `json:"from" validate:"required"`

	// The end time of a word from the transcript. The value matches the end time of a word from the `timestamps` array.
	To *float32 `json:"to" validate:"required"`

	// The numeric identifier that the service assigns to a speaker from the audio. Speaker IDs begin at `0` initially but
	// can evolve and change across interim results (if supported by the method) and between interim and final results as
	// the service processes the audio. They are not guaranteed to be sequential, contiguous, or ordered.
	Speaker *int64 `json:"speaker" validate:"required"`

	// A score that indicates the service's confidence in its identification of the speaker in the range of 0.0 to 1.0.
	Confidence *float32 `json:"confidence" validate:"required"`

	// An indication of whether the service might further change word and speaker-label results. A value of `true` means
	// that the service guarantees not to send any further updates for the current or any preceding results; `false` means
	// that the service might send further updates to the results.
	Final *bool `json:"final" validate:"required"`
}

// UnmarshalSpeakerLabelsResult unmarshals an instance of SpeakerLabelsResult from the specified map of raw messages.
func UnmarshalSpeakerLabelsResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SpeakerLabelsResult)
	err = core.UnmarshalPrimitive(m, "from", &obj.From)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "to", &obj.To)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "speaker", &obj.Speaker)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "confidence", &obj.Confidence)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "final", &obj.Final)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SpeechModel : Information about an available language model.
type SpeechModel struct {
	// The name of the model for use as an identifier in calls to the service (for example, `en-US_BroadbandModel`).
	Name *string `json:"name" validate:"required"`

	// The language identifier of the model (for example, `en-US`).
	Language *string `json:"language" validate:"required"`

	// The sampling rate (minimum acceptable rate for audio) used by the model in Hertz.
	Rate *int64 `json:"rate" validate:"required"`

	// The URI for the model.
	URL *string `json:"url" validate:"required"`

	// Indicates whether select service features are supported with the model.
	SupportedFeatures *SupportedFeatures `json:"supported_features" validate:"required"`

	// A brief description of the model.
	Description *string `json:"description" validate:"required"`
}

// UnmarshalSpeechModel unmarshals an instance of SpeechModel from the specified map of raw messages.
func UnmarshalSpeechModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SpeechModel)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "language", &obj.Language)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rate", &obj.Rate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "supported_features", &obj.SupportedFeatures, UnmarshalSupportedFeatures)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SpeechModels : Information about the available language models.
type SpeechModels struct {
	// An array of `SpeechModel` objects that provides information about each available model.
	Models []SpeechModel `json:"models" validate:"required"`
}

// UnmarshalSpeechModels unmarshals an instance of SpeechModels from the specified map of raw messages.
func UnmarshalSpeechModels(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SpeechModels)
	err = core.UnmarshalModel(m, "models", &obj.Models, UnmarshalSpeechModel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SpeechRecognitionAlternative : An alternative transcript from speech recognition results.
type SpeechRecognitionAlternative struct {
	// A transcription of the audio.
	Transcript *string `json:"transcript" validate:"required"`

	// A score that indicates the service's confidence in the transcript in the range of 0.0 to 1.0. The service returns a
	// confidence score only for the best alternative and only with results marked as final.
	Confidence *float64 `json:"confidence,omitempty"`

	// Time alignments for each word from the transcript as a list of lists. Each inner list consists of three elements:
	// the word followed by its start and end time in seconds, for example: `[["hello",0.0,1.2],["world",1.2,2.5]]`.
	// Timestamps are returned only for the best alternative.
	Timestamps interface{} `json:"timestamps,omitempty"`

	// A confidence score for each word of the transcript as a list of lists. Each inner list consists of two elements: the
	// word and its confidence score in the range of 0.0 to 1.0, for example: `[["hello",0.95],["world",0.86]]`. Confidence
	// scores are returned only for the best alternative and only with results marked as final.
	WordConfidence interface{} `json:"word_confidence,omitempty"`
}

// UnmarshalSpeechRecognitionAlternative unmarshals an instance of SpeechRecognitionAlternative from the specified map of raw messages.
func UnmarshalSpeechRecognitionAlternative(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SpeechRecognitionAlternative)
	err = core.UnmarshalPrimitive(m, "transcript", &obj.Transcript)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "confidence", &obj.Confidence)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "timestamps", &obj.Timestamps)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "word_confidence", &obj.WordConfidence)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SpeechRecognitionResult : Component results for a speech recognition request.
type SpeechRecognitionResult struct {
	// An indication of whether the transcription results are final:
	// * If `true`, the results for this utterance are final. They are guaranteed not to be updated further.
	// * If `false`, the results are interim. They can be updated with further interim results until final results are
	// eventually sent.
	//
	// **Note:** Because `final` is a reserved word in Java and Swift, the field is renamed `xFinal` in Java and is escaped
	// with back quotes in Swift.
	Final *bool `json:"final" validate:"required"`

	// An array of alternative transcripts. The `alternatives` array can include additional requested output such as word
	// confidence or timestamps.
	Alternatives []SpeechRecognitionAlternative `json:"alternatives" validate:"required"`

	// A dictionary (or associative array) whose keys are the strings specified for `keywords` if both that parameter and
	// `keywords_threshold` are specified. The value for each key is an array of matches spotted in the audio for that
	// keyword. Each match is described by a `KeywordResult` object. A keyword for which no matches are found is omitted
	// from the dictionary. The dictionary is omitted entirely if no matches are found for any keywords.
	KeywordsResult map[string][]KeywordResult `json:"keywords_result,omitempty"`

	// An array of alternative hypotheses found for words of the input audio if a `word_alternatives_threshold` is
	// specified.
	WordAlternatives []WordAlternativeResults `json:"word_alternatives,omitempty"`

	// If the `split_transcript_at_phrase_end` parameter is `true`, describes the reason for the split:
	// * `end_of_data` - The end of the input audio stream.
	// * `full_stop` - A full semantic stop, such as for the conclusion of a grammatical sentence. The insertion of splits
	// is influenced by the base language model and biased by custom language models and grammars.
	// * `reset` - The amount of audio that is currently being processed exceeds the two-minute maximum. The service splits
	// the transcript to avoid excessive memory use.
	// * `silence` - A pause or silence that is at least as long as the pause interval.
	EndOfUtterance *string `json:"end_of_utterance,omitempty"`
}

// Constants associated with the SpeechRecognitionResult.EndOfUtterance property.
// If the `split_transcript_at_phrase_end` parameter is `true`, describes the reason for the split:
// * `end_of_data` - The end of the input audio stream.
// * `full_stop` - A full semantic stop, such as for the conclusion of a grammatical sentence. The insertion of splits
// is influenced by the base language model and biased by custom language models and grammars.
// * `reset` - The amount of audio that is currently being processed exceeds the two-minute maximum. The service splits
// the transcript to avoid excessive memory use.
// * `silence` - A pause or silence that is at least as long as the pause interval.
const (
	SpeechRecognitionResultEndOfUtteranceEndOfDataConst = "end_of_data"
	SpeechRecognitionResultEndOfUtteranceFullStopConst  = "full_stop"
	SpeechRecognitionResultEndOfUtteranceResetConst     = "reset"
	SpeechRecognitionResultEndOfUtteranceSilenceConst   = "silence"
)

// UnmarshalSpeechRecognitionResult unmarshals an instance of SpeechRecognitionResult from the specified map of raw messages.
func UnmarshalSpeechRecognitionResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SpeechRecognitionResult)
	err = core.UnmarshalPrimitive(m, "final", &obj.Final)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "alternatives", &obj.Alternatives, UnmarshalSpeechRecognitionAlternative)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "keywords_result", &obj.KeywordsResult, UnmarshalKeywordResult)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "word_alternatives", &obj.WordAlternatives, UnmarshalWordAlternativeResults)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "end_of_utterance", &obj.EndOfUtterance)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SpeechRecognitionResults : The complete results for a speech recognition request.
type SpeechRecognitionResults struct {
	// An array of `SpeechRecognitionResult` objects that can include interim and final results (interim results are
	// returned only if supported by the method). Final results are guaranteed not to change; interim results might be
	// replaced by further interim results and eventually final results.
	//
	// For the HTTP interfaces, all results arrive at the same time. For the WebSocket interface, results can be sent as
	// multiple separate responses. The service periodically sends updates to the results list. The `result_index` is
	// incremented to the lowest index in the array that has changed for new results.
	//
	// For more information, see [Understanding speech recognition
	// results](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-basic-response).
	Results []SpeechRecognitionResult `json:"results,omitempty"`

	// An index that indicates a change point in the `results` array. The service increments the index for additional
	// results that it sends for new audio for the same request. All results with the same index are delivered at the same
	// time. The same index can include multiple final results that are delivered with the same response.
	ResultIndex *int64 `json:"result_index,omitempty"`

	// An array of `SpeakerLabelsResult` objects that identifies which words were spoken by which speakers in a
	// multi-person exchange. The array is returned only if the `speaker_labels` parameter is `true`. When interim results
	// are also requested for methods that support them, it is possible for a `SpeechRecognitionResults` object to include
	// only the `speaker_labels` field.
	SpeakerLabels []SpeakerLabelsResult `json:"speaker_labels,omitempty"`

	// If processing metrics are requested, information about the service's processing of the input audio. Processing
	// metrics are not available with the synchronous [Recognize audio](#recognize) method.
	ProcessingMetrics *ProcessingMetrics `json:"processing_metrics,omitempty"`

	// If audio metrics are requested, information about the signal characteristics of the input audio.
	AudioMetrics *AudioMetrics `json:"audio_metrics,omitempty"`

	// An array of warning messages associated with the request:
	// * Warnings for invalid parameters or fields can include a descriptive message and a list of invalid argument
	// strings, for example, `"Unknown arguments:"` or `"Unknown url query arguments:"` followed by a list of the form
	// `"{invalid_arg_1}, {invalid_arg_2}."`
	// * The following warning is returned if the request passes a custom model that is based on an older version of a base
	// model for which an updated version is available: `"Using previous version of base model, because your custom model
	// has been built with it. Please note that this version will be supported only for a limited time. Consider updating
	// your custom model to the new base model. If you do not do that you will be automatically switched to base model when
	// you used the non-updated custom model."`
	//
	// In both cases, the request succeeds despite the warnings.
	Warnings []string `json:"warnings,omitempty"`
}

// UnmarshalSpeechRecognitionResults unmarshals an instance of SpeechRecognitionResults from the specified map of raw messages.
func UnmarshalSpeechRecognitionResults(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SpeechRecognitionResults)
	err = core.UnmarshalModel(m, "results", &obj.Results, UnmarshalSpeechRecognitionResult)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "result_index", &obj.ResultIndex)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "speaker_labels", &obj.SpeakerLabels, UnmarshalSpeakerLabelsResult)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "processing_metrics", &obj.ProcessingMetrics, UnmarshalProcessingMetrics)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "audio_metrics", &obj.AudioMetrics, UnmarshalAudioMetrics)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "warnings", &obj.Warnings)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SupportedFeatures : Indicates whether select service features are supported with the model.
type SupportedFeatures struct {
	// Indicates whether the customization interface can be used to create a custom language model based on the language
	// model.
	CustomLanguageModel *bool `json:"custom_language_model" validate:"required"`

	// Indicates whether the customization interface can be used to create a custom acoustic model based on the language
	// model.
	CustomAcousticModel *bool `json:"custom_acoustic_model" validate:"required"`

	// Indicates whether the `speaker_labels` parameter can be used with the language model.
	//
	// **Note:** The field returns `true` for all models. However, speaker labels are supported for use only with the
	// following languages and models:
	// * _For previous-generation models,_ the parameter can be used with Australian English, US English, German, Japanese,
	// Korean, and Spanish (both broadband and narrowband models) and UK English (narrowband model) transcription only.
	// * _For next-generation models,_ the parameter can be used with Czech, English (Australian, Indian, UK, and US),
	// German, Japanese, Korean, and Spanish transcription only.
	//
	// Speaker labels are not supported for use with any other languages or models.
	SpeakerLabels *bool `json:"speaker_labels" validate:"required"`

	// Indicates whether the `low_latency` parameter can be used with a next-generation language model. The field is
	// returned only for next-generation models. Previous-generation models do not support the `low_latency` parameter.
	LowLatency *bool `json:"low_latency,omitempty"`
}

// UnmarshalSupportedFeatures unmarshals an instance of SupportedFeatures from the specified map of raw messages.
func UnmarshalSupportedFeatures(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SupportedFeatures)
	err = core.UnmarshalPrimitive(m, "custom_language_model", &obj.CustomLanguageModel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "custom_acoustic_model", &obj.CustomAcousticModel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "speaker_labels", &obj.SpeakerLabels)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "low_latency", &obj.LowLatency)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TrainAcousticModelOptions : The TrainAcousticModel options.
type TrainAcousticModelOptions struct {
	// The customization ID (GUID) of the custom acoustic model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// The customization ID (GUID) of a custom language model that is to be used during training of the custom acoustic
	// model. Specify a custom language model that has been trained with verbatim transcriptions of the audio resources or
	// that contains words that are relevant to the contents of the audio resources. The custom language model must be
	// based on the same version of the same base model as the custom acoustic model, and the custom language model must be
	// fully trained and available. The credentials specified with the request must own both custom models.
	CustomLanguageModelID *string `json:"custom_language_model_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewTrainAcousticModelOptions : Instantiate TrainAcousticModelOptions
func (*SpeechToTextV1) NewTrainAcousticModelOptions(customizationID string) *TrainAcousticModelOptions {
	return &TrainAcousticModelOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *TrainAcousticModelOptions) SetCustomizationID(customizationID string) *TrainAcousticModelOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetCustomLanguageModelID : Allow user to set CustomLanguageModelID
func (_options *TrainAcousticModelOptions) SetCustomLanguageModelID(customLanguageModelID string) *TrainAcousticModelOptions {
	_options.CustomLanguageModelID = core.StringPtr(customLanguageModelID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *TrainAcousticModelOptions) SetHeaders(param map[string]string) *TrainAcousticModelOptions {
	options.Headers = param
	return options
}

// TrainLanguageModelOptions : The TrainLanguageModel options.
type TrainLanguageModelOptions struct {
	// The customization ID (GUID) of the custom language model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// _For custom models that are based on previous-generation models_, the type of words from the custom language model's
	// words resource on which to train the model:
	// * `all` (the default) trains the model on all new words, regardless of whether they were extracted from corpora or
	// grammars or were added or modified by the user.
	// * `user` trains the model only on custom words that were added or modified by the user directly. The model is not
	// trained on new words extracted from corpora or grammars.
	//
	// _For custom models that are based on next-generation models_, the service ignores the parameter. The words resource
	// contains only custom words that the user adds or modifies directly, so the parameter is unnecessary.
	WordTypeToAdd *string `json:"word_type_to_add,omitempty"`

	// Specifies a customization weight for the custom language model. The customization weight tells the service how much
	// weight to give to words from the custom language model compared to those from the base model for speech recognition.
	// Specify a value between 0.0 and 1.0; the default is 0.3.
	//
	// The default value yields the best performance in general. Assign a higher value if your audio makes frequent use of
	// OOV words from the custom model. Use caution when setting the weight: a higher value can improve the accuracy of
	// phrases from the custom model's domain, but it can negatively affect performance on non-domain phrases.
	//
	// The value that you assign is used for all recognition requests that use the model. You can override it for any
	// recognition request by specifying a customization weight for that request.
	//
	// See [Using customization weight](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-languageUse#weight).
	CustomizationWeight *float64 `json:"customization_weight,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the TrainLanguageModelOptions.WordTypeToAdd property.
// _For custom models that are based on previous-generation models_, the type of words from the custom language model's
// words resource on which to train the model:
// * `all` (the default) trains the model on all new words, regardless of whether they were extracted from corpora or
// grammars or were added or modified by the user.
// * `user` trains the model only on custom words that were added or modified by the user directly. The model is not
// trained on new words extracted from corpora or grammars.
//
// _For custom models that are based on next-generation models_, the service ignores the parameter. The words resource
// contains only custom words that the user adds or modifies directly, so the parameter is unnecessary.
const (
	TrainLanguageModelOptionsWordTypeToAddAllConst  = "all"
	TrainLanguageModelOptionsWordTypeToAddUserConst = "user"
)

// NewTrainLanguageModelOptions : Instantiate TrainLanguageModelOptions
func (*SpeechToTextV1) NewTrainLanguageModelOptions(customizationID string) *TrainLanguageModelOptions {
	return &TrainLanguageModelOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *TrainLanguageModelOptions) SetCustomizationID(customizationID string) *TrainLanguageModelOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetWordTypeToAdd : Allow user to set WordTypeToAdd
func (_options *TrainLanguageModelOptions) SetWordTypeToAdd(wordTypeToAdd string) *TrainLanguageModelOptions {
	_options.WordTypeToAdd = core.StringPtr(wordTypeToAdd)
	return _options
}

// SetCustomizationWeight : Allow user to set CustomizationWeight
func (_options *TrainLanguageModelOptions) SetCustomizationWeight(customizationWeight float64) *TrainLanguageModelOptions {
	_options.CustomizationWeight = core.Float64Ptr(customizationWeight)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *TrainLanguageModelOptions) SetHeaders(param map[string]string) *TrainLanguageModelOptions {
	options.Headers = param
	return options
}

// TrainingResponse : The response from training of a custom language or custom acoustic model.
type TrainingResponse struct {
	// An array of `TrainingWarning` objects that lists any invalid resources contained in the custom model. For custom
	// language models, invalid resources are grouped and identified by type of resource. The method can return warnings
	// only if the `strict` parameter is set to `false`.
	Warnings []TrainingWarning `json:"warnings,omitempty"`
}

// UnmarshalTrainingResponse unmarshals an instance of TrainingResponse from the specified map of raw messages.
func UnmarshalTrainingResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TrainingResponse)
	err = core.UnmarshalModel(m, "warnings", &obj.Warnings, UnmarshalTrainingWarning)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TrainingWarning : A warning from training of a custom language or custom acoustic model.
type TrainingWarning struct {
	// An identifier for the type of invalid resources listed in the `description` field.
	Code *string `json:"code" validate:"required"`

	// A warning message that lists the invalid resources that are excluded from the custom model's training. The message
	// has the following format: `Analysis of the following {resource_type} has not completed successfully:
	// [{resource_names}]. They will be excluded from custom {model_type} model training.`.
	Message *string `json:"message" validate:"required"`
}

// Constants associated with the TrainingWarning.Code property.
// An identifier for the type of invalid resources listed in the `description` field.
const (
	TrainingWarningCodeInvalidAudioFilesConst   = "invalid_audio_files"
	TrainingWarningCodeInvalidCorpusFilesConst  = "invalid_corpus_files"
	TrainingWarningCodeInvalidGrammarFilesConst = "invalid_grammar_files"
	TrainingWarningCodeInvalidWordsConst        = "invalid_words"
)

// UnmarshalTrainingWarning unmarshals an instance of TrainingWarning from the specified map of raw messages.
func UnmarshalTrainingWarning(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TrainingWarning)
	err = core.UnmarshalPrimitive(m, "code", &obj.Code)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UnregisterCallbackOptions : The UnregisterCallback options.
type UnregisterCallbackOptions struct {
	// The callback URL that is to be unregistered.
	CallbackURL *string `json:"callback_url" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUnregisterCallbackOptions : Instantiate UnregisterCallbackOptions
func (*SpeechToTextV1) NewUnregisterCallbackOptions(callbackURL string) *UnregisterCallbackOptions {
	return &UnregisterCallbackOptions{
		CallbackURL: core.StringPtr(callbackURL),
	}
}

// SetCallbackURL : Allow user to set CallbackURL
func (_options *UnregisterCallbackOptions) SetCallbackURL(callbackURL string) *UnregisterCallbackOptions {
	_options.CallbackURL = core.StringPtr(callbackURL)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UnregisterCallbackOptions) SetHeaders(param map[string]string) *UnregisterCallbackOptions {
	options.Headers = param
	return options
}

// UpgradeAcousticModelOptions : The UpgradeAcousticModel options.
type UpgradeAcousticModelOptions struct {
	// The customization ID (GUID) of the custom acoustic model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// If the custom acoustic model was trained with a custom language model, the customization ID (GUID) of that custom
	// language model. The custom language model must be upgraded before the custom acoustic model can be upgraded. The
	// custom language model must be fully trained and available. The credentials specified with the request must own both
	// custom models.
	CustomLanguageModelID *string `json:"custom_language_model_id,omitempty"`

	// If `true`, forces the upgrade of a custom acoustic model for which no input data has been modified since it was last
	// trained. Use this parameter only to force the upgrade of a custom acoustic model that is trained with a custom
	// language model, and only if you receive a 400 response code and the message `No input data modified since last
	// training`. See [Upgrading a custom acoustic
	// model](https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-custom-upgrade#custom-upgrade-acoustic).
	Force *bool `json:"force,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpgradeAcousticModelOptions : Instantiate UpgradeAcousticModelOptions
func (*SpeechToTextV1) NewUpgradeAcousticModelOptions(customizationID string) *UpgradeAcousticModelOptions {
	return &UpgradeAcousticModelOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *UpgradeAcousticModelOptions) SetCustomizationID(customizationID string) *UpgradeAcousticModelOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetCustomLanguageModelID : Allow user to set CustomLanguageModelID
func (_options *UpgradeAcousticModelOptions) SetCustomLanguageModelID(customLanguageModelID string) *UpgradeAcousticModelOptions {
	_options.CustomLanguageModelID = core.StringPtr(customLanguageModelID)
	return _options
}

// SetForce : Allow user to set Force
func (_options *UpgradeAcousticModelOptions) SetForce(force bool) *UpgradeAcousticModelOptions {
	_options.Force = core.BoolPtr(force)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpgradeAcousticModelOptions) SetHeaders(param map[string]string) *UpgradeAcousticModelOptions {
	options.Headers = param
	return options
}

// UpgradeLanguageModelOptions : The UpgradeLanguageModel options.
type UpgradeLanguageModelOptions struct {
	// The customization ID (GUID) of the custom language model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpgradeLanguageModelOptions : Instantiate UpgradeLanguageModelOptions
func (*SpeechToTextV1) NewUpgradeLanguageModelOptions(customizationID string) *UpgradeLanguageModelOptions {
	return &UpgradeLanguageModelOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *UpgradeLanguageModelOptions) SetCustomizationID(customizationID string) *UpgradeLanguageModelOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpgradeLanguageModelOptions) SetHeaders(param map[string]string) *UpgradeLanguageModelOptions {
	options.Headers = param
	return options
}

// Word : Information about a word from a custom language model.
type Word struct {
	// A word from the custom model's words resource. The spelling of the word is used to train the model.
	Word *string `json:"word" validate:"required"`

	// _For a custom model that is based on a previous-generation model_, an array of as many as five pronunciations for
	// the word. The array can include the sounds-like pronunciation that is automatically generated by the service if none
	// is provided when the word is added to the custom model; the service adds this pronunciation when it finishes
	// processing the word.
	//
	// _For a custom model that is based on a next-generation model_, this field does not apply. Custom models based on
	// next-generation models do not support the `sounds_like` field, which is ignored.
	SoundsLike []string `json:"sounds_like" validate:"required"`

	// The spelling of the word that the service uses to display the word in a transcript. The field contains an empty
	// string if no display-as value is provided for the word, in which case the word is displayed as it is spelled.
	DisplayAs *string `json:"display_as" validate:"required"`

	// _For a custom model that is based on a previous-generation model_, a sum of the number of times the word is found
	// across all corpora and grammars. For example, if the word occurs five times in one corpus and seven times in
	// another, its count is `12`. If you add a custom word to a model before it is added by any corpora or grammars, the
	// count begins at `1`; if the word is added from a corpus or grammar first and later modified, the count reflects only
	// the number of times it is found in corpora and grammars.
	//
	// _For a custom model that is based on a next-generation model_, the `count` field for any word is always `1`.
	Count *int64 `json:"count" validate:"required"`

	// An array of sources that describes how the word was added to the custom model's words resource.
	// * _For a custom model that is based on previous-generation model,_ the field includes the name of each corpus and
	// grammar from which the service extracted the word. For OOV that are added by multiple corpora or grammars, the names
	// of all corpora and grammars are listed. If you modified or added the word directly, the field includes the string
	// `user`.
	// * _For a custom model that is based on a next-generation model,_ this field shows only `user` for custom words that
	// were added directly to the custom model. Words from corpora and grammars are not added to the words resource for
	// custom models that are based on next-generation models.
	Source []string `json:"source" validate:"required"`

	// If the service discovered one or more problems that you need to correct for the word's definition, an array that
	// describes each of the errors.
	Error []WordError `json:"error,omitempty"`
}

// UnmarshalWord unmarshals an instance of Word from the specified map of raw messages.
func UnmarshalWord(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Word)
	err = core.UnmarshalPrimitive(m, "word", &obj.Word)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sounds_like", &obj.SoundsLike)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "display_as", &obj.DisplayAs)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "count", &obj.Count)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source", &obj.Source)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "error", &obj.Error, UnmarshalWordError)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// WordAlternativeResult : An alternative hypothesis for a word from speech recognition results.
type WordAlternativeResult struct {
	// A confidence score for the word alternative hypothesis in the range of 0.0 to 1.0.
	Confidence *float64 `json:"confidence" validate:"required"`

	// An alternative hypothesis for a word from the input audio.
	Word *string `json:"word" validate:"required"`
}

// UnmarshalWordAlternativeResult unmarshals an instance of WordAlternativeResult from the specified map of raw messages.
func UnmarshalWordAlternativeResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(WordAlternativeResult)
	err = core.UnmarshalPrimitive(m, "confidence", &obj.Confidence)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "word", &obj.Word)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// WordAlternativeResults : Information about alternative hypotheses for words from speech recognition results.
type WordAlternativeResults struct {
	// The start time in seconds of the word from the input audio that corresponds to the word alternatives.
	StartTime *float64 `json:"start_time" validate:"required"`

	// The end time in seconds of the word from the input audio that corresponds to the word alternatives.
	EndTime *float64 `json:"end_time" validate:"required"`

	// An array of alternative hypotheses for a word from the input audio.
	Alternatives []WordAlternativeResult `json:"alternatives" validate:"required"`
}

// UnmarshalWordAlternativeResults unmarshals an instance of WordAlternativeResults from the specified map of raw messages.
func UnmarshalWordAlternativeResults(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(WordAlternativeResults)
	err = core.UnmarshalPrimitive(m, "start_time", &obj.StartTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "end_time", &obj.EndTime)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "alternatives", &obj.Alternatives, UnmarshalWordAlternativeResult)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// WordError : An error associated with a word from a custom language model.
type WordError struct {
	// A key-value pair that describes an error associated with the definition of a word in the words resource. The pair
	// has the format `"element": "message"`, where `element` is the aspect of the definition that caused the problem and
	// `message` describes the problem. The following example describes a problem with one of the word's sounds-like
	// definitions: `"{sounds_like_string}": "Numbers are not allowed in sounds-like. You can try for example
	// '{suggested_string}'."`.
	Element *string `json:"element" validate:"required"`
}

// UnmarshalWordError unmarshals an instance of WordError from the specified map of raw messages.
func UnmarshalWordError(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(WordError)
	err = core.UnmarshalPrimitive(m, "element", &obj.Element)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Words : Information about the words from a custom language model.
type Words struct {
	// An array of `Word` objects that provides information about each word in the custom model's words resource. The array
	// is empty if the custom model has no words.
	Words []Word `json:"words" validate:"required"`
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
