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

// Package texttospeechv1 : Operations and models for the TextToSpeechV1 service
package texttospeechv1

import (
	"fmt"
	"github.com/IBM/go-sdk-core/core"
	common "github.com/watson-developer-cloud/go-sdk/common"
)

// TextToSpeechV1 : The IBM&reg; Text to Speech service provides APIs that use IBM's speech-synthesis capabilities to
// synthesize text into natural-sounding speech in a variety of languages, dialects, and voices. The service supports at
// least one male or female voice, sometimes both, for each language. The audio is streamed back to the client with
// minimal delay.
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
// Representation (SPR). The Arabic, Chinese, Dutch, and Korean languages support only IPA.
//
// Version: 1.0.0
// See: https://cloud.ibm.com/docs/text-to-speech/
type TextToSpeechV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://stream.watsonplatform.net/text-to-speech/api"

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

// SetServiceURL sets the service URL
func (textToSpeech *TextToSpeechV1) SetServiceURL(url string) error {
	return textToSpeech.Service.SetServiceURL(url)
}

// DisableSSLVerification bypasses verification of the server's SSL certificate
func (textToSpeech *TextToSpeechV1) DisableSSLVerification() {
	textToSpeech.Service.DisableSSLVerification()
}

// ListVoices : List voices
// Lists all voices available for use with the service. The information includes the name, language, gender, and other
// details about the voice. To see information about a specific voice, use the **Get a voice** method.
//
// **See also:** [Listing all available
// voices](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-voices#listVoices).
func (textToSpeech *TextToSpeechV1) ListVoices(listVoicesOptions *ListVoicesOptions) (result *Voices, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listVoicesOptions, "listVoicesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/voices"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(textToSpeech.Service.Options.URL, pathSegments, pathParameters)
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

	response, err = textToSpeech.Service.Request(request, new(Voices))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*Voices)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// GetVoice : Get a voice
// Gets information about the specified voice. The information includes the name, language, gender, and other details
// about the voice. Specify a customization ID to obtain information for a custom voice model that is defined for the
// language of the specified voice. To list information about all available voices, use the **List voices** method.
//
// **See also:** [Listing a specific
// voice](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-voices#listVoice).
func (textToSpeech *TextToSpeechV1) GetVoice(getVoiceOptions *GetVoiceOptions) (result *Voice, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getVoiceOptions, "getVoiceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getVoiceOptions, "getVoiceOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/voices"}
	pathParameters := []string{*getVoiceOptions.Voice}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(textToSpeech.Service.Options.URL, pathSegments, pathParameters)
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

	response, err = textToSpeech.Service.Request(request, new(Voice))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*Voice)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

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
// ### Warning messages
//
//  If a request includes invalid query parameters, the service returns a `Warnings` response header that provides
// messages about the invalid parameters. The warning includes a descriptive message and a list of invalid argument
// strings. For example, a message such as `"Unknown arguments:"` or `"Unknown url query arguments:"` followed by a list
// of the form `"{invalid_arg_1}, {invalid_arg_2}."` The request succeeds despite the warnings.
func (textToSpeech *TextToSpeechV1) Synthesize(synthesizeOptions *SynthesizeOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(synthesizeOptions, "synthesizeOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(synthesizeOptions, "synthesizeOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/synthesize"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(textToSpeech.Service.Options.URL, pathSegments, pathParameters)
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

	response, err = textToSpeech.Service.Request(request, new(io.ReadCloser))
	if err == nil {
		var ok bool
		result, ok = response.Result.(io.ReadCloser)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// GetPronunciation : Get pronunciation
// Gets the phonetic pronunciation for the specified word. You can request the pronunciation for a specific format. You
// can also request the pronunciation for a specific voice to see the default translation for the language of that voice
// or for a specific custom voice model to see the translation for that voice model.
//
// **Note:** This method is currently a beta release.
//
// **See also:** [Querying a word from a
// language](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customWords#cuWordsQueryLanguage).
func (textToSpeech *TextToSpeechV1) GetPronunciation(getPronunciationOptions *GetPronunciationOptions) (result *Pronunciation, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getPronunciationOptions, "getPronunciationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getPronunciationOptions, "getPronunciationOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/pronunciation"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(textToSpeech.Service.Options.URL, pathSegments, pathParameters)
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

	response, err = textToSpeech.Service.Request(request, new(Pronunciation))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*Pronunciation)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// CreateVoiceModel : Create a custom model
// Creates a new empty custom voice model. You must specify a name for the new custom model. You can optionally specify
// the language and a description for the new model. The model is owned by the instance of the service whose credentials
// are used to create it.
//
// **Note:** This method is currently a beta release.
//
// **See also:** [Creating a custom
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customModels#cuModelsCreate).
func (textToSpeech *TextToSpeechV1) CreateVoiceModel(createVoiceModelOptions *CreateVoiceModelOptions) (result *VoiceModel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createVoiceModelOptions, "createVoiceModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createVoiceModelOptions, "createVoiceModelOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/customizations"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(textToSpeech.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

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
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = textToSpeech.Service.Request(request, new(VoiceModel))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*VoiceModel)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
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
// models](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customModels#cuModelsQueryAll).
func (textToSpeech *TextToSpeechV1) ListVoiceModels(listVoiceModelsOptions *ListVoiceModelsOptions) (result *VoiceModels, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listVoiceModelsOptions, "listVoiceModelsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/customizations"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(textToSpeech.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

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
		return
	}

	response, err = textToSpeech.Service.Request(request, new(VoiceModels))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*VoiceModels)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
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
//   <code>&lt;phoneme alphabet="ipa" ph="t&#601;m&#712;&#593;to"&gt;&lt;/phoneme&gt;</code>
//
//   or in the proprietary IBM Symbolic Phonetic Representation (SPR)
//
//   <code>&lt;phoneme alphabet="ibm" ph="1gAstroEntxrYFXs"&gt;&lt;/phoneme&gt;</code>
//
// **Note:** This method is currently a beta release.
//
// **See also:**
// * [Updating a custom
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customModels#cuModelsUpdate)
// * [Adding words to a Japanese custom
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customWords#cuJapaneseAdd)
// * [Understanding
// customization](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customIntro#customIntro).
func (textToSpeech *TextToSpeechV1) UpdateVoiceModel(updateVoiceModelOptions *UpdateVoiceModelOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateVoiceModelOptions, "updateVoiceModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateVoiceModelOptions, "updateVoiceModelOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/customizations"}
	pathParameters := []string{*updateVoiceModelOptions.CustomizationID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(textToSpeech.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

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

// GetVoiceModel : Get a custom model
// Gets all information about a specified custom voice model. In addition to metadata such as the name and description
// of the voice model, the output includes the words and their translations as defined in the model. To see just the
// metadata for a voice model, use the **List custom models** method.
//
// **Note:** This method is currently a beta release.
//
// **See also:** [Querying a custom
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customModels#cuModelsQuery).
func (textToSpeech *TextToSpeechV1) GetVoiceModel(getVoiceModelOptions *GetVoiceModelOptions) (result *VoiceModel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getVoiceModelOptions, "getVoiceModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getVoiceModelOptions, "getVoiceModelOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/customizations"}
	pathParameters := []string{*getVoiceModelOptions.CustomizationID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(textToSpeech.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

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
		return
	}

	response, err = textToSpeech.Service.Request(request, new(VoiceModel))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*VoiceModel)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// DeleteVoiceModel : Delete a custom model
// Deletes the specified custom voice model. You must use credentials for the instance of the service that owns a model
// to delete it.
//
// **Note:** This method is currently a beta release.
//
// **See also:** [Deleting a custom
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customModels#cuModelsDelete).
func (textToSpeech *TextToSpeechV1) DeleteVoiceModel(deleteVoiceModelOptions *DeleteVoiceModelOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteVoiceModelOptions, "deleteVoiceModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteVoiceModelOptions, "deleteVoiceModelOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/customizations"}
	pathParameters := []string{*deleteVoiceModelOptions.CustomizationID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(textToSpeech.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteVoiceModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "DeleteVoiceModel")
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
// Adds one or more words and their translations to the specified custom voice model. Adding a new translation for a
// word that already exists in a custom model overwrites the word's existing translation. A custom model can contain no
// more than 20,000 entries. You must use credentials for the instance of the service that owns a model to add words to
// it.
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
// **Note:** This method is currently a beta release.
//
// **See also:**
// * [Adding multiple words to a custom
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customWords#cuWordsAdd)
// * [Adding words to a Japanese custom
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customWords#cuJapaneseAdd)
// * [Understanding
// customization](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customIntro#customIntro).
func (textToSpeech *TextToSpeechV1) AddWords(addWordsOptions *AddWordsOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addWordsOptions, "addWordsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addWordsOptions, "addWordsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/customizations", "words"}
	pathParameters := []string{*addWordsOptions.CustomizationID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(textToSpeech.Service.Options.URL, pathSegments, pathParameters)
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
// Lists all of the words and their translations for the specified custom voice model. The output shows the translations
// as they are defined in the model. You must use credentials for the instance of the service that owns a model to list
// its words.
//
// **Note:** This method is currently a beta release.
//
// **See also:** [Querying all words from a custom
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customWords#cuWordsQueryModel).
func (textToSpeech *TextToSpeechV1) ListWords(listWordsOptions *ListWordsOptions) (result *Words, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listWordsOptions, "listWordsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listWordsOptions, "listWordsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/customizations", "words"}
	pathParameters := []string{*listWordsOptions.CustomizationID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(textToSpeech.Service.Options.URL, pathSegments, pathParameters)
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

	response, err = textToSpeech.Service.Request(request, new(Words))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*Words)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
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
//   <code>&lt;phoneme alphabet="ipa" ph="t&#601;m&#712;&#593;to"&gt;&lt;/phoneme&gt;</code>
//
//   or in the proprietary IBM Symbolic Phonetic Representation (SPR)
//
//   <code>&lt;phoneme alphabet="ibm" ph="1gAstroEntxrYFXs"&gt;&lt;/phoneme&gt;</code>
//
// **Note:** This method is currently a beta release.
//
// **See also:**
// * [Adding a single word to a custom
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customWords#cuWordAdd)
// * [Adding words to a Japanese custom
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customWords#cuJapaneseAdd)
// * [Understanding
// customization](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customIntro#customIntro).
func (textToSpeech *TextToSpeechV1) AddWord(addWordOptions *AddWordOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addWordOptions, "addWordOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addWordOptions, "addWordOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/customizations", "words"}
	pathParameters := []string{*addWordOptions.CustomizationID, *addWordOptions.Word}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(textToSpeech.Service.Options.URL, pathSegments, pathParameters)
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
// **Note:** This method is currently a beta release.
//
// **See also:** [Querying a single word from a custom
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customWords#cuWordQueryModel).
func (textToSpeech *TextToSpeechV1) GetWord(getWordOptions *GetWordOptions) (result *Translation, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getWordOptions, "getWordOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getWordOptions, "getWordOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/customizations", "words"}
	pathParameters := []string{*getWordOptions.CustomizationID, *getWordOptions.Word}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(textToSpeech.Service.Options.URL, pathSegments, pathParameters)
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

	response, err = textToSpeech.Service.Request(request, new(Translation))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*Translation)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// DeleteWord : Delete a custom word
// Deletes a single word from the specified custom voice model. You must use credentials for the instance of the service
// that owns a model to delete its words.
//
// **Note:** This method is currently a beta release.
//
// **See also:** [Deleting a word from a custom
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customWords#cuWordDelete).
func (textToSpeech *TextToSpeechV1) DeleteWord(deleteWordOptions *DeleteWordOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteWordOptions, "deleteWordOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteWordOptions, "deleteWordOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/customizations", "words"}
	pathParameters := []string{*deleteWordOptions.CustomizationID, *deleteWordOptions.Word}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(textToSpeech.Service.Options.URL, pathSegments, pathParameters)
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
// associate the customer ID with the data.
//
// You associate a customer ID with data by passing the `X-Watson-Metadata` header with a request that passes the data.
//
// **See also:** [Information
// security](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-information-security#information-security).
func (textToSpeech *TextToSpeechV1) DeleteUserData(deleteUserDataOptions *DeleteUserDataOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteUserDataOptions, "deleteUserDataOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteUserDataOptions, "deleteUserDataOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/user_data"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(textToSpeech.Service.Options.URL, pathSegments, pathParameters)
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

	// The customization ID (GUID) of the custom voice model. You must make the request with credentials for the instance
	// of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// The word that is to be added or updated for the custom voice model.
	Word *string `json:"word" validate:"required"`

	// The phonetic or sounds-like translation for the word. A phonetic translation is based on the SSML format for
	// representing the phonetic string of a word either as an IPA translation or as an IBM SPR translation. The Arabic,
	// Chinese, Dutch, and Korean languages support only IPA. A sounds-like is one or more words that, when combined, sound
	// like the word.
	Translation *string `json:"translation" validate:"required"`

	// **Japanese only.** The part of speech for the word. The service uses the value to produce the correct intonation for
	// the word. You can create only a single entry, with or without a single part of speech, for any word; you cannot
	// create multiple entries with different parts of speech for the same word. For more information, see [Working with
	// Japanese entries](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-rules#jaNotes).
	PartOfSpeech *string `json:"part_of_speech,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the AddWordOptions.PartOfSpeech property.
// **Japanese only.** The part of speech for the word. The service uses the value to produce the correct intonation for
// the word. You can create only a single entry, with or without a single part of speech, for any word; you cannot
// create multiple entries with different parts of speech for the same word. For more information, see [Working with
// Japanese entries](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-rules#jaNotes).
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

// AddWordsOptions : The AddWords options.
type AddWordsOptions struct {

	// The customization ID (GUID) of the custom voice model. You must make the request with credentials for the instance
	// of the service that owns the custom model.
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

// CreateVoiceModelOptions : The CreateVoiceModel options.
type CreateVoiceModelOptions struct {

	// The name of the new custom voice model.
	Name *string `json:"name" validate:"required"`

	// The language of the new custom voice model. You create a custom voice model for a specific language, not for a
	// specific voice. A custom model can be used with any voice, standard or neural, for its specified language. Omit the
	// parameter to use the the default language, `en-US`.
	Language *string `json:"language,omitempty"`

	// A description of the new custom voice model. Specifying a description is recommended.
	Description *string `json:"description,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the CreateVoiceModelOptions.Language property.
// The language of the new custom voice model. You create a custom voice model for a specific language, not for a
// specific voice. A custom model can be used with any voice, standard or neural, for its specified language. Omit the
// parameter to use the the default language, `en-US`.
const (
	CreateVoiceModelOptions_Language_ArAr = "ar-AR"
	CreateVoiceModelOptions_Language_DeDe = "de-DE"
	CreateVoiceModelOptions_Language_EnGb = "en-GB"
	CreateVoiceModelOptions_Language_EnUs = "en-US"
	CreateVoiceModelOptions_Language_EsEs = "es-ES"
	CreateVoiceModelOptions_Language_EsLa = "es-LA"
	CreateVoiceModelOptions_Language_EsUs = "es-US"
	CreateVoiceModelOptions_Language_FrFr = "fr-FR"
	CreateVoiceModelOptions_Language_ItIt = "it-IT"
	CreateVoiceModelOptions_Language_JaJp = "ja-JP"
	CreateVoiceModelOptions_Language_KoKr = "ko-KR"
	CreateVoiceModelOptions_Language_NlNl = "nl-NL"
	CreateVoiceModelOptions_Language_PtBr = "pt-BR"
	CreateVoiceModelOptions_Language_ZhCn = "zh-CN"
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

// DeleteUserDataOptions : The DeleteUserData options.
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

// DeleteVoiceModelOptions : The DeleteVoiceModel options.
type DeleteVoiceModelOptions struct {

	// The customization ID (GUID) of the custom voice model. You must make the request with credentials for the instance
	// of the service that owns the custom model.
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

// DeleteWordOptions : The DeleteWord options.
type DeleteWordOptions struct {

	// The customization ID (GUID) of the custom voice model. You must make the request with credentials for the instance
	// of the service that owns the custom model.
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

// GetPronunciationOptions : The GetPronunciation options.
type GetPronunciationOptions struct {

	// The word for which the pronunciation is requested.
	Text *string `json:"text" validate:"required"`

	// A voice that specifies the language in which the pronunciation is to be returned. All voices for the same language
	// (for example, `en-US`) return the same translation.
	Voice *string `json:"voice,omitempty"`

	// The phoneme format in which to return the pronunciation. The Arabic, Chinese, Dutch, and Korean languages support
	// only IPA. Omit the parameter to obtain the pronunciation in the default format.
	Format *string `json:"format,omitempty"`

	// The customization ID (GUID) of a custom voice model for which the pronunciation is to be returned. The language of a
	// specified custom model must match the language of the specified voice. If the word is not defined in the specified
	// custom model, the service returns the default translation for the custom model's language. You must make the request
	// with credentials for the instance of the service that owns the custom model. Omit the parameter to see the
	// translation for the specified voice with no customization.
	CustomizationID *string `json:"customization_id,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the GetPronunciationOptions.Voice property.
// A voice that specifies the language in which the pronunciation is to be returned. All voices for the same language
// (for example, `en-US`) return the same translation.
const (
	GetPronunciationOptions_Voice_ArArOmarvoice        = "ar-AR_OmarVoice"
	GetPronunciationOptions_Voice_DeDeBirgitv3voice    = "de-DE_BirgitV3Voice"
	GetPronunciationOptions_Voice_DeDeBirgitvoice      = "de-DE_BirgitVoice"
	GetPronunciationOptions_Voice_DeDeDieterv3voice    = "de-DE_DieterV3Voice"
	GetPronunciationOptions_Voice_DeDeDietervoice      = "de-DE_DieterVoice"
	GetPronunciationOptions_Voice_DeDeErikav3voice     = "de-DE_ErikaV3Voice"
	GetPronunciationOptions_Voice_EnGbKatev3voice      = "en-GB_KateV3Voice"
	GetPronunciationOptions_Voice_EnGbKatevoice        = "en-GB_KateVoice"
	GetPronunciationOptions_Voice_EnUsAllisonv3voice   = "en-US_AllisonV3Voice"
	GetPronunciationOptions_Voice_EnUsAllisonvoice     = "en-US_AllisonVoice"
	GetPronunciationOptions_Voice_EnUsEmilyv3voice     = "en-US_EmilyV3Voice"
	GetPronunciationOptions_Voice_EnUsHenryv3voice     = "en-US_HenryV3Voice"
	GetPronunciationOptions_Voice_EnUsKevinv3voice     = "en-US_KevinV3Voice"
	GetPronunciationOptions_Voice_EnUsLisav3voice      = "en-US_LisaV3Voice"
	GetPronunciationOptions_Voice_EnUsLisavoice        = "en-US_LisaVoice"
	GetPronunciationOptions_Voice_EnUsMichaelv3voice   = "en-US_MichaelV3Voice"
	GetPronunciationOptions_Voice_EnUsMichaelvoice     = "en-US_MichaelVoice"
	GetPronunciationOptions_Voice_EnUsOliviav3voice    = "en-US_OliviaV3Voice"
	GetPronunciationOptions_Voice_EsEsEnriquev3voice   = "es-ES_EnriqueV3Voice"
	GetPronunciationOptions_Voice_EsEsEnriquevoice     = "es-ES_EnriqueVoice"
	GetPronunciationOptions_Voice_EsEsLaurav3voice     = "es-ES_LauraV3Voice"
	GetPronunciationOptions_Voice_EsEsLauravoice       = "es-ES_LauraVoice"
	GetPronunciationOptions_Voice_EsLaSofiav3voice     = "es-LA_SofiaV3Voice"
	GetPronunciationOptions_Voice_EsLaSofiavoice       = "es-LA_SofiaVoice"
	GetPronunciationOptions_Voice_EsUsSofiav3voice     = "es-US_SofiaV3Voice"
	GetPronunciationOptions_Voice_EsUsSofiavoice       = "es-US_SofiaVoice"
	GetPronunciationOptions_Voice_FrFrReneev3voice     = "fr-FR_ReneeV3Voice"
	GetPronunciationOptions_Voice_FrFrReneevoice       = "fr-FR_ReneeVoice"
	GetPronunciationOptions_Voice_ItItFrancescav3voice = "it-IT_FrancescaV3Voice"
	GetPronunciationOptions_Voice_ItItFrancescavoice   = "it-IT_FrancescaVoice"
	GetPronunciationOptions_Voice_JaJpEmiv3voice       = "ja-JP_EmiV3Voice"
	GetPronunciationOptions_Voice_JaJpEmivoice         = "ja-JP_EmiVoice"
	GetPronunciationOptions_Voice_KoKrYoungmivoice     = "ko-KR_YoungmiVoice"
	GetPronunciationOptions_Voice_KoKrYunavoice        = "ko-KR_YunaVoice"
	GetPronunciationOptions_Voice_NlNlEmmavoice        = "nl-NL_EmmaVoice"
	GetPronunciationOptions_Voice_NlNlLiamvoice        = "nl-NL_LiamVoice"
	GetPronunciationOptions_Voice_PtBrIsabelav3voice   = "pt-BR_IsabelaV3Voice"
	GetPronunciationOptions_Voice_PtBrIsabelavoice     = "pt-BR_IsabelaVoice"
	GetPronunciationOptions_Voice_ZhCnLinavoice        = "zh-CN_LiNaVoice"
	GetPronunciationOptions_Voice_ZhCnWangweivoice     = "zh-CN_WangWeiVoice"
	GetPronunciationOptions_Voice_ZhCnZhangjingvoice   = "zh-CN_ZhangJingVoice"
)

// Constants associated with the GetPronunciationOptions.Format property.
// The phoneme format in which to return the pronunciation. The Arabic, Chinese, Dutch, and Korean languages support
// only IPA. Omit the parameter to obtain the pronunciation in the default format.
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

// GetVoiceModelOptions : The GetVoiceModel options.
type GetVoiceModelOptions struct {

	// The customization ID (GUID) of the custom voice model. You must make the request with credentials for the instance
	// of the service that owns the custom model.
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

// GetVoiceOptions : The GetVoice options.
type GetVoiceOptions struct {

	// The voice for which information is to be returned.
	Voice *string `json:"voice" validate:"required"`

	// The customization ID (GUID) of a custom voice model for which information is to be returned. You must make the
	// request with credentials for the instance of the service that owns the custom model. Omit the parameter to see
	// information about the specified voice with no customization.
	CustomizationID *string `json:"customization_id,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the GetVoiceOptions.Voice property.
// The voice for which information is to be returned.
const (
	GetVoiceOptions_Voice_ArArOmarvoice        = "ar-AR_OmarVoice"
	GetVoiceOptions_Voice_DeDeBirgitv3voice    = "de-DE_BirgitV3Voice"
	GetVoiceOptions_Voice_DeDeBirgitvoice      = "de-DE_BirgitVoice"
	GetVoiceOptions_Voice_DeDeDieterv3voice    = "de-DE_DieterV3Voice"
	GetVoiceOptions_Voice_DeDeDietervoice      = "de-DE_DieterVoice"
	GetVoiceOptions_Voice_DeDeErikav3voice     = "de-DE_ErikaV3Voice"
	GetVoiceOptions_Voice_EnGbKatev3voice      = "en-GB_KateV3Voice"
	GetVoiceOptions_Voice_EnGbKatevoice        = "en-GB_KateVoice"
	GetVoiceOptions_Voice_EnUsAllisonv3voice   = "en-US_AllisonV3Voice"
	GetVoiceOptions_Voice_EnUsAllisonvoice     = "en-US_AllisonVoice"
	GetVoiceOptions_Voice_EnUsEmilyv3voice     = "en-US_EmilyV3Voice"
	GetVoiceOptions_Voice_EnUsHenryv3voice     = "en-US_HenryV3Voice"
	GetVoiceOptions_Voice_EnUsKevinv3voice     = "en-US_KevinV3Voice"
	GetVoiceOptions_Voice_EnUsLisav3voice      = "en-US_LisaV3Voice"
	GetVoiceOptions_Voice_EnUsLisavoice        = "en-US_LisaVoice"
	GetVoiceOptions_Voice_EnUsMichaelv3voice   = "en-US_MichaelV3Voice"
	GetVoiceOptions_Voice_EnUsMichaelvoice     = "en-US_MichaelVoice"
	GetVoiceOptions_Voice_EnUsOliviav3voice    = "en-US_OliviaV3Voice"
	GetVoiceOptions_Voice_EsEsEnriquev3voice   = "es-ES_EnriqueV3Voice"
	GetVoiceOptions_Voice_EsEsEnriquevoice     = "es-ES_EnriqueVoice"
	GetVoiceOptions_Voice_EsEsLaurav3voice     = "es-ES_LauraV3Voice"
	GetVoiceOptions_Voice_EsEsLauravoice       = "es-ES_LauraVoice"
	GetVoiceOptions_Voice_EsLaSofiav3voice     = "es-LA_SofiaV3Voice"
	GetVoiceOptions_Voice_EsLaSofiavoice       = "es-LA_SofiaVoice"
	GetVoiceOptions_Voice_EsUsSofiav3voice     = "es-US_SofiaV3Voice"
	GetVoiceOptions_Voice_EsUsSofiavoice       = "es-US_SofiaVoice"
	GetVoiceOptions_Voice_FrFrReneev3voice     = "fr-FR_ReneeV3Voice"
	GetVoiceOptions_Voice_FrFrReneevoice       = "fr-FR_ReneeVoice"
	GetVoiceOptions_Voice_ItItFrancescav3voice = "it-IT_FrancescaV3Voice"
	GetVoiceOptions_Voice_ItItFrancescavoice   = "it-IT_FrancescaVoice"
	GetVoiceOptions_Voice_JaJpEmiv3voice       = "ja-JP_EmiV3Voice"
	GetVoiceOptions_Voice_JaJpEmivoice         = "ja-JP_EmiVoice"
	GetVoiceOptions_Voice_KoKrYoungmivoice     = "ko-KR_YoungmiVoice"
	GetVoiceOptions_Voice_KoKrYunavoice        = "ko-KR_YunaVoice"
	GetVoiceOptions_Voice_NlNlEmmavoice        = "nl-NL_EmmaVoice"
	GetVoiceOptions_Voice_NlNlLiamvoice        = "nl-NL_LiamVoice"
	GetVoiceOptions_Voice_PtBrIsabelav3voice   = "pt-BR_IsabelaV3Voice"
	GetVoiceOptions_Voice_PtBrIsabelavoice     = "pt-BR_IsabelaVoice"
	GetVoiceOptions_Voice_ZhCnLinavoice        = "zh-CN_LiNaVoice"
	GetVoiceOptions_Voice_ZhCnWangweivoice     = "zh-CN_WangWeiVoice"
	GetVoiceOptions_Voice_ZhCnZhangjingvoice   = "zh-CN_ZhangJingVoice"
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

// GetWordOptions : The GetWord options.
type GetWordOptions struct {

	// The customization ID (GUID) of the custom voice model. You must make the request with credentials for the instance
	// of the service that owns the custom model.
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

// ListVoiceModelsOptions : The ListVoiceModels options.
type ListVoiceModelsOptions struct {

	// The language for which custom voice models that are owned by the requesting credentials are to be returned. Omit the
	// parameter to see all custom voice models that are owned by the requester.
	Language *string `json:"language,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the ListVoiceModelsOptions.Language property.
// The language for which custom voice models that are owned by the requesting credentials are to be returned. Omit the
// parameter to see all custom voice models that are owned by the requester.
const (
	ListVoiceModelsOptions_Language_ArAr = "ar-AR"
	ListVoiceModelsOptions_Language_DeDe = "de-DE"
	ListVoiceModelsOptions_Language_EnGb = "en-GB"
	ListVoiceModelsOptions_Language_EnUs = "en-US"
	ListVoiceModelsOptions_Language_EsEs = "es-ES"
	ListVoiceModelsOptions_Language_EsLa = "es-LA"
	ListVoiceModelsOptions_Language_EsUs = "es-US"
	ListVoiceModelsOptions_Language_FrFr = "fr-FR"
	ListVoiceModelsOptions_Language_ItIt = "it-IT"
	ListVoiceModelsOptions_Language_JaJp = "ja-JP"
	ListVoiceModelsOptions_Language_KoKr = "ko-KR"
	ListVoiceModelsOptions_Language_NlNl = "nl-NL"
	ListVoiceModelsOptions_Language_PtBr = "pt-BR"
	ListVoiceModelsOptions_Language_ZhCn = "zh-CN"
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

// ListVoicesOptions : The ListVoices options.
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

// ListWordsOptions : The ListWords options.
type ListWordsOptions struct {

	// The customization ID (GUID) of the custom voice model. You must make the request with credentials for the instance
	// of the service that owns the custom model.
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

// Pronunciation : The pronunciation of the specified text.
type Pronunciation struct {

	// The pronunciation of the specified text in the requested voice and format. If a custom voice model is specified, the
	// pronunciation also reflects that custom voice.
	Pronunciation *string `json:"pronunciation" validate:"required"`
}

// SupportedFeatures : Additional service features that are supported with the voice.
type SupportedFeatures struct {

	// If `true`, the voice can be customized; if `false`, the voice cannot be customized. (Same as `customizable`.).
	CustomPronunciation *bool `json:"custom_pronunciation" validate:"required"`

	// If `true`, the voice can be transformed by using the SSML &lt;voice-transformation&gt; element; if `false`, the
	// voice cannot be transformed.
	VoiceTransformation *bool `json:"voice_transformation" validate:"required"`
}

// SynthesizeOptions : The Synthesize options.
type SynthesizeOptions struct {

	// The text to synthesize.
	Text *string `json:"text" validate:"required"`

	// The requested format (MIME type) of the audio. You can use the `Accept` header or the `accept` parameter to specify
	// the audio format. For more information about specifying an audio format, see **Audio formats (accept types)** in the
	// method description.
	Accept *string `json:"Accept,omitempty"`

	// The voice to use for synthesis.
	Voice *string `json:"voice,omitempty"`

	// The customization ID (GUID) of a custom voice model to use for the synthesis. If a custom voice model is specified,
	// it works only if it matches the language of the indicated voice. You must make the request with credentials for the
	// instance of the service that owns the custom model. Omit the parameter to use the specified voice with no
	// customization.
	CustomizationID *string `json:"customization_id,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the SynthesizeOptions.Voice property.
// The voice to use for synthesis.
const (
	SynthesizeOptions_Voice_ArArOmarvoice        = "ar-AR_OmarVoice"
	SynthesizeOptions_Voice_DeDeBirgitv3voice    = "de-DE_BirgitV3Voice"
	SynthesizeOptions_Voice_DeDeBirgitvoice      = "de-DE_BirgitVoice"
	SynthesizeOptions_Voice_DeDeDieterv3voice    = "de-DE_DieterV3Voice"
	SynthesizeOptions_Voice_DeDeDietervoice      = "de-DE_DieterVoice"
	SynthesizeOptions_Voice_DeDeErikav3voice     = "de-DE_ErikaV3Voice"
	SynthesizeOptions_Voice_EnGbKatev3voice      = "en-GB_KateV3Voice"
	SynthesizeOptions_Voice_EnGbKatevoice        = "en-GB_KateVoice"
	SynthesizeOptions_Voice_EnUsAllisonv3voice   = "en-US_AllisonV3Voice"
	SynthesizeOptions_Voice_EnUsAllisonvoice     = "en-US_AllisonVoice"
	SynthesizeOptions_Voice_EnUsEmilyv3voice     = "en-US_EmilyV3Voice"
	SynthesizeOptions_Voice_EnUsHenryv3voice     = "en-US_HenryV3Voice"
	SynthesizeOptions_Voice_EnUsKevinv3voice     = "en-US_KevinV3Voice"
	SynthesizeOptions_Voice_EnUsLisav3voice      = "en-US_LisaV3Voice"
	SynthesizeOptions_Voice_EnUsLisavoice        = "en-US_LisaVoice"
	SynthesizeOptions_Voice_EnUsMichaelv3voice   = "en-US_MichaelV3Voice"
	SynthesizeOptions_Voice_EnUsMichaelvoice     = "en-US_MichaelVoice"
	SynthesizeOptions_Voice_EnUsOliviav3voice    = "en-US_OliviaV3Voice"
	SynthesizeOptions_Voice_EsEsEnriquev3voice   = "es-ES_EnriqueV3Voice"
	SynthesizeOptions_Voice_EsEsEnriquevoice     = "es-ES_EnriqueVoice"
	SynthesizeOptions_Voice_EsEsLaurav3voice     = "es-ES_LauraV3Voice"
	SynthesizeOptions_Voice_EsEsLauravoice       = "es-ES_LauraVoice"
	SynthesizeOptions_Voice_EsLaSofiav3voice     = "es-LA_SofiaV3Voice"
	SynthesizeOptions_Voice_EsLaSofiavoice       = "es-LA_SofiaVoice"
	SynthesizeOptions_Voice_EsUsSofiav3voice     = "es-US_SofiaV3Voice"
	SynthesizeOptions_Voice_EsUsSofiavoice       = "es-US_SofiaVoice"
	SynthesizeOptions_Voice_FrFrReneev3voice     = "fr-FR_ReneeV3Voice"
	SynthesizeOptions_Voice_FrFrReneevoice       = "fr-FR_ReneeVoice"
	SynthesizeOptions_Voice_ItItFrancescav3voice = "it-IT_FrancescaV3Voice"
	SynthesizeOptions_Voice_ItItFrancescavoice   = "it-IT_FrancescaVoice"
	SynthesizeOptions_Voice_JaJpEmiv3voice       = "ja-JP_EmiV3Voice"
	SynthesizeOptions_Voice_JaJpEmivoice         = "ja-JP_EmiVoice"
	SynthesizeOptions_Voice_KoKrYoungmivoice     = "ko-KR_YoungmiVoice"
	SynthesizeOptions_Voice_KoKrYunavoice        = "ko-KR_YunaVoice"
	SynthesizeOptions_Voice_NlNlEmmavoice        = "nl-NL_EmmaVoice"
	SynthesizeOptions_Voice_NlNlLiamvoice        = "nl-NL_LiamVoice"
	SynthesizeOptions_Voice_PtBrIsabelav3voice   = "pt-BR_IsabelaV3Voice"
	SynthesizeOptions_Voice_PtBrIsabelavoice     = "pt-BR_IsabelaVoice"
	SynthesizeOptions_Voice_ZhCnLinavoice        = "zh-CN_LiNaVoice"
	SynthesizeOptions_Voice_ZhCnWangweivoice     = "zh-CN_WangWeiVoice"
	SynthesizeOptions_Voice_ZhCnZhangjingvoice   = "zh-CN_ZhangJingVoice"
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
	// Chinese, Dutch, and Korean languages support only IPA. A sounds-like is one or more words that, when combined, sound
	// like the word.
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

// NewTranslation : Instantiate Translation (Generic Model Constructor)
func (textToSpeech *TextToSpeechV1) NewTranslation(translation string) (model *Translation, err error) {
	model = &Translation{
		Translation: core.StringPtr(translation),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UpdateVoiceModelOptions : The UpdateVoiceModel options.
type UpdateVoiceModelOptions struct {

	// The customization ID (GUID) of the custom voice model. You must make the request with credentials for the instance
	// of the service that owns the custom model.
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

// Voice : Information about an available voice model.
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

	// Returns information about a specified custom voice model. This field is returned only by the **Get a voice** method
	// and only when you specify the customization ID of a custom voice model.
	Customization *VoiceModel `json:"customization,omitempty"`
}

// VoiceModel : Information about an existing custom voice model.
type VoiceModel struct {

	// The customization ID (GUID) of the custom voice model. The **Create a custom model** method returns only this field.
	// It does not not return the other fields of this object.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// The name of the custom voice model.
	Name *string `json:"name,omitempty"`

	// The language identifier of the custom voice model (for example, `en-US`).
	Language *string `json:"language,omitempty"`

	// The GUID of the credentials for the instance of the service that owns the custom voice model.
	Owner *string `json:"owner,omitempty"`

	// The date and time in Coordinated Universal Time (UTC) at which the custom voice model was created. The value is
	// provided in full ISO 8601 format (`YYYY-MM-DDThh:mm:ss.sTZD`).
	Created *string `json:"created,omitempty"`

	// The date and time in Coordinated Universal Time (UTC) at which the custom voice model was last modified. The
	// `created` and `updated` fields are equal when a voice model is first added but has yet to be updated. The value is
	// provided in full ISO 8601 format (`YYYY-MM-DDThh:mm:ss.sTZD`).
	LastModified *string `json:"last_modified,omitempty"`

	// The description of the custom voice model.
	Description *string `json:"description,omitempty"`

	// An array of `Word` objects that lists the words and their translations from the custom voice model. The words are
	// listed in alphabetical order, with uppercase letters listed before lowercase letters. The array is empty if the
	// custom model contains no words. This field is returned only by the **Get a voice** method and only when you specify
	// the customization ID of a custom voice model.
	Words []Word `json:"words,omitempty"`
}

// VoiceModels : Information about existing custom voice models.
type VoiceModels struct {

	// An array of `VoiceModel` objects that provides information about each available custom voice model. The array is
	// empty if the requesting credentials own no custom voice models (if no language is specified) or own no custom voice
	// models for the specified language.
	Customizations []VoiceModel `json:"customizations" validate:"required"`
}

// Voices : Information about all available voice models.
type Voices struct {

	// A list of available voices.
	Voices []Voice `json:"voices" validate:"required"`
}

// Word : Information about a word for the custom voice model.
type Word struct {

	// The word for the custom voice model. The maximum length of a word is 49 characters.
	Word *string `json:"word" validate:"required"`

	// The phonetic or sounds-like translation for the word. A phonetic translation is based on the SSML format for
	// representing the phonetic string of a word either as an IPA or IBM SPR translation. The Arabic, Chinese, Dutch, and
	// Korean languages support only IPA. A sounds-like translation consists of one or more words that, when combined,
	// sound like the word. The maximum length of a translation is 499 characters.
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

// NewWord : Instantiate Word (Generic Model Constructor)
func (textToSpeech *TextToSpeechV1) NewWord(word string, translation string) (model *Word, err error) {
	model = &Word{
		Word:        core.StringPtr(word),
		Translation: core.StringPtr(translation),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// Words : For the **Add custom words** method, one or more words that are to be added or updated for the custom voice model and
// the translation for each specified word.
//
// For the **List custom words** method, the words and their translations from the custom voice model.
type Words struct {

	// The **Add custom words** method accepts an array of `Word` objects. Each object provides a word that is to be added
	// or updated for the custom voice model and the word's translation.
	//
	// The **List custom words** method returns an array of `Word` objects. Each object shows a word and its translation
	// from the custom voice model. The words are listed in alphabetical order, with uppercase letters listed before
	// lowercase letters. The array is empty if the custom model contains no words.
	Words []Word `json:"words" validate:"required"`
}

// NewWords : Instantiate Words (Generic Model Constructor)
func (textToSpeech *TextToSpeechV1) NewWords(words []Word) (model *Words, err error) {
	model = &Words{
		Words: words,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}
