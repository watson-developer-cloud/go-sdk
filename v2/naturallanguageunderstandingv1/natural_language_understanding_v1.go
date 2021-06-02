/**
 * (C) Copyright IBM Corp. 2021.
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
 * IBM OpenAPI SDK Code Generator Version: 3.31.0-902c9336-20210504-161156
 */

// Package naturallanguageunderstandingv1 : Operations and models for the NaturalLanguageUnderstandingV1 service
package naturallanguageunderstandingv1

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

// NaturalLanguageUnderstandingV1 : Analyze various features of text content at scale. Provide text, raw HTML, or a
// public URL and IBM Watson Natural Language Understanding will give you results for the features you request. The
// service cleans HTML content before analysis by default, so the results can ignore most advertisements and other
// unwanted content.
//
// You can create [custom
// models](https://cloud.ibm.com/docs/natural-language-understanding?topic=natural-language-understanding-customizing)
// with Watson Knowledge Studio to detect custom entities and relations in Natural Language Understanding.
//
// Version: 1.0
// See: https://cloud.ibm.com/docs/natural-language-understanding
type NaturalLanguageUnderstandingV1 struct {
	Service *core.BaseService

	// Release date of the API version you want to use. Specify dates in YYYY-MM-DD format. The current version is
	// `2021-03-25`.
	Version *string
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.us-south.natural-language-understanding.watson.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "natural-language-understanding"

// NaturalLanguageUnderstandingV1Options : Service options
type NaturalLanguageUnderstandingV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator

	// Release date of the API version you want to use. Specify dates in YYYY-MM-DD format. The current version is
	// `2021-03-25`.
	Version *string `validate:"required"`
}

// NewNaturalLanguageUnderstandingV1 : constructs an instance of NaturalLanguageUnderstandingV1 with passed in options.
func NewNaturalLanguageUnderstandingV1(options *NaturalLanguageUnderstandingV1Options) (service *NaturalLanguageUnderstandingV1, err error) {
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

	service = &NaturalLanguageUnderstandingV1{
		Service: baseService,
		Version: options.Version,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "naturalLanguageUnderstanding" suitable for processing requests.
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) Clone() *NaturalLanguageUnderstandingV1 {
	if core.IsNil(naturalLanguageUnderstanding) {
		return nil
	}
	clone := *naturalLanguageUnderstanding
	clone.Service = naturalLanguageUnderstanding.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) SetServiceURL(url string) error {
	return naturalLanguageUnderstanding.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) GetServiceURL() string {
	return naturalLanguageUnderstanding.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) SetDefaultHeaders(headers http.Header) {
	naturalLanguageUnderstanding.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) SetEnableGzipCompression(enableGzip bool) {
	naturalLanguageUnderstanding.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) GetEnableGzipCompression() bool {
	return naturalLanguageUnderstanding.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	naturalLanguageUnderstanding.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) DisableRetries() {
	naturalLanguageUnderstanding.Service.DisableRetries()
}

// DisableSSLVerification bypasses verification of the server's SSL certificate
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) DisableSSLVerification() {
	naturalLanguageUnderstanding.Service.DisableSSLVerification()
}

// Analyze : Analyze text
// Analyzes text, HTML, or a public webpage for the following features:
// - Categories
// - Classifications
// - Concepts
// - Emotion
// - Entities
// - Keywords
// - Metadata
// - Relations
// - Semantic roles
// - Sentiment
// - Syntax
// - Summarization (Experimental)
//
// If a language for the input text is not specified with the `language` parameter, the service [automatically detects
// the
// language](https://cloud.ibm.com/docs/natural-language-understanding?topic=natural-language-understanding-detectable-languages).
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) Analyze(analyzeOptions *AnalyzeOptions) (result *AnalysisResults, response *core.DetailedResponse, err error) {
	return naturalLanguageUnderstanding.AnalyzeWithContext(context.Background(), analyzeOptions)
}

// AnalyzeWithContext is an alternate form of the Analyze method which supports a Context parameter
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) AnalyzeWithContext(ctx context.Context, analyzeOptions *AnalyzeOptions) (result *AnalysisResults, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = naturalLanguageUnderstanding.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(naturalLanguageUnderstanding.Service.Options.URL, `/v1/analyze`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range analyzeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("natural-language-understanding", "V1", "Analyze")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*naturalLanguageUnderstanding.Version))

	body := make(map[string]interface{})
	if analyzeOptions.Features != nil {
		body["features"] = analyzeOptions.Features
	}
	if analyzeOptions.Text != nil {
		body["text"] = analyzeOptions.Text
	}
	if analyzeOptions.HTML != nil {
		body["html"] = analyzeOptions.HTML
	}
	if analyzeOptions.URL != nil {
		body["url"] = analyzeOptions.URL
	}
	if analyzeOptions.Clean != nil {
		body["clean"] = analyzeOptions.Clean
	}
	if analyzeOptions.Xpath != nil {
		body["xpath"] = analyzeOptions.Xpath
	}
	if analyzeOptions.FallbackToRaw != nil {
		body["fallback_to_raw"] = analyzeOptions.FallbackToRaw
	}
	if analyzeOptions.ReturnAnalyzedText != nil {
		body["return_analyzed_text"] = analyzeOptions.ReturnAnalyzedText
	}
	if analyzeOptions.Language != nil {
		body["language"] = analyzeOptions.Language
	}
	if analyzeOptions.LimitTextCharacters != nil {
		body["limit_text_characters"] = analyzeOptions.LimitTextCharacters
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
	response, err = naturalLanguageUnderstanding.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAnalysisResults)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListModels : List models
// Lists Watson Knowledge Studio [custom entities and relations
// models](https://cloud.ibm.com/docs/natural-language-understanding?topic=natural-language-understanding-customizing)
// that are deployed to your Natural Language Understanding service.
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) ListModels(listModelsOptions *ListModelsOptions) (result *ListModelsResults, response *core.DetailedResponse, err error) {
	return naturalLanguageUnderstanding.ListModelsWithContext(context.Background(), listModelsOptions)
}

// ListModelsWithContext is an alternate form of the ListModels method which supports a Context parameter
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) ListModelsWithContext(ctx context.Context, listModelsOptions *ListModelsOptions) (result *ListModelsResults, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listModelsOptions, "listModelsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = naturalLanguageUnderstanding.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(naturalLanguageUnderstanding.Service.Options.URL, `/v1/models`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listModelsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("natural-language-understanding", "V1", "ListModels")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*naturalLanguageUnderstanding.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = naturalLanguageUnderstanding.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListModelsResults)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteModel : Delete model
// Deletes a custom model.
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) DeleteModel(deleteModelOptions *DeleteModelOptions) (result *DeleteModelResults, response *core.DetailedResponse, err error) {
	return naturalLanguageUnderstanding.DeleteModelWithContext(context.Background(), deleteModelOptions)
}

// DeleteModelWithContext is an alternate form of the DeleteModel method which supports a Context parameter
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) DeleteModelWithContext(ctx context.Context, deleteModelOptions *DeleteModelOptions) (result *DeleteModelResults, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteModelOptions, "deleteModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteModelOptions, "deleteModelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"model_id": *deleteModelOptions.ModelID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = naturalLanguageUnderstanding.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(naturalLanguageUnderstanding.Service.Options.URL, `/v1/models/{model_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("natural-language-understanding", "V1", "DeleteModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*naturalLanguageUnderstanding.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = naturalLanguageUnderstanding.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeleteModelResults)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateSentimentModel : Create sentiment model
// (Beta) Creates a custom sentiment model by uploading training data and associated metadata. The model begins the
// training and deploying process and is ready to use when the `status` is `available`.
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) CreateSentimentModel(createSentimentModelOptions *CreateSentimentModelOptions) (result *SentimentModel, response *core.DetailedResponse, err error) {
	return naturalLanguageUnderstanding.CreateSentimentModelWithContext(context.Background(), createSentimentModelOptions)
}

// CreateSentimentModelWithContext is an alternate form of the CreateSentimentModel method which supports a Context parameter
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) CreateSentimentModelWithContext(ctx context.Context, createSentimentModelOptions *CreateSentimentModelOptions) (result *SentimentModel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createSentimentModelOptions, "createSentimentModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createSentimentModelOptions, "createSentimentModelOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = naturalLanguageUnderstanding.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(naturalLanguageUnderstanding.Service.Options.URL, `/v1/models/sentiment`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createSentimentModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("natural-language-understanding", "V1", "CreateSentimentModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*naturalLanguageUnderstanding.Version))

	builder.AddFormData("language", "", "", fmt.Sprint(*createSentimentModelOptions.Language))
	builder.AddFormData("training_data", "filename",
		"text/csv", createSentimentModelOptions.TrainingData)
	if createSentimentModelOptions.Name != nil {
		builder.AddFormData("name", "", "", fmt.Sprint(*createSentimentModelOptions.Name))
	}
	if createSentimentModelOptions.Description != nil {
		builder.AddFormData("description", "", "", fmt.Sprint(*createSentimentModelOptions.Description))
	}
	if createSentimentModelOptions.ModelVersion != nil {
		builder.AddFormData("model_version", "", "", fmt.Sprint(*createSentimentModelOptions.ModelVersion))
	}
	if createSentimentModelOptions.WorkspaceID != nil {
		builder.AddFormData("workspace_id", "", "", fmt.Sprint(*createSentimentModelOptions.WorkspaceID))
	}
	if createSentimentModelOptions.VersionDescription != nil {
		builder.AddFormData("version_description", "", "", fmt.Sprint(*createSentimentModelOptions.VersionDescription))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = naturalLanguageUnderstanding.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSentimentModel)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListSentimentModels : List sentiment models
// (Beta) Returns all custom sentiment models associated with this service instance.
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) ListSentimentModels(listSentimentModelsOptions *ListSentimentModelsOptions) (result *ListSentimentModelsResponse, response *core.DetailedResponse, err error) {
	return naturalLanguageUnderstanding.ListSentimentModelsWithContext(context.Background(), listSentimentModelsOptions)
}

// ListSentimentModelsWithContext is an alternate form of the ListSentimentModels method which supports a Context parameter
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) ListSentimentModelsWithContext(ctx context.Context, listSentimentModelsOptions *ListSentimentModelsOptions) (result *ListSentimentModelsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listSentimentModelsOptions, "listSentimentModelsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = naturalLanguageUnderstanding.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(naturalLanguageUnderstanding.Service.Options.URL, `/v1/models/sentiment`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listSentimentModelsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("natural-language-understanding", "V1", "ListSentimentModels")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*naturalLanguageUnderstanding.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = naturalLanguageUnderstanding.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListSentimentModelsResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetSentimentModel : Get sentiment model details
// (Beta) Returns the status of the sentiment model with the given model ID.
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) GetSentimentModel(getSentimentModelOptions *GetSentimentModelOptions) (result *SentimentModel, response *core.DetailedResponse, err error) {
	return naturalLanguageUnderstanding.GetSentimentModelWithContext(context.Background(), getSentimentModelOptions)
}

// GetSentimentModelWithContext is an alternate form of the GetSentimentModel method which supports a Context parameter
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) GetSentimentModelWithContext(ctx context.Context, getSentimentModelOptions *GetSentimentModelOptions) (result *SentimentModel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getSentimentModelOptions, "getSentimentModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getSentimentModelOptions, "getSentimentModelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"model_id": *getSentimentModelOptions.ModelID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = naturalLanguageUnderstanding.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(naturalLanguageUnderstanding.Service.Options.URL, `/v1/models/sentiment/{model_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getSentimentModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("natural-language-understanding", "V1", "GetSentimentModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*naturalLanguageUnderstanding.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = naturalLanguageUnderstanding.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSentimentModel)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateSentimentModel : Update sentiment model
// (Beta) Overwrites the training data associated with this custom sentiment model and retrains the model. The new model
// replaces the current deployment.
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) UpdateSentimentModel(updateSentimentModelOptions *UpdateSentimentModelOptions) (result *SentimentModel, response *core.DetailedResponse, err error) {
	return naturalLanguageUnderstanding.UpdateSentimentModelWithContext(context.Background(), updateSentimentModelOptions)
}

// UpdateSentimentModelWithContext is an alternate form of the UpdateSentimentModel method which supports a Context parameter
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) UpdateSentimentModelWithContext(ctx context.Context, updateSentimentModelOptions *UpdateSentimentModelOptions) (result *SentimentModel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateSentimentModelOptions, "updateSentimentModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateSentimentModelOptions, "updateSentimentModelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"model_id": *updateSentimentModelOptions.ModelID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = naturalLanguageUnderstanding.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(naturalLanguageUnderstanding.Service.Options.URL, `/v1/models/sentiment/{model_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateSentimentModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("natural-language-understanding", "V1", "UpdateSentimentModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*naturalLanguageUnderstanding.Version))

	builder.AddFormData("language", "", "", fmt.Sprint(*updateSentimentModelOptions.Language))
	builder.AddFormData("training_data", "filename",
		"text/csv", updateSentimentModelOptions.TrainingData)
	if updateSentimentModelOptions.Name != nil {
		builder.AddFormData("name", "", "", fmt.Sprint(*updateSentimentModelOptions.Name))
	}
	if updateSentimentModelOptions.Description != nil {
		builder.AddFormData("description", "", "", fmt.Sprint(*updateSentimentModelOptions.Description))
	}
	if updateSentimentModelOptions.ModelVersion != nil {
		builder.AddFormData("model_version", "", "", fmt.Sprint(*updateSentimentModelOptions.ModelVersion))
	}
	if updateSentimentModelOptions.WorkspaceID != nil {
		builder.AddFormData("workspace_id", "", "", fmt.Sprint(*updateSentimentModelOptions.WorkspaceID))
	}
	if updateSentimentModelOptions.VersionDescription != nil {
		builder.AddFormData("version_description", "", "", fmt.Sprint(*updateSentimentModelOptions.VersionDescription))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = naturalLanguageUnderstanding.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSentimentModel)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteSentimentModel : Delete sentiment model
// (Beta) Un-deploys the custom sentiment model with the given model ID and deletes all associated customer data,
// including any training data or binary artifacts.
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) DeleteSentimentModel(deleteSentimentModelOptions *DeleteSentimentModelOptions) (result *DeleteModelResults, response *core.DetailedResponse, err error) {
	return naturalLanguageUnderstanding.DeleteSentimentModelWithContext(context.Background(), deleteSentimentModelOptions)
}

// DeleteSentimentModelWithContext is an alternate form of the DeleteSentimentModel method which supports a Context parameter
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) DeleteSentimentModelWithContext(ctx context.Context, deleteSentimentModelOptions *DeleteSentimentModelOptions) (result *DeleteModelResults, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteSentimentModelOptions, "deleteSentimentModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteSentimentModelOptions, "deleteSentimentModelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"model_id": *deleteSentimentModelOptions.ModelID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = naturalLanguageUnderstanding.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(naturalLanguageUnderstanding.Service.Options.URL, `/v1/models/sentiment/{model_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteSentimentModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("natural-language-understanding", "V1", "DeleteSentimentModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*naturalLanguageUnderstanding.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = naturalLanguageUnderstanding.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeleteModelResults)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateCategoriesModel : Create categories model
// (Beta) Creates a custom categories model by uploading training data and associated metadata. The model begins the
// training and deploying process and is ready to use when the `status` is `available`.
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) CreateCategoriesModel(createCategoriesModelOptions *CreateCategoriesModelOptions) (result *CategoriesModel, response *core.DetailedResponse, err error) {
	return naturalLanguageUnderstanding.CreateCategoriesModelWithContext(context.Background(), createCategoriesModelOptions)
}

// CreateCategoriesModelWithContext is an alternate form of the CreateCategoriesModel method which supports a Context parameter
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) CreateCategoriesModelWithContext(ctx context.Context, createCategoriesModelOptions *CreateCategoriesModelOptions) (result *CategoriesModel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createCategoriesModelOptions, "createCategoriesModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createCategoriesModelOptions, "createCategoriesModelOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = naturalLanguageUnderstanding.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(naturalLanguageUnderstanding.Service.Options.URL, `/v1/models/categories`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createCategoriesModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("natural-language-understanding", "V1", "CreateCategoriesModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*naturalLanguageUnderstanding.Version))

	builder.AddFormData("language", "", "", fmt.Sprint(*createCategoriesModelOptions.Language))
	builder.AddFormData("training_data", "filename",
		core.StringNilMapper(createCategoriesModelOptions.TrainingDataContentType), createCategoriesModelOptions.TrainingData)
	if createCategoriesModelOptions.Name != nil {
		builder.AddFormData("name", "", "", fmt.Sprint(*createCategoriesModelOptions.Name))
	}
	if createCategoriesModelOptions.Description != nil {
		builder.AddFormData("description", "", "", fmt.Sprint(*createCategoriesModelOptions.Description))
	}
	if createCategoriesModelOptions.ModelVersion != nil {
		builder.AddFormData("model_version", "", "", fmt.Sprint(*createCategoriesModelOptions.ModelVersion))
	}
	if createCategoriesModelOptions.WorkspaceID != nil {
		builder.AddFormData("workspace_id", "", "", fmt.Sprint(*createCategoriesModelOptions.WorkspaceID))
	}
	if createCategoriesModelOptions.VersionDescription != nil {
		builder.AddFormData("version_description", "", "", fmt.Sprint(*createCategoriesModelOptions.VersionDescription))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = naturalLanguageUnderstanding.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCategoriesModel)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListCategoriesModels : List categories models
// (Beta) Returns all custom categories models associated with this service instance.
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) ListCategoriesModels(listCategoriesModelsOptions *ListCategoriesModelsOptions) (result *CategoriesModelList, response *core.DetailedResponse, err error) {
	return naturalLanguageUnderstanding.ListCategoriesModelsWithContext(context.Background(), listCategoriesModelsOptions)
}

// ListCategoriesModelsWithContext is an alternate form of the ListCategoriesModels method which supports a Context parameter
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) ListCategoriesModelsWithContext(ctx context.Context, listCategoriesModelsOptions *ListCategoriesModelsOptions) (result *CategoriesModelList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listCategoriesModelsOptions, "listCategoriesModelsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = naturalLanguageUnderstanding.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(naturalLanguageUnderstanding.Service.Options.URL, `/v1/models/categories`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listCategoriesModelsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("natural-language-understanding", "V1", "ListCategoriesModels")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*naturalLanguageUnderstanding.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = naturalLanguageUnderstanding.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCategoriesModelList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetCategoriesModel : Get categories model details
// (Beta) Returns the status of the categories model with the given model ID.
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) GetCategoriesModel(getCategoriesModelOptions *GetCategoriesModelOptions) (result *CategoriesModel, response *core.DetailedResponse, err error) {
	return naturalLanguageUnderstanding.GetCategoriesModelWithContext(context.Background(), getCategoriesModelOptions)
}

// GetCategoriesModelWithContext is an alternate form of the GetCategoriesModel method which supports a Context parameter
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) GetCategoriesModelWithContext(ctx context.Context, getCategoriesModelOptions *GetCategoriesModelOptions) (result *CategoriesModel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getCategoriesModelOptions, "getCategoriesModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getCategoriesModelOptions, "getCategoriesModelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"model_id": *getCategoriesModelOptions.ModelID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = naturalLanguageUnderstanding.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(naturalLanguageUnderstanding.Service.Options.URL, `/v1/models/categories/{model_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCategoriesModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("natural-language-understanding", "V1", "GetCategoriesModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*naturalLanguageUnderstanding.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = naturalLanguageUnderstanding.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCategoriesModel)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateCategoriesModel : Update categories model
// (Beta) Overwrites the training data associated with this custom categories model and retrains the model. The new
// model replaces the current deployment.
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) UpdateCategoriesModel(updateCategoriesModelOptions *UpdateCategoriesModelOptions) (result *CategoriesModel, response *core.DetailedResponse, err error) {
	return naturalLanguageUnderstanding.UpdateCategoriesModelWithContext(context.Background(), updateCategoriesModelOptions)
}

// UpdateCategoriesModelWithContext is an alternate form of the UpdateCategoriesModel method which supports a Context parameter
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) UpdateCategoriesModelWithContext(ctx context.Context, updateCategoriesModelOptions *UpdateCategoriesModelOptions) (result *CategoriesModel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateCategoriesModelOptions, "updateCategoriesModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateCategoriesModelOptions, "updateCategoriesModelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"model_id": *updateCategoriesModelOptions.ModelID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = naturalLanguageUnderstanding.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(naturalLanguageUnderstanding.Service.Options.URL, `/v1/models/categories/{model_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateCategoriesModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("natural-language-understanding", "V1", "UpdateCategoriesModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*naturalLanguageUnderstanding.Version))

	builder.AddFormData("language", "", "", fmt.Sprint(*updateCategoriesModelOptions.Language))
	builder.AddFormData("training_data", "filename",
		core.StringNilMapper(updateCategoriesModelOptions.TrainingDataContentType), updateCategoriesModelOptions.TrainingData)
	if updateCategoriesModelOptions.Name != nil {
		builder.AddFormData("name", "", "", fmt.Sprint(*updateCategoriesModelOptions.Name))
	}
	if updateCategoriesModelOptions.Description != nil {
		builder.AddFormData("description", "", "", fmt.Sprint(*updateCategoriesModelOptions.Description))
	}
	if updateCategoriesModelOptions.ModelVersion != nil {
		builder.AddFormData("model_version", "", "", fmt.Sprint(*updateCategoriesModelOptions.ModelVersion))
	}
	if updateCategoriesModelOptions.WorkspaceID != nil {
		builder.AddFormData("workspace_id", "", "", fmt.Sprint(*updateCategoriesModelOptions.WorkspaceID))
	}
	if updateCategoriesModelOptions.VersionDescription != nil {
		builder.AddFormData("version_description", "", "", fmt.Sprint(*updateCategoriesModelOptions.VersionDescription))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = naturalLanguageUnderstanding.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCategoriesModel)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteCategoriesModel : Delete categories model
// (Beta) Un-deploys the custom categories model with the given model ID and deletes all associated customer data,
// including any training data or binary artifacts.
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) DeleteCategoriesModel(deleteCategoriesModelOptions *DeleteCategoriesModelOptions) (result *DeleteModelResults, response *core.DetailedResponse, err error) {
	return naturalLanguageUnderstanding.DeleteCategoriesModelWithContext(context.Background(), deleteCategoriesModelOptions)
}

// DeleteCategoriesModelWithContext is an alternate form of the DeleteCategoriesModel method which supports a Context parameter
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) DeleteCategoriesModelWithContext(ctx context.Context, deleteCategoriesModelOptions *DeleteCategoriesModelOptions) (result *DeleteModelResults, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteCategoriesModelOptions, "deleteCategoriesModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteCategoriesModelOptions, "deleteCategoriesModelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"model_id": *deleteCategoriesModelOptions.ModelID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = naturalLanguageUnderstanding.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(naturalLanguageUnderstanding.Service.Options.URL, `/v1/models/categories/{model_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteCategoriesModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("natural-language-understanding", "V1", "DeleteCategoriesModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*naturalLanguageUnderstanding.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = naturalLanguageUnderstanding.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeleteModelResults)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateClassificationsModel : Create classifications model
// (Beta) Creates a custom classifications model by uploading training data and associated metadata. The model begins
// the training and deploying process and is ready to use when the `status` is `available`.
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) CreateClassificationsModel(createClassificationsModelOptions *CreateClassificationsModelOptions) (result *ClassificationsModel, response *core.DetailedResponse, err error) {
	return naturalLanguageUnderstanding.CreateClassificationsModelWithContext(context.Background(), createClassificationsModelOptions)
}

// CreateClassificationsModelWithContext is an alternate form of the CreateClassificationsModel method which supports a Context parameter
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) CreateClassificationsModelWithContext(ctx context.Context, createClassificationsModelOptions *CreateClassificationsModelOptions) (result *ClassificationsModel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createClassificationsModelOptions, "createClassificationsModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createClassificationsModelOptions, "createClassificationsModelOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = naturalLanguageUnderstanding.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(naturalLanguageUnderstanding.Service.Options.URL, `/v1/models/classifications`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createClassificationsModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("natural-language-understanding", "V1", "CreateClassificationsModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*naturalLanguageUnderstanding.Version))

	builder.AddFormData("language", "", "", fmt.Sprint(*createClassificationsModelOptions.Language))
	builder.AddFormData("training_data", "filename",
		core.StringNilMapper(createClassificationsModelOptions.TrainingDataContentType), createClassificationsModelOptions.TrainingData)
	if createClassificationsModelOptions.Name != nil {
		builder.AddFormData("name", "", "", fmt.Sprint(*createClassificationsModelOptions.Name))
	}
	if createClassificationsModelOptions.Description != nil {
		builder.AddFormData("description", "", "", fmt.Sprint(*createClassificationsModelOptions.Description))
	}
	if createClassificationsModelOptions.ModelVersion != nil {
		builder.AddFormData("model_version", "", "", fmt.Sprint(*createClassificationsModelOptions.ModelVersion))
	}
	if createClassificationsModelOptions.WorkspaceID != nil {
		builder.AddFormData("workspace_id", "", "", fmt.Sprint(*createClassificationsModelOptions.WorkspaceID))
	}
	if createClassificationsModelOptions.VersionDescription != nil {
		builder.AddFormData("version_description", "", "", fmt.Sprint(*createClassificationsModelOptions.VersionDescription))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = naturalLanguageUnderstanding.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalClassificationsModel)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListClassificationsModels : List classifications models
// (Beta) Returns all custom classifications models associated with this service instance.
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) ListClassificationsModels(listClassificationsModelsOptions *ListClassificationsModelsOptions) (result *ListClassificationsModelsResponse, response *core.DetailedResponse, err error) {
	return naturalLanguageUnderstanding.ListClassificationsModelsWithContext(context.Background(), listClassificationsModelsOptions)
}

// ListClassificationsModelsWithContext is an alternate form of the ListClassificationsModels method which supports a Context parameter
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) ListClassificationsModelsWithContext(ctx context.Context, listClassificationsModelsOptions *ListClassificationsModelsOptions) (result *ListClassificationsModelsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listClassificationsModelsOptions, "listClassificationsModelsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = naturalLanguageUnderstanding.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(naturalLanguageUnderstanding.Service.Options.URL, `/v1/models/classifications`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listClassificationsModelsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("natural-language-understanding", "V1", "ListClassificationsModels")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*naturalLanguageUnderstanding.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = naturalLanguageUnderstanding.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListClassificationsModelsResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetClassificationsModel : Get classifications model details
// (Beta) Returns the status of the classifications model with the given model ID.
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) GetClassificationsModel(getClassificationsModelOptions *GetClassificationsModelOptions) (result *ClassificationsModel, response *core.DetailedResponse, err error) {
	return naturalLanguageUnderstanding.GetClassificationsModelWithContext(context.Background(), getClassificationsModelOptions)
}

// GetClassificationsModelWithContext is an alternate form of the GetClassificationsModel method which supports a Context parameter
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) GetClassificationsModelWithContext(ctx context.Context, getClassificationsModelOptions *GetClassificationsModelOptions) (result *ClassificationsModel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getClassificationsModelOptions, "getClassificationsModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getClassificationsModelOptions, "getClassificationsModelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"model_id": *getClassificationsModelOptions.ModelID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = naturalLanguageUnderstanding.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(naturalLanguageUnderstanding.Service.Options.URL, `/v1/models/classifications/{model_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getClassificationsModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("natural-language-understanding", "V1", "GetClassificationsModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*naturalLanguageUnderstanding.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = naturalLanguageUnderstanding.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalClassificationsModel)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateClassificationsModel : Update classifications model
// (Beta) Overwrites the training data associated with this custom classifications model and retrains the model. The new
// model replaces the current deployment.
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) UpdateClassificationsModel(updateClassificationsModelOptions *UpdateClassificationsModelOptions) (result *ClassificationsModel, response *core.DetailedResponse, err error) {
	return naturalLanguageUnderstanding.UpdateClassificationsModelWithContext(context.Background(), updateClassificationsModelOptions)
}

// UpdateClassificationsModelWithContext is an alternate form of the UpdateClassificationsModel method which supports a Context parameter
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) UpdateClassificationsModelWithContext(ctx context.Context, updateClassificationsModelOptions *UpdateClassificationsModelOptions) (result *ClassificationsModel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateClassificationsModelOptions, "updateClassificationsModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateClassificationsModelOptions, "updateClassificationsModelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"model_id": *updateClassificationsModelOptions.ModelID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = naturalLanguageUnderstanding.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(naturalLanguageUnderstanding.Service.Options.URL, `/v1/models/classifications/{model_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateClassificationsModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("natural-language-understanding", "V1", "UpdateClassificationsModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*naturalLanguageUnderstanding.Version))

	builder.AddFormData("language", "", "", fmt.Sprint(*updateClassificationsModelOptions.Language))
	builder.AddFormData("training_data", "filename",
		core.StringNilMapper(updateClassificationsModelOptions.TrainingDataContentType), updateClassificationsModelOptions.TrainingData)
	if updateClassificationsModelOptions.Name != nil {
		builder.AddFormData("name", "", "", fmt.Sprint(*updateClassificationsModelOptions.Name))
	}
	if updateClassificationsModelOptions.Description != nil {
		builder.AddFormData("description", "", "", fmt.Sprint(*updateClassificationsModelOptions.Description))
	}
	if updateClassificationsModelOptions.ModelVersion != nil {
		builder.AddFormData("model_version", "", "", fmt.Sprint(*updateClassificationsModelOptions.ModelVersion))
	}
	if updateClassificationsModelOptions.WorkspaceID != nil {
		builder.AddFormData("workspace_id", "", "", fmt.Sprint(*updateClassificationsModelOptions.WorkspaceID))
	}
	if updateClassificationsModelOptions.VersionDescription != nil {
		builder.AddFormData("version_description", "", "", fmt.Sprint(*updateClassificationsModelOptions.VersionDescription))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = naturalLanguageUnderstanding.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalClassificationsModel)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteClassificationsModel : Delete classifications model
// (Beta) Un-deploys the custom classifications model with the given model ID and deletes all associated customer data,
// including any training data or binary artifacts.
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) DeleteClassificationsModel(deleteClassificationsModelOptions *DeleteClassificationsModelOptions) (result *DeleteModelResults, response *core.DetailedResponse, err error) {
	return naturalLanguageUnderstanding.DeleteClassificationsModelWithContext(context.Background(), deleteClassificationsModelOptions)
}

// DeleteClassificationsModelWithContext is an alternate form of the DeleteClassificationsModel method which supports a Context parameter
func (naturalLanguageUnderstanding *NaturalLanguageUnderstandingV1) DeleteClassificationsModelWithContext(ctx context.Context, deleteClassificationsModelOptions *DeleteClassificationsModelOptions) (result *DeleteModelResults, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteClassificationsModelOptions, "deleteClassificationsModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteClassificationsModelOptions, "deleteClassificationsModelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"model_id": *deleteClassificationsModelOptions.ModelID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = naturalLanguageUnderstanding.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(naturalLanguageUnderstanding.Service.Options.URL, `/v1/models/classifications/{model_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteClassificationsModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("natural-language-understanding", "V1", "DeleteClassificationsModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*naturalLanguageUnderstanding.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = naturalLanguageUnderstanding.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeleteModelResults)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// AnalysisResults : Results of the analysis, organized by feature.
type AnalysisResults struct {
	// Language used to analyze the text.
	Language *string `json:"language,omitempty"`

	// Text that was used in the analysis.
	AnalyzedText *string `json:"analyzed_text,omitempty"`

	// URL of the webpage that was analyzed.
	RetrievedURL *string `json:"retrieved_url,omitempty"`

	// API usage information for the request.
	Usage *AnalysisResultsUsage `json:"usage,omitempty"`

	// The general concepts referenced or alluded to in the analyzed text.
	Concepts []ConceptsResult `json:"concepts,omitempty"`

	// The entities detected in the analyzed text.
	Entities []EntitiesResult `json:"entities,omitempty"`

	// The keywords from the analyzed text.
	Keywords []KeywordsResult `json:"keywords,omitempty"`

	// The categories that the service assigned to the analyzed text.
	Categories []CategoriesResult `json:"categories,omitempty"`

	// The classifications assigned to the analyzed text.
	Classifications []ClassificationsResult `json:"classifications,omitempty"`

	// The anger, disgust, fear, joy, or sadness conveyed by the content.
	Emotion *EmotionResult `json:"emotion,omitempty"`

	// Webpage metadata, such as the author and the title of the page.
	Metadata *FeaturesResultsMetadata `json:"metadata,omitempty"`

	// The relationships between entities in the content.
	Relations []RelationsResult `json:"relations,omitempty"`

	// Sentences parsed into `subject`, `action`, and `object` form.
	SemanticRoles []SemanticRolesResult `json:"semantic_roles,omitempty"`

	// The sentiment of the content.
	Sentiment *SentimentResult `json:"sentiment,omitempty"`

	// Tokens and sentences returned from syntax analysis.
	Syntax *SyntaxResult `json:"syntax,omitempty"`
}

// UnmarshalAnalysisResults unmarshals an instance of AnalysisResults from the specified map of raw messages.
func UnmarshalAnalysisResults(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AnalysisResults)
	err = core.UnmarshalPrimitive(m, "language", &obj.Language)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "analyzed_text", &obj.AnalyzedText)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "retrieved_url", &obj.RetrievedURL)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "usage", &obj.Usage, UnmarshalAnalysisResultsUsage)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "concepts", &obj.Concepts, UnmarshalConceptsResult)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "entities", &obj.Entities, UnmarshalEntitiesResult)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "keywords", &obj.Keywords, UnmarshalKeywordsResult)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "categories", &obj.Categories, UnmarshalCategoriesResult)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "classifications", &obj.Classifications, UnmarshalClassificationsResult)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "emotion", &obj.Emotion, UnmarshalEmotionResult)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalFeaturesResultsMetadata)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "relations", &obj.Relations, UnmarshalRelationsResult)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "semantic_roles", &obj.SemanticRoles, UnmarshalSemanticRolesResult)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "sentiment", &obj.Sentiment, UnmarshalSentimentResult)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "syntax", &obj.Syntax, UnmarshalSyntaxResult)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AnalysisResultsUsage : API usage information for the request.
type AnalysisResultsUsage struct {
	// Number of features used in the API call.
	Features *int64 `json:"features,omitempty"`

	// Number of text characters processed.
	TextCharacters *int64 `json:"text_characters,omitempty"`

	// Number of 10,000-character units processed.
	TextUnits *int64 `json:"text_units,omitempty"`
}

// UnmarshalAnalysisResultsUsage unmarshals an instance of AnalysisResultsUsage from the specified map of raw messages.
func UnmarshalAnalysisResultsUsage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AnalysisResultsUsage)
	err = core.UnmarshalPrimitive(m, "features", &obj.Features)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text_characters", &obj.TextCharacters)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text_units", &obj.TextUnits)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AnalyzeOptions : The Analyze options.
type AnalyzeOptions struct {
	// Specific features to analyze the document for.
	Features *Features `validate:"required"`

	// The plain text to analyze. One of the `text`, `html`, or `url` parameters is required.
	Text *string

	// The HTML file to analyze. One of the `text`, `html`, or `url` parameters is required.
	HTML *string

	// The webpage to analyze. One of the `text`, `html`, or `url` parameters is required.
	URL *string

	// Set this to `false` to disable webpage cleaning. For more information about webpage cleaning, see [Analyzing
	// webpages](https://cloud.ibm.com/docs/natural-language-understanding?topic=natural-language-understanding-analyzing-webpages).
	Clean *bool

	// An [XPath
	// query](https://cloud.ibm.com/docs/natural-language-understanding?topic=natural-language-understanding-analyzing-webpages#xpath)
	// to perform on `html` or `url` input. Results of the query will be appended to the cleaned webpage text before it is
	// analyzed. To analyze only the results of the XPath query, set the `clean` parameter to `false`.
	Xpath *string

	// Whether to use raw HTML content if text cleaning fails.
	FallbackToRaw *bool

	// Whether or not to return the analyzed text.
	ReturnAnalyzedText *bool

	// ISO 639-1 code that specifies the language of your text. This overrides automatic language detection. Language
	// support differs depending on the features you include in your analysis. For more information, see [Language
	// support](https://cloud.ibm.com/docs/natural-language-understanding?topic=natural-language-understanding-language-support).
	Language *string

	// Sets the maximum number of characters that are processed by the service.
	LimitTextCharacters *int64

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAnalyzeOptions : Instantiate AnalyzeOptions
func (*NaturalLanguageUnderstandingV1) NewAnalyzeOptions(features *Features) *AnalyzeOptions {
	return &AnalyzeOptions{
		Features: features,
	}
}

// SetFeatures : Allow user to set Features
func (options *AnalyzeOptions) SetFeatures(features *Features) *AnalyzeOptions {
	options.Features = features
	return options
}

// SetText : Allow user to set Text
func (options *AnalyzeOptions) SetText(text string) *AnalyzeOptions {
	options.Text = core.StringPtr(text)
	return options
}

// SetHTML : Allow user to set HTML
func (options *AnalyzeOptions) SetHTML(html string) *AnalyzeOptions {
	options.HTML = core.StringPtr(html)
	return options
}

// SetURL : Allow user to set URL
func (options *AnalyzeOptions) SetURL(url string) *AnalyzeOptions {
	options.URL = core.StringPtr(url)
	return options
}

// SetClean : Allow user to set Clean
func (options *AnalyzeOptions) SetClean(clean bool) *AnalyzeOptions {
	options.Clean = core.BoolPtr(clean)
	return options
}

// SetXpath : Allow user to set Xpath
func (options *AnalyzeOptions) SetXpath(xpath string) *AnalyzeOptions {
	options.Xpath = core.StringPtr(xpath)
	return options
}

// SetFallbackToRaw : Allow user to set FallbackToRaw
func (options *AnalyzeOptions) SetFallbackToRaw(fallbackToRaw bool) *AnalyzeOptions {
	options.FallbackToRaw = core.BoolPtr(fallbackToRaw)
	return options
}

// SetReturnAnalyzedText : Allow user to set ReturnAnalyzedText
func (options *AnalyzeOptions) SetReturnAnalyzedText(returnAnalyzedText bool) *AnalyzeOptions {
	options.ReturnAnalyzedText = core.BoolPtr(returnAnalyzedText)
	return options
}

// SetLanguage : Allow user to set Language
func (options *AnalyzeOptions) SetLanguage(language string) *AnalyzeOptions {
	options.Language = core.StringPtr(language)
	return options
}

// SetLimitTextCharacters : Allow user to set LimitTextCharacters
func (options *AnalyzeOptions) SetLimitTextCharacters(limitTextCharacters int64) *AnalyzeOptions {
	options.LimitTextCharacters = core.Int64Ptr(limitTextCharacters)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *AnalyzeOptions) SetHeaders(param map[string]string) *AnalyzeOptions {
	options.Headers = param
	return options
}

// Author : The author of the analyzed content.
type Author struct {
	// Name of the author.
	Name *string `json:"name,omitempty"`
}

// UnmarshalAuthor unmarshals an instance of Author from the specified map of raw messages.
func UnmarshalAuthor(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Author)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CategoriesModel : Categories model.
type CategoriesModel struct {
	// An optional name for the model.
	Name *string `json:"name,omitempty"`

	// An optional map of metadata key-value pairs to store with this model.
	UserMetadata map[string]interface{} `json:"user_metadata,omitempty"`

	// The 2-letter language code of this model.
	Language *string `json:"language" validate:"required"`

	// An optional description of the model.
	Description *string `json:"description,omitempty"`

	// An optional version string.
	ModelVersion *string `json:"model_version,omitempty"`

	// ID of the Watson Knowledge Studio workspace that deployed this model to Natural Language Understanding.
	WorkspaceID *string `json:"workspace_id,omitempty"`

	// The description of the version.
	VersionDescription *string `json:"version_description,omitempty"`

	// The service features that are supported by the custom model.
	Features []string `json:"features,omitempty"`

	// When the status is `available`, the model is ready to use.
	Status *string `json:"status" validate:"required"`

	// Unique model ID.
	ModelID *string `json:"model_id" validate:"required"`

	// dateTime indicating when the model was created.
	Created *strfmt.DateTime `json:"created" validate:"required"`

	Notices []Notice `json:"notices,omitempty"`

	// dateTime of last successful model training.
	LastTrained *strfmt.DateTime `json:"last_trained,omitempty"`

	// dateTime of last successful model deployment.
	LastDeployed *strfmt.DateTime `json:"last_deployed,omitempty"`
}

// Constants associated with the CategoriesModel.Status property.
// When the status is `available`, the model is ready to use.
const (
	CategoriesModelStatusAvailableConst = "available"
	CategoriesModelStatusDeletedConst   = "deleted"
	CategoriesModelStatusDeployingConst = "deploying"
	CategoriesModelStatusErrorConst     = "error"
	CategoriesModelStatusStartingConst  = "starting"
	CategoriesModelStatusTrainingConst  = "training"
)

// UnmarshalCategoriesModel unmarshals an instance of CategoriesModel from the specified map of raw messages.
func UnmarshalCategoriesModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CategoriesModel)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "user_metadata", &obj.UserMetadata)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "language", &obj.Language)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "model_version", &obj.ModelVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "workspace_id", &obj.WorkspaceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version_description", &obj.VersionDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "features", &obj.Features)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "model_id", &obj.ModelID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created", &obj.Created)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "notices", &obj.Notices, UnmarshalNotice)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_trained", &obj.LastTrained)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_deployed", &obj.LastDeployed)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CategoriesModelList : List of categories models.
type CategoriesModelList struct {
	// The categories models.
	Models []CategoriesModel `json:"models,omitempty"`
}

// UnmarshalCategoriesModelList unmarshals an instance of CategoriesModelList from the specified map of raw messages.
func UnmarshalCategoriesModelList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CategoriesModelList)
	err = core.UnmarshalModel(m, "models", &obj.Models, UnmarshalCategoriesModel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CategoriesOptions : Returns a five-level taxonomy of the content. The top three categories are returned.
//
// Supported languages: Arabic, English, French, German, Italian, Japanese, Korean, Portuguese, Spanish.
type CategoriesOptions struct {
	// Set this to `true` to return explanations for each categorization. **This is available only for English
	// categories.**.
	Explanation *bool `json:"explanation,omitempty"`

	// Maximum number of categories to return.
	Limit *int64 `json:"limit,omitempty"`

	// (Beta) Enter a [custom
	// model](https://cloud.ibm.com/docs/natural-language-understanding?topic=natural-language-understanding-customizing)
	// ID to override the standard categories model. **This is available only for English categories.**.
	Model *string `json:"model,omitempty"`
}

// UnmarshalCategoriesOptions unmarshals an instance of CategoriesOptions from the specified map of raw messages.
func UnmarshalCategoriesOptions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CategoriesOptions)
	err = core.UnmarshalPrimitive(m, "explanation", &obj.Explanation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "model", &obj.Model)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CategoriesRelevantText : Relevant text that contributed to the categorization.
type CategoriesRelevantText struct {
	// Text from the analyzed source that supports the categorization.
	Text *string `json:"text,omitempty"`
}

// UnmarshalCategoriesRelevantText unmarshals an instance of CategoriesRelevantText from the specified map of raw messages.
func UnmarshalCategoriesRelevantText(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CategoriesRelevantText)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CategoriesResult : A categorization of the analyzed text.
type CategoriesResult struct {
	// The path to the category through the 5-level taxonomy hierarchy. For more information about the categories, see
	// [Categories
	// hierarchy](https://cloud.ibm.com/docs/natural-language-understanding?topic=natural-language-understanding-categories#categories-hierarchy).
	Label *string `json:"label,omitempty"`

	// Confidence score for the category classification. Higher values indicate greater confidence.
	Score *float64 `json:"score,omitempty"`

	// Information that helps to explain what contributed to the categories result.
	Explanation *CategoriesResultExplanation `json:"explanation,omitempty"`
}

// UnmarshalCategoriesResult unmarshals an instance of CategoriesResult from the specified map of raw messages.
func UnmarshalCategoriesResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CategoriesResult)
	err = core.UnmarshalPrimitive(m, "label", &obj.Label)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "score", &obj.Score)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "explanation", &obj.Explanation, UnmarshalCategoriesResultExplanation)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CategoriesResultExplanation : Information that helps to explain what contributed to the categories result.
type CategoriesResultExplanation struct {
	// An array of relevant text from the source that contributed to the categorization. The sorted array begins with the
	// phrase that contributed most significantly to the result, followed by phrases that were less and less impactful.
	RelevantText []CategoriesRelevantText `json:"relevant_text,omitempty"`
}

// UnmarshalCategoriesResultExplanation unmarshals an instance of CategoriesResultExplanation from the specified map of raw messages.
func UnmarshalCategoriesResultExplanation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CategoriesResultExplanation)
	err = core.UnmarshalModel(m, "relevant_text", &obj.RelevantText, UnmarshalCategoriesRelevantText)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ClassificationsModel : Classifications model.
type ClassificationsModel struct {
	// An optional name for the model.
	Name *string `json:"name,omitempty"`

	// An optional map of metadata key-value pairs to store with this model.
	UserMetadata map[string]interface{} `json:"user_metadata,omitempty"`

	// The 2-letter language code of this model.
	Language *string `json:"language" validate:"required"`

	// An optional description of the model.
	Description *string `json:"description,omitempty"`

	// An optional version string.
	ModelVersion *string `json:"model_version,omitempty"`

	// ID of the Watson Knowledge Studio workspace that deployed this model to Natural Language Understanding.
	WorkspaceID *string `json:"workspace_id,omitempty"`

	// The description of the version.
	VersionDescription *string `json:"version_description,omitempty"`

	// The service features that are supported by the custom model.
	Features []string `json:"features,omitempty"`

	// When the status is `available`, the model is ready to use.
	Status *string `json:"status" validate:"required"`

	// Unique model ID.
	ModelID *string `json:"model_id" validate:"required"`

	// dateTime indicating when the model was created.
	Created *strfmt.DateTime `json:"created" validate:"required"`

	Notices []Notice `json:"notices,omitempty"`

	// dateTime of last successful model training.
	LastTrained *strfmt.DateTime `json:"last_trained,omitempty"`

	// dateTime of last successful model deployment.
	LastDeployed *strfmt.DateTime `json:"last_deployed,omitempty"`
}

// Constants associated with the ClassificationsModel.Status property.
// When the status is `available`, the model is ready to use.
const (
	ClassificationsModelStatusAvailableConst = "available"
	ClassificationsModelStatusDeletedConst   = "deleted"
	ClassificationsModelStatusDeployingConst = "deploying"
	ClassificationsModelStatusErrorConst     = "error"
	ClassificationsModelStatusStartingConst  = "starting"
	ClassificationsModelStatusTrainingConst  = "training"
)

// UnmarshalClassificationsModel unmarshals an instance of ClassificationsModel from the specified map of raw messages.
func UnmarshalClassificationsModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ClassificationsModel)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "user_metadata", &obj.UserMetadata)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "language", &obj.Language)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "model_version", &obj.ModelVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "workspace_id", &obj.WorkspaceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version_description", &obj.VersionDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "features", &obj.Features)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "model_id", &obj.ModelID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created", &obj.Created)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "notices", &obj.Notices, UnmarshalNotice)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_trained", &obj.LastTrained)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_deployed", &obj.LastDeployed)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ClassificationsModelList : List of classifications models.
type ClassificationsModelList struct {
	// The classifications models.
	Models []ClassificationsModel `json:"models,omitempty"`
}

// UnmarshalClassificationsModelList unmarshals an instance of ClassificationsModelList from the specified map of raw messages.
func UnmarshalClassificationsModelList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ClassificationsModelList)
	err = core.UnmarshalModel(m, "models", &obj.Models, UnmarshalClassificationsModel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ClassificationsOptions : Returns text classifications for the content.
//
// Supported languages: English only.
type ClassificationsOptions struct {
	// (Beta) Enter a [custom
	// model](https://cloud.ibm.com/docs/natural-language-understanding?topic=natural-language-understanding-customizing)
	// ID of the classification model to be used.
	Model *string `json:"model,omitempty"`
}

// UnmarshalClassificationsOptions unmarshals an instance of ClassificationsOptions from the specified map of raw messages.
func UnmarshalClassificationsOptions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ClassificationsOptions)
	err = core.UnmarshalPrimitive(m, "model", &obj.Model)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ClassificationsResult : A classification of the analyzed text.
type ClassificationsResult struct {
	// Classification assigned to the text.
	ClassName *string `json:"class_name,omitempty"`

	// Confidence score for the classification. Higher values indicate greater confidence.
	Confidence *float64 `json:"confidence,omitempty"`
}

// UnmarshalClassificationsResult unmarshals an instance of ClassificationsResult from the specified map of raw messages.
func UnmarshalClassificationsResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ClassificationsResult)
	err = core.UnmarshalPrimitive(m, "class_name", &obj.ClassName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "confidence", &obj.Confidence)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConceptsOptions : Returns high-level concepts in the content. For example, a research paper about deep learning might return the
// concept, "Artificial Intelligence" although the term is not mentioned.
//
// Supported languages: English, French, German, Italian, Japanese, Korean, Portuguese, Spanish.
type ConceptsOptions struct {
	// Maximum number of concepts to return.
	Limit *int64 `json:"limit,omitempty"`
}

// UnmarshalConceptsOptions unmarshals an instance of ConceptsOptions from the specified map of raw messages.
func UnmarshalConceptsOptions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConceptsOptions)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConceptsResult : The general concepts referenced or alluded to in the analyzed text.
type ConceptsResult struct {
	// Name of the concept.
	Text *string `json:"text,omitempty"`

	// Relevance score between 0 and 1. Higher scores indicate greater relevance.
	Relevance *float64 `json:"relevance,omitempty"`

	// Link to the corresponding DBpedia resource.
	DbpediaResource *string `json:"dbpedia_resource,omitempty"`
}

// UnmarshalConceptsResult unmarshals an instance of ConceptsResult from the specified map of raw messages.
func UnmarshalConceptsResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConceptsResult)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "relevance", &obj.Relevance)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "dbpedia_resource", &obj.DbpediaResource)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateCategoriesModelOptions : The CreateCategoriesModel options.
type CreateCategoriesModelOptions struct {
	// The 2-letter language code of this model.
	Language *string `validate:"required"`

	// Training data in JSON format. For more information, see [Categories training data
	// requirements](https://cloud.ibm.com/docs/natural-language-understanding?topic=natural-language-understanding-categories##categories-training-data-requirements).
	TrainingData io.ReadCloser `validate:"required"`

	// The content type of trainingData.
	TrainingDataContentType *string

	// An optional name for the model.
	Name *string

	// An optional map of metadata key-value pairs to store with this model.
	UserMetadata map[string]interface{}

	// An optional description of the model.
	Description *string

	// An optional version string.
	ModelVersion *string

	// ID of the Watson Knowledge Studio workspace that deployed this model to Natural Language Understanding.
	WorkspaceID *string

	// The description of the version.
	VersionDescription *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateCategoriesModelOptions : Instantiate CreateCategoriesModelOptions
func (*NaturalLanguageUnderstandingV1) NewCreateCategoriesModelOptions(language string, trainingData io.ReadCloser) *CreateCategoriesModelOptions {
	return &CreateCategoriesModelOptions{
		Language:     core.StringPtr(language),
		TrainingData: trainingData,
	}
}

// SetLanguage : Allow user to set Language
func (options *CreateCategoriesModelOptions) SetLanguage(language string) *CreateCategoriesModelOptions {
	options.Language = core.StringPtr(language)
	return options
}

// SetTrainingData : Allow user to set TrainingData
func (options *CreateCategoriesModelOptions) SetTrainingData(trainingData io.ReadCloser) *CreateCategoriesModelOptions {
	options.TrainingData = trainingData
	return options
}

// SetTrainingDataContentType : Allow user to set TrainingDataContentType
func (options *CreateCategoriesModelOptions) SetTrainingDataContentType(trainingDataContentType string) *CreateCategoriesModelOptions {
	options.TrainingDataContentType = core.StringPtr(trainingDataContentType)
	return options
}

// SetName : Allow user to set Name
func (options *CreateCategoriesModelOptions) SetName(name string) *CreateCategoriesModelOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetUserMetadata : Allow user to set UserMetadata
func (options *CreateCategoriesModelOptions) SetUserMetadata(userMetadata map[string]interface{}) *CreateCategoriesModelOptions {
	options.UserMetadata = userMetadata
	return options
}

// SetDescription : Allow user to set Description
func (options *CreateCategoriesModelOptions) SetDescription(description string) *CreateCategoriesModelOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetModelVersion : Allow user to set ModelVersion
func (options *CreateCategoriesModelOptions) SetModelVersion(modelVersion string) *CreateCategoriesModelOptions {
	options.ModelVersion = core.StringPtr(modelVersion)
	return options
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *CreateCategoriesModelOptions) SetWorkspaceID(workspaceID string) *CreateCategoriesModelOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetVersionDescription : Allow user to set VersionDescription
func (options *CreateCategoriesModelOptions) SetVersionDescription(versionDescription string) *CreateCategoriesModelOptions {
	options.VersionDescription = core.StringPtr(versionDescription)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateCategoriesModelOptions) SetHeaders(param map[string]string) *CreateCategoriesModelOptions {
	options.Headers = param
	return options
}

// CreateClassificationsModelOptions : The CreateClassificationsModel options.
type CreateClassificationsModelOptions struct {
	// The 2-letter language code of this model.
	Language *string `validate:"required"`

	// Training data in JSON format. For more information, see [Classifications training data
	// requirements](https://cloud.ibm.com/docs/natural-language-understanding?topic=natural-language-understanding-classifications#classification-training-data-requirements).
	TrainingData io.ReadCloser `validate:"required"`

	// The content type of trainingData.
	TrainingDataContentType *string

	// An optional name for the model.
	Name *string

	// An optional map of metadata key-value pairs to store with this model.
	UserMetadata map[string]interface{}

	// An optional description of the model.
	Description *string

	// An optional version string.
	ModelVersion *string

	// ID of the Watson Knowledge Studio workspace that deployed this model to Natural Language Understanding.
	WorkspaceID *string

	// The description of the version.
	VersionDescription *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateClassificationsModelOptions : Instantiate CreateClassificationsModelOptions
func (*NaturalLanguageUnderstandingV1) NewCreateClassificationsModelOptions(language string, trainingData io.ReadCloser) *CreateClassificationsModelOptions {
	return &CreateClassificationsModelOptions{
		Language:     core.StringPtr(language),
		TrainingData: trainingData,
	}
}

// SetLanguage : Allow user to set Language
func (options *CreateClassificationsModelOptions) SetLanguage(language string) *CreateClassificationsModelOptions {
	options.Language = core.StringPtr(language)
	return options
}

// SetTrainingData : Allow user to set TrainingData
func (options *CreateClassificationsModelOptions) SetTrainingData(trainingData io.ReadCloser) *CreateClassificationsModelOptions {
	options.TrainingData = trainingData
	return options
}

// SetTrainingDataContentType : Allow user to set TrainingDataContentType
func (options *CreateClassificationsModelOptions) SetTrainingDataContentType(trainingDataContentType string) *CreateClassificationsModelOptions {
	options.TrainingDataContentType = core.StringPtr(trainingDataContentType)
	return options
}

// SetName : Allow user to set Name
func (options *CreateClassificationsModelOptions) SetName(name string) *CreateClassificationsModelOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetUserMetadata : Allow user to set UserMetadata
func (options *CreateClassificationsModelOptions) SetUserMetadata(userMetadata map[string]interface{}) *CreateClassificationsModelOptions {
	options.UserMetadata = userMetadata
	return options
}

// SetDescription : Allow user to set Description
func (options *CreateClassificationsModelOptions) SetDescription(description string) *CreateClassificationsModelOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetModelVersion : Allow user to set ModelVersion
func (options *CreateClassificationsModelOptions) SetModelVersion(modelVersion string) *CreateClassificationsModelOptions {
	options.ModelVersion = core.StringPtr(modelVersion)
	return options
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *CreateClassificationsModelOptions) SetWorkspaceID(workspaceID string) *CreateClassificationsModelOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetVersionDescription : Allow user to set VersionDescription
func (options *CreateClassificationsModelOptions) SetVersionDescription(versionDescription string) *CreateClassificationsModelOptions {
	options.VersionDescription = core.StringPtr(versionDescription)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateClassificationsModelOptions) SetHeaders(param map[string]string) *CreateClassificationsModelOptions {
	options.Headers = param
	return options
}

// CreateSentimentModelOptions : The CreateSentimentModel options.
type CreateSentimentModelOptions struct {
	// The 2-letter language code of this model.
	Language *string `validate:"required"`

	// Training data in CSV format. For more information, see [Sentiment training data
	// requirements](https://cloud.ibm.com/docs/natural-language-understanding?topic=natural-language-understanding-custom-sentiment#sentiment-training-data-requirements).
	TrainingData io.ReadCloser `validate:"required"`

	// An optional name for the model.
	Name *string

	// An optional map of metadata key-value pairs to store with this model.
	UserMetadata map[string]interface{}

	// An optional description of the model.
	Description *string

	// An optional version string.
	ModelVersion *string

	// ID of the Watson Knowledge Studio workspace that deployed this model to Natural Language Understanding.
	WorkspaceID *string

	// The description of the version.
	VersionDescription *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateSentimentModelOptions : Instantiate CreateSentimentModelOptions
func (*NaturalLanguageUnderstandingV1) NewCreateSentimentModelOptions(language string, trainingData io.ReadCloser) *CreateSentimentModelOptions {
	return &CreateSentimentModelOptions{
		Language:     core.StringPtr(language),
		TrainingData: trainingData,
	}
}

// SetLanguage : Allow user to set Language
func (options *CreateSentimentModelOptions) SetLanguage(language string) *CreateSentimentModelOptions {
	options.Language = core.StringPtr(language)
	return options
}

// SetTrainingData : Allow user to set TrainingData
func (options *CreateSentimentModelOptions) SetTrainingData(trainingData io.ReadCloser) *CreateSentimentModelOptions {
	options.TrainingData = trainingData
	return options
}

// SetName : Allow user to set Name
func (options *CreateSentimentModelOptions) SetName(name string) *CreateSentimentModelOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetUserMetadata : Allow user to set UserMetadata
func (options *CreateSentimentModelOptions) SetUserMetadata(userMetadata map[string]interface{}) *CreateSentimentModelOptions {
	options.UserMetadata = userMetadata
	return options
}

// SetDescription : Allow user to set Description
func (options *CreateSentimentModelOptions) SetDescription(description string) *CreateSentimentModelOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetModelVersion : Allow user to set ModelVersion
func (options *CreateSentimentModelOptions) SetModelVersion(modelVersion string) *CreateSentimentModelOptions {
	options.ModelVersion = core.StringPtr(modelVersion)
	return options
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *CreateSentimentModelOptions) SetWorkspaceID(workspaceID string) *CreateSentimentModelOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetVersionDescription : Allow user to set VersionDescription
func (options *CreateSentimentModelOptions) SetVersionDescription(versionDescription string) *CreateSentimentModelOptions {
	options.VersionDescription = core.StringPtr(versionDescription)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateSentimentModelOptions) SetHeaders(param map[string]string) *CreateSentimentModelOptions {
	options.Headers = param
	return options
}

// DeleteCategoriesModelOptions : The DeleteCategoriesModel options.
type DeleteCategoriesModelOptions struct {
	// ID of the model.
	ModelID *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteCategoriesModelOptions : Instantiate DeleteCategoriesModelOptions
func (*NaturalLanguageUnderstandingV1) NewDeleteCategoriesModelOptions(modelID string) *DeleteCategoriesModelOptions {
	return &DeleteCategoriesModelOptions{
		ModelID: core.StringPtr(modelID),
	}
}

// SetModelID : Allow user to set ModelID
func (options *DeleteCategoriesModelOptions) SetModelID(modelID string) *DeleteCategoriesModelOptions {
	options.ModelID = core.StringPtr(modelID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteCategoriesModelOptions) SetHeaders(param map[string]string) *DeleteCategoriesModelOptions {
	options.Headers = param
	return options
}

// DeleteClassificationsModelOptions : The DeleteClassificationsModel options.
type DeleteClassificationsModelOptions struct {
	// ID of the model.
	ModelID *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteClassificationsModelOptions : Instantiate DeleteClassificationsModelOptions
func (*NaturalLanguageUnderstandingV1) NewDeleteClassificationsModelOptions(modelID string) *DeleteClassificationsModelOptions {
	return &DeleteClassificationsModelOptions{
		ModelID: core.StringPtr(modelID),
	}
}

// SetModelID : Allow user to set ModelID
func (options *DeleteClassificationsModelOptions) SetModelID(modelID string) *DeleteClassificationsModelOptions {
	options.ModelID = core.StringPtr(modelID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteClassificationsModelOptions) SetHeaders(param map[string]string) *DeleteClassificationsModelOptions {
	options.Headers = param
	return options
}

// DeleteModelOptions : The DeleteModel options.
type DeleteModelOptions struct {
	// Model ID of the model to delete.
	ModelID *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteModelOptions : Instantiate DeleteModelOptions
func (*NaturalLanguageUnderstandingV1) NewDeleteModelOptions(modelID string) *DeleteModelOptions {
	return &DeleteModelOptions{
		ModelID: core.StringPtr(modelID),
	}
}

// SetModelID : Allow user to set ModelID
func (options *DeleteModelOptions) SetModelID(modelID string) *DeleteModelOptions {
	options.ModelID = core.StringPtr(modelID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteModelOptions) SetHeaders(param map[string]string) *DeleteModelOptions {
	options.Headers = param
	return options
}

// DeleteModelResults : Delete model results.
type DeleteModelResults struct {
	// model_id of the deleted model.
	Deleted *string `json:"deleted,omitempty"`
}

// UnmarshalDeleteModelResults unmarshals an instance of DeleteModelResults from the specified map of raw messages.
func UnmarshalDeleteModelResults(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteModelResults)
	err = core.UnmarshalPrimitive(m, "deleted", &obj.Deleted)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteSentimentModelOptions : The DeleteSentimentModel options.
type DeleteSentimentModelOptions struct {
	// ID of the model.
	ModelID *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteSentimentModelOptions : Instantiate DeleteSentimentModelOptions
func (*NaturalLanguageUnderstandingV1) NewDeleteSentimentModelOptions(modelID string) *DeleteSentimentModelOptions {
	return &DeleteSentimentModelOptions{
		ModelID: core.StringPtr(modelID),
	}
}

// SetModelID : Allow user to set ModelID
func (options *DeleteSentimentModelOptions) SetModelID(modelID string) *DeleteSentimentModelOptions {
	options.ModelID = core.StringPtr(modelID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteSentimentModelOptions) SetHeaders(param map[string]string) *DeleteSentimentModelOptions {
	options.Headers = param
	return options
}

// DisambiguationResult : Disambiguation information for the entity.
type DisambiguationResult struct {
	// Common entity name.
	Name *string `json:"name,omitempty"`

	// Link to the corresponding DBpedia resource.
	DbpediaResource *string `json:"dbpedia_resource,omitempty"`

	// Entity subtype information.
	Subtype []string `json:"subtype,omitempty"`
}

// UnmarshalDisambiguationResult unmarshals an instance of DisambiguationResult from the specified map of raw messages.
func UnmarshalDisambiguationResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DisambiguationResult)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "dbpedia_resource", &obj.DbpediaResource)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "subtype", &obj.Subtype)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DocumentEmotionResults : Emotion results for the document as a whole.
type DocumentEmotionResults struct {
	// Emotion results for the document as a whole.
	Emotion *EmotionScores `json:"emotion,omitempty"`
}

// UnmarshalDocumentEmotionResults unmarshals an instance of DocumentEmotionResults from the specified map of raw messages.
func UnmarshalDocumentEmotionResults(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DocumentEmotionResults)
	err = core.UnmarshalModel(m, "emotion", &obj.Emotion, UnmarshalEmotionScores)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DocumentSentimentResults : DocumentSentimentResults struct
type DocumentSentimentResults struct {
	// Indicates whether the sentiment is positive, neutral, or negative.
	Label *string `json:"label,omitempty"`

	// Sentiment score from -1 (negative) to 1 (positive).
	Score *float64 `json:"score,omitempty"`
}

// UnmarshalDocumentSentimentResults unmarshals an instance of DocumentSentimentResults from the specified map of raw messages.
func UnmarshalDocumentSentimentResults(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DocumentSentimentResults)
	err = core.UnmarshalPrimitive(m, "label", &obj.Label)
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

// EmotionOptions : Detects anger, disgust, fear, joy, or sadness that is conveyed in the content or by the context around target phrases
// specified in the targets parameter. You can analyze emotion for detected entities with `entities.emotion` and for
// keywords with `keywords.emotion`.
//
// Supported languages: English.
type EmotionOptions struct {
	// Set this to `false` to hide document-level emotion results.
	Document *bool `json:"document,omitempty"`

	// Emotion results will be returned for each target string that is found in the document.
	Targets []string `json:"targets,omitempty"`
}

// UnmarshalEmotionOptions unmarshals an instance of EmotionOptions from the specified map of raw messages.
func UnmarshalEmotionOptions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EmotionOptions)
	err = core.UnmarshalPrimitive(m, "document", &obj.Document)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "targets", &obj.Targets)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EmotionResult : The detected anger, disgust, fear, joy, or sadness that is conveyed by the content. Emotion information can be
// returned for detected entities, keywords, or user-specified target phrases found in the text.
type EmotionResult struct {
	// Emotion results for the document as a whole.
	Document *DocumentEmotionResults `json:"document,omitempty"`

	// Emotion results for specified targets.
	Targets []TargetedEmotionResults `json:"targets,omitempty"`
}

// UnmarshalEmotionResult unmarshals an instance of EmotionResult from the specified map of raw messages.
func UnmarshalEmotionResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EmotionResult)
	err = core.UnmarshalModel(m, "document", &obj.Document, UnmarshalDocumentEmotionResults)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "targets", &obj.Targets, UnmarshalTargetedEmotionResults)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EmotionScores : EmotionScores struct
type EmotionScores struct {
	// Anger score from 0 to 1. A higher score means that the text is more likely to convey anger.
	Anger *float64 `json:"anger,omitempty"`

	// Disgust score from 0 to 1. A higher score means that the text is more likely to convey disgust.
	Disgust *float64 `json:"disgust,omitempty"`

	// Fear score from 0 to 1. A higher score means that the text is more likely to convey fear.
	Fear *float64 `json:"fear,omitempty"`

	// Joy score from 0 to 1. A higher score means that the text is more likely to convey joy.
	Joy *float64 `json:"joy,omitempty"`

	// Sadness score from 0 to 1. A higher score means that the text is more likely to convey sadness.
	Sadness *float64 `json:"sadness,omitempty"`
}

// UnmarshalEmotionScores unmarshals an instance of EmotionScores from the specified map of raw messages.
func UnmarshalEmotionScores(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EmotionScores)
	err = core.UnmarshalPrimitive(m, "anger", &obj.Anger)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "disgust", &obj.Disgust)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "fear", &obj.Fear)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "joy", &obj.Joy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sadness", &obj.Sadness)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EntitiesOptions : Identifies people, cities, organizations, and other entities in the content. For more information, see [Entity types
// and
// subtypes](https://cloud.ibm.com/docs/natural-language-understanding?topic=natural-language-understanding-entity-types).
//
// Supported languages: English, French, German, Italian, Japanese, Korean, Portuguese, Russian, Spanish, Swedish.
// Arabic, Chinese, and Dutch are supported only through custom models.
type EntitiesOptions struct {
	// Maximum number of entities to return.
	Limit *int64 `json:"limit,omitempty"`

	// Set this to `true` to return locations of entity mentions.
	Mentions *bool `json:"mentions,omitempty"`

	// Enter a [custom
	// model](https://cloud.ibm.com/docs/natural-language-understanding?topic=natural-language-understanding-customizing)
	// ID to override the standard entity detection model.
	Model *string `json:"model,omitempty"`

	// Set this to `true` to return sentiment information for detected entities.
	Sentiment *bool `json:"sentiment,omitempty"`

	// Set this to `true` to analyze emotion for detected keywords.
	Emotion *bool `json:"emotion,omitempty"`
}

// UnmarshalEntitiesOptions unmarshals an instance of EntitiesOptions from the specified map of raw messages.
func UnmarshalEntitiesOptions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EntitiesOptions)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "mentions", &obj.Mentions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "model", &obj.Model)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sentiment", &obj.Sentiment)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "emotion", &obj.Emotion)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EntitiesResult : The important people, places, geopolitical entities and other types of entities in your content.
type EntitiesResult struct {
	// Entity type.
	Type *string `json:"type,omitempty"`

	// The name of the entity.
	Text *string `json:"text,omitempty"`

	// Relevance score from 0 to 1. Higher values indicate greater relevance.
	Relevance *float64 `json:"relevance,omitempty"`

	// Confidence in the entity identification from 0 to 1. Higher values indicate higher confidence. In standard entities
	// requests, confidence is returned only for English text. All entities requests that use custom models return the
	// confidence score.
	Confidence *float64 `json:"confidence,omitempty"`

	// Entity mentions and locations.
	Mentions []EntityMention `json:"mentions,omitempty"`

	// How many times the entity was mentioned in the text.
	Count *int64 `json:"count,omitempty"`

	// Emotion analysis results for the entity, enabled with the `emotion` option.
	Emotion *EmotionScores `json:"emotion,omitempty"`

	// Sentiment analysis results for the entity, enabled with the `sentiment` option.
	Sentiment *FeatureSentimentResults `json:"sentiment,omitempty"`

	// Disambiguation information for the entity.
	Disambiguation *DisambiguationResult `json:"disambiguation,omitempty"`
}

// UnmarshalEntitiesResult unmarshals an instance of EntitiesResult from the specified map of raw messages.
func UnmarshalEntitiesResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EntitiesResult)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "relevance", &obj.Relevance)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "confidence", &obj.Confidence)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "mentions", &obj.Mentions, UnmarshalEntityMention)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "count", &obj.Count)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "emotion", &obj.Emotion, UnmarshalEmotionScores)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "sentiment", &obj.Sentiment, UnmarshalFeatureSentimentResults)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "disambiguation", &obj.Disambiguation, UnmarshalDisambiguationResult)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EntityMention : EntityMention struct
type EntityMention struct {
	// Entity mention text.
	Text *string `json:"text,omitempty"`

	// Character offsets indicating the beginning and end of the mention in the analyzed text.
	Location []int64 `json:"location,omitempty"`

	// Confidence in the entity identification from 0 to 1. Higher values indicate higher confidence. In standard entities
	// requests, confidence is returned only for English text. All entities requests that use custom models return the
	// confidence score.
	Confidence *float64 `json:"confidence,omitempty"`
}

// UnmarshalEntityMention unmarshals an instance of EntityMention from the specified map of raw messages.
func UnmarshalEntityMention(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EntityMention)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "location", &obj.Location)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "confidence", &obj.Confidence)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FeatureSentimentResults : FeatureSentimentResults struct
type FeatureSentimentResults struct {
	// Sentiment score from -1 (negative) to 1 (positive).
	Score *float64 `json:"score,omitempty"`
}

// UnmarshalFeatureSentimentResults unmarshals an instance of FeatureSentimentResults from the specified map of raw messages.
func UnmarshalFeatureSentimentResults(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FeatureSentimentResults)
	err = core.UnmarshalPrimitive(m, "score", &obj.Score)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Features : Analysis features and options.
type Features struct {
	// Returns text classifications for the content.
	//
	// Supported languages: English only.
	Classifications *ClassificationsOptions `json:"classifications,omitempty"`

	// Returns high-level concepts in the content. For example, a research paper about deep learning might return the
	// concept, "Artificial Intelligence" although the term is not mentioned.
	//
	// Supported languages: English, French, German, Italian, Japanese, Korean, Portuguese, Spanish.
	Concepts *ConceptsOptions `json:"concepts,omitempty"`

	// Detects anger, disgust, fear, joy, or sadness that is conveyed in the content or by the context around target
	// phrases specified in the targets parameter. You can analyze emotion for detected entities with `entities.emotion`
	// and for keywords with `keywords.emotion`.
	//
	// Supported languages: English.
	Emotion *EmotionOptions `json:"emotion,omitempty"`

	// Identifies people, cities, organizations, and other entities in the content. For more information, see [Entity types
	// and
	// subtypes](https://cloud.ibm.com/docs/natural-language-understanding?topic=natural-language-understanding-entity-types).
	//
	// Supported languages: English, French, German, Italian, Japanese, Korean, Portuguese, Russian, Spanish, Swedish.
	// Arabic, Chinese, and Dutch are supported only through custom models.
	Entities *EntitiesOptions `json:"entities,omitempty"`

	// Returns important keywords in the content.
	//
	// Supported languages: English, French, German, Italian, Japanese, Korean, Portuguese, Russian, Spanish, Swedish.
	Keywords *KeywordsOptions `json:"keywords,omitempty"`

	// Returns information from the document, including author name, title, RSS/ATOM feeds, prominent page image, and
	// publication date. Supports URL and HTML input types only.
	Metadata interface{} `json:"metadata,omitempty"`

	// Recognizes when two entities are related and identifies the type of relation. For example, an `awardedTo` relation
	// might connect the entities "Nobel Prize" and "Albert Einstein". For more information, see [Relation
	// types](https://cloud.ibm.com/docs/natural-language-understanding?topic=natural-language-understanding-relations).
	//
	// Supported languages: Arabic, English, German, Japanese, Korean, Spanish. Chinese, Dutch, French, Italian, and
	// Portuguese custom models are also supported.
	Relations *RelationsOptions `json:"relations,omitempty"`

	// Parses sentences into subject, action, and object form.
	//
	// Supported languages: English, German, Japanese, Korean, Spanish.
	SemanticRoles *SemanticRolesOptions `json:"semantic_roles,omitempty"`

	// Analyzes the general sentiment of your content or the sentiment toward specific target phrases. You can analyze
	// sentiment for detected entities with `entities.sentiment` and for keywords with `keywords.sentiment`.
	//
	//  Supported languages: Arabic, English, French, German, Italian, Japanese, Korean, Portuguese, Russian, Spanish.
	Sentiment *SentimentOptions `json:"sentiment,omitempty"`

	// (Experimental) Returns a summary of content.
	//
	// Supported languages: English only.
	Summarization *SummarizationOptions `json:"summarization,omitempty"`

	// Returns a five-level taxonomy of the content. The top three categories are returned.
	//
	// Supported languages: Arabic, English, French, German, Italian, Japanese, Korean, Portuguese, Spanish.
	Categories *CategoriesOptions `json:"categories,omitempty"`

	// Returns tokens and sentences from the input text.
	Syntax *SyntaxOptions `json:"syntax,omitempty"`
}

// UnmarshalFeatures unmarshals an instance of Features from the specified map of raw messages.
func UnmarshalFeatures(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Features)
	err = core.UnmarshalModel(m, "classifications", &obj.Classifications, UnmarshalClassificationsOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "concepts", &obj.Concepts, UnmarshalConceptsOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "emotion", &obj.Emotion, UnmarshalEmotionOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "entities", &obj.Entities, UnmarshalEntitiesOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "keywords", &obj.Keywords, UnmarshalKeywordsOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalMetadataOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "relations", &obj.Relations, UnmarshalRelationsOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "semantic_roles", &obj.SemanticRoles, UnmarshalSemanticRolesOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "sentiment", &obj.Sentiment, UnmarshalSentimentOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "summarization", &obj.Summarization, UnmarshalSummarizationOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "categories", &obj.Categories, UnmarshalCategoriesOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "syntax", &obj.Syntax, UnmarshalSyntaxOptions)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FeaturesResultsMetadata : Webpage metadata, such as the author and the title of the page.
type FeaturesResultsMetadata struct {
	// The authors of the document.
	Authors []Author `json:"authors,omitempty"`

	// The publication date in the format ISO 8601.
	PublicationDate *string `json:"publication_date,omitempty"`

	// The title of the document.
	Title *string `json:"title,omitempty"`

	// URL of a prominent image on the webpage.
	Image *string `json:"image,omitempty"`

	// RSS/ATOM feeds found on the webpage.
	Feeds []Feed `json:"feeds,omitempty"`
}

// UnmarshalFeaturesResultsMetadata unmarshals an instance of FeaturesResultsMetadata from the specified map of raw messages.
func UnmarshalFeaturesResultsMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FeaturesResultsMetadata)
	err = core.UnmarshalModel(m, "authors", &obj.Authors, UnmarshalAuthor)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "publication_date", &obj.PublicationDate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "title", &obj.Title)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "image", &obj.Image)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "feeds", &obj.Feeds, UnmarshalFeed)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Feed : RSS or ATOM feed found on the webpage.
type Feed struct {
	// URL of the RSS or ATOM feed.
	Link *string `json:"link,omitempty"`
}

// UnmarshalFeed unmarshals an instance of Feed from the specified map of raw messages.
func UnmarshalFeed(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Feed)
	err = core.UnmarshalPrimitive(m, "link", &obj.Link)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetCategoriesModelOptions : The GetCategoriesModel options.
type GetCategoriesModelOptions struct {
	// ID of the model.
	ModelID *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCategoriesModelOptions : Instantiate GetCategoriesModelOptions
func (*NaturalLanguageUnderstandingV1) NewGetCategoriesModelOptions(modelID string) *GetCategoriesModelOptions {
	return &GetCategoriesModelOptions{
		ModelID: core.StringPtr(modelID),
	}
}

// SetModelID : Allow user to set ModelID
func (options *GetCategoriesModelOptions) SetModelID(modelID string) *GetCategoriesModelOptions {
	options.ModelID = core.StringPtr(modelID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetCategoriesModelOptions) SetHeaders(param map[string]string) *GetCategoriesModelOptions {
	options.Headers = param
	return options
}

// GetClassificationsModelOptions : The GetClassificationsModel options.
type GetClassificationsModelOptions struct {
	// ID of the model.
	ModelID *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetClassificationsModelOptions : Instantiate GetClassificationsModelOptions
func (*NaturalLanguageUnderstandingV1) NewGetClassificationsModelOptions(modelID string) *GetClassificationsModelOptions {
	return &GetClassificationsModelOptions{
		ModelID: core.StringPtr(modelID),
	}
}

// SetModelID : Allow user to set ModelID
func (options *GetClassificationsModelOptions) SetModelID(modelID string) *GetClassificationsModelOptions {
	options.ModelID = core.StringPtr(modelID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetClassificationsModelOptions) SetHeaders(param map[string]string) *GetClassificationsModelOptions {
	options.Headers = param
	return options
}

// GetSentimentModelOptions : The GetSentimentModel options.
type GetSentimentModelOptions struct {
	// ID of the model.
	ModelID *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetSentimentModelOptions : Instantiate GetSentimentModelOptions
func (*NaturalLanguageUnderstandingV1) NewGetSentimentModelOptions(modelID string) *GetSentimentModelOptions {
	return &GetSentimentModelOptions{
		ModelID: core.StringPtr(modelID),
	}
}

// SetModelID : Allow user to set ModelID
func (options *GetSentimentModelOptions) SetModelID(modelID string) *GetSentimentModelOptions {
	options.ModelID = core.StringPtr(modelID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetSentimentModelOptions) SetHeaders(param map[string]string) *GetSentimentModelOptions {
	options.Headers = param
	return options
}

// KeywordsOptions : Returns important keywords in the content.
//
// Supported languages: English, French, German, Italian, Japanese, Korean, Portuguese, Russian, Spanish, Swedish.
type KeywordsOptions struct {
	// Maximum number of keywords to return.
	Limit *int64 `json:"limit,omitempty"`

	// Set this to `true` to return sentiment information for detected keywords.
	Sentiment *bool `json:"sentiment,omitempty"`

	// Set this to `true` to analyze emotion for detected keywords.
	Emotion *bool `json:"emotion,omitempty"`
}

// UnmarshalKeywordsOptions unmarshals an instance of KeywordsOptions from the specified map of raw messages.
func UnmarshalKeywordsOptions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeywordsOptions)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sentiment", &obj.Sentiment)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "emotion", &obj.Emotion)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeywordsResult : The important keywords in the content, organized by relevance.
type KeywordsResult struct {
	// Number of times the keyword appears in the analyzed text.
	Count *int64 `json:"count,omitempty"`

	// Relevance score from 0 to 1. Higher values indicate greater relevance.
	Relevance *float64 `json:"relevance,omitempty"`

	// The keyword text.
	Text *string `json:"text,omitempty"`

	// Emotion analysis results for the keyword, enabled with the `emotion` option.
	Emotion *EmotionScores `json:"emotion,omitempty"`

	// Sentiment analysis results for the keyword, enabled with the `sentiment` option.
	Sentiment *FeatureSentimentResults `json:"sentiment,omitempty"`
}

// UnmarshalKeywordsResult unmarshals an instance of KeywordsResult from the specified map of raw messages.
func UnmarshalKeywordsResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeywordsResult)
	err = core.UnmarshalPrimitive(m, "count", &obj.Count)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "relevance", &obj.Relevance)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "emotion", &obj.Emotion, UnmarshalEmotionScores)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "sentiment", &obj.Sentiment, UnmarshalFeatureSentimentResults)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListCategoriesModelsOptions : The ListCategoriesModels options.
type ListCategoriesModelsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListCategoriesModelsOptions : Instantiate ListCategoriesModelsOptions
func (*NaturalLanguageUnderstandingV1) NewListCategoriesModelsOptions() *ListCategoriesModelsOptions {
	return &ListCategoriesModelsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListCategoriesModelsOptions) SetHeaders(param map[string]string) *ListCategoriesModelsOptions {
	options.Headers = param
	return options
}

// ListClassificationsModelsOptions : The ListClassificationsModels options.
type ListClassificationsModelsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListClassificationsModelsOptions : Instantiate ListClassificationsModelsOptions
func (*NaturalLanguageUnderstandingV1) NewListClassificationsModelsOptions() *ListClassificationsModelsOptions {
	return &ListClassificationsModelsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListClassificationsModelsOptions) SetHeaders(param map[string]string) *ListClassificationsModelsOptions {
	options.Headers = param
	return options
}

// ListClassificationsModelsResponse : ListClassificationsModelsResponse struct
type ListClassificationsModelsResponse struct {
	Models []ClassificationsModelList `json:"models,omitempty"`
}

// UnmarshalListClassificationsModelsResponse unmarshals an instance of ListClassificationsModelsResponse from the specified map of raw messages.
func UnmarshalListClassificationsModelsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListClassificationsModelsResponse)
	err = core.UnmarshalModel(m, "models", &obj.Models, UnmarshalClassificationsModelList)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListModelsOptions : The ListModels options.
type ListModelsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListModelsOptions : Instantiate ListModelsOptions
func (*NaturalLanguageUnderstandingV1) NewListModelsOptions() *ListModelsOptions {
	return &ListModelsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListModelsOptions) SetHeaders(param map[string]string) *ListModelsOptions {
	options.Headers = param
	return options
}

// ListModelsResults : Custom models that are available for entities and relations.
type ListModelsResults struct {
	// An array of available models.
	Models []Model `json:"models,omitempty"`
}

// UnmarshalListModelsResults unmarshals an instance of ListModelsResults from the specified map of raw messages.
func UnmarshalListModelsResults(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListModelsResults)
	err = core.UnmarshalModel(m, "models", &obj.Models, UnmarshalModel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListSentimentModelsOptions : The ListSentimentModels options.
type ListSentimentModelsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListSentimentModelsOptions : Instantiate ListSentimentModelsOptions
func (*NaturalLanguageUnderstandingV1) NewListSentimentModelsOptions() *ListSentimentModelsOptions {
	return &ListSentimentModelsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListSentimentModelsOptions) SetHeaders(param map[string]string) *ListSentimentModelsOptions {
	options.Headers = param
	return options
}

// ListSentimentModelsResponse : ListSentimentModelsResponse struct
type ListSentimentModelsResponse struct {
	Models []SentimentModel `json:"models,omitempty"`
}

// UnmarshalListSentimentModelsResponse unmarshals an instance of ListSentimentModelsResponse from the specified map of raw messages.
func UnmarshalListSentimentModelsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListSentimentModelsResponse)
	err = core.UnmarshalModel(m, "models", &obj.Models, UnmarshalSentimentModel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MetadataOptions : Returns information from the document, including author name, title, RSS/ATOM feeds, prominent page image, and
// publication date. Supports URL and HTML input types only.
type MetadataOptions struct {
}

// UnmarshalMetadataOptions unmarshals an instance of MetadataOptions from the specified map of raw messages.
func UnmarshalMetadataOptions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MetadataOptions)
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Model : Model struct
type Model struct {
	// When the status is `available`, the model is ready to use.
	Status *string `json:"status,omitempty"`

	// Unique model ID.
	ModelID *string `json:"model_id,omitempty"`

	// ISO 639-1 code that indicates the language of the model.
	Language *string `json:"language,omitempty"`

	// Model description.
	Description *string `json:"description,omitempty"`

	// ID of the Watson Knowledge Studio workspace that deployed this model to Natural Language Understanding.
	WorkspaceID *string `json:"workspace_id,omitempty"`

	// The model version, if it was manually provided in Watson Knowledge Studio.
	ModelVersion *string `json:"model_version,omitempty"`

	// Deprecated — use `model_version`.
	Version *string `json:"version,omitempty"`

	// The description of the version, if it was manually provided in Watson Knowledge Studio.
	VersionDescription *string `json:"version_description,omitempty"`

	// A dateTime indicating when the model was created.
	Created *strfmt.DateTime `json:"created,omitempty"`
}

// Constants associated with the Model.Status property.
// When the status is `available`, the model is ready to use.
const (
	ModelStatusAvailableConst = "available"
	ModelStatusDeletedConst   = "deleted"
	ModelStatusDeployingConst = "deploying"
	ModelStatusErrorConst     = "error"
	ModelStatusStartingConst  = "starting"
	ModelStatusTrainingConst  = "training"
)

// UnmarshalModel unmarshals an instance of Model from the specified map of raw messages.
func UnmarshalModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Model)
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "model_id", &obj.ModelID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "language", &obj.Language)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "workspace_id", &obj.WorkspaceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "model_version", &obj.ModelVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version_description", &obj.VersionDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created", &obj.Created)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Notice : A list of messages describing model training issues when model status is `error`.
type Notice struct {
	// Describes deficiencies or inconsistencies in training data.
	Message *string `json:"message,omitempty"`
}

// UnmarshalNotice unmarshals an instance of Notice from the specified map of raw messages.
func UnmarshalNotice(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Notice)
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RelationArgument : RelationArgument struct
type RelationArgument struct {
	// An array of extracted entities.
	Entities []RelationEntity `json:"entities,omitempty"`

	// Character offsets indicating the beginning and end of the mention in the analyzed text.
	Location []int64 `json:"location,omitempty"`

	// Text that corresponds to the argument.
	Text *string `json:"text,omitempty"`
}

// UnmarshalRelationArgument unmarshals an instance of RelationArgument from the specified map of raw messages.
func UnmarshalRelationArgument(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RelationArgument)
	err = core.UnmarshalModel(m, "entities", &obj.Entities, UnmarshalRelationEntity)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "location", &obj.Location)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RelationEntity : An entity that corresponds with an argument in a relation.
type RelationEntity struct {
	// Text that corresponds to the entity.
	Text *string `json:"text,omitempty"`

	// Entity type.
	Type *string `json:"type,omitempty"`
}

// UnmarshalRelationEntity unmarshals an instance of RelationEntity from the specified map of raw messages.
func UnmarshalRelationEntity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RelationEntity)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RelationsOptions : Recognizes when two entities are related and identifies the type of relation. For example, an `awardedTo` relation
// might connect the entities "Nobel Prize" and "Albert Einstein". For more information, see [Relation
// types](https://cloud.ibm.com/docs/natural-language-understanding?topic=natural-language-understanding-relations).
//
// Supported languages: Arabic, English, German, Japanese, Korean, Spanish. Chinese, Dutch, French, Italian, and
// Portuguese custom models are also supported.
type RelationsOptions struct {
	// Enter a [custom
	// model](https://cloud.ibm.com/docs/natural-language-understanding?topic=natural-language-understanding-customizing)
	// ID to override the default model.
	Model *string `json:"model,omitempty"`
}

// UnmarshalRelationsOptions unmarshals an instance of RelationsOptions from the specified map of raw messages.
func UnmarshalRelationsOptions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RelationsOptions)
	err = core.UnmarshalPrimitive(m, "model", &obj.Model)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RelationsResult : The relations between entities found in the content.
type RelationsResult struct {
	// Confidence score for the relation. Higher values indicate greater confidence.
	Score *float64 `json:"score,omitempty"`

	// The sentence that contains the relation.
	Sentence *string `json:"sentence,omitempty"`

	// The type of the relation.
	Type *string `json:"type,omitempty"`

	// Entity mentions that are involved in the relation.
	Arguments []RelationArgument `json:"arguments,omitempty"`
}

// UnmarshalRelationsResult unmarshals an instance of RelationsResult from the specified map of raw messages.
func UnmarshalRelationsResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RelationsResult)
	err = core.UnmarshalPrimitive(m, "score", &obj.Score)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sentence", &obj.Sentence)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "arguments", &obj.Arguments, UnmarshalRelationArgument)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SemanticRolesEntity : SemanticRolesEntity struct
type SemanticRolesEntity struct {
	// Entity type.
	Type *string `json:"type,omitempty"`

	// The entity text.
	Text *string `json:"text,omitempty"`
}

// UnmarshalSemanticRolesEntity unmarshals an instance of SemanticRolesEntity from the specified map of raw messages.
func UnmarshalSemanticRolesEntity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SemanticRolesEntity)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SemanticRolesKeyword : SemanticRolesKeyword struct
type SemanticRolesKeyword struct {
	// The keyword text.
	Text *string `json:"text,omitempty"`
}

// UnmarshalSemanticRolesKeyword unmarshals an instance of SemanticRolesKeyword from the specified map of raw messages.
func UnmarshalSemanticRolesKeyword(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SemanticRolesKeyword)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SemanticRolesOptions : Parses sentences into subject, action, and object form.
//
// Supported languages: English, German, Japanese, Korean, Spanish.
type SemanticRolesOptions struct {
	// Maximum number of semantic_roles results to return.
	Limit *int64 `json:"limit,omitempty"`

	// Set this to `true` to return keyword information for subjects and objects.
	Keywords *bool `json:"keywords,omitempty"`

	// Set this to `true` to return entity information for subjects and objects.
	Entities *bool `json:"entities,omitempty"`
}

// UnmarshalSemanticRolesOptions unmarshals an instance of SemanticRolesOptions from the specified map of raw messages.
func UnmarshalSemanticRolesOptions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SemanticRolesOptions)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "keywords", &obj.Keywords)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "entities", &obj.Entities)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SemanticRolesResult : The object containing the actions and the objects the actions act upon.
type SemanticRolesResult struct {
	// Sentence from the source that contains the subject, action, and object.
	Sentence *string `json:"sentence,omitempty"`

	// The extracted subject from the sentence.
	Subject *SemanticRolesResultSubject `json:"subject,omitempty"`

	// The extracted action from the sentence.
	Action *SemanticRolesResultAction `json:"action,omitempty"`

	// The extracted object from the sentence.
	Object *SemanticRolesResultObject `json:"object,omitempty"`
}

// UnmarshalSemanticRolesResult unmarshals an instance of SemanticRolesResult from the specified map of raw messages.
func UnmarshalSemanticRolesResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SemanticRolesResult)
	err = core.UnmarshalPrimitive(m, "sentence", &obj.Sentence)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "subject", &obj.Subject, UnmarshalSemanticRolesResultSubject)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "action", &obj.Action, UnmarshalSemanticRolesResultAction)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "object", &obj.Object, UnmarshalSemanticRolesResultObject)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SemanticRolesResultAction : The extracted action from the sentence.
type SemanticRolesResultAction struct {
	// Analyzed text that corresponds to the action.
	Text *string `json:"text,omitempty"`

	// normalized version of the action.
	Normalized *string `json:"normalized,omitempty"`

	Verb *SemanticRolesVerb `json:"verb,omitempty"`
}

// UnmarshalSemanticRolesResultAction unmarshals an instance of SemanticRolesResultAction from the specified map of raw messages.
func UnmarshalSemanticRolesResultAction(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SemanticRolesResultAction)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "normalized", &obj.Normalized)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "verb", &obj.Verb, UnmarshalSemanticRolesVerb)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SemanticRolesResultObject : The extracted object from the sentence.
type SemanticRolesResultObject struct {
	// Object text.
	Text *string `json:"text,omitempty"`

	// An array of extracted keywords.
	Keywords []SemanticRolesKeyword `json:"keywords,omitempty"`
}

// UnmarshalSemanticRolesResultObject unmarshals an instance of SemanticRolesResultObject from the specified map of raw messages.
func UnmarshalSemanticRolesResultObject(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SemanticRolesResultObject)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "keywords", &obj.Keywords, UnmarshalSemanticRolesKeyword)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SemanticRolesResultSubject : The extracted subject from the sentence.
type SemanticRolesResultSubject struct {
	// Text that corresponds to the subject role.
	Text *string `json:"text,omitempty"`

	// An array of extracted entities.
	Entities []SemanticRolesEntity `json:"entities,omitempty"`

	// An array of extracted keywords.
	Keywords []SemanticRolesKeyword `json:"keywords,omitempty"`
}

// UnmarshalSemanticRolesResultSubject unmarshals an instance of SemanticRolesResultSubject from the specified map of raw messages.
func UnmarshalSemanticRolesResultSubject(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SemanticRolesResultSubject)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "entities", &obj.Entities, UnmarshalSemanticRolesEntity)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "keywords", &obj.Keywords, UnmarshalSemanticRolesKeyword)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SemanticRolesVerb : SemanticRolesVerb struct
type SemanticRolesVerb struct {
	// The keyword text.
	Text *string `json:"text,omitempty"`

	// Verb tense.
	Tense *string `json:"tense,omitempty"`
}

// UnmarshalSemanticRolesVerb unmarshals an instance of SemanticRolesVerb from the specified map of raw messages.
func UnmarshalSemanticRolesVerb(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SemanticRolesVerb)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tense", &obj.Tense)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SentenceResult : SentenceResult struct
type SentenceResult struct {
	// The sentence.
	Text *string `json:"text,omitempty"`

	// Character offsets indicating the beginning and end of the sentence in the analyzed text.
	Location []int64 `json:"location,omitempty"`
}

// UnmarshalSentenceResult unmarshals an instance of SentenceResult from the specified map of raw messages.
func UnmarshalSentenceResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SentenceResult)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "location", &obj.Location)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SentimentModel : SentimentModel struct
type SentimentModel struct {
	// The service features that are supported by the custom model.
	Features []string `json:"features,omitempty"`

	// When the status is `available`, the model is ready to use.
	Status *string `json:"status,omitempty"`

	// Unique model ID.
	ModelID *string `json:"model_id,omitempty"`

	// dateTime indicating when the model was created.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// dateTime of last successful model training.
	LastTrained *strfmt.DateTime `json:"last_trained,omitempty"`

	// dateTime of last successful model deployment.
	LastDeployed *strfmt.DateTime `json:"last_deployed,omitempty"`

	// A name for the model.
	Name *string `json:"name,omitempty"`

	// An optional map of metadata key-value pairs to store with this model.
	UserMetadata map[string]interface{} `json:"user_metadata,omitempty"`

	// The 2-letter language code of this model.
	Language *string `json:"language,omitempty"`

	// An optional description of the model.
	Description *string `json:"description,omitempty"`

	// An optional version string.
	ModelVersion *string `json:"model_version,omitempty"`

	Notices []Notice `json:"notices,omitempty"`

	// ID of the Watson Knowledge Studio workspace that deployed this model to Natural Language Understanding.
	WorkspaceID *string `json:"workspace_id,omitempty"`

	// The description of the version.
	VersionDescription *string `json:"version_description,omitempty"`
}

// Constants associated with the SentimentModel.Status property.
// When the status is `available`, the model is ready to use.
const (
	SentimentModelStatusAvailableConst = "available"
	SentimentModelStatusDeletedConst   = "deleted"
	SentimentModelStatusDeployingConst = "deploying"
	SentimentModelStatusErrorConst     = "error"
	SentimentModelStatusStartingConst  = "starting"
	SentimentModelStatusTrainingConst  = "training"
)

// UnmarshalSentimentModel unmarshals an instance of SentimentModel from the specified map of raw messages.
func UnmarshalSentimentModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SentimentModel)
	err = core.UnmarshalPrimitive(m, "features", &obj.Features)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "model_id", &obj.ModelID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created", &obj.Created)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_trained", &obj.LastTrained)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_deployed", &obj.LastDeployed)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "user_metadata", &obj.UserMetadata)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "language", &obj.Language)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "model_version", &obj.ModelVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "notices", &obj.Notices, UnmarshalNotice)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "workspace_id", &obj.WorkspaceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version_description", &obj.VersionDescription)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SentimentOptions : Analyzes the general sentiment of your content or the sentiment toward specific target phrases. You can analyze
// sentiment for detected entities with `entities.sentiment` and for keywords with `keywords.sentiment`.
//
//  Supported languages: Arabic, English, French, German, Italian, Japanese, Korean, Portuguese, Russian, Spanish.
type SentimentOptions struct {
	// Set this to `false` to hide document-level sentiment results.
	Document *bool `json:"document,omitempty"`

	// Sentiment results will be returned for each target string that is found in the document.
	Targets []string `json:"targets,omitempty"`

	// (Beta) Enter a [custom
	// model](https://cloud.ibm.com/docs/natural-language-understanding?topic=natural-language-understanding-customizing)
	// ID to override the standard sentiment model for all sentiment analysis operations in the request, including targeted
	// sentiment for entities and keywords.
	Model *string `json:"model,omitempty"`
}

// UnmarshalSentimentOptions unmarshals an instance of SentimentOptions from the specified map of raw messages.
func UnmarshalSentimentOptions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SentimentOptions)
	err = core.UnmarshalPrimitive(m, "document", &obj.Document)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "targets", &obj.Targets)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "model", &obj.Model)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SentimentResult : The sentiment of the content.
type SentimentResult struct {
	// The document level sentiment.
	Document *DocumentSentimentResults `json:"document,omitempty"`

	// The targeted sentiment to analyze.
	Targets []TargetedSentimentResults `json:"targets,omitempty"`
}

// UnmarshalSentimentResult unmarshals an instance of SentimentResult from the specified map of raw messages.
func UnmarshalSentimentResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SentimentResult)
	err = core.UnmarshalModel(m, "document", &obj.Document, UnmarshalDocumentSentimentResults)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "targets", &obj.Targets, UnmarshalTargetedSentimentResults)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SummarizationOptions : (Experimental) Returns a summary of content.
//
// Supported languages: English only.
type SummarizationOptions struct {
	// Maximum number of summary sentences to return.
	Limit *int64 `json:"limit,omitempty"`
}

// UnmarshalSummarizationOptions unmarshals an instance of SummarizationOptions from the specified map of raw messages.
func UnmarshalSummarizationOptions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SummarizationOptions)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SyntaxOptions : Returns tokens and sentences from the input text.
type SyntaxOptions struct {
	// Tokenization options.
	Tokens *SyntaxOptionsTokens `json:"tokens,omitempty"`

	// Set this to `true` to return sentence information.
	Sentences *bool `json:"sentences,omitempty"`
}

// UnmarshalSyntaxOptions unmarshals an instance of SyntaxOptions from the specified map of raw messages.
func UnmarshalSyntaxOptions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SyntaxOptions)
	err = core.UnmarshalModel(m, "tokens", &obj.Tokens, UnmarshalSyntaxOptionsTokens)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sentences", &obj.Sentences)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SyntaxOptionsTokens : Tokenization options.
type SyntaxOptionsTokens struct {
	// Set this to `true` to return the lemma for each token.
	Lemma *bool `json:"lemma,omitempty"`

	// Set this to `true` to return the part of speech for each token.
	PartOfSpeech *bool `json:"part_of_speech,omitempty"`
}

// UnmarshalSyntaxOptionsTokens unmarshals an instance of SyntaxOptionsTokens from the specified map of raw messages.
func UnmarshalSyntaxOptionsTokens(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SyntaxOptionsTokens)
	err = core.UnmarshalPrimitive(m, "lemma", &obj.Lemma)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "part_of_speech", &obj.PartOfSpeech)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SyntaxResult : Tokens and sentences returned from syntax analysis.
type SyntaxResult struct {
	Tokens []TokenResult `json:"tokens,omitempty"`

	Sentences []SentenceResult `json:"sentences,omitempty"`
}

// UnmarshalSyntaxResult unmarshals an instance of SyntaxResult from the specified map of raw messages.
func UnmarshalSyntaxResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SyntaxResult)
	err = core.UnmarshalModel(m, "tokens", &obj.Tokens, UnmarshalTokenResult)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "sentences", &obj.Sentences, UnmarshalSentenceResult)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TargetedEmotionResults : Emotion results for a specified target.
type TargetedEmotionResults struct {
	// Targeted text.
	Text *string `json:"text,omitempty"`

	// The emotion results for the target.
	Emotion *EmotionScores `json:"emotion,omitempty"`
}

// UnmarshalTargetedEmotionResults unmarshals an instance of TargetedEmotionResults from the specified map of raw messages.
func UnmarshalTargetedEmotionResults(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TargetedEmotionResults)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "emotion", &obj.Emotion, UnmarshalEmotionScores)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TargetedSentimentResults : TargetedSentimentResults struct
type TargetedSentimentResults struct {
	// Targeted text.
	Text *string `json:"text,omitempty"`

	// Sentiment score from -1 (negative) to 1 (positive).
	Score *float64 `json:"score,omitempty"`
}

// UnmarshalTargetedSentimentResults unmarshals an instance of TargetedSentimentResults from the specified map of raw messages.
func UnmarshalTargetedSentimentResults(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TargetedSentimentResults)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
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

// TokenResult : TokenResult struct
type TokenResult struct {
	// The token as it appears in the analyzed text.
	Text *string `json:"text,omitempty"`

	// The part of speech of the token. For more information about the values, see [Universal Dependencies POS
	// tags](https://universaldependencies.org/u/pos/).
	PartOfSpeech *string `json:"part_of_speech,omitempty"`

	// Character offsets indicating the beginning and end of the token in the analyzed text.
	Location []int64 `json:"location,omitempty"`

	// The [lemma](https://wikipedia.org/wiki/Lemma_%28morphology%29) of the token.
	Lemma *string `json:"lemma,omitempty"`
}

// Constants associated with the TokenResult.PartOfSpeech property.
// The part of speech of the token. For more information about the values, see [Universal Dependencies POS
// tags](https://universaldependencies.org/u/pos/).
const (
	TokenResultPartOfSpeechAdjConst   = "ADJ"
	TokenResultPartOfSpeechAdpConst   = "ADP"
	TokenResultPartOfSpeechAdvConst   = "ADV"
	TokenResultPartOfSpeechAuxConst   = "AUX"
	TokenResultPartOfSpeechCconjConst = "CCONJ"
	TokenResultPartOfSpeechDetConst   = "DET"
	TokenResultPartOfSpeechIntjConst  = "INTJ"
	TokenResultPartOfSpeechNounConst  = "NOUN"
	TokenResultPartOfSpeechNumConst   = "NUM"
	TokenResultPartOfSpeechPartConst  = "PART"
	TokenResultPartOfSpeechPronConst  = "PRON"
	TokenResultPartOfSpeechPropnConst = "PROPN"
	TokenResultPartOfSpeechPunctConst = "PUNCT"
	TokenResultPartOfSpeechSconjConst = "SCONJ"
	TokenResultPartOfSpeechSymConst   = "SYM"
	TokenResultPartOfSpeechVerbConst  = "VERB"
	TokenResultPartOfSpeechXConst     = "X"
)

// UnmarshalTokenResult unmarshals an instance of TokenResult from the specified map of raw messages.
func UnmarshalTokenResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TokenResult)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "part_of_speech", &obj.PartOfSpeech)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "location", &obj.Location)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "lemma", &obj.Lemma)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateCategoriesModelOptions : The UpdateCategoriesModel options.
type UpdateCategoriesModelOptions struct {
	// ID of the model.
	ModelID *string `validate:"required,ne="`

	// The 2-letter language code of this model.
	Language *string `validate:"required"`

	// Training data in JSON format. For more information, see [Categories training data
	// requirements](https://cloud.ibm.com/docs/natural-language-understanding?topic=natural-language-understanding-categories##categories-training-data-requirements).
	TrainingData io.ReadCloser `validate:"required"`

	// The content type of trainingData.
	TrainingDataContentType *string

	// An optional name for the model.
	Name *string

	// An optional map of metadata key-value pairs to store with this model.
	UserMetadata map[string]interface{}

	// An optional description of the model.
	Description *string

	// An optional version string.
	ModelVersion *string

	// ID of the Watson Knowledge Studio workspace that deployed this model to Natural Language Understanding.
	WorkspaceID *string

	// The description of the version.
	VersionDescription *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateCategoriesModelOptions : Instantiate UpdateCategoriesModelOptions
func (*NaturalLanguageUnderstandingV1) NewUpdateCategoriesModelOptions(modelID string, language string, trainingData io.ReadCloser) *UpdateCategoriesModelOptions {
	return &UpdateCategoriesModelOptions{
		ModelID:      core.StringPtr(modelID),
		Language:     core.StringPtr(language),
		TrainingData: trainingData,
	}
}

// SetModelID : Allow user to set ModelID
func (options *UpdateCategoriesModelOptions) SetModelID(modelID string) *UpdateCategoriesModelOptions {
	options.ModelID = core.StringPtr(modelID)
	return options
}

// SetLanguage : Allow user to set Language
func (options *UpdateCategoriesModelOptions) SetLanguage(language string) *UpdateCategoriesModelOptions {
	options.Language = core.StringPtr(language)
	return options
}

// SetTrainingData : Allow user to set TrainingData
func (options *UpdateCategoriesModelOptions) SetTrainingData(trainingData io.ReadCloser) *UpdateCategoriesModelOptions {
	options.TrainingData = trainingData
	return options
}

// SetTrainingDataContentType : Allow user to set TrainingDataContentType
func (options *UpdateCategoriesModelOptions) SetTrainingDataContentType(trainingDataContentType string) *UpdateCategoriesModelOptions {
	options.TrainingDataContentType = core.StringPtr(trainingDataContentType)
	return options
}

// SetName : Allow user to set Name
func (options *UpdateCategoriesModelOptions) SetName(name string) *UpdateCategoriesModelOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetUserMetadata : Allow user to set UserMetadata
func (options *UpdateCategoriesModelOptions) SetUserMetadata(userMetadata map[string]interface{}) *UpdateCategoriesModelOptions {
	options.UserMetadata = userMetadata
	return options
}

// SetDescription : Allow user to set Description
func (options *UpdateCategoriesModelOptions) SetDescription(description string) *UpdateCategoriesModelOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetModelVersion : Allow user to set ModelVersion
func (options *UpdateCategoriesModelOptions) SetModelVersion(modelVersion string) *UpdateCategoriesModelOptions {
	options.ModelVersion = core.StringPtr(modelVersion)
	return options
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *UpdateCategoriesModelOptions) SetWorkspaceID(workspaceID string) *UpdateCategoriesModelOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetVersionDescription : Allow user to set VersionDescription
func (options *UpdateCategoriesModelOptions) SetVersionDescription(versionDescription string) *UpdateCategoriesModelOptions {
	options.VersionDescription = core.StringPtr(versionDescription)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateCategoriesModelOptions) SetHeaders(param map[string]string) *UpdateCategoriesModelOptions {
	options.Headers = param
	return options
}

// UpdateClassificationsModelOptions : The UpdateClassificationsModel options.
type UpdateClassificationsModelOptions struct {
	// ID of the model.
	ModelID *string `validate:"required,ne="`

	// The 2-letter language code of this model.
	Language *string `validate:"required"`

	// Training data in JSON format. For more information, see [Classifications training data
	// requirements](https://cloud.ibm.com/docs/natural-language-understanding?topic=natural-language-understanding-classifications#classification-training-data-requirements).
	TrainingData io.ReadCloser `validate:"required"`

	// The content type of trainingData.
	TrainingDataContentType *string

	// An optional name for the model.
	Name *string

	// An optional map of metadata key-value pairs to store with this model.
	UserMetadata map[string]interface{}

	// An optional description of the model.
	Description *string

	// An optional version string.
	ModelVersion *string

	// ID of the Watson Knowledge Studio workspace that deployed this model to Natural Language Understanding.
	WorkspaceID *string

	// The description of the version.
	VersionDescription *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateClassificationsModelOptions : Instantiate UpdateClassificationsModelOptions
func (*NaturalLanguageUnderstandingV1) NewUpdateClassificationsModelOptions(modelID string, language string, trainingData io.ReadCloser) *UpdateClassificationsModelOptions {
	return &UpdateClassificationsModelOptions{
		ModelID:      core.StringPtr(modelID),
		Language:     core.StringPtr(language),
		TrainingData: trainingData,
	}
}

// SetModelID : Allow user to set ModelID
func (options *UpdateClassificationsModelOptions) SetModelID(modelID string) *UpdateClassificationsModelOptions {
	options.ModelID = core.StringPtr(modelID)
	return options
}

// SetLanguage : Allow user to set Language
func (options *UpdateClassificationsModelOptions) SetLanguage(language string) *UpdateClassificationsModelOptions {
	options.Language = core.StringPtr(language)
	return options
}

// SetTrainingData : Allow user to set TrainingData
func (options *UpdateClassificationsModelOptions) SetTrainingData(trainingData io.ReadCloser) *UpdateClassificationsModelOptions {
	options.TrainingData = trainingData
	return options
}

// SetTrainingDataContentType : Allow user to set TrainingDataContentType
func (options *UpdateClassificationsModelOptions) SetTrainingDataContentType(trainingDataContentType string) *UpdateClassificationsModelOptions {
	options.TrainingDataContentType = core.StringPtr(trainingDataContentType)
	return options
}

// SetName : Allow user to set Name
func (options *UpdateClassificationsModelOptions) SetName(name string) *UpdateClassificationsModelOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetUserMetadata : Allow user to set UserMetadata
func (options *UpdateClassificationsModelOptions) SetUserMetadata(userMetadata map[string]interface{}) *UpdateClassificationsModelOptions {
	options.UserMetadata = userMetadata
	return options
}

// SetDescription : Allow user to set Description
func (options *UpdateClassificationsModelOptions) SetDescription(description string) *UpdateClassificationsModelOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetModelVersion : Allow user to set ModelVersion
func (options *UpdateClassificationsModelOptions) SetModelVersion(modelVersion string) *UpdateClassificationsModelOptions {
	options.ModelVersion = core.StringPtr(modelVersion)
	return options
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *UpdateClassificationsModelOptions) SetWorkspaceID(workspaceID string) *UpdateClassificationsModelOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetVersionDescription : Allow user to set VersionDescription
func (options *UpdateClassificationsModelOptions) SetVersionDescription(versionDescription string) *UpdateClassificationsModelOptions {
	options.VersionDescription = core.StringPtr(versionDescription)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateClassificationsModelOptions) SetHeaders(param map[string]string) *UpdateClassificationsModelOptions {
	options.Headers = param
	return options
}

// UpdateSentimentModelOptions : The UpdateSentimentModel options.
type UpdateSentimentModelOptions struct {
	// ID of the model.
	ModelID *string `validate:"required,ne="`

	// The 2-letter language code of this model.
	Language *string `validate:"required"`

	// Training data in CSV format. For more information, see [Sentiment training data
	// requirements](https://cloud.ibm.com/docs/natural-language-understanding?topic=natural-language-understanding-custom-sentiment#sentiment-training-data-requirements).
	TrainingData io.ReadCloser `validate:"required"`

	// An optional name for the model.
	Name *string

	// An optional map of metadata key-value pairs to store with this model.
	UserMetadata map[string]interface{}

	// An optional description of the model.
	Description *string

	// An optional version string.
	ModelVersion *string

	// ID of the Watson Knowledge Studio workspace that deployed this model to Natural Language Understanding.
	WorkspaceID *string

	// The description of the version.
	VersionDescription *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateSentimentModelOptions : Instantiate UpdateSentimentModelOptions
func (*NaturalLanguageUnderstandingV1) NewUpdateSentimentModelOptions(modelID string, language string, trainingData io.ReadCloser) *UpdateSentimentModelOptions {
	return &UpdateSentimentModelOptions{
		ModelID:      core.StringPtr(modelID),
		Language:     core.StringPtr(language),
		TrainingData: trainingData,
	}
}

// SetModelID : Allow user to set ModelID
func (options *UpdateSentimentModelOptions) SetModelID(modelID string) *UpdateSentimentModelOptions {
	options.ModelID = core.StringPtr(modelID)
	return options
}

// SetLanguage : Allow user to set Language
func (options *UpdateSentimentModelOptions) SetLanguage(language string) *UpdateSentimentModelOptions {
	options.Language = core.StringPtr(language)
	return options
}

// SetTrainingData : Allow user to set TrainingData
func (options *UpdateSentimentModelOptions) SetTrainingData(trainingData io.ReadCloser) *UpdateSentimentModelOptions {
	options.TrainingData = trainingData
	return options
}

// SetName : Allow user to set Name
func (options *UpdateSentimentModelOptions) SetName(name string) *UpdateSentimentModelOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetUserMetadata : Allow user to set UserMetadata
func (options *UpdateSentimentModelOptions) SetUserMetadata(userMetadata map[string]interface{}) *UpdateSentimentModelOptions {
	options.UserMetadata = userMetadata
	return options
}

// SetDescription : Allow user to set Description
func (options *UpdateSentimentModelOptions) SetDescription(description string) *UpdateSentimentModelOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetModelVersion : Allow user to set ModelVersion
func (options *UpdateSentimentModelOptions) SetModelVersion(modelVersion string) *UpdateSentimentModelOptions {
	options.ModelVersion = core.StringPtr(modelVersion)
	return options
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *UpdateSentimentModelOptions) SetWorkspaceID(workspaceID string) *UpdateSentimentModelOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetVersionDescription : Allow user to set VersionDescription
func (options *UpdateSentimentModelOptions) SetVersionDescription(versionDescription string) *UpdateSentimentModelOptions {
	options.VersionDescription = core.StringPtr(versionDescription)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateSentimentModelOptions) SetHeaders(param map[string]string) *UpdateSentimentModelOptions {
	options.Headers = param
	return options
}
