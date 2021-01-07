/**
 * (C) Copyright IBM Corp. 2018, 2020.
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

/*
 * IBM OpenAPI SDK Code Generator Version: 99-SNAPSHOT-9dacd99b-20201204-091925
 */

// Package naturallanguageclassifierv1 : Operations and models for the NaturalLanguageClassifierV1 service
package naturallanguageclassifierv1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	common "github.com/watson-developer-cloud/go-sdk/common"
	"io"
	"net/http"
	"reflect"
	"time"
)

// NaturalLanguageClassifierV1 : IBM Watson&trade; Natural Language Classifier uses machine learning algorithms to
// return the top matching predefined classes for short text input. You create and train a classifier to connect
// predefined classes to example texts so that the service can apply those classes to new inputs.
//
// Version: 1.0
// See: https://cloud.ibm.com/docs/natural-language-classifier
type NaturalLanguageClassifierV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.us-south.natural-language-classifier.watson.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "natural_language_classifier"

// NaturalLanguageClassifierV1Options : Service options
type NaturalLanguageClassifierV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewNaturalLanguageClassifierV1 : constructs an instance of NaturalLanguageClassifierV1 with passed in options.
func NewNaturalLanguageClassifierV1(options *NaturalLanguageClassifierV1Options) (service *NaturalLanguageClassifierV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	if serviceOptions.Authenticator == nil {
		serviceOptions.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	err = baseService.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &NaturalLanguageClassifierV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "naturalLanguageClassifier" suitable for processing requests.
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) Clone() *NaturalLanguageClassifierV1 {
	if core.IsNil(naturalLanguageClassifier) {
		return nil
	}
	clone := *naturalLanguageClassifier
	clone.Service = naturalLanguageClassifier.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) SetServiceURL(url string) error {
	return naturalLanguageClassifier.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) GetServiceURL() string {
	return naturalLanguageClassifier.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) SetDefaultHeaders(headers http.Header) {
	naturalLanguageClassifier.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) SetEnableGzipCompression(enableGzip bool) {
	naturalLanguageClassifier.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) GetEnableGzipCompression() bool {
	return naturalLanguageClassifier.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	naturalLanguageClassifier.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) DisableRetries() {
	naturalLanguageClassifier.Service.DisableRetries()
}

// DisableSSLVerification bypasses verification of the server's SSL certificate
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) DisableSSLVerification() {
	naturalLanguageClassifier.Service.DisableSSLVerification()
}

// Classify : Classify a phrase
// Returns label information for the input. The status must be `Available` before you can use the classifier to classify
// text.
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) Classify(classifyOptions *ClassifyOptions) (result *Classification, response *core.DetailedResponse, err error) {
	return naturalLanguageClassifier.ClassifyWithContext(context.Background(), classifyOptions)
}

// ClassifyWithContext is an alternate form of the Classify method which supports a Context parameter
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) ClassifyWithContext(ctx context.Context, classifyOptions *ClassifyOptions) (result *Classification, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(classifyOptions, "classifyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(classifyOptions, "classifyOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"classifier_id": *classifyOptions.ClassifierID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = naturalLanguageClassifier.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(naturalLanguageClassifier.Service.Options.URL, `/v1/classifiers/{classifier_id}/classify`, pathParamsMap)
	if err != nil {
		return
	}

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
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = naturalLanguageClassifier.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalClassification)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ClassifyCollection : Classify multiple phrases
// Returns label information for multiple phrases. The status must be `Available` before you can use the classifier to
// classify text.
//
// Note that classifying Japanese texts is a beta feature.
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) ClassifyCollection(classifyCollectionOptions *ClassifyCollectionOptions) (result *ClassificationCollection, response *core.DetailedResponse, err error) {
	return naturalLanguageClassifier.ClassifyCollectionWithContext(context.Background(), classifyCollectionOptions)
}

// ClassifyCollectionWithContext is an alternate form of the ClassifyCollection method which supports a Context parameter
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) ClassifyCollectionWithContext(ctx context.Context, classifyCollectionOptions *ClassifyCollectionOptions) (result *ClassificationCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(classifyCollectionOptions, "classifyCollectionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(classifyCollectionOptions, "classifyCollectionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"classifier_id": *classifyCollectionOptions.ClassifierID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = naturalLanguageClassifier.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(naturalLanguageClassifier.Service.Options.URL, `/v1/classifiers/{classifier_id}/classify_collection`, pathParamsMap)
	if err != nil {
		return
	}

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
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = naturalLanguageClassifier.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalClassificationCollection)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CreateClassifier : Create classifier
// Sends data to create and train a classifier and returns information about the new classifier.
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) CreateClassifier(createClassifierOptions *CreateClassifierOptions) (result *Classifier, response *core.DetailedResponse, err error) {
	return naturalLanguageClassifier.CreateClassifierWithContext(context.Background(), createClassifierOptions)
}

// CreateClassifierWithContext is an alternate form of the CreateClassifier method which supports a Context parameter
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) CreateClassifierWithContext(ctx context.Context, createClassifierOptions *CreateClassifierOptions) (result *Classifier, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createClassifierOptions, "createClassifierOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createClassifierOptions, "createClassifierOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = naturalLanguageClassifier.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(naturalLanguageClassifier.Service.Options.URL, `/v1/classifiers`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createClassifierOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("natural_language_classifier", "V1", "CreateClassifier")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddFormData("training_metadata", "filename",
		"application/json", createClassifierOptions.TrainingMetadata)
	builder.AddFormData("training_data", "filename",
		"text/csv", createClassifierOptions.TrainingData)

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = naturalLanguageClassifier.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalClassifier)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ListClassifiers : List classifiers
// Returns an empty array if no classifiers are available.
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) ListClassifiers(listClassifiersOptions *ListClassifiersOptions) (result *ClassifierList, response *core.DetailedResponse, err error) {
	return naturalLanguageClassifier.ListClassifiersWithContext(context.Background(), listClassifiersOptions)
}

// ListClassifiersWithContext is an alternate form of the ListClassifiers method which supports a Context parameter
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) ListClassifiersWithContext(ctx context.Context, listClassifiersOptions *ListClassifiersOptions) (result *ClassifierList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listClassifiersOptions, "listClassifiersOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = naturalLanguageClassifier.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(naturalLanguageClassifier.Service.Options.URL, `/v1/classifiers`, nil)
	if err != nil {
		return
	}

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
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = naturalLanguageClassifier.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalClassifierList)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetClassifier : Get information about a classifier
// Returns status and other information about a classifier.
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) GetClassifier(getClassifierOptions *GetClassifierOptions) (result *Classifier, response *core.DetailedResponse, err error) {
	return naturalLanguageClassifier.GetClassifierWithContext(context.Background(), getClassifierOptions)
}

// GetClassifierWithContext is an alternate form of the GetClassifier method which supports a Context parameter
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) GetClassifierWithContext(ctx context.Context, getClassifierOptions *GetClassifierOptions) (result *Classifier, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getClassifierOptions, "getClassifierOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getClassifierOptions, "getClassifierOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"classifier_id": *getClassifierOptions.ClassifierID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = naturalLanguageClassifier.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(naturalLanguageClassifier.Service.Options.URL, `/v1/classifiers/{classifier_id}`, pathParamsMap)
	if err != nil {
		return
	}

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
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = naturalLanguageClassifier.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalClassifier)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteClassifier : Delete classifier
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) DeleteClassifier(deleteClassifierOptions *DeleteClassifierOptions) (response *core.DetailedResponse, err error) {
	return naturalLanguageClassifier.DeleteClassifierWithContext(context.Background(), deleteClassifierOptions)
}

// DeleteClassifierWithContext is an alternate form of the DeleteClassifier method which supports a Context parameter
func (naturalLanguageClassifier *NaturalLanguageClassifierV1) DeleteClassifierWithContext(ctx context.Context, deleteClassifierOptions *DeleteClassifierOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteClassifierOptions, "deleteClassifierOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteClassifierOptions, "deleteClassifierOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"classifier_id": *deleteClassifierOptions.ClassifierID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = naturalLanguageClassifier.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(naturalLanguageClassifier.Service.Options.URL, `/v1/classifiers/{classifier_id}`, pathParamsMap)
	if err != nil {
		return
	}

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
		return
	}

	response, err = naturalLanguageClassifier.Service.Request(request, nil)

	return
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

// UnmarshalClassification unmarshals an instance of Classification from the specified map of raw messages.
func UnmarshalClassification(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Classification)
	err = core.UnmarshalPrimitive(m, "classifier_id", &obj.ClassifierID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "top_class", &obj.TopClass)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "classes", &obj.Classes, UnmarshalClassifiedClass)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
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

// UnmarshalClassificationCollection unmarshals an instance of ClassificationCollection from the specified map of raw messages.
func UnmarshalClassificationCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ClassificationCollection)
	err = core.UnmarshalPrimitive(m, "classifier_id", &obj.ClassifierID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "collection", &obj.Collection, UnmarshalCollectionItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ClassifiedClass : Class and confidence.
type ClassifiedClass struct {
	// A decimal percentage that represents the confidence that Watson has in this class. Higher values represent higher
	// confidences.
	Confidence *float64 `json:"confidence,omitempty"`

	// Class label.
	ClassName *string `json:"class_name,omitempty"`
}

// UnmarshalClassifiedClass unmarshals an instance of ClassifiedClass from the specified map of raw messages.
func UnmarshalClassifiedClass(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ClassifiedClass)
	err = core.UnmarshalPrimitive(m, "confidence", &obj.Confidence)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "class_name", &obj.ClassName)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
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
	ClassifierStatusAvailableConst   = "Available"
	ClassifierStatusFailedConst      = "Failed"
	ClassifierStatusNonExistentConst = "Non Existent"
	ClassifierStatusTrainingConst    = "Training"
	ClassifierStatusUnavailableConst = "Unavailable"
)

// UnmarshalClassifier unmarshals an instance of Classifier from the specified map of raw messages.
func UnmarshalClassifier(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Classifier)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "classifier_id", &obj.ClassifierID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created", &obj.Created)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status_description", &obj.StatusDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "language", &obj.Language)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ClassifierList : List of available classifiers.
type ClassifierList struct {
	// The classifiers available to the user. Returns an empty array if no classifiers are available.
	Classifiers []Classifier `json:"classifiers" validate:"required"`
}

// UnmarshalClassifierList unmarshals an instance of ClassifierList from the specified map of raw messages.
func UnmarshalClassifierList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ClassifierList)
	err = core.UnmarshalModel(m, "classifiers", &obj.Classifiers, UnmarshalClassifier)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ClassifyCollectionOptions : The ClassifyCollection options.
type ClassifyCollectionOptions struct {
	// Classifier ID to use.
	ClassifierID *string `json:"classifier_id" validate:"required,ne="`

	// The submitted phrases.
	Collection []ClassifyInput `json:"collection" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewClassifyCollectionOptions : Instantiate ClassifyCollectionOptions
func (*NaturalLanguageClassifierV1) NewClassifyCollectionOptions(classifierID string, collection []ClassifyInput) *ClassifyCollectionOptions {
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

// NewClassifyInput : Instantiate ClassifyInput (Generic Model Constructor)
func (*NaturalLanguageClassifierV1) NewClassifyInput(text string) (model *ClassifyInput, err error) {
	model = &ClassifyInput{
		Text: core.StringPtr(text),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalClassifyInput unmarshals an instance of ClassifyInput from the specified map of raw messages.
func UnmarshalClassifyInput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ClassifyInput)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ClassifyOptions : The Classify options.
type ClassifyOptions struct {
	// Classifier ID to use.
	ClassifierID *string `json:"classifier_id" validate:"required,ne="`

	// The submitted phrase. The maximum length is 2048 characters.
	Text *string `json:"text" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewClassifyOptions : Instantiate ClassifyOptions
func (*NaturalLanguageClassifierV1) NewClassifyOptions(classifierID string, text string) *ClassifyOptions {
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

// UnmarshalCollectionItem unmarshals an instance of CollectionItem from the specified map of raw messages.
func UnmarshalCollectionItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CollectionItem)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "top_class", &obj.TopClass)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "classes", &obj.Classes, UnmarshalClassifiedClass)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateClassifierOptions : The CreateClassifier options.
type CreateClassifierOptions struct {
	// Metadata in JSON format. The metadata identifies the language of the data, and an optional name to identify the
	// classifier. Specify the language with the 2-letter primary language code as assigned in ISO standard 639.
	//
	// Supported languages are English (`en`), Arabic (`ar`), French (`fr`), German, (`de`), Italian (`it`), Japanese
	// (`ja`), Korean (`ko`), Brazilian Portuguese (`pt`), and Spanish (`es`).
	TrainingMetadata io.ReadCloser `json:"training_metadata" validate:"required"`

	// Training data in CSV format. Each text value must have at least one class. The data can include up to 3,000 classes
	// and 20,000 records. For details, see [Data
	// preparation](https://cloud.ibm.com/docs/natural-language-classifier?topic=natural-language-classifier-using-your-data).
	TrainingData io.ReadCloser `json:"training_data" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateClassifierOptions : Instantiate CreateClassifierOptions
func (*NaturalLanguageClassifierV1) NewCreateClassifierOptions(trainingMetadata io.ReadCloser, trainingData io.ReadCloser) *CreateClassifierOptions {
	return &CreateClassifierOptions{
		TrainingMetadata: trainingMetadata,
		TrainingData:     trainingData,
	}
}

// SetTrainingMetadata : Allow user to set TrainingMetadata
func (options *CreateClassifierOptions) SetTrainingMetadata(trainingMetadata io.ReadCloser) *CreateClassifierOptions {
	options.TrainingMetadata = trainingMetadata
	return options
}

// SetTrainingData : Allow user to set TrainingData
func (options *CreateClassifierOptions) SetTrainingData(trainingData io.ReadCloser) *CreateClassifierOptions {
	options.TrainingData = trainingData
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateClassifierOptions) SetHeaders(param map[string]string) *CreateClassifierOptions {
	options.Headers = param
	return options
}

// DeleteClassifierOptions : The DeleteClassifier options.
type DeleteClassifierOptions struct {
	// Classifier ID to delete.
	ClassifierID *string `json:"classifier_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteClassifierOptions : Instantiate DeleteClassifierOptions
func (*NaturalLanguageClassifierV1) NewDeleteClassifierOptions(classifierID string) *DeleteClassifierOptions {
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

// GetClassifierOptions : The GetClassifier options.
type GetClassifierOptions struct {
	// Classifier ID to query.
	ClassifierID *string `json:"classifier_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetClassifierOptions : Instantiate GetClassifierOptions
func (*NaturalLanguageClassifierV1) NewGetClassifierOptions(classifierID string) *GetClassifierOptions {
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

// ListClassifiersOptions : The ListClassifiers options.
type ListClassifiersOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListClassifiersOptions : Instantiate ListClassifiersOptions
func (*NaturalLanguageClassifierV1) NewListClassifiersOptions() *ListClassifiersOptions {
	return &ListClassifiersOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListClassifiersOptions) SetHeaders(param map[string]string) *ListClassifiersOptions {
	options.Headers = param
	return options
}
