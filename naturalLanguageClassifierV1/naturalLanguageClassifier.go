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
package naturalLanguageClassifierV1

import (
    "bytes"
    "fmt"
    "github.com/go-openapi/strfmt"
    "os"
    "strings"
    req "github.com/parnurzeal/gorequest"
    watson "golang-sdk"
)

type NaturalLanguageClassifierV1 struct {
	client *watson.Client
}

func NewNaturalLanguageClassifierV1(creds watson.Credentials) (*NaturalLanguageClassifierV1, error) {
	client, clientErr := watson.NewClient(creds, "natural_language_classifier")

	if clientErr != nil {
		return nil, clientErr
	}

	return &NaturalLanguageClassifierV1{ client: client }, nil
}

// Classify a phrase
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) Classify(classifierID string, body *ClassifyInput) (*watson.WatsonResponse, []error) {
    path := "/v1/classifiers/{classifier_id}/classify"
    creds := naturalLanguageClassifier.client.Creds
    useTM := naturalLanguageClassifier.client.UseTM
    tokenManager := naturalLanguageClassifier.client.TokenManager

    path = strings.Replace(path, "{classifier_id}", classifierID, 1)

    request := req.New().Post(creds.ServiceURL + path).
        Set("Accept", "application/json").
        Query("version=" + creds.Version)

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

func GetClassifyResult(response *watson.WatsonResponse) *Classification {
    result, ok := response.Result.(*Classification)

    if ok {
        return result
    }

    return nil
}

// Classify multiple phrases
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) ClassifyCollection(classifierID string, body *ClassifyCollectionInput) (*watson.WatsonResponse, []error) {
    path := "/v1/classifiers/{classifier_id}/classify_collection"
    creds := naturalLanguageClassifier.client.Creds
    useTM := naturalLanguageClassifier.client.UseTM
    tokenManager := naturalLanguageClassifier.client.TokenManager

    path = strings.Replace(path, "{classifier_id}", classifierID, 1)

    request := req.New().Post(creds.ServiceURL + path).
        Set("Accept", "application/json").
        Query("version=" + creds.Version)

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

func GetClassifyCollectionResult(response *watson.WatsonResponse) *ClassificationCollection {
    result, ok := response.Result.(*ClassificationCollection)

    if ok {
        return result
    }

    return nil
}

// Create classifier
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) CreateClassifier(metadata os.File, trainingData os.File) (*watson.WatsonResponse, []error) {
    path := "/v1/classifiers"
    creds := naturalLanguageClassifier.client.Creds
    useTM := naturalLanguageClassifier.client.UseTM
    tokenManager := naturalLanguageClassifier.client.TokenManager


    request := req.New().Post(creds.ServiceURL + path).
        Set("Accept", "application/json").
        Query("version=" + creds.Version)


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

func GetCreateClassifierResult(response *watson.WatsonResponse) *Classifier {
    result, ok := response.Result.(*Classifier)

    if ok {
        return result
    }

    return nil
}

// Delete classifier
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) DeleteClassifier(classifierID string) (*watson.WatsonResponse, []error) {
    path := "/v1/classifiers/{classifier_id}"
    creds := naturalLanguageClassifier.client.Creds
    useTM := naturalLanguageClassifier.client.UseTM
    tokenManager := naturalLanguageClassifier.client.TokenManager

    path = strings.Replace(path, "{classifier_id}", classifierID, 1)

    request := req.New().Delete(creds.ServiceURL + path).
        Set("Accept", "application/json").
        Query("version=" + creds.Version)


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


// Get information about a classifier
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) GetClassifier(classifierID string) (*watson.WatsonResponse, []error) {
    path := "/v1/classifiers/{classifier_id}"
    creds := naturalLanguageClassifier.client.Creds
    useTM := naturalLanguageClassifier.client.UseTM
    tokenManager := naturalLanguageClassifier.client.TokenManager

    path = strings.Replace(path, "{classifier_id}", classifierID, 1)

    request := req.New().Get(creds.ServiceURL + path).
        Set("Accept", "application/json").
        Query("version=" + creds.Version)


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

func GetGetClassifierResult(response *watson.WatsonResponse) *Classifier {
    result, ok := response.Result.(*Classifier)

    if ok {
        return result
    }

    return nil
}

// List classifiers
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) ListClassifiers() (*watson.WatsonResponse, []error) {
    path := "/v1/classifiers"
    creds := naturalLanguageClassifier.client.Creds
    useTM := naturalLanguageClassifier.client.UseTM
    tokenManager := naturalLanguageClassifier.client.TokenManager


    request := req.New().Get(creds.ServiceURL + path).
        Set("Accept", "application/json").
        Query("version=" + creds.Version)


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

func GetListClassifiersResult(response *watson.WatsonResponse) *ClassifierList {
    result, ok := response.Result.(*ClassifierList)

    if ok {
        return result
    }

    return nil
}


type Classification struct {

	// Unique identifier for this classifier.
	ClassifierId string `json:"classifier_id,omitempty"`

	// Link to the classifier.
	Url string `json:"url,omitempty"`

	// The submitted phrase.
	Text string `json:"text,omitempty"`

	// The class with the highest confidence.
	TopClass string `json:"top_class,omitempty"`

	// An array of up to ten class-confidence pairs sorted in descending order of confidence.
	Classes []ClassifiedClass `json:"classes,omitempty"`
}

type ClassificationCollection struct {

	// Unique identifier for this classifier.
	ClassifierId string `json:"classifier_id,omitempty"`

	// Link to the classifier.
	Url string `json:"url,omitempty"`

	// An array of classifier responses for each submitted phrase.
	Collection []CollectionItem `json:"collection,omitempty"`
}

type ClassifiedClass struct {

	// A decimal percentage that represents the confidence that Watson has in this class. Higher values represent higher confidences.
	Confidence float64 `json:"confidence,omitempty"`

	// Class label.
	ClassName string `json:"class_name,omitempty"`
}

type Classifier struct {

	// User-supplied name for the classifier.
	Name string `json:"name,omitempty"`

	// Link to the classifier.
	Url string `json:"url"`

	// The state of the classifier.
	Status string `json:"status,omitempty"`

	// Unique identifier for this classifier.
	ClassifierId string `json:"classifier_id"`

	// Date and time (UTC) the classifier was created.
	Created strfmt.DateTime `json:"created,omitempty"`

	// Additional detail about the status.
	StatusDescription string `json:"status_description,omitempty"`

	// The language used for the classifier.
	Language string `json:"language,omitempty"`
}

type ClassifierList struct {

	// The classifiers available to the user. Returns an empty array if no classifiers are available.
	Classifiers []Classifier `json:"classifiers"`
}

type ClassifyCollectionInput struct {

	// The submitted phrases.
	Collection []ClassifyInput `json:"collection"`
}

type ClassifyInput struct {

	// The submitted phrase.
	Text string `json:"text"`
}

type CollectionItem struct {

	// The submitted phrase.
	Text string `json:"text,omitempty"`

	// The class with the highest confidence.
	TopClass string `json:"top_class,omitempty"`

	// An array of up to ten class-confidence pairs sorted in descending order of confidence.
	Classes []ClassifiedClass `json:"classes,omitempty"`
}
