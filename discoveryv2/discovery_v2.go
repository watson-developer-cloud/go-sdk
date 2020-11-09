/**
 * (C) Copyright IBM Corp. 2020.
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
 

// Package discoveryv2 : Operations and models for the DiscoveryV2 service
package discoveryv2

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
	"strings"
	"time"
)

// DiscoveryV2 : IBM Watson&trade; Discovery is a cognitive search and content analytics engine that you can add to
// applications to identify patterns, trends and actionable insights to drive better decision-making. Securely unify
// structured and unstructured data with pre-enriched content, and use a simplified query language to eliminate the need
// for manual filtering of results.
//
// Version: 2.0
// See: https://cloud.ibm.com/docs/discovery-data
type DiscoveryV2 struct {
	Service *core.BaseService

	// Release date of the version of the API you want to use. Specify dates in YYYY-MM-DD format. The current version is
	// `2019-11-22`.
	Version *string
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.us-south.discovery.watson.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "discovery"

// DiscoveryV2Options : Service options
type DiscoveryV2Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator

	// Release date of the version of the API you want to use. Specify dates in YYYY-MM-DD format. The current version is
	// `2019-11-22`.
	Version *string `validate:"required"`
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

	service = &DiscoveryV2{
		Service: baseService,
		Version: options.Version,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "discovery" suitable for processing requests.
func (discovery *DiscoveryV2) Clone() *DiscoveryV2 {
	if core.IsNil(discovery) {
		return nil
	}
	clone := *discovery
	clone.Service = discovery.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (discovery *DiscoveryV2) SetServiceURL(url string) error {
	return discovery.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (discovery *DiscoveryV2) GetServiceURL() string {
	return discovery.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (discovery *DiscoveryV2) SetDefaultHeaders(headers http.Header) {
	discovery.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (discovery *DiscoveryV2) SetEnableGzipCompression(enableGzip bool) {
	discovery.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (discovery *DiscoveryV2) GetEnableGzipCompression() bool {
	return discovery.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (discovery *DiscoveryV2) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	discovery.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (discovery *DiscoveryV2) DisableRetries() {
	discovery.Service.DisableRetries()
}

// DisableSSLVerification bypasses verification of the server's SSL certificate
func (discovery *DiscoveryV2) DisableSSLVerification() {
	discovery.Service.DisableSSLVerification()
}

// ListCollections : List collections
// Lists existing collections for the specified project.
func (discovery *DiscoveryV2) ListCollections(listCollectionsOptions *ListCollectionsOptions) (result *ListCollectionsResponse, response *core.DetailedResponse, err error) {
	return discovery.ListCollectionsWithContext(context.Background(), listCollectionsOptions)
}

// ListCollectionsWithContext is an alternate form of the ListCollections method which supports a Context parameter
func (discovery *DiscoveryV2) ListCollectionsWithContext(ctx context.Context, listCollectionsOptions *ListCollectionsOptions) (result *ListCollectionsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listCollectionsOptions, "listCollectionsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listCollectionsOptions, "listCollectionsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *listCollectionsOptions.ProjectID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v2/projects/{project_id}/collections`, pathParamsMap)
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

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = discovery.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListCollectionsResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CreateCollection : Create a collection
// Create a new collection in the specified project.
func (discovery *DiscoveryV2) CreateCollection(createCollectionOptions *CreateCollectionOptions) (result *CollectionDetails, response *core.DetailedResponse, err error) {
	return discovery.CreateCollectionWithContext(context.Background(), createCollectionOptions)
}

// CreateCollectionWithContext is an alternate form of the CreateCollection method which supports a Context parameter
func (discovery *DiscoveryV2) CreateCollectionWithContext(ctx context.Context, createCollectionOptions *CreateCollectionOptions) (result *CollectionDetails, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createCollectionOptions, "createCollectionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createCollectionOptions, "createCollectionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *createCollectionOptions.ProjectID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v2/projects/{project_id}/collections`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createCollectionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V2", "CreateCollection")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

	body := make(map[string]interface{})
	if createCollectionOptions.Name != nil {
		body["name"] = createCollectionOptions.Name
	}
	if createCollectionOptions.Description != nil {
		body["description"] = createCollectionOptions.Description
	}
	if createCollectionOptions.Language != nil {
		body["language"] = createCollectionOptions.Language
	}
	if createCollectionOptions.Enrichments != nil {
		body["enrichments"] = createCollectionOptions.Enrichments
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
	response, err = discovery.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCollectionDetails)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetCollection : Get collection
// Get details about the specified collection.
func (discovery *DiscoveryV2) GetCollection(getCollectionOptions *GetCollectionOptions) (result *CollectionDetails, response *core.DetailedResponse, err error) {
	return discovery.GetCollectionWithContext(context.Background(), getCollectionOptions)
}

// GetCollectionWithContext is an alternate form of the GetCollection method which supports a Context parameter
func (discovery *DiscoveryV2) GetCollectionWithContext(ctx context.Context, getCollectionOptions *GetCollectionOptions) (result *CollectionDetails, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getCollectionOptions, "getCollectionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getCollectionOptions, "getCollectionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *getCollectionOptions.ProjectID,
		"collection_id": *getCollectionOptions.CollectionID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v2/projects/{project_id}/collections/{collection_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCollectionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V2", "GetCollection")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = discovery.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCollectionDetails)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UpdateCollection : Update a collection
// Updates the specified collection's name, description, and enrichments.
func (discovery *DiscoveryV2) UpdateCollection(updateCollectionOptions *UpdateCollectionOptions) (result *CollectionDetails, response *core.DetailedResponse, err error) {
	return discovery.UpdateCollectionWithContext(context.Background(), updateCollectionOptions)
}

// UpdateCollectionWithContext is an alternate form of the UpdateCollection method which supports a Context parameter
func (discovery *DiscoveryV2) UpdateCollectionWithContext(ctx context.Context, updateCollectionOptions *UpdateCollectionOptions) (result *CollectionDetails, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateCollectionOptions, "updateCollectionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateCollectionOptions, "updateCollectionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *updateCollectionOptions.ProjectID,
		"collection_id": *updateCollectionOptions.CollectionID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v2/projects/{project_id}/collections/{collection_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateCollectionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V2", "UpdateCollection")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

	body := make(map[string]interface{})
	if updateCollectionOptions.Name != nil {
		body["name"] = updateCollectionOptions.Name
	}
	if updateCollectionOptions.Description != nil {
		body["description"] = updateCollectionOptions.Description
	}
	if updateCollectionOptions.Enrichments != nil {
		body["enrichments"] = updateCollectionOptions.Enrichments
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
	response, err = discovery.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCollectionDetails)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteCollection : Delete a collection
// Deletes the specified collection from the project. All documents stored in the specified collection and not shared is
// also deleted.
func (discovery *DiscoveryV2) DeleteCollection(deleteCollectionOptions *DeleteCollectionOptions) (response *core.DetailedResponse, err error) {
	return discovery.DeleteCollectionWithContext(context.Background(), deleteCollectionOptions)
}

// DeleteCollectionWithContext is an alternate form of the DeleteCollection method which supports a Context parameter
func (discovery *DiscoveryV2) DeleteCollectionWithContext(ctx context.Context, deleteCollectionOptions *DeleteCollectionOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteCollectionOptions, "deleteCollectionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteCollectionOptions, "deleteCollectionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *deleteCollectionOptions.ProjectID,
		"collection_id": *deleteCollectionOptions.CollectionID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v2/projects/{project_id}/collections/{collection_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteCollectionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V2", "DeleteCollection")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = discovery.Service.Request(request, nil)

	return
}

// Query : Query a project
// By using this method, you can construct queries. For details, see the [Discovery
// documentation](https://cloud.ibm.com/docs/discovery-data?topic=discovery-data-query-concepts). The default query
// parameters are defined by the settings for this project, see the [Discovery
// documentation](https://cloud.ibm.com/docs/discovery-data?topic=discovery-data-project-defaults) for an overview of
// the standard default settings, and see [the Projects API documentation](#create-project) for details about how to set
// custom default query settings.
func (discovery *DiscoveryV2) Query(queryOptions *QueryOptions) (result *QueryResponse, response *core.DetailedResponse, err error) {
	return discovery.QueryWithContext(context.Background(), queryOptions)
}

// QueryWithContext is an alternate form of the Query method which supports a Context parameter
func (discovery *DiscoveryV2) QueryWithContext(ctx context.Context, queryOptions *QueryOptions) (result *QueryResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(queryOptions, "queryOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(queryOptions, "queryOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *queryOptions.ProjectID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v2/projects/{project_id}/query`, pathParamsMap)
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

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

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

	var rawResponse map[string]json.RawMessage
	response, err = discovery.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalQueryResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetAutocompletion : Get Autocomplete Suggestions
// Returns completion query suggestions for the specified prefix.
func (discovery *DiscoveryV2) GetAutocompletion(getAutocompletionOptions *GetAutocompletionOptions) (result *Completions, response *core.DetailedResponse, err error) {
	return discovery.GetAutocompletionWithContext(context.Background(), getAutocompletionOptions)
}

// GetAutocompletionWithContext is an alternate form of the GetAutocompletion method which supports a Context parameter
func (discovery *DiscoveryV2) GetAutocompletionWithContext(ctx context.Context, getAutocompletionOptions *GetAutocompletionOptions) (result *Completions, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAutocompletionOptions, "getAutocompletionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAutocompletionOptions, "getAutocompletionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *getAutocompletionOptions.ProjectID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v2/projects/{project_id}/autocompletion`, pathParamsMap)
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

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))
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

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = discovery.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCompletions)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// QueryNotices : Query system notices
// Queries for notices (errors or warnings) that might have been generated by the system. Notices are generated when
// ingesting documents and performing relevance training.
func (discovery *DiscoveryV2) QueryNotices(queryNoticesOptions *QueryNoticesOptions) (result *QueryNoticesResponse, response *core.DetailedResponse, err error) {
	return discovery.QueryNoticesWithContext(context.Background(), queryNoticesOptions)
}

// QueryNoticesWithContext is an alternate form of the QueryNotices method which supports a Context parameter
func (discovery *DiscoveryV2) QueryNoticesWithContext(ctx context.Context, queryNoticesOptions *QueryNoticesOptions) (result *QueryNoticesResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(queryNoticesOptions, "queryNoticesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(queryNoticesOptions, "queryNoticesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *queryNoticesOptions.ProjectID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v2/projects/{project_id}/notices`, pathParamsMap)
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

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))
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

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = discovery.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalQueryNoticesResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ListFields : List fields
// Gets a list of the unique fields (and their types) stored in the the specified collections.
func (discovery *DiscoveryV2) ListFields(listFieldsOptions *ListFieldsOptions) (result *ListFieldsResponse, response *core.DetailedResponse, err error) {
	return discovery.ListFieldsWithContext(context.Background(), listFieldsOptions)
}

// ListFieldsWithContext is an alternate form of the ListFields method which supports a Context parameter
func (discovery *DiscoveryV2) ListFieldsWithContext(ctx context.Context, listFieldsOptions *ListFieldsOptions) (result *ListFieldsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listFieldsOptions, "listFieldsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listFieldsOptions, "listFieldsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *listFieldsOptions.ProjectID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v2/projects/{project_id}/fields`, pathParamsMap)
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

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))
	if listFieldsOptions.CollectionIds != nil {
		builder.AddQuery("collection_ids", strings.Join(listFieldsOptions.CollectionIds, ","))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = discovery.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListFieldsResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetComponentSettings : List component settings
// Returns default configuration settings for components.
func (discovery *DiscoveryV2) GetComponentSettings(getComponentSettingsOptions *GetComponentSettingsOptions) (result *ComponentSettingsResponse, response *core.DetailedResponse, err error) {
	return discovery.GetComponentSettingsWithContext(context.Background(), getComponentSettingsOptions)
}

// GetComponentSettingsWithContext is an alternate form of the GetComponentSettings method which supports a Context parameter
func (discovery *DiscoveryV2) GetComponentSettingsWithContext(ctx context.Context, getComponentSettingsOptions *GetComponentSettingsOptions) (result *ComponentSettingsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getComponentSettingsOptions, "getComponentSettingsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getComponentSettingsOptions, "getComponentSettingsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *getComponentSettingsOptions.ProjectID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v2/projects/{project_id}/component_settings`, pathParamsMap)
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

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = discovery.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalComponentSettingsResponse)
	if err != nil {
		return
	}
	response.Result = result

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
	return discovery.AddDocumentWithContext(context.Background(), addDocumentOptions)
}

// AddDocumentWithContext is an alternate form of the AddDocument method which supports a Context parameter
func (discovery *DiscoveryV2) AddDocumentWithContext(ctx context.Context, addDocumentOptions *AddDocumentOptions) (result *DocumentAccepted, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addDocumentOptions, "addDocumentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addDocumentOptions, "addDocumentOptions")
	if err != nil {
		return
	}
	if (addDocumentOptions.File == nil) && (addDocumentOptions.Metadata == nil) {
		err = fmt.Errorf("at least one of file or metadata must be supplied")
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *addDocumentOptions.ProjectID,
		"collection_id": *addDocumentOptions.CollectionID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v2/projects/{project_id}/collections/{collection_id}/documents`, pathParamsMap)
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

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

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

	var rawResponse map[string]json.RawMessage
	response, err = discovery.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDocumentAccepted)
	if err != nil {
		return
	}
	response.Result = result

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
//
// **Note:** If an uploaded document is segmented, all segments will be overwritten, even if the updated version of the
// document has fewer segments.
func (discovery *DiscoveryV2) UpdateDocument(updateDocumentOptions *UpdateDocumentOptions) (result *DocumentAccepted, response *core.DetailedResponse, err error) {
	return discovery.UpdateDocumentWithContext(context.Background(), updateDocumentOptions)
}

// UpdateDocumentWithContext is an alternate form of the UpdateDocument method which supports a Context parameter
func (discovery *DiscoveryV2) UpdateDocumentWithContext(ctx context.Context, updateDocumentOptions *UpdateDocumentOptions) (result *DocumentAccepted, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateDocumentOptions, "updateDocumentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateDocumentOptions, "updateDocumentOptions")
	if err != nil {
		return
	}
	if (updateDocumentOptions.File == nil) && (updateDocumentOptions.Metadata == nil) {
		err = fmt.Errorf("at least one of file or metadata must be supplied")
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *updateDocumentOptions.ProjectID,
		"collection_id": *updateDocumentOptions.CollectionID,
		"document_id": *updateDocumentOptions.DocumentID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v2/projects/{project_id}/collections/{collection_id}/documents/{document_id}`, pathParamsMap)
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

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

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

	var rawResponse map[string]json.RawMessage
	response, err = discovery.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDocumentAccepted)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteDocument : Delete a document
// If the given document ID is invalid, or if the document is not found, then the a success response is returned (HTTP
// status code `200`) with the status set to 'deleted'.
//
// **Note:** This operation only works on collections created to accept direct file uploads. It cannot be used to modify
// a collection that connects to an external source such as Microsoft SharePoint.
//
// **Note:** Segments of an uploaded document cannot be deleted individually. Delete all segments by deleting using the
// `parent_document_id` of a segment result.
func (discovery *DiscoveryV2) DeleteDocument(deleteDocumentOptions *DeleteDocumentOptions) (result *DeleteDocumentResponse, response *core.DetailedResponse, err error) {
	return discovery.DeleteDocumentWithContext(context.Background(), deleteDocumentOptions)
}

// DeleteDocumentWithContext is an alternate form of the DeleteDocument method which supports a Context parameter
func (discovery *DiscoveryV2) DeleteDocumentWithContext(ctx context.Context, deleteDocumentOptions *DeleteDocumentOptions) (result *DeleteDocumentResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteDocumentOptions, "deleteDocumentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteDocumentOptions, "deleteDocumentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *deleteDocumentOptions.ProjectID,
		"collection_id": *deleteDocumentOptions.CollectionID,
		"document_id": *deleteDocumentOptions.DocumentID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v2/projects/{project_id}/collections/{collection_id}/documents/{document_id}`, pathParamsMap)
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

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = discovery.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeleteDocumentResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ListTrainingQueries : List training queries
// List the training queries for the specified project.
func (discovery *DiscoveryV2) ListTrainingQueries(listTrainingQueriesOptions *ListTrainingQueriesOptions) (result *TrainingQuerySet, response *core.DetailedResponse, err error) {
	return discovery.ListTrainingQueriesWithContext(context.Background(), listTrainingQueriesOptions)
}

// ListTrainingQueriesWithContext is an alternate form of the ListTrainingQueries method which supports a Context parameter
func (discovery *DiscoveryV2) ListTrainingQueriesWithContext(ctx context.Context, listTrainingQueriesOptions *ListTrainingQueriesOptions) (result *TrainingQuerySet, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listTrainingQueriesOptions, "listTrainingQueriesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listTrainingQueriesOptions, "listTrainingQueriesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *listTrainingQueriesOptions.ProjectID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v2/projects/{project_id}/training_data/queries`, pathParamsMap)
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

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = discovery.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTrainingQuerySet)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteTrainingQueries : Delete training queries
// Removes all training queries for the specified project.
func (discovery *DiscoveryV2) DeleteTrainingQueries(deleteTrainingQueriesOptions *DeleteTrainingQueriesOptions) (response *core.DetailedResponse, err error) {
	return discovery.DeleteTrainingQueriesWithContext(context.Background(), deleteTrainingQueriesOptions)
}

// DeleteTrainingQueriesWithContext is an alternate form of the DeleteTrainingQueries method which supports a Context parameter
func (discovery *DiscoveryV2) DeleteTrainingQueriesWithContext(ctx context.Context, deleteTrainingQueriesOptions *DeleteTrainingQueriesOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteTrainingQueriesOptions, "deleteTrainingQueriesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteTrainingQueriesOptions, "deleteTrainingQueriesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *deleteTrainingQueriesOptions.ProjectID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v2/projects/{project_id}/training_data/queries`, pathParamsMap)
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

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

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
	return discovery.CreateTrainingQueryWithContext(context.Background(), createTrainingQueryOptions)
}

// CreateTrainingQueryWithContext is an alternate form of the CreateTrainingQuery method which supports a Context parameter
func (discovery *DiscoveryV2) CreateTrainingQueryWithContext(ctx context.Context, createTrainingQueryOptions *CreateTrainingQueryOptions) (result *TrainingQuery, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createTrainingQueryOptions, "createTrainingQueryOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createTrainingQueryOptions, "createTrainingQueryOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *createTrainingQueryOptions.ProjectID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v2/projects/{project_id}/training_data/queries`, pathParamsMap)
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

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

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

	var rawResponse map[string]json.RawMessage
	response, err = discovery.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTrainingQuery)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetTrainingQuery : Get a training data query
// Get details for a specific training data query, including the query string and all examples.
func (discovery *DiscoveryV2) GetTrainingQuery(getTrainingQueryOptions *GetTrainingQueryOptions) (result *TrainingQuery, response *core.DetailedResponse, err error) {
	return discovery.GetTrainingQueryWithContext(context.Background(), getTrainingQueryOptions)
}

// GetTrainingQueryWithContext is an alternate form of the GetTrainingQuery method which supports a Context parameter
func (discovery *DiscoveryV2) GetTrainingQueryWithContext(ctx context.Context, getTrainingQueryOptions *GetTrainingQueryOptions) (result *TrainingQuery, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getTrainingQueryOptions, "getTrainingQueryOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getTrainingQueryOptions, "getTrainingQueryOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *getTrainingQueryOptions.ProjectID,
		"query_id": *getTrainingQueryOptions.QueryID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v2/projects/{project_id}/training_data/queries/{query_id}`, pathParamsMap)
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

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = discovery.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTrainingQuery)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UpdateTrainingQuery : Update a training query
// Updates an existing training query and it's examples.
func (discovery *DiscoveryV2) UpdateTrainingQuery(updateTrainingQueryOptions *UpdateTrainingQueryOptions) (result *TrainingQuery, response *core.DetailedResponse, err error) {
	return discovery.UpdateTrainingQueryWithContext(context.Background(), updateTrainingQueryOptions)
}

// UpdateTrainingQueryWithContext is an alternate form of the UpdateTrainingQuery method which supports a Context parameter
func (discovery *DiscoveryV2) UpdateTrainingQueryWithContext(ctx context.Context, updateTrainingQueryOptions *UpdateTrainingQueryOptions) (result *TrainingQuery, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateTrainingQueryOptions, "updateTrainingQueryOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateTrainingQueryOptions, "updateTrainingQueryOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *updateTrainingQueryOptions.ProjectID,
		"query_id": *updateTrainingQueryOptions.QueryID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v2/projects/{project_id}/training_data/queries/{query_id}`, pathParamsMap)
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

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

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

	var rawResponse map[string]json.RawMessage
	response, err = discovery.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTrainingQuery)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// AnalyzeDocument : Analyze a Document
// Process a document using the specified collection's settings and return it for realtime use.
//
// **Note:** Documents processed using this method are not added to the specified collection.
//
// **Note:** This method is only supported on IBM Cloud Pak for Data instances of Discovery.
func (discovery *DiscoveryV2) AnalyzeDocument(analyzeDocumentOptions *AnalyzeDocumentOptions) (result *AnalyzedDocument, response *core.DetailedResponse, err error) {
	return discovery.AnalyzeDocumentWithContext(context.Background(), analyzeDocumentOptions)
}

// AnalyzeDocumentWithContext is an alternate form of the AnalyzeDocument method which supports a Context parameter
func (discovery *DiscoveryV2) AnalyzeDocumentWithContext(ctx context.Context, analyzeDocumentOptions *AnalyzeDocumentOptions) (result *AnalyzedDocument, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(analyzeDocumentOptions, "analyzeDocumentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(analyzeDocumentOptions, "analyzeDocumentOptions")
	if err != nil {
		return
	}
	if (analyzeDocumentOptions.File == nil) && (analyzeDocumentOptions.Metadata == nil) {
		err = fmt.Errorf("at least one of file or metadata must be supplied")
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *analyzeDocumentOptions.ProjectID,
		"collection_id": *analyzeDocumentOptions.CollectionID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v2/projects/{project_id}/collections/{collection_id}/analyze`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range analyzeDocumentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V2", "AnalyzeDocument")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

	if analyzeDocumentOptions.File != nil {
		builder.AddFormData("file", core.StringNilMapper(analyzeDocumentOptions.Filename),
			core.StringNilMapper(analyzeDocumentOptions.FileContentType), analyzeDocumentOptions.File)
	}
	if analyzeDocumentOptions.Metadata != nil {
		builder.AddFormData("metadata", "", "", fmt.Sprint(*analyzeDocumentOptions.Metadata))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = discovery.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAnalyzedDocument)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ListEnrichments : List Enrichments
// List the enrichments available to this project.
func (discovery *DiscoveryV2) ListEnrichments(listEnrichmentsOptions *ListEnrichmentsOptions) (result *Enrichments, response *core.DetailedResponse, err error) {
	return discovery.ListEnrichmentsWithContext(context.Background(), listEnrichmentsOptions)
}

// ListEnrichmentsWithContext is an alternate form of the ListEnrichments method which supports a Context parameter
func (discovery *DiscoveryV2) ListEnrichmentsWithContext(ctx context.Context, listEnrichmentsOptions *ListEnrichmentsOptions) (result *Enrichments, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listEnrichmentsOptions, "listEnrichmentsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listEnrichmentsOptions, "listEnrichmentsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *listEnrichmentsOptions.ProjectID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v2/projects/{project_id}/enrichments`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listEnrichmentsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V2", "ListEnrichments")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = discovery.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalEnrichments)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CreateEnrichment : Create an enrichment
// Create an enrichment for use with the specified project/.
func (discovery *DiscoveryV2) CreateEnrichment(createEnrichmentOptions *CreateEnrichmentOptions) (result *Enrichment, response *core.DetailedResponse, err error) {
	return discovery.CreateEnrichmentWithContext(context.Background(), createEnrichmentOptions)
}

// CreateEnrichmentWithContext is an alternate form of the CreateEnrichment method which supports a Context parameter
func (discovery *DiscoveryV2) CreateEnrichmentWithContext(ctx context.Context, createEnrichmentOptions *CreateEnrichmentOptions) (result *Enrichment, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createEnrichmentOptions, "createEnrichmentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createEnrichmentOptions, "createEnrichmentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *createEnrichmentOptions.ProjectID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v2/projects/{project_id}/enrichments`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createEnrichmentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V2", "CreateEnrichment")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

	builder.AddFormData("enrichment", "", "application/json", createEnrichmentOptions.Enrichment)
	if createEnrichmentOptions.File != nil {
		builder.AddFormData("file", "filename",
			"application/octet-stream", createEnrichmentOptions.File)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = discovery.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalEnrichment)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetEnrichment : Get enrichment
// Get details about a specific enrichment.
func (discovery *DiscoveryV2) GetEnrichment(getEnrichmentOptions *GetEnrichmentOptions) (result *Enrichment, response *core.DetailedResponse, err error) {
	return discovery.GetEnrichmentWithContext(context.Background(), getEnrichmentOptions)
}

// GetEnrichmentWithContext is an alternate form of the GetEnrichment method which supports a Context parameter
func (discovery *DiscoveryV2) GetEnrichmentWithContext(ctx context.Context, getEnrichmentOptions *GetEnrichmentOptions) (result *Enrichment, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getEnrichmentOptions, "getEnrichmentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getEnrichmentOptions, "getEnrichmentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *getEnrichmentOptions.ProjectID,
		"enrichment_id": *getEnrichmentOptions.EnrichmentID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v2/projects/{project_id}/enrichments/{enrichment_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getEnrichmentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V2", "GetEnrichment")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = discovery.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalEnrichment)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UpdateEnrichment : Update an enrichment
// Updates an existing enrichment's name and description.
func (discovery *DiscoveryV2) UpdateEnrichment(updateEnrichmentOptions *UpdateEnrichmentOptions) (result *Enrichment, response *core.DetailedResponse, err error) {
	return discovery.UpdateEnrichmentWithContext(context.Background(), updateEnrichmentOptions)
}

// UpdateEnrichmentWithContext is an alternate form of the UpdateEnrichment method which supports a Context parameter
func (discovery *DiscoveryV2) UpdateEnrichmentWithContext(ctx context.Context, updateEnrichmentOptions *UpdateEnrichmentOptions) (result *Enrichment, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateEnrichmentOptions, "updateEnrichmentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateEnrichmentOptions, "updateEnrichmentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *updateEnrichmentOptions.ProjectID,
		"enrichment_id": *updateEnrichmentOptions.EnrichmentID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v2/projects/{project_id}/enrichments/{enrichment_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateEnrichmentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V2", "UpdateEnrichment")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

	body := make(map[string]interface{})
	if updateEnrichmentOptions.Name != nil {
		body["name"] = updateEnrichmentOptions.Name
	}
	if updateEnrichmentOptions.Description != nil {
		body["description"] = updateEnrichmentOptions.Description
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
	response, err = discovery.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalEnrichment)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteEnrichment : Delete an enrichment
// Deletes an existing enrichment from the specified project.
//
// **Note:** Only enrichments that have been manually created can be deleted.
func (discovery *DiscoveryV2) DeleteEnrichment(deleteEnrichmentOptions *DeleteEnrichmentOptions) (response *core.DetailedResponse, err error) {
	return discovery.DeleteEnrichmentWithContext(context.Background(), deleteEnrichmentOptions)
}

// DeleteEnrichmentWithContext is an alternate form of the DeleteEnrichment method which supports a Context parameter
func (discovery *DiscoveryV2) DeleteEnrichmentWithContext(ctx context.Context, deleteEnrichmentOptions *DeleteEnrichmentOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteEnrichmentOptions, "deleteEnrichmentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteEnrichmentOptions, "deleteEnrichmentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *deleteEnrichmentOptions.ProjectID,
		"enrichment_id": *deleteEnrichmentOptions.EnrichmentID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v2/projects/{project_id}/enrichments/{enrichment_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteEnrichmentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V2", "DeleteEnrichment")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = discovery.Service.Request(request, nil)

	return
}

// ListProjects : List projects
// Lists existing projects for this instance.
func (discovery *DiscoveryV2) ListProjects(listProjectsOptions *ListProjectsOptions) (result *ListProjectsResponse, response *core.DetailedResponse, err error) {
	return discovery.ListProjectsWithContext(context.Background(), listProjectsOptions)
}

// ListProjectsWithContext is an alternate form of the ListProjects method which supports a Context parameter
func (discovery *DiscoveryV2) ListProjectsWithContext(ctx context.Context, listProjectsOptions *ListProjectsOptions) (result *ListProjectsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listProjectsOptions, "listProjectsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v2/projects`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listProjectsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V2", "ListProjects")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = discovery.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListProjectsResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CreateProject : Create a Project
// Create a new project for this instance.
func (discovery *DiscoveryV2) CreateProject(createProjectOptions *CreateProjectOptions) (result *ProjectDetails, response *core.DetailedResponse, err error) {
	return discovery.CreateProjectWithContext(context.Background(), createProjectOptions)
}

// CreateProjectWithContext is an alternate form of the CreateProject method which supports a Context parameter
func (discovery *DiscoveryV2) CreateProjectWithContext(ctx context.Context, createProjectOptions *CreateProjectOptions) (result *ProjectDetails, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createProjectOptions, "createProjectOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createProjectOptions, "createProjectOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v2/projects`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createProjectOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V2", "CreateProject")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

	body := make(map[string]interface{})
	if createProjectOptions.Name != nil {
		body["name"] = createProjectOptions.Name
	}
	if createProjectOptions.Type != nil {
		body["type"] = createProjectOptions.Type
	}
	if createProjectOptions.DefaultQueryParameters != nil {
		body["default_query_parameters"] = createProjectOptions.DefaultQueryParameters
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
	response, err = discovery.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectDetails)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetProject : Get project
// Get details on the specified project.
func (discovery *DiscoveryV2) GetProject(getProjectOptions *GetProjectOptions) (result *ProjectDetails, response *core.DetailedResponse, err error) {
	return discovery.GetProjectWithContext(context.Background(), getProjectOptions)
}

// GetProjectWithContext is an alternate form of the GetProject method which supports a Context parameter
func (discovery *DiscoveryV2) GetProjectWithContext(ctx context.Context, getProjectOptions *GetProjectOptions) (result *ProjectDetails, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getProjectOptions, "getProjectOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getProjectOptions, "getProjectOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *getProjectOptions.ProjectID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v2/projects/{project_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getProjectOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V2", "GetProject")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = discovery.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectDetails)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UpdateProject : Update a project
// Update the specified project's name.
func (discovery *DiscoveryV2) UpdateProject(updateProjectOptions *UpdateProjectOptions) (result *ProjectDetails, response *core.DetailedResponse, err error) {
	return discovery.UpdateProjectWithContext(context.Background(), updateProjectOptions)
}

// UpdateProjectWithContext is an alternate form of the UpdateProject method which supports a Context parameter
func (discovery *DiscoveryV2) UpdateProjectWithContext(ctx context.Context, updateProjectOptions *UpdateProjectOptions) (result *ProjectDetails, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateProjectOptions, "updateProjectOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateProjectOptions, "updateProjectOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *updateProjectOptions.ProjectID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v2/projects/{project_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateProjectOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V2", "UpdateProject")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

	body := make(map[string]interface{})
	if updateProjectOptions.Name != nil {
		body["name"] = updateProjectOptions.Name
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
	response, err = discovery.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectDetails)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteProject : Delete a project
// Deletes the specified project.
//
// **Important:** Deleting a project deletes everything that is part of the specified project, including all
// collections.
func (discovery *DiscoveryV2) DeleteProject(deleteProjectOptions *DeleteProjectOptions) (response *core.DetailedResponse, err error) {
	return discovery.DeleteProjectWithContext(context.Background(), deleteProjectOptions)
}

// DeleteProjectWithContext is an alternate form of the DeleteProject method which supports a Context parameter
func (discovery *DiscoveryV2) DeleteProjectWithContext(ctx context.Context, deleteProjectOptions *DeleteProjectOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteProjectOptions, "deleteProjectOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteProjectOptions, "deleteProjectOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *deleteProjectOptions.ProjectID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v2/projects/{project_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteProjectOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V2", "DeleteProject")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = discovery.Service.Request(request, nil)

	return
}

// DeleteUserData : Delete labeled data
// Deletes all data associated with a specified customer ID. The method has no effect if no data is associated with the
// customer ID.
//
// You associate a customer ID with data by passing the **X-Watson-Metadata** header with a request that passes data.
// For more information about personal data and customer IDs, see [Information
// security](https://cloud.ibm.com/docs/discovery-data?topic=discovery-data-information-security#information-security).
//
// **Note:** This method is only supported on IBM Cloud instances of Discovery.
func (discovery *DiscoveryV2) DeleteUserData(deleteUserDataOptions *DeleteUserDataOptions) (response *core.DetailedResponse, err error) {
	return discovery.DeleteUserDataWithContext(context.Background(), deleteUserDataOptions)
}

// DeleteUserDataWithContext is an alternate form of the DeleteUserData method which supports a Context parameter
func (discovery *DiscoveryV2) DeleteUserDataWithContext(ctx context.Context, deleteUserDataOptions *DeleteUserDataOptions) (response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v2/user_data`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteUserDataOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V2", "DeleteUserData")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))
	builder.AddQuery("customer_id", fmt.Sprint(*deleteUserDataOptions.CustomerID))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = discovery.Service.Request(request, nil)

	return
}

// AddDocumentOptions : The AddDocument options.
type AddDocumentOptions struct {
	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// The ID of the collection.
	CollectionID *string `json:"collection_id" validate:"required,ne="`

	// The content of the document to ingest. The maximum supported file size when adding a file to a collection is 50
	// megabytes, the maximum supported file size when testing a configuration is 1 megabyte. Files larger than the
	// supported size are rejected.
	File io.ReadCloser `json:"file,omitempty"`

	// The filename for file.
	Filename *string `json:"filename,omitempty"`

	// The content type of file.
	FileContentType *string `json:"file_content_type,omitempty"`

	// The maximum supported metadata file size is 1 MB. Metadata parts larger than 1 MB are rejected.
	//
	//
	// Example:  ``` {
	//   "Creator": "Johnny Appleseed",
	//   "Subject": "Apples"
	// } ```.
	Metadata *string `json:"metadata,omitempty"`

	// When `true`, the uploaded document is added to the collection even if the data for that collection is shared with
	// other collections.
	XWatsonDiscoveryForce *bool `json:"X-Watson-Discovery-Force,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAddDocumentOptions : Instantiate AddDocumentOptions
func (*DiscoveryV2) NewAddDocumentOptions(projectID string, collectionID string) *AddDocumentOptions {
	return &AddDocumentOptions{
		ProjectID: core.StringPtr(projectID),
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

// AnalyzeDocumentOptions : The AnalyzeDocument options.
type AnalyzeDocumentOptions struct {
	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// The ID of the collection.
	CollectionID *string `json:"collection_id" validate:"required,ne="`

	// The content of the document to ingest. The maximum supported file size when adding a file to a collection is 50
	// megabytes, the maximum supported file size when testing a configuration is 1 megabyte. Files larger than the
	// supported size are rejected.
	File io.ReadCloser `json:"file,omitempty"`

	// The filename for file.
	Filename *string `json:"filename,omitempty"`

	// The content type of file.
	FileContentType *string `json:"file_content_type,omitempty"`

	// The maximum supported metadata file size is 1 MB. Metadata parts larger than 1 MB are rejected.
	//
	//
	// Example:  ``` {
	//   "Creator": "Johnny Appleseed",
	//   "Subject": "Apples"
	// } ```.
	Metadata *string `json:"metadata,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAnalyzeDocumentOptions : Instantiate AnalyzeDocumentOptions
func (*DiscoveryV2) NewAnalyzeDocumentOptions(projectID string, collectionID string) *AnalyzeDocumentOptions {
	return &AnalyzeDocumentOptions{
		ProjectID: core.StringPtr(projectID),
		CollectionID: core.StringPtr(collectionID),
	}
}

// SetProjectID : Allow user to set ProjectID
func (options *AnalyzeDocumentOptions) SetProjectID(projectID string) *AnalyzeDocumentOptions {
	options.ProjectID = core.StringPtr(projectID)
	return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *AnalyzeDocumentOptions) SetCollectionID(collectionID string) *AnalyzeDocumentOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
}

// SetFile : Allow user to set File
func (options *AnalyzeDocumentOptions) SetFile(file io.ReadCloser) *AnalyzeDocumentOptions {
	options.File = file
	return options
}

// SetFilename : Allow user to set Filename
func (options *AnalyzeDocumentOptions) SetFilename(filename string) *AnalyzeDocumentOptions {
	options.Filename = core.StringPtr(filename)
	return options
}

// SetFileContentType : Allow user to set FileContentType
func (options *AnalyzeDocumentOptions) SetFileContentType(fileContentType string) *AnalyzeDocumentOptions {
	options.FileContentType = core.StringPtr(fileContentType)
	return options
}

// SetMetadata : Allow user to set Metadata
func (options *AnalyzeDocumentOptions) SetMetadata(metadata string) *AnalyzeDocumentOptions {
	options.Metadata = core.StringPtr(metadata)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *AnalyzeDocumentOptions) SetHeaders(param map[string]string) *AnalyzeDocumentOptions {
	options.Headers = param
	return options
}

// AnalyzedDocument : An object containing the converted document and any identified enrichments.
type AnalyzedDocument struct {
	// Array of document results that match the query.
	Notices []Notice `json:"notices,omitempty"`

	// Result of the document analysis.
	Result *AnalyzedResult `json:"result,omitempty"`
}


// UnmarshalAnalyzedDocument unmarshals an instance of AnalyzedDocument from the specified map of raw messages.
func UnmarshalAnalyzedDocument(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AnalyzedDocument)
	err = core.UnmarshalModel(m, "notices", &obj.Notices, UnmarshalNotice)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalAnalyzedResult)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AnalyzedResult : Result of the document analysis.
type AnalyzedResult struct {
	// Metadata of the document.
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// Allows users to set arbitrary properties
	additionalProperties map[string]interface{}
}


// SetProperty allows the user to set an arbitrary property on an instance of AnalyzedResult
func (o *AnalyzedResult) SetProperty(key string, value interface{}) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string]interface{})
	}
	o.additionalProperties[key] = value
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of AnalyzedResult
func (o *AnalyzedResult) GetProperty(key string) interface{} {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of AnalyzedResult
func (o *AnalyzedResult) GetProperties() map[string]interface{} {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of AnalyzedResult
func (o *AnalyzedResult) MarshalJSON() (buffer []byte, err error) {
	m := make(map[string]interface{})
	if len(o.additionalProperties) > 0 {
		for k, v := range o.additionalProperties {
			m[k] = v
		}
	}
	if o.Metadata != nil {
		m["metadata"] = o.Metadata
	}
	buffer, err = json.Marshal(m)
	return
}

// UnmarshalAnalyzedResult unmarshals an instance of AnalyzedResult from the specified map of raw messages.
func UnmarshalAnalyzedResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AnalyzedResult)
	err = core.UnmarshalPrimitive(m, "metadata", &obj.Metadata)
	if err != nil {
		return
	}
	delete(m, "metadata")
	for k := range m {
		var v interface{}
		e := core.UnmarshalPrimitive(m, k, &v)
		if e != nil {
			err = e
			return
		}
		obj.SetProperty(k, v)
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Collection : A collection for storing documents.
type Collection struct {
	// The unique identifier of the collection.
	CollectionID *string `json:"collection_id,omitempty"`

	// The name of the collection.
	Name *string `json:"name,omitempty"`
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CollectionDetails : A collection for storing documents.
type CollectionDetails struct {
	// The unique identifier of the collection.
	CollectionID *string `json:"collection_id,omitempty"`

	// The name of the collection.
	Name *string `json:"name" validate:"required"`

	// A description of the collection.
	Description *string `json:"description,omitempty"`

	// The date that the collection was created.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// The language of the collection.
	Language *string `json:"language,omitempty"`

	// An array of enrichments that are applied to this collection.
	Enrichments []CollectionEnrichment `json:"enrichments,omitempty"`
}


// NewCollectionDetails : Instantiate CollectionDetails (Generic Model Constructor)
func (*DiscoveryV2) NewCollectionDetails(name string) (model *CollectionDetails, err error) {
	model = &CollectionDetails{
		Name: core.StringPtr(name),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalCollectionDetails unmarshals an instance of CollectionDetails from the specified map of raw messages.
func UnmarshalCollectionDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CollectionDetails)
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
	err = core.UnmarshalPrimitive(m, "language", &obj.Language)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "enrichments", &obj.Enrichments, UnmarshalCollectionEnrichment)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CollectionEnrichment : An object describing an Enrichment for a collection.
type CollectionEnrichment struct {
	// The unique identifier of this enrichment.
	EnrichmentID *string `json:"enrichment_id,omitempty"`

	// An array of field names that the enrichment is applied to.
	Fields []string `json:"fields,omitempty"`
}


// UnmarshalCollectionEnrichment unmarshals an instance of CollectionEnrichment from the specified map of raw messages.
func UnmarshalCollectionEnrichment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CollectionEnrichment)
	err = core.UnmarshalPrimitive(m, "enrichment_id", &obj.EnrichmentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "fields", &obj.Fields)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Completions : An object containing an array of autocompletion suggestions.
type Completions struct {
	// Array of autocomplete suggestion based on the provided prefix.
	Completions []string `json:"completions,omitempty"`
}


// UnmarshalCompletions unmarshals an instance of Completions from the specified map of raw messages.
func UnmarshalCompletions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Completions)
	err = core.UnmarshalPrimitive(m, "completions", &obj.Completions)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
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
	ComponentSettingsAggregationVisualizationTypeAutoConst = "auto"
	ComponentSettingsAggregationVisualizationTypeFacetTableConst = "facet_table"
	ComponentSettingsAggregationVisualizationTypeMapConst = "map"
	ComponentSettingsAggregationVisualizationTypeWordCloudConst = "word_cloud"
)


// UnmarshalComponentSettingsAggregation unmarshals an instance of ComponentSettingsAggregation from the specified map of raw messages.
func UnmarshalComponentSettingsAggregation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ComponentSettingsAggregation)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "label", &obj.Label)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "multiple_selections_allowed", &obj.MultipleSelectionsAllowed)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "visualization_type", &obj.VisualizationType)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ComponentSettingsFieldsShown : Fields shown in the results section of the UI.
type ComponentSettingsFieldsShown struct {
	// Body label.
	Body *ComponentSettingsFieldsShownBody `json:"body,omitempty"`

	// Title label.
	Title *ComponentSettingsFieldsShownTitle `json:"title,omitempty"`
}


// UnmarshalComponentSettingsFieldsShown unmarshals an instance of ComponentSettingsFieldsShown from the specified map of raw messages.
func UnmarshalComponentSettingsFieldsShown(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ComponentSettingsFieldsShown)
	err = core.UnmarshalModel(m, "body", &obj.Body, UnmarshalComponentSettingsFieldsShownBody)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "title", &obj.Title, UnmarshalComponentSettingsFieldsShownTitle)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ComponentSettingsFieldsShownBody : Body label.
type ComponentSettingsFieldsShownBody struct {
	// Use the whole passage as the body.
	UsePassage *bool `json:"use_passage,omitempty"`

	// Use a specific field as the title.
	Field *string `json:"field,omitempty"`
}


// UnmarshalComponentSettingsFieldsShownBody unmarshals an instance of ComponentSettingsFieldsShownBody from the specified map of raw messages.
func UnmarshalComponentSettingsFieldsShownBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ComponentSettingsFieldsShownBody)
	err = core.UnmarshalPrimitive(m, "use_passage", &obj.UsePassage)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "field", &obj.Field)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ComponentSettingsFieldsShownTitle : Title label.
type ComponentSettingsFieldsShownTitle struct {
	// Use a specific field as the title.
	Field *string `json:"field,omitempty"`
}


// UnmarshalComponentSettingsFieldsShownTitle unmarshals an instance of ComponentSettingsFieldsShownTitle from the specified map of raw messages.
func UnmarshalComponentSettingsFieldsShownTitle(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ComponentSettingsFieldsShownTitle)
	err = core.UnmarshalPrimitive(m, "field", &obj.Field)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ComponentSettingsResponse : The default component settings for this project.
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


// UnmarshalComponentSettingsResponse unmarshals an instance of ComponentSettingsResponse from the specified map of raw messages.
func UnmarshalComponentSettingsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ComponentSettingsResponse)
	err = core.UnmarshalModel(m, "fields_shown", &obj.FieldsShown, UnmarshalComponentSettingsFieldsShown)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "autocomplete", &obj.Autocomplete)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "structured_search", &obj.StructuredSearch)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "results_per_page", &obj.ResultsPerPage)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "aggregations", &obj.Aggregations, UnmarshalComponentSettingsAggregation)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateCollectionOptions : The CreateCollection options.
type CreateCollectionOptions struct {
	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// The name of the collection.
	Name *string `json:"name" validate:"required"`

	// A description of the collection.
	Description *string `json:"description,omitempty"`

	// The language of the collection.
	Language *string `json:"language,omitempty"`

	// An array of enrichments that are applied to this collection.
	Enrichments []CollectionEnrichment `json:"enrichments,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateCollectionOptions : Instantiate CreateCollectionOptions
func (*DiscoveryV2) NewCreateCollectionOptions(projectID string, name string) *CreateCollectionOptions {
	return &CreateCollectionOptions{
		ProjectID: core.StringPtr(projectID),
		Name: core.StringPtr(name),
	}
}

// SetProjectID : Allow user to set ProjectID
func (options *CreateCollectionOptions) SetProjectID(projectID string) *CreateCollectionOptions {
	options.ProjectID = core.StringPtr(projectID)
	return options
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

// SetLanguage : Allow user to set Language
func (options *CreateCollectionOptions) SetLanguage(language string) *CreateCollectionOptions {
	options.Language = core.StringPtr(language)
	return options
}

// SetEnrichments : Allow user to set Enrichments
func (options *CreateCollectionOptions) SetEnrichments(enrichments []CollectionEnrichment) *CreateCollectionOptions {
	options.Enrichments = enrichments
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateCollectionOptions) SetHeaders(param map[string]string) *CreateCollectionOptions {
	options.Headers = param
	return options
}

// CreateEnrichment : Information about a specific enrichment.
type CreateEnrichment struct {
	// The human readable name for this enrichment.
	Name *string `json:"name,omitempty"`

	// The description of this enrichment.
	Description *string `json:"description,omitempty"`

	// The type of this enrichment.
	Type *string `json:"type,omitempty"`

	// A object containing options for the current enrichment.
	Options *EnrichmentOptions `json:"options,omitempty"`
}

// Constants associated with the CreateEnrichment.Type property.
// The type of this enrichment.
const (
	CreateEnrichmentTypeDictionaryConst = "dictionary"
	CreateEnrichmentTypeRegularExpressionConst = "regular_expression"
	CreateEnrichmentTypeRuleBasedConst = "rule_based"
	CreateEnrichmentTypeUimaAnnotatorConst = "uima_annotator"
	CreateEnrichmentTypeWatsonKnowledgeStudioModelConst = "watson_knowledge_studio_model"
)


// UnmarshalCreateEnrichment unmarshals an instance of CreateEnrichment from the specified map of raw messages.
func UnmarshalCreateEnrichment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateEnrichment)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "options", &obj.Options, UnmarshalEnrichmentOptions)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateEnrichmentOptions : The CreateEnrichment options.
type CreateEnrichmentOptions struct {
	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// Information about a specific enrichment.
	Enrichment *CreateEnrichment `json:"enrichment" validate:"required"`

	// The enrichment file to upload.
	File io.ReadCloser `json:"file,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateEnrichmentOptions : Instantiate CreateEnrichmentOptions
func (*DiscoveryV2) NewCreateEnrichmentOptions(projectID string, enrichment *CreateEnrichment) *CreateEnrichmentOptions {
	return &CreateEnrichmentOptions{
		ProjectID: core.StringPtr(projectID),
		Enrichment: enrichment,
	}
}

// SetProjectID : Allow user to set ProjectID
func (options *CreateEnrichmentOptions) SetProjectID(projectID string) *CreateEnrichmentOptions {
	options.ProjectID = core.StringPtr(projectID)
	return options
}

// SetEnrichment : Allow user to set Enrichment
func (options *CreateEnrichmentOptions) SetEnrichment(enrichment *CreateEnrichment) *CreateEnrichmentOptions {
	options.Enrichment = enrichment
	return options
}

// SetFile : Allow user to set File
func (options *CreateEnrichmentOptions) SetFile(file io.ReadCloser) *CreateEnrichmentOptions {
	options.File = file
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateEnrichmentOptions) SetHeaders(param map[string]string) *CreateEnrichmentOptions {
	options.Headers = param
	return options
}

// CreateProjectOptions : The CreateProject options.
type CreateProjectOptions struct {
	// The human readable name of this project.
	Name *string `json:"name" validate:"required"`

	// The project type of this project.
	Type *string `json:"type" validate:"required"`

	// Default query parameters for this project.
	DefaultQueryParameters *DefaultQueryParams `json:"default_query_parameters,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateProjectOptions.Type property.
// The project type of this project.
const (
	CreateProjectOptionsTypeAnswerRetrievalConst = "answer_retrieval"
	CreateProjectOptionsTypeContentMiningConst = "content_mining"
	CreateProjectOptionsTypeDocumentRetrievalConst = "document_retrieval"
	CreateProjectOptionsTypeOtherConst = "other"
)

// NewCreateProjectOptions : Instantiate CreateProjectOptions
func (*DiscoveryV2) NewCreateProjectOptions(name string, typeVar string) *CreateProjectOptions {
	return &CreateProjectOptions{
		Name: core.StringPtr(name),
		Type: core.StringPtr(typeVar),
	}
}

// SetName : Allow user to set Name
func (options *CreateProjectOptions) SetName(name string) *CreateProjectOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetType : Allow user to set Type
func (options *CreateProjectOptions) SetType(typeVar string) *CreateProjectOptions {
	options.Type = core.StringPtr(typeVar)
	return options
}

// SetDefaultQueryParameters : Allow user to set DefaultQueryParameters
func (options *CreateProjectOptions) SetDefaultQueryParameters(defaultQueryParameters *DefaultQueryParams) *CreateProjectOptions {
	options.DefaultQueryParameters = defaultQueryParameters
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateProjectOptions) SetHeaders(param map[string]string) *CreateProjectOptions {
	options.Headers = param
	return options
}

// CreateTrainingQueryOptions : The CreateTrainingQuery options.
type CreateTrainingQueryOptions struct {
	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// The natural text query for the training query.
	NaturalLanguageQuery *string `json:"natural_language_query" validate:"required"`

	// Array of training examples.
	Examples []TrainingExample `json:"examples" validate:"required"`

	// The filter used on the collection before the **natural_language_query** is applied.
	Filter *string `json:"filter,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateTrainingQueryOptions : Instantiate CreateTrainingQueryOptions
func (*DiscoveryV2) NewCreateTrainingQueryOptions(projectID string, naturalLanguageQuery string, examples []TrainingExample) *CreateTrainingQueryOptions {
	return &CreateTrainingQueryOptions{
		ProjectID: core.StringPtr(projectID),
		NaturalLanguageQuery: core.StringPtr(naturalLanguageQuery),
		Examples: examples,
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

// DefaultQueryParams : Default query parameters for this project.
type DefaultQueryParams struct {
	// An array of collection identifiers to query. If empty or omitted all collections in the project are queried.
	CollectionIds []string `json:"collection_ids,omitempty"`

	// Default settings configuration for passage search options.
	Passages *DefaultQueryParamsPassages `json:"passages,omitempty"`

	// Default project query settings for table results.
	TableResults *DefaultQueryParamsTableResults `json:"table_results,omitempty"`

	// A string representing the default aggregation query for the project.
	Aggregation *string `json:"aggregation,omitempty"`

	// Object containing suggested refinement settings.
	SuggestedRefinements *DefaultQueryParamsSuggestedRefinements `json:"suggested_refinements,omitempty"`

	// When `true`, a spelling suggestions for the query are returned by default.
	SpellingSuggestions *bool `json:"spelling_suggestions,omitempty"`

	// When `true`, a highlights for the query are returned by default.
	Highlight *bool `json:"highlight,omitempty"`

	// The number of document results returned by default.
	Count *int64 `json:"count,omitempty"`

	// A comma separated list of document fields to sort results by default.
	Sort *string `json:"sort,omitempty"`

	// An array of field names to return in document results if present by default.
	Return []string `json:"return,omitempty"`
}


// UnmarshalDefaultQueryParams unmarshals an instance of DefaultQueryParams from the specified map of raw messages.
func UnmarshalDefaultQueryParams(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DefaultQueryParams)
	err = core.UnmarshalPrimitive(m, "collection_ids", &obj.CollectionIds)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "passages", &obj.Passages, UnmarshalDefaultQueryParamsPassages)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "table_results", &obj.TableResults, UnmarshalDefaultQueryParamsTableResults)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "aggregation", &obj.Aggregation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "suggested_refinements", &obj.SuggestedRefinements, UnmarshalDefaultQueryParamsSuggestedRefinements)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spelling_suggestions", &obj.SpellingSuggestions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "highlight", &obj.Highlight)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "count", &obj.Count)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sort", &obj.Sort)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "return", &obj.Return)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DefaultQueryParamsPassages : Default settings configuration for passage search options.
type DefaultQueryParamsPassages struct {
	// When `true`, a passage search is performed by default.
	Enabled *bool `json:"enabled,omitempty"`

	// The number of passages to return.
	Count *int64 `json:"count,omitempty"`

	// An array of field names to perform the passage search on.
	Fields []string `json:"fields,omitempty"`

	// The approximate number of characters that each returned passage will contain.
	Characters *int64 `json:"characters,omitempty"`

	// When `true` the number of passages that can be returned from a single document is restricted to the
	// *max_per_document* value.
	PerDocument *bool `json:"per_document,omitempty"`

	// The default maximum number of passages that can be taken from a single document as the result of a passage query.
	MaxPerDocument *int64 `json:"max_per_document,omitempty"`
}


// UnmarshalDefaultQueryParamsPassages unmarshals an instance of DefaultQueryParamsPassages from the specified map of raw messages.
func UnmarshalDefaultQueryParamsPassages(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DefaultQueryParamsPassages)
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "count", &obj.Count)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "fields", &obj.Fields)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "characters", &obj.Characters)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "per_document", &obj.PerDocument)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "max_per_document", &obj.MaxPerDocument)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DefaultQueryParamsSuggestedRefinements : Object containing suggested refinement settings.
type DefaultQueryParamsSuggestedRefinements struct {
	// When `true`, a suggested refinements for the query are returned by default.
	Enabled *bool `json:"enabled,omitempty"`

	// The number of suggested refinements to return by default.
	Count *int64 `json:"count,omitempty"`
}


// UnmarshalDefaultQueryParamsSuggestedRefinements unmarshals an instance of DefaultQueryParamsSuggestedRefinements from the specified map of raw messages.
func UnmarshalDefaultQueryParamsSuggestedRefinements(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DefaultQueryParamsSuggestedRefinements)
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
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

// DefaultQueryParamsTableResults : Default project query settings for table results.
type DefaultQueryParamsTableResults struct {
	// When `true`, a table results for the query are returned by default.
	Enabled *bool `json:"enabled,omitempty"`

	// The number of table results to return by default.
	Count *int64 `json:"count,omitempty"`

	// The number of table results to include in each result document.
	PerDocument *int64 `json:"per_document,omitempty"`
}


// UnmarshalDefaultQueryParamsTableResults unmarshals an instance of DefaultQueryParamsTableResults from the specified map of raw messages.
func UnmarshalDefaultQueryParamsTableResults(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DefaultQueryParamsTableResults)
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "count", &obj.Count)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "per_document", &obj.PerDocument)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteCollectionOptions : The DeleteCollection options.
type DeleteCollectionOptions struct {
	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// The ID of the collection.
	CollectionID *string `json:"collection_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteCollectionOptions : Instantiate DeleteCollectionOptions
func (*DiscoveryV2) NewDeleteCollectionOptions(projectID string, collectionID string) *DeleteCollectionOptions {
	return &DeleteCollectionOptions{
		ProjectID: core.StringPtr(projectID),
		CollectionID: core.StringPtr(collectionID),
	}
}

// SetProjectID : Allow user to set ProjectID
func (options *DeleteCollectionOptions) SetProjectID(projectID string) *DeleteCollectionOptions {
	options.ProjectID = core.StringPtr(projectID)
	return options
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

// DeleteDocumentOptions : The DeleteDocument options.
type DeleteDocumentOptions struct {
	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// The ID of the collection.
	CollectionID *string `json:"collection_id" validate:"required,ne="`

	// The ID of the document.
	DocumentID *string `json:"document_id" validate:"required,ne="`

	// When `true`, the uploaded document is added to the collection even if the data for that collection is shared with
	// other collections.
	XWatsonDiscoveryForce *bool `json:"X-Watson-Discovery-Force,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteDocumentOptions : Instantiate DeleteDocumentOptions
func (*DiscoveryV2) NewDeleteDocumentOptions(projectID string, collectionID string, documentID string) *DeleteDocumentOptions {
	return &DeleteDocumentOptions{
		ProjectID: core.StringPtr(projectID),
		CollectionID: core.StringPtr(collectionID),
		DocumentID: core.StringPtr(documentID),
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
	DeleteDocumentResponseStatusDeletedConst = "deleted"
)


// UnmarshalDeleteDocumentResponse unmarshals an instance of DeleteDocumentResponse from the specified map of raw messages.
func UnmarshalDeleteDocumentResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteDocumentResponse)
	err = core.UnmarshalPrimitive(m, "document_id", &obj.DocumentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteEnrichmentOptions : The DeleteEnrichment options.
type DeleteEnrichmentOptions struct {
	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// The ID of the enrichment.
	EnrichmentID *string `json:"enrichment_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteEnrichmentOptions : Instantiate DeleteEnrichmentOptions
func (*DiscoveryV2) NewDeleteEnrichmentOptions(projectID string, enrichmentID string) *DeleteEnrichmentOptions {
	return &DeleteEnrichmentOptions{
		ProjectID: core.StringPtr(projectID),
		EnrichmentID: core.StringPtr(enrichmentID),
	}
}

// SetProjectID : Allow user to set ProjectID
func (options *DeleteEnrichmentOptions) SetProjectID(projectID string) *DeleteEnrichmentOptions {
	options.ProjectID = core.StringPtr(projectID)
	return options
}

// SetEnrichmentID : Allow user to set EnrichmentID
func (options *DeleteEnrichmentOptions) SetEnrichmentID(enrichmentID string) *DeleteEnrichmentOptions {
	options.EnrichmentID = core.StringPtr(enrichmentID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteEnrichmentOptions) SetHeaders(param map[string]string) *DeleteEnrichmentOptions {
	options.Headers = param
	return options
}

// DeleteProjectOptions : The DeleteProject options.
type DeleteProjectOptions struct {
	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteProjectOptions : Instantiate DeleteProjectOptions
func (*DiscoveryV2) NewDeleteProjectOptions(projectID string) *DeleteProjectOptions {
	return &DeleteProjectOptions{
		ProjectID: core.StringPtr(projectID),
	}
}

// SetProjectID : Allow user to set ProjectID
func (options *DeleteProjectOptions) SetProjectID(projectID string) *DeleteProjectOptions {
	options.ProjectID = core.StringPtr(projectID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteProjectOptions) SetHeaders(param map[string]string) *DeleteProjectOptions {
	options.Headers = param
	return options
}

// DeleteTrainingQueriesOptions : The DeleteTrainingQueries options.
type DeleteTrainingQueriesOptions struct {
	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteTrainingQueriesOptions : Instantiate DeleteTrainingQueriesOptions
func (*DiscoveryV2) NewDeleteTrainingQueriesOptions(projectID string) *DeleteTrainingQueriesOptions {
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

// DeleteUserDataOptions : The DeleteUserData options.
type DeleteUserDataOptions struct {
	// The customer ID for which all data is to be deleted.
	CustomerID *string `json:"customer_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteUserDataOptions : Instantiate DeleteUserDataOptions
func (*DiscoveryV2) NewDeleteUserDataOptions(customerID string) *DeleteUserDataOptions {
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
	DocumentAcceptedStatusPendingConst = "pending"
	DocumentAcceptedStatusProcessingConst = "processing"
)


// UnmarshalDocumentAccepted unmarshals an instance of DocumentAccepted from the specified map of raw messages.
func UnmarshalDocumentAccepted(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DocumentAccepted)
	err = core.UnmarshalPrimitive(m, "document_id", &obj.DocumentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

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


// UnmarshalDocumentAttribute unmarshals an instance of DocumentAttribute from the specified map of raw messages.
func UnmarshalDocumentAttribute(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DocumentAttribute)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "location", &obj.Location, UnmarshalTableElementLocation)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Enrichment : Information about a specific enrichment.
type Enrichment struct {
	// The unique identifier of this enrichment.
	EnrichmentID *string `json:"enrichment_id,omitempty"`

	// The human readable name for this enrichment.
	Name *string `json:"name,omitempty"`

	// The description of this enrichment.
	Description *string `json:"description,omitempty"`

	// The type of this enrichment.
	Type *string `json:"type,omitempty"`

	// A object containing options for the current enrichment.
	Options *EnrichmentOptions `json:"options,omitempty"`
}

// Constants associated with the Enrichment.Type property.
// The type of this enrichment.
const (
	EnrichmentTypeDictionaryConst = "dictionary"
	EnrichmentTypeNaturalLanguageUnderstandingConst = "natural_language_understanding"
	EnrichmentTypePartOfSpeechConst = "part_of_speech"
	EnrichmentTypeRegularExpressionConst = "regular_expression"
	EnrichmentTypeRuleBasedConst = "rule_based"
	EnrichmentTypeSentimentConst = "sentiment"
	EnrichmentTypeUimaAnnotatorConst = "uima_annotator"
	EnrichmentTypeWatsonKnowledgeStudioModelConst = "watson_knowledge_studio_model"
)


// UnmarshalEnrichment unmarshals an instance of Enrichment from the specified map of raw messages.
func UnmarshalEnrichment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Enrichment)
	err = core.UnmarshalPrimitive(m, "enrichment_id", &obj.EnrichmentID)
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
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "options", &obj.Options, UnmarshalEnrichmentOptions)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EnrichmentOptions : A object containing options for the current enrichment.
type EnrichmentOptions struct {
	// An array of supported languages for this enrichment.
	Languages []string `json:"languages,omitempty"`

	// The type of entity. Required when creating `dictionary` and `regular_expression` **type** enrichment. Not valid when
	// creating any other type of enrichment.
	EntityType *string `json:"entity_type,omitempty"`

	// The regular expression to apply for this enrichment. Required only when the **type** of enrichment being created is
	// a `regular_expression`. Not valid when creating any other type of enrichment.
	RegularExpression *string `json:"regular_expression,omitempty"`

	// The name of the result document field that this enrichment creates. Required only when the enrichment **type** is
	// `rule_based`. Not valid when creating any other type of enrichment.
	ResultField *string `json:"result_field,omitempty"`
}


// UnmarshalEnrichmentOptions unmarshals an instance of EnrichmentOptions from the specified map of raw messages.
func UnmarshalEnrichmentOptions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EnrichmentOptions)
	err = core.UnmarshalPrimitive(m, "languages", &obj.Languages)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "entity_type", &obj.EntityType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "regular_expression", &obj.RegularExpression)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "result_field", &obj.ResultField)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Enrichments : An object containing an array of enrichment definitions.
type Enrichments struct {
	// An array of enrichment definitions.
	Enrichments []Enrichment `json:"enrichments,omitempty"`
}


// UnmarshalEnrichments unmarshals an instance of Enrichments from the specified map of raw messages.
func UnmarshalEnrichments(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Enrichments)
	err = core.UnmarshalModel(m, "enrichments", &obj.Enrichments, UnmarshalEnrichment)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
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
	FieldTypeBinaryConst = "binary"
	FieldTypeBooleanConst = "boolean"
	FieldTypeByteConst = "byte"
	FieldTypeDateConst = "date"
	FieldTypeDoubleConst = "double"
	FieldTypeFloatConst = "float"
	FieldTypeIntegerConst = "integer"
	FieldTypeLongConst = "long"
	FieldTypeNestedConst = "nested"
	FieldTypeShortConst = "short"
	FieldTypeStringConst = "string"
)


// UnmarshalField unmarshals an instance of Field from the specified map of raw messages.
func UnmarshalField(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Field)
	err = core.UnmarshalPrimitive(m, "field", &obj.Field)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "collection_id", &obj.CollectionID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetAutocompletionOptions : The GetAutocompletion options.
type GetAutocompletionOptions struct {
	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required,ne="`

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

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAutocompletionOptions : Instantiate GetAutocompletionOptions
func (*DiscoveryV2) NewGetAutocompletionOptions(projectID string, prefix string) *GetAutocompletionOptions {
	return &GetAutocompletionOptions{
		ProjectID: core.StringPtr(projectID),
		Prefix: core.StringPtr(prefix),
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

// GetCollectionOptions : The GetCollection options.
type GetCollectionOptions struct {
	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// The ID of the collection.
	CollectionID *string `json:"collection_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCollectionOptions : Instantiate GetCollectionOptions
func (*DiscoveryV2) NewGetCollectionOptions(projectID string, collectionID string) *GetCollectionOptions {
	return &GetCollectionOptions{
		ProjectID: core.StringPtr(projectID),
		CollectionID: core.StringPtr(collectionID),
	}
}

// SetProjectID : Allow user to set ProjectID
func (options *GetCollectionOptions) SetProjectID(projectID string) *GetCollectionOptions {
	options.ProjectID = core.StringPtr(projectID)
	return options
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

// GetComponentSettingsOptions : The GetComponentSettings options.
type GetComponentSettingsOptions struct {
	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetComponentSettingsOptions : Instantiate GetComponentSettingsOptions
func (*DiscoveryV2) NewGetComponentSettingsOptions(projectID string) *GetComponentSettingsOptions {
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

// GetEnrichmentOptions : The GetEnrichment options.
type GetEnrichmentOptions struct {
	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// The ID of the enrichment.
	EnrichmentID *string `json:"enrichment_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetEnrichmentOptions : Instantiate GetEnrichmentOptions
func (*DiscoveryV2) NewGetEnrichmentOptions(projectID string, enrichmentID string) *GetEnrichmentOptions {
	return &GetEnrichmentOptions{
		ProjectID: core.StringPtr(projectID),
		EnrichmentID: core.StringPtr(enrichmentID),
	}
}

// SetProjectID : Allow user to set ProjectID
func (options *GetEnrichmentOptions) SetProjectID(projectID string) *GetEnrichmentOptions {
	options.ProjectID = core.StringPtr(projectID)
	return options
}

// SetEnrichmentID : Allow user to set EnrichmentID
func (options *GetEnrichmentOptions) SetEnrichmentID(enrichmentID string) *GetEnrichmentOptions {
	options.EnrichmentID = core.StringPtr(enrichmentID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetEnrichmentOptions) SetHeaders(param map[string]string) *GetEnrichmentOptions {
	options.Headers = param
	return options
}

// GetProjectOptions : The GetProject options.
type GetProjectOptions struct {
	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetProjectOptions : Instantiate GetProjectOptions
func (*DiscoveryV2) NewGetProjectOptions(projectID string) *GetProjectOptions {
	return &GetProjectOptions{
		ProjectID: core.StringPtr(projectID),
	}
}

// SetProjectID : Allow user to set ProjectID
func (options *GetProjectOptions) SetProjectID(projectID string) *GetProjectOptions {
	options.ProjectID = core.StringPtr(projectID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetProjectOptions) SetHeaders(param map[string]string) *GetProjectOptions {
	options.Headers = param
	return options
}

// GetTrainingQueryOptions : The GetTrainingQuery options.
type GetTrainingQueryOptions struct {
	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// The ID of the query used for training.
	QueryID *string `json:"query_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetTrainingQueryOptions : Instantiate GetTrainingQueryOptions
func (*DiscoveryV2) NewGetTrainingQueryOptions(projectID string, queryID string) *GetTrainingQueryOptions {
	return &GetTrainingQueryOptions{
		ProjectID: core.StringPtr(projectID),
		QueryID: core.StringPtr(queryID),
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
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListCollectionsOptions : Instantiate ListCollectionsOptions
func (*DiscoveryV2) NewListCollectionsOptions(projectID string) *ListCollectionsOptions {
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


// UnmarshalListCollectionsResponse unmarshals an instance of ListCollectionsResponse from the specified map of raw messages.
func UnmarshalListCollectionsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListCollectionsResponse)
	err = core.UnmarshalModel(m, "collections", &obj.Collections, UnmarshalCollection)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListEnrichmentsOptions : The ListEnrichments options.
type ListEnrichmentsOptions struct {
	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListEnrichmentsOptions : Instantiate ListEnrichmentsOptions
func (*DiscoveryV2) NewListEnrichmentsOptions(projectID string) *ListEnrichmentsOptions {
	return &ListEnrichmentsOptions{
		ProjectID: core.StringPtr(projectID),
	}
}

// SetProjectID : Allow user to set ProjectID
func (options *ListEnrichmentsOptions) SetProjectID(projectID string) *ListEnrichmentsOptions {
	options.ProjectID = core.StringPtr(projectID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListEnrichmentsOptions) SetHeaders(param map[string]string) *ListEnrichmentsOptions {
	options.Headers = param
	return options
}

// ListFieldsOptions : The ListFields options.
type ListFieldsOptions struct {
	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// Comma separated list of the collection IDs. If this parameter is not specified, all collections in the project are
	// used.
	CollectionIds []string `json:"collection_ids,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListFieldsOptions : Instantiate ListFieldsOptions
func (*DiscoveryV2) NewListFieldsOptions(projectID string) *ListFieldsOptions {
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


// UnmarshalListFieldsResponse unmarshals an instance of ListFieldsResponse from the specified map of raw messages.
func UnmarshalListFieldsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListFieldsResponse)
	err = core.UnmarshalModel(m, "fields", &obj.Fields, UnmarshalField)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListProjectsOptions : The ListProjects options.
type ListProjectsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListProjectsOptions : Instantiate ListProjectsOptions
func (*DiscoveryV2) NewListProjectsOptions() *ListProjectsOptions {
	return &ListProjectsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListProjectsOptions) SetHeaders(param map[string]string) *ListProjectsOptions {
	options.Headers = param
	return options
}

// ListProjectsResponse : A list of projects in this instance.
type ListProjectsResponse struct {
	// An array of project details.
	Projects []ProjectListDetails `json:"projects,omitempty"`
}


// UnmarshalListProjectsResponse unmarshals an instance of ListProjectsResponse from the specified map of raw messages.
func UnmarshalListProjectsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListProjectsResponse)
	err = core.UnmarshalModel(m, "projects", &obj.Projects, UnmarshalProjectListDetails)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListTrainingQueriesOptions : The ListTrainingQueries options.
type ListTrainingQueriesOptions struct {
	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListTrainingQueriesOptions : Instantiate ListTrainingQueriesOptions
func (*DiscoveryV2) NewListTrainingQueriesOptions(projectID string) *ListTrainingQueriesOptions {
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
	NoticeSeverityErrorConst = "error"
	NoticeSeverityWarningConst = "warning"
)


// UnmarshalNotice unmarshals an instance of Notice from the specified map of raw messages.
func UnmarshalNotice(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Notice)
	err = core.UnmarshalPrimitive(m, "notice_id", &obj.NoticeID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created", &obj.Created)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "document_id", &obj.DocumentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "collection_id", &obj.CollectionID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "query_id", &obj.QueryID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "severity", &obj.Severity)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "step", &obj.Step)
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

// ProjectDetails : Detailed information about the specified project.
type ProjectDetails struct {
	// The unique identifier of this project.
	ProjectID *string `json:"project_id,omitempty"`

	// The human readable name of this project.
	Name *string `json:"name,omitempty"`

	// The project type of this project.
	Type *string `json:"type,omitempty"`

	// Relevancy training status information for this project.
	RelevancyTrainingStatus *ProjectListDetailsRelevancyTrainingStatus `json:"relevancy_training_status,omitempty"`

	// The number of collections configured in this project.
	CollectionCount *int64 `json:"collection_count,omitempty"`

	// Default query parameters for this project.
	DefaultQueryParameters *DefaultQueryParams `json:"default_query_parameters,omitempty"`
}

// Constants associated with the ProjectDetails.Type property.
// The project type of this project.
const (
	ProjectDetailsTypeAnswerRetrievalConst = "answer_retrieval"
	ProjectDetailsTypeContentMiningConst = "content_mining"
	ProjectDetailsTypeDocumentRetrievalConst = "document_retrieval"
	ProjectDetailsTypeOtherConst = "other"
)


// UnmarshalProjectDetails unmarshals an instance of ProjectDetails from the specified map of raw messages.
func UnmarshalProjectDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectDetails)
	err = core.UnmarshalPrimitive(m, "project_id", &obj.ProjectID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "relevancy_training_status", &obj.RelevancyTrainingStatus, UnmarshalProjectListDetailsRelevancyTrainingStatus)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "collection_count", &obj.CollectionCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "default_query_parameters", &obj.DefaultQueryParameters, UnmarshalDefaultQueryParams)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectListDetails : Details about a specific project.
type ProjectListDetails struct {
	// The unique identifier of this project.
	ProjectID *string `json:"project_id,omitempty"`

	// The human readable name of this project.
	Name *string `json:"name,omitempty"`

	// The project type of this project.
	Type *string `json:"type,omitempty"`

	// Relevancy training status information for this project.
	RelevancyTrainingStatus *ProjectListDetailsRelevancyTrainingStatus `json:"relevancy_training_status,omitempty"`

	// The number of collections configured in this project.
	CollectionCount *int64 `json:"collection_count,omitempty"`
}

// Constants associated with the ProjectListDetails.Type property.
// The project type of this project.
const (
	ProjectListDetailsTypeAnswerRetrievalConst = "answer_retrieval"
	ProjectListDetailsTypeContentMiningConst = "content_mining"
	ProjectListDetailsTypeDocumentRetrievalConst = "document_retrieval"
	ProjectListDetailsTypeOtherConst = "other"
)


// UnmarshalProjectListDetails unmarshals an instance of ProjectListDetails from the specified map of raw messages.
func UnmarshalProjectListDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectListDetails)
	err = core.UnmarshalPrimitive(m, "project_id", &obj.ProjectID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "relevancy_training_status", &obj.RelevancyTrainingStatus, UnmarshalProjectListDetailsRelevancyTrainingStatus)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "collection_count", &obj.CollectionCount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectListDetailsRelevancyTrainingStatus : Relevancy training status information for this project.
type ProjectListDetailsRelevancyTrainingStatus struct {
	// When the training data was updated.
	DataUpdated *string `json:"data_updated,omitempty"`

	// The total number of examples.
	TotalExamples *int64 `json:"total_examples,omitempty"`

	// When `true`, sufficient label diversity is present to allow training for this project.
	SufficientLabelDiversity *bool `json:"sufficient_label_diversity,omitempty"`

	// When `true`, the relevancy training is in processing.
	Processing *bool `json:"processing,omitempty"`

	// When `true`, the minimum number of examples required to train has been met.
	MinimumExamplesAdded *bool `json:"minimum_examples_added,omitempty"`

	// The time that the most recent successful training occurred.
	SuccessfullyTrained *string `json:"successfully_trained,omitempty"`

	// When `true`, relevancy training is available when querying collections in the project.
	Available *bool `json:"available,omitempty"`

	// The number of notices generated during the relevancy training.
	Notices *int64 `json:"notices,omitempty"`

	// When `true`, the minimum number of queries required to train has been met.
	MinimumQueriesAdded *bool `json:"minimum_queries_added,omitempty"`
}


// UnmarshalProjectListDetailsRelevancyTrainingStatus unmarshals an instance of ProjectListDetailsRelevancyTrainingStatus from the specified map of raw messages.
func UnmarshalProjectListDetailsRelevancyTrainingStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectListDetailsRelevancyTrainingStatus)
	err = core.UnmarshalPrimitive(m, "data_updated", &obj.DataUpdated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_examples", &obj.TotalExamples)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sufficient_label_diversity", &obj.SufficientLabelDiversity)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "processing", &obj.Processing)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "minimum_examples_added", &obj.MinimumExamplesAdded)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "successfully_trained", &obj.SuccessfullyTrained)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "available", &obj.Available)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "notices", &obj.Notices)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "minimum_queries_added", &obj.MinimumQueriesAdded)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// QueryAggregation : An abstract aggregation type produced by Discovery to analyze the input provided.
type QueryAggregation struct {
	// The type of aggregation command used. Options include: term, histogram, timeslice, nested, filter, min, max, sum,
	// average, unique_count, and top_hits.
	Type *string `json:"type" validate:"required"`
}

func (*QueryAggregation) isaQueryAggregation() bool {
	return true
}

type QueryAggregationIntf interface {
	isaQueryAggregation() bool
}

// UnmarshalQueryAggregation unmarshals an instance of QueryAggregation from the specified map of raw messages.
func UnmarshalQueryAggregation(m map[string]json.RawMessage, result interface{}) (err error) {
	// Retrieve discriminator value to determine correct "subclass".
	var discValue string
	err = core.UnmarshalPrimitive(m, "type", &discValue)
	if err != nil {
		err = fmt.Errorf("error unmarshalling discriminator property 'type': %s", err.Error())
		return
	}
	if discValue == "" {
		err = fmt.Errorf("required discriminator property 'type' not found in JSON object")
		return
	}
	if discValue == "term" {
		err = core.UnmarshalModel(m, "", result, UnmarshalQueryTermAggregation)
	} else if discValue == "histogram" {
		err = core.UnmarshalModel(m, "", result, UnmarshalQueryHistogramAggregation)
	} else if discValue == "timeslice" {
		err = core.UnmarshalModel(m, "", result, UnmarshalQueryTimesliceAggregation)
	} else if discValue == "nested" {
		err = core.UnmarshalModel(m, "", result, UnmarshalQueryNestedAggregation)
	} else if discValue == "filter" {
		err = core.UnmarshalModel(m, "", result, UnmarshalQueryFilterAggregation)
	} else if discValue == "min" {
		err = core.UnmarshalModel(m, "", result, UnmarshalQueryCalculationAggregation)
	} else if discValue == "max" {
		err = core.UnmarshalModel(m, "", result, UnmarshalQueryCalculationAggregation)
	} else if discValue == "sum" {
		err = core.UnmarshalModel(m, "", result, UnmarshalQueryCalculationAggregation)
	} else if discValue == "average" {
		err = core.UnmarshalModel(m, "", result, UnmarshalQueryCalculationAggregation)
	} else if discValue == "unique_count" {
		err = core.UnmarshalModel(m, "", result, UnmarshalQueryCalculationAggregation)
	} else if discValue == "top_hits" {
		err = core.UnmarshalModel(m, "", result, UnmarshalQueryTopHitsAggregation)
	} else if discValue == "group_by" {
		err = core.UnmarshalModel(m, "", result, UnmarshalQueryGroupByAggregation)
	} else {
		err = fmt.Errorf("unrecognized value for discriminator property 'type': %s", discValue)
	}
	return
}

// QueryGroupByAggregationResult : Top value result for the term aggregation.
type QueryGroupByAggregationResult struct {
	// Value of the field with a non-zero frequency in the document set.
	Key *string `json:"key" validate:"required"`

	// Number of documents containing the 'key'.
	MatchingResults *int64 `json:"matching_results" validate:"required"`

	// The relevancy for this group.
	Relevancy *float64 `json:"relevancy,omitempty"`

	// The number of documents which have the group as the value of specified field in the whole set of documents in this
	// collection. Returned only when the `relevancy` parameter is set to `true`.
	TotalMatchingDocuments *int64 `json:"total_matching_documents,omitempty"`

	// The estimated number of documents which would match the query and also meet the condition. Returned only when the
	// `relevancy` parameter is set to `true`.
	EstimatedMatchingDocuments *int64 `json:"estimated_matching_documents,omitempty"`

	// An array of sub aggregations.
	Aggregations []QueryAggregationIntf `json:"aggregations,omitempty"`
}


// UnmarshalQueryGroupByAggregationResult unmarshals an instance of QueryGroupByAggregationResult from the specified map of raw messages.
func UnmarshalQueryGroupByAggregationResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(QueryGroupByAggregationResult)
	err = core.UnmarshalPrimitive(m, "key", &obj.Key)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "matching_results", &obj.MatchingResults)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "relevancy", &obj.Relevancy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_matching_documents", &obj.TotalMatchingDocuments)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "estimated_matching_documents", &obj.EstimatedMatchingDocuments)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "aggregations", &obj.Aggregations, UnmarshalQueryAggregation)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// QueryHistogramAggregationResult : Histogram numeric interval result.
type QueryHistogramAggregationResult struct {
	// The value of the upper bound for the numeric segment.
	Key *int64 `json:"key" validate:"required"`

	// Number of documents with the specified key as the upper bound.
	MatchingResults *int64 `json:"matching_results" validate:"required"`

	// An array of sub aggregations.
	Aggregations []QueryAggregationIntf `json:"aggregations,omitempty"`
}


// UnmarshalQueryHistogramAggregationResult unmarshals an instance of QueryHistogramAggregationResult from the specified map of raw messages.
func UnmarshalQueryHistogramAggregationResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(QueryHistogramAggregationResult)
	err = core.UnmarshalPrimitive(m, "key", &obj.Key)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "matching_results", &obj.MatchingResults)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "aggregations", &obj.Aggregations, UnmarshalQueryAggregation)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
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
	// maximum is `100`.
	Count *int64 `json:"count,omitempty"`

	// The approximate number of characters that any one passage will have.
	Characters *int64 `json:"characters,omitempty"`
}


// UnmarshalQueryLargePassages unmarshals an instance of QueryLargePassages from the specified map of raw messages.
func UnmarshalQueryLargePassages(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(QueryLargePassages)
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "per_document", &obj.PerDocument)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "max_per_document", &obj.MaxPerDocument)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "fields", &obj.Fields)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "count", &obj.Count)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "characters", &obj.Characters)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// QueryLargeSuggestedRefinements : Configuration for suggested refinements.
type QueryLargeSuggestedRefinements struct {
	// Whether to perform suggested refinements.
	Enabled *bool `json:"enabled,omitempty"`

	// Maximum number of suggested refinements texts to be returned. The maximum is `100`.
	Count *int64 `json:"count,omitempty"`
}


// UnmarshalQueryLargeSuggestedRefinements unmarshals an instance of QueryLargeSuggestedRefinements from the specified map of raw messages.
func UnmarshalQueryLargeSuggestedRefinements(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(QueryLargeSuggestedRefinements)
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
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

// QueryLargeTableResults : Configuration for table retrieval.
type QueryLargeTableResults struct {
	// Whether to enable table retrieval.
	Enabled *bool `json:"enabled,omitempty"`

	// Maximum number of tables to return.
	Count *int64 `json:"count,omitempty"`
}


// UnmarshalQueryLargeTableResults unmarshals an instance of QueryLargeTableResults from the specified map of raw messages.
func UnmarshalQueryLargeTableResults(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(QueryLargeTableResults)
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
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

// QueryNoticesOptions : The QueryNotices options.
type QueryNoticesOptions struct {
	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required,ne="`

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

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewQueryNoticesOptions : Instantiate QueryNoticesOptions
func (*DiscoveryV2) NewQueryNoticesOptions(projectID string) *QueryNoticesOptions {
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


// UnmarshalQueryNoticesResponse unmarshals an instance of QueryNoticesResponse from the specified map of raw messages.
func UnmarshalQueryNoticesResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(QueryNoticesResponse)
	err = core.UnmarshalPrimitive(m, "matching_results", &obj.MatchingResults)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "notices", &obj.Notices, UnmarshalNotice)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// QueryOptions : The Query options.
type QueryOptions struct {
	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required,ne="`

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

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewQueryOptions : Instantiate QueryOptions
func (*DiscoveryV2) NewQueryOptions(projectID string) *QueryOptions {
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
	Aggregations []QueryAggregationIntf `json:"aggregations,omitempty"`

	// An object contain retrieval type information.
	RetrievalDetails *RetrievalDetails `json:"retrieval_details,omitempty"`

	// Suggested correction to the submitted **natural_language_query** value.
	SuggestedQuery *string `json:"suggested_query,omitempty"`

	// Array of suggested refinements.
	SuggestedRefinements []QuerySuggestedRefinement `json:"suggested_refinements,omitempty"`

	// Array of table results.
	TableResults []QueryTableResult `json:"table_results,omitempty"`

	// Passages returned by Discovery.
	Passages []QueryResponsePassage `json:"passages,omitempty"`
}


// UnmarshalQueryResponse unmarshals an instance of QueryResponse from the specified map of raw messages.
func UnmarshalQueryResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(QueryResponse)
	err = core.UnmarshalPrimitive(m, "matching_results", &obj.MatchingResults)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "results", &obj.Results, UnmarshalQueryResult)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "aggregations", &obj.Aggregations, UnmarshalQueryAggregation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "retrieval_details", &obj.RetrievalDetails, UnmarshalRetrievalDetails)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "suggested_query", &obj.SuggestedQuery)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "suggested_refinements", &obj.SuggestedRefinements, UnmarshalQuerySuggestedRefinement)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "table_results", &obj.TableResults, UnmarshalQueryTableResult)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "passages", &obj.Passages, UnmarshalQueryResponsePassage)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// QueryResponsePassage : A passage query response.
type QueryResponsePassage struct {
	// The content of the extracted passage.
	PassageText *string `json:"passage_text,omitempty"`

	// The confidence score of the passage's analysis. A higher score indicates greater confidence.
	PassageScore *float64 `json:"passage_score,omitempty"`

	// The unique identifier of the ingested document.
	DocumentID *string `json:"document_id,omitempty"`

	// The unique identifier of the collection.
	CollectionID *string `json:"collection_id,omitempty"`

	// The position of the first character of the extracted passage in the originating field.
	StartOffset *int64 `json:"start_offset,omitempty"`

	// The position of the last character of the extracted passage in the originating field.
	EndOffset *int64 `json:"end_offset,omitempty"`

	// The label of the field from which the passage has been extracted.
	Field *string `json:"field,omitempty"`
}


// UnmarshalQueryResponsePassage unmarshals an instance of QueryResponsePassage from the specified map of raw messages.
func UnmarshalQueryResponsePassage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(QueryResponsePassage)
	err = core.UnmarshalPrimitive(m, "passage_text", &obj.PassageText)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "passage_score", &obj.PassageScore)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "document_id", &obj.DocumentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "collection_id", &obj.CollectionID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "start_offset", &obj.StartOffset)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "end_offset", &obj.EndOffset)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "field", &obj.Field)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// QueryResult : Result document for the specified query.
type QueryResult struct {
	// The unique identifier of the document.
	DocumentID *string `json:"document_id" validate:"required"`

	// Metadata of the document.
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// Metadata of a query result.
	ResultMetadata *QueryResultMetadata `json:"result_metadata" validate:"required"`

	// Passages returned by Discovery.
	DocumentPassages []QueryResultPassage `json:"document_passages,omitempty"`

	// Allows users to set arbitrary properties
	additionalProperties map[string]interface{}
}


// SetProperty allows the user to set an arbitrary property on an instance of QueryResult
func (o *QueryResult) SetProperty(key string, value interface{}) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string]interface{})
	}
	o.additionalProperties[key] = value
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of QueryResult
func (o *QueryResult) GetProperty(key string) interface{} {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of QueryResult
func (o *QueryResult) GetProperties() map[string]interface{} {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of QueryResult
func (o *QueryResult) MarshalJSON() (buffer []byte, err error) {
	m := make(map[string]interface{})
	if len(o.additionalProperties) > 0 {
		for k, v := range o.additionalProperties {
			m[k] = v
		}
	}
	if o.DocumentID != nil {
		m["document_id"] = o.DocumentID
	}
	if o.Metadata != nil {
		m["metadata"] = o.Metadata
	}
	if o.ResultMetadata != nil {
		m["result_metadata"] = o.ResultMetadata
	}
	if o.DocumentPassages != nil {
		m["document_passages"] = o.DocumentPassages
	}
	buffer, err = json.Marshal(m)
	return
}

// UnmarshalQueryResult unmarshals an instance of QueryResult from the specified map of raw messages.
func UnmarshalQueryResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(QueryResult)
	err = core.UnmarshalPrimitive(m, "document_id", &obj.DocumentID)
	if err != nil {
		return
	}
	delete(m, "document_id")
	err = core.UnmarshalPrimitive(m, "metadata", &obj.Metadata)
	if err != nil {
		return
	}
	delete(m, "metadata")
	err = core.UnmarshalModel(m, "result_metadata", &obj.ResultMetadata, UnmarshalQueryResultMetadata)
	if err != nil {
		return
	}
	delete(m, "result_metadata")
	err = core.UnmarshalModel(m, "document_passages", &obj.DocumentPassages, UnmarshalQueryResultPassage)
	if err != nil {
		return
	}
	delete(m, "document_passages")
	for k := range m {
		var v interface{}
		e := core.UnmarshalPrimitive(m, k, &v)
		if e != nil {
			err = e
			return
		}
		obj.SetProperty(k, v)
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
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
	QueryResultMetadataDocumentRetrievalSourceCurationConst = "curation"
	QueryResultMetadataDocumentRetrievalSourceSearchConst = "search"
)


// UnmarshalQueryResultMetadata unmarshals an instance of QueryResultMetadata from the specified map of raw messages.
func UnmarshalQueryResultMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(QueryResultMetadata)
	err = core.UnmarshalPrimitive(m, "document_retrieval_source", &obj.DocumentRetrievalSource)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "collection_id", &obj.CollectionID)
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


// UnmarshalQueryResultPassage unmarshals an instance of QueryResultPassage from the specified map of raw messages.
func UnmarshalQueryResultPassage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(QueryResultPassage)
	err = core.UnmarshalPrimitive(m, "passage_text", &obj.PassageText)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "start_offset", &obj.StartOffset)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "end_offset", &obj.EndOffset)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "field", &obj.Field)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// QuerySuggestedRefinement : A suggested additional query term or terms user to filter results.
type QuerySuggestedRefinement struct {
	// The text used to filter.
	Text *string `json:"text,omitempty"`
}


// UnmarshalQuerySuggestedRefinement unmarshals an instance of QuerySuggestedRefinement from the specified map of raw messages.
func UnmarshalQuerySuggestedRefinement(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(QuerySuggestedRefinement)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
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


// UnmarshalQueryTableResult unmarshals an instance of QueryTableResult from the specified map of raw messages.
func UnmarshalQueryTableResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(QueryTableResult)
	err = core.UnmarshalPrimitive(m, "table_id", &obj.TableID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source_document_id", &obj.SourceDocumentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "collection_id", &obj.CollectionID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "table_html", &obj.TableHTML)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "table_html_offset", &obj.TableHTMLOffset)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "table", &obj.Table, UnmarshalTableResultTable)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// QueryTermAggregationResult : Top value result for the term aggregation.
type QueryTermAggregationResult struct {
	// Value of the field with a non-zero frequency in the document set.
	Key *string `json:"key" validate:"required"`

	// Number of documents containing the 'key'.
	MatchingResults *int64 `json:"matching_results" validate:"required"`

	// The relevancy for this term.
	Relevancy *float64 `json:"relevancy,omitempty"`

	// The number of documents which have the term as the value of specified field in the whole set of documents in this
	// collection. Returned only when the `relevancy` parameter is set to `true`.
	TotalMatchingDocuments *int64 `json:"total_matching_documents,omitempty"`

	// The estimated number of documents which would match the query and also meet the condition. Returned only when the
	// `relevancy` parameter is set to `true`.
	EstimatedMatchingDocuments *int64 `json:"estimated_matching_documents,omitempty"`

	// An array of sub aggregations.
	Aggregations []QueryAggregationIntf `json:"aggregations,omitempty"`
}


// UnmarshalQueryTermAggregationResult unmarshals an instance of QueryTermAggregationResult from the specified map of raw messages.
func UnmarshalQueryTermAggregationResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(QueryTermAggregationResult)
	err = core.UnmarshalPrimitive(m, "key", &obj.Key)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "matching_results", &obj.MatchingResults)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "relevancy", &obj.Relevancy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_matching_documents", &obj.TotalMatchingDocuments)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "estimated_matching_documents", &obj.EstimatedMatchingDocuments)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "aggregations", &obj.Aggregations, UnmarshalQueryAggregation)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
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
	Aggregations []QueryAggregationIntf `json:"aggregations,omitempty"`
}


// UnmarshalQueryTimesliceAggregationResult unmarshals an instance of QueryTimesliceAggregationResult from the specified map of raw messages.
func UnmarshalQueryTimesliceAggregationResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(QueryTimesliceAggregationResult)
	err = core.UnmarshalPrimitive(m, "key_as_string", &obj.KeyAsString)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "key", &obj.Key)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "matching_results", &obj.MatchingResults)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "aggregations", &obj.Aggregations, UnmarshalQueryAggregation)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// QueryTopHitsAggregationResult : A query response containing the matching documents for the preceding aggregations.
type QueryTopHitsAggregationResult struct {
	// Number of matching results.
	MatchingResults *int64 `json:"matching_results" validate:"required"`

	// An array of the document results.
	Hits []map[string]interface{} `json:"hits,omitempty"`
}


// UnmarshalQueryTopHitsAggregationResult unmarshals an instance of QueryTopHitsAggregationResult from the specified map of raw messages.
func UnmarshalQueryTopHitsAggregationResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(QueryTopHitsAggregationResult)
	err = core.UnmarshalPrimitive(m, "matching_results", &obj.MatchingResults)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hits", &obj.Hits)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
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
	RetrievalDetailsDocumentRetrievalStrategyRelevancyTrainingConst = "relevancy_training"
	RetrievalDetailsDocumentRetrievalStrategyUntrainedConst = "untrained"
)


// UnmarshalRetrievalDetails unmarshals an instance of RetrievalDetails from the specified map of raw messages.
func UnmarshalRetrievalDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RetrievalDetails)
	err = core.UnmarshalPrimitive(m, "document_retrieval_strategy", &obj.DocumentRetrievalStrategy)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

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


// UnmarshalTableBodyCells unmarshals an instance of TableBodyCells from the specified map of raw messages.
func UnmarshalTableBodyCells(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TableBodyCells)
	err = core.UnmarshalPrimitive(m, "cell_id", &obj.CellID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "location", &obj.Location, UnmarshalTableElementLocation)
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
	err = core.UnmarshalModel(m, "row_header_ids", &obj.RowHeaderIds, UnmarshalTableRowHeaderIds)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "row_header_texts", &obj.RowHeaderTexts, UnmarshalTableRowHeaderTexts)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "row_header_texts_normalized", &obj.RowHeaderTextsNormalized, UnmarshalTableRowHeaderTextsNormalized)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "column_header_ids", &obj.ColumnHeaderIds, UnmarshalTableColumnHeaderIds)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "column_header_texts", &obj.ColumnHeaderTexts, UnmarshalTableColumnHeaderTexts)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "column_header_texts_normalized", &obj.ColumnHeaderTextsNormalized, UnmarshalTableColumnHeaderTextsNormalized)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "attributes", &obj.Attributes, UnmarshalDocumentAttribute)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
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


// UnmarshalTableCellKey unmarshals an instance of TableCellKey from the specified map of raw messages.
func UnmarshalTableCellKey(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TableCellKey)
	err = core.UnmarshalPrimitive(m, "cell_id", &obj.CellID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "location", &obj.Location, UnmarshalTableElementLocation)
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


// UnmarshalTableCellValues unmarshals an instance of TableCellValues from the specified map of raw messages.
func UnmarshalTableCellValues(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TableCellValues)
	err = core.UnmarshalPrimitive(m, "cell_id", &obj.CellID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "location", &obj.Location, UnmarshalTableElementLocation)
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

// TableColumnHeaderIds : An array of values, each being the `id` value of a column header that is applicable to the current cell.
type TableColumnHeaderIds struct {
	// The `id` value of a column header.
	ID *string `json:"id,omitempty"`
}


// UnmarshalTableColumnHeaderIds unmarshals an instance of TableColumnHeaderIds from the specified map of raw messages.
func UnmarshalTableColumnHeaderIds(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TableColumnHeaderIds)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TableColumnHeaderTexts : An array of values, each being the `text` value of a column header that is applicable to the current cell.
type TableColumnHeaderTexts struct {
	// The `text` value of a column header.
	Text *string `json:"text,omitempty"`
}


// UnmarshalTableColumnHeaderTexts unmarshals an instance of TableColumnHeaderTexts from the specified map of raw messages.
func UnmarshalTableColumnHeaderTexts(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TableColumnHeaderTexts)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TableColumnHeaderTextsNormalized : If you provide customization input, the normalized version of the column header texts according to the customization;
// otherwise, the same value as `column_header_texts`.
type TableColumnHeaderTextsNormalized struct {
	// The normalized version of a column header text.
	TextNormalized *string `json:"text_normalized,omitempty"`
}


// UnmarshalTableColumnHeaderTextsNormalized unmarshals an instance of TableColumnHeaderTextsNormalized from the specified map of raw messages.
func UnmarshalTableColumnHeaderTextsNormalized(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TableColumnHeaderTextsNormalized)
	err = core.UnmarshalPrimitive(m, "text_normalized", &obj.TextNormalized)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
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


// UnmarshalTableColumnHeaders unmarshals an instance of TableColumnHeaders from the specified map of raw messages.
func UnmarshalTableColumnHeaders(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TableColumnHeaders)
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

// TableElementLocation : The numeric location of the identified element in the document, represented with two integers labeled `begin` and
// `end`.
type TableElementLocation struct {
	// The element's `begin` index.
	Begin *int64 `json:"begin" validate:"required"`

	// The element's `end` index.
	End *int64 `json:"end" validate:"required"`
}


// UnmarshalTableElementLocation unmarshals an instance of TableElementLocation from the specified map of raw messages.
func UnmarshalTableElementLocation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TableElementLocation)
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

// TableKeyValuePairs : Key-value pairs detected across cell boundaries.
type TableKeyValuePairs struct {
	// A key in a key-value pair.
	Key *TableCellKey `json:"key,omitempty"`

	// A list of values in a key-value pair.
	Value []TableCellValues `json:"value,omitempty"`
}


// UnmarshalTableKeyValuePairs unmarshals an instance of TableKeyValuePairs from the specified map of raw messages.
func UnmarshalTableKeyValuePairs(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TableKeyValuePairs)
	err = core.UnmarshalModel(m, "key", &obj.Key, UnmarshalTableCellKey)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "value", &obj.Value, UnmarshalTableCellValues)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
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


// UnmarshalTableResultTable unmarshals an instance of TableResultTable from the specified map of raw messages.
func UnmarshalTableResultTable(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TableResultTable)
	err = core.UnmarshalModel(m, "location", &obj.Location, UnmarshalTableElementLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "section_title", &obj.SectionTitle, UnmarshalTableTextLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "title", &obj.Title, UnmarshalTableTextLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "table_headers", &obj.TableHeaders, UnmarshalTableHeaders)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "row_headers", &obj.RowHeaders, UnmarshalTableRowHeaders)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "column_headers", &obj.ColumnHeaders, UnmarshalTableColumnHeaders)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "key_value_pairs", &obj.KeyValuePairs, UnmarshalTableKeyValuePairs)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "body_cells", &obj.BodyCells, UnmarshalTableBodyCells)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "contexts", &obj.Contexts, UnmarshalTableTextLocation)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TableRowHeaderIds : An array of values, each being the `id` value of a row header that is applicable to this body cell.
type TableRowHeaderIds struct {
	// The `id` values of a row header.
	ID *string `json:"id,omitempty"`
}


// UnmarshalTableRowHeaderIds unmarshals an instance of TableRowHeaderIds from the specified map of raw messages.
func UnmarshalTableRowHeaderIds(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TableRowHeaderIds)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TableRowHeaderTexts : An array of values, each being the `text` value of a row header that is applicable to this body cell.
type TableRowHeaderTexts struct {
	// The `text` value of a row header.
	Text *string `json:"text,omitempty"`
}


// UnmarshalTableRowHeaderTexts unmarshals an instance of TableRowHeaderTexts from the specified map of raw messages.
func UnmarshalTableRowHeaderTexts(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TableRowHeaderTexts)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TableRowHeaderTextsNormalized : If you provide customization input, the normalized version of the row header texts according to the customization;
// otherwise, the same value as `row_header_texts`.
type TableRowHeaderTextsNormalized struct {
	// The normalized version of a row header text.
	TextNormalized *string `json:"text_normalized,omitempty"`
}


// UnmarshalTableRowHeaderTextsNormalized unmarshals an instance of TableRowHeaderTextsNormalized from the specified map of raw messages.
func UnmarshalTableRowHeaderTextsNormalized(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TableRowHeaderTextsNormalized)
	err = core.UnmarshalPrimitive(m, "text_normalized", &obj.TextNormalized)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
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


// UnmarshalTableRowHeaders unmarshals an instance of TableRowHeaders from the specified map of raw messages.
func UnmarshalTableRowHeaders(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TableRowHeaders)
	err = core.UnmarshalPrimitive(m, "cell_id", &obj.CellID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "location", &obj.Location, UnmarshalTableElementLocation)
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

// TableTextLocation : Text and associated location within a table.
type TableTextLocation struct {
	// The text retrieved.
	Text *string `json:"text,omitempty"`

	// The numeric location of the identified element in the document, represented with two integers labeled `begin` and
	// `end`.
	Location *TableElementLocation `json:"location,omitempty"`
}


// UnmarshalTableTextLocation unmarshals an instance of TableTextLocation from the specified map of raw messages.
func UnmarshalTableTextLocation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TableTextLocation)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "location", &obj.Location, UnmarshalTableElementLocation)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
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
func (*DiscoveryV2) NewTrainingExample(documentID string, collectionID string, relevance int64) (model *TrainingExample, err error) {
	model = &TrainingExample{
		DocumentID: core.StringPtr(documentID),
		CollectionID: core.StringPtr(collectionID),
		Relevance: core.Int64Ptr(relevance),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalTrainingExample unmarshals an instance of TrainingExample from the specified map of raw messages.
func UnmarshalTrainingExample(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TrainingExample)
	err = core.UnmarshalPrimitive(m, "document_id", &obj.DocumentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "collection_id", &obj.CollectionID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "relevance", &obj.Relevance)
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
func (*DiscoveryV2) NewTrainingQuery(naturalLanguageQuery string, examples []TrainingExample) (model *TrainingQuery, err error) {
	model = &TrainingQuery{
		NaturalLanguageQuery: core.StringPtr(naturalLanguageQuery),
		Examples: examples,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalTrainingQuery unmarshals an instance of TrainingQuery from the specified map of raw messages.
func UnmarshalTrainingQuery(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TrainingQuery)
	err = core.UnmarshalPrimitive(m, "query_id", &obj.QueryID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "natural_language_query", &obj.NaturalLanguageQuery)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "filter", &obj.Filter)
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
	err = core.UnmarshalModel(m, "examples", &obj.Examples, UnmarshalTrainingExample)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TrainingQuerySet : Object specifying the training queries contained in the identified training set.
type TrainingQuerySet struct {
	// Array of training queries.
	Queries []TrainingQuery `json:"queries,omitempty"`
}


// UnmarshalTrainingQuerySet unmarshals an instance of TrainingQuerySet from the specified map of raw messages.
func UnmarshalTrainingQuerySet(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TrainingQuerySet)
	err = core.UnmarshalModel(m, "queries", &obj.Queries, UnmarshalTrainingQuery)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateCollectionOptions : The UpdateCollection options.
type UpdateCollectionOptions struct {
	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// The ID of the collection.
	CollectionID *string `json:"collection_id" validate:"required,ne="`

	// The name of the collection.
	Name *string `json:"name,omitempty"`

	// A description of the collection.
	Description *string `json:"description,omitempty"`

	// An array of enrichments that are applied to this collection.
	Enrichments []CollectionEnrichment `json:"enrichments,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateCollectionOptions : Instantiate UpdateCollectionOptions
func (*DiscoveryV2) NewUpdateCollectionOptions(projectID string, collectionID string) *UpdateCollectionOptions {
	return &UpdateCollectionOptions{
		ProjectID: core.StringPtr(projectID),
		CollectionID: core.StringPtr(collectionID),
	}
}

// SetProjectID : Allow user to set ProjectID
func (options *UpdateCollectionOptions) SetProjectID(projectID string) *UpdateCollectionOptions {
	options.ProjectID = core.StringPtr(projectID)
	return options
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

// SetEnrichments : Allow user to set Enrichments
func (options *UpdateCollectionOptions) SetEnrichments(enrichments []CollectionEnrichment) *UpdateCollectionOptions {
	options.Enrichments = enrichments
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateCollectionOptions) SetHeaders(param map[string]string) *UpdateCollectionOptions {
	options.Headers = param
	return options
}

// UpdateDocumentOptions : The UpdateDocument options.
type UpdateDocumentOptions struct {
	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// The ID of the collection.
	CollectionID *string `json:"collection_id" validate:"required,ne="`

	// The ID of the document.
	DocumentID *string `json:"document_id" validate:"required,ne="`

	// The content of the document to ingest. The maximum supported file size when adding a file to a collection is 50
	// megabytes, the maximum supported file size when testing a configuration is 1 megabyte. Files larger than the
	// supported size are rejected.
	File io.ReadCloser `json:"file,omitempty"`

	// The filename for file.
	Filename *string `json:"filename,omitempty"`

	// The content type of file.
	FileContentType *string `json:"file_content_type,omitempty"`

	// The maximum supported metadata file size is 1 MB. Metadata parts larger than 1 MB are rejected.
	//
	//
	// Example:  ``` {
	//   "Creator": "Johnny Appleseed",
	//   "Subject": "Apples"
	// } ```.
	Metadata *string `json:"metadata,omitempty"`

	// When `true`, the uploaded document is added to the collection even if the data for that collection is shared with
	// other collections.
	XWatsonDiscoveryForce *bool `json:"X-Watson-Discovery-Force,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateDocumentOptions : Instantiate UpdateDocumentOptions
func (*DiscoveryV2) NewUpdateDocumentOptions(projectID string, collectionID string, documentID string) *UpdateDocumentOptions {
	return &UpdateDocumentOptions{
		ProjectID: core.StringPtr(projectID),
		CollectionID: core.StringPtr(collectionID),
		DocumentID: core.StringPtr(documentID),
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

// UpdateEnrichmentOptions : The UpdateEnrichment options.
type UpdateEnrichmentOptions struct {
	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// The ID of the enrichment.
	EnrichmentID *string `json:"enrichment_id" validate:"required,ne="`

	// A new name for the enrichment.
	Name *string `json:"name" validate:"required"`

	// A new description for the enrichment.
	Description *string `json:"description,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateEnrichmentOptions : Instantiate UpdateEnrichmentOptions
func (*DiscoveryV2) NewUpdateEnrichmentOptions(projectID string, enrichmentID string, name string) *UpdateEnrichmentOptions {
	return &UpdateEnrichmentOptions{
		ProjectID: core.StringPtr(projectID),
		EnrichmentID: core.StringPtr(enrichmentID),
		Name: core.StringPtr(name),
	}
}

// SetProjectID : Allow user to set ProjectID
func (options *UpdateEnrichmentOptions) SetProjectID(projectID string) *UpdateEnrichmentOptions {
	options.ProjectID = core.StringPtr(projectID)
	return options
}

// SetEnrichmentID : Allow user to set EnrichmentID
func (options *UpdateEnrichmentOptions) SetEnrichmentID(enrichmentID string) *UpdateEnrichmentOptions {
	options.EnrichmentID = core.StringPtr(enrichmentID)
	return options
}

// SetName : Allow user to set Name
func (options *UpdateEnrichmentOptions) SetName(name string) *UpdateEnrichmentOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetDescription : Allow user to set Description
func (options *UpdateEnrichmentOptions) SetDescription(description string) *UpdateEnrichmentOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateEnrichmentOptions) SetHeaders(param map[string]string) *UpdateEnrichmentOptions {
	options.Headers = param
	return options
}

// UpdateProjectOptions : The UpdateProject options.
type UpdateProjectOptions struct {
	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// The new name to give this project.
	Name *string `json:"name,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateProjectOptions : Instantiate UpdateProjectOptions
func (*DiscoveryV2) NewUpdateProjectOptions(projectID string) *UpdateProjectOptions {
	return &UpdateProjectOptions{
		ProjectID: core.StringPtr(projectID),
	}
}

// SetProjectID : Allow user to set ProjectID
func (options *UpdateProjectOptions) SetProjectID(projectID string) *UpdateProjectOptions {
	options.ProjectID = core.StringPtr(projectID)
	return options
}

// SetName : Allow user to set Name
func (options *UpdateProjectOptions) SetName(name string) *UpdateProjectOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateProjectOptions) SetHeaders(param map[string]string) *UpdateProjectOptions {
	options.Headers = param
	return options
}

// UpdateTrainingQueryOptions : The UpdateTrainingQuery options.
type UpdateTrainingQueryOptions struct {
	// The ID of the project. This information can be found from the deploy page of the Discovery administrative tooling.
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// The ID of the query used for training.
	QueryID *string `json:"query_id" validate:"required,ne="`

	// The natural text query for the training query.
	NaturalLanguageQuery *string `json:"natural_language_query" validate:"required"`

	// Array of training examples.
	Examples []TrainingExample `json:"examples" validate:"required"`

	// The filter used on the collection before the **natural_language_query** is applied.
	Filter *string `json:"filter,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateTrainingQueryOptions : Instantiate UpdateTrainingQueryOptions
func (*DiscoveryV2) NewUpdateTrainingQueryOptions(projectID string, queryID string, naturalLanguageQuery string, examples []TrainingExample) *UpdateTrainingQueryOptions {
	return &UpdateTrainingQueryOptions{
		ProjectID: core.StringPtr(projectID),
		QueryID: core.StringPtr(queryID),
		NaturalLanguageQuery: core.StringPtr(naturalLanguageQuery),
		Examples: examples,
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
// This model "extends" QueryAggregation
type QueryCalculationAggregation struct {
	// The type of aggregation command used. Options include: term, histogram, timeslice, nested, filter, min, max, sum,
	// average, unique_count, and top_hits.
	Type *string `json:"type" validate:"required"`

	// The field to perform the calculation on.
	Field *string `json:"field" validate:"required"`

	// The value of the calculation.
	Value *float64 `json:"value,omitempty"`
}


func (*QueryCalculationAggregation) isaQueryAggregation() bool {
	return true
}

// UnmarshalQueryCalculationAggregation unmarshals an instance of QueryCalculationAggregation from the specified map of raw messages.
func UnmarshalQueryCalculationAggregation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(QueryCalculationAggregation)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "field", &obj.Field)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// QueryFilterAggregation : A modifier that will narrow down the document set of the sub aggregations it precedes.
// This model "extends" QueryAggregation
type QueryFilterAggregation struct {
	// The type of aggregation command used. Options include: term, histogram, timeslice, nested, filter, min, max, sum,
	// average, unique_count, and top_hits.
	Type *string `json:"type" validate:"required"`

	// The filter written in Discovery Query Language syntax applied to the documents before sub aggregations are run.
	Match *string `json:"match" validate:"required"`

	// Number of documents matching the filter.
	MatchingResults *int64 `json:"matching_results" validate:"required"`

	// An array of sub aggregations.
	Aggregations []QueryAggregationIntf `json:"aggregations,omitempty"`
}


func (*QueryFilterAggregation) isaQueryAggregation() bool {
	return true
}

// UnmarshalQueryFilterAggregation unmarshals an instance of QueryFilterAggregation from the specified map of raw messages.
func UnmarshalQueryFilterAggregation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(QueryFilterAggregation)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "match", &obj.Match)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "matching_results", &obj.MatchingResults)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "aggregations", &obj.Aggregations, UnmarshalQueryAggregation)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// QueryGroupByAggregation : Returns the top values for the field specified.
// This model "extends" QueryAggregation
type QueryGroupByAggregation struct {
	// The type of aggregation command used. Options include: term, histogram, timeslice, nested, filter, min, max, sum,
	// average, unique_count, and top_hits.
	Type *string `json:"type" validate:"required"`

	// Array of top values for the field.
	Results []QueryGroupByAggregationResult `json:"results,omitempty"`
}


func (*QueryGroupByAggregation) isaQueryAggregation() bool {
	return true
}

// UnmarshalQueryGroupByAggregation unmarshals an instance of QueryGroupByAggregation from the specified map of raw messages.
func UnmarshalQueryGroupByAggregation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(QueryGroupByAggregation)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "results", &obj.Results, UnmarshalQueryGroupByAggregationResult)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// QueryHistogramAggregation : Numeric interval segments to categorize documents by using field values from a single numeric field to describe the
// category.
// This model "extends" QueryAggregation
type QueryHistogramAggregation struct {
	// The type of aggregation command used. Options include: term, histogram, timeslice, nested, filter, min, max, sum,
	// average, unique_count, and top_hits.
	Type *string `json:"type" validate:"required"`

	// The numeric field name used to create the histogram.
	Field *string `json:"field" validate:"required"`

	// The size of the sections the results are split into.
	Interval *int64 `json:"interval" validate:"required"`

	// Identifier specified in the query request of this aggregation.
	Name *string `json:"name,omitempty"`

	// Array of numeric intervals.
	Results []QueryHistogramAggregationResult `json:"results,omitempty"`
}


func (*QueryHistogramAggregation) isaQueryAggregation() bool {
	return true
}

// UnmarshalQueryHistogramAggregation unmarshals an instance of QueryHistogramAggregation from the specified map of raw messages.
func UnmarshalQueryHistogramAggregation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(QueryHistogramAggregation)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "field", &obj.Field)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "interval", &obj.Interval)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "results", &obj.Results, UnmarshalQueryHistogramAggregationResult)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// QueryNestedAggregation : A restriction that alter the document set used for sub aggregations it precedes to nested documents found in the
// field specified.
// This model "extends" QueryAggregation
type QueryNestedAggregation struct {
	// The type of aggregation command used. Options include: term, histogram, timeslice, nested, filter, min, max, sum,
	// average, unique_count, and top_hits.
	Type *string `json:"type" validate:"required"`

	// The path to the document field to scope sub aggregations to.
	Path *string `json:"path" validate:"required"`

	// Number of nested documents found in the specified field.
	MatchingResults *int64 `json:"matching_results" validate:"required"`

	// An array of sub aggregations.
	Aggregations []QueryAggregationIntf `json:"aggregations,omitempty"`
}


func (*QueryNestedAggregation) isaQueryAggregation() bool {
	return true
}

// UnmarshalQueryNestedAggregation unmarshals an instance of QueryNestedAggregation from the specified map of raw messages.
func UnmarshalQueryNestedAggregation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(QueryNestedAggregation)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "path", &obj.Path)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "matching_results", &obj.MatchingResults)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "aggregations", &obj.Aggregations, UnmarshalQueryAggregation)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// QueryTermAggregation : Returns the top values for the field specified.
// This model "extends" QueryAggregation
type QueryTermAggregation struct {
	// The type of aggregation command used. Options include: term, histogram, timeslice, nested, filter, min, max, sum,
	// average, unique_count, and top_hits.
	Type *string `json:"type" validate:"required"`

	// The field in the document used to generate top values from.
	Field *string `json:"field" validate:"required"`

	// The number of top values returned.
	Count *int64 `json:"count,omitempty"`

	// Identifier specified in the query request of this aggregation.
	Name *string `json:"name,omitempty"`

	// Array of top values for the field.
	Results []QueryTermAggregationResult `json:"results,omitempty"`
}


func (*QueryTermAggregation) isaQueryAggregation() bool {
	return true
}

// UnmarshalQueryTermAggregation unmarshals an instance of QueryTermAggregation from the specified map of raw messages.
func UnmarshalQueryTermAggregation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(QueryTermAggregation)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "field", &obj.Field)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "count", &obj.Count)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "results", &obj.Results, UnmarshalQueryTermAggregationResult)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// QueryTimesliceAggregation : A specialized histogram aggregation that uses dates to create interval segments.
// This model "extends" QueryAggregation
type QueryTimesliceAggregation struct {
	// The type of aggregation command used. Options include: term, histogram, timeslice, nested, filter, min, max, sum,
	// average, unique_count, and top_hits.
	Type *string `json:"type" validate:"required"`

	// The date field name used to create the timeslice.
	Field *string `json:"field" validate:"required"`

	// The date interval value. Valid values are seconds, minutes, hours, days, weeks, and years.
	Interval *string `json:"interval" validate:"required"`

	// Identifier specified in the query request of this aggregation.
	Name *string `json:"name,omitempty"`

	// Array of aggregation results.
	Results []QueryTimesliceAggregationResult `json:"results,omitempty"`
}


func (*QueryTimesliceAggregation) isaQueryAggregation() bool {
	return true
}

// UnmarshalQueryTimesliceAggregation unmarshals an instance of QueryTimesliceAggregation from the specified map of raw messages.
func UnmarshalQueryTimesliceAggregation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(QueryTimesliceAggregation)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "field", &obj.Field)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "interval", &obj.Interval)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "results", &obj.Results, UnmarshalQueryTimesliceAggregationResult)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// QueryTopHitsAggregation : Returns the top documents ranked by the score of the query.
// This model "extends" QueryAggregation
type QueryTopHitsAggregation struct {
	// The type of aggregation command used. Options include: term, histogram, timeslice, nested, filter, min, max, sum,
	// average, unique_count, and top_hits.
	Type *string `json:"type" validate:"required"`

	// The number of documents to return.
	Size *int64 `json:"size" validate:"required"`

	// Identifier specified in the query request of this aggregation.
	Name *string `json:"name,omitempty"`

	Hits *QueryTopHitsAggregationResult `json:"hits,omitempty"`
}


func (*QueryTopHitsAggregation) isaQueryAggregation() bool {
	return true
}

// UnmarshalQueryTopHitsAggregation unmarshals an instance of QueryTopHitsAggregation from the specified map of raw messages.
func UnmarshalQueryTopHitsAggregation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(QueryTopHitsAggregation)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "size", &obj.Size)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "hits", &obj.Hits, UnmarshalQueryTopHitsAggregationResult)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
