// Package texttospeechv1 : Operations and models for the TextToSpeechV1 service
package texttospeechv1

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
	"io"
)

// TextToSpeechV1 : ### Service Overview
// The IBM&reg; Text to Speech service provides APIs that use IBM's speech-synthesis capabilities to synthesize text
// into natural-sounding speech in a variety of languages, dialects, and voices. The service supports at least one male
// or female voice, sometimes both, for each language. The audio is streamed back to the client with minimal delay.
//
// For speech synthesis, the service supports a synchronous HTTP Representational State Transfer (REST) interface. It
// also supports a WebSocket interface that provides both plain text and SSML input, including the SSML &lt;mark&gt;
// element and word timings. SSML is an XML-based markup language that provides text annotation for speech-synthesis
// applications.
//
// The service also offers a customization interface. You can use the interface to define sounds-like or phonetic
// translations for words. A sounds-like translation consists of one or more words that, when combined, sound like the
// word. A phonetic translation is based on the SSML phoneme format for representing a word. You can specify a phonetic
// translation in standard International Phonetic Alphabet (IPA) representation or in the proprietary IBM Symbolic
// Phonetic Representation (SPR).
//
// Version: V1
// See: http://www.ibm.com/watson/developercloud/text-to-speech.html
type TextToSpeechV1 struct {
	Service *core.BaseService
}

// TextToSpeechV1Options : Service options
type TextToSpeechV1Options struct {
	URL            string
	Username       string
	Password       string
	IAMApiKey      string
	IAMAccessToken string
	IAMURL         string
}

// NewTextToSpeechV1 : Instantiate TextToSpeechV1
func NewTextToSpeechV1(options *TextToSpeechV1Options) (*TextToSpeechV1, error) {
	if options.URL == "" {
		options.URL = "https://stream.watsonplatform.net/text-to-speech/api"
	}

	serviceOptions := &core.ServiceOptions{
		URL:            options.URL,
		Username:       options.Username,
		Password:       options.Password,
		IAMApiKey:      options.IAMApiKey,
		IAMAccessToken: options.IAMAccessToken,
		IAMURL:         options.IAMURL,
	}
	service, serviceErr := core.NewBaseService(serviceOptions, "text_to_speech", "Text to Speech")
	if serviceErr != nil {
		return nil, serviceErr
	}

	return &TextToSpeechV1{Service: service}, nil
}

// GetVoice : Get a voice
// Gets information about the specified voice. The information includes the name, language, gender, and other details
// about the voice. Specify a customization ID to obtain information for that custom voice model of the specified voice.
// To list information about all available voices, use the **List voices** method.
//
// **See also:** [Listing a specific voice](https://cloud.ibm.com/docs/services/text-to-speech/voices.html#listVoice).
func (textToSpeech *TextToSpeechV1) GetVoice(getVoiceOptions *GetVoiceOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(getVoiceOptions, "getVoiceOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(getVoiceOptions, "getVoiceOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/voices"}
	pathParameters := []string{*getVoiceOptions.Voice}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(textToSpeech.Service.Options.URL, pathSegments, pathParameters)

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
		return nil, err
	}

	response, err := textToSpeech.Service.Request(request, new(Voice))
	return response, err
}

// GetGetVoiceResult : Retrieve result of GetVoice operation
func (textToSpeech *TextToSpeechV1) GetGetVoiceResult(response *core.DetailedResponse) *Voice {
	result, ok := response.Result.(*Voice)
	if ok {
		return result
	}
	return nil
}

// ListVoices : List voices
// Lists all voices available for use with the service. The information includes the name, language, gender, and other
// details about the voice. To see information about a specific voice, use the **Get a voice** method.
//
// **See also:** [Listing all available
// voices](https://cloud.ibm.com/docs/services/text-to-speech/voices.html#listVoices).
func (textToSpeech *TextToSpeechV1) ListVoices(listVoicesOptions *ListVoicesOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateStruct(listVoicesOptions, "listVoicesOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/voices"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(textToSpeech.Service.Options.URL, pathSegments, pathParameters)

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
		return nil, err
	}

	response, err := textToSpeech.Service.Request(request, new(Voices))
	return response, err
}

// GetListVoicesResult : Retrieve result of ListVoices operation
func (textToSpeech *TextToSpeechV1) GetListVoicesResult(response *core.DetailedResponse) *Voices {
	result, ok := response.Result.(*Voices)
	if ok {
		return result
	}
	return nil
}

// Synthesize : Synthesize audio
// Synthesizes text to audio that is spoken in the specified voice. The service bases its understanding of the language
// for the input text on the specified voice. Use a voice that matches the language of the input text.
//
// The method accepts a maximum of 5 KB of input text in the body of the request, and 8 KB for the URL and headers. The
// 5 KB limit includes any SSML tags that you specify. The service returns the synthesized audio stream as an array of
// bytes.
//
// **See also:** [The HTTP interface](https://cloud.ibm.com/docs/services/text-to-speech/http.html).
//
// ### Audio formats (accept types)
//
//  The service can return audio in the following formats (MIME types).
// * Where indicated, you can optionally specify the sampling rate (`rate`) of the audio. You must specify a sampling
// rate for the `audio/l16` and `audio/mulaw` formats. A specified sampling rate must lie in the range of 8 kHz to 192
// kHz.
// * For the `audio/l16` format, you can optionally specify the endianness (`endianness`) of the audio:
// `endianness=big-endian` or `endianness=little-endian`.
//
// Use the `Accept` header or the `accept` parameter to specify the requested format of the response audio. If you omit
// an audio format altogether, the service returns the audio in Ogg format with the Opus codec
// (`audio/ogg;codecs=opus`). The service always returns single-channel audio.
// * `audio/basic`
//
//   The service returns audio with a sampling rate of 8000 Hz.
// * `audio/flac`
//
//   You can optionally specify the `rate` of the audio. The default sampling rate is 22,050 Hz.
// * `audio/l16`
//
//   You must specify the `rate` of the audio. You can optionally specify the `endianness` of the audio. The default
// endianness is `little-endian`.
// * `audio/mp3`
//
//   You can optionally specify the `rate` of the audio. The default sampling rate is 22,050 Hz.
// * `audio/mpeg`
//
//   You can optionally specify the `rate` of the audio. The default sampling rate is 22,050 Hz.
// * `audio/mulaw`
//
//   You must specify the `rate` of the audio.
// * `audio/ogg`
//
//   The service returns the audio in the `vorbis` codec. You can optionally specify the `rate` of the audio. The
// default sampling rate is 22,050 Hz.
// * `audio/ogg;codecs=opus`
//
//   You can optionally specify the `rate` of the audio. The default sampling rate is 22,050 Hz.
// * `audio/ogg;codecs=vorbis`
//
//   You can optionally specify the `rate` of the audio. The default sampling rate is 22,050 Hz.
// * `audio/wav`
//
//   You can optionally specify the `rate` of the audio. The default sampling rate is 22,050 Hz.
// * `audio/webm`
//
//   The service returns the audio in the `opus` codec. The service returns audio with a sampling rate of 48,000 Hz.
// * `audio/webm;codecs=opus`
//
//   The service returns audio with a sampling rate of 48,000 Hz.
// * `audio/webm;codecs=vorbis`
//
//   You can optionally specify the `rate` of the audio. The default sampling rate is 22,050 Hz.
//
// For more information about specifying an audio format, including additional details about some of the formats, see
// [Audio formats](https://cloud.ibm.com/docs/services/text-to-speech/audio-formats.html).
//
// ### Warning messages
//
//  If a request includes invalid query parameters, the service returns a `Warnings` response header that provides
// messages about the invalid parameters. The warning includes a descriptive message and a list of invalid argument
// strings. For example, a message such as `\"Unknown arguments:\"` or `\"Unknown url query arguments:\"` followed by a
// list of the form `\"{invalid_arg_1}, {invalid_arg_2}.\"` The request succeeds despite the warnings.
func (textToSpeech *TextToSpeechV1) Synthesize(synthesizeOptions *SynthesizeOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(synthesizeOptions, "synthesizeOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(synthesizeOptions, "synthesizeOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/synthesize"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(textToSpeech.Service.Options.URL, pathSegments, pathParameters)

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
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := textToSpeech.Service.Request(request, new(io.ReadCloser))
	return response, err
}

// GetSynthesizeResult : Retrieve result of Synthesize operation
func (textToSpeech *TextToSpeechV1) GetSynthesizeResult(response *core.DetailedResponse) io.ReadCloser {
	result, ok := response.Result.(io.ReadCloser)
	if ok {
		return result
	}
	return nil
}

// GetPronunciation : Get pronunciation
// Gets the phonetic pronunciation for the specified word. You can request the pronunciation for a specific format. You
// can also request the pronunciation for a specific voice to see the default translation for the language of that voice
// or for a specific custom voice model to see the translation for that voice model.
//
// **Note:** This method is currently a beta release.
//
// **See also:** [Querying a word from a
// language](https://cloud.ibm.com/docs/services/text-to-speech/custom-entries.html#cuWordsQueryLanguage).
func (textToSpeech *TextToSpeechV1) GetPronunciation(getPronunciationOptions *GetPronunciationOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(getPronunciationOptions, "getPronunciationOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(getPronunciationOptions, "getPronunciationOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/pronunciation"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(textToSpeech.Service.Options.URL, pathSegments, pathParameters)

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
		return nil, err
	}

	response, err := textToSpeech.Service.Request(request, new(Pronunciation))
	return response, err
}

// GetGetPronunciationResult : Retrieve result of GetPronunciation operation
func (textToSpeech *TextToSpeechV1) GetGetPronunciationResult(response *core.DetailedResponse) *Pronunciation {
	result, ok := response.Result.(*Pronunciation)
	if ok {
		return result
	}
	return nil
}

// CreateVoiceModel : Create a custom model
// Creates a new empty custom voice model. You must specify a name for the new custom model. You can optionally specify
// the language and a description for the new model. The model is owned by the instance of the service whose credentials
// are used to create it.
//
// **Note:** This method is currently a beta release.
//
// **See also:** [Creating a custom
// model](https://cloud.ibm.com/docs/services/text-to-speech/custom-models.html#cuModelsCreate).
func (textToSpeech *TextToSpeechV1) CreateVoiceModel(createVoiceModelOptions *CreateVoiceModelOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(createVoiceModelOptions, "createVoiceModelOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(createVoiceModelOptions, "createVoiceModelOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/customizations"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(textToSpeech.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range createVoiceModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "CreateVoiceModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createVoiceModelOptions.Name != nil {
		body["name"] = createVoiceModelOptions.Name
	}
	if createVoiceModelOptions.Language != nil {
		body["language"] = createVoiceModelOptions.Language
	}
	if createVoiceModelOptions.Description != nil {
		body["description"] = createVoiceModelOptions.Description
	}
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := textToSpeech.Service.Request(request, new(VoiceModel))
	return response, err
}

// GetCreateVoiceModelResult : Retrieve result of CreateVoiceModel operation
func (textToSpeech *TextToSpeechV1) GetCreateVoiceModelResult(response *core.DetailedResponse) *VoiceModel {
	result, ok := response.Result.(*VoiceModel)
	if ok {
		return result
	}
	return nil
}

// DeleteVoiceModel : Delete a custom model
// Deletes the specified custom voice model. You must use credentials for the instance of the service that owns a model
// to delete it.
//
// **Note:** This method is currently a beta release.
//
// **See also:** [Deleting a custom
// model](https://cloud.ibm.com/docs/services/text-to-speech/custom-models.html#cuModelsDelete).
func (textToSpeech *TextToSpeechV1) DeleteVoiceModel(deleteVoiceModelOptions *DeleteVoiceModelOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(deleteVoiceModelOptions, "deleteVoiceModelOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(deleteVoiceModelOptions, "deleteVoiceModelOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/customizations"}
	pathParameters := []string{*deleteVoiceModelOptions.CustomizationID}

	builder := core.NewRequestBuilder(core.DELETE)
	builder.ConstructHTTPURL(textToSpeech.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteVoiceModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "DeleteVoiceModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "")

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := textToSpeech.Service.Request(request, nil)
	return response, err
}

// GetVoiceModel : Get a custom model
// Gets all information about a specified custom voice model. In addition to metadata such as the name and description
// of the voice model, the output includes the words and their translations as defined in the model. To see just the
// metadata for a voice model, use the **List custom models** method.
//
// **Note:** This method is currently a beta release.
//
// **See also:** [Querying a custom
// model](https://cloud.ibm.com/docs/services/text-to-speech/custom-models.html#cuModelsQuery).
func (textToSpeech *TextToSpeechV1) GetVoiceModel(getVoiceModelOptions *GetVoiceModelOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(getVoiceModelOptions, "getVoiceModelOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(getVoiceModelOptions, "getVoiceModelOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/customizations"}
	pathParameters := []string{*getVoiceModelOptions.CustomizationID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(textToSpeech.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range getVoiceModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "GetVoiceModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := textToSpeech.Service.Request(request, new(VoiceModel))
	return response, err
}

// GetGetVoiceModelResult : Retrieve result of GetVoiceModel operation
func (textToSpeech *TextToSpeechV1) GetGetVoiceModelResult(response *core.DetailedResponse) *VoiceModel {
	result, ok := response.Result.(*VoiceModel)
	if ok {
		return result
	}
	return nil
}

// ListVoiceModels : List custom models
// Lists metadata such as the name and description for all custom voice models that are owned by an instance of the
// service. Specify a language to list the voice models for that language only. To see the words in addition to the
// metadata for a specific voice model, use the **List a custom model** method. You must use credentials for the
// instance of the service that owns a model to list information about it.
//
// **Note:** This method is currently a beta release.
//
// **See also:** [Querying all custom
// models](https://cloud.ibm.com/docs/services/text-to-speech/custom-models.html#cuModelsQueryAll).
func (textToSpeech *TextToSpeechV1) ListVoiceModels(listVoiceModelsOptions *ListVoiceModelsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateStruct(listVoiceModelsOptions, "listVoiceModelsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/customizations"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(textToSpeech.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range listVoiceModelsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "ListVoiceModels")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")

	if listVoiceModelsOptions.Language != nil {
		builder.AddQuery("language", fmt.Sprint(*listVoiceModelsOptions.Language))
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := textToSpeech.Service.Request(request, new(VoiceModels))
	return response, err
}

// GetListVoiceModelsResult : Retrieve result of ListVoiceModels operation
func (textToSpeech *TextToSpeechV1) GetListVoiceModelsResult(response *core.DetailedResponse) *VoiceModels {
	result, ok := response.Result.(*VoiceModels)
	if ok {
		return result
	}
	return nil
}

// UpdateVoiceModel : Update a custom model
// Updates information for the specified custom voice model. You can update metadata such as the name and description of
// the voice model. You can also update the words in the model and their translations. Adding a new translation for a
// word that already exists in a custom model overwrites the word's existing translation. A custom model can contain no
// more than 20,000 entries. You must use credentials for the instance of the service that owns a model to update it.
//
// You can define sounds-like or phonetic translations for words. A sounds-like translation consists of one or more
// words that, when combined, sound like the word. Phonetic translations are based on the SSML phoneme format for
// representing a word. You can specify them in standard International Phonetic Alphabet (IPA) representation
//
//   <code>&lt;phoneme alphabet=\"ipa\" ph=\"t&#601;m&#712;&#593;to\"&gt;&lt;/phoneme&gt;</code>
//
//   or in the proprietary IBM Symbolic Phonetic Representation (SPR)
//
//   <code>&lt;phoneme alphabet=\"ibm\" ph=\"1gAstroEntxrYFXs\"&gt;&lt;/phoneme&gt;</code>
//
// **Note:** This method is currently a beta release.
//
// **See also:**
// * [Updating a custom model](https://cloud.ibm.com/docs/services/text-to-speech/custom-models.html#cuModelsUpdate)
// * [Adding words to a Japanese custom
// model](https://cloud.ibm.com/docs/services/text-to-speech/custom-entries.html#cuJapaneseAdd)
// * [Understanding customization](https://cloud.ibm.com/docs/services/text-to-speech/custom-intro.html).
func (textToSpeech *TextToSpeechV1) UpdateVoiceModel(updateVoiceModelOptions *UpdateVoiceModelOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(updateVoiceModelOptions, "updateVoiceModelOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(updateVoiceModelOptions, "updateVoiceModelOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/customizations"}
	pathParameters := []string{*updateVoiceModelOptions.CustomizationID}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(textToSpeech.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range updateVoiceModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "UpdateVoiceModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateVoiceModelOptions.Name != nil {
		body["name"] = updateVoiceModelOptions.Name
	}
	if updateVoiceModelOptions.Description != nil {
		body["description"] = updateVoiceModelOptions.Description
	}
	if updateVoiceModelOptions.Words != nil {
		body["words"] = updateVoiceModelOptions.Words
	}
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := textToSpeech.Service.Request(request, nil)
	return response, err
}

// AddWord : Add a custom word
// Adds a single word and its translation to the specified custom voice model. Adding a new translation for a word that
// already exists in a custom model overwrites the word's existing translation. A custom model can contain no more than
// 20,000 entries. You must use credentials for the instance of the service that owns a model to add a word to it.
//
// You can define sounds-like or phonetic translations for words. A sounds-like translation consists of one or more
// words that, when combined, sound like the word. Phonetic translations are based on the SSML phoneme format for
// representing a word. You can specify them in standard International Phonetic Alphabet (IPA) representation
//
//   <code>&lt;phoneme alphabet=\"ipa\" ph=\"t&#601;m&#712;&#593;to\"&gt;&lt;/phoneme&gt;</code>
//
//   or in the proprietary IBM Symbolic Phonetic Representation (SPR)
//
//   <code>&lt;phoneme alphabet=\"ibm\" ph=\"1gAstroEntxrYFXs\"&gt;&lt;/phoneme&gt;</code>
//
// **Note:** This method is currently a beta release.
//
// **See also:**
// * [Adding a single word to a custom
// model](https://cloud.ibm.com/docs/services/text-to-speech/custom-entries.html#cuWordAdd)
// * [Adding words to a Japanese custom
// model](https://cloud.ibm.com/docs/services/text-to-speech/custom-entries.html#cuJapaneseAdd)
// * [Understanding customization](https://cloud.ibm.com/docs/services/text-to-speech/custom-intro.html).
func (textToSpeech *TextToSpeechV1) AddWord(addWordOptions *AddWordOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(addWordOptions, "addWordOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(addWordOptions, "addWordOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/customizations", "words"}
	pathParameters := []string{*addWordOptions.CustomizationID, *addWordOptions.Word}

	builder := core.NewRequestBuilder(core.PUT)
	builder.ConstructHTTPURL(textToSpeech.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range addWordOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "AddWord")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if addWordOptions.Translation != nil {
		body["translation"] = addWordOptions.Translation
	}
	if addWordOptions.PartOfSpeech != nil {
		body["part_of_speech"] = addWordOptions.PartOfSpeech
	}
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := textToSpeech.Service.Request(request, nil)
	return response, err
}

// AddWords : Add custom words
// Adds one or more words and their translations to the specified custom voice model. Adding a new translation for a
// word that already exists in a custom model overwrites the word's existing translation. A custom model can contain no
// more than 20,000 entries. You must use credentials for the instance of the service that owns a model to add words to
// it.
//
// You can define sounds-like or phonetic translations for words. A sounds-like translation consists of one or more
// words that, when combined, sound like the word. Phonetic translations are based on the SSML phoneme format for
// representing a word. You can specify them in standard International Phonetic Alphabet (IPA) representation
//
//   <code>&lt;phoneme alphabet=\"ipa\" ph=\"t&#601;m&#712;&#593;to\"&gt;&lt;/phoneme&gt;</code>
//
//   or in the proprietary IBM Symbolic Phonetic Representation (SPR)
//
//   <code>&lt;phoneme alphabet=\"ibm\" ph=\"1gAstroEntxrYFXs\"&gt;&lt;/phoneme&gt;</code>
//
// **Note:** This method is currently a beta release.
//
// **See also:**
// * [Adding multiple words to a custom
// model](https://cloud.ibm.com/docs/services/text-to-speech/custom-entries.html#cuWordsAdd)
// * [Adding words to a Japanese custom
// model](https://cloud.ibm.com/docs/services/text-to-speech/custom-entries.html#cuJapaneseAdd)
// * [Understanding customization](https://cloud.ibm.com/docs/services/text-to-speech/custom-intro.html).
func (textToSpeech *TextToSpeechV1) AddWords(addWordsOptions *AddWordsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(addWordsOptions, "addWordsOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(addWordsOptions, "addWordsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/customizations", "words"}
	pathParameters := []string{*addWordsOptions.CustomizationID}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(textToSpeech.Service.Options.URL, pathSegments, pathParameters)

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
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := textToSpeech.Service.Request(request, nil)
	return response, err
}

// DeleteWord : Delete a custom word
// Deletes a single word from the specified custom voice model. You must use credentials for the instance of the service
// that owns a model to delete its words.
//
// **Note:** This method is currently a beta release.
//
// **See also:** [Deleting a word from a custom
// model](https://cloud.ibm.com/docs/services/text-to-speech/custom-entries.html#cuWordDelete).
func (textToSpeech *TextToSpeechV1) DeleteWord(deleteWordOptions *DeleteWordOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(deleteWordOptions, "deleteWordOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(deleteWordOptions, "deleteWordOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/customizations", "words"}
	pathParameters := []string{*deleteWordOptions.CustomizationID, *deleteWordOptions.Word}

	builder := core.NewRequestBuilder(core.DELETE)
	builder.ConstructHTTPURL(textToSpeech.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteWordOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "DeleteWord")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "")

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := textToSpeech.Service.Request(request, nil)
	return response, err
}

// GetWord : Get a custom word
// Gets the translation for a single word from the specified custom model. The output shows the translation as it is
// defined in the model. You must use credentials for the instance of the service that owns a model to list its words.
//
// **Note:** This method is currently a beta release.
//
// **See also:** [Querying a single word from a custom
// model](https://cloud.ibm.com/docs/services/text-to-speech/custom-entries.html#cuWordQueryModel).
func (textToSpeech *TextToSpeechV1) GetWord(getWordOptions *GetWordOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(getWordOptions, "getWordOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(getWordOptions, "getWordOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/customizations", "words"}
	pathParameters := []string{*getWordOptions.CustomizationID, *getWordOptions.Word}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(textToSpeech.Service.Options.URL, pathSegments, pathParameters)

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
		return nil, err
	}

	response, err := textToSpeech.Service.Request(request, new(Translation))
	return response, err
}

// GetGetWordResult : Retrieve result of GetWord operation
func (textToSpeech *TextToSpeechV1) GetGetWordResult(response *core.DetailedResponse) *Translation {
	result, ok := response.Result.(*Translation)
	if ok {
		return result
	}
	return nil
}

// ListWords : List custom words
// Lists all of the words and their translations for the specified custom voice model. The output shows the translations
// as they are defined in the model. You must use credentials for the instance of the service that owns a model to list
// its words.
//
// **Note:** This method is currently a beta release.
//
// **See also:** [Querying all words from a custom
// model](https://cloud.ibm.com/docs/services/text-to-speech/custom-entries.html#cuWordsQueryModel).
func (textToSpeech *TextToSpeechV1) ListWords(listWordsOptions *ListWordsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(listWordsOptions, "listWordsOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(listWordsOptions, "listWordsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/customizations", "words"}
	pathParameters := []string{*listWordsOptions.CustomizationID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(textToSpeech.Service.Options.URL, pathSegments, pathParameters)

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
		return nil, err
	}

	response, err := textToSpeech.Service.Request(request, new(Words))
	return response, err
}

// GetListWordsResult : Retrieve result of ListWords operation
func (textToSpeech *TextToSpeechV1) GetListWordsResult(response *core.DetailedResponse) *Words {
	result, ok := response.Result.(*Words)
	if ok {
		return result
	}
	return nil
}

// DeleteUserData : Delete labeled data
// Deletes all data that is associated with a specified customer ID. The method deletes all data for the customer ID,
// regardless of the method by which the information was added. The method has no effect if no data is associated with
// the customer ID. You must issue the request with credentials for the same instance of the service that was used to
// associate the customer ID with the data.
//
// You associate a customer ID with data by passing the `X-Watson-Metadata` header with a request that passes the data.
//
// **See also:** [Information security](https://cloud.ibm.com/docs/services/text-to-speech/information-security.html).
func (textToSpeech *TextToSpeechV1) DeleteUserData(deleteUserDataOptions *DeleteUserDataOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(deleteUserDataOptions, "deleteUserDataOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(deleteUserDataOptions, "deleteUserDataOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/user_data"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.DELETE)
	builder.ConstructHTTPURL(textToSpeech.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteUserDataOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "DeleteUserData")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "")

	builder.AddQuery("customer_id", fmt.Sprint(*deleteUserDataOptions.CustomerID))

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := textToSpeech.Service.Request(request, nil)
	return response, err
}

// AddWordOptions : The addWord options.
type AddWordOptions struct {

	// The customization ID (GUID) of the custom voice model. You must make the request with service credentials created
	// for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// The word that is to be added or updated for the custom voice model.
	Word *string `json:"word" validate:"required"`

	// The phonetic or sounds-like translation for the word. A phonetic translation is based on the SSML format for
	// representing the phonetic string of a word either as an IPA translation or as an IBM SPR translation. A sounds-like
	// is one or more words that, when combined, sound like the word.
	Translation *string `json:"translation" validate:"required"`

	// **Japanese only.** The part of speech for the word. The service uses the value to produce the correct intonation for
	// the word. You can create only a single entry, with or without a single part of speech, for any word; you cannot
	// create multiple entries with different parts of speech for the same word. For more information, see [Working with
	// Japanese entries](https://cloud.ibm.com/docs/services/text-to-speech/custom-rules.html#jaNotes).
	PartOfSpeech *string `json:"part_of_speech,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the AddWordOptions.PartOfSpeech property.
// **Japanese only.** The part of speech for the word. The service uses the value to produce the correct intonation for
// the word. You can create only a single entry, with or without a single part of speech, for any word; you cannot
// create multiple entries with different parts of speech for the same word. For more information, see [Working with
// Japanese entries](https://cloud.ibm.com/docs/services/text-to-speech/custom-rules.html#jaNotes).
const (
	AddWordOptions_PartOfSpeech_Dosi = "Dosi"
	AddWordOptions_PartOfSpeech_Fuku = "Fuku"
	AddWordOptions_PartOfSpeech_Gobi = "Gobi"
	AddWordOptions_PartOfSpeech_Hoka = "Hoka"
	AddWordOptions_PartOfSpeech_Jodo = "Jodo"
	AddWordOptions_PartOfSpeech_Josi = "Josi"
	AddWordOptions_PartOfSpeech_Kato = "Kato"
	AddWordOptions_PartOfSpeech_Kedo = "Kedo"
	AddWordOptions_PartOfSpeech_Keyo = "Keyo"
	AddWordOptions_PartOfSpeech_Kigo = "Kigo"
	AddWordOptions_PartOfSpeech_Koyu = "Koyu"
	AddWordOptions_PartOfSpeech_Mesi = "Mesi"
	AddWordOptions_PartOfSpeech_Reta = "Reta"
	AddWordOptions_PartOfSpeech_Stbi = "Stbi"
	AddWordOptions_PartOfSpeech_Stto = "Stto"
	AddWordOptions_PartOfSpeech_Stzo = "Stzo"
	AddWordOptions_PartOfSpeech_Suji = "Suji"
)

// NewAddWordOptions : Instantiate AddWordOptions
func (textToSpeech *TextToSpeechV1) NewAddWordOptions(customizationID string, word string, translation string) *AddWordOptions {
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

// AddWordsOptions : The addWords options.
type AddWordsOptions struct {

	// The customization ID (GUID) of the custom voice model. You must make the request with service credentials created
	// for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// The **Add custom words** method accepts an array of `Word` objects. Each object provides a word that is to be added
	// or updated for the custom voice model and the word's translation.
	//
	// The **List custom words** method returns an array of `Word` objects. Each object shows a word and its translation
	// from the custom voice model. The words are listed in alphabetical order, with uppercase letters listed before
	// lowercase letters. The array is empty if the custom model contains no words.
	Words []Word `json:"words" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewAddWordsOptions : Instantiate AddWordsOptions
func (textToSpeech *TextToSpeechV1) NewAddWordsOptions(customizationID string, words []Word) *AddWordsOptions {
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

// CreateVoiceModelOptions : The createVoiceModel options.
type CreateVoiceModelOptions struct {

	// The name of the new custom voice model.
	Name *string `json:"name" validate:"required"`

	// The language of the new custom voice model. Omit the parameter to use the the default language, `en-US`.
	Language *string `json:"language,omitempty"`

	// A description of the new custom voice model. Specifying a description is recommended.
	Description *string `json:"description,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the CreateVoiceModelOptions.Language property.
// The language of the new custom voice model. Omit the parameter to use the the default language, `en-US`.
const (
	CreateVoiceModelOptions_Language_DeDe = "de-DE"
	CreateVoiceModelOptions_Language_EnGb = "en-GB"
	CreateVoiceModelOptions_Language_EnUs = "en-US"
	CreateVoiceModelOptions_Language_EsEs = "es-ES"
	CreateVoiceModelOptions_Language_EsLa = "es-LA"
	CreateVoiceModelOptions_Language_EsUs = "es-US"
	CreateVoiceModelOptions_Language_FrFr = "fr-FR"
	CreateVoiceModelOptions_Language_ItIt = "it-IT"
	CreateVoiceModelOptions_Language_JaJp = "ja-JP"
	CreateVoiceModelOptions_Language_PtBr = "pt-BR"
)

// NewCreateVoiceModelOptions : Instantiate CreateVoiceModelOptions
func (textToSpeech *TextToSpeechV1) NewCreateVoiceModelOptions(name string) *CreateVoiceModelOptions {
	return &CreateVoiceModelOptions{
		Name: core.StringPtr(name),
	}
}

// SetName : Allow user to set Name
func (options *CreateVoiceModelOptions) SetName(name string) *CreateVoiceModelOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetLanguage : Allow user to set Language
func (options *CreateVoiceModelOptions) SetLanguage(language string) *CreateVoiceModelOptions {
	options.Language = core.StringPtr(language)
	return options
}

// SetDescription : Allow user to set Description
func (options *CreateVoiceModelOptions) SetDescription(description string) *CreateVoiceModelOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateVoiceModelOptions) SetHeaders(param map[string]string) *CreateVoiceModelOptions {
	options.Headers = param
	return options
}

// DeleteUserDataOptions : The deleteUserData options.
type DeleteUserDataOptions struct {

	// The customer ID for which all data is to be deleted.
	CustomerID *string `json:"customer_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteUserDataOptions : Instantiate DeleteUserDataOptions
func (textToSpeech *TextToSpeechV1) NewDeleteUserDataOptions(customerID string) *DeleteUserDataOptions {
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

// DeleteVoiceModelOptions : The deleteVoiceModel options.
type DeleteVoiceModelOptions struct {

	// The customization ID (GUID) of the custom voice model. You must make the request with service credentials created
	// for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteVoiceModelOptions : Instantiate DeleteVoiceModelOptions
func (textToSpeech *TextToSpeechV1) NewDeleteVoiceModelOptions(customizationID string) *DeleteVoiceModelOptions {
	return &DeleteVoiceModelOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *DeleteVoiceModelOptions) SetCustomizationID(customizationID string) *DeleteVoiceModelOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteVoiceModelOptions) SetHeaders(param map[string]string) *DeleteVoiceModelOptions {
	options.Headers = param
	return options
}

// DeleteWordOptions : The deleteWord options.
type DeleteWordOptions struct {

	// The customization ID (GUID) of the custom voice model. You must make the request with service credentials created
	// for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// The word that is to be deleted from the custom voice model.
	Word *string `json:"word" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteWordOptions : Instantiate DeleteWordOptions
func (textToSpeech *TextToSpeechV1) NewDeleteWordOptions(customizationID string, word string) *DeleteWordOptions {
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

// GetPronunciationOptions : The getPronunciation options.
type GetPronunciationOptions struct {

	// The word for which the pronunciation is requested.
	Text *string `json:"text" validate:"required"`

	// A voice that specifies the language in which the pronunciation is to be returned. All voices for the same language
	// (for example, `en-US`) return the same translation.
	Voice *string `json:"voice,omitempty"`

	// The phoneme format in which to return the pronunciation. Omit the parameter to obtain the pronunciation in the
	// default format.
	Format *string `json:"format,omitempty"`

	// The customization ID (GUID) of a custom voice model for which the pronunciation is to be returned. The language of a
	// specified custom model must match the language of the specified voice. If the word is not defined in the specified
	// custom model, the service returns the default translation for the custom model's language. You must make the request
	// with service credentials created for the instance of the service that owns the custom model. Omit the parameter to
	// see the translation for the specified voice with no customization.
	CustomizationID *string `json:"customization_id,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the GetPronunciationOptions.Voice property.
// A voice that specifies the language in which the pronunciation is to be returned. All voices for the same language
// (for example, `en-US`) return the same translation.
const (
	GetPronunciationOptions_Voice_DeDeBirgitv2voice    = "de-DE_BirgitV2Voice"
	GetPronunciationOptions_Voice_DeDeBirgitvoice      = "de-DE_BirgitVoice"
	GetPronunciationOptions_Voice_DeDeDieterv2voice    = "de-DE_DieterV2Voice"
	GetPronunciationOptions_Voice_DeDeDietervoice      = "de-DE_DieterVoice"
	GetPronunciationOptions_Voice_EnGbKatevoice        = "en-GB_KateVoice"
	GetPronunciationOptions_Voice_EnUsAllisonv2voice   = "en-US_AllisonV2Voice"
	GetPronunciationOptions_Voice_EnUsAllisonvoice     = "en-US_AllisonVoice"
	GetPronunciationOptions_Voice_EnUsLisav2voice      = "en-US_LisaV2Voice"
	GetPronunciationOptions_Voice_EnUsLisavoice        = "en-US_LisaVoice"
	GetPronunciationOptions_Voice_EnUsMichaelv2voice   = "en-US_MichaelV2Voice"
	GetPronunciationOptions_Voice_EnUsMichaelvoice     = "en-US_MichaelVoice"
	GetPronunciationOptions_Voice_EsEsEnriquevoice     = "es-ES_EnriqueVoice"
	GetPronunciationOptions_Voice_EsEsLauravoice       = "es-ES_LauraVoice"
	GetPronunciationOptions_Voice_EsLaSofiavoice       = "es-LA_SofiaVoice"
	GetPronunciationOptions_Voice_EsUsSofiavoice       = "es-US_SofiaVoice"
	GetPronunciationOptions_Voice_FrFrReneevoice       = "fr-FR_ReneeVoice"
	GetPronunciationOptions_Voice_ItItFrancescav2voice = "it-IT_FrancescaV2Voice"
	GetPronunciationOptions_Voice_ItItFrancescavoice   = "it-IT_FrancescaVoice"
	GetPronunciationOptions_Voice_JaJpEmivoice         = "ja-JP_EmiVoice"
	GetPronunciationOptions_Voice_PtBrIsabelavoice     = "pt-BR_IsabelaVoice"
)

// Constants associated with the GetPronunciationOptions.Format property.
// The phoneme format in which to return the pronunciation. Omit the parameter to obtain the pronunciation in the
// default format.
const (
	GetPronunciationOptions_Format_Ibm = "ibm"
	GetPronunciationOptions_Format_Ipa = "ipa"
)

// NewGetPronunciationOptions : Instantiate GetPronunciationOptions
func (textToSpeech *TextToSpeechV1) NewGetPronunciationOptions(text string) *GetPronunciationOptions {
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

// GetVoiceModelOptions : The getVoiceModel options.
type GetVoiceModelOptions struct {

	// The customization ID (GUID) of the custom voice model. You must make the request with service credentials created
	// for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetVoiceModelOptions : Instantiate GetVoiceModelOptions
func (textToSpeech *TextToSpeechV1) NewGetVoiceModelOptions(customizationID string) *GetVoiceModelOptions {
	return &GetVoiceModelOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *GetVoiceModelOptions) SetCustomizationID(customizationID string) *GetVoiceModelOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetVoiceModelOptions) SetHeaders(param map[string]string) *GetVoiceModelOptions {
	options.Headers = param
	return options
}

// GetVoiceOptions : The getVoice options.
type GetVoiceOptions struct {

	// The voice for which information is to be returned.
	Voice *string `json:"voice" validate:"required"`

	// The customization ID (GUID) of a custom voice model for which information is to be returned. You must make the
	// request with service credentials created for the instance of the service that owns the custom model. Omit the
	// parameter to see information about the specified voice with no customization.
	CustomizationID *string `json:"customization_id,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the GetVoiceOptions.Voice property.
// The voice for which information is to be returned.
const (
	GetVoiceOptions_Voice_DeDeBirgitv2voice    = "de-DE_BirgitV2Voice"
	GetVoiceOptions_Voice_DeDeBirgitvoice      = "de-DE_BirgitVoice"
	GetVoiceOptions_Voice_DeDeDieterv2voice    = "de-DE_DieterV2Voice"
	GetVoiceOptions_Voice_DeDeDietervoice      = "de-DE_DieterVoice"
	GetVoiceOptions_Voice_EnGbKatevoice        = "en-GB_KateVoice"
	GetVoiceOptions_Voice_EnUsAllisonv2voice   = "en-US_AllisonV2Voice"
	GetVoiceOptions_Voice_EnUsAllisonvoice     = "en-US_AllisonVoice"
	GetVoiceOptions_Voice_EnUsLisav2voice      = "en-US_LisaV2Voice"
	GetVoiceOptions_Voice_EnUsLisavoice        = "en-US_LisaVoice"
	GetVoiceOptions_Voice_EnUsMichaelv2voice   = "en-US_MichaelV2Voice"
	GetVoiceOptions_Voice_EnUsMichaelvoice     = "en-US_MichaelVoice"
	GetVoiceOptions_Voice_EsEsEnriquevoice     = "es-ES_EnriqueVoice"
	GetVoiceOptions_Voice_EsEsLauravoice       = "es-ES_LauraVoice"
	GetVoiceOptions_Voice_EsLaSofiavoice       = "es-LA_SofiaVoice"
	GetVoiceOptions_Voice_EsUsSofiavoice       = "es-US_SofiaVoice"
	GetVoiceOptions_Voice_FrFrReneevoice       = "fr-FR_ReneeVoice"
	GetVoiceOptions_Voice_ItItFrancescav2voice = "it-IT_FrancescaV2Voice"
	GetVoiceOptions_Voice_ItItFrancescavoice   = "it-IT_FrancescaVoice"
	GetVoiceOptions_Voice_JaJpEmivoice         = "ja-JP_EmiVoice"
	GetVoiceOptions_Voice_PtBrIsabelavoice     = "pt-BR_IsabelaVoice"
)

// NewGetVoiceOptions : Instantiate GetVoiceOptions
func (textToSpeech *TextToSpeechV1) NewGetVoiceOptions(voice string) *GetVoiceOptions {
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

// GetWordOptions : The getWord options.
type GetWordOptions struct {

	// The customization ID (GUID) of the custom voice model. You must make the request with service credentials created
	// for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// The word that is to be queried from the custom voice model.
	Word *string `json:"word" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetWordOptions : Instantiate GetWordOptions
func (textToSpeech *TextToSpeechV1) NewGetWordOptions(customizationID string, word string) *GetWordOptions {
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

// ListVoiceModelsOptions : The listVoiceModels options.
type ListVoiceModelsOptions struct {

	// The language for which custom voice models that are owned by the requesting service credentials are to be returned.
	// Omit the parameter to see all custom voice models that are owned by the requester.
	Language *string `json:"language,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the ListVoiceModelsOptions.Language property.
// The language for which custom voice models that are owned by the requesting service credentials are to be returned.
// Omit the parameter to see all custom voice models that are owned by the requester.
const (
	ListVoiceModelsOptions_Language_DeDe = "de-DE"
	ListVoiceModelsOptions_Language_EnGb = "en-GB"
	ListVoiceModelsOptions_Language_EnUs = "en-US"
	ListVoiceModelsOptions_Language_EsEs = "es-ES"
	ListVoiceModelsOptions_Language_EsLa = "es-LA"
	ListVoiceModelsOptions_Language_EsUs = "es-US"
	ListVoiceModelsOptions_Language_FrFr = "fr-FR"
	ListVoiceModelsOptions_Language_ItIt = "it-IT"
	ListVoiceModelsOptions_Language_JaJp = "ja-JP"
	ListVoiceModelsOptions_Language_PtBr = "pt-BR"
)

// NewListVoiceModelsOptions : Instantiate ListVoiceModelsOptions
func (textToSpeech *TextToSpeechV1) NewListVoiceModelsOptions() *ListVoiceModelsOptions {
	return &ListVoiceModelsOptions{}
}

// SetLanguage : Allow user to set Language
func (options *ListVoiceModelsOptions) SetLanguage(language string) *ListVoiceModelsOptions {
	options.Language = core.StringPtr(language)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListVoiceModelsOptions) SetHeaders(param map[string]string) *ListVoiceModelsOptions {
	options.Headers = param
	return options
}

// ListVoicesOptions : The listVoices options.
type ListVoicesOptions struct {

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListVoicesOptions : Instantiate ListVoicesOptions
func (textToSpeech *TextToSpeechV1) NewListVoicesOptions() *ListVoicesOptions {
	return &ListVoicesOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListVoicesOptions) SetHeaders(param map[string]string) *ListVoicesOptions {
	options.Headers = param
	return options
}

// ListWordsOptions : The listWords options.
type ListWordsOptions struct {

	// The customization ID (GUID) of the custom voice model. You must make the request with service credentials created
	// for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListWordsOptions : Instantiate ListWordsOptions
func (textToSpeech *TextToSpeechV1) NewListWordsOptions(customizationID string) *ListWordsOptions {
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

// Pronunciation : Pronunciation struct
type Pronunciation struct {

	// The pronunciation of the specified text in the requested voice and format. If a custom voice model is specified, the
	// pronunciation also reflects that custom voice.
	Pronunciation *string `json:"pronunciation" validate:"required"`
}

// SupportedFeatures : Describes the additional service features that are supported with the voice.
type SupportedFeatures struct {

	// If `true`, the voice can be customized; if `false`, the voice cannot be customized. (Same as `customizable`.).
	CustomPronunciation *bool `json:"custom_pronunciation" validate:"required"`

	// If `true`, the voice can be transformed by using the SSML &lt;voice-transformation&gt; element; if `false`, the
	// voice cannot be transformed.
	VoiceTransformation *bool `json:"voice_transformation" validate:"required"`
}

// SynthesizeOptions : The synthesize options.
type SynthesizeOptions struct {

	// The text to synthesize.
	Text *string `json:"text" validate:"required"`

	// The voice to use for synthesis.
	Voice *string `json:"voice,omitempty"`

	// The customization ID (GUID) of a custom voice model to use for the synthesis. If a custom voice model is specified,
	// it is guaranteed to work only if it matches the language of the indicated voice. You must make the request with
	// service credentials created for the instance of the service that owns the custom model. Omit the parameter to use
	// the specified voice with no customization.
	CustomizationID *string `json:"customization_id,omitempty"`

	// The requested format (MIME type) of the audio. You can use the `Accept` header or the `accept` parameter to specify
	// the audio format. For more information about specifying an audio format, see **Audio formats (accept types)** in the
	// method description.
	//
	// Default: `audio/ogg;codecs=opus`.
	Accept *string `json:"Accept,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the SynthesizeOptions.Voice property.
// The voice to use for synthesis.
const (
	SynthesizeOptions_Voice_DeDeBirgitv2voice    = "de-DE_BirgitV2Voice"
	SynthesizeOptions_Voice_DeDeBirgitvoice      = "de-DE_BirgitVoice"
	SynthesizeOptions_Voice_DeDeDieterv2voice    = "de-DE_DieterV2Voice"
	SynthesizeOptions_Voice_DeDeDietervoice      = "de-DE_DieterVoice"
	SynthesizeOptions_Voice_EnGbKatevoice        = "en-GB_KateVoice"
	SynthesizeOptions_Voice_EnUsAllisonv2voice   = "en-US_AllisonV2Voice"
	SynthesizeOptions_Voice_EnUsAllisonvoice     = "en-US_AllisonVoice"
	SynthesizeOptions_Voice_EnUsLisav2voice      = "en-US_LisaV2Voice"
	SynthesizeOptions_Voice_EnUsLisavoice        = "en-US_LisaVoice"
	SynthesizeOptions_Voice_EnUsMichaelv2voice   = "en-US_MichaelV2Voice"
	SynthesizeOptions_Voice_EnUsMichaelvoice     = "en-US_MichaelVoice"
	SynthesizeOptions_Voice_EsEsEnriquevoice     = "es-ES_EnriqueVoice"
	SynthesizeOptions_Voice_EsEsLauravoice       = "es-ES_LauraVoice"
	SynthesizeOptions_Voice_EsLaSofiavoice       = "es-LA_SofiaVoice"
	SynthesizeOptions_Voice_EsUsSofiavoice       = "es-US_SofiaVoice"
	SynthesizeOptions_Voice_FrFrReneevoice       = "fr-FR_ReneeVoice"
	SynthesizeOptions_Voice_ItItFrancescav2voice = "it-IT_FrancescaV2Voice"
	SynthesizeOptions_Voice_ItItFrancescavoice   = "it-IT_FrancescaVoice"
	SynthesizeOptions_Voice_JaJpEmivoice         = "ja-JP_EmiVoice"
	SynthesizeOptions_Voice_PtBrIsabelavoice     = "pt-BR_IsabelaVoice"
)

// Constants associated with the SynthesizeOptions.Accept property.
// The requested format (MIME type) of the audio. You can use the `Accept` header or the `accept` parameter to specify
// the audio format. For more information about specifying an audio format, see **Audio formats (accept types)** in the
// method description.
//
// Default: `audio/ogg;codecs=opus`.
const (
	SynthesizeOptions_Accept_AudioBasic            = "audio/basic"
	SynthesizeOptions_Accept_AudioFlac             = "audio/flac"
	SynthesizeOptions_Accept_AudioL16              = "audio/l16"
	SynthesizeOptions_Accept_AudioMp3              = "audio/mp3"
	SynthesizeOptions_Accept_AudioMpeg             = "audio/mpeg"
	SynthesizeOptions_Accept_AudioMulaw            = "audio/mulaw"
	SynthesizeOptions_Accept_AudioOgg              = "audio/ogg"
	SynthesizeOptions_Accept_AudioOggCodecsOpus    = "audio/ogg;codecs=opus"
	SynthesizeOptions_Accept_AudioOggCodecsVorbis  = "audio/ogg;codecs=vorbis"
	SynthesizeOptions_Accept_AudioWav              = "audio/wav"
	SynthesizeOptions_Accept_AudioWebm             = "audio/webm"
	SynthesizeOptions_Accept_AudioWebmCodecsOpus   = "audio/webm;codecs=opus"
	SynthesizeOptions_Accept_AudioWebmCodecsVorbis = "audio/webm;codecs=vorbis"
)

// NewSynthesizeOptions : Instantiate SynthesizeOptions
func (textToSpeech *TextToSpeechV1) NewSynthesizeOptions(text string) *SynthesizeOptions {
	return &SynthesizeOptions{
		Text: core.StringPtr(text),
	}
}

// SetText : Allow user to set Text
func (options *SynthesizeOptions) SetText(text string) *SynthesizeOptions {
	options.Text = core.StringPtr(text)
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

// SetAccept : Allow user to set Accept
func (options *SynthesizeOptions) SetAccept(accept string) *SynthesizeOptions {
	options.Accept = core.StringPtr(accept)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *SynthesizeOptions) SetHeaders(param map[string]string) *SynthesizeOptions {
	options.Headers = param
	return options
}

// Translation : Translation struct
type Translation struct {

	// The phonetic or sounds-like translation for the word. A phonetic translation is based on the SSML format for
	// representing the phonetic string of a word either as an IPA translation or as an IBM SPR translation. A sounds-like
	// is one or more words that, when combined, sound like the word.
	Translation *string `json:"translation" validate:"required"`

	// **Japanese only.** The part of speech for the word. The service uses the value to produce the correct intonation for
	// the word. You can create only a single entry, with or without a single part of speech, for any word; you cannot
	// create multiple entries with different parts of speech for the same word. For more information, see [Working with
	// Japanese entries](https://cloud.ibm.com/docs/services/text-to-speech/custom-rules.html#jaNotes).
	PartOfSpeech *string `json:"part_of_speech,omitempty"`
}

// Constants associated with the Translation.PartOfSpeech property.
// **Japanese only.** The part of speech for the word. The service uses the value to produce the correct intonation for
// the word. You can create only a single entry, with or without a single part of speech, for any word; you cannot
// create multiple entries with different parts of speech for the same word. For more information, see [Working with
// Japanese entries](https://cloud.ibm.com/docs/services/text-to-speech/custom-rules.html#jaNotes).
const (
	Translation_PartOfSpeech_Dosi = "Dosi"
	Translation_PartOfSpeech_Fuku = "Fuku"
	Translation_PartOfSpeech_Gobi = "Gobi"
	Translation_PartOfSpeech_Hoka = "Hoka"
	Translation_PartOfSpeech_Jodo = "Jodo"
	Translation_PartOfSpeech_Josi = "Josi"
	Translation_PartOfSpeech_Kato = "Kato"
	Translation_PartOfSpeech_Kedo = "Kedo"
	Translation_PartOfSpeech_Keyo = "Keyo"
	Translation_PartOfSpeech_Kigo = "Kigo"
	Translation_PartOfSpeech_Koyu = "Koyu"
	Translation_PartOfSpeech_Mesi = "Mesi"
	Translation_PartOfSpeech_Reta = "Reta"
	Translation_PartOfSpeech_Stbi = "Stbi"
	Translation_PartOfSpeech_Stto = "Stto"
	Translation_PartOfSpeech_Stzo = "Stzo"
	Translation_PartOfSpeech_Suji = "Suji"
)

// UpdateVoiceModelOptions : The updateVoiceModel options.
type UpdateVoiceModelOptions struct {

	// The customization ID (GUID) of the custom voice model. You must make the request with service credentials created
	// for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// A new name for the custom voice model.
	Name *string `json:"name,omitempty"`

	// A new description for the custom voice model.
	Description *string `json:"description,omitempty"`

	// An array of `Word` objects that provides the words and their translations that are to be added or updated for the
	// custom voice model. Pass an empty array to make no additions or updates.
	Words []Word `json:"words,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewUpdateVoiceModelOptions : Instantiate UpdateVoiceModelOptions
func (textToSpeech *TextToSpeechV1) NewUpdateVoiceModelOptions(customizationID string) *UpdateVoiceModelOptions {
	return &UpdateVoiceModelOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *UpdateVoiceModelOptions) SetCustomizationID(customizationID string) *UpdateVoiceModelOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetName : Allow user to set Name
func (options *UpdateVoiceModelOptions) SetName(name string) *UpdateVoiceModelOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetDescription : Allow user to set Description
func (options *UpdateVoiceModelOptions) SetDescription(description string) *UpdateVoiceModelOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetWords : Allow user to set Words
func (options *UpdateVoiceModelOptions) SetWords(words []Word) *UpdateVoiceModelOptions {
	options.Words = words
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateVoiceModelOptions) SetHeaders(param map[string]string) *UpdateVoiceModelOptions {
	options.Headers = param
	return options
}

// Voice : Voice struct
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

	// Describes the additional service features that are supported with the voice.
	SupportedFeatures *SupportedFeatures `json:"supported_features" validate:"required"`

	// Returns information about a specified custom voice model. This field is returned only by the **Get a voice** method
	// and only when you specify the customization ID of a custom voice model.
	Customization *VoiceModel `json:"customization,omitempty"`
}

// VoiceModel : VoiceModel struct
type VoiceModel struct {

	// The customization ID (GUID) of the custom voice model. The **Create a custom model** method returns only this field.
	// It does not not return the other fields of this object.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// The name of the custom voice model.
	Name *string `json:"name,omitempty"`

	// The language identifier of the custom voice model (for example, `en-US`).
	Language *string `json:"language,omitempty"`

	// The GUID of the service credentials for the instance of the service that owns the custom voice model.
	Owner *string `json:"owner,omitempty"`

	// The date and time in Coordinated Universal Time (UTC) at which the custom voice model was created. The value is
	// provided in full ISO 8601 format (`YYYY-MM-DDThh:mm:ss.sTZD`).
	Created *string `json:"created,omitempty"`

	// The date and time in Coordinated Universal Time (UTC) at which the custom voice model was last modified. Equals
	// `created` when a new voice model is first added but has yet to be updated. The value is provided in full ISO 8601
	// format (`YYYY-MM-DDThh:mm:ss.sTZD`).
	LastModified *string `json:"last_modified,omitempty"`

	// The description of the custom voice model.
	Description *string `json:"description,omitempty"`

	// An array of `Word` objects that lists the words and their translations from the custom voice model. The words are
	// listed in alphabetical order, with uppercase letters listed before lowercase letters. The array is empty if the
	// custom model contains no words. This field is returned only by the **Get a voice** method and only when you specify
	// the customization ID of a custom voice model.
	Words []Word `json:"words,omitempty"`
}

// VoiceModels : VoiceModels struct
type VoiceModels struct {

	// An array of `VoiceModel` objects that provides information about each available custom voice model. The array is
	// empty if the requesting service credentials own no custom voice models (if no language is specified) or own no
	// custom voice models for the specified language.
	Customizations []VoiceModel `json:"customizations" validate:"required"`
}

// Voices : Voices struct
type Voices struct {

	// A list of available voices.
	Voices []Voice `json:"voices" validate:"required"`
}

// Word : Word struct
type Word struct {

	// A word from the custom voice model.
	Word *string `json:"word" validate:"required"`

	// The phonetic or sounds-like translation for the word. A phonetic translation is based on the SSML format for
	// representing the phonetic string of a word either as an IPA or IBM SPR translation. A sounds-like translation
	// consists of one or more words that, when combined, sound like the word.
	Translation *string `json:"translation" validate:"required"`

	// **Japanese only.** The part of speech for the word. The service uses the value to produce the correct intonation for
	// the word. You can create only a single entry, with or without a single part of speech, for any word; you cannot
	// create multiple entries with different parts of speech for the same word. For more information, see [Working with
	// Japanese entries](https://cloud.ibm.com/docs/services/text-to-speech/custom-rules.html#jaNotes).
	PartOfSpeech *string `json:"part_of_speech,omitempty"`
}

// Constants associated with the Word.PartOfSpeech property.
// **Japanese only.** The part of speech for the word. The service uses the value to produce the correct intonation for
// the word. You can create only a single entry, with or without a single part of speech, for any word; you cannot
// create multiple entries with different parts of speech for the same word. For more information, see [Working with
// Japanese entries](https://cloud.ibm.com/docs/services/text-to-speech/custom-rules.html#jaNotes).
const (
	Word_PartOfSpeech_Dosi = "Dosi"
	Word_PartOfSpeech_Fuku = "Fuku"
	Word_PartOfSpeech_Gobi = "Gobi"
	Word_PartOfSpeech_Hoka = "Hoka"
	Word_PartOfSpeech_Jodo = "Jodo"
	Word_PartOfSpeech_Josi = "Josi"
	Word_PartOfSpeech_Kato = "Kato"
	Word_PartOfSpeech_Kedo = "Kedo"
	Word_PartOfSpeech_Keyo = "Keyo"
	Word_PartOfSpeech_Kigo = "Kigo"
	Word_PartOfSpeech_Koyu = "Koyu"
	Word_PartOfSpeech_Mesi = "Mesi"
	Word_PartOfSpeech_Reta = "Reta"
	Word_PartOfSpeech_Stbi = "Stbi"
	Word_PartOfSpeech_Stto = "Stto"
	Word_PartOfSpeech_Stzo = "Stzo"
	Word_PartOfSpeech_Suji = "Suji"
)

// Words : Words struct
type Words struct {

	// The **Add custom words** method accepts an array of `Word` objects. Each object provides a word that is to be added
	// or updated for the custom voice model and the word's translation.
	//
	// The **List custom words** method returns an array of `Word` objects. Each object shows a word and its translation
	// from the custom voice model. The words are listed in alphabetical order, with uppercase letters listed before
	// lowercase letters. The array is empty if the custom model contains no words.
	Words []Word `json:"words" validate:"required"`
}
