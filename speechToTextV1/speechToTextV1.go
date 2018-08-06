// Package speechToTextV1 : Operations and models for the SpeechToTextV1 service
package speechToTextV1
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
    "io"
    "os"
    "strings"
    req "github.com/parnurzeal/gorequest"
    watson "go-sdk"
)

// SpeechToTextV1 : The SpeechToTextV1 service
type SpeechToTextV1 struct {
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

// NewSpeechToTextV1 : Instantiate SpeechToTextV1
func NewSpeechToTextV1(serviceCreds *ServiceCredentials) (*SpeechToTextV1, error) {
    if serviceCreds.ServiceURL == "" {
        serviceCreds.ServiceURL = "https://stream.watsonplatform.net/speech-to-text/api"
    }

    creds := watson.Credentials(*serviceCreds)
    client, clientErr := watson.NewClient(&creds, "speech_to_text")

    if clientErr != nil {
        return nil, clientErr
    }

    return &SpeechToTextV1{ client: client }, nil
}

// GetModel : Get a model
func (speechToText *SpeechToTextV1) GetModel(options *GetModelOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/models/{model_id}"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{model_id}", options.ModelID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")

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

    response.Result = new(SpeechModel)
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

// GetGetModelResult : Cast result of GetModel operation
func GetGetModelResult(response *watson.WatsonResponse) *SpeechModel {
    result, ok := response.Result.(*SpeechModel)

    if ok {
        return result
    }

    return nil
}

// ListModels : List models
func (speechToText *SpeechToTextV1) ListModels(options *ListModelsOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/models"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")

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

    response.Result = new(SpeechModels)
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

// GetListModelsResult : Cast result of ListModels operation
func GetListModelsResult(response *watson.WatsonResponse) *SpeechModels {
    result, ok := response.Result.(*SpeechModels)

    if ok {
        return result
    }

    return nil
}

// Recognize : Recognize audio
func (speechToText *SpeechToTextV1) Recognize(options *RecognizeOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/recognize"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", fmt.Sprint(options.ContentType))
    if options.IsModelSet {
        request.Query("model=" + fmt.Sprint(options.Model))
    }
    if options.IsCustomizationIDSet {
        request.Query("customization_id=" + fmt.Sprint(options.CustomizationID))
    }
    if options.IsAcousticCustomizationIDSet {
        request.Query("acoustic_customization_id=" + fmt.Sprint(options.AcousticCustomizationID))
    }
    if options.IsBaseModelVersionSet {
        request.Query("base_model_version=" + fmt.Sprint(options.BaseModelVersion))
    }
    if options.IsCustomizationWeightSet {
        request.Query("customization_weight=" + fmt.Sprint(options.CustomizationWeight))
    }
    if options.IsInactivityTimeoutSet {
        request.Query("inactivity_timeout=" + fmt.Sprint(options.InactivityTimeout))
    }
    if options.IsKeywordsSet {
        request.Query("keywords=" + fmt.Sprint(options.Keywords))
    }
    if options.IsKeywordsThresholdSet {
        request.Query("keywords_threshold=" + fmt.Sprint(options.KeywordsThreshold))
    }
    if options.IsMaxAlternativesSet {
        request.Query("max_alternatives=" + fmt.Sprint(options.MaxAlternatives))
    }
    if options.IsWordAlternativesThresholdSet {
        request.Query("word_alternatives_threshold=" + fmt.Sprint(options.WordAlternativesThreshold))
    }
    if options.IsWordConfidenceSet {
        request.Query("word_confidence=" + fmt.Sprint(options.WordConfidence))
    }
    if options.IsTimestampsSet {
        request.Query("timestamps=" + fmt.Sprint(options.Timestamps))
    }
    if options.IsProfanityFilterSet {
        request.Query("profanity_filter=" + fmt.Sprint(options.ProfanityFilter))
    }
    if options.IsSmartFormattingSet {
        request.Query("smart_formatting=" + fmt.Sprint(options.SmartFormatting))
    }
    if options.IsSpeakerLabelsSet {
        request.Query("speaker_labels=" + fmt.Sprint(options.SpeakerLabels))
    }
    audio := new(bytes.Buffer)
    audio.ReadFrom(options.Audio)
    request.Send(audio.String())

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

    response.Result = new(SpeechRecognitionResults)
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

// GetRecognizeResult : Cast result of Recognize operation
func GetRecognizeResult(response *watson.WatsonResponse) *SpeechRecognitionResults {
    result, ok := response.Result.(*SpeechRecognitionResults)

    if ok {
        return result
    }

    return nil
}

// CheckJob : Check a job
func (speechToText *SpeechToTextV1) CheckJob(options *CheckJobOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/recognitions/{id}"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{id}", options.ID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")

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

    response.Result = new(RecognitionJob)
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

// GetCheckJobResult : Cast result of CheckJob operation
func GetCheckJobResult(response *watson.WatsonResponse) *RecognitionJob {
    result, ok := response.Result.(*RecognitionJob)

    if ok {
        return result
    }

    return nil
}

// CheckJobs : Check jobs
func (speechToText *SpeechToTextV1) CheckJobs(options *CheckJobsOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/recognitions"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")

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

    response.Result = new(RecognitionJobs)
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

// GetCheckJobsResult : Cast result of CheckJobs operation
func GetCheckJobsResult(response *watson.WatsonResponse) *RecognitionJobs {
    result, ok := response.Result.(*RecognitionJobs)

    if ok {
        return result
    }

    return nil
}

// CreateJob : Create a job
func (speechToText *SpeechToTextV1) CreateJob(options *CreateJobOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/recognitions"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", fmt.Sprint(options.ContentType))
    if options.IsModelSet {
        request.Query("model=" + fmt.Sprint(options.Model))
    }
    if options.IsCallbackURLSet {
        request.Query("callback_url=" + fmt.Sprint(options.CallbackURL))
    }
    if options.IsEventsSet {
        request.Query("events=" + fmt.Sprint(options.Events))
    }
    if options.IsUserTokenSet {
        request.Query("user_token=" + fmt.Sprint(options.UserToken))
    }
    if options.IsResultsTTLSet {
        request.Query("results_ttl=" + fmt.Sprint(options.ResultsTTL))
    }
    if options.IsCustomizationIDSet {
        request.Query("customization_id=" + fmt.Sprint(options.CustomizationID))
    }
    if options.IsAcousticCustomizationIDSet {
        request.Query("acoustic_customization_id=" + fmt.Sprint(options.AcousticCustomizationID))
    }
    if options.IsBaseModelVersionSet {
        request.Query("base_model_version=" + fmt.Sprint(options.BaseModelVersion))
    }
    if options.IsCustomizationWeightSet {
        request.Query("customization_weight=" + fmt.Sprint(options.CustomizationWeight))
    }
    if options.IsInactivityTimeoutSet {
        request.Query("inactivity_timeout=" + fmt.Sprint(options.InactivityTimeout))
    }
    if options.IsKeywordsSet {
        request.Query("keywords=" + fmt.Sprint(options.Keywords))
    }
    if options.IsKeywordsThresholdSet {
        request.Query("keywords_threshold=" + fmt.Sprint(options.KeywordsThreshold))
    }
    if options.IsMaxAlternativesSet {
        request.Query("max_alternatives=" + fmt.Sprint(options.MaxAlternatives))
    }
    if options.IsWordAlternativesThresholdSet {
        request.Query("word_alternatives_threshold=" + fmt.Sprint(options.WordAlternativesThreshold))
    }
    if options.IsWordConfidenceSet {
        request.Query("word_confidence=" + fmt.Sprint(options.WordConfidence))
    }
    if options.IsTimestampsSet {
        request.Query("timestamps=" + fmt.Sprint(options.Timestamps))
    }
    if options.IsProfanityFilterSet {
        request.Query("profanity_filter=" + fmt.Sprint(options.ProfanityFilter))
    }
    if options.IsSmartFormattingSet {
        request.Query("smart_formatting=" + fmt.Sprint(options.SmartFormatting))
    }
    if options.IsSpeakerLabelsSet {
        request.Query("speaker_labels=" + fmt.Sprint(options.SpeakerLabels))
    }
    audio := new(bytes.Buffer)
    audio.ReadFrom(options.Audio)
    request.Send(audio.String())

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

    response.Result = new(RecognitionJob)
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

// GetCreateJobResult : Cast result of CreateJob operation
func GetCreateJobResult(response *watson.WatsonResponse) *RecognitionJob {
    result, ok := response.Result.(*RecognitionJob)

    if ok {
        return result
    }

    return nil
}

// DeleteJob : Delete a job
func (speechToText *SpeechToTextV1) DeleteJob(options *DeleteJobOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/recognitions/{id}"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{id}", options.ID, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")

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

    res, _, err := request.End()

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


// RegisterCallback : Register a callback
func (speechToText *SpeechToTextV1) RegisterCallback(options *RegisterCallbackOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/register_callback"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("callback_url=" + fmt.Sprint(options.CallbackURL))
    if options.IsUserSecretSet {
        request.Query("user_secret=" + fmt.Sprint(options.UserSecret))
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

    response.Result = new(RegisterStatus)
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

// GetRegisterCallbackResult : Cast result of RegisterCallback operation
func GetRegisterCallbackResult(response *watson.WatsonResponse) *RegisterStatus {
    result, ok := response.Result.(*RegisterStatus)

    if ok {
        return result
    }

    return nil
}

// UnregisterCallback : Unregister a callback
func (speechToText *SpeechToTextV1) UnregisterCallback(options *UnregisterCallbackOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/unregister_callback"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("callback_url=" + fmt.Sprint(options.CallbackURL))

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

    res, _, err := request.End()

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


// CreateLanguageModel : Create a custom language model
func (speechToText *SpeechToTextV1) CreateLanguageModel(options *CreateLanguageModelOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    body := map[string]interface{}{}
    body["name"] = options.Name
    body["base_model_name"] = options.BaseModelName
    if options.IsDialectSet {
        body["dialect"] = options.Dialect
    }
    if options.IsDescriptionSet {
        body["description"] = options.Description
    }
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

    response.Result = new(LanguageModel)
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

// GetCreateLanguageModelResult : Cast result of CreateLanguageModel operation
func GetCreateLanguageModelResult(response *watson.WatsonResponse) *LanguageModel {
    result, ok := response.Result.(*LanguageModel)

    if ok {
        return result
    }

    return nil
}

// DeleteLanguageModel : Delete a custom language model
func (speechToText *SpeechToTextV1) DeleteLanguageModel(options *DeleteLanguageModelOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", options.CustomizationID, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")

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

    res, _, err := request.End()

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


// GetLanguageModel : Get a custom language model
func (speechToText *SpeechToTextV1) GetLanguageModel(options *GetLanguageModelOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", options.CustomizationID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")

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

    response.Result = new(LanguageModel)
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

// GetGetLanguageModelResult : Cast result of GetLanguageModel operation
func GetGetLanguageModelResult(response *watson.WatsonResponse) *LanguageModel {
    result, ok := response.Result.(*LanguageModel)

    if ok {
        return result
    }

    return nil
}

// ListLanguageModels : List custom language models
func (speechToText *SpeechToTextV1) ListLanguageModels(options *ListLanguageModelsOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    if options.IsLanguageSet {
        request.Query("language=" + fmt.Sprint(options.Language))
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

    response.Result = new(LanguageModels)
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

// GetListLanguageModelsResult : Cast result of ListLanguageModels operation
func GetListLanguageModelsResult(response *watson.WatsonResponse) *LanguageModels {
    result, ok := response.Result.(*LanguageModels)

    if ok {
        return result
    }

    return nil
}

// ResetLanguageModel : Reset a custom language model
func (speechToText *SpeechToTextV1) ResetLanguageModel(options *ResetLanguageModelOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}/reset"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", options.CustomizationID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")

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

    res, _, err := request.End()

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


// TrainLanguageModel : Train a custom language model
func (speechToText *SpeechToTextV1) TrainLanguageModel(options *TrainLanguageModelOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}/train"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", options.CustomizationID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    if options.IsWordTypeToAddSet {
        request.Query("word_type_to_add=" + fmt.Sprint(options.WordTypeToAdd))
    }
    if options.IsCustomizationWeightSet {
        request.Query("customization_weight=" + fmt.Sprint(options.CustomizationWeight))
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

    res, _, err := request.End()

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


// UpgradeLanguageModel : Upgrade a custom language model
func (speechToText *SpeechToTextV1) UpgradeLanguageModel(options *UpgradeLanguageModelOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}/upgrade_model"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", options.CustomizationID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")

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

    res, _, err := request.End()

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


// AddCorpus : Add a corpus
func (speechToText *SpeechToTextV1) AddCorpus(options *AddCorpusOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}/corpora/{corpus_name}"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", options.CustomizationID, 1)
    path = strings.Replace(path, "{corpus_name}", options.CorpusName, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    if options.IsAllowOverwriteSet {
        request.Query("allow_overwrite=" + fmt.Sprint(options.AllowOverwrite))
    }
    request.Type("multipart")
    request.SendFile(options.CorpusFile, "", "corpus_file")

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

    res, _, err := request.End()

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


// DeleteCorpus : Delete a corpus
func (speechToText *SpeechToTextV1) DeleteCorpus(options *DeleteCorpusOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}/corpora/{corpus_name}"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", options.CustomizationID, 1)
    path = strings.Replace(path, "{corpus_name}", options.CorpusName, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")

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

    res, _, err := request.End()

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


// GetCorpus : Get a corpus
func (speechToText *SpeechToTextV1) GetCorpus(options *GetCorpusOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}/corpora/{corpus_name}"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", options.CustomizationID, 1)
    path = strings.Replace(path, "{corpus_name}", options.CorpusName, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")

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

    response.Result = new(Corpus)
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

// GetGetCorpusResult : Cast result of GetCorpus operation
func GetGetCorpusResult(response *watson.WatsonResponse) *Corpus {
    result, ok := response.Result.(*Corpus)

    if ok {
        return result
    }

    return nil
}

// ListCorpora : List corpora
func (speechToText *SpeechToTextV1) ListCorpora(options *ListCorporaOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}/corpora"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", options.CustomizationID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")

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

    response.Result = new(Corpora)
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

// GetListCorporaResult : Cast result of ListCorpora operation
func GetListCorporaResult(response *watson.WatsonResponse) *Corpora {
    result, ok := response.Result.(*Corpora)

    if ok {
        return result
    }

    return nil
}

// AddWord : Add a custom word
func (speechToText *SpeechToTextV1) AddWord(options *AddWordOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}/words/{word_name}"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", options.CustomizationID, 1)
    path = strings.Replace(path, "{word_name}", options.WordName, 1)
    request := req.New().Put(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    body := map[string]interface{}{}
    if options.IsWordSet {
        body["word"] = options.Word
    }
    if options.IsSoundsLikeSet {
        body["sounds_like"] = options.SoundsLike
    }
    if options.IsDisplayAsSet {
        body["display_as"] = options.DisplayAs
    }
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

    res, _, err := request.End()

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


// AddWords : Add custom words
func (speechToText *SpeechToTextV1) AddWords(options *AddWordsOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}/words"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", options.CustomizationID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    body := map[string]interface{}{}
    body["words"] = options.Words
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

    res, _, err := request.End()

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


// DeleteWord : Delete a custom word
func (speechToText *SpeechToTextV1) DeleteWord(options *DeleteWordOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}/words/{word_name}"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", options.CustomizationID, 1)
    path = strings.Replace(path, "{word_name}", options.WordName, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")

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

    res, _, err := request.End()

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


// GetWord : Get a custom word
func (speechToText *SpeechToTextV1) GetWord(options *GetWordOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}/words/{word_name}"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", options.CustomizationID, 1)
    path = strings.Replace(path, "{word_name}", options.WordName, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")

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

    response.Result = new(Word)
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

// GetGetWordResult : Cast result of GetWord operation
func GetGetWordResult(response *watson.WatsonResponse) *Word {
    result, ok := response.Result.(*Word)

    if ok {
        return result
    }

    return nil
}

// ListWords : List custom words
func (speechToText *SpeechToTextV1) ListWords(options *ListWordsOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}/words"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", options.CustomizationID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    if options.IsWordTypeSet {
        request.Query("word_type=" + fmt.Sprint(options.WordType))
    }
    if options.IsSortSet {
        request.Query("sort=" + fmt.Sprint(options.Sort))
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

    response.Result = new(Words)
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

// GetListWordsResult : Cast result of ListWords operation
func GetListWordsResult(response *watson.WatsonResponse) *Words {
    result, ok := response.Result.(*Words)

    if ok {
        return result
    }

    return nil
}

// CreateAcousticModel : Create a custom acoustic model
func (speechToText *SpeechToTextV1) CreateAcousticModel(options *CreateAcousticModelOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/acoustic_customizations"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    body := map[string]interface{}{}
    body["name"] = options.Name
    body["base_model_name"] = options.BaseModelName
    if options.IsDescriptionSet {
        body["description"] = options.Description
    }
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

    response.Result = new(AcousticModel)
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

// GetCreateAcousticModelResult : Cast result of CreateAcousticModel operation
func GetCreateAcousticModelResult(response *watson.WatsonResponse) *AcousticModel {
    result, ok := response.Result.(*AcousticModel)

    if ok {
        return result
    }

    return nil
}

// DeleteAcousticModel : Delete a custom acoustic model
func (speechToText *SpeechToTextV1) DeleteAcousticModel(options *DeleteAcousticModelOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/acoustic_customizations/{customization_id}"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", options.CustomizationID, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")

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

    res, _, err := request.End()

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


// GetAcousticModel : Get a custom acoustic model
func (speechToText *SpeechToTextV1) GetAcousticModel(options *GetAcousticModelOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/acoustic_customizations/{customization_id}"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", options.CustomizationID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")

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

    response.Result = new(AcousticModel)
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

// GetGetAcousticModelResult : Cast result of GetAcousticModel operation
func GetGetAcousticModelResult(response *watson.WatsonResponse) *AcousticModel {
    result, ok := response.Result.(*AcousticModel)

    if ok {
        return result
    }

    return nil
}

// ListAcousticModels : List custom acoustic models
func (speechToText *SpeechToTextV1) ListAcousticModels(options *ListAcousticModelsOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/acoustic_customizations"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    if options.IsLanguageSet {
        request.Query("language=" + fmt.Sprint(options.Language))
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

    response.Result = new(AcousticModels)
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

// GetListAcousticModelsResult : Cast result of ListAcousticModels operation
func GetListAcousticModelsResult(response *watson.WatsonResponse) *AcousticModels {
    result, ok := response.Result.(*AcousticModels)

    if ok {
        return result
    }

    return nil
}

// ResetAcousticModel : Reset a custom acoustic model
func (speechToText *SpeechToTextV1) ResetAcousticModel(options *ResetAcousticModelOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/acoustic_customizations/{customization_id}/reset"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", options.CustomizationID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")

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

    res, _, err := request.End()

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


// TrainAcousticModel : Train a custom acoustic model
func (speechToText *SpeechToTextV1) TrainAcousticModel(options *TrainAcousticModelOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/acoustic_customizations/{customization_id}/train"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", options.CustomizationID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    if options.IsCustomLanguageModelIDSet {
        request.Query("custom_language_model_id=" + fmt.Sprint(options.CustomLanguageModelID))
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

    res, _, err := request.End()

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


// UpgradeAcousticModel : Upgrade a custom acoustic model
func (speechToText *SpeechToTextV1) UpgradeAcousticModel(options *UpgradeAcousticModelOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/acoustic_customizations/{customization_id}/upgrade_model"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", options.CustomizationID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    if options.IsCustomLanguageModelIDSet {
        request.Query("custom_language_model_id=" + fmt.Sprint(options.CustomLanguageModelID))
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

    res, _, err := request.End()

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


// AddAudio : Add an audio resource
func (speechToText *SpeechToTextV1) AddAudio(options *AddAudioOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/acoustic_customizations/{customization_id}/audio/{audio_name}"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", options.CustomizationID, 1)
    path = strings.Replace(path, "{audio_name}", options.AudioName, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", fmt.Sprint(options.ContentType))
    if options.IsContainedContentTypeSet {
        request.Set("Contained-Content-Type", fmt.Sprint(options.ContainedContentType))
    }
    if options.IsAllowOverwriteSet {
        request.Query("allow_overwrite=" + fmt.Sprint(options.AllowOverwrite))
    }
    audioResource := new(bytes.Buffer)
    audioResource.ReadFrom(options.AudioResource)
    request.Send(audioResource.String())

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

    res, _, err := request.End()

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


// DeleteAudio : Delete an audio resource
func (speechToText *SpeechToTextV1) DeleteAudio(options *DeleteAudioOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/acoustic_customizations/{customization_id}/audio/{audio_name}"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", options.CustomizationID, 1)
    path = strings.Replace(path, "{audio_name}", options.AudioName, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")

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

    res, _, err := request.End()

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


// GetAudio : Get an audio resource
func (speechToText *SpeechToTextV1) GetAudio(options *GetAudioOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/acoustic_customizations/{customization_id}/audio/{audio_name}"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", options.CustomizationID, 1)
    path = strings.Replace(path, "{audio_name}", options.AudioName, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")

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

    response.Result = new(AudioListing)
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

// GetGetAudioResult : Cast result of GetAudio operation
func GetGetAudioResult(response *watson.WatsonResponse) *AudioListing {
    result, ok := response.Result.(*AudioListing)

    if ok {
        return result
    }

    return nil
}

// ListAudio : List audio resources
func (speechToText *SpeechToTextV1) ListAudio(options *ListAudioOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/acoustic_customizations/{customization_id}/audio"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", options.CustomizationID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")

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

    response.Result = new(AudioResources)
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

// GetListAudioResult : Cast result of ListAudio operation
func GetListAudioResult(response *watson.WatsonResponse) *AudioResources {
    result, ok := response.Result.(*AudioResources)

    if ok {
        return result
    }

    return nil
}

// DeleteUserData : Delete labeled data
func (speechToText *SpeechToTextV1) DeleteUserData(options *DeleteUserDataOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/user_data"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    request := req.New().Delete(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("customer_id=" + fmt.Sprint(options.CustomerID))

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

    res, _, err := request.End()

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



// AcousticModel : AcousticModel struct
type AcousticModel struct {

	// The customization ID (GUID) of the custom acoustic model. The **Create a custom acoustic model** method returns only this field of the object; it does not return the other fields.
	CustomizationID string `json:"customization_id"`

	// The date and time in Coordinated Universal Time (UTC) at which the custom acoustic model was created. The value is provided in full ISO 8601 format (`YYYY-MM-DDThh:mm:ss.sTZD`).
	Created string `json:"created,omitempty"`

	// The language identifier of the custom acoustic model (for example, `en-US`).
	Language string `json:"language,omitempty"`

	// A list of the available versions of the custom acoustic model. Each element of the array indicates a version of the base model with which the custom model can be used. Multiple versions exist only if the custom model has been upgraded; otherwise, only a single version is shown.
	Versions []string `json:"versions,omitempty"`

	// The GUID of the service credentials for the instance of the service that owns the custom acoustic model.
	Owner string `json:"owner,omitempty"`

	// The name of the custom acoustic model.
	Name string `json:"name,omitempty"`

	// The description of the custom acoustic model.
	Description string `json:"description,omitempty"`

	// The name of the language model for which the custom acoustic model was created.
	BaseModelName string `json:"base_model_name,omitempty"`

	// The current status of the custom acoustic model: * `pending` indicates that the model was created but is waiting either for training data to be added or for the service to finish analyzing added data. * `ready` indicates that the model contains data and is ready to be trained. * `training` indicates that the model is currently being trained. * `available` indicates that the model is trained and ready to use. * `upgrading` indicates that the model is currently being upgraded. * `failed` indicates that training of the model failed.
	Status string `json:"status,omitempty"`

	// A percentage that indicates the progress of the custom acoustic model's current training. A value of `100` means that the model is fully trained. **Note:** The `progress` field does not currently reflect the progress of the training. The field changes from `0` to `100` when training is complete.
	Progress int64 `json:"progress,omitempty"`

	// If the request included unknown parameters, the following message: `Unexpected query parameter(s) ['parameters'] detected`, where `parameters` is a list that includes a quoted string for each unknown parameter.
	Warnings string `json:"warnings,omitempty"`
}

// AcousticModels : AcousticModels struct
type AcousticModels struct {

	// An array of objects that provides information about each available custom acoustic model. The array is empty if the requesting service credentials own no custom acoustic models (if no language is specified) or own no custom acoustic models for the specified language.
	Customizations []AcousticModel `json:"customizations"`
}

// AddAudioOptions : The addAudio options.
type AddAudioOptions struct {

	// The customization ID (GUID) of the custom acoustic model. You must make the request with service credentials created for the instance of the service that owns the custom model.
	CustomizationID string `json:"customization_id"`

	// The name of the audio resource for the custom acoustic model. When adding an audio resource, do not include spaces in the name; use a localized name that matches the language of the custom model.
	AudioName string `json:"audio_name"`

	AudioResource io.ReadCloser `json:"audio_resource"`

	// The type of the input.
	ContentType string `json:"Content-Type"`

	// For an archive-type resource, specifies the format of the audio files contained in the archive file. The parameter accepts all of the audio formats supported for use with speech recognition, including the `rate`, `channels`, and `endianness` parameters that are used with some formats. For a complete list of supported audio formats, see [Audio formats](/docs/services/speech-to-text/input.html#formats).
	ContainedContentType string `json:"Contained-Content-Type,omitempty"`

    // Indicates whether user set optional parameter ContainedContentType
    IsContainedContentTypeSet bool

	// If `true`, the specified corpus or audio resource overwrites an existing corpus or audio resource with the same name. If `false`, the request fails if a corpus or audio resource with the same name already exists. The parameter has no effect if a corpus or audio resource with the same name does not already exist.
	AllowOverwrite bool `json:"allow_overwrite,omitempty"`

    // Indicates whether user set optional parameter AllowOverwrite
    IsAllowOverwriteSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewAddAudioOptionsForZip : Instantiate AddAudioOptionsForZip
func NewAddAudioOptionsForZip(audioResource io.ReadCloser) *AddAudioOptions {
    return &AddAudioOptions{
        AudioResource: audioResource,
        ContentType: "application/zip",
    }
}

// NewAddAudioOptionsForGzip : Instantiate AddAudioOptionsForGzip
func NewAddAudioOptionsForGzip(audioResource io.ReadCloser) *AddAudioOptions {
    return &AddAudioOptions{
        AudioResource: audioResource,
        ContentType: "application/gzip",
    }
}

// NewAddAudioOptionsForBasic : Instantiate AddAudioOptionsForBasic
func NewAddAudioOptionsForBasic(audioResource io.ReadCloser) *AddAudioOptions {
    return &AddAudioOptions{
        AudioResource: audioResource,
        ContentType: "audio/basic",
    }
}

// NewAddAudioOptionsForFlac : Instantiate AddAudioOptionsForFlac
func NewAddAudioOptionsForFlac(audioResource io.ReadCloser) *AddAudioOptions {
    return &AddAudioOptions{
        AudioResource: audioResource,
        ContentType: "audio/flac",
    }
}

// NewAddAudioOptionsForL16 : Instantiate AddAudioOptionsForL16
func NewAddAudioOptionsForL16(audioResource io.ReadCloser) *AddAudioOptions {
    return &AddAudioOptions{
        AudioResource: audioResource,
        ContentType: "audio/l16",
    }
}

// NewAddAudioOptionsForMp3 : Instantiate AddAudioOptionsForMp3
func NewAddAudioOptionsForMp3(audioResource io.ReadCloser) *AddAudioOptions {
    return &AddAudioOptions{
        AudioResource: audioResource,
        ContentType: "audio/mp3",
    }
}

// NewAddAudioOptionsForMpeg : Instantiate AddAudioOptionsForMpeg
func NewAddAudioOptionsForMpeg(audioResource io.ReadCloser) *AddAudioOptions {
    return &AddAudioOptions{
        AudioResource: audioResource,
        ContentType: "audio/mpeg",
    }
}

// NewAddAudioOptionsForMulaw : Instantiate AddAudioOptionsForMulaw
func NewAddAudioOptionsForMulaw(audioResource io.ReadCloser) *AddAudioOptions {
    return &AddAudioOptions{
        AudioResource: audioResource,
        ContentType: "audio/mulaw",
    }
}

// NewAddAudioOptionsForOgg : Instantiate AddAudioOptionsForOgg
func NewAddAudioOptionsForOgg(audioResource io.ReadCloser) *AddAudioOptions {
    return &AddAudioOptions{
        AudioResource: audioResource,
        ContentType: "audio/ogg",
    }
}

// NewAddAudioOptionsForOggcodecsopus : Instantiate AddAudioOptionsForOggcodecsopus
func NewAddAudioOptionsForOggcodecsopus(audioResource io.ReadCloser) *AddAudioOptions {
    return &AddAudioOptions{
        AudioResource: audioResource,
        ContentType: "audio/ogg;codecs=opus",
    }
}

// NewAddAudioOptionsForOggcodecsvorbis : Instantiate AddAudioOptionsForOggcodecsvorbis
func NewAddAudioOptionsForOggcodecsvorbis(audioResource io.ReadCloser) *AddAudioOptions {
    return &AddAudioOptions{
        AudioResource: audioResource,
        ContentType: "audio/ogg;codecs=vorbis",
    }
}

// NewAddAudioOptionsForWav : Instantiate AddAudioOptionsForWav
func NewAddAudioOptionsForWav(audioResource io.ReadCloser) *AddAudioOptions {
    return &AddAudioOptions{
        AudioResource: audioResource,
        ContentType: "audio/wav",
    }
}

// NewAddAudioOptionsForWebm : Instantiate AddAudioOptionsForWebm
func NewAddAudioOptionsForWebm(audioResource io.ReadCloser) *AddAudioOptions {
    return &AddAudioOptions{
        AudioResource: audioResource,
        ContentType: "audio/webm",
    }
}

// NewAddAudioOptionsForWebmcodecsopus : Instantiate AddAudioOptionsForWebmcodecsopus
func NewAddAudioOptionsForWebmcodecsopus(audioResource io.ReadCloser) *AddAudioOptions {
    return &AddAudioOptions{
        AudioResource: audioResource,
        ContentType: "audio/webm;codecs=opus",
    }
}

// NewAddAudioOptionsForWebmcodecsvorbis : Instantiate AddAudioOptionsForWebmcodecsvorbis
func NewAddAudioOptionsForWebmcodecsvorbis(audioResource io.ReadCloser) *AddAudioOptions {
    return &AddAudioOptions{
        AudioResource: audioResource,
        ContentType: "audio/webm;codecs=vorbis",
    }
}

// SetAudioResource : Allow user to set AudioResource with specified ContentType
func (options *AddAudioOptions) SetAudioResource(audioResource io.ReadCloser, contentType string) *AddAudioOptions {
    options.AudioResource = audioResource
    options.ContentType = contentType
    return options
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *AddAudioOptions) SetCustomizationID(param string) *AddAudioOptions {
    options.CustomizationID = param
    return options
}

// SetAudioName : Allow user to set AudioName
func (options *AddAudioOptions) SetAudioName(param string) *AddAudioOptions {
    options.AudioName = param
    return options
}

// SetContainedContentType : Allow user to set ContainedContentType
func (options *AddAudioOptions) SetContainedContentType(param string) *AddAudioOptions {
    options.ContainedContentType = param
    options.IsContainedContentTypeSet = true
    return options
}

// SetAllowOverwrite : Allow user to set AllowOverwrite
func (options *AddAudioOptions) SetAllowOverwrite(param bool) *AddAudioOptions {
    options.AllowOverwrite = param
    options.IsAllowOverwriteSet = true
    return options
}

// SetHeaders : Allow user to set Headers
func (options *AddAudioOptions) SetHeaders(param map[string]string) *AddAudioOptions {
    options.Headers = param
    return options
}

// AddCorpusOptions : The addCorpus options.
type AddCorpusOptions struct {

	// The customization ID (GUID) of the custom language model. You must make the request with service credentials created for the instance of the service that owns the custom model.
	CustomizationID string `json:"customization_id"`

	// The name of the corpus for the custom language model. When adding a corpus, do not include spaces in the name; use a localized name that matches the language of the custom model; and do not use the name `user`, which is reserved by the service to denote custom words added or modified by the user.
	CorpusName string `json:"corpus_name"`

	// A plain text file that contains the training data for the corpus. Encode the file in UTF-8 if it contains non-ASCII characters; the service assumes UTF-8 encoding if it encounters non-ASCII characters. With cURL, use the `--data-binary` option to upload the file for the request.
	CorpusFile os.File `json:"corpus_file"`

	// If `true`, the specified corpus or audio resource overwrites an existing corpus or audio resource with the same name. If `false`, the request fails if a corpus or audio resource with the same name already exists. The parameter has no effect if a corpus or audio resource with the same name does not already exist.
	AllowOverwrite bool `json:"allow_overwrite,omitempty"`

    // Indicates whether user set optional parameter AllowOverwrite
    IsAllowOverwriteSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewAddCorpusOptions : Instantiate AddCorpusOptions
func NewAddCorpusOptions(customizationID string, corpusName string, corpusFile os.File) *AddCorpusOptions {
    return &AddCorpusOptions{
        CustomizationID: customizationID,
        CorpusName: corpusName,
        CorpusFile: corpusFile,
    }
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *AddCorpusOptions) SetCustomizationID(param string) *AddCorpusOptions {
    options.CustomizationID = param
    return options
}

// SetCorpusName : Allow user to set CorpusName
func (options *AddCorpusOptions) SetCorpusName(param string) *AddCorpusOptions {
    options.CorpusName = param
    return options
}

// SetCorpusFile : Allow user to set CorpusFile
func (options *AddCorpusOptions) SetCorpusFile(param os.File) *AddCorpusOptions {
    options.CorpusFile = param
    return options
}

// SetAllowOverwrite : Allow user to set AllowOverwrite
func (options *AddCorpusOptions) SetAllowOverwrite(param bool) *AddCorpusOptions {
    options.AllowOverwrite = param
    options.IsAllowOverwriteSet = true
    return options
}

// SetHeaders : Allow user to set Headers
func (options *AddCorpusOptions) SetHeaders(param map[string]string) *AddCorpusOptions {
    options.Headers = param
    return options
}

// AddWordOptions : The addWord options.
type AddWordOptions struct {

	// The customization ID (GUID) of the custom language model. You must make the request with service credentials created for the instance of the service that owns the custom model.
	CustomizationID string `json:"customization_id"`

	// The custom word for the custom language model. When you add or update a custom word with the **Add a custom word** method, do not include spaces in the word. Use a `-` (dash) or `_` (underscore) to connect the tokens of compound words.
	WordName string `json:"word_name"`

	// For the **Add custom words** method, you must specify the custom word that is to be added to or updated in the custom model. Do not include spaces in the word. Use a `-` (dash) or `_` (underscore) to connect the tokens of compound words. Omit this field for the **Add a custom word** method.
	Word string `json:"word,omitempty"`

    // Indicates whether user set optional parameter Word
    IsWordSet bool

	// An array of sounds-like pronunciations for the custom word. Specify how words that are difficult to pronounce, foreign words, acronyms, and so on can be pronounced by users. For a word that is not in the service's base vocabulary, omit the parameter to have the service automatically generate a sounds-like pronunciation for the word. For a word that is in the service's base vocabulary, use the parameter to specify additional pronunciations for the word. You cannot override the default pronunciation of a word; pronunciations you add augment the pronunciation from the base vocabulary. A word can have at most five sounds-like pronunciations, and a pronunciation can include at most 40 characters not including spaces.
	SoundsLike []string `json:"sounds_like,omitempty"`

    // Indicates whether user set optional parameter SoundsLike
    IsSoundsLikeSet bool

	// An alternative spelling for the custom word when it appears in a transcript. Use the parameter when you want the word to have a spelling that is different from its usual representation or from its spelling in corpora training data.
	DisplayAs string `json:"display_as,omitempty"`

    // Indicates whether user set optional parameter DisplayAs
    IsDisplayAsSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewAddWordOptions : Instantiate AddWordOptions
func NewAddWordOptions(customizationID string, wordName string) *AddWordOptions {
    return &AddWordOptions{
        CustomizationID: customizationID,
        WordName: wordName,
    }
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *AddWordOptions) SetCustomizationID(param string) *AddWordOptions {
    options.CustomizationID = param
    return options
}

// SetWordName : Allow user to set WordName
func (options *AddWordOptions) SetWordName(param string) *AddWordOptions {
    options.WordName = param
    return options
}

// SetWord : Allow user to set Word
func (options *AddWordOptions) SetWord(param string) *AddWordOptions {
    options.Word = param
    options.IsWordSet = true
    return options
}

// SetSoundsLike : Allow user to set SoundsLike
func (options *AddWordOptions) SetSoundsLike(param []string) *AddWordOptions {
    options.SoundsLike = param
    options.IsSoundsLikeSet = true
    return options
}

// SetDisplayAs : Allow user to set DisplayAs
func (options *AddWordOptions) SetDisplayAs(param string) *AddWordOptions {
    options.DisplayAs = param
    options.IsDisplayAsSet = true
    return options
}

// SetHeaders : Allow user to set Headers
func (options *AddWordOptions) SetHeaders(param map[string]string) *AddWordOptions {
    options.Headers = param
    return options
}

// AddWordsOptions : The addWords options.
type AddWordsOptions struct {

	// The customization ID (GUID) of the custom language model. You must make the request with service credentials created for the instance of the service that owns the custom model.
	CustomizationID string `json:"customization_id"`

	// An array of objects that provides information about each custom word that is to be added to or updated in the custom language model.
	Words []CustomWord `json:"words"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewAddWordsOptions : Instantiate AddWordsOptions
func NewAddWordsOptions(customizationID string, words []CustomWord) *AddWordsOptions {
    return &AddWordsOptions{
        CustomizationID: customizationID,
        Words: words,
    }
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *AddWordsOptions) SetCustomizationID(param string) *AddWordsOptions {
    options.CustomizationID = param
    return options
}

// SetWords : Allow user to set Words
func (options *AddWordsOptions) SetWords(param []CustomWord) *AddWordsOptions {
    options.Words = param
    return options
}

// SetHeaders : Allow user to set Headers
func (options *AddWordsOptions) SetHeaders(param map[string]string) *AddWordsOptions {
    options.Headers = param
    return options
}

// AudioDetails : AudioDetails struct
type AudioDetails struct {

	// The type of the audio resource: * `audio` for an individual audio file * `archive` for an archive (**.zip** or **.tar.gz**) file that contains audio files * `undetermined` for a resource that the service cannot validate (for example, if the user mistakenly passes a file that does not contain audio, such as a JPEG file).
	TypeVar string `json:"type,omitempty"`

	// **For an audio-type resource,** the codec in which the audio is encoded. Omitted for an archive-type resource.
	Codec string `json:"codec,omitempty"`

	// **For an audio-type resource,** the sampling rate of the audio in Hertz (samples per second). Omitted for an archive-type resource.
	Frequency int64 `json:"frequency,omitempty"`

	// **For an archive-type resource,** the format of the compressed archive: * `zip` for a **.zip** file * `gzip` for a **.tar.gz** file Omitted for an audio-type resource.
	Compression string `json:"compression,omitempty"`
}

// AudioListing : AudioListing struct
type AudioListing struct {

	// **For an audio-type resource,**  the total seconds of audio in the resource. The value is always a whole number. Omitted for an archive-type resource.
	Duration float64 `json:"duration,omitempty"`

	// **For an audio-type resource,** the user-specified name of the resource. Omitted for an archive-type resource.
	Name string `json:"name,omitempty"`

	// **For an audio-type resource,** an `AudioDetails` object that provides detailed information about the resource. The object is empty until the service finishes processing the audio. Omitted for an archive-type resource.
	Details AudioDetails `json:"details,omitempty"`

	// **For an audio-type resource,** the status of the resource: * `ok` indicates that the service has successfully analyzed the audio data. The data can be used to train the custom model. * `being_processed` indicates that the service is still analyzing the audio data. The service cannot accept requests to add new audio resources or to train the custom model until its analysis is complete. * `invalid` indicates that the audio data is not valid for training the custom model (possibly because it has the wrong format or sampling rate, or because it is corrupted). Omitted for an archive-type resource.
	Status string `json:"status,omitempty"`

	// **For an archive-type resource,** an object of type `AudioResource` that provides information about the resource. Omitted for an audio-type resource.
	Container AudioResource `json:"container,omitempty"`

	// **For an archive-type resource,** an array of `AudioResource` objects that provides information about the audio-type resources that are contained in the resource. Omitted for an audio-type resource.
	Audio []AudioResource `json:"audio,omitempty"`
}

// AudioResource : AudioResource struct
type AudioResource struct {

	// The total seconds of audio in the audio resource. The value is always a whole number.
	Duration float64 `json:"duration"`

	// **For an archive-type resource,** the user-specified name of the resource. **For an audio-type resource,** the user-specified name of the resource or the name of the audio file that the user added for the resource. The value depends on the method that is called.
	Name string `json:"name"`

	// An `AudioDetails` object that provides detailed information about the audio resource. The object is empty until the service finishes processing the audio.
	Details AudioDetails `json:"details"`

	// The status of the audio resource: * `ok` indicates that the service has successfully analyzed the audio data. The data can be used to train the custom model. * `being_processed` indicates that the service is still analyzing the audio data. The service cannot accept requests to add new audio resources or to train the custom model until its analysis is complete. * `invalid` indicates that the audio data is not valid for training the custom model (possibly because it has the wrong format or sampling rate, or because it is corrupted). For an archive file, the entire archive is invalid if any of its audio files are invalid.
	Status string `json:"status"`
}

// AudioResources : AudioResources struct
type AudioResources struct {

	// The total minutes of accumulated audio summed over all of the valid audio resources for the custom acoustic model. You can use this value to determine whether the custom model has too little or too much audio to begin training.
	TotalMinutesOfAudio float64 `json:"total_minutes_of_audio"`

	// An array of objects that provides information about the audio resources of the custom acoustic model. The array is empty if the custom model has no audio resources.
	Audio []AudioResource `json:"audio"`
}

// CheckJobOptions : The checkJob options.
type CheckJobOptions struct {

	// The ID of the asynchronous job.
	ID string `json:"id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewCheckJobOptions : Instantiate CheckJobOptions
func NewCheckJobOptions(id string) *CheckJobOptions {
    return &CheckJobOptions{
        ID: id,
    }
}

// SetID : Allow user to set ID
func (options *CheckJobOptions) SetID(param string) *CheckJobOptions {
    options.ID = param
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
func NewCheckJobsOptions() *CheckJobsOptions {
    return &CheckJobsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *CheckJobsOptions) SetHeaders(param map[string]string) *CheckJobsOptions {
    options.Headers = param
    return options
}

// Corpora : Corpora struct
type Corpora struct {

	// An array of objects that provides information about the corpora for the custom model. The array is empty if the custom model has no corpora.
	Corpora []Corpus `json:"corpora"`
}

// Corpus : Corpus struct
type Corpus struct {

	// The name of the corpus.
	Name string `json:"name"`

	// The total number of words in the corpus. The value is `0` while the corpus is being processed.
	TotalWords int64 `json:"total_words"`

	// The number of OOV words in the corpus. The value is `0` while the corpus is being processed.
	OutOfVocabularyWords int64 `json:"out_of_vocabulary_words"`

	// The status of the corpus: * `analyzed` indicates that the service has successfully analyzed the corpus; the custom model can be trained with data from the corpus. * `being_processed` indicates that the service is still analyzing the corpus; the service cannot accept requests to add new corpora or words, or to train the custom model. * `undetermined` indicates that the service encountered an error while processing the corpus.
	Status string `json:"status"`

	// If the status of the corpus is `undetermined`, the following message: `Analysis of corpus 'name' failed. Please try adding the corpus again by setting the 'allow_overwrite' flag to 'true'`.
	Error string `json:"error,omitempty"`
}

// CreateAcousticModelOptions : The createAcousticModel options.
type CreateAcousticModelOptions struct {

	// A user-defined name for the new custom acoustic model. Use a name that is unique among all custom acoustic models that you own. Use a localized name that matches the language of the custom model. Use a name that describes the acoustic environment of the custom model, such as `Mobile custom model` or `Noisy car custom model`.
	Name string `json:"name"`

	// The name of the base language model that is to be customized by the new custom acoustic model. The new custom model can be used only with the base model that it customizes. To determine whether a base model supports acoustic model customization, refer to [Language support for customization](https://console.bluemix.net/docs/services/speech-to-text/custom.html#languageSupport).
	BaseModelName string `json:"base_model_name"`

	// A description of the new custom acoustic model. Use a localized description that matches the language of the custom model.
	Description string `json:"description,omitempty"`

    // Indicates whether user set optional parameter Description
    IsDescriptionSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewCreateAcousticModelOptions : Instantiate CreateAcousticModelOptions
func NewCreateAcousticModelOptions(name string, baseModelName string) *CreateAcousticModelOptions {
    return &CreateAcousticModelOptions{
        Name: name,
        BaseModelName: baseModelName,
    }
}

// SetName : Allow user to set Name
func (options *CreateAcousticModelOptions) SetName(param string) *CreateAcousticModelOptions {
    options.Name = param
    return options
}

// SetBaseModelName : Allow user to set BaseModelName
func (options *CreateAcousticModelOptions) SetBaseModelName(param string) *CreateAcousticModelOptions {
    options.BaseModelName = param
    return options
}

// SetDescription : Allow user to set Description
func (options *CreateAcousticModelOptions) SetDescription(param string) *CreateAcousticModelOptions {
    options.Description = param
    options.IsDescriptionSet = true
    return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateAcousticModelOptions) SetHeaders(param map[string]string) *CreateAcousticModelOptions {
    options.Headers = param
    return options
}

// CreateJobOptions : The createJob options.
type CreateJobOptions struct {

	Audio io.ReadCloser `json:"audio,omitempty"`

	// The type of the input.
	ContentType string `json:"Content-Type"`

	// The identifier of the model that is to be used for the recognition request or, for the **Create a session** method, with the new session.
	Model string `json:"model,omitempty"`

    // Indicates whether user set optional parameter Model
    IsModelSet bool

	// A URL to which callback notifications are to be sent. The URL must already be successfully white-listed by using the **Register a callback** method. You can include the same callback URL with any number of job creation requests. Omit the parameter to poll the service for job completion and results. Use the `user_token` parameter to specify a unique user-specified string with each job to differentiate the callback notifications for the jobs.
	CallbackURL string `json:"callback_url,omitempty"`

    // Indicates whether user set optional parameter CallbackURL
    IsCallbackURLSet bool

	// If the job includes a callback URL, a comma-separated list of notification events to which to subscribe. Valid events are * `recognitions.started` generates a callback notification when the service begins to process the job. * `recognitions.completed` generates a callback notification when the job is complete. You must use the **Check a job** method to retrieve the results before they time out or are deleted. * `recognitions.completed_with_results` generates a callback notification when the job is complete. The notification includes the results of the request. * `recognitions.failed` generates a callback notification if the service experiences an error while processing the job. Omit the parameter to subscribe to the default events: `recognitions.started`, `recognitions.completed`, and `recognitions.failed`. The `recognitions.completed` and `recognitions.completed_with_results` events are incompatible; you can specify only of the two events. If the job does not include a callback URL, omit the parameter.
	Events string `json:"events,omitempty"`

    // Indicates whether user set optional parameter Events
    IsEventsSet bool

	// If the job includes a callback URL, a user-specified string that the service is to include with each callback notification for the job; the token allows the user to maintain an internal mapping between jobs and notification events. If the job does not include a callback URL, omit the parameter.
	UserToken string `json:"user_token,omitempty"`

    // Indicates whether user set optional parameter UserToken
    IsUserTokenSet bool

	// The number of minutes for which the results are to be available after the job has finished. If not delivered via a callback, the results must be retrieved within this time. Omit the parameter to use a time to live of one week. The parameter is valid with or without a callback URL.
	ResultsTTL int64 `json:"results_ttl,omitempty"`

    // Indicates whether user set optional parameter ResultsTTL
    IsResultsTTLSet bool

	// The customization ID (GUID) of a custom language model that is to be used with the recognition request or, for the **Create a session** method, with the new session. The base model of the specified custom language model must match the model specified with the `model` parameter. You must make the request with service credentials created for the instance of the service that owns the custom model. By default, no custom language model is used.
	CustomizationID string `json:"customization_id,omitempty"`

    // Indicates whether user set optional parameter CustomizationID
    IsCustomizationIDSet bool

	// The customization ID (GUID) of a custom acoustic model that is to be used with the recognition request or, for the **Create a session** method, with the new session. The base model of the specified custom acoustic model must match the model specified with the `model` parameter. You must make the request with service credentials created for the instance of the service that owns the custom model. By default, no custom acoustic model is used.
	AcousticCustomizationID string `json:"acoustic_customization_id,omitempty"`

    // Indicates whether user set optional parameter AcousticCustomizationID
    IsAcousticCustomizationIDSet bool

	// The version of the specified base model that is to be used with recognition request or, for the **Create a session** method, with the new session. Multiple versions of a base model can exist when a model is updated for internal improvements. The parameter is intended primarily for use with custom models that have been upgraded for a new base model. The default value depends on whether the parameter is used with or without a custom model. For more information, see [Base model version](https://console.bluemix.net/docs/services/speech-to-text/input.html#version).
	BaseModelVersion string `json:"base_model_version,omitempty"`

    // Indicates whether user set optional parameter BaseModelVersion
    IsBaseModelVersionSet bool

	// If you specify the customization ID (GUID) of a custom language model with the recognition request or, for sessions, with the **Create a session** method, the customization weight tells the service how much weight to give to words from the custom language model compared to those from the base model for the current request. Specify a value between 0.0 and 1.0. Unless a different customization weight was specified for the custom model when it was trained, the default value is 0.3. A customization weight that you specify overrides a weight that was specified when the custom model was trained. The default value yields the best performance in general. Assign a higher value if your audio makes frequent use of OOV words from the custom model. Use caution when setting the weight: a higher value can improve the accuracy of phrases from the custom model's domain, but it can negatively affect performance on non-domain phrases.
	CustomizationWeight float64 `json:"customization_weight,omitempty"`

    // Indicates whether user set optional parameter CustomizationWeight
    IsCustomizationWeightSet bool

	// The time in seconds after which, if only silence (no speech) is detected in submitted audio, the connection is closed with a 400 error. The parameter is useful for stopping audio submission from a live microphone when a user simply walks away. Use `-1` for infinity.
	InactivityTimeout int64 `json:"inactivity_timeout,omitempty"`

    // Indicates whether user set optional parameter InactivityTimeout
    IsInactivityTimeoutSet bool

	// An array of keyword strings to spot in the audio. Each keyword string can include one or more string tokens. Keywords are spotted only in the final results, not in interim hypotheses. If you specify any keywords, you must also specify a keywords threshold. You can spot a maximum of 1000 keywords. Omit the parameter or specify an empty array if you do not need to spot keywords.
	Keywords []string `json:"keywords,omitempty"`

    // Indicates whether user set optional parameter Keywords
    IsKeywordsSet bool

	// A confidence value that is the lower bound for spotting a keyword. A word is considered to match a keyword if its confidence is greater than or equal to the threshold. Specify a probability between 0.0 and 1.0. No keyword spotting is performed if you omit the parameter. If you specify a threshold, you must also specify one or more keywords.
	KeywordsThreshold float32 `json:"keywords_threshold,omitempty"`

    // Indicates whether user set optional parameter KeywordsThreshold
    IsKeywordsThresholdSet bool

	// The maximum number of alternative transcripts that the service is to return. By default, a single transcription is returned.
	MaxAlternatives int64 `json:"max_alternatives,omitempty"`

    // Indicates whether user set optional parameter MaxAlternatives
    IsMaxAlternativesSet bool

	// A confidence value that is the lower bound for identifying a hypothesis as a possible word alternative (also known as "Confusion Networks"). An alternative word is considered if its confidence is greater than or equal to the threshold. Specify a probability between 0.0 and 1.0. No alternative words are computed if you omit the parameter.
	WordAlternativesThreshold float32 `json:"word_alternatives_threshold,omitempty"`

    // Indicates whether user set optional parameter WordAlternativesThreshold
    IsWordAlternativesThresholdSet bool

	// If `true`, the service returns a confidence measure in the range of 0.0 to 1.0 for each word. By default, no word confidence measures are returned.
	WordConfidence bool `json:"word_confidence,omitempty"`

    // Indicates whether user set optional parameter WordConfidence
    IsWordConfidenceSet bool

	// If `true`, the service returns time alignment for each word. By default, no timestamps are returned.
	Timestamps bool `json:"timestamps,omitempty"`

    // Indicates whether user set optional parameter Timestamps
    IsTimestampsSet bool

	// If `true`, the service filters profanity from all output except for keyword results by replacing inappropriate words with a series of asterisks. Set the parameter to `false` to return results with no censoring. Applies to US English transcription only.
	ProfanityFilter bool `json:"profanity_filter,omitempty"`

    // Indicates whether user set optional parameter ProfanityFilter
    IsProfanityFilterSet bool

	// If `true`, the service converts dates, times, series of digits and numbers, phone numbers, currency values, and internet addresses into more readable, conventional representations in the final transcript of a recognition request. For US English, the service also converts certain keyword strings to punctuation symbols. By default, no smart formatting is performed. Applies to US English and Spanish transcription only.
	SmartFormatting bool `json:"smart_formatting,omitempty"`

    // Indicates whether user set optional parameter SmartFormatting
    IsSmartFormattingSet bool

	// If `true`, the response includes labels that identify which words were spoken by which participants in a multi-person exchange. By default, no speaker labels are returned. Setting `speaker_labels` to `true` forces the `timestamps` parameter to be `true`, regardless of whether you specify `false` for the parameter. To determine whether a language model supports speaker labels, use the **Get models** method and check that the attribute `speaker_labels` is set to `true`. You can also refer to [Speaker labels](https://console.bluemix.net/docs/services/speech-to-text/output.html#speaker_labels).
	SpeakerLabels bool `json:"speaker_labels,omitempty"`

    // Indicates whether user set optional parameter SpeakerLabels
    IsSpeakerLabelsSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewCreateJobOptionsForBasic : Instantiate CreateJobOptionsForBasic
func NewCreateJobOptionsForBasic(audio io.ReadCloser) *CreateJobOptions {
    return &CreateJobOptions{
        Audio: audio,
        ContentType: "audio/basic",
    }
}

// NewCreateJobOptionsForFlac : Instantiate CreateJobOptionsForFlac
func NewCreateJobOptionsForFlac(audio io.ReadCloser) *CreateJobOptions {
    return &CreateJobOptions{
        Audio: audio,
        ContentType: "audio/flac",
    }
}

// NewCreateJobOptionsForL16 : Instantiate CreateJobOptionsForL16
func NewCreateJobOptionsForL16(audio io.ReadCloser) *CreateJobOptions {
    return &CreateJobOptions{
        Audio: audio,
        ContentType: "audio/l16",
    }
}

// NewCreateJobOptionsForMp3 : Instantiate CreateJobOptionsForMp3
func NewCreateJobOptionsForMp3(audio io.ReadCloser) *CreateJobOptions {
    return &CreateJobOptions{
        Audio: audio,
        ContentType: "audio/mp3",
    }
}

// NewCreateJobOptionsForMpeg : Instantiate CreateJobOptionsForMpeg
func NewCreateJobOptionsForMpeg(audio io.ReadCloser) *CreateJobOptions {
    return &CreateJobOptions{
        Audio: audio,
        ContentType: "audio/mpeg",
    }
}

// NewCreateJobOptionsForMulaw : Instantiate CreateJobOptionsForMulaw
func NewCreateJobOptionsForMulaw(audio io.ReadCloser) *CreateJobOptions {
    return &CreateJobOptions{
        Audio: audio,
        ContentType: "audio/mulaw",
    }
}

// NewCreateJobOptionsForOgg : Instantiate CreateJobOptionsForOgg
func NewCreateJobOptionsForOgg(audio io.ReadCloser) *CreateJobOptions {
    return &CreateJobOptions{
        Audio: audio,
        ContentType: "audio/ogg",
    }
}

// NewCreateJobOptionsForOggcodecsopus : Instantiate CreateJobOptionsForOggcodecsopus
func NewCreateJobOptionsForOggcodecsopus(audio io.ReadCloser) *CreateJobOptions {
    return &CreateJobOptions{
        Audio: audio,
        ContentType: "audio/ogg;codecs=opus",
    }
}

// NewCreateJobOptionsForOggcodecsvorbis : Instantiate CreateJobOptionsForOggcodecsvorbis
func NewCreateJobOptionsForOggcodecsvorbis(audio io.ReadCloser) *CreateJobOptions {
    return &CreateJobOptions{
        Audio: audio,
        ContentType: "audio/ogg;codecs=vorbis",
    }
}

// NewCreateJobOptionsForWav : Instantiate CreateJobOptionsForWav
func NewCreateJobOptionsForWav(audio io.ReadCloser) *CreateJobOptions {
    return &CreateJobOptions{
        Audio: audio,
        ContentType: "audio/wav",
    }
}

// NewCreateJobOptionsForWebm : Instantiate CreateJobOptionsForWebm
func NewCreateJobOptionsForWebm(audio io.ReadCloser) *CreateJobOptions {
    return &CreateJobOptions{
        Audio: audio,
        ContentType: "audio/webm",
    }
}

// NewCreateJobOptionsForWebmcodecsopus : Instantiate CreateJobOptionsForWebmcodecsopus
func NewCreateJobOptionsForWebmcodecsopus(audio io.ReadCloser) *CreateJobOptions {
    return &CreateJobOptions{
        Audio: audio,
        ContentType: "audio/webm;codecs=opus",
    }
}

// NewCreateJobOptionsForWebmcodecsvorbis : Instantiate CreateJobOptionsForWebmcodecsvorbis
func NewCreateJobOptionsForWebmcodecsvorbis(audio io.ReadCloser) *CreateJobOptions {
    return &CreateJobOptions{
        Audio: audio,
        ContentType: "audio/webm;codecs=vorbis",
    }
}

// SetAudio : Allow user to set Audio with specified ContentType
func (options *CreateJobOptions) SetAudio(audio io.ReadCloser, contentType string) *CreateJobOptions {
    options.Audio = audio
    options.ContentType = contentType
    return options
}

// SetModel : Allow user to set Model
func (options *CreateJobOptions) SetModel(param string) *CreateJobOptions {
    options.Model = param
    options.IsModelSet = true
    return options
}

// SetCallbackURL : Allow user to set CallbackURL
func (options *CreateJobOptions) SetCallbackURL(param string) *CreateJobOptions {
    options.CallbackURL = param
    options.IsCallbackURLSet = true
    return options
}

// SetEvents : Allow user to set Events
func (options *CreateJobOptions) SetEvents(param string) *CreateJobOptions {
    options.Events = param
    options.IsEventsSet = true
    return options
}

// SetUserToken : Allow user to set UserToken
func (options *CreateJobOptions) SetUserToken(param string) *CreateJobOptions {
    options.UserToken = param
    options.IsUserTokenSet = true
    return options
}

// SetResultsTTL : Allow user to set ResultsTTL
func (options *CreateJobOptions) SetResultsTTL(param int64) *CreateJobOptions {
    options.ResultsTTL = param
    options.IsResultsTTLSet = true
    return options
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *CreateJobOptions) SetCustomizationID(param string) *CreateJobOptions {
    options.CustomizationID = param
    options.IsCustomizationIDSet = true
    return options
}

// SetAcousticCustomizationID : Allow user to set AcousticCustomizationID
func (options *CreateJobOptions) SetAcousticCustomizationID(param string) *CreateJobOptions {
    options.AcousticCustomizationID = param
    options.IsAcousticCustomizationIDSet = true
    return options
}

// SetBaseModelVersion : Allow user to set BaseModelVersion
func (options *CreateJobOptions) SetBaseModelVersion(param string) *CreateJobOptions {
    options.BaseModelVersion = param
    options.IsBaseModelVersionSet = true
    return options
}

// SetCustomizationWeight : Allow user to set CustomizationWeight
func (options *CreateJobOptions) SetCustomizationWeight(param float64) *CreateJobOptions {
    options.CustomizationWeight = param
    options.IsCustomizationWeightSet = true
    return options
}

// SetInactivityTimeout : Allow user to set InactivityTimeout
func (options *CreateJobOptions) SetInactivityTimeout(param int64) *CreateJobOptions {
    options.InactivityTimeout = param
    options.IsInactivityTimeoutSet = true
    return options
}

// SetKeywords : Allow user to set Keywords
func (options *CreateJobOptions) SetKeywords(param []string) *CreateJobOptions {
    options.Keywords = param
    options.IsKeywordsSet = true
    return options
}

// SetKeywordsThreshold : Allow user to set KeywordsThreshold
func (options *CreateJobOptions) SetKeywordsThreshold(param float32) *CreateJobOptions {
    options.KeywordsThreshold = param
    options.IsKeywordsThresholdSet = true
    return options
}

// SetMaxAlternatives : Allow user to set MaxAlternatives
func (options *CreateJobOptions) SetMaxAlternatives(param int64) *CreateJobOptions {
    options.MaxAlternatives = param
    options.IsMaxAlternativesSet = true
    return options
}

// SetWordAlternativesThreshold : Allow user to set WordAlternativesThreshold
func (options *CreateJobOptions) SetWordAlternativesThreshold(param float32) *CreateJobOptions {
    options.WordAlternativesThreshold = param
    options.IsWordAlternativesThresholdSet = true
    return options
}

// SetWordConfidence : Allow user to set WordConfidence
func (options *CreateJobOptions) SetWordConfidence(param bool) *CreateJobOptions {
    options.WordConfidence = param
    options.IsWordConfidenceSet = true
    return options
}

// SetTimestamps : Allow user to set Timestamps
func (options *CreateJobOptions) SetTimestamps(param bool) *CreateJobOptions {
    options.Timestamps = param
    options.IsTimestampsSet = true
    return options
}

// SetProfanityFilter : Allow user to set ProfanityFilter
func (options *CreateJobOptions) SetProfanityFilter(param bool) *CreateJobOptions {
    options.ProfanityFilter = param
    options.IsProfanityFilterSet = true
    return options
}

// SetSmartFormatting : Allow user to set SmartFormatting
func (options *CreateJobOptions) SetSmartFormatting(param bool) *CreateJobOptions {
    options.SmartFormatting = param
    options.IsSmartFormattingSet = true
    return options
}

// SetSpeakerLabels : Allow user to set SpeakerLabels
func (options *CreateJobOptions) SetSpeakerLabels(param bool) *CreateJobOptions {
    options.SpeakerLabels = param
    options.IsSpeakerLabelsSet = true
    return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateJobOptions) SetHeaders(param map[string]string) *CreateJobOptions {
    options.Headers = param
    return options
}

// CreateLanguageModelOptions : The createLanguageModel options.
type CreateLanguageModelOptions struct {

	// A user-defined name for the new custom language model. Use a name that is unique among all custom language models that you own. Use a localized name that matches the language of the custom model. Use a name that describes the domain of the custom model, such as `Medical custom model` or `Legal custom model`.
	Name string `json:"name"`

	// The name of the base language model that is to be customized by the new custom language model. The new custom model can be used only with the base model that it customizes. To determine whether a base model supports language model customization, request information about the base model and check that the attribute `custom_language_model` is set to `true`, or refer to [Language support for customization](https://console.bluemix.net/docs/services/speech-to-text/custom.html#languageSupport).
	BaseModelName string `json:"base_model_name"`

	// The dialect of the specified language that is to be used with the custom language model. The parameter is meaningful only for Spanish models, for which the service creates a custom language model that is suited for speech in one of the following dialects: * `es-ES` for Castilian Spanish (the default) * `es-LA` for Latin American Spanish * `es-US` for North American (Mexican) Spanish A specified dialect must be valid for the base model. By default, the dialect matches the language of the base model; for example, `en-US` for either of the US English language models.
	Dialect string `json:"dialect,omitempty"`

    // Indicates whether user set optional parameter Dialect
    IsDialectSet bool

	// A description of the new custom language model. Use a localized description that matches the language of the custom model.
	Description string `json:"description,omitempty"`

    // Indicates whether user set optional parameter Description
    IsDescriptionSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewCreateLanguageModelOptions : Instantiate CreateLanguageModelOptions
func NewCreateLanguageModelOptions(name string, baseModelName string) *CreateLanguageModelOptions {
    return &CreateLanguageModelOptions{
        Name: name,
        BaseModelName: baseModelName,
    }
}

// SetName : Allow user to set Name
func (options *CreateLanguageModelOptions) SetName(param string) *CreateLanguageModelOptions {
    options.Name = param
    return options
}

// SetBaseModelName : Allow user to set BaseModelName
func (options *CreateLanguageModelOptions) SetBaseModelName(param string) *CreateLanguageModelOptions {
    options.BaseModelName = param
    return options
}

// SetDialect : Allow user to set Dialect
func (options *CreateLanguageModelOptions) SetDialect(param string) *CreateLanguageModelOptions {
    options.Dialect = param
    options.IsDialectSet = true
    return options
}

// SetDescription : Allow user to set Description
func (options *CreateLanguageModelOptions) SetDescription(param string) *CreateLanguageModelOptions {
    options.Description = param
    options.IsDescriptionSet = true
    return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateLanguageModelOptions) SetHeaders(param map[string]string) *CreateLanguageModelOptions {
    options.Headers = param
    return options
}

// CustomWord : CustomWord struct
type CustomWord struct {

	// For the **Add custom words** method, you must specify the custom word that is to be added to or updated in the custom model. Do not include spaces in the word. Use a `-` (dash) or `_` (underscore) to connect the tokens of compound words. Omit this field for the **Add a custom word** method.
	Word string `json:"word,omitempty"`

	// An array of sounds-like pronunciations for the custom word. Specify how words that are difficult to pronounce, foreign words, acronyms, and so on can be pronounced by users. For a word that is not in the service's base vocabulary, omit the parameter to have the service automatically generate a sounds-like pronunciation for the word. For a word that is in the service's base vocabulary, use the parameter to specify additional pronunciations for the word. You cannot override the default pronunciation of a word; pronunciations you add augment the pronunciation from the base vocabulary. A word can have at most five sounds-like pronunciations, and a pronunciation can include at most 40 characters not including spaces.
	SoundsLike []string `json:"sounds_like,omitempty"`

	// An alternative spelling for the custom word when it appears in a transcript. Use the parameter when you want the word to have a spelling that is different from its usual representation or from its spelling in corpora training data.
	DisplayAs string `json:"display_as,omitempty"`
}

// DeleteAcousticModelOptions : The deleteAcousticModel options.
type DeleteAcousticModelOptions struct {

	// The customization ID (GUID) of the custom acoustic model. You must make the request with service credentials created for the instance of the service that owns the custom model.
	CustomizationID string `json:"customization_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewDeleteAcousticModelOptions : Instantiate DeleteAcousticModelOptions
func NewDeleteAcousticModelOptions(customizationID string) *DeleteAcousticModelOptions {
    return &DeleteAcousticModelOptions{
        CustomizationID: customizationID,
    }
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *DeleteAcousticModelOptions) SetCustomizationID(param string) *DeleteAcousticModelOptions {
    options.CustomizationID = param
    return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteAcousticModelOptions) SetHeaders(param map[string]string) *DeleteAcousticModelOptions {
    options.Headers = param
    return options
}

// DeleteAudioOptions : The deleteAudio options.
type DeleteAudioOptions struct {

	// The customization ID (GUID) of the custom acoustic model. You must make the request with service credentials created for the instance of the service that owns the custom model.
	CustomizationID string `json:"customization_id"`

	// The name of the audio resource for the custom acoustic model. When adding an audio resource, do not include spaces in the name; use a localized name that matches the language of the custom model.
	AudioName string `json:"audio_name"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewDeleteAudioOptions : Instantiate DeleteAudioOptions
func NewDeleteAudioOptions(customizationID string, audioName string) *DeleteAudioOptions {
    return &DeleteAudioOptions{
        CustomizationID: customizationID,
        AudioName: audioName,
    }
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *DeleteAudioOptions) SetCustomizationID(param string) *DeleteAudioOptions {
    options.CustomizationID = param
    return options
}

// SetAudioName : Allow user to set AudioName
func (options *DeleteAudioOptions) SetAudioName(param string) *DeleteAudioOptions {
    options.AudioName = param
    return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteAudioOptions) SetHeaders(param map[string]string) *DeleteAudioOptions {
    options.Headers = param
    return options
}

// DeleteCorpusOptions : The deleteCorpus options.
type DeleteCorpusOptions struct {

	// The customization ID (GUID) of the custom language model. You must make the request with service credentials created for the instance of the service that owns the custom model.
	CustomizationID string `json:"customization_id"`

	// The name of the corpus for the custom language model. When adding a corpus, do not include spaces in the name; use a localized name that matches the language of the custom model; and do not use the name `user`, which is reserved by the service to denote custom words added or modified by the user.
	CorpusName string `json:"corpus_name"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewDeleteCorpusOptions : Instantiate DeleteCorpusOptions
func NewDeleteCorpusOptions(customizationID string, corpusName string) *DeleteCorpusOptions {
    return &DeleteCorpusOptions{
        CustomizationID: customizationID,
        CorpusName: corpusName,
    }
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *DeleteCorpusOptions) SetCustomizationID(param string) *DeleteCorpusOptions {
    options.CustomizationID = param
    return options
}

// SetCorpusName : Allow user to set CorpusName
func (options *DeleteCorpusOptions) SetCorpusName(param string) *DeleteCorpusOptions {
    options.CorpusName = param
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
	ID string `json:"id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewDeleteJobOptions : Instantiate DeleteJobOptions
func NewDeleteJobOptions(id string) *DeleteJobOptions {
    return &DeleteJobOptions{
        ID: id,
    }
}

// SetID : Allow user to set ID
func (options *DeleteJobOptions) SetID(param string) *DeleteJobOptions {
    options.ID = param
    return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteJobOptions) SetHeaders(param map[string]string) *DeleteJobOptions {
    options.Headers = param
    return options
}

// DeleteLanguageModelOptions : The deleteLanguageModel options.
type DeleteLanguageModelOptions struct {

	// The customization ID (GUID) of the custom language model. You must make the request with service credentials created for the instance of the service that owns the custom model.
	CustomizationID string `json:"customization_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewDeleteLanguageModelOptions : Instantiate DeleteLanguageModelOptions
func NewDeleteLanguageModelOptions(customizationID string) *DeleteLanguageModelOptions {
    return &DeleteLanguageModelOptions{
        CustomizationID: customizationID,
    }
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *DeleteLanguageModelOptions) SetCustomizationID(param string) *DeleteLanguageModelOptions {
    options.CustomizationID = param
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
	CustomerID string `json:"customer_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewDeleteUserDataOptions : Instantiate DeleteUserDataOptions
func NewDeleteUserDataOptions(customerID string) *DeleteUserDataOptions {
    return &DeleteUserDataOptions{
        CustomerID: customerID,
    }
}

// SetCustomerID : Allow user to set CustomerID
func (options *DeleteUserDataOptions) SetCustomerID(param string) *DeleteUserDataOptions {
    options.CustomerID = param
    return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteUserDataOptions) SetHeaders(param map[string]string) *DeleteUserDataOptions {
    options.Headers = param
    return options
}

// DeleteWordOptions : The deleteWord options.
type DeleteWordOptions struct {

	// The customization ID (GUID) of the custom language model. You must make the request with service credentials created for the instance of the service that owns the custom model.
	CustomizationID string `json:"customization_id"`

	// The custom word for the custom language model. When you add or update a custom word with the **Add a custom word** method, do not include spaces in the word. Use a `-` (dash) or `_` (underscore) to connect the tokens of compound words.
	WordName string `json:"word_name"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewDeleteWordOptions : Instantiate DeleteWordOptions
func NewDeleteWordOptions(customizationID string, wordName string) *DeleteWordOptions {
    return &DeleteWordOptions{
        CustomizationID: customizationID,
        WordName: wordName,
    }
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *DeleteWordOptions) SetCustomizationID(param string) *DeleteWordOptions {
    options.CustomizationID = param
    return options
}

// SetWordName : Allow user to set WordName
func (options *DeleteWordOptions) SetWordName(param string) *DeleteWordOptions {
    options.WordName = param
    return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteWordOptions) SetHeaders(param map[string]string) *DeleteWordOptions {
    options.Headers = param
    return options
}

// GetAcousticModelOptions : The getAcousticModel options.
type GetAcousticModelOptions struct {

	// The customization ID (GUID) of the custom acoustic model. You must make the request with service credentials created for the instance of the service that owns the custom model.
	CustomizationID string `json:"customization_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewGetAcousticModelOptions : Instantiate GetAcousticModelOptions
func NewGetAcousticModelOptions(customizationID string) *GetAcousticModelOptions {
    return &GetAcousticModelOptions{
        CustomizationID: customizationID,
    }
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *GetAcousticModelOptions) SetCustomizationID(param string) *GetAcousticModelOptions {
    options.CustomizationID = param
    return options
}

// SetHeaders : Allow user to set Headers
func (options *GetAcousticModelOptions) SetHeaders(param map[string]string) *GetAcousticModelOptions {
    options.Headers = param
    return options
}

// GetAudioOptions : The getAudio options.
type GetAudioOptions struct {

	// The customization ID (GUID) of the custom acoustic model. You must make the request with service credentials created for the instance of the service that owns the custom model.
	CustomizationID string `json:"customization_id"`

	// The name of the audio resource for the custom acoustic model. When adding an audio resource, do not include spaces in the name; use a localized name that matches the language of the custom model.
	AudioName string `json:"audio_name"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewGetAudioOptions : Instantiate GetAudioOptions
func NewGetAudioOptions(customizationID string, audioName string) *GetAudioOptions {
    return &GetAudioOptions{
        CustomizationID: customizationID,
        AudioName: audioName,
    }
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *GetAudioOptions) SetCustomizationID(param string) *GetAudioOptions {
    options.CustomizationID = param
    return options
}

// SetAudioName : Allow user to set AudioName
func (options *GetAudioOptions) SetAudioName(param string) *GetAudioOptions {
    options.AudioName = param
    return options
}

// SetHeaders : Allow user to set Headers
func (options *GetAudioOptions) SetHeaders(param map[string]string) *GetAudioOptions {
    options.Headers = param
    return options
}

// GetCorpusOptions : The getCorpus options.
type GetCorpusOptions struct {

	// The customization ID (GUID) of the custom language model. You must make the request with service credentials created for the instance of the service that owns the custom model.
	CustomizationID string `json:"customization_id"`

	// The name of the corpus for the custom language model. When adding a corpus, do not include spaces in the name; use a localized name that matches the language of the custom model; and do not use the name `user`, which is reserved by the service to denote custom words added or modified by the user.
	CorpusName string `json:"corpus_name"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewGetCorpusOptions : Instantiate GetCorpusOptions
func NewGetCorpusOptions(customizationID string, corpusName string) *GetCorpusOptions {
    return &GetCorpusOptions{
        CustomizationID: customizationID,
        CorpusName: corpusName,
    }
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *GetCorpusOptions) SetCustomizationID(param string) *GetCorpusOptions {
    options.CustomizationID = param
    return options
}

// SetCorpusName : Allow user to set CorpusName
func (options *GetCorpusOptions) SetCorpusName(param string) *GetCorpusOptions {
    options.CorpusName = param
    return options
}

// SetHeaders : Allow user to set Headers
func (options *GetCorpusOptions) SetHeaders(param map[string]string) *GetCorpusOptions {
    options.Headers = param
    return options
}

// GetLanguageModelOptions : The getLanguageModel options.
type GetLanguageModelOptions struct {

	// The customization ID (GUID) of the custom language model. You must make the request with service credentials created for the instance of the service that owns the custom model.
	CustomizationID string `json:"customization_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewGetLanguageModelOptions : Instantiate GetLanguageModelOptions
func NewGetLanguageModelOptions(customizationID string) *GetLanguageModelOptions {
    return &GetLanguageModelOptions{
        CustomizationID: customizationID,
    }
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *GetLanguageModelOptions) SetCustomizationID(param string) *GetLanguageModelOptions {
    options.CustomizationID = param
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
	ModelID string `json:"model_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewGetModelOptions : Instantiate GetModelOptions
func NewGetModelOptions(modelID string) *GetModelOptions {
    return &GetModelOptions{
        ModelID: modelID,
    }
}

// SetModelID : Allow user to set ModelID
func (options *GetModelOptions) SetModelID(param string) *GetModelOptions {
    options.ModelID = param
    return options
}

// SetHeaders : Allow user to set Headers
func (options *GetModelOptions) SetHeaders(param map[string]string) *GetModelOptions {
    options.Headers = param
    return options
}

// GetWordOptions : The getWord options.
type GetWordOptions struct {

	// The customization ID (GUID) of the custom language model. You must make the request with service credentials created for the instance of the service that owns the custom model.
	CustomizationID string `json:"customization_id"`

	// The custom word for the custom language model. When you add or update a custom word with the **Add a custom word** method, do not include spaces in the word. Use a `-` (dash) or `_` (underscore) to connect the tokens of compound words.
	WordName string `json:"word_name"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewGetWordOptions : Instantiate GetWordOptions
func NewGetWordOptions(customizationID string, wordName string) *GetWordOptions {
    return &GetWordOptions{
        CustomizationID: customizationID,
        WordName: wordName,
    }
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *GetWordOptions) SetCustomizationID(param string) *GetWordOptions {
    options.CustomizationID = param
    return options
}

// SetWordName : Allow user to set WordName
func (options *GetWordOptions) SetWordName(param string) *GetWordOptions {
    options.WordName = param
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
	NormalizedText string `json:"normalized_text"`

	// The start time in seconds of the keyword match.
	StartTime float64 `json:"start_time"`

	// The end time in seconds of the keyword match.
	EndTime float64 `json:"end_time"`

	// A confidence score for the keyword match in the range of 0.0 to 1.0.
	Confidence float64 `json:"confidence"`
}

// LanguageModel : LanguageModel struct
type LanguageModel struct {

	// The customization ID (GUID) of the custom language model. The **Create a custom language model** method returns only this field of the object; it does not return the other fields.
	CustomizationID string `json:"customization_id"`

	// The date and time in Coordinated Universal Time (UTC) at which the custom language model was created. The value is provided in full ISO 8601 format (`YYYY-MM-DDThh:mm:ss.sTZD`).
	Created string `json:"created,omitempty"`

	// The language identifier of the custom language model (for example, `en-US`).
	Language string `json:"language,omitempty"`

	// The dialect of the language for the custom language model. By default, the dialect matches the language of the base model; for example, `en-US` for either of the US English language models. For Spanish models, the field indicates the dialect for which the model was created: * `es-ES` for Castilian Spanish (the default) * `es-LA` for Latin American Spanish * `es-US` for North American (Mexican) Spanish.
	Dialect string `json:"dialect,omitempty"`

	// A list of the available versions of the custom language model. Each element of the array indicates a version of the base model with which the custom model can be used. Multiple versions exist only if the custom model has been upgraded; otherwise, only a single version is shown.
	Versions []string `json:"versions,omitempty"`

	// The GUID of the service credentials for the instance of the service that owns the custom language model.
	Owner string `json:"owner,omitempty"`

	// The name of the custom language model.
	Name string `json:"name,omitempty"`

	// The description of the custom language model.
	Description string `json:"description,omitempty"`

	// The name of the language model for which the custom language model was created.
	BaseModelName string `json:"base_model_name,omitempty"`

	// The current status of the custom language model: * `pending` indicates that the model was created but is waiting either for training data to be added or for the service to finish analyzing added data. * `ready` indicates that the model contains data and is ready to be trained. * `training` indicates that the model is currently being trained. * `available` indicates that the model is trained and ready to use. * `upgrading` indicates that the model is currently being upgraded. * `failed` indicates that training of the model failed.
	Status string `json:"status,omitempty"`

	// A percentage that indicates the progress of the custom language model's current training. A value of `100` means that the model is fully trained. **Note:** The `progress` field does not currently reflect the progress of the training. The field changes from `0` to `100` when training is complete.
	Progress int64 `json:"progress,omitempty"`

	// If the request included unknown parameters, the following message: `Unexpected query parameter(s) ['parameters'] detected`, where `parameters` is a list that includes a quoted string for each unknown parameter.
	Warnings string `json:"warnings,omitempty"`
}

// LanguageModels : LanguageModels struct
type LanguageModels struct {

	// An array of objects that provides information about each available custom language model. The array is empty if the requesting service credentials own no custom language models (if no language is specified) or own no custom language models for the specified language.
	Customizations []LanguageModel `json:"customizations"`
}

// ListAcousticModelsOptions : The listAcousticModels options.
type ListAcousticModelsOptions struct {

	// The identifier of the language for which custom language or custom acoustic models are to be returned (for example, `en-US`). Omit the parameter to see all custom language or custom acoustic models owned by the requesting service credentials.
	Language string `json:"language,omitempty"`

    // Indicates whether user set optional parameter Language
    IsLanguageSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewListAcousticModelsOptions : Instantiate ListAcousticModelsOptions
func NewListAcousticModelsOptions() *ListAcousticModelsOptions {
    return &ListAcousticModelsOptions{}
}

// SetLanguage : Allow user to set Language
func (options *ListAcousticModelsOptions) SetLanguage(param string) *ListAcousticModelsOptions {
    options.Language = param
    options.IsLanguageSet = true
    return options
}

// SetHeaders : Allow user to set Headers
func (options *ListAcousticModelsOptions) SetHeaders(param map[string]string) *ListAcousticModelsOptions {
    options.Headers = param
    return options
}

// ListAudioOptions : The listAudio options.
type ListAudioOptions struct {

	// The customization ID (GUID) of the custom acoustic model. You must make the request with service credentials created for the instance of the service that owns the custom model.
	CustomizationID string `json:"customization_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewListAudioOptions : Instantiate ListAudioOptions
func NewListAudioOptions(customizationID string) *ListAudioOptions {
    return &ListAudioOptions{
        CustomizationID: customizationID,
    }
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *ListAudioOptions) SetCustomizationID(param string) *ListAudioOptions {
    options.CustomizationID = param
    return options
}

// SetHeaders : Allow user to set Headers
func (options *ListAudioOptions) SetHeaders(param map[string]string) *ListAudioOptions {
    options.Headers = param
    return options
}

// ListCorporaOptions : The listCorpora options.
type ListCorporaOptions struct {

	// The customization ID (GUID) of the custom language model. You must make the request with service credentials created for the instance of the service that owns the custom model.
	CustomizationID string `json:"customization_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewListCorporaOptions : Instantiate ListCorporaOptions
func NewListCorporaOptions(customizationID string) *ListCorporaOptions {
    return &ListCorporaOptions{
        CustomizationID: customizationID,
    }
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *ListCorporaOptions) SetCustomizationID(param string) *ListCorporaOptions {
    options.CustomizationID = param
    return options
}

// SetHeaders : Allow user to set Headers
func (options *ListCorporaOptions) SetHeaders(param map[string]string) *ListCorporaOptions {
    options.Headers = param
    return options
}

// ListLanguageModelsOptions : The listLanguageModels options.
type ListLanguageModelsOptions struct {

	// The identifier of the language for which custom language or custom acoustic models are to be returned (for example, `en-US`). Omit the parameter to see all custom language or custom acoustic models owned by the requesting service credentials.
	Language string `json:"language,omitempty"`

    // Indicates whether user set optional parameter Language
    IsLanguageSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewListLanguageModelsOptions : Instantiate ListLanguageModelsOptions
func NewListLanguageModelsOptions() *ListLanguageModelsOptions {
    return &ListLanguageModelsOptions{}
}

// SetLanguage : Allow user to set Language
func (options *ListLanguageModelsOptions) SetLanguage(param string) *ListLanguageModelsOptions {
    options.Language = param
    options.IsLanguageSet = true
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
func NewListModelsOptions() *ListModelsOptions {
    return &ListModelsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListModelsOptions) SetHeaders(param map[string]string) *ListModelsOptions {
    options.Headers = param
    return options
}

// ListWordsOptions : The listWords options.
type ListWordsOptions struct {

	// The customization ID (GUID) of the custom language model. You must make the request with service credentials created for the instance of the service that owns the custom model.
	CustomizationID string `json:"customization_id"`

	// The type of words to be listed from the custom language model's words resource: * `all` (the default) shows all words. * `user` shows only custom words that were added or modified by the user. * `corpora` shows only OOV that were extracted from corpora.
	WordType string `json:"word_type,omitempty"`

    // Indicates whether user set optional parameter WordType
    IsWordTypeSet bool

	// Indicates the order in which the words are to be listed, `alphabetical` or by `count`. You can prepend an optional `+` or `-` to an argument to indicate whether the results are to be sorted in ascending or descending order. By default, words are sorted in ascending alphabetical order. For alphabetical ordering, the lexicographical precedence is numeric values, uppercase letters, and lowercase letters. For count ordering, values with the same count are ordered alphabetically. With cURL, URL encode the `+` symbol as `%2B`.
	Sort string `json:"sort,omitempty"`

    // Indicates whether user set optional parameter Sort
    IsSortSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewListWordsOptions : Instantiate ListWordsOptions
func NewListWordsOptions(customizationID string) *ListWordsOptions {
    return &ListWordsOptions{
        CustomizationID: customizationID,
    }
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *ListWordsOptions) SetCustomizationID(param string) *ListWordsOptions {
    options.CustomizationID = param
    return options
}

// SetWordType : Allow user to set WordType
func (options *ListWordsOptions) SetWordType(param string) *ListWordsOptions {
    options.WordType = param
    options.IsWordTypeSet = true
    return options
}

// SetSort : Allow user to set Sort
func (options *ListWordsOptions) SetSort(param string) *ListWordsOptions {
    options.Sort = param
    options.IsSortSet = true
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
	ID string `json:"id"`

	// The current status of the job: * `waiting`: The service is preparing the job for processing. The service returns this status when the job is initially created or when it is waiting for capacity to process the job. The job remains in this state until the service has the capacity to begin processing it. * `processing`: The service is actively processing the job. * `completed`: The service has finished processing the job. If the job specified a callback URL and the event `recognitions.completed_with_results`, the service sent the results with the callback notification; otherwise, you must retrieve the results by checking the individual job. * `failed`: The job failed.
	Status string `json:"status"`

	// The date and time in Coordinated Universal Time (UTC) at which the job was created. The value is provided in full ISO 8601 format (`YYYY-MM-DDThh:mm:ss.sTZD`).
	Created string `json:"created"`

	// The date and time in Coordinated Universal Time (UTC) at which the job was last updated by the service. The value is provided in full ISO 8601 format (`YYYY-MM-DDThh:mm:ss.sTZD`). This field is returned only by the **Check jobs** and **Check a job** methods.
	Updated string `json:"updated,omitempty"`

	// The URL to use to request information about the job with the **Check a job** method. This field is returned only by the **Create a job** method.
	URL string `json:"url,omitempty"`

	// The user token associated with a job that was created with a callback URL and a user token. This field can be returned only by the **Check jobs** method.
	UserToken string `json:"user_token,omitempty"`

	// If the status is `completed`, the results of the recognition request as an array that includes a single instance of a `SpeechRecognitionResults` object. This field is returned only by the **Check a job** method.
	Results []SpeechRecognitionResults `json:"results,omitempty"`

	// An array of warning messages about invalid parameters included with the request. Each warning includes a descriptive message and a list of invalid argument strings, for example, `"unexpected query parameter 'user_token', query parameter 'callback_url' was not specified"`. The request succeeds despite the warnings. This field can be returned only by the **Create a job** method.
	Warnings []string `json:"warnings,omitempty"`
}

// RecognitionJobs : RecognitionJobs struct
type RecognitionJobs struct {

	// An array of objects that provides the status for each of the user's current jobs. The array is empty if the user has no current jobs.
	Recognitions []RecognitionJob `json:"recognitions"`
}

// RecognizeOptions : The recognize options.
type RecognizeOptions struct {

	Audio io.ReadCloser `json:"audio,omitempty"`

	// The type of the input.
	ContentType string `json:"Content-Type"`

	// The identifier of the model that is to be used for the recognition request or, for the **Create a session** method, with the new session.
	Model string `json:"model,omitempty"`

    // Indicates whether user set optional parameter Model
    IsModelSet bool

	// The customization ID (GUID) of a custom language model that is to be used with the recognition request or, for the **Create a session** method, with the new session. The base model of the specified custom language model must match the model specified with the `model` parameter. You must make the request with service credentials created for the instance of the service that owns the custom model. By default, no custom language model is used.
	CustomizationID string `json:"customization_id,omitempty"`

    // Indicates whether user set optional parameter CustomizationID
    IsCustomizationIDSet bool

	// The customization ID (GUID) of a custom acoustic model that is to be used with the recognition request or, for the **Create a session** method, with the new session. The base model of the specified custom acoustic model must match the model specified with the `model` parameter. You must make the request with service credentials created for the instance of the service that owns the custom model. By default, no custom acoustic model is used.
	AcousticCustomizationID string `json:"acoustic_customization_id,omitempty"`

    // Indicates whether user set optional parameter AcousticCustomizationID
    IsAcousticCustomizationIDSet bool

	// The version of the specified base model that is to be used with recognition request or, for the **Create a session** method, with the new session. Multiple versions of a base model can exist when a model is updated for internal improvements. The parameter is intended primarily for use with custom models that have been upgraded for a new base model. The default value depends on whether the parameter is used with or without a custom model. For more information, see [Base model version](https://console.bluemix.net/docs/services/speech-to-text/input.html#version).
	BaseModelVersion string `json:"base_model_version,omitempty"`

    // Indicates whether user set optional parameter BaseModelVersion
    IsBaseModelVersionSet bool

	// If you specify the customization ID (GUID) of a custom language model with the recognition request or, for sessions, with the **Create a session** method, the customization weight tells the service how much weight to give to words from the custom language model compared to those from the base model for the current request. Specify a value between 0.0 and 1.0. Unless a different customization weight was specified for the custom model when it was trained, the default value is 0.3. A customization weight that you specify overrides a weight that was specified when the custom model was trained. The default value yields the best performance in general. Assign a higher value if your audio makes frequent use of OOV words from the custom model. Use caution when setting the weight: a higher value can improve the accuracy of phrases from the custom model's domain, but it can negatively affect performance on non-domain phrases.
	CustomizationWeight float64 `json:"customization_weight,omitempty"`

    // Indicates whether user set optional parameter CustomizationWeight
    IsCustomizationWeightSet bool

	// The time in seconds after which, if only silence (no speech) is detected in submitted audio, the connection is closed with a 400 error. The parameter is useful for stopping audio submission from a live microphone when a user simply walks away. Use `-1` for infinity.
	InactivityTimeout int64 `json:"inactivity_timeout,omitempty"`

    // Indicates whether user set optional parameter InactivityTimeout
    IsInactivityTimeoutSet bool

	// An array of keyword strings to spot in the audio. Each keyword string can include one or more string tokens. Keywords are spotted only in the final results, not in interim hypotheses. If you specify any keywords, you must also specify a keywords threshold. You can spot a maximum of 1000 keywords. Omit the parameter or specify an empty array if you do not need to spot keywords.
	Keywords []string `json:"keywords,omitempty"`

    // Indicates whether user set optional parameter Keywords
    IsKeywordsSet bool

	// A confidence value that is the lower bound for spotting a keyword. A word is considered to match a keyword if its confidence is greater than or equal to the threshold. Specify a probability between 0.0 and 1.0. No keyword spotting is performed if you omit the parameter. If you specify a threshold, you must also specify one or more keywords.
	KeywordsThreshold float32 `json:"keywords_threshold,omitempty"`

    // Indicates whether user set optional parameter KeywordsThreshold
    IsKeywordsThresholdSet bool

	// The maximum number of alternative transcripts that the service is to return. By default, a single transcription is returned.
	MaxAlternatives int64 `json:"max_alternatives,omitempty"`

    // Indicates whether user set optional parameter MaxAlternatives
    IsMaxAlternativesSet bool

	// A confidence value that is the lower bound for identifying a hypothesis as a possible word alternative (also known as "Confusion Networks"). An alternative word is considered if its confidence is greater than or equal to the threshold. Specify a probability between 0.0 and 1.0. No alternative words are computed if you omit the parameter.
	WordAlternativesThreshold float32 `json:"word_alternatives_threshold,omitempty"`

    // Indicates whether user set optional parameter WordAlternativesThreshold
    IsWordAlternativesThresholdSet bool

	// If `true`, the service returns a confidence measure in the range of 0.0 to 1.0 for each word. By default, no word confidence measures are returned.
	WordConfidence bool `json:"word_confidence,omitempty"`

    // Indicates whether user set optional parameter WordConfidence
    IsWordConfidenceSet bool

	// If `true`, the service returns time alignment for each word. By default, no timestamps are returned.
	Timestamps bool `json:"timestamps,omitempty"`

    // Indicates whether user set optional parameter Timestamps
    IsTimestampsSet bool

	// If `true`, the service filters profanity from all output except for keyword results by replacing inappropriate words with a series of asterisks. Set the parameter to `false` to return results with no censoring. Applies to US English transcription only.
	ProfanityFilter bool `json:"profanity_filter,omitempty"`

    // Indicates whether user set optional parameter ProfanityFilter
    IsProfanityFilterSet bool

	// If `true`, the service converts dates, times, series of digits and numbers, phone numbers, currency values, and internet addresses into more readable, conventional representations in the final transcript of a recognition request. For US English, the service also converts certain keyword strings to punctuation symbols. By default, no smart formatting is performed. Applies to US English and Spanish transcription only.
	SmartFormatting bool `json:"smart_formatting,omitempty"`

    // Indicates whether user set optional parameter SmartFormatting
    IsSmartFormattingSet bool

	// If `true`, the response includes labels that identify which words were spoken by which participants in a multi-person exchange. By default, no speaker labels are returned. Setting `speaker_labels` to `true` forces the `timestamps` parameter to be `true`, regardless of whether you specify `false` for the parameter. To determine whether a language model supports speaker labels, use the **Get models** method and check that the attribute `speaker_labels` is set to `true`. You can also refer to [Speaker labels](https://console.bluemix.net/docs/services/speech-to-text/output.html#speaker_labels).
	SpeakerLabels bool `json:"speaker_labels,omitempty"`

    // Indicates whether user set optional parameter SpeakerLabels
    IsSpeakerLabelsSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewRecognizeOptionsForBasic : Instantiate RecognizeOptionsForBasic
func NewRecognizeOptionsForBasic(audio io.ReadCloser) *RecognizeOptions {
    return &RecognizeOptions{
        Audio: audio,
        ContentType: "audio/basic",
    }
}

// NewRecognizeOptionsForFlac : Instantiate RecognizeOptionsForFlac
func NewRecognizeOptionsForFlac(audio io.ReadCloser) *RecognizeOptions {
    return &RecognizeOptions{
        Audio: audio,
        ContentType: "audio/flac",
    }
}

// NewRecognizeOptionsForL16 : Instantiate RecognizeOptionsForL16
func NewRecognizeOptionsForL16(audio io.ReadCloser) *RecognizeOptions {
    return &RecognizeOptions{
        Audio: audio,
        ContentType: "audio/l16",
    }
}

// NewRecognizeOptionsForMp3 : Instantiate RecognizeOptionsForMp3
func NewRecognizeOptionsForMp3(audio io.ReadCloser) *RecognizeOptions {
    return &RecognizeOptions{
        Audio: audio,
        ContentType: "audio/mp3",
    }
}

// NewRecognizeOptionsForMpeg : Instantiate RecognizeOptionsForMpeg
func NewRecognizeOptionsForMpeg(audio io.ReadCloser) *RecognizeOptions {
    return &RecognizeOptions{
        Audio: audio,
        ContentType: "audio/mpeg",
    }
}

// NewRecognizeOptionsForMulaw : Instantiate RecognizeOptionsForMulaw
func NewRecognizeOptionsForMulaw(audio io.ReadCloser) *RecognizeOptions {
    return &RecognizeOptions{
        Audio: audio,
        ContentType: "audio/mulaw",
    }
}

// NewRecognizeOptionsForOgg : Instantiate RecognizeOptionsForOgg
func NewRecognizeOptionsForOgg(audio io.ReadCloser) *RecognizeOptions {
    return &RecognizeOptions{
        Audio: audio,
        ContentType: "audio/ogg",
    }
}

// NewRecognizeOptionsForOggcodecsopus : Instantiate RecognizeOptionsForOggcodecsopus
func NewRecognizeOptionsForOggcodecsopus(audio io.ReadCloser) *RecognizeOptions {
    return &RecognizeOptions{
        Audio: audio,
        ContentType: "audio/ogg;codecs=opus",
    }
}

// NewRecognizeOptionsForOggcodecsvorbis : Instantiate RecognizeOptionsForOggcodecsvorbis
func NewRecognizeOptionsForOggcodecsvorbis(audio io.ReadCloser) *RecognizeOptions {
    return &RecognizeOptions{
        Audio: audio,
        ContentType: "audio/ogg;codecs=vorbis",
    }
}

// NewRecognizeOptionsForWav : Instantiate RecognizeOptionsForWav
func NewRecognizeOptionsForWav(audio io.ReadCloser) *RecognizeOptions {
    return &RecognizeOptions{
        Audio: audio,
        ContentType: "audio/wav",
    }
}

// NewRecognizeOptionsForWebm : Instantiate RecognizeOptionsForWebm
func NewRecognizeOptionsForWebm(audio io.ReadCloser) *RecognizeOptions {
    return &RecognizeOptions{
        Audio: audio,
        ContentType: "audio/webm",
    }
}

// NewRecognizeOptionsForWebmcodecsopus : Instantiate RecognizeOptionsForWebmcodecsopus
func NewRecognizeOptionsForWebmcodecsopus(audio io.ReadCloser) *RecognizeOptions {
    return &RecognizeOptions{
        Audio: audio,
        ContentType: "audio/webm;codecs=opus",
    }
}

// NewRecognizeOptionsForWebmcodecsvorbis : Instantiate RecognizeOptionsForWebmcodecsvorbis
func NewRecognizeOptionsForWebmcodecsvorbis(audio io.ReadCloser) *RecognizeOptions {
    return &RecognizeOptions{
        Audio: audio,
        ContentType: "audio/webm;codecs=vorbis",
    }
}

// SetAudio : Allow user to set Audio with specified ContentType
func (options *RecognizeOptions) SetAudio(audio io.ReadCloser, contentType string) *RecognizeOptions {
    options.Audio = audio
    options.ContentType = contentType
    return options
}

// SetModel : Allow user to set Model
func (options *RecognizeOptions) SetModel(param string) *RecognizeOptions {
    options.Model = param
    options.IsModelSet = true
    return options
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *RecognizeOptions) SetCustomizationID(param string) *RecognizeOptions {
    options.CustomizationID = param
    options.IsCustomizationIDSet = true
    return options
}

// SetAcousticCustomizationID : Allow user to set AcousticCustomizationID
func (options *RecognizeOptions) SetAcousticCustomizationID(param string) *RecognizeOptions {
    options.AcousticCustomizationID = param
    options.IsAcousticCustomizationIDSet = true
    return options
}

// SetBaseModelVersion : Allow user to set BaseModelVersion
func (options *RecognizeOptions) SetBaseModelVersion(param string) *RecognizeOptions {
    options.BaseModelVersion = param
    options.IsBaseModelVersionSet = true
    return options
}

// SetCustomizationWeight : Allow user to set CustomizationWeight
func (options *RecognizeOptions) SetCustomizationWeight(param float64) *RecognizeOptions {
    options.CustomizationWeight = param
    options.IsCustomizationWeightSet = true
    return options
}

// SetInactivityTimeout : Allow user to set InactivityTimeout
func (options *RecognizeOptions) SetInactivityTimeout(param int64) *RecognizeOptions {
    options.InactivityTimeout = param
    options.IsInactivityTimeoutSet = true
    return options
}

// SetKeywords : Allow user to set Keywords
func (options *RecognizeOptions) SetKeywords(param []string) *RecognizeOptions {
    options.Keywords = param
    options.IsKeywordsSet = true
    return options
}

// SetKeywordsThreshold : Allow user to set KeywordsThreshold
func (options *RecognizeOptions) SetKeywordsThreshold(param float32) *RecognizeOptions {
    options.KeywordsThreshold = param
    options.IsKeywordsThresholdSet = true
    return options
}

// SetMaxAlternatives : Allow user to set MaxAlternatives
func (options *RecognizeOptions) SetMaxAlternatives(param int64) *RecognizeOptions {
    options.MaxAlternatives = param
    options.IsMaxAlternativesSet = true
    return options
}

// SetWordAlternativesThreshold : Allow user to set WordAlternativesThreshold
func (options *RecognizeOptions) SetWordAlternativesThreshold(param float32) *RecognizeOptions {
    options.WordAlternativesThreshold = param
    options.IsWordAlternativesThresholdSet = true
    return options
}

// SetWordConfidence : Allow user to set WordConfidence
func (options *RecognizeOptions) SetWordConfidence(param bool) *RecognizeOptions {
    options.WordConfidence = param
    options.IsWordConfidenceSet = true
    return options
}

// SetTimestamps : Allow user to set Timestamps
func (options *RecognizeOptions) SetTimestamps(param bool) *RecognizeOptions {
    options.Timestamps = param
    options.IsTimestampsSet = true
    return options
}

// SetProfanityFilter : Allow user to set ProfanityFilter
func (options *RecognizeOptions) SetProfanityFilter(param bool) *RecognizeOptions {
    options.ProfanityFilter = param
    options.IsProfanityFilterSet = true
    return options
}

// SetSmartFormatting : Allow user to set SmartFormatting
func (options *RecognizeOptions) SetSmartFormatting(param bool) *RecognizeOptions {
    options.SmartFormatting = param
    options.IsSmartFormattingSet = true
    return options
}

// SetSpeakerLabels : Allow user to set SpeakerLabels
func (options *RecognizeOptions) SetSpeakerLabels(param bool) *RecognizeOptions {
    options.SpeakerLabels = param
    options.IsSpeakerLabelsSet = true
    return options
}

// SetHeaders : Allow user to set Headers
func (options *RecognizeOptions) SetHeaders(param map[string]string) *RecognizeOptions {
    options.Headers = param
    return options
}

// RegisterCallbackOptions : The registerCallback options.
type RegisterCallbackOptions struct {

	// An HTTP or HTTPS URL to which callback notifications are to be sent. To be white-listed, the URL must successfully echo the challenge string during URL verification. During verification, the client can also check the signature that the service sends in the `X-Callback-Signature` header to verify the origin of the request.
	CallbackURL string `json:"callback_url"`

	// A user-specified string that the service uses to generate the HMAC-SHA1 signature that it sends via the `X-Callback-Signature` header. The service includes the header during URL verification and with every notification sent to the callback URL. It calculates the signature over the payload of the notification. If you omit the parameter, the service does not send the header.
	UserSecret string `json:"user_secret,omitempty"`

    // Indicates whether user set optional parameter UserSecret
    IsUserSecretSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewRegisterCallbackOptions : Instantiate RegisterCallbackOptions
func NewRegisterCallbackOptions(callbackURL string) *RegisterCallbackOptions {
    return &RegisterCallbackOptions{
        CallbackURL: callbackURL,
    }
}

// SetCallbackURL : Allow user to set CallbackURL
func (options *RegisterCallbackOptions) SetCallbackURL(param string) *RegisterCallbackOptions {
    options.CallbackURL = param
    return options
}

// SetUserSecret : Allow user to set UserSecret
func (options *RegisterCallbackOptions) SetUserSecret(param string) *RegisterCallbackOptions {
    options.UserSecret = param
    options.IsUserSecretSet = true
    return options
}

// SetHeaders : Allow user to set Headers
func (options *RegisterCallbackOptions) SetHeaders(param map[string]string) *RegisterCallbackOptions {
    options.Headers = param
    return options
}

// RegisterStatus : RegisterStatus struct
type RegisterStatus struct {

	// The current status of the job: * `created` if the callback URL was successfully white-listed as a result of the call. * `already created` if the URL was already white-listed.
	Status string `json:"status"`

	// The callback URL that is successfully registered.
	URL string `json:"url"`
}

// ResetAcousticModelOptions : The resetAcousticModel options.
type ResetAcousticModelOptions struct {

	// The customization ID (GUID) of the custom acoustic model. You must make the request with service credentials created for the instance of the service that owns the custom model.
	CustomizationID string `json:"customization_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewResetAcousticModelOptions : Instantiate ResetAcousticModelOptions
func NewResetAcousticModelOptions(customizationID string) *ResetAcousticModelOptions {
    return &ResetAcousticModelOptions{
        CustomizationID: customizationID,
    }
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *ResetAcousticModelOptions) SetCustomizationID(param string) *ResetAcousticModelOptions {
    options.CustomizationID = param
    return options
}

// SetHeaders : Allow user to set Headers
func (options *ResetAcousticModelOptions) SetHeaders(param map[string]string) *ResetAcousticModelOptions {
    options.Headers = param
    return options
}

// ResetLanguageModelOptions : The resetLanguageModel options.
type ResetLanguageModelOptions struct {

	// The customization ID (GUID) of the custom language model. You must make the request with service credentials created for the instance of the service that owns the custom model.
	CustomizationID string `json:"customization_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewResetLanguageModelOptions : Instantiate ResetLanguageModelOptions
func NewResetLanguageModelOptions(customizationID string) *ResetLanguageModelOptions {
    return &ResetLanguageModelOptions{
        CustomizationID: customizationID,
    }
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *ResetLanguageModelOptions) SetCustomizationID(param string) *ResetLanguageModelOptions {
    options.CustomizationID = param
    return options
}

// SetHeaders : Allow user to set Headers
func (options *ResetLanguageModelOptions) SetHeaders(param map[string]string) *ResetLanguageModelOptions {
    options.Headers = param
    return options
}

// SpeakerLabelsResult : SpeakerLabelsResult struct
type SpeakerLabelsResult struct {

	// The start time of a word from the transcript. The value matches the start time of a word from the `timestamps` array.
	From float32 `json:"from"`

	// The end time of a word from the transcript. The value matches the end time of a word from the `timestamps` array.
	To float32 `json:"to"`

	// The numeric identifier that the service assigns to a speaker from the audio. Speaker IDs begin at `0` initially but can evolve and change across interim results (if supported by the method) and between interim and final results as the service processes the audio. They are not guaranteed to be sequential, contiguous, or ordered.
	Speaker int64 `json:"speaker"`

	// A score that indicates the service's confidence in its identification of the speaker in the range of 0.0 to 1.0.
	Confidence float32 `json:"confidence"`

	// An indication of whether the service might further change word and speaker-label results. A value of `true` means that the service guarantees not to send any further updates for the current or any preceding results; `false` means that the service might send further updates to the results.
	FinalResults bool `json:"final"`
}

// SpeechModel : SpeechModel struct
type SpeechModel struct {

	// The name of the model for use as an identifier in calls to the service (for example, `en-US_BroadbandModel`).
	Name string `json:"name"`

	// The language identifier of the model (for example, `en-US`).
	Language string `json:"language"`

	// The sampling rate (minimum acceptable rate for audio) used by the model in Hertz.
	Rate int64 `json:"rate"`

	// The URI for the model.
	URL string `json:"url"`

	// Describes the additional service features supported with the model.
	SupportedFeatures SupportedFeatures `json:"supported_features"`

	// Brief description of the model.
	Description string `json:"description"`
}

// SpeechModels : SpeechModels struct
type SpeechModels struct {

	// An array of objects that provides information about each available model.
	Models []SpeechModel `json:"models"`
}

// SpeechRecognitionAlternative : SpeechRecognitionAlternative struct
type SpeechRecognitionAlternative struct {

	// A transcription of the audio.
	Transcript string `json:"transcript"`

	// A score that indicates the service's confidence in the transcript in the range of 0.0 to 1.0. Returned only for the best alternative and only with results marked as final.
	Confidence float64 `json:"confidence,omitempty"`

	// Time alignments for each word from the transcript as a list of lists. Each inner list consists of three elements: the word followed by its start and end time in seconds, for example: `[["hello",0.0,1.2],["world",1.2,2.5]]`. Returned only for the best alternative.
	Timestamps []string `json:"timestamps,omitempty"`

	// A confidence score for each word of the transcript as a list of lists. Each inner list consists of two elements: the word and its confidence score in the range of 0.0 to 1.0, for example: `[["hello",0.95],["world",0.866]]`. Returned only for the best alternative and only with results marked as final.
	WordConfidence []string `json:"word_confidence,omitempty"`
}

// SpeechRecognitionResult : SpeechRecognitionResult struct
type SpeechRecognitionResult struct {

	// An indication of whether the transcription results are final. If `true`, the results for this utterance are not updated further; no additional results are sent for a `result_index` once its results are indicated as final.
	FinalResults bool `json:"final"`

	// An array of alternative transcripts. The `alternatives` array can include additional requested output such as word confidence or timestamps.
	Alternatives []SpeechRecognitionAlternative `json:"alternatives"`

	// A dictionary (or associative array) whose keys are the strings specified for `keywords` if both that parameter and `keywords_threshold` are specified. A keyword for which no matches are found is omitted from the array. The array is omitted if no matches are found for any keywords.
	KeywordsResult map[string][]KeywordResult `json:"keywords_result,omitempty"`

	// An array of alternative hypotheses found for words of the input audio if a `word_alternatives_threshold` is specified.
	WordAlternatives []WordAlternativeResults `json:"word_alternatives,omitempty"`
}

// SpeechRecognitionResults : SpeechRecognitionResults struct
type SpeechRecognitionResults struct {

	// An array that can include interim and final results (interim results are returned only if supported by the method). Final results are guaranteed not to change; interim results might be replaced by further interim results and final results. The service periodically sends updates to the results list; the `result_index` is set to the lowest index in the array that has changed; it is incremented for new results.
	Results []SpeechRecognitionResult `json:"results,omitempty"`

	// An index that indicates a change point in the `results` array. The service increments the index only for additional results that it sends for new audio for the same request.
	ResultIndex int64 `json:"result_index,omitempty"`

	// An array that identifies which words were spoken by which speakers in a multi-person exchange. Returned in the response only if `speaker_labels` is `true`. When interim results are also requested for methods that support them, it is possible for a `SpeechRecognitionResults` object to include only the `speaker_labels` field.
	SpeakerLabels []SpeakerLabelsResult `json:"speaker_labels,omitempty"`

	// An array of warning messages associated with the request: * Warnings for invalid parameters or fields can include a descriptive message and a list of invalid argument strings, for example, `"Unknown arguments:"` or `"Unknown url query arguments:"` followed by a list of the form `"invalid_arg_1, invalid_arg_2."` * The following warning is returned if the request passes a custom model that is based on an older version of a base model for which an updated version is available: `"Using previous version of base model, because your custom model has been built with it. Please note that this version will be supported only for a limited time. Consider updating your custom model to the new base model. If you do not do that you will be automatically switched to base model when you used the non-updated custom model."` In both cases, the request succeeds despite the warnings.
	Warnings []string `json:"warnings,omitempty"`
}

// SupportedFeatures : SupportedFeatures struct
type SupportedFeatures struct {

	// Indicates whether the customization interface can be used to create a custom language model based on the language model.
	CustomLanguageModel bool `json:"custom_language_model"`

	// Indicates whether the `speaker_labels` parameter can be used with the language model.
	SpeakerLabels bool `json:"speaker_labels"`
}

// TrainAcousticModelOptions : The trainAcousticModel options.
type TrainAcousticModelOptions struct {

	// The customization ID (GUID) of the custom acoustic model. You must make the request with service credentials created for the instance of the service that owns the custom model.
	CustomizationID string `json:"customization_id"`

	// The customization ID (GUID) of a custom language model that is to be used during training of the custom acoustic model. Specify a custom language model that has been trained with verbatim transcriptions of the audio resources or that contains words that are relevant to the contents of the audio resources.
	CustomLanguageModelID string `json:"custom_language_model_id,omitempty"`

    // Indicates whether user set optional parameter CustomLanguageModelID
    IsCustomLanguageModelIDSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewTrainAcousticModelOptions : Instantiate TrainAcousticModelOptions
func NewTrainAcousticModelOptions(customizationID string) *TrainAcousticModelOptions {
    return &TrainAcousticModelOptions{
        CustomizationID: customizationID,
    }
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *TrainAcousticModelOptions) SetCustomizationID(param string) *TrainAcousticModelOptions {
    options.CustomizationID = param
    return options
}

// SetCustomLanguageModelID : Allow user to set CustomLanguageModelID
func (options *TrainAcousticModelOptions) SetCustomLanguageModelID(param string) *TrainAcousticModelOptions {
    options.CustomLanguageModelID = param
    options.IsCustomLanguageModelIDSet = true
    return options
}

// SetHeaders : Allow user to set Headers
func (options *TrainAcousticModelOptions) SetHeaders(param map[string]string) *TrainAcousticModelOptions {
    options.Headers = param
    return options
}

// TrainLanguageModelOptions : The trainLanguageModel options.
type TrainLanguageModelOptions struct {

	// The customization ID (GUID) of the custom language model. You must make the request with service credentials created for the instance of the service that owns the custom model.
	CustomizationID string `json:"customization_id"`

	// The type of words from the custom language model's words resource on which to train the model: * `all` (the default) trains the model on all new words, regardless of whether they were extracted from corpora or were added or modified by the user. * `user` trains the model only on new words that were added or modified by the user; the model is not trained on new words extracted from corpora.
	WordTypeToAdd string `json:"word_type_to_add,omitempty"`

    // Indicates whether user set optional parameter WordTypeToAdd
    IsWordTypeToAddSet bool

	// Specifies a customization weight for the custom language model. The customization weight tells the service how much weight to give to words from the custom language model compared to those from the base model for speech recognition. Specify a value between 0.0 and 1.0; the default is 0.3. The default value yields the best performance in general. Assign a higher value if your audio makes frequent use of OOV words from the custom model. Use caution when setting the weight: a higher value can improve the accuracy of phrases from the custom model's domain, but it can negatively affect performance on non-domain phrases. The value that you assign is used for all recognition requests that use the model. You can override it for any recognition request by specifying a customization weight for that request.
	CustomizationWeight float64 `json:"customization_weight,omitempty"`

    // Indicates whether user set optional parameter CustomizationWeight
    IsCustomizationWeightSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewTrainLanguageModelOptions : Instantiate TrainLanguageModelOptions
func NewTrainLanguageModelOptions(customizationID string) *TrainLanguageModelOptions {
    return &TrainLanguageModelOptions{
        CustomizationID: customizationID,
    }
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *TrainLanguageModelOptions) SetCustomizationID(param string) *TrainLanguageModelOptions {
    options.CustomizationID = param
    return options
}

// SetWordTypeToAdd : Allow user to set WordTypeToAdd
func (options *TrainLanguageModelOptions) SetWordTypeToAdd(param string) *TrainLanguageModelOptions {
    options.WordTypeToAdd = param
    options.IsWordTypeToAddSet = true
    return options
}

// SetCustomizationWeight : Allow user to set CustomizationWeight
func (options *TrainLanguageModelOptions) SetCustomizationWeight(param float64) *TrainLanguageModelOptions {
    options.CustomizationWeight = param
    options.IsCustomizationWeightSet = true
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
	CallbackURL string `json:"callback_url"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewUnregisterCallbackOptions : Instantiate UnregisterCallbackOptions
func NewUnregisterCallbackOptions(callbackURL string) *UnregisterCallbackOptions {
    return &UnregisterCallbackOptions{
        CallbackURL: callbackURL,
    }
}

// SetCallbackURL : Allow user to set CallbackURL
func (options *UnregisterCallbackOptions) SetCallbackURL(param string) *UnregisterCallbackOptions {
    options.CallbackURL = param
    return options
}

// SetHeaders : Allow user to set Headers
func (options *UnregisterCallbackOptions) SetHeaders(param map[string]string) *UnregisterCallbackOptions {
    options.Headers = param
    return options
}

// UpgradeAcousticModelOptions : The upgradeAcousticModel options.
type UpgradeAcousticModelOptions struct {

	// The customization ID (GUID) of the custom acoustic model. You must make the request with service credentials created for the instance of the service that owns the custom model.
	CustomizationID string `json:"customization_id"`

	// If the custom acoustic model was trained with a custom language model, the customization ID (GUID) of that custom language model. The custom language model must be upgraded before the custom acoustic model can be upgraded.
	CustomLanguageModelID string `json:"custom_language_model_id,omitempty"`

    // Indicates whether user set optional parameter CustomLanguageModelID
    IsCustomLanguageModelIDSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewUpgradeAcousticModelOptions : Instantiate UpgradeAcousticModelOptions
func NewUpgradeAcousticModelOptions(customizationID string) *UpgradeAcousticModelOptions {
    return &UpgradeAcousticModelOptions{
        CustomizationID: customizationID,
    }
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *UpgradeAcousticModelOptions) SetCustomizationID(param string) *UpgradeAcousticModelOptions {
    options.CustomizationID = param
    return options
}

// SetCustomLanguageModelID : Allow user to set CustomLanguageModelID
func (options *UpgradeAcousticModelOptions) SetCustomLanguageModelID(param string) *UpgradeAcousticModelOptions {
    options.CustomLanguageModelID = param
    options.IsCustomLanguageModelIDSet = true
    return options
}

// SetHeaders : Allow user to set Headers
func (options *UpgradeAcousticModelOptions) SetHeaders(param map[string]string) *UpgradeAcousticModelOptions {
    options.Headers = param
    return options
}

// UpgradeLanguageModelOptions : The upgradeLanguageModel options.
type UpgradeLanguageModelOptions struct {

	// The customization ID (GUID) of the custom language model. You must make the request with service credentials created for the instance of the service that owns the custom model.
	CustomizationID string `json:"customization_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewUpgradeLanguageModelOptions : Instantiate UpgradeLanguageModelOptions
func NewUpgradeLanguageModelOptions(customizationID string) *UpgradeLanguageModelOptions {
    return &UpgradeLanguageModelOptions{
        CustomizationID: customizationID,
    }
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *UpgradeLanguageModelOptions) SetCustomizationID(param string) *UpgradeLanguageModelOptions {
    options.CustomizationID = param
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
	Word string `json:"word"`

	// An array of pronunciations for the word. The array can include the sounds-like pronunciation automatically generated by the service if none is provided for the word; the service adds this pronunciation when it finishes processing the word.
	SoundsLike []string `json:"sounds_like"`

	// The spelling of the word that the service uses to display the word in a transcript. The field contains an empty string if no display-as value is provided for the word, in which case the word is displayed as it is spelled.
	DisplayAs string `json:"display_as"`

	// A sum of the number of times the word is found across all corpora. For example, if the word occurs five times in one corpus and seven times in another, its count is `12`. If you add a custom word to a model before it is added by any corpora, the count begins at `1`; if the word is added from a corpus first and later modified, the count reflects only the number of times it is found in corpora.
	Count int64 `json:"count"`

	// An array of sources that describes how the word was added to the custom model's words resource. For OOV words added from a corpus, includes the name of the corpus; if the word was added by multiple corpora, the names of all corpora are listed. If the word was modified or added by the user directly, the field includes the string `user`.
	Source []string `json:"source"`

	// If the service discovered one or more problems that you need to correct for the word's definition, an array that describes each of the errors.
	Error []WordError `json:"error,omitempty"`
}

// WordAlternativeResult : WordAlternativeResult struct
type WordAlternativeResult struct {

	// A confidence score for the word alternative hypothesis in the range of 0.0 to 1.0.
	Confidence float64 `json:"confidence"`

	// An alternative hypothesis for a word from the input audio.
	Word string `json:"word"`
}

// WordAlternativeResults : WordAlternativeResults struct
type WordAlternativeResults struct {

	// The start time in seconds of the word from the input audio that corresponds to the word alternatives.
	StartTime float64 `json:"start_time"`

	// The end time in seconds of the word from the input audio that corresponds to the word alternatives.
	EndTime float64 `json:"end_time"`

	// An array of alternative hypotheses for a word from the input audio.
	Alternatives []WordAlternativeResult `json:"alternatives"`
}

// WordError : WordError struct
type WordError struct {

	// A key-value pair that describes an error associated with the definition of a word in the words resource. Each pair has the format `"element": "message"`, where `element` is the aspect of the definition that caused the problem and `message` describes the problem. The following example describes a problem with one of the word's sounds-like definitions: `"{sounds_like_string}": "Numbers are not allowed in sounds-like. You can try for example '{suggested_string}'."` You must correct the error before you can train the model.
	Element string `json:"element"`
}

// Words : Words struct
type Words struct {

	// An array of objects that provides information about each word in the custom model's words resource. The array is empty if the custom model has no words.
	Words []Word `json:"words"`
}
