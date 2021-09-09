/**
 * (C) Copyright IBM Corp. 2019, 2021.
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
 * IBM OpenAPI SDK Code Generator Version: 3.38.0-07189efd-20210827-205025
 */

// Package visualrecognitionv4 : Operations and models for the VisualRecognitionV4 service
package visualrecognitionv4

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	common "github.com/watson-developer-cloud/go-sdk/v2/common"
)

// VisualRecognitionV4 : IBM Watson&trade; Visual Recognition is discontinued. Existing instances are supported until 1
// December 2021, but as of 7 January 2021, you can't create instances. Any instance that is provisioned on 1 December
// 2021 will be deleted.
// {: deprecated}
//
// Provide images to the IBM Watson Visual Recognition service for analysis. The service detects objects based on a set
// of images with training data.
//
// API Version: 4.0
// See: https://cloud.ibm.com/docs/visual-recognition?topic=visual-recognition-object-detection-overview
type VisualRecognitionV4 struct {
	Service *core.BaseService

	// Release date of the API version you want to use. Specify dates in YYYY-MM-DD format. The current version is
	// `2019-02-11`.
	Version *string
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.us-south.visual-recognition.watson.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "watson_vision_combined"

// VisualRecognitionV4Options : Service options
type VisualRecognitionV4Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator

	// Release date of the API version you want to use. Specify dates in YYYY-MM-DD format. The current version is
	// `2019-02-11`.
	Version *string `validate:"required"`
}

// NewVisualRecognitionV4 : constructs an instance of VisualRecognitionV4 with passed in options.
func NewVisualRecognitionV4(options *VisualRecognitionV4Options) (service *VisualRecognitionV4, err error) {
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

	service = &VisualRecognitionV4{
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
func (visualRecognition *VisualRecognitionV4) Clone() *VisualRecognitionV4 {
	if core.IsNil(visualRecognition) {
		return nil
	}
	clone := *visualRecognition
	clone.Service = visualRecognition.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (visualRecognition *VisualRecognitionV4) SetServiceURL(url string) error {
	return visualRecognition.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (visualRecognition *VisualRecognitionV4) GetServiceURL() string {
	return visualRecognition.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (visualRecognition *VisualRecognitionV4) SetDefaultHeaders(headers http.Header) {
	visualRecognition.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (visualRecognition *VisualRecognitionV4) SetEnableGzipCompression(enableGzip bool) {
	visualRecognition.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (visualRecognition *VisualRecognitionV4) GetEnableGzipCompression() bool {
	return visualRecognition.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (visualRecognition *VisualRecognitionV4) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	visualRecognition.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (visualRecognition *VisualRecognitionV4) DisableRetries() {
	visualRecognition.Service.DisableRetries()
}

// DisableSSLVerification bypasses verification of the server's SSL certificate
func (visualRecognition *VisualRecognitionV4) DisableSSLVerification() {
	visualRecognition.Service.DisableSSLVerification()
}

// Analyze : Analyze images
// Analyze images by URL, by file, or both against your own collection. Make sure that **training_status.objects.ready**
// is `true` for the feature before you use a collection to analyze images.
//
// Encode the image and .zip file names in UTF-8 if they contain non-ASCII characters. The service assumes UTF-8
// encoding if it encounters non-ASCII characters.
func (visualRecognition *VisualRecognitionV4) Analyze(analyzeOptions *AnalyzeOptions) (result *AnalyzeResponse, response *core.DetailedResponse, err error) {
	return visualRecognition.AnalyzeWithContext(context.Background(), analyzeOptions)
}

// AnalyzeWithContext is an alternate form of the Analyze method which supports a Context parameter
func (visualRecognition *VisualRecognitionV4) AnalyzeWithContext(ctx context.Context, analyzeOptions *AnalyzeOptions) (result *AnalyzeResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(analyzeOptions, "analyzeOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(analyzeOptions, "analyzeOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = visualRecognition.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(visualRecognition.Service.Options.URL, `/v4/analyze`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range analyzeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_vision_combined", "V4", "Analyze")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*visualRecognition.Version))

	for _, item := range analyzeOptions.CollectionIds {
		builder.AddFormData("collection_ids", "", "", fmt.Sprint(item))
	}
	for _, item := range analyzeOptions.Features {
		builder.AddFormData("features", "", "", fmt.Sprint(item))
	}
	if analyzeOptions.ImagesFile != nil {
		for _, item := range analyzeOptions.ImagesFile {
			builder.AddFormData("images_file", core.StringNilMapper(item.Filename), core.StringNilMapper(item.ContentType), item.Data)
		}
	}
	if analyzeOptions.ImageURL != nil {
		for _, item := range analyzeOptions.ImageURL {
			builder.AddFormData("image_url", "", "", fmt.Sprint(item))
		}
	}
	if analyzeOptions.Threshold != nil {
		builder.AddFormData("threshold", "", "", fmt.Sprint(*analyzeOptions.Threshold))
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAnalyzeResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateCollection : Create a collection
// Create a collection that can be used to store images.
//
// To create a collection without specifying a name and description, include an empty JSON object in the request body.
//
// Encode the name and description in UTF-8 if they contain non-ASCII characters. The service assumes UTF-8 encoding if
// it encounters non-ASCII characters.
func (visualRecognition *VisualRecognitionV4) CreateCollection(createCollectionOptions *CreateCollectionOptions) (result *Collection, response *core.DetailedResponse, err error) {
	return visualRecognition.CreateCollectionWithContext(context.Background(), createCollectionOptions)
}

// CreateCollectionWithContext is an alternate form of the CreateCollection method which supports a Context parameter
func (visualRecognition *VisualRecognitionV4) CreateCollectionWithContext(ctx context.Context, createCollectionOptions *CreateCollectionOptions) (result *Collection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createCollectionOptions, "createCollectionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createCollectionOptions, "createCollectionOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = visualRecognition.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(visualRecognition.Service.Options.URL, `/v4/collections`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createCollectionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_vision_combined", "V4", "CreateCollection")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*visualRecognition.Version))

	body := make(map[string]interface{})
	if createCollectionOptions.Name != nil {
		body["name"] = createCollectionOptions.Name
	}
	if createCollectionOptions.Description != nil {
		body["description"] = createCollectionOptions.Description
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
	response, err = visualRecognition.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListCollections : List collections
// Retrieves a list of collections for the service instance.
func (visualRecognition *VisualRecognitionV4) ListCollections(listCollectionsOptions *ListCollectionsOptions) (result *CollectionsList, response *core.DetailedResponse, err error) {
	return visualRecognition.ListCollectionsWithContext(context.Background(), listCollectionsOptions)
}

// ListCollectionsWithContext is an alternate form of the ListCollections method which supports a Context parameter
func (visualRecognition *VisualRecognitionV4) ListCollectionsWithContext(ctx context.Context, listCollectionsOptions *ListCollectionsOptions) (result *CollectionsList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listCollectionsOptions, "listCollectionsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = visualRecognition.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(visualRecognition.Service.Options.URL, `/v4/collections`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listCollectionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_vision_combined", "V4", "ListCollections")
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCollectionsList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetCollection : Get collection details
// Get details of one collection.
func (visualRecognition *VisualRecognitionV4) GetCollection(getCollectionOptions *GetCollectionOptions) (result *Collection, response *core.DetailedResponse, err error) {
	return visualRecognition.GetCollectionWithContext(context.Background(), getCollectionOptions)
}

// GetCollectionWithContext is an alternate form of the GetCollection method which supports a Context parameter
func (visualRecognition *VisualRecognitionV4) GetCollectionWithContext(ctx context.Context, getCollectionOptions *GetCollectionOptions) (result *Collection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getCollectionOptions, "getCollectionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getCollectionOptions, "getCollectionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"collection_id": *getCollectionOptions.CollectionID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = visualRecognition.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(visualRecognition.Service.Options.URL, `/v4/collections/{collection_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCollectionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_vision_combined", "V4", "GetCollection")
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateCollection : Update a collection
// Update the name or description of a collection.
//
// Encode the name and description in UTF-8 if they contain non-ASCII characters. The service assumes UTF-8 encoding if
// it encounters non-ASCII characters.
func (visualRecognition *VisualRecognitionV4) UpdateCollection(updateCollectionOptions *UpdateCollectionOptions) (result *Collection, response *core.DetailedResponse, err error) {
	return visualRecognition.UpdateCollectionWithContext(context.Background(), updateCollectionOptions)
}

// UpdateCollectionWithContext is an alternate form of the UpdateCollection method which supports a Context parameter
func (visualRecognition *VisualRecognitionV4) UpdateCollectionWithContext(ctx context.Context, updateCollectionOptions *UpdateCollectionOptions) (result *Collection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateCollectionOptions, "updateCollectionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateCollectionOptions, "updateCollectionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"collection_id": *updateCollectionOptions.CollectionID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = visualRecognition.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(visualRecognition.Service.Options.URL, `/v4/collections/{collection_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateCollectionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_vision_combined", "V4", "UpdateCollection")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*visualRecognition.Version))

	body := make(map[string]interface{})
	if updateCollectionOptions.Name != nil {
		body["name"] = updateCollectionOptions.Name
	}
	if updateCollectionOptions.Description != nil {
		body["description"] = updateCollectionOptions.Description
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
	response, err = visualRecognition.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteCollection : Delete a collection
// Delete a collection from the service instance.
func (visualRecognition *VisualRecognitionV4) DeleteCollection(deleteCollectionOptions *DeleteCollectionOptions) (response *core.DetailedResponse, err error) {
	return visualRecognition.DeleteCollectionWithContext(context.Background(), deleteCollectionOptions)
}

// DeleteCollectionWithContext is an alternate form of the DeleteCollection method which supports a Context parameter
func (visualRecognition *VisualRecognitionV4) DeleteCollectionWithContext(ctx context.Context, deleteCollectionOptions *DeleteCollectionOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteCollectionOptions, "deleteCollectionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteCollectionOptions, "deleteCollectionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"collection_id": *deleteCollectionOptions.CollectionID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = visualRecognition.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(visualRecognition.Service.Options.URL, `/v4/collections/{collection_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteCollectionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_vision_combined", "V4", "DeleteCollection")
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

// GetModelFile : Get a model
// Download a model that you can deploy to detect objects in images. The collection must include a generated model,
// which is indicated in the response for the collection details as `"rscnn_ready": true`. If the value is `false`,
// train or retrain the collection to generate the model.
//
// Currently, the model format is specific to Android apps. For more information about how to deploy the model to your
// app, see the [Watson Visual Recognition on Android](https://github.com/matt-ny/rscnn) project in GitHub.
func (visualRecognition *VisualRecognitionV4) GetModelFile(getModelFileOptions *GetModelFileOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	return visualRecognition.GetModelFileWithContext(context.Background(), getModelFileOptions)
}

// GetModelFileWithContext is an alternate form of the GetModelFile method which supports a Context parameter
func (visualRecognition *VisualRecognitionV4) GetModelFileWithContext(ctx context.Context, getModelFileOptions *GetModelFileOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getModelFileOptions, "getModelFileOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getModelFileOptions, "getModelFileOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"collection_id": *getModelFileOptions.CollectionID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = visualRecognition.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(visualRecognition.Service.Options.URL, `/v4/collections/{collection_id}/model`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getModelFileOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_vision_combined", "V4", "GetModelFile")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/octet-stream")

	builder.AddQuery("version", fmt.Sprint(*visualRecognition.Version))
	builder.AddQuery("feature", fmt.Sprint(*getModelFileOptions.Feature))
	builder.AddQuery("model_format", fmt.Sprint(*getModelFileOptions.ModelFormat))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = visualRecognition.Service.Request(request, &result)

	return
}

// AddImages : Add images
// Add images to a collection by URL, by file, or both.
//
// Encode the image and .zip file names in UTF-8 if they contain non-ASCII characters. The service assumes UTF-8
// encoding if it encounters non-ASCII characters.
func (visualRecognition *VisualRecognitionV4) AddImages(addImagesOptions *AddImagesOptions) (result *ImageDetailsList, response *core.DetailedResponse, err error) {
	return visualRecognition.AddImagesWithContext(context.Background(), addImagesOptions)
}

// AddImagesWithContext is an alternate form of the AddImages method which supports a Context parameter
func (visualRecognition *VisualRecognitionV4) AddImagesWithContext(ctx context.Context, addImagesOptions *AddImagesOptions) (result *ImageDetailsList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addImagesOptions, "addImagesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addImagesOptions, "addImagesOptions")
	if err != nil {
		return
	}
	if (addImagesOptions.ImagesFile == nil) && (addImagesOptions.ImageURL == nil) && (addImagesOptions.TrainingData == nil) {
		err = fmt.Errorf("at least one of imagesFile, imageURL, or trainingData must be supplied")
		return
	}

	pathParamsMap := map[string]string{
		"collection_id": *addImagesOptions.CollectionID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = visualRecognition.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(visualRecognition.Service.Options.URL, `/v4/collections/{collection_id}/images`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range addImagesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_vision_combined", "V4", "AddImages")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*visualRecognition.Version))

	if addImagesOptions.ImagesFile != nil {
		for _, item := range addImagesOptions.ImagesFile {
			builder.AddFormData("images_file", core.StringNilMapper(item.Filename), core.StringNilMapper(item.ContentType), item.Data)
		}
	}
	if addImagesOptions.ImageURL != nil {
		for _, item := range addImagesOptions.ImageURL {
			builder.AddFormData("image_url", "", "", fmt.Sprint(item))
		}
	}
	if addImagesOptions.TrainingData != nil {
		builder.AddFormData("training_data", "", "", fmt.Sprint(*addImagesOptions.TrainingData))
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalImageDetailsList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListImages : List images
// Retrieves a list of images in a collection.
func (visualRecognition *VisualRecognitionV4) ListImages(listImagesOptions *ListImagesOptions) (result *ImageSummaryList, response *core.DetailedResponse, err error) {
	return visualRecognition.ListImagesWithContext(context.Background(), listImagesOptions)
}

// ListImagesWithContext is an alternate form of the ListImages method which supports a Context parameter
func (visualRecognition *VisualRecognitionV4) ListImagesWithContext(ctx context.Context, listImagesOptions *ListImagesOptions) (result *ImageSummaryList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listImagesOptions, "listImagesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listImagesOptions, "listImagesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"collection_id": *listImagesOptions.CollectionID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = visualRecognition.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(visualRecognition.Service.Options.URL, `/v4/collections/{collection_id}/images`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listImagesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_vision_combined", "V4", "ListImages")
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalImageSummaryList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetImageDetails : Get image details
// Get the details of an image in a collection.
func (visualRecognition *VisualRecognitionV4) GetImageDetails(getImageDetailsOptions *GetImageDetailsOptions) (result *ImageDetails, response *core.DetailedResponse, err error) {
	return visualRecognition.GetImageDetailsWithContext(context.Background(), getImageDetailsOptions)
}

// GetImageDetailsWithContext is an alternate form of the GetImageDetails method which supports a Context parameter
func (visualRecognition *VisualRecognitionV4) GetImageDetailsWithContext(ctx context.Context, getImageDetailsOptions *GetImageDetailsOptions) (result *ImageDetails, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getImageDetailsOptions, "getImageDetailsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getImageDetailsOptions, "getImageDetailsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"collection_id": *getImageDetailsOptions.CollectionID,
		"image_id":      *getImageDetailsOptions.ImageID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = visualRecognition.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(visualRecognition.Service.Options.URL, `/v4/collections/{collection_id}/images/{image_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getImageDetailsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_vision_combined", "V4", "GetImageDetails")
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalImageDetails)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteImage : Delete an image
// Delete one image from a collection.
func (visualRecognition *VisualRecognitionV4) DeleteImage(deleteImageOptions *DeleteImageOptions) (response *core.DetailedResponse, err error) {
	return visualRecognition.DeleteImageWithContext(context.Background(), deleteImageOptions)
}

// DeleteImageWithContext is an alternate form of the DeleteImage method which supports a Context parameter
func (visualRecognition *VisualRecognitionV4) DeleteImageWithContext(ctx context.Context, deleteImageOptions *DeleteImageOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteImageOptions, "deleteImageOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteImageOptions, "deleteImageOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"collection_id": *deleteImageOptions.CollectionID,
		"image_id":      *deleteImageOptions.ImageID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = visualRecognition.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(visualRecognition.Service.Options.URL, `/v4/collections/{collection_id}/images/{image_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteImageOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_vision_combined", "V4", "DeleteImage")
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

// GetJpegImage : Get a JPEG file of an image
// Download a JPEG representation of an image.
func (visualRecognition *VisualRecognitionV4) GetJpegImage(getJpegImageOptions *GetJpegImageOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	return visualRecognition.GetJpegImageWithContext(context.Background(), getJpegImageOptions)
}

// GetJpegImageWithContext is an alternate form of the GetJpegImage method which supports a Context parameter
func (visualRecognition *VisualRecognitionV4) GetJpegImageWithContext(ctx context.Context, getJpegImageOptions *GetJpegImageOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getJpegImageOptions, "getJpegImageOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getJpegImageOptions, "getJpegImageOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"collection_id": *getJpegImageOptions.CollectionID,
		"image_id":      *getJpegImageOptions.ImageID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = visualRecognition.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(visualRecognition.Service.Options.URL, `/v4/collections/{collection_id}/images/{image_id}/jpeg`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getJpegImageOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_vision_combined", "V4", "GetJpegImage")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "image/jpeg")

	builder.AddQuery("version", fmt.Sprint(*visualRecognition.Version))
	if getJpegImageOptions.Size != nil {
		builder.AddQuery("size", fmt.Sprint(*getJpegImageOptions.Size))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = visualRecognition.Service.Request(request, &result)

	return
}

// ListObjectMetadata : List object metadata
// Retrieves a list of object names in a collection.
func (visualRecognition *VisualRecognitionV4) ListObjectMetadata(listObjectMetadataOptions *ListObjectMetadataOptions) (result *ObjectMetadataList, response *core.DetailedResponse, err error) {
	return visualRecognition.ListObjectMetadataWithContext(context.Background(), listObjectMetadataOptions)
}

// ListObjectMetadataWithContext is an alternate form of the ListObjectMetadata method which supports a Context parameter
func (visualRecognition *VisualRecognitionV4) ListObjectMetadataWithContext(ctx context.Context, listObjectMetadataOptions *ListObjectMetadataOptions) (result *ObjectMetadataList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listObjectMetadataOptions, "listObjectMetadataOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listObjectMetadataOptions, "listObjectMetadataOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"collection_id": *listObjectMetadataOptions.CollectionID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = visualRecognition.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(visualRecognition.Service.Options.URL, `/v4/collections/{collection_id}/objects`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listObjectMetadataOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_vision_combined", "V4", "ListObjectMetadata")
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalObjectMetadataList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateObjectMetadata : Update an object name
// Update the name of an object. A successful request updates the training data for all images that use the object.
func (visualRecognition *VisualRecognitionV4) UpdateObjectMetadata(updateObjectMetadataOptions *UpdateObjectMetadataOptions) (result *UpdateObjectMetadata, response *core.DetailedResponse, err error) {
	return visualRecognition.UpdateObjectMetadataWithContext(context.Background(), updateObjectMetadataOptions)
}

// UpdateObjectMetadataWithContext is an alternate form of the UpdateObjectMetadata method which supports a Context parameter
func (visualRecognition *VisualRecognitionV4) UpdateObjectMetadataWithContext(ctx context.Context, updateObjectMetadataOptions *UpdateObjectMetadataOptions) (result *UpdateObjectMetadata, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateObjectMetadataOptions, "updateObjectMetadataOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateObjectMetadataOptions, "updateObjectMetadataOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"collection_id": *updateObjectMetadataOptions.CollectionID,
		"object":        *updateObjectMetadataOptions.Object,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = visualRecognition.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(visualRecognition.Service.Options.URL, `/v4/collections/{collection_id}/objects/{object}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateObjectMetadataOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_vision_combined", "V4", "UpdateObjectMetadata")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*visualRecognition.Version))

	body := make(map[string]interface{})
	if updateObjectMetadataOptions.NewObject != nil {
		body["object"] = updateObjectMetadataOptions.NewObject
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
	response, err = visualRecognition.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalUpdateObjectMetadata)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetObjectMetadata : Get object metadata
// Get the number of bounding boxes for a single object in a collection.
func (visualRecognition *VisualRecognitionV4) GetObjectMetadata(getObjectMetadataOptions *GetObjectMetadataOptions) (result *ObjectMetadata, response *core.DetailedResponse, err error) {
	return visualRecognition.GetObjectMetadataWithContext(context.Background(), getObjectMetadataOptions)
}

// GetObjectMetadataWithContext is an alternate form of the GetObjectMetadata method which supports a Context parameter
func (visualRecognition *VisualRecognitionV4) GetObjectMetadataWithContext(ctx context.Context, getObjectMetadataOptions *GetObjectMetadataOptions) (result *ObjectMetadata, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getObjectMetadataOptions, "getObjectMetadataOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getObjectMetadataOptions, "getObjectMetadataOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"collection_id": *getObjectMetadataOptions.CollectionID,
		"object":        *getObjectMetadataOptions.Object,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = visualRecognition.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(visualRecognition.Service.Options.URL, `/v4/collections/{collection_id}/objects/{object}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getObjectMetadataOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_vision_combined", "V4", "GetObjectMetadata")
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalObjectMetadata)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteObject : Delete an object
// Delete one object from a collection. A successful request deletes the training data from all images that use the
// object.
func (visualRecognition *VisualRecognitionV4) DeleteObject(deleteObjectOptions *DeleteObjectOptions) (response *core.DetailedResponse, err error) {
	return visualRecognition.DeleteObjectWithContext(context.Background(), deleteObjectOptions)
}

// DeleteObjectWithContext is an alternate form of the DeleteObject method which supports a Context parameter
func (visualRecognition *VisualRecognitionV4) DeleteObjectWithContext(ctx context.Context, deleteObjectOptions *DeleteObjectOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteObjectOptions, "deleteObjectOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteObjectOptions, "deleteObjectOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"collection_id": *deleteObjectOptions.CollectionID,
		"object":        *deleteObjectOptions.Object,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = visualRecognition.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(visualRecognition.Service.Options.URL, `/v4/collections/{collection_id}/objects/{object}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteObjectOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_vision_combined", "V4", "DeleteObject")
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

// Train : Train a collection
// Start training on images in a collection. The collection must have enough training data and untrained data (the
// **training_status.objects.data_changed** is `true`). If training is in progress, the request queues the next training
// job.
func (visualRecognition *VisualRecognitionV4) Train(trainOptions *TrainOptions) (result *Collection, response *core.DetailedResponse, err error) {
	return visualRecognition.TrainWithContext(context.Background(), trainOptions)
}

// TrainWithContext is an alternate form of the Train method which supports a Context parameter
func (visualRecognition *VisualRecognitionV4) TrainWithContext(ctx context.Context, trainOptions *TrainOptions) (result *Collection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(trainOptions, "trainOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(trainOptions, "trainOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"collection_id": *trainOptions.CollectionID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = visualRecognition.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(visualRecognition.Service.Options.URL, `/v4/collections/{collection_id}/train`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range trainOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_vision_combined", "V4", "Train")
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// AddImageTrainingData : Add training data to an image
// Add, update, or delete training data for an image. Encode the object name in UTF-8 if it contains non-ASCII
// characters. The service assumes UTF-8 encoding if it encounters non-ASCII characters.
//
// Elements in the request replace the existing elements.
//
// - To update the training data, provide both the unchanged and the new or changed values.
//
// - To delete the training data, provide an empty value for the training data.
func (visualRecognition *VisualRecognitionV4) AddImageTrainingData(addImageTrainingDataOptions *AddImageTrainingDataOptions) (result *TrainingDataObjects, response *core.DetailedResponse, err error) {
	return visualRecognition.AddImageTrainingDataWithContext(context.Background(), addImageTrainingDataOptions)
}

// AddImageTrainingDataWithContext is an alternate form of the AddImageTrainingData method which supports a Context parameter
func (visualRecognition *VisualRecognitionV4) AddImageTrainingDataWithContext(ctx context.Context, addImageTrainingDataOptions *AddImageTrainingDataOptions) (result *TrainingDataObjects, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addImageTrainingDataOptions, "addImageTrainingDataOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addImageTrainingDataOptions, "addImageTrainingDataOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"collection_id": *addImageTrainingDataOptions.CollectionID,
		"image_id":      *addImageTrainingDataOptions.ImageID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = visualRecognition.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(visualRecognition.Service.Options.URL, `/v4/collections/{collection_id}/images/{image_id}/training_data`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range addImageTrainingDataOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_vision_combined", "V4", "AddImageTrainingData")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*visualRecognition.Version))

	body := make(map[string]interface{})
	if addImageTrainingDataOptions.Objects != nil {
		body["objects"] = addImageTrainingDataOptions.Objects
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
	response, err = visualRecognition.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTrainingDataObjects)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetTrainingUsage : Get training usage
// Information about the completed training events. You can use this information to determine how close you are to the
// training limits for the month.
func (visualRecognition *VisualRecognitionV4) GetTrainingUsage(getTrainingUsageOptions *GetTrainingUsageOptions) (result *TrainingEvents, response *core.DetailedResponse, err error) {
	return visualRecognition.GetTrainingUsageWithContext(context.Background(), getTrainingUsageOptions)
}

// GetTrainingUsageWithContext is an alternate form of the GetTrainingUsage method which supports a Context parameter
func (visualRecognition *VisualRecognitionV4) GetTrainingUsageWithContext(ctx context.Context, getTrainingUsageOptions *GetTrainingUsageOptions) (result *TrainingEvents, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getTrainingUsageOptions, "getTrainingUsageOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = visualRecognition.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(visualRecognition.Service.Options.URL, `/v4/training_usage`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getTrainingUsageOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_vision_combined", "V4", "GetTrainingUsage")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*visualRecognition.Version))
	if getTrainingUsageOptions.StartTime != nil {
		builder.AddQuery("start_time", fmt.Sprint(*getTrainingUsageOptions.StartTime))
	}
	if getTrainingUsageOptions.EndTime != nil {
		builder.AddQuery("end_time", fmt.Sprint(*getTrainingUsageOptions.EndTime))
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTrainingEvents)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteUserData : Delete labeled data
// Deletes all data associated with a specified customer ID. The method has no effect if no data is associated with the
// customer ID.
//
// You associate a customer ID with data by passing the `X-Watson-Metadata` header with a request that passes data. For
// more information about personal data and customer IDs, see [Information
// security](https://cloud.ibm.com/docs/visual-recognition?topic=visual-recognition-information-security).
func (visualRecognition *VisualRecognitionV4) DeleteUserData(deleteUserDataOptions *DeleteUserDataOptions) (response *core.DetailedResponse, err error) {
	return visualRecognition.DeleteUserDataWithContext(context.Background(), deleteUserDataOptions)
}

// DeleteUserDataWithContext is an alternate form of the DeleteUserData method which supports a Context parameter
func (visualRecognition *VisualRecognitionV4) DeleteUserDataWithContext(ctx context.Context, deleteUserDataOptions *DeleteUserDataOptions) (response *core.DetailedResponse, err error) {
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
	_, err = builder.ResolveRequestURL(visualRecognition.Service.Options.URL, `/v4/user_data`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteUserDataOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_vision_combined", "V4", "DeleteUserData")
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

// AddImageTrainingDataOptions : The AddImageTrainingData options.
type AddImageTrainingDataOptions struct {
	// The identifier of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// The identifier of the image.
	ImageID *string `json:"-" validate:"required,ne="`

	// Training data for specific objects.
	Objects []TrainingDataObject `json:"objects,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAddImageTrainingDataOptions : Instantiate AddImageTrainingDataOptions
func (*VisualRecognitionV4) NewAddImageTrainingDataOptions(collectionID string, imageID string) *AddImageTrainingDataOptions {
	return &AddImageTrainingDataOptions{
		CollectionID: core.StringPtr(collectionID),
		ImageID:      core.StringPtr(imageID),
	}
}

// SetCollectionID : Allow user to set CollectionID
func (_options *AddImageTrainingDataOptions) SetCollectionID(collectionID string) *AddImageTrainingDataOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetImageID : Allow user to set ImageID
func (_options *AddImageTrainingDataOptions) SetImageID(imageID string) *AddImageTrainingDataOptions {
	_options.ImageID = core.StringPtr(imageID)
	return _options
}

// SetObjects : Allow user to set Objects
func (_options *AddImageTrainingDataOptions) SetObjects(objects []TrainingDataObject) *AddImageTrainingDataOptions {
	_options.Objects = objects
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *AddImageTrainingDataOptions) SetHeaders(param map[string]string) *AddImageTrainingDataOptions {
	options.Headers = param
	return options
}

// AddImagesOptions : The AddImages options.
type AddImagesOptions struct {
	// The identifier of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// An array of image files (.jpg or .png) or .zip files with images.
	// - Include a maximum of 20 images in a request.
	// - Limit the .zip file to 100 MB.
	// - Limit each image file to 10 MB.
	//
	// You can also include an image with the **image_url** parameter.
	ImagesFile []FileWithMetadata `json:"-"`

	// The array of URLs of image files (.jpg or .png).
	// - Include a maximum of 20 images in a request.
	// - Limit each image file to 10 MB.
	// - Minimum width and height is 30 pixels, but the service tends to perform better with images that are at least 300 x
	// 300 pixels. Maximum is 5400 pixels for either height or width.
	//
	// You can also include images with the **images_file** parameter.
	ImageURL []string `json:"-"`

	// Training data for a single image. Include training data only if you add one image with the request.
	//
	// The `object` property can contain alphanumeric, underscore, hyphen, space, and dot characters. It cannot begin with
	// the reserved prefix `sys-` and must be no longer than 32 characters.
	TrainingData *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAddImagesOptions : Instantiate AddImagesOptions
func (*VisualRecognitionV4) NewAddImagesOptions(collectionID string) *AddImagesOptions {
	return &AddImagesOptions{
		CollectionID: core.StringPtr(collectionID),
	}
}

// SetCollectionID : Allow user to set CollectionID
func (_options *AddImagesOptions) SetCollectionID(collectionID string) *AddImagesOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetImagesFile : Allow user to set ImagesFile
func (_options *AddImagesOptions) SetImagesFile(imagesFile []FileWithMetadata) *AddImagesOptions {
	_options.ImagesFile = imagesFile
	return _options
}

// SetImageURL : Allow user to set ImageURL
func (_options *AddImagesOptions) SetImageURL(imageURL []string) *AddImagesOptions {
	_options.ImageURL = imageURL
	return _options
}

// SetTrainingData : Allow user to set TrainingData
func (_options *AddImagesOptions) SetTrainingData(trainingData string) *AddImagesOptions {
	_options.TrainingData = core.StringPtr(trainingData)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *AddImagesOptions) SetHeaders(param map[string]string) *AddImagesOptions {
	options.Headers = param
	return options
}

// AnalyzeOptions : The Analyze options.
type AnalyzeOptions struct {
	// The IDs of the collections to analyze.
	CollectionIds []string `json:"-" validate:"required"`

	// The features to analyze.
	Features []string `json:"-" validate:"required"`

	// An array of image files (.jpg or .png) or .zip files with images.
	// - Include a maximum of 20 images in a request.
	// - Limit the .zip file to 100 MB.
	// - Limit each image file to 10 MB.
	//
	// You can also include an image with the **image_url** parameter.
	ImagesFile []FileWithMetadata `json:"-"`

	// An array of URLs of image files (.jpg or .png).
	// - Include a maximum of 20 images in a request.
	// - Limit each image file to 10 MB.
	// - Minimum width and height is 30 pixels, but the service tends to perform better with images that are at least 300 x
	// 300 pixels. Maximum is 5400 pixels for either height or width.
	//
	// You can also include images with the **images_file** parameter.
	ImageURL []string `json:"-"`

	// The minimum score a feature must have to be returned.
	Threshold *float32 `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the AnalyzeOptions.Features property.
const (
	AnalyzeOptionsFeaturesObjectsConst = "objects"
)

// NewAnalyzeOptions : Instantiate AnalyzeOptions
func (*VisualRecognitionV4) NewAnalyzeOptions(collectionIds []string, features []string) *AnalyzeOptions {
	return &AnalyzeOptions{
		CollectionIds: collectionIds,
		Features:      features,
	}
}

// SetCollectionIds : Allow user to set CollectionIds
func (_options *AnalyzeOptions) SetCollectionIds(collectionIds []string) *AnalyzeOptions {
	_options.CollectionIds = collectionIds
	return _options
}

// SetFeatures : Allow user to set Features
func (_options *AnalyzeOptions) SetFeatures(features []string) *AnalyzeOptions {
	_options.Features = features
	return _options
}

// SetImagesFile : Allow user to set ImagesFile
func (_options *AnalyzeOptions) SetImagesFile(imagesFile []FileWithMetadata) *AnalyzeOptions {
	_options.ImagesFile = imagesFile
	return _options
}

// SetImageURL : Allow user to set ImageURL
func (_options *AnalyzeOptions) SetImageURL(imageURL []string) *AnalyzeOptions {
	_options.ImageURL = imageURL
	return _options
}

// SetThreshold : Allow user to set Threshold
func (_options *AnalyzeOptions) SetThreshold(threshold float32) *AnalyzeOptions {
	_options.Threshold = core.Float32Ptr(threshold)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *AnalyzeOptions) SetHeaders(param map[string]string) *AnalyzeOptions {
	options.Headers = param
	return options
}

// AnalyzeResponse : Results for all images.
type AnalyzeResponse struct {
	// Analyzed images.
	Images []Image `json:"images" validate:"required"`

	// Information about what might cause less than optimal output.
	Warnings []Warning `json:"warnings,omitempty"`

	// A unique identifier of the request. Included only when an error or warning is returned.
	Trace *string `json:"trace,omitempty"`
}

// UnmarshalAnalyzeResponse unmarshals an instance of AnalyzeResponse from the specified map of raw messages.
func UnmarshalAnalyzeResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AnalyzeResponse)
	err = core.UnmarshalModel(m, "images", &obj.Images, UnmarshalImage)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "warnings", &obj.Warnings, UnmarshalWarning)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "trace", &obj.Trace)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Collection : Details about a collection.
type Collection struct {
	// The identifier of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// The name of the collection.
	Name *string `json:"name" validate:"required"`

	// The description of the collection.
	Description *string `json:"description" validate:"required"`

	// Date and time in Coordinated Universal Time (UTC) that the collection was created.
	Created *strfmt.DateTime `json:"created" validate:"required"`

	// Date and time in Coordinated Universal Time (UTC) that the collection was most recently updated.
	Updated *strfmt.DateTime `json:"updated" validate:"required"`

	// Number of images in the collection.
	ImageCount *int64 `json:"image_count" validate:"required"`

	// Training status information for the collection.
	TrainingStatus *CollectionTrainingStatus `json:"training_status" validate:"required"`
}

// UnmarshalCollection unmarshals an instance of Collection from the specified map of raw messages.
func UnmarshalCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Collection)
	err = core.UnmarshalPrimitive(m, "collection_id", &obj.CollectionID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created", &obj.Created)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated", &obj.Updated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "image_count", &obj.ImageCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "training_status", &obj.TrainingStatus, UnmarshalCollectionTrainingStatus)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CollectionObjects : The objects in a collection that are detected in an image.
type CollectionObjects struct {
	// The identifier of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// The identified objects in a collection.
	Objects []ObjectDetail `json:"objects" validate:"required"`
}

// UnmarshalCollectionObjects unmarshals an instance of CollectionObjects from the specified map of raw messages.
func UnmarshalCollectionObjects(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CollectionObjects)
	err = core.UnmarshalPrimitive(m, "collection_id", &obj.CollectionID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "objects", &obj.Objects, UnmarshalObjectDetail)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CollectionTrainingStatus : Training status information for the collection.
type CollectionTrainingStatus struct {
	// Training status for the objects in the collection.
	Objects *ObjectTrainingStatus `json:"objects" validate:"required"`
}

// UnmarshalCollectionTrainingStatus unmarshals an instance of CollectionTrainingStatus from the specified map of raw messages.
func UnmarshalCollectionTrainingStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CollectionTrainingStatus)
	err = core.UnmarshalModel(m, "objects", &obj.Objects, UnmarshalObjectTrainingStatus)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CollectionsList : A container for the list of collections.
type CollectionsList struct {
	// The collections in this service instance.
	Collections []Collection `json:"collections" validate:"required"`
}

// UnmarshalCollectionsList unmarshals an instance of CollectionsList from the specified map of raw messages.
func UnmarshalCollectionsList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CollectionsList)
	err = core.UnmarshalModel(m, "collections", &obj.Collections, UnmarshalCollection)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateCollectionOptions : The CreateCollection options.
type CreateCollectionOptions struct {
	// The name of the collection. The name can contain alphanumeric, underscore, hyphen, and dot characters. It cannot
	// begin with the reserved prefix `sys-`.
	Name *string `json:"name,omitempty"`

	// The description of the collection.
	Description *string `json:"description,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateCollectionOptions : Instantiate CreateCollectionOptions
func (*VisualRecognitionV4) NewCreateCollectionOptions() *CreateCollectionOptions {
	return &CreateCollectionOptions{}
}

// SetName : Allow user to set Name
func (_options *CreateCollectionOptions) SetName(name string) *CreateCollectionOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateCollectionOptions) SetDescription(description string) *CreateCollectionOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateCollectionOptions) SetHeaders(param map[string]string) *CreateCollectionOptions {
	options.Headers = param
	return options
}

// DeleteCollectionOptions : The DeleteCollection options.
type DeleteCollectionOptions struct {
	// The identifier of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteCollectionOptions : Instantiate DeleteCollectionOptions
func (*VisualRecognitionV4) NewDeleteCollectionOptions(collectionID string) *DeleteCollectionOptions {
	return &DeleteCollectionOptions{
		CollectionID: core.StringPtr(collectionID),
	}
}

// SetCollectionID : Allow user to set CollectionID
func (_options *DeleteCollectionOptions) SetCollectionID(collectionID string) *DeleteCollectionOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteCollectionOptions) SetHeaders(param map[string]string) *DeleteCollectionOptions {
	options.Headers = param
	return options
}

// DeleteImageOptions : The DeleteImage options.
type DeleteImageOptions struct {
	// The identifier of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// The identifier of the image.
	ImageID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteImageOptions : Instantiate DeleteImageOptions
func (*VisualRecognitionV4) NewDeleteImageOptions(collectionID string, imageID string) *DeleteImageOptions {
	return &DeleteImageOptions{
		CollectionID: core.StringPtr(collectionID),
		ImageID:      core.StringPtr(imageID),
	}
}

// SetCollectionID : Allow user to set CollectionID
func (_options *DeleteImageOptions) SetCollectionID(collectionID string) *DeleteImageOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetImageID : Allow user to set ImageID
func (_options *DeleteImageOptions) SetImageID(imageID string) *DeleteImageOptions {
	_options.ImageID = core.StringPtr(imageID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteImageOptions) SetHeaders(param map[string]string) *DeleteImageOptions {
	options.Headers = param
	return options
}

// DeleteObjectOptions : The DeleteObject options.
type DeleteObjectOptions struct {
	// The identifier of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// The name of the object.
	Object *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteObjectOptions : Instantiate DeleteObjectOptions
func (*VisualRecognitionV4) NewDeleteObjectOptions(collectionID string, object string) *DeleteObjectOptions {
	return &DeleteObjectOptions{
		CollectionID: core.StringPtr(collectionID),
		Object:       core.StringPtr(object),
	}
}

// SetCollectionID : Allow user to set CollectionID
func (_options *DeleteObjectOptions) SetCollectionID(collectionID string) *DeleteObjectOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetObject : Allow user to set Object
func (_options *DeleteObjectOptions) SetObject(object string) *DeleteObjectOptions {
	_options.Object = core.StringPtr(object)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteObjectOptions) SetHeaders(param map[string]string) *DeleteObjectOptions {
	options.Headers = param
	return options
}

// DeleteUserDataOptions : The DeleteUserData options.
type DeleteUserDataOptions struct {
	// The customer ID for which all data is to be deleted.
	CustomerID *string `json:"-" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteUserDataOptions : Instantiate DeleteUserDataOptions
func (*VisualRecognitionV4) NewDeleteUserDataOptions(customerID string) *DeleteUserDataOptions {
	return &DeleteUserDataOptions{
		CustomerID: core.StringPtr(customerID),
	}
}

// SetCustomerID : Allow user to set CustomerID
func (_options *DeleteUserDataOptions) SetCustomerID(customerID string) *DeleteUserDataOptions {
	_options.CustomerID = core.StringPtr(customerID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteUserDataOptions) SetHeaders(param map[string]string) *DeleteUserDataOptions {
	options.Headers = param
	return options
}

// DetectedObjects : Container for the list of collections that have objects detected in an image.
type DetectedObjects struct {
	// The collections with identified objects.
	Collections []CollectionObjects `json:"collections,omitempty"`
}

// UnmarshalDetectedObjects unmarshals an instance of DetectedObjects from the specified map of raw messages.
func UnmarshalDetectedObjects(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DetectedObjects)
	err = core.UnmarshalModel(m, "collections", &obj.Collections, UnmarshalCollectionObjects)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Error : Details about an error.
type Error struct {
	// Identifier of the problem.
	Code *string `json:"code" validate:"required"`

	// An explanation of the problem with possible solutions.
	Message *string `json:"message" validate:"required"`

	// A URL for more information about the solution.
	MoreInfo *string `json:"more_info,omitempty"`

	// Details about the specific area of the problem.
	Target *ErrorTarget `json:"target,omitempty"`
}

// Constants associated with the Error.Code property.
// Identifier of the problem.
const (
	ErrorCodeInvalidFieldConst  = "invalid_field"
	ErrorCodeInvalidHeaderConst = "invalid_header"
	ErrorCodeInvalidMethodConst = "invalid_method"
	ErrorCodeMissingFieldConst  = "missing_field"
	ErrorCodeServerErrorConst   = "server_error"
)

// UnmarshalError unmarshals an instance of Error from the specified map of raw messages.
func UnmarshalError(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Error)
	err = core.UnmarshalPrimitive(m, "code", &obj.Code)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "more_info", &obj.MoreInfo)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "target", &obj.Target, UnmarshalErrorTarget)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ErrorTarget : Details about the specific area of the problem.
type ErrorTarget struct {
	// The parameter or property that is the focus of the problem.
	Type *string `json:"type" validate:"required"`

	// The property that is identified with the problem.
	Name *string `json:"name" validate:"required"`
}

// Constants associated with the ErrorTarget.Type property.
// The parameter or property that is the focus of the problem.
const (
	ErrorTargetTypeFieldConst     = "field"
	ErrorTargetTypeHeaderConst    = "header"
	ErrorTargetTypeParameterConst = "parameter"
)

// UnmarshalErrorTarget unmarshals an instance of ErrorTarget from the specified map of raw messages.
func UnmarshalErrorTarget(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ErrorTarget)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FileWithMetadata : A file with its associated metadata.
type FileWithMetadata struct {
	// The data / content for the file.
	Data io.ReadCloser `json:"data" validate:"required"`

	// The filename of the file.
	Filename *string `json:"filename,omitempty"`

	// The content type of the file.
	ContentType *string `json:"content_type,omitempty"`
}

// NewFileWithMetadata : Instantiate FileWithMetadata (Generic Model Constructor)
func (*VisualRecognitionV4) NewFileWithMetadata(data io.ReadCloser) (_model *FileWithMetadata, err error) {
	_model = &FileWithMetadata{
		Data: data,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalFileWithMetadata unmarshals an instance of FileWithMetadata from the specified map of raw messages.
func UnmarshalFileWithMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(core.FileWithMetadata)
	err = core.UnmarshalFileWithMetadata(m, &obj)
	if err != nil {
		return
	}

	// do a simple conversion from the core type to the service type
	// they have identical fields
	convertedModel := FileWithMetadata(*obj)
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(&convertedModel))

	return
}

// GetCollectionOptions : The GetCollection options.
type GetCollectionOptions struct {
	// The identifier of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCollectionOptions : Instantiate GetCollectionOptions
func (*VisualRecognitionV4) NewGetCollectionOptions(collectionID string) *GetCollectionOptions {
	return &GetCollectionOptions{
		CollectionID: core.StringPtr(collectionID),
	}
}

// SetCollectionID : Allow user to set CollectionID
func (_options *GetCollectionOptions) SetCollectionID(collectionID string) *GetCollectionOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetCollectionOptions) SetHeaders(param map[string]string) *GetCollectionOptions {
	options.Headers = param
	return options
}

// GetImageDetailsOptions : The GetImageDetails options.
type GetImageDetailsOptions struct {
	// The identifier of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// The identifier of the image.
	ImageID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetImageDetailsOptions : Instantiate GetImageDetailsOptions
func (*VisualRecognitionV4) NewGetImageDetailsOptions(collectionID string, imageID string) *GetImageDetailsOptions {
	return &GetImageDetailsOptions{
		CollectionID: core.StringPtr(collectionID),
		ImageID:      core.StringPtr(imageID),
	}
}

// SetCollectionID : Allow user to set CollectionID
func (_options *GetImageDetailsOptions) SetCollectionID(collectionID string) *GetImageDetailsOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetImageID : Allow user to set ImageID
func (_options *GetImageDetailsOptions) SetImageID(imageID string) *GetImageDetailsOptions {
	_options.ImageID = core.StringPtr(imageID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetImageDetailsOptions) SetHeaders(param map[string]string) *GetImageDetailsOptions {
	options.Headers = param
	return options
}

// GetJpegImageOptions : The GetJpegImage options.
type GetJpegImageOptions struct {
	// The identifier of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// The identifier of the image.
	ImageID *string `json:"-" validate:"required,ne="`

	// The image size. Specify `thumbnail` to return a version that maintains the original aspect ratio but is no larger
	// than 200 pixels in the larger dimension. For example, an original 800 x 1000 image is resized to 160 x 200 pixels.
	Size *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetJpegImageOptions.Size property.
// The image size. Specify `thumbnail` to return a version that maintains the original aspect ratio but is no larger
// than 200 pixels in the larger dimension. For example, an original 800 x 1000 image is resized to 160 x 200 pixels.
const (
	GetJpegImageOptionsSizeFullConst      = "full"
	GetJpegImageOptionsSizeThumbnailConst = "thumbnail"
)

// NewGetJpegImageOptions : Instantiate GetJpegImageOptions
func (*VisualRecognitionV4) NewGetJpegImageOptions(collectionID string, imageID string) *GetJpegImageOptions {
	return &GetJpegImageOptions{
		CollectionID: core.StringPtr(collectionID),
		ImageID:      core.StringPtr(imageID),
	}
}

// SetCollectionID : Allow user to set CollectionID
func (_options *GetJpegImageOptions) SetCollectionID(collectionID string) *GetJpegImageOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetImageID : Allow user to set ImageID
func (_options *GetJpegImageOptions) SetImageID(imageID string) *GetJpegImageOptions {
	_options.ImageID = core.StringPtr(imageID)
	return _options
}

// SetSize : Allow user to set Size
func (_options *GetJpegImageOptions) SetSize(size string) *GetJpegImageOptions {
	_options.Size = core.StringPtr(size)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetJpegImageOptions) SetHeaders(param map[string]string) *GetJpegImageOptions {
	options.Headers = param
	return options
}

// GetModelFileOptions : The GetModelFile options.
type GetModelFileOptions struct {
	// The identifier of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// The feature for the model.
	Feature *string `json:"-" validate:"required"`

	// The format of the returned model.
	ModelFormat *string `json:"-" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetModelFileOptions.Feature property.
// The feature for the model.
const (
	GetModelFileOptionsFeatureObjectsConst = "objects"
)

// Constants associated with the GetModelFileOptions.ModelFormat property.
// The format of the returned model.
const (
	GetModelFileOptionsModelFormatRscnnConst = "rscnn"
)

// NewGetModelFileOptions : Instantiate GetModelFileOptions
func (*VisualRecognitionV4) NewGetModelFileOptions(collectionID string, feature string, modelFormat string) *GetModelFileOptions {
	return &GetModelFileOptions{
		CollectionID: core.StringPtr(collectionID),
		Feature:      core.StringPtr(feature),
		ModelFormat:  core.StringPtr(modelFormat),
	}
}

// SetCollectionID : Allow user to set CollectionID
func (_options *GetModelFileOptions) SetCollectionID(collectionID string) *GetModelFileOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetFeature : Allow user to set Feature
func (_options *GetModelFileOptions) SetFeature(feature string) *GetModelFileOptions {
	_options.Feature = core.StringPtr(feature)
	return _options
}

// SetModelFormat : Allow user to set ModelFormat
func (_options *GetModelFileOptions) SetModelFormat(modelFormat string) *GetModelFileOptions {
	_options.ModelFormat = core.StringPtr(modelFormat)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetModelFileOptions) SetHeaders(param map[string]string) *GetModelFileOptions {
	options.Headers = param
	return options
}

// GetObjectMetadataOptions : The GetObjectMetadata options.
type GetObjectMetadataOptions struct {
	// The identifier of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// The name of the object.
	Object *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetObjectMetadataOptions : Instantiate GetObjectMetadataOptions
func (*VisualRecognitionV4) NewGetObjectMetadataOptions(collectionID string, object string) *GetObjectMetadataOptions {
	return &GetObjectMetadataOptions{
		CollectionID: core.StringPtr(collectionID),
		Object:       core.StringPtr(object),
	}
}

// SetCollectionID : Allow user to set CollectionID
func (_options *GetObjectMetadataOptions) SetCollectionID(collectionID string) *GetObjectMetadataOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetObject : Allow user to set Object
func (_options *GetObjectMetadataOptions) SetObject(object string) *GetObjectMetadataOptions {
	_options.Object = core.StringPtr(object)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetObjectMetadataOptions) SetHeaders(param map[string]string) *GetObjectMetadataOptions {
	options.Headers = param
	return options
}

// GetTrainingUsageOptions : The GetTrainingUsage options.
type GetTrainingUsageOptions struct {
	// The earliest day to include training events. Specify dates in YYYY-MM-DD format. If empty or not specified, the
	// earliest training event is included.
	StartTime *strfmt.Date `json:"-"`

	// The most recent day to include training events. Specify dates in YYYY-MM-DD format. All events for the day are
	// included. If empty or not specified, the current day is used. Specify the same value as `start_time` to request
	// events for a single day.
	EndTime *strfmt.Date `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetTrainingUsageOptions : Instantiate GetTrainingUsageOptions
func (*VisualRecognitionV4) NewGetTrainingUsageOptions() *GetTrainingUsageOptions {
	return &GetTrainingUsageOptions{}
}

// SetStartTime : Allow user to set StartTime
func (_options *GetTrainingUsageOptions) SetStartTime(startTime *strfmt.Date) *GetTrainingUsageOptions {
	_options.StartTime = startTime
	return _options
}

// SetEndTime : Allow user to set EndTime
func (_options *GetTrainingUsageOptions) SetEndTime(endTime *strfmt.Date) *GetTrainingUsageOptions {
	_options.EndTime = endTime
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetTrainingUsageOptions) SetHeaders(param map[string]string) *GetTrainingUsageOptions {
	options.Headers = param
	return options
}

// Image : Details about an image.
type Image struct {
	// The source type of the image.
	Source *ImageSource `json:"source" validate:"required"`

	// Height and width of an image.
	Dimensions *ImageDimensions `json:"dimensions" validate:"required"`

	// Container for the list of collections that have objects detected in an image.
	Objects *DetectedObjects `json:"objects" validate:"required"`

	// A container for the problems in the request.
	Errors []Error `json:"errors,omitempty"`
}

// UnmarshalImage unmarshals an instance of Image from the specified map of raw messages.
func UnmarshalImage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Image)
	err = core.UnmarshalModel(m, "source", &obj.Source, UnmarshalImageSource)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "dimensions", &obj.Dimensions, UnmarshalImageDimensions)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "objects", &obj.Objects, UnmarshalDetectedObjects)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "errors", &obj.Errors, UnmarshalError)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ImageDetails : Details about an image.
type ImageDetails struct {
	// The identifier of the image.
	ImageID *string `json:"image_id,omitempty"`

	// Date and time in Coordinated Universal Time (UTC) that the image was most recently updated.
	Updated *strfmt.DateTime `json:"updated,omitempty"`

	// Date and time in Coordinated Universal Time (UTC) that the image was created.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// The source type of the image.
	Source *ImageSource `json:"source" validate:"required"`

	// Height and width of an image.
	Dimensions *ImageDimensions `json:"dimensions,omitempty"`

	// Details about the errors.
	Errors []Error `json:"errors,omitempty"`

	// Training data for all objects.
	TrainingData *TrainingDataObjects `json:"training_data,omitempty"`
}

// UnmarshalImageDetails unmarshals an instance of ImageDetails from the specified map of raw messages.
func UnmarshalImageDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ImageDetails)
	err = core.UnmarshalPrimitive(m, "image_id", &obj.ImageID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated", &obj.Updated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created", &obj.Created)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "source", &obj.Source, UnmarshalImageSource)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "dimensions", &obj.Dimensions, UnmarshalImageDimensions)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "errors", &obj.Errors, UnmarshalError)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "training_data", &obj.TrainingData, UnmarshalTrainingDataObjects)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ImageDetailsList : List of information about the images.
type ImageDetailsList struct {
	// The images in the collection.
	Images []ImageDetails `json:"images,omitempty"`

	// Information about what might cause less than optimal output.
	Warnings []Warning `json:"warnings,omitempty"`

	// A unique identifier of the request. Included only when an error or warning is returned.
	Trace *string `json:"trace,omitempty"`
}

// UnmarshalImageDetailsList unmarshals an instance of ImageDetailsList from the specified map of raw messages.
func UnmarshalImageDetailsList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ImageDetailsList)
	err = core.UnmarshalModel(m, "images", &obj.Images, UnmarshalImageDetails)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "warnings", &obj.Warnings, UnmarshalWarning)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "trace", &obj.Trace)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ImageDimensions : Height and width of an image.
type ImageDimensions struct {
	// Height in pixels of the image.
	Height *int64 `json:"height,omitempty"`

	// Width in pixels of the image.
	Width *int64 `json:"width,omitempty"`
}

// UnmarshalImageDimensions unmarshals an instance of ImageDimensions from the specified map of raw messages.
func UnmarshalImageDimensions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ImageDimensions)
	err = core.UnmarshalPrimitive(m, "height", &obj.Height)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "width", &obj.Width)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ImageSource : The source type of the image.
type ImageSource struct {
	// The source type of the image.
	Type *string `json:"type" validate:"required"`

	// Name of the image file if uploaded. Not returned when the image is passed by URL.
	Filename *string `json:"filename,omitempty"`

	// Name of the .zip file of images if uploaded. Not returned when the image is passed directly or by URL.
	ArchiveFilename *string `json:"archive_filename,omitempty"`

	// Source of the image before any redirects. Not returned when the image is uploaded.
	SourceURL *string `json:"source_url,omitempty"`

	// Fully resolved URL of the image after redirects are followed. Not returned when the image is uploaded.
	ResolvedURL *string `json:"resolved_url,omitempty"`
}

// Constants associated with the ImageSource.Type property.
// The source type of the image.
const (
	ImageSourceTypeFileConst = "file"
	ImageSourceTypeURLConst  = "url"
)

// UnmarshalImageSource unmarshals an instance of ImageSource from the specified map of raw messages.
func UnmarshalImageSource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ImageSource)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "filename", &obj.Filename)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "archive_filename", &obj.ArchiveFilename)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source_url", &obj.SourceURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resolved_url", &obj.ResolvedURL)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ImageSummary : Basic information about an image.
type ImageSummary struct {
	// The identifier of the image.
	ImageID *string `json:"image_id,omitempty"`

	// Date and time in Coordinated Universal Time (UTC) that the image was most recently updated.
	Updated *strfmt.DateTime `json:"updated,omitempty"`
}

// UnmarshalImageSummary unmarshals an instance of ImageSummary from the specified map of raw messages.
func UnmarshalImageSummary(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ImageSummary)
	err = core.UnmarshalPrimitive(m, "image_id", &obj.ImageID)
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

// ImageSummaryList : List of images.
type ImageSummaryList struct {
	// The images in the collection.
	Images []ImageSummary `json:"images" validate:"required"`
}

// UnmarshalImageSummaryList unmarshals an instance of ImageSummaryList from the specified map of raw messages.
func UnmarshalImageSummaryList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ImageSummaryList)
	err = core.UnmarshalModel(m, "images", &obj.Images, UnmarshalImageSummary)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListCollectionsOptions : The ListCollections options.
type ListCollectionsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListCollectionsOptions : Instantiate ListCollectionsOptions
func (*VisualRecognitionV4) NewListCollectionsOptions() *ListCollectionsOptions {
	return &ListCollectionsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListCollectionsOptions) SetHeaders(param map[string]string) *ListCollectionsOptions {
	options.Headers = param
	return options
}

// ListImagesOptions : The ListImages options.
type ListImagesOptions struct {
	// The identifier of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListImagesOptions : Instantiate ListImagesOptions
func (*VisualRecognitionV4) NewListImagesOptions(collectionID string) *ListImagesOptions {
	return &ListImagesOptions{
		CollectionID: core.StringPtr(collectionID),
	}
}

// SetCollectionID : Allow user to set CollectionID
func (_options *ListImagesOptions) SetCollectionID(collectionID string) *ListImagesOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListImagesOptions) SetHeaders(param map[string]string) *ListImagesOptions {
	options.Headers = param
	return options
}

// ListObjectMetadataOptions : The ListObjectMetadata options.
type ListObjectMetadataOptions struct {
	// The identifier of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListObjectMetadataOptions : Instantiate ListObjectMetadataOptions
func (*VisualRecognitionV4) NewListObjectMetadataOptions(collectionID string) *ListObjectMetadataOptions {
	return &ListObjectMetadataOptions{
		CollectionID: core.StringPtr(collectionID),
	}
}

// SetCollectionID : Allow user to set CollectionID
func (_options *ListObjectMetadataOptions) SetCollectionID(collectionID string) *ListObjectMetadataOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListObjectMetadataOptions) SetHeaders(param map[string]string) *ListObjectMetadataOptions {
	options.Headers = param
	return options
}

// Location : Defines the location of the bounding box around the object.
type Location struct {
	// Y-position of top-left pixel of the bounding box.
	Top *int64 `json:"top" validate:"required"`

	// X-position of top-left pixel of the bounding box.
	Left *int64 `json:"left" validate:"required"`

	// Width in pixels of of the bounding box.
	Width *int64 `json:"width" validate:"required"`

	// Height in pixels of the bounding box.
	Height *int64 `json:"height" validate:"required"`
}

// NewLocation : Instantiate Location (Generic Model Constructor)
func (*VisualRecognitionV4) NewLocation(top int64, left int64, width int64, height int64) (_model *Location, err error) {
	_model = &Location{
		Top:    core.Int64Ptr(top),
		Left:   core.Int64Ptr(left),
		Width:  core.Int64Ptr(width),
		Height: core.Int64Ptr(height),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalLocation unmarshals an instance of Location from the specified map of raw messages.
func UnmarshalLocation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Location)
	err = core.UnmarshalPrimitive(m, "top", &obj.Top)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "left", &obj.Left)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "width", &obj.Width)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "height", &obj.Height)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ObjectDetail : Details about an object in the collection.
type ObjectDetail struct {
	// The label for the object.
	Object *string `json:"object" validate:"required"`

	// Defines the location of the bounding box around the object.
	Location *ObjectDetailLocation `json:"location" validate:"required"`

	// Confidence score for the object in the range of 0 to 1. A higher score indicates greater likelihood that the object
	// is depicted at this location in the image.
	Score *float32 `json:"score" validate:"required"`
}

// UnmarshalObjectDetail unmarshals an instance of ObjectDetail from the specified map of raw messages.
func UnmarshalObjectDetail(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ObjectDetail)
	err = core.UnmarshalPrimitive(m, "object", &obj.Object)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "location", &obj.Location, UnmarshalObjectDetailLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "score", &obj.Score)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ObjectDetailLocation : Defines the location of the bounding box around the object.
type ObjectDetailLocation struct {
	// Y-position of top-left pixel of the bounding box.
	Top *int64 `json:"top" validate:"required"`

	// X-position of top-left pixel of the bounding box.
	Left *int64 `json:"left" validate:"required"`

	// Width in pixels of of the bounding box.
	Width *int64 `json:"width" validate:"required"`

	// Height in pixels of the bounding box.
	Height *int64 `json:"height" validate:"required"`
}

// UnmarshalObjectDetailLocation unmarshals an instance of ObjectDetailLocation from the specified map of raw messages.
func UnmarshalObjectDetailLocation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ObjectDetailLocation)
	err = core.UnmarshalPrimitive(m, "top", &obj.Top)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "left", &obj.Left)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "width", &obj.Width)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "height", &obj.Height)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ObjectMetadata : Basic information about an object.
type ObjectMetadata struct {
	// The name of the object.
	Object *string `json:"object,omitempty"`

	// Number of bounding boxes with this object name in the collection.
	Count *int64 `json:"count,omitempty"`
}

// UnmarshalObjectMetadata unmarshals an instance of ObjectMetadata from the specified map of raw messages.
func UnmarshalObjectMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ObjectMetadata)
	err = core.UnmarshalPrimitive(m, "object", &obj.Object)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "count", &obj.Count)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ObjectMetadataList : List of objects.
type ObjectMetadataList struct {
	// Number of unique named objects in the collection.
	ObjectCount *int64 `json:"object_count" validate:"required"`

	// The objects in the collection.
	Objects []ObjectMetadata `json:"objects,omitempty"`
}

// UnmarshalObjectMetadataList unmarshals an instance of ObjectMetadataList from the specified map of raw messages.
func UnmarshalObjectMetadataList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ObjectMetadataList)
	err = core.UnmarshalPrimitive(m, "object_count", &obj.ObjectCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "objects", &obj.Objects, UnmarshalObjectMetadata)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ObjectTrainingStatus : Training status for the objects in the collection.
type ObjectTrainingStatus struct {
	// Whether you can analyze images in the collection with the **objects** feature.
	Ready *bool `json:"ready" validate:"required"`

	// Whether training is in progress.
	InProgress *bool `json:"in_progress" validate:"required"`

	// Whether there are changes to the training data since the most recent training.
	DataChanged *bool `json:"data_changed" validate:"required"`

	// Whether the most recent training failed.
	LatestFailed *bool `json:"latest_failed" validate:"required"`

	// Whether the model can be downloaded after the training status is `ready`.
	RscnnReady *bool `json:"rscnn_ready" validate:"required"`

	// Details about the training. If training is in progress, includes information about the status. If training is not in
	// progress, includes a success message or information about why training failed.
	Description *string `json:"description" validate:"required"`
}

// NewObjectTrainingStatus : Instantiate ObjectTrainingStatus (Generic Model Constructor)
func (*VisualRecognitionV4) NewObjectTrainingStatus(ready bool, inProgress bool, dataChanged bool, latestFailed bool, rscnnReady bool, description string) (_model *ObjectTrainingStatus, err error) {
	_model = &ObjectTrainingStatus{
		Ready:        core.BoolPtr(ready),
		InProgress:   core.BoolPtr(inProgress),
		DataChanged:  core.BoolPtr(dataChanged),
		LatestFailed: core.BoolPtr(latestFailed),
		RscnnReady:   core.BoolPtr(rscnnReady),
		Description:  core.StringPtr(description),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalObjectTrainingStatus unmarshals an instance of ObjectTrainingStatus from the specified map of raw messages.
func UnmarshalObjectTrainingStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ObjectTrainingStatus)
	err = core.UnmarshalPrimitive(m, "ready", &obj.Ready)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "in_progress", &obj.InProgress)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "data_changed", &obj.DataChanged)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "latest_failed", &obj.LatestFailed)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rscnn_ready", &obj.RscnnReady)
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

// TrainOptions : The Train options.
type TrainOptions struct {
	// The identifier of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewTrainOptions : Instantiate TrainOptions
func (*VisualRecognitionV4) NewTrainOptions(collectionID string) *TrainOptions {
	return &TrainOptions{
		CollectionID: core.StringPtr(collectionID),
	}
}

// SetCollectionID : Allow user to set CollectionID
func (_options *TrainOptions) SetCollectionID(collectionID string) *TrainOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *TrainOptions) SetHeaders(param map[string]string) *TrainOptions {
	options.Headers = param
	return options
}

// TrainingDataObject : Details about the training data.
type TrainingDataObject struct {
	// The name of the object.
	Object *string `json:"object,omitempty"`

	// Defines the location of the bounding box around the object.
	Location *Location `json:"location,omitempty"`
}

// UnmarshalTrainingDataObject unmarshals an instance of TrainingDataObject from the specified map of raw messages.
func UnmarshalTrainingDataObject(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TrainingDataObject)
	err = core.UnmarshalPrimitive(m, "object", &obj.Object)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "location", &obj.Location, UnmarshalLocation)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TrainingDataObjects : Training data for all objects.
type TrainingDataObjects struct {
	// Training data for specific objects.
	Objects []TrainingDataObject `json:"objects,omitempty"`
}

// UnmarshalTrainingDataObjects unmarshals an instance of TrainingDataObjects from the specified map of raw messages.
func UnmarshalTrainingDataObjects(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TrainingDataObjects)
	err = core.UnmarshalModel(m, "objects", &obj.Objects, UnmarshalTrainingDataObject)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TrainingEvent : Details about the training event.
type TrainingEvent struct {
	// Trained object type. Only `objects` is currently supported.
	Type *string `json:"type,omitempty"`

	// Identifier of the trained collection.
	CollectionID *string `json:"collection_id,omitempty"`

	// Date and time in Coordinated Universal Time (UTC) that training on the collection finished.
	CompletionTime *strfmt.DateTime `json:"completion_time,omitempty"`

	// Training status of the training event.
	Status *string `json:"status,omitempty"`

	// The total number of images that were used in training for this training event.
	ImageCount *int64 `json:"image_count,omitempty"`
}

// Constants associated with the TrainingEvent.Type property.
// Trained object type. Only `objects` is currently supported.
const (
	TrainingEventTypeObjectsConst = "objects"
)

// Constants associated with the TrainingEvent.Status property.
// Training status of the training event.
const (
	TrainingEventStatusFailedConst    = "failed"
	TrainingEventStatusSucceededConst = "succeeded"
)

// UnmarshalTrainingEvent unmarshals an instance of TrainingEvent from the specified map of raw messages.
func UnmarshalTrainingEvent(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TrainingEvent)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "collection_id", &obj.CollectionID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "completion_time", &obj.CompletionTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "image_count", &obj.ImageCount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TrainingEvents : Details about the training events.
type TrainingEvents struct {
	// The starting day for the returned training events in Coordinated Universal Time (UTC). If not specified in the
	// request, it identifies the earliest training event.
	StartTime *strfmt.DateTime `json:"start_time,omitempty"`

	// The ending day for the returned training events in Coordinated Universal Time (UTC). If not specified in the
	// request, it lists the current time.
	EndTime *strfmt.DateTime `json:"end_time,omitempty"`

	// The total number of training events in the response for the start and end times.
	CompletedEvents *int64 `json:"completed_events,omitempty"`

	// The total number of images that were used in training for the start and end times.
	TrainedImages *int64 `json:"trained_images,omitempty"`

	// The completed training events for the start and end time.
	Events []TrainingEvent `json:"events,omitempty"`
}

// UnmarshalTrainingEvents unmarshals an instance of TrainingEvents from the specified map of raw messages.
func UnmarshalTrainingEvents(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TrainingEvents)
	err = core.UnmarshalPrimitive(m, "start_time", &obj.StartTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "end_time", &obj.EndTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "completed_events", &obj.CompletedEvents)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "trained_images", &obj.TrainedImages)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "events", &obj.Events, UnmarshalTrainingEvent)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TrainingStatus : Training status information for the collection.
type TrainingStatus struct {
	// Training status for the objects in the collection.
	Objects *ObjectTrainingStatus `json:"objects" validate:"required"`
}

// NewTrainingStatus : Instantiate TrainingStatus (Generic Model Constructor)
func (*VisualRecognitionV4) NewTrainingStatus(objects *ObjectTrainingStatus) (_model *TrainingStatus, err error) {
	_model = &TrainingStatus{
		Objects: objects,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalTrainingStatus unmarshals an instance of TrainingStatus from the specified map of raw messages.
func UnmarshalTrainingStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TrainingStatus)
	err = core.UnmarshalModel(m, "objects", &obj.Objects, UnmarshalObjectTrainingStatus)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateCollectionOptions : The UpdateCollection options.
type UpdateCollectionOptions struct {
	// The identifier of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// The name of the collection. The name can contain alphanumeric, underscore, hyphen, and dot characters. It cannot
	// begin with the reserved prefix `sys-`.
	Name *string `json:"name,omitempty"`

	// The description of the collection.
	Description *string `json:"description,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateCollectionOptions : Instantiate UpdateCollectionOptions
func (*VisualRecognitionV4) NewUpdateCollectionOptions(collectionID string) *UpdateCollectionOptions {
	return &UpdateCollectionOptions{
		CollectionID: core.StringPtr(collectionID),
	}
}

// SetCollectionID : Allow user to set CollectionID
func (_options *UpdateCollectionOptions) SetCollectionID(collectionID string) *UpdateCollectionOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetName : Allow user to set Name
func (_options *UpdateCollectionOptions) SetName(name string) *UpdateCollectionOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *UpdateCollectionOptions) SetDescription(description string) *UpdateCollectionOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateCollectionOptions) SetHeaders(param map[string]string) *UpdateCollectionOptions {
	options.Headers = param
	return options
}

// UpdateObjectMetadata : Basic information about an updated object.
type UpdateObjectMetadata struct {
	// The updated name of the object. The name can contain alphanumeric, underscore, hyphen, space, and dot characters. It
	// cannot begin with the reserved prefix `sys-`.
	Object *string `json:"object" validate:"required"`

	// Number of bounding boxes in the collection with the updated object name.
	Count *int64 `json:"count,omitempty"`
}

// NewUpdateObjectMetadata : Instantiate UpdateObjectMetadata (Generic Model Constructor)
func (*VisualRecognitionV4) NewUpdateObjectMetadata(object string) (_model *UpdateObjectMetadata, err error) {
	_model = &UpdateObjectMetadata{
		Object: core.StringPtr(object),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalUpdateObjectMetadata unmarshals an instance of UpdateObjectMetadata from the specified map of raw messages.
func UnmarshalUpdateObjectMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UpdateObjectMetadata)
	err = core.UnmarshalPrimitive(m, "object", &obj.Object)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "count", &obj.Count)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateObjectMetadataOptions : The UpdateObjectMetadata options.
type UpdateObjectMetadataOptions struct {
	// The identifier of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// The name of the object.
	Object *string `json:"-" validate:"required,ne="`

	// The updated name of the object. The name can contain alphanumeric, underscore, hyphen, space, and dot characters. It
	// cannot begin with the reserved prefix `sys-`.
	NewObject *string `json:"object" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateObjectMetadataOptions : Instantiate UpdateObjectMetadataOptions
func (*VisualRecognitionV4) NewUpdateObjectMetadataOptions(collectionID string, object string, newObject string) *UpdateObjectMetadataOptions {
	return &UpdateObjectMetadataOptions{
		CollectionID: core.StringPtr(collectionID),
		Object:       core.StringPtr(object),
		NewObject:    core.StringPtr(newObject),
	}
}

// SetCollectionID : Allow user to set CollectionID
func (_options *UpdateObjectMetadataOptions) SetCollectionID(collectionID string) *UpdateObjectMetadataOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetObject : Allow user to set Object
func (_options *UpdateObjectMetadataOptions) SetObject(object string) *UpdateObjectMetadataOptions {
	_options.Object = core.StringPtr(object)
	return _options
}

// SetNewObject : Allow user to set NewObject
func (_options *UpdateObjectMetadataOptions) SetNewObject(newObject string) *UpdateObjectMetadataOptions {
	_options.NewObject = core.StringPtr(newObject)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateObjectMetadataOptions) SetHeaders(param map[string]string) *UpdateObjectMetadataOptions {
	options.Headers = param
	return options
}

// Warning : Details about a problem.
type Warning struct {
	// Identifier of the problem.
	Code *string `json:"code" validate:"required"`

	// An explanation of the problem with possible solutions.
	Message *string `json:"message" validate:"required"`

	// A URL for more information about the solution.
	MoreInfo *string `json:"more_info,omitempty"`
}

// Constants associated with the Warning.Code property.
// Identifier of the problem.
const (
	WarningCodeInvalidFieldConst  = "invalid_field"
	WarningCodeInvalidHeaderConst = "invalid_header"
	WarningCodeInvalidMethodConst = "invalid_method"
	WarningCodeMissingFieldConst  = "missing_field"
	WarningCodeServerErrorConst   = "server_error"
)

// UnmarshalWarning unmarshals an instance of Warning from the specified map of raw messages.
func UnmarshalWarning(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Warning)
	err = core.UnmarshalPrimitive(m, "code", &obj.Code)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "more_info", &obj.MoreInfo)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
