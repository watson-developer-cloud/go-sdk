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
    "bytes"
    "fmt"
    "io"
    "os"
    "strings"
    req "github.com/parnurzeal/gorequest"
    watson "golang-sdk"
)

// SpeechToTextV1 : The SpeechToTextV1 service
type SpeechToTextV1 struct {
	client *watson.Client
}

// NewSpeechToTextV1 : Instantiate SpeechToTextV1
func NewSpeechToTextV1(creds watson.Credentials) (*SpeechToTextV1, error) {
    if creds.ServiceURL == "" {
        creds.ServiceURL = "https://stream.watsonplatform.net/speech-to-text/api"
    }

	client, clientErr := watson.NewClient(creds, "speech_to_text")

	if clientErr != nil {
		return nil, clientErr
	}

	return &SpeechToTextV1{ client: client }, nil
}

// GetModel : Get a model
func (speechToText *SpeechToTextV1) GetModel(modelID string) (*watson.WatsonResponse, []error) {
    path := "/v1/models/{model_id}"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{model_id}", modelID, 1)
    request := req.New().Get(creds.ServiceURL + path)

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

// GetGetModelResult : Cast result of GetModel operation
func GetGetModelResult(response *watson.WatsonResponse) *SpeechModel {
    result, ok := response.Result.(*SpeechModel)

    if ok {
        return result
    }

    return nil
}

// ListModels : List models
func (speechToText *SpeechToTextV1) ListModels() (*watson.WatsonResponse, []error) {
    path := "/v1/models"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    request := req.New().Get(creds.ServiceURL + path)

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

// GetListModelsResult : Cast result of ListModels operation
func GetListModelsResult(response *watson.WatsonResponse) *SpeechModels {
    result, ok := response.Result.(*SpeechModels)

    if ok {
        return result
    }

    return nil
}

// Recognize : Recognize audio
func (speechToText *SpeechToTextV1) Recognize(body *io.ReadCloser, contentType string, model string, customizationID string, acousticCustomizationID string, baseModelVersion string, customizationWeight float64, inactivityTimeout int64, keywords []string, keywordsThreshold float32, maxAlternatives int64, wordAlternativesThreshold float32, wordConfidence bool, timestamps bool, profanityFilter bool, smartFormatting bool, speakerLabels bool) (*watson.WatsonResponse, []error) {
    path := "/v1/recognize"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "audio/basic")
    request.Set("Content-Type", fmt.Sprint(contentType))
    request.Query("model=" + fmt.Sprint(model))
    request.Query("customization_id=" + fmt.Sprint(customizationID))
    request.Query("acoustic_customization_id=" + fmt.Sprint(acousticCustomizationID))
    request.Query("base_model_version=" + fmt.Sprint(baseModelVersion))
    request.Query("customization_weight=" + fmt.Sprint(customizationWeight))
    request.Query("inactivity_timeout=" + fmt.Sprint(inactivityTimeout))
    request.Query("keywords=" + fmt.Sprint(keywords))
    request.Query("keywords_threshold=" + fmt.Sprint(keywordsThreshold))
    request.Query("max_alternatives=" + fmt.Sprint(maxAlternatives))
    request.Query("word_alternatives_threshold=" + fmt.Sprint(wordAlternativesThreshold))
    request.Query("word_confidence=" + fmt.Sprint(wordConfidence))
    request.Query("timestamps=" + fmt.Sprint(timestamps))
    request.Query("profanity_filter=" + fmt.Sprint(profanityFilter))
    request.Query("smart_formatting=" + fmt.Sprint(smartFormatting))
    request.Query("speaker_labels=" + fmt.Sprint(speakerLabels))
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

    response.Result = new(SpeechRecognitionResults)
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

// GetRecognizeResult : Cast result of Recognize operation
func GetRecognizeResult(response *watson.WatsonResponse) *SpeechRecognitionResults {
    result, ok := response.Result.(*SpeechRecognitionResults)

    if ok {
        return result
    }

    return nil
}

// CheckJob : Check a job
func (speechToText *SpeechToTextV1) CheckJob(id string) (*watson.WatsonResponse, []error) {
    path := "/v1/recognitions/{id}"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{id}", id, 1)
    request := req.New().Get(creds.ServiceURL + path)

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

// GetCheckJobResult : Cast result of CheckJob operation
func GetCheckJobResult(response *watson.WatsonResponse) *RecognitionJob {
    result, ok := response.Result.(*RecognitionJob)

    if ok {
        return result
    }

    return nil
}

// CheckJobs : Check jobs
func (speechToText *SpeechToTextV1) CheckJobs() (*watson.WatsonResponse, []error) {
    path := "/v1/recognitions"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    request := req.New().Get(creds.ServiceURL + path)

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

// GetCheckJobsResult : Cast result of CheckJobs operation
func GetCheckJobsResult(response *watson.WatsonResponse) *RecognitionJobs {
    result, ok := response.Result.(*RecognitionJobs)

    if ok {
        return result
    }

    return nil
}

// CreateJob : Create a job
func (speechToText *SpeechToTextV1) CreateJob(body *io.ReadCloser, contentType string, model string, callbackURL string, events string, userToken string, resultsTTL int64, customizationID string, acousticCustomizationID string, baseModelVersion string, customizationWeight float64, inactivityTimeout int64, keywords []string, keywordsThreshold float32, maxAlternatives int64, wordAlternativesThreshold float32, wordConfidence bool, timestamps bool, profanityFilter bool, smartFormatting bool, speakerLabels bool) (*watson.WatsonResponse, []error) {
    path := "/v1/recognitions"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "audio/basic")
    request.Set("Content-Type", fmt.Sprint(contentType))
    request.Query("model=" + fmt.Sprint(model))
    request.Query("callback_url=" + fmt.Sprint(callbackURL))
    request.Query("events=" + fmt.Sprint(events))
    request.Query("user_token=" + fmt.Sprint(userToken))
    request.Query("results_ttl=" + fmt.Sprint(resultsTTL))
    request.Query("customization_id=" + fmt.Sprint(customizationID))
    request.Query("acoustic_customization_id=" + fmt.Sprint(acousticCustomizationID))
    request.Query("base_model_version=" + fmt.Sprint(baseModelVersion))
    request.Query("customization_weight=" + fmt.Sprint(customizationWeight))
    request.Query("inactivity_timeout=" + fmt.Sprint(inactivityTimeout))
    request.Query("keywords=" + fmt.Sprint(keywords))
    request.Query("keywords_threshold=" + fmt.Sprint(keywordsThreshold))
    request.Query("max_alternatives=" + fmt.Sprint(maxAlternatives))
    request.Query("word_alternatives_threshold=" + fmt.Sprint(wordAlternativesThreshold))
    request.Query("word_confidence=" + fmt.Sprint(wordConfidence))
    request.Query("timestamps=" + fmt.Sprint(timestamps))
    request.Query("profanity_filter=" + fmt.Sprint(profanityFilter))
    request.Query("smart_formatting=" + fmt.Sprint(smartFormatting))
    request.Query("speaker_labels=" + fmt.Sprint(speakerLabels))
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

    response.Result = new(RecognitionJob)
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

// GetCreateJobResult : Cast result of CreateJob operation
func GetCreateJobResult(response *watson.WatsonResponse) *RecognitionJob {
    result, ok := response.Result.(*RecognitionJob)

    if ok {
        return result
    }

    return nil
}

// DeleteJob : Delete a job
func (speechToText *SpeechToTextV1) DeleteJob(id string) (*watson.WatsonResponse, []error) {
    path := "/v1/recognitions/{id}"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{id}", id, 1)
    request := req.New().Delete(creds.ServiceURL + path)

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


// RegisterCallback : Register a callback
func (speechToText *SpeechToTextV1) RegisterCallback(callbackURL string, userSecret string) (*watson.WatsonResponse, []error) {
    path := "/v1/register_callback"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("callback_url=" + fmt.Sprint(callbackURL))
    request.Query("user_secret=" + fmt.Sprint(userSecret))

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

// GetRegisterCallbackResult : Cast result of RegisterCallback operation
func GetRegisterCallbackResult(response *watson.WatsonResponse) *RegisterStatus {
    result, ok := response.Result.(*RegisterStatus)

    if ok {
        return result
    }

    return nil
}

// UnregisterCallback : Unregister a callback
func (speechToText *SpeechToTextV1) UnregisterCallback(callbackURL string) (*watson.WatsonResponse, []error) {
    path := "/v1/unregister_callback"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("callback_url=" + fmt.Sprint(callbackURL))

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


// CreateLanguageModel : Create a custom language model
func (speechToText *SpeechToTextV1) CreateLanguageModel(body *CreateLanguageModel) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
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

// GetCreateLanguageModelResult : Cast result of CreateLanguageModel operation
func GetCreateLanguageModelResult(response *watson.WatsonResponse) *LanguageModel {
    result, ok := response.Result.(*LanguageModel)

    if ok {
        return result
    }

    return nil
}

// DeleteLanguageModel : Delete a custom language model
func (speechToText *SpeechToTextV1) DeleteLanguageModel(customizationID string) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", customizationID, 1)
    request := req.New().Delete(creds.ServiceURL + path)

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


// GetLanguageModel : Get a custom language model
func (speechToText *SpeechToTextV1) GetLanguageModel(customizationID string) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", customizationID, 1)
    request := req.New().Get(creds.ServiceURL + path)

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

// GetGetLanguageModelResult : Cast result of GetLanguageModel operation
func GetGetLanguageModelResult(response *watson.WatsonResponse) *LanguageModel {
    result, ok := response.Result.(*LanguageModel)

    if ok {
        return result
    }

    return nil
}

// ListLanguageModels : List custom language models
func (speechToText *SpeechToTextV1) ListLanguageModels(language string) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("language=" + fmt.Sprint(language))

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

// GetListLanguageModelsResult : Cast result of ListLanguageModels operation
func GetListLanguageModelsResult(response *watson.WatsonResponse) *LanguageModels {
    result, ok := response.Result.(*LanguageModels)

    if ok {
        return result
    }

    return nil
}

// ResetLanguageModel : Reset a custom language model
func (speechToText *SpeechToTextV1) ResetLanguageModel(customizationID string) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}/reset"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", customizationID, 1)
    request := req.New().Post(creds.ServiceURL + path)

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


// TrainLanguageModel : Train a custom language model
func (speechToText *SpeechToTextV1) TrainLanguageModel(customizationID string, wordTypeToAdd string, customizationWeight float64) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}/train"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", customizationID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("word_type_to_add=" + fmt.Sprint(wordTypeToAdd))
    request.Query("customization_weight=" + fmt.Sprint(customizationWeight))

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


// UpgradeLanguageModel : Upgrade a custom language model
func (speechToText *SpeechToTextV1) UpgradeLanguageModel(customizationID string) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}/upgrade_model"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", customizationID, 1)
    request := req.New().Post(creds.ServiceURL + path)

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


// AddCorpus : Add a corpus
func (speechToText *SpeechToTextV1) AddCorpus(customizationID string, corpusName string, corpusFile os.File, allowOverwrite bool) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}/corpora/{corpus_name}"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", customizationID, 1)
    path = strings.Replace(path, "{corpus_name}", corpusName, 1)
    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "multipart/form-data")
    request.Query("allow_overwrite=" + fmt.Sprint(allowOverwrite))
    request.Type("multipart")
    request.SendFile(corpusFile)

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


// DeleteCorpus : Delete a corpus
func (speechToText *SpeechToTextV1) DeleteCorpus(customizationID string, corpusName string) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}/corpora/{corpus_name}"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", customizationID, 1)
    path = strings.Replace(path, "{corpus_name}", corpusName, 1)
    request := req.New().Delete(creds.ServiceURL + path)

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


// GetCorpus : Get a corpus
func (speechToText *SpeechToTextV1) GetCorpus(customizationID string, corpusName string) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}/corpora/{corpus_name}"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", customizationID, 1)
    path = strings.Replace(path, "{corpus_name}", corpusName, 1)
    request := req.New().Get(creds.ServiceURL + path)

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

// GetGetCorpusResult : Cast result of GetCorpus operation
func GetGetCorpusResult(response *watson.WatsonResponse) *Corpus {
    result, ok := response.Result.(*Corpus)

    if ok {
        return result
    }

    return nil
}

// ListCorpora : List corpora
func (speechToText *SpeechToTextV1) ListCorpora(customizationID string) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}/corpora"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", customizationID, 1)
    request := req.New().Get(creds.ServiceURL + path)

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

// GetListCorporaResult : Cast result of ListCorpora operation
func GetListCorporaResult(response *watson.WatsonResponse) *Corpora {
    result, ok := response.Result.(*Corpora)

    if ok {
        return result
    }

    return nil
}

// AddWord : Add a custom word
func (speechToText *SpeechToTextV1) AddWord(customizationID string, wordName string, body *CustomWord) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}/words/{word_name}"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", customizationID, 1)
    path = strings.Replace(path, "{word_name}", wordName, 1)
    request := req.New().Put(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
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


// AddWords : Add custom words
func (speechToText *SpeechToTextV1) AddWords(customizationID string, body *CustomWords) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}/words"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", customizationID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
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


// DeleteWord : Delete a custom word
func (speechToText *SpeechToTextV1) DeleteWord(customizationID string, wordName string) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}/words/{word_name}"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", customizationID, 1)
    path = strings.Replace(path, "{word_name}", wordName, 1)
    request := req.New().Delete(creds.ServiceURL + path)

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


// GetWord : Get a custom word
func (speechToText *SpeechToTextV1) GetWord(customizationID string, wordName string) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}/words/{word_name}"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", customizationID, 1)
    path = strings.Replace(path, "{word_name}", wordName, 1)
    request := req.New().Get(creds.ServiceURL + path)

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

// GetGetWordResult : Cast result of GetWord operation
func GetGetWordResult(response *watson.WatsonResponse) *Word {
    result, ok := response.Result.(*Word)

    if ok {
        return result
    }

    return nil
}

// ListWords : List custom words
func (speechToText *SpeechToTextV1) ListWords(customizationID string, wordType string, sort string) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}/words"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", customizationID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("word_type=" + fmt.Sprint(wordType))
    request.Query("sort=" + fmt.Sprint(sort))

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

// GetListWordsResult : Cast result of ListWords operation
func GetListWordsResult(response *watson.WatsonResponse) *Words {
    result, ok := response.Result.(*Words)

    if ok {
        return result
    }

    return nil
}

// CreateAcousticModel : Create a custom acoustic model
func (speechToText *SpeechToTextV1) CreateAcousticModel(body *CreateAcousticModel) (*watson.WatsonResponse, []error) {
    path := "/v1/acoustic_customizations"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
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

// GetCreateAcousticModelResult : Cast result of CreateAcousticModel operation
func GetCreateAcousticModelResult(response *watson.WatsonResponse) *AcousticModel {
    result, ok := response.Result.(*AcousticModel)

    if ok {
        return result
    }

    return nil
}

// DeleteAcousticModel : Delete a custom acoustic model
func (speechToText *SpeechToTextV1) DeleteAcousticModel(customizationID string) (*watson.WatsonResponse, []error) {
    path := "/v1/acoustic_customizations/{customization_id}"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", customizationID, 1)
    request := req.New().Delete(creds.ServiceURL + path)

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


// GetAcousticModel : Get a custom acoustic model
func (speechToText *SpeechToTextV1) GetAcousticModel(customizationID string) (*watson.WatsonResponse, []error) {
    path := "/v1/acoustic_customizations/{customization_id}"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", customizationID, 1)
    request := req.New().Get(creds.ServiceURL + path)

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

// GetGetAcousticModelResult : Cast result of GetAcousticModel operation
func GetGetAcousticModelResult(response *watson.WatsonResponse) *AcousticModel {
    result, ok := response.Result.(*AcousticModel)

    if ok {
        return result
    }

    return nil
}

// ListAcousticModels : List custom acoustic models
func (speechToText *SpeechToTextV1) ListAcousticModels(language string) (*watson.WatsonResponse, []error) {
    path := "/v1/acoustic_customizations"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("language=" + fmt.Sprint(language))

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

// GetListAcousticModelsResult : Cast result of ListAcousticModels operation
func GetListAcousticModelsResult(response *watson.WatsonResponse) *AcousticModels {
    result, ok := response.Result.(*AcousticModels)

    if ok {
        return result
    }

    return nil
}

// ResetAcousticModel : Reset a custom acoustic model
func (speechToText *SpeechToTextV1) ResetAcousticModel(customizationID string) (*watson.WatsonResponse, []error) {
    path := "/v1/acoustic_customizations/{customization_id}/reset"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", customizationID, 1)
    request := req.New().Post(creds.ServiceURL + path)

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


// TrainAcousticModel : Train a custom acoustic model
func (speechToText *SpeechToTextV1) TrainAcousticModel(customizationID string, customLanguageModelID string) (*watson.WatsonResponse, []error) {
    path := "/v1/acoustic_customizations/{customization_id}/train"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", customizationID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("custom_language_model_id=" + fmt.Sprint(customLanguageModelID))

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


// UpgradeAcousticModel : Upgrade a custom acoustic model
func (speechToText *SpeechToTextV1) UpgradeAcousticModel(customizationID string, customLanguageModelID string) (*watson.WatsonResponse, []error) {
    path := "/v1/acoustic_customizations/{customization_id}/upgrade_model"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", customizationID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("custom_language_model_id=" + fmt.Sprint(customLanguageModelID))

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


// AddAudio : Add an audio resource
func (speechToText *SpeechToTextV1) AddAudio(customizationID string, audioName string, body *[][]byte, contentType string, containedContentType string, allowOverwrite bool) (*watson.WatsonResponse, []error) {
    path := "/v1/acoustic_customizations/{customization_id}/audio/{audio_name}"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", customizationID, 1)
    path = strings.Replace(path, "{audio_name}", audioName, 1)
    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/zip")
    request.Set("Content-Type", fmt.Sprint(contentType))
    request.Set("Contained-Content-Type", fmt.Sprint(containedContentType))
    request.Query("allow_overwrite=" + fmt.Sprint(allowOverwrite))
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


// DeleteAudio : Delete an audio resource
func (speechToText *SpeechToTextV1) DeleteAudio(customizationID string, audioName string) (*watson.WatsonResponse, []error) {
    path := "/v1/acoustic_customizations/{customization_id}/audio/{audio_name}"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", customizationID, 1)
    path = strings.Replace(path, "{audio_name}", audioName, 1)
    request := req.New().Delete(creds.ServiceURL + path)

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


// GetAudio : Get an audio resource
func (speechToText *SpeechToTextV1) GetAudio(customizationID string, audioName string) (*watson.WatsonResponse, []error) {
    path := "/v1/acoustic_customizations/{customization_id}/audio/{audio_name}"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", customizationID, 1)
    path = strings.Replace(path, "{audio_name}", audioName, 1)
    request := req.New().Get(creds.ServiceURL + path)

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

// GetGetAudioResult : Cast result of GetAudio operation
func GetGetAudioResult(response *watson.WatsonResponse) *AudioListing {
    result, ok := response.Result.(*AudioListing)

    if ok {
        return result
    }

    return nil
}

// ListAudio : List audio resources
func (speechToText *SpeechToTextV1) ListAudio(customizationID string) (*watson.WatsonResponse, []error) {
    path := "/v1/acoustic_customizations/{customization_id}/audio"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    path = strings.Replace(path, "{customization_id}", customizationID, 1)
    request := req.New().Get(creds.ServiceURL + path)

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

// GetListAudioResult : Cast result of ListAudio operation
func GetListAudioResult(response *watson.WatsonResponse) *AudioResources {
    result, ok := response.Result.(*AudioResources)

    if ok {
        return result
    }

    return nil
}

// DeleteUserData : Delete labeled data
func (speechToText *SpeechToTextV1) DeleteUserData(customerID string) (*watson.WatsonResponse, []error) {
    path := "/v1/user_data"
    creds := speechToText.client.Creds
    useTM := speechToText.client.UseTM
    tokenManager := speechToText.client.TokenManager

    request := req.New().Delete(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("customer_id=" + fmt.Sprint(customerID))

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

// AudioDetails : AudioDetails struct
type AudioDetails struct {

	// The type of the audio resource: * `audio` for an individual audio file * `archive` for an archive (**.zip** or **.tar.gz**) file that contains audio files * `undetermined` for a resource that the service cannot validate (for example, if the user mistakenly passes a file that does not contain audio, such as a JPEG file).
	TypeVar string `json:"type_var,omitempty"`

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

// CreateAcousticModel : CreateAcousticModel struct
type CreateAcousticModel struct {

	// A user-defined name for the new custom acoustic model. Use a name that is unique among all custom acoustic models that you own. Use a localized name that matches the language of the custom model. Use a name that describes the acoustic environment of the custom model, such as `Mobile custom model` or `Noisy car custom model`.
	Name string `json:"name"`

	// The name of the base language model that is to be customized by the new custom acoustic model. The new custom model can be used only with the base model that it customizes. To determine whether a base model supports acoustic model customization, refer to [Language support for customization](https://console.bluemix.net/docs/services/speech-to-text/custom.html#languageSupport).
	BaseModelName string `json:"base_model_name"`

	// A description of the new custom acoustic model. Use a localized description that matches the language of the custom model.
	Description string `json:"description,omitempty"`
}

// CreateLanguageModel : CreateLanguageModel struct
type CreateLanguageModel struct {

	// A user-defined name for the new custom language model. Use a name that is unique among all custom language models that you own. Use a localized name that matches the language of the custom model. Use a name that describes the domain of the custom model, such as `Medical custom model` or `Legal custom model`.
	Name string `json:"name"`

	// The name of the base language model that is to be customized by the new custom language model. The new custom model can be used only with the base model that it customizes. To determine whether a base model supports language model customization, request information about the base model and check that the attribute `custom_language_model` is set to `true`, or refer to [Language support for customization](https://console.bluemix.net/docs/services/speech-to-text/custom.html#languageSupport).
	BaseModelName string `json:"base_model_name"`

	// The dialect of the specified language that is to be used with the custom language model. The parameter is meaningful only for Spanish models, for which the service creates a custom language model that is suited for speech in one of the following dialects: * `es-ES` for Castilian Spanish (the default) * `es-LA` for Latin American Spanish * `es-US` for North American (Mexican) Spanish A specified dialect must be valid for the base model. By default, the dialect matches the language of the base model; for example, `en-US` for either of the US English language models.
	Dialect string `json:"dialect,omitempty"`

	// A description of the new custom language model. Use a localized description that matches the language of the custom model.
	Description string `json:"description,omitempty"`
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

// CustomWords : CustomWords struct
type CustomWords struct {

	// An array of objects that provides information about each custom word that is to be added to or updated in the custom language model.
	Words []CustomWord `json:"words"`
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

// RegisterStatus : RegisterStatus struct
type RegisterStatus struct {

	// The current status of the job: * `created` if the callback URL was successfully white-listed as a result of the call. * `already created` if the URL was already white-listed.
	Status string `json:"status"`

	// The callback URL that is successfully registered.
	URL string `json:"url"`
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
	FinalResults bool `json:"final_results"`
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
	FinalResults bool `json:"final_results"`

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
