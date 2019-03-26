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
	"github.com/IBM/go-sdk-core/core"
	common "github.com/watson-developer-cloud/go-sdk/common"
	"io"
	"os"
	"strings"
)

// SpeechToTextV1 : The IBM&reg; Speech to Text service provides APIs that use IBM's speech-recognition capabilities to
// produce transcripts of spoken audio. The service can transcribe speech from various languages and audio formats. In
// addition to basic transcription, the service can produce detailed information about many different aspects of the
// audio. For most languages, the service supports two sampling rates, broadband and narrowband. It returns all JSON
// response content in the UTF-8 character set.
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
// Language model customization is generally available for production use with most supported languages. Acoustic model
// customization is beta functionality that is available for all supported languages.
//
// Version: V1
// See: http://www.ibm.com/watson/developercloud/speech-to-text.html
type SpeechToTextV1 struct {
	Service *core.BaseService
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
	service, serviceErr := core.NewBaseService(serviceOptions, "speech_to_text", "Speech to Text")
	if serviceErr != nil {
		return nil, serviceErr
	}

	return &SpeechToTextV1{Service: service}, nil
}

// GetModel : Get a model
// Gets information for a single specified language model that is available for use with the service. The information
// includes the name of the model and its minimum sampling rate in Hertz, among other things.
//
// **See also:** [Languages and models](https://cloud.ibm.com/docs/services/speech-to-text/models.html).
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
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

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
		return nil, err
	}

	response, err := speechToText.Service.Request(request, new(SpeechModel))
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
// Lists all language models that are available for use with the service. The information includes the name of the model
// and its minimum sampling rate in Hertz, among other things.
//
// **See also:** [Languages and models](https://cloud.ibm.com/docs/services/speech-to-text/models.html).
func (speechToText *SpeechToTextV1) ListModels(listModelsOptions *ListModelsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateStruct(listModelsOptions, "listModelsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/models"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

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
		return nil, err
	}

	response, err := speechToText.Service.Request(request, new(SpeechModels))
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
// Sends audio and returns transcription results for a recognition request. You can pass a maximum of 100 MB and a
// minimum of 100 bytes of audio with a request. The service automatically detects the endianness of the incoming audio
// and, for audio that includes multiple channels, downmixes the audio to one-channel mono during transcoding. The
// method returns only final results; to enable interim results, use the WebSocket API.
//
// **See also:** [Making a basic HTTP request](https://cloud.ibm.com/docs/services/speech-to-text/http.html#HTTP-basic).
//
//
// ### Streaming mode
//
//  For requests to transcribe live audio as it becomes available, you must set the `Transfer-Encoding` header to
// `chunked` to use streaming mode. In streaming mode, the service closes the connection (status code 408) if it does
// not receive at least 15 seconds of audio (including silence) in any 30-second period. The service also closes the
// connection (status code 400) if it detects no speech for `inactivity_timeout` seconds of streaming audio; use the
// `inactivity_timeout` parameter to change the default of 30 seconds.
//
// **See also:**
// * [Audio transmission](https://cloud.ibm.com/docs/services/speech-to-text/input.html#transmission)
// * [Timeouts](https://cloud.ibm.com/docs/services/speech-to-text/input.html#timeouts)
//
// ### Audio formats (content types)
//
//  The service accepts audio in the following formats (MIME types).
// * For formats that are labeled **Required**, you must use the `Content-Type` header with the request to specify the
// format of the audio.
// * For all other formats, you can omit the `Content-Type` header or specify `application/octet-stream` with the header
// to have the service automatically detect the format of the audio. (With the `curl` command, you can specify either
// `\"Content-Type:\"` or `\"Content-Type: application/octet-stream\"`.)
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
//  **See also:** [Audio formats](https://cloud.ibm.com/docs/services/speech-to-text/audio-formats.html).
//
// ### Multipart speech recognition
//
//  **Note:** The Watson SDKs do not support multipart speech recognition.
//
// The HTTP `POST` method of the service also supports multipart speech recognition. With multipart requests, you pass
// all audio data as multipart form data. You specify some parameters as request headers and query parameters, but you
// pass JSON metadata as form data to control most aspects of the transcription.
//
// The multipart approach is intended for use with browsers for which JavaScript is disabled or when the parameters used
// with the request are greater than the 8 KB limit imposed by most HTTP servers and proxies. You can encounter this
// limit, for example, if you want to spot a very large number of keywords.
//
// **See also:** [Making a multipart HTTP
// request](https://cloud.ibm.com/docs/services/speech-to-text/http.html#HTTP-multi).
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
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

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

	_, err := builder.SetBodyContent(core.StringNilMapper(recognizeOptions.ContentType), nil, nil, recognizeOptions.Audio)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.Service.Request(request, new(SpeechRecognitionResults))
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
// Returns information about the specified job. The response always includes the status of the job and its creation and
// update times. If the status is `completed`, the response includes the results of the recognition request. You must
// use credentials for the instance of the service that owns a job to list information about it.
//
// You can use the method to retrieve the results of any job, regardless of whether it was submitted with a callback URL
// and the `recognitions.completed_with_results` event, and you can retrieve the results multiple times for as long as
// they remain available. Use the **Check jobs** method to request information about the most recent jobs associated
// with the calling credentials.
//
// **See also:** [Checking the status and retrieving the results of a
// job](https://cloud.ibm.com/docs/services/speech-to-text/async.html#job).
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
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

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
		return nil, err
	}

	response, err := speechToText.Service.Request(request, new(RecognitionJob))
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
// Returns the ID and status of the latest 100 outstanding jobs associated with the credentials with which it is called.
// The method also returns the creation and update times of each job, and, if a job was created with a callback URL and
// a user token, the user token for the job. To obtain the results for a job whose status is `completed` or not one of
// the latest 100 outstanding jobs, use the **Check a job** method. A job and its results remain available until you
// delete them with the **Delete a job** method or until the job's time to live expires, whichever comes first.
//
// **See also:** [Checking the status of the latest
// jobs](https://cloud.ibm.com/docs/services/speech-to-text/async.html#jobs).
func (speechToText *SpeechToTextV1) CheckJobs(checkJobsOptions *CheckJobsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateStruct(checkJobsOptions, "checkJobsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/recognitions"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

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
		return nil, err
	}

	response, err := speechToText.Service.Request(request, new(RecognitionJobs))
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
// Creates a job for a new asynchronous recognition request. The job is owned by the instance of the service whose
// credentials are used to create it. How you learn the status and results of a job depends on the parameters you
// include with the job creation request:
// * By callback notification: Include the `callback_url` parameter to specify a URL to which the service is to send
// callback notifications when the status of the job changes. Optionally, you can also include the `events` and
// `user_token` parameters to subscribe to specific events and to specify a string that is to be included with each
// notification for the job.
// * By polling the service: Omit the `callback_url`, `events`, and `user_token` parameters. You must then use the
// **Check jobs** or **Check a job** methods to check the status of the job, using the latter to retrieve the results
// when the job is complete.
//
// The two approaches are not mutually exclusive. You can poll the service for job status or obtain results from the
// service manually even if you include a callback URL. In both cases, you can include the `results_ttl` parameter to
// specify how long the results are to remain available after the job is complete. Using the HTTPS **Check a job**
// method to retrieve results is more secure than receiving them via callback notification over HTTP because it provides
// confidentiality in addition to authentication and data integrity.
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
// WebSocket API.
//
// **See also:** [Creating a job](https://cloud.ibm.com/docs/services/speech-to-text/async.html#create).
//
// ### Streaming mode
//
//  For requests to transcribe live audio as it becomes available, you must set the `Transfer-Encoding` header to
// `chunked` to use streaming mode. In streaming mode, the service closes the connection (status code 408) if it does
// not receive at least 15 seconds of audio (including silence) in any 30-second period. The service also closes the
// connection (status code 400) if it detects no speech for `inactivity_timeout` seconds of streaming audio; use the
// `inactivity_timeout` parameter to change the default of 30 seconds.
//
// **See also:**
// * [Audio transmission](https://cloud.ibm.com/docs/services/speech-to-text/input.html#transmission)
// * [Timeouts](https://cloud.ibm.com/docs/services/speech-to-text/input.html#timeouts)
//
// ### Audio formats (content types)
//
//  The service accepts audio in the following formats (MIME types).
// * For formats that are labeled **Required**, you must use the `Content-Type` header with the request to specify the
// format of the audio.
// * For all other formats, you can omit the `Content-Type` header or specify `application/octet-stream` with the header
// to have the service automatically detect the format of the audio. (With the `curl` command, you can specify either
// `\"Content-Type:\"` or `\"Content-Type: application/octet-stream\"`.)
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
//  **See also:** [Audio formats](https://cloud.ibm.com/docs/services/speech-to-text/audio-formats.html).
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
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

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

	_, err := builder.SetBodyContent(core.StringNilMapper(createJobOptions.ContentType), nil, nil, createJobOptions.Audio)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.Service.Request(request, new(RecognitionJob))
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
// Deletes the specified job. You cannot delete a job that the service is actively processing. Once you delete a job,
// its results are no longer available. The service automatically deletes a job and its results when the time to live
// for the results expires. You must use credentials for the instance of the service that owns a job to delete it.
//
// **See also:** [Deleting a job](https://cloud.ibm.com/docs/services/speech-to-text/async.html#delete-async).
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
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteJobOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "DeleteJob")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "")

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.Service.Request(request, nil)
	return response, err
}

// RegisterCallback : Register a callback
// Registers a callback URL with the service for use with subsequent asynchronous recognition requests. The service
// attempts to register, or white-list, the callback URL if it is not already registered by sending a `GET` request to
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
// white-list the URL; it instead sends status code 400 in response to the **Register a callback** request. If the
// requested callback URL is already white-listed, the service responds to the initial registration request with
// response code 200.
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
// **See also:** [Registering a callback URL](https://cloud.ibm.com/docs/services/speech-to-text/async.html#register).
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
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

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
		return nil, err
	}

	response, err := speechToText.Service.Request(request, new(RegisterStatus))
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
// Unregisters a callback URL that was previously white-listed with a **Register a callback** request for use with the
// asynchronous interface. Once unregistered, the URL can no longer be used with asynchronous recognition requests.
//
// **See also:** [Unregistering a callback
// URL](https://cloud.ibm.com/docs/services/speech-to-text/async.html#unregister).
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
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range unregisterCallbackOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "UnregisterCallback")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "")

	builder.AddQuery("callback_url", fmt.Sprint(*unregisterCallbackOptions.CallbackURL))

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.Service.Request(request, nil)
	return response, err
}

// CreateLanguageModel : Create a custom language model
// Creates a new custom language model for a specified base model. The custom language model can be used only with the
// base model for which it is created. The model is owned by the instance of the service whose credentials are used to
// create it.
//
// **See also:** [Create a custom language
// model](https://cloud.ibm.com/docs/services/speech-to-text/language-create.html#createModel-language).
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
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

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
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.Service.Request(request, new(LanguageModel))
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
// Deletes an existing custom language model. The custom model cannot be deleted if another request, such as adding a
// corpus or grammar to the model, is currently being processed. You must use credentials for the instance of the
// service that owns a model to delete it.
//
// **See also:** [Deleting a custom language
// model](https://cloud.ibm.com/docs/services/speech-to-text/language-models.html#deleteModel-language).
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
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

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
		return nil, err
	}

	response, err := speechToText.Service.Request(request, nil)
	return response, err
}

// GetLanguageModel : Get a custom language model
// Gets information about a specified custom language model. You must use credentials for the instance of the service
// that owns a model to list information about it.
//
// **See also:** [Listing custom language
// models](https://cloud.ibm.com/docs/services/speech-to-text/language-models.html#listModels-language).
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
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

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
		return nil, err
	}

	response, err := speechToText.Service.Request(request, new(LanguageModel))
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
// Lists information about all custom language models that are owned by an instance of the service. Use the `language`
// parameter to see all custom language models for the specified language. Omit the parameter to see all custom language
// models for all languages. You must use credentials for the instance of the service that owns a model to list
// information about it.
//
// **See also:** [Listing custom language
// models](https://cloud.ibm.com/docs/services/speech-to-text/language-models.html#listModels-language).
func (speechToText *SpeechToTextV1) ListLanguageModels(listLanguageModelsOptions *ListLanguageModelsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateStruct(listLanguageModelsOptions, "listLanguageModelsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/customizations"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

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
		return nil, err
	}

	response, err := speechToText.Service.Request(request, new(LanguageModels))
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
// Resets a custom language model by removing all corpora, grammars, and words from the model. Resetting a custom
// language model initializes the model to its state when it was first created. Metadata such as the name and language
// of the model are preserved, but the model's words resource is removed and must be re-created. You must use
// credentials for the instance of the service that owns a model to reset it.
//
// **See also:** [Resetting a custom language
// model](https://cloud.ibm.com/docs/services/speech-to-text/language-models.html#resetModel-language).
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
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

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
		return nil, err
	}

	response, err := speechToText.Service.Request(request, nil)
	return response, err
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
// You can monitor the status of the training by using the **Get a custom language model** method to poll the model's
// status. Use a loop to check the status every 10 seconds. The method returns a `LanguageModel` object that includes
// `status` and `progress` fields. A status of `available` means that the custom model is trained and ready to use. The
// service cannot accept subsequent training requests or requests to add new resources until the existing request
// completes.
//
// Training can fail to start for the following reasons:
// * The service is currently handling another request for the custom model, such as another training request or a
// request to add a corpus or grammar to the model.
// * No training data have been added to the custom model.
// * One or more words that were added to the custom model have invalid sounds-like pronunciations that you must fix.
//
// **See also:** [Train the custom language
// model](https://cloud.ibm.com/docs/services/speech-to-text/language-create.html#trainModel-language).
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
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

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
		return nil, err
	}

	response, err := speechToText.Service.Request(request, nil)
	return response, err
}

// UpgradeLanguageModel : Upgrade a custom language model
// Initiates the upgrade of a custom language model to the latest version of its base language model. The upgrade method
// is asynchronous. It can take on the order of minutes to complete depending on the amount of data in the custom model
// and the current load on the service. A custom model must be in the `ready` or `available` state to be upgraded. You
// must use credentials for the instance of the service that owns a model to upgrade it.
//
// The method returns an HTTP 200 response code to indicate that the upgrade process has begun successfully. You can
// monitor the status of the upgrade by using the **Get a custom language model** method to poll the model's status. The
// method returns a `LanguageModel` object that includes `status` and `progress` fields. Use a loop to check the status
// every 10 seconds. While it is being upgraded, the custom model has the status `upgrading`. When the upgrade is
// complete, the model resumes the status that it had prior to upgrade. The service cannot accept subsequent requests
// for the model until the upgrade completes.
//
// **See also:** [Upgrading a custom language
// model](https://cloud.ibm.com/docs/services/speech-to-text/custom-upgrade.html#upgradeLanguage).
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
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

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
		return nil, err
	}

	response, err := speechToText.Service.Request(request, nil)
	return response, err
}

// AddCorpus : Add a corpus
// Adds a single corpus text file of new training data to a custom language model. Use multiple requests to submit
// multiple corpus text files. You must use credentials for the instance of the service that owns a model to add a
// corpus to it. Adding a corpus does not affect the custom language model until you train the model for the new data by
// using the **Train a custom language model** method.
//
// Submit a plain text file that contains sample sentences from the domain of interest to enable the service to extract
// words in context. The more sentences you add that represent the context in which speakers use words from the domain,
// the better the service's recognition accuracy.
//
// The call returns an HTTP 201 response code if the corpus is valid. The service then asynchronously processes the
// contents of the corpus and automatically extracts new words that it finds. This can take on the order of a minute or
// two to complete depending on the total number of words and the number of new words in the corpus, as well as the
// current load on the service. You cannot submit requests to add additional resources to the custom model or to train
// the model until the service's analysis of the corpus for the current request completes. Use the **List a corpus**
// method to check the status of the analysis.
//
// The service auto-populates the model's words resource with words from the corpus that are not found in its base
// vocabulary. These are referred to as out-of-vocabulary (OOV) words. You can use the **List custom words** method to
// examine the words resource. You can use other words method to eliminate typos and modify how words are pronounced as
// needed.
//
// To add a corpus file that has the same name as an existing corpus, set the `allow_overwrite` parameter to `true`;
// otherwise, the request fails. Overwriting an existing corpus causes the service to process the corpus text file and
// extract OOV words anew. Before doing so, it removes any OOV words associated with the existing corpus from the
// model's words resource unless they were also added by another corpus or grammar, or they have been modified in some
// way with the **Add custom words** or **Add a custom word** method.
//
// The service limits the overall amount of data that you can add to a custom model to a maximum of 10 million total
// words from all sources combined. Also, you can add no more than 30 thousand custom (OOV) words to a model. This
// includes words that the service extracts from corpora and grammars, and words that you add directly.
//
// **See also:**
// * [Working with corpora](https://cloud.ibm.com/docs/services/speech-to-text/language-resource.html#workingCorpora)
// * [Add corpora to the custom language
// model](https://cloud.ibm.com/docs/services/speech-to-text/language-create.html#addCorpora).
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
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

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
		return nil, err
	}

	response, err := speechToText.Service.Request(request, nil)
	return response, err
}

// DeleteCorpus : Delete a corpus
// Deletes an existing corpus from a custom language model. The service removes any out-of-vocabulary (OOV) words that
// are associated with the corpus from the custom model's words resource unless they were also added by another corpus
// or grammar, or they were modified in some way with the **Add custom words** or **Add a custom word** method. Removing
// a corpus does not affect the custom model until you train the model with the **Train a custom language model**
// method. You must use credentials for the instance of the service that owns a model to delete its corpora.
//
// **See also:** [Deleting a corpus from a custom language
// model](https://cloud.ibm.com/docs/services/speech-to-text/language-corpora.html#deleteCorpus).
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
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

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
		return nil, err
	}

	response, err := speechToText.Service.Request(request, nil)
	return response, err
}

// GetCorpus : Get a corpus
// Gets information about a corpus from a custom language model. The information includes the total number of words and
// out-of-vocabulary (OOV) words, name, and status of the corpus. You must use credentials for the instance of the
// service that owns a model to list its corpora.
//
// **See also:** [Listing corpora for a custom language
// model](https://cloud.ibm.com/docs/services/speech-to-text/language-corpora.html#listCorpora).
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
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

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
		return nil, err
	}

	response, err := speechToText.Service.Request(request, new(Corpus))
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
// Lists information about all corpora from a custom language model. The information includes the total number of words
// and out-of-vocabulary (OOV) words, name, and status of each corpus. You must use credentials for the instance of the
// service that owns a model to list its corpora.
//
// **See also:** [Listing corpora for a custom language
// model](https://cloud.ibm.com/docs/services/speech-to-text/language-corpora.html#listCorpora).
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
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

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
		return nil, err
	}

	response, err := speechToText.Service.Request(request, new(Corpora))
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
// Adds a custom word to a custom language model. The service populates the words resource for a custom model with
// out-of-vocabulary (OOV) words from each corpus or grammar that is added to the model. You can use this method to add
// a word or to modify an existing word in the words resource. The words resource for a model can contain a maximum of
// 30 thousand custom (OOV) words. This includes words that the service extracts from corpora and grammars and words
// that you add directly.
//
// You must use credentials for the instance of the service that owns a model to add or modify a custom word for the
// model. Adding or modifying a custom word does not affect the custom model until you train the model for the new data
// by using the **Train a custom language model** method.
//
// Use the `word_name` parameter to specify the custom word that is to be added or modified. Use the `CustomWord` object
// to provide one or both of the optional `sounds_like` and `display_as` fields for the word.
// * The `sounds_like` field provides an array of one or more pronunciations for the word. Use the parameter to specify
// how the word can be pronounced by users. Use the parameter for words that are difficult to pronounce, foreign words,
// acronyms, and so on. For example, you might specify that the word `IEEE` can sound like `i triple e`. You can specify
// a maximum of five sounds-like pronunciations for a word.
// * The `display_as` field provides a different way of spelling the word in a transcript. Use the parameter when you
// want the word to appear different from its usual representation or from its spelling in training data. For example,
// you might indicate that the word `IBM(trademark)` is to be displayed as `IBM&trade;`.
//
// If you add a custom word that already exists in the words resource for the custom model, the new definition
// overwrites the existing data for the word. If the service encounters an error, it does not add the word to the words
// resource. Use the **List a custom word** method to review the word that you add.
//
// **See also:**
// * [Working with custom words](https://cloud.ibm.com/docs/services/speech-to-text/language-resource.html#workingWords)
// * [Add words to the custom language
// model](https://cloud.ibm.com/docs/services/speech-to-text/language-create.html#addWords).
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
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

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
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.Service.Request(request, nil)
	return response, err
}

// AddWords : Add custom words
// Adds one or more custom words to a custom language model. The service populates the words resource for a custom model
// with out-of-vocabulary (OOV) words from each corpus or grammar that is added to the model. You can use this method to
// add additional words or to modify existing words in the words resource. The words resource for a model can contain a
// maximum of 30 thousand custom (OOV) words. This includes words that the service extracts from corpora and grammars
// and words that you add directly.
//
// You must use credentials for the instance of the service that owns a model to add or modify custom words for the
// model. Adding or modifying custom words does not affect the custom model until you train the model for the new data
// by using the **Train a custom language model** method.
//
// You add custom words by providing a `CustomWords` object, which is an array of `CustomWord` objects, one per word.
// You must use the object's `word` parameter to identify the word that is to be added. You can also provide one or both
// of the optional `sounds_like` and `display_as` fields for each word.
// * The `sounds_like` field provides an array of one or more pronunciations for the word. Use the parameter to specify
// how the word can be pronounced by users. Use the parameter for words that are difficult to pronounce, foreign words,
// acronyms, and so on. For example, you might specify that the word `IEEE` can sound like `i triple e`. You can specify
// a maximum of five sounds-like pronunciations for a word.
// * The `display_as` field provides a different way of spelling the word in a transcript. Use the parameter when you
// want the word to appear different from its usual representation or from its spelling in training data. For example,
// you might indicate that the word `IBM(trademark)` is to be displayed as `IBM&trade;`.
//
// If you add a custom word that already exists in the words resource for the custom model, the new definition
// overwrites the existing data for the word. If the service encounters an error with the input data, it returns a
// failure code and does not add any of the words to the words resource.
//
// The call returns an HTTP 201 response code if the input data is valid. It then asynchronously processes the words to
// add them to the model's words resource. The time that it takes for the analysis to complete depends on the number of
// new words that you add but is generally faster than adding a corpus or grammar.
//
// You can monitor the status of the request by using the **List a custom language model** method to poll the model's
// status. Use a loop to check the status every 10 seconds. The method returns a `Customization` object that includes a
// `status` field. A status of `ready` means that the words have been added to the custom model. The service cannot
// accept requests to add new data or to train the model until the existing request completes.
//
// You can use the **List custom words** or **List a custom word** method to review the words that you add. Words with
// an invalid `sounds_like` field include an `error` field that describes the problem. You can use other words-related
// methods to correct errors, eliminate typos, and modify how words are pronounced as needed.
//
// **See also:**
// * [Working with custom words](https://cloud.ibm.com/docs/services/speech-to-text/language-resource.html#workingWords)
// * [Add words to the custom language
// model](https://cloud.ibm.com/docs/services/speech-to-text/language-create.html#addWords).
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
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

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
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.Service.Request(request, nil)
	return response, err
}

// DeleteWord : Delete a custom word
// Deletes a custom word from a custom language model. You can remove any word that you added to the custom model's
// words resource via any means. However, if the word also exists in the service's base vocabulary, the service removes
// only the custom pronunciation for the word; the word remains in the base vocabulary. Removing a custom word does not
// affect the custom model until you train the model with the **Train a custom language model** method. You must use
// credentials for the instance of the service that owns a model to delete its words.
//
// **See also:** [Deleting a word from a custom language
// model](https://cloud.ibm.com/docs/services/speech-to-text/language-words.html#deleteWord).
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
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

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
		return nil, err
	}

	response, err := speechToText.Service.Request(request, nil)
	return response, err
}

// GetWord : Get a custom word
// Gets information about a custom word from a custom language model. You must use credentials for the instance of the
// service that owns a model to list information about its words.
//
// **See also:** [Listing words from a custom language
// model](https://cloud.ibm.com/docs/services/speech-to-text/language-words.html#listWords).
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
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

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
		return nil, err
	}

	response, err := speechToText.Service.Request(request, new(Word))
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
// Lists information about custom words from a custom language model. You can list all words from the custom model's
// words resource, only custom words that were added or modified by the user, or only out-of-vocabulary (OOV) words that
// were extracted from corpora or are recognized by grammars. You can also indicate the order in which the service is to
// return words; by default, the service lists words in ascending alphabetical order. You must use credentials for the
// instance of the service that owns a model to list information about its words.
//
// **See also:** [Listing words from a custom language
// model](https://cloud.ibm.com/docs/services/speech-to-text/language-words.html#listWords).
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
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

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
		return nil, err
	}

	response, err := speechToText.Service.Request(request, new(Words))
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

// AddGrammar : Add a grammar
// Adds a single grammar file to a custom language model. Submit a plain text file in UTF-8 format that defines the
// grammar. Use multiple requests to submit multiple grammar files. You must use credentials for the instance of the
// service that owns a model to add a grammar to it. Adding a grammar does not affect the custom language model until
// you train the model for the new data by using the **Train a custom language model** method.
//
// The call returns an HTTP 201 response code if the grammar is valid. The service then asynchronously processes the
// contents of the grammar and automatically extracts new words that it finds. This can take a few seconds to complete
// depending on the size and complexity of the grammar, as well as the current load on the service. You cannot submit
// requests to add additional resources to the custom model or to train the model until the service's analysis of the
// grammar for the current request completes. Use the **Get a grammar** method to check the status of the analysis.
//
// The service populates the model's words resource with any word that is recognized by the grammar that is not found in
// the model's base vocabulary. These are referred to as out-of-vocabulary (OOV) words. You can use the **List custom
// words** method to examine the words resource and use other words-related methods to eliminate typos and modify how
// words are pronounced as needed.
//
// To add a grammar that has the same name as an existing grammar, set the `allow_overwrite` parameter to `true`;
// otherwise, the request fails. Overwriting an existing grammar causes the service to process the grammar file and
// extract OOV words anew. Before doing so, it removes any OOV words associated with the existing grammar from the
// model's words resource unless they were also added by another resource or they have been modified in some way with
// the **Add custom words** or **Add a custom word** method.
//
// The service limits the overall amount of data that you can add to a custom model to a maximum of 10 million total
// words from all sources combined. Also, you can add no more than 30 thousand OOV words to a model. This includes words
// that the service extracts from corpora and grammars and words that you add directly.
//
// **See also:**
// * [Working with grammars](https://cloud.ibm.com/docs/services/speech-to-text/)
// * [Add grammars to the custom language model](https://cloud.ibm.com/docs/services/speech-to-text/).
func (speechToText *SpeechToTextV1) AddGrammar(addGrammarOptions *AddGrammarOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(addGrammarOptions, "addGrammarOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(addGrammarOptions, "addGrammarOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/customizations", "grammars"}
	pathParameters := []string{*addGrammarOptions.CustomizationID, *addGrammarOptions.GrammarName}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

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

	_, err := builder.SetBodyContent(core.StringNilMapper(addGrammarOptions.ContentType), nil, nil, addGrammarOptions.GrammarFile)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.Service.Request(request, nil)
	return response, err
}

// DeleteGrammar : Delete a grammar
// Deletes an existing grammar from a custom language model. The service removes any out-of-vocabulary (OOV) words
// associated with the grammar from the custom model's words resource unless they were also added by another resource or
// they were modified in some way with the **Add custom words** or **Add a custom word** method. Removing a grammar does
// not affect the custom model until you train the model with the **Train a custom language model** method. You must use
// credentials for the instance of the service that owns a model to delete its grammar.
//
// **See also:** [Deleting a grammar from a custom language model](https://cloud.ibm.com/docs/services/speech-to-text/).
func (speechToText *SpeechToTextV1) DeleteGrammar(deleteGrammarOptions *DeleteGrammarOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(deleteGrammarOptions, "deleteGrammarOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(deleteGrammarOptions, "deleteGrammarOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/customizations", "grammars"}
	pathParameters := []string{*deleteGrammarOptions.CustomizationID, *deleteGrammarOptions.GrammarName}

	builder := core.NewRequestBuilder(core.DELETE)
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

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
		return nil, err
	}

	response, err := speechToText.Service.Request(request, nil)
	return response, err
}

// GetGrammar : Get a grammar
// Gets information about a grammar from a custom language model. The information includes the total number of
// out-of-vocabulary (OOV) words, name, and status of the grammar. You must use credentials for the instance of the
// service that owns a model to list its grammars.
//
// **See also:** [Listing grammars from a custom language model](https://cloud.ibm.com/docs/services/speech-to-text/).
func (speechToText *SpeechToTextV1) GetGrammar(getGrammarOptions *GetGrammarOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(getGrammarOptions, "getGrammarOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(getGrammarOptions, "getGrammarOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/customizations", "grammars"}
	pathParameters := []string{*getGrammarOptions.CustomizationID, *getGrammarOptions.GrammarName}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

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
		return nil, err
	}

	response, err := speechToText.Service.Request(request, new(Grammar))
	return response, err
}

// GetGetGrammarResult : Retrieve result of GetGrammar operation
func (speechToText *SpeechToTextV1) GetGetGrammarResult(response *core.DetailedResponse) *Grammar {
	result, ok := response.Result.(*Grammar)
	if ok {
		return result
	}
	return nil
}

// ListGrammars : List grammars
// Lists information about all grammars from a custom language model. The information includes the total number of
// out-of-vocabulary (OOV) words, name, and status of each grammar. You must use credentials for the instance of the
// service that owns a model to list its grammars.
//
// **See also:** [Listing grammars from a custom language model](https://cloud.ibm.com/docs/services/speech-to-text/).
func (speechToText *SpeechToTextV1) ListGrammars(listGrammarsOptions *ListGrammarsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(listGrammarsOptions, "listGrammarsOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(listGrammarsOptions, "listGrammarsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/customizations", "grammars"}
	pathParameters := []string{*listGrammarsOptions.CustomizationID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

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
		return nil, err
	}

	response, err := speechToText.Service.Request(request, new(Grammars))
	return response, err
}

// GetListGrammarsResult : Retrieve result of ListGrammars operation
func (speechToText *SpeechToTextV1) GetListGrammarsResult(response *core.DetailedResponse) *Grammars {
	result, ok := response.Result.(*Grammars)
	if ok {
		return result
	}
	return nil
}

// CreateAcousticModel : Create a custom acoustic model
// Creates a new custom acoustic model for a specified base model. The custom acoustic model can be used only with the
// base model for which it is created. The model is owned by the instance of the service whose credentials are used to
// create it.
//
// **See also:** [Create a custom acoustic
// model](https://cloud.ibm.com/docs/services/speech-to-text/acoustic-create.html#createModel-acoustic).
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
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

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
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.Service.Request(request, new(AcousticModel))
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
// Deletes an existing custom acoustic model. The custom model cannot be deleted if another request, such as adding an
// audio resource to the model, is currently being processed. You must use credentials for the instance of the service
// that owns a model to delete it.
//
// **See also:** [Deleting a custom acoustic
// model](https://cloud.ibm.com/docs/services/speech-to-text/acoustic-models.html#deleteModel-acoustic).
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
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

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
		return nil, err
	}

	response, err := speechToText.Service.Request(request, nil)
	return response, err
}

// GetAcousticModel : Get a custom acoustic model
// Gets information about a specified custom acoustic model. You must use credentials for the instance of the service
// that owns a model to list information about it.
//
// **See also:** [Listing custom acoustic
// models](https://cloud.ibm.com/docs/services/speech-to-text/acoustic-models.html#listModels-acoustic).
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
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

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
		return nil, err
	}

	response, err := speechToText.Service.Request(request, new(AcousticModel))
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
// Lists information about all custom acoustic models that are owned by an instance of the service. Use the `language`
// parameter to see all custom acoustic models for the specified language. Omit the parameter to see all custom acoustic
// models for all languages. You must use credentials for the instance of the service that owns a model to list
// information about it.
//
// **See also:** [Listing custom acoustic
// models](https://cloud.ibm.com/docs/services/speech-to-text/acoustic-models.html#listModels-acoustic).
func (speechToText *SpeechToTextV1) ListAcousticModels(listAcousticModelsOptions *ListAcousticModelsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateStruct(listAcousticModelsOptions, "listAcousticModelsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/acoustic_customizations"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

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
		return nil, err
	}

	response, err := speechToText.Service.Request(request, new(AcousticModels))
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
// Resets a custom acoustic model by removing all audio resources from the model. Resetting a custom acoustic model
// initializes the model to its state when it was first created. Metadata such as the name and language of the model are
// preserved, but the model's audio resources are removed and must be re-created. You must use credentials for the
// instance of the service that owns a model to reset it.
//
// **See also:** [Resetting a custom acoustic
// model](https://cloud.ibm.com/docs/services/speech-to-text/acoustic-models.html#resetModel-acoustic).
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
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

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
		return nil, err
	}

	response, err := speechToText.Service.Request(request, nil)
	return response, err
}

// TrainAcousticModel : Train a custom acoustic model
// Initiates the training of a custom acoustic model with new or changed audio resources. After adding or deleting audio
// resources for a custom acoustic model, use this method to begin the actual training of the model on the latest audio
// data. The custom acoustic model does not reflect its changed data until you train it. You must use credentials for
// the instance of the service that owns a model to train it.
//
// The training method is asynchronous. It can take on the order of minutes or hours to complete depending on the total
// amount of audio data on which the custom acoustic model is being trained and the current load on the service.
// Typically, training a custom acoustic model takes approximately two to four times the length of its audio data. The
// range of time depends on the model being trained and the nature of the audio, such as whether the audio is clean or
// noisy. The method returns an HTTP 200 response code to indicate that the training process has begun.
//
// You can monitor the status of the training by using the **Get a custom acoustic model** method to poll the model's
// status. Use a loop to check the status once a minute. The method returns an `AcousticModel` object that includes
// `status` and `progress` fields. A status of `available` indicates that the custom model is trained and ready to use.
// The service cannot accept subsequent training requests, or requests to add new audio resources, until the existing
// request completes.
//
// You can use the optional `custom_language_model_id` parameter to specify the GUID of a separately created custom
// language model that is to be used during training. Train with a custom language model if you have verbatim
// transcriptions of the audio files that you have added to the custom model or you have either corpora (text files) or
// a list of words that are relevant to the contents of the audio files. Both of the custom models must be based on the
// same version of the same base model for training to succeed.
//
// Training can fail to start for the following reasons:
// * The service is currently handling another request for the custom model, such as another training request or a
// request to add audio resources to the model.
// * The custom model contains less than 10 minutes or more than 100 hours of audio data.
// * One or more of the custom model's audio resources is invalid.
// * You passed an incompatible custom language model with the `custom_language_model_id` query parameter. Both custom
// models must be based on the same version of the same base model.
//
// **See also:** [Train the custom acoustic
// model](https://cloud.ibm.com/docs/services/speech-to-text/acoustic-create.html#trainModel-acoustic).
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
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

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
		return nil, err
	}

	response, err := speechToText.Service.Request(request, nil)
	return response, err
}

// UpgradeAcousticModel : Upgrade a custom acoustic model
// Initiates the upgrade of a custom acoustic model to the latest version of its base language model. The upgrade method
// is asynchronous. It can take on the order of minutes or hours to complete depending on the amount of data in the
// custom model and the current load on the service; typically, upgrade takes approximately twice the length of the
// total audio contained in the custom model. A custom model must be in the `ready` or `available` state to be upgraded.
// You must use credentials for the instance of the service that owns a model to upgrade it.
//
// The method returns an HTTP 200 response code to indicate that the upgrade process has begun successfully. You can
// monitor the status of the upgrade by using the **Get a custom acoustic model** method to poll the model's status. The
// method returns an `AcousticModel` object that includes `status` and `progress` fields. Use a loop to check the status
// once a minute. While it is being upgraded, the custom model has the status `upgrading`. When the upgrade is complete,
// the model resumes the status that it had prior to upgrade. The service cannot accept subsequent requests for the
// model until the upgrade completes.
//
// If the custom acoustic model was trained with a separately created custom language model, you must use the
// `custom_language_model_id` parameter to specify the GUID of that custom language model. The custom language model
// must be upgraded before the custom acoustic model can be upgraded. Omit the parameter if the custom acoustic model
// was not trained with a custom language model.
//
// **See also:** [Upgrading a custom acoustic
// model](https://cloud.ibm.com/docs/services/speech-to-text/custom-upgrade.html#upgradeAcoustic).
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
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

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
		return nil, err
	}

	response, err := speechToText.Service.Request(request, nil)
	return response, err
}

// AddAudio : Add an audio resource
// Adds an audio resource to a custom acoustic model. Add audio content that reflects the acoustic characteristics of
// the audio that you plan to transcribe. You must use credentials for the instance of the service that owns a model to
// add an audio resource to it. Adding audio data does not affect the custom acoustic model until you train the model
// for the new data by using the **Train a custom acoustic model** method.
//
// You can add individual audio files or an archive file that contains multiple audio files. Adding multiple audio files
// via a single archive file is significantly more efficient than adding each file individually. You can add audio
// resources in any format that the service supports for speech recognition.
//
// You can use this method to add any number of audio resources to a custom model by calling the method once for each
// audio or archive file. But the addition of one audio resource must be fully complete before you can add another. You
// must add a minimum of 10 minutes and a maximum of 100 hours of audio that includes speech, not just silence, to a
// custom acoustic model before you can train it. No audio resource, audio- or archive-type, can be larger than 100 MB.
// To add an audio resource that has the same name as an existing audio resource, set the `allow_overwrite` parameter to
// `true`; otherwise, the request fails.
//
// The method is asynchronous. It can take several seconds to complete depending on the duration of the audio and, in
// the case of an archive file, the total number of audio files being processed. The service returns a 201 response code
// if the audio is valid. It then asynchronously analyzes the contents of the audio file or files and automatically
// extracts information about the audio such as its length, sampling rate, and encoding. You cannot submit requests to
// add additional audio resources to a custom acoustic model, or to train the model, until the service's analysis of all
// audio files for the current request completes.
//
// To determine the status of the service's analysis of the audio, use the **Get an audio resource** method to poll the
// status of the audio. The method accepts the customization ID of the custom model and the name of the audio resource,
// and it returns the status of the resource. Use a loop to check the status of the audio every few seconds until it
// becomes `ok`.
//
// **See also:** [Add audio to the custom acoustic
// model](https://cloud.ibm.com/docs/services/speech-to-text/acoustic-create.html#addAudio).
//
// ### Content types for audio-type resources
//
//  You can add an individual audio file in any format that the service supports for speech recognition. For an
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
//  **See also:** [Audio formats](https://cloud.ibm.com/docs/services/speech-to-text/audio-formats.html).
//
// ### Content types for archive-type resources
//
//  You can add an archive file (**.zip** or **.tar.gz** file) that contains audio files in any format that the service
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
//  The name of an audio file that is embedded within an archive-type resource must meet the following restrictions:
// * Include a maximum of 128 characters in the file name; this includes the file extension.
// * Do not include spaces, slashes, or backslashes in the file name.
// * Do not use the name of an audio file that has already been added to the custom model as part of an archive-type
// resource.
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
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range addAudioOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "AddAudio")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	if addAudioOptions.ContainedContentType != nil {
		builder.AddHeader("Contained-Content-Type", fmt.Sprint(*addAudioOptions.ContainedContentType))
	}
	if addAudioOptions.ContentType != nil {
		builder.AddHeader("Content-Type", fmt.Sprint(*addAudioOptions.ContentType))
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

	response, err := speechToText.Service.Request(request, nil)
	return response, err
}

// DeleteAudio : Delete an audio resource
// Deletes an existing audio resource from a custom acoustic model. Deleting an archive-type audio resource removes the
// entire archive of files; the current interface does not allow deletion of individual files from an archive resource.
// Removing an audio resource does not affect the custom model until you train the model on its updated data by using
// the **Train a custom acoustic model** method. You must use credentials for the instance of the service that owns a
// model to delete its audio resources.
//
// **See also:** [Deleting an audio resource from a custom acoustic
// model](https://cloud.ibm.com/docs/services/speech-to-text/acoustic-audio.html#deleteAudio).
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
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

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
		return nil, err
	}

	response, err := speechToText.Service.Request(request, nil)
	return response, err
}

// GetAudio : Get an audio resource
// Gets information about an audio resource from a custom acoustic model. The method returns an `AudioListing` object
// whose fields depend on the type of audio resource that you specify with the method's `audio_name` parameter:
// * **For an audio-type resource,** the object's fields match those of an `AudioResource` object: `duration`, `name`,
// `details`, and `status`.
// * **For an archive-type resource,** the object includes a `container` field whose fields match those of an
// `AudioResource` object. It also includes an `audio` field, which contains an array of `AudioResource` objects that
// provides information about the audio files that are contained in the archive.
//
// The information includes the status of the specified audio resource. The status is important for checking the
// service's analysis of a resource that you add to the custom model.
// * For an audio-type resource, the `status` field is located in the `AudioListing` object.
// * For an archive-type resource, the `status` field is located in the `AudioResource` object that is returned in the
// `container` field.
//
// You must use credentials for the instance of the service that owns a model to list its audio resources.
//
// **See also:** [Listing audio resources for a custom acoustic
// model](https://cloud.ibm.com/docs/services/speech-to-text/acoustic-audio.html#listAudio).
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
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

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
		return nil, err
	}

	response, err := speechToText.Service.Request(request, new(AudioListing))
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
// Lists information about all audio resources from a custom acoustic model. The information includes the name of the
// resource and information about its audio data, such as its duration. It also includes the status of the audio
// resource, which is important for checking the service's analysis of the resource in response to a request to add it
// to the custom acoustic model. You must use credentials for the instance of the service that owns a model to list its
// audio resources.
//
// **See also:** [Listing audio resources for a custom acoustic
// model](https://cloud.ibm.com/docs/services/speech-to-text/acoustic-audio.html#listAudio).
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
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

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
		return nil, err
	}

	response, err := speechToText.Service.Request(request, new(AudioResources))
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
// Deletes all data that is associated with a specified customer ID. The method deletes all data for the customer ID,
// regardless of the method by which the information was added. The method has no effect if no data is associated with
// the customer ID. You must issue the request with credentials for the same instance of the service that was used to
// associate the customer ID with the data.
//
// You associate a customer ID with data by passing the `X-Watson-Metadata` header with a request that passes the data.
//
// **See also:** [Information security](https://cloud.ibm.com/docs/services/speech-to-text/information-security.html).
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
	builder.ConstructHTTPURL(speechToText.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteUserDataOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("speech_to_text", "V1", "DeleteUserData")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "")

	builder.AddQuery("customer_id", fmt.Sprint(*deleteUserDataOptions.CustomerID))

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := speechToText.Service.Request(request, nil)
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

	// The GUID of the credentials for the instance of the service that owns the custom acoustic model.
	Owner *string `json:"owner,omitempty"`

	// The name of the custom acoustic model.
	Name *string `json:"name,omitempty"`

	// The description of the custom acoustic model.
	Description *string `json:"description,omitempty"`

	// The name of the language model for which the custom acoustic model was created.
	BaseModelName *string `json:"base_model_name,omitempty"`

	// The current status of the custom acoustic model:
	// * `pending`: The model was created but is waiting either for training data to be added or for the service to finish
	// analyzing added data.
	// * `ready`: The model contains data and is ready to be trained.
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
// * `pending`: The model was created but is waiting either for training data to be added or for the service to finish
// analyzing added data.
// * `ready`: The model contains data and is ready to be trained.
// * `training`: The model is currently being trained.
// * `available`: The model is trained and ready to use.
// * `upgrading`: The model is currently being upgraded.
// * `failed`: Training of the model failed.
const (
	AcousticModel_Status_Available = "available"
	AcousticModel_Status_Failed    = "failed"
	AcousticModel_Status_Pending   = "pending"
	AcousticModel_Status_Ready     = "ready"
	AcousticModel_Status_Training  = "training"
	AcousticModel_Status_Upgrading = "upgrading"
)

// AcousticModels : AcousticModels struct
type AcousticModels struct {

	// An array of `AcousticModel` objects that provides information about each available custom acoustic model. The array
	// is empty if the requesting credentials own no custom acoustic models (if no language is specified) or own no custom
	// acoustic models for the specified language.
	Customizations []AcousticModel `json:"customizations" validate:"required"`
}

// AddAudioOptions : The addAudio options.
type AddAudioOptions struct {

	// The customization ID (GUID) of the custom acoustic model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// The name of the new audio resource for the custom acoustic model. Use a localized name that matches the language of
	// the custom model and reflects the contents of the resource.
	// * Include a maximum of 128 characters in the name.
	// * Do not include spaces, slashes, or backslashes in the name.
	// * Do not use the name of an audio resource that has already been added to the custom model.
	AudioName *string `json:"audio_name" validate:"required"`

	// The audio resource that is to be added to the custom acoustic model, an individual audio file or an archive file.
	AudioResource *io.ReadCloser `json:"audio_resource" validate:"required"`

	// **For an archive-type resource,** specify the format of the audio files that are contained in the archive file if
	// they are of type `audio/alaw`, `audio/basic`, `audio/l16`, or `audio/mulaw`. Include the `rate`, `channels`, and
	// `endianness` parameters where necessary. In this case, all audio files that are contained in the archive file must
	// be of the indicated type.
	//
	// For all other audio formats, you can omit the header. In this case, the audio files can be of multiple types as long
	// as they are not of the types listed in the previous paragraph.
	//
	// The parameter accepts all of the audio formats that are supported for use with speech recognition. For more
	// information, see **Content types for audio-type resources** in the method description.
	//
	// **For an audio-type resource,** omit the header.
	ContainedContentType *string `json:"Contained-Content-Type,omitempty"`

	// If `true`, the specified audio resource overwrites an existing audio resource with the same name. If `false`, the
	// request fails if an audio resource with the same name already exists. The parameter has no effect if an audio
	// resource with the same name does not already exist.
	AllowOverwrite *bool `json:"allow_overwrite,omitempty"`

	// For an audio-type resource, the format (MIME type) of the audio. For more information, see **Content types for
	// audio-type resources** in the method description.
	//
	// For an archive-type resource, the media type of the archive file. For more information, see **Content types for
	// archive-type resources** in the method description.
	ContentType *string `json:"Content-Type,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the AddAudioOptions.ContainedContentType property.
// **For an archive-type resource,** specify the format of the audio files that are contained in the archive file if
// they are of type `audio/alaw`, `audio/basic`, `audio/l16`, or `audio/mulaw`. Include the `rate`, `channels`, and
// `endianness` parameters where necessary. In this case, all audio files that are contained in the archive file must be
// of the indicated type.
//
// For all other audio formats, you can omit the header. In this case, the audio files can be of multiple types as long
// as they are not of the types listed in the previous paragraph.
//
// The parameter accepts all of the audio formats that are supported for use with speech recognition. For more
// information, see **Content types for audio-type resources** in the method description.
//
// **For an audio-type resource,** omit the header.
const (
	AddAudioOptions_ContainedContentType_AudioAlaw             = "audio/alaw"
	AddAudioOptions_ContainedContentType_AudioBasic            = "audio/basic"
	AddAudioOptions_ContainedContentType_AudioFlac             = "audio/flac"
	AddAudioOptions_ContainedContentType_AudioG729             = "audio/g729"
	AddAudioOptions_ContainedContentType_AudioL16              = "audio/l16"
	AddAudioOptions_ContainedContentType_AudioMp3              = "audio/mp3"
	AddAudioOptions_ContainedContentType_AudioMpeg             = "audio/mpeg"
	AddAudioOptions_ContainedContentType_AudioMulaw            = "audio/mulaw"
	AddAudioOptions_ContainedContentType_AudioOgg              = "audio/ogg"
	AddAudioOptions_ContainedContentType_AudioOggCodecsOpus    = "audio/ogg;codecs=opus"
	AddAudioOptions_ContainedContentType_AudioOggCodecsVorbis  = "audio/ogg;codecs=vorbis"
	AddAudioOptions_ContainedContentType_AudioWav              = "audio/wav"
	AddAudioOptions_ContainedContentType_AudioWebm             = "audio/webm"
	AddAudioOptions_ContainedContentType_AudioWebmCodecsOpus   = "audio/webm;codecs=opus"
	AddAudioOptions_ContainedContentType_AudioWebmCodecsVorbis = "audio/webm;codecs=vorbis"
)

// Constants associated with the AddAudioOptions.ContentType property.
// For an audio-type resource, the format (MIME type) of the audio. For more information, see **Content types for
// audio-type resources** in the method description.
//
// For an archive-type resource, the media type of the archive file. For more information, see **Content types for
// archive-type resources** in the method description.
const (
	AddAudioOptions_ContentType_ApplicationGzip       = "application/gzip"
	AddAudioOptions_ContentType_ApplicationZip        = "application/zip"
	AddAudioOptions_ContentType_AudioAlaw             = "audio/alaw"
	AddAudioOptions_ContentType_AudioBasic            = "audio/basic"
	AddAudioOptions_ContentType_AudioFlac             = "audio/flac"
	AddAudioOptions_ContentType_AudioG729             = "audio/g729"
	AddAudioOptions_ContentType_AudioL16              = "audio/l16"
	AddAudioOptions_ContentType_AudioMp3              = "audio/mp3"
	AddAudioOptions_ContentType_AudioMpeg             = "audio/mpeg"
	AddAudioOptions_ContentType_AudioMulaw            = "audio/mulaw"
	AddAudioOptions_ContentType_AudioOgg              = "audio/ogg"
	AddAudioOptions_ContentType_AudioOggCodecsOpus    = "audio/ogg;codecs=opus"
	AddAudioOptions_ContentType_AudioOggCodecsVorbis  = "audio/ogg;codecs=vorbis"
	AddAudioOptions_ContentType_AudioWav              = "audio/wav"
	AddAudioOptions_ContentType_AudioWebm             = "audio/webm"
	AddAudioOptions_ContentType_AudioWebmCodecsOpus   = "audio/webm;codecs=opus"
	AddAudioOptions_ContentType_AudioWebmCodecsVorbis = "audio/webm;codecs=vorbis"
)

// NewAddAudioOptions : Instantiate AddAudioOptions
func (speechToText *SpeechToTextV1) NewAddAudioOptions(customizationID string, audioName string, audioResource io.ReadCloser) *AddAudioOptions {
	return &AddAudioOptions{
		CustomizationID: core.StringPtr(customizationID),
		AudioName:       core.StringPtr(audioName),
		AudioResource:   &audioResource,
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

// SetAudioResource : Allow user to set AudioResource
func (options *AddAudioOptions) SetAudioResource(audioResource io.ReadCloser) *AddAudioOptions {
	options.AudioResource = &audioResource
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

// SetContentType : Allow user to set ContentType
func (options *AddAudioOptions) SetContentType(contentType string) *AddAudioOptions {
	options.ContentType = core.StringPtr(contentType)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *AddAudioOptions) SetHeaders(param map[string]string) *AddAudioOptions {
	options.Headers = param
	return options
}

// AddCorpusOptions : The addCorpus options.
type AddCorpusOptions struct {

	// The customization ID (GUID) of the custom language model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// The name of the new corpus for the custom language model. Use a localized name that matches the language of the
	// custom model and reflects the contents of the corpus.
	// * Include a maximum of 128 characters in the name.
	// * Do not include spaces, slashes, or backslashes in the name.
	// * Do not use the name of an existing corpus or grammar that is already defined for the custom model.
	// * Do not use the name `user`, which is reserved by the service to denote custom words that are added or modified by
	// the user.
	CorpusName *string `json:"corpus_name" validate:"required"`

	// A plain text file that contains the training data for the corpus. Encode the file in UTF-8 if it contains non-ASCII
	// characters; the service assumes UTF-8 encoding if it encounters non-ASCII characters.
	//
	// Make sure that you know the character encoding of the file. You must use that encoding when working with the words
	// in the custom language model. For more information, see [Character
	// encoding](https://cloud.ibm.com/docs/services/speech-to-text/language-resource.html#charEncoding).
	//
	// With the `curl` command, use the `--data-binary` option to upload the file for the request.
	CorpusFile *os.File `json:"corpus_file" validate:"required"`

	// If `true`, the specified corpus overwrites an existing corpus with the same name. If `false`, the request fails if a
	// corpus with the same name already exists. The parameter has no effect if a corpus with the same name does not
	// already exist.
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

// AddGrammarOptions : The addGrammar options.
type AddGrammarOptions struct {

	// The customization ID (GUID) of the custom language model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// The name of the new grammar for the custom language model. Use a localized name that matches the language of the
	// custom model and reflects the contents of the grammar.
	// * Include a maximum of 128 characters in the name.
	// * Do not include spaces, slashes, or backslashes in the name.
	// * Do not use the name of an existing grammar or corpus that is already defined for the custom model.
	// * Do not use the name `user`, which is reserved by the service to denote custom words that are added or modified by
	// the user.
	GrammarName *string `json:"grammar_name" validate:"required"`

	// A plain text file that contains the grammar in the format specified by the `Content-Type` header. Encode the file in
	// UTF-8 (ASCII is a subset of UTF-8). Using any other encoding can lead to issues when compiling the grammar or to
	// unexpected results in decoding. The service ignores an encoding that is specified in the header of the grammar.
	GrammarFile *io.ReadCloser `json:"grammar_file" validate:"required"`

	// The format (MIME type) of the grammar file:
	// * `application/srgs` for Augmented Backus-Naur Form (ABNF), which uses a plain-text representation that is similar
	// to traditional BNF grammars.
	// * `application/srgs+xml` for XML Form, which uses XML elements to represent the grammar.
	ContentType *string `json:"Content-Type" validate:"required"`

	// If `true`, the specified grammar overwrites an existing grammar with the same name. If `false`, the request fails if
	// a grammar with the same name already exists. The parameter has no effect if a grammar with the same name does not
	// already exist.
	AllowOverwrite *bool `json:"allow_overwrite,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the AddGrammarOptions.ContentType property.
// The format (MIME type) of the grammar file:
// * `application/srgs` for Augmented Backus-Naur Form (ABNF), which uses a plain-text representation that is similar to
// traditional BNF grammars.
// * `application/srgs+xml` for XML Form, which uses XML elements to represent the grammar.
const (
	AddGrammarOptions_ContentType_ApplicationSrgs    = "application/srgs"
	AddGrammarOptions_ContentType_ApplicationSrgsXml = "application/srgs+xml"
)

// NewAddGrammarOptions : Instantiate AddGrammarOptions
func (speechToText *SpeechToTextV1) NewAddGrammarOptions(customizationID string, grammarName string, grammarFile io.ReadCloser, contentType string) *AddGrammarOptions {
	return &AddGrammarOptions{
		CustomizationID: core.StringPtr(customizationID),
		GrammarName:     core.StringPtr(grammarName),
		GrammarFile:     &grammarFile,
		ContentType:     core.StringPtr(contentType),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *AddGrammarOptions) SetCustomizationID(customizationID string) *AddGrammarOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetGrammarName : Allow user to set GrammarName
func (options *AddGrammarOptions) SetGrammarName(grammarName string) *AddGrammarOptions {
	options.GrammarName = core.StringPtr(grammarName)
	return options
}

// SetGrammarFile : Allow user to set GrammarFile
func (options *AddGrammarOptions) SetGrammarFile(grammarFile io.ReadCloser) *AddGrammarOptions {
	options.GrammarFile = &grammarFile
	return options
}

// SetContentType : Allow user to set ContentType
func (options *AddGrammarOptions) SetContentType(contentType string) *AddGrammarOptions {
	options.ContentType = core.StringPtr(contentType)
	return options
}

// SetAllowOverwrite : Allow user to set AllowOverwrite
func (options *AddGrammarOptions) SetAllowOverwrite(allowOverwrite bool) *AddGrammarOptions {
	options.AllowOverwrite = core.BoolPtr(allowOverwrite)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *AddGrammarOptions) SetHeaders(param map[string]string) *AddGrammarOptions {
	options.Headers = param
	return options
}

// AddWordOptions : The addWord options.
type AddWordOptions struct {

	// The customization ID (GUID) of the custom language model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// The custom word that is to be added to or updated in the custom language model. Do not include spaces in the word.
	// Use a `-` (dash) or `_` (underscore) to connect the tokens of compound words. URL-encode the word if it includes
	// non-ASCII characters. For more information, see [Character
	// encoding](https://cloud.ibm.com/docs/services/speech-to-text/language-resource.html#charEncoding).
	WordName *string `json:"word_name" validate:"required"`

	// For the **Add custom words** method, you must specify the custom word that is to be added to or updated in the
	// custom model. Do not include spaces in the word. Use a `-` (dash) or `_` (underscore) to connect the tokens of
	// compound words.
	//
	// Omit this parameter for the **Add a custom word** method.
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

	// The customization ID (GUID) of the custom language model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// An array of `CustomWord` objects that provides information about each custom word that is to be added to or updated
	// in the custom language model.
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

// Constants associated with the AudioDetails.Type property.
// The type of the audio resource:
// * `audio` for an individual audio file
// * `archive` for an archive (**.zip** or **.tar.gz**) file that contains audio files
// * `undetermined` for a resource that the service cannot validate (for example, if the user mistakenly passes a file
// that does not contain audio, such as a JPEG file).
const (
	AudioDetails_Type_Archive      = "archive"
	AudioDetails_Type_Audio        = "audio"
	AudioDetails_Type_Undetermined = "undetermined"
)

// Constants associated with the AudioDetails.Compression property.
// **For an archive-type resource,** the format of the compressed archive:
// * `zip` for a **.zip** file
// * `gzip` for a **.tar.gz** file
//
// Omitted for an audio-type resource.
const (
	AudioDetails_Compression_Gzip = "gzip"
	AudioDetails_Compression_Zip  = "zip"
)

// AudioListing : AudioListing struct
type AudioListing struct {

	// **For an audio-type resource,**  the total seconds of audio in the resource. Omitted for an archive-type resource.
	Duration *int64 `json:"duration,omitempty"`

	// **For an audio-type resource,** the user-specified name of the resource. Omitted for an archive-type resource.
	Name *string `json:"name,omitempty"`

	// **For an audio-type resource,** an `AudioDetails` object that provides detailed information about the resource. The
	// object is empty until the service finishes processing the audio. Omitted for an archive-type resource.
	Details *AudioDetails `json:"details,omitempty"`

	// **For an audio-type resource,** the status of the resource:
	// * `ok`: The service successfully analyzed the audio data. The data can be used to train the custom model.
	// * `being_processed`: The service is still analyzing the audio data. The service cannot accept requests to add new
	// audio resources or to train the custom model until its analysis is complete.
	// * `invalid`: The audio data is not valid for training the custom model (possibly because it has the wrong format or
	// sampling rate, or because it is corrupted).
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

// Constants associated with the AudioListing.Status property.
// **For an audio-type resource,** the status of the resource:
// * `ok`: The service successfully analyzed the audio data. The data can be used to train the custom model.
// * `being_processed`: The service is still analyzing the audio data. The service cannot accept requests to add new
// audio resources or to train the custom model until its analysis is complete.
// * `invalid`: The audio data is not valid for training the custom model (possibly because it has the wrong format or
// sampling rate, or because it is corrupted).
//
// Omitted for an archive-type resource.
const (
	AudioListing_Status_BeingProcessed = "being_processed"
	AudioListing_Status_Invalid        = "invalid"
	AudioListing_Status_Ok             = "ok"
)

// AudioResource : AudioResource struct
type AudioResource struct {

	// The total seconds of audio in the audio resource.
	Duration *int64 `json:"duration" validate:"required"`

	// **For an archive-type resource,** the user-specified name of the resource.
	//
	// **For an audio-type resource,** the user-specified name of the resource or the name of the audio file that the user
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
	AudioResource_Status_BeingProcessed = "being_processed"
	AudioResource_Status_Invalid        = "invalid"
	AudioResource_Status_Ok             = "ok"
)

// AudioResources : AudioResources struct
type AudioResources struct {

	// The total minutes of accumulated audio summed over all of the valid audio resources for the custom acoustic model.
	// You can use this value to determine whether the custom model has too little or too much audio to begin training.
	TotalMinutesOfAudio *float64 `json:"total_minutes_of_audio" validate:"required"`

	// An array of `AudioResource` objects that provides information about the audio resources of the custom acoustic
	// model. The array is empty if the custom model has no audio resources.
	Audio []AudioResource `json:"audio" validate:"required"`
}

// CheckJobOptions : The checkJob options.
type CheckJobOptions struct {

	// The identifier of the asynchronous job that is to be used for the request. You must make the request with
	// credentials for the instance of the service that owns the job.
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

	// An array of `Corpus` objects that provides information about the corpora for the custom model. The array is empty if
	// the custom model has no corpora.
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
	Corpus_Status_Analyzed       = "analyzed"
	Corpus_Status_BeingProcessed = "being_processed"
	Corpus_Status_Undetermined   = "undetermined"
)

// CreateAcousticModelOptions : The createAcousticModel options.
type CreateAcousticModelOptions struct {

	// A user-defined name for the new custom acoustic model. Use a name that is unique among all custom acoustic models
	// that you own. Use a localized name that matches the language of the custom model. Use a name that describes the
	// acoustic environment of the custom model, such as `Mobile custom model` or `Noisy car custom model`.
	Name *string `json:"name" validate:"required"`

	// The name of the base language model that is to be customized by the new custom acoustic model. The new custom model
	// can be used only with the base model that it customizes.
	//
	// To determine whether a base model supports acoustic model customization, refer to [Language support for
	// customization](https://cloud.ibm.com/docs/services/speech-to-text/custom.html#languageSupport).
	BaseModelName *string `json:"base_model_name" validate:"required"`

	// A description of the new custom acoustic model. Use a localized description that matches the language of the custom
	// model.
	Description *string `json:"description,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the CreateAcousticModelOptions.BaseModelName property.
// The name of the base language model that is to be customized by the new custom acoustic model. The new custom model
// can be used only with the base model that it customizes.
//
// To determine whether a base model supports acoustic model customization, refer to [Language support for
// customization](https://cloud.ibm.com/docs/services/speech-to-text/custom.html#languageSupport).
const (
	CreateAcousticModelOptions_BaseModelName_ArArBroadbandmodel           = "ar-AR_BroadbandModel"
	CreateAcousticModelOptions_BaseModelName_DeDeBroadbandmodel           = "de-DE_BroadbandModel"
	CreateAcousticModelOptions_BaseModelName_DeDeNarrowbandmodel          = "de-DE_NarrowbandModel"
	CreateAcousticModelOptions_BaseModelName_EnGbBroadbandmodel           = "en-GB_BroadbandModel"
	CreateAcousticModelOptions_BaseModelName_EnGbNarrowbandmodel          = "en-GB_NarrowbandModel"
	CreateAcousticModelOptions_BaseModelName_EnUsBroadbandmodel           = "en-US_BroadbandModel"
	CreateAcousticModelOptions_BaseModelName_EnUsNarrowbandmodel          = "en-US_NarrowbandModel"
	CreateAcousticModelOptions_BaseModelName_EnUsShortformNarrowbandmodel = "en-US_ShortForm_NarrowbandModel"
	CreateAcousticModelOptions_BaseModelName_EsEsBroadbandmodel           = "es-ES_BroadbandModel"
	CreateAcousticModelOptions_BaseModelName_EsEsNarrowbandmodel          = "es-ES_NarrowbandModel"
	CreateAcousticModelOptions_BaseModelName_FrFrBroadbandmodel           = "fr-FR_BroadbandModel"
	CreateAcousticModelOptions_BaseModelName_FrFrNarrowbandmodel          = "fr-FR_NarrowbandModel"
	CreateAcousticModelOptions_BaseModelName_JaJpBroadbandmodel           = "ja-JP_BroadbandModel"
	CreateAcousticModelOptions_BaseModelName_JaJpNarrowbandmodel          = "ja-JP_NarrowbandModel"
	CreateAcousticModelOptions_BaseModelName_KoKrBroadbandmodel           = "ko-KR_BroadbandModel"
	CreateAcousticModelOptions_BaseModelName_KoKrNarrowbandmodel          = "ko-KR_NarrowbandModel"
	CreateAcousticModelOptions_BaseModelName_PtBrBroadbandmodel           = "pt-BR_BroadbandModel"
	CreateAcousticModelOptions_BaseModelName_PtBrNarrowbandmodel          = "pt-BR_NarrowbandModel"
	CreateAcousticModelOptions_BaseModelName_ZhCnBroadbandmodel           = "zh-CN_BroadbandModel"
	CreateAcousticModelOptions_BaseModelName_ZhCnNarrowbandmodel          = "zh-CN_NarrowbandModel"
)

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

	// The audio to transcribe.
	Audio *io.ReadCloser `json:"audio" validate:"required"`

	// The identifier of the model that is to be used for the recognition request. See [Languages and
	// models](https://cloud.ibm.com/docs/services/speech-to-text/models.html).
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
	// make the request with credentials for the instance of the service that owns the custom model. By default, no custom
	// language model is used. See [Custom
	// models](https://cloud.ibm.com/docs/services/speech-to-text/input.html#custom-input).
	//
	// **Note:** Use this parameter instead of the deprecated `customization_id` parameter.
	LanguageCustomizationID *string `json:"language_customization_id,omitempty"`

	// The customization ID (GUID) of a custom acoustic model that is to be used with the recognition request. The base
	// model of the specified custom acoustic model must match the model specified with the `model` parameter. You must
	// make the request with credentials for the instance of the service that owns the custom model. By default, no custom
	// acoustic model is used. See [Custom
	// models](https://cloud.ibm.com/docs/services/speech-to-text/input.html#custom-input).
	AcousticCustomizationID *string `json:"acoustic_customization_id,omitempty"`

	// The version of the specified base model that is to be used with recognition request. Multiple versions of a base
	// model can exist when a model is updated for internal improvements. The parameter is intended primarily for use with
	// custom models that have been upgraded for a new base model. The default value depends on whether the parameter is
	// used with or without a custom model. See [Base model
	// version](https://cloud.ibm.com/docs/services/speech-to-text/input.html#version).
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
	// See [Custom models](https://cloud.ibm.com/docs/services/speech-to-text/input.html#custom-input).
	CustomizationWeight *float64 `json:"customization_weight,omitempty"`

	// The time in seconds after which, if only silence (no speech) is detected in streaming audio, the connection is
	// closed with a 400 error. The parameter is useful for stopping audio submission from a live microphone when a user
	// simply walks away. Use `-1` for infinity. See [Inactivity
	// timeout](https://cloud.ibm.com/docs/services/speech-to-text/input.html#timeouts-inactivity).
	InactivityTimeout *int64 `json:"inactivity_timeout,omitempty"`

	// An array of keyword strings to spot in the audio. Each keyword string can include one or more string tokens.
	// Keywords are spotted only in the final results, not in interim hypotheses. If you specify any keywords, you must
	// also specify a keywords threshold. You can spot a maximum of 1000 keywords. Omit the parameter or specify an empty
	// array if you do not need to spot keywords. See [Keyword
	// spotting](https://cloud.ibm.com/docs/services/speech-to-text/output.html#keyword_spotting).
	Keywords []string `json:"keywords,omitempty"`

	// A confidence value that is the lower bound for spotting a keyword. A word is considered to match a keyword if its
	// confidence is greater than or equal to the threshold. Specify a probability between 0.0 and 1.0. If you specify a
	// threshold, you must also specify one or more keywords. The service performs no keyword spotting if you omit either
	// parameter. See [Keyword spotting](https://cloud.ibm.com/docs/services/speech-to-text/output.html#keyword_spotting).
	KeywordsThreshold *float32 `json:"keywords_threshold,omitempty"`

	// The maximum number of alternative transcripts that the service is to return. By default, the service returns a
	// single transcript. If you specify a value of `0`, the service uses the default value, `1`. See [Maximum
	// alternatives](https://cloud.ibm.com/docs/services/speech-to-text/output.html#max_alternatives).
	MaxAlternatives *int64 `json:"max_alternatives,omitempty"`

	// A confidence value that is the lower bound for identifying a hypothesis as a possible word alternative (also known
	// as "Confusion Networks"). An alternative word is considered if its confidence is greater than or equal to the
	// threshold. Specify a probability between 0.0 and 1.0. By default, the service computes no alternative words. See
	// [Word alternatives](https://cloud.ibm.com/docs/services/speech-to-text/output.html#word_alternatives).
	WordAlternativesThreshold *float32 `json:"word_alternatives_threshold,omitempty"`

	// If `true`, the service returns a confidence measure in the range of 0.0 to 1.0 for each word. By default, the
	// service returns no word confidence scores. See [Word
	// confidence](https://cloud.ibm.com/docs/services/speech-to-text/output.html#word_confidence).
	WordConfidence *bool `json:"word_confidence,omitempty"`

	// If `true`, the service returns time alignment for each word. By default, no timestamps are returned. See [Word
	// timestamps](https://cloud.ibm.com/docs/services/speech-to-text/output.html#word_timestamps).
	Timestamps *bool `json:"timestamps,omitempty"`

	// If `true`, the service filters profanity from all output except for keyword results by replacing inappropriate words
	// with a series of asterisks. Set the parameter to `false` to return results with no censoring. Applies to US English
	// transcription only. See [Profanity
	// filtering](https://cloud.ibm.com/docs/services/speech-to-text/output.html#profanity_filter).
	ProfanityFilter *bool `json:"profanity_filter,omitempty"`

	// If `true`, the service converts dates, times, series of digits and numbers, phone numbers, currency values, and
	// internet addresses into more readable, conventional representations in the final transcript of a recognition
	// request. For US English, the service also converts certain keyword strings to punctuation symbols. By default, the
	// service performs no smart formatting.
	//
	// **Note:** Applies to US English, Japanese, and Spanish transcription only.
	//
	// See [Smart formatting](https://cloud.ibm.com/docs/services/speech-to-text/output.html#smart_formatting).
	SmartFormatting *bool `json:"smart_formatting,omitempty"`

	// If `true`, the response includes labels that identify which words were spoken by which participants in a
	// multi-person exchange. By default, the service returns no speaker labels. Setting `speaker_labels` to `true` forces
	// the `timestamps` parameter to be `true`, regardless of whether you specify `false` for the parameter.
	//
	// **Note:** Applies to US English, Japanese, and Spanish transcription only. To determine whether a language model
	// supports speaker labels, you can also use the **Get a model** method and check that the attribute `speaker_labels`
	// is set to `true`.
	//
	// See [Speaker labels](https://cloud.ibm.com/docs/services/speech-to-text/output.html#speaker_labels).
	SpeakerLabels *bool `json:"speaker_labels,omitempty"`

	// **Deprecated.** Use the `language_customization_id` parameter to specify the customization ID (GUID) of a custom
	// language model that is to be used with the recognition request. Do not specify both parameters with a request.
	CustomizationID *string `json:"customization_id,omitempty"`

	// The name of a grammar that is to be used with the recognition request. If you specify a grammar, you must also use
	// the `language_customization_id` parameter to specify the name of the custom language model for which the grammar is
	// defined. The service recognizes only strings that are recognized by the specified grammar; it does not recognize
	// other custom words from the model's words resource. See
	// [Grammars](https://cloud.ibm.com/docs/services/speech-to-text/input.html#grammars-input).
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
	// **Note:** Applies to US English, Japanese, and Korean transcription only.
	//
	// See [Numeric redaction](https://cloud.ibm.com/docs/services/speech-to-text/output.html#redaction).
	Redaction *bool `json:"redaction,omitempty"`

	// The format (MIME type) of the audio. For more information about specifying an audio format, see **Audio formats
	// (content types)** in the method description.
	ContentType *string `json:"Content-Type,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the CreateJobOptions.Model property.
// The identifier of the model that is to be used for the recognition request. See [Languages and
// models](https://cloud.ibm.com/docs/services/speech-to-text/models.html).
const (
	CreateJobOptions_Model_ArArBroadbandmodel           = "ar-AR_BroadbandModel"
	CreateJobOptions_Model_DeDeBroadbandmodel           = "de-DE_BroadbandModel"
	CreateJobOptions_Model_DeDeNarrowbandmodel          = "de-DE_NarrowbandModel"
	CreateJobOptions_Model_EnGbBroadbandmodel           = "en-GB_BroadbandModel"
	CreateJobOptions_Model_EnGbNarrowbandmodel          = "en-GB_NarrowbandModel"
	CreateJobOptions_Model_EnUsBroadbandmodel           = "en-US_BroadbandModel"
	CreateJobOptions_Model_EnUsNarrowbandmodel          = "en-US_NarrowbandModel"
	CreateJobOptions_Model_EnUsShortformNarrowbandmodel = "en-US_ShortForm_NarrowbandModel"
	CreateJobOptions_Model_EsEsBroadbandmodel           = "es-ES_BroadbandModel"
	CreateJobOptions_Model_EsEsNarrowbandmodel          = "es-ES_NarrowbandModel"
	CreateJobOptions_Model_FrFrBroadbandmodel           = "fr-FR_BroadbandModel"
	CreateJobOptions_Model_FrFrNarrowbandmodel          = "fr-FR_NarrowbandModel"
	CreateJobOptions_Model_JaJpBroadbandmodel           = "ja-JP_BroadbandModel"
	CreateJobOptions_Model_JaJpNarrowbandmodel          = "ja-JP_NarrowbandModel"
	CreateJobOptions_Model_KoKrBroadbandmodel           = "ko-KR_BroadbandModel"
	CreateJobOptions_Model_KoKrNarrowbandmodel          = "ko-KR_NarrowbandModel"
	CreateJobOptions_Model_PtBrBroadbandmodel           = "pt-BR_BroadbandModel"
	CreateJobOptions_Model_PtBrNarrowbandmodel          = "pt-BR_NarrowbandModel"
	CreateJobOptions_Model_ZhCnBroadbandmodel           = "zh-CN_BroadbandModel"
	CreateJobOptions_Model_ZhCnNarrowbandmodel          = "zh-CN_NarrowbandModel"
)

// Constants associated with the CreateJobOptions.Events property.
// If the job includes a callback URL, a comma-separated list of notification events to which to subscribe. Valid events
// are
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
const (
	CreateJobOptions_Events_RecognitionsCompleted            = "recognitions.completed"
	CreateJobOptions_Events_RecognitionsCompletedWithResults = "recognitions.completed_with_results"
	CreateJobOptions_Events_RecognitionsFailed               = "recognitions.failed"
	CreateJobOptions_Events_RecognitionsStarted              = "recognitions.started"
)

// Constants associated with the CreateJobOptions.ContentType property.
// The format (MIME type) of the audio. For more information about specifying an audio format, see **Audio formats
// (content types)** in the method description.
const (
	CreateJobOptions_ContentType_ApplicationOctetStream = "application/octet-stream"
	CreateJobOptions_ContentType_AudioAlaw              = "audio/alaw"
	CreateJobOptions_ContentType_AudioBasic             = "audio/basic"
	CreateJobOptions_ContentType_AudioFlac              = "audio/flac"
	CreateJobOptions_ContentType_AudioG729              = "audio/g729"
	CreateJobOptions_ContentType_AudioL16               = "audio/l16"
	CreateJobOptions_ContentType_AudioMp3               = "audio/mp3"
	CreateJobOptions_ContentType_AudioMpeg              = "audio/mpeg"
	CreateJobOptions_ContentType_AudioMulaw             = "audio/mulaw"
	CreateJobOptions_ContentType_AudioOgg               = "audio/ogg"
	CreateJobOptions_ContentType_AudioOggCodecsOpus     = "audio/ogg;codecs=opus"
	CreateJobOptions_ContentType_AudioOggCodecsVorbis   = "audio/ogg;codecs=vorbis"
	CreateJobOptions_ContentType_AudioWav               = "audio/wav"
	CreateJobOptions_ContentType_AudioWebm              = "audio/webm"
	CreateJobOptions_ContentType_AudioWebmCodecsOpus    = "audio/webm;codecs=opus"
	CreateJobOptions_ContentType_AudioWebmCodecsVorbis  = "audio/webm;codecs=vorbis"
)

// NewCreateJobOptions : Instantiate CreateJobOptions
func (speechToText *SpeechToTextV1) NewCreateJobOptions(audio io.ReadCloser) *CreateJobOptions {
	return &CreateJobOptions{
		Audio: &audio,
	}
}

// SetAudio : Allow user to set Audio
func (options *CreateJobOptions) SetAudio(audio io.ReadCloser) *CreateJobOptions {
	options.Audio = &audio
	return options
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

// SetLanguageCustomizationID : Allow user to set LanguageCustomizationID
func (options *CreateJobOptions) SetLanguageCustomizationID(languageCustomizationID string) *CreateJobOptions {
	options.LanguageCustomizationID = core.StringPtr(languageCustomizationID)
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

// SetCustomizationID : Allow user to set CustomizationID
func (options *CreateJobOptions) SetCustomizationID(customizationID string) *CreateJobOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetGrammarName : Allow user to set GrammarName
func (options *CreateJobOptions) SetGrammarName(grammarName string) *CreateJobOptions {
	options.GrammarName = core.StringPtr(grammarName)
	return options
}

// SetRedaction : Allow user to set Redaction
func (options *CreateJobOptions) SetRedaction(redaction bool) *CreateJobOptions {
	options.Redaction = core.BoolPtr(redaction)
	return options
}

// SetContentType : Allow user to set ContentType
func (options *CreateJobOptions) SetContentType(contentType string) *CreateJobOptions {
	options.ContentType = core.StringPtr(contentType)
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
	// can be used only with the base model that it customizes.
	//
	// To determine whether a base model supports language model customization, use the **Get a model** method and check
	// that the attribute `custom_language_model` is set to `true`. You can also refer to [Language support for
	// customization](https://cloud.ibm.com/docs/services/speech-to-text/custom.html#languageSupport).
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

// Constants associated with the CreateLanguageModelOptions.BaseModelName property.
// The name of the base language model that is to be customized by the new custom language model. The new custom model
// can be used only with the base model that it customizes.
//
// To determine whether a base model supports language model customization, use the **Get a model** method and check
// that the attribute `custom_language_model` is set to `true`. You can also refer to [Language support for
// customization](https://cloud.ibm.com/docs/services/speech-to-text/custom.html#languageSupport).
const (
	CreateLanguageModelOptions_BaseModelName_DeDeBroadbandmodel           = "de-DE_BroadbandModel"
	CreateLanguageModelOptions_BaseModelName_DeDeNarrowbandmodel          = "de-DE_NarrowbandModel"
	CreateLanguageModelOptions_BaseModelName_EnGbBroadbandmodel           = "en-GB_BroadbandModel"
	CreateLanguageModelOptions_BaseModelName_EnGbNarrowbandmodel          = "en-GB_NarrowbandModel"
	CreateLanguageModelOptions_BaseModelName_EnUsBroadbandmodel           = "en-US_BroadbandModel"
	CreateLanguageModelOptions_BaseModelName_EnUsNarrowbandmodel          = "en-US_NarrowbandModel"
	CreateLanguageModelOptions_BaseModelName_EnUsShortformNarrowbandmodel = "en-US_ShortForm_NarrowbandModel"
	CreateLanguageModelOptions_BaseModelName_EsEsBroadbandmodel           = "es-ES_BroadbandModel"
	CreateLanguageModelOptions_BaseModelName_EsEsNarrowbandmodel          = "es-ES_NarrowbandModel"
	CreateLanguageModelOptions_BaseModelName_FrFrBroadbandmodel           = "fr-FR_BroadbandModel"
	CreateLanguageModelOptions_BaseModelName_FrFrNarrowbandmodel          = "fr-FR_NarrowbandModel"
	CreateLanguageModelOptions_BaseModelName_JaJpBroadbandmodel           = "ja-JP_BroadbandModel"
	CreateLanguageModelOptions_BaseModelName_JaJpNarrowbandmodel          = "ja-JP_NarrowbandModel"
	CreateLanguageModelOptions_BaseModelName_KoKrBroadbandmodel           = "ko-KR_BroadbandModel"
	CreateLanguageModelOptions_BaseModelName_KoKrNarrowbandmodel          = "ko-KR_NarrowbandModel"
	CreateLanguageModelOptions_BaseModelName_PtBrBroadbandmodel           = "pt-BR_BroadbandModel"
	CreateLanguageModelOptions_BaseModelName_PtBrNarrowbandmodel          = "pt-BR_NarrowbandModel"
)

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
	// Omit this parameter for the **Add a custom word** method.
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

	// The customization ID (GUID) of the custom acoustic model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
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

	// The customization ID (GUID) of the custom acoustic model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
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

	// The customization ID (GUID) of the custom language model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
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

// DeleteGrammarOptions : The deleteGrammar options.
type DeleteGrammarOptions struct {

	// The customization ID (GUID) of the custom language model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// The name of the grammar for the custom language model.
	GrammarName *string `json:"grammar_name" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteGrammarOptions : Instantiate DeleteGrammarOptions
func (speechToText *SpeechToTextV1) NewDeleteGrammarOptions(customizationID string, grammarName string) *DeleteGrammarOptions {
	return &DeleteGrammarOptions{
		CustomizationID: core.StringPtr(customizationID),
		GrammarName:     core.StringPtr(grammarName),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *DeleteGrammarOptions) SetCustomizationID(customizationID string) *DeleteGrammarOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetGrammarName : Allow user to set GrammarName
func (options *DeleteGrammarOptions) SetGrammarName(grammarName string) *DeleteGrammarOptions {
	options.GrammarName = core.StringPtr(grammarName)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteGrammarOptions) SetHeaders(param map[string]string) *DeleteGrammarOptions {
	options.Headers = param
	return options
}

// DeleteJobOptions : The deleteJob options.
type DeleteJobOptions struct {

	// The identifier of the asynchronous job that is to be used for the request. You must make the request with
	// credentials for the instance of the service that owns the job.
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

	// The customization ID (GUID) of the custom language model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
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

	// The customization ID (GUID) of the custom language model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// The custom word that is to be deleted from the custom language model. URL-encode the word if it includes non-ASCII
	// characters. For more information, see [Character
	// encoding](https://cloud.ibm.com/docs/services/speech-to-text/language-resource.html#charEncoding).
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

	// The customization ID (GUID) of the custom acoustic model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
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

	// The customization ID (GUID) of the custom acoustic model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
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

	// The customization ID (GUID) of the custom language model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
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

// GetGrammarOptions : The getGrammar options.
type GetGrammarOptions struct {

	// The customization ID (GUID) of the custom language model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// The name of the grammar for the custom language model.
	GrammarName *string `json:"grammar_name" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetGrammarOptions : Instantiate GetGrammarOptions
func (speechToText *SpeechToTextV1) NewGetGrammarOptions(customizationID string, grammarName string) *GetGrammarOptions {
	return &GetGrammarOptions{
		CustomizationID: core.StringPtr(customizationID),
		GrammarName:     core.StringPtr(grammarName),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *GetGrammarOptions) SetCustomizationID(customizationID string) *GetGrammarOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetGrammarName : Allow user to set GrammarName
func (options *GetGrammarOptions) SetGrammarName(grammarName string) *GetGrammarOptions {
	options.GrammarName = core.StringPtr(grammarName)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetGrammarOptions) SetHeaders(param map[string]string) *GetGrammarOptions {
	options.Headers = param
	return options
}

// GetLanguageModelOptions : The getLanguageModel options.
type GetLanguageModelOptions struct {

	// The customization ID (GUID) of the custom language model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
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

	// The identifier of the model in the form of its name from the output of the **Get a model** method.
	ModelID *string `json:"model_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the GetModelOptions.ModelID property.
// The identifier of the model in the form of its name from the output of the **Get a model** method.
const (
	GetModelOptions_ModelID_ArArBroadbandmodel           = "ar-AR_BroadbandModel"
	GetModelOptions_ModelID_DeDeBroadbandmodel           = "de-DE_BroadbandModel"
	GetModelOptions_ModelID_DeDeNarrowbandmodel          = "de-DE_NarrowbandModel"
	GetModelOptions_ModelID_EnGbBroadbandmodel           = "en-GB_BroadbandModel"
	GetModelOptions_ModelID_EnGbNarrowbandmodel          = "en-GB_NarrowbandModel"
	GetModelOptions_ModelID_EnUsBroadbandmodel           = "en-US_BroadbandModel"
	GetModelOptions_ModelID_EnUsNarrowbandmodel          = "en-US_NarrowbandModel"
	GetModelOptions_ModelID_EnUsShortformNarrowbandmodel = "en-US_ShortForm_NarrowbandModel"
	GetModelOptions_ModelID_EsEsBroadbandmodel           = "es-ES_BroadbandModel"
	GetModelOptions_ModelID_EsEsNarrowbandmodel          = "es-ES_NarrowbandModel"
	GetModelOptions_ModelID_FrFrBroadbandmodel           = "fr-FR_BroadbandModel"
	GetModelOptions_ModelID_FrFrNarrowbandmodel          = "fr-FR_NarrowbandModel"
	GetModelOptions_ModelID_JaJpBroadbandmodel           = "ja-JP_BroadbandModel"
	GetModelOptions_ModelID_JaJpNarrowbandmodel          = "ja-JP_NarrowbandModel"
	GetModelOptions_ModelID_KoKrBroadbandmodel           = "ko-KR_BroadbandModel"
	GetModelOptions_ModelID_KoKrNarrowbandmodel          = "ko-KR_NarrowbandModel"
	GetModelOptions_ModelID_PtBrBroadbandmodel           = "pt-BR_BroadbandModel"
	GetModelOptions_ModelID_PtBrNarrowbandmodel          = "pt-BR_NarrowbandModel"
	GetModelOptions_ModelID_ZhCnBroadbandmodel           = "zh-CN_BroadbandModel"
	GetModelOptions_ModelID_ZhCnNarrowbandmodel          = "zh-CN_NarrowbandModel"
)

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

	// The customization ID (GUID) of the custom language model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// The custom word that is to be read from the custom language model. URL-encode the word if it includes non-ASCII
	// characters. For more information, see [Character
	// encoding](https://cloud.ibm.com/docs/services/speech-to-text/language-resource.html#charEncoding).
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

// Grammar : Grammar struct
type Grammar struct {

	// The name of the grammar.
	Name *string `json:"name" validate:"required"`

	// The number of OOV words in the grammar. The value is `0` while the grammar is being processed.
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
	Grammar_Status_Analyzed       = "analyzed"
	Grammar_Status_BeingProcessed = "being_processed"
	Grammar_Status_Undetermined   = "undetermined"
)

// Grammars : Grammars struct
type Grammars struct {

	// An array of `Grammar` objects that provides information about the grammars for the custom model. The array is empty
	// if the custom model has no grammars.
	Grammars []Grammar `json:"grammars" validate:"required"`
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

	// The GUID of the credentials for the instance of the service that owns the custom language model.
	Owner *string `json:"owner,omitempty"`

	// The name of the custom language model.
	Name *string `json:"name,omitempty"`

	// The description of the custom language model.
	Description *string `json:"description,omitempty"`

	// The name of the language model for which the custom language model was created.
	BaseModelName *string `json:"base_model_name,omitempty"`

	// The current status of the custom language model:
	// * `pending`: The model was created but is waiting either for training data to be added or for the service to finish
	// analyzing added data.
	// * `ready`: The model contains data and is ready to be trained.
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
// * `pending`: The model was created but is waiting either for training data to be added or for the service to finish
// analyzing added data.
// * `ready`: The model contains data and is ready to be trained.
// * `training`: The model is currently being trained.
// * `available`: The model is trained and ready to use.
// * `upgrading`: The model is currently being upgraded.
// * `failed`: Training of the model failed.
const (
	LanguageModel_Status_Available = "available"
	LanguageModel_Status_Failed    = "failed"
	LanguageModel_Status_Pending   = "pending"
	LanguageModel_Status_Ready     = "ready"
	LanguageModel_Status_Training  = "training"
	LanguageModel_Status_Upgrading = "upgrading"
)

// LanguageModels : LanguageModels struct
type LanguageModels struct {

	// An array of `LanguageModel` objects that provides information about each available custom language model. The array
	// is empty if the requesting credentials own no custom language models (if no language is specified) or own no custom
	// language models for the specified language.
	Customizations []LanguageModel `json:"customizations" validate:"required"`
}

// ListAcousticModelsOptions : The listAcousticModels options.
type ListAcousticModelsOptions struct {

	// The identifier of the language for which custom language or custom acoustic models are to be returned (for example,
	// `en-US`). Omit the parameter to see all custom language or custom acoustic models that are owned by the requesting
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

	// The customization ID (GUID) of the custom acoustic model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
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

	// The customization ID (GUID) of the custom language model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
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

// ListGrammarsOptions : The listGrammars options.
type ListGrammarsOptions struct {

	// The customization ID (GUID) of the custom language model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListGrammarsOptions : Instantiate ListGrammarsOptions
func (speechToText *SpeechToTextV1) NewListGrammarsOptions(customizationID string) *ListGrammarsOptions {
	return &ListGrammarsOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *ListGrammarsOptions) SetCustomizationID(customizationID string) *ListGrammarsOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListGrammarsOptions) SetHeaders(param map[string]string) *ListGrammarsOptions {
	options.Headers = param
	return options
}

// ListLanguageModelsOptions : The listLanguageModels options.
type ListLanguageModelsOptions struct {

	// The identifier of the language for which custom language or custom acoustic models are to be returned (for example,
	// `en-US`). Omit the parameter to see all custom language or custom acoustic models that are owned by the requesting
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

	// The customization ID (GUID) of the custom language model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// The type of words to be listed from the custom language model's words resource:
	// * `all` (the default) shows all words.
	// * `user` shows only custom words that were added or modified by the user directly.
	// * `corpora` shows only OOV that were extracted from corpora.
	// * `grammars` shows only OOV words that are recognized by grammars.
	WordType *string `json:"word_type,omitempty"`

	// Indicates the order in which the words are to be listed, `alphabetical` or by `count`. You can prepend an optional
	// `+` or `-` to an argument to indicate whether the results are to be sorted in ascending or descending order. By
	// default, words are sorted in ascending alphabetical order. For alphabetical ordering, the lexicographical precedence
	// is numeric values, uppercase letters, and lowercase letters. For count ordering, values with the same count are
	// ordered alphabetically. With the `curl` command, URL encode the `+` symbol as `%2B`.
	Sort *string `json:"sort,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the ListWordsOptions.WordType property.
// The type of words to be listed from the custom language model's words resource:
// * `all` (the default) shows all words.
// * `user` shows only custom words that were added or modified by the user directly.
// * `corpora` shows only OOV that were extracted from corpora.
// * `grammars` shows only OOV words that are recognized by grammars.
const (
	ListWordsOptions_WordType_All      = "all"
	ListWordsOptions_WordType_Corpora  = "corpora"
	ListWordsOptions_WordType_Grammars = "grammars"
	ListWordsOptions_WordType_User     = "user"
)

// Constants associated with the ListWordsOptions.Sort property.
// Indicates the order in which the words are to be listed, `alphabetical` or by `count`. You can prepend an optional
// `+` or `-` to an argument to indicate whether the results are to be sorted in ascending or descending order. By
// default, words are sorted in ascending alphabetical order. For alphabetical ordering, the lexicographical precedence
// is numeric values, uppercase letters, and lowercase letters. For count ordering, values with the same count are
// ordered alphabetically. With the `curl` command, URL encode the `+` symbol as `%2B`.
const (
	ListWordsOptions_Sort_Alphabetical = "alphabetical"
	ListWordsOptions_Sort_Count        = "count"
)

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
	// `recognitions.completed_with_results`, the service sent the results with the callback notification. Otherwise, you
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
	RecognitionJob_Status_Completed  = "completed"
	RecognitionJob_Status_Failed     = "failed"
	RecognitionJob_Status_Processing = "processing"
	RecognitionJob_Status_Waiting    = "waiting"
)

// RecognitionJobs : RecognitionJobs struct
type RecognitionJobs struct {

	// An array of `RecognitionJob` objects that provides the status for each of the user's current jobs. The array is
	// empty if the user has no current jobs.
	Recognitions []RecognitionJob `json:"recognitions" validate:"required"`
}

// RecognizeOptions : The recognize options.
type RecognizeOptions struct {

	// The audio to transcribe.
	Audio *io.ReadCloser `json:"audio" validate:"required"`

	// The identifier of the model that is to be used for the recognition request. See [Languages and
	// models](https://cloud.ibm.com/docs/services/speech-to-text/models.html).
	Model *string `json:"model,omitempty"`

	// The customization ID (GUID) of a custom language model that is to be used with the recognition request. The base
	// model of the specified custom language model must match the model specified with the `model` parameter. You must
	// make the request with credentials for the instance of the service that owns the custom model. By default, no custom
	// language model is used. See [Custom
	// models](https://cloud.ibm.com/docs/services/speech-to-text/input.html#custom-input).
	//
	// **Note:** Use this parameter instead of the deprecated `customization_id` parameter.
	LanguageCustomizationID *string `json:"language_customization_id,omitempty"`

	// The customization ID (GUID) of a custom acoustic model that is to be used with the recognition request. The base
	// model of the specified custom acoustic model must match the model specified with the `model` parameter. You must
	// make the request with credentials for the instance of the service that owns the custom model. By default, no custom
	// acoustic model is used. See [Custom
	// models](https://cloud.ibm.com/docs/services/speech-to-text/input.html#custom-input).
	AcousticCustomizationID *string `json:"acoustic_customization_id,omitempty"`

	// The version of the specified base model that is to be used with recognition request. Multiple versions of a base
	// model can exist when a model is updated for internal improvements. The parameter is intended primarily for use with
	// custom models that have been upgraded for a new base model. The default value depends on whether the parameter is
	// used with or without a custom model. See [Base model
	// version](https://cloud.ibm.com/docs/services/speech-to-text/input.html#version).
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
	// See [Custom models](https://cloud.ibm.com/docs/services/speech-to-text/input.html#custom-input).
	CustomizationWeight *float64 `json:"customization_weight,omitempty"`

	// The time in seconds after which, if only silence (no speech) is detected in streaming audio, the connection is
	// closed with a 400 error. The parameter is useful for stopping audio submission from a live microphone when a user
	// simply walks away. Use `-1` for infinity. See [Inactivity
	// timeout](https://cloud.ibm.com/docs/services/speech-to-text/input.html#timeouts-inactivity).
	InactivityTimeout *int64 `json:"inactivity_timeout,omitempty"`

	// An array of keyword strings to spot in the audio. Each keyword string can include one or more string tokens.
	// Keywords are spotted only in the final results, not in interim hypotheses. If you specify any keywords, you must
	// also specify a keywords threshold. You can spot a maximum of 1000 keywords. Omit the parameter or specify an empty
	// array if you do not need to spot keywords. See [Keyword
	// spotting](https://cloud.ibm.com/docs/services/speech-to-text/output.html#keyword_spotting).
	Keywords []string `json:"keywords,omitempty"`

	// A confidence value that is the lower bound for spotting a keyword. A word is considered to match a keyword if its
	// confidence is greater than or equal to the threshold. Specify a probability between 0.0 and 1.0. If you specify a
	// threshold, you must also specify one or more keywords. The service performs no keyword spotting if you omit either
	// parameter. See [Keyword spotting](https://cloud.ibm.com/docs/services/speech-to-text/output.html#keyword_spotting).
	KeywordsThreshold *float32 `json:"keywords_threshold,omitempty"`

	// The maximum number of alternative transcripts that the service is to return. By default, the service returns a
	// single transcript. If you specify a value of `0`, the service uses the default value, `1`. See [Maximum
	// alternatives](https://cloud.ibm.com/docs/services/speech-to-text/output.html#max_alternatives).
	MaxAlternatives *int64 `json:"max_alternatives,omitempty"`

	// A confidence value that is the lower bound for identifying a hypothesis as a possible word alternative (also known
	// as "Confusion Networks"). An alternative word is considered if its confidence is greater than or equal to the
	// threshold. Specify a probability between 0.0 and 1.0. By default, the service computes no alternative words. See
	// [Word alternatives](https://cloud.ibm.com/docs/services/speech-to-text/output.html#word_alternatives).
	WordAlternativesThreshold *float32 `json:"word_alternatives_threshold,omitempty"`

	// If `true`, the service returns a confidence measure in the range of 0.0 to 1.0 for each word. By default, the
	// service returns no word confidence scores. See [Word
	// confidence](https://cloud.ibm.com/docs/services/speech-to-text/output.html#word_confidence).
	WordConfidence *bool `json:"word_confidence,omitempty"`

	// If `true`, the service returns time alignment for each word. By default, no timestamps are returned. See [Word
	// timestamps](https://cloud.ibm.com/docs/services/speech-to-text/output.html#word_timestamps).
	Timestamps *bool `json:"timestamps,omitempty"`

	// If `true`, the service filters profanity from all output except for keyword results by replacing inappropriate words
	// with a series of asterisks. Set the parameter to `false` to return results with no censoring. Applies to US English
	// transcription only. See [Profanity
	// filtering](https://cloud.ibm.com/docs/services/speech-to-text/output.html#profanity_filter).
	ProfanityFilter *bool `json:"profanity_filter,omitempty"`

	// If `true`, the service converts dates, times, series of digits and numbers, phone numbers, currency values, and
	// internet addresses into more readable, conventional representations in the final transcript of a recognition
	// request. For US English, the service also converts certain keyword strings to punctuation symbols. By default, the
	// service performs no smart formatting.
	//
	// **Note:** Applies to US English, Japanese, and Spanish transcription only.
	//
	// See [Smart formatting](https://cloud.ibm.com/docs/services/speech-to-text/output.html#smart_formatting).
	SmartFormatting *bool `json:"smart_formatting,omitempty"`

	// If `true`, the response includes labels that identify which words were spoken by which participants in a
	// multi-person exchange. By default, the service returns no speaker labels. Setting `speaker_labels` to `true` forces
	// the `timestamps` parameter to be `true`, regardless of whether you specify `false` for the parameter.
	//
	// **Note:** Applies to US English, Japanese, and Spanish transcription only. To determine whether a language model
	// supports speaker labels, you can also use the **Get a model** method and check that the attribute `speaker_labels`
	// is set to `true`.
	//
	// See [Speaker labels](https://cloud.ibm.com/docs/services/speech-to-text/output.html#speaker_labels).
	SpeakerLabels *bool `json:"speaker_labels,omitempty"`

	// **Deprecated.** Use the `language_customization_id` parameter to specify the customization ID (GUID) of a custom
	// language model that is to be used with the recognition request. Do not specify both parameters with a request.
	CustomizationID *string `json:"customization_id,omitempty"`

	// The name of a grammar that is to be used with the recognition request. If you specify a grammar, you must also use
	// the `language_customization_id` parameter to specify the name of the custom language model for which the grammar is
	// defined. The service recognizes only strings that are recognized by the specified grammar; it does not recognize
	// other custom words from the model's words resource. See
	// [Grammars](https://cloud.ibm.com/docs/services/speech-to-text/input.html#grammars-input).
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
	// **Note:** Applies to US English, Japanese, and Korean transcription only.
	//
	// See [Numeric redaction](https://cloud.ibm.com/docs/services/speech-to-text/output.html#redaction).
	Redaction *bool `json:"redaction,omitempty"`

	// The format (MIME type) of the audio. For more information about specifying an audio format, see **Audio formats
	// (content types)** in the method description.
	ContentType *string `json:"Content-Type,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the RecognizeOptions.Model property.
// The identifier of the model that is to be used for the recognition request. See [Languages and
// models](https://cloud.ibm.com/docs/services/speech-to-text/models.html).
const (
	RecognizeOptions_Model_ArArBroadbandmodel           = "ar-AR_BroadbandModel"
	RecognizeOptions_Model_DeDeBroadbandmodel           = "de-DE_BroadbandModel"
	RecognizeOptions_Model_DeDeNarrowbandmodel          = "de-DE_NarrowbandModel"
	RecognizeOptions_Model_EnGbBroadbandmodel           = "en-GB_BroadbandModel"
	RecognizeOptions_Model_EnGbNarrowbandmodel          = "en-GB_NarrowbandModel"
	RecognizeOptions_Model_EnUsBroadbandmodel           = "en-US_BroadbandModel"
	RecognizeOptions_Model_EnUsNarrowbandmodel          = "en-US_NarrowbandModel"
	RecognizeOptions_Model_EnUsShortformNarrowbandmodel = "en-US_ShortForm_NarrowbandModel"
	RecognizeOptions_Model_EsEsBroadbandmodel           = "es-ES_BroadbandModel"
	RecognizeOptions_Model_EsEsNarrowbandmodel          = "es-ES_NarrowbandModel"
	RecognizeOptions_Model_FrFrBroadbandmodel           = "fr-FR_BroadbandModel"
	RecognizeOptions_Model_FrFrNarrowbandmodel          = "fr-FR_NarrowbandModel"
	RecognizeOptions_Model_JaJpBroadbandmodel           = "ja-JP_BroadbandModel"
	RecognizeOptions_Model_JaJpNarrowbandmodel          = "ja-JP_NarrowbandModel"
	RecognizeOptions_Model_KoKrBroadbandmodel           = "ko-KR_BroadbandModel"
	RecognizeOptions_Model_KoKrNarrowbandmodel          = "ko-KR_NarrowbandModel"
	RecognizeOptions_Model_PtBrBroadbandmodel           = "pt-BR_BroadbandModel"
	RecognizeOptions_Model_PtBrNarrowbandmodel          = "pt-BR_NarrowbandModel"
	RecognizeOptions_Model_ZhCnBroadbandmodel           = "zh-CN_BroadbandModel"
	RecognizeOptions_Model_ZhCnNarrowbandmodel          = "zh-CN_NarrowbandModel"
)

// Constants associated with the RecognizeOptions.ContentType property.
// The format (MIME type) of the audio. For more information about specifying an audio format, see **Audio formats
// (content types)** in the method description.
const (
	RecognizeOptions_ContentType_ApplicationOctetStream = "application/octet-stream"
	RecognizeOptions_ContentType_AudioAlaw              = "audio/alaw"
	RecognizeOptions_ContentType_AudioBasic             = "audio/basic"
	RecognizeOptions_ContentType_AudioFlac              = "audio/flac"
	RecognizeOptions_ContentType_AudioG729              = "audio/g729"
	RecognizeOptions_ContentType_AudioL16               = "audio/l16"
	RecognizeOptions_ContentType_AudioMp3               = "audio/mp3"
	RecognizeOptions_ContentType_AudioMpeg              = "audio/mpeg"
	RecognizeOptions_ContentType_AudioMulaw             = "audio/mulaw"
	RecognizeOptions_ContentType_AudioOgg               = "audio/ogg"
	RecognizeOptions_ContentType_AudioOggCodecsOpus     = "audio/ogg;codecs=opus"
	RecognizeOptions_ContentType_AudioOggCodecsVorbis   = "audio/ogg;codecs=vorbis"
	RecognizeOptions_ContentType_AudioWav               = "audio/wav"
	RecognizeOptions_ContentType_AudioWebm              = "audio/webm"
	RecognizeOptions_ContentType_AudioWebmCodecsOpus    = "audio/webm;codecs=opus"
	RecognizeOptions_ContentType_AudioWebmCodecsVorbis  = "audio/webm;codecs=vorbis"
)

// NewRecognizeOptions : Instantiate RecognizeOptions
func (speechToText *SpeechToTextV1) NewRecognizeOptions(audio io.ReadCloser) *RecognizeOptions {
	return &RecognizeOptions{
		Audio: &audio,
	}
}

// SetAudio : Allow user to set Audio
func (options *RecognizeOptions) SetAudio(audio io.ReadCloser) *RecognizeOptions {
	options.Audio = &audio
	return options
}

// SetModel : Allow user to set Model
func (options *RecognizeOptions) SetModel(model string) *RecognizeOptions {
	options.Model = core.StringPtr(model)
	return options
}

// SetLanguageCustomizationID : Allow user to set LanguageCustomizationID
func (options *RecognizeOptions) SetLanguageCustomizationID(languageCustomizationID string) *RecognizeOptions {
	options.LanguageCustomizationID = core.StringPtr(languageCustomizationID)
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

// SetCustomizationID : Allow user to set CustomizationID
func (options *RecognizeOptions) SetCustomizationID(customizationID string) *RecognizeOptions {
	options.CustomizationID = core.StringPtr(customizationID)
	return options
}

// SetGrammarName : Allow user to set GrammarName
func (options *RecognizeOptions) SetGrammarName(grammarName string) *RecognizeOptions {
	options.GrammarName = core.StringPtr(grammarName)
	return options
}

// SetRedaction : Allow user to set Redaction
func (options *RecognizeOptions) SetRedaction(redaction bool) *RecognizeOptions {
	options.Redaction = core.BoolPtr(redaction)
	return options
}

// SetContentType : Allow user to set ContentType
func (options *RecognizeOptions) SetContentType(contentType string) *RecognizeOptions {
	options.ContentType = core.StringPtr(contentType)
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
	// * `created`: The service successfully white-listed the callback URL as a result of the call.
	// * `already created`: The URL was already white-listed.
	Status *string `json:"status" validate:"required"`

	// The callback URL that is successfully registered.
	URL *string `json:"url" validate:"required"`
}

// Constants associated with the RegisterStatus.Status property.
// The current status of the job:
// * `created`: The service successfully white-listed the callback URL as a result of the call.
// * `already created`: The URL was already white-listed.
const (
	RegisterStatus_Status_AlreadyCreated = "already created"
	RegisterStatus_Status_Created        = "created"
)

// ResetAcousticModelOptions : The resetAcousticModel options.
type ResetAcousticModelOptions struct {

	// The customization ID (GUID) of the custom acoustic model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
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

	// The customization ID (GUID) of the custom language model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
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

	// Describes the additional service features that are supported with the model.
	SupportedFeatures *SupportedFeatures `json:"supported_features" validate:"required"`

	// A brief description of the model.
	Description *string `json:"description" validate:"required"`
}

// SpeechModels : SpeechModels struct
type SpeechModels struct {

	// An array of `SpeechModel` objects that provides information about each available model.
	Models []SpeechModel `json:"models" validate:"required"`
}

// SpeechRecognitionAlternative : SpeechRecognitionAlternative struct
type SpeechRecognitionAlternative struct {

	// A transcription of the audio.
	Transcript *string `json:"transcript" validate:"required"`

	// A score that indicates the service's confidence in the transcript in the range of 0.0 to 1.0. A confidence score is
	// returned only for the best alternative and only with results marked as final.
	Confidence *float64 `json:"confidence,omitempty"`

	// Time alignments for each word from the transcript as a list of lists. Each inner list consists of three elements:
	// the word followed by its start and end time in seconds, for example: `[["hello",0.0,1.2],["world",1.2,2.5]]`.
	// Timestamps are returned only for the best alternative.
	Timestamps []interface{} `json:"timestamps,omitempty"`

	// A confidence score for each word of the transcript as a list of lists. Each inner list consists of two elements: the
	// word and its confidence score in the range of 0.0 to 1.0, for example: `[["hello",0.95],["world",0.866]]`.
	// Confidence scores are returned only for the best alternative and only with results marked as final.
	WordConfidence []interface{} `json:"word_confidence,omitempty"`
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
	// `keywords_threshold` are specified. The value for each key is an array of matches spotted in the audio for that
	// keyword. Each match is described by a `KeywordResult` object. A keyword for which no matches are found is omitted
	// from the dictionary. The dictionary is omitted entirely if no matches are found for any keywords.
	KeywordsResult map[string][]KeywordResult `json:"keywords_result,omitempty"`

	// An array of alternative hypotheses found for words of the input audio if a `word_alternatives_threshold` is
	// specified.
	WordAlternatives []WordAlternativeResults `json:"word_alternatives,omitempty"`
}

// SpeechRecognitionResults : SpeechRecognitionResults struct
type SpeechRecognitionResults struct {

	// An array of `SpeechRecognitionResult` objects that can include interim and final results (interim results are
	// returned only if supported by the method). Final results are guaranteed not to change; interim results might be
	// replaced by further interim results and final results. The service periodically sends updates to the results list;
	// the `result_index` is set to the lowest index in the array that has changed; it is incremented for new results.
	Results []SpeechRecognitionResult `json:"results,omitempty"`

	// An index that indicates a change point in the `results` array. The service increments the index only for additional
	// results that it sends for new audio for the same request.
	ResultIndex *int64 `json:"result_index,omitempty"`

	// An array of `SpeakerLabelsResult` objects that identifies which words were spoken by which speakers in a
	// multi-person exchange. The array is returned only if the `speaker_labels` parameter is `true`. When interim results
	// are also requested for methods that support them, it is possible for a `SpeechRecognitionResults` object to include
	// only the `speaker_labels` field.
	SpeakerLabels []SpeakerLabelsResult `json:"speaker_labels,omitempty"`

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

// SupportedFeatures : Describes the additional service features that are supported with the model.
type SupportedFeatures struct {

	// Indicates whether the customization interface can be used to create a custom language model based on the language
	// model.
	CustomLanguageModel *bool `json:"custom_language_model" validate:"required"`

	// Indicates whether the `speaker_labels` parameter can be used with the language model.
	SpeakerLabels *bool `json:"speaker_labels" validate:"required"`
}

// TrainAcousticModelOptions : The trainAcousticModel options.
type TrainAcousticModelOptions struct {

	// The customization ID (GUID) of the custom acoustic model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// The customization ID (GUID) of a custom language model that is to be used during training of the custom acoustic
	// model. Specify a custom language model that has been trained with verbatim transcriptions of the audio resources or
	// that contains words that are relevant to the contents of the audio resources. The custom language model must be
	// based on the same version of the same base model as the custom acoustic model. The credentials specified with the
	// request must own both custom models.
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

	// The customization ID (GUID) of the custom language model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// The type of words from the custom language model's words resource on which to train the model:
	// * `all` (the default) trains the model on all new words, regardless of whether they were extracted from corpora or
	// grammars or were added or modified by the user.
	// * `user` trains the model only on new words that were added or modified by the user directly. The model is not
	// trained on new words extracted from corpora or grammars.
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

// Constants associated with the TrainLanguageModelOptions.WordTypeToAdd property.
// The type of words from the custom language model's words resource on which to train the model:
// * `all` (the default) trains the model on all new words, regardless of whether they were extracted from corpora or
// grammars or were added or modified by the user.
// * `user` trains the model only on new words that were added or modified by the user directly. The model is not
// trained on new words extracted from corpora or grammars.
const (
	TrainLanguageModelOptions_WordTypeToAdd_All  = "all"
	TrainLanguageModelOptions_WordTypeToAdd_User = "user"
)

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

	// The customization ID (GUID) of the custom acoustic model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// If the custom acoustic model was trained with a custom language model, the customization ID (GUID) of that custom
	// language model. The custom language model must be upgraded before the custom acoustic model can be upgraded. The
	// credentials specified with the request must own both custom models.
	CustomLanguageModelID *string `json:"custom_language_model_id,omitempty"`

	// If `true`, forces the upgrade of a custom acoustic model for which no input data has been modified since it was last
	// trained. Use this parameter only to force the upgrade of a custom acoustic model that is trained with a custom
	// language model, and only if you receive a 400 response code and the message `No input data modified since last
	// training`. See [Upgrading a custom acoustic
	// model](https://cloud.ibm.com/docs/services/speech-to-text/custom-upgrade.html#upgradeAcoustic).
	Force *bool `json:"force,omitempty"`

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

// SetForce : Allow user to set Force
func (options *UpgradeAcousticModelOptions) SetForce(force bool) *UpgradeAcousticModelOptions {
	options.Force = core.BoolPtr(force)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpgradeAcousticModelOptions) SetHeaders(param map[string]string) *UpgradeAcousticModelOptions {
	options.Headers = param
	return options
}

// UpgradeLanguageModelOptions : The upgradeLanguageModel options.
type UpgradeLanguageModelOptions struct {

	// The customization ID (GUID) of the custom language model that is to be used for the request. You must make the
	// request with credentials for the instance of the service that owns the custom model.
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

	// An array of `Word` objects that provides information about each word in the custom model's words resource. The array
	// is empty if the custom model has no words.
	Words []Word `json:"words" validate:"required"`
}
