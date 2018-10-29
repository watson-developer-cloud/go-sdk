// Package discoveryv1 : Operations and models for the DiscoveryV1 service
package discoveryv1

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
	"github.com/go-openapi/strfmt"
	core "github.com/watson-developer-cloud/go-sdk/core"
	"os"
	"strings"
)

// DiscoveryV1 : The IBM Watson&trade; Discovery Service is a cognitive search and content analytics engine that you can
// add to applications to identify patterns, trends and actionable insights to drive better decision-making. Securely
// unify structured and unstructured data with pre-enriched content, and use a simplified query language to eliminate
// the need for manual filtering of results.
//
// Version: V1
// See: http://www.ibm.com/watson/developercloud/discovery.html
type DiscoveryV1 struct {
	Service *core.WatsonService
}

// DiscoveryV1Options : Service options
type DiscoveryV1Options struct {
	Version        string
	URL            string
	Username       string
	Password       string
	IAMApiKey      string
	IAMAccessToken string
	IAMURL         string
}

// NewDiscoveryV1 : Instantiate DiscoveryV1
func NewDiscoveryV1(options *DiscoveryV1Options) (*DiscoveryV1, error) {
	if options.URL == "" {
		options.URL = "https://gateway.watsonplatform.net/discovery/api"
	}

	serviceOptions := &core.ServiceOptions{
		URL:            options.URL,
		Version:        options.Version,
		Username:       options.Username,
		Password:       options.Password,
		IAMApiKey:      options.IAMApiKey,
		IAMAccessToken: options.IAMAccessToken,
		IAMURL:         options.IAMURL,
	}
	service, serviceErr := core.NewWatsonService(serviceOptions, "discovery")
	if serviceErr != nil {
		return nil, serviceErr
	}

	return &DiscoveryV1{Service: service}, nil
}

// CreateEnvironment : Create an environment
// Creates a new environment for private data. An environment must be created before collections can be created.
//
// **Note**: You can create only one environment for private data per service instance. An attempt to create another
// environment results in an error.
func (discovery *DiscoveryV1) CreateEnvironment(createEnvironmentOptions *CreateEnvironmentOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(createEnvironmentOptions, "createEnvironmentOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(createEnvironmentOptions, "createEnvironmentOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range createEnvironmentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

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
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(Environment))
	return response, err
}

// GetCreateEnvironmentResult : Retrieve result of CreateEnvironment operation
func (discovery *DiscoveryV1) GetCreateEnvironmentResult(response *core.DetailedResponse) *Environment {
	result, ok := response.Result.(*Environment)
	if ok {
		return result
	}
	return nil
}

// DeleteEnvironment : Delete environment
func (discovery *DiscoveryV1) DeleteEnvironment(deleteEnvironmentOptions *DeleteEnvironmentOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(deleteEnvironmentOptions, "deleteEnvironmentOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(deleteEnvironmentOptions, "deleteEnvironmentOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments"}
	pathParameters := []string{*deleteEnvironmentOptions.EnvironmentID}

	builder := core.NewRequestBuilder(core.DELETE)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteEnvironmentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(DeleteEnvironmentResponse))
	return response, err
}

// GetDeleteEnvironmentResult : Retrieve result of DeleteEnvironment operation
func (discovery *DiscoveryV1) GetDeleteEnvironmentResult(response *core.DetailedResponse) *DeleteEnvironmentResponse {
	result, ok := response.Result.(*DeleteEnvironmentResponse)
	if ok {
		return result
	}
	return nil
}

// GetEnvironment : Get environment info
func (discovery *DiscoveryV1) GetEnvironment(getEnvironmentOptions *GetEnvironmentOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(getEnvironmentOptions, "getEnvironmentOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(getEnvironmentOptions, "getEnvironmentOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments"}
	pathParameters := []string{*getEnvironmentOptions.EnvironmentID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range getEnvironmentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(Environment))
	return response, err
}

// GetGetEnvironmentResult : Retrieve result of GetEnvironment operation
func (discovery *DiscoveryV1) GetGetEnvironmentResult(response *core.DetailedResponse) *Environment {
	result, ok := response.Result.(*Environment)
	if ok {
		return result
	}
	return nil
}

// ListEnvironments : List environments
// List existing environments for the service instance.
func (discovery *DiscoveryV1) ListEnvironments(listEnvironmentsOptions *ListEnvironmentsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateStruct(listEnvironmentsOptions, "listEnvironmentsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range listEnvironmentsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listEnvironmentsOptions.Name != nil {
		builder.AddQuery("name", fmt.Sprint(*listEnvironmentsOptions.Name))
	}
	builder.AddQuery("version", discovery.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(ListEnvironmentsResponse))
	return response, err
}

// GetListEnvironmentsResult : Retrieve result of ListEnvironments operation
func (discovery *DiscoveryV1) GetListEnvironmentsResult(response *core.DetailedResponse) *ListEnvironmentsResponse {
	result, ok := response.Result.(*ListEnvironmentsResponse)
	if ok {
		return result
	}
	return nil
}

// ListFields : List fields across collections
// Gets a list of the unique fields (and their types) stored in the indexes of the specified collections.
func (discovery *DiscoveryV1) ListFields(listFieldsOptions *ListFieldsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(listFieldsOptions, "listFieldsOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(listFieldsOptions, "listFieldsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "fields"}
	pathParameters := []string{*listFieldsOptions.EnvironmentID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range listFieldsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("collection_ids", strings.Join(listFieldsOptions.CollectionIds, ","))
	builder.AddQuery("version", discovery.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(ListCollectionFieldsResponse))
	return response, err
}

// GetListFieldsResult : Retrieve result of ListFields operation
func (discovery *DiscoveryV1) GetListFieldsResult(response *core.DetailedResponse) *ListCollectionFieldsResponse {
	result, ok := response.Result.(*ListCollectionFieldsResponse)
	if ok {
		return result
	}
	return nil
}

// UpdateEnvironment : Update an environment
// Updates an environment. The environment's **name** and  **description** parameters can be changed. You must specify a
// **name** for the environment.
func (discovery *DiscoveryV1) UpdateEnvironment(updateEnvironmentOptions *UpdateEnvironmentOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(updateEnvironmentOptions, "updateEnvironmentOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(updateEnvironmentOptions, "updateEnvironmentOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments"}
	pathParameters := []string{*updateEnvironmentOptions.EnvironmentID}

	builder := core.NewRequestBuilder(core.PUT)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range updateEnvironmentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

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
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(Environment))
	return response, err
}

// GetUpdateEnvironmentResult : Retrieve result of UpdateEnvironment operation
func (discovery *DiscoveryV1) GetUpdateEnvironmentResult(response *core.DetailedResponse) *Environment {
	result, ok := response.Result.(*Environment)
	if ok {
		return result
	}
	return nil
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
func (discovery *DiscoveryV1) CreateConfiguration(createConfigurationOptions *CreateConfigurationOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(createConfigurationOptions, "createConfigurationOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(createConfigurationOptions, "createConfigurationOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "configurations"}
	pathParameters := []string{*createConfigurationOptions.EnvironmentID}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range createConfigurationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

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
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(Configuration))
	return response, err
}

// GetCreateConfigurationResult : Retrieve result of CreateConfiguration operation
func (discovery *DiscoveryV1) GetCreateConfigurationResult(response *core.DetailedResponse) *Configuration {
	result, ok := response.Result.(*Configuration)
	if ok {
		return result
	}
	return nil
}

// DeleteConfiguration : Delete a configuration
// The deletion is performed unconditionally. A configuration deletion request succeeds even if the configuration is
// referenced by a collection or document ingestion. However, documents that have already been submitted for processing
// continue to use the deleted configuration. Documents are always processed with a snapshot of the configuration as it
// existed at the time the document was submitted.
func (discovery *DiscoveryV1) DeleteConfiguration(deleteConfigurationOptions *DeleteConfigurationOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(deleteConfigurationOptions, "deleteConfigurationOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(deleteConfigurationOptions, "deleteConfigurationOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "configurations"}
	pathParameters := []string{*deleteConfigurationOptions.EnvironmentID, *deleteConfigurationOptions.ConfigurationID}

	builder := core.NewRequestBuilder(core.DELETE)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteConfigurationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(DeleteConfigurationResponse))
	return response, err
}

// GetDeleteConfigurationResult : Retrieve result of DeleteConfiguration operation
func (discovery *DiscoveryV1) GetDeleteConfigurationResult(response *core.DetailedResponse) *DeleteConfigurationResponse {
	result, ok := response.Result.(*DeleteConfigurationResponse)
	if ok {
		return result
	}
	return nil
}

// GetConfiguration : Get configuration details
func (discovery *DiscoveryV1) GetConfiguration(getConfigurationOptions *GetConfigurationOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(getConfigurationOptions, "getConfigurationOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(getConfigurationOptions, "getConfigurationOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "configurations"}
	pathParameters := []string{*getConfigurationOptions.EnvironmentID, *getConfigurationOptions.ConfigurationID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range getConfigurationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(Configuration))
	return response, err
}

// GetGetConfigurationResult : Retrieve result of GetConfiguration operation
func (discovery *DiscoveryV1) GetGetConfigurationResult(response *core.DetailedResponse) *Configuration {
	result, ok := response.Result.(*Configuration)
	if ok {
		return result
	}
	return nil
}

// ListConfigurations : List configurations
// Lists existing configurations for the service instance.
func (discovery *DiscoveryV1) ListConfigurations(listConfigurationsOptions *ListConfigurationsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(listConfigurationsOptions, "listConfigurationsOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(listConfigurationsOptions, "listConfigurationsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "configurations"}
	pathParameters := []string{*listConfigurationsOptions.EnvironmentID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range listConfigurationsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listConfigurationsOptions.Name != nil {
		builder.AddQuery("name", fmt.Sprint(*listConfigurationsOptions.Name))
	}
	builder.AddQuery("version", discovery.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(ListConfigurationsResponse))
	return response, err
}

// GetListConfigurationsResult : Retrieve result of ListConfigurations operation
func (discovery *DiscoveryV1) GetListConfigurationsResult(response *core.DetailedResponse) *ListConfigurationsResponse {
	result, ok := response.Result.(*ListConfigurationsResponse)
	if ok {
		return result
	}
	return nil
}

// UpdateConfiguration : Update a configuration
// Replaces an existing configuration.
//   * Completely replaces the original configuration.
//   * The **configuration_id**, **updated**, and **created** fields are accepted in the request, but they are ignored,
// and an error is not generated. It is also acceptable for users to submit an updated configuration with none of the
// three properties.
//   * Documents are processed with a snapshot of the configuration as it was at the time the document was submitted to
// be ingested. This means that already submitted documents will not see any updates made to the configuration.
func (discovery *DiscoveryV1) UpdateConfiguration(updateConfigurationOptions *UpdateConfigurationOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(updateConfigurationOptions, "updateConfigurationOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(updateConfigurationOptions, "updateConfigurationOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "configurations"}
	pathParameters := []string{*updateConfigurationOptions.EnvironmentID, *updateConfigurationOptions.ConfigurationID}

	builder := core.NewRequestBuilder(core.PUT)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range updateConfigurationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

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
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(Configuration))
	return response, err
}

// GetUpdateConfigurationResult : Retrieve result of UpdateConfiguration operation
func (discovery *DiscoveryV1) GetUpdateConfigurationResult(response *core.DetailedResponse) *Configuration {
	result, ok := response.Result.(*Configuration)
	if ok {
		return result
	}
	return nil
}

// TestConfigurationInEnvironment : Test configuration
// Runs a sample document through the default or your configuration and returns diagnostic information designed to help
// you understand how the document was processed. The document is not added to the index.
func (discovery *DiscoveryV1) TestConfigurationInEnvironment(testConfigurationInEnvironmentOptions *TestConfigurationInEnvironmentOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(testConfigurationInEnvironmentOptions, "testConfigurationInEnvironmentOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(testConfigurationInEnvironmentOptions, "testConfigurationInEnvironmentOptions"); err != nil {
		return nil, err
	}
	if (testConfigurationInEnvironmentOptions.Configuration == nil) && (testConfigurationInEnvironmentOptions.File == nil) && (testConfigurationInEnvironmentOptions.Metadata == nil) {
		return nil, fmt.Errorf("At least one of configuration, file, or metadata must be supplied")
	}

	pathSegments := []string{"v1/environments", "preview"}
	pathParameters := []string{*testConfigurationInEnvironmentOptions.EnvironmentID}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range testConfigurationInEnvironmentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if testConfigurationInEnvironmentOptions.Step != nil {
		builder.AddQuery("step", fmt.Sprint(*testConfigurationInEnvironmentOptions.Step))
	}
	if testConfigurationInEnvironmentOptions.ConfigurationID != nil {
		builder.AddQuery("configuration_id", fmt.Sprint(*testConfigurationInEnvironmentOptions.ConfigurationID))
	}
	builder.AddQuery("version", discovery.Service.Options.Version)

	if testConfigurationInEnvironmentOptions.Configuration != nil {
		builder.AddFormData("configuration", "", "", fmt.Sprint(*testConfigurationInEnvironmentOptions.Configuration))
	}
	if testConfigurationInEnvironmentOptions.File != nil {
		builder.AddFormData("file", core.StringNilMapper(testConfigurationInEnvironmentOptions.Filename),
			core.StringNilMapper(testConfigurationInEnvironmentOptions.FileContentType), testConfigurationInEnvironmentOptions.File)
	}
	if testConfigurationInEnvironmentOptions.Metadata != nil {
		builder.AddFormData("metadata", "", "", fmt.Sprint(*testConfigurationInEnvironmentOptions.Metadata))
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(TestDocument))
	return response, err
}

// GetTestConfigurationInEnvironmentResult : Retrieve result of TestConfigurationInEnvironment operation
func (discovery *DiscoveryV1) GetTestConfigurationInEnvironmentResult(response *core.DetailedResponse) *TestDocument {
	result, ok := response.Result.(*TestDocument)
	if ok {
		return result
	}
	return nil
}

// CreateCollection : Create a collection
func (discovery *DiscoveryV1) CreateCollection(createCollectionOptions *CreateCollectionOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(createCollectionOptions, "createCollectionOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(createCollectionOptions, "createCollectionOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "collections"}
	pathParameters := []string{*createCollectionOptions.EnvironmentID}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range createCollectionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

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
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(Collection))
	return response, err
}

// GetCreateCollectionResult : Retrieve result of CreateCollection operation
func (discovery *DiscoveryV1) GetCreateCollectionResult(response *core.DetailedResponse) *Collection {
	result, ok := response.Result.(*Collection)
	if ok {
		return result
	}
	return nil
}

// DeleteCollection : Delete a collection
func (discovery *DiscoveryV1) DeleteCollection(deleteCollectionOptions *DeleteCollectionOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(deleteCollectionOptions, "deleteCollectionOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(deleteCollectionOptions, "deleteCollectionOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "collections"}
	pathParameters := []string{*deleteCollectionOptions.EnvironmentID, *deleteCollectionOptions.CollectionID}

	builder := core.NewRequestBuilder(core.DELETE)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteCollectionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(DeleteCollectionResponse))
	return response, err
}

// GetDeleteCollectionResult : Retrieve result of DeleteCollection operation
func (discovery *DiscoveryV1) GetDeleteCollectionResult(response *core.DetailedResponse) *DeleteCollectionResponse {
	result, ok := response.Result.(*DeleteCollectionResponse)
	if ok {
		return result
	}
	return nil
}

// GetCollection : Get collection details
func (discovery *DiscoveryV1) GetCollection(getCollectionOptions *GetCollectionOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(getCollectionOptions, "getCollectionOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(getCollectionOptions, "getCollectionOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "collections"}
	pathParameters := []string{*getCollectionOptions.EnvironmentID, *getCollectionOptions.CollectionID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range getCollectionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(Collection))
	return response, err
}

// GetGetCollectionResult : Retrieve result of GetCollection operation
func (discovery *DiscoveryV1) GetGetCollectionResult(response *core.DetailedResponse) *Collection {
	result, ok := response.Result.(*Collection)
	if ok {
		return result
	}
	return nil
}

// ListCollectionFields : List collection fields
// Gets a list of the unique fields (and their types) stored in the index.
func (discovery *DiscoveryV1) ListCollectionFields(listCollectionFieldsOptions *ListCollectionFieldsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(listCollectionFieldsOptions, "listCollectionFieldsOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(listCollectionFieldsOptions, "listCollectionFieldsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "collections", "fields"}
	pathParameters := []string{*listCollectionFieldsOptions.EnvironmentID, *listCollectionFieldsOptions.CollectionID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range listCollectionFieldsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(ListCollectionFieldsResponse))
	return response, err
}

// GetListCollectionFieldsResult : Retrieve result of ListCollectionFields operation
func (discovery *DiscoveryV1) GetListCollectionFieldsResult(response *core.DetailedResponse) *ListCollectionFieldsResponse {
	result, ok := response.Result.(*ListCollectionFieldsResponse)
	if ok {
		return result
	}
	return nil
}

// ListCollections : List collections
// Lists existing collections for the service instance.
func (discovery *DiscoveryV1) ListCollections(listCollectionsOptions *ListCollectionsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(listCollectionsOptions, "listCollectionsOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(listCollectionsOptions, "listCollectionsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "collections"}
	pathParameters := []string{*listCollectionsOptions.EnvironmentID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range listCollectionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listCollectionsOptions.Name != nil {
		builder.AddQuery("name", fmt.Sprint(*listCollectionsOptions.Name))
	}
	builder.AddQuery("version", discovery.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(ListCollectionsResponse))
	return response, err
}

// GetListCollectionsResult : Retrieve result of ListCollections operation
func (discovery *DiscoveryV1) GetListCollectionsResult(response *core.DetailedResponse) *ListCollectionsResponse {
	result, ok := response.Result.(*ListCollectionsResponse)
	if ok {
		return result
	}
	return nil
}

// UpdateCollection : Update a collection
func (discovery *DiscoveryV1) UpdateCollection(updateCollectionOptions *UpdateCollectionOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(updateCollectionOptions, "updateCollectionOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(updateCollectionOptions, "updateCollectionOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "collections"}
	pathParameters := []string{*updateCollectionOptions.EnvironmentID, *updateCollectionOptions.CollectionID}

	builder := core.NewRequestBuilder(core.PUT)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range updateCollectionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

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
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(Collection))
	return response, err
}

// GetUpdateCollectionResult : Retrieve result of UpdateCollection operation
func (discovery *DiscoveryV1) GetUpdateCollectionResult(response *core.DetailedResponse) *Collection {
	result, ok := response.Result.(*Collection)
	if ok {
		return result
	}
	return nil
}

// CreateExpansions : Create or update expansion list
// Create or replace the Expansion list for this collection. The maximum number of expanded terms per collection is
// `500`.
// The current expansion list is replaced with the uploaded content.
func (discovery *DiscoveryV1) CreateExpansions(createExpansionsOptions *CreateExpansionsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(createExpansionsOptions, "createExpansionsOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(createExpansionsOptions, "createExpansionsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "collections", "expansions"}
	pathParameters := []string{*createExpansionsOptions.EnvironmentID, *createExpansionsOptions.CollectionID}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range createExpansionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

	body := make(map[string]interface{})
	if createExpansionsOptions.Expansions != nil {
		body["expansions"] = createExpansionsOptions.Expansions
	}
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(Expansions))
	return response, err
}

// GetCreateExpansionsResult : Retrieve result of CreateExpansions operation
func (discovery *DiscoveryV1) GetCreateExpansionsResult(response *core.DetailedResponse) *Expansions {
	result, ok := response.Result.(*Expansions)
	if ok {
		return result
	}
	return nil
}

// CreateTokenizationDictionary : Create tokenization dictionary
// Upload a custom tokenization dictionary to use with the specified collection.
func (discovery *DiscoveryV1) CreateTokenizationDictionary(createTokenizationDictionaryOptions *CreateTokenizationDictionaryOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(createTokenizationDictionaryOptions, "createTokenizationDictionaryOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(createTokenizationDictionaryOptions, "createTokenizationDictionaryOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "collections", "word_lists/tokenization_dictionary"}
	pathParameters := []string{*createTokenizationDictionaryOptions.EnvironmentID, *createTokenizationDictionaryOptions.CollectionID}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range createTokenizationDictionaryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

	body := make(map[string]interface{})
	if createTokenizationDictionaryOptions.TokenizationRules != nil {
		body["tokenization_rules"] = createTokenizationDictionaryOptions.TokenizationRules
	}
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(TokenDictStatusResponse))
	return response, err
}

// GetCreateTokenizationDictionaryResult : Retrieve result of CreateTokenizationDictionary operation
func (discovery *DiscoveryV1) GetCreateTokenizationDictionaryResult(response *core.DetailedResponse) *TokenDictStatusResponse {
	result, ok := response.Result.(*TokenDictStatusResponse)
	if ok {
		return result
	}
	return nil
}

// DeleteExpansions : Delete the expansion list
// Remove the expansion information for this collection. The expansion list must be deleted to disable query expansion
// for a collection.
func (discovery *DiscoveryV1) DeleteExpansions(deleteExpansionsOptions *DeleteExpansionsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(deleteExpansionsOptions, "deleteExpansionsOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(deleteExpansionsOptions, "deleteExpansionsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "collections", "expansions"}
	pathParameters := []string{*deleteExpansionsOptions.EnvironmentID, *deleteExpansionsOptions.CollectionID}

	builder := core.NewRequestBuilder(core.DELETE)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteExpansionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, nil)
	return response, err
}

// DeleteTokenizationDictionary : Delete tokenization dictionary
// Delete the tokenization dictionary from the collection.
func (discovery *DiscoveryV1) DeleteTokenizationDictionary(deleteTokenizationDictionaryOptions *DeleteTokenizationDictionaryOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(deleteTokenizationDictionaryOptions, "deleteTokenizationDictionaryOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(deleteTokenizationDictionaryOptions, "deleteTokenizationDictionaryOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "collections", "word_lists/tokenization_dictionary"}
	pathParameters := []string{*deleteTokenizationDictionaryOptions.EnvironmentID, *deleteTokenizationDictionaryOptions.CollectionID}

	builder := core.NewRequestBuilder(core.DELETE)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteTokenizationDictionaryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, nil)
	return response, err
}

// GetTokenizationDictionaryStatus : Get tokenization dictionary status
// Returns the current status of the tokenization dictionary for the specified collection.
func (discovery *DiscoveryV1) GetTokenizationDictionaryStatus(getTokenizationDictionaryStatusOptions *GetTokenizationDictionaryStatusOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(getTokenizationDictionaryStatusOptions, "getTokenizationDictionaryStatusOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(getTokenizationDictionaryStatusOptions, "getTokenizationDictionaryStatusOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "collections", "word_lists/tokenization_dictionary"}
	pathParameters := []string{*getTokenizationDictionaryStatusOptions.EnvironmentID, *getTokenizationDictionaryStatusOptions.CollectionID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range getTokenizationDictionaryStatusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(TokenDictStatusResponse))
	return response, err
}

// GetGetTokenizationDictionaryStatusResult : Retrieve result of GetTokenizationDictionaryStatus operation
func (discovery *DiscoveryV1) GetGetTokenizationDictionaryStatusResult(response *core.DetailedResponse) *TokenDictStatusResponse {
	result, ok := response.Result.(*TokenDictStatusResponse)
	if ok {
		return result
	}
	return nil
}

// ListExpansions : Get the expansion list
// Returns the current expansion list for the specified collection. If an expansion list is not specified, an object
// with empty expansion arrays is returned.
func (discovery *DiscoveryV1) ListExpansions(listExpansionsOptions *ListExpansionsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(listExpansionsOptions, "listExpansionsOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(listExpansionsOptions, "listExpansionsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "collections", "expansions"}
	pathParameters := []string{*listExpansionsOptions.EnvironmentID, *listExpansionsOptions.CollectionID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range listExpansionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(Expansions))
	return response, err
}

// GetListExpansionsResult : Retrieve result of ListExpansions operation
func (discovery *DiscoveryV1) GetListExpansionsResult(response *core.DetailedResponse) *Expansions {
	result, ok := response.Result.(*Expansions)
	if ok {
		return result
	}
	return nil
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
//   * Fields containing the following characters after normalization are filtered out before indexing: `#` and `,`.
func (discovery *DiscoveryV1) AddDocument(addDocumentOptions *AddDocumentOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(addDocumentOptions, "addDocumentOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(addDocumentOptions, "addDocumentOptions"); err != nil {
		return nil, err
	}
	if (addDocumentOptions.File == nil) && (addDocumentOptions.Metadata == nil) {
		return nil, fmt.Errorf("At least one of file or metadata must be supplied")
	}

	pathSegments := []string{"v1/environments", "collections", "documents"}
	pathParameters := []string{*addDocumentOptions.EnvironmentID, *addDocumentOptions.CollectionID}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range addDocumentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

	if addDocumentOptions.File != nil {
		builder.AddFormData("file", core.StringNilMapper(addDocumentOptions.Filename),
			core.StringNilMapper(addDocumentOptions.FileContentType), addDocumentOptions.File)
	}
	if addDocumentOptions.Metadata != nil {
		builder.AddFormData("metadata", "", "", fmt.Sprint(*addDocumentOptions.Metadata))
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(DocumentAccepted))
	return response, err
}

// GetAddDocumentResult : Retrieve result of AddDocument operation
func (discovery *DiscoveryV1) GetAddDocumentResult(response *core.DetailedResponse) *DocumentAccepted {
	result, ok := response.Result.(*DocumentAccepted)
	if ok {
		return result
	}
	return nil
}

// DeleteDocument : Delete a document
// If the given document ID is invalid, or if the document is not found, then the a success response is returned (HTTP
// status code `200`) with the status set to 'deleted'.
func (discovery *DiscoveryV1) DeleteDocument(deleteDocumentOptions *DeleteDocumentOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(deleteDocumentOptions, "deleteDocumentOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(deleteDocumentOptions, "deleteDocumentOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "collections", "documents"}
	pathParameters := []string{*deleteDocumentOptions.EnvironmentID, *deleteDocumentOptions.CollectionID, *deleteDocumentOptions.DocumentID}

	builder := core.NewRequestBuilder(core.DELETE)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteDocumentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(DeleteDocumentResponse))
	return response, err
}

// GetDeleteDocumentResult : Retrieve result of DeleteDocument operation
func (discovery *DiscoveryV1) GetDeleteDocumentResult(response *core.DetailedResponse) *DeleteDocumentResponse {
	result, ok := response.Result.(*DeleteDocumentResponse)
	if ok {
		return result
	}
	return nil
}

// GetDocumentStatus : Get document details
// Fetch status details about a submitted document. **Note:** this operation does not return the document itself.
// Instead, it returns only the document's processing status and any notices (warnings or errors) that were generated
// when the document was ingested. Use the query API to retrieve the actual document content.
func (discovery *DiscoveryV1) GetDocumentStatus(getDocumentStatusOptions *GetDocumentStatusOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(getDocumentStatusOptions, "getDocumentStatusOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(getDocumentStatusOptions, "getDocumentStatusOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "collections", "documents"}
	pathParameters := []string{*getDocumentStatusOptions.EnvironmentID, *getDocumentStatusOptions.CollectionID, *getDocumentStatusOptions.DocumentID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range getDocumentStatusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(DocumentStatus))
	return response, err
}

// GetGetDocumentStatusResult : Retrieve result of GetDocumentStatus operation
func (discovery *DiscoveryV1) GetGetDocumentStatusResult(response *core.DetailedResponse) *DocumentStatus {
	result, ok := response.Result.(*DocumentStatus)
	if ok {
		return result
	}
	return nil
}

// UpdateDocument : Update a document
// Replace an existing document. Starts ingesting a document with optional metadata.
func (discovery *DiscoveryV1) UpdateDocument(updateDocumentOptions *UpdateDocumentOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(updateDocumentOptions, "updateDocumentOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(updateDocumentOptions, "updateDocumentOptions"); err != nil {
		return nil, err
	}
	if (updateDocumentOptions.File == nil) && (updateDocumentOptions.Metadata == nil) {
		return nil, fmt.Errorf("At least one of file or metadata must be supplied")
	}

	pathSegments := []string{"v1/environments", "collections", "documents"}
	pathParameters := []string{*updateDocumentOptions.EnvironmentID, *updateDocumentOptions.CollectionID, *updateDocumentOptions.DocumentID}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range updateDocumentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

	if updateDocumentOptions.File != nil {
		builder.AddFormData("file", core.StringNilMapper(updateDocumentOptions.Filename),
			core.StringNilMapper(updateDocumentOptions.FileContentType), updateDocumentOptions.File)
	}
	if updateDocumentOptions.Metadata != nil {
		builder.AddFormData("metadata", "", "", fmt.Sprint(*updateDocumentOptions.Metadata))
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(DocumentAccepted))
	return response, err
}

// GetUpdateDocumentResult : Retrieve result of UpdateDocument operation
func (discovery *DiscoveryV1) GetUpdateDocumentResult(response *core.DetailedResponse) *DocumentAccepted {
	result, ok := response.Result.(*DocumentAccepted)
	if ok {
		return result
	}
	return nil
}

// FederatedQuery : Long environment queries
// Complex queries might be too long for a standard method query. By using this method, you can construct longer
// queries. However, these queries may take longer to complete than the standard method. For details, see the [Discovery
// service documentation](https://console.bluemix.net/docs/services/discovery/using.html).
func (discovery *DiscoveryV1) FederatedQuery(federatedQueryOptions *FederatedQueryOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(federatedQueryOptions, "federatedQueryOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(federatedQueryOptions, "federatedQueryOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "query"}
	pathParameters := []string{*federatedQueryOptions.EnvironmentID}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range federatedQueryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if federatedQueryOptions.LoggingOptOut != nil {
		builder.AddHeader("X-Watson-Logging-Opt-Out", fmt.Sprint(*federatedQueryOptions.LoggingOptOut))
	}
	builder.AddQuery("version", discovery.Service.Options.Version)

	body := make(map[string]interface{})
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
	if federatedQueryOptions.ReturnFields != nil {
		body["return"] = federatedQueryOptions.ReturnFields
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
	if federatedQueryOptions.CollectionIds != nil {
		body["collection_ids"] = federatedQueryOptions.CollectionIds
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
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(QueryResponse))
	return response, err
}

// GetFederatedQueryResult : Retrieve result of FederatedQuery operation
func (discovery *DiscoveryV1) GetFederatedQueryResult(response *core.DetailedResponse) *QueryResponse {
	result, ok := response.Result.(*QueryResponse)
	if ok {
		return result
	}
	return nil
}

// FederatedQueryNotices : Query multiple collection system notices
// Queries for notices (errors or warnings) that might have been generated by the system. Notices are generated when
// ingesting documents and performing relevance training. See the [Discovery service
// documentation](https://console.bluemix.net/docs/services/discovery/using.html) for more details on the query
// language.
func (discovery *DiscoveryV1) FederatedQueryNotices(federatedQueryNoticesOptions *FederatedQueryNoticesOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(federatedQueryNoticesOptions, "federatedQueryNoticesOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(federatedQueryNoticesOptions, "federatedQueryNoticesOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "notices"}
	pathParameters := []string{*federatedQueryNoticesOptions.EnvironmentID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range federatedQueryNoticesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

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
	if federatedQueryNoticesOptions.ReturnFields != nil {
		builder.AddQuery("return", strings.Join(federatedQueryNoticesOptions.ReturnFields, ","))
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
	builder.AddQuery("version", discovery.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(QueryNoticesResponse))
	return response, err
}

// GetFederatedQueryNoticesResult : Retrieve result of FederatedQueryNotices operation
func (discovery *DiscoveryV1) GetFederatedQueryNoticesResult(response *core.DetailedResponse) *QueryNoticesResponse {
	result, ok := response.Result.(*QueryNoticesResponse)
	if ok {
		return result
	}
	return nil
}

// Query : Long collection queries
// Complex queries might be too long for a standard method query. By using this method, you can construct longer
// queries. However, these queries may take longer to complete than the standard method. For details, see the [Discovery
// service documentation](https://console.bluemix.net/docs/services/discovery/using.html).
func (discovery *DiscoveryV1) Query(queryOptions *QueryOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(queryOptions, "queryOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(queryOptions, "queryOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "collections", "query"}
	pathParameters := []string{*queryOptions.EnvironmentID, *queryOptions.CollectionID}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range queryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if queryOptions.LoggingOptOut != nil {
		builder.AddHeader("X-Watson-Logging-Opt-Out", fmt.Sprint(*queryOptions.LoggingOptOut))
	}
	builder.AddQuery("version", discovery.Service.Options.Version)

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
	if queryOptions.ReturnFields != nil {
		body["return"] = queryOptions.ReturnFields
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
	if queryOptions.CollectionIds != nil {
		body["collection_ids"] = queryOptions.CollectionIds
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
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(QueryResponse))
	return response, err
}

// GetQueryResult : Retrieve result of Query operation
func (discovery *DiscoveryV1) GetQueryResult(response *core.DetailedResponse) *QueryResponse {
	result, ok := response.Result.(*QueryResponse)
	if ok {
		return result
	}
	return nil
}

// QueryEntities : Knowledge Graph entity query
// See the [Knowledge Graph documentation](https://console.bluemix.net/docs/services/discovery/building-kg.html) for
// more details.
func (discovery *DiscoveryV1) QueryEntities(queryEntitiesOptions *QueryEntitiesOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(queryEntitiesOptions, "queryEntitiesOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(queryEntitiesOptions, "queryEntitiesOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "collections", "query_entities"}
	pathParameters := []string{*queryEntitiesOptions.EnvironmentID, *queryEntitiesOptions.CollectionID}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range queryEntitiesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

	body := make(map[string]interface{})
	if queryEntitiesOptions.Feature != nil {
		body["feature"] = queryEntitiesOptions.Feature
	}
	if queryEntitiesOptions.Entity != nil {
		body["entity"] = queryEntitiesOptions.Entity
	}
	if queryEntitiesOptions.Context != nil {
		body["context"] = queryEntitiesOptions.Context
	}
	if queryEntitiesOptions.Count != nil {
		body["count"] = queryEntitiesOptions.Count
	}
	if queryEntitiesOptions.EvidenceCount != nil {
		body["evidence_count"] = queryEntitiesOptions.EvidenceCount
	}
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(QueryEntitiesResponse))
	return response, err
}

// GetQueryEntitiesResult : Retrieve result of QueryEntities operation
func (discovery *DiscoveryV1) GetQueryEntitiesResult(response *core.DetailedResponse) *QueryEntitiesResponse {
	result, ok := response.Result.(*QueryEntitiesResponse)
	if ok {
		return result
	}
	return nil
}

// QueryNotices : Query system notices
// Queries for notices (errors or warnings) that might have been generated by the system. Notices are generated when
// ingesting documents and performing relevance training. See the [Discovery service
// documentation](https://console.bluemix.net/docs/services/discovery/using.html) for more details on the query
// language.
func (discovery *DiscoveryV1) QueryNotices(queryNoticesOptions *QueryNoticesOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(queryNoticesOptions, "queryNoticesOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(queryNoticesOptions, "queryNoticesOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "collections", "notices"}
	pathParameters := []string{*queryNoticesOptions.EnvironmentID, *queryNoticesOptions.CollectionID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range queryNoticesOptions.Headers {
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
	if queryNoticesOptions.Passages != nil {
		builder.AddQuery("passages", fmt.Sprint(*queryNoticesOptions.Passages))
	}
	if queryNoticesOptions.Aggregation != nil {
		builder.AddQuery("aggregation", fmt.Sprint(*queryNoticesOptions.Aggregation))
	}
	if queryNoticesOptions.Count != nil {
		builder.AddQuery("count", fmt.Sprint(*queryNoticesOptions.Count))
	}
	if queryNoticesOptions.ReturnFields != nil {
		builder.AddQuery("return", strings.Join(queryNoticesOptions.ReturnFields, ","))
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
	builder.AddQuery("version", discovery.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(QueryNoticesResponse))
	return response, err
}

// GetQueryNoticesResult : Retrieve result of QueryNotices operation
func (discovery *DiscoveryV1) GetQueryNoticesResult(response *core.DetailedResponse) *QueryNoticesResponse {
	result, ok := response.Result.(*QueryNoticesResponse)
	if ok {
		return result
	}
	return nil
}

// QueryRelations : Knowledge Graph relationship query
// See the [Knowledge Graph documentation](https://console.bluemix.net/docs/services/discovery/building-kg.html) for
// more details.
func (discovery *DiscoveryV1) QueryRelations(queryRelationsOptions *QueryRelationsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(queryRelationsOptions, "queryRelationsOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(queryRelationsOptions, "queryRelationsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "collections", "query_relations"}
	pathParameters := []string{*queryRelationsOptions.EnvironmentID, *queryRelationsOptions.CollectionID}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range queryRelationsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

	body := make(map[string]interface{})
	if queryRelationsOptions.Entities != nil {
		body["entities"] = queryRelationsOptions.Entities
	}
	if queryRelationsOptions.Context != nil {
		body["context"] = queryRelationsOptions.Context
	}
	if queryRelationsOptions.Sort != nil {
		body["sort"] = queryRelationsOptions.Sort
	}
	if queryRelationsOptions.Filter != nil {
		body["filter"] = queryRelationsOptions.Filter
	}
	if queryRelationsOptions.Count != nil {
		body["count"] = queryRelationsOptions.Count
	}
	if queryRelationsOptions.EvidenceCount != nil {
		body["evidence_count"] = queryRelationsOptions.EvidenceCount
	}
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(QueryRelationsResponse))
	return response, err
}

// GetQueryRelationsResult : Retrieve result of QueryRelations operation
func (discovery *DiscoveryV1) GetQueryRelationsResult(response *core.DetailedResponse) *QueryRelationsResponse {
	result, ok := response.Result.(*QueryRelationsResponse)
	if ok {
		return result
	}
	return nil
}

// AddTrainingData : Add query to training data
// Adds a query to the training data for this collection. The query can contain a filter and natural language query.
func (discovery *DiscoveryV1) AddTrainingData(addTrainingDataOptions *AddTrainingDataOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(addTrainingDataOptions, "addTrainingDataOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(addTrainingDataOptions, "addTrainingDataOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "collections", "training_data"}
	pathParameters := []string{*addTrainingDataOptions.EnvironmentID, *addTrainingDataOptions.CollectionID}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range addTrainingDataOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

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
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(TrainingQuery))
	return response, err
}

// GetAddTrainingDataResult : Retrieve result of AddTrainingData operation
func (discovery *DiscoveryV1) GetAddTrainingDataResult(response *core.DetailedResponse) *TrainingQuery {
	result, ok := response.Result.(*TrainingQuery)
	if ok {
		return result
	}
	return nil
}

// CreateTrainingExample : Add example to training data query
// Adds a example to this training data query.
func (discovery *DiscoveryV1) CreateTrainingExample(createTrainingExampleOptions *CreateTrainingExampleOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(createTrainingExampleOptions, "createTrainingExampleOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(createTrainingExampleOptions, "createTrainingExampleOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "collections", "training_data", "examples"}
	pathParameters := []string{*createTrainingExampleOptions.EnvironmentID, *createTrainingExampleOptions.CollectionID, *createTrainingExampleOptions.QueryID}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range createTrainingExampleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

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
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(TrainingExample))
	return response, err
}

// GetCreateTrainingExampleResult : Retrieve result of CreateTrainingExample operation
func (discovery *DiscoveryV1) GetCreateTrainingExampleResult(response *core.DetailedResponse) *TrainingExample {
	result, ok := response.Result.(*TrainingExample)
	if ok {
		return result
	}
	return nil
}

// DeleteAllTrainingData : Delete all training data
// Deletes all training data from a collection.
func (discovery *DiscoveryV1) DeleteAllTrainingData(deleteAllTrainingDataOptions *DeleteAllTrainingDataOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(deleteAllTrainingDataOptions, "deleteAllTrainingDataOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(deleteAllTrainingDataOptions, "deleteAllTrainingDataOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "collections", "training_data"}
	pathParameters := []string{*deleteAllTrainingDataOptions.EnvironmentID, *deleteAllTrainingDataOptions.CollectionID}

	builder := core.NewRequestBuilder(core.DELETE)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteAllTrainingDataOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, nil)
	return response, err
}

// DeleteTrainingData : Delete a training data query
// Removes the training data query and all associated examples from the training data set.
func (discovery *DiscoveryV1) DeleteTrainingData(deleteTrainingDataOptions *DeleteTrainingDataOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(deleteTrainingDataOptions, "deleteTrainingDataOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(deleteTrainingDataOptions, "deleteTrainingDataOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "collections", "training_data"}
	pathParameters := []string{*deleteTrainingDataOptions.EnvironmentID, *deleteTrainingDataOptions.CollectionID, *deleteTrainingDataOptions.QueryID}

	builder := core.NewRequestBuilder(core.DELETE)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteTrainingDataOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, nil)
	return response, err
}

// DeleteTrainingExample : Delete example for training data query
// Deletes the example document with the given ID from the training data query.
func (discovery *DiscoveryV1) DeleteTrainingExample(deleteTrainingExampleOptions *DeleteTrainingExampleOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(deleteTrainingExampleOptions, "deleteTrainingExampleOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(deleteTrainingExampleOptions, "deleteTrainingExampleOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "collections", "training_data", "examples"}
	pathParameters := []string{*deleteTrainingExampleOptions.EnvironmentID, *deleteTrainingExampleOptions.CollectionID, *deleteTrainingExampleOptions.QueryID, *deleteTrainingExampleOptions.ExampleID}

	builder := core.NewRequestBuilder(core.DELETE)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteTrainingExampleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, nil)
	return response, err
}

// GetTrainingData : Get details about a query
// Gets details for a specific training data query, including the query string and all examples.
func (discovery *DiscoveryV1) GetTrainingData(getTrainingDataOptions *GetTrainingDataOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(getTrainingDataOptions, "getTrainingDataOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(getTrainingDataOptions, "getTrainingDataOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "collections", "training_data"}
	pathParameters := []string{*getTrainingDataOptions.EnvironmentID, *getTrainingDataOptions.CollectionID, *getTrainingDataOptions.QueryID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range getTrainingDataOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(TrainingQuery))
	return response, err
}

// GetGetTrainingDataResult : Retrieve result of GetTrainingData operation
func (discovery *DiscoveryV1) GetGetTrainingDataResult(response *core.DetailedResponse) *TrainingQuery {
	result, ok := response.Result.(*TrainingQuery)
	if ok {
		return result
	}
	return nil
}

// GetTrainingExample : Get details for training data example
// Gets the details for this training example.
func (discovery *DiscoveryV1) GetTrainingExample(getTrainingExampleOptions *GetTrainingExampleOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(getTrainingExampleOptions, "getTrainingExampleOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(getTrainingExampleOptions, "getTrainingExampleOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "collections", "training_data", "examples"}
	pathParameters := []string{*getTrainingExampleOptions.EnvironmentID, *getTrainingExampleOptions.CollectionID, *getTrainingExampleOptions.QueryID, *getTrainingExampleOptions.ExampleID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range getTrainingExampleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(TrainingExample))
	return response, err
}

// GetGetTrainingExampleResult : Retrieve result of GetTrainingExample operation
func (discovery *DiscoveryV1) GetGetTrainingExampleResult(response *core.DetailedResponse) *TrainingExample {
	result, ok := response.Result.(*TrainingExample)
	if ok {
		return result
	}
	return nil
}

// ListTrainingData : List training data
// Lists the training data for the specified collection.
func (discovery *DiscoveryV1) ListTrainingData(listTrainingDataOptions *ListTrainingDataOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(listTrainingDataOptions, "listTrainingDataOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(listTrainingDataOptions, "listTrainingDataOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "collections", "training_data"}
	pathParameters := []string{*listTrainingDataOptions.EnvironmentID, *listTrainingDataOptions.CollectionID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range listTrainingDataOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(TrainingDataSet))
	return response, err
}

// GetListTrainingDataResult : Retrieve result of ListTrainingData operation
func (discovery *DiscoveryV1) GetListTrainingDataResult(response *core.DetailedResponse) *TrainingDataSet {
	result, ok := response.Result.(*TrainingDataSet)
	if ok {
		return result
	}
	return nil
}

// ListTrainingExamples : List examples for a training data query
// List all examples for this training data query.
func (discovery *DiscoveryV1) ListTrainingExamples(listTrainingExamplesOptions *ListTrainingExamplesOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(listTrainingExamplesOptions, "listTrainingExamplesOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(listTrainingExamplesOptions, "listTrainingExamplesOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "collections", "training_data", "examples"}
	pathParameters := []string{*listTrainingExamplesOptions.EnvironmentID, *listTrainingExamplesOptions.CollectionID, *listTrainingExamplesOptions.QueryID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range listTrainingExamplesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(TrainingExampleList))
	return response, err
}

// GetListTrainingExamplesResult : Retrieve result of ListTrainingExamples operation
func (discovery *DiscoveryV1) GetListTrainingExamplesResult(response *core.DetailedResponse) *TrainingExampleList {
	result, ok := response.Result.(*TrainingExampleList)
	if ok {
		return result
	}
	return nil
}

// UpdateTrainingExample : Change label or cross reference for example
// Changes the label or cross reference query for this training data example.
func (discovery *DiscoveryV1) UpdateTrainingExample(updateTrainingExampleOptions *UpdateTrainingExampleOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(updateTrainingExampleOptions, "updateTrainingExampleOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(updateTrainingExampleOptions, "updateTrainingExampleOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "collections", "training_data", "examples"}
	pathParameters := []string{*updateTrainingExampleOptions.EnvironmentID, *updateTrainingExampleOptions.CollectionID, *updateTrainingExampleOptions.QueryID, *updateTrainingExampleOptions.ExampleID}

	builder := core.NewRequestBuilder(core.PUT)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range updateTrainingExampleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

	body := make(map[string]interface{})
	if updateTrainingExampleOptions.CrossReference != nil {
		body["cross_reference"] = updateTrainingExampleOptions.CrossReference
	}
	if updateTrainingExampleOptions.Relevance != nil {
		body["relevance"] = updateTrainingExampleOptions.Relevance
	}
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(TrainingExample))
	return response, err
}

// GetUpdateTrainingExampleResult : Retrieve result of UpdateTrainingExample operation
func (discovery *DiscoveryV1) GetUpdateTrainingExampleResult(response *core.DetailedResponse) *TrainingExample {
	result, ok := response.Result.(*TrainingExample)
	if ok {
		return result
	}
	return nil
}

// DeleteUserData : Delete labeled data
// Deletes all data associated with a specified customer ID. The method has no effect if no data is associated with the
// customer ID.
//
// You associate a customer ID with data by passing the **X-Watson-Metadata** header with a request that passes data.
// For more information about personal data and customer IDs, see [Information
// security](https://console.bluemix.net/docs/services/discovery/information-security.html).
func (discovery *DiscoveryV1) DeleteUserData(deleteUserDataOptions *DeleteUserDataOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(deleteUserDataOptions, "deleteUserDataOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(deleteUserDataOptions, "deleteUserDataOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/user_data"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.DELETE)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteUserDataOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("customer_id", fmt.Sprint(*deleteUserDataOptions.CustomerID))
	builder.AddQuery("version", discovery.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, nil)
	return response, err
}

// CreateEvent : Create event
// The **Events** API can be used to create log entries that are associated with specific queries. For example, you can
// record which documents in the results set were \"clicked\" by a user and when that click occured.
func (discovery *DiscoveryV1) CreateEvent(createEventOptions *CreateEventOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(createEventOptions, "createEventOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(createEventOptions, "createEventOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/events"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range createEventOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

	body := make(map[string]interface{})
	if createEventOptions.Type != nil {
		body["type"] = createEventOptions.Type
	}
	if createEventOptions.Data != nil {
		body["data"] = createEventOptions.Data
	}
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(CreateEventResponse))
	return response, err
}

// GetCreateEventResult : Retrieve result of CreateEvent operation
func (discovery *DiscoveryV1) GetCreateEventResult(response *core.DetailedResponse) *CreateEventResponse {
	result, ok := response.Result.(*CreateEventResponse)
	if ok {
		return result
	}
	return nil
}

// GetMetricsEventRate : Percentage of queries with an associated event
// The percentage of queries using the **natural_language_query** parameter that have a corresponding \"click\" event
// over a specified time window.  This metric requires having integrated event tracking in your application using the
// **Events** API.
func (discovery *DiscoveryV1) GetMetricsEventRate(getMetricsEventRateOptions *GetMetricsEventRateOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateStruct(getMetricsEventRateOptions, "getMetricsEventRateOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/metrics/event_rate"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range getMetricsEventRateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getMetricsEventRateOptions.StartTime != nil {
		builder.AddQuery("start_time", fmt.Sprint(*getMetricsEventRateOptions.StartTime))
	}
	if getMetricsEventRateOptions.EndTime != nil {
		builder.AddQuery("end_time", fmt.Sprint(*getMetricsEventRateOptions.EndTime))
	}
	if getMetricsEventRateOptions.ResultType != nil {
		builder.AddQuery("result_type", fmt.Sprint(*getMetricsEventRateOptions.ResultType))
	}
	builder.AddQuery("version", discovery.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(MetricResponse))
	return response, err
}

// GetGetMetricsEventRateResult : Retrieve result of GetMetricsEventRate operation
func (discovery *DiscoveryV1) GetGetMetricsEventRateResult(response *core.DetailedResponse) *MetricResponse {
	result, ok := response.Result.(*MetricResponse)
	if ok {
		return result
	}
	return nil
}

// GetMetricsQuery : Number of queries over time
// Total number of queries using the **natural_language_query** parameter over a specific time window.
func (discovery *DiscoveryV1) GetMetricsQuery(getMetricsQueryOptions *GetMetricsQueryOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateStruct(getMetricsQueryOptions, "getMetricsQueryOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/metrics/number_of_queries"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range getMetricsQueryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getMetricsQueryOptions.StartTime != nil {
		builder.AddQuery("start_time", fmt.Sprint(*getMetricsQueryOptions.StartTime))
	}
	if getMetricsQueryOptions.EndTime != nil {
		builder.AddQuery("end_time", fmt.Sprint(*getMetricsQueryOptions.EndTime))
	}
	if getMetricsQueryOptions.ResultType != nil {
		builder.AddQuery("result_type", fmt.Sprint(*getMetricsQueryOptions.ResultType))
	}
	builder.AddQuery("version", discovery.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(MetricResponse))
	return response, err
}

// GetGetMetricsQueryResult : Retrieve result of GetMetricsQuery operation
func (discovery *DiscoveryV1) GetGetMetricsQueryResult(response *core.DetailedResponse) *MetricResponse {
	result, ok := response.Result.(*MetricResponse)
	if ok {
		return result
	}
	return nil
}

// GetMetricsQueryEvent : Number of queries with an event over time
// Total number of queries using the **natural_language_query** parameter that have a corresponding \"click\" event over
// a specified time window. This metric requires having integrated event tracking in your application using the
// **Events** API.
func (discovery *DiscoveryV1) GetMetricsQueryEvent(getMetricsQueryEventOptions *GetMetricsQueryEventOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateStruct(getMetricsQueryEventOptions, "getMetricsQueryEventOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/metrics/number_of_queries_with_event"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range getMetricsQueryEventOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getMetricsQueryEventOptions.StartTime != nil {
		builder.AddQuery("start_time", fmt.Sprint(*getMetricsQueryEventOptions.StartTime))
	}
	if getMetricsQueryEventOptions.EndTime != nil {
		builder.AddQuery("end_time", fmt.Sprint(*getMetricsQueryEventOptions.EndTime))
	}
	if getMetricsQueryEventOptions.ResultType != nil {
		builder.AddQuery("result_type", fmt.Sprint(*getMetricsQueryEventOptions.ResultType))
	}
	builder.AddQuery("version", discovery.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(MetricResponse))
	return response, err
}

// GetGetMetricsQueryEventResult : Retrieve result of GetMetricsQueryEvent operation
func (discovery *DiscoveryV1) GetGetMetricsQueryEventResult(response *core.DetailedResponse) *MetricResponse {
	result, ok := response.Result.(*MetricResponse)
	if ok {
		return result
	}
	return nil
}

// GetMetricsQueryNoResults : Number of queries with no search results over time
// Total number of queries using the **natural_language_query** parameter that have no results returned over a specified
// time window.
func (discovery *DiscoveryV1) GetMetricsQueryNoResults(getMetricsQueryNoResultsOptions *GetMetricsQueryNoResultsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateStruct(getMetricsQueryNoResultsOptions, "getMetricsQueryNoResultsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/metrics/number_of_queries_with_no_search_results"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range getMetricsQueryNoResultsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getMetricsQueryNoResultsOptions.StartTime != nil {
		builder.AddQuery("start_time", fmt.Sprint(*getMetricsQueryNoResultsOptions.StartTime))
	}
	if getMetricsQueryNoResultsOptions.EndTime != nil {
		builder.AddQuery("end_time", fmt.Sprint(*getMetricsQueryNoResultsOptions.EndTime))
	}
	if getMetricsQueryNoResultsOptions.ResultType != nil {
		builder.AddQuery("result_type", fmt.Sprint(*getMetricsQueryNoResultsOptions.ResultType))
	}
	builder.AddQuery("version", discovery.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(MetricResponse))
	return response, err
}

// GetGetMetricsQueryNoResultsResult : Retrieve result of GetMetricsQueryNoResults operation
func (discovery *DiscoveryV1) GetGetMetricsQueryNoResultsResult(response *core.DetailedResponse) *MetricResponse {
	result, ok := response.Result.(*MetricResponse)
	if ok {
		return result
	}
	return nil
}

// GetMetricsQueryTokenEvent : Most frequent query tokens with an event
// The most frequent query tokens parsed from the **natural_language_query** parameter and their corresponding \"click\"
// event rate within the recording period (queries and events are stored for 30 days). A query token is an individual
// word or unigram within the query string.
func (discovery *DiscoveryV1) GetMetricsQueryTokenEvent(getMetricsQueryTokenEventOptions *GetMetricsQueryTokenEventOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateStruct(getMetricsQueryTokenEventOptions, "getMetricsQueryTokenEventOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/metrics/top_query_tokens_with_event_rate"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range getMetricsQueryTokenEventOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getMetricsQueryTokenEventOptions.Count != nil {
		builder.AddQuery("count", fmt.Sprint(*getMetricsQueryTokenEventOptions.Count))
	}
	builder.AddQuery("version", discovery.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(MetricTokenResponse))
	return response, err
}

// GetGetMetricsQueryTokenEventResult : Retrieve result of GetMetricsQueryTokenEvent operation
func (discovery *DiscoveryV1) GetGetMetricsQueryTokenEventResult(response *core.DetailedResponse) *MetricTokenResponse {
	result, ok := response.Result.(*MetricTokenResponse)
	if ok {
		return result
	}
	return nil
}

// QueryLog : Search the query and event log
// Searches the query and event log to find query sessions that match the specified criteria. Searching the **logs**
// endpoint uses the standard Discovery query syntax for the parameters that are supported.
func (discovery *DiscoveryV1) QueryLog(queryLogOptions *QueryLogOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateStruct(queryLogOptions, "queryLogOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/logs"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range queryLogOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

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
	builder.AddQuery("version", discovery.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(LogQueryResponse))
	return response, err
}

// GetQueryLogResult : Retrieve result of QueryLog operation
func (discovery *DiscoveryV1) GetQueryLogResult(response *core.DetailedResponse) *LogQueryResponse {
	result, ok := response.Result.(*LogQueryResponse)
	if ok {
		return result
	}
	return nil
}

// CreateCredentials : Create credentials
// Creates a set of credentials to connect to a remote source. Created credentials are used in a configuration to
// associate a collection with the remote source.
//
// **Note:** All credentials are sent over an encrypted connection and encrypted at rest.
func (discovery *DiscoveryV1) CreateCredentials(createCredentialsOptions *CreateCredentialsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(createCredentialsOptions, "createCredentialsOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(createCredentialsOptions, "createCredentialsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "credentials"}
	pathParameters := []string{*createCredentialsOptions.EnvironmentID}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range createCredentialsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

	body := make(map[string]interface{})
	if createCredentialsOptions.SourceType != nil {
		body["source_type"] = createCredentialsOptions.SourceType
	}
	if createCredentialsOptions.CredentialDetails != nil {
		body["credential_details"] = createCredentialsOptions.CredentialDetails
	}
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(Credentials))
	return response, err
}

// GetCreateCredentialsResult : Retrieve result of CreateCredentials operation
func (discovery *DiscoveryV1) GetCreateCredentialsResult(response *core.DetailedResponse) *Credentials {
	result, ok := response.Result.(*Credentials)
	if ok {
		return result
	}
	return nil
}

// DeleteCredentials : Delete credentials
// Deletes a set of stored credentials from your Discovery instance.
func (discovery *DiscoveryV1) DeleteCredentials(deleteCredentialsOptions *DeleteCredentialsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(deleteCredentialsOptions, "deleteCredentialsOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(deleteCredentialsOptions, "deleteCredentialsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "credentials"}
	pathParameters := []string{*deleteCredentialsOptions.EnvironmentID, *deleteCredentialsOptions.CredentialID}

	builder := core.NewRequestBuilder(core.DELETE)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteCredentialsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(DeleteCredentials))
	return response, err
}

// GetDeleteCredentialsResult : Retrieve result of DeleteCredentials operation
func (discovery *DiscoveryV1) GetDeleteCredentialsResult(response *core.DetailedResponse) *DeleteCredentials {
	result, ok := response.Result.(*DeleteCredentials)
	if ok {
		return result
	}
	return nil
}

// GetCredentials : View Credentials
// Returns details about the specified credentials.
//
//  **Note:** Secure credential information such as a password or SSH key is never returned and must be obtained from
// the source system.
func (discovery *DiscoveryV1) GetCredentials(getCredentialsOptions *GetCredentialsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(getCredentialsOptions, "getCredentialsOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(getCredentialsOptions, "getCredentialsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "credentials"}
	pathParameters := []string{*getCredentialsOptions.EnvironmentID, *getCredentialsOptions.CredentialID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range getCredentialsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(Credentials))
	return response, err
}

// GetGetCredentialsResult : Retrieve result of GetCredentials operation
func (discovery *DiscoveryV1) GetGetCredentialsResult(response *core.DetailedResponse) *Credentials {
	result, ok := response.Result.(*Credentials)
	if ok {
		return result
	}
	return nil
}

// ListCredentials : List credentials
// List all the source credentials that have been created for this service instance.
//
//  **Note:**  All credentials are sent over an encrypted connection and encrypted at rest.
func (discovery *DiscoveryV1) ListCredentials(listCredentialsOptions *ListCredentialsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(listCredentialsOptions, "listCredentialsOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(listCredentialsOptions, "listCredentialsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "credentials"}
	pathParameters := []string{*listCredentialsOptions.EnvironmentID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range listCredentialsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(CredentialsList))
	return response, err
}

// GetListCredentialsResult : Retrieve result of ListCredentials operation
func (discovery *DiscoveryV1) GetListCredentialsResult(response *core.DetailedResponse) *CredentialsList {
	result, ok := response.Result.(*CredentialsList)
	if ok {
		return result
	}
	return nil
}

// UpdateCredentials : Update credentials
// Updates an existing set of source credentials.
//
// **Note:** All credentials are sent over an encrypted connection and encrypted at rest.
func (discovery *DiscoveryV1) UpdateCredentials(updateCredentialsOptions *UpdateCredentialsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(updateCredentialsOptions, "updateCredentialsOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(updateCredentialsOptions, "updateCredentialsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/environments", "credentials"}
	pathParameters := []string{*updateCredentialsOptions.EnvironmentID, *updateCredentialsOptions.CredentialID}

	builder := core.NewRequestBuilder(core.PUT)
	builder.ConstructHTTPURL(discovery.Service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range updateCredentialsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", discovery.Service.Options.Version)

	body := make(map[string]interface{})
	if updateCredentialsOptions.SourceType != nil {
		body["source_type"] = updateCredentialsOptions.SourceType
	}
	if updateCredentialsOptions.CredentialDetails != nil {
		body["credential_details"] = updateCredentialsOptions.CredentialDetails
	}
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := discovery.Service.Request(request, new(Credentials))
	return response, err
}

// GetUpdateCredentialsResult : Retrieve result of UpdateCredentials operation
func (discovery *DiscoveryV1) GetUpdateCredentialsResult(response *core.DetailedResponse) *Credentials {
	result, ok := response.Result.(*Credentials)
	if ok {
		return result
	}
	return nil
}

// AddDocumentOptions : The addDocument options.
type AddDocumentOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The ID of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// The content of the document to ingest. The maximum supported file size is 50 megabytes. Files larger than 50
	// megabytes is rejected.
	File *os.File `json:"file,omitempty"`

	// The filename for file.
	Filename *string `json:"filename,omitempty"`

	// If you're using the Data Crawler to upload your documents, you can test a document against the type of metadata that
	// the Data Crawler might send. The maximum supported metadata file size is 1 MB. Metadata parts larger than 1 MB are
	// rejected.
	// Example:  ``` {
	//   "Creator": "Johnny Appleseed",
	//   "Subject": "Apples"
	// } ```.
	Metadata *string `json:"metadata,omitempty"`

	// The content type of file. Values for this parameter can be obtained from the HttpMediaType class.
	FileContentType *string `json:"file_content_type,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewAddDocumentOptions : Instantiate AddDocumentOptions
func (discovery *DiscoveryV1) NewAddDocumentOptions(environmentID string, collectionID string) *AddDocumentOptions {
	return &AddDocumentOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *AddDocumentOptions) SetEnvironmentID(environmentID string) *AddDocumentOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *AddDocumentOptions) SetCollectionID(collectionID string) *AddDocumentOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
}

// SetFile : Allow user to set File
func (options *AddDocumentOptions) SetFile(file *os.File) *AddDocumentOptions {
	options.File = file
	return options
}

// SetFilename : Allow user to set Filename
func (options *AddDocumentOptions) SetFilename(filename string) *AddDocumentOptions {
	options.Filename = core.StringPtr(filename)
	return options
}

// SetMetadata : Allow user to set Metadata
func (options *AddDocumentOptions) SetMetadata(metadata string) *AddDocumentOptions {
	options.Metadata = core.StringPtr(metadata)
	return options
}

// SetFileContentType : Allow user to set FileContentType
func (options *AddDocumentOptions) SetFileContentType(fileContentType string) *AddDocumentOptions {
	options.FileContentType = core.StringPtr(fileContentType)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *AddDocumentOptions) SetHeaders(param map[string]string) *AddDocumentOptions {
	options.Headers = param
	return options
}

// AddTrainingDataOptions : The addTrainingData options.
type AddTrainingDataOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The ID of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	NaturalLanguageQuery *string `json:"natural_language_query,omitempty"`

	Filter *string `json:"filter,omitempty"`

	Examples []TrainingExample `json:"examples,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewAddTrainingDataOptions : Instantiate AddTrainingDataOptions
func (discovery *DiscoveryV1) NewAddTrainingDataOptions(environmentID string, collectionID string) *AddTrainingDataOptions {
	return &AddTrainingDataOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *AddTrainingDataOptions) SetEnvironmentID(environmentID string) *AddTrainingDataOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *AddTrainingDataOptions) SetCollectionID(collectionID string) *AddTrainingDataOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
}

// SetNaturalLanguageQuery : Allow user to set NaturalLanguageQuery
func (options *AddTrainingDataOptions) SetNaturalLanguageQuery(naturalLanguageQuery string) *AddTrainingDataOptions {
	options.NaturalLanguageQuery = core.StringPtr(naturalLanguageQuery)
	return options
}

// SetFilter : Allow user to set Filter
func (options *AddTrainingDataOptions) SetFilter(filter string) *AddTrainingDataOptions {
	options.Filter = core.StringPtr(filter)
	return options
}

// SetExamples : Allow user to set Examples
func (options *AddTrainingDataOptions) SetExamples(examples []TrainingExample) *AddTrainingDataOptions {
	options.Examples = examples
	return options
}

// SetHeaders : Allow user to set Headers
func (options *AddTrainingDataOptions) SetHeaders(param map[string]string) *AddTrainingDataOptions {
	options.Headers = param
	return options
}

// AggregationResult : AggregationResult struct
type AggregationResult struct {

	// Key that matched the aggregation type.
	Key *string `json:"key,omitempty"`

	// Number of matching results.
	MatchingResults *int64 `json:"matching_results,omitempty"`

	// Aggregations returned in the case of chained aggregations.
	Aggregations []QueryAggregation `json:"aggregations,omitempty"`
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

	// The object providing information about the documents in the collection. Present only when retrieving details of a
	// collection.
	DocumentCounts *DocumentCounts `json:"document_counts,omitempty"`

	// The object providing information about the disk usage of the collection. Present only when retrieving details of a
	// collection.
	DiskUsage *CollectionDiskUsage `json:"disk_usage,omitempty"`

	// Provides information about the status of relevance training for collection.
	TrainingStatus *TrainingStatus `json:"training_status,omitempty"`

	// Object containing source crawl status information.
	SourceCrawl *SourceStatus `json:"source_crawl,omitempty"`
}

// Constants associated with the Collection.Status property.
// The status of the collection.
const (
	Collection_Status_Active      = "active"
	Collection_Status_Maintenance = "maintenance"
	Collection_Status_Pending     = "pending"
)

// CollectionDiskUsage : Summary of the disk usage statistics for this collection.
type CollectionDiskUsage struct {

	// Number of bytes used by the collection.
	UsedBytes *int64 `json:"used_bytes,omitempty"`
}

// CollectionUsage : Summary of the collection usage in the environment.
type CollectionUsage struct {

	// Number of active collections in the environment.
	Available *int64 `json:"available,omitempty"`

	// Total number of collections allowed in the environment.
	MaximumAllowed *int64 `json:"maximum_allowed,omitempty"`
}

// Configuration : A custom configuration for the environment.
type Configuration struct {

	// The unique identifier of the configuration.
	ConfigurationID *string `json:"configuration_id,omitempty"`

	// The name of the configuration.
	Name *string `json:"name,omitempty"`

	// The creation date of the configuration in the format yyyy-MM-dd'T'HH:mm:ss.SSS'Z'.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// The timestamp of when the configuration was last updated in the format yyyy-MM-dd'T'HH:mm:ss.SSS'Z'.
	Updated *strfmt.DateTime `json:"updated,omitempty"`

	// The description of the configuration, if available.
	Description *string `json:"description,omitempty"`

	// The document conversion settings for the configuration.
	Conversions *Conversions `json:"conversions,omitempty"`

	// An array of document enrichment settings for the configuration.
	Enrichments []Enrichment `json:"enrichments,omitempty"`

	// Defines operations that can be used to transform the final output JSON into a normalized form. Operations are
	// executed in the order that they appear in the array.
	Normalizations []NormalizationOperation `json:"normalizations,omitempty"`

	// Object containing source parameters for the configuration.
	Source *Source `json:"source,omitempty"`
}

// Conversions : Document conversion settings.
type Conversions struct {

	// A list of PDF conversion settings.
	Pdf *PdfSettings `json:"pdf,omitempty"`

	// A list of Word conversion settings.
	Word *WordSettings `json:"word,omitempty"`

	// A list of HTML conversion settings.
	HTML *HTMLSettings `json:"html,omitempty"`

	// A list of Document Segmentation settings.
	Segment *SegmentSettings `json:"segment,omitempty"`

	// Defines operations that can be used to transform the final output JSON into a normalized form. Operations are
	// executed in the order that they appear in the array.
	JSONNormalizations []NormalizationOperation `json:"json_normalizations,omitempty"`
}

// CreateCollectionOptions : The createCollection options.
type CreateCollectionOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The name of the collection to be created.
	Name *string `json:"name" validate:"required"`

	// A description of the collection.
	Description *string `json:"description,omitempty"`

	// The ID of the configuration in which the collection is to be created.
	ConfigurationID *string `json:"configuration_id,omitempty"`

	// The language of the documents stored in the collection, in the form of an ISO 639-1 language code.
	Language *string `json:"language,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the CreateCollectionOptions.Language property.
// The language of the documents stored in the collection, in the form of an ISO 639-1 language code.
const (
	CreateCollectionOptions_Language_Ar = "ar"
	CreateCollectionOptions_Language_De = "de"
	CreateCollectionOptions_Language_En = "en"
	CreateCollectionOptions_Language_Es = "es"
	CreateCollectionOptions_Language_Fr = "fr"
	CreateCollectionOptions_Language_It = "it"
	CreateCollectionOptions_Language_Ja = "ja"
	CreateCollectionOptions_Language_Ko = "ko"
	CreateCollectionOptions_Language_Nl = "nl"
	CreateCollectionOptions_Language_Pt = "pt"
)

// NewCreateCollectionOptions : Instantiate CreateCollectionOptions
func (discovery *DiscoveryV1) NewCreateCollectionOptions(environmentID string, name string) *CreateCollectionOptions {
	return &CreateCollectionOptions{
		EnvironmentID: core.StringPtr(environmentID),
		Name:          core.StringPtr(name),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *CreateCollectionOptions) SetEnvironmentID(environmentID string) *CreateCollectionOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
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

// SetConfigurationID : Allow user to set ConfigurationID
func (options *CreateCollectionOptions) SetConfigurationID(configurationID string) *CreateCollectionOptions {
	options.ConfigurationID = core.StringPtr(configurationID)
	return options
}

// SetLanguage : Allow user to set Language
func (options *CreateCollectionOptions) SetLanguage(language string) *CreateCollectionOptions {
	options.Language = core.StringPtr(language)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateCollectionOptions) SetHeaders(param map[string]string) *CreateCollectionOptions {
	options.Headers = param
	return options
}

// CreateConfigurationOptions : The createConfiguration options.
type CreateConfigurationOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The name of the configuration.
	Name *string `json:"name,omitempty"`

	// The description of the configuration, if available.
	Description *string `json:"description,omitempty"`

	// The document conversion settings for the configuration.
	Conversions *Conversions `json:"conversions,omitempty"`

	// An array of document enrichment settings for the configuration.
	Enrichments []Enrichment `json:"enrichments,omitempty"`

	// Defines operations that can be used to transform the final output JSON into a normalized form. Operations are
	// executed in the order that they appear in the array.
	Normalizations []NormalizationOperation `json:"normalizations,omitempty"`

	// Object containing source parameters for the configuration.
	Source *Source `json:"source,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewCreateConfigurationOptions : Instantiate CreateConfigurationOptions
func (discovery *DiscoveryV1) NewCreateConfigurationOptions(environmentID string) *CreateConfigurationOptions {
	return &CreateConfigurationOptions{
		EnvironmentID: core.StringPtr(environmentID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *CreateConfigurationOptions) SetEnvironmentID(environmentID string) *CreateConfigurationOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetName : Allow user to set Name
func (options *CreateConfigurationOptions) SetName(name string) *CreateConfigurationOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetDescription : Allow user to set Description
func (options *CreateConfigurationOptions) SetDescription(description string) *CreateConfigurationOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetConversions : Allow user to set Conversions
func (options *CreateConfigurationOptions) SetConversions(conversions *Conversions) *CreateConfigurationOptions {
	options.Conversions = conversions
	return options
}

// SetEnrichments : Allow user to set Enrichments
func (options *CreateConfigurationOptions) SetEnrichments(enrichments []Enrichment) *CreateConfigurationOptions {
	options.Enrichments = enrichments
	return options
}

// SetNormalizations : Allow user to set Normalizations
func (options *CreateConfigurationOptions) SetNormalizations(normalizations []NormalizationOperation) *CreateConfigurationOptions {
	options.Normalizations = normalizations
	return options
}

// SetSource : Allow user to set Source
func (options *CreateConfigurationOptions) SetSource(source *Source) *CreateConfigurationOptions {
	options.Source = source
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateConfigurationOptions) SetHeaders(param map[string]string) *CreateConfigurationOptions {
	options.Headers = param
	return options
}

// CreateCredentialsOptions : The createCredentials options.
type CreateCredentialsOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The source that this credentials object connects to.
	// -  `box` indicates the credentials are used to connect an instance of Enterprise Box.
	// -  `salesforce` indicates the credentials are used to connect to Salesforce.
	// -  `sharepoint` indicates the credentials are used to connect to Microsoft SharePoint Online.
	SourceType *string `json:"source_type,omitempty"`

	// Object containing details of the stored credentials.
	//
	// Obtain credentials for your source from the administrator of the source.
	CredentialDetails *CredentialDetails `json:"credential_details,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the CreateCredentialsOptions.SourceType property.
// The source that this credentials object connects to.
// -  `box` indicates the credentials are used to connect an instance of Enterprise Box.
// -  `salesforce` indicates the credentials are used to connect to Salesforce.
// -  `sharepoint` indicates the credentials are used to connect to Microsoft SharePoint Online.
const (
	CreateCredentialsOptions_SourceType_Box        = "box"
	CreateCredentialsOptions_SourceType_Salesforce = "salesforce"
	CreateCredentialsOptions_SourceType_Sharepoint = "sharepoint"
)

// NewCreateCredentialsOptions : Instantiate CreateCredentialsOptions
func (discovery *DiscoveryV1) NewCreateCredentialsOptions(environmentID string) *CreateCredentialsOptions {
	return &CreateCredentialsOptions{
		EnvironmentID: core.StringPtr(environmentID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *CreateCredentialsOptions) SetEnvironmentID(environmentID string) *CreateCredentialsOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetSourceType : Allow user to set SourceType
func (options *CreateCredentialsOptions) SetSourceType(sourceType string) *CreateCredentialsOptions {
	options.SourceType = core.StringPtr(sourceType)
	return options
}

// SetCredentialDetails : Allow user to set CredentialDetails
func (options *CreateCredentialsOptions) SetCredentialDetails(credentialDetails *CredentialDetails) *CreateCredentialsOptions {
	options.CredentialDetails = credentialDetails
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateCredentialsOptions) SetHeaders(param map[string]string) *CreateCredentialsOptions {
	options.Headers = param
	return options
}

// CreateEnvironmentOptions : The createEnvironment options.
type CreateEnvironmentOptions struct {

	// Name that identifies the environment.
	Name *string `json:"name" validate:"required"`

	// Description of the environment.
	Description *string `json:"description,omitempty"`

	// Size of the environment. In the Lite plan the default and only accepted value is `LT`, in all other plans the
	// default is `S`.
	Size *string `json:"size,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the CreateEnvironmentOptions.Size property.
// Size of the environment. In the Lite plan the default and only accepted value is `LT`, in all other plans the default
// is `S`.
const (
	CreateEnvironmentOptions_Size_L    = "L"
	CreateEnvironmentOptions_Size_LT   = "LT"
	CreateEnvironmentOptions_Size_M    = "M"
	CreateEnvironmentOptions_Size_ML   = "ML"
	CreateEnvironmentOptions_Size_MS   = "MS"
	CreateEnvironmentOptions_Size_S    = "S"
	CreateEnvironmentOptions_Size_XL   = "XL"
	CreateEnvironmentOptions_Size_XS   = "XS"
	CreateEnvironmentOptions_Size_XXL  = "XXL"
	CreateEnvironmentOptions_Size_XXXL = "XXXL"
)

// NewCreateEnvironmentOptions : Instantiate CreateEnvironmentOptions
func (discovery *DiscoveryV1) NewCreateEnvironmentOptions(name string) *CreateEnvironmentOptions {
	return &CreateEnvironmentOptions{
		Name: core.StringPtr(name),
	}
}

// SetName : Allow user to set Name
func (options *CreateEnvironmentOptions) SetName(name string) *CreateEnvironmentOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetDescription : Allow user to set Description
func (options *CreateEnvironmentOptions) SetDescription(description string) *CreateEnvironmentOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetSize : Allow user to set Size
func (options *CreateEnvironmentOptions) SetSize(size string) *CreateEnvironmentOptions {
	options.Size = core.StringPtr(size)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateEnvironmentOptions) SetHeaders(param map[string]string) *CreateEnvironmentOptions {
	options.Headers = param
	return options
}

// CreateEventOptions : The createEvent options.
type CreateEventOptions struct {

	// The event type to be created.
	Type *string `json:"type" validate:"required"`

	// Data object used to create a query event.
	Data *EventData `json:"data" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the CreateEventOptions.Type property.
// The event type to be created.
const (
	CreateEventOptions_Type_Click = "click"
)

// NewCreateEventOptions : Instantiate CreateEventOptions
func (discovery *DiscoveryV1) NewCreateEventOptions(typeVar string, data *EventData) *CreateEventOptions {
	return &CreateEventOptions{
		Type: core.StringPtr(typeVar),
		Data: data,
	}
}

// SetType : Allow user to set Type
func (options *CreateEventOptions) SetType(typeVar string) *CreateEventOptions {
	options.Type = core.StringPtr(typeVar)
	return options
}

// SetData : Allow user to set Data
func (options *CreateEventOptions) SetData(data *EventData) *CreateEventOptions {
	options.Data = data
	return options
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
	CreateEventResponse_Type_Click = "click"
)

// CreateExpansionsOptions : The createExpansions options.
type CreateExpansionsOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The ID of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

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
	Expansions []Expansion `json:"expansions,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewCreateExpansionsOptions : Instantiate CreateExpansionsOptions
func (discovery *DiscoveryV1) NewCreateExpansionsOptions(environmentID string, collectionID string) *CreateExpansionsOptions {
	return &CreateExpansionsOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *CreateExpansionsOptions) SetEnvironmentID(environmentID string) *CreateExpansionsOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *CreateExpansionsOptions) SetCollectionID(collectionID string) *CreateExpansionsOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
}

// SetExpansions : Allow user to set Expansions
func (options *CreateExpansionsOptions) SetExpansions(expansions []Expansion) *CreateExpansionsOptions {
	options.Expansions = expansions
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateExpansionsOptions) SetHeaders(param map[string]string) *CreateExpansionsOptions {
	options.Headers = param
	return options
}

// CreateTokenizationDictionaryOptions : The createTokenizationDictionary options.
type CreateTokenizationDictionaryOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The ID of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// An array of tokenization rules. Each rule contains, the original `text` string, component `tokens`, any alternate
	// character set `readings`, and which `part_of_speech` the text is from.
	TokenizationRules []TokenDictRule `json:"tokenization_rules,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewCreateTokenizationDictionaryOptions : Instantiate CreateTokenizationDictionaryOptions
func (discovery *DiscoveryV1) NewCreateTokenizationDictionaryOptions(environmentID string, collectionID string) *CreateTokenizationDictionaryOptions {
	return &CreateTokenizationDictionaryOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *CreateTokenizationDictionaryOptions) SetEnvironmentID(environmentID string) *CreateTokenizationDictionaryOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *CreateTokenizationDictionaryOptions) SetCollectionID(collectionID string) *CreateTokenizationDictionaryOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
}

// SetTokenizationRules : Allow user to set TokenizationRules
func (options *CreateTokenizationDictionaryOptions) SetTokenizationRules(tokenizationRules []TokenDictRule) *CreateTokenizationDictionaryOptions {
	options.TokenizationRules = tokenizationRules
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateTokenizationDictionaryOptions) SetHeaders(param map[string]string) *CreateTokenizationDictionaryOptions {
	options.Headers = param
	return options
}

// CreateTrainingExampleOptions : The createTrainingExample options.
type CreateTrainingExampleOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The ID of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// The ID of the query used for training.
	QueryID *string `json:"query_id" validate:"required"`

	DocumentID *string `json:"document_id,omitempty"`

	CrossReference *string `json:"cross_reference,omitempty"`

	Relevance *int64 `json:"relevance,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewCreateTrainingExampleOptions : Instantiate CreateTrainingExampleOptions
func (discovery *DiscoveryV1) NewCreateTrainingExampleOptions(environmentID string, collectionID string, queryID string) *CreateTrainingExampleOptions {
	return &CreateTrainingExampleOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
		QueryID:       core.StringPtr(queryID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *CreateTrainingExampleOptions) SetEnvironmentID(environmentID string) *CreateTrainingExampleOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *CreateTrainingExampleOptions) SetCollectionID(collectionID string) *CreateTrainingExampleOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
}

// SetQueryID : Allow user to set QueryID
func (options *CreateTrainingExampleOptions) SetQueryID(queryID string) *CreateTrainingExampleOptions {
	options.QueryID = core.StringPtr(queryID)
	return options
}

// SetDocumentID : Allow user to set DocumentID
func (options *CreateTrainingExampleOptions) SetDocumentID(documentID string) *CreateTrainingExampleOptions {
	options.DocumentID = core.StringPtr(documentID)
	return options
}

// SetCrossReference : Allow user to set CrossReference
func (options *CreateTrainingExampleOptions) SetCrossReference(crossReference string) *CreateTrainingExampleOptions {
	options.CrossReference = core.StringPtr(crossReference)
	return options
}

// SetRelevance : Allow user to set Relevance
func (options *CreateTrainingExampleOptions) SetRelevance(relevance int64) *CreateTrainingExampleOptions {
	options.Relevance = core.Int64Ptr(relevance)
	return options
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
	// -  `"source_type": "sharepoint"` - valid `credential_type`s: `saml`.
	CredentialType *string `json:"credential_type,omitempty"`

	// The **client_id** of the source that these credentials connect to. Only valid, and required, with a
	// **credential_type** of `oauth2`.
	ClientID *string `json:"client_id,omitempty"`

	// The **enterprise_id** of the Box site that these credentials connect to. Only valid, and required, with a
	// **source_type** of `box`.
	EnterpriseID *string `json:"enterprise_id,omitempty"`

	// The **url** of the source that these credentials connect to. Only valid, and required, with a **credential_type** of
	// `username_password`.
	URL *string `json:"url,omitempty"`

	// The **username** of the source that these credentials connect to. Only valid, and required, with a
	// **credential_type** of `saml` and `username_password`.
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
	// **credential_type**s of `saml` and `username_password`.
	//
	// **Note:** When used with a **source_type** of `salesforce`, the password consists of the Salesforce password and a
	// valid Salesforce security token concatenated. This value is never returned and is only used when creating or
	// modifying **credentials**.
	Password *string `json:"password,omitempty"`
}

// Constants associated with the CredentialDetails.CredentialType property.
// The authentication method for this credentials definition. The  **credential_type** specified must be supported by
// the **source_type**. The following combinations are possible:
//
// -  `"source_type": "box"` - valid `credential_type`s: `oauth2`
// -  `"source_type": "salesforce"` - valid `credential_type`s: `username_password`
// -  `"source_type": "sharepoint"` - valid `credential_type`s: `saml`.
const (
	CredentialDetails_CredentialType_Oauth2           = "oauth2"
	CredentialDetails_CredentialType_Saml             = "saml"
	CredentialDetails_CredentialType_UsernamePassword = "username_password"
)

// Credentials : Object containing credential information.
type Credentials struct {

	// Unique identifier for this set of credentials.
	CredentialID *string `json:"credential_id,omitempty"`

	// The source that this credentials object connects to.
	// -  `box` indicates the credentials are used to connect an instance of Enterprise Box.
	// -  `salesforce` indicates the credentials are used to connect to Salesforce.
	// -  `sharepoint` indicates the credentials are used to connect to Microsoft SharePoint Online.
	SourceType *string `json:"source_type,omitempty"`

	// Object containing details of the stored credentials.
	//
	// Obtain credentials for your source from the administrator of the source.
	CredentialDetails *CredentialDetails `json:"credential_details,omitempty"`
}

// Constants associated with the Credentials.SourceType property.
// The source that this credentials object connects to.
// -  `box` indicates the credentials are used to connect an instance of Enterprise Box.
// -  `salesforce` indicates the credentials are used to connect to Salesforce.
// -  `sharepoint` indicates the credentials are used to connect to Microsoft SharePoint Online.
const (
	Credentials_SourceType_Box        = "box"
	Credentials_SourceType_Salesforce = "salesforce"
	Credentials_SourceType_Sharepoint = "sharepoint"
)

// CredentialsList : CredentialsList struct
type CredentialsList struct {

	// An array of credential definitions that were created for this instance.
	Credentials []Credentials `json:"credentials,omitempty"`
}

// DeleteAllTrainingDataOptions : The deleteAllTrainingData options.
type DeleteAllTrainingDataOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The ID of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteAllTrainingDataOptions : Instantiate DeleteAllTrainingDataOptions
func (discovery *DiscoveryV1) NewDeleteAllTrainingDataOptions(environmentID string, collectionID string) *DeleteAllTrainingDataOptions {
	return &DeleteAllTrainingDataOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *DeleteAllTrainingDataOptions) SetEnvironmentID(environmentID string) *DeleteAllTrainingDataOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *DeleteAllTrainingDataOptions) SetCollectionID(collectionID string) *DeleteAllTrainingDataOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteAllTrainingDataOptions) SetHeaders(param map[string]string) *DeleteAllTrainingDataOptions {
	options.Headers = param
	return options
}

// DeleteCollectionOptions : The deleteCollection options.
type DeleteCollectionOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The ID of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteCollectionOptions : Instantiate DeleteCollectionOptions
func (discovery *DiscoveryV1) NewDeleteCollectionOptions(environmentID string, collectionID string) *DeleteCollectionOptions {
	return &DeleteCollectionOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *DeleteCollectionOptions) SetEnvironmentID(environmentID string) *DeleteCollectionOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
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

// DeleteCollectionResponse : DeleteCollectionResponse struct
type DeleteCollectionResponse struct {

	// The unique identifier of the collection that is being deleted.
	CollectionID *string `json:"collection_id" validate:"required"`

	// The status of the collection. The status of a successful deletion operation is `deleted`.
	Status *string `json:"status" validate:"required"`
}

// Constants associated with the DeleteCollectionResponse.Status property.
// The status of the collection. The status of a successful deletion operation is `deleted`.
const (
	DeleteCollectionResponse_Status_Deleted = "deleted"
)

// DeleteConfigurationOptions : The deleteConfiguration options.
type DeleteConfigurationOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The ID of the configuration.
	ConfigurationID *string `json:"configuration_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteConfigurationOptions : Instantiate DeleteConfigurationOptions
func (discovery *DiscoveryV1) NewDeleteConfigurationOptions(environmentID string, configurationID string) *DeleteConfigurationOptions {
	return &DeleteConfigurationOptions{
		EnvironmentID:   core.StringPtr(environmentID),
		ConfigurationID: core.StringPtr(configurationID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *DeleteConfigurationOptions) SetEnvironmentID(environmentID string) *DeleteConfigurationOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetConfigurationID : Allow user to set ConfigurationID
func (options *DeleteConfigurationOptions) SetConfigurationID(configurationID string) *DeleteConfigurationOptions {
	options.ConfigurationID = core.StringPtr(configurationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteConfigurationOptions) SetHeaders(param map[string]string) *DeleteConfigurationOptions {
	options.Headers = param
	return options
}

// DeleteConfigurationResponse : DeleteConfigurationResponse struct
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
	DeleteConfigurationResponse_Status_Deleted = "deleted"
)

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
	DeleteCredentials_Status_Deleted = "deleted"
)

// DeleteCredentialsOptions : The deleteCredentials options.
type DeleteCredentialsOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The unique identifier for a set of source credentials.
	CredentialID *string `json:"credential_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteCredentialsOptions : Instantiate DeleteCredentialsOptions
func (discovery *DiscoveryV1) NewDeleteCredentialsOptions(environmentID string, credentialID string) *DeleteCredentialsOptions {
	return &DeleteCredentialsOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CredentialID:  core.StringPtr(credentialID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *DeleteCredentialsOptions) SetEnvironmentID(environmentID string) *DeleteCredentialsOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetCredentialID : Allow user to set CredentialID
func (options *DeleteCredentialsOptions) SetCredentialID(credentialID string) *DeleteCredentialsOptions {
	options.CredentialID = core.StringPtr(credentialID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteCredentialsOptions) SetHeaders(param map[string]string) *DeleteCredentialsOptions {
	options.Headers = param
	return options
}

// DeleteDocumentOptions : The deleteDocument options.
type DeleteDocumentOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The ID of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// The ID of the document.
	DocumentID *string `json:"document_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteDocumentOptions : Instantiate DeleteDocumentOptions
func (discovery *DiscoveryV1) NewDeleteDocumentOptions(environmentID string, collectionID string, documentID string) *DeleteDocumentOptions {
	return &DeleteDocumentOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
		DocumentID:    core.StringPtr(documentID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *DeleteDocumentOptions) SetEnvironmentID(environmentID string) *DeleteDocumentOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
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

// SetHeaders : Allow user to set Headers
func (options *DeleteDocumentOptions) SetHeaders(param map[string]string) *DeleteDocumentOptions {
	options.Headers = param
	return options
}

// DeleteDocumentResponse : DeleteDocumentResponse struct
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

// DeleteEnvironmentOptions : The deleteEnvironment options.
type DeleteEnvironmentOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteEnvironmentOptions : Instantiate DeleteEnvironmentOptions
func (discovery *DiscoveryV1) NewDeleteEnvironmentOptions(environmentID string) *DeleteEnvironmentOptions {
	return &DeleteEnvironmentOptions{
		EnvironmentID: core.StringPtr(environmentID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *DeleteEnvironmentOptions) SetEnvironmentID(environmentID string) *DeleteEnvironmentOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteEnvironmentOptions) SetHeaders(param map[string]string) *DeleteEnvironmentOptions {
	options.Headers = param
	return options
}

// DeleteEnvironmentResponse : DeleteEnvironmentResponse struct
type DeleteEnvironmentResponse struct {

	// The unique identifier for the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// Status of the environment.
	Status *string `json:"status" validate:"required"`
}

// Constants associated with the DeleteEnvironmentResponse.Status property.
// Status of the environment.
const (
	DeleteEnvironmentResponse_Status_Deleted = "deleted"
)

// DeleteExpansionsOptions : The deleteExpansions options.
type DeleteExpansionsOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The ID of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteExpansionsOptions : Instantiate DeleteExpansionsOptions
func (discovery *DiscoveryV1) NewDeleteExpansionsOptions(environmentID string, collectionID string) *DeleteExpansionsOptions {
	return &DeleteExpansionsOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *DeleteExpansionsOptions) SetEnvironmentID(environmentID string) *DeleteExpansionsOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *DeleteExpansionsOptions) SetCollectionID(collectionID string) *DeleteExpansionsOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteExpansionsOptions) SetHeaders(param map[string]string) *DeleteExpansionsOptions {
	options.Headers = param
	return options
}

// DeleteTokenizationDictionaryOptions : The deleteTokenizationDictionary options.
type DeleteTokenizationDictionaryOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The ID of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteTokenizationDictionaryOptions : Instantiate DeleteTokenizationDictionaryOptions
func (discovery *DiscoveryV1) NewDeleteTokenizationDictionaryOptions(environmentID string, collectionID string) *DeleteTokenizationDictionaryOptions {
	return &DeleteTokenizationDictionaryOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *DeleteTokenizationDictionaryOptions) SetEnvironmentID(environmentID string) *DeleteTokenizationDictionaryOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *DeleteTokenizationDictionaryOptions) SetCollectionID(collectionID string) *DeleteTokenizationDictionaryOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteTokenizationDictionaryOptions) SetHeaders(param map[string]string) *DeleteTokenizationDictionaryOptions {
	options.Headers = param
	return options
}

// DeleteTrainingDataOptions : The deleteTrainingData options.
type DeleteTrainingDataOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The ID of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// The ID of the query used for training.
	QueryID *string `json:"query_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteTrainingDataOptions : Instantiate DeleteTrainingDataOptions
func (discovery *DiscoveryV1) NewDeleteTrainingDataOptions(environmentID string, collectionID string, queryID string) *DeleteTrainingDataOptions {
	return &DeleteTrainingDataOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
		QueryID:       core.StringPtr(queryID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *DeleteTrainingDataOptions) SetEnvironmentID(environmentID string) *DeleteTrainingDataOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *DeleteTrainingDataOptions) SetCollectionID(collectionID string) *DeleteTrainingDataOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
}

// SetQueryID : Allow user to set QueryID
func (options *DeleteTrainingDataOptions) SetQueryID(queryID string) *DeleteTrainingDataOptions {
	options.QueryID = core.StringPtr(queryID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteTrainingDataOptions) SetHeaders(param map[string]string) *DeleteTrainingDataOptions {
	options.Headers = param
	return options
}

// DeleteTrainingExampleOptions : The deleteTrainingExample options.
type DeleteTrainingExampleOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The ID of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// The ID of the query used for training.
	QueryID *string `json:"query_id" validate:"required"`

	// The ID of the document as it is indexed.
	ExampleID *string `json:"example_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteTrainingExampleOptions : Instantiate DeleteTrainingExampleOptions
func (discovery *DiscoveryV1) NewDeleteTrainingExampleOptions(environmentID string, collectionID string, queryID string, exampleID string) *DeleteTrainingExampleOptions {
	return &DeleteTrainingExampleOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
		QueryID:       core.StringPtr(queryID),
		ExampleID:     core.StringPtr(exampleID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *DeleteTrainingExampleOptions) SetEnvironmentID(environmentID string) *DeleteTrainingExampleOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *DeleteTrainingExampleOptions) SetCollectionID(collectionID string) *DeleteTrainingExampleOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
}

// SetQueryID : Allow user to set QueryID
func (options *DeleteTrainingExampleOptions) SetQueryID(queryID string) *DeleteTrainingExampleOptions {
	options.QueryID = core.StringPtr(queryID)
	return options
}

// SetExampleID : Allow user to set ExampleID
func (options *DeleteTrainingExampleOptions) SetExampleID(exampleID string) *DeleteTrainingExampleOptions {
	options.ExampleID = core.StringPtr(exampleID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteTrainingExampleOptions) SetHeaders(param map[string]string) *DeleteTrainingExampleOptions {
	options.Headers = param
	return options
}

// DeleteUserDataOptions : The deleteUserData options.
type DeleteUserDataOptions struct {

	// The customer ID for which all data is to be deleted.
	CustomerID *string `json:"customer_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteUserDataOptions : Instantiate DeleteUserDataOptions
func (discovery *DiscoveryV1) NewDeleteUserDataOptions(customerID string) *DeleteUserDataOptions {
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

// DiskUsage : Summary of the disk usage statistics for the environment.
type DiskUsage struct {

	// Number of bytes within the environment's disk capacity that are currently used to store data.
	UsedBytes *int64 `json:"used_bytes,omitempty"`

	// Total number of bytes available in the environment's disk capacity.
	MaximumAllowedBytes *int64 `json:"maximum_allowed_bytes,omitempty"`

	// **Deprecated**: Total number of bytes available in the environment's disk capacity.
	TotalBytes *int64 `json:"total_bytes,omitempty"`

	// **Deprecated**: Amount of disk capacity used, in KB or GB format.
	Used *string `json:"used,omitempty"`

	// **Deprecated**: Total amount of the environment's disk capacity, in KB or GB format.
	Total *string `json:"total,omitempty"`

	// **Deprecated**: Percentage of the environment's disk capacity that is being used.
	PercentUsed *float64 `json:"percent_used,omitempty"`
}

// DocumentAccepted : DocumentAccepted struct
type DocumentAccepted struct {

	// The unique identifier of the ingested document.
	DocumentID *string `json:"document_id,omitempty"`

	// Status of the document in the ingestion process.
	Status *string `json:"status,omitempty"`

	// Array of notices produced by the document-ingestion process.
	Notices []Notice `json:"notices,omitempty"`
}

// Constants associated with the DocumentAccepted.Status property.
// Status of the document in the ingestion process.
const (
	DocumentAccepted_Status_Processing = "processing"
)

// DocumentCounts : DocumentCounts struct
type DocumentCounts struct {

	// The total number of available documents in the collection.
	Available *int64 `json:"available,omitempty"`

	// The number of documents in the collection that are currently being processed.
	Processing *int64 `json:"processing,omitempty"`

	// The number of documents in the collection that failed to be ingested.
	Failed *int64 `json:"failed,omitempty"`
}

// DocumentSnapshot : DocumentSnapshot struct
type DocumentSnapshot struct {
	Step *string `json:"step,omitempty"`

	Snapshot interface{} `json:"snapshot,omitempty"`
}

// Constants associated with the DocumentSnapshot.Step property.
const (
	DocumentSnapshot_Step_EnrichmentsOutput        = "enrichments_output"
	DocumentSnapshot_Step_HTMLInput                = "html_input"
	DocumentSnapshot_Step_HTMLOutput               = "html_output"
	DocumentSnapshot_Step_JSONNormalizationsOutput = "json_normalizations_output"
	DocumentSnapshot_Step_JSONOutput               = "json_output"
	DocumentSnapshot_Step_NormalizationsOutput     = "normalizations_output"
)

// DocumentStatus : Status information about a submitted document.
type DocumentStatus struct {

	// The unique identifier of the document.
	DocumentID *string `json:"document_id" validate:"required"`

	// The unique identifier for the configuration.
	ConfigurationID *string `json:"configuration_id,omitempty"`

	// The creation date of the document in the format yyyy-MM-dd'T'HH:mm:ss.SSS'Z'.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// Date of the most recent document update, in the format yyyy-MM-dd'T'HH:mm:ss.SSS'Z'.
	Updated *strfmt.DateTime `json:"updated,omitempty"`

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
	DocumentStatus_Status_Available            = "available"
	DocumentStatus_Status_AvailableWithNotices = "available with notices"
	DocumentStatus_Status_Failed               = "failed"
	DocumentStatus_Status_Processing           = "processing"
)

// Constants associated with the DocumentStatus.FileType property.
// The type of the original source file.
const (
	DocumentStatus_FileType_HTML = "html"
	DocumentStatus_FileType_JSON = "json"
	DocumentStatus_FileType_Pdf  = "pdf"
	DocumentStatus_FileType_Word = "word"
)

// Enrichment : Enrichment struct
type Enrichment struct {

	// Describes what the enrichment step does.
	Description *string `json:"description,omitempty"`

	// Field where enrichments will be stored. This field must already exist or be at most 1 level deeper than an existing
	// field. For example, if `text` is a top-level field with no sub-fields, `text.foo` is a valid destination but
	// `text.foo.bar` is not.
	DestinationField *string `json:"destination_field" validate:"required"`

	// Field to be enriched.
	SourceField *string `json:"source_field" validate:"required"`

	// Indicates that the enrichments will overwrite the destination_field field if it already exists.
	Overwrite *bool `json:"overwrite,omitempty"`

	// Name of the enrichment service to call. Current options are `natural_language_understanding` and `elements`.
	//
	//  When using `natual_language_understanding`, the **options** object must contain Natural Language Understanding
	// options.
	//
	//  When using `elements` the **options** object must contain Element Classification options. Additionally, when using
	// the `elements` enrichment the configuration specified and files ingested must meet all the criteria specified in
	// [the documentation](https://console.bluemix.net/docs/services/discovery/element-classification.html)
	//
	//
	//
	//  Previous API versions also supported `alchemy_language`.
	EnrichmentName *string `json:"enrichment" validate:"required"`

	// If true, then most errors generated during the enrichment process will be treated as warnings and will not cause the
	// document to fail processing.
	IgnoreDownstreamErrors *bool `json:"ignore_downstream_errors,omitempty"`

	// A list of options specific to the enrichment.
	Options *EnrichmentOptions `json:"options,omitempty"`
}

// EnrichmentOptions : Options which are specific to a particular enrichment.
type EnrichmentOptions struct {

	// An object representing the enrichment features that will be applied to the specified field.
	Features *NluEnrichmentFeatures `json:"features,omitempty"`

	// ISO 639-1 code indicating the language to use for the analysis. This code overrides the automatic language detection
	// performed by the service. Valid codes are `ar` (Arabic), `en` (English), `fr` (French), `de` (German), `it`
	// (Italian), `pt` (Portuguese), `ru` (Russian), `es` (Spanish), and `sv` (Swedish). **Note:** Not all features support
	// all languages, automatic detection is recommended.
	Language *string `json:"language,omitempty"`

	// *For use with `elements` enrichments only.* The element extraction model to use. Models available are: `contract`.
	Model *string `json:"model,omitempty"`
}

// Constants associated with the EnrichmentOptions.Language property.
// ISO 639-1 code indicating the language to use for the analysis. This code overrides the automatic language detection
// performed by the service. Valid codes are `ar` (Arabic), `en` (English), `fr` (French), `de` (German), `it`
// (Italian), `pt` (Portuguese), `ru` (Russian), `es` (Spanish), and `sv` (Swedish). **Note:** Not all features support
// all languages, automatic detection is recommended.
const (
	EnrichmentOptions_Language_Ar = "ar"
	EnrichmentOptions_Language_De = "de"
	EnrichmentOptions_Language_En = "en"
	EnrichmentOptions_Language_Es = "es"
	EnrichmentOptions_Language_Fr = "fr"
	EnrichmentOptions_Language_It = "it"
	EnrichmentOptions_Language_Pt = "pt"
	EnrichmentOptions_Language_Ru = "ru"
	EnrichmentOptions_Language_Sv = "sv"
)

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

	// Information about Continuous Relevancy Training for this environment.
	SearchStatus *SearchStatus `json:"search_status,omitempty"`
}

// Constants associated with the Environment.Status property.
// Current status of the environment. `resizing` is displayed when a request to increase the environment size has been
// made, but is still in the process of being completed.
const (
	Environment_Status_Active      = "active"
	Environment_Status_Maintenance = "maintenance"
	Environment_Status_Pending     = "pending"
	Environment_Status_Resizing    = "resizing"
)

// Constants associated with the Environment.Size property.
// Current size of the environment.
const (
	Environment_Size_L    = "L"
	Environment_Size_LT   = "LT"
	Environment_Size_M    = "M"
	Environment_Size_ML   = "ML"
	Environment_Size_MS   = "MS"
	Environment_Size_S    = "S"
	Environment_Size_XL   = "XL"
	Environment_Size_XS   = "XS"
	Environment_Size_XXL  = "XXL"
	Environment_Size_XXXL = "XXXL"
)

// EnvironmentDocuments : Summary of the document usage statistics for the environment.
type EnvironmentDocuments struct {

	// Number of documents indexed for the environment.
	Indexed *int64 `json:"indexed,omitempty"`

	// Total number of documents allowed in the environment's capacity.
	MaximumAllowed *int64 `json:"maximum_allowed,omitempty"`
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

// Expansion : An expansion definition. Each object respresents one set of expandable strings. For example, you could have
// expansions for the word `hot` in one object, and expansions for the word `cold` in another.
type Expansion struct {

	// A list of terms that will be expanded for this expansion. If specified, only the items in this list are expanded.
	InputTerms []string `json:"input_terms,omitempty"`

	// A list of terms that this expansion will be expanded to. If specified without **input_terms**, it also functions as
	// the input term list.
	ExpandedTerms []string `json:"expanded_terms" validate:"required"`
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
	Expansions []Expansion `json:"expansions,omitempty"`
}

// FederatedQueryNoticesOptions : The federatedQueryNotices options.
type FederatedQueryNoticesOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// A comma-separated list of collection IDs to be queried against.
	CollectionIds []string `json:"collection_ids" validate:"required"`

	// A cacheable query that excludes documents that don't mention the query content. Filter searches are better for
	// metadata-type searches and for assessing the concepts in the data set.
	Filter *string `json:"filter,omitempty"`

	// A query search returns all documents in your data set with full enrichments and full text, but with the most
	// relevant documents listed first. Use a query search when you want to find the most relevant search results. You
	// cannot use **natural_language_query** and **query** at the same time.
	Query *string `json:"query,omitempty"`

	// A natural language query that returns relevant documents by utilizing training data and natural language
	// understanding. You cannot use **natural_language_query** and **query** at the same time.
	NaturalLanguageQuery *string `json:"natural_language_query,omitempty"`

	// An aggregation search that returns an exact answer by combining query search with filters. Useful for applications
	// to build lists, tables, and time series. For a full list of possible aggregations, see the Query reference.
	Aggregation *string `json:"aggregation,omitempty"`

	// Number of results to return.
	Count *int64 `json:"count,omitempty"`

	// A comma-separated list of the portion of the document hierarchy to return.
	ReturnFields []string `json:"return,omitempty"`

	// The number of query results to skip at the beginning. For example, if the total number of results that are returned
	// is 10 and the offset is 8, it returns the last two results.
	Offset *int64 `json:"offset,omitempty"`

	// A comma-separated list of fields in the document to sort on. You can optionally specify a sort direction by
	// prefixing the field with `-` for descending or `+` for ascending. Ascending is the default sort direction if no
	// prefix is specified.
	Sort []string `json:"sort,omitempty"`

	// When true, a highlight field is returned for each result which contains the fields which match the query with
	// `<em></em>` tags around the matching query terms.
	Highlight *bool `json:"highlight,omitempty"`

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
	SimilarDocumentIds []string `json:"similar.document_ids,omitempty"`

	// A comma-separated list of field names that are used as a basis for comparison to identify similar documents. If not
	// specified, the entire document is used for comparison.
	SimilarFields []string `json:"similar.fields,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewFederatedQueryNoticesOptions : Instantiate FederatedQueryNoticesOptions
func (discovery *DiscoveryV1) NewFederatedQueryNoticesOptions(environmentID string, collectionIds []string) *FederatedQueryNoticesOptions {
	return &FederatedQueryNoticesOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionIds: collectionIds,
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *FederatedQueryNoticesOptions) SetEnvironmentID(environmentID string) *FederatedQueryNoticesOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetCollectionIds : Allow user to set CollectionIds
func (options *FederatedQueryNoticesOptions) SetCollectionIds(collectionIds []string) *FederatedQueryNoticesOptions {
	options.CollectionIds = collectionIds
	return options
}

// SetFilter : Allow user to set Filter
func (options *FederatedQueryNoticesOptions) SetFilter(filter string) *FederatedQueryNoticesOptions {
	options.Filter = core.StringPtr(filter)
	return options
}

// SetQuery : Allow user to set Query
func (options *FederatedQueryNoticesOptions) SetQuery(query string) *FederatedQueryNoticesOptions {
	options.Query = core.StringPtr(query)
	return options
}

// SetNaturalLanguageQuery : Allow user to set NaturalLanguageQuery
func (options *FederatedQueryNoticesOptions) SetNaturalLanguageQuery(naturalLanguageQuery string) *FederatedQueryNoticesOptions {
	options.NaturalLanguageQuery = core.StringPtr(naturalLanguageQuery)
	return options
}

// SetAggregation : Allow user to set Aggregation
func (options *FederatedQueryNoticesOptions) SetAggregation(aggregation string) *FederatedQueryNoticesOptions {
	options.Aggregation = core.StringPtr(aggregation)
	return options
}

// SetCount : Allow user to set Count
func (options *FederatedQueryNoticesOptions) SetCount(count int64) *FederatedQueryNoticesOptions {
	options.Count = core.Int64Ptr(count)
	return options
}

// SetReturnFields : Allow user to set ReturnFields
func (options *FederatedQueryNoticesOptions) SetReturnFields(returnFields []string) *FederatedQueryNoticesOptions {
	options.ReturnFields = returnFields
	return options
}

// SetOffset : Allow user to set Offset
func (options *FederatedQueryNoticesOptions) SetOffset(offset int64) *FederatedQueryNoticesOptions {
	options.Offset = core.Int64Ptr(offset)
	return options
}

// SetSort : Allow user to set Sort
func (options *FederatedQueryNoticesOptions) SetSort(sort []string) *FederatedQueryNoticesOptions {
	options.Sort = sort
	return options
}

// SetHighlight : Allow user to set Highlight
func (options *FederatedQueryNoticesOptions) SetHighlight(highlight bool) *FederatedQueryNoticesOptions {
	options.Highlight = core.BoolPtr(highlight)
	return options
}

// SetDeduplicateField : Allow user to set DeduplicateField
func (options *FederatedQueryNoticesOptions) SetDeduplicateField(deduplicateField string) *FederatedQueryNoticesOptions {
	options.DeduplicateField = core.StringPtr(deduplicateField)
	return options
}

// SetSimilar : Allow user to set Similar
func (options *FederatedQueryNoticesOptions) SetSimilar(similar bool) *FederatedQueryNoticesOptions {
	options.Similar = core.BoolPtr(similar)
	return options
}

// SetSimilarDocumentIds : Allow user to set SimilarDocumentIds
func (options *FederatedQueryNoticesOptions) SetSimilarDocumentIds(similarDocumentIds []string) *FederatedQueryNoticesOptions {
	options.SimilarDocumentIds = similarDocumentIds
	return options
}

// SetSimilarFields : Allow user to set SimilarFields
func (options *FederatedQueryNoticesOptions) SetSimilarFields(similarFields []string) *FederatedQueryNoticesOptions {
	options.SimilarFields = similarFields
	return options
}

// SetHeaders : Allow user to set Headers
func (options *FederatedQueryNoticesOptions) SetHeaders(param map[string]string) *FederatedQueryNoticesOptions {
	options.Headers = param
	return options
}

// FederatedQueryOptions : The federatedQuery options.
type FederatedQueryOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// A cacheable query that excludes documents that don't mention the query content. Filter searches are better for
	// metadata-type searches and for assessing the concepts in the data set.
	Filter *string `json:"filter,omitempty"`

	// A query search returns all documents in your data set with full enrichments and full text, but with the most
	// relevant documents listed first. Use a query search when you want to find the most relevant search results. You
	// cannot use **natural_language_query** and **query** at the same time.
	Query *string `json:"query,omitempty"`

	// A natural language query that returns relevant documents by utilizing training data and natural language
	// understanding. You cannot use **natural_language_query** and **query** at the same time.
	NaturalLanguageQuery *string `json:"natural_language_query,omitempty"`

	// A passages query that returns the most relevant passages from the results.
	Passages *bool `json:"passages,omitempty"`

	// An aggregation search that returns an exact answer by combining query search with filters. Useful for applications
	// to build lists, tables, and time series. For a full list of possible aggregations, see the Query reference.
	Aggregation *string `json:"aggregation,omitempty"`

	// Number of results to return.
	Count *int64 `json:"count,omitempty"`

	// A comma-separated list of the portion of the document hierarchy to return.
	ReturnFields *string `json:"return,omitempty"`

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

	// When `true` and used with a Watson Discovery News collection, duplicate results (based on the contents of the
	// **title** field) are removed. Duplicate comparison is limited to the current query only; **offset** is not
	// considered. This parameter is currently Beta functionality.
	Deduplicate *bool `json:"deduplicate,omitempty"`

	// When specified, duplicate results based on the field specified are removed from the returned results. Duplicate
	// comparison is limited to the current query only, **offset** is not considered. This parameter is currently Beta
	// functionality.
	DeduplicateField *string `json:"deduplicate.field,omitempty"`

	// A comma-separated list of collection IDs to be queried against. Required when querying multiple collections, invalid
	// when performing a single collection query.
	CollectionIds *string `json:"collection_ids,omitempty"`

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
	LoggingOptOut *bool `json:"X-Watson-Logging-Opt-Out,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewFederatedQueryOptions : Instantiate FederatedQueryOptions
func (discovery *DiscoveryV1) NewFederatedQueryOptions(environmentID string) *FederatedQueryOptions {
	return &FederatedQueryOptions{
		EnvironmentID: core.StringPtr(environmentID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *FederatedQueryOptions) SetEnvironmentID(environmentID string) *FederatedQueryOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetFilter : Allow user to set Filter
func (options *FederatedQueryOptions) SetFilter(filter string) *FederatedQueryOptions {
	options.Filter = core.StringPtr(filter)
	return options
}

// SetQuery : Allow user to set Query
func (options *FederatedQueryOptions) SetQuery(query string) *FederatedQueryOptions {
	options.Query = core.StringPtr(query)
	return options
}

// SetNaturalLanguageQuery : Allow user to set NaturalLanguageQuery
func (options *FederatedQueryOptions) SetNaturalLanguageQuery(naturalLanguageQuery string) *FederatedQueryOptions {
	options.NaturalLanguageQuery = core.StringPtr(naturalLanguageQuery)
	return options
}

// SetPassages : Allow user to set Passages
func (options *FederatedQueryOptions) SetPassages(passages bool) *FederatedQueryOptions {
	options.Passages = core.BoolPtr(passages)
	return options
}

// SetAggregation : Allow user to set Aggregation
func (options *FederatedQueryOptions) SetAggregation(aggregation string) *FederatedQueryOptions {
	options.Aggregation = core.StringPtr(aggregation)
	return options
}

// SetCount : Allow user to set Count
func (options *FederatedQueryOptions) SetCount(count int64) *FederatedQueryOptions {
	options.Count = core.Int64Ptr(count)
	return options
}

// SetReturnFields : Allow user to set ReturnFields
func (options *FederatedQueryOptions) SetReturnFields(returnFields string) *FederatedQueryOptions {
	options.ReturnFields = core.StringPtr(returnFields)
	return options
}

// SetOffset : Allow user to set Offset
func (options *FederatedQueryOptions) SetOffset(offset int64) *FederatedQueryOptions {
	options.Offset = core.Int64Ptr(offset)
	return options
}

// SetSort : Allow user to set Sort
func (options *FederatedQueryOptions) SetSort(sort string) *FederatedQueryOptions {
	options.Sort = core.StringPtr(sort)
	return options
}

// SetHighlight : Allow user to set Highlight
func (options *FederatedQueryOptions) SetHighlight(highlight bool) *FederatedQueryOptions {
	options.Highlight = core.BoolPtr(highlight)
	return options
}

// SetPassagesFields : Allow user to set PassagesFields
func (options *FederatedQueryOptions) SetPassagesFields(passagesFields string) *FederatedQueryOptions {
	options.PassagesFields = core.StringPtr(passagesFields)
	return options
}

// SetPassagesCount : Allow user to set PassagesCount
func (options *FederatedQueryOptions) SetPassagesCount(passagesCount int64) *FederatedQueryOptions {
	options.PassagesCount = core.Int64Ptr(passagesCount)
	return options
}

// SetPassagesCharacters : Allow user to set PassagesCharacters
func (options *FederatedQueryOptions) SetPassagesCharacters(passagesCharacters int64) *FederatedQueryOptions {
	options.PassagesCharacters = core.Int64Ptr(passagesCharacters)
	return options
}

// SetDeduplicate : Allow user to set Deduplicate
func (options *FederatedQueryOptions) SetDeduplicate(deduplicate bool) *FederatedQueryOptions {
	options.Deduplicate = core.BoolPtr(deduplicate)
	return options
}

// SetDeduplicateField : Allow user to set DeduplicateField
func (options *FederatedQueryOptions) SetDeduplicateField(deduplicateField string) *FederatedQueryOptions {
	options.DeduplicateField = core.StringPtr(deduplicateField)
	return options
}

// SetCollectionIds : Allow user to set CollectionIds
func (options *FederatedQueryOptions) SetCollectionIds(collectionIds string) *FederatedQueryOptions {
	options.CollectionIds = core.StringPtr(collectionIds)
	return options
}

// SetSimilar : Allow user to set Similar
func (options *FederatedQueryOptions) SetSimilar(similar bool) *FederatedQueryOptions {
	options.Similar = core.BoolPtr(similar)
	return options
}

// SetSimilarDocumentIds : Allow user to set SimilarDocumentIds
func (options *FederatedQueryOptions) SetSimilarDocumentIds(similarDocumentIds string) *FederatedQueryOptions {
	options.SimilarDocumentIds = core.StringPtr(similarDocumentIds)
	return options
}

// SetSimilarFields : Allow user to set SimilarFields
func (options *FederatedQueryOptions) SetSimilarFields(similarFields string) *FederatedQueryOptions {
	options.SimilarFields = core.StringPtr(similarFields)
	return options
}

// SetBias : Allow user to set Bias
func (options *FederatedQueryOptions) SetBias(bias string) *FederatedQueryOptions {
	options.Bias = core.StringPtr(bias)
	return options
}

// SetLoggingOptOut : Allow user to set LoggingOptOut
func (options *FederatedQueryOptions) SetLoggingOptOut(loggingOptOut bool) *FederatedQueryOptions {
	options.LoggingOptOut = core.BoolPtr(loggingOptOut)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *FederatedQueryOptions) SetHeaders(param map[string]string) *FederatedQueryOptions {
	options.Headers = param
	return options
}

// Field : Field struct
type Field struct {

	// The name of the field.
	FieldName *string `json:"field,omitempty"`

	// The type of the field.
	FieldType *string `json:"type,omitempty"`
}

// Constants associated with the Field.FieldType property.
// The type of the field.
const (
	Field_FieldType_Binary  = "binary"
	Field_FieldType_Boolean = "boolean"
	Field_FieldType_Byte    = "byte"
	Field_FieldType_Date    = "date"
	Field_FieldType_Double  = "double"
	Field_FieldType_Float   = "float"
	Field_FieldType_Integer = "integer"
	Field_FieldType_Long    = "long"
	Field_FieldType_Nested  = "nested"
	Field_FieldType_Short   = "short"
	Field_FieldType_String  = "string"
)

// FontSetting : FontSetting struct
type FontSetting struct {
	Level *int64 `json:"level,omitempty"`

	MinSize *int64 `json:"min_size,omitempty"`

	MaxSize *int64 `json:"max_size,omitempty"`

	Bold *bool `json:"bold,omitempty"`

	Italic *bool `json:"italic,omitempty"`

	Name *string `json:"name,omitempty"`
}

// GetCollectionOptions : The getCollection options.
type GetCollectionOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The ID of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetCollectionOptions : Instantiate GetCollectionOptions
func (discovery *DiscoveryV1) NewGetCollectionOptions(environmentID string, collectionID string) *GetCollectionOptions {
	return &GetCollectionOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *GetCollectionOptions) SetEnvironmentID(environmentID string) *GetCollectionOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
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

// GetConfigurationOptions : The getConfiguration options.
type GetConfigurationOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The ID of the configuration.
	ConfigurationID *string `json:"configuration_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetConfigurationOptions : Instantiate GetConfigurationOptions
func (discovery *DiscoveryV1) NewGetConfigurationOptions(environmentID string, configurationID string) *GetConfigurationOptions {
	return &GetConfigurationOptions{
		EnvironmentID:   core.StringPtr(environmentID),
		ConfigurationID: core.StringPtr(configurationID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *GetConfigurationOptions) SetEnvironmentID(environmentID string) *GetConfigurationOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetConfigurationID : Allow user to set ConfigurationID
func (options *GetConfigurationOptions) SetConfigurationID(configurationID string) *GetConfigurationOptions {
	options.ConfigurationID = core.StringPtr(configurationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetConfigurationOptions) SetHeaders(param map[string]string) *GetConfigurationOptions {
	options.Headers = param
	return options
}

// GetCredentialsOptions : The getCredentials options.
type GetCredentialsOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The unique identifier for a set of source credentials.
	CredentialID *string `json:"credential_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetCredentialsOptions : Instantiate GetCredentialsOptions
func (discovery *DiscoveryV1) NewGetCredentialsOptions(environmentID string, credentialID string) *GetCredentialsOptions {
	return &GetCredentialsOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CredentialID:  core.StringPtr(credentialID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *GetCredentialsOptions) SetEnvironmentID(environmentID string) *GetCredentialsOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetCredentialID : Allow user to set CredentialID
func (options *GetCredentialsOptions) SetCredentialID(credentialID string) *GetCredentialsOptions {
	options.CredentialID = core.StringPtr(credentialID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetCredentialsOptions) SetHeaders(param map[string]string) *GetCredentialsOptions {
	options.Headers = param
	return options
}

// GetDocumentStatusOptions : The getDocumentStatus options.
type GetDocumentStatusOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The ID of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// The ID of the document.
	DocumentID *string `json:"document_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetDocumentStatusOptions : Instantiate GetDocumentStatusOptions
func (discovery *DiscoveryV1) NewGetDocumentStatusOptions(environmentID string, collectionID string, documentID string) *GetDocumentStatusOptions {
	return &GetDocumentStatusOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
		DocumentID:    core.StringPtr(documentID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *GetDocumentStatusOptions) SetEnvironmentID(environmentID string) *GetDocumentStatusOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *GetDocumentStatusOptions) SetCollectionID(collectionID string) *GetDocumentStatusOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
}

// SetDocumentID : Allow user to set DocumentID
func (options *GetDocumentStatusOptions) SetDocumentID(documentID string) *GetDocumentStatusOptions {
	options.DocumentID = core.StringPtr(documentID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetDocumentStatusOptions) SetHeaders(param map[string]string) *GetDocumentStatusOptions {
	options.Headers = param
	return options
}

// GetEnvironmentOptions : The getEnvironment options.
type GetEnvironmentOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetEnvironmentOptions : Instantiate GetEnvironmentOptions
func (discovery *DiscoveryV1) NewGetEnvironmentOptions(environmentID string) *GetEnvironmentOptions {
	return &GetEnvironmentOptions{
		EnvironmentID: core.StringPtr(environmentID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *GetEnvironmentOptions) SetEnvironmentID(environmentID string) *GetEnvironmentOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetEnvironmentOptions) SetHeaders(param map[string]string) *GetEnvironmentOptions {
	options.Headers = param
	return options
}

// GetMetricsEventRateOptions : The getMetricsEventRate options.
type GetMetricsEventRateOptions struct {

	// Metric is computed from data recorded after this timestamp; must be in `YYYY-MM-DDThh:mm:ssZ` format.
	StartTime *strfmt.DateTime `json:"start_time,omitempty"`

	// Metric is computed from data recorded before this timestamp; must be in `YYYY-MM-DDThh:mm:ssZ` format.
	EndTime *strfmt.DateTime `json:"end_time,omitempty"`

	// The type of result to consider when calculating the metric.
	ResultType *string `json:"result_type,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the GetMetricsEventRateOptions.ResultType property.
// The type of result to consider when calculating the metric.
const (
	GetMetricsEventRateOptions_ResultType_Document = "document"
)

// NewGetMetricsEventRateOptions : Instantiate GetMetricsEventRateOptions
func (discovery *DiscoveryV1) NewGetMetricsEventRateOptions() *GetMetricsEventRateOptions {
	return &GetMetricsEventRateOptions{}
}

// SetStartTime : Allow user to set StartTime
func (options *GetMetricsEventRateOptions) SetStartTime(startTime *strfmt.DateTime) *GetMetricsEventRateOptions {
	options.StartTime = startTime
	return options
}

// SetEndTime : Allow user to set EndTime
func (options *GetMetricsEventRateOptions) SetEndTime(endTime *strfmt.DateTime) *GetMetricsEventRateOptions {
	options.EndTime = endTime
	return options
}

// SetResultType : Allow user to set ResultType
func (options *GetMetricsEventRateOptions) SetResultType(resultType string) *GetMetricsEventRateOptions {
	options.ResultType = core.StringPtr(resultType)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetMetricsEventRateOptions) SetHeaders(param map[string]string) *GetMetricsEventRateOptions {
	options.Headers = param
	return options
}

// GetMetricsQueryEventOptions : The getMetricsQueryEvent options.
type GetMetricsQueryEventOptions struct {

	// Metric is computed from data recorded after this timestamp; must be in `YYYY-MM-DDThh:mm:ssZ` format.
	StartTime *strfmt.DateTime `json:"start_time,omitempty"`

	// Metric is computed from data recorded before this timestamp; must be in `YYYY-MM-DDThh:mm:ssZ` format.
	EndTime *strfmt.DateTime `json:"end_time,omitempty"`

	// The type of result to consider when calculating the metric.
	ResultType *string `json:"result_type,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the GetMetricsQueryEventOptions.ResultType property.
// The type of result to consider when calculating the metric.
const (
	GetMetricsQueryEventOptions_ResultType_Document = "document"
)

// NewGetMetricsQueryEventOptions : Instantiate GetMetricsQueryEventOptions
func (discovery *DiscoveryV1) NewGetMetricsQueryEventOptions() *GetMetricsQueryEventOptions {
	return &GetMetricsQueryEventOptions{}
}

// SetStartTime : Allow user to set StartTime
func (options *GetMetricsQueryEventOptions) SetStartTime(startTime *strfmt.DateTime) *GetMetricsQueryEventOptions {
	options.StartTime = startTime
	return options
}

// SetEndTime : Allow user to set EndTime
func (options *GetMetricsQueryEventOptions) SetEndTime(endTime *strfmt.DateTime) *GetMetricsQueryEventOptions {
	options.EndTime = endTime
	return options
}

// SetResultType : Allow user to set ResultType
func (options *GetMetricsQueryEventOptions) SetResultType(resultType string) *GetMetricsQueryEventOptions {
	options.ResultType = core.StringPtr(resultType)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetMetricsQueryEventOptions) SetHeaders(param map[string]string) *GetMetricsQueryEventOptions {
	options.Headers = param
	return options
}

// GetMetricsQueryNoResultsOptions : The getMetricsQueryNoResults options.
type GetMetricsQueryNoResultsOptions struct {

	// Metric is computed from data recorded after this timestamp; must be in `YYYY-MM-DDThh:mm:ssZ` format.
	StartTime *strfmt.DateTime `json:"start_time,omitempty"`

	// Metric is computed from data recorded before this timestamp; must be in `YYYY-MM-DDThh:mm:ssZ` format.
	EndTime *strfmt.DateTime `json:"end_time,omitempty"`

	// The type of result to consider when calculating the metric.
	ResultType *string `json:"result_type,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the GetMetricsQueryNoResultsOptions.ResultType property.
// The type of result to consider when calculating the metric.
const (
	GetMetricsQueryNoResultsOptions_ResultType_Document = "document"
)

// NewGetMetricsQueryNoResultsOptions : Instantiate GetMetricsQueryNoResultsOptions
func (discovery *DiscoveryV1) NewGetMetricsQueryNoResultsOptions() *GetMetricsQueryNoResultsOptions {
	return &GetMetricsQueryNoResultsOptions{}
}

// SetStartTime : Allow user to set StartTime
func (options *GetMetricsQueryNoResultsOptions) SetStartTime(startTime *strfmt.DateTime) *GetMetricsQueryNoResultsOptions {
	options.StartTime = startTime
	return options
}

// SetEndTime : Allow user to set EndTime
func (options *GetMetricsQueryNoResultsOptions) SetEndTime(endTime *strfmt.DateTime) *GetMetricsQueryNoResultsOptions {
	options.EndTime = endTime
	return options
}

// SetResultType : Allow user to set ResultType
func (options *GetMetricsQueryNoResultsOptions) SetResultType(resultType string) *GetMetricsQueryNoResultsOptions {
	options.ResultType = core.StringPtr(resultType)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetMetricsQueryNoResultsOptions) SetHeaders(param map[string]string) *GetMetricsQueryNoResultsOptions {
	options.Headers = param
	return options
}

// GetMetricsQueryOptions : The getMetricsQuery options.
type GetMetricsQueryOptions struct {

	// Metric is computed from data recorded after this timestamp; must be in `YYYY-MM-DDThh:mm:ssZ` format.
	StartTime *strfmt.DateTime `json:"start_time,omitempty"`

	// Metric is computed from data recorded before this timestamp; must be in `YYYY-MM-DDThh:mm:ssZ` format.
	EndTime *strfmt.DateTime `json:"end_time,omitempty"`

	// The type of result to consider when calculating the metric.
	ResultType *string `json:"result_type,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the GetMetricsQueryOptions.ResultType property.
// The type of result to consider when calculating the metric.
const (
	GetMetricsQueryOptions_ResultType_Document = "document"
)

// NewGetMetricsQueryOptions : Instantiate GetMetricsQueryOptions
func (discovery *DiscoveryV1) NewGetMetricsQueryOptions() *GetMetricsQueryOptions {
	return &GetMetricsQueryOptions{}
}

// SetStartTime : Allow user to set StartTime
func (options *GetMetricsQueryOptions) SetStartTime(startTime *strfmt.DateTime) *GetMetricsQueryOptions {
	options.StartTime = startTime
	return options
}

// SetEndTime : Allow user to set EndTime
func (options *GetMetricsQueryOptions) SetEndTime(endTime *strfmt.DateTime) *GetMetricsQueryOptions {
	options.EndTime = endTime
	return options
}

// SetResultType : Allow user to set ResultType
func (options *GetMetricsQueryOptions) SetResultType(resultType string) *GetMetricsQueryOptions {
	options.ResultType = core.StringPtr(resultType)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetMetricsQueryOptions) SetHeaders(param map[string]string) *GetMetricsQueryOptions {
	options.Headers = param
	return options
}

// GetMetricsQueryTokenEventOptions : The getMetricsQueryTokenEvent options.
type GetMetricsQueryTokenEventOptions struct {

	// Number of results to return.
	Count *int64 `json:"count,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetMetricsQueryTokenEventOptions : Instantiate GetMetricsQueryTokenEventOptions
func (discovery *DiscoveryV1) NewGetMetricsQueryTokenEventOptions() *GetMetricsQueryTokenEventOptions {
	return &GetMetricsQueryTokenEventOptions{}
}

// SetCount : Allow user to set Count
func (options *GetMetricsQueryTokenEventOptions) SetCount(count int64) *GetMetricsQueryTokenEventOptions {
	options.Count = core.Int64Ptr(count)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetMetricsQueryTokenEventOptions) SetHeaders(param map[string]string) *GetMetricsQueryTokenEventOptions {
	options.Headers = param
	return options
}

// GetTokenizationDictionaryStatusOptions : The getTokenizationDictionaryStatus options.
type GetTokenizationDictionaryStatusOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The ID of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetTokenizationDictionaryStatusOptions : Instantiate GetTokenizationDictionaryStatusOptions
func (discovery *DiscoveryV1) NewGetTokenizationDictionaryStatusOptions(environmentID string, collectionID string) *GetTokenizationDictionaryStatusOptions {
	return &GetTokenizationDictionaryStatusOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *GetTokenizationDictionaryStatusOptions) SetEnvironmentID(environmentID string) *GetTokenizationDictionaryStatusOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *GetTokenizationDictionaryStatusOptions) SetCollectionID(collectionID string) *GetTokenizationDictionaryStatusOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetTokenizationDictionaryStatusOptions) SetHeaders(param map[string]string) *GetTokenizationDictionaryStatusOptions {
	options.Headers = param
	return options
}

// GetTrainingDataOptions : The getTrainingData options.
type GetTrainingDataOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The ID of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// The ID of the query used for training.
	QueryID *string `json:"query_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetTrainingDataOptions : Instantiate GetTrainingDataOptions
func (discovery *DiscoveryV1) NewGetTrainingDataOptions(environmentID string, collectionID string, queryID string) *GetTrainingDataOptions {
	return &GetTrainingDataOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
		QueryID:       core.StringPtr(queryID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *GetTrainingDataOptions) SetEnvironmentID(environmentID string) *GetTrainingDataOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *GetTrainingDataOptions) SetCollectionID(collectionID string) *GetTrainingDataOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
}

// SetQueryID : Allow user to set QueryID
func (options *GetTrainingDataOptions) SetQueryID(queryID string) *GetTrainingDataOptions {
	options.QueryID = core.StringPtr(queryID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetTrainingDataOptions) SetHeaders(param map[string]string) *GetTrainingDataOptions {
	options.Headers = param
	return options
}

// GetTrainingExampleOptions : The getTrainingExample options.
type GetTrainingExampleOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The ID of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// The ID of the query used for training.
	QueryID *string `json:"query_id" validate:"required"`

	// The ID of the document as it is indexed.
	ExampleID *string `json:"example_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetTrainingExampleOptions : Instantiate GetTrainingExampleOptions
func (discovery *DiscoveryV1) NewGetTrainingExampleOptions(environmentID string, collectionID string, queryID string, exampleID string) *GetTrainingExampleOptions {
	return &GetTrainingExampleOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
		QueryID:       core.StringPtr(queryID),
		ExampleID:     core.StringPtr(exampleID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *GetTrainingExampleOptions) SetEnvironmentID(environmentID string) *GetTrainingExampleOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *GetTrainingExampleOptions) SetCollectionID(collectionID string) *GetTrainingExampleOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
}

// SetQueryID : Allow user to set QueryID
func (options *GetTrainingExampleOptions) SetQueryID(queryID string) *GetTrainingExampleOptions {
	options.QueryID = core.StringPtr(queryID)
	return options
}

// SetExampleID : Allow user to set ExampleID
func (options *GetTrainingExampleOptions) SetExampleID(exampleID string) *GetTrainingExampleOptions {
	options.ExampleID = core.StringPtr(exampleID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetTrainingExampleOptions) SetHeaders(param map[string]string) *GetTrainingExampleOptions {
	options.Headers = param
	return options
}

// HTMLSettings : A list of HTML conversion settings.
type HTMLSettings struct {
	ExcludeTagsCompletely []string `json:"exclude_tags_completely,omitempty"`

	ExcludeTagsKeepContent []string `json:"exclude_tags_keep_content,omitempty"`

	KeepContent *XPathPatterns `json:"keep_content,omitempty"`

	ExcludeContent *XPathPatterns `json:"exclude_content,omitempty"`

	KeepTagAttributes []string `json:"keep_tag_attributes,omitempty"`

	ExcludeTagAttributes []string `json:"exclude_tag_attributes,omitempty"`
}

// IndexCapacity : Details about the resource usage and capacity of the environment.
type IndexCapacity struct {

	// Summary of the document usage statistics for the environment.
	Documents *EnvironmentDocuments `json:"documents,omitempty"`

	// Summary of the disk usage of the environment.
	DiskUsage *DiskUsage `json:"disk_usage,omitempty"`

	// Summary of the collection usage in the environment.
	Collections *CollectionUsage `json:"collections,omitempty"`

	// **Deprecated**: Summary of the memory usage of the environment.
	MemoryUsage *MemoryUsage `json:"memory_usage,omitempty"`
}

// ListCollectionFieldsOptions : The listCollectionFields options.
type ListCollectionFieldsOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The ID of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListCollectionFieldsOptions : Instantiate ListCollectionFieldsOptions
func (discovery *DiscoveryV1) NewListCollectionFieldsOptions(environmentID string, collectionID string) *ListCollectionFieldsOptions {
	return &ListCollectionFieldsOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *ListCollectionFieldsOptions) SetEnvironmentID(environmentID string) *ListCollectionFieldsOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *ListCollectionFieldsOptions) SetCollectionID(collectionID string) *ListCollectionFieldsOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
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

// ListCollectionsOptions : The listCollections options.
type ListCollectionsOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// Find collections with the given name.
	Name *string `json:"name,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListCollectionsOptions : Instantiate ListCollectionsOptions
func (discovery *DiscoveryV1) NewListCollectionsOptions(environmentID string) *ListCollectionsOptions {
	return &ListCollectionsOptions{
		EnvironmentID: core.StringPtr(environmentID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *ListCollectionsOptions) SetEnvironmentID(environmentID string) *ListCollectionsOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetName : Allow user to set Name
func (options *ListCollectionsOptions) SetName(name string) *ListCollectionsOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListCollectionsOptions) SetHeaders(param map[string]string) *ListCollectionsOptions {
	options.Headers = param
	return options
}

// ListCollectionsResponse : ListCollectionsResponse struct
type ListCollectionsResponse struct {

	// An array containing information about each collection in the environment.
	Collections []Collection `json:"collections,omitempty"`
}

// ListConfigurationsOptions : The listConfigurations options.
type ListConfigurationsOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// Find configurations with the given name.
	Name *string `json:"name,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListConfigurationsOptions : Instantiate ListConfigurationsOptions
func (discovery *DiscoveryV1) NewListConfigurationsOptions(environmentID string) *ListConfigurationsOptions {
	return &ListConfigurationsOptions{
		EnvironmentID: core.StringPtr(environmentID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *ListConfigurationsOptions) SetEnvironmentID(environmentID string) *ListConfigurationsOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetName : Allow user to set Name
func (options *ListConfigurationsOptions) SetName(name string) *ListConfigurationsOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListConfigurationsOptions) SetHeaders(param map[string]string) *ListConfigurationsOptions {
	options.Headers = param
	return options
}

// ListConfigurationsResponse : ListConfigurationsResponse struct
type ListConfigurationsResponse struct {

	// An array of Configurations that are available for the service instance.
	Configurations []Configuration `json:"configurations,omitempty"`
}

// ListCredentialsOptions : The listCredentials options.
type ListCredentialsOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListCredentialsOptions : Instantiate ListCredentialsOptions
func (discovery *DiscoveryV1) NewListCredentialsOptions(environmentID string) *ListCredentialsOptions {
	return &ListCredentialsOptions{
		EnvironmentID: core.StringPtr(environmentID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *ListCredentialsOptions) SetEnvironmentID(environmentID string) *ListCredentialsOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListCredentialsOptions) SetHeaders(param map[string]string) *ListCredentialsOptions {
	options.Headers = param
	return options
}

// ListEnvironmentsOptions : The listEnvironments options.
type ListEnvironmentsOptions struct {

	// Show only the environment with the given name.
	Name *string `json:"name,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListEnvironmentsOptions : Instantiate ListEnvironmentsOptions
func (discovery *DiscoveryV1) NewListEnvironmentsOptions() *ListEnvironmentsOptions {
	return &ListEnvironmentsOptions{}
}

// SetName : Allow user to set Name
func (options *ListEnvironmentsOptions) SetName(name string) *ListEnvironmentsOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListEnvironmentsOptions) SetHeaders(param map[string]string) *ListEnvironmentsOptions {
	options.Headers = param
	return options
}

// ListEnvironmentsResponse : ListEnvironmentsResponse struct
type ListEnvironmentsResponse struct {

	// An array of [environments] that are available for the service instance.
	Environments []Environment `json:"environments,omitempty"`
}

// ListExpansionsOptions : The listExpansions options.
type ListExpansionsOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The ID of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListExpansionsOptions : Instantiate ListExpansionsOptions
func (discovery *DiscoveryV1) NewListExpansionsOptions(environmentID string, collectionID string) *ListExpansionsOptions {
	return &ListExpansionsOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *ListExpansionsOptions) SetEnvironmentID(environmentID string) *ListExpansionsOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *ListExpansionsOptions) SetCollectionID(collectionID string) *ListExpansionsOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListExpansionsOptions) SetHeaders(param map[string]string) *ListExpansionsOptions {
	options.Headers = param
	return options
}

// ListFieldsOptions : The listFields options.
type ListFieldsOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// A comma-separated list of collection IDs to be queried against.
	CollectionIds []string `json:"collection_ids" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListFieldsOptions : Instantiate ListFieldsOptions
func (discovery *DiscoveryV1) NewListFieldsOptions(environmentID string, collectionIds []string) *ListFieldsOptions {
	return &ListFieldsOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionIds: collectionIds,
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *ListFieldsOptions) SetEnvironmentID(environmentID string) *ListFieldsOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
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

// ListTrainingDataOptions : The listTrainingData options.
type ListTrainingDataOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The ID of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListTrainingDataOptions : Instantiate ListTrainingDataOptions
func (discovery *DiscoveryV1) NewListTrainingDataOptions(environmentID string, collectionID string) *ListTrainingDataOptions {
	return &ListTrainingDataOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *ListTrainingDataOptions) SetEnvironmentID(environmentID string) *ListTrainingDataOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *ListTrainingDataOptions) SetCollectionID(collectionID string) *ListTrainingDataOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListTrainingDataOptions) SetHeaders(param map[string]string) *ListTrainingDataOptions {
	options.Headers = param
	return options
}

// ListTrainingExamplesOptions : The listTrainingExamples options.
type ListTrainingExamplesOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The ID of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// The ID of the query used for training.
	QueryID *string `json:"query_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListTrainingExamplesOptions : Instantiate ListTrainingExamplesOptions
func (discovery *DiscoveryV1) NewListTrainingExamplesOptions(environmentID string, collectionID string, queryID string) *ListTrainingExamplesOptions {
	return &ListTrainingExamplesOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
		QueryID:       core.StringPtr(queryID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *ListTrainingExamplesOptions) SetEnvironmentID(environmentID string) *ListTrainingExamplesOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *ListTrainingExamplesOptions) SetCollectionID(collectionID string) *ListTrainingExamplesOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
}

// SetQueryID : Allow user to set QueryID
func (options *ListTrainingExamplesOptions) SetQueryID(queryID string) *ListTrainingExamplesOptions {
	options.QueryID = core.StringPtr(queryID)
	return options
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

	Results []LogQueryResponseResult `json:"results,omitempty"`
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
	LogQueryResponseResult_DocumentType_Event = "event"
	LogQueryResponseResult_DocumentType_Query = "query"
)

// Constants associated with the LogQueryResponseResult.EventType property.
// The type of event that this object respresents. Possible values are
//
//  -  `query` the log of a query to a collection
//
//  -  `click` the result of a call to the **events** endpoint.
const (
	LogQueryResponseResult_EventType_Click = "click"
	LogQueryResponseResult_EventType_Query = "query"
)

// Constants associated with the LogQueryResponseResult.ResultType property.
// The type of result that this **event** is associated with. Only returned with logs of type `event`.
const (
	LogQueryResponseResult_ResultType_Document = "document"
)

// LogQueryResponseResultDocuments : Object containing result information that was returned by the query used to create this log entry. Only returned with
// logs of type `query`.
type LogQueryResponseResultDocuments struct {
	Results []LogQueryResponseResultDocumentsResult `json:"results,omitempty"`

	// The number of results returned in the query associate with this log.
	Count *int64 `json:"count,omitempty"`
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

// MemoryUsage : **Deprecated**: Summary of the memory usage statistics for this environment.
type MemoryUsage struct {

	// **Deprecated**: Number of bytes used in the environment's memory capacity.
	UsedBytes *int64 `json:"used_bytes,omitempty"`

	// **Deprecated**: Total number of bytes available in the environment's memory capacity.
	TotalBytes *int64 `json:"total_bytes,omitempty"`

	// **Deprecated**: Amount of memory capacity used, in KB or GB format.
	Used *string `json:"used,omitempty"`

	// **Deprecated**: Total amount of the environment's memory capacity, in KB or GB format.
	Total *string `json:"total,omitempty"`

	// **Deprecated**: Percentage of the environment's memory capacity that is being used.
	PercentUsed *float64 `json:"percent_used,omitempty"`
}

// MetricAggregation : An aggregation analyzing log information for queries and events.
type MetricAggregation struct {

	// The measurement interval for this metric. Metric intervals are always 1 day (`1d`).
	Interval *string `json:"interval,omitempty"`

	// The event type associated with this metric result. This field, when present, will always be `click`.
	EventType *string `json:"event_type,omitempty"`

	Results []MetricAggregationResult `json:"results,omitempty"`
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

// MetricResponse : The response generated from a call to a **metrics** method.
type MetricResponse struct {
	Aggregations []MetricAggregation `json:"aggregations,omitempty"`
}

// MetricTokenAggregation : An aggregation analyzing log information for queries and events.
type MetricTokenAggregation struct {

	// The event type associated with this metric result. This field, when present, will always be `click`.
	EventType *string `json:"event_type,omitempty"`

	Results []MetricTokenAggregationResult `json:"results,omitempty"`
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

// MetricTokenResponse : The response generated from a call to a **metrics** method that evaluates tokens.
type MetricTokenResponse struct {
	Aggregations []MetricTokenAggregation `json:"aggregations,omitempty"`
}

// NluEnrichmentCategories : An object that indicates the Categories enrichment will be applied to the specified field.
type NluEnrichmentCategories struct {
}

// NluEnrichmentEmotion : An object specifying the emotion detection enrichment and related parameters.
type NluEnrichmentEmotion struct {

	// When `true`, emotion detection is performed on the entire field.
	Document *bool `json:"document,omitempty"`

	// A comma-separated list of target strings that will have any associated emotions detected.
	Targets []string `json:"targets,omitempty"`
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
	SentenceLocation *bool `json:"sentence_location,omitempty"`

	// The enrichement model to use with entity extraction. May be a custom model provided by Watson Knowledge Studio, the
	// public model for use with Knowledge Graph `en-news`, or the default public model `alchemy`.
	Model *string `json:"model,omitempty"`
}

// NluEnrichmentFeatures : NluEnrichmentFeatures struct
type NluEnrichmentFeatures struct {

	// An object specifying the Keyword enrichment and related parameters.
	Keywords *NluEnrichmentKeywords `json:"keywords,omitempty"`

	// An object speficying the Entities enrichment and related parameters.
	Entities *NluEnrichmentEntities `json:"entities,omitempty"`

	// An object specifying the sentiment extraction enrichment and related parameters.
	Sentiment *NluEnrichmentSentiment `json:"sentiment,omitempty"`

	// An object specifying the emotion detection enrichment and related parameters.
	Emotion *NluEnrichmentEmotion `json:"emotion,omitempty"`

	// An object specifying the categories enrichment and related parameters.
	Categories *NluEnrichmentCategories `json:"categories,omitempty"`

	// An object specifiying the semantic roles enrichment and related parameters.
	SemanticRoles *NluEnrichmentSemanticRoles `json:"semantic_roles,omitempty"`

	// An object specifying the relations enrichment and related parameters.
	Relations *NluEnrichmentRelations `json:"relations,omitempty"`
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

// NluEnrichmentRelations : An object specifying the relations enrichment and related parameters.
type NluEnrichmentRelations struct {

	// *For use with `natural_language_understanding` enrichments only.* The enrichement model to use with relationship
	// extraction. May be a custom model provided by Watson Knowledge Studio, the public model for use with Knowledge Graph
	// `en-news`, the default is`en-news`.
	Model *string `json:"model,omitempty"`
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

// NluEnrichmentSentiment : An object specifying the sentiment extraction enrichment and related parameters.
type NluEnrichmentSentiment struct {

	// When `true`, sentiment analysis is performed on the entire field.
	Document *bool `json:"document,omitempty"`

	// A comma-separated list of target strings that will have any associated sentiment analyzed.
	Targets []string `json:"targets,omitempty"`
}

// NormalizationOperation : NormalizationOperation struct
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
	// **remove_nulls** - Removes all nested null (blank) field values from the JSON tree. **source_field** and
	// **destination_field** are ignored by this operation because _remove_nulls_ operates on the entire JSON tree.
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
// **remove_nulls** - Removes all nested null (blank) field values from the JSON tree. **source_field** and
// **destination_field** are ignored by this operation because _remove_nulls_ operates on the entire JSON tree.
// Typically, **remove_nulls** is invoked as the last normalization operation (if it is invoked at all, it can be
// time-expensive).
const (
	NormalizationOperation_Operation_Copy        = "copy"
	NormalizationOperation_Operation_Merge       = "merge"
	NormalizationOperation_Operation_Move        = "move"
	NormalizationOperation_Operation_Remove      = "remove"
	NormalizationOperation_Operation_RemoveNulls = "remove_nulls"
)

// Notice : A notice produced for the collection.
type Notice struct {

	// Identifies the notice. Many notices might have the same ID. This field exists so that user applications can
	// programmatically identify a notice and take automatic corrective action.
	NoticeID *string `json:"notice_id,omitempty"`

	// The creation date of the collection in the format yyyy-MM-dd'T'HH:mm:ss.SSS'Z'.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// Unique identifier of the document.
	DocumentID *string `json:"document_id,omitempty"`

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

// PdfHeadingDetection : PdfHeadingDetection struct
type PdfHeadingDetection struct {
	Fonts []FontSetting `json:"fonts,omitempty"`
}

// PdfSettings : A list of PDF conversion settings.
type PdfSettings struct {
	Heading *PdfHeadingDetection `json:"heading,omitempty"`
}

// QueryAggregation : An aggregation produced by the Discovery service to analyze the input provided.
type QueryAggregation struct {

	// The type of aggregation command used. For example: term, filter, max, min, etc.
	Type *string `json:"type,omitempty"`

	Results []AggregationResult `json:"results,omitempty"`

	// Number of matching results.
	MatchingResults *int64 `json:"matching_results,omitempty"`

	// Aggregations returned by the Discovery service.
	Aggregations []QueryAggregation `json:"aggregations,omitempty"`
}

// QueryEntitiesContext : Entity text to provide context for the queried entity and rank based on that association. For example, if you wanted
// to query the city of London in England your query would look for `London` with the context of `England`.
type QueryEntitiesContext struct {

	// Entity text to provide context for the queried entity and rank based on that association. For example, if you wanted
	// to query the city of London in England your query would look for `London` with the context of `England`.
	Text *string `json:"text,omitempty"`
}

// QueryEntitiesEntity : A text string that appears within the entity text field.
type QueryEntitiesEntity struct {

	// Entity text content.
	Text *string `json:"text,omitempty"`

	// The type of the specified entity.
	Type *string `json:"type,omitempty"`
}

// QueryEntitiesOptions : The queryEntities options.
type QueryEntitiesOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The ID of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// The entity query feature to perform. Supported features are `disambiguate` and `similar_entities`.
	Feature *string `json:"feature,omitempty"`

	// A text string that appears within the entity text field.
	Entity *QueryEntitiesEntity `json:"entity,omitempty"`

	// Entity text to provide context for the queried entity and rank based on that association. For example, if you wanted
	// to query the city of London in England your query would look for `London` with the context of `England`.
	Context *QueryEntitiesContext `json:"context,omitempty"`

	// The number of results to return. The default is `10`. The maximum is `1000`.
	Count *int64 `json:"count,omitempty"`

	// The number of evidence items to return for each result. The default is `0`. The maximum number of evidence items per
	// query is 10,000.
	EvidenceCount *int64 `json:"evidence_count,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewQueryEntitiesOptions : Instantiate QueryEntitiesOptions
func (discovery *DiscoveryV1) NewQueryEntitiesOptions(environmentID string, collectionID string) *QueryEntitiesOptions {
	return &QueryEntitiesOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *QueryEntitiesOptions) SetEnvironmentID(environmentID string) *QueryEntitiesOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *QueryEntitiesOptions) SetCollectionID(collectionID string) *QueryEntitiesOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
}

// SetFeature : Allow user to set Feature
func (options *QueryEntitiesOptions) SetFeature(feature string) *QueryEntitiesOptions {
	options.Feature = core.StringPtr(feature)
	return options
}

// SetEntity : Allow user to set Entity
func (options *QueryEntitiesOptions) SetEntity(entity *QueryEntitiesEntity) *QueryEntitiesOptions {
	options.Entity = entity
	return options
}

// SetContext : Allow user to set Context
func (options *QueryEntitiesOptions) SetContext(context *QueryEntitiesContext) *QueryEntitiesOptions {
	options.Context = context
	return options
}

// SetCount : Allow user to set Count
func (options *QueryEntitiesOptions) SetCount(count int64) *QueryEntitiesOptions {
	options.Count = core.Int64Ptr(count)
	return options
}

// SetEvidenceCount : Allow user to set EvidenceCount
func (options *QueryEntitiesOptions) SetEvidenceCount(evidenceCount int64) *QueryEntitiesOptions {
	options.EvidenceCount = core.Int64Ptr(evidenceCount)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *QueryEntitiesOptions) SetHeaders(param map[string]string) *QueryEntitiesOptions {
	options.Headers = param
	return options
}

// QueryEntitiesResponse : An array of entities resulting from the query.
type QueryEntitiesResponse struct {
	Entities []QueryEntitiesResponseItem `json:"entities,omitempty"`
}

// QueryEntitiesResponseItem : Object containing Entity query response information.
type QueryEntitiesResponseItem struct {

	// Entity text content.
	Text *string `json:"text,omitempty"`

	// The type of the result entity.
	Type *string `json:"type,omitempty"`

	// List of different evidentiary items to support the result.
	Evidence []QueryEvidence `json:"evidence,omitempty"`
}

// QueryEvidence : Description of evidence location supporting Knoweldge Graph query result.
type QueryEvidence struct {

	// The docuemnt ID (as indexed in Discovery) of the evidence location.
	DocumentID *string `json:"document_id,omitempty"`

	// The field of the document where the supporting evidence was identified.
	Field *string `json:"field,omitempty"`

	// The start location of the evidence in the identified field. This value is inclusive.
	StartOffset *int64 `json:"start_offset,omitempty"`

	// The end location of the evidence in the identified field. This value is inclusive.
	EndOffset *int64 `json:"end_offset,omitempty"`

	// An array of entity objects that show evidence of the result.
	Entities []QueryEvidenceEntity `json:"entities,omitempty"`
}

// QueryEvidenceEntity : Entity description and location within evidence field.
type QueryEvidenceEntity struct {

	// The entity type for this entity. Possible types vary based on model used.
	Type *string `json:"type,omitempty"`

	// The original text of this entity as found in the evidence field.
	Text *string `json:"text,omitempty"`

	// The start location of the entity text in the identified field. This value is inclusive.
	StartOffset *int64 `json:"start_offset,omitempty"`

	// The end location of the entity text in the identified field. This value is exclusive.
	EndOffset *int64 `json:"end_offset,omitempty"`
}

// QueryFilterType : QueryFilterType struct
type QueryFilterType struct {

	// A comma-separated list of types to exclude.
	Exclude []string `json:"exclude,omitempty"`

	// A comma-separated list of types to include. All other types are excluded.
	Include []string `json:"include,omitempty"`
}

// QueryLogOptions : The queryLog options.
type QueryLogOptions struct {

	// A cacheable query that excludes documents that don't mention the query content. Filter searches are better for
	// metadata-type searches and for assessing the concepts in the data set.
	Filter *string `json:"filter,omitempty"`

	// A query search returns all documents in your data set with full enrichments and full text, but with the most
	// relevant documents listed first. Use a query search when you want to find the most relevant search results. You
	// cannot use **natural_language_query** and **query** at the same time.
	Query *string `json:"query,omitempty"`

	// Number of results to return.
	Count *int64 `json:"count,omitempty"`

	// The number of query results to skip at the beginning. For example, if the total number of results that are returned
	// is 10 and the offset is 8, it returns the last two results.
	Offset *int64 `json:"offset,omitempty"`

	// A comma-separated list of fields in the document to sort on. You can optionally specify a sort direction by
	// prefixing the field with `-` for descending or `+` for ascending. Ascending is the default sort direction if no
	// prefix is specified.
	Sort []string `json:"sort,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewQueryLogOptions : Instantiate QueryLogOptions
func (discovery *DiscoveryV1) NewQueryLogOptions() *QueryLogOptions {
	return &QueryLogOptions{}
}

// SetFilter : Allow user to set Filter
func (options *QueryLogOptions) SetFilter(filter string) *QueryLogOptions {
	options.Filter = core.StringPtr(filter)
	return options
}

// SetQuery : Allow user to set Query
func (options *QueryLogOptions) SetQuery(query string) *QueryLogOptions {
	options.Query = core.StringPtr(query)
	return options
}

// SetCount : Allow user to set Count
func (options *QueryLogOptions) SetCount(count int64) *QueryLogOptions {
	options.Count = core.Int64Ptr(count)
	return options
}

// SetOffset : Allow user to set Offset
func (options *QueryLogOptions) SetOffset(offset int64) *QueryLogOptions {
	options.Offset = core.Int64Ptr(offset)
	return options
}

// SetSort : Allow user to set Sort
func (options *QueryLogOptions) SetSort(sort []string) *QueryLogOptions {
	options.Sort = sort
	return options
}

// SetHeaders : Allow user to set Headers
func (options *QueryLogOptions) SetHeaders(param map[string]string) *QueryLogOptions {
	options.Headers = param
	return options
}

// QueryNoticesOptions : The queryNotices options.
type QueryNoticesOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The ID of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// A cacheable query that excludes documents that don't mention the query content. Filter searches are better for
	// metadata-type searches and for assessing the concepts in the data set.
	Filter *string `json:"filter,omitempty"`

	// A query search returns all documents in your data set with full enrichments and full text, but with the most
	// relevant documents listed first. Use a query search when you want to find the most relevant search results. You
	// cannot use **natural_language_query** and **query** at the same time.
	Query *string `json:"query,omitempty"`

	// A natural language query that returns relevant documents by utilizing training data and natural language
	// understanding. You cannot use **natural_language_query** and **query** at the same time.
	NaturalLanguageQuery *string `json:"natural_language_query,omitempty"`

	// A passages query that returns the most relevant passages from the results.
	Passages *bool `json:"passages,omitempty"`

	// An aggregation search that returns an exact answer by combining query search with filters. Useful for applications
	// to build lists, tables, and time series. For a full list of possible aggregations, see the Query reference.
	Aggregation *string `json:"aggregation,omitempty"`

	// Number of results to return.
	Count *int64 `json:"count,omitempty"`

	// A comma-separated list of the portion of the document hierarchy to return.
	ReturnFields []string `json:"return,omitempty"`

	// The number of query results to skip at the beginning. For example, if the total number of results that are returned
	// is 10 and the offset is 8, it returns the last two results.
	Offset *int64 `json:"offset,omitempty"`

	// A comma-separated list of fields in the document to sort on. You can optionally specify a sort direction by
	// prefixing the field with `-` for descending or `+` for ascending. Ascending is the default sort direction if no
	// prefix is specified.
	Sort []string `json:"sort,omitempty"`

	// When true, a highlight field is returned for each result which contains the fields which match the query with
	// `<em></em>` tags around the matching query terms.
	Highlight *bool `json:"highlight,omitempty"`

	// A comma-separated list of fields that passages are drawn from. If this parameter not specified, then all top-level
	// fields are included.
	PassagesFields []string `json:"passages.fields,omitempty"`

	// The maximum number of passages to return. The search returns fewer passages if the requested total is not found.
	PassagesCount *int64 `json:"passages.count,omitempty"`

	// The approximate number of characters that any one passage will have.
	PassagesCharacters *int64 `json:"passages.characters,omitempty"`

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
	SimilarDocumentIds []string `json:"similar.document_ids,omitempty"`

	// A comma-separated list of field names that are used as a basis for comparison to identify similar documents. If not
	// specified, the entire document is used for comparison.
	SimilarFields []string `json:"similar.fields,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewQueryNoticesOptions : Instantiate QueryNoticesOptions
func (discovery *DiscoveryV1) NewQueryNoticesOptions(environmentID string, collectionID string) *QueryNoticesOptions {
	return &QueryNoticesOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *QueryNoticesOptions) SetEnvironmentID(environmentID string) *QueryNoticesOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *QueryNoticesOptions) SetCollectionID(collectionID string) *QueryNoticesOptions {
	options.CollectionID = core.StringPtr(collectionID)
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

// SetPassages : Allow user to set Passages
func (options *QueryNoticesOptions) SetPassages(passages bool) *QueryNoticesOptions {
	options.Passages = core.BoolPtr(passages)
	return options
}

// SetAggregation : Allow user to set Aggregation
func (options *QueryNoticesOptions) SetAggregation(aggregation string) *QueryNoticesOptions {
	options.Aggregation = core.StringPtr(aggregation)
	return options
}

// SetCount : Allow user to set Count
func (options *QueryNoticesOptions) SetCount(count int64) *QueryNoticesOptions {
	options.Count = core.Int64Ptr(count)
	return options
}

// SetReturnFields : Allow user to set ReturnFields
func (options *QueryNoticesOptions) SetReturnFields(returnFields []string) *QueryNoticesOptions {
	options.ReturnFields = returnFields
	return options
}

// SetOffset : Allow user to set Offset
func (options *QueryNoticesOptions) SetOffset(offset int64) *QueryNoticesOptions {
	options.Offset = core.Int64Ptr(offset)
	return options
}

// SetSort : Allow user to set Sort
func (options *QueryNoticesOptions) SetSort(sort []string) *QueryNoticesOptions {
	options.Sort = sort
	return options
}

// SetHighlight : Allow user to set Highlight
func (options *QueryNoticesOptions) SetHighlight(highlight bool) *QueryNoticesOptions {
	options.Highlight = core.BoolPtr(highlight)
	return options
}

// SetPassagesFields : Allow user to set PassagesFields
func (options *QueryNoticesOptions) SetPassagesFields(passagesFields []string) *QueryNoticesOptions {
	options.PassagesFields = passagesFields
	return options
}

// SetPassagesCount : Allow user to set PassagesCount
func (options *QueryNoticesOptions) SetPassagesCount(passagesCount int64) *QueryNoticesOptions {
	options.PassagesCount = core.Int64Ptr(passagesCount)
	return options
}

// SetPassagesCharacters : Allow user to set PassagesCharacters
func (options *QueryNoticesOptions) SetPassagesCharacters(passagesCharacters int64) *QueryNoticesOptions {
	options.PassagesCharacters = core.Int64Ptr(passagesCharacters)
	return options
}

// SetDeduplicateField : Allow user to set DeduplicateField
func (options *QueryNoticesOptions) SetDeduplicateField(deduplicateField string) *QueryNoticesOptions {
	options.DeduplicateField = core.StringPtr(deduplicateField)
	return options
}

// SetSimilar : Allow user to set Similar
func (options *QueryNoticesOptions) SetSimilar(similar bool) *QueryNoticesOptions {
	options.Similar = core.BoolPtr(similar)
	return options
}

// SetSimilarDocumentIds : Allow user to set SimilarDocumentIds
func (options *QueryNoticesOptions) SetSimilarDocumentIds(similarDocumentIds []string) *QueryNoticesOptions {
	options.SimilarDocumentIds = similarDocumentIds
	return options
}

// SetSimilarFields : Allow user to set SimilarFields
func (options *QueryNoticesOptions) SetSimilarFields(similarFields []string) *QueryNoticesOptions {
	options.SimilarFields = similarFields
	return options
}

// SetHeaders : Allow user to set Headers
func (options *QueryNoticesOptions) SetHeaders(param map[string]string) *QueryNoticesOptions {
	options.Headers = param
	return options
}

// QueryNoticesResponse : QueryNoticesResponse struct
type QueryNoticesResponse struct {
	MatchingResults *int64 `json:"matching_results,omitempty"`

	Results []QueryNoticesResult `json:"results,omitempty"`

	Aggregations []QueryAggregation `json:"aggregations,omitempty"`

	Passages []QueryPassages `json:"passages,omitempty"`

	DuplicatesRemoved *int64 `json:"duplicates_removed,omitempty"`
}

// QueryNoticesResult : QueryNoticesResult struct
type QueryNoticesResult struct {

	// The unique identifier of the document.
	ID *string `json:"id,omitempty"`

	// *Deprecated* This field is now part of the **result_metadata** object.
	Score *float64 `json:"score,omitempty"`

	// Metadata of the document.
	Metadata interface{} `json:"metadata,omitempty"`

	// The collection ID of the collection containing the document for this result.
	CollectionID *string `json:"collection_id,omitempty"`

	// Metadata of the query result.
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
}

// Constants associated with the QueryNoticesResult.FileType property.
// The type of the original source file.
const (
	QueryNoticesResult_FileType_HTML = "html"
	QueryNoticesResult_FileType_JSON = "json"
	QueryNoticesResult_FileType_Pdf  = "pdf"
	QueryNoticesResult_FileType_Word = "word"
)

// QueryOptions : The query options.
type QueryOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The ID of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// A cacheable query that excludes documents that don't mention the query content. Filter searches are better for
	// metadata-type searches and for assessing the concepts in the data set.
	Filter *string `json:"filter,omitempty"`

	// A query search returns all documents in your data set with full enrichments and full text, but with the most
	// relevant documents listed first. Use a query search when you want to find the most relevant search results. You
	// cannot use **natural_language_query** and **query** at the same time.
	Query *string `json:"query,omitempty"`

	// A natural language query that returns relevant documents by utilizing training data and natural language
	// understanding. You cannot use **natural_language_query** and **query** at the same time.
	NaturalLanguageQuery *string `json:"natural_language_query,omitempty"`

	// A passages query that returns the most relevant passages from the results.
	Passages *bool `json:"passages,omitempty"`

	// An aggregation search that returns an exact answer by combining query search with filters. Useful for applications
	// to build lists, tables, and time series. For a full list of possible aggregations, see the Query reference.
	Aggregation *string `json:"aggregation,omitempty"`

	// Number of results to return.
	Count *int64 `json:"count,omitempty"`

	// A comma-separated list of the portion of the document hierarchy to return.
	ReturnFields *string `json:"return,omitempty"`

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

	// When `true` and used with a Watson Discovery News collection, duplicate results (based on the contents of the
	// **title** field) are removed. Duplicate comparison is limited to the current query only; **offset** is not
	// considered. This parameter is currently Beta functionality.
	Deduplicate *bool `json:"deduplicate,omitempty"`

	// When specified, duplicate results based on the field specified are removed from the returned results. Duplicate
	// comparison is limited to the current query only, **offset** is not considered. This parameter is currently Beta
	// functionality.
	DeduplicateField *string `json:"deduplicate.field,omitempty"`

	// A comma-separated list of collection IDs to be queried against. Required when querying multiple collections, invalid
	// when performing a single collection query.
	CollectionIds *string `json:"collection_ids,omitempty"`

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
	LoggingOptOut *bool `json:"X-Watson-Logging-Opt-Out,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewQueryOptions : Instantiate QueryOptions
func (discovery *DiscoveryV1) NewQueryOptions(environmentID string, collectionID string) *QueryOptions {
	return &QueryOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *QueryOptions) SetEnvironmentID(environmentID string) *QueryOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *QueryOptions) SetCollectionID(collectionID string) *QueryOptions {
	options.CollectionID = core.StringPtr(collectionID)
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

// SetPassages : Allow user to set Passages
func (options *QueryOptions) SetPassages(passages bool) *QueryOptions {
	options.Passages = core.BoolPtr(passages)
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

// SetReturnFields : Allow user to set ReturnFields
func (options *QueryOptions) SetReturnFields(returnFields string) *QueryOptions {
	options.ReturnFields = core.StringPtr(returnFields)
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

// SetPassagesFields : Allow user to set PassagesFields
func (options *QueryOptions) SetPassagesFields(passagesFields string) *QueryOptions {
	options.PassagesFields = core.StringPtr(passagesFields)
	return options
}

// SetPassagesCount : Allow user to set PassagesCount
func (options *QueryOptions) SetPassagesCount(passagesCount int64) *QueryOptions {
	options.PassagesCount = core.Int64Ptr(passagesCount)
	return options
}

// SetPassagesCharacters : Allow user to set PassagesCharacters
func (options *QueryOptions) SetPassagesCharacters(passagesCharacters int64) *QueryOptions {
	options.PassagesCharacters = core.Int64Ptr(passagesCharacters)
	return options
}

// SetDeduplicate : Allow user to set Deduplicate
func (options *QueryOptions) SetDeduplicate(deduplicate bool) *QueryOptions {
	options.Deduplicate = core.BoolPtr(deduplicate)
	return options
}

// SetDeduplicateField : Allow user to set DeduplicateField
func (options *QueryOptions) SetDeduplicateField(deduplicateField string) *QueryOptions {
	options.DeduplicateField = core.StringPtr(deduplicateField)
	return options
}

// SetCollectionIds : Allow user to set CollectionIds
func (options *QueryOptions) SetCollectionIds(collectionIds string) *QueryOptions {
	options.CollectionIds = core.StringPtr(collectionIds)
	return options
}

// SetSimilar : Allow user to set Similar
func (options *QueryOptions) SetSimilar(similar bool) *QueryOptions {
	options.Similar = core.BoolPtr(similar)
	return options
}

// SetSimilarDocumentIds : Allow user to set SimilarDocumentIds
func (options *QueryOptions) SetSimilarDocumentIds(similarDocumentIds string) *QueryOptions {
	options.SimilarDocumentIds = core.StringPtr(similarDocumentIds)
	return options
}

// SetSimilarFields : Allow user to set SimilarFields
func (options *QueryOptions) SetSimilarFields(similarFields string) *QueryOptions {
	options.SimilarFields = core.StringPtr(similarFields)
	return options
}

// SetBias : Allow user to set Bias
func (options *QueryOptions) SetBias(bias string) *QueryOptions {
	options.Bias = core.StringPtr(bias)
	return options
}

// SetLoggingOptOut : Allow user to set LoggingOptOut
func (options *QueryOptions) SetLoggingOptOut(loggingOptOut bool) *QueryOptions {
	options.LoggingOptOut = core.BoolPtr(loggingOptOut)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *QueryOptions) SetHeaders(param map[string]string) *QueryOptions {
	options.Headers = param
	return options
}

// QueryPassages : QueryPassages struct
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

// QueryRelationsArgument : QueryRelationsArgument struct
type QueryRelationsArgument struct {
	Entities []QueryEntitiesEntity `json:"entities,omitempty"`
}

// QueryRelationsEntity : QueryRelationsEntity struct
type QueryRelationsEntity struct {

	// Entity text content.
	Text *string `json:"text,omitempty"`

	// The type of the specified entity.
	Type *string `json:"type,omitempty"`

	// If false, implicit querying is performed. The default is `false`.
	Exact *bool `json:"exact,omitempty"`
}

// QueryRelationsFilter : QueryRelationsFilter struct
type QueryRelationsFilter struct {

	// A list of relation types to include or exclude from the query.
	RelationTypes *QueryFilterType `json:"relation_types,omitempty"`

	// A list of entity types to include or exclude from the query.
	EntityTypes *QueryFilterType `json:"entity_types,omitempty"`

	// A comma-separated list of document IDs to include in the query.
	DocumentIds []string `json:"document_ids,omitempty"`
}

// QueryRelationsOptions : The queryRelations options.
type QueryRelationsOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The ID of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// An array of entities to find relationships for.
	Entities []QueryRelationsEntity `json:"entities,omitempty"`

	// Entity text to provide context for the queried entity and rank based on that association. For example, if you wanted
	// to query the city of London in England your query would look for `London` with the context of `England`.
	Context *QueryEntitiesContext `json:"context,omitempty"`

	// The sorting method for the relationships, can be `score` or `frequency`. `frequency` is the number of unique times
	// each entity is identified. The default is `score`. This parameter cannot be used in the same query as the **bias**
	// parameter.
	Sort *string `json:"sort,omitempty"`

	// Filters to apply to the relationship query.
	Filter *QueryRelationsFilter `json:"filter,omitempty"`

	// The number of results to return. The default is `10`. The maximum is `1000`.
	Count *int64 `json:"count,omitempty"`

	// The number of evidence items to return for each result. The default is `0`. The maximum number of evidence items per
	// query is 10,000.
	EvidenceCount *int64 `json:"evidence_count,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the QueryRelationsOptions.Sort property.
// The sorting method for the relationships, can be `score` or `frequency`. `frequency` is the number of unique times
// each entity is identified. The default is `score`. This parameter cannot be used in the same query as the **bias**
// parameter.
const (
	QueryRelationsOptions_Sort_Frequency = "frequency"
	QueryRelationsOptions_Sort_Score     = "score"
)

// NewQueryRelationsOptions : Instantiate QueryRelationsOptions
func (discovery *DiscoveryV1) NewQueryRelationsOptions(environmentID string, collectionID string) *QueryRelationsOptions {
	return &QueryRelationsOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *QueryRelationsOptions) SetEnvironmentID(environmentID string) *QueryRelationsOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *QueryRelationsOptions) SetCollectionID(collectionID string) *QueryRelationsOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
}

// SetEntities : Allow user to set Entities
func (options *QueryRelationsOptions) SetEntities(entities []QueryRelationsEntity) *QueryRelationsOptions {
	options.Entities = entities
	return options
}

// SetContext : Allow user to set Context
func (options *QueryRelationsOptions) SetContext(context *QueryEntitiesContext) *QueryRelationsOptions {
	options.Context = context
	return options
}

// SetSort : Allow user to set Sort
func (options *QueryRelationsOptions) SetSort(sort string) *QueryRelationsOptions {
	options.Sort = core.StringPtr(sort)
	return options
}

// SetFilter : Allow user to set Filter
func (options *QueryRelationsOptions) SetFilter(filter *QueryRelationsFilter) *QueryRelationsOptions {
	options.Filter = filter
	return options
}

// SetCount : Allow user to set Count
func (options *QueryRelationsOptions) SetCount(count int64) *QueryRelationsOptions {
	options.Count = core.Int64Ptr(count)
	return options
}

// SetEvidenceCount : Allow user to set EvidenceCount
func (options *QueryRelationsOptions) SetEvidenceCount(evidenceCount int64) *QueryRelationsOptions {
	options.EvidenceCount = core.Int64Ptr(evidenceCount)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *QueryRelationsOptions) SetHeaders(param map[string]string) *QueryRelationsOptions {
	options.Headers = param
	return options
}

// QueryRelationsRelationship : QueryRelationsRelationship struct
type QueryRelationsRelationship struct {

	// The identified relationship type.
	Type *string `json:"type,omitempty"`

	// The number of times the relationship is mentioned.
	Frequency *int64 `json:"frequency,omitempty"`

	// Information about the relationship.
	Arguments []QueryRelationsArgument `json:"arguments,omitempty"`

	// List of different evidentiary items to support the result.
	Evidence []QueryEvidence `json:"evidence,omitempty"`
}

// QueryRelationsResponse : QueryRelationsResponse struct
type QueryRelationsResponse struct {
	Relations []QueryRelationsRelationship `json:"relations,omitempty"`
}

// QueryResponse : A response containing the documents and aggregations for the query.
type QueryResponse struct {
	MatchingResults *int64 `json:"matching_results,omitempty"`

	Results []QueryResult `json:"results,omitempty"`

	Aggregations []QueryAggregation `json:"aggregations,omitempty"`

	Passages []QueryPassages `json:"passages,omitempty"`

	DuplicatesRemoved *int64 `json:"duplicates_removed,omitempty"`

	// The session token for this query. The session token can be used to add events associated with this query to the
	// query and event log.
	//
	// **Important:** Session tokens are case sensitive.
	SessionToken *string `json:"session_token,omitempty"`
}

// QueryResult : QueryResult struct
type QueryResult struct {

	// The unique identifier of the document.
	ID *string `json:"id,omitempty"`

	// *Deprecated* This field is now part of the **result_metadata** object.
	Score *float64 `json:"score,omitempty"`

	// Metadata of the document.
	Metadata interface{} `json:"metadata,omitempty"`

	// The collection ID of the collection containing the document for this result.
	CollectionID *string `json:"collection_id,omitempty"`

	// Metadata of the query result.
	ResultMetadata *QueryResultMetadata `json:"result_metadata,omitempty"`
}

// QueryResultMetadata : Metadata of a query result.
type QueryResultMetadata struct {

	// An unbounded measure of the relevance of a particular result, dependent on the query and matching document. A higher
	// score indicates a greater match to the query parameters.
	Score *float64 `json:"score,omitempty"`

	// The confidence score for the given result. Calculated based on how relevant the result is estimated to be, compared
	// to a trained relevancy model. confidence can range from `0.0` to `1.0`. The higher the number, the more relevant the
	// document.
	Confidence *float64 `json:"confidence,omitempty"`
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
	SearchStatus_Status_INSUFFICENTDATA = "INSUFFICENT_DATA"
	SearchStatus_Status_NOTAPPLICABLE   = "NOT_APPLICABLE"
	SearchStatus_Status_NODATA          = "NO_DATA"
	SearchStatus_Status_TRAINED         = "TRAINED"
	SearchStatus_Status_TRAINING        = "TRAINING"
)

// SegmentSettings : A list of Document Segmentation settings.
type SegmentSettings struct {

	// Enables/disables the Document Segmentation feature.
	Enabled *bool `json:"enabled,omitempty"`

	// Defines the heading level that splits into document segments. Valid values are h1, h2, h3, h4, h5, h6.
	SelectorTags []string `json:"selector_tags,omitempty"`
}

// Source : Object containing source parameters for the configuration.
type Source struct {

	// The type of source to connect to.
	// -  `box` indicates the configuration is to connect an instance of Enterprise Box.
	// -  `salesforce` indicates the configuration is to connect to Salesforce.
	// -  `sharepoint` indicates the configuration is to connect to Microsoft SharePoint Online.
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
const (
	Source_Type_Box        = "box"
	Source_Type_Salesforce = "salesforce"
	Source_Type_Sharepoint = "sharepoint"
)

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

// SourceOptionsObject : Object that defines a Salesforce document object type crawl with this configuration.
type SourceOptionsObject struct {

	// The name of the Salesforce document object to crawl. For example, `case`.
	Name *string `json:"name" validate:"required"`

	// The maximum number of documents to crawl for this document object. By default, all documents in the document object
	// are crawled.
	Limit *int64 `json:"limit,omitempty"`
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
	// -  `daily`: Runs every day between 00:00 and 06:00.
	// -  `weekly`: Runs every week on Sunday between 00:00 and 06:00.
	// -  `monthly`: Runs the on the first Sunday of every month between 00:00 and 06:00.
	Frequency *string `json:"frequency,omitempty"`
}

// Constants associated with the SourceSchedule.Frequency property.
// The crawl schedule in the specified **time_zone**.
//
// -  `daily`: Runs every day between 00:00 and 06:00.
// -  `weekly`: Runs every week on Sunday between 00:00 and 06:00.
// -  `monthly`: Runs the on the first Sunday of every month between 00:00 and 06:00.
const (
	SourceSchedule_Frequency_Daily   = "daily"
	SourceSchedule_Frequency_Monthly = "monthly"
	SourceSchedule_Frequency_Weekly  = "weekly"
)

// SourceStatus : Object containing source crawl status information.
type SourceStatus struct {

	// The current status of the source crawl for this collection. This field returns `not_configured` if the default
	// configuration for this source does not have a **source** object defined.
	//
	// -  `running` indicates that a crawl to fetch more documents is in progress.
	// -  `complete` indicates that the crawl has completed with no errors.
	// -  `complete_with_notices` indicates that some notices were generated during the crawl. Notices can be checked by
	// using the **notices** query method.
	// -  `stopped` indicates that the crawl has stopped but is not complete.
	Status *string `json:"status,omitempty"`

	// Date in UTC format indicating when the last crawl was attempted. If `null`, no crawl was completed.
	LastUpdated *strfmt.DateTime `json:"last_updated,omitempty"`
}

// Constants associated with the SourceStatus.Status property.
// The current status of the source crawl for this collection. This field returns `not_configured` if the default
// configuration for this source does not have a **source** object defined.
//
// -  `running` indicates that a crawl to fetch more documents is in progress.
// -  `complete` indicates that the crawl has completed with no errors.
// -  `complete_with_notices` indicates that some notices were generated during the crawl. Notices can be checked by
// using the **notices** query method.
// -  `stopped` indicates that the crawl has stopped but is not complete.
const (
	SourceStatus_Status_Complete            = "complete"
	SourceStatus_Status_CompleteWithNotices = "complete_with_notices"
	SourceStatus_Status_NotConfigured       = "not_configured"
	SourceStatus_Status_Running             = "running"
	SourceStatus_Status_Stopped             = "stopped"
)

// TestConfigurationInEnvironmentOptions : The testConfigurationInEnvironment options.
type TestConfigurationInEnvironmentOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The configuration to use to process the document. If this part is provided, then the provided configuration is used
	// to process the document. If the **configuration_id** is also provided (both are present at the same time), then
	// request is rejected. The maximum supported configuration size is 1 MB. Configuration parts larger than 1 MB are
	// rejected.
	// See the `GET /configurations/{configuration_id}` operation for an example configuration.
	Configuration *string `json:"configuration,omitempty"`

	// Specify to only run the input document through the given step instead of running the input document through the
	// entire ingestion workflow. Valid values are `convert`, `enrich`, and `normalize`.
	Step *string `json:"step,omitempty"`

	// The ID of the configuration to use to process the document. If the **configuration** form part is also provided
	// (both are present at the same time), then the request will be rejected.
	ConfigurationID *string `json:"configuration_id,omitempty"`

	// The content of the document to ingest. The maximum supported file size is 50 megabytes. Files larger than 50
	// megabytes is rejected.
	File *os.File `json:"file,omitempty"`

	// The filename for file.
	Filename *string `json:"filename,omitempty"`

	// If you're using the Data Crawler to upload your documents, you can test a document against the type of metadata that
	// the Data Crawler might send. The maximum supported metadata file size is 1 MB. Metadata parts larger than 1 MB are
	// rejected.
	// Example:  ``` {
	//   "Creator": "Johnny Appleseed",
	//   "Subject": "Apples"
	// } ```.
	Metadata *string `json:"metadata,omitempty"`

	// The content type of file. Values for this parameter can be obtained from the HttpMediaType class.
	FileContentType *string `json:"file_content_type,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the TestConfigurationInEnvironmentOptions.Step property.
// Specify to only run the input document through the given step instead of running the input document through the
// entire ingestion workflow. Valid values are `convert`, `enrich`, and `normalize`.
const (
	TestConfigurationInEnvironmentOptions_Step_EnrichmentsOutput        = "enrichments_output"
	TestConfigurationInEnvironmentOptions_Step_HTMLInput                = "html_input"
	TestConfigurationInEnvironmentOptions_Step_HTMLOutput               = "html_output"
	TestConfigurationInEnvironmentOptions_Step_JSONNormalizationsOutput = "json_normalizations_output"
	TestConfigurationInEnvironmentOptions_Step_JSONOutput               = "json_output"
	TestConfigurationInEnvironmentOptions_Step_NormalizationsOutput     = "normalizations_output"
)

// NewTestConfigurationInEnvironmentOptions : Instantiate TestConfigurationInEnvironmentOptions
func (discovery *DiscoveryV1) NewTestConfigurationInEnvironmentOptions(environmentID string) *TestConfigurationInEnvironmentOptions {
	return &TestConfigurationInEnvironmentOptions{
		EnvironmentID: core.StringPtr(environmentID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *TestConfigurationInEnvironmentOptions) SetEnvironmentID(environmentID string) *TestConfigurationInEnvironmentOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetConfiguration : Allow user to set Configuration
func (options *TestConfigurationInEnvironmentOptions) SetConfiguration(configuration string) *TestConfigurationInEnvironmentOptions {
	options.Configuration = core.StringPtr(configuration)
	return options
}

// SetStep : Allow user to set Step
func (options *TestConfigurationInEnvironmentOptions) SetStep(step string) *TestConfigurationInEnvironmentOptions {
	options.Step = core.StringPtr(step)
	return options
}

// SetConfigurationID : Allow user to set ConfigurationID
func (options *TestConfigurationInEnvironmentOptions) SetConfigurationID(configurationID string) *TestConfigurationInEnvironmentOptions {
	options.ConfigurationID = core.StringPtr(configurationID)
	return options
}

// SetFile : Allow user to set File
func (options *TestConfigurationInEnvironmentOptions) SetFile(file *os.File) *TestConfigurationInEnvironmentOptions {
	options.File = file
	return options
}

// SetFilename : Allow user to set Filename
func (options *TestConfigurationInEnvironmentOptions) SetFilename(filename string) *TestConfigurationInEnvironmentOptions {
	options.Filename = core.StringPtr(filename)
	return options
}

// SetMetadata : Allow user to set Metadata
func (options *TestConfigurationInEnvironmentOptions) SetMetadata(metadata string) *TestConfigurationInEnvironmentOptions {
	options.Metadata = core.StringPtr(metadata)
	return options
}

// SetFileContentType : Allow user to set FileContentType
func (options *TestConfigurationInEnvironmentOptions) SetFileContentType(fileContentType string) *TestConfigurationInEnvironmentOptions {
	options.FileContentType = core.StringPtr(fileContentType)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *TestConfigurationInEnvironmentOptions) SetHeaders(param map[string]string) *TestConfigurationInEnvironmentOptions {
	options.Headers = param
	return options
}

// TestDocument : TestDocument struct
type TestDocument struct {

	// The unique identifier for the configuration.
	ConfigurationID *string `json:"configuration_id,omitempty"`

	// Status of the preview operation.
	Status *string `json:"status,omitempty"`

	// The number of 10-kB chunks of field data that were enriched. This can be used to estimate the cost of running a real
	// ingestion.
	EnrichedFieldUnits *int64 `json:"enriched_field_units,omitempty"`

	// Format of the test document.
	OriginalMediaType *string `json:"original_media_type,omitempty"`

	// An array of objects that describe each step in the preview process.
	Snapshots []DocumentSnapshot `json:"snapshots,omitempty"`

	// An array of notice messages about the preview operation.
	Notices []Notice `json:"notices,omitempty"`
}

// TokenDictRule : An object defining a single tokenizaion rule.
type TokenDictRule struct {

	// The string to tokenize.
	Text *string `json:"text,omitempty"`

	// Array of tokens that the `text` field is split into when found.
	Tokens []string `json:"tokens,omitempty"`

	// Array of tokens that represent the content of the `text` field in an alternate character set.
	Readings []string `json:"readings,omitempty"`

	// The part of speech that the `text` string belongs to. For example `noun`. Custom parts of speech can be specified.
	PartOfSpeech *string `json:"part_of_speech,omitempty"`
}

// TokenDictStatusResponse : Object describing the current status of the tokenization dictionary.
type TokenDictStatusResponse struct {

	// Current tokenization dictionary status for the specified collection.
	Status *string `json:"status,omitempty"`

	// The type for this dictionary. Always returns `tokenization_dictionary`.
	Type *string `json:"type,omitempty"`
}

// Constants associated with the TokenDictStatusResponse.Status property.
// Current tokenization dictionary status for the specified collection.
const (
	TokenDictStatusResponse_Status_Active   = "active"
	TokenDictStatusResponse_Status_NotFound = "not found"
	TokenDictStatusResponse_Status_Pending  = "pending"
)

// TopHitsResults : TopHitsResults struct
type TopHitsResults struct {

	// Number of matching results.
	MatchingResults *int64 `json:"matching_results,omitempty"`

	// Top results returned by the aggregation.
	Hits []QueryResult `json:"hits,omitempty"`
}

// TrainingDataSet : TrainingDataSet struct
type TrainingDataSet struct {
	EnvironmentID *string `json:"environment_id,omitempty"`

	CollectionID *string `json:"collection_id,omitempty"`

	Queries []TrainingQuery `json:"queries,omitempty"`
}

// TrainingExample : TrainingExample struct
type TrainingExample struct {
	DocumentID *string `json:"document_id,omitempty"`

	CrossReference *string `json:"cross_reference,omitempty"`

	Relevance *int64 `json:"relevance,omitempty"`
}

// TrainingExampleList : TrainingExampleList struct
type TrainingExampleList struct {
	Examples []TrainingExample `json:"examples,omitempty"`
}

// TrainingQuery : TrainingQuery struct
type TrainingQuery struct {
	QueryID *string `json:"query_id,omitempty"`

	NaturalLanguageQuery *string `json:"natural_language_query,omitempty"`

	Filter *string `json:"filter,omitempty"`

	Examples []TrainingExample `json:"examples,omitempty"`
}

// TrainingStatus : TrainingStatus struct
type TrainingStatus struct {
	TotalExamples *int64 `json:"total_examples,omitempty"`

	Available *bool `json:"available,omitempty"`

	Processing *bool `json:"processing,omitempty"`

	MinimumQueriesAdded *bool `json:"minimum_queries_added,omitempty"`

	MinimumExamplesAdded *bool `json:"minimum_examples_added,omitempty"`

	SufficientLabelDiversity *bool `json:"sufficient_label_diversity,omitempty"`

	Notices *int64 `json:"notices,omitempty"`

	SuccessfullyTrained *strfmt.DateTime `json:"successfully_trained,omitempty"`

	DataUpdated *strfmt.DateTime `json:"data_updated,omitempty"`
}

// UpdateCollectionOptions : The updateCollection options.
type UpdateCollectionOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The ID of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// The name of the collection.
	Name *string `json:"name,omitempty"`

	// A description of the collection.
	Description *string `json:"description,omitempty"`

	// The ID of the configuration in which the collection is to be updated.
	ConfigurationID *string `json:"configuration_id,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewUpdateCollectionOptions : Instantiate UpdateCollectionOptions
func (discovery *DiscoveryV1) NewUpdateCollectionOptions(environmentID string, collectionID string) *UpdateCollectionOptions {
	return &UpdateCollectionOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *UpdateCollectionOptions) SetEnvironmentID(environmentID string) *UpdateCollectionOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
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

// SetConfigurationID : Allow user to set ConfigurationID
func (options *UpdateCollectionOptions) SetConfigurationID(configurationID string) *UpdateCollectionOptions {
	options.ConfigurationID = core.StringPtr(configurationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateCollectionOptions) SetHeaders(param map[string]string) *UpdateCollectionOptions {
	options.Headers = param
	return options
}

// UpdateConfigurationOptions : The updateConfiguration options.
type UpdateConfigurationOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The ID of the configuration.
	ConfigurationID *string `json:"configuration_id" validate:"required"`

	// The name of the configuration.
	Name *string `json:"name,omitempty"`

	// The description of the configuration, if available.
	Description *string `json:"description,omitempty"`

	// The document conversion settings for the configuration.
	Conversions *Conversions `json:"conversions,omitempty"`

	// An array of document enrichment settings for the configuration.
	Enrichments []Enrichment `json:"enrichments,omitempty"`

	// Defines operations that can be used to transform the final output JSON into a normalized form. Operations are
	// executed in the order that they appear in the array.
	Normalizations []NormalizationOperation `json:"normalizations,omitempty"`

	// Object containing source parameters for the configuration.
	Source *Source `json:"source,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewUpdateConfigurationOptions : Instantiate UpdateConfigurationOptions
func (discovery *DiscoveryV1) NewUpdateConfigurationOptions(environmentID string, configurationID string) *UpdateConfigurationOptions {
	return &UpdateConfigurationOptions{
		EnvironmentID:   core.StringPtr(environmentID),
		ConfigurationID: core.StringPtr(configurationID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *UpdateConfigurationOptions) SetEnvironmentID(environmentID string) *UpdateConfigurationOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetConfigurationID : Allow user to set ConfigurationID
func (options *UpdateConfigurationOptions) SetConfigurationID(configurationID string) *UpdateConfigurationOptions {
	options.ConfigurationID = core.StringPtr(configurationID)
	return options
}

// SetName : Allow user to set Name
func (options *UpdateConfigurationOptions) SetName(name string) *UpdateConfigurationOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetDescription : Allow user to set Description
func (options *UpdateConfigurationOptions) SetDescription(description string) *UpdateConfigurationOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetConversions : Allow user to set Conversions
func (options *UpdateConfigurationOptions) SetConversions(conversions *Conversions) *UpdateConfigurationOptions {
	options.Conversions = conversions
	return options
}

// SetEnrichments : Allow user to set Enrichments
func (options *UpdateConfigurationOptions) SetEnrichments(enrichments []Enrichment) *UpdateConfigurationOptions {
	options.Enrichments = enrichments
	return options
}

// SetNormalizations : Allow user to set Normalizations
func (options *UpdateConfigurationOptions) SetNormalizations(normalizations []NormalizationOperation) *UpdateConfigurationOptions {
	options.Normalizations = normalizations
	return options
}

// SetSource : Allow user to set Source
func (options *UpdateConfigurationOptions) SetSource(source *Source) *UpdateConfigurationOptions {
	options.Source = source
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateConfigurationOptions) SetHeaders(param map[string]string) *UpdateConfigurationOptions {
	options.Headers = param
	return options
}

// UpdateCredentialsOptions : The updateCredentials options.
type UpdateCredentialsOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The unique identifier for a set of source credentials.
	CredentialID *string `json:"credential_id" validate:"required"`

	// The source that this credentials object connects to.
	// -  `box` indicates the credentials are used to connect an instance of Enterprise Box.
	// -  `salesforce` indicates the credentials are used to connect to Salesforce.
	// -  `sharepoint` indicates the credentials are used to connect to Microsoft SharePoint Online.
	SourceType *string `json:"source_type,omitempty"`

	// Object containing details of the stored credentials.
	//
	// Obtain credentials for your source from the administrator of the source.
	CredentialDetails *CredentialDetails `json:"credential_details,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the UpdateCredentialsOptions.SourceType property.
// The source that this credentials object connects to.
// -  `box` indicates the credentials are used to connect an instance of Enterprise Box.
// -  `salesforce` indicates the credentials are used to connect to Salesforce.
// -  `sharepoint` indicates the credentials are used to connect to Microsoft SharePoint Online.
const (
	UpdateCredentialsOptions_SourceType_Box        = "box"
	UpdateCredentialsOptions_SourceType_Salesforce = "salesforce"
	UpdateCredentialsOptions_SourceType_Sharepoint = "sharepoint"
)

// NewUpdateCredentialsOptions : Instantiate UpdateCredentialsOptions
func (discovery *DiscoveryV1) NewUpdateCredentialsOptions(environmentID string, credentialID string) *UpdateCredentialsOptions {
	return &UpdateCredentialsOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CredentialID:  core.StringPtr(credentialID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *UpdateCredentialsOptions) SetEnvironmentID(environmentID string) *UpdateCredentialsOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetCredentialID : Allow user to set CredentialID
func (options *UpdateCredentialsOptions) SetCredentialID(credentialID string) *UpdateCredentialsOptions {
	options.CredentialID = core.StringPtr(credentialID)
	return options
}

// SetSourceType : Allow user to set SourceType
func (options *UpdateCredentialsOptions) SetSourceType(sourceType string) *UpdateCredentialsOptions {
	options.SourceType = core.StringPtr(sourceType)
	return options
}

// SetCredentialDetails : Allow user to set CredentialDetails
func (options *UpdateCredentialsOptions) SetCredentialDetails(credentialDetails *CredentialDetails) *UpdateCredentialsOptions {
	options.CredentialDetails = credentialDetails
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateCredentialsOptions) SetHeaders(param map[string]string) *UpdateCredentialsOptions {
	options.Headers = param
	return options
}

// UpdateDocumentOptions : The updateDocument options.
type UpdateDocumentOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The ID of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// The ID of the document.
	DocumentID *string `json:"document_id" validate:"required"`

	// The content of the document to ingest. The maximum supported file size is 50 megabytes. Files larger than 50
	// megabytes is rejected.
	File *os.File `json:"file,omitempty"`

	// The filename for file.
	Filename *string `json:"filename,omitempty"`

	// If you're using the Data Crawler to upload your documents, you can test a document against the type of metadata that
	// the Data Crawler might send. The maximum supported metadata file size is 1 MB. Metadata parts larger than 1 MB are
	// rejected.
	// Example:  ``` {
	//   "Creator": "Johnny Appleseed",
	//   "Subject": "Apples"
	// } ```.
	Metadata *string `json:"metadata,omitempty"`

	// The content type of file. Values for this parameter can be obtained from the HttpMediaType class.
	FileContentType *string `json:"file_content_type,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewUpdateDocumentOptions : Instantiate UpdateDocumentOptions
func (discovery *DiscoveryV1) NewUpdateDocumentOptions(environmentID string, collectionID string, documentID string) *UpdateDocumentOptions {
	return &UpdateDocumentOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
		DocumentID:    core.StringPtr(documentID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *UpdateDocumentOptions) SetEnvironmentID(environmentID string) *UpdateDocumentOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
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
func (options *UpdateDocumentOptions) SetFile(file *os.File) *UpdateDocumentOptions {
	options.File = file
	return options
}

// SetFilename : Allow user to set Filename
func (options *UpdateDocumentOptions) SetFilename(filename string) *UpdateDocumentOptions {
	options.Filename = core.StringPtr(filename)
	return options
}

// SetMetadata : Allow user to set Metadata
func (options *UpdateDocumentOptions) SetMetadata(metadata string) *UpdateDocumentOptions {
	options.Metadata = core.StringPtr(metadata)
	return options
}

// SetFileContentType : Allow user to set FileContentType
func (options *UpdateDocumentOptions) SetFileContentType(fileContentType string) *UpdateDocumentOptions {
	options.FileContentType = core.StringPtr(fileContentType)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateDocumentOptions) SetHeaders(param map[string]string) *UpdateDocumentOptions {
	options.Headers = param
	return options
}

// UpdateEnvironmentOptions : The updateEnvironment options.
type UpdateEnvironmentOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// Name that identifies the environment.
	Name *string `json:"name,omitempty"`

	// Description of the environment.
	Description *string `json:"description,omitempty"`

	// Size that the environment should be increased to. Environment size cannot be modified when using a Lite plan.
	// Environment size can only increased and not decreased.
	Size *string `json:"size,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the UpdateEnvironmentOptions.Size property.
// Size that the environment should be increased to. Environment size cannot be modified when using a Lite plan.
// Environment size can only increased and not decreased.
const (
	UpdateEnvironmentOptions_Size_L    = "L"
	UpdateEnvironmentOptions_Size_M    = "M"
	UpdateEnvironmentOptions_Size_ML   = "ML"
	UpdateEnvironmentOptions_Size_MS   = "MS"
	UpdateEnvironmentOptions_Size_S    = "S"
	UpdateEnvironmentOptions_Size_XL   = "XL"
	UpdateEnvironmentOptions_Size_XXL  = "XXL"
	UpdateEnvironmentOptions_Size_XXXL = "XXXL"
)

// NewUpdateEnvironmentOptions : Instantiate UpdateEnvironmentOptions
func (discovery *DiscoveryV1) NewUpdateEnvironmentOptions(environmentID string) *UpdateEnvironmentOptions {
	return &UpdateEnvironmentOptions{
		EnvironmentID: core.StringPtr(environmentID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *UpdateEnvironmentOptions) SetEnvironmentID(environmentID string) *UpdateEnvironmentOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetName : Allow user to set Name
func (options *UpdateEnvironmentOptions) SetName(name string) *UpdateEnvironmentOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetDescription : Allow user to set Description
func (options *UpdateEnvironmentOptions) SetDescription(description string) *UpdateEnvironmentOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetSize : Allow user to set Size
func (options *UpdateEnvironmentOptions) SetSize(size string) *UpdateEnvironmentOptions {
	options.Size = core.StringPtr(size)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateEnvironmentOptions) SetHeaders(param map[string]string) *UpdateEnvironmentOptions {
	options.Headers = param
	return options
}

// UpdateTrainingExampleOptions : The updateTrainingExample options.
type UpdateTrainingExampleOptions struct {

	// The ID of the environment.
	EnvironmentID *string `json:"environment_id" validate:"required"`

	// The ID of the collection.
	CollectionID *string `json:"collection_id" validate:"required"`

	// The ID of the query used for training.
	QueryID *string `json:"query_id" validate:"required"`

	// The ID of the document as it is indexed.
	ExampleID *string `json:"example_id" validate:"required"`

	CrossReference *string `json:"cross_reference,omitempty"`

	Relevance *int64 `json:"relevance,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewUpdateTrainingExampleOptions : Instantiate UpdateTrainingExampleOptions
func (discovery *DiscoveryV1) NewUpdateTrainingExampleOptions(environmentID string, collectionID string, queryID string, exampleID string) *UpdateTrainingExampleOptions {
	return &UpdateTrainingExampleOptions{
		EnvironmentID: core.StringPtr(environmentID),
		CollectionID:  core.StringPtr(collectionID),
		QueryID:       core.StringPtr(queryID),
		ExampleID:     core.StringPtr(exampleID),
	}
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *UpdateTrainingExampleOptions) SetEnvironmentID(environmentID string) *UpdateTrainingExampleOptions {
	options.EnvironmentID = core.StringPtr(environmentID)
	return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *UpdateTrainingExampleOptions) SetCollectionID(collectionID string) *UpdateTrainingExampleOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
}

// SetQueryID : Allow user to set QueryID
func (options *UpdateTrainingExampleOptions) SetQueryID(queryID string) *UpdateTrainingExampleOptions {
	options.QueryID = core.StringPtr(queryID)
	return options
}

// SetExampleID : Allow user to set ExampleID
func (options *UpdateTrainingExampleOptions) SetExampleID(exampleID string) *UpdateTrainingExampleOptions {
	options.ExampleID = core.StringPtr(exampleID)
	return options
}

// SetCrossReference : Allow user to set CrossReference
func (options *UpdateTrainingExampleOptions) SetCrossReference(crossReference string) *UpdateTrainingExampleOptions {
	options.CrossReference = core.StringPtr(crossReference)
	return options
}

// SetRelevance : Allow user to set Relevance
func (options *UpdateTrainingExampleOptions) SetRelevance(relevance int64) *UpdateTrainingExampleOptions {
	options.Relevance = core.Int64Ptr(relevance)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateTrainingExampleOptions) SetHeaders(param map[string]string) *UpdateTrainingExampleOptions {
	options.Headers = param
	return options
}

// WordHeadingDetection : WordHeadingDetection struct
type WordHeadingDetection struct {
	Fonts []FontSetting `json:"fonts,omitempty"`

	Styles []WordStyle `json:"styles,omitempty"`
}

// WordSettings : A list of Word conversion settings.
type WordSettings struct {
	Heading *WordHeadingDetection `json:"heading,omitempty"`
}

// WordStyle : WordStyle struct
type WordStyle struct {
	Level *int64 `json:"level,omitempty"`

	Names []string `json:"names,omitempty"`
}

// XPathPatterns : XPathPatterns struct
type XPathPatterns struct {
	Xpaths []string `json:"xpaths,omitempty"`
}

// Calculation : Calculation struct
type Calculation struct {

	// The field where the aggregation is located in the document.
	Field *string `json:"field,omitempty"`

	// Value of the aggregation.
	Value *float64 `json:"value,omitempty"`
}

// Filter : Filter struct
type Filter struct {

	// The match the aggregated results queried for.
	Match *string `json:"match,omitempty"`
}

// Histogram : Histogram struct
type Histogram struct {

	// The field where the aggregation is located in the document.
	Field *string `json:"field,omitempty"`

	// Interval of the aggregation. (For 'histogram' type).
	Interval *int64 `json:"interval,omitempty"`
}

// Nested : Nested struct
type Nested struct {

	// The area of the results the aggregation was restricted to.
	Path *string `json:"path,omitempty"`
}

// Term : Term struct
type Term struct {

	// The field where the aggregation is located in the document.
	Field *string `json:"field,omitempty"`

	Count *int64 `json:"count,omitempty"`
}

// Timeslice : Timeslice struct
type Timeslice struct {

	// The field where the aggregation is located in the document.
	Field *string `json:"field,omitempty"`

	// Interval of the aggregation. Valid date interval values are second/seconds minute/minutes, hour/hours, day/days,
	// week/weeks, month/months, and year/years.
	Interval *string `json:"interval,omitempty"`

	// Used to indicate that anomaly detection should be performed. Anomaly detection is used to locate unusual datapoints
	// within a time series.
	Anomaly *bool `json:"anomaly,omitempty"`
}

// TopHits : TopHits struct
type TopHits struct {

	// Number of top hits returned by the aggregation.
	Size *int64 `json:"size,omitempty"`

	Hits *TopHitsResults `json:"hits,omitempty"`
}
