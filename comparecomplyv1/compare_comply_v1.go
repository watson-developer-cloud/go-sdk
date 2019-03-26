// Package comparecomplyv1 : Operations and models for the CompareComplyV1 service
package comparecomplyv1

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
	"fmt"
	"github.com/IBM/go-sdk-core/core"
	"github.com/go-openapi/strfmt"
	common "github.com/watson-developer-cloud/go-sdk/common"
	"os"
)

// CompareComplyV1 : IBM Watson&trade; Compare and Comply analyzes governing documents to provide details about critical
// aspects of the documents.
//
// Version: V1
// See: http://www.ibm.com/watson/developercloud/compare-comply.html
type CompareComplyV1 struct {
	Service *core.BaseService
}

// CompareComplyV1Options : Service options
type CompareComplyV1Options struct {
	Version        string
	URL            string
	IAMApiKey      string
	IAMAccessToken string
	IAMURL         string
}

// NewCompareComplyV1 : Instantiate CompareComplyV1
func NewCompareComplyV1(options *CompareComplyV1Options) (*CompareComplyV1, error) {
	if options.URL == "" {
		options.URL = "https://gateway.watsonplatform.net/compare-comply/api"
	}

	serviceOptions := &core.ServiceOptions{
		URL:            options.URL,
		Version:        options.Version,
		IAMApiKey:      options.IAMApiKey,
		IAMAccessToken: options.IAMAccessToken,
		IAMURL:         options.IAMURL,
	}
	service, serviceErr := core.NewBaseService(serviceOptions, "compare-comply", "Compare Comply")
	if serviceErr != nil {
		return nil, serviceErr
	}

	return &CompareComplyV1{Service: service}, nil
}

// ConvertToHTML : Convert document to HTML
// Converts a document to HTML.
func (compareComply *CompareComplyV1) ConvertToHTML(convertToHTMLOptions *ConvertToHTMLOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(convertToHTMLOptions, "convertToHTMLOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(convertToHTMLOptions, "convertToHTMLOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/html_conversion"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(compareComply.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range convertToHTMLOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compare-comply", "V1", "ConvertToHTML")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")

	if convertToHTMLOptions.Model != nil {
		builder.AddQuery("model", fmt.Sprint(*convertToHTMLOptions.Model))
	}
	builder.AddQuery("version", compareComply.Service.Options.Version)

	builder.AddFormData("file", core.StringNilMapper(convertToHTMLOptions.Filename),
		core.StringNilMapper(convertToHTMLOptions.FileContentType), convertToHTMLOptions.File)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := compareComply.Service.Request(request, new(HTMLReturn))
	return response, err
}

// GetConvertToHTMLResult : Retrieve result of ConvertToHTML operation
func (compareComply *CompareComplyV1) GetConvertToHTMLResult(response *core.DetailedResponse) *HTMLReturn {
	result, ok := response.Result.(*HTMLReturn)
	if ok {
		return result
	}
	return nil
}

// ClassifyElements : Classify the elements of a document
// Analyzes the structural and semantic elements of a document.
func (compareComply *CompareComplyV1) ClassifyElements(classifyElementsOptions *ClassifyElementsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(classifyElementsOptions, "classifyElementsOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(classifyElementsOptions, "classifyElementsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/element_classification"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(compareComply.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range classifyElementsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compare-comply", "V1", "ClassifyElements")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")

	if classifyElementsOptions.Model != nil {
		builder.AddQuery("model", fmt.Sprint(*classifyElementsOptions.Model))
	}
	builder.AddQuery("version", compareComply.Service.Options.Version)

	builder.AddFormData("file", "filename",
		core.StringNilMapper(classifyElementsOptions.FileContentType), classifyElementsOptions.File)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := compareComply.Service.Request(request, new(ClassifyReturn))
	return response, err
}

// GetClassifyElementsResult : Retrieve result of ClassifyElements operation
func (compareComply *CompareComplyV1) GetClassifyElementsResult(response *core.DetailedResponse) *ClassifyReturn {
	result, ok := response.Result.(*ClassifyReturn)
	if ok {
		return result
	}
	return nil
}

// ExtractTables : Extract a document's tables
// Analyzes the tables in a document.
func (compareComply *CompareComplyV1) ExtractTables(extractTablesOptions *ExtractTablesOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(extractTablesOptions, "extractTablesOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(extractTablesOptions, "extractTablesOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/tables"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(compareComply.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range extractTablesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compare-comply", "V1", "ExtractTables")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")

	if extractTablesOptions.Model != nil {
		builder.AddQuery("model", fmt.Sprint(*extractTablesOptions.Model))
	}
	builder.AddQuery("version", compareComply.Service.Options.Version)

	builder.AddFormData("file", "filename",
		core.StringNilMapper(extractTablesOptions.FileContentType), extractTablesOptions.File)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := compareComply.Service.Request(request, new(TableReturn))
	return response, err
}

// GetExtractTablesResult : Retrieve result of ExtractTables operation
func (compareComply *CompareComplyV1) GetExtractTablesResult(response *core.DetailedResponse) *TableReturn {
	result, ok := response.Result.(*TableReturn)
	if ok {
		return result
	}
	return nil
}

// CompareDocuments : Compare two documents
// Compares two input documents. Documents must be in the same format.
func (compareComply *CompareComplyV1) CompareDocuments(compareDocumentsOptions *CompareDocumentsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(compareDocumentsOptions, "compareDocumentsOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(compareDocumentsOptions, "compareDocumentsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/comparison"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(compareComply.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range compareDocumentsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compare-comply", "V1", "CompareDocuments")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")

	if compareDocumentsOptions.File1Label != nil {
		builder.AddQuery("file_1_label", fmt.Sprint(*compareDocumentsOptions.File1Label))
	}
	if compareDocumentsOptions.File2Label != nil {
		builder.AddQuery("file_2_label", fmt.Sprint(*compareDocumentsOptions.File2Label))
	}
	if compareDocumentsOptions.Model != nil {
		builder.AddQuery("model", fmt.Sprint(*compareDocumentsOptions.Model))
	}
	builder.AddQuery("version", compareComply.Service.Options.Version)

	builder.AddFormData("file_1", "filename",
		core.StringNilMapper(compareDocumentsOptions.File1ContentType), compareDocumentsOptions.File1)
	builder.AddFormData("file_2", "filename",
		core.StringNilMapper(compareDocumentsOptions.File2ContentType), compareDocumentsOptions.File2)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := compareComply.Service.Request(request, new(CompareReturn))
	return response, err
}

// GetCompareDocumentsResult : Retrieve result of CompareDocuments operation
func (compareComply *CompareComplyV1) GetCompareDocumentsResult(response *core.DetailedResponse) *CompareReturn {
	result, ok := response.Result.(*CompareReturn)
	if ok {
		return result
	}
	return nil
}

// AddFeedback : Add feedback
// Adds feedback in the form of _labels_ from a subject-matter expert (SME) to a governing document.
// **Important:** Feedback is not immediately incorporated into the training model, nor is it guaranteed to be
// incorporated at a later date. Instead, submitted feedback is used to suggest future updates to the training model.
func (compareComply *CompareComplyV1) AddFeedback(addFeedbackOptions *AddFeedbackOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(addFeedbackOptions, "addFeedbackOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(addFeedbackOptions, "addFeedbackOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/feedback"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(compareComply.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range addFeedbackOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compare-comply", "V1", "AddFeedback")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", compareComply.Service.Options.Version)

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
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := compareComply.Service.Request(request, new(FeedbackReturn))
	return response, err
}

// GetAddFeedbackResult : Retrieve result of AddFeedback operation
func (compareComply *CompareComplyV1) GetAddFeedbackResult(response *core.DetailedResponse) *FeedbackReturn {
	result, ok := response.Result.(*FeedbackReturn)
	if ok {
		return result
	}
	return nil
}

// DeleteFeedback : Delete a specified feedback entry
// Deletes a feedback entry with a specified `feedback_id`.
func (compareComply *CompareComplyV1) DeleteFeedback(deleteFeedbackOptions *DeleteFeedbackOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(deleteFeedbackOptions, "deleteFeedbackOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(deleteFeedbackOptions, "deleteFeedbackOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/feedback"}
	pathParameters := []string{*deleteFeedbackOptions.FeedbackID}

	builder := core.NewRequestBuilder(core.DELETE)
	builder.ConstructHTTPURL(compareComply.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteFeedbackOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compare-comply", "V1", "DeleteFeedback")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")

	if deleteFeedbackOptions.Model != nil {
		builder.AddQuery("model", fmt.Sprint(*deleteFeedbackOptions.Model))
	}
	builder.AddQuery("version", compareComply.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := compareComply.Service.Request(request, new(FeedbackDeleted))
	return response, err
}

// GetDeleteFeedbackResult : Retrieve result of DeleteFeedback operation
func (compareComply *CompareComplyV1) GetDeleteFeedbackResult(response *core.DetailedResponse) *FeedbackDeleted {
	result, ok := response.Result.(*FeedbackDeleted)
	if ok {
		return result
	}
	return nil
}

// GetFeedback : List a specified feedback entry
// Lists a feedback entry with a specified `feedback_id`.
func (compareComply *CompareComplyV1) GetFeedback(getFeedbackOptions *GetFeedbackOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(getFeedbackOptions, "getFeedbackOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(getFeedbackOptions, "getFeedbackOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/feedback"}
	pathParameters := []string{*getFeedbackOptions.FeedbackID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(compareComply.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range getFeedbackOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compare-comply", "V1", "GetFeedback")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")

	if getFeedbackOptions.Model != nil {
		builder.AddQuery("model", fmt.Sprint(*getFeedbackOptions.Model))
	}
	builder.AddQuery("version", compareComply.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := compareComply.Service.Request(request, new(GetFeedback))
	return response, err
}

// GetGetFeedbackResult : Retrieve result of GetFeedback operation
func (compareComply *CompareComplyV1) GetGetFeedbackResult(response *core.DetailedResponse) *GetFeedback {
	result, ok := response.Result.(*GetFeedback)
	if ok {
		return result
	}
	return nil
}

// ListFeedback : List the feedback in a document
// Lists the feedback in a document.
func (compareComply *CompareComplyV1) ListFeedback(listFeedbackOptions *ListFeedbackOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateStruct(listFeedbackOptions, "listFeedbackOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/feedback"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(compareComply.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range listFeedbackOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compare-comply", "V1", "ListFeedback")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")

	if listFeedbackOptions.FeedbackType != nil {
		builder.AddQuery("feedback_type", fmt.Sprint(*listFeedbackOptions.FeedbackType))
	}
	if listFeedbackOptions.Before != nil {
		builder.AddQuery("before", fmt.Sprint(*listFeedbackOptions.Before))
	}
	if listFeedbackOptions.After != nil {
		builder.AddQuery("after", fmt.Sprint(*listFeedbackOptions.After))
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
	builder.AddQuery("version", compareComply.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := compareComply.Service.Request(request, new(FeedbackList))
	return response, err
}

// GetListFeedbackResult : Retrieve result of ListFeedback operation
func (compareComply *CompareComplyV1) GetListFeedbackResult(response *core.DetailedResponse) *FeedbackList {
	result, ok := response.Result.(*FeedbackList)
	if ok {
		return result
	}
	return nil
}

// CreateBatch : Submit a batch-processing request
// Run Compare and Comply methods over a collection of input documents.
// **Important:** Batch processing requires the use of the [IBM Cloud Object Storage
// service](https://cloud.ibm.com/docs/services/cloud-object-storage/about-cos.html#about-ibm-cloud-object-storage). The
// use of IBM Cloud Object Storage with Compare and Comply is discussed at [Using batch
// processing](https://cloud.ibm.com/docs/services/compare-comply/batching.html#before-you-batch).
func (compareComply *CompareComplyV1) CreateBatch(createBatchOptions *CreateBatchOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(createBatchOptions, "createBatchOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(createBatchOptions, "createBatchOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/batches"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(compareComply.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range createBatchOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compare-comply", "V1", "CreateBatch")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("function", fmt.Sprint(*createBatchOptions.Function))
	if createBatchOptions.Model != nil {
		builder.AddQuery("model", fmt.Sprint(*createBatchOptions.Model))
	}
	builder.AddQuery("version", compareComply.Service.Options.Version)

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
		return nil, err
	}

	response, err := compareComply.Service.Request(request, new(BatchStatus))
	return response, err
}

// GetCreateBatchResult : Retrieve result of CreateBatch operation
func (compareComply *CompareComplyV1) GetCreateBatchResult(response *core.DetailedResponse) *BatchStatus {
	result, ok := response.Result.(*BatchStatus)
	if ok {
		return result
	}
	return nil
}

// GetBatch : Get information about a specific batch-processing job
// Gets information about a batch-processing job with a specified ID.
func (compareComply *CompareComplyV1) GetBatch(getBatchOptions *GetBatchOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(getBatchOptions, "getBatchOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(getBatchOptions, "getBatchOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/batches"}
	pathParameters := []string{*getBatchOptions.BatchID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(compareComply.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range getBatchOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compare-comply", "V1", "GetBatch")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", compareComply.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := compareComply.Service.Request(request, new(BatchStatus))
	return response, err
}

// GetGetBatchResult : Retrieve result of GetBatch operation
func (compareComply *CompareComplyV1) GetGetBatchResult(response *core.DetailedResponse) *BatchStatus {
	result, ok := response.Result.(*BatchStatus)
	if ok {
		return result
	}
	return nil
}

// ListBatches : List submitted batch-processing jobs
// Lists batch-processing jobs submitted by users.
func (compareComply *CompareComplyV1) ListBatches(listBatchesOptions *ListBatchesOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateStruct(listBatchesOptions, "listBatchesOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/batches"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(compareComply.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range listBatchesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compare-comply", "V1", "ListBatches")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", compareComply.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := compareComply.Service.Request(request, new(Batches))
	return response, err
}

// GetListBatchesResult : Retrieve result of ListBatches operation
func (compareComply *CompareComplyV1) GetListBatchesResult(response *core.DetailedResponse) *Batches {
	result, ok := response.Result.(*Batches)
	if ok {
		return result
	}
	return nil
}

// UpdateBatch : Update a pending or active batch-processing job
// Updates a pending or active batch-processing job. You can rescan the input bucket to check for new documents or
// cancel a job.
func (compareComply *CompareComplyV1) UpdateBatch(updateBatchOptions *UpdateBatchOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(updateBatchOptions, "updateBatchOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(updateBatchOptions, "updateBatchOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/batches"}
	pathParameters := []string{*updateBatchOptions.BatchID}

	builder := core.NewRequestBuilder(core.PUT)
	builder.ConstructHTTPURL(compareComply.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range updateBatchOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compare-comply", "V1", "UpdateBatch")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("action", fmt.Sprint(*updateBatchOptions.Action))
	if updateBatchOptions.Model != nil {
		builder.AddQuery("model", fmt.Sprint(*updateBatchOptions.Model))
	}
	builder.AddQuery("version", compareComply.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := compareComply.Service.Request(request, new(BatchStatus))
	return response, err
}

// GetUpdateBatchResult : Retrieve result of UpdateBatch operation
func (compareComply *CompareComplyV1) GetUpdateBatchResult(response *core.DetailedResponse) *BatchStatus {
	result, ok := response.Result.(*BatchStatus)
	if ok {
		return result
	}
	return nil
}

// AddFeedbackOptions : The addFeedback options.
type AddFeedbackOptions struct {

	// Feedback data for submission.
	FeedbackData *FeedbackDataInput `json:"feedback_data" validate:"required"`

	// An optional string identifying the user.
	UserID *string `json:"user_id,omitempty"`

	// An optional comment on or description of the feedback.
	Comment *string `json:"comment,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewAddFeedbackOptions : Instantiate AddFeedbackOptions
func (compareComply *CompareComplyV1) NewAddFeedbackOptions(feedbackData *FeedbackDataInput) *AddFeedbackOptions {
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

// AlignedElement : AlignedElement struct
type AlignedElement struct {

	// Identifies two elements that semantically align between the compared documents.
	ElementPair []ElementPair `json:"element_pair,omitempty"`

	// Specifies whether the aligned element is identical. Elements are considered identical despite minor differences such
	// as leading punctuation, end-of-sentence punctuation, whitespace, the presence or absence of definite or indefinite
	// articles, and others.
	IdenticalText *bool `json:"identical_text,omitempty"`

	// One or more hashed values that you can send to IBM to provide feedback or receive support.
	ProvenanceIds []string `json:"provenance_ids,omitempty"`

	// Indicates that the elements aligned are contractual clauses of significance.
	SignificantElements *bool `json:"significant_elements,omitempty"`
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
	Attribute_Type_Currency     = "Currency"
	Attribute_Type_Datetime     = "DateTime"
	Attribute_Type_Duration     = "Duration"
	Attribute_Type_Location     = "Location"
	Attribute_Type_Organization = "Organization"
	Attribute_Type_Percentage   = "Percentage"
	Attribute_Type_Person       = "Person"
)

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
	BatchStatus_Function_ElementClassification = "element_classification"
	BatchStatus_Function_HTMLConversion        = "html_conversion"
	BatchStatus_Function_Tables                = "tables"
)

// Batches : The results of a successful `GET /v1/batches` request.
type Batches struct {

	// A list of the status of all batch requests.
	Batches []BatchStatus `json:"batches,omitempty"`
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

	RowHeaderIds []RowHeaderIds `json:"row_header_ids,omitempty"`

	RowHeaderTexts []RowHeaderTexts `json:"row_header_texts,omitempty"`

	RowHeaderTextsNormalized []RowHeaderTextsNormalized `json:"row_header_texts_normalized,omitempty"`

	ColumnHeaderIds []ColumnHeaderIds `json:"column_header_ids,omitempty"`

	ColumnHeaderTexts []ColumnHeaderTexts `json:"column_header_texts,omitempty"`

	ColumnHeaderTextsNormalized []ColumnHeaderTextsNormalized `json:"column_header_texts_normalized,omitempty"`

	Attributes []Attribute `json:"attributes,omitempty"`
}

// Category : Information defining an element's subject matter.
type Category struct {

	// The category of the associated element.
	Label *string `json:"label,omitempty"`

	// One or more hashed values that you can send to IBM to provide feedback or receive support.
	ProvenanceIds []string `json:"provenance_ids,omitempty"`
}

// Constants associated with the Category.Label property.
// The category of the associated element.
const (
	Category_Label_Amendments           = "Amendments"
	Category_Label_AssetUse             = "Asset Use"
	Category_Label_Assignments          = "Assignments"
	Category_Label_Audits               = "Audits"
	Category_Label_BusinessContinuity   = "Business Continuity"
	Category_Label_Communication        = "Communication"
	Category_Label_Confidentiality      = "Confidentiality"
	Category_Label_Deliverables         = "Deliverables"
	Category_Label_Delivery             = "Delivery"
	Category_Label_DisputeResolution    = "Dispute Resolution"
	Category_Label_ForceMajeure         = "Force Majeure"
	Category_Label_Indemnification      = "Indemnification"
	Category_Label_Insurance            = "Insurance"
	Category_Label_IntellectualProperty = "Intellectual Property"
	Category_Label_Liability            = "Liability"
	Category_Label_PaymentTermsBilling  = "Payment Terms & Billing"
	Category_Label_PricingTaxes         = "Pricing & Taxes"
	Category_Label_Privacy              = "Privacy"
	Category_Label_Responsibilities     = "Responsibilities"
	Category_Label_SafetyAndSecurity    = "Safety and Security"
	Category_Label_ScopeOfWork          = "Scope of Work"
	Category_Label_Subcontracts         = "Subcontracts"
	Category_Label_TermTermination      = "Term & Termination"
	Category_Label_Warranties           = "Warranties"
)

// CategoryComparison : Information defining an element's subject matter.
type CategoryComparison struct {

	// The category of the associated element.
	Label *string `json:"label,omitempty"`
}

// Constants associated with the CategoryComparison.Label property.
// The category of the associated element.
const (
	CategoryComparison_Label_Amendments           = "Amendments"
	CategoryComparison_Label_AssetUse             = "Asset Use"
	CategoryComparison_Label_Assignments          = "Assignments"
	CategoryComparison_Label_Audits               = "Audits"
	CategoryComparison_Label_BusinessContinuity   = "Business Continuity"
	CategoryComparison_Label_Communication        = "Communication"
	CategoryComparison_Label_Confidentiality      = "Confidentiality"
	CategoryComparison_Label_Deliverables         = "Deliverables"
	CategoryComparison_Label_Delivery             = "Delivery"
	CategoryComparison_Label_DisputeResolution    = "Dispute Resolution"
	CategoryComparison_Label_ForceMajeure         = "Force Majeure"
	CategoryComparison_Label_Indemnification      = "Indemnification"
	CategoryComparison_Label_Insurance            = "Insurance"
	CategoryComparison_Label_IntellectualProperty = "Intellectual Property"
	CategoryComparison_Label_Liability            = "Liability"
	CategoryComparison_Label_PaymentTermsBilling  = "Payment Terms & Billing"
	CategoryComparison_Label_PricingTaxes         = "Pricing & Taxes"
	CategoryComparison_Label_Privacy              = "Privacy"
	CategoryComparison_Label_Responsibilities     = "Responsibilities"
	CategoryComparison_Label_SafetyAndSecurity    = "Safety and Security"
	CategoryComparison_Label_ScopeOfWork          = "Scope of Work"
	CategoryComparison_Label_Subcontracts         = "Subcontracts"
	CategoryComparison_Label_TermTermination      = "Term & Termination"
	CategoryComparison_Label_Warranties           = "Warranties"
)

// ClassifyElementsOptions : The classifyElements options.
type ClassifyElementsOptions struct {

	// The document to classify.
	File *os.File `json:"file" validate:"required"`

	// The content type of file.
	FileContentType *string `json:"file_content_type,omitempty"`

	// The analysis model to be used by the service. For the **Element classification** and **Compare two documents**
	// methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults
	// apply to the standalone methods as well as to the methods' use in batch-processing requests.
	Model *string `json:"model,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the ClassifyElementsOptions.Model property.
// The analysis model to be used by the service. For the **Element classification** and **Compare two documents**
// methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults apply
// to the standalone methods as well as to the methods' use in batch-processing requests.
const (
	ClassifyElementsOptions_Model_Contracts = "contracts"
	ClassifyElementsOptions_Model_Tables    = "tables"
)

// NewClassifyElementsOptions : Instantiate ClassifyElementsOptions
func (compareComply *CompareComplyV1) NewClassifyElementsOptions(file *os.File) *ClassifyElementsOptions {
	return &ClassifyElementsOptions{
		File: file,
	}
}

// SetFile : Allow user to set File
func (options *ClassifyElementsOptions) SetFile(file *os.File) *ClassifyElementsOptions {
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

	// Definition of tables identified in the input document.
	Tables []Tables `json:"tables,omitempty"`

	// The structure of the input document.
	DocumentStructure *DocStructure `json:"document_structure,omitempty"`

	// Definitions of the parties identified in the input document.
	Parties []Parties `json:"parties,omitempty"`

	// The date or dates on which the document becomes effective.
	EffectiveDates []EffectiveDates `json:"effective_dates,omitempty"`

	// The monetary amounts that identify the total amount of the contract that needs to be paid from one party to another.
	ContractAmounts []ContractAmts `json:"contract_amounts,omitempty"`

	// The date or dates on which the document is to be terminated.
	TerminationDates []TerminationDates `json:"termination_dates,omitempty"`

	// The document's contract type or types as declared in the document.
	ContractType []ContractType `json:"contract_type,omitempty"`
}

// ColumnHeaderIds : An array of values, each being the `id` value of a column header that is applicable to the current cell.
type ColumnHeaderIds struct {

	// The `id` value of a column header.
	ID *string `json:"id,omitempty"`
}

// ColumnHeaderTexts : An array of values, each being the `text` value of a column header that is applicable to the current cell.
type ColumnHeaderTexts struct {

	// The `text` value of a column header.
	Text *string `json:"text,omitempty"`
}

// ColumnHeaderTextsNormalized : If you provide customization input, the normalized version of the column header texts according to the customization;
// otherwise, the same value as `column_header_texts`.
type ColumnHeaderTextsNormalized struct {

	// The normalized version of a column header text.
	TextNormalized *string `json:"text_normalized,omitempty"`
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

// CompareDocumentsOptions : The compareDocuments options.
type CompareDocumentsOptions struct {

	// The first document to compare.
	File1 *os.File `json:"file_1" validate:"required"`

	// The second document to compare.
	File2 *os.File `json:"file_2" validate:"required"`

	// The content type of file1.
	File1ContentType *string `json:"file_1_content_type,omitempty"`

	// The content type of file2.
	File2ContentType *string `json:"file_2_content_type,omitempty"`

	// A text label for the first document.
	File1Label *string `json:"file_1_label,omitempty"`

	// A text label for the second document.
	File2Label *string `json:"file_2_label,omitempty"`

	// The analysis model to be used by the service. For the **Element classification** and **Compare two documents**
	// methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults
	// apply to the standalone methods as well as to the methods' use in batch-processing requests.
	Model *string `json:"model,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the CompareDocumentsOptions.Model property.
// The analysis model to be used by the service. For the **Element classification** and **Compare two documents**
// methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults apply
// to the standalone methods as well as to the methods' use in batch-processing requests.
const (
	CompareDocumentsOptions_Model_Contracts = "contracts"
	CompareDocumentsOptions_Model_Tables    = "tables"
)

// NewCompareDocumentsOptions : Instantiate CompareDocumentsOptions
func (compareComply *CompareComplyV1) NewCompareDocumentsOptions(file1 *os.File, file2 *os.File) *CompareDocumentsOptions {
	return &CompareDocumentsOptions{
		File1: file1,
		File2: file2,
	}
}

// SetFile1 : Allow user to set File1
func (options *CompareDocumentsOptions) SetFile1(file1 *os.File) *CompareDocumentsOptions {
	options.File1 = file1
	return options
}

// SetFile2 : Allow user to set File2
func (options *CompareDocumentsOptions) SetFile2(file2 *os.File) *CompareDocumentsOptions {
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

// Contact : A contact.
type Contact struct {

	// A string listing the name of the contact.
	Name *string `json:"name,omitempty"`

	// A string listing the role of the contact.
	Role *string `json:"role,omitempty"`
}

// ContractAmts : A monetary amount identified in the input document.
type ContractAmts struct {

	// The monetary amount.
	Text *string `json:"text,omitempty"`

	// The confidence level in the identification of the contract amount.
	ConfidenceLevel *string `json:"confidence_level,omitempty"`

	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *Location `json:"location,omitempty"`
}

// Constants associated with the ContractAmts.ConfidenceLevel property.
// The confidence level in the identification of the contract amount.
const (
	ContractAmts_ConfidenceLevel_High   = "High"
	ContractAmts_ConfidenceLevel_Low    = "Low"
	ContractAmts_ConfidenceLevel_Medium = "Medium"
)

// ContractType : The contract type identified in the input document.
type ContractType struct {

	// The contract type.
	Text *string `json:"text,omitempty"`

	// The confidence level in the identification of the termination date.
	ConfidenceLevel *string `json:"confidence_level,omitempty"`

	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *Location `json:"location,omitempty"`
}

// Constants associated with the ContractType.ConfidenceLevel property.
// The confidence level in the identification of the termination date.
const (
	ContractType_ConfidenceLevel_High   = "High"
	ContractType_ConfidenceLevel_Low    = "Low"
	ContractType_ConfidenceLevel_Medium = "Medium"
)

// ConvertToHTMLOptions : The convertToHtml options.
type ConvertToHTMLOptions struct {

	// The document to convert.
	File *os.File `json:"file" validate:"required"`

	// The filename for file.
	Filename *string `json:"filename" validate:"required"`

	// The content type of file.
	FileContentType *string `json:"file_content_type,omitempty"`

	// The analysis model to be used by the service. For the **Element classification** and **Compare two documents**
	// methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults
	// apply to the standalone methods as well as to the methods' use in batch-processing requests.
	Model *string `json:"model,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the ConvertToHTMLOptions.Model property.
// The analysis model to be used by the service. For the **Element classification** and **Compare two documents**
// methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults apply
// to the standalone methods as well as to the methods' use in batch-processing requests.
const (
	ConvertToHTMLOptions_Model_Contracts = "contracts"
	ConvertToHTMLOptions_Model_Tables    = "tables"
)

// NewConvertToHTMLOptions : Instantiate ConvertToHTMLOptions
func (compareComply *CompareComplyV1) NewConvertToHTMLOptions(file *os.File, filename string) *ConvertToHTMLOptions {
	return &ConvertToHTMLOptions{
		File:     file,
		Filename: core.StringPtr(filename),
	}
}

// SetFile : Allow user to set File
func (options *ConvertToHTMLOptions) SetFile(file *os.File) *ConvertToHTMLOptions {
	options.File = file
	return options
}

// SetFilename : Allow user to set Filename
func (options *ConvertToHTMLOptions) SetFilename(filename string) *ConvertToHTMLOptions {
	options.Filename = core.StringPtr(filename)
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

// CreateBatchOptions : The createBatch options.
type CreateBatchOptions struct {

	// The Compare and Comply method to run across the submitted input documents.
	Function *string `json:"function" validate:"required"`

	// A JSON file containing the input Cloud Object Storage credentials. At a minimum, the credentials must enable `READ`
	// permissions on the bucket defined by the `input_bucket_name` parameter.
	InputCredentialsFile *os.File `json:"input_credentials_file" validate:"required"`

	// The geographical location of the Cloud Object Storage input bucket as listed on the **Endpoint** tab of your Cloud
	// Object Storage instance; for example, `us-geo`, `eu-geo`, or `ap-geo`.
	InputBucketLocation *string `json:"input_bucket_location" validate:"required"`

	// The name of the Cloud Object Storage input bucket.
	InputBucketName *string `json:"input_bucket_name" validate:"required"`

	// A JSON file that lists the Cloud Object Storage output credentials. At a minimum, the credentials must enable `READ`
	// and `WRITE` permissions on the bucket defined by the `output_bucket_name` parameter.
	OutputCredentialsFile *os.File `json:"output_credentials_file" validate:"required"`

	// The geographical location of the Cloud Object Storage output bucket as listed on the **Endpoint** tab of your Cloud
	// Object Storage instance; for example, `us-geo`, `eu-geo`, or `ap-geo`.
	OutputBucketLocation *string `json:"output_bucket_location" validate:"required"`

	// The name of the Cloud Object Storage output bucket.
	OutputBucketName *string `json:"output_bucket_name" validate:"required"`

	// The analysis model to be used by the service. For the **Element classification** and **Compare two documents**
	// methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults
	// apply to the standalone methods as well as to the methods' use in batch-processing requests.
	Model *string `json:"model,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the CreateBatchOptions.Function property.
// The Compare and Comply method to run across the submitted input documents.
const (
	CreateBatchOptions_Function_ElementClassification = "element_classification"
	CreateBatchOptions_Function_HTMLConversion        = "html_conversion"
	CreateBatchOptions_Function_Tables                = "tables"
)

// Constants associated with the CreateBatchOptions.Model property.
// The analysis model to be used by the service. For the **Element classification** and **Compare two documents**
// methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults apply
// to the standalone methods as well as to the methods' use in batch-processing requests.
const (
	CreateBatchOptions_Model_Contracts = "contracts"
	CreateBatchOptions_Model_Tables    = "tables"
)

// NewCreateBatchOptions : Instantiate CreateBatchOptions
func (compareComply *CompareComplyV1) NewCreateBatchOptions(function string, inputCredentialsFile *os.File, inputBucketLocation string, inputBucketName string, outputCredentialsFile *os.File, outputBucketLocation string, outputBucketName string) *CreateBatchOptions {
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
func (options *CreateBatchOptions) SetInputCredentialsFile(inputCredentialsFile *os.File) *CreateBatchOptions {
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
func (options *CreateBatchOptions) SetOutputCredentialsFile(outputCredentialsFile *os.File) *CreateBatchOptions {
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

// DeleteFeedbackOptions : The deleteFeedback options.
type DeleteFeedbackOptions struct {

	// A string that specifies the feedback entry to be deleted from the document.
	FeedbackID *string `json:"feedback_id" validate:"required"`

	// The analysis model to be used by the service. For the **Element classification** and **Compare two documents**
	// methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults
	// apply to the standalone methods as well as to the methods' use in batch-processing requests.
	Model *string `json:"model,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the DeleteFeedbackOptions.Model property.
// The analysis model to be used by the service. For the **Element classification** and **Compare two documents**
// methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults apply
// to the standalone methods as well as to the methods' use in batch-processing requests.
const (
	DeleteFeedbackOptions_Model_Contracts = "contracts"
	DeleteFeedbackOptions_Model_Tables    = "tables"
)

// NewDeleteFeedbackOptions : Instantiate DeleteFeedbackOptions
func (compareComply *CompareComplyV1) NewDeleteFeedbackOptions(feedbackID string) *DeleteFeedbackOptions {
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

// DocInfo : Information about the parsed input document.
type DocInfo struct {

	// The full text of the parsed document in HTML format.
	HTML *string `json:"html,omitempty"`

	// The title of the parsed document. If the service did not detect a title, the value of this element is `null`.
	Title *string `json:"title,omitempty"`

	// The MD5 hash of the input document.
	Hash *string `json:"hash,omitempty"`
}

// DocStructure : The structure of the input document.
type DocStructure struct {

	// An array containing one object per section or subsection identified in the input document.
	SectionTitles []SectionTitles `json:"section_titles,omitempty"`

	// An array containing one object per section or subsection, in parallel with the `section_titles` array, that details
	// the leading sentences in the corresponding section or subsection.
	LeadingSentences []LeadingSentence `json:"leading_sentences,omitempty"`
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

// EffectiveDates : An effective date.
type EffectiveDates struct {

	// The effective date, listed as a string.
	Text *string `json:"text,omitempty"`

	// The confidence level in the identification of the effective date.
	ConfidenceLevel *string `json:"confidence_level,omitempty"`

	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *Location `json:"location,omitempty"`
}

// Constants associated with the EffectiveDates.ConfidenceLevel property.
// The confidence level in the identification of the effective date.
const (
	EffectiveDates_ConfidenceLevel_High   = "High"
	EffectiveDates_ConfidenceLevel_Low    = "Low"
	EffectiveDates_ConfidenceLevel_Medium = "Medium"
)

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

// ElementLocations : A list of `begin` and `end` indexes that indicate the locations of the elements in the input document.
type ElementLocations struct {

	// An integer that indicates the starting position of the element in the input document.
	Begin *int64 `json:"begin,omitempty"`

	// An integer that indicates the ending position of the element in the input document.
	End *int64 `json:"end,omitempty"`
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

// ExtractTablesOptions : The extractTables options.
type ExtractTablesOptions struct {

	// The document on which to run table extraction.
	File *os.File `json:"file" validate:"required"`

	// The content type of file.
	FileContentType *string `json:"file_content_type,omitempty"`

	// The analysis model to be used by the service. For the **Element classification** and **Compare two documents**
	// methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults
	// apply to the standalone methods as well as to the methods' use in batch-processing requests.
	Model *string `json:"model,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the ExtractTablesOptions.Model property.
// The analysis model to be used by the service. For the **Element classification** and **Compare two documents**
// methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults apply
// to the standalone methods as well as to the methods' use in batch-processing requests.
const (
	ExtractTablesOptions_Model_Contracts = "contracts"
	ExtractTablesOptions_Model_Tables    = "tables"
)

// NewExtractTablesOptions : Instantiate ExtractTablesOptions
func (compareComply *CompareComplyV1) NewExtractTablesOptions(file *os.File) *ExtractTablesOptions {
	return &ExtractTablesOptions{
		File: file,
	}
}

// SetFile : Allow user to set File
func (options *ExtractTablesOptions) SetFile(file *os.File) *ExtractTablesOptions {
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

// FeedbackDataOutput : Information returned from the `POST /v1/feedback` method.
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

// FeedbackDeleted : The status and message of the deletion request.
type FeedbackDeleted struct {

	// HTTP return code.
	Status *int64 `json:"status,omitempty"`

	// Status message returned from the service.
	Message *string `json:"message,omitempty"`
}

// FeedbackList : The results of a successful `GET /v1/feedback` request.
type FeedbackList struct {

	// A list of all feedback for the document.
	Feedback []GetFeedback `json:"feedback,omitempty"`
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

	// Information returned from the `POST /v1/feedback` method.
	FeedbackData *FeedbackDataOutput `json:"feedback_data,omitempty"`
}

// GetBatchOptions : The getBatch options.
type GetBatchOptions struct {

	// The ID of the batch-processing job whose information you want to retrieve.
	BatchID *string `json:"batch_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetBatchOptions : Instantiate GetBatchOptions
func (compareComply *CompareComplyV1) NewGetBatchOptions(batchID string) *GetBatchOptions {
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

// GetFeedback : The results of a single feedback query.
type GetFeedback struct {

	// A string uniquely identifying the feedback entry.
	FeedbackID *string `json:"feedback_id,omitempty"`

	// A timestamp identifying the creation time of the feedback entry.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// A string containing the user's comment about the feedback entry.
	Comment *string `json:"comment,omitempty"`

	// Information returned from the `POST /v1/feedback` method.
	FeedbackData *FeedbackDataOutput `json:"feedback_data,omitempty"`
}

// GetFeedbackOptions : The getFeedback options.
type GetFeedbackOptions struct {

	// A string that specifies the feedback entry to be included in the output.
	FeedbackID *string `json:"feedback_id" validate:"required"`

	// The analysis model to be used by the service. For the **Element classification** and **Compare two documents**
	// methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults
	// apply to the standalone methods as well as to the methods' use in batch-processing requests.
	Model *string `json:"model,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the GetFeedbackOptions.Model property.
// The analysis model to be used by the service. For the **Element classification** and **Compare two documents**
// methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults apply
// to the standalone methods as well as to the methods' use in batch-processing requests.
const (
	GetFeedbackOptions_Model_Contracts = "contracts"
	GetFeedbackOptions_Model_Tables    = "tables"
)

// NewGetFeedbackOptions : Instantiate GetFeedbackOptions
func (compareComply *CompareComplyV1) NewGetFeedbackOptions(feedbackID string) *GetFeedbackOptions {
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

// KeyValuePair : Key-value pairs detected across cell boundaries.
type KeyValuePair struct {

	// A key in a key-value pair.
	Key *Key `json:"key,omitempty"`

	// A value in a key-value pair.
	Value *Value `json:"value,omitempty"`
}

// Label : A pair of `nature` and `party` objects. The `nature` object identifies the effect of the element on the identified
// `party`, and the `party` object identifies the affected party.
type Label struct {

	// The identified `nature` of the element.
	Nature *string `json:"nature" validate:"required"`

	// The identified `party` of the element.
	Party *string `json:"party" validate:"required"`
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

// ListBatchesOptions : The listBatches options.
type ListBatchesOptions struct {

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListBatchesOptions : Instantiate ListBatchesOptions
func (compareComply *CompareComplyV1) NewListBatchesOptions() *ListBatchesOptions {
	return &ListBatchesOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListBatchesOptions) SetHeaders(param map[string]string) *ListBatchesOptions {
	options.Headers = param
	return options
}

// ListFeedbackOptions : The listFeedback options.
type ListFeedbackOptions struct {

	// An optional string that filters the output to include only feedback with the specified feedback type. The only
	// permitted value is `element_classification`.
	FeedbackType *string `json:"feedback_type,omitempty"`

	// An optional string in the format `YYYY-MM-DD` that filters the output to include only feedback that was added before
	// the specified date.
	Before *strfmt.Date `json:"before,omitempty"`

	// An optional string in the format `YYYY-MM-DD` that filters the output to include only feedback that was added after
	// the specified date.
	After *strfmt.Date `json:"after,omitempty"`

	// An optional string that filters the output to include only feedback from the document with the specified
	// `document_title`.
	DocumentTitle *string `json:"document_title,omitempty"`

	// An optional string that filters the output to include only feedback with the specified `model_id`. The only
	// permitted value is `contracts`.
	ModelID *string `json:"model_id,omitempty"`

	// An optional string that filters the output to include only feedback with the specified `model_version`.
	ModelVersion *string `json:"model_version,omitempty"`

	// An optional string in the form of a comma-separated list of categories. If this is specified, the service filters
	// the output to include only feedback that has at least one category from the list removed.
	CategoryRemoved *string `json:"category_removed,omitempty"`

	// An optional string in the form of a comma-separated list of categories. If this is specified, the service filters
	// the output to include only feedback that has at least one category from the list added.
	CategoryAdded *string `json:"category_added,omitempty"`

	// An optional string in the form of a comma-separated list of categories. If this is specified, the service filters
	// the output to include only feedback that has at least one category from the list unchanged.
	CategoryNotChanged *string `json:"category_not_changed,omitempty"`

	// An optional string of comma-separated `nature`:`party` pairs. If this is specified, the service filters the output
	// to include only feedback that has at least one `nature`:`party` pair from the list removed.
	TypeRemoved *string `json:"type_removed,omitempty"`

	// An optional string of comma-separated `nature`:`party` pairs. If this is specified, the service filters the output
	// to include only feedback that has at least one `nature`:`party` pair from the list removed.
	TypeAdded *string `json:"type_added,omitempty"`

	// An optional string of comma-separated `nature`:`party` pairs. If this is specified, the service filters the output
	// to include only feedback that has at least one `nature`:`party` pair from the list unchanged.
	TypeNotChanged *string `json:"type_not_changed,omitempty"`

	// An optional integer specifying the number of documents that you want the service to return.
	PageLimit *int64 `json:"page_limit,omitempty"`

	// An optional string that returns the set of documents after the previous set. Use this parameter with the
	// `page_limit` parameter.
	Cursor *string `json:"cursor,omitempty"`

	// An optional comma-separated list of fields in the document to sort on. You can optionally specify the sort direction
	// by prefixing the value of the field with `-` for descending order or `+` for ascending order (the default).
	// Currently permitted sorting fields are `created`, `user_id`, and `document_title`.
	Sort *string `json:"sort,omitempty"`

	// An optional boolean value. If specified as `true`, the `pagination` object in the output includes a value called
	// `total` that gives the total count of feedback created.
	IncludeTotal *bool `json:"include_total,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListFeedbackOptions : Instantiate ListFeedbackOptions
func (compareComply *CompareComplyV1) NewListFeedbackOptions() *ListFeedbackOptions {
	return &ListFeedbackOptions{}
}

// SetFeedbackType : Allow user to set FeedbackType
func (options *ListFeedbackOptions) SetFeedbackType(feedbackType string) *ListFeedbackOptions {
	options.FeedbackType = core.StringPtr(feedbackType)
	return options
}

// SetBefore : Allow user to set Before
func (options *ListFeedbackOptions) SetBefore(before *strfmt.Date) *ListFeedbackOptions {
	options.Before = before
	return options
}

// SetAfter : Allow user to set After
func (options *ListFeedbackOptions) SetAfter(after *strfmt.Date) *ListFeedbackOptions {
	options.After = after
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

// OriginalLabelsIn : The original labeling from the input document, without the submitted feedback.
type OriginalLabelsIn struct {

	// Description of the action specified by the element and whom it affects.
	Types []TypeLabel `json:"types" validate:"required"`

	// List of functional categories into which the element falls; in other words, the subject matter of the element.
	Categories []Category `json:"categories" validate:"required"`
}

// OriginalLabelsOut : The original labeling from the input document, without the submitted feedback.
type OriginalLabelsOut struct {

	// Description of the action specified by the element and whom it affects.
	Types []TypeLabel `json:"types,omitempty"`

	// List of functional categories into which the element falls; in other words, the subject matter of the element.
	Categories []Category `json:"categories,omitempty"`

	// A string identifying the type of modification the feedback entry in the `updated_labels` array. Possible values are
	// `added`, `not_changed`, and `removed`.
	Modification *string `json:"modification,omitempty"`
}

// Constants associated with the OriginalLabelsOut.Modification property.
// A string identifying the type of modification the feedback entry in the `updated_labels` array. Possible values are
// `added`, `not_changed`, and `removed`.
const (
	OriginalLabelsOut_Modification_Added      = "added"
	OriginalLabelsOut_Modification_NotChanged = "not_changed"
	OriginalLabelsOut_Modification_Removed    = "removed"
)

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

// Parties : A party and its corresponding role, including address and contact information if identified.
type Parties struct {

	// A string identifying the party.
	Party *string `json:"party,omitempty"`

	// A string that identifies the importance of the party.
	Importance *string `json:"importance,omitempty"`

	// A string identifying the party's role.
	Role *string `json:"role,omitempty"`

	// List of the party's address or addresses.
	Addresses []Address `json:"addresses,omitempty"`

	// List of the names and roles of contacts identified in the input document.
	Contacts []Contact `json:"contacts,omitempty"`
}

// Constants associated with the Parties.Importance property.
// A string that identifies the importance of the party.
const (
	Parties_Importance_Primary = "Primary"
	Parties_Importance_Unknown = "Unknown"
)

// RowHeaderIds : An array of values, each being the `id` value of a row header that is applicable to this body cell.
type RowHeaderIds struct {

	// The `id` values of a row header.
	ID *string `json:"id,omitempty"`
}

// RowHeaderTexts : An array of values, each being the `text` value of a row header that is applicable to this body cell.
type RowHeaderTexts struct {

	// The `text` value of a row header.
	Text *string `json:"text,omitempty"`
}

// RowHeaderTextsNormalized : If you provide customization input, the normalized version of the row header texts according to the customization;
// otherwise, the same value as `row_header_texts`.
type RowHeaderTextsNormalized struct {

	// The normalized version of a row header text.
	TextNormalized *string `json:"text_normalized,omitempty"`
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

// SectionTitle : The table's section title, if identified.
type SectionTitle struct {

	// The text of the section title, if identified.
	Text *string `json:"text,omitempty"`

	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *Location `json:"location,omitempty"`
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

// ShortDoc : Brief information about the input document.
type ShortDoc struct {

	// The title of the input document, if identified.
	Title *string `json:"title,omitempty"`

	// The MD5 hash of the input document.
	Hash *string `json:"hash,omitempty"`
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

// Tables : The contents of the tables extracted from a document.
type Tables struct {

	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *Location `json:"location,omitempty"`

	// The textual contents of the current table from the input document without associated markup content.
	Text *string `json:"text,omitempty"`

	// The table's section title, if identified.
	SectionTitle *SectionTitle `json:"section_title,omitempty"`

	// An array of table-level cells that apply as headers to all the other cells in the current table.
	TableHeaders []TableHeaders `json:"table_headers,omitempty"`

	// An array of row-level cells, each applicable as a header to other cells in the same row as itself, of the current
	// table.
	RowHeaders []RowHeaders `json:"row_headers,omitempty"`

	// An array of column-level cells, each applicable as a header to other cells in the same column as itself, of the
	// current table.
	ColumnHeaders []ColumnHeaders `json:"column_headers,omitempty"`

	// An array of key-value pairs identified in the current table.
	KeyValuePairs []KeyValuePair `json:"key_value_pairs,omitempty"`

	// An array of cells that are neither table header nor column header nor row header cells, of the current table with
	// corresponding row and column header associations.
	BodyCells []BodyCells `json:"body_cells,omitempty"`
}

// TerminationDates : Termination dates identified in the input document.
type TerminationDates struct {

	// The termination date.
	Text *string `json:"text,omitempty"`

	// The confidence level in the identification of the termination date.
	ConfidenceLevel *string `json:"confidence_level,omitempty"`

	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *Location `json:"location,omitempty"`
}

// Constants associated with the TerminationDates.ConfidenceLevel property.
// The confidence level in the identification of the termination date.
const (
	TerminationDates_ConfidenceLevel_High   = "High"
	TerminationDates_ConfidenceLevel_Low    = "Low"
	TerminationDates_ConfidenceLevel_Medium = "Medium"
)

// TypeLabel : Identification of a specific type.
type TypeLabel struct {

	// A pair of `nature` and `party` objects. The `nature` object identifies the effect of the element on the identified
	// `party`, and the `party` object identifies the affected party.
	Label *Label `json:"label,omitempty"`

	// One or more hash values that you can send to IBM to provide feedback or receive support.
	ProvenanceIds []string `json:"provenance_ids,omitempty"`
}

// TypeLabelComparison : Identification of a specific type.
type TypeLabelComparison struct {

	// A pair of `nature` and `party` objects. The `nature` object identifies the effect of the element on the identified
	// `party`, and the `party` object identifies the affected party.
	Label *Label `json:"label,omitempty"`
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

// UpdateBatchOptions : The updateBatch options.
type UpdateBatchOptions struct {

	// The ID of the batch-processing job you want to update.
	BatchID *string `json:"batch_id" validate:"required"`

	// The action you want to perform on the specified batch-processing job.
	Action *string `json:"action" validate:"required"`

	// The analysis model to be used by the service. For the **Element classification** and **Compare two documents**
	// methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults
	// apply to the standalone methods as well as to the methods' use in batch-processing requests.
	Model *string `json:"model,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the UpdateBatchOptions.Action property.
// The action you want to perform on the specified batch-processing job.
const (
	UpdateBatchOptions_Action_Cancel = "cancel"
	UpdateBatchOptions_Action_Rescan = "rescan"
)

// Constants associated with the UpdateBatchOptions.Model property.
// The analysis model to be used by the service. For the **Element classification** and **Compare two documents**
// methods, the default is `contracts`. For the **Extract tables** method, the default is `tables`. These defaults apply
// to the standalone methods as well as to the methods' use in batch-processing requests.
const (
	UpdateBatchOptions_Model_Contracts = "contracts"
	UpdateBatchOptions_Model_Tables    = "tables"
)

// NewUpdateBatchOptions : Instantiate UpdateBatchOptions
func (compareComply *CompareComplyV1) NewUpdateBatchOptions(batchID string, action string) *UpdateBatchOptions {
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

// UpdatedLabelsOut : The updated labeling from the input document, accounting for the submitted feedback.
type UpdatedLabelsOut struct {

	// Description of the action specified by the element and whom it affects.
	Types []TypeLabel `json:"types,omitempty"`

	// List of functional categories into which the element falls; in other words, the subject matter of the element.
	Categories []Category `json:"categories,omitempty"`

	// The type of modification the feedback entry in the `updated_labels` array. Possible values are `added`,
	// `not_changed`, and `removed`.
	Modification *string `json:"modification,omitempty"`
}

// Constants associated with the UpdatedLabelsOut.Modification property.
// The type of modification the feedback entry in the `updated_labels` array. Possible values are `added`,
// `not_changed`, and `removed`.
const (
	UpdatedLabelsOut_Modification_Added      = "added"
	UpdatedLabelsOut_Modification_NotChanged = "not_changed"
	UpdatedLabelsOut_Modification_Removed    = "removed"
)

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
