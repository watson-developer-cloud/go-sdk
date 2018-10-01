// Package assistantv1 : Operations and models for the AssistantV1 service
package assistantv1

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
	core "github.com/ibm-watson/go-sdk/core"
)

// AssistantV1 : Assistant service combines machine learning, natural language understanding, and integrated
// dialog tools to create conversation flows between your apps and your users.
//
// Version: V1
// See: http://www.ibm.com/watson/developercloud/assistant.html
//
type AssistantV1 struct {
	service *core.WatsonService
}

// AssistantV1Options : Service options
type AssistantV1Options struct {
	Version        string
	URL            string
	Username       string
	Password       string
	IAMApiKey      string
	IAMAccessToken string
	IAMURL         string
}

// NewAssistantV1 : Instantiate AssistantV1
func NewAssistantV1(options *AssistantV1Options) (*AssistantV1, error) {
	if options.URL == "" {
		options.URL = "https://gateway.watsonplatform.net/assistant/api"
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
	service, serviceErr := core.NewWatsonService(serviceOptions, "conversation")
	if serviceErr != nil {
		return nil, serviceErr
	}

	return &AssistantV1{service: service}, nil
}

// Message : Get response to user input
func (assistant *AssistantV1) Message(messageOptions *MessageOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(messageOptions, "messageOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(messageOptions, "messageOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces", "message"}
	pathParameters := []string{*messageOptions.WorkspaceID}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range messageOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	if messageOptions.NodesVisitedDetails != nil {
		builder.AddQuery("nodes_visited_details", fmt.Sprint(*messageOptions.NodesVisitedDetails))
	}
	builder.AddQuery("version", assistant.service.Options.Version)

	body := make(map[string]interface{})
	if messageOptions.Input != nil {
		body["input"] = messageOptions.Input
	}
	if messageOptions.AlternateIntents != nil {
		body["alternate_intents"] = messageOptions.AlternateIntents
	}
	if messageOptions.Context != nil {
		body["context"] = messageOptions.Context
	}
	if messageOptions.Entities != nil {
		body["entities"] = messageOptions.Entities
	}
	if messageOptions.Intents != nil {
		body["intents"] = messageOptions.Intents
	}
	if messageOptions.Output != nil {
		body["output"] = messageOptions.Output
	}
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, new(MessageResponse))
	return response, err
}

// GetMessageResult : Retrieve result of Message operation
func (assistant *AssistantV1) GetMessageResult(response *core.DetailedResponse) *MessageResponse {
	result, ok := response.Result.(*MessageResponse)
	if ok {
		return result
	}
	return nil
}

// CreateWorkspace : Create workspace
func (assistant *AssistantV1) CreateWorkspace(createWorkspaceOptions *CreateWorkspaceOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateStruct(createWorkspaceOptions, "createWorkspaceOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range createWorkspaceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", assistant.service.Options.Version)

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
	if createWorkspaceOptions.Metadata != nil {
		body["metadata"] = createWorkspaceOptions.Metadata
	}
	if createWorkspaceOptions.LearningOptOut != nil {
		body["learning_opt_out"] = createWorkspaceOptions.LearningOptOut
	}
	if createWorkspaceOptions.SystemSettings != nil {
		body["system_settings"] = createWorkspaceOptions.SystemSettings
	}
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, new(Workspace))
	return response, err
}

// GetCreateWorkspaceResult : Retrieve result of CreateWorkspace operation
func (assistant *AssistantV1) GetCreateWorkspaceResult(response *core.DetailedResponse) *Workspace {
	result, ok := response.Result.(*Workspace)
	if ok {
		return result
	}
	return nil
}

// DeleteWorkspace : Delete workspace
func (assistant *AssistantV1) DeleteWorkspace(deleteWorkspaceOptions *DeleteWorkspaceOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(deleteWorkspaceOptions, "deleteWorkspaceOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(deleteWorkspaceOptions, "deleteWorkspaceOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces"}
	pathParameters := []string{*deleteWorkspaceOptions.WorkspaceID}

	builder := core.NewRequestBuilder(core.DELETE)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteWorkspaceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", assistant.service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, nil)
	return response, err
}

// GetWorkspace : Get information about a workspace
func (assistant *AssistantV1) GetWorkspace(getWorkspaceOptions *GetWorkspaceOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(getWorkspaceOptions, "getWorkspaceOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(getWorkspaceOptions, "getWorkspaceOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces"}
	pathParameters := []string{*getWorkspaceOptions.WorkspaceID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range getWorkspaceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getWorkspaceOptions.Export != nil {
		builder.AddQuery("export", fmt.Sprint(*getWorkspaceOptions.Export))
	}
	if getWorkspaceOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*getWorkspaceOptions.IncludeAudit))
	}
	builder.AddQuery("version", assistant.service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, new(WorkspaceExport))
	return response, err
}

// GetGetWorkspaceResult : Retrieve result of GetWorkspace operation
func (assistant *AssistantV1) GetGetWorkspaceResult(response *core.DetailedResponse) *WorkspaceExport {
	result, ok := response.Result.(*WorkspaceExport)
	if ok {
		return result
	}
	return nil
}

// ListWorkspaces : List workspaces
func (assistant *AssistantV1) ListWorkspaces(listWorkspacesOptions *ListWorkspacesOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateStruct(listWorkspacesOptions, "listWorkspacesOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range listWorkspacesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

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
	builder.AddQuery("version", assistant.service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, new(WorkspaceCollection))
	return response, err
}

// GetListWorkspacesResult : Retrieve result of ListWorkspaces operation
func (assistant *AssistantV1) GetListWorkspacesResult(response *core.DetailedResponse) *WorkspaceCollection {
	result, ok := response.Result.(*WorkspaceCollection)
	if ok {
		return result
	}
	return nil
}

// UpdateWorkspace : Update workspace
func (assistant *AssistantV1) UpdateWorkspace(updateWorkspaceOptions *UpdateWorkspaceOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(updateWorkspaceOptions, "updateWorkspaceOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(updateWorkspaceOptions, "updateWorkspaceOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces"}
	pathParameters := []string{*updateWorkspaceOptions.WorkspaceID}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range updateWorkspaceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	if updateWorkspaceOptions.Append != nil {
		builder.AddQuery("append", fmt.Sprint(*updateWorkspaceOptions.Append))
	}
	builder.AddQuery("version", assistant.service.Options.Version)

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
	if updateWorkspaceOptions.Metadata != nil {
		body["metadata"] = updateWorkspaceOptions.Metadata
	}
	if updateWorkspaceOptions.LearningOptOut != nil {
		body["learning_opt_out"] = updateWorkspaceOptions.LearningOptOut
	}
	if updateWorkspaceOptions.SystemSettings != nil {
		body["system_settings"] = updateWorkspaceOptions.SystemSettings
	}
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, new(Workspace))
	return response, err
}

// GetUpdateWorkspaceResult : Retrieve result of UpdateWorkspace operation
func (assistant *AssistantV1) GetUpdateWorkspaceResult(response *core.DetailedResponse) *Workspace {
	result, ok := response.Result.(*Workspace)
	if ok {
		return result
	}
	return nil
}

// CreateIntent : Create intent
func (assistant *AssistantV1) CreateIntent(createIntentOptions *CreateIntentOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(createIntentOptions, "createIntentOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(createIntentOptions, "createIntentOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces", "intents"}
	pathParameters := []string{*createIntentOptions.WorkspaceID}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range createIntentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", assistant.service.Options.Version)

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
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, new(Intent))
	return response, err
}

// GetCreateIntentResult : Retrieve result of CreateIntent operation
func (assistant *AssistantV1) GetCreateIntentResult(response *core.DetailedResponse) *Intent {
	result, ok := response.Result.(*Intent)
	if ok {
		return result
	}
	return nil
}

// DeleteIntent : Delete intent
func (assistant *AssistantV1) DeleteIntent(deleteIntentOptions *DeleteIntentOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(deleteIntentOptions, "deleteIntentOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(deleteIntentOptions, "deleteIntentOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces", "intents"}
	pathParameters := []string{*deleteIntentOptions.WorkspaceID, *deleteIntentOptions.Intent}

	builder := core.NewRequestBuilder(core.DELETE)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteIntentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", assistant.service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, nil)
	return response, err
}

// GetIntent : Get intent
func (assistant *AssistantV1) GetIntent(getIntentOptions *GetIntentOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(getIntentOptions, "getIntentOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(getIntentOptions, "getIntentOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces", "intents"}
	pathParameters := []string{*getIntentOptions.WorkspaceID, *getIntentOptions.Intent}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range getIntentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getIntentOptions.Export != nil {
		builder.AddQuery("export", fmt.Sprint(*getIntentOptions.Export))
	}
	if getIntentOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*getIntentOptions.IncludeAudit))
	}
	builder.AddQuery("version", assistant.service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, new(IntentExport))
	return response, err
}

// GetGetIntentResult : Retrieve result of GetIntent operation
func (assistant *AssistantV1) GetGetIntentResult(response *core.DetailedResponse) *IntentExport {
	result, ok := response.Result.(*IntentExport)
	if ok {
		return result
	}
	return nil
}

// ListIntents : List intents
func (assistant *AssistantV1) ListIntents(listIntentsOptions *ListIntentsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(listIntentsOptions, "listIntentsOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(listIntentsOptions, "listIntentsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces", "intents"}
	pathParameters := []string{*listIntentsOptions.WorkspaceID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range listIntentsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

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
	builder.AddQuery("version", assistant.service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, new(IntentCollection))
	return response, err
}

// GetListIntentsResult : Retrieve result of ListIntents operation
func (assistant *AssistantV1) GetListIntentsResult(response *core.DetailedResponse) *IntentCollection {
	result, ok := response.Result.(*IntentCollection)
	if ok {
		return result
	}
	return nil
}

// UpdateIntent : Update intent
func (assistant *AssistantV1) UpdateIntent(updateIntentOptions *UpdateIntentOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(updateIntentOptions, "updateIntentOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(updateIntentOptions, "updateIntentOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces", "intents"}
	pathParameters := []string{*updateIntentOptions.WorkspaceID, *updateIntentOptions.Intent}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range updateIntentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", assistant.service.Options.Version)

	body := make(map[string]interface{})
	if updateIntentOptions.NewIntent != nil {
		body["intent"] = updateIntentOptions.NewIntent
	}
	if updateIntentOptions.NewExamples != nil {
		body["examples"] = updateIntentOptions.NewExamples
	}
	if updateIntentOptions.NewDescription != nil {
		body["description"] = updateIntentOptions.NewDescription
	}
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, new(Intent))
	return response, err
}

// GetUpdateIntentResult : Retrieve result of UpdateIntent operation
func (assistant *AssistantV1) GetUpdateIntentResult(response *core.DetailedResponse) *Intent {
	result, ok := response.Result.(*Intent)
	if ok {
		return result
	}
	return nil
}

// CreateExample : Create user input example
func (assistant *AssistantV1) CreateExample(createExampleOptions *CreateExampleOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(createExampleOptions, "createExampleOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(createExampleOptions, "createExampleOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces", "intents", "examples"}
	pathParameters := []string{*createExampleOptions.WorkspaceID, *createExampleOptions.Intent}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range createExampleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", assistant.service.Options.Version)

	body := make(map[string]interface{})
	if createExampleOptions.Text != nil {
		body["text"] = createExampleOptions.Text
	}
	if createExampleOptions.Mentions != nil {
		body["mentions"] = createExampleOptions.Mentions
	}
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, new(Example))
	return response, err
}

// GetCreateExampleResult : Retrieve result of CreateExample operation
func (assistant *AssistantV1) GetCreateExampleResult(response *core.DetailedResponse) *Example {
	result, ok := response.Result.(*Example)
	if ok {
		return result
	}
	return nil
}

// DeleteExample : Delete user input example
func (assistant *AssistantV1) DeleteExample(deleteExampleOptions *DeleteExampleOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(deleteExampleOptions, "deleteExampleOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(deleteExampleOptions, "deleteExampleOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces", "intents", "examples"}
	pathParameters := []string{*deleteExampleOptions.WorkspaceID, *deleteExampleOptions.Intent, *deleteExampleOptions.Text}

	builder := core.NewRequestBuilder(core.DELETE)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteExampleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", assistant.service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, nil)
	return response, err
}

// GetExample : Get user input example
func (assistant *AssistantV1) GetExample(getExampleOptions *GetExampleOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(getExampleOptions, "getExampleOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(getExampleOptions, "getExampleOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces", "intents", "examples"}
	pathParameters := []string{*getExampleOptions.WorkspaceID, *getExampleOptions.Intent, *getExampleOptions.Text}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range getExampleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getExampleOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*getExampleOptions.IncludeAudit))
	}
	builder.AddQuery("version", assistant.service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, new(Example))
	return response, err
}

// GetGetExampleResult : Retrieve result of GetExample operation
func (assistant *AssistantV1) GetGetExampleResult(response *core.DetailedResponse) *Example {
	result, ok := response.Result.(*Example)
	if ok {
		return result
	}
	return nil
}

// ListExamples : List user input examples
func (assistant *AssistantV1) ListExamples(listExamplesOptions *ListExamplesOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(listExamplesOptions, "listExamplesOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(listExamplesOptions, "listExamplesOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces", "intents", "examples"}
	pathParameters := []string{*listExamplesOptions.WorkspaceID, *listExamplesOptions.Intent}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range listExamplesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

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
	builder.AddQuery("version", assistant.service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, new(ExampleCollection))
	return response, err
}

// GetListExamplesResult : Retrieve result of ListExamples operation
func (assistant *AssistantV1) GetListExamplesResult(response *core.DetailedResponse) *ExampleCollection {
	result, ok := response.Result.(*ExampleCollection)
	if ok {
		return result
	}
	return nil
}

// UpdateExample : Update user input example
func (assistant *AssistantV1) UpdateExample(updateExampleOptions *UpdateExampleOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(updateExampleOptions, "updateExampleOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(updateExampleOptions, "updateExampleOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces", "intents", "examples"}
	pathParameters := []string{*updateExampleOptions.WorkspaceID, *updateExampleOptions.Intent, *updateExampleOptions.Text}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range updateExampleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", assistant.service.Options.Version)

	body := make(map[string]interface{})
	if updateExampleOptions.NewText != nil {
		body["text"] = updateExampleOptions.NewText
	}
	if updateExampleOptions.NewMentions != nil {
		body["mentions"] = updateExampleOptions.NewMentions
	}
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, new(Example))
	return response, err
}

// GetUpdateExampleResult : Retrieve result of UpdateExample operation
func (assistant *AssistantV1) GetUpdateExampleResult(response *core.DetailedResponse) *Example {
	result, ok := response.Result.(*Example)
	if ok {
		return result
	}
	return nil
}

// CreateCounterexample : Create counterexample
func (assistant *AssistantV1) CreateCounterexample(createCounterexampleOptions *CreateCounterexampleOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(createCounterexampleOptions, "createCounterexampleOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(createCounterexampleOptions, "createCounterexampleOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces", "counterexamples"}
	pathParameters := []string{*createCounterexampleOptions.WorkspaceID}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range createCounterexampleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", assistant.service.Options.Version)

	body := make(map[string]interface{})
	if createCounterexampleOptions.Text != nil {
		body["text"] = createCounterexampleOptions.Text
	}
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, new(Counterexample))
	return response, err
}

// GetCreateCounterexampleResult : Retrieve result of CreateCounterexample operation
func (assistant *AssistantV1) GetCreateCounterexampleResult(response *core.DetailedResponse) *Counterexample {
	result, ok := response.Result.(*Counterexample)
	if ok {
		return result
	}
	return nil
}

// DeleteCounterexample : Delete counterexample
func (assistant *AssistantV1) DeleteCounterexample(deleteCounterexampleOptions *DeleteCounterexampleOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(deleteCounterexampleOptions, "deleteCounterexampleOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(deleteCounterexampleOptions, "deleteCounterexampleOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces", "counterexamples"}
	pathParameters := []string{*deleteCounterexampleOptions.WorkspaceID, *deleteCounterexampleOptions.Text}

	builder := core.NewRequestBuilder(core.DELETE)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteCounterexampleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", assistant.service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, nil)
	return response, err
}

// GetCounterexample : Get counterexample
func (assistant *AssistantV1) GetCounterexample(getCounterexampleOptions *GetCounterexampleOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(getCounterexampleOptions, "getCounterexampleOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(getCounterexampleOptions, "getCounterexampleOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces", "counterexamples"}
	pathParameters := []string{*getCounterexampleOptions.WorkspaceID, *getCounterexampleOptions.Text}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range getCounterexampleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getCounterexampleOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*getCounterexampleOptions.IncludeAudit))
	}
	builder.AddQuery("version", assistant.service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, new(Counterexample))
	return response, err
}

// GetGetCounterexampleResult : Retrieve result of GetCounterexample operation
func (assistant *AssistantV1) GetGetCounterexampleResult(response *core.DetailedResponse) *Counterexample {
	result, ok := response.Result.(*Counterexample)
	if ok {
		return result
	}
	return nil
}

// ListCounterexamples : List counterexamples
func (assistant *AssistantV1) ListCounterexamples(listCounterexamplesOptions *ListCounterexamplesOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(listCounterexamplesOptions, "listCounterexamplesOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(listCounterexamplesOptions, "listCounterexamplesOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces", "counterexamples"}
	pathParameters := []string{*listCounterexamplesOptions.WorkspaceID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range listCounterexamplesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

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
	builder.AddQuery("version", assistant.service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, new(CounterexampleCollection))
	return response, err
}

// GetListCounterexamplesResult : Retrieve result of ListCounterexamples operation
func (assistant *AssistantV1) GetListCounterexamplesResult(response *core.DetailedResponse) *CounterexampleCollection {
	result, ok := response.Result.(*CounterexampleCollection)
	if ok {
		return result
	}
	return nil
}

// UpdateCounterexample : Update counterexample
func (assistant *AssistantV1) UpdateCounterexample(updateCounterexampleOptions *UpdateCounterexampleOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(updateCounterexampleOptions, "updateCounterexampleOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(updateCounterexampleOptions, "updateCounterexampleOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces", "counterexamples"}
	pathParameters := []string{*updateCounterexampleOptions.WorkspaceID, *updateCounterexampleOptions.Text}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range updateCounterexampleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", assistant.service.Options.Version)

	body := make(map[string]interface{})
	if updateCounterexampleOptions.NewText != nil {
		body["text"] = updateCounterexampleOptions.NewText
	}
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, new(Counterexample))
	return response, err
}

// GetUpdateCounterexampleResult : Retrieve result of UpdateCounterexample operation
func (assistant *AssistantV1) GetUpdateCounterexampleResult(response *core.DetailedResponse) *Counterexample {
	result, ok := response.Result.(*Counterexample)
	if ok {
		return result
	}
	return nil
}

// CreateEntity : Create entity
func (assistant *AssistantV1) CreateEntity(createEntityOptions *CreateEntityOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(createEntityOptions, "createEntityOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(createEntityOptions, "createEntityOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces", "entities"}
	pathParameters := []string{*createEntityOptions.WorkspaceID}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range createEntityOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", assistant.service.Options.Version)

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
	if createEntityOptions.Values != nil {
		body["values"] = createEntityOptions.Values
	}
	if createEntityOptions.FuzzyMatch != nil {
		body["fuzzy_match"] = createEntityOptions.FuzzyMatch
	}
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, new(Entity))
	return response, err
}

// GetCreateEntityResult : Retrieve result of CreateEntity operation
func (assistant *AssistantV1) GetCreateEntityResult(response *core.DetailedResponse) *Entity {
	result, ok := response.Result.(*Entity)
	if ok {
		return result
	}
	return nil
}

// DeleteEntity : Delete entity
func (assistant *AssistantV1) DeleteEntity(deleteEntityOptions *DeleteEntityOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(deleteEntityOptions, "deleteEntityOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(deleteEntityOptions, "deleteEntityOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces", "entities"}
	pathParameters := []string{*deleteEntityOptions.WorkspaceID, *deleteEntityOptions.Entity}

	builder := core.NewRequestBuilder(core.DELETE)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteEntityOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", assistant.service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, nil)
	return response, err
}

// GetEntity : Get entity
func (assistant *AssistantV1) GetEntity(getEntityOptions *GetEntityOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(getEntityOptions, "getEntityOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(getEntityOptions, "getEntityOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces", "entities"}
	pathParameters := []string{*getEntityOptions.WorkspaceID, *getEntityOptions.Entity}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range getEntityOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getEntityOptions.Export != nil {
		builder.AddQuery("export", fmt.Sprint(*getEntityOptions.Export))
	}
	if getEntityOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*getEntityOptions.IncludeAudit))
	}
	builder.AddQuery("version", assistant.service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, new(EntityExport))
	return response, err
}

// GetGetEntityResult : Retrieve result of GetEntity operation
func (assistant *AssistantV1) GetGetEntityResult(response *core.DetailedResponse) *EntityExport {
	result, ok := response.Result.(*EntityExport)
	if ok {
		return result
	}
	return nil
}

// ListEntities : List entities
func (assistant *AssistantV1) ListEntities(listEntitiesOptions *ListEntitiesOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(listEntitiesOptions, "listEntitiesOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(listEntitiesOptions, "listEntitiesOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces", "entities"}
	pathParameters := []string{*listEntitiesOptions.WorkspaceID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range listEntitiesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

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
	builder.AddQuery("version", assistant.service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, new(EntityCollection))
	return response, err
}

// GetListEntitiesResult : Retrieve result of ListEntities operation
func (assistant *AssistantV1) GetListEntitiesResult(response *core.DetailedResponse) *EntityCollection {
	result, ok := response.Result.(*EntityCollection)
	if ok {
		return result
	}
	return nil
}

// UpdateEntity : Update entity
func (assistant *AssistantV1) UpdateEntity(updateEntityOptions *UpdateEntityOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(updateEntityOptions, "updateEntityOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(updateEntityOptions, "updateEntityOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces", "entities"}
	pathParameters := []string{*updateEntityOptions.WorkspaceID, *updateEntityOptions.Entity}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range updateEntityOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", assistant.service.Options.Version)

	body := make(map[string]interface{})
	if updateEntityOptions.NewFuzzyMatch != nil {
		body["fuzzy_match"] = updateEntityOptions.NewFuzzyMatch
	}
	if updateEntityOptions.NewEntity != nil {
		body["entity"] = updateEntityOptions.NewEntity
	}
	if updateEntityOptions.NewMetadata != nil {
		body["metadata"] = updateEntityOptions.NewMetadata
	}
	if updateEntityOptions.NewValues != nil {
		body["values"] = updateEntityOptions.NewValues
	}
	if updateEntityOptions.NewDescription != nil {
		body["description"] = updateEntityOptions.NewDescription
	}
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, new(Entity))
	return response, err
}

// GetUpdateEntityResult : Retrieve result of UpdateEntity operation
func (assistant *AssistantV1) GetUpdateEntityResult(response *core.DetailedResponse) *Entity {
	result, ok := response.Result.(*Entity)
	if ok {
		return result
	}
	return nil
}

// ListMentions : List entity mentions
func (assistant *AssistantV1) ListMentions(listMentionsOptions *ListMentionsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(listMentionsOptions, "listMentionsOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(listMentionsOptions, "listMentionsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces", "entities", "mentions"}
	pathParameters := []string{*listMentionsOptions.WorkspaceID, *listMentionsOptions.Entity}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range listMentionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listMentionsOptions.Export != nil {
		builder.AddQuery("export", fmt.Sprint(*listMentionsOptions.Export))
	}
	if listMentionsOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*listMentionsOptions.IncludeAudit))
	}
	builder.AddQuery("version", assistant.service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, new(EntityMentionCollection))
	return response, err
}

// GetListMentionsResult : Retrieve result of ListMentions operation
func (assistant *AssistantV1) GetListMentionsResult(response *core.DetailedResponse) *EntityMentionCollection {
	result, ok := response.Result.(*EntityMentionCollection)
	if ok {
		return result
	}
	return nil
}

// CreateValue : Add entity value
func (assistant *AssistantV1) CreateValue(createValueOptions *CreateValueOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(createValueOptions, "createValueOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(createValueOptions, "createValueOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces", "entities", "values"}
	pathParameters := []string{*createValueOptions.WorkspaceID, *createValueOptions.Entity}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range createValueOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", assistant.service.Options.Version)

	body := make(map[string]interface{})
	if createValueOptions.Value != nil {
		body["value"] = createValueOptions.Value
	}
	if createValueOptions.Metadata != nil {
		body["metadata"] = createValueOptions.Metadata
	}
	if createValueOptions.Synonyms != nil {
		body["synonyms"] = createValueOptions.Synonyms
	}
	if createValueOptions.Patterns != nil {
		body["patterns"] = createValueOptions.Patterns
	}
	if createValueOptions.ValueType != nil {
		body["type"] = createValueOptions.ValueType
	}
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, new(Value))
	return response, err
}

// GetCreateValueResult : Retrieve result of CreateValue operation
func (assistant *AssistantV1) GetCreateValueResult(response *core.DetailedResponse) *Value {
	result, ok := response.Result.(*Value)
	if ok {
		return result
	}
	return nil
}

// DeleteValue : Delete entity value
func (assistant *AssistantV1) DeleteValue(deleteValueOptions *DeleteValueOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(deleteValueOptions, "deleteValueOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(deleteValueOptions, "deleteValueOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces", "entities", "values"}
	pathParameters := []string{*deleteValueOptions.WorkspaceID, *deleteValueOptions.Entity, *deleteValueOptions.Value}

	builder := core.NewRequestBuilder(core.DELETE)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteValueOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", assistant.service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, nil)
	return response, err
}

// GetValue : Get entity value
func (assistant *AssistantV1) GetValue(getValueOptions *GetValueOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(getValueOptions, "getValueOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(getValueOptions, "getValueOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces", "entities", "values"}
	pathParameters := []string{*getValueOptions.WorkspaceID, *getValueOptions.Entity, *getValueOptions.Value}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range getValueOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getValueOptions.Export != nil {
		builder.AddQuery("export", fmt.Sprint(*getValueOptions.Export))
	}
	if getValueOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*getValueOptions.IncludeAudit))
	}
	builder.AddQuery("version", assistant.service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, new(ValueExport))
	return response, err
}

// GetGetValueResult : Retrieve result of GetValue operation
func (assistant *AssistantV1) GetGetValueResult(response *core.DetailedResponse) *ValueExport {
	result, ok := response.Result.(*ValueExport)
	if ok {
		return result
	}
	return nil
}

// ListValues : List entity values
func (assistant *AssistantV1) ListValues(listValuesOptions *ListValuesOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(listValuesOptions, "listValuesOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(listValuesOptions, "listValuesOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces", "entities", "values"}
	pathParameters := []string{*listValuesOptions.WorkspaceID, *listValuesOptions.Entity}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range listValuesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

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
	builder.AddQuery("version", assistant.service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, new(ValueCollection))
	return response, err
}

// GetListValuesResult : Retrieve result of ListValues operation
func (assistant *AssistantV1) GetListValuesResult(response *core.DetailedResponse) *ValueCollection {
	result, ok := response.Result.(*ValueCollection)
	if ok {
		return result
	}
	return nil
}

// UpdateValue : Update entity value
func (assistant *AssistantV1) UpdateValue(updateValueOptions *UpdateValueOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(updateValueOptions, "updateValueOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(updateValueOptions, "updateValueOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces", "entities", "values"}
	pathParameters := []string{*updateValueOptions.WorkspaceID, *updateValueOptions.Entity, *updateValueOptions.Value}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range updateValueOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", assistant.service.Options.Version)

	body := make(map[string]interface{})
	if updateValueOptions.NewSynonyms != nil {
		body["synonyms"] = updateValueOptions.NewSynonyms
	}
	if updateValueOptions.ValueType != nil {
		body["type"] = updateValueOptions.ValueType
	}
	if updateValueOptions.NewMetadata != nil {
		body["metadata"] = updateValueOptions.NewMetadata
	}
	if updateValueOptions.NewPatterns != nil {
		body["patterns"] = updateValueOptions.NewPatterns
	}
	if updateValueOptions.NewValue != nil {
		body["value"] = updateValueOptions.NewValue
	}
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, new(Value))
	return response, err
}

// GetUpdateValueResult : Retrieve result of UpdateValue operation
func (assistant *AssistantV1) GetUpdateValueResult(response *core.DetailedResponse) *Value {
	result, ok := response.Result.(*Value)
	if ok {
		return result
	}
	return nil
}

// CreateSynonym : Add entity value synonym
func (assistant *AssistantV1) CreateSynonym(createSynonymOptions *CreateSynonymOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(createSynonymOptions, "createSynonymOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(createSynonymOptions, "createSynonymOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces", "entities", "values", "synonyms"}
	pathParameters := []string{*createSynonymOptions.WorkspaceID, *createSynonymOptions.Entity, *createSynonymOptions.Value}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range createSynonymOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", assistant.service.Options.Version)

	body := make(map[string]interface{})
	if createSynonymOptions.Synonym != nil {
		body["synonym"] = createSynonymOptions.Synonym
	}
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, new(Synonym))
	return response, err
}

// GetCreateSynonymResult : Retrieve result of CreateSynonym operation
func (assistant *AssistantV1) GetCreateSynonymResult(response *core.DetailedResponse) *Synonym {
	result, ok := response.Result.(*Synonym)
	if ok {
		return result
	}
	return nil
}

// DeleteSynonym : Delete entity value synonym
func (assistant *AssistantV1) DeleteSynonym(deleteSynonymOptions *DeleteSynonymOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(deleteSynonymOptions, "deleteSynonymOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(deleteSynonymOptions, "deleteSynonymOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces", "entities", "values", "synonyms"}
	pathParameters := []string{*deleteSynonymOptions.WorkspaceID, *deleteSynonymOptions.Entity, *deleteSynonymOptions.Value, *deleteSynonymOptions.Synonym}

	builder := core.NewRequestBuilder(core.DELETE)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteSynonymOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", assistant.service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, nil)
	return response, err
}

// GetSynonym : Get entity value synonym
func (assistant *AssistantV1) GetSynonym(getSynonymOptions *GetSynonymOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(getSynonymOptions, "getSynonymOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(getSynonymOptions, "getSynonymOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces", "entities", "values", "synonyms"}
	pathParameters := []string{*getSynonymOptions.WorkspaceID, *getSynonymOptions.Entity, *getSynonymOptions.Value, *getSynonymOptions.Synonym}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range getSynonymOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getSynonymOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*getSynonymOptions.IncludeAudit))
	}
	builder.AddQuery("version", assistant.service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, new(Synonym))
	return response, err
}

// GetGetSynonymResult : Retrieve result of GetSynonym operation
func (assistant *AssistantV1) GetGetSynonymResult(response *core.DetailedResponse) *Synonym {
	result, ok := response.Result.(*Synonym)
	if ok {
		return result
	}
	return nil
}

// ListSynonyms : List entity value synonyms
func (assistant *AssistantV1) ListSynonyms(listSynonymsOptions *ListSynonymsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(listSynonymsOptions, "listSynonymsOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(listSynonymsOptions, "listSynonymsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces", "entities", "values", "synonyms"}
	pathParameters := []string{*listSynonymsOptions.WorkspaceID, *listSynonymsOptions.Entity, *listSynonymsOptions.Value}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range listSynonymsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

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
	builder.AddQuery("version", assistant.service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, new(SynonymCollection))
	return response, err
}

// GetListSynonymsResult : Retrieve result of ListSynonyms operation
func (assistant *AssistantV1) GetListSynonymsResult(response *core.DetailedResponse) *SynonymCollection {
	result, ok := response.Result.(*SynonymCollection)
	if ok {
		return result
	}
	return nil
}

// UpdateSynonym : Update entity value synonym
func (assistant *AssistantV1) UpdateSynonym(updateSynonymOptions *UpdateSynonymOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(updateSynonymOptions, "updateSynonymOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(updateSynonymOptions, "updateSynonymOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces", "entities", "values", "synonyms"}
	pathParameters := []string{*updateSynonymOptions.WorkspaceID, *updateSynonymOptions.Entity, *updateSynonymOptions.Value, *updateSynonymOptions.Synonym}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range updateSynonymOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", assistant.service.Options.Version)

	body := make(map[string]interface{})
	if updateSynonymOptions.NewSynonym != nil {
		body["synonym"] = updateSynonymOptions.NewSynonym
	}
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, new(Synonym))
	return response, err
}

// GetUpdateSynonymResult : Retrieve result of UpdateSynonym operation
func (assistant *AssistantV1) GetUpdateSynonymResult(response *core.DetailedResponse) *Synonym {
	result, ok := response.Result.(*Synonym)
	if ok {
		return result
	}
	return nil
}

// CreateDialogNode : Create dialog node
func (assistant *AssistantV1) CreateDialogNode(createDialogNodeOptions *CreateDialogNodeOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(createDialogNodeOptions, "createDialogNodeOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(createDialogNodeOptions, "createDialogNodeOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces", "dialog_nodes"}
	pathParameters := []string{*createDialogNodeOptions.WorkspaceID}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range createDialogNodeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", assistant.service.Options.Version)

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
	if createDialogNodeOptions.Actions != nil {
		body["actions"] = createDialogNodeOptions.Actions
	}
	if createDialogNodeOptions.Title != nil {
		body["title"] = createDialogNodeOptions.Title
	}
	if createDialogNodeOptions.NodeType != nil {
		body["type"] = createDialogNodeOptions.NodeType
	}
	if createDialogNodeOptions.EventName != nil {
		body["event_name"] = createDialogNodeOptions.EventName
	}
	if createDialogNodeOptions.Variable != nil {
		body["variable"] = createDialogNodeOptions.Variable
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
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, new(DialogNode))
	return response, err
}

// GetCreateDialogNodeResult : Retrieve result of CreateDialogNode operation
func (assistant *AssistantV1) GetCreateDialogNodeResult(response *core.DetailedResponse) *DialogNode {
	result, ok := response.Result.(*DialogNode)
	if ok {
		return result
	}
	return nil
}

// DeleteDialogNode : Delete dialog node
func (assistant *AssistantV1) DeleteDialogNode(deleteDialogNodeOptions *DeleteDialogNodeOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(deleteDialogNodeOptions, "deleteDialogNodeOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(deleteDialogNodeOptions, "deleteDialogNodeOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces", "dialog_nodes"}
	pathParameters := []string{*deleteDialogNodeOptions.WorkspaceID, *deleteDialogNodeOptions.DialogNode}

	builder := core.NewRequestBuilder(core.DELETE)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteDialogNodeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", assistant.service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, nil)
	return response, err
}

// GetDialogNode : Get dialog node
func (assistant *AssistantV1) GetDialogNode(getDialogNodeOptions *GetDialogNodeOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(getDialogNodeOptions, "getDialogNodeOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(getDialogNodeOptions, "getDialogNodeOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces", "dialog_nodes"}
	pathParameters := []string{*getDialogNodeOptions.WorkspaceID, *getDialogNodeOptions.DialogNode}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range getDialogNodeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getDialogNodeOptions.IncludeAudit != nil {
		builder.AddQuery("include_audit", fmt.Sprint(*getDialogNodeOptions.IncludeAudit))
	}
	builder.AddQuery("version", assistant.service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, new(DialogNode))
	return response, err
}

// GetGetDialogNodeResult : Retrieve result of GetDialogNode operation
func (assistant *AssistantV1) GetGetDialogNodeResult(response *core.DetailedResponse) *DialogNode {
	result, ok := response.Result.(*DialogNode)
	if ok {
		return result
	}
	return nil
}

// ListDialogNodes : List dialog nodes
func (assistant *AssistantV1) ListDialogNodes(listDialogNodesOptions *ListDialogNodesOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(listDialogNodesOptions, "listDialogNodesOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(listDialogNodesOptions, "listDialogNodesOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces", "dialog_nodes"}
	pathParameters := []string{*listDialogNodesOptions.WorkspaceID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range listDialogNodesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

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
	builder.AddQuery("version", assistant.service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, new(DialogNodeCollection))
	return response, err
}

// GetListDialogNodesResult : Retrieve result of ListDialogNodes operation
func (assistant *AssistantV1) GetListDialogNodesResult(response *core.DetailedResponse) *DialogNodeCollection {
	result, ok := response.Result.(*DialogNodeCollection)
	if ok {
		return result
	}
	return nil
}

// UpdateDialogNode : Update dialog node
func (assistant *AssistantV1) UpdateDialogNode(updateDialogNodeOptions *UpdateDialogNodeOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(updateDialogNodeOptions, "updateDialogNodeOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(updateDialogNodeOptions, "updateDialogNodeOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces", "dialog_nodes"}
	pathParameters := []string{*updateDialogNodeOptions.WorkspaceID, *updateDialogNodeOptions.DialogNode}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range updateDialogNodeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", assistant.service.Options.Version)

	body := make(map[string]interface{})
	if updateDialogNodeOptions.NodeType != nil {
		body["type"] = updateDialogNodeOptions.NodeType
	}
	if updateDialogNodeOptions.NewActions != nil {
		body["actions"] = updateDialogNodeOptions.NewActions
	}
	if updateDialogNodeOptions.NewConditions != nil {
		body["conditions"] = updateDialogNodeOptions.NewConditions
	}
	if updateDialogNodeOptions.NewContext != nil {
		body["context"] = updateDialogNodeOptions.NewContext
	}
	if updateDialogNodeOptions.NewPreviousSibling != nil {
		body["previous_sibling"] = updateDialogNodeOptions.NewPreviousSibling
	}
	if updateDialogNodeOptions.NewVariable != nil {
		body["variable"] = updateDialogNodeOptions.NewVariable
	}
	if updateDialogNodeOptions.NewUserLabel != nil {
		body["user_label"] = updateDialogNodeOptions.NewUserLabel
	}
	if updateDialogNodeOptions.NewMetadata != nil {
		body["metadata"] = updateDialogNodeOptions.NewMetadata
	}
	if updateDialogNodeOptions.NewTitle != nil {
		body["title"] = updateDialogNodeOptions.NewTitle
	}
	if updateDialogNodeOptions.NewDescription != nil {
		body["description"] = updateDialogNodeOptions.NewDescription
	}
	if updateDialogNodeOptions.NewDigressOut != nil {
		body["digress_out"] = updateDialogNodeOptions.NewDigressOut
	}
	if updateDialogNodeOptions.NewEventName != nil {
		body["event_name"] = updateDialogNodeOptions.NewEventName
	}
	if updateDialogNodeOptions.NewDigressOutSlots != nil {
		body["digress_out_slots"] = updateDialogNodeOptions.NewDigressOutSlots
	}
	if updateDialogNodeOptions.NewNextStep != nil {
		body["next_step"] = updateDialogNodeOptions.NewNextStep
	}
	if updateDialogNodeOptions.NewDigressIn != nil {
		body["digress_in"] = updateDialogNodeOptions.NewDigressIn
	}
	if updateDialogNodeOptions.NewOutput != nil {
		body["output"] = updateDialogNodeOptions.NewOutput
	}
	if updateDialogNodeOptions.NewParent != nil {
		body["parent"] = updateDialogNodeOptions.NewParent
	}
	if updateDialogNodeOptions.NewDialogNode != nil {
		body["dialog_node"] = updateDialogNodeOptions.NewDialogNode
	}
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, new(DialogNode))
	return response, err
}

// GetUpdateDialogNodeResult : Retrieve result of UpdateDialogNode operation
func (assistant *AssistantV1) GetUpdateDialogNodeResult(response *core.DetailedResponse) *DialogNode {
	result, ok := response.Result.(*DialogNode)
	if ok {
		return result
	}
	return nil
}

// ListAllLogs : List log events in all workspaces
func (assistant *AssistantV1) ListAllLogs(listAllLogsOptions *ListAllLogsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(listAllLogsOptions, "listAllLogsOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(listAllLogsOptions, "listAllLogsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/logs"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range listAllLogsOptions.Headers {
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
	builder.AddQuery("version", assistant.service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, new(LogCollection))
	return response, err
}

// GetListAllLogsResult : Retrieve result of ListAllLogs operation
func (assistant *AssistantV1) GetListAllLogsResult(response *core.DetailedResponse) *LogCollection {
	result, ok := response.Result.(*LogCollection)
	if ok {
		return result
	}
	return nil
}

// ListLogs : List log events in a workspace
func (assistant *AssistantV1) ListLogs(listLogsOptions *ListLogsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(listLogsOptions, "listLogsOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(listLogsOptions, "listLogsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/workspaces", "logs"}
	pathParameters := []string{*listLogsOptions.WorkspaceID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range listLogsOptions.Headers {
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
	builder.AddQuery("version", assistant.service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, new(LogCollection))
	return response, err
}

// GetListLogsResult : Retrieve result of ListLogs operation
func (assistant *AssistantV1) GetListLogsResult(response *core.DetailedResponse) *LogCollection {
	result, ok := response.Result.(*LogCollection)
	if ok {
		return result
	}
	return nil
}

// DeleteUserData : Delete labeled data
func (assistant *AssistantV1) DeleteUserData(deleteUserDataOptions *DeleteUserDataOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(deleteUserDataOptions, "deleteUserDataOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(deleteUserDataOptions, "deleteUserDataOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v1/user_data"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.DELETE)
	builder.ConstructHTTPURL(assistant.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteUserDataOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("customer_id", fmt.Sprint(*deleteUserDataOptions.CustomerID))
	builder.AddQuery("version", assistant.service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := assistant.service.Request(request, nil)
	return response, err
}

// CaptureGroup : CaptureGroup struct
type CaptureGroup struct {

	// A recognized capture group for the entity.
	Group *string `json:"group" validate:"required"`

	// Zero-based character offsets that indicate where the entity value begins and ends in the input text.
	Location []int64 `json:"location,omitempty"`
}

// Context : State information for the conversation. To maintain state, include the context from the previous response.
type Context struct {

	// The unique identifier of the conversation.
	ConversationID *string `json:"conversation_id,omitempty"`

	// For internal use only.
	System *SystemResponse `json:"system,omitempty"`
}

// Counterexample : Counterexample struct
type Counterexample struct {

	// The text of the counterexample.
	Text *string `json:"text" validate:"required"`

	// The timestamp for creation of the counterexample.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// The timestamp for the last update to the counterexample.
	Updated *strfmt.DateTime `json:"updated,omitempty"`
}

// CounterexampleCollection : CounterexampleCollection struct
type CounterexampleCollection struct {

	// An array of objects describing the examples marked as irrelevant input.
	Counterexamples []Counterexample `json:"counterexamples" validate:"required"`

	// The pagination data for the returned objects.
	Pagination *Pagination `json:"pagination" validate:"required"`
}

// CreateCounterexample : CreateCounterexample struct
type CreateCounterexample struct {

	// The text of a user input marked as irrelevant input. This string must conform to the following restrictions:
	// - It cannot contain carriage return, newline, or tab characters
	// - It cannot consist of only whitespace characters
	// - It must be no longer than 1024 characters.
	Text *string `json:"text" validate:"required"`
}

// CreateCounterexampleOptions : The createCounterexample options.
type CreateCounterexampleOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The text of a user input marked as irrelevant input. This string must conform to the following restrictions:
	// - It cannot contain carriage return, newline, or tab characters
	// - It cannot consist of only whitespace characters
	// - It must be no longer than 1024 characters.
	Text *string `json:"text" validate:"required"`

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
func (options *CreateCounterexampleOptions) SetWorkspaceID(param string) *CreateCounterexampleOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetText : Allow user to set Text
func (options *CreateCounterexampleOptions) SetText(param string) *CreateCounterexampleOptions {
	options.Text = core.StringPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateCounterexampleOptions) SetHeaders(param map[string]string) *CreateCounterexampleOptions {
	options.Headers = param
	return options
}

// CreateDialogNode : CreateDialogNode struct
type CreateDialogNode struct {

	// The dialog node ID. This string must conform to the following restrictions:
	// - It can contain only Unicode alphanumeric, space, underscore, hyphen, and dot characters.
	// - It must be no longer than 1024 characters.
	DialogNode *string `json:"dialog_node" validate:"required"`

	// The description of the dialog node. This string cannot contain carriage return, newline, or tab characters, and it
	// must be no longer than 128 characters.
	Description *string `json:"description,omitempty"`

	// The condition that will trigger the dialog node. This string cannot contain carriage return, newline, or tab
	// characters, and it must be no longer than 2048 characters.
	Conditions *string `json:"conditions,omitempty"`

	// The ID of the parent dialog node.
	Parent *string `json:"parent,omitempty"`

	// The ID of the previous dialog node.
	PreviousSibling *string `json:"previous_sibling,omitempty"`

	// The output of the dialog node. For more information about how to specify dialog node output, see the
	// [documentation](https://console.bluemix.net/docs/services/conversation/dialog-overview.html#complex).
	Output *DialogNodeOutput `json:"output,omitempty"`

	// The context for the dialog node.
	Context interface{} `json:"context,omitempty"`

	// The metadata for the dialog node.
	Metadata interface{} `json:"metadata,omitempty"`

	// The next step to be executed in dialog processing.
	NextStep *DialogNodeNextStep `json:"next_step,omitempty"`

	// An array of objects describing any actions to be invoked by the dialog node.
	Actions []DialogNodeAction `json:"actions,omitempty"`

	// The alias used to identify the dialog node. This string must conform to the following restrictions:
	// - It can contain only Unicode alphanumeric, space, underscore, hyphen, and dot characters.
	// - It must be no longer than 64 characters.
	Title *string `json:"title,omitempty"`

	// How the dialog node is processed.
	NodeType *string `json:"type,omitempty"`

	// How an `event_handler` node is processed.
	EventName *string `json:"event_name,omitempty"`

	// The location in the dialog context where output is stored.
	Variable *string `json:"variable,omitempty"`

	// Whether this top-level dialog node can be digressed into.
	DigressIn *string `json:"digress_in,omitempty"`

	// Whether this dialog node can be returned to after a digression.
	DigressOut *string `json:"digress_out,omitempty"`

	// Whether the user can digress to top-level nodes while filling out slots.
	DigressOutSlots *string `json:"digress_out_slots,omitempty"`

	// A label that can be displayed externally to describe the purpose of the node to users. This string must be no longer
	// than 512 characters.
	UserLabel *string `json:"user_label,omitempty"`
}

// CreateDialogNodeOptions : The createDialogNode options.
type CreateDialogNodeOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The dialog node ID. This string must conform to the following restrictions:
	// - It can contain only Unicode alphanumeric, space, underscore, hyphen, and dot characters.
	// - It must be no longer than 1024 characters.
	DialogNode *string `json:"dialog_node" validate:"required"`

	// The description of the dialog node. This string cannot contain carriage return, newline, or tab characters, and it
	// must be no longer than 128 characters.
	Description *string `json:"description,omitempty"`

	// The condition that will trigger the dialog node. This string cannot contain carriage return, newline, or tab
	// characters, and it must be no longer than 2048 characters.
	Conditions *string `json:"conditions,omitempty"`

	// The ID of the parent dialog node.
	Parent *string `json:"parent,omitempty"`

	// The ID of the previous dialog node.
	PreviousSibling *string `json:"previous_sibling,omitempty"`

	// The output of the dialog node. For more information about how to specify dialog node output, see the
	// [documentation](https://console.bluemix.net/docs/services/conversation/dialog-overview.html#complex).
	Output *DialogNodeOutput `json:"output,omitempty"`

	// The context for the dialog node.
	Context interface{} `json:"context,omitempty"`

	// The metadata for the dialog node.
	Metadata interface{} `json:"metadata,omitempty"`

	// The next step to be executed in dialog processing.
	NextStep *DialogNodeNextStep `json:"next_step,omitempty"`

	// An array of objects describing any actions to be invoked by the dialog node.
	Actions []DialogNodeAction `json:"actions,omitempty"`

	// The alias used to identify the dialog node. This string must conform to the following restrictions:
	// - It can contain only Unicode alphanumeric, space, underscore, hyphen, and dot characters.
	// - It must be no longer than 64 characters.
	Title *string `json:"title,omitempty"`

	// How the dialog node is processed.
	NodeType *string `json:"type,omitempty"`

	// How an `event_handler` node is processed.
	EventName *string `json:"event_name,omitempty"`

	// The location in the dialog context where output is stored.
	Variable *string `json:"variable,omitempty"`

	// Whether this top-level dialog node can be digressed into.
	DigressIn *string `json:"digress_in,omitempty"`

	// Whether this dialog node can be returned to after a digression.
	DigressOut *string `json:"digress_out,omitempty"`

	// Whether the user can digress to top-level nodes while filling out slots.
	DigressOutSlots *string `json:"digress_out_slots,omitempty"`

	// A label that can be displayed externally to describe the purpose of the node to users. This string must be no longer
	// than 512 characters.
	UserLabel *string `json:"user_label,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewCreateDialogNodeOptions : Instantiate CreateDialogNodeOptions
func (assistant *AssistantV1) NewCreateDialogNodeOptions(workspaceID string, dialogNode string) *CreateDialogNodeOptions {
	return &CreateDialogNodeOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		DialogNode:  core.StringPtr(dialogNode),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *CreateDialogNodeOptions) SetWorkspaceID(param string) *CreateDialogNodeOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetDialogNode : Allow user to set DialogNode
func (options *CreateDialogNodeOptions) SetDialogNode(param string) *CreateDialogNodeOptions {
	options.DialogNode = core.StringPtr(param)
	return options
}

// SetDescription : Allow user to set Description
func (options *CreateDialogNodeOptions) SetDescription(param string) *CreateDialogNodeOptions {
	options.Description = core.StringPtr(param)
	return options
}

// SetConditions : Allow user to set Conditions
func (options *CreateDialogNodeOptions) SetConditions(param string) *CreateDialogNodeOptions {
	options.Conditions = core.StringPtr(param)
	return options
}

// SetParent : Allow user to set Parent
func (options *CreateDialogNodeOptions) SetParent(param string) *CreateDialogNodeOptions {
	options.Parent = core.StringPtr(param)
	return options
}

// SetPreviousSibling : Allow user to set PreviousSibling
func (options *CreateDialogNodeOptions) SetPreviousSibling(param string) *CreateDialogNodeOptions {
	options.PreviousSibling = core.StringPtr(param)
	return options
}

// SetOutput : Allow user to set Output
func (options *CreateDialogNodeOptions) SetOutput(param DialogNodeOutput) *CreateDialogNodeOptions {
	options.Output = &param
	return options
}

// SetContext : Allow user to set Context
func (options *CreateDialogNodeOptions) SetContext(param interface{}) *CreateDialogNodeOptions {
	options.Context = &param
	return options
}

// SetMetadata : Allow user to set Metadata
func (options *CreateDialogNodeOptions) SetMetadata(param interface{}) *CreateDialogNodeOptions {
	options.Metadata = &param
	return options
}

// SetNextStep : Allow user to set NextStep
func (options *CreateDialogNodeOptions) SetNextStep(param DialogNodeNextStep) *CreateDialogNodeOptions {
	options.NextStep = &param
	return options
}

// SetActions : Allow user to set Actions
func (options *CreateDialogNodeOptions) SetActions(param []DialogNodeAction) *CreateDialogNodeOptions {
	options.Actions = param
	return options
}

// SetTitle : Allow user to set Title
func (options *CreateDialogNodeOptions) SetTitle(param string) *CreateDialogNodeOptions {
	options.Title = core.StringPtr(param)
	return options
}

// SetNodeType : Allow user to set NodeType
func (options *CreateDialogNodeOptions) SetNodeType(param string) *CreateDialogNodeOptions {
	options.NodeType = core.StringPtr(param)
	return options
}

// SetEventName : Allow user to set EventName
func (options *CreateDialogNodeOptions) SetEventName(param string) *CreateDialogNodeOptions {
	options.EventName = core.StringPtr(param)
	return options
}

// SetVariable : Allow user to set Variable
func (options *CreateDialogNodeOptions) SetVariable(param string) *CreateDialogNodeOptions {
	options.Variable = core.StringPtr(param)
	return options
}

// SetDigressIn : Allow user to set DigressIn
func (options *CreateDialogNodeOptions) SetDigressIn(param string) *CreateDialogNodeOptions {
	options.DigressIn = core.StringPtr(param)
	return options
}

// SetDigressOut : Allow user to set DigressOut
func (options *CreateDialogNodeOptions) SetDigressOut(param string) *CreateDialogNodeOptions {
	options.DigressOut = core.StringPtr(param)
	return options
}

// SetDigressOutSlots : Allow user to set DigressOutSlots
func (options *CreateDialogNodeOptions) SetDigressOutSlots(param string) *CreateDialogNodeOptions {
	options.DigressOutSlots = core.StringPtr(param)
	return options
}

// SetUserLabel : Allow user to set UserLabel
func (options *CreateDialogNodeOptions) SetUserLabel(param string) *CreateDialogNodeOptions {
	options.UserLabel = core.StringPtr(param)
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
	// - It cannot begin with the reserved prefix `sys-`.
	// - It must be no longer than 64 characters.
	Entity *string `json:"entity" validate:"required"`

	// The description of the entity. This string cannot contain carriage return, newline, or tab characters, and it must
	// be no longer than 128 characters.
	Description *string `json:"description,omitempty"`

	// Any metadata related to the value.
	Metadata interface{} `json:"metadata,omitempty"`

	// An array of objects describing the entity values.
	Values []CreateValue `json:"values,omitempty"`

	// Whether to use fuzzy matching for the entity.
	FuzzyMatch *bool `json:"fuzzy_match,omitempty"`
}

// CreateEntityOptions : The createEntity options.
type CreateEntityOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The name of the entity. This string must conform to the following restrictions:
	// - It can contain only Unicode alphanumeric, underscore, and hyphen characters.
	// - It cannot begin with the reserved prefix `sys-`.
	// - It must be no longer than 64 characters.
	Entity *string `json:"entity" validate:"required"`

	// The description of the entity. This string cannot contain carriage return, newline, or tab characters, and it must
	// be no longer than 128 characters.
	Description *string `json:"description,omitempty"`

	// Any metadata related to the value.
	Metadata interface{} `json:"metadata,omitempty"`

	// An array of objects describing the entity values.
	Values []CreateValue `json:"values,omitempty"`

	// Whether to use fuzzy matching for the entity.
	FuzzyMatch *bool `json:"fuzzy_match,omitempty"`

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
func (options *CreateEntityOptions) SetWorkspaceID(param string) *CreateEntityOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetEntity : Allow user to set Entity
func (options *CreateEntityOptions) SetEntity(param string) *CreateEntityOptions {
	options.Entity = core.StringPtr(param)
	return options
}

// SetDescription : Allow user to set Description
func (options *CreateEntityOptions) SetDescription(param string) *CreateEntityOptions {
	options.Description = core.StringPtr(param)
	return options
}

// SetMetadata : Allow user to set Metadata
func (options *CreateEntityOptions) SetMetadata(param interface{}) *CreateEntityOptions {
	options.Metadata = &param
	return options
}

// SetValues : Allow user to set Values
func (options *CreateEntityOptions) SetValues(param []CreateValue) *CreateEntityOptions {
	options.Values = param
	return options
}

// SetFuzzyMatch : Allow user to set FuzzyMatch
func (options *CreateEntityOptions) SetFuzzyMatch(param bool) *CreateEntityOptions {
	options.FuzzyMatch = core.BoolPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateEntityOptions) SetHeaders(param map[string]string) *CreateEntityOptions {
	options.Headers = param
	return options
}

// CreateExample : CreateExample struct
type CreateExample struct {

	// The text of a user input example. This string must conform to the following restrictions:
	// - It cannot contain carriage return, newline, or tab characters.
	// - It cannot consist of only whitespace characters.
	// - It must be no longer than 1024 characters.
	Text *string `json:"text" validate:"required"`

	// An array of contextual entity mentions.
	Mentions []Mentions `json:"mentions,omitempty"`
}

// CreateExampleOptions : The createExample options.
type CreateExampleOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The intent name.
	Intent *string `json:"intent" validate:"required"`

	// The text of a user input example. This string must conform to the following restrictions:
	// - It cannot contain carriage return, newline, or tab characters.
	// - It cannot consist of only whitespace characters.
	// - It must be no longer than 1024 characters.
	Text *string `json:"text" validate:"required"`

	// An array of contextual entity mentions.
	Mentions []Mentions `json:"mentions,omitempty"`

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
func (options *CreateExampleOptions) SetWorkspaceID(param string) *CreateExampleOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetIntent : Allow user to set Intent
func (options *CreateExampleOptions) SetIntent(param string) *CreateExampleOptions {
	options.Intent = core.StringPtr(param)
	return options
}

// SetText : Allow user to set Text
func (options *CreateExampleOptions) SetText(param string) *CreateExampleOptions {
	options.Text = core.StringPtr(param)
	return options
}

// SetMentions : Allow user to set Mentions
func (options *CreateExampleOptions) SetMentions(param []Mentions) *CreateExampleOptions {
	options.Mentions = param
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
	// - It must be no longer than 128 characters.
	Intent *string `json:"intent" validate:"required"`

	// The description of the intent. This string cannot contain carriage return, newline, or tab characters, and it must
	// be no longer than 128 characters.
	Description *string `json:"description,omitempty"`

	// An array of user input examples for the intent.
	Examples []CreateExample `json:"examples,omitempty"`
}

// CreateIntentOptions : The createIntent options.
type CreateIntentOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The name of the intent. This string must conform to the following restrictions:
	// - It can contain only Unicode alphanumeric, underscore, hyphen, and dot characters.
	// - It cannot begin with the reserved prefix `sys-`.
	// - It must be no longer than 128 characters.
	Intent *string `json:"intent" validate:"required"`

	// The description of the intent. This string cannot contain carriage return, newline, or tab characters, and it must
	// be no longer than 128 characters.
	Description *string `json:"description,omitempty"`

	// An array of user input examples for the intent.
	Examples []CreateExample `json:"examples,omitempty"`

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
func (options *CreateIntentOptions) SetWorkspaceID(param string) *CreateIntentOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetIntent : Allow user to set Intent
func (options *CreateIntentOptions) SetIntent(param string) *CreateIntentOptions {
	options.Intent = core.StringPtr(param)
	return options
}

// SetDescription : Allow user to set Description
func (options *CreateIntentOptions) SetDescription(param string) *CreateIntentOptions {
	options.Description = core.StringPtr(param)
	return options
}

// SetExamples : Allow user to set Examples
func (options *CreateIntentOptions) SetExamples(param []CreateExample) *CreateIntentOptions {
	options.Examples = param
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateIntentOptions) SetHeaders(param map[string]string) *CreateIntentOptions {
	options.Headers = param
	return options
}

// CreateSynonymOptions : The createSynonym options.
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
	// - It must be no longer than 64 characters.
	Synonym *string `json:"synonym" validate:"required"`

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
func (options *CreateSynonymOptions) SetWorkspaceID(param string) *CreateSynonymOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetEntity : Allow user to set Entity
func (options *CreateSynonymOptions) SetEntity(param string) *CreateSynonymOptions {
	options.Entity = core.StringPtr(param)
	return options
}

// SetValue : Allow user to set Value
func (options *CreateSynonymOptions) SetValue(param string) *CreateSynonymOptions {
	options.Value = core.StringPtr(param)
	return options
}

// SetSynonym : Allow user to set Synonym
func (options *CreateSynonymOptions) SetSynonym(param string) *CreateSynonymOptions {
	options.Synonym = core.StringPtr(param)
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
	// - It must be no longer than 64 characters.
	Value *string `json:"value" validate:"required"`

	// Any metadata related to the entity value.
	Metadata interface{} `json:"metadata,omitempty"`

	// An array containing any synonyms for the entity value. You can provide either synonyms or patterns (as indicated by
	// **type**), but not both. A synonym must conform to the following restrictions:
	// - It cannot contain carriage return, newline, or tab characters.
	// - It cannot consist of only whitespace characters.
	// - It must be no longer than 64 characters.
	Synonyms []string `json:"synonyms,omitempty"`

	// An array of patterns for the entity value. You can provide either synonyms or patterns (as indicated by **type**),
	// but not both. A pattern is a regular expression no longer than 512 characters. For more information about how to
	// specify a pattern, see the
	// [documentation](https://console.bluemix.net/docs/services/conversation/entities.html#creating-entities).
	Patterns []string `json:"patterns,omitempty"`

	// Specifies the type of value.
	ValueType *string `json:"type,omitempty"`
}

// CreateValueOptions : The createValue options.
type CreateValueOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The name of the entity.
	Entity *string `json:"entity" validate:"required"`

	// The text of the entity value. This string must conform to the following restrictions:
	// - It cannot contain carriage return, newline, or tab characters.
	// - It cannot consist of only whitespace characters.
	// - It must be no longer than 64 characters.
	Value *string `json:"value" validate:"required"`

	// Any metadata related to the entity value.
	Metadata interface{} `json:"metadata,omitempty"`

	// An array containing any synonyms for the entity value. You can provide either synonyms or patterns (as indicated by
	// **type**), but not both. A synonym must conform to the following restrictions:
	// - It cannot contain carriage return, newline, or tab characters.
	// - It cannot consist of only whitespace characters.
	// - It must be no longer than 64 characters.
	Synonyms []string `json:"synonyms,omitempty"`

	// An array of patterns for the entity value. You can provide either synonyms or patterns (as indicated by **type**),
	// but not both. A pattern is a regular expression no longer than 512 characters. For more information about how to
	// specify a pattern, see the
	// [documentation](https://console.bluemix.net/docs/services/conversation/entities.html#creating-entities).
	Patterns []string `json:"patterns,omitempty"`

	// Specifies the type of value.
	ValueType *string `json:"type,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewCreateValueOptions : Instantiate CreateValueOptions
func (assistant *AssistantV1) NewCreateValueOptions(workspaceID string, entity string, value string) *CreateValueOptions {
	return &CreateValueOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Entity:      core.StringPtr(entity),
		Value:       core.StringPtr(value),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *CreateValueOptions) SetWorkspaceID(param string) *CreateValueOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetEntity : Allow user to set Entity
func (options *CreateValueOptions) SetEntity(param string) *CreateValueOptions {
	options.Entity = core.StringPtr(param)
	return options
}

// SetValue : Allow user to set Value
func (options *CreateValueOptions) SetValue(param string) *CreateValueOptions {
	options.Value = core.StringPtr(param)
	return options
}

// SetMetadata : Allow user to set Metadata
func (options *CreateValueOptions) SetMetadata(param interface{}) *CreateValueOptions {
	options.Metadata = &param
	return options
}

// SetSynonyms : Allow user to set Synonyms
func (options *CreateValueOptions) SetSynonyms(param []string) *CreateValueOptions {
	options.Synonyms = param
	return options
}

// SetPatterns : Allow user to set Patterns
func (options *CreateValueOptions) SetPatterns(param []string) *CreateValueOptions {
	options.Patterns = param
	return options
}

// SetValueType : Allow user to set ValueType
func (options *CreateValueOptions) SetValueType(param string) *CreateValueOptions {
	options.ValueType = core.StringPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateValueOptions) SetHeaders(param map[string]string) *CreateValueOptions {
	options.Headers = param
	return options
}

// CreateWorkspaceOptions : The createWorkspace options.
type CreateWorkspaceOptions struct {

	// The name of the workspace. This string cannot contain carriage return, newline, or tab characters, and it must be no
	// longer than 64 characters.
	Name *string `json:"name,omitempty"`

	// The description of the workspace. This string cannot contain carriage return, newline, or tab characters, and it
	// must be no longer than 128 characters.
	Description *string `json:"description,omitempty"`

	// The language of the workspace.
	Language *string `json:"language,omitempty"`

	// An array of objects defining the intents for the workspace.
	Intents []CreateIntent `json:"intents,omitempty"`

	// An array of objects defining the entities for the workspace.
	Entities []CreateEntity `json:"entities,omitempty"`

	// An array of objects defining the nodes in the workspace dialog.
	DialogNodes []CreateDialogNode `json:"dialog_nodes,omitempty"`

	// An array of objects defining input examples that have been marked as irrelevant input.
	Counterexamples []CreateCounterexample `json:"counterexamples,omitempty"`

	// Any metadata related to the workspace.
	Metadata interface{} `json:"metadata,omitempty"`

	// Whether training data from the workspace can be used by IBM for general service improvements. `true` indicates that
	// workspace training data is not to be used.
	LearningOptOut *bool `json:"learning_opt_out,omitempty"`

	// Global settings for the workspace.
	SystemSettings *WorkspaceSystemSettings `json:"system_settings,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewCreateWorkspaceOptions : Instantiate CreateWorkspaceOptions
func (assistant *AssistantV1) NewCreateWorkspaceOptions() *CreateWorkspaceOptions {
	return &CreateWorkspaceOptions{}
}

// SetName : Allow user to set Name
func (options *CreateWorkspaceOptions) SetName(param string) *CreateWorkspaceOptions {
	options.Name = core.StringPtr(param)
	return options
}

// SetDescription : Allow user to set Description
func (options *CreateWorkspaceOptions) SetDescription(param string) *CreateWorkspaceOptions {
	options.Description = core.StringPtr(param)
	return options
}

// SetLanguage : Allow user to set Language
func (options *CreateWorkspaceOptions) SetLanguage(param string) *CreateWorkspaceOptions {
	options.Language = core.StringPtr(param)
	return options
}

// SetIntents : Allow user to set Intents
func (options *CreateWorkspaceOptions) SetIntents(param []CreateIntent) *CreateWorkspaceOptions {
	options.Intents = param
	return options
}

// SetEntities : Allow user to set Entities
func (options *CreateWorkspaceOptions) SetEntities(param []CreateEntity) *CreateWorkspaceOptions {
	options.Entities = param
	return options
}

// SetDialogNodes : Allow user to set DialogNodes
func (options *CreateWorkspaceOptions) SetDialogNodes(param []CreateDialogNode) *CreateWorkspaceOptions {
	options.DialogNodes = param
	return options
}

// SetCounterexamples : Allow user to set Counterexamples
func (options *CreateWorkspaceOptions) SetCounterexamples(param []CreateCounterexample) *CreateWorkspaceOptions {
	options.Counterexamples = param
	return options
}

// SetMetadata : Allow user to set Metadata
func (options *CreateWorkspaceOptions) SetMetadata(param interface{}) *CreateWorkspaceOptions {
	options.Metadata = &param
	return options
}

// SetLearningOptOut : Allow user to set LearningOptOut
func (options *CreateWorkspaceOptions) SetLearningOptOut(param bool) *CreateWorkspaceOptions {
	options.LearningOptOut = core.BoolPtr(param)
	return options
}

// SetSystemSettings : Allow user to set SystemSettings
func (options *CreateWorkspaceOptions) SetSystemSettings(param WorkspaceSystemSettings) *CreateWorkspaceOptions {
	options.SystemSettings = &param
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateWorkspaceOptions) SetHeaders(param map[string]string) *CreateWorkspaceOptions {
	options.Headers = param
	return options
}

// DeleteCounterexampleOptions : The deleteCounterexample options.
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
func (options *DeleteCounterexampleOptions) SetWorkspaceID(param string) *DeleteCounterexampleOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetText : Allow user to set Text
func (options *DeleteCounterexampleOptions) SetText(param string) *DeleteCounterexampleOptions {
	options.Text = core.StringPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteCounterexampleOptions) SetHeaders(param map[string]string) *DeleteCounterexampleOptions {
	options.Headers = param
	return options
}

// DeleteDialogNodeOptions : The deleteDialogNode options.
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
func (options *DeleteDialogNodeOptions) SetWorkspaceID(param string) *DeleteDialogNodeOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetDialogNode : Allow user to set DialogNode
func (options *DeleteDialogNodeOptions) SetDialogNode(param string) *DeleteDialogNodeOptions {
	options.DialogNode = core.StringPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteDialogNodeOptions) SetHeaders(param map[string]string) *DeleteDialogNodeOptions {
	options.Headers = param
	return options
}

// DeleteEntityOptions : The deleteEntity options.
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
func (options *DeleteEntityOptions) SetWorkspaceID(param string) *DeleteEntityOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetEntity : Allow user to set Entity
func (options *DeleteEntityOptions) SetEntity(param string) *DeleteEntityOptions {
	options.Entity = core.StringPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteEntityOptions) SetHeaders(param map[string]string) *DeleteEntityOptions {
	options.Headers = param
	return options
}

// DeleteExampleOptions : The deleteExample options.
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
func (options *DeleteExampleOptions) SetWorkspaceID(param string) *DeleteExampleOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetIntent : Allow user to set Intent
func (options *DeleteExampleOptions) SetIntent(param string) *DeleteExampleOptions {
	options.Intent = core.StringPtr(param)
	return options
}

// SetText : Allow user to set Text
func (options *DeleteExampleOptions) SetText(param string) *DeleteExampleOptions {
	options.Text = core.StringPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteExampleOptions) SetHeaders(param map[string]string) *DeleteExampleOptions {
	options.Headers = param
	return options
}

// DeleteIntentOptions : The deleteIntent options.
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
func (options *DeleteIntentOptions) SetWorkspaceID(param string) *DeleteIntentOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetIntent : Allow user to set Intent
func (options *DeleteIntentOptions) SetIntent(param string) *DeleteIntentOptions {
	options.Intent = core.StringPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteIntentOptions) SetHeaders(param map[string]string) *DeleteIntentOptions {
	options.Headers = param
	return options
}

// DeleteSynonymOptions : The deleteSynonym options.
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
func (options *DeleteSynonymOptions) SetWorkspaceID(param string) *DeleteSynonymOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetEntity : Allow user to set Entity
func (options *DeleteSynonymOptions) SetEntity(param string) *DeleteSynonymOptions {
	options.Entity = core.StringPtr(param)
	return options
}

// SetValue : Allow user to set Value
func (options *DeleteSynonymOptions) SetValue(param string) *DeleteSynonymOptions {
	options.Value = core.StringPtr(param)
	return options
}

// SetSynonym : Allow user to set Synonym
func (options *DeleteSynonymOptions) SetSynonym(param string) *DeleteSynonymOptions {
	options.Synonym = core.StringPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteSynonymOptions) SetHeaders(param map[string]string) *DeleteSynonymOptions {
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
func (assistant *AssistantV1) NewDeleteUserDataOptions(customerID string) *DeleteUserDataOptions {
	return &DeleteUserDataOptions{
		CustomerID: core.StringPtr(customerID),
	}
}

// SetCustomerID : Allow user to set CustomerID
func (options *DeleteUserDataOptions) SetCustomerID(param string) *DeleteUserDataOptions {
	options.CustomerID = core.StringPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteUserDataOptions) SetHeaders(param map[string]string) *DeleteUserDataOptions {
	options.Headers = param
	return options
}

// DeleteValueOptions : The deleteValue options.
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
func (options *DeleteValueOptions) SetWorkspaceID(param string) *DeleteValueOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetEntity : Allow user to set Entity
func (options *DeleteValueOptions) SetEntity(param string) *DeleteValueOptions {
	options.Entity = core.StringPtr(param)
	return options
}

// SetValue : Allow user to set Value
func (options *DeleteValueOptions) SetValue(param string) *DeleteValueOptions {
	options.Value = core.StringPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteValueOptions) SetHeaders(param map[string]string) *DeleteValueOptions {
	options.Headers = param
	return options
}

// DeleteWorkspaceOptions : The deleteWorkspace options.
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
func (options *DeleteWorkspaceOptions) SetWorkspaceID(param string) *DeleteWorkspaceOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteWorkspaceOptions) SetHeaders(param map[string]string) *DeleteWorkspaceOptions {
	options.Headers = param
	return options
}

// DialogNode : DialogNode struct
type DialogNode struct {

	// The dialog node ID.
	DialogNodeID *string `json:"dialog_node" validate:"required"`

	// The description of the dialog node.
	Description *string `json:"description,omitempty"`

	// The condition that triggers the dialog node.
	Conditions *string `json:"conditions,omitempty"`

	// The ID of the parent dialog node. This property is not returned if the dialog node has no parent.
	Parent *string `json:"parent,omitempty"`

	// The ID of the previous sibling dialog node. This property is not returned if the dialog node has no previous
	// sibling.
	PreviousSibling *string `json:"previous_sibling,omitempty"`

	// The output of the dialog node. For more information about how to specify dialog node output, see the
	// [documentation](https://console.bluemix.net/docs/services/conversation/dialog-overview.html#complex).
	Output *DialogNodeOutput `json:"output,omitempty"`

	// The context (if defined) for the dialog node.
	Context interface{} `json:"context,omitempty"`

	// Any metadata for the dialog node.
	Metadata interface{} `json:"metadata,omitempty"`

	// The next step to execute following this dialog node.
	NextStep *DialogNodeNextStep `json:"next_step,omitempty"`

	// The timestamp for creation of the dialog node.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// The timestamp for the most recent update to the dialog node.
	Updated *strfmt.DateTime `json:"updated,omitempty"`

	// The actions for the dialog node.
	Actions []DialogNodeAction `json:"actions,omitempty"`

	// The alias used to identify the dialog node.
	Title *string `json:"title,omitempty"`

	// How the dialog node is processed.
	NodeType *string `json:"type,omitempty"`

	// How an `event_handler` node is processed.
	EventName *string `json:"event_name,omitempty"`

	// The location in the dialog context where output is stored.
	Variable *string `json:"variable,omitempty"`

	// Whether this top-level dialog node can be digressed into.
	DigressIn *string `json:"digress_in,omitempty"`

	// Whether this dialog node can be returned to after a digression.
	DigressOut *string `json:"digress_out,omitempty"`

	// Whether the user can digress to top-level nodes while filling out slots.
	DigressOutSlots *string `json:"digress_out_slots,omitempty"`

	// A label that can be displayed externally to describe the purpose of the node to users. This string must be no longer
	// than 512 characters.
	UserLabel *string `json:"user_label,omitempty"`
}

// DialogNodeAction : DialogNodeAction struct
type DialogNodeAction struct {

	// The name of the action.
	Name *string `json:"name" validate:"required"`

	// The type of action to invoke.
	ActionType *string `json:"type,omitempty"`

	// A map of key/value pairs to be provided to the action.
	Parameters interface{} `json:"parameters,omitempty"`

	// The location in the dialog context where the result of the action is stored.
	ResultVariable *string `json:"result_variable" validate:"required"`

	// The name of the context variable that the client application will use to pass in credentials for the action.
	Credentials *string `json:"credentials,omitempty"`
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
	//
	// If you specify `jump_to`, then you must also specify a value for the `dialog_node` property.
	Behavior *string `json:"behavior" validate:"required"`

	// The ID of the dialog node to process next. This parameter is required if **behavior**=`jump_to`.
	DialogNode *string `json:"dialog_node,omitempty"`

	// Which part of the dialog node to process next.
	Selector *string `json:"selector,omitempty"`
}

// DialogNodeOutput : The output of the dialog node. For more information about how to specify dialog node output, see the
// [documentation](https://console.bluemix.net/docs/services/conversation/dialog-overview.html#complex).
type DialogNodeOutput struct {

	// An array of objects describing the output defined for the dialog node.
	Generic []DialogNodeOutputGeneric `json:"generic,omitempty"`

	// Options that modify how specified output is handled.
	Modifiers *DialogNodeOutputModifiers `json:"modifiers,omitempty"`
}

// DialogNodeOutputGeneric : DialogNodeOutputGeneric struct
type DialogNodeOutputGeneric struct {

	// The type of response returned by the dialog node. The specified response type must be supported by the client
	// application or channel.
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

	// An optional title to show before the response. Valid only when **response_type**=`image` or `option`. This string
	// must be no longer than 512 characters.
	Title *string `json:"title,omitempty"`

	// An optional description to show with the response. Valid only when **response_type**=`image` or `option`. This
	// string must be no longer than 256 characters.
	Description *string `json:"description,omitempty"`

	// The preferred type of control to display, if supported by the channel. Valid only when **response_type**=`option`.
	Preference *string `json:"preference,omitempty"`

	// An array of objects describing the options from which the user can choose. You can include up to 20 options.
	// Required when **response_type**=`option`.
	Options []DialogNodeOutputOptionsElement `json:"options,omitempty"`

	// An optional message to be sent to the human agent who will be taking over the conversation. Valid only when
	// **reponse_type**=`connect_to_agent`. This string must be no longer than 256 characters.
	MessageToHumanAgent *string `json:"message_to_human_agent,omitempty"`
}

// DialogNodeOutputModifiers : Options that modify how specified output is handled.
type DialogNodeOutputModifiers struct {

	// Whether values in the output will overwrite output values in an array specified by previously executed dialog nodes.
	// If this option is set to **false**, new values will be appended to previously specified values.
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

// DialogNodeOutputOptionsElementValue : An object defining the message input to be sent to the Watson Assistant service if the user selects the corresponding
// option.
type DialogNodeOutputOptionsElementValue struct {

	// The user input.
	Input *InputData `json:"input,omitempty"`
}

// DialogNodeOutputTextValuesElement : DialogNodeOutputTextValuesElement struct
type DialogNodeOutputTextValuesElement struct {

	// The text of a response. This string can include newline characters (`
	// `), Markdown tagging, or other special characters, if supported by the channel. It must be no longer than 4096
	// characters.
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

// DialogRuntimeResponseGeneric : DialogRuntimeResponseGeneric struct
type DialogRuntimeResponseGeneric struct {

	// The type of response returned by the dialog node. The specified response type must be supported by the client
	// application or channel.
	//
	// **Note:** The **suggestion** response type is part of the disambiguation feature, which is only available for
	// Premium users.
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

	// A label identifying the topic of the conversation, derived from the **user_label** property of the relevant node.
	Topic *string `json:"topic,omitempty"`

	// An array of objects describing the possible matching dialog nodes from which the user can choose.
	//
	// **Note:** The **suggestions** property is part of the disambiguation feature, which is only available for Premium
	// users.
	Suggestions []DialogSuggestion `json:"suggestions,omitempty"`
}

// DialogSuggestion : DialogSuggestion struct
type DialogSuggestion struct {

	// The user-facing label for the disambiguation option. This label is taken from the **user_label** property of the
	// corresponding dialog node.
	Label *string `json:"label" validate:"required"`

	// An object defining the message input, intents, and entities to be sent to the Watson Assistant service if the user
	// selects the corresponding disambiguation option.
	Value *DialogSuggestionValue `json:"value" validate:"required"`

	// The dialog output that will be returned from the Watson Assistant service if the user selects the corresponding
	// option.
	Output interface{} `json:"output,omitempty"`
}

// DialogSuggestionValue : An object defining the message input, intents, and entities to be sent to the Watson Assistant service if the user
// selects the corresponding disambiguation option.
type DialogSuggestionValue struct {

	// The user input.
	Input *InputData `json:"input,omitempty"`

	// An array of intents to be sent along with the user input.
	Intents []RuntimeIntent `json:"intents,omitempty"`

	// An array of entities to be sent along with the user input.
	Entities []RuntimeEntity `json:"entities,omitempty"`
}

// Entity : Entity struct
type Entity struct {

	// The name of the entity.
	EntityName *string `json:"entity" validate:"required"`

	// The timestamp for creation of the entity.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// The timestamp for the last update to the entity.
	Updated *strfmt.DateTime `json:"updated,omitempty"`

	// The description of the entity.
	Description *string `json:"description,omitempty"`

	// Any metadata related to the entity.
	Metadata interface{} `json:"metadata,omitempty"`

	// Whether fuzzy matching is used for the entity.
	FuzzyMatch *bool `json:"fuzzy_match,omitempty"`
}

// EntityCollection : An array of entities.
type EntityCollection struct {

	// An array of objects describing the entities defined for the workspace.
	Entities []EntityExport `json:"entities" validate:"required"`

	// The pagination data for the returned objects.
	Pagination *Pagination `json:"pagination" validate:"required"`
}

// EntityExport : EntityExport struct
type EntityExport struct {

	// The name of the entity.
	EntityName *string `json:"entity" validate:"required"`

	// The timestamp for creation of the entity.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// The timestamp for the last update to the entity.
	Updated *strfmt.DateTime `json:"updated,omitempty"`

	// The description of the entity.
	Description *string `json:"description,omitempty"`

	// Any metadata related to the entity.
	Metadata interface{} `json:"metadata,omitempty"`

	// Whether fuzzy matching is used for the entity.
	FuzzyMatch *bool `json:"fuzzy_match,omitempty"`

	// An array objects describing the entity values.
	Values []ValueExport `json:"values,omitempty"`
}

// EntityMention : An object describing a contextual entity mention.
type EntityMention struct {

	// The text of the user input example.
	ExampleText *string `json:"text" validate:"required"`

	// The name of the intent.
	IntentName *string `json:"intent" validate:"required"`

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

	// The text of the user input example.
	ExampleText *string `json:"text" validate:"required"`

	// The timestamp for creation of the example.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// The timestamp for the last update to the example.
	Updated *strfmt.DateTime `json:"updated,omitempty"`

	// An array of contextual entity mentions.
	Mentions []Mentions `json:"mentions,omitempty"`
}

// ExampleCollection : ExampleCollection struct
type ExampleCollection struct {

	// An array of objects describing the examples defined for the intent.
	Examples []Example `json:"examples" validate:"required"`

	// The pagination data for the returned objects.
	Pagination *Pagination `json:"pagination" validate:"required"`
}

// GetCounterexampleOptions : The getCounterexample options.
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
func (options *GetCounterexampleOptions) SetWorkspaceID(param string) *GetCounterexampleOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetText : Allow user to set Text
func (options *GetCounterexampleOptions) SetText(param string) *GetCounterexampleOptions {
	options.Text = core.StringPtr(param)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *GetCounterexampleOptions) SetIncludeAudit(param bool) *GetCounterexampleOptions {
	options.IncludeAudit = core.BoolPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetCounterexampleOptions) SetHeaders(param map[string]string) *GetCounterexampleOptions {
	options.Headers = param
	return options
}

// GetDialogNodeOptions : The getDialogNode options.
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
func (options *GetDialogNodeOptions) SetWorkspaceID(param string) *GetDialogNodeOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetDialogNode : Allow user to set DialogNode
func (options *GetDialogNodeOptions) SetDialogNode(param string) *GetDialogNodeOptions {
	options.DialogNode = core.StringPtr(param)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *GetDialogNodeOptions) SetIncludeAudit(param bool) *GetDialogNodeOptions {
	options.IncludeAudit = core.BoolPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetDialogNodeOptions) SetHeaders(param map[string]string) *GetDialogNodeOptions {
	options.Headers = param
	return options
}

// GetEntityOptions : The getEntity options.
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
func (options *GetEntityOptions) SetWorkspaceID(param string) *GetEntityOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetEntity : Allow user to set Entity
func (options *GetEntityOptions) SetEntity(param string) *GetEntityOptions {
	options.Entity = core.StringPtr(param)
	return options
}

// SetExport : Allow user to set Export
func (options *GetEntityOptions) SetExport(param bool) *GetEntityOptions {
	options.Export = core.BoolPtr(param)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *GetEntityOptions) SetIncludeAudit(param bool) *GetEntityOptions {
	options.IncludeAudit = core.BoolPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetEntityOptions) SetHeaders(param map[string]string) *GetEntityOptions {
	options.Headers = param
	return options
}

// GetExampleOptions : The getExample options.
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
func (options *GetExampleOptions) SetWorkspaceID(param string) *GetExampleOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetIntent : Allow user to set Intent
func (options *GetExampleOptions) SetIntent(param string) *GetExampleOptions {
	options.Intent = core.StringPtr(param)
	return options
}

// SetText : Allow user to set Text
func (options *GetExampleOptions) SetText(param string) *GetExampleOptions {
	options.Text = core.StringPtr(param)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *GetExampleOptions) SetIncludeAudit(param bool) *GetExampleOptions {
	options.IncludeAudit = core.BoolPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetExampleOptions) SetHeaders(param map[string]string) *GetExampleOptions {
	options.Headers = param
	return options
}

// GetIntentOptions : The getIntent options.
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
func (options *GetIntentOptions) SetWorkspaceID(param string) *GetIntentOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetIntent : Allow user to set Intent
func (options *GetIntentOptions) SetIntent(param string) *GetIntentOptions {
	options.Intent = core.StringPtr(param)
	return options
}

// SetExport : Allow user to set Export
func (options *GetIntentOptions) SetExport(param bool) *GetIntentOptions {
	options.Export = core.BoolPtr(param)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *GetIntentOptions) SetIncludeAudit(param bool) *GetIntentOptions {
	options.IncludeAudit = core.BoolPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetIntentOptions) SetHeaders(param map[string]string) *GetIntentOptions {
	options.Headers = param
	return options
}

// GetSynonymOptions : The getSynonym options.
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
func (options *GetSynonymOptions) SetWorkspaceID(param string) *GetSynonymOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetEntity : Allow user to set Entity
func (options *GetSynonymOptions) SetEntity(param string) *GetSynonymOptions {
	options.Entity = core.StringPtr(param)
	return options
}

// SetValue : Allow user to set Value
func (options *GetSynonymOptions) SetValue(param string) *GetSynonymOptions {
	options.Value = core.StringPtr(param)
	return options
}

// SetSynonym : Allow user to set Synonym
func (options *GetSynonymOptions) SetSynonym(param string) *GetSynonymOptions {
	options.Synonym = core.StringPtr(param)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *GetSynonymOptions) SetIncludeAudit(param bool) *GetSynonymOptions {
	options.IncludeAudit = core.BoolPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetSynonymOptions) SetHeaders(param map[string]string) *GetSynonymOptions {
	options.Headers = param
	return options
}

// GetValueOptions : The getValue options.
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
func (options *GetValueOptions) SetWorkspaceID(param string) *GetValueOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetEntity : Allow user to set Entity
func (options *GetValueOptions) SetEntity(param string) *GetValueOptions {
	options.Entity = core.StringPtr(param)
	return options
}

// SetValue : Allow user to set Value
func (options *GetValueOptions) SetValue(param string) *GetValueOptions {
	options.Value = core.StringPtr(param)
	return options
}

// SetExport : Allow user to set Export
func (options *GetValueOptions) SetExport(param bool) *GetValueOptions {
	options.Export = core.BoolPtr(param)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *GetValueOptions) SetIncludeAudit(param bool) *GetValueOptions {
	options.IncludeAudit = core.BoolPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetValueOptions) SetHeaders(param map[string]string) *GetValueOptions {
	options.Headers = param
	return options
}

// GetWorkspaceOptions : The getWorkspace options.
type GetWorkspaceOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// Whether to include all element content in the returned data. If **export**=`false`, the returned data includes only
	// information about the element itself. If **export**=`true`, all content, including subelements, is included.
	Export *bool `json:"export,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetWorkspaceOptions : Instantiate GetWorkspaceOptions
func (assistant *AssistantV1) NewGetWorkspaceOptions(workspaceID string) *GetWorkspaceOptions {
	return &GetWorkspaceOptions{
		WorkspaceID: core.StringPtr(workspaceID),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *GetWorkspaceOptions) SetWorkspaceID(param string) *GetWorkspaceOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetExport : Allow user to set Export
func (options *GetWorkspaceOptions) SetExport(param bool) *GetWorkspaceOptions {
	options.Export = core.BoolPtr(param)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *GetWorkspaceOptions) SetIncludeAudit(param bool) *GetWorkspaceOptions {
	options.IncludeAudit = core.BoolPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetWorkspaceOptions) SetHeaders(param map[string]string) *GetWorkspaceOptions {
	options.Headers = param
	return options
}

// InputData : The user input.
type InputData struct {

	// The text of the user input. This string cannot contain carriage return, newline, or tab characters, and it must be
	// no longer than 2048 characters.
	Text *string `json:"text" validate:"required"`
}

// Intent : Intent struct
type Intent struct {

	// The name of the intent.
	IntentName *string `json:"intent" validate:"required"`

	// The timestamp for creation of the intent.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// The timestamp for the last update to the intent.
	Updated *strfmt.DateTime `json:"updated,omitempty"`

	// The description of the intent.
	Description *string `json:"description,omitempty"`
}

// IntentCollection : IntentCollection struct
type IntentCollection struct {

	// An array of objects describing the intents defined for the workspace.
	Intents []IntentExport `json:"intents" validate:"required"`

	// The pagination data for the returned objects.
	Pagination *Pagination `json:"pagination" validate:"required"`
}

// IntentExport : IntentExport struct
type IntentExport struct {

	// The name of the intent.
	IntentName *string `json:"intent" validate:"required"`

	// The timestamp for creation of the intent.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// The timestamp for the last update to the intent.
	Updated *strfmt.DateTime `json:"updated,omitempty"`

	// The description of the intent.
	Description *string `json:"description,omitempty"`

	// An array of objects describing the user input examples for the intent.
	Examples []Example `json:"examples,omitempty"`
}

// ListAllLogsOptions : The listAllLogs options.
type ListAllLogsOptions struct {

	// A cacheable parameter that limits the results to those matching the specified filter. You must specify a filter
	// query that includes a value for `language`, as well as a value for `workspace_id` or
	// `request.context.metadata.deployment`. For more information, see the
	// [documentation](https://console.bluemix.net/docs/services/conversation/filter-reference.html#filter-query-syntax).
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
func (options *ListAllLogsOptions) SetFilter(param string) *ListAllLogsOptions {
	options.Filter = core.StringPtr(param)
	return options
}

// SetSort : Allow user to set Sort
func (options *ListAllLogsOptions) SetSort(param string) *ListAllLogsOptions {
	options.Sort = core.StringPtr(param)
	return options
}

// SetPageLimit : Allow user to set PageLimit
func (options *ListAllLogsOptions) SetPageLimit(param int64) *ListAllLogsOptions {
	options.PageLimit = core.Int64Ptr(param)
	return options
}

// SetCursor : Allow user to set Cursor
func (options *ListAllLogsOptions) SetCursor(param string) *ListAllLogsOptions {
	options.Cursor = core.StringPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListAllLogsOptions) SetHeaders(param map[string]string) *ListAllLogsOptions {
	options.Headers = param
	return options
}

// ListCounterexamplesOptions : The listCounterexamples options.
type ListCounterexamplesOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The number of records to return in each page of results.
	PageLimit *int64 `json:"page_limit,omitempty"`

	// Whether to include information about the number of records returned.
	IncludeCount *bool `json:"include_count,omitempty"`

	// The attribute by which returned results will be sorted. To reverse the sort order, prefix the value with a minus
	// sign (`-`).
	Sort *string `json:"sort,omitempty"`

	// A token identifying the page of results to retrieve.
	Cursor *string `json:"cursor,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListCounterexamplesOptions : Instantiate ListCounterexamplesOptions
func (assistant *AssistantV1) NewListCounterexamplesOptions(workspaceID string) *ListCounterexamplesOptions {
	return &ListCounterexamplesOptions{
		WorkspaceID: core.StringPtr(workspaceID),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *ListCounterexamplesOptions) SetWorkspaceID(param string) *ListCounterexamplesOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetPageLimit : Allow user to set PageLimit
func (options *ListCounterexamplesOptions) SetPageLimit(param int64) *ListCounterexamplesOptions {
	options.PageLimit = core.Int64Ptr(param)
	return options
}

// SetIncludeCount : Allow user to set IncludeCount
func (options *ListCounterexamplesOptions) SetIncludeCount(param bool) *ListCounterexamplesOptions {
	options.IncludeCount = core.BoolPtr(param)
	return options
}

// SetSort : Allow user to set Sort
func (options *ListCounterexamplesOptions) SetSort(param string) *ListCounterexamplesOptions {
	options.Sort = core.StringPtr(param)
	return options
}

// SetCursor : Allow user to set Cursor
func (options *ListCounterexamplesOptions) SetCursor(param string) *ListCounterexamplesOptions {
	options.Cursor = core.StringPtr(param)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *ListCounterexamplesOptions) SetIncludeAudit(param bool) *ListCounterexamplesOptions {
	options.IncludeAudit = core.BoolPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListCounterexamplesOptions) SetHeaders(param map[string]string) *ListCounterexamplesOptions {
	options.Headers = param
	return options
}

// ListDialogNodesOptions : The listDialogNodes options.
type ListDialogNodesOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The number of records to return in each page of results.
	PageLimit *int64 `json:"page_limit,omitempty"`

	// Whether to include information about the number of records returned.
	IncludeCount *bool `json:"include_count,omitempty"`

	// The attribute by which returned results will be sorted. To reverse the sort order, prefix the value with a minus
	// sign (`-`).
	Sort *string `json:"sort,omitempty"`

	// A token identifying the page of results to retrieve.
	Cursor *string `json:"cursor,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListDialogNodesOptions : Instantiate ListDialogNodesOptions
func (assistant *AssistantV1) NewListDialogNodesOptions(workspaceID string) *ListDialogNodesOptions {
	return &ListDialogNodesOptions{
		WorkspaceID: core.StringPtr(workspaceID),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *ListDialogNodesOptions) SetWorkspaceID(param string) *ListDialogNodesOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetPageLimit : Allow user to set PageLimit
func (options *ListDialogNodesOptions) SetPageLimit(param int64) *ListDialogNodesOptions {
	options.PageLimit = core.Int64Ptr(param)
	return options
}

// SetIncludeCount : Allow user to set IncludeCount
func (options *ListDialogNodesOptions) SetIncludeCount(param bool) *ListDialogNodesOptions {
	options.IncludeCount = core.BoolPtr(param)
	return options
}

// SetSort : Allow user to set Sort
func (options *ListDialogNodesOptions) SetSort(param string) *ListDialogNodesOptions {
	options.Sort = core.StringPtr(param)
	return options
}

// SetCursor : Allow user to set Cursor
func (options *ListDialogNodesOptions) SetCursor(param string) *ListDialogNodesOptions {
	options.Cursor = core.StringPtr(param)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *ListDialogNodesOptions) SetIncludeAudit(param bool) *ListDialogNodesOptions {
	options.IncludeAudit = core.BoolPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListDialogNodesOptions) SetHeaders(param map[string]string) *ListDialogNodesOptions {
	options.Headers = param
	return options
}

// ListEntitiesOptions : The listEntities options.
type ListEntitiesOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// Whether to include all element content in the returned data. If **export**=`false`, the returned data includes only
	// information about the element itself. If **export**=`true`, all content, including subelements, is included.
	Export *bool `json:"export,omitempty"`

	// The number of records to return in each page of results.
	PageLimit *int64 `json:"page_limit,omitempty"`

	// Whether to include information about the number of records returned.
	IncludeCount *bool `json:"include_count,omitempty"`

	// The attribute by which returned results will be sorted. To reverse the sort order, prefix the value with a minus
	// sign (`-`).
	Sort *string `json:"sort,omitempty"`

	// A token identifying the page of results to retrieve.
	Cursor *string `json:"cursor,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListEntitiesOptions : Instantiate ListEntitiesOptions
func (assistant *AssistantV1) NewListEntitiesOptions(workspaceID string) *ListEntitiesOptions {
	return &ListEntitiesOptions{
		WorkspaceID: core.StringPtr(workspaceID),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *ListEntitiesOptions) SetWorkspaceID(param string) *ListEntitiesOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetExport : Allow user to set Export
func (options *ListEntitiesOptions) SetExport(param bool) *ListEntitiesOptions {
	options.Export = core.BoolPtr(param)
	return options
}

// SetPageLimit : Allow user to set PageLimit
func (options *ListEntitiesOptions) SetPageLimit(param int64) *ListEntitiesOptions {
	options.PageLimit = core.Int64Ptr(param)
	return options
}

// SetIncludeCount : Allow user to set IncludeCount
func (options *ListEntitiesOptions) SetIncludeCount(param bool) *ListEntitiesOptions {
	options.IncludeCount = core.BoolPtr(param)
	return options
}

// SetSort : Allow user to set Sort
func (options *ListEntitiesOptions) SetSort(param string) *ListEntitiesOptions {
	options.Sort = core.StringPtr(param)
	return options
}

// SetCursor : Allow user to set Cursor
func (options *ListEntitiesOptions) SetCursor(param string) *ListEntitiesOptions {
	options.Cursor = core.StringPtr(param)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *ListEntitiesOptions) SetIncludeAudit(param bool) *ListEntitiesOptions {
	options.IncludeAudit = core.BoolPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListEntitiesOptions) SetHeaders(param map[string]string) *ListEntitiesOptions {
	options.Headers = param
	return options
}

// ListExamplesOptions : The listExamples options.
type ListExamplesOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The intent name.
	Intent *string `json:"intent" validate:"required"`

	// The number of records to return in each page of results.
	PageLimit *int64 `json:"page_limit,omitempty"`

	// Whether to include information about the number of records returned.
	IncludeCount *bool `json:"include_count,omitempty"`

	// The attribute by which returned results will be sorted. To reverse the sort order, prefix the value with a minus
	// sign (`-`).
	Sort *string `json:"sort,omitempty"`

	// A token identifying the page of results to retrieve.
	Cursor *string `json:"cursor,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListExamplesOptions : Instantiate ListExamplesOptions
func (assistant *AssistantV1) NewListExamplesOptions(workspaceID string, intent string) *ListExamplesOptions {
	return &ListExamplesOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Intent:      core.StringPtr(intent),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *ListExamplesOptions) SetWorkspaceID(param string) *ListExamplesOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetIntent : Allow user to set Intent
func (options *ListExamplesOptions) SetIntent(param string) *ListExamplesOptions {
	options.Intent = core.StringPtr(param)
	return options
}

// SetPageLimit : Allow user to set PageLimit
func (options *ListExamplesOptions) SetPageLimit(param int64) *ListExamplesOptions {
	options.PageLimit = core.Int64Ptr(param)
	return options
}

// SetIncludeCount : Allow user to set IncludeCount
func (options *ListExamplesOptions) SetIncludeCount(param bool) *ListExamplesOptions {
	options.IncludeCount = core.BoolPtr(param)
	return options
}

// SetSort : Allow user to set Sort
func (options *ListExamplesOptions) SetSort(param string) *ListExamplesOptions {
	options.Sort = core.StringPtr(param)
	return options
}

// SetCursor : Allow user to set Cursor
func (options *ListExamplesOptions) SetCursor(param string) *ListExamplesOptions {
	options.Cursor = core.StringPtr(param)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *ListExamplesOptions) SetIncludeAudit(param bool) *ListExamplesOptions {
	options.IncludeAudit = core.BoolPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListExamplesOptions) SetHeaders(param map[string]string) *ListExamplesOptions {
	options.Headers = param
	return options
}

// ListIntentsOptions : The listIntents options.
type ListIntentsOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// Whether to include all element content in the returned data. If **export**=`false`, the returned data includes only
	// information about the element itself. If **export**=`true`, all content, including subelements, is included.
	Export *bool `json:"export,omitempty"`

	// The number of records to return in each page of results.
	PageLimit *int64 `json:"page_limit,omitempty"`

	// Whether to include information about the number of records returned.
	IncludeCount *bool `json:"include_count,omitempty"`

	// The attribute by which returned results will be sorted. To reverse the sort order, prefix the value with a minus
	// sign (`-`).
	Sort *string `json:"sort,omitempty"`

	// A token identifying the page of results to retrieve.
	Cursor *string `json:"cursor,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListIntentsOptions : Instantiate ListIntentsOptions
func (assistant *AssistantV1) NewListIntentsOptions(workspaceID string) *ListIntentsOptions {
	return &ListIntentsOptions{
		WorkspaceID: core.StringPtr(workspaceID),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *ListIntentsOptions) SetWorkspaceID(param string) *ListIntentsOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetExport : Allow user to set Export
func (options *ListIntentsOptions) SetExport(param bool) *ListIntentsOptions {
	options.Export = core.BoolPtr(param)
	return options
}

// SetPageLimit : Allow user to set PageLimit
func (options *ListIntentsOptions) SetPageLimit(param int64) *ListIntentsOptions {
	options.PageLimit = core.Int64Ptr(param)
	return options
}

// SetIncludeCount : Allow user to set IncludeCount
func (options *ListIntentsOptions) SetIncludeCount(param bool) *ListIntentsOptions {
	options.IncludeCount = core.BoolPtr(param)
	return options
}

// SetSort : Allow user to set Sort
func (options *ListIntentsOptions) SetSort(param string) *ListIntentsOptions {
	options.Sort = core.StringPtr(param)
	return options
}

// SetCursor : Allow user to set Cursor
func (options *ListIntentsOptions) SetCursor(param string) *ListIntentsOptions {
	options.Cursor = core.StringPtr(param)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *ListIntentsOptions) SetIncludeAudit(param bool) *ListIntentsOptions {
	options.IncludeAudit = core.BoolPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListIntentsOptions) SetHeaders(param map[string]string) *ListIntentsOptions {
	options.Headers = param
	return options
}

// ListLogsOptions : The listLogs options.
type ListLogsOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// How to sort the returned log events. You can sort by **request_timestamp**. To reverse the sort order, prefix the
	// parameter value with a minus sign (`-`).
	Sort *string `json:"sort,omitempty"`

	// A cacheable parameter that limits the results to those matching the specified filter. For more information, see the
	// [documentation](https://console.bluemix.net/docs/services/conversation/filter-reference.html#filter-query-syntax).
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
func (options *ListLogsOptions) SetWorkspaceID(param string) *ListLogsOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetSort : Allow user to set Sort
func (options *ListLogsOptions) SetSort(param string) *ListLogsOptions {
	options.Sort = core.StringPtr(param)
	return options
}

// SetFilter : Allow user to set Filter
func (options *ListLogsOptions) SetFilter(param string) *ListLogsOptions {
	options.Filter = core.StringPtr(param)
	return options
}

// SetPageLimit : Allow user to set PageLimit
func (options *ListLogsOptions) SetPageLimit(param int64) *ListLogsOptions {
	options.PageLimit = core.Int64Ptr(param)
	return options
}

// SetCursor : Allow user to set Cursor
func (options *ListLogsOptions) SetCursor(param string) *ListLogsOptions {
	options.Cursor = core.StringPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListLogsOptions) SetHeaders(param map[string]string) *ListLogsOptions {
	options.Headers = param
	return options
}

// ListMentionsOptions : The listMentions options.
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
func (options *ListMentionsOptions) SetWorkspaceID(param string) *ListMentionsOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetEntity : Allow user to set Entity
func (options *ListMentionsOptions) SetEntity(param string) *ListMentionsOptions {
	options.Entity = core.StringPtr(param)
	return options
}

// SetExport : Allow user to set Export
func (options *ListMentionsOptions) SetExport(param bool) *ListMentionsOptions {
	options.Export = core.BoolPtr(param)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *ListMentionsOptions) SetIncludeAudit(param bool) *ListMentionsOptions {
	options.IncludeAudit = core.BoolPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListMentionsOptions) SetHeaders(param map[string]string) *ListMentionsOptions {
	options.Headers = param
	return options
}

// ListSynonymsOptions : The listSynonyms options.
type ListSynonymsOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The name of the entity.
	Entity *string `json:"entity" validate:"required"`

	// The text of the entity value.
	Value *string `json:"value" validate:"required"`

	// The number of records to return in each page of results.
	PageLimit *int64 `json:"page_limit,omitempty"`

	// Whether to include information about the number of records returned.
	IncludeCount *bool `json:"include_count,omitempty"`

	// The attribute by which returned results will be sorted. To reverse the sort order, prefix the value with a minus
	// sign (`-`).
	Sort *string `json:"sort,omitempty"`

	// A token identifying the page of results to retrieve.
	Cursor *string `json:"cursor,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListSynonymsOptions : Instantiate ListSynonymsOptions
func (assistant *AssistantV1) NewListSynonymsOptions(workspaceID string, entity string, value string) *ListSynonymsOptions {
	return &ListSynonymsOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Entity:      core.StringPtr(entity),
		Value:       core.StringPtr(value),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *ListSynonymsOptions) SetWorkspaceID(param string) *ListSynonymsOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetEntity : Allow user to set Entity
func (options *ListSynonymsOptions) SetEntity(param string) *ListSynonymsOptions {
	options.Entity = core.StringPtr(param)
	return options
}

// SetValue : Allow user to set Value
func (options *ListSynonymsOptions) SetValue(param string) *ListSynonymsOptions {
	options.Value = core.StringPtr(param)
	return options
}

// SetPageLimit : Allow user to set PageLimit
func (options *ListSynonymsOptions) SetPageLimit(param int64) *ListSynonymsOptions {
	options.PageLimit = core.Int64Ptr(param)
	return options
}

// SetIncludeCount : Allow user to set IncludeCount
func (options *ListSynonymsOptions) SetIncludeCount(param bool) *ListSynonymsOptions {
	options.IncludeCount = core.BoolPtr(param)
	return options
}

// SetSort : Allow user to set Sort
func (options *ListSynonymsOptions) SetSort(param string) *ListSynonymsOptions {
	options.Sort = core.StringPtr(param)
	return options
}

// SetCursor : Allow user to set Cursor
func (options *ListSynonymsOptions) SetCursor(param string) *ListSynonymsOptions {
	options.Cursor = core.StringPtr(param)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *ListSynonymsOptions) SetIncludeAudit(param bool) *ListSynonymsOptions {
	options.IncludeAudit = core.BoolPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListSynonymsOptions) SetHeaders(param map[string]string) *ListSynonymsOptions {
	options.Headers = param
	return options
}

// ListValuesOptions : The listValues options.
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

	// Whether to include information about the number of records returned.
	IncludeCount *bool `json:"include_count,omitempty"`

	// The attribute by which returned results will be sorted. To reverse the sort order, prefix the value with a minus
	// sign (`-`).
	Sort *string `json:"sort,omitempty"`

	// A token identifying the page of results to retrieve.
	Cursor *string `json:"cursor,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListValuesOptions : Instantiate ListValuesOptions
func (assistant *AssistantV1) NewListValuesOptions(workspaceID string, entity string) *ListValuesOptions {
	return &ListValuesOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Entity:      core.StringPtr(entity),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *ListValuesOptions) SetWorkspaceID(param string) *ListValuesOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetEntity : Allow user to set Entity
func (options *ListValuesOptions) SetEntity(param string) *ListValuesOptions {
	options.Entity = core.StringPtr(param)
	return options
}

// SetExport : Allow user to set Export
func (options *ListValuesOptions) SetExport(param bool) *ListValuesOptions {
	options.Export = core.BoolPtr(param)
	return options
}

// SetPageLimit : Allow user to set PageLimit
func (options *ListValuesOptions) SetPageLimit(param int64) *ListValuesOptions {
	options.PageLimit = core.Int64Ptr(param)
	return options
}

// SetIncludeCount : Allow user to set IncludeCount
func (options *ListValuesOptions) SetIncludeCount(param bool) *ListValuesOptions {
	options.IncludeCount = core.BoolPtr(param)
	return options
}

// SetSort : Allow user to set Sort
func (options *ListValuesOptions) SetSort(param string) *ListValuesOptions {
	options.Sort = core.StringPtr(param)
	return options
}

// SetCursor : Allow user to set Cursor
func (options *ListValuesOptions) SetCursor(param string) *ListValuesOptions {
	options.Cursor = core.StringPtr(param)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *ListValuesOptions) SetIncludeAudit(param bool) *ListValuesOptions {
	options.IncludeAudit = core.BoolPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListValuesOptions) SetHeaders(param map[string]string) *ListValuesOptions {
	options.Headers = param
	return options
}

// ListWorkspacesOptions : The listWorkspaces options.
type ListWorkspacesOptions struct {

	// The number of records to return in each page of results.
	PageLimit *int64 `json:"page_limit,omitempty"`

	// Whether to include information about the number of records returned.
	IncludeCount *bool `json:"include_count,omitempty"`

	// The attribute by which returned results will be sorted. To reverse the sort order, prefix the value with a minus
	// sign (`-`).
	Sort *string `json:"sort,omitempty"`

	// A token identifying the page of results to retrieve.
	Cursor *string `json:"cursor,omitempty"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit *bool `json:"include_audit,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListWorkspacesOptions : Instantiate ListWorkspacesOptions
func (assistant *AssistantV1) NewListWorkspacesOptions() *ListWorkspacesOptions {
	return &ListWorkspacesOptions{}
}

// SetPageLimit : Allow user to set PageLimit
func (options *ListWorkspacesOptions) SetPageLimit(param int64) *ListWorkspacesOptions {
	options.PageLimit = core.Int64Ptr(param)
	return options
}

// SetIncludeCount : Allow user to set IncludeCount
func (options *ListWorkspacesOptions) SetIncludeCount(param bool) *ListWorkspacesOptions {
	options.IncludeCount = core.BoolPtr(param)
	return options
}

// SetSort : Allow user to set Sort
func (options *ListWorkspacesOptions) SetSort(param string) *ListWorkspacesOptions {
	options.Sort = core.StringPtr(param)
	return options
}

// SetCursor : Allow user to set Cursor
func (options *ListWorkspacesOptions) SetCursor(param string) *ListWorkspacesOptions {
	options.Cursor = core.StringPtr(param)
	return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *ListWorkspacesOptions) SetIncludeAudit(param bool) *ListWorkspacesOptions {
	options.IncludeAudit = core.BoolPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListWorkspacesOptions) SetHeaders(param map[string]string) *ListWorkspacesOptions {
	options.Headers = param
	return options
}

// LogCollection : LogCollection struct
type LogCollection struct {

	// An array of objects describing log events.
	Logs []LogExport `json:"logs" validate:"required"`

	// The pagination data for the returned objects.
	Pagination *LogPagination `json:"pagination" validate:"required"`
}

// LogExport : LogExport struct
type LogExport struct {

	// A request received by the workspace, including the user input and context.
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

// LogMessage : Log message details.
type LogMessage struct {

	// The severity of the log message.
	Level *string `json:"level" validate:"required"`

	// The text of the log message.
	Msg *string `json:"msg" validate:"required"`
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

// Mentions : A mention of a contextual entity.
type Mentions struct {

	// The name of the entity.
	Entity *string `json:"entity" validate:"required"`

	// An array of zero-based character offsets that indicate where the entity mentions begin and end in the input text.
	Location []int64 `json:"location" validate:"required"`
}

// MessageInput : The text of the user input.
type MessageInput struct {

	// The user's input.
	Text *string `json:"text,omitempty"`
}

// MessageOptions : The message options.
type MessageOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// An input object that includes the input text.
	Input *InputData `json:"input,omitempty"`

	// Whether to return more than one intent. Set to `true` to return all matching intents.
	AlternateIntents *bool `json:"alternate_intents,omitempty"`

	// State information for the conversation. Continue a conversation by including the context object from the previous
	// response.
	Context *Context `json:"context,omitempty"`

	// Entities to use when evaluating the message. Include entities from the previous response to continue using those
	// entities rather than detecting entities in the new input.
	Entities []RuntimeEntity `json:"entities,omitempty"`

	// Intents to use when evaluating the user input. Include intents from the previous response to continue using those
	// intents rather than trying to recognize intents in the new input.
	Intents []RuntimeIntent `json:"intents,omitempty"`

	// System output. Include the output from the previous response to maintain intermediate information over multiple
	// requests.
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
func (options *MessageOptions) SetWorkspaceID(param string) *MessageOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetInput : Allow user to set Input
func (options *MessageOptions) SetInput(param InputData) *MessageOptions {
	options.Input = &param
	return options
}

// SetAlternateIntents : Allow user to set AlternateIntents
func (options *MessageOptions) SetAlternateIntents(param bool) *MessageOptions {
	options.AlternateIntents = core.BoolPtr(param)
	return options
}

// SetContext : Allow user to set Context
func (options *MessageOptions) SetContext(param Context) *MessageOptions {
	options.Context = &param
	return options
}

// SetEntities : Allow user to set Entities
func (options *MessageOptions) SetEntities(param []RuntimeEntity) *MessageOptions {
	options.Entities = param
	return options
}

// SetIntents : Allow user to set Intents
func (options *MessageOptions) SetIntents(param []RuntimeIntent) *MessageOptions {
	options.Intents = param
	return options
}

// SetOutput : Allow user to set Output
func (options *MessageOptions) SetOutput(param OutputData) *MessageOptions {
	options.Output = &param
	return options
}

// SetNodesVisitedDetails : Allow user to set NodesVisitedDetails
func (options *MessageOptions) SetNodesVisitedDetails(param bool) *MessageOptions {
	options.NodesVisitedDetails = core.BoolPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *MessageOptions) SetHeaders(param map[string]string) *MessageOptions {
	options.Headers = param
	return options
}

// MessageRequest : A message request formatted for the Watson Assistant service.
type MessageRequest struct {

	// An input object that includes the input text.
	Input *InputData `json:"input,omitempty"`

	// Whether to return more than one intent. Set to `true` to return all matching intents.
	AlternateIntents *bool `json:"alternate_intents,omitempty"`

	// State information for the conversation. Continue a conversation by including the context object from the previous
	// response.
	Context *Context `json:"context,omitempty"`

	// Entities to use when evaluating the message. Include entities from the previous response to continue using those
	// entities rather than detecting entities in the new input.
	Entities []RuntimeEntity `json:"entities,omitempty"`

	// Intents to use when evaluating the user input. Include intents from the previous response to continue using those
	// intents rather than trying to recognize intents in the new input.
	Intents []RuntimeIntent `json:"intents,omitempty"`

	// System output. Include the output from the previous response to maintain intermediate information over multiple
	// requests.
	Output *OutputData `json:"output,omitempty"`
}

// MessageResponse : A response from the Watson Assistant service.
type MessageResponse struct {

	// The user input from the request.
	Input *MessageInput `json:"input,omitempty"`

	// An array of intents recognized in the user input, sorted in descending order of confidence.
	Intents []RuntimeIntent `json:"intents" validate:"required"`

	// An array of entities identified in the user input.
	Entities []RuntimeEntity `json:"entities" validate:"required"`

	// Whether to return more than one intent. A value of `true` indicates that all matching intents are returned.
	AlternateIntents *bool `json:"alternate_intents,omitempty"`

	// State information for the conversation.
	Context *Context `json:"context" validate:"required"`

	// Output from the dialog, including the response to the user, the nodes that were triggered, and log messages.
	Output *OutputData `json:"output" validate:"required"`

	// An array of objects describing any actions requested by the dialog node.
	Actions []DialogNodeAction `json:"actions,omitempty"`
}

// OutputData : An output object that includes the response to the user, the dialog nodes that were triggered, and messages from the
// log.
type OutputData struct {

	// An array of up to 50 messages logged with the request.
	LogMessages []LogMessage `json:"log_messages" validate:"required"`

	// An array of responses to the user.
	Text []string `json:"text" validate:"required"`

	// Output intended for any channel. It is the responsibility of the client application to implement the supported
	// response types.
	Generic []DialogRuntimeResponseGeneric `json:"generic,omitempty"`

	// An array of the nodes that were triggered to create the response, in the order in which they were visited. This
	// information is useful for debugging and for tracing the path taken through the node tree.
	NodesVisited []string `json:"nodes_visited,omitempty"`

	// An array of objects containing detailed diagnostic information about the nodes that were triggered during processing
	// of the input message. Included only if **nodes_visited_details** is set to `true` in the message request.
	NodesVisitedDetails []DialogNodeVisitedDetails `json:"nodes_visited_details,omitempty"`
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

	// The term in the input text that was recognized as an entity value.
	Value *string `json:"value" validate:"required"`

	// A decimal percentage that represents Watson's confidence in the entity.
	Confidence *float64 `json:"confidence,omitempty"`

	// Any metadata for the entity.
	Metadata interface{} `json:"metadata,omitempty"`

	// The recognized capture groups for the entity, as defined by the entity pattern.
	Groups []CaptureGroup `json:"groups,omitempty"`
}

// RuntimeIntent : An intent identified in the user input.
type RuntimeIntent struct {

	// The name of the recognized intent.
	Intent *string `json:"intent" validate:"required"`

	// A decimal percentage that represents Watson's confidence in the intent.
	Confidence *float64 `json:"confidence" validate:"required"`
}

// Synonym : Synonym struct
type Synonym struct {

	// The text of the synonym.
	SynonymText *string `json:"synonym" validate:"required"`

	// The timestamp for creation of the synonym.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// The timestamp for the most recent update to the synonym.
	Updated *strfmt.DateTime `json:"updated,omitempty"`
}

// SynonymCollection : SynonymCollection struct
type SynonymCollection struct {

	// An array of synonyms.
	Synonyms []Synonym `json:"synonyms" validate:"required"`

	// The pagination data for the returned objects.
	Pagination *Pagination `json:"pagination" validate:"required"`
}

// SystemResponse : For internal use only.
type SystemResponse struct {
}

// UpdateCounterexampleOptions : The updateCounterexample options.
type UpdateCounterexampleOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The text of a user input counterexample (for example, `What are you wearing?`).
	Text *string `json:"text" validate:"required"`

	// The text of a user input counterexample.
	NewText *string `json:"new_text,omitempty"`

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
func (options *UpdateCounterexampleOptions) SetWorkspaceID(param string) *UpdateCounterexampleOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetText : Allow user to set Text
func (options *UpdateCounterexampleOptions) SetText(param string) *UpdateCounterexampleOptions {
	options.Text = core.StringPtr(param)
	return options
}

// SetNewText : Allow user to set NewText
func (options *UpdateCounterexampleOptions) SetNewText(param string) *UpdateCounterexampleOptions {
	options.NewText = core.StringPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateCounterexampleOptions) SetHeaders(param map[string]string) *UpdateCounterexampleOptions {
	options.Headers = param
	return options
}

// UpdateDialogNodeOptions : The updateDialogNode options.
type UpdateDialogNodeOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The dialog node ID (for example, `get_order`).
	DialogNode *string `json:"dialog_node" validate:"required"`

	// How the dialog node is processed.
	NodeType *string `json:"new_type,omitempty"`

	// An array of objects describing any actions to be invoked by the dialog node.
	NewActions []DialogNodeAction `json:"new_actions,omitempty"`

	// The condition that will trigger the dialog node. This string cannot contain carriage return, newline, or tab
	// characters, and it must be no longer than 2048 characters.
	NewConditions *string `json:"new_conditions,omitempty"`

	// The context for the dialog node.
	NewContext interface{} `json:"new_context,omitempty"`

	// The ID of the previous sibling dialog node.
	NewPreviousSibling *string `json:"new_previous_sibling,omitempty"`

	// The location in the dialog context where output is stored.
	NewVariable *string `json:"new_variable,omitempty"`

	// A label that can be displayed externally to describe the purpose of the node to users. This string must be no longer
	// than 512 characters.
	NewUserLabel *string `json:"new_user_label,omitempty"`

	// The metadata for the dialog node.
	NewMetadata interface{} `json:"new_metadata,omitempty"`

	// The alias used to identify the dialog node. This string must conform to the following restrictions:
	// - It can contain only Unicode alphanumeric, space, underscore, hyphen, and dot characters.
	// - It must be no longer than 64 characters.
	NewTitle *string `json:"new_title,omitempty"`

	// The description of the dialog node. This string cannot contain carriage return, newline, or tab characters, and it
	// must be no longer than 128 characters.
	NewDescription *string `json:"new_description,omitempty"`

	// Whether this dialog node can be returned to after a digression.
	NewDigressOut *string `json:"new_digress_out,omitempty"`

	// How an `event_handler` node is processed.
	NewEventName *string `json:"new_event_name,omitempty"`

	// Whether the user can digress to top-level nodes while filling out slots.
	NewDigressOutSlots *string `json:"new_digress_out_slots,omitempty"`

	// The next step to be executed in dialog processing.
	NewNextStep *DialogNodeNextStep `json:"new_next_step,omitempty"`

	// Whether this top-level dialog node can be digressed into.
	NewDigressIn *string `json:"new_digress_in,omitempty"`

	// The output of the dialog node. For more information about how to specify dialog node output, see the
	// [documentation](https://console.bluemix.net/docs/services/conversation/dialog-overview.html#complex).
	NewOutput *DialogNodeOutput `json:"new_output,omitempty"`

	// The ID of the parent dialog node.
	NewParent *string `json:"new_parent,omitempty"`

	// The dialog node ID. This string must conform to the following restrictions:
	// - It can contain only Unicode alphanumeric, space, underscore, hyphen, and dot characters.
	// - It must be no longer than 1024 characters.
	NewDialogNode *string `json:"new_dialog_node,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewUpdateDialogNodeOptions : Instantiate UpdateDialogNodeOptions
func (assistant *AssistantV1) NewUpdateDialogNodeOptions(workspaceID string, dialogNode string) *UpdateDialogNodeOptions {
	return &UpdateDialogNodeOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		DialogNode:  core.StringPtr(dialogNode),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *UpdateDialogNodeOptions) SetWorkspaceID(param string) *UpdateDialogNodeOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetDialogNode : Allow user to set DialogNode
func (options *UpdateDialogNodeOptions) SetDialogNode(param string) *UpdateDialogNodeOptions {
	options.DialogNode = core.StringPtr(param)
	return options
}

// SetNodeType : Allow user to set NodeType
func (options *UpdateDialogNodeOptions) SetNodeType(param string) *UpdateDialogNodeOptions {
	options.NodeType = core.StringPtr(param)
	return options
}

// SetNewActions : Allow user to set NewActions
func (options *UpdateDialogNodeOptions) SetNewActions(param []DialogNodeAction) *UpdateDialogNodeOptions {
	options.NewActions = param
	return options
}

// SetNewConditions : Allow user to set NewConditions
func (options *UpdateDialogNodeOptions) SetNewConditions(param string) *UpdateDialogNodeOptions {
	options.NewConditions = core.StringPtr(param)
	return options
}

// SetNewContext : Allow user to set NewContext
func (options *UpdateDialogNodeOptions) SetNewContext(param interface{}) *UpdateDialogNodeOptions {
	options.NewContext = &param
	return options
}

// SetNewPreviousSibling : Allow user to set NewPreviousSibling
func (options *UpdateDialogNodeOptions) SetNewPreviousSibling(param string) *UpdateDialogNodeOptions {
	options.NewPreviousSibling = core.StringPtr(param)
	return options
}

// SetNewVariable : Allow user to set NewVariable
func (options *UpdateDialogNodeOptions) SetNewVariable(param string) *UpdateDialogNodeOptions {
	options.NewVariable = core.StringPtr(param)
	return options
}

// SetNewUserLabel : Allow user to set NewUserLabel
func (options *UpdateDialogNodeOptions) SetNewUserLabel(param string) *UpdateDialogNodeOptions {
	options.NewUserLabel = core.StringPtr(param)
	return options
}

// SetNewMetadata : Allow user to set NewMetadata
func (options *UpdateDialogNodeOptions) SetNewMetadata(param interface{}) *UpdateDialogNodeOptions {
	options.NewMetadata = &param
	return options
}

// SetNewTitle : Allow user to set NewTitle
func (options *UpdateDialogNodeOptions) SetNewTitle(param string) *UpdateDialogNodeOptions {
	options.NewTitle = core.StringPtr(param)
	return options
}

// SetNewDescription : Allow user to set NewDescription
func (options *UpdateDialogNodeOptions) SetNewDescription(param string) *UpdateDialogNodeOptions {
	options.NewDescription = core.StringPtr(param)
	return options
}

// SetNewDigressOut : Allow user to set NewDigressOut
func (options *UpdateDialogNodeOptions) SetNewDigressOut(param string) *UpdateDialogNodeOptions {
	options.NewDigressOut = core.StringPtr(param)
	return options
}

// SetNewEventName : Allow user to set NewEventName
func (options *UpdateDialogNodeOptions) SetNewEventName(param string) *UpdateDialogNodeOptions {
	options.NewEventName = core.StringPtr(param)
	return options
}

// SetNewDigressOutSlots : Allow user to set NewDigressOutSlots
func (options *UpdateDialogNodeOptions) SetNewDigressOutSlots(param string) *UpdateDialogNodeOptions {
	options.NewDigressOutSlots = core.StringPtr(param)
	return options
}

// SetNewNextStep : Allow user to set NewNextStep
func (options *UpdateDialogNodeOptions) SetNewNextStep(param DialogNodeNextStep) *UpdateDialogNodeOptions {
	options.NewNextStep = &param
	return options
}

// SetNewDigressIn : Allow user to set NewDigressIn
func (options *UpdateDialogNodeOptions) SetNewDigressIn(param string) *UpdateDialogNodeOptions {
	options.NewDigressIn = core.StringPtr(param)
	return options
}

// SetNewOutput : Allow user to set NewOutput
func (options *UpdateDialogNodeOptions) SetNewOutput(param DialogNodeOutput) *UpdateDialogNodeOptions {
	options.NewOutput = &param
	return options
}

// SetNewParent : Allow user to set NewParent
func (options *UpdateDialogNodeOptions) SetNewParent(param string) *UpdateDialogNodeOptions {
	options.NewParent = core.StringPtr(param)
	return options
}

// SetNewDialogNode : Allow user to set NewDialogNode
func (options *UpdateDialogNodeOptions) SetNewDialogNode(param string) *UpdateDialogNodeOptions {
	options.NewDialogNode = core.StringPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateDialogNodeOptions) SetHeaders(param map[string]string) *UpdateDialogNodeOptions {
	options.Headers = param
	return options
}

// UpdateEntityOptions : The updateEntity options.
type UpdateEntityOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The name of the entity.
	Entity *string `json:"entity" validate:"required"`

	// Whether to use fuzzy matching for the entity.
	NewFuzzyMatch *bool `json:"new_fuzzy_match,omitempty"`

	// The name of the entity. This string must conform to the following restrictions:
	// - It can contain only Unicode alphanumeric, underscore, and hyphen characters.
	// - It cannot begin with the reserved prefix `sys-`.
	// - It must be no longer than 64 characters.
	NewEntity *string `json:"new_entity,omitempty"`

	// Any metadata related to the entity.
	NewMetadata interface{} `json:"new_metadata,omitempty"`

	// An array of entity values.
	NewValues []CreateValue `json:"new_values,omitempty"`

	// The description of the entity. This string cannot contain carriage return, newline, or tab characters, and it must
	// be no longer than 128 characters.
	NewDescription *string `json:"new_description,omitempty"`

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
func (options *UpdateEntityOptions) SetWorkspaceID(param string) *UpdateEntityOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetEntity : Allow user to set Entity
func (options *UpdateEntityOptions) SetEntity(param string) *UpdateEntityOptions {
	options.Entity = core.StringPtr(param)
	return options
}

// SetNewFuzzyMatch : Allow user to set NewFuzzyMatch
func (options *UpdateEntityOptions) SetNewFuzzyMatch(param bool) *UpdateEntityOptions {
	options.NewFuzzyMatch = core.BoolPtr(param)
	return options
}

// SetNewEntity : Allow user to set NewEntity
func (options *UpdateEntityOptions) SetNewEntity(param string) *UpdateEntityOptions {
	options.NewEntity = core.StringPtr(param)
	return options
}

// SetNewMetadata : Allow user to set NewMetadata
func (options *UpdateEntityOptions) SetNewMetadata(param interface{}) *UpdateEntityOptions {
	options.NewMetadata = &param
	return options
}

// SetNewValues : Allow user to set NewValues
func (options *UpdateEntityOptions) SetNewValues(param []CreateValue) *UpdateEntityOptions {
	options.NewValues = param
	return options
}

// SetNewDescription : Allow user to set NewDescription
func (options *UpdateEntityOptions) SetNewDescription(param string) *UpdateEntityOptions {
	options.NewDescription = core.StringPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateEntityOptions) SetHeaders(param map[string]string) *UpdateEntityOptions {
	options.Headers = param
	return options
}

// UpdateExampleOptions : The updateExample options.
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
	// - It must be no longer than 1024 characters.
	NewText *string `json:"new_text,omitempty"`

	// An array of contextual entity mentions.
	NewMentions []Mentions `json:"new_mentions,omitempty"`

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
func (options *UpdateExampleOptions) SetWorkspaceID(param string) *UpdateExampleOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetIntent : Allow user to set Intent
func (options *UpdateExampleOptions) SetIntent(param string) *UpdateExampleOptions {
	options.Intent = core.StringPtr(param)
	return options
}

// SetText : Allow user to set Text
func (options *UpdateExampleOptions) SetText(param string) *UpdateExampleOptions {
	options.Text = core.StringPtr(param)
	return options
}

// SetNewText : Allow user to set NewText
func (options *UpdateExampleOptions) SetNewText(param string) *UpdateExampleOptions {
	options.NewText = core.StringPtr(param)
	return options
}

// SetNewMentions : Allow user to set NewMentions
func (options *UpdateExampleOptions) SetNewMentions(param []Mentions) *UpdateExampleOptions {
	options.NewMentions = param
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateExampleOptions) SetHeaders(param map[string]string) *UpdateExampleOptions {
	options.Headers = param
	return options
}

// UpdateIntentOptions : The updateIntent options.
type UpdateIntentOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The intent name.
	Intent *string `json:"intent" validate:"required"`

	// The name of the intent. This string must conform to the following restrictions:
	// - It can contain only Unicode alphanumeric, underscore, hyphen, and dot characters.
	// - It cannot begin with the reserved prefix `sys-`.
	// - It must be no longer than 128 characters.
	NewIntent *string `json:"new_intent,omitempty"`

	// An array of user input examples for the intent.
	NewExamples []CreateExample `json:"new_examples,omitempty"`

	// The description of the intent.
	NewDescription *string `json:"new_description,omitempty"`

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
func (options *UpdateIntentOptions) SetWorkspaceID(param string) *UpdateIntentOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetIntent : Allow user to set Intent
func (options *UpdateIntentOptions) SetIntent(param string) *UpdateIntentOptions {
	options.Intent = core.StringPtr(param)
	return options
}

// SetNewIntent : Allow user to set NewIntent
func (options *UpdateIntentOptions) SetNewIntent(param string) *UpdateIntentOptions {
	options.NewIntent = core.StringPtr(param)
	return options
}

// SetNewExamples : Allow user to set NewExamples
func (options *UpdateIntentOptions) SetNewExamples(param []CreateExample) *UpdateIntentOptions {
	options.NewExamples = param
	return options
}

// SetNewDescription : Allow user to set NewDescription
func (options *UpdateIntentOptions) SetNewDescription(param string) *UpdateIntentOptions {
	options.NewDescription = core.StringPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateIntentOptions) SetHeaders(param map[string]string) *UpdateIntentOptions {
	options.Headers = param
	return options
}

// UpdateSynonymOptions : The updateSynonym options.
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
	// - It must be no longer than 64 characters.
	NewSynonym *string `json:"new_synonym,omitempty"`

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
func (options *UpdateSynonymOptions) SetWorkspaceID(param string) *UpdateSynonymOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetEntity : Allow user to set Entity
func (options *UpdateSynonymOptions) SetEntity(param string) *UpdateSynonymOptions {
	options.Entity = core.StringPtr(param)
	return options
}

// SetValue : Allow user to set Value
func (options *UpdateSynonymOptions) SetValue(param string) *UpdateSynonymOptions {
	options.Value = core.StringPtr(param)
	return options
}

// SetSynonym : Allow user to set Synonym
func (options *UpdateSynonymOptions) SetSynonym(param string) *UpdateSynonymOptions {
	options.Synonym = core.StringPtr(param)
	return options
}

// SetNewSynonym : Allow user to set NewSynonym
func (options *UpdateSynonymOptions) SetNewSynonym(param string) *UpdateSynonymOptions {
	options.NewSynonym = core.StringPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateSynonymOptions) SetHeaders(param map[string]string) *UpdateSynonymOptions {
	options.Headers = param
	return options
}

// UpdateValueOptions : The updateValue options.
type UpdateValueOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The name of the entity.
	Entity *string `json:"entity" validate:"required"`

	// The text of the entity value.
	Value *string `json:"value" validate:"required"`

	// An array of synonyms for the entity value. You can provide either synonyms or patterns (as indicated by **type**),
	// but not both. A synonym must conform to the following resrictions:
	// - It cannot contain carriage return, newline, or tab characters.
	// - It cannot consist of only whitespace characters.
	// - It must be no longer than 64 characters.
	NewSynonyms []string `json:"new_synonyms,omitempty"`

	// Specifies the type of value.
	ValueType *string `json:"new_type,omitempty"`

	// Any metadata related to the entity value.
	NewMetadata interface{} `json:"new_metadata,omitempty"`

	// An array of patterns for the entity value. You can provide either synonyms or patterns (as indicated by **type**),
	// but not both. A pattern is a regular expression no longer than 512 characters. For more information about how to
	// specify a pattern, see the
	// [documentation](https://console.bluemix.net/docs/services/conversation/entities.html#creating-entities).
	NewPatterns []string `json:"new_patterns,omitempty"`

	// The text of the entity value. This string must conform to the following restrictions:
	// - It cannot contain carriage return, newline, or tab characters.
	// - It cannot consist of only whitespace characters.
	// - It must be no longer than 64 characters.
	NewValue *string `json:"new_value,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewUpdateValueOptions : Instantiate UpdateValueOptions
func (assistant *AssistantV1) NewUpdateValueOptions(workspaceID string, entity string, value string) *UpdateValueOptions {
	return &UpdateValueOptions{
		WorkspaceID: core.StringPtr(workspaceID),
		Entity:      core.StringPtr(entity),
		Value:       core.StringPtr(value),
	}
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *UpdateValueOptions) SetWorkspaceID(param string) *UpdateValueOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetEntity : Allow user to set Entity
func (options *UpdateValueOptions) SetEntity(param string) *UpdateValueOptions {
	options.Entity = core.StringPtr(param)
	return options
}

// SetValue : Allow user to set Value
func (options *UpdateValueOptions) SetValue(param string) *UpdateValueOptions {
	options.Value = core.StringPtr(param)
	return options
}

// SetNewSynonyms : Allow user to set NewSynonyms
func (options *UpdateValueOptions) SetNewSynonyms(param []string) *UpdateValueOptions {
	options.NewSynonyms = param
	return options
}

// SetValueType : Allow user to set ValueType
func (options *UpdateValueOptions) SetValueType(param string) *UpdateValueOptions {
	options.ValueType = core.StringPtr(param)
	return options
}

// SetNewMetadata : Allow user to set NewMetadata
func (options *UpdateValueOptions) SetNewMetadata(param interface{}) *UpdateValueOptions {
	options.NewMetadata = &param
	return options
}

// SetNewPatterns : Allow user to set NewPatterns
func (options *UpdateValueOptions) SetNewPatterns(param []string) *UpdateValueOptions {
	options.NewPatterns = param
	return options
}

// SetNewValue : Allow user to set NewValue
func (options *UpdateValueOptions) SetNewValue(param string) *UpdateValueOptions {
	options.NewValue = core.StringPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateValueOptions) SetHeaders(param map[string]string) *UpdateValueOptions {
	options.Headers = param
	return options
}

// UpdateWorkspaceOptions : The updateWorkspace options.
type UpdateWorkspaceOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The name of the workspace. This string cannot contain carriage return, newline, or tab characters, and it must be no
	// longer than 64 characters.
	Name *string `json:"name,omitempty"`

	// The description of the workspace. This string cannot contain carriage return, newline, or tab characters, and it
	// must be no longer than 128 characters.
	Description *string `json:"description,omitempty"`

	// The language of the workspace.
	Language *string `json:"language,omitempty"`

	// An array of objects defining the intents for the workspace.
	Intents []CreateIntent `json:"intents,omitempty"`

	// An array of objects defining the entities for the workspace.
	Entities []CreateEntity `json:"entities,omitempty"`

	// An array of objects defining the nodes in the workspace dialog.
	DialogNodes []CreateDialogNode `json:"dialog_nodes,omitempty"`

	// An array of objects defining input examples that have been marked as irrelevant input.
	Counterexamples []CreateCounterexample `json:"counterexamples,omitempty"`

	// Any metadata related to the workspace.
	Metadata interface{} `json:"metadata,omitempty"`

	// Whether training data from the workspace can be used by IBM for general service improvements. `true` indicates that
	// workspace training data is not to be used.
	LearningOptOut *bool `json:"learning_opt_out,omitempty"`

	// Global settings for the workspace.
	SystemSettings *WorkspaceSystemSettings `json:"system_settings,omitempty"`

	// Whether the new data is to be appended to the existing data in the workspace. If **append**=`false`, elements
	// included in the new data completely replace the corresponding existing elements, including all subelements. For
	// example, if the new data includes **entities** and **append**=`false`, all existing entities in the workspace are
	// discarded and replaced with the new entities.
	//
	// If **append**=`true`, existing elements are preserved, and the new elements are added. If any elements in the new
	// data collide with existing elements, the update request fails.
	Append *bool `json:"append,omitempty"`

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
func (options *UpdateWorkspaceOptions) SetWorkspaceID(param string) *UpdateWorkspaceOptions {
	options.WorkspaceID = core.StringPtr(param)
	return options
}

// SetName : Allow user to set Name
func (options *UpdateWorkspaceOptions) SetName(param string) *UpdateWorkspaceOptions {
	options.Name = core.StringPtr(param)
	return options
}

// SetDescription : Allow user to set Description
func (options *UpdateWorkspaceOptions) SetDescription(param string) *UpdateWorkspaceOptions {
	options.Description = core.StringPtr(param)
	return options
}

// SetLanguage : Allow user to set Language
func (options *UpdateWorkspaceOptions) SetLanguage(param string) *UpdateWorkspaceOptions {
	options.Language = core.StringPtr(param)
	return options
}

// SetIntents : Allow user to set Intents
func (options *UpdateWorkspaceOptions) SetIntents(param []CreateIntent) *UpdateWorkspaceOptions {
	options.Intents = param
	return options
}

// SetEntities : Allow user to set Entities
func (options *UpdateWorkspaceOptions) SetEntities(param []CreateEntity) *UpdateWorkspaceOptions {
	options.Entities = param
	return options
}

// SetDialogNodes : Allow user to set DialogNodes
func (options *UpdateWorkspaceOptions) SetDialogNodes(param []CreateDialogNode) *UpdateWorkspaceOptions {
	options.DialogNodes = param
	return options
}

// SetCounterexamples : Allow user to set Counterexamples
func (options *UpdateWorkspaceOptions) SetCounterexamples(param []CreateCounterexample) *UpdateWorkspaceOptions {
	options.Counterexamples = param
	return options
}

// SetMetadata : Allow user to set Metadata
func (options *UpdateWorkspaceOptions) SetMetadata(param interface{}) *UpdateWorkspaceOptions {
	options.Metadata = &param
	return options
}

// SetLearningOptOut : Allow user to set LearningOptOut
func (options *UpdateWorkspaceOptions) SetLearningOptOut(param bool) *UpdateWorkspaceOptions {
	options.LearningOptOut = core.BoolPtr(param)
	return options
}

// SetSystemSettings : Allow user to set SystemSettings
func (options *UpdateWorkspaceOptions) SetSystemSettings(param WorkspaceSystemSettings) *UpdateWorkspaceOptions {
	options.SystemSettings = &param
	return options
}

// SetAppend : Allow user to set Append
func (options *UpdateWorkspaceOptions) SetAppend(param bool) *UpdateWorkspaceOptions {
	options.Append = core.BoolPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateWorkspaceOptions) SetHeaders(param map[string]string) *UpdateWorkspaceOptions {
	options.Headers = param
	return options
}

// Value : Value struct
type Value struct {

	// The text of the entity value.
	ValueText *string `json:"value" validate:"required"`

	// Any metadata related to the entity value.
	Metadata interface{} `json:"metadata,omitempty"`

	// The timestamp for creation of the entity value.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// The timestamp for the last update to the entity value.
	Updated *strfmt.DateTime `json:"updated,omitempty"`

	// An array containing any synonyms for the entity value.
	Synonyms []string `json:"synonyms,omitempty"`

	// An array containing any patterns for the entity value.
	Patterns []string `json:"patterns,omitempty"`

	// Specifies the type of value.
	ValueType *string `json:"type" validate:"required"`
}

// ValueCollection : ValueCollection struct
type ValueCollection struct {

	// An array of entity values.
	Values []ValueExport `json:"values" validate:"required"`

	// An object defining the pagination data for the returned objects.
	Pagination *Pagination `json:"pagination" validate:"required"`
}

// ValueExport : ValueExport struct
type ValueExport struct {

	// The text of the entity value.
	ValueText *string `json:"value" validate:"required"`

	// Any metadata related to the entity value.
	Metadata interface{} `json:"metadata,omitempty"`

	// The timestamp for creation of the entity value.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// The timestamp for the last update to the entity value.
	Updated *strfmt.DateTime `json:"updated,omitempty"`

	// An array containing any synonyms for the entity value.
	Synonyms []string `json:"synonyms,omitempty"`

	// An array containing any patterns for the entity value.
	Patterns []string `json:"patterns,omitempty"`

	// Specifies the type of value.
	ValueType *string `json:"type" validate:"required"`
}

// Workspace : Workspace struct
type Workspace struct {

	// The name of the workspace.
	Name *string `json:"name" validate:"required"`

	// The language of the workspace.
	Language *string `json:"language" validate:"required"`

	// The timestamp for creation of the workspace.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// The timestamp for the last update to the workspace.
	Updated *strfmt.DateTime `json:"updated,omitempty"`

	// The workspace ID.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The description of the workspace.
	Description *string `json:"description,omitempty"`

	// Any metadata related to the workspace.
	Metadata interface{} `json:"metadata,omitempty"`

	// Whether training data from the workspace (including artifacts such as intents and entities) can be used by IBM for
	// general service improvements. `true` indicates that workspace training data is not to be used.
	LearningOptOut *bool `json:"learning_opt_out,omitempty"`

	// Global settings for the workspace.
	SystemSettings *WorkspaceSystemSettings `json:"system_settings,omitempty"`
}

// WorkspaceCollection : WorkspaceCollection struct
type WorkspaceCollection struct {

	// An array of objects describing the workspaces associated with the service instance.
	Workspaces []Workspace `json:"workspaces" validate:"required"`

	// An object defining the pagination data for the returned objects.
	Pagination *Pagination `json:"pagination" validate:"required"`
}

// WorkspaceExport : WorkspaceExport struct
type WorkspaceExport struct {

	// The name of the workspace.
	Name *string `json:"name" validate:"required"`

	// The description of the workspace.
	Description *string `json:"description" validate:"required"`

	// The language of the workspace.
	Language *string `json:"language" validate:"required"`

	// Any metadata that is required by the workspace.
	Metadata interface{} `json:"metadata" validate:"required"`

	// The timestamp for creation of the workspace.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// The timestamp for the last update to the workspace.
	Updated *strfmt.DateTime `json:"updated,omitempty"`

	// The workspace ID.
	WorkspaceID *string `json:"workspace_id" validate:"required"`

	// The current status of the workspace.
	Status *string `json:"status" validate:"required"`

	// Whether training data from the workspace can be used by IBM for general service improvements. `true` indicates that
	// workspace training data is not to be used.
	LearningOptOut *bool `json:"learning_opt_out" validate:"required"`

	// Global settings for the workspace.
	SystemSettings *WorkspaceSystemSettings `json:"system_settings,omitempty"`

	// An array of intents.
	Intents []IntentExport `json:"intents,omitempty"`

	// An array of entities.
	Entities []EntityExport `json:"entities,omitempty"`

	// An array of counterexamples.
	Counterexamples []Counterexample `json:"counterexamples,omitempty"`

	// An array of objects describing the dialog nodes in the workspace.
	DialogNodes []DialogNode `json:"dialog_nodes,omitempty"`
}

// WorkspaceSystemSettings : WorkspaceSystemSettings struct
type WorkspaceSystemSettings struct {

	// Workspace settings related to the Watson Assistant tool.
	Tooling *WorkspaceSystemSettingsTooling `json:"tooling,omitempty"`

	// Workspace settings related to the disambiguation feature.
	//
	// **Note:** This feature is available only to Premium users.
	Disambiguation *WorkspaceSystemSettingsDisambiguation `json:"disambiguation,omitempty"`

	// For internal use only.
	HumanAgentAssist interface{} `json:"human_agent_assist,omitempty"`
}

// WorkspaceSystemSettingsDisambiguation : WorkspaceSystemSettingsDisambiguation struct
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
}

// WorkspaceSystemSettingsTooling : WorkspaceSystemSettingsTooling struct
type WorkspaceSystemSettingsTooling struct {

	// Whether the dialog JSON editor displays text responses within the `output.generic` object.
	StoreGenericResponses *bool `json:"store_generic_responses,omitempty"`
}
