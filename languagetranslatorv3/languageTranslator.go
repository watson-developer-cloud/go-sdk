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
func (languageTranslator *LanguageTranslatorV3) Translate(options *TranslateOptions) (*watson.WatsonResponse, []error) {
    path := "/v3/translate"
    creds := languageTranslator.client.Creds
    useTM := languageTranslator.client.UseTM
    tokenManager := languageTranslator.client.TokenManager

    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    body := map[string]interface{}{}
    body["text"] = options.Text
    if options.IsModelIDSet {
        body["model_id"] = options.ModelID
    }
    if options.IsSourceSet {
        body["source"] = options.Source
    }
    if options.IsTargetSet {
        body["target"] = options.Target
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
func (languageTranslator *LanguageTranslatorV3) Identify(options *IdentifyOptions) (*watson.WatsonResponse, []error) {
    path := "/v3/identify"
    creds := languageTranslator.client.Creds
    useTM := languageTranslator.client.UseTM
    tokenManager := languageTranslator.client.TokenManager

    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "text/plain")
    request.Query("version=" + creds.Version)
    request.Send(options.Text)

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
func (languageTranslator *LanguageTranslatorV3) ListIdentifiableLanguages(options *ListIdentifiableLanguagesOptions) (*watson.WatsonResponse, []error) {
    path := "/v3/identifiable_languages"
    creds := languageTranslator.client.Creds
    useTM := languageTranslator.client.UseTM
    tokenManager := languageTranslator.client.TokenManager

    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

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
func (languageTranslator *LanguageTranslatorV3) CreateModel(options *CreateModelOptions) (*watson.WatsonResponse, []error) {
    path := "/v3/models"
    creds := languageTranslator.client.Creds
    useTM := languageTranslator.client.UseTM
    tokenManager := languageTranslator.client.TokenManager

    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("base_model_id=" + fmt.Sprint(options.BaseModelID))
    if options.IsNameSet {
        request.Query("name=" + fmt.Sprint(options.Name))
    }
    request.Type("multipart")
    if options.IsForcedGlossarySet {
        request.SendFile(options.ForcedGlossary, "", "forced_glossary")
    }
    if options.IsParallelCorpusSet {
        request.SendFile(options.ParallelCorpus, "", "parallel_corpus")
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
func (languageTranslator *LanguageTranslatorV3) DeleteModel(options *DeleteModelOptions) (*watson.WatsonResponse, []error) {
    path := "/v3/models/{model_id}"
    creds := languageTranslator.client.Creds
    useTM := languageTranslator.client.UseTM
    tokenManager := languageTranslator.client.TokenManager

    path = strings.Replace(path, "{model_id}", options.ModelID, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

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
func (languageTranslator *LanguageTranslatorV3) GetModel(options *GetModelOptions) (*watson.WatsonResponse, []error) {
    path := "/v3/models/{model_id}"
    creds := languageTranslator.client.Creds
    useTM := languageTranslator.client.UseTM
    tokenManager := languageTranslator.client.TokenManager

    path = strings.Replace(path, "{model_id}", options.ModelID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

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
func (languageTranslator *LanguageTranslatorV3) ListModels(options *ListModelsOptions) (*watson.WatsonResponse, []error) {
    path := "/v3/models"
    creds := languageTranslator.client.Creds
    useTM := languageTranslator.client.UseTM
    tokenManager := languageTranslator.client.TokenManager

    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    if options.IsSourceSet {
        request.Query("source=" + fmt.Sprint(options.Source))
    }
    if options.IsTargetSet {
        request.Query("target=" + fmt.Sprint(options.Target))
    }
    if options.IsDefaultModelsSet {
        request.Query("default=" + fmt.Sprint(options.DefaultModels))
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


// CreateModelOptions : The createModel options.
type CreateModelOptions struct {

	// The model ID of the model to use as the base for customization. To see available models, use the `List models` method. Usually all IBM provided models are customizable. In addition, all your models that have been created via parallel corpus customization, can be further customized with a forced glossary.
	BaseModelID string `json:"base_model_id"`

	// An optional model name that you can use to identify the model. Valid characters are letters, numbers, dashes, underscores, spaces and apostrophes. The maximum length is 32 characters.
	Name string `json:"name,omitempty"`

    // Indicates whether user set optional parameter Name
    IsNameSet bool

	// A TMX file with your customizations. The customizations in the file completely overwrite the domain translaton data, including high frequency or high confidence phrase translations. You can upload only one glossary with a file size less than 10 MB per call. A forced glossary should contain single words or short phrases.
	ForcedGlossary os.File `json:"forced_glossary,omitempty"`

    // Indicates whether user set optional parameter ForcedGlossary
    IsForcedGlossarySet bool

	// A TMX file with parallel sentences for source and target language. You can upload multiple parallel_corpus files in one request. All uploaded parallel_corpus files combined, your parallel corpus must contain at least 5,000 parallel sentences to train successfully.
	ParallelCorpus os.File `json:"parallel_corpus,omitempty"`

    // Indicates whether user set optional parameter ParallelCorpus
    IsParallelCorpusSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewCreateModelOptions : Instantiate CreateModelOptions
func NewCreateModelOptions(baseModelID string) *CreateModelOptions {
    return &CreateModelOptions{
        BaseModelID: baseModelID,
    }
}

// SetBaseModelID : Allow user to set BaseModelID
func (options *CreateModelOptions) SetBaseModelID(param string) *CreateModelOptions {
    options.BaseModelID = param
    return options
}

// SetName : Allow user to set Name
func (options *CreateModelOptions) SetName(param string) *CreateModelOptions {
    options.Name = param
    options.IsNameSet = true
    return options
}

// SetForcedGlossary : Allow user to set ForcedGlossary
func (options *CreateModelOptions) SetForcedGlossary(param os.File) *CreateModelOptions {
    options.ForcedGlossary = param
    options.IsForcedGlossarySet = true
    return options
}

// SetParallelCorpus : Allow user to set ParallelCorpus
func (options *CreateModelOptions) SetParallelCorpus(param os.File) *CreateModelOptions {
    options.ParallelCorpus = param
    options.IsParallelCorpusSet = true
    return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateModelOptions) SetHeaders(param map[string]string) *CreateModelOptions {
    options.Headers = param
    return options
}

// DeleteModelOptions : The deleteModel options.
type DeleteModelOptions struct {

	// Model ID of the model to delete.
	ModelID string `json:"model_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewDeleteModelOptions : Instantiate DeleteModelOptions
func NewDeleteModelOptions(modelID string) *DeleteModelOptions {
    return &DeleteModelOptions{
        ModelID: modelID,
    }
}

// SetModelID : Allow user to set ModelID
func (options *DeleteModelOptions) SetModelID(param string) *DeleteModelOptions {
    options.ModelID = param
    return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteModelOptions) SetHeaders(param map[string]string) *DeleteModelOptions {
    options.Headers = param
    return options
}

// DeleteModelResult : DeleteModelResult struct
type DeleteModelResult struct {

	// "OK" indicates that the model was successfully deleted.
	Status string `json:"status"`
}

// GetModelOptions : The getModel options.
type GetModelOptions struct {

	// Model ID of the model to get.
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

// IdentifyOptions : The identify options.
type IdentifyOptions struct {

	// Input text in UTF-8 format.
	Text string `json:"text"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewIdentifyOptions : Instantiate IdentifyOptions
func NewIdentifyOptions(text string) *IdentifyOptions {
    return &IdentifyOptions{
        Text: text,
    }
}

// SetText : Allow user to set Text
func (options *IdentifyOptions) SetText(param string) *IdentifyOptions {
    options.Text = param
    return options
}

// SetHeaders : Allow user to set Headers
func (options *IdentifyOptions) SetHeaders(param map[string]string) *IdentifyOptions {
    options.Headers = param
    return options
}

// ListIdentifiableLanguagesOptions : The listIdentifiableLanguages options.
type ListIdentifiableLanguagesOptions struct {

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewListIdentifiableLanguagesOptions : Instantiate ListIdentifiableLanguagesOptions
func NewListIdentifiableLanguagesOptions() *ListIdentifiableLanguagesOptions {
    return &ListIdentifiableLanguagesOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListIdentifiableLanguagesOptions) SetHeaders(param map[string]string) *ListIdentifiableLanguagesOptions {
    options.Headers = param
    return options
}

// ListModelsOptions : The listModels options.
type ListModelsOptions struct {

	// Specify a language code to filter results by source language.
	Source string `json:"source,omitempty"`

    // Indicates whether user set optional parameter Source
    IsSourceSet bool

	// Specify a language code to filter results by target language.
	Target string `json:"target,omitempty"`

    // Indicates whether user set optional parameter Target
    IsTargetSet bool

	// If the default parameter isn't specified, the service will return all models (default and non-default) for each language pair. To return only default models, set this to `true`. To return only non-default models, set this to `false`. There is exactly one default model per language pair, the IBM provided base model.
	DefaultModels bool `json:"default,omitempty"`

    // Indicates whether user set optional parameter DefaultModels
    IsDefaultModelsSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewListModelsOptions : Instantiate ListModelsOptions
func NewListModelsOptions() *ListModelsOptions {
    return &ListModelsOptions{}
}

// SetSource : Allow user to set Source
func (options *ListModelsOptions) SetSource(param string) *ListModelsOptions {
    options.Source = param
    options.IsSourceSet = true
    return options
}

// SetTarget : Allow user to set Target
func (options *ListModelsOptions) SetTarget(param string) *ListModelsOptions {
    options.Target = param
    options.IsTargetSet = true
    return options
}

// SetDefaultModels : Allow user to set DefaultModels
func (options *ListModelsOptions) SetDefaultModels(param bool) *ListModelsOptions {
    options.DefaultModels = param
    options.IsDefaultModelsSet = true
    return options
}

// SetHeaders : Allow user to set Headers
func (options *ListModelsOptions) SetHeaders(param map[string]string) *ListModelsOptions {
    options.Headers = param
    return options
}

// TranslateOptions : The translate options.
type TranslateOptions struct {

	// Input text in UTF-8 encoding. Multiple entries will result in multiple translations in the response.
	Text []string `json:"text"`

	// Model ID of the translation model to use. If this is specified, the **source** and **target** parameters will be ignored. The method requires either a model ID or both the **source** and **target** parameters.
	ModelID string `json:"model_id,omitempty"`

    // Indicates whether user set optional parameter ModelID
    IsModelIDSet bool

	// Language code of the source text language. Use with `target` as an alternative way to select a translation model. When `source` and `target` are set, and a model ID is not set, the system chooses a default model for the language pair (usually the model based on the news domain).
	Source string `json:"source,omitempty"`

    // Indicates whether user set optional parameter Source
    IsSourceSet bool

	// Language code of the translation target language. Use with source as an alternative way to select a translation model.
	Target string `json:"target,omitempty"`

    // Indicates whether user set optional parameter Target
    IsTargetSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewTranslateOptions : Instantiate TranslateOptions
func NewTranslateOptions(text []string) *TranslateOptions {
    return &TranslateOptions{
        Text: text,
    }
}

// SetText : Allow user to set Text
func (options *TranslateOptions) SetText(param []string) *TranslateOptions {
    options.Text = param
    return options
}

// SetModelID : Allow user to set ModelID
func (options *TranslateOptions) SetModelID(param string) *TranslateOptions {
    options.ModelID = param
    options.IsModelIDSet = true
    return options
}

// SetSource : Allow user to set Source
func (options *TranslateOptions) SetSource(param string) *TranslateOptions {
    options.Source = param
    options.IsSourceSet = true
    return options
}

// SetTarget : Allow user to set Target
func (options *TranslateOptions) SetTarget(param string) *TranslateOptions {
    options.Target = param
    options.IsTargetSet = true
    return options
}

// SetHeaders : Allow user to set Headers
func (options *TranslateOptions) SetHeaders(param map[string]string) *TranslateOptions {
    options.Headers = param
    return options
}

// Translation : Translation struct
type Translation struct {

	// Translation output in UTF-8.
	TranslationOutput string `json:"translation"`
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
