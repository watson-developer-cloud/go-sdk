/**
 * (C) Copyright IBM Corp. 2018, 2021.
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

// Package comparecomplyv1 : Operations and models for the CompareComplyV1 service
package comparecomplyv1

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

// CompareComplyV1 : IBM Watson&trade; Compare and Comply is discontinued. Existing instances are supported until 30
// November 2021, but as of 1 December 2020, you can't create instances. Any instance that exists on 30 November 2021
// will be deleted. Consider migrating to Watson Discovery Premium on IBM Cloud for your Compare and Comply use cases.
// To start the migration process, visit [https://ibm.biz/contact-wdc-premium](https://ibm.biz/contact-wdc-premium).
// {: deprecated}
//
// Compare and Comply analyzes governing documents to provide details about critical aspects of the documents.
//
// Version: 1.0
// See: https://cloud.ibm.com/docs/compare-comply?topic=compare-comply-about
type CompareComplyV1 struct {
	Service *core.BaseService

	// Release date of the version of the API you want to use. Specify dates in YYYY-MM-DD format. The current version is
	// `2018-10-15`.
	Version *string
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.us-south.compare-comply.watson.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "compare-comply"

// CompareComplyV1Options : Service options
type CompareComplyV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator

	// Release date of the version of the API you want to use. Specify dates in YYYY-MM-DD format. The current version is
	// `2018-10-15`.
	Version *string `validate:"required"`
}

// NewCompareComplyV1 : constructs an instance of CompareComplyV1 with passed in options.
func NewCompareComplyV1(options *CompareComplyV1Options) (service *CompareComplyV1, err error) {
	// Log deprecation warning
	core.GetLogger().Log(core.LevelWarn, "", "On 30 November 2021, Compare and Comply will no longer be available. For more information, see Compare and Comply deprecation.")

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

	service = &CompareComplyV1{
		Service: baseService,
		Version: options.Version,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "compareComply" suitable for processing requests.
func (compareComply *CompareComplyV1) Clone() *CompareComplyV1 {
	if core.IsNil(compareComply) {
		return nil
	}
	clone := *compareComply
	clone.Service = compareComply.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (compareComply *CompareComplyV1) SetServiceURL(url string) error {
	return compareComply.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (compareComply *CompareComplyV1) GetServiceURL() string {
	return compareComply.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (compareComply *CompareComplyV1) SetDefaultHeaders(headers http.Header) {
	compareComply.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (compareComply *CompareComplyV1) SetEnableGzipCompression(enableGzip bool) {
	compareComply.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (compareComply *CompareComplyV1) GetEnableGzipCompression() bool {
	return compareComply.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (compareComply *CompareComplyV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	compareComply.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (compareComply *CompareComplyV1) DisableRetries() {
	compareComply.Service.DisableRetries()
}

// DisableSSLVerification bypasses verification of the server's SSL certificate
func (compareComply *CompareComplyV1) DisableSSLVerification() {
	compareComply.Service.DisableSSLVerification()
}

// ConvertToHTML : Convert document to HTML
// Converts a document to HTML.
func (compareComply *CompareComplyV1) ConvertToHTML(convertToHTMLOptions *ConvertToHTMLOptions) (result *HTMLReturn, response *core.DetailedResponse, err error) {
	return compareComply.ConvertToHTMLWithContext(context.Background(), convertToHTMLOptions)
}

// ConvertToHTMLWithContext is an alternate form of the ConvertToHTML method which supports a Context parameter
func (compareComply *CompareComplyV1) ConvertToHTMLWithContext(ctx context.Context, convertToHTMLOptions *ConvertToHTMLOptions) (result *HTMLReturn, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(convertToHTMLOptions, "convertToHTMLOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(convertToHTMLOptions, "convertToHTMLOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = compareComply.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(compareComply.Service.Options.URL, `/v1/html_conversion`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range convertToHTMLOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compare-comply", "V1", "ConvertToHTML")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*compareComply.Version))
	if convertToHTMLOptions.Model != nil {
		builder.AddQuery("model", fmt.Sprint(*convertToHTMLOptions.Model))
	}

	builder.AddFormData("file", "filename",
		core.StringNilMapper(convertToHTMLOptions.FileContentType), convertToHTMLOptions.File)

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = compareComply.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalHTMLReturn)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ClassifyElements : Classify the elements of a document
// Analyzes the structural and semantic elements of a document.
func (compareComply *CompareComplyV1) ClassifyElements(classifyElementsOptions *ClassifyElementsOptions) (result *ClassifyReturn, response *core.DetailedResponse, err error) {
	return compareComply.ClassifyElementsWithContext(context.Background(), classifyElementsOptions)
}

// ClassifyElementsWithContext is an alternate form of the ClassifyElements method which supports a Context parameter
func (compareComply *CompareComplyV1) ClassifyElementsWithContext(ctx context.Context, classifyElementsOptions *ClassifyElementsOptions) (result *ClassifyReturn, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(classifyElementsOptions, "classifyElementsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(classifyElementsOptions, "classifyElementsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = compareComply.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(compareComply.Service.Options.URL, `/v1/element_classification`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range classifyElementsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compare-comply", "V1", "ClassifyElements")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*compareComply.Version))
	if classifyElementsOptions.Model != nil {
		builder.AddQuery("model", fmt.Sprint(*classifyElementsOptions.Model))
	}

	builder.AddFormData("file", "filename",
		core.StringNilMapper(classifyElementsOptions.FileContentType), classifyElementsOptions.File)

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = compareComply.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalClassifyReturn)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ExtractTables : Extract a document's tables
// Analyzes the tables in a document.
func (compareComply *CompareComplyV1) ExtractTables(extractTablesOptions *ExtractTablesOptions) (result *TableReturn, response *core.DetailedResponse, err error) {
	return compareComply.ExtractTablesWithContext(context.Background(), extractTablesOptions)
}

// ExtractTablesWithContext is an alternate form of the ExtractTables method which supports a Context parameter
func (compareComply *CompareComplyV1) ExtractTablesWithContext(ctx context.Context, extractTablesOptions *ExtractTablesOptions) (result *TableReturn, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(extractTablesOptions, "extractTablesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(extractTablesOptions, "extractTablesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = compareComply.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(compareComply.Service.Options.URL, `/v1/tables`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range extractTablesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compare-comply", "V1", "ExtractTables")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*compareComply.Version))
	if extractTablesOptions.Model != nil {
		builder.AddQuery("model", fmt.Sprint(*extractTablesOptions.Model))
	}

	builder.AddFormData("file", "filename",
		core.StringNilMapper(extractTablesOptions.FileContentType), extractTablesOptions.File)

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = compareComply.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTableReturn)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CompareDocuments : Compare two documents
// Compares two input documents. Documents must be in the same format.
func (compareComply *CompareComplyV1) CompareDocuments(compareDocumentsOptions *CompareDocumentsOptions) (result *CompareReturn, response *core.DetailedResponse, err error) {
	return compareComply.CompareDocumentsWithContext(context.Background(), compareDocumentsOptions)
}

// CompareDocumentsWithContext is an alternate form of the CompareDocuments method which supports a Context parameter
func (compareComply *CompareComplyV1) CompareDocumentsWithContext(ctx context.Context, compareDocumentsOptions *CompareDocumentsOptions) (result *CompareReturn, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(compareDocumentsOptions, "compareDocumentsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(compareDocumentsOptions, "compareDocumentsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = compareComply.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(compareComply.Service.Options.URL, `/v1/comparison`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range compareDocumentsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compare-comply", "V1", "CompareDocuments")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*compareComply.Version))
	if compareDocumentsOptions.File1Label != nil {
		builder.AddQuery("file_1_label", fmt.Sprint(*compareDocumentsOptions.File1Label))
	}
	if compareDocumentsOptions.File2Label != nil {
		builder.AddQuery("file_2_label", fmt.Sprint(*compareDocumentsOptions.File2Label))
	}
	if compareDocumentsOptions.Model != nil {
		builder.AddQuery("model", fmt.Sprint(*compareDocumentsOptions.Model))
	}

	builder.AddFormData("file_1", "filename",
		core.StringNilMapper(compareDocumentsOptions.File1ContentType), compareDocumentsOptions.File1)
	builder.AddFormData("file_2", "filename",
		core.StringNilMapper(compareDocumentsOptions.File2ContentType), compareDocumentsOptions.File2)

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = compareComply.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCompareReturn)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// AddFeedback : Add feedback
// Adds feedback in the form of _labels_ from a subject-matter expert (SME) to a governing document.
// **Important:** Feedback is not immediately incorporated into the training model, nor is it guaranteed to be
// incorporated at a later date. Instead, submitted feedback is used to suggest future updates to the training model.
func (compareComply *CompareComplyV1) AddFeedback(addFeedbackOptions *AddFeedbackOptions) (result *FeedbackReturn, response *core.DetailedResponse, err error) {
	return compareComply.AddFeedbackWithContext(context.Background(), addFeedbackOptions)
}

// AddFeedbackWithContext is an alternate form of the AddFeedback method which supports a Context parameter
func (compareComply *CompareComplyV1) AddFeedbackWithContext(ctx context.Context, addFeedbackOptions *AddFeedbackOptions) (result *FeedbackReturn, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addFeedbackOptions, "addFeedbackOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addFeedbackOptions, "addFeedbackOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = compareComply.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(compareComply.Service.Options.URL, `/v1/feedback`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range addFeedbackOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compare-comply", "V1", "AddFeedback")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*compareComply.Version))

	body := make(map[string]interface{})
	if addFeedbackOptions.FeedbackData != nil {
		body["feedback_data"] = addFeedbackOptions.FeedbackData
	}
	if addFeedbackOptions.UserID != nil {
		body["user_id"] = addFeedbackOptions.UserID
	}
	if addFeedbackOptions.Comment != nil {
		body["comment"] = addFeedbackOptions.Comment
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
	response, err = compareComply.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalFeedbackReturn)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListFeedback : List the feedback in a document
// Lists the feedback in a document.
func (compareComply *CompareComplyV1) ListFeedback(listFeedbackOptions *ListFeedbackOptions) (result *FeedbackList, response *core.DetailedResponse, err error) {
	return compareComply.ListFeedbackWithContext(context.Background(), listFeedbackOptions)
}

// ListFeedbackWithContext is an alternate form of the ListFeedback method which supports a Context parameter
func (compareComply *CompareComplyV1) ListFeedbackWithContext(ctx context.Context, listFeedbackOptions *ListFeedbackOptions) (result *FeedbackList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listFeedbackOptions, "listFeedbackOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = compareComply.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(compareComply.Service.Options.URL, `/v1/feedback`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listFeedbackOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compare-comply", "V1", "ListFeedback")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*compareComply.Version))
	if listFeedbackOptions.FeedbackType != nil {
		builder.AddQuery("feedback_type", fmt.Sprint(*listFeedbackOptions.FeedbackType))
	}
	if listFeedbackOptions.DocumentTitle != nil {
		builder.AddQuery("document_title", fmt.Sprint(*listFeedbackOptions.DocumentTitle))
	}
	if listFeedbackOptions.ModelID != nil {
		builder.AddQuery("model_id", fmt.Sprint(*listFeedbackOptions.ModelID))
	}
	if listFeedbackOptions.ModelVersion != nil {
		builder.AddQuery("model_version", fmt.Sprint(*listFeedbackOptions.ModelVersion))
	}
	if listFeedbackOptions.CategoryRemoved != nil {
		builder.AddQuery("category_removed", fmt.Sprint(*listFeedbackOptions.CategoryRemoved))
	}
	if listFeedbackOptions.CategoryAdded != nil {
		builder.AddQuery("category_added", fmt.Sprint(*listFeedbackOptions.CategoryAdded))
	}
	if listFeedbackOptions.CategoryNotChanged != nil {
		builder.AddQuery("category_not_changed", fmt.Sprint(*listFeedbackOptions.CategoryNotChanged))
	}
	if listFeedbackOptions.TypeRemoved != nil {
		builder.AddQuery("type_removed", fmt.Sprint(*listFeedbackOptions.TypeRemoved))
	}
	if listFeedbackOptions.TypeAdded != nil {
		builder.AddQuery("type_added", fmt.Sprint(*listFeedbackOptions.TypeAdded))
	}
	if listFeedbackOptions.TypeNotChanged != nil {
		builder.AddQuery("type_not_changed", fmt.Sprint(*listFeedbackOptions.TypeNotChanged))
	}
	if listFeedbackOptions.PageLimit != nil {
		builder.AddQuery("page_limit", fmt.Sprint(*listFeedbackOptions.PageLimit))
	}
	if listFeedbackOptions.Cursor != nil {
		builder.AddQuery("cursor", fmt.Sprint(*listFeedbackOptions.Cursor))
	}
	if listFeedbackOptions.Sort != nil {
		builder.AddQuery("sort", fmt.Sprint(*listFeedbackOptions.Sort))
	}
	if listFeedbackOptions.IncludeTotal != nil {
		builder.AddQuery("include_total", fmt.Sprint(*listFeedbackOptions.IncludeTotal))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = compareComply.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalFeedbackList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetFeedback : Get a specified feedback entry
// Gets a feedback entry with a specified `feedback_id`.
func (compareComply *CompareComplyV1) GetFeedback(getFeedbackOptions *GetFeedbackOptions) (result *GetFeedback, response *core.DetailedResponse, err error) {
	return compareComply.GetFeedbackWithContext(context.Background(), getFeedbackOptions)
}

// GetFeedbackWithContext is an alternate form of the GetFeedback method which supports a Context parameter
func (compareComply *CompareComplyV1) GetFeedbackWithContext(ctx context.Context, getFeedbackOptions *GetFeedbackOptions) (result *GetFeedback, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getFeedbackOptions, "getFeedbackOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getFeedbackOptions, "getFeedbackOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"feedback_id": *getFeedbackOptions.FeedbackID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = compareComply.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(compareComply.Service.Options.URL, `/v1/feedback/{feedback_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getFeedbackOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compare-comply", "V1", "GetFeedback")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*compareComply.Version))
	if getFeedbackOptions.Model != nil {
		builder.AddQuery("model", fmt.Sprint(*getFeedbackOptions.Model))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = compareComply.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetFeedback)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteFeedback : Delete a specified feedback entry
// Deletes a feedback entry with a specified `feedback_id`.
func (compareComply *CompareComplyV1) DeleteFeedback(deleteFeedbackOptions *DeleteFeedbackOptions) (result *FeedbackDeleted, response *core.DetailedResponse, err error) {
	return compareComply.DeleteFeedbackWithContext(context.Background(), deleteFeedbackOptions)
}

// DeleteFeedbackWithContext is an alternate form of the DeleteFeedback method which supports a Context parameter
func (compareComply *CompareComplyV1) DeleteFeedbackWithContext(ctx context.Context, deleteFeedbackOptions *DeleteFeedbackOptions) (result *FeedbackDeleted, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteFeedbackOptions, "deleteFeedbackOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteFeedbackOptions, "deleteFeedbackOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"feedback_id": *deleteFeedbackOptions.FeedbackID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = compareComply.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(compareComply.Service.Options.URL, `/v1/feedback/{feedback_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteFeedbackOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compare-comply", "V1", "DeleteFeedback")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*compareComply.Version))
	if deleteFeedbackOptions.Model != nil {
		builder.AddQuery("model", fmt.Sprint(*deleteFeedbackOptions.Model))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = compareComply.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalFeedbackDeleted)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateBatch : Submit a batch-processing request
// Run Compare and Comply methods over a collection of input documents.
//
// **Important:** Batch processing requires the use of the [IBM Cloud Object Storage
// service](https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-about#about-ibm-cloud-object-storage).
// The use of IBM Cloud Object Storage with Compare and Comply is discussed at [Using batch
// processing](https://cloud.ibm.com/docs/compare-comply?topic=compare-comply-batching#before-you-batch).
func (compareComply *CompareComplyV1) CreateBatch(createBatchOptions *CreateBatchOptions) (result *BatchStatus, response *core.DetailedResponse, err error) {
	return compareComply.CreateBatchWithContext(context.Background(), createBatchOptions)
}

// CreateBatchWithContext is an alternate form of the CreateBatch method which supports a Context parameter
func (compareComply *CompareComplyV1) CreateBatchWithContext(ctx context.Context, createBatchOptions *CreateBatchOptions) (result *BatchStatus, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createBatchOptions, "createBatchOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createBatchOptions, "createBatchOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = compareComply.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(compareComply.Service.Options.URL, `/v1/batches`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createBatchOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compare-comply", "V1", "CreateBatch")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*compareComply.Version))
	builder.AddQuery("function", fmt.Sprint(*createBatchOptions.Function))
	if createBatchOptions.Model != nil {
		builder.AddQuery("model", fmt.Sprint(*createBatchOptions.Model))
	}

	builder.AddFormData("input_credentials_file", "filename",
		"application/json", createBatchOptions.InputCredentialsFile)
	builder.AddFormData("input_bucket_location", "", "", fmt.Sprint(*createBatchOptions.InputBucketLocation))
	builder.AddFormData("input_bucket_name", "", "", fmt.Sprint(*createBatchOptions.InputBucketName))
	builder.AddFormData("output_credentials_file", "filename",
		"application/json", createBatchOptions.OutputCredentialsFile)
	builder.AddFormData("output_bucket_location", "", "", fmt.Sprint(*createBatchOptions.OutputBucketLocation))
	builder.AddFormData("output_bucket_name", "", "", fmt.Sprint(*createBatchOptions.OutputBucketName))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = compareComply.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBatchStatus)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListBatches : List submitted batch-processing jobs
// Lists batch-processing jobs submitted by users.
func (compareComply *CompareComplyV1) ListBatches(listBatchesOptions *ListBatchesOptions) (result *Batches, response *core.DetailedResponse, err error) {
	return compareComply.ListBatchesWithContext(context.Background(), listBatchesOptions)
}

// ListBatchesWithContext is an alternate form of the ListBatches method which supports a Context parameter
func (compareComply *CompareComplyV1) ListBatchesWithContext(ctx context.Context, listBatchesOptions *ListBatchesOptions) (result *Batches, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listBatchesOptions, "listBatchesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = compareComply.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(compareComply.Service.Options.URL, `/v1/batches`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listBatchesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compare-comply", "V1", "ListBatches")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*compareComply.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = compareComply.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBatches)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetBatch : Get information about a specific batch-processing job
// Gets information about a batch-processing job with a specified ID.
func (compareComply *CompareComplyV1) GetBatch(getBatchOptions *GetBatchOptions) (result *BatchStatus, response *core.DetailedResponse, err error) {
	return compareComply.GetBatchWithContext(context.Background(), getBatchOptions)
}

// GetBatchWithContext is an alternate form of the GetBatch method which supports a Context parameter
func (compareComply *CompareComplyV1) GetBatchWithContext(ctx context.Context, getBatchOptions *GetBatchOptions) (result *BatchStatus, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getBatchOptions, "getBatchOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getBatchOptions, "getBatchOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"batch_id": *getBatchOptions.BatchID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = compareComply.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(compareComply.Service.Options.URL, `/v1/batches/{batch_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getBatchOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compare-comply", "V1", "GetBatch")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*compareComply.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = compareComply.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBatchStatus)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateBatch : Update a pending or active batch-processing job
// Updates a pending or active batch-processing job. You can rescan the input bucket to check for new documents or
// cancel a job.
func (compareComply *CompareComplyV1) UpdateBatch(updateBatchOptions *UpdateBatchOptions) (result *BatchStatus, response *core.DetailedResponse, err error) {
	return compareComply.UpdateBatchWithContext(context.Background(), updateBatchOptions)
}

// UpdateBatchWithContext is an alternate form of the UpdateBatch method which supports a Context parameter
func (compareComply *CompareComplyV1) UpdateBatchWithContext(ctx context.Context, updateBatchOptions *UpdateBatchOptions) (result *BatchStatus, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateBatchOptions, "updateBatchOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateBatchOptions, "updateBatchOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"batch_id": *updateBatchOptions.BatchID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = compareComply.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(compareComply.Service.Options.URL, `/v1/batches/{batch_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateBatchOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compare-comply", "V1", "UpdateBatch")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*compareComply.Version))
	builder.AddQuery("action", fmt.Sprint(*updateBatchOptions.Action))
	if updateBatchOptions.Model != nil {
		builder.AddQuery("model", fmt.Sprint(*updateBatchOptions.Model))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = compareComply.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBatchStatus)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// AddFeedbackOptions : The AddFeedback options.
type AddFeedbackOptions struct {
	// Feedback data for submission.
	FeedbackData *FeedbackDataInput `validate:"required"`

	// An optional string identifying the user.
	UserID *string

	// An optional comment on or description of the feedback.
	Comment *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAddFeedbackOptions : Instantiate AddFeedbackOptions
func (*CompareComplyV1) NewAddFeedbackOptions(feedbackData *FeedbackDataInput) *AddFeedbackOptions {
	return &AddFeedbackOptions{
		FeedbackData: feedbackData,
	}
}

// SetFeedbackData : Allow user to set FeedbackData
func (options *AddFeedbackOptions) SetFeedbackData(feedbackData *FeedbackDataInput) *AddFeedbackOptions {
	options.FeedbackData = feedbackData
	return options
}

// SetUserID : Allow user to set UserID
func (options *AddFeedbackOptions) SetUserID(userID string) *AddFeedbackOptions {
	options.UserID = core.StringPtr(userID)
	return options
}

// SetComment : Allow user to set Comment
func (options *AddFeedbackOptions) SetComment(comment string) *AddFeedbackOptions {
	options.Comment = core.StringPtr(comment)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *AddFeedbackOptions) SetHeaders(param map[string]string) *AddFeedbackOptions {
	options.Headers = param
	return options
}

// Address : A party's address.
type Address struct {
	// A string listing the address.
	Text *string `json:"text,omitempty"`

	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *Location `json:"location,omitempty"`
}

// UnmarshalAddress unmarshals an instance of Address from the specified map of raw messages.
func UnmarshalAddress(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Address)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
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

// AlignedElement : AlignedElement struct
type AlignedElement struct {
	// Identifies two elements that semantically align between the compared documents.
	ElementPair []ElementPair `json:"element_pair,omitempty"`

	// Specifies whether the aligned element is identical. Elements are considered identical despite minor differences such
	// as leading punctuation, end-of-sentence punctuation, whitespace, the presence or absence of definite or indefinite
	// articles, and others.
	IdenticalText *bool `json:"identical_text,omitempty"`

	// Hashed values that you can send to IBM to provide feedback or receive support.
	ProvenanceIds []string `json:"provenance_ids,omitempty"`

	// Indicates that the elements aligned are contractual clauses of significance.
	SignificantElements *bool `json:"significant_elements,omitempty"`
}

// UnmarshalAlignedElement unmarshals an instance of AlignedElement from the specified map of raw messages.
func UnmarshalAlignedElement(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AlignedElement)
	err = core.UnmarshalModel(m, "element_pair", &obj.ElementPair, UnmarshalElementPair)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "identical_text", &obj.IdenticalText)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "provenance_ids", &obj.ProvenanceIds)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "significant_elements", &obj.SignificantElements)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Attribute : List of document attributes.
type Attribute struct {
	// The type of attribute.
	Type *string `json:"type,omitempty"`

	// The text associated with the attribute.
	Text *string `json:"text,omitempty"`

	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *Location `json:"location,omitempty"`
}

// Constants associated with the Attribute.Type property.
// The type of attribute.
const (
	AttributeTypeCurrencyConst     = "Currency"
	AttributeTypeDatetimeConst     = "DateTime"
	AttributeTypeDefinedtermConst  = "DefinedTerm"
	AttributeTypeDurationConst     = "Duration"
	AttributeTypeLocationConst     = "Location"
	AttributeTypeNumberConst       = "Number"
	AttributeTypeOrganizationConst = "Organization"
	AttributeTypePercentageConst   = "Percentage"
	AttributeTypePersonConst       = "Person"
)

// UnmarshalAttribute unmarshals an instance of Attribute from the specified map of raw messages.
func UnmarshalAttribute(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Attribute)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
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

// BatchStatus : The batch-request status.
type BatchStatus struct {
	// The method to be run against the documents. Possible values are `html_conversion`, `element_classification`, and
	// `tables`.
	Function *string `json:"function,omitempty"`

	// The geographical location of the Cloud Object Storage input bucket as listed on the **Endpoint** tab of your COS
	// instance; for example, `us-geo`, `eu-geo`, or `ap-geo`.
	InputBucketLocation *string `json:"input_bucket_location,omitempty"`

	// The name of the Cloud Object Storage input bucket.
	InputBucketName *string `json:"input_bucket_name,omitempty"`

	// The geographical location of the Cloud Object Storage output bucket as listed on the **Endpoint** tab of your COS
	// instance; for example, `us-geo`, `eu-geo`, or `ap-geo`.
	OutputBucketLocation *string `json:"output_bucket_location,omitempty"`

	// The name of the Cloud Object Storage output bucket.
	OutputBucketName *string `json:"output_bucket_name,omitempty"`

	// The unique identifier for the batch request.
	BatchID *string `json:"batch_id,omitempty"`

	// Document counts.
	DocumentCounts *DocCounts `json:"document_counts,omitempty"`

	// The status of the batch request.
	Status *string `json:"status,omitempty"`

	// The creation time of the batch request.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// The time of the most recent update to the batch request.
	Updated *strfmt.DateTime `json:"updated,omitempty"`
}

// Constants associated with the BatchStatus.Function property.
// The method to be run against the documents. Possible values are `html_conversion`, `element_classification`, and
// `tables`.
const (
	BatchStatusFunctionElementClassificationConst = "element_classification"
	BatchStatusFunctionHTMLConversionConst        = "html_conversion"
	BatchStatusFunctionTablesConst                = "tables"
)

// UnmarshalBatchStatus unmarshals an instance of BatchStatus from the specified map of raw messages.
func UnmarshalBatchStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BatchStatus)
	err = core.UnmarshalPrimitive(m, "function", &obj.Function)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "input_bucket_location", &obj.InputBucketLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "input_bucket_name", &obj.InputBucketName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "output_bucket_location", &obj.OutputBucketLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "output_bucket_name", &obj.OutputBucketName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "batch_id", &obj.BatchID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "document_counts", &obj.DocumentCounts, UnmarshalDocCounts)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Batches : The results of a successful **List Batches** request.
type Batches struct {
	// A list of the status of all batch requests.
	Batches []BatchStatus `json:"batches,omitempty"`
}

// UnmarshalBatches unmarshals an instance of Batches from the specified map of raw messages.
func UnmarshalBatches(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Batches)
	err = core.UnmarshalModel(m, "batches", &obj.Batches, UnmarshalBatchStatus)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BodyCells : Cells that are not table header, column header, or row header cells.
type BodyCells struct {
	// The unique ID of the cell in the current table.
	CellID *string `json:"cell_id,omitempty"`

	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *Location `json:"location,omitempty"`

	// The textual contents of this cell from the input document without associated markup content.
	Text *string `json:"text,omitempty"`

	// The `begin` index of this cell's `row` location in the current table.
	RowIndexBegin *int64 `json:"row_index_begin,omitempty"`

	// The `end` index of this cell's `row` location in the current table.
	RowIndexEnd *int64 `json:"row_index_end,omitempty"`

	// The `begin` index of this cell's `column` location in the current table.
	ColumnIndexBegin *int64 `json:"column_index_begin,omitempty"`

	// The `end` index of this cell's `column` location in the current table.
	ColumnIndexEnd *int64 `json:"column_index_end,omitempty"`

	// An array that contains the `id` value of a row header that is applicable to this body cell.
	RowHeaderIds []string `json:"row_header_ids,omitempty"`

	// An array that contains the `text` value of a row header that is applicable to this body cell.
	RowHeaderTexts []string `json:"row_header_texts,omitempty"`

	// If you provide customization input, the normalized version of the row header texts according to the customization;
	// otherwise, the same value as `row_header_texts`.
	RowHeaderTextsNormalized []string `json:"row_header_texts_normalized,omitempty"`

	// An array that contains the `id` value of a column header that is applicable to the current cell.
	ColumnHeaderIds []string `json:"column_header_ids,omitempty"`

	// An array that contains the `text` value of a column header that is applicable to the current cell.
	ColumnHeaderTexts []string `json:"column_header_texts,omitempty"`

	// If you provide customization input, the normalized version of the column header texts according to the
	// customization; otherwise, the same value as `column_header_texts`.
	ColumnHeaderTextsNormalized []string `json:"column_header_texts_normalized,omitempty"`

	Attributes []Attribute `json:"attributes,omitempty"`
}

// UnmarshalBodyCells unmarshals an instance of BodyCells from the specified map of raw messages.
func UnmarshalBodyCells(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BodyCells)
	err = core.UnmarshalPrimitive(m, "cell_id", &obj.CellID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "location", &obj.Location, UnmarshalLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "row_index_begin", &obj.RowIndexBegin)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "row_index_end", &obj.RowIndexEnd)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "column_index_begin", &obj.ColumnIndexBegin)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "column_index_end", &obj.ColumnIndexEnd)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "row_header_ids", &obj.RowHeaderIds)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "row_header_texts", &obj.RowHeaderTexts)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "row_header_texts_normalized", &obj.RowHeaderTextsNormalized)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "column_header_ids", &obj.ColumnHeaderIds)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "column_header_texts", &obj.ColumnHeaderTexts)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "column_header_texts_normalized", &obj.ColumnHeaderTextsNormalized)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "attributes", &obj.Attributes, UnmarshalAttribute)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Category : Information defining an element's subject matter.
type Category struct {
	// The category of the associated element.
	Label *string `json:"label,omitempty"`

	// Hashed values that you can send to IBM to provide feedback or receive support.
	ProvenanceIds []string `json:"provenance_ids,omitempty"`

	// The type of modification of the feedback entry in the updated labels response.
	Modification *string `json:"modification,omitempty"`
}

// Constants associated with the Category.Label property.
// The category of the associated element.
const (
	CategoryLabelAmendmentsConst           = "Amendments"
	CategoryLabelAssetUseConst             = "Asset Use"
	CategoryLabelAssignmentsConst          = "Assignments"
	CategoryLabelAuditsConst               = "Audits"
	CategoryLabelBusinessContinuityConst   = "Business Continuity"
	CategoryLabelCommunicationConst        = "Communication"
	CategoryLabelConfidentialityConst      = "Confidentiality"
	CategoryLabelDeliverablesConst         = "Deliverables"
	CategoryLabelDeliveryConst             = "Delivery"
	CategoryLabelDisputeResolutionConst    = "Dispute Resolution"
	CategoryLabelForceMajeureConst         = "Force Majeure"
	CategoryLabelIndemnificationConst      = "Indemnification"
	CategoryLabelInsuranceConst            = "Insurance"
	CategoryLabelIntellectualPropertyConst = "Intellectual Property"
	CategoryLabelLiabilityConst            = "Liability"
	CategoryLabelOrderOfPrecedenceConst    = "Order of Precedence"
	CategoryLabelPaymentTermsBillingConst  = "Payment Terms & Billing"
	CategoryLabelPricingTaxesConst         = "Pricing & Taxes"
	CategoryLabelPrivacyConst              = "Privacy"
	CategoryLabelResponsibilitiesConst     = "Responsibilities"
	CategoryLabelSafetyAndSecurityConst    = "Safety and Security"
	CategoryLabelScopeOfWorkConst          = "Scope of Work"
	CategoryLabelSubcontractsConst         = "Subcontracts"
	CategoryLabelTermTerminationConst      = "Term & Termination"
	CategoryLabelWarrantiesConst           = "Warranties"
)

// Constants associated with the Category.Modification property.
// The type of modification of the feedback entry in the updated labels response.
const (
	CategoryModificationAddedConst     = "added"
	CategoryModificationRemovedConst   = "removed"
	CategoryModificationUnchangedConst = "unchanged"
)

// UnmarshalCategory unmarshals an instance of Category from the specified map of raw messages.
func UnmarshalCategory(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Category)
	err = core.UnmarshalPrimitive(m, "label", &obj.Label)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "provenance_ids", &obj.ProvenanceIds)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modification", &obj.Modification)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CategoryComparison : Information defining an element's subject matter.
type CategoryComparison struct {
	// The category of the associated element.
	Label *string `json:"label,omitempty"`
}

// Constants associated with the CategoryComparison.Label property.
// The category of the associated element.
const (
	CategoryComparisonLabelAmendmentsConst           = "Amendments"
	CategoryComparisonLabelAssetUseConst             = "Asset Use"
	CategoryComparisonLabelAssignmentsConst          = "Assignments"
	CategoryComparisonLabelAuditsConst               = "Audits"
	CategoryComparisonLabelBusinessContinuityConst   = "Business Continuity"
	CategoryComparisonLabelCommunicationConst        = "Communication"
	CategoryComparisonLabelConfidentialityConst      = "Confidentiality"
	CategoryComparisonLabelDeliverablesConst         = "Deliverables"
	CategoryComparisonLabelDeliveryConst             = "Delivery"
	CategoryComparisonLabelDisputeResolutionConst    = "Dispute Resolution"
	CategoryComparisonLabelForceMajeureConst         = "Force Majeure"
	CategoryComparisonLabelIndemnificationConst      = "Indemnification"
	CategoryComparisonLabelInsuranceConst            = "Insurance"
	CategoryComparisonLabelIntellectualPropertyConst = "Intellectual Property"
	CategoryComparisonLabelLiabilityConst            = "Liability"
	CategoryComparisonLabelOrderOfPrecedenceConst    = "Order of Precedence"
	CategoryComparisonLabelPaymentTermsBillingConst  = "Payment Terms & Billing"
	CategoryComparisonLabelPricingTaxesConst         = "Pricing & Taxes"
	CategoryComparisonLabelPrivacyConst              = "Privacy"
	CategoryComparisonLabelResponsibilitiesConst     = "Responsibilities"
	CategoryComparisonLabelSafetyAndSecurityConst    = "Safety and Security"
	CategoryComparisonLabelScopeOfWorkConst          = "Scope of Work"
	CategoryComparisonLabelSubcontractsConst         = "Subcontracts"
	CategoryComparisonLabelTermTerminationConst      = "Term & Termination"
	CategoryComparisonLabelWarrantiesConst           = "Warranties"
)

// UnmarshalCategoryComparison unmarshals an instance of CategoryComparison from the specified map of raw messages.
func UnmarshalCategoryComparison(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CategoryComparison)
	err = core.UnmarshalPrimitive(m, "label", &obj.Label)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ClassifyElementsOptions : The ClassifyElements options.
type ClassifyElementsOptions struct {
	// The document to classify.
	File io.ReadCloser `validate:"required"`

	// The content type of file.
	FileContentType *string

	// The analysis model to be used by the service. For the **Element classification** and **Compare two documents**
	// methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults
	// apply to the standalone methods as well as to the methods' use in batch-processing requests.
	Model *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ClassifyElementsOptions.Model property.
// The analysis model to be used by the service. For the **Element classification** and **Compare two documents**
// methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults apply
// to the standalone methods as well as to the methods' use in batch-processing requests.
const (
	ClassifyElementsOptionsModelContractsConst = "contracts"
	ClassifyElementsOptionsModelTablesConst    = "tables"
)

// NewClassifyElementsOptions : Instantiate ClassifyElementsOptions
func (*CompareComplyV1) NewClassifyElementsOptions(file io.ReadCloser) *ClassifyElementsOptions {
	return &ClassifyElementsOptions{
		File: file,
	}
}

// SetFile : Allow user to set File
func (options *ClassifyElementsOptions) SetFile(file io.ReadCloser) *ClassifyElementsOptions {
	options.File = file
	return options
}

// SetFileContentType : Allow user to set FileContentType
func (options *ClassifyElementsOptions) SetFileContentType(fileContentType string) *ClassifyElementsOptions {
	options.FileContentType = core.StringPtr(fileContentType)
	return options
}

// SetModel : Allow user to set Model
func (options *ClassifyElementsOptions) SetModel(model string) *ClassifyElementsOptions {
	options.Model = core.StringPtr(model)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ClassifyElementsOptions) SetHeaders(param map[string]string) *ClassifyElementsOptions {
	options.Headers = param
	return options
}

// ClassifyReturn : The analysis of objects returned by the **Element classification** method.
type ClassifyReturn struct {
	// Basic information about the input document.
	Document *Document `json:"document,omitempty"`

	// The analysis model used to classify the input document. For the **Element classification** method, the only valid
	// value is `contracts`.
	ModelID *string `json:"model_id,omitempty"`

	// The version of the analysis model identified by the value of the `model_id` key.
	ModelVersion *string `json:"model_version,omitempty"`

	// Document elements identified by the service.
	Elements []Element `json:"elements,omitempty"`

	// The date or dates on which the document becomes effective.
	EffectiveDates []EffectiveDates `json:"effective_dates,omitempty"`

	// The monetary amounts that identify the total amount of the contract that needs to be paid from one party to another.
	ContractAmounts []ContractAmts `json:"contract_amounts,omitempty"`

	// The dates on which the document is to be terminated.
	TerminationDates []TerminationDates `json:"termination_dates,omitempty"`

	// The contract type as declared in the document.
	ContractTypes []ContractTypes `json:"contract_types,omitempty"`

	// The durations of the contract.
	ContractTerms []ContractTerms `json:"contract_terms,omitempty"`

	// The document's payment durations.
	PaymentTerms []PaymentTerms `json:"payment_terms,omitempty"`

	// The contract currencies as declared in the document.
	ContractCurrencies []ContractCurrencies `json:"contract_currencies,omitempty"`

	// Definition of tables identified in the input document.
	Tables []Tables `json:"tables,omitempty"`

	// The structure of the input document.
	DocumentStructure *DocStructure `json:"document_structure,omitempty"`

	// Definitions of the parties identified in the input document.
	Parties []Parties `json:"parties,omitempty"`
}

// UnmarshalClassifyReturn unmarshals an instance of ClassifyReturn from the specified map of raw messages.
func UnmarshalClassifyReturn(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ClassifyReturn)
	err = core.UnmarshalModel(m, "document", &obj.Document, UnmarshalDocument)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "model_id", &obj.ModelID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "model_version", &obj.ModelVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "elements", &obj.Elements, UnmarshalElement)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "effective_dates", &obj.EffectiveDates, UnmarshalEffectiveDates)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "contract_amounts", &obj.ContractAmounts, UnmarshalContractAmts)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "termination_dates", &obj.TerminationDates, UnmarshalTerminationDates)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "contract_types", &obj.ContractTypes, UnmarshalContractTypes)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "contract_terms", &obj.ContractTerms, UnmarshalContractTerms)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "payment_terms", &obj.PaymentTerms, UnmarshalPaymentTerms)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "contract_currencies", &obj.ContractCurrencies, UnmarshalContractCurrencies)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "tables", &obj.Tables, UnmarshalTables)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "document_structure", &obj.DocumentStructure, UnmarshalDocStructure)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "parties", &obj.Parties, UnmarshalParties)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ColumnHeaders : Column-level cells, each applicable as a header to other cells in the same column as itself, of the current table.
type ColumnHeaders struct {
	// The unique ID of the cell in the current table.
	CellID *string `json:"cell_id,omitempty"`

	// The location of the column header cell in the current table as defined by its `begin` and `end` offsets,
	// respectfully, in the input document.
	Location interface{} `json:"location,omitempty"`

	// The textual contents of this cell from the input document without associated markup content.
	Text *string `json:"text,omitempty"`

	// If you provide customization input, the normalized version of the cell text according to the customization;
	// otherwise, the same value as `text`.
	TextNormalized *string `json:"text_normalized,omitempty"`

	// The `begin` index of this cell's `row` location in the current table.
	RowIndexBegin *int64 `json:"row_index_begin,omitempty"`

	// The `end` index of this cell's `row` location in the current table.
	RowIndexEnd *int64 `json:"row_index_end,omitempty"`

	// The `begin` index of this cell's `column` location in the current table.
	ColumnIndexBegin *int64 `json:"column_index_begin,omitempty"`

	// The `end` index of this cell's `column` location in the current table.
	ColumnIndexEnd *int64 `json:"column_index_end,omitempty"`
}

// UnmarshalColumnHeaders unmarshals an instance of ColumnHeaders from the specified map of raw messages.
func UnmarshalColumnHeaders(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ColumnHeaders)
	err = core.UnmarshalPrimitive(m, "cell_id", &obj.CellID)
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
	err = core.UnmarshalPrimitive(m, "text_normalized", &obj.TextNormalized)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "row_index_begin", &obj.RowIndexBegin)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "row_index_end", &obj.RowIndexEnd)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "column_index_begin", &obj.ColumnIndexBegin)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "column_index_end", &obj.ColumnIndexEnd)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CompareDocumentsOptions : The CompareDocuments options.
type CompareDocumentsOptions struct {
	// The first document to compare.
	File1 io.ReadCloser `validate:"required"`

	// The second document to compare.
	File2 io.ReadCloser `validate:"required"`

	// The content type of file1.
	File1ContentType *string

	// The content type of file2.
	File2ContentType *string

	// A text label for the first document.
	File1Label *string

	// A text label for the second document.
	File2Label *string

	// The analysis model to be used by the service. For the **Element classification** and **Compare two documents**
	// methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults
	// apply to the standalone methods as well as to the methods' use in batch-processing requests.
	Model *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CompareDocumentsOptions.Model property.
// The analysis model to be used by the service. For the **Element classification** and **Compare two documents**
// methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults apply
// to the standalone methods as well as to the methods' use in batch-processing requests.
const (
	CompareDocumentsOptionsModelContractsConst = "contracts"
	CompareDocumentsOptionsModelTablesConst    = "tables"
)

// NewCompareDocumentsOptions : Instantiate CompareDocumentsOptions
func (*CompareComplyV1) NewCompareDocumentsOptions(file1 io.ReadCloser, file2 io.ReadCloser) *CompareDocumentsOptions {
	return &CompareDocumentsOptions{
		File1: file1,
		File2: file2,
	}
}

// SetFile1 : Allow user to set File1
func (options *CompareDocumentsOptions) SetFile1(file1 io.ReadCloser) *CompareDocumentsOptions {
	options.File1 = file1
	return options
}

// SetFile2 : Allow user to set File2
func (options *CompareDocumentsOptions) SetFile2(file2 io.ReadCloser) *CompareDocumentsOptions {
	options.File2 = file2
	return options
}

// SetFile1ContentType : Allow user to set File1ContentType
func (options *CompareDocumentsOptions) SetFile1ContentType(file1ContentType string) *CompareDocumentsOptions {
	options.File1ContentType = core.StringPtr(file1ContentType)
	return options
}

// SetFile2ContentType : Allow user to set File2ContentType
func (options *CompareDocumentsOptions) SetFile2ContentType(file2ContentType string) *CompareDocumentsOptions {
	options.File2ContentType = core.StringPtr(file2ContentType)
	return options
}

// SetFile1Label : Allow user to set File1Label
func (options *CompareDocumentsOptions) SetFile1Label(file1Label string) *CompareDocumentsOptions {
	options.File1Label = core.StringPtr(file1Label)
	return options
}

// SetFile2Label : Allow user to set File2Label
func (options *CompareDocumentsOptions) SetFile2Label(file2Label string) *CompareDocumentsOptions {
	options.File2Label = core.StringPtr(file2Label)
	return options
}

// SetModel : Allow user to set Model
func (options *CompareDocumentsOptions) SetModel(model string) *CompareDocumentsOptions {
	options.Model = core.StringPtr(model)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CompareDocumentsOptions) SetHeaders(param map[string]string) *CompareDocumentsOptions {
	options.Headers = param
	return options
}

// CompareReturn : The comparison of the two submitted documents.
type CompareReturn struct {
	// The analysis model used to compare the input documents. For the **Compare two documents** method, the only valid
	// value is `contracts`.
	ModelID *string `json:"model_id,omitempty"`

	// The version of the analysis model identified by the value of the `model_id` key.
	ModelVersion *string `json:"model_version,omitempty"`

	// Information about the documents being compared.
	Documents []Document `json:"documents,omitempty"`

	// A list of pairs of elements that semantically align between the compared documents.
	AlignedElements []AlignedElement `json:"aligned_elements,omitempty"`

	// A list of elements that do not semantically align between the compared documents.
	UnalignedElements []UnalignedElement `json:"unaligned_elements,omitempty"`
}

// UnmarshalCompareReturn unmarshals an instance of CompareReturn from the specified map of raw messages.
func UnmarshalCompareReturn(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CompareReturn)
	err = core.UnmarshalPrimitive(m, "model_id", &obj.ModelID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "model_version", &obj.ModelVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "documents", &obj.Documents, UnmarshalDocument)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "aligned_elements", &obj.AlignedElements, UnmarshalAlignedElement)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "unaligned_elements", &obj.UnalignedElements, UnmarshalUnalignedElement)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Contact : A contact.
type Contact struct {
	// A string listing the name of the contact.
	Name *string `json:"name,omitempty"`

	// A string listing the role of the contact.
	Role *string `json:"role,omitempty"`
}

// UnmarshalContact unmarshals an instance of Contact from the specified map of raw messages.
func UnmarshalContact(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Contact)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "role", &obj.Role)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Contexts : Text that is related to the contents of the table and that precedes or follows the current table.
type Contexts struct {
	// The related text.
	Text *string `json:"text,omitempty"`

	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *Location `json:"location,omitempty"`
}

// UnmarshalContexts unmarshals an instance of Contexts from the specified map of raw messages.
func UnmarshalContexts(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Contexts)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
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

// ContractAmts : A monetary amount identified in the input document.
type ContractAmts struct {
	// The confidence level in the identification of the contract amount.
	ConfidenceLevel *string `json:"confidence_level,omitempty"`

	// The monetary amount.
	Text *string `json:"text,omitempty"`

	// The normalized form of the amount, which is listed as a string. This element is optional; it is returned only if
	// normalized text exists.
	TextNormalized *string `json:"text_normalized,omitempty"`

	// The details of the normalized text, if applicable. This element is optional; it is returned only if normalized text
	// exists.
	Interpretation *Interpretation `json:"interpretation,omitempty"`

	// Hashed values that you can send to IBM to provide feedback or receive support.
	ProvenanceIds []string `json:"provenance_ids,omitempty"`

	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *Location `json:"location,omitempty"`
}

// Constants associated with the ContractAmts.ConfidenceLevel property.
// The confidence level in the identification of the contract amount.
const (
	ContractAmtsConfidenceLevelHighConst   = "High"
	ContractAmtsConfidenceLevelLowConst    = "Low"
	ContractAmtsConfidenceLevelMediumConst = "Medium"
)

// UnmarshalContractAmts unmarshals an instance of ContractAmts from the specified map of raw messages.
func UnmarshalContractAmts(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ContractAmts)
	err = core.UnmarshalPrimitive(m, "confidence_level", &obj.ConfidenceLevel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text_normalized", &obj.TextNormalized)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "interpretation", &obj.Interpretation, UnmarshalInterpretation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "provenance_ids", &obj.ProvenanceIds)
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

// ContractCurrencies : The contract currencies that are declared in the document.
type ContractCurrencies struct {
	// The confidence level in the identification of the contract currency.
	ConfidenceLevel *string `json:"confidence_level,omitempty"`

	// The contract currency.
	Text *string `json:"text,omitempty"`

	// The normalized form of the contract currency, which is listed as a string in
	// [ISO-4217](https://www.iso.org/iso-4217-currency-codes.html) format. This element is optional; it is returned only
	// if normalized text exists.
	TextNormalized *string `json:"text_normalized,omitempty"`

	// Hashed values that you can send to IBM to provide feedback or receive support.
	ProvenanceIds []string `json:"provenance_ids,omitempty"`

	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *Location `json:"location,omitempty"`
}

// Constants associated with the ContractCurrencies.ConfidenceLevel property.
// The confidence level in the identification of the contract currency.
const (
	ContractCurrenciesConfidenceLevelHighConst   = "High"
	ContractCurrenciesConfidenceLevelLowConst    = "Low"
	ContractCurrenciesConfidenceLevelMediumConst = "Medium"
)

// UnmarshalContractCurrencies unmarshals an instance of ContractCurrencies from the specified map of raw messages.
func UnmarshalContractCurrencies(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ContractCurrencies)
	err = core.UnmarshalPrimitive(m, "confidence_level", &obj.ConfidenceLevel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text_normalized", &obj.TextNormalized)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "provenance_ids", &obj.ProvenanceIds)
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

// ContractTerms : The duration or durations of the contract.
type ContractTerms struct {
	// The confidence level in the identification of the contract term.
	ConfidenceLevel *string `json:"confidence_level,omitempty"`

	// The contract term (duration).
	Text *string `json:"text,omitempty"`

	// The normalized form of the contract term, which is listed as a string. This element is optional; it is returned only
	// if normalized text exists.
	TextNormalized *string `json:"text_normalized,omitempty"`

	// The details of the normalized text, if applicable. This element is optional; it is returned only if normalized text
	// exists.
	Interpretation *Interpretation `json:"interpretation,omitempty"`

	// Hashed values that you can send to IBM to provide feedback or receive support.
	ProvenanceIds []string `json:"provenance_ids,omitempty"`

	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *Location `json:"location,omitempty"`
}

// Constants associated with the ContractTerms.ConfidenceLevel property.
// The confidence level in the identification of the contract term.
const (
	ContractTermsConfidenceLevelHighConst   = "High"
	ContractTermsConfidenceLevelLowConst    = "Low"
	ContractTermsConfidenceLevelMediumConst = "Medium"
)

// UnmarshalContractTerms unmarshals an instance of ContractTerms from the specified map of raw messages.
func UnmarshalContractTerms(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ContractTerms)
	err = core.UnmarshalPrimitive(m, "confidence_level", &obj.ConfidenceLevel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text_normalized", &obj.TextNormalized)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "interpretation", &obj.Interpretation, UnmarshalInterpretation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "provenance_ids", &obj.ProvenanceIds)
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

// ContractTypes : The contract type identified in the input document.
type ContractTypes struct {
	// The confidence level in the identification of the contract type.
	ConfidenceLevel *string `json:"confidence_level,omitempty"`

	// The contract type.
	Text *string `json:"text,omitempty"`

	// Hashed values that you can send to IBM to provide feedback or receive support.
	ProvenanceIds []string `json:"provenance_ids,omitempty"`

	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *Location `json:"location,omitempty"`
}

// Constants associated with the ContractTypes.ConfidenceLevel property.
// The confidence level in the identification of the contract type.
const (
	ContractTypesConfidenceLevelHighConst   = "High"
	ContractTypesConfidenceLevelLowConst    = "Low"
	ContractTypesConfidenceLevelMediumConst = "Medium"
)

// UnmarshalContractTypes unmarshals an instance of ContractTypes from the specified map of raw messages.
func UnmarshalContractTypes(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ContractTypes)
	err = core.UnmarshalPrimitive(m, "confidence_level", &obj.ConfidenceLevel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "provenance_ids", &obj.ProvenanceIds)
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

// ConvertToHTMLOptions : The ConvertToHTML options.
type ConvertToHTMLOptions struct {
	// The document to convert.
	File io.ReadCloser `validate:"required"`

	// The content type of file.
	FileContentType *string

	// The analysis model to be used by the service. For the **Element classification** and **Compare two documents**
	// methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults
	// apply to the standalone methods as well as to the methods' use in batch-processing requests.
	Model *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ConvertToHTMLOptions.Model property.
// The analysis model to be used by the service. For the **Element classification** and **Compare two documents**
// methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults apply
// to the standalone methods as well as to the methods' use in batch-processing requests.
const (
	ConvertToHTMLOptionsModelContractsConst = "contracts"
	ConvertToHTMLOptionsModelTablesConst    = "tables"
)

// NewConvertToHTMLOptions : Instantiate ConvertToHTMLOptions
func (*CompareComplyV1) NewConvertToHTMLOptions(file io.ReadCloser) *ConvertToHTMLOptions {
	return &ConvertToHTMLOptions{
		File: file,
	}
}

// SetFile : Allow user to set File
func (options *ConvertToHTMLOptions) SetFile(file io.ReadCloser) *ConvertToHTMLOptions {
	options.File = file
	return options
}

// SetFileContentType : Allow user to set FileContentType
func (options *ConvertToHTMLOptions) SetFileContentType(fileContentType string) *ConvertToHTMLOptions {
	options.FileContentType = core.StringPtr(fileContentType)
	return options
}

// SetModel : Allow user to set Model
func (options *ConvertToHTMLOptions) SetModel(model string) *ConvertToHTMLOptions {
	options.Model = core.StringPtr(model)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ConvertToHTMLOptions) SetHeaders(param map[string]string) *ConvertToHTMLOptions {
	options.Headers = param
	return options
}

// CreateBatchOptions : The CreateBatch options.
type CreateBatchOptions struct {
	// The Compare and Comply method to run across the submitted input documents.
	Function *string `validate:"required"`

	// A JSON file containing the input Cloud Object Storage credentials. At a minimum, the credentials must enable `READ`
	// permissions on the bucket defined by the `input_bucket_name` parameter.
	InputCredentialsFile io.ReadCloser `validate:"required"`

	// The geographical location of the Cloud Object Storage input bucket as listed on the **Endpoint** tab of your Cloud
	// Object Storage instance; for example, `us-geo`, `eu-geo`, or `ap-geo`.
	InputBucketLocation *string `validate:"required"`

	// The name of the Cloud Object Storage input bucket.
	InputBucketName *string `validate:"required"`

	// A JSON file that lists the Cloud Object Storage output credentials. At a minimum, the credentials must enable `READ`
	// and `WRITE` permissions on the bucket defined by the `output_bucket_name` parameter.
	OutputCredentialsFile io.ReadCloser `validate:"required"`

	// The geographical location of the Cloud Object Storage output bucket as listed on the **Endpoint** tab of your Cloud
	// Object Storage instance; for example, `us-geo`, `eu-geo`, or `ap-geo`.
	OutputBucketLocation *string `validate:"required"`

	// The name of the Cloud Object Storage output bucket.
	OutputBucketName *string `validate:"required"`

	// The analysis model to be used by the service. For the **Element classification** and **Compare two documents**
	// methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults
	// apply to the standalone methods as well as to the methods' use in batch-processing requests.
	Model *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateBatchOptions.Function property.
// The Compare and Comply method to run across the submitted input documents.
const (
	CreateBatchOptionsFunctionElementClassificationConst = "element_classification"
	CreateBatchOptionsFunctionHTMLConversionConst        = "html_conversion"
	CreateBatchOptionsFunctionTablesConst                = "tables"
)

// Constants associated with the CreateBatchOptions.Model property.
// The analysis model to be used by the service. For the **Element classification** and **Compare two documents**
// methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults apply
// to the standalone methods as well as to the methods' use in batch-processing requests.
const (
	CreateBatchOptionsModelContractsConst = "contracts"
	CreateBatchOptionsModelTablesConst    = "tables"
)

// NewCreateBatchOptions : Instantiate CreateBatchOptions
func (*CompareComplyV1) NewCreateBatchOptions(function string, inputCredentialsFile io.ReadCloser, inputBucketLocation string, inputBucketName string, outputCredentialsFile io.ReadCloser, outputBucketLocation string, outputBucketName string) *CreateBatchOptions {
	return &CreateBatchOptions{
		Function:              core.StringPtr(function),
		InputCredentialsFile:  inputCredentialsFile,
		InputBucketLocation:   core.StringPtr(inputBucketLocation),
		InputBucketName:       core.StringPtr(inputBucketName),
		OutputCredentialsFile: outputCredentialsFile,
		OutputBucketLocation:  core.StringPtr(outputBucketLocation),
		OutputBucketName:      core.StringPtr(outputBucketName),
	}
}

// SetFunction : Allow user to set Function
func (options *CreateBatchOptions) SetFunction(function string) *CreateBatchOptions {
	options.Function = core.StringPtr(function)
	return options
}

// SetInputCredentialsFile : Allow user to set InputCredentialsFile
func (options *CreateBatchOptions) SetInputCredentialsFile(inputCredentialsFile io.ReadCloser) *CreateBatchOptions {
	options.InputCredentialsFile = inputCredentialsFile
	return options
}

// SetInputBucketLocation : Allow user to set InputBucketLocation
func (options *CreateBatchOptions) SetInputBucketLocation(inputBucketLocation string) *CreateBatchOptions {
	options.InputBucketLocation = core.StringPtr(inputBucketLocation)
	return options
}

// SetInputBucketName : Allow user to set InputBucketName
func (options *CreateBatchOptions) SetInputBucketName(inputBucketName string) *CreateBatchOptions {
	options.InputBucketName = core.StringPtr(inputBucketName)
	return options
}

// SetOutputCredentialsFile : Allow user to set OutputCredentialsFile
func (options *CreateBatchOptions) SetOutputCredentialsFile(outputCredentialsFile io.ReadCloser) *CreateBatchOptions {
	options.OutputCredentialsFile = outputCredentialsFile
	return options
}

// SetOutputBucketLocation : Allow user to set OutputBucketLocation
func (options *CreateBatchOptions) SetOutputBucketLocation(outputBucketLocation string) *CreateBatchOptions {
	options.OutputBucketLocation = core.StringPtr(outputBucketLocation)
	return options
}

// SetOutputBucketName : Allow user to set OutputBucketName
func (options *CreateBatchOptions) SetOutputBucketName(outputBucketName string) *CreateBatchOptions {
	options.OutputBucketName = core.StringPtr(outputBucketName)
	return options
}

// SetModel : Allow user to set Model
func (options *CreateBatchOptions) SetModel(model string) *CreateBatchOptions {
	options.Model = core.StringPtr(model)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateBatchOptions) SetHeaders(param map[string]string) *CreateBatchOptions {
	options.Headers = param
	return options
}

// DeleteFeedbackOptions : The DeleteFeedback options.
type DeleteFeedbackOptions struct {
	// A string that specifies the feedback entry to be deleted from the document.
	FeedbackID *string `validate:"required,ne="`

	// The analysis model to be used by the service. For the **Element classification** and **Compare two documents**
	// methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults
	// apply to the standalone methods as well as to the methods' use in batch-processing requests.
	Model *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the DeleteFeedbackOptions.Model property.
// The analysis model to be used by the service. For the **Element classification** and **Compare two documents**
// methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults apply
// to the standalone methods as well as to the methods' use in batch-processing requests.
const (
	DeleteFeedbackOptionsModelContractsConst = "contracts"
	DeleteFeedbackOptionsModelTablesConst    = "tables"
)

// NewDeleteFeedbackOptions : Instantiate DeleteFeedbackOptions
func (*CompareComplyV1) NewDeleteFeedbackOptions(feedbackID string) *DeleteFeedbackOptions {
	return &DeleteFeedbackOptions{
		FeedbackID: core.StringPtr(feedbackID),
	}
}

// SetFeedbackID : Allow user to set FeedbackID
func (options *DeleteFeedbackOptions) SetFeedbackID(feedbackID string) *DeleteFeedbackOptions {
	options.FeedbackID = core.StringPtr(feedbackID)
	return options
}

// SetModel : Allow user to set Model
func (options *DeleteFeedbackOptions) SetModel(model string) *DeleteFeedbackOptions {
	options.Model = core.StringPtr(model)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteFeedbackOptions) SetHeaders(param map[string]string) *DeleteFeedbackOptions {
	options.Headers = param
	return options
}

// DocCounts : Document counts.
type DocCounts struct {
	// Total number of documents.
	Total *int64 `json:"total,omitempty"`

	// Number of pending documents.
	Pending *int64 `json:"pending,omitempty"`

	// Number of documents successfully processed.
	Successful *int64 `json:"successful,omitempty"`

	// Number of documents not successfully processed.
	Failed *int64 `json:"failed,omitempty"`
}

// UnmarshalDocCounts unmarshals an instance of DocCounts from the specified map of raw messages.
func UnmarshalDocCounts(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DocCounts)
	err = core.UnmarshalPrimitive(m, "total", &obj.Total)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pending", &obj.Pending)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "successful", &obj.Successful)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "failed", &obj.Failed)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DocInfo : Information about the parsed input document.
type DocInfo struct {
	// The full text of the parsed document in HTML format.
	HTML *string `json:"html,omitempty"`

	// The title of the parsed document. If the service did not detect a title, the value of this element is `null`.
	Title *string `json:"title,omitempty"`

	// The MD5 hash of the input document.
	Hash *string `json:"hash,omitempty"`
}

// UnmarshalDocInfo unmarshals an instance of DocInfo from the specified map of raw messages.
func UnmarshalDocInfo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DocInfo)
	err = core.UnmarshalPrimitive(m, "html", &obj.HTML)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "title", &obj.Title)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hash", &obj.Hash)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DocStructure : The structure of the input document.
type DocStructure struct {
	// An array containing one object per section or subsection identified in the input document.
	SectionTitles []SectionTitles `json:"section_titles,omitempty"`

	// An array containing one object per section or subsection, in parallel with the `section_titles` array, that details
	// the leading sentences in the corresponding section or subsection.
	LeadingSentences []LeadingSentence `json:"leading_sentences,omitempty"`

	// An array containing one object per paragraph, in parallel with the `section_titles` and `leading_sentences` arrays.
	Paragraphs []Paragraphs `json:"paragraphs,omitempty"`
}

// UnmarshalDocStructure unmarshals an instance of DocStructure from the specified map of raw messages.
func UnmarshalDocStructure(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DocStructure)
	err = core.UnmarshalModel(m, "section_titles", &obj.SectionTitles, UnmarshalSectionTitles)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "leading_sentences", &obj.LeadingSentences, UnmarshalLeadingSentence)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "paragraphs", &obj.Paragraphs, UnmarshalParagraphs)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Document : Basic information about the input document.
type Document struct {
	// Document title, if detected.
	Title *string `json:"title,omitempty"`

	// The input document converted into HTML format.
	HTML *string `json:"html,omitempty"`

	// The MD5 hash value of the input document.
	Hash *string `json:"hash,omitempty"`

	// The label applied to the input document with the calling method's `file_1_label` or `file_2_label` value. This field
	// is specified only in the output of the **Comparing two documents** method.
	Label *string `json:"label,omitempty"`
}

// UnmarshalDocument unmarshals an instance of Document from the specified map of raw messages.
func UnmarshalDocument(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Document)
	err = core.UnmarshalPrimitive(m, "title", &obj.Title)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "html", &obj.HTML)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hash", &obj.Hash)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "label", &obj.Label)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EffectiveDates : An effective date.
type EffectiveDates struct {
	// The confidence level in the identification of the effective date.
	ConfidenceLevel *string `json:"confidence_level,omitempty"`

	// The effective date, listed as a string.
	Text *string `json:"text,omitempty"`

	// The normalized form of the effective date, which is listed as a string. This element is optional; it is returned
	// only if normalized text exists.
	TextNormalized *string `json:"text_normalized,omitempty"`

	// Hashed values that you can send to IBM to provide feedback or receive support.
	ProvenanceIds []string `json:"provenance_ids,omitempty"`

	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *Location `json:"location,omitempty"`
}

// Constants associated with the EffectiveDates.ConfidenceLevel property.
// The confidence level in the identification of the effective date.
const (
	EffectiveDatesConfidenceLevelHighConst   = "High"
	EffectiveDatesConfidenceLevelLowConst    = "Low"
	EffectiveDatesConfidenceLevelMediumConst = "Medium"
)

// UnmarshalEffectiveDates unmarshals an instance of EffectiveDates from the specified map of raw messages.
func UnmarshalEffectiveDates(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EffectiveDates)
	err = core.UnmarshalPrimitive(m, "confidence_level", &obj.ConfidenceLevel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text_normalized", &obj.TextNormalized)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "provenance_ids", &obj.ProvenanceIds)
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

// Element : A component part of the document.
type Element struct {
	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *Location `json:"location,omitempty"`

	// The text of the element.
	Text *string `json:"text,omitempty"`

	// Description of the action specified by the element  and whom it affects.
	Types []TypeLabel `json:"types,omitempty"`

	// List of functional categories into which the element falls; in other words, the subject matter of the element.
	Categories []Category `json:"categories,omitempty"`

	// List of document attributes.
	Attributes []Attribute `json:"attributes,omitempty"`
}

// UnmarshalElement unmarshals an instance of Element from the specified map of raw messages.
func UnmarshalElement(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Element)
	err = core.UnmarshalModel(m, "location", &obj.Location, UnmarshalLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "types", &obj.Types, UnmarshalTypeLabel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "categories", &obj.Categories, UnmarshalCategory)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "attributes", &obj.Attributes, UnmarshalAttribute)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ElementLocations : A list of `begin` and `end` indexes that indicate the locations of the elements in the input document.
type ElementLocations struct {
	// An integer that indicates the starting position of the element in the input document.
	Begin *int64 `json:"begin,omitempty"`

	// An integer that indicates the ending position of the element in the input document.
	End *int64 `json:"end,omitempty"`
}

// UnmarshalElementLocations unmarshals an instance of ElementLocations from the specified map of raw messages.
func UnmarshalElementLocations(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ElementLocations)
	err = core.UnmarshalPrimitive(m, "begin", &obj.Begin)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "end", &obj.End)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ElementPair : Details of semantically aligned elements.
type ElementPair struct {
	// The label of the document (that is, the value of either the `file_1_label` or `file_2_label` parameters) in which
	// the element occurs.
	DocumentLabel *string `json:"document_label,omitempty"`

	// The contents of the element.
	Text *string `json:"text,omitempty"`

	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *Location `json:"location,omitempty"`

	// Description of the action specified by the element and whom it affects.
	Types []TypeLabelComparison `json:"types,omitempty"`

	// List of functional categories into which the element falls; in other words, the subject matter of the element.
	Categories []CategoryComparison `json:"categories,omitempty"`

	// List of document attributes.
	Attributes []Attribute `json:"attributes,omitempty"`
}

// UnmarshalElementPair unmarshals an instance of ElementPair from the specified map of raw messages.
func UnmarshalElementPair(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ElementPair)
	err = core.UnmarshalPrimitive(m, "document_label", &obj.DocumentLabel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "location", &obj.Location, UnmarshalLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "types", &obj.Types, UnmarshalTypeLabelComparison)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "categories", &obj.Categories, UnmarshalCategoryComparison)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "attributes", &obj.Attributes, UnmarshalAttribute)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ExtractTablesOptions : The ExtractTables options.
type ExtractTablesOptions struct {
	// The document on which to run table extraction.
	File io.ReadCloser `validate:"required"`

	// The content type of file.
	FileContentType *string

	// The analysis model to be used by the service. For the **Element classification** and **Compare two documents**
	// methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults
	// apply to the standalone methods as well as to the methods' use in batch-processing requests.
	Model *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ExtractTablesOptions.Model property.
// The analysis model to be used by the service. For the **Element classification** and **Compare two documents**
// methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults apply
// to the standalone methods as well as to the methods' use in batch-processing requests.
const (
	ExtractTablesOptionsModelContractsConst = "contracts"
	ExtractTablesOptionsModelTablesConst    = "tables"
)

// NewExtractTablesOptions : Instantiate ExtractTablesOptions
func (*CompareComplyV1) NewExtractTablesOptions(file io.ReadCloser) *ExtractTablesOptions {
	return &ExtractTablesOptions{
		File: file,
	}
}

// SetFile : Allow user to set File
func (options *ExtractTablesOptions) SetFile(file io.ReadCloser) *ExtractTablesOptions {
	options.File = file
	return options
}

// SetFileContentType : Allow user to set FileContentType
func (options *ExtractTablesOptions) SetFileContentType(fileContentType string) *ExtractTablesOptions {
	options.FileContentType = core.StringPtr(fileContentType)
	return options
}

// SetModel : Allow user to set Model
func (options *ExtractTablesOptions) SetModel(model string) *ExtractTablesOptions {
	options.Model = core.StringPtr(model)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ExtractTablesOptions) SetHeaders(param map[string]string) *ExtractTablesOptions {
	options.Headers = param
	return options
}

// FeedbackDataInput : Feedback data for submission.
type FeedbackDataInput struct {
	// The type of feedback. The only permitted value is `element_classification`.
	FeedbackType *string `json:"feedback_type" validate:"required"`

	// Brief information about the input document.
	Document *ShortDoc `json:"document,omitempty"`

	// An optional string identifying the model ID. The only permitted value is `contracts`.
	ModelID *string `json:"model_id,omitempty"`

	// An optional string identifying the version of the model used.
	ModelVersion *string `json:"model_version,omitempty"`

	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *Location `json:"location" validate:"required"`

	// The text on which to submit feedback.
	Text *string `json:"text" validate:"required"`

	// The original labeling from the input document, without the submitted feedback.
	OriginalLabels *OriginalLabelsIn `json:"original_labels" validate:"required"`

	// The updated labeling from the input document, accounting for the submitted feedback.
	UpdatedLabels *UpdatedLabelsIn `json:"updated_labels" validate:"required"`
}

// NewFeedbackDataInput : Instantiate FeedbackDataInput (Generic Model Constructor)
func (*CompareComplyV1) NewFeedbackDataInput(feedbackType string, location *Location, text string, originalLabels *OriginalLabelsIn, updatedLabels *UpdatedLabelsIn) (model *FeedbackDataInput, err error) {
	model = &FeedbackDataInput{
		FeedbackType:   core.StringPtr(feedbackType),
		Location:       location,
		Text:           core.StringPtr(text),
		OriginalLabels: originalLabels,
		UpdatedLabels:  updatedLabels,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalFeedbackDataInput unmarshals an instance of FeedbackDataInput from the specified map of raw messages.
func UnmarshalFeedbackDataInput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FeedbackDataInput)
	err = core.UnmarshalPrimitive(m, "feedback_type", &obj.FeedbackType)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "document", &obj.Document, UnmarshalShortDoc)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "model_id", &obj.ModelID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "model_version", &obj.ModelVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "location", &obj.Location, UnmarshalLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "original_labels", &obj.OriginalLabels, UnmarshalOriginalLabelsIn)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "updated_labels", &obj.UpdatedLabels, UnmarshalUpdatedLabelsIn)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FeedbackDataOutput : Information returned from the **Add Feedback** method.
type FeedbackDataOutput struct {
	// A string identifying the user adding the feedback. The only permitted value is `element_classification`.
	FeedbackType *string `json:"feedback_type,omitempty"`

	// Brief information about the input document.
	Document *ShortDoc `json:"document,omitempty"`

	// An optional string identifying the model ID. The only permitted value is `contracts`.
	ModelID *string `json:"model_id,omitempty"`

	// An optional string identifying the version of the model used.
	ModelVersion *string `json:"model_version,omitempty"`

	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *Location `json:"location,omitempty"`

	// The text to which the feedback applies.
	Text *string `json:"text,omitempty"`

	// The original labeling from the input document, without the submitted feedback.
	OriginalLabels *OriginalLabelsOut `json:"original_labels,omitempty"`

	// The updated labeling from the input document, accounting for the submitted feedback.
	UpdatedLabels *UpdatedLabelsOut `json:"updated_labels,omitempty"`

	// Pagination details, if required by the length of the output.
	Pagination *Pagination `json:"pagination,omitempty"`
}

// UnmarshalFeedbackDataOutput unmarshals an instance of FeedbackDataOutput from the specified map of raw messages.
func UnmarshalFeedbackDataOutput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FeedbackDataOutput)
	err = core.UnmarshalPrimitive(m, "feedback_type", &obj.FeedbackType)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "document", &obj.Document, UnmarshalShortDoc)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "model_id", &obj.ModelID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "model_version", &obj.ModelVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "location", &obj.Location, UnmarshalLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "original_labels", &obj.OriginalLabels, UnmarshalOriginalLabelsOut)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "updated_labels", &obj.UpdatedLabels, UnmarshalUpdatedLabelsOut)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "pagination", &obj.Pagination, UnmarshalPagination)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FeedbackDeleted : The status and message of the deletion request.
type FeedbackDeleted struct {
	// HTTP return code.
	Status *int64 `json:"status,omitempty"`

	// Status message returned from the service.
	Message *string `json:"message,omitempty"`
}

// UnmarshalFeedbackDeleted unmarshals an instance of FeedbackDeleted from the specified map of raw messages.
func UnmarshalFeedbackDeleted(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FeedbackDeleted)
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FeedbackList : The results of a successful **List Feedback** request for all feedback.
type FeedbackList struct {
	// A list of all feedback for the document.
	Feedback []GetFeedback `json:"feedback,omitempty"`
}

// UnmarshalFeedbackList unmarshals an instance of FeedbackList from the specified map of raw messages.
func UnmarshalFeedbackList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FeedbackList)
	err = core.UnmarshalModel(m, "feedback", &obj.Feedback, UnmarshalGetFeedback)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FeedbackReturn : Information about the document and the submitted feedback.
type FeedbackReturn struct {
	// The unique ID of the feedback object.
	FeedbackID *string `json:"feedback_id,omitempty"`

	// An optional string identifying the person submitting feedback.
	UserID *string `json:"user_id,omitempty"`

	// An optional comment from the person submitting the feedback.
	Comment *string `json:"comment,omitempty"`

	// Timestamp listing the creation time of the feedback submission.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// Information returned from the **Add Feedback** method.
	FeedbackData *FeedbackDataOutput `json:"feedback_data,omitempty"`
}

// UnmarshalFeedbackReturn unmarshals an instance of FeedbackReturn from the specified map of raw messages.
func UnmarshalFeedbackReturn(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FeedbackReturn)
	err = core.UnmarshalPrimitive(m, "feedback_id", &obj.FeedbackID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "user_id", &obj.UserID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "comment", &obj.Comment)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created", &obj.Created)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "feedback_data", &obj.FeedbackData, UnmarshalFeedbackDataOutput)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetBatchOptions : The GetBatch options.
type GetBatchOptions struct {
	// The ID of the batch-processing job whose information you want to retrieve.
	BatchID *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetBatchOptions : Instantiate GetBatchOptions
func (*CompareComplyV1) NewGetBatchOptions(batchID string) *GetBatchOptions {
	return &GetBatchOptions{
		BatchID: core.StringPtr(batchID),
	}
}

// SetBatchID : Allow user to set BatchID
func (options *GetBatchOptions) SetBatchID(batchID string) *GetBatchOptions {
	options.BatchID = core.StringPtr(batchID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetBatchOptions) SetHeaders(param map[string]string) *GetBatchOptions {
	options.Headers = param
	return options
}

// GetFeedback : The results of a successful **Get Feedback** request for a single feedback entry.
type GetFeedback struct {
	// A string uniquely identifying the feedback entry.
	FeedbackID *string `json:"feedback_id,omitempty"`

	// A timestamp identifying the creation time of the feedback entry.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// A string containing the user's comment about the feedback entry.
	Comment *string `json:"comment,omitempty"`

	// Information returned from the **Add Feedback** method.
	FeedbackData *FeedbackDataOutput `json:"feedback_data,omitempty"`
}

// UnmarshalGetFeedback unmarshals an instance of GetFeedback from the specified map of raw messages.
func UnmarshalGetFeedback(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetFeedback)
	err = core.UnmarshalPrimitive(m, "feedback_id", &obj.FeedbackID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created", &obj.Created)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "comment", &obj.Comment)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "feedback_data", &obj.FeedbackData, UnmarshalFeedbackDataOutput)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetFeedbackOptions : The GetFeedback options.
type GetFeedbackOptions struct {
	// A string that specifies the feedback entry to be included in the output.
	FeedbackID *string `validate:"required,ne="`

	// The analysis model to be used by the service. For the **Element classification** and **Compare two documents**
	// methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults
	// apply to the standalone methods as well as to the methods' use in batch-processing requests.
	Model *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetFeedbackOptions.Model property.
// The analysis model to be used by the service. For the **Element classification** and **Compare two documents**
// methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults apply
// to the standalone methods as well as to the methods' use in batch-processing requests.
const (
	GetFeedbackOptionsModelContractsConst = "contracts"
	GetFeedbackOptionsModelTablesConst    = "tables"
)

// NewGetFeedbackOptions : Instantiate GetFeedbackOptions
func (*CompareComplyV1) NewGetFeedbackOptions(feedbackID string) *GetFeedbackOptions {
	return &GetFeedbackOptions{
		FeedbackID: core.StringPtr(feedbackID),
	}
}

// SetFeedbackID : Allow user to set FeedbackID
func (options *GetFeedbackOptions) SetFeedbackID(feedbackID string) *GetFeedbackOptions {
	options.FeedbackID = core.StringPtr(feedbackID)
	return options
}

// SetModel : Allow user to set Model
func (options *GetFeedbackOptions) SetModel(model string) *GetFeedbackOptions {
	options.Model = core.StringPtr(model)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetFeedbackOptions) SetHeaders(param map[string]string) *GetFeedbackOptions {
	options.Headers = param
	return options
}

// HTMLReturn : The HTML converted from an input document.
type HTMLReturn struct {
	// The number of pages in the input document.
	NumPages *string `json:"num_pages,omitempty"`

	// The author of the input document, if identified.
	Author *string `json:"author,omitempty"`

	// The publication date of the input document, if identified.
	PublicationDate *string `json:"publication_date,omitempty"`

	// The title of the input document, if identified.
	Title *string `json:"title,omitempty"`

	// The HTML version of the input document.
	HTML *string `json:"html,omitempty"`
}

// UnmarshalHTMLReturn unmarshals an instance of HTMLReturn from the specified map of raw messages.
func UnmarshalHTMLReturn(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(HTMLReturn)
	err = core.UnmarshalPrimitive(m, "num_pages", &obj.NumPages)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "author", &obj.Author)
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
	err = core.UnmarshalPrimitive(m, "html", &obj.HTML)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Interpretation : The details of the normalized text, if applicable. This element is optional; it is returned only if normalized text
// exists.
type Interpretation struct {
	// The value that was located in the normalized text.
	Value *string `json:"value,omitempty"`

	// An integer or float expressing the numeric value of the `value` key.
	NumericValue *float64 `json:"numeric_value,omitempty"`

	// A string listing the unit of the value that was found in the normalized text.
	//
	// **Note:** The value of `unit` is the [ISO-4217 currency code](https://www.iso.org/iso-4217-currency-codes.html)
	// identified for the currency amount (for example, `USD` or `EUR`). If the service cannot disambiguate a currency
	// symbol (for example, `$` or `£`), the value of `unit` contains the ambiguous symbol as-is.
	Unit *string `json:"unit,omitempty"`
}

// UnmarshalInterpretation unmarshals an instance of Interpretation from the specified map of raw messages.
func UnmarshalInterpretation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Interpretation)
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "numeric_value", &obj.NumericValue)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "unit", &obj.Unit)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Key : A key in a key-value pair.
type Key struct {
	// The unique ID of the key in the table.
	CellID *string `json:"cell_id,omitempty"`

	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *Location `json:"location,omitempty"`

	// The text content of the table cell without HTML markup.
	Text *string `json:"text,omitempty"`
}

// UnmarshalKey unmarshals an instance of Key from the specified map of raw messages.
func UnmarshalKey(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Key)
	err = core.UnmarshalPrimitive(m, "cell_id", &obj.CellID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "location", &obj.Location, UnmarshalLocation)
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

// KeyValuePair : Key-value pairs detected across cell boundaries.
type KeyValuePair struct {
	// A key in a key-value pair.
	Key *Key `json:"key,omitempty"`

	// A list of values in a key-value pair.
	Value []Value `json:"value,omitempty"`
}

// UnmarshalKeyValuePair unmarshals an instance of KeyValuePair from the specified map of raw messages.
func UnmarshalKeyValuePair(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeyValuePair)
	err = core.UnmarshalModel(m, "key", &obj.Key, UnmarshalKey)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "value", &obj.Value, UnmarshalValue)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Label : A pair of `nature` and `party` objects. The `nature` object identifies the effect of the element on the identified
// `party`, and the `party` object identifies the affected party.
type Label struct {
	// The identified `nature` of the element.
	Nature *string `json:"nature" validate:"required"`

	// The identified `party` of the element.
	Party *string `json:"party" validate:"required"`
}

// NewLabel : Instantiate Label (Generic Model Constructor)
func (*CompareComplyV1) NewLabel(nature string, party string) (model *Label, err error) {
	model = &Label{
		Nature: core.StringPtr(nature),
		Party:  core.StringPtr(party),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalLabel unmarshals an instance of Label from the specified map of raw messages.
func UnmarshalLabel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Label)
	err = core.UnmarshalPrimitive(m, "nature", &obj.Nature)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "party", &obj.Party)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LeadingSentence : The leading sentences in a section or subsection of the input document.
type LeadingSentence struct {
	// The text of the leading sentence.
	Text *string `json:"text,omitempty"`

	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *Location `json:"location,omitempty"`

	// An array of `location` objects that lists the locations of detected leading sentences.
	ElementLocations []ElementLocations `json:"element_locations,omitempty"`
}

// UnmarshalLeadingSentence unmarshals an instance of LeadingSentence from the specified map of raw messages.
func UnmarshalLeadingSentence(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LeadingSentence)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "location", &obj.Location, UnmarshalLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "element_locations", &obj.ElementLocations, UnmarshalElementLocations)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListBatchesOptions : The ListBatches options.
type ListBatchesOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListBatchesOptions : Instantiate ListBatchesOptions
func (*CompareComplyV1) NewListBatchesOptions() *ListBatchesOptions {
	return &ListBatchesOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListBatchesOptions) SetHeaders(param map[string]string) *ListBatchesOptions {
	options.Headers = param
	return options
}

// ListFeedbackOptions : The ListFeedback options.
type ListFeedbackOptions struct {
	// An optional string that filters the output to include only feedback with the specified feedback type. The only
	// permitted value is `element_classification`.
	FeedbackType *string

	// An optional string that filters the output to include only feedback from the document with the specified
	// `document_title`.
	DocumentTitle *string

	// An optional string that filters the output to include only feedback with the specified `model_id`. The only
	// permitted value is `contracts`.
	ModelID *string

	// An optional string that filters the output to include only feedback with the specified `model_version`.
	ModelVersion *string

	// An optional string in the form of a comma-separated list of categories. If it is specified, the service filters the
	// output to include only feedback that has at least one category from the list removed.
	CategoryRemoved *string

	// An optional string in the form of a comma-separated list of categories. If this is specified, the service filters
	// the output to include only feedback that has at least one category from the list added.
	CategoryAdded *string

	// An optional string in the form of a comma-separated list of categories. If this is specified, the service filters
	// the output to include only feedback that has at least one category from the list unchanged.
	CategoryNotChanged *string

	// An optional string of comma-separated `nature`:`party` pairs. If this is specified, the service filters the output
	// to include only feedback that has at least one `nature`:`party` pair from the list removed.
	TypeRemoved *string

	// An optional string of comma-separated `nature`:`party` pairs. If this is specified, the service filters the output
	// to include only feedback that has at least one `nature`:`party` pair from the list removed.
	TypeAdded *string

	// An optional string of comma-separated `nature`:`party` pairs. If this is specified, the service filters the output
	// to include only feedback that has at least one `nature`:`party` pair from the list unchanged.
	TypeNotChanged *string

	// An optional integer specifying the number of documents that you want the service to return.
	PageLimit *int64

	// An optional string that returns the set of documents after the previous set. Use this parameter with the
	// `page_limit` parameter.
	Cursor *string

	// An optional comma-separated list of fields in the document to sort on. You can optionally specify the sort direction
	// by prefixing the value of the field with `-` for descending order or `+` for ascending order (the default).
	// Currently permitted sorting fields are `created`, `user_id`, and `document_title`.
	Sort *string

	// An optional boolean value. If specified as `true`, the `pagination` object in the output includes a value called
	// `total` that gives the total count of feedback created.
	IncludeTotal *bool

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListFeedbackOptions : Instantiate ListFeedbackOptions
func (*CompareComplyV1) NewListFeedbackOptions() *ListFeedbackOptions {
	return &ListFeedbackOptions{}
}

// SetFeedbackType : Allow user to set FeedbackType
func (options *ListFeedbackOptions) SetFeedbackType(feedbackType string) *ListFeedbackOptions {
	options.FeedbackType = core.StringPtr(feedbackType)
	return options
}

// SetDocumentTitle : Allow user to set DocumentTitle
func (options *ListFeedbackOptions) SetDocumentTitle(documentTitle string) *ListFeedbackOptions {
	options.DocumentTitle = core.StringPtr(documentTitle)
	return options
}

// SetModelID : Allow user to set ModelID
func (options *ListFeedbackOptions) SetModelID(modelID string) *ListFeedbackOptions {
	options.ModelID = core.StringPtr(modelID)
	return options
}

// SetModelVersion : Allow user to set ModelVersion
func (options *ListFeedbackOptions) SetModelVersion(modelVersion string) *ListFeedbackOptions {
	options.ModelVersion = core.StringPtr(modelVersion)
	return options
}

// SetCategoryRemoved : Allow user to set CategoryRemoved
func (options *ListFeedbackOptions) SetCategoryRemoved(categoryRemoved string) *ListFeedbackOptions {
	options.CategoryRemoved = core.StringPtr(categoryRemoved)
	return options
}

// SetCategoryAdded : Allow user to set CategoryAdded
func (options *ListFeedbackOptions) SetCategoryAdded(categoryAdded string) *ListFeedbackOptions {
	options.CategoryAdded = core.StringPtr(categoryAdded)
	return options
}

// SetCategoryNotChanged : Allow user to set CategoryNotChanged
func (options *ListFeedbackOptions) SetCategoryNotChanged(categoryNotChanged string) *ListFeedbackOptions {
	options.CategoryNotChanged = core.StringPtr(categoryNotChanged)
	return options
}

// SetTypeRemoved : Allow user to set TypeRemoved
func (options *ListFeedbackOptions) SetTypeRemoved(typeRemoved string) *ListFeedbackOptions {
	options.TypeRemoved = core.StringPtr(typeRemoved)
	return options
}

// SetTypeAdded : Allow user to set TypeAdded
func (options *ListFeedbackOptions) SetTypeAdded(typeAdded string) *ListFeedbackOptions {
	options.TypeAdded = core.StringPtr(typeAdded)
	return options
}

// SetTypeNotChanged : Allow user to set TypeNotChanged
func (options *ListFeedbackOptions) SetTypeNotChanged(typeNotChanged string) *ListFeedbackOptions {
	options.TypeNotChanged = core.StringPtr(typeNotChanged)
	return options
}

// SetPageLimit : Allow user to set PageLimit
func (options *ListFeedbackOptions) SetPageLimit(pageLimit int64) *ListFeedbackOptions {
	options.PageLimit = core.Int64Ptr(pageLimit)
	return options
}

// SetCursor : Allow user to set Cursor
func (options *ListFeedbackOptions) SetCursor(cursor string) *ListFeedbackOptions {
	options.Cursor = core.StringPtr(cursor)
	return options
}

// SetSort : Allow user to set Sort
func (options *ListFeedbackOptions) SetSort(sort string) *ListFeedbackOptions {
	options.Sort = core.StringPtr(sort)
	return options
}

// SetIncludeTotal : Allow user to set IncludeTotal
func (options *ListFeedbackOptions) SetIncludeTotal(includeTotal bool) *ListFeedbackOptions {
	options.IncludeTotal = core.BoolPtr(includeTotal)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListFeedbackOptions) SetHeaders(param map[string]string) *ListFeedbackOptions {
	options.Headers = param
	return options
}

// Location : The numeric location of the identified element in the document, represented with two integers labeled `begin` and
// `end`.
type Location struct {
	// The element's `begin` index.
	Begin *int64 `json:"begin" validate:"required"`

	// The element's `end` index.
	End *int64 `json:"end" validate:"required"`
}

// NewLocation : Instantiate Location (Generic Model Constructor)
func (*CompareComplyV1) NewLocation(begin int64, end int64) (model *Location, err error) {
	model = &Location{
		Begin: core.Int64Ptr(begin),
		End:   core.Int64Ptr(end),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalLocation unmarshals an instance of Location from the specified map of raw messages.
func UnmarshalLocation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Location)
	err = core.UnmarshalPrimitive(m, "begin", &obj.Begin)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "end", &obj.End)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Mention : A mention of a party.
type Mention struct {
	// The name of the party.
	Text *string `json:"text,omitempty"`

	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *Location `json:"location,omitempty"`
}

// UnmarshalMention unmarshals an instance of Mention from the specified map of raw messages.
func UnmarshalMention(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Mention)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
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

// OriginalLabelsIn : The original labeling from the input document, without the submitted feedback.
type OriginalLabelsIn struct {
	// Description of the action specified by the element and whom it affects.
	Types []TypeLabel `json:"types" validate:"required"`

	// List of functional categories into which the element falls; in other words, the subject matter of the element.
	Categories []Category `json:"categories" validate:"required"`
}

// NewOriginalLabelsIn : Instantiate OriginalLabelsIn (Generic Model Constructor)
func (*CompareComplyV1) NewOriginalLabelsIn(types []TypeLabel, categories []Category) (model *OriginalLabelsIn, err error) {
	model = &OriginalLabelsIn{
		Types:      types,
		Categories: categories,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalOriginalLabelsIn unmarshals an instance of OriginalLabelsIn from the specified map of raw messages.
func UnmarshalOriginalLabelsIn(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OriginalLabelsIn)
	err = core.UnmarshalModel(m, "types", &obj.Types, UnmarshalTypeLabel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "categories", &obj.Categories, UnmarshalCategory)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// OriginalLabelsOut : The original labeling from the input document, without the submitted feedback.
type OriginalLabelsOut struct {
	// Description of the action specified by the element and whom it affects.
	Types []TypeLabel `json:"types,omitempty"`

	// List of functional categories into which the element falls; in other words, the subject matter of the element.
	Categories []Category `json:"categories,omitempty"`
}

// UnmarshalOriginalLabelsOut unmarshals an instance of OriginalLabelsOut from the specified map of raw messages.
func UnmarshalOriginalLabelsOut(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OriginalLabelsOut)
	err = core.UnmarshalModel(m, "types", &obj.Types, UnmarshalTypeLabel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "categories", &obj.Categories, UnmarshalCategory)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Pagination : Pagination details, if required by the length of the output.
type Pagination struct {
	// A token identifying the current page of results.
	RefreshCursor *string `json:"refresh_cursor,omitempty"`

	// A token identifying the next page of results.
	NextCursor *string `json:"next_cursor,omitempty"`

	// The URL that returns the current page of results.
	RefreshURL *string `json:"refresh_url,omitempty"`

	// The URL that returns the next page of results.
	NextURL *string `json:"next_url,omitempty"`

	// Reserved for future use.
	Total *int64 `json:"total,omitempty"`
}

// UnmarshalPagination unmarshals an instance of Pagination from the specified map of raw messages.
func UnmarshalPagination(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Pagination)
	err = core.UnmarshalPrimitive(m, "refresh_cursor", &obj.RefreshCursor)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "next_cursor", &obj.NextCursor)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "refresh_url", &obj.RefreshURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "next_url", &obj.NextURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total", &obj.Total)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Paragraphs : The locations of each paragraph in the input document.
type Paragraphs struct {
	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *Location `json:"location,omitempty"`
}

// UnmarshalParagraphs unmarshals an instance of Paragraphs from the specified map of raw messages.
func UnmarshalParagraphs(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Paragraphs)
	err = core.UnmarshalModel(m, "location", &obj.Location, UnmarshalLocation)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Parties : A party and its corresponding role, including address and contact information if identified.
type Parties struct {
	// The normalized form of the party's name.
	Party *string `json:"party,omitempty"`

	// A string identifying the party's role.
	Role *string `json:"role,omitempty"`

	// A string that identifies the importance of the party.
	Importance *string `json:"importance,omitempty"`

	// A list of the party's address or addresses.
	Addresses []Address `json:"addresses,omitempty"`

	// A list of the names and roles of contacts identified in the input document.
	Contacts []Contact `json:"contacts,omitempty"`

	// A list of the party's mentions in the input document.
	Mentions []Mention `json:"mentions,omitempty"`
}

// Constants associated with the Parties.Importance property.
// A string that identifies the importance of the party.
const (
	PartiesImportancePrimaryConst = "Primary"
	PartiesImportanceUnknownConst = "Unknown"
)

// UnmarshalParties unmarshals an instance of Parties from the specified map of raw messages.
func UnmarshalParties(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Parties)
	err = core.UnmarshalPrimitive(m, "party", &obj.Party)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "role", &obj.Role)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "importance", &obj.Importance)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "addresses", &obj.Addresses, UnmarshalAddress)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "contacts", &obj.Contacts, UnmarshalContact)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "mentions", &obj.Mentions, UnmarshalMention)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PaymentTerms : The document's payment duration or durations.
type PaymentTerms struct {
	// The confidence level in the identification of the payment term.
	ConfidenceLevel *string `json:"confidence_level,omitempty"`

	// The payment term (duration).
	Text *string `json:"text,omitempty"`

	// The normalized form of the payment term, which is listed as a string. This element is optional; it is returned only
	// if normalized text exists.
	TextNormalized *string `json:"text_normalized,omitempty"`

	// The details of the normalized text, if applicable. This element is optional; it is returned only if normalized text
	// exists.
	Interpretation *Interpretation `json:"interpretation,omitempty"`

	// Hashed values that you can send to IBM to provide feedback or receive support.
	ProvenanceIds []string `json:"provenance_ids,omitempty"`

	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *Location `json:"location,omitempty"`
}

// Constants associated with the PaymentTerms.ConfidenceLevel property.
// The confidence level in the identification of the payment term.
const (
	PaymentTermsConfidenceLevelHighConst   = "High"
	PaymentTermsConfidenceLevelLowConst    = "Low"
	PaymentTermsConfidenceLevelMediumConst = "Medium"
)

// UnmarshalPaymentTerms unmarshals an instance of PaymentTerms from the specified map of raw messages.
func UnmarshalPaymentTerms(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PaymentTerms)
	err = core.UnmarshalPrimitive(m, "confidence_level", &obj.ConfidenceLevel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text_normalized", &obj.TextNormalized)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "interpretation", &obj.Interpretation, UnmarshalInterpretation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "provenance_ids", &obj.ProvenanceIds)
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

// RowHeaders : Row-level cells, each applicable as a header to other cells in the same row as itself, of the current table.
type RowHeaders struct {
	// The unique ID of the cell in the current table.
	CellID *string `json:"cell_id,omitempty"`

	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *Location `json:"location,omitempty"`

	// The textual contents of this cell from the input document without associated markup content.
	Text *string `json:"text,omitempty"`

	// If you provide customization input, the normalized version of the cell text according to the customization;
	// otherwise, the same value as `text`.
	TextNormalized *string `json:"text_normalized,omitempty"`

	// The `begin` index of this cell's `row` location in the current table.
	RowIndexBegin *int64 `json:"row_index_begin,omitempty"`

	// The `end` index of this cell's `row` location in the current table.
	RowIndexEnd *int64 `json:"row_index_end,omitempty"`

	// The `begin` index of this cell's `column` location in the current table.
	ColumnIndexBegin *int64 `json:"column_index_begin,omitempty"`

	// The `end` index of this cell's `column` location in the current table.
	ColumnIndexEnd *int64 `json:"column_index_end,omitempty"`
}

// UnmarshalRowHeaders unmarshals an instance of RowHeaders from the specified map of raw messages.
func UnmarshalRowHeaders(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RowHeaders)
	err = core.UnmarshalPrimitive(m, "cell_id", &obj.CellID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "location", &obj.Location, UnmarshalLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text_normalized", &obj.TextNormalized)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "row_index_begin", &obj.RowIndexBegin)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "row_index_end", &obj.RowIndexEnd)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "column_index_begin", &obj.ColumnIndexBegin)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "column_index_end", &obj.ColumnIndexEnd)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SectionTitle : The table's section title, if identified.
type SectionTitle struct {
	// The text of the section title, if identified.
	Text *string `json:"text,omitempty"`

	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *Location `json:"location,omitempty"`
}

// UnmarshalSectionTitle unmarshals an instance of SectionTitle from the specified map of raw messages.
func UnmarshalSectionTitle(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SectionTitle)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
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

// SectionTitles : An array containing one object per section or subsection detected in the input document. Sections and subsections are
// not nested; instead, they are flattened out and can be placed back in order by using the `begin` and `end` values of
// the element and the `level` value of the section.
type SectionTitles struct {
	// The text of the section title, if identified.
	Text *string `json:"text,omitempty"`

	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *Location `json:"location,omitempty"`

	// An integer indicating the level at which the section is located in the input document. For example, `1` represents a
	// top-level section, `2` represents a subsection within the level `1` section, and so forth.
	Level *int64 `json:"level,omitempty"`

	// An array of `location` objects that lists the locations of detected section titles.
	ElementLocations []ElementLocations `json:"element_locations,omitempty"`
}

// UnmarshalSectionTitles unmarshals an instance of SectionTitles from the specified map of raw messages.
func UnmarshalSectionTitles(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SectionTitles)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "location", &obj.Location, UnmarshalLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "level", &obj.Level)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "element_locations", &obj.ElementLocations, UnmarshalElementLocations)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ShortDoc : Brief information about the input document.
type ShortDoc struct {
	// The title of the input document, if identified.
	Title *string `json:"title,omitempty"`

	// The MD5 hash of the input document.
	Hash *string `json:"hash,omitempty"`
}

// UnmarshalShortDoc unmarshals an instance of ShortDoc from the specified map of raw messages.
func UnmarshalShortDoc(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ShortDoc)
	err = core.UnmarshalPrimitive(m, "title", &obj.Title)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hash", &obj.Hash)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TableHeaders : The contents of the current table's header.
type TableHeaders struct {
	// The unique ID of the cell in the current table.
	CellID *string `json:"cell_id,omitempty"`

	// The location of the table header cell in the current table as defined by its `begin` and `end` offsets,
	// respectfully, in the input document.
	Location interface{} `json:"location,omitempty"`

	// The textual contents of the cell from the input document without associated markup content.
	Text *string `json:"text,omitempty"`

	// The `begin` index of this cell's `row` location in the current table.
	RowIndexBegin *int64 `json:"row_index_begin,omitempty"`

	// The `end` index of this cell's `row` location in the current table.
	RowIndexEnd *int64 `json:"row_index_end,omitempty"`

	// The `begin` index of this cell's `column` location in the current table.
	ColumnIndexBegin *int64 `json:"column_index_begin,omitempty"`

	// The `end` index of this cell's `column` location in the current table.
	ColumnIndexEnd *int64 `json:"column_index_end,omitempty"`
}

// UnmarshalTableHeaders unmarshals an instance of TableHeaders from the specified map of raw messages.
func UnmarshalTableHeaders(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TableHeaders)
	err = core.UnmarshalPrimitive(m, "cell_id", &obj.CellID)
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
	err = core.UnmarshalPrimitive(m, "row_index_begin", &obj.RowIndexBegin)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "row_index_end", &obj.RowIndexEnd)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "column_index_begin", &obj.ColumnIndexBegin)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "column_index_end", &obj.ColumnIndexEnd)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TableReturn : The analysis of the document's tables.
type TableReturn struct {
	// Information about the parsed input document.
	Document *DocInfo `json:"document,omitempty"`

	// The ID of the model used to extract the table contents. The value for table extraction is `tables`.
	ModelID *string `json:"model_id,omitempty"`

	// The version of the `tables` model ID.
	ModelVersion *string `json:"model_version,omitempty"`

	// Definitions of the tables identified in the input document.
	Tables []Tables `json:"tables,omitempty"`
}

// UnmarshalTableReturn unmarshals an instance of TableReturn from the specified map of raw messages.
func UnmarshalTableReturn(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TableReturn)
	err = core.UnmarshalModel(m, "document", &obj.Document, UnmarshalDocInfo)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "model_id", &obj.ModelID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "model_version", &obj.ModelVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "tables", &obj.Tables, UnmarshalTables)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TableTitle : If identified, the title or caption of the current table of the form `Table x.: ...`. Empty when no title is
// identified. When exposed, the `title` is also excluded from the `contexts` array of the same table.
type TableTitle struct {
	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *Location `json:"location,omitempty"`

	// The text of the identified table title or caption.
	Text *string `json:"text,omitempty"`
}

// UnmarshalTableTitle unmarshals an instance of TableTitle from the specified map of raw messages.
func UnmarshalTableTitle(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TableTitle)
	err = core.UnmarshalModel(m, "location", &obj.Location, UnmarshalLocation)
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

// Tables : The contents of the tables extracted from a document.
type Tables struct {
	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *Location `json:"location,omitempty"`

	// The textual contents of the current table from the input document without associated markup content.
	Text *string `json:"text,omitempty"`

	// The table's section title, if identified.
	SectionTitle *SectionTitle `json:"section_title,omitempty"`

	// If identified, the title or caption of the current table of the form `Table x.: ...`. Empty when no title is
	// identified. When exposed, the `title` is also excluded from the `contexts` array of the same table.
	Title *TableTitle `json:"title,omitempty"`

	// An array of table-level cells that apply as headers to all the other cells in the current table.
	TableHeaders []TableHeaders `json:"table_headers,omitempty"`

	// An array of row-level cells, each applicable as a header to other cells in the same row as itself, of the current
	// table.
	RowHeaders []RowHeaders `json:"row_headers,omitempty"`

	// An array of column-level cells, each applicable as a header to other cells in the same column as itself, of the
	// current table.
	ColumnHeaders []ColumnHeaders `json:"column_headers,omitempty"`

	// An array of cells that are neither table header nor column header nor row header cells, of the current table with
	// corresponding row and column header associations.
	BodyCells []BodyCells `json:"body_cells,omitempty"`

	// An array of objects that list text that is related to the table contents and that precedes or follows the current
	// table.
	Contexts []Contexts `json:"contexts,omitempty"`

	// An array of key-value pairs identified in the current table.
	KeyValuePairs []KeyValuePair `json:"key_value_pairs,omitempty"`
}

// UnmarshalTables unmarshals an instance of Tables from the specified map of raw messages.
func UnmarshalTables(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Tables)
	err = core.UnmarshalModel(m, "location", &obj.Location, UnmarshalLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "section_title", &obj.SectionTitle, UnmarshalSectionTitle)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "title", &obj.Title, UnmarshalTableTitle)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "table_headers", &obj.TableHeaders, UnmarshalTableHeaders)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "row_headers", &obj.RowHeaders, UnmarshalRowHeaders)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "column_headers", &obj.ColumnHeaders, UnmarshalColumnHeaders)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "body_cells", &obj.BodyCells, UnmarshalBodyCells)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "contexts", &obj.Contexts, UnmarshalContexts)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "key_value_pairs", &obj.KeyValuePairs, UnmarshalKeyValuePair)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TerminationDates : Termination dates identified in the input document.
type TerminationDates struct {
	// The confidence level in the identification of the termination date.
	ConfidenceLevel *string `json:"confidence_level,omitempty"`

	// The termination date.
	Text *string `json:"text,omitempty"`

	// The normalized form of the termination date, which is listed as a string. This element is optional; it is returned
	// only if normalized text exists.
	TextNormalized *string `json:"text_normalized,omitempty"`

	// Hashed values that you can send to IBM to provide feedback or receive support.
	ProvenanceIds []string `json:"provenance_ids,omitempty"`

	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *Location `json:"location,omitempty"`
}

// Constants associated with the TerminationDates.ConfidenceLevel property.
// The confidence level in the identification of the termination date.
const (
	TerminationDatesConfidenceLevelHighConst   = "High"
	TerminationDatesConfidenceLevelLowConst    = "Low"
	TerminationDatesConfidenceLevelMediumConst = "Medium"
)

// UnmarshalTerminationDates unmarshals an instance of TerminationDates from the specified map of raw messages.
func UnmarshalTerminationDates(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TerminationDates)
	err = core.UnmarshalPrimitive(m, "confidence_level", &obj.ConfidenceLevel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text_normalized", &obj.TextNormalized)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "provenance_ids", &obj.ProvenanceIds)
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

// TypeLabel : Identification of a specific type.
type TypeLabel struct {
	// A pair of `nature` and `party` objects. The `nature` object identifies the effect of the element on the identified
	// `party`, and the `party` object identifies the affected party.
	Label *Label `json:"label,omitempty"`

	// Hashed values that you can send to IBM to provide feedback or receive support.
	ProvenanceIds []string `json:"provenance_ids,omitempty"`

	// The type of modification of the feedback entry in the updated labels response.
	Modification *string `json:"modification,omitempty"`
}

// Constants associated with the TypeLabel.Modification property.
// The type of modification of the feedback entry in the updated labels response.
const (
	TypeLabelModificationAddedConst     = "added"
	TypeLabelModificationRemovedConst   = "removed"
	TypeLabelModificationUnchangedConst = "unchanged"
)

// UnmarshalTypeLabel unmarshals an instance of TypeLabel from the specified map of raw messages.
func UnmarshalTypeLabel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TypeLabel)
	err = core.UnmarshalModel(m, "label", &obj.Label, UnmarshalLabel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "provenance_ids", &obj.ProvenanceIds)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modification", &obj.Modification)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TypeLabelComparison : Identification of a specific type.
type TypeLabelComparison struct {
	// A pair of `nature` and `party` objects. The `nature` object identifies the effect of the element on the identified
	// `party`, and the `party` object identifies the affected party.
	Label *Label `json:"label,omitempty"`
}

// UnmarshalTypeLabelComparison unmarshals an instance of TypeLabelComparison from the specified map of raw messages.
func UnmarshalTypeLabelComparison(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TypeLabelComparison)
	err = core.UnmarshalModel(m, "label", &obj.Label, UnmarshalLabel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UnalignedElement : Element that does not align semantically between two compared documents.
type UnalignedElement struct {
	// The label assigned to the document by the value of the `file_1_label` or `file_2_label` parameters on the **Compare
	// two documents** method.
	DocumentLabel *string `json:"document_label,omitempty"`

	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *Location `json:"location,omitempty"`

	// The text of the element.
	Text *string `json:"text,omitempty"`

	// Description of the action specified by the element and whom it affects.
	Types []TypeLabelComparison `json:"types,omitempty"`

	// List of functional categories into which the element falls; in other words, the subject matter of the element.
	Categories []CategoryComparison `json:"categories,omitempty"`

	// List of document attributes.
	Attributes []Attribute `json:"attributes,omitempty"`
}

// UnmarshalUnalignedElement unmarshals an instance of UnalignedElement from the specified map of raw messages.
func UnmarshalUnalignedElement(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UnalignedElement)
	err = core.UnmarshalPrimitive(m, "document_label", &obj.DocumentLabel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "location", &obj.Location, UnmarshalLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "types", &obj.Types, UnmarshalTypeLabelComparison)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "categories", &obj.Categories, UnmarshalCategoryComparison)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "attributes", &obj.Attributes, UnmarshalAttribute)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateBatchOptions : The UpdateBatch options.
type UpdateBatchOptions struct {
	// The ID of the batch-processing job you want to update.
	BatchID *string `validate:"required,ne="`

	// The action you want to perform on the specified batch-processing job.
	Action *string `validate:"required"`

	// The analysis model to be used by the service. For the **Element classification** and **Compare two documents**
	// methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults
	// apply to the standalone methods as well as to the methods' use in batch-processing requests.
	Model *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateBatchOptions.Action property.
// The action you want to perform on the specified batch-processing job.
const (
	UpdateBatchOptionsActionCancelConst = "cancel"
	UpdateBatchOptionsActionRescanConst = "rescan"
)

// Constants associated with the UpdateBatchOptions.Model property.
// The analysis model to be used by the service. For the **Element classification** and **Compare two documents**
// methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults apply
// to the standalone methods as well as to the methods' use in batch-processing requests.
const (
	UpdateBatchOptionsModelContractsConst = "contracts"
	UpdateBatchOptionsModelTablesConst    = "tables"
)

// NewUpdateBatchOptions : Instantiate UpdateBatchOptions
func (*CompareComplyV1) NewUpdateBatchOptions(batchID string, action string) *UpdateBatchOptions {
	return &UpdateBatchOptions{
		BatchID: core.StringPtr(batchID),
		Action:  core.StringPtr(action),
	}
}

// SetBatchID : Allow user to set BatchID
func (options *UpdateBatchOptions) SetBatchID(batchID string) *UpdateBatchOptions {
	options.BatchID = core.StringPtr(batchID)
	return options
}

// SetAction : Allow user to set Action
func (options *UpdateBatchOptions) SetAction(action string) *UpdateBatchOptions {
	options.Action = core.StringPtr(action)
	return options
}

// SetModel : Allow user to set Model
func (options *UpdateBatchOptions) SetModel(model string) *UpdateBatchOptions {
	options.Model = core.StringPtr(model)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateBatchOptions) SetHeaders(param map[string]string) *UpdateBatchOptions {
	options.Headers = param
	return options
}

// UpdatedLabelsIn : The updated labeling from the input document, accounting for the submitted feedback.
type UpdatedLabelsIn struct {
	// Description of the action specified by the element and whom it affects.
	Types []TypeLabel `json:"types" validate:"required"`

	// List of functional categories into which the element falls; in other words, the subject matter of the element.
	Categories []Category `json:"categories" validate:"required"`
}

// NewUpdatedLabelsIn : Instantiate UpdatedLabelsIn (Generic Model Constructor)
func (*CompareComplyV1) NewUpdatedLabelsIn(types []TypeLabel, categories []Category) (model *UpdatedLabelsIn, err error) {
	model = &UpdatedLabelsIn{
		Types:      types,
		Categories: categories,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalUpdatedLabelsIn unmarshals an instance of UpdatedLabelsIn from the specified map of raw messages.
func UnmarshalUpdatedLabelsIn(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UpdatedLabelsIn)
	err = core.UnmarshalModel(m, "types", &obj.Types, UnmarshalTypeLabel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "categories", &obj.Categories, UnmarshalCategory)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdatedLabelsOut : The updated labeling from the input document, accounting for the submitted feedback.
type UpdatedLabelsOut struct {
	// Description of the action specified by the element and whom it affects.
	Types []TypeLabel `json:"types,omitempty"`

	// List of functional categories into which the element falls; in other words, the subject matter of the element.
	Categories []Category `json:"categories,omitempty"`
}

// UnmarshalUpdatedLabelsOut unmarshals an instance of UpdatedLabelsOut from the specified map of raw messages.
func UnmarshalUpdatedLabelsOut(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UpdatedLabelsOut)
	err = core.UnmarshalModel(m, "types", &obj.Types, UnmarshalTypeLabel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "categories", &obj.Categories, UnmarshalCategory)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Value : A value in a key-value pair.
type Value struct {
	// The unique ID of the value in the table.
	CellID *string `json:"cell_id,omitempty"`

	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *Location `json:"location,omitempty"`

	// The text content of the table cell without HTML markup.
	Text *string `json:"text,omitempty"`
}

// UnmarshalValue unmarshals an instance of Value from the specified map of raw messages.
func UnmarshalValue(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Value)
	err = core.UnmarshalPrimitive(m, "cell_id", &obj.CellID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "location", &obj.Location, UnmarshalLocation)
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
