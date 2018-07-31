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
    "bytes"
    "fmt"
    "os"
    "strings"
    req "github.com/parnurzeal/gorequest"
    watson "golang-sdk"
)

// TextToSpeechV1 : The TextToSpeechV1 service
type TextToSpeechV1 struct {
	client *watson.Client
}

// NewTextToSpeechV1 : Instantiate TextToSpeechV1
func NewTextToSpeechV1(creds watson.Credentials) (*TextToSpeechV1, error) {
    if creds.ServiceURL == "" {
        creds.ServiceURL = "https://stream.watsonplatform.net/text-to-speech/api"
    }

	client, clientErr := watson.NewClient(creds, "text_to_speech")

	if clientErr != nil {
		return nil, clientErr
	}

	return &TextToSpeechV1{ client: client }, nil
}

// GetVoice : Get a voice
func (textToSpeech *TextToSpeechV1) GetVoice(options *GetVoiceOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/voices/{voice}"
    creds := textToSpeech.client.Creds
    useTM := textToSpeech.client.UseTM
    tokenManager := textToSpeech.client.TokenManager

    path = strings.Replace(path, "{voice}", options.Voice, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    if options.IsCustomizationIDSet {
        request.Query("customization_id=" + fmt.Sprint(options.CustomizationID))
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

    response.Result = new(Voice)
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

// GetGetVoiceResult : Cast result of GetVoice operation
func GetGetVoiceResult(response *watson.WatsonResponse) *Voice {
    result, ok := response.Result.(*Voice)

    if ok {
        return result
    }

    return nil
}

// ListVoices : List voices
func (textToSpeech *TextToSpeechV1) ListVoices(options *ListVoicesOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/voices"
    creds := textToSpeech.client.Creds
    useTM := textToSpeech.client.UseTM
    tokenManager := textToSpeech.client.TokenManager

    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")

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

    response.Result = new(Voices)
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

// GetListVoicesResult : Cast result of ListVoices operation
func GetListVoicesResult(response *watson.WatsonResponse) *Voices {
    result, ok := response.Result.(*Voices)

    if ok {
        return result
    }

    return nil
}

// Synthesize : Synthesize audio
func (textToSpeech *TextToSpeechV1) Synthesize(options *SynthesizeOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/synthesize"
    creds := textToSpeech.client.Creds
    useTM := textToSpeech.client.UseTM
    tokenManager := textToSpeech.client.TokenManager

    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "audio/basic")
    request.Set("Content-Type", "application/json")
    if options.IsAcceptSet {
        request.Set("Accept", fmt.Sprint(options.Accept))
    }
    if options.IsVoiceSet {
        request.Query("voice=" + fmt.Sprint(options.Voice))
    }
    if options.IsCustomizationIDSet {
        request.Query("customization_id=" + fmt.Sprint(options.CustomizationID))
    }
    body := map[string]interface{}{}
    body["text"] = options.Text
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

    response.Result = new(os.File)
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

// GetSynthesizeResult : Cast result of Synthesize operation
func GetSynthesizeResult(response *watson.WatsonResponse) *os.File {
    result, ok := response.Result.(*os.File)

    if ok {
        return result
    }

    return nil
}

// GetPronunciation : Get pronunciation
func (textToSpeech *TextToSpeechV1) GetPronunciation(options *GetPronunciationOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/pronunciation"
    creds := textToSpeech.client.Creds
    useTM := textToSpeech.client.UseTM
    tokenManager := textToSpeech.client.TokenManager

    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Query("text=" + fmt.Sprint(options.Text))
    if options.IsVoiceSet {
        request.Query("voice=" + fmt.Sprint(options.Voice))
    }
    if options.IsFormatSet {
        request.Query("format=" + fmt.Sprint(options.Format))
    }
    if options.IsCustomizationIDSet {
        request.Query("customization_id=" + fmt.Sprint(options.CustomizationID))
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

    response.Result = new(Pronunciation)
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

// GetGetPronunciationResult : Cast result of GetPronunciation operation
func GetGetPronunciationResult(response *watson.WatsonResponse) *Pronunciation {
    result, ok := response.Result.(*Pronunciation)

    if ok {
        return result
    }

    return nil
}

// CreateVoiceModel : Create a custom model
func (textToSpeech *TextToSpeechV1) CreateVoiceModel(options *CreateVoiceModelOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations"
    creds := textToSpeech.client.Creds
    useTM := textToSpeech.client.UseTM
    tokenManager := textToSpeech.client.TokenManager

    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    body := map[string]interface{}{}
    body["name"] = options.Name
    if options.IsLanguageSet {
        body["language"] = options.Language
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

    response.Result = new(VoiceModel)
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

// GetCreateVoiceModelResult : Cast result of CreateVoiceModel operation
func GetCreateVoiceModelResult(response *watson.WatsonResponse) *VoiceModel {
    result, ok := response.Result.(*VoiceModel)

    if ok {
        return result
    }

    return nil
}

// DeleteVoiceModel : Delete a custom model
func (textToSpeech *TextToSpeechV1) DeleteVoiceModel(options *DeleteVoiceModelOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}"
    creds := textToSpeech.client.Creds
    useTM := textToSpeech.client.UseTM
    tokenManager := textToSpeech.client.TokenManager

    path = strings.Replace(path, "{customization_id}", options.CustomizationID, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
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


// GetVoiceModel : Get a custom model
func (textToSpeech *TextToSpeechV1) GetVoiceModel(options *GetVoiceModelOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}"
    creds := textToSpeech.client.Creds
    useTM := textToSpeech.client.UseTM
    tokenManager := textToSpeech.client.TokenManager

    path = strings.Replace(path, "{customization_id}", options.CustomizationID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")

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

    response.Result = new(VoiceModel)
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

// GetGetVoiceModelResult : Cast result of GetVoiceModel operation
func GetGetVoiceModelResult(response *watson.WatsonResponse) *VoiceModel {
    result, ok := response.Result.(*VoiceModel)

    if ok {
        return result
    }

    return nil
}

// ListVoiceModels : List custom models
func (textToSpeech *TextToSpeechV1) ListVoiceModels(options *ListVoiceModelsOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations"
    creds := textToSpeech.client.Creds
    useTM := textToSpeech.client.UseTM
    tokenManager := textToSpeech.client.TokenManager

    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
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

    response.Result = new(VoiceModels)
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

// GetListVoiceModelsResult : Cast result of ListVoiceModels operation
func GetListVoiceModelsResult(response *watson.WatsonResponse) *VoiceModels {
    result, ok := response.Result.(*VoiceModels)

    if ok {
        return result
    }

    return nil
}

// UpdateVoiceModel : Update a custom model
func (textToSpeech *TextToSpeechV1) UpdateVoiceModel(options *UpdateVoiceModelOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}"
    creds := textToSpeech.client.Creds
    useTM := textToSpeech.client.UseTM
    tokenManager := textToSpeech.client.TokenManager

    path = strings.Replace(path, "{customization_id}", options.CustomizationID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    body := map[string]interface{}{}
    if options.IsNameSet {
        body["name"] = options.Name
    }
    if options.IsDescriptionSet {
        body["description"] = options.Description
    }
    if options.IsWordsSet {
        body["words"] = options.Words
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


// AddWord : Add a custom word
func (textToSpeech *TextToSpeechV1) AddWord(options *AddWordOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}/words/{word}"
    creds := textToSpeech.client.Creds
    useTM := textToSpeech.client.UseTM
    tokenManager := textToSpeech.client.TokenManager

    path = strings.Replace(path, "{customization_id}", options.CustomizationID, 1)
    path = strings.Replace(path, "{word}", options.Word, 1)
    request := req.New().Put(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Content-Type", "application/json")
    body := map[string]interface{}{}
    if options.IsTranslationSet {
        body["translation"] = options.Translation
    }
    if options.IsPartOfSpeechSet {
        body["part_of_speech"] = options.PartOfSpeech
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
func (textToSpeech *TextToSpeechV1) AddWords(options *AddWordsOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}/words"
    creds := textToSpeech.client.Creds
    useTM := textToSpeech.client.UseTM
    tokenManager := textToSpeech.client.TokenManager

    path = strings.Replace(path, "{customization_id}", options.CustomizationID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    body := map[string]interface{}{}
    if options.IsWordsSet {
        body["words"] = options.Words
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
func (textToSpeech *TextToSpeechV1) DeleteWord(options *DeleteWordOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}/words/{word}"
    creds := textToSpeech.client.Creds
    useTM := textToSpeech.client.UseTM
    tokenManager := textToSpeech.client.TokenManager

    path = strings.Replace(path, "{customization_id}", options.CustomizationID, 1)
    path = strings.Replace(path, "{word}", options.Word, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
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
func (textToSpeech *TextToSpeechV1) GetWord(options *GetWordOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}/words/{word}"
    creds := textToSpeech.client.Creds
    useTM := textToSpeech.client.UseTM
    tokenManager := textToSpeech.client.TokenManager

    path = strings.Replace(path, "{customization_id}", options.CustomizationID, 1)
    path = strings.Replace(path, "{word}", options.Word, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")

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

    response.Result = new(Translation)
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
func GetGetWordResult(response *watson.WatsonResponse) *Translation {
    result, ok := response.Result.(*Translation)

    if ok {
        return result
    }

    return nil
}

// ListWords : List custom words
func (textToSpeech *TextToSpeechV1) ListWords(options *ListWordsOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}/words"
    creds := textToSpeech.client.Creds
    useTM := textToSpeech.client.UseTM
    tokenManager := textToSpeech.client.TokenManager

    path = strings.Replace(path, "{customization_id}", options.CustomizationID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")

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

// DeleteUserData : Delete labeled data
func (textToSpeech *TextToSpeechV1) DeleteUserData(options *DeleteUserDataOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/user_data"
    creds := textToSpeech.client.Creds
    useTM := textToSpeech.client.UseTM
    tokenManager := textToSpeech.client.TokenManager

    request := req.New().Delete(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

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



// AddWordOptions : The addWord options.
type AddWordOptions struct {

	// The customization ID (GUID) of the custom voice model. You must make the request with service credentials created for the instance of the service that owns the custom model.
	CustomizationID string `json:"customization_id"`

	// The word that is to be added or updated for the custom voice model.
	Word string `json:"word"`

	// The phonetic or sounds-like translation for the word. A phonetic translation is based on the SSML format for representing the phonetic string of a word either as an IPA translation or as an IBM SPR translation. A sounds-like is one or more words that, when combined, sound like the word.
	Translation string `json:"translation,omitempty"`

    // Indicates whether user set optional parameter Translation
    IsTranslationSet bool

	// **Japanese only.** The part of speech for the word. The service uses the value to produce the correct intonation for the word. You can create only a single entry, with or without a single part of speech, for any word; you cannot create multiple entries with different parts of speech for the same word. For more information, see [Working with Japanese entries](https://console.bluemix.net/docs/services/text-to-speech/custom-rules.html#jaNotes).
	PartOfSpeech string `json:"part_of_speech,omitempty"`

    // Indicates whether user set optional parameter PartOfSpeech
    IsPartOfSpeechSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewAddWordOptions : Instantiate AddWordOptions
func NewAddWordOptions(customizationID string, word string) *AddWordOptions {
    return &AddWordOptions{
        CustomizationID: customizationID,
        Word: word,
    }
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *AddWordOptions) SetCustomizationID(param string) *AddWordOptions {
    options.CustomizationID = param
    return options
}

// SetWord : Allow user to set Word
func (options *AddWordOptions) SetWord(param string) *AddWordOptions {
    options.Word = param
    return options
}

// SetTranslation : Allow user to set Translation
func (options *AddWordOptions) SetTranslation(param string) *AddWordOptions {
    options.Translation = param
    options.IsTranslationSet = true
    return options
}

// SetPartOfSpeech : Allow user to set PartOfSpeech
func (options *AddWordOptions) SetPartOfSpeech(param string) *AddWordOptions {
    options.PartOfSpeech = param
    options.IsPartOfSpeechSet = true
    return options
}

// SetHeaders : Allow user to set Headers
func (options *AddWordOptions) SetHeaders(param map[string]string) *AddWordOptions {
    options.Headers = param
    return options
}

// AddWordsOptions : The addWords options.
type AddWordsOptions struct {

	// The customization ID (GUID) of the custom voice model. You must make the request with service credentials created for the instance of the service that owns the custom model.
	CustomizationID string `json:"customization_id"`

	// The **Add custom words** method accepts an array of `Word` objects. Each object provides a word that is to be added or updated for the custom voice model and the word's translation. The **List custom words** method returns an array of `Word` objects. Each object shows a word and its translation from the custom voice model. The words are listed in alphabetical order, with uppercase letters listed before lowercase letters. The array is empty if the custom model contains no words.
	Words []Word `json:"words,omitempty"`

    // Indicates whether user set optional parameter Words
    IsWordsSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewAddWordsOptions : Instantiate AddWordsOptions
func NewAddWordsOptions(customizationID string) *AddWordsOptions {
    return &AddWordsOptions{
        CustomizationID: customizationID,
    }
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *AddWordsOptions) SetCustomizationID(param string) *AddWordsOptions {
    options.CustomizationID = param
    return options
}

// SetWords : Allow user to set Words
func (options *AddWordsOptions) SetWords(param []Word) *AddWordsOptions {
    options.Words = param
    options.IsWordsSet = true
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
	Name string `json:"name"`

	// The language of the new custom voice model. Omit the parameter to use the the default language, `en-US`.
	Language string `json:"language,omitempty"`

    // Indicates whether user set optional parameter Language
    IsLanguageSet bool

	// A description of the new custom voice model. Specifying a description is recommended.
	Description string `json:"description,omitempty"`

    // Indicates whether user set optional parameter Description
    IsDescriptionSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewCreateVoiceModelOptions : Instantiate CreateVoiceModelOptions
func NewCreateVoiceModelOptions(name string) *CreateVoiceModelOptions {
    return &CreateVoiceModelOptions{
        Name: name,
    }
}

// SetName : Allow user to set Name
func (options *CreateVoiceModelOptions) SetName(param string) *CreateVoiceModelOptions {
    options.Name = param
    return options
}

// SetLanguage : Allow user to set Language
func (options *CreateVoiceModelOptions) SetLanguage(param string) *CreateVoiceModelOptions {
    options.Language = param
    options.IsLanguageSet = true
    return options
}

// SetDescription : Allow user to set Description
func (options *CreateVoiceModelOptions) SetDescription(param string) *CreateVoiceModelOptions {
    options.Description = param
    options.IsDescriptionSet = true
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

// DeleteVoiceModelOptions : The deleteVoiceModel options.
type DeleteVoiceModelOptions struct {

	// The customization ID (GUID) of the custom voice model. You must make the request with service credentials created for the instance of the service that owns the custom model.
	CustomizationID string `json:"customization_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewDeleteVoiceModelOptions : Instantiate DeleteVoiceModelOptions
func NewDeleteVoiceModelOptions(customizationID string) *DeleteVoiceModelOptions {
    return &DeleteVoiceModelOptions{
        CustomizationID: customizationID,
    }
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *DeleteVoiceModelOptions) SetCustomizationID(param string) *DeleteVoiceModelOptions {
    options.CustomizationID = param
    return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteVoiceModelOptions) SetHeaders(param map[string]string) *DeleteVoiceModelOptions {
    options.Headers = param
    return options
}

// DeleteWordOptions : The deleteWord options.
type DeleteWordOptions struct {

	// The customization ID (GUID) of the custom voice model. You must make the request with service credentials created for the instance of the service that owns the custom model.
	CustomizationID string `json:"customization_id"`

	// The word that is to be deleted from the custom voice model.
	Word string `json:"word"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewDeleteWordOptions : Instantiate DeleteWordOptions
func NewDeleteWordOptions(customizationID string, word string) *DeleteWordOptions {
    return &DeleteWordOptions{
        CustomizationID: customizationID,
        Word: word,
    }
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *DeleteWordOptions) SetCustomizationID(param string) *DeleteWordOptions {
    options.CustomizationID = param
    return options
}

// SetWord : Allow user to set Word
func (options *DeleteWordOptions) SetWord(param string) *DeleteWordOptions {
    options.Word = param
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
	Text string `json:"text"`

	// A voice that specifies the language in which the pronunciation is to be returned. All voices for the same language (for example, `en-US`) return the same translation.
	Voice string `json:"voice,omitempty"`

    // Indicates whether user set optional parameter Voice
    IsVoiceSet bool

	// The phoneme format in which to return the pronunciation. Omit the parameter to obtain the pronunciation in the default format.
	Format string `json:"format,omitempty"`

    // Indicates whether user set optional parameter Format
    IsFormatSet bool

	// The customization ID (GUID) of a custom voice model for which the pronunciation is to be returned. The language of a specified custom model must match the language of the specified voice. If the word is not defined in the specified custom model, the service returns the default translation for the custom model's language. You must make the request with service credentials created for the instance of the service that owns the custom model. Omit the parameter to see the translation for the specified voice with no customization.
	CustomizationID string `json:"customization_id,omitempty"`

    // Indicates whether user set optional parameter CustomizationID
    IsCustomizationIDSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewGetPronunciationOptions : Instantiate GetPronunciationOptions
func NewGetPronunciationOptions(text string) *GetPronunciationOptions {
    return &GetPronunciationOptions{
        Text: text,
    }
}

// SetText : Allow user to set Text
func (options *GetPronunciationOptions) SetText(param string) *GetPronunciationOptions {
    options.Text = param
    return options
}

// SetVoice : Allow user to set Voice
func (options *GetPronunciationOptions) SetVoice(param string) *GetPronunciationOptions {
    options.Voice = param
    options.IsVoiceSet = true
    return options
}

// SetFormat : Allow user to set Format
func (options *GetPronunciationOptions) SetFormat(param string) *GetPronunciationOptions {
    options.Format = param
    options.IsFormatSet = true
    return options
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *GetPronunciationOptions) SetCustomizationID(param string) *GetPronunciationOptions {
    options.CustomizationID = param
    options.IsCustomizationIDSet = true
    return options
}

// SetHeaders : Allow user to set Headers
func (options *GetPronunciationOptions) SetHeaders(param map[string]string) *GetPronunciationOptions {
    options.Headers = param
    return options
}

// GetVoiceModelOptions : The getVoiceModel options.
type GetVoiceModelOptions struct {

	// The customization ID (GUID) of the custom voice model. You must make the request with service credentials created for the instance of the service that owns the custom model.
	CustomizationID string `json:"customization_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewGetVoiceModelOptions : Instantiate GetVoiceModelOptions
func NewGetVoiceModelOptions(customizationID string) *GetVoiceModelOptions {
    return &GetVoiceModelOptions{
        CustomizationID: customizationID,
    }
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *GetVoiceModelOptions) SetCustomizationID(param string) *GetVoiceModelOptions {
    options.CustomizationID = param
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
	Voice string `json:"voice"`

	// The customization ID (GUID) of a custom voice model for which information is to be returned. You must make the request with service credentials created for the instance of the service that owns the custom model. Omit the parameter to see information about the specified voice with no customization.
	CustomizationID string `json:"customization_id,omitempty"`

    // Indicates whether user set optional parameter CustomizationID
    IsCustomizationIDSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewGetVoiceOptions : Instantiate GetVoiceOptions
func NewGetVoiceOptions(voice string) *GetVoiceOptions {
    return &GetVoiceOptions{
        Voice: voice,
    }
}

// SetVoice : Allow user to set Voice
func (options *GetVoiceOptions) SetVoice(param string) *GetVoiceOptions {
    options.Voice = param
    return options
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *GetVoiceOptions) SetCustomizationID(param string) *GetVoiceOptions {
    options.CustomizationID = param
    options.IsCustomizationIDSet = true
    return options
}

// SetHeaders : Allow user to set Headers
func (options *GetVoiceOptions) SetHeaders(param map[string]string) *GetVoiceOptions {
    options.Headers = param
    return options
}

// GetWordOptions : The getWord options.
type GetWordOptions struct {

	// The customization ID (GUID) of the custom voice model. You must make the request with service credentials created for the instance of the service that owns the custom model.
	CustomizationID string `json:"customization_id"`

	// The word that is to be queried from the custom voice model.
	Word string `json:"word"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewGetWordOptions : Instantiate GetWordOptions
func NewGetWordOptions(customizationID string, word string) *GetWordOptions {
    return &GetWordOptions{
        CustomizationID: customizationID,
        Word: word,
    }
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *GetWordOptions) SetCustomizationID(param string) *GetWordOptions {
    options.CustomizationID = param
    return options
}

// SetWord : Allow user to set Word
func (options *GetWordOptions) SetWord(param string) *GetWordOptions {
    options.Word = param
    return options
}

// SetHeaders : Allow user to set Headers
func (options *GetWordOptions) SetHeaders(param map[string]string) *GetWordOptions {
    options.Headers = param
    return options
}

// ListVoiceModelsOptions : The listVoiceModels options.
type ListVoiceModelsOptions struct {

	// The language for which custom voice models that are owned by the requesting service credentials are to be returned. Omit the parameter to see all custom voice models that are owned by the requester.
	Language string `json:"language,omitempty"`

    // Indicates whether user set optional parameter Language
    IsLanguageSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewListVoiceModelsOptions : Instantiate ListVoiceModelsOptions
func NewListVoiceModelsOptions() *ListVoiceModelsOptions {
    return &ListVoiceModelsOptions{}
}

// SetLanguage : Allow user to set Language
func (options *ListVoiceModelsOptions) SetLanguage(param string) *ListVoiceModelsOptions {
    options.Language = param
    options.IsLanguageSet = true
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
func NewListVoicesOptions() *ListVoicesOptions {
    return &ListVoicesOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListVoicesOptions) SetHeaders(param map[string]string) *ListVoicesOptions {
    options.Headers = param
    return options
}

// ListWordsOptions : The listWords options.
type ListWordsOptions struct {

	// The customization ID (GUID) of the custom voice model. You must make the request with service credentials created for the instance of the service that owns the custom model.
	CustomizationID string `json:"customization_id"`

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

// SetHeaders : Allow user to set Headers
func (options *ListWordsOptions) SetHeaders(param map[string]string) *ListWordsOptions {
    options.Headers = param
    return options
}

// Pronunciation : Pronunciation struct
type Pronunciation struct {

	// The pronunciation of the specified text in the requested voice and format. If a custom voice model is specified, the pronunciation also reflects that custom voice.
	Pronunciation string `json:"pronunciation"`
}

// SupportedFeatures : SupportedFeatures struct
type SupportedFeatures struct {

	// If `true`, the voice can be customized; if `false`, the voice cannot be customized. (Same as `customizable`.).
	CustomPronunciation bool `json:"custom_pronunciation"`

	// If `true`, the voice can be transformed by using the SSML &lt;voice-transformation&gt; element; if `false`, the voice cannot be transformed.
	VoiceTransformation bool `json:"voice_transformation"`
}

// SynthesizeOptions : The synthesize options.
type SynthesizeOptions struct {

	// The text to synthesize.
	Text string `json:"text"`

	// The requested audio format (MIME type) of the audio. You can use the `Accept` header or the `accept` query parameter to specify the audio format. (For the `audio/l16` format, you can optionally specify `endianness=big-endian` or `endianness=little-endian`; the default is little endian.) For detailed information about the supported audio formats and sampling rates, see [Specifying an audio format](https://console.bluemix.net/docs/services/text-to-speech/http.html#format).
	Accept string `json:"Accept,omitempty"`

    // Indicates whether user set optional parameter Accept
    IsAcceptSet bool

	// The voice to use for synthesis.
	Voice string `json:"voice,omitempty"`

    // Indicates whether user set optional parameter Voice
    IsVoiceSet bool

	// The customization ID (GUID) of a custom voice model to use for the synthesis. If a custom voice model is specified, it is guaranteed to work only if it matches the language of the indicated voice. You must make the request with service credentials created for the instance of the service that owns the custom model. Omit the parameter to use the specified voice with no customization.
	CustomizationID string `json:"customization_id,omitempty"`

    // Indicates whether user set optional parameter CustomizationID
    IsCustomizationIDSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewSynthesizeOptions : Instantiate SynthesizeOptions
func NewSynthesizeOptions(text string) *SynthesizeOptions {
    return &SynthesizeOptions{
        Text: text,
    }
}

// SetText : Allow user to set Text
func (options *SynthesizeOptions) SetText(param string) *SynthesizeOptions {
    options.Text = param
    return options
}

// SetAccept : Allow user to set Accept
func (options *SynthesizeOptions) SetAccept(param string) *SynthesizeOptions {
    options.Accept = param
    options.IsAcceptSet = true
    return options
}

// SetVoice : Allow user to set Voice
func (options *SynthesizeOptions) SetVoice(param string) *SynthesizeOptions {
    options.Voice = param
    options.IsVoiceSet = true
    return options
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *SynthesizeOptions) SetCustomizationID(param string) *SynthesizeOptions {
    options.CustomizationID = param
    options.IsCustomizationIDSet = true
    return options
}

// SetHeaders : Allow user to set Headers
func (options *SynthesizeOptions) SetHeaders(param map[string]string) *SynthesizeOptions {
    options.Headers = param
    return options
}

// Translation : Translation struct
type Translation struct {

	// The phonetic or sounds-like translation for the word. A phonetic translation is based on the SSML format for representing the phonetic string of a word either as an IPA translation or as an IBM SPR translation. A sounds-like is one or more words that, when combined, sound like the word.
	Translation string `json:"translation,omitempty"`

	// **Japanese only.** The part of speech for the word. The service uses the value to produce the correct intonation for the word. You can create only a single entry, with or without a single part of speech, for any word; you cannot create multiple entries with different parts of speech for the same word. For more information, see [Working with Japanese entries](https://console.bluemix.net/docs/services/text-to-speech/custom-rules.html#jaNotes).
	PartOfSpeech string `json:"part_of_speech,omitempty"`
}

// UpdateVoiceModelOptions : The updateVoiceModel options.
type UpdateVoiceModelOptions struct {

	// The customization ID (GUID) of the custom voice model. You must make the request with service credentials created for the instance of the service that owns the custom model.
	CustomizationID string `json:"customization_id"`

	// A new name for the custom voice model.
	Name string `json:"name,omitempty"`

    // Indicates whether user set optional parameter Name
    IsNameSet bool

	// A new description for the custom voice model.
	Description string `json:"description,omitempty"`

    // Indicates whether user set optional parameter Description
    IsDescriptionSet bool

	// An array of `Word` objects that provides the words and their translations that are to be added or updated for the custom voice model. Pass an empty array to make no additions or updates.
	Words []Word `json:"words,omitempty"`

    // Indicates whether user set optional parameter Words
    IsWordsSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewUpdateVoiceModelOptions : Instantiate UpdateVoiceModelOptions
func NewUpdateVoiceModelOptions(customizationID string) *UpdateVoiceModelOptions {
    return &UpdateVoiceModelOptions{
        CustomizationID: customizationID,
    }
}

// SetCustomizationID : Allow user to set CustomizationID
func (options *UpdateVoiceModelOptions) SetCustomizationID(param string) *UpdateVoiceModelOptions {
    options.CustomizationID = param
    return options
}

// SetName : Allow user to set Name
func (options *UpdateVoiceModelOptions) SetName(param string) *UpdateVoiceModelOptions {
    options.Name = param
    options.IsNameSet = true
    return options
}

// SetDescription : Allow user to set Description
func (options *UpdateVoiceModelOptions) SetDescription(param string) *UpdateVoiceModelOptions {
    options.Description = param
    options.IsDescriptionSet = true
    return options
}

// SetWords : Allow user to set Words
func (options *UpdateVoiceModelOptions) SetWords(param []Word) *UpdateVoiceModelOptions {
    options.Words = param
    options.IsWordsSet = true
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
	URL string `json:"url"`

	// The gender of the voice: `male` or `female`.
	Gender string `json:"gender"`

	// The name of the voice. Use this as the voice identifier in all requests.
	Name string `json:"name"`

	// The language and region of the voice (for example, `en-US`).
	Language string `json:"language"`

	// A textual description of the voice.
	Description string `json:"description"`

	// If `true`, the voice can be customized; if `false`, the voice cannot be customized. (Same as `custom_pronunciation`; maintained for backward compatibility.).
	Customizable bool `json:"customizable"`

	// Describes the additional service features supported with the voice.
	SupportedFeatures SupportedFeatures `json:"supported_features"`

	// Returns information about a specified custom voice model. This field is returned only by the **Get a voice** method and only when you specify the customization ID of a custom voice model.
	Customization VoiceModel `json:"customization,omitempty"`
}

// VoiceModel : VoiceModel struct
type VoiceModel struct {

	// The customization ID (GUID) of the custom voice model. The **Create a custom model** method returns only this field. It does not not return the other fields of this object.
	CustomizationID string `json:"customization_id"`

	// The name of the custom voice model.
	Name string `json:"name,omitempty"`

	// The language identifier of the custom voice model (for example, `en-US`).
	Language string `json:"language,omitempty"`

	// The GUID of the service credentials for the instance of the service that owns the custom voice model.
	Owner string `json:"owner,omitempty"`

	// The date and time in Coordinated Universal Time (UTC) at which the custom voice model was created. The value is provided in full ISO 8601 format (`YYYY-MM-DDThh:mm:ss.sTZD`).
	Created string `json:"created,omitempty"`

	// The date and time in Coordinated Universal Time (UTC) at which the custom voice model was last modified. Equals `created` when a new voice model is first added but has yet to be updated. The value is provided in full ISO 8601 format (`YYYY-MM-DDThh:mm:ss.sTZD`).
	LastModified string `json:"last_modified,omitempty"`

	// The description of the custom voice model.
	Description string `json:"description,omitempty"`

	// An array of `Word` objects that lists the words and their translations from the custom voice model. The words are listed in alphabetical order, with uppercase letters listed before lowercase letters. The array is empty if the custom model contains no words. This field is returned only by the **Get a voice** method and only when you specify the customization ID of a custom voice model.
	Words []Word `json:"words,omitempty"`
}

// VoiceModels : VoiceModels struct
type VoiceModels struct {

	// An array of `VoiceModel` objects that provides information about each available custom voice model. The array is empty if the requesting service credentials own no custom voice models (if no language is specified) or own no custom voice models for the specified language.
	Customizations []VoiceModel `json:"customizations"`
}

// Voices : Voices struct
type Voices struct {

	// A list of available voices.
	Voices []Voice `json:"voices"`
}

// Word : Word struct
type Word struct {

	// A word from the custom voice model.
	Word string `json:"word"`

	// The phonetic or sounds-like translation for the word. A phonetic translation is based on the SSML format for representing the phonetic string of a word either as an IPA or IBM SPR translation. A sounds-like translation consists of one or more words that, when combined, sound like the word.
	Translation string `json:"translation"`

	// **Japanese only.** The part of speech for the word. The service uses the value to produce the correct intonation for the word. You can create only a single entry, with or without a single part of speech, for any word; you cannot create multiple entries with different parts of speech for the same word. For more information, see [Working with Japanese entries](https://console.bluemix.net/docs/services/text-to-speech/custom-rules.html#jaNotes).
	PartOfSpeech string `json:"part_of_speech,omitempty"`
}

// Words : Words struct
type Words struct {

	// The **Add custom words** method accepts an array of `Word` objects. Each object provides a word that is to be added or updated for the custom voice model and the word's translation. The **List custom words** method returns an array of `Word` objects. Each object shows a word and its translation from the custom voice model. The words are listed in alphabetical order, with uppercase letters listed before lowercase letters. The array is empty if the custom model contains no words.
	Words []Word `json:"words,omitempty"`
}
