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

// Package assistantv1 : Operations and models for the AssistantV1 service
package assistantv1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	common "github.com/watson-developer-cloud/go-sdk/v2/common"
)

// AssistantV1 : The IBM Watson&trade; Assistant service combines machine learning, natural language understanding, and
// an integrated dialog editor to create conversation flows between your apps and your users.
//
// The Assistant v1 API provides authoring methods your application can use to create or update a workspace.
//
// API Version: 1.0
// See: https://cloud.ibm.com/docs/assistant
type AssistantV1 struct {
	Service *core.BaseService

	// Release date of the API version you want to use. Specify dates in YYYY-MM-DD format. The current version is
	// `2021-06-14`.
	Version *string
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.us-south.assistant.watson.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "conversation"

// AssistantV1Options : Service options
type AssistantV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator

	// Release date of the API version you want to use. Specify dates in YYYY-MM-DD format. The current version is
	// `2021-06-14`.
	Version *string `validate:"required"`
}

// NewAssistantV1 : constructs an instance of AssistantV1 with passed in options.
func NewAssistantV1(options *AssistantV1Options) (service *AssistantV1, err error) {
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

	service = &AssistantV1{
		Service: baseService,
		Version: options.Version,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "assistant" suitable for processing requests.
func (assistant *AssistantV1) Clone() *AssistantV1 {
	if core.IsNil(assistant) {
		return nil
	}
	clone := *assistant
	clone.Service = assistant.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (assistant *AssistantV1) SetServiceURL(url string) error {
	return assistant.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (assistant *AssistantV1) GetServiceURL() string {
	return assistant.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (assistant *AssistantV1) SetDefaultHeaders(headers http.Header) {
	assistant.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (assistant *AssistantV1) SetEnableGzipCompression(enableGzip bool) {
	assistant.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (assistant *AssistantV1) GetEnableGzipCompression() bool {
	return assistant.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (assistant *AssistantV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	assistant.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (assistant *AssistantV1) DisableRetries() {
	assistant.Service.DisableRetries()
}

// DisableSSLVerification bypasses verification of the server's SSL certificate
func (assistant *AssistantV1) DisableSSLVerification() {
	assistant.Service.DisableSSLVerification()
}

// Message : Get response to user input
// Send user input to a workspace and receive a response.
//
// **Important:** This method has been superseded by the new v2 runtime API. The v2 API offers significant advantages,
// including ease of deployment, automatic state management, versioning, and search capabilities. For more information,
// see the [documentation](https://cloud.ibm.com/docs/assistant?topic=assistant-api-overview).
func (assistant *AssistantV1) Message(messageOptions *MessageOptions) (result *MessageResponse, response *core.DetailedResponse, err error) {
	return assistant.MessageWithContext(context.Background(), messageOptions)
}

// MessageWithContext is an alternate form of the Message method which supports a Context parameter
func (assistant *AssistantV1) MessageWithContext(ctx context.Context, messageOptions *MessageOptions) (result *MessageResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(messageOptions, "messageOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(messageOptions, "messageOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *messageOptions.WorkspaceID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/message`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range messageOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "Message")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	if messageOptions.NodesVisitedDetails != nil {
		builder.AddQuery("nodes_visited_details", fmt.Sprint(*messageOptions.NodesVisitedDetails))
	}

	body := make(map[string]interface{})
	if messageOptions.Input != nil {
		body["input"] = messageOptions.Input
	}
	if messageOptions.Intents != nil {
		body["intents"] = messageOptions.Intents
	}
	if messageOptions.Entities != nil {
		body["entities"] = messageOptions.Entities
	}
	if messageOptions.AlternateIntents != nil {
		body["alternate_intents"] = messageOptions.AlternateIntents
	}
	if messageOptions.Context != nil {
		body["context"] = messageOptions.Context
	}
	if messageOptions.Output != nil {
		body["output"] = messageOptions.Output
	}
	if messageOptions.UserID != nil {
		body["user_id"] = messageOptions.UserID
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
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMessageResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// BulkClassify : Identify intents and entities in multiple user utterances
// Send multiple user inputs to a workspace in a single request and receive information about the intents and entities
// recognized in each input. This method is useful for testing and comparing the performance of different workspaces.
//
// This method is available only with Enterprise with Data Isolation plans.
func (assistant *AssistantV1) BulkClassify(bulkClassifyOptions *BulkClassifyOptions) (result *BulkClassifyResponse, response *core.DetailedResponse, err error) {
	return assistant.BulkClassifyWithContext(context.Background(), bulkClassifyOptions)
}

// BulkClassifyWithContext is an alternate form of the BulkClassify method which supports a Context parameter
func (assistant *AssistantV1) BulkClassifyWithContext(ctx context.Context, bulkClassifyOptions *BulkClassifyOptions) (result *BulkClassifyResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(bulkClassifyOptions, "bulkClassifyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(bulkClassifyOptions, "bulkClassifyOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *bulkClassifyOptions.WorkspaceID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/bulk_classify`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range bulkClassifyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "BulkClassify")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))

	body := make(map[string]interface{})
	if bulkClassifyOptions.Input != nil {
		body["input"] = bulkClassifyOptions.Input
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
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBulkClassifyResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListWorkspaces : List workspaces
// List the workspaces associated with a Watson Assistant service instance.
func (assistant *AssistantV1) ListWorkspaces(listWorkspacesOptions *ListWorkspacesOptions) (result *WorkspaceCollection, response *core.DetailedResponse, err error) {
	return assistant.ListWorkspacesWithContext(context.Background(), listWorkspacesOptions)
}

// ListWorkspacesWithContext is an alternate form of the ListWorkspaces method which supports a Context parameter
func (assistant *AssistantV1) ListWorkspacesWithContext(ctx context.Context, listWorkspacesOptions *ListWorkspacesOptions) (result *WorkspaceCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listWorkspacesOptions, "listWorkspacesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listWorkspacesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "ListWorkspaces")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	if listWorkspacesOptions.PageLimit != nil {
		builder.AddQuery("page_limit", fmt.Sprint(*listWorkspacesOptions.PageLimit))
	}
	if listWorkspacesOptions.IncludeCount != nil {
		builder.AddQuery("include_count", fmt.Sprint(*listWorkspacesOptions.IncludeCount))
	}
	if listWorkspacesOptions.Sort != nil {
		builder.AddQuery("sort", fmt.Sprint(*listWorkspacesOptions.Sort))
	}
	if listWorkspacesOptions.Cursor != nil {
		builder.AddQuery("cursor", fmt.Sprint(*listWorkspacesOptions.Cursor))
	}
	if listWorkspacesOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*listWorkspacesOptions.IncludeAudit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalWorkspaceCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateWorkspace : Create workspace
// Create a workspace based on component objects. You must provide workspace components defining the content of the new
// workspace.
func (assistant *AssistantV1) CreateWorkspace(createWorkspaceOptions *CreateWorkspaceOptions) (result *Workspace, response *core.DetailedResponse, err error) {
	return assistant.CreateWorkspaceWithContext(context.Background(), createWorkspaceOptions)
}

// CreateWorkspaceWithContext is an alternate form of the CreateWorkspace method which supports a Context parameter
func (assistant *AssistantV1) CreateWorkspaceWithContext(ctx context.Context, createWorkspaceOptions *CreateWorkspaceOptions) (result *Workspace, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(createWorkspaceOptions, "createWorkspaceOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createWorkspaceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "CreateWorkspace")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	if createWorkspaceOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*createWorkspaceOptions.IncludeAudit))
	}

	body := make(map[string]interface{})
	if createWorkspaceOptions.Name != nil {
		body["name"] = createWorkspaceOptions.Name
	}
	if createWorkspaceOptions.Description != nil {
		body["description"] = createWorkspaceOptions.Description
	}
	if createWorkspaceOptions.Language != nil {
		body["language"] = createWorkspaceOptions.Language
	}
	if createWorkspaceOptions.DialogNodes != nil {
		body["dialog_nodes"] = createWorkspaceOptions.DialogNodes
	}
	if createWorkspaceOptions.Counterexamples != nil {
		body["counterexamples"] = createWorkspaceOptions.Counterexamples
	}
	if createWorkspaceOptions.Metadata != nil {
		body["metadata"] = createWorkspaceOptions.Metadata
	}
	if createWorkspaceOptions.LearningOptOut != nil {
		body["learning_opt_out"] = createWorkspaceOptions.LearningOptOut
	}
	if createWorkspaceOptions.SystemSettings != nil {
		body["system_settings"] = createWorkspaceOptions.SystemSettings
	}
	if createWorkspaceOptions.Webhooks != nil {
		body["webhooks"] = createWorkspaceOptions.Webhooks
	}
	if createWorkspaceOptions.Intents != nil {
		body["intents"] = createWorkspaceOptions.Intents
	}
	if createWorkspaceOptions.Entities != nil {
		body["entities"] = createWorkspaceOptions.Entities
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
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalWorkspace)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetWorkspace : Get information about a workspace
// Get information about a workspace, optionally including all workspace content.
func (assistant *AssistantV1) GetWorkspace(getWorkspaceOptions *GetWorkspaceOptions) (result *Workspace, response *core.DetailedResponse, err error) {
	return assistant.GetWorkspaceWithContext(context.Background(), getWorkspaceOptions)
}

// GetWorkspaceWithContext is an alternate form of the GetWorkspace method which supports a Context parameter
func (assistant *AssistantV1) GetWorkspaceWithContext(ctx context.Context, getWorkspaceOptions *GetWorkspaceOptions) (result *Workspace, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getWorkspaceOptions, "getWorkspaceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getWorkspaceOptions, "getWorkspaceOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *getWorkspaceOptions.WorkspaceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getWorkspaceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "GetWorkspace")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	if getWorkspaceOptions.Export != nil {
		builder.AddQuery("export", fmt.Sprint(*getWorkspaceOptions.Export))
	}
	if getWorkspaceOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*getWorkspaceOptions.IncludeAudit))
	}
	if getWorkspaceOptions.Sort != nil {
		builder.AddQuery("sort", fmt.Sprint(*getWorkspaceOptions.Sort))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalWorkspace)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateWorkspace : Update workspace
// Update an existing workspace with new or modified data. You must provide component objects defining the content of
// the updated workspace.
func (assistant *AssistantV1) UpdateWorkspace(updateWorkspaceOptions *UpdateWorkspaceOptions) (result *Workspace, response *core.DetailedResponse, err error) {
	return assistant.UpdateWorkspaceWithContext(context.Background(), updateWorkspaceOptions)
}

// UpdateWorkspaceWithContext is an alternate form of the UpdateWorkspace method which supports a Context parameter
func (assistant *AssistantV1) UpdateWorkspaceWithContext(ctx context.Context, updateWorkspaceOptions *UpdateWorkspaceOptions) (result *Workspace, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateWorkspaceOptions, "updateWorkspaceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateWorkspaceOptions, "updateWorkspaceOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *updateWorkspaceOptions.WorkspaceID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateWorkspaceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "UpdateWorkspace")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	if updateWorkspaceOptions.Append != nil {
		builder.AddQuery("append", fmt.Sprint(*updateWorkspaceOptions.Append))
	}
	if updateWorkspaceOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*updateWorkspaceOptions.IncludeAudit))
	}

	body := make(map[string]interface{})
	if updateWorkspaceOptions.Name != nil {
		body["name"] = updateWorkspaceOptions.Name
	}
	if updateWorkspaceOptions.Description != nil {
		body["description"] = updateWorkspaceOptions.Description
	}
	if updateWorkspaceOptions.Language != nil {
		body["language"] = updateWorkspaceOptions.Language
	}
	if updateWorkspaceOptions.DialogNodes != nil {
		body["dialog_nodes"] = updateWorkspaceOptions.DialogNodes
	}
	if updateWorkspaceOptions.Counterexamples != nil {
		body["counterexamples"] = updateWorkspaceOptions.Counterexamples
	}
	if updateWorkspaceOptions.Metadata != nil {
		body["metadata"] = updateWorkspaceOptions.Metadata
	}
	if updateWorkspaceOptions.LearningOptOut != nil {
		body["learning_opt_out"] = updateWorkspaceOptions.LearningOptOut
	}
	if updateWorkspaceOptions.SystemSettings != nil {
		body["system_settings"] = updateWorkspaceOptions.SystemSettings
	}
	if updateWorkspaceOptions.Webhooks != nil {
		body["webhooks"] = updateWorkspaceOptions.Webhooks
	}
	if updateWorkspaceOptions.Intents != nil {
		body["intents"] = updateWorkspaceOptions.Intents
	}
	if updateWorkspaceOptions.Entities != nil {
		body["entities"] = updateWorkspaceOptions.Entities
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
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalWorkspace)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteWorkspace : Delete workspace
// Delete a workspace from the service instance.
func (assistant *AssistantV1) DeleteWorkspace(deleteWorkspaceOptions *DeleteWorkspaceOptions) (response *core.DetailedResponse, err error) {
	return assistant.DeleteWorkspaceWithContext(context.Background(), deleteWorkspaceOptions)
}

// DeleteWorkspaceWithContext is an alternate form of the DeleteWorkspace method which supports a Context parameter
func (assistant *AssistantV1) DeleteWorkspaceWithContext(ctx context.Context, deleteWorkspaceOptions *DeleteWorkspaceOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteWorkspaceOptions, "deleteWorkspaceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteWorkspaceOptions, "deleteWorkspaceOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *deleteWorkspaceOptions.WorkspaceID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteWorkspaceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "DeleteWorkspace")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, nil)

	return
}

// ListIntents : List intents
// List the intents for a workspace.
func (assistant *AssistantV1) ListIntents(listIntentsOptions *ListIntentsOptions) (result *IntentCollection, response *core.DetailedResponse, err error) {
	return assistant.ListIntentsWithContext(context.Background(), listIntentsOptions)
}

// ListIntentsWithContext is an alternate form of the ListIntents method which supports a Context parameter
func (assistant *AssistantV1) ListIntentsWithContext(ctx context.Context, listIntentsOptions *ListIntentsOptions) (result *IntentCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listIntentsOptions, "listIntentsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listIntentsOptions, "listIntentsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *listIntentsOptions.WorkspaceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/intents`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listIntentsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "ListIntents")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	if listIntentsOptions.Export != nil {
		builder.AddQuery("export", fmt.Sprint(*listIntentsOptions.Export))
	}
	if listIntentsOptions.PageLimit != nil {
		builder.AddQuery("page_limit", fmt.Sprint(*listIntentsOptions.PageLimit))
	}
	if listIntentsOptions.IncludeCount != nil {
		builder.AddQuery("include_count", fmt.Sprint(*listIntentsOptions.IncludeCount))
	}
	if listIntentsOptions.Sort != nil {
		builder.AddQuery("sort", fmt.Sprint(*listIntentsOptions.Sort))
	}
	if listIntentsOptions.Cursor != nil {
		builder.AddQuery("cursor", fmt.Sprint(*listIntentsOptions.Cursor))
	}
	if listIntentsOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*listIntentsOptions.IncludeAudit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalIntentCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateIntent : Create intent
// Create a new intent.
//
// If you want to create multiple intents with a single API call, consider using the **[Update
// workspace](#update-workspace)** method instead.
func (assistant *AssistantV1) CreateIntent(createIntentOptions *CreateIntentOptions) (result *Intent, response *core.DetailedResponse, err error) {
	return assistant.CreateIntentWithContext(context.Background(), createIntentOptions)
}

// CreateIntentWithContext is an alternate form of the CreateIntent method which supports a Context parameter
func (assistant *AssistantV1) CreateIntentWithContext(ctx context.Context, createIntentOptions *CreateIntentOptions) (result *Intent, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createIntentOptions, "createIntentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createIntentOptions, "createIntentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *createIntentOptions.WorkspaceID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/intents`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createIntentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "CreateIntent")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	if createIntentOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*createIntentOptions.IncludeAudit))
	}

	body := make(map[string]interface{})
	if createIntentOptions.Intent != nil {
		body["intent"] = createIntentOptions.Intent
	}
	if createIntentOptions.Description != nil {
		body["description"] = createIntentOptions.Description
	}
	if createIntentOptions.Examples != nil {
		body["examples"] = createIntentOptions.Examples
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
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalIntent)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetIntent : Get intent
// Get information about an intent, optionally including all intent content.
func (assistant *AssistantV1) GetIntent(getIntentOptions *GetIntentOptions) (result *Intent, response *core.DetailedResponse, err error) {
	return assistant.GetIntentWithContext(context.Background(), getIntentOptions)
}

// GetIntentWithContext is an alternate form of the GetIntent method which supports a Context parameter
func (assistant *AssistantV1) GetIntentWithContext(ctx context.Context, getIntentOptions *GetIntentOptions) (result *Intent, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getIntentOptions, "getIntentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getIntentOptions, "getIntentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *getIntentOptions.WorkspaceID,
		"intent":       *getIntentOptions.Intent,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/intents/{intent}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getIntentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "GetIntent")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	if getIntentOptions.Export != nil {
		builder.AddQuery("export", fmt.Sprint(*getIntentOptions.Export))
	}
	if getIntentOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*getIntentOptions.IncludeAudit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalIntent)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateIntent : Update intent
// Update an existing intent with new or modified data. You must provide component objects defining the content of the
// updated intent.
//
// If you want to update multiple intents with a single API call, consider using the **[Update
// workspace](#update-workspace)** method instead.
func (assistant *AssistantV1) UpdateIntent(updateIntentOptions *UpdateIntentOptions) (result *Intent, response *core.DetailedResponse, err error) {
	return assistant.UpdateIntentWithContext(context.Background(), updateIntentOptions)
}

// UpdateIntentWithContext is an alternate form of the UpdateIntent method which supports a Context parameter
func (assistant *AssistantV1) UpdateIntentWithContext(ctx context.Context, updateIntentOptions *UpdateIntentOptions) (result *Intent, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateIntentOptions, "updateIntentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateIntentOptions, "updateIntentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *updateIntentOptions.WorkspaceID,
		"intent":       *updateIntentOptions.Intent,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/intents/{intent}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateIntentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "UpdateIntent")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	if updateIntentOptions.Append != nil {
		builder.AddQuery("append", fmt.Sprint(*updateIntentOptions.Append))
	}
	if updateIntentOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*updateIntentOptions.IncludeAudit))
	}

	body := make(map[string]interface{})
	if updateIntentOptions.NewIntent != nil {
		body["intent"] = updateIntentOptions.NewIntent
	}
	if updateIntentOptions.NewDescription != nil {
		body["description"] = updateIntentOptions.NewDescription
	}
	if updateIntentOptions.NewExamples != nil {
		body["examples"] = updateIntentOptions.NewExamples
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
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalIntent)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteIntent : Delete intent
// Delete an intent from a workspace.
func (assistant *AssistantV1) DeleteIntent(deleteIntentOptions *DeleteIntentOptions) (response *core.DetailedResponse, err error) {
	return assistant.DeleteIntentWithContext(context.Background(), deleteIntentOptions)
}

// DeleteIntentWithContext is an alternate form of the DeleteIntent method which supports a Context parameter
func (assistant *AssistantV1) DeleteIntentWithContext(ctx context.Context, deleteIntentOptions *DeleteIntentOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteIntentOptions, "deleteIntentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteIntentOptions, "deleteIntentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *deleteIntentOptions.WorkspaceID,
		"intent":       *deleteIntentOptions.Intent,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/intents/{intent}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteIntentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "DeleteIntent")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, nil)

	return
}

// ListExamples : List user input examples
// List the user input examples for an intent, optionally including contextual entity mentions.
func (assistant *AssistantV1) ListExamples(listExamplesOptions *ListExamplesOptions) (result *ExampleCollection, response *core.DetailedResponse, err error) {
	return assistant.ListExamplesWithContext(context.Background(), listExamplesOptions)
}

// ListExamplesWithContext is an alternate form of the ListExamples method which supports a Context parameter
func (assistant *AssistantV1) ListExamplesWithContext(ctx context.Context, listExamplesOptions *ListExamplesOptions) (result *ExampleCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listExamplesOptions, "listExamplesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listExamplesOptions, "listExamplesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *listExamplesOptions.WorkspaceID,
		"intent":       *listExamplesOptions.Intent,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/intents/{intent}/examples`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listExamplesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "ListExamples")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	if listExamplesOptions.PageLimit != nil {
		builder.AddQuery("page_limit", fmt.Sprint(*listExamplesOptions.PageLimit))
	}
	if listExamplesOptions.IncludeCount != nil {
		builder.AddQuery("include_count", fmt.Sprint(*listExamplesOptions.IncludeCount))
	}
	if listExamplesOptions.Sort != nil {
		builder.AddQuery("sort", fmt.Sprint(*listExamplesOptions.Sort))
	}
	if listExamplesOptions.Cursor != nil {
		builder.AddQuery("cursor", fmt.Sprint(*listExamplesOptions.Cursor))
	}
	if listExamplesOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*listExamplesOptions.IncludeAudit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalExampleCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateExample : Create user input example
// Add a new user input example to an intent.
//
// If you want to add multiple examples with a single API call, consider using the **[Update intent](#update-intent)**
// method instead.
func (assistant *AssistantV1) CreateExample(createExampleOptions *CreateExampleOptions) (result *Example, response *core.DetailedResponse, err error) {
	return assistant.CreateExampleWithContext(context.Background(), createExampleOptions)
}

// CreateExampleWithContext is an alternate form of the CreateExample method which supports a Context parameter
func (assistant *AssistantV1) CreateExampleWithContext(ctx context.Context, createExampleOptions *CreateExampleOptions) (result *Example, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createExampleOptions, "createExampleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createExampleOptions, "createExampleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *createExampleOptions.WorkspaceID,
		"intent":       *createExampleOptions.Intent,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/intents/{intent}/examples`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createExampleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "CreateExample")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	if createExampleOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*createExampleOptions.IncludeAudit))
	}

	body := make(map[string]interface{})
	if createExampleOptions.Text != nil {
		body["text"] = createExampleOptions.Text
	}
	if createExampleOptions.Mentions != nil {
		body["mentions"] = createExampleOptions.Mentions
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
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalExample)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetExample : Get user input example
// Get information about a user input example.
func (assistant *AssistantV1) GetExample(getExampleOptions *GetExampleOptions) (result *Example, response *core.DetailedResponse, err error) {
	return assistant.GetExampleWithContext(context.Background(), getExampleOptions)
}

// GetExampleWithContext is an alternate form of the GetExample method which supports a Context parameter
func (assistant *AssistantV1) GetExampleWithContext(ctx context.Context, getExampleOptions *GetExampleOptions) (result *Example, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getExampleOptions, "getExampleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getExampleOptions, "getExampleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *getExampleOptions.WorkspaceID,
		"intent":       *getExampleOptions.Intent,
		"text":         *getExampleOptions.Text,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/intents/{intent}/examples/{text}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getExampleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "GetExample")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	if getExampleOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*getExampleOptions.IncludeAudit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalExample)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateExample : Update user input example
// Update the text of a user input example.
//
// If you want to update multiple examples with a single API call, consider using the **[Update
// intent](#update-intent)** method instead.
func (assistant *AssistantV1) UpdateExample(updateExampleOptions *UpdateExampleOptions) (result *Example, response *core.DetailedResponse, err error) {
	return assistant.UpdateExampleWithContext(context.Background(), updateExampleOptions)
}

// UpdateExampleWithContext is an alternate form of the UpdateExample method which supports a Context parameter
func (assistant *AssistantV1) UpdateExampleWithContext(ctx context.Context, updateExampleOptions *UpdateExampleOptions) (result *Example, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateExampleOptions, "updateExampleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateExampleOptions, "updateExampleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *updateExampleOptions.WorkspaceID,
		"intent":       *updateExampleOptions.Intent,
		"text":         *updateExampleOptions.Text,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/intents/{intent}/examples/{text}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateExampleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "UpdateExample")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	if updateExampleOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*updateExampleOptions.IncludeAudit))
	}

	body := make(map[string]interface{})
	if updateExampleOptions.NewText != nil {
		body["text"] = updateExampleOptions.NewText
	}
	if updateExampleOptions.NewMentions != nil {
		body["mentions"] = updateExampleOptions.NewMentions
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
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalExample)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteExample : Delete user input example
// Delete a user input example from an intent.
func (assistant *AssistantV1) DeleteExample(deleteExampleOptions *DeleteExampleOptions) (response *core.DetailedResponse, err error) {
	return assistant.DeleteExampleWithContext(context.Background(), deleteExampleOptions)
}

// DeleteExampleWithContext is an alternate form of the DeleteExample method which supports a Context parameter
func (assistant *AssistantV1) DeleteExampleWithContext(ctx context.Context, deleteExampleOptions *DeleteExampleOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteExampleOptions, "deleteExampleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteExampleOptions, "deleteExampleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *deleteExampleOptions.WorkspaceID,
		"intent":       *deleteExampleOptions.Intent,
		"text":         *deleteExampleOptions.Text,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/intents/{intent}/examples/{text}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteExampleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "DeleteExample")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, nil)

	return
}

// ListCounterexamples : List counterexamples
// List the counterexamples for a workspace. Counterexamples are examples that have been marked as irrelevant input.
func (assistant *AssistantV1) ListCounterexamples(listCounterexamplesOptions *ListCounterexamplesOptions) (result *CounterexampleCollection, response *core.DetailedResponse, err error) {
	return assistant.ListCounterexamplesWithContext(context.Background(), listCounterexamplesOptions)
}

// ListCounterexamplesWithContext is an alternate form of the ListCounterexamples method which supports a Context parameter
func (assistant *AssistantV1) ListCounterexamplesWithContext(ctx context.Context, listCounterexamplesOptions *ListCounterexamplesOptions) (result *CounterexampleCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listCounterexamplesOptions, "listCounterexamplesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listCounterexamplesOptions, "listCounterexamplesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *listCounterexamplesOptions.WorkspaceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/counterexamples`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listCounterexamplesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "ListCounterexamples")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	if listCounterexamplesOptions.PageLimit != nil {
		builder.AddQuery("page_limit", fmt.Sprint(*listCounterexamplesOptions.PageLimit))
	}
	if listCounterexamplesOptions.IncludeCount != nil {
		builder.AddQuery("include_count", fmt.Sprint(*listCounterexamplesOptions.IncludeCount))
	}
	if listCounterexamplesOptions.Sort != nil {
		builder.AddQuery("sort", fmt.Sprint(*listCounterexamplesOptions.Sort))
	}
	if listCounterexamplesOptions.Cursor != nil {
		builder.AddQuery("cursor", fmt.Sprint(*listCounterexamplesOptions.Cursor))
	}
	if listCounterexamplesOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*listCounterexamplesOptions.IncludeAudit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCounterexampleCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateCounterexample : Create counterexample
// Add a new counterexample to a workspace. Counterexamples are examples that have been marked as irrelevant input.
//
// If you want to add multiple counterexamples with a single API call, consider using the **[Update
// workspace](#update-workspace)** method instead.
func (assistant *AssistantV1) CreateCounterexample(createCounterexampleOptions *CreateCounterexampleOptions) (result *Counterexample, response *core.DetailedResponse, err error) {
	return assistant.CreateCounterexampleWithContext(context.Background(), createCounterexampleOptions)
}

// CreateCounterexampleWithContext is an alternate form of the CreateCounterexample method which supports a Context parameter
func (assistant *AssistantV1) CreateCounterexampleWithContext(ctx context.Context, createCounterexampleOptions *CreateCounterexampleOptions) (result *Counterexample, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createCounterexampleOptions, "createCounterexampleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createCounterexampleOptions, "createCounterexampleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *createCounterexampleOptions.WorkspaceID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/counterexamples`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createCounterexampleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "CreateCounterexample")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	if createCounterexampleOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*createCounterexampleOptions.IncludeAudit))
	}

	body := make(map[string]interface{})
	if createCounterexampleOptions.Text != nil {
		body["text"] = createCounterexampleOptions.Text
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
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCounterexample)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetCounterexample : Get counterexample
// Get information about a counterexample. Counterexamples are examples that have been marked as irrelevant input.
func (assistant *AssistantV1) GetCounterexample(getCounterexampleOptions *GetCounterexampleOptions) (result *Counterexample, response *core.DetailedResponse, err error) {
	return assistant.GetCounterexampleWithContext(context.Background(), getCounterexampleOptions)
}

// GetCounterexampleWithContext is an alternate form of the GetCounterexample method which supports a Context parameter
func (assistant *AssistantV1) GetCounterexampleWithContext(ctx context.Context, getCounterexampleOptions *GetCounterexampleOptions) (result *Counterexample, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getCounterexampleOptions, "getCounterexampleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getCounterexampleOptions, "getCounterexampleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *getCounterexampleOptions.WorkspaceID,
		"text":         *getCounterexampleOptions.Text,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/counterexamples/{text}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCounterexampleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "GetCounterexample")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	if getCounterexampleOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*getCounterexampleOptions.IncludeAudit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCounterexample)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateCounterexample : Update counterexample
// Update the text of a counterexample. Counterexamples are examples that have been marked as irrelevant input.
func (assistant *AssistantV1) UpdateCounterexample(updateCounterexampleOptions *UpdateCounterexampleOptions) (result *Counterexample, response *core.DetailedResponse, err error) {
	return assistant.UpdateCounterexampleWithContext(context.Background(), updateCounterexampleOptions)
}

// UpdateCounterexampleWithContext is an alternate form of the UpdateCounterexample method which supports a Context parameter
func (assistant *AssistantV1) UpdateCounterexampleWithContext(ctx context.Context, updateCounterexampleOptions *UpdateCounterexampleOptions) (result *Counterexample, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateCounterexampleOptions, "updateCounterexampleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateCounterexampleOptions, "updateCounterexampleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *updateCounterexampleOptions.WorkspaceID,
		"text":         *updateCounterexampleOptions.Text,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/counterexamples/{text}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateCounterexampleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "UpdateCounterexample")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	if updateCounterexampleOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*updateCounterexampleOptions.IncludeAudit))
	}

	body := make(map[string]interface{})
	if updateCounterexampleOptions.NewText != nil {
		body["text"] = updateCounterexampleOptions.NewText
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
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCounterexample)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteCounterexample : Delete counterexample
// Delete a counterexample from a workspace. Counterexamples are examples that have been marked as irrelevant input.
func (assistant *AssistantV1) DeleteCounterexample(deleteCounterexampleOptions *DeleteCounterexampleOptions) (response *core.DetailedResponse, err error) {
	return assistant.DeleteCounterexampleWithContext(context.Background(), deleteCounterexampleOptions)
}

// DeleteCounterexampleWithContext is an alternate form of the DeleteCounterexample method which supports a Context parameter
func (assistant *AssistantV1) DeleteCounterexampleWithContext(ctx context.Context, deleteCounterexampleOptions *DeleteCounterexampleOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteCounterexampleOptions, "deleteCounterexampleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteCounterexampleOptions, "deleteCounterexampleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *deleteCounterexampleOptions.WorkspaceID,
		"text":         *deleteCounterexampleOptions.Text,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/counterexamples/{text}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteCounterexampleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "DeleteCounterexample")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, nil)

	return
}

// ListEntities : List entities
// List the entities for a workspace.
func (assistant *AssistantV1) ListEntities(listEntitiesOptions *ListEntitiesOptions) (result *EntityCollection, response *core.DetailedResponse, err error) {
	return assistant.ListEntitiesWithContext(context.Background(), listEntitiesOptions)
}

// ListEntitiesWithContext is an alternate form of the ListEntities method which supports a Context parameter
func (assistant *AssistantV1) ListEntitiesWithContext(ctx context.Context, listEntitiesOptions *ListEntitiesOptions) (result *EntityCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listEntitiesOptions, "listEntitiesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listEntitiesOptions, "listEntitiesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *listEntitiesOptions.WorkspaceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/entities`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listEntitiesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "ListEntities")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	if listEntitiesOptions.Export != nil {
		builder.AddQuery("export", fmt.Sprint(*listEntitiesOptions.Export))
	}
	if listEntitiesOptions.PageLimit != nil {
		builder.AddQuery("page_limit", fmt.Sprint(*listEntitiesOptions.PageLimit))
	}
	if listEntitiesOptions.IncludeCount != nil {
		builder.AddQuery("include_count", fmt.Sprint(*listEntitiesOptions.IncludeCount))
	}
	if listEntitiesOptions.Sort != nil {
		builder.AddQuery("sort", fmt.Sprint(*listEntitiesOptions.Sort))
	}
	if listEntitiesOptions.Cursor != nil {
		builder.AddQuery("cursor", fmt.Sprint(*listEntitiesOptions.Cursor))
	}
	if listEntitiesOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*listEntitiesOptions.IncludeAudit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalEntityCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateEntity : Create entity
// Create a new entity, or enable a system entity.
//
// If you want to create multiple entities with a single API call, consider using the **[Update
// workspace](#update-workspace)** method instead.
func (assistant *AssistantV1) CreateEntity(createEntityOptions *CreateEntityOptions) (result *Entity, response *core.DetailedResponse, err error) {
	return assistant.CreateEntityWithContext(context.Background(), createEntityOptions)
}

// CreateEntityWithContext is an alternate form of the CreateEntity method which supports a Context parameter
func (assistant *AssistantV1) CreateEntityWithContext(ctx context.Context, createEntityOptions *CreateEntityOptions) (result *Entity, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createEntityOptions, "createEntityOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createEntityOptions, "createEntityOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *createEntityOptions.WorkspaceID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/entities`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createEntityOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "CreateEntity")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	if createEntityOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*createEntityOptions.IncludeAudit))
	}

	body := make(map[string]interface{})
	if createEntityOptions.Entity != nil {
		body["entity"] = createEntityOptions.Entity
	}
	if createEntityOptions.Description != nil {
		body["description"] = createEntityOptions.Description
	}
	if createEntityOptions.Metadata != nil {
		body["metadata"] = createEntityOptions.Metadata
	}
	if createEntityOptions.FuzzyMatch != nil {
		body["fuzzy_match"] = createEntityOptions.FuzzyMatch
	}
	if createEntityOptions.Values != nil {
		body["values"] = createEntityOptions.Values
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
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalEntity)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetEntity : Get entity
// Get information about an entity, optionally including all entity content.
func (assistant *AssistantV1) GetEntity(getEntityOptions *GetEntityOptions) (result *Entity, response *core.DetailedResponse, err error) {
	return assistant.GetEntityWithContext(context.Background(), getEntityOptions)
}

// GetEntityWithContext is an alternate form of the GetEntity method which supports a Context parameter
func (assistant *AssistantV1) GetEntityWithContext(ctx context.Context, getEntityOptions *GetEntityOptions) (result *Entity, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getEntityOptions, "getEntityOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getEntityOptions, "getEntityOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *getEntityOptions.WorkspaceID,
		"entity":       *getEntityOptions.Entity,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/entities/{entity}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getEntityOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "GetEntity")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	if getEntityOptions.Export != nil {
		builder.AddQuery("export", fmt.Sprint(*getEntityOptions.Export))
	}
	if getEntityOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*getEntityOptions.IncludeAudit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalEntity)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateEntity : Update entity
// Update an existing entity with new or modified data. You must provide component objects defining the content of the
// updated entity.
//
// If you want to update multiple entities with a single API call, consider using the **[Update
// workspace](#update-workspace)** method instead.
func (assistant *AssistantV1) UpdateEntity(updateEntityOptions *UpdateEntityOptions) (result *Entity, response *core.DetailedResponse, err error) {
	return assistant.UpdateEntityWithContext(context.Background(), updateEntityOptions)
}

// UpdateEntityWithContext is an alternate form of the UpdateEntity method which supports a Context parameter
func (assistant *AssistantV1) UpdateEntityWithContext(ctx context.Context, updateEntityOptions *UpdateEntityOptions) (result *Entity, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateEntityOptions, "updateEntityOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateEntityOptions, "updateEntityOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *updateEntityOptions.WorkspaceID,
		"entity":       *updateEntityOptions.Entity,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/entities/{entity}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateEntityOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "UpdateEntity")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	if updateEntityOptions.Append != nil {
		builder.AddQuery("append", fmt.Sprint(*updateEntityOptions.Append))
	}
	if updateEntityOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*updateEntityOptions.IncludeAudit))
	}

	body := make(map[string]interface{})
	if updateEntityOptions.NewEntity != nil {
		body["entity"] = updateEntityOptions.NewEntity
	}
	if updateEntityOptions.NewDescription != nil {
		body["description"] = updateEntityOptions.NewDescription
	}
	if updateEntityOptions.NewMetadata != nil {
		body["metadata"] = updateEntityOptions.NewMetadata
	}
	if updateEntityOptions.NewFuzzyMatch != nil {
		body["fuzzy_match"] = updateEntityOptions.NewFuzzyMatch
	}
	if updateEntityOptions.NewValues != nil {
		body["values"] = updateEntityOptions.NewValues
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
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalEntity)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteEntity : Delete entity
// Delete an entity from a workspace, or disable a system entity.
func (assistant *AssistantV1) DeleteEntity(deleteEntityOptions *DeleteEntityOptions) (response *core.DetailedResponse, err error) {
	return assistant.DeleteEntityWithContext(context.Background(), deleteEntityOptions)
}

// DeleteEntityWithContext is an alternate form of the DeleteEntity method which supports a Context parameter
func (assistant *AssistantV1) DeleteEntityWithContext(ctx context.Context, deleteEntityOptions *DeleteEntityOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteEntityOptions, "deleteEntityOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteEntityOptions, "deleteEntityOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *deleteEntityOptions.WorkspaceID,
		"entity":       *deleteEntityOptions.Entity,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/entities/{entity}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteEntityOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "DeleteEntity")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, nil)

	return
}

// ListMentions : List entity mentions
// List mentions for a contextual entity. An entity mention is an occurrence of a contextual entity in the context of an
// intent user input example.
func (assistant *AssistantV1) ListMentions(listMentionsOptions *ListMentionsOptions) (result *EntityMentionCollection, response *core.DetailedResponse, err error) {
	return assistant.ListMentionsWithContext(context.Background(), listMentionsOptions)
}

// ListMentionsWithContext is an alternate form of the ListMentions method which supports a Context parameter
func (assistant *AssistantV1) ListMentionsWithContext(ctx context.Context, listMentionsOptions *ListMentionsOptions) (result *EntityMentionCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listMentionsOptions, "listMentionsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listMentionsOptions, "listMentionsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *listMentionsOptions.WorkspaceID,
		"entity":       *listMentionsOptions.Entity,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/entities/{entity}/mentions`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listMentionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "ListMentions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	if listMentionsOptions.Export != nil {
		builder.AddQuery("export", fmt.Sprint(*listMentionsOptions.Export))
	}
	if listMentionsOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*listMentionsOptions.IncludeAudit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalEntityMentionCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListValues : List entity values
// List the values for an entity.
func (assistant *AssistantV1) ListValues(listValuesOptions *ListValuesOptions) (result *ValueCollection, response *core.DetailedResponse, err error) {
	return assistant.ListValuesWithContext(context.Background(), listValuesOptions)
}

// ListValuesWithContext is an alternate form of the ListValues method which supports a Context parameter
func (assistant *AssistantV1) ListValuesWithContext(ctx context.Context, listValuesOptions *ListValuesOptions) (result *ValueCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listValuesOptions, "listValuesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listValuesOptions, "listValuesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *listValuesOptions.WorkspaceID,
		"entity":       *listValuesOptions.Entity,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/entities/{entity}/values`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listValuesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "ListValues")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	if listValuesOptions.Export != nil {
		builder.AddQuery("export", fmt.Sprint(*listValuesOptions.Export))
	}
	if listValuesOptions.PageLimit != nil {
		builder.AddQuery("page_limit", fmt.Sprint(*listValuesOptions.PageLimit))
	}
	if listValuesOptions.IncludeCount != nil {
		builder.AddQuery("include_count", fmt.Sprint(*listValuesOptions.IncludeCount))
	}
	if listValuesOptions.Sort != nil {
		builder.AddQuery("sort", fmt.Sprint(*listValuesOptions.Sort))
	}
	if listValuesOptions.Cursor != nil {
		builder.AddQuery("cursor", fmt.Sprint(*listValuesOptions.Cursor))
	}
	if listValuesOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*listValuesOptions.IncludeAudit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalValueCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateValue : Create entity value
// Create a new value for an entity.
//
// If you want to create multiple entity values with a single API call, consider using the **[Update
// entity](#update-entity)** method instead.
func (assistant *AssistantV1) CreateValue(createValueOptions *CreateValueOptions) (result *Value, response *core.DetailedResponse, err error) {
	return assistant.CreateValueWithContext(context.Background(), createValueOptions)
}

// CreateValueWithContext is an alternate form of the CreateValue method which supports a Context parameter
func (assistant *AssistantV1) CreateValueWithContext(ctx context.Context, createValueOptions *CreateValueOptions) (result *Value, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createValueOptions, "createValueOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createValueOptions, "createValueOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *createValueOptions.WorkspaceID,
		"entity":       *createValueOptions.Entity,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/entities/{entity}/values`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createValueOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "CreateValue")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	if createValueOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*createValueOptions.IncludeAudit))
	}

	body := make(map[string]interface{})
	if createValueOptions.Value != nil {
		body["value"] = createValueOptions.Value
	}
	if createValueOptions.Metadata != nil {
		body["metadata"] = createValueOptions.Metadata
	}
	if createValueOptions.Type != nil {
		body["type"] = createValueOptions.Type
	}
	if createValueOptions.Synonyms != nil {
		body["synonyms"] = createValueOptions.Synonyms
	}
	if createValueOptions.Patterns != nil {
		body["patterns"] = createValueOptions.Patterns
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
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalValue)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetValue : Get entity value
// Get information about an entity value.
func (assistant *AssistantV1) GetValue(getValueOptions *GetValueOptions) (result *Value, response *core.DetailedResponse, err error) {
	return assistant.GetValueWithContext(context.Background(), getValueOptions)
}

// GetValueWithContext is an alternate form of the GetValue method which supports a Context parameter
func (assistant *AssistantV1) GetValueWithContext(ctx context.Context, getValueOptions *GetValueOptions) (result *Value, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getValueOptions, "getValueOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getValueOptions, "getValueOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *getValueOptions.WorkspaceID,
		"entity":       *getValueOptions.Entity,
		"value":        *getValueOptions.Value,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getValueOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "GetValue")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	if getValueOptions.Export != nil {
		builder.AddQuery("export", fmt.Sprint(*getValueOptions.Export))
	}
	if getValueOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*getValueOptions.IncludeAudit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalValue)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateValue : Update entity value
// Update an existing entity value with new or modified data. You must provide component objects defining the content of
// the updated entity value.
//
// If you want to update multiple entity values with a single API call, consider using the **[Update
// entity](#update-entity)** method instead.
func (assistant *AssistantV1) UpdateValue(updateValueOptions *UpdateValueOptions) (result *Value, response *core.DetailedResponse, err error) {
	return assistant.UpdateValueWithContext(context.Background(), updateValueOptions)
}

// UpdateValueWithContext is an alternate form of the UpdateValue method which supports a Context parameter
func (assistant *AssistantV1) UpdateValueWithContext(ctx context.Context, updateValueOptions *UpdateValueOptions) (result *Value, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateValueOptions, "updateValueOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateValueOptions, "updateValueOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *updateValueOptions.WorkspaceID,
		"entity":       *updateValueOptions.Entity,
		"value":        *updateValueOptions.Value,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateValueOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "UpdateValue")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	if updateValueOptions.Append != nil {
		builder.AddQuery("append", fmt.Sprint(*updateValueOptions.Append))
	}
	if updateValueOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*updateValueOptions.IncludeAudit))
	}

	body := make(map[string]interface{})
	if updateValueOptions.NewValue != nil {
		body["value"] = updateValueOptions.NewValue
	}
	if updateValueOptions.NewMetadata != nil {
		body["metadata"] = updateValueOptions.NewMetadata
	}
	if updateValueOptions.NewType != nil {
		body["type"] = updateValueOptions.NewType
	}
	if updateValueOptions.NewSynonyms != nil {
		body["synonyms"] = updateValueOptions.NewSynonyms
	}
	if updateValueOptions.NewPatterns != nil {
		body["patterns"] = updateValueOptions.NewPatterns
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
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalValue)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteValue : Delete entity value
// Delete a value from an entity.
func (assistant *AssistantV1) DeleteValue(deleteValueOptions *DeleteValueOptions) (response *core.DetailedResponse, err error) {
	return assistant.DeleteValueWithContext(context.Background(), deleteValueOptions)
}

// DeleteValueWithContext is an alternate form of the DeleteValue method which supports a Context parameter
func (assistant *AssistantV1) DeleteValueWithContext(ctx context.Context, deleteValueOptions *DeleteValueOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteValueOptions, "deleteValueOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteValueOptions, "deleteValueOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *deleteValueOptions.WorkspaceID,
		"entity":       *deleteValueOptions.Entity,
		"value":        *deleteValueOptions.Value,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteValueOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "DeleteValue")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, nil)

	return
}

// ListSynonyms : List entity value synonyms
// List the synonyms for an entity value.
func (assistant *AssistantV1) ListSynonyms(listSynonymsOptions *ListSynonymsOptions) (result *SynonymCollection, response *core.DetailedResponse, err error) {
	return assistant.ListSynonymsWithContext(context.Background(), listSynonymsOptions)
}

// ListSynonymsWithContext is an alternate form of the ListSynonyms method which supports a Context parameter
func (assistant *AssistantV1) ListSynonymsWithContext(ctx context.Context, listSynonymsOptions *ListSynonymsOptions) (result *SynonymCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listSynonymsOptions, "listSynonymsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listSynonymsOptions, "listSynonymsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *listSynonymsOptions.WorkspaceID,
		"entity":       *listSynonymsOptions.Entity,
		"value":        *listSynonymsOptions.Value,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}/synonyms`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listSynonymsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "ListSynonyms")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	if listSynonymsOptions.PageLimit != nil {
		builder.AddQuery("page_limit", fmt.Sprint(*listSynonymsOptions.PageLimit))
	}
	if listSynonymsOptions.IncludeCount != nil {
		builder.AddQuery("include_count", fmt.Sprint(*listSynonymsOptions.IncludeCount))
	}
	if listSynonymsOptions.Sort != nil {
		builder.AddQuery("sort", fmt.Sprint(*listSynonymsOptions.Sort))
	}
	if listSynonymsOptions.Cursor != nil {
		builder.AddQuery("cursor", fmt.Sprint(*listSynonymsOptions.Cursor))
	}
	if listSynonymsOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*listSynonymsOptions.IncludeAudit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSynonymCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateSynonym : Create entity value synonym
// Add a new synonym to an entity value.
//
// If you want to create multiple synonyms with a single API call, consider using the **[Update
// entity](#update-entity)** or **[Update entity value](#update-entity-value)** method instead.
func (assistant *AssistantV1) CreateSynonym(createSynonymOptions *CreateSynonymOptions) (result *Synonym, response *core.DetailedResponse, err error) {
	return assistant.CreateSynonymWithContext(context.Background(), createSynonymOptions)
}

// CreateSynonymWithContext is an alternate form of the CreateSynonym method which supports a Context parameter
func (assistant *AssistantV1) CreateSynonymWithContext(ctx context.Context, createSynonymOptions *CreateSynonymOptions) (result *Synonym, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createSynonymOptions, "createSynonymOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createSynonymOptions, "createSynonymOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *createSynonymOptions.WorkspaceID,
		"entity":       *createSynonymOptions.Entity,
		"value":        *createSynonymOptions.Value,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}/synonyms`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createSynonymOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "CreateSynonym")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	if createSynonymOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*createSynonymOptions.IncludeAudit))
	}

	body := make(map[string]interface{})
	if createSynonymOptions.Synonym != nil {
		body["synonym"] = createSynonymOptions.Synonym
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
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSynonym)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetSynonym : Get entity value synonym
// Get information about a synonym of an entity value.
func (assistant *AssistantV1) GetSynonym(getSynonymOptions *GetSynonymOptions) (result *Synonym, response *core.DetailedResponse, err error) {
	return assistant.GetSynonymWithContext(context.Background(), getSynonymOptions)
}

// GetSynonymWithContext is an alternate form of the GetSynonym method which supports a Context parameter
func (assistant *AssistantV1) GetSynonymWithContext(ctx context.Context, getSynonymOptions *GetSynonymOptions) (result *Synonym, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getSynonymOptions, "getSynonymOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getSynonymOptions, "getSynonymOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *getSynonymOptions.WorkspaceID,
		"entity":       *getSynonymOptions.Entity,
		"value":        *getSynonymOptions.Value,
		"synonym":      *getSynonymOptions.Synonym,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}/synonyms/{synonym}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getSynonymOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "GetSynonym")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	if getSynonymOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*getSynonymOptions.IncludeAudit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSynonym)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateSynonym : Update entity value synonym
// Update an existing entity value synonym with new text.
//
// If you want to update multiple synonyms with a single API call, consider using the **[Update
// entity](#update-entity)** or **[Update entity value](#update-entity-value)** method instead.
func (assistant *AssistantV1) UpdateSynonym(updateSynonymOptions *UpdateSynonymOptions) (result *Synonym, response *core.DetailedResponse, err error) {
	return assistant.UpdateSynonymWithContext(context.Background(), updateSynonymOptions)
}

// UpdateSynonymWithContext is an alternate form of the UpdateSynonym method which supports a Context parameter
func (assistant *AssistantV1) UpdateSynonymWithContext(ctx context.Context, updateSynonymOptions *UpdateSynonymOptions) (result *Synonym, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateSynonymOptions, "updateSynonymOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateSynonymOptions, "updateSynonymOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *updateSynonymOptions.WorkspaceID,
		"entity":       *updateSynonymOptions.Entity,
		"value":        *updateSynonymOptions.Value,
		"synonym":      *updateSynonymOptions.Synonym,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}/synonyms/{synonym}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateSynonymOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "UpdateSynonym")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	if updateSynonymOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*updateSynonymOptions.IncludeAudit))
	}

	body := make(map[string]interface{})
	if updateSynonymOptions.NewSynonym != nil {
		body["synonym"] = updateSynonymOptions.NewSynonym
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
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSynonym)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteSynonym : Delete entity value synonym
// Delete a synonym from an entity value.
func (assistant *AssistantV1) DeleteSynonym(deleteSynonymOptions *DeleteSynonymOptions) (response *core.DetailedResponse, err error) {
	return assistant.DeleteSynonymWithContext(context.Background(), deleteSynonymOptions)
}

// DeleteSynonymWithContext is an alternate form of the DeleteSynonym method which supports a Context parameter
func (assistant *AssistantV1) DeleteSynonymWithContext(ctx context.Context, deleteSynonymOptions *DeleteSynonymOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteSynonymOptions, "deleteSynonymOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteSynonymOptions, "deleteSynonymOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *deleteSynonymOptions.WorkspaceID,
		"entity":       *deleteSynonymOptions.Entity,
		"value":        *deleteSynonymOptions.Value,
		"synonym":      *deleteSynonymOptions.Synonym,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}/synonyms/{synonym}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteSynonymOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "DeleteSynonym")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, nil)

	return
}

// ListDialogNodes : List dialog nodes
// List the dialog nodes for a workspace.
func (assistant *AssistantV1) ListDialogNodes(listDialogNodesOptions *ListDialogNodesOptions) (result *DialogNodeCollection, response *core.DetailedResponse, err error) {
	return assistant.ListDialogNodesWithContext(context.Background(), listDialogNodesOptions)
}

// ListDialogNodesWithContext is an alternate form of the ListDialogNodes method which supports a Context parameter
func (assistant *AssistantV1) ListDialogNodesWithContext(ctx context.Context, listDialogNodesOptions *ListDialogNodesOptions) (result *DialogNodeCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listDialogNodesOptions, "listDialogNodesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listDialogNodesOptions, "listDialogNodesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *listDialogNodesOptions.WorkspaceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/dialog_nodes`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listDialogNodesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "ListDialogNodes")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	if listDialogNodesOptions.PageLimit != nil {
		builder.AddQuery("page_limit", fmt.Sprint(*listDialogNodesOptions.PageLimit))
	}
	if listDialogNodesOptions.IncludeCount != nil {
		builder.AddQuery("include_count", fmt.Sprint(*listDialogNodesOptions.IncludeCount))
	}
	if listDialogNodesOptions.Sort != nil {
		builder.AddQuery("sort", fmt.Sprint(*listDialogNodesOptions.Sort))
	}
	if listDialogNodesOptions.Cursor != nil {
		builder.AddQuery("cursor", fmt.Sprint(*listDialogNodesOptions.Cursor))
	}
	if listDialogNodesOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*listDialogNodesOptions.IncludeAudit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDialogNodeCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateDialogNode : Create dialog node
// Create a new dialog node.
//
// If you want to create multiple dialog nodes with a single API call, consider using the **[Update
// workspace](#update-workspace)** method instead.
func (assistant *AssistantV1) CreateDialogNode(createDialogNodeOptions *CreateDialogNodeOptions) (result *DialogNode, response *core.DetailedResponse, err error) {
	return assistant.CreateDialogNodeWithContext(context.Background(), createDialogNodeOptions)
}

// CreateDialogNodeWithContext is an alternate form of the CreateDialogNode method which supports a Context parameter
func (assistant *AssistantV1) CreateDialogNodeWithContext(ctx context.Context, createDialogNodeOptions *CreateDialogNodeOptions) (result *DialogNode, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createDialogNodeOptions, "createDialogNodeOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createDialogNodeOptions, "createDialogNodeOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *createDialogNodeOptions.WorkspaceID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/dialog_nodes`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createDialogNodeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "CreateDialogNode")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	if createDialogNodeOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*createDialogNodeOptions.IncludeAudit))
	}

	body := make(map[string]interface{})
	if createDialogNodeOptions.DialogNode != nil {
		body["dialog_node"] = createDialogNodeOptions.DialogNode
	}
	if createDialogNodeOptions.Description != nil {
		body["description"] = createDialogNodeOptions.Description
	}
	if createDialogNodeOptions.Conditions != nil {
		body["conditions"] = createDialogNodeOptions.Conditions
	}
	if createDialogNodeOptions.Parent != nil {
		body["parent"] = createDialogNodeOptions.Parent
	}
	if createDialogNodeOptions.PreviousSibling != nil {
		body["previous_sibling"] = createDialogNodeOptions.PreviousSibling
	}
	if createDialogNodeOptions.Output != nil {
		body["output"] = createDialogNodeOptions.Output
	}
	if createDialogNodeOptions.Context != nil {
		body["context"] = createDialogNodeOptions.Context
	}
	if createDialogNodeOptions.Metadata != nil {
		body["metadata"] = createDialogNodeOptions.Metadata
	}
	if createDialogNodeOptions.NextStep != nil {
		body["next_step"] = createDialogNodeOptions.NextStep
	}
	if createDialogNodeOptions.Title != nil {
		body["title"] = createDialogNodeOptions.Title
	}
	if createDialogNodeOptions.Type != nil {
		body["type"] = createDialogNodeOptions.Type
	}
	if createDialogNodeOptions.EventName != nil {
		body["event_name"] = createDialogNodeOptions.EventName
	}
	if createDialogNodeOptions.Variable != nil {
		body["variable"] = createDialogNodeOptions.Variable
	}
	if createDialogNodeOptions.Actions != nil {
		body["actions"] = createDialogNodeOptions.Actions
	}
	if createDialogNodeOptions.DigressIn != nil {
		body["digress_in"] = createDialogNodeOptions.DigressIn
	}
	if createDialogNodeOptions.DigressOut != nil {
		body["digress_out"] = createDialogNodeOptions.DigressOut
	}
	if createDialogNodeOptions.DigressOutSlots != nil {
		body["digress_out_slots"] = createDialogNodeOptions.DigressOutSlots
	}
	if createDialogNodeOptions.UserLabel != nil {
		body["user_label"] = createDialogNodeOptions.UserLabel
	}
	if createDialogNodeOptions.DisambiguationOptOut != nil {
		body["disambiguation_opt_out"] = createDialogNodeOptions.DisambiguationOptOut
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
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDialogNode)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetDialogNode : Get dialog node
// Get information about a dialog node.
func (assistant *AssistantV1) GetDialogNode(getDialogNodeOptions *GetDialogNodeOptions) (result *DialogNode, response *core.DetailedResponse, err error) {
	return assistant.GetDialogNodeWithContext(context.Background(), getDialogNodeOptions)
}

// GetDialogNodeWithContext is an alternate form of the GetDialogNode method which supports a Context parameter
func (assistant *AssistantV1) GetDialogNodeWithContext(ctx context.Context, getDialogNodeOptions *GetDialogNodeOptions) (result *DialogNode, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDialogNodeOptions, "getDialogNodeOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getDialogNodeOptions, "getDialogNodeOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *getDialogNodeOptions.WorkspaceID,
		"dialog_node":  *getDialogNodeOptions.DialogNode,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/dialog_nodes/{dialog_node}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getDialogNodeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "GetDialogNode")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	if getDialogNodeOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*getDialogNodeOptions.IncludeAudit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDialogNode)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateDialogNode : Update dialog node
// Update an existing dialog node with new or modified data.
//
// If you want to update multiple dialog nodes with a single API call, consider using the **[Update
// workspace](#update-workspace)** method instead.
func (assistant *AssistantV1) UpdateDialogNode(updateDialogNodeOptions *UpdateDialogNodeOptions) (result *DialogNode, response *core.DetailedResponse, err error) {
	return assistant.UpdateDialogNodeWithContext(context.Background(), updateDialogNodeOptions)
}

// UpdateDialogNodeWithContext is an alternate form of the UpdateDialogNode method which supports a Context parameter
func (assistant *AssistantV1) UpdateDialogNodeWithContext(ctx context.Context, updateDialogNodeOptions *UpdateDialogNodeOptions) (result *DialogNode, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateDialogNodeOptions, "updateDialogNodeOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateDialogNodeOptions, "updateDialogNodeOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *updateDialogNodeOptions.WorkspaceID,
		"dialog_node":  *updateDialogNodeOptions.DialogNode,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/dialog_nodes/{dialog_node}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateDialogNodeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "UpdateDialogNode")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	if updateDialogNodeOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*updateDialogNodeOptions.IncludeAudit))
	}

	body := make(map[string]interface{})
	if updateDialogNodeOptions.NewDialogNode != nil {
		body["dialog_node"] = updateDialogNodeOptions.NewDialogNode
	}
	if updateDialogNodeOptions.NewDescription != nil {
		body["description"] = updateDialogNodeOptions.NewDescription
	}
	if updateDialogNodeOptions.NewConditions != nil {
		body["conditions"] = updateDialogNodeOptions.NewConditions
	}
	if updateDialogNodeOptions.NewParent != nil {
		body["parent"] = updateDialogNodeOptions.NewParent
	}
	if updateDialogNodeOptions.NewPreviousSibling != nil {
		body["previous_sibling"] = updateDialogNodeOptions.NewPreviousSibling
	}
	if updateDialogNodeOptions.NewOutput != nil {
		body["output"] = updateDialogNodeOptions.NewOutput
	}
	if updateDialogNodeOptions.NewContext != nil {
		body["context"] = updateDialogNodeOptions.NewContext
	}
	if updateDialogNodeOptions.NewMetadata != nil {
		body["metadata"] = updateDialogNodeOptions.NewMetadata
	}
	if updateDialogNodeOptions.NewNextStep != nil {
		body["next_step"] = updateDialogNodeOptions.NewNextStep
	}
	if updateDialogNodeOptions.NewTitle != nil {
		body["title"] = updateDialogNodeOptions.NewTitle
	}
	if updateDialogNodeOptions.NewType != nil {
		body["type"] = updateDialogNodeOptions.NewType
	}
	if updateDialogNodeOptions.NewEventName != nil {
		body["event_name"] = updateDialogNodeOptions.NewEventName
	}
	if updateDialogNodeOptions.NewVariable != nil {
		body["variable"] = updateDialogNodeOptions.NewVariable
	}
	if updateDialogNodeOptions.NewActions != nil {
		body["actions"] = updateDialogNodeOptions.NewActions
	}
	if updateDialogNodeOptions.NewDigressIn != nil {
		body["digress_in"] = updateDialogNodeOptions.NewDigressIn
	}
	if updateDialogNodeOptions.NewDigressOut != nil {
		body["digress_out"] = updateDialogNodeOptions.NewDigressOut
	}
	if updateDialogNodeOptions.NewDigressOutSlots != nil {
		body["digress_out_slots"] = updateDialogNodeOptions.NewDigressOutSlots
	}
	if updateDialogNodeOptions.NewUserLabel != nil {
		body["user_label"] = updateDialogNodeOptions.NewUserLabel
	}
	if updateDialogNodeOptions.NewDisambiguationOptOut != nil {
		body["disambiguation_opt_out"] = updateDialogNodeOptions.NewDisambiguationOptOut
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
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDialogNode)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteDialogNode : Delete dialog node
// Delete a dialog node from a workspace.
func (assistant *AssistantV1) DeleteDialogNode(deleteDialogNodeOptions *DeleteDialogNodeOptions) (response *core.DetailedResponse, err error) {
	return assistant.DeleteDialogNodeWithContext(context.Background(), deleteDialogNodeOptions)
}

// DeleteDialogNodeWithContext is an alternate form of the DeleteDialogNode method which supports a Context parameter
func (assistant *AssistantV1) DeleteDialogNodeWithContext(ctx context.Context, deleteDialogNodeOptions *DeleteDialogNodeOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteDialogNodeOptions, "deleteDialogNodeOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteDialogNodeOptions, "deleteDialogNodeOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *deleteDialogNodeOptions.WorkspaceID,
		"dialog_node":  *deleteDialogNodeOptions.DialogNode,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/dialog_nodes/{dialog_node}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteDialogNodeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "DeleteDialogNode")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, nil)

	return
}

// ListLogs : List log events in a workspace
// List the events from the log of a specific workspace.
//
// This method requires Manager access.
func (assistant *AssistantV1) ListLogs(listLogsOptions *ListLogsOptions) (result *LogCollection, response *core.DetailedResponse, err error) {
	return assistant.ListLogsWithContext(context.Background(), listLogsOptions)
}

// ListLogsWithContext is an alternate form of the ListLogs method which supports a Context parameter
func (assistant *AssistantV1) ListLogsWithContext(ctx context.Context, listLogsOptions *ListLogsOptions) (result *LogCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listLogsOptions, "listLogsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listLogsOptions, "listLogsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"workspace_id": *listLogsOptions.WorkspaceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/workspaces/{workspace_id}/logs`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listLogsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "ListLogs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	if listLogsOptions.Sort != nil {
		builder.AddQuery("sort", fmt.Sprint(*listLogsOptions.Sort))
	}
	if listLogsOptions.Filter != nil {
		builder.AddQuery("filter", fmt.Sprint(*listLogsOptions.Filter))
	}
	if listLogsOptions.PageLimit != nil {
		builder.AddQuery("page_limit", fmt.Sprint(*listLogsOptions.PageLimit))
	}
	if listLogsOptions.Cursor != nil {
		builder.AddQuery("cursor", fmt.Sprint(*listLogsOptions.Cursor))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalLogCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListAllLogs : List log events in all workspaces
// List the events from the logs of all workspaces in the service instance.
func (assistant *AssistantV1) ListAllLogs(listAllLogsOptions *ListAllLogsOptions) (result *LogCollection, response *core.DetailedResponse, err error) {
	return assistant.ListAllLogsWithContext(context.Background(), listAllLogsOptions)
}

// ListAllLogsWithContext is an alternate form of the ListAllLogs method which supports a Context parameter
func (assistant *AssistantV1) ListAllLogsWithContext(ctx context.Context, listAllLogsOptions *ListAllLogsOptions) (result *LogCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listAllLogsOptions, "listAllLogsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listAllLogsOptions, "listAllLogsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/logs`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listAllLogsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "ListAllLogs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	builder.AddQuery("filter", fmt.Sprint(*listAllLogsOptions.Filter))
	if listAllLogsOptions.Sort != nil {
		builder.AddQuery("sort", fmt.Sprint(*listAllLogsOptions.Sort))
	}
	if listAllLogsOptions.PageLimit != nil {
		builder.AddQuery("page_limit", fmt.Sprint(*listAllLogsOptions.PageLimit))
	}
	if listAllLogsOptions.Cursor != nil {
		builder.AddQuery("cursor", fmt.Sprint(*listAllLogsOptions.Cursor))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalLogCollection)
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
// security](https://cloud.ibm.com/docs/assistant?topic=assistant-information-security#information-security).
//
// **Note:** This operation is intended only for deleting data associated with a single specific customer, not for
// deleting data associated with multiple customers or for any other purpose. For more information, see [Labeling and
// deleting data in Watson
// Assistant](https://cloud.ibm.com/docs/assistant?topic=assistant-information-security#information-security-gdpr-wa).
func (assistant *AssistantV1) DeleteUserData(deleteUserDataOptions *DeleteUserDataOptions) (response *core.DetailedResponse, err error) {
	return assistant.DeleteUserDataWithContext(context.Background(), deleteUserDataOptions)
}

// DeleteUserDataWithContext is an alternate form of the DeleteUserData method which supports a Context parameter
func (assistant *AssistantV1) DeleteUserDataWithContext(ctx context.Context, deleteUserDataOptions *DeleteUserDataOptions) (response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v1/user_data`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteUserDataOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V1", "DeleteUserData")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	builder.AddQuery("customer_id", fmt.Sprint(*deleteUserDataOptions.CustomerID))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, nil)

	return
}

// AgentAvailabilityMessage : AgentAvailabilityMessage struct
type AgentAvailabilityMessage struct {
	// The text of the message.
	Message *string `json:"message,omitempty"`
}

// UnmarshalAgentAvailabilityMessage unmarshals an instance of AgentAvailabilityMessage from the specified map of raw messages.
func UnmarshalAgentAvailabilityMessage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AgentAvailabilityMessage)
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BulkClassifyOptions : The BulkClassify options.
type BulkClassifyOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// An array of input utterances to classify.
	Input []BulkClassifyUtterance `json:"input,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewBulkClassifyOptions : Instantiate BulkClassifyOptions
func (*AssistantV1) NewBulkClassifyOptions(workspaceID string) *BulkClassifyOptions {
	return &BulkClassifyOptions{
		WorkspaceID: core.StringPtr(workspaceID),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *BulkClassifyOptions) SetWorkspaceID(workspaceID string) *BulkClassifyOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetInput : Allow user to set Input
func (_options *BulkClassifyOptions) SetInput(input []BulkClassifyUtterance) *BulkClassifyOptions {
	_options.Input = input
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *BulkClassifyOptions) SetHeaders(param map[string]string) *BulkClassifyOptions {
	options.Headers = param
	return options
}

// BulkClassifyOutput : BulkClassifyOutput struct
type BulkClassifyOutput struct {
	// The user input utterance to classify.
	Input *BulkClassifyUtterance `json:"input,omitempty"`

	// An array of entities identified in the utterance.
	Entities []RuntimeEntity `json:"entities,omitempty"`

	// An array of intents recognized in the utterance.
	Intents []RuntimeIntent `json:"intents,omitempty"`
}

// UnmarshalBulkClassifyOutput unmarshals an instance of BulkClassifyOutput from the specified map of raw messages.
func UnmarshalBulkClassifyOutput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BulkClassifyOutput)
	err = core.UnmarshalModel(m, "input", &obj.Input, UnmarshalBulkClassifyUtterance)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "entities", &obj.Entities, UnmarshalRuntimeEntity)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "intents", &obj.Intents, UnmarshalRuntimeIntent)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BulkClassifyResponse : BulkClassifyResponse struct
type BulkClassifyResponse struct {
	// An array of objects that contain classification information for the submitted input utterances.
	Output []BulkClassifyOutput `json:"output,omitempty"`
}

// UnmarshalBulkClassifyResponse unmarshals an instance of BulkClassifyResponse from the specified map of raw messages.
func UnmarshalBulkClassifyResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BulkClassifyResponse)
	err = core.UnmarshalModel(m, "output", &obj.Output, UnmarshalBulkClassifyOutput)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BulkClassifyUtterance : The user input utterance to classify.
type BulkClassifyUtterance struct {
	// The text of the input utterance.
	Text *string `json:"text" validate:"required"`
}

// NewBulkClassifyUtterance : Instantiate BulkClassifyUtterance (Generic Model Constructor)
func (*AssistantV1) NewBulkClassifyUtterance(text string) (_model *BulkClassifyUtterance, err error) {
	_model = &BulkClassifyUtterance{
		Text: core.StringPtr(text),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalBulkClassifyUtterance unmarshals an instance of BulkClassifyUtterance from the specified map of raw messages.
func UnmarshalBulkClassifyUtterance(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BulkClassifyUtterance)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CaptureGroup : A recognized capture group for a pattern-based entity.
type CaptureGroup struct {
	// A recognized capture group for the entity.
	Group *string `json:"group" validate:"required"`

	// Zero-based character offsets that indicate where the entity value begins and ends in the input text.
	Location []int64 `json:"location,omitempty"`
}

// NewCaptureGroup : Instantiate CaptureGroup (Generic Model Constructor)
func (*AssistantV1) NewCaptureGroup(group string) (_model *CaptureGroup, err error) {
	_model = &CaptureGroup{
		Group: core.StringPtr(group),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalCaptureGroup unmarshals an instance of CaptureGroup from the specified map of raw messages.
func UnmarshalCaptureGroup(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CaptureGroup)
	err = core.UnmarshalPrimitive(m, "group", &obj.Group)
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

// ChannelTransferInfo : Information used by an integration to transfer the conversation to a different channel.
type ChannelTransferInfo struct {
	// An object specifying target channels available for the transfer. Each property of this object represents an
	// available transfer target. Currently, the only supported property is **chat**, representing the web chat
	// integration.
	Target *ChannelTransferTarget `json:"target" validate:"required"`
}

// NewChannelTransferInfo : Instantiate ChannelTransferInfo (Generic Model Constructor)
func (*AssistantV1) NewChannelTransferInfo(target *ChannelTransferTarget) (_model *ChannelTransferInfo, err error) {
	_model = &ChannelTransferInfo{
		Target: target,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalChannelTransferInfo unmarshals an instance of ChannelTransferInfo from the specified map of raw messages.
func UnmarshalChannelTransferInfo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ChannelTransferInfo)
	err = core.UnmarshalModel(m, "target", &obj.Target, UnmarshalChannelTransferTarget)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ChannelTransferTarget : An object specifying target channels available for the transfer. Each property of this object represents an available
// transfer target. Currently, the only supported property is **chat**, representing the web chat integration.
type ChannelTransferTarget struct {
	// Information for transferring to the web chat integration.
	Chat *ChannelTransferTargetChat `json:"chat,omitempty"`
}

// UnmarshalChannelTransferTarget unmarshals an instance of ChannelTransferTarget from the specified map of raw messages.
func UnmarshalChannelTransferTarget(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ChannelTransferTarget)
	err = core.UnmarshalModel(m, "chat", &obj.Chat, UnmarshalChannelTransferTargetChat)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ChannelTransferTargetChat : Information for transferring to the web chat integration.
type ChannelTransferTargetChat struct {
	// The URL of the target web chat.
	URL *string `json:"url,omitempty"`
}

// UnmarshalChannelTransferTargetChat unmarshals an instance of ChannelTransferTargetChat from the specified map of raw messages.
func UnmarshalChannelTransferTargetChat(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ChannelTransferTargetChat)
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Context : State information for the conversation. To maintain state, include the context from the previous response.
type Context struct {
	// The unique identifier of the conversation.
	ConversationID *string `json:"conversation_id,omitempty"`

	// For internal use only.
	System map[string]interface{} `json:"system,omitempty"`

	// Metadata related to the message.
	Metadata *MessageContextMetadata `json:"metadata,omitempty"`

	// Allows users to set arbitrary properties
	additionalProperties map[string]interface{}
}

// SetProperty allows the user to set an arbitrary property on an instance of Context
func (o *Context) SetProperty(key string, value interface{}) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string]interface{})
	}
	o.additionalProperties[key] = value
}

// SetProperties allows the user to set a map of arbitrary properties on an instance of Context
func (o *Context) SetProperties(m map[string]interface{}) {
	o.additionalProperties = make(map[string]interface{})
	for k, v := range m {
		o.additionalProperties[k] = v
	}
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of Context
func (o *Context) GetProperty(key string) interface{} {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of Context
func (o *Context) GetProperties() map[string]interface{} {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of Context
func (o *Context) MarshalJSON() (buffer []byte, err error) {
	m := make(map[string]interface{})
	if len(o.additionalProperties) > 0 {
		for k, v := range o.additionalProperties {
			m[k] = v
		}
	}
	if o.ConversationID != nil {
		m["conversation_id"] = o.ConversationID
	}
	if o.System != nil {
		m["system"] = o.System
	}
	if o.Metadata != nil {
		m["metadata"] = o.Metadata
	}
	buffer, err = json.Marshal(m)
	return
}

// UnmarshalContext unmarshals an instance of Context from the specified map of raw messages.
func UnmarshalContext(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Context)
	err = core.UnmarshalPrimitive(m, "conversation_id", &obj.ConversationID)
	if err != nil {
		return
	}
	delete(m, "conversation_id")
	err = core.UnmarshalPrimitive(m, "system", &obj.System)
	if err != nil {
		return
	}
	delete(m, "system")
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalMessageContextMetadata)
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

// Counterexample : Counterexample struct
type Counterexample struct {
	// The text of a user input marked as irrelevant input. This string must conform to the following restrictions:
	// - It cannot contain carriage return, newline, or tab characters.
	// - It cannot consist of only whitespace characters.
	Text *string `json:"text" validate:"required"`

	// The timestamp for creation of the object.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// The timestamp for the most recent update to the object.
	Updated *strfmt.DateTime `json:"updated,omitempty"`
}

// NewCounterexample : Instantiate Counterexample (Generic Model Constructor)
func (*AssistantV1) NewCounterexample(text string) (_model *Counterexample, err error) {
	_model = &Counterexample{
		Text: core.StringPtr(text),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalCounterexample unmarshals an instance of Counterexample from the specified map of raw messages.
func UnmarshalCounterexample(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Counterexample)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
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

// CounterexampleCollection : CounterexampleCollection struct
type CounterexampleCollection struct {
	// An array of objects describing the examples marked as irrelevant input.
	Counterexamples []Counterexample `json:"counterexamples" validate:"required"`

	// The pagination data for the returned objects.
	Pagination *Pagination `json:"pagination" validate:"required"`
}

// UnmarshalCounterexampleCollection unmarshals an instance of CounterexampleCollection from the specified map of raw messages.
func UnmarshalCounterexampleCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CounterexampleCollection)
	err = core.UnmarshalModel(m, "counterexamples", &obj.Counterexamples, UnmarshalCounterexample)
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

// CreateCounterexampleOptions : The CreateCounterexample options.
type CreateCounterexampleOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// The text of a user input marked as irrelevant input. This string must conform to the following restrictions:
	// - It cannot contain carriage return, newline, or tab characters.
	// - It cannot consist of only whitespace characters.
	Text *string `json:"text" validate:"required"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateCounterexampleOptions : Instantiate CreateCounterexampleOptions
func (*AssistantV1) NewCreateCounterexampleOptions(workspaceID string, text string) *CreateCounterexampleOptions {
	return &CreateCounterexampleOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Text:        core.StringPtr(text),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *CreateCounterexampleOptions) SetWorkspaceID(workspaceID string) *CreateCounterexampleOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetText : Allow user to set Text
func (_options *CreateCounterexampleOptions) SetText(text string) *CreateCounterexampleOptions {
	_options.Text = core.StringPtr(text)
	return _options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (_options *CreateCounterexampleOptions) SetIncludeAudit(includeAudit bool) *CreateCounterexampleOptions {
	_options.IncludeAudit = core.BoolPtr(includeAudit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateCounterexampleOptions) SetHeaders(param map[string]string) *CreateCounterexampleOptions {
	options.Headers = param
	return options
}

// CreateDialogNodeOptions : The CreateDialogNode options.
type CreateDialogNodeOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// The unique ID of the dialog node. This is an internal identifier used to refer to the dialog node from other dialog
	// nodes and in the diagnostic information included with message responses.
	//
	// This string can contain only Unicode alphanumeric, space, underscore, hyphen, and dot characters.
	DialogNode *string `json:"dialog_node" validate:"required"`

	// The description of the dialog node. This string cannot contain carriage return, newline, or tab characters.
	Description *string `json:"description,omitempty"`

	// The condition that will trigger the dialog node. This string cannot contain carriage return, newline, or tab
	// characters.
	Conditions *string `json:"conditions,omitempty"`

	// The unique ID of the parent dialog node. This property is omitted if the dialog node has no parent.
	Parent *string `json:"parent,omitempty"`

	// The unique ID of the previous sibling dialog node. This property is omitted if the dialog node has no previous
	// sibling.
	PreviousSibling *string `json:"previous_sibling,omitempty"`

	// The output of the dialog node. For more information about how to specify dialog node output, see the
	// [documentation](https://cloud.ibm.com/docs/assistant?topic=assistant-dialog-overview#dialog-overview-responses).
	Output *DialogNodeOutput `json:"output,omitempty"`

	// The context for the dialog node.
	Context *DialogNodeContext `json:"context,omitempty"`

	// The metadata for the dialog node.
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// The next step to execute following this dialog node.
	NextStep *DialogNodeNextStep `json:"next_step,omitempty"`

	// A human-readable name for the dialog node. If the node is included in disambiguation, this title is used to populate
	// the **label** property of the corresponding suggestion in the `suggestion` response type (unless it is overridden by
	// the **user_label** property). The title is also used to populate the **topic** property in the `connect_to_agent`
	// response type.
	//
	// This string can contain only Unicode alphanumeric, space, underscore, hyphen, and dot characters.
	Title *string `json:"title,omitempty"`

	// How the dialog node is processed.
	Type *string `json:"type,omitempty"`

	// How an `event_handler` node is processed.
	EventName *string `json:"event_name,omitempty"`

	// The location in the dialog context where output is stored.
	Variable *string `json:"variable,omitempty"`

	// An array of objects describing any actions to be invoked by the dialog node.
	Actions []DialogNodeAction `json:"actions,omitempty"`

	// Whether this top-level dialog node can be digressed into.
	DigressIn *string `json:"digress_in,omitempty"`

	// Whether this dialog node can be returned to after a digression.
	DigressOut *string `json:"digress_out,omitempty"`

	// Whether the user can digress to top-level nodes while filling out slots.
	DigressOutSlots *string `json:"digress_out_slots,omitempty"`

	// A label that can be displayed externally to describe the purpose of the node to users. If set, this label is used to
	// identify the node in disambiguation responses (overriding the value of the **title** property).
	UserLabel *string `json:"user_label,omitempty"`

	// Whether the dialog node should be excluded from disambiguation suggestions. Valid only when **type**=`standard` or
	// `frame`.
	DisambiguationOptOut *bool `json:"disambiguation_opt_out,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateDialogNodeOptions.Type property.
// How the dialog node is processed.
const (
	CreateDialogNodeOptionsTypeEventHandlerConst      = "event_handler"
	CreateDialogNodeOptionsTypeFolderConst            = "folder"
	CreateDialogNodeOptionsTypeFrameConst             = "frame"
	CreateDialogNodeOptionsTypeResponseConditionConst = "response_condition"
	CreateDialogNodeOptionsTypeSlotConst              = "slot"
	CreateDialogNodeOptionsTypeStandardConst          = "standard"
)

// Constants associated with the CreateDialogNodeOptions.EventName property.
// How an `event_handler` node is processed.
const (
	CreateDialogNodeOptionsEventNameDigressionReturnPromptConst   = "digression_return_prompt"
	CreateDialogNodeOptionsEventNameFilledConst                   = "filled"
	CreateDialogNodeOptionsEventNameFilledMultipleConst           = "filled_multiple"
	CreateDialogNodeOptionsEventNameFocusConst                    = "focus"
	CreateDialogNodeOptionsEventNameGenericConst                  = "generic"
	CreateDialogNodeOptionsEventNameInputConst                    = "input"
	CreateDialogNodeOptionsEventNameNomatchConst                  = "nomatch"
	CreateDialogNodeOptionsEventNameNomatchResponsesDepletedConst = "nomatch_responses_depleted"
	CreateDialogNodeOptionsEventNameValidateConst                 = "validate"
)

// Constants associated with the CreateDialogNodeOptions.DigressIn property.
// Whether this top-level dialog node can be digressed into.
const (
	CreateDialogNodeOptionsDigressInDoesNotReturnConst = "does_not_return"
	CreateDialogNodeOptionsDigressInNotAvailableConst  = "not_available"
	CreateDialogNodeOptionsDigressInReturnsConst       = "returns"
)

// Constants associated with the CreateDialogNodeOptions.DigressOut property.
// Whether this dialog node can be returned to after a digression.
const (
	CreateDialogNodeOptionsDigressOutAllowAllConst            = "allow_all"
	CreateDialogNodeOptionsDigressOutAllowAllNeverReturnConst = "allow_all_never_return"
	CreateDialogNodeOptionsDigressOutAllowReturningConst      = "allow_returning"
)

// Constants associated with the CreateDialogNodeOptions.DigressOutSlots property.
// Whether the user can digress to top-level nodes while filling out slots.
const (
	CreateDialogNodeOptionsDigressOutSlotsAllowAllConst       = "allow_all"
	CreateDialogNodeOptionsDigressOutSlotsAllowReturningConst = "allow_returning"
	CreateDialogNodeOptionsDigressOutSlotsNotAllowedConst     = "not_allowed"
)

// NewCreateDialogNodeOptions : Instantiate CreateDialogNodeOptions
func (*AssistantV1) NewCreateDialogNodeOptions(workspaceID string, dialogNode string) *CreateDialogNodeOptions {
	return &CreateDialogNodeOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		DialogNode:  core.StringPtr(dialogNode),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *CreateDialogNodeOptions) SetWorkspaceID(workspaceID string) *CreateDialogNodeOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetDialogNode : Allow user to set DialogNode
func (_options *CreateDialogNodeOptions) SetDialogNode(dialogNode string) *CreateDialogNodeOptions {
	_options.DialogNode = core.StringPtr(dialogNode)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateDialogNodeOptions) SetDescription(description string) *CreateDialogNodeOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetConditions : Allow user to set Conditions
func (_options *CreateDialogNodeOptions) SetConditions(conditions string) *CreateDialogNodeOptions {
	_options.Conditions = core.StringPtr(conditions)
	return _options
}

// SetParent : Allow user to set Parent
func (_options *CreateDialogNodeOptions) SetParent(parent string) *CreateDialogNodeOptions {
	_options.Parent = core.StringPtr(parent)
	return _options
}

// SetPreviousSibling : Allow user to set PreviousSibling
func (_options *CreateDialogNodeOptions) SetPreviousSibling(previousSibling string) *CreateDialogNodeOptions {
	_options.PreviousSibling = core.StringPtr(previousSibling)
	return _options
}

// SetOutput : Allow user to set Output
func (_options *CreateDialogNodeOptions) SetOutput(output *DialogNodeOutput) *CreateDialogNodeOptions {
	_options.Output = output
	return _options
}

// SetContext : Allow user to set Context
func (_options *CreateDialogNodeOptions) SetContext(context *DialogNodeContext) *CreateDialogNodeOptions {
	_options.Context = context
	return _options
}

// SetMetadata : Allow user to set Metadata
func (_options *CreateDialogNodeOptions) SetMetadata(metadata map[string]interface{}) *CreateDialogNodeOptions {
	_options.Metadata = metadata
	return _options
}

// SetNextStep : Allow user to set NextStep
func (_options *CreateDialogNodeOptions) SetNextStep(nextStep *DialogNodeNextStep) *CreateDialogNodeOptions {
	_options.NextStep = nextStep
	return _options
}

// SetTitle : Allow user to set Title
func (_options *CreateDialogNodeOptions) SetTitle(title string) *CreateDialogNodeOptions {
	_options.Title = core.StringPtr(title)
	return _options
}

// SetType : Allow user to set Type
func (_options *CreateDialogNodeOptions) SetType(typeVar string) *CreateDialogNodeOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetEventName : Allow user to set EventName
func (_options *CreateDialogNodeOptions) SetEventName(eventName string) *CreateDialogNodeOptions {
	_options.EventName = core.StringPtr(eventName)
	return _options
}

// SetVariable : Allow user to set Variable
func (_options *CreateDialogNodeOptions) SetVariable(variable string) *CreateDialogNodeOptions {
	_options.Variable = core.StringPtr(variable)
	return _options
}

// SetActions : Allow user to set Actions
func (_options *CreateDialogNodeOptions) SetActions(actions []DialogNodeAction) *CreateDialogNodeOptions {
	_options.Actions = actions
	return _options
}

// SetDigressIn : Allow user to set DigressIn
func (_options *CreateDialogNodeOptions) SetDigressIn(digressIn string) *CreateDialogNodeOptions {
	_options.DigressIn = core.StringPtr(digressIn)
	return _options
}

// SetDigressOut : Allow user to set DigressOut
func (_options *CreateDialogNodeOptions) SetDigressOut(digressOut string) *CreateDialogNodeOptions {
	_options.DigressOut = core.StringPtr(digressOut)
	return _options
}

// SetDigressOutSlots : Allow user to set DigressOutSlots
func (_options *CreateDialogNodeOptions) SetDigressOutSlots(digressOutSlots string) *CreateDialogNodeOptions {
	_options.DigressOutSlots = core.StringPtr(digressOutSlots)
	return _options
}

// SetUserLabel : Allow user to set UserLabel
func (_options *CreateDialogNodeOptions) SetUserLabel(userLabel string) *CreateDialogNodeOptions {
	_options.UserLabel = core.StringPtr(userLabel)
	return _options
}

// SetDisambiguationOptOut : Allow user to set DisambiguationOptOut
func (_options *CreateDialogNodeOptions) SetDisambiguationOptOut(disambiguationOptOut bool) *CreateDialogNodeOptions {
	_options.DisambiguationOptOut = core.BoolPtr(disambiguationOptOut)
	return _options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (_options *CreateDialogNodeOptions) SetIncludeAudit(includeAudit bool) *CreateDialogNodeOptions {
	_options.IncludeAudit = core.BoolPtr(includeAudit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateDialogNodeOptions) SetHeaders(param map[string]string) *CreateDialogNodeOptions {
	options.Headers = param
	return options
}

// CreateEntity : CreateEntity struct
type CreateEntity struct {
	// The name of the entity. This string must conform to the following restrictions:
	// - It can contain only Unicode alphanumeric, underscore, and hyphen characters.
	// - If you specify an entity name beginning with the reserved prefix `sys-`, it must be the name of a system entity
	// that you want to enable. (Any entity content specified with the request is ignored.).
	Entity *string `json:"entity" validate:"required"`

	// The description of the entity. This string cannot contain carriage return, newline, or tab characters.
	Description *string `json:"description,omitempty"`

	// Any metadata related to the entity.
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// Whether to use fuzzy matching for the entity.
	FuzzyMatch *bool `json:"fuzzy_match,omitempty"`

	// The timestamp for creation of the object.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// The timestamp for the most recent update to the object.
	Updated *strfmt.DateTime `json:"updated,omitempty"`

	// An array of objects describing the entity values.
	Values []CreateValue `json:"values,omitempty"`
}

// NewCreateEntity : Instantiate CreateEntity (Generic Model Constructor)
func (*AssistantV1) NewCreateEntity(entity string) (_model *CreateEntity, err error) {
	_model = &CreateEntity{
		Entity: core.StringPtr(entity),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalCreateEntity unmarshals an instance of CreateEntity from the specified map of raw messages.
func UnmarshalCreateEntity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateEntity)
	err = core.UnmarshalPrimitive(m, "entity", &obj.Entity)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "metadata", &obj.Metadata)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "fuzzy_match", &obj.FuzzyMatch)
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
	err = core.UnmarshalModel(m, "values", &obj.Values, UnmarshalCreateValue)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateEntityOptions : The CreateEntity options.
type CreateEntityOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// The name of the entity. This string must conform to the following restrictions:
	// - It can contain only Unicode alphanumeric, underscore, and hyphen characters.
	// - If you specify an entity name beginning with the reserved prefix `sys-`, it must be the name of a system entity
	// that you want to enable. (Any entity content specified with the request is ignored.).
	Entity *string `json:"entity" validate:"required"`

	// The description of the entity. This string cannot contain carriage return, newline, or tab characters.
	Description *string `json:"description,omitempty"`

	// Any metadata related to the entity.
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// Whether to use fuzzy matching for the entity.
	FuzzyMatch *bool `json:"fuzzy_match,omitempty"`

	// An array of objects describing the entity values.
	Values []CreateValue `json:"values,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateEntityOptions : Instantiate CreateEntityOptions
func (*AssistantV1) NewCreateEntityOptions(workspaceID string, entity string) *CreateEntityOptions {
	return &CreateEntityOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Entity:      core.StringPtr(entity),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *CreateEntityOptions) SetWorkspaceID(workspaceID string) *CreateEntityOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetEntity : Allow user to set Entity
func (_options *CreateEntityOptions) SetEntity(entity string) *CreateEntityOptions {
	_options.Entity = core.StringPtr(entity)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateEntityOptions) SetDescription(description string) *CreateEntityOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetMetadata : Allow user to set Metadata
func (_options *CreateEntityOptions) SetMetadata(metadata map[string]interface{}) *CreateEntityOptions {
	_options.Metadata = metadata
	return _options
}

// SetFuzzyMatch : Allow user to set FuzzyMatch
func (_options *CreateEntityOptions) SetFuzzyMatch(fuzzyMatch bool) *CreateEntityOptions {
	_options.FuzzyMatch = core.BoolPtr(fuzzyMatch)
	return _options
}

// SetValues : Allow user to set Values
func (_options *CreateEntityOptions) SetValues(values []CreateValue) *CreateEntityOptions {
	_options.Values = values
	return _options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (_options *CreateEntityOptions) SetIncludeAudit(includeAudit bool) *CreateEntityOptions {
	_options.IncludeAudit = core.BoolPtr(includeAudit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateEntityOptions) SetHeaders(param map[string]string) *CreateEntityOptions {
	options.Headers = param
	return options
}

// CreateExampleOptions : The CreateExample options.
type CreateExampleOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// The intent name.
	Intent *string `json:"-" validate:"required,ne="`

	// The text of a user input example. This string must conform to the following restrictions:
	// - It cannot contain carriage return, newline, or tab characters.
	// - It cannot consist of only whitespace characters.
	Text *string `json:"text" validate:"required"`

	// An array of contextual entity mentions.
	Mentions []Mention `json:"mentions,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateExampleOptions : Instantiate CreateExampleOptions
func (*AssistantV1) NewCreateExampleOptions(workspaceID string, intent string, text string) *CreateExampleOptions {
	return &CreateExampleOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Intent:      core.StringPtr(intent),
		Text:        core.StringPtr(text),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *CreateExampleOptions) SetWorkspaceID(workspaceID string) *CreateExampleOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetIntent : Allow user to set Intent
func (_options *CreateExampleOptions) SetIntent(intent string) *CreateExampleOptions {
	_options.Intent = core.StringPtr(intent)
	return _options
}

// SetText : Allow user to set Text
func (_options *CreateExampleOptions) SetText(text string) *CreateExampleOptions {
	_options.Text = core.StringPtr(text)
	return _options
}

// SetMentions : Allow user to set Mentions
func (_options *CreateExampleOptions) SetMentions(mentions []Mention) *CreateExampleOptions {
	_options.Mentions = mentions
	return _options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (_options *CreateExampleOptions) SetIncludeAudit(includeAudit bool) *CreateExampleOptions {
	_options.IncludeAudit = core.BoolPtr(includeAudit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateExampleOptions) SetHeaders(param map[string]string) *CreateExampleOptions {
	options.Headers = param
	return options
}

// CreateIntent : CreateIntent struct
type CreateIntent struct {
	// The name of the intent. This string must conform to the following restrictions:
	// - It can contain only Unicode alphanumeric, underscore, hyphen, and dot characters.
	// - It cannot begin with the reserved prefix `sys-`.
	Intent *string `json:"intent" validate:"required"`

	// The description of the intent. This string cannot contain carriage return, newline, or tab characters.
	Description *string `json:"description,omitempty"`

	// The timestamp for creation of the object.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// The timestamp for the most recent update to the object.
	Updated *strfmt.DateTime `json:"updated,omitempty"`

	// An array of user input examples for the intent.
	Examples []Example `json:"examples,omitempty"`
}

// NewCreateIntent : Instantiate CreateIntent (Generic Model Constructor)
func (*AssistantV1) NewCreateIntent(intent string) (_model *CreateIntent, err error) {
	_model = &CreateIntent{
		Intent: core.StringPtr(intent),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalCreateIntent unmarshals an instance of CreateIntent from the specified map of raw messages.
func UnmarshalCreateIntent(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateIntent)
	err = core.UnmarshalPrimitive(m, "intent", &obj.Intent)
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
	err = core.UnmarshalModel(m, "examples", &obj.Examples, UnmarshalExample)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateIntentOptions : The CreateIntent options.
type CreateIntentOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// The name of the intent. This string must conform to the following restrictions:
	// - It can contain only Unicode alphanumeric, underscore, hyphen, and dot characters.
	// - It cannot begin with the reserved prefix `sys-`.
	Intent *string `json:"intent" validate:"required"`

	// The description of the intent. This string cannot contain carriage return, newline, or tab characters.
	Description *string `json:"description,omitempty"`

	// An array of user input examples for the intent.
	Examples []Example `json:"examples,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateIntentOptions : Instantiate CreateIntentOptions
func (*AssistantV1) NewCreateIntentOptions(workspaceID string, intent string) *CreateIntentOptions {
	return &CreateIntentOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Intent:      core.StringPtr(intent),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *CreateIntentOptions) SetWorkspaceID(workspaceID string) *CreateIntentOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetIntent : Allow user to set Intent
func (_options *CreateIntentOptions) SetIntent(intent string) *CreateIntentOptions {
	_options.Intent = core.StringPtr(intent)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateIntentOptions) SetDescription(description string) *CreateIntentOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetExamples : Allow user to set Examples
func (_options *CreateIntentOptions) SetExamples(examples []Example) *CreateIntentOptions {
	_options.Examples = examples
	return _options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (_options *CreateIntentOptions) SetIncludeAudit(includeAudit bool) *CreateIntentOptions {
	_options.IncludeAudit = core.BoolPtr(includeAudit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateIntentOptions) SetHeaders(param map[string]string) *CreateIntentOptions {
	options.Headers = param
	return options
}

// CreateSynonymOptions : The CreateSynonym options.
type CreateSynonymOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// The name of the entity.
	Entity *string `json:"-" validate:"required,ne="`

	// The text of the entity value.
	Value *string `json:"-" validate:"required,ne="`

	// The text of the synonym. This string must conform to the following restrictions:
	// - It cannot contain carriage return, newline, or tab characters.
	// - It cannot consist of only whitespace characters.
	Synonym *string `json:"synonym" validate:"required"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateSynonymOptions : Instantiate CreateSynonymOptions
func (*AssistantV1) NewCreateSynonymOptions(workspaceID string, entity string, value string, synonym string) *CreateSynonymOptions {
	return &CreateSynonymOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Entity:      core.StringPtr(entity),
		Value:       core.StringPtr(value),
		Synonym:     core.StringPtr(synonym),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *CreateSynonymOptions) SetWorkspaceID(workspaceID string) *CreateSynonymOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetEntity : Allow user to set Entity
func (_options *CreateSynonymOptions) SetEntity(entity string) *CreateSynonymOptions {
	_options.Entity = core.StringPtr(entity)
	return _options
}

// SetValue : Allow user to set Value
func (_options *CreateSynonymOptions) SetValue(value string) *CreateSynonymOptions {
	_options.Value = core.StringPtr(value)
	return _options
}

// SetSynonym : Allow user to set Synonym
func (_options *CreateSynonymOptions) SetSynonym(synonym string) *CreateSynonymOptions {
	_options.Synonym = core.StringPtr(synonym)
	return _options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (_options *CreateSynonymOptions) SetIncludeAudit(includeAudit bool) *CreateSynonymOptions {
	_options.IncludeAudit = core.BoolPtr(includeAudit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateSynonymOptions) SetHeaders(param map[string]string) *CreateSynonymOptions {
	options.Headers = param
	return options
}

// CreateValue : CreateValue struct
type CreateValue struct {
	// The text of the entity value. This string must conform to the following restrictions:
	// - It cannot contain carriage return, newline, or tab characters.
	// - It cannot consist of only whitespace characters.
	Value *string `json:"value" validate:"required"`

	// Any metadata related to the entity value.
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// Specifies the type of entity value.
	Type *string `json:"type,omitempty"`

	// An array of synonyms for the entity value. A value can specify either synonyms or patterns (depending on the value
	// type), but not both. A synonym must conform to the following resrictions:
	// - It cannot contain carriage return, newline, or tab characters.
	// - It cannot consist of only whitespace characters.
	Synonyms []string `json:"synonyms,omitempty"`

	// An array of patterns for the entity value. A value can specify either synonyms or patterns (depending on the value
	// type), but not both. A pattern is a regular expression; for more information about how to specify a pattern, see the
	// [documentation](https://cloud.ibm.com/docs/assistant?topic=assistant-entities#entities-create-dictionary-based).
	Patterns []string `json:"patterns,omitempty"`

	// The timestamp for creation of the object.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// The timestamp for the most recent update to the object.
	Updated *strfmt.DateTime `json:"updated,omitempty"`
}

// Constants associated with the CreateValue.Type property.
// Specifies the type of entity value.
const (
	CreateValueTypePatternsConst = "patterns"
	CreateValueTypeSynonymsConst = "synonyms"
)

// NewCreateValue : Instantiate CreateValue (Generic Model Constructor)
func (*AssistantV1) NewCreateValue(value string) (_model *CreateValue, err error) {
	_model = &CreateValue{
		Value: core.StringPtr(value),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalCreateValue unmarshals an instance of CreateValue from the specified map of raw messages.
func UnmarshalCreateValue(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateValue)
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "metadata", &obj.Metadata)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "synonyms", &obj.Synonyms)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "patterns", &obj.Patterns)
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

// CreateValueOptions : The CreateValue options.
type CreateValueOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// The name of the entity.
	Entity *string `json:"-" validate:"required,ne="`

	// The text of the entity value. This string must conform to the following restrictions:
	// - It cannot contain carriage return, newline, or tab characters.
	// - It cannot consist of only whitespace characters.
	Value *string `json:"value" validate:"required"`

	// Any metadata related to the entity value.
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// Specifies the type of entity value.
	Type *string `json:"type,omitempty"`

	// An array of synonyms for the entity value. A value can specify either synonyms or patterns (depending on the value
	// type), but not both. A synonym must conform to the following resrictions:
	// - It cannot contain carriage return, newline, or tab characters.
	// - It cannot consist of only whitespace characters.
	Synonyms []string `json:"synonyms,omitempty"`

	// An array of patterns for the entity value. A value can specify either synonyms or patterns (depending on the value
	// type), but not both. A pattern is a regular expression; for more information about how to specify a pattern, see the
	// [documentation](https://cloud.ibm.com/docs/assistant?topic=assistant-entities#entities-create-dictionary-based).
	Patterns []string `json:"patterns,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateValueOptions.Type property.
// Specifies the type of entity value.
const (
	CreateValueOptionsTypePatternsConst = "patterns"
	CreateValueOptionsTypeSynonymsConst = "synonyms"
)

// NewCreateValueOptions : Instantiate CreateValueOptions
func (*AssistantV1) NewCreateValueOptions(workspaceID string, entity string, value string) *CreateValueOptions {
	return &CreateValueOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Entity:      core.StringPtr(entity),
		Value:       core.StringPtr(value),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *CreateValueOptions) SetWorkspaceID(workspaceID string) *CreateValueOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetEntity : Allow user to set Entity
func (_options *CreateValueOptions) SetEntity(entity string) *CreateValueOptions {
	_options.Entity = core.StringPtr(entity)
	return _options
}

// SetValue : Allow user to set Value
func (_options *CreateValueOptions) SetValue(value string) *CreateValueOptions {
	_options.Value = core.StringPtr(value)
	return _options
}

// SetMetadata : Allow user to set Metadata
func (_options *CreateValueOptions) SetMetadata(metadata map[string]interface{}) *CreateValueOptions {
	_options.Metadata = metadata
	return _options
}

// SetType : Allow user to set Type
func (_options *CreateValueOptions) SetType(typeVar string) *CreateValueOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetSynonyms : Allow user to set Synonyms
func (_options *CreateValueOptions) SetSynonyms(synonyms []string) *CreateValueOptions {
	_options.Synonyms = synonyms
	return _options
}

// SetPatterns : Allow user to set Patterns
func (_options *CreateValueOptions) SetPatterns(patterns []string) *CreateValueOptions {
	_options.Patterns = patterns
	return _options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (_options *CreateValueOptions) SetIncludeAudit(includeAudit bool) *CreateValueOptions {
	_options.IncludeAudit = core.BoolPtr(includeAudit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateValueOptions) SetHeaders(param map[string]string) *CreateValueOptions {
	options.Headers = param
	return options
}

// CreateWorkspaceOptions : The CreateWorkspace options.
type CreateWorkspaceOptions struct {
	// The name of the workspace. This string cannot contain carriage return, newline, or tab characters.
	Name *string `json:"name,omitempty"`

	// The description of the workspace. This string cannot contain carriage return, newline, or tab characters.
	Description *string `json:"description,omitempty"`

	// The language of the workspace.
	Language *string `json:"language,omitempty"`

	// An array of objects describing the dialog nodes in the workspace.
	DialogNodes []DialogNode `json:"dialog_nodes,omitempty"`

	// An array of objects defining input examples that have been marked as irrelevant input.
	Counterexamples []Counterexample `json:"counterexamples,omitempty"`

	// Any metadata related to the workspace.
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// Whether training data from the workspace (including artifacts such as intents and entities) can be used by IBM for
	// general service improvements. `true` indicates that workspace training data is not to be used.
	LearningOptOut *bool `json:"learning_opt_out,omitempty"`

	// Global settings for the workspace.
	SystemSettings *WorkspaceSystemSettings `json:"system_settings,omitempty"`

	Webhooks []Webhook `json:"webhooks,omitempty"`

	// An array of objects defining the intents for the workspace.
	Intents []CreateIntent `json:"intents,omitempty"`

	// An array of objects describing the entities for the workspace.
	Entities []CreateEntity `json:"entities,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateWorkspaceOptions : Instantiate CreateWorkspaceOptions
func (*AssistantV1) NewCreateWorkspaceOptions() *CreateWorkspaceOptions {
	return &CreateWorkspaceOptions{}
}

// SetName : Allow user to set Name
func (_options *CreateWorkspaceOptions) SetName(name string) *CreateWorkspaceOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateWorkspaceOptions) SetDescription(description string) *CreateWorkspaceOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetLanguage : Allow user to set Language
func (_options *CreateWorkspaceOptions) SetLanguage(language string) *CreateWorkspaceOptions {
	_options.Language = core.StringPtr(language)
	return _options
}

// SetDialogNodes : Allow user to set DialogNodes
func (_options *CreateWorkspaceOptions) SetDialogNodes(dialogNodes []DialogNode) *CreateWorkspaceOptions {
	_options.DialogNodes = dialogNodes
	return _options
}

// SetCounterexamples : Allow user to set Counterexamples
func (_options *CreateWorkspaceOptions) SetCounterexamples(counterexamples []Counterexample) *CreateWorkspaceOptions {
	_options.Counterexamples = counterexamples
	return _options
}

// SetMetadata : Allow user to set Metadata
func (_options *CreateWorkspaceOptions) SetMetadata(metadata map[string]interface{}) *CreateWorkspaceOptions {
	_options.Metadata = metadata
	return _options
}

// SetLearningOptOut : Allow user to set LearningOptOut
func (_options *CreateWorkspaceOptions) SetLearningOptOut(learningOptOut bool) *CreateWorkspaceOptions {
	_options.LearningOptOut = core.BoolPtr(learningOptOut)
	return _options
}

// SetSystemSettings : Allow user to set SystemSettings
func (_options *CreateWorkspaceOptions) SetSystemSettings(systemSettings *WorkspaceSystemSettings) *CreateWorkspaceOptions {
	_options.SystemSettings = systemSettings
	return _options
}

// SetWebhooks : Allow user to set Webhooks
func (_options *CreateWorkspaceOptions) SetWebhooks(webhooks []Webhook) *CreateWorkspaceOptions {
	_options.Webhooks = webhooks
	return _options
}

// SetIntents : Allow user to set Intents
func (_options *CreateWorkspaceOptions) SetIntents(intents []CreateIntent) *CreateWorkspaceOptions {
	_options.Intents = intents
	return _options
}

// SetEntities : Allow user to set Entities
func (_options *CreateWorkspaceOptions) SetEntities(entities []CreateEntity) *CreateWorkspaceOptions {
	_options.Entities = entities
	return _options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (_options *CreateWorkspaceOptions) SetIncludeAudit(includeAudit bool) *CreateWorkspaceOptions {
	_options.IncludeAudit = core.BoolPtr(includeAudit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateWorkspaceOptions) SetHeaders(param map[string]string) *CreateWorkspaceOptions {
	options.Headers = param
	return options
}

// DeleteCounterexampleOptions : The DeleteCounterexample options.
type DeleteCounterexampleOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// The text of a user input counterexample (for example, `What are you wearing?`).
	Text *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteCounterexampleOptions : Instantiate DeleteCounterexampleOptions
func (*AssistantV1) NewDeleteCounterexampleOptions(workspaceID string, text string) *DeleteCounterexampleOptions {
	return &DeleteCounterexampleOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Text:        core.StringPtr(text),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *DeleteCounterexampleOptions) SetWorkspaceID(workspaceID string) *DeleteCounterexampleOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetText : Allow user to set Text
func (_options *DeleteCounterexampleOptions) SetText(text string) *DeleteCounterexampleOptions {
	_options.Text = core.StringPtr(text)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteCounterexampleOptions) SetHeaders(param map[string]string) *DeleteCounterexampleOptions {
	options.Headers = param
	return options
}

// DeleteDialogNodeOptions : The DeleteDialogNode options.
type DeleteDialogNodeOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// The dialog node ID (for example, `node_1_1479323581900`).
	DialogNode *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteDialogNodeOptions : Instantiate DeleteDialogNodeOptions
func (*AssistantV1) NewDeleteDialogNodeOptions(workspaceID string, dialogNode string) *DeleteDialogNodeOptions {
	return &DeleteDialogNodeOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		DialogNode:  core.StringPtr(dialogNode),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *DeleteDialogNodeOptions) SetWorkspaceID(workspaceID string) *DeleteDialogNodeOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetDialogNode : Allow user to set DialogNode
func (_options *DeleteDialogNodeOptions) SetDialogNode(dialogNode string) *DeleteDialogNodeOptions {
	_options.DialogNode = core.StringPtr(dialogNode)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteDialogNodeOptions) SetHeaders(param map[string]string) *DeleteDialogNodeOptions {
	options.Headers = param
	return options
}

// DeleteEntityOptions : The DeleteEntity options.
type DeleteEntityOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// The name of the entity.
	Entity *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteEntityOptions : Instantiate DeleteEntityOptions
func (*AssistantV1) NewDeleteEntityOptions(workspaceID string, entity string) *DeleteEntityOptions {
	return &DeleteEntityOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Entity:      core.StringPtr(entity),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *DeleteEntityOptions) SetWorkspaceID(workspaceID string) *DeleteEntityOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetEntity : Allow user to set Entity
func (_options *DeleteEntityOptions) SetEntity(entity string) *DeleteEntityOptions {
	_options.Entity = core.StringPtr(entity)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteEntityOptions) SetHeaders(param map[string]string) *DeleteEntityOptions {
	options.Headers = param
	return options
}

// DeleteExampleOptions : The DeleteExample options.
type DeleteExampleOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// The intent name.
	Intent *string `json:"-" validate:"required,ne="`

	// The text of the user input example.
	Text *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteExampleOptions : Instantiate DeleteExampleOptions
func (*AssistantV1) NewDeleteExampleOptions(workspaceID string, intent string, text string) *DeleteExampleOptions {
	return &DeleteExampleOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Intent:      core.StringPtr(intent),
		Text:        core.StringPtr(text),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *DeleteExampleOptions) SetWorkspaceID(workspaceID string) *DeleteExampleOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetIntent : Allow user to set Intent
func (_options *DeleteExampleOptions) SetIntent(intent string) *DeleteExampleOptions {
	_options.Intent = core.StringPtr(intent)
	return _options
}

// SetText : Allow user to set Text
func (_options *DeleteExampleOptions) SetText(text string) *DeleteExampleOptions {
	_options.Text = core.StringPtr(text)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteExampleOptions) SetHeaders(param map[string]string) *DeleteExampleOptions {
	options.Headers = param
	return options
}

// DeleteIntentOptions : The DeleteIntent options.
type DeleteIntentOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// The intent name.
	Intent *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteIntentOptions : Instantiate DeleteIntentOptions
func (*AssistantV1) NewDeleteIntentOptions(workspaceID string, intent string) *DeleteIntentOptions {
	return &DeleteIntentOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Intent:      core.StringPtr(intent),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *DeleteIntentOptions) SetWorkspaceID(workspaceID string) *DeleteIntentOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetIntent : Allow user to set Intent
func (_options *DeleteIntentOptions) SetIntent(intent string) *DeleteIntentOptions {
	_options.Intent = core.StringPtr(intent)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteIntentOptions) SetHeaders(param map[string]string) *DeleteIntentOptions {
	options.Headers = param
	return options
}

// DeleteSynonymOptions : The DeleteSynonym options.
type DeleteSynonymOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// The name of the entity.
	Entity *string `json:"-" validate:"required,ne="`

	// The text of the entity value.
	Value *string `json:"-" validate:"required,ne="`

	// The text of the synonym.
	Synonym *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteSynonymOptions : Instantiate DeleteSynonymOptions
func (*AssistantV1) NewDeleteSynonymOptions(workspaceID string, entity string, value string, synonym string) *DeleteSynonymOptions {
	return &DeleteSynonymOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Entity:      core.StringPtr(entity),
		Value:       core.StringPtr(value),
		Synonym:     core.StringPtr(synonym),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *DeleteSynonymOptions) SetWorkspaceID(workspaceID string) *DeleteSynonymOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetEntity : Allow user to set Entity
func (_options *DeleteSynonymOptions) SetEntity(entity string) *DeleteSynonymOptions {
	_options.Entity = core.StringPtr(entity)
	return _options
}

// SetValue : Allow user to set Value
func (_options *DeleteSynonymOptions) SetValue(value string) *DeleteSynonymOptions {
	_options.Value = core.StringPtr(value)
	return _options
}

// SetSynonym : Allow user to set Synonym
func (_options *DeleteSynonymOptions) SetSynonym(synonym string) *DeleteSynonymOptions {
	_options.Synonym = core.StringPtr(synonym)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteSynonymOptions) SetHeaders(param map[string]string) *DeleteSynonymOptions {
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
func (*AssistantV1) NewDeleteUserDataOptions(customerID string) *DeleteUserDataOptions {
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

// DeleteValueOptions : The DeleteValue options.
type DeleteValueOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// The name of the entity.
	Entity *string `json:"-" validate:"required,ne="`

	// The text of the entity value.
	Value *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteValueOptions : Instantiate DeleteValueOptions
func (*AssistantV1) NewDeleteValueOptions(workspaceID string, entity string, value string) *DeleteValueOptions {
	return &DeleteValueOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Entity:      core.StringPtr(entity),
		Value:       core.StringPtr(value),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *DeleteValueOptions) SetWorkspaceID(workspaceID string) *DeleteValueOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetEntity : Allow user to set Entity
func (_options *DeleteValueOptions) SetEntity(entity string) *DeleteValueOptions {
	_options.Entity = core.StringPtr(entity)
	return _options
}

// SetValue : Allow user to set Value
func (_options *DeleteValueOptions) SetValue(value string) *DeleteValueOptions {
	_options.Value = core.StringPtr(value)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteValueOptions) SetHeaders(param map[string]string) *DeleteValueOptions {
	options.Headers = param
	return options
}

// DeleteWorkspaceOptions : The DeleteWorkspace options.
type DeleteWorkspaceOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteWorkspaceOptions : Instantiate DeleteWorkspaceOptions
func (*AssistantV1) NewDeleteWorkspaceOptions(workspaceID string) *DeleteWorkspaceOptions {
	return &DeleteWorkspaceOptions{
		WorkspaceID: core.StringPtr(workspaceID),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *DeleteWorkspaceOptions) SetWorkspaceID(workspaceID string) *DeleteWorkspaceOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteWorkspaceOptions) SetHeaders(param map[string]string) *DeleteWorkspaceOptions {
	options.Headers = param
	return options
}

// DialogNode : DialogNode struct
type DialogNode struct {
	// The unique ID of the dialog node. This is an internal identifier used to refer to the dialog node from other dialog
	// nodes and in the diagnostic information included with message responses.
	//
	// This string can contain only Unicode alphanumeric, space, underscore, hyphen, and dot characters.
	DialogNode *string `json:"dialog_node" validate:"required"`

	// The description of the dialog node. This string cannot contain carriage return, newline, or tab characters.
	Description *string `json:"description,omitempty"`

	// The condition that will trigger the dialog node. This string cannot contain carriage return, newline, or tab
	// characters.
	Conditions *string `json:"conditions,omitempty"`

	// The unique ID of the parent dialog node. This property is omitted if the dialog node has no parent.
	Parent *string `json:"parent,omitempty"`

	// The unique ID of the previous sibling dialog node. This property is omitted if the dialog node has no previous
	// sibling.
	PreviousSibling *string `json:"previous_sibling,omitempty"`

	// The output of the dialog node. For more information about how to specify dialog node output, see the
	// [documentation](https://cloud.ibm.com/docs/assistant?topic=assistant-dialog-overview#dialog-overview-responses).
	Output *DialogNodeOutput `json:"output,omitempty"`

	// The context for the dialog node.
	Context *DialogNodeContext `json:"context,omitempty"`

	// The metadata for the dialog node.
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// The next step to execute following this dialog node.
	NextStep *DialogNodeNextStep `json:"next_step,omitempty"`

	// A human-readable name for the dialog node. If the node is included in disambiguation, this title is used to populate
	// the **label** property of the corresponding suggestion in the `suggestion` response type (unless it is overridden by
	// the **user_label** property). The title is also used to populate the **topic** property in the `connect_to_agent`
	// response type.
	//
	// This string can contain only Unicode alphanumeric, space, underscore, hyphen, and dot characters.
	Title *string `json:"title,omitempty"`

	// How the dialog node is processed.
	Type *string `json:"type,omitempty"`

	// How an `event_handler` node is processed.
	EventName *string `json:"event_name,omitempty"`

	// The location in the dialog context where output is stored.
	Variable *string `json:"variable,omitempty"`

	// An array of objects describing any actions to be invoked by the dialog node.
	Actions []DialogNodeAction `json:"actions,omitempty"`

	// Whether this top-level dialog node can be digressed into.
	DigressIn *string `json:"digress_in,omitempty"`

	// Whether this dialog node can be returned to after a digression.
	DigressOut *string `json:"digress_out,omitempty"`

	// Whether the user can digress to top-level nodes while filling out slots.
	DigressOutSlots *string `json:"digress_out_slots,omitempty"`

	// A label that can be displayed externally to describe the purpose of the node to users. If set, this label is used to
	// identify the node in disambiguation responses (overriding the value of the **title** property).
	UserLabel *string `json:"user_label,omitempty"`

	// Whether the dialog node should be excluded from disambiguation suggestions. Valid only when **type**=`standard` or
	// `frame`.
	DisambiguationOptOut *bool `json:"disambiguation_opt_out,omitempty"`

	// For internal use only.
	Disabled *bool `json:"disabled,omitempty"`

	// The timestamp for creation of the object.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// The timestamp for the most recent update to the object.
	Updated *strfmt.DateTime `json:"updated,omitempty"`
}

// Constants associated with the DialogNode.Type property.
// How the dialog node is processed.
const (
	DialogNodeTypeEventHandlerConst      = "event_handler"
	DialogNodeTypeFolderConst            = "folder"
	DialogNodeTypeFrameConst             = "frame"
	DialogNodeTypeResponseConditionConst = "response_condition"
	DialogNodeTypeSlotConst              = "slot"
	DialogNodeTypeStandardConst          = "standard"
)

// Constants associated with the DialogNode.EventName property.
// How an `event_handler` node is processed.
const (
	DialogNodeEventNameDigressionReturnPromptConst   = "digression_return_prompt"
	DialogNodeEventNameFilledConst                   = "filled"
	DialogNodeEventNameFilledMultipleConst           = "filled_multiple"
	DialogNodeEventNameFocusConst                    = "focus"
	DialogNodeEventNameGenericConst                  = "generic"
	DialogNodeEventNameInputConst                    = "input"
	DialogNodeEventNameNomatchConst                  = "nomatch"
	DialogNodeEventNameNomatchResponsesDepletedConst = "nomatch_responses_depleted"
	DialogNodeEventNameValidateConst                 = "validate"
)

// Constants associated with the DialogNode.DigressIn property.
// Whether this top-level dialog node can be digressed into.
const (
	DialogNodeDigressInDoesNotReturnConst = "does_not_return"
	DialogNodeDigressInNotAvailableConst  = "not_available"
	DialogNodeDigressInReturnsConst       = "returns"
)

// Constants associated with the DialogNode.DigressOut property.
// Whether this dialog node can be returned to after a digression.
const (
	DialogNodeDigressOutAllowAllConst            = "allow_all"
	DialogNodeDigressOutAllowAllNeverReturnConst = "allow_all_never_return"
	DialogNodeDigressOutAllowReturningConst      = "allow_returning"
)

// Constants associated with the DialogNode.DigressOutSlots property.
// Whether the user can digress to top-level nodes while filling out slots.
const (
	DialogNodeDigressOutSlotsAllowAllConst       = "allow_all"
	DialogNodeDigressOutSlotsAllowReturningConst = "allow_returning"
	DialogNodeDigressOutSlotsNotAllowedConst     = "not_allowed"
)

// NewDialogNode : Instantiate DialogNode (Generic Model Constructor)
func (*AssistantV1) NewDialogNode(dialogNode string) (_model *DialogNode, err error) {
	_model = &DialogNode{
		DialogNode: core.StringPtr(dialogNode),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalDialogNode unmarshals an instance of DialogNode from the specified map of raw messages.
func UnmarshalDialogNode(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DialogNode)
	err = core.UnmarshalPrimitive(m, "dialog_node", &obj.DialogNode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "conditions", &obj.Conditions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parent", &obj.Parent)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "previous_sibling", &obj.PreviousSibling)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "output", &obj.Output, UnmarshalDialogNodeOutput)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "context", &obj.Context, UnmarshalDialogNodeContext)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "metadata", &obj.Metadata)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next_step", &obj.NextStep, UnmarshalDialogNodeNextStep)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "title", &obj.Title)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "event_name", &obj.EventName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "variable", &obj.Variable)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "actions", &obj.Actions, UnmarshalDialogNodeAction)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "digress_in", &obj.DigressIn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "digress_out", &obj.DigressOut)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "digress_out_slots", &obj.DigressOutSlots)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "user_label", &obj.UserLabel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "disambiguation_opt_out", &obj.DisambiguationOptOut)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "disabled", &obj.Disabled)
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

// DialogNodeAction : DialogNodeAction struct
type DialogNodeAction struct {
	// The name of the action.
	Name *string `json:"name" validate:"required"`

	// The type of action to invoke.
	Type *string `json:"type,omitempty"`

	// A map of key/value pairs to be provided to the action.
	Parameters map[string]interface{} `json:"parameters,omitempty"`

	// The location in the dialog context where the result of the action is stored.
	ResultVariable *string `json:"result_variable" validate:"required"`

	// The name of the context variable that the client application will use to pass in credentials for the action.
	Credentials *string `json:"credentials,omitempty"`
}

// Constants associated with the DialogNodeAction.Type property.
// The type of action to invoke.
const (
	DialogNodeActionTypeClientConst        = "client"
	DialogNodeActionTypeCloudFunctionConst = "cloud_function"
	DialogNodeActionTypeServerConst        = "server"
	DialogNodeActionTypeWebActionConst     = "web_action"
	DialogNodeActionTypeWebhookConst       = "webhook"
)

// NewDialogNodeAction : Instantiate DialogNodeAction (Generic Model Constructor)
func (*AssistantV1) NewDialogNodeAction(name string, resultVariable string) (_model *DialogNodeAction, err error) {
	_model = &DialogNodeAction{
		Name:           core.StringPtr(name),
		ResultVariable: core.StringPtr(resultVariable),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalDialogNodeAction unmarshals an instance of DialogNodeAction from the specified map of raw messages.
func UnmarshalDialogNodeAction(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DialogNodeAction)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parameters", &obj.Parameters)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "result_variable", &obj.ResultVariable)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "credentials", &obj.Credentials)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DialogNodeCollection : An array of dialog nodes.
type DialogNodeCollection struct {
	// An array of objects describing the dialog nodes defined for the workspace.
	DialogNodes []DialogNode `json:"dialog_nodes" validate:"required"`

	// The pagination data for the returned objects.
	Pagination *Pagination `json:"pagination" validate:"required"`
}

// UnmarshalDialogNodeCollection unmarshals an instance of DialogNodeCollection from the specified map of raw messages.
func UnmarshalDialogNodeCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DialogNodeCollection)
	err = core.UnmarshalModel(m, "dialog_nodes", &obj.DialogNodes, UnmarshalDialogNode)
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

// DialogNodeContext : The context for the dialog node.
type DialogNodeContext struct {
	// Context data intended for specific integrations.
	Integrations map[string]map[string]interface{} `json:"integrations,omitempty"`

	// Allows users to set arbitrary properties
	additionalProperties map[string]interface{}
}

// SetProperty allows the user to set an arbitrary property on an instance of DialogNodeContext
func (o *DialogNodeContext) SetProperty(key string, value interface{}) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string]interface{})
	}
	o.additionalProperties[key] = value
}

// SetProperties allows the user to set a map of arbitrary properties on an instance of DialogNodeContext
func (o *DialogNodeContext) SetProperties(m map[string]interface{}) {
	o.additionalProperties = make(map[string]interface{})
	for k, v := range m {
		o.additionalProperties[k] = v
	}
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of DialogNodeContext
func (o *DialogNodeContext) GetProperty(key string) interface{} {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of DialogNodeContext
func (o *DialogNodeContext) GetProperties() map[string]interface{} {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of DialogNodeContext
func (o *DialogNodeContext) MarshalJSON() (buffer []byte, err error) {
	m := make(map[string]interface{})
	if len(o.additionalProperties) > 0 {
		for k, v := range o.additionalProperties {
			m[k] = v
		}
	}
	if o.Integrations != nil {
		m["integrations"] = o.Integrations
	}
	buffer, err = json.Marshal(m)
	return
}

// UnmarshalDialogNodeContext unmarshals an instance of DialogNodeContext from the specified map of raw messages.
func UnmarshalDialogNodeContext(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DialogNodeContext)
	err = core.UnmarshalPrimitive(m, "integrations", &obj.Integrations)
	if err != nil {
		return
	}
	delete(m, "integrations")
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

// DialogNodeNextStep : The next step to execute following this dialog node.
type DialogNodeNextStep struct {
	// What happens after the dialog node completes. The valid values depend on the node type:
	// - The following values are valid for any node:
	//   - `get_user_input`
	//   - `skip_user_input`
	//   - `jump_to`
	// - If the node is of type `event_handler` and its parent node is of type `slot` or `frame`, additional values are
	// also valid:
	//   - if **event_name**=`filled` and the type of the parent node is `slot`:
	//     - `reprompt`
	//     - `skip_all_slots`
	// - if **event_name**=`nomatch` and the type of the parent node is `slot`:
	//     - `reprompt`
	//     - `skip_slot`
	//     - `skip_all_slots`
	// - if **event_name**=`generic` and the type of the parent node is `frame`:
	//     - `reprompt`
	//     - `skip_slot`
	//     - `skip_all_slots`
	//      If you specify `jump_to`, then you must also specify a value for the `dialog_node` property.
	Behavior *string `json:"behavior" validate:"required"`

	// The unique ID of the dialog node to process next. This parameter is required if **behavior**=`jump_to`.
	DialogNode *string `json:"dialog_node,omitempty"`

	// Which part of the dialog node to process next.
	Selector *string `json:"selector,omitempty"`
}

// Constants associated with the DialogNodeNextStep.Behavior property.
// What happens after the dialog node completes. The valid values depend on the node type:
// - The following values are valid for any node:
//   - `get_user_input`
//   - `skip_user_input`
//   - `jump_to`
// - If the node is of type `event_handler` and its parent node is of type `slot` or `frame`, additional values are also
// valid:
//   - if **event_name**=`filled` and the type of the parent node is `slot`:
//     - `reprompt`
//     - `skip_all_slots`
// - if **event_name**=`nomatch` and the type of the parent node is `slot`:
//     - `reprompt`
//     - `skip_slot`
//     - `skip_all_slots`
// - if **event_name**=`generic` and the type of the parent node is `frame`:
//     - `reprompt`
//     - `skip_slot`
//     - `skip_all_slots`
//      If you specify `jump_to`, then you must also specify a value for the `dialog_node` property.
const (
	DialogNodeNextStepBehaviorGetUserInputConst  = "get_user_input"
	DialogNodeNextStepBehaviorJumpToConst        = "jump_to"
	DialogNodeNextStepBehaviorRepromptConst      = "reprompt"
	DialogNodeNextStepBehaviorSkipAllSlotsConst  = "skip_all_slots"
	DialogNodeNextStepBehaviorSkipSlotConst      = "skip_slot"
	DialogNodeNextStepBehaviorSkipUserInputConst = "skip_user_input"
)

// Constants associated with the DialogNodeNextStep.Selector property.
// Which part of the dialog node to process next.
const (
	DialogNodeNextStepSelectorBodyConst      = "body"
	DialogNodeNextStepSelectorClientConst    = "client"
	DialogNodeNextStepSelectorConditionConst = "condition"
	DialogNodeNextStepSelectorUserInputConst = "user_input"
)

// NewDialogNodeNextStep : Instantiate DialogNodeNextStep (Generic Model Constructor)
func (*AssistantV1) NewDialogNodeNextStep(behavior string) (_model *DialogNodeNextStep, err error) {
	_model = &DialogNodeNextStep{
		Behavior: core.StringPtr(behavior),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalDialogNodeNextStep unmarshals an instance of DialogNodeNextStep from the specified map of raw messages.
func UnmarshalDialogNodeNextStep(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DialogNodeNextStep)
	err = core.UnmarshalPrimitive(m, "behavior", &obj.Behavior)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "dialog_node", &obj.DialogNode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "selector", &obj.Selector)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DialogNodeOutput : The output of the dialog node. For more information about how to specify dialog node output, see the
// [documentation](https://cloud.ibm.com/docs/assistant?topic=assistant-dialog-overview#dialog-overview-responses).
type DialogNodeOutput struct {
	// An array of objects describing the output defined for the dialog node.
	Generic []DialogNodeOutputGenericIntf `json:"generic,omitempty"`

	// Output intended for specific integrations. For more information, see the
	// [documentation](https://cloud.ibm.com/docs/assistant?topic=assistant-dialog-responses-json).
	Integrations map[string]map[string]interface{} `json:"integrations,omitempty"`

	// Options that modify how specified output is handled.
	Modifiers *DialogNodeOutputModifiers `json:"modifiers,omitempty"`

	// Allows users to set arbitrary properties
	additionalProperties map[string]interface{}
}

// SetProperty allows the user to set an arbitrary property on an instance of DialogNodeOutput
func (o *DialogNodeOutput) SetProperty(key string, value interface{}) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string]interface{})
	}
	o.additionalProperties[key] = value
}

// SetProperties allows the user to set a map of arbitrary properties on an instance of DialogNodeOutput
func (o *DialogNodeOutput) SetProperties(m map[string]interface{}) {
	o.additionalProperties = make(map[string]interface{})
	for k, v := range m {
		o.additionalProperties[k] = v
	}
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of DialogNodeOutput
func (o *DialogNodeOutput) GetProperty(key string) interface{} {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of DialogNodeOutput
func (o *DialogNodeOutput) GetProperties() map[string]interface{} {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of DialogNodeOutput
func (o *DialogNodeOutput) MarshalJSON() (buffer []byte, err error) {
	m := make(map[string]interface{})
	if len(o.additionalProperties) > 0 {
		for k, v := range o.additionalProperties {
			m[k] = v
		}
	}
	if o.Generic != nil {
		m["generic"] = o.Generic
	}
	if o.Integrations != nil {
		m["integrations"] = o.Integrations
	}
	if o.Modifiers != nil {
		m["modifiers"] = o.Modifiers
	}
	buffer, err = json.Marshal(m)
	return
}

// UnmarshalDialogNodeOutput unmarshals an instance of DialogNodeOutput from the specified map of raw messages.
func UnmarshalDialogNodeOutput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DialogNodeOutput)
	err = core.UnmarshalModel(m, "generic", &obj.Generic, UnmarshalDialogNodeOutputGeneric)
	if err != nil {
		return
	}
	delete(m, "generic")
	err = core.UnmarshalPrimitive(m, "integrations", &obj.Integrations)
	if err != nil {
		return
	}
	delete(m, "integrations")
	err = core.UnmarshalModel(m, "modifiers", &obj.Modifiers, UnmarshalDialogNodeOutputModifiers)
	if err != nil {
		return
	}
	delete(m, "modifiers")
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

// DialogNodeOutputConnectToAgentTransferInfo : Routing or other contextual information to be used by target service desk systems.
type DialogNodeOutputConnectToAgentTransferInfo struct {
	Target map[string]map[string]interface{} `json:"target,omitempty"`
}

// UnmarshalDialogNodeOutputConnectToAgentTransferInfo unmarshals an instance of DialogNodeOutputConnectToAgentTransferInfo from the specified map of raw messages.
func UnmarshalDialogNodeOutputConnectToAgentTransferInfo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DialogNodeOutputConnectToAgentTransferInfo)
	err = core.UnmarshalPrimitive(m, "target", &obj.Target)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DialogNodeOutputGeneric : DialogNodeOutputGeneric struct
// Models which "extend" this model:
// - DialogNodeOutputGenericDialogNodeOutputResponseTypeText
// - DialogNodeOutputGenericDialogNodeOutputResponseTypePause
// - DialogNodeOutputGenericDialogNodeOutputResponseTypeImage
// - DialogNodeOutputGenericDialogNodeOutputResponseTypeOption
// - DialogNodeOutputGenericDialogNodeOutputResponseTypeConnectToAgent
// - DialogNodeOutputGenericDialogNodeOutputResponseTypeSearchSkill
// - DialogNodeOutputGenericDialogNodeOutputResponseTypeChannelTransfer
// - DialogNodeOutputGenericDialogNodeOutputResponseTypeUserDefined
type DialogNodeOutputGeneric struct {
	// The type of response returned by the dialog node. The specified response type must be supported by the client
	// application or channel.
	ResponseType *string `json:"response_type,omitempty"`

	// A list of one or more objects defining text responses.
	Values []DialogNodeOutputTextValuesElement `json:"values,omitempty"`

	// How a response is selected from the list, if more than one response is specified.
	SelectionPolicy *string `json:"selection_policy,omitempty"`

	// The delimiter to use as a separator between responses when `selection_policy`=`multiline`.
	Delimiter *string `json:"delimiter,omitempty"`

	// An array of objects specifying channels for which the response is intended.
	Channels []ResponseGenericChannel `json:"channels,omitempty"`

	// How long to pause, in milliseconds. The valid values are from 0 to 10000.
	Time *int64 `json:"time,omitempty"`

	// Whether to send a "user is typing" event during the pause. Ignored if the channel does not support this event.
	Typing *bool `json:"typing,omitempty"`

	// The `https:` URL of the image.
	Source *string `json:"source,omitempty"`

	// An optional title to show before the response.
	Title *string `json:"title,omitempty"`

	// An optional description to show with the response.
	Description *string `json:"description,omitempty"`

	// Descriptive text that can be used for screen readers or other situations where the image cannot be seen.
	AltText *string `json:"alt_text,omitempty"`

	// The preferred type of control to display, if supported by the channel.
	Preference *string `json:"preference,omitempty"`

	// An array of objects describing the options from which the user can choose. You can include up to 20 options.
	Options []DialogNodeOutputOptionsElement `json:"options,omitempty"`

	// An optional message to be sent to the human agent who will be taking over the conversation.
	MessageToHumanAgent *string `json:"message_to_human_agent,omitempty"`

	// An optional message to be displayed to the user to indicate that the conversation will be transferred to the next
	// available agent.
	AgentAvailable *AgentAvailabilityMessage `json:"agent_available,omitempty"`

	// An optional message to be displayed to the user to indicate that no online agent is available to take over the
	// conversation.
	AgentUnavailable *AgentAvailabilityMessage `json:"agent_unavailable,omitempty"`

	// Routing or other contextual information to be used by target service desk systems.
	TransferInfo *DialogNodeOutputConnectToAgentTransferInfo `json:"transfer_info,omitempty"`

	// The text of the search query. This can be either a natural-language query or a query that uses the Discovery query
	// language syntax, depending on the value of the **query_type** property. For more information, see the [Discovery
	// service documentation](https://cloud.ibm.com/docs/discovery?topic=discovery-query-operators#query-operators).
	Query *string `json:"query,omitempty"`

	// The type of the search query.
	QueryType *string `json:"query_type,omitempty"`

	// An optional filter that narrows the set of documents to be searched. For more information, see the [Discovery
	// service documentation]([Discovery service
	// documentation](https://cloud.ibm.com/docs/discovery?topic=discovery-query-parameters#filter).
	Filter *string `json:"filter,omitempty"`

	// The version of the Discovery service API to use for the query.
	DiscoveryVersion *string `json:"discovery_version,omitempty"`

	// The message to display to the user when initiating a channel transfer.
	MessageToUser *string `json:"message_to_user,omitempty"`

	// An object containing any properties for the user-defined response type. The total size of this object cannot exceed
	// 5000 bytes.
	UserDefined map[string]interface{} `json:"user_defined,omitempty"`
}

// Constants associated with the DialogNodeOutputGeneric.SelectionPolicy property.
// How a response is selected from the list, if more than one response is specified.
const (
	DialogNodeOutputGenericSelectionPolicyMultilineConst  = "multiline"
	DialogNodeOutputGenericSelectionPolicyRandomConst     = "random"
	DialogNodeOutputGenericSelectionPolicySequentialConst = "sequential"
)

// Constants associated with the DialogNodeOutputGeneric.Preference property.
// The preferred type of control to display, if supported by the channel.
const (
	DialogNodeOutputGenericPreferenceButtonConst   = "button"
	DialogNodeOutputGenericPreferenceDropdownConst = "dropdown"
)

// Constants associated with the DialogNodeOutputGeneric.QueryType property.
// The type of the search query.
const (
	DialogNodeOutputGenericQueryTypeDiscoveryQueryLanguageConst = "discovery_query_language"
	DialogNodeOutputGenericQueryTypeNaturalLanguageConst        = "natural_language"
)

func (*DialogNodeOutputGeneric) isaDialogNodeOutputGeneric() bool {
	return true
}

type DialogNodeOutputGenericIntf interface {
	isaDialogNodeOutputGeneric() bool
}

// UnmarshalDialogNodeOutputGeneric unmarshals an instance of DialogNodeOutputGeneric from the specified map of raw messages.
func UnmarshalDialogNodeOutputGeneric(m map[string]json.RawMessage, result interface{}) (err error) {
	// Retrieve discriminator value to determine correct "subclass".
	var discValue string
	err = core.UnmarshalPrimitive(m, "response_type", &discValue)
	if err != nil {
		err = fmt.Errorf("error unmarshalling discriminator property 'response_type': %s", err.Error())
		return
	}
	if discValue == "" {
		err = fmt.Errorf("required discriminator property 'response_type' not found in JSON object")
		return
	}
	if discValue == "channel_transfer" {
		err = core.UnmarshalModel(m, "", result, UnmarshalDialogNodeOutputGenericDialogNodeOutputResponseTypeChannelTransfer)
	} else if discValue == "connect_to_agent" {
		err = core.UnmarshalModel(m, "", result, UnmarshalDialogNodeOutputGenericDialogNodeOutputResponseTypeConnectToAgent)
	} else if discValue == "image" {
		err = core.UnmarshalModel(m, "", result, UnmarshalDialogNodeOutputGenericDialogNodeOutputResponseTypeImage)
	} else if discValue == "option" {
		err = core.UnmarshalModel(m, "", result, UnmarshalDialogNodeOutputGenericDialogNodeOutputResponseTypeOption)
	} else if discValue == "pause" {
		err = core.UnmarshalModel(m, "", result, UnmarshalDialogNodeOutputGenericDialogNodeOutputResponseTypePause)
	} else if discValue == "search_skill" {
		err = core.UnmarshalModel(m, "", result, UnmarshalDialogNodeOutputGenericDialogNodeOutputResponseTypeSearchSkill)
	} else if discValue == "text" {
		err = core.UnmarshalModel(m, "", result, UnmarshalDialogNodeOutputGenericDialogNodeOutputResponseTypeText)
	} else if discValue == "user_defined" {
		err = core.UnmarshalModel(m, "", result, UnmarshalDialogNodeOutputGenericDialogNodeOutputResponseTypeUserDefined)
	} else {
		err = fmt.Errorf("unrecognized value for discriminator property 'response_type': %s", discValue)
	}
	return
}

// DialogNodeOutputModifiers : Options that modify how specified output is handled.
type DialogNodeOutputModifiers struct {
	// Whether values in the output will overwrite output values in an array specified by previously executed dialog nodes.
	// If this option is set to `false`, new values will be appended to previously specified values.
	Overwrite *bool `json:"overwrite,omitempty"`
}

// UnmarshalDialogNodeOutputModifiers unmarshals an instance of DialogNodeOutputModifiers from the specified map of raw messages.
func UnmarshalDialogNodeOutputModifiers(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DialogNodeOutputModifiers)
	err = core.UnmarshalPrimitive(m, "overwrite", &obj.Overwrite)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DialogNodeOutputOptionsElement : DialogNodeOutputOptionsElement struct
type DialogNodeOutputOptionsElement struct {
	// The user-facing label for the option.
	Label *string `json:"label" validate:"required"`

	// An object defining the message input to be sent to the Watson Assistant service if the user selects the
	// corresponding option.
	Value *DialogNodeOutputOptionsElementValue `json:"value" validate:"required"`
}

// NewDialogNodeOutputOptionsElement : Instantiate DialogNodeOutputOptionsElement (Generic Model Constructor)
func (*AssistantV1) NewDialogNodeOutputOptionsElement(label string, value *DialogNodeOutputOptionsElementValue) (_model *DialogNodeOutputOptionsElement, err error) {
	_model = &DialogNodeOutputOptionsElement{
		Label: core.StringPtr(label),
		Value: value,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalDialogNodeOutputOptionsElement unmarshals an instance of DialogNodeOutputOptionsElement from the specified map of raw messages.
func UnmarshalDialogNodeOutputOptionsElement(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DialogNodeOutputOptionsElement)
	err = core.UnmarshalPrimitive(m, "label", &obj.Label)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "value", &obj.Value, UnmarshalDialogNodeOutputOptionsElementValue)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DialogNodeOutputOptionsElementValue : An object defining the message input to be sent to the Watson Assistant service if the user selects the corresponding
// option.
type DialogNodeOutputOptionsElementValue struct {
	// An input object that includes the input text.
	Input *MessageInput `json:"input,omitempty"`

	// An array of intents to be used while processing the input.
	//
	// **Note:** This property is supported for backward compatibility with applications that use the v1 **Get response to
	// user input** method.
	Intents []RuntimeIntent `json:"intents,omitempty"`

	// An array of entities to be used while processing the user input.
	//
	// **Note:** This property is supported for backward compatibility with applications that use the v1 **Get response to
	// user input** method.
	Entities []RuntimeEntity `json:"entities,omitempty"`
}

// UnmarshalDialogNodeOutputOptionsElementValue unmarshals an instance of DialogNodeOutputOptionsElementValue from the specified map of raw messages.
func UnmarshalDialogNodeOutputOptionsElementValue(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DialogNodeOutputOptionsElementValue)
	err = core.UnmarshalModel(m, "input", &obj.Input, UnmarshalMessageInput)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "intents", &obj.Intents, UnmarshalRuntimeIntent)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "entities", &obj.Entities, UnmarshalRuntimeEntity)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DialogNodeOutputTextValuesElement : DialogNodeOutputTextValuesElement struct
type DialogNodeOutputTextValuesElement struct {
	// The text of a response. This string can include newline characters (`\n`), Markdown tagging, or other special
	// characters, if supported by the channel.
	Text *string `json:"text,omitempty"`
}

// UnmarshalDialogNodeOutputTextValuesElement unmarshals an instance of DialogNodeOutputTextValuesElement from the specified map of raw messages.
func UnmarshalDialogNodeOutputTextValuesElement(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DialogNodeOutputTextValuesElement)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DialogNodeVisitedDetails : DialogNodeVisitedDetails struct
type DialogNodeVisitedDetails struct {
	// The unique ID of a dialog node that was triggered during processing of the input message.
	DialogNode *string `json:"dialog_node,omitempty"`

	// The title of the dialog node.
	Title *string `json:"title,omitempty"`

	// The conditions that trigger the dialog node.
	Conditions *string `json:"conditions,omitempty"`
}

// UnmarshalDialogNodeVisitedDetails unmarshals an instance of DialogNodeVisitedDetails from the specified map of raw messages.
func UnmarshalDialogNodeVisitedDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DialogNodeVisitedDetails)
	err = core.UnmarshalPrimitive(m, "dialog_node", &obj.DialogNode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "title", &obj.Title)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "conditions", &obj.Conditions)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DialogSuggestion : DialogSuggestion struct
type DialogSuggestion struct {
	// The user-facing label for the disambiguation option. This label is taken from the **title** or **user_label**
	// property of the corresponding dialog node.
	Label *string `json:"label" validate:"required"`

	// An object defining the message input, intents, and entities to be sent to the Watson Assistant service if the user
	// selects the corresponding disambiguation option.
	Value *DialogSuggestionValue `json:"value" validate:"required"`

	// The dialog output that will be returned from the Watson Assistant service if the user selects the corresponding
	// option.
	Output map[string]interface{} `json:"output,omitempty"`

	// The unique ID of the dialog node that the **label** property is taken from. The **label** property is populated
	// using the value of the dialog node's **title** or **user_label** property.
	DialogNode *string `json:"dialog_node,omitempty"`
}

// NewDialogSuggestion : Instantiate DialogSuggestion (Generic Model Constructor)
func (*AssistantV1) NewDialogSuggestion(label string, value *DialogSuggestionValue) (_model *DialogSuggestion, err error) {
	_model = &DialogSuggestion{
		Label: core.StringPtr(label),
		Value: value,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalDialogSuggestion unmarshals an instance of DialogSuggestion from the specified map of raw messages.
func UnmarshalDialogSuggestion(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DialogSuggestion)
	err = core.UnmarshalPrimitive(m, "label", &obj.Label)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "value", &obj.Value, UnmarshalDialogSuggestionValue)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "output", &obj.Output)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "dialog_node", &obj.DialogNode)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DialogSuggestionValue : An object defining the message input, intents, and entities to be sent to the Watson Assistant service if the user
// selects the corresponding disambiguation option.
type DialogSuggestionValue struct {
	// An input object that includes the input text.
	Input *MessageInput `json:"input,omitempty"`

	// An array of intents to be sent along with the user input.
	Intents []RuntimeIntent `json:"intents,omitempty"`

	// An array of entities to be sent along with the user input.
	Entities []RuntimeEntity `json:"entities,omitempty"`
}

// UnmarshalDialogSuggestionValue unmarshals an instance of DialogSuggestionValue from the specified map of raw messages.
func UnmarshalDialogSuggestionValue(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DialogSuggestionValue)
	err = core.UnmarshalModel(m, "input", &obj.Input, UnmarshalMessageInput)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "intents", &obj.Intents, UnmarshalRuntimeIntent)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "entities", &obj.Entities, UnmarshalRuntimeEntity)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Entity : Entity struct
type Entity struct {
	// The name of the entity. This string must conform to the following restrictions:
	// - It can contain only Unicode alphanumeric, underscore, and hyphen characters.
	// - If you specify an entity name beginning with the reserved prefix `sys-`, it must be the name of a system entity
	// that you want to enable. (Any entity content specified with the request is ignored.).
	Entity *string `json:"entity" validate:"required"`

	// The description of the entity. This string cannot contain carriage return, newline, or tab characters.
	Description *string `json:"description,omitempty"`

	// Any metadata related to the entity.
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// Whether to use fuzzy matching for the entity.
	FuzzyMatch *bool `json:"fuzzy_match,omitempty"`

	// The timestamp for creation of the object.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// The timestamp for the most recent update to the object.
	Updated *strfmt.DateTime `json:"updated,omitempty"`

	// An array of objects describing the entity values.
	Values []Value `json:"values,omitempty"`
}

// UnmarshalEntity unmarshals an instance of Entity from the specified map of raw messages.
func UnmarshalEntity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Entity)
	err = core.UnmarshalPrimitive(m, "entity", &obj.Entity)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "metadata", &obj.Metadata)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "fuzzy_match", &obj.FuzzyMatch)
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
	err = core.UnmarshalModel(m, "values", &obj.Values, UnmarshalValue)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EntityCollection : An array of objects describing the entities for the workspace.
type EntityCollection struct {
	// An array of objects describing the entities defined for the workspace.
	Entities []Entity `json:"entities" validate:"required"`

	// The pagination data for the returned objects.
	Pagination *Pagination `json:"pagination" validate:"required"`
}

// UnmarshalEntityCollection unmarshals an instance of EntityCollection from the specified map of raw messages.
func UnmarshalEntityCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EntityCollection)
	err = core.UnmarshalModel(m, "entities", &obj.Entities, UnmarshalEntity)
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

// EntityMention : An object describing a contextual entity mention.
type EntityMention struct {
	// The text of the user input example.
	Text *string `json:"text" validate:"required"`

	// The name of the intent.
	Intent *string `json:"intent" validate:"required"`

	// An array of zero-based character offsets that indicate where the entity mentions begin and end in the input text.
	Location []int64 `json:"location" validate:"required"`
}

// UnmarshalEntityMention unmarshals an instance of EntityMention from the specified map of raw messages.
func UnmarshalEntityMention(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EntityMention)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "intent", &obj.Intent)
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

// EntityMentionCollection : EntityMentionCollection struct
type EntityMentionCollection struct {
	// An array of objects describing the entity mentions defined for an entity.
	Examples []EntityMention `json:"examples" validate:"required"`

	// The pagination data for the returned objects.
	Pagination *Pagination `json:"pagination" validate:"required"`
}

// UnmarshalEntityMentionCollection unmarshals an instance of EntityMentionCollection from the specified map of raw messages.
func UnmarshalEntityMentionCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EntityMentionCollection)
	err = core.UnmarshalModel(m, "examples", &obj.Examples, UnmarshalEntityMention)
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

// Example : Example struct
type Example struct {
	// The text of a user input example. This string must conform to the following restrictions:
	// - It cannot contain carriage return, newline, or tab characters.
	// - It cannot consist of only whitespace characters.
	Text *string `json:"text" validate:"required"`

	// An array of contextual entity mentions.
	Mentions []Mention `json:"mentions,omitempty"`

	// The timestamp for creation of the object.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// The timestamp for the most recent update to the object.
	Updated *strfmt.DateTime `json:"updated,omitempty"`
}

// NewExample : Instantiate Example (Generic Model Constructor)
func (*AssistantV1) NewExample(text string) (_model *Example, err error) {
	_model = &Example{
		Text: core.StringPtr(text),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalExample unmarshals an instance of Example from the specified map of raw messages.
func UnmarshalExample(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Example)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "mentions", &obj.Mentions, UnmarshalMention)
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

// ExampleCollection : ExampleCollection struct
type ExampleCollection struct {
	// An array of objects describing the examples defined for the intent.
	Examples []Example `json:"examples" validate:"required"`

	// The pagination data for the returned objects.
	Pagination *Pagination `json:"pagination" validate:"required"`
}

// UnmarshalExampleCollection unmarshals an instance of ExampleCollection from the specified map of raw messages.
func UnmarshalExampleCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ExampleCollection)
	err = core.UnmarshalModel(m, "examples", &obj.Examples, UnmarshalExample)
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

// GetCounterexampleOptions : The GetCounterexample options.
type GetCounterexampleOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// The text of a user input counterexample (for example, `What are you wearing?`).
	Text *string `json:"-" validate:"required,ne="`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCounterexampleOptions : Instantiate GetCounterexampleOptions
func (*AssistantV1) NewGetCounterexampleOptions(workspaceID string, text string) *GetCounterexampleOptions {
	return &GetCounterexampleOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Text:        core.StringPtr(text),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *GetCounterexampleOptions) SetWorkspaceID(workspaceID string) *GetCounterexampleOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetText : Allow user to set Text
func (_options *GetCounterexampleOptions) SetText(text string) *GetCounterexampleOptions {
	_options.Text = core.StringPtr(text)
	return _options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (_options *GetCounterexampleOptions) SetIncludeAudit(includeAudit bool) *GetCounterexampleOptions {
	_options.IncludeAudit = core.BoolPtr(includeAudit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetCounterexampleOptions) SetHeaders(param map[string]string) *GetCounterexampleOptions {
	options.Headers = param
	return options
}

// GetDialogNodeOptions : The GetDialogNode options.
type GetDialogNodeOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// The dialog node ID (for example, `node_1_1479323581900`).
	DialogNode *string `json:"-" validate:"required,ne="`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetDialogNodeOptions : Instantiate GetDialogNodeOptions
func (*AssistantV1) NewGetDialogNodeOptions(workspaceID string, dialogNode string) *GetDialogNodeOptions {
	return &GetDialogNodeOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		DialogNode:  core.StringPtr(dialogNode),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *GetDialogNodeOptions) SetWorkspaceID(workspaceID string) *GetDialogNodeOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetDialogNode : Allow user to set DialogNode
func (_options *GetDialogNodeOptions) SetDialogNode(dialogNode string) *GetDialogNodeOptions {
	_options.DialogNode = core.StringPtr(dialogNode)
	return _options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (_options *GetDialogNodeOptions) SetIncludeAudit(includeAudit bool) *GetDialogNodeOptions {
	_options.IncludeAudit = core.BoolPtr(includeAudit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDialogNodeOptions) SetHeaders(param map[string]string) *GetDialogNodeOptions {
	options.Headers = param
	return options
}

// GetEntityOptions : The GetEntity options.
type GetEntityOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// The name of the entity.
	Entity *string `json:"-" validate:"required,ne="`

	// Whether to include all element content in the returned data. If **export**=`false`, the returned data includes only
	// information about the element itself. If **export**=`true`, all content, including subelements, is included.
	Export *bool `json:"-"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetEntityOptions : Instantiate GetEntityOptions
func (*AssistantV1) NewGetEntityOptions(workspaceID string, entity string) *GetEntityOptions {
	return &GetEntityOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Entity:      core.StringPtr(entity),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *GetEntityOptions) SetWorkspaceID(workspaceID string) *GetEntityOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetEntity : Allow user to set Entity
func (_options *GetEntityOptions) SetEntity(entity string) *GetEntityOptions {
	_options.Entity = core.StringPtr(entity)
	return _options
}

// SetExport : Allow user to set Export
func (_options *GetEntityOptions) SetExport(export bool) *GetEntityOptions {
	_options.Export = core.BoolPtr(export)
	return _options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (_options *GetEntityOptions) SetIncludeAudit(includeAudit bool) *GetEntityOptions {
	_options.IncludeAudit = core.BoolPtr(includeAudit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetEntityOptions) SetHeaders(param map[string]string) *GetEntityOptions {
	options.Headers = param
	return options
}

// GetExampleOptions : The GetExample options.
type GetExampleOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// The intent name.
	Intent *string `json:"-" validate:"required,ne="`

	// The text of the user input example.
	Text *string `json:"-" validate:"required,ne="`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetExampleOptions : Instantiate GetExampleOptions
func (*AssistantV1) NewGetExampleOptions(workspaceID string, intent string, text string) *GetExampleOptions {
	return &GetExampleOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Intent:      core.StringPtr(intent),
		Text:        core.StringPtr(text),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *GetExampleOptions) SetWorkspaceID(workspaceID string) *GetExampleOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetIntent : Allow user to set Intent
func (_options *GetExampleOptions) SetIntent(intent string) *GetExampleOptions {
	_options.Intent = core.StringPtr(intent)
	return _options
}

// SetText : Allow user to set Text
func (_options *GetExampleOptions) SetText(text string) *GetExampleOptions {
	_options.Text = core.StringPtr(text)
	return _options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (_options *GetExampleOptions) SetIncludeAudit(includeAudit bool) *GetExampleOptions {
	_options.IncludeAudit = core.BoolPtr(includeAudit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetExampleOptions) SetHeaders(param map[string]string) *GetExampleOptions {
	options.Headers = param
	return options
}

// GetIntentOptions : The GetIntent options.
type GetIntentOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// The intent name.
	Intent *string `json:"-" validate:"required,ne="`

	// Whether to include all element content in the returned data. If **export**=`false`, the returned data includes only
	// information about the element itself. If **export**=`true`, all content, including subelements, is included.
	Export *bool `json:"-"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetIntentOptions : Instantiate GetIntentOptions
func (*AssistantV1) NewGetIntentOptions(workspaceID string, intent string) *GetIntentOptions {
	return &GetIntentOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Intent:      core.StringPtr(intent),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *GetIntentOptions) SetWorkspaceID(workspaceID string) *GetIntentOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetIntent : Allow user to set Intent
func (_options *GetIntentOptions) SetIntent(intent string) *GetIntentOptions {
	_options.Intent = core.StringPtr(intent)
	return _options
}

// SetExport : Allow user to set Export
func (_options *GetIntentOptions) SetExport(export bool) *GetIntentOptions {
	_options.Export = core.BoolPtr(export)
	return _options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (_options *GetIntentOptions) SetIncludeAudit(includeAudit bool) *GetIntentOptions {
	_options.IncludeAudit = core.BoolPtr(includeAudit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetIntentOptions) SetHeaders(param map[string]string) *GetIntentOptions {
	options.Headers = param
	return options
}

// GetSynonymOptions : The GetSynonym options.
type GetSynonymOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// The name of the entity.
	Entity *string `json:"-" validate:"required,ne="`

	// The text of the entity value.
	Value *string `json:"-" validate:"required,ne="`

	// The text of the synonym.
	Synonym *string `json:"-" validate:"required,ne="`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetSynonymOptions : Instantiate GetSynonymOptions
func (*AssistantV1) NewGetSynonymOptions(workspaceID string, entity string, value string, synonym string) *GetSynonymOptions {
	return &GetSynonymOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Entity:      core.StringPtr(entity),
		Value:       core.StringPtr(value),
		Synonym:     core.StringPtr(synonym),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *GetSynonymOptions) SetWorkspaceID(workspaceID string) *GetSynonymOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetEntity : Allow user to set Entity
func (_options *GetSynonymOptions) SetEntity(entity string) *GetSynonymOptions {
	_options.Entity = core.StringPtr(entity)
	return _options
}

// SetValue : Allow user to set Value
func (_options *GetSynonymOptions) SetValue(value string) *GetSynonymOptions {
	_options.Value = core.StringPtr(value)
	return _options
}

// SetSynonym : Allow user to set Synonym
func (_options *GetSynonymOptions) SetSynonym(synonym string) *GetSynonymOptions {
	_options.Synonym = core.StringPtr(synonym)
	return _options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (_options *GetSynonymOptions) SetIncludeAudit(includeAudit bool) *GetSynonymOptions {
	_options.IncludeAudit = core.BoolPtr(includeAudit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetSynonymOptions) SetHeaders(param map[string]string) *GetSynonymOptions {
	options.Headers = param
	return options
}

// GetValueOptions : The GetValue options.
type GetValueOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// The name of the entity.
	Entity *string `json:"-" validate:"required,ne="`

	// The text of the entity value.
	Value *string `json:"-" validate:"required,ne="`

	// Whether to include all element content in the returned data. If **export**=`false`, the returned data includes only
	// information about the element itself. If **export**=`true`, all content, including subelements, is included.
	Export *bool `json:"-"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetValueOptions : Instantiate GetValueOptions
func (*AssistantV1) NewGetValueOptions(workspaceID string, entity string, value string) *GetValueOptions {
	return &GetValueOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Entity:      core.StringPtr(entity),
		Value:       core.StringPtr(value),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *GetValueOptions) SetWorkspaceID(workspaceID string) *GetValueOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetEntity : Allow user to set Entity
func (_options *GetValueOptions) SetEntity(entity string) *GetValueOptions {
	_options.Entity = core.StringPtr(entity)
	return _options
}

// SetValue : Allow user to set Value
func (_options *GetValueOptions) SetValue(value string) *GetValueOptions {
	_options.Value = core.StringPtr(value)
	return _options
}

// SetExport : Allow user to set Export
func (_options *GetValueOptions) SetExport(export bool) *GetValueOptions {
	_options.Export = core.BoolPtr(export)
	return _options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (_options *GetValueOptions) SetIncludeAudit(includeAudit bool) *GetValueOptions {
	_options.IncludeAudit = core.BoolPtr(includeAudit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetValueOptions) SetHeaders(param map[string]string) *GetValueOptions {
	options.Headers = param
	return options
}

// GetWorkspaceOptions : The GetWorkspace options.
type GetWorkspaceOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// Whether to include all element content in the returned data. If **export**=`false`, the returned data includes only
	// information about the element itself. If **export**=`true`, all content, including subelements, is included.
	Export *bool `json:"-"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"-"`

	// Indicates how the returned workspace data will be sorted. This parameter is valid only if **export**=`true`. Specify
	// `sort=stable` to sort all workspace objects by unique identifier, in ascending alphabetical order.
	Sort *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetWorkspaceOptions.Sort property.
// Indicates how the returned workspace data will be sorted. This parameter is valid only if **export**=`true`. Specify
// `sort=stable` to sort all workspace objects by unique identifier, in ascending alphabetical order.
const (
	GetWorkspaceOptionsSortStableConst = "stable"
)

// NewGetWorkspaceOptions : Instantiate GetWorkspaceOptions
func (*AssistantV1) NewGetWorkspaceOptions(workspaceID string) *GetWorkspaceOptions {
	return &GetWorkspaceOptions{
		WorkspaceID: core.StringPtr(workspaceID),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *GetWorkspaceOptions) SetWorkspaceID(workspaceID string) *GetWorkspaceOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetExport : Allow user to set Export
func (_options *GetWorkspaceOptions) SetExport(export bool) *GetWorkspaceOptions {
	_options.Export = core.BoolPtr(export)
	return _options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (_options *GetWorkspaceOptions) SetIncludeAudit(includeAudit bool) *GetWorkspaceOptions {
	_options.IncludeAudit = core.BoolPtr(includeAudit)
	return _options
}

// SetSort : Allow user to set Sort
func (_options *GetWorkspaceOptions) SetSort(sort string) *GetWorkspaceOptions {
	_options.Sort = core.StringPtr(sort)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetWorkspaceOptions) SetHeaders(param map[string]string) *GetWorkspaceOptions {
	options.Headers = param
	return options
}

// Intent : Intent struct
type Intent struct {
	// The name of the intent. This string must conform to the following restrictions:
	// - It can contain only Unicode alphanumeric, underscore, hyphen, and dot characters.
	// - It cannot begin with the reserved prefix `sys-`.
	Intent *string `json:"intent" validate:"required"`

	// The description of the intent. This string cannot contain carriage return, newline, or tab characters.
	Description *string `json:"description,omitempty"`

	// The timestamp for creation of the object.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// The timestamp for the most recent update to the object.
	Updated *strfmt.DateTime `json:"updated,omitempty"`

	// An array of user input examples for the intent.
	Examples []Example `json:"examples,omitempty"`
}

// UnmarshalIntent unmarshals an instance of Intent from the specified map of raw messages.
func UnmarshalIntent(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Intent)
	err = core.UnmarshalPrimitive(m, "intent", &obj.Intent)
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
	err = core.UnmarshalModel(m, "examples", &obj.Examples, UnmarshalExample)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// IntentCollection : IntentCollection struct
type IntentCollection struct {
	// An array of objects describing the intents defined for the workspace.
	Intents []Intent `json:"intents" validate:"required"`

	// The pagination data for the returned objects.
	Pagination *Pagination `json:"pagination" validate:"required"`
}

// UnmarshalIntentCollection unmarshals an instance of IntentCollection from the specified map of raw messages.
func UnmarshalIntentCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(IntentCollection)
	err = core.UnmarshalModel(m, "intents", &obj.Intents, UnmarshalIntent)
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

// ListAllLogsOptions : The ListAllLogs options.
type ListAllLogsOptions struct {
	// A cacheable parameter that limits the results to those matching the specified filter. You must specify a filter
	// query that includes a value for `language`, as well as a value for `request.context.system.assistant_id`,
	// `workspace_id`, or `request.context.metadata.deployment`. These required filters must be specified using the exact
	// match (`::`) operator. For more information, see the
	// [documentation](https://cloud.ibm.com/docs/assistant?topic=assistant-filter-reference#filter-reference).
	Filter *string `json:"-" validate:"required"`

	// How to sort the returned log events. You can sort by **request_timestamp**. To reverse the sort order, prefix the
	// parameter value with a minus sign (`-`).
	Sort *string `json:"-"`

	// The number of records to return in each page of results.
	PageLimit *int64 `json:"-"`

	// A token identifying the page of results to retrieve.
	Cursor *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListAllLogsOptions : Instantiate ListAllLogsOptions
func (*AssistantV1) NewListAllLogsOptions(filter string) *ListAllLogsOptions {
	return &ListAllLogsOptions{
		Filter: core.StringPtr(filter),
	}
}

// SetFilter : Allow user to set Filter
func (_options *ListAllLogsOptions) SetFilter(filter string) *ListAllLogsOptions {
	_options.Filter = core.StringPtr(filter)
	return _options
}

// SetSort : Allow user to set Sort
func (_options *ListAllLogsOptions) SetSort(sort string) *ListAllLogsOptions {
	_options.Sort = core.StringPtr(sort)
	return _options
}

// SetPageLimit : Allow user to set PageLimit
func (_options *ListAllLogsOptions) SetPageLimit(pageLimit int64) *ListAllLogsOptions {
	_options.PageLimit = core.Int64Ptr(pageLimit)
	return _options
}

// SetCursor : Allow user to set Cursor
func (_options *ListAllLogsOptions) SetCursor(cursor string) *ListAllLogsOptions {
	_options.Cursor = core.StringPtr(cursor)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListAllLogsOptions) SetHeaders(param map[string]string) *ListAllLogsOptions {
	options.Headers = param
	return options
}

// ListCounterexamplesOptions : The ListCounterexamples options.
type ListCounterexamplesOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// The number of records to return in each page of results.
	PageLimit *int64 `json:"-"`

	// Whether to include information about the number of records that satisfy the request, regardless of the page limit.
	// If this parameter is `true`, the `pagination` object in the response includes the `total` property.
	IncludeCount *bool `json:"-"`

	// The attribute by which returned counterexamples will be sorted. To reverse the sort order, prefix the value with a
	// minus sign (`-`).
	Sort *string `json:"-"`

	// A token identifying the page of results to retrieve.
	Cursor *string `json:"-"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListCounterexamplesOptions.Sort property.
// The attribute by which returned counterexamples will be sorted. To reverse the sort order, prefix the value with a
// minus sign (`-`).
const (
	ListCounterexamplesOptionsSortTextConst    = "text"
	ListCounterexamplesOptionsSortUpdatedConst = "updated"
)

// NewListCounterexamplesOptions : Instantiate ListCounterexamplesOptions
func (*AssistantV1) NewListCounterexamplesOptions(workspaceID string) *ListCounterexamplesOptions {
	return &ListCounterexamplesOptions{
		WorkspaceID: core.StringPtr(workspaceID),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *ListCounterexamplesOptions) SetWorkspaceID(workspaceID string) *ListCounterexamplesOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetPageLimit : Allow user to set PageLimit
func (_options *ListCounterexamplesOptions) SetPageLimit(pageLimit int64) *ListCounterexamplesOptions {
	_options.PageLimit = core.Int64Ptr(pageLimit)
	return _options
}

// SetIncludeCount : Allow user to set IncludeCount
func (_options *ListCounterexamplesOptions) SetIncludeCount(includeCount bool) *ListCounterexamplesOptions {
	_options.IncludeCount = core.BoolPtr(includeCount)
	return _options
}

// SetSort : Allow user to set Sort
func (_options *ListCounterexamplesOptions) SetSort(sort string) *ListCounterexamplesOptions {
	_options.Sort = core.StringPtr(sort)
	return _options
}

// SetCursor : Allow user to set Cursor
func (_options *ListCounterexamplesOptions) SetCursor(cursor string) *ListCounterexamplesOptions {
	_options.Cursor = core.StringPtr(cursor)
	return _options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (_options *ListCounterexamplesOptions) SetIncludeAudit(includeAudit bool) *ListCounterexamplesOptions {
	_options.IncludeAudit = core.BoolPtr(includeAudit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListCounterexamplesOptions) SetHeaders(param map[string]string) *ListCounterexamplesOptions {
	options.Headers = param
	return options
}

// ListDialogNodesOptions : The ListDialogNodes options.
type ListDialogNodesOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// The number of records to return in each page of results.
	PageLimit *int64 `json:"-"`

	// Whether to include information about the number of records that satisfy the request, regardless of the page limit.
	// If this parameter is `true`, the `pagination` object in the response includes the `total` property.
	IncludeCount *bool `json:"-"`

	// The attribute by which returned dialog nodes will be sorted. To reverse the sort order, prefix the value with a
	// minus sign (`-`).
	Sort *string `json:"-"`

	// A token identifying the page of results to retrieve.
	Cursor *string `json:"-"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListDialogNodesOptions.Sort property.
// The attribute by which returned dialog nodes will be sorted. To reverse the sort order, prefix the value with a minus
// sign (`-`).
const (
	ListDialogNodesOptionsSortDialogNodeConst = "dialog_node"
	ListDialogNodesOptionsSortUpdatedConst    = "updated"
)

// NewListDialogNodesOptions : Instantiate ListDialogNodesOptions
func (*AssistantV1) NewListDialogNodesOptions(workspaceID string) *ListDialogNodesOptions {
	return &ListDialogNodesOptions{
		WorkspaceID: core.StringPtr(workspaceID),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *ListDialogNodesOptions) SetWorkspaceID(workspaceID string) *ListDialogNodesOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetPageLimit : Allow user to set PageLimit
func (_options *ListDialogNodesOptions) SetPageLimit(pageLimit int64) *ListDialogNodesOptions {
	_options.PageLimit = core.Int64Ptr(pageLimit)
	return _options
}

// SetIncludeCount : Allow user to set IncludeCount
func (_options *ListDialogNodesOptions) SetIncludeCount(includeCount bool) *ListDialogNodesOptions {
	_options.IncludeCount = core.BoolPtr(includeCount)
	return _options
}

// SetSort : Allow user to set Sort
func (_options *ListDialogNodesOptions) SetSort(sort string) *ListDialogNodesOptions {
	_options.Sort = core.StringPtr(sort)
	return _options
}

// SetCursor : Allow user to set Cursor
func (_options *ListDialogNodesOptions) SetCursor(cursor string) *ListDialogNodesOptions {
	_options.Cursor = core.StringPtr(cursor)
	return _options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (_options *ListDialogNodesOptions) SetIncludeAudit(includeAudit bool) *ListDialogNodesOptions {
	_options.IncludeAudit = core.BoolPtr(includeAudit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListDialogNodesOptions) SetHeaders(param map[string]string) *ListDialogNodesOptions {
	options.Headers = param
	return options
}

// ListEntitiesOptions : The ListEntities options.
type ListEntitiesOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// Whether to include all element content in the returned data. If **export**=`false`, the returned data includes only
	// information about the element itself. If **export**=`true`, all content, including subelements, is included.
	Export *bool `json:"-"`

	// The number of records to return in each page of results.
	PageLimit *int64 `json:"-"`

	// Whether to include information about the number of records that satisfy the request, regardless of the page limit.
	// If this parameter is `true`, the `pagination` object in the response includes the `total` property.
	IncludeCount *bool `json:"-"`

	// The attribute by which returned entities will be sorted. To reverse the sort order, prefix the value with a minus
	// sign (`-`).
	Sort *string `json:"-"`

	// A token identifying the page of results to retrieve.
	Cursor *string `json:"-"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListEntitiesOptions.Sort property.
// The attribute by which returned entities will be sorted. To reverse the sort order, prefix the value with a minus
// sign (`-`).
const (
	ListEntitiesOptionsSortEntityConst  = "entity"
	ListEntitiesOptionsSortUpdatedConst = "updated"
)

// NewListEntitiesOptions : Instantiate ListEntitiesOptions
func (*AssistantV1) NewListEntitiesOptions(workspaceID string) *ListEntitiesOptions {
	return &ListEntitiesOptions{
		WorkspaceID: core.StringPtr(workspaceID),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *ListEntitiesOptions) SetWorkspaceID(workspaceID string) *ListEntitiesOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetExport : Allow user to set Export
func (_options *ListEntitiesOptions) SetExport(export bool) *ListEntitiesOptions {
	_options.Export = core.BoolPtr(export)
	return _options
}

// SetPageLimit : Allow user to set PageLimit
func (_options *ListEntitiesOptions) SetPageLimit(pageLimit int64) *ListEntitiesOptions {
	_options.PageLimit = core.Int64Ptr(pageLimit)
	return _options
}

// SetIncludeCount : Allow user to set IncludeCount
func (_options *ListEntitiesOptions) SetIncludeCount(includeCount bool) *ListEntitiesOptions {
	_options.IncludeCount = core.BoolPtr(includeCount)
	return _options
}

// SetSort : Allow user to set Sort
func (_options *ListEntitiesOptions) SetSort(sort string) *ListEntitiesOptions {
	_options.Sort = core.StringPtr(sort)
	return _options
}

// SetCursor : Allow user to set Cursor
func (_options *ListEntitiesOptions) SetCursor(cursor string) *ListEntitiesOptions {
	_options.Cursor = core.StringPtr(cursor)
	return _options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (_options *ListEntitiesOptions) SetIncludeAudit(includeAudit bool) *ListEntitiesOptions {
	_options.IncludeAudit = core.BoolPtr(includeAudit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListEntitiesOptions) SetHeaders(param map[string]string) *ListEntitiesOptions {
	options.Headers = param
	return options
}

// ListExamplesOptions : The ListExamples options.
type ListExamplesOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// The intent name.
	Intent *string `json:"-" validate:"required,ne="`

	// The number of records to return in each page of results.
	PageLimit *int64 `json:"-"`

	// Whether to include information about the number of records that satisfy the request, regardless of the page limit.
	// If this parameter is `true`, the `pagination` object in the response includes the `total` property.
	IncludeCount *bool `json:"-"`

	// The attribute by which returned examples will be sorted. To reverse the sort order, prefix the value with a minus
	// sign (`-`).
	Sort *string `json:"-"`

	// A token identifying the page of results to retrieve.
	Cursor *string `json:"-"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListExamplesOptions.Sort property.
// The attribute by which returned examples will be sorted. To reverse the sort order, prefix the value with a minus
// sign (`-`).
const (
	ListExamplesOptionsSortTextConst    = "text"
	ListExamplesOptionsSortUpdatedConst = "updated"
)

// NewListExamplesOptions : Instantiate ListExamplesOptions
func (*AssistantV1) NewListExamplesOptions(workspaceID string, intent string) *ListExamplesOptions {
	return &ListExamplesOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Intent:      core.StringPtr(intent),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *ListExamplesOptions) SetWorkspaceID(workspaceID string) *ListExamplesOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetIntent : Allow user to set Intent
func (_options *ListExamplesOptions) SetIntent(intent string) *ListExamplesOptions {
	_options.Intent = core.StringPtr(intent)
	return _options
}

// SetPageLimit : Allow user to set PageLimit
func (_options *ListExamplesOptions) SetPageLimit(pageLimit int64) *ListExamplesOptions {
	_options.PageLimit = core.Int64Ptr(pageLimit)
	return _options
}

// SetIncludeCount : Allow user to set IncludeCount
func (_options *ListExamplesOptions) SetIncludeCount(includeCount bool) *ListExamplesOptions {
	_options.IncludeCount = core.BoolPtr(includeCount)
	return _options
}

// SetSort : Allow user to set Sort
func (_options *ListExamplesOptions) SetSort(sort string) *ListExamplesOptions {
	_options.Sort = core.StringPtr(sort)
	return _options
}

// SetCursor : Allow user to set Cursor
func (_options *ListExamplesOptions) SetCursor(cursor string) *ListExamplesOptions {
	_options.Cursor = core.StringPtr(cursor)
	return _options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (_options *ListExamplesOptions) SetIncludeAudit(includeAudit bool) *ListExamplesOptions {
	_options.IncludeAudit = core.BoolPtr(includeAudit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListExamplesOptions) SetHeaders(param map[string]string) *ListExamplesOptions {
	options.Headers = param
	return options
}

// ListIntentsOptions : The ListIntents options.
type ListIntentsOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// Whether to include all element content in the returned data. If **export**=`false`, the returned data includes only
	// information about the element itself. If **export**=`true`, all content, including subelements, is included.
	Export *bool `json:"-"`

	// The number of records to return in each page of results.
	PageLimit *int64 `json:"-"`

	// Whether to include information about the number of records that satisfy the request, regardless of the page limit.
	// If this parameter is `true`, the `pagination` object in the response includes the `total` property.
	IncludeCount *bool `json:"-"`

	// The attribute by which returned intents will be sorted. To reverse the sort order, prefix the value with a minus
	// sign (`-`).
	Sort *string `json:"-"`

	// A token identifying the page of results to retrieve.
	Cursor *string `json:"-"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListIntentsOptions.Sort property.
// The attribute by which returned intents will be sorted. To reverse the sort order, prefix the value with a minus sign
// (`-`).
const (
	ListIntentsOptionsSortIntentConst  = "intent"
	ListIntentsOptionsSortUpdatedConst = "updated"
)

// NewListIntentsOptions : Instantiate ListIntentsOptions
func (*AssistantV1) NewListIntentsOptions(workspaceID string) *ListIntentsOptions {
	return &ListIntentsOptions{
		WorkspaceID: core.StringPtr(workspaceID),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *ListIntentsOptions) SetWorkspaceID(workspaceID string) *ListIntentsOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetExport : Allow user to set Export
func (_options *ListIntentsOptions) SetExport(export bool) *ListIntentsOptions {
	_options.Export = core.BoolPtr(export)
	return _options
}

// SetPageLimit : Allow user to set PageLimit
func (_options *ListIntentsOptions) SetPageLimit(pageLimit int64) *ListIntentsOptions {
	_options.PageLimit = core.Int64Ptr(pageLimit)
	return _options
}

// SetIncludeCount : Allow user to set IncludeCount
func (_options *ListIntentsOptions) SetIncludeCount(includeCount bool) *ListIntentsOptions {
	_options.IncludeCount = core.BoolPtr(includeCount)
	return _options
}

// SetSort : Allow user to set Sort
func (_options *ListIntentsOptions) SetSort(sort string) *ListIntentsOptions {
	_options.Sort = core.StringPtr(sort)
	return _options
}

// SetCursor : Allow user to set Cursor
func (_options *ListIntentsOptions) SetCursor(cursor string) *ListIntentsOptions {
	_options.Cursor = core.StringPtr(cursor)
	return _options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (_options *ListIntentsOptions) SetIncludeAudit(includeAudit bool) *ListIntentsOptions {
	_options.IncludeAudit = core.BoolPtr(includeAudit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListIntentsOptions) SetHeaders(param map[string]string) *ListIntentsOptions {
	options.Headers = param
	return options
}

// ListLogsOptions : The ListLogs options.
type ListLogsOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// How to sort the returned log events. You can sort by **request_timestamp**. To reverse the sort order, prefix the
	// parameter value with a minus sign (`-`).
	Sort *string `json:"-"`

	// A cacheable parameter that limits the results to those matching the specified filter. For more information, see the
	// [documentation](https://cloud.ibm.com/docs/assistant?topic=assistant-filter-reference#filter-reference).
	Filter *string `json:"-"`

	// The number of records to return in each page of results.
	PageLimit *int64 `json:"-"`

	// A token identifying the page of results to retrieve.
	Cursor *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListLogsOptions : Instantiate ListLogsOptions
func (*AssistantV1) NewListLogsOptions(workspaceID string) *ListLogsOptions {
	return &ListLogsOptions{
		WorkspaceID: core.StringPtr(workspaceID),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *ListLogsOptions) SetWorkspaceID(workspaceID string) *ListLogsOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetSort : Allow user to set Sort
func (_options *ListLogsOptions) SetSort(sort string) *ListLogsOptions {
	_options.Sort = core.StringPtr(sort)
	return _options
}

// SetFilter : Allow user to set Filter
func (_options *ListLogsOptions) SetFilter(filter string) *ListLogsOptions {
	_options.Filter = core.StringPtr(filter)
	return _options
}

// SetPageLimit : Allow user to set PageLimit
func (_options *ListLogsOptions) SetPageLimit(pageLimit int64) *ListLogsOptions {
	_options.PageLimit = core.Int64Ptr(pageLimit)
	return _options
}

// SetCursor : Allow user to set Cursor
func (_options *ListLogsOptions) SetCursor(cursor string) *ListLogsOptions {
	_options.Cursor = core.StringPtr(cursor)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListLogsOptions) SetHeaders(param map[string]string) *ListLogsOptions {
	options.Headers = param
	return options
}

// ListMentionsOptions : The ListMentions options.
type ListMentionsOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// The name of the entity.
	Entity *string `json:"-" validate:"required,ne="`

	// Whether to include all element content in the returned data. If **export**=`false`, the returned data includes only
	// information about the element itself. If **export**=`true`, all content, including subelements, is included.
	Export *bool `json:"-"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListMentionsOptions : Instantiate ListMentionsOptions
func (*AssistantV1) NewListMentionsOptions(workspaceID string, entity string) *ListMentionsOptions {
	return &ListMentionsOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Entity:      core.StringPtr(entity),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *ListMentionsOptions) SetWorkspaceID(workspaceID string) *ListMentionsOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetEntity : Allow user to set Entity
func (_options *ListMentionsOptions) SetEntity(entity string) *ListMentionsOptions {
	_options.Entity = core.StringPtr(entity)
	return _options
}

// SetExport : Allow user to set Export
func (_options *ListMentionsOptions) SetExport(export bool) *ListMentionsOptions {
	_options.Export = core.BoolPtr(export)
	return _options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (_options *ListMentionsOptions) SetIncludeAudit(includeAudit bool) *ListMentionsOptions {
	_options.IncludeAudit = core.BoolPtr(includeAudit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListMentionsOptions) SetHeaders(param map[string]string) *ListMentionsOptions {
	options.Headers = param
	return options
}

// ListSynonymsOptions : The ListSynonyms options.
type ListSynonymsOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// The name of the entity.
	Entity *string `json:"-" validate:"required,ne="`

	// The text of the entity value.
	Value *string `json:"-" validate:"required,ne="`

	// The number of records to return in each page of results.
	PageLimit *int64 `json:"-"`

	// Whether to include information about the number of records that satisfy the request, regardless of the page limit.
	// If this parameter is `true`, the `pagination` object in the response includes the `total` property.
	IncludeCount *bool `json:"-"`

	// The attribute by which returned entity value synonyms will be sorted. To reverse the sort order, prefix the value
	// with a minus sign (`-`).
	Sort *string `json:"-"`

	// A token identifying the page of results to retrieve.
	Cursor *string `json:"-"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListSynonymsOptions.Sort property.
// The attribute by which returned entity value synonyms will be sorted. To reverse the sort order, prefix the value
// with a minus sign (`-`).
const (
	ListSynonymsOptionsSortSynonymConst = "synonym"
	ListSynonymsOptionsSortUpdatedConst = "updated"
)

// NewListSynonymsOptions : Instantiate ListSynonymsOptions
func (*AssistantV1) NewListSynonymsOptions(workspaceID string, entity string, value string) *ListSynonymsOptions {
	return &ListSynonymsOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Entity:      core.StringPtr(entity),
		Value:       core.StringPtr(value),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *ListSynonymsOptions) SetWorkspaceID(workspaceID string) *ListSynonymsOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetEntity : Allow user to set Entity
func (_options *ListSynonymsOptions) SetEntity(entity string) *ListSynonymsOptions {
	_options.Entity = core.StringPtr(entity)
	return _options
}

// SetValue : Allow user to set Value
func (_options *ListSynonymsOptions) SetValue(value string) *ListSynonymsOptions {
	_options.Value = core.StringPtr(value)
	return _options
}

// SetPageLimit : Allow user to set PageLimit
func (_options *ListSynonymsOptions) SetPageLimit(pageLimit int64) *ListSynonymsOptions {
	_options.PageLimit = core.Int64Ptr(pageLimit)
	return _options
}

// SetIncludeCount : Allow user to set IncludeCount
func (_options *ListSynonymsOptions) SetIncludeCount(includeCount bool) *ListSynonymsOptions {
	_options.IncludeCount = core.BoolPtr(includeCount)
	return _options
}

// SetSort : Allow user to set Sort
func (_options *ListSynonymsOptions) SetSort(sort string) *ListSynonymsOptions {
	_options.Sort = core.StringPtr(sort)
	return _options
}

// SetCursor : Allow user to set Cursor
func (_options *ListSynonymsOptions) SetCursor(cursor string) *ListSynonymsOptions {
	_options.Cursor = core.StringPtr(cursor)
	return _options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (_options *ListSynonymsOptions) SetIncludeAudit(includeAudit bool) *ListSynonymsOptions {
	_options.IncludeAudit = core.BoolPtr(includeAudit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListSynonymsOptions) SetHeaders(param map[string]string) *ListSynonymsOptions {
	options.Headers = param
	return options
}

// ListValuesOptions : The ListValues options.
type ListValuesOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// The name of the entity.
	Entity *string `json:"-" validate:"required,ne="`

	// Whether to include all element content in the returned data. If **export**=`false`, the returned data includes only
	// information about the element itself. If **export**=`true`, all content, including subelements, is included.
	Export *bool `json:"-"`

	// The number of records to return in each page of results.
	PageLimit *int64 `json:"-"`

	// Whether to include information about the number of records that satisfy the request, regardless of the page limit.
	// If this parameter is `true`, the `pagination` object in the response includes the `total` property.
	IncludeCount *bool `json:"-"`

	// The attribute by which returned entity values will be sorted. To reverse the sort order, prefix the value with a
	// minus sign (`-`).
	Sort *string `json:"-"`

	// A token identifying the page of results to retrieve.
	Cursor *string `json:"-"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListValuesOptions.Sort property.
// The attribute by which returned entity values will be sorted. To reverse the sort order, prefix the value with a
// minus sign (`-`).
const (
	ListValuesOptionsSortUpdatedConst = "updated"
	ListValuesOptionsSortValueConst   = "value"
)

// NewListValuesOptions : Instantiate ListValuesOptions
func (*AssistantV1) NewListValuesOptions(workspaceID string, entity string) *ListValuesOptions {
	return &ListValuesOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Entity:      core.StringPtr(entity),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *ListValuesOptions) SetWorkspaceID(workspaceID string) *ListValuesOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetEntity : Allow user to set Entity
func (_options *ListValuesOptions) SetEntity(entity string) *ListValuesOptions {
	_options.Entity = core.StringPtr(entity)
	return _options
}

// SetExport : Allow user to set Export
func (_options *ListValuesOptions) SetExport(export bool) *ListValuesOptions {
	_options.Export = core.BoolPtr(export)
	return _options
}

// SetPageLimit : Allow user to set PageLimit
func (_options *ListValuesOptions) SetPageLimit(pageLimit int64) *ListValuesOptions {
	_options.PageLimit = core.Int64Ptr(pageLimit)
	return _options
}

// SetIncludeCount : Allow user to set IncludeCount
func (_options *ListValuesOptions) SetIncludeCount(includeCount bool) *ListValuesOptions {
	_options.IncludeCount = core.BoolPtr(includeCount)
	return _options
}

// SetSort : Allow user to set Sort
func (_options *ListValuesOptions) SetSort(sort string) *ListValuesOptions {
	_options.Sort = core.StringPtr(sort)
	return _options
}

// SetCursor : Allow user to set Cursor
func (_options *ListValuesOptions) SetCursor(cursor string) *ListValuesOptions {
	_options.Cursor = core.StringPtr(cursor)
	return _options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (_options *ListValuesOptions) SetIncludeAudit(includeAudit bool) *ListValuesOptions {
	_options.IncludeAudit = core.BoolPtr(includeAudit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListValuesOptions) SetHeaders(param map[string]string) *ListValuesOptions {
	options.Headers = param
	return options
}

// ListWorkspacesOptions : The ListWorkspaces options.
type ListWorkspacesOptions struct {
	// The number of records to return in each page of results.
	PageLimit *int64 `json:"-"`

	// Whether to include information about the number of records that satisfy the request, regardless of the page limit.
	// If this parameter is `true`, the `pagination` object in the response includes the `total` property.
	IncludeCount *bool `json:"-"`

	// The attribute by which returned workspaces will be sorted. To reverse the sort order, prefix the value with a minus
	// sign (`-`).
	Sort *string `json:"-"`

	// A token identifying the page of results to retrieve.
	Cursor *string `json:"-"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListWorkspacesOptions.Sort property.
// The attribute by which returned workspaces will be sorted. To reverse the sort order, prefix the value with a minus
// sign (`-`).
const (
	ListWorkspacesOptionsSortNameConst    = "name"
	ListWorkspacesOptionsSortUpdatedConst = "updated"
)

// NewListWorkspacesOptions : Instantiate ListWorkspacesOptions
func (*AssistantV1) NewListWorkspacesOptions() *ListWorkspacesOptions {
	return &ListWorkspacesOptions{}
}

// SetPageLimit : Allow user to set PageLimit
func (_options *ListWorkspacesOptions) SetPageLimit(pageLimit int64) *ListWorkspacesOptions {
	_options.PageLimit = core.Int64Ptr(pageLimit)
	return _options
}

// SetIncludeCount : Allow user to set IncludeCount
func (_options *ListWorkspacesOptions) SetIncludeCount(includeCount bool) *ListWorkspacesOptions {
	_options.IncludeCount = core.BoolPtr(includeCount)
	return _options
}

// SetSort : Allow user to set Sort
func (_options *ListWorkspacesOptions) SetSort(sort string) *ListWorkspacesOptions {
	_options.Sort = core.StringPtr(sort)
	return _options
}

// SetCursor : Allow user to set Cursor
func (_options *ListWorkspacesOptions) SetCursor(cursor string) *ListWorkspacesOptions {
	_options.Cursor = core.StringPtr(cursor)
	return _options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (_options *ListWorkspacesOptions) SetIncludeAudit(includeAudit bool) *ListWorkspacesOptions {
	_options.IncludeAudit = core.BoolPtr(includeAudit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListWorkspacesOptions) SetHeaders(param map[string]string) *ListWorkspacesOptions {
	options.Headers = param
	return options
}

// Log : Log struct
type Log struct {
	// A request sent to the workspace, including the user input and context.
	Request *MessageRequest `json:"request" validate:"required"`

	// The response sent by the workspace, including the output text, detected intents and entities, and context.
	Response *MessageResponse `json:"response" validate:"required"`

	// A unique identifier for the logged event.
	LogID *string `json:"log_id" validate:"required"`

	// The timestamp for receipt of the message.
	RequestTimestamp *string `json:"request_timestamp" validate:"required"`

	// The timestamp for the system response to the message.
	ResponseTimestamp *string `json:"response_timestamp" validate:"required"`

	// The unique identifier of the workspace where the request was made.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The language of the workspace where the message request was made.
	Language *string `json:"language" validate:"required"`
}

// UnmarshalLog unmarshals an instance of Log from the specified map of raw messages.
func UnmarshalLog(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Log)
	err = core.UnmarshalModel(m, "request", &obj.Request, UnmarshalMessageRequest)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "response", &obj.Response, UnmarshalMessageResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "log_id", &obj.LogID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "request_timestamp", &obj.RequestTimestamp)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "response_timestamp", &obj.ResponseTimestamp)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "workspace_id", &obj.WorkspaceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "language", &obj.Language)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LogCollection : LogCollection struct
type LogCollection struct {
	// An array of objects describing log events.
	Logs []Log `json:"logs" validate:"required"`

	// The pagination data for the returned objects.
	Pagination *LogPagination `json:"pagination" validate:"required"`
}

// UnmarshalLogCollection unmarshals an instance of LogCollection from the specified map of raw messages.
func UnmarshalLogCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LogCollection)
	err = core.UnmarshalModel(m, "logs", &obj.Logs, UnmarshalLog)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "pagination", &obj.Pagination, UnmarshalLogPagination)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LogMessage : Log message details.
type LogMessage struct {
	// The severity of the log message.
	Level *string `json:"level" validate:"required"`

	// The text of the log message.
	Msg *string `json:"msg" validate:"required"`

	// A code that indicates the category to which the error message belongs.
	Code *string `json:"code" validate:"required"`

	// An object that identifies the dialog element that generated the error message.
	Source *LogMessageSource `json:"source,omitempty"`
}

// Constants associated with the LogMessage.Level property.
// The severity of the log message.
const (
	LogMessageLevelErrorConst = "error"
	LogMessageLevelInfoConst  = "info"
	LogMessageLevelWarnConst  = "warn"
)

// NewLogMessage : Instantiate LogMessage (Generic Model Constructor)
func (*AssistantV1) NewLogMessage(level string, msg string, code string) (_model *LogMessage, err error) {
	_model = &LogMessage{
		Level: core.StringPtr(level),
		Msg:   core.StringPtr(msg),
		Code:  core.StringPtr(code),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalLogMessage unmarshals an instance of LogMessage from the specified map of raw messages.
func UnmarshalLogMessage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LogMessage)
	err = core.UnmarshalPrimitive(m, "level", &obj.Level)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "msg", &obj.Msg)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "code", &obj.Code)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "source", &obj.Source, UnmarshalLogMessageSource)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LogMessageSource : An object that identifies the dialog element that generated the error message.
type LogMessageSource struct {
	// A string that indicates the type of dialog element that generated the error message.
	Type *string `json:"type,omitempty"`

	// The unique identifier of the dialog node that generated the error message.
	DialogNode *string `json:"dialog_node,omitempty"`
}

// Constants associated with the LogMessageSource.Type property.
// A string that indicates the type of dialog element that generated the error message.
const (
	LogMessageSourceTypeDialogNodeConst = "dialog_node"
)

// UnmarshalLogMessageSource unmarshals an instance of LogMessageSource from the specified map of raw messages.
func UnmarshalLogMessageSource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LogMessageSource)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "dialog_node", &obj.DialogNode)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LogPagination : The pagination data for the returned objects.
type LogPagination struct {
	// The URL that will return the next page of results, if any.
	NextURL *string `json:"next_url,omitempty"`

	// Reserved for future use.
	Matched *int64 `json:"matched,omitempty"`

	// A token identifying the next page of results.
	NextCursor *string `json:"next_cursor,omitempty"`
}

// UnmarshalLogPagination unmarshals an instance of LogPagination from the specified map of raw messages.
func UnmarshalLogPagination(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LogPagination)
	err = core.UnmarshalPrimitive(m, "next_url", &obj.NextURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "matched", &obj.Matched)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "next_cursor", &obj.NextCursor)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Mention : A mention of a contextual entity.
type Mention struct {
	// The name of the entity.
	Entity *string `json:"entity" validate:"required"`

	// An array of zero-based character offsets that indicate where the entity mentions begin and end in the input text.
	Location []int64 `json:"location" validate:"required"`
}

// NewMention : Instantiate Mention (Generic Model Constructor)
func (*AssistantV1) NewMention(entity string, location []int64) (_model *Mention, err error) {
	_model = &Mention{
		Entity:   core.StringPtr(entity),
		Location: location,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalMention unmarshals an instance of Mention from the specified map of raw messages.
func UnmarshalMention(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Mention)
	err = core.UnmarshalPrimitive(m, "entity", &obj.Entity)
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

// MessageContextMetadata : Metadata related to the message.
type MessageContextMetadata struct {
	// A label identifying the deployment environment, used for filtering log data. This string cannot contain carriage
	// return, newline, or tab characters.
	Deployment *string `json:"deployment,omitempty"`

	// A string value that identifies the user who is interacting with the workspace. The client must provide a unique
	// identifier for each individual end user who accesses the application. For user-based plans, this user ID is used to
	// identify unique users for billing purposes. This string cannot contain carriage return, newline, or tab characters.
	// If no value is specified in the input, **user_id** is automatically set to the value of **context.conversation_id**.
	//
	// **Note:** This property is the same as the **user_id** property at the root of the message body. If **user_id** is
	// specified in both locations in a message request, the value specified at the root is used.
	UserID *string `json:"user_id,omitempty"`
}

// UnmarshalMessageContextMetadata unmarshals an instance of MessageContextMetadata from the specified map of raw messages.
func UnmarshalMessageContextMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MessageContextMetadata)
	err = core.UnmarshalPrimitive(m, "deployment", &obj.Deployment)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "user_id", &obj.UserID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MessageInput : An input object that includes the input text.
type MessageInput struct {
	// The text of the user input. This string cannot contain carriage return, newline, or tab characters.
	Text *string `json:"text,omitempty"`

	// Whether to use spelling correction when processing the input. This property overrides the value of the
	// **spelling_suggestions** property in the workspace settings.
	SpellingSuggestions *bool `json:"spelling_suggestions,omitempty"`

	// Whether to use autocorrection when processing the input. If spelling correction is used and this property is
	// `false`, any suggested corrections are returned in the **suggested_text** property of the message response. If this
	// property is `true`, any corrections are automatically applied to the user input, and the original text is returned
	// in the **original_text** property of the message response. This property overrides the value of the
	// **spelling_auto_correct** property in the workspace settings.
	SpellingAutoCorrect *bool `json:"spelling_auto_correct,omitempty"`

	// Any suggested corrections of the input text. This property is returned only if spelling correction is enabled and
	// autocorrection is disabled.
	SuggestedText *string `json:"suggested_text,omitempty"`

	// The original user input text. This property is returned only if autocorrection is enabled and the user input was
	// corrected.
	OriginalText *string `json:"original_text,omitempty"`

	// Allows users to set arbitrary properties
	additionalProperties map[string]interface{}
}

// SetProperty allows the user to set an arbitrary property on an instance of MessageInput
func (o *MessageInput) SetProperty(key string, value interface{}) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string]interface{})
	}
	o.additionalProperties[key] = value
}

// SetProperties allows the user to set a map of arbitrary properties on an instance of MessageInput
func (o *MessageInput) SetProperties(m map[string]interface{}) {
	o.additionalProperties = make(map[string]interface{})
	for k, v := range m {
		o.additionalProperties[k] = v
	}
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of MessageInput
func (o *MessageInput) GetProperty(key string) interface{} {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of MessageInput
func (o *MessageInput) GetProperties() map[string]interface{} {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of MessageInput
func (o *MessageInput) MarshalJSON() (buffer []byte, err error) {
	m := make(map[string]interface{})
	if len(o.additionalProperties) > 0 {
		for k, v := range o.additionalProperties {
			m[k] = v
		}
	}
	if o.Text != nil {
		m["text"] = o.Text
	}
	if o.SpellingSuggestions != nil {
		m["spelling_suggestions"] = o.SpellingSuggestions
	}
	if o.SpellingAutoCorrect != nil {
		m["spelling_auto_correct"] = o.SpellingAutoCorrect
	}
	if o.SuggestedText != nil {
		m["suggested_text"] = o.SuggestedText
	}
	if o.OriginalText != nil {
		m["original_text"] = o.OriginalText
	}
	buffer, err = json.Marshal(m)
	return
}

// UnmarshalMessageInput unmarshals an instance of MessageInput from the specified map of raw messages.
func UnmarshalMessageInput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MessageInput)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	delete(m, "text")
	err = core.UnmarshalPrimitive(m, "spelling_suggestions", &obj.SpellingSuggestions)
	if err != nil {
		return
	}
	delete(m, "spelling_suggestions")
	err = core.UnmarshalPrimitive(m, "spelling_auto_correct", &obj.SpellingAutoCorrect)
	if err != nil {
		return
	}
	delete(m, "spelling_auto_correct")
	err = core.UnmarshalPrimitive(m, "suggested_text", &obj.SuggestedText)
	if err != nil {
		return
	}
	delete(m, "suggested_text")
	err = core.UnmarshalPrimitive(m, "original_text", &obj.OriginalText)
	if err != nil {
		return
	}
	delete(m, "original_text")
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

// MessageOptions : The Message options.
type MessageOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// An input object that includes the input text.
	Input *MessageInput `json:"input,omitempty"`

	// Intents to use when evaluating the user input. Include intents from the previous response to continue using those
	// intents rather than trying to recognize intents in the new input.
	Intents []RuntimeIntent `json:"intents,omitempty"`

	// Entities to use when evaluating the message. Include entities from the previous response to continue using those
	// entities rather than detecting entities in the new input.
	Entities []RuntimeEntity `json:"entities,omitempty"`

	// Whether to return more than one intent. A value of `true` indicates that all matching intents are returned.
	AlternateIntents *bool `json:"alternate_intents,omitempty"`

	// State information for the conversation. To maintain state, include the context from the previous response.
	Context *Context `json:"context,omitempty"`

	// An output object that includes the response to the user, the dialog nodes that were triggered, and messages from the
	// log.
	Output *OutputData `json:"output,omitempty"`

	// A string value that identifies the user who is interacting with the workspace. The client must provide a unique
	// identifier for each individual end user who accesses the application. For user-based plans, this user ID is used to
	// identify unique users for billing purposes. This string cannot contain carriage return, newline, or tab characters.
	// If no value is specified in the input, **user_id** is automatically set to the value of **context.conversation_id**.
	//
	// **Note:** This property is the same as the **user_id** property in the context metadata. If **user_id** is specified
	// in both locations in a message request, the value specified at the root is used.
	UserID *string `json:"user_id,omitempty"`

	// Whether to include additional diagnostic information about the dialog nodes that were visited during processing of
	// the message.
	NodesVisitedDetails *bool `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewMessageOptions : Instantiate MessageOptions
func (*AssistantV1) NewMessageOptions(workspaceID string) *MessageOptions {
	return &MessageOptions{
		WorkspaceID: core.StringPtr(workspaceID),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *MessageOptions) SetWorkspaceID(workspaceID string) *MessageOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetInput : Allow user to set Input
func (_options *MessageOptions) SetInput(input *MessageInput) *MessageOptions {
	_options.Input = input
	return _options
}

// SetIntents : Allow user to set Intents
func (_options *MessageOptions) SetIntents(intents []RuntimeIntent) *MessageOptions {
	_options.Intents = intents
	return _options
}

// SetEntities : Allow user to set Entities
func (_options *MessageOptions) SetEntities(entities []RuntimeEntity) *MessageOptions {
	_options.Entities = entities
	return _options
}

// SetAlternateIntents : Allow user to set AlternateIntents
func (_options *MessageOptions) SetAlternateIntents(alternateIntents bool) *MessageOptions {
	_options.AlternateIntents = core.BoolPtr(alternateIntents)
	return _options
}

// SetContext : Allow user to set Context
func (_options *MessageOptions) SetContext(context *Context) *MessageOptions {
	_options.Context = context
	return _options
}

// SetOutput : Allow user to set Output
func (_options *MessageOptions) SetOutput(output *OutputData) *MessageOptions {
	_options.Output = output
	return _options
}

// SetUserID : Allow user to set UserID
func (_options *MessageOptions) SetUserID(userID string) *MessageOptions {
	_options.UserID = core.StringPtr(userID)
	return _options
}

// SetNodesVisitedDetails : Allow user to set NodesVisitedDetails
func (_options *MessageOptions) SetNodesVisitedDetails(nodesVisitedDetails bool) *MessageOptions {
	_options.NodesVisitedDetails = core.BoolPtr(nodesVisitedDetails)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *MessageOptions) SetHeaders(param map[string]string) *MessageOptions {
	options.Headers = param
	return options
}

// MessageRequest : A request sent to the workspace, including the user input and context.
type MessageRequest struct {
	// An input object that includes the input text.
	Input *MessageInput `json:"input,omitempty"`

	// Intents to use when evaluating the user input. Include intents from the previous response to continue using those
	// intents rather than trying to recognize intents in the new input.
	Intents []RuntimeIntent `json:"intents,omitempty"`

	// Entities to use when evaluating the message. Include entities from the previous response to continue using those
	// entities rather than detecting entities in the new input.
	Entities []RuntimeEntity `json:"entities,omitempty"`

	// Whether to return more than one intent. A value of `true` indicates that all matching intents are returned.
	AlternateIntents *bool `json:"alternate_intents,omitempty"`

	// State information for the conversation. To maintain state, include the context from the previous response.
	Context *Context `json:"context,omitempty"`

	// An output object that includes the response to the user, the dialog nodes that were triggered, and messages from the
	// log.
	Output *OutputData `json:"output,omitempty"`

	// An array of objects describing any actions requested by the dialog node.
	Actions []DialogNodeAction `json:"actions,omitempty"`

	// A string value that identifies the user who is interacting with the workspace. The client must provide a unique
	// identifier for each individual end user who accesses the application. For user-based plans, this user ID is used to
	// identify unique users for billing purposes. This string cannot contain carriage return, newline, or tab characters.
	// If no value is specified in the input, **user_id** is automatically set to the value of **context.conversation_id**.
	//
	// **Note:** This property is the same as the **user_id** property in the context metadata. If **user_id** is specified
	// in both locations in a message request, the value specified at the root is used.
	UserID *string `json:"user_id,omitempty"`
}

// UnmarshalMessageRequest unmarshals an instance of MessageRequest from the specified map of raw messages.
func UnmarshalMessageRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MessageRequest)
	err = core.UnmarshalModel(m, "input", &obj.Input, UnmarshalMessageInput)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "intents", &obj.Intents, UnmarshalRuntimeIntent)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "entities", &obj.Entities, UnmarshalRuntimeEntity)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "alternate_intents", &obj.AlternateIntents)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "context", &obj.Context, UnmarshalContext)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "output", &obj.Output, UnmarshalOutputData)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "actions", &obj.Actions, UnmarshalDialogNodeAction)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "user_id", &obj.UserID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MessageResponse : The response sent by the workspace, including the output text, detected intents and entities, and context.
type MessageResponse struct {
	// An input object that includes the input text.
	Input *MessageInput `json:"input" validate:"required"`

	// An array of intents recognized in the user input, sorted in descending order of confidence.
	Intents []RuntimeIntent `json:"intents" validate:"required"`

	// An array of entities identified in the user input.
	Entities []RuntimeEntity `json:"entities" validate:"required"`

	// Whether to return more than one intent. A value of `true` indicates that all matching intents are returned.
	AlternateIntents *bool `json:"alternate_intents,omitempty"`

	// State information for the conversation. To maintain state, include the context from the previous response.
	Context *Context `json:"context" validate:"required"`

	// An output object that includes the response to the user, the dialog nodes that were triggered, and messages from the
	// log.
	Output *OutputData `json:"output" validate:"required"`

	// An array of objects describing any actions requested by the dialog node.
	Actions []DialogNodeAction `json:"actions,omitempty"`

	// A string value that identifies the user who is interacting with the workspace. The client must provide a unique
	// identifier for each individual end user who accesses the application. For user-based plans, this user ID is used to
	// identify unique users for billing purposes. This string cannot contain carriage return, newline, or tab characters.
	// If no value is specified in the input, **user_id** is automatically set to the value of **context.conversation_id**.
	//
	// **Note:** This property is the same as the **user_id** property in the context metadata. If **user_id** is specified
	// in both locations in a message request, the value specified at the root is used.
	UserID *string `json:"user_id" validate:"required"`
}

// UnmarshalMessageResponse unmarshals an instance of MessageResponse from the specified map of raw messages.
func UnmarshalMessageResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MessageResponse)
	err = core.UnmarshalModel(m, "input", &obj.Input, UnmarshalMessageInput)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "intents", &obj.Intents, UnmarshalRuntimeIntent)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "entities", &obj.Entities, UnmarshalRuntimeEntity)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "alternate_intents", &obj.AlternateIntents)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "context", &obj.Context, UnmarshalContext)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "output", &obj.Output, UnmarshalOutputData)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "actions", &obj.Actions, UnmarshalDialogNodeAction)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "user_id", &obj.UserID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// OutputData : An output object that includes the response to the user, the dialog nodes that were triggered, and messages from the
// log.
type OutputData struct {
	// An array of the nodes that were triggered to create the response, in the order in which they were visited. This
	// information is useful for debugging and for tracing the path taken through the node tree.
	NodesVisited []string `json:"nodes_visited,omitempty"`

	// An array of objects containing detailed diagnostic information about the nodes that were triggered during processing
	// of the input message. Included only if **nodes_visited_details** is set to `true` in the message request.
	NodesVisitedDetails []DialogNodeVisitedDetails `json:"nodes_visited_details,omitempty"`

	// An array of up to 50 messages logged with the request.
	LogMessages []LogMessage `json:"log_messages" validate:"required"`

	// An array of responses to the user.
	Text []string `json:"text" validate:"required"`

	// Output intended for any channel. It is the responsibility of the client application to implement the supported
	// response types.
	Generic []RuntimeResponseGenericIntf `json:"generic,omitempty"`

	// Allows users to set arbitrary properties
	additionalProperties map[string]interface{}
}

// NewOutputData : Instantiate OutputData (Generic Model Constructor)
func (*AssistantV1) NewOutputData(logMessages []LogMessage, text []string) (_model *OutputData, err error) {
	_model = &OutputData{
		LogMessages: logMessages,
		Text:        text,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// SetProperty allows the user to set an arbitrary property on an instance of OutputData
func (o *OutputData) SetProperty(key string, value interface{}) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string]interface{})
	}
	o.additionalProperties[key] = value
}

// SetProperties allows the user to set a map of arbitrary properties on an instance of OutputData
func (o *OutputData) SetProperties(m map[string]interface{}) {
	o.additionalProperties = make(map[string]interface{})
	for k, v := range m {
		o.additionalProperties[k] = v
	}
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of OutputData
func (o *OutputData) GetProperty(key string) interface{} {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of OutputData
func (o *OutputData) GetProperties() map[string]interface{} {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of OutputData
func (o *OutputData) MarshalJSON() (buffer []byte, err error) {
	m := make(map[string]interface{})
	if len(o.additionalProperties) > 0 {
		for k, v := range o.additionalProperties {
			m[k] = v
		}
	}
	if o.NodesVisited != nil {
		m["nodes_visited"] = o.NodesVisited
	}
	if o.NodesVisitedDetails != nil {
		m["nodes_visited_details"] = o.NodesVisitedDetails
	}
	if o.LogMessages != nil {
		m["log_messages"] = o.LogMessages
	}
	if o.Text != nil {
		m["text"] = o.Text
	}
	if o.Generic != nil {
		m["generic"] = o.Generic
	}
	buffer, err = json.Marshal(m)
	return
}

// UnmarshalOutputData unmarshals an instance of OutputData from the specified map of raw messages.
func UnmarshalOutputData(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OutputData)
	err = core.UnmarshalPrimitive(m, "nodes_visited", &obj.NodesVisited)
	if err != nil {
		return
	}
	delete(m, "nodes_visited")
	err = core.UnmarshalModel(m, "nodes_visited_details", &obj.NodesVisitedDetails, UnmarshalDialogNodeVisitedDetails)
	if err != nil {
		return
	}
	delete(m, "nodes_visited_details")
	err = core.UnmarshalModel(m, "log_messages", &obj.LogMessages, UnmarshalLogMessage)
	if err != nil {
		return
	}
	delete(m, "log_messages")
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	delete(m, "text")
	err = core.UnmarshalModel(m, "generic", &obj.Generic, UnmarshalRuntimeResponseGeneric)
	if err != nil {
		return
	}
	delete(m, "generic")
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

// Pagination : The pagination data for the returned objects.
type Pagination struct {
	// The URL that will return the same page of results.
	RefreshURL *string `json:"refresh_url" validate:"required"`

	// The URL that will return the next page of results.
	NextURL *string `json:"next_url,omitempty"`

	// The total number of objects that satisfy the request. This total includes all results, not just those included in
	// the current page.
	Total *int64 `json:"total,omitempty"`

	// Reserved for future use.
	Matched *int64 `json:"matched,omitempty"`

	// A token identifying the current page of results.
	RefreshCursor *string `json:"refresh_cursor,omitempty"`

	// A token identifying the next page of results.
	NextCursor *string `json:"next_cursor,omitempty"`
}

// UnmarshalPagination unmarshals an instance of Pagination from the specified map of raw messages.
func UnmarshalPagination(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Pagination)
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
	err = core.UnmarshalPrimitive(m, "matched", &obj.Matched)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "refresh_cursor", &obj.RefreshCursor)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "next_cursor", &obj.NextCursor)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResponseGenericChannel : ResponseGenericChannel struct
type ResponseGenericChannel struct {
	// A channel for which the response is intended.
	Channel *string `json:"channel,omitempty"`
}

// Constants associated with the ResponseGenericChannel.Channel property.
// A channel for which the response is intended.
const (
	ResponseGenericChannelChannelChatConst           = "chat"
	ResponseGenericChannelChannelFacebookConst       = "facebook"
	ResponseGenericChannelChannelIntercomConst       = "intercom"
	ResponseGenericChannelChannelSlackConst          = "slack"
	ResponseGenericChannelChannelTextMessagingConst  = "text_messaging"
	ResponseGenericChannelChannelVoiceTelephonyConst = "voice_telephony"
	ResponseGenericChannelChannelWhatsappConst       = "whatsapp"
)

// UnmarshalResponseGenericChannel unmarshals an instance of ResponseGenericChannel from the specified map of raw messages.
func UnmarshalResponseGenericChannel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResponseGenericChannel)
	err = core.UnmarshalPrimitive(m, "channel", &obj.Channel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuntimeEntity : A term from the request that was identified as an entity.
type RuntimeEntity struct {
	// An entity detected in the input.
	Entity *string `json:"entity" validate:"required"`

	// An array of zero-based character offsets that indicate where the detected entity values begin and end in the input
	// text.
	Location []int64 `json:"location,omitempty"`

	// The entity value that was recognized in the user input.
	Value *string `json:"value" validate:"required"`

	// A decimal percentage that represents Watson's confidence in the recognized entity.
	Confidence *float64 `json:"confidence,omitempty"`

	// **Deprecated.** Any metadata for the entity.
	//
	// Beginning with the `2021-06-14` API version, the `metadata` property is no longer returned. For information about
	// system entities recognized in the user input, see the `interpretation` property.
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// The recognized capture groups for the entity, as defined by the entity pattern.
	Groups []CaptureGroup `json:"groups,omitempty"`

	// An object containing detailed information about the entity recognized in the user input.
	//
	// For more information about how system entities are interpreted, see the
	// [documentation](https://cloud.ibm.com/docs/assistant?topic=assistant-system-entities).
	Interpretation *RuntimeEntityInterpretation `json:"interpretation,omitempty"`

	// An array of possible alternative values that the user might have intended instead of the value returned in the
	// **value** property. This property is returned only for `@sys-time` and `@sys-date` entities when the user's input is
	// ambiguous.
	//
	// This property is included only if the new system entities are enabled for the workspace.
	Alternatives []RuntimeEntityAlternative `json:"alternatives,omitempty"`

	// An object describing the role played by a system entity that is specifies the beginning or end of a range recognized
	// in the user input. This property is included only if the new system entities are enabled for the workspace.
	Role *RuntimeEntityRole `json:"role,omitempty"`
}

// NewRuntimeEntity : Instantiate RuntimeEntity (Generic Model Constructor)
func (*AssistantV1) NewRuntimeEntity(entity string, value string) (_model *RuntimeEntity, err error) {
	_model = &RuntimeEntity{
		Entity: core.StringPtr(entity),
		Value:  core.StringPtr(value),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalRuntimeEntity unmarshals an instance of RuntimeEntity from the specified map of raw messages.
func UnmarshalRuntimeEntity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuntimeEntity)
	err = core.UnmarshalPrimitive(m, "entity", &obj.Entity)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "location", &obj.Location)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "confidence", &obj.Confidence)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "metadata", &obj.Metadata)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "groups", &obj.Groups, UnmarshalCaptureGroup)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "interpretation", &obj.Interpretation, UnmarshalRuntimeEntityInterpretation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "alternatives", &obj.Alternatives, UnmarshalRuntimeEntityAlternative)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "role", &obj.Role, UnmarshalRuntimeEntityRole)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuntimeEntityAlternative : An alternative value for the recognized entity.
type RuntimeEntityAlternative struct {
	// The entity value that was recognized in the user input.
	Value *string `json:"value,omitempty"`

	// A decimal percentage that represents Watson's confidence in the recognized entity.
	Confidence *float64 `json:"confidence,omitempty"`
}

// UnmarshalRuntimeEntityAlternative unmarshals an instance of RuntimeEntityAlternative from the specified map of raw messages.
func UnmarshalRuntimeEntityAlternative(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuntimeEntityAlternative)
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
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

// RuntimeEntityInterpretation : RuntimeEntityInterpretation struct
type RuntimeEntityInterpretation struct {
	// The calendar used to represent a recognized date (for example, `Gregorian`).
	CalendarType *string `json:"calendar_type,omitempty"`

	// A unique identifier used to associate a recognized time and date. If the user input contains a date and time that
	// are mentioned together (for example, `Today at 5`, the same **datetime_link** value is returned for both the
	// `@sys-date` and `@sys-time` entities).
	DatetimeLink *string `json:"datetime_link,omitempty"`

	// A locale-specific holiday name (such as `thanksgiving` or `christmas`). This property is included when a `@sys-date`
	// entity is recognized based on a holiday name in the user input.
	Festival *string `json:"festival,omitempty"`

	// The precision or duration of a time range specified by a recognized `@sys-time` or `@sys-date` entity.
	Granularity *string `json:"granularity,omitempty"`

	// A unique identifier used to associate multiple recognized `@sys-date`, `@sys-time`, or `@sys-number` entities that
	// are recognized as a range of values in the user's input (for example, `from July 4 until July 14` or `from 20 to
	// 25`).
	RangeLink *string `json:"range_link,omitempty"`

	// The word in the user input that indicates that a `sys-date` or `sys-time` entity is part of an implied range where
	// only one date or time is specified (for example, `since` or `until`).
	RangeModifier *string `json:"range_modifier,omitempty"`

	// A recognized mention of a relative day, represented numerically as an offset from the current date (for example,
	// `-1` for `yesterday` or `10` for `in ten days`).
	RelativeDay *float64 `json:"relative_day,omitempty"`

	// A recognized mention of a relative month, represented numerically as an offset from the current month (for example,
	// `1` for `next month` or `-3` for `three months ago`).
	RelativeMonth *float64 `json:"relative_month,omitempty"`

	// A recognized mention of a relative week, represented numerically as an offset from the current week (for example,
	// `2` for `in two weeks` or `-1` for `last week).
	RelativeWeek *float64 `json:"relative_week,omitempty"`

	// A recognized mention of a relative date range for a weekend, represented numerically as an offset from the current
	// weekend (for example, `0` for `this weekend` or `-1` for `last weekend`).
	RelativeWeekend *float64 `json:"relative_weekend,omitempty"`

	// A recognized mention of a relative year, represented numerically as an offset from the current year (for example,
	// `1` for `next year` or `-5` for `five years ago`).
	RelativeYear *float64 `json:"relative_year,omitempty"`

	// A recognized mention of a specific date, represented numerically as the date within the month (for example, `30` for
	// `June 30`.).
	SpecificDay *float64 `json:"specific_day,omitempty"`

	// A recognized mention of a specific day of the week as a lowercase string (for example, `monday`).
	SpecificDayOfWeek *string `json:"specific_day_of_week,omitempty"`

	// A recognized mention of a specific month, represented numerically (for example, `7` for `July`).
	SpecificMonth *float64 `json:"specific_month,omitempty"`

	// A recognized mention of a specific quarter, represented numerically (for example, `3` for `the third quarter`).
	SpecificQuarter *float64 `json:"specific_quarter,omitempty"`

	// A recognized mention of a specific year (for example, `2016`).
	SpecificYear *float64 `json:"specific_year,omitempty"`

	// A recognized numeric value, represented as an integer or double.
	NumericValue *float64 `json:"numeric_value,omitempty"`

	// The type of numeric value recognized in the user input (`integer` or `rational`).
	Subtype *string `json:"subtype,omitempty"`

	// A recognized term for a time that was mentioned as a part of the day in the user's input (for example, `morning` or
	// `afternoon`).
	PartOfDay *string `json:"part_of_day,omitempty"`

	// A recognized mention of a relative hour, represented numerically as an offset from the current hour (for example,
	// `3` for `in three hours` or `-1` for `an hour ago`).
	RelativeHour *float64 `json:"relative_hour,omitempty"`

	// A recognized mention of a relative time, represented numerically as an offset in minutes from the current time (for
	// example, `5` for `in five minutes` or `-15` for `fifteen minutes ago`).
	RelativeMinute *float64 `json:"relative_minute,omitempty"`

	// A recognized mention of a relative time, represented numerically as an offset in seconds from the current time (for
	// example, `10` for `in ten seconds` or `-30` for `thirty seconds ago`).
	RelativeSecond *float64 `json:"relative_second,omitempty"`

	// A recognized specific hour mentioned as part of a time value (for example, `10` for `10:15 AM`.).
	SpecificHour *float64 `json:"specific_hour,omitempty"`

	// A recognized specific minute mentioned as part of a time value (for example, `15` for `10:15 AM`.).
	SpecificMinute *float64 `json:"specific_minute,omitempty"`

	// A recognized specific second mentioned as part of a time value (for example, `30` for `10:15:30 AM`.).
	SpecificSecond *float64 `json:"specific_second,omitempty"`

	// A recognized time zone mentioned as part of a time value (for example, `EST`).
	Timezone *string `json:"timezone,omitempty"`
}

// Constants associated with the RuntimeEntityInterpretation.Granularity property.
// The precision or duration of a time range specified by a recognized `@sys-time` or `@sys-date` entity.
const (
	RuntimeEntityInterpretationGranularityDayConst       = "day"
	RuntimeEntityInterpretationGranularityFortnightConst = "fortnight"
	RuntimeEntityInterpretationGranularityHourConst      = "hour"
	RuntimeEntityInterpretationGranularityInstantConst   = "instant"
	RuntimeEntityInterpretationGranularityMinuteConst    = "minute"
	RuntimeEntityInterpretationGranularityMonthConst     = "month"
	RuntimeEntityInterpretationGranularityQuarterConst   = "quarter"
	RuntimeEntityInterpretationGranularitySecondConst    = "second"
	RuntimeEntityInterpretationGranularityWeekConst      = "week"
	RuntimeEntityInterpretationGranularityWeekendConst   = "weekend"
	RuntimeEntityInterpretationGranularityYearConst      = "year"
)

// UnmarshalRuntimeEntityInterpretation unmarshals an instance of RuntimeEntityInterpretation from the specified map of raw messages.
func UnmarshalRuntimeEntityInterpretation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuntimeEntityInterpretation)
	err = core.UnmarshalPrimitive(m, "calendar_type", &obj.CalendarType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "datetime_link", &obj.DatetimeLink)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "festival", &obj.Festival)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "granularity", &obj.Granularity)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "range_link", &obj.RangeLink)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "range_modifier", &obj.RangeModifier)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "relative_day", &obj.RelativeDay)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "relative_month", &obj.RelativeMonth)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "relative_week", &obj.RelativeWeek)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "relative_weekend", &obj.RelativeWeekend)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "relative_year", &obj.RelativeYear)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "specific_day", &obj.SpecificDay)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "specific_day_of_week", &obj.SpecificDayOfWeek)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "specific_month", &obj.SpecificMonth)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "specific_quarter", &obj.SpecificQuarter)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "specific_year", &obj.SpecificYear)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "numeric_value", &obj.NumericValue)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "subtype", &obj.Subtype)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "part_of_day", &obj.PartOfDay)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "relative_hour", &obj.RelativeHour)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "relative_minute", &obj.RelativeMinute)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "relative_second", &obj.RelativeSecond)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "specific_hour", &obj.SpecificHour)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "specific_minute", &obj.SpecificMinute)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "specific_second", &obj.SpecificSecond)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "timezone", &obj.Timezone)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuntimeEntityRole : An object describing the role played by a system entity that is specifies the beginning or end of a range recognized
// in the user input. This property is included only if the new system entities are enabled for the workspace.
type RuntimeEntityRole struct {
	// The relationship of the entity to the range.
	Type *string `json:"type,omitempty"`
}

// Constants associated with the RuntimeEntityRole.Type property.
// The relationship of the entity to the range.
const (
	RuntimeEntityRoleTypeDateFromConst   = "date_from"
	RuntimeEntityRoleTypeDateToConst     = "date_to"
	RuntimeEntityRoleTypeNumberFromConst = "number_from"
	RuntimeEntityRoleTypeNumberToConst   = "number_to"
	RuntimeEntityRoleTypeTimeFromConst   = "time_from"
	RuntimeEntityRoleTypeTimeToConst     = "time_to"
)

// UnmarshalRuntimeEntityRole unmarshals an instance of RuntimeEntityRole from the specified map of raw messages.
func UnmarshalRuntimeEntityRole(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuntimeEntityRole)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuntimeIntent : An intent identified in the user input.
type RuntimeIntent struct {
	// The name of the recognized intent.
	Intent *string `json:"intent" validate:"required"`

	// A decimal percentage that represents Watson's confidence in the intent.
	Confidence *float64 `json:"confidence" validate:"required"`
}

// NewRuntimeIntent : Instantiate RuntimeIntent (Generic Model Constructor)
func (*AssistantV1) NewRuntimeIntent(intent string, confidence float64) (_model *RuntimeIntent, err error) {
	_model = &RuntimeIntent{
		Intent:     core.StringPtr(intent),
		Confidence: core.Float64Ptr(confidence),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalRuntimeIntent unmarshals an instance of RuntimeIntent from the specified map of raw messages.
func UnmarshalRuntimeIntent(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuntimeIntent)
	err = core.UnmarshalPrimitive(m, "intent", &obj.Intent)
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

// RuntimeResponseGeneric : RuntimeResponseGeneric struct
// Models which "extend" this model:
// - RuntimeResponseGenericRuntimeResponseTypeText
// - RuntimeResponseGenericRuntimeResponseTypePause
// - RuntimeResponseGenericRuntimeResponseTypeImage
// - RuntimeResponseGenericRuntimeResponseTypeOption
// - RuntimeResponseGenericRuntimeResponseTypeConnectToAgent
// - RuntimeResponseGenericRuntimeResponseTypeSuggestion
// - RuntimeResponseGenericRuntimeResponseTypeChannelTransfer
// - RuntimeResponseGenericRuntimeResponseTypeUserDefined
type RuntimeResponseGeneric struct {
	// The type of response returned by the dialog node. The specified response type must be supported by the client
	// application or channel.
	ResponseType *string `json:"response_type,omitempty"`

	// The text of the response.
	Text *string `json:"text,omitempty"`

	// An array of objects specifying channels for which the response is intended. If **channels** is present, the response
	// is intended for a built-in integration and should not be handled by an API client.
	Channels []ResponseGenericChannel `json:"channels,omitempty"`

	// How long to pause, in milliseconds.
	Time *int64 `json:"time,omitempty"`

	// Whether to send a "user is typing" event during the pause.
	Typing *bool `json:"typing,omitempty"`

	// The `https:` URL of the image.
	Source *string `json:"source,omitempty"`

	// The title or introductory text to show before the response.
	Title *string `json:"title,omitempty"`

	// The description to show with the the response.
	Description *string `json:"description,omitempty"`

	// Descriptive text that can be used for screen readers or other situations where the image cannot be seen.
	AltText *string `json:"alt_text,omitempty"`

	// The preferred type of control to display.
	Preference *string `json:"preference,omitempty"`

	// An array of objects describing the options from which the user can choose.
	Options []DialogNodeOutputOptionsElement `json:"options,omitempty"`

	// A message to be sent to the human agent who will be taking over the conversation.
	MessageToHumanAgent *string `json:"message_to_human_agent,omitempty"`

	// An optional message to be displayed to the user to indicate that the conversation will be transferred to the next
	// available agent.
	AgentAvailable *AgentAvailabilityMessage `json:"agent_available,omitempty"`

	// An optional message to be displayed to the user to indicate that no online agent is available to take over the
	// conversation.
	AgentUnavailable *AgentAvailabilityMessage `json:"agent_unavailable,omitempty"`

	// Routing or other contextual information to be used by target service desk systems.
	TransferInfo *DialogNodeOutputConnectToAgentTransferInfo `json:"transfer_info,omitempty"`

	// A label identifying the topic of the conversation, derived from the **title** property of the relevant node or the
	// **topic** property of the dialog node response.
	Topic *string `json:"topic,omitempty"`

	// The unique ID of the dialog node that the **topic** property is taken from. The **topic** property is populated
	// using the value of the dialog node's **title** property.
	DialogNode *string `json:"dialog_node,omitempty"`

	// An array of objects describing the possible matching dialog nodes from which the user can choose.
	Suggestions []DialogSuggestion `json:"suggestions,omitempty"`

	// The message to display to the user when initiating a channel transfer.
	MessageToUser *string `json:"message_to_user,omitempty"`

	// An object containing any properties for the user-defined response type.
	UserDefined map[string]interface{} `json:"user_defined,omitempty"`
}

// Constants associated with the RuntimeResponseGeneric.Preference property.
// The preferred type of control to display.
const (
	RuntimeResponseGenericPreferenceButtonConst   = "button"
	RuntimeResponseGenericPreferenceDropdownConst = "dropdown"
)

func (*RuntimeResponseGeneric) isaRuntimeResponseGeneric() bool {
	return true
}

type RuntimeResponseGenericIntf interface {
	isaRuntimeResponseGeneric() bool
}

// UnmarshalRuntimeResponseGeneric unmarshals an instance of RuntimeResponseGeneric from the specified map of raw messages.
func UnmarshalRuntimeResponseGeneric(m map[string]json.RawMessage, result interface{}) (err error) {
	// Retrieve discriminator value to determine correct "subclass".
	var discValue string
	err = core.UnmarshalPrimitive(m, "response_type", &discValue)
	if err != nil {
		err = fmt.Errorf("error unmarshalling discriminator property 'response_type': %s", err.Error())
		return
	}
	if discValue == "" {
		err = fmt.Errorf("required discriminator property 'response_type' not found in JSON object")
		return
	}
	if discValue == "channel_transfer" {
		err = core.UnmarshalModel(m, "", result, UnmarshalRuntimeResponseGenericRuntimeResponseTypeChannelTransfer)
	} else if discValue == "connect_to_agent" {
		err = core.UnmarshalModel(m, "", result, UnmarshalRuntimeResponseGenericRuntimeResponseTypeConnectToAgent)
	} else if discValue == "image" {
		err = core.UnmarshalModel(m, "", result, UnmarshalRuntimeResponseGenericRuntimeResponseTypeImage)
	} else if discValue == "option" {
		err = core.UnmarshalModel(m, "", result, UnmarshalRuntimeResponseGenericRuntimeResponseTypeOption)
	} else if discValue == "suggestion" {
		err = core.UnmarshalModel(m, "", result, UnmarshalRuntimeResponseGenericRuntimeResponseTypeSuggestion)
	} else if discValue == "pause" {
		err = core.UnmarshalModel(m, "", result, UnmarshalRuntimeResponseGenericRuntimeResponseTypePause)
	} else if discValue == "text" {
		err = core.UnmarshalModel(m, "", result, UnmarshalRuntimeResponseGenericRuntimeResponseTypeText)
	} else if discValue == "user_defined" {
		err = core.UnmarshalModel(m, "", result, UnmarshalRuntimeResponseGenericRuntimeResponseTypeUserDefined)
	} else {
		err = fmt.Errorf("unrecognized value for discriminator property 'response_type': %s", discValue)
	}
	return
}

// Synonym : Synonym struct
type Synonym struct {
	// The text of the synonym. This string must conform to the following restrictions:
	// - It cannot contain carriage return, newline, or tab characters.
	// - It cannot consist of only whitespace characters.
	Synonym *string `json:"synonym" validate:"required"`

	// The timestamp for creation of the object.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// The timestamp for the most recent update to the object.
	Updated *strfmt.DateTime `json:"updated,omitempty"`
}

// NewSynonym : Instantiate Synonym (Generic Model Constructor)
func (*AssistantV1) NewSynonym(synonym string) (_model *Synonym, err error) {
	_model = &Synonym{
		Synonym: core.StringPtr(synonym),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalSynonym unmarshals an instance of Synonym from the specified map of raw messages.
func UnmarshalSynonym(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Synonym)
	err = core.UnmarshalPrimitive(m, "synonym", &obj.Synonym)
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

// SynonymCollection : SynonymCollection struct
type SynonymCollection struct {
	// An array of synonyms.
	Synonyms []Synonym `json:"synonyms" validate:"required"`

	// The pagination data for the returned objects.
	Pagination *Pagination `json:"pagination" validate:"required"`
}

// UnmarshalSynonymCollection unmarshals an instance of SynonymCollection from the specified map of raw messages.
func UnmarshalSynonymCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SynonymCollection)
	err = core.UnmarshalModel(m, "synonyms", &obj.Synonyms, UnmarshalSynonym)
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

// UpdateCounterexampleOptions : The UpdateCounterexample options.
type UpdateCounterexampleOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// The text of a user input counterexample (for example, `What are you wearing?`).
	Text *string `json:"-" validate:"required,ne="`

	// The text of a user input marked as irrelevant input. This string must conform to the following restrictions:
	// - It cannot contain carriage return, newline, or tab characters.
	// - It cannot consist of only whitespace characters.
	NewText *string `json:"text,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateCounterexampleOptions : Instantiate UpdateCounterexampleOptions
func (*AssistantV1) NewUpdateCounterexampleOptions(workspaceID string, text string) *UpdateCounterexampleOptions {
	return &UpdateCounterexampleOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Text:        core.StringPtr(text),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *UpdateCounterexampleOptions) SetWorkspaceID(workspaceID string) *UpdateCounterexampleOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetText : Allow user to set Text
func (_options *UpdateCounterexampleOptions) SetText(text string) *UpdateCounterexampleOptions {
	_options.Text = core.StringPtr(text)
	return _options
}

// SetNewText : Allow user to set NewText
func (_options *UpdateCounterexampleOptions) SetNewText(newText string) *UpdateCounterexampleOptions {
	_options.NewText = core.StringPtr(newText)
	return _options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (_options *UpdateCounterexampleOptions) SetIncludeAudit(includeAudit bool) *UpdateCounterexampleOptions {
	_options.IncludeAudit = core.BoolPtr(includeAudit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateCounterexampleOptions) SetHeaders(param map[string]string) *UpdateCounterexampleOptions {
	options.Headers = param
	return options
}

// UpdateDialogNodeOptions : The UpdateDialogNode options.
type UpdateDialogNodeOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// The dialog node ID (for example, `node_1_1479323581900`).
	DialogNode *string `json:"-" validate:"required,ne="`

	// The unique ID of the dialog node. This is an internal identifier used to refer to the dialog node from other dialog
	// nodes and in the diagnostic information included with message responses.
	//
	// This string can contain only Unicode alphanumeric, space, underscore, hyphen, and dot characters.
	NewDialogNode *string `json:"dialog_node,omitempty"`

	// The description of the dialog node. This string cannot contain carriage return, newline, or tab characters.
	NewDescription *string `json:"description,omitempty"`

	// The condition that will trigger the dialog node. This string cannot contain carriage return, newline, or tab
	// characters.
	NewConditions *string `json:"conditions,omitempty"`

	// The unique ID of the parent dialog node. This property is omitted if the dialog node has no parent.
	NewParent *string `json:"parent,omitempty"`

	// The unique ID of the previous sibling dialog node. This property is omitted if the dialog node has no previous
	// sibling.
	NewPreviousSibling *string `json:"previous_sibling,omitempty"`

	// The output of the dialog node. For more information about how to specify dialog node output, see the
	// [documentation](https://cloud.ibm.com/docs/assistant?topic=assistant-dialog-overview#dialog-overview-responses).
	NewOutput *DialogNodeOutput `json:"output,omitempty"`

	// The context for the dialog node.
	NewContext *DialogNodeContext `json:"context,omitempty"`

	// The metadata for the dialog node.
	NewMetadata map[string]interface{} `json:"metadata,omitempty"`

	// The next step to execute following this dialog node.
	NewNextStep *DialogNodeNextStep `json:"next_step,omitempty"`

	// A human-readable name for the dialog node. If the node is included in disambiguation, this title is used to populate
	// the **label** property of the corresponding suggestion in the `suggestion` response type (unless it is overridden by
	// the **user_label** property). The title is also used to populate the **topic** property in the `connect_to_agent`
	// response type.
	//
	// This string can contain only Unicode alphanumeric, space, underscore, hyphen, and dot characters.
	NewTitle *string `json:"title,omitempty"`

	// How the dialog node is processed.
	NewType *string `json:"type,omitempty"`

	// How an `event_handler` node is processed.
	NewEventName *string `json:"event_name,omitempty"`

	// The location in the dialog context where output is stored.
	NewVariable *string `json:"variable,omitempty"`

	// An array of objects describing any actions to be invoked by the dialog node.
	NewActions []DialogNodeAction `json:"actions,omitempty"`

	// Whether this top-level dialog node can be digressed into.
	NewDigressIn *string `json:"digress_in,omitempty"`

	// Whether this dialog node can be returned to after a digression.
	NewDigressOut *string `json:"digress_out,omitempty"`

	// Whether the user can digress to top-level nodes while filling out slots.
	NewDigressOutSlots *string `json:"digress_out_slots,omitempty"`

	// A label that can be displayed externally to describe the purpose of the node to users. If set, this label is used to
	// identify the node in disambiguation responses (overriding the value of the **title** property).
	NewUserLabel *string `json:"user_label,omitempty"`

	// Whether the dialog node should be excluded from disambiguation suggestions. Valid only when **type**=`standard` or
	// `frame`.
	NewDisambiguationOptOut *bool `json:"disambiguation_opt_out,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateDialogNodeOptions.NewType property.
// How the dialog node is processed.
const (
	UpdateDialogNodeOptionsNewTypeEventHandlerConst      = "event_handler"
	UpdateDialogNodeOptionsNewTypeFolderConst            = "folder"
	UpdateDialogNodeOptionsNewTypeFrameConst             = "frame"
	UpdateDialogNodeOptionsNewTypeResponseConditionConst = "response_condition"
	UpdateDialogNodeOptionsNewTypeSlotConst              = "slot"
	UpdateDialogNodeOptionsNewTypeStandardConst          = "standard"
)

// Constants associated with the UpdateDialogNodeOptions.NewEventName property.
// How an `event_handler` node is processed.
const (
	UpdateDialogNodeOptionsNewEventNameDigressionReturnPromptConst   = "digression_return_prompt"
	UpdateDialogNodeOptionsNewEventNameFilledConst                   = "filled"
	UpdateDialogNodeOptionsNewEventNameFilledMultipleConst           = "filled_multiple"
	UpdateDialogNodeOptionsNewEventNameFocusConst                    = "focus"
	UpdateDialogNodeOptionsNewEventNameGenericConst                  = "generic"
	UpdateDialogNodeOptionsNewEventNameInputConst                    = "input"
	UpdateDialogNodeOptionsNewEventNameNomatchConst                  = "nomatch"
	UpdateDialogNodeOptionsNewEventNameNomatchResponsesDepletedConst = "nomatch_responses_depleted"
	UpdateDialogNodeOptionsNewEventNameValidateConst                 = "validate"
)

// Constants associated with the UpdateDialogNodeOptions.NewDigressIn property.
// Whether this top-level dialog node can be digressed into.
const (
	UpdateDialogNodeOptionsNewDigressInDoesNotReturnConst = "does_not_return"
	UpdateDialogNodeOptionsNewDigressInNotAvailableConst  = "not_available"
	UpdateDialogNodeOptionsNewDigressInReturnsConst       = "returns"
)

// Constants associated with the UpdateDialogNodeOptions.NewDigressOut property.
// Whether this dialog node can be returned to after a digression.
const (
	UpdateDialogNodeOptionsNewDigressOutAllowAllConst            = "allow_all"
	UpdateDialogNodeOptionsNewDigressOutAllowAllNeverReturnConst = "allow_all_never_return"
	UpdateDialogNodeOptionsNewDigressOutAllowReturningConst      = "allow_returning"
)

// Constants associated with the UpdateDialogNodeOptions.NewDigressOutSlots property.
// Whether the user can digress to top-level nodes while filling out slots.
const (
	UpdateDialogNodeOptionsNewDigressOutSlotsAllowAllConst       = "allow_all"
	UpdateDialogNodeOptionsNewDigressOutSlotsAllowReturningConst = "allow_returning"
	UpdateDialogNodeOptionsNewDigressOutSlotsNotAllowedConst     = "not_allowed"
)

// NewUpdateDialogNodeOptions : Instantiate UpdateDialogNodeOptions
func (*AssistantV1) NewUpdateDialogNodeOptions(workspaceID string, dialogNode string) *UpdateDialogNodeOptions {
	return &UpdateDialogNodeOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		DialogNode:  core.StringPtr(dialogNode),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *UpdateDialogNodeOptions) SetWorkspaceID(workspaceID string) *UpdateDialogNodeOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetDialogNode : Allow user to set DialogNode
func (_options *UpdateDialogNodeOptions) SetDialogNode(dialogNode string) *UpdateDialogNodeOptions {
	_options.DialogNode = core.StringPtr(dialogNode)
	return _options
}

// SetNewDialogNode : Allow user to set NewDialogNode
func (_options *UpdateDialogNodeOptions) SetNewDialogNode(newDialogNode string) *UpdateDialogNodeOptions {
	_options.NewDialogNode = core.StringPtr(newDialogNode)
	return _options
}

// SetNewDescription : Allow user to set NewDescription
func (_options *UpdateDialogNodeOptions) SetNewDescription(newDescription string) *UpdateDialogNodeOptions {
	_options.NewDescription = core.StringPtr(newDescription)
	return _options
}

// SetNewConditions : Allow user to set NewConditions
func (_options *UpdateDialogNodeOptions) SetNewConditions(newConditions string) *UpdateDialogNodeOptions {
	_options.NewConditions = core.StringPtr(newConditions)
	return _options
}

// SetNewParent : Allow user to set NewParent
func (_options *UpdateDialogNodeOptions) SetNewParent(newParent string) *UpdateDialogNodeOptions {
	_options.NewParent = core.StringPtr(newParent)
	return _options
}

// SetNewPreviousSibling : Allow user to set NewPreviousSibling
func (_options *UpdateDialogNodeOptions) SetNewPreviousSibling(newPreviousSibling string) *UpdateDialogNodeOptions {
	_options.NewPreviousSibling = core.StringPtr(newPreviousSibling)
	return _options
}

// SetNewOutput : Allow user to set NewOutput
func (_options *UpdateDialogNodeOptions) SetNewOutput(newOutput *DialogNodeOutput) *UpdateDialogNodeOptions {
	_options.NewOutput = newOutput
	return _options
}

// SetNewContext : Allow user to set NewContext
func (_options *UpdateDialogNodeOptions) SetNewContext(newContext *DialogNodeContext) *UpdateDialogNodeOptions {
	_options.NewContext = newContext
	return _options
}

// SetNewMetadata : Allow user to set NewMetadata
func (_options *UpdateDialogNodeOptions) SetNewMetadata(newMetadata map[string]interface{}) *UpdateDialogNodeOptions {
	_options.NewMetadata = newMetadata
	return _options
}

// SetNewNextStep : Allow user to set NewNextStep
func (_options *UpdateDialogNodeOptions) SetNewNextStep(newNextStep *DialogNodeNextStep) *UpdateDialogNodeOptions {
	_options.NewNextStep = newNextStep
	return _options
}

// SetNewTitle : Allow user to set NewTitle
func (_options *UpdateDialogNodeOptions) SetNewTitle(newTitle string) *UpdateDialogNodeOptions {
	_options.NewTitle = core.StringPtr(newTitle)
	return _options
}

// SetNewType : Allow user to set NewType
func (_options *UpdateDialogNodeOptions) SetNewType(newType string) *UpdateDialogNodeOptions {
	_options.NewType = core.StringPtr(newType)
	return _options
}

// SetNewEventName : Allow user to set NewEventName
func (_options *UpdateDialogNodeOptions) SetNewEventName(newEventName string) *UpdateDialogNodeOptions {
	_options.NewEventName = core.StringPtr(newEventName)
	return _options
}

// SetNewVariable : Allow user to set NewVariable
func (_options *UpdateDialogNodeOptions) SetNewVariable(newVariable string) *UpdateDialogNodeOptions {
	_options.NewVariable = core.StringPtr(newVariable)
	return _options
}

// SetNewActions : Allow user to set NewActions
func (_options *UpdateDialogNodeOptions) SetNewActions(newActions []DialogNodeAction) *UpdateDialogNodeOptions {
	_options.NewActions = newActions
	return _options
}

// SetNewDigressIn : Allow user to set NewDigressIn
func (_options *UpdateDialogNodeOptions) SetNewDigressIn(newDigressIn string) *UpdateDialogNodeOptions {
	_options.NewDigressIn = core.StringPtr(newDigressIn)
	return _options
}

// SetNewDigressOut : Allow user to set NewDigressOut
func (_options *UpdateDialogNodeOptions) SetNewDigressOut(newDigressOut string) *UpdateDialogNodeOptions {
	_options.NewDigressOut = core.StringPtr(newDigressOut)
	return _options
}

// SetNewDigressOutSlots : Allow user to set NewDigressOutSlots
func (_options *UpdateDialogNodeOptions) SetNewDigressOutSlots(newDigressOutSlots string) *UpdateDialogNodeOptions {
	_options.NewDigressOutSlots = core.StringPtr(newDigressOutSlots)
	return _options
}

// SetNewUserLabel : Allow user to set NewUserLabel
func (_options *UpdateDialogNodeOptions) SetNewUserLabel(newUserLabel string) *UpdateDialogNodeOptions {
	_options.NewUserLabel = core.StringPtr(newUserLabel)
	return _options
}

// SetNewDisambiguationOptOut : Allow user to set NewDisambiguationOptOut
func (_options *UpdateDialogNodeOptions) SetNewDisambiguationOptOut(newDisambiguationOptOut bool) *UpdateDialogNodeOptions {
	_options.NewDisambiguationOptOut = core.BoolPtr(newDisambiguationOptOut)
	return _options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (_options *UpdateDialogNodeOptions) SetIncludeAudit(includeAudit bool) *UpdateDialogNodeOptions {
	_options.IncludeAudit = core.BoolPtr(includeAudit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateDialogNodeOptions) SetHeaders(param map[string]string) *UpdateDialogNodeOptions {
	options.Headers = param
	return options
}

// UpdateEntityOptions : The UpdateEntity options.
type UpdateEntityOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// The name of the entity.
	Entity *string `json:"-" validate:"required,ne="`

	// The name of the entity. This string must conform to the following restrictions:
	// - It can contain only Unicode alphanumeric, underscore, and hyphen characters.
	// - It cannot begin with the reserved prefix `sys-`.
	NewEntity *string `json:"entity,omitempty"`

	// The description of the entity. This string cannot contain carriage return, newline, or tab characters.
	NewDescription *string `json:"description,omitempty"`

	// Any metadata related to the entity.
	NewMetadata map[string]interface{} `json:"metadata,omitempty"`

	// Whether to use fuzzy matching for the entity.
	NewFuzzyMatch *bool `json:"fuzzy_match,omitempty"`

	// An array of objects describing the entity values.
	NewValues []CreateValue `json:"values,omitempty"`

	// Whether the new data is to be appended to the existing data in the entity. If **append**=`false`, elements included
	// in the new data completely replace the corresponding existing elements, including all subelements. For example, if
	// the new data for the entity includes **values** and **append**=`false`, all existing values for the entity are
	// discarded and replaced with the new values.
	//
	// If **append**=`true`, existing elements are preserved, and the new elements are added. If any elements in the new
	// data collide with existing elements, the update request fails.
	Append *bool `json:"-"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateEntityOptions : Instantiate UpdateEntityOptions
func (*AssistantV1) NewUpdateEntityOptions(workspaceID string, entity string) *UpdateEntityOptions {
	return &UpdateEntityOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Entity:      core.StringPtr(entity),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *UpdateEntityOptions) SetWorkspaceID(workspaceID string) *UpdateEntityOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetEntity : Allow user to set Entity
func (_options *UpdateEntityOptions) SetEntity(entity string) *UpdateEntityOptions {
	_options.Entity = core.StringPtr(entity)
	return _options
}

// SetNewEntity : Allow user to set NewEntity
func (_options *UpdateEntityOptions) SetNewEntity(newEntity string) *UpdateEntityOptions {
	_options.NewEntity = core.StringPtr(newEntity)
	return _options
}

// SetNewDescription : Allow user to set NewDescription
func (_options *UpdateEntityOptions) SetNewDescription(newDescription string) *UpdateEntityOptions {
	_options.NewDescription = core.StringPtr(newDescription)
	return _options
}

// SetNewMetadata : Allow user to set NewMetadata
func (_options *UpdateEntityOptions) SetNewMetadata(newMetadata map[string]interface{}) *UpdateEntityOptions {
	_options.NewMetadata = newMetadata
	return _options
}

// SetNewFuzzyMatch : Allow user to set NewFuzzyMatch
func (_options *UpdateEntityOptions) SetNewFuzzyMatch(newFuzzyMatch bool) *UpdateEntityOptions {
	_options.NewFuzzyMatch = core.BoolPtr(newFuzzyMatch)
	return _options
}

// SetNewValues : Allow user to set NewValues
func (_options *UpdateEntityOptions) SetNewValues(newValues []CreateValue) *UpdateEntityOptions {
	_options.NewValues = newValues
	return _options
}

// SetAppend : Allow user to set Append
func (_options *UpdateEntityOptions) SetAppend(append bool) *UpdateEntityOptions {
	_options.Append = core.BoolPtr(append)
	return _options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (_options *UpdateEntityOptions) SetIncludeAudit(includeAudit bool) *UpdateEntityOptions {
	_options.IncludeAudit = core.BoolPtr(includeAudit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateEntityOptions) SetHeaders(param map[string]string) *UpdateEntityOptions {
	options.Headers = param
	return options
}

// UpdateExampleOptions : The UpdateExample options.
type UpdateExampleOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// The intent name.
	Intent *string `json:"-" validate:"required,ne="`

	// The text of the user input example.
	Text *string `json:"-" validate:"required,ne="`

	// The text of the user input example. This string must conform to the following restrictions:
	// - It cannot contain carriage return, newline, or tab characters.
	// - It cannot consist of only whitespace characters.
	NewText *string `json:"text,omitempty"`

	// An array of contextual entity mentions.
	NewMentions []Mention `json:"mentions,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateExampleOptions : Instantiate UpdateExampleOptions
func (*AssistantV1) NewUpdateExampleOptions(workspaceID string, intent string, text string) *UpdateExampleOptions {
	return &UpdateExampleOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Intent:      core.StringPtr(intent),
		Text:        core.StringPtr(text),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *UpdateExampleOptions) SetWorkspaceID(workspaceID string) *UpdateExampleOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetIntent : Allow user to set Intent
func (_options *UpdateExampleOptions) SetIntent(intent string) *UpdateExampleOptions {
	_options.Intent = core.StringPtr(intent)
	return _options
}

// SetText : Allow user to set Text
func (_options *UpdateExampleOptions) SetText(text string) *UpdateExampleOptions {
	_options.Text = core.StringPtr(text)
	return _options
}

// SetNewText : Allow user to set NewText
func (_options *UpdateExampleOptions) SetNewText(newText string) *UpdateExampleOptions {
	_options.NewText = core.StringPtr(newText)
	return _options
}

// SetNewMentions : Allow user to set NewMentions
func (_options *UpdateExampleOptions) SetNewMentions(newMentions []Mention) *UpdateExampleOptions {
	_options.NewMentions = newMentions
	return _options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (_options *UpdateExampleOptions) SetIncludeAudit(includeAudit bool) *UpdateExampleOptions {
	_options.IncludeAudit = core.BoolPtr(includeAudit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateExampleOptions) SetHeaders(param map[string]string) *UpdateExampleOptions {
	options.Headers = param
	return options
}

// UpdateIntentOptions : The UpdateIntent options.
type UpdateIntentOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// The intent name.
	Intent *string `json:"-" validate:"required,ne="`

	// The name of the intent. This string must conform to the following restrictions:
	// - It can contain only Unicode alphanumeric, underscore, hyphen, and dot characters.
	// - It cannot begin with the reserved prefix `sys-`.
	NewIntent *string `json:"intent,omitempty"`

	// The description of the intent. This string cannot contain carriage return, newline, or tab characters.
	NewDescription *string `json:"description,omitempty"`

	// An array of user input examples for the intent.
	NewExamples []Example `json:"examples,omitempty"`

	// Whether the new data is to be appended to the existing data in the object. If **append**=`false`, elements included
	// in the new data completely replace the corresponding existing elements, including all subelements. For example, if
	// the new data for the intent includes **examples** and **append**=`false`, all existing examples for the intent are
	// discarded and replaced with the new examples.
	//
	// If **append**=`true`, existing elements are preserved, and the new elements are added. If any elements in the new
	// data collide with existing elements, the update request fails.
	Append *bool `json:"-"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateIntentOptions : Instantiate UpdateIntentOptions
func (*AssistantV1) NewUpdateIntentOptions(workspaceID string, intent string) *UpdateIntentOptions {
	return &UpdateIntentOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Intent:      core.StringPtr(intent),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *UpdateIntentOptions) SetWorkspaceID(workspaceID string) *UpdateIntentOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetIntent : Allow user to set Intent
func (_options *UpdateIntentOptions) SetIntent(intent string) *UpdateIntentOptions {
	_options.Intent = core.StringPtr(intent)
	return _options
}

// SetNewIntent : Allow user to set NewIntent
func (_options *UpdateIntentOptions) SetNewIntent(newIntent string) *UpdateIntentOptions {
	_options.NewIntent = core.StringPtr(newIntent)
	return _options
}

// SetNewDescription : Allow user to set NewDescription
func (_options *UpdateIntentOptions) SetNewDescription(newDescription string) *UpdateIntentOptions {
	_options.NewDescription = core.StringPtr(newDescription)
	return _options
}

// SetNewExamples : Allow user to set NewExamples
func (_options *UpdateIntentOptions) SetNewExamples(newExamples []Example) *UpdateIntentOptions {
	_options.NewExamples = newExamples
	return _options
}

// SetAppend : Allow user to set Append
func (_options *UpdateIntentOptions) SetAppend(append bool) *UpdateIntentOptions {
	_options.Append = core.BoolPtr(append)
	return _options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (_options *UpdateIntentOptions) SetIncludeAudit(includeAudit bool) *UpdateIntentOptions {
	_options.IncludeAudit = core.BoolPtr(includeAudit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateIntentOptions) SetHeaders(param map[string]string) *UpdateIntentOptions {
	options.Headers = param
	return options
}

// UpdateSynonymOptions : The UpdateSynonym options.
type UpdateSynonymOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// The name of the entity.
	Entity *string `json:"-" validate:"required,ne="`

	// The text of the entity value.
	Value *string `json:"-" validate:"required,ne="`

	// The text of the synonym.
	Synonym *string `json:"-" validate:"required,ne="`

	// The text of the synonym. This string must conform to the following restrictions:
	// - It cannot contain carriage return, newline, or tab characters.
	// - It cannot consist of only whitespace characters.
	NewSynonym *string `json:"synonym,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateSynonymOptions : Instantiate UpdateSynonymOptions
func (*AssistantV1) NewUpdateSynonymOptions(workspaceID string, entity string, value string, synonym string) *UpdateSynonymOptions {
	return &UpdateSynonymOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Entity:      core.StringPtr(entity),
		Value:       core.StringPtr(value),
		Synonym:     core.StringPtr(synonym),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *UpdateSynonymOptions) SetWorkspaceID(workspaceID string) *UpdateSynonymOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetEntity : Allow user to set Entity
func (_options *UpdateSynonymOptions) SetEntity(entity string) *UpdateSynonymOptions {
	_options.Entity = core.StringPtr(entity)
	return _options
}

// SetValue : Allow user to set Value
func (_options *UpdateSynonymOptions) SetValue(value string) *UpdateSynonymOptions {
	_options.Value = core.StringPtr(value)
	return _options
}

// SetSynonym : Allow user to set Synonym
func (_options *UpdateSynonymOptions) SetSynonym(synonym string) *UpdateSynonymOptions {
	_options.Synonym = core.StringPtr(synonym)
	return _options
}

// SetNewSynonym : Allow user to set NewSynonym
func (_options *UpdateSynonymOptions) SetNewSynonym(newSynonym string) *UpdateSynonymOptions {
	_options.NewSynonym = core.StringPtr(newSynonym)
	return _options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (_options *UpdateSynonymOptions) SetIncludeAudit(includeAudit bool) *UpdateSynonymOptions {
	_options.IncludeAudit = core.BoolPtr(includeAudit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateSynonymOptions) SetHeaders(param map[string]string) *UpdateSynonymOptions {
	options.Headers = param
	return options
}

// UpdateValueOptions : The UpdateValue options.
type UpdateValueOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// The name of the entity.
	Entity *string `json:"-" validate:"required,ne="`

	// The text of the entity value.
	Value *string `json:"-" validate:"required,ne="`

	// The text of the entity value. This string must conform to the following restrictions:
	// - It cannot contain carriage return, newline, or tab characters.
	// - It cannot consist of only whitespace characters.
	NewValue *string `json:"value,omitempty"`

	// Any metadata related to the entity value.
	NewMetadata map[string]interface{} `json:"metadata,omitempty"`

	// Specifies the type of entity value.
	NewType *string `json:"type,omitempty"`

	// An array of synonyms for the entity value. A value can specify either synonyms or patterns (depending on the value
	// type), but not both. A synonym must conform to the following resrictions:
	// - It cannot contain carriage return, newline, or tab characters.
	// - It cannot consist of only whitespace characters.
	NewSynonyms []string `json:"synonyms,omitempty"`

	// An array of patterns for the entity value. A value can specify either synonyms or patterns (depending on the value
	// type), but not both. A pattern is a regular expression; for more information about how to specify a pattern, see the
	// [documentation](https://cloud.ibm.com/docs/assistant?topic=assistant-entities#entities-create-dictionary-based).
	NewPatterns []string `json:"patterns,omitempty"`

	// Whether the new data is to be appended to the existing data in the entity value. If **append**=`false`, elements
	// included in the new data completely replace the corresponding existing elements, including all subelements. For
	// example, if the new data for the entity value includes **synonyms** and **append**=`false`, all existing synonyms
	// for the entity value are discarded and replaced with the new synonyms.
	//
	// If **append**=`true`, existing elements are preserved, and the new elements are added. If any elements in the new
	// data collide with existing elements, the update request fails.
	Append *bool `json:"-"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateValueOptions.NewType property.
// Specifies the type of entity value.
const (
	UpdateValueOptionsNewTypePatternsConst = "patterns"
	UpdateValueOptionsNewTypeSynonymsConst = "synonyms"
)

// NewUpdateValueOptions : Instantiate UpdateValueOptions
func (*AssistantV1) NewUpdateValueOptions(workspaceID string, entity string, value string) *UpdateValueOptions {
	return &UpdateValueOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Entity:      core.StringPtr(entity),
		Value:       core.StringPtr(value),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *UpdateValueOptions) SetWorkspaceID(workspaceID string) *UpdateValueOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetEntity : Allow user to set Entity
func (_options *UpdateValueOptions) SetEntity(entity string) *UpdateValueOptions {
	_options.Entity = core.StringPtr(entity)
	return _options
}

// SetValue : Allow user to set Value
func (_options *UpdateValueOptions) SetValue(value string) *UpdateValueOptions {
	_options.Value = core.StringPtr(value)
	return _options
}

// SetNewValue : Allow user to set NewValue
func (_options *UpdateValueOptions) SetNewValue(newValue string) *UpdateValueOptions {
	_options.NewValue = core.StringPtr(newValue)
	return _options
}

// SetNewMetadata : Allow user to set NewMetadata
func (_options *UpdateValueOptions) SetNewMetadata(newMetadata map[string]interface{}) *UpdateValueOptions {
	_options.NewMetadata = newMetadata
	return _options
}

// SetNewType : Allow user to set NewType
func (_options *UpdateValueOptions) SetNewType(newType string) *UpdateValueOptions {
	_options.NewType = core.StringPtr(newType)
	return _options
}

// SetNewSynonyms : Allow user to set NewSynonyms
func (_options *UpdateValueOptions) SetNewSynonyms(newSynonyms []string) *UpdateValueOptions {
	_options.NewSynonyms = newSynonyms
	return _options
}

// SetNewPatterns : Allow user to set NewPatterns
func (_options *UpdateValueOptions) SetNewPatterns(newPatterns []string) *UpdateValueOptions {
	_options.NewPatterns = newPatterns
	return _options
}

// SetAppend : Allow user to set Append
func (_options *UpdateValueOptions) SetAppend(append bool) *UpdateValueOptions {
	_options.Append = core.BoolPtr(append)
	return _options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (_options *UpdateValueOptions) SetIncludeAudit(includeAudit bool) *UpdateValueOptions {
	_options.IncludeAudit = core.BoolPtr(includeAudit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateValueOptions) SetHeaders(param map[string]string) *UpdateValueOptions {
	options.Headers = param
	return options
}

// UpdateWorkspaceOptions : The UpdateWorkspace options.
type UpdateWorkspaceOptions struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `json:"-" validate:"required,ne="`

	// The name of the workspace. This string cannot contain carriage return, newline, or tab characters.
	Name *string `json:"name,omitempty"`

	// The description of the workspace. This string cannot contain carriage return, newline, or tab characters.
	Description *string `json:"description,omitempty"`

	// The language of the workspace.
	Language *string `json:"language,omitempty"`

	// An array of objects describing the dialog nodes in the workspace.
	DialogNodes []DialogNode `json:"dialog_nodes,omitempty"`

	// An array of objects defining input examples that have been marked as irrelevant input.
	Counterexamples []Counterexample `json:"counterexamples,omitempty"`

	// Any metadata related to the workspace.
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// Whether training data from the workspace (including artifacts such as intents and entities) can be used by IBM for
	// general service improvements. `true` indicates that workspace training data is not to be used.
	LearningOptOut *bool `json:"learning_opt_out,omitempty"`

	// Global settings for the workspace.
	SystemSettings *WorkspaceSystemSettings `json:"system_settings,omitempty"`

	Webhooks []Webhook `json:"webhooks,omitempty"`

	// An array of objects defining the intents for the workspace.
	Intents []CreateIntent `json:"intents,omitempty"`

	// An array of objects describing the entities for the workspace.
	Entities []CreateEntity `json:"entities,omitempty"`

	// Whether the new data is to be appended to the existing data in the object. If **append**=`false`, elements included
	// in the new data completely replace the corresponding existing elements, including all subelements. For example, if
	// the new data for a workspace includes **entities** and **append**=`false`, all existing entities in the workspace
	// are discarded and replaced with the new entities.
	//
	// If **append**=`true`, existing elements are preserved, and the new elements are added. If any elements in the new
	// data collide with existing elements, the update request fails.
	Append *bool `json:"-"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateWorkspaceOptions : Instantiate UpdateWorkspaceOptions
func (*AssistantV1) NewUpdateWorkspaceOptions(workspaceID string) *UpdateWorkspaceOptions {
	return &UpdateWorkspaceOptions{
		WorkspaceID: core.StringPtr(workspaceID),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (_options *UpdateWorkspaceOptions) SetWorkspaceID(workspaceID string) *UpdateWorkspaceOptions {
	_options.WorkspaceID = core.StringPtr(workspaceID)
	return _options
}

// SetName : Allow user to set Name
func (_options *UpdateWorkspaceOptions) SetName(name string) *UpdateWorkspaceOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *UpdateWorkspaceOptions) SetDescription(description string) *UpdateWorkspaceOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetLanguage : Allow user to set Language
func (_options *UpdateWorkspaceOptions) SetLanguage(language string) *UpdateWorkspaceOptions {
	_options.Language = core.StringPtr(language)
	return _options
}

// SetDialogNodes : Allow user to set DialogNodes
func (_options *UpdateWorkspaceOptions) SetDialogNodes(dialogNodes []DialogNode) *UpdateWorkspaceOptions {
	_options.DialogNodes = dialogNodes
	return _options
}

// SetCounterexamples : Allow user to set Counterexamples
func (_options *UpdateWorkspaceOptions) SetCounterexamples(counterexamples []Counterexample) *UpdateWorkspaceOptions {
	_options.Counterexamples = counterexamples
	return _options
}

// SetMetadata : Allow user to set Metadata
func (_options *UpdateWorkspaceOptions) SetMetadata(metadata map[string]interface{}) *UpdateWorkspaceOptions {
	_options.Metadata = metadata
	return _options
}

// SetLearningOptOut : Allow user to set LearningOptOut
func (_options *UpdateWorkspaceOptions) SetLearningOptOut(learningOptOut bool) *UpdateWorkspaceOptions {
	_options.LearningOptOut = core.BoolPtr(learningOptOut)
	return _options
}

// SetSystemSettings : Allow user to set SystemSettings
func (_options *UpdateWorkspaceOptions) SetSystemSettings(systemSettings *WorkspaceSystemSettings) *UpdateWorkspaceOptions {
	_options.SystemSettings = systemSettings
	return _options
}

// SetWebhooks : Allow user to set Webhooks
func (_options *UpdateWorkspaceOptions) SetWebhooks(webhooks []Webhook) *UpdateWorkspaceOptions {
	_options.Webhooks = webhooks
	return _options
}

// SetIntents : Allow user to set Intents
func (_options *UpdateWorkspaceOptions) SetIntents(intents []CreateIntent) *UpdateWorkspaceOptions {
	_options.Intents = intents
	return _options
}

// SetEntities : Allow user to set Entities
func (_options *UpdateWorkspaceOptions) SetEntities(entities []CreateEntity) *UpdateWorkspaceOptions {
	_options.Entities = entities
	return _options
}

// SetAppend : Allow user to set Append
func (_options *UpdateWorkspaceOptions) SetAppend(append bool) *UpdateWorkspaceOptions {
	_options.Append = core.BoolPtr(append)
	return _options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (_options *UpdateWorkspaceOptions) SetIncludeAudit(includeAudit bool) *UpdateWorkspaceOptions {
	_options.IncludeAudit = core.BoolPtr(includeAudit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateWorkspaceOptions) SetHeaders(param map[string]string) *UpdateWorkspaceOptions {
	options.Headers = param
	return options
}

// Value : Value struct
type Value struct {
	// The text of the entity value. This string must conform to the following restrictions:
	// - It cannot contain carriage return, newline, or tab characters.
	// - It cannot consist of only whitespace characters.
	Value *string `json:"value" validate:"required"`

	// Any metadata related to the entity value.
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// Specifies the type of entity value.
	Type *string `json:"type" validate:"required"`

	// An array of synonyms for the entity value. A value can specify either synonyms or patterns (depending on the value
	// type), but not both. A synonym must conform to the following resrictions:
	// - It cannot contain carriage return, newline, or tab characters.
	// - It cannot consist of only whitespace characters.
	Synonyms []string `json:"synonyms,omitempty"`

	// An array of patterns for the entity value. A value can specify either synonyms or patterns (depending on the value
	// type), but not both. A pattern is a regular expression; for more information about how to specify a pattern, see the
	// [documentation](https://cloud.ibm.com/docs/assistant?topic=assistant-entities#entities-create-dictionary-based).
	Patterns []string `json:"patterns,omitempty"`

	// The timestamp for creation of the object.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// The timestamp for the most recent update to the object.
	Updated *strfmt.DateTime `json:"updated,omitempty"`
}

// Constants associated with the Value.Type property.
// Specifies the type of entity value.
const (
	ValueTypePatternsConst = "patterns"
	ValueTypeSynonymsConst = "synonyms"
)

// UnmarshalValue unmarshals an instance of Value from the specified map of raw messages.
func UnmarshalValue(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Value)
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "metadata", &obj.Metadata)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "synonyms", &obj.Synonyms)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "patterns", &obj.Patterns)
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

// ValueCollection : ValueCollection struct
type ValueCollection struct {
	// An array of entity values.
	Values []Value `json:"values" validate:"required"`

	// The pagination data for the returned objects.
	Pagination *Pagination `json:"pagination" validate:"required"`
}

// UnmarshalValueCollection unmarshals an instance of ValueCollection from the specified map of raw messages.
func UnmarshalValueCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ValueCollection)
	err = core.UnmarshalModel(m, "values", &obj.Values, UnmarshalValue)
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

// Webhook : A webhook that can be used by dialog nodes to make programmatic calls to an external function.
//
// **Note:** Currently, only a single webhook named `main_webhook` is supported.
type Webhook struct {
	// The URL for the external service or application to which you want to send HTTP POST requests.
	URL *string `json:"url" validate:"required"`

	// The name of the webhook. Currently, `main_webhook` is the only supported value.
	Name *string `json:"name" validate:"required"`

	// An optional array of HTTP headers to pass with the HTTP request.
	HeadersVar []WebhookHeader `json:"headers,omitempty"`
}

// NewWebhook : Instantiate Webhook (Generic Model Constructor)
func (*AssistantV1) NewWebhook(url string, name string) (_model *Webhook, err error) {
	_model = &Webhook{
		URL:  core.StringPtr(url),
		Name: core.StringPtr(name),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalWebhook unmarshals an instance of Webhook from the specified map of raw messages.
func UnmarshalWebhook(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Webhook)
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "headers", &obj.HeadersVar, UnmarshalWebhookHeader)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// WebhookHeader : A key/value pair defining an HTTP header and a value.
type WebhookHeader struct {
	// The name of an HTTP header (for example, `Authorization`).
	Name *string `json:"name" validate:"required"`

	// The value of an HTTP header.
	Value *string `json:"value" validate:"required"`
}

// NewWebhookHeader : Instantiate WebhookHeader (Generic Model Constructor)
func (*AssistantV1) NewWebhookHeader(name string, value string) (_model *WebhookHeader, err error) {
	_model = &WebhookHeader{
		Name:  core.StringPtr(name),
		Value: core.StringPtr(value),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalWebhookHeader unmarshals an instance of WebhookHeader from the specified map of raw messages.
func UnmarshalWebhookHeader(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(WebhookHeader)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
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

// Workspace : Workspace struct
type Workspace struct {
	// The name of the workspace. This string cannot contain carriage return, newline, or tab characters.
	Name *string `json:"name" validate:"required"`

	// The description of the workspace. This string cannot contain carriage return, newline, or tab characters.
	Description *string `json:"description,omitempty"`

	// The language of the workspace.
	Language *string `json:"language" validate:"required"`

	// The workspace ID of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// An array of objects describing the dialog nodes in the workspace.
	DialogNodes []DialogNode `json:"dialog_nodes,omitempty"`

	// An array of objects defining input examples that have been marked as irrelevant input.
	Counterexamples []Counterexample `json:"counterexamples,omitempty"`

	// The timestamp for creation of the object.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// The timestamp for the most recent update to the object.
	Updated *strfmt.DateTime `json:"updated,omitempty"`

	// Any metadata related to the workspace.
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// Whether training data from the workspace (including artifacts such as intents and entities) can be used by IBM for
	// general service improvements. `true` indicates that workspace training data is not to be used.
	LearningOptOut *bool `json:"learning_opt_out" validate:"required"`

	// Global settings for the workspace.
	SystemSettings *WorkspaceSystemSettings `json:"system_settings,omitempty"`

	// The current status of the workspace.
	Status *string `json:"status,omitempty"`

	Webhooks []Webhook `json:"webhooks,omitempty"`

	// An array of intents.
	Intents []Intent `json:"intents,omitempty"`

	// An array of objects describing the entities for the workspace.
	Entities []Entity `json:"entities,omitempty"`
}

// Constants associated with the Workspace.Status property.
// The current status of the workspace.
const (
	WorkspaceStatusAvailableConst   = "Available"
	WorkspaceStatusFailedConst      = "Failed"
	WorkspaceStatusNonExistentConst = "Non Existent"
	WorkspaceStatusTrainingConst    = "Training"
	WorkspaceStatusUnavailableConst = "Unavailable"
)

// UnmarshalWorkspace unmarshals an instance of Workspace from the specified map of raw messages.
func UnmarshalWorkspace(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Workspace)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "language", &obj.Language)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "workspace_id", &obj.WorkspaceID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "dialog_nodes", &obj.DialogNodes, UnmarshalDialogNode)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "counterexamples", &obj.Counterexamples, UnmarshalCounterexample)
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
	err = core.UnmarshalPrimitive(m, "metadata", &obj.Metadata)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "learning_opt_out", &obj.LearningOptOut)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "system_settings", &obj.SystemSettings, UnmarshalWorkspaceSystemSettings)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "webhooks", &obj.Webhooks, UnmarshalWebhook)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "intents", &obj.Intents, UnmarshalIntent)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "entities", &obj.Entities, UnmarshalEntity)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// WorkspaceCollection : WorkspaceCollection struct
type WorkspaceCollection struct {
	// An array of objects describing the workspaces associated with the service instance.
	Workspaces []Workspace `json:"workspaces" validate:"required"`

	// The pagination data for the returned objects.
	Pagination *Pagination `json:"pagination" validate:"required"`
}

// UnmarshalWorkspaceCollection unmarshals an instance of WorkspaceCollection from the specified map of raw messages.
func UnmarshalWorkspaceCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(WorkspaceCollection)
	err = core.UnmarshalModel(m, "workspaces", &obj.Workspaces, UnmarshalWorkspace)
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

// WorkspaceSystemSettings : Global settings for the workspace.
type WorkspaceSystemSettings struct {
	// Workspace settings related to the Watson Assistant user interface.
	Tooling *WorkspaceSystemSettingsTooling `json:"tooling,omitempty"`

	// Workspace settings related to the disambiguation feature.
	Disambiguation *WorkspaceSystemSettingsDisambiguation `json:"disambiguation,omitempty"`

	// For internal use only.
	HumanAgentAssist map[string]interface{} `json:"human_agent_assist,omitempty"`

	// Whether spelling correction is enabled for the workspace.
	SpellingSuggestions *bool `json:"spelling_suggestions,omitempty"`

	// Whether autocorrection is enabled for the workspace. If spelling correction is enabled and this property is `false`,
	// any suggested corrections are returned in the **suggested_text** property of the message response. If this property
	// is `true`, any corrections are automatically applied to the user input, and the original text is returned in the
	// **original_text** property of the message response.
	SpellingAutoCorrect *bool `json:"spelling_auto_correct,omitempty"`

	// Workspace settings related to the behavior of system entities.
	SystemEntities *WorkspaceSystemSettingsSystemEntities `json:"system_entities,omitempty"`

	// Workspace settings related to detection of irrelevant input.
	OffTopic *WorkspaceSystemSettingsOffTopic `json:"off_topic,omitempty"`
}

// UnmarshalWorkspaceSystemSettings unmarshals an instance of WorkspaceSystemSettings from the specified map of raw messages.
func UnmarshalWorkspaceSystemSettings(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(WorkspaceSystemSettings)
	err = core.UnmarshalModel(m, "tooling", &obj.Tooling, UnmarshalWorkspaceSystemSettingsTooling)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "disambiguation", &obj.Disambiguation, UnmarshalWorkspaceSystemSettingsDisambiguation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "human_agent_assist", &obj.HumanAgentAssist)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spelling_suggestions", &obj.SpellingSuggestions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spelling_auto_correct", &obj.SpellingAutoCorrect)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "system_entities", &obj.SystemEntities, UnmarshalWorkspaceSystemSettingsSystemEntities)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "off_topic", &obj.OffTopic, UnmarshalWorkspaceSystemSettingsOffTopic)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// WorkspaceSystemSettingsDisambiguation : Workspace settings related to the disambiguation feature.
type WorkspaceSystemSettingsDisambiguation struct {
	// The text of the introductory prompt that accompanies disambiguation options presented to the user.
	Prompt *string `json:"prompt,omitempty"`

	// The user-facing label for the option users can select if none of the suggested options is correct. If no value is
	// specified for this property, this option does not appear.
	NoneOfTheAbovePrompt *string `json:"none_of_the_above_prompt,omitempty"`

	// Whether the disambiguation feature is enabled for the workspace.
	Enabled *bool `json:"enabled,omitempty"`

	// The sensitivity of the disambiguation feature to intent detection uncertainty. Higher sensitivity means that the
	// disambiguation feature is triggered more often and includes more choices.
	Sensitivity *string `json:"sensitivity,omitempty"`

	// Whether the order in which disambiguation suggestions are presented should be randomized (but still influenced by
	// relative confidence).
	Randomize *bool `json:"randomize,omitempty"`

	// The maximum number of disambigation suggestions that can be included in a `suggestion` response.
	MaxSuggestions *int64 `json:"max_suggestions,omitempty"`

	// For internal use only.
	SuggestionTextPolicy *string `json:"suggestion_text_policy,omitempty"`
}

// Constants associated with the WorkspaceSystemSettingsDisambiguation.Sensitivity property.
// The sensitivity of the disambiguation feature to intent detection uncertainty. Higher sensitivity means that the
// disambiguation feature is triggered more often and includes more choices.
const (
	WorkspaceSystemSettingsDisambiguationSensitivityAutoConst       = "auto"
	WorkspaceSystemSettingsDisambiguationSensitivityHighConst       = "high"
	WorkspaceSystemSettingsDisambiguationSensitivityLowConst        = "low"
	WorkspaceSystemSettingsDisambiguationSensitivityMediumConst     = "medium"
	WorkspaceSystemSettingsDisambiguationSensitivityMediumHighConst = "medium_high"
	WorkspaceSystemSettingsDisambiguationSensitivityMediumLowConst  = "medium_low"
)

// UnmarshalWorkspaceSystemSettingsDisambiguation unmarshals an instance of WorkspaceSystemSettingsDisambiguation from the specified map of raw messages.
func UnmarshalWorkspaceSystemSettingsDisambiguation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(WorkspaceSystemSettingsDisambiguation)
	err = core.UnmarshalPrimitive(m, "prompt", &obj.Prompt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "none_of_the_above_prompt", &obj.NoneOfTheAbovePrompt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sensitivity", &obj.Sensitivity)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "randomize", &obj.Randomize)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "max_suggestions", &obj.MaxSuggestions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "suggestion_text_policy", &obj.SuggestionTextPolicy)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// WorkspaceSystemSettingsOffTopic : Workspace settings related to detection of irrelevant input.
type WorkspaceSystemSettingsOffTopic struct {
	// Whether enhanced irrelevance detection is enabled for the workspace.
	Enabled *bool `json:"enabled,omitempty"`
}

// UnmarshalWorkspaceSystemSettingsOffTopic unmarshals an instance of WorkspaceSystemSettingsOffTopic from the specified map of raw messages.
func UnmarshalWorkspaceSystemSettingsOffTopic(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(WorkspaceSystemSettingsOffTopic)
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// WorkspaceSystemSettingsSystemEntities : Workspace settings related to the behavior of system entities.
type WorkspaceSystemSettingsSystemEntities struct {
	// Whether the new system entities are enabled for the workspace.
	Enabled *bool `json:"enabled,omitempty"`
}

// UnmarshalWorkspaceSystemSettingsSystemEntities unmarshals an instance of WorkspaceSystemSettingsSystemEntities from the specified map of raw messages.
func UnmarshalWorkspaceSystemSettingsSystemEntities(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(WorkspaceSystemSettingsSystemEntities)
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// WorkspaceSystemSettingsTooling : Workspace settings related to the Watson Assistant user interface.
type WorkspaceSystemSettingsTooling struct {
	// Whether the dialog JSON editor displays text responses within the `output.generic` object.
	StoreGenericResponses *bool `json:"store_generic_responses,omitempty"`
}

// UnmarshalWorkspaceSystemSettingsTooling unmarshals an instance of WorkspaceSystemSettingsTooling from the specified map of raw messages.
func UnmarshalWorkspaceSystemSettingsTooling(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(WorkspaceSystemSettingsTooling)
	err = core.UnmarshalPrimitive(m, "store_generic_responses", &obj.StoreGenericResponses)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DialogNodeOutputGenericDialogNodeOutputResponseTypeChannelTransfer : DialogNodeOutputGenericDialogNodeOutputResponseTypeChannelTransfer struct
// This model "extends" DialogNodeOutputGeneric
type DialogNodeOutputGenericDialogNodeOutputResponseTypeChannelTransfer struct {
	// The type of response returned by the dialog node. The specified response type must be supported by the client
	// application or channel.
	ResponseType *string `json:"response_type" validate:"required"`

	// The message to display to the user when initiating a channel transfer.
	MessageToUser *string `json:"message_to_user" validate:"required"`

	// Information used by an integration to transfer the conversation to a different channel.
	TransferInfo *ChannelTransferInfo `json:"transfer_info" validate:"required"`

	// An array of objects specifying channels for which the response is intended.
	Channels []ResponseGenericChannel `json:"channels,omitempty"`
}

// NewDialogNodeOutputGenericDialogNodeOutputResponseTypeChannelTransfer : Instantiate DialogNodeOutputGenericDialogNodeOutputResponseTypeChannelTransfer (Generic Model Constructor)
func (*AssistantV1) NewDialogNodeOutputGenericDialogNodeOutputResponseTypeChannelTransfer(responseType string, messageToUser string, transferInfo *ChannelTransferInfo) (_model *DialogNodeOutputGenericDialogNodeOutputResponseTypeChannelTransfer, err error) {
	_model = &DialogNodeOutputGenericDialogNodeOutputResponseTypeChannelTransfer{
		ResponseType:  core.StringPtr(responseType),
		MessageToUser: core.StringPtr(messageToUser),
		TransferInfo:  transferInfo,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*DialogNodeOutputGenericDialogNodeOutputResponseTypeChannelTransfer) isaDialogNodeOutputGeneric() bool {
	return true
}

// UnmarshalDialogNodeOutputGenericDialogNodeOutputResponseTypeChannelTransfer unmarshals an instance of DialogNodeOutputGenericDialogNodeOutputResponseTypeChannelTransfer from the specified map of raw messages.
func UnmarshalDialogNodeOutputGenericDialogNodeOutputResponseTypeChannelTransfer(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DialogNodeOutputGenericDialogNodeOutputResponseTypeChannelTransfer)
	err = core.UnmarshalPrimitive(m, "response_type", &obj.ResponseType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "message_to_user", &obj.MessageToUser)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "transfer_info", &obj.TransferInfo, UnmarshalChannelTransferInfo)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "channels", &obj.Channels, UnmarshalResponseGenericChannel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DialogNodeOutputGenericDialogNodeOutputResponseTypeConnectToAgent : DialogNodeOutputGenericDialogNodeOutputResponseTypeConnectToAgent struct
// This model "extends" DialogNodeOutputGeneric
type DialogNodeOutputGenericDialogNodeOutputResponseTypeConnectToAgent struct {
	// The type of response returned by the dialog node. The specified response type must be supported by the client
	// application or channel.
	ResponseType *string `json:"response_type" validate:"required"`

	// An optional message to be sent to the human agent who will be taking over the conversation.
	MessageToHumanAgent *string `json:"message_to_human_agent,omitempty"`

	// An optional message to be displayed to the user to indicate that the conversation will be transferred to the next
	// available agent.
	AgentAvailable *AgentAvailabilityMessage `json:"agent_available,omitempty"`

	// An optional message to be displayed to the user to indicate that no online agent is available to take over the
	// conversation.
	AgentUnavailable *AgentAvailabilityMessage `json:"agent_unavailable,omitempty"`

	// Routing or other contextual information to be used by target service desk systems.
	TransferInfo *DialogNodeOutputConnectToAgentTransferInfo `json:"transfer_info,omitempty"`

	// An array of objects specifying channels for which the response is intended.
	Channels []ResponseGenericChannel `json:"channels,omitempty"`
}

// NewDialogNodeOutputGenericDialogNodeOutputResponseTypeConnectToAgent : Instantiate DialogNodeOutputGenericDialogNodeOutputResponseTypeConnectToAgent (Generic Model Constructor)
func (*AssistantV1) NewDialogNodeOutputGenericDialogNodeOutputResponseTypeConnectToAgent(responseType string) (_model *DialogNodeOutputGenericDialogNodeOutputResponseTypeConnectToAgent, err error) {
	_model = &DialogNodeOutputGenericDialogNodeOutputResponseTypeConnectToAgent{
		ResponseType: core.StringPtr(responseType),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*DialogNodeOutputGenericDialogNodeOutputResponseTypeConnectToAgent) isaDialogNodeOutputGeneric() bool {
	return true
}

// UnmarshalDialogNodeOutputGenericDialogNodeOutputResponseTypeConnectToAgent unmarshals an instance of DialogNodeOutputGenericDialogNodeOutputResponseTypeConnectToAgent from the specified map of raw messages.
func UnmarshalDialogNodeOutputGenericDialogNodeOutputResponseTypeConnectToAgent(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DialogNodeOutputGenericDialogNodeOutputResponseTypeConnectToAgent)
	err = core.UnmarshalPrimitive(m, "response_type", &obj.ResponseType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "message_to_human_agent", &obj.MessageToHumanAgent)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "agent_available", &obj.AgentAvailable, UnmarshalAgentAvailabilityMessage)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "agent_unavailable", &obj.AgentUnavailable, UnmarshalAgentAvailabilityMessage)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "transfer_info", &obj.TransferInfo, UnmarshalDialogNodeOutputConnectToAgentTransferInfo)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "channels", &obj.Channels, UnmarshalResponseGenericChannel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DialogNodeOutputGenericDialogNodeOutputResponseTypeImage : DialogNodeOutputGenericDialogNodeOutputResponseTypeImage struct
// This model "extends" DialogNodeOutputGeneric
type DialogNodeOutputGenericDialogNodeOutputResponseTypeImage struct {
	// The type of response returned by the dialog node. The specified response type must be supported by the client
	// application or channel.
	ResponseType *string `json:"response_type" validate:"required"`

	// The `https:` URL of the image.
	Source *string `json:"source" validate:"required"`

	// An optional title to show before the response.
	Title *string `json:"title,omitempty"`

	// An optional description to show with the response.
	Description *string `json:"description,omitempty"`

	// An array of objects specifying channels for which the response is intended.
	Channels []ResponseGenericChannel `json:"channels,omitempty"`

	// Descriptive text that can be used for screen readers or other situations where the image cannot be seen.
	AltText *string `json:"alt_text,omitempty"`
}

// NewDialogNodeOutputGenericDialogNodeOutputResponseTypeImage : Instantiate DialogNodeOutputGenericDialogNodeOutputResponseTypeImage (Generic Model Constructor)
func (*AssistantV1) NewDialogNodeOutputGenericDialogNodeOutputResponseTypeImage(responseType string, source string) (_model *DialogNodeOutputGenericDialogNodeOutputResponseTypeImage, err error) {
	_model = &DialogNodeOutputGenericDialogNodeOutputResponseTypeImage{
		ResponseType: core.StringPtr(responseType),
		Source:       core.StringPtr(source),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*DialogNodeOutputGenericDialogNodeOutputResponseTypeImage) isaDialogNodeOutputGeneric() bool {
	return true
}

// UnmarshalDialogNodeOutputGenericDialogNodeOutputResponseTypeImage unmarshals an instance of DialogNodeOutputGenericDialogNodeOutputResponseTypeImage from the specified map of raw messages.
func UnmarshalDialogNodeOutputGenericDialogNodeOutputResponseTypeImage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DialogNodeOutputGenericDialogNodeOutputResponseTypeImage)
	err = core.UnmarshalPrimitive(m, "response_type", &obj.ResponseType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source", &obj.Source)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "title", &obj.Title)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "channels", &obj.Channels, UnmarshalResponseGenericChannel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "alt_text", &obj.AltText)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DialogNodeOutputGenericDialogNodeOutputResponseTypeOption : DialogNodeOutputGenericDialogNodeOutputResponseTypeOption struct
// This model "extends" DialogNodeOutputGeneric
type DialogNodeOutputGenericDialogNodeOutputResponseTypeOption struct {
	// The type of response returned by the dialog node. The specified response type must be supported by the client
	// application or channel.
	ResponseType *string `json:"response_type" validate:"required"`

	// An optional title to show before the response.
	Title *string `json:"title" validate:"required"`

	// An optional description to show with the response.
	Description *string `json:"description,omitempty"`

	// The preferred type of control to display, if supported by the channel.
	Preference *string `json:"preference,omitempty"`

	// An array of objects describing the options from which the user can choose. You can include up to 20 options.
	Options []DialogNodeOutputOptionsElement `json:"options" validate:"required"`

	// An array of objects specifying channels for which the response is intended.
	Channels []ResponseGenericChannel `json:"channels,omitempty"`
}

// Constants associated with the DialogNodeOutputGenericDialogNodeOutputResponseTypeOption.Preference property.
// The preferred type of control to display, if supported by the channel.
const (
	DialogNodeOutputGenericDialogNodeOutputResponseTypeOptionPreferenceButtonConst   = "button"
	DialogNodeOutputGenericDialogNodeOutputResponseTypeOptionPreferenceDropdownConst = "dropdown"
)

// NewDialogNodeOutputGenericDialogNodeOutputResponseTypeOption : Instantiate DialogNodeOutputGenericDialogNodeOutputResponseTypeOption (Generic Model Constructor)
func (*AssistantV1) NewDialogNodeOutputGenericDialogNodeOutputResponseTypeOption(responseType string, title string, options []DialogNodeOutputOptionsElement) (_model *DialogNodeOutputGenericDialogNodeOutputResponseTypeOption, err error) {
	_model = &DialogNodeOutputGenericDialogNodeOutputResponseTypeOption{
		ResponseType: core.StringPtr(responseType),
		Title:        core.StringPtr(title),
		Options:      options,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*DialogNodeOutputGenericDialogNodeOutputResponseTypeOption) isaDialogNodeOutputGeneric() bool {
	return true
}

// UnmarshalDialogNodeOutputGenericDialogNodeOutputResponseTypeOption unmarshals an instance of DialogNodeOutputGenericDialogNodeOutputResponseTypeOption from the specified map of raw messages.
func UnmarshalDialogNodeOutputGenericDialogNodeOutputResponseTypeOption(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DialogNodeOutputGenericDialogNodeOutputResponseTypeOption)
	err = core.UnmarshalPrimitive(m, "response_type", &obj.ResponseType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "title", &obj.Title)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "preference", &obj.Preference)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "options", &obj.Options, UnmarshalDialogNodeOutputOptionsElement)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "channels", &obj.Channels, UnmarshalResponseGenericChannel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DialogNodeOutputGenericDialogNodeOutputResponseTypePause : DialogNodeOutputGenericDialogNodeOutputResponseTypePause struct
// This model "extends" DialogNodeOutputGeneric
type DialogNodeOutputGenericDialogNodeOutputResponseTypePause struct {
	// The type of response returned by the dialog node. The specified response type must be supported by the client
	// application or channel.
	ResponseType *string `json:"response_type" validate:"required"`

	// How long to pause, in milliseconds. The valid values are from 0 to 10000.
	Time *int64 `json:"time" validate:"required"`

	// Whether to send a "user is typing" event during the pause. Ignored if the channel does not support this event.
	Typing *bool `json:"typing,omitempty"`

	// An array of objects specifying channels for which the response is intended.
	Channels []ResponseGenericChannel `json:"channels,omitempty"`
}

// NewDialogNodeOutputGenericDialogNodeOutputResponseTypePause : Instantiate DialogNodeOutputGenericDialogNodeOutputResponseTypePause (Generic Model Constructor)
func (*AssistantV1) NewDialogNodeOutputGenericDialogNodeOutputResponseTypePause(responseType string, time int64) (_model *DialogNodeOutputGenericDialogNodeOutputResponseTypePause, err error) {
	_model = &DialogNodeOutputGenericDialogNodeOutputResponseTypePause{
		ResponseType: core.StringPtr(responseType),
		Time:         core.Int64Ptr(time),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*DialogNodeOutputGenericDialogNodeOutputResponseTypePause) isaDialogNodeOutputGeneric() bool {
	return true
}

// UnmarshalDialogNodeOutputGenericDialogNodeOutputResponseTypePause unmarshals an instance of DialogNodeOutputGenericDialogNodeOutputResponseTypePause from the specified map of raw messages.
func UnmarshalDialogNodeOutputGenericDialogNodeOutputResponseTypePause(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DialogNodeOutputGenericDialogNodeOutputResponseTypePause)
	err = core.UnmarshalPrimitive(m, "response_type", &obj.ResponseType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "time", &obj.Time)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "typing", &obj.Typing)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "channels", &obj.Channels, UnmarshalResponseGenericChannel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DialogNodeOutputGenericDialogNodeOutputResponseTypeSearchSkill : DialogNodeOutputGenericDialogNodeOutputResponseTypeSearchSkill struct
// This model "extends" DialogNodeOutputGeneric
type DialogNodeOutputGenericDialogNodeOutputResponseTypeSearchSkill struct {
	// The type of response returned by the dialog node. The specified response type must be supported by the client
	// application or channel.
	//
	// **Note:** The **search_skill** response type is used only by the v2 runtime API.
	ResponseType *string `json:"response_type" validate:"required"`

	// The text of the search query. This can be either a natural-language query or a query that uses the Discovery query
	// language syntax, depending on the value of the **query_type** property. For more information, see the [Discovery
	// service documentation](https://cloud.ibm.com/docs/discovery?topic=discovery-query-operators#query-operators).
	Query *string `json:"query" validate:"required"`

	// The type of the search query.
	QueryType *string `json:"query_type" validate:"required"`

	// An optional filter that narrows the set of documents to be searched. For more information, see the [Discovery
	// service documentation]([Discovery service
	// documentation](https://cloud.ibm.com/docs/discovery?topic=discovery-query-parameters#filter).
	Filter *string `json:"filter,omitempty"`

	// The version of the Discovery service API to use for the query.
	DiscoveryVersion *string `json:"discovery_version,omitempty"`

	// An array of objects specifying channels for which the response is intended.
	Channels []ResponseGenericChannel `json:"channels,omitempty"`
}

// Constants associated with the DialogNodeOutputGenericDialogNodeOutputResponseTypeSearchSkill.QueryType property.
// The type of the search query.
const (
	DialogNodeOutputGenericDialogNodeOutputResponseTypeSearchSkillQueryTypeDiscoveryQueryLanguageConst = "discovery_query_language"
	DialogNodeOutputGenericDialogNodeOutputResponseTypeSearchSkillQueryTypeNaturalLanguageConst        = "natural_language"
)

// NewDialogNodeOutputGenericDialogNodeOutputResponseTypeSearchSkill : Instantiate DialogNodeOutputGenericDialogNodeOutputResponseTypeSearchSkill (Generic Model Constructor)
func (*AssistantV1) NewDialogNodeOutputGenericDialogNodeOutputResponseTypeSearchSkill(responseType string, query string, queryType string) (_model *DialogNodeOutputGenericDialogNodeOutputResponseTypeSearchSkill, err error) {
	_model = &DialogNodeOutputGenericDialogNodeOutputResponseTypeSearchSkill{
		ResponseType: core.StringPtr(responseType),
		Query:        core.StringPtr(query),
		QueryType:    core.StringPtr(queryType),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*DialogNodeOutputGenericDialogNodeOutputResponseTypeSearchSkill) isaDialogNodeOutputGeneric() bool {
	return true
}

// UnmarshalDialogNodeOutputGenericDialogNodeOutputResponseTypeSearchSkill unmarshals an instance of DialogNodeOutputGenericDialogNodeOutputResponseTypeSearchSkill from the specified map of raw messages.
func UnmarshalDialogNodeOutputGenericDialogNodeOutputResponseTypeSearchSkill(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DialogNodeOutputGenericDialogNodeOutputResponseTypeSearchSkill)
	err = core.UnmarshalPrimitive(m, "response_type", &obj.ResponseType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "query", &obj.Query)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "query_type", &obj.QueryType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "filter", &obj.Filter)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "discovery_version", &obj.DiscoveryVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "channels", &obj.Channels, UnmarshalResponseGenericChannel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DialogNodeOutputGenericDialogNodeOutputResponseTypeText : DialogNodeOutputGenericDialogNodeOutputResponseTypeText struct
// This model "extends" DialogNodeOutputGeneric
type DialogNodeOutputGenericDialogNodeOutputResponseTypeText struct {
	// The type of response returned by the dialog node. The specified response type must be supported by the client
	// application or channel.
	ResponseType *string `json:"response_type" validate:"required"`

	// A list of one or more objects defining text responses.
	Values []DialogNodeOutputTextValuesElement `json:"values" validate:"required"`

	// How a response is selected from the list, if more than one response is specified.
	SelectionPolicy *string `json:"selection_policy,omitempty"`

	// The delimiter to use as a separator between responses when `selection_policy`=`multiline`.
	Delimiter *string `json:"delimiter,omitempty"`

	// An array of objects specifying channels for which the response is intended.
	Channels []ResponseGenericChannel `json:"channels,omitempty"`
}

// Constants associated with the DialogNodeOutputGenericDialogNodeOutputResponseTypeText.SelectionPolicy property.
// How a response is selected from the list, if more than one response is specified.
const (
	DialogNodeOutputGenericDialogNodeOutputResponseTypeTextSelectionPolicyMultilineConst  = "multiline"
	DialogNodeOutputGenericDialogNodeOutputResponseTypeTextSelectionPolicyRandomConst     = "random"
	DialogNodeOutputGenericDialogNodeOutputResponseTypeTextSelectionPolicySequentialConst = "sequential"
)

// NewDialogNodeOutputGenericDialogNodeOutputResponseTypeText : Instantiate DialogNodeOutputGenericDialogNodeOutputResponseTypeText (Generic Model Constructor)
func (*AssistantV1) NewDialogNodeOutputGenericDialogNodeOutputResponseTypeText(responseType string, values []DialogNodeOutputTextValuesElement) (_model *DialogNodeOutputGenericDialogNodeOutputResponseTypeText, err error) {
	_model = &DialogNodeOutputGenericDialogNodeOutputResponseTypeText{
		ResponseType: core.StringPtr(responseType),
		Values:       values,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*DialogNodeOutputGenericDialogNodeOutputResponseTypeText) isaDialogNodeOutputGeneric() bool {
	return true
}

// UnmarshalDialogNodeOutputGenericDialogNodeOutputResponseTypeText unmarshals an instance of DialogNodeOutputGenericDialogNodeOutputResponseTypeText from the specified map of raw messages.
func UnmarshalDialogNodeOutputGenericDialogNodeOutputResponseTypeText(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DialogNodeOutputGenericDialogNodeOutputResponseTypeText)
	err = core.UnmarshalPrimitive(m, "response_type", &obj.ResponseType)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "values", &obj.Values, UnmarshalDialogNodeOutputTextValuesElement)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "selection_policy", &obj.SelectionPolicy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "delimiter", &obj.Delimiter)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "channels", &obj.Channels, UnmarshalResponseGenericChannel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DialogNodeOutputGenericDialogNodeOutputResponseTypeUserDefined : DialogNodeOutputGenericDialogNodeOutputResponseTypeUserDefined struct
// This model "extends" DialogNodeOutputGeneric
type DialogNodeOutputGenericDialogNodeOutputResponseTypeUserDefined struct {
	// The type of response returned by the dialog node. The specified response type must be supported by the client
	// application or channel.
	ResponseType *string `json:"response_type" validate:"required"`

	// An object containing any properties for the user-defined response type. The total size of this object cannot exceed
	// 5000 bytes.
	UserDefined map[string]interface{} `json:"user_defined" validate:"required"`

	// An array of objects specifying channels for which the response is intended.
	Channels []ResponseGenericChannel `json:"channels,omitempty"`
}

// NewDialogNodeOutputGenericDialogNodeOutputResponseTypeUserDefined : Instantiate DialogNodeOutputGenericDialogNodeOutputResponseTypeUserDefined (Generic Model Constructor)
func (*AssistantV1) NewDialogNodeOutputGenericDialogNodeOutputResponseTypeUserDefined(responseType string, userDefined map[string]interface{}) (_model *DialogNodeOutputGenericDialogNodeOutputResponseTypeUserDefined, err error) {
	_model = &DialogNodeOutputGenericDialogNodeOutputResponseTypeUserDefined{
		ResponseType: core.StringPtr(responseType),
		UserDefined:  userDefined,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*DialogNodeOutputGenericDialogNodeOutputResponseTypeUserDefined) isaDialogNodeOutputGeneric() bool {
	return true
}

// UnmarshalDialogNodeOutputGenericDialogNodeOutputResponseTypeUserDefined unmarshals an instance of DialogNodeOutputGenericDialogNodeOutputResponseTypeUserDefined from the specified map of raw messages.
func UnmarshalDialogNodeOutputGenericDialogNodeOutputResponseTypeUserDefined(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DialogNodeOutputGenericDialogNodeOutputResponseTypeUserDefined)
	err = core.UnmarshalPrimitive(m, "response_type", &obj.ResponseType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "user_defined", &obj.UserDefined)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "channels", &obj.Channels, UnmarshalResponseGenericChannel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuntimeResponseGenericRuntimeResponseTypeChannelTransfer : RuntimeResponseGenericRuntimeResponseTypeChannelTransfer struct
// This model "extends" RuntimeResponseGeneric
type RuntimeResponseGenericRuntimeResponseTypeChannelTransfer struct {
	// The type of response returned by the dialog node. The specified response type must be supported by the client
	// application or channel.
	ResponseType *string `json:"response_type" validate:"required"`

	// The message to display to the user when initiating a channel transfer.
	MessageToUser *string `json:"message_to_user" validate:"required"`

	// Information used by an integration to transfer the conversation to a different channel.
	TransferInfo *ChannelTransferInfo `json:"transfer_info" validate:"required"`

	// An array of objects specifying channels for which the response is intended. If **channels** is present, the response
	// is intended only for a built-in integration and should not be handled by an API client.
	Channels []ResponseGenericChannel `json:"channels,omitempty"`
}

// NewRuntimeResponseGenericRuntimeResponseTypeChannelTransfer : Instantiate RuntimeResponseGenericRuntimeResponseTypeChannelTransfer (Generic Model Constructor)
func (*AssistantV1) NewRuntimeResponseGenericRuntimeResponseTypeChannelTransfer(responseType string, messageToUser string, transferInfo *ChannelTransferInfo) (_model *RuntimeResponseGenericRuntimeResponseTypeChannelTransfer, err error) {
	_model = &RuntimeResponseGenericRuntimeResponseTypeChannelTransfer{
		ResponseType:  core.StringPtr(responseType),
		MessageToUser: core.StringPtr(messageToUser),
		TransferInfo:  transferInfo,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*RuntimeResponseGenericRuntimeResponseTypeChannelTransfer) isaRuntimeResponseGeneric() bool {
	return true
}

// UnmarshalRuntimeResponseGenericRuntimeResponseTypeChannelTransfer unmarshals an instance of RuntimeResponseGenericRuntimeResponseTypeChannelTransfer from the specified map of raw messages.
func UnmarshalRuntimeResponseGenericRuntimeResponseTypeChannelTransfer(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuntimeResponseGenericRuntimeResponseTypeChannelTransfer)
	err = core.UnmarshalPrimitive(m, "response_type", &obj.ResponseType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "message_to_user", &obj.MessageToUser)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "transfer_info", &obj.TransferInfo, UnmarshalChannelTransferInfo)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "channels", &obj.Channels, UnmarshalResponseGenericChannel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuntimeResponseGenericRuntimeResponseTypeConnectToAgent : RuntimeResponseGenericRuntimeResponseTypeConnectToAgent struct
// This model "extends" RuntimeResponseGeneric
type RuntimeResponseGenericRuntimeResponseTypeConnectToAgent struct {
	// The type of response returned by the dialog node. The specified response type must be supported by the client
	// application or channel.
	ResponseType *string `json:"response_type" validate:"required"`

	// A message to be sent to the human agent who will be taking over the conversation.
	MessageToHumanAgent *string `json:"message_to_human_agent,omitempty"`

	// An optional message to be displayed to the user to indicate that the conversation will be transferred to the next
	// available agent.
	AgentAvailable *AgentAvailabilityMessage `json:"agent_available,omitempty"`

	// An optional message to be displayed to the user to indicate that no online agent is available to take over the
	// conversation.
	AgentUnavailable *AgentAvailabilityMessage `json:"agent_unavailable,omitempty"`

	// Routing or other contextual information to be used by target service desk systems.
	TransferInfo *DialogNodeOutputConnectToAgentTransferInfo `json:"transfer_info,omitempty"`

	// A label identifying the topic of the conversation, derived from the **title** property of the relevant node or the
	// **topic** property of the dialog node response.
	Topic *string `json:"topic,omitempty"`

	// The unique ID of the dialog node that the **topic** property is taken from. The **topic** property is populated
	// using the value of the dialog node's **title** property.
	DialogNode *string `json:"dialog_node,omitempty"`

	// An array of objects specifying channels for which the response is intended. If **channels** is present, the response
	// is intended for a built-in integration and should not be handled by an API client.
	Channels []ResponseGenericChannel `json:"channels,omitempty"`
}

// NewRuntimeResponseGenericRuntimeResponseTypeConnectToAgent : Instantiate RuntimeResponseGenericRuntimeResponseTypeConnectToAgent (Generic Model Constructor)
func (*AssistantV1) NewRuntimeResponseGenericRuntimeResponseTypeConnectToAgent(responseType string) (_model *RuntimeResponseGenericRuntimeResponseTypeConnectToAgent, err error) {
	_model = &RuntimeResponseGenericRuntimeResponseTypeConnectToAgent{
		ResponseType: core.StringPtr(responseType),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*RuntimeResponseGenericRuntimeResponseTypeConnectToAgent) isaRuntimeResponseGeneric() bool {
	return true
}

// UnmarshalRuntimeResponseGenericRuntimeResponseTypeConnectToAgent unmarshals an instance of RuntimeResponseGenericRuntimeResponseTypeConnectToAgent from the specified map of raw messages.
func UnmarshalRuntimeResponseGenericRuntimeResponseTypeConnectToAgent(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuntimeResponseGenericRuntimeResponseTypeConnectToAgent)
	err = core.UnmarshalPrimitive(m, "response_type", &obj.ResponseType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "message_to_human_agent", &obj.MessageToHumanAgent)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "agent_available", &obj.AgentAvailable, UnmarshalAgentAvailabilityMessage)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "agent_unavailable", &obj.AgentUnavailable, UnmarshalAgentAvailabilityMessage)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "transfer_info", &obj.TransferInfo, UnmarshalDialogNodeOutputConnectToAgentTransferInfo)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "topic", &obj.Topic)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "dialog_node", &obj.DialogNode)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "channels", &obj.Channels, UnmarshalResponseGenericChannel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuntimeResponseGenericRuntimeResponseTypeImage : RuntimeResponseGenericRuntimeResponseTypeImage struct
// This model "extends" RuntimeResponseGeneric
type RuntimeResponseGenericRuntimeResponseTypeImage struct {
	// The type of response returned by the dialog node. The specified response type must be supported by the client
	// application or channel.
	ResponseType *string `json:"response_type" validate:"required"`

	// The `https:` URL of the image.
	Source *string `json:"source" validate:"required"`

	// The title or introductory text to show before the response.
	Title *string `json:"title,omitempty"`

	// The description to show with the the response.
	Description *string `json:"description,omitempty"`

	// An array of objects specifying channels for which the response is intended. If **channels** is present, the response
	// is intended for a built-in integration and should not be handled by an API client.
	Channels []ResponseGenericChannel `json:"channels,omitempty"`

	// Descriptive text that can be used for screen readers or other situations where the image cannot be seen.
	AltText *string `json:"alt_text,omitempty"`
}

// NewRuntimeResponseGenericRuntimeResponseTypeImage : Instantiate RuntimeResponseGenericRuntimeResponseTypeImage (Generic Model Constructor)
func (*AssistantV1) NewRuntimeResponseGenericRuntimeResponseTypeImage(responseType string, source string) (_model *RuntimeResponseGenericRuntimeResponseTypeImage, err error) {
	_model = &RuntimeResponseGenericRuntimeResponseTypeImage{
		ResponseType: core.StringPtr(responseType),
		Source:       core.StringPtr(source),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*RuntimeResponseGenericRuntimeResponseTypeImage) isaRuntimeResponseGeneric() bool {
	return true
}

// UnmarshalRuntimeResponseGenericRuntimeResponseTypeImage unmarshals an instance of RuntimeResponseGenericRuntimeResponseTypeImage from the specified map of raw messages.
func UnmarshalRuntimeResponseGenericRuntimeResponseTypeImage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuntimeResponseGenericRuntimeResponseTypeImage)
	err = core.UnmarshalPrimitive(m, "response_type", &obj.ResponseType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source", &obj.Source)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "title", &obj.Title)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "channels", &obj.Channels, UnmarshalResponseGenericChannel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "alt_text", &obj.AltText)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuntimeResponseGenericRuntimeResponseTypeOption : RuntimeResponseGenericRuntimeResponseTypeOption struct
// This model "extends" RuntimeResponseGeneric
type RuntimeResponseGenericRuntimeResponseTypeOption struct {
	// The type of response returned by the dialog node. The specified response type must be supported by the client
	// application or channel.
	ResponseType *string `json:"response_type" validate:"required"`

	// The title or introductory text to show before the response.
	Title *string `json:"title" validate:"required"`

	// The description to show with the the response.
	Description *string `json:"description,omitempty"`

	// The preferred type of control to display.
	Preference *string `json:"preference,omitempty"`

	// An array of objects describing the options from which the user can choose.
	Options []DialogNodeOutputOptionsElement `json:"options" validate:"required"`

	// An array of objects specifying channels for which the response is intended. If **channels** is present, the response
	// is intended for a built-in integration and should not be handled by an API client.
	Channels []ResponseGenericChannel `json:"channels,omitempty"`
}

// Constants associated with the RuntimeResponseGenericRuntimeResponseTypeOption.Preference property.
// The preferred type of control to display.
const (
	RuntimeResponseGenericRuntimeResponseTypeOptionPreferenceButtonConst   = "button"
	RuntimeResponseGenericRuntimeResponseTypeOptionPreferenceDropdownConst = "dropdown"
)

// NewRuntimeResponseGenericRuntimeResponseTypeOption : Instantiate RuntimeResponseGenericRuntimeResponseTypeOption (Generic Model Constructor)
func (*AssistantV1) NewRuntimeResponseGenericRuntimeResponseTypeOption(responseType string, title string, options []DialogNodeOutputOptionsElement) (_model *RuntimeResponseGenericRuntimeResponseTypeOption, err error) {
	_model = &RuntimeResponseGenericRuntimeResponseTypeOption{
		ResponseType: core.StringPtr(responseType),
		Title:        core.StringPtr(title),
		Options:      options,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*RuntimeResponseGenericRuntimeResponseTypeOption) isaRuntimeResponseGeneric() bool {
	return true
}

// UnmarshalRuntimeResponseGenericRuntimeResponseTypeOption unmarshals an instance of RuntimeResponseGenericRuntimeResponseTypeOption from the specified map of raw messages.
func UnmarshalRuntimeResponseGenericRuntimeResponseTypeOption(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuntimeResponseGenericRuntimeResponseTypeOption)
	err = core.UnmarshalPrimitive(m, "response_type", &obj.ResponseType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "title", &obj.Title)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "preference", &obj.Preference)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "options", &obj.Options, UnmarshalDialogNodeOutputOptionsElement)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "channels", &obj.Channels, UnmarshalResponseGenericChannel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuntimeResponseGenericRuntimeResponseTypePause : RuntimeResponseGenericRuntimeResponseTypePause struct
// This model "extends" RuntimeResponseGeneric
type RuntimeResponseGenericRuntimeResponseTypePause struct {
	// The type of response returned by the dialog node. The specified response type must be supported by the client
	// application or channel.
	ResponseType *string `json:"response_type" validate:"required"`

	// How long to pause, in milliseconds.
	Time *int64 `json:"time" validate:"required"`

	// Whether to send a "user is typing" event during the pause.
	Typing *bool `json:"typing,omitempty"`

	// An array of objects specifying channels for which the response is intended. If **channels** is present, the response
	// is intended for a built-in integration and should not be handled by an API client.
	Channels []ResponseGenericChannel `json:"channels,omitempty"`
}

// NewRuntimeResponseGenericRuntimeResponseTypePause : Instantiate RuntimeResponseGenericRuntimeResponseTypePause (Generic Model Constructor)
func (*AssistantV1) NewRuntimeResponseGenericRuntimeResponseTypePause(responseType string, time int64) (_model *RuntimeResponseGenericRuntimeResponseTypePause, err error) {
	_model = &RuntimeResponseGenericRuntimeResponseTypePause{
		ResponseType: core.StringPtr(responseType),
		Time:         core.Int64Ptr(time),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*RuntimeResponseGenericRuntimeResponseTypePause) isaRuntimeResponseGeneric() bool {
	return true
}

// UnmarshalRuntimeResponseGenericRuntimeResponseTypePause unmarshals an instance of RuntimeResponseGenericRuntimeResponseTypePause from the specified map of raw messages.
func UnmarshalRuntimeResponseGenericRuntimeResponseTypePause(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuntimeResponseGenericRuntimeResponseTypePause)
	err = core.UnmarshalPrimitive(m, "response_type", &obj.ResponseType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "time", &obj.Time)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "typing", &obj.Typing)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "channels", &obj.Channels, UnmarshalResponseGenericChannel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuntimeResponseGenericRuntimeResponseTypeSuggestion : RuntimeResponseGenericRuntimeResponseTypeSuggestion struct
// This model "extends" RuntimeResponseGeneric
type RuntimeResponseGenericRuntimeResponseTypeSuggestion struct {
	// The type of response returned by the dialog node. The specified response type must be supported by the client
	// application or channel.
	ResponseType *string `json:"response_type" validate:"required"`

	// The title or introductory text to show before the response.
	Title *string `json:"title" validate:"required"`

	// An array of objects describing the possible matching dialog nodes from which the user can choose.
	Suggestions []DialogSuggestion `json:"suggestions" validate:"required"`

	// An array of objects specifying channels for which the response is intended. If **channels** is present, the response
	// is intended for a built-in integration and should not be handled by an API client.
	Channels []ResponseGenericChannel `json:"channels,omitempty"`
}

// NewRuntimeResponseGenericRuntimeResponseTypeSuggestion : Instantiate RuntimeResponseGenericRuntimeResponseTypeSuggestion (Generic Model Constructor)
func (*AssistantV1) NewRuntimeResponseGenericRuntimeResponseTypeSuggestion(responseType string, title string, suggestions []DialogSuggestion) (_model *RuntimeResponseGenericRuntimeResponseTypeSuggestion, err error) {
	_model = &RuntimeResponseGenericRuntimeResponseTypeSuggestion{
		ResponseType: core.StringPtr(responseType),
		Title:        core.StringPtr(title),
		Suggestions:  suggestions,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*RuntimeResponseGenericRuntimeResponseTypeSuggestion) isaRuntimeResponseGeneric() bool {
	return true
}

// UnmarshalRuntimeResponseGenericRuntimeResponseTypeSuggestion unmarshals an instance of RuntimeResponseGenericRuntimeResponseTypeSuggestion from the specified map of raw messages.
func UnmarshalRuntimeResponseGenericRuntimeResponseTypeSuggestion(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuntimeResponseGenericRuntimeResponseTypeSuggestion)
	err = core.UnmarshalPrimitive(m, "response_type", &obj.ResponseType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "title", &obj.Title)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "suggestions", &obj.Suggestions, UnmarshalDialogSuggestion)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "channels", &obj.Channels, UnmarshalResponseGenericChannel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuntimeResponseGenericRuntimeResponseTypeText : RuntimeResponseGenericRuntimeResponseTypeText struct
// This model "extends" RuntimeResponseGeneric
type RuntimeResponseGenericRuntimeResponseTypeText struct {
	// The type of response returned by the dialog node. The specified response type must be supported by the client
	// application or channel.
	ResponseType *string `json:"response_type" validate:"required"`

	// The text of the response.
	Text *string `json:"text" validate:"required"`

	// An array of objects specifying channels for which the response is intended. If **channels** is present, the response
	// is intended for a built-in integration and should not be handled by an API client.
	Channels []ResponseGenericChannel `json:"channels,omitempty"`
}

// NewRuntimeResponseGenericRuntimeResponseTypeText : Instantiate RuntimeResponseGenericRuntimeResponseTypeText (Generic Model Constructor)
func (*AssistantV1) NewRuntimeResponseGenericRuntimeResponseTypeText(responseType string, text string) (_model *RuntimeResponseGenericRuntimeResponseTypeText, err error) {
	_model = &RuntimeResponseGenericRuntimeResponseTypeText{
		ResponseType: core.StringPtr(responseType),
		Text:         core.StringPtr(text),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*RuntimeResponseGenericRuntimeResponseTypeText) isaRuntimeResponseGeneric() bool {
	return true
}

// UnmarshalRuntimeResponseGenericRuntimeResponseTypeText unmarshals an instance of RuntimeResponseGenericRuntimeResponseTypeText from the specified map of raw messages.
func UnmarshalRuntimeResponseGenericRuntimeResponseTypeText(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuntimeResponseGenericRuntimeResponseTypeText)
	err = core.UnmarshalPrimitive(m, "response_type", &obj.ResponseType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "channels", &obj.Channels, UnmarshalResponseGenericChannel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuntimeResponseGenericRuntimeResponseTypeUserDefined : RuntimeResponseGenericRuntimeResponseTypeUserDefined struct
// This model "extends" RuntimeResponseGeneric
type RuntimeResponseGenericRuntimeResponseTypeUserDefined struct {
	// The type of response returned by the dialog node. The specified response type must be supported by the client
	// application or channel.
	ResponseType *string `json:"response_type" validate:"required"`

	// An object containing any properties for the user-defined response type.
	UserDefined map[string]interface{} `json:"user_defined" validate:"required"`

	// An array of objects specifying channels for which the response is intended. If **channels** is present, the response
	// is intended for a built-in integration and should not be handled by an API client.
	Channels []ResponseGenericChannel `json:"channels,omitempty"`
}

// NewRuntimeResponseGenericRuntimeResponseTypeUserDefined : Instantiate RuntimeResponseGenericRuntimeResponseTypeUserDefined (Generic Model Constructor)
func (*AssistantV1) NewRuntimeResponseGenericRuntimeResponseTypeUserDefined(responseType string, userDefined map[string]interface{}) (_model *RuntimeResponseGenericRuntimeResponseTypeUserDefined, err error) {
	_model = &RuntimeResponseGenericRuntimeResponseTypeUserDefined{
		ResponseType: core.StringPtr(responseType),
		UserDefined:  userDefined,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*RuntimeResponseGenericRuntimeResponseTypeUserDefined) isaRuntimeResponseGeneric() bool {
	return true
}

// UnmarshalRuntimeResponseGenericRuntimeResponseTypeUserDefined unmarshals an instance of RuntimeResponseGenericRuntimeResponseTypeUserDefined from the specified map of raw messages.
func UnmarshalRuntimeResponseGenericRuntimeResponseTypeUserDefined(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuntimeResponseGenericRuntimeResponseTypeUserDefined)
	err = core.UnmarshalPrimitive(m, "response_type", &obj.ResponseType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "user_defined", &obj.UserDefined)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "channels", &obj.Channels, UnmarshalResponseGenericChannel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
