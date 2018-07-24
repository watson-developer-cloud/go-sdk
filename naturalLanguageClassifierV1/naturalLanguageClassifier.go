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
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) Classify(classifierID string, body *ClassifyInput) (*watson.WatsonResponse, []error) {
    path := "/v1/classifiers/{classifier_id}/classify"
    creds := naturalLanguageClassifier.client.Creds
    useTM := naturalLanguageClassifier.client.UseTM
    tokenManager := naturalLanguageClassifier.client.TokenManager

    path = strings.Replace(path, "{classifier_id}", classifierID, 1)
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
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) ClassifyCollection(classifierID string, body *ClassifyCollectionInput) (*watson.WatsonResponse, []error) {
    path := "/v1/classifiers/{classifier_id}/classify_collection"
    creds := naturalLanguageClassifier.client.Creds
    useTM := naturalLanguageClassifier.client.UseTM
    tokenManager := naturalLanguageClassifier.client.TokenManager

    path = strings.Replace(path, "{classifier_id}", classifierID, 1)
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
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) CreateClassifier(metadata os.File, trainingData os.File) (*watson.WatsonResponse, []error) {
    path := "/v1/classifiers"
    creds := naturalLanguageClassifier.client.Creds
    useTM := naturalLanguageClassifier.client.UseTM
    tokenManager := naturalLanguageClassifier.client.TokenManager

    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "multipart/form-data")
    request.Type("multipart")
    request.SendFile(metadata)
    request.SendFile(trainingData)

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
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) DeleteClassifier(classifierID string) (*watson.WatsonResponse, []error) {
    path := "/v1/classifiers/{classifier_id}"
    creds := naturalLanguageClassifier.client.Creds
    useTM := naturalLanguageClassifier.client.UseTM
    tokenManager := naturalLanguageClassifier.client.TokenManager

    path = strings.Replace(path, "{classifier_id}", classifierID, 1)
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


// GetClassifier : Get information about a classifier
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) GetClassifier(classifierID string) (*watson.WatsonResponse, []error) {
    path := "/v1/classifiers/{classifier_id}"
    creds := naturalLanguageClassifier.client.Creds
    useTM := naturalLanguageClassifier.client.UseTM
    tokenManager := naturalLanguageClassifier.client.TokenManager

    path = strings.Replace(path, "{classifier_id}", classifierID, 1)
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
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) ListClassifiers() (*watson.WatsonResponse, []error) {
    path := "/v1/classifiers"
    creds := naturalLanguageClassifier.client.Creds
    useTM := naturalLanguageClassifier.client.UseTM
    tokenManager := naturalLanguageClassifier.client.TokenManager

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

// ClassifyCollectionInput : Request payload to classify.
type ClassifyCollectionInput struct {

	// The submitted phrases.
	Collection []ClassifyInput `json:"collection"`
}

// ClassifyInput : Request payload to classify.
type ClassifyInput struct {

	// The submitted phrase.
	Text string `json:"text"`
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
