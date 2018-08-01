// Package naturallanguageclassifierv1 : Operations and models for the NaturalLanguageClassifierV1 service
package naturallanguageclassifierv1
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
    "github.com/go-openapi/strfmt"
    "os"
    "strings"
    req "github.com/parnurzeal/gorequest"
    watson "golang-sdk"
)

// NaturalLanguageClassifierV1 : The NaturalLanguageClassifierV1 service
type NaturalLanguageClassifierV1 struct {
	client *watson.Client
}

// NewNaturalLanguageClassifierV1 : Instantiate NaturalLanguageClassifierV1
func NewNaturalLanguageClassifierV1(creds watson.Credentials) (*NaturalLanguageClassifierV1, error) {
    if creds.ServiceURL == "" {
        creds.ServiceURL = "https://gateway.watsonplatform.net/natural-language-classifier/api"
    }

	client, clientErr := watson.NewClient(creds, "natural_language_classifier")

	if clientErr != nil {
		return nil, clientErr
	}

	return &NaturalLanguageClassifierV1{ client: client }, nil
}

// Classify : Classify a phrase
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) Classify(options *ClassifyOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/classifiers/{classifier_id}/classify"
    creds := naturalLanguageClassifier.client.Creds
    useTM := naturalLanguageClassifier.client.UseTM
    tokenManager := naturalLanguageClassifier.client.TokenManager

    path = strings.Replace(path, "{classifier_id}", options.ClassifierID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
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

    response.Result = new(Classification)
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

// GetClassifyResult : Cast result of Classify operation
func GetClassifyResult(response *watson.WatsonResponse) *Classification {
    result, ok := response.Result.(*Classification)

    if ok {
        return result
    }

    return nil
}

// ClassifyCollection : Classify multiple phrases
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) ClassifyCollection(options *ClassifyCollectionOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/classifiers/{classifier_id}/classify_collection"
    creds := naturalLanguageClassifier.client.Creds
    useTM := naturalLanguageClassifier.client.UseTM
    tokenManager := naturalLanguageClassifier.client.TokenManager

    path = strings.Replace(path, "{classifier_id}", options.ClassifierID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    body := map[string]interface{}{}
    body["collection"] = options.Collection
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

    response.Result = new(ClassificationCollection)
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

// GetClassifyCollectionResult : Cast result of ClassifyCollection operation
func GetClassifyCollectionResult(response *watson.WatsonResponse) *ClassificationCollection {
    result, ok := response.Result.(*ClassificationCollection)

    if ok {
        return result
    }

    return nil
}

// CreateClassifier : Create classifier
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) CreateClassifier(options *CreateClassifierOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/classifiers"
    creds := naturalLanguageClassifier.client.Creds
    useTM := naturalLanguageClassifier.client.UseTM
    tokenManager := naturalLanguageClassifier.client.TokenManager

    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Type("multipart")
    request.SendFile(options.Metadata, "", "training_metadata")
    request.SendFile(options.TrainingData, "", "training_data")

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

    response.Result = new(Classifier)
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

// GetCreateClassifierResult : Cast result of CreateClassifier operation
func GetCreateClassifierResult(response *watson.WatsonResponse) *Classifier {
    result, ok := response.Result.(*Classifier)

    if ok {
        return result
    }

    return nil
}

// DeleteClassifier : Delete classifier
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) DeleteClassifier(options *DeleteClassifierOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/classifiers/{classifier_id}"
    creds := naturalLanguageClassifier.client.Creds
    useTM := naturalLanguageClassifier.client.UseTM
    tokenManager := naturalLanguageClassifier.client.TokenManager

    path = strings.Replace(path, "{classifier_id}", options.ClassifierID, 1)
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


// GetClassifier : Get information about a classifier
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) GetClassifier(options *GetClassifierOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/classifiers/{classifier_id}"
    creds := naturalLanguageClassifier.client.Creds
    useTM := naturalLanguageClassifier.client.UseTM
    tokenManager := naturalLanguageClassifier.client.TokenManager

    path = strings.Replace(path, "{classifier_id}", options.ClassifierID, 1)
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

    response.Result = new(Classifier)
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

// GetGetClassifierResult : Cast result of GetClassifier operation
func GetGetClassifierResult(response *watson.WatsonResponse) *Classifier {
    result, ok := response.Result.(*Classifier)

    if ok {
        return result
    }

    return nil
}

// ListClassifiers : List classifiers
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) ListClassifiers(options *ListClassifiersOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/classifiers"
    creds := naturalLanguageClassifier.client.Creds
    useTM := naturalLanguageClassifier.client.UseTM
    tokenManager := naturalLanguageClassifier.client.TokenManager

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

    response.Result = new(ClassifierList)
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

// GetListClassifiersResult : Cast result of ListClassifiers operation
func GetListClassifiersResult(response *watson.WatsonResponse) *ClassifierList {
    result, ok := response.Result.(*ClassifierList)

    if ok {
        return result
    }

    return nil
}


// Classification : Response from the classifier for a phrase.
type Classification struct {

	// Unique identifier for this classifier.
	ClassifierID string `json:"classifier_id,omitempty"`

	// Link to the classifier.
	URL string `json:"url,omitempty"`

	// The submitted phrase.
	Text string `json:"text,omitempty"`

	// The class with the highest confidence.
	TopClass string `json:"top_class,omitempty"`

	// An array of up to ten class-confidence pairs sorted in descending order of confidence.
	Classes []ClassifiedClass `json:"classes,omitempty"`
}

// ClassificationCollection : Response from the classifier for multiple phrases.
type ClassificationCollection struct {

	// Unique identifier for this classifier.
	ClassifierID string `json:"classifier_id,omitempty"`

	// Link to the classifier.
	URL string `json:"url,omitempty"`

	// An array of classifier responses for each submitted phrase.
	Collection []CollectionItem `json:"collection,omitempty"`
}

// ClassifiedClass : Class and confidence.
type ClassifiedClass struct {

	// A decimal percentage that represents the confidence that Watson has in this class. Higher values represent higher confidences.
	Confidence float64 `json:"confidence,omitempty"`

	// Class label.
	ClassName string `json:"class_name,omitempty"`
}

// Classifier : A classifier for natural language phrases.
type Classifier struct {

	// User-supplied name for the classifier.
	Name string `json:"name,omitempty"`

	// Link to the classifier.
	URL string `json:"url"`

	// The state of the classifier.
	Status string `json:"status,omitempty"`

	// Unique identifier for this classifier.
	ClassifierID string `json:"classifier_id"`

	// Date and time (UTC) the classifier was created.
	Created strfmt.DateTime `json:"created,omitempty"`

	// Additional detail about the status.
	StatusDescription string `json:"status_description,omitempty"`

	// The language used for the classifier.
	Language string `json:"language,omitempty"`
}

// ClassifierList : List of available classifiers.
type ClassifierList struct {

	// The classifiers available to the user. Returns an empty array if no classifiers are available.
	Classifiers []Classifier `json:"classifiers"`
}

// ClassifyCollectionOptions : The classifyCollection options.
type ClassifyCollectionOptions struct {

	// Classifier ID to use.
	ClassifierID string `json:"classifier_id"`

	// The submitted phrases.
	Collection []ClassifyInput `json:"collection"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewClassifyCollectionOptions : Instantiate ClassifyCollectionOptions
func NewClassifyCollectionOptions(classifierID string, collection []ClassifyInput) *ClassifyCollectionOptions {
    return &ClassifyCollectionOptions{
        ClassifierID: classifierID,
        Collection: collection,
    }
}

// SetClassifierID : Allow user to set ClassifierID
func (options *ClassifyCollectionOptions) SetClassifierID(param string) *ClassifyCollectionOptions {
    options.ClassifierID = param
    return options
}

// SetCollection : Allow user to set Collection
func (options *ClassifyCollectionOptions) SetCollection(param []ClassifyInput) *ClassifyCollectionOptions {
    options.Collection = param
    return options
}

// SetHeaders : Allow user to set Headers
func (options *ClassifyCollectionOptions) SetHeaders(param map[string]string) *ClassifyCollectionOptions {
    options.Headers = param
    return options
}

// ClassifyInput : Request payload to classify.
type ClassifyInput struct {

	// The submitted phrase.
	Text string `json:"text"`
}

// ClassifyOptions : The classify options.
type ClassifyOptions struct {

	// Classifier ID to use.
	ClassifierID string `json:"classifier_id"`

	// The submitted phrase.
	Text string `json:"text"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewClassifyOptions : Instantiate ClassifyOptions
func NewClassifyOptions(classifierID string, text string) *ClassifyOptions {
    return &ClassifyOptions{
        ClassifierID: classifierID,
        Text: text,
    }
}

// SetClassifierID : Allow user to set ClassifierID
func (options *ClassifyOptions) SetClassifierID(param string) *ClassifyOptions {
    options.ClassifierID = param
    return options
}

// SetText : Allow user to set Text
func (options *ClassifyOptions) SetText(param string) *ClassifyOptions {
    options.Text = param
    return options
}

// SetHeaders : Allow user to set Headers
func (options *ClassifyOptions) SetHeaders(param map[string]string) *ClassifyOptions {
    options.Headers = param
    return options
}

// CollectionItem : Response from the classifier for a phrase in a collection.
type CollectionItem struct {

	// The submitted phrase.
	Text string `json:"text,omitempty"`

	// The class with the highest confidence.
	TopClass string `json:"top_class,omitempty"`

	// An array of up to ten class-confidence pairs sorted in descending order of confidence.
	Classes []ClassifiedClass `json:"classes,omitempty"`
}

// CreateClassifierOptions : The createClassifier options.
type CreateClassifierOptions struct {

	// Metadata in JSON format. The metadata identifies the language of the data, and an optional name to identify the classifier. Specify the language with the 2-letter primary language code as assigned in ISO standard 639. Supported languages are English (`en`), Arabic (`ar`), French (`fr`), German, (`de`), Italian (`it`), Japanese (`ja`), Korean (`ko`), Brazilian Portuguese (`pt`), and Spanish (`es`).
	Metadata os.File `json:"training_metadata"`

	// Training data in CSV format. Each text value must have at least one class. The data can include up to 20,000 records. For details, see [Data preparation](https://console.bluemix.net/docs/services/natural-language-classifier/using-your-data.html).
	TrainingData os.File `json:"training_data"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewCreateClassifierOptions : Instantiate CreateClassifierOptions
func NewCreateClassifierOptions(trainingMetadata os.File, trainingData os.File) *CreateClassifierOptions {
    return &CreateClassifierOptions{
        Metadata: trainingMetadata,
        TrainingData: trainingData,
    }
}

// SetMetadata : Allow user to set Metadata
func (options *CreateClassifierOptions) SetMetadata(param os.File) *CreateClassifierOptions {
    options.Metadata = param
    return options
}

// SetTrainingData : Allow user to set TrainingData
func (options *CreateClassifierOptions) SetTrainingData(param os.File) *CreateClassifierOptions {
    options.TrainingData = param
    return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateClassifierOptions) SetHeaders(param map[string]string) *CreateClassifierOptions {
    options.Headers = param
    return options
}

// DeleteClassifierOptions : The deleteClassifier options.
type DeleteClassifierOptions struct {

	// Classifier ID to delete.
	ClassifierID string `json:"classifier_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewDeleteClassifierOptions : Instantiate DeleteClassifierOptions
func NewDeleteClassifierOptions(classifierID string) *DeleteClassifierOptions {
    return &DeleteClassifierOptions{
        ClassifierID: classifierID,
    }
}

// SetClassifierID : Allow user to set ClassifierID
func (options *DeleteClassifierOptions) SetClassifierID(param string) *DeleteClassifierOptions {
    options.ClassifierID = param
    return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteClassifierOptions) SetHeaders(param map[string]string) *DeleteClassifierOptions {
    options.Headers = param
    return options
}

// GetClassifierOptions : The getClassifier options.
type GetClassifierOptions struct {

	// Classifier ID to query.
	ClassifierID string `json:"classifier_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewGetClassifierOptions : Instantiate GetClassifierOptions
func NewGetClassifierOptions(classifierID string) *GetClassifierOptions {
    return &GetClassifierOptions{
        ClassifierID: classifierID,
    }
}

// SetClassifierID : Allow user to set ClassifierID
func (options *GetClassifierOptions) SetClassifierID(param string) *GetClassifierOptions {
    options.ClassifierID = param
    return options
}

// SetHeaders : Allow user to set Headers
func (options *GetClassifierOptions) SetHeaders(param map[string]string) *GetClassifierOptions {
    options.Headers = param
    return options
}

// ListClassifiersOptions : The listClassifiers options.
type ListClassifiersOptions struct {

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewListClassifiersOptions : Instantiate ListClassifiersOptions
func NewListClassifiersOptions() *ListClassifiersOptions {
    return &ListClassifiersOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListClassifiersOptions) SetHeaders(param map[string]string) *ListClassifiersOptions {
    options.Headers = param
    return options
}
