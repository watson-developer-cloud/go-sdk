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

// Package visualrecognitionv3 : Operations and models for the VisualRecognitionV3 service
package visualrecognitionv3

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	common "github.com/watson-developer-cloud/go-sdk/common"
)

// VisualRecognitionV3 : IBM Watson&trade; Visual Recognition is discontinued. Existing instances are supported until 1
// December 2021, but as of 7 January 2021, you can't create instances. Any instance that is provisioned on 1 December
// 2021 will be deleted.
// {: deprecated}
//
// The IBM Watson Visual Recognition service uses deep learning algorithms to identify scenes and objects in images that
// you upload to the service. You can create and train a custom classifier to identify subjects that suit your needs.
//
// Version: 3.0
// See: https://cloud.ibm.com/docs/visual-recognition
type VisualRecognitionV3 struct {
	Service *core.BaseService

	// Release date of the API version you want to use. Specify dates in YYYY-MM-DD format. The current version is
	// `2018-03-19`.
	Version *string
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.us-south.visual-recognition.watson.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "watson_vision_combined"

// VisualRecognitionV3Options : Service options
type VisualRecognitionV3Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator

	// Release date of the API version you want to use. Specify dates in YYYY-MM-DD format. The current version is
	// `2018-03-19`.
	Version *string `validate:"required"`
}

// NewVisualRecognitionV3 : constructs an instance of VisualRecognitionV3 with passed in options.
func NewVisualRecognitionV3(options *VisualRecognitionV3Options) (service *VisualRecognitionV3, err error) {
	// Log deprecation warning
	core.GetLogger().Log(core.LevelWarn, "", "On 1 December 2021, Visual Recognition will no longer be available. For more information, see Visual Recognition Deprecation at https://github.com/watson-developer-cloud/go-sdk/tree/master#visual-recognition-deprecation.")

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

	err = core.ValidateStruct(options, "options")
	if err != nil {
		return
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

	service = &VisualRecognitionV3{
		Service: baseService,
		Version: options.Version,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "visualRecognition" suitable for processing requests.
func (visualRecognition *VisualRecognitionV3) Clone() *VisualRecognitionV3 {
	if core.IsNil(visualRecognition) {
		return nil
	}
	clone := *visualRecognition
	clone.Service = visualRecognition.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (visualRecognition *VisualRecognitionV3) SetServiceURL(url string) error {
	return visualRecognition.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (visualRecognition *VisualRecognitionV3) GetServiceURL() string {
	return visualRecognition.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (visualRecognition *VisualRecognitionV3) SetDefaultHeaders(headers http.Header) {
	visualRecognition.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (visualRecognition *VisualRecognitionV3) SetEnableGzipCompression(enableGzip bool) {
	visualRecognition.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (visualRecognition *VisualRecognitionV3) GetEnableGzipCompression() bool {
	return visualRecognition.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (visualRecognition *VisualRecognitionV3) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	visualRecognition.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (visualRecognition *VisualRecognitionV3) DisableRetries() {
	visualRecognition.Service.DisableRetries()
}

// DisableSSLVerification bypasses verification of the server's SSL certificate
func (visualRecognition *VisualRecognitionV3) DisableSSLVerification() {
	visualRecognition.Service.DisableSSLVerification()
}

// Classify : Classify images
// Classify images with built-in or custom classifiers.
func (visualRecognition *VisualRecognitionV3) Classify(classifyOptions *ClassifyOptions) (result *ClassifiedImages, response *core.DetailedResponse, err error) {
	return visualRecognition.ClassifyWithContext(context.Background(), classifyOptions)
}

// ClassifyWithContext is an alternate form of the Classify method which supports a Context parameter
func (visualRecognition *VisualRecognitionV3) ClassifyWithContext(ctx context.Context, classifyOptions *ClassifyOptions) (result *ClassifiedImages, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(classifyOptions, "classifyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(classifyOptions, "classifyOptions")
	if err != nil {
		return
	}
	if (classifyOptions.ImagesFile == nil) && (classifyOptions.URL == nil) && (classifyOptions.Threshold == nil) && (classifyOptions.Owners == nil) && (classifyOptions.ClassifierIds == nil) {
		err = fmt.Errorf("at least one of imagesFile, url, threshold, owners, or classifierIds must be supplied")
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = visualRecognition.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(visualRecognition.Service.Options.URL, `/v3/classify`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range classifyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_vision_combined", "V3", "Classify")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if classifyOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*classifyOptions.AcceptLanguage))
	}

	builder.AddQuery("version", fmt.Sprint(*visualRecognition.Version))

	if classifyOptions.ImagesFile != nil {
		builder.AddFormData("images_file", core.StringNilMapper(classifyOptions.ImagesFilename),
			core.StringNilMapper(classifyOptions.ImagesFileContentType), classifyOptions.ImagesFile)
	}
	if classifyOptions.URL != nil {
		builder.AddFormData("url", "", "", fmt.Sprint(*classifyOptions.URL))
	}
	if classifyOptions.Threshold != nil {
		builder.AddFormData("threshold", "", "", fmt.Sprint(*classifyOptions.Threshold))
	}
	if classifyOptions.Owners != nil {
		for _, item := range classifyOptions.Owners {
			builder.AddFormData("owners", "", "", fmt.Sprint(item))
		}
	}
	if classifyOptions.ClassifierIds != nil {
		for _, item := range classifyOptions.ClassifierIds {
			builder.AddFormData("classifier_ids", "", "", fmt.Sprint(item))
		}
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = visualRecognition.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalClassifiedImages)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CreateClassifier : Create a classifier
// Train a new multi-faceted classifier on the uploaded image data. Create your custom classifier with positive or
// negative example training images. Include at least two sets of examples, either two positive example files or one
// positive and one negative file. You can upload a maximum of 256 MB per call.
//
// **Tips when creating:**
//
// - If you set the **X-Watson-Learning-Opt-Out** header parameter to `true` when you create a classifier, the example
// training images are not stored. Save your training images locally. For more information, see [Data
// collection](#data-collection).
//
// - Encode all names in UTF-8 if they contain non-ASCII characters (.zip and image file names, and classifier and class
// names). The service assumes UTF-8 encoding if it encounters non-ASCII characters.
func (visualRecognition *VisualRecognitionV3) CreateClassifier(createClassifierOptions *CreateClassifierOptions) (result *Classifier, response *core.DetailedResponse, err error) {
	return visualRecognition.CreateClassifierWithContext(context.Background(), createClassifierOptions)
}

// CreateClassifierWithContext is an alternate form of the CreateClassifier method which supports a Context parameter
func (visualRecognition *VisualRecognitionV3) CreateClassifierWithContext(ctx context.Context, createClassifierOptions *CreateClassifierOptions) (result *Classifier, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = visualRecognition.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(visualRecognition.Service.Options.URL, `/v3/classifiers`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createClassifierOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_vision_combined", "V3", "CreateClassifier")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*visualRecognition.Version))

	builder.AddFormData("name", "", "", fmt.Sprint(*createClassifierOptions.Name))
	for key, value := range createClassifierOptions.PositiveExamples {
		partName := fmt.Sprintf("%s_positive_examples", key)
		builder.AddFormData(partName, key, "application/octet-stream", value)
	}
	if createClassifierOptions.NegativeExamples != nil {
		builder.AddFormData("negative_examples", core.StringNilMapper(createClassifierOptions.NegativeExamplesFilename),
			"application/octet-stream", createClassifierOptions.NegativeExamples)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = visualRecognition.Service.Request(request, &rawResponse)
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

// ListClassifiers : Retrieve a list of classifiers
func (visualRecognition *VisualRecognitionV3) ListClassifiers(listClassifiersOptions *ListClassifiersOptions) (result *Classifiers, response *core.DetailedResponse, err error) {
	return visualRecognition.ListClassifiersWithContext(context.Background(), listClassifiersOptions)
}

// ListClassifiersWithContext is an alternate form of the ListClassifiers method which supports a Context parameter
func (visualRecognition *VisualRecognitionV3) ListClassifiersWithContext(ctx context.Context, listClassifiersOptions *ListClassifiersOptions) (result *Classifiers, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listClassifiersOptions, "listClassifiersOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = visualRecognition.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(visualRecognition.Service.Options.URL, `/v3/classifiers`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listClassifiersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_vision_combined", "V3", "ListClassifiers")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*visualRecognition.Version))
	if listClassifiersOptions.Verbose != nil {
		builder.AddQuery("verbose", fmt.Sprint(*listClassifiersOptions.Verbose))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = visualRecognition.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalClassifiers)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetClassifier : Retrieve classifier details
// Retrieve information about a custom classifier.
func (visualRecognition *VisualRecognitionV3) GetClassifier(getClassifierOptions *GetClassifierOptions) (result *Classifier, response *core.DetailedResponse, err error) {
	return visualRecognition.GetClassifierWithContext(context.Background(), getClassifierOptions)
}

// GetClassifierWithContext is an alternate form of the GetClassifier method which supports a Context parameter
func (visualRecognition *VisualRecognitionV3) GetClassifierWithContext(ctx context.Context, getClassifierOptions *GetClassifierOptions) (result *Classifier, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = visualRecognition.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(visualRecognition.Service.Options.URL, `/v3/classifiers/{classifier_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getClassifierOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_vision_combined", "V3", "GetClassifier")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*visualRecognition.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = visualRecognition.Service.Request(request, &rawResponse)
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

// UpdateClassifier : Update a classifier
// Update a custom classifier by adding new positive or negative classes or by adding new images to existing classes.
// You must supply at least one set of positive or negative examples. For details, see [Updating custom
// classifiers](https://cloud.ibm.com/docs/visual-recognition?topic=visual-recognition-customizing#updating-custom-classifiers).
//
// Encode all names in UTF-8 if they contain non-ASCII characters (.zip and image file names, and classifier and class
// names). The service assumes UTF-8 encoding if it encounters non-ASCII characters.
//
// **Tips about retraining:**
//
// - You can't update the classifier if the **X-Watson-Learning-Opt-Out** header parameter was set to `true` when the
// classifier was created. Training images are not stored in that case. Instead, create another classifier. For more
// information, see [Data collection](#data-collection).
//
// - Don't make retraining calls on a classifier until the status is ready. When you submit retraining requests in
// parallel, the last request overwrites the previous requests. The `retrained` property shows the last time the
// classifier retraining finished.
func (visualRecognition *VisualRecognitionV3) UpdateClassifier(updateClassifierOptions *UpdateClassifierOptions) (result *Classifier, response *core.DetailedResponse, err error) {
	return visualRecognition.UpdateClassifierWithContext(context.Background(), updateClassifierOptions)
}

// UpdateClassifierWithContext is an alternate form of the UpdateClassifier method which supports a Context parameter
func (visualRecognition *VisualRecognitionV3) UpdateClassifierWithContext(ctx context.Context, updateClassifierOptions *UpdateClassifierOptions) (result *Classifier, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateClassifierOptions, "updateClassifierOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateClassifierOptions, "updateClassifierOptions")
	if err != nil {
		return
	}
	if (updateClassifierOptions.PositiveExamples == nil) && (updateClassifierOptions.NegativeExamples == nil) {
		err = fmt.Errorf("at least one of positiveExamples or negativeExamples must be supplied")
		return
	}

	pathParamsMap := map[string]string{
		"classifier_id": *updateClassifierOptions.ClassifierID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = visualRecognition.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(visualRecognition.Service.Options.URL, `/v3/classifiers/{classifier_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateClassifierOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_vision_combined", "V3", "UpdateClassifier")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*visualRecognition.Version))

	for key, value := range updateClassifierOptions.PositiveExamples {
		partName := fmt.Sprintf("%s_positive_examples", key)
		builder.AddFormData(partName, key, "application/octet-stream", value)
	}
	if updateClassifierOptions.NegativeExamples != nil {
		builder.AddFormData("negative_examples", core.StringNilMapper(updateClassifierOptions.NegativeExamplesFilename),
			"application/octet-stream", updateClassifierOptions.NegativeExamples)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = visualRecognition.Service.Request(request, &rawResponse)
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

// DeleteClassifier : Delete a classifier
func (visualRecognition *VisualRecognitionV3) DeleteClassifier(deleteClassifierOptions *DeleteClassifierOptions) (response *core.DetailedResponse, err error) {
	return visualRecognition.DeleteClassifierWithContext(context.Background(), deleteClassifierOptions)
}

// DeleteClassifierWithContext is an alternate form of the DeleteClassifier method which supports a Context parameter
func (visualRecognition *VisualRecognitionV3) DeleteClassifierWithContext(ctx context.Context, deleteClassifierOptions *DeleteClassifierOptions) (response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = visualRecognition.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(visualRecognition.Service.Options.URL, `/v3/classifiers/{classifier_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteClassifierOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_vision_combined", "V3", "DeleteClassifier")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*visualRecognition.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = visualRecognition.Service.Request(request, nil)

	return
}

// GetCoreMlModel : Retrieve a Core ML model of a classifier
// Download a Core ML model file (.mlmodel) of a custom classifier that returns <tt>"core_ml_enabled": true</tt> in the
// classifier details.
func (visualRecognition *VisualRecognitionV3) GetCoreMlModel(getCoreMlModelOptions *GetCoreMlModelOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	return visualRecognition.GetCoreMlModelWithContext(context.Background(), getCoreMlModelOptions)
}

// GetCoreMlModelWithContext is an alternate form of the GetCoreMlModel method which supports a Context parameter
func (visualRecognition *VisualRecognitionV3) GetCoreMlModelWithContext(ctx context.Context, getCoreMlModelOptions *GetCoreMlModelOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getCoreMlModelOptions, "getCoreMlModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getCoreMlModelOptions, "getCoreMlModelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"classifier_id": *getCoreMlModelOptions.ClassifierID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = visualRecognition.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(visualRecognition.Service.Options.URL, `/v3/classifiers/{classifier_id}/core_ml_model`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCoreMlModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_vision_combined", "V3", "GetCoreMlModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/octet-stream")

	builder.AddQuery("version", fmt.Sprint(*visualRecognition.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = visualRecognition.Service.Request(request, &result)

	return
}

// DeleteUserData : Delete labeled data
// Deletes all data associated with a specified customer ID. The method has no effect if no data is associated with the
// customer ID.
//
// You associate a customer ID with data by passing the `X-Watson-Metadata` header with a request that passes data. For
// more information about personal data and customer IDs, see [Information
// security](https://cloud.ibm.com/docs/visual-recognition?topic=visual-recognition-information-security).
func (visualRecognition *VisualRecognitionV3) DeleteUserData(deleteUserDataOptions *DeleteUserDataOptions) (response *core.DetailedResponse, err error) {
	return visualRecognition.DeleteUserDataWithContext(context.Background(), deleteUserDataOptions)
}

// DeleteUserDataWithContext is an alternate form of the DeleteUserData method which supports a Context parameter
func (visualRecognition *VisualRecognitionV3) DeleteUserDataWithContext(ctx context.Context, deleteUserDataOptions *DeleteUserDataOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteUserDataOptions, "deleteUserDataOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteUserDataOptions, "deleteUserDataOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = visualRecognition.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(visualRecognition.Service.Options.URL, `/v3/user_data`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteUserDataOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_vision_combined", "V3", "DeleteUserData")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*visualRecognition.Version))
	builder.AddQuery("customer_id", fmt.Sprint(*deleteUserDataOptions.CustomerID))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = visualRecognition.Service.Request(request, nil)

	return
}

// Class : A category within a classifier.
type Class struct {
	// The name of the class.
	Class *string `json:"class" validate:"required"`
}

// UnmarshalClass unmarshals an instance of Class from the specified map of raw messages.
func UnmarshalClass(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Class)
	err = core.UnmarshalPrimitive(m, "class", &obj.Class)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ClassResult : Result of a class within a classifier.
type ClassResult struct {
	// Name of the class.
	//
	// Class names are translated in the language defined by the **Accept-Language** request header for the build-in
	// classifier IDs (`default`, `food`, and `explicit`). Class names of custom classifiers are not translated. The
	// response might not be in the specified language when the requested language is not supported or when there is no
	// translation for the class name.
	Class *string `json:"class" validate:"required"`

	// Confidence score for the property in the range of 0 to 1. A higher score indicates greater likelihood that the class
	// is depicted in the image. The default threshold for returning scores from a classifier is 0.5.
	Score *float32 `json:"score" validate:"required"`

	// Knowledge graph of the property. For example, `/fruit/pome/apple/eating apple/Granny Smith`. Included only if
	// identified.
	TypeHierarchy *string `json:"type_hierarchy,omitempty"`
}

// UnmarshalClassResult unmarshals an instance of ClassResult from the specified map of raw messages.
func UnmarshalClassResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ClassResult)
	err = core.UnmarshalPrimitive(m, "class", &obj.Class)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "score", &obj.Score)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type_hierarchy", &obj.TypeHierarchy)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ClassifiedImage : Results for one image.
type ClassifiedImage struct {
	// Source of the image before any redirects. Not returned when the image is uploaded.
	SourceURL *string `json:"source_url,omitempty"`

	// Fully resolved URL of the image after redirects are followed. Not returned when the image is uploaded.
	ResolvedURL *string `json:"resolved_url,omitempty"`

	// Relative path of the image file if uploaded directly. Not returned when the image is passed by URL.
	Image *string `json:"image,omitempty"`

	// Information about what might have caused a failure, such as an image that is too large. Not returned when there is
	// no error.
	Error *ErrorInfo `json:"error,omitempty"`

	// The classifiers.
	Classifiers []ClassifierResult `json:"classifiers" validate:"required"`
}

// UnmarshalClassifiedImage unmarshals an instance of ClassifiedImage from the specified map of raw messages.
func UnmarshalClassifiedImage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ClassifiedImage)
	err = core.UnmarshalPrimitive(m, "source_url", &obj.SourceURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resolved_url", &obj.ResolvedURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "image", &obj.Image)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "error", &obj.Error, UnmarshalErrorInfo)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "classifiers", &obj.Classifiers, UnmarshalClassifierResult)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ClassifiedImages : Results for all images.
type ClassifiedImages struct {
	// Number of custom classes identified in the images.
	CustomClasses *int64 `json:"custom_classes,omitempty"`

	// Number of images processed for the API call.
	ImagesProcessed *int64 `json:"images_processed,omitempty"`

	// Classified images.
	Images []ClassifiedImage `json:"images" validate:"required"`

	// Information about what might cause less than optimal output. For example, a request sent with a corrupt .zip file
	// and a list of image URLs will still complete, but does not return the expected output. Not returned when there is no
	// warning.
	Warnings []WarningInfo `json:"warnings,omitempty"`
}

// UnmarshalClassifiedImages unmarshals an instance of ClassifiedImages from the specified map of raw messages.
func UnmarshalClassifiedImages(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ClassifiedImages)
	err = core.UnmarshalPrimitive(m, "custom_classes", &obj.CustomClasses)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "images_processed", &obj.ImagesProcessed)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "images", &obj.Images, UnmarshalClassifiedImage)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "warnings", &obj.Warnings, UnmarshalWarningInfo)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Classifier : Information about a classifier.
type Classifier struct {
	// ID of a classifier identified in the image.
	ClassifierID *string `json:"classifier_id" validate:"required"`

	// Name of the classifier.
	Name *string `json:"name" validate:"required"`

	// Unique ID of the account who owns the classifier. Might not be returned by some requests.
	Owner *string `json:"owner,omitempty"`

	// Training status of classifier.
	Status *string `json:"status,omitempty"`

	// Whether the classifier can be downloaded as a Core ML model after the training status is `ready`.
	CoreMlEnabled *bool `json:"core_ml_enabled,omitempty"`

	// If classifier training has failed, this field might explain why.
	Explanation *string `json:"explanation,omitempty"`

	// Date and time in Coordinated Universal Time (UTC) that the classifier was created.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// Classes that define a classifier.
	Classes []Class `json:"classes,omitempty"`

	// Date and time in Coordinated Universal Time (UTC) that the classifier was updated. Might not be returned by some
	// requests. Identical to `updated` and retained for backward compatibility.
	Retrained *strfmt.DateTime `json:"retrained,omitempty"`

	// Date and time in Coordinated Universal Time (UTC) that the classifier was most recently updated. The field matches
	// either `retrained` or `created`. Might not be returned by some requests.
	Updated *strfmt.DateTime `json:"updated,omitempty"`
}

// Constants associated with the Classifier.Status property.
// Training status of classifier.
const (
	ClassifierStatusFailedConst     = "failed"
	ClassifierStatusReadyConst      = "ready"
	ClassifierStatusRetrainingConst = "retraining"
	ClassifierStatusTrainingConst   = "training"
)

// UnmarshalClassifier unmarshals an instance of Classifier from the specified map of raw messages.
func UnmarshalClassifier(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Classifier)
	err = core.UnmarshalPrimitive(m, "classifier_id", &obj.ClassifierID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "owner", &obj.Owner)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "core_ml_enabled", &obj.CoreMlEnabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "explanation", &obj.Explanation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created", &obj.Created)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "classes", &obj.Classes, UnmarshalClass)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "retrained", &obj.Retrained)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated", &obj.Updated)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ClassifierResult : Classifier and score combination.
type ClassifierResult struct {
	// Name of the classifier.
	Name *string `json:"name" validate:"required"`

	// ID of a classifier identified in the image.
	ClassifierID *string `json:"classifier_id" validate:"required"`

	// Classes within the classifier.
	Classes []ClassResult `json:"classes" validate:"required"`
}

// UnmarshalClassifierResult unmarshals an instance of ClassifierResult from the specified map of raw messages.
func UnmarshalClassifierResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ClassifierResult)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "classifier_id", &obj.ClassifierID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "classes", &obj.Classes, UnmarshalClassResult)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Classifiers : A container for the list of classifiers.
type Classifiers struct {
	// List of classifiers.
	Classifiers []Classifier `json:"classifiers" validate:"required"`
}

// UnmarshalClassifiers unmarshals an instance of Classifiers from the specified map of raw messages.
func UnmarshalClassifiers(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Classifiers)
	err = core.UnmarshalModel(m, "classifiers", &obj.Classifiers, UnmarshalClassifier)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ClassifyOptions : The Classify options.
type ClassifyOptions struct {
	// An image file (.gif, .jpg, .png, .tif) or .zip file with images. Maximum image size is 10 MB. Include no more than
	// 20 images and limit the .zip file to 100 MB. Encode the image and .zip file names in UTF-8 if they contain non-ASCII
	// characters. The service assumes UTF-8 encoding if it encounters non-ASCII characters.
	//
	// You can also include an image with the **url** parameter.
	ImagesFile io.ReadCloser `json:"images_file,omitempty"`

	// The filename for imagesFile.
	ImagesFilename *string `json:"images_filename,omitempty"`

	// The content type of imagesFile.
	ImagesFileContentType *string `json:"images_file_content_type,omitempty"`

	// The URL of an image (.gif, .jpg, .png, .tif) to analyze. The minimum recommended pixel density is 32X32 pixels, but
	// the service tends to perform better with images that are at least 224 x 224 pixels. The maximum image size is 10 MB.
	//
	// You can also include images with the **images_file** parameter.
	URL *string `json:"url,omitempty"`

	// The minimum score a class must have to be displayed in the response. Set the threshold to `0.0` to return all
	// identified classes.
	Threshold *float32 `json:"threshold,omitempty"`

	// The categories of classifiers to apply. The **classifier_ids** parameter overrides **owners**, so make sure that
	// **classifier_ids** is empty.
	// - Use `IBM` to classify against the `default` general classifier. You get the same result if both **classifier_ids**
	// and **owners** parameters are empty.
	// - Use `me` to classify against all your custom classifiers. However, for better performance use **classifier_ids**
	// to specify the specific custom classifiers to apply.
	// - Use both `IBM` and `me` to analyze the image against both classifier categories.
	Owners []string `json:"owners,omitempty"`

	// Which classifiers to apply. Overrides the **owners** parameter. You can specify both custom and built-in classifier
	// IDs. The built-in `default` classifier is used if both **classifier_ids** and **owners** parameters are empty.
	//
	// The following built-in classifier IDs require no training:
	// - `default`: Returns classes from thousands of general tags.
	// - `food`: Enhances specificity and accuracy for images of food items.
	// - `explicit`: Evaluates whether the image might be pornographic.
	ClassifierIds []string `json:"classifier_ids,omitempty"`

	// The desired language of parts of the response. See the response for details.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ClassifyOptions.AcceptLanguage property.
// The desired language of parts of the response. See the response for details.
const (
	ClassifyOptionsAcceptLanguageArConst   = "ar"
	ClassifyOptionsAcceptLanguageDeConst   = "de"
	ClassifyOptionsAcceptLanguageEnConst   = "en"
	ClassifyOptionsAcceptLanguageEsConst   = "es"
	ClassifyOptionsAcceptLanguageFrConst   = "fr"
	ClassifyOptionsAcceptLanguageItConst   = "it"
	ClassifyOptionsAcceptLanguageJaConst   = "ja"
	ClassifyOptionsAcceptLanguageKoConst   = "ko"
	ClassifyOptionsAcceptLanguagePtBrConst = "pt-br"
	ClassifyOptionsAcceptLanguageZhCnConst = "zh-cn"
	ClassifyOptionsAcceptLanguageZhTwConst = "zh-tw"
)

// NewClassifyOptions : Instantiate ClassifyOptions
func (*VisualRecognitionV3) NewClassifyOptions() *ClassifyOptions {
	return &ClassifyOptions{}
}

// SetImagesFile : Allow user to set ImagesFile
func (options *ClassifyOptions) SetImagesFile(imagesFile io.ReadCloser) *ClassifyOptions {
	options.ImagesFile = imagesFile
	return options
}

// SetImagesFilename : Allow user to set ImagesFilename
func (options *ClassifyOptions) SetImagesFilename(imagesFilename string) *ClassifyOptions {
	options.ImagesFilename = core.StringPtr(imagesFilename)
	return options
}

// SetImagesFileContentType : Allow user to set ImagesFileContentType
func (options *ClassifyOptions) SetImagesFileContentType(imagesFileContentType string) *ClassifyOptions {
	options.ImagesFileContentType = core.StringPtr(imagesFileContentType)
	return options
}

// SetURL : Allow user to set URL
func (options *ClassifyOptions) SetURL(url string) *ClassifyOptions {
	options.URL = core.StringPtr(url)
	return options
}

// SetThreshold : Allow user to set Threshold
func (options *ClassifyOptions) SetThreshold(threshold float32) *ClassifyOptions {
	options.Threshold = core.Float32Ptr(threshold)
	return options
}

// SetOwners : Allow user to set Owners
func (options *ClassifyOptions) SetOwners(owners []string) *ClassifyOptions {
	options.Owners = owners
	return options
}

// SetClassifierIds : Allow user to set ClassifierIds
func (options *ClassifyOptions) SetClassifierIds(classifierIds []string) *ClassifyOptions {
	options.ClassifierIds = classifierIds
	return options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (options *ClassifyOptions) SetAcceptLanguage(acceptLanguage string) *ClassifyOptions {
	options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ClassifyOptions) SetHeaders(param map[string]string) *ClassifyOptions {
	options.Headers = param
	return options
}

// CreateClassifierOptions : The CreateClassifier options.
type CreateClassifierOptions struct {
	// The name of the new classifier. Encode special characters in UTF-8.
	Name *string `json:"name" validate:"required"`

	// A .zip file of images that depict the visual subject of a class in the new classifier. You can include more than one
	// positive example file in a call.
	//
	// Specify the parameter name by appending `_positive_examples` to the class name. For example,
	// `goldenretriever_positive_examples` creates the class **goldenretriever**. The string cannot contain the following
	// characters: ``$ * - { } \ | / ' " ` [ ]``.
	//
	// Include at least 10 images in .jpg or .png format. The minimum recommended image resolution is 32X32 pixels. The
	// maximum number of images is 10,000 images or 100 MB per .zip file.
	//
	// Encode special characters in the file name in UTF-8.
	PositiveExamples map[string]io.ReadCloser `json:"positive_examples" validate:"required"`

	// A .zip file of images that do not depict the visual subject of any of the classes of the new classifier. Must
	// contain a minimum of 10 images.
	//
	// Encode special characters in the file name in UTF-8.
	NegativeExamples io.ReadCloser `json:"negative_examples,omitempty"`

	// The filename for negativeExamples.
	NegativeExamplesFilename *string `json:"negative_examples_filename,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateClassifierOptions : Instantiate CreateClassifierOptions
func (*VisualRecognitionV3) NewCreateClassifierOptions(name string) *CreateClassifierOptions {
	return &CreateClassifierOptions{
		Name: core.StringPtr(name),
	}
}

// SetName : Allow user to set Name
func (options *CreateClassifierOptions) SetName(name string) *CreateClassifierOptions {
	options.Name = core.StringPtr(name)
	return options
}

// AddPositiveExamples : Allow user to add a new entry to the PositiveExamples map
func (options *CreateClassifierOptions) AddPositiveExamples(classname string, positiveExamples io.ReadCloser) *CreateClassifierOptions {
	if options.PositiveExamples == nil {
		options.PositiveExamples = make(map[string]io.ReadCloser)
	}
	options.PositiveExamples[classname] = positiveExamples
	return options
}

// SetNegativeExamples : Allow user to set NegativeExamples
func (options *CreateClassifierOptions) SetNegativeExamples(negativeExamples io.ReadCloser) *CreateClassifierOptions {
	options.NegativeExamples = negativeExamples
	return options
}

// SetNegativeExamplesFilename : Allow user to set NegativeExamplesFilename
func (options *CreateClassifierOptions) SetNegativeExamplesFilename(negativeExamplesFilename string) *CreateClassifierOptions {
	options.NegativeExamplesFilename = core.StringPtr(negativeExamplesFilename)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateClassifierOptions) SetHeaders(param map[string]string) *CreateClassifierOptions {
	options.Headers = param
	return options
}

// DeleteClassifierOptions : The DeleteClassifier options.
type DeleteClassifierOptions struct {
	// The ID of the classifier.
	ClassifierID *string `json:"classifier_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteClassifierOptions : Instantiate DeleteClassifierOptions
func (*VisualRecognitionV3) NewDeleteClassifierOptions(classifierID string) *DeleteClassifierOptions {
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

// DeleteUserDataOptions : The DeleteUserData options.
type DeleteUserDataOptions struct {
	// The customer ID for which all data is to be deleted.
	CustomerID *string `json:"customer_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteUserDataOptions : Instantiate DeleteUserDataOptions
func (*VisualRecognitionV3) NewDeleteUserDataOptions(customerID string) *DeleteUserDataOptions {
	return &DeleteUserDataOptions{
		CustomerID: core.StringPtr(customerID),
	}
}

// SetCustomerID : Allow user to set CustomerID
func (options *DeleteUserDataOptions) SetCustomerID(customerID string) *DeleteUserDataOptions {
	options.CustomerID = core.StringPtr(customerID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteUserDataOptions) SetHeaders(param map[string]string) *DeleteUserDataOptions {
	options.Headers = param
	return options
}

// ErrorInfo : Information about what might have caused a failure, such as an image that is too large. Not returned when there is no
// error.
type ErrorInfo struct {
	// HTTP status code.
	Code *int64 `json:"code" validate:"required"`

	// Human-readable error description. For example, `File size limit exceeded`.
	Description *string `json:"description" validate:"required"`

	// Codified error string. For example, `limit_exceeded`.
	ErrorID *string `json:"error_id" validate:"required"`
}

// UnmarshalErrorInfo unmarshals an instance of ErrorInfo from the specified map of raw messages.
func UnmarshalErrorInfo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ErrorInfo)
	err = core.UnmarshalPrimitive(m, "code", &obj.Code)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "error_id", &obj.ErrorID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetClassifierOptions : The GetClassifier options.
type GetClassifierOptions struct {
	// The ID of the classifier.
	ClassifierID *string `json:"classifier_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetClassifierOptions : Instantiate GetClassifierOptions
func (*VisualRecognitionV3) NewGetClassifierOptions(classifierID string) *GetClassifierOptions {
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

// GetCoreMlModelOptions : The GetCoreMlModel options.
type GetCoreMlModelOptions struct {
	// The ID of the classifier.
	ClassifierID *string `json:"classifier_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCoreMlModelOptions : Instantiate GetCoreMlModelOptions
func (*VisualRecognitionV3) NewGetCoreMlModelOptions(classifierID string) *GetCoreMlModelOptions {
	return &GetCoreMlModelOptions{
		ClassifierID: core.StringPtr(classifierID),
	}
}

// SetClassifierID : Allow user to set ClassifierID
func (options *GetCoreMlModelOptions) SetClassifierID(classifierID string) *GetCoreMlModelOptions {
	options.ClassifierID = core.StringPtr(classifierID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetCoreMlModelOptions) SetHeaders(param map[string]string) *GetCoreMlModelOptions {
	options.Headers = param
	return options
}

// ListClassifiersOptions : The ListClassifiers options.
type ListClassifiersOptions struct {
	// Specify `true` to return details about the classifiers. Omit this parameter to return a brief list of classifiers.
	Verbose *bool `json:"verbose,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListClassifiersOptions : Instantiate ListClassifiersOptions
func (*VisualRecognitionV3) NewListClassifiersOptions() *ListClassifiersOptions {
	return &ListClassifiersOptions{}
}

// SetVerbose : Allow user to set Verbose
func (options *ListClassifiersOptions) SetVerbose(verbose bool) *ListClassifiersOptions {
	options.Verbose = core.BoolPtr(verbose)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListClassifiersOptions) SetHeaders(param map[string]string) *ListClassifiersOptions {
	options.Headers = param
	return options
}

// UpdateClassifierOptions : The UpdateClassifier options.
type UpdateClassifierOptions struct {
	// The ID of the classifier.
	ClassifierID *string `json:"classifier_id" validate:"required,ne="`

	// A .zip file of images that depict the visual subject of a class in the classifier. The positive examples create or
	// update classes in the classifier. You can include more than one positive example file in a call.
	//
	// Specify the parameter name by appending `_positive_examples` to the class name. For example,
	// `goldenretriever_positive_examples` creates the class `goldenretriever`. The string cannot contain the following
	// characters: ``$ * - { } \ | / ' " ` [ ]``.
	//
	// Include at least 10 images in .jpg or .png format. The minimum recommended image resolution is 32X32 pixels. The
	// maximum number of images is 10,000 images or 100 MB per .zip file.
	//
	// Encode special characters in the file name in UTF-8.
	PositiveExamples map[string]io.ReadCloser `json:"positive_examples,omitempty"`

	// A .zip file of images that do not depict the visual subject of any of the classes of the new classifier. Must
	// contain a minimum of 10 images.
	//
	// Encode special characters in the file name in UTF-8.
	NegativeExamples io.ReadCloser `json:"negative_examples,omitempty"`

	// The filename for negativeExamples.
	NegativeExamplesFilename *string `json:"negative_examples_filename,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateClassifierOptions : Instantiate UpdateClassifierOptions
func (*VisualRecognitionV3) NewUpdateClassifierOptions(classifierID string) *UpdateClassifierOptions {
	return &UpdateClassifierOptions{
		ClassifierID: core.StringPtr(classifierID),
	}
}

// SetClassifierID : Allow user to set ClassifierID
func (options *UpdateClassifierOptions) SetClassifierID(classifierID string) *UpdateClassifierOptions {
	options.ClassifierID = core.StringPtr(classifierID)
	return options
}

// AddPositiveExamples : Allow user to add a new entry to the PositiveExamples map
func (options *UpdateClassifierOptions) AddPositiveExamples(classname string, positiveExamples io.ReadCloser) *UpdateClassifierOptions {
	if options.PositiveExamples == nil {
		options.PositiveExamples = make(map[string]io.ReadCloser)
	}
	options.PositiveExamples[classname] = positiveExamples
	return options
}

// SetNegativeExamples : Allow user to set NegativeExamples
func (options *UpdateClassifierOptions) SetNegativeExamples(negativeExamples io.ReadCloser) *UpdateClassifierOptions {
	options.NegativeExamples = negativeExamples
	return options
}

// SetNegativeExamplesFilename : Allow user to set NegativeExamplesFilename
func (options *UpdateClassifierOptions) SetNegativeExamplesFilename(negativeExamplesFilename string) *UpdateClassifierOptions {
	options.NegativeExamplesFilename = core.StringPtr(negativeExamplesFilename)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateClassifierOptions) SetHeaders(param map[string]string) *UpdateClassifierOptions {
	options.Headers = param
	return options
}

// WarningInfo : Information about something that went wrong.
type WarningInfo struct {
	// Codified warning string, such as `limit_reached`.
	WarningID *string `json:"warning_id" validate:"required"`

	// Information about the error.
	Description *string `json:"description" validate:"required"`
}

// UnmarshalWarningInfo unmarshals an instance of WarningInfo from the specified map of raw messages.
func UnmarshalWarningInfo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(WarningInfo)
	err = core.UnmarshalPrimitive(m, "warning_id", &obj.WarningID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
