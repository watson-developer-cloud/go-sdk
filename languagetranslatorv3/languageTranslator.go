// Package languagetranslatorv3 : Operations and models for the LanguageTranslatorV3 service
package languagetranslatorv3
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

// LanguageTranslatorV3 : The LanguageTranslatorV3 service
type LanguageTranslatorV3 struct {
	client *watson.Client
}

// NewLanguageTranslatorV3 : Instantiate LanguageTranslatorV3
func NewLanguageTranslatorV3(creds watson.Credentials) (*LanguageTranslatorV3, error) {
    if creds.ServiceURL == "" {
        creds.ServiceURL = "https://gateway.watsonplatform.net/language-translator/api"
    }

	client, clientErr := watson.NewClient(creds, "language_translator")

	if clientErr != nil {
		return nil, clientErr
	}

	return &LanguageTranslatorV3{ client: client }, nil
}

// Translate : Translate
func (languageTranslator *LanguageTranslatorV3) Translate(body *TranslateRequest) (*watson.WatsonResponse, []error) {
    path := "/v3/translate"
    creds := languageTranslator.client.Creds
    useTM := languageTranslator.client.UseTM
    tokenManager := languageTranslator.client.TokenManager

    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
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

    response.Result = new(TranslationResult)
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

// GetTranslateResult : Cast result of Translate operation
func GetTranslateResult(response *watson.WatsonResponse) *TranslationResult {
    result, ok := response.Result.(*TranslationResult)

    if ok {
        return result
    }

    return nil
}

// Identify : Identify language
func (languageTranslator *LanguageTranslatorV3) Identify(body *string) (*watson.WatsonResponse, []error) {
    path := "/v3/identify"
    creds := languageTranslator.client.Creds
    useTM := languageTranslator.client.UseTM
    tokenManager := languageTranslator.client.TokenManager

    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "text/plain")
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

    response.Result = new(IdentifiedLanguages)
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

// GetIdentifyResult : Cast result of Identify operation
func GetIdentifyResult(response *watson.WatsonResponse) *IdentifiedLanguages {
    result, ok := response.Result.(*IdentifiedLanguages)

    if ok {
        return result
    }

    return nil
}

// ListIdentifiableLanguages : List identifiable languages
func (languageTranslator *LanguageTranslatorV3) ListIdentifiableLanguages() (*watson.WatsonResponse, []error) {
    path := "/v3/identifiable_languages"
    creds := languageTranslator.client.Creds
    useTM := languageTranslator.client.UseTM
    tokenManager := languageTranslator.client.TokenManager

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

    response.Result = new(IdentifiableLanguages)
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

// GetListIdentifiableLanguagesResult : Cast result of ListIdentifiableLanguages operation
func GetListIdentifiableLanguagesResult(response *watson.WatsonResponse) *IdentifiableLanguages {
    result, ok := response.Result.(*IdentifiableLanguages)

    if ok {
        return result
    }

    return nil
}

// CreateModel : Create model
func (languageTranslator *LanguageTranslatorV3) CreateModel(baseModelID string, name string, forcedGlossary os.File, parallelCorpus os.File) (*watson.WatsonResponse, []error) {
    path := "/v3/models"
    creds := languageTranslator.client.Creds
    useTM := languageTranslator.client.UseTM
    tokenManager := languageTranslator.client.TokenManager

    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "multipart/form-data")
    request.Query("version=" + creds.Version)
    request.Query("base_model_id=" + fmt.Sprint(baseModelID))
    request.Query("name=" + fmt.Sprint(name))
    request.Type("multipart")
    request.SendFile(forcedGlossary)
    request.SendFile(parallelCorpus)

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

    response.Result = new(TranslationModel)
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

// GetCreateModelResult : Cast result of CreateModel operation
func GetCreateModelResult(response *watson.WatsonResponse) *TranslationModel {
    result, ok := response.Result.(*TranslationModel)

    if ok {
        return result
    }

    return nil
}

// DeleteModel : Delete model
func (languageTranslator *LanguageTranslatorV3) DeleteModel(modelID string) (*watson.WatsonResponse, []error) {
    path := "/v3/models/{model_id}"
    creds := languageTranslator.client.Creds
    useTM := languageTranslator.client.UseTM
    tokenManager := languageTranslator.client.TokenManager

    path = strings.Replace(path, "{model_id}", modelID, 1)
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

    response.Result = new(DeleteModelResult)
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

// GetDeleteModelResult : Cast result of DeleteModel operation
func GetDeleteModelResult(response *watson.WatsonResponse) *DeleteModelResult {
    result, ok := response.Result.(*DeleteModelResult)

    if ok {
        return result
    }

    return nil
}

// GetModel : Get model details
func (languageTranslator *LanguageTranslatorV3) GetModel(modelID string) (*watson.WatsonResponse, []error) {
    path := "/v3/models/{model_id}"
    creds := languageTranslator.client.Creds
    useTM := languageTranslator.client.UseTM
    tokenManager := languageTranslator.client.TokenManager

    path = strings.Replace(path, "{model_id}", modelID, 1)
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

    response.Result = new(TranslationModel)
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
func GetGetModelResult(response *watson.WatsonResponse) *TranslationModel {
    result, ok := response.Result.(*TranslationModel)

    if ok {
        return result
    }

    return nil
}

// ListModels : List models
func (languageTranslator *LanguageTranslatorV3) ListModels(source string, target string, defaultModels bool) (*watson.WatsonResponse, []error) {
    path := "/v3/models"
    creds := languageTranslator.client.Creds
    useTM := languageTranslator.client.UseTM
    tokenManager := languageTranslator.client.TokenManager

    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("source=" + fmt.Sprint(source))
    request.Query("target=" + fmt.Sprint(target))
    request.Query("default=" + fmt.Sprint(defaultModels))

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

    response.Result = new(TranslationModels)
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
func GetListModelsResult(response *watson.WatsonResponse) *TranslationModels {
    result, ok := response.Result.(*TranslationModels)

    if ok {
        return result
    }

    return nil
}


// DeleteModelResult : DeleteModelResult struct
type DeleteModelResult struct {

	// "OK" indicates that the model was successfully deleted.
	Status string `json:"status"`
}

// IdentifiableLanguage : IdentifiableLanguage struct
type IdentifiableLanguage struct {

	// The language code for an identifiable language.
	Language string `json:"language"`

	// The name of the identifiable language.
	Name string `json:"name"`
}

// IdentifiableLanguages : IdentifiableLanguages struct
type IdentifiableLanguages struct {

	// A list of all languages that the service can identify.
	Languages []IdentifiableLanguage `json:"languages"`
}

// IdentifiedLanguage : IdentifiedLanguage struct
type IdentifiedLanguage struct {

	// The language code for an identified language.
	Language string `json:"language"`

	// The confidence score for the identified language.
	Confidence float64 `json:"confidence"`
}

// IdentifiedLanguages : IdentifiedLanguages struct
type IdentifiedLanguages struct {

	// A ranking of identified languages with confidence scores.
	Languages []IdentifiedLanguage `json:"languages"`
}

// TranslateRequest : TranslateRequest struct
type TranslateRequest struct {

	// Input text in UTF-8 encoding. Multiple entries will result in multiple translations in the response.
	Text []string `json:"text"`

	// Model ID of the translation model to use. If this is specified, the **source** and **target** parameters will be ignored. The method requires either a model ID or both the **source** and **target** parameters.
	ModelID string `json:"model_id,omitempty"`

	// Language code of the source text language. Use with `target` as an alternative way to select a translation model. When `source` and `target` are set, and a model ID is not set, the system chooses a default model for the language pair (usually the model based on the news domain).
	Source string `json:"source,omitempty"`

	// Language code of the translation target language. Use with source as an alternative way to select a translation model.
	Target string `json:"target,omitempty"`
}

// Translation : Translation struct
type Translation struct {

	// Translation output in UTF-8.
	TranslationOutput string `json:"translation_output"`
}

// TranslationModel : Response payload for models.
type TranslationModel struct {

	// A globally unique string that identifies the underlying model that is used for translation.
	ModelID string `json:"model_id"`

	// Optional name that can be specified when the model is created.
	Name string `json:"name,omitempty"`

	// Translation source language code.
	Source string `json:"source,omitempty"`

	// Translation target language code.
	Target string `json:"target,omitempty"`

	// Model ID of the base model that was used to customize the model. If the model is not a custom model, this will be an empty string.
	BaseModelID string `json:"base_model_id,omitempty"`

	// The domain of the translation model.
	Domain string `json:"domain,omitempty"`

	// Whether this model can be used as a base for customization. Customized models are not further customizable, and some base models are not customizable.
	Customizable bool `json:"customizable,omitempty"`

	// Whether or not the model is a default model. A default model is the model for a given language pair that will be used when that language pair is specified in the source and target parameters.
	DefaultModel bool `json:"default_model,omitempty"`

	// Either an empty string, indicating the model is not a custom model, or the ID of the service instance that created the model.
	Owner string `json:"owner,omitempty"`

	// Availability of a model.
	Status string `json:"status,omitempty"`
}

// TranslationModels : The response type for listing existing translation models.
type TranslationModels struct {

	// An array of available models.
	Models []TranslationModel `json:"models"`
}

// TranslationResult : TranslationResult struct
type TranslationResult struct {

	// Number of words in the input text.
	WordCount int64 `json:"word_count"`

	// Number of characters in the input text.
	CharacterCount int64 `json:"character_count"`

	// List of translation output in UTF-8, corresponding to the input text entries.
	Translations []Translation `json:"translations"`
}
