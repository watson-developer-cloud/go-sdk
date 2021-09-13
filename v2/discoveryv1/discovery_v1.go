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
 * IBM OpenAPI SDK Code Generator Version: 3.38.0-07189efd-20210827-205025
 */

// Package discoveryv1 : Operations and models for the DiscoveryV1 service
package discoveryv1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	common "github.com/watson-developer-cloud/go-sdk/v2/common"
)

// DiscoveryV1 : IBM Watson&trade; Discovery is a cognitive search and content analytics engine that you can add to
// applications to identify patterns, trends and actionable insights to drive better decision-making. Securely unify
// structured and unstructured data with pre-enriched content, and use a simplified query language to eliminate the need
// for manual filtering of results.
//
// API Version: 1.0
// See: https://cloud.ibm.com/docs/discovery
type DiscoveryV1 struct {
	Service *core.BaseService

	// Release date of the version of the API you want to use. Specify dates in YYYY-MM-DD format. The current version is
	// `2019-04-30`.
	Version *string
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.us-south.discovery.watson.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "discovery"

// DiscoveryV1Options : Service options
type DiscoveryV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator

	// Release date of the version of the API you want to use. Specify dates in YYYY-MM-DD format. The current version is
	// `2019-04-30`.
	Version *string `validate:"required"`
}

// NewDiscoveryV1 : constructs an instance of DiscoveryV1 with passed in options.
func NewDiscoveryV1(options *DiscoveryV1Options) (service *DiscoveryV1, err error) {
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

	service = &DiscoveryV1{
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
func (discovery *DiscoveryV1) Clone() *DiscoveryV1 {
	if core.IsNil(discovery) {
		return nil
	}
	clone := *discovery
	clone.Service = discovery.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (discovery *DiscoveryV1) SetServiceURL(url string) error {
	return discovery.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (discovery *DiscoveryV1) GetServiceURL() string {
	return discovery.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (discovery *DiscoveryV1) SetDefaultHeaders(headers http.Header) {
	discovery.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (discovery *DiscoveryV1) SetEnableGzipCompression(enableGzip bool) {
	discovery.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (discovery *DiscoveryV1) GetEnableGzipCompression() bool {
	return discovery.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (discovery *DiscoveryV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	discovery.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (discovery *DiscoveryV1) DisableRetries() {
	discovery.Service.DisableRetries()
}

// DisableSSLVerification bypasses verification of the server's SSL certificate
func (discovery *DiscoveryV1) DisableSSLVerification() {
	discovery.Service.DisableSSLVerification()
}

// CreateEnvironment : Create an environment
// Creates a new environment for private data. An environment must be created before collections can be created.
//
// **Note**: You can create only one environment for private data per service instance. An attempt to create another
// environment results in an error.
func (discovery *DiscoveryV1) CreateEnvironment(createEnvironmentOptions *CreateEnvironmentOptions) (result *Environment, response *core.DetailedResponse, err error) {
	return discovery.CreateEnvironmentWithContext(context.Background(), createEnvironmentOptions)
}

// CreateEnvironmentWithContext is an alternate form of the CreateEnvironment method which supports a Context parameter
func (discovery *DiscoveryV1) CreateEnvironmentWithContext(ctx context.Context, createEnvironmentOptions *CreateEnvironmentOptions) (result *Environment, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createEnvironmentOptions, "createEnvironmentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createEnvironmentOptions, "createEnvironmentOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createEnvironmentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "CreateEnvironment")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

	body := make(map[string]interface{})
	if createEnvironmentOptions.Name != nil {
		body["name"] = createEnvironmentOptions.Name
	}
	if createEnvironmentOptions.Description != nil {
		body["description"] = createEnvironmentOptions.Description
	}
	if createEnvironmentOptions.Size != nil {
		body["size"] = createEnvironmentOptions.Size
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalEnvironment)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListEnvironments : List environments
// List existing environments for the service instance.
func (discovery *DiscoveryV1) ListEnvironments(listEnvironmentsOptions *ListEnvironmentsOptions) (result *ListEnvironmentsResponse, response *core.DetailedResponse, err error) {
	return discovery.ListEnvironmentsWithContext(context.Background(), listEnvironmentsOptions)
}

// ListEnvironmentsWithContext is an alternate form of the ListEnvironments method which supports a Context parameter
func (discovery *DiscoveryV1) ListEnvironmentsWithContext(ctx context.Context, listEnvironmentsOptions *ListEnvironmentsOptions) (result *ListEnvironmentsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listEnvironmentsOptions, "listEnvironmentsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listEnvironmentsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "ListEnvironments")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))
	if listEnvironmentsOptions.Name != nil {
		builder.AddQuery("name", fmt.Sprint(*listEnvironmentsOptions.Name))
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListEnvironmentsResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetEnvironment : Get environment info
func (discovery *DiscoveryV1) GetEnvironment(getEnvironmentOptions *GetEnvironmentOptions) (result *Environment, response *core.DetailedResponse, err error) {
	return discovery.GetEnvironmentWithContext(context.Background(), getEnvironmentOptions)
}

// GetEnvironmentWithContext is an alternate form of the GetEnvironment method which supports a Context parameter
func (discovery *DiscoveryV1) GetEnvironmentWithContext(ctx context.Context, getEnvironmentOptions *GetEnvironmentOptions) (result *Environment, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getEnvironmentOptions, "getEnvironmentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getEnvironmentOptions, "getEnvironmentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *getEnvironmentOptions.EnvironmentID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getEnvironmentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "GetEnvironment")
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalEnvironment)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateEnvironment : Update an environment
// Updates an environment. The environment's **name** and  **description** parameters can be changed. You must specify a
// **name** for the environment.
func (discovery *DiscoveryV1) UpdateEnvironment(updateEnvironmentOptions *UpdateEnvironmentOptions) (result *Environment, response *core.DetailedResponse, err error) {
	return discovery.UpdateEnvironmentWithContext(context.Background(), updateEnvironmentOptions)
}

// UpdateEnvironmentWithContext is an alternate form of the UpdateEnvironment method which supports a Context parameter
func (discovery *DiscoveryV1) UpdateEnvironmentWithContext(ctx context.Context, updateEnvironmentOptions *UpdateEnvironmentOptions) (result *Environment, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateEnvironmentOptions, "updateEnvironmentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateEnvironmentOptions, "updateEnvironmentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *updateEnvironmentOptions.EnvironmentID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateEnvironmentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "UpdateEnvironment")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

	body := make(map[string]interface{})
	if updateEnvironmentOptions.Name != nil {
		body["name"] = updateEnvironmentOptions.Name
	}
	if updateEnvironmentOptions.Description != nil {
		body["description"] = updateEnvironmentOptions.Description
	}
	if updateEnvironmentOptions.Size != nil {
		body["size"] = updateEnvironmentOptions.Size
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalEnvironment)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteEnvironment : Delete environment
func (discovery *DiscoveryV1) DeleteEnvironment(deleteEnvironmentOptions *DeleteEnvironmentOptions) (result *DeleteEnvironmentResponse, response *core.DetailedResponse, err error) {
	return discovery.DeleteEnvironmentWithContext(context.Background(), deleteEnvironmentOptions)
}

// DeleteEnvironmentWithContext is an alternate form of the DeleteEnvironment method which supports a Context parameter
func (discovery *DiscoveryV1) DeleteEnvironmentWithContext(ctx context.Context, deleteEnvironmentOptions *DeleteEnvironmentOptions) (result *DeleteEnvironmentResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteEnvironmentOptions, "deleteEnvironmentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteEnvironmentOptions, "deleteEnvironmentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *deleteEnvironmentOptions.EnvironmentID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteEnvironmentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "DeleteEnvironment")
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeleteEnvironmentResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListFields : List fields across collections
// Gets a list of the unique fields (and their types) stored in the indexes of the specified collections.
func (discovery *DiscoveryV1) ListFields(listFieldsOptions *ListFieldsOptions) (result *ListCollectionFieldsResponse, response *core.DetailedResponse, err error) {
	return discovery.ListFieldsWithContext(context.Background(), listFieldsOptions)
}

// ListFieldsWithContext is an alternate form of the ListFields method which supports a Context parameter
func (discovery *DiscoveryV1) ListFieldsWithContext(ctx context.Context, listFieldsOptions *ListFieldsOptions) (result *ListCollectionFieldsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listFieldsOptions, "listFieldsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listFieldsOptions, "listFieldsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *listFieldsOptions.EnvironmentID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/fields`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listFieldsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "ListFields")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))
	builder.AddQuery("collection_ids", strings.Join(listFieldsOptions.CollectionIds, ","))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = discovery.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListCollectionFieldsResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateConfiguration : Add configuration
// Creates a new configuration.
//
// If the input configuration contains the **configuration_id**, **created**, or **updated** properties, then they are
// ignored and overridden by the system, and an error is not returned so that the overridden fields do not need to be
// removed when copying a configuration.
//
// The configuration can contain unrecognized JSON fields. Any such fields are ignored and do not generate an error.
// This makes it easier to use newer configuration files with older versions of the API and the service. It also makes
// it possible for the tooling to add additional metadata and information to the configuration.
func (discovery *DiscoveryV1) CreateConfiguration(createConfigurationOptions *CreateConfigurationOptions) (result *Configuration, response *core.DetailedResponse, err error) {
	return discovery.CreateConfigurationWithContext(context.Background(), createConfigurationOptions)
}

// CreateConfigurationWithContext is an alternate form of the CreateConfiguration method which supports a Context parameter
func (discovery *DiscoveryV1) CreateConfigurationWithContext(ctx context.Context, createConfigurationOptions *CreateConfigurationOptions) (result *Configuration, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createConfigurationOptions, "createConfigurationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createConfigurationOptions, "createConfigurationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *createConfigurationOptions.EnvironmentID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/configurations`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createConfigurationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "CreateConfiguration")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

	body := make(map[string]interface{})
	if createConfigurationOptions.Name != nil {
		body["name"] = createConfigurationOptions.Name
	}
	if createConfigurationOptions.Description != nil {
		body["description"] = createConfigurationOptions.Description
	}
	if createConfigurationOptions.Conversions != nil {
		body["conversions"] = createConfigurationOptions.Conversions
	}
	if createConfigurationOptions.Enrichments != nil {
		body["enrichments"] = createConfigurationOptions.Enrichments
	}
	if createConfigurationOptions.Normalizations != nil {
		body["normalizations"] = createConfigurationOptions.Normalizations
	}
	if createConfigurationOptions.Source != nil {
		body["source"] = createConfigurationOptions.Source
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalConfiguration)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListConfigurations : List configurations
// Lists existing configurations for the service instance.
func (discovery *DiscoveryV1) ListConfigurations(listConfigurationsOptions *ListConfigurationsOptions) (result *ListConfigurationsResponse, response *core.DetailedResponse, err error) {
	return discovery.ListConfigurationsWithContext(context.Background(), listConfigurationsOptions)
}

// ListConfigurationsWithContext is an alternate form of the ListConfigurations method which supports a Context parameter
func (discovery *DiscoveryV1) ListConfigurationsWithContext(ctx context.Context, listConfigurationsOptions *ListConfigurationsOptions) (result *ListConfigurationsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listConfigurationsOptions, "listConfigurationsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listConfigurationsOptions, "listConfigurationsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *listConfigurationsOptions.EnvironmentID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/configurations`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listConfigurationsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "ListConfigurations")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))
	if listConfigurationsOptions.Name != nil {
		builder.AddQuery("name", fmt.Sprint(*listConfigurationsOptions.Name))
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListConfigurationsResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetConfiguration : Get configuration details
func (discovery *DiscoveryV1) GetConfiguration(getConfigurationOptions *GetConfigurationOptions) (result *Configuration, response *core.DetailedResponse, err error) {
	return discovery.GetConfigurationWithContext(context.Background(), getConfigurationOptions)
}

// GetConfigurationWithContext is an alternate form of the GetConfiguration method which supports a Context parameter
func (discovery *DiscoveryV1) GetConfigurationWithContext(ctx context.Context, getConfigurationOptions *GetConfigurationOptions) (result *Configuration, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getConfigurationOptions, "getConfigurationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getConfigurationOptions, "getConfigurationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id":   *getConfigurationOptions.EnvironmentID,
		"configuration_id": *getConfigurationOptions.ConfigurationID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/configurations/{configuration_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getConfigurationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "GetConfiguration")
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalConfiguration)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateConfiguration : Update a configuration
// Replaces an existing configuration.
//   * Completely replaces the original configuration.
//   * The **configuration_id**, **updated**, and **created** fields are accepted in the request, but they are ignored,
// and an error is not generated. It is also acceptable for users to submit an updated configuration with none of the
// three properties.
//   * Documents are processed with a snapshot of the configuration as it was at the time the document was submitted to
// be ingested. This means that already submitted documents will not see any updates made to the configuration.
func (discovery *DiscoveryV1) UpdateConfiguration(updateConfigurationOptions *UpdateConfigurationOptions) (result *Configuration, response *core.DetailedResponse, err error) {
	return discovery.UpdateConfigurationWithContext(context.Background(), updateConfigurationOptions)
}

// UpdateConfigurationWithContext is an alternate form of the UpdateConfiguration method which supports a Context parameter
func (discovery *DiscoveryV1) UpdateConfigurationWithContext(ctx context.Context, updateConfigurationOptions *UpdateConfigurationOptions) (result *Configuration, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateConfigurationOptions, "updateConfigurationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateConfigurationOptions, "updateConfigurationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id":   *updateConfigurationOptions.EnvironmentID,
		"configuration_id": *updateConfigurationOptions.ConfigurationID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/configurations/{configuration_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateConfigurationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "UpdateConfiguration")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

	body := make(map[string]interface{})
	if updateConfigurationOptions.Name != nil {
		body["name"] = updateConfigurationOptions.Name
	}
	if updateConfigurationOptions.Description != nil {
		body["description"] = updateConfigurationOptions.Description
	}
	if updateConfigurationOptions.Conversions != nil {
		body["conversions"] = updateConfigurationOptions.Conversions
	}
	if updateConfigurationOptions.Enrichments != nil {
		body["enrichments"] = updateConfigurationOptions.Enrichments
	}
	if updateConfigurationOptions.Normalizations != nil {
		body["normalizations"] = updateConfigurationOptions.Normalizations
	}
	if updateConfigurationOptions.Source != nil {
		body["source"] = updateConfigurationOptions.Source
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalConfiguration)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteConfiguration : Delete a configuration
// The deletion is performed unconditionally. A configuration deletion request succeeds even if the configuration is
// referenced by a collection or document ingestion. However, documents that have already been submitted for processing
// continue to use the deleted configuration. Documents are always processed with a snapshot of the configuration as it
// existed at the time the document was submitted.
func (discovery *DiscoveryV1) DeleteConfiguration(deleteConfigurationOptions *DeleteConfigurationOptions) (result *DeleteConfigurationResponse, response *core.DetailedResponse, err error) {
	return discovery.DeleteConfigurationWithContext(context.Background(), deleteConfigurationOptions)
}

// DeleteConfigurationWithContext is an alternate form of the DeleteConfiguration method which supports a Context parameter
func (discovery *DiscoveryV1) DeleteConfigurationWithContext(ctx context.Context, deleteConfigurationOptions *DeleteConfigurationOptions) (result *DeleteConfigurationResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteConfigurationOptions, "deleteConfigurationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteConfigurationOptions, "deleteConfigurationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id":   *deleteConfigurationOptions.EnvironmentID,
		"configuration_id": *deleteConfigurationOptions.ConfigurationID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/configurations/{configuration_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteConfigurationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "DeleteConfiguration")
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeleteConfigurationResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateCollection : Create a collection
func (discovery *DiscoveryV1) CreateCollection(createCollectionOptions *CreateCollectionOptions) (result *Collection, response *core.DetailedResponse, err error) {
	return discovery.CreateCollectionWithContext(context.Background(), createCollectionOptions)
}

// CreateCollectionWithContext is an alternate form of the CreateCollection method which supports a Context parameter
func (discovery *DiscoveryV1) CreateCollectionWithContext(ctx context.Context, createCollectionOptions *CreateCollectionOptions) (result *Collection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createCollectionOptions, "createCollectionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createCollectionOptions, "createCollectionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *createCollectionOptions.EnvironmentID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/collections`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createCollectionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "CreateCollection")
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
	if createCollectionOptions.ConfigurationID != nil {
		body["configuration_id"] = createCollectionOptions.ConfigurationID
	}
	if createCollectionOptions.Language != nil {
		body["language"] = createCollectionOptions.Language
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
// Lists existing collections for the service instance.
func (discovery *DiscoveryV1) ListCollections(listCollectionsOptions *ListCollectionsOptions) (result *ListCollectionsResponse, response *core.DetailedResponse, err error) {
	return discovery.ListCollectionsWithContext(context.Background(), listCollectionsOptions)
}

// ListCollectionsWithContext is an alternate form of the ListCollections method which supports a Context parameter
func (discovery *DiscoveryV1) ListCollectionsWithContext(ctx context.Context, listCollectionsOptions *ListCollectionsOptions) (result *ListCollectionsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listCollectionsOptions, "listCollectionsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listCollectionsOptions, "listCollectionsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *listCollectionsOptions.EnvironmentID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/collections`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listCollectionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "ListCollections")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))
	if listCollectionsOptions.Name != nil {
		builder.AddQuery("name", fmt.Sprint(*listCollectionsOptions.Name))
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListCollectionsResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetCollection : Get collection details
func (discovery *DiscoveryV1) GetCollection(getCollectionOptions *GetCollectionOptions) (result *Collection, response *core.DetailedResponse, err error) {
	return discovery.GetCollectionWithContext(context.Background(), getCollectionOptions)
}

// GetCollectionWithContext is an alternate form of the GetCollection method which supports a Context parameter
func (discovery *DiscoveryV1) GetCollectionWithContext(ctx context.Context, getCollectionOptions *GetCollectionOptions) (result *Collection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getCollectionOptions, "getCollectionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getCollectionOptions, "getCollectionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *getCollectionOptions.EnvironmentID,
		"collection_id":  *getCollectionOptions.CollectionID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/collections/{collection_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCollectionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "GetCollection")
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
func (discovery *DiscoveryV1) UpdateCollection(updateCollectionOptions *UpdateCollectionOptions) (result *Collection, response *core.DetailedResponse, err error) {
	return discovery.UpdateCollectionWithContext(context.Background(), updateCollectionOptions)
}

// UpdateCollectionWithContext is an alternate form of the UpdateCollection method which supports a Context parameter
func (discovery *DiscoveryV1) UpdateCollectionWithContext(ctx context.Context, updateCollectionOptions *UpdateCollectionOptions) (result *Collection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateCollectionOptions, "updateCollectionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateCollectionOptions, "updateCollectionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *updateCollectionOptions.EnvironmentID,
		"collection_id":  *updateCollectionOptions.CollectionID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/collections/{collection_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateCollectionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "UpdateCollection")
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
	if updateCollectionOptions.ConfigurationID != nil {
		body["configuration_id"] = updateCollectionOptions.ConfigurationID
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
func (discovery *DiscoveryV1) DeleteCollection(deleteCollectionOptions *DeleteCollectionOptions) (result *DeleteCollectionResponse, response *core.DetailedResponse, err error) {
	return discovery.DeleteCollectionWithContext(context.Background(), deleteCollectionOptions)
}

// DeleteCollectionWithContext is an alternate form of the DeleteCollection method which supports a Context parameter
func (discovery *DiscoveryV1) DeleteCollectionWithContext(ctx context.Context, deleteCollectionOptions *DeleteCollectionOptions) (result *DeleteCollectionResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteCollectionOptions, "deleteCollectionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteCollectionOptions, "deleteCollectionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *deleteCollectionOptions.EnvironmentID,
		"collection_id":  *deleteCollectionOptions.CollectionID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/collections/{collection_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteCollectionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "DeleteCollection")
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeleteCollectionResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListCollectionFields : List collection fields
// Gets a list of the unique fields (and their types) stored in the index.
func (discovery *DiscoveryV1) ListCollectionFields(listCollectionFieldsOptions *ListCollectionFieldsOptions) (result *ListCollectionFieldsResponse, response *core.DetailedResponse, err error) {
	return discovery.ListCollectionFieldsWithContext(context.Background(), listCollectionFieldsOptions)
}

// ListCollectionFieldsWithContext is an alternate form of the ListCollectionFields method which supports a Context parameter
func (discovery *DiscoveryV1) ListCollectionFieldsWithContext(ctx context.Context, listCollectionFieldsOptions *ListCollectionFieldsOptions) (result *ListCollectionFieldsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listCollectionFieldsOptions, "listCollectionFieldsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listCollectionFieldsOptions, "listCollectionFieldsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *listCollectionFieldsOptions.EnvironmentID,
		"collection_id":  *listCollectionFieldsOptions.CollectionID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/collections/{collection_id}/fields`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listCollectionFieldsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "ListCollectionFields")
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListCollectionFieldsResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListExpansions : Get the expansion list
// Returns the current expansion list for the specified collection. If an expansion list is not specified, an object
// with empty expansion arrays is returned.
func (discovery *DiscoveryV1) ListExpansions(listExpansionsOptions *ListExpansionsOptions) (result *Expansions, response *core.DetailedResponse, err error) {
	return discovery.ListExpansionsWithContext(context.Background(), listExpansionsOptions)
}

// ListExpansionsWithContext is an alternate form of the ListExpansions method which supports a Context parameter
func (discovery *DiscoveryV1) ListExpansionsWithContext(ctx context.Context, listExpansionsOptions *ListExpansionsOptions) (result *Expansions, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listExpansionsOptions, "listExpansionsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listExpansionsOptions, "listExpansionsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *listExpansionsOptions.EnvironmentID,
		"collection_id":  *listExpansionsOptions.CollectionID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/collections/{collection_id}/expansions`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listExpansionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "ListExpansions")
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalExpansions)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateExpansions : Create or update expansion list
// Create or replace the Expansion list for this collection. The maximum number of expanded terms per collection is
// `500`. The current expansion list is replaced with the uploaded content.
func (discovery *DiscoveryV1) CreateExpansions(createExpansionsOptions *CreateExpansionsOptions) (result *Expansions, response *core.DetailedResponse, err error) {
	return discovery.CreateExpansionsWithContext(context.Background(), createExpansionsOptions)
}

// CreateExpansionsWithContext is an alternate form of the CreateExpansions method which supports a Context parameter
func (discovery *DiscoveryV1) CreateExpansionsWithContext(ctx context.Context, createExpansionsOptions *CreateExpansionsOptions) (result *Expansions, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createExpansionsOptions, "createExpansionsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createExpansionsOptions, "createExpansionsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *createExpansionsOptions.EnvironmentID,
		"collection_id":  *createExpansionsOptions.CollectionID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/collections/{collection_id}/expansions`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createExpansionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "CreateExpansions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

	body := make(map[string]interface{})
	if createExpansionsOptions.Expansions != nil {
		body["expansions"] = createExpansionsOptions.Expansions
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalExpansions)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteExpansions : Delete the expansion list
// Remove the expansion information for this collection. The expansion list must be deleted to disable query expansion
// for a collection.
func (discovery *DiscoveryV1) DeleteExpansions(deleteExpansionsOptions *DeleteExpansionsOptions) (response *core.DetailedResponse, err error) {
	return discovery.DeleteExpansionsWithContext(context.Background(), deleteExpansionsOptions)
}

// DeleteExpansionsWithContext is an alternate form of the DeleteExpansions method which supports a Context parameter
func (discovery *DiscoveryV1) DeleteExpansionsWithContext(ctx context.Context, deleteExpansionsOptions *DeleteExpansionsOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteExpansionsOptions, "deleteExpansionsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteExpansionsOptions, "deleteExpansionsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *deleteExpansionsOptions.EnvironmentID,
		"collection_id":  *deleteExpansionsOptions.CollectionID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/collections/{collection_id}/expansions`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteExpansionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "DeleteExpansions")
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

// GetTokenizationDictionaryStatus : Get tokenization dictionary status
// Returns the current status of the tokenization dictionary for the specified collection.
func (discovery *DiscoveryV1) GetTokenizationDictionaryStatus(getTokenizationDictionaryStatusOptions *GetTokenizationDictionaryStatusOptions) (result *TokenDictStatusResponse, response *core.DetailedResponse, err error) {
	return discovery.GetTokenizationDictionaryStatusWithContext(context.Background(), getTokenizationDictionaryStatusOptions)
}

// GetTokenizationDictionaryStatusWithContext is an alternate form of the GetTokenizationDictionaryStatus method which supports a Context parameter
func (discovery *DiscoveryV1) GetTokenizationDictionaryStatusWithContext(ctx context.Context, getTokenizationDictionaryStatusOptions *GetTokenizationDictionaryStatusOptions) (result *TokenDictStatusResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getTokenizationDictionaryStatusOptions, "getTokenizationDictionaryStatusOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getTokenizationDictionaryStatusOptions, "getTokenizationDictionaryStatusOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *getTokenizationDictionaryStatusOptions.EnvironmentID,
		"collection_id":  *getTokenizationDictionaryStatusOptions.CollectionID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/collections/{collection_id}/word_lists/tokenization_dictionary`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getTokenizationDictionaryStatusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "GetTokenizationDictionaryStatus")
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTokenDictStatusResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateTokenizationDictionary : Create tokenization dictionary
// Upload a custom tokenization dictionary to use with the specified collection.
func (discovery *DiscoveryV1) CreateTokenizationDictionary(createTokenizationDictionaryOptions *CreateTokenizationDictionaryOptions) (result *TokenDictStatusResponse, response *core.DetailedResponse, err error) {
	return discovery.CreateTokenizationDictionaryWithContext(context.Background(), createTokenizationDictionaryOptions)
}

// CreateTokenizationDictionaryWithContext is an alternate form of the CreateTokenizationDictionary method which supports a Context parameter
func (discovery *DiscoveryV1) CreateTokenizationDictionaryWithContext(ctx context.Context, createTokenizationDictionaryOptions *CreateTokenizationDictionaryOptions) (result *TokenDictStatusResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createTokenizationDictionaryOptions, "createTokenizationDictionaryOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createTokenizationDictionaryOptions, "createTokenizationDictionaryOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *createTokenizationDictionaryOptions.EnvironmentID,
		"collection_id":  *createTokenizationDictionaryOptions.CollectionID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/collections/{collection_id}/word_lists/tokenization_dictionary`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createTokenizationDictionaryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "CreateTokenizationDictionary")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

	body := make(map[string]interface{})
	if createTokenizationDictionaryOptions.TokenizationRules != nil {
		body["tokenization_rules"] = createTokenizationDictionaryOptions.TokenizationRules
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTokenDictStatusResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteTokenizationDictionary : Delete tokenization dictionary
// Delete the tokenization dictionary from the collection.
func (discovery *DiscoveryV1) DeleteTokenizationDictionary(deleteTokenizationDictionaryOptions *DeleteTokenizationDictionaryOptions) (response *core.DetailedResponse, err error) {
	return discovery.DeleteTokenizationDictionaryWithContext(context.Background(), deleteTokenizationDictionaryOptions)
}

// DeleteTokenizationDictionaryWithContext is an alternate form of the DeleteTokenizationDictionary method which supports a Context parameter
func (discovery *DiscoveryV1) DeleteTokenizationDictionaryWithContext(ctx context.Context, deleteTokenizationDictionaryOptions *DeleteTokenizationDictionaryOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteTokenizationDictionaryOptions, "deleteTokenizationDictionaryOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteTokenizationDictionaryOptions, "deleteTokenizationDictionaryOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *deleteTokenizationDictionaryOptions.EnvironmentID,
		"collection_id":  *deleteTokenizationDictionaryOptions.CollectionID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/collections/{collection_id}/word_lists/tokenization_dictionary`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteTokenizationDictionaryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "DeleteTokenizationDictionary")
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

// GetStopwordListStatus : Get stopword list status
// Returns the current status of the stopword list for the specified collection.
func (discovery *DiscoveryV1) GetStopwordListStatus(getStopwordListStatusOptions *GetStopwordListStatusOptions) (result *TokenDictStatusResponse, response *core.DetailedResponse, err error) {
	return discovery.GetStopwordListStatusWithContext(context.Background(), getStopwordListStatusOptions)
}

// GetStopwordListStatusWithContext is an alternate form of the GetStopwordListStatus method which supports a Context parameter
func (discovery *DiscoveryV1) GetStopwordListStatusWithContext(ctx context.Context, getStopwordListStatusOptions *GetStopwordListStatusOptions) (result *TokenDictStatusResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getStopwordListStatusOptions, "getStopwordListStatusOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getStopwordListStatusOptions, "getStopwordListStatusOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *getStopwordListStatusOptions.EnvironmentID,
		"collection_id":  *getStopwordListStatusOptions.CollectionID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/collections/{collection_id}/word_lists/stopwords`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getStopwordListStatusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "GetStopwordListStatus")
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTokenDictStatusResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateStopwordList : Create stopword list
// Upload a custom stopword list to use with the specified collection.
func (discovery *DiscoveryV1) CreateStopwordList(createStopwordListOptions *CreateStopwordListOptions) (result *TokenDictStatusResponse, response *core.DetailedResponse, err error) {
	return discovery.CreateStopwordListWithContext(context.Background(), createStopwordListOptions)
}

// CreateStopwordListWithContext is an alternate form of the CreateStopwordList method which supports a Context parameter
func (discovery *DiscoveryV1) CreateStopwordListWithContext(ctx context.Context, createStopwordListOptions *CreateStopwordListOptions) (result *TokenDictStatusResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createStopwordListOptions, "createStopwordListOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createStopwordListOptions, "createStopwordListOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *createStopwordListOptions.EnvironmentID,
		"collection_id":  *createStopwordListOptions.CollectionID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/collections/{collection_id}/word_lists/stopwords`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createStopwordListOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "CreateStopwordList")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

	builder.AddFormData("stopword_file", core.StringNilMapper(createStopwordListOptions.StopwordFilename),
		"application/octet-stream", createStopwordListOptions.StopwordFile)

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = discovery.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTokenDictStatusResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteStopwordList : Delete a custom stopword list
// Delete a custom stopword list from the collection. After a custom stopword list is deleted, the default list is used
// for the collection.
func (discovery *DiscoveryV1) DeleteStopwordList(deleteStopwordListOptions *DeleteStopwordListOptions) (response *core.DetailedResponse, err error) {
	return discovery.DeleteStopwordListWithContext(context.Background(), deleteStopwordListOptions)
}

// DeleteStopwordListWithContext is an alternate form of the DeleteStopwordList method which supports a Context parameter
func (discovery *DiscoveryV1) DeleteStopwordListWithContext(ctx context.Context, deleteStopwordListOptions *DeleteStopwordListOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteStopwordListOptions, "deleteStopwordListOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteStopwordListOptions, "deleteStopwordListOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *deleteStopwordListOptions.EnvironmentID,
		"collection_id":  *deleteStopwordListOptions.CollectionID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/collections/{collection_id}/word_lists/stopwords`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteStopwordListOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "DeleteStopwordList")
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

// AddDocument : Add a document
// Add a document to a collection with optional metadata.
//
//   * The **version** query parameter is still required.
//
//   * Returns immediately after the system has accepted the document for processing.
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
//  **Note:** Documents can be added with a specific **document_id** by using the
// **_/v1/environments/{environment_id}/collections/{collection_id}/documents** method.
func (discovery *DiscoveryV1) AddDocument(addDocumentOptions *AddDocumentOptions) (result *DocumentAccepted, response *core.DetailedResponse, err error) {
	return discovery.AddDocumentWithContext(context.Background(), addDocumentOptions)
}

// AddDocumentWithContext is an alternate form of the AddDocument method which supports a Context parameter
func (discovery *DiscoveryV1) AddDocumentWithContext(ctx context.Context, addDocumentOptions *AddDocumentOptions) (result *DocumentAccepted, response *core.DetailedResponse, err error) {
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
		"environment_id": *addDocumentOptions.EnvironmentID,
		"collection_id":  *addDocumentOptions.CollectionID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/collections/{collection_id}/documents`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range addDocumentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "AddDocument")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDocumentAccepted)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetDocumentStatus : Get document details
// Fetch status details about a submitted document. **Note:** this operation does not return the document itself.
// Instead, it returns only the document's processing status and any notices (warnings or errors) that were generated
// when the document was ingested. Use the query API to retrieve the actual document content.
func (discovery *DiscoveryV1) GetDocumentStatus(getDocumentStatusOptions *GetDocumentStatusOptions) (result *DocumentStatus, response *core.DetailedResponse, err error) {
	return discovery.GetDocumentStatusWithContext(context.Background(), getDocumentStatusOptions)
}

// GetDocumentStatusWithContext is an alternate form of the GetDocumentStatus method which supports a Context parameter
func (discovery *DiscoveryV1) GetDocumentStatusWithContext(ctx context.Context, getDocumentStatusOptions *GetDocumentStatusOptions) (result *DocumentStatus, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDocumentStatusOptions, "getDocumentStatusOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getDocumentStatusOptions, "getDocumentStatusOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *getDocumentStatusOptions.EnvironmentID,
		"collection_id":  *getDocumentStatusOptions.CollectionID,
		"document_id":    *getDocumentStatusOptions.DocumentID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/collections/{collection_id}/documents/{document_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getDocumentStatusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "GetDocumentStatus")
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDocumentStatus)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateDocument : Update a document
// Replace an existing document or add a document with a specified **document_id**. Starts ingesting a document with
// optional metadata.
//
// **Note:** When uploading a new document with this method it automatically replaces any document stored with the same
// **document_id** if it exists.
func (discovery *DiscoveryV1) UpdateDocument(updateDocumentOptions *UpdateDocumentOptions) (result *DocumentAccepted, response *core.DetailedResponse, err error) {
	return discovery.UpdateDocumentWithContext(context.Background(), updateDocumentOptions)
}

// UpdateDocumentWithContext is an alternate form of the UpdateDocument method which supports a Context parameter
func (discovery *DiscoveryV1) UpdateDocumentWithContext(ctx context.Context, updateDocumentOptions *UpdateDocumentOptions) (result *DocumentAccepted, response *core.DetailedResponse, err error) {
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
		"environment_id": *updateDocumentOptions.EnvironmentID,
		"collection_id":  *updateDocumentOptions.CollectionID,
		"document_id":    *updateDocumentOptions.DocumentID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/collections/{collection_id}/documents/{document_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateDocumentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "UpdateDocument")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDocumentAccepted)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteDocument : Delete a document
// If the given document ID is invalid, or if the document is not found, then the a success response is returned (HTTP
// status code `200`) with the status set to 'deleted'.
func (discovery *DiscoveryV1) DeleteDocument(deleteDocumentOptions *DeleteDocumentOptions) (result *DeleteDocumentResponse, response *core.DetailedResponse, err error) {
	return discovery.DeleteDocumentWithContext(context.Background(), deleteDocumentOptions)
}

// DeleteDocumentWithContext is an alternate form of the DeleteDocument method which supports a Context parameter
func (discovery *DiscoveryV1) DeleteDocumentWithContext(ctx context.Context, deleteDocumentOptions *DeleteDocumentOptions) (result *DeleteDocumentResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteDocumentOptions, "deleteDocumentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteDocumentOptions, "deleteDocumentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *deleteDocumentOptions.EnvironmentID,
		"collection_id":  *deleteDocumentOptions.CollectionID,
		"document_id":    *deleteDocumentOptions.DocumentID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/collections/{collection_id}/documents/{document_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteDocumentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "DeleteDocument")
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeleteDocumentResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// Query : Query a collection
// By using this method, you can construct long queries. For details, see the [Discovery
// documentation](https://cloud.ibm.com/docs/discovery?topic=discovery-query-concepts#query-concepts).
func (discovery *DiscoveryV1) Query(queryOptions *QueryOptions) (result *QueryResponse, response *core.DetailedResponse, err error) {
	return discovery.QueryWithContext(context.Background(), queryOptions)
}

// QueryWithContext is an alternate form of the Query method which supports a Context parameter
func (discovery *DiscoveryV1) QueryWithContext(ctx context.Context, queryOptions *QueryOptions) (result *QueryResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(queryOptions, "queryOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(queryOptions, "queryOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *queryOptions.EnvironmentID,
		"collection_id":  *queryOptions.CollectionID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/collections/{collection_id}/query`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range queryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "Query")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if queryOptions.XWatsonLoggingOptOut != nil {
		builder.AddHeader("X-Watson-Logging-Opt-Out", fmt.Sprint(*queryOptions.XWatsonLoggingOptOut))
	}

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

	body := make(map[string]interface{})
	if queryOptions.Filter != nil {
		body["filter"] = queryOptions.Filter
	}
	if queryOptions.Query != nil {
		body["query"] = queryOptions.Query
	}
	if queryOptions.NaturalLanguageQuery != nil {
		body["natural_language_query"] = queryOptions.NaturalLanguageQuery
	}
	if queryOptions.Passages != nil {
		body["passages"] = queryOptions.Passages
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
	if queryOptions.PassagesFields != nil {
		body["passages.fields"] = queryOptions.PassagesFields
	}
	if queryOptions.PassagesCount != nil {
		body["passages.count"] = queryOptions.PassagesCount
	}
	if queryOptions.PassagesCharacters != nil {
		body["passages.characters"] = queryOptions.PassagesCharacters
	}
	if queryOptions.Deduplicate != nil {
		body["deduplicate"] = queryOptions.Deduplicate
	}
	if queryOptions.DeduplicateField != nil {
		body["deduplicate.field"] = queryOptions.DeduplicateField
	}
	if queryOptions.Similar != nil {
		body["similar"] = queryOptions.Similar
	}
	if queryOptions.SimilarDocumentIds != nil {
		body["similar.document_ids"] = queryOptions.SimilarDocumentIds
	}
	if queryOptions.SimilarFields != nil {
		body["similar.fields"] = queryOptions.SimilarFields
	}
	if queryOptions.Bias != nil {
		body["bias"] = queryOptions.Bias
	}
	if queryOptions.SpellingSuggestions != nil {
		body["spelling_suggestions"] = queryOptions.SpellingSuggestions
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalQueryResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// QueryNotices : Query system notices
// Queries for notices (errors or warnings) that might have been generated by the system. Notices are generated when
// ingesting documents and performing relevance training. See the [Discovery
// documentation](https://cloud.ibm.com/docs/discovery?topic=discovery-query-concepts#query-concepts) for more details
// on the query language.
func (discovery *DiscoveryV1) QueryNotices(queryNoticesOptions *QueryNoticesOptions) (result *QueryNoticesResponse, response *core.DetailedResponse, err error) {
	return discovery.QueryNoticesWithContext(context.Background(), queryNoticesOptions)
}

// QueryNoticesWithContext is an alternate form of the QueryNotices method which supports a Context parameter
func (discovery *DiscoveryV1) QueryNoticesWithContext(ctx context.Context, queryNoticesOptions *QueryNoticesOptions) (result *QueryNoticesResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(queryNoticesOptions, "queryNoticesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(queryNoticesOptions, "queryNoticesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *queryNoticesOptions.EnvironmentID,
		"collection_id":  *queryNoticesOptions.CollectionID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/collections/{collection_id}/notices`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range queryNoticesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "QueryNotices")
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
	if queryNoticesOptions.Passages != nil {
		builder.AddQuery("passages", fmt.Sprint(*queryNoticesOptions.Passages))
	}
	if queryNoticesOptions.Aggregation != nil {
		builder.AddQuery("aggregation", fmt.Sprint(*queryNoticesOptions.Aggregation))
	}
	if queryNoticesOptions.Count != nil {
		builder.AddQuery("count", fmt.Sprint(*queryNoticesOptions.Count))
	}
	if queryNoticesOptions.Return != nil {
		builder.AddQuery("return", strings.Join(queryNoticesOptions.Return, ","))
	}
	if queryNoticesOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*queryNoticesOptions.Offset))
	}
	if queryNoticesOptions.Sort != nil {
		builder.AddQuery("sort", strings.Join(queryNoticesOptions.Sort, ","))
	}
	if queryNoticesOptions.Highlight != nil {
		builder.AddQuery("highlight", fmt.Sprint(*queryNoticesOptions.Highlight))
	}
	if queryNoticesOptions.PassagesFields != nil {
		builder.AddQuery("passages.fields", strings.Join(queryNoticesOptions.PassagesFields, ","))
	}
	if queryNoticesOptions.PassagesCount != nil {
		builder.AddQuery("passages.count", fmt.Sprint(*queryNoticesOptions.PassagesCount))
	}
	if queryNoticesOptions.PassagesCharacters != nil {
		builder.AddQuery("passages.characters", fmt.Sprint(*queryNoticesOptions.PassagesCharacters))
	}
	if queryNoticesOptions.DeduplicateField != nil {
		builder.AddQuery("deduplicate.field", fmt.Sprint(*queryNoticesOptions.DeduplicateField))
	}
	if queryNoticesOptions.Similar != nil {
		builder.AddQuery("similar", fmt.Sprint(*queryNoticesOptions.Similar))
	}
	if queryNoticesOptions.SimilarDocumentIds != nil {
		builder.AddQuery("similar.document_ids", strings.Join(queryNoticesOptions.SimilarDocumentIds, ","))
	}
	if queryNoticesOptions.SimilarFields != nil {
		builder.AddQuery("similar.fields", strings.Join(queryNoticesOptions.SimilarFields, ","))
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalQueryNoticesResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// FederatedQuery : Query multiple collections
// By using this method, you can construct long queries that search multiple collection. For details, see the [Discovery
// documentation](https://cloud.ibm.com/docs/discovery?topic=discovery-query-concepts#query-concepts).
func (discovery *DiscoveryV1) FederatedQuery(federatedQueryOptions *FederatedQueryOptions) (result *QueryResponse, response *core.DetailedResponse, err error) {
	return discovery.FederatedQueryWithContext(context.Background(), federatedQueryOptions)
}

// FederatedQueryWithContext is an alternate form of the FederatedQuery method which supports a Context parameter
func (discovery *DiscoveryV1) FederatedQueryWithContext(ctx context.Context, federatedQueryOptions *FederatedQueryOptions) (result *QueryResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(federatedQueryOptions, "federatedQueryOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(federatedQueryOptions, "federatedQueryOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *federatedQueryOptions.EnvironmentID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/query`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range federatedQueryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "FederatedQuery")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if federatedQueryOptions.XWatsonLoggingOptOut != nil {
		builder.AddHeader("X-Watson-Logging-Opt-Out", fmt.Sprint(*federatedQueryOptions.XWatsonLoggingOptOut))
	}

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

	body := make(map[string]interface{})
	if federatedQueryOptions.CollectionIds != nil {
		body["collection_ids"] = federatedQueryOptions.CollectionIds
	}
	if federatedQueryOptions.Filter != nil {
		body["filter"] = federatedQueryOptions.Filter
	}
	if federatedQueryOptions.Query != nil {
		body["query"] = federatedQueryOptions.Query
	}
	if federatedQueryOptions.NaturalLanguageQuery != nil {
		body["natural_language_query"] = federatedQueryOptions.NaturalLanguageQuery
	}
	if federatedQueryOptions.Passages != nil {
		body["passages"] = federatedQueryOptions.Passages
	}
	if federatedQueryOptions.Aggregation != nil {
		body["aggregation"] = federatedQueryOptions.Aggregation
	}
	if federatedQueryOptions.Count != nil {
		body["count"] = federatedQueryOptions.Count
	}
	if federatedQueryOptions.Return != nil {
		body["return"] = federatedQueryOptions.Return
	}
	if federatedQueryOptions.Offset != nil {
		body["offset"] = federatedQueryOptions.Offset
	}
	if federatedQueryOptions.Sort != nil {
		body["sort"] = federatedQueryOptions.Sort
	}
	if federatedQueryOptions.Highlight != nil {
		body["highlight"] = federatedQueryOptions.Highlight
	}
	if federatedQueryOptions.PassagesFields != nil {
		body["passages.fields"] = federatedQueryOptions.PassagesFields
	}
	if federatedQueryOptions.PassagesCount != nil {
		body["passages.count"] = federatedQueryOptions.PassagesCount
	}
	if federatedQueryOptions.PassagesCharacters != nil {
		body["passages.characters"] = federatedQueryOptions.PassagesCharacters
	}
	if federatedQueryOptions.Deduplicate != nil {
		body["deduplicate"] = federatedQueryOptions.Deduplicate
	}
	if federatedQueryOptions.DeduplicateField != nil {
		body["deduplicate.field"] = federatedQueryOptions.DeduplicateField
	}
	if federatedQueryOptions.Similar != nil {
		body["similar"] = federatedQueryOptions.Similar
	}
	if federatedQueryOptions.SimilarDocumentIds != nil {
		body["similar.document_ids"] = federatedQueryOptions.SimilarDocumentIds
	}
	if federatedQueryOptions.SimilarFields != nil {
		body["similar.fields"] = federatedQueryOptions.SimilarFields
	}
	if federatedQueryOptions.Bias != nil {
		body["bias"] = federatedQueryOptions.Bias
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalQueryResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// FederatedQueryNotices : Query multiple collection system notices
// Queries for notices (errors or warnings) that might have been generated by the system. Notices are generated when
// ingesting documents and performing relevance training. See the [Discovery
// documentation](https://cloud.ibm.com/docs/discovery?topic=discovery-query-concepts#query-concepts) for more details
// on the query language.
func (discovery *DiscoveryV1) FederatedQueryNotices(federatedQueryNoticesOptions *FederatedQueryNoticesOptions) (result *QueryNoticesResponse, response *core.DetailedResponse, err error) {
	return discovery.FederatedQueryNoticesWithContext(context.Background(), federatedQueryNoticesOptions)
}

// FederatedQueryNoticesWithContext is an alternate form of the FederatedQueryNotices method which supports a Context parameter
func (discovery *DiscoveryV1) FederatedQueryNoticesWithContext(ctx context.Context, federatedQueryNoticesOptions *FederatedQueryNoticesOptions) (result *QueryNoticesResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(federatedQueryNoticesOptions, "federatedQueryNoticesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(federatedQueryNoticesOptions, "federatedQueryNoticesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *federatedQueryNoticesOptions.EnvironmentID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/notices`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range federatedQueryNoticesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "FederatedQueryNotices")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))
	builder.AddQuery("collection_ids", strings.Join(federatedQueryNoticesOptions.CollectionIds, ","))
	if federatedQueryNoticesOptions.Filter != nil {
		builder.AddQuery("filter", fmt.Sprint(*federatedQueryNoticesOptions.Filter))
	}
	if federatedQueryNoticesOptions.Query != nil {
		builder.AddQuery("query", fmt.Sprint(*federatedQueryNoticesOptions.Query))
	}
	if federatedQueryNoticesOptions.NaturalLanguageQuery != nil {
		builder.AddQuery("natural_language_query", fmt.Sprint(*federatedQueryNoticesOptions.NaturalLanguageQuery))
	}
	if federatedQueryNoticesOptions.Aggregation != nil {
		builder.AddQuery("aggregation", fmt.Sprint(*federatedQueryNoticesOptions.Aggregation))
	}
	if federatedQueryNoticesOptions.Count != nil {
		builder.AddQuery("count", fmt.Sprint(*federatedQueryNoticesOptions.Count))
	}
	if federatedQueryNoticesOptions.Return != nil {
		builder.AddQuery("return", strings.Join(federatedQueryNoticesOptions.Return, ","))
	}
	if federatedQueryNoticesOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*federatedQueryNoticesOptions.Offset))
	}
	if federatedQueryNoticesOptions.Sort != nil {
		builder.AddQuery("sort", strings.Join(federatedQueryNoticesOptions.Sort, ","))
	}
	if federatedQueryNoticesOptions.Highlight != nil {
		builder.AddQuery("highlight", fmt.Sprint(*federatedQueryNoticesOptions.Highlight))
	}
	if federatedQueryNoticesOptions.DeduplicateField != nil {
		builder.AddQuery("deduplicate.field", fmt.Sprint(*federatedQueryNoticesOptions.DeduplicateField))
	}
	if federatedQueryNoticesOptions.Similar != nil {
		builder.AddQuery("similar", fmt.Sprint(*federatedQueryNoticesOptions.Similar))
	}
	if federatedQueryNoticesOptions.SimilarDocumentIds != nil {
		builder.AddQuery("similar.document_ids", strings.Join(federatedQueryNoticesOptions.SimilarDocumentIds, ","))
	}
	if federatedQueryNoticesOptions.SimilarFields != nil {
		builder.AddQuery("similar.fields", strings.Join(federatedQueryNoticesOptions.SimilarFields, ","))
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalQueryNoticesResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetAutocompletion : Get Autocomplete Suggestions
// Returns completion query suggestions for the specified prefix.  /n/n **Important:** this method is only valid when
// using the Cloud Pak version of Discovery.
func (discovery *DiscoveryV1) GetAutocompletion(getAutocompletionOptions *GetAutocompletionOptions) (result *Completions, response *core.DetailedResponse, err error) {
	return discovery.GetAutocompletionWithContext(context.Background(), getAutocompletionOptions)
}

// GetAutocompletionWithContext is an alternate form of the GetAutocompletion method which supports a Context parameter
func (discovery *DiscoveryV1) GetAutocompletionWithContext(ctx context.Context, getAutocompletionOptions *GetAutocompletionOptions) (result *Completions, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAutocompletionOptions, "getAutocompletionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAutocompletionOptions, "getAutocompletionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *getAutocompletionOptions.EnvironmentID,
		"collection_id":  *getAutocompletionOptions.CollectionID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/collections/{collection_id}/autocompletion`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAutocompletionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "GetAutocompletion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))
	builder.AddQuery("prefix", fmt.Sprint(*getAutocompletionOptions.Prefix))
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCompletions)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListTrainingData : List training data
// Lists the training data for the specified collection.
func (discovery *DiscoveryV1) ListTrainingData(listTrainingDataOptions *ListTrainingDataOptions) (result *TrainingDataSet, response *core.DetailedResponse, err error) {
	return discovery.ListTrainingDataWithContext(context.Background(), listTrainingDataOptions)
}

// ListTrainingDataWithContext is an alternate form of the ListTrainingData method which supports a Context parameter
func (discovery *DiscoveryV1) ListTrainingDataWithContext(ctx context.Context, listTrainingDataOptions *ListTrainingDataOptions) (result *TrainingDataSet, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listTrainingDataOptions, "listTrainingDataOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listTrainingDataOptions, "listTrainingDataOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *listTrainingDataOptions.EnvironmentID,
		"collection_id":  *listTrainingDataOptions.CollectionID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/collections/{collection_id}/training_data`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listTrainingDataOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "ListTrainingData")
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTrainingDataSet)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// AddTrainingData : Add query to training data
// Adds a query to the training data for this collection. The query can contain a filter and natural language query.
func (discovery *DiscoveryV1) AddTrainingData(addTrainingDataOptions *AddTrainingDataOptions) (result *TrainingQuery, response *core.DetailedResponse, err error) {
	return discovery.AddTrainingDataWithContext(context.Background(), addTrainingDataOptions)
}

// AddTrainingDataWithContext is an alternate form of the AddTrainingData method which supports a Context parameter
func (discovery *DiscoveryV1) AddTrainingDataWithContext(ctx context.Context, addTrainingDataOptions *AddTrainingDataOptions) (result *TrainingQuery, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addTrainingDataOptions, "addTrainingDataOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addTrainingDataOptions, "addTrainingDataOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *addTrainingDataOptions.EnvironmentID,
		"collection_id":  *addTrainingDataOptions.CollectionID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/collections/{collection_id}/training_data`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range addTrainingDataOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "AddTrainingData")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

	body := make(map[string]interface{})
	if addTrainingDataOptions.NaturalLanguageQuery != nil {
		body["natural_language_query"] = addTrainingDataOptions.NaturalLanguageQuery
	}
	if addTrainingDataOptions.Filter != nil {
		body["filter"] = addTrainingDataOptions.Filter
	}
	if addTrainingDataOptions.Examples != nil {
		body["examples"] = addTrainingDataOptions.Examples
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTrainingQuery)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteAllTrainingData : Delete all training data
// Deletes all training data from a collection.
func (discovery *DiscoveryV1) DeleteAllTrainingData(deleteAllTrainingDataOptions *DeleteAllTrainingDataOptions) (response *core.DetailedResponse, err error) {
	return discovery.DeleteAllTrainingDataWithContext(context.Background(), deleteAllTrainingDataOptions)
}

// DeleteAllTrainingDataWithContext is an alternate form of the DeleteAllTrainingData method which supports a Context parameter
func (discovery *DiscoveryV1) DeleteAllTrainingDataWithContext(ctx context.Context, deleteAllTrainingDataOptions *DeleteAllTrainingDataOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteAllTrainingDataOptions, "deleteAllTrainingDataOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteAllTrainingDataOptions, "deleteAllTrainingDataOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *deleteAllTrainingDataOptions.EnvironmentID,
		"collection_id":  *deleteAllTrainingDataOptions.CollectionID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/collections/{collection_id}/training_data`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteAllTrainingDataOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "DeleteAllTrainingData")
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

// GetTrainingData : Get details about a query
// Gets details for a specific training data query, including the query string and all examples.
func (discovery *DiscoveryV1) GetTrainingData(getTrainingDataOptions *GetTrainingDataOptions) (result *TrainingQuery, response *core.DetailedResponse, err error) {
	return discovery.GetTrainingDataWithContext(context.Background(), getTrainingDataOptions)
}

// GetTrainingDataWithContext is an alternate form of the GetTrainingData method which supports a Context parameter
func (discovery *DiscoveryV1) GetTrainingDataWithContext(ctx context.Context, getTrainingDataOptions *GetTrainingDataOptions) (result *TrainingQuery, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getTrainingDataOptions, "getTrainingDataOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getTrainingDataOptions, "getTrainingDataOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *getTrainingDataOptions.EnvironmentID,
		"collection_id":  *getTrainingDataOptions.CollectionID,
		"query_id":       *getTrainingDataOptions.QueryID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getTrainingDataOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "GetTrainingData")
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTrainingQuery)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteTrainingData : Delete a training data query
// Removes the training data query and all associated examples from the training data set.
func (discovery *DiscoveryV1) DeleteTrainingData(deleteTrainingDataOptions *DeleteTrainingDataOptions) (response *core.DetailedResponse, err error) {
	return discovery.DeleteTrainingDataWithContext(context.Background(), deleteTrainingDataOptions)
}

// DeleteTrainingDataWithContext is an alternate form of the DeleteTrainingData method which supports a Context parameter
func (discovery *DiscoveryV1) DeleteTrainingDataWithContext(ctx context.Context, deleteTrainingDataOptions *DeleteTrainingDataOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteTrainingDataOptions, "deleteTrainingDataOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteTrainingDataOptions, "deleteTrainingDataOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *deleteTrainingDataOptions.EnvironmentID,
		"collection_id":  *deleteTrainingDataOptions.CollectionID,
		"query_id":       *deleteTrainingDataOptions.QueryID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteTrainingDataOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "DeleteTrainingData")
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

// ListTrainingExamples : List examples for a training data query
// List all examples for this training data query.
func (discovery *DiscoveryV1) ListTrainingExamples(listTrainingExamplesOptions *ListTrainingExamplesOptions) (result *TrainingExampleList, response *core.DetailedResponse, err error) {
	return discovery.ListTrainingExamplesWithContext(context.Background(), listTrainingExamplesOptions)
}

// ListTrainingExamplesWithContext is an alternate form of the ListTrainingExamples method which supports a Context parameter
func (discovery *DiscoveryV1) ListTrainingExamplesWithContext(ctx context.Context, listTrainingExamplesOptions *ListTrainingExamplesOptions) (result *TrainingExampleList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listTrainingExamplesOptions, "listTrainingExamplesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listTrainingExamplesOptions, "listTrainingExamplesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *listTrainingExamplesOptions.EnvironmentID,
		"collection_id":  *listTrainingExamplesOptions.CollectionID,
		"query_id":       *listTrainingExamplesOptions.QueryID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}/examples`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listTrainingExamplesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "ListTrainingExamples")
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTrainingExampleList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateTrainingExample : Add example to training data query
// Adds a example to this training data query.
func (discovery *DiscoveryV1) CreateTrainingExample(createTrainingExampleOptions *CreateTrainingExampleOptions) (result *TrainingExample, response *core.DetailedResponse, err error) {
	return discovery.CreateTrainingExampleWithContext(context.Background(), createTrainingExampleOptions)
}

// CreateTrainingExampleWithContext is an alternate form of the CreateTrainingExample method which supports a Context parameter
func (discovery *DiscoveryV1) CreateTrainingExampleWithContext(ctx context.Context, createTrainingExampleOptions *CreateTrainingExampleOptions) (result *TrainingExample, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createTrainingExampleOptions, "createTrainingExampleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createTrainingExampleOptions, "createTrainingExampleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *createTrainingExampleOptions.EnvironmentID,
		"collection_id":  *createTrainingExampleOptions.CollectionID,
		"query_id":       *createTrainingExampleOptions.QueryID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}/examples`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createTrainingExampleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "CreateTrainingExample")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

	body := make(map[string]interface{})
	if createTrainingExampleOptions.DocumentID != nil {
		body["document_id"] = createTrainingExampleOptions.DocumentID
	}
	if createTrainingExampleOptions.CrossReference != nil {
		body["cross_reference"] = createTrainingExampleOptions.CrossReference
	}
	if createTrainingExampleOptions.Relevance != nil {
		body["relevance"] = createTrainingExampleOptions.Relevance
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTrainingExample)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteTrainingExample : Delete example for training data query
// Deletes the example document with the given ID from the training data query.
func (discovery *DiscoveryV1) DeleteTrainingExample(deleteTrainingExampleOptions *DeleteTrainingExampleOptions) (response *core.DetailedResponse, err error) {
	return discovery.DeleteTrainingExampleWithContext(context.Background(), deleteTrainingExampleOptions)
}

// DeleteTrainingExampleWithContext is an alternate form of the DeleteTrainingExample method which supports a Context parameter
func (discovery *DiscoveryV1) DeleteTrainingExampleWithContext(ctx context.Context, deleteTrainingExampleOptions *DeleteTrainingExampleOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteTrainingExampleOptions, "deleteTrainingExampleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteTrainingExampleOptions, "deleteTrainingExampleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *deleteTrainingExampleOptions.EnvironmentID,
		"collection_id":  *deleteTrainingExampleOptions.CollectionID,
		"query_id":       *deleteTrainingExampleOptions.QueryID,
		"example_id":     *deleteTrainingExampleOptions.ExampleID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}/examples/{example_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteTrainingExampleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "DeleteTrainingExample")
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

// UpdateTrainingExample : Change label or cross reference for example
// Changes the label or cross reference query for this training data example.
func (discovery *DiscoveryV1) UpdateTrainingExample(updateTrainingExampleOptions *UpdateTrainingExampleOptions) (result *TrainingExample, response *core.DetailedResponse, err error) {
	return discovery.UpdateTrainingExampleWithContext(context.Background(), updateTrainingExampleOptions)
}

// UpdateTrainingExampleWithContext is an alternate form of the UpdateTrainingExample method which supports a Context parameter
func (discovery *DiscoveryV1) UpdateTrainingExampleWithContext(ctx context.Context, updateTrainingExampleOptions *UpdateTrainingExampleOptions) (result *TrainingExample, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateTrainingExampleOptions, "updateTrainingExampleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateTrainingExampleOptions, "updateTrainingExampleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *updateTrainingExampleOptions.EnvironmentID,
		"collection_id":  *updateTrainingExampleOptions.CollectionID,
		"query_id":       *updateTrainingExampleOptions.QueryID,
		"example_id":     *updateTrainingExampleOptions.ExampleID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}/examples/{example_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateTrainingExampleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "UpdateTrainingExample")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

	body := make(map[string]interface{})
	if updateTrainingExampleOptions.CrossReference != nil {
		body["cross_reference"] = updateTrainingExampleOptions.CrossReference
	}
	if updateTrainingExampleOptions.Relevance != nil {
		body["relevance"] = updateTrainingExampleOptions.Relevance
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTrainingExample)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetTrainingExample : Get details for training data example
// Gets the details for this training example.
func (discovery *DiscoveryV1) GetTrainingExample(getTrainingExampleOptions *GetTrainingExampleOptions) (result *TrainingExample, response *core.DetailedResponse, err error) {
	return discovery.GetTrainingExampleWithContext(context.Background(), getTrainingExampleOptions)
}

// GetTrainingExampleWithContext is an alternate form of the GetTrainingExample method which supports a Context parameter
func (discovery *DiscoveryV1) GetTrainingExampleWithContext(ctx context.Context, getTrainingExampleOptions *GetTrainingExampleOptions) (result *TrainingExample, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getTrainingExampleOptions, "getTrainingExampleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getTrainingExampleOptions, "getTrainingExampleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *getTrainingExampleOptions.EnvironmentID,
		"collection_id":  *getTrainingExampleOptions.CollectionID,
		"query_id":       *getTrainingExampleOptions.QueryID,
		"example_id":     *getTrainingExampleOptions.ExampleID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}/examples/{example_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getTrainingExampleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "GetTrainingExample")
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTrainingExample)
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
// You associate a customer ID with data by passing the **X-Watson-Metadata** header with a request that passes data.
// For more information about personal data and customer IDs, see [Information
// security](https://cloud.ibm.com/docs/discovery?topic=discovery-information-security#information-security).
func (discovery *DiscoveryV1) DeleteUserData(deleteUserDataOptions *DeleteUserDataOptions) (response *core.DetailedResponse, err error) {
	return discovery.DeleteUserDataWithContext(context.Background(), deleteUserDataOptions)
}

// DeleteUserDataWithContext is an alternate form of the DeleteUserData method which supports a Context parameter
func (discovery *DiscoveryV1) DeleteUserDataWithContext(ctx context.Context, deleteUserDataOptions *DeleteUserDataOptions) (response *core.DetailedResponse, err error) {
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
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/user_data`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteUserDataOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "DeleteUserData")
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

// CreateEvent : Create event
// The **Events** API can be used to create log entries that are associated with specific queries. For example, you can
// record which documents in the results set were "clicked" by a user and when that click occurred.
func (discovery *DiscoveryV1) CreateEvent(createEventOptions *CreateEventOptions) (result *CreateEventResponse, response *core.DetailedResponse, err error) {
	return discovery.CreateEventWithContext(context.Background(), createEventOptions)
}

// CreateEventWithContext is an alternate form of the CreateEvent method which supports a Context parameter
func (discovery *DiscoveryV1) CreateEventWithContext(ctx context.Context, createEventOptions *CreateEventOptions) (result *CreateEventResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createEventOptions, "createEventOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createEventOptions, "createEventOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/events`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createEventOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "CreateEvent")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

	body := make(map[string]interface{})
	if createEventOptions.Type != nil {
		body["type"] = createEventOptions.Type
	}
	if createEventOptions.Data != nil {
		body["data"] = createEventOptions.Data
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreateEventResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// QueryLog : Search the query and event log
// Searches the query and event log to find query sessions that match the specified criteria. Searching the **logs**
// endpoint uses the standard Discovery query syntax for the parameters that are supported.
func (discovery *DiscoveryV1) QueryLog(queryLogOptions *QueryLogOptions) (result *LogQueryResponse, response *core.DetailedResponse, err error) {
	return discovery.QueryLogWithContext(context.Background(), queryLogOptions)
}

// QueryLogWithContext is an alternate form of the QueryLog method which supports a Context parameter
func (discovery *DiscoveryV1) QueryLogWithContext(ctx context.Context, queryLogOptions *QueryLogOptions) (result *LogQueryResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(queryLogOptions, "queryLogOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/logs`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range queryLogOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "QueryLog")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))
	if queryLogOptions.Filter != nil {
		builder.AddQuery("filter", fmt.Sprint(*queryLogOptions.Filter))
	}
	if queryLogOptions.Query != nil {
		builder.AddQuery("query", fmt.Sprint(*queryLogOptions.Query))
	}
	if queryLogOptions.Count != nil {
		builder.AddQuery("count", fmt.Sprint(*queryLogOptions.Count))
	}
	if queryLogOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*queryLogOptions.Offset))
	}
	if queryLogOptions.Sort != nil {
		builder.AddQuery("sort", strings.Join(queryLogOptions.Sort, ","))
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalLogQueryResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetMetricsQuery : Number of queries over time
// Total number of queries using the **natural_language_query** parameter over a specific time window.
func (discovery *DiscoveryV1) GetMetricsQuery(getMetricsQueryOptions *GetMetricsQueryOptions) (result *MetricResponse, response *core.DetailedResponse, err error) {
	return discovery.GetMetricsQueryWithContext(context.Background(), getMetricsQueryOptions)
}

// GetMetricsQueryWithContext is an alternate form of the GetMetricsQuery method which supports a Context parameter
func (discovery *DiscoveryV1) GetMetricsQueryWithContext(ctx context.Context, getMetricsQueryOptions *GetMetricsQueryOptions) (result *MetricResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getMetricsQueryOptions, "getMetricsQueryOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/metrics/number_of_queries`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getMetricsQueryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "GetMetricsQuery")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))
	if getMetricsQueryOptions.StartTime != nil {
		builder.AddQuery("start_time", fmt.Sprint(*getMetricsQueryOptions.StartTime))
	}
	if getMetricsQueryOptions.EndTime != nil {
		builder.AddQuery("end_time", fmt.Sprint(*getMetricsQueryOptions.EndTime))
	}
	if getMetricsQueryOptions.ResultType != nil {
		builder.AddQuery("result_type", fmt.Sprint(*getMetricsQueryOptions.ResultType))
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMetricResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetMetricsQueryEvent : Number of queries with an event over time
// Total number of queries using the **natural_language_query** parameter that have a corresponding "click" event over a
// specified time window. This metric requires having integrated event tracking in your application using the **Events**
// API.
func (discovery *DiscoveryV1) GetMetricsQueryEvent(getMetricsQueryEventOptions *GetMetricsQueryEventOptions) (result *MetricResponse, response *core.DetailedResponse, err error) {
	return discovery.GetMetricsQueryEventWithContext(context.Background(), getMetricsQueryEventOptions)
}

// GetMetricsQueryEventWithContext is an alternate form of the GetMetricsQueryEvent method which supports a Context parameter
func (discovery *DiscoveryV1) GetMetricsQueryEventWithContext(ctx context.Context, getMetricsQueryEventOptions *GetMetricsQueryEventOptions) (result *MetricResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getMetricsQueryEventOptions, "getMetricsQueryEventOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/metrics/number_of_queries_with_event`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getMetricsQueryEventOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "GetMetricsQueryEvent")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))
	if getMetricsQueryEventOptions.StartTime != nil {
		builder.AddQuery("start_time", fmt.Sprint(*getMetricsQueryEventOptions.StartTime))
	}
	if getMetricsQueryEventOptions.EndTime != nil {
		builder.AddQuery("end_time", fmt.Sprint(*getMetricsQueryEventOptions.EndTime))
	}
	if getMetricsQueryEventOptions.ResultType != nil {
		builder.AddQuery("result_type", fmt.Sprint(*getMetricsQueryEventOptions.ResultType))
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMetricResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetMetricsQueryNoResults : Number of queries with no search results over time
// Total number of queries using the **natural_language_query** parameter that have no results returned over a specified
// time window.
func (discovery *DiscoveryV1) GetMetricsQueryNoResults(getMetricsQueryNoResultsOptions *GetMetricsQueryNoResultsOptions) (result *MetricResponse, response *core.DetailedResponse, err error) {
	return discovery.GetMetricsQueryNoResultsWithContext(context.Background(), getMetricsQueryNoResultsOptions)
}

// GetMetricsQueryNoResultsWithContext is an alternate form of the GetMetricsQueryNoResults method which supports a Context parameter
func (discovery *DiscoveryV1) GetMetricsQueryNoResultsWithContext(ctx context.Context, getMetricsQueryNoResultsOptions *GetMetricsQueryNoResultsOptions) (result *MetricResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getMetricsQueryNoResultsOptions, "getMetricsQueryNoResultsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/metrics/number_of_queries_with_no_search_results`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getMetricsQueryNoResultsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "GetMetricsQueryNoResults")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))
	if getMetricsQueryNoResultsOptions.StartTime != nil {
		builder.AddQuery("start_time", fmt.Sprint(*getMetricsQueryNoResultsOptions.StartTime))
	}
	if getMetricsQueryNoResultsOptions.EndTime != nil {
		builder.AddQuery("end_time", fmt.Sprint(*getMetricsQueryNoResultsOptions.EndTime))
	}
	if getMetricsQueryNoResultsOptions.ResultType != nil {
		builder.AddQuery("result_type", fmt.Sprint(*getMetricsQueryNoResultsOptions.ResultType))
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMetricResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetMetricsEventRate : Percentage of queries with an associated event
// The percentage of queries using the **natural_language_query** parameter that have a corresponding "click" event over
// a specified time window.  This metric requires having integrated event tracking in your application using the
// **Events** API.
func (discovery *DiscoveryV1) GetMetricsEventRate(getMetricsEventRateOptions *GetMetricsEventRateOptions) (result *MetricResponse, response *core.DetailedResponse, err error) {
	return discovery.GetMetricsEventRateWithContext(context.Background(), getMetricsEventRateOptions)
}

// GetMetricsEventRateWithContext is an alternate form of the GetMetricsEventRate method which supports a Context parameter
func (discovery *DiscoveryV1) GetMetricsEventRateWithContext(ctx context.Context, getMetricsEventRateOptions *GetMetricsEventRateOptions) (result *MetricResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getMetricsEventRateOptions, "getMetricsEventRateOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/metrics/event_rate`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getMetricsEventRateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "GetMetricsEventRate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))
	if getMetricsEventRateOptions.StartTime != nil {
		builder.AddQuery("start_time", fmt.Sprint(*getMetricsEventRateOptions.StartTime))
	}
	if getMetricsEventRateOptions.EndTime != nil {
		builder.AddQuery("end_time", fmt.Sprint(*getMetricsEventRateOptions.EndTime))
	}
	if getMetricsEventRateOptions.ResultType != nil {
		builder.AddQuery("result_type", fmt.Sprint(*getMetricsEventRateOptions.ResultType))
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMetricResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetMetricsQueryTokenEvent : Most frequent query tokens with an event
// The most frequent query tokens parsed from the **natural_language_query** parameter and their corresponding "click"
// event rate within the recording period (queries and events are stored for 30 days). A query token is an individual
// word or unigram within the query string.
func (discovery *DiscoveryV1) GetMetricsQueryTokenEvent(getMetricsQueryTokenEventOptions *GetMetricsQueryTokenEventOptions) (result *MetricTokenResponse, response *core.DetailedResponse, err error) {
	return discovery.GetMetricsQueryTokenEventWithContext(context.Background(), getMetricsQueryTokenEventOptions)
}

// GetMetricsQueryTokenEventWithContext is an alternate form of the GetMetricsQueryTokenEvent method which supports a Context parameter
func (discovery *DiscoveryV1) GetMetricsQueryTokenEventWithContext(ctx context.Context, getMetricsQueryTokenEventOptions *GetMetricsQueryTokenEventOptions) (result *MetricTokenResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getMetricsQueryTokenEventOptions, "getMetricsQueryTokenEventOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/metrics/top_query_tokens_with_event_rate`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getMetricsQueryTokenEventOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "GetMetricsQueryTokenEvent")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))
	if getMetricsQueryTokenEventOptions.Count != nil {
		builder.AddQuery("count", fmt.Sprint(*getMetricsQueryTokenEventOptions.Count))
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMetricTokenResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListCredentials : List credentials
// List all the source credentials that have been created for this service instance.
//
//  **Note:**  All credentials are sent over an encrypted connection and encrypted at rest.
func (discovery *DiscoveryV1) ListCredentials(listCredentialsOptions *ListCredentialsOptions) (result *CredentialsList, response *core.DetailedResponse, err error) {
	return discovery.ListCredentialsWithContext(context.Background(), listCredentialsOptions)
}

// ListCredentialsWithContext is an alternate form of the ListCredentials method which supports a Context parameter
func (discovery *DiscoveryV1) ListCredentialsWithContext(ctx context.Context, listCredentialsOptions *ListCredentialsOptions) (result *CredentialsList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listCredentialsOptions, "listCredentialsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listCredentialsOptions, "listCredentialsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *listCredentialsOptions.EnvironmentID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/credentials`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listCredentialsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "ListCredentials")
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCredentialsList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateCredentials : Create credentials
// Creates a set of credentials to connect to a remote source. Created credentials are used in a configuration to
// associate a collection with the remote source.
//
// **Note:** All credentials are sent over an encrypted connection and encrypted at rest.
func (discovery *DiscoveryV1) CreateCredentials(createCredentialsOptions *CreateCredentialsOptions) (result *Credentials, response *core.DetailedResponse, err error) {
	return discovery.CreateCredentialsWithContext(context.Background(), createCredentialsOptions)
}

// CreateCredentialsWithContext is an alternate form of the CreateCredentials method which supports a Context parameter
func (discovery *DiscoveryV1) CreateCredentialsWithContext(ctx context.Context, createCredentialsOptions *CreateCredentialsOptions) (result *Credentials, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createCredentialsOptions, "createCredentialsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createCredentialsOptions, "createCredentialsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *createCredentialsOptions.EnvironmentID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/credentials`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createCredentialsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "CreateCredentials")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

	body := make(map[string]interface{})
	if createCredentialsOptions.SourceType != nil {
		body["source_type"] = createCredentialsOptions.SourceType
	}
	if createCredentialsOptions.CredentialDetails != nil {
		body["credential_details"] = createCredentialsOptions.CredentialDetails
	}
	if createCredentialsOptions.Status != nil {
		body["status"] = createCredentialsOptions.Status
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCredentials)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetCredentials : View Credentials
// Returns details about the specified credentials.
//
//  **Note:** Secure credential information such as a password or SSH key is never returned and must be obtained from
// the source system.
func (discovery *DiscoveryV1) GetCredentials(getCredentialsOptions *GetCredentialsOptions) (result *Credentials, response *core.DetailedResponse, err error) {
	return discovery.GetCredentialsWithContext(context.Background(), getCredentialsOptions)
}

// GetCredentialsWithContext is an alternate form of the GetCredentials method which supports a Context parameter
func (discovery *DiscoveryV1) GetCredentialsWithContext(ctx context.Context, getCredentialsOptions *GetCredentialsOptions) (result *Credentials, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getCredentialsOptions, "getCredentialsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getCredentialsOptions, "getCredentialsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *getCredentialsOptions.EnvironmentID,
		"credential_id":  *getCredentialsOptions.CredentialID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/credentials/{credential_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCredentialsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "GetCredentials")
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCredentials)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateCredentials : Update credentials
// Updates an existing set of source credentials.
//
// **Note:** All credentials are sent over an encrypted connection and encrypted at rest.
func (discovery *DiscoveryV1) UpdateCredentials(updateCredentialsOptions *UpdateCredentialsOptions) (result *Credentials, response *core.DetailedResponse, err error) {
	return discovery.UpdateCredentialsWithContext(context.Background(), updateCredentialsOptions)
}

// UpdateCredentialsWithContext is an alternate form of the UpdateCredentials method which supports a Context parameter
func (discovery *DiscoveryV1) UpdateCredentialsWithContext(ctx context.Context, updateCredentialsOptions *UpdateCredentialsOptions) (result *Credentials, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateCredentialsOptions, "updateCredentialsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateCredentialsOptions, "updateCredentialsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *updateCredentialsOptions.EnvironmentID,
		"credential_id":  *updateCredentialsOptions.CredentialID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/credentials/{credential_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateCredentialsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "UpdateCredentials")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

	body := make(map[string]interface{})
	if updateCredentialsOptions.SourceType != nil {
		body["source_type"] = updateCredentialsOptions.SourceType
	}
	if updateCredentialsOptions.CredentialDetails != nil {
		body["credential_details"] = updateCredentialsOptions.CredentialDetails
	}
	if updateCredentialsOptions.Status != nil {
		body["status"] = updateCredentialsOptions.Status
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCredentials)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteCredentials : Delete credentials
// Deletes a set of stored credentials from your Discovery instance.
func (discovery *DiscoveryV1) DeleteCredentials(deleteCredentialsOptions *DeleteCredentialsOptions) (result *DeleteCredentials, response *core.DetailedResponse, err error) {
	return discovery.DeleteCredentialsWithContext(context.Background(), deleteCredentialsOptions)
}

// DeleteCredentialsWithContext is an alternate form of the DeleteCredentials method which supports a Context parameter
func (discovery *DiscoveryV1) DeleteCredentialsWithContext(ctx context.Context, deleteCredentialsOptions *DeleteCredentialsOptions) (result *DeleteCredentials, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteCredentialsOptions, "deleteCredentialsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteCredentialsOptions, "deleteCredentialsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *deleteCredentialsOptions.EnvironmentID,
		"credential_id":  *deleteCredentialsOptions.CredentialID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/credentials/{credential_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteCredentialsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "DeleteCredentials")
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeleteCredentials)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListGateways : List Gateways
// List the currently configured gateways.
func (discovery *DiscoveryV1) ListGateways(listGatewaysOptions *ListGatewaysOptions) (result *GatewayList, response *core.DetailedResponse, err error) {
	return discovery.ListGatewaysWithContext(context.Background(), listGatewaysOptions)
}

// ListGatewaysWithContext is an alternate form of the ListGateways method which supports a Context parameter
func (discovery *DiscoveryV1) ListGatewaysWithContext(ctx context.Context, listGatewaysOptions *ListGatewaysOptions) (result *GatewayList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listGatewaysOptions, "listGatewaysOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listGatewaysOptions, "listGatewaysOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *listGatewaysOptions.EnvironmentID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/gateways`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listGatewaysOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "ListGateways")
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGatewayList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateGateway : Create Gateway
// Create a gateway configuration to use with a remotely installed gateway.
func (discovery *DiscoveryV1) CreateGateway(createGatewayOptions *CreateGatewayOptions) (result *Gateway, response *core.DetailedResponse, err error) {
	return discovery.CreateGatewayWithContext(context.Background(), createGatewayOptions)
}

// CreateGatewayWithContext is an alternate form of the CreateGateway method which supports a Context parameter
func (discovery *DiscoveryV1) CreateGatewayWithContext(ctx context.Context, createGatewayOptions *CreateGatewayOptions) (result *Gateway, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createGatewayOptions, "createGatewayOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createGatewayOptions, "createGatewayOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *createGatewayOptions.EnvironmentID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/gateways`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createGatewayOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "CreateGateway")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*discovery.Version))

	body := make(map[string]interface{})
	if createGatewayOptions.Name != nil {
		body["name"] = createGatewayOptions.Name
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGateway)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetGateway : List Gateway Details
// List information about the specified gateway.
func (discovery *DiscoveryV1) GetGateway(getGatewayOptions *GetGatewayOptions) (result *Gateway, response *core.DetailedResponse, err error) {
	return discovery.GetGatewayWithContext(context.Background(), getGatewayOptions)
}

// GetGatewayWithContext is an alternate form of the GetGateway method which supports a Context parameter
func (discovery *DiscoveryV1) GetGatewayWithContext(ctx context.Context, getGatewayOptions *GetGatewayOptions) (result *Gateway, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getGatewayOptions, "getGatewayOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getGatewayOptions, "getGatewayOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *getGatewayOptions.EnvironmentID,
		"gateway_id":     *getGatewayOptions.GatewayID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/gateways/{gateway_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getGatewayOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "GetGateway")
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGateway)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteGateway : Delete Gateway
// Delete the specified gateway configuration.
func (discovery *DiscoveryV1) DeleteGateway(deleteGatewayOptions *DeleteGatewayOptions) (result *GatewayDelete, response *core.DetailedResponse, err error) {
	return discovery.DeleteGatewayWithContext(context.Background(), deleteGatewayOptions)
}

// DeleteGatewayWithContext is an alternate form of the DeleteGateway method which supports a Context parameter
func (discovery *DiscoveryV1) DeleteGatewayWithContext(ctx context.Context, deleteGatewayOptions *DeleteGatewayOptions) (result *GatewayDelete, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteGatewayOptions, "deleteGatewayOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteGatewayOptions, "deleteGatewayOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"environment_id": *deleteGatewayOptions.EnvironmentID,
		"gateway_id":     *deleteGatewayOptions.GatewayID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = discovery.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(discovery.Service.Options.URL, `/v1/environments/{environment_id}/gateways/{gateway_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteGatewayOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("discovery", "V1", "DeleteGateway")
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
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGatewayDelete)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// AddDocumentOptions : The AddDocument options.
type AddDocumentOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The ID of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// The content of the document to ingest. The maximum supported file size when adding a file to a collection is 50
	// megabytes, the maximum supported file size when testing a configuration is 1 megabyte. Files larger than the
	// supported size are rejected.
	File io.ReadCloser `json:"-"`

	// The filename for file.
	Filename *string `json:"-"`

	// The content type of file.
	FileContentType *string `json:"-"`

	// The maximum supported metadata file size is 1 MB. Metadata parts larger than 1 MB are rejected. Example:  ``` {
	//   "Creator": "Johnny Appleseed",
	//   "Subject": "Apples"
	// } ```.
	Metadata *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAddDocumentOptions : Instantiate AddDocumentOptions
func (*DiscoveryV1) NewAddDocumentOptions(environmentID string, collectionID string) *AddDocumentOptions {
	return &AddDocumentOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *AddDocumentOptions) SetEnvironmentID(environmentID string) *AddDocumentOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetCollectionID : Allow user to set CollectionID
func (_options *AddDocumentOptions) SetCollectionID(collectionID string) *AddDocumentOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetFile : Allow user to set File
func (_options *AddDocumentOptions) SetFile(file io.ReadCloser) *AddDocumentOptions {
	_options.File = file
	return _options
}

// SetFilename : Allow user to set Filename
func (_options *AddDocumentOptions) SetFilename(filename string) *AddDocumentOptions {
	_options.Filename = core.StringPtr(filename)
	return _options
}

// SetFileContentType : Allow user to set FileContentType
func (_options *AddDocumentOptions) SetFileContentType(fileContentType string) *AddDocumentOptions {
	_options.FileContentType = core.StringPtr(fileContentType)
	return _options
}

// SetMetadata : Allow user to set Metadata
func (_options *AddDocumentOptions) SetMetadata(metadata string) *AddDocumentOptions {
	_options.Metadata = core.StringPtr(metadata)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *AddDocumentOptions) SetHeaders(param map[string]string) *AddDocumentOptions {
	options.Headers = param
	return options
}

// AddTrainingDataOptions : The AddTrainingData options.
type AddTrainingDataOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The ID of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// The natural text query for the new training query.
	NaturalLanguageQuery *string `json:"natural_language_query,omitempty"`

	// The filter used on the collection before the **natural_language_query** is applied.
	Filter *string `json:"filter,omitempty"`

	// Array of training examples.
	Examples []TrainingExample `json:"examples,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAddTrainingDataOptions : Instantiate AddTrainingDataOptions
func (*DiscoveryV1) NewAddTrainingDataOptions(environmentID string, collectionID string) *AddTrainingDataOptions {
	return &AddTrainingDataOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *AddTrainingDataOptions) SetEnvironmentID(environmentID string) *AddTrainingDataOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetCollectionID : Allow user to set CollectionID
func (_options *AddTrainingDataOptions) SetCollectionID(collectionID string) *AddTrainingDataOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetNaturalLanguageQuery : Allow user to set NaturalLanguageQuery
func (_options *AddTrainingDataOptions) SetNaturalLanguageQuery(naturalLanguageQuery string) *AddTrainingDataOptions {
	_options.NaturalLanguageQuery = core.StringPtr(naturalLanguageQuery)
	return _options
}

// SetFilter : Allow user to set Filter
func (_options *AddTrainingDataOptions) SetFilter(filter string) *AddTrainingDataOptions {
	_options.Filter = core.StringPtr(filter)
	return _options
}

// SetExamples : Allow user to set Examples
func (_options *AddTrainingDataOptions) SetExamples(examples []TrainingExample) *AddTrainingDataOptions {
	_options.Examples = examples
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *AddTrainingDataOptions) SetHeaders(param map[string]string) *AddTrainingDataOptions {
	options.Headers = param
	return options
}

// AggregationResult : Aggregation results for the specified query.
type AggregationResult struct {
	// Key that matched the aggregation type.
	Key *string `json:"key,omitempty"`

	// Number of matching results.
	MatchingResults *int64 `json:"matching_results,omitempty"`

	// Aggregations returned in the case of chained aggregations.
	Aggregations []QueryAggregationIntf `json:"aggregations,omitempty"`
}

// UnmarshalAggregationResult unmarshals an instance of AggregationResult from the specified map of raw messages.
func UnmarshalAggregationResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AggregationResult)
	err = core.UnmarshalPrimitive(m, "key", &obj.Key)
	if err != nil {
		err = core.UnmarshalPrimitive(m, "key_as_string", &obj.Key)
		if err != nil {
			return
		}
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

// Collection : A collection for storing documents.
type Collection struct {
	// The unique identifier of the collection.
	CollectionID *string `json:"collection_id,omitempty"`

	// The name of the collection.
	Name *string `json:"name,omitempty"`

	// The description of the collection.
	Description *string `json:"description,omitempty"`

	// The creation date of the collection in the format yyyy-MM-dd'T'HH:mmcon:ss.SSS'Z'.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// The timestamp of when the collection was last updated in the format yyyy-MM-dd'T'HH:mm:ss.SSS'Z'.
	Updated *strfmt.DateTime `json:"updated,omitempty"`

	// The status of the collection.
	Status *string `json:"status,omitempty"`

	// The unique identifier of the collection's configuration.
	ConfigurationID *string `json:"configuration_id,omitempty"`

	// The language of the documents stored in the collection. Permitted values include `en` (English), `de` (German), and
	// `es` (Spanish).
	Language *string `json:"language,omitempty"`

	// Object containing collection document count information.
	DocumentCounts *DocumentCounts `json:"document_counts,omitempty"`

	// Summary of the disk usage statistics for this collection.
	DiskUsage *CollectionDiskUsage `json:"disk_usage,omitempty"`

	// Training status details.
	TrainingStatus *TrainingStatus `json:"training_status,omitempty"`

	// Object containing information about the crawl status of this collection.
	CrawlStatus *CollectionCrawlStatus `json:"crawl_status,omitempty"`

	// Object containing smart document understanding information for this collection.
	SmartDocumentUnderstanding *SduStatus `json:"smart_document_understanding,omitempty"`
}

// Constants associated with the Collection.Status property.
// The status of the collection.
const (
	CollectionStatusActiveConst      = "active"
	CollectionStatusMaintenanceConst = "maintenance"
	CollectionStatusPendingConst     = "pending"
)

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
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "configuration_id", &obj.ConfigurationID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "language", &obj.Language)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "document_counts", &obj.DocumentCounts, UnmarshalDocumentCounts)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "disk_usage", &obj.DiskUsage, UnmarshalCollectionDiskUsage)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "training_status", &obj.TrainingStatus, UnmarshalTrainingStatus)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "crawl_status", &obj.CrawlStatus, UnmarshalCollectionCrawlStatus)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "smart_document_understanding", &obj.SmartDocumentUnderstanding, UnmarshalSduStatus)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CollectionCrawlStatus : Object containing information about the crawl status of this collection.
type CollectionCrawlStatus struct {
	// Object containing source crawl status information.
	SourceCrawl *SourceStatus `json:"source_crawl,omitempty"`
}

// UnmarshalCollectionCrawlStatus unmarshals an instance of CollectionCrawlStatus from the specified map of raw messages.
func UnmarshalCollectionCrawlStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CollectionCrawlStatus)
	err = core.UnmarshalModel(m, "source_crawl", &obj.SourceCrawl, UnmarshalSourceStatus)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CollectionDiskUsage : Summary of the disk usage statistics for this collection.
type CollectionDiskUsage struct {
	// Number of bytes used by the collection.
	UsedBytes *int64 `json:"used_bytes,omitempty"`
}

// UnmarshalCollectionDiskUsage unmarshals an instance of CollectionDiskUsage from the specified map of raw messages.
func UnmarshalCollectionDiskUsage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CollectionDiskUsage)
	err = core.UnmarshalPrimitive(m, "used_bytes", &obj.UsedBytes)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CollectionUsage : Summary of the collection usage in the environment.
type CollectionUsage struct {
	// Number of active collections in the environment.
	Available *int64 `json:"available,omitempty"`

	// Total number of collections allowed in the environment.
	MaximumAllowed *int64 `json:"maximum_allowed,omitempty"`
}

// UnmarshalCollectionUsage unmarshals an instance of CollectionUsage from the specified map of raw messages.
func UnmarshalCollectionUsage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CollectionUsage)
	err = core.UnmarshalPrimitive(m, "available", &obj.Available)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "maximum_allowed", &obj.MaximumAllowed)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Completions : An object containing an array of autocompletion suggestions.
type Completions struct {
	// Array of autcomplete suggestion based on the provided prefix.
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

// Configuration : A custom configuration for the environment.
type Configuration struct {
	// The unique identifier of the configuration.
	ConfigurationID *string `json:"configuration_id,omitempty"`

	// The name of the configuration.
	Name *string `json:"name" validate:"required"`

	// The creation date of the configuration in the format yyyy-MM-dd'T'HH:mm:ss.SSS'Z'.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// The timestamp of when the configuration was last updated in the format yyyy-MM-dd'T'HH:mm:ss.SSS'Z'.
	Updated *strfmt.DateTime `json:"updated,omitempty"`

	// The description of the configuration, if available.
	Description *string `json:"description,omitempty"`

	// Document conversion settings.
	Conversions *Conversions `json:"conversions,omitempty"`

	// An array of document enrichment settings for the configuration.
	Enrichments []Enrichment `json:"enrichments,omitempty"`

	// Defines operations that can be used to transform the final output JSON into a normalized form. Operations are
	// executed in the order that they appear in the array.
	Normalizations []NormalizationOperation `json:"normalizations,omitempty"`

	// Object containing source parameters for the configuration.
	Source *Source `json:"source,omitempty"`
}

// NewConfiguration : Instantiate Configuration (Generic Model Constructor)
func (*DiscoveryV1) NewConfiguration(name string) (_model *Configuration, err error) {
	_model = &Configuration{
		Name: core.StringPtr(name),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalConfiguration unmarshals an instance of Configuration from the specified map of raw messages.
func UnmarshalConfiguration(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Configuration)
	err = core.UnmarshalPrimitive(m, "configuration_id", &obj.ConfigurationID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
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
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "conversions", &obj.Conversions, UnmarshalConversions)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "enrichments", &obj.Enrichments, UnmarshalEnrichment)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "normalizations", &obj.Normalizations, UnmarshalNormalizationOperation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "source", &obj.Source, UnmarshalSource)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Conversions : Document conversion settings.
type Conversions struct {
	// A list of PDF conversion settings.
	PDF *PDFSettings `json:"pdf,omitempty"`

	// A list of Word conversion settings.
	Word *WordSettings `json:"word,omitempty"`

	// A list of HTML conversion settings.
	HTML *HTMLSettings `json:"html,omitempty"`

	// A list of Document Segmentation settings.
	Segment *SegmentSettings `json:"segment,omitempty"`

	// Defines operations that can be used to transform the final output JSON into a normalized form. Operations are
	// executed in the order that they appear in the array.
	JSONNormalizations []NormalizationOperation `json:"json_normalizations,omitempty"`

	// When `true`, automatic text extraction from images (this includes images embedded in supported document formats, for
	// example PDF, and suppported image formats, for example TIFF) is performed on documents uploaded to the collection.
	// This field is supported on **Advanced** and higher plans only. **Lite** plans do not support image text recognition.
	ImageTextRecognition *bool `json:"image_text_recognition,omitempty"`
}

// UnmarshalConversions unmarshals an instance of Conversions from the specified map of raw messages.
func UnmarshalConversions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Conversions)
	err = core.UnmarshalModel(m, "pdf", &obj.PDF, UnmarshalPDFSettings)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "word", &obj.Word, UnmarshalWordSettings)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "html", &obj.HTML, UnmarshalHTMLSettings)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "segment", &obj.Segment, UnmarshalSegmentSettings)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "json_normalizations", &obj.JSONNormalizations, UnmarshalNormalizationOperation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "image_text_recognition", &obj.ImageTextRecognition)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateCollectionOptions : The CreateCollection options.
type CreateCollectionOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The name of the collection to be created.
	Name *string `json:"name" validate:"required"`

	// A description of the collection.
	Description *string `json:"description,omitempty"`

	// The ID of the configuration in which the collection is to be created.
	ConfigurationID *string `json:"configuration_id,omitempty"`

	// The language of the documents stored in the collection, in the form of an ISO 639-1 language code.
	Language *string `json:"language,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateCollectionOptions.Language property.
// The language of the documents stored in the collection, in the form of an ISO 639-1 language code.
const (
	CreateCollectionOptionsLanguageArConst   = "ar"
	CreateCollectionOptionsLanguageDeConst   = "de"
	CreateCollectionOptionsLanguageEnConst   = "en"
	CreateCollectionOptionsLanguageEsConst   = "es"
	CreateCollectionOptionsLanguageFrConst   = "fr"
	CreateCollectionOptionsLanguageItConst   = "it"
	CreateCollectionOptionsLanguageJaConst   = "ja"
	CreateCollectionOptionsLanguageKoConst   = "ko"
	CreateCollectionOptionsLanguageNlConst   = "nl"
	CreateCollectionOptionsLanguagePtConst   = "pt"
	CreateCollectionOptionsLanguageZhCnConst = "zh-CN"
)

// NewCreateCollectionOptions : Instantiate CreateCollectionOptions
func (*DiscoveryV1) NewCreateCollectionOptions(environmentID string, name string) *CreateCollectionOptions {
	return &CreateCollectionOptions{
		EnvironmentID: core.StringPtr(environmentID),
		Name:          core.StringPtr(name),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *CreateCollectionOptions) SetEnvironmentID(environmentID string) *CreateCollectionOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
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

// SetConfigurationID : Allow user to set ConfigurationID
func (_options *CreateCollectionOptions) SetConfigurationID(configurationID string) *CreateCollectionOptions {
	_options.ConfigurationID = core.StringPtr(configurationID)
	return _options
}

// SetLanguage : Allow user to set Language
func (_options *CreateCollectionOptions) SetLanguage(language string) *CreateCollectionOptions {
	_options.Language = core.StringPtr(language)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateCollectionOptions) SetHeaders(param map[string]string) *CreateCollectionOptions {
	options.Headers = param
	return options
}

// CreateConfigurationOptions : The CreateConfiguration options.
type CreateConfigurationOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The name of the configuration.
	Name *string `json:"name" validate:"required"`

	// The description of the configuration, if available.
	Description *string `json:"description,omitempty"`

	// Document conversion settings.
	Conversions *Conversions `json:"conversions,omitempty"`

	// An array of document enrichment settings for the configuration.
	Enrichments []Enrichment `json:"enrichments,omitempty"`

	// Defines operations that can be used to transform the final output JSON into a normalized form. Operations are
	// executed in the order that they appear in the array.
	Normalizations []NormalizationOperation `json:"normalizations,omitempty"`

	// Object containing source parameters for the configuration.
	Source *Source `json:"source,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateConfigurationOptions : Instantiate CreateConfigurationOptions
func (*DiscoveryV1) NewCreateConfigurationOptions(environmentID string, name string) *CreateConfigurationOptions {
	return &CreateConfigurationOptions{
		EnvironmentID: core.StringPtr(environmentID),
		Name:          core.StringPtr(name),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *CreateConfigurationOptions) SetEnvironmentID(environmentID string) *CreateConfigurationOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateConfigurationOptions) SetName(name string) *CreateConfigurationOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateConfigurationOptions) SetDescription(description string) *CreateConfigurationOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetConversions : Allow user to set Conversions
func (_options *CreateConfigurationOptions) SetConversions(conversions *Conversions) *CreateConfigurationOptions {
	_options.Conversions = conversions
	return _options
}

// SetEnrichments : Allow user to set Enrichments
func (_options *CreateConfigurationOptions) SetEnrichments(enrichments []Enrichment) *CreateConfigurationOptions {
	_options.Enrichments = enrichments
	return _options
}

// SetNormalizations : Allow user to set Normalizations
func (_options *CreateConfigurationOptions) SetNormalizations(normalizations []NormalizationOperation) *CreateConfigurationOptions {
	_options.Normalizations = normalizations
	return _options
}

// SetSource : Allow user to set Source
func (_options *CreateConfigurationOptions) SetSource(source *Source) *CreateConfigurationOptions {
	_options.Source = source
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateConfigurationOptions) SetHeaders(param map[string]string) *CreateConfigurationOptions {
	options.Headers = param
	return options
}

// CreateCredentialsOptions : The CreateCredentials options.
type CreateCredentialsOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The source that this credentials object connects to.
	// -  `box` indicates the credentials are used to connect an instance of Enterprise Box.
	// -  `salesforce` indicates the credentials are used to connect to Salesforce.
	// -  `sharepoint` indicates the credentials are used to connect to Microsoft SharePoint Online.
	// -  `web_crawl` indicates the credentials are used to perform a web crawl.
	// =  `cloud_object_storage` indicates the credentials are used to connect to an IBM Cloud Object Store.
	SourceType *string `json:"source_type,omitempty"`

	// Object containing details of the stored credentials.
	//
	// Obtain credentials for your source from the administrator of the source.
	CredentialDetails *CredentialDetails `json:"credential_details,omitempty"`

	// Object that contains details about the status of the authentication process.
	Status *StatusDetails `json:"status,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateCredentialsOptions.SourceType property.
// The source that this credentials object connects to.
// -  `box` indicates the credentials are used to connect an instance of Enterprise Box.
// -  `salesforce` indicates the credentials are used to connect to Salesforce.
// -  `sharepoint` indicates the credentials are used to connect to Microsoft SharePoint Online.
// -  `web_crawl` indicates the credentials are used to perform a web crawl.
// =  `cloud_object_storage` indicates the credentials are used to connect to an IBM Cloud Object Store.
const (
	CreateCredentialsOptionsSourceTypeBoxConst                = "box"
	CreateCredentialsOptionsSourceTypeCloudObjectStorageConst = "cloud_object_storage"
	CreateCredentialsOptionsSourceTypeSalesforceConst         = "salesforce"
	CreateCredentialsOptionsSourceTypeSharepointConst         = "sharepoint"
	CreateCredentialsOptionsSourceTypeWebCrawlConst           = "web_crawl"
)

// NewCreateCredentialsOptions : Instantiate CreateCredentialsOptions
func (*DiscoveryV1) NewCreateCredentialsOptions(environmentID string) *CreateCredentialsOptions {
	return &CreateCredentialsOptions{
		EnvironmentID: core.StringPtr(environmentID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *CreateCredentialsOptions) SetEnvironmentID(environmentID string) *CreateCredentialsOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetSourceType : Allow user to set SourceType
func (_options *CreateCredentialsOptions) SetSourceType(sourceType string) *CreateCredentialsOptions {
	_options.SourceType = core.StringPtr(sourceType)
	return _options
}

// SetCredentialDetails : Allow user to set CredentialDetails
func (_options *CreateCredentialsOptions) SetCredentialDetails(credentialDetails *CredentialDetails) *CreateCredentialsOptions {
	_options.CredentialDetails = credentialDetails
	return _options
}

// SetStatus : Allow user to set Status
func (_options *CreateCredentialsOptions) SetStatus(status *StatusDetails) *CreateCredentialsOptions {
	_options.Status = status
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateCredentialsOptions) SetHeaders(param map[string]string) *CreateCredentialsOptions {
	options.Headers = param
	return options
}

// CreateEnvironmentOptions : The CreateEnvironment options.
type CreateEnvironmentOptions struct {
	// Name that identifies the environment.
	Name *string `json:"name" validate:"required"`

	// Description of the environment.
	Description *string `json:"description,omitempty"`

	// Size of the environment. In the Lite plan the default and only accepted value is `LT`, in all other plans the
	// default is `S`.
	Size *string `json:"size,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateEnvironmentOptions.Size property.
// Size of the environment. In the Lite plan the default and only accepted value is `LT`, in all other plans the default
// is `S`.
const (
	CreateEnvironmentOptionsSizeLConst    = "L"
	CreateEnvironmentOptionsSizeLtConst   = "LT"
	CreateEnvironmentOptionsSizeMConst    = "M"
	CreateEnvironmentOptionsSizeMlConst   = "ML"
	CreateEnvironmentOptionsSizeMsConst   = "MS"
	CreateEnvironmentOptionsSizeSConst    = "S"
	CreateEnvironmentOptionsSizeXlConst   = "XL"
	CreateEnvironmentOptionsSizeXsConst   = "XS"
	CreateEnvironmentOptionsSizeXxlConst  = "XXL"
	CreateEnvironmentOptionsSizeXxxlConst = "XXXL"
)

// NewCreateEnvironmentOptions : Instantiate CreateEnvironmentOptions
func (*DiscoveryV1) NewCreateEnvironmentOptions(name string) *CreateEnvironmentOptions {
	return &CreateEnvironmentOptions{
		Name: core.StringPtr(name),
	}
}

// SetName : Allow user to set Name
func (_options *CreateEnvironmentOptions) SetName(name string) *CreateEnvironmentOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateEnvironmentOptions) SetDescription(description string) *CreateEnvironmentOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetSize : Allow user to set Size
func (_options *CreateEnvironmentOptions) SetSize(size string) *CreateEnvironmentOptions {
	_options.Size = core.StringPtr(size)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateEnvironmentOptions) SetHeaders(param map[string]string) *CreateEnvironmentOptions {
	options.Headers = param
	return options
}

// CreateEventOptions : The CreateEvent options.
type CreateEventOptions struct {
	// The event type to be created.
	Type *string `json:"type" validate:"required"`

	// Query event data object.
	Data *EventData `json:"data" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateEventOptions.Type property.
// The event type to be created.
const (
	CreateEventOptionsTypeClickConst = "click"
)

// NewCreateEventOptions : Instantiate CreateEventOptions
func (*DiscoveryV1) NewCreateEventOptions(typeVar string, data *EventData) *CreateEventOptions {
	return &CreateEventOptions{
		Type: core.StringPtr(typeVar),
		Data: data,
	}
}

// SetType : Allow user to set Type
func (_options *CreateEventOptions) SetType(typeVar string) *CreateEventOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetData : Allow user to set Data
func (_options *CreateEventOptions) SetData(data *EventData) *CreateEventOptions {
	_options.Data = data
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateEventOptions) SetHeaders(param map[string]string) *CreateEventOptions {
	options.Headers = param
	return options
}

// CreateEventResponse : An object defining the event being created.
type CreateEventResponse struct {
	// The event type that was created.
	Type *string `json:"type,omitempty"`

	// Query event data object.
	Data *EventData `json:"data,omitempty"`
}

// Constants associated with the CreateEventResponse.Type property.
// The event type that was created.
const (
	CreateEventResponseTypeClickConst = "click"
)

// UnmarshalCreateEventResponse unmarshals an instance of CreateEventResponse from the specified map of raw messages.
func UnmarshalCreateEventResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateEventResponse)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "data", &obj.Data, UnmarshalEventData)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateExpansionsOptions : The CreateExpansions options.
type CreateExpansionsOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The ID of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// An array of query expansion definitions.
	//
	//  Each object in the **expansions** array represents a term or set of terms that will be expanded into other terms.
	// Each expansion object can be configured as bidirectional or unidirectional. Bidirectional means that all terms are
	// expanded to all other terms in the object. Unidirectional means that a set list of terms can be expanded into a
	// second list of terms.
	//
	//  To create a bi-directional expansion specify an **expanded_terms** array. When found in a query, all items in the
	// **expanded_terms** array are then expanded to the other items in the same array.
	//
	//  To create a uni-directional expansion, specify both an array of **input_terms** and an array of **expanded_terms**.
	// When items in the **input_terms** array are present in a query, they are expanded using the items listed in the
	// **expanded_terms** array.
	Expansions []Expansion `json:"expansions" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateExpansionsOptions : Instantiate CreateExpansionsOptions
func (*DiscoveryV1) NewCreateExpansionsOptions(environmentID string, collectionID string, expansions []Expansion) *CreateExpansionsOptions {
	return &CreateExpansionsOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
		Expansions:    expansions,
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *CreateExpansionsOptions) SetEnvironmentID(environmentID string) *CreateExpansionsOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetCollectionID : Allow user to set CollectionID
func (_options *CreateExpansionsOptions) SetCollectionID(collectionID string) *CreateExpansionsOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetExpansions : Allow user to set Expansions
func (_options *CreateExpansionsOptions) SetExpansions(expansions []Expansion) *CreateExpansionsOptions {
	_options.Expansions = expansions
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateExpansionsOptions) SetHeaders(param map[string]string) *CreateExpansionsOptions {
	options.Headers = param
	return options
}

// CreateGatewayOptions : The CreateGateway options.
type CreateGatewayOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// User-defined name.
	Name *string `json:"name,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateGatewayOptions : Instantiate CreateGatewayOptions
func (*DiscoveryV1) NewCreateGatewayOptions(environmentID string) *CreateGatewayOptions {
	return &CreateGatewayOptions{
		EnvironmentID: core.StringPtr(environmentID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *CreateGatewayOptions) SetEnvironmentID(environmentID string) *CreateGatewayOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateGatewayOptions) SetName(name string) *CreateGatewayOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateGatewayOptions) SetHeaders(param map[string]string) *CreateGatewayOptions {
	options.Headers = param
	return options
}

// CreateStopwordListOptions : The CreateStopwordList options.
type CreateStopwordListOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The ID of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// The content of the stopword list to ingest.
	StopwordFile io.ReadCloser `json:"-" validate:"required"`

	// The filename for stopwordFile.
	StopwordFilename *string `json:"-" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateStopwordListOptions : Instantiate CreateStopwordListOptions
func (*DiscoveryV1) NewCreateStopwordListOptions(environmentID string, collectionID string, stopwordFile io.ReadCloser, stopwordFilename string) *CreateStopwordListOptions {
	return &CreateStopwordListOptions{
		EnvironmentID:    core.StringPtr(environmentID),
		CollectionID:     core.StringPtr(collectionID),
		StopwordFile:     stopwordFile,
		StopwordFilename: core.StringPtr(stopwordFilename),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *CreateStopwordListOptions) SetEnvironmentID(environmentID string) *CreateStopwordListOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetCollectionID : Allow user to set CollectionID
func (_options *CreateStopwordListOptions) SetCollectionID(collectionID string) *CreateStopwordListOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetStopwordFile : Allow user to set StopwordFile
func (_options *CreateStopwordListOptions) SetStopwordFile(stopwordFile io.ReadCloser) *CreateStopwordListOptions {
	_options.StopwordFile = stopwordFile
	return _options
}

// SetStopwordFilename : Allow user to set StopwordFilename
func (_options *CreateStopwordListOptions) SetStopwordFilename(stopwordFilename string) *CreateStopwordListOptions {
	_options.StopwordFilename = core.StringPtr(stopwordFilename)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateStopwordListOptions) SetHeaders(param map[string]string) *CreateStopwordListOptions {
	options.Headers = param
	return options
}

// CreateTokenizationDictionaryOptions : The CreateTokenizationDictionary options.
type CreateTokenizationDictionaryOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The ID of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// An array of tokenization rules. Each rule contains, the original `text` string, component `tokens`, any alternate
	// character set `readings`, and which `part_of_speech` the text is from.
	TokenizationRules []TokenDictRule `json:"tokenization_rules,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateTokenizationDictionaryOptions : Instantiate CreateTokenizationDictionaryOptions
func (*DiscoveryV1) NewCreateTokenizationDictionaryOptions(environmentID string, collectionID string) *CreateTokenizationDictionaryOptions {
	return &CreateTokenizationDictionaryOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *CreateTokenizationDictionaryOptions) SetEnvironmentID(environmentID string) *CreateTokenizationDictionaryOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetCollectionID : Allow user to set CollectionID
func (_options *CreateTokenizationDictionaryOptions) SetCollectionID(collectionID string) *CreateTokenizationDictionaryOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetTokenizationRules : Allow user to set TokenizationRules
func (_options *CreateTokenizationDictionaryOptions) SetTokenizationRules(tokenizationRules []TokenDictRule) *CreateTokenizationDictionaryOptions {
	_options.TokenizationRules = tokenizationRules
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateTokenizationDictionaryOptions) SetHeaders(param map[string]string) *CreateTokenizationDictionaryOptions {
	options.Headers = param
	return options
}

// CreateTrainingExampleOptions : The CreateTrainingExample options.
type CreateTrainingExampleOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The ID of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// The ID of the query used for training.
	QueryID *string `json:"-" validate:"required,ne="`

	// The document ID associated with this training example.
	DocumentID *string `json:"document_id,omitempty"`

	// The cross reference associated with this training example.
	CrossReference *string `json:"cross_reference,omitempty"`

	// The relevance of the training example.
	Relevance *int64 `json:"relevance,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateTrainingExampleOptions : Instantiate CreateTrainingExampleOptions
func (*DiscoveryV1) NewCreateTrainingExampleOptions(environmentID string, collectionID string, queryID string) *CreateTrainingExampleOptions {
	return &CreateTrainingExampleOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
		QueryID:       core.StringPtr(queryID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *CreateTrainingExampleOptions) SetEnvironmentID(environmentID string) *CreateTrainingExampleOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetCollectionID : Allow user to set CollectionID
func (_options *CreateTrainingExampleOptions) SetCollectionID(collectionID string) *CreateTrainingExampleOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetQueryID : Allow user to set QueryID
func (_options *CreateTrainingExampleOptions) SetQueryID(queryID string) *CreateTrainingExampleOptions {
	_options.QueryID = core.StringPtr(queryID)
	return _options
}

// SetDocumentID : Allow user to set DocumentID
func (_options *CreateTrainingExampleOptions) SetDocumentID(documentID string) *CreateTrainingExampleOptions {
	_options.DocumentID = core.StringPtr(documentID)
	return _options
}

// SetCrossReference : Allow user to set CrossReference
func (_options *CreateTrainingExampleOptions) SetCrossReference(crossReference string) *CreateTrainingExampleOptions {
	_options.CrossReference = core.StringPtr(crossReference)
	return _options
}

// SetRelevance : Allow user to set Relevance
func (_options *CreateTrainingExampleOptions) SetRelevance(relevance int64) *CreateTrainingExampleOptions {
	_options.Relevance = core.Int64Ptr(relevance)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateTrainingExampleOptions) SetHeaders(param map[string]string) *CreateTrainingExampleOptions {
	options.Headers = param
	return options
}

// CredentialDetails : Object containing details of the stored credentials.
//
// Obtain credentials for your source from the administrator of the source.
type CredentialDetails struct {
	// The authentication method for this credentials definition. The  **credential_type** specified must be supported by
	// the **source_type**. The following combinations are possible:
	//
	// -  `"source_type": "box"` - valid `credential_type`s: `oauth2`
	// -  `"source_type": "salesforce"` - valid `credential_type`s: `username_password`
	// -  `"source_type": "sharepoint"` - valid `credential_type`s: `saml` with **source_version** of `online`, or
	// `ntlm_v1` with **source_version** of `2016`
	// -  `"source_type": "web_crawl"` - valid `credential_type`s: `noauth` or `basic`
	// -  "source_type": "cloud_object_storage"` - valid `credential_type`s: `aws4_hmac`.
	CredentialType *string `json:"credential_type,omitempty"`

	// The **client_id** of the source that these credentials connect to. Only valid, and required, with a
	// **credential_type** of `oauth2`.
	ClientID *string `json:"client_id,omitempty"`

	// The **enterprise_id** of the Box site that these credentials connect to. Only valid, and required, with a
	// **source_type** of `box`.
	EnterpriseID *string `json:"enterprise_id,omitempty"`

	// The **url** of the source that these credentials connect to. Only valid, and required, with a **credential_type** of
	// `username_password`, `noauth`, and `basic`.
	URL *string `json:"url,omitempty"`

	// The **username** of the source that these credentials connect to. Only valid, and required, with a
	// **credential_type** of `saml`, `username_password`, `basic`, or `ntlm_v1`.
	Username *string `json:"username,omitempty"`

	// The **organization_url** of the source that these credentials connect to. Only valid, and required, with a
	// **credential_type** of `saml`.
	OrganizationURL *string `json:"organization_url,omitempty"`

	// The **site_collection.path** of the source that these credentials connect to. Only valid, and required, with a
	// **source_type** of `sharepoint`.
	SiteCollectionPath *string `json:"site_collection.path,omitempty"`

	// The **client_secret** of the source that these credentials connect to. Only valid, and required, with a
	// **credential_type** of `oauth2`. This value is never returned and is only used when creating or modifying
	// **credentials**.
	ClientSecret *string `json:"client_secret,omitempty"`

	// The **public_key_id** of the source that these credentials connect to. Only valid, and required, with a
	// **credential_type** of `oauth2`. This value is never returned and is only used when creating or modifying
	// **credentials**.
	PublicKeyID *string `json:"public_key_id,omitempty"`

	// The **private_key** of the source that these credentials connect to. Only valid, and required, with a
	// **credential_type** of `oauth2`. This value is never returned and is only used when creating or modifying
	// **credentials**.
	PrivateKey *string `json:"private_key,omitempty"`

	// The **passphrase** of the source that these credentials connect to. Only valid, and required, with a
	// **credential_type** of `oauth2`. This value is never returned and is only used when creating or modifying
	// **credentials**.
	Passphrase *string `json:"passphrase,omitempty"`

	// The **password** of the source that these credentials connect to. Only valid, and required, with
	// **credential_type**s of `saml`, `username_password`, `basic`, or `ntlm_v1`.
	//
	// **Note:** When used with a **source_type** of `salesforce`, the password consists of the Salesforce password and a
	// valid Salesforce security token concatenated. This value is never returned and is only used when creating or
	// modifying **credentials**.
	Password *string `json:"password,omitempty"`

	// The ID of the **gateway** to be connected through (when connecting to intranet sites). Only valid with a
	// **credential_type** of `noauth`, `basic`, or `ntlm_v1`. Gateways are created using the
	// `/v1/environments/{environment_id}/gateways` methods.
	GatewayID *string `json:"gateway_id,omitempty"`

	// The type of Sharepoint repository to connect to. Only valid, and required, with a **source_type** of `sharepoint`.
	SourceVersion *string `json:"source_version,omitempty"`

	// SharePoint OnPrem WebApplication URL. Only valid, and required, with a **source_version** of `2016`. If a port is
	// not supplied, the default to port `80` for http and port `443` for https connections are used.
	WebApplicationURL *string `json:"web_application_url,omitempty"`

	// The domain used to log in to your OnPrem SharePoint account. Only valid, and required, with a **source_version** of
	// `2016`.
	Domain *string `json:"domain,omitempty"`

	// The endpoint associated with the cloud object store that your are connecting to. Only valid, and required, with a
	// **credential_type** of `aws4_hmac`.
	Endpoint *string `json:"endpoint,omitempty"`

	// The access key ID associated with the cloud object store. Only valid, and required, with a **credential_type** of
	// `aws4_hmac`. This value is never returned and is only used when creating or modifying **credentials**. For more
	// infomation, see the [cloud object store
	// documentation](https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-using-hmac-credentials#using-hmac-credentials).
	AccessKeyID *string `json:"access_key_id,omitempty"`

	// The secret access key associated with the cloud object store. Only valid, and required, with a **credential_type**
	// of `aws4_hmac`. This value is never returned and is only used when creating or modifying **credentials**. For more
	// infomation, see the [cloud object store
	// documentation](https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-using-hmac-credentials#using-hmac-credentials).
	SecretAccessKey *string `json:"secret_access_key,omitempty"`
}

// Constants associated with the CredentialDetails.CredentialType property.
// The authentication method for this credentials definition. The  **credential_type** specified must be supported by
// the **source_type**. The following combinations are possible:
//
// -  `"source_type": "box"` - valid `credential_type`s: `oauth2`
// -  `"source_type": "salesforce"` - valid `credential_type`s: `username_password`
// -  `"source_type": "sharepoint"` - valid `credential_type`s: `saml` with **source_version** of `online`, or `ntlm_v1`
// with **source_version** of `2016`
// -  `"source_type": "web_crawl"` - valid `credential_type`s: `noauth` or `basic`
// -  "source_type": "cloud_object_storage"` - valid `credential_type`s: `aws4_hmac`.
const (
	CredentialDetailsCredentialTypeAws4HmacConst         = "aws4_hmac"
	CredentialDetailsCredentialTypeBasicConst            = "basic"
	CredentialDetailsCredentialTypeNoauthConst           = "noauth"
	CredentialDetailsCredentialTypeNtlmV1Const           = "ntlm_v1"
	CredentialDetailsCredentialTypeOauth2Const           = "oauth2"
	CredentialDetailsCredentialTypeSamlConst             = "saml"
	CredentialDetailsCredentialTypeUsernamePasswordConst = "username_password"
)

// Constants associated with the CredentialDetails.SourceVersion property.
// The type of Sharepoint repository to connect to. Only valid, and required, with a **source_type** of `sharepoint`.
const (
	CredentialDetailsSourceVersionOnlineConst = "online"
)

// UnmarshalCredentialDetails unmarshals an instance of CredentialDetails from the specified map of raw messages.
func UnmarshalCredentialDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CredentialDetails)
	err = core.UnmarshalPrimitive(m, "credential_type", &obj.CredentialType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "client_id", &obj.ClientID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enterprise_id", &obj.EnterpriseID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "username", &obj.Username)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "organization_url", &obj.OrganizationURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "site_collection.path", &obj.SiteCollectionPath)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "client_secret", &obj.ClientSecret)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "public_key_id", &obj.PublicKeyID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "private_key", &obj.PrivateKey)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "passphrase", &obj.Passphrase)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "gateway_id", &obj.GatewayID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source_version", &obj.SourceVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "web_application_url", &obj.WebApplicationURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "domain", &obj.Domain)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "endpoint", &obj.Endpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "access_key_id", &obj.AccessKeyID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "secret_access_key", &obj.SecretAccessKey)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Credentials : Object containing credential information.
type Credentials struct {
	// Unique identifier for this set of credentials.
	CredentialID *string `json:"credential_id,omitempty"`

	// The source that this credentials object connects to.
	// -  `box` indicates the credentials are used to connect an instance of Enterprise Box.
	// -  `salesforce` indicates the credentials are used to connect to Salesforce.
	// -  `sharepoint` indicates the credentials are used to connect to Microsoft SharePoint Online.
	// -  `web_crawl` indicates the credentials are used to perform a web crawl.
	// =  `cloud_object_storage` indicates the credentials are used to connect to an IBM Cloud Object Store.
	SourceType *string `json:"source_type,omitempty"`

	// Object containing details of the stored credentials.
	//
	// Obtain credentials for your source from the administrator of the source.
	CredentialDetails *CredentialDetails `json:"credential_details,omitempty"`

	// Object that contains details about the status of the authentication process.
	Status *StatusDetails `json:"status,omitempty"`
}

// Constants associated with the Credentials.SourceType property.
// The source that this credentials object connects to.
// -  `box` indicates the credentials are used to connect an instance of Enterprise Box.
// -  `salesforce` indicates the credentials are used to connect to Salesforce.
// -  `sharepoint` indicates the credentials are used to connect to Microsoft SharePoint Online.
// -  `web_crawl` indicates the credentials are used to perform a web crawl.
// =  `cloud_object_storage` indicates the credentials are used to connect to an IBM Cloud Object Store.
const (
	CredentialsSourceTypeBoxConst                = "box"
	CredentialsSourceTypeCloudObjectStorageConst = "cloud_object_storage"
	CredentialsSourceTypeSalesforceConst         = "salesforce"
	CredentialsSourceTypeSharepointConst         = "sharepoint"
	CredentialsSourceTypeWebCrawlConst           = "web_crawl"
)

// UnmarshalCredentials unmarshals an instance of Credentials from the specified map of raw messages.
func UnmarshalCredentials(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Credentials)
	err = core.UnmarshalPrimitive(m, "credential_id", &obj.CredentialID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source_type", &obj.SourceType)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "credential_details", &obj.CredentialDetails, UnmarshalCredentialDetails)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "status", &obj.Status, UnmarshalStatusDetails)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CredentialsList : Object containing array of credential definitions.
type CredentialsList struct {
	// An array of credential definitions that were created for this instance.
	Credentials []Credentials `json:"credentials,omitempty"`
}

// UnmarshalCredentialsList unmarshals an instance of CredentialsList from the specified map of raw messages.
func UnmarshalCredentialsList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CredentialsList)
	err = core.UnmarshalModel(m, "credentials", &obj.Credentials, UnmarshalCredentials)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteAllTrainingDataOptions : The DeleteAllTrainingData options.
type DeleteAllTrainingDataOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The ID of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteAllTrainingDataOptions : Instantiate DeleteAllTrainingDataOptions
func (*DiscoveryV1) NewDeleteAllTrainingDataOptions(environmentID string, collectionID string) *DeleteAllTrainingDataOptions {
	return &DeleteAllTrainingDataOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *DeleteAllTrainingDataOptions) SetEnvironmentID(environmentID string) *DeleteAllTrainingDataOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetCollectionID : Allow user to set CollectionID
func (_options *DeleteAllTrainingDataOptions) SetCollectionID(collectionID string) *DeleteAllTrainingDataOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteAllTrainingDataOptions) SetHeaders(param map[string]string) *DeleteAllTrainingDataOptions {
	options.Headers = param
	return options
}

// DeleteCollectionOptions : The DeleteCollection options.
type DeleteCollectionOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The ID of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteCollectionOptions : Instantiate DeleteCollectionOptions
func (*DiscoveryV1) NewDeleteCollectionOptions(environmentID string, collectionID string) *DeleteCollectionOptions {
	return &DeleteCollectionOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *DeleteCollectionOptions) SetEnvironmentID(environmentID string) *DeleteCollectionOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
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

// DeleteCollectionResponse : Response object returned when deleting a colleciton.
type DeleteCollectionResponse struct {
	// The unique identifier of the collection that is being deleted.
	CollectionID *string `json:"collection_id" validate:"required"`

	// The status of the collection. The status of a successful deletion operation is `deleted`.
	Status *string `json:"status" validate:"required"`
}

// Constants associated with the DeleteCollectionResponse.Status property.
// The status of the collection. The status of a successful deletion operation is `deleted`.
const (
	DeleteCollectionResponseStatusDeletedConst = "deleted"
)

// UnmarshalDeleteCollectionResponse unmarshals an instance of DeleteCollectionResponse from the specified map of raw messages.
func UnmarshalDeleteCollectionResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteCollectionResponse)
	err = core.UnmarshalPrimitive(m, "collection_id", &obj.CollectionID)
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

// DeleteConfigurationOptions : The DeleteConfiguration options.
type DeleteConfigurationOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The ID of the configuration.
	ConfigurationID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteConfigurationOptions : Instantiate DeleteConfigurationOptions
func (*DiscoveryV1) NewDeleteConfigurationOptions(environmentID string, configurationID string) *DeleteConfigurationOptions {
	return &DeleteConfigurationOptions{
		EnvironmentID:   core.StringPtr(environmentID),
		ConfigurationID: core.StringPtr(configurationID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *DeleteConfigurationOptions) SetEnvironmentID(environmentID string) *DeleteConfigurationOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetConfigurationID : Allow user to set ConfigurationID
func (_options *DeleteConfigurationOptions) SetConfigurationID(configurationID string) *DeleteConfigurationOptions {
	_options.ConfigurationID = core.StringPtr(configurationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteConfigurationOptions) SetHeaders(param map[string]string) *DeleteConfigurationOptions {
	options.Headers = param
	return options
}

// DeleteConfigurationResponse : Information returned when a configuration is deleted.
type DeleteConfigurationResponse struct {
	// The unique identifier for the configuration.
	ConfigurationID *string `json:"configuration_id" validate:"required"`

	// Status of the configuration. A deleted configuration has the status deleted.
	Status *string `json:"status" validate:"required"`

	// An array of notice messages, if any.
	Notices []Notice `json:"notices,omitempty"`
}

// Constants associated with the DeleteConfigurationResponse.Status property.
// Status of the configuration. A deleted configuration has the status deleted.
const (
	DeleteConfigurationResponseStatusDeletedConst = "deleted"
)

// UnmarshalDeleteConfigurationResponse unmarshals an instance of DeleteConfigurationResponse from the specified map of raw messages.
func UnmarshalDeleteConfigurationResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteConfigurationResponse)
	err = core.UnmarshalPrimitive(m, "configuration_id", &obj.ConfigurationID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
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

// DeleteCredentials : Object returned after credentials are deleted.
type DeleteCredentials struct {
	// The unique identifier of the credentials that have been deleted.
	CredentialID *string `json:"credential_id,omitempty"`

	// The status of the deletion request.
	Status *string `json:"status,omitempty"`
}

// Constants associated with the DeleteCredentials.Status property.
// The status of the deletion request.
const (
	DeleteCredentialsStatusDeletedConst = "deleted"
)

// UnmarshalDeleteCredentials unmarshals an instance of DeleteCredentials from the specified map of raw messages.
func UnmarshalDeleteCredentials(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteCredentials)
	err = core.UnmarshalPrimitive(m, "credential_id", &obj.CredentialID)
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

// DeleteCredentialsOptions : The DeleteCredentials options.
type DeleteCredentialsOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The unique identifier for a set of source credentials.
	CredentialID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteCredentialsOptions : Instantiate DeleteCredentialsOptions
func (*DiscoveryV1) NewDeleteCredentialsOptions(environmentID string, credentialID string) *DeleteCredentialsOptions {
	return &DeleteCredentialsOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CredentialID:  core.StringPtr(credentialID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *DeleteCredentialsOptions) SetEnvironmentID(environmentID string) *DeleteCredentialsOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetCredentialID : Allow user to set CredentialID
func (_options *DeleteCredentialsOptions) SetCredentialID(credentialID string) *DeleteCredentialsOptions {
	_options.CredentialID = core.StringPtr(credentialID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteCredentialsOptions) SetHeaders(param map[string]string) *DeleteCredentialsOptions {
	options.Headers = param
	return options
}

// DeleteDocumentOptions : The DeleteDocument options.
type DeleteDocumentOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The ID of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// The ID of the document.
	DocumentID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteDocumentOptions : Instantiate DeleteDocumentOptions
func (*DiscoveryV1) NewDeleteDocumentOptions(environmentID string, collectionID string, documentID string) *DeleteDocumentOptions {
	return &DeleteDocumentOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
		DocumentID:    core.StringPtr(documentID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *DeleteDocumentOptions) SetEnvironmentID(environmentID string) *DeleteDocumentOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetCollectionID : Allow user to set CollectionID
func (_options *DeleteDocumentOptions) SetCollectionID(collectionID string) *DeleteDocumentOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetDocumentID : Allow user to set DocumentID
func (_options *DeleteDocumentOptions) SetDocumentID(documentID string) *DeleteDocumentOptions {
	_options.DocumentID = core.StringPtr(documentID)
	return _options
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

// DeleteEnvironmentOptions : The DeleteEnvironment options.
type DeleteEnvironmentOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteEnvironmentOptions : Instantiate DeleteEnvironmentOptions
func (*DiscoveryV1) NewDeleteEnvironmentOptions(environmentID string) *DeleteEnvironmentOptions {
	return &DeleteEnvironmentOptions{
		EnvironmentID: core.StringPtr(environmentID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *DeleteEnvironmentOptions) SetEnvironmentID(environmentID string) *DeleteEnvironmentOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteEnvironmentOptions) SetHeaders(param map[string]string) *DeleteEnvironmentOptions {
	options.Headers = param
	return options
}

// DeleteEnvironmentResponse : Response object returned when deleting an environment.
type DeleteEnvironmentResponse struct {
	// The unique identifier for the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// Status of the environment.
	Status *string `json:"status" validate:"required"`
}

// Constants associated with the DeleteEnvironmentResponse.Status property.
// Status of the environment.
const (
	DeleteEnvironmentResponseStatusDeletedConst = "deleted"
)

// UnmarshalDeleteEnvironmentResponse unmarshals an instance of DeleteEnvironmentResponse from the specified map of raw messages.
func UnmarshalDeleteEnvironmentResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteEnvironmentResponse)
	err = core.UnmarshalPrimitive(m, "environment_id", &obj.EnvironmentID)
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

// DeleteExpansionsOptions : The DeleteExpansions options.
type DeleteExpansionsOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The ID of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteExpansionsOptions : Instantiate DeleteExpansionsOptions
func (*DiscoveryV1) NewDeleteExpansionsOptions(environmentID string, collectionID string) *DeleteExpansionsOptions {
	return &DeleteExpansionsOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *DeleteExpansionsOptions) SetEnvironmentID(environmentID string) *DeleteExpansionsOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetCollectionID : Allow user to set CollectionID
func (_options *DeleteExpansionsOptions) SetCollectionID(collectionID string) *DeleteExpansionsOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteExpansionsOptions) SetHeaders(param map[string]string) *DeleteExpansionsOptions {
	options.Headers = param
	return options
}

// DeleteGatewayOptions : The DeleteGateway options.
type DeleteGatewayOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The requested gateway ID.
	GatewayID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteGatewayOptions : Instantiate DeleteGatewayOptions
func (*DiscoveryV1) NewDeleteGatewayOptions(environmentID string, gatewayID string) *DeleteGatewayOptions {
	return &DeleteGatewayOptions{
		EnvironmentID: core.StringPtr(environmentID),
		GatewayID:     core.StringPtr(gatewayID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *DeleteGatewayOptions) SetEnvironmentID(environmentID string) *DeleteGatewayOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetGatewayID : Allow user to set GatewayID
func (_options *DeleteGatewayOptions) SetGatewayID(gatewayID string) *DeleteGatewayOptions {
	_options.GatewayID = core.StringPtr(gatewayID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteGatewayOptions) SetHeaders(param map[string]string) *DeleteGatewayOptions {
	options.Headers = param
	return options
}

// DeleteStopwordListOptions : The DeleteStopwordList options.
type DeleteStopwordListOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The ID of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteStopwordListOptions : Instantiate DeleteStopwordListOptions
func (*DiscoveryV1) NewDeleteStopwordListOptions(environmentID string, collectionID string) *DeleteStopwordListOptions {
	return &DeleteStopwordListOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *DeleteStopwordListOptions) SetEnvironmentID(environmentID string) *DeleteStopwordListOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetCollectionID : Allow user to set CollectionID
func (_options *DeleteStopwordListOptions) SetCollectionID(collectionID string) *DeleteStopwordListOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteStopwordListOptions) SetHeaders(param map[string]string) *DeleteStopwordListOptions {
	options.Headers = param
	return options
}

// DeleteTokenizationDictionaryOptions : The DeleteTokenizationDictionary options.
type DeleteTokenizationDictionaryOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The ID of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteTokenizationDictionaryOptions : Instantiate DeleteTokenizationDictionaryOptions
func (*DiscoveryV1) NewDeleteTokenizationDictionaryOptions(environmentID string, collectionID string) *DeleteTokenizationDictionaryOptions {
	return &DeleteTokenizationDictionaryOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *DeleteTokenizationDictionaryOptions) SetEnvironmentID(environmentID string) *DeleteTokenizationDictionaryOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetCollectionID : Allow user to set CollectionID
func (_options *DeleteTokenizationDictionaryOptions) SetCollectionID(collectionID string) *DeleteTokenizationDictionaryOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteTokenizationDictionaryOptions) SetHeaders(param map[string]string) *DeleteTokenizationDictionaryOptions {
	options.Headers = param
	return options
}

// DeleteTrainingDataOptions : The DeleteTrainingData options.
type DeleteTrainingDataOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The ID of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// The ID of the query used for training.
	QueryID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteTrainingDataOptions : Instantiate DeleteTrainingDataOptions
func (*DiscoveryV1) NewDeleteTrainingDataOptions(environmentID string, collectionID string, queryID string) *DeleteTrainingDataOptions {
	return &DeleteTrainingDataOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
		QueryID:       core.StringPtr(queryID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *DeleteTrainingDataOptions) SetEnvironmentID(environmentID string) *DeleteTrainingDataOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetCollectionID : Allow user to set CollectionID
func (_options *DeleteTrainingDataOptions) SetCollectionID(collectionID string) *DeleteTrainingDataOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetQueryID : Allow user to set QueryID
func (_options *DeleteTrainingDataOptions) SetQueryID(queryID string) *DeleteTrainingDataOptions {
	_options.QueryID = core.StringPtr(queryID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteTrainingDataOptions) SetHeaders(param map[string]string) *DeleteTrainingDataOptions {
	options.Headers = param
	return options
}

// DeleteTrainingExampleOptions : The DeleteTrainingExample options.
type DeleteTrainingExampleOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The ID of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// The ID of the query used for training.
	QueryID *string `json:"-" validate:"required,ne="`

	// The ID of the document as it is indexed.
	ExampleID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteTrainingExampleOptions : Instantiate DeleteTrainingExampleOptions
func (*DiscoveryV1) NewDeleteTrainingExampleOptions(environmentID string, collectionID string, queryID string, exampleID string) *DeleteTrainingExampleOptions {
	return &DeleteTrainingExampleOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
		QueryID:       core.StringPtr(queryID),
		ExampleID:     core.StringPtr(exampleID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *DeleteTrainingExampleOptions) SetEnvironmentID(environmentID string) *DeleteTrainingExampleOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetCollectionID : Allow user to set CollectionID
func (_options *DeleteTrainingExampleOptions) SetCollectionID(collectionID string) *DeleteTrainingExampleOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetQueryID : Allow user to set QueryID
func (_options *DeleteTrainingExampleOptions) SetQueryID(queryID string) *DeleteTrainingExampleOptions {
	_options.QueryID = core.StringPtr(queryID)
	return _options
}

// SetExampleID : Allow user to set ExampleID
func (_options *DeleteTrainingExampleOptions) SetExampleID(exampleID string) *DeleteTrainingExampleOptions {
	_options.ExampleID = core.StringPtr(exampleID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteTrainingExampleOptions) SetHeaders(param map[string]string) *DeleteTrainingExampleOptions {
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
func (*DiscoveryV1) NewDeleteUserDataOptions(customerID string) *DeleteUserDataOptions {
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

// DiskUsage : Summary of the disk usage statistics for the environment.
type DiskUsage struct {
	// Number of bytes within the environment's disk capacity that are currently used to store data.
	UsedBytes *int64 `json:"used_bytes,omitempty"`

	// Total number of bytes available in the environment's disk capacity.
	MaximumAllowedBytes *int64 `json:"maximum_allowed_bytes,omitempty"`
}

// UnmarshalDiskUsage unmarshals an instance of DiskUsage from the specified map of raw messages.
func UnmarshalDiskUsage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DiskUsage)
	err = core.UnmarshalPrimitive(m, "used_bytes", &obj.UsedBytes)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "maximum_allowed_bytes", &obj.MaximumAllowedBytes)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DocumentAccepted : Information returned after an uploaded document is accepted.
type DocumentAccepted struct {
	// The unique identifier of the ingested document.
	DocumentID *string `json:"document_id,omitempty"`

	// Status of the document in the ingestion process. A status of `processing` is returned for documents that are
	// ingested with a *version* date before `2019-01-01`. The `pending` status is returned for all others.
	Status *string `json:"status,omitempty"`

	// Array of notices produced by the document-ingestion process.
	Notices []Notice `json:"notices,omitempty"`
}

// Constants associated with the DocumentAccepted.Status property.
// Status of the document in the ingestion process. A status of `processing` is returned for documents that are ingested
// with a *version* date before `2019-01-01`. The `pending` status is returned for all others.
const (
	DocumentAcceptedStatusPendingConst    = "pending"
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
	err = core.UnmarshalModel(m, "notices", &obj.Notices, UnmarshalNotice)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DocumentCounts : Object containing collection document count information.
type DocumentCounts struct {
	// The total number of available documents in the collection.
	Available *int64 `json:"available,omitempty"`

	// The number of documents in the collection that are currently being processed.
	Processing *int64 `json:"processing,omitempty"`

	// The number of documents in the collection that failed to be ingested.
	Failed *int64 `json:"failed,omitempty"`

	// The number of documents that have been uploaded to the collection, but have not yet started processing.
	Pending *int64 `json:"pending,omitempty"`
}

// UnmarshalDocumentCounts unmarshals an instance of DocumentCounts from the specified map of raw messages.
func UnmarshalDocumentCounts(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DocumentCounts)
	err = core.UnmarshalPrimitive(m, "available", &obj.Available)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "processing", &obj.Processing)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "failed", &obj.Failed)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pending", &obj.Pending)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DocumentStatus : Status information about a submitted document.
type DocumentStatus struct {
	// The unique identifier of the document.
	DocumentID *string `json:"document_id" validate:"required"`

	// The unique identifier for the configuration.
	ConfigurationID *string `json:"configuration_id,omitempty"`

	// Status of the document in the ingestion process.
	Status *string `json:"status" validate:"required"`

	// Description of the document status.
	StatusDescription *string `json:"status_description" validate:"required"`

	// Name of the original source file (if available).
	Filename *string `json:"filename,omitempty"`

	// The type of the original source file.
	FileType *string `json:"file_type,omitempty"`

	// The SHA-1 hash of the original source file (formatted as a hexadecimal string).
	Sha1 *string `json:"sha1,omitempty"`

	// Array of notices produced by the document-ingestion process.
	Notices []Notice `json:"notices" validate:"required"`
}

// Constants associated with the DocumentStatus.Status property.
// Status of the document in the ingestion process.
const (
	DocumentStatusStatusAvailableConst            = "available"
	DocumentStatusStatusAvailableWithNoticesConst = "available with notices"
	DocumentStatusStatusFailedConst               = "failed"
	DocumentStatusStatusPendingConst              = "pending"
	DocumentStatusStatusProcessingConst           = "processing"
)

// Constants associated with the DocumentStatus.FileType property.
// The type of the original source file.
const (
	DocumentStatusFileTypeHTMLConst = "html"
	DocumentStatusFileTypeJSONConst = "json"
	DocumentStatusFileTypePDFConst  = "pdf"
	DocumentStatusFileTypeWordConst = "word"
)

// UnmarshalDocumentStatus unmarshals an instance of DocumentStatus from the specified map of raw messages.
func UnmarshalDocumentStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DocumentStatus)
	err = core.UnmarshalPrimitive(m, "document_id", &obj.DocumentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "configuration_id", &obj.ConfigurationID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status_description", &obj.StatusDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "filename", &obj.Filename)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "file_type", &obj.FileType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sha1", &obj.Sha1)
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

// Enrichment : Enrichment step to perform on the document. Each enrichment is performed on the specified field in the order that
// they are listed in the configuration.
type Enrichment struct {
	// Describes what the enrichment step does.
	Description *string `json:"description,omitempty"`

	// Field where enrichments will be stored. This field must already exist or be at most 1 level deeper than an existing
	// field. For example, if `text` is a top-level field with no sub-fields, `text.foo` is a valid destination but
	// `text.foo.bar` is not.
	DestinationField *string `json:"destination_field" validate:"required"`

	// Field to be enriched.
	//
	// Arrays can be specified as the **source_field** if the **enrichment** service for this enrichment is set to
	// `natural_language_undstanding`.
	SourceField *string `json:"source_field" validate:"required"`

	// Indicates that the enrichments will overwrite the destination_field field if it already exists.
	Overwrite *bool `json:"overwrite,omitempty"`

	// Name of the enrichment service to call. Current options are `natural_language_understanding` and `elements`.
	//
	//  When using `natual_language_understanding`, the **options** object must contain Natural Language Understanding
	// options.
	//
	// When using `elements` the **options** object must contain Element Classification options. Additionally, when using
	// the `elements` enrichment the configuration specified and files ingested must meet all the criteria specified in
	// [the
	// documentation](https://cloud.ibm.com/docs/discovery?topic=discovery-element-classification#element-classification).
	Enrichment *string `json:"enrichment" validate:"required"`

	// If true, then most errors generated during the enrichment process will be treated as warnings and will not cause the
	// document to fail processing.
	IgnoreDownstreamErrors *bool `json:"ignore_downstream_errors,omitempty"`

	// Options which are specific to a particular enrichment.
	Options *EnrichmentOptions `json:"options,omitempty"`
}

// NewEnrichment : Instantiate Enrichment (Generic Model Constructor)
func (*DiscoveryV1) NewEnrichment(destinationField string, sourceField string, enrichment string) (_model *Enrichment, err error) {
	_model = &Enrichment{
		DestinationField: core.StringPtr(destinationField),
		SourceField:      core.StringPtr(sourceField),
		Enrichment:       core.StringPtr(enrichment),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalEnrichment unmarshals an instance of Enrichment from the specified map of raw messages.
func UnmarshalEnrichment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Enrichment)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "destination_field", &obj.DestinationField)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source_field", &obj.SourceField)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "overwrite", &obj.Overwrite)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enrichment", &obj.Enrichment)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ignore_downstream_errors", &obj.IgnoreDownstreamErrors)
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

// EnrichmentOptions : Options which are specific to a particular enrichment.
type EnrichmentOptions struct {
	// Object containing Natural Language Understanding features to be used.
	Features *NluEnrichmentFeatures `json:"features,omitempty"`

	// ISO 639-1 code indicating the language to use for the analysis. This code overrides the automatic language detection
	// performed by the service. Valid codes are `ar` (Arabic), `en` (English), `fr` (French), `de` (German), `it`
	// (Italian), `pt` (Portuguese), `ru` (Russian), `es` (Spanish), and `sv` (Swedish). **Note:** Not all features support
	// all languages, automatic detection is recommended.
	Language *string `json:"language,omitempty"`

	// For use with `elements` enrichments only. The element extraction model to use. The only model available is
	// `contract`.
	Model *string `json:"model,omitempty"`
}

// Constants associated with the EnrichmentOptions.Language property.
// ISO 639-1 code indicating the language to use for the analysis. This code overrides the automatic language detection
// performed by the service. Valid codes are `ar` (Arabic), `en` (English), `fr` (French), `de` (German), `it`
// (Italian), `pt` (Portuguese), `ru` (Russian), `es` (Spanish), and `sv` (Swedish). **Note:** Not all features support
// all languages, automatic detection is recommended.
const (
	EnrichmentOptionsLanguageArConst = "ar"
	EnrichmentOptionsLanguageDeConst = "de"
	EnrichmentOptionsLanguageEnConst = "en"
	EnrichmentOptionsLanguageEsConst = "es"
	EnrichmentOptionsLanguageFrConst = "fr"
	EnrichmentOptionsLanguageItConst = "it"
	EnrichmentOptionsLanguagePtConst = "pt"
	EnrichmentOptionsLanguageRuConst = "ru"
	EnrichmentOptionsLanguageSvConst = "sv"
)

// UnmarshalEnrichmentOptions unmarshals an instance of EnrichmentOptions from the specified map of raw messages.
func UnmarshalEnrichmentOptions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EnrichmentOptions)
	err = core.UnmarshalModel(m, "features", &obj.Features, UnmarshalNluEnrichmentFeatures)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "language", &obj.Language)
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

// Environment : Details about an environment.
type Environment struct {
	// Unique identifier for the environment.
	EnvironmentID *string `json:"environment_id,omitempty"`

	// Name that identifies the environment.
	Name *string `json:"name,omitempty"`

	// Description of the environment.
	Description *string `json:"description,omitempty"`

	// Creation date of the environment, in the format `yyyy-MM-dd'T'HH:mm:ss.SSS'Z'`.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// Date of most recent environment update, in the format `yyyy-MM-dd'T'HH:mm:ss.SSS'Z'`.
	Updated *strfmt.DateTime `json:"updated,omitempty"`

	// Current status of the environment. `resizing` is displayed when a request to increase the environment size has been
	// made, but is still in the process of being completed.
	Status *string `json:"status,omitempty"`

	// If `true`, the environment contains read-only collections that are maintained by IBM.
	ReadOnly *bool `json:"read_only,omitempty"`

	// Current size of the environment.
	Size *string `json:"size,omitempty"`

	// The new size requested for this environment. Only returned when the environment *status* is `resizing`.
	//
	// *Note:* Querying and indexing can still be performed during an environment upsize.
	RequestedSize *string `json:"requested_size,omitempty"`

	// Details about the resource usage and capacity of the environment.
	IndexCapacity *IndexCapacity `json:"index_capacity,omitempty"`

	// Information about the Continuous Relevancy Training for this environment.
	SearchStatus *SearchStatus `json:"search_status,omitempty"`
}

// Constants associated with the Environment.Status property.
// Current status of the environment. `resizing` is displayed when a request to increase the environment size has been
// made, but is still in the process of being completed.
const (
	EnvironmentStatusActiveConst      = "active"
	EnvironmentStatusMaintenanceConst = "maintenance"
	EnvironmentStatusPendingConst     = "pending"
	EnvironmentStatusResizingConst    = "resizing"
)

// Constants associated with the Environment.Size property.
// Current size of the environment.
const (
	EnvironmentSizeLConst    = "L"
	EnvironmentSizeLtConst   = "LT"
	EnvironmentSizeMConst    = "M"
	EnvironmentSizeMlConst   = "ML"
	EnvironmentSizeMsConst   = "MS"
	EnvironmentSizeSConst    = "S"
	EnvironmentSizeXlConst   = "XL"
	EnvironmentSizeXsConst   = "XS"
	EnvironmentSizeXxlConst  = "XXL"
	EnvironmentSizeXxxlConst = "XXXL"
)

// UnmarshalEnvironment unmarshals an instance of Environment from the specified map of raw messages.
func UnmarshalEnvironment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Environment)
	err = core.UnmarshalPrimitive(m, "environment_id", &obj.EnvironmentID)
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
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "read_only", &obj.ReadOnly)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "size", &obj.Size)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "requested_size", &obj.RequestedSize)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "index_capacity", &obj.IndexCapacity, UnmarshalIndexCapacity)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "search_status", &obj.SearchStatus, UnmarshalSearchStatus)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EnvironmentDocuments : Summary of the document usage statistics for the environment.
type EnvironmentDocuments struct {
	// Number of documents indexed for the environment.
	Available *int64 `json:"available,omitempty"`

	// Total number of documents allowed in the environment's capacity.
	MaximumAllowed *int64 `json:"maximum_allowed,omitempty"`
}

// UnmarshalEnvironmentDocuments unmarshals an instance of EnvironmentDocuments from the specified map of raw messages.
func UnmarshalEnvironmentDocuments(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EnvironmentDocuments)
	err = core.UnmarshalPrimitive(m, "available", &obj.Available)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "maximum_allowed", &obj.MaximumAllowed)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EventData : Query event data object.
type EventData struct {
	// The **environment_id** associated with the query that the event is associated with.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The session token that was returned as part of the query results that this event is associated with.
	SessionToken *string `json:"session_token" validate:"required"`

	// The optional timestamp for the event that was created. If not provided, the time that the event was created in the
	// log was used.
	ClientTimestamp *strfmt.DateTime `json:"client_timestamp,omitempty"`

	// The rank of the result item which the event is associated with.
	DisplayRank *int64 `json:"display_rank,omitempty"`

	// The **collection_id** of the document that this event is associated with.
	CollectionID *string `json:"collection_id" validate:"required"`

	// The **document_id** of the document that this event is associated with.
	DocumentID *string `json:"document_id" validate:"required"`

	// The query identifier stored in the log. The query and any events associated with that query are stored with the same
	// **query_id**.
	QueryID *string `json:"query_id,omitempty"`
}

// NewEventData : Instantiate EventData (Generic Model Constructor)
func (*DiscoveryV1) NewEventData(environmentID string, sessionToken string, collectionID string, documentID string) (_model *EventData, err error) {
	_model = &EventData{
		EnvironmentID: core.StringPtr(environmentID),
		SessionToken:  core.StringPtr(sessionToken),
		CollectionID:  core.StringPtr(collectionID),
		DocumentID:    core.StringPtr(documentID),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalEventData unmarshals an instance of EventData from the specified map of raw messages.
func UnmarshalEventData(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EventData)
	err = core.UnmarshalPrimitive(m, "environment_id", &obj.EnvironmentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "session_token", &obj.SessionToken)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "client_timestamp", &obj.ClientTimestamp)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "display_rank", &obj.DisplayRank)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "collection_id", &obj.CollectionID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "document_id", &obj.DocumentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "query_id", &obj.QueryID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Expansion : An expansion definition. Each object respresents one set of expandable strings. For example, you could have
// expansions for the word `hot` in one object, and expansions for the word `cold` in another.
type Expansion struct {
	// A list of terms that will be expanded for this expansion. If specified, only the items in this list are expanded.
	InputTerms []string `json:"input_terms,omitempty"`

	// A list of terms that this expansion will be expanded to. If specified without **input_terms**, it also functions as
	// the input term list.
	ExpandedTerms []string `json:"expanded_terms" validate:"required"`
}

// NewExpansion : Instantiate Expansion (Generic Model Constructor)
func (*DiscoveryV1) NewExpansion(expandedTerms []string) (_model *Expansion, err error) {
	_model = &Expansion{
		ExpandedTerms: expandedTerms,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalExpansion unmarshals an instance of Expansion from the specified map of raw messages.
func UnmarshalExpansion(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Expansion)
	err = core.UnmarshalPrimitive(m, "input_terms", &obj.InputTerms)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "expanded_terms", &obj.ExpandedTerms)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Expansions : The query expansion definitions for the specified collection.
type Expansions struct {
	// An array of query expansion definitions.
	//
	//  Each object in the **expansions** array represents a term or set of terms that will be expanded into other terms.
	// Each expansion object can be configured as bidirectional or unidirectional. Bidirectional means that all terms are
	// expanded to all other terms in the object. Unidirectional means that a set list of terms can be expanded into a
	// second list of terms.
	//
	//  To create a bi-directional expansion specify an **expanded_terms** array. When found in a query, all items in the
	// **expanded_terms** array are then expanded to the other items in the same array.
	//
	//  To create a uni-directional expansion, specify both an array of **input_terms** and an array of **expanded_terms**.
	// When items in the **input_terms** array are present in a query, they are expanded using the items listed in the
	// **expanded_terms** array.
	Expansions []Expansion `json:"expansions" validate:"required"`
}

// NewExpansions : Instantiate Expansions (Generic Model Constructor)
func (*DiscoveryV1) NewExpansions(expansions []Expansion) (_model *Expansions, err error) {
	_model = &Expansions{
		Expansions: expansions,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalExpansions unmarshals an instance of Expansions from the specified map of raw messages.
func UnmarshalExpansions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Expansions)
	err = core.UnmarshalModel(m, "expansions", &obj.Expansions, UnmarshalExpansion)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FederatedQueryNoticesOptions : The FederatedQueryNotices options.
type FederatedQueryNoticesOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// A comma-separated list of collection IDs to be queried against.
	CollectionIds []string `json:"-" validate:"required"`

	// A cacheable query that excludes documents that don't mention the query content. Filter searches are better for
	// metadata-type searches and for assessing the concepts in the data set.
	Filter *string `json:"-"`

	// A query search returns all documents in your data set with full enrichments and full text, but with the most
	// relevant documents listed first.
	Query *string `json:"-"`

	// A natural language query that returns relevant documents by utilizing training data and natural language
	// understanding.
	NaturalLanguageQuery *string `json:"-"`

	// An aggregation search that returns an exact answer by combining query search with filters. Useful for applications
	// to build lists, tables, and time series. For a full list of possible aggregations, see the Query reference.
	Aggregation *string `json:"-"`

	// Number of results to return. The maximum for the **count** and **offset** values together in any one query is
	// **10000**.
	Count *int64 `json:"-"`

	// A comma-separated list of the portion of the document hierarchy to return.
	Return []string `json:"-"`

	// The number of query results to skip at the beginning. For example, if the total number of results that are returned
	// is 10 and the offset is 8, it returns the last two results. The maximum for the **count** and **offset** values
	// together in any one query is **10000**.
	Offset *int64 `json:"-"`

	// A comma-separated list of fields in the document to sort on. You can optionally specify a sort direction by
	// prefixing the field with `-` for descending or `+` for ascending. Ascending is the default sort direction if no
	// prefix is specified.
	Sort []string `json:"-"`

	// When true, a highlight field is returned for each result which contains the fields which match the query with
	// `<em></em>` tags around the matching query terms.
	Highlight *bool `json:"-"`

	// When specified, duplicate results based on the field specified are removed from the returned results. Duplicate
	// comparison is limited to the current query only, **offset** is not considered. This parameter is currently Beta
	// functionality.
	DeduplicateField *string `json:"-"`

	// When `true`, results are returned based on their similarity to the document IDs specified in the
	// **similar.document_ids** parameter.
	Similar *bool `json:"-"`

	// A comma-separated list of document IDs to find similar documents.
	//
	// **Tip:** Include the **natural_language_query** parameter to expand the scope of the document similarity search with
	// the natural language query. Other query parameters, such as **filter** and **query**, are subsequently applied and
	// reduce the scope.
	SimilarDocumentIds []string `json:"-"`

	// A comma-separated list of field names that are used as a basis for comparison to identify similar documents. If not
	// specified, the entire document is used for comparison.
	SimilarFields []string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewFederatedQueryNoticesOptions : Instantiate FederatedQueryNoticesOptions
func (*DiscoveryV1) NewFederatedQueryNoticesOptions(environmentID string, collectionIds []string) *FederatedQueryNoticesOptions {
	return &FederatedQueryNoticesOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionIds: collectionIds,
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *FederatedQueryNoticesOptions) SetEnvironmentID(environmentID string) *FederatedQueryNoticesOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetCollectionIds : Allow user to set CollectionIds
func (_options *FederatedQueryNoticesOptions) SetCollectionIds(collectionIds []string) *FederatedQueryNoticesOptions {
	_options.CollectionIds = collectionIds
	return _options
}

// SetFilter : Allow user to set Filter
func (_options *FederatedQueryNoticesOptions) SetFilter(filter string) *FederatedQueryNoticesOptions {
	_options.Filter = core.StringPtr(filter)
	return _options
}

// SetQuery : Allow user to set Query
func (_options *FederatedQueryNoticesOptions) SetQuery(query string) *FederatedQueryNoticesOptions {
	_options.Query = core.StringPtr(query)
	return _options
}

// SetNaturalLanguageQuery : Allow user to set NaturalLanguageQuery
func (_options *FederatedQueryNoticesOptions) SetNaturalLanguageQuery(naturalLanguageQuery string) *FederatedQueryNoticesOptions {
	_options.NaturalLanguageQuery = core.StringPtr(naturalLanguageQuery)
	return _options
}

// SetAggregation : Allow user to set Aggregation
func (_options *FederatedQueryNoticesOptions) SetAggregation(aggregation string) *FederatedQueryNoticesOptions {
	_options.Aggregation = core.StringPtr(aggregation)
	return _options
}

// SetCount : Allow user to set Count
func (_options *FederatedQueryNoticesOptions) SetCount(count int64) *FederatedQueryNoticesOptions {
	_options.Count = core.Int64Ptr(count)
	return _options
}

// SetReturn : Allow user to set Return
func (_options *FederatedQueryNoticesOptions) SetReturn(returnVar []string) *FederatedQueryNoticesOptions {
	_options.Return = returnVar
	return _options
}

// SetOffset : Allow user to set Offset
func (_options *FederatedQueryNoticesOptions) SetOffset(offset int64) *FederatedQueryNoticesOptions {
	_options.Offset = core.Int64Ptr(offset)
	return _options
}

// SetSort : Allow user to set Sort
func (_options *FederatedQueryNoticesOptions) SetSort(sort []string) *FederatedQueryNoticesOptions {
	_options.Sort = sort
	return _options
}

// SetHighlight : Allow user to set Highlight
func (_options *FederatedQueryNoticesOptions) SetHighlight(highlight bool) *FederatedQueryNoticesOptions {
	_options.Highlight = core.BoolPtr(highlight)
	return _options
}

// SetDeduplicateField : Allow user to set DeduplicateField
func (_options *FederatedQueryNoticesOptions) SetDeduplicateField(deduplicateField string) *FederatedQueryNoticesOptions {
	_options.DeduplicateField = core.StringPtr(deduplicateField)
	return _options
}

// SetSimilar : Allow user to set Similar
func (_options *FederatedQueryNoticesOptions) SetSimilar(similar bool) *FederatedQueryNoticesOptions {
	_options.Similar = core.BoolPtr(similar)
	return _options
}

// SetSimilarDocumentIds : Allow user to set SimilarDocumentIds
func (_options *FederatedQueryNoticesOptions) SetSimilarDocumentIds(similarDocumentIds []string) *FederatedQueryNoticesOptions {
	_options.SimilarDocumentIds = similarDocumentIds
	return _options
}

// SetSimilarFields : Allow user to set SimilarFields
func (_options *FederatedQueryNoticesOptions) SetSimilarFields(similarFields []string) *FederatedQueryNoticesOptions {
	_options.SimilarFields = similarFields
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *FederatedQueryNoticesOptions) SetHeaders(param map[string]string) *FederatedQueryNoticesOptions {
	options.Headers = param
	return options
}

// FederatedQueryOptions : The FederatedQuery options.
type FederatedQueryOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// A comma-separated list of collection IDs to be queried against.
	CollectionIds *string `json:"collection_ids" validate:"required"`

	// A cacheable query that excludes documents that don't mention the query content. Filter searches are better for
	// metadata-type searches and for assessing the concepts in the data set.
	Filter *string `json:"filter,omitempty"`

	// A query search returns all documents in your data set with full enrichments and full text, but with the most
	// relevant documents listed first. Use a query search when you want to find the most relevant search results.
	Query *string `json:"query,omitempty"`

	// A natural language query that returns relevant documents by utilizing training data and natural language
	// understanding.
	NaturalLanguageQuery *string `json:"natural_language_query,omitempty"`

	// A passages query that returns the most relevant passages from the results.
	Passages *bool `json:"passages,omitempty"`

	// An aggregation search that returns an exact answer by combining query search with filters. Useful for applications
	// to build lists, tables, and time series. For a full list of possible aggregations, see the Query reference.
	Aggregation *string `json:"aggregation,omitempty"`

	// Number of results to return.
	Count *int64 `json:"count,omitempty"`

	// A comma-separated list of the portion of the document hierarchy to return.
	Return *string `json:"return,omitempty"`

	// The number of query results to skip at the beginning. For example, if the total number of results that are returned
	// is 10 and the offset is 8, it returns the last two results.
	Offset *int64 `json:"offset,omitempty"`

	// A comma-separated list of fields in the document to sort on. You can optionally specify a sort direction by
	// prefixing the field with `-` for descending or `+` for ascending. Ascending is the default sort direction if no
	// prefix is specified. This parameter cannot be used in the same query as the **bias** parameter.
	Sort *string `json:"sort,omitempty"`

	// When true, a highlight field is returned for each result which contains the fields which match the query with
	// `<em></em>` tags around the matching query terms.
	Highlight *bool `json:"highlight,omitempty"`

	// A comma-separated list of fields that passages are drawn from. If this parameter not specified, then all top-level
	// fields are included.
	PassagesFields *string `json:"passages.fields,omitempty"`

	// The maximum number of passages to return. The search returns fewer passages if the requested total is not found. The
	// default is `10`. The maximum is `100`.
	PassagesCount *int64 `json:"passages.count,omitempty"`

	// The approximate number of characters that any one passage will have.
	PassagesCharacters *int64 `json:"passages.characters,omitempty"`

	// When `true`, and used with a Watson Discovery News collection, duplicate results (based on the contents of the
	// **title** field) are removed. Duplicate comparison is limited to the current query only; **offset** is not
	// considered. This parameter is currently Beta functionality.
	Deduplicate *bool `json:"deduplicate,omitempty"`

	// When specified, duplicate results based on the field specified are removed from the returned results. Duplicate
	// comparison is limited to the current query only, **offset** is not considered. This parameter is currently Beta
	// functionality.
	DeduplicateField *string `json:"deduplicate.field,omitempty"`

	// When `true`, results are returned based on their similarity to the document IDs specified in the
	// **similar.document_ids** parameter.
	Similar *bool `json:"similar,omitempty"`

	// A comma-separated list of document IDs to find similar documents.
	//
	// **Tip:** Include the **natural_language_query** parameter to expand the scope of the document similarity search with
	// the natural language query. Other query parameters, such as **filter** and **query**, are subsequently applied and
	// reduce the scope.
	SimilarDocumentIds *string `json:"similar.document_ids,omitempty"`

	// A comma-separated list of field names that are used as a basis for comparison to identify similar documents. If not
	// specified, the entire document is used for comparison.
	SimilarFields *string `json:"similar.fields,omitempty"`

	// Field which the returned results will be biased against. The specified field must be either a **date** or **number**
	// format. When a **date** type field is specified returned results are biased towards field values closer to the
	// current date. When a **number** type field is specified, returned results are biased towards higher field values.
	// This parameter cannot be used in the same query as the **sort** parameter.
	Bias *string `json:"bias,omitempty"`

	// If `true`, queries are not stored in the Discovery **Logs** endpoint.
	XWatsonLoggingOptOut *bool `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewFederatedQueryOptions : Instantiate FederatedQueryOptions
func (*DiscoveryV1) NewFederatedQueryOptions(environmentID string, collectionIds string) *FederatedQueryOptions {
	return &FederatedQueryOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionIds: core.StringPtr(collectionIds),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *FederatedQueryOptions) SetEnvironmentID(environmentID string) *FederatedQueryOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetCollectionIds : Allow user to set CollectionIds
func (_options *FederatedQueryOptions) SetCollectionIds(collectionIds string) *FederatedQueryOptions {
	_options.CollectionIds = core.StringPtr(collectionIds)
	return _options
}

// SetFilter : Allow user to set Filter
func (_options *FederatedQueryOptions) SetFilter(filter string) *FederatedQueryOptions {
	_options.Filter = core.StringPtr(filter)
	return _options
}

// SetQuery : Allow user to set Query
func (_options *FederatedQueryOptions) SetQuery(query string) *FederatedQueryOptions {
	_options.Query = core.StringPtr(query)
	return _options
}

// SetNaturalLanguageQuery : Allow user to set NaturalLanguageQuery
func (_options *FederatedQueryOptions) SetNaturalLanguageQuery(naturalLanguageQuery string) *FederatedQueryOptions {
	_options.NaturalLanguageQuery = core.StringPtr(naturalLanguageQuery)
	return _options
}

// SetPassages : Allow user to set Passages
func (_options *FederatedQueryOptions) SetPassages(passages bool) *FederatedQueryOptions {
	_options.Passages = core.BoolPtr(passages)
	return _options
}

// SetAggregation : Allow user to set Aggregation
func (_options *FederatedQueryOptions) SetAggregation(aggregation string) *FederatedQueryOptions {
	_options.Aggregation = core.StringPtr(aggregation)
	return _options
}

// SetCount : Allow user to set Count
func (_options *FederatedQueryOptions) SetCount(count int64) *FederatedQueryOptions {
	_options.Count = core.Int64Ptr(count)
	return _options
}

// SetReturn : Allow user to set Return
func (_options *FederatedQueryOptions) SetReturn(returnVar string) *FederatedQueryOptions {
	_options.Return = core.StringPtr(returnVar)
	return _options
}

// SetOffset : Allow user to set Offset
func (_options *FederatedQueryOptions) SetOffset(offset int64) *FederatedQueryOptions {
	_options.Offset = core.Int64Ptr(offset)
	return _options
}

// SetSort : Allow user to set Sort
func (_options *FederatedQueryOptions) SetSort(sort string) *FederatedQueryOptions {
	_options.Sort = core.StringPtr(sort)
	return _options
}

// SetHighlight : Allow user to set Highlight
func (_options *FederatedQueryOptions) SetHighlight(highlight bool) *FederatedQueryOptions {
	_options.Highlight = core.BoolPtr(highlight)
	return _options
}

// SetPassagesFields : Allow user to set PassagesFields
func (_options *FederatedQueryOptions) SetPassagesFields(passagesFields string) *FederatedQueryOptions {
	_options.PassagesFields = core.StringPtr(passagesFields)
	return _options
}

// SetPassagesCount : Allow user to set PassagesCount
func (_options *FederatedQueryOptions) SetPassagesCount(passagesCount int64) *FederatedQueryOptions {
	_options.PassagesCount = core.Int64Ptr(passagesCount)
	return _options
}

// SetPassagesCharacters : Allow user to set PassagesCharacters
func (_options *FederatedQueryOptions) SetPassagesCharacters(passagesCharacters int64) *FederatedQueryOptions {
	_options.PassagesCharacters = core.Int64Ptr(passagesCharacters)
	return _options
}

// SetDeduplicate : Allow user to set Deduplicate
func (_options *FederatedQueryOptions) SetDeduplicate(deduplicate bool) *FederatedQueryOptions {
	_options.Deduplicate = core.BoolPtr(deduplicate)
	return _options
}

// SetDeduplicateField : Allow user to set DeduplicateField
func (_options *FederatedQueryOptions) SetDeduplicateField(deduplicateField string) *FederatedQueryOptions {
	_options.DeduplicateField = core.StringPtr(deduplicateField)
	return _options
}

// SetSimilar : Allow user to set Similar
func (_options *FederatedQueryOptions) SetSimilar(similar bool) *FederatedQueryOptions {
	_options.Similar = core.BoolPtr(similar)
	return _options
}

// SetSimilarDocumentIds : Allow user to set SimilarDocumentIds
func (_options *FederatedQueryOptions) SetSimilarDocumentIds(similarDocumentIds string) *FederatedQueryOptions {
	_options.SimilarDocumentIds = core.StringPtr(similarDocumentIds)
	return _options
}

// SetSimilarFields : Allow user to set SimilarFields
func (_options *FederatedQueryOptions) SetSimilarFields(similarFields string) *FederatedQueryOptions {
	_options.SimilarFields = core.StringPtr(similarFields)
	return _options
}

// SetBias : Allow user to set Bias
func (_options *FederatedQueryOptions) SetBias(bias string) *FederatedQueryOptions {
	_options.Bias = core.StringPtr(bias)
	return _options
}

// SetXWatsonLoggingOptOut : Allow user to set XWatsonLoggingOptOut
func (_options *FederatedQueryOptions) SetXWatsonLoggingOptOut(xWatsonLoggingOptOut bool) *FederatedQueryOptions {
	_options.XWatsonLoggingOptOut = core.BoolPtr(xWatsonLoggingOptOut)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *FederatedQueryOptions) SetHeaders(param map[string]string) *FederatedQueryOptions {
	options.Headers = param
	return options
}

// Field : Object containing field details.
type Field struct {
	// The name of the field.
	Field *string `json:"field,omitempty"`

	// The type of the field.
	Type *string `json:"type,omitempty"`
}

// Constants associated with the Field.Type property.
// The type of the field.
const (
	FieldTypeBinaryConst  = "binary"
	FieldTypeBooleanConst = "boolean"
	FieldTypeByteConst    = "byte"
	FieldTypeDateConst    = "date"
	FieldTypeDoubleConst  = "double"
	FieldTypeFloatConst   = "float"
	FieldTypeIntegerConst = "integer"
	FieldTypeLongConst    = "long"
	FieldTypeNestedConst  = "nested"
	FieldTypeShortConst   = "short"
	FieldTypeStringConst  = "string"
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FontSetting : Font matching configuration.
type FontSetting struct {
	// The HTML heading level that any content with the matching font is converted to.
	Level *int64 `json:"level,omitempty"`

	// The minimum size of the font to match.
	MinSize *int64 `json:"min_size,omitempty"`

	// The maximum size of the font to match.
	MaxSize *int64 `json:"max_size,omitempty"`

	// When `true`, the font is matched if it is bold.
	Bold *bool `json:"bold,omitempty"`

	// When `true`, the font is matched if it is italic.
	Italic *bool `json:"italic,omitempty"`

	// The name of the font.
	Name *string `json:"name,omitempty"`
}

// UnmarshalFontSetting unmarshals an instance of FontSetting from the specified map of raw messages.
func UnmarshalFontSetting(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FontSetting)
	err = core.UnmarshalPrimitive(m, "level", &obj.Level)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "min_size", &obj.MinSize)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "max_size", &obj.MaxSize)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "bold", &obj.Bold)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "italic", &obj.Italic)
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

// Gateway : Object describing a specific gateway.
type Gateway struct {
	// The gateway ID of the gateway.
	GatewayID *string `json:"gateway_id,omitempty"`

	// The user defined name of the gateway.
	Name *string `json:"name,omitempty"`

	// The current status of the gateway. `connected` means the gateway is connected to the remotly installed gateway.
	// `idle` means this gateway is not currently in use.
	Status *string `json:"status,omitempty"`

	// The generated **token** for this gateway. The value of this field is used when configuring the remotly installed
	// gateway.
	Token *string `json:"token,omitempty"`

	// The generated **token_id** for this gateway. The value of this field is used when configuring the remotly installed
	// gateway.
	TokenID *string `json:"token_id,omitempty"`
}

// Constants associated with the Gateway.Status property.
// The current status of the gateway. `connected` means the gateway is connected to the remotly installed gateway.
// `idle` means this gateway is not currently in use.
const (
	GatewayStatusConnectedConst = "connected"
	GatewayStatusIdleConst      = "idle"
)

// UnmarshalGateway unmarshals an instance of Gateway from the specified map of raw messages.
func UnmarshalGateway(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Gateway)
	err = core.UnmarshalPrimitive(m, "gateway_id", &obj.GatewayID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "token", &obj.Token)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "token_id", &obj.TokenID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GatewayDelete : Gatway deletion confirmation.
type GatewayDelete struct {
	// The gateway ID of the deleted gateway.
	GatewayID *string `json:"gateway_id,omitempty"`

	// The status of the request.
	Status *string `json:"status,omitempty"`
}

// UnmarshalGatewayDelete unmarshals an instance of GatewayDelete from the specified map of raw messages.
func UnmarshalGatewayDelete(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GatewayDelete)
	err = core.UnmarshalPrimitive(m, "gateway_id", &obj.GatewayID)
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

// GatewayList : Object containing gateways array.
type GatewayList struct {
	// Array of configured gateway connections.
	Gateways []Gateway `json:"gateways,omitempty"`
}

// UnmarshalGatewayList unmarshals an instance of GatewayList from the specified map of raw messages.
func UnmarshalGatewayList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GatewayList)
	err = core.UnmarshalModel(m, "gateways", &obj.Gateways, UnmarshalGateway)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetAutocompletionOptions : The GetAutocompletion options.
type GetAutocompletionOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The ID of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// The prefix to use for autocompletion. For example, the prefix `Ho` could autocomplete to `hot`, `housing`, or `how`.
	Prefix *string `json:"-" validate:"required"`

	// The field in the result documents that autocompletion suggestions are identified from.
	Field *string `json:"-"`

	// The number of autocompletion suggestions to return.
	Count *int64 `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAutocompletionOptions : Instantiate GetAutocompletionOptions
func (*DiscoveryV1) NewGetAutocompletionOptions(environmentID string, collectionID string, prefix string) *GetAutocompletionOptions {
	return &GetAutocompletionOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
		Prefix:        core.StringPtr(prefix),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *GetAutocompletionOptions) SetEnvironmentID(environmentID string) *GetAutocompletionOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetCollectionID : Allow user to set CollectionID
func (_options *GetAutocompletionOptions) SetCollectionID(collectionID string) *GetAutocompletionOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetPrefix : Allow user to set Prefix
func (_options *GetAutocompletionOptions) SetPrefix(prefix string) *GetAutocompletionOptions {
	_options.Prefix = core.StringPtr(prefix)
	return _options
}

// SetField : Allow user to set Field
func (_options *GetAutocompletionOptions) SetField(field string) *GetAutocompletionOptions {
	_options.Field = core.StringPtr(field)
	return _options
}

// SetCount : Allow user to set Count
func (_options *GetAutocompletionOptions) SetCount(count int64) *GetAutocompletionOptions {
	_options.Count = core.Int64Ptr(count)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetAutocompletionOptions) SetHeaders(param map[string]string) *GetAutocompletionOptions {
	options.Headers = param
	return options
}

// GetCollectionOptions : The GetCollection options.
type GetCollectionOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The ID of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCollectionOptions : Instantiate GetCollectionOptions
func (*DiscoveryV1) NewGetCollectionOptions(environmentID string, collectionID string) *GetCollectionOptions {
	return &GetCollectionOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *GetCollectionOptions) SetEnvironmentID(environmentID string) *GetCollectionOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
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

// GetConfigurationOptions : The GetConfiguration options.
type GetConfigurationOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The ID of the configuration.
	ConfigurationID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetConfigurationOptions : Instantiate GetConfigurationOptions
func (*DiscoveryV1) NewGetConfigurationOptions(environmentID string, configurationID string) *GetConfigurationOptions {
	return &GetConfigurationOptions{
		EnvironmentID:   core.StringPtr(environmentID),
		ConfigurationID: core.StringPtr(configurationID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *GetConfigurationOptions) SetEnvironmentID(environmentID string) *GetConfigurationOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetConfigurationID : Allow user to set ConfigurationID
func (_options *GetConfigurationOptions) SetConfigurationID(configurationID string) *GetConfigurationOptions {
	_options.ConfigurationID = core.StringPtr(configurationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetConfigurationOptions) SetHeaders(param map[string]string) *GetConfigurationOptions {
	options.Headers = param
	return options
}

// GetCredentialsOptions : The GetCredentials options.
type GetCredentialsOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The unique identifier for a set of source credentials.
	CredentialID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCredentialsOptions : Instantiate GetCredentialsOptions
func (*DiscoveryV1) NewGetCredentialsOptions(environmentID string, credentialID string) *GetCredentialsOptions {
	return &GetCredentialsOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CredentialID:  core.StringPtr(credentialID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *GetCredentialsOptions) SetEnvironmentID(environmentID string) *GetCredentialsOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetCredentialID : Allow user to set CredentialID
func (_options *GetCredentialsOptions) SetCredentialID(credentialID string) *GetCredentialsOptions {
	_options.CredentialID = core.StringPtr(credentialID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetCredentialsOptions) SetHeaders(param map[string]string) *GetCredentialsOptions {
	options.Headers = param
	return options
}

// GetDocumentStatusOptions : The GetDocumentStatus options.
type GetDocumentStatusOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The ID of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// The ID of the document.
	DocumentID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetDocumentStatusOptions : Instantiate GetDocumentStatusOptions
func (*DiscoveryV1) NewGetDocumentStatusOptions(environmentID string, collectionID string, documentID string) *GetDocumentStatusOptions {
	return &GetDocumentStatusOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
		DocumentID:    core.StringPtr(documentID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *GetDocumentStatusOptions) SetEnvironmentID(environmentID string) *GetDocumentStatusOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetCollectionID : Allow user to set CollectionID
func (_options *GetDocumentStatusOptions) SetCollectionID(collectionID string) *GetDocumentStatusOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetDocumentID : Allow user to set DocumentID
func (_options *GetDocumentStatusOptions) SetDocumentID(documentID string) *GetDocumentStatusOptions {
	_options.DocumentID = core.StringPtr(documentID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDocumentStatusOptions) SetHeaders(param map[string]string) *GetDocumentStatusOptions {
	options.Headers = param
	return options
}

// GetEnvironmentOptions : The GetEnvironment options.
type GetEnvironmentOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetEnvironmentOptions : Instantiate GetEnvironmentOptions
func (*DiscoveryV1) NewGetEnvironmentOptions(environmentID string) *GetEnvironmentOptions {
	return &GetEnvironmentOptions{
		EnvironmentID: core.StringPtr(environmentID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *GetEnvironmentOptions) SetEnvironmentID(environmentID string) *GetEnvironmentOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetEnvironmentOptions) SetHeaders(param map[string]string) *GetEnvironmentOptions {
	options.Headers = param
	return options
}

// GetGatewayOptions : The GetGateway options.
type GetGatewayOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The requested gateway ID.
	GatewayID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetGatewayOptions : Instantiate GetGatewayOptions
func (*DiscoveryV1) NewGetGatewayOptions(environmentID string, gatewayID string) *GetGatewayOptions {
	return &GetGatewayOptions{
		EnvironmentID: core.StringPtr(environmentID),
		GatewayID:     core.StringPtr(gatewayID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *GetGatewayOptions) SetEnvironmentID(environmentID string) *GetGatewayOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetGatewayID : Allow user to set GatewayID
func (_options *GetGatewayOptions) SetGatewayID(gatewayID string) *GetGatewayOptions {
	_options.GatewayID = core.StringPtr(gatewayID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetGatewayOptions) SetHeaders(param map[string]string) *GetGatewayOptions {
	options.Headers = param
	return options
}

// GetMetricsEventRateOptions : The GetMetricsEventRate options.
type GetMetricsEventRateOptions struct {
	// Metric is computed from data recorded after this timestamp; must be in `YYYY-MM-DDThh:mm:ssZ` format.
	StartTime *strfmt.DateTime `json:"-"`

	// Metric is computed from data recorded before this timestamp; must be in `YYYY-MM-DDThh:mm:ssZ` format.
	EndTime *strfmt.DateTime `json:"-"`

	// The type of result to consider when calculating the metric.
	ResultType *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetMetricsEventRateOptions.ResultType property.
// The type of result to consider when calculating the metric.
const (
	GetMetricsEventRateOptionsResultTypeDocumentConst = "document"
)

// NewGetMetricsEventRateOptions : Instantiate GetMetricsEventRateOptions
func (*DiscoveryV1) NewGetMetricsEventRateOptions() *GetMetricsEventRateOptions {
	return &GetMetricsEventRateOptions{}
}

// SetStartTime : Allow user to set StartTime
func (_options *GetMetricsEventRateOptions) SetStartTime(startTime *strfmt.DateTime) *GetMetricsEventRateOptions {
	_options.StartTime = startTime
	return _options
}

// SetEndTime : Allow user to set EndTime
func (_options *GetMetricsEventRateOptions) SetEndTime(endTime *strfmt.DateTime) *GetMetricsEventRateOptions {
	_options.EndTime = endTime
	return _options
}

// SetResultType : Allow user to set ResultType
func (_options *GetMetricsEventRateOptions) SetResultType(resultType string) *GetMetricsEventRateOptions {
	_options.ResultType = core.StringPtr(resultType)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetMetricsEventRateOptions) SetHeaders(param map[string]string) *GetMetricsEventRateOptions {
	options.Headers = param
	return options
}

// GetMetricsQueryEventOptions : The GetMetricsQueryEvent options.
type GetMetricsQueryEventOptions struct {
	// Metric is computed from data recorded after this timestamp; must be in `YYYY-MM-DDThh:mm:ssZ` format.
	StartTime *strfmt.DateTime `json:"-"`

	// Metric is computed from data recorded before this timestamp; must be in `YYYY-MM-DDThh:mm:ssZ` format.
	EndTime *strfmt.DateTime `json:"-"`

	// The type of result to consider when calculating the metric.
	ResultType *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetMetricsQueryEventOptions.ResultType property.
// The type of result to consider when calculating the metric.
const (
	GetMetricsQueryEventOptionsResultTypeDocumentConst = "document"
)

// NewGetMetricsQueryEventOptions : Instantiate GetMetricsQueryEventOptions
func (*DiscoveryV1) NewGetMetricsQueryEventOptions() *GetMetricsQueryEventOptions {
	return &GetMetricsQueryEventOptions{}
}

// SetStartTime : Allow user to set StartTime
func (_options *GetMetricsQueryEventOptions) SetStartTime(startTime *strfmt.DateTime) *GetMetricsQueryEventOptions {
	_options.StartTime = startTime
	return _options
}

// SetEndTime : Allow user to set EndTime
func (_options *GetMetricsQueryEventOptions) SetEndTime(endTime *strfmt.DateTime) *GetMetricsQueryEventOptions {
	_options.EndTime = endTime
	return _options
}

// SetResultType : Allow user to set ResultType
func (_options *GetMetricsQueryEventOptions) SetResultType(resultType string) *GetMetricsQueryEventOptions {
	_options.ResultType = core.StringPtr(resultType)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetMetricsQueryEventOptions) SetHeaders(param map[string]string) *GetMetricsQueryEventOptions {
	options.Headers = param
	return options
}

// GetMetricsQueryNoResultsOptions : The GetMetricsQueryNoResults options.
type GetMetricsQueryNoResultsOptions struct {
	// Metric is computed from data recorded after this timestamp; must be in `YYYY-MM-DDThh:mm:ssZ` format.
	StartTime *strfmt.DateTime `json:"-"`

	// Metric is computed from data recorded before this timestamp; must be in `YYYY-MM-DDThh:mm:ssZ` format.
	EndTime *strfmt.DateTime `json:"-"`

	// The type of result to consider when calculating the metric.
	ResultType *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetMetricsQueryNoResultsOptions.ResultType property.
// The type of result to consider when calculating the metric.
const (
	GetMetricsQueryNoResultsOptionsResultTypeDocumentConst = "document"
)

// NewGetMetricsQueryNoResultsOptions : Instantiate GetMetricsQueryNoResultsOptions
func (*DiscoveryV1) NewGetMetricsQueryNoResultsOptions() *GetMetricsQueryNoResultsOptions {
	return &GetMetricsQueryNoResultsOptions{}
}

// SetStartTime : Allow user to set StartTime
func (_options *GetMetricsQueryNoResultsOptions) SetStartTime(startTime *strfmt.DateTime) *GetMetricsQueryNoResultsOptions {
	_options.StartTime = startTime
	return _options
}

// SetEndTime : Allow user to set EndTime
func (_options *GetMetricsQueryNoResultsOptions) SetEndTime(endTime *strfmt.DateTime) *GetMetricsQueryNoResultsOptions {
	_options.EndTime = endTime
	return _options
}

// SetResultType : Allow user to set ResultType
func (_options *GetMetricsQueryNoResultsOptions) SetResultType(resultType string) *GetMetricsQueryNoResultsOptions {
	_options.ResultType = core.StringPtr(resultType)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetMetricsQueryNoResultsOptions) SetHeaders(param map[string]string) *GetMetricsQueryNoResultsOptions {
	options.Headers = param
	return options
}

// GetMetricsQueryOptions : The GetMetricsQuery options.
type GetMetricsQueryOptions struct {
	// Metric is computed from data recorded after this timestamp; must be in `YYYY-MM-DDThh:mm:ssZ` format.
	StartTime *strfmt.DateTime `json:"-"`

	// Metric is computed from data recorded before this timestamp; must be in `YYYY-MM-DDThh:mm:ssZ` format.
	EndTime *strfmt.DateTime `json:"-"`

	// The type of result to consider when calculating the metric.
	ResultType *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetMetricsQueryOptions.ResultType property.
// The type of result to consider when calculating the metric.
const (
	GetMetricsQueryOptionsResultTypeDocumentConst = "document"
)

// NewGetMetricsQueryOptions : Instantiate GetMetricsQueryOptions
func (*DiscoveryV1) NewGetMetricsQueryOptions() *GetMetricsQueryOptions {
	return &GetMetricsQueryOptions{}
}

// SetStartTime : Allow user to set StartTime
func (_options *GetMetricsQueryOptions) SetStartTime(startTime *strfmt.DateTime) *GetMetricsQueryOptions {
	_options.StartTime = startTime
	return _options
}

// SetEndTime : Allow user to set EndTime
func (_options *GetMetricsQueryOptions) SetEndTime(endTime *strfmt.DateTime) *GetMetricsQueryOptions {
	_options.EndTime = endTime
	return _options
}

// SetResultType : Allow user to set ResultType
func (_options *GetMetricsQueryOptions) SetResultType(resultType string) *GetMetricsQueryOptions {
	_options.ResultType = core.StringPtr(resultType)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetMetricsQueryOptions) SetHeaders(param map[string]string) *GetMetricsQueryOptions {
	options.Headers = param
	return options
}

// GetMetricsQueryTokenEventOptions : The GetMetricsQueryTokenEvent options.
type GetMetricsQueryTokenEventOptions struct {
	// Number of results to return. The maximum for the **count** and **offset** values together in any one query is
	// **10000**.
	Count *int64 `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetMetricsQueryTokenEventOptions : Instantiate GetMetricsQueryTokenEventOptions
func (*DiscoveryV1) NewGetMetricsQueryTokenEventOptions() *GetMetricsQueryTokenEventOptions {
	return &GetMetricsQueryTokenEventOptions{}
}

// SetCount : Allow user to set Count
func (_options *GetMetricsQueryTokenEventOptions) SetCount(count int64) *GetMetricsQueryTokenEventOptions {
	_options.Count = core.Int64Ptr(count)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetMetricsQueryTokenEventOptions) SetHeaders(param map[string]string) *GetMetricsQueryTokenEventOptions {
	options.Headers = param
	return options
}

// GetStopwordListStatusOptions : The GetStopwordListStatus options.
type GetStopwordListStatusOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The ID of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetStopwordListStatusOptions : Instantiate GetStopwordListStatusOptions
func (*DiscoveryV1) NewGetStopwordListStatusOptions(environmentID string, collectionID string) *GetStopwordListStatusOptions {
	return &GetStopwordListStatusOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *GetStopwordListStatusOptions) SetEnvironmentID(environmentID string) *GetStopwordListStatusOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetCollectionID : Allow user to set CollectionID
func (_options *GetStopwordListStatusOptions) SetCollectionID(collectionID string) *GetStopwordListStatusOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetStopwordListStatusOptions) SetHeaders(param map[string]string) *GetStopwordListStatusOptions {
	options.Headers = param
	return options
}

// GetTokenizationDictionaryStatusOptions : The GetTokenizationDictionaryStatus options.
type GetTokenizationDictionaryStatusOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The ID of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetTokenizationDictionaryStatusOptions : Instantiate GetTokenizationDictionaryStatusOptions
func (*DiscoveryV1) NewGetTokenizationDictionaryStatusOptions(environmentID string, collectionID string) *GetTokenizationDictionaryStatusOptions {
	return &GetTokenizationDictionaryStatusOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *GetTokenizationDictionaryStatusOptions) SetEnvironmentID(environmentID string) *GetTokenizationDictionaryStatusOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetCollectionID : Allow user to set CollectionID
func (_options *GetTokenizationDictionaryStatusOptions) SetCollectionID(collectionID string) *GetTokenizationDictionaryStatusOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetTokenizationDictionaryStatusOptions) SetHeaders(param map[string]string) *GetTokenizationDictionaryStatusOptions {
	options.Headers = param
	return options
}

// GetTrainingDataOptions : The GetTrainingData options.
type GetTrainingDataOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The ID of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// The ID of the query used for training.
	QueryID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetTrainingDataOptions : Instantiate GetTrainingDataOptions
func (*DiscoveryV1) NewGetTrainingDataOptions(environmentID string, collectionID string, queryID string) *GetTrainingDataOptions {
	return &GetTrainingDataOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
		QueryID:       core.StringPtr(queryID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *GetTrainingDataOptions) SetEnvironmentID(environmentID string) *GetTrainingDataOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetCollectionID : Allow user to set CollectionID
func (_options *GetTrainingDataOptions) SetCollectionID(collectionID string) *GetTrainingDataOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetQueryID : Allow user to set QueryID
func (_options *GetTrainingDataOptions) SetQueryID(queryID string) *GetTrainingDataOptions {
	_options.QueryID = core.StringPtr(queryID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetTrainingDataOptions) SetHeaders(param map[string]string) *GetTrainingDataOptions {
	options.Headers = param
	return options
}

// GetTrainingExampleOptions : The GetTrainingExample options.
type GetTrainingExampleOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The ID of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// The ID of the query used for training.
	QueryID *string `json:"-" validate:"required,ne="`

	// The ID of the document as it is indexed.
	ExampleID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetTrainingExampleOptions : Instantiate GetTrainingExampleOptions
func (*DiscoveryV1) NewGetTrainingExampleOptions(environmentID string, collectionID string, queryID string, exampleID string) *GetTrainingExampleOptions {
	return &GetTrainingExampleOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
		QueryID:       core.StringPtr(queryID),
		ExampleID:     core.StringPtr(exampleID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *GetTrainingExampleOptions) SetEnvironmentID(environmentID string) *GetTrainingExampleOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetCollectionID : Allow user to set CollectionID
func (_options *GetTrainingExampleOptions) SetCollectionID(collectionID string) *GetTrainingExampleOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetQueryID : Allow user to set QueryID
func (_options *GetTrainingExampleOptions) SetQueryID(queryID string) *GetTrainingExampleOptions {
	_options.QueryID = core.StringPtr(queryID)
	return _options
}

// SetExampleID : Allow user to set ExampleID
func (_options *GetTrainingExampleOptions) SetExampleID(exampleID string) *GetTrainingExampleOptions {
	_options.ExampleID = core.StringPtr(exampleID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetTrainingExampleOptions) SetHeaders(param map[string]string) *GetTrainingExampleOptions {
	options.Headers = param
	return options
}

// HTMLSettings : A list of HTML conversion settings.
type HTMLSettings struct {
	// Array of HTML tags that are excluded completely.
	ExcludeTagsCompletely []string `json:"exclude_tags_completely,omitempty"`

	// Array of HTML tags which are excluded but still retain content.
	ExcludeTagsKeepContent []string `json:"exclude_tags_keep_content,omitempty"`

	// Object containing an array of XPaths.
	KeepContent *XPathPatterns `json:"keep_content,omitempty"`

	// Object containing an array of XPaths.
	ExcludeContent *XPathPatterns `json:"exclude_content,omitempty"`

	// An array of HTML tag attributes to keep in the converted document.
	KeepTagAttributes []string `json:"keep_tag_attributes,omitempty"`

	// Array of HTML tag attributes to exclude.
	ExcludeTagAttributes []string `json:"exclude_tag_attributes,omitempty"`
}

// UnmarshalHTMLSettings unmarshals an instance of HTMLSettings from the specified map of raw messages.
func UnmarshalHTMLSettings(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(HTMLSettings)
	err = core.UnmarshalPrimitive(m, "exclude_tags_completely", &obj.ExcludeTagsCompletely)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "exclude_tags_keep_content", &obj.ExcludeTagsKeepContent)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "keep_content", &obj.KeepContent, UnmarshalXPathPatterns)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "exclude_content", &obj.ExcludeContent, UnmarshalXPathPatterns)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "keep_tag_attributes", &obj.KeepTagAttributes)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "exclude_tag_attributes", &obj.ExcludeTagAttributes)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// IndexCapacity : Details about the resource usage and capacity of the environment.
type IndexCapacity struct {
	// Summary of the document usage statistics for the environment.
	Documents *EnvironmentDocuments `json:"documents,omitempty"`

	// Summary of the disk usage statistics for the environment.
	DiskUsage *DiskUsage `json:"disk_usage,omitempty"`

	// Summary of the collection usage in the environment.
	Collections *CollectionUsage `json:"collections,omitempty"`
}

// UnmarshalIndexCapacity unmarshals an instance of IndexCapacity from the specified map of raw messages.
func UnmarshalIndexCapacity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(IndexCapacity)
	err = core.UnmarshalModel(m, "documents", &obj.Documents, UnmarshalEnvironmentDocuments)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "disk_usage", &obj.DiskUsage, UnmarshalDiskUsage)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "collections", &obj.Collections, UnmarshalCollectionUsage)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListCollectionFieldsOptions : The ListCollectionFields options.
type ListCollectionFieldsOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The ID of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListCollectionFieldsOptions : Instantiate ListCollectionFieldsOptions
func (*DiscoveryV1) NewListCollectionFieldsOptions(environmentID string, collectionID string) *ListCollectionFieldsOptions {
	return &ListCollectionFieldsOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *ListCollectionFieldsOptions) SetEnvironmentID(environmentID string) *ListCollectionFieldsOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetCollectionID : Allow user to set CollectionID
func (_options *ListCollectionFieldsOptions) SetCollectionID(collectionID string) *ListCollectionFieldsOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListCollectionFieldsOptions) SetHeaders(param map[string]string) *ListCollectionFieldsOptions {
	options.Headers = param
	return options
}

// ListCollectionFieldsResponse : The list of fetched fields.
//
// The fields are returned using a fully qualified name format, however, the format differs slightly from that used by
// the query operations.
//
//   * Fields which contain nested JSON objects are assigned a type of "nested".
//
//   * Fields which belong to a nested object are prefixed with `.properties` (for example,
// `warnings.properties.severity` means that the `warnings` object has a property called `severity`).
//
//   * Fields returned from the News collection are prefixed with `v{N}-fullnews-t3-{YEAR}.mappings` (for example,
// `v5-fullnews-t3-2016.mappings.text.properties.author`).
type ListCollectionFieldsResponse struct {
	// An array containing information about each field in the collections.
	Fields []Field `json:"fields,omitempty"`
}

// UnmarshalListCollectionFieldsResponse unmarshals an instance of ListCollectionFieldsResponse from the specified map of raw messages.
func UnmarshalListCollectionFieldsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListCollectionFieldsResponse)
	err = core.UnmarshalModel(m, "fields", &obj.Fields, UnmarshalField)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListCollectionsOptions : The ListCollections options.
type ListCollectionsOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// Find collections with the given name.
	Name *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListCollectionsOptions : Instantiate ListCollectionsOptions
func (*DiscoveryV1) NewListCollectionsOptions(environmentID string) *ListCollectionsOptions {
	return &ListCollectionsOptions{
		EnvironmentID: core.StringPtr(environmentID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *ListCollectionsOptions) SetEnvironmentID(environmentID string) *ListCollectionsOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetName : Allow user to set Name
func (_options *ListCollectionsOptions) SetName(name string) *ListCollectionsOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListCollectionsOptions) SetHeaders(param map[string]string) *ListCollectionsOptions {
	options.Headers = param
	return options
}

// ListCollectionsResponse : Response object containing an array of collection details.
type ListCollectionsResponse struct {
	// An array containing information about each collection in the environment.
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

// ListConfigurationsOptions : The ListConfigurations options.
type ListConfigurationsOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// Find configurations with the given name.
	Name *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListConfigurationsOptions : Instantiate ListConfigurationsOptions
func (*DiscoveryV1) NewListConfigurationsOptions(environmentID string) *ListConfigurationsOptions {
	return &ListConfigurationsOptions{
		EnvironmentID: core.StringPtr(environmentID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *ListConfigurationsOptions) SetEnvironmentID(environmentID string) *ListConfigurationsOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetName : Allow user to set Name
func (_options *ListConfigurationsOptions) SetName(name string) *ListConfigurationsOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListConfigurationsOptions) SetHeaders(param map[string]string) *ListConfigurationsOptions {
	options.Headers = param
	return options
}

// ListConfigurationsResponse : Object containing an array of available configurations.
type ListConfigurationsResponse struct {
	// An array of configurations that are available for the service instance.
	Configurations []Configuration `json:"configurations,omitempty"`
}

// UnmarshalListConfigurationsResponse unmarshals an instance of ListConfigurationsResponse from the specified map of raw messages.
func UnmarshalListConfigurationsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListConfigurationsResponse)
	err = core.UnmarshalModel(m, "configurations", &obj.Configurations, UnmarshalConfiguration)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListCredentialsOptions : The ListCredentials options.
type ListCredentialsOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListCredentialsOptions : Instantiate ListCredentialsOptions
func (*DiscoveryV1) NewListCredentialsOptions(environmentID string) *ListCredentialsOptions {
	return &ListCredentialsOptions{
		EnvironmentID: core.StringPtr(environmentID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *ListCredentialsOptions) SetEnvironmentID(environmentID string) *ListCredentialsOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListCredentialsOptions) SetHeaders(param map[string]string) *ListCredentialsOptions {
	options.Headers = param
	return options
}

// ListEnvironmentsOptions : The ListEnvironments options.
type ListEnvironmentsOptions struct {
	// Show only the environment with the given name.
	Name *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListEnvironmentsOptions : Instantiate ListEnvironmentsOptions
func (*DiscoveryV1) NewListEnvironmentsOptions() *ListEnvironmentsOptions {
	return &ListEnvironmentsOptions{}
}

// SetName : Allow user to set Name
func (_options *ListEnvironmentsOptions) SetName(name string) *ListEnvironmentsOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListEnvironmentsOptions) SetHeaders(param map[string]string) *ListEnvironmentsOptions {
	options.Headers = param
	return options
}

// ListEnvironmentsResponse : Response object containing an array of configured environments.
type ListEnvironmentsResponse struct {
	// An array of [environments] that are available for the service instance.
	Environments []Environment `json:"environments,omitempty"`
}

// UnmarshalListEnvironmentsResponse unmarshals an instance of ListEnvironmentsResponse from the specified map of raw messages.
func UnmarshalListEnvironmentsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListEnvironmentsResponse)
	err = core.UnmarshalModel(m, "environments", &obj.Environments, UnmarshalEnvironment)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListExpansionsOptions : The ListExpansions options.
type ListExpansionsOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The ID of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListExpansionsOptions : Instantiate ListExpansionsOptions
func (*DiscoveryV1) NewListExpansionsOptions(environmentID string, collectionID string) *ListExpansionsOptions {
	return &ListExpansionsOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *ListExpansionsOptions) SetEnvironmentID(environmentID string) *ListExpansionsOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetCollectionID : Allow user to set CollectionID
func (_options *ListExpansionsOptions) SetCollectionID(collectionID string) *ListExpansionsOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListExpansionsOptions) SetHeaders(param map[string]string) *ListExpansionsOptions {
	options.Headers = param
	return options
}

// ListFieldsOptions : The ListFields options.
type ListFieldsOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// A comma-separated list of collection IDs to be queried against.
	CollectionIds []string `json:"-" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListFieldsOptions : Instantiate ListFieldsOptions
func (*DiscoveryV1) NewListFieldsOptions(environmentID string, collectionIds []string) *ListFieldsOptions {
	return &ListFieldsOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionIds: collectionIds,
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *ListFieldsOptions) SetEnvironmentID(environmentID string) *ListFieldsOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetCollectionIds : Allow user to set CollectionIds
func (_options *ListFieldsOptions) SetCollectionIds(collectionIds []string) *ListFieldsOptions {
	_options.CollectionIds = collectionIds
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListFieldsOptions) SetHeaders(param map[string]string) *ListFieldsOptions {
	options.Headers = param
	return options
}

// ListGatewaysOptions : The ListGateways options.
type ListGatewaysOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListGatewaysOptions : Instantiate ListGatewaysOptions
func (*DiscoveryV1) NewListGatewaysOptions(environmentID string) *ListGatewaysOptions {
	return &ListGatewaysOptions{
		EnvironmentID: core.StringPtr(environmentID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *ListGatewaysOptions) SetEnvironmentID(environmentID string) *ListGatewaysOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListGatewaysOptions) SetHeaders(param map[string]string) *ListGatewaysOptions {
	options.Headers = param
	return options
}

// ListTrainingDataOptions : The ListTrainingData options.
type ListTrainingDataOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The ID of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListTrainingDataOptions : Instantiate ListTrainingDataOptions
func (*DiscoveryV1) NewListTrainingDataOptions(environmentID string, collectionID string) *ListTrainingDataOptions {
	return &ListTrainingDataOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *ListTrainingDataOptions) SetEnvironmentID(environmentID string) *ListTrainingDataOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetCollectionID : Allow user to set CollectionID
func (_options *ListTrainingDataOptions) SetCollectionID(collectionID string) *ListTrainingDataOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListTrainingDataOptions) SetHeaders(param map[string]string) *ListTrainingDataOptions {
	options.Headers = param
	return options
}

// ListTrainingExamplesOptions : The ListTrainingExamples options.
type ListTrainingExamplesOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The ID of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// The ID of the query used for training.
	QueryID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListTrainingExamplesOptions : Instantiate ListTrainingExamplesOptions
func (*DiscoveryV1) NewListTrainingExamplesOptions(environmentID string, collectionID string, queryID string) *ListTrainingExamplesOptions {
	return &ListTrainingExamplesOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
		QueryID:       core.StringPtr(queryID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *ListTrainingExamplesOptions) SetEnvironmentID(environmentID string) *ListTrainingExamplesOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetCollectionID : Allow user to set CollectionID
func (_options *ListTrainingExamplesOptions) SetCollectionID(collectionID string) *ListTrainingExamplesOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetQueryID : Allow user to set QueryID
func (_options *ListTrainingExamplesOptions) SetQueryID(queryID string) *ListTrainingExamplesOptions {
	_options.QueryID = core.StringPtr(queryID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListTrainingExamplesOptions) SetHeaders(param map[string]string) *ListTrainingExamplesOptions {
	options.Headers = param
	return options
}

// LogQueryResponse : Object containing results that match the requested **logs** query.
type LogQueryResponse struct {
	// Number of matching results.
	MatchingResults *int64 `json:"matching_results,omitempty"`

	// Array of log query response results.
	Results []LogQueryResponseResult `json:"results,omitempty"`
}

// UnmarshalLogQueryResponse unmarshals an instance of LogQueryResponse from the specified map of raw messages.
func UnmarshalLogQueryResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LogQueryResponse)
	err = core.UnmarshalPrimitive(m, "matching_results", &obj.MatchingResults)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "results", &obj.Results, UnmarshalLogQueryResponseResult)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LogQueryResponseResult : Individual result object for a **logs** query. Each object represents either a query to a Discovery collection or an
// event that is associated with a query.
type LogQueryResponseResult struct {
	// The environment ID that is associated with this log entry.
	EnvironmentID *string `json:"environment_id,omitempty"`

	// The **customer_id** label that was specified in the header of the query or event API call that corresponds to this
	// log entry.
	CustomerID *string `json:"customer_id,omitempty"`

	// The type of log entry returned.
	//
	//  **query** indicates that the log represents the results of a call to the single collection **query** method.
	//
	//  **event** indicates that the log represents  a call to the **events** API.
	DocumentType *string `json:"document_type,omitempty"`

	// The value of the **natural_language_query** query parameter that was used to create these results. Only returned
	// with logs of type **query**.
	//
	// **Note:** Other query parameters (such as **filter** or **deduplicate**) might  have been used with this query, but
	// are not recorded.
	NaturalLanguageQuery *string `json:"natural_language_query,omitempty"`

	// Object containing result information that was returned by the query used to create this log entry. Only returned
	// with logs of type `query`.
	DocumentResults *LogQueryResponseResultDocuments `json:"document_results,omitempty"`

	// Date that the log result was created. Returned in `YYYY-MM-DDThh:mm:ssZ` format.
	CreatedTimestamp *strfmt.DateTime `json:"created_timestamp,omitempty"`

	// Date specified by the user when recording an event. Returned in `YYYY-MM-DDThh:mm:ssZ` format. Only returned with
	// logs of type **event**.
	ClientTimestamp *strfmt.DateTime `json:"client_timestamp,omitempty"`

	// Identifier that corresponds to the **natural_language_query** string used in the original or associated query. All
	// **event** and **query** log entries that have the same original **natural_language_query** string also have them
	// same **query_id**. This field can be used to recall all **event** and **query** log results that have the same
	// original query (**event** logs do not contain the original **natural_language_query** field).
	QueryID *string `json:"query_id,omitempty"`

	// Unique identifier (within a 24-hour period) that identifies a single `query` log and any `event` logs that were
	// created for it.
	//
	// **Note:** If the exact same query is run at the exact same time on different days, the **session_token** for those
	// queries might be identical. However, the **created_timestamp** differs.
	//
	// **Note:** Session tokens are case sensitive. To avoid matching on session tokens that are identical except for case,
	// use the exact match operator (`::`) when you query for a specific session token.
	SessionToken *string `json:"session_token,omitempty"`

	// The collection ID of the document associated with this event. Only returned with logs of type `event`.
	CollectionID *string `json:"collection_id,omitempty"`

	// The original display rank of the document associated with this event. Only returned with logs of type `event`.
	DisplayRank *int64 `json:"display_rank,omitempty"`

	// The document ID of the document associated with this event. Only returned with logs of type `event`.
	DocumentID *string `json:"document_id,omitempty"`

	// The type of event that this object respresents. Possible values are
	//
	//  -  `query` the log of a query to a collection
	//
	//  -  `click` the result of a call to the **events** endpoint.
	EventType *string `json:"event_type,omitempty"`

	// The type of result that this **event** is associated with. Only returned with logs of type `event`.
	ResultType *string `json:"result_type,omitempty"`
}

// Constants associated with the LogQueryResponseResult.DocumentType property.
// The type of log entry returned.
//
//  **query** indicates that the log represents the results of a call to the single collection **query** method.
//
//  **event** indicates that the log represents  a call to the **events** API.
const (
	LogQueryResponseResultDocumentTypeEventConst = "event"
	LogQueryResponseResultDocumentTypeQueryConst = "query"
)

// Constants associated with the LogQueryResponseResult.EventType property.
// The type of event that this object respresents. Possible values are
//
//  -  `query` the log of a query to a collection
//
//  -  `click` the result of a call to the **events** endpoint.
const (
	LogQueryResponseResultEventTypeClickConst = "click"
	LogQueryResponseResultEventTypeQueryConst = "query"
)

// Constants associated with the LogQueryResponseResult.ResultType property.
// The type of result that this **event** is associated with. Only returned with logs of type `event`.
const (
	LogQueryResponseResultResultTypeDocumentConst = "document"
)

// UnmarshalLogQueryResponseResult unmarshals an instance of LogQueryResponseResult from the specified map of raw messages.
func UnmarshalLogQueryResponseResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LogQueryResponseResult)
	err = core.UnmarshalPrimitive(m, "environment_id", &obj.EnvironmentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "customer_id", &obj.CustomerID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "document_type", &obj.DocumentType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "natural_language_query", &obj.NaturalLanguageQuery)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "document_results", &obj.DocumentResults, UnmarshalLogQueryResponseResultDocuments)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_timestamp", &obj.CreatedTimestamp)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "client_timestamp", &obj.ClientTimestamp)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "query_id", &obj.QueryID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "session_token", &obj.SessionToken)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "collection_id", &obj.CollectionID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "display_rank", &obj.DisplayRank)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "document_id", &obj.DocumentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "event_type", &obj.EventType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "result_type", &obj.ResultType)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LogQueryResponseResultDocuments : Object containing result information that was returned by the query used to create this log entry. Only returned with
// logs of type `query`.
type LogQueryResponseResultDocuments struct {
	// Array of log query response results.
	Results []LogQueryResponseResultDocumentsResult `json:"results,omitempty"`

	// The number of results returned in the query associate with this log.
	Count *int64 `json:"count,omitempty"`
}

// UnmarshalLogQueryResponseResultDocuments unmarshals an instance of LogQueryResponseResultDocuments from the specified map of raw messages.
func UnmarshalLogQueryResponseResultDocuments(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LogQueryResponseResultDocuments)
	err = core.UnmarshalModel(m, "results", &obj.Results, UnmarshalLogQueryResponseResultDocumentsResult)
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

// LogQueryResponseResultDocumentsResult : Each object in the **results** array corresponds to an individual document returned by the original query.
type LogQueryResponseResultDocumentsResult struct {
	// The result rank of this document. A position of `1` indicates that it was the first returned result.
	Position *int64 `json:"position,omitempty"`

	// The **document_id** of the document that this result represents.
	DocumentID *string `json:"document_id,omitempty"`

	// The raw score of this result. A higher score indicates a greater match to the query parameters.
	Score *float64 `json:"score,omitempty"`

	// The confidence score of the result's analysis. A higher score indicating greater confidence.
	Confidence *float64 `json:"confidence,omitempty"`

	// The **collection_id** of the document represented by this result.
	CollectionID *string `json:"collection_id,omitempty"`
}

// UnmarshalLogQueryResponseResultDocumentsResult unmarshals an instance of LogQueryResponseResultDocumentsResult from the specified map of raw messages.
func UnmarshalLogQueryResponseResultDocumentsResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LogQueryResponseResultDocumentsResult)
	err = core.UnmarshalPrimitive(m, "position", &obj.Position)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "document_id", &obj.DocumentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "score", &obj.Score)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "confidence", &obj.Confidence)
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

// MetricAggregation : An aggregation analyzing log information for queries and events.
type MetricAggregation struct {
	// The measurement interval for this metric. Metric intervals are always 1 day (`1d`).
	Interval *string `json:"interval,omitempty"`

	// The event type associated with this metric result. This field, when present, will always be `click`.
	EventType *string `json:"event_type,omitempty"`

	// Array of metric aggregation query results.
	Results []MetricAggregationResult `json:"results,omitempty"`
}

// UnmarshalMetricAggregation unmarshals an instance of MetricAggregation from the specified map of raw messages.
func UnmarshalMetricAggregation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MetricAggregation)
	err = core.UnmarshalPrimitive(m, "interval", &obj.Interval)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "event_type", &obj.EventType)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "results", &obj.Results, UnmarshalMetricAggregationResult)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MetricAggregationResult : Aggregation result data for the requested metric.
type MetricAggregationResult struct {
	// Date in string form representing the start of this interval.
	KeyAsString *strfmt.DateTime `json:"key_as_string,omitempty"`

	// Unix epoch time equivalent of the **key_as_string**, that represents the start of this interval.
	Key *int64 `json:"key,omitempty"`

	// Number of matching results.
	MatchingResults *int64 `json:"matching_results,omitempty"`

	// The number of queries with associated events divided by the total number of queries for the interval. Only returned
	// with **event_rate** metrics.
	EventRate *float64 `json:"event_rate,omitempty"`
}

// UnmarshalMetricAggregationResult unmarshals an instance of MetricAggregationResult from the specified map of raw messages.
func UnmarshalMetricAggregationResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MetricAggregationResult)
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
	err = core.UnmarshalPrimitive(m, "event_rate", &obj.EventRate)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MetricResponse : The response generated from a call to a **metrics** method.
type MetricResponse struct {
	// Array of metric aggregations.
	Aggregations []MetricAggregation `json:"aggregations,omitempty"`
}

// UnmarshalMetricResponse unmarshals an instance of MetricResponse from the specified map of raw messages.
func UnmarshalMetricResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MetricResponse)
	err = core.UnmarshalModel(m, "aggregations", &obj.Aggregations, UnmarshalMetricAggregation)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MetricTokenAggregation : An aggregation analyzing log information for queries and events.
type MetricTokenAggregation struct {
	// The event type associated with this metric result. This field, when present, will always be `click`.
	EventType *string `json:"event_type,omitempty"`

	// Array of results for the metric token aggregation.
	Results []MetricTokenAggregationResult `json:"results,omitempty"`
}

// UnmarshalMetricTokenAggregation unmarshals an instance of MetricTokenAggregation from the specified map of raw messages.
func UnmarshalMetricTokenAggregation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MetricTokenAggregation)
	err = core.UnmarshalPrimitive(m, "event_type", &obj.EventType)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "results", &obj.Results, UnmarshalMetricTokenAggregationResult)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MetricTokenAggregationResult : Aggregation result data for the requested metric.
type MetricTokenAggregationResult struct {
	// The content of the **natural_language_query** parameter used in the query that this result represents.
	Key *string `json:"key,omitempty"`

	// Number of matching results.
	MatchingResults *int64 `json:"matching_results,omitempty"`

	// The number of queries with associated events divided by the total number of queries currently stored (queries and
	// events are stored in the log for 30 days).
	EventRate *float64 `json:"event_rate,omitempty"`
}

// UnmarshalMetricTokenAggregationResult unmarshals an instance of MetricTokenAggregationResult from the specified map of raw messages.
func UnmarshalMetricTokenAggregationResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MetricTokenAggregationResult)
	err = core.UnmarshalPrimitive(m, "key", &obj.Key)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "matching_results", &obj.MatchingResults)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "event_rate", &obj.EventRate)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MetricTokenResponse : The response generated from a call to a **metrics** method that evaluates tokens.
type MetricTokenResponse struct {
	// Array of metric token aggregations.
	Aggregations []MetricTokenAggregation `json:"aggregations,omitempty"`
}

// UnmarshalMetricTokenResponse unmarshals an instance of MetricTokenResponse from the specified map of raw messages.
func UnmarshalMetricTokenResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MetricTokenResponse)
	err = core.UnmarshalModel(m, "aggregations", &obj.Aggregations, UnmarshalMetricTokenAggregation)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// NluEnrichmentConcepts : An object specifiying the concepts enrichment and related parameters.
type NluEnrichmentConcepts struct {
	// The maximum number of concepts enrichments to extact from each instance of the specified field.
	Limit *int64 `json:"limit,omitempty"`
}

// UnmarshalNluEnrichmentConcepts unmarshals an instance of NluEnrichmentConcepts from the specified map of raw messages.
func UnmarshalNluEnrichmentConcepts(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NluEnrichmentConcepts)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// NluEnrichmentEmotion : An object specifying the emotion detection enrichment and related parameters.
type NluEnrichmentEmotion struct {
	// When `true`, emotion detection is performed on the entire field.
	Document *bool `json:"document,omitempty"`

	// A comma-separated list of target strings that will have any associated emotions detected.
	Targets []string `json:"targets,omitempty"`
}

// UnmarshalNluEnrichmentEmotion unmarshals an instance of NluEnrichmentEmotion from the specified map of raw messages.
func UnmarshalNluEnrichmentEmotion(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NluEnrichmentEmotion)
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

// NluEnrichmentEntities : An object speficying the Entities enrichment and related parameters.
type NluEnrichmentEntities struct {
	// When `true`, sentiment analysis of entities will be performed on the specified field.
	Sentiment *bool `json:"sentiment,omitempty"`

	// When `true`, emotion detection of entities will be performed on the specified field.
	Emotion *bool `json:"emotion,omitempty"`

	// The maximum number of entities to extract for each instance of the specified field.
	Limit *int64 `json:"limit,omitempty"`

	// When `true`, the number of mentions of each identified entity is recorded. The default is `false`.
	Mentions *bool `json:"mentions,omitempty"`

	// When `true`, the types of mentions for each idetifieid entity is recorded. The default is `false`.
	MentionTypes *bool `json:"mention_types,omitempty"`

	// When `true`, a list of sentence locations for each instance of each identified entity is recorded. The default is
	// `false`.
	SentenceLocations *bool `json:"sentence_locations,omitempty"`

	// The enrichement model to use with entity extraction. May be a custom model provided by Watson Knowledge Studio, or
	// the default public model `alchemy`.
	Model *string `json:"model,omitempty"`
}

// UnmarshalNluEnrichmentEntities unmarshals an instance of NluEnrichmentEntities from the specified map of raw messages.
func UnmarshalNluEnrichmentEntities(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NluEnrichmentEntities)
	err = core.UnmarshalPrimitive(m, "sentiment", &obj.Sentiment)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "emotion", &obj.Emotion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "mentions", &obj.Mentions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "mention_types", &obj.MentionTypes)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sentence_locations", &obj.SentenceLocations)
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

// NluEnrichmentFeatures : Object containing Natural Language Understanding features to be used.
type NluEnrichmentFeatures struct {
	// An object specifying the Keyword enrichment and related parameters.
	Keywords *NluEnrichmentKeywords `json:"keywords,omitempty"`

	// An object speficying the Entities enrichment and related parameters.
	Entities *NluEnrichmentEntities `json:"entities,omitempty"`

	// An object specifying the sentiment extraction enrichment and related parameters.
	Sentiment *NluEnrichmentSentiment `json:"sentiment,omitempty"`

	// An object specifying the emotion detection enrichment and related parameters.
	Emotion *NluEnrichmentEmotion `json:"emotion,omitempty"`

	// An object that indicates the Categories enrichment will be applied to the specified field.
	Categories map[string]interface{} `json:"categories,omitempty"`

	// An object specifiying the semantic roles enrichment and related parameters.
	SemanticRoles *NluEnrichmentSemanticRoles `json:"semantic_roles,omitempty"`

	// An object specifying the relations enrichment and related parameters.
	Relations *NluEnrichmentRelations `json:"relations,omitempty"`

	// An object specifiying the concepts enrichment and related parameters.
	Concepts *NluEnrichmentConcepts `json:"concepts,omitempty"`
}

// UnmarshalNluEnrichmentFeatures unmarshals an instance of NluEnrichmentFeatures from the specified map of raw messages.
func UnmarshalNluEnrichmentFeatures(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NluEnrichmentFeatures)
	err = core.UnmarshalModel(m, "keywords", &obj.Keywords, UnmarshalNluEnrichmentKeywords)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "entities", &obj.Entities, UnmarshalNluEnrichmentEntities)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "sentiment", &obj.Sentiment, UnmarshalNluEnrichmentSentiment)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "emotion", &obj.Emotion, UnmarshalNluEnrichmentEmotion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "categories", &obj.Categories)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "semantic_roles", &obj.SemanticRoles, UnmarshalNluEnrichmentSemanticRoles)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "relations", &obj.Relations, UnmarshalNluEnrichmentRelations)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "concepts", &obj.Concepts, UnmarshalNluEnrichmentConcepts)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// NluEnrichmentKeywords : An object specifying the Keyword enrichment and related parameters.
type NluEnrichmentKeywords struct {
	// When `true`, sentiment analysis of keywords will be performed on the specified field.
	Sentiment *bool `json:"sentiment,omitempty"`

	// When `true`, emotion detection of keywords will be performed on the specified field.
	Emotion *bool `json:"emotion,omitempty"`

	// The maximum number of keywords to extract for each instance of the specified field.
	Limit *int64 `json:"limit,omitempty"`
}

// UnmarshalNluEnrichmentKeywords unmarshals an instance of NluEnrichmentKeywords from the specified map of raw messages.
func UnmarshalNluEnrichmentKeywords(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NluEnrichmentKeywords)
	err = core.UnmarshalPrimitive(m, "sentiment", &obj.Sentiment)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "emotion", &obj.Emotion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// NluEnrichmentRelations : An object specifying the relations enrichment and related parameters.
type NluEnrichmentRelations struct {
	// *For use with `natural_language_understanding` enrichments only.* The enrichement model to use with relationship
	// extraction. May be a custom model provided by Watson Knowledge Studio, the default public model is`en-news`.
	Model *string `json:"model,omitempty"`
}

// UnmarshalNluEnrichmentRelations unmarshals an instance of NluEnrichmentRelations from the specified map of raw messages.
func UnmarshalNluEnrichmentRelations(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NluEnrichmentRelations)
	err = core.UnmarshalPrimitive(m, "model", &obj.Model)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// NluEnrichmentSemanticRoles : An object specifiying the semantic roles enrichment and related parameters.
type NluEnrichmentSemanticRoles struct {
	// When `true`, entities are extracted from the identified sentence parts.
	Entities *bool `json:"entities,omitempty"`

	// When `true`, keywords are extracted from the identified sentence parts.
	Keywords *bool `json:"keywords,omitempty"`

	// The maximum number of semantic roles enrichments to extact from each instance of the specified field.
	Limit *int64 `json:"limit,omitempty"`
}

// UnmarshalNluEnrichmentSemanticRoles unmarshals an instance of NluEnrichmentSemanticRoles from the specified map of raw messages.
func UnmarshalNluEnrichmentSemanticRoles(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NluEnrichmentSemanticRoles)
	err = core.UnmarshalPrimitive(m, "entities", &obj.Entities)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "keywords", &obj.Keywords)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// NluEnrichmentSentiment : An object specifying the sentiment extraction enrichment and related parameters.
type NluEnrichmentSentiment struct {
	// When `true`, sentiment analysis is performed on the entire field.
	Document *bool `json:"document,omitempty"`

	// A comma-separated list of target strings that will have any associated sentiment analyzed.
	Targets []string `json:"targets,omitempty"`
}

// UnmarshalNluEnrichmentSentiment unmarshals an instance of NluEnrichmentSentiment from the specified map of raw messages.
func UnmarshalNluEnrichmentSentiment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NluEnrichmentSentiment)
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

// NormalizationOperation : Object containing normalization operations.
type NormalizationOperation struct {
	// Identifies what type of operation to perform.
	//
	// **copy** - Copies the value of the **source_field** to the **destination_field** field. If the **destination_field**
	// already exists, then the value of the **source_field** overwrites the original value of the **destination_field**.
	//
	// **move** - Renames (moves) the **source_field** to the **destination_field**. If the **destination_field** already
	// exists, then the value of the **source_field** overwrites the original value of the **destination_field**. Rename is
	// identical to copy, except that the **source_field** is removed after the value has been copied to the
	// **destination_field** (it is the same as a _copy_ followed by a _remove_).
	//
	// **merge** - Merges the value of the **source_field** with the value of the **destination_field**. The
	// **destination_field** is converted into an array if it is not already an array, and the value of the
	// **source_field** is appended to the array. This operation removes the **source_field** after the merge. If the
	// **source_field** does not exist in the current document, then the **destination_field** is still converted into an
	// array (if it is not an array already). This conversion ensures the type for **destination_field** is consistent
	// across all documents.
	//
	// **remove** - Deletes the **source_field** field. The **destination_field** is ignored for this operation.
	//
	// **remove_nulls** - Removes all nested null (blank) field values from the ingested document. **source_field** and
	// **destination_field** are ignored by this operation because _remove_nulls_ operates on the entire ingested document.
	// Typically, **remove_nulls** is invoked as the last normalization operation (if it is invoked at all, it can be
	// time-expensive).
	Operation *string `json:"operation,omitempty"`

	// The source field for the operation.
	SourceField *string `json:"source_field,omitempty"`

	// The destination field for the operation.
	DestinationField *string `json:"destination_field,omitempty"`
}

// Constants associated with the NormalizationOperation.Operation property.
// Identifies what type of operation to perform.
//
// **copy** - Copies the value of the **source_field** to the **destination_field** field. If the **destination_field**
// already exists, then the value of the **source_field** overwrites the original value of the **destination_field**.
//
// **move** - Renames (moves) the **source_field** to the **destination_field**. If the **destination_field** already
// exists, then the value of the **source_field** overwrites the original value of the **destination_field**. Rename is
// identical to copy, except that the **source_field** is removed after the value has been copied to the
// **destination_field** (it is the same as a _copy_ followed by a _remove_).
//
// **merge** - Merges the value of the **source_field** with the value of the **destination_field**. The
// **destination_field** is converted into an array if it is not already an array, and the value of the **source_field**
// is appended to the array. This operation removes the **source_field** after the merge. If the **source_field** does
// not exist in the current document, then the **destination_field** is still converted into an array (if it is not an
// array already). This conversion ensures the type for **destination_field** is consistent across all documents.
//
// **remove** - Deletes the **source_field** field. The **destination_field** is ignored for this operation.
//
// **remove_nulls** - Removes all nested null (blank) field values from the ingested document. **source_field** and
// **destination_field** are ignored by this operation because _remove_nulls_ operates on the entire ingested document.
// Typically, **remove_nulls** is invoked as the last normalization operation (if it is invoked at all, it can be
// time-expensive).
const (
	NormalizationOperationOperationCopyConst        = "copy"
	NormalizationOperationOperationMergeConst       = "merge"
	NormalizationOperationOperationMoveConst        = "move"
	NormalizationOperationOperationRemoveConst      = "remove"
	NormalizationOperationOperationRemoveNullsConst = "remove_nulls"
)

// UnmarshalNormalizationOperation unmarshals an instance of NormalizationOperation from the specified map of raw messages.
func UnmarshalNormalizationOperation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NormalizationOperation)
	err = core.UnmarshalPrimitive(m, "operation", &obj.Operation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source_field", &obj.SourceField)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "destination_field", &obj.DestinationField)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
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
	// complete list; other values might be returned.
	NoticeID *string `json:"notice_id,omitempty"`

	// The creation date of the collection in the format yyyy-MM-dd'T'HH:mm:ss.SSS'Z'.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// Unique identifier of the document.
	DocumentID *string `json:"document_id,omitempty"`

	// Unique identifier of the query used for relevance training.
	QueryID *string `json:"query_id,omitempty"`

	// Severity level of the notice.
	Severity *string `json:"severity,omitempty"`

	// Ingestion or training step in which the notice occurred. Typical step values include: `smartDocumentUnderstanding`,
	// `ingestion`, `indexing`, `convert`. **Note:** This is not a complete list; other values might be returned.
	Step *string `json:"step,omitempty"`

	// The description of the notice.
	Description *string `json:"description,omitempty"`
}

// Constants associated with the Notice.Severity property.
// Severity level of the notice.
const (
	NoticeSeverityErrorConst   = "error"
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

// PDFHeadingDetection : Object containing heading detection conversion settings for PDF documents.
type PDFHeadingDetection struct {
	// Array of font matching configurations.
	Fonts []FontSetting `json:"fonts,omitempty"`
}

// UnmarshalPDFHeadingDetection unmarshals an instance of PDFHeadingDetection from the specified map of raw messages.
func UnmarshalPDFHeadingDetection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PDFHeadingDetection)
	err = core.UnmarshalModel(m, "fonts", &obj.Fonts, UnmarshalFontSetting)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PDFSettings : A list of PDF conversion settings.
type PDFSettings struct {
	// Object containing heading detection conversion settings for PDF documents.
	Heading *PDFHeadingDetection `json:"heading,omitempty"`
}

// UnmarshalPDFSettings unmarshals an instance of PDFSettings from the specified map of raw messages.
func UnmarshalPDFSettings(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PDFSettings)
	err = core.UnmarshalModel(m, "heading", &obj.Heading, UnmarshalPDFHeadingDetection)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// QueryAggregation : An aggregation produced by  Discovery to analyze the input provided.
type QueryAggregation struct {
	// The type of aggregation command used. For example: term, filter, max, min, etc.
	Type *string `json:"type,omitempty"`

	// Array of aggregation results.
	Results []AggregationResult `json:"results,omitempty"`

	// Number of matching results.
	MatchingResults *int64 `json:"matching_results,omitempty"`

	// Aggregations returned by Discovery.
	Aggregations []QueryAggregationIntf `json:"aggregations,omitempty"`
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
	if discValue == "histogram" {
		err = core.UnmarshalModel(m, "", result, UnmarshalHistogram)
	} else if discValue == "max" {
		err = core.UnmarshalModel(m, "", result, UnmarshalCalculation)
	} else if discValue == "min" {
		err = core.UnmarshalModel(m, "", result, UnmarshalCalculation)
	} else if discValue == "average" {
		err = core.UnmarshalModel(m, "", result, UnmarshalCalculation)
	} else if discValue == "sum" {
		err = core.UnmarshalModel(m, "", result, UnmarshalCalculation)
	} else if discValue == "unique_count" {
		err = core.UnmarshalModel(m, "", result, UnmarshalCalculation)
	} else if discValue == "term" {
		err = core.UnmarshalModel(m, "", result, UnmarshalTerm)
	} else if discValue == "filter" {
		err = core.UnmarshalModel(m, "", result, UnmarshalFilter)
	} else if discValue == "nested" {
		err = core.UnmarshalModel(m, "", result, UnmarshalNested)
	} else if discValue == "timeslice" {
		err = core.UnmarshalModel(m, "", result, UnmarshalTimeslice)
	} else if discValue == "top_hits" {
		err = core.UnmarshalModel(m, "", result, UnmarshalTopHits)
	} else {
		err = fmt.Errorf("unrecognized value for discriminator property 'type': %s", discValue)
	}
	return
}

// QueryLogOptions : The QueryLog options.
type QueryLogOptions struct {
	// A cacheable query that excludes documents that don't mention the query content. Filter searches are better for
	// metadata-type searches and for assessing the concepts in the data set.
	Filter *string `json:"-"`

	// A query search returns all documents in your data set with full enrichments and full text, but with the most
	// relevant documents listed first.
	Query *string `json:"-"`

	// Number of results to return. The maximum for the **count** and **offset** values together in any one query is
	// **10000**.
	Count *int64 `json:"-"`

	// The number of query results to skip at the beginning. For example, if the total number of results that are returned
	// is 10 and the offset is 8, it returns the last two results. The maximum for the **count** and **offset** values
	// together in any one query is **10000**.
	Offset *int64 `json:"-"`

	// A comma-separated list of fields in the document to sort on. You can optionally specify a sort direction by
	// prefixing the field with `-` for descending or `+` for ascending. Ascending is the default sort direction if no
	// prefix is specified.
	Sort []string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewQueryLogOptions : Instantiate QueryLogOptions
func (*DiscoveryV1) NewQueryLogOptions() *QueryLogOptions {
	return &QueryLogOptions{}
}

// SetFilter : Allow user to set Filter
func (_options *QueryLogOptions) SetFilter(filter string) *QueryLogOptions {
	_options.Filter = core.StringPtr(filter)
	return _options
}

// SetQuery : Allow user to set Query
func (_options *QueryLogOptions) SetQuery(query string) *QueryLogOptions {
	_options.Query = core.StringPtr(query)
	return _options
}

// SetCount : Allow user to set Count
func (_options *QueryLogOptions) SetCount(count int64) *QueryLogOptions {
	_options.Count = core.Int64Ptr(count)
	return _options
}

// SetOffset : Allow user to set Offset
func (_options *QueryLogOptions) SetOffset(offset int64) *QueryLogOptions {
	_options.Offset = core.Int64Ptr(offset)
	return _options
}

// SetSort : Allow user to set Sort
func (_options *QueryLogOptions) SetSort(sort []string) *QueryLogOptions {
	_options.Sort = sort
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *QueryLogOptions) SetHeaders(param map[string]string) *QueryLogOptions {
	options.Headers = param
	return options
}

// QueryNoticesOptions : The QueryNotices options.
type QueryNoticesOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The ID of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// A cacheable query that excludes documents that don't mention the query content. Filter searches are better for
	// metadata-type searches and for assessing the concepts in the data set.
	Filter *string `json:"-"`

	// A query search returns all documents in your data set with full enrichments and full text, but with the most
	// relevant documents listed first.
	Query *string `json:"-"`

	// A natural language query that returns relevant documents by utilizing training data and natural language
	// understanding.
	NaturalLanguageQuery *string `json:"-"`

	// A passages query that returns the most relevant passages from the results.
	Passages *bool `json:"-"`

	// An aggregation search that returns an exact answer by combining query search with filters. Useful for applications
	// to build lists, tables, and time series. For a full list of possible aggregations, see the Query reference.
	Aggregation *string `json:"-"`

	// Number of results to return. The maximum for the **count** and **offset** values together in any one query is
	// **10000**.
	Count *int64 `json:"-"`

	// A comma-separated list of the portion of the document hierarchy to return.
	Return []string `json:"-"`

	// The number of query results to skip at the beginning. For example, if the total number of results that are returned
	// is 10 and the offset is 8, it returns the last two results. The maximum for the **count** and **offset** values
	// together in any one query is **10000**.
	Offset *int64 `json:"-"`

	// A comma-separated list of fields in the document to sort on. You can optionally specify a sort direction by
	// prefixing the field with `-` for descending or `+` for ascending. Ascending is the default sort direction if no
	// prefix is specified.
	Sort []string `json:"-"`

	// When true, a highlight field is returned for each result which contains the fields which match the query with
	// `<em></em>` tags around the matching query terms.
	Highlight *bool `json:"-"`

	// A comma-separated list of fields that passages are drawn from. If this parameter not specified, then all top-level
	// fields are included.
	PassagesFields []string `json:"-"`

	// The maximum number of passages to return. The search returns fewer passages if the requested total is not found.
	PassagesCount *int64 `json:"-"`

	// The approximate number of characters that any one passage will have.
	PassagesCharacters *int64 `json:"-"`

	// When specified, duplicate results based on the field specified are removed from the returned results. Duplicate
	// comparison is limited to the current query only, **offset** is not considered. This parameter is currently Beta
	// functionality.
	DeduplicateField *string `json:"-"`

	// When `true`, results are returned based on their similarity to the document IDs specified in the
	// **similar.document_ids** parameter.
	Similar *bool `json:"-"`

	// A comma-separated list of document IDs to find similar documents.
	//
	// **Tip:** Include the **natural_language_query** parameter to expand the scope of the document similarity search with
	// the natural language query. Other query parameters, such as **filter** and **query**, are subsequently applied and
	// reduce the scope.
	SimilarDocumentIds []string `json:"-"`

	// A comma-separated list of field names that are used as a basis for comparison to identify similar documents. If not
	// specified, the entire document is used for comparison.
	SimilarFields []string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewQueryNoticesOptions : Instantiate QueryNoticesOptions
func (*DiscoveryV1) NewQueryNoticesOptions(environmentID string, collectionID string) *QueryNoticesOptions {
	return &QueryNoticesOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *QueryNoticesOptions) SetEnvironmentID(environmentID string) *QueryNoticesOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetCollectionID : Allow user to set CollectionID
func (_options *QueryNoticesOptions) SetCollectionID(collectionID string) *QueryNoticesOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetFilter : Allow user to set Filter
func (_options *QueryNoticesOptions) SetFilter(filter string) *QueryNoticesOptions {
	_options.Filter = core.StringPtr(filter)
	return _options
}

// SetQuery : Allow user to set Query
func (_options *QueryNoticesOptions) SetQuery(query string) *QueryNoticesOptions {
	_options.Query = core.StringPtr(query)
	return _options
}

// SetNaturalLanguageQuery : Allow user to set NaturalLanguageQuery
func (_options *QueryNoticesOptions) SetNaturalLanguageQuery(naturalLanguageQuery string) *QueryNoticesOptions {
	_options.NaturalLanguageQuery = core.StringPtr(naturalLanguageQuery)
	return _options
}

// SetPassages : Allow user to set Passages
func (_options *QueryNoticesOptions) SetPassages(passages bool) *QueryNoticesOptions {
	_options.Passages = core.BoolPtr(passages)
	return _options
}

// SetAggregation : Allow user to set Aggregation
func (_options *QueryNoticesOptions) SetAggregation(aggregation string) *QueryNoticesOptions {
	_options.Aggregation = core.StringPtr(aggregation)
	return _options
}

// SetCount : Allow user to set Count
func (_options *QueryNoticesOptions) SetCount(count int64) *QueryNoticesOptions {
	_options.Count = core.Int64Ptr(count)
	return _options
}

// SetReturn : Allow user to set Return
func (_options *QueryNoticesOptions) SetReturn(returnVar []string) *QueryNoticesOptions {
	_options.Return = returnVar
	return _options
}

// SetOffset : Allow user to set Offset
func (_options *QueryNoticesOptions) SetOffset(offset int64) *QueryNoticesOptions {
	_options.Offset = core.Int64Ptr(offset)
	return _options
}

// SetSort : Allow user to set Sort
func (_options *QueryNoticesOptions) SetSort(sort []string) *QueryNoticesOptions {
	_options.Sort = sort
	return _options
}

// SetHighlight : Allow user to set Highlight
func (_options *QueryNoticesOptions) SetHighlight(highlight bool) *QueryNoticesOptions {
	_options.Highlight = core.BoolPtr(highlight)
	return _options
}

// SetPassagesFields : Allow user to set PassagesFields
func (_options *QueryNoticesOptions) SetPassagesFields(passagesFields []string) *QueryNoticesOptions {
	_options.PassagesFields = passagesFields
	return _options
}

// SetPassagesCount : Allow user to set PassagesCount
func (_options *QueryNoticesOptions) SetPassagesCount(passagesCount int64) *QueryNoticesOptions {
	_options.PassagesCount = core.Int64Ptr(passagesCount)
	return _options
}

// SetPassagesCharacters : Allow user to set PassagesCharacters
func (_options *QueryNoticesOptions) SetPassagesCharacters(passagesCharacters int64) *QueryNoticesOptions {
	_options.PassagesCharacters = core.Int64Ptr(passagesCharacters)
	return _options
}

// SetDeduplicateField : Allow user to set DeduplicateField
func (_options *QueryNoticesOptions) SetDeduplicateField(deduplicateField string) *QueryNoticesOptions {
	_options.DeduplicateField = core.StringPtr(deduplicateField)
	return _options
}

// SetSimilar : Allow user to set Similar
func (_options *QueryNoticesOptions) SetSimilar(similar bool) *QueryNoticesOptions {
	_options.Similar = core.BoolPtr(similar)
	return _options
}

// SetSimilarDocumentIds : Allow user to set SimilarDocumentIds
func (_options *QueryNoticesOptions) SetSimilarDocumentIds(similarDocumentIds []string) *QueryNoticesOptions {
	_options.SimilarDocumentIds = similarDocumentIds
	return _options
}

// SetSimilarFields : Allow user to set SimilarFields
func (_options *QueryNoticesOptions) SetSimilarFields(similarFields []string) *QueryNoticesOptions {
	_options.SimilarFields = similarFields
	return _options
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
	Results []QueryNoticesResult `json:"results,omitempty"`

	// Array of aggregation results that match the query.
	Aggregations []QueryAggregationIntf `json:"aggregations,omitempty"`

	// Array of passage results that match the query.
	Passages []QueryPassages `json:"passages,omitempty"`

	// The number of duplicates removed from this notices query.
	DuplicatesRemoved *int64 `json:"duplicates_removed,omitempty"`
}

// UnmarshalQueryNoticesResponse unmarshals an instance of QueryNoticesResponse from the specified map of raw messages.
func UnmarshalQueryNoticesResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(QueryNoticesResponse)
	err = core.UnmarshalPrimitive(m, "matching_results", &obj.MatchingResults)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "results", &obj.Results, UnmarshalQueryNoticesResult)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "aggregations", &obj.Aggregations, UnmarshalQueryAggregation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "passages", &obj.Passages, UnmarshalQueryPassages)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "duplicates_removed", &obj.DuplicatesRemoved)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// QueryNoticesResult : Query result object.
type QueryNoticesResult struct {
	// The unique identifier of the document.
	ID *string `json:"id,omitempty"`

	// Metadata of the document.
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// The collection ID of the collection containing the document for this result.
	CollectionID *string `json:"collection_id,omitempty"`

	// Metadata of a query result.
	ResultMetadata *QueryResultMetadata `json:"result_metadata,omitempty"`

	// The internal status code returned by the ingestion subsystem indicating the overall result of ingesting the source
	// document.
	Code *int64 `json:"code,omitempty"`

	// Name of the original source file (if available).
	Filename *string `json:"filename,omitempty"`

	// The type of the original source file.
	FileType *string `json:"file_type,omitempty"`

	// The SHA-1 hash of the original source file (formatted as a hexadecimal string).
	Sha1 *string `json:"sha1,omitempty"`

	// Array of notices for the document.
	Notices []Notice `json:"notices,omitempty"`

	// Allows users to set arbitrary properties
	additionalProperties map[string]interface{}
}

// Constants associated with the QueryNoticesResult.FileType property.
// The type of the original source file.
const (
	QueryNoticesResultFileTypeHTMLConst = "html"
	QueryNoticesResultFileTypeJSONConst = "json"
	QueryNoticesResultFileTypePDFConst  = "pdf"
	QueryNoticesResultFileTypeWordConst = "word"
)

// SetProperty allows the user to set an arbitrary property on an instance of QueryNoticesResult
func (o *QueryNoticesResult) SetProperty(key string, value interface{}) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string]interface{})
	}
	o.additionalProperties[key] = value
}

// SetProperties allows the user to set a map of arbitrary properties on an instance of QueryNoticesResult
func (o *QueryNoticesResult) SetProperties(m map[string]interface{}) {
	o.additionalProperties = make(map[string]interface{})
	for k, v := range m {
		o.additionalProperties[k] = v
	}
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of QueryNoticesResult
func (o *QueryNoticesResult) GetProperty(key string) interface{} {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of QueryNoticesResult
func (o *QueryNoticesResult) GetProperties() map[string]interface{} {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of QueryNoticesResult
func (o *QueryNoticesResult) MarshalJSON() (buffer []byte, err error) {
	m := make(map[string]interface{})
	if len(o.additionalProperties) > 0 {
		for k, v := range o.additionalProperties {
			m[k] = v
		}
	}
	if o.ID != nil {
		m["id"] = o.ID
	}
	if o.Metadata != nil {
		m["metadata"] = o.Metadata
	}
	if o.CollectionID != nil {
		m["collection_id"] = o.CollectionID
	}
	if o.ResultMetadata != nil {
		m["result_metadata"] = o.ResultMetadata
	}
	if o.Code != nil {
		m["code"] = o.Code
	}
	if o.Filename != nil {
		m["filename"] = o.Filename
	}
	if o.FileType != nil {
		m["file_type"] = o.FileType
	}
	if o.Sha1 != nil {
		m["sha1"] = o.Sha1
	}
	if o.Notices != nil {
		m["notices"] = o.Notices
	}
	buffer, err = json.Marshal(m)
	return
}

// UnmarshalQueryNoticesResult unmarshals an instance of QueryNoticesResult from the specified map of raw messages.
func UnmarshalQueryNoticesResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(QueryNoticesResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	delete(m, "id")
	err = core.UnmarshalPrimitive(m, "metadata", &obj.Metadata)
	if err != nil {
		return
	}
	delete(m, "metadata")
	err = core.UnmarshalPrimitive(m, "collection_id", &obj.CollectionID)
	if err != nil {
		return
	}
	delete(m, "collection_id")
	err = core.UnmarshalModel(m, "result_metadata", &obj.ResultMetadata, UnmarshalQueryResultMetadata)
	if err != nil {
		return
	}
	delete(m, "result_metadata")
	err = core.UnmarshalPrimitive(m, "code", &obj.Code)
	if err != nil {
		return
	}
	delete(m, "code")
	err = core.UnmarshalPrimitive(m, "filename", &obj.Filename)
	if err != nil {
		return
	}
	delete(m, "filename")
	err = core.UnmarshalPrimitive(m, "file_type", &obj.FileType)
	if err != nil {
		return
	}
	delete(m, "file_type")
	err = core.UnmarshalPrimitive(m, "sha1", &obj.Sha1)
	if err != nil {
		return
	}
	delete(m, "sha1")
	err = core.UnmarshalModel(m, "notices", &obj.Notices, UnmarshalNotice)
	if err != nil {
		return
	}
	delete(m, "notices")
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

// QueryOptions : The Query options.
type QueryOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The ID of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// A cacheable query that excludes documents that don't mention the query content. Filter searches are better for
	// metadata-type searches and for assessing the concepts in the data set.
	Filter *string `json:"filter,omitempty"`

	// A query search returns all documents in your data set with full enrichments and full text, but with the most
	// relevant documents listed first. Use a query search when you want to find the most relevant search results.
	Query *string `json:"query,omitempty"`

	// A natural language query that returns relevant documents by utilizing training data and natural language
	// understanding.
	NaturalLanguageQuery *string `json:"natural_language_query,omitempty"`

	// A passages query that returns the most relevant passages from the results.
	Passages *bool `json:"passages,omitempty"`

	// An aggregation search that returns an exact answer by combining query search with filters. Useful for applications
	// to build lists, tables, and time series. For a full list of possible aggregations, see the Query reference.
	Aggregation *string `json:"aggregation,omitempty"`

	// Number of results to return.
	Count *int64 `json:"count,omitempty"`

	// A comma-separated list of the portion of the document hierarchy to return.
	Return *string `json:"return,omitempty"`

	// The number of query results to skip at the beginning. For example, if the total number of results that are returned
	// is 10 and the offset is 8, it returns the last two results.
	Offset *int64 `json:"offset,omitempty"`

	// A comma-separated list of fields in the document to sort on. You can optionally specify a sort direction by
	// prefixing the field with `-` for descending or `+` for ascending. Ascending is the default sort direction if no
	// prefix is specified. This parameter cannot be used in the same query as the **bias** parameter.
	Sort *string `json:"sort,omitempty"`

	// When true, a highlight field is returned for each result which contains the fields which match the query with
	// `<em></em>` tags around the matching query terms.
	Highlight *bool `json:"highlight,omitempty"`

	// A comma-separated list of fields that passages are drawn from. If this parameter not specified, then all top-level
	// fields are included.
	PassagesFields *string `json:"passages.fields,omitempty"`

	// The maximum number of passages to return. The search returns fewer passages if the requested total is not found. The
	// default is `10`. The maximum is `100`.
	PassagesCount *int64 `json:"passages.count,omitempty"`

	// The approximate number of characters that any one passage will have.
	PassagesCharacters *int64 `json:"passages.characters,omitempty"`

	// When `true`, and used with a Watson Discovery News collection, duplicate results (based on the contents of the
	// **title** field) are removed. Duplicate comparison is limited to the current query only; **offset** is not
	// considered. This parameter is currently Beta functionality.
	Deduplicate *bool `json:"deduplicate,omitempty"`

	// When specified, duplicate results based on the field specified are removed from the returned results. Duplicate
	// comparison is limited to the current query only, **offset** is not considered. This parameter is currently Beta
	// functionality.
	DeduplicateField *string `json:"deduplicate.field,omitempty"`

	// When `true`, results are returned based on their similarity to the document IDs specified in the
	// **similar.document_ids** parameter.
	Similar *bool `json:"similar,omitempty"`

	// A comma-separated list of document IDs to find similar documents.
	//
	// **Tip:** Include the **natural_language_query** parameter to expand the scope of the document similarity search with
	// the natural language query. Other query parameters, such as **filter** and **query**, are subsequently applied and
	// reduce the scope.
	SimilarDocumentIds *string `json:"similar.document_ids,omitempty"`

	// A comma-separated list of field names that are used as a basis for comparison to identify similar documents. If not
	// specified, the entire document is used for comparison.
	SimilarFields *string `json:"similar.fields,omitempty"`

	// Field which the returned results will be biased against. The specified field must be either a **date** or **number**
	// format. When a **date** type field is specified returned results are biased towards field values closer to the
	// current date. When a **number** type field is specified, returned results are biased towards higher field values.
	// This parameter cannot be used in the same query as the **sort** parameter.
	Bias *string `json:"bias,omitempty"`

	// When `true` and the **natural_language_query** parameter is used, the **natural_languge_query** parameter is spell
	// checked. The most likely correction is returned in the **suggested_query** field of the response (if one exists).
	//
	// **Important:** this parameter is only valid when using the Cloud Pak version of Discovery.
	SpellingSuggestions *bool `json:"spelling_suggestions,omitempty"`

	// If `true`, queries are not stored in the Discovery **Logs** endpoint.
	XWatsonLoggingOptOut *bool `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewQueryOptions : Instantiate QueryOptions
func (*DiscoveryV1) NewQueryOptions(environmentID string, collectionID string) *QueryOptions {
	return &QueryOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *QueryOptions) SetEnvironmentID(environmentID string) *QueryOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetCollectionID : Allow user to set CollectionID
func (_options *QueryOptions) SetCollectionID(collectionID string) *QueryOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetFilter : Allow user to set Filter
func (_options *QueryOptions) SetFilter(filter string) *QueryOptions {
	_options.Filter = core.StringPtr(filter)
	return _options
}

// SetQuery : Allow user to set Query
func (_options *QueryOptions) SetQuery(query string) *QueryOptions {
	_options.Query = core.StringPtr(query)
	return _options
}

// SetNaturalLanguageQuery : Allow user to set NaturalLanguageQuery
func (_options *QueryOptions) SetNaturalLanguageQuery(naturalLanguageQuery string) *QueryOptions {
	_options.NaturalLanguageQuery = core.StringPtr(naturalLanguageQuery)
	return _options
}

// SetPassages : Allow user to set Passages
func (_options *QueryOptions) SetPassages(passages bool) *QueryOptions {
	_options.Passages = core.BoolPtr(passages)
	return _options
}

// SetAggregation : Allow user to set Aggregation
func (_options *QueryOptions) SetAggregation(aggregation string) *QueryOptions {
	_options.Aggregation = core.StringPtr(aggregation)
	return _options
}

// SetCount : Allow user to set Count
func (_options *QueryOptions) SetCount(count int64) *QueryOptions {
	_options.Count = core.Int64Ptr(count)
	return _options
}

// SetReturn : Allow user to set Return
func (_options *QueryOptions) SetReturn(returnVar string) *QueryOptions {
	_options.Return = core.StringPtr(returnVar)
	return _options
}

// SetOffset : Allow user to set Offset
func (_options *QueryOptions) SetOffset(offset int64) *QueryOptions {
	_options.Offset = core.Int64Ptr(offset)
	return _options
}

// SetSort : Allow user to set Sort
func (_options *QueryOptions) SetSort(sort string) *QueryOptions {
	_options.Sort = core.StringPtr(sort)
	return _options
}

// SetHighlight : Allow user to set Highlight
func (_options *QueryOptions) SetHighlight(highlight bool) *QueryOptions {
	_options.Highlight = core.BoolPtr(highlight)
	return _options
}

// SetPassagesFields : Allow user to set PassagesFields
func (_options *QueryOptions) SetPassagesFields(passagesFields string) *QueryOptions {
	_options.PassagesFields = core.StringPtr(passagesFields)
	return _options
}

// SetPassagesCount : Allow user to set PassagesCount
func (_options *QueryOptions) SetPassagesCount(passagesCount int64) *QueryOptions {
	_options.PassagesCount = core.Int64Ptr(passagesCount)
	return _options
}

// SetPassagesCharacters : Allow user to set PassagesCharacters
func (_options *QueryOptions) SetPassagesCharacters(passagesCharacters int64) *QueryOptions {
	_options.PassagesCharacters = core.Int64Ptr(passagesCharacters)
	return _options
}

// SetDeduplicate : Allow user to set Deduplicate
func (_options *QueryOptions) SetDeduplicate(deduplicate bool) *QueryOptions {
	_options.Deduplicate = core.BoolPtr(deduplicate)
	return _options
}

// SetDeduplicateField : Allow user to set DeduplicateField
func (_options *QueryOptions) SetDeduplicateField(deduplicateField string) *QueryOptions {
	_options.DeduplicateField = core.StringPtr(deduplicateField)
	return _options
}

// SetSimilar : Allow user to set Similar
func (_options *QueryOptions) SetSimilar(similar bool) *QueryOptions {
	_options.Similar = core.BoolPtr(similar)
	return _options
}

// SetSimilarDocumentIds : Allow user to set SimilarDocumentIds
func (_options *QueryOptions) SetSimilarDocumentIds(similarDocumentIds string) *QueryOptions {
	_options.SimilarDocumentIds = core.StringPtr(similarDocumentIds)
	return _options
}

// SetSimilarFields : Allow user to set SimilarFields
func (_options *QueryOptions) SetSimilarFields(similarFields string) *QueryOptions {
	_options.SimilarFields = core.StringPtr(similarFields)
	return _options
}

// SetBias : Allow user to set Bias
func (_options *QueryOptions) SetBias(bias string) *QueryOptions {
	_options.Bias = core.StringPtr(bias)
	return _options
}

// SetSpellingSuggestions : Allow user to set SpellingSuggestions
func (_options *QueryOptions) SetSpellingSuggestions(spellingSuggestions bool) *QueryOptions {
	_options.SpellingSuggestions = core.BoolPtr(spellingSuggestions)
	return _options
}

// SetXWatsonLoggingOptOut : Allow user to set XWatsonLoggingOptOut
func (_options *QueryOptions) SetXWatsonLoggingOptOut(xWatsonLoggingOptOut bool) *QueryOptions {
	_options.XWatsonLoggingOptOut = core.BoolPtr(xWatsonLoggingOptOut)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *QueryOptions) SetHeaders(param map[string]string) *QueryOptions {
	options.Headers = param
	return options
}

// QueryPassages : A passage query result.
type QueryPassages struct {
	// The unique identifier of the document from which the passage has been extracted.
	DocumentID *string `json:"document_id,omitempty"`

	// The confidence score of the passages's analysis. A higher score indicates greater confidence.
	PassageScore *float64 `json:"passage_score,omitempty"`

	// The content of the extracted passage.
	PassageText *string `json:"passage_text,omitempty"`

	// The position of the first character of the extracted passage in the originating field.
	StartOffset *int64 `json:"start_offset,omitempty"`

	// The position of the last character of the extracted passage in the originating field.
	EndOffset *int64 `json:"end_offset,omitempty"`

	// The label of the field from which the passage has been extracted.
	Field *string `json:"field,omitempty"`
}

// UnmarshalQueryPassages unmarshals an instance of QueryPassages from the specified map of raw messages.
func UnmarshalQueryPassages(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(QueryPassages)
	err = core.UnmarshalPrimitive(m, "document_id", &obj.DocumentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "passage_score", &obj.PassageScore)
	if err != nil {
		return
	}
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

// QueryResponse : A response containing the documents and aggregations for the query.
type QueryResponse struct {
	// The number of matching results for the query.
	MatchingResults *int64 `json:"matching_results,omitempty"`

	// Array of document results for the query.
	Results []QueryResult `json:"results,omitempty"`

	// Array of aggregation results for the query.
	Aggregations []QueryAggregationIntf `json:"aggregations,omitempty"`

	// Array of passage results for the query.
	Passages []QueryPassages `json:"passages,omitempty"`

	// The number of duplicate results removed.
	DuplicatesRemoved *int64 `json:"duplicates_removed,omitempty"`

	// The session token for this query. The session token can be used to add events associated with this query to the
	// query and event log.
	//
	// **Important:** Session tokens are case sensitive.
	SessionToken *string `json:"session_token,omitempty"`

	// An object contain retrieval type information.
	RetrievalDetails *RetrievalDetails `json:"retrieval_details,omitempty"`

	// The suggestions for a misspelled natural language query.
	SuggestedQuery *string `json:"suggested_query,omitempty"`
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
	err = core.UnmarshalModel(m, "passages", &obj.Passages, UnmarshalQueryPassages)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "duplicates_removed", &obj.DuplicatesRemoved)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "session_token", &obj.SessionToken)
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// QueryResult : Query result object.
type QueryResult struct {
	// The unique identifier of the document.
	ID *string `json:"id,omitempty"`

	// Metadata of the document.
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// The collection ID of the collection containing the document for this result.
	CollectionID *string `json:"collection_id,omitempty"`

	// Metadata of a query result.
	ResultMetadata *QueryResultMetadata `json:"result_metadata,omitempty"`

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

// SetProperties allows the user to set a map of arbitrary properties on an instance of QueryResult
func (o *QueryResult) SetProperties(m map[string]interface{}) {
	o.additionalProperties = make(map[string]interface{})
	for k, v := range m {
		o.additionalProperties[k] = v
	}
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
	if o.ID != nil {
		m["id"] = o.ID
	}
	if o.Metadata != nil {
		m["metadata"] = o.Metadata
	}
	if o.CollectionID != nil {
		m["collection_id"] = o.CollectionID
	}
	if o.ResultMetadata != nil {
		m["result_metadata"] = o.ResultMetadata
	}
	buffer, err = json.Marshal(m)
	return
}

// UnmarshalQueryResult unmarshals an instance of QueryResult from the specified map of raw messages.
func UnmarshalQueryResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(QueryResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	delete(m, "id")
	err = core.UnmarshalPrimitive(m, "metadata", &obj.Metadata)
	if err != nil {
		return
	}
	delete(m, "metadata")
	err = core.UnmarshalPrimitive(m, "collection_id", &obj.CollectionID)
	if err != nil {
		return
	}
	delete(m, "collection_id")
	err = core.UnmarshalModel(m, "result_metadata", &obj.ResultMetadata, UnmarshalQueryResultMetadata)
	if err != nil {
		return
	}
	delete(m, "result_metadata")
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
	// An unbounded measure of the relevance of a particular result, dependent on the query and matching document. A higher
	// score indicates a greater match to the query parameters.
	Score *float64 `json:"score" validate:"required"`

	// The confidence score for the given result. Calculated based on how relevant the result is estimated to be.
	// confidence can range from `0.0` to `1.0`. The higher the number, the more relevant the document. The `confidence`
	// value for a result was calculated using the model specified in the `document_retrieval_strategy` field of the result
	// set.
	Confidence *float64 `json:"confidence,omitempty"`
}

// UnmarshalQueryResultMetadata unmarshals an instance of QueryResultMetadata from the specified map of raw messages.
func UnmarshalQueryResultMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(QueryResultMetadata)
	err = core.UnmarshalPrimitive(m, "score", &obj.Score)
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

// RetrievalDetails : An object contain retrieval type information.
type RetrievalDetails struct {
	// Indentifies the document retrieval strategy used for this query. `relevancy_training` indicates that the results
	// were returned using a relevancy trained model. `continuous_relevancy_training` indicates that the results were
	// returned using the continuous relevancy training model created by result feedback analysis. `untrained` means the
	// results were returned using the standard untrained model.
	//
	//  **Note**: In the event of trained collections being queried, but the trained model is not used to return results,
	// the **document_retrieval_strategy** will be listed as `untrained`.
	DocumentRetrievalStrategy *string `json:"document_retrieval_strategy,omitempty"`
}

// Constants associated with the RetrievalDetails.DocumentRetrievalStrategy property.
// Indentifies the document retrieval strategy used for this query. `relevancy_training` indicates that the results were
// returned using a relevancy trained model. `continuous_relevancy_training` indicates that the results were returned
// using the continuous relevancy training model created by result feedback analysis. `untrained` means the results were
// returned using the standard untrained model.
//
//  **Note**: In the event of trained collections being queried, but the trained model is not used to return results,
// the **document_retrieval_strategy** will be listed as `untrained`.
const (
	RetrievalDetailsDocumentRetrievalStrategyContinuousRelevancyTrainingConst = "continuous_relevancy_training"
	RetrievalDetailsDocumentRetrievalStrategyRelevancyTrainingConst           = "relevancy_training"
	RetrievalDetailsDocumentRetrievalStrategyUntrainedConst                   = "untrained"
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

// SduStatus : Object containing smart document understanding information for this collection.
type SduStatus struct {
	// When `true`, smart document understanding conversion is enabled for this collection. All collections created with a
	// version date after `2019-04-30` have smart document understanding enabled. If `false`, documents added to the
	// collection are converted using the **conversion** settings specified in the configuration associated with the
	// collection.
	Enabled *bool `json:"enabled,omitempty"`

	// The total number of pages annotated using smart document understanding in this collection.
	TotalAnnotatedPages *int64 `json:"total_annotated_pages,omitempty"`

	// The current number of pages that can be used for training smart document understanding. The `total_pages` number is
	// calculated as the total number of pages identified from the documents listed in the **total_documents** field.
	TotalPages *int64 `json:"total_pages,omitempty"`

	// The total number of documents in this collection that can be used to train smart document understanding. For
	// **lite** plan collections, the maximum is the first 20 uploaded documents (not including HTML or JSON documents).
	// For other plans, the maximum is the first 40 uploaded documents (not including HTML or JSON documents). When the
	// maximum is reached, additional documents uploaded to the collection are not considered for training smart document
	// understanding.
	TotalDocuments *int64 `json:"total_documents,omitempty"`

	// Information about custom smart document understanding fields that exist in this collection.
	CustomFields *SduStatusCustomFields `json:"custom_fields,omitempty"`
}

// UnmarshalSduStatus unmarshals an instance of SduStatus from the specified map of raw messages.
func UnmarshalSduStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SduStatus)
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_annotated_pages", &obj.TotalAnnotatedPages)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_pages", &obj.TotalPages)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_documents", &obj.TotalDocuments)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "custom_fields", &obj.CustomFields, UnmarshalSduStatusCustomFields)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SduStatusCustomFields : Information about custom smart document understanding fields that exist in this collection.
type SduStatusCustomFields struct {
	// The number of custom fields defined for this collection.
	Defined *int64 `json:"defined,omitempty"`

	// The maximum number of custom fields that are allowed in this collection.
	MaximumAllowed *int64 `json:"maximum_allowed,omitempty"`
}

// UnmarshalSduStatusCustomFields unmarshals an instance of SduStatusCustomFields from the specified map of raw messages.
func UnmarshalSduStatusCustomFields(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SduStatusCustomFields)
	err = core.UnmarshalPrimitive(m, "defined", &obj.Defined)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "maximum_allowed", &obj.MaximumAllowed)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SearchStatus : Information about the Continuous Relevancy Training for this environment.
type SearchStatus struct {
	// Current scope of the training. Always returned as `environment`.
	Scope *string `json:"scope,omitempty"`

	// The current status of Continuous Relevancy Training for this environment.
	Status *string `json:"status,omitempty"`

	// Long description of the current Continuous Relevancy Training status.
	StatusDescription *string `json:"status_description,omitempty"`

	// The date stamp of the most recent completed training for this environment.
	LastTrained *strfmt.Date `json:"last_trained,omitempty"`
}

// Constants associated with the SearchStatus.Status property.
// The current status of Continuous Relevancy Training for this environment.
const (
	SearchStatusStatusInsufficentDataConst = "INSUFFICENT_DATA"
	SearchStatusStatusNoDataConst          = "NO_DATA"
	SearchStatusStatusNotApplicableConst   = "NOT_APPLICABLE"
	SearchStatusStatusTrainedConst         = "TRAINED"
	SearchStatusStatusTrainingConst        = "TRAINING"
)

// UnmarshalSearchStatus unmarshals an instance of SearchStatus from the specified map of raw messages.
func UnmarshalSearchStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SearchStatus)
	err = core.UnmarshalPrimitive(m, "scope", &obj.Scope)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status_description", &obj.StatusDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_trained", &obj.LastTrained)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SegmentSettings : A list of Document Segmentation settings.
type SegmentSettings struct {
	// Enables/disables the Document Segmentation feature.
	Enabled *bool `json:"enabled,omitempty"`

	// Defines the heading level that splits into document segments. Valid values are h1, h2, h3, h4, h5, h6. The content
	// of the header field that the segmentation splits at is used as the **title** field for that segmented result. Only
	// valid if used with a collection that has **enabled** set to `false` in the **smart_document_understanding** object.
	SelectorTags []string `json:"selector_tags,omitempty"`

	// Defines the annotated smart document understanding fields that the document is split on. The content of the
	// annotated field that the segmentation splits at is used as the **title** field for that segmented result. For
	// example, if the field `sub-title` is specified, when a document is uploaded each time the smart documement
	// understanding conversion encounters a field of type `sub-title` the document is split at that point and the content
	// of the field used as the title of the remaining content. Thnis split is performed for all instances of the listed
	// fields in the uploaded document. Only valid if used with a collection that has **enabled** set to `true` in the
	// **smart_document_understanding** object.
	AnnotatedFields []string `json:"annotated_fields,omitempty"`
}

// UnmarshalSegmentSettings unmarshals an instance of SegmentSettings from the specified map of raw messages.
func UnmarshalSegmentSettings(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SegmentSettings)
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "selector_tags", &obj.SelectorTags)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "annotated_fields", &obj.AnnotatedFields)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Source : Object containing source parameters for the configuration.
type Source struct {
	// The type of source to connect to.
	// -  `box` indicates the configuration is to connect an instance of Enterprise Box.
	// -  `salesforce` indicates the configuration is to connect to Salesforce.
	// -  `sharepoint` indicates the configuration is to connect to Microsoft SharePoint Online.
	// -  `web_crawl` indicates the configuration is to perform a web page crawl.
	// -  `cloud_object_storage` indicates the configuration is to connect to a cloud object store.
	Type *string `json:"type,omitempty"`

	// The **credential_id** of the credentials to use to connect to the source. Credentials are defined using the
	// **credentials** method. The **source_type** of the credentials used must match the **type** field specified in this
	// object.
	CredentialID *string `json:"credential_id,omitempty"`

	// Object containing the schedule information for the source.
	Schedule *SourceSchedule `json:"schedule,omitempty"`

	// The **options** object defines which items to crawl from the source system.
	Options *SourceOptions `json:"options,omitempty"`
}

// Constants associated with the Source.Type property.
// The type of source to connect to.
// -  `box` indicates the configuration is to connect an instance of Enterprise Box.
// -  `salesforce` indicates the configuration is to connect to Salesforce.
// -  `sharepoint` indicates the configuration is to connect to Microsoft SharePoint Online.
// -  `web_crawl` indicates the configuration is to perform a web page crawl.
// -  `cloud_object_storage` indicates the configuration is to connect to a cloud object store.
const (
	SourceTypeBoxConst                = "box"
	SourceTypeCloudObjectStorageConst = "cloud_object_storage"
	SourceTypeSalesforceConst         = "salesforce"
	SourceTypeSharepointConst         = "sharepoint"
	SourceTypeWebCrawlConst           = "web_crawl"
)

// UnmarshalSource unmarshals an instance of Source from the specified map of raw messages.
func UnmarshalSource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Source)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "credential_id", &obj.CredentialID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "schedule", &obj.Schedule, UnmarshalSourceSchedule)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "options", &obj.Options, UnmarshalSourceOptions)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SourceOptions : The **options** object defines which items to crawl from the source system.
type SourceOptions struct {
	// Array of folders to crawl from the Box source. Only valid, and required, when the **type** field of the **source**
	// object is set to `box`.
	Folders []SourceOptionsFolder `json:"folders,omitempty"`

	// Array of Salesforce document object types to crawl from the Salesforce source. Only valid, and required, when the
	// **type** field of the **source** object is set to `salesforce`.
	Objects []SourceOptionsObject `json:"objects,omitempty"`

	// Array of Microsoft SharePointoint Online site collections to crawl from the SharePoint source. Only valid and
	// required when the **type** field of the **source** object is set to `sharepoint`.
	SiteCollections []SourceOptionsSiteColl `json:"site_collections,omitempty"`

	// Array of Web page URLs to begin crawling the web from. Only valid and required when the **type** field of the
	// **source** object is set to `web_crawl`.
	Urls []SourceOptionsWebCrawl `json:"urls,omitempty"`

	// Array of cloud object store buckets to begin crawling. Only valid and required when the **type** field of the
	// **source** object is set to `cloud_object_store`, and the **crawl_all_buckets** field is `false` or not specified.
	Buckets []SourceOptionsBuckets `json:"buckets,omitempty"`

	// When `true`, all buckets in the specified cloud object store are crawled. If set to `true`, the **buckets** array
	// must not be specified.
	CrawlAllBuckets *bool `json:"crawl_all_buckets,omitempty"`
}

// UnmarshalSourceOptions unmarshals an instance of SourceOptions from the specified map of raw messages.
func UnmarshalSourceOptions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SourceOptions)
	err = core.UnmarshalModel(m, "folders", &obj.Folders, UnmarshalSourceOptionsFolder)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "objects", &obj.Objects, UnmarshalSourceOptionsObject)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "site_collections", &obj.SiteCollections, UnmarshalSourceOptionsSiteColl)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "urls", &obj.Urls, UnmarshalSourceOptionsWebCrawl)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "buckets", &obj.Buckets, UnmarshalSourceOptionsBuckets)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crawl_all_buckets", &obj.CrawlAllBuckets)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SourceOptionsBuckets : Object defining a cloud object store bucket to crawl.
type SourceOptionsBuckets struct {
	// The name of the cloud object store bucket to crawl.
	Name *string `json:"name" validate:"required"`

	// The number of documents to crawl from this cloud object store bucket. If not specified, all documents in the bucket
	// are crawled.
	Limit *int64 `json:"limit,omitempty"`
}

// NewSourceOptionsBuckets : Instantiate SourceOptionsBuckets (Generic Model Constructor)
func (*DiscoveryV1) NewSourceOptionsBuckets(name string) (_model *SourceOptionsBuckets, err error) {
	_model = &SourceOptionsBuckets{
		Name: core.StringPtr(name),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalSourceOptionsBuckets unmarshals an instance of SourceOptionsBuckets from the specified map of raw messages.
func UnmarshalSourceOptionsBuckets(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SourceOptionsBuckets)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SourceOptionsFolder : Object that defines a box folder to crawl with this configuration.
type SourceOptionsFolder struct {
	// The Box user ID of the user who owns the folder to crawl.
	OwnerUserID *string `json:"owner_user_id" validate:"required"`

	// The Box folder ID of the folder to crawl.
	FolderID *string `json:"folder_id" validate:"required"`

	// The maximum number of documents to crawl for this folder. By default, all documents in the folder are crawled.
	Limit *int64 `json:"limit,omitempty"`
}

// NewSourceOptionsFolder : Instantiate SourceOptionsFolder (Generic Model Constructor)
func (*DiscoveryV1) NewSourceOptionsFolder(ownerUserID string, folderID string) (_model *SourceOptionsFolder, err error) {
	_model = &SourceOptionsFolder{
		OwnerUserID: core.StringPtr(ownerUserID),
		FolderID:    core.StringPtr(folderID),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalSourceOptionsFolder unmarshals an instance of SourceOptionsFolder from the specified map of raw messages.
func UnmarshalSourceOptionsFolder(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SourceOptionsFolder)
	err = core.UnmarshalPrimitive(m, "owner_user_id", &obj.OwnerUserID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "folder_id", &obj.FolderID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SourceOptionsObject : Object that defines a Salesforce document object type crawl with this configuration.
type SourceOptionsObject struct {
	// The name of the Salesforce document object to crawl. For example, `case`.
	Name *string `json:"name" validate:"required"`

	// The maximum number of documents to crawl for this document object. By default, all documents in the document object
	// are crawled.
	Limit *int64 `json:"limit,omitempty"`
}

// NewSourceOptionsObject : Instantiate SourceOptionsObject (Generic Model Constructor)
func (*DiscoveryV1) NewSourceOptionsObject(name string) (_model *SourceOptionsObject, err error) {
	_model = &SourceOptionsObject{
		Name: core.StringPtr(name),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalSourceOptionsObject unmarshals an instance of SourceOptionsObject from the specified map of raw messages.
func UnmarshalSourceOptionsObject(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SourceOptionsObject)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SourceOptionsSiteColl : Object that defines a Microsoft SharePoint site collection to crawl with this configuration.
type SourceOptionsSiteColl struct {
	// The Microsoft SharePoint Online site collection path to crawl. The path must be be relative to the
	// **organization_url** that was specified in the credentials associated with this source configuration.
	SiteCollectionPath *string `json:"site_collection_path" validate:"required"`

	// The maximum number of documents to crawl for this site collection. By default, all documents in the site collection
	// are crawled.
	Limit *int64 `json:"limit,omitempty"`
}

// NewSourceOptionsSiteColl : Instantiate SourceOptionsSiteColl (Generic Model Constructor)
func (*DiscoveryV1) NewSourceOptionsSiteColl(siteCollectionPath string) (_model *SourceOptionsSiteColl, err error) {
	_model = &SourceOptionsSiteColl{
		SiteCollectionPath: core.StringPtr(siteCollectionPath),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalSourceOptionsSiteColl unmarshals an instance of SourceOptionsSiteColl from the specified map of raw messages.
func UnmarshalSourceOptionsSiteColl(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SourceOptionsSiteColl)
	err = core.UnmarshalPrimitive(m, "site_collection_path", &obj.SiteCollectionPath)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SourceOptionsWebCrawl : Object defining which URL to crawl and how to crawl it.
type SourceOptionsWebCrawl struct {
	// The starting URL to crawl.
	URL *string `json:"url" validate:"required"`

	// When `true`, crawls of the specified URL are limited to the host part of the **url** field.
	LimitToStartingHosts *bool `json:"limit_to_starting_hosts,omitempty"`

	// The number of concurrent URLs to fetch. `gentle` means one URL is fetched at a time with a delay between each call.
	// `normal` means as many as two URLs are fectched concurrently with a short delay between fetch calls. `aggressive`
	// means that up to ten URLs are fetched concurrently with a short delay between fetch calls.
	CrawlSpeed *string `json:"crawl_speed,omitempty"`

	// When `true`, allows the crawl to interact with HTTPS sites with SSL certificates with untrusted signers.
	AllowUntrustedCertificate *bool `json:"allow_untrusted_certificate,omitempty"`

	// The maximum number of hops to make from the initial URL. When a page is crawled each link on that page will also be
	// crawled if it is within the **maximum_hops** from the initial URL. The first page crawled is 0 hops, each link
	// crawled from the first page is 1 hop, each link crawled from those pages is 2 hops, and so on.
	MaximumHops *int64 `json:"maximum_hops,omitempty"`

	// The maximum milliseconds to wait for a response from the web server.
	RequestTimeout *int64 `json:"request_timeout,omitempty"`

	// When `true`, the crawler will ignore any `robots.txt` encountered by the crawler. This should only ever be done when
	// crawling a web site the user owns. This must be be set to `true` when a **gateway_id** is specied in the
	// **credentials**.
	OverrideRobotsTxt *bool `json:"override_robots_txt,omitempty"`

	// Array of URL's to be excluded while crawling. The crawler will not follow links which contains this string. For
	// example, listing `https://ibm.com/watson` also excludes `https://ibm.com/watson/discovery`.
	Blacklist []string `json:"blacklist,omitempty"`
}

// Constants associated with the SourceOptionsWebCrawl.CrawlSpeed property.
// The number of concurrent URLs to fetch. `gentle` means one URL is fetched at a time with a delay between each call.
// `normal` means as many as two URLs are fectched concurrently with a short delay between fetch calls. `aggressive`
// means that up to ten URLs are fetched concurrently with a short delay between fetch calls.
const (
	SourceOptionsWebCrawlCrawlSpeedAggressiveConst = "aggressive"
	SourceOptionsWebCrawlCrawlSpeedGentleConst     = "gentle"
	SourceOptionsWebCrawlCrawlSpeedNormalConst     = "normal"
)

// NewSourceOptionsWebCrawl : Instantiate SourceOptionsWebCrawl (Generic Model Constructor)
func (*DiscoveryV1) NewSourceOptionsWebCrawl(url string) (_model *SourceOptionsWebCrawl, err error) {
	_model = &SourceOptionsWebCrawl{
		URL: core.StringPtr(url),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalSourceOptionsWebCrawl unmarshals an instance of SourceOptionsWebCrawl from the specified map of raw messages.
func UnmarshalSourceOptionsWebCrawl(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SourceOptionsWebCrawl)
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit_to_starting_hosts", &obj.LimitToStartingHosts)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crawl_speed", &obj.CrawlSpeed)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "allow_untrusted_certificate", &obj.AllowUntrustedCertificate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "maximum_hops", &obj.MaximumHops)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "request_timeout", &obj.RequestTimeout)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "override_robots_txt", &obj.OverrideRobotsTxt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "blacklist", &obj.Blacklist)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SourceSchedule : Object containing the schedule information for the source.
type SourceSchedule struct {
	// When `true`, the source is re-crawled based on the **frequency** field in this object. When `false` the source is
	// not re-crawled; When `false` and connecting to Salesforce the source is crawled annually.
	Enabled *bool `json:"enabled,omitempty"`

	// The time zone to base source crawl times on. Possible values correspond to the IANA (Internet Assigned Numbers
	// Authority) time zones list.
	TimeZone *string `json:"time_zone,omitempty"`

	// The crawl schedule in the specified **time_zone**.
	//
	// -  `five_minutes`: Runs every five minutes.
	// -  `hourly`: Runs every hour.
	// -  `daily`: Runs every day between 00:00 and 06:00.
	// -  `weekly`: Runs every week on Sunday between 00:00 and 06:00.
	// -  `monthly`: Runs the on the first Sunday of every month between 00:00 and 06:00.
	Frequency *string `json:"frequency,omitempty"`
}

// Constants associated with the SourceSchedule.Frequency property.
// The crawl schedule in the specified **time_zone**.
//
// -  `five_minutes`: Runs every five minutes.
// -  `hourly`: Runs every hour.
// -  `daily`: Runs every day between 00:00 and 06:00.
// -  `weekly`: Runs every week on Sunday between 00:00 and 06:00.
// -  `monthly`: Runs the on the first Sunday of every month between 00:00 and 06:00.
const (
	SourceScheduleFrequencyDailyConst       = "daily"
	SourceScheduleFrequencyFiveMinutesConst = "five_minutes"
	SourceScheduleFrequencyHourlyConst      = "hourly"
	SourceScheduleFrequencyMonthlyConst     = "monthly"
	SourceScheduleFrequencyWeeklyConst      = "weekly"
)

// UnmarshalSourceSchedule unmarshals an instance of SourceSchedule from the specified map of raw messages.
func UnmarshalSourceSchedule(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SourceSchedule)
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "time_zone", &obj.TimeZone)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "frequency", &obj.Frequency)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SourceStatus : Object containing source crawl status information.
type SourceStatus struct {
	// The current status of the source crawl for this collection. This field returns `not_configured` if the default
	// configuration for this source does not have a **source** object defined.
	//
	// -  `running` indicates that a crawl to fetch more documents is in progress.
	// -  `complete` indicates that the crawl has completed with no errors.
	// -  `queued` indicates that the crawl has been paused by the system and will automatically restart when possible.
	// -  `unknown` indicates that an unidentified error has occured in the service.
	Status *string `json:"status,omitempty"`

	// Date in `RFC 3339` format indicating the time of the next crawl attempt.
	NextCrawl *strfmt.DateTime `json:"next_crawl,omitempty"`
}

// Constants associated with the SourceStatus.Status property.
// The current status of the source crawl for this collection. This field returns `not_configured` if the default
// configuration for this source does not have a **source** object defined.
//
// -  `running` indicates that a crawl to fetch more documents is in progress.
// -  `complete` indicates that the crawl has completed with no errors.
// -  `queued` indicates that the crawl has been paused by the system and will automatically restart when possible.
// -  `unknown` indicates that an unidentified error has occured in the service.
const (
	SourceStatusStatusCompleteConst      = "complete"
	SourceStatusStatusNotConfiguredConst = "not_configured"
	SourceStatusStatusQueuedConst        = "queued"
	SourceStatusStatusRunningConst       = "running"
	SourceStatusStatusUnknownConst       = "unknown"
)

// UnmarshalSourceStatus unmarshals an instance of SourceStatus from the specified map of raw messages.
func UnmarshalSourceStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SourceStatus)
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "next_crawl", &obj.NextCrawl)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// StatusDetails : Object that contains details about the status of the authentication process.
type StatusDetails struct {
	// Indicates whether the credential is accepted by the target data source.
	Authenticated *bool `json:"authenticated,omitempty"`

	// If `authenticated` is `false`, a message describes why the authentication was unsuccessful.
	ErrorMessage *string `json:"error_message,omitempty"`
}

// UnmarshalStatusDetails unmarshals an instance of StatusDetails from the specified map of raw messages.
func UnmarshalStatusDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(StatusDetails)
	err = core.UnmarshalPrimitive(m, "authenticated", &obj.Authenticated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "error_message", &obj.ErrorMessage)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TokenDictRule : An object defining a single tokenizaion rule.
type TokenDictRule struct {
	// The string to tokenize.
	Text *string `json:"text" validate:"required"`

	// Array of tokens that the `text` field is split into when found.
	Tokens []string `json:"tokens" validate:"required"`

	// Array of tokens that represent the content of the `text` field in an alternate character set.
	Readings []string `json:"readings,omitempty"`

	// The part of speech that the `text` string belongs to. For example `noun`. Custom parts of speech can be specified.
	PartOfSpeech *string `json:"part_of_speech" validate:"required"`
}

// NewTokenDictRule : Instantiate TokenDictRule (Generic Model Constructor)
func (*DiscoveryV1) NewTokenDictRule(text string, tokens []string, partOfSpeech string) (_model *TokenDictRule, err error) {
	_model = &TokenDictRule{
		Text:         core.StringPtr(text),
		Tokens:       tokens,
		PartOfSpeech: core.StringPtr(partOfSpeech),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalTokenDictRule unmarshals an instance of TokenDictRule from the specified map of raw messages.
func UnmarshalTokenDictRule(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TokenDictRule)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tokens", &obj.Tokens)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "readings", &obj.Readings)
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

// TokenDictStatusResponse : Object describing the current status of the wordlist.
type TokenDictStatusResponse struct {
	// Current wordlist status for the specified collection.
	Status *string `json:"status,omitempty"`

	// The type for this wordlist. Can be `tokenization_dictionary` or `stopwords`.
	Type *string `json:"type,omitempty"`
}

// Constants associated with the TokenDictStatusResponse.Status property.
// Current wordlist status for the specified collection.
const (
	TokenDictStatusResponseStatusActiveConst   = "active"
	TokenDictStatusResponseStatusNotFoundConst = "not found"
	TokenDictStatusResponseStatusPendingConst  = "pending"
)

// UnmarshalTokenDictStatusResponse unmarshals an instance of TokenDictStatusResponse from the specified map of raw messages.
func UnmarshalTokenDictStatusResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TokenDictStatusResponse)
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
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

// TopHitsResults : Top hit information for this query.
type TopHitsResults struct {
	// Number of matching results.
	MatchingResults *int64 `json:"matching_results,omitempty"`

	// Top results returned by the aggregation.
	Hits []QueryResult `json:"hits,omitempty"`
}

// UnmarshalTopHitsResults unmarshals an instance of TopHitsResults from the specified map of raw messages.
func UnmarshalTopHitsResults(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TopHitsResults)
	err = core.UnmarshalPrimitive(m, "matching_results", &obj.MatchingResults)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "hits", &obj.Hits, UnmarshalQueryResult)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TrainingDataSet : Training information for a specific collection.
type TrainingDataSet struct {
	// The environment id associated with this training data set.
	EnvironmentID *string `json:"environment_id,omitempty"`

	// The collection id associated with this training data set.
	CollectionID *string `json:"collection_id,omitempty"`

	// Array of training queries.
	Queries []TrainingQuery `json:"queries,omitempty"`
}

// UnmarshalTrainingDataSet unmarshals an instance of TrainingDataSet from the specified map of raw messages.
func UnmarshalTrainingDataSet(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TrainingDataSet)
	err = core.UnmarshalPrimitive(m, "environment_id", &obj.EnvironmentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "collection_id", &obj.CollectionID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "queries", &obj.Queries, UnmarshalTrainingQuery)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TrainingExample : Training example details.
type TrainingExample struct {
	// The document ID associated with this training example.
	DocumentID *string `json:"document_id,omitempty"`

	// The cross reference associated with this training example.
	CrossReference *string `json:"cross_reference,omitempty"`

	// The relevance of the training example.
	Relevance *int64 `json:"relevance,omitempty"`
}

// UnmarshalTrainingExample unmarshals an instance of TrainingExample from the specified map of raw messages.
func UnmarshalTrainingExample(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TrainingExample)
	err = core.UnmarshalPrimitive(m, "document_id", &obj.DocumentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cross_reference", &obj.CrossReference)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "relevance", &obj.Relevance)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TrainingExampleList : Object containing an array of training examples.
type TrainingExampleList struct {
	// Array of training examples.
	Examples []TrainingExample `json:"examples,omitempty"`
}

// UnmarshalTrainingExampleList unmarshals an instance of TrainingExampleList from the specified map of raw messages.
func UnmarshalTrainingExampleList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TrainingExampleList)
	err = core.UnmarshalModel(m, "examples", &obj.Examples, UnmarshalTrainingExample)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TrainingQuery : Training query details.
type TrainingQuery struct {
	// The query ID associated with the training query.
	QueryID *string `json:"query_id,omitempty"`

	// The natural text query for the training query.
	NaturalLanguageQuery *string `json:"natural_language_query,omitempty"`

	// The filter used on the collection before the **natural_language_query** is applied.
	Filter *string `json:"filter,omitempty"`

	// Array of training examples.
	Examples []TrainingExample `json:"examples,omitempty"`
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
	err = core.UnmarshalModel(m, "examples", &obj.Examples, UnmarshalTrainingExample)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TrainingStatus : Training status details.
type TrainingStatus struct {
	// The total number of training examples uploaded to this collection.
	TotalExamples *int64 `json:"total_examples,omitempty"`

	// When `true`, the collection has been successfully trained.
	Available *bool `json:"available,omitempty"`

	// When `true`, the collection is currently processing training.
	Processing *bool `json:"processing,omitempty"`

	// When `true`, the collection has a sufficent amount of queries added for training to occur.
	MinimumQueriesAdded *bool `json:"minimum_queries_added,omitempty"`

	// When `true`, the collection has a sufficent amount of examples added for training to occur.
	MinimumExamplesAdded *bool `json:"minimum_examples_added,omitempty"`

	// When `true`, the collection has a sufficent amount of diversity in labeled results for training to occur.
	SufficientLabelDiversity *bool `json:"sufficient_label_diversity,omitempty"`

	// The number of notices associated with this data set.
	Notices *int64 `json:"notices,omitempty"`

	// The timestamp of when the collection was successfully trained.
	SuccessfullyTrained *strfmt.DateTime `json:"successfully_trained,omitempty"`

	// The timestamp of when the data was uploaded.
	DataUpdated *strfmt.DateTime `json:"data_updated,omitempty"`
}

// UnmarshalTrainingStatus unmarshals an instance of TrainingStatus from the specified map of raw messages.
func UnmarshalTrainingStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TrainingStatus)
	err = core.UnmarshalPrimitive(m, "total_examples", &obj.TotalExamples)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "available", &obj.Available)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "processing", &obj.Processing)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "minimum_queries_added", &obj.MinimumQueriesAdded)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "minimum_examples_added", &obj.MinimumExamplesAdded)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sufficient_label_diversity", &obj.SufficientLabelDiversity)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "notices", &obj.Notices)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "successfully_trained", &obj.SuccessfullyTrained)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "data_updated", &obj.DataUpdated)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateCollectionOptions : The UpdateCollection options.
type UpdateCollectionOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The ID of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// The name of the collection.
	Name *string `json:"name" validate:"required"`

	// A description of the collection.
	Description *string `json:"description,omitempty"`

	// The ID of the configuration in which the collection is to be updated.
	ConfigurationID *string `json:"configuration_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateCollectionOptions : Instantiate UpdateCollectionOptions
func (*DiscoveryV1) NewUpdateCollectionOptions(environmentID string, collectionID string, name string) *UpdateCollectionOptions {
	return &UpdateCollectionOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
		Name:          core.StringPtr(name),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *UpdateCollectionOptions) SetEnvironmentID(environmentID string) *UpdateCollectionOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
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

// SetConfigurationID : Allow user to set ConfigurationID
func (_options *UpdateCollectionOptions) SetConfigurationID(configurationID string) *UpdateCollectionOptions {
	_options.ConfigurationID = core.StringPtr(configurationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateCollectionOptions) SetHeaders(param map[string]string) *UpdateCollectionOptions {
	options.Headers = param
	return options
}

// UpdateConfigurationOptions : The UpdateConfiguration options.
type UpdateConfigurationOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The ID of the configuration.
	ConfigurationID *string `json:"-" validate:"required,ne="`

	// The name of the configuration.
	Name *string `json:"name" validate:"required"`

	// The description of the configuration, if available.
	Description *string `json:"description,omitempty"`

	// Document conversion settings.
	Conversions *Conversions `json:"conversions,omitempty"`

	// An array of document enrichment settings for the configuration.
	Enrichments []Enrichment `json:"enrichments,omitempty"`

	// Defines operations that can be used to transform the final output JSON into a normalized form. Operations are
	// executed in the order that they appear in the array.
	Normalizations []NormalizationOperation `json:"normalizations,omitempty"`

	// Object containing source parameters for the configuration.
	Source *Source `json:"source,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateConfigurationOptions : Instantiate UpdateConfigurationOptions
func (*DiscoveryV1) NewUpdateConfigurationOptions(environmentID string, configurationID string, name string) *UpdateConfigurationOptions {
	return &UpdateConfigurationOptions{
		EnvironmentID:   core.StringPtr(environmentID),
		ConfigurationID: core.StringPtr(configurationID),
		Name:            core.StringPtr(name),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *UpdateConfigurationOptions) SetEnvironmentID(environmentID string) *UpdateConfigurationOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetConfigurationID : Allow user to set ConfigurationID
func (_options *UpdateConfigurationOptions) SetConfigurationID(configurationID string) *UpdateConfigurationOptions {
	_options.ConfigurationID = core.StringPtr(configurationID)
	return _options
}

// SetName : Allow user to set Name
func (_options *UpdateConfigurationOptions) SetName(name string) *UpdateConfigurationOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *UpdateConfigurationOptions) SetDescription(description string) *UpdateConfigurationOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetConversions : Allow user to set Conversions
func (_options *UpdateConfigurationOptions) SetConversions(conversions *Conversions) *UpdateConfigurationOptions {
	_options.Conversions = conversions
	return _options
}

// SetEnrichments : Allow user to set Enrichments
func (_options *UpdateConfigurationOptions) SetEnrichments(enrichments []Enrichment) *UpdateConfigurationOptions {
	_options.Enrichments = enrichments
	return _options
}

// SetNormalizations : Allow user to set Normalizations
func (_options *UpdateConfigurationOptions) SetNormalizations(normalizations []NormalizationOperation) *UpdateConfigurationOptions {
	_options.Normalizations = normalizations
	return _options
}

// SetSource : Allow user to set Source
func (_options *UpdateConfigurationOptions) SetSource(source *Source) *UpdateConfigurationOptions {
	_options.Source = source
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateConfigurationOptions) SetHeaders(param map[string]string) *UpdateConfigurationOptions {
	options.Headers = param
	return options
}

// UpdateCredentialsOptions : The UpdateCredentials options.
type UpdateCredentialsOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The unique identifier for a set of source credentials.
	CredentialID *string `json:"-" validate:"required,ne="`

	// The source that this credentials object connects to.
	// -  `box` indicates the credentials are used to connect an instance of Enterprise Box.
	// -  `salesforce` indicates the credentials are used to connect to Salesforce.
	// -  `sharepoint` indicates the credentials are used to connect to Microsoft SharePoint Online.
	// -  `web_crawl` indicates the credentials are used to perform a web crawl.
	// =  `cloud_object_storage` indicates the credentials are used to connect to an IBM Cloud Object Store.
	SourceType *string `json:"source_type,omitempty"`

	// Object containing details of the stored credentials.
	//
	// Obtain credentials for your source from the administrator of the source.
	CredentialDetails *CredentialDetails `json:"credential_details,omitempty"`

	// Object that contains details about the status of the authentication process.
	Status *StatusDetails `json:"status,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateCredentialsOptions.SourceType property.
// The source that this credentials object connects to.
// -  `box` indicates the credentials are used to connect an instance of Enterprise Box.
// -  `salesforce` indicates the credentials are used to connect to Salesforce.
// -  `sharepoint` indicates the credentials are used to connect to Microsoft SharePoint Online.
// -  `web_crawl` indicates the credentials are used to perform a web crawl.
// =  `cloud_object_storage` indicates the credentials are used to connect to an IBM Cloud Object Store.
const (
	UpdateCredentialsOptionsSourceTypeBoxConst                = "box"
	UpdateCredentialsOptionsSourceTypeCloudObjectStorageConst = "cloud_object_storage"
	UpdateCredentialsOptionsSourceTypeSalesforceConst         = "salesforce"
	UpdateCredentialsOptionsSourceTypeSharepointConst         = "sharepoint"
	UpdateCredentialsOptionsSourceTypeWebCrawlConst           = "web_crawl"
)

// NewUpdateCredentialsOptions : Instantiate UpdateCredentialsOptions
func (*DiscoveryV1) NewUpdateCredentialsOptions(environmentID string, credentialID string) *UpdateCredentialsOptions {
	return &UpdateCredentialsOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CredentialID:  core.StringPtr(credentialID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *UpdateCredentialsOptions) SetEnvironmentID(environmentID string) *UpdateCredentialsOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetCredentialID : Allow user to set CredentialID
func (_options *UpdateCredentialsOptions) SetCredentialID(credentialID string) *UpdateCredentialsOptions {
	_options.CredentialID = core.StringPtr(credentialID)
	return _options
}

// SetSourceType : Allow user to set SourceType
func (_options *UpdateCredentialsOptions) SetSourceType(sourceType string) *UpdateCredentialsOptions {
	_options.SourceType = core.StringPtr(sourceType)
	return _options
}

// SetCredentialDetails : Allow user to set CredentialDetails
func (_options *UpdateCredentialsOptions) SetCredentialDetails(credentialDetails *CredentialDetails) *UpdateCredentialsOptions {
	_options.CredentialDetails = credentialDetails
	return _options
}

// SetStatus : Allow user to set Status
func (_options *UpdateCredentialsOptions) SetStatus(status *StatusDetails) *UpdateCredentialsOptions {
	_options.Status = status
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateCredentialsOptions) SetHeaders(param map[string]string) *UpdateCredentialsOptions {
	options.Headers = param
	return options
}

// UpdateDocumentOptions : The UpdateDocument options.
type UpdateDocumentOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The ID of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// The ID of the document.
	DocumentID *string `json:"-" validate:"required,ne="`

	// The content of the document to ingest. The maximum supported file size when adding a file to a collection is 50
	// megabytes, the maximum supported file size when testing a configuration is 1 megabyte. Files larger than the
	// supported size are rejected.
	File io.ReadCloser `json:"-"`

	// The filename for file.
	Filename *string `json:"-"`

	// The content type of file.
	FileContentType *string `json:"-"`

	// The maximum supported metadata file size is 1 MB. Metadata parts larger than 1 MB are rejected. Example:  ``` {
	//   "Creator": "Johnny Appleseed",
	//   "Subject": "Apples"
	// } ```.
	Metadata *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateDocumentOptions : Instantiate UpdateDocumentOptions
func (*DiscoveryV1) NewUpdateDocumentOptions(environmentID string, collectionID string, documentID string) *UpdateDocumentOptions {
	return &UpdateDocumentOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
		DocumentID:    core.StringPtr(documentID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *UpdateDocumentOptions) SetEnvironmentID(environmentID string) *UpdateDocumentOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetCollectionID : Allow user to set CollectionID
func (_options *UpdateDocumentOptions) SetCollectionID(collectionID string) *UpdateDocumentOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetDocumentID : Allow user to set DocumentID
func (_options *UpdateDocumentOptions) SetDocumentID(documentID string) *UpdateDocumentOptions {
	_options.DocumentID = core.StringPtr(documentID)
	return _options
}

// SetFile : Allow user to set File
func (_options *UpdateDocumentOptions) SetFile(file io.ReadCloser) *UpdateDocumentOptions {
	_options.File = file
	return _options
}

// SetFilename : Allow user to set Filename
func (_options *UpdateDocumentOptions) SetFilename(filename string) *UpdateDocumentOptions {
	_options.Filename = core.StringPtr(filename)
	return _options
}

// SetFileContentType : Allow user to set FileContentType
func (_options *UpdateDocumentOptions) SetFileContentType(fileContentType string) *UpdateDocumentOptions {
	_options.FileContentType = core.StringPtr(fileContentType)
	return _options
}

// SetMetadata : Allow user to set Metadata
func (_options *UpdateDocumentOptions) SetMetadata(metadata string) *UpdateDocumentOptions {
	_options.Metadata = core.StringPtr(metadata)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateDocumentOptions) SetHeaders(param map[string]string) *UpdateDocumentOptions {
	options.Headers = param
	return options
}

// UpdateEnvironmentOptions : The UpdateEnvironment options.
type UpdateEnvironmentOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// Name that identifies the environment.
	Name *string `json:"name,omitempty"`

	// Description of the environment.
	Description *string `json:"description,omitempty"`

	// Size that the environment should be increased to. Environment size cannot be modified when using a Lite plan.
	// Environment size can only increased and not decreased.
	Size *string `json:"size,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateEnvironmentOptions.Size property.
// Size that the environment should be increased to. Environment size cannot be modified when using a Lite plan.
// Environment size can only increased and not decreased.
const (
	UpdateEnvironmentOptionsSizeLConst    = "L"
	UpdateEnvironmentOptionsSizeMConst    = "M"
	UpdateEnvironmentOptionsSizeMlConst   = "ML"
	UpdateEnvironmentOptionsSizeMsConst   = "MS"
	UpdateEnvironmentOptionsSizeSConst    = "S"
	UpdateEnvironmentOptionsSizeXlConst   = "XL"
	UpdateEnvironmentOptionsSizeXxlConst  = "XXL"
	UpdateEnvironmentOptionsSizeXxxlConst = "XXXL"
)

// NewUpdateEnvironmentOptions : Instantiate UpdateEnvironmentOptions
func (*DiscoveryV1) NewUpdateEnvironmentOptions(environmentID string) *UpdateEnvironmentOptions {
	return &UpdateEnvironmentOptions{
		EnvironmentID: core.StringPtr(environmentID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *UpdateEnvironmentOptions) SetEnvironmentID(environmentID string) *UpdateEnvironmentOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetName : Allow user to set Name
func (_options *UpdateEnvironmentOptions) SetName(name string) *UpdateEnvironmentOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *UpdateEnvironmentOptions) SetDescription(description string) *UpdateEnvironmentOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetSize : Allow user to set Size
func (_options *UpdateEnvironmentOptions) SetSize(size string) *UpdateEnvironmentOptions {
	_options.Size = core.StringPtr(size)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateEnvironmentOptions) SetHeaders(param map[string]string) *UpdateEnvironmentOptions {
	options.Headers = param
	return options
}

// UpdateTrainingExampleOptions : The UpdateTrainingExample options.
type UpdateTrainingExampleOptions struct {
	// The ID of the environment.
	EnvironmentID *string `json:"-" validate:"required,ne="`

	// The ID of the collection.
	CollectionID *string `json:"-" validate:"required,ne="`

	// The ID of the query used for training.
	QueryID *string `json:"-" validate:"required,ne="`

	// The ID of the document as it is indexed.
	ExampleID *string `json:"-" validate:"required,ne="`

	// The example to add.
	CrossReference *string `json:"cross_reference,omitempty"`

	// The relevance value for this example.
	Relevance *int64 `json:"relevance,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateTrainingExampleOptions : Instantiate UpdateTrainingExampleOptions
func (*DiscoveryV1) NewUpdateTrainingExampleOptions(environmentID string, collectionID string, queryID string, exampleID string) *UpdateTrainingExampleOptions {
	return &UpdateTrainingExampleOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
		QueryID:       core.StringPtr(queryID),
		ExampleID:     core.StringPtr(exampleID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (_options *UpdateTrainingExampleOptions) SetEnvironmentID(environmentID string) *UpdateTrainingExampleOptions {
	_options.EnvironmentID = core.StringPtr(environmentID)
	return _options
}

// SetCollectionID : Allow user to set CollectionID
func (_options *UpdateTrainingExampleOptions) SetCollectionID(collectionID string) *UpdateTrainingExampleOptions {
	_options.CollectionID = core.StringPtr(collectionID)
	return _options
}

// SetQueryID : Allow user to set QueryID
func (_options *UpdateTrainingExampleOptions) SetQueryID(queryID string) *UpdateTrainingExampleOptions {
	_options.QueryID = core.StringPtr(queryID)
	return _options
}

// SetExampleID : Allow user to set ExampleID
func (_options *UpdateTrainingExampleOptions) SetExampleID(exampleID string) *UpdateTrainingExampleOptions {
	_options.ExampleID = core.StringPtr(exampleID)
	return _options
}

// SetCrossReference : Allow user to set CrossReference
func (_options *UpdateTrainingExampleOptions) SetCrossReference(crossReference string) *UpdateTrainingExampleOptions {
	_options.CrossReference = core.StringPtr(crossReference)
	return _options
}

// SetRelevance : Allow user to set Relevance
func (_options *UpdateTrainingExampleOptions) SetRelevance(relevance int64) *UpdateTrainingExampleOptions {
	_options.Relevance = core.Int64Ptr(relevance)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateTrainingExampleOptions) SetHeaders(param map[string]string) *UpdateTrainingExampleOptions {
	options.Headers = param
	return options
}

// WordHeadingDetection : Object containing heading detection conversion settings for Microsoft Word documents.
type WordHeadingDetection struct {
	// Array of font matching configurations.
	Fonts []FontSetting `json:"fonts,omitempty"`

	// Array of Microsoft Word styles to convert.
	Styles []WordStyle `json:"styles,omitempty"`
}

// UnmarshalWordHeadingDetection unmarshals an instance of WordHeadingDetection from the specified map of raw messages.
func UnmarshalWordHeadingDetection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(WordHeadingDetection)
	err = core.UnmarshalModel(m, "fonts", &obj.Fonts, UnmarshalFontSetting)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "styles", &obj.Styles, UnmarshalWordStyle)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// WordSettings : A list of Word conversion settings.
type WordSettings struct {
	// Object containing heading detection conversion settings for Microsoft Word documents.
	Heading *WordHeadingDetection `json:"heading,omitempty"`
}

// UnmarshalWordSettings unmarshals an instance of WordSettings from the specified map of raw messages.
func UnmarshalWordSettings(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(WordSettings)
	err = core.UnmarshalModel(m, "heading", &obj.Heading, UnmarshalWordHeadingDetection)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// WordStyle : Microsoft Word styles to convert into a specified HTML head level.
type WordStyle struct {
	// HTML head level that content matching this style is tagged with.
	Level *int64 `json:"level,omitempty"`

	// Array of word style names to convert.
	Names []string `json:"names,omitempty"`
}

// UnmarshalWordStyle unmarshals an instance of WordStyle from the specified map of raw messages.
func UnmarshalWordStyle(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(WordStyle)
	err = core.UnmarshalPrimitive(m, "level", &obj.Level)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "names", &obj.Names)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// XPathPatterns : Object containing an array of XPaths.
type XPathPatterns struct {
	// An array to XPaths.
	Xpaths []string `json:"xpaths,omitempty"`
}

// UnmarshalXPathPatterns unmarshals an instance of XPathPatterns from the specified map of raw messages.
func UnmarshalXPathPatterns(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(XPathPatterns)
	err = core.UnmarshalPrimitive(m, "xpaths", &obj.Xpaths)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Calculation : Calculation struct
// This model "extends" QueryAggregation
type Calculation struct {
	// The type of aggregation command used. For example: term, filter, max, min, etc.
	Type *string `json:"type,omitempty"`

	// Array of aggregation results.
	Results []AggregationResult `json:"results,omitempty"`

	// Number of matching results.
	MatchingResults *int64 `json:"matching_results,omitempty"`

	// Aggregations returned by Discovery.
	Aggregations []QueryAggregationIntf `json:"aggregations,omitempty"`

	// The field where the aggregation is located in the document.
	Field *string `json:"field,omitempty"`

	// Value of the aggregation.
	Value *float64 `json:"value,omitempty"`
}

func (*Calculation) isaQueryAggregation() bool {
	return true
}

// UnmarshalCalculation unmarshals an instance of Calculation from the specified map of raw messages.
func UnmarshalCalculation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Calculation)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "results", &obj.Results, UnmarshalAggregationResult)
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

// Filter : Filter struct
// This model "extends" QueryAggregation
type Filter struct {
	// The type of aggregation command used. For example: term, filter, max, min, etc.
	Type *string `json:"type,omitempty"`

	// Array of aggregation results.
	Results []AggregationResult `json:"results,omitempty"`

	// Number of matching results.
	MatchingResults *int64 `json:"matching_results,omitempty"`

	// Aggregations returned by Discovery.
	Aggregations []QueryAggregationIntf `json:"aggregations,omitempty"`

	// The match the aggregated results queried for.
	Match *string `json:"match,omitempty"`
}

func (*Filter) isaQueryAggregation() bool {
	return true
}

// UnmarshalFilter unmarshals an instance of Filter from the specified map of raw messages.
func UnmarshalFilter(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Filter)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "results", &obj.Results, UnmarshalAggregationResult)
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
	err = core.UnmarshalPrimitive(m, "match", &obj.Match)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Histogram : Histogram struct
// This model "extends" QueryAggregation
type Histogram struct {
	// The type of aggregation command used. For example: term, filter, max, min, etc.
	Type *string `json:"type,omitempty"`

	// Array of aggregation results.
	Results []AggregationResult `json:"results,omitempty"`

	// Number of matching results.
	MatchingResults *int64 `json:"matching_results,omitempty"`

	// Aggregations returned by Discovery.
	Aggregations []QueryAggregationIntf `json:"aggregations,omitempty"`

	// The field where the aggregation is located in the document.
	Field *string `json:"field,omitempty"`

	// Interval of the aggregation. (For 'histogram' type).
	Interval *int64 `json:"interval,omitempty"`
}

func (*Histogram) isaQueryAggregation() bool {
	return true
}

// UnmarshalHistogram unmarshals an instance of Histogram from the specified map of raw messages.
func UnmarshalHistogram(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Histogram)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "results", &obj.Results, UnmarshalAggregationResult)
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
	err = core.UnmarshalPrimitive(m, "field", &obj.Field)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "interval", &obj.Interval)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Nested : Nested struct
// This model "extends" QueryAggregation
type Nested struct {
	// The type of aggregation command used. For example: term, filter, max, min, etc.
	Type *string `json:"type,omitempty"`

	// Array of aggregation results.
	Results []AggregationResult `json:"results,omitempty"`

	// Number of matching results.
	MatchingResults *int64 `json:"matching_results,omitempty"`

	// Aggregations returned by Discovery.
	Aggregations []QueryAggregationIntf `json:"aggregations,omitempty"`

	// The area of the results the aggregation was restricted to.
	Path *string `json:"path,omitempty"`
}

func (*Nested) isaQueryAggregation() bool {
	return true
}

// UnmarshalNested unmarshals an instance of Nested from the specified map of raw messages.
func UnmarshalNested(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Nested)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "results", &obj.Results, UnmarshalAggregationResult)
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
	err = core.UnmarshalPrimitive(m, "path", &obj.Path)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Term : Term struct
// This model "extends" QueryAggregation
type Term struct {
	// The type of aggregation command used. For example: term, filter, max, min, etc.
	Type *string `json:"type,omitempty"`

	// Array of aggregation results.
	Results []AggregationResult `json:"results,omitempty"`

	// Number of matching results.
	MatchingResults *int64 `json:"matching_results,omitempty"`

	// Aggregations returned by Discovery.
	Aggregations []QueryAggregationIntf `json:"aggregations,omitempty"`

	// The field where the aggregation is located in the document.
	Field *string `json:"field,omitempty"`

	// The number of terms identified.
	Count *int64 `json:"count,omitempty"`
}

func (*Term) isaQueryAggregation() bool {
	return true
}

// UnmarshalTerm unmarshals an instance of Term from the specified map of raw messages.
func UnmarshalTerm(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Term)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "results", &obj.Results, UnmarshalAggregationResult)
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
	err = core.UnmarshalPrimitive(m, "field", &obj.Field)
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

// Timeslice : Timeslice struct
// This model "extends" QueryAggregation
type Timeslice struct {
	// The type of aggregation command used. For example: term, filter, max, min, etc.
	Type *string `json:"type,omitempty"`

	// Array of aggregation results.
	Results []AggregationResult `json:"results,omitempty"`

	// Number of matching results.
	MatchingResults *int64 `json:"matching_results,omitempty"`

	// Aggregations returned by Discovery.
	Aggregations []QueryAggregationIntf `json:"aggregations,omitempty"`

	// The field where the aggregation is located in the document.
	Field *string `json:"field,omitempty"`

	// Interval of the aggregation. Valid date interval values are second/seconds minute/minutes, hour/hours, day/days,
	// week/weeks, month/months, and year/years.
	Interval *string `json:"interval,omitempty"`

	// Used to indicate that anomaly detection should be performed. Anomaly detection is used to locate unusual datapoints
	// within a time series.
	Anomaly *bool `json:"anomaly,omitempty"`
}

func (*Timeslice) isaQueryAggregation() bool {
	return true
}

// UnmarshalTimeslice unmarshals an instance of Timeslice from the specified map of raw messages.
func UnmarshalTimeslice(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Timeslice)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "results", &obj.Results, UnmarshalAggregationResult)
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
	err = core.UnmarshalPrimitive(m, "field", &obj.Field)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "interval", &obj.Interval)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "anomaly", &obj.Anomaly)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TopHits : TopHits struct
// This model "extends" QueryAggregation
type TopHits struct {
	// The type of aggregation command used. For example: term, filter, max, min, etc.
	Type *string `json:"type,omitempty"`

	// Array of aggregation results.
	Results []AggregationResult `json:"results,omitempty"`

	// Number of matching results.
	MatchingResults *int64 `json:"matching_results,omitempty"`

	// Aggregations returned by Discovery.
	Aggregations []QueryAggregationIntf `json:"aggregations,omitempty"`

	// Number of top hits returned by the aggregation.
	Size *int64 `json:"size,omitempty"`

	Hits *TopHitsResults `json:"hits,omitempty"`
}

func (*TopHits) isaQueryAggregation() bool {
	return true
}

// UnmarshalTopHits unmarshals an instance of TopHits from the specified map of raw messages.
func UnmarshalTopHits(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TopHits)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "results", &obj.Results, UnmarshalAggregationResult)
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
	err = core.UnmarshalPrimitive(m, "size", &obj.Size)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "hits", &obj.Hits, UnmarshalTopHitsResults)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
