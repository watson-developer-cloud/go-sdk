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
package textToSpeechV1

import (
    "bytes"
    "fmt"
    "os"
    "strings"
    req "github.com/parnurzeal/gorequest"
    watson "golang-sdk"
)

type TextToSpeechV1 struct {
	client *watson.Client
}

func NewTextToSpeechV1(creds watson.Credentials) (*TextToSpeechV1, error) {
	client, clientErr := watson.NewClient(creds, "text_to_speech")

	if clientErr != nil {
		return nil, clientErr
	}

	return &TextToSpeechV1{ client: client }, nil
}

// Get a voice
func (textToSpeech *TextToSpeechV1) GetVoice(voice string, customizationID string) (*watson.WatsonResponse, []error) {
    path := "/v1/voices/{voice}"
    creds := textToSpeech.client.Creds
    useTM := textToSpeech.client.UseTM
    tokenManager := textToSpeech.client.TokenManager

    path = strings.Replace(path, "{voice}", voice, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("customization_id=" + fmt.Sprint(customizationID))

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

func GetGetVoiceResult(response *watson.WatsonResponse) *Voice {
    result, ok := response.Result.(*Voice)

    if ok {
        return result
    }

    return nil
}

// List voices
func (textToSpeech *TextToSpeechV1) ListVoices() (*watson.WatsonResponse, []error) {
    path := "/v1/voices"
    creds := textToSpeech.client.Creds
    useTM := textToSpeech.client.UseTM
    tokenManager := textToSpeech.client.TokenManager

    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)

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

func GetListVoicesResult(response *watson.WatsonResponse) *Voices {
    result, ok := response.Result.(*Voices)

    if ok {
        return result
    }

    return nil
}

// Synthesize audio
func (textToSpeech *TextToSpeechV1) Synthesize(body *Text, accept string, voice string, customizationID string) (*watson.WatsonResponse, []error) {
    path := "/v1/synthesize"
    creds := textToSpeech.client.Creds
    useTM := textToSpeech.client.UseTM
    tokenManager := textToSpeech.client.TokenManager

    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Accept", fmt.Sprint(accept))
    request.Query("version=" + creds.Version)
    request.Query("voice=" + fmt.Sprint(voice))
    request.Query("customization_id=" + fmt.Sprint(customizationID))
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

func GetSynthesizeResult(response *watson.WatsonResponse) *os.File {
    result, ok := response.Result.(*os.File)

    if ok {
        return result
    }

    return nil
}

// Get pronunciation
func (textToSpeech *TextToSpeechV1) GetPronunciation(text string, voice string, format string, customizationID string) (*watson.WatsonResponse, []error) {
    path := "/v1/pronunciation"
    creds := textToSpeech.client.Creds
    useTM := textToSpeech.client.UseTM
    tokenManager := textToSpeech.client.TokenManager

    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("text=" + fmt.Sprint(text))
    request.Query("voice=" + fmt.Sprint(voice))
    request.Query("format=" + fmt.Sprint(format))
    request.Query("customization_id=" + fmt.Sprint(customizationID))

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

func GetGetPronunciationResult(response *watson.WatsonResponse) *Pronunciation {
    result, ok := response.Result.(*Pronunciation)

    if ok {
        return result
    }

    return nil
}

// Create a custom model
func (textToSpeech *TextToSpeechV1) CreateVoiceModel(body *CreateVoiceModel) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations"
    creds := textToSpeech.client.Creds
    useTM := textToSpeech.client.UseTM
    tokenManager := textToSpeech.client.TokenManager

    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
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

func GetCreateVoiceModelResult(response *watson.WatsonResponse) *VoiceModel {
    result, ok := response.Result.(*VoiceModel)

    if ok {
        return result
    }

    return nil
}

// Delete a custom model
func (textToSpeech *TextToSpeechV1) DeleteVoiceModel(customizationID string) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}"
    creds := textToSpeech.client.Creds
    useTM := textToSpeech.client.UseTM
    tokenManager := textToSpeech.client.TokenManager

    path = strings.Replace(path, "{customization_id}", customizationID, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)

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


// Get a custom model
func (textToSpeech *TextToSpeechV1) GetVoiceModel(customizationID string) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}"
    creds := textToSpeech.client.Creds
    useTM := textToSpeech.client.UseTM
    tokenManager := textToSpeech.client.TokenManager

    path = strings.Replace(path, "{customization_id}", customizationID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)

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

func GetGetVoiceModelResult(response *watson.WatsonResponse) *VoiceModel {
    result, ok := response.Result.(*VoiceModel)

    if ok {
        return result
    }

    return nil
}

// List custom models
func (textToSpeech *TextToSpeechV1) ListVoiceModels(language string) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations"
    creds := textToSpeech.client.Creds
    useTM := textToSpeech.client.UseTM
    tokenManager := textToSpeech.client.TokenManager

    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
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

func GetListVoiceModelsResult(response *watson.WatsonResponse) *VoiceModels {
    result, ok := response.Result.(*VoiceModels)

    if ok {
        return result
    }

    return nil
}

// Update a custom model
func (textToSpeech *TextToSpeechV1) UpdateVoiceModel(customizationID string, body *UpdateVoiceModel) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}"
    creds := textToSpeech.client.Creds
    useTM := textToSpeech.client.UseTM
    tokenManager := textToSpeech.client.TokenManager

    path = strings.Replace(path, "{customization_id}", customizationID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
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


// Add a custom word
func (textToSpeech *TextToSpeechV1) AddWord(customizationID string, word string, body *Translation) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}/words/{word}"
    creds := textToSpeech.client.Creds
    useTM := textToSpeech.client.UseTM
    tokenManager := textToSpeech.client.TokenManager

    path = strings.Replace(path, "{customization_id}", customizationID, 1)
    path = strings.Replace(path, "{word}", word, 1)
    request := req.New().Put(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
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


// Add custom words
func (textToSpeech *TextToSpeechV1) AddWords(customizationID string, body *Words) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}/words"
    creds := textToSpeech.client.Creds
    useTM := textToSpeech.client.UseTM
    tokenManager := textToSpeech.client.TokenManager

    path = strings.Replace(path, "{customization_id}", customizationID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
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


// Delete a custom word
func (textToSpeech *TextToSpeechV1) DeleteWord(customizationID string, word string) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}/words/{word}"
    creds := textToSpeech.client.Creds
    useTM := textToSpeech.client.UseTM
    tokenManager := textToSpeech.client.TokenManager

    path = strings.Replace(path, "{customization_id}", customizationID, 1)
    path = strings.Replace(path, "{word}", word, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)

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


// Get a custom word
func (textToSpeech *TextToSpeechV1) GetWord(customizationID string, word string) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}/words/{word}"
    creds := textToSpeech.client.Creds
    useTM := textToSpeech.client.UseTM
    tokenManager := textToSpeech.client.TokenManager

    path = strings.Replace(path, "{customization_id}", customizationID, 1)
    path = strings.Replace(path, "{word}", word, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)

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

func GetGetWordResult(response *watson.WatsonResponse) *Translation {
    result, ok := response.Result.(*Translation)

    if ok {
        return result
    }

    return nil
}

// List custom words
func (textToSpeech *TextToSpeechV1) ListWords(customizationID string) (*watson.WatsonResponse, []error) {
    path := "/v1/customizations/{customization_id}/words"
    creds := textToSpeech.client.Creds
    useTM := textToSpeech.client.UseTM
    tokenManager := textToSpeech.client.TokenManager

    path = strings.Replace(path, "{customization_id}", customizationID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)

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

func GetListWordsResult(response *watson.WatsonResponse) *Words {
    result, ok := response.Result.(*Words)

    if ok {
        return result
    }

    return nil
}

// Delete labeled data
func (textToSpeech *TextToSpeechV1) DeleteUserData(customerID string) (*watson.WatsonResponse, []error) {
    path := "/v1/user_data"
    creds := textToSpeech.client.Creds
    useTM := textToSpeech.client.UseTM
    tokenManager := textToSpeech.client.TokenManager

    request := req.New().Delete(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
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



type CreateVoiceModel struct {

	// The name of the new custom voice model.
	Name string `json:"name"`

	// The language of the new custom voice model. Omit the parameter to use the the default language, `en-US`.
	Language string `json:"language,omitempty"`

	// A description of the new custom voice model. Specifying a description is recommended.
	Description string `json:"description,omitempty"`
}

type Pronunciation struct {

	// The pronunciation of the specified text in the requested voice and format. If a custom voice model is specified, the pronunciation also reflects that custom voice.
	Pronunciation string `json:"pronunciation"`
}

type SupportedFeatures struct {

	// If `true`, the voice can be customized; if `false`, the voice cannot be customized. (Same as `customizable`.).
	CustomPronunciation bool `json:"custom_pronunciation"`

	// If `true`, the voice can be transformed by using the SSML &lt;voice-transformation&gt; element; if `false`, the voice cannot be transformed.
	VoiceTransformation bool `json:"voice_transformation"`
}

type Text struct {

	// The text to synthesize.
	Text string `json:"text"`
}

type Translation struct {

	// The phonetic or sounds-like translation for the word. A phonetic translation is based on the SSML format for representing the phonetic string of a word either as an IPA translation or as an IBM SPR translation. A sounds-like is one or more words that, when combined, sound like the word.
	Translation string `json:"translation"`

	// **Japanese only.** The part of speech for the word. The service uses the value to produce the correct intonation for the word. You can create only a single entry, with or without a single part of speech, for any word; you cannot create multiple entries with different parts of speech for the same word. For more information, see [Working with Japanese entries](https://console.bluemix.net/docs/services/text-to-speech/custom-rules.html#jaNotes).
	PartOfSpeech string `json:"part_of_speech,omitempty"`
}

type UpdateVoiceModel struct {

	// A new name for the custom voice model.
	Name string `json:"name,omitempty"`

	// A new description for the custom voice model.
	Description string `json:"description,omitempty"`

	// An array of `Word` objects that provides the words and their translations that are to be added or updated for the custom voice model. Pass an empty array to make no additions or updates.
	Words []Word `json:"words,omitempty"`
}

type Voice struct {

	// The URI of the voice.
	Url string `json:"url"`

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

type VoiceModel struct {

	// The customization ID (GUID) of the custom voice model. The **Create a custom model** method returns only this field. It does not not return the other fields of this object.
	CustomizationId string `json:"customization_id"`

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

type VoiceModels struct {

	// An array of `VoiceModel` objects that provides information about each available custom voice model. The array is empty if the requesting service credentials own no custom voice models (if no language is specified) or own no custom voice models for the specified language.
	Customizations []VoiceModel `json:"customizations"`
}

type Voices struct {

	// A list of available voices.
	Voices []Voice `json:"voices"`
}

type Word struct {

	// A word from the custom voice model.
	Word string `json:"word"`

	// The phonetic or sounds-like translation for the word. A phonetic translation is based on the SSML format for representing the phonetic string of a word either as an IPA or IBM SPR translation. A sounds-like translation consists of one or more words that, when combined, sound like the word.
	Translation string `json:"translation"`

	// **Japanese only.** The part of speech for the word. The service uses the value to produce the correct intonation for the word. You can create only a single entry, with or without a single part of speech, for any word; you cannot create multiple entries with different parts of speech for the same word. For more information, see [Working with Japanese entries](https://console.bluemix.net/docs/services/text-to-speech/custom-rules.html#jaNotes).
	PartOfSpeech string `json:"part_of_speech,omitempty"`
}

type Words struct {

	// The **Add custom words** method accepts an array of `Word` objects. Each object provides a word that is to be added or updated for the custom voice model and the word's translation. The **List custom words** method returns an array of `Word` objects. Each object shows a word and its translation from the custom voice model. The words are listed in alphabetical order, with uppercase letters listed before lowercase letters. The array is empty if the custom model contains no words.
	Words []Word `json:"words"`
}
