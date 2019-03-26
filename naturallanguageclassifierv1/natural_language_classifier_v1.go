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
	"github.com/IBM/go-sdk-core/core"
	"github.com/go-openapi/strfmt"
	common "github.com/watson-developer-cloud/go-sdk/common"
	"os"
)

// NaturalLanguageClassifierV1 : IBM Watson&trade; Natural Language Classifier uses machine learning algorithms to
// return the top matching predefined classes for short text input. You create and train a classifier to connect
// predefined classes to example texts so that the service can apply those classes to new inputs.
//
// Version: V1
// See: http://www.ibm.com/watson/developercloud/natural-language-classifier.html
type NaturalLanguageClassifierV1 struct {
	Service *core.BaseService
}

// NaturalLanguageClassifierV1Options : Service options
type NaturalLanguageClassifierV1Options struct {
	URL            string
	Username       string
	Password       string
	IAMApiKey      string
	IAMAccessToken string
	IAMURL         string
}

// NewNaturalLanguageClassifierV1 : Instantiate NaturalLanguageClassifierV1
func NewNaturalLanguageClassifierV1(options *NaturalLanguageClassifierV1Options) (*NaturalLanguageClassifierV1, error) {
	if options.URL == "" {
		options.URL = "https://gateway.watsonplatform.net/natural-language-classifier/api"
	}

	serviceOptions := &core.ServiceOptions{
		URL:            options.URL,
		Username:       options.Username,
		Password:       options.Password,
		IAMApiKey:      options.IAMApiKey,
		IAMAccessToken: options.IAMAccessToken,
		IAMURL:         options.IAMURL,
	}
	service, serviceErr := core.NewBaseService(serviceOptions, "natural_language_classifier", "Natural Language Classifier")
	if serviceErr != nil {
		return nil, serviceErr
	}

	return &NaturalLanguageClassifierV1{Service: service}, nil
}

// Classify : Classify a phrase
// Returns label information for the input. The status must be `Available` before you can use the classifier to classify
// text.
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) Classify(classifyOptions *ClassifyOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(classifyOptions, "classifyOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(classifyOptions, "classifyOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/classifiers", "classify"}
	pathParameters := []string{*classifyOptions.ClassifierID}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(naturalLanguageClassifier.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range classifyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("natural_language_classifier", "V1", "Classify")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if classifyOptions.Text != nil {
		body["text"] = classifyOptions.Text
	}
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := naturalLanguageClassifier.Service.Request(request, new(Classification))
	return response, err
}

// GetClassifyResult : Retrieve result of Classify operation
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) GetClassifyResult(response *core.DetailedResponse) *Classification {
	result, ok := response.Result.(*Classification)
	if ok {
		return result
	}
	return nil
}

// ClassifyCollection : Classify multiple phrases
// Returns label information for multiple phrases. The status must be `Available` before you can use the classifier to
// classify text.
//
// Note that classifying Japanese texts is a beta feature.
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) ClassifyCollection(classifyCollectionOptions *ClassifyCollectionOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(classifyCollectionOptions, "classifyCollectionOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(classifyCollectionOptions, "classifyCollectionOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/classifiers", "classify_collection"}
	pathParameters := []string{*classifyCollectionOptions.ClassifierID}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(naturalLanguageClassifier.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range classifyCollectionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("natural_language_classifier", "V1", "ClassifyCollection")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if classifyCollectionOptions.Collection != nil {
		body["collection"] = classifyCollectionOptions.Collection
	}
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := naturalLanguageClassifier.Service.Request(request, new(ClassificationCollection))
	return response, err
}

// GetClassifyCollectionResult : Retrieve result of ClassifyCollection operation
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) GetClassifyCollectionResult(response *core.DetailedResponse) *ClassificationCollection {
	result, ok := response.Result.(*ClassificationCollection)
	if ok {
		return result
	}
	return nil
}

// CreateClassifier : Create classifier
// Sends data to create and train a classifier and returns information about the new classifier.
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) CreateClassifier(createClassifierOptions *CreateClassifierOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(createClassifierOptions, "createClassifierOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(createClassifierOptions, "createClassifierOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/classifiers"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(naturalLanguageClassifier.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range createClassifierOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("natural_language_classifier", "V1", "CreateClassifier")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")

	builder.AddFormData("training_metadata", "filename",
		"application/json", createClassifierOptions.Metadata)
	builder.AddFormData("training_data", "filename",
		"text/csv", createClassifierOptions.TrainingData)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := naturalLanguageClassifier.Service.Request(request, new(Classifier))
	return response, err
}

// GetCreateClassifierResult : Retrieve result of CreateClassifier operation
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) GetCreateClassifierResult(response *core.DetailedResponse) *Classifier {
	result, ok := response.Result.(*Classifier)
	if ok {
		return result
	}
	return nil
}

// DeleteClassifier : Delete classifier
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) DeleteClassifier(deleteClassifierOptions *DeleteClassifierOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(deleteClassifierOptions, "deleteClassifierOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(deleteClassifierOptions, "deleteClassifierOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/classifiers"}
	pathParameters := []string{*deleteClassifierOptions.ClassifierID}

	builder := core.NewRequestBuilder(core.DELETE)
	builder.ConstructHTTPURL(naturalLanguageClassifier.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteClassifierOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("natural_language_classifier", "V1", "DeleteClassifier")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := naturalLanguageClassifier.Service.Request(request, nil)
	return response, err
}

// GetClassifier : Get information about a classifier
// Returns status and other information about a classifier.
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) GetClassifier(getClassifierOptions *GetClassifierOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(getClassifierOptions, "getClassifierOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(getClassifierOptions, "getClassifierOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/classifiers"}
	pathParameters := []string{*getClassifierOptions.ClassifierID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(naturalLanguageClassifier.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range getClassifierOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("natural_language_classifier", "V1", "GetClassifier")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := naturalLanguageClassifier.Service.Request(request, new(Classifier))
	return response, err
}

// GetGetClassifierResult : Retrieve result of GetClassifier operation
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) GetGetClassifierResult(response *core.DetailedResponse) *Classifier {
	result, ok := response.Result.(*Classifier)
	if ok {
		return result
	}
	return nil
}

// ListClassifiers : List classifiers
// Returns an empty array if no classifiers are available.
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) ListClassifiers(listClassifiersOptions *ListClassifiersOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateStruct(listClassifiersOptions, "listClassifiersOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/classifiers"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(naturalLanguageClassifier.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range listClassifiersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("natural_language_classifier", "V1", "ListClassifiers")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := naturalLanguageClassifier.Service.Request(request, new(ClassifierList))
	return response, err
}

// GetListClassifiersResult : Retrieve result of ListClassifiers operation
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) GetListClassifiersResult(response *core.DetailedResponse) *ClassifierList {
	result, ok := response.Result.(*ClassifierList)
	if ok {
		return result
	}
	return nil
}

// Classification : Response from the classifier for a phrase.
type Classification struct {

	// Unique identifier for this classifier.
	ClassifierID *string `json:"classifier_id,omitempty"`

	// Link to the classifier.
	URL *string `json:"url,omitempty"`

	// The submitted phrase.
	Text *string `json:"text,omitempty"`

	// The class with the highest confidence.
	TopClass *string `json:"top_class,omitempty"`

	// An array of up to ten class-confidence pairs sorted in descending order of confidence.
	Classes []ClassifiedClass `json:"classes,omitempty"`
}

// ClassificationCollection : Response from the classifier for multiple phrases.
type ClassificationCollection struct {

	// Unique identifier for this classifier.
	ClassifierID *string `json:"classifier_id,omitempty"`

	// Link to the classifier.
	URL *string `json:"url,omitempty"`

	// An array of classifier responses for each submitted phrase.
	Collection []CollectionItem `json:"collection,omitempty"`
}

// ClassifiedClass : Class and confidence.
type ClassifiedClass struct {

	// A decimal percentage that represents the confidence that Watson has in this class. Higher values represent higher
	// confidences.
	Confidence *float64 `json:"confidence,omitempty"`

	// Class label.
	ClassName *string `json:"class_name,omitempty"`
}

// Classifier : A classifier for natural language phrases.
type Classifier struct {

	// User-supplied name for the classifier.
	Name *string `json:"name,omitempty"`

	// Link to the classifier.
	URL *string `json:"url" validate:"required"`

	// The state of the classifier.
	Status *string `json:"status,omitempty"`

	// Unique identifier for this classifier.
	ClassifierID *string `json:"classifier_id" validate:"required"`

	// Date and time (UTC) the classifier was created.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// Additional detail about the status.
	StatusDescription *string `json:"status_description,omitempty"`

	// The language used for the classifier.
	Language *string `json:"language,omitempty"`
}

// Constants associated with the Classifier.Status property.
// The state of the classifier.
const (
	Classifier_Status_Available   = "Available"
	Classifier_Status_Failed      = "Failed"
	Classifier_Status_NonExistent = "Non Existent"
	Classifier_Status_Training    = "Training"
	Classifier_Status_Unavailable = "Unavailable"
)

// ClassifierList : List of available classifiers.
type ClassifierList struct {

	// The classifiers available to the user. Returns an empty array if no classifiers are available.
	Classifiers []Classifier `json:"classifiers" validate:"required"`
}

// ClassifyCollectionOptions : The classifyCollection options.
type ClassifyCollectionOptions struct {

	// Classifier ID to use.
	ClassifierID *string `json:"classifier_id" validate:"required"`

	// The submitted phrases.
	Collection []ClassifyInput `json:"collection" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewClassifyCollectionOptions : Instantiate ClassifyCollectionOptions
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) NewClassifyCollectionOptions(classifierID string, collection []ClassifyInput) *ClassifyCollectionOptions {
	return &ClassifyCollectionOptions{
		ClassifierID: core.StringPtr(classifierID),
		Collection:   collection,
	}
}

// SetClassifierID : Allow user to set ClassifierID
func (options *ClassifyCollectionOptions) SetClassifierID(classifierID string) *ClassifyCollectionOptions {
	options.ClassifierID = core.StringPtr(classifierID)
	return options
}

// SetCollection : Allow user to set Collection
func (options *ClassifyCollectionOptions) SetCollection(collection []ClassifyInput) *ClassifyCollectionOptions {
	options.Collection = collection
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ClassifyCollectionOptions) SetHeaders(param map[string]string) *ClassifyCollectionOptions {
	options.Headers = param
	return options
}

// ClassifyInput : Request payload to classify.
type ClassifyInput struct {

	// The submitted phrase. The maximum length is 2048 characters.
	Text *string `json:"text" validate:"required"`
}

// ClassifyOptions : The classify options.
type ClassifyOptions struct {

	// Classifier ID to use.
	ClassifierID *string `json:"classifier_id" validate:"required"`

	// The submitted phrase. The maximum length is 2048 characters.
	Text *string `json:"text" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewClassifyOptions : Instantiate ClassifyOptions
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) NewClassifyOptions(classifierID string, text string) *ClassifyOptions {
	return &ClassifyOptions{
		ClassifierID: core.StringPtr(classifierID),
		Text:         core.StringPtr(text),
	}
}

// SetClassifierID : Allow user to set ClassifierID
func (options *ClassifyOptions) SetClassifierID(classifierID string) *ClassifyOptions {
	options.ClassifierID = core.StringPtr(classifierID)
	return options
}

// SetText : Allow user to set Text
func (options *ClassifyOptions) SetText(text string) *ClassifyOptions {
	options.Text = core.StringPtr(text)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ClassifyOptions) SetHeaders(param map[string]string) *ClassifyOptions {
	options.Headers = param
	return options
}

// CollectionItem : Response from the classifier for a phrase in a collection.
type CollectionItem struct {

	// The submitted phrase. The maximum length is 2048 characters.
	Text *string `json:"text,omitempty"`

	// The class with the highest confidence.
	TopClass *string `json:"top_class,omitempty"`

	// An array of up to ten class-confidence pairs sorted in descending order of confidence.
	Classes []ClassifiedClass `json:"classes,omitempty"`
}

// CreateClassifierOptions : The createClassifier options.
type CreateClassifierOptions struct {

	// Metadata in JSON format. The metadata identifies the language of the data, and an optional name to identify the
	// classifier. Specify the language with the 2-letter primary language code as assigned in ISO standard 639.
	//
	// Supported languages are English (`en`), Arabic (`ar`), French (`fr`), German, (`de`), Italian (`it`), Japanese
	// (`ja`), Korean (`ko`), Brazilian Portuguese (`pt`), and Spanish (`es`).
	Metadata *os.File `json:"training_metadata" validate:"required"`

	// Training data in CSV format. Each text value must have at least one class. The data can include up to 3,000 classes
	// and 20,000 records. For details, see [Data
	// preparation](https://cloud.ibm.com/docs/services/natural-language-classifier/using-your-data.html).
	TrainingData *os.File `json:"training_data" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewCreateClassifierOptions : Instantiate CreateClassifierOptions
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) NewCreateClassifierOptions(metadata *os.File, trainingData *os.File) *CreateClassifierOptions {
	return &CreateClassifierOptions{
		Metadata:     metadata,
		TrainingData: trainingData,
	}
}

// SetMetadata : Allow user to set Metadata
func (options *CreateClassifierOptions) SetMetadata(metadata *os.File) *CreateClassifierOptions {
	options.Metadata = metadata
	return options
}

// SetTrainingData : Allow user to set TrainingData
func (options *CreateClassifierOptions) SetTrainingData(trainingData *os.File) *CreateClassifierOptions {
	options.TrainingData = trainingData
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
	ClassifierID *string `json:"classifier_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteClassifierOptions : Instantiate DeleteClassifierOptions
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) NewDeleteClassifierOptions(classifierID string) *DeleteClassifierOptions {
	return &DeleteClassifierOptions{
		ClassifierID: core.StringPtr(classifierID),
	}
}

// SetClassifierID : Allow user to set ClassifierID
func (options *DeleteClassifierOptions) SetClassifierID(classifierID string) *DeleteClassifierOptions {
	options.ClassifierID = core.StringPtr(classifierID)
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
	ClassifierID *string `json:"classifier_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetClassifierOptions : Instantiate GetClassifierOptions
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) NewGetClassifierOptions(classifierID string) *GetClassifierOptions {
	return &GetClassifierOptions{
		ClassifierID: core.StringPtr(classifierID),
	}
}

// SetClassifierID : Allow user to set ClassifierID
func (options *GetClassifierOptions) SetClassifierID(classifierID string) *GetClassifierOptions {
	options.ClassifierID = core.StringPtr(classifierID)
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
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) NewListClassifiersOptions() *ListClassifiersOptions {
	return &ListClassifiersOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListClassifiersOptions) SetHeaders(param map[string]string) *ListClassifiersOptions {
	options.Headers = param
	return options
}
