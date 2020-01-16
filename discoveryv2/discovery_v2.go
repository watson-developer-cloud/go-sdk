/**
 * (C) Copyright IBM Corp. 2019, 2020.
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

// Package discoveryv2 : Operations and models for the DiscoveryV2 service
package discoveryv2

import (
	"fmt"
	"github.com/IBM/go-sdk-core/core"
	"github.com/go-openapi/strfmt"
	common "github.com/watson-developer-cloud/go-sdk/common"
	"io"
	"strings"
)

// DiscoveryV2 : IBM Watson&trade; Discovery for IBM Cloud Pak for Data is a cognitive search and content analytics
// engine that you can add to applications to identify patterns, trends and actionable insights to drive better
// decision-making. Securely unify structured and unstructured data with pre-enriched content, and use a simplified
// query language to eliminate the need for manual filtering of results.
//
// Version: 2.0
// See: https://cloud.ibm.com/docs/services/discovery-data/
type DiscoveryV2 struct {
	Service *core.BaseService
	Version string
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = ""

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "discovery"

// DiscoveryV2Options : Service options
type DiscoveryV2Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
	Version       string
}

// NewDiscoveryV2 : constructs an instance of DiscoveryV2 with passed in options.
func NewDiscoveryV2(options *DiscoveryV2Options) (service *DiscoveryV2, err error) {
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
		baseService.SetServiceURL(options.URL)
	}

	service = &DiscoveryV2{
		Service: baseService,
		Version: options.Version,
	}

	return
}

// SetServiceURL sets the service URL
func (discovery *DiscoveryV2) SetServiceURL(url string) error {
	return discovery.Service.SetServiceURL(url)
}

// DisableSSLVerification bypasses verification of the server's SSL certificate
func (discovery *DiscoveryV2) DisableSSLVerification() {
	discovery.Service.DisableSSLVerification()
}

// ListCollections : List collections
// Lists existing collections for the specified project.
func (discovery *DiscoveryV2) ListCollections(listCollectionsOptions *ListCollectionsOptions) (result *ListCollectionsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listCollectionsOptions, "listCollectionsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listCollectionsOptions, "listCollectionsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/projects", "collections"}
	pathParameters := []string{*listCollectionsOptions.ProjectID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listCollectionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V2", "ListCollections")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", discovery.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = discovery.Service.Request(request, new(ListCollectionsResponse))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*ListCollectionsResponse)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// Query : Query a project
// By using this method, you can construct queries. For details, see the [Discovery
// documentation](https://cloud.ibm.com/docs/services/discovery-data?topic=discovery-data-query-concepts).
func (discovery *DiscoveryV2) Query(queryOptions *QueryOptions) (result *QueryResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(queryOptions, "queryOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(queryOptions, "queryOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/projects", "query"}
	pathParameters := []string{*queryOptions.ProjectID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range queryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V2", "Query")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", discovery.Version)

	body := make(map[string]interface{})
	if queryOptions.CollectionIds != nil {
		body["collection_ids"] = queryOptions.CollectionIds
	}
	if queryOptions.Filter != nil {
		body["filter"] = queryOptions.Filter
	}
	if queryOptions.Query != nil {
		body["query"] = queryOptions.Query
	}
	if queryOptions.NaturalLanguageQuery != nil {
		body["natural_language_query"] = queryOptions.NaturalLanguageQuery
	}
	if queryOptions.Aggregation != nil {
		body["aggregation"] = queryOptions.Aggregation
	}
	if queryOptions.Count != nil {
		body["count"] = queryOptions.Count
	}
	if queryOptions.Return != nil {
		body["return"] = queryOptions.Return
	}
	if queryOptions.Offset != nil {
		body["offset"] = queryOptions.Offset
	}
	if queryOptions.Sort != nil {
		body["sort"] = queryOptions.Sort
	}
	if queryOptions.Highlight != nil {
		body["highlight"] = queryOptions.Highlight
	}
	if queryOptions.SpellingSuggestions != nil {
		body["spelling_suggestions"] = queryOptions.SpellingSuggestions
	}
	if queryOptions.TableResults != nil {
		body["table_results"] = queryOptions.TableResults
	}
	if queryOptions.SuggestedRefinements != nil {
		body["suggested_refinements"] = queryOptions.SuggestedRefinements
	}
	if queryOptions.Passages != nil {
		body["passages"] = queryOptions.Passages
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = discovery.Service.Request(request, new(QueryResponse))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*QueryResponse)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// GetAutocompletion : Get Autocomplete Suggestions
// Returns completion query suggestions for the specified prefix.
func (discovery *DiscoveryV2) GetAutocompletion(getAutocompletionOptions *GetAutocompletionOptions) (result *Completions, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAutocompletionOptions, "getAutocompletionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAutocompletionOptions, "getAutocompletionOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/projects", "autocompletion"}
	pathParameters := []string{*getAutocompletionOptions.ProjectID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAutocompletionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V2", "GetAutocompletion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("prefix", fmt.Sprint(*getAutocompletionOptions.Prefix))
	if getAutocompletionOptions.CollectionIds != nil {
		builder.AddQuery("collection_ids", strings.Join(getAutocompletionOptions.CollectionIds, ","))
	}
	if getAutocompletionOptions.Field != nil {
		builder.AddQuery("field", fmt.Sprint(*getAutocompletionOptions.Field))
	}
	if getAutocompletionOptions.Count != nil {
		builder.AddQuery("count", fmt.Sprint(*getAutocompletionOptions.Count))
	}
	builder.AddQuery("version", discovery.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = discovery.Service.Request(request, new(Completions))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*Completions)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// QueryNotices : Query system notices
// Queries for notices (errors or warnings) that might have been generated by the system. Notices are generated when
// ingesting documents and performing relevance training.
func (discovery *DiscoveryV2) QueryNotices(queryNoticesOptions *QueryNoticesOptions) (result *QueryNoticesResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(queryNoticesOptions, "queryNoticesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(queryNoticesOptions, "queryNoticesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/projects", "notices"}
	pathParameters := []string{*queryNoticesOptions.ProjectID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range queryNoticesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V2", "QueryNotices")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")

	if queryNoticesOptions.Filter != nil {
		builder.AddQuery("filter", fmt.Sprint(*queryNoticesOptions.Filter))
	}
	if queryNoticesOptions.Query != nil {
		builder.AddQuery("query", fmt.Sprint(*queryNoticesOptions.Query))
	}
	if queryNoticesOptions.NaturalLanguageQuery != nil {
		builder.AddQuery("natural_language_query", fmt.Sprint(*queryNoticesOptions.NaturalLanguageQuery))
	}
	if queryNoticesOptions.Count != nil {
		builder.AddQuery("count", fmt.Sprint(*queryNoticesOptions.Count))
	}
	if queryNoticesOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*queryNoticesOptions.Offset))
	}
	builder.AddQuery("version", discovery.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = discovery.Service.Request(request, new(QueryNoticesResponse))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*QueryNoticesResponse)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// ListFields : List fields
// Gets a list of the unique fields (and their types) stored in the the specified collections.
func (discovery *DiscoveryV2) ListFields(listFieldsOptions *ListFieldsOptions) (result *ListFieldsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listFieldsOptions, "listFieldsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listFieldsOptions, "listFieldsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/projects", "fields"}
	pathParameters := []string{*listFieldsOptions.ProjectID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listFieldsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V2", "ListFields")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")

	if listFieldsOptions.CollectionIds != nil {
		builder.AddQuery("collection_ids", strings.Join(listFieldsOptions.CollectionIds, ","))
	}
	builder.AddQuery("version", discovery.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = discovery.Service.Request(request, new(ListFieldsResponse))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*ListFieldsResponse)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// GetComponentSettings : Configuration settings for components
// Returns default configuration settings for components.
func (discovery *DiscoveryV2) GetComponentSettings(getComponentSettingsOptions *GetComponentSettingsOptions) (result *ComponentSettingsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getComponentSettingsOptions, "getComponentSettingsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getComponentSettingsOptions, "getComponentSettingsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/projects", "component_settings"}
	pathParameters := []string{*getComponentSettingsOptions.ProjectID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getComponentSettingsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V2", "GetComponentSettings")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", discovery.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = discovery.Service.Request(request, new(ComponentSettingsResponse))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*ComponentSettingsResponse)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// AddDocument : Add a document
// Add a document to a collection with optional metadata.
//
//  Returns immediately after the system has accepted the document for processing.
//
//   * The user must provide document content, metadata, or both. If the request is missing both document content and
// metadata, it is rejected.
//
//   * The user can set the **Content-Type** parameter on the **file** part to indicate the media type of the document.
// If the **Content-Type** parameter is missing or is one of the generic media types (for example,
// `application/octet-stream`), then the service attempts to automatically detect the document's media type.
//
//   * The following field names are reserved and will be filtered out if present after normalization: `id`, `score`,
// `highlight`, and any field with the prefix of: `_`, `+`, or `-`
//
//   * Fields with empty name values after normalization are filtered out before indexing.
//
//   * Fields containing the following characters after normalization are filtered out before indexing: `#` and `,`
//
//   If the document is uploaded to a collection that has it's data shared with another collection, the
// **X-Watson-Discovery-Force** header must be set to `true`.
//
//  **Note:** Documents can be added with a specific **document_id** by using the
// **_/v2/projects/{project_id}/collections/{collection_id}/documents** method.
//
// **Note:** This operation only works on collections created to accept direct file uploads. It cannot be used to modify
// a collection that connects to an external source such as Microsoft SharePoint.
func (discovery *DiscoveryV2) AddDocument(addDocumentOptions *AddDocumentOptions) (result *DocumentAccepted, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addDocumentOptions, "addDocumentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addDocumentOptions, "addDocumentOptions")
	if err != nil {
		return
	}
	if (addDocumentOptions.File == nil) && (addDocumentOptions.Metadata == nil) {
		err = fmt.Errorf("At least one of file or metadata must be supplied")
		return
	}

	pathSegments := []string{"v2/projects", "collections", "documents"}
	pathParameters := []string{*addDocumentOptions.ProjectID, *addDocumentOptions.CollectionID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range addDocumentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V2", "AddDocument")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	if addDocumentOptions.XWatsonDiscoveryForce != nil {
		builder.AddHeader("X-Watson-Discovery-Force", fmt.Sprint(*addDocumentOptions.XWatsonDiscoveryForce))
	}
	builder.AddQuery("version", discovery.Version)

	if addDocumentOptions.File != nil {
		builder.AddFormData("file", core.StringNilMapper(addDocumentOptions.Filename),
			core.StringNilMapper(addDocumentOptions.FileContentType), addDocumentOptions.File)
	}
	if addDocumentOptions.Metadata != nil {
		builder.AddFormData("metadata", "", "", fmt.Sprint(*addDocumentOptions.Metadata))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = discovery.Service.Request(request, new(DocumentAccepted))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*DocumentAccepted)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// UpdateDocument : Update a document
// Replace an existing document or add a document with a specified **document_id**. Starts ingesting a document with
// optional metadata.
//
// If the document is uploaded to a collection that has it's data shared with another collection, the
// **X-Watson-Discovery-Force** header must be set to `true`.
//
// **Note:** When uploading a new document with this method it automatically replaces any document stored with the same
// **document_id** if it exists.
//
// **Note:** This operation only works on collections created to accept direct file uploads. It cannot be used to modify
// a collection that connects to an external source such as Microsoft SharePoint.
func (discovery *DiscoveryV2) UpdateDocument(updateDocumentOptions *UpdateDocumentOptions) (result *DocumentAccepted, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateDocumentOptions, "updateDocumentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateDocumentOptions, "updateDocumentOptions")
	if err != nil {
		return
	}
	if (updateDocumentOptions.File == nil) && (updateDocumentOptions.Metadata == nil) {
		err = fmt.Errorf("At least one of file or metadata must be supplied")
		return
	}

	pathSegments := []string{"v2/projects", "collections", "documents"}
	pathParameters := []string{*updateDocumentOptions.ProjectID, *updateDocumentOptions.CollectionID, *updateDocumentOptions.DocumentID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateDocumentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V2", "UpdateDocument")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	if updateDocumentOptions.XWatsonDiscoveryForce != nil {
		builder.AddHeader("X-Watson-Discovery-Force", fmt.Sprint(*updateDocumentOptions.XWatsonDiscoveryForce))
	}
	builder.AddQuery("version", discovery.Version)

	if updateDocumentOptions.File != nil {
		builder.AddFormData("file", core.StringNilMapper(updateDocumentOptions.Filename),
			core.StringNilMapper(updateDocumentOptions.FileContentType), updateDocumentOptions.File)
	}
	if updateDocumentOptions.Metadata != nil {
		builder.AddFormData("metadata", "", "", fmt.Sprint(*updateDocumentOptions.Metadata))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = discovery.Service.Request(request, new(DocumentAccepted))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*DocumentAccepted)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// DeleteDocument : Delete a document
// If the given document ID is invalid, or if the document is not found, then the a success response is returned (HTTP
// status code `200`) with the status set to 'deleted'.
//
// **Note:** This operation only works on collections created to accept direct file uploads. It cannot be used to modify
// a collection that connects to an external source such as Microsoft SharePoint.
func (discovery *DiscoveryV2) DeleteDocument(deleteDocumentOptions *DeleteDocumentOptions) (result *DeleteDocumentResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteDocumentOptions, "deleteDocumentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteDocumentOptions, "deleteDocumentOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/projects", "collections", "documents"}
	pathParameters := []string{*deleteDocumentOptions.ProjectID, *deleteDocumentOptions.CollectionID, *deleteDocumentOptions.DocumentID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteDocumentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V2", "DeleteDocument")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	if deleteDocumentOptions.XWatsonDiscoveryForce != nil {
		builder.AddHeader("X-Watson-Discovery-Force", fmt.Sprint(*deleteDocumentOptions.XWatsonDiscoveryForce))
	}
	builder.AddQuery("version", discovery.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = discovery.Service.Request(request, new(DeleteDocumentResponse))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*DeleteDocumentResponse)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// ListTrainingQueries : List training queries
// List the training queries for the specified project.
func (discovery *DiscoveryV2) ListTrainingQueries(listTrainingQueriesOptions *ListTrainingQueriesOptions) (result *TrainingQuerySet, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listTrainingQueriesOptions, "listTrainingQueriesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listTrainingQueriesOptions, "listTrainingQueriesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/projects", "training_data/queries"}
	pathParameters := []string{*listTrainingQueriesOptions.ProjectID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listTrainingQueriesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V2", "ListTrainingQueries")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", discovery.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = discovery.Service.Request(request, new(TrainingQuerySet))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*TrainingQuerySet)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// DeleteTrainingQueries : Delete training queries
// Removes all training queries for the specified project.
func (discovery *DiscoveryV2) DeleteTrainingQueries(deleteTrainingQueriesOptions *DeleteTrainingQueriesOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteTrainingQueriesOptions, "deleteTrainingQueriesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteTrainingQueriesOptions, "deleteTrainingQueriesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/projects", "training_data/queries"}
	pathParameters := []string{*deleteTrainingQueriesOptions.ProjectID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteTrainingQueriesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V2", "DeleteTrainingQueries")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("version", discovery.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = discovery.Service.Request(request, nil)

	return
}

// CreateTrainingQuery : Create training query
// Add a query to the training data for this project. The query can contain a filter and natural language query.
func (discovery *DiscoveryV2) CreateTrainingQuery(createTrainingQueryOptions *CreateTrainingQueryOptions) (result *TrainingQuery, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createTrainingQueryOptions, "createTrainingQueryOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createTrainingQueryOptions, "createTrainingQueryOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/projects", "training_data/queries"}
	pathParameters := []string{*createTrainingQueryOptions.ProjectID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createTrainingQueryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V2", "CreateTrainingQuery")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", discovery.Version)

	body := make(map[string]interface{})
	if createTrainingQueryOptions.NaturalLanguageQuery != nil {
		body["natural_language_query"] = createTrainingQueryOptions.NaturalLanguageQuery
	}
	if createTrainingQueryOptions.Examples != nil {
		body["examples"] = createTrainingQueryOptions.Examples
	}
	if createTrainingQueryOptions.Filter != nil {
		body["filter"] = createTrainingQueryOptions.Filter
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = discovery.Service.Request(request, new(TrainingQuery))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*TrainingQuery)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// GetTrainingQuery : Get a training data query
// Get details for a specific training data query, including the query string and all examples.
func (discovery *DiscoveryV2) GetTrainingQuery(getTrainingQueryOptions *GetTrainingQueryOptions) (result *TrainingQuery, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getTrainingQueryOptions, "getTrainingQueryOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getTrainingQueryOptions, "getTrainingQueryOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/projects", "training_data/queries"}
	pathParameters := []string{*getTrainingQueryOptions.ProjectID, *getTrainingQueryOptions.QueryID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getTrainingQueryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V2", "GetTrainingQuery")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", discovery.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = discovery.Service.Request(request, new(TrainingQuery))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*TrainingQuery)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// UpdateTrainingQuery : Update a training query
// Updates an existing training query and it's examples.
func (discovery *DiscoveryV2) UpdateTrainingQuery(updateTrainingQueryOptions *UpdateTrainingQueryOptions) (result *TrainingQuery, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateTrainingQueryOptions, "updateTrainingQueryOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateTrainingQueryOptions, "updateTrainingQueryOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/projects", "training_data/queries"}
	pathParameters := []string{*updateTrainingQueryOptions.ProjectID, *updateTrainingQueryOptions.QueryID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateTrainingQueryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V2", "UpdateTrainingQuery")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", discovery.Version)

	body := make(map[string]interface{})
	if updateTrainingQueryOptions.NaturalLanguageQuery != nil {
		body["natural_language_query"] = updateTrainingQueryOptions.NaturalLanguageQuery
	}
	if updateTrainingQueryOptions.Examples != nil {
		body["examples"] = updateTrainingQueryOptions.Examples
	}
	if updateTrainingQueryOptions.Filter != nil {
		body["filter"] = updateTrainingQueryOptions.Filter
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = discovery.Service.Request(request, new(TrainingQuery))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*TrainingQuery)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// AddDocumentOptions : The AddDocument options.
type AddDocumentOptions struct {

	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required"`

	// The ID of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// The content of the document to ingest. The maximum supported file size when adding a file to a collection is 50
	// megabytes, the maximum supported file size when testing a configuration is 1 megabyte. Files larger than the
	// supported size are rejected.
	File io.ReadCloser `json:"file,omitempty"`

	// The filename for file.
	Filename *string `json:"filename,omitempty"`

	// The content type of file.
	FileContentType *string `json:"file_content_type,omitempty"`

	// The maximum supported metadata file size is 1 MB. Metadata parts larger than 1 MB are rejected. Example:  ``` {
	//   "Creator": "Johnny Appleseed",
	//   "Subject": "Apples"
	// } ```.
	Metadata *string `json:"metadata,omitempty"`

	// When `true`, the uploaded document is added to the collection even if the data for that collection is shared with
	// other collections.
	XWatsonDiscoveryForce *bool `json:"X-Watson-Discovery-Force,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewAddDocumentOptions : Instantiate AddDocumentOptions
func (discovery *DiscoveryV2) NewAddDocumentOptions(projectID string, collectionID string) *AddDocumentOptions {
	return &AddDocumentOptions{
		ProjectID:    core.StringPtr(projectID),
		CollectionID: core.StringPtr(collectionID),
	}
}

// SetProjectID : Allow user to set ProjectID
func (options *AddDocumentOptions) SetProjectID(projectID string) *AddDocumentOptions {
	options.ProjectID = core.StringPtr(projectID)
	return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *AddDocumentOptions) SetCollectionID(collectionID string) *AddDocumentOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
}

// SetFile : Allow user to set File
func (options *AddDocumentOptions) SetFile(file io.ReadCloser) *AddDocumentOptions {
	options.File = file
	return options
}

// SetFilename : Allow user to set Filename
func (options *AddDocumentOptions) SetFilename(filename string) *AddDocumentOptions {
	options.Filename = core.StringPtr(filename)
	return options
}

// SetFileContentType : Allow user to set FileContentType
func (options *AddDocumentOptions) SetFileContentType(fileContentType string) *AddDocumentOptions {
	options.FileContentType = core.StringPtr(fileContentType)
	return options
}

// SetMetadata : Allow user to set Metadata
func (options *AddDocumentOptions) SetMetadata(metadata string) *AddDocumentOptions {
	options.Metadata = core.StringPtr(metadata)
	return options
}

// SetXWatsonDiscoveryForce : Allow user to set XWatsonDiscoveryForce
func (options *AddDocumentOptions) SetXWatsonDiscoveryForce(xWatsonDiscoveryForce bool) *AddDocumentOptions {
	options.XWatsonDiscoveryForce = core.BoolPtr(xWatsonDiscoveryForce)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *AddDocumentOptions) SetHeaders(param map[string]string) *AddDocumentOptions {
	options.Headers = param
	return options
}

// Collection : A collection for storing documents.
type Collection struct {

	// The unique identifier of the collection.
	CollectionID *string `json:"collection_id,omitempty"`

	// The name of the collection.
	Name *string `json:"name,omitempty"`
}

// Completions : An object containing an array of autocompletion suggestions.
type Completions struct {

	// Array of autcomplete suggestion based on the provided prefix.
	Completions []string `json:"completions,omitempty"`
}

// ComponentSettingsAggregation : Display settings for aggregations.
type ComponentSettingsAggregation struct {

	// Identifier used to map aggregation settings to aggregation configuration.
	Name *string `json:"name,omitempty"`

	// User-friendly alias for the aggregation.
	Label *string `json:"label,omitempty"`

	// Whether users is allowed to select more than one of the aggregation terms.
	MultipleSelectionsAllowed *bool `json:"multiple_selections_allowed,omitempty"`

	// Type of visualization to use when rendering the aggregation.
	VisualizationType *string `json:"visualization_type,omitempty"`
}

// Constants associated with the ComponentSettingsAggregation.VisualizationType property.
// Type of visualization to use when rendering the aggregation.
const (
	ComponentSettingsAggregation_VisualizationType_Auto       = "auto"
	ComponentSettingsAggregation_VisualizationType_FacetTable = "facet_table"
	ComponentSettingsAggregation_VisualizationType_Map        = "map"
	ComponentSettingsAggregation_VisualizationType_WordCloud  = "word_cloud"
)

// ComponentSettingsFieldsShown : Fields shown in the results section of the UI.
type ComponentSettingsFieldsShown struct {

	// Body label.
	Body *ComponentSettingsFieldsShownBody `json:"body,omitempty"`

	// Title label.
	Title *ComponentSettingsFieldsShownTitle `json:"title,omitempty"`
}

// ComponentSettingsFieldsShownBody : Body label.
type ComponentSettingsFieldsShownBody struct {

	// Use the whole passage as the body.
	UsePassage *bool `json:"use_passage,omitempty"`

	// Use a specific field as the title.
	Field *string `json:"field,omitempty"`
}

// ComponentSettingsFieldsShownTitle : Title label.
type ComponentSettingsFieldsShownTitle struct {

	// Use a specific field as the title.
	Field *string `json:"field,omitempty"`
}

// ComponentSettingsResponse : A response containing the default component settings.
type ComponentSettingsResponse struct {

	// Fields shown in the results section of the UI.
	FieldsShown *ComponentSettingsFieldsShown `json:"fields_shown,omitempty"`

	// Whether or not autocomplete is enabled.
	Autocomplete *bool `json:"autocomplete,omitempty"`

	// Whether or not structured search is enabled.
	StructuredSearch *bool `json:"structured_search,omitempty"`

	// Number or results shown per page.
	ResultsPerPage *int64 `json:"results_per_page,omitempty"`

	// a list of component setting aggregations.
	Aggregations []ComponentSettingsAggregation `json:"aggregations,omitempty"`
}

// CreateTrainingQueryOptions : The CreateTrainingQuery options.
type CreateTrainingQueryOptions struct {

	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required"`

	// The natural text query for the training query.
	NaturalLanguageQuery *string `json:"natural_language_query" validate:"required"`

	// Array of training examples.
	Examples []TrainingExample `json:"examples" validate:"required"`

	// The filter used on the collection before the **natural_language_query** is applied.
	Filter *string `json:"filter,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewCreateTrainingQueryOptions : Instantiate CreateTrainingQueryOptions
func (discovery *DiscoveryV2) NewCreateTrainingQueryOptions(projectID string, naturalLanguageQuery string, examples []TrainingExample) *CreateTrainingQueryOptions {
	return &CreateTrainingQueryOptions{
		ProjectID:            core.StringPtr(projectID),
		NaturalLanguageQuery: core.StringPtr(naturalLanguageQuery),
		Examples:             examples,
	}
}

// SetProjectID : Allow user to set ProjectID
func (options *CreateTrainingQueryOptions) SetProjectID(projectID string) *CreateTrainingQueryOptions {
	options.ProjectID = core.StringPtr(projectID)
	return options
}

// SetNaturalLanguageQuery : Allow user to set NaturalLanguageQuery
func (options *CreateTrainingQueryOptions) SetNaturalLanguageQuery(naturalLanguageQuery string) *CreateTrainingQueryOptions {
	options.NaturalLanguageQuery = core.StringPtr(naturalLanguageQuery)
	return options
}

// SetExamples : Allow user to set Examples
func (options *CreateTrainingQueryOptions) SetExamples(examples []TrainingExample) *CreateTrainingQueryOptions {
	options.Examples = examples
	return options
}

// SetFilter : Allow user to set Filter
func (options *CreateTrainingQueryOptions) SetFilter(filter string) *CreateTrainingQueryOptions {
	options.Filter = core.StringPtr(filter)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateTrainingQueryOptions) SetHeaders(param map[string]string) *CreateTrainingQueryOptions {
	options.Headers = param
	return options
}

// DeleteDocumentOptions : The DeleteDocument options.
type DeleteDocumentOptions struct {

	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required"`

	// The ID of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// The ID of the document.
	DocumentID *string `json:"document_id" validate:"required"`

	// When `true`, the uploaded document is added to the collection even if the data for that collection is shared with
	// other collections.
	XWatsonDiscoveryForce *bool `json:"X-Watson-Discovery-Force,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteDocumentOptions : Instantiate DeleteDocumentOptions
func (discovery *DiscoveryV2) NewDeleteDocumentOptions(projectID string, collectionID string, documentID string) *DeleteDocumentOptions {
	return &DeleteDocumentOptions{
		ProjectID:    core.StringPtr(projectID),
		CollectionID: core.StringPtr(collectionID),
		DocumentID:   core.StringPtr(documentID),
	}
}

// SetProjectID : Allow user to set ProjectID
func (options *DeleteDocumentOptions) SetProjectID(projectID string) *DeleteDocumentOptions {
	options.ProjectID = core.StringPtr(projectID)
	return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *DeleteDocumentOptions) SetCollectionID(collectionID string) *DeleteDocumentOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
}

// SetDocumentID : Allow user to set DocumentID
func (options *DeleteDocumentOptions) SetDocumentID(documentID string) *DeleteDocumentOptions {
	options.DocumentID = core.StringPtr(documentID)
	return options
}

// SetXWatsonDiscoveryForce : Allow user to set XWatsonDiscoveryForce
func (options *DeleteDocumentOptions) SetXWatsonDiscoveryForce(xWatsonDiscoveryForce bool) *DeleteDocumentOptions {
	options.XWatsonDiscoveryForce = core.BoolPtr(xWatsonDiscoveryForce)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteDocumentOptions) SetHeaders(param map[string]string) *DeleteDocumentOptions {
	options.Headers = param
	return options
}

// DeleteDocumentResponse : Information returned when a document is deleted.
type DeleteDocumentResponse struct {

	// The unique identifier of the document.
	DocumentID *string `json:"document_id,omitempty"`

	// Status of the document. A deleted document has the status deleted.
	Status *string `json:"status,omitempty"`
}

// Constants associated with the DeleteDocumentResponse.Status property.
// Status of the document. A deleted document has the status deleted.
const (
	DeleteDocumentResponse_Status_Deleted = "deleted"
)

// DeleteTrainingQueriesOptions : The DeleteTrainingQueries options.
type DeleteTrainingQueriesOptions struct {

	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteTrainingQueriesOptions : Instantiate DeleteTrainingQueriesOptions
func (discovery *DiscoveryV2) NewDeleteTrainingQueriesOptions(projectID string) *DeleteTrainingQueriesOptions {
	return &DeleteTrainingQueriesOptions{
		ProjectID: core.StringPtr(projectID),
	}
}

// SetProjectID : Allow user to set ProjectID
func (options *DeleteTrainingQueriesOptions) SetProjectID(projectID string) *DeleteTrainingQueriesOptions {
	options.ProjectID = core.StringPtr(projectID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteTrainingQueriesOptions) SetHeaders(param map[string]string) *DeleteTrainingQueriesOptions {
	options.Headers = param
	return options
}

// DocumentAccepted : Information returned after an uploaded document is accepted.
type DocumentAccepted struct {

	// The unique identifier of the ingested document.
	DocumentID *string `json:"document_id,omitempty"`

	// Status of the document in the ingestion process. A status of `processing` is returned for documents that are
	// ingested with a *version* date before `2019-01-01`. The `pending` status is returned for all others.
	Status *string `json:"status,omitempty"`
}

// Constants associated with the DocumentAccepted.Status property.
// Status of the document in the ingestion process. A status of `processing` is returned for documents that are ingested
// with a *version* date before `2019-01-01`. The `pending` status is returned for all others.
const (
	DocumentAccepted_Status_Pending    = "pending"
	DocumentAccepted_Status_Processing = "processing"
)

// DocumentAttribute : List of document attributes.
type DocumentAttribute struct {

	// The type of attribute.
	Type *string `json:"type,omitempty"`

	// The text associated with the attribute.
	Text *string `json:"text,omitempty"`

	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *TableElementLocation `json:"location,omitempty"`
}

// Field : Object containing field details.
type Field struct {

	// The name of the field.
	Field *string `json:"field,omitempty"`

	// The type of the field.
	Type *string `json:"type,omitempty"`

	// The collection Id of the collection where the field was found.
	CollectionID *string `json:"collection_id,omitempty"`
}

// Constants associated with the Field.Type property.
// The type of the field.
const (
	Field_Type_Binary  = "binary"
	Field_Type_Boolean = "boolean"
	Field_Type_Byte    = "byte"
	Field_Type_Date    = "date"
	Field_Type_Double  = "double"
	Field_Type_Float   = "float"
	Field_Type_Integer = "integer"
	Field_Type_Long    = "long"
	Field_Type_Nested  = "nested"
	Field_Type_Short   = "short"
	Field_Type_String  = "string"
)

// GetAutocompletionOptions : The GetAutocompletion options.
type GetAutocompletionOptions struct {

	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required"`

	// The prefix to use for autocompletion. For example, the prefix `Ho` could autocomplete to `Hot`, `Housing`, or `How
	// do I upgrade`. Possible completions are.
	Prefix *string `json:"prefix" validate:"required"`

	// Comma separated list of the collection IDs. If this parameter is not specified, all collections in the project are
	// used.
	CollectionIds []string `json:"collection_ids,omitempty"`

	// The field in the result documents that autocompletion suggestions are identified from.
	Field *string `json:"field,omitempty"`

	// The number of autocompletion suggestions to return.
	Count *int64 `json:"count,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetAutocompletionOptions : Instantiate GetAutocompletionOptions
func (discovery *DiscoveryV2) NewGetAutocompletionOptions(projectID string, prefix string) *GetAutocompletionOptions {
	return &GetAutocompletionOptions{
		ProjectID: core.StringPtr(projectID),
		Prefix:    core.StringPtr(prefix),
	}
}

// SetProjectID : Allow user to set ProjectID
func (options *GetAutocompletionOptions) SetProjectID(projectID string) *GetAutocompletionOptions {
	options.ProjectID = core.StringPtr(projectID)
	return options
}

// SetPrefix : Allow user to set Prefix
func (options *GetAutocompletionOptions) SetPrefix(prefix string) *GetAutocompletionOptions {
	options.Prefix = core.StringPtr(prefix)
	return options
}

// SetCollectionIds : Allow user to set CollectionIds
func (options *GetAutocompletionOptions) SetCollectionIds(collectionIds []string) *GetAutocompletionOptions {
	options.CollectionIds = collectionIds
	return options
}

// SetField : Allow user to set Field
func (options *GetAutocompletionOptions) SetField(field string) *GetAutocompletionOptions {
	options.Field = core.StringPtr(field)
	return options
}

// SetCount : Allow user to set Count
func (options *GetAutocompletionOptions) SetCount(count int64) *GetAutocompletionOptions {
	options.Count = core.Int64Ptr(count)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetAutocompletionOptions) SetHeaders(param map[string]string) *GetAutocompletionOptions {
	options.Headers = param
	return options
}

// GetComponentSettingsOptions : The GetComponentSettings options.
type GetComponentSettingsOptions struct {

	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetComponentSettingsOptions : Instantiate GetComponentSettingsOptions
func (discovery *DiscoveryV2) NewGetComponentSettingsOptions(projectID string) *GetComponentSettingsOptions {
	return &GetComponentSettingsOptions{
		ProjectID: core.StringPtr(projectID),
	}
}

// SetProjectID : Allow user to set ProjectID
func (options *GetComponentSettingsOptions) SetProjectID(projectID string) *GetComponentSettingsOptions {
	options.ProjectID = core.StringPtr(projectID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetComponentSettingsOptions) SetHeaders(param map[string]string) *GetComponentSettingsOptions {
	options.Headers = param
	return options
}

// GetTrainingQueryOptions : The GetTrainingQuery options.
type GetTrainingQueryOptions struct {

	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required"`

	// The ID of the query used for training.
	QueryID *string `json:"query_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetTrainingQueryOptions : Instantiate GetTrainingQueryOptions
func (discovery *DiscoveryV2) NewGetTrainingQueryOptions(projectID string, queryID string) *GetTrainingQueryOptions {
	return &GetTrainingQueryOptions{
		ProjectID: core.StringPtr(projectID),
		QueryID:   core.StringPtr(queryID),
	}
}

// SetProjectID : Allow user to set ProjectID
func (options *GetTrainingQueryOptions) SetProjectID(projectID string) *GetTrainingQueryOptions {
	options.ProjectID = core.StringPtr(projectID)
	return options
}

// SetQueryID : Allow user to set QueryID
func (options *GetTrainingQueryOptions) SetQueryID(queryID string) *GetTrainingQueryOptions {
	options.QueryID = core.StringPtr(queryID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetTrainingQueryOptions) SetHeaders(param map[string]string) *GetTrainingQueryOptions {
	options.Headers = param
	return options
}

// ListCollectionsOptions : The ListCollections options.
type ListCollectionsOptions struct {

	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListCollectionsOptions : Instantiate ListCollectionsOptions
func (discovery *DiscoveryV2) NewListCollectionsOptions(projectID string) *ListCollectionsOptions {
	return &ListCollectionsOptions{
		ProjectID: core.StringPtr(projectID),
	}
}

// SetProjectID : Allow user to set ProjectID
func (options *ListCollectionsOptions) SetProjectID(projectID string) *ListCollectionsOptions {
	options.ProjectID = core.StringPtr(projectID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListCollectionsOptions) SetHeaders(param map[string]string) *ListCollectionsOptions {
	options.Headers = param
	return options
}

// ListCollectionsResponse : Response object containing an array of collection details.
type ListCollectionsResponse struct {

	// An array containing information about each collection in the project.
	Collections []Collection `json:"collections,omitempty"`
}

// ListFieldsOptions : The ListFields options.
type ListFieldsOptions struct {

	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required"`

	// Comma separated list of the collection IDs. If this parameter is not specified, all collections in the project are
	// used.
	CollectionIds []string `json:"collection_ids,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListFieldsOptions : Instantiate ListFieldsOptions
func (discovery *DiscoveryV2) NewListFieldsOptions(projectID string) *ListFieldsOptions {
	return &ListFieldsOptions{
		ProjectID: core.StringPtr(projectID),
	}
}

// SetProjectID : Allow user to set ProjectID
func (options *ListFieldsOptions) SetProjectID(projectID string) *ListFieldsOptions {
	options.ProjectID = core.StringPtr(projectID)
	return options
}

// SetCollectionIds : Allow user to set CollectionIds
func (options *ListFieldsOptions) SetCollectionIds(collectionIds []string) *ListFieldsOptions {
	options.CollectionIds = collectionIds
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListFieldsOptions) SetHeaders(param map[string]string) *ListFieldsOptions {
	options.Headers = param
	return options
}

// ListFieldsResponse : The list of fetched fields.
//
// The fields are returned using a fully qualified name format, however, the format differs slightly from that used by
// the query operations.
//
//   * Fields which contain nested objects are assigned a type of "nested".
//
//   * Fields which belong to a nested object are prefixed with `.properties` (for example,
// `warnings.properties.severity` means that the `warnings` object has a property called `severity`).
type ListFieldsResponse struct {

	// An array containing information about each field in the collections.
	Fields []Field `json:"fields,omitempty"`
}

// ListTrainingQueriesOptions : The ListTrainingQueries options.
type ListTrainingQueriesOptions struct {

	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListTrainingQueriesOptions : Instantiate ListTrainingQueriesOptions
func (discovery *DiscoveryV2) NewListTrainingQueriesOptions(projectID string) *ListTrainingQueriesOptions {
	return &ListTrainingQueriesOptions{
		ProjectID: core.StringPtr(projectID),
	}
}

// SetProjectID : Allow user to set ProjectID
func (options *ListTrainingQueriesOptions) SetProjectID(projectID string) *ListTrainingQueriesOptions {
	options.ProjectID = core.StringPtr(projectID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListTrainingQueriesOptions) SetHeaders(param map[string]string) *ListTrainingQueriesOptions {
	options.Headers = param
	return options
}

// Notice : A notice produced for the collection.
type Notice struct {

	// Identifies the notice. Many notices might have the same ID. This field exists so that user applications can
	// programmatically identify a notice and take automatic corrective action. Typical notice IDs include: `index_failed`,
	// `index_failed_too_many_requests`, `index_failed_incompatible_field`, `index_failed_cluster_unavailable`,
	// `ingestion_timeout`, `ingestion_error`, `bad_request`, `internal_error`, `missing_model`, `unsupported_model`,
	// `smart_document_understanding_failed_incompatible_field`, `smart_document_understanding_failed_internal_error`,
	// `smart_document_understanding_failed_internal_error`, `smart_document_understanding_failed_warning`,
	// `smart_document_understanding_page_error`, `smart_document_understanding_page_warning`. **Note:** This is not a
	// complete list, other values might be returned.
	NoticeID *string `json:"notice_id,omitempty"`

	// The creation date of the collection in the format yyyy-MM-dd'T'HH:mm:ss.SSS'Z'.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// Unique identifier of the document.
	DocumentID *string `json:"document_id,omitempty"`

	// Unique identifier of the collection.
	CollectionID *string `json:"collection_id,omitempty"`

	// Unique identifier of the query used for relevance training.
	QueryID *string `json:"query_id,omitempty"`

	// Severity level of the notice.
	Severity *string `json:"severity,omitempty"`

	// Ingestion or training step in which the notice occurred.
	Step *string `json:"step,omitempty"`

	// The description of the notice.
	Description *string `json:"description,omitempty"`
}

// Constants associated with the Notice.Severity property.
// Severity level of the notice.
const (
	Notice_Severity_Error   = "error"
	Notice_Severity_Warning = "warning"
)

// QueryAggregation : An abstract aggregation type produced by Discovery to analyze the input provided.
type QueryAggregation struct {

	// The type of aggregation command used. Options include: term, histogram, timeslice, nested, filter, min, max, sum,
	// average, unique_count, and top_hits.
	Type *string `json:"type" validate:"required"`
}

// QueryHistogramAggregationResult : Histogram numeric interval result.
type QueryHistogramAggregationResult struct {

	// The value of the upper bound for the numeric segment.
	Key *int64 `json:"key" validate:"required"`

	// Number of documents with the specified key as the upper bound.
	MatchingResults *int64 `json:"matching_results" validate:"required"`

	// An array of sub aggregations.
	Aggregations []QueryAggregation `json:"aggregations,omitempty"`
}

// QueryLargePassages : Configuration for passage retrieval.
type QueryLargePassages struct {

	// A passages query that returns the most relevant passages from the results.
	Enabled *bool `json:"enabled,omitempty"`

	// When `true`, passages will be returned within their respective result.
	PerDocument *bool `json:"per_document,omitempty"`

	// Maximum number of passages to return per result.
	MaxPerDocument *int64 `json:"max_per_document,omitempty"`

	// A list of fields that passages are drawn from. If this parameter not specified, then all top-level fields are
	// included.
	Fields []string `json:"fields,omitempty"`

	// The maximum number of passages to return. The search returns fewer passages if the requested total is not found. The
	// default is `10`. The maximum is `100`.
	Count *int64 `json:"count,omitempty"`

	// The approximate number of characters that any one passage will have.
	Characters *int64 `json:"characters,omitempty"`
}

// QueryLargeSuggestedRefinements : Configuration for suggested refinements.
type QueryLargeSuggestedRefinements struct {

	// Whether to perform suggested refinements.
	Enabled *bool `json:"enabled,omitempty"`

	// Maximum number of suggested refinements texts to be returned. The default is `10`. The maximum is `100`.
	Count *int64 `json:"count,omitempty"`
}

// QueryLargeTableResults : Configuration for table retrieval.
type QueryLargeTableResults struct {

	// Whether to enable table retrieval.
	Enabled *bool `json:"enabled,omitempty"`

	// Maximum number of tables to return.
	Count *int64 `json:"count,omitempty"`
}

// QueryNoticesOptions : The QueryNotices options.
type QueryNoticesOptions struct {

	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required"`

	// A cacheable query that excludes documents that don't mention the query content. Filter searches are better for
	// metadata-type searches and for assessing the concepts in the data set.
	Filter *string `json:"filter,omitempty"`

	// A query search returns all documents in your data set with full enrichments and full text, but with the most
	// relevant documents listed first.
	Query *string `json:"query,omitempty"`

	// A natural language query that returns relevant documents by utilizing training data and natural language
	// understanding.
	NaturalLanguageQuery *string `json:"natural_language_query,omitempty"`

	// Number of results to return. The maximum for the **count** and **offset** values together in any one query is
	// **10000**.
	Count *int64 `json:"count,omitempty"`

	// The number of query results to skip at the beginning. For example, if the total number of results that are returned
	// is 10 and the offset is 8, it returns the last two results. The maximum for the **count** and **offset** values
	// together in any one query is **10000**.
	Offset *int64 `json:"offset,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewQueryNoticesOptions : Instantiate QueryNoticesOptions
func (discovery *DiscoveryV2) NewQueryNoticesOptions(projectID string) *QueryNoticesOptions {
	return &QueryNoticesOptions{
		ProjectID: core.StringPtr(projectID),
	}
}

// SetProjectID : Allow user to set ProjectID
func (options *QueryNoticesOptions) SetProjectID(projectID string) *QueryNoticesOptions {
	options.ProjectID = core.StringPtr(projectID)
	return options
}

// SetFilter : Allow user to set Filter
func (options *QueryNoticesOptions) SetFilter(filter string) *QueryNoticesOptions {
	options.Filter = core.StringPtr(filter)
	return options
}

// SetQuery : Allow user to set Query
func (options *QueryNoticesOptions) SetQuery(query string) *QueryNoticesOptions {
	options.Query = core.StringPtr(query)
	return options
}

// SetNaturalLanguageQuery : Allow user to set NaturalLanguageQuery
func (options *QueryNoticesOptions) SetNaturalLanguageQuery(naturalLanguageQuery string) *QueryNoticesOptions {
	options.NaturalLanguageQuery = core.StringPtr(naturalLanguageQuery)
	return options
}

// SetCount : Allow user to set Count
func (options *QueryNoticesOptions) SetCount(count int64) *QueryNoticesOptions {
	options.Count = core.Int64Ptr(count)
	return options
}

// SetOffset : Allow user to set Offset
func (options *QueryNoticesOptions) SetOffset(offset int64) *QueryNoticesOptions {
	options.Offset = core.Int64Ptr(offset)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *QueryNoticesOptions) SetHeaders(param map[string]string) *QueryNoticesOptions {
	options.Headers = param
	return options
}

// QueryNoticesResponse : Object containing notice query results.
type QueryNoticesResponse struct {

	// The number of matching results.
	MatchingResults *int64 `json:"matching_results,omitempty"`

	// Array of document results that match the query.
	Notices []Notice `json:"notices,omitempty"`
}

// QueryOptions : The Query options.
type QueryOptions struct {

	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required"`

	// A comma-separated list of collection IDs to be queried against.
	CollectionIds []string `json:"collection_ids,omitempty"`

	// A cacheable query that excludes documents that don't mention the query content. Filter searches are better for
	// metadata-type searches and for assessing the concepts in the data set.
	Filter *string `json:"filter,omitempty"`

	// A query search returns all documents in your data set with full enrichments and full text, but with the most
	// relevant documents listed first. Use a query search when you want to find the most relevant search results.
	Query *string `json:"query,omitempty"`

	// A natural language query that returns relevant documents by utilizing training data and natural language
	// understanding.
	NaturalLanguageQuery *string `json:"natural_language_query,omitempty"`

	// An aggregation search that returns an exact answer by combining query search with filters. Useful for applications
	// to build lists, tables, and time series. For a full list of possible aggregations, see the Query reference.
	Aggregation *string `json:"aggregation,omitempty"`

	// Number of results to return.
	Count *int64 `json:"count,omitempty"`

	// A list of the fields in the document hierarchy to return. If this parameter not specified, then all top-level fields
	// are returned.
	Return []string `json:"return,omitempty"`

	// The number of query results to skip at the beginning. For example, if the total number of results that are returned
	// is 10 and the offset is 8, it returns the last two results.
	Offset *int64 `json:"offset,omitempty"`

	// A comma-separated list of fields in the document to sort on. You can optionally specify a sort direction by
	// prefixing the field with `-` for descending or `+` for ascending. Ascending is the default sort direction if no
	// prefix is specified. This parameter cannot be used in the same query as the **bias** parameter.
	Sort *string `json:"sort,omitempty"`

	// When `true`, a highlight field is returned for each result which contains the fields which match the query with
	// `<em></em>` tags around the matching query terms.
	Highlight *bool `json:"highlight,omitempty"`

	// When `true` and the **natural_language_query** parameter is used, the **natural_language_query** parameter is spell
	// checked. The most likely correction is returned in the **suggested_query** field of the response (if one exists).
	SpellingSuggestions *bool `json:"spelling_suggestions,omitempty"`

	// Configuration for table retrieval.
	TableResults *QueryLargeTableResults `json:"table_results,omitempty"`

	// Configuration for suggested refinements.
	SuggestedRefinements *QueryLargeSuggestedRefinements `json:"suggested_refinements,omitempty"`

	// Configuration for passage retrieval.
	Passages *QueryLargePassages `json:"passages,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewQueryOptions : Instantiate QueryOptions
func (discovery *DiscoveryV2) NewQueryOptions(projectID string) *QueryOptions {
	return &QueryOptions{
		ProjectID: core.StringPtr(projectID),
	}
}

// SetProjectID : Allow user to set ProjectID
func (options *QueryOptions) SetProjectID(projectID string) *QueryOptions {
	options.ProjectID = core.StringPtr(projectID)
	return options
}

// SetCollectionIds : Allow user to set CollectionIds
func (options *QueryOptions) SetCollectionIds(collectionIds []string) *QueryOptions {
	options.CollectionIds = collectionIds
	return options
}

// SetFilter : Allow user to set Filter
func (options *QueryOptions) SetFilter(filter string) *QueryOptions {
	options.Filter = core.StringPtr(filter)
	return options
}

// SetQuery : Allow user to set Query
func (options *QueryOptions) SetQuery(query string) *QueryOptions {
	options.Query = core.StringPtr(query)
	return options
}

// SetNaturalLanguageQuery : Allow user to set NaturalLanguageQuery
func (options *QueryOptions) SetNaturalLanguageQuery(naturalLanguageQuery string) *QueryOptions {
	options.NaturalLanguageQuery = core.StringPtr(naturalLanguageQuery)
	return options
}

// SetAggregation : Allow user to set Aggregation
func (options *QueryOptions) SetAggregation(aggregation string) *QueryOptions {
	options.Aggregation = core.StringPtr(aggregation)
	return options
}

// SetCount : Allow user to set Count
func (options *QueryOptions) SetCount(count int64) *QueryOptions {
	options.Count = core.Int64Ptr(count)
	return options
}

// SetReturn : Allow user to set Return
func (options *QueryOptions) SetReturn(returnVar []string) *QueryOptions {
	options.Return = returnVar
	return options
}

// SetOffset : Allow user to set Offset
func (options *QueryOptions) SetOffset(offset int64) *QueryOptions {
	options.Offset = core.Int64Ptr(offset)
	return options
}

// SetSort : Allow user to set Sort
func (options *QueryOptions) SetSort(sort string) *QueryOptions {
	options.Sort = core.StringPtr(sort)
	return options
}

// SetHighlight : Allow user to set Highlight
func (options *QueryOptions) SetHighlight(highlight bool) *QueryOptions {
	options.Highlight = core.BoolPtr(highlight)
	return options
}

// SetSpellingSuggestions : Allow user to set SpellingSuggestions
func (options *QueryOptions) SetSpellingSuggestions(spellingSuggestions bool) *QueryOptions {
	options.SpellingSuggestions = core.BoolPtr(spellingSuggestions)
	return options
}

// SetTableResults : Allow user to set TableResults
func (options *QueryOptions) SetTableResults(tableResults *QueryLargeTableResults) *QueryOptions {
	options.TableResults = tableResults
	return options
}

// SetSuggestedRefinements : Allow user to set SuggestedRefinements
func (options *QueryOptions) SetSuggestedRefinements(suggestedRefinements *QueryLargeSuggestedRefinements) *QueryOptions {
	options.SuggestedRefinements = suggestedRefinements
	return options
}

// SetPassages : Allow user to set Passages
func (options *QueryOptions) SetPassages(passages *QueryLargePassages) *QueryOptions {
	options.Passages = passages
	return options
}

// SetHeaders : Allow user to set Headers
func (options *QueryOptions) SetHeaders(param map[string]string) *QueryOptions {
	options.Headers = param
	return options
}

// QueryResponse : A response containing the documents and aggregations for the query.
type QueryResponse struct {

	// The number of matching results for the query.
	MatchingResults *int64 `json:"matching_results,omitempty"`

	// Array of document results for the query.
	Results []QueryResult `json:"results,omitempty"`

	// Array of aggregations for the query.
	Aggregations []QueryAggregation `json:"aggregations,omitempty"`

	// An object contain retrieval type information.
	RetrievalDetails *RetrievalDetails `json:"retrieval_details,omitempty"`

	// Suggested correction to the submitted **natural_language_query** value.
	SuggestedQuery *string `json:"suggested_query,omitempty"`

	// Array of suggested refinements.
	SuggestedRefinements []QuerySuggestedRefinement `json:"suggested_refinements,omitempty"`

	// Array of table results.
	TableResults []QueryTableResult `json:"table_results,omitempty"`
}

// QueryResult : Result document for the specified query.
type QueryResult map[string]interface{}

// SetDocumentID : Allow user to set DocumentID
func (this *QueryResult) SetDocumentID(DocumentID *string) {
	(*this)["document_id"] = DocumentID
}

// GetDocumentID : Allow user to get DocumentID
func (this *QueryResult) GetDocumentID() *string {
	return (*this)["document_id"].(*string)
}

// SetMetadata : Allow user to set Metadata
func (this *QueryResult) SetMetadata(Metadata *map[string]interface{}) {
	(*this)["metadata"] = Metadata
}

// GetMetadata : Allow user to get Metadata
func (this *QueryResult) GetMetadata() *map[string]interface{} {
	return (*this)["metadata"].(*map[string]interface{})
}

// SetResultMetadata : Allow user to set ResultMetadata
func (this *QueryResult) SetResultMetadata(ResultMetadata *QueryResultMetadata) {
	(*this)["result_metadata"] = ResultMetadata
}

// GetResultMetadata : Allow user to get ResultMetadata
func (this *QueryResult) GetResultMetadata() *QueryResultMetadata {
	return (*this)["result_metadata"].(*QueryResultMetadata)
}

// SetDocumentPassages : Allow user to set DocumentPassages
func (this *QueryResult) SetDocumentPassages(DocumentPassages *[]QueryResultPassage) {
	(*this)["document_passages"] = DocumentPassages
}

// GetDocumentPassages : Allow user to get DocumentPassages
func (this *QueryResult) GetDocumentPassages() *[]QueryResultPassage {
	return (*this)["document_passages"].(*[]QueryResultPassage)
}

// SetProperty : Allow user to set arbitrary property
func (this *QueryResult) SetProperty(Key string, Value *interface{}) {
	(*this)[Key] = Value
}

// GetProperty : Allow user to get arbitrary property
func (this *QueryResult) GetProperty(Key string) *interface{} {
	return (*this)[Key].(*interface{})
}

// QueryResultMetadata : Metadata of a query result.
type QueryResultMetadata struct {

	// The document retrieval source that produced this search result.
	DocumentRetrievalSource *string `json:"document_retrieval_source,omitempty"`

	// The collection id associated with this training data set.
	CollectionID *string `json:"collection_id" validate:"required"`

	// The confidence score for the given result. Calculated based on how relevant the result is estimated to be.
	// confidence can range from `0.0` to `1.0`. The higher the number, the more relevant the document. The `confidence`
	// value for a result was calculated using the model specified in the `document_retrieval_strategy` field of the result
	// set. This field is only returned if the **natural_language_query** parameter is specified in the query.
	Confidence *float64 `json:"confidence,omitempty"`
}

// Constants associated with the QueryResultMetadata.DocumentRetrievalSource property.
// The document retrieval source that produced this search result.
const (
	QueryResultMetadata_DocumentRetrievalSource_Curation = "curation"
	QueryResultMetadata_DocumentRetrievalSource_Search   = "search"
)

// QueryResultPassage : A passage query result.
type QueryResultPassage struct {

	// The content of the extracted passage.
	PassageText *string `json:"passage_text,omitempty"`

	// The position of the first character of the extracted passage in the originating field.
	StartOffset *int64 `json:"start_offset,omitempty"`

	// The position of the last character of the extracted passage in the originating field.
	EndOffset *int64 `json:"end_offset,omitempty"`

	// The label of the field from which the passage has been extracted.
	Field *string `json:"field,omitempty"`
}

// QuerySuggestedRefinement : A suggested additional query term or terms user to filter results.
type QuerySuggestedRefinement struct {

	// The text used to filter.
	Text *string `json:"text,omitempty"`
}

// QueryTableResult : A tables whose content or context match a search query.
type QueryTableResult struct {

	// The identifier for the retrieved table.
	TableID *string `json:"table_id,omitempty"`

	// The identifier of the document the table was retrieved from.
	SourceDocumentID *string `json:"source_document_id,omitempty"`

	// The identifier of the collection the table was retrieved from.
	CollectionID *string `json:"collection_id,omitempty"`

	// HTML snippet of the table info.
	TableHTML *string `json:"table_html,omitempty"`

	// The offset of the table html snippet in the original document html.
	TableHTMLOffset *int64 `json:"table_html_offset,omitempty"`

	// Full table object retrieved from Table Understanding Enrichment.
	Table *TableResultTable `json:"table,omitempty"`
}

// QueryTermAggregationResult : Top value result for the term aggregation.
type QueryTermAggregationResult struct {

	// Value of the field with a non-zero frequency in the document set.
	Key *string `json:"key" validate:"required"`

	// Number of documents containing the 'key'.
	MatchingResults *int64 `json:"matching_results" validate:"required"`

	// An array of sub aggregations.
	Aggregations []QueryAggregation `json:"aggregations,omitempty"`
}

// QueryTimesliceAggregationResult : A timeslice interval segment.
type QueryTimesliceAggregationResult struct {

	// String date value of the upper bound for the timeslice interval in ISO-8601 format.
	KeyAsString *string `json:"key_as_string" validate:"required"`

	// Numeric date value of the upper bound for the timeslice interval in UNIX milliseconds since epoch.
	Key *int64 `json:"key" validate:"required"`

	// Number of documents with the specified key as the upper bound.
	MatchingResults *int64 `json:"matching_results" validate:"required"`

	// An array of sub aggregations.
	Aggregations []QueryAggregation `json:"aggregations,omitempty"`
}

// QueryTopHitsAggregationResult : A query response containing the matching documents for the preceding aggregations.
type QueryTopHitsAggregationResult struct {

	// Number of matching results.
	MatchingResults *int64 `json:"matching_results" validate:"required"`

	// An array of the document results.
	Hits []map[string]interface{} `json:"hits,omitempty"`
}

// RetrievalDetails : An object contain retrieval type information.
type RetrievalDetails struct {

	// Identifies the document retrieval strategy used for this query. `relevancy_training` indicates that the results were
	// returned using a relevancy trained model.
	//
	//  **Note**: In the event of trained collections being queried, but the trained model is not used to return results,
	// the **document_retrieval_strategy** will be listed as `untrained`.
	DocumentRetrievalStrategy *string `json:"document_retrieval_strategy,omitempty"`
}

// Constants associated with the RetrievalDetails.DocumentRetrievalStrategy property.
// Identifies the document retrieval strategy used for this query. `relevancy_training` indicates that the results were
// returned using a relevancy trained model.
//
//  **Note**: In the event of trained collections being queried, but the trained model is not used to return results,
// the **document_retrieval_strategy** will be listed as `untrained`.
const (
	RetrievalDetails_DocumentRetrievalStrategy_RelevancyTraining = "relevancy_training"
	RetrievalDetails_DocumentRetrievalStrategy_Untrained         = "untrained"
)

// TableBodyCells : Cells that are not table header, column header, or row header cells.
type TableBodyCells struct {

	// The unique ID of the cell in the current table.
	CellID *string `json:"cell_id,omitempty"`

	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *TableElementLocation `json:"location,omitempty"`

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

	// A list of table row header ids.
	RowHeaderIds []TableRowHeaderIds `json:"row_header_ids,omitempty"`

	// A list of table row header texts.
	RowHeaderTexts []TableRowHeaderTexts `json:"row_header_texts,omitempty"`

	// A list of table row header texts normalized.
	RowHeaderTextsNormalized []TableRowHeaderTextsNormalized `json:"row_header_texts_normalized,omitempty"`

	// A list of table column header ids.
	ColumnHeaderIds []TableColumnHeaderIds `json:"column_header_ids,omitempty"`

	// A list of table column header texts.
	ColumnHeaderTexts []TableColumnHeaderTexts `json:"column_header_texts,omitempty"`

	// A list of table column header texts normalized.
	ColumnHeaderTextsNormalized []TableColumnHeaderTextsNormalized `json:"column_header_texts_normalized,omitempty"`

	// A list of document attributes.
	Attributes []DocumentAttribute `json:"attributes,omitempty"`
}

// TableCellKey : A key in a key-value pair.
type TableCellKey struct {

	// The unique ID of the key in the table.
	CellID *string `json:"cell_id,omitempty"`

	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *TableElementLocation `json:"location,omitempty"`

	// The text content of the table cell without HTML markup.
	Text *string `json:"text,omitempty"`
}

// TableCellValues : A value in a key-value pair.
type TableCellValues struct {

	// The unique ID of the value in the table.
	CellID *string `json:"cell_id,omitempty"`

	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *TableElementLocation `json:"location,omitempty"`

	// The text content of the table cell without HTML markup.
	Text *string `json:"text,omitempty"`
}

// TableColumnHeaderIds : An array of values, each being the `id` value of a column header that is applicable to the current cell.
type TableColumnHeaderIds struct {

	// The `id` value of a column header.
	ID *string `json:"id,omitempty"`
}

// TableColumnHeaderTexts : An array of values, each being the `text` value of a column header that is applicable to the current cell.
type TableColumnHeaderTexts struct {

	// The `text` value of a column header.
	Text *string `json:"text,omitempty"`
}

// TableColumnHeaderTextsNormalized : If you provide customization input, the normalized version of the column header texts according to the customization;
// otherwise, the same value as `column_header_texts`.
type TableColumnHeaderTextsNormalized struct {

	// The normalized version of a column header text.
	TextNormalized *string `json:"text_normalized,omitempty"`
}

// TableColumnHeaders : Column-level cells, each applicable as a header to other cells in the same column as itself, of the current table.
type TableColumnHeaders struct {

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

// TableElementLocation : The numeric location of the identified element in the document, represented with two integers labeled `begin` and
// `end`.
type TableElementLocation struct {

	// The element's `begin` index.
	Begin *int64 `json:"begin" validate:"required"`

	// The element's `end` index.
	End *int64 `json:"end" validate:"required"`
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

// TableKeyValuePairs : Key-value pairs detected across cell boundaries.
type TableKeyValuePairs struct {

	// A key in a key-value pair.
	Key *TableCellKey `json:"key,omitempty"`

	// A list of values in a key-value pair.
	Value []TableCellValues `json:"value,omitempty"`
}

// TableResultTable : Full table object retrieved from Table Understanding Enrichment.
type TableResultTable struct {

	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *TableElementLocation `json:"location,omitempty"`

	// The textual contents of the current table from the input document without associated markup content.
	Text *string `json:"text,omitempty"`

	// Text and associated location within a table.
	SectionTitle *TableTextLocation `json:"section_title,omitempty"`

	// Text and associated location within a table.
	Title *TableTextLocation `json:"title,omitempty"`

	// An array of table-level cells that apply as headers to all the other cells in the current table.
	TableHeaders []TableHeaders `json:"table_headers,omitempty"`

	// An array of row-level cells, each applicable as a header to other cells in the same row as itself, of the current
	// table.
	RowHeaders []TableRowHeaders `json:"row_headers,omitempty"`

	// An array of column-level cells, each applicable as a header to other cells in the same column as itself, of the
	// current table.
	ColumnHeaders []TableColumnHeaders `json:"column_headers,omitempty"`

	// An array of key-value pairs identified in the current table.
	KeyValuePairs []TableKeyValuePairs `json:"key_value_pairs,omitempty"`

	// An array of cells that are neither table header nor column header nor row header cells, of the current table with
	// corresponding row and column header associations.
	BodyCells []TableBodyCells `json:"body_cells,omitempty"`

	// An array of lists of textual entries across the document related to the current table being parsed.
	Contexts []TableTextLocation `json:"contexts,omitempty"`
}

// TableRowHeaderIds : An array of values, each being the `id` value of a row header that is applicable to this body cell.
type TableRowHeaderIds struct {

	// The `id` values of a row header.
	ID *string `json:"id,omitempty"`
}

// TableRowHeaderTexts : An array of values, each being the `text` value of a row header that is applicable to this body cell.
type TableRowHeaderTexts struct {

	// The `text` value of a row header.
	Text *string `json:"text,omitempty"`
}

// TableRowHeaderTextsNormalized : If you provide customization input, the normalized version of the row header texts according to the customization;
// otherwise, the same value as `row_header_texts`.
type TableRowHeaderTextsNormalized struct {

	// The normalized version of a row header text.
	TextNormalized *string `json:"text_normalized,omitempty"`
}

// TableRowHeaders : Row-level cells, each applicable as a header to other cells in the same row as itself, of the current table.
type TableRowHeaders struct {

	// The unique ID of the cell in the current table.
	CellID *string `json:"cell_id,omitempty"`

	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *TableElementLocation `json:"location,omitempty"`

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

// TableTextLocation : Text and associated location within a table.
type TableTextLocation struct {

	// The text retrieved.
	Text *string `json:"text,omitempty"`

	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *TableElementLocation `json:"location,omitempty"`
}

// TrainingExample : Object containing example response details for a training query.
type TrainingExample struct {

	// The document ID associated with this training example.
	DocumentID *string `json:"document_id" validate:"required"`

	// The collection ID associated with this training example.
	CollectionID *string `json:"collection_id" validate:"required"`

	// The relevance of the training example.
	Relevance *int64 `json:"relevance" validate:"required"`

	// The date and time the example was created.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// The date and time the example was updated.
	Updated *strfmt.DateTime `json:"updated,omitempty"`
}

// NewTrainingExample : Instantiate TrainingExample (Generic Model Constructor)
func (discovery *DiscoveryV2) NewTrainingExample(documentID string, collectionID string, relevance int64) (model *TrainingExample, err error) {
	model = &TrainingExample{
		DocumentID:   core.StringPtr(documentID),
		CollectionID: core.StringPtr(collectionID),
		Relevance:    core.Int64Ptr(relevance),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// TrainingQuery : Object containing training query details.
type TrainingQuery struct {

	// The query ID associated with the training query.
	QueryID *string `json:"query_id,omitempty"`

	// The natural text query for the training query.
	NaturalLanguageQuery *string `json:"natural_language_query" validate:"required"`

	// The filter used on the collection before the **natural_language_query** is applied.
	Filter *string `json:"filter,omitempty"`

	// The date and time the query was created.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// The date and time the query was updated.
	Updated *strfmt.DateTime `json:"updated,omitempty"`

	// Array of training examples.
	Examples []TrainingExample `json:"examples" validate:"required"`
}

// NewTrainingQuery : Instantiate TrainingQuery (Generic Model Constructor)
func (discovery *DiscoveryV2) NewTrainingQuery(naturalLanguageQuery string, examples []TrainingExample) (model *TrainingQuery, err error) {
	model = &TrainingQuery{
		NaturalLanguageQuery: core.StringPtr(naturalLanguageQuery),
		Examples:             examples,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// TrainingQuerySet : Object specifying the training queries contained in the identified training set.
type TrainingQuerySet struct {

	// Array of training queries.
	Queries []TrainingQuery `json:"queries,omitempty"`
}

// UpdateDocumentOptions : The UpdateDocument options.
type UpdateDocumentOptions struct {

	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required"`

	// The ID of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// The ID of the document.
	DocumentID *string `json:"document_id" validate:"required"`

	// The content of the document to ingest. The maximum supported file size when adding a file to a collection is 50
	// megabytes, the maximum supported file size when testing a configuration is 1 megabyte. Files larger than the
	// supported size are rejected.
	File io.ReadCloser `json:"file,omitempty"`

	// The filename for file.
	Filename *string `json:"filename,omitempty"`

	// The content type of file.
	FileContentType *string `json:"file_content_type,omitempty"`

	// The maximum supported metadata file size is 1 MB. Metadata parts larger than 1 MB are rejected. Example:  ``` {
	//   "Creator": "Johnny Appleseed",
	//   "Subject": "Apples"
	// } ```.
	Metadata *string `json:"metadata,omitempty"`

	// When `true`, the uploaded document is added to the collection even if the data for that collection is shared with
	// other collections.
	XWatsonDiscoveryForce *bool `json:"X-Watson-Discovery-Force,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewUpdateDocumentOptions : Instantiate UpdateDocumentOptions
func (discovery *DiscoveryV2) NewUpdateDocumentOptions(projectID string, collectionID string, documentID string) *UpdateDocumentOptions {
	return &UpdateDocumentOptions{
		ProjectID:    core.StringPtr(projectID),
		CollectionID: core.StringPtr(collectionID),
		DocumentID:   core.StringPtr(documentID),
	}
}

// SetProjectID : Allow user to set ProjectID
func (options *UpdateDocumentOptions) SetProjectID(projectID string) *UpdateDocumentOptions {
	options.ProjectID = core.StringPtr(projectID)
	return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *UpdateDocumentOptions) SetCollectionID(collectionID string) *UpdateDocumentOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
}

// SetDocumentID : Allow user to set DocumentID
func (options *UpdateDocumentOptions) SetDocumentID(documentID string) *UpdateDocumentOptions {
	options.DocumentID = core.StringPtr(documentID)
	return options
}

// SetFile : Allow user to set File
func (options *UpdateDocumentOptions) SetFile(file io.ReadCloser) *UpdateDocumentOptions {
	options.File = file
	return options
}

// SetFilename : Allow user to set Filename
func (options *UpdateDocumentOptions) SetFilename(filename string) *UpdateDocumentOptions {
	options.Filename = core.StringPtr(filename)
	return options
}

// SetFileContentType : Allow user to set FileContentType
func (options *UpdateDocumentOptions) SetFileContentType(fileContentType string) *UpdateDocumentOptions {
	options.FileContentType = core.StringPtr(fileContentType)
	return options
}

// SetMetadata : Allow user to set Metadata
func (options *UpdateDocumentOptions) SetMetadata(metadata string) *UpdateDocumentOptions {
	options.Metadata = core.StringPtr(metadata)
	return options
}

// SetXWatsonDiscoveryForce : Allow user to set XWatsonDiscoveryForce
func (options *UpdateDocumentOptions) SetXWatsonDiscoveryForce(xWatsonDiscoveryForce bool) *UpdateDocumentOptions {
	options.XWatsonDiscoveryForce = core.BoolPtr(xWatsonDiscoveryForce)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateDocumentOptions) SetHeaders(param map[string]string) *UpdateDocumentOptions {
	options.Headers = param
	return options
}

// UpdateTrainingQueryOptions : The UpdateTrainingQuery options.
type UpdateTrainingQueryOptions struct {

	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required"`

	// The ID of the query used for training.
	QueryID *string `json:"query_id" validate:"required"`

	// The natural text query for the training query.
	NaturalLanguageQuery *string `json:"natural_language_query" validate:"required"`

	// Array of training examples.
	Examples []TrainingExample `json:"examples" validate:"required"`

	// The filter used on the collection before the **natural_language_query** is applied.
	Filter *string `json:"filter,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewUpdateTrainingQueryOptions : Instantiate UpdateTrainingQueryOptions
func (discovery *DiscoveryV2) NewUpdateTrainingQueryOptions(projectID string, queryID string, naturalLanguageQuery string, examples []TrainingExample) *UpdateTrainingQueryOptions {
	return &UpdateTrainingQueryOptions{
		ProjectID:            core.StringPtr(projectID),
		QueryID:              core.StringPtr(queryID),
		NaturalLanguageQuery: core.StringPtr(naturalLanguageQuery),
		Examples:             examples,
	}
}

// SetProjectID : Allow user to set ProjectID
func (options *UpdateTrainingQueryOptions) SetProjectID(projectID string) *UpdateTrainingQueryOptions {
	options.ProjectID = core.StringPtr(projectID)
	return options
}

// SetQueryID : Allow user to set QueryID
func (options *UpdateTrainingQueryOptions) SetQueryID(queryID string) *UpdateTrainingQueryOptions {
	options.QueryID = core.StringPtr(queryID)
	return options
}

// SetNaturalLanguageQuery : Allow user to set NaturalLanguageQuery
func (options *UpdateTrainingQueryOptions) SetNaturalLanguageQuery(naturalLanguageQuery string) *UpdateTrainingQueryOptions {
	options.NaturalLanguageQuery = core.StringPtr(naturalLanguageQuery)
	return options
}

// SetExamples : Allow user to set Examples
func (options *UpdateTrainingQueryOptions) SetExamples(examples []TrainingExample) *UpdateTrainingQueryOptions {
	options.Examples = examples
	return options
}

// SetFilter : Allow user to set Filter
func (options *UpdateTrainingQueryOptions) SetFilter(filter string) *UpdateTrainingQueryOptions {
	options.Filter = core.StringPtr(filter)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateTrainingQueryOptions) SetHeaders(param map[string]string) *UpdateTrainingQueryOptions {
	options.Headers = param
	return options
}

// QueryCalculationAggregation : Returns a scalar calculation across all documents for the field specified. Possible calculations include min, max,
// sum, average, and unique_count.
type QueryCalculationAggregation struct {

	// The field to perform the calculation on.
	Field *string `json:"field" validate:"required"`

	// The value of the calculation.
	Value *float64 `json:"value,omitempty"`
}

// QueryFilterAggregation : A modifier that will narrow down the document set of the sub aggregations it precedes.
type QueryFilterAggregation struct {

	// The filter written in Discovery Query Language syntax applied to the documents before sub aggregations are run.
	Match *string `json:"match" validate:"required"`

	// Number of documents matching the filter.
	MatchingResults *int64 `json:"matching_results" validate:"required"`

	// An array of sub aggregations.
	Aggregations []QueryAggregation `json:"aggregations,omitempty"`
}

// QueryHistogramAggregation : Numeric interval segments to categorize documents by using field values from a single numeric field to describe the
// category.
type QueryHistogramAggregation struct {

	// The numeric field name used to create the histogram.
	Field *string `json:"field" validate:"required"`

	// The size of the sections the results are split into.
	Interval *int64 `json:"interval" validate:"required"`

	// Array of numeric intervals.
	Results []QueryHistogramAggregationResult `json:"results,omitempty"`
}

// QueryNestedAggregation : A restriction that alter the document set used for sub aggregations it precedes to nested documents found in the
// field specified.
type QueryNestedAggregation struct {

	// The path to the document field to scope sub aggregations to.
	Path *string `json:"path" validate:"required"`

	// Number of nested documents found in the specified field.
	MatchingResults *int64 `json:"matching_results" validate:"required"`

	// An array of sub aggregations.
	Aggregations []QueryAggregation `json:"aggregations,omitempty"`
}

// QueryTermAggregation : Returns the top values for the field specified.
type QueryTermAggregation struct {

	// The field in the document used to generate top values from.
	Field *string `json:"field" validate:"required"`

	// The number of top values returned.
	Count *int64 `json:"count,omitempty"`

	// Array of top values for the field.
	Results []QueryTermAggregationResult `json:"results,omitempty"`
}

// QueryTimesliceAggregation : A specialized histogram aggregation that uses dates to create interval segments.
type QueryTimesliceAggregation struct {

	// The date field name used to create the timeslice.
	Field *string `json:"field" validate:"required"`

	// The date interval value. Valid values are seconds, minutes, hours, days, weeks, and years.
	Interval *string `json:"interval" validate:"required"`

	// Array of aggregation results.
	Results []QueryTimesliceAggregationResult `json:"results,omitempty"`
}

// QueryTopHitsAggregation : Returns the top documents ranked by the score of the query.
type QueryTopHitsAggregation struct {

	// The number of documents to return.
	Size *int64 `json:"size" validate:"required"`

	Hits *QueryTopHitsAggregationResult `json:"hits,omitempty"`
}
