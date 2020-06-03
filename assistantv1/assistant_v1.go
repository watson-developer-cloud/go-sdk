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

// Package assistantv1 : Operations and models for the AssistantV1 service
package assistantv1

import (
	"fmt"
	"github.com/IBM/go-sdk-core/core"
	"github.com/go-openapi/strfmt"
	common "github.com/watson-developer-cloud/go-sdk/common"
)

// AssistantV1 : The IBM Watson&trade; Assistant service combines machine learning, natural language understanding, and
// an integrated dialog editor to create conversation flows between your apps and your users.
//
// The Assistant v1 API provides authoring methods your application can use to create or update a workspace.
//
// Version: 1.0
// See: https://cloud.ibm.com/docs/assistant/
type AssistantV1 struct {
	Service *core.BaseService
	Version string
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://gateway.watsonplatform.net/assistant/api"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "conversation"

// AssistantV1Options : Service options
type AssistantV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
	Version       string
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

// SetServiceURL sets the service URL
func (assistant *AssistantV1) SetServiceURL(url string) error {
	return assistant.Service.SetServiceURL(url)
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
//
// There is no rate limit for this operation.
func (assistant *AssistantV1) Message(messageOptions *MessageOptions) (result *MessageResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(messageOptions, "messageOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(messageOptions, "messageOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces", "message"}
	pathParameters := []string{*messageOptions.WorkspaceID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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

	if messageOptions.NodesVisitedDetails != nil {
		builder.AddQuery("nodes_visited_details", fmt.Sprint(*messageOptions.NodesVisitedDetails))
	}
	builder.AddQuery("version", assistant.Version)

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
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, new(MessageResponse))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*MessageResponse)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// ListWorkspaces : List workspaces
// List the workspaces associated with a Watson Assistant service instance.
//
// This operation is limited to 500 requests per 30 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) ListWorkspaces(listWorkspacesOptions *ListWorkspacesOptions) (result *WorkspaceCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listWorkspacesOptions, "listWorkspacesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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

	if listWorkspacesOptions.PageLimit != nil {
		builder.AddQuery("page_limit", fmt.Sprint(*listWorkspacesOptions.PageLimit))
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
	builder.AddQuery("version", assistant.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, new(WorkspaceCollection))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*WorkspaceCollection)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// CreateWorkspace : Create workspace
// Create a workspace based on component objects. You must provide workspace components defining the content of the new
// workspace.
//
// This operation is limited to 30 requests per 30 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) CreateWorkspace(createWorkspaceOptions *CreateWorkspaceOptions) (result *Workspace, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(createWorkspaceOptions, "createWorkspaceOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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

	if createWorkspaceOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*createWorkspaceOptions.IncludeAudit))
	}
	builder.AddQuery("version", assistant.Version)

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
	if createWorkspaceOptions.Metadata != nil {
		body["metadata"] = createWorkspaceOptions.Metadata
	}
	if createWorkspaceOptions.LearningOptOut != nil {
		body["learning_opt_out"] = createWorkspaceOptions.LearningOptOut
	}
	if createWorkspaceOptions.SystemSettings != nil {
		body["system_settings"] = createWorkspaceOptions.SystemSettings
	}
	if createWorkspaceOptions.Intents != nil {
		body["intents"] = createWorkspaceOptions.Intents
	}
	if createWorkspaceOptions.Entities != nil {
		body["entities"] = createWorkspaceOptions.Entities
	}
	if createWorkspaceOptions.DialogNodes != nil {
		body["dialog_nodes"] = createWorkspaceOptions.DialogNodes
	}
	if createWorkspaceOptions.Counterexamples != nil {
		body["counterexamples"] = createWorkspaceOptions.Counterexamples
	}
	if createWorkspaceOptions.Webhooks != nil {
		body["webhooks"] = createWorkspaceOptions.Webhooks
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, new(Workspace))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*Workspace)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// GetWorkspace : Get information about a workspace
// Get information about a workspace, optionally including all workspace content.
//
// With **export**=`false`, this operation is limited to 6000 requests per 5 minutes. With **export**=`true`, the limit
// is 20 requests per 30 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) GetWorkspace(getWorkspaceOptions *GetWorkspaceOptions) (result *Workspace, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getWorkspaceOptions, "getWorkspaceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getWorkspaceOptions, "getWorkspaceOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces"}
	pathParameters := []string{*getWorkspaceOptions.WorkspaceID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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

	if getWorkspaceOptions.Export != nil {
		builder.AddQuery("export", fmt.Sprint(*getWorkspaceOptions.Export))
	}
	if getWorkspaceOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*getWorkspaceOptions.IncludeAudit))
	}
	if getWorkspaceOptions.Sort != nil {
		builder.AddQuery("sort", fmt.Sprint(*getWorkspaceOptions.Sort))
	}
	builder.AddQuery("version", assistant.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, new(Workspace))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*Workspace)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// UpdateWorkspace : Update workspace
// Update an existing workspace with new or modified data. You must provide component objects defining the content of
// the updated workspace.
//
// This operation is limited to 30 request per 30 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) UpdateWorkspace(updateWorkspaceOptions *UpdateWorkspaceOptions) (result *Workspace, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateWorkspaceOptions, "updateWorkspaceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateWorkspaceOptions, "updateWorkspaceOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces"}
	pathParameters := []string{*updateWorkspaceOptions.WorkspaceID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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

	if updateWorkspaceOptions.Append != nil {
		builder.AddQuery("append", fmt.Sprint(*updateWorkspaceOptions.Append))
	}
	if updateWorkspaceOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*updateWorkspaceOptions.IncludeAudit))
	}
	builder.AddQuery("version", assistant.Version)

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
	if updateWorkspaceOptions.Metadata != nil {
		body["metadata"] = updateWorkspaceOptions.Metadata
	}
	if updateWorkspaceOptions.LearningOptOut != nil {
		body["learning_opt_out"] = updateWorkspaceOptions.LearningOptOut
	}
	if updateWorkspaceOptions.SystemSettings != nil {
		body["system_settings"] = updateWorkspaceOptions.SystemSettings
	}
	if updateWorkspaceOptions.Intents != nil {
		body["intents"] = updateWorkspaceOptions.Intents
	}
	if updateWorkspaceOptions.Entities != nil {
		body["entities"] = updateWorkspaceOptions.Entities
	}
	if updateWorkspaceOptions.DialogNodes != nil {
		body["dialog_nodes"] = updateWorkspaceOptions.DialogNodes
	}
	if updateWorkspaceOptions.Counterexamples != nil {
		body["counterexamples"] = updateWorkspaceOptions.Counterexamples
	}
	if updateWorkspaceOptions.Webhooks != nil {
		body["webhooks"] = updateWorkspaceOptions.Webhooks
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, new(Workspace))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*Workspace)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// DeleteWorkspace : Delete workspace
// Delete a workspace from the service instance.
//
// This operation is limited to 30 requests per 30 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) DeleteWorkspace(deleteWorkspaceOptions *DeleteWorkspaceOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteWorkspaceOptions, "deleteWorkspaceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteWorkspaceOptions, "deleteWorkspaceOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces"}
	pathParameters := []string{*deleteWorkspaceOptions.WorkspaceID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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
	builder.AddQuery("version", assistant.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, nil)

	return
}

// ListIntents : List intents
// List the intents for a workspace.
//
// With **export**=`false`, this operation is limited to 2000 requests per 30 minutes. With **export**=`true`, the limit
// is 400 requests per 30 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) ListIntents(listIntentsOptions *ListIntentsOptions) (result *IntentCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listIntentsOptions, "listIntentsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listIntentsOptions, "listIntentsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces", "intents"}
	pathParameters := []string{*listIntentsOptions.WorkspaceID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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

	if listIntentsOptions.Export != nil {
		builder.AddQuery("export", fmt.Sprint(*listIntentsOptions.Export))
	}
	if listIntentsOptions.PageLimit != nil {
		builder.AddQuery("page_limit", fmt.Sprint(*listIntentsOptions.PageLimit))
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
	builder.AddQuery("version", assistant.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, new(IntentCollection))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*IntentCollection)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// CreateIntent : Create intent
// Create a new intent.
//
// If you want to create multiple intents with a single API call, consider using the **[Update
// workspace](#update-workspace)** method instead.
//
// This operation is limited to 2000 requests per 30 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) CreateIntent(createIntentOptions *CreateIntentOptions) (result *Intent, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createIntentOptions, "createIntentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createIntentOptions, "createIntentOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces", "intents"}
	pathParameters := []string{*createIntentOptions.WorkspaceID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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

	if createIntentOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*createIntentOptions.IncludeAudit))
	}
	builder.AddQuery("version", assistant.Version)

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

	response, err = assistant.Service.Request(request, new(Intent))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*Intent)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// GetIntent : Get intent
// Get information about an intent, optionally including all intent content.
//
// With **export**=`false`, this operation is limited to 6000 requests per 5 minutes. With **export**=`true`, the limit
// is 400 requests per 30 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) GetIntent(getIntentOptions *GetIntentOptions) (result *Intent, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getIntentOptions, "getIntentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getIntentOptions, "getIntentOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces", "intents"}
	pathParameters := []string{*getIntentOptions.WorkspaceID, *getIntentOptions.Intent}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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

	if getIntentOptions.Export != nil {
		builder.AddQuery("export", fmt.Sprint(*getIntentOptions.Export))
	}
	if getIntentOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*getIntentOptions.IncludeAudit))
	}
	builder.AddQuery("version", assistant.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, new(Intent))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*Intent)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// UpdateIntent : Update intent
// Update an existing intent with new or modified data. You must provide component objects defining the content of the
// updated intent.
//
// If you want to update multiple intents with a single API call, consider using the **[Update
// workspace](#update-workspace)** method instead.
//
// This operation is limited to 2000 requests per 30 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) UpdateIntent(updateIntentOptions *UpdateIntentOptions) (result *Intent, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateIntentOptions, "updateIntentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateIntentOptions, "updateIntentOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces", "intents"}
	pathParameters := []string{*updateIntentOptions.WorkspaceID, *updateIntentOptions.Intent}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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

	if updateIntentOptions.Append != nil {
		builder.AddQuery("append", fmt.Sprint(*updateIntentOptions.Append))
	}
	if updateIntentOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*updateIntentOptions.IncludeAudit))
	}
	builder.AddQuery("version", assistant.Version)

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

	response, err = assistant.Service.Request(request, new(Intent))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*Intent)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// DeleteIntent : Delete intent
// Delete an intent from a workspace.
//
// This operation is limited to 2000 requests per 30 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) DeleteIntent(deleteIntentOptions *DeleteIntentOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteIntentOptions, "deleteIntentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteIntentOptions, "deleteIntentOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces", "intents"}
	pathParameters := []string{*deleteIntentOptions.WorkspaceID, *deleteIntentOptions.Intent}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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
	builder.AddQuery("version", assistant.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, nil)

	return
}

// ListExamples : List user input examples
// List the user input examples for an intent, optionally including contextual entity mentions.
//
// This operation is limited to 2500 requests per 30 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) ListExamples(listExamplesOptions *ListExamplesOptions) (result *ExampleCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listExamplesOptions, "listExamplesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listExamplesOptions, "listExamplesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces", "intents", "examples"}
	pathParameters := []string{*listExamplesOptions.WorkspaceID, *listExamplesOptions.Intent}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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

	if listExamplesOptions.PageLimit != nil {
		builder.AddQuery("page_limit", fmt.Sprint(*listExamplesOptions.PageLimit))
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
	builder.AddQuery("version", assistant.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, new(ExampleCollection))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*ExampleCollection)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// CreateExample : Create user input example
// Add a new user input example to an intent.
//
// If you want to add multiple examples with a single API call, consider using the **[Update intent](#update-intent)**
// method instead.
//
// This operation is limited to 1000 requests per 30 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) CreateExample(createExampleOptions *CreateExampleOptions) (result *Example, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createExampleOptions, "createExampleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createExampleOptions, "createExampleOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces", "intents", "examples"}
	pathParameters := []string{*createExampleOptions.WorkspaceID, *createExampleOptions.Intent}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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

	if createExampleOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*createExampleOptions.IncludeAudit))
	}
	builder.AddQuery("version", assistant.Version)

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

	response, err = assistant.Service.Request(request, new(Example))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*Example)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// GetExample : Get user input example
// Get information about a user input example.
//
// This operation is limited to 6000 requests per 5 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) GetExample(getExampleOptions *GetExampleOptions) (result *Example, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getExampleOptions, "getExampleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getExampleOptions, "getExampleOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces", "intents", "examples"}
	pathParameters := []string{*getExampleOptions.WorkspaceID, *getExampleOptions.Intent, *getExampleOptions.Text}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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

	if getExampleOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*getExampleOptions.IncludeAudit))
	}
	builder.AddQuery("version", assistant.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, new(Example))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*Example)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// UpdateExample : Update user input example
// Update the text of a user input example.
//
// If you want to update multiple examples with a single API call, consider using the **[Update
// intent](#update-intent)** method instead.
//
// This operation is limited to 1000 requests per 30 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) UpdateExample(updateExampleOptions *UpdateExampleOptions) (result *Example, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateExampleOptions, "updateExampleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateExampleOptions, "updateExampleOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces", "intents", "examples"}
	pathParameters := []string{*updateExampleOptions.WorkspaceID, *updateExampleOptions.Intent, *updateExampleOptions.Text}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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

	if updateExampleOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*updateExampleOptions.IncludeAudit))
	}
	builder.AddQuery("version", assistant.Version)

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

	response, err = assistant.Service.Request(request, new(Example))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*Example)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// DeleteExample : Delete user input example
// Delete a user input example from an intent.
//
// This operation is limited to 1000 requests per 30 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) DeleteExample(deleteExampleOptions *DeleteExampleOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteExampleOptions, "deleteExampleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteExampleOptions, "deleteExampleOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces", "intents", "examples"}
	pathParameters := []string{*deleteExampleOptions.WorkspaceID, *deleteExampleOptions.Intent, *deleteExampleOptions.Text}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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
	builder.AddQuery("version", assistant.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, nil)

	return
}

// ListCounterexamples : List counterexamples
// List the counterexamples for a workspace. Counterexamples are examples that have been marked as irrelevant input.
//
// This operation is limited to 2500 requests per 30 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) ListCounterexamples(listCounterexamplesOptions *ListCounterexamplesOptions) (result *CounterexampleCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listCounterexamplesOptions, "listCounterexamplesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listCounterexamplesOptions, "listCounterexamplesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces", "counterexamples"}
	pathParameters := []string{*listCounterexamplesOptions.WorkspaceID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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

	if listCounterexamplesOptions.PageLimit != nil {
		builder.AddQuery("page_limit", fmt.Sprint(*listCounterexamplesOptions.PageLimit))
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
	builder.AddQuery("version", assistant.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, new(CounterexampleCollection))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*CounterexampleCollection)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// CreateCounterexample : Create counterexample
// Add a new counterexample to a workspace. Counterexamples are examples that have been marked as irrelevant input.
//
// If you want to add multiple counterexamples with a single API call, consider using the **[Update
// workspace](#update-workspace)** method instead.
//
// This operation is limited to 1000 requests per 30 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) CreateCounterexample(createCounterexampleOptions *CreateCounterexampleOptions) (result *Counterexample, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createCounterexampleOptions, "createCounterexampleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createCounterexampleOptions, "createCounterexampleOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces", "counterexamples"}
	pathParameters := []string{*createCounterexampleOptions.WorkspaceID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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

	if createCounterexampleOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*createCounterexampleOptions.IncludeAudit))
	}
	builder.AddQuery("version", assistant.Version)

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

	response, err = assistant.Service.Request(request, new(Counterexample))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*Counterexample)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// GetCounterexample : Get counterexample
// Get information about a counterexample. Counterexamples are examples that have been marked as irrelevant input.
//
// This operation is limited to 6000 requests per 5 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) GetCounterexample(getCounterexampleOptions *GetCounterexampleOptions) (result *Counterexample, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getCounterexampleOptions, "getCounterexampleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getCounterexampleOptions, "getCounterexampleOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces", "counterexamples"}
	pathParameters := []string{*getCounterexampleOptions.WorkspaceID, *getCounterexampleOptions.Text}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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

	if getCounterexampleOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*getCounterexampleOptions.IncludeAudit))
	}
	builder.AddQuery("version", assistant.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, new(Counterexample))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*Counterexample)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// UpdateCounterexample : Update counterexample
// Update the text of a counterexample. Counterexamples are examples that have been marked as irrelevant input.
//
// If you want to update multiple counterexamples with a single API call, consider using the **[Update
// workspace](#update-workspace)** method instead.
//
// This operation is limited to 1000 requests per 30 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) UpdateCounterexample(updateCounterexampleOptions *UpdateCounterexampleOptions) (result *Counterexample, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateCounterexampleOptions, "updateCounterexampleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateCounterexampleOptions, "updateCounterexampleOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces", "counterexamples"}
	pathParameters := []string{*updateCounterexampleOptions.WorkspaceID, *updateCounterexampleOptions.Text}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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

	if updateCounterexampleOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*updateCounterexampleOptions.IncludeAudit))
	}
	builder.AddQuery("version", assistant.Version)

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

	response, err = assistant.Service.Request(request, new(Counterexample))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*Counterexample)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// DeleteCounterexample : Delete counterexample
// Delete a counterexample from a workspace. Counterexamples are examples that have been marked as irrelevant input.
//
// This operation is limited to 1000 requests per 30 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) DeleteCounterexample(deleteCounterexampleOptions *DeleteCounterexampleOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteCounterexampleOptions, "deleteCounterexampleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteCounterexampleOptions, "deleteCounterexampleOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces", "counterexamples"}
	pathParameters := []string{*deleteCounterexampleOptions.WorkspaceID, *deleteCounterexampleOptions.Text}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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
	builder.AddQuery("version", assistant.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, nil)

	return
}

// ListEntities : List entities
// List the entities for a workspace.
//
// With **export**=`false`, this operation is limited to 1000 requests per 30 minutes. With **export**=`true`, the limit
// is 200 requests per 30 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) ListEntities(listEntitiesOptions *ListEntitiesOptions) (result *EntityCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listEntitiesOptions, "listEntitiesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listEntitiesOptions, "listEntitiesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces", "entities"}
	pathParameters := []string{*listEntitiesOptions.WorkspaceID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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

	if listEntitiesOptions.Export != nil {
		builder.AddQuery("export", fmt.Sprint(*listEntitiesOptions.Export))
	}
	if listEntitiesOptions.PageLimit != nil {
		builder.AddQuery("page_limit", fmt.Sprint(*listEntitiesOptions.PageLimit))
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
	builder.AddQuery("version", assistant.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, new(EntityCollection))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*EntityCollection)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// CreateEntity : Create entity
// Create a new entity, or enable a system entity.
//
// If you want to create multiple entities with a single API call, consider using the **[Update
// workspace](#update-workspace)** method instead.
//
// This operation is limited to 1000 requests per 30 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) CreateEntity(createEntityOptions *CreateEntityOptions) (result *Entity, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createEntityOptions, "createEntityOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createEntityOptions, "createEntityOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces", "entities"}
	pathParameters := []string{*createEntityOptions.WorkspaceID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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

	if createEntityOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*createEntityOptions.IncludeAudit))
	}
	builder.AddQuery("version", assistant.Version)

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

	response, err = assistant.Service.Request(request, new(Entity))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*Entity)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// GetEntity : Get entity
// Get information about an entity, optionally including all entity content.
//
// With **export**=`false`, this operation is limited to 6000 requests per 5 minutes. With **export**=`true`, the limit
// is 200 requests per 30 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) GetEntity(getEntityOptions *GetEntityOptions) (result *Entity, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getEntityOptions, "getEntityOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getEntityOptions, "getEntityOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces", "entities"}
	pathParameters := []string{*getEntityOptions.WorkspaceID, *getEntityOptions.Entity}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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

	if getEntityOptions.Export != nil {
		builder.AddQuery("export", fmt.Sprint(*getEntityOptions.Export))
	}
	if getEntityOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*getEntityOptions.IncludeAudit))
	}
	builder.AddQuery("version", assistant.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, new(Entity))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*Entity)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// UpdateEntity : Update entity
// Update an existing entity with new or modified data. You must provide component objects defining the content of the
// updated entity.
//
// If you want to update multiple entities with a single API call, consider using the **[Update
// workspace](#update-workspace)** method instead.
//
// This operation is limited to 1000 requests per 30 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) UpdateEntity(updateEntityOptions *UpdateEntityOptions) (result *Entity, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateEntityOptions, "updateEntityOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateEntityOptions, "updateEntityOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces", "entities"}
	pathParameters := []string{*updateEntityOptions.WorkspaceID, *updateEntityOptions.Entity}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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

	if updateEntityOptions.Append != nil {
		builder.AddQuery("append", fmt.Sprint(*updateEntityOptions.Append))
	}
	if updateEntityOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*updateEntityOptions.IncludeAudit))
	}
	builder.AddQuery("version", assistant.Version)

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

	response, err = assistant.Service.Request(request, new(Entity))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*Entity)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// DeleteEntity : Delete entity
// Delete an entity from a workspace, or disable a system entity.
//
// This operation is limited to 1000 requests per 30 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) DeleteEntity(deleteEntityOptions *DeleteEntityOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteEntityOptions, "deleteEntityOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteEntityOptions, "deleteEntityOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces", "entities"}
	pathParameters := []string{*deleteEntityOptions.WorkspaceID, *deleteEntityOptions.Entity}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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
	builder.AddQuery("version", assistant.Version)

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
//
// This operation is limited to 200 requests per 30 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) ListMentions(listMentionsOptions *ListMentionsOptions) (result *EntityMentionCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listMentionsOptions, "listMentionsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listMentionsOptions, "listMentionsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces", "entities", "mentions"}
	pathParameters := []string{*listMentionsOptions.WorkspaceID, *listMentionsOptions.Entity}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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

	if listMentionsOptions.Export != nil {
		builder.AddQuery("export", fmt.Sprint(*listMentionsOptions.Export))
	}
	if listMentionsOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*listMentionsOptions.IncludeAudit))
	}
	builder.AddQuery("version", assistant.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, new(EntityMentionCollection))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*EntityMentionCollection)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// ListValues : List entity values
// List the values for an entity.
//
// This operation is limited to 2500 requests per 30 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) ListValues(listValuesOptions *ListValuesOptions) (result *ValueCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listValuesOptions, "listValuesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listValuesOptions, "listValuesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces", "entities", "values"}
	pathParameters := []string{*listValuesOptions.WorkspaceID, *listValuesOptions.Entity}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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

	if listValuesOptions.Export != nil {
		builder.AddQuery("export", fmt.Sprint(*listValuesOptions.Export))
	}
	if listValuesOptions.PageLimit != nil {
		builder.AddQuery("page_limit", fmt.Sprint(*listValuesOptions.PageLimit))
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
	builder.AddQuery("version", assistant.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, new(ValueCollection))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*ValueCollection)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// CreateValue : Create entity value
// Create a new value for an entity.
//
// If you want to create multiple entity values with a single API call, consider using the **[Update
// entity](#update-entity)** method instead.
//
// This operation is limited to 1000 requests per 30 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) CreateValue(createValueOptions *CreateValueOptions) (result *Value, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createValueOptions, "createValueOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createValueOptions, "createValueOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces", "entities", "values"}
	pathParameters := []string{*createValueOptions.WorkspaceID, *createValueOptions.Entity}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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

	if createValueOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*createValueOptions.IncludeAudit))
	}
	builder.AddQuery("version", assistant.Version)

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

	response, err = assistant.Service.Request(request, new(Value))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*Value)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// GetValue : Get entity value
// Get information about an entity value.
//
// This operation is limited to 6000 requests per 5 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) GetValue(getValueOptions *GetValueOptions) (result *Value, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getValueOptions, "getValueOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getValueOptions, "getValueOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces", "entities", "values"}
	pathParameters := []string{*getValueOptions.WorkspaceID, *getValueOptions.Entity, *getValueOptions.Value}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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

	if getValueOptions.Export != nil {
		builder.AddQuery("export", fmt.Sprint(*getValueOptions.Export))
	}
	if getValueOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*getValueOptions.IncludeAudit))
	}
	builder.AddQuery("version", assistant.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, new(Value))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*Value)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// UpdateValue : Update entity value
// Update an existing entity value with new or modified data. You must provide component objects defining the content of
// the updated entity value.
//
// If you want to update multiple entity values with a single API call, consider using the **[Update
// entity](#update-entity)** method instead.
//
// This operation is limited to 1000 requests per 30 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) UpdateValue(updateValueOptions *UpdateValueOptions) (result *Value, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateValueOptions, "updateValueOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateValueOptions, "updateValueOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces", "entities", "values"}
	pathParameters := []string{*updateValueOptions.WorkspaceID, *updateValueOptions.Entity, *updateValueOptions.Value}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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

	if updateValueOptions.Append != nil {
		builder.AddQuery("append", fmt.Sprint(*updateValueOptions.Append))
	}
	if updateValueOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*updateValueOptions.IncludeAudit))
	}
	builder.AddQuery("version", assistant.Version)

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

	response, err = assistant.Service.Request(request, new(Value))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*Value)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// DeleteValue : Delete entity value
// Delete a value from an entity.
//
// This operation is limited to 1000 requests per 30 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) DeleteValue(deleteValueOptions *DeleteValueOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteValueOptions, "deleteValueOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteValueOptions, "deleteValueOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces", "entities", "values"}
	pathParameters := []string{*deleteValueOptions.WorkspaceID, *deleteValueOptions.Entity, *deleteValueOptions.Value}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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
	builder.AddQuery("version", assistant.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, nil)

	return
}

// ListSynonyms : List entity value synonyms
// List the synonyms for an entity value.
//
// This operation is limited to 2500 requests per 30 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) ListSynonyms(listSynonymsOptions *ListSynonymsOptions) (result *SynonymCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listSynonymsOptions, "listSynonymsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listSynonymsOptions, "listSynonymsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces", "entities", "values", "synonyms"}
	pathParameters := []string{*listSynonymsOptions.WorkspaceID, *listSynonymsOptions.Entity, *listSynonymsOptions.Value}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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

	if listSynonymsOptions.PageLimit != nil {
		builder.AddQuery("page_limit", fmt.Sprint(*listSynonymsOptions.PageLimit))
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
	builder.AddQuery("version", assistant.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, new(SynonymCollection))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*SynonymCollection)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// CreateSynonym : Create entity value synonym
// Add a new synonym to an entity value.
//
// If you want to create multiple synonyms with a single API call, consider using the **[Update
// entity](#update-entity)** or **[Update entity value](#update-entity-value)** method instead.
//
// This operation is limited to 1000 requests per 30 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) CreateSynonym(createSynonymOptions *CreateSynonymOptions) (result *Synonym, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createSynonymOptions, "createSynonymOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createSynonymOptions, "createSynonymOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces", "entities", "values", "synonyms"}
	pathParameters := []string{*createSynonymOptions.WorkspaceID, *createSynonymOptions.Entity, *createSynonymOptions.Value}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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

	if createSynonymOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*createSynonymOptions.IncludeAudit))
	}
	builder.AddQuery("version", assistant.Version)

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

	response, err = assistant.Service.Request(request, new(Synonym))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*Synonym)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// GetSynonym : Get entity value synonym
// Get information about a synonym of an entity value.
//
// This operation is limited to 6000 requests per 5 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) GetSynonym(getSynonymOptions *GetSynonymOptions) (result *Synonym, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getSynonymOptions, "getSynonymOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getSynonymOptions, "getSynonymOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces", "entities", "values", "synonyms"}
	pathParameters := []string{*getSynonymOptions.WorkspaceID, *getSynonymOptions.Entity, *getSynonymOptions.Value, *getSynonymOptions.Synonym}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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

	if getSynonymOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*getSynonymOptions.IncludeAudit))
	}
	builder.AddQuery("version", assistant.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, new(Synonym))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*Synonym)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// UpdateSynonym : Update entity value synonym
// Update an existing entity value synonym with new text.
//
// If you want to update multiple synonyms with a single API call, consider using the **[Update
// entity](#update-entity)** or **[Update entity value](#update-entity-value)** method instead.
//
// This operation is limited to 1000 requests per 30 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) UpdateSynonym(updateSynonymOptions *UpdateSynonymOptions) (result *Synonym, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateSynonymOptions, "updateSynonymOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateSynonymOptions, "updateSynonymOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces", "entities", "values", "synonyms"}
	pathParameters := []string{*updateSynonymOptions.WorkspaceID, *updateSynonymOptions.Entity, *updateSynonymOptions.Value, *updateSynonymOptions.Synonym}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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

	if updateSynonymOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*updateSynonymOptions.IncludeAudit))
	}
	builder.AddQuery("version", assistant.Version)

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

	response, err = assistant.Service.Request(request, new(Synonym))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*Synonym)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// DeleteSynonym : Delete entity value synonym
// Delete a synonym from an entity value.
//
// This operation is limited to 1000 requests per 30 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) DeleteSynonym(deleteSynonymOptions *DeleteSynonymOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteSynonymOptions, "deleteSynonymOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteSynonymOptions, "deleteSynonymOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces", "entities", "values", "synonyms"}
	pathParameters := []string{*deleteSynonymOptions.WorkspaceID, *deleteSynonymOptions.Entity, *deleteSynonymOptions.Value, *deleteSynonymOptions.Synonym}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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
	builder.AddQuery("version", assistant.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, nil)

	return
}

// ListDialogNodes : List dialog nodes
// List the dialog nodes for a workspace.
//
// This operation is limited to 2500 requests per 30 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) ListDialogNodes(listDialogNodesOptions *ListDialogNodesOptions) (result *DialogNodeCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listDialogNodesOptions, "listDialogNodesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listDialogNodesOptions, "listDialogNodesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces", "dialog_nodes"}
	pathParameters := []string{*listDialogNodesOptions.WorkspaceID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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

	if listDialogNodesOptions.PageLimit != nil {
		builder.AddQuery("page_limit", fmt.Sprint(*listDialogNodesOptions.PageLimit))
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
	builder.AddQuery("version", assistant.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, new(DialogNodeCollection))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*DialogNodeCollection)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// CreateDialogNode : Create dialog node
// Create a new dialog node.
//
// If you want to create multiple dialog nodes with a single API call, consider using the **[Update
// workspace](#update-workspace)** method instead.
//
// This operation is limited to 500 requests per 30 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) CreateDialogNode(createDialogNodeOptions *CreateDialogNodeOptions) (result *DialogNode, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createDialogNodeOptions, "createDialogNodeOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createDialogNodeOptions, "createDialogNodeOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces", "dialog_nodes"}
	pathParameters := []string{*createDialogNodeOptions.WorkspaceID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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

	if createDialogNodeOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*createDialogNodeOptions.IncludeAudit))
	}
	builder.AddQuery("version", assistant.Version)

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

	response, err = assistant.Service.Request(request, new(DialogNode))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*DialogNode)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// GetDialogNode : Get dialog node
// Get information about a dialog node.
//
// This operation is limited to 6000 requests per 5 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) GetDialogNode(getDialogNodeOptions *GetDialogNodeOptions) (result *DialogNode, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDialogNodeOptions, "getDialogNodeOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getDialogNodeOptions, "getDialogNodeOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces", "dialog_nodes"}
	pathParameters := []string{*getDialogNodeOptions.WorkspaceID, *getDialogNodeOptions.DialogNode}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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

	if getDialogNodeOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*getDialogNodeOptions.IncludeAudit))
	}
	builder.AddQuery("version", assistant.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, new(DialogNode))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*DialogNode)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// UpdateDialogNode : Update dialog node
// Update an existing dialog node with new or modified data.
//
// If you want to update multiple dialog nodes with a single API call, consider using the **[Update
// workspace](#update-workspace)** method instead.
//
// This operation is limited to 500 requests per 30 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) UpdateDialogNode(updateDialogNodeOptions *UpdateDialogNodeOptions) (result *DialogNode, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateDialogNodeOptions, "updateDialogNodeOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateDialogNodeOptions, "updateDialogNodeOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces", "dialog_nodes"}
	pathParameters := []string{*updateDialogNodeOptions.WorkspaceID, *updateDialogNodeOptions.DialogNode}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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

	if updateDialogNodeOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*updateDialogNodeOptions.IncludeAudit))
	}
	builder.AddQuery("version", assistant.Version)

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

	response, err = assistant.Service.Request(request, new(DialogNode))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*DialogNode)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// DeleteDialogNode : Delete dialog node
// Delete a dialog node from a workspace.
//
// This operation is limited to 500 requests per 30 minutes. For more information, see **Rate limiting**.
func (assistant *AssistantV1) DeleteDialogNode(deleteDialogNodeOptions *DeleteDialogNodeOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteDialogNodeOptions, "deleteDialogNodeOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteDialogNodeOptions, "deleteDialogNodeOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces", "dialog_nodes"}
	pathParameters := []string{*deleteDialogNodeOptions.WorkspaceID, *deleteDialogNodeOptions.DialogNode}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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
	builder.AddQuery("version", assistant.Version)

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
// If **cursor** is not specified, this operation is limited to 40 requests per 30 minutes. If **cursor** is specified,
// the limit is 120 requests per minute. For more information, see **Rate limiting**.
func (assistant *AssistantV1) ListLogs(listLogsOptions *ListLogsOptions) (result *LogCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listLogsOptions, "listLogsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listLogsOptions, "listLogsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/workspaces", "logs"}
	pathParameters := []string{*listLogsOptions.WorkspaceID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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
	builder.AddQuery("version", assistant.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, new(LogCollection))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*LogCollection)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// ListAllLogs : List log events in all workspaces
// List the events from the logs of all workspaces in the service instance.
//
// If **cursor** is not specified, this operation is limited to 40 requests per 30 minutes. If **cursor** is specified,
// the limit is 120 requests per minute. For more information, see **Rate limiting**.
func (assistant *AssistantV1) ListAllLogs(listAllLogsOptions *ListAllLogsOptions) (result *LogCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listAllLogsOptions, "listAllLogsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listAllLogsOptions, "listAllLogsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/logs"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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
	builder.AddQuery("version", assistant.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, new(LogCollection))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*LogCollection)
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
// security](https://cloud.ibm.com/docs/assistant?topic=assistant-information-security#information-security).
//
// This operation is limited to 4 requests per minute. For more information, see **Rate limiting**.
func (assistant *AssistantV1) DeleteUserData(deleteUserDataOptions *DeleteUserDataOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteUserDataOptions, "deleteUserDataOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteUserDataOptions, "deleteUserDataOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/user_data"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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

	builder.AddQuery("customer_id", fmt.Sprint(*deleteUserDataOptions.CustomerID))
	builder.AddQuery("version", assistant.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, nil)

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
func (assistant *AssistantV1) NewCaptureGroup(group string) (model *CaptureGroup, err error) {
	model = &CaptureGroup{
		Group: core.StringPtr(group),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// Context : State information for the conversation. To maintain state, include the context from the previous response.
type Context map[string]interface{}

// SetConversationID : Allow user to set ConversationID
func (this *Context) SetConversationID(ConversationID *string) {
	(*this)["conversation_id"] = ConversationID
}

// GetConversationID : Allow user to get ConversationID
func (this *Context) GetConversationID() *string {
	return (*this)["conversation_id"].(*string)
}

// SetSystem : Allow user to set System
func (this *Context) SetSystem(System *SystemResponse) {
	(*this)["system"] = System
}

// GetSystem : Allow user to get System
func (this *Context) GetSystem() *SystemResponse {
	return (*this)["system"].(*SystemResponse)
}

// SetMetadata : Allow user to set Metadata
func (this *Context) SetMetadata(Metadata *MessageContextMetadata) {
	(*this)["metadata"] = Metadata
}

// GetMetadata : Allow user to get Metadata
func (this *Context) GetMetadata() *MessageContextMetadata {
	return (*this)["metadata"].(*MessageContextMetadata)
}

// SetProperty : Allow user to set arbitrary property
func (this *Context) SetProperty(Key string, Value *interface{}) {
	(*this)[Key] = Value
}

// GetProperty : Allow user to get arbitrary property
func (this *Context) GetProperty(Key string) *interface{} {
	return (*this)[Key].(*interface{})
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
func (assistant *AssistantV1) NewCounterexample(text string) (model *Counterexample, err error) {
	model = &Counterexample{
		Text: core.StringPtr(text),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// CounterexampleCollection : CounterexampleCollection struct
type CounterexampleCollection struct {

	// An array of objects describing the examples marked as irrelevant input.
	Counterexamples []Counterexample `json:"counterexamples" validate:"required"`

	// The pagination data for the returned objects.
	Pagination *Pagination `json:"pagination" validate:"required"`
}

// CreateCounterexampleOptions : The CreateCounterexample options.
type CreateCounterexampleOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The text of a user input marked as irrelevant input. This string must conform to the following restrictions:
	// - It cannot contain carriage return, newline, or tab characters.
	// - It cannot consist of only whitespace characters.
	Text *string `json:"text" validate:"required"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewCreateCounterexampleOptions : Instantiate CreateCounterexampleOptions
func (assistant *AssistantV1) NewCreateCounterexampleOptions(workspaceID string, text string) *CreateCounterexampleOptions {
	return &CreateCounterexampleOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Text:        core.StringPtr(text),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *CreateCounterexampleOptions) SetWorkspaceID(workspaceID string) *CreateCounterexampleOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetText : Allow user to set Text
func (options *CreateCounterexampleOptions) SetText(text string) *CreateCounterexampleOptions {
	options.Text = core.StringPtr(text)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *CreateCounterexampleOptions) SetIncludeAudit(includeAudit bool) *CreateCounterexampleOptions {
	options.IncludeAudit = core.BoolPtr(includeAudit)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateCounterexampleOptions) SetHeaders(param map[string]string) *CreateCounterexampleOptions {
	options.Headers = param
	return options
}

// CreateDialogNodeOptions : The CreateDialogNode options.
type CreateDialogNodeOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The dialog node ID. This string must conform to the following restrictions:
	// - It can contain only Unicode alphanumeric, space, underscore, hyphen, and dot characters.
	DialogNode *string `json:"dialog_node" validate:"required"`

	// The description of the dialog node. This string cannot contain carriage return, newline, or tab characters.
	Description *string `json:"description,omitempty"`

	// The condition that will trigger the dialog node. This string cannot contain carriage return, newline, or tab
	// characters.
	Conditions *string `json:"conditions,omitempty"`

	// The ID of the parent dialog node. This property is omitted if the dialog node has no parent.
	Parent *string `json:"parent,omitempty"`

	// The ID of the previous sibling dialog node. This property is omitted if the dialog node has no previous sibling.
	PreviousSibling *string `json:"previous_sibling,omitempty"`

	// The output of the dialog node. For more information about how to specify dialog node output, see the
	// [documentation](https://cloud.ibm.com/docs/assistant?topic=assistant-dialog-overview#dialog-overview-responses).
	Output *DialogNodeOutput `json:"output,omitempty"`

	// The context for the dialog node.
	Context map[string]interface{} `json:"context,omitempty"`

	// The metadata for the dialog node.
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// The next step to execute following this dialog node.
	NextStep *DialogNodeNextStep `json:"next_step,omitempty"`

	// The alias used to identify the dialog node. This string must conform to the following restrictions:
	// - It can contain only Unicode alphanumeric, space, underscore, hyphen, and dot characters.
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

	// A label that can be displayed externally to describe the purpose of the node to users.
	UserLabel *string `json:"user_label,omitempty"`

	// Whether the dialog node should be excluded from disambiguation suggestions. Valid only when **type**=`standard` or
	// `frame`.
	DisambiguationOptOut *bool `json:"disambiguation_opt_out,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the CreateDialogNodeOptions.Type property.
// How the dialog node is processed.
const (
	CreateDialogNodeOptions_Type_EventHandler      = "event_handler"
	CreateDialogNodeOptions_Type_Folder            = "folder"
	CreateDialogNodeOptions_Type_Frame             = "frame"
	CreateDialogNodeOptions_Type_ResponseCondition = "response_condition"
	CreateDialogNodeOptions_Type_Slot              = "slot"
	CreateDialogNodeOptions_Type_Standard          = "standard"
)

// Constants associated with the CreateDialogNodeOptions.EventName property.
// How an `event_handler` node is processed.
const (
	CreateDialogNodeOptions_EventName_DigressionReturnPrompt   = "digression_return_prompt"
	CreateDialogNodeOptions_EventName_Filled                   = "filled"
	CreateDialogNodeOptions_EventName_FilledMultiple           = "filled_multiple"
	CreateDialogNodeOptions_EventName_Focus                    = "focus"
	CreateDialogNodeOptions_EventName_Generic                  = "generic"
	CreateDialogNodeOptions_EventName_Input                    = "input"
	CreateDialogNodeOptions_EventName_Nomatch                  = "nomatch"
	CreateDialogNodeOptions_EventName_NomatchResponsesDepleted = "nomatch_responses_depleted"
	CreateDialogNodeOptions_EventName_Validate                 = "validate"
)

// Constants associated with the CreateDialogNodeOptions.DigressIn property.
// Whether this top-level dialog node can be digressed into.
const (
	CreateDialogNodeOptions_DigressIn_DoesNotReturn = "does_not_return"
	CreateDialogNodeOptions_DigressIn_NotAvailable  = "not_available"
	CreateDialogNodeOptions_DigressIn_Returns       = "returns"
)

// Constants associated with the CreateDialogNodeOptions.DigressOut property.
// Whether this dialog node can be returned to after a digression.
const (
	CreateDialogNodeOptions_DigressOut_AllowAll            = "allow_all"
	CreateDialogNodeOptions_DigressOut_AllowAllNeverReturn = "allow_all_never_return"
	CreateDialogNodeOptions_DigressOut_AllowReturning      = "allow_returning"
)

// Constants associated with the CreateDialogNodeOptions.DigressOutSlots property.
// Whether the user can digress to top-level nodes while filling out slots.
const (
	CreateDialogNodeOptions_DigressOutSlots_AllowAll       = "allow_all"
	CreateDialogNodeOptions_DigressOutSlots_AllowReturning = "allow_returning"
	CreateDialogNodeOptions_DigressOutSlots_NotAllowed     = "not_allowed"
)

// NewCreateDialogNodeOptions : Instantiate CreateDialogNodeOptions
func (assistant *AssistantV1) NewCreateDialogNodeOptions(workspaceID string, dialogNode string) *CreateDialogNodeOptions {
	return &CreateDialogNodeOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		DialogNode:  core.StringPtr(dialogNode),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *CreateDialogNodeOptions) SetWorkspaceID(workspaceID string) *CreateDialogNodeOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetDialogNode : Allow user to set DialogNode
func (options *CreateDialogNodeOptions) SetDialogNode(dialogNode string) *CreateDialogNodeOptions {
	options.DialogNode = core.StringPtr(dialogNode)
	return options
}

// SetDescription : Allow user to set Description
func (options *CreateDialogNodeOptions) SetDescription(description string) *CreateDialogNodeOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetConditions : Allow user to set Conditions
func (options *CreateDialogNodeOptions) SetConditions(conditions string) *CreateDialogNodeOptions {
	options.Conditions = core.StringPtr(conditions)
	return options
}

// SetParent : Allow user to set Parent
func (options *CreateDialogNodeOptions) SetParent(parent string) *CreateDialogNodeOptions {
	options.Parent = core.StringPtr(parent)
	return options
}

// SetPreviousSibling : Allow user to set PreviousSibling
func (options *CreateDialogNodeOptions) SetPreviousSibling(previousSibling string) *CreateDialogNodeOptions {
	options.PreviousSibling = core.StringPtr(previousSibling)
	return options
}

// SetOutput : Allow user to set Output
func (options *CreateDialogNodeOptions) SetOutput(output *DialogNodeOutput) *CreateDialogNodeOptions {
	options.Output = output
	return options
}

// SetContext : Allow user to set Context
func (options *CreateDialogNodeOptions) SetContext(context map[string]interface{}) *CreateDialogNodeOptions {
	options.Context = context
	return options
}

// SetMetadata : Allow user to set Metadata
func (options *CreateDialogNodeOptions) SetMetadata(metadata map[string]interface{}) *CreateDialogNodeOptions {
	options.Metadata = metadata
	return options
}

// SetNextStep : Allow user to set NextStep
func (options *CreateDialogNodeOptions) SetNextStep(nextStep *DialogNodeNextStep) *CreateDialogNodeOptions {
	options.NextStep = nextStep
	return options
}

// SetTitle : Allow user to set Title
func (options *CreateDialogNodeOptions) SetTitle(title string) *CreateDialogNodeOptions {
	options.Title = core.StringPtr(title)
	return options
}

// SetType : Allow user to set Type
func (options *CreateDialogNodeOptions) SetType(typeVar string) *CreateDialogNodeOptions {
	options.Type = core.StringPtr(typeVar)
	return options
}

// SetEventName : Allow user to set EventName
func (options *CreateDialogNodeOptions) SetEventName(eventName string) *CreateDialogNodeOptions {
	options.EventName = core.StringPtr(eventName)
	return options
}

// SetVariable : Allow user to set Variable
func (options *CreateDialogNodeOptions) SetVariable(variable string) *CreateDialogNodeOptions {
	options.Variable = core.StringPtr(variable)
	return options
}

// SetActions : Allow user to set Actions
func (options *CreateDialogNodeOptions) SetActions(actions []DialogNodeAction) *CreateDialogNodeOptions {
	options.Actions = actions
	return options
}

// SetDigressIn : Allow user to set DigressIn
func (options *CreateDialogNodeOptions) SetDigressIn(digressIn string) *CreateDialogNodeOptions {
	options.DigressIn = core.StringPtr(digressIn)
	return options
}

// SetDigressOut : Allow user to set DigressOut
func (options *CreateDialogNodeOptions) SetDigressOut(digressOut string) *CreateDialogNodeOptions {
	options.DigressOut = core.StringPtr(digressOut)
	return options
}

// SetDigressOutSlots : Allow user to set DigressOutSlots
func (options *CreateDialogNodeOptions) SetDigressOutSlots(digressOutSlots string) *CreateDialogNodeOptions {
	options.DigressOutSlots = core.StringPtr(digressOutSlots)
	return options
}

// SetUserLabel : Allow user to set UserLabel
func (options *CreateDialogNodeOptions) SetUserLabel(userLabel string) *CreateDialogNodeOptions {
	options.UserLabel = core.StringPtr(userLabel)
	return options
}

// SetDisambiguationOptOut : Allow user to set DisambiguationOptOut
func (options *CreateDialogNodeOptions) SetDisambiguationOptOut(disambiguationOptOut bool) *CreateDialogNodeOptions {
	options.DisambiguationOptOut = core.BoolPtr(disambiguationOptOut)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *CreateDialogNodeOptions) SetIncludeAudit(includeAudit bool) *CreateDialogNodeOptions {
	options.IncludeAudit = core.BoolPtr(includeAudit)
	return options
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
func (assistant *AssistantV1) NewCreateEntity(entity string) (model *CreateEntity, err error) {
	model = &CreateEntity{
		Entity: core.StringPtr(entity),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// CreateEntityOptions : The CreateEntity options.
type CreateEntityOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

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
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewCreateEntityOptions : Instantiate CreateEntityOptions
func (assistant *AssistantV1) NewCreateEntityOptions(workspaceID string, entity string) *CreateEntityOptions {
	return &CreateEntityOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Entity:      core.StringPtr(entity),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *CreateEntityOptions) SetWorkspaceID(workspaceID string) *CreateEntityOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetEntity : Allow user to set Entity
func (options *CreateEntityOptions) SetEntity(entity string) *CreateEntityOptions {
	options.Entity = core.StringPtr(entity)
	return options
}

// SetDescription : Allow user to set Description
func (options *CreateEntityOptions) SetDescription(description string) *CreateEntityOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetMetadata : Allow user to set Metadata
func (options *CreateEntityOptions) SetMetadata(metadata map[string]interface{}) *CreateEntityOptions {
	options.Metadata = metadata
	return options
}

// SetFuzzyMatch : Allow user to set FuzzyMatch
func (options *CreateEntityOptions) SetFuzzyMatch(fuzzyMatch bool) *CreateEntityOptions {
	options.FuzzyMatch = core.BoolPtr(fuzzyMatch)
	return options
}

// SetValues : Allow user to set Values
func (options *CreateEntityOptions) SetValues(values []CreateValue) *CreateEntityOptions {
	options.Values = values
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *CreateEntityOptions) SetIncludeAudit(includeAudit bool) *CreateEntityOptions {
	options.IncludeAudit = core.BoolPtr(includeAudit)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateEntityOptions) SetHeaders(param map[string]string) *CreateEntityOptions {
	options.Headers = param
	return options
}

// CreateExampleOptions : The CreateExample options.
type CreateExampleOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The intent name.
	Intent *string `json:"intent" validate:"required"`

	// The text of a user input example. This string must conform to the following restrictions:
	// - It cannot contain carriage return, newline, or tab characters.
	// - It cannot consist of only whitespace characters.
	Text *string `json:"text" validate:"required"`

	// An array of contextual entity mentions.
	Mentions []Mention `json:"mentions,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewCreateExampleOptions : Instantiate CreateExampleOptions
func (assistant *AssistantV1) NewCreateExampleOptions(workspaceID string, intent string, text string) *CreateExampleOptions {
	return &CreateExampleOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Intent:      core.StringPtr(intent),
		Text:        core.StringPtr(text),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *CreateExampleOptions) SetWorkspaceID(workspaceID string) *CreateExampleOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetIntent : Allow user to set Intent
func (options *CreateExampleOptions) SetIntent(intent string) *CreateExampleOptions {
	options.Intent = core.StringPtr(intent)
	return options
}

// SetText : Allow user to set Text
func (options *CreateExampleOptions) SetText(text string) *CreateExampleOptions {
	options.Text = core.StringPtr(text)
	return options
}

// SetMentions : Allow user to set Mentions
func (options *CreateExampleOptions) SetMentions(mentions []Mention) *CreateExampleOptions {
	options.Mentions = mentions
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *CreateExampleOptions) SetIncludeAudit(includeAudit bool) *CreateExampleOptions {
	options.IncludeAudit = core.BoolPtr(includeAudit)
	return options
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
func (assistant *AssistantV1) NewCreateIntent(intent string) (model *CreateIntent, err error) {
	model = &CreateIntent{
		Intent: core.StringPtr(intent),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// CreateIntentOptions : The CreateIntent options.
type CreateIntentOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The name of the intent. This string must conform to the following restrictions:
	// - It can contain only Unicode alphanumeric, underscore, hyphen, and dot characters.
	// - It cannot begin with the reserved prefix `sys-`.
	Intent *string `json:"intent" validate:"required"`

	// The description of the intent. This string cannot contain carriage return, newline, or tab characters.
	Description *string `json:"description,omitempty"`

	// An array of user input examples for the intent.
	Examples []Example `json:"examples,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewCreateIntentOptions : Instantiate CreateIntentOptions
func (assistant *AssistantV1) NewCreateIntentOptions(workspaceID string, intent string) *CreateIntentOptions {
	return &CreateIntentOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Intent:      core.StringPtr(intent),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *CreateIntentOptions) SetWorkspaceID(workspaceID string) *CreateIntentOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetIntent : Allow user to set Intent
func (options *CreateIntentOptions) SetIntent(intent string) *CreateIntentOptions {
	options.Intent = core.StringPtr(intent)
	return options
}

// SetDescription : Allow user to set Description
func (options *CreateIntentOptions) SetDescription(description string) *CreateIntentOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetExamples : Allow user to set Examples
func (options *CreateIntentOptions) SetExamples(examples []Example) *CreateIntentOptions {
	options.Examples = examples
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *CreateIntentOptions) SetIncludeAudit(includeAudit bool) *CreateIntentOptions {
	options.IncludeAudit = core.BoolPtr(includeAudit)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateIntentOptions) SetHeaders(param map[string]string) *CreateIntentOptions {
	options.Headers = param
	return options
}

// CreateSynonymOptions : The CreateSynonym options.
type CreateSynonymOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The name of the entity.
	Entity *string `json:"entity" validate:"required"`

	// The text of the entity value.
	Value *string `json:"value" validate:"required"`

	// The text of the synonym. This string must conform to the following restrictions:
	// - It cannot contain carriage return, newline, or tab characters.
	// - It cannot consist of only whitespace characters.
	Synonym *string `json:"synonym" validate:"required"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewCreateSynonymOptions : Instantiate CreateSynonymOptions
func (assistant *AssistantV1) NewCreateSynonymOptions(workspaceID string, entity string, value string, synonym string) *CreateSynonymOptions {
	return &CreateSynonymOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Entity:      core.StringPtr(entity),
		Value:       core.StringPtr(value),
		Synonym:     core.StringPtr(synonym),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *CreateSynonymOptions) SetWorkspaceID(workspaceID string) *CreateSynonymOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetEntity : Allow user to set Entity
func (options *CreateSynonymOptions) SetEntity(entity string) *CreateSynonymOptions {
	options.Entity = core.StringPtr(entity)
	return options
}

// SetValue : Allow user to set Value
func (options *CreateSynonymOptions) SetValue(value string) *CreateSynonymOptions {
	options.Value = core.StringPtr(value)
	return options
}

// SetSynonym : Allow user to set Synonym
func (options *CreateSynonymOptions) SetSynonym(synonym string) *CreateSynonymOptions {
	options.Synonym = core.StringPtr(synonym)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *CreateSynonymOptions) SetIncludeAudit(includeAudit bool) *CreateSynonymOptions {
	options.IncludeAudit = core.BoolPtr(includeAudit)
	return options
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
	CreateValue_Type_Patterns = "patterns"
	CreateValue_Type_Synonyms = "synonyms"
)

// NewCreateValue : Instantiate CreateValue (Generic Model Constructor)
func (assistant *AssistantV1) NewCreateValue(value string) (model *CreateValue, err error) {
	model = &CreateValue{
		Value: core.StringPtr(value),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// CreateValueOptions : The CreateValue options.
type CreateValueOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The name of the entity.
	Entity *string `json:"entity" validate:"required"`

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
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the CreateValueOptions.Type property.
// Specifies the type of entity value.
const (
	CreateValueOptions_Type_Patterns = "patterns"
	CreateValueOptions_Type_Synonyms = "synonyms"
)

// NewCreateValueOptions : Instantiate CreateValueOptions
func (assistant *AssistantV1) NewCreateValueOptions(workspaceID string, entity string, value string) *CreateValueOptions {
	return &CreateValueOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Entity:      core.StringPtr(entity),
		Value:       core.StringPtr(value),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *CreateValueOptions) SetWorkspaceID(workspaceID string) *CreateValueOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetEntity : Allow user to set Entity
func (options *CreateValueOptions) SetEntity(entity string) *CreateValueOptions {
	options.Entity = core.StringPtr(entity)
	return options
}

// SetValue : Allow user to set Value
func (options *CreateValueOptions) SetValue(value string) *CreateValueOptions {
	options.Value = core.StringPtr(value)
	return options
}

// SetMetadata : Allow user to set Metadata
func (options *CreateValueOptions) SetMetadata(metadata map[string]interface{}) *CreateValueOptions {
	options.Metadata = metadata
	return options
}

// SetType : Allow user to set Type
func (options *CreateValueOptions) SetType(typeVar string) *CreateValueOptions {
	options.Type = core.StringPtr(typeVar)
	return options
}

// SetSynonyms : Allow user to set Synonyms
func (options *CreateValueOptions) SetSynonyms(synonyms []string) *CreateValueOptions {
	options.Synonyms = synonyms
	return options
}

// SetPatterns : Allow user to set Patterns
func (options *CreateValueOptions) SetPatterns(patterns []string) *CreateValueOptions {
	options.Patterns = patterns
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *CreateValueOptions) SetIncludeAudit(includeAudit bool) *CreateValueOptions {
	options.IncludeAudit = core.BoolPtr(includeAudit)
	return options
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

	// Any metadata related to the workspace.
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// Whether training data from the workspace (including artifacts such as intents and entities) can be used by IBM for
	// general service improvements. `true` indicates that workspace training data is not to be used.
	LearningOptOut *bool `json:"learning_opt_out,omitempty"`

	// Global settings for the workspace.
	SystemSettings *WorkspaceSystemSettings `json:"system_settings,omitempty"`

	// An array of objects defining the intents for the workspace.
	Intents []CreateIntent `json:"intents,omitempty"`

	// An array of objects describing the entities for the workspace.
	Entities []CreateEntity `json:"entities,omitempty"`

	// An array of objects describing the dialog nodes in the workspace.
	DialogNodes []DialogNode `json:"dialog_nodes,omitempty"`

	// An array of objects defining input examples that have been marked as irrelevant input.
	Counterexamples []Counterexample `json:"counterexamples,omitempty"`

	Webhooks []Webhook `json:"webhooks,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewCreateWorkspaceOptions : Instantiate CreateWorkspaceOptions
func (assistant *AssistantV1) NewCreateWorkspaceOptions() *CreateWorkspaceOptions {
	return &CreateWorkspaceOptions{}
}

// SetName : Allow user to set Name
func (options *CreateWorkspaceOptions) SetName(name string) *CreateWorkspaceOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetDescription : Allow user to set Description
func (options *CreateWorkspaceOptions) SetDescription(description string) *CreateWorkspaceOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetLanguage : Allow user to set Language
func (options *CreateWorkspaceOptions) SetLanguage(language string) *CreateWorkspaceOptions {
	options.Language = core.StringPtr(language)
	return options
}

// SetMetadata : Allow user to set Metadata
func (options *CreateWorkspaceOptions) SetMetadata(metadata map[string]interface{}) *CreateWorkspaceOptions {
	options.Metadata = metadata
	return options
}

// SetLearningOptOut : Allow user to set LearningOptOut
func (options *CreateWorkspaceOptions) SetLearningOptOut(learningOptOut bool) *CreateWorkspaceOptions {
	options.LearningOptOut = core.BoolPtr(learningOptOut)
	return options
}

// SetSystemSettings : Allow user to set SystemSettings
func (options *CreateWorkspaceOptions) SetSystemSettings(systemSettings *WorkspaceSystemSettings) *CreateWorkspaceOptions {
	options.SystemSettings = systemSettings
	return options
}

// SetIntents : Allow user to set Intents
func (options *CreateWorkspaceOptions) SetIntents(intents []CreateIntent) *CreateWorkspaceOptions {
	options.Intents = intents
	return options
}

// SetEntities : Allow user to set Entities
func (options *CreateWorkspaceOptions) SetEntities(entities []CreateEntity) *CreateWorkspaceOptions {
	options.Entities = entities
	return options
}

// SetDialogNodes : Allow user to set DialogNodes
func (options *CreateWorkspaceOptions) SetDialogNodes(dialogNodes []DialogNode) *CreateWorkspaceOptions {
	options.DialogNodes = dialogNodes
	return options
}

// SetCounterexamples : Allow user to set Counterexamples
func (options *CreateWorkspaceOptions) SetCounterexamples(counterexamples []Counterexample) *CreateWorkspaceOptions {
	options.Counterexamples = counterexamples
	return options
}

// SetWebhooks : Allow user to set Webhooks
func (options *CreateWorkspaceOptions) SetWebhooks(webhooks []Webhook) *CreateWorkspaceOptions {
	options.Webhooks = webhooks
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *CreateWorkspaceOptions) SetIncludeAudit(includeAudit bool) *CreateWorkspaceOptions {
	options.IncludeAudit = core.BoolPtr(includeAudit)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateWorkspaceOptions) SetHeaders(param map[string]string) *CreateWorkspaceOptions {
	options.Headers = param
	return options
}

// DeleteCounterexampleOptions : The DeleteCounterexample options.
type DeleteCounterexampleOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The text of a user input counterexample (for example, `What are you wearing?`).
	Text *string `json:"text" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteCounterexampleOptions : Instantiate DeleteCounterexampleOptions
func (assistant *AssistantV1) NewDeleteCounterexampleOptions(workspaceID string, text string) *DeleteCounterexampleOptions {
	return &DeleteCounterexampleOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Text:        core.StringPtr(text),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *DeleteCounterexampleOptions) SetWorkspaceID(workspaceID string) *DeleteCounterexampleOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetText : Allow user to set Text
func (options *DeleteCounterexampleOptions) SetText(text string) *DeleteCounterexampleOptions {
	options.Text = core.StringPtr(text)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteCounterexampleOptions) SetHeaders(param map[string]string) *DeleteCounterexampleOptions {
	options.Headers = param
	return options
}

// DeleteDialogNodeOptions : The DeleteDialogNode options.
type DeleteDialogNodeOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The dialog node ID (for example, `get_order`).
	DialogNode *string `json:"dialog_node" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteDialogNodeOptions : Instantiate DeleteDialogNodeOptions
func (assistant *AssistantV1) NewDeleteDialogNodeOptions(workspaceID string, dialogNode string) *DeleteDialogNodeOptions {
	return &DeleteDialogNodeOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		DialogNode:  core.StringPtr(dialogNode),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *DeleteDialogNodeOptions) SetWorkspaceID(workspaceID string) *DeleteDialogNodeOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetDialogNode : Allow user to set DialogNode
func (options *DeleteDialogNodeOptions) SetDialogNode(dialogNode string) *DeleteDialogNodeOptions {
	options.DialogNode = core.StringPtr(dialogNode)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteDialogNodeOptions) SetHeaders(param map[string]string) *DeleteDialogNodeOptions {
	options.Headers = param
	return options
}

// DeleteEntityOptions : The DeleteEntity options.
type DeleteEntityOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The name of the entity.
	Entity *string `json:"entity" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteEntityOptions : Instantiate DeleteEntityOptions
func (assistant *AssistantV1) NewDeleteEntityOptions(workspaceID string, entity string) *DeleteEntityOptions {
	return &DeleteEntityOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Entity:      core.StringPtr(entity),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *DeleteEntityOptions) SetWorkspaceID(workspaceID string) *DeleteEntityOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetEntity : Allow user to set Entity
func (options *DeleteEntityOptions) SetEntity(entity string) *DeleteEntityOptions {
	options.Entity = core.StringPtr(entity)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteEntityOptions) SetHeaders(param map[string]string) *DeleteEntityOptions {
	options.Headers = param
	return options
}

// DeleteExampleOptions : The DeleteExample options.
type DeleteExampleOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The intent name.
	Intent *string `json:"intent" validate:"required"`

	// The text of the user input example.
	Text *string `json:"text" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteExampleOptions : Instantiate DeleteExampleOptions
func (assistant *AssistantV1) NewDeleteExampleOptions(workspaceID string, intent string, text string) *DeleteExampleOptions {
	return &DeleteExampleOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Intent:      core.StringPtr(intent),
		Text:        core.StringPtr(text),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *DeleteExampleOptions) SetWorkspaceID(workspaceID string) *DeleteExampleOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetIntent : Allow user to set Intent
func (options *DeleteExampleOptions) SetIntent(intent string) *DeleteExampleOptions {
	options.Intent = core.StringPtr(intent)
	return options
}

// SetText : Allow user to set Text
func (options *DeleteExampleOptions) SetText(text string) *DeleteExampleOptions {
	options.Text = core.StringPtr(text)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteExampleOptions) SetHeaders(param map[string]string) *DeleteExampleOptions {
	options.Headers = param
	return options
}

// DeleteIntentOptions : The DeleteIntent options.
type DeleteIntentOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The intent name.
	Intent *string `json:"intent" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteIntentOptions : Instantiate DeleteIntentOptions
func (assistant *AssistantV1) NewDeleteIntentOptions(workspaceID string, intent string) *DeleteIntentOptions {
	return &DeleteIntentOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Intent:      core.StringPtr(intent),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *DeleteIntentOptions) SetWorkspaceID(workspaceID string) *DeleteIntentOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetIntent : Allow user to set Intent
func (options *DeleteIntentOptions) SetIntent(intent string) *DeleteIntentOptions {
	options.Intent = core.StringPtr(intent)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteIntentOptions) SetHeaders(param map[string]string) *DeleteIntentOptions {
	options.Headers = param
	return options
}

// DeleteSynonymOptions : The DeleteSynonym options.
type DeleteSynonymOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The name of the entity.
	Entity *string `json:"entity" validate:"required"`

	// The text of the entity value.
	Value *string `json:"value" validate:"required"`

	// The text of the synonym.
	Synonym *string `json:"synonym" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteSynonymOptions : Instantiate DeleteSynonymOptions
func (assistant *AssistantV1) NewDeleteSynonymOptions(workspaceID string, entity string, value string, synonym string) *DeleteSynonymOptions {
	return &DeleteSynonymOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Entity:      core.StringPtr(entity),
		Value:       core.StringPtr(value),
		Synonym:     core.StringPtr(synonym),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *DeleteSynonymOptions) SetWorkspaceID(workspaceID string) *DeleteSynonymOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetEntity : Allow user to set Entity
func (options *DeleteSynonymOptions) SetEntity(entity string) *DeleteSynonymOptions {
	options.Entity = core.StringPtr(entity)
	return options
}

// SetValue : Allow user to set Value
func (options *DeleteSynonymOptions) SetValue(value string) *DeleteSynonymOptions {
	options.Value = core.StringPtr(value)
	return options
}

// SetSynonym : Allow user to set Synonym
func (options *DeleteSynonymOptions) SetSynonym(synonym string) *DeleteSynonymOptions {
	options.Synonym = core.StringPtr(synonym)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteSynonymOptions) SetHeaders(param map[string]string) *DeleteSynonymOptions {
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
func (assistant *AssistantV1) NewDeleteUserDataOptions(customerID string) *DeleteUserDataOptions {
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

// DeleteValueOptions : The DeleteValue options.
type DeleteValueOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The name of the entity.
	Entity *string `json:"entity" validate:"required"`

	// The text of the entity value.
	Value *string `json:"value" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteValueOptions : Instantiate DeleteValueOptions
func (assistant *AssistantV1) NewDeleteValueOptions(workspaceID string, entity string, value string) *DeleteValueOptions {
	return &DeleteValueOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Entity:      core.StringPtr(entity),
		Value:       core.StringPtr(value),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *DeleteValueOptions) SetWorkspaceID(workspaceID string) *DeleteValueOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetEntity : Allow user to set Entity
func (options *DeleteValueOptions) SetEntity(entity string) *DeleteValueOptions {
	options.Entity = core.StringPtr(entity)
	return options
}

// SetValue : Allow user to set Value
func (options *DeleteValueOptions) SetValue(value string) *DeleteValueOptions {
	options.Value = core.StringPtr(value)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteValueOptions) SetHeaders(param map[string]string) *DeleteValueOptions {
	options.Headers = param
	return options
}

// DeleteWorkspaceOptions : The DeleteWorkspace options.
type DeleteWorkspaceOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteWorkspaceOptions : Instantiate DeleteWorkspaceOptions
func (assistant *AssistantV1) NewDeleteWorkspaceOptions(workspaceID string) *DeleteWorkspaceOptions {
	return &DeleteWorkspaceOptions{
		WorkspaceID: core.StringPtr(workspaceID),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *DeleteWorkspaceOptions) SetWorkspaceID(workspaceID string) *DeleteWorkspaceOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteWorkspaceOptions) SetHeaders(param map[string]string) *DeleteWorkspaceOptions {
	options.Headers = param
	return options
}

// DialogNode : DialogNode struct
type DialogNode struct {

	// The dialog node ID. This string must conform to the following restrictions:
	// - It can contain only Unicode alphanumeric, space, underscore, hyphen, and dot characters.
	DialogNode *string `json:"dialog_node" validate:"required"`

	// The description of the dialog node. This string cannot contain carriage return, newline, or tab characters.
	Description *string `json:"description,omitempty"`

	// The condition that will trigger the dialog node. This string cannot contain carriage return, newline, or tab
	// characters.
	Conditions *string `json:"conditions,omitempty"`

	// The ID of the parent dialog node. This property is omitted if the dialog node has no parent.
	Parent *string `json:"parent,omitempty"`

	// The ID of the previous sibling dialog node. This property is omitted if the dialog node has no previous sibling.
	PreviousSibling *string `json:"previous_sibling,omitempty"`

	// The output of the dialog node. For more information about how to specify dialog node output, see the
	// [documentation](https://cloud.ibm.com/docs/assistant?topic=assistant-dialog-overview#dialog-overview-responses).
	Output *DialogNodeOutput `json:"output,omitempty"`

	// The context for the dialog node.
	Context map[string]interface{} `json:"context,omitempty"`

	// The metadata for the dialog node.
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// The next step to execute following this dialog node.
	NextStep *DialogNodeNextStep `json:"next_step,omitempty"`

	// The alias used to identify the dialog node. This string must conform to the following restrictions:
	// - It can contain only Unicode alphanumeric, space, underscore, hyphen, and dot characters.
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

	// A label that can be displayed externally to describe the purpose of the node to users.
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
	DialogNode_Type_EventHandler      = "event_handler"
	DialogNode_Type_Folder            = "folder"
	DialogNode_Type_Frame             = "frame"
	DialogNode_Type_ResponseCondition = "response_condition"
	DialogNode_Type_Slot              = "slot"
	DialogNode_Type_Standard          = "standard"
)

// Constants associated with the DialogNode.EventName property.
// How an `event_handler` node is processed.
const (
	DialogNode_EventName_DigressionReturnPrompt   = "digression_return_prompt"
	DialogNode_EventName_Filled                   = "filled"
	DialogNode_EventName_FilledMultiple           = "filled_multiple"
	DialogNode_EventName_Focus                    = "focus"
	DialogNode_EventName_Generic                  = "generic"
	DialogNode_EventName_Input                    = "input"
	DialogNode_EventName_Nomatch                  = "nomatch"
	DialogNode_EventName_NomatchResponsesDepleted = "nomatch_responses_depleted"
	DialogNode_EventName_Validate                 = "validate"
)

// Constants associated with the DialogNode.DigressIn property.
// Whether this top-level dialog node can be digressed into.
const (
	DialogNode_DigressIn_DoesNotReturn = "does_not_return"
	DialogNode_DigressIn_NotAvailable  = "not_available"
	DialogNode_DigressIn_Returns       = "returns"
)

// Constants associated with the DialogNode.DigressOut property.
// Whether this dialog node can be returned to after a digression.
const (
	DialogNode_DigressOut_AllowAll            = "allow_all"
	DialogNode_DigressOut_AllowAllNeverReturn = "allow_all_never_return"
	DialogNode_DigressOut_AllowReturning      = "allow_returning"
)

// Constants associated with the DialogNode.DigressOutSlots property.
// Whether the user can digress to top-level nodes while filling out slots.
const (
	DialogNode_DigressOutSlots_AllowAll       = "allow_all"
	DialogNode_DigressOutSlots_AllowReturning = "allow_returning"
	DialogNode_DigressOutSlots_NotAllowed     = "not_allowed"
)

// NewDialogNode : Instantiate DialogNode (Generic Model Constructor)
func (assistant *AssistantV1) NewDialogNode(dialogNode string) (model *DialogNode, err error) {
	model = &DialogNode{
		DialogNode: core.StringPtr(dialogNode),
	}
	err = core.ValidateStruct(model, "required parameters")
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
	DialogNodeAction_Type_Client        = "client"
	DialogNodeAction_Type_CloudFunction = "cloud_function"
	DialogNodeAction_Type_Server        = "server"
	DialogNodeAction_Type_WebAction     = "web_action"
	DialogNodeAction_Type_Webhook       = "webhook"
)

// NewDialogNodeAction : Instantiate DialogNodeAction (Generic Model Constructor)
func (assistant *AssistantV1) NewDialogNodeAction(name string, resultVariable string) (model *DialogNodeAction, err error) {
	model = &DialogNodeAction{
		Name:           core.StringPtr(name),
		ResultVariable: core.StringPtr(resultVariable),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// DialogNodeCollection : An array of dialog nodes.
type DialogNodeCollection struct {

	// An array of objects describing the dialog nodes defined for the workspace.
	DialogNodes []DialogNode `json:"dialog_nodes" validate:"required"`

	// The pagination data for the returned objects.
	Pagination *Pagination `json:"pagination" validate:"required"`
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

	// The ID of the dialog node to process next. This parameter is required if **behavior**=`jump_to`.
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
	DialogNodeNextStep_Behavior_GetUserInput  = "get_user_input"
	DialogNodeNextStep_Behavior_JumpTo        = "jump_to"
	DialogNodeNextStep_Behavior_Reprompt      = "reprompt"
	DialogNodeNextStep_Behavior_SkipAllSlots  = "skip_all_slots"
	DialogNodeNextStep_Behavior_SkipSlot      = "skip_slot"
	DialogNodeNextStep_Behavior_SkipUserInput = "skip_user_input"
)

// Constants associated with the DialogNodeNextStep.Selector property.
// Which part of the dialog node to process next.
const (
	DialogNodeNextStep_Selector_Body      = "body"
	DialogNodeNextStep_Selector_Client    = "client"
	DialogNodeNextStep_Selector_Condition = "condition"
	DialogNodeNextStep_Selector_UserInput = "user_input"
)

// NewDialogNodeNextStep : Instantiate DialogNodeNextStep (Generic Model Constructor)
func (assistant *AssistantV1) NewDialogNodeNextStep(behavior string) (model *DialogNodeNextStep, err error) {
	model = &DialogNodeNextStep{
		Behavior: core.StringPtr(behavior),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// DialogNodeOutput : The output of the dialog node. For more information about how to specify dialog node output, see the
// [documentation](https://cloud.ibm.com/docs/assistant?topic=assistant-dialog-overview#dialog-overview-responses).
type DialogNodeOutput map[string]interface{}

// SetGeneric : Allow user to set Generic
func (this *DialogNodeOutput) SetGeneric(Generic *[]DialogNodeOutputGeneric) {
	(*this)["generic"] = Generic
}

// GetGeneric : Allow user to get Generic
func (this *DialogNodeOutput) GetGeneric() *[]DialogNodeOutputGeneric {
	return (*this)["generic"].(*[]DialogNodeOutputGeneric)
}

// SetModifiers : Allow user to set Modifiers
func (this *DialogNodeOutput) SetModifiers(Modifiers *DialogNodeOutputModifiers) {
	(*this)["modifiers"] = Modifiers
}

// GetModifiers : Allow user to get Modifiers
func (this *DialogNodeOutput) GetModifiers() *DialogNodeOutputModifiers {
	return (*this)["modifiers"].(*DialogNodeOutputModifiers)
}

// SetProperty : Allow user to set arbitrary property
func (this *DialogNodeOutput) SetProperty(Key string, Value *interface{}) {
	(*this)[Key] = Value
}

// GetProperty : Allow user to get arbitrary property
func (this *DialogNodeOutput) GetProperty(Key string) *interface{} {
	return (*this)[Key].(*interface{})
}

// DialogNodeOutputGeneric : DialogNodeOutputGeneric struct
type DialogNodeOutputGeneric struct {

	// The type of response returned by the dialog node. The specified response type must be supported by the client
	// application or channel.
	//
	// **Note:** The **search_skill** response type is available only for Plus and Premium users, and is used only by the
	// v2 runtime API.
	ResponseType *string `json:"response_type" validate:"required"`

	// A list of one or more objects defining text responses. Required when **response_type**=`text`.
	Values []DialogNodeOutputTextValuesElement `json:"values,omitempty"`

	// How a response is selected from the list, if more than one response is specified. Valid only when
	// **response_type**=`text`.
	SelectionPolicy *string `json:"selection_policy,omitempty"`

	// The delimiter to use as a separator between responses when `selection_policy`=`multiline`.
	Delimiter *string `json:"delimiter,omitempty"`

	// How long to pause, in milliseconds. The valid values are from 0 to 10000. Valid only when **response_type**=`pause`.
	Time *int64 `json:"time,omitempty"`

	// Whether to send a "user is typing" event during the pause. Ignored if the channel does not support this event. Valid
	// only when **response_type**=`pause`.
	Typing *bool `json:"typing,omitempty"`

	// The URL of the image. Required when **response_type**=`image`.
	Source *string `json:"source,omitempty"`

	// An optional title to show before the response. Valid only when **response_type**=`image` or `option`.
	Title *string `json:"title,omitempty"`

	// An optional description to show with the response. Valid only when **response_type**=`image` or `option`.
	Description *string `json:"description,omitempty"`

	// The preferred type of control to display, if supported by the channel. Valid only when **response_type**=`option`.
	Preference *string `json:"preference,omitempty"`

	// An array of objects describing the options from which the user can choose. You can include up to 20 options.
	// Required when **response_type**=`option`.
	Options []DialogNodeOutputOptionsElement `json:"options,omitempty"`

	// An optional message to be sent to the human agent who will be taking over the conversation. Valid only when
	// **reponse_type**=`connect_to_agent`.
	MessageToHumanAgent *string `json:"message_to_human_agent,omitempty"`

	// The text of the search query. This can be either a natural-language query or a query that uses the Discovery query
	// language syntax, depending on the value of the **query_type** property. For more information, see the [Discovery
	// service documentation](https://cloud.ibm.com/docs/discovery?topic=discovery-query-operators#query-operators).
	// Required when **response_type**=`search_skill`.
	Query *string `json:"query,omitempty"`

	// The type of the search query. Required when **response_type**=`search_skill`.
	QueryType *string `json:"query_type,omitempty"`

	// An optional filter that narrows the set of documents to be searched. For more information, see the [Discovery
	// service documentation]([Discovery service
	// documentation](https://cloud.ibm.com/docs/discovery?topic=discovery-query-parameters#filter).
	Filter *string `json:"filter,omitempty"`

	// The version of the Discovery service API to use for the query.
	DiscoveryVersion *string `json:"discovery_version,omitempty"`
}

// Constants associated with the DialogNodeOutputGeneric.ResponseType property.
// The type of response returned by the dialog node. The specified response type must be supported by the client
// application or channel.
//
// **Note:** The **search_skill** response type is available only for Plus and Premium users, and is used only by the v2
// runtime API.
const (
	DialogNodeOutputGeneric_ResponseType_ConnectToAgent = "connect_to_agent"
	DialogNodeOutputGeneric_ResponseType_Image          = "image"
	DialogNodeOutputGeneric_ResponseType_Option         = "option"
	DialogNodeOutputGeneric_ResponseType_Pause          = "pause"
	DialogNodeOutputGeneric_ResponseType_SearchSkill    = "search_skill"
	DialogNodeOutputGeneric_ResponseType_Text           = "text"
)

// Constants associated with the DialogNodeOutputGeneric.SelectionPolicy property.
// How a response is selected from the list, if more than one response is specified. Valid only when
// **response_type**=`text`.
const (
	DialogNodeOutputGeneric_SelectionPolicy_Multiline  = "multiline"
	DialogNodeOutputGeneric_SelectionPolicy_Random     = "random"
	DialogNodeOutputGeneric_SelectionPolicy_Sequential = "sequential"
)

// Constants associated with the DialogNodeOutputGeneric.Preference property.
// The preferred type of control to display, if supported by the channel. Valid only when **response_type**=`option`.
const (
	DialogNodeOutputGeneric_Preference_Button   = "button"
	DialogNodeOutputGeneric_Preference_Dropdown = "dropdown"
)

// Constants associated with the DialogNodeOutputGeneric.QueryType property.
// The type of the search query. Required when **response_type**=`search_skill`.
const (
	DialogNodeOutputGeneric_QueryType_DiscoveryQueryLanguage = "discovery_query_language"
	DialogNodeOutputGeneric_QueryType_NaturalLanguage        = "natural_language"
)

// NewDialogNodeOutputGeneric : Instantiate DialogNodeOutputGeneric (Generic Model Constructor)
func (assistant *AssistantV1) NewDialogNodeOutputGeneric(responseType string) (model *DialogNodeOutputGeneric, err error) {
	model = &DialogNodeOutputGeneric{
		ResponseType: core.StringPtr(responseType),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// DialogNodeOutputModifiers : Options that modify how specified output is handled.
type DialogNodeOutputModifiers struct {

	// Whether values in the output will overwrite output values in an array specified by previously executed dialog nodes.
	// If this option is set to `false`, new values will be appended to previously specified values.
	Overwrite *bool `json:"overwrite,omitempty"`
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
func (assistant *AssistantV1) NewDialogNodeOutputOptionsElement(label string, value *DialogNodeOutputOptionsElementValue) (model *DialogNodeOutputOptionsElement, err error) {
	model = &DialogNodeOutputOptionsElement{
		Label: core.StringPtr(label),
		Value: value,
	}
	err = core.ValidateStruct(model, "required parameters")
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

// DialogNodeOutputTextValuesElement : DialogNodeOutputTextValuesElement struct
type DialogNodeOutputTextValuesElement struct {

	// The text of a response. This string can include newline characters (`\n`), Markdown tagging, or other special
	// characters, if supported by the channel.
	Text *string `json:"text,omitempty"`
}

// DialogNodeVisitedDetails : DialogNodeVisitedDetails struct
type DialogNodeVisitedDetails struct {

	// A dialog node that was triggered during processing of the input message.
	DialogNode *string `json:"dialog_node,omitempty"`

	// The title of the dialog node.
	Title *string `json:"title,omitempty"`

	// The conditions that trigger the dialog node.
	Conditions *string `json:"conditions,omitempty"`
}

// DialogSuggestion : DialogSuggestion struct
type DialogSuggestion struct {

	// The user-facing label for the disambiguation option. This label is taken from the **title** or **user_label**
	// property of the corresponding dialog node, depending on the disambiguation options.
	Label *string `json:"label" validate:"required"`

	// An object defining the message input, intents, and entities to be sent to the Watson Assistant service if the user
	// selects the corresponding disambiguation option.
	Value *DialogSuggestionValue `json:"value" validate:"required"`

	// The dialog output that will be returned from the Watson Assistant service if the user selects the corresponding
	// option.
	Output *DialogSuggestionOutput `json:"output,omitempty"`

	// The ID of the dialog node that the **label** property is taken from. The **label** property is populated using the
	// value of the dialog node's **user_label** property.
	DialogNode *string `json:"dialog_node,omitempty"`
}

// NewDialogSuggestion : Instantiate DialogSuggestion (Generic Model Constructor)
func (assistant *AssistantV1) NewDialogSuggestion(label string, value *DialogSuggestionValue) (model *DialogSuggestion, err error) {
	model = &DialogSuggestion{
		Label: core.StringPtr(label),
		Value: value,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// DialogSuggestionOutput : The dialog output that will be returned from the Watson Assistant service if the user selects the corresponding
// option.
type DialogSuggestionOutput map[string]interface{}

// SetNodesVisited : Allow user to set NodesVisited
func (this *DialogSuggestionOutput) SetNodesVisited(NodesVisited *[]string) {
	(*this)["nodes_visited"] = NodesVisited
}

// GetNodesVisited : Allow user to get NodesVisited
func (this *DialogSuggestionOutput) GetNodesVisited() *[]string {
	return (*this)["nodes_visited"].(*[]string)
}

// SetNodesVisitedDetails : Allow user to set NodesVisitedDetails
func (this *DialogSuggestionOutput) SetNodesVisitedDetails(NodesVisitedDetails *[]DialogNodeVisitedDetails) {
	(*this)["nodes_visited_details"] = NodesVisitedDetails
}

// GetNodesVisitedDetails : Allow user to get NodesVisitedDetails
func (this *DialogSuggestionOutput) GetNodesVisitedDetails() *[]DialogNodeVisitedDetails {
	return (*this)["nodes_visited_details"].(*[]DialogNodeVisitedDetails)
}

// SetText : Allow user to set Text
func (this *DialogSuggestionOutput) SetText(Text *[]string) {
	(*this)["text"] = Text
}

// GetText : Allow user to get Text
func (this *DialogSuggestionOutput) GetText() *[]string {
	return (*this)["text"].(*[]string)
}

// SetGeneric : Allow user to set Generic
func (this *DialogSuggestionOutput) SetGeneric(Generic *[]DialogSuggestionResponseGeneric) {
	(*this)["generic"] = Generic
}

// GetGeneric : Allow user to get Generic
func (this *DialogSuggestionOutput) GetGeneric() *[]DialogSuggestionResponseGeneric {
	return (*this)["generic"].(*[]DialogSuggestionResponseGeneric)
}

// SetProperty : Allow user to set arbitrary property
func (this *DialogSuggestionOutput) SetProperty(Key string, Value *interface{}) {
	(*this)[Key] = Value
}

// GetProperty : Allow user to get arbitrary property
func (this *DialogSuggestionOutput) GetProperty(Key string) *interface{} {
	return (*this)[Key].(*interface{})
}

// DialogSuggestionResponseGeneric : DialogSuggestionResponseGeneric struct
type DialogSuggestionResponseGeneric struct {

	// The type of response returned by the dialog node. The specified response type must be supported by the client
	// application or channel.
	//
	// **Note:** The **suggestion** response type is part of the disambiguation feature, which is only available for Plus
	// and Premium users. The **search_skill** response type is available only for Plus and Premium users, and is used only
	// by the v2 runtime API.
	ResponseType *string `json:"response_type" validate:"required"`

	// The text of the response.
	Text *string `json:"text,omitempty"`

	// How long to pause, in milliseconds.
	Time *int64 `json:"time,omitempty"`

	// Whether to send a "user is typing" event during the pause.
	Typing *bool `json:"typing,omitempty"`

	// The URL of the image.
	Source *string `json:"source,omitempty"`

	// The title or introductory text to show before the response.
	Title *string `json:"title,omitempty"`

	// The description to show with the the response.
	Description *string `json:"description,omitempty"`

	// The preferred type of control to display.
	Preference *string `json:"preference,omitempty"`

	// An array of objects describing the options from which the user can choose.
	Options []DialogNodeOutputOptionsElement `json:"options,omitempty"`

	// A message to be sent to the human agent who will be taking over the conversation.
	MessageToHumanAgent *string `json:"message_to_human_agent,omitempty"`

	// A label identifying the topic of the conversation, derived from the **title** property of the relevant node.
	Topic *string `json:"topic,omitempty"`

	// The ID of the dialog node that the **topic** property is taken from. The **topic** property is populated using the
	// value of the dialog node's **title** property.
	DialogNode *string `json:"dialog_node,omitempty"`
}

// Constants associated with the DialogSuggestionResponseGeneric.ResponseType property.
// The type of response returned by the dialog node. The specified response type must be supported by the client
// application or channel.
//
// **Note:** The **suggestion** response type is part of the disambiguation feature, which is only available for Plus
// and Premium users. The **search_skill** response type is available only for Plus and Premium users, and is used only
// by the v2 runtime API.
const (
	DialogSuggestionResponseGeneric_ResponseType_ConnectToAgent = "connect_to_agent"
	DialogSuggestionResponseGeneric_ResponseType_Image          = "image"
	DialogSuggestionResponseGeneric_ResponseType_Option         = "option"
	DialogSuggestionResponseGeneric_ResponseType_Pause          = "pause"
	DialogSuggestionResponseGeneric_ResponseType_SearchSkill    = "search_skill"
	DialogSuggestionResponseGeneric_ResponseType_Text           = "text"
)

// Constants associated with the DialogSuggestionResponseGeneric.Preference property.
// The preferred type of control to display.
const (
	DialogSuggestionResponseGeneric_Preference_Button   = "button"
	DialogSuggestionResponseGeneric_Preference_Dropdown = "dropdown"
)

// NewDialogSuggestionResponseGeneric : Instantiate DialogSuggestionResponseGeneric (Generic Model Constructor)
func (assistant *AssistantV1) NewDialogSuggestionResponseGeneric(responseType string) (model *DialogSuggestionResponseGeneric, err error) {
	model = &DialogSuggestionResponseGeneric{
		ResponseType: core.StringPtr(responseType),
	}
	err = core.ValidateStruct(model, "required parameters")
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

// EntityCollection : An array of objects describing the entities for the workspace.
type EntityCollection struct {

	// An array of objects describing the entities defined for the workspace.
	Entities []Entity `json:"entities" validate:"required"`

	// The pagination data for the returned objects.
	Pagination *Pagination `json:"pagination" validate:"required"`
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

// EntityMentionCollection : EntityMentionCollection struct
type EntityMentionCollection struct {

	// An array of objects describing the entity mentions defined for an entity.
	Examples []EntityMention `json:"examples" validate:"required"`

	// The pagination data for the returned objects.
	Pagination *Pagination `json:"pagination" validate:"required"`
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
func (assistant *AssistantV1) NewExample(text string) (model *Example, err error) {
	model = &Example{
		Text: core.StringPtr(text),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// ExampleCollection : ExampleCollection struct
type ExampleCollection struct {

	// An array of objects describing the examples defined for the intent.
	Examples []Example `json:"examples" validate:"required"`

	// The pagination data for the returned objects.
	Pagination *Pagination `json:"pagination" validate:"required"`
}

// GetCounterexampleOptions : The GetCounterexample options.
type GetCounterexampleOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The text of a user input counterexample (for example, `What are you wearing?`).
	Text *string `json:"text" validate:"required"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetCounterexampleOptions : Instantiate GetCounterexampleOptions
func (assistant *AssistantV1) NewGetCounterexampleOptions(workspaceID string, text string) *GetCounterexampleOptions {
	return &GetCounterexampleOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Text:        core.StringPtr(text),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *GetCounterexampleOptions) SetWorkspaceID(workspaceID string) *GetCounterexampleOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetText : Allow user to set Text
func (options *GetCounterexampleOptions) SetText(text string) *GetCounterexampleOptions {
	options.Text = core.StringPtr(text)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *GetCounterexampleOptions) SetIncludeAudit(includeAudit bool) *GetCounterexampleOptions {
	options.IncludeAudit = core.BoolPtr(includeAudit)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetCounterexampleOptions) SetHeaders(param map[string]string) *GetCounterexampleOptions {
	options.Headers = param
	return options
}

// GetDialogNodeOptions : The GetDialogNode options.
type GetDialogNodeOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The dialog node ID (for example, `get_order`).
	DialogNode *string `json:"dialog_node" validate:"required"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetDialogNodeOptions : Instantiate GetDialogNodeOptions
func (assistant *AssistantV1) NewGetDialogNodeOptions(workspaceID string, dialogNode string) *GetDialogNodeOptions {
	return &GetDialogNodeOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		DialogNode:  core.StringPtr(dialogNode),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *GetDialogNodeOptions) SetWorkspaceID(workspaceID string) *GetDialogNodeOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetDialogNode : Allow user to set DialogNode
func (options *GetDialogNodeOptions) SetDialogNode(dialogNode string) *GetDialogNodeOptions {
	options.DialogNode = core.StringPtr(dialogNode)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *GetDialogNodeOptions) SetIncludeAudit(includeAudit bool) *GetDialogNodeOptions {
	options.IncludeAudit = core.BoolPtr(includeAudit)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetDialogNodeOptions) SetHeaders(param map[string]string) *GetDialogNodeOptions {
	options.Headers = param
	return options
}

// GetEntityOptions : The GetEntity options.
type GetEntityOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The name of the entity.
	Entity *string `json:"entity" validate:"required"`

	// Whether to include all element content in the returned data. If **export**=`false`, the returned data includes only
	// information about the element itself. If **export**=`true`, all content, including subelements, is included.
	Export *bool `json:"export,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetEntityOptions : Instantiate GetEntityOptions
func (assistant *AssistantV1) NewGetEntityOptions(workspaceID string, entity string) *GetEntityOptions {
	return &GetEntityOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Entity:      core.StringPtr(entity),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *GetEntityOptions) SetWorkspaceID(workspaceID string) *GetEntityOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetEntity : Allow user to set Entity
func (options *GetEntityOptions) SetEntity(entity string) *GetEntityOptions {
	options.Entity = core.StringPtr(entity)
	return options
}

// SetExport : Allow user to set Export
func (options *GetEntityOptions) SetExport(export bool) *GetEntityOptions {
	options.Export = core.BoolPtr(export)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *GetEntityOptions) SetIncludeAudit(includeAudit bool) *GetEntityOptions {
	options.IncludeAudit = core.BoolPtr(includeAudit)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetEntityOptions) SetHeaders(param map[string]string) *GetEntityOptions {
	options.Headers = param
	return options
}

// GetExampleOptions : The GetExample options.
type GetExampleOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The intent name.
	Intent *string `json:"intent" validate:"required"`

	// The text of the user input example.
	Text *string `json:"text" validate:"required"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetExampleOptions : Instantiate GetExampleOptions
func (assistant *AssistantV1) NewGetExampleOptions(workspaceID string, intent string, text string) *GetExampleOptions {
	return &GetExampleOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Intent:      core.StringPtr(intent),
		Text:        core.StringPtr(text),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *GetExampleOptions) SetWorkspaceID(workspaceID string) *GetExampleOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetIntent : Allow user to set Intent
func (options *GetExampleOptions) SetIntent(intent string) *GetExampleOptions {
	options.Intent = core.StringPtr(intent)
	return options
}

// SetText : Allow user to set Text
func (options *GetExampleOptions) SetText(text string) *GetExampleOptions {
	options.Text = core.StringPtr(text)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *GetExampleOptions) SetIncludeAudit(includeAudit bool) *GetExampleOptions {
	options.IncludeAudit = core.BoolPtr(includeAudit)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetExampleOptions) SetHeaders(param map[string]string) *GetExampleOptions {
	options.Headers = param
	return options
}

// GetIntentOptions : The GetIntent options.
type GetIntentOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The intent name.
	Intent *string `json:"intent" validate:"required"`

	// Whether to include all element content in the returned data. If **export**=`false`, the returned data includes only
	// information about the element itself. If **export**=`true`, all content, including subelements, is included.
	Export *bool `json:"export,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetIntentOptions : Instantiate GetIntentOptions
func (assistant *AssistantV1) NewGetIntentOptions(workspaceID string, intent string) *GetIntentOptions {
	return &GetIntentOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Intent:      core.StringPtr(intent),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *GetIntentOptions) SetWorkspaceID(workspaceID string) *GetIntentOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetIntent : Allow user to set Intent
func (options *GetIntentOptions) SetIntent(intent string) *GetIntentOptions {
	options.Intent = core.StringPtr(intent)
	return options
}

// SetExport : Allow user to set Export
func (options *GetIntentOptions) SetExport(export bool) *GetIntentOptions {
	options.Export = core.BoolPtr(export)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *GetIntentOptions) SetIncludeAudit(includeAudit bool) *GetIntentOptions {
	options.IncludeAudit = core.BoolPtr(includeAudit)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetIntentOptions) SetHeaders(param map[string]string) *GetIntentOptions {
	options.Headers = param
	return options
}

// GetSynonymOptions : The GetSynonym options.
type GetSynonymOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The name of the entity.
	Entity *string `json:"entity" validate:"required"`

	// The text of the entity value.
	Value *string `json:"value" validate:"required"`

	// The text of the synonym.
	Synonym *string `json:"synonym" validate:"required"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetSynonymOptions : Instantiate GetSynonymOptions
func (assistant *AssistantV1) NewGetSynonymOptions(workspaceID string, entity string, value string, synonym string) *GetSynonymOptions {
	return &GetSynonymOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Entity:      core.StringPtr(entity),
		Value:       core.StringPtr(value),
		Synonym:     core.StringPtr(synonym),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *GetSynonymOptions) SetWorkspaceID(workspaceID string) *GetSynonymOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetEntity : Allow user to set Entity
func (options *GetSynonymOptions) SetEntity(entity string) *GetSynonymOptions {
	options.Entity = core.StringPtr(entity)
	return options
}

// SetValue : Allow user to set Value
func (options *GetSynonymOptions) SetValue(value string) *GetSynonymOptions {
	options.Value = core.StringPtr(value)
	return options
}

// SetSynonym : Allow user to set Synonym
func (options *GetSynonymOptions) SetSynonym(synonym string) *GetSynonymOptions {
	options.Synonym = core.StringPtr(synonym)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *GetSynonymOptions) SetIncludeAudit(includeAudit bool) *GetSynonymOptions {
	options.IncludeAudit = core.BoolPtr(includeAudit)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetSynonymOptions) SetHeaders(param map[string]string) *GetSynonymOptions {
	options.Headers = param
	return options
}

// GetValueOptions : The GetValue options.
type GetValueOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The name of the entity.
	Entity *string `json:"entity" validate:"required"`

	// The text of the entity value.
	Value *string `json:"value" validate:"required"`

	// Whether to include all element content in the returned data. If **export**=`false`, the returned data includes only
	// information about the element itself. If **export**=`true`, all content, including subelements, is included.
	Export *bool `json:"export,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetValueOptions : Instantiate GetValueOptions
func (assistant *AssistantV1) NewGetValueOptions(workspaceID string, entity string, value string) *GetValueOptions {
	return &GetValueOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Entity:      core.StringPtr(entity),
		Value:       core.StringPtr(value),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *GetValueOptions) SetWorkspaceID(workspaceID string) *GetValueOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetEntity : Allow user to set Entity
func (options *GetValueOptions) SetEntity(entity string) *GetValueOptions {
	options.Entity = core.StringPtr(entity)
	return options
}

// SetValue : Allow user to set Value
func (options *GetValueOptions) SetValue(value string) *GetValueOptions {
	options.Value = core.StringPtr(value)
	return options
}

// SetExport : Allow user to set Export
func (options *GetValueOptions) SetExport(export bool) *GetValueOptions {
	options.Export = core.BoolPtr(export)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *GetValueOptions) SetIncludeAudit(includeAudit bool) *GetValueOptions {
	options.IncludeAudit = core.BoolPtr(includeAudit)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetValueOptions) SetHeaders(param map[string]string) *GetValueOptions {
	options.Headers = param
	return options
}

// GetWorkspaceOptions : The GetWorkspace options.
type GetWorkspaceOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// Whether to include all element content in the returned data. If **export**=`false`, the returned data includes only
	// information about the element itself. If **export**=`true`, all content, including subelements, is included.
	Export *bool `json:"export,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Indicates how the returned workspace data will be sorted. This parameter is valid only if **export**=`true`. Specify
	// `sort=stable` to sort all workspace objects by unique identifier, in ascending alphabetical order.
	Sort *string `json:"sort,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the GetWorkspaceOptions.Sort property.
// Indicates how the returned workspace data will be sorted. This parameter is valid only if **export**=`true`. Specify
// `sort=stable` to sort all workspace objects by unique identifier, in ascending alphabetical order.
const (
	GetWorkspaceOptions_Sort_Stable = "stable"
)

// NewGetWorkspaceOptions : Instantiate GetWorkspaceOptions
func (assistant *AssistantV1) NewGetWorkspaceOptions(workspaceID string) *GetWorkspaceOptions {
	return &GetWorkspaceOptions{
		WorkspaceID: core.StringPtr(workspaceID),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *GetWorkspaceOptions) SetWorkspaceID(workspaceID string) *GetWorkspaceOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetExport : Allow user to set Export
func (options *GetWorkspaceOptions) SetExport(export bool) *GetWorkspaceOptions {
	options.Export = core.BoolPtr(export)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *GetWorkspaceOptions) SetIncludeAudit(includeAudit bool) *GetWorkspaceOptions {
	options.IncludeAudit = core.BoolPtr(includeAudit)
	return options
}

// SetSort : Allow user to set Sort
func (options *GetWorkspaceOptions) SetSort(sort string) *GetWorkspaceOptions {
	options.Sort = core.StringPtr(sort)
	return options
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

// IntentCollection : IntentCollection struct
type IntentCollection struct {

	// An array of objects describing the intents defined for the workspace.
	Intents []Intent `json:"intents" validate:"required"`

	// The pagination data for the returned objects.
	Pagination *Pagination `json:"pagination" validate:"required"`
}

// ListAllLogsOptions : The ListAllLogs options.
type ListAllLogsOptions struct {

	// A cacheable parameter that limits the results to those matching the specified filter. You must specify a filter
	// query that includes a value for `language`, as well as a value for `request.context.system.assistant_id`,
	// `workspace_id`, or `request.context.metadata.deployment`. For more information, see the
	// [documentation](https://cloud.ibm.com/docs/assistant?topic=assistant-filter-reference#filter-reference).
	Filter *string `json:"filter" validate:"required"`

	// How to sort the returned log events. You can sort by **request_timestamp**. To reverse the sort order, prefix the
	// parameter value with a minus sign (`-`).
	Sort *string `json:"sort,omitempty"`

	// The number of records to return in each page of results.
	PageLimit *int64 `json:"page_limit,omitempty"`

	// A token identifying the page of results to retrieve.
	Cursor *string `json:"cursor,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListAllLogsOptions : Instantiate ListAllLogsOptions
func (assistant *AssistantV1) NewListAllLogsOptions(filter string) *ListAllLogsOptions {
	return &ListAllLogsOptions{
		Filter: core.StringPtr(filter),
	}
}

// SetFilter : Allow user to set Filter
func (options *ListAllLogsOptions) SetFilter(filter string) *ListAllLogsOptions {
	options.Filter = core.StringPtr(filter)
	return options
}

// SetSort : Allow user to set Sort
func (options *ListAllLogsOptions) SetSort(sort string) *ListAllLogsOptions {
	options.Sort = core.StringPtr(sort)
	return options
}

// SetPageLimit : Allow user to set PageLimit
func (options *ListAllLogsOptions) SetPageLimit(pageLimit int64) *ListAllLogsOptions {
	options.PageLimit = core.Int64Ptr(pageLimit)
	return options
}

// SetCursor : Allow user to set Cursor
func (options *ListAllLogsOptions) SetCursor(cursor string) *ListAllLogsOptions {
	options.Cursor = core.StringPtr(cursor)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListAllLogsOptions) SetHeaders(param map[string]string) *ListAllLogsOptions {
	options.Headers = param
	return options
}

// ListCounterexamplesOptions : The ListCounterexamples options.
type ListCounterexamplesOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The number of records to return in each page of results.
	PageLimit *int64 `json:"page_limit,omitempty"`

	// The attribute by which returned counterexamples will be sorted. To reverse the sort order, prefix the value with a
	// minus sign (`-`).
	Sort *string `json:"sort,omitempty"`

	// A token identifying the page of results to retrieve.
	Cursor *string `json:"cursor,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the ListCounterexamplesOptions.Sort property.
// The attribute by which returned counterexamples will be sorted. To reverse the sort order, prefix the value with a
// minus sign (`-`).
const (
	ListCounterexamplesOptions_Sort_Text    = "text"
	ListCounterexamplesOptions_Sort_Updated = "updated"
)

// NewListCounterexamplesOptions : Instantiate ListCounterexamplesOptions
func (assistant *AssistantV1) NewListCounterexamplesOptions(workspaceID string) *ListCounterexamplesOptions {
	return &ListCounterexamplesOptions{
		WorkspaceID: core.StringPtr(workspaceID),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *ListCounterexamplesOptions) SetWorkspaceID(workspaceID string) *ListCounterexamplesOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetPageLimit : Allow user to set PageLimit
func (options *ListCounterexamplesOptions) SetPageLimit(pageLimit int64) *ListCounterexamplesOptions {
	options.PageLimit = core.Int64Ptr(pageLimit)
	return options
}

// SetSort : Allow user to set Sort
func (options *ListCounterexamplesOptions) SetSort(sort string) *ListCounterexamplesOptions {
	options.Sort = core.StringPtr(sort)
	return options
}

// SetCursor : Allow user to set Cursor
func (options *ListCounterexamplesOptions) SetCursor(cursor string) *ListCounterexamplesOptions {
	options.Cursor = core.StringPtr(cursor)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *ListCounterexamplesOptions) SetIncludeAudit(includeAudit bool) *ListCounterexamplesOptions {
	options.IncludeAudit = core.BoolPtr(includeAudit)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListCounterexamplesOptions) SetHeaders(param map[string]string) *ListCounterexamplesOptions {
	options.Headers = param
	return options
}

// ListDialogNodesOptions : The ListDialogNodes options.
type ListDialogNodesOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The number of records to return in each page of results.
	PageLimit *int64 `json:"page_limit,omitempty"`

	// The attribute by which returned dialog nodes will be sorted. To reverse the sort order, prefix the value with a
	// minus sign (`-`).
	Sort *string `json:"sort,omitempty"`

	// A token identifying the page of results to retrieve.
	Cursor *string `json:"cursor,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the ListDialogNodesOptions.Sort property.
// The attribute by which returned dialog nodes will be sorted. To reverse the sort order, prefix the value with a minus
// sign (`-`).
const (
	ListDialogNodesOptions_Sort_DialogNode = "dialog_node"
	ListDialogNodesOptions_Sort_Updated    = "updated"
)

// NewListDialogNodesOptions : Instantiate ListDialogNodesOptions
func (assistant *AssistantV1) NewListDialogNodesOptions(workspaceID string) *ListDialogNodesOptions {
	return &ListDialogNodesOptions{
		WorkspaceID: core.StringPtr(workspaceID),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *ListDialogNodesOptions) SetWorkspaceID(workspaceID string) *ListDialogNodesOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetPageLimit : Allow user to set PageLimit
func (options *ListDialogNodesOptions) SetPageLimit(pageLimit int64) *ListDialogNodesOptions {
	options.PageLimit = core.Int64Ptr(pageLimit)
	return options
}

// SetSort : Allow user to set Sort
func (options *ListDialogNodesOptions) SetSort(sort string) *ListDialogNodesOptions {
	options.Sort = core.StringPtr(sort)
	return options
}

// SetCursor : Allow user to set Cursor
func (options *ListDialogNodesOptions) SetCursor(cursor string) *ListDialogNodesOptions {
	options.Cursor = core.StringPtr(cursor)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *ListDialogNodesOptions) SetIncludeAudit(includeAudit bool) *ListDialogNodesOptions {
	options.IncludeAudit = core.BoolPtr(includeAudit)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListDialogNodesOptions) SetHeaders(param map[string]string) *ListDialogNodesOptions {
	options.Headers = param
	return options
}

// ListEntitiesOptions : The ListEntities options.
type ListEntitiesOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// Whether to include all element content in the returned data. If **export**=`false`, the returned data includes only
	// information about the element itself. If **export**=`true`, all content, including subelements, is included.
	Export *bool `json:"export,omitempty"`

	// The number of records to return in each page of results.
	PageLimit *int64 `json:"page_limit,omitempty"`

	// The attribute by which returned entities will be sorted. To reverse the sort order, prefix the value with a minus
	// sign (`-`).
	Sort *string `json:"sort,omitempty"`

	// A token identifying the page of results to retrieve.
	Cursor *string `json:"cursor,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the ListEntitiesOptions.Sort property.
// The attribute by which returned entities will be sorted. To reverse the sort order, prefix the value with a minus
// sign (`-`).
const (
	ListEntitiesOptions_Sort_Entity  = "entity"
	ListEntitiesOptions_Sort_Updated = "updated"
)

// NewListEntitiesOptions : Instantiate ListEntitiesOptions
func (assistant *AssistantV1) NewListEntitiesOptions(workspaceID string) *ListEntitiesOptions {
	return &ListEntitiesOptions{
		WorkspaceID: core.StringPtr(workspaceID),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *ListEntitiesOptions) SetWorkspaceID(workspaceID string) *ListEntitiesOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetExport : Allow user to set Export
func (options *ListEntitiesOptions) SetExport(export bool) *ListEntitiesOptions {
	options.Export = core.BoolPtr(export)
	return options
}

// SetPageLimit : Allow user to set PageLimit
func (options *ListEntitiesOptions) SetPageLimit(pageLimit int64) *ListEntitiesOptions {
	options.PageLimit = core.Int64Ptr(pageLimit)
	return options
}

// SetSort : Allow user to set Sort
func (options *ListEntitiesOptions) SetSort(sort string) *ListEntitiesOptions {
	options.Sort = core.StringPtr(sort)
	return options
}

// SetCursor : Allow user to set Cursor
func (options *ListEntitiesOptions) SetCursor(cursor string) *ListEntitiesOptions {
	options.Cursor = core.StringPtr(cursor)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *ListEntitiesOptions) SetIncludeAudit(includeAudit bool) *ListEntitiesOptions {
	options.IncludeAudit = core.BoolPtr(includeAudit)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListEntitiesOptions) SetHeaders(param map[string]string) *ListEntitiesOptions {
	options.Headers = param
	return options
}

// ListExamplesOptions : The ListExamples options.
type ListExamplesOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The intent name.
	Intent *string `json:"intent" validate:"required"`

	// The number of records to return in each page of results.
	PageLimit *int64 `json:"page_limit,omitempty"`

	// The attribute by which returned examples will be sorted. To reverse the sort order, prefix the value with a minus
	// sign (`-`).
	Sort *string `json:"sort,omitempty"`

	// A token identifying the page of results to retrieve.
	Cursor *string `json:"cursor,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the ListExamplesOptions.Sort property.
// The attribute by which returned examples will be sorted. To reverse the sort order, prefix the value with a minus
// sign (`-`).
const (
	ListExamplesOptions_Sort_Text    = "text"
	ListExamplesOptions_Sort_Updated = "updated"
)

// NewListExamplesOptions : Instantiate ListExamplesOptions
func (assistant *AssistantV1) NewListExamplesOptions(workspaceID string, intent string) *ListExamplesOptions {
	return &ListExamplesOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Intent:      core.StringPtr(intent),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *ListExamplesOptions) SetWorkspaceID(workspaceID string) *ListExamplesOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetIntent : Allow user to set Intent
func (options *ListExamplesOptions) SetIntent(intent string) *ListExamplesOptions {
	options.Intent = core.StringPtr(intent)
	return options
}

// SetPageLimit : Allow user to set PageLimit
func (options *ListExamplesOptions) SetPageLimit(pageLimit int64) *ListExamplesOptions {
	options.PageLimit = core.Int64Ptr(pageLimit)
	return options
}

// SetSort : Allow user to set Sort
func (options *ListExamplesOptions) SetSort(sort string) *ListExamplesOptions {
	options.Sort = core.StringPtr(sort)
	return options
}

// SetCursor : Allow user to set Cursor
func (options *ListExamplesOptions) SetCursor(cursor string) *ListExamplesOptions {
	options.Cursor = core.StringPtr(cursor)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *ListExamplesOptions) SetIncludeAudit(includeAudit bool) *ListExamplesOptions {
	options.IncludeAudit = core.BoolPtr(includeAudit)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListExamplesOptions) SetHeaders(param map[string]string) *ListExamplesOptions {
	options.Headers = param
	return options
}

// ListIntentsOptions : The ListIntents options.
type ListIntentsOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// Whether to include all element content in the returned data. If **export**=`false`, the returned data includes only
	// information about the element itself. If **export**=`true`, all content, including subelements, is included.
	Export *bool `json:"export,omitempty"`

	// The number of records to return in each page of results.
	PageLimit *int64 `json:"page_limit,omitempty"`

	// The attribute by which returned intents will be sorted. To reverse the sort order, prefix the value with a minus
	// sign (`-`).
	Sort *string `json:"sort,omitempty"`

	// A token identifying the page of results to retrieve.
	Cursor *string `json:"cursor,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the ListIntentsOptions.Sort property.
// The attribute by which returned intents will be sorted. To reverse the sort order, prefix the value with a minus sign
// (`-`).
const (
	ListIntentsOptions_Sort_Intent  = "intent"
	ListIntentsOptions_Sort_Updated = "updated"
)

// NewListIntentsOptions : Instantiate ListIntentsOptions
func (assistant *AssistantV1) NewListIntentsOptions(workspaceID string) *ListIntentsOptions {
	return &ListIntentsOptions{
		WorkspaceID: core.StringPtr(workspaceID),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *ListIntentsOptions) SetWorkspaceID(workspaceID string) *ListIntentsOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetExport : Allow user to set Export
func (options *ListIntentsOptions) SetExport(export bool) *ListIntentsOptions {
	options.Export = core.BoolPtr(export)
	return options
}

// SetPageLimit : Allow user to set PageLimit
func (options *ListIntentsOptions) SetPageLimit(pageLimit int64) *ListIntentsOptions {
	options.PageLimit = core.Int64Ptr(pageLimit)
	return options
}

// SetSort : Allow user to set Sort
func (options *ListIntentsOptions) SetSort(sort string) *ListIntentsOptions {
	options.Sort = core.StringPtr(sort)
	return options
}

// SetCursor : Allow user to set Cursor
func (options *ListIntentsOptions) SetCursor(cursor string) *ListIntentsOptions {
	options.Cursor = core.StringPtr(cursor)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *ListIntentsOptions) SetIncludeAudit(includeAudit bool) *ListIntentsOptions {
	options.IncludeAudit = core.BoolPtr(includeAudit)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListIntentsOptions) SetHeaders(param map[string]string) *ListIntentsOptions {
	options.Headers = param
	return options
}

// ListLogsOptions : The ListLogs options.
type ListLogsOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// How to sort the returned log events. You can sort by **request_timestamp**. To reverse the sort order, prefix the
	// parameter value with a minus sign (`-`).
	Sort *string `json:"sort,omitempty"`

	// A cacheable parameter that limits the results to those matching the specified filter. For more information, see the
	// [documentation](https://cloud.ibm.com/docs/assistant?topic=assistant-filter-reference#filter-reference).
	Filter *string `json:"filter,omitempty"`

	// The number of records to return in each page of results.
	PageLimit *int64 `json:"page_limit,omitempty"`

	// A token identifying the page of results to retrieve.
	Cursor *string `json:"cursor,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListLogsOptions : Instantiate ListLogsOptions
func (assistant *AssistantV1) NewListLogsOptions(workspaceID string) *ListLogsOptions {
	return &ListLogsOptions{
		WorkspaceID: core.StringPtr(workspaceID),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *ListLogsOptions) SetWorkspaceID(workspaceID string) *ListLogsOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetSort : Allow user to set Sort
func (options *ListLogsOptions) SetSort(sort string) *ListLogsOptions {
	options.Sort = core.StringPtr(sort)
	return options
}

// SetFilter : Allow user to set Filter
func (options *ListLogsOptions) SetFilter(filter string) *ListLogsOptions {
	options.Filter = core.StringPtr(filter)
	return options
}

// SetPageLimit : Allow user to set PageLimit
func (options *ListLogsOptions) SetPageLimit(pageLimit int64) *ListLogsOptions {
	options.PageLimit = core.Int64Ptr(pageLimit)
	return options
}

// SetCursor : Allow user to set Cursor
func (options *ListLogsOptions) SetCursor(cursor string) *ListLogsOptions {
	options.Cursor = core.StringPtr(cursor)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListLogsOptions) SetHeaders(param map[string]string) *ListLogsOptions {
	options.Headers = param
	return options
}

// ListMentionsOptions : The ListMentions options.
type ListMentionsOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The name of the entity.
	Entity *string `json:"entity" validate:"required"`

	// Whether to include all element content in the returned data. If **export**=`false`, the returned data includes only
	// information about the element itself. If **export**=`true`, all content, including subelements, is included.
	Export *bool `json:"export,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListMentionsOptions : Instantiate ListMentionsOptions
func (assistant *AssistantV1) NewListMentionsOptions(workspaceID string, entity string) *ListMentionsOptions {
	return &ListMentionsOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Entity:      core.StringPtr(entity),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *ListMentionsOptions) SetWorkspaceID(workspaceID string) *ListMentionsOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetEntity : Allow user to set Entity
func (options *ListMentionsOptions) SetEntity(entity string) *ListMentionsOptions {
	options.Entity = core.StringPtr(entity)
	return options
}

// SetExport : Allow user to set Export
func (options *ListMentionsOptions) SetExport(export bool) *ListMentionsOptions {
	options.Export = core.BoolPtr(export)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *ListMentionsOptions) SetIncludeAudit(includeAudit bool) *ListMentionsOptions {
	options.IncludeAudit = core.BoolPtr(includeAudit)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListMentionsOptions) SetHeaders(param map[string]string) *ListMentionsOptions {
	options.Headers = param
	return options
}

// ListSynonymsOptions : The ListSynonyms options.
type ListSynonymsOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The name of the entity.
	Entity *string `json:"entity" validate:"required"`

	// The text of the entity value.
	Value *string `json:"value" validate:"required"`

	// The number of records to return in each page of results.
	PageLimit *int64 `json:"page_limit,omitempty"`

	// The attribute by which returned entity value synonyms will be sorted. To reverse the sort order, prefix the value
	// with a minus sign (`-`).
	Sort *string `json:"sort,omitempty"`

	// A token identifying the page of results to retrieve.
	Cursor *string `json:"cursor,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the ListSynonymsOptions.Sort property.
// The attribute by which returned entity value synonyms will be sorted. To reverse the sort order, prefix the value
// with a minus sign (`-`).
const (
	ListSynonymsOptions_Sort_Synonym = "synonym"
	ListSynonymsOptions_Sort_Updated = "updated"
)

// NewListSynonymsOptions : Instantiate ListSynonymsOptions
func (assistant *AssistantV1) NewListSynonymsOptions(workspaceID string, entity string, value string) *ListSynonymsOptions {
	return &ListSynonymsOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Entity:      core.StringPtr(entity),
		Value:       core.StringPtr(value),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *ListSynonymsOptions) SetWorkspaceID(workspaceID string) *ListSynonymsOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetEntity : Allow user to set Entity
func (options *ListSynonymsOptions) SetEntity(entity string) *ListSynonymsOptions {
	options.Entity = core.StringPtr(entity)
	return options
}

// SetValue : Allow user to set Value
func (options *ListSynonymsOptions) SetValue(value string) *ListSynonymsOptions {
	options.Value = core.StringPtr(value)
	return options
}

// SetPageLimit : Allow user to set PageLimit
func (options *ListSynonymsOptions) SetPageLimit(pageLimit int64) *ListSynonymsOptions {
	options.PageLimit = core.Int64Ptr(pageLimit)
	return options
}

// SetSort : Allow user to set Sort
func (options *ListSynonymsOptions) SetSort(sort string) *ListSynonymsOptions {
	options.Sort = core.StringPtr(sort)
	return options
}

// SetCursor : Allow user to set Cursor
func (options *ListSynonymsOptions) SetCursor(cursor string) *ListSynonymsOptions {
	options.Cursor = core.StringPtr(cursor)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *ListSynonymsOptions) SetIncludeAudit(includeAudit bool) *ListSynonymsOptions {
	options.IncludeAudit = core.BoolPtr(includeAudit)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListSynonymsOptions) SetHeaders(param map[string]string) *ListSynonymsOptions {
	options.Headers = param
	return options
}

// ListValuesOptions : The ListValues options.
type ListValuesOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The name of the entity.
	Entity *string `json:"entity" validate:"required"`

	// Whether to include all element content in the returned data. If **export**=`false`, the returned data includes only
	// information about the element itself. If **export**=`true`, all content, including subelements, is included.
	Export *bool `json:"export,omitempty"`

	// The number of records to return in each page of results.
	PageLimit *int64 `json:"page_limit,omitempty"`

	// The attribute by which returned entity values will be sorted. To reverse the sort order, prefix the value with a
	// minus sign (`-`).
	Sort *string `json:"sort,omitempty"`

	// A token identifying the page of results to retrieve.
	Cursor *string `json:"cursor,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the ListValuesOptions.Sort property.
// The attribute by which returned entity values will be sorted. To reverse the sort order, prefix the value with a
// minus sign (`-`).
const (
	ListValuesOptions_Sort_Updated = "updated"
	ListValuesOptions_Sort_Value   = "value"
)

// NewListValuesOptions : Instantiate ListValuesOptions
func (assistant *AssistantV1) NewListValuesOptions(workspaceID string, entity string) *ListValuesOptions {
	return &ListValuesOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Entity:      core.StringPtr(entity),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *ListValuesOptions) SetWorkspaceID(workspaceID string) *ListValuesOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetEntity : Allow user to set Entity
func (options *ListValuesOptions) SetEntity(entity string) *ListValuesOptions {
	options.Entity = core.StringPtr(entity)
	return options
}

// SetExport : Allow user to set Export
func (options *ListValuesOptions) SetExport(export bool) *ListValuesOptions {
	options.Export = core.BoolPtr(export)
	return options
}

// SetPageLimit : Allow user to set PageLimit
func (options *ListValuesOptions) SetPageLimit(pageLimit int64) *ListValuesOptions {
	options.PageLimit = core.Int64Ptr(pageLimit)
	return options
}

// SetSort : Allow user to set Sort
func (options *ListValuesOptions) SetSort(sort string) *ListValuesOptions {
	options.Sort = core.StringPtr(sort)
	return options
}

// SetCursor : Allow user to set Cursor
func (options *ListValuesOptions) SetCursor(cursor string) *ListValuesOptions {
	options.Cursor = core.StringPtr(cursor)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *ListValuesOptions) SetIncludeAudit(includeAudit bool) *ListValuesOptions {
	options.IncludeAudit = core.BoolPtr(includeAudit)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListValuesOptions) SetHeaders(param map[string]string) *ListValuesOptions {
	options.Headers = param
	return options
}

// ListWorkspacesOptions : The ListWorkspaces options.
type ListWorkspacesOptions struct {

	// The number of records to return in each page of results.
	PageLimit *int64 `json:"page_limit,omitempty"`

	// The attribute by which returned workspaces will be sorted. To reverse the sort order, prefix the value with a minus
	// sign (`-`).
	Sort *string `json:"sort,omitempty"`

	// A token identifying the page of results to retrieve.
	Cursor *string `json:"cursor,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the ListWorkspacesOptions.Sort property.
// The attribute by which returned workspaces will be sorted. To reverse the sort order, prefix the value with a minus
// sign (`-`).
const (
	ListWorkspacesOptions_Sort_Name    = "name"
	ListWorkspacesOptions_Sort_Updated = "updated"
)

// NewListWorkspacesOptions : Instantiate ListWorkspacesOptions
func (assistant *AssistantV1) NewListWorkspacesOptions() *ListWorkspacesOptions {
	return &ListWorkspacesOptions{}
}

// SetPageLimit : Allow user to set PageLimit
func (options *ListWorkspacesOptions) SetPageLimit(pageLimit int64) *ListWorkspacesOptions {
	options.PageLimit = core.Int64Ptr(pageLimit)
	return options
}

// SetSort : Allow user to set Sort
func (options *ListWorkspacesOptions) SetSort(sort string) *ListWorkspacesOptions {
	options.Sort = core.StringPtr(sort)
	return options
}

// SetCursor : Allow user to set Cursor
func (options *ListWorkspacesOptions) SetCursor(cursor string) *ListWorkspacesOptions {
	options.Cursor = core.StringPtr(cursor)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *ListWorkspacesOptions) SetIncludeAudit(includeAudit bool) *ListWorkspacesOptions {
	options.IncludeAudit = core.BoolPtr(includeAudit)
	return options
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

// LogCollection : LogCollection struct
type LogCollection struct {

	// An array of objects describing log events.
	Logs []Log `json:"logs" validate:"required"`

	// The pagination data for the returned objects.
	Pagination *LogPagination `json:"pagination" validate:"required"`
}

// LogMessage : Log message details.
type LogMessage struct {

	// The severity of the log message.
	Level *string `json:"level" validate:"required"`

	// The text of the log message.
	Msg *string `json:"msg" validate:"required"`
}

// Constants associated with the LogMessage.Level property.
// The severity of the log message.
const (
	LogMessage_Level_Error = "error"
	LogMessage_Level_Info  = "info"
	LogMessage_Level_Warn  = "warn"
)

// NewLogMessage : Instantiate LogMessage (Generic Model Constructor)
func (assistant *AssistantV1) NewLogMessage(level string, msg string) (model *LogMessage, err error) {
	model = &LogMessage{
		Level: core.StringPtr(level),
		Msg:   core.StringPtr(msg),
	}
	err = core.ValidateStruct(model, "required parameters")
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

// Mention : A mention of a contextual entity.
type Mention struct {

	// The name of the entity.
	Entity *string `json:"entity" validate:"required"`

	// An array of zero-based character offsets that indicate where the entity mentions begin and end in the input text.
	Location []int64 `json:"location" validate:"required"`
}

// NewMention : Instantiate Mention (Generic Model Constructor)
func (assistant *AssistantV1) NewMention(entity string, location []int64) (model *Mention, err error) {
	model = &Mention{
		Entity:   core.StringPtr(entity),
		Location: location,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// MessageContextMetadata : Metadata related to the message.
type MessageContextMetadata struct {

	// A label identifying the deployment environment, used for filtering log data. This string cannot contain carriage
	// return, newline, or tab characters.
	Deployment *string `json:"deployment,omitempty"`

	// A string value that identifies the user who is interacting with the workspace. The client must provide a unique
	// identifier for each individual end user who accesses the application. For Plus and Premium plans, this user ID is
	// used to identify unique users for billing purposes. This string cannot contain carriage return, newline, or tab
	// characters.
	UserID *string `json:"user_id,omitempty"`
}

// MessageInput : An input object that includes the input text.
type MessageInput map[string]interface{}

// SetText : Allow user to set Text
func (this *MessageInput) SetText(Text *string) {
	(*this)["text"] = Text
}

// GetText : Allow user to get Text
func (this *MessageInput) GetText() *string {
	return (*this)["text"].(*string)
}

// SetSpellingSuggestions : Allow user to set SpellingSuggestions
func (this *MessageInput) SetSpellingSuggestions(SpellingSuggestions *bool) {
	(*this)["spelling_suggestions"] = SpellingSuggestions
}

// GetSpellingSuggestions : Allow user to get SpellingSuggestions
func (this *MessageInput) GetSpellingSuggestions() *bool {
	return (*this)["spelling_suggestions"].(*bool)
}

// SetSpellingAutoCorrect : Allow user to set SpellingAutoCorrect
func (this *MessageInput) SetSpellingAutoCorrect(SpellingAutoCorrect *bool) {
	(*this)["spelling_auto_correct"] = SpellingAutoCorrect
}

// GetSpellingAutoCorrect : Allow user to get SpellingAutoCorrect
func (this *MessageInput) GetSpellingAutoCorrect() *bool {
	return (*this)["spelling_auto_correct"].(*bool)
}

// SetSuggestedText : Allow user to set SuggestedText
func (this *MessageInput) SetSuggestedText(SuggestedText *string) {
	(*this)["suggested_text"] = SuggestedText
}

// GetSuggestedText : Allow user to get SuggestedText
func (this *MessageInput) GetSuggestedText() *string {
	return (*this)["suggested_text"].(*string)
}

// SetOriginalText : Allow user to set OriginalText
func (this *MessageInput) SetOriginalText(OriginalText *string) {
	(*this)["original_text"] = OriginalText
}

// GetOriginalText : Allow user to get OriginalText
func (this *MessageInput) GetOriginalText() *string {
	return (*this)["original_text"].(*string)
}

// SetProperty : Allow user to set arbitrary property
func (this *MessageInput) SetProperty(Key string, Value *interface{}) {
	(*this)[Key] = Value
}

// GetProperty : Allow user to get arbitrary property
func (this *MessageInput) GetProperty(Key string) *interface{} {
	return (*this)[Key].(*interface{})
}

// MessageOptions : The Message options.
type MessageOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

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

	// Whether to include additional diagnostic information about the dialog nodes that were visited during processing of
	// the message.
	NodesVisitedDetails *bool `json:"nodes_visited_details,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewMessageOptions : Instantiate MessageOptions
func (assistant *AssistantV1) NewMessageOptions(workspaceID string) *MessageOptions {
	return &MessageOptions{
		WorkspaceID: core.StringPtr(workspaceID),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *MessageOptions) SetWorkspaceID(workspaceID string) *MessageOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetInput : Allow user to set Input
func (options *MessageOptions) SetInput(input *MessageInput) *MessageOptions {
	options.Input = input
	return options
}

// SetIntents : Allow user to set Intents
func (options *MessageOptions) SetIntents(intents []RuntimeIntent) *MessageOptions {
	options.Intents = intents
	return options
}

// SetEntities : Allow user to set Entities
func (options *MessageOptions) SetEntities(entities []RuntimeEntity) *MessageOptions {
	options.Entities = entities
	return options
}

// SetAlternateIntents : Allow user to set AlternateIntents
func (options *MessageOptions) SetAlternateIntents(alternateIntents bool) *MessageOptions {
	options.AlternateIntents = core.BoolPtr(alternateIntents)
	return options
}

// SetContext : Allow user to set Context
func (options *MessageOptions) SetContext(context *Context) *MessageOptions {
	options.Context = context
	return options
}

// SetOutput : Allow user to set Output
func (options *MessageOptions) SetOutput(output *OutputData) *MessageOptions {
	options.Output = output
	return options
}

// SetNodesVisitedDetails : Allow user to set NodesVisitedDetails
func (options *MessageOptions) SetNodesVisitedDetails(nodesVisitedDetails bool) *MessageOptions {
	options.NodesVisitedDetails = core.BoolPtr(nodesVisitedDetails)
	return options
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
}

// OutputData : An output object that includes the response to the user, the dialog nodes that were triggered, and messages from the
// log.
type OutputData map[string]interface{}

// SetNodesVisited : Allow user to set NodesVisited
func (this *OutputData) SetNodesVisited(NodesVisited *[]string) {
	(*this)["nodes_visited"] = NodesVisited
}

// GetNodesVisited : Allow user to get NodesVisited
func (this *OutputData) GetNodesVisited() *[]string {
	return (*this)["nodes_visited"].(*[]string)
}

// SetNodesVisitedDetails : Allow user to set NodesVisitedDetails
func (this *OutputData) SetNodesVisitedDetails(NodesVisitedDetails *[]DialogNodeVisitedDetails) {
	(*this)["nodes_visited_details"] = NodesVisitedDetails
}

// GetNodesVisitedDetails : Allow user to get NodesVisitedDetails
func (this *OutputData) GetNodesVisitedDetails() *[]DialogNodeVisitedDetails {
	return (*this)["nodes_visited_details"].(*[]DialogNodeVisitedDetails)
}

// SetLogMessages : Allow user to set LogMessages
func (this *OutputData) SetLogMessages(LogMessages *[]LogMessage) {
	(*this)["log_messages"] = LogMessages
}

// GetLogMessages : Allow user to get LogMessages
func (this *OutputData) GetLogMessages() *[]LogMessage {
	return (*this)["log_messages"].(*[]LogMessage)
}

// SetText : Allow user to set Text
func (this *OutputData) SetText(Text *[]string) {
	(*this)["text"] = Text
}

// GetText : Allow user to get Text
func (this *OutputData) GetText() *[]string {
	return (*this)["text"].(*[]string)
}

// SetGeneric : Allow user to set Generic
func (this *OutputData) SetGeneric(Generic *[]RuntimeResponseGeneric) {
	(*this)["generic"] = Generic
}

// GetGeneric : Allow user to get Generic
func (this *OutputData) GetGeneric() *[]RuntimeResponseGeneric {
	return (*this)["generic"].(*[]RuntimeResponseGeneric)
}

// SetProperty : Allow user to set arbitrary property
func (this *OutputData) SetProperty(Key string, Value *interface{}) {
	(*this)[Key] = Value
}

// GetProperty : Allow user to get arbitrary property
func (this *OutputData) GetProperty(Key string) *interface{} {
	return (*this)[Key].(*interface{})
}

// Pagination : The pagination data for the returned objects.
type Pagination struct {

	// The URL that will return the same page of results.
	RefreshURL *string `json:"refresh_url" validate:"required"`

	// The URL that will return the next page of results.
	NextURL *string `json:"next_url,omitempty"`

	// Reserved for future use.
	Total *int64 `json:"total,omitempty"`

	// Reserved for future use.
	Matched *int64 `json:"matched,omitempty"`

	// A token identifying the current page of results.
	RefreshCursor *string `json:"refresh_cursor,omitempty"`

	// A token identifying the next page of results.
	NextCursor *string `json:"next_cursor,omitempty"`
}

// RuntimeEntity : A term from the request that was identified as an entity.
type RuntimeEntity struct {

	// An entity detected in the input.
	Entity *string `json:"entity" validate:"required"`

	// An array of zero-based character offsets that indicate where the detected entity values begin and end in the input
	// text.
	Location []int64 `json:"location" validate:"required"`

	// The entity value that was recognized in the user input.
	Value *string `json:"value" validate:"required"`

	// A decimal percentage that represents Watson's confidence in the recognized entity.
	Confidence *float64 `json:"confidence,omitempty"`

	// Any metadata for the entity.
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// The recognized capture groups for the entity, as defined by the entity pattern.
	Groups []CaptureGroup `json:"groups,omitempty"`

	// An object containing detailed information about the entity recognized in the user input. This property is included
	// only if the new system entities are enabled for the workspace.
	//
	// For more information about how the new system entities are interpreted, see the
	// [documentation](https://cloud.ibm.com/docs/assistant?topic=assistant-beta-system-entities).
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
func (assistant *AssistantV1) NewRuntimeEntity(entity string, location []int64, value string) (model *RuntimeEntity, err error) {
	model = &RuntimeEntity{
		Entity:   core.StringPtr(entity),
		Location: location,
		Value:    core.StringPtr(value),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// RuntimeEntityAlternative : An alternative value for the recognized entity.
type RuntimeEntityAlternative struct {

	// The entity value that was recognized in the user input.
	Value *string `json:"value,omitempty"`

	// A decimal percentage that represents Watson's confidence in the recognized entity.
	Confidence *float64 `json:"confidence,omitempty"`
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
	RuntimeEntityInterpretation_Granularity_Day       = "day"
	RuntimeEntityInterpretation_Granularity_Fortnight = "fortnight"
	RuntimeEntityInterpretation_Granularity_Hour      = "hour"
	RuntimeEntityInterpretation_Granularity_Instant   = "instant"
	RuntimeEntityInterpretation_Granularity_Minute    = "minute"
	RuntimeEntityInterpretation_Granularity_Month     = "month"
	RuntimeEntityInterpretation_Granularity_Quarter   = "quarter"
	RuntimeEntityInterpretation_Granularity_Second    = "second"
	RuntimeEntityInterpretation_Granularity_Week      = "week"
	RuntimeEntityInterpretation_Granularity_Weekend   = "weekend"
	RuntimeEntityInterpretation_Granularity_Year      = "year"
)

// RuntimeEntityRole : An object describing the role played by a system entity that is specifies the beginning or end of a range recognized
// in the user input. This property is included only if the new system entities are enabled for the workspace.
type RuntimeEntityRole struct {

	// The relationship of the entity to the range.
	Type *string `json:"type,omitempty"`
}

// Constants associated with the RuntimeEntityRole.Type property.
// The relationship of the entity to the range.
const (
	RuntimeEntityRole_Type_DateFrom   = "date_from"
	RuntimeEntityRole_Type_DateTo     = "date_to"
	RuntimeEntityRole_Type_NumberFrom = "number_from"
	RuntimeEntityRole_Type_NumberTo   = "number_to"
	RuntimeEntityRole_Type_TimeFrom   = "time_from"
	RuntimeEntityRole_Type_TimeTo     = "time_to"
)

// RuntimeIntent : An intent identified in the user input.
type RuntimeIntent struct {

	// The name of the recognized intent.
	Intent *string `json:"intent" validate:"required"`

	// A decimal percentage that represents Watson's confidence in the intent.
	Confidence *float64 `json:"confidence" validate:"required"`
}

// NewRuntimeIntent : Instantiate RuntimeIntent (Generic Model Constructor)
func (assistant *AssistantV1) NewRuntimeIntent(intent string, confidence float64) (model *RuntimeIntent, err error) {
	model = &RuntimeIntent{
		Intent:     core.StringPtr(intent),
		Confidence: core.Float64Ptr(confidence),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// RuntimeResponseGeneric : RuntimeResponseGeneric struct
type RuntimeResponseGeneric struct {

	// The type of response returned by the dialog node. The specified response type must be supported by the client
	// application or channel.
	//
	// **Note:** The **suggestion** response type is part of the disambiguation feature, which is only available for Plus
	// and Premium users.
	ResponseType *string `json:"response_type" validate:"required"`

	// The text of the response.
	Text *string `json:"text,omitempty"`

	// How long to pause, in milliseconds.
	Time *int64 `json:"time,omitempty"`

	// Whether to send a "user is typing" event during the pause.
	Typing *bool `json:"typing,omitempty"`

	// The URL of the image.
	Source *string `json:"source,omitempty"`

	// The title or introductory text to show before the response.
	Title *string `json:"title,omitempty"`

	// The description to show with the the response.
	Description *string `json:"description,omitempty"`

	// The preferred type of control to display.
	Preference *string `json:"preference,omitempty"`

	// An array of objects describing the options from which the user can choose.
	Options []DialogNodeOutputOptionsElement `json:"options,omitempty"`

	// A message to be sent to the human agent who will be taking over the conversation.
	MessageToHumanAgent *string `json:"message_to_human_agent,omitempty"`

	// A label identifying the topic of the conversation, derived from the **title** property of the relevant node.
	Topic *string `json:"topic,omitempty"`

	// The ID of the dialog node that the **topic** property is taken from. The **topic** property is populated using the
	// value of the dialog node's **title** property.
	DialogNode *string `json:"dialog_node,omitempty"`

	// An array of objects describing the possible matching dialog nodes from which the user can choose.
	//
	// **Note:** The **suggestions** property is part of the disambiguation feature, which is only available for Plus and
	// Premium users.
	Suggestions []DialogSuggestion `json:"suggestions,omitempty"`
}

// Constants associated with the RuntimeResponseGeneric.ResponseType property.
// The type of response returned by the dialog node. The specified response type must be supported by the client
// application or channel.
//
// **Note:** The **suggestion** response type is part of the disambiguation feature, which is only available for Plus
// and Premium users.
const (
	RuntimeResponseGeneric_ResponseType_ConnectToAgent = "connect_to_agent"
	RuntimeResponseGeneric_ResponseType_Image          = "image"
	RuntimeResponseGeneric_ResponseType_Option         = "option"
	RuntimeResponseGeneric_ResponseType_Pause          = "pause"
	RuntimeResponseGeneric_ResponseType_Suggestion     = "suggestion"
	RuntimeResponseGeneric_ResponseType_Text           = "text"
)

// Constants associated with the RuntimeResponseGeneric.Preference property.
// The preferred type of control to display.
const (
	RuntimeResponseGeneric_Preference_Button   = "button"
	RuntimeResponseGeneric_Preference_Dropdown = "dropdown"
)

// NewRuntimeResponseGeneric : Instantiate RuntimeResponseGeneric (Generic Model Constructor)
func (assistant *AssistantV1) NewRuntimeResponseGeneric(responseType string) (model *RuntimeResponseGeneric, err error) {
	model = &RuntimeResponseGeneric{
		ResponseType: core.StringPtr(responseType),
	}
	err = core.ValidateStruct(model, "required parameters")
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
func (assistant *AssistantV1) NewSynonym(synonym string) (model *Synonym, err error) {
	model = &Synonym{
		Synonym: core.StringPtr(synonym),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// SynonymCollection : SynonymCollection struct
type SynonymCollection struct {

	// An array of synonyms.
	Synonyms []Synonym `json:"synonyms" validate:"required"`

	// The pagination data for the returned objects.
	Pagination *Pagination `json:"pagination" validate:"required"`
}

// SystemResponse : For internal use only.
type SystemResponse map[string]interface{}

// SetProperty : Allow user to set arbitrary property
func (this *SystemResponse) SetProperty(Key string, Value *interface{}) {
	(*this)[Key] = Value
}

// GetProperty : Allow user to get arbitrary property
func (this *SystemResponse) GetProperty(Key string) *interface{} {
	return (*this)[Key].(*interface{})
}

// UpdateCounterexampleOptions : The UpdateCounterexample options.
type UpdateCounterexampleOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The text of a user input counterexample (for example, `What are you wearing?`).
	Text *string `json:"text" validate:"required"`

	// The text of a user input marked as irrelevant input. This string must conform to the following restrictions:
	// - It cannot contain carriage return, newline, or tab characters.
	// - It cannot consist of only whitespace characters.
	NewText *string `json:"new_text,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewUpdateCounterexampleOptions : Instantiate UpdateCounterexampleOptions
func (assistant *AssistantV1) NewUpdateCounterexampleOptions(workspaceID string, text string) *UpdateCounterexampleOptions {
	return &UpdateCounterexampleOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Text:        core.StringPtr(text),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *UpdateCounterexampleOptions) SetWorkspaceID(workspaceID string) *UpdateCounterexampleOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetText : Allow user to set Text
func (options *UpdateCounterexampleOptions) SetText(text string) *UpdateCounterexampleOptions {
	options.Text = core.StringPtr(text)
	return options
}

// SetNewText : Allow user to set NewText
func (options *UpdateCounterexampleOptions) SetNewText(newText string) *UpdateCounterexampleOptions {
	options.NewText = core.StringPtr(newText)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *UpdateCounterexampleOptions) SetIncludeAudit(includeAudit bool) *UpdateCounterexampleOptions {
	options.IncludeAudit = core.BoolPtr(includeAudit)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateCounterexampleOptions) SetHeaders(param map[string]string) *UpdateCounterexampleOptions {
	options.Headers = param
	return options
}

// UpdateDialogNodeOptions : The UpdateDialogNode options.
type UpdateDialogNodeOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The dialog node ID (for example, `get_order`).
	DialogNode *string `json:"dialog_node" validate:"required"`

	// The dialog node ID. This string must conform to the following restrictions:
	// - It can contain only Unicode alphanumeric, space, underscore, hyphen, and dot characters.
	NewDialogNode *string `json:"new_dialog_node,omitempty"`

	// The description of the dialog node. This string cannot contain carriage return, newline, or tab characters.
	NewDescription *string `json:"new_description,omitempty"`

	// The condition that will trigger the dialog node. This string cannot contain carriage return, newline, or tab
	// characters.
	NewConditions *string `json:"new_conditions,omitempty"`

	// The ID of the parent dialog node. This property is omitted if the dialog node has no parent.
	NewParent *string `json:"new_parent,omitempty"`

	// The ID of the previous sibling dialog node. This property is omitted if the dialog node has no previous sibling.
	NewPreviousSibling *string `json:"new_previous_sibling,omitempty"`

	// The output of the dialog node. For more information about how to specify dialog node output, see the
	// [documentation](https://cloud.ibm.com/docs/assistant?topic=assistant-dialog-overview#dialog-overview-responses).
	NewOutput *DialogNodeOutput `json:"new_output,omitempty"`

	// The context for the dialog node.
	NewContext map[string]interface{} `json:"new_context,omitempty"`

	// The metadata for the dialog node.
	NewMetadata map[string]interface{} `json:"new_metadata,omitempty"`

	// The next step to execute following this dialog node.
	NewNextStep *DialogNodeNextStep `json:"new_next_step,omitempty"`

	// The alias used to identify the dialog node. This string must conform to the following restrictions:
	// - It can contain only Unicode alphanumeric, space, underscore, hyphen, and dot characters.
	NewTitle *string `json:"new_title,omitempty"`

	// How the dialog node is processed.
	NewType *string `json:"new_type,omitempty"`

	// How an `event_handler` node is processed.
	NewEventName *string `json:"new_event_name,omitempty"`

	// The location in the dialog context where output is stored.
	NewVariable *string `json:"new_variable,omitempty"`

	// An array of objects describing any actions to be invoked by the dialog node.
	NewActions []DialogNodeAction `json:"new_actions,omitempty"`

	// Whether this top-level dialog node can be digressed into.
	NewDigressIn *string `json:"new_digress_in,omitempty"`

	// Whether this dialog node can be returned to after a digression.
	NewDigressOut *string `json:"new_digress_out,omitempty"`

	// Whether the user can digress to top-level nodes while filling out slots.
	NewDigressOutSlots *string `json:"new_digress_out_slots,omitempty"`

	// A label that can be displayed externally to describe the purpose of the node to users.
	NewUserLabel *string `json:"new_user_label,omitempty"`

	// Whether the dialog node should be excluded from disambiguation suggestions. Valid only when **type**=`standard` or
	// `frame`.
	NewDisambiguationOptOut *bool `json:"new_disambiguation_opt_out,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the UpdateDialogNodeOptions.NewType property.
// How the dialog node is processed.
const (
	UpdateDialogNodeOptions_NewType_EventHandler      = "event_handler"
	UpdateDialogNodeOptions_NewType_Folder            = "folder"
	UpdateDialogNodeOptions_NewType_Frame             = "frame"
	UpdateDialogNodeOptions_NewType_ResponseCondition = "response_condition"
	UpdateDialogNodeOptions_NewType_Slot              = "slot"
	UpdateDialogNodeOptions_NewType_Standard          = "standard"
)

// Constants associated with the UpdateDialogNodeOptions.NewEventName property.
// How an `event_handler` node is processed.
const (
	UpdateDialogNodeOptions_NewEventName_DigressionReturnPrompt   = "digression_return_prompt"
	UpdateDialogNodeOptions_NewEventName_Filled                   = "filled"
	UpdateDialogNodeOptions_NewEventName_FilledMultiple           = "filled_multiple"
	UpdateDialogNodeOptions_NewEventName_Focus                    = "focus"
	UpdateDialogNodeOptions_NewEventName_Generic                  = "generic"
	UpdateDialogNodeOptions_NewEventName_Input                    = "input"
	UpdateDialogNodeOptions_NewEventName_Nomatch                  = "nomatch"
	UpdateDialogNodeOptions_NewEventName_NomatchResponsesDepleted = "nomatch_responses_depleted"
	UpdateDialogNodeOptions_NewEventName_Validate                 = "validate"
)

// Constants associated with the UpdateDialogNodeOptions.NewDigressIn property.
// Whether this top-level dialog node can be digressed into.
const (
	UpdateDialogNodeOptions_NewDigressIn_DoesNotReturn = "does_not_return"
	UpdateDialogNodeOptions_NewDigressIn_NotAvailable  = "not_available"
	UpdateDialogNodeOptions_NewDigressIn_Returns       = "returns"
)

// Constants associated with the UpdateDialogNodeOptions.NewDigressOut property.
// Whether this dialog node can be returned to after a digression.
const (
	UpdateDialogNodeOptions_NewDigressOut_AllowAll            = "allow_all"
	UpdateDialogNodeOptions_NewDigressOut_AllowAllNeverReturn = "allow_all_never_return"
	UpdateDialogNodeOptions_NewDigressOut_AllowReturning      = "allow_returning"
)

// Constants associated with the UpdateDialogNodeOptions.NewDigressOutSlots property.
// Whether the user can digress to top-level nodes while filling out slots.
const (
	UpdateDialogNodeOptions_NewDigressOutSlots_AllowAll       = "allow_all"
	UpdateDialogNodeOptions_NewDigressOutSlots_AllowReturning = "allow_returning"
	UpdateDialogNodeOptions_NewDigressOutSlots_NotAllowed     = "not_allowed"
)

// NewUpdateDialogNodeOptions : Instantiate UpdateDialogNodeOptions
func (assistant *AssistantV1) NewUpdateDialogNodeOptions(workspaceID string, dialogNode string) *UpdateDialogNodeOptions {
	return &UpdateDialogNodeOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		DialogNode:  core.StringPtr(dialogNode),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *UpdateDialogNodeOptions) SetWorkspaceID(workspaceID string) *UpdateDialogNodeOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetDialogNode : Allow user to set DialogNode
func (options *UpdateDialogNodeOptions) SetDialogNode(dialogNode string) *UpdateDialogNodeOptions {
	options.DialogNode = core.StringPtr(dialogNode)
	return options
}

// SetNewDialogNode : Allow user to set NewDialogNode
func (options *UpdateDialogNodeOptions) SetNewDialogNode(newDialogNode string) *UpdateDialogNodeOptions {
	options.NewDialogNode = core.StringPtr(newDialogNode)
	return options
}

// SetNewDescription : Allow user to set NewDescription
func (options *UpdateDialogNodeOptions) SetNewDescription(newDescription string) *UpdateDialogNodeOptions {
	options.NewDescription = core.StringPtr(newDescription)
	return options
}

// SetNewConditions : Allow user to set NewConditions
func (options *UpdateDialogNodeOptions) SetNewConditions(newConditions string) *UpdateDialogNodeOptions {
	options.NewConditions = core.StringPtr(newConditions)
	return options
}

// SetNewParent : Allow user to set NewParent
func (options *UpdateDialogNodeOptions) SetNewParent(newParent string) *UpdateDialogNodeOptions {
	options.NewParent = core.StringPtr(newParent)
	return options
}

// SetNewPreviousSibling : Allow user to set NewPreviousSibling
func (options *UpdateDialogNodeOptions) SetNewPreviousSibling(newPreviousSibling string) *UpdateDialogNodeOptions {
	options.NewPreviousSibling = core.StringPtr(newPreviousSibling)
	return options
}

// SetNewOutput : Allow user to set NewOutput
func (options *UpdateDialogNodeOptions) SetNewOutput(newOutput *DialogNodeOutput) *UpdateDialogNodeOptions {
	options.NewOutput = newOutput
	return options
}

// SetNewContext : Allow user to set NewContext
func (options *UpdateDialogNodeOptions) SetNewContext(newContext map[string]interface{}) *UpdateDialogNodeOptions {
	options.NewContext = newContext
	return options
}

// SetNewMetadata : Allow user to set NewMetadata
func (options *UpdateDialogNodeOptions) SetNewMetadata(newMetadata map[string]interface{}) *UpdateDialogNodeOptions {
	options.NewMetadata = newMetadata
	return options
}

// SetNewNextStep : Allow user to set NewNextStep
func (options *UpdateDialogNodeOptions) SetNewNextStep(newNextStep *DialogNodeNextStep) *UpdateDialogNodeOptions {
	options.NewNextStep = newNextStep
	return options
}

// SetNewTitle : Allow user to set NewTitle
func (options *UpdateDialogNodeOptions) SetNewTitle(newTitle string) *UpdateDialogNodeOptions {
	options.NewTitle = core.StringPtr(newTitle)
	return options
}

// SetNewType : Allow user to set NewType
func (options *UpdateDialogNodeOptions) SetNewType(newType string) *UpdateDialogNodeOptions {
	options.NewType = core.StringPtr(newType)
	return options
}

// SetNewEventName : Allow user to set NewEventName
func (options *UpdateDialogNodeOptions) SetNewEventName(newEventName string) *UpdateDialogNodeOptions {
	options.NewEventName = core.StringPtr(newEventName)
	return options
}

// SetNewVariable : Allow user to set NewVariable
func (options *UpdateDialogNodeOptions) SetNewVariable(newVariable string) *UpdateDialogNodeOptions {
	options.NewVariable = core.StringPtr(newVariable)
	return options
}

// SetNewActions : Allow user to set NewActions
func (options *UpdateDialogNodeOptions) SetNewActions(newActions []DialogNodeAction) *UpdateDialogNodeOptions {
	options.NewActions = newActions
	return options
}

// SetNewDigressIn : Allow user to set NewDigressIn
func (options *UpdateDialogNodeOptions) SetNewDigressIn(newDigressIn string) *UpdateDialogNodeOptions {
	options.NewDigressIn = core.StringPtr(newDigressIn)
	return options
}

// SetNewDigressOut : Allow user to set NewDigressOut
func (options *UpdateDialogNodeOptions) SetNewDigressOut(newDigressOut string) *UpdateDialogNodeOptions {
	options.NewDigressOut = core.StringPtr(newDigressOut)
	return options
}

// SetNewDigressOutSlots : Allow user to set NewDigressOutSlots
func (options *UpdateDialogNodeOptions) SetNewDigressOutSlots(newDigressOutSlots string) *UpdateDialogNodeOptions {
	options.NewDigressOutSlots = core.StringPtr(newDigressOutSlots)
	return options
}

// SetNewUserLabel : Allow user to set NewUserLabel
func (options *UpdateDialogNodeOptions) SetNewUserLabel(newUserLabel string) *UpdateDialogNodeOptions {
	options.NewUserLabel = core.StringPtr(newUserLabel)
	return options
}

// SetNewDisambiguationOptOut : Allow user to set NewDisambiguationOptOut
func (options *UpdateDialogNodeOptions) SetNewDisambiguationOptOut(newDisambiguationOptOut bool) *UpdateDialogNodeOptions {
	options.NewDisambiguationOptOut = core.BoolPtr(newDisambiguationOptOut)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *UpdateDialogNodeOptions) SetIncludeAudit(includeAudit bool) *UpdateDialogNodeOptions {
	options.IncludeAudit = core.BoolPtr(includeAudit)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateDialogNodeOptions) SetHeaders(param map[string]string) *UpdateDialogNodeOptions {
	options.Headers = param
	return options
}

// UpdateEntityOptions : The UpdateEntity options.
type UpdateEntityOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The name of the entity.
	Entity *string `json:"entity" validate:"required"`

	// The name of the entity. This string must conform to the following restrictions:
	// - It can contain only Unicode alphanumeric, underscore, and hyphen characters.
	// - It cannot begin with the reserved prefix `sys-`.
	NewEntity *string `json:"new_entity,omitempty"`

	// The description of the entity. This string cannot contain carriage return, newline, or tab characters.
	NewDescription *string `json:"new_description,omitempty"`

	// Any metadata related to the entity.
	NewMetadata map[string]interface{} `json:"new_metadata,omitempty"`

	// Whether to use fuzzy matching for the entity.
	NewFuzzyMatch *bool `json:"new_fuzzy_match,omitempty"`

	// An array of objects describing the entity values.
	NewValues []CreateValue `json:"new_values,omitempty"`

	// Whether the new data is to be appended to the existing data in the entity. If **append**=`false`, elements included
	// in the new data completely replace the corresponding existing elements, including all subelements. For example, if
	// the new data for the entity includes **values** and **append**=`false`, all existing values for the entity are
	// discarded and replaced with the new values.
	//
	// If **append**=`true`, existing elements are preserved, and the new elements are added. If any elements in the new
	// data collide with existing elements, the update request fails.
	Append *bool `json:"append,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewUpdateEntityOptions : Instantiate UpdateEntityOptions
func (assistant *AssistantV1) NewUpdateEntityOptions(workspaceID string, entity string) *UpdateEntityOptions {
	return &UpdateEntityOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Entity:      core.StringPtr(entity),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *UpdateEntityOptions) SetWorkspaceID(workspaceID string) *UpdateEntityOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetEntity : Allow user to set Entity
func (options *UpdateEntityOptions) SetEntity(entity string) *UpdateEntityOptions {
	options.Entity = core.StringPtr(entity)
	return options
}

// SetNewEntity : Allow user to set NewEntity
func (options *UpdateEntityOptions) SetNewEntity(newEntity string) *UpdateEntityOptions {
	options.NewEntity = core.StringPtr(newEntity)
	return options
}

// SetNewDescription : Allow user to set NewDescription
func (options *UpdateEntityOptions) SetNewDescription(newDescription string) *UpdateEntityOptions {
	options.NewDescription = core.StringPtr(newDescription)
	return options
}

// SetNewMetadata : Allow user to set NewMetadata
func (options *UpdateEntityOptions) SetNewMetadata(newMetadata map[string]interface{}) *UpdateEntityOptions {
	options.NewMetadata = newMetadata
	return options
}

// SetNewFuzzyMatch : Allow user to set NewFuzzyMatch
func (options *UpdateEntityOptions) SetNewFuzzyMatch(newFuzzyMatch bool) *UpdateEntityOptions {
	options.NewFuzzyMatch = core.BoolPtr(newFuzzyMatch)
	return options
}

// SetNewValues : Allow user to set NewValues
func (options *UpdateEntityOptions) SetNewValues(newValues []CreateValue) *UpdateEntityOptions {
	options.NewValues = newValues
	return options
}

// SetAppend : Allow user to set Append
func (options *UpdateEntityOptions) SetAppend(append bool) *UpdateEntityOptions {
	options.Append = core.BoolPtr(append)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *UpdateEntityOptions) SetIncludeAudit(includeAudit bool) *UpdateEntityOptions {
	options.IncludeAudit = core.BoolPtr(includeAudit)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateEntityOptions) SetHeaders(param map[string]string) *UpdateEntityOptions {
	options.Headers = param
	return options
}

// UpdateExampleOptions : The UpdateExample options.
type UpdateExampleOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The intent name.
	Intent *string `json:"intent" validate:"required"`

	// The text of the user input example.
	Text *string `json:"text" validate:"required"`

	// The text of the user input example. This string must conform to the following restrictions:
	// - It cannot contain carriage return, newline, or tab characters.
	// - It cannot consist of only whitespace characters.
	NewText *string `json:"new_text,omitempty"`

	// An array of contextual entity mentions.
	NewMentions []Mention `json:"new_mentions,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewUpdateExampleOptions : Instantiate UpdateExampleOptions
func (assistant *AssistantV1) NewUpdateExampleOptions(workspaceID string, intent string, text string) *UpdateExampleOptions {
	return &UpdateExampleOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Intent:      core.StringPtr(intent),
		Text:        core.StringPtr(text),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *UpdateExampleOptions) SetWorkspaceID(workspaceID string) *UpdateExampleOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetIntent : Allow user to set Intent
func (options *UpdateExampleOptions) SetIntent(intent string) *UpdateExampleOptions {
	options.Intent = core.StringPtr(intent)
	return options
}

// SetText : Allow user to set Text
func (options *UpdateExampleOptions) SetText(text string) *UpdateExampleOptions {
	options.Text = core.StringPtr(text)
	return options
}

// SetNewText : Allow user to set NewText
func (options *UpdateExampleOptions) SetNewText(newText string) *UpdateExampleOptions {
	options.NewText = core.StringPtr(newText)
	return options
}

// SetNewMentions : Allow user to set NewMentions
func (options *UpdateExampleOptions) SetNewMentions(newMentions []Mention) *UpdateExampleOptions {
	options.NewMentions = newMentions
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *UpdateExampleOptions) SetIncludeAudit(includeAudit bool) *UpdateExampleOptions {
	options.IncludeAudit = core.BoolPtr(includeAudit)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateExampleOptions) SetHeaders(param map[string]string) *UpdateExampleOptions {
	options.Headers = param
	return options
}

// UpdateIntentOptions : The UpdateIntent options.
type UpdateIntentOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The intent name.
	Intent *string `json:"intent" validate:"required"`

	// The name of the intent. This string must conform to the following restrictions:
	// - It can contain only Unicode alphanumeric, underscore, hyphen, and dot characters.
	// - It cannot begin with the reserved prefix `sys-`.
	NewIntent *string `json:"new_intent,omitempty"`

	// The description of the intent. This string cannot contain carriage return, newline, or tab characters.
	NewDescription *string `json:"new_description,omitempty"`

	// An array of user input examples for the intent.
	NewExamples []Example `json:"new_examples,omitempty"`

	// Whether the new data is to be appended to the existing data in the object. If **append**=`false`, elements included
	// in the new data completely replace the corresponding existing elements, including all subelements. For example, if
	// the new data for the intent includes **examples** and **append**=`false`, all existing examples for the intent are
	// discarded and replaced with the new examples.
	//
	// If **append**=`true`, existing elements are preserved, and the new elements are added. If any elements in the new
	// data collide with existing elements, the update request fails.
	Append *bool `json:"append,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewUpdateIntentOptions : Instantiate UpdateIntentOptions
func (assistant *AssistantV1) NewUpdateIntentOptions(workspaceID string, intent string) *UpdateIntentOptions {
	return &UpdateIntentOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Intent:      core.StringPtr(intent),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *UpdateIntentOptions) SetWorkspaceID(workspaceID string) *UpdateIntentOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetIntent : Allow user to set Intent
func (options *UpdateIntentOptions) SetIntent(intent string) *UpdateIntentOptions {
	options.Intent = core.StringPtr(intent)
	return options
}

// SetNewIntent : Allow user to set NewIntent
func (options *UpdateIntentOptions) SetNewIntent(newIntent string) *UpdateIntentOptions {
	options.NewIntent = core.StringPtr(newIntent)
	return options
}

// SetNewDescription : Allow user to set NewDescription
func (options *UpdateIntentOptions) SetNewDescription(newDescription string) *UpdateIntentOptions {
	options.NewDescription = core.StringPtr(newDescription)
	return options
}

// SetNewExamples : Allow user to set NewExamples
func (options *UpdateIntentOptions) SetNewExamples(newExamples []Example) *UpdateIntentOptions {
	options.NewExamples = newExamples
	return options
}

// SetAppend : Allow user to set Append
func (options *UpdateIntentOptions) SetAppend(append bool) *UpdateIntentOptions {
	options.Append = core.BoolPtr(append)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *UpdateIntentOptions) SetIncludeAudit(includeAudit bool) *UpdateIntentOptions {
	options.IncludeAudit = core.BoolPtr(includeAudit)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateIntentOptions) SetHeaders(param map[string]string) *UpdateIntentOptions {
	options.Headers = param
	return options
}

// UpdateSynonymOptions : The UpdateSynonym options.
type UpdateSynonymOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The name of the entity.
	Entity *string `json:"entity" validate:"required"`

	// The text of the entity value.
	Value *string `json:"value" validate:"required"`

	// The text of the synonym.
	Synonym *string `json:"synonym" validate:"required"`

	// The text of the synonym. This string must conform to the following restrictions:
	// - It cannot contain carriage return, newline, or tab characters.
	// - It cannot consist of only whitespace characters.
	NewSynonym *string `json:"new_synonym,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewUpdateSynonymOptions : Instantiate UpdateSynonymOptions
func (assistant *AssistantV1) NewUpdateSynonymOptions(workspaceID string, entity string, value string, synonym string) *UpdateSynonymOptions {
	return &UpdateSynonymOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Entity:      core.StringPtr(entity),
		Value:       core.StringPtr(value),
		Synonym:     core.StringPtr(synonym),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *UpdateSynonymOptions) SetWorkspaceID(workspaceID string) *UpdateSynonymOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetEntity : Allow user to set Entity
func (options *UpdateSynonymOptions) SetEntity(entity string) *UpdateSynonymOptions {
	options.Entity = core.StringPtr(entity)
	return options
}

// SetValue : Allow user to set Value
func (options *UpdateSynonymOptions) SetValue(value string) *UpdateSynonymOptions {
	options.Value = core.StringPtr(value)
	return options
}

// SetSynonym : Allow user to set Synonym
func (options *UpdateSynonymOptions) SetSynonym(synonym string) *UpdateSynonymOptions {
	options.Synonym = core.StringPtr(synonym)
	return options
}

// SetNewSynonym : Allow user to set NewSynonym
func (options *UpdateSynonymOptions) SetNewSynonym(newSynonym string) *UpdateSynonymOptions {
	options.NewSynonym = core.StringPtr(newSynonym)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *UpdateSynonymOptions) SetIncludeAudit(includeAudit bool) *UpdateSynonymOptions {
	options.IncludeAudit = core.BoolPtr(includeAudit)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateSynonymOptions) SetHeaders(param map[string]string) *UpdateSynonymOptions {
	options.Headers = param
	return options
}

// UpdateValueOptions : The UpdateValue options.
type UpdateValueOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The name of the entity.
	Entity *string `json:"entity" validate:"required"`

	// The text of the entity value.
	Value *string `json:"value" validate:"required"`

	// The text of the entity value. This string must conform to the following restrictions:
	// - It cannot contain carriage return, newline, or tab characters.
	// - It cannot consist of only whitespace characters.
	NewValue *string `json:"new_value,omitempty"`

	// Any metadata related to the entity value.
	NewMetadata map[string]interface{} `json:"new_metadata,omitempty"`

	// Specifies the type of entity value.
	NewType *string `json:"new_type,omitempty"`

	// An array of synonyms for the entity value. A value can specify either synonyms or patterns (depending on the value
	// type), but not both. A synonym must conform to the following resrictions:
	// - It cannot contain carriage return, newline, or tab characters.
	// - It cannot consist of only whitespace characters.
	NewSynonyms []string `json:"new_synonyms,omitempty"`

	// An array of patterns for the entity value. A value can specify either synonyms or patterns (depending on the value
	// type), but not both. A pattern is a regular expression; for more information about how to specify a pattern, see the
	// [documentation](https://cloud.ibm.com/docs/assistant?topic=assistant-entities#entities-create-dictionary-based).
	NewPatterns []string `json:"new_patterns,omitempty"`

	// Whether the new data is to be appended to the existing data in the entity value. If **append**=`false`, elements
	// included in the new data completely replace the corresponding existing elements, including all subelements. For
	// example, if the new data for the entity value includes **synonyms** and **append**=`false`, all existing synonyms
	// for the entity value are discarded and replaced with the new synonyms.
	//
	// If **append**=`true`, existing elements are preserved, and the new elements are added. If any elements in the new
	// data collide with existing elements, the update request fails.
	Append *bool `json:"append,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// Constants associated with the UpdateValueOptions.NewType property.
// Specifies the type of entity value.
const (
	UpdateValueOptions_NewType_Patterns = "patterns"
	UpdateValueOptions_NewType_Synonyms = "synonyms"
)

// NewUpdateValueOptions : Instantiate UpdateValueOptions
func (assistant *AssistantV1) NewUpdateValueOptions(workspaceID string, entity string, value string) *UpdateValueOptions {
	return &UpdateValueOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Entity:      core.StringPtr(entity),
		Value:       core.StringPtr(value),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *UpdateValueOptions) SetWorkspaceID(workspaceID string) *UpdateValueOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetEntity : Allow user to set Entity
func (options *UpdateValueOptions) SetEntity(entity string) *UpdateValueOptions {
	options.Entity = core.StringPtr(entity)
	return options
}

// SetValue : Allow user to set Value
func (options *UpdateValueOptions) SetValue(value string) *UpdateValueOptions {
	options.Value = core.StringPtr(value)
	return options
}

// SetNewValue : Allow user to set NewValue
func (options *UpdateValueOptions) SetNewValue(newValue string) *UpdateValueOptions {
	options.NewValue = core.StringPtr(newValue)
	return options
}

// SetNewMetadata : Allow user to set NewMetadata
func (options *UpdateValueOptions) SetNewMetadata(newMetadata map[string]interface{}) *UpdateValueOptions {
	options.NewMetadata = newMetadata
	return options
}

// SetNewType : Allow user to set NewType
func (options *UpdateValueOptions) SetNewType(newType string) *UpdateValueOptions {
	options.NewType = core.StringPtr(newType)
	return options
}

// SetNewSynonyms : Allow user to set NewSynonyms
func (options *UpdateValueOptions) SetNewSynonyms(newSynonyms []string) *UpdateValueOptions {
	options.NewSynonyms = newSynonyms
	return options
}

// SetNewPatterns : Allow user to set NewPatterns
func (options *UpdateValueOptions) SetNewPatterns(newPatterns []string) *UpdateValueOptions {
	options.NewPatterns = newPatterns
	return options
}

// SetAppend : Allow user to set Append
func (options *UpdateValueOptions) SetAppend(append bool) *UpdateValueOptions {
	options.Append = core.BoolPtr(append)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *UpdateValueOptions) SetIncludeAudit(includeAudit bool) *UpdateValueOptions {
	options.IncludeAudit = core.BoolPtr(includeAudit)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateValueOptions) SetHeaders(param map[string]string) *UpdateValueOptions {
	options.Headers = param
	return options
}

// UpdateWorkspaceOptions : The UpdateWorkspace options.
type UpdateWorkspaceOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The name of the workspace. This string cannot contain carriage return, newline, or tab characters.
	Name *string `json:"name,omitempty"`

	// The description of the workspace. This string cannot contain carriage return, newline, or tab characters.
	Description *string `json:"description,omitempty"`

	// The language of the workspace.
	Language *string `json:"language,omitempty"`

	// Any metadata related to the workspace.
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// Whether training data from the workspace (including artifacts such as intents and entities) can be used by IBM for
	// general service improvements. `true` indicates that workspace training data is not to be used.
	LearningOptOut *bool `json:"learning_opt_out,omitempty"`

	// Global settings for the workspace.
	SystemSettings *WorkspaceSystemSettings `json:"system_settings,omitempty"`

	// An array of objects defining the intents for the workspace.
	Intents []CreateIntent `json:"intents,omitempty"`

	// An array of objects describing the entities for the workspace.
	Entities []CreateEntity `json:"entities,omitempty"`

	// An array of objects describing the dialog nodes in the workspace.
	DialogNodes []DialogNode `json:"dialog_nodes,omitempty"`

	// An array of objects defining input examples that have been marked as irrelevant input.
	Counterexamples []Counterexample `json:"counterexamples,omitempty"`

	Webhooks []Webhook `json:"webhooks,omitempty"`

	// Whether the new data is to be appended to the existing data in the object. If **append**=`false`, elements included
	// in the new data completely replace the corresponding existing elements, including all subelements. For example, if
	// the new data for a workspace includes **entities** and **append**=`false`, all existing entities in the workspace
	// are discarded and replaced with the new entities.
	//
	// If **append**=`true`, existing elements are preserved, and the new elements are added. If any elements in the new
	// data collide with existing elements, the update request fails.
	Append *bool `json:"append,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewUpdateWorkspaceOptions : Instantiate UpdateWorkspaceOptions
func (assistant *AssistantV1) NewUpdateWorkspaceOptions(workspaceID string) *UpdateWorkspaceOptions {
	return &UpdateWorkspaceOptions{
		WorkspaceID: core.StringPtr(workspaceID),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *UpdateWorkspaceOptions) SetWorkspaceID(workspaceID string) *UpdateWorkspaceOptions {
	options.WorkspaceID = core.StringPtr(workspaceID)
	return options
}

// SetName : Allow user to set Name
func (options *UpdateWorkspaceOptions) SetName(name string) *UpdateWorkspaceOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetDescription : Allow user to set Description
func (options *UpdateWorkspaceOptions) SetDescription(description string) *UpdateWorkspaceOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetLanguage : Allow user to set Language
func (options *UpdateWorkspaceOptions) SetLanguage(language string) *UpdateWorkspaceOptions {
	options.Language = core.StringPtr(language)
	return options
}

// SetMetadata : Allow user to set Metadata
func (options *UpdateWorkspaceOptions) SetMetadata(metadata map[string]interface{}) *UpdateWorkspaceOptions {
	options.Metadata = metadata
	return options
}

// SetLearningOptOut : Allow user to set LearningOptOut
func (options *UpdateWorkspaceOptions) SetLearningOptOut(learningOptOut bool) *UpdateWorkspaceOptions {
	options.LearningOptOut = core.BoolPtr(learningOptOut)
	return options
}

// SetSystemSettings : Allow user to set SystemSettings
func (options *UpdateWorkspaceOptions) SetSystemSettings(systemSettings *WorkspaceSystemSettings) *UpdateWorkspaceOptions {
	options.SystemSettings = systemSettings
	return options
}

// SetIntents : Allow user to set Intents
func (options *UpdateWorkspaceOptions) SetIntents(intents []CreateIntent) *UpdateWorkspaceOptions {
	options.Intents = intents
	return options
}

// SetEntities : Allow user to set Entities
func (options *UpdateWorkspaceOptions) SetEntities(entities []CreateEntity) *UpdateWorkspaceOptions {
	options.Entities = entities
	return options
}

// SetDialogNodes : Allow user to set DialogNodes
func (options *UpdateWorkspaceOptions) SetDialogNodes(dialogNodes []DialogNode) *UpdateWorkspaceOptions {
	options.DialogNodes = dialogNodes
	return options
}

// SetCounterexamples : Allow user to set Counterexamples
func (options *UpdateWorkspaceOptions) SetCounterexamples(counterexamples []Counterexample) *UpdateWorkspaceOptions {
	options.Counterexamples = counterexamples
	return options
}

// SetWebhooks : Allow user to set Webhooks
func (options *UpdateWorkspaceOptions) SetWebhooks(webhooks []Webhook) *UpdateWorkspaceOptions {
	options.Webhooks = webhooks
	return options
}

// SetAppend : Allow user to set Append
func (options *UpdateWorkspaceOptions) SetAppend(append bool) *UpdateWorkspaceOptions {
	options.Append = core.BoolPtr(append)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *UpdateWorkspaceOptions) SetIncludeAudit(includeAudit bool) *UpdateWorkspaceOptions {
	options.IncludeAudit = core.BoolPtr(includeAudit)
	return options
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
	Value_Type_Patterns = "patterns"
	Value_Type_Synonyms = "synonyms"
)

// ValueCollection : ValueCollection struct
type ValueCollection struct {

	// An array of entity values.
	Values []Value `json:"values" validate:"required"`

	// The pagination data for the returned objects.
	Pagination *Pagination `json:"pagination" validate:"required"`
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
	Headers []WebhookHeader `json:"headers,omitempty"`
}

// NewWebhook : Instantiate Webhook (Generic Model Constructor)
func (assistant *AssistantV1) NewWebhook(URL string, name string) (model *Webhook, err error) {
	model = &Webhook{
		URL:  core.StringPtr(URL),
		Name: core.StringPtr(name),
	}
	err = core.ValidateStruct(model, "required parameters")
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
func (assistant *AssistantV1) NewWebhookHeader(name string, value string) (model *WebhookHeader, err error) {
	model = &WebhookHeader{
		Name:  core.StringPtr(name),
		Value: core.StringPtr(value),
	}
	err = core.ValidateStruct(model, "required parameters")
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

	// Any metadata related to the workspace.
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// Whether training data from the workspace (including artifacts such as intents and entities) can be used by IBM for
	// general service improvements. `true` indicates that workspace training data is not to be used.
	LearningOptOut *bool `json:"learning_opt_out" validate:"required"`

	// Global settings for the workspace.
	SystemSettings *WorkspaceSystemSettings `json:"system_settings,omitempty"`

	// The workspace ID of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The current status of the workspace.
	Status *string `json:"status,omitempty"`

	// The timestamp for creation of the object.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// The timestamp for the most recent update to the object.
	Updated *strfmt.DateTime `json:"updated,omitempty"`

	// An array of intents.
	Intents []Intent `json:"intents,omitempty"`

	// An array of objects describing the entities for the workspace.
	Entities []Entity `json:"entities,omitempty"`

	// An array of objects describing the dialog nodes in the workspace.
	DialogNodes []DialogNode `json:"dialog_nodes,omitempty"`

	// An array of counterexamples.
	Counterexamples []Counterexample `json:"counterexamples,omitempty"`

	Webhooks []Webhook `json:"webhooks,omitempty"`
}

// Constants associated with the Workspace.Status property.
// The current status of the workspace.
const (
	Workspace_Status_Available   = "Available"
	Workspace_Status_Failed      = "Failed"
	Workspace_Status_NonExistent = "Non Existent"
	Workspace_Status_Training    = "Training"
	Workspace_Status_Unavailable = "Unavailable"
)

// WorkspaceCollection : WorkspaceCollection struct
type WorkspaceCollection struct {

	// An array of objects describing the workspaces associated with the service instance.
	Workspaces []Workspace `json:"workspaces" validate:"required"`

	// The pagination data for the returned objects.
	Pagination *Pagination `json:"pagination" validate:"required"`
}

// WorkspaceSystemSettings : Global settings for the workspace.
type WorkspaceSystemSettings struct {

	// Workspace settings related to the Watson Assistant user interface.
	Tooling *WorkspaceSystemSettingsTooling `json:"tooling,omitempty"`

	// Workspace settings related to the disambiguation feature.
	//
	// **Note:** This feature is available only to Plus and Premium users.
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

// WorkspaceSystemSettingsDisambiguation : Workspace settings related to the disambiguation feature.
//
// **Note:** This feature is available only to Plus and Premium users.
type WorkspaceSystemSettingsDisambiguation struct {

	// The text of the introductory prompt that accompanies disambiguation options presented to the user.
	Prompt *string `json:"prompt,omitempty"`

	// The user-facing label for the option users can select if none of the suggested options is correct. If no value is
	// specified for this property, this option does not appear.
	NoneOfTheAbovePrompt *string `json:"none_of_the_above_prompt,omitempty"`

	// Whether the disambiguation feature is enabled for the workspace.
	Enabled *bool `json:"enabled,omitempty"`

	// The sensitivity of the disambiguation feature to intent detection conflicts. Set to **high** if you want the
	// disambiguation feature to be triggered more often. This can be useful for testing or demonstration purposes.
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
// The sensitivity of the disambiguation feature to intent detection conflicts. Set to **high** if you want the
// disambiguation feature to be triggered more often. This can be useful for testing or demonstration purposes.
const (
	WorkspaceSystemSettingsDisambiguation_Sensitivity_Auto = "auto"
	WorkspaceSystemSettingsDisambiguation_Sensitivity_High = "high"
)

// WorkspaceSystemSettingsOffTopic : Workspace settings related to detection of irrelevant input.
type WorkspaceSystemSettingsOffTopic struct {

	// Whether enhanced irrelevance detection is enabled for the workspace.
	Enabled *bool `json:"enabled,omitempty"`
}

// WorkspaceSystemSettingsSystemEntities : Workspace settings related to the behavior of system entities.
type WorkspaceSystemSettingsSystemEntities struct {

	// Whether the new system entities are enabled for the workspace.
	Enabled *bool `json:"enabled,omitempty"`
}

// WorkspaceSystemSettingsTooling : Workspace settings related to the Watson Assistant user interface.
type WorkspaceSystemSettingsTooling struct {

	// Whether the dialog JSON editor displays text responses within the `output.generic` object.
	StoreGenericResponses *bool `json:"store_generic_responses,omitempty"`
}
