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
	core "github.com/ibm-watson/go-sdk/core"
	"io"
)

// TextToSpeechV1: ### Service Overview
// The IBM&reg; Text to Speech service provides an API that uses IBM's speech-synthesis capabilities to synthesize text
// into natural-sounding speech in a variety of languages, dialects, and voices. The service supports at least one male
// or female voice, sometimes both, for each language. The audio is streamed back to the client with minimal delay. For
// more information about the service, see the [IBM&reg; Cloud
// documentation](https://console.bluemix.net/docs/services/text-to-speech/index.html).
//
// ### API usage guidelines
// * **Audio formats:** The service can produce audio in many formats (MIME types). See [Specifying an audio
// format](https://console.bluemix.net/docs/services/text-to-speech/http.html#format).
// * **SSML:** Many methods refer to the Speech Synthesis Markup Language (SSML). SSML is an XML-based markup language
// that provides text annotation for speech-synthesis applications. See [Using
// SSML](https://console.bluemix.net/docs/services/text-to-speech/SSML.html) and [Using IBM
// SPR](https://console.bluemix.net/docs/services/text-to-speech/SPRs.html).
// * **Word translations:** Many customization methods accept sounds-like or phonetic translations for words. Phonetic
// translations are based on the SSML phoneme format for representing a word. You can specify them in standard
// International Phonetic Alphabet (IPA) representation
//
//   &lt;phoneme alphabet="ipa" ph="t&#601;m&#712;&#593;to"&gt;&lt;/phoneme&gt;
//
//   or in the proprietary IBM Symbolic Phonetic Representation (SPR)
//
//   &lt;phoneme alphabet="ibm" ph="1gAstroEntxrYFXs"&gt;&lt;/phoneme&gt;
//
//   See [Understanding customization](https://console.bluemix.net/docs/services/text-to-speech/custom-intro.html).
// * **WebSocket interface:** The service also offers a WebSocket interface for speech synthesis. The WebSocket
// interface supports both plain text and SSML input, including the SSML &lt;mark&gt; element and word timings. See [The
// WebSocket interface](https://console.bluemix.net/docs/services/text-to-speech/websockets.html).
// * **Customization IDs:** Many methods accept a customization ID, which is a Globally Unique Identifier (GUID).
// Customization IDs are hexadecimal strings that have the format `xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx`.
// * **`X-Watson-Learning-Opt-Out`:** By default, all Watson services log requests and their results. Logging is done
// only to improve the services for future users. The logged data is not shared or made public. To prevent IBM from
// accessing your data for general service improvements, set the `X-Watson-Learning-Opt-Out` request header to `true`
// for all requests. You must set the header on each request that you do not want IBM to access for general service
// improvements.
//
//   Methods of the customization interface do not log words and translations that you use to build custom voice models.
// Your training data is never used to improve the service's base models. However, the service does log such data when a
// custom model is used with a synthesize request. You must set the `X-Watson-Learning-Opt-Out` request header to `true`
// to prevent IBM from accessing the data to improve the service.
// * **`X-Watson-Metadata`:** This header allows you to associate a customer ID with data that is passed with a request.
// If necessary, you can use the **Delete labeled data** method to delete the data for a customer ID. See [Information
// security](https://console.bluemix.net/docs/services/text-to-speech/information-security.html).
//
// Version: V1
// See: http://www.ibm.com/watson/developercloud/text-to-speech.html
type TextToSpeechV1 struct {
	service *core.WatsonService
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
	service, serviceErr := core.NewWatsonService(serviceOptions, "text_to_speech")
	if serviceErr != nil {
		return nil, serviceErr
	}

	return &TextToSpeechV1{service: service}, nil
}

// GetVoice : Get a voice
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
	builder.ConstructHTTPURL(textToSpeech.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range getVoiceOptions.Headers {
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

	response, err := textToSpeech.service.Request(request, new(Voice))
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
func (textToSpeech *TextToSpeechV1) ListVoices(listVoicesOptions *ListVoicesOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateStruct(listVoicesOptions, "listVoicesOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/voices"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(textToSpeech.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range listVoicesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := textToSpeech.service.Request(request, new(Voices))
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
	builder.ConstructHTTPURL(textToSpeech.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range synthesizeOptions.Headers {
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

	response, err := textToSpeech.service.Request(request, new(io.ReadCloser))
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
	builder.ConstructHTTPURL(textToSpeech.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range getPronunciationOptions.Headers {
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

	response, err := textToSpeech.service.Request(request, new(Pronunciation))
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
	builder.ConstructHTTPURL(textToSpeech.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range createVoiceModelOptions.Headers {
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

	response, err := textToSpeech.service.Request(request, new(VoiceModel))
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
	builder.ConstructHTTPURL(textToSpeech.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteVoiceModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := textToSpeech.service.Request(request, nil)
	return response, err
}

// GetVoiceModel : Get a custom model
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
	builder.ConstructHTTPURL(textToSpeech.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range getVoiceModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := textToSpeech.service.Request(request, new(VoiceModel))
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
func (textToSpeech *TextToSpeechV1) ListVoiceModels(listVoiceModelsOptions *ListVoiceModelsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateStruct(listVoiceModelsOptions, "listVoiceModelsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/customizations"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(textToSpeech.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range listVoiceModelsOptions.Headers {
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

	response, err := textToSpeech.service.Request(request, new(VoiceModels))
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
	builder.ConstructHTTPURL(textToSpeech.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range updateVoiceModelOptions.Headers {
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

	response, err := textToSpeech.service.Request(request, nil)
	return response, err
}

// AddWord : Add a custom word
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
	builder.ConstructHTTPURL(textToSpeech.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range addWordOptions.Headers {
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
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := textToSpeech.service.Request(request, nil)
	return response, err
}

// AddWords : Add custom words
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
	builder.ConstructHTTPURL(textToSpeech.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range addWordsOptions.Headers {
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

	response, err := textToSpeech.service.Request(request, nil)
	return response, err
}

// DeleteWord : Delete a custom word
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
	builder.ConstructHTTPURL(textToSpeech.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteWordOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := textToSpeech.service.Request(request, nil)
	return response, err
}

// GetWord : Get a custom word
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
	builder.ConstructHTTPURL(textToSpeech.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range getWordOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := textToSpeech.service.Request(request, new(Translation))
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
	builder.ConstructHTTPURL(textToSpeech.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range listWordsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := textToSpeech.service.Request(request, new(Words))
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
	builder.ConstructHTTPURL(textToSpeech.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteUserDataOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("customer_id", fmt.Sprint(*deleteUserDataOptions.CustomerID))

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := textToSpeech.service.Request(request, nil)
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
	Translation *string `json:"translation,omitempty"`

	// **Japanese only.** The part of speech for the word. The service uses the value to produce the correct intonation for
	// the word. You can create only a single entry, with or without a single part of speech, for any word; you cannot
	// create multiple entries with different parts of speech for the same word. For more information, see [Working with
	// Japanese entries](https://console.bluemix.net/docs/services/text-to-speech/custom-rules.html#jaNotes).
	PartOfSpeech *string `json:"part_of_speech,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewAddWordOptions : Instantiate AddWordOptions
func (textToSpeech *TextToSpeechV1) NewAddWordOptions(customizationID string, word string) *AddWordOptions {
	return &AddWordOptions{
		CustomizationID: core.StringPtr(customizationID),
		Word: core.StringPtr(word),
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
	Words []Word `json:"words,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewAddWordsOptions : Instantiate AddWordsOptions
func (textToSpeech *TextToSpeechV1) NewAddWordsOptions(customizationID string) *AddWordsOptions {
	return &AddWordsOptions{
		CustomizationID: core.StringPtr(customizationID),
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
		Word: core.StringPtr(word),
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
		Word: core.StringPtr(word),
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

// SupportedFeatures : SupportedFeatures struct
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

	// The requested audio format (MIME type) of the audio. You can use the `Accept` header or the `accept` query parameter
	// to specify the audio format. (For the `audio/l16` format, you can optionally specify `endianness=big-endian` or
	// `endianness=little-endian`; the default is little endian.) For detailed information about the supported audio
	// formats and sampling rates, see [Specifying an audio
	// format](https://console.bluemix.net/docs/services/text-to-speech/http.html#format).
	Accept *string `json:"Accept,omitempty"`

	// The voice to use for synthesis.
	Voice *string `json:"voice,omitempty"`

	// The customization ID (GUID) of a custom voice model to use for the synthesis. If a custom voice model is specified,
	// it is guaranteed to work only if it matches the language of the indicated voice. You must make the request with
	// service credentials created for the instance of the service that owns the custom model. Omit the parameter to use
	// the specified voice with no customization.
	CustomizationID *string `json:"customization_id,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

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

// Translation : Translation struct
type Translation struct {

	// The phonetic or sounds-like translation for the word. A phonetic translation is based on the SSML format for
	// representing the phonetic string of a word either as an IPA translation or as an IBM SPR translation. A sounds-like
	// is one or more words that, when combined, sound like the word.
	Translation *string `json:"translation,omitempty"`

	// **Japanese only.** The part of speech for the word. The service uses the value to produce the correct intonation for
	// the word. You can create only a single entry, with or without a single part of speech, for any word; you cannot
	// create multiple entries with different parts of speech for the same word. For more information, see [Working with
	// Japanese entries](https://console.bluemix.net/docs/services/text-to-speech/custom-rules.html#jaNotes).
	PartOfSpeech *string `json:"part_of_speech,omitempty"`
}

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

	// Describes the additional service features supported with the voice.
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
	// Japanese entries](https://console.bluemix.net/docs/services/text-to-speech/custom-rules.html#jaNotes).
	PartOfSpeech *string `json:"part_of_speech,omitempty"`
}

// Words : Words struct
type Words struct {

	// The **Add custom words** method accepts an array of `Word` objects. Each object provides a word that is to be added
	// or updated for the custom voice model and the word's translation.
	//
	// The **List custom words** method returns an array of `Word` objects. Each object shows a word and its translation
	// from the custom voice model. The words are listed in alphabetical order, with uppercase letters listed before
	// lowercase letters. The array is empty if the custom model contains no words.
	Words []Word `json:"words,omitempty"`
}
