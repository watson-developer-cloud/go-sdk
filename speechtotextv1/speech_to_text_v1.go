// Package speechtotextv1 : Operations and models for the SpeechToTextV1 service
package speechtotextv1

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
	"io"
	"os"
	"strings"

	core "github.com/ibm-watson/go-sdk/core"
)

// SpeechToTextV1 : The IBM&reg; Speech to Text service provides an API that uses IBM's speech-recognition capabilities to produce
// transcripts of spoken audio. The service can transcribe speech from various languages and audio formats. It addition
// to basic transcription, the service can produce detailed information about many different aspects of the audio. For
// most languages, the service supports two sampling rates, broadband and narrowband. It returns all JSON response
// content in the UTF-8 character set.
//
//  For more information about the service, see the [IBM&reg; Cloud
// documentation](https://console.bluemix.net/docs/services/speech-to-text/index.html).
//
// ### API usage guidelines
// * **Audio formats:** The service accepts audio in many formats (MIME types). See [Audio
// formats](https://console.bluemix.net/docs/services/speech-to-text/audio-formats.html).
// * **HTTP interfaces:** The service provides two HTTP Representational State Transfer (REST) interfaces for speech
// recognition. The basic interface includes a single synchronous method. The asynchronous interface provides multiple
// methods that use registered callbacks and polling for non-blocking recognition. See [The HTTP
// interface](https://console.bluemix.net/docs/services/speech-to-text/http.html) and [The asynchronous HTTP
// interface](https://console.bluemix.net/docs/services/speech-to-text/async.html).
// * **WebSocket interface:** The service also offers a WebSocket interface for speech recognition. The WebSocket
// interface provides a full-duplex, low-latency communication channel. Clients send requests and audio to the service
// and receive results over a single connection in an asynchronous fashion. See [The WebSocket
// interface](https://console.bluemix.net/docs/services/speech-to-text/websockets.html).
// * **Customization:** The service offers two customization interfaces. Use language model customization to expand the
// vocabulary of a base model with domain-specific terminology. Use acoustic model customization to adapt a base model
// for the acoustic characteristics of your audio. Language model customization is generally available for production
// use by most supported languages; acoustic model customization is beta functionality that is available for all
// supported languages. See [The customization
// interface](https://console.bluemix.net/docs/services/speech-to-text/custom.html).
// * **Customization IDs:** Many methods accept a customization ID to identify a custom language or custom acoustic
// model. Customization IDs are Globally Unique Identifiers (GUIDs). They are hexadecimal strings that have the format
// `xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx`.
// * **`X-Watson-Learning-Opt-Out`:** By default, all Watson services log requests and their results. Logging is done
// only to improve the services for future users. The logged data is not shared or made public. To prevent IBM from
// accessing your data for general service improvements, set the `X-Watson-Learning-Opt-Out` request header to `true`
// for all requests. You must set the header on each request that you do not want IBM to access for general service
// improvements.
//
//   Methods of the customization interface do not log corpora, words, and audio resources that you use to build custom
// models. Your training data is never used to improve the service's base models. However, the service does log such
// data when a custom model is used with a recognition request. You must set the `X-Watson-Learning-Opt-Out` request
// header to `true` to prevent IBM from accessing the data to improve the service.
// * **`X-Watson-Metadata`**: This header allows you to associate a customer ID with data that is passed with a request.
// If necessary, you can use the **Delete labeled data** method to delete the data for a customer ID. See [Information
// security](https://console.bluemix.net/docs/services/speech-to-text/information-security.html).
//
// Version: V1
// See: http://www.ibm.com/watson/developercloud/speech-to-text.html
type SpeechToTextV1 struct {
	service *core.WatsonService
}

// SpeechToTextV1Options : Service options
type SpeechToTextV1Options struct {
	URL            string
	Username       string
	Password       string
	IAMApiKey      string
	IAMAccessToken string
	IAMURL         string
}

// NewSpeechToTextV1 : Instantiate SpeechToTextV1
func NewSpeechToTextV1(options *SpeechToTextV1Options) (*SpeechToTextV1, error) {
	if options.URL == "" {
		options.URL = "https://stream.watsonplatform.net/speech-to-text/api"
	}

	serviceOptions := &core.ServiceOptions{
		URL:            options.URL,
		Username:       options.Username,
		Password:       options.Password,
		IAMApiKey:      options.IAMApiKey,
		IAMAccessToken: options.IAMAccessToken,
		IAMURL:         options.IAMURL,
	}
	service, serviceErr := core.NewWatsonService(serviceOptions, "speech_to_text")
	if serviceErr != nil {
		return nil, serviceErr
	}

	return &SpeechToTextV1{service: service}, nil
}

// GetModel : Get a model
func (speechToText *SpeechToTextV1) GetModel(getModelOptions *GetModelOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(getModelOptions, "getModelOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(getModelOptions, "getModelOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/models"}
	pathParameters := []string{*getModelOptions.ModelID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(speechToText.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range getModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.service.Request(request, new(SpeechModel))
	return response, err
}

// GetGetModelResult : Retrieve result of GetModel operation
func (speechToText *SpeechToTextV1) GetGetModelResult(response *core.DetailedResponse) *SpeechModel {
	result, ok := response.Result.(*SpeechModel)
	if ok {
		return result
	}
	return nil
}

// ListModels : List models
func (speechToText *SpeechToTextV1) ListModels(listModelsOptions *ListModelsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateStruct(listModelsOptions, "listModelsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/models"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(speechToText.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range listModelsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.service.Request(request, new(SpeechModels))
	return response, err
}

// GetListModelsResult : Retrieve result of ListModels operation
func (speechToText *SpeechToTextV1) GetListModelsResult(response *core.DetailedResponse) *SpeechModels {
	result, ok := response.Result.(*SpeechModels)
	if ok {
		return result
	}
	return nil
}

// Recognize : Recognize audio
func (speechToText *SpeechToTextV1) Recognize(recognizeOptions *RecognizeOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(recognizeOptions, "recognizeOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(recognizeOptions, "recognizeOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/recognize"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(speechToText.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range recognizeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if recognizeOptions.ContentType != nil {
		builder.AddHeader("Content-Type", fmt.Sprint(*recognizeOptions.ContentType))
	}

	if recognizeOptions.Model != nil {
		builder.AddQuery("model", fmt.Sprint(*recognizeOptions.Model))
	}
	if recognizeOptions.CustomizationID != nil {
		builder.AddQuery("customization_id", fmt.Sprint(*recognizeOptions.CustomizationID))
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

	_, err := builder.SetBodyContent(core.StringNilMapper(recognizeOptions.ContentType), nil, nil, recognizeOptions.Audio)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.service.Request(request, new(SpeechRecognitionResults))
	return response, err
}

// GetRecognizeResult : Retrieve result of Recognize operation
func (speechToText *SpeechToTextV1) GetRecognizeResult(response *core.DetailedResponse) *SpeechRecognitionResults {
	result, ok := response.Result.(*SpeechRecognitionResults)
	if ok {
		return result
	}
	return nil
}

// CheckJob : Check a job
func (speechToText *SpeechToTextV1) CheckJob(checkJobOptions *CheckJobOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(checkJobOptions, "checkJobOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(checkJobOptions, "checkJobOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/recognitions"}
	pathParameters := []string{*checkJobOptions.ID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(speechToText.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range checkJobOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.service.Request(request, new(RecognitionJob))
	return response, err
}

// GetCheckJobResult : Retrieve result of CheckJob operation
func (speechToText *SpeechToTextV1) GetCheckJobResult(response *core.DetailedResponse) *RecognitionJob {
	result, ok := response.Result.(*RecognitionJob)
	if ok {
		return result
	}
	return nil
}

// CheckJobs : Check jobs
func (speechToText *SpeechToTextV1) CheckJobs(checkJobsOptions *CheckJobsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateStruct(checkJobsOptions, "checkJobsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/recognitions"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(speechToText.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range checkJobsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.service.Request(request, new(RecognitionJobs))
	return response, err
}

// GetCheckJobsResult : Retrieve result of CheckJobs operation
func (speechToText *SpeechToTextV1) GetCheckJobsResult(response *core.DetailedResponse) *RecognitionJobs {
	result, ok := response.Result.(*RecognitionJobs)
	if ok {
		return result
	}
	return nil
}

// CreateJob : Create a job
func (speechToText *SpeechToTextV1) CreateJob(createJobOptions *CreateJobOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(createJobOptions, "createJobOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(createJobOptions, "createJobOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/recognitions"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(speechToText.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range createJobOptions.Headers {
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
	if createJobOptions.CustomizationID != nil {
		builder.AddQuery("customization_id", fmt.Sprint(*createJobOptions.CustomizationID))
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

	_, err := builder.SetBodyContent(core.StringNilMapper(createJobOptions.ContentType), nil, nil, createJobOptions.Audio)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.service.Request(request, new(RecognitionJob))
	return response, err
}

// GetCreateJobResult : Retrieve result of CreateJob operation
func (speechToText *SpeechToTextV1) GetCreateJobResult(response *core.DetailedResponse) *RecognitionJob {
	result, ok := response.Result.(*RecognitionJob)
	if ok {
		return result
	}
	return nil
}

// DeleteJob : Delete a job
func (speechToText *SpeechToTextV1) DeleteJob(deleteJobOptions *DeleteJobOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(deleteJobOptions, "deleteJobOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(deleteJobOptions, "deleteJobOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/recognitions"}
	pathParameters := []string{*deleteJobOptions.ID}

	builder := core.NewRequestBuilder(core.DELETE)
	builder.ConstructHTTPURL(speechToText.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteJobOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.service.Request(request, nil)
	return response, err
}

// RegisterCallback : Register a callback
func (speechToText *SpeechToTextV1) RegisterCallback(registerCallbackOptions *RegisterCallbackOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(registerCallbackOptions, "registerCallbackOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(registerCallbackOptions, "registerCallbackOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/register_callback"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(speechToText.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range registerCallbackOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("callback_url", fmt.Sprint(*registerCallbackOptions.CallbackURL))
	if registerCallbackOptions.UserSecret != nil {
		builder.AddQuery("user_secret", fmt.Sprint(*registerCallbackOptions.UserSecret))
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.service.Request(request, new(RegisterStatus))
	return response, err
}

// GetRegisterCallbackResult : Retrieve result of RegisterCallback operation
func (speechToText *SpeechToTextV1) GetRegisterCallbackResult(response *core.DetailedResponse) *RegisterStatus {
	result, ok := response.Result.(*RegisterStatus)
	if ok {
		return result
	}
	return nil
}

// UnregisterCallback : Unregister a callback
func (speechToText *SpeechToTextV1) UnregisterCallback(unregisterCallbackOptions *UnregisterCallbackOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(unregisterCallbackOptions, "unregisterCallbackOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(unregisterCallbackOptions, "unregisterCallbackOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/unregister_callback"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(speechToText.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range unregisterCallbackOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("callback_url", fmt.Sprint(*unregisterCallbackOptions.CallbackURL))

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.service.Request(request, nil)
	return response, err
}

// CreateLanguageModel : Create a custom language model
func (speechToText *SpeechToTextV1) CreateLanguageModel(createLanguageModelOptions *CreateLanguageModelOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(createLanguageModelOptions, "createLanguageModelOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(createLanguageModelOptions, "createLanguageModelOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/customizations"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(speechToText.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range createLanguageModelOptions.Headers {
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
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.service.Request(request, new(LanguageModel))
	return response, err
}

// GetCreateLanguageModelResult : Retrieve result of CreateLanguageModel operation
func (speechToText *SpeechToTextV1) GetCreateLanguageModelResult(response *core.DetailedResponse) *LanguageModel {
	result, ok := response.Result.(*LanguageModel)
	if ok {
		return result
	}
	return nil
}

// DeleteLanguageModel : Delete a custom language model
func (speechToText *SpeechToTextV1) DeleteLanguageModel(deleteLanguageModelOptions *DeleteLanguageModelOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(deleteLanguageModelOptions, "deleteLanguageModelOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(deleteLanguageModelOptions, "deleteLanguageModelOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/customizations"}
	pathParameters := []string{*deleteLanguageModelOptions.CustomizationID}

	builder := core.NewRequestBuilder(core.DELETE)
	builder.ConstructHTTPURL(speechToText.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteLanguageModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.service.Request(request, nil)
	return response, err
}

// GetLanguageModel : Get a custom language model
func (speechToText *SpeechToTextV1) GetLanguageModel(getLanguageModelOptions *GetLanguageModelOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(getLanguageModelOptions, "getLanguageModelOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(getLanguageModelOptions, "getLanguageModelOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/customizations"}
	pathParameters := []string{*getLanguageModelOptions.CustomizationID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(speechToText.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range getLanguageModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.service.Request(request, new(LanguageModel))
	return response, err
}

// GetGetLanguageModelResult : Retrieve result of GetLanguageModel operation
func (speechToText *SpeechToTextV1) GetGetLanguageModelResult(response *core.DetailedResponse) *LanguageModel {
	result, ok := response.Result.(*LanguageModel)
	if ok {
		return result
	}
	return nil
}

// ListLanguageModels : List custom language models
func (speechToText *SpeechToTextV1) ListLanguageModels(listLanguageModelsOptions *ListLanguageModelsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateStruct(listLanguageModelsOptions, "listLanguageModelsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/customizations"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(speechToText.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range listLanguageModelsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listLanguageModelsOptions.Language != nil {
		builder.AddQuery("language", fmt.Sprint(*listLanguageModelsOptions.Language))
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.service.Request(request, new(LanguageModels))
	return response, err
}

// GetListLanguageModelsResult : Retrieve result of ListLanguageModels operation
func (speechToText *SpeechToTextV1) GetListLanguageModelsResult(response *core.DetailedResponse) *LanguageModels {
	result, ok := response.Result.(*LanguageModels)
	if ok {
		return result
	}
	return nil
}

// ResetLanguageModel : Reset a custom language model
func (speechToText *SpeechToTextV1) ResetLanguageModel(resetLanguageModelOptions *ResetLanguageModelOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(resetLanguageModelOptions, "resetLanguageModelOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(resetLanguageModelOptions, "resetLanguageModelOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/customizations", "reset"}
	pathParameters := []string{*resetLanguageModelOptions.CustomizationID}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(speechToText.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range resetLanguageModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.service.Request(request, nil)
	return response, err
}

// TrainLanguageModel : Train a custom language model
func (speechToText *SpeechToTextV1) TrainLanguageModel(trainLanguageModelOptions *TrainLanguageModelOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(trainLanguageModelOptions, "trainLanguageModelOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(trainLanguageModelOptions, "trainLanguageModelOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/customizations", "train"}
	pathParameters := []string{*trainLanguageModelOptions.CustomizationID}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(speechToText.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range trainLanguageModelOptions.Headers {
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
		return nil, err
	}

	response, err := speechToText.service.Request(request, nil)
	return response, err
}

// UpgradeLanguageModel : Upgrade a custom language model
func (speechToText *SpeechToTextV1) UpgradeLanguageModel(upgradeLanguageModelOptions *UpgradeLanguageModelOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(upgradeLanguageModelOptions, "upgradeLanguageModelOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(upgradeLanguageModelOptions, "upgradeLanguageModelOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/customizations", "upgrade_model"}
	pathParameters := []string{*upgradeLanguageModelOptions.CustomizationID}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(speechToText.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range upgradeLanguageModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.service.Request(request, nil)
	return response, err
}

// AddCorpus : Add a corpus
func (speechToText *SpeechToTextV1) AddCorpus(addCorpusOptions *AddCorpusOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(addCorpusOptions, "addCorpusOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(addCorpusOptions, "addCorpusOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/customizations", "corpora"}
	pathParameters := []string{*addCorpusOptions.CustomizationID, *addCorpusOptions.CorpusName}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(speechToText.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range addCorpusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if addCorpusOptions.AllowOverwrite != nil {
		builder.AddQuery("allow_overwrite", fmt.Sprint(*addCorpusOptions.AllowOverwrite))
	}

	builder.AddFormData("corpus_file", core.StringNilMapper(addCorpusOptions.CorpusFilename),
		"text/plain", addCorpusOptions.CorpusFile)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.service.Request(request, nil)
	return response, err
}

// DeleteCorpus : Delete a corpus
func (speechToText *SpeechToTextV1) DeleteCorpus(deleteCorpusOptions *DeleteCorpusOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(deleteCorpusOptions, "deleteCorpusOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(deleteCorpusOptions, "deleteCorpusOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/customizations", "corpora"}
	pathParameters := []string{*deleteCorpusOptions.CustomizationID, *deleteCorpusOptions.CorpusName}

	builder := core.NewRequestBuilder(core.DELETE)
	builder.ConstructHTTPURL(speechToText.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteCorpusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.service.Request(request, nil)
	return response, err
}

// GetCorpus : Get a corpus
func (speechToText *SpeechToTextV1) GetCorpus(getCorpusOptions *GetCorpusOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(getCorpusOptions, "getCorpusOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(getCorpusOptions, "getCorpusOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/customizations", "corpora"}
	pathParameters := []string{*getCorpusOptions.CustomizationID, *getCorpusOptions.CorpusName}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(speechToText.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range getCorpusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.service.Request(request, new(Corpus))
	return response, err
}

// GetGetCorpusResult : Retrieve result of GetCorpus operation
func (speechToText *SpeechToTextV1) GetGetCorpusResult(response *core.DetailedResponse) *Corpus {
	result, ok := response.Result.(*Corpus)
	if ok {
		return result
	}
	return nil
}

// ListCorpora : List corpora
func (speechToText *SpeechToTextV1) ListCorpora(listCorporaOptions *ListCorporaOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(listCorporaOptions, "listCorporaOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(listCorporaOptions, "listCorporaOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/customizations", "corpora"}
	pathParameters := []string{*listCorporaOptions.CustomizationID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(speechToText.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range listCorporaOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.service.Request(request, new(Corpora))
	return response, err
}

// GetListCorporaResult : Retrieve result of ListCorpora operation
func (speechToText *SpeechToTextV1) GetListCorporaResult(response *core.DetailedResponse) *Corpora {
	result, ok := response.Result.(*Corpora)
	if ok {
		return result
	}
	return nil
}

// AddWord : Add a custom word
func (speechToText *SpeechToTextV1) AddWord(addWordOptions *AddWordOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(addWordOptions, "addWordOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(addWordOptions, "addWordOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/customizations", "words"}
	pathParameters := []string{*addWordOptions.CustomizationID, *addWordOptions.WordName}

	builder := core.NewRequestBuilder(core.PUT)
	builder.ConstructHTTPURL(speechToText.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range addWordOptions.Headers {
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
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.service.Request(request, nil)
	return response, err
}

// AddWords : Add custom words
func (speechToText *SpeechToTextV1) AddWords(addWordsOptions *AddWordsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(addWordsOptions, "addWordsOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(addWordsOptions, "addWordsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/customizations", "words"}
	pathParameters := []string{*addWordsOptions.CustomizationID}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(speechToText.service.Options.URL, pathSegments, pathParameters)

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

	response, err := speechToText.service.Request(request, nil)
	return response, err
}

// DeleteWord : Delete a custom word
func (speechToText *SpeechToTextV1) DeleteWord(deleteWordOptions *DeleteWordOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(deleteWordOptions, "deleteWordOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(deleteWordOptions, "deleteWordOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/customizations", "words"}
	pathParameters := []string{*deleteWordOptions.CustomizationID, *deleteWordOptions.WordName}

	builder := core.NewRequestBuilder(core.DELETE)
	builder.ConstructHTTPURL(speechToText.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteWordOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.service.Request(request, nil)
	return response, err
}

// GetWord : Get a custom word
func (speechToText *SpeechToTextV1) GetWord(getWordOptions *GetWordOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(getWordOptions, "getWordOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(getWordOptions, "getWordOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/customizations", "words"}
	pathParameters := []string{*getWordOptions.CustomizationID, *getWordOptions.WordName}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(speechToText.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range getWordOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.service.Request(request, new(Word))
	return response, err
}

// GetGetWordResult : Retrieve result of GetWord operation
func (speechToText *SpeechToTextV1) GetGetWordResult(response *core.DetailedResponse) *Word {
	result, ok := response.Result.(*Word)
	if ok {
		return result
	}
	return nil
}

// ListWords : List custom words
func (speechToText *SpeechToTextV1) ListWords(listWordsOptions *ListWordsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(listWordsOptions, "listWordsOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(listWordsOptions, "listWordsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/customizations", "words"}
	pathParameters := []string{*listWordsOptions.CustomizationID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(speechToText.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range listWordsOptions.Headers {
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
		return nil, err
	}

	response, err := speechToText.service.Request(request, new(Words))
	return response, err
}

// GetListWordsResult : Retrieve result of ListWords operation
func (speechToText *SpeechToTextV1) GetListWordsResult(response *core.DetailedResponse) *Words {
	result, ok := response.Result.(*Words)
	if ok {
		return result
	}
	return nil
}

// CreateAcousticModel : Create a custom acoustic model
func (speechToText *SpeechToTextV1) CreateAcousticModel(createAcousticModelOptions *CreateAcousticModelOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(createAcousticModelOptions, "createAcousticModelOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(createAcousticModelOptions, "createAcousticModelOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/acoustic_customizations"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(speechToText.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range createAcousticModelOptions.Headers {
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
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.service.Request(request, new(AcousticModel))
	return response, err
}

// GetCreateAcousticModelResult : Retrieve result of CreateAcousticModel operation
func (speechToText *SpeechToTextV1) GetCreateAcousticModelResult(response *core.DetailedResponse) *AcousticModel {
	result, ok := response.Result.(*AcousticModel)
	if ok {
		return result
	}
	return nil
}

// DeleteAcousticModel : Delete a custom acoustic model
func (speechToText *SpeechToTextV1) DeleteAcousticModel(deleteAcousticModelOptions *DeleteAcousticModelOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(deleteAcousticModelOptions, "deleteAcousticModelOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(deleteAcousticModelOptions, "deleteAcousticModelOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/acoustic_customizations"}
	pathParameters := []string{*deleteAcousticModelOptions.CustomizationID}

	builder := core.NewRequestBuilder(core.DELETE)
	builder.ConstructHTTPURL(speechToText.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteAcousticModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.service.Request(request, nil)
	return response, err
}

// GetAcousticModel : Get a custom acoustic model
func (speechToText *SpeechToTextV1) GetAcousticModel(getAcousticModelOptions *GetAcousticModelOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(getAcousticModelOptions, "getAcousticModelOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(getAcousticModelOptions, "getAcousticModelOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/acoustic_customizations"}
	pathParameters := []string{*getAcousticModelOptions.CustomizationID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(speechToText.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range getAcousticModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.service.Request(request, new(AcousticModel))
	return response, err
}

// GetGetAcousticModelResult : Retrieve result of GetAcousticModel operation
func (speechToText *SpeechToTextV1) GetGetAcousticModelResult(response *core.DetailedResponse) *AcousticModel {
	result, ok := response.Result.(*AcousticModel)
	if ok {
		return result
	}
	return nil
}

// ListAcousticModels : List custom acoustic models
func (speechToText *SpeechToTextV1) ListAcousticModels(listAcousticModelsOptions *ListAcousticModelsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateStruct(listAcousticModelsOptions, "listAcousticModelsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/acoustic_customizations"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(speechToText.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range listAcousticModelsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listAcousticModelsOptions.Language != nil {
		builder.AddQuery("language", fmt.Sprint(*listAcousticModelsOptions.Language))
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.service.Request(request, new(AcousticModels))
	return response, err
}

// GetListAcousticModelsResult : Retrieve result of ListAcousticModels operation
func (speechToText *SpeechToTextV1) GetListAcousticModelsResult(response *core.DetailedResponse) *AcousticModels {
	result, ok := response.Result.(*AcousticModels)
	if ok {
		return result
	}
	return nil
}

// ResetAcousticModel : Reset a custom acoustic model
func (speechToText *SpeechToTextV1) ResetAcousticModel(resetAcousticModelOptions *ResetAcousticModelOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(resetAcousticModelOptions, "resetAcousticModelOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(resetAcousticModelOptions, "resetAcousticModelOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/acoustic_customizations", "reset"}
	pathParameters := []string{*resetAcousticModelOptions.CustomizationID}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(speechToText.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range resetAcousticModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.service.Request(request, nil)
	return response, err
}

// TrainAcousticModel : Train a custom acoustic model
func (speechToText *SpeechToTextV1) TrainAcousticModel(trainAcousticModelOptions *TrainAcousticModelOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(trainAcousticModelOptions, "trainAcousticModelOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(trainAcousticModelOptions, "trainAcousticModelOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/acoustic_customizations", "train"}
	pathParameters := []string{*trainAcousticModelOptions.CustomizationID}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(speechToText.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range trainAcousticModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if trainAcousticModelOptions.CustomLanguageModelID != nil {
		builder.AddQuery("custom_language_model_id", fmt.Sprint(*trainAcousticModelOptions.CustomLanguageModelID))
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.service.Request(request, nil)
	return response, err
}

// UpgradeAcousticModel : Upgrade a custom acoustic model
func (speechToText *SpeechToTextV1) UpgradeAcousticModel(upgradeAcousticModelOptions *UpgradeAcousticModelOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(upgradeAcousticModelOptions, "upgradeAcousticModelOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(upgradeAcousticModelOptions, "upgradeAcousticModelOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/acoustic_customizations", "upgrade_model"}
	pathParameters := []string{*upgradeAcousticModelOptions.CustomizationID}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(speechToText.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range upgradeAcousticModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if upgradeAcousticModelOptions.CustomLanguageModelID != nil {
		builder.AddQuery("custom_language_model_id", fmt.Sprint(*upgradeAcousticModelOptions.CustomLanguageModelID))
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.service.Request(request, nil)
	return response, err
}

// AddAudio : Add an audio resource
func (speechToText *SpeechToTextV1) AddAudio(addAudioOptions *AddAudioOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(addAudioOptions, "addAudioOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(addAudioOptions, "addAudioOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/acoustic_customizations", "audio"}
	pathParameters := []string{*addAudioOptions.CustomizationID, *addAudioOptions.AudioName}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(speechToText.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range addAudioOptions.Headers {
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

	_, err := builder.SetBodyContent(core.StringNilMapper(addAudioOptions.ContentType), nil, nil, addAudioOptions.AudioResource)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.service.Request(request, nil)
	return response, err
}

// DeleteAudio : Delete an audio resource
func (speechToText *SpeechToTextV1) DeleteAudio(deleteAudioOptions *DeleteAudioOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(deleteAudioOptions, "deleteAudioOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(deleteAudioOptions, "deleteAudioOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/acoustic_customizations", "audio"}
	pathParameters := []string{*deleteAudioOptions.CustomizationID, *deleteAudioOptions.AudioName}

	builder := core.NewRequestBuilder(core.DELETE)
	builder.ConstructHTTPURL(speechToText.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteAudioOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.service.Request(request, nil)
	return response, err
}

// GetAudio : Get an audio resource
func (speechToText *SpeechToTextV1) GetAudio(getAudioOptions *GetAudioOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(getAudioOptions, "getAudioOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(getAudioOptions, "getAudioOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/acoustic_customizations", "audio"}
	pathParameters := []string{*getAudioOptions.CustomizationID, *getAudioOptions.AudioName}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(speechToText.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range getAudioOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.service.Request(request, new(AudioListing))
	return response, err
}

// GetGetAudioResult : Retrieve result of GetAudio operation
func (speechToText *SpeechToTextV1) GetGetAudioResult(response *core.DetailedResponse) *AudioListing {
	result, ok := response.Result.(*AudioListing)
	if ok {
		return result
	}
	return nil
}

// ListAudio : List audio resources
func (speechToText *SpeechToTextV1) ListAudio(listAudioOptions *ListAudioOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(listAudioOptions, "listAudioOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(listAudioOptions, "listAudioOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/acoustic_customizations", "audio"}
	pathParameters := []string{*listAudioOptions.CustomizationID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(speechToText.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range listAudioOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.service.Request(request, new(AudioResources))
	return response, err
}

// GetListAudioResult : Retrieve result of ListAudio operation
func (speechToText *SpeechToTextV1) GetListAudioResult(response *core.DetailedResponse) *AudioResources {
	result, ok := response.Result.(*AudioResources)
	if ok {
		return result
	}
	return nil
}

// DeleteUserData : Delete labeled data
func (speechToText *SpeechToTextV1) DeleteUserData(deleteUserDataOptions *DeleteUserDataOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(deleteUserDataOptions, "deleteUserDataOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(deleteUserDataOptions, "deleteUserDataOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/user_data"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.DELETE)
	builder.ConstructHTTPURL(speechToText.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteUserDataOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("customer_id", fmt.Sprint(*deleteUserDataOptions.CustomerID))

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.service.Request(request, nil)
	return response, err
}

// AcousticModel : AcousticModel struct
type AcousticModel struct {

	// The customization ID (GUID) of the custom acoustic model. The **Create a custom acoustic model** method returns only
	// this field of the object; it does not return the other fields.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// The date and time in Coordinated Universal Time (UTC) at which the custom acoustic model was created. The value is
	// provided in full ISO 8601 format (`YYYY-MM-DDThh:mm:ss.sTZD`).
	Created *string `json:"created,omitempty"`

	// The language identifier of the custom acoustic model (for example, `en-US`).
	Language *string `json:"language,omitempty"`

	// A list of the available versions of the custom acoustic model. Each element of the array indicates a version of the
	// base model with which the custom model can be used. Multiple versions exist only if the custom model has been
	// upgraded; otherwise, only a single version is shown.
	Versions []string `json:"versions,omitempty"`

	// The GUID of the service credentials for the instance of the service that owns the custom acoustic model.
	Owner *string `json:"owner,omitempty"`

	// The name of the custom acoustic model.
	Name *string `json:"name,omitempty"`

	// The description of the custom acoustic model.
	Description *string `json:"description,omitempty"`

	// The name of the language model for which the custom acoustic model was created.
	BaseModelName *string `json:"base_model_name,omitempty"`

	// The current status of the custom acoustic model:
	// * `pending` indicates that the model was created but is waiting either for training data to be added or for the
	// service to finish analyzing added data.
	// * `ready` indicates that the model contains data and is ready to be trained.
	// * `training` indicates that the model is currently being trained.
	// * `available` indicates that the model is trained and ready to use.
	// * `upgrading` indicates that the model is currently being upgraded.
	// * `failed` indicates that training of the model failed.
	Status *string `json:"status,omitempty"`

	// A percentage that indicates the progress of the custom acoustic model's current training. A value of `100` means
	// that the model is fully trained. **Note:** The `progress` field does not currently reflect the progress of the
	// training. The field changes from `0` to `100` when training is complete.
	Progress *int64 `json:"progress,omitempty"`

	// If the request included unknown parameters, the following message: `Unexpected query parameter(s) ['parameters']
	// detected`, where `parameters` is a list that includes a quoted string for each unknown parameter.
	Warnings *string `json:"warnings,omitempty"`
}

// AcousticModels : AcousticModels struct
type AcousticModels struct {

	// An array of objects that provides information about each available custom acoustic model. The array is empty if the
	// requesting service credentials own no custom acoustic models (if no language is specified) or own no custom acoustic
	// models for the specified language.
	Customizations []AcousticModel `json:"customizations" validate:"required"`
}

// AddAudioOptions : The addAudio options.
type AddAudioOptions struct {

	// The customization ID (GUID) of the custom acoustic model. You must make the request with service credentials created
	// for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// The name of the new audio resource for the custom acoustic model. Use a localized name that matches the language of
	// the custom model and reflects the contents of the resource.
	// * Include a maximum of 128 characters in the name.
	// * Do not include spaces, slashes, or backslashes in the name.
	// * Do not use the name of an audio resource that has already been added to the custom model.
	AudioName *string `json:"audio_name" validate:"required"`

	// The audio resource that is to be added to the custom acoustic model, an individual audio file or an archive file.
	AudioResource *io.ReadCloser `json:"audio_resource,omitempty"`

	// The type of the input.
	ContentType *string `json:"Content-Type" validate:"required"`

	// For an archive-type resource, specifies the format of the audio files contained in the archive file. The parameter
	// accepts all of the audio formats supported for use with speech recognition, including the `rate`, `channels`, and
	// `endianness` parameters that are used with some formats. For a complete list of supported audio formats, see [Audio
	// formats](/docs/services/speech-to-text/input.html#formats).
	ContainedContentType *string `json:"Contained-Content-Type,omitempty"`

	// If `true`, the specified corpus or audio resource overwrites an existing corpus or audio resource with the same
	// name. If `false`, the request fails if a corpus or audio resource with the same name already exists. The parameter
	// has no effect if a corpus or audio resource with the same name does not already exist.
	AllowOverwrite *bool `json:"allow_overwrite,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewAddAudioOptionsForZip : Instantiate AddAudioOptionsForZip
func (speechToText *SpeechToTextV1) NewAddAudioOptionsForZip(audioResource io.ReadCloser) *AddAudioOptions {
	return &AddAudioOptions{
		AudioResource: &audioResource,
		ContentType:   core.StringPtr("application/zip"),
	}
}

// NewAddAudioOptionsForGzip : Instantiate AddAudioOptionsForGzip
func (speechToText *SpeechToTextV1) NewAddAudioOptionsForGzip(audioResource io.ReadCloser) *AddAudioOptions {
	return &AddAudioOptions{
		AudioResource: &audioResource,
		ContentType:   core.StringPtr("application/gzip"),
	}
}

// NewAddAudioOptionsForBasic : Instantiate AddAudioOptionsForBasic
func (speechToText *SpeechToTextV1) NewAddAudioOptionsForBasic(audioResource io.ReadCloser) *AddAudioOptions {
	return &AddAudioOptions{
		AudioResource: &audioResource,
		ContentType:   core.StringPtr("audio/basic"),
	}
}

// NewAddAudioOptionsForFlac : Instantiate AddAudioOptionsForFlac
func (speechToText *SpeechToTextV1) NewAddAudioOptionsForFlac(audioResource io.ReadCloser) *AddAudioOptions {
	return &AddAudioOptions{
		AudioResource: &audioResource,
		ContentType:   core.StringPtr("audio/flac"),
	}
}

// NewAddAudioOptionsForL16 : Instantiate AddAudioOptionsForL16
func (speechToText *SpeechToTextV1) NewAddAudioOptionsForL16(audioResource io.ReadCloser) *AddAudioOptions {
	return &AddAudioOptions{
		AudioResource: &audioResource,
		ContentType:   core.StringPtr("audio/l16"),
	}
}

// NewAddAudioOptionsForMp3 : Instantiate AddAudioOptionsForMp3
func (speechToText *SpeechToTextV1) NewAddAudioOptionsForMp3(audioResource io.ReadCloser) *AddAudioOptions {
	return &AddAudioOptions{
		AudioResource: &audioResource,
		ContentType:   core.StringPtr("audio/mp3"),
	}
}

// NewAddAudioOptionsForMpeg : Instantiate AddAudioOptionsForMpeg
func (speechToText *SpeechToTextV1) NewAddAudioOptionsForMpeg(audioResource io.ReadCloser) *AddAudioOptions {
	return &AddAudioOptions{
		AudioResource: &audioResource,
		ContentType:   core.StringPtr("audio/mpeg"),
	}
}

// NewAddAudioOptionsForMulaw : Instantiate AddAudioOptionsForMulaw
func (speechToText *SpeechToTextV1) NewAddAudioOptionsForMulaw(audioResource io.ReadCloser) *AddAudioOptions {
	return &AddAudioOptions{
		AudioResource: &audioResource,
		ContentType:   core.StringPtr("audio/mulaw"),
	}
}

// NewAddAudioOptionsForOgg : Instantiate AddAudioOptionsForOgg
func (speechToText *SpeechToTextV1) NewAddAudioOptionsForOgg(audioResource io.ReadCloser) *AddAudioOptions {
	return &AddAudioOptions{
		AudioResource: &audioResource,
		ContentType:   core.StringPtr("audio/ogg"),
	}
}

// NewAddAudioOptionsForOggcodecsopus : Instantiate AddAudioOptionsForOggcodecsopus
func (speechToText *SpeechToTextV1) NewAddAudioOptionsForOggcodecsopus(audioResource io.ReadCloser) *AddAudioOptions {
	return &AddAudioOptions{
		AudioResource: &audioResource,
		ContentType:   core.StringPtr("audio/ogg;codecs=opus"),
	}
}

// NewAddAudioOptionsForOggcodecsvorbis : Instantiate AddAudioOptionsForOggcodecsvorbis
func (speechToText *SpeechToTextV1) NewAddAudioOptionsForOggcodecsvorbis(audioResource io.ReadCloser) *AddAudioOptions {
	return &AddAudioOptions{
		AudioResource: &audioResource,
		ContentType:   core.StringPtr("audio/ogg;codecs=vorbis"),
	}
}

// NewAddAudioOptionsForWav : Instantiate AddAudioOptionsForWav
func (speechToText *SpeechToTextV1) NewAddAudioOptionsForWav(audioResource io.ReadCloser) *AddAudioOptions {
	return &AddAudioOptions{
		AudioResource: &audioResource,
		ContentType:   core.StringPtr("audio/wav"),
	}
}

// NewAddAudioOptionsForWebm : Instantiate AddAudioOptionsForWebm
func (speechToText *SpeechToTextV1) NewAddAudioOptionsForWebm(audioResource io.ReadCloser) *AddAudioOptions {
	return &AddAudioOptions{
		AudioResource: &audioResource,
		ContentType:   core.StringPtr("audio/webm"),
	}
}

// NewAddAudioOptionsForWebmcodecsopus : Instantiate AddAudioOptionsForWebmcodecsopus
func (speechToText *SpeechToTextV1) NewAddAudioOptionsForWebmcodecsopus(audioResource io.ReadCloser) *AddAudioOptions {
	return &AddAudioOptions{
		AudioResource: &audioResource,
		ContentType:   core.StringPtr("audio/webm;codecs=opus"),
	}
}

// NewAddAudioOptionsForWebmcodecsvorbis : Instantiate AddAudioOptionsForWebmcodecsvorbis
func (speechToText *SpeechToTextV1) NewAddAudioOptionsForWebmcodecsvorbis(audioResource io.ReadCloser) *AddAudioOptions {
	return &AddAudioOptions{
		AudioResource: &audioResource,
		ContentType:   core.StringPtr("audio/webm;codecs=vorbis"),
	}
}

// SetAudioResource : Allow user to set AudioResource with the specified content type
func (options *AddAudioOptions) SetAudioResource(audioResource io.ReadCloser, contentType string) *AddAudioOptions {
	options.AudioResource = &audioResource
	options.ContentType = core.StringPtr(contentType)
	return options
}

// NewAddAudioOptions : Instantiate AddAudioOptions
func (speechToText *SpeechToTextV1) NewAddAudioOptions(customizationID string, audioName string, contentType string) *AddAudioOptions {
	return &AddAudioOptions{
		CustomizationID: core.StringPtr(customizationID),
		AudioName:       core.StringPtr(audioName),
		ContentType:     core.StringPtr(contentType),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *AddAudioOptions) SetCustomizationID(customizationID string) *AddAudioOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetAudioName : Allow user to set AudioName
func (options *AddAudioOptions) SetAudioName(audioName string) *AddAudioOptions {
	options.AudioName = core.StringPtr(audioName)
	return options
}

// SetContainedContentType : Allow user to set ContainedContentType
func (options *AddAudioOptions) SetContainedContentType(containedContentType string) *AddAudioOptions {
	options.ContainedContentType = core.StringPtr(containedContentType)
	return options
}

// SetAllowOverwrite : Allow user to set AllowOverwrite
func (options *AddAudioOptions) SetAllowOverwrite(allowOverwrite bool) *AddAudioOptions {
	options.AllowOverwrite = core.BoolPtr(allowOverwrite)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *AddAudioOptions) SetHeaders(param map[string]string) *AddAudioOptions {
	options.Headers = param
	return options
}

// AddCorpusOptions : The addCorpus options.
type AddCorpusOptions struct {

	// The customization ID (GUID) of the custom language model. You must make the request with service credentials created
	// for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// The name of the new corpus for the custom language model. Use a localized name that matches the language of the
	// custom model and reflects the contents of the corpus.
	// * Include a maximum of 128 characters in the name.
	// * Do not include spaces, slashes, or backslashes in the name.
	// * Do not use the name of a corpus that has already been added to the custom model.
	// * Do not use the name `user`, which is reserved by the service to denote custom words that are added or modified by
	// the user.
	CorpusName *string `json:"corpus_name" validate:"required"`

	// A plain text file that contains the training data for the corpus. Encode the file in UTF-8 if it contains non-ASCII
	// characters; the service assumes UTF-8 encoding if it encounters non-ASCII characters. With cURL, use the
	// `--data-binary` option to upload the file for the request.
	CorpusFile *os.File `json:"corpus_file" validate:"required"`

	// The filename for corpusFile.
	CorpusFilename *string `json:"corpus_filename,omitempty"`

	// If `true`, the specified corpus or audio resource overwrites an existing corpus or audio resource with the same
	// name. If `false`, the request fails if a corpus or audio resource with the same name already exists. The parameter
	// has no effect if a corpus or audio resource with the same name does not already exist.
	AllowOverwrite *bool `json:"allow_overwrite,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewAddCorpusOptions : Instantiate AddCorpusOptions
func (speechToText *SpeechToTextV1) NewAddCorpusOptions(customizationID string, corpusName string, corpusFile *os.File) *AddCorpusOptions {
	return &AddCorpusOptions{
		CustomizationID: core.StringPtr(customizationID),
		CorpusName:      core.StringPtr(corpusName),
		CorpusFile:      corpusFile,
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *AddCorpusOptions) SetCustomizationID(customizationID string) *AddCorpusOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetCorpusName : Allow user to set CorpusName
func (options *AddCorpusOptions) SetCorpusName(corpusName string) *AddCorpusOptions {
	options.CorpusName = core.StringPtr(corpusName)
	return options
}

// SetCorpusFile : Allow user to set CorpusFile
func (options *AddCorpusOptions) SetCorpusFile(corpusFile *os.File) *AddCorpusOptions {
	options.CorpusFile = corpusFile
	return options
}

// SetCorpusFilename : Allow user to set CorpusFilename
func (options *AddCorpusOptions) SetCorpusFilename(corpusFilename string) *AddCorpusOptions {
	options.CorpusFilename = core.StringPtr(corpusFilename)
	return options
}

// SetAllowOverwrite : Allow user to set AllowOverwrite
func (options *AddCorpusOptions) SetAllowOverwrite(allowOverwrite bool) *AddCorpusOptions {
	options.AllowOverwrite = core.BoolPtr(allowOverwrite)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *AddCorpusOptions) SetHeaders(param map[string]string) *AddCorpusOptions {
	options.Headers = param
	return options
}

// AddWordOptions : The addWord options.
type AddWordOptions struct {

	// The customization ID (GUID) of the custom language model. You must make the request with service credentials created
	// for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// The custom word for the custom language model. When you add or update a custom word with the **Add a custom word**
	// method, do not include spaces in the word. Use a `-` (dash) or `_` (underscore) to connect the tokens of compound
	// words.
	WordName *string `json:"word_name" validate:"required"`

	// For the **Add custom words** method, you must specify the custom word that is to be added to or updated in the
	// custom model. Do not include spaces in the word. Use a `-` (dash) or `_` (underscore) to connect the tokens of
	// compound words.
	//
	// Omit this field for the **Add a custom word** method.
	Word *string `json:"word,omitempty"`

	// An array of sounds-like pronunciations for the custom word. Specify how words that are difficult to pronounce,
	// foreign words, acronyms, and so on can be pronounced by users.
	// * For a word that is not in the service's base vocabulary, omit the parameter to have the service automatically
	// generate a sounds-like pronunciation for the word.
	// * For a word that is in the service's base vocabulary, use the parameter to specify additional pronunciations for
	// the word. You cannot override the default pronunciation of a word; pronunciations you add augment the pronunciation
	// from the base vocabulary.
	//
	// A word can have at most five sounds-like pronunciations. A pronunciation can include at most 40 characters not
	// including spaces.
	SoundsLike []string `json:"sounds_like,omitempty"`

	// An alternative spelling for the custom word when it appears in a transcript. Use the parameter when you want the
	// word to have a spelling that is different from its usual representation or from its spelling in corpora training
	// data.
	DisplayAs *string `json:"display_as,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewAddWordOptions : Instantiate AddWordOptions
func (speechToText *SpeechToTextV1) NewAddWordOptions(customizationID string, wordName string) *AddWordOptions {
	return &AddWordOptions{
		CustomizationID: core.StringPtr(customizationID),
		WordName:        core.StringPtr(wordName),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *AddWordOptions) SetCustomizationID(customizationID string) *AddWordOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetWordName : Allow user to set WordName
func (options *AddWordOptions) SetWordName(wordName string) *AddWordOptions {
	options.WordName = core.StringPtr(wordName)
	return options
}

// SetWord : Allow user to set Word
func (options *AddWordOptions) SetWord(word string) *AddWordOptions {
	options.Word = core.StringPtr(word)
	return options
}

// SetSoundsLike : Allow user to set SoundsLike
func (options *AddWordOptions) SetSoundsLike(soundsLike []string) *AddWordOptions {
	options.SoundsLike = soundsLike
	return options
}

// SetDisplayAs : Allow user to set DisplayAs
func (options *AddWordOptions) SetDisplayAs(displayAs string) *AddWordOptions {
	options.DisplayAs = core.StringPtr(displayAs)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *AddWordOptions) SetHeaders(param map[string]string) *AddWordOptions {
	options.Headers = param
	return options
}

// AddWordsOptions : The addWords options.
type AddWordsOptions struct {

	// The customization ID (GUID) of the custom language model. You must make the request with service credentials created
	// for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// An array of objects that provides information about each custom word that is to be added to or updated in the custom
	// language model.
	Words []CustomWord `json:"words" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewAddWordsOptions : Instantiate AddWordsOptions
func (speechToText *SpeechToTextV1) NewAddWordsOptions(customizationID string, words []CustomWord) *AddWordsOptions {
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
func (options *AddWordsOptions) SetWords(words []CustomWord) *AddWordsOptions {
	options.Words = words
	return options
}

// SetHeaders : Allow user to set Headers
func (options *AddWordsOptions) SetHeaders(param map[string]string) *AddWordsOptions {
	options.Headers = param
	return options
}

// AudioDetails : AudioDetails struct
type AudioDetails struct {

	// The type of the audio resource:
	// * `audio` for an individual audio file
	// * `archive` for an archive (**.zip** or **.tar.gz**) file that contains audio files
	// * `undetermined` for a resource that the service cannot validate (for example, if the user mistakenly passes a file
	// that does not contain audio, such as a JPEG file).
	Type *string `json:"type,omitempty"`

	// **For an audio-type resource,** the codec in which the audio is encoded. Omitted for an archive-type resource.
	Codec *string `json:"codec,omitempty"`

	// **For an audio-type resource,** the sampling rate of the audio in Hertz (samples per second). Omitted for an
	// archive-type resource.
	Frequency *int64 `json:"frequency,omitempty"`

	// **For an archive-type resource,** the format of the compressed archive:
	// * `zip` for a **.zip** file
	// * `gzip` for a **.tar.gz** file
	//
	// Omitted for an audio-type resource.
	Compression *string `json:"compression,omitempty"`
}

// AudioListing : AudioListing struct
type AudioListing struct {

	// **For an audio-type resource,**  the total seconds of audio in the resource. The value is always a whole number.
	// Omitted for an archive-type resource.
	Duration *float64 `json:"duration,omitempty"`

	// **For an audio-type resource,** the user-specified name of the resource. Omitted for an archive-type resource.
	Name *string `json:"name,omitempty"`

	// **For an audio-type resource,** an `AudioDetails` object that provides detailed information about the resource. The
	// object is empty until the service finishes processing the audio. Omitted for an archive-type resource.
	Details *AudioDetails `json:"details,omitempty"`

	// **For an audio-type resource,** the status of the resource:
	// * `ok` indicates that the service has successfully analyzed the audio data. The data can be used to train the custom
	// model.
	// * `being_processed` indicates that the service is still analyzing the audio data. The service cannot accept requests
	// to add new audio resources or to train the custom model until its analysis is complete.
	// * `invalid` indicates that the audio data is not valid for training the custom model (possibly because it has the
	// wrong format or sampling rate, or because it is corrupted).
	//
	// Omitted for an archive-type resource.
	Status *string `json:"status,omitempty"`

	// **For an archive-type resource,** an object of type `AudioResource` that provides information about the resource.
	// Omitted for an audio-type resource.
	Container *AudioResource `json:"container,omitempty"`

	// **For an archive-type resource,** an array of `AudioResource` objects that provides information about the audio-type
	// resources that are contained in the resource. Omitted for an audio-type resource.
	Audio []AudioResource `json:"audio,omitempty"`
}

// AudioResource : AudioResource struct
type AudioResource struct {

	// The total seconds of audio in the audio resource. The value is always a whole number.
	Duration *float64 `json:"duration" validate:"required"`

	// **For an archive-type resource,** the user-specified name of the resource.
	//
	// **For an audio-type resource,** the user-specified name of the resource or the name of the audio file that the user
	// added for the resource. The value depends on the method that is called.
	Name *string `json:"name" validate:"required"`

	// An `AudioDetails` object that provides detailed information about the audio resource. The object is empty until the
	// service finishes processing the audio.
	Details *AudioDetails `json:"details" validate:"required"`

	// The status of the audio resource:
	// * `ok` indicates that the service has successfully analyzed the audio data. The data can be used to train the custom
	// model.
	// * `being_processed` indicates that the service is still analyzing the audio data. The service cannot accept requests
	// to add new audio resources or to train the custom model until its analysis is complete.
	// * `invalid` indicates that the audio data is not valid for training the custom model (possibly because it has the
	// wrong format or sampling rate, or because it is corrupted). For an archive file, the entire archive is invalid if
	// any of its audio files are invalid.
	Status *string `json:"status" validate:"required"`
}

// AudioResources : AudioResources struct
type AudioResources struct {

	// The total minutes of accumulated audio summed over all of the valid audio resources for the custom acoustic model.
	// You can use this value to determine whether the custom model has too little or too much audio to begin training.
	TotalMinutesOfAudio *float64 `json:"total_minutes_of_audio" validate:"required"`

	// An array of objects that provides information about the audio resources of the custom acoustic model. The array is
	// empty if the custom model has no audio resources.
	Audio []AudioResource `json:"audio" validate:"required"`
}

// CheckJobOptions : The checkJob options.
type CheckJobOptions struct {

	// The ID of the asynchronous job.
	ID *string `json:"id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewCheckJobOptions : Instantiate CheckJobOptions
func (speechToText *SpeechToTextV1) NewCheckJobOptions(ID string) *CheckJobOptions {
	return &CheckJobOptions{
		ID: core.StringPtr(ID),
	}
}

// SetID : Allow user to set ID
func (options *CheckJobOptions) SetID(ID string) *CheckJobOptions {
	options.ID = core.StringPtr(ID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CheckJobOptions) SetHeaders(param map[string]string) *CheckJobOptions {
	options.Headers = param
	return options
}

// CheckJobsOptions : The checkJobs options.
type CheckJobsOptions struct {

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewCheckJobsOptions : Instantiate CheckJobsOptions
func (speechToText *SpeechToTextV1) NewCheckJobsOptions() *CheckJobsOptions {
	return &CheckJobsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *CheckJobsOptions) SetHeaders(param map[string]string) *CheckJobsOptions {
	options.Headers = param
	return options
}

// Corpora : Corpora struct
type Corpora struct {

	// An array of objects that provides information about the corpora for the custom model. The array is empty if the
	// custom model has no corpora.
	Corpora []Corpus `json:"corpora" validate:"required"`
}

// Corpus : Corpus struct
type Corpus struct {

	// The name of the corpus.
	Name *string `json:"name" validate:"required"`

	// The total number of words in the corpus. The value is `0` while the corpus is being processed.
	TotalWords *int64 `json:"total_words" validate:"required"`

	// The number of OOV words in the corpus. The value is `0` while the corpus is being processed.
	OutOfVocabularyWords *int64 `json:"out_of_vocabulary_words" validate:"required"`

	// The status of the corpus:
	// * `analyzed` indicates that the service has successfully analyzed the corpus; the custom model can be trained with
	// data from the corpus.
	// * `being_processed` indicates that the service is still analyzing the corpus; the service cannot accept requests to
	// add new corpora or words, or to train the custom model.
	// * `undetermined` indicates that the service encountered an error while processing the corpus.
	Status *string `json:"status" validate:"required"`

	// If the status of the corpus is `undetermined`, the following message: `Analysis of corpus 'name' failed. Please try
	// adding the corpus again by setting the 'allow_overwrite' flag to 'true'`.
	Error *string `json:"error,omitempty"`
}

// CreateAcousticModelOptions : The createAcousticModel options.
type CreateAcousticModelOptions struct {

	// A user-defined name for the new custom acoustic model. Use a name that is unique among all custom acoustic models
	// that you own. Use a localized name that matches the language of the custom model. Use a name that describes the
	// acoustic environment of the custom model, such as `Mobile custom model` or `Noisy car custom model`.
	Name *string `json:"name" validate:"required"`

	// The name of the base language model that is to be customized by the new custom acoustic model. The new custom model
	// can be used only with the base model that it customizes. To determine whether a base model supports acoustic model
	// customization, refer to [Language support for
	// customization](https://console.bluemix.net/docs/services/speech-to-text/custom.html#languageSupport).
	BaseModelName *string `json:"base_model_name" validate:"required"`

	// A description of the new custom acoustic model. Use a localized description that matches the language of the custom
	// model.
	Description *string `json:"description,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewCreateAcousticModelOptions : Instantiate CreateAcousticModelOptions
func (speechToText *SpeechToTextV1) NewCreateAcousticModelOptions(name string, baseModelName string) *CreateAcousticModelOptions {
	return &CreateAcousticModelOptions{
		Name:          core.StringPtr(name),
		BaseModelName: core.StringPtr(baseModelName),
	}
}

// SetName : Allow user to set Name
func (options *CreateAcousticModelOptions) SetName(name string) *CreateAcousticModelOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetBaseModelName : Allow user to set BaseModelName
func (options *CreateAcousticModelOptions) SetBaseModelName(baseModelName string) *CreateAcousticModelOptions {
	options.BaseModelName = core.StringPtr(baseModelName)
	return options
}

// SetDescription : Allow user to set Description
func (options *CreateAcousticModelOptions) SetDescription(description string) *CreateAcousticModelOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateAcousticModelOptions) SetHeaders(param map[string]string) *CreateAcousticModelOptions {
	options.Headers = param
	return options
}

// CreateJobOptions : The createJob options.
type CreateJobOptions struct {

	// The audio to transcribe in the format specified by the `Content-Type` header.
	Audio *io.ReadCloser `json:"audio,omitempty"`

	// The type of the input.
	ContentType *string `json:"Content-Type" validate:"required"`

	// The identifier of the model that is to be used for the recognition request.
	Model *string `json:"model,omitempty"`

	// A URL to which callback notifications are to be sent. The URL must already be successfully white-listed by using the
	// **Register a callback** method. You can include the same callback URL with any number of job creation requests. Omit
	// the parameter to poll the service for job completion and results.
	//
	// Use the `user_token` parameter to specify a unique user-specified string with each job to differentiate the callback
	// notifications for the jobs.
	CallbackURL *string `json:"callback_url,omitempty"`

	// If the job includes a callback URL, a comma-separated list of notification events to which to subscribe. Valid
	// events are
	// * `recognitions.started` generates a callback notification when the service begins to process the job.
	// * `recognitions.completed` generates a callback notification when the job is complete. You must use the **Check a
	// job** method to retrieve the results before they time out or are deleted.
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
	// make the request with service credentials created for the instance of the service that owns the custom model. By
	// default, no custom language model is used.
	CustomizationID *string `json:"customization_id,omitempty"`

	// The customization ID (GUID) of a custom acoustic model that is to be used with the recognition request. The base
	// model of the specified custom acoustic model must match the model specified with the `model` parameter. You must
	// make the request with service credentials created for the instance of the service that owns the custom model. By
	// default, no custom acoustic model is used.
	AcousticCustomizationID *string `json:"acoustic_customization_id,omitempty"`

	// The version of the specified base model that is to be used with recognition request. Multiple versions of a base
	// model can exist when a model is updated for internal improvements. The parameter is intended primarily for use with
	// custom models that have been upgraded for a new base model. The default value depends on whether the parameter is
	// used with or without a custom model. For more information, see [Base model
	// version](https://console.bluemix.net/docs/services/speech-to-text/input.html#version).
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
	CustomizationWeight *float64 `json:"customization_weight,omitempty"`

	// The time in seconds after which, if only silence (no speech) is detected in submitted audio, the connection is
	// closed with a 400 error. The parameter is useful for stopping audio submission from a live microphone when a user
	// simply walks away. Use `-1` for infinity.
	InactivityTimeout *int64 `json:"inactivity_timeout,omitempty"`

	// An array of keyword strings to spot in the audio. Each keyword string can include one or more string tokens.
	// Keywords are spotted only in the final results, not in interim hypotheses. If you specify any keywords, you must
	// also specify a keywords threshold. You can spot a maximum of 1000 keywords. Omit the parameter or specify an empty
	// array if you do not need to spot keywords.
	Keywords []string `json:"keywords,omitempty"`

	// A confidence value that is the lower bound for spotting a keyword. A word is considered to match a keyword if its
	// confidence is greater than or equal to the threshold. Specify a probability between 0.0 and 1.0. No keyword spotting
	// is performed if you omit the parameter. If you specify a threshold, you must also specify one or more keywords.
	KeywordsThreshold *float32 `json:"keywords_threshold,omitempty"`

	// The maximum number of alternative transcripts that the service is to return. By default, a single transcription is
	// returned.
	MaxAlternatives *int64 `json:"max_alternatives,omitempty"`

	// A confidence value that is the lower bound for identifying a hypothesis as a possible word alternative (also known
	// as "Confusion Networks"). An alternative word is considered if its confidence is greater than or equal to the
	// threshold. Specify a probability between 0.0 and 1.0. No alternative words are computed if you omit the parameter.
	WordAlternativesThreshold *float32 `json:"word_alternatives_threshold,omitempty"`

	// If `true`, the service returns a confidence measure in the range of 0.0 to 1.0 for each word. By default, no word
	// confidence measures are returned.
	WordConfidence *bool `json:"word_confidence,omitempty"`

	// If `true`, the service returns time alignment for each word. By default, no timestamps are returned.
	Timestamps *bool `json:"timestamps,omitempty"`

	// If `true`, the service filters profanity from all output except for keyword results by replacing inappropriate words
	// with a series of asterisks. Set the parameter to `false` to return results with no censoring. Applies to US English
	// transcription only.
	ProfanityFilter *bool `json:"profanity_filter,omitempty"`

	// If `true`, the service converts dates, times, series of digits and numbers, phone numbers, currency values, and
	// internet addresses into more readable, conventional representations in the final transcript of a recognition
	// request. For US English, the service also converts certain keyword strings to punctuation symbols. By default, no
	// smart formatting is performed. Applies to US English and Spanish transcription only.
	SmartFormatting *bool `json:"smart_formatting,omitempty"`

	// If `true`, the response includes labels that identify which words were spoken by which participants in a
	// multi-person exchange. By default, no speaker labels are returned. Setting `speaker_labels` to `true` forces the
	// `timestamps` parameter to be `true`, regardless of whether you specify `false` for the parameter.
	//
	//  To determine whether a language model supports speaker labels, use the **Get models** method and check that the
	// attribute `speaker_labels` is set to `true`. You can also refer to [Speaker
	// labels](https://console.bluemix.net/docs/services/speech-to-text/output.html#speaker_labels).
	SpeakerLabels *bool `json:"speaker_labels,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewCreateJobOptionsForBasic : Instantiate CreateJobOptionsForBasic
func (speechToText *SpeechToTextV1) NewCreateJobOptionsForBasic(audio io.ReadCloser) *CreateJobOptions {
	return &CreateJobOptions{
		Audio:       &audio,
		ContentType: core.StringPtr("audio/basic"),
	}
}

// NewCreateJobOptionsForFlac : Instantiate CreateJobOptionsForFlac
func (speechToText *SpeechToTextV1) NewCreateJobOptionsForFlac(audio io.ReadCloser) *CreateJobOptions {
	return &CreateJobOptions{
		Audio:       &audio,
		ContentType: core.StringPtr("audio/flac"),
	}
}

// NewCreateJobOptionsForL16 : Instantiate CreateJobOptionsForL16
func (speechToText *SpeechToTextV1) NewCreateJobOptionsForL16(audio io.ReadCloser) *CreateJobOptions {
	return &CreateJobOptions{
		Audio:       &audio,
		ContentType: core.StringPtr("audio/l16"),
	}
}

// NewCreateJobOptionsForMp3 : Instantiate CreateJobOptionsForMp3
func (speechToText *SpeechToTextV1) NewCreateJobOptionsForMp3(audio io.ReadCloser) *CreateJobOptions {
	return &CreateJobOptions{
		Audio:       &audio,
		ContentType: core.StringPtr("audio/mp3"),
	}
}

// NewCreateJobOptionsForMpeg : Instantiate CreateJobOptionsForMpeg
func (speechToText *SpeechToTextV1) NewCreateJobOptionsForMpeg(audio io.ReadCloser) *CreateJobOptions {
	return &CreateJobOptions{
		Audio:       &audio,
		ContentType: core.StringPtr("audio/mpeg"),
	}
}

// NewCreateJobOptionsForMulaw : Instantiate CreateJobOptionsForMulaw
func (speechToText *SpeechToTextV1) NewCreateJobOptionsForMulaw(audio io.ReadCloser) *CreateJobOptions {
	return &CreateJobOptions{
		Audio:       &audio,
		ContentType: core.StringPtr("audio/mulaw"),
	}
}

// NewCreateJobOptionsForOgg : Instantiate CreateJobOptionsForOgg
func (speechToText *SpeechToTextV1) NewCreateJobOptionsForOgg(audio io.ReadCloser) *CreateJobOptions {
	return &CreateJobOptions{
		Audio:       &audio,
		ContentType: core.StringPtr("audio/ogg"),
	}
}

// NewCreateJobOptionsForOggcodecsopus : Instantiate CreateJobOptionsForOggcodecsopus
func (speechToText *SpeechToTextV1) NewCreateJobOptionsForOggcodecsopus(audio io.ReadCloser) *CreateJobOptions {
	return &CreateJobOptions{
		Audio:       &audio,
		ContentType: core.StringPtr("audio/ogg;codecs=opus"),
	}
}

// NewCreateJobOptionsForOggcodecsvorbis : Instantiate CreateJobOptionsForOggcodecsvorbis
func (speechToText *SpeechToTextV1) NewCreateJobOptionsForOggcodecsvorbis(audio io.ReadCloser) *CreateJobOptions {
	return &CreateJobOptions{
		Audio:       &audio,
		ContentType: core.StringPtr("audio/ogg;codecs=vorbis"),
	}
}

// NewCreateJobOptionsForWav : Instantiate CreateJobOptionsForWav
func (speechToText *SpeechToTextV1) NewCreateJobOptionsForWav(audio io.ReadCloser) *CreateJobOptions {
	return &CreateJobOptions{
		Audio:       &audio,
		ContentType: core.StringPtr("audio/wav"),
	}
}

// NewCreateJobOptionsForWebm : Instantiate CreateJobOptionsForWebm
func (speechToText *SpeechToTextV1) NewCreateJobOptionsForWebm(audio io.ReadCloser) *CreateJobOptions {
	return &CreateJobOptions{
		Audio:       &audio,
		ContentType: core.StringPtr("audio/webm"),
	}
}

// NewCreateJobOptionsForWebmcodecsopus : Instantiate CreateJobOptionsForWebmcodecsopus
func (speechToText *SpeechToTextV1) NewCreateJobOptionsForWebmcodecsopus(audio io.ReadCloser) *CreateJobOptions {
	return &CreateJobOptions{
		Audio:       &audio,
		ContentType: core.StringPtr("audio/webm;codecs=opus"),
	}
}

// NewCreateJobOptionsForWebmcodecsvorbis : Instantiate CreateJobOptionsForWebmcodecsvorbis
func (speechToText *SpeechToTextV1) NewCreateJobOptionsForWebmcodecsvorbis(audio io.ReadCloser) *CreateJobOptions {
	return &CreateJobOptions{
		Audio:       &audio,
		ContentType: core.StringPtr("audio/webm;codecs=vorbis"),
	}
}

// SetAudio : Allow user to set Audio with the specified content type
func (options *CreateJobOptions) SetAudio(audio io.ReadCloser, contentType string) *CreateJobOptions {
	options.Audio = &audio
	options.ContentType = core.StringPtr(contentType)
	return options
}

// NewCreateJobOptions : Instantiate CreateJobOptions
func (speechToText *SpeechToTextV1) NewCreateJobOptions(contentType string) *CreateJobOptions {
	return &CreateJobOptions{
		ContentType: core.StringPtr(contentType),
	}
}

// SetModel : Allow user to set Model
func (options *CreateJobOptions) SetModel(model string) *CreateJobOptions {
	options.Model = core.StringPtr(model)
	return options
}

// SetCallbackURL : Allow user to set CallbackURL
func (options *CreateJobOptions) SetCallbackURL(callbackURL string) *CreateJobOptions {
	options.CallbackURL = core.StringPtr(callbackURL)
	return options
}

// SetEvents : Allow user to set Events
func (options *CreateJobOptions) SetEvents(events string) *CreateJobOptions {
	options.Events = core.StringPtr(events)
	return options
}

// SetUserToken : Allow user to set UserToken
func (options *CreateJobOptions) SetUserToken(userToken string) *CreateJobOptions {
	options.UserToken = core.StringPtr(userToken)
	return options
}

// SetResultsTTL : Allow user to set ResultsTTL
func (options *CreateJobOptions) SetResultsTTL(resultsTTL int64) *CreateJobOptions {
	options.ResultsTTL = core.Int64Ptr(resultsTTL)
	return options
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *CreateJobOptions) SetCustomizationID(customizationID string) *CreateJobOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetAcousticCustomizationID : Allow user to set AcousticCustomizationID
func (options *CreateJobOptions) SetAcousticCustomizationID(acousticCustomizationID string) *CreateJobOptions {
	options.AcousticCustomizationID = core.StringPtr(acousticCustomizationID)
	return options
}

// SetBaseModelVersion : Allow user to set BaseModelVersion
func (options *CreateJobOptions) SetBaseModelVersion(baseModelVersion string) *CreateJobOptions {
	options.BaseModelVersion = core.StringPtr(baseModelVersion)
	return options
}

// SetCustomizationWeight : Allow user to set CustomizationWeight
func (options *CreateJobOptions) SetCustomizationWeight(customizationWeight float64) *CreateJobOptions {
	options.CustomizationWeight = core.Float64Ptr(customizationWeight)
	return options
}

// SetInactivityTimeout : Allow user to set InactivityTimeout
func (options *CreateJobOptions) SetInactivityTimeout(inactivityTimeout int64) *CreateJobOptions {
	options.InactivityTimeout = core.Int64Ptr(inactivityTimeout)
	return options
}

// SetKeywords : Allow user to set Keywords
func (options *CreateJobOptions) SetKeywords(keywords []string) *CreateJobOptions {
	options.Keywords = keywords
	return options
}

// SetKeywordsThreshold : Allow user to set KeywordsThreshold
func (options *CreateJobOptions) SetKeywordsThreshold(keywordsThreshold float32) *CreateJobOptions {
	options.KeywordsThreshold = core.Float32Ptr(keywordsThreshold)
	return options
}

// SetMaxAlternatives : Allow user to set MaxAlternatives
func (options *CreateJobOptions) SetMaxAlternatives(maxAlternatives int64) *CreateJobOptions {
	options.MaxAlternatives = core.Int64Ptr(maxAlternatives)
	return options
}

// SetWordAlternativesThreshold : Allow user to set WordAlternativesThreshold
func (options *CreateJobOptions) SetWordAlternativesThreshold(wordAlternativesThreshold float32) *CreateJobOptions {
	options.WordAlternativesThreshold = core.Float32Ptr(wordAlternativesThreshold)
	return options
}

// SetWordConfidence : Allow user to set WordConfidence
func (options *CreateJobOptions) SetWordConfidence(wordConfidence bool) *CreateJobOptions {
	options.WordConfidence = core.BoolPtr(wordConfidence)
	return options
}

// SetTimestamps : Allow user to set Timestamps
func (options *CreateJobOptions) SetTimestamps(timestamps bool) *CreateJobOptions {
	options.Timestamps = core.BoolPtr(timestamps)
	return options
}

// SetProfanityFilter : Allow user to set ProfanityFilter
func (options *CreateJobOptions) SetProfanityFilter(profanityFilter bool) *CreateJobOptions {
	options.ProfanityFilter = core.BoolPtr(profanityFilter)
	return options
}

// SetSmartFormatting : Allow user to set SmartFormatting
func (options *CreateJobOptions) SetSmartFormatting(smartFormatting bool) *CreateJobOptions {
	options.SmartFormatting = core.BoolPtr(smartFormatting)
	return options
}

// SetSpeakerLabels : Allow user to set SpeakerLabels
func (options *CreateJobOptions) SetSpeakerLabels(speakerLabels bool) *CreateJobOptions {
	options.SpeakerLabels = core.BoolPtr(speakerLabels)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateJobOptions) SetHeaders(param map[string]string) *CreateJobOptions {
	options.Headers = param
	return options
}

// CreateLanguageModelOptions : The createLanguageModel options.
type CreateLanguageModelOptions struct {

	// A user-defined name for the new custom language model. Use a name that is unique among all custom language models
	// that you own. Use a localized name that matches the language of the custom model. Use a name that describes the
	// domain of the custom model, such as `Medical custom model` or `Legal custom model`.
	Name *string `json:"name" validate:"required"`

	// The name of the base language model that is to be customized by the new custom language model. The new custom model
	// can be used only with the base model that it customizes. To determine whether a base model supports language model
	// customization, request information about the base model and check that the attribute `custom_language_model` is set
	// to `true`, or refer to [Language support for
	// customization](https://console.bluemix.net/docs/services/speech-to-text/custom.html#languageSupport).
	BaseModelName *string `json:"base_model_name" validate:"required"`

	// The dialect of the specified language that is to be used with the custom language model. The parameter is meaningful
	// only for Spanish models, for which the service creates a custom language model that is suited for speech in one of
	// the following dialects:
	// * `es-ES` for Castilian Spanish (the default)
	// * `es-LA` for Latin American Spanish
	// * `es-US` for North American (Mexican) Spanish
	//
	// A specified dialect must be valid for the base model. By default, the dialect matches the language of the base
	// model; for example, `en-US` for either of the US English language models.
	Dialect *string `json:"dialect,omitempty"`

	// A description of the new custom language model. Use a localized description that matches the language of the custom
	// model.
	Description *string `json:"description,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewCreateLanguageModelOptions : Instantiate CreateLanguageModelOptions
func (speechToText *SpeechToTextV1) NewCreateLanguageModelOptions(name string, baseModelName string) *CreateLanguageModelOptions {
	return &CreateLanguageModelOptions{
		Name:          core.StringPtr(name),
		BaseModelName: core.StringPtr(baseModelName),
	}
}

// SetName : Allow user to set Name
func (options *CreateLanguageModelOptions) SetName(name string) *CreateLanguageModelOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetBaseModelName : Allow user to set BaseModelName
func (options *CreateLanguageModelOptions) SetBaseModelName(baseModelName string) *CreateLanguageModelOptions {
	options.BaseModelName = core.StringPtr(baseModelName)
	return options
}

// SetDialect : Allow user to set Dialect
func (options *CreateLanguageModelOptions) SetDialect(dialect string) *CreateLanguageModelOptions {
	options.Dialect = core.StringPtr(dialect)
	return options
}

// SetDescription : Allow user to set Description
func (options *CreateLanguageModelOptions) SetDescription(description string) *CreateLanguageModelOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateLanguageModelOptions) SetHeaders(param map[string]string) *CreateLanguageModelOptions {
	options.Headers = param
	return options
}

// CustomWord : CustomWord struct
type CustomWord struct {

	// For the **Add custom words** method, you must specify the custom word that is to be added to or updated in the
	// custom model. Do not include spaces in the word. Use a `-` (dash) or `_` (underscore) to connect the tokens of
	// compound words.
	//
	// Omit this field for the **Add a custom word** method.
	Word *string `json:"word,omitempty"`

	// An array of sounds-like pronunciations for the custom word. Specify how words that are difficult to pronounce,
	// foreign words, acronyms, and so on can be pronounced by users.
	// * For a word that is not in the service's base vocabulary, omit the parameter to have the service automatically
	// generate a sounds-like pronunciation for the word.
	// * For a word that is in the service's base vocabulary, use the parameter to specify additional pronunciations for
	// the word. You cannot override the default pronunciation of a word; pronunciations you add augment the pronunciation
	// from the base vocabulary.
	//
	// A word can have at most five sounds-like pronunciations. A pronunciation can include at most 40 characters not
	// including spaces.
	SoundsLike []string `json:"sounds_like,omitempty"`

	// An alternative spelling for the custom word when it appears in a transcript. Use the parameter when you want the
	// word to have a spelling that is different from its usual representation or from its spelling in corpora training
	// data.
	DisplayAs *string `json:"display_as,omitempty"`
}

// DeleteAcousticModelOptions : The deleteAcousticModel options.
type DeleteAcousticModelOptions struct {

	// The customization ID (GUID) of the custom acoustic model. You must make the request with service credentials created
	// for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteAcousticModelOptions : Instantiate DeleteAcousticModelOptions
func (speechToText *SpeechToTextV1) NewDeleteAcousticModelOptions(customizationID string) *DeleteAcousticModelOptions {
	return &DeleteAcousticModelOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *DeleteAcousticModelOptions) SetCustomizationID(customizationID string) *DeleteAcousticModelOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteAcousticModelOptions) SetHeaders(param map[string]string) *DeleteAcousticModelOptions {
	options.Headers = param
	return options
}

// DeleteAudioOptions : The deleteAudio options.
type DeleteAudioOptions struct {

	// The customization ID (GUID) of the custom acoustic model. You must make the request with service credentials created
	// for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// The name of the audio resource for the custom acoustic model.
	AudioName *string `json:"audio_name" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteAudioOptions : Instantiate DeleteAudioOptions
func (speechToText *SpeechToTextV1) NewDeleteAudioOptions(customizationID string, audioName string) *DeleteAudioOptions {
	return &DeleteAudioOptions{
		CustomizationID: core.StringPtr(customizationID),
		AudioName:       core.StringPtr(audioName),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *DeleteAudioOptions) SetCustomizationID(customizationID string) *DeleteAudioOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetAudioName : Allow user to set AudioName
func (options *DeleteAudioOptions) SetAudioName(audioName string) *DeleteAudioOptions {
	options.AudioName = core.StringPtr(audioName)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteAudioOptions) SetHeaders(param map[string]string) *DeleteAudioOptions {
	options.Headers = param
	return options
}

// DeleteCorpusOptions : The deleteCorpus options.
type DeleteCorpusOptions struct {

	// The customization ID (GUID) of the custom language model. You must make the request with service credentials created
	// for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// The name of the corpus for the custom language model.
	CorpusName *string `json:"corpus_name" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteCorpusOptions : Instantiate DeleteCorpusOptions
func (speechToText *SpeechToTextV1) NewDeleteCorpusOptions(customizationID string, corpusName string) *DeleteCorpusOptions {
	return &DeleteCorpusOptions{
		CustomizationID: core.StringPtr(customizationID),
		CorpusName:      core.StringPtr(corpusName),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *DeleteCorpusOptions) SetCustomizationID(customizationID string) *DeleteCorpusOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetCorpusName : Allow user to set CorpusName
func (options *DeleteCorpusOptions) SetCorpusName(corpusName string) *DeleteCorpusOptions {
	options.CorpusName = core.StringPtr(corpusName)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteCorpusOptions) SetHeaders(param map[string]string) *DeleteCorpusOptions {
	options.Headers = param
	return options
}

// DeleteJobOptions : The deleteJob options.
type DeleteJobOptions struct {

	// The ID of the asynchronous job.
	ID *string `json:"id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteJobOptions : Instantiate DeleteJobOptions
func (speechToText *SpeechToTextV1) NewDeleteJobOptions(ID string) *DeleteJobOptions {
	return &DeleteJobOptions{
		ID: core.StringPtr(ID),
	}
}

// SetID : Allow user to set ID
func (options *DeleteJobOptions) SetID(ID string) *DeleteJobOptions {
	options.ID = core.StringPtr(ID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteJobOptions) SetHeaders(param map[string]string) *DeleteJobOptions {
	options.Headers = param
	return options
}

// DeleteLanguageModelOptions : The deleteLanguageModel options.
type DeleteLanguageModelOptions struct {

	// The customization ID (GUID) of the custom language model. You must make the request with service credentials created
	// for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteLanguageModelOptions : Instantiate DeleteLanguageModelOptions
func (speechToText *SpeechToTextV1) NewDeleteLanguageModelOptions(customizationID string) *DeleteLanguageModelOptions {
	return &DeleteLanguageModelOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *DeleteLanguageModelOptions) SetCustomizationID(customizationID string) *DeleteLanguageModelOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteLanguageModelOptions) SetHeaders(param map[string]string) *DeleteLanguageModelOptions {
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
func (speechToText *SpeechToTextV1) NewDeleteUserDataOptions(customerID string) *DeleteUserDataOptions {
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

// DeleteWordOptions : The deleteWord options.
type DeleteWordOptions struct {

	// The customization ID (GUID) of the custom language model. You must make the request with service credentials created
	// for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// The custom word for the custom language model. When you add or update a custom word with the **Add a custom word**
	// method, do not include spaces in the word. Use a `-` (dash) or `_` (underscore) to connect the tokens of compound
	// words.
	WordName *string `json:"word_name" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteWordOptions : Instantiate DeleteWordOptions
func (speechToText *SpeechToTextV1) NewDeleteWordOptions(customizationID string, wordName string) *DeleteWordOptions {
	return &DeleteWordOptions{
		CustomizationID: core.StringPtr(customizationID),
		WordName:        core.StringPtr(wordName),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *DeleteWordOptions) SetCustomizationID(customizationID string) *DeleteWordOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetWordName : Allow user to set WordName
func (options *DeleteWordOptions) SetWordName(wordName string) *DeleteWordOptions {
	options.WordName = core.StringPtr(wordName)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteWordOptions) SetHeaders(param map[string]string) *DeleteWordOptions {
	options.Headers = param
	return options
}

// GetAcousticModelOptions : The getAcousticModel options.
type GetAcousticModelOptions struct {

	// The customization ID (GUID) of the custom acoustic model. You must make the request with service credentials created
	// for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetAcousticModelOptions : Instantiate GetAcousticModelOptions
func (speechToText *SpeechToTextV1) NewGetAcousticModelOptions(customizationID string) *GetAcousticModelOptions {
	return &GetAcousticModelOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *GetAcousticModelOptions) SetCustomizationID(customizationID string) *GetAcousticModelOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetAcousticModelOptions) SetHeaders(param map[string]string) *GetAcousticModelOptions {
	options.Headers = param
	return options
}

// GetAudioOptions : The getAudio options.
type GetAudioOptions struct {

	// The customization ID (GUID) of the custom acoustic model. You must make the request with service credentials created
	// for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// The name of the audio resource for the custom acoustic model.
	AudioName *string `json:"audio_name" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetAudioOptions : Instantiate GetAudioOptions
func (speechToText *SpeechToTextV1) NewGetAudioOptions(customizationID string, audioName string) *GetAudioOptions {
	return &GetAudioOptions{
		CustomizationID: core.StringPtr(customizationID),
		AudioName:       core.StringPtr(audioName),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *GetAudioOptions) SetCustomizationID(customizationID string) *GetAudioOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetAudioName : Allow user to set AudioName
func (options *GetAudioOptions) SetAudioName(audioName string) *GetAudioOptions {
	options.AudioName = core.StringPtr(audioName)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetAudioOptions) SetHeaders(param map[string]string) *GetAudioOptions {
	options.Headers = param
	return options
}

// GetCorpusOptions : The getCorpus options.
type GetCorpusOptions struct {

	// The customization ID (GUID) of the custom language model. You must make the request with service credentials created
	// for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// The name of the corpus for the custom language model.
	CorpusName *string `json:"corpus_name" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetCorpusOptions : Instantiate GetCorpusOptions
func (speechToText *SpeechToTextV1) NewGetCorpusOptions(customizationID string, corpusName string) *GetCorpusOptions {
	return &GetCorpusOptions{
		CustomizationID: core.StringPtr(customizationID),
		CorpusName:      core.StringPtr(corpusName),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *GetCorpusOptions) SetCustomizationID(customizationID string) *GetCorpusOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetCorpusName : Allow user to set CorpusName
func (options *GetCorpusOptions) SetCorpusName(corpusName string) *GetCorpusOptions {
	options.CorpusName = core.StringPtr(corpusName)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetCorpusOptions) SetHeaders(param map[string]string) *GetCorpusOptions {
	options.Headers = param
	return options
}

// GetLanguageModelOptions : The getLanguageModel options.
type GetLanguageModelOptions struct {

	// The customization ID (GUID) of the custom language model. You must make the request with service credentials created
	// for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetLanguageModelOptions : Instantiate GetLanguageModelOptions
func (speechToText *SpeechToTextV1) NewGetLanguageModelOptions(customizationID string) *GetLanguageModelOptions {
	return &GetLanguageModelOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *GetLanguageModelOptions) SetCustomizationID(customizationID string) *GetLanguageModelOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetLanguageModelOptions) SetHeaders(param map[string]string) *GetLanguageModelOptions {
	options.Headers = param
	return options
}

// GetModelOptions : The getModel options.
type GetModelOptions struct {

	// The identifier of the model in the form of its name from the output of the **Get models** method.
	ModelID *string `json:"model_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetModelOptions : Instantiate GetModelOptions
func (speechToText *SpeechToTextV1) NewGetModelOptions(modelID string) *GetModelOptions {
	return &GetModelOptions{
		ModelID: core.StringPtr(modelID),
	}
}

// SetModelID : Allow user to set ModelID
func (options *GetModelOptions) SetModelID(modelID string) *GetModelOptions {
	options.ModelID = core.StringPtr(modelID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetModelOptions) SetHeaders(param map[string]string) *GetModelOptions {
	options.Headers = param
	return options
}

// GetWordOptions : The getWord options.
type GetWordOptions struct {

	// The customization ID (GUID) of the custom language model. You must make the request with service credentials created
	// for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// The custom word for the custom language model. When you add or update a custom word with the **Add a custom word**
	// method, do not include spaces in the word. Use a `-` (dash) or `_` (underscore) to connect the tokens of compound
	// words.
	WordName *string `json:"word_name" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetWordOptions : Instantiate GetWordOptions
func (speechToText *SpeechToTextV1) NewGetWordOptions(customizationID string, wordName string) *GetWordOptions {
	return &GetWordOptions{
		CustomizationID: core.StringPtr(customizationID),
		WordName:        core.StringPtr(wordName),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *GetWordOptions) SetCustomizationID(customizationID string) *GetWordOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetWordName : Allow user to set WordName
func (options *GetWordOptions) SetWordName(wordName string) *GetWordOptions {
	options.WordName = core.StringPtr(wordName)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetWordOptions) SetHeaders(param map[string]string) *GetWordOptions {
	options.Headers = param
	return options
}

// KeywordResult : KeywordResult struct
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

// LanguageModel : LanguageModel struct
type LanguageModel struct {

	// The customization ID (GUID) of the custom language model. The **Create a custom language model** method returns only
	// this field of the object; it does not return the other fields.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// The date and time in Coordinated Universal Time (UTC) at which the custom language model was created. The value is
	// provided in full ISO 8601 format (`YYYY-MM-DDThh:mm:ss.sTZD`).
	Created *string `json:"created,omitempty"`

	// The language identifier of the custom language model (for example, `en-US`).
	Language *string `json:"language,omitempty"`

	// The dialect of the language for the custom language model. By default, the dialect matches the language of the base
	// model; for example, `en-US` for either of the US English language models. For Spanish models, the field indicates
	// the dialect for which the model was created:
	// * `es-ES` for Castilian Spanish (the default)
	// * `es-LA` for Latin American Spanish
	// * `es-US` for North American (Mexican) Spanish.
	Dialect *string `json:"dialect,omitempty"`

	// A list of the available versions of the custom language model. Each element of the array indicates a version of the
	// base model with which the custom model can be used. Multiple versions exist only if the custom model has been
	// upgraded; otherwise, only a single version is shown.
	Versions []string `json:"versions,omitempty"`

	// The GUID of the service credentials for the instance of the service that owns the custom language model.
	Owner *string `json:"owner,omitempty"`

	// The name of the custom language model.
	Name *string `json:"name,omitempty"`

	// The description of the custom language model.
	Description *string `json:"description,omitempty"`

	// The name of the language model for which the custom language model was created.
	BaseModelName *string `json:"base_model_name,omitempty"`

	// The current status of the custom language model:
	// * `pending` indicates that the model was created but is waiting either for training data to be added or for the
	// service to finish analyzing added data.
	// * `ready` indicates that the model contains data and is ready to be trained.
	// * `training` indicates that the model is currently being trained.
	// * `available` indicates that the model is trained and ready to use.
	// * `upgrading` indicates that the model is currently being upgraded.
	// * `failed` indicates that training of the model failed.
	Status *string `json:"status,omitempty"`

	// A percentage that indicates the progress of the custom language model's current training. A value of `100` means
	// that the model is fully trained. **Note:** The `progress` field does not currently reflect the progress of the
	// training. The field changes from `0` to `100` when training is complete.
	Progress *int64 `json:"progress,omitempty"`

	// If the request included unknown parameters, the following message: `Unexpected query parameter(s) ['parameters']
	// detected`, where `parameters` is a list that includes a quoted string for each unknown parameter.
	Warnings *string `json:"warnings,omitempty"`
}

// LanguageModels : LanguageModels struct
type LanguageModels struct {

	// An array of objects that provides information about each available custom language model. The array is empty if the
	// requesting service credentials own no custom language models (if no language is specified) or own no custom language
	// models for the specified language.
	Customizations []LanguageModel `json:"customizations" validate:"required"`
}

// ListAcousticModelsOptions : The listAcousticModels options.
type ListAcousticModelsOptions struct {

	// The identifier of the language for which custom language or custom acoustic models are to be returned (for example,
	// `en-US`). Omit the parameter to see all custom language or custom acoustic models owned by the requesting service
	// credentials.
	Language *string `json:"language,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListAcousticModelsOptions : Instantiate ListAcousticModelsOptions
func (speechToText *SpeechToTextV1) NewListAcousticModelsOptions() *ListAcousticModelsOptions {
	return &ListAcousticModelsOptions{}
}

// SetLanguage : Allow user to set Language
func (options *ListAcousticModelsOptions) SetLanguage(language string) *ListAcousticModelsOptions {
	options.Language = core.StringPtr(language)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListAcousticModelsOptions) SetHeaders(param map[string]string) *ListAcousticModelsOptions {
	options.Headers = param
	return options
}

// ListAudioOptions : The listAudio options.
type ListAudioOptions struct {

	// The customization ID (GUID) of the custom acoustic model. You must make the request with service credentials created
	// for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListAudioOptions : Instantiate ListAudioOptions
func (speechToText *SpeechToTextV1) NewListAudioOptions(customizationID string) *ListAudioOptions {
	return &ListAudioOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *ListAudioOptions) SetCustomizationID(customizationID string) *ListAudioOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListAudioOptions) SetHeaders(param map[string]string) *ListAudioOptions {
	options.Headers = param
	return options
}

// ListCorporaOptions : The listCorpora options.
type ListCorporaOptions struct {

	// The customization ID (GUID) of the custom language model. You must make the request with service credentials created
	// for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListCorporaOptions : Instantiate ListCorporaOptions
func (speechToText *SpeechToTextV1) NewListCorporaOptions(customizationID string) *ListCorporaOptions {
	return &ListCorporaOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *ListCorporaOptions) SetCustomizationID(customizationID string) *ListCorporaOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListCorporaOptions) SetHeaders(param map[string]string) *ListCorporaOptions {
	options.Headers = param
	return options
}

// ListLanguageModelsOptions : The listLanguageModels options.
type ListLanguageModelsOptions struct {

	// The identifier of the language for which custom language or custom acoustic models are to be returned (for example,
	// `en-US`). Omit the parameter to see all custom language or custom acoustic models owned by the requesting service
	// credentials.
	Language *string `json:"language,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListLanguageModelsOptions : Instantiate ListLanguageModelsOptions
func (speechToText *SpeechToTextV1) NewListLanguageModelsOptions() *ListLanguageModelsOptions {
	return &ListLanguageModelsOptions{}
}

// SetLanguage : Allow user to set Language
func (options *ListLanguageModelsOptions) SetLanguage(language string) *ListLanguageModelsOptions {
	options.Language = core.StringPtr(language)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListLanguageModelsOptions) SetHeaders(param map[string]string) *ListLanguageModelsOptions {
	options.Headers = param
	return options
}

// ListModelsOptions : The listModels options.
type ListModelsOptions struct {

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListModelsOptions : Instantiate ListModelsOptions
func (speechToText *SpeechToTextV1) NewListModelsOptions() *ListModelsOptions {
	return &ListModelsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListModelsOptions) SetHeaders(param map[string]string) *ListModelsOptions {
	options.Headers = param
	return options
}

// ListWordsOptions : The listWords options.
type ListWordsOptions struct {

	// The customization ID (GUID) of the custom language model. You must make the request with service credentials created
	// for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// The type of words to be listed from the custom language model's words resource:
	// * `all` (the default) shows all words.
	// * `user` shows only custom words that were added or modified by the user.
	// * `corpora` shows only OOV that were extracted from corpora.
	WordType *string `json:"word_type,omitempty"`

	// Indicates the order in which the words are to be listed, `alphabetical` or by `count`. You can prepend an optional
	// `+` or `-` to an argument to indicate whether the results are to be sorted in ascending or descending order. By
	// default, words are sorted in ascending alphabetical order. For alphabetical ordering, the lexicographical precedence
	// is numeric values, uppercase letters, and lowercase letters. For count ordering, values with the same count are
	// ordered alphabetically. With cURL, URL encode the `+` symbol as `%2B`.
	Sort *string `json:"sort,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListWordsOptions : Instantiate ListWordsOptions
func (speechToText *SpeechToTextV1) NewListWordsOptions(customizationID string) *ListWordsOptions {
	return &ListWordsOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *ListWordsOptions) SetCustomizationID(customizationID string) *ListWordsOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetWordType : Allow user to set WordType
func (options *ListWordsOptions) SetWordType(wordType string) *ListWordsOptions {
	options.WordType = core.StringPtr(wordType)
	return options
}

// SetSort : Allow user to set Sort
func (options *ListWordsOptions) SetSort(sort string) *ListWordsOptions {
	options.Sort = core.StringPtr(sort)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListWordsOptions) SetHeaders(param map[string]string) *ListWordsOptions {
	options.Headers = param
	return options
}

// RecognitionJob : RecognitionJob struct
type RecognitionJob struct {

	// The ID of the asynchronous job.
	ID *string `json:"id" validate:"required"`

	// The current status of the job:
	// * `waiting`: The service is preparing the job for processing. The service returns this status when the job is
	// initially created or when it is waiting for capacity to process the job. The job remains in this state until the
	// service has the capacity to begin processing it.
	// * `processing`: The service is actively processing the job.
	// * `completed`: The service has finished processing the job. If the job specified a callback URL and the event
	// `recognitions.completed_with_results`, the service sent the results with the callback notification; otherwise, you
	// must retrieve the results by checking the individual job.
	// * `failed`: The job failed.
	Status *string `json:"status" validate:"required"`

	// The date and time in Coordinated Universal Time (UTC) at which the job was created. The value is provided in full
	// ISO 8601 format (`YYYY-MM-DDThh:mm:ss.sTZD`).
	Created *string `json:"created" validate:"required"`

	// The date and time in Coordinated Universal Time (UTC) at which the job was last updated by the service. The value is
	// provided in full ISO 8601 format (`YYYY-MM-DDThh:mm:ss.sTZD`). This field is returned only by the **Check jobs** and
	// **Check a job** methods.
	Updated *string `json:"updated,omitempty"`

	// The URL to use to request information about the job with the **Check a job** method. This field is returned only by
	// the **Create a job** method.
	URL *string `json:"url,omitempty"`

	// The user token associated with a job that was created with a callback URL and a user token. This field can be
	// returned only by the **Check jobs** method.
	UserToken *string `json:"user_token,omitempty"`

	// If the status is `completed`, the results of the recognition request as an array that includes a single instance of
	// a `SpeechRecognitionResults` object. This field is returned only by the **Check a job** method.
	Results []SpeechRecognitionResults `json:"results,omitempty"`

	// An array of warning messages about invalid parameters included with the request. Each warning includes a descriptive
	// message and a list of invalid argument strings, for example, `"unexpected query parameter 'user_token', query
	// parameter 'callback_url' was not specified"`. The request succeeds despite the warnings. This field can be returned
	// only by the **Create a job** method.
	Warnings []string `json:"warnings,omitempty"`
}

// RecognitionJobs : RecognitionJobs struct
type RecognitionJobs struct {

	// An array of objects that provides the status for each of the user's current jobs. The array is empty if the user has
	// no current jobs.
	Recognitions []RecognitionJob `json:"recognitions" validate:"required"`
}

// RecognizeOptions : The recognize options.
type RecognizeOptions struct {

	// The audio to transcribe in the format specified by the `Content-Type` header.
	Audio *io.ReadCloser `json:"audio,omitempty"`

	// The type of the input.
	ContentType *string `json:"Content-Type" validate:"required"`

	// The identifier of the model that is to be used for the recognition request.
	Model *string `json:"model,omitempty"`

	// The customization ID (GUID) of a custom language model that is to be used with the recognition request. The base
	// model of the specified custom language model must match the model specified with the `model` parameter. You must
	// make the request with service credentials created for the instance of the service that owns the custom model. By
	// default, no custom language model is used.
	CustomizationID *string `json:"customization_id,omitempty"`

	// The customization ID (GUID) of a custom acoustic model that is to be used with the recognition request. The base
	// model of the specified custom acoustic model must match the model specified with the `model` parameter. You must
	// make the request with service credentials created for the instance of the service that owns the custom model. By
	// default, no custom acoustic model is used.
	AcousticCustomizationID *string `json:"acoustic_customization_id,omitempty"`

	// The version of the specified base model that is to be used with recognition request. Multiple versions of a base
	// model can exist when a model is updated for internal improvements. The parameter is intended primarily for use with
	// custom models that have been upgraded for a new base model. The default value depends on whether the parameter is
	// used with or without a custom model. For more information, see [Base model
	// version](https://console.bluemix.net/docs/services/speech-to-text/input.html#version).
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
	CustomizationWeight *float64 `json:"customization_weight,omitempty"`

	// The time in seconds after which, if only silence (no speech) is detected in submitted audio, the connection is
	// closed with a 400 error. The parameter is useful for stopping audio submission from a live microphone when a user
	// simply walks away. Use `-1` for infinity.
	InactivityTimeout *int64 `json:"inactivity_timeout,omitempty"`

	// An array of keyword strings to spot in the audio. Each keyword string can include one or more string tokens.
	// Keywords are spotted only in the final results, not in interim hypotheses. If you specify any keywords, you must
	// also specify a keywords threshold. You can spot a maximum of 1000 keywords. Omit the parameter or specify an empty
	// array if you do not need to spot keywords.
	Keywords []string `json:"keywords,omitempty"`

	// A confidence value that is the lower bound for spotting a keyword. A word is considered to match a keyword if its
	// confidence is greater than or equal to the threshold. Specify a probability between 0.0 and 1.0. No keyword spotting
	// is performed if you omit the parameter. If you specify a threshold, you must also specify one or more keywords.
	KeywordsThreshold *float32 `json:"keywords_threshold,omitempty"`

	// The maximum number of alternative transcripts that the service is to return. By default, a single transcription is
	// returned.
	MaxAlternatives *int64 `json:"max_alternatives,omitempty"`

	// A confidence value that is the lower bound for identifying a hypothesis as a possible word alternative (also known
	// as "Confusion Networks"). An alternative word is considered if its confidence is greater than or equal to the
	// threshold. Specify a probability between 0.0 and 1.0. No alternative words are computed if you omit the parameter.
	WordAlternativesThreshold *float32 `json:"word_alternatives_threshold,omitempty"`

	// If `true`, the service returns a confidence measure in the range of 0.0 to 1.0 for each word. By default, no word
	// confidence measures are returned.
	WordConfidence *bool `json:"word_confidence,omitempty"`

	// If `true`, the service returns time alignment for each word. By default, no timestamps are returned.
	Timestamps *bool `json:"timestamps,omitempty"`

	// If `true`, the service filters profanity from all output except for keyword results by replacing inappropriate words
	// with a series of asterisks. Set the parameter to `false` to return results with no censoring. Applies to US English
	// transcription only.
	ProfanityFilter *bool `json:"profanity_filter,omitempty"`

	// If `true`, the service converts dates, times, series of digits and numbers, phone numbers, currency values, and
	// internet addresses into more readable, conventional representations in the final transcript of a recognition
	// request. For US English, the service also converts certain keyword strings to punctuation symbols. By default, no
	// smart formatting is performed. Applies to US English and Spanish transcription only.
	SmartFormatting *bool `json:"smart_formatting,omitempty"`

	// If `true`, the response includes labels that identify which words were spoken by which participants in a
	// multi-person exchange. By default, no speaker labels are returned. Setting `speaker_labels` to `true` forces the
	// `timestamps` parameter to be `true`, regardless of whether you specify `false` for the parameter.
	//
	//  To determine whether a language model supports speaker labels, use the **Get models** method and check that the
	// attribute `speaker_labels` is set to `true`. You can also refer to [Speaker
	// labels](https://console.bluemix.net/docs/services/speech-to-text/output.html#speaker_labels).
	SpeakerLabels *bool `json:"speaker_labels,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewRecognizeOptionsForBasic : Instantiate RecognizeOptionsForBasic
func (speechToText *SpeechToTextV1) NewRecognizeOptionsForBasic(audio io.ReadCloser) *RecognizeOptions {
	return &RecognizeOptions{
		Audio:       &audio,
		ContentType: core.StringPtr("audio/basic"),
	}
}

// NewRecognizeOptionsForFlac : Instantiate RecognizeOptionsForFlac
func (speechToText *SpeechToTextV1) NewRecognizeOptionsForFlac(audio io.ReadCloser) *RecognizeOptions {
	return &RecognizeOptions{
		Audio:       &audio,
		ContentType: core.StringPtr("audio/flac"),
	}
}

// NewRecognizeOptionsForL16 : Instantiate RecognizeOptionsForL16
func (speechToText *SpeechToTextV1) NewRecognizeOptionsForL16(audio io.ReadCloser) *RecognizeOptions {
	return &RecognizeOptions{
		Audio:       &audio,
		ContentType: core.StringPtr("audio/l16"),
	}
}

// NewRecognizeOptionsForMp3 : Instantiate RecognizeOptionsForMp3
func (speechToText *SpeechToTextV1) NewRecognizeOptionsForMp3(audio io.ReadCloser) *RecognizeOptions {
	return &RecognizeOptions{
		Audio:       &audio,
		ContentType: core.StringPtr("audio/mp3"),
	}
}

// NewRecognizeOptionsForMpeg : Instantiate RecognizeOptionsForMpeg
func (speechToText *SpeechToTextV1) NewRecognizeOptionsForMpeg(audio io.ReadCloser) *RecognizeOptions {
	return &RecognizeOptions{
		Audio:       &audio,
		ContentType: core.StringPtr("audio/mpeg"),
	}
}

// NewRecognizeOptionsForMulaw : Instantiate RecognizeOptionsForMulaw
func (speechToText *SpeechToTextV1) NewRecognizeOptionsForMulaw(audio io.ReadCloser) *RecognizeOptions {
	return &RecognizeOptions{
		Audio:       &audio,
		ContentType: core.StringPtr("audio/mulaw"),
	}
}

// NewRecognizeOptionsForOgg : Instantiate RecognizeOptionsForOgg
func (speechToText *SpeechToTextV1) NewRecognizeOptionsForOgg(audio io.ReadCloser) *RecognizeOptions {
	return &RecognizeOptions{
		Audio:       &audio,
		ContentType: core.StringPtr("audio/ogg"),
	}
}

// NewRecognizeOptionsForOggcodecsopus : Instantiate RecognizeOptionsForOggcodecsopus
func (speechToText *SpeechToTextV1) NewRecognizeOptionsForOggcodecsopus(audio io.ReadCloser) *RecognizeOptions {
	return &RecognizeOptions{
		Audio:       &audio,
		ContentType: core.StringPtr("audio/ogg;codecs=opus"),
	}
}

// NewRecognizeOptionsForOggcodecsvorbis : Instantiate RecognizeOptionsForOggcodecsvorbis
func (speechToText *SpeechToTextV1) NewRecognizeOptionsForOggcodecsvorbis(audio io.ReadCloser) *RecognizeOptions {
	return &RecognizeOptions{
		Audio:       &audio,
		ContentType: core.StringPtr("audio/ogg;codecs=vorbis"),
	}
}

// NewRecognizeOptionsForWav : Instantiate RecognizeOptionsForWav
func (speechToText *SpeechToTextV1) NewRecognizeOptionsForWav(audio io.ReadCloser) *RecognizeOptions {
	return &RecognizeOptions{
		Audio:       &audio,
		ContentType: core.StringPtr("audio/wav"),
	}
}

// NewRecognizeOptionsForWebm : Instantiate RecognizeOptionsForWebm
func (speechToText *SpeechToTextV1) NewRecognizeOptionsForWebm(audio io.ReadCloser) *RecognizeOptions {
	return &RecognizeOptions{
		Audio:       &audio,
		ContentType: core.StringPtr("audio/webm"),
	}
}

// NewRecognizeOptionsForWebmcodecsopus : Instantiate RecognizeOptionsForWebmcodecsopus
func (speechToText *SpeechToTextV1) NewRecognizeOptionsForWebmcodecsopus(audio io.ReadCloser) *RecognizeOptions {
	return &RecognizeOptions{
		Audio:       &audio,
		ContentType: core.StringPtr("audio/webm;codecs=opus"),
	}
}

// NewRecognizeOptionsForWebmcodecsvorbis : Instantiate RecognizeOptionsForWebmcodecsvorbis
func (speechToText *SpeechToTextV1) NewRecognizeOptionsForWebmcodecsvorbis(audio io.ReadCloser) *RecognizeOptions {
	return &RecognizeOptions{
		Audio:       &audio,
		ContentType: core.StringPtr("audio/webm;codecs=vorbis"),
	}
}

// SetAudio : Allow user to set Audio with the specified content type
func (options *RecognizeOptions) SetAudio(audio io.ReadCloser, contentType string) *RecognizeOptions {
	options.Audio = &audio
	options.ContentType = core.StringPtr(contentType)
	return options
}

// NewRecognizeOptions : Instantiate RecognizeOptions
func (speechToText *SpeechToTextV1) NewRecognizeOptions(contentType string) *RecognizeOptions {
	return &RecognizeOptions{
		ContentType: core.StringPtr(contentType),
	}
}

// SetModel : Allow user to set Model
func (options *RecognizeOptions) SetModel(model string) *RecognizeOptions {
	options.Model = core.StringPtr(model)
	return options
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *RecognizeOptions) SetCustomizationID(customizationID string) *RecognizeOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetAcousticCustomizationID : Allow user to set AcousticCustomizationID
func (options *RecognizeOptions) SetAcousticCustomizationID(acousticCustomizationID string) *RecognizeOptions {
	options.AcousticCustomizationID = core.StringPtr(acousticCustomizationID)
	return options
}

// SetBaseModelVersion : Allow user to set BaseModelVersion
func (options *RecognizeOptions) SetBaseModelVersion(baseModelVersion string) *RecognizeOptions {
	options.BaseModelVersion = core.StringPtr(baseModelVersion)
	return options
}

// SetCustomizationWeight : Allow user to set CustomizationWeight
func (options *RecognizeOptions) SetCustomizationWeight(customizationWeight float64) *RecognizeOptions {
	options.CustomizationWeight = core.Float64Ptr(customizationWeight)
	return options
}

// SetInactivityTimeout : Allow user to set InactivityTimeout
func (options *RecognizeOptions) SetInactivityTimeout(inactivityTimeout int64) *RecognizeOptions {
	options.InactivityTimeout = core.Int64Ptr(inactivityTimeout)
	return options
}

// SetKeywords : Allow user to set Keywords
func (options *RecognizeOptions) SetKeywords(keywords []string) *RecognizeOptions {
	options.Keywords = keywords
	return options
}

// SetKeywordsThreshold : Allow user to set KeywordsThreshold
func (options *RecognizeOptions) SetKeywordsThreshold(keywordsThreshold float32) *RecognizeOptions {
	options.KeywordsThreshold = core.Float32Ptr(keywordsThreshold)
	return options
}

// SetMaxAlternatives : Allow user to set MaxAlternatives
func (options *RecognizeOptions) SetMaxAlternatives(maxAlternatives int64) *RecognizeOptions {
	options.MaxAlternatives = core.Int64Ptr(maxAlternatives)
	return options
}

// SetWordAlternativesThreshold : Allow user to set WordAlternativesThreshold
func (options *RecognizeOptions) SetWordAlternativesThreshold(wordAlternativesThreshold float32) *RecognizeOptions {
	options.WordAlternativesThreshold = core.Float32Ptr(wordAlternativesThreshold)
	return options
}

// SetWordConfidence : Allow user to set WordConfidence
func (options *RecognizeOptions) SetWordConfidence(wordConfidence bool) *RecognizeOptions {
	options.WordConfidence = core.BoolPtr(wordConfidence)
	return options
}

// SetTimestamps : Allow user to set Timestamps
func (options *RecognizeOptions) SetTimestamps(timestamps bool) *RecognizeOptions {
	options.Timestamps = core.BoolPtr(timestamps)
	return options
}

// SetProfanityFilter : Allow user to set ProfanityFilter
func (options *RecognizeOptions) SetProfanityFilter(profanityFilter bool) *RecognizeOptions {
	options.ProfanityFilter = core.BoolPtr(profanityFilter)
	return options
}

// SetSmartFormatting : Allow user to set SmartFormatting
func (options *RecognizeOptions) SetSmartFormatting(smartFormatting bool) *RecognizeOptions {
	options.SmartFormatting = core.BoolPtr(smartFormatting)
	return options
}

// SetSpeakerLabels : Allow user to set SpeakerLabels
func (options *RecognizeOptions) SetSpeakerLabels(speakerLabels bool) *RecognizeOptions {
	options.SpeakerLabels = core.BoolPtr(speakerLabels)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *RecognizeOptions) SetHeaders(param map[string]string) *RecognizeOptions {
	options.Headers = param
	return options
}

// RegisterCallbackOptions : The registerCallback options.
type RegisterCallbackOptions struct {

	// An HTTP or HTTPS URL to which callback notifications are to be sent. To be white-listed, the URL must successfully
	// echo the challenge string during URL verification. During verification, the client can also check the signature that
	// the service sends in the `X-Callback-Signature` header to verify the origin of the request.
	CallbackURL *string `json:"callback_url" validate:"required"`

	// A user-specified string that the service uses to generate the HMAC-SHA1 signature that it sends via the
	// `X-Callback-Signature` header. The service includes the header during URL verification and with every notification
	// sent to the callback URL. It calculates the signature over the payload of the notification. If you omit the
	// parameter, the service does not send the header.
	UserSecret *string `json:"user_secret,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewRegisterCallbackOptions : Instantiate RegisterCallbackOptions
func (speechToText *SpeechToTextV1) NewRegisterCallbackOptions(callbackURL string) *RegisterCallbackOptions {
	return &RegisterCallbackOptions{
		CallbackURL: core.StringPtr(callbackURL),
	}
}

// SetCallbackURL : Allow user to set CallbackURL
func (options *RegisterCallbackOptions) SetCallbackURL(callbackURL string) *RegisterCallbackOptions {
	options.CallbackURL = core.StringPtr(callbackURL)
	return options
}

// SetUserSecret : Allow user to set UserSecret
func (options *RegisterCallbackOptions) SetUserSecret(userSecret string) *RegisterCallbackOptions {
	options.UserSecret = core.StringPtr(userSecret)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *RegisterCallbackOptions) SetHeaders(param map[string]string) *RegisterCallbackOptions {
	options.Headers = param
	return options
}

// RegisterStatus : RegisterStatus struct
type RegisterStatus struct {

	// The current status of the job:
	// * `created` if the callback URL was successfully white-listed as a result of the call.
	// * `already created` if the URL was already white-listed.
	Status *string `json:"status" validate:"required"`

	// The callback URL that is successfully registered.
	URL *string `json:"url" validate:"required"`
}

// ResetAcousticModelOptions : The resetAcousticModel options.
type ResetAcousticModelOptions struct {

	// The customization ID (GUID) of the custom acoustic model. You must make the request with service credentials created
	// for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewResetAcousticModelOptions : Instantiate ResetAcousticModelOptions
func (speechToText *SpeechToTextV1) NewResetAcousticModelOptions(customizationID string) *ResetAcousticModelOptions {
	return &ResetAcousticModelOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *ResetAcousticModelOptions) SetCustomizationID(customizationID string) *ResetAcousticModelOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ResetAcousticModelOptions) SetHeaders(param map[string]string) *ResetAcousticModelOptions {
	options.Headers = param
	return options
}

// ResetLanguageModelOptions : The resetLanguageModel options.
type ResetLanguageModelOptions struct {

	// The customization ID (GUID) of the custom language model. You must make the request with service credentials created
	// for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewResetLanguageModelOptions : Instantiate ResetLanguageModelOptions
func (speechToText *SpeechToTextV1) NewResetLanguageModelOptions(customizationID string) *ResetLanguageModelOptions {
	return &ResetLanguageModelOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *ResetLanguageModelOptions) SetCustomizationID(customizationID string) *ResetLanguageModelOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ResetLanguageModelOptions) SetHeaders(param map[string]string) *ResetLanguageModelOptions {
	options.Headers = param
	return options
}

// SpeakerLabelsResult : SpeakerLabelsResult struct
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
	FinalResults *bool `json:"final" validate:"required"`
}

// SpeechModel : SpeechModel struct
type SpeechModel struct {

	// The name of the model for use as an identifier in calls to the service (for example, `en-US_BroadbandModel`).
	Name *string `json:"name" validate:"required"`

	// The language identifier of the model (for example, `en-US`).
	Language *string `json:"language" validate:"required"`

	// The sampling rate (minimum acceptable rate for audio) used by the model in Hertz.
	Rate *int64 `json:"rate" validate:"required"`

	// The URI for the model.
	URL *string `json:"url" validate:"required"`

	// Describes the additional service features supported with the model.
	SupportedFeatures *SupportedFeatures `json:"supported_features" validate:"required"`

	// Brief description of the model.
	Description *string `json:"description" validate:"required"`
}

// SpeechModels : SpeechModels struct
type SpeechModels struct {

	// An array of objects that provides information about each available model.
	Models []SpeechModel `json:"models" validate:"required"`
}

// SpeechRecognitionAlternative : SpeechRecognitionAlternative struct
type SpeechRecognitionAlternative struct {

	// A transcription of the audio.
	Transcript *string `json:"transcript" validate:"required"`

	// A score that indicates the service's confidence in the transcript in the range of 0.0 to 1.0. Returned only for the
	// best alternative and only with results marked as final.
	Confidence *float64 `json:"confidence,omitempty"`

	// Time alignments for each word from the transcript as a list of lists. Each inner list consists of three elements:
	// the word followed by its start and end time in seconds, for example: `[["hello",0.0,1.2],["world",1.2,2.5]]`.
	// Returned only for the best alternative.
	Timestamps []string `json:"timestamps,omitempty"`

	// A confidence score for each word of the transcript as a list of lists. Each inner list consists of two elements: the
	// word and its confidence score in the range of 0.0 to 1.0, for example: `[["hello",0.95],["world",0.866]]`. Returned
	// only for the best alternative and only with results marked as final.
	WordConfidence []string `json:"word_confidence,omitempty"`
}

// SpeechRecognitionResult : SpeechRecognitionResult struct
type SpeechRecognitionResult struct {

	// An indication of whether the transcription results are final. If `true`, the results for this utterance are not
	// updated further; no additional results are sent for a `result_index` once its results are indicated as final.
	FinalResults *bool `json:"final" validate:"required"`

	// An array of alternative transcripts. The `alternatives` array can include additional requested output such as word
	// confidence or timestamps.
	Alternatives []SpeechRecognitionAlternative `json:"alternatives" validate:"required"`

	// A dictionary (or associative array) whose keys are the strings specified for `keywords` if both that parameter and
	// `keywords_threshold` are specified. A keyword for which no matches are found is omitted from the array. The array is
	// omitted if no matches are found for any keywords.
	KeywordsResult map[string][]KeywordResult `json:"keywords_result,omitempty"`

	// An array of alternative hypotheses found for words of the input audio if a `word_alternatives_threshold` is
	// specified.
	WordAlternatives []WordAlternativeResults `json:"word_alternatives,omitempty"`
}

// SpeechRecognitionResults : SpeechRecognitionResults struct
type SpeechRecognitionResults struct {

	// An array that can include interim and final results (interim results are returned only if supported by the method).
	// Final results are guaranteed not to change; interim results might be replaced by further interim results and final
	// results. The service periodically sends updates to the results list; the `result_index` is set to the lowest index
	// in the array that has changed; it is incremented for new results.
	Results []SpeechRecognitionResult `json:"results,omitempty"`

	// An index that indicates a change point in the `results` array. The service increments the index only for additional
	// results that it sends for new audio for the same request.
	ResultIndex *int64 `json:"result_index,omitempty"`

	// An array that identifies which words were spoken by which speakers in a multi-person exchange. Returned in the
	// response only if `speaker_labels` is `true`. When interim results are also requested for methods that support them,
	// it is possible for a `SpeechRecognitionResults` object to include only the `speaker_labels` field.
	SpeakerLabels []SpeakerLabelsResult `json:"speaker_labels,omitempty"`

	// An array of warning messages associated with the request:
	// * Warnings for invalid parameters or fields can include a descriptive message and a list of invalid argument
	// strings, for example, `"Unknown arguments:"` or `"Unknown url query arguments:"` followed by a list of the form
	// `"invalid_arg_1, invalid_arg_2."`
	// * The following warning is returned if the request passes a custom model that is based on an older version of a base
	// model for which an updated version is available: `"Using previous version of base model, because your custom model
	// has been built with it. Please note that this version will be supported only for a limited time. Consider updating
	// your custom model to the new base model. If you do not do that you will be automatically switched to base model when
	// you used the non-updated custom model."`
	//
	// In both cases, the request succeeds despite the warnings.
	Warnings []string `json:"warnings,omitempty"`
}

// SupportedFeatures : SupportedFeatures struct
type SupportedFeatures struct {

	// Indicates whether the customization interface can be used to create a custom language model based on the language
	// model.
	CustomLanguageModel *bool `json:"custom_language_model" validate:"required"`

	// Indicates whether the `speaker_labels` parameter can be used with the language model.
	SpeakerLabels *bool `json:"speaker_labels" validate:"required"`
}

// TrainAcousticModelOptions : The trainAcousticModel options.
type TrainAcousticModelOptions struct {

	// The customization ID (GUID) of the custom acoustic model. You must make the request with service credentials created
	// for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// The customization ID (GUID) of a custom language model that is to be used during training of the custom acoustic
	// model. Specify a custom language model that has been trained with verbatim transcriptions of the audio resources or
	// that contains words that are relevant to the contents of the audio resources.
	CustomLanguageModelID *string `json:"custom_language_model_id,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewTrainAcousticModelOptions : Instantiate TrainAcousticModelOptions
func (speechToText *SpeechToTextV1) NewTrainAcousticModelOptions(customizationID string) *TrainAcousticModelOptions {
	return &TrainAcousticModelOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *TrainAcousticModelOptions) SetCustomizationID(customizationID string) *TrainAcousticModelOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetCustomLanguageModelID : Allow user to set CustomLanguageModelID
func (options *TrainAcousticModelOptions) SetCustomLanguageModelID(customLanguageModelID string) *TrainAcousticModelOptions {
	options.CustomLanguageModelID = core.StringPtr(customLanguageModelID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *TrainAcousticModelOptions) SetHeaders(param map[string]string) *TrainAcousticModelOptions {
	options.Headers = param
	return options
}

// TrainLanguageModelOptions : The trainLanguageModel options.
type TrainLanguageModelOptions struct {

	// The customization ID (GUID) of the custom language model. You must make the request with service credentials created
	// for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// The type of words from the custom language model's words resource on which to train the model:
	// * `all` (the default) trains the model on all new words, regardless of whether they were extracted from corpora or
	// were added or modified by the user.
	// * `user` trains the model only on new words that were added or modified by the user; the model is not trained on new
	// words extracted from corpora.
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
	CustomizationWeight *float64 `json:"customization_weight,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewTrainLanguageModelOptions : Instantiate TrainLanguageModelOptions
func (speechToText *SpeechToTextV1) NewTrainLanguageModelOptions(customizationID string) *TrainLanguageModelOptions {
	return &TrainLanguageModelOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *TrainLanguageModelOptions) SetCustomizationID(customizationID string) *TrainLanguageModelOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetWordTypeToAdd : Allow user to set WordTypeToAdd
func (options *TrainLanguageModelOptions) SetWordTypeToAdd(wordTypeToAdd string) *TrainLanguageModelOptions {
	options.WordTypeToAdd = core.StringPtr(wordTypeToAdd)
	return options
}

// SetCustomizationWeight : Allow user to set CustomizationWeight
func (options *TrainLanguageModelOptions) SetCustomizationWeight(customizationWeight float64) *TrainLanguageModelOptions {
	options.CustomizationWeight = core.Float64Ptr(customizationWeight)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *TrainLanguageModelOptions) SetHeaders(param map[string]string) *TrainLanguageModelOptions {
	options.Headers = param
	return options
}

// UnregisterCallbackOptions : The unregisterCallback options.
type UnregisterCallbackOptions struct {

	// The callback URL that is to be unregistered.
	CallbackURL *string `json:"callback_url" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewUnregisterCallbackOptions : Instantiate UnregisterCallbackOptions
func (speechToText *SpeechToTextV1) NewUnregisterCallbackOptions(callbackURL string) *UnregisterCallbackOptions {
	return &UnregisterCallbackOptions{
		CallbackURL: core.StringPtr(callbackURL),
	}
}

// SetCallbackURL : Allow user to set CallbackURL
func (options *UnregisterCallbackOptions) SetCallbackURL(callbackURL string) *UnregisterCallbackOptions {
	options.CallbackURL = core.StringPtr(callbackURL)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UnregisterCallbackOptions) SetHeaders(param map[string]string) *UnregisterCallbackOptions {
	options.Headers = param
	return options
}

// UpgradeAcousticModelOptions : The upgradeAcousticModel options.
type UpgradeAcousticModelOptions struct {

	// The customization ID (GUID) of the custom acoustic model. You must make the request with service credentials created
	// for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// If the custom acoustic model was trained with a custom language model, the customization ID (GUID) of that custom
	// language model. The custom language model must be upgraded before the custom acoustic model can be upgraded.
	CustomLanguageModelID *string `json:"custom_language_model_id,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewUpgradeAcousticModelOptions : Instantiate UpgradeAcousticModelOptions
func (speechToText *SpeechToTextV1) NewUpgradeAcousticModelOptions(customizationID string) *UpgradeAcousticModelOptions {
	return &UpgradeAcousticModelOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *UpgradeAcousticModelOptions) SetCustomizationID(customizationID string) *UpgradeAcousticModelOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetCustomLanguageModelID : Allow user to set CustomLanguageModelID
func (options *UpgradeAcousticModelOptions) SetCustomLanguageModelID(customLanguageModelID string) *UpgradeAcousticModelOptions {
	options.CustomLanguageModelID = core.StringPtr(customLanguageModelID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpgradeAcousticModelOptions) SetHeaders(param map[string]string) *UpgradeAcousticModelOptions {
	options.Headers = param
	return options
}

// UpgradeLanguageModelOptions : The upgradeLanguageModel options.
type UpgradeLanguageModelOptions struct {

	// The customization ID (GUID) of the custom language model. You must make the request with service credentials created
	// for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewUpgradeLanguageModelOptions : Instantiate UpgradeLanguageModelOptions
func (speechToText *SpeechToTextV1) NewUpgradeLanguageModelOptions(customizationID string) *UpgradeLanguageModelOptions {
	return &UpgradeLanguageModelOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *UpgradeLanguageModelOptions) SetCustomizationID(customizationID string) *UpgradeLanguageModelOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpgradeLanguageModelOptions) SetHeaders(param map[string]string) *UpgradeLanguageModelOptions {
	options.Headers = param
	return options
}

// Word : Word struct
type Word struct {

	// A word from the custom model's words resource. The spelling of the word is used to train the model.
	Word *string `json:"word" validate:"required"`

	// An array of pronunciations for the word. The array can include the sounds-like pronunciation automatically generated
	// by the service if none is provided for the word; the service adds this pronunciation when it finishes processing the
	// word.
	SoundsLike []string `json:"sounds_like" validate:"required"`

	// The spelling of the word that the service uses to display the word in a transcript. The field contains an empty
	// string if no display-as value is provided for the word, in which case the word is displayed as it is spelled.
	DisplayAs *string `json:"display_as" validate:"required"`

	// A sum of the number of times the word is found across all corpora. For example, if the word occurs five times in one
	// corpus and seven times in another, its count is `12`. If you add a custom word to a model before it is added by any
	// corpora, the count begins at `1`; if the word is added from a corpus first and later modified, the count reflects
	// only the number of times it is found in corpora.
	Count *int64 `json:"count" validate:"required"`

	// An array of sources that describes how the word was added to the custom model's words resource. For OOV words added
	// from a corpus, includes the name of the corpus; if the word was added by multiple corpora, the names of all corpora
	// are listed. If the word was modified or added by the user directly, the field includes the string `user`.
	Source []string `json:"source" validate:"required"`

	// If the service discovered one or more problems that you need to correct for the word's definition, an array that
	// describes each of the errors.
	Error []WordError `json:"error,omitempty"`
}

// WordAlternativeResult : WordAlternativeResult struct
type WordAlternativeResult struct {

	// A confidence score for the word alternative hypothesis in the range of 0.0 to 1.0.
	Confidence *float64 `json:"confidence" validate:"required"`

	// An alternative hypothesis for a word from the input audio.
	Word *string `json:"word" validate:"required"`
}

// WordAlternativeResults : WordAlternativeResults struct
type WordAlternativeResults struct {

	// The start time in seconds of the word from the input audio that corresponds to the word alternatives.
	StartTime *float64 `json:"start_time" validate:"required"`

	// The end time in seconds of the word from the input audio that corresponds to the word alternatives.
	EndTime *float64 `json:"end_time" validate:"required"`

	// An array of alternative hypotheses for a word from the input audio.
	Alternatives []WordAlternativeResult `json:"alternatives" validate:"required"`
}

// WordError : WordError struct
type WordError struct {

	// A key-value pair that describes an error associated with the definition of a word in the words resource. Each pair
	// has the format `"element": "message"`, where `element` is the aspect of the definition that caused the problem and
	// `message` describes the problem. The following example describes a problem with one of the word's sounds-like
	// definitions: `"{sounds_like_string}": "Numbers are not allowed in sounds-like. You can try for example
	// '{suggested_string}'."` You must correct the error before you can train the model.
	Element *string `json:"element" validate:"required"`
}

// Words : Words struct
type Words struct {

	// An array of objects that provides information about each word in the custom model's words resource. The array is
	// empty if the custom model has no words.
	Words []Word `json:"words" validate:"required"`
}
