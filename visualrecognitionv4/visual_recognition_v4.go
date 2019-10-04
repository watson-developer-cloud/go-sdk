/**
 * (C) Copyright IBM Corp. 2019.
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

// Package visualrecognitionv4 : Operations and models for the VisualRecognitionV4 service
package visualrecognitionv4

import (
	"fmt"
	"io"
	"strings"

	"github.com/IBM/go-sdk-core/core"
	"github.com/go-openapi/strfmt"
	common "github.com/watson-developer-cloud/go-sdk/common"
)

// VisualRecognitionV4 : Provide images to the IBM Watson&trade; Visual Recognition service for analysis. The service
// detects objects based on a set of images with training data.
//
// **Beta:** The Visual Recognition v4 API and Object Detection model are beta features. For more information about beta
// features, see the [Release
// notes](https://cloud.ibm.com/docs/services/visual-recognition?topic=visual-recognition-release-notes#beta).
// {: important}
//
// Version: 4.0
// See: https://cloud.ibm.com/docs/services/visual-recognition?topic=visual-recognition-object-detection-overview
type VisualRecognitionV4 struct {
	Service *core.BaseService
	Version string
}

const defaultServiceURL = "https://gateway.watsonplatform.net/visual-recognition/api"

// VisualRecognitionV4Options : Service options
type VisualRecognitionV4Options struct {
	URL           string
	Authenticator core.Authenticator
	Version       string
}

// NewVisualRecognitionV4 : Instantiate VisualRecognitionV4
func NewVisualRecognitionV4(options *VisualRecognitionV4Options) (service *VisualRecognitionV4, err error) {
	if options.URL == "" {
		options.URL = defaultServiceURL
	}

	serviceOptions := &core.ServiceOptions{
		URL:           options.URL,
		Authenticator: options.Authenticator,
	}

	if serviceOptions.Authenticator == nil {
		serviceOptions.Authenticator, err = core.GetAuthenticatorFromEnvironment("watson_vision_combined")
		if err != nil {
			return
		}
	}

	baseService, err := core.NewBaseService(serviceOptions, "watson_vision_combined", "Visual Recognition")
	if err != nil {
		return
	}

	service = &VisualRecognitionV4{
		Service: baseService,
		Version: options.Version,
	}

	return
}

// SetServiceURL sets the service URL
func (visualRecognition *VisualRecognitionV4) SetServiceURL(url string) error {
	return visualRecognition.Service.SetServiceURL(url)
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
	err = core.ValidateNotNil(analyzeOptions, "analyzeOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(analyzeOptions, "analyzeOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v4/analyze"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(visualRecognition.Service.Options.URL, pathSegments, pathParameters)
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
	builder.AddQuery("version", visualRecognition.Version)

	if analyzeOptions.CollectionIds != nil {
		builder.AddFormData("collection_ids", "", "", strings.Join(analyzeOptions.CollectionIds, ","))
	}
	if analyzeOptions.Features != nil {
		builder.AddFormData("features", "", "", strings.Join(analyzeOptions.Features, ","))
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

	response, err = visualRecognition.Service.Request(request, new(AnalyzeResponse))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*AnalyzeResponse)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
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
	err = core.ValidateNotNil(createCollectionOptions, "createCollectionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createCollectionOptions, "createCollectionOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v4/collections"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(visualRecognition.Service.Options.URL, pathSegments, pathParameters)
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
	builder.AddQuery("version", visualRecognition.Version)

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

	response, err = visualRecognition.Service.Request(request, new(Collection))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*Collection)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// ListCollections : List collections
// Retrieves a list of collections for the service instance.
func (visualRecognition *VisualRecognitionV4) ListCollections(listCollectionsOptions *ListCollectionsOptions) (result *CollectionsList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listCollectionsOptions, "listCollectionsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v4/collections"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(visualRecognition.Service.Options.URL, pathSegments, pathParameters)
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
	builder.AddQuery("version", visualRecognition.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = visualRecognition.Service.Request(request, new(CollectionsList))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*CollectionsList)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// GetCollection : Get collection details
// Get details of one collection.
func (visualRecognition *VisualRecognitionV4) GetCollection(getCollectionOptions *GetCollectionOptions) (result *Collection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getCollectionOptions, "getCollectionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getCollectionOptions, "getCollectionOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v4/collections"}
	pathParameters := []string{*getCollectionOptions.CollectionID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(visualRecognition.Service.Options.URL, pathSegments, pathParameters)
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
	builder.AddQuery("version", visualRecognition.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = visualRecognition.Service.Request(request, new(Collection))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*Collection)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// UpdateCollection : Update a collection
// Update the name or description of a collection.
//
// Encode the name and description in UTF-8 if they contain non-ASCII characters. The service assumes UTF-8 encoding if
// it encounters non-ASCII characters.
func (visualRecognition *VisualRecognitionV4) UpdateCollection(updateCollectionOptions *UpdateCollectionOptions) (result *Collection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateCollectionOptions, "updateCollectionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateCollectionOptions, "updateCollectionOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v4/collections"}
	pathParameters := []string{*updateCollectionOptions.CollectionID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(visualRecognition.Service.Options.URL, pathSegments, pathParameters)
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
	builder.AddQuery("version", visualRecognition.Version)

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

	response, err = visualRecognition.Service.Request(request, new(Collection))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*Collection)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// DeleteCollection : Delete a collection
// Delete a collection from the service instance.
func (visualRecognition *VisualRecognitionV4) DeleteCollection(deleteCollectionOptions *DeleteCollectionOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteCollectionOptions, "deleteCollectionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteCollectionOptions, "deleteCollectionOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v4/collections"}
	pathParameters := []string{*deleteCollectionOptions.CollectionID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(visualRecognition.Service.Options.URL, pathSegments, pathParameters)
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
	builder.AddQuery("version", visualRecognition.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = visualRecognition.Service.Request(request, nil)

	return
}

// AddImages : Add images
// Add images to a collection by URL, by file, or both.
//
// Encode the image and .zip file names in UTF-8 if they contain non-ASCII characters. The service assumes UTF-8
// encoding if it encounters non-ASCII characters.
func (visualRecognition *VisualRecognitionV4) AddImages(addImagesOptions *AddImagesOptions) (result *ImageDetailsList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addImagesOptions, "addImagesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addImagesOptions, "addImagesOptions")
	if err != nil {
		return
	}
	if (addImagesOptions.ImagesFile == nil) && (addImagesOptions.ImageURL == nil) && (addImagesOptions.TrainingData == nil) {
		err = fmt.Errorf("At least one of imagesFile, imageURL, or trainingData must be supplied")
		return
	}

	pathSegments := []string{"v4/collections", "images"}
	pathParameters := []string{*addImagesOptions.CollectionID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(visualRecognition.Service.Options.URL, pathSegments, pathParameters)
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
	builder.AddQuery("version", visualRecognition.Version)

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

	response, err = visualRecognition.Service.Request(request, new(ImageDetailsList))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*ImageDetailsList)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// ListImages : List images
// Retrieves a list of images in a collection.
func (visualRecognition *VisualRecognitionV4) ListImages(listImagesOptions *ListImagesOptions) (result *ImageSummaryList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listImagesOptions, "listImagesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listImagesOptions, "listImagesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v4/collections", "images"}
	pathParameters := []string{*listImagesOptions.CollectionID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(visualRecognition.Service.Options.URL, pathSegments, pathParameters)
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
	builder.AddQuery("version", visualRecognition.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = visualRecognition.Service.Request(request, new(ImageSummaryList))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*ImageSummaryList)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// GetImageDetails : Get image details
// Get the details of an image in a collection.
func (visualRecognition *VisualRecognitionV4) GetImageDetails(getImageDetailsOptions *GetImageDetailsOptions) (result *ImageDetails, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getImageDetailsOptions, "getImageDetailsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getImageDetailsOptions, "getImageDetailsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v4/collections", "images"}
	pathParameters := []string{*getImageDetailsOptions.CollectionID, *getImageDetailsOptions.ImageID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(visualRecognition.Service.Options.URL, pathSegments, pathParameters)
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
	builder.AddQuery("version", visualRecognition.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = visualRecognition.Service.Request(request, new(ImageDetails))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*ImageDetails)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// DeleteImage : Delete an image
// Delete one image from a collection.
func (visualRecognition *VisualRecognitionV4) DeleteImage(deleteImageOptions *DeleteImageOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteImageOptions, "deleteImageOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteImageOptions, "deleteImageOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v4/collections", "images"}
	pathParameters := []string{*deleteImageOptions.CollectionID, *deleteImageOptions.ImageID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(visualRecognition.Service.Options.URL, pathSegments, pathParameters)
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
	builder.AddQuery("version", visualRecognition.Version)

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
	err = core.ValidateNotNil(getJpegImageOptions, "getJpegImageOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getJpegImageOptions, "getJpegImageOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v4/collections", "images", "jpeg"}
	pathParameters := []string{*getJpegImageOptions.CollectionID, *getJpegImageOptions.ImageID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(visualRecognition.Service.Options.URL, pathSegments, pathParameters)
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

	if getJpegImageOptions.Size != nil {
		builder.AddQuery("size", fmt.Sprint(*getJpegImageOptions.Size))
	}
	builder.AddQuery("version", visualRecognition.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = visualRecognition.Service.Request(request, new(io.ReadCloser))
	if err == nil {
		var ok bool
		result, ok = response.Result.(io.ReadCloser)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// Train : Train a collection
// Start training on images in a collection. The collection must have enough training data and untrained data (the
// **training_status.objects.data_changed** is `true`). If training is in progress, the request queues the next training
// job.
func (visualRecognition *VisualRecognitionV4) Train(trainOptions *TrainOptions) (result *Collection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(trainOptions, "trainOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(trainOptions, "trainOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v4/collections", "train"}
	pathParameters := []string{*trainOptions.CollectionID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(visualRecognition.Service.Options.URL, pathSegments, pathParameters)
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
	builder.AddQuery("version", visualRecognition.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = visualRecognition.Service.Request(request, new(Collection))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*Collection)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
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
	err = core.ValidateNotNil(addImageTrainingDataOptions, "addImageTrainingDataOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addImageTrainingDataOptions, "addImageTrainingDataOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v4/collections", "images", "training_data"}
	pathParameters := []string{*addImageTrainingDataOptions.CollectionID, *addImageTrainingDataOptions.ImageID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(visualRecognition.Service.Options.URL, pathSegments, pathParameters)
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
	builder.AddQuery("version", visualRecognition.Version)

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

	response, err = visualRecognition.Service.Request(request, new(TrainingDataObjects))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*TrainingDataObjects)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// DeleteUserData : Delete labeled data
// Deletes all data associated with a specified customer ID. The method has no effect if no data is associated with the
// customer ID.
//
// You associate a customer ID with data by passing the `X-Watson-Metadata` header with a request that passes data. For
// more information about personal data and customer IDs, see [Information
// security](https://cloud.ibm.com/docs/services/visual-recognition?topic=visual-recognition-information-security).
func (visualRecognition *VisualRecognitionV4) DeleteUserData(deleteUserDataOptions *DeleteUserDataOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteUserDataOptions, "deleteUserDataOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteUserDataOptions, "deleteUserDataOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v4/user_data"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(visualRecognition.Service.Options.URL, pathSegments, pathParameters)
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

	builder.AddQuery("customer_id", fmt.Sprint(*deleteUserDataOptions.CustomerID))
	builder.AddQuery("version", visualRecognition.Version)

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
	CollectionID *string `json:"collection_id" validate:"required"`

	// The identifier of the image.
	ImageID *string `json:"image_id" validate:"required"`

	// Training data for specific objects.
	Objects []TrainingDataObject `json:"objects,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewAddImageTrainingDataOptions : Instantiate AddImageTrainingDataOptions
func (visualRecognition *VisualRecognitionV4) NewAddImageTrainingDataOptions(collectionID string, imageID string) *AddImageTrainingDataOptions {
	return &AddImageTrainingDataOptions{
		CollectionID: core.StringPtr(collectionID),
		ImageID:      core.StringPtr(imageID),
	}
}

// SetCollectionID : Allow user to set CollectionID
func (options *AddImageTrainingDataOptions) SetCollectionID(collectionID string) *AddImageTrainingDataOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
}

// SetImageID : Allow user to set ImageID
func (options *AddImageTrainingDataOptions) SetImageID(imageID string) *AddImageTrainingDataOptions {
	options.ImageID = core.StringPtr(imageID)
	return options
}

// SetObjects : Allow user to set Objects
func (options *AddImageTrainingDataOptions) SetObjects(objects []TrainingDataObject) *AddImageTrainingDataOptions {
	options.Objects = objects
	return options
}

// SetHeaders : Allow user to set Headers
func (options *AddImageTrainingDataOptions) SetHeaders(param map[string]string) *AddImageTrainingDataOptions {
	options.Headers = param
	return options
}

// AddImagesOptions : The AddImages options.
type AddImagesOptions struct {

	// The identifier of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// An array of image files (.jpg or .png) or .zip files with images.
	// - Include a maximum of 20 images in a request.
	// - Limit the .zip file to 100 MB.
	// - Limit each image file to 10 MB.
	//
	// You can also include an image with the **image_url** parameter.
	ImagesFile []FileWithMetadata `json:"images_file,omitempty"`

	// The array of URLs of image files (.jpg or .png).
	// - Include a maximum of 20 images in a request.
	// - Limit each image file to 10 MB.
	// - Minimum width and height is 30 pixels, but the service tends to perform better with images that are at least 300 x
	// 300 pixels. Maximum is 5400 pixels for either height or width.
	//
	// You can also include images with the **images_file** parameter.
	ImageURL []string `json:"image_url,omitempty"`

	// Training data for a single image. Include training data only if you add one image with the request.
	//
	// The `object` property can contain alphanumeric, underscore, hyphen, space, and dot characters. It cannot begin with
	// the reserved prefix `sys-` and must be no longer than 32 characters.
	TrainingData *string `json:"training_data,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewAddImagesOptions : Instantiate AddImagesOptions
func (visualRecognition *VisualRecognitionV4) NewAddImagesOptions(collectionID string) *AddImagesOptions {
	return &AddImagesOptions{
		CollectionID: core.StringPtr(collectionID),
	}
}

// SetCollectionID : Allow user to set CollectionID
func (options *AddImagesOptions) SetCollectionID(collectionID string) *AddImagesOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
}

// SetImagesFile : Allow user to set ImagesFile
func (options *AddImagesOptions) SetImagesFile(imagesFile []FileWithMetadata) *AddImagesOptions {
	options.ImagesFile = imagesFile
	return options
}

// SetImageURL : Allow user to set ImageURL
func (options *AddImagesOptions) SetImageURL(imageURL []string) *AddImagesOptions {
	options.ImageURL = imageURL
	return options
}

// SetTrainingData : Allow user to set TrainingData
func (options *AddImagesOptions) SetTrainingData(trainingData string) *AddImagesOptions {
	options.TrainingData = core.StringPtr(trainingData)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *AddImagesOptions) SetHeaders(param map[string]string) *AddImagesOptions {
	options.Headers = param
	return options
}

// AnalyzeOptions : The Analyze options.
type AnalyzeOptions struct {

	// The IDs of the collections to analyze.
	CollectionIds []string `json:"collection_ids" validate:"required"`

	// The features to analyze.
	Features []string `json:"features" validate:"required"`

	// An array of image files (.jpg or .png) or .zip files with images.
	// - Include a maximum of 20 images in a request.
	// - Limit the .zip file to 100 MB.
	// - Limit each image file to 10 MB.
	//
	// You can also include an image with the **image_url** parameter.
	ImagesFile []FileWithMetadata `json:"images_file,omitempty"`

	// An array of URLs of image files (.jpg or .png).
	// - Include a maximum of 20 images in a request.
	// - Limit each image file to 10 MB.
	// - Minimum width and height is 30 pixels, but the service tends to perform better with images that are at least 300 x
	// 300 pixels. Maximum is 5400 pixels for either height or width.
	//
	// You can also include images with the **images_file** parameter.
	ImageURL []string `json:"image_url,omitempty"`

	// The minimum score a feature must have to be returned.
	Threshold *float32 `json:"threshold,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the AnalyzeOptions.Features property.
const (
	AnalyzeOptions_Features_Objects = "objects"
)

// NewAnalyzeOptions : Instantiate AnalyzeOptions
func (visualRecognition *VisualRecognitionV4) NewAnalyzeOptions(collectionIds []string, features []string) *AnalyzeOptions {
	return &AnalyzeOptions{
		CollectionIds: collectionIds,
		Features:      features,
	}
}

// SetCollectionIds : Allow user to set CollectionIds
func (options *AnalyzeOptions) SetCollectionIds(collectionIds []string) *AnalyzeOptions {
	options.CollectionIds = collectionIds
	return options
}

// SetFeatures : Allow user to set Features
func (options *AnalyzeOptions) SetFeatures(features []string) *AnalyzeOptions {
	options.Features = features
	return options
}

// SetImagesFile : Allow user to set ImagesFile
func (options *AnalyzeOptions) SetImagesFile(imagesFile []FileWithMetadata) *AnalyzeOptions {
	options.ImagesFile = imagesFile
	return options
}

// SetImageURL : Allow user to set ImageURL
func (options *AnalyzeOptions) SetImageURL(imageURL []string) *AnalyzeOptions {
	options.ImageURL = imageURL
	return options
}

// SetThreshold : Allow user to set Threshold
func (options *AnalyzeOptions) SetThreshold(threshold float32) *AnalyzeOptions {
	options.Threshold = core.Float32Ptr(threshold)
	return options
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
	TrainingStatus *TrainingStatus `json:"training_status" validate:"required"`
}

// CollectionObjects : The objects in a collection that are detected in an image.
type CollectionObjects struct {

	// The identifier of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// The identified objects in a collection.
	Objects []ObjectDetail `json:"objects" validate:"required"`
}

// CollectionsList : A container for the list of collections.
type CollectionsList struct {

	// The collections in this service instance.
	Collections []Collection `json:"collections" validate:"required"`
}

// CreateCollectionOptions : The CreateCollection options.
type CreateCollectionOptions struct {

	// The name of the collection. The name can contain alphanumeric, underscore, hyphen, and dot characters. It cannot
	// begin with the reserved prefix `sys-`.
	Name *string `json:"name,omitempty"`

	// The description of the collection.
	Description *string `json:"description,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewCreateCollectionOptions : Instantiate CreateCollectionOptions
func (visualRecognition *VisualRecognitionV4) NewCreateCollectionOptions() *CreateCollectionOptions {
	return &CreateCollectionOptions{}
}

// SetName : Allow user to set Name
func (options *CreateCollectionOptions) SetName(name string) *CreateCollectionOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetDescription : Allow user to set Description
func (options *CreateCollectionOptions) SetDescription(description string) *CreateCollectionOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateCollectionOptions) SetHeaders(param map[string]string) *CreateCollectionOptions {
	options.Headers = param
	return options
}

// DeleteCollectionOptions : The DeleteCollection options.
type DeleteCollectionOptions struct {

	// The identifier of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteCollectionOptions : Instantiate DeleteCollectionOptions
func (visualRecognition *VisualRecognitionV4) NewDeleteCollectionOptions(collectionID string) *DeleteCollectionOptions {
	return &DeleteCollectionOptions{
		CollectionID: core.StringPtr(collectionID),
	}
}

// SetCollectionID : Allow user to set CollectionID
func (options *DeleteCollectionOptions) SetCollectionID(collectionID string) *DeleteCollectionOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteCollectionOptions) SetHeaders(param map[string]string) *DeleteCollectionOptions {
	options.Headers = param
	return options
}

// DeleteImageOptions : The DeleteImage options.
type DeleteImageOptions struct {

	// The identifier of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// The identifier of the image.
	ImageID *string `json:"image_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteImageOptions : Instantiate DeleteImageOptions
func (visualRecognition *VisualRecognitionV4) NewDeleteImageOptions(collectionID string, imageID string) *DeleteImageOptions {
	return &DeleteImageOptions{
		CollectionID: core.StringPtr(collectionID),
		ImageID:      core.StringPtr(imageID),
	}
}

// SetCollectionID : Allow user to set CollectionID
func (options *DeleteImageOptions) SetCollectionID(collectionID string) *DeleteImageOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
}

// SetImageID : Allow user to set ImageID
func (options *DeleteImageOptions) SetImageID(imageID string) *DeleteImageOptions {
	options.ImageID = core.StringPtr(imageID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteImageOptions) SetHeaders(param map[string]string) *DeleteImageOptions {
	options.Headers = param
	return options
}

// DeleteUserDataOptions : The DeleteUserData options.
type DeleteUserDataOptions struct {

	// The customer ID for which all data is to be deleted.
	CustomerID *string `json:"customer_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteUserDataOptions : Instantiate DeleteUserDataOptions
func (visualRecognition *VisualRecognitionV4) NewDeleteUserDataOptions(customerID string) *DeleteUserDataOptions {
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

// DetectedObjects : Container for the list of collections that have objects detected in an image.
type DetectedObjects struct {

	// The collections with identified objects.
	Collections []CollectionObjects `json:"collections,omitempty"`
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
	Error_Code_InvalidField  = "invalid_field"
	Error_Code_InvalidHeader = "invalid_header"
	Error_Code_InvalidMethod = "invalid_method"
	Error_Code_MissingField  = "missing_field"
	Error_Code_ServerError   = "server_error"
)

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
	ErrorTarget_Type_Field     = "field"
	ErrorTarget_Type_Header    = "header"
	ErrorTarget_Type_Parameter = "parameter"
)

// FileWithMetadata : A file with its associated metadata.
type FileWithMetadata struct {

	// The data / content for the file.
	Data io.ReadCloser `json:"data" validate:"required"`

	// The filename of the file.
	Filename *string `json:"filename,omitempty"`

	// The content type of the file.
	ContentType *string `json:"content_type,omitempty"`
}

// GetCollectionOptions : The GetCollection options.
type GetCollectionOptions struct {

	// The identifier of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetCollectionOptions : Instantiate GetCollectionOptions
func (visualRecognition *VisualRecognitionV4) NewGetCollectionOptions(collectionID string) *GetCollectionOptions {
	return &GetCollectionOptions{
		CollectionID: core.StringPtr(collectionID),
	}
}

// SetCollectionID : Allow user to set CollectionID
func (options *GetCollectionOptions) SetCollectionID(collectionID string) *GetCollectionOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetCollectionOptions) SetHeaders(param map[string]string) *GetCollectionOptions {
	options.Headers = param
	return options
}

// GetImageDetailsOptions : The GetImageDetails options.
type GetImageDetailsOptions struct {

	// The identifier of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// The identifier of the image.
	ImageID *string `json:"image_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetImageDetailsOptions : Instantiate GetImageDetailsOptions
func (visualRecognition *VisualRecognitionV4) NewGetImageDetailsOptions(collectionID string, imageID string) *GetImageDetailsOptions {
	return &GetImageDetailsOptions{
		CollectionID: core.StringPtr(collectionID),
		ImageID:      core.StringPtr(imageID),
	}
}

// SetCollectionID : Allow user to set CollectionID
func (options *GetImageDetailsOptions) SetCollectionID(collectionID string) *GetImageDetailsOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
}

// SetImageID : Allow user to set ImageID
func (options *GetImageDetailsOptions) SetImageID(imageID string) *GetImageDetailsOptions {
	options.ImageID = core.StringPtr(imageID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetImageDetailsOptions) SetHeaders(param map[string]string) *GetImageDetailsOptions {
	options.Headers = param
	return options
}

// GetJpegImageOptions : The GetJpegImage options.
type GetJpegImageOptions struct {

	// The identifier of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// The identifier of the image.
	ImageID *string `json:"image_id" validate:"required"`

	// Specify the image size.
	Size *string `json:"size,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the GetJpegImageOptions.Size property.
// Specify the image size.
const (
	GetJpegImageOptions_Size_Full = "full"
)

// NewGetJpegImageOptions : Instantiate GetJpegImageOptions
func (visualRecognition *VisualRecognitionV4) NewGetJpegImageOptions(collectionID string, imageID string) *GetJpegImageOptions {
	return &GetJpegImageOptions{
		CollectionID: core.StringPtr(collectionID),
		ImageID:      core.StringPtr(imageID),
	}
}

// SetCollectionID : Allow user to set CollectionID
func (options *GetJpegImageOptions) SetCollectionID(collectionID string) *GetJpegImageOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
}

// SetImageID : Allow user to set ImageID
func (options *GetJpegImageOptions) SetImageID(imageID string) *GetJpegImageOptions {
	options.ImageID = core.StringPtr(imageID)
	return options
}

// SetSize : Allow user to set Size
func (options *GetJpegImageOptions) SetSize(size string) *GetJpegImageOptions {
	options.Size = core.StringPtr(size)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetJpegImageOptions) SetHeaders(param map[string]string) *GetJpegImageOptions {
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

	// Details about an error.
	Errors *Error `json:"errors,omitempty"`
}

// ImageDetails : Details about an image.
type ImageDetails struct {

	// The identifier of the image.
	ImageID *string `json:"image_id" validate:"required"`

	// Date and time in Coordinated Universal Time (UTC) that the image was most recently updated.
	Updated *strfmt.DateTime `json:"updated" validate:"required"`

	// Date and time in Coordinated Universal Time (UTC) that the image was created.
	Created *strfmt.DateTime `json:"created" validate:"required"`

	// The source type of the image.
	Source *ImageSource `json:"source" validate:"required"`

	// Height and width of an image.
	Dimensions *ImageDimensions `json:"dimensions" validate:"required"`

	// Details about an error.
	Errors *Error `json:"errors,omitempty"`

	// Training data for all objects.
	TrainingData *TrainingDataObjects `json:"training_data" validate:"required"`
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

// ImageDimensions : Height and width of an image.
type ImageDimensions struct {

	// Height in pixels of the image.
	Height *int64 `json:"height" validate:"required"`

	// Width in pixels of the image.
	Width *int64 `json:"width" validate:"required"`
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
	ImageSource_Type_File = "file"
	ImageSource_Type_URL  = "url"
)

// ImageSummary : Basic information about an image.
type ImageSummary struct {

	// The identifier of the image.
	ImageID *string `json:"image_id,omitempty"`

	// Date and time in Coordinated Universal Time (UTC) that the image was most recently updated.
	Updated *strfmt.DateTime `json:"updated,omitempty"`
}

// ImageSummaryList : List of images.
type ImageSummaryList struct {

	// The images in the collection.
	Images []ImageSummary `json:"images" validate:"required"`
}

// ListCollectionsOptions : The ListCollections options.
type ListCollectionsOptions struct {

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListCollectionsOptions : Instantiate ListCollectionsOptions
func (visualRecognition *VisualRecognitionV4) NewListCollectionsOptions() *ListCollectionsOptions {
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
	CollectionID *string `json:"collection_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListImagesOptions : Instantiate ListImagesOptions
func (visualRecognition *VisualRecognitionV4) NewListImagesOptions(collectionID string) *ListImagesOptions {
	return &ListImagesOptions{
		CollectionID: core.StringPtr(collectionID),
	}
}

// SetCollectionID : Allow user to set CollectionID
func (options *ListImagesOptions) SetCollectionID(collectionID string) *ListImagesOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListImagesOptions) SetHeaders(param map[string]string) *ListImagesOptions {
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

// ObjectDetail : Details about an object in the collection.
type ObjectDetail struct {

	// The label for the object.
	Object *string `json:"object" validate:"required"`

	// Defines the location of the bounding box around the object.
	Location *Location `json:"location" validate:"required"`

	// Confidence score for the object in the range of 0 to 1. A higher score indicates greater likelihood that the object
	// is depicted at this location in the image.
	Score *float32 `json:"score" validate:"required"`
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

	// Details about the training. If training is in progress, includes information about the status. If training is not in
	// progress, includes a success message or information about why training failed.
	Description *string `json:"description" validate:"required"`
}

// TrainOptions : The Train options.
type TrainOptions struct {

	// The identifier of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewTrainOptions : Instantiate TrainOptions
func (visualRecognition *VisualRecognitionV4) NewTrainOptions(collectionID string) *TrainOptions {
	return &TrainOptions{
		CollectionID: core.StringPtr(collectionID),
	}
}

// SetCollectionID : Allow user to set CollectionID
func (options *TrainOptions) SetCollectionID(collectionID string) *TrainOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
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

// TrainingDataObjects : Training data for all objects.
type TrainingDataObjects struct {

	// Training data for specific objects.
	Objects []TrainingDataObject `json:"objects,omitempty"`
}

// TrainingStatus : Training status information for the collection.
type TrainingStatus struct {

	// Training status for the objects in the collection.
	Objects *ObjectTrainingStatus `json:"objects" validate:"required"`
}

// UpdateCollectionOptions : The UpdateCollection options.
type UpdateCollectionOptions struct {

	// The identifier of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// The name of the collection. The name can contain alphanumeric, underscore, hyphen, and dot characters. It cannot
	// begin with the reserved prefix `sys-`.
	Name *string `json:"name,omitempty"`

	// The description of the collection.
	Description *string `json:"description,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewUpdateCollectionOptions : Instantiate UpdateCollectionOptions
func (visualRecognition *VisualRecognitionV4) NewUpdateCollectionOptions(collectionID string) *UpdateCollectionOptions {
	return &UpdateCollectionOptions{
		CollectionID: core.StringPtr(collectionID),
	}
}

// SetCollectionID : Allow user to set CollectionID
func (options *UpdateCollectionOptions) SetCollectionID(collectionID string) *UpdateCollectionOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
}

// SetName : Allow user to set Name
func (options *UpdateCollectionOptions) SetName(name string) *UpdateCollectionOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetDescription : Allow user to set Description
func (options *UpdateCollectionOptions) SetDescription(description string) *UpdateCollectionOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateCollectionOptions) SetHeaders(param map[string]string) *UpdateCollectionOptions {
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
	Warning_Code_InvalidField  = "invalid_field"
	Warning_Code_InvalidHeader = "invalid_header"
	Warning_Code_InvalidMethod = "invalid_method"
	Warning_Code_MissingField  = "missing_field"
	Warning_Code_ServerError   = "server_error"
)
