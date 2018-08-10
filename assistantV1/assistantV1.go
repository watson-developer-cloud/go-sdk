// Package assistantV1 : Operations and models for the AssistantV1 service
package assistantV1
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
    "bytes"
    "fmt"
    "github.com/go-openapi/strfmt"
    "runtime"
    "strings"
    req "github.com/parnurzeal/gorequest"
    watson "go-sdk"
)

// AssistantV1 : The AssistantV1 service
type AssistantV1 struct {
    client *watson.Client
}

// ServiceCredentials : Service credentials
type ServiceCredentials struct {
    ServiceURL string
    Version string
    Username string
    Password string
    APIkey string
    IAMtoken string
}

// NewAssistantV1 : Instantiate AssistantV1
func NewAssistantV1(serviceCreds *ServiceCredentials) (*AssistantV1, error) {
    if serviceCreds.ServiceURL == "" {
        serviceCreds.ServiceURL = "https://gateway.watsonplatform.net/assistant/api"
    }

    creds := watson.Credentials(*serviceCreds)
    client, clientErr := watson.NewClient(&creds, "conversation")

    if clientErr != nil {
        return nil, clientErr
    }

    return &AssistantV1{ client: client }, nil
}

// Message : Get response to user input
func (assistant *AssistantV1) Message(options *MessageOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/message"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    if options.IsNodesVisitedDetailsSet {
        request.Query("nodes_visited_details=" + fmt.Sprint(options.NodesVisitedDetails))
    }
    body := map[string]interface{}{}
    if options.IsInputSet {
        body["input"] = options.Input
    }
    if options.IsAlternateIntentsSet {
        body["alternate_intents"] = options.AlternateIntents
    }
    if options.IsContextSet {
        body["context"] = options.Context
    }
    if options.IsEntitiesSet {
        body["entities"] = options.Entities
    }
    if options.IsIntentsSet {
        body["intents"] = options.Intents
    }
    if options.IsOutputSet {
        body["output"] = options.Output
    }
    request.Send(body)

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(MessageResponse)
    res, _, err := request.EndStruct(&response.Result)

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetMessageResult : Cast result of Message operation
func GetMessageResult(response *watson.WatsonResponse) *MessageResponse {
    result, ok := response.Result.(*MessageResponse)

    if ok {
        return result
    }

    return nil
}

// CreateWorkspace : Create workspace
func (assistant *AssistantV1) CreateWorkspace(options *CreateWorkspaceOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    body := map[string]interface{}{}
    if options.IsNameSet {
        body["name"] = options.Name
    }
    if options.IsDescriptionSet {
        body["description"] = options.Description
    }
    if options.IsLanguageSet {
        body["language"] = options.Language
    }
    if options.IsIntentsSet {
        body["intents"] = options.Intents
    }
    if options.IsEntitiesSet {
        body["entities"] = options.Entities
    }
    if options.IsDialogNodesSet {
        body["dialog_nodes"] = options.DialogNodes
    }
    if options.IsCounterexamplesSet {
        body["counterexamples"] = options.Counterexamples
    }
    if options.IsMetadataSet {
        body["metadata"] = options.Metadata
    }
    if options.IsLearningOptOutSet {
        body["learning_opt_out"] = options.LearningOptOut
    }
    if options.IsSystemSettingsSet {
        body["system_settings"] = options.SystemSettings
    }
    request.Send(body)

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(Workspace)
    res, _, err := request.EndStruct(&response.Result)

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetCreateWorkspaceResult : Cast result of CreateWorkspace operation
func GetCreateWorkspaceResult(response *watson.WatsonResponse) *Workspace {
    result, ok := response.Result.(*Workspace)

    if ok {
        return result
    }

    return nil
}

// DeleteWorkspace : Delete workspace
func (assistant *AssistantV1) DeleteWorkspace(options *DeleteWorkspaceOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    res, _, err := request.End()

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}


// GetWorkspace : Get information about a workspace
func (assistant *AssistantV1) GetWorkspace(options *GetWorkspaceOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    if options.IsExportSet {
        request.Query("export=" + fmt.Sprint(options.Export))
    }
    if options.IsIncludeAuditSet {
        request.Query("include_audit=" + fmt.Sprint(options.IncludeAudit))
    }

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(WorkspaceExport)
    res, _, err := request.EndStruct(&response.Result)

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetGetWorkspaceResult : Cast result of GetWorkspace operation
func GetGetWorkspaceResult(response *watson.WatsonResponse) *WorkspaceExport {
    result, ok := response.Result.(*WorkspaceExport)

    if ok {
        return result
    }

    return nil
}

// ListWorkspaces : List workspaces
func (assistant *AssistantV1) ListWorkspaces(options *ListWorkspacesOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    if options.IsPageLimitSet {
        request.Query("page_limit=" + fmt.Sprint(options.PageLimit))
    }
    if options.IsIncludeCountSet {
        request.Query("include_count=" + fmt.Sprint(options.IncludeCount))
    }
    if options.IsSortSet {
        request.Query("sort=" + fmt.Sprint(options.Sort))
    }
    if options.IsCursorSet {
        request.Query("cursor=" + fmt.Sprint(options.Cursor))
    }
    if options.IsIncludeAuditSet {
        request.Query("include_audit=" + fmt.Sprint(options.IncludeAudit))
    }

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(WorkspaceCollection)
    res, _, err := request.EndStruct(&response.Result)

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetListWorkspacesResult : Cast result of ListWorkspaces operation
func GetListWorkspacesResult(response *watson.WatsonResponse) *WorkspaceCollection {
    result, ok := response.Result.(*WorkspaceCollection)

    if ok {
        return result
    }

    return nil
}

// UpdateWorkspace : Update workspace
func (assistant *AssistantV1) UpdateWorkspace(options *UpdateWorkspaceOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    if options.IsAppendVarSet {
        request.Query("append=" + fmt.Sprint(options.AppendVar))
    }
    body := map[string]interface{}{}
    if options.IsNameSet {
        body["name"] = options.Name
    }
    if options.IsDescriptionSet {
        body["description"] = options.Description
    }
    if options.IsLanguageSet {
        body["language"] = options.Language
    }
    if options.IsIntentsSet {
        body["intents"] = options.Intents
    }
    if options.IsEntitiesSet {
        body["entities"] = options.Entities
    }
    if options.IsDialogNodesSet {
        body["dialog_nodes"] = options.DialogNodes
    }
    if options.IsCounterexamplesSet {
        body["counterexamples"] = options.Counterexamples
    }
    if options.IsMetadataSet {
        body["metadata"] = options.Metadata
    }
    if options.IsLearningOptOutSet {
        body["learning_opt_out"] = options.LearningOptOut
    }
    if options.IsSystemSettingsSet {
        body["system_settings"] = options.SystemSettings
    }
    request.Send(body)

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(Workspace)
    res, _, err := request.EndStruct(&response.Result)

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetUpdateWorkspaceResult : Cast result of UpdateWorkspace operation
func GetUpdateWorkspaceResult(response *watson.WatsonResponse) *Workspace {
    result, ok := response.Result.(*Workspace)

    if ok {
        return result
    }

    return nil
}

// CreateIntent : Create intent
func (assistant *AssistantV1) CreateIntent(options *CreateIntentOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/intents"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    body := map[string]interface{}{}
    body["intent"] = options.Intent
    if options.IsDescriptionSet {
        body["description"] = options.Description
    }
    if options.IsExamplesSet {
        body["examples"] = options.Examples
    }
    request.Send(body)

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(Intent)
    res, _, err := request.EndStruct(&response.Result)

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetCreateIntentResult : Cast result of CreateIntent operation
func GetCreateIntentResult(response *watson.WatsonResponse) *Intent {
    result, ok := response.Result.(*Intent)

    if ok {
        return result
    }

    return nil
}

// DeleteIntent : Delete intent
func (assistant *AssistantV1) DeleteIntent(options *DeleteIntentOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/intents/{intent}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    path = strings.Replace(path, "{intent}", options.Intent, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    res, _, err := request.End()

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}


// GetIntent : Get intent
func (assistant *AssistantV1) GetIntent(options *GetIntentOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/intents/{intent}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    path = strings.Replace(path, "{intent}", options.Intent, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    if options.IsExportSet {
        request.Query("export=" + fmt.Sprint(options.Export))
    }
    if options.IsIncludeAuditSet {
        request.Query("include_audit=" + fmt.Sprint(options.IncludeAudit))
    }

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(IntentExport)
    res, _, err := request.EndStruct(&response.Result)

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetGetIntentResult : Cast result of GetIntent operation
func GetGetIntentResult(response *watson.WatsonResponse) *IntentExport {
    result, ok := response.Result.(*IntentExport)

    if ok {
        return result
    }

    return nil
}

// ListIntents : List intents
func (assistant *AssistantV1) ListIntents(options *ListIntentsOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/intents"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    if options.IsExportSet {
        request.Query("export=" + fmt.Sprint(options.Export))
    }
    if options.IsPageLimitSet {
        request.Query("page_limit=" + fmt.Sprint(options.PageLimit))
    }
    if options.IsIncludeCountSet {
        request.Query("include_count=" + fmt.Sprint(options.IncludeCount))
    }
    if options.IsSortSet {
        request.Query("sort=" + fmt.Sprint(options.Sort))
    }
    if options.IsCursorSet {
        request.Query("cursor=" + fmt.Sprint(options.Cursor))
    }
    if options.IsIncludeAuditSet {
        request.Query("include_audit=" + fmt.Sprint(options.IncludeAudit))
    }

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(IntentCollection)
    res, _, err := request.EndStruct(&response.Result)

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetListIntentsResult : Cast result of ListIntents operation
func GetListIntentsResult(response *watson.WatsonResponse) *IntentCollection {
    result, ok := response.Result.(*IntentCollection)

    if ok {
        return result
    }

    return nil
}

// UpdateIntent : Update intent
func (assistant *AssistantV1) UpdateIntent(options *UpdateIntentOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/intents/{intent}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    path = strings.Replace(path, "{intent}", options.Intent, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    body := map[string]interface{}{}
    if options.IsNewIntentSet {
        body["intent"] = options.NewIntent
    }
    if options.IsNewExamplesSet {
        body["examples"] = options.NewExamples
    }
    if options.IsNewDescriptionSet {
        body["description"] = options.NewDescription
    }
    request.Send(body)

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(Intent)
    res, _, err := request.EndStruct(&response.Result)

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetUpdateIntentResult : Cast result of UpdateIntent operation
func GetUpdateIntentResult(response *watson.WatsonResponse) *Intent {
    result, ok := response.Result.(*Intent)

    if ok {
        return result
    }

    return nil
}

// CreateExample : Create user input example
func (assistant *AssistantV1) CreateExample(options *CreateExampleOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/intents/{intent}/examples"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    path = strings.Replace(path, "{intent}", options.Intent, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    body := map[string]interface{}{}
    body["text"] = options.Text
    if options.IsMentionsSet {
        body["mentions"] = options.Mentions
    }
    request.Send(body)

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(Example)
    res, _, err := request.EndStruct(&response.Result)

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetCreateExampleResult : Cast result of CreateExample operation
func GetCreateExampleResult(response *watson.WatsonResponse) *Example {
    result, ok := response.Result.(*Example)

    if ok {
        return result
    }

    return nil
}

// DeleteExample : Delete user input example
func (assistant *AssistantV1) DeleteExample(options *DeleteExampleOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/intents/{intent}/examples/{text}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    path = strings.Replace(path, "{intent}", options.Intent, 1)
    path = strings.Replace(path, "{text}", options.Text, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    res, _, err := request.End()

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}


// GetExample : Get user input example
func (assistant *AssistantV1) GetExample(options *GetExampleOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/intents/{intent}/examples/{text}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    path = strings.Replace(path, "{intent}", options.Intent, 1)
    path = strings.Replace(path, "{text}", options.Text, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    if options.IsIncludeAuditSet {
        request.Query("include_audit=" + fmt.Sprint(options.IncludeAudit))
    }

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(Example)
    res, _, err := request.EndStruct(&response.Result)

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetGetExampleResult : Cast result of GetExample operation
func GetGetExampleResult(response *watson.WatsonResponse) *Example {
    result, ok := response.Result.(*Example)

    if ok {
        return result
    }

    return nil
}

// ListExamples : List user input examples
func (assistant *AssistantV1) ListExamples(options *ListExamplesOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/intents/{intent}/examples"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    path = strings.Replace(path, "{intent}", options.Intent, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    if options.IsPageLimitSet {
        request.Query("page_limit=" + fmt.Sprint(options.PageLimit))
    }
    if options.IsIncludeCountSet {
        request.Query("include_count=" + fmt.Sprint(options.IncludeCount))
    }
    if options.IsSortSet {
        request.Query("sort=" + fmt.Sprint(options.Sort))
    }
    if options.IsCursorSet {
        request.Query("cursor=" + fmt.Sprint(options.Cursor))
    }
    if options.IsIncludeAuditSet {
        request.Query("include_audit=" + fmt.Sprint(options.IncludeAudit))
    }

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(ExampleCollection)
    res, _, err := request.EndStruct(&response.Result)

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetListExamplesResult : Cast result of ListExamples operation
func GetListExamplesResult(response *watson.WatsonResponse) *ExampleCollection {
    result, ok := response.Result.(*ExampleCollection)

    if ok {
        return result
    }

    return nil
}

// UpdateExample : Update user input example
func (assistant *AssistantV1) UpdateExample(options *UpdateExampleOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/intents/{intent}/examples/{text}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    path = strings.Replace(path, "{intent}", options.Intent, 1)
    path = strings.Replace(path, "{text}", options.Text, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    body := map[string]interface{}{}
    if options.IsNewTextSet {
        body["text"] = options.NewText
    }
    if options.IsNewMentionsSet {
        body["mentions"] = options.NewMentions
    }
    request.Send(body)

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(Example)
    res, _, err := request.EndStruct(&response.Result)

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetUpdateExampleResult : Cast result of UpdateExample operation
func GetUpdateExampleResult(response *watson.WatsonResponse) *Example {
    result, ok := response.Result.(*Example)

    if ok {
        return result
    }

    return nil
}

// CreateCounterexample : Create counterexample
func (assistant *AssistantV1) CreateCounterexample(options *CreateCounterexampleOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/counterexamples"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    body := map[string]interface{}{}
    body["text"] = options.Text
    request.Send(body)

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(Counterexample)
    res, _, err := request.EndStruct(&response.Result)

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetCreateCounterexampleResult : Cast result of CreateCounterexample operation
func GetCreateCounterexampleResult(response *watson.WatsonResponse) *Counterexample {
    result, ok := response.Result.(*Counterexample)

    if ok {
        return result
    }

    return nil
}

// DeleteCounterexample : Delete counterexample
func (assistant *AssistantV1) DeleteCounterexample(options *DeleteCounterexampleOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/counterexamples/{text}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    path = strings.Replace(path, "{text}", options.Text, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    res, _, err := request.End()

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}


// GetCounterexample : Get counterexample
func (assistant *AssistantV1) GetCounterexample(options *GetCounterexampleOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/counterexamples/{text}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    path = strings.Replace(path, "{text}", options.Text, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    if options.IsIncludeAuditSet {
        request.Query("include_audit=" + fmt.Sprint(options.IncludeAudit))
    }

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(Counterexample)
    res, _, err := request.EndStruct(&response.Result)

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetGetCounterexampleResult : Cast result of GetCounterexample operation
func GetGetCounterexampleResult(response *watson.WatsonResponse) *Counterexample {
    result, ok := response.Result.(*Counterexample)

    if ok {
        return result
    }

    return nil
}

// ListCounterexamples : List counterexamples
func (assistant *AssistantV1) ListCounterexamples(options *ListCounterexamplesOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/counterexamples"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    if options.IsPageLimitSet {
        request.Query("page_limit=" + fmt.Sprint(options.PageLimit))
    }
    if options.IsIncludeCountSet {
        request.Query("include_count=" + fmt.Sprint(options.IncludeCount))
    }
    if options.IsSortSet {
        request.Query("sort=" + fmt.Sprint(options.Sort))
    }
    if options.IsCursorSet {
        request.Query("cursor=" + fmt.Sprint(options.Cursor))
    }
    if options.IsIncludeAuditSet {
        request.Query("include_audit=" + fmt.Sprint(options.IncludeAudit))
    }

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(CounterexampleCollection)
    res, _, err := request.EndStruct(&response.Result)

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetListCounterexamplesResult : Cast result of ListCounterexamples operation
func GetListCounterexamplesResult(response *watson.WatsonResponse) *CounterexampleCollection {
    result, ok := response.Result.(*CounterexampleCollection)

    if ok {
        return result
    }

    return nil
}

// UpdateCounterexample : Update counterexample
func (assistant *AssistantV1) UpdateCounterexample(options *UpdateCounterexampleOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/counterexamples/{text}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    path = strings.Replace(path, "{text}", options.Text, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    body := map[string]interface{}{}
    if options.IsNewTextSet {
        body["text"] = options.NewText
    }
    request.Send(body)

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(Counterexample)
    res, _, err := request.EndStruct(&response.Result)

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetUpdateCounterexampleResult : Cast result of UpdateCounterexample operation
func GetUpdateCounterexampleResult(response *watson.WatsonResponse) *Counterexample {
    result, ok := response.Result.(*Counterexample)

    if ok {
        return result
    }

    return nil
}

// CreateEntity : Create entity
func (assistant *AssistantV1) CreateEntity(options *CreateEntityOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/entities"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    body := map[string]interface{}{}
    body["entity"] = options.Entity
    if options.IsDescriptionSet {
        body["description"] = options.Description
    }
    if options.IsMetadataSet {
        body["metadata"] = options.Metadata
    }
    if options.IsValuesSet {
        body["values"] = options.Values
    }
    if options.IsFuzzyMatchSet {
        body["fuzzy_match"] = options.FuzzyMatch
    }
    request.Send(body)

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(Entity)
    res, _, err := request.EndStruct(&response.Result)

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetCreateEntityResult : Cast result of CreateEntity operation
func GetCreateEntityResult(response *watson.WatsonResponse) *Entity {
    result, ok := response.Result.(*Entity)

    if ok {
        return result
    }

    return nil
}

// DeleteEntity : Delete entity
func (assistant *AssistantV1) DeleteEntity(options *DeleteEntityOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/entities/{entity}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    path = strings.Replace(path, "{entity}", options.Entity, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    res, _, err := request.End()

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}


// GetEntity : Get entity
func (assistant *AssistantV1) GetEntity(options *GetEntityOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/entities/{entity}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    path = strings.Replace(path, "{entity}", options.Entity, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    if options.IsExportSet {
        request.Query("export=" + fmt.Sprint(options.Export))
    }
    if options.IsIncludeAuditSet {
        request.Query("include_audit=" + fmt.Sprint(options.IncludeAudit))
    }

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(EntityExport)
    res, _, err := request.EndStruct(&response.Result)

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetGetEntityResult : Cast result of GetEntity operation
func GetGetEntityResult(response *watson.WatsonResponse) *EntityExport {
    result, ok := response.Result.(*EntityExport)

    if ok {
        return result
    }

    return nil
}

// ListEntities : List entities
func (assistant *AssistantV1) ListEntities(options *ListEntitiesOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/entities"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    if options.IsExportSet {
        request.Query("export=" + fmt.Sprint(options.Export))
    }
    if options.IsPageLimitSet {
        request.Query("page_limit=" + fmt.Sprint(options.PageLimit))
    }
    if options.IsIncludeCountSet {
        request.Query("include_count=" + fmt.Sprint(options.IncludeCount))
    }
    if options.IsSortSet {
        request.Query("sort=" + fmt.Sprint(options.Sort))
    }
    if options.IsCursorSet {
        request.Query("cursor=" + fmt.Sprint(options.Cursor))
    }
    if options.IsIncludeAuditSet {
        request.Query("include_audit=" + fmt.Sprint(options.IncludeAudit))
    }

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(EntityCollection)
    res, _, err := request.EndStruct(&response.Result)

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetListEntitiesResult : Cast result of ListEntities operation
func GetListEntitiesResult(response *watson.WatsonResponse) *EntityCollection {
    result, ok := response.Result.(*EntityCollection)

    if ok {
        return result
    }

    return nil
}

// UpdateEntity : Update entity
func (assistant *AssistantV1) UpdateEntity(options *UpdateEntityOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/entities/{entity}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    path = strings.Replace(path, "{entity}", options.Entity, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    body := map[string]interface{}{}
    if options.IsNewFuzzyMatchSet {
        body["fuzzy_match"] = options.NewFuzzyMatch
    }
    if options.IsNewEntitySet {
        body["entity"] = options.NewEntity
    }
    if options.IsNewMetadataSet {
        body["metadata"] = options.NewMetadata
    }
    if options.IsNewValuesSet {
        body["values"] = options.NewValues
    }
    if options.IsNewDescriptionSet {
        body["description"] = options.NewDescription
    }
    request.Send(body)

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(Entity)
    res, _, err := request.EndStruct(&response.Result)

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetUpdateEntityResult : Cast result of UpdateEntity operation
func GetUpdateEntityResult(response *watson.WatsonResponse) *Entity {
    result, ok := response.Result.(*Entity)

    if ok {
        return result
    }

    return nil
}

// ListEntityMentions : List entity mentions
func (assistant *AssistantV1) ListEntityMentions(options *ListEntityMentionsOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/entities/{entity}/mentions"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    path = strings.Replace(path, "{entity}", options.Entity, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    if options.IsExportSet {
        request.Query("export=" + fmt.Sprint(options.Export))
    }
    if options.IsIncludeAuditSet {
        request.Query("include_audit=" + fmt.Sprint(options.IncludeAudit))
    }

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(EntityMentionCollection)
    res, _, err := request.EndStruct(&response.Result)

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetListEntityMentionsResult : Cast result of ListEntityMentions operation
func GetListEntityMentionsResult(response *watson.WatsonResponse) *EntityMentionCollection {
    result, ok := response.Result.(*EntityMentionCollection)

    if ok {
        return result
    }

    return nil
}

// CreateValue : Add entity value
func (assistant *AssistantV1) CreateValue(options *CreateValueOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/entities/{entity}/values"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    path = strings.Replace(path, "{entity}", options.Entity, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    body := map[string]interface{}{}
    body["value"] = options.Value
    if options.IsMetadataSet {
        body["metadata"] = options.Metadata
    }
    if options.IsSynonymsSet {
        body["synonyms"] = options.Synonyms
    }
    if options.IsPatternsSet {
        body["patterns"] = options.Patterns
    }
    if options.IsValueTypeSet {
        body["type"] = options.ValueType
    }
    request.Send(body)

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(Value)
    res, _, err := request.EndStruct(&response.Result)

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetCreateValueResult : Cast result of CreateValue operation
func GetCreateValueResult(response *watson.WatsonResponse) *Value {
    result, ok := response.Result.(*Value)

    if ok {
        return result
    }

    return nil
}

// DeleteValue : Delete entity value
func (assistant *AssistantV1) DeleteValue(options *DeleteValueOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    path = strings.Replace(path, "{entity}", options.Entity, 1)
    path = strings.Replace(path, "{value}", options.Value, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    res, _, err := request.End()

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}


// GetValue : Get entity value
func (assistant *AssistantV1) GetValue(options *GetValueOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    path = strings.Replace(path, "{entity}", options.Entity, 1)
    path = strings.Replace(path, "{value}", options.Value, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    if options.IsExportSet {
        request.Query("export=" + fmt.Sprint(options.Export))
    }
    if options.IsIncludeAuditSet {
        request.Query("include_audit=" + fmt.Sprint(options.IncludeAudit))
    }

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(ValueExport)
    res, _, err := request.EndStruct(&response.Result)

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetGetValueResult : Cast result of GetValue operation
func GetGetValueResult(response *watson.WatsonResponse) *ValueExport {
    result, ok := response.Result.(*ValueExport)

    if ok {
        return result
    }

    return nil
}

// ListValues : List entity values
func (assistant *AssistantV1) ListValues(options *ListValuesOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/entities/{entity}/values"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    path = strings.Replace(path, "{entity}", options.Entity, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    if options.IsExportSet {
        request.Query("export=" + fmt.Sprint(options.Export))
    }
    if options.IsPageLimitSet {
        request.Query("page_limit=" + fmt.Sprint(options.PageLimit))
    }
    if options.IsIncludeCountSet {
        request.Query("include_count=" + fmt.Sprint(options.IncludeCount))
    }
    if options.IsSortSet {
        request.Query("sort=" + fmt.Sprint(options.Sort))
    }
    if options.IsCursorSet {
        request.Query("cursor=" + fmt.Sprint(options.Cursor))
    }
    if options.IsIncludeAuditSet {
        request.Query("include_audit=" + fmt.Sprint(options.IncludeAudit))
    }

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(ValueCollection)
    res, _, err := request.EndStruct(&response.Result)

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetListValuesResult : Cast result of ListValues operation
func GetListValuesResult(response *watson.WatsonResponse) *ValueCollection {
    result, ok := response.Result.(*ValueCollection)

    if ok {
        return result
    }

    return nil
}

// UpdateValue : Update entity value
func (assistant *AssistantV1) UpdateValue(options *UpdateValueOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    path = strings.Replace(path, "{entity}", options.Entity, 1)
    path = strings.Replace(path, "{value}", options.Value, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    body := map[string]interface{}{}
    if options.IsValueTypeSet {
        body["type"] = options.ValueType
    }
    if options.IsNewSynonymsSet {
        body["synonyms"] = options.NewSynonyms
    }
    if options.IsNewMetadataSet {
        body["metadata"] = options.NewMetadata
    }
    if options.IsNewValueSet {
        body["value"] = options.NewValue
    }
    if options.IsNewPatternsSet {
        body["patterns"] = options.NewPatterns
    }
    request.Send(body)

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(Value)
    res, _, err := request.EndStruct(&response.Result)

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetUpdateValueResult : Cast result of UpdateValue operation
func GetUpdateValueResult(response *watson.WatsonResponse) *Value {
    result, ok := response.Result.(*Value)

    if ok {
        return result
    }

    return nil
}

// CreateSynonym : Add entity value synonym
func (assistant *AssistantV1) CreateSynonym(options *CreateSynonymOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}/synonyms"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    path = strings.Replace(path, "{entity}", options.Entity, 1)
    path = strings.Replace(path, "{value}", options.Value, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    body := map[string]interface{}{}
    body["synonym"] = options.Synonym
    request.Send(body)

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(Synonym)
    res, _, err := request.EndStruct(&response.Result)

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetCreateSynonymResult : Cast result of CreateSynonym operation
func GetCreateSynonymResult(response *watson.WatsonResponse) *Synonym {
    result, ok := response.Result.(*Synonym)

    if ok {
        return result
    }

    return nil
}

// DeleteSynonym : Delete entity value synonym
func (assistant *AssistantV1) DeleteSynonym(options *DeleteSynonymOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}/synonyms/{synonym}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    path = strings.Replace(path, "{entity}", options.Entity, 1)
    path = strings.Replace(path, "{value}", options.Value, 1)
    path = strings.Replace(path, "{synonym}", options.Synonym, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    res, _, err := request.End()

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}


// GetSynonym : Get entity value synonym
func (assistant *AssistantV1) GetSynonym(options *GetSynonymOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}/synonyms/{synonym}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    path = strings.Replace(path, "{entity}", options.Entity, 1)
    path = strings.Replace(path, "{value}", options.Value, 1)
    path = strings.Replace(path, "{synonym}", options.Synonym, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    if options.IsIncludeAuditSet {
        request.Query("include_audit=" + fmt.Sprint(options.IncludeAudit))
    }

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(Synonym)
    res, _, err := request.EndStruct(&response.Result)

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetGetSynonymResult : Cast result of GetSynonym operation
func GetGetSynonymResult(response *watson.WatsonResponse) *Synonym {
    result, ok := response.Result.(*Synonym)

    if ok {
        return result
    }

    return nil
}

// ListSynonyms : List entity value synonyms
func (assistant *AssistantV1) ListSynonyms(options *ListSynonymsOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}/synonyms"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    path = strings.Replace(path, "{entity}", options.Entity, 1)
    path = strings.Replace(path, "{value}", options.Value, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    if options.IsPageLimitSet {
        request.Query("page_limit=" + fmt.Sprint(options.PageLimit))
    }
    if options.IsIncludeCountSet {
        request.Query("include_count=" + fmt.Sprint(options.IncludeCount))
    }
    if options.IsSortSet {
        request.Query("sort=" + fmt.Sprint(options.Sort))
    }
    if options.IsCursorSet {
        request.Query("cursor=" + fmt.Sprint(options.Cursor))
    }
    if options.IsIncludeAuditSet {
        request.Query("include_audit=" + fmt.Sprint(options.IncludeAudit))
    }

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(SynonymCollection)
    res, _, err := request.EndStruct(&response.Result)

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetListSynonymsResult : Cast result of ListSynonyms operation
func GetListSynonymsResult(response *watson.WatsonResponse) *SynonymCollection {
    result, ok := response.Result.(*SynonymCollection)

    if ok {
        return result
    }

    return nil
}

// UpdateSynonym : Update entity value synonym
func (assistant *AssistantV1) UpdateSynonym(options *UpdateSynonymOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}/synonyms/{synonym}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    path = strings.Replace(path, "{entity}", options.Entity, 1)
    path = strings.Replace(path, "{value}", options.Value, 1)
    path = strings.Replace(path, "{synonym}", options.Synonym, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    body := map[string]interface{}{}
    if options.IsNewSynonymSet {
        body["synonym"] = options.NewSynonym
    }
    request.Send(body)

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(Synonym)
    res, _, err := request.EndStruct(&response.Result)

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetUpdateSynonymResult : Cast result of UpdateSynonym operation
func GetUpdateSynonymResult(response *watson.WatsonResponse) *Synonym {
    result, ok := response.Result.(*Synonym)

    if ok {
        return result
    }

    return nil
}

// CreateDialogNode : Create dialog node
func (assistant *AssistantV1) CreateDialogNode(options *CreateDialogNodeOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/dialog_nodes"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    body := map[string]interface{}{}
    body["dialog_node"] = options.DialogNode
    if options.IsDescriptionSet {
        body["description"] = options.Description
    }
    if options.IsConditionsSet {
        body["conditions"] = options.Conditions
    }
    if options.IsParentSet {
        body["parent"] = options.Parent
    }
    if options.IsPreviousSiblingSet {
        body["previous_sibling"] = options.PreviousSibling
    }
    if options.IsOutputSet {
        body["output"] = options.Output
    }
    if options.IsContextSet {
        body["context"] = options.Context
    }
    if options.IsMetadataSet {
        body["metadata"] = options.Metadata
    }
    if options.IsNextStepSet {
        body["next_step"] = options.NextStep
    }
    if options.IsActionsSet {
        body["actions"] = options.Actions
    }
    if options.IsTitleSet {
        body["title"] = options.Title
    }
    if options.IsNodeTypeSet {
        body["type"] = options.NodeType
    }
    if options.IsEventNameSet {
        body["event_name"] = options.EventName
    }
    if options.IsVariableSet {
        body["variable"] = options.Variable
    }
    if options.IsDigressInSet {
        body["digress_in"] = options.DigressIn
    }
    if options.IsDigressOutSet {
        body["digress_out"] = options.DigressOut
    }
    if options.IsDigressOutSlotsSet {
        body["digress_out_slots"] = options.DigressOutSlots
    }
    if options.IsUserLabelSet {
        body["user_label"] = options.UserLabel
    }
    request.Send(body)

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(DialogNode)
    res, _, err := request.EndStruct(&response.Result)

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetCreateDialogNodeResult : Cast result of CreateDialogNode operation
func GetCreateDialogNodeResult(response *watson.WatsonResponse) *DialogNode {
    result, ok := response.Result.(*DialogNode)

    if ok {
        return result
    }

    return nil
}

// DeleteDialogNode : Delete dialog node
func (assistant *AssistantV1) DeleteDialogNode(options *DeleteDialogNodeOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/dialog_nodes/{dialog_node}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    path = strings.Replace(path, "{dialog_node}", options.DialogNode, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    res, _, err := request.End()

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}


// GetDialogNode : Get dialog node
func (assistant *AssistantV1) GetDialogNode(options *GetDialogNodeOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/dialog_nodes/{dialog_node}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    path = strings.Replace(path, "{dialog_node}", options.DialogNode, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    if options.IsIncludeAuditSet {
        request.Query("include_audit=" + fmt.Sprint(options.IncludeAudit))
    }

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(DialogNode)
    res, _, err := request.EndStruct(&response.Result)

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetGetDialogNodeResult : Cast result of GetDialogNode operation
func GetGetDialogNodeResult(response *watson.WatsonResponse) *DialogNode {
    result, ok := response.Result.(*DialogNode)

    if ok {
        return result
    }

    return nil
}

// ListDialogNodes : List dialog nodes
func (assistant *AssistantV1) ListDialogNodes(options *ListDialogNodesOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/dialog_nodes"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    if options.IsPageLimitSet {
        request.Query("page_limit=" + fmt.Sprint(options.PageLimit))
    }
    if options.IsIncludeCountSet {
        request.Query("include_count=" + fmt.Sprint(options.IncludeCount))
    }
    if options.IsSortSet {
        request.Query("sort=" + fmt.Sprint(options.Sort))
    }
    if options.IsCursorSet {
        request.Query("cursor=" + fmt.Sprint(options.Cursor))
    }
    if options.IsIncludeAuditSet {
        request.Query("include_audit=" + fmt.Sprint(options.IncludeAudit))
    }

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(DialogNodeCollection)
    res, _, err := request.EndStruct(&response.Result)

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetListDialogNodesResult : Cast result of ListDialogNodes operation
func GetListDialogNodesResult(response *watson.WatsonResponse) *DialogNodeCollection {
    result, ok := response.Result.(*DialogNodeCollection)

    if ok {
        return result
    }

    return nil
}

// UpdateDialogNode : Update dialog node
func (assistant *AssistantV1) UpdateDialogNode(options *UpdateDialogNodeOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/dialog_nodes/{dialog_node}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    path = strings.Replace(path, "{dialog_node}", options.DialogNode, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    body := map[string]interface{}{}
    if options.IsNodeTypeSet {
        body["type"] = options.NodeType
    }
    if options.IsNewConditionsSet {
        body["conditions"] = options.NewConditions
    }
    if options.IsNewActionsSet {
        body["actions"] = options.NewActions
    }
    if options.IsNewPreviousSiblingSet {
        body["previous_sibling"] = options.NewPreviousSibling
    }
    if options.IsNewContextSet {
        body["context"] = options.NewContext
    }
    if options.IsNewVariableSet {
        body["variable"] = options.NewVariable
    }
    if options.IsNewUserLabelSet {
        body["user_label"] = options.NewUserLabel
    }
    if options.IsNewMetadataSet {
        body["metadata"] = options.NewMetadata
    }
    if options.IsNewTitleSet {
        body["title"] = options.NewTitle
    }
    if options.IsNewDescriptionSet {
        body["description"] = options.NewDescription
    }
    if options.IsNewDigressOutSet {
        body["digress_out"] = options.NewDigressOut
    }
    if options.IsNewEventNameSet {
        body["event_name"] = options.NewEventName
    }
    if options.IsNewDigressOutSlotsSet {
        body["digress_out_slots"] = options.NewDigressOutSlots
    }
    if options.IsNewNextStepSet {
        body["next_step"] = options.NewNextStep
    }
    if options.IsNewOutputSet {
        body["output"] = options.NewOutput
    }
    if options.IsNewDigressInSet {
        body["digress_in"] = options.NewDigressIn
    }
    if options.IsNewParentSet {
        body["parent"] = options.NewParent
    }
    if options.IsNewDialogNodeSet {
        body["dialog_node"] = options.NewDialogNode
    }
    request.Send(body)

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(DialogNode)
    res, _, err := request.EndStruct(&response.Result)

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetUpdateDialogNodeResult : Cast result of UpdateDialogNode operation
func GetUpdateDialogNodeResult(response *watson.WatsonResponse) *DialogNode {
    result, ok := response.Result.(*DialogNode)

    if ok {
        return result
    }

    return nil
}

// ListAllLogs : List log events in all workspaces
func (assistant *AssistantV1) ListAllLogs(options *ListAllLogsOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/logs"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("filter=" + fmt.Sprint(options.Filter))
    if options.IsSortSet {
        request.Query("sort=" + fmt.Sprint(options.Sort))
    }
    if options.IsPageLimitSet {
        request.Query("page_limit=" + fmt.Sprint(options.PageLimit))
    }
    if options.IsCursorSet {
        request.Query("cursor=" + fmt.Sprint(options.Cursor))
    }

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(LogCollection)
    res, _, err := request.EndStruct(&response.Result)

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetListAllLogsResult : Cast result of ListAllLogs operation
func GetListAllLogsResult(response *watson.WatsonResponse) *LogCollection {
    result, ok := response.Result.(*LogCollection)

    if ok {
        return result
    }

    return nil
}

// ListLogs : List log events in a workspace
func (assistant *AssistantV1) ListLogs(options *ListLogsOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/logs"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", options.WorkspaceID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    if options.IsSortSet {
        request.Query("sort=" + fmt.Sprint(options.Sort))
    }
    if options.IsFilterSet {
        request.Query("filter=" + fmt.Sprint(options.Filter))
    }
    if options.IsPageLimitSet {
        request.Query("page_limit=" + fmt.Sprint(options.PageLimit))
    }
    if options.IsCursorSet {
        request.Query("cursor=" + fmt.Sprint(options.Cursor))
    }

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(LogCollection)
    res, _, err := request.EndStruct(&response.Result)

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetListLogsResult : Cast result of ListLogs operation
func GetListLogsResult(response *watson.WatsonResponse) *LogCollection {
    result, ok := response.Result.(*LogCollection)

    if ok {
        return result
    }

    return nil
}

// DeleteUserData : Delete labeled data
func (assistant *AssistantV1) DeleteUserData(options *DeleteUserDataOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/user_data"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    request := req.New().Delete(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("User-Agent", "watson-apis-go-sdk 0.0.1 " + runtime.GOOS)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("customer_id=" + fmt.Sprint(options.CustomerID))

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    res, _, err := request.End()

    if err != nil {
        return nil, err
    }

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}



// CaptureGroup : CaptureGroup struct
type CaptureGroup struct {

	// A recognized capture group for the entity.
	Group string `json:"group"`

	// Zero-based character offsets that indicate where the entity value begins and ends in the input text.
	Location []int64 `json:"location,omitempty"`
}

// Context : State information for the conversation. To maintain state, include the context from the previous response.
type Context struct {

	// The unique identifier of the conversation.
	ConversationID string `json:"conversation_id,omitempty"`

	// For internal use only.
	System SystemResponse `json:"system,omitempty"`
}

// Counterexample : Counterexample struct
type Counterexample struct {

	// The text of the counterexample.
	Text string `json:"text"`

	// The timestamp for creation of the counterexample.
	Created strfmt.DateTime `json:"created,omitempty"`

	// The timestamp for the last update to the counterexample.
	Updated strfmt.DateTime `json:"updated,omitempty"`
}

// CounterexampleCollection : CounterexampleCollection struct
type CounterexampleCollection struct {

	// An array of objects describing the examples marked as irrelevant input.
	Counterexamples []Counterexample `json:"counterexamples"`

	// The pagination data for the returned objects.
	Pagination Pagination `json:"pagination"`
}

// CreateCounterexample : CreateCounterexample struct
type CreateCounterexample struct {

	// The text of a user input marked as irrelevant input. This string must conform to the following restrictions: - It cannot contain carriage return, newline, or tab characters - It cannot consist of only whitespace characters - It must be no longer than 1024 characters.
	Text string `json:"text"`
}

// CreateCounterexampleOptions : The createCounterexample options.
type CreateCounterexampleOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID string `json:"workspace_id"`

	// The text of a user input marked as irrelevant input. This string must conform to the following restrictions: - It cannot contain carriage return, newline, or tab characters - It cannot consist of only whitespace characters - It must be no longer than 1024 characters.
	Text string `json:"text"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewCreateCounterexampleOptions : Instantiate CreateCounterexampleOptions
func NewCreateCounterexampleOptions(workspaceID string, text string) *CreateCounterexampleOptions {
    return &CreateCounterexampleOptions{
        WorkspaceID: workspaceID,
        Text: text,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *CreateCounterexampleOptions) SetWorkspaceID(param string) *CreateCounterexampleOptions {
    options.WorkspaceID = param
    return options
}

// SetText : Allow user to set Text
func (options *CreateCounterexampleOptions) SetText(param string) *CreateCounterexampleOptions {
    options.Text = param
    return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateCounterexampleOptions) SetHeaders(param map[string]string) *CreateCounterexampleOptions {
    options.Headers = param
    return options
}

// CreateDialogNode : CreateDialogNode struct
type CreateDialogNode struct {

	// The dialog node ID. This string must conform to the following restrictions: - It can contain only Unicode alphanumeric, space, underscore, hyphen, and dot characters. - It must be no longer than 1024 characters.
	DialogNode string `json:"dialog_node"`

	// The description of the dialog node. This string cannot contain carriage return, newline, or tab characters, and it must be no longer than 128 characters.
	Description string `json:"description,omitempty"`

	// The condition that will trigger the dialog node. This string cannot contain carriage return, newline, or tab characters, and it must be no longer than 2048 characters.
	Conditions string `json:"conditions,omitempty"`

	// The ID of the parent dialog node.
	Parent string `json:"parent,omitempty"`

	// The ID of the previous dialog node.
	PreviousSibling string `json:"previous_sibling,omitempty"`

	// The output of the dialog node. For more information about how to specify dialog node output, see the [documentation](https://console.bluemix.net/docs/services/conversation/dialog-overview.html#complex).
	Output DialogNodeOutput `json:"output,omitempty"`

	// The context for the dialog node.
	Context interface{} `json:"context,omitempty"`

	// The metadata for the dialog node.
	Metadata interface{} `json:"metadata,omitempty"`

	// The next step to be executed in dialog processing.
	NextStep DialogNodeNextStep `json:"next_step,omitempty"`

	// An array of objects describing any actions to be invoked by the dialog node.
	Actions []DialogNodeAction `json:"actions,omitempty"`

	// The alias used to identify the dialog node. This string must conform to the following restrictions: - It can contain only Unicode alphanumeric, space, underscore, hyphen, and dot characters. - It must be no longer than 64 characters.
	Title string `json:"title,omitempty"`

	// How the dialog node is processed.
	NodeType string `json:"type,omitempty"`

	// How an `event_handler` node is processed.
	EventName string `json:"event_name,omitempty"`

	// The location in the dialog context where output is stored.
	Variable string `json:"variable,omitempty"`

	// Whether this top-level dialog node can be digressed into.
	DigressIn string `json:"digress_in,omitempty"`

	// Whether this dialog node can be returned to after a digression.
	DigressOut string `json:"digress_out,omitempty"`

	// Whether the user can digress to top-level nodes while filling out slots.
	DigressOutSlots string `json:"digress_out_slots,omitempty"`

	// A label that can be displayed externally to describe the purpose of the node to users.
	UserLabel string `json:"user_label,omitempty"`
}

// CreateDialogNodeOptions : The createDialogNode options.
type CreateDialogNodeOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID string `json:"workspace_id"`

	// The dialog node ID. This string must conform to the following restrictions: - It can contain only Unicode alphanumeric, space, underscore, hyphen, and dot characters. - It must be no longer than 1024 characters.
	DialogNode string `json:"dialog_node"`

	// The description of the dialog node. This string cannot contain carriage return, newline, or tab characters, and it must be no longer than 128 characters.
	Description string `json:"description,omitempty"`

    // Indicates whether user set optional parameter Description
    IsDescriptionSet bool

	// The condition that will trigger the dialog node. This string cannot contain carriage return, newline, or tab characters, and it must be no longer than 2048 characters.
	Conditions string `json:"conditions,omitempty"`

    // Indicates whether user set optional parameter Conditions
    IsConditionsSet bool

	// The ID of the parent dialog node.
	Parent string `json:"parent,omitempty"`

    // Indicates whether user set optional parameter Parent
    IsParentSet bool

	// The ID of the previous dialog node.
	PreviousSibling string `json:"previous_sibling,omitempty"`

    // Indicates whether user set optional parameter PreviousSibling
    IsPreviousSiblingSet bool

	// The output of the dialog node. For more information about how to specify dialog node output, see the [documentation](https://console.bluemix.net/docs/services/conversation/dialog-overview.html#complex).
	Output DialogNodeOutput `json:"output,omitempty"`

    // Indicates whether user set optional parameter Output
    IsOutputSet bool

	// The context for the dialog node.
	Context interface{} `json:"context,omitempty"`

    // Indicates whether user set optional parameter Context
    IsContextSet bool

	// The metadata for the dialog node.
	Metadata interface{} `json:"metadata,omitempty"`

    // Indicates whether user set optional parameter Metadata
    IsMetadataSet bool

	// The next step to be executed in dialog processing.
	NextStep DialogNodeNextStep `json:"next_step,omitempty"`

    // Indicates whether user set optional parameter NextStep
    IsNextStepSet bool

	// An array of objects describing any actions to be invoked by the dialog node.
	Actions []DialogNodeAction `json:"actions,omitempty"`

    // Indicates whether user set optional parameter Actions
    IsActionsSet bool

	// The alias used to identify the dialog node. This string must conform to the following restrictions: - It can contain only Unicode alphanumeric, space, underscore, hyphen, and dot characters. - It must be no longer than 64 characters.
	Title string `json:"title,omitempty"`

    // Indicates whether user set optional parameter Title
    IsTitleSet bool

	// How the dialog node is processed.
	NodeType string `json:"type,omitempty"`

    // Indicates whether user set optional parameter NodeType
    IsNodeTypeSet bool

	// How an `event_handler` node is processed.
	EventName string `json:"event_name,omitempty"`

    // Indicates whether user set optional parameter EventName
    IsEventNameSet bool

	// The location in the dialog context where output is stored.
	Variable string `json:"variable,omitempty"`

    // Indicates whether user set optional parameter Variable
    IsVariableSet bool

	// Whether this top-level dialog node can be digressed into.
	DigressIn string `json:"digress_in,omitempty"`

    // Indicates whether user set optional parameter DigressIn
    IsDigressInSet bool

	// Whether this dialog node can be returned to after a digression.
	DigressOut string `json:"digress_out,omitempty"`

    // Indicates whether user set optional parameter DigressOut
    IsDigressOutSet bool

	// Whether the user can digress to top-level nodes while filling out slots.
	DigressOutSlots string `json:"digress_out_slots,omitempty"`

    // Indicates whether user set optional parameter DigressOutSlots
    IsDigressOutSlotsSet bool

	// A label that can be displayed externally to describe the purpose of the node to users.
	UserLabel string `json:"user_label,omitempty"`

    // Indicates whether user set optional parameter UserLabel
    IsUserLabelSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewCreateDialogNodeOptions : Instantiate CreateDialogNodeOptions
func NewCreateDialogNodeOptions(workspaceID string, dialogNode string) *CreateDialogNodeOptions {
    return &CreateDialogNodeOptions{
        WorkspaceID: workspaceID,
        DialogNode: dialogNode,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *CreateDialogNodeOptions) SetWorkspaceID(param string) *CreateDialogNodeOptions {
    options.WorkspaceID = param
    return options
}

// SetDialogNode : Allow user to set DialogNode
func (options *CreateDialogNodeOptions) SetDialogNode(param string) *CreateDialogNodeOptions {
    options.DialogNode = param
    return options
}

// SetDescription : Allow user to set Description
func (options *CreateDialogNodeOptions) SetDescription(param string) *CreateDialogNodeOptions {
    options.Description = param
    options.IsDescriptionSet = true
    return options
}

// SetConditions : Allow user to set Conditions
func (options *CreateDialogNodeOptions) SetConditions(param string) *CreateDialogNodeOptions {
    options.Conditions = param
    options.IsConditionsSet = true
    return options
}

// SetParent : Allow user to set Parent
func (options *CreateDialogNodeOptions) SetParent(param string) *CreateDialogNodeOptions {
    options.Parent = param
    options.IsParentSet = true
    return options
}

// SetPreviousSibling : Allow user to set PreviousSibling
func (options *CreateDialogNodeOptions) SetPreviousSibling(param string) *CreateDialogNodeOptions {
    options.PreviousSibling = param
    options.IsPreviousSiblingSet = true
    return options
}

// SetOutput : Allow user to set Output
func (options *CreateDialogNodeOptions) SetOutput(param DialogNodeOutput) *CreateDialogNodeOptions {
    options.Output = param
    options.IsOutputSet = true
    return options
}

// SetContext : Allow user to set Context
func (options *CreateDialogNodeOptions) SetContext(param interface{}) *CreateDialogNodeOptions {
    options.Context = param
    options.IsContextSet = true
    return options
}

// SetMetadata : Allow user to set Metadata
func (options *CreateDialogNodeOptions) SetMetadata(param interface{}) *CreateDialogNodeOptions {
    options.Metadata = param
    options.IsMetadataSet = true
    return options
}

// SetNextStep : Allow user to set NextStep
func (options *CreateDialogNodeOptions) SetNextStep(param DialogNodeNextStep) *CreateDialogNodeOptions {
    options.NextStep = param
    options.IsNextStepSet = true
    return options
}

// SetActions : Allow user to set Actions
func (options *CreateDialogNodeOptions) SetActions(param []DialogNodeAction) *CreateDialogNodeOptions {
    options.Actions = param
    options.IsActionsSet = true
    return options
}

// SetTitle : Allow user to set Title
func (options *CreateDialogNodeOptions) SetTitle(param string) *CreateDialogNodeOptions {
    options.Title = param
    options.IsTitleSet = true
    return options
}

// SetNodeType : Allow user to set NodeType
func (options *CreateDialogNodeOptions) SetNodeType(param string) *CreateDialogNodeOptions {
    options.NodeType = param
    options.IsNodeTypeSet = true
    return options
}

// SetEventName : Allow user to set EventName
func (options *CreateDialogNodeOptions) SetEventName(param string) *CreateDialogNodeOptions {
    options.EventName = param
    options.IsEventNameSet = true
    return options
}

// SetVariable : Allow user to set Variable
func (options *CreateDialogNodeOptions) SetVariable(param string) *CreateDialogNodeOptions {
    options.Variable = param
    options.IsVariableSet = true
    return options
}

// SetDigressIn : Allow user to set DigressIn
func (options *CreateDialogNodeOptions) SetDigressIn(param string) *CreateDialogNodeOptions {
    options.DigressIn = param
    options.IsDigressInSet = true
    return options
}

// SetDigressOut : Allow user to set DigressOut
func (options *CreateDialogNodeOptions) SetDigressOut(param string) *CreateDialogNodeOptions {
    options.DigressOut = param
    options.IsDigressOutSet = true
    return options
}

// SetDigressOutSlots : Allow user to set DigressOutSlots
func (options *CreateDialogNodeOptions) SetDigressOutSlots(param string) *CreateDialogNodeOptions {
    options.DigressOutSlots = param
    options.IsDigressOutSlotsSet = true
    return options
}

// SetUserLabel : Allow user to set UserLabel
func (options *CreateDialogNodeOptions) SetUserLabel(param string) *CreateDialogNodeOptions {
    options.UserLabel = param
    options.IsUserLabelSet = true
    return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateDialogNodeOptions) SetHeaders(param map[string]string) *CreateDialogNodeOptions {
    options.Headers = param
    return options
}

// CreateEntity : CreateEntity struct
type CreateEntity struct {

	// The name of the entity. This string must conform to the following restrictions: - It can contain only Unicode alphanumeric, underscore, and hyphen characters. - It cannot begin with the reserved prefix `sys-`. - It must be no longer than 64 characters.
	Entity string `json:"entity"`

	// The description of the entity. This string cannot contain carriage return, newline, or tab characters, and it must be no longer than 128 characters.
	Description string `json:"description,omitempty"`

	// Any metadata related to the value.
	Metadata interface{} `json:"metadata,omitempty"`

	// An array of objects describing the entity values.
	Values []CreateValue `json:"values,omitempty"`

	// Whether to use fuzzy matching for the entity.
	FuzzyMatch bool `json:"fuzzy_match,omitempty"`
}

// CreateEntityOptions : The createEntity options.
type CreateEntityOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID string `json:"workspace_id"`

	// The name of the entity. This string must conform to the following restrictions: - It can contain only Unicode alphanumeric, underscore, and hyphen characters. - It cannot begin with the reserved prefix `sys-`. - It must be no longer than 64 characters.
	Entity string `json:"entity"`

	// The description of the entity. This string cannot contain carriage return, newline, or tab characters, and it must be no longer than 128 characters.
	Description string `json:"description,omitempty"`

    // Indicates whether user set optional parameter Description
    IsDescriptionSet bool

	// Any metadata related to the value.
	Metadata interface{} `json:"metadata,omitempty"`

    // Indicates whether user set optional parameter Metadata
    IsMetadataSet bool

	// An array of objects describing the entity values.
	Values []CreateValue `json:"values,omitempty"`

    // Indicates whether user set optional parameter Values
    IsValuesSet bool

	// Whether to use fuzzy matching for the entity.
	FuzzyMatch bool `json:"fuzzy_match,omitempty"`

    // Indicates whether user set optional parameter FuzzyMatch
    IsFuzzyMatchSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewCreateEntityOptions : Instantiate CreateEntityOptions
func NewCreateEntityOptions(workspaceID string, entity string) *CreateEntityOptions {
    return &CreateEntityOptions{
        WorkspaceID: workspaceID,
        Entity: entity,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *CreateEntityOptions) SetWorkspaceID(param string) *CreateEntityOptions {
    options.WorkspaceID = param
    return options
}

// SetEntity : Allow user to set Entity
func (options *CreateEntityOptions) SetEntity(param string) *CreateEntityOptions {
    options.Entity = param
    return options
}

// SetDescription : Allow user to set Description
func (options *CreateEntityOptions) SetDescription(param string) *CreateEntityOptions {
    options.Description = param
    options.IsDescriptionSet = true
    return options
}

// SetMetadata : Allow user to set Metadata
func (options *CreateEntityOptions) SetMetadata(param interface{}) *CreateEntityOptions {
    options.Metadata = param
    options.IsMetadataSet = true
    return options
}

// SetValues : Allow user to set Values
func (options *CreateEntityOptions) SetValues(param []CreateValue) *CreateEntityOptions {
    options.Values = param
    options.IsValuesSet = true
    return options
}

// SetFuzzyMatch : Allow user to set FuzzyMatch
func (options *CreateEntityOptions) SetFuzzyMatch(param bool) *CreateEntityOptions {
    options.FuzzyMatch = param
    options.IsFuzzyMatchSet = true
    return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateEntityOptions) SetHeaders(param map[string]string) *CreateEntityOptions {
    options.Headers = param
    return options
}

// CreateExample : CreateExample struct
type CreateExample struct {

	// The text of a user input example. This string must conform to the following restrictions: - It cannot contain carriage return, newline, or tab characters. - It cannot consist of only whitespace characters. - It must be no longer than 1024 characters.
	Text string `json:"text"`

	// An array of contextual entity mentions.
	Mentions []Mentions `json:"mentions,omitempty"`
}

// CreateExampleOptions : The createExample options.
type CreateExampleOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID string `json:"workspace_id"`

	// The intent name.
	Intent string `json:"intent"`

	// The text of a user input example. This string must conform to the following restrictions: - It cannot contain carriage return, newline, or tab characters. - It cannot consist of only whitespace characters. - It must be no longer than 1024 characters.
	Text string `json:"text"`

	// An array of contextual entity mentions.
	Mentions []Mentions `json:"mentions,omitempty"`

    // Indicates whether user set optional parameter Mentions
    IsMentionsSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewCreateExampleOptions : Instantiate CreateExampleOptions
func NewCreateExampleOptions(workspaceID string, intent string, text string) *CreateExampleOptions {
    return &CreateExampleOptions{
        WorkspaceID: workspaceID,
        Intent: intent,
        Text: text,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *CreateExampleOptions) SetWorkspaceID(param string) *CreateExampleOptions {
    options.WorkspaceID = param
    return options
}

// SetIntent : Allow user to set Intent
func (options *CreateExampleOptions) SetIntent(param string) *CreateExampleOptions {
    options.Intent = param
    return options
}

// SetText : Allow user to set Text
func (options *CreateExampleOptions) SetText(param string) *CreateExampleOptions {
    options.Text = param
    return options
}

// SetMentions : Allow user to set Mentions
func (options *CreateExampleOptions) SetMentions(param []Mentions) *CreateExampleOptions {
    options.Mentions = param
    options.IsMentionsSet = true
    return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateExampleOptions) SetHeaders(param map[string]string) *CreateExampleOptions {
    options.Headers = param
    return options
}

// CreateIntent : CreateIntent struct
type CreateIntent struct {

	// The name of the intent. This string must conform to the following restrictions: - It can contain only Unicode alphanumeric, underscore, hyphen, and dot characters. - It cannot begin with the reserved prefix `sys-`. - It must be no longer than 128 characters.
	Intent string `json:"intent"`

	// The description of the intent. This string cannot contain carriage return, newline, or tab characters, and it must be no longer than 128 characters.
	Description string `json:"description,omitempty"`

	// An array of user input examples for the intent.
	Examples []CreateExample `json:"examples,omitempty"`
}

// CreateIntentOptions : The createIntent options.
type CreateIntentOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID string `json:"workspace_id"`

	// The name of the intent. This string must conform to the following restrictions: - It can contain only Unicode alphanumeric, underscore, hyphen, and dot characters. - It cannot begin with the reserved prefix `sys-`. - It must be no longer than 128 characters.
	Intent string `json:"intent"`

	// The description of the intent. This string cannot contain carriage return, newline, or tab characters, and it must be no longer than 128 characters.
	Description string `json:"description,omitempty"`

    // Indicates whether user set optional parameter Description
    IsDescriptionSet bool

	// An array of user input examples for the intent.
	Examples []CreateExample `json:"examples,omitempty"`

    // Indicates whether user set optional parameter Examples
    IsExamplesSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewCreateIntentOptions : Instantiate CreateIntentOptions
func NewCreateIntentOptions(workspaceID string, intent string) *CreateIntentOptions {
    return &CreateIntentOptions{
        WorkspaceID: workspaceID,
        Intent: intent,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *CreateIntentOptions) SetWorkspaceID(param string) *CreateIntentOptions {
    options.WorkspaceID = param
    return options
}

// SetIntent : Allow user to set Intent
func (options *CreateIntentOptions) SetIntent(param string) *CreateIntentOptions {
    options.Intent = param
    return options
}

// SetDescription : Allow user to set Description
func (options *CreateIntentOptions) SetDescription(param string) *CreateIntentOptions {
    options.Description = param
    options.IsDescriptionSet = true
    return options
}

// SetExamples : Allow user to set Examples
func (options *CreateIntentOptions) SetExamples(param []CreateExample) *CreateIntentOptions {
    options.Examples = param
    options.IsExamplesSet = true
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
	WorkspaceID string `json:"workspace_id"`

	// The name of the entity.
	Entity string `json:"entity"`

	// The text of the entity value.
	Value string `json:"value"`

	// The text of the synonym. This string must conform to the following restrictions: - It cannot contain carriage return, newline, or tab characters. - It cannot consist of only whitespace characters. - It must be no longer than 64 characters.
	Synonym string `json:"synonym"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewCreateSynonymOptions : Instantiate CreateSynonymOptions
func NewCreateSynonymOptions(workspaceID string, entity string, value string, synonym string) *CreateSynonymOptions {
    return &CreateSynonymOptions{
        WorkspaceID: workspaceID,
        Entity: entity,
        Value: value,
        Synonym: synonym,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *CreateSynonymOptions) SetWorkspaceID(param string) *CreateSynonymOptions {
    options.WorkspaceID = param
    return options
}

// SetEntity : Allow user to set Entity
func (options *CreateSynonymOptions) SetEntity(param string) *CreateSynonymOptions {
    options.Entity = param
    return options
}

// SetValue : Allow user to set Value
func (options *CreateSynonymOptions) SetValue(param string) *CreateSynonymOptions {
    options.Value = param
    return options
}

// SetSynonym : Allow user to set Synonym
func (options *CreateSynonymOptions) SetSynonym(param string) *CreateSynonymOptions {
    options.Synonym = param
    return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateSynonymOptions) SetHeaders(param map[string]string) *CreateSynonymOptions {
    options.Headers = param
    return options
}

// CreateValue : CreateValue struct
type CreateValue struct {

	// The text of the entity value. This string must conform to the following restrictions: - It cannot contain carriage return, newline, or tab characters. - It cannot consist of only whitespace characters. - It must be no longer than 64 characters.
	Value string `json:"value"`

	// Any metadata related to the entity value.
	Metadata interface{} `json:"metadata,omitempty"`

	// An array containing any synonyms for the entity value. You can provide either synonyms or patterns (as indicated by **type**), but not both. A synonym must conform to the following restrictions: - It cannot contain carriage return, newline, or tab characters. - It cannot consist of only whitespace characters. - It must be no longer than 64 characters.
	Synonyms []string `json:"synonyms,omitempty"`

	// An array of patterns for the entity value. You can provide either synonyms or patterns (as indicated by **type**), but not both. A pattern is a regular expression no longer than 512 characters. For more information about how to specify a pattern, see the [documentation](https://console.bluemix.net/docs/services/conversation/entities.html#creating-entities).
	Patterns []string `json:"patterns,omitempty"`

	// Specifies the type of value.
	ValueType string `json:"type,omitempty"`
}

// CreateValueOptions : The createValue options.
type CreateValueOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID string `json:"workspace_id"`

	// The name of the entity.
	Entity string `json:"entity"`

	// The text of the entity value. This string must conform to the following restrictions: - It cannot contain carriage return, newline, or tab characters. - It cannot consist of only whitespace characters. - It must be no longer than 64 characters.
	Value string `json:"value"`

	// Any metadata related to the entity value.
	Metadata interface{} `json:"metadata,omitempty"`

    // Indicates whether user set optional parameter Metadata
    IsMetadataSet bool

	// An array containing any synonyms for the entity value. You can provide either synonyms or patterns (as indicated by **type**), but not both. A synonym must conform to the following restrictions: - It cannot contain carriage return, newline, or tab characters. - It cannot consist of only whitespace characters. - It must be no longer than 64 characters.
	Synonyms []string `json:"synonyms,omitempty"`

    // Indicates whether user set optional parameter Synonyms
    IsSynonymsSet bool

	// An array of patterns for the entity value. You can provide either synonyms or patterns (as indicated by **type**), but not both. A pattern is a regular expression no longer than 512 characters. For more information about how to specify a pattern, see the [documentation](https://console.bluemix.net/docs/services/conversation/entities.html#creating-entities).
	Patterns []string `json:"patterns,omitempty"`

    // Indicates whether user set optional parameter Patterns
    IsPatternsSet bool

	// Specifies the type of value.
	ValueType string `json:"type,omitempty"`

    // Indicates whether user set optional parameter ValueType
    IsValueTypeSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewCreateValueOptions : Instantiate CreateValueOptions
func NewCreateValueOptions(workspaceID string, entity string, value string) *CreateValueOptions {
    return &CreateValueOptions{
        WorkspaceID: workspaceID,
        Entity: entity,
        Value: value,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *CreateValueOptions) SetWorkspaceID(param string) *CreateValueOptions {
    options.WorkspaceID = param
    return options
}

// SetEntity : Allow user to set Entity
func (options *CreateValueOptions) SetEntity(param string) *CreateValueOptions {
    options.Entity = param
    return options
}

// SetValue : Allow user to set Value
func (options *CreateValueOptions) SetValue(param string) *CreateValueOptions {
    options.Value = param
    return options
}

// SetMetadata : Allow user to set Metadata
func (options *CreateValueOptions) SetMetadata(param interface{}) *CreateValueOptions {
    options.Metadata = param
    options.IsMetadataSet = true
    return options
}

// SetSynonyms : Allow user to set Synonyms
func (options *CreateValueOptions) SetSynonyms(param []string) *CreateValueOptions {
    options.Synonyms = param
    options.IsSynonymsSet = true
    return options
}

// SetPatterns : Allow user to set Patterns
func (options *CreateValueOptions) SetPatterns(param []string) *CreateValueOptions {
    options.Patterns = param
    options.IsPatternsSet = true
    return options
}

// SetValueType : Allow user to set ValueType
func (options *CreateValueOptions) SetValueType(param string) *CreateValueOptions {
    options.ValueType = param
    options.IsValueTypeSet = true
    return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateValueOptions) SetHeaders(param map[string]string) *CreateValueOptions {
    options.Headers = param
    return options
}

// CreateWorkspaceOptions : The createWorkspace options.
type CreateWorkspaceOptions struct {

	// The name of the workspace. This string cannot contain carriage return, newline, or tab characters, and it must be no longer than 64 characters.
	Name string `json:"name,omitempty"`

    // Indicates whether user set optional parameter Name
    IsNameSet bool

	// The description of the workspace. This string cannot contain carriage return, newline, or tab characters, and it must be no longer than 128 characters.
	Description string `json:"description,omitempty"`

    // Indicates whether user set optional parameter Description
    IsDescriptionSet bool

	// The language of the workspace.
	Language string `json:"language,omitempty"`

    // Indicates whether user set optional parameter Language
    IsLanguageSet bool

	// An array of objects defining the intents for the workspace.
	Intents []CreateIntent `json:"intents,omitempty"`

    // Indicates whether user set optional parameter Intents
    IsIntentsSet bool

	// An array of objects defining the entities for the workspace.
	Entities []CreateEntity `json:"entities,omitempty"`

    // Indicates whether user set optional parameter Entities
    IsEntitiesSet bool

	// An array of objects defining the nodes in the workspace dialog.
	DialogNodes []CreateDialogNode `json:"dialog_nodes,omitempty"`

    // Indicates whether user set optional parameter DialogNodes
    IsDialogNodesSet bool

	// An array of objects defining input examples that have been marked as irrelevant input.
	Counterexamples []CreateCounterexample `json:"counterexamples,omitempty"`

    // Indicates whether user set optional parameter Counterexamples
    IsCounterexamplesSet bool

	// Any metadata related to the workspace.
	Metadata interface{} `json:"metadata,omitempty"`

    // Indicates whether user set optional parameter Metadata
    IsMetadataSet bool

	// Whether training data from the workspace can be used by IBM for general service improvements. `true` indicates that workspace training data is not to be used.
	LearningOptOut bool `json:"learning_opt_out,omitempty"`

    // Indicates whether user set optional parameter LearningOptOut
    IsLearningOptOutSet bool

	// Global settings for the workspace.
	SystemSettings WorkspaceSystemSettings `json:"system_settings,omitempty"`

    // Indicates whether user set optional parameter SystemSettings
    IsSystemSettingsSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewCreateWorkspaceOptions : Instantiate CreateWorkspaceOptions
func NewCreateWorkspaceOptions() *CreateWorkspaceOptions {
    return &CreateWorkspaceOptions{}
}

// SetName : Allow user to set Name
func (options *CreateWorkspaceOptions) SetName(param string) *CreateWorkspaceOptions {
    options.Name = param
    options.IsNameSet = true
    return options
}

// SetDescription : Allow user to set Description
func (options *CreateWorkspaceOptions) SetDescription(param string) *CreateWorkspaceOptions {
    options.Description = param
    options.IsDescriptionSet = true
    return options
}

// SetLanguage : Allow user to set Language
func (options *CreateWorkspaceOptions) SetLanguage(param string) *CreateWorkspaceOptions {
    options.Language = param
    options.IsLanguageSet = true
    return options
}

// SetIntents : Allow user to set Intents
func (options *CreateWorkspaceOptions) SetIntents(param []CreateIntent) *CreateWorkspaceOptions {
    options.Intents = param
    options.IsIntentsSet = true
    return options
}

// SetEntities : Allow user to set Entities
func (options *CreateWorkspaceOptions) SetEntities(param []CreateEntity) *CreateWorkspaceOptions {
    options.Entities = param
    options.IsEntitiesSet = true
    return options
}

// SetDialogNodes : Allow user to set DialogNodes
func (options *CreateWorkspaceOptions) SetDialogNodes(param []CreateDialogNode) *CreateWorkspaceOptions {
    options.DialogNodes = param
    options.IsDialogNodesSet = true
    return options
}

// SetCounterexamples : Allow user to set Counterexamples
func (options *CreateWorkspaceOptions) SetCounterexamples(param []CreateCounterexample) *CreateWorkspaceOptions {
    options.Counterexamples = param
    options.IsCounterexamplesSet = true
    return options
}

// SetMetadata : Allow user to set Metadata
func (options *CreateWorkspaceOptions) SetMetadata(param interface{}) *CreateWorkspaceOptions {
    options.Metadata = param
    options.IsMetadataSet = true
    return options
}

// SetLearningOptOut : Allow user to set LearningOptOut
func (options *CreateWorkspaceOptions) SetLearningOptOut(param bool) *CreateWorkspaceOptions {
    options.LearningOptOut = param
    options.IsLearningOptOutSet = true
    return options
}

// SetSystemSettings : Allow user to set SystemSettings
func (options *CreateWorkspaceOptions) SetSystemSettings(param WorkspaceSystemSettings) *CreateWorkspaceOptions {
    options.SystemSettings = param
    options.IsSystemSettingsSet = true
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
	WorkspaceID string `json:"workspace_id"`

	// The text of a user input counterexample (for example, `What are you wearing?`).
	Text string `json:"text"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewDeleteCounterexampleOptions : Instantiate DeleteCounterexampleOptions
func NewDeleteCounterexampleOptions(workspaceID string, text string) *DeleteCounterexampleOptions {
    return &DeleteCounterexampleOptions{
        WorkspaceID: workspaceID,
        Text: text,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *DeleteCounterexampleOptions) SetWorkspaceID(param string) *DeleteCounterexampleOptions {
    options.WorkspaceID = param
    return options
}

// SetText : Allow user to set Text
func (options *DeleteCounterexampleOptions) SetText(param string) *DeleteCounterexampleOptions {
    options.Text = param
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
	WorkspaceID string `json:"workspace_id"`

	// The dialog node ID (for example, `get_order`).
	DialogNode string `json:"dialog_node"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewDeleteDialogNodeOptions : Instantiate DeleteDialogNodeOptions
func NewDeleteDialogNodeOptions(workspaceID string, dialogNode string) *DeleteDialogNodeOptions {
    return &DeleteDialogNodeOptions{
        WorkspaceID: workspaceID,
        DialogNode: dialogNode,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *DeleteDialogNodeOptions) SetWorkspaceID(param string) *DeleteDialogNodeOptions {
    options.WorkspaceID = param
    return options
}

// SetDialogNode : Allow user to set DialogNode
func (options *DeleteDialogNodeOptions) SetDialogNode(param string) *DeleteDialogNodeOptions {
    options.DialogNode = param
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
	WorkspaceID string `json:"workspace_id"`

	// The name of the entity.
	Entity string `json:"entity"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewDeleteEntityOptions : Instantiate DeleteEntityOptions
func NewDeleteEntityOptions(workspaceID string, entity string) *DeleteEntityOptions {
    return &DeleteEntityOptions{
        WorkspaceID: workspaceID,
        Entity: entity,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *DeleteEntityOptions) SetWorkspaceID(param string) *DeleteEntityOptions {
    options.WorkspaceID = param
    return options
}

// SetEntity : Allow user to set Entity
func (options *DeleteEntityOptions) SetEntity(param string) *DeleteEntityOptions {
    options.Entity = param
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
	WorkspaceID string `json:"workspace_id"`

	// The intent name.
	Intent string `json:"intent"`

	// The text of the user input example.
	Text string `json:"text"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewDeleteExampleOptions : Instantiate DeleteExampleOptions
func NewDeleteExampleOptions(workspaceID string, intent string, text string) *DeleteExampleOptions {
    return &DeleteExampleOptions{
        WorkspaceID: workspaceID,
        Intent: intent,
        Text: text,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *DeleteExampleOptions) SetWorkspaceID(param string) *DeleteExampleOptions {
    options.WorkspaceID = param
    return options
}

// SetIntent : Allow user to set Intent
func (options *DeleteExampleOptions) SetIntent(param string) *DeleteExampleOptions {
    options.Intent = param
    return options
}

// SetText : Allow user to set Text
func (options *DeleteExampleOptions) SetText(param string) *DeleteExampleOptions {
    options.Text = param
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
	WorkspaceID string `json:"workspace_id"`

	// The intent name.
	Intent string `json:"intent"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewDeleteIntentOptions : Instantiate DeleteIntentOptions
func NewDeleteIntentOptions(workspaceID string, intent string) *DeleteIntentOptions {
    return &DeleteIntentOptions{
        WorkspaceID: workspaceID,
        Intent: intent,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *DeleteIntentOptions) SetWorkspaceID(param string) *DeleteIntentOptions {
    options.WorkspaceID = param
    return options
}

// SetIntent : Allow user to set Intent
func (options *DeleteIntentOptions) SetIntent(param string) *DeleteIntentOptions {
    options.Intent = param
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
	WorkspaceID string `json:"workspace_id"`

	// The name of the entity.
	Entity string `json:"entity"`

	// The text of the entity value.
	Value string `json:"value"`

	// The text of the synonym.
	Synonym string `json:"synonym"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewDeleteSynonymOptions : Instantiate DeleteSynonymOptions
func NewDeleteSynonymOptions(workspaceID string, entity string, value string, synonym string) *DeleteSynonymOptions {
    return &DeleteSynonymOptions{
        WorkspaceID: workspaceID,
        Entity: entity,
        Value: value,
        Synonym: synonym,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *DeleteSynonymOptions) SetWorkspaceID(param string) *DeleteSynonymOptions {
    options.WorkspaceID = param
    return options
}

// SetEntity : Allow user to set Entity
func (options *DeleteSynonymOptions) SetEntity(param string) *DeleteSynonymOptions {
    options.Entity = param
    return options
}

// SetValue : Allow user to set Value
func (options *DeleteSynonymOptions) SetValue(param string) *DeleteSynonymOptions {
    options.Value = param
    return options
}

// SetSynonym : Allow user to set Synonym
func (options *DeleteSynonymOptions) SetSynonym(param string) *DeleteSynonymOptions {
    options.Synonym = param
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
	CustomerID string `json:"customer_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewDeleteUserDataOptions : Instantiate DeleteUserDataOptions
func NewDeleteUserDataOptions(customerID string) *DeleteUserDataOptions {
    return &DeleteUserDataOptions{
        CustomerID: customerID,
    }
}

// SetCustomerID : Allow user to set CustomerID
func (options *DeleteUserDataOptions) SetCustomerID(param string) *DeleteUserDataOptions {
    options.CustomerID = param
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
	WorkspaceID string `json:"workspace_id"`

	// The name of the entity.
	Entity string `json:"entity"`

	// The text of the entity value.
	Value string `json:"value"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewDeleteValueOptions : Instantiate DeleteValueOptions
func NewDeleteValueOptions(workspaceID string, entity string, value string) *DeleteValueOptions {
    return &DeleteValueOptions{
        WorkspaceID: workspaceID,
        Entity: entity,
        Value: value,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *DeleteValueOptions) SetWorkspaceID(param string) *DeleteValueOptions {
    options.WorkspaceID = param
    return options
}

// SetEntity : Allow user to set Entity
func (options *DeleteValueOptions) SetEntity(param string) *DeleteValueOptions {
    options.Entity = param
    return options
}

// SetValue : Allow user to set Value
func (options *DeleteValueOptions) SetValue(param string) *DeleteValueOptions {
    options.Value = param
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
	WorkspaceID string `json:"workspace_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewDeleteWorkspaceOptions : Instantiate DeleteWorkspaceOptions
func NewDeleteWorkspaceOptions(workspaceID string) *DeleteWorkspaceOptions {
    return &DeleteWorkspaceOptions{
        WorkspaceID: workspaceID,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *DeleteWorkspaceOptions) SetWorkspaceID(param string) *DeleteWorkspaceOptions {
    options.WorkspaceID = param
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
	DialogNodeID string `json:"dialog_node"`

	// The description of the dialog node.
	Description string `json:"description,omitempty"`

	// The condition that triggers the dialog node.
	Conditions string `json:"conditions,omitempty"`

	// The ID of the parent dialog node. This property is not returned if the dialog node has no parent.
	Parent string `json:"parent,omitempty"`

	// The ID of the previous sibling dialog node. This property is not returned if the dialog node has no previous sibling.
	PreviousSibling string `json:"previous_sibling,omitempty"`

	// The output of the dialog node. For more information about how to specify dialog node output, see the [documentation](https://console.bluemix.net/docs/services/conversation/dialog-overview.html#complex).
	Output DialogNodeOutput `json:"output,omitempty"`

	// The context (if defined) for the dialog node.
	Context interface{} `json:"context,omitempty"`

	// Any metadata for the dialog node.
	Metadata interface{} `json:"metadata,omitempty"`

	// The next step to execute following this dialog node.
	NextStep DialogNodeNextStep `json:"next_step,omitempty"`

	// The timestamp for creation of the dialog node.
	Created strfmt.DateTime `json:"created,omitempty"`

	// The timestamp for the most recent update to the dialog node.
	Updated strfmt.DateTime `json:"updated,omitempty"`

	// The actions for the dialog node.
	Actions []DialogNodeAction `json:"actions,omitempty"`

	// The alias used to identify the dialog node.
	Title string `json:"title,omitempty"`

	// How the dialog node is processed.
	NodeType string `json:"type,omitempty"`

	// How an `event_handler` node is processed.
	EventName string `json:"event_name,omitempty"`

	// The location in the dialog context where output is stored.
	Variable string `json:"variable,omitempty"`

	// Whether this top-level dialog node can be digressed into.
	DigressIn string `json:"digress_in,omitempty"`

	// Whether this dialog node can be returned to after a digression.
	DigressOut string `json:"digress_out,omitempty"`

	// Whether the user can digress to top-level nodes while filling out slots.
	DigressOutSlots string `json:"digress_out_slots,omitempty"`

	// A label that can be displayed externally to describe the purpose of the node to users.
	UserLabel string `json:"user_label,omitempty"`
}

// DialogNodeAction : DialogNodeAction struct
type DialogNodeAction struct {

	// The name of the action.
	Name string `json:"name"`

	// The type of action to invoke.
	ActionType string `json:"type,omitempty"`

	// A map of key/value pairs to be provided to the action.
	Parameters interface{} `json:"parameters,omitempty"`

	// The location in the dialog context where the result of the action is stored.
	ResultVariable string `json:"result_variable"`

	// The name of the context variable that the client application will use to pass in credentials for the action.
	Credentials string `json:"credentials,omitempty"`
}

// DialogNodeCollection : An array of dialog nodes.
type DialogNodeCollection struct {

	// An array of objects describing the dialog nodes defined for the workspace.
	DialogNodes []DialogNode `json:"dialog_nodes"`

	// The pagination data for the returned objects.
	Pagination Pagination `json:"pagination"`
}

// DialogNodeNextStep : The next step to execute following this dialog node.
type DialogNodeNextStep struct {

	// What happens after the dialog node completes. The valid values depend on the node type: - The following values are valid for any node: - `get_user_input` - `skip_user_input` - `jump_to` - If the node is of type `event_handler` and its parent node is of type `slot` or `frame`, additional values are also valid: - if **event_name**=`filled` and the type of the parent node is `slot`: - `reprompt` - `skip_all_slots` - if **event_name**=`nomatch` and the type of the parent node is `slot`: - `reprompt` - `skip_slot` - `skip_all_slots` - if **event_name**=`generic` and the type of the parent node is `frame`: - `reprompt` - `skip_slot` - `skip_all_slots` If you specify `jump_to`, then you must also specify a value for the `dialog_node` property.
	Behavior string `json:"behavior"`

	// The ID of the dialog node to process next. This parameter is required if **behavior**=`jump_to`.
	DialogNode string `json:"dialog_node,omitempty"`

	// Which part of the dialog node to process next.
	Selector string `json:"selector,omitempty"`
}

// DialogNodeOutput : The output of the dialog node. For more information about how to specify dialog node output, see the [documentation](https://console.bluemix.net/docs/services/conversation/dialog-overview.html#complex).
type DialogNodeOutput struct {

	// An array of objects describing the output defined for the dialog node.
	Generic []DialogNodeOutputGeneric `json:"generic,omitempty"`

	// Options that modify how specified output is handled.
	Modifiers DialogNodeOutputModifiers `json:"modifiers,omitempty"`

	// An object defining text responses in dialog nodes that do not use the `output.generic` object to define responses. New dialog nodes should use `output.generic`. For more information about how to specify dialog node output, see the [documentation](https://console.bluemix.net/docs/services/conversation/dialog-overview.html#complex).
	Text interface{} `json:"text,omitempty"`
}

// DialogNodeOutputGeneric : DialogNodeOutputGeneric struct
type DialogNodeOutputGeneric struct {

	// The type of response returned by the dialog node. The specified response type must be supported by the client application or channel.
	ResponseType string `json:"response_type"`

	// A list of one or more objects defining text responses. Required when **response_type**=`text`.
	Values []DialogNodeOutputTextValuesElement `json:"values,omitempty"`

	// How a response is selected from the list, if more than one response is specified. Valid only when **response_type**=`text`.
	SelectionPolicy string `json:"selection_policy,omitempty"`

	// The delimiter to use as a separator between responses when `selection_policy`=`multiline`.
	Delimiter string `json:"delimiter,omitempty"`

	// How long to pause, in milliseconds. The valid values are from 0 to 10000. Valid only when **response_type**=`pause`.
	Time int64 `json:"time,omitempty"`

	// Whether to send a "user is typing" event during the pause. Ignored if the channel does not support this event. Valid only when **response_type**=`pause`.
	Typing bool `json:"typing,omitempty"`

	// The URL of the image. Required when **response_type**=`image`.
	Source string `json:"source,omitempty"`

	// An optional title to show before the response. Valid only when **response_type**=`image` or `option`.
	Title string `json:"title,omitempty"`

	// An optional description to show with the response. Valid only when **response_type**=`image` or `option`.
	Description string `json:"description,omitempty"`

	// The preferred type of control to display, if supported by the channel. Valid only when **response_type**=`option`.
	Preference string `json:"preference,omitempty"`

	// An array of objects describing the options from which the user can choose. Required when **response_type**=`option`.
	Options []DialogNodeOutputOptionsElement `json:"options,omitempty"`

	// An optional message to be sent to the human agent who will be taking over the conversation. Valid only when **reponse_type**=`connect_to_agent`.
	MessageToHumanAgent string `json:"message_to_human_agent,omitempty"`
}

// DialogNodeOutputModifiers : Options that modify how specified output is handled.
type DialogNodeOutputModifiers struct {

	// Whether values in the output will overwrite output values in an array specified by previously executed dialog nodes. If this option is set to **false**, new values will be appended to previously specified values.
	Overwrite bool `json:"overwrite,omitempty"`
}

// DialogNodeOutputOptionsElement : DialogNodeOutputOptionsElement struct
type DialogNodeOutputOptionsElement struct {

	// The user-facing label for the option.
	Label string `json:"label"`

	// An object defining the message input to be sent to the Watson Assistant service if the user selects the corresponding option.
	Value DialogNodeOutputOptionsElementValue `json:"value"`
}

// DialogNodeOutputOptionsElementValue : An object defining the message input to be sent to the Watson Assistant service if the user selects the corresponding option.
type DialogNodeOutputOptionsElementValue struct {

	// The user input.
	Input InputData `json:"input,omitempty"`
}

// DialogNodeOutputTextValuesElement : DialogNodeOutputTextValuesElement struct
type DialogNodeOutputTextValuesElement struct {

	// The text of a response. This can include newline characters (` `), Markdown tagging, or other special characters, if supported by the channel.
	Text string `json:"text,omitempty"`
}

// DialogNodeVisitedDetails : DialogNodeVisitedDetails struct
type DialogNodeVisitedDetails struct {

	// A dialog node that was triggered during processing of the input message.
	DialogNode string `json:"dialog_node,omitempty"`

	// The title of the dialog node.
	Title string `json:"title,omitempty"`

	// The conditions that trigger the dialog node.
	Conditions string `json:"conditions,omitempty"`
}

// DialogRuntimeResponseGeneric : DialogRuntimeResponseGeneric struct
type DialogRuntimeResponseGeneric struct {

	// The type of response returned by the dialog node. The specified response type must be supported by the client application or channel. **Note:** The **suggestion** response type is part of the disambiguation feature, which is only available for Premium users.
	ResponseType string `json:"response_type"`

	// The text of the response.
	Text string `json:"text,omitempty"`

	// How long to pause, in milliseconds.
	Time int64 `json:"time,omitempty"`

	// Whether to send a "user is typing" event during the pause.
	Typing bool `json:"typing,omitempty"`

	// The URL of the image.
	Source string `json:"source,omitempty"`

	// The title to show before the response.
	Title string `json:"title,omitempty"`

	// The description to show with the the response.
	Description string `json:"description,omitempty"`

	// The preferred type of control to display.
	Preference string `json:"preference,omitempty"`

	// An array of objects describing the options from which the user can choose.
	Options []DialogNodeOutputOptionsElement `json:"options,omitempty"`

	// A message to be sent to the human agent who will be taking over the conversation.
	MessageToHumanAgent string `json:"message_to_human_agent,omitempty"`

	// A label identifying the topic of the conversation, derived from the **user_label** property of the relevant node.
	Topic string `json:"topic,omitempty"`

	// An array of objects describing the possible matching dialog nodes from which the user can choose. **Note:** The **suggestions** property is part of the disambiguation feature, which is only available for Premium users.
	Suggestions []DialogSuggestion `json:"suggestions,omitempty"`
}

// DialogSuggestion : DialogSuggestion struct
type DialogSuggestion struct {

	// The user-facing label for the disambiguation option. This label is taken from the **user_label** property of the corresponding dialog node.
	Label string `json:"label"`

	// An object defining the message input, intents, and entities to be sent to the Watson Assistant service if the user selects the corresponding disambiguation option.
	Value DialogSuggestionValue `json:"value"`

	// The dialog output that will be returned from the Watson Assistant service if the user selects the corresponding option.
	Output interface{} `json:"output,omitempty"`
}

// DialogSuggestionValue : An object defining the message input, intents, and entities to be sent to the Watson Assistant service if the user selects the corresponding disambiguation option.
type DialogSuggestionValue struct {

	// The user input.
	Input InputData `json:"input,omitempty"`

	// An array of intents to be sent along with the user input.
	Intents []RuntimeIntent `json:"intents,omitempty"`

	// An array of entities to be sent along with the user input.
	Entities []RuntimeEntity `json:"entities,omitempty"`
}

// Entity : Entity struct
type Entity struct {

	// The name of the entity.
	EntityName string `json:"entity"`

	// The timestamp for creation of the entity.
	Created strfmt.DateTime `json:"created,omitempty"`

	// The timestamp for the last update to the entity.
	Updated strfmt.DateTime `json:"updated,omitempty"`

	// The description of the entity.
	Description string `json:"description,omitempty"`

	// Any metadata related to the entity.
	Metadata interface{} `json:"metadata,omitempty"`

	// Whether fuzzy matching is used for the entity.
	FuzzyMatch bool `json:"fuzzy_match,omitempty"`
}

// EntityCollection : An array of entities.
type EntityCollection struct {

	// An array of objects describing the entities defined for the workspace.
	Entities []EntityExport `json:"entities"`

	// The pagination data for the returned objects.
	Pagination Pagination `json:"pagination"`
}

// EntityExport : EntityExport struct
type EntityExport struct {

	// The name of the entity.
	EntityName string `json:"entity"`

	// The timestamp for creation of the entity.
	Created strfmt.DateTime `json:"created,omitempty"`

	// The timestamp for the last update to the entity.
	Updated strfmt.DateTime `json:"updated,omitempty"`

	// The description of the entity.
	Description string `json:"description,omitempty"`

	// Any metadata related to the entity.
	Metadata interface{} `json:"metadata,omitempty"`

	// Whether fuzzy matching is used for the entity.
	FuzzyMatch bool `json:"fuzzy_match,omitempty"`

	// An array objects describing the entity values.
	Values []ValueExport `json:"values,omitempty"`
}

// EntityMention : An object describing a contextual entity mention.
type EntityMention struct {

	// The text of the user input example.
	ExampleText string `json:"text"`

	// The name of the intent.
	IntentName string `json:"intent"`

	// An array of zero-based character offsets that indicate where the entity mentions begin and end in the input text.
	Location []int64 `json:"location"`
}

// EntityMentionCollection : EntityMentionCollection struct
type EntityMentionCollection struct {

	// An array of objects describing the entity mentions defined for an entity.
	Examples []EntityMention `json:"examples"`

	// The pagination data for the returned objects.
	Pagination Pagination `json:"pagination"`
}

// Example : Example struct
type Example struct {

	// The text of the user input example.
	ExampleText string `json:"text"`

	// The timestamp for creation of the example.
	Created strfmt.DateTime `json:"created,omitempty"`

	// The timestamp for the last update to the example.
	Updated strfmt.DateTime `json:"updated,omitempty"`

	// An array of contextual entity mentions.
	Mentions []Mentions `json:"mentions,omitempty"`
}

// ExampleCollection : ExampleCollection struct
type ExampleCollection struct {

	// An array of objects describing the examples defined for the intent.
	Examples []Example `json:"examples"`

	// The pagination data for the returned objects.
	Pagination Pagination `json:"pagination"`
}

// GetCounterexampleOptions : The getCounterexample options.
type GetCounterexampleOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID string `json:"workspace_id"`

	// The text of a user input counterexample (for example, `What are you wearing?`).
	Text string `json:"text"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit bool `json:"include_audit,omitempty"`

    // Indicates whether user set optional parameter IncludeAudit
    IsIncludeAuditSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewGetCounterexampleOptions : Instantiate GetCounterexampleOptions
func NewGetCounterexampleOptions(workspaceID string, text string) *GetCounterexampleOptions {
    return &GetCounterexampleOptions{
        WorkspaceID: workspaceID,
        Text: text,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *GetCounterexampleOptions) SetWorkspaceID(param string) *GetCounterexampleOptions {
    options.WorkspaceID = param
    return options
}

// SetText : Allow user to set Text
func (options *GetCounterexampleOptions) SetText(param string) *GetCounterexampleOptions {
    options.Text = param
    return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *GetCounterexampleOptions) SetIncludeAudit(param bool) *GetCounterexampleOptions {
    options.IncludeAudit = param
    options.IsIncludeAuditSet = true
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
	WorkspaceID string `json:"workspace_id"`

	// The dialog node ID (for example, `get_order`).
	DialogNode string `json:"dialog_node"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit bool `json:"include_audit,omitempty"`

    // Indicates whether user set optional parameter IncludeAudit
    IsIncludeAuditSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewGetDialogNodeOptions : Instantiate GetDialogNodeOptions
func NewGetDialogNodeOptions(workspaceID string, dialogNode string) *GetDialogNodeOptions {
    return &GetDialogNodeOptions{
        WorkspaceID: workspaceID,
        DialogNode: dialogNode,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *GetDialogNodeOptions) SetWorkspaceID(param string) *GetDialogNodeOptions {
    options.WorkspaceID = param
    return options
}

// SetDialogNode : Allow user to set DialogNode
func (options *GetDialogNodeOptions) SetDialogNode(param string) *GetDialogNodeOptions {
    options.DialogNode = param
    return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *GetDialogNodeOptions) SetIncludeAudit(param bool) *GetDialogNodeOptions {
    options.IncludeAudit = param
    options.IsIncludeAuditSet = true
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
	WorkspaceID string `json:"workspace_id"`

	// The name of the entity.
	Entity string `json:"entity"`

	// Whether to include all element content in the returned data. If **export**=`false`, the returned data includes only information about the element itself. If **export**=`true`, all content, including subelements, is included.
	Export bool `json:"export,omitempty"`

    // Indicates whether user set optional parameter Export
    IsExportSet bool

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit bool `json:"include_audit,omitempty"`

    // Indicates whether user set optional parameter IncludeAudit
    IsIncludeAuditSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewGetEntityOptions : Instantiate GetEntityOptions
func NewGetEntityOptions(workspaceID string, entity string) *GetEntityOptions {
    return &GetEntityOptions{
        WorkspaceID: workspaceID,
        Entity: entity,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *GetEntityOptions) SetWorkspaceID(param string) *GetEntityOptions {
    options.WorkspaceID = param
    return options
}

// SetEntity : Allow user to set Entity
func (options *GetEntityOptions) SetEntity(param string) *GetEntityOptions {
    options.Entity = param
    return options
}

// SetExport : Allow user to set Export
func (options *GetEntityOptions) SetExport(param bool) *GetEntityOptions {
    options.Export = param
    options.IsExportSet = true
    return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *GetEntityOptions) SetIncludeAudit(param bool) *GetEntityOptions {
    options.IncludeAudit = param
    options.IsIncludeAuditSet = true
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
	WorkspaceID string `json:"workspace_id"`

	// The intent name.
	Intent string `json:"intent"`

	// The text of the user input example.
	Text string `json:"text"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit bool `json:"include_audit,omitempty"`

    // Indicates whether user set optional parameter IncludeAudit
    IsIncludeAuditSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewGetExampleOptions : Instantiate GetExampleOptions
func NewGetExampleOptions(workspaceID string, intent string, text string) *GetExampleOptions {
    return &GetExampleOptions{
        WorkspaceID: workspaceID,
        Intent: intent,
        Text: text,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *GetExampleOptions) SetWorkspaceID(param string) *GetExampleOptions {
    options.WorkspaceID = param
    return options
}

// SetIntent : Allow user to set Intent
func (options *GetExampleOptions) SetIntent(param string) *GetExampleOptions {
    options.Intent = param
    return options
}

// SetText : Allow user to set Text
func (options *GetExampleOptions) SetText(param string) *GetExampleOptions {
    options.Text = param
    return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *GetExampleOptions) SetIncludeAudit(param bool) *GetExampleOptions {
    options.IncludeAudit = param
    options.IsIncludeAuditSet = true
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
	WorkspaceID string `json:"workspace_id"`

	// The intent name.
	Intent string `json:"intent"`

	// Whether to include all element content in the returned data. If **export**=`false`, the returned data includes only information about the element itself. If **export**=`true`, all content, including subelements, is included.
	Export bool `json:"export,omitempty"`

    // Indicates whether user set optional parameter Export
    IsExportSet bool

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit bool `json:"include_audit,omitempty"`

    // Indicates whether user set optional parameter IncludeAudit
    IsIncludeAuditSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewGetIntentOptions : Instantiate GetIntentOptions
func NewGetIntentOptions(workspaceID string, intent string) *GetIntentOptions {
    return &GetIntentOptions{
        WorkspaceID: workspaceID,
        Intent: intent,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *GetIntentOptions) SetWorkspaceID(param string) *GetIntentOptions {
    options.WorkspaceID = param
    return options
}

// SetIntent : Allow user to set Intent
func (options *GetIntentOptions) SetIntent(param string) *GetIntentOptions {
    options.Intent = param
    return options
}

// SetExport : Allow user to set Export
func (options *GetIntentOptions) SetExport(param bool) *GetIntentOptions {
    options.Export = param
    options.IsExportSet = true
    return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *GetIntentOptions) SetIncludeAudit(param bool) *GetIntentOptions {
    options.IncludeAudit = param
    options.IsIncludeAuditSet = true
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
	WorkspaceID string `json:"workspace_id"`

	// The name of the entity.
	Entity string `json:"entity"`

	// The text of the entity value.
	Value string `json:"value"`

	// The text of the synonym.
	Synonym string `json:"synonym"`

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit bool `json:"include_audit,omitempty"`

    // Indicates whether user set optional parameter IncludeAudit
    IsIncludeAuditSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewGetSynonymOptions : Instantiate GetSynonymOptions
func NewGetSynonymOptions(workspaceID string, entity string, value string, synonym string) *GetSynonymOptions {
    return &GetSynonymOptions{
        WorkspaceID: workspaceID,
        Entity: entity,
        Value: value,
        Synonym: synonym,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *GetSynonymOptions) SetWorkspaceID(param string) *GetSynonymOptions {
    options.WorkspaceID = param
    return options
}

// SetEntity : Allow user to set Entity
func (options *GetSynonymOptions) SetEntity(param string) *GetSynonymOptions {
    options.Entity = param
    return options
}

// SetValue : Allow user to set Value
func (options *GetSynonymOptions) SetValue(param string) *GetSynonymOptions {
    options.Value = param
    return options
}

// SetSynonym : Allow user to set Synonym
func (options *GetSynonymOptions) SetSynonym(param string) *GetSynonymOptions {
    options.Synonym = param
    return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *GetSynonymOptions) SetIncludeAudit(param bool) *GetSynonymOptions {
    options.IncludeAudit = param
    options.IsIncludeAuditSet = true
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
	WorkspaceID string `json:"workspace_id"`

	// The name of the entity.
	Entity string `json:"entity"`

	// The text of the entity value.
	Value string `json:"value"`

	// Whether to include all element content in the returned data. If **export**=`false`, the returned data includes only information about the element itself. If **export**=`true`, all content, including subelements, is included.
	Export bool `json:"export,omitempty"`

    // Indicates whether user set optional parameter Export
    IsExportSet bool

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit bool `json:"include_audit,omitempty"`

    // Indicates whether user set optional parameter IncludeAudit
    IsIncludeAuditSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewGetValueOptions : Instantiate GetValueOptions
func NewGetValueOptions(workspaceID string, entity string, value string) *GetValueOptions {
    return &GetValueOptions{
        WorkspaceID: workspaceID,
        Entity: entity,
        Value: value,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *GetValueOptions) SetWorkspaceID(param string) *GetValueOptions {
    options.WorkspaceID = param
    return options
}

// SetEntity : Allow user to set Entity
func (options *GetValueOptions) SetEntity(param string) *GetValueOptions {
    options.Entity = param
    return options
}

// SetValue : Allow user to set Value
func (options *GetValueOptions) SetValue(param string) *GetValueOptions {
    options.Value = param
    return options
}

// SetExport : Allow user to set Export
func (options *GetValueOptions) SetExport(param bool) *GetValueOptions {
    options.Export = param
    options.IsExportSet = true
    return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *GetValueOptions) SetIncludeAudit(param bool) *GetValueOptions {
    options.IncludeAudit = param
    options.IsIncludeAuditSet = true
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
	WorkspaceID string `json:"workspace_id"`

	// Whether to include all element content in the returned data. If **export**=`false`, the returned data includes only information about the element itself. If **export**=`true`, all content, including subelements, is included.
	Export bool `json:"export,omitempty"`

    // Indicates whether user set optional parameter Export
    IsExportSet bool

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit bool `json:"include_audit,omitempty"`

    // Indicates whether user set optional parameter IncludeAudit
    IsIncludeAuditSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewGetWorkspaceOptions : Instantiate GetWorkspaceOptions
func NewGetWorkspaceOptions(workspaceID string) *GetWorkspaceOptions {
    return &GetWorkspaceOptions{
        WorkspaceID: workspaceID,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *GetWorkspaceOptions) SetWorkspaceID(param string) *GetWorkspaceOptions {
    options.WorkspaceID = param
    return options
}

// SetExport : Allow user to set Export
func (options *GetWorkspaceOptions) SetExport(param bool) *GetWorkspaceOptions {
    options.Export = param
    options.IsExportSet = true
    return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *GetWorkspaceOptions) SetIncludeAudit(param bool) *GetWorkspaceOptions {
    options.IncludeAudit = param
    options.IsIncludeAuditSet = true
    return options
}

// SetHeaders : Allow user to set Headers
func (options *GetWorkspaceOptions) SetHeaders(param map[string]string) *GetWorkspaceOptions {
    options.Headers = param
    return options
}

// InputData : The user input.
type InputData struct {

	// The text of the user input. This string cannot contain carriage return, newline, or tab characters, and it must be no longer than 2048 characters.
	Text string `json:"text"`
}

// Intent : Intent struct
type Intent struct {

	// The name of the intent.
	IntentName string `json:"intent"`

	// The timestamp for creation of the intent.
	Created strfmt.DateTime `json:"created,omitempty"`

	// The timestamp for the last update to the intent.
	Updated strfmt.DateTime `json:"updated,omitempty"`

	// The description of the intent.
	Description string `json:"description,omitempty"`
}

// IntentCollection : IntentCollection struct
type IntentCollection struct {

	// An array of objects describing the intents defined for the workspace.
	Intents []IntentExport `json:"intents"`

	// The pagination data for the returned objects.
	Pagination Pagination `json:"pagination"`
}

// IntentExport : IntentExport struct
type IntentExport struct {

	// The name of the intent.
	IntentName string `json:"intent"`

	// The timestamp for creation of the intent.
	Created strfmt.DateTime `json:"created,omitempty"`

	// The timestamp for the last update to the intent.
	Updated strfmt.DateTime `json:"updated,omitempty"`

	// The description of the intent.
	Description string `json:"description,omitempty"`

	// An array of objects describing the user input examples for the intent.
	Examples []Example `json:"examples,omitempty"`
}

// ListAllLogsOptions : The listAllLogs options.
type ListAllLogsOptions struct {

	// A cacheable parameter that limits the results to those matching the specified filter. You must specify a filter query that includes a value for `language`, as well as a value for `workspace_id` or `request.context.metadata.deployment`. For more information, see the [documentation](https://console.bluemix.net/docs/services/conversation/filter-reference.html#filter-query-syntax).
	Filter string `json:"filter"`

	// The attribute by which returned results will be sorted. To reverse the sort order, prefix the value with a minus sign (`-`). Supported values are `name`, `updated`, and `workspace_id`.
	Sort string `json:"sort,omitempty"`

    // Indicates whether user set optional parameter Sort
    IsSortSet bool

	// The number of records to return in each page of results.
	PageLimit int64 `json:"page_limit,omitempty"`

    // Indicates whether user set optional parameter PageLimit
    IsPageLimitSet bool

	// A token identifying the page of results to retrieve.
	Cursor string `json:"cursor,omitempty"`

    // Indicates whether user set optional parameter Cursor
    IsCursorSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewListAllLogsOptions : Instantiate ListAllLogsOptions
func NewListAllLogsOptions(filter string) *ListAllLogsOptions {
    return &ListAllLogsOptions{
        Filter: filter,
    }
}

// SetFilter : Allow user to set Filter
func (options *ListAllLogsOptions) SetFilter(param string) *ListAllLogsOptions {
    options.Filter = param
    return options
}

// SetSort : Allow user to set Sort
func (options *ListAllLogsOptions) SetSort(param string) *ListAllLogsOptions {
    options.Sort = param
    options.IsSortSet = true
    return options
}

// SetPageLimit : Allow user to set PageLimit
func (options *ListAllLogsOptions) SetPageLimit(param int64) *ListAllLogsOptions {
    options.PageLimit = param
    options.IsPageLimitSet = true
    return options
}

// SetCursor : Allow user to set Cursor
func (options *ListAllLogsOptions) SetCursor(param string) *ListAllLogsOptions {
    options.Cursor = param
    options.IsCursorSet = true
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
	WorkspaceID string `json:"workspace_id"`

	// The number of records to return in each page of results.
	PageLimit int64 `json:"page_limit,omitempty"`

    // Indicates whether user set optional parameter PageLimit
    IsPageLimitSet bool

	// Whether to include information about the number of records returned.
	IncludeCount bool `json:"include_count,omitempty"`

    // Indicates whether user set optional parameter IncludeCount
    IsIncludeCountSet bool

	// The attribute by which returned results will be sorted. To reverse the sort order, prefix the value with a minus sign (`-`). Supported values are `name`, `updated`, and `workspace_id`.
	Sort string `json:"sort,omitempty"`

    // Indicates whether user set optional parameter Sort
    IsSortSet bool

	// A token identifying the page of results to retrieve.
	Cursor string `json:"cursor,omitempty"`

    // Indicates whether user set optional parameter Cursor
    IsCursorSet bool

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit bool `json:"include_audit,omitempty"`

    // Indicates whether user set optional parameter IncludeAudit
    IsIncludeAuditSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewListCounterexamplesOptions : Instantiate ListCounterexamplesOptions
func NewListCounterexamplesOptions(workspaceID string) *ListCounterexamplesOptions {
    return &ListCounterexamplesOptions{
        WorkspaceID: workspaceID,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *ListCounterexamplesOptions) SetWorkspaceID(param string) *ListCounterexamplesOptions {
    options.WorkspaceID = param
    return options
}

// SetPageLimit : Allow user to set PageLimit
func (options *ListCounterexamplesOptions) SetPageLimit(param int64) *ListCounterexamplesOptions {
    options.PageLimit = param
    options.IsPageLimitSet = true
    return options
}

// SetIncludeCount : Allow user to set IncludeCount
func (options *ListCounterexamplesOptions) SetIncludeCount(param bool) *ListCounterexamplesOptions {
    options.IncludeCount = param
    options.IsIncludeCountSet = true
    return options
}

// SetSort : Allow user to set Sort
func (options *ListCounterexamplesOptions) SetSort(param string) *ListCounterexamplesOptions {
    options.Sort = param
    options.IsSortSet = true
    return options
}

// SetCursor : Allow user to set Cursor
func (options *ListCounterexamplesOptions) SetCursor(param string) *ListCounterexamplesOptions {
    options.Cursor = param
    options.IsCursorSet = true
    return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *ListCounterexamplesOptions) SetIncludeAudit(param bool) *ListCounterexamplesOptions {
    options.IncludeAudit = param
    options.IsIncludeAuditSet = true
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
	WorkspaceID string `json:"workspace_id"`

	// The number of records to return in each page of results.
	PageLimit int64 `json:"page_limit,omitempty"`

    // Indicates whether user set optional parameter PageLimit
    IsPageLimitSet bool

	// Whether to include information about the number of records returned.
	IncludeCount bool `json:"include_count,omitempty"`

    // Indicates whether user set optional parameter IncludeCount
    IsIncludeCountSet bool

	// The attribute by which returned results will be sorted. To reverse the sort order, prefix the value with a minus sign (`-`). Supported values are `name`, `updated`, and `workspace_id`.
	Sort string `json:"sort,omitempty"`

    // Indicates whether user set optional parameter Sort
    IsSortSet bool

	// A token identifying the page of results to retrieve.
	Cursor string `json:"cursor,omitempty"`

    // Indicates whether user set optional parameter Cursor
    IsCursorSet bool

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit bool `json:"include_audit,omitempty"`

    // Indicates whether user set optional parameter IncludeAudit
    IsIncludeAuditSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewListDialogNodesOptions : Instantiate ListDialogNodesOptions
func NewListDialogNodesOptions(workspaceID string) *ListDialogNodesOptions {
    return &ListDialogNodesOptions{
        WorkspaceID: workspaceID,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *ListDialogNodesOptions) SetWorkspaceID(param string) *ListDialogNodesOptions {
    options.WorkspaceID = param
    return options
}

// SetPageLimit : Allow user to set PageLimit
func (options *ListDialogNodesOptions) SetPageLimit(param int64) *ListDialogNodesOptions {
    options.PageLimit = param
    options.IsPageLimitSet = true
    return options
}

// SetIncludeCount : Allow user to set IncludeCount
func (options *ListDialogNodesOptions) SetIncludeCount(param bool) *ListDialogNodesOptions {
    options.IncludeCount = param
    options.IsIncludeCountSet = true
    return options
}

// SetSort : Allow user to set Sort
func (options *ListDialogNodesOptions) SetSort(param string) *ListDialogNodesOptions {
    options.Sort = param
    options.IsSortSet = true
    return options
}

// SetCursor : Allow user to set Cursor
func (options *ListDialogNodesOptions) SetCursor(param string) *ListDialogNodesOptions {
    options.Cursor = param
    options.IsCursorSet = true
    return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *ListDialogNodesOptions) SetIncludeAudit(param bool) *ListDialogNodesOptions {
    options.IncludeAudit = param
    options.IsIncludeAuditSet = true
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
	WorkspaceID string `json:"workspace_id"`

	// Whether to include all element content in the returned data. If **export**=`false`, the returned data includes only information about the element itself. If **export**=`true`, all content, including subelements, is included.
	Export bool `json:"export,omitempty"`

    // Indicates whether user set optional parameter Export
    IsExportSet bool

	// The number of records to return in each page of results.
	PageLimit int64 `json:"page_limit,omitempty"`

    // Indicates whether user set optional parameter PageLimit
    IsPageLimitSet bool

	// Whether to include information about the number of records returned.
	IncludeCount bool `json:"include_count,omitempty"`

    // Indicates whether user set optional parameter IncludeCount
    IsIncludeCountSet bool

	// The attribute by which returned results will be sorted. To reverse the sort order, prefix the value with a minus sign (`-`). Supported values are `name`, `updated`, and `workspace_id`.
	Sort string `json:"sort,omitempty"`

    // Indicates whether user set optional parameter Sort
    IsSortSet bool

	// A token identifying the page of results to retrieve.
	Cursor string `json:"cursor,omitempty"`

    // Indicates whether user set optional parameter Cursor
    IsCursorSet bool

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit bool `json:"include_audit,omitempty"`

    // Indicates whether user set optional parameter IncludeAudit
    IsIncludeAuditSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewListEntitiesOptions : Instantiate ListEntitiesOptions
func NewListEntitiesOptions(workspaceID string) *ListEntitiesOptions {
    return &ListEntitiesOptions{
        WorkspaceID: workspaceID,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *ListEntitiesOptions) SetWorkspaceID(param string) *ListEntitiesOptions {
    options.WorkspaceID = param
    return options
}

// SetExport : Allow user to set Export
func (options *ListEntitiesOptions) SetExport(param bool) *ListEntitiesOptions {
    options.Export = param
    options.IsExportSet = true
    return options
}

// SetPageLimit : Allow user to set PageLimit
func (options *ListEntitiesOptions) SetPageLimit(param int64) *ListEntitiesOptions {
    options.PageLimit = param
    options.IsPageLimitSet = true
    return options
}

// SetIncludeCount : Allow user to set IncludeCount
func (options *ListEntitiesOptions) SetIncludeCount(param bool) *ListEntitiesOptions {
    options.IncludeCount = param
    options.IsIncludeCountSet = true
    return options
}

// SetSort : Allow user to set Sort
func (options *ListEntitiesOptions) SetSort(param string) *ListEntitiesOptions {
    options.Sort = param
    options.IsSortSet = true
    return options
}

// SetCursor : Allow user to set Cursor
func (options *ListEntitiesOptions) SetCursor(param string) *ListEntitiesOptions {
    options.Cursor = param
    options.IsCursorSet = true
    return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *ListEntitiesOptions) SetIncludeAudit(param bool) *ListEntitiesOptions {
    options.IncludeAudit = param
    options.IsIncludeAuditSet = true
    return options
}

// SetHeaders : Allow user to set Headers
func (options *ListEntitiesOptions) SetHeaders(param map[string]string) *ListEntitiesOptions {
    options.Headers = param
    return options
}

// ListEntityMentionsOptions : The listEntityMentions options.
type ListEntityMentionsOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID string `json:"workspace_id"`

	// The name of the entity.
	Entity string `json:"entity"`

	// Whether to include all element content in the returned data. If **export**=`false`, the returned data includes only information about the element itself. If **export**=`true`, all content, including subelements, is included.
	Export bool `json:"export,omitempty"`

    // Indicates whether user set optional parameter Export
    IsExportSet bool

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit bool `json:"include_audit,omitempty"`

    // Indicates whether user set optional parameter IncludeAudit
    IsIncludeAuditSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewListEntityMentionsOptions : Instantiate ListEntityMentionsOptions
func NewListEntityMentionsOptions(workspaceID string, entity string) *ListEntityMentionsOptions {
    return &ListEntityMentionsOptions{
        WorkspaceID: workspaceID,
        Entity: entity,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *ListEntityMentionsOptions) SetWorkspaceID(param string) *ListEntityMentionsOptions {
    options.WorkspaceID = param
    return options
}

// SetEntity : Allow user to set Entity
func (options *ListEntityMentionsOptions) SetEntity(param string) *ListEntityMentionsOptions {
    options.Entity = param
    return options
}

// SetExport : Allow user to set Export
func (options *ListEntityMentionsOptions) SetExport(param bool) *ListEntityMentionsOptions {
    options.Export = param
    options.IsExportSet = true
    return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *ListEntityMentionsOptions) SetIncludeAudit(param bool) *ListEntityMentionsOptions {
    options.IncludeAudit = param
    options.IsIncludeAuditSet = true
    return options
}

// SetHeaders : Allow user to set Headers
func (options *ListEntityMentionsOptions) SetHeaders(param map[string]string) *ListEntityMentionsOptions {
    options.Headers = param
    return options
}

// ListExamplesOptions : The listExamples options.
type ListExamplesOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID string `json:"workspace_id"`

	// The intent name.
	Intent string `json:"intent"`

	// The number of records to return in each page of results.
	PageLimit int64 `json:"page_limit,omitempty"`

    // Indicates whether user set optional parameter PageLimit
    IsPageLimitSet bool

	// Whether to include information about the number of records returned.
	IncludeCount bool `json:"include_count,omitempty"`

    // Indicates whether user set optional parameter IncludeCount
    IsIncludeCountSet bool

	// The attribute by which returned results will be sorted. To reverse the sort order, prefix the value with a minus sign (`-`). Supported values are `name`, `updated`, and `workspace_id`.
	Sort string `json:"sort,omitempty"`

    // Indicates whether user set optional parameter Sort
    IsSortSet bool

	// A token identifying the page of results to retrieve.
	Cursor string `json:"cursor,omitempty"`

    // Indicates whether user set optional parameter Cursor
    IsCursorSet bool

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit bool `json:"include_audit,omitempty"`

    // Indicates whether user set optional parameter IncludeAudit
    IsIncludeAuditSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewListExamplesOptions : Instantiate ListExamplesOptions
func NewListExamplesOptions(workspaceID string, intent string) *ListExamplesOptions {
    return &ListExamplesOptions{
        WorkspaceID: workspaceID,
        Intent: intent,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *ListExamplesOptions) SetWorkspaceID(param string) *ListExamplesOptions {
    options.WorkspaceID = param
    return options
}

// SetIntent : Allow user to set Intent
func (options *ListExamplesOptions) SetIntent(param string) *ListExamplesOptions {
    options.Intent = param
    return options
}

// SetPageLimit : Allow user to set PageLimit
func (options *ListExamplesOptions) SetPageLimit(param int64) *ListExamplesOptions {
    options.PageLimit = param
    options.IsPageLimitSet = true
    return options
}

// SetIncludeCount : Allow user to set IncludeCount
func (options *ListExamplesOptions) SetIncludeCount(param bool) *ListExamplesOptions {
    options.IncludeCount = param
    options.IsIncludeCountSet = true
    return options
}

// SetSort : Allow user to set Sort
func (options *ListExamplesOptions) SetSort(param string) *ListExamplesOptions {
    options.Sort = param
    options.IsSortSet = true
    return options
}

// SetCursor : Allow user to set Cursor
func (options *ListExamplesOptions) SetCursor(param string) *ListExamplesOptions {
    options.Cursor = param
    options.IsCursorSet = true
    return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *ListExamplesOptions) SetIncludeAudit(param bool) *ListExamplesOptions {
    options.IncludeAudit = param
    options.IsIncludeAuditSet = true
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
	WorkspaceID string `json:"workspace_id"`

	// Whether to include all element content in the returned data. If **export**=`false`, the returned data includes only information about the element itself. If **export**=`true`, all content, including subelements, is included.
	Export bool `json:"export,omitempty"`

    // Indicates whether user set optional parameter Export
    IsExportSet bool

	// The number of records to return in each page of results.
	PageLimit int64 `json:"page_limit,omitempty"`

    // Indicates whether user set optional parameter PageLimit
    IsPageLimitSet bool

	// Whether to include information about the number of records returned.
	IncludeCount bool `json:"include_count,omitempty"`

    // Indicates whether user set optional parameter IncludeCount
    IsIncludeCountSet bool

	// The attribute by which returned results will be sorted. To reverse the sort order, prefix the value with a minus sign (`-`). Supported values are `name`, `updated`, and `workspace_id`.
	Sort string `json:"sort,omitempty"`

    // Indicates whether user set optional parameter Sort
    IsSortSet bool

	// A token identifying the page of results to retrieve.
	Cursor string `json:"cursor,omitempty"`

    // Indicates whether user set optional parameter Cursor
    IsCursorSet bool

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit bool `json:"include_audit,omitempty"`

    // Indicates whether user set optional parameter IncludeAudit
    IsIncludeAuditSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewListIntentsOptions : Instantiate ListIntentsOptions
func NewListIntentsOptions(workspaceID string) *ListIntentsOptions {
    return &ListIntentsOptions{
        WorkspaceID: workspaceID,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *ListIntentsOptions) SetWorkspaceID(param string) *ListIntentsOptions {
    options.WorkspaceID = param
    return options
}

// SetExport : Allow user to set Export
func (options *ListIntentsOptions) SetExport(param bool) *ListIntentsOptions {
    options.Export = param
    options.IsExportSet = true
    return options
}

// SetPageLimit : Allow user to set PageLimit
func (options *ListIntentsOptions) SetPageLimit(param int64) *ListIntentsOptions {
    options.PageLimit = param
    options.IsPageLimitSet = true
    return options
}

// SetIncludeCount : Allow user to set IncludeCount
func (options *ListIntentsOptions) SetIncludeCount(param bool) *ListIntentsOptions {
    options.IncludeCount = param
    options.IsIncludeCountSet = true
    return options
}

// SetSort : Allow user to set Sort
func (options *ListIntentsOptions) SetSort(param string) *ListIntentsOptions {
    options.Sort = param
    options.IsSortSet = true
    return options
}

// SetCursor : Allow user to set Cursor
func (options *ListIntentsOptions) SetCursor(param string) *ListIntentsOptions {
    options.Cursor = param
    options.IsCursorSet = true
    return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *ListIntentsOptions) SetIncludeAudit(param bool) *ListIntentsOptions {
    options.IncludeAudit = param
    options.IsIncludeAuditSet = true
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
	WorkspaceID string `json:"workspace_id"`

	// The attribute by which returned results will be sorted. To reverse the sort order, prefix the value with a minus sign (`-`). Supported values are `name`, `updated`, and `workspace_id`.
	Sort string `json:"sort,omitempty"`

    // Indicates whether user set optional parameter Sort
    IsSortSet bool

	// A cacheable parameter that limits the results to those matching the specified filter. For more information, see the [documentation](https://console.bluemix.net/docs/services/conversation/filter-reference.html#filter-query-syntax).
	Filter string `json:"filter,omitempty"`

    // Indicates whether user set optional parameter Filter
    IsFilterSet bool

	// The number of records to return in each page of results.
	PageLimit int64 `json:"page_limit,omitempty"`

    // Indicates whether user set optional parameter PageLimit
    IsPageLimitSet bool

	// A token identifying the page of results to retrieve.
	Cursor string `json:"cursor,omitempty"`

    // Indicates whether user set optional parameter Cursor
    IsCursorSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewListLogsOptions : Instantiate ListLogsOptions
func NewListLogsOptions(workspaceID string) *ListLogsOptions {
    return &ListLogsOptions{
        WorkspaceID: workspaceID,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *ListLogsOptions) SetWorkspaceID(param string) *ListLogsOptions {
    options.WorkspaceID = param
    return options
}

// SetSort : Allow user to set Sort
func (options *ListLogsOptions) SetSort(param string) *ListLogsOptions {
    options.Sort = param
    options.IsSortSet = true
    return options
}

// SetFilter : Allow user to set Filter
func (options *ListLogsOptions) SetFilter(param string) *ListLogsOptions {
    options.Filter = param
    options.IsFilterSet = true
    return options
}

// SetPageLimit : Allow user to set PageLimit
func (options *ListLogsOptions) SetPageLimit(param int64) *ListLogsOptions {
    options.PageLimit = param
    options.IsPageLimitSet = true
    return options
}

// SetCursor : Allow user to set Cursor
func (options *ListLogsOptions) SetCursor(param string) *ListLogsOptions {
    options.Cursor = param
    options.IsCursorSet = true
    return options
}

// SetHeaders : Allow user to set Headers
func (options *ListLogsOptions) SetHeaders(param map[string]string) *ListLogsOptions {
    options.Headers = param
    return options
}

// ListSynonymsOptions : The listSynonyms options.
type ListSynonymsOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID string `json:"workspace_id"`

	// The name of the entity.
	Entity string `json:"entity"`

	// The text of the entity value.
	Value string `json:"value"`

	// The number of records to return in each page of results.
	PageLimit int64 `json:"page_limit,omitempty"`

    // Indicates whether user set optional parameter PageLimit
    IsPageLimitSet bool

	// Whether to include information about the number of records returned.
	IncludeCount bool `json:"include_count,omitempty"`

    // Indicates whether user set optional parameter IncludeCount
    IsIncludeCountSet bool

	// The attribute by which returned results will be sorted. To reverse the sort order, prefix the value with a minus sign (`-`). Supported values are `name`, `updated`, and `workspace_id`.
	Sort string `json:"sort,omitempty"`

    // Indicates whether user set optional parameter Sort
    IsSortSet bool

	// A token identifying the page of results to retrieve.
	Cursor string `json:"cursor,omitempty"`

    // Indicates whether user set optional parameter Cursor
    IsCursorSet bool

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit bool `json:"include_audit,omitempty"`

    // Indicates whether user set optional parameter IncludeAudit
    IsIncludeAuditSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewListSynonymsOptions : Instantiate ListSynonymsOptions
func NewListSynonymsOptions(workspaceID string, entity string, value string) *ListSynonymsOptions {
    return &ListSynonymsOptions{
        WorkspaceID: workspaceID,
        Entity: entity,
        Value: value,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *ListSynonymsOptions) SetWorkspaceID(param string) *ListSynonymsOptions {
    options.WorkspaceID = param
    return options
}

// SetEntity : Allow user to set Entity
func (options *ListSynonymsOptions) SetEntity(param string) *ListSynonymsOptions {
    options.Entity = param
    return options
}

// SetValue : Allow user to set Value
func (options *ListSynonymsOptions) SetValue(param string) *ListSynonymsOptions {
    options.Value = param
    return options
}

// SetPageLimit : Allow user to set PageLimit
func (options *ListSynonymsOptions) SetPageLimit(param int64) *ListSynonymsOptions {
    options.PageLimit = param
    options.IsPageLimitSet = true
    return options
}

// SetIncludeCount : Allow user to set IncludeCount
func (options *ListSynonymsOptions) SetIncludeCount(param bool) *ListSynonymsOptions {
    options.IncludeCount = param
    options.IsIncludeCountSet = true
    return options
}

// SetSort : Allow user to set Sort
func (options *ListSynonymsOptions) SetSort(param string) *ListSynonymsOptions {
    options.Sort = param
    options.IsSortSet = true
    return options
}

// SetCursor : Allow user to set Cursor
func (options *ListSynonymsOptions) SetCursor(param string) *ListSynonymsOptions {
    options.Cursor = param
    options.IsCursorSet = true
    return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *ListSynonymsOptions) SetIncludeAudit(param bool) *ListSynonymsOptions {
    options.IncludeAudit = param
    options.IsIncludeAuditSet = true
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
	WorkspaceID string `json:"workspace_id"`

	// The name of the entity.
	Entity string `json:"entity"`

	// Whether to include all element content in the returned data. If **export**=`false`, the returned data includes only information about the element itself. If **export**=`true`, all content, including subelements, is included.
	Export bool `json:"export,omitempty"`

    // Indicates whether user set optional parameter Export
    IsExportSet bool

	// The number of records to return in each page of results.
	PageLimit int64 `json:"page_limit,omitempty"`

    // Indicates whether user set optional parameter PageLimit
    IsPageLimitSet bool

	// Whether to include information about the number of records returned.
	IncludeCount bool `json:"include_count,omitempty"`

    // Indicates whether user set optional parameter IncludeCount
    IsIncludeCountSet bool

	// The attribute by which returned results will be sorted. To reverse the sort order, prefix the value with a minus sign (`-`). Supported values are `name`, `updated`, and `workspace_id`.
	Sort string `json:"sort,omitempty"`

    // Indicates whether user set optional parameter Sort
    IsSortSet bool

	// A token identifying the page of results to retrieve.
	Cursor string `json:"cursor,omitempty"`

    // Indicates whether user set optional parameter Cursor
    IsCursorSet bool

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit bool `json:"include_audit,omitempty"`

    // Indicates whether user set optional parameter IncludeAudit
    IsIncludeAuditSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewListValuesOptions : Instantiate ListValuesOptions
func NewListValuesOptions(workspaceID string, entity string) *ListValuesOptions {
    return &ListValuesOptions{
        WorkspaceID: workspaceID,
        Entity: entity,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *ListValuesOptions) SetWorkspaceID(param string) *ListValuesOptions {
    options.WorkspaceID = param
    return options
}

// SetEntity : Allow user to set Entity
func (options *ListValuesOptions) SetEntity(param string) *ListValuesOptions {
    options.Entity = param
    return options
}

// SetExport : Allow user to set Export
func (options *ListValuesOptions) SetExport(param bool) *ListValuesOptions {
    options.Export = param
    options.IsExportSet = true
    return options
}

// SetPageLimit : Allow user to set PageLimit
func (options *ListValuesOptions) SetPageLimit(param int64) *ListValuesOptions {
    options.PageLimit = param
    options.IsPageLimitSet = true
    return options
}

// SetIncludeCount : Allow user to set IncludeCount
func (options *ListValuesOptions) SetIncludeCount(param bool) *ListValuesOptions {
    options.IncludeCount = param
    options.IsIncludeCountSet = true
    return options
}

// SetSort : Allow user to set Sort
func (options *ListValuesOptions) SetSort(param string) *ListValuesOptions {
    options.Sort = param
    options.IsSortSet = true
    return options
}

// SetCursor : Allow user to set Cursor
func (options *ListValuesOptions) SetCursor(param string) *ListValuesOptions {
    options.Cursor = param
    options.IsCursorSet = true
    return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *ListValuesOptions) SetIncludeAudit(param bool) *ListValuesOptions {
    options.IncludeAudit = param
    options.IsIncludeAuditSet = true
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
	PageLimit int64 `json:"page_limit,omitempty"`

    // Indicates whether user set optional parameter PageLimit
    IsPageLimitSet bool

	// Whether to include information about the number of records returned.
	IncludeCount bool `json:"include_count,omitempty"`

    // Indicates whether user set optional parameter IncludeCount
    IsIncludeCountSet bool

	// The attribute by which returned results will be sorted. To reverse the sort order, prefix the value with a minus sign (`-`). Supported values are `name`, `updated`, and `workspace_id`.
	Sort string `json:"sort,omitempty"`

    // Indicates whether user set optional parameter Sort
    IsSortSet bool

	// A token identifying the page of results to retrieve.
	Cursor string `json:"cursor,omitempty"`

    // Indicates whether user set optional parameter Cursor
    IsCursorSet bool

	// Whether to include the audit properties (`created` and `updated` timestamps) in the response.
	IncludeAudit bool `json:"include_audit,omitempty"`

    // Indicates whether user set optional parameter IncludeAudit
    IsIncludeAuditSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewListWorkspacesOptions : Instantiate ListWorkspacesOptions
func NewListWorkspacesOptions() *ListWorkspacesOptions {
    return &ListWorkspacesOptions{}
}

// SetPageLimit : Allow user to set PageLimit
func (options *ListWorkspacesOptions) SetPageLimit(param int64) *ListWorkspacesOptions {
    options.PageLimit = param
    options.IsPageLimitSet = true
    return options
}

// SetIncludeCount : Allow user to set IncludeCount
func (options *ListWorkspacesOptions) SetIncludeCount(param bool) *ListWorkspacesOptions {
    options.IncludeCount = param
    options.IsIncludeCountSet = true
    return options
}

// SetSort : Allow user to set Sort
func (options *ListWorkspacesOptions) SetSort(param string) *ListWorkspacesOptions {
    options.Sort = param
    options.IsSortSet = true
    return options
}

// SetCursor : Allow user to set Cursor
func (options *ListWorkspacesOptions) SetCursor(param string) *ListWorkspacesOptions {
    options.Cursor = param
    options.IsCursorSet = true
    return options
}

// SetIncludeAudit : Allow user to set IncludeAudit
func (options *ListWorkspacesOptions) SetIncludeAudit(param bool) *ListWorkspacesOptions {
    options.IncludeAudit = param
    options.IsIncludeAuditSet = true
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
	Logs []LogExport `json:"logs"`

	// The pagination data for the returned objects.
	Pagination LogPagination `json:"pagination"`
}

// LogExport : LogExport struct
type LogExport struct {

	// A request received by the workspace, including the user input and context.
	Request MessageRequest `json:"request"`

	// The response sent by the workspace, including the output text, detected intents and entities, and context.
	Response MessageResponse `json:"response"`

	// A unique identifier for the logged event.
	LogID string `json:"log_id"`

	// The timestamp for receipt of the message.
	RequestTimestamp string `json:"request_timestamp"`

	// The timestamp for the system response to the message.
	ResponseTimestamp string `json:"response_timestamp"`

	// The unique identifier of the workspace where the request was made.
	WorkspaceID string `json:"workspace_id"`

	// The language of the workspace where the message request was made.
	Language string `json:"language"`
}

// LogMessage : Log message details.
type LogMessage struct {

	// The severity of the log message.
	Level string `json:"level"`

	// The text of the log message.
	Msg string `json:"msg"`
}

// LogPagination : The pagination data for the returned objects.
type LogPagination struct {

	// The URL that will return the next page of results, if any.
	NextURL string `json:"next_url,omitempty"`

	// Reserved for future use.
	Matched int64 `json:"matched,omitempty"`

	// A token identifying the next page of results.
	NextCursor string `json:"next_cursor,omitempty"`
}

// Mentions : A mention of a contextual entity.
type Mentions struct {

	// The name of the entity.
	Entity string `json:"entity"`

	// An array of zero-based character offsets that indicate where the entity mentions begin and end in the input text.
	Location []int64 `json:"location"`
}

// MessageInput : The text of the user input.
type MessageInput struct {

	// The user's input.
	Text string `json:"text,omitempty"`
}

// MessageOptions : The message options.
type MessageOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID string `json:"workspace_id"`

	// An input object that includes the input text.
	Input InputData `json:"input,omitempty"`

    // Indicates whether user set optional parameter Input
    IsInputSet bool

	// Whether to return more than one intent. Set to `true` to return all matching intents.
	AlternateIntents bool `json:"alternate_intents,omitempty"`

    // Indicates whether user set optional parameter AlternateIntents
    IsAlternateIntentsSet bool

	// State information for the conversation. Continue a conversation by including the context object from the previous response.
	Context Context `json:"context,omitempty"`

    // Indicates whether user set optional parameter Context
    IsContextSet bool

	// Entities to use when evaluating the message. Include entities from the previous response to continue using those entities rather than detecting entities in the new input.
	Entities []RuntimeEntity `json:"entities,omitempty"`

    // Indicates whether user set optional parameter Entities
    IsEntitiesSet bool

	// Intents to use when evaluating the user input. Include intents from the previous response to continue using those intents rather than trying to recognize intents in the new input.
	Intents []RuntimeIntent `json:"intents,omitempty"`

    // Indicates whether user set optional parameter Intents
    IsIntentsSet bool

	// System output. Include the output from the previous response to maintain intermediate information over multiple requests.
	Output OutputData `json:"output,omitempty"`

    // Indicates whether user set optional parameter Output
    IsOutputSet bool

	// Whether to include additional diagnostic information about the dialog nodes that were visited during processing of the message.
	NodesVisitedDetails bool `json:"nodes_visited_details,omitempty"`

    // Indicates whether user set optional parameter NodesVisitedDetails
    IsNodesVisitedDetailsSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewMessageOptions : Instantiate MessageOptions
func NewMessageOptions(workspaceID string) *MessageOptions {
    return &MessageOptions{
        WorkspaceID: workspaceID,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *MessageOptions) SetWorkspaceID(param string) *MessageOptions {
    options.WorkspaceID = param
    return options
}

// SetInput : Allow user to set Input
func (options *MessageOptions) SetInput(param InputData) *MessageOptions {
    options.Input = param
    options.IsInputSet = true
    return options
}

// SetAlternateIntents : Allow user to set AlternateIntents
func (options *MessageOptions) SetAlternateIntents(param bool) *MessageOptions {
    options.AlternateIntents = param
    options.IsAlternateIntentsSet = true
    return options
}

// SetContext : Allow user to set Context
func (options *MessageOptions) SetContext(param Context) *MessageOptions {
    options.Context = param
    options.IsContextSet = true
    return options
}

// SetEntities : Allow user to set Entities
func (options *MessageOptions) SetEntities(param []RuntimeEntity) *MessageOptions {
    options.Entities = param
    options.IsEntitiesSet = true
    return options
}

// SetIntents : Allow user to set Intents
func (options *MessageOptions) SetIntents(param []RuntimeIntent) *MessageOptions {
    options.Intents = param
    options.IsIntentsSet = true
    return options
}

// SetOutput : Allow user to set Output
func (options *MessageOptions) SetOutput(param OutputData) *MessageOptions {
    options.Output = param
    options.IsOutputSet = true
    return options
}

// SetNodesVisitedDetails : Allow user to set NodesVisitedDetails
func (options *MessageOptions) SetNodesVisitedDetails(param bool) *MessageOptions {
    options.NodesVisitedDetails = param
    options.IsNodesVisitedDetailsSet = true
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
	Input InputData `json:"input,omitempty"`

	// Whether to return more than one intent. Set to `true` to return all matching intents.
	AlternateIntents bool `json:"alternate_intents,omitempty"`

	// State information for the conversation. Continue a conversation by including the context object from the previous response.
	Context Context `json:"context,omitempty"`

	// Entities to use when evaluating the message. Include entities from the previous response to continue using those entities rather than detecting entities in the new input.
	Entities []RuntimeEntity `json:"entities,omitempty"`

	// Intents to use when evaluating the user input. Include intents from the previous response to continue using those intents rather than trying to recognize intents in the new input.
	Intents []RuntimeIntent `json:"intents,omitempty"`

	// System output. Include the output from the previous response to maintain intermediate information over multiple requests.
	Output OutputData `json:"output,omitempty"`
}

// MessageResponse : A response from the Watson Assistant service.
type MessageResponse struct {

	// The user input from the request.
	Input MessageInput `json:"input,omitempty"`

	// An array of intents recognized in the user input, sorted in descending order of confidence.
	Intents []RuntimeIntent `json:"intents"`

	// An array of entities identified in the user input.
	Entities []RuntimeEntity `json:"entities"`

	// Whether to return more than one intent. A value of `true` indicates that all matching intents are returned.
	AlternateIntents bool `json:"alternate_intents,omitempty"`

	// State information for the conversation.
	Context Context `json:"context"`

	// Output from the dialog, including the response to the user, the nodes that were triggered, and log messages.
	Output OutputData `json:"output"`
}

// OutputData : An output object that includes the response to the user, the dialog nodes that were triggered, and messages from the log.
type OutputData struct {

	// An array of up to 50 messages logged with the request.
	LogMessages []LogMessage `json:"log_messages"`

	// An array of responses to the user.
	Text []string `json:"text"`

	// Output intended for any channel. It is the responsibility of the client application to implement the supported response types.
	Generic []DialogRuntimeResponseGeneric `json:"generic,omitempty"`

	// An array of the nodes that were triggered to create the response, in the order in which they were visited. This information is useful for debugging and for tracing the path taken through the node tree.
	NodesVisited []string `json:"nodes_visited,omitempty"`

	// An array of objects containing detailed diagnostic information about the nodes that were triggered during processing of the input message. Included only if **nodes_visited_details** is set to `true` in the message request.
	NodesVisitedDetails []DialogNodeVisitedDetails `json:"nodes_visited_details,omitempty"`
}

// Pagination : The pagination data for the returned objects.
type Pagination struct {

	// The URL that will return the same page of results.
	RefreshURL string `json:"refresh_url"`

	// The URL that will return the next page of results.
	NextURL string `json:"next_url,omitempty"`

	// Reserved for future use.
	Total int64 `json:"total,omitempty"`

	// Reserved for future use.
	Matched int64 `json:"matched,omitempty"`

	// A token identifying the current page of results.
	RefreshCursor string `json:"refresh_cursor,omitempty"`

	// A token identifying the next page of results.
	NextCursor string `json:"next_cursor,omitempty"`
}

// RuntimeEntity : A term from the request that was identified as an entity.
type RuntimeEntity struct {

	// An entity detected in the input.
	Entity string `json:"entity"`

	// An array of zero-based character offsets that indicate where the detected entity values begin and end in the input text.
	Location []int64 `json:"location"`

	// The term in the input text that was recognized as an entity value.
	Value string `json:"value"`

	// A decimal percentage that represents Watson's confidence in the entity.
	Confidence float64 `json:"confidence,omitempty"`

	// Any metadata for the entity.
	Metadata interface{} `json:"metadata,omitempty"`

	// The recognized capture groups for the entity, as defined by the entity pattern.
	Groups []CaptureGroup `json:"groups,omitempty"`
}

// RuntimeIntent : An intent identified in the user input.
type RuntimeIntent struct {

	// The name of the recognized intent.
	Intent string `json:"intent"`

	// A decimal percentage that represents Watson's confidence in the intent.
	Confidence float64 `json:"confidence"`
}

// Synonym : Synonym struct
type Synonym struct {

	// The text of the synonym.
	SynonymText string `json:"synonym"`

	// The timestamp for creation of the synonym.
	Created strfmt.DateTime `json:"created,omitempty"`

	// The timestamp for the most recent update to the synonym.
	Updated strfmt.DateTime `json:"updated,omitempty"`
}

// SynonymCollection : SynonymCollection struct
type SynonymCollection struct {

	// An array of synonyms.
	Synonyms []Synonym `json:"synonyms"`

	// The pagination data for the returned objects.
	Pagination Pagination `json:"pagination"`
}

// SystemResponse : For internal use only.
type SystemResponse struct {
}

// UpdateCounterexampleOptions : The updateCounterexample options.
type UpdateCounterexampleOptions struct {

	// Unique identifier of the workspace.
	WorkspaceID string `json:"workspace_id"`

	// The text of a user input counterexample (for example, `What are you wearing?`).
	Text string `json:"text"`

	// The text of a user input counterexample.
	NewText string `json:"text,omitempty"`

    // Indicates whether user set optional parameter NewText
    IsNewTextSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewUpdateCounterexampleOptions : Instantiate UpdateCounterexampleOptions
func NewUpdateCounterexampleOptions(workspaceID string, text string) *UpdateCounterexampleOptions {
    return &UpdateCounterexampleOptions{
        WorkspaceID: workspaceID,
        Text: text,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *UpdateCounterexampleOptions) SetWorkspaceID(param string) *UpdateCounterexampleOptions {
    options.WorkspaceID = param
    return options
}

// SetText : Allow user to set Text
func (options *UpdateCounterexampleOptions) SetText(param string) *UpdateCounterexampleOptions {
    options.Text = param
    return options
}

// SetNewText : Allow user to set NewText
func (options *UpdateCounterexampleOptions) SetNewText(param string) *UpdateCounterexampleOptions {
    options.NewText = param
    options.IsNewTextSet = true
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
	WorkspaceID string `json:"workspace_id"`

	// The dialog node ID (for example, `get_order`).
	DialogNode string `json:"dialog_node"`

	// How the dialog node is processed.
	NodeType string `json:"type,omitempty"`

    // Indicates whether user set optional parameter NodeType
    IsNodeTypeSet bool

	// The condition that will trigger the dialog node. This string cannot contain carriage return, newline, or tab characters, and it must be no longer than 2048 characters.
	NewConditions string `json:"conditions,omitempty"`

    // Indicates whether user set optional parameter NewConditions
    IsNewConditionsSet bool

	// An array of objects describing any actions to be invoked by the dialog node.
	NewActions []DialogNodeAction `json:"actions,omitempty"`

    // Indicates whether user set optional parameter NewActions
    IsNewActionsSet bool

	// The ID of the previous sibling dialog node.
	NewPreviousSibling string `json:"previous_sibling,omitempty"`

    // Indicates whether user set optional parameter NewPreviousSibling
    IsNewPreviousSiblingSet bool

	// The context for the dialog node.
	NewContext interface{} `json:"context,omitempty"`

    // Indicates whether user set optional parameter NewContext
    IsNewContextSet bool

	// The location in the dialog context where output is stored.
	NewVariable string `json:"variable,omitempty"`

    // Indicates whether user set optional parameter NewVariable
    IsNewVariableSet bool

	// A label that can be displayed externally to describe the purpose of the node to users.
	NewUserLabel string `json:"user_label,omitempty"`

    // Indicates whether user set optional parameter NewUserLabel
    IsNewUserLabelSet bool

	// The metadata for the dialog node.
	NewMetadata interface{} `json:"metadata,omitempty"`

    // Indicates whether user set optional parameter NewMetadata
    IsNewMetadataSet bool

	// The alias used to identify the dialog node. This string must conform to the following restrictions: - It can contain only Unicode alphanumeric, space, underscore, hyphen, and dot characters. - It must be no longer than 64 characters.
	NewTitle string `json:"title,omitempty"`

    // Indicates whether user set optional parameter NewTitle
    IsNewTitleSet bool

	// The description of the dialog node. This string cannot contain carriage return, newline, or tab characters, and it must be no longer than 128 characters.
	NewDescription string `json:"description,omitempty"`

    // Indicates whether user set optional parameter NewDescription
    IsNewDescriptionSet bool

	// Whether this dialog node can be returned to after a digression.
	NewDigressOut string `json:"digress_out,omitempty"`

    // Indicates whether user set optional parameter NewDigressOut
    IsNewDigressOutSet bool

	// How an `event_handler` node is processed.
	NewEventName string `json:"event_name,omitempty"`

    // Indicates whether user set optional parameter NewEventName
    IsNewEventNameSet bool

	// Whether the user can digress to top-level nodes while filling out slots.
	NewDigressOutSlots string `json:"digress_out_slots,omitempty"`

    // Indicates whether user set optional parameter NewDigressOutSlots
    IsNewDigressOutSlotsSet bool

	// The next step to be executed in dialog processing.
	NewNextStep DialogNodeNextStep `json:"next_step,omitempty"`

    // Indicates whether user set optional parameter NewNextStep
    IsNewNextStepSet bool

	// The output of the dialog node. For more information about how to specify dialog node output, see the [documentation](https://console.bluemix.net/docs/services/conversation/dialog-overview.html#complex).
	NewOutput DialogNodeOutput `json:"output,omitempty"`

    // Indicates whether user set optional parameter NewOutput
    IsNewOutputSet bool

	// Whether this top-level dialog node can be digressed into.
	NewDigressIn string `json:"digress_in,omitempty"`

    // Indicates whether user set optional parameter NewDigressIn
    IsNewDigressInSet bool

	// The ID of the parent dialog node.
	NewParent string `json:"parent,omitempty"`

    // Indicates whether user set optional parameter NewParent
    IsNewParentSet bool

	// The dialog node ID. This string must conform to the following restrictions: - It can contain only Unicode alphanumeric, space, underscore, hyphen, and dot characters. - It must be no longer than 1024 characters.
	NewDialogNode string `json:"dialog_node,omitempty"`

    // Indicates whether user set optional parameter NewDialogNode
    IsNewDialogNodeSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewUpdateDialogNodeOptions : Instantiate UpdateDialogNodeOptions
func NewUpdateDialogNodeOptions(workspaceID string, dialogNode string) *UpdateDialogNodeOptions {
    return &UpdateDialogNodeOptions{
        WorkspaceID: workspaceID,
        DialogNode: dialogNode,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *UpdateDialogNodeOptions) SetWorkspaceID(param string) *UpdateDialogNodeOptions {
    options.WorkspaceID = param
    return options
}

// SetDialogNode : Allow user to set DialogNode
func (options *UpdateDialogNodeOptions) SetDialogNode(param string) *UpdateDialogNodeOptions {
    options.DialogNode = param
    return options
}

// SetNodeType : Allow user to set NodeType
func (options *UpdateDialogNodeOptions) SetNodeType(param string) *UpdateDialogNodeOptions {
    options.NodeType = param
    options.IsNodeTypeSet = true
    return options
}

// SetNewConditions : Allow user to set NewConditions
func (options *UpdateDialogNodeOptions) SetNewConditions(param string) *UpdateDialogNodeOptions {
    options.NewConditions = param
    options.IsNewConditionsSet = true
    return options
}

// SetNewActions : Allow user to set NewActions
func (options *UpdateDialogNodeOptions) SetNewActions(param []DialogNodeAction) *UpdateDialogNodeOptions {
    options.NewActions = param
    options.IsNewActionsSet = true
    return options
}

// SetNewPreviousSibling : Allow user to set NewPreviousSibling
func (options *UpdateDialogNodeOptions) SetNewPreviousSibling(param string) *UpdateDialogNodeOptions {
    options.NewPreviousSibling = param
    options.IsNewPreviousSiblingSet = true
    return options
}

// SetNewContext : Allow user to set NewContext
func (options *UpdateDialogNodeOptions) SetNewContext(param interface{}) *UpdateDialogNodeOptions {
    options.NewContext = param
    options.IsNewContextSet = true
    return options
}

// SetNewVariable : Allow user to set NewVariable
func (options *UpdateDialogNodeOptions) SetNewVariable(param string) *UpdateDialogNodeOptions {
    options.NewVariable = param
    options.IsNewVariableSet = true
    return options
}

// SetNewUserLabel : Allow user to set NewUserLabel
func (options *UpdateDialogNodeOptions) SetNewUserLabel(param string) *UpdateDialogNodeOptions {
    options.NewUserLabel = param
    options.IsNewUserLabelSet = true
    return options
}

// SetNewMetadata : Allow user to set NewMetadata
func (options *UpdateDialogNodeOptions) SetNewMetadata(param interface{}) *UpdateDialogNodeOptions {
    options.NewMetadata = param
    options.IsNewMetadataSet = true
    return options
}

// SetNewTitle : Allow user to set NewTitle
func (options *UpdateDialogNodeOptions) SetNewTitle(param string) *UpdateDialogNodeOptions {
    options.NewTitle = param
    options.IsNewTitleSet = true
    return options
}

// SetNewDescription : Allow user to set NewDescription
func (options *UpdateDialogNodeOptions) SetNewDescription(param string) *UpdateDialogNodeOptions {
    options.NewDescription = param
    options.IsNewDescriptionSet = true
    return options
}

// SetNewDigressOut : Allow user to set NewDigressOut
func (options *UpdateDialogNodeOptions) SetNewDigressOut(param string) *UpdateDialogNodeOptions {
    options.NewDigressOut = param
    options.IsNewDigressOutSet = true
    return options
}

// SetNewEventName : Allow user to set NewEventName
func (options *UpdateDialogNodeOptions) SetNewEventName(param string) *UpdateDialogNodeOptions {
    options.NewEventName = param
    options.IsNewEventNameSet = true
    return options
}

// SetNewDigressOutSlots : Allow user to set NewDigressOutSlots
func (options *UpdateDialogNodeOptions) SetNewDigressOutSlots(param string) *UpdateDialogNodeOptions {
    options.NewDigressOutSlots = param
    options.IsNewDigressOutSlotsSet = true
    return options
}

// SetNewNextStep : Allow user to set NewNextStep
func (options *UpdateDialogNodeOptions) SetNewNextStep(param DialogNodeNextStep) *UpdateDialogNodeOptions {
    options.NewNextStep = param
    options.IsNewNextStepSet = true
    return options
}

// SetNewOutput : Allow user to set NewOutput
func (options *UpdateDialogNodeOptions) SetNewOutput(param DialogNodeOutput) *UpdateDialogNodeOptions {
    options.NewOutput = param
    options.IsNewOutputSet = true
    return options
}

// SetNewDigressIn : Allow user to set NewDigressIn
func (options *UpdateDialogNodeOptions) SetNewDigressIn(param string) *UpdateDialogNodeOptions {
    options.NewDigressIn = param
    options.IsNewDigressInSet = true
    return options
}

// SetNewParent : Allow user to set NewParent
func (options *UpdateDialogNodeOptions) SetNewParent(param string) *UpdateDialogNodeOptions {
    options.NewParent = param
    options.IsNewParentSet = true
    return options
}

// SetNewDialogNode : Allow user to set NewDialogNode
func (options *UpdateDialogNodeOptions) SetNewDialogNode(param string) *UpdateDialogNodeOptions {
    options.NewDialogNode = param
    options.IsNewDialogNodeSet = true
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
	WorkspaceID string `json:"workspace_id"`

	// The name of the entity.
	Entity string `json:"entity"`

	// Whether to use fuzzy matching for the entity.
	NewFuzzyMatch bool `json:"fuzzy_match,omitempty"`

    // Indicates whether user set optional parameter NewFuzzyMatch
    IsNewFuzzyMatchSet bool

	// The name of the entity. This string must conform to the following restrictions: - It can contain only Unicode alphanumeric, underscore, and hyphen characters. - It cannot begin with the reserved prefix `sys-`. - It must be no longer than 64 characters.
	NewEntity string `json:"entity,omitempty"`

    // Indicates whether user set optional parameter NewEntity
    IsNewEntitySet bool

	// Any metadata related to the entity.
	NewMetadata interface{} `json:"metadata,omitempty"`

    // Indicates whether user set optional parameter NewMetadata
    IsNewMetadataSet bool

	// An array of entity values.
	NewValues []CreateValue `json:"values,omitempty"`

    // Indicates whether user set optional parameter NewValues
    IsNewValuesSet bool

	// The description of the entity. This string cannot contain carriage return, newline, or tab characters, and it must be no longer than 128 characters.
	NewDescription string `json:"description,omitempty"`

    // Indicates whether user set optional parameter NewDescription
    IsNewDescriptionSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewUpdateEntityOptions : Instantiate UpdateEntityOptions
func NewUpdateEntityOptions(workspaceID string, entity string) *UpdateEntityOptions {
    return &UpdateEntityOptions{
        WorkspaceID: workspaceID,
        Entity: entity,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *UpdateEntityOptions) SetWorkspaceID(param string) *UpdateEntityOptions {
    options.WorkspaceID = param
    return options
}

// SetEntity : Allow user to set Entity
func (options *UpdateEntityOptions) SetEntity(param string) *UpdateEntityOptions {
    options.Entity = param
    return options
}

// SetNewFuzzyMatch : Allow user to set NewFuzzyMatch
func (options *UpdateEntityOptions) SetNewFuzzyMatch(param bool) *UpdateEntityOptions {
    options.NewFuzzyMatch = param
    options.IsNewFuzzyMatchSet = true
    return options
}

// SetNewEntity : Allow user to set NewEntity
func (options *UpdateEntityOptions) SetNewEntity(param string) *UpdateEntityOptions {
    options.NewEntity = param
    options.IsNewEntitySet = true
    return options
}

// SetNewMetadata : Allow user to set NewMetadata
func (options *UpdateEntityOptions) SetNewMetadata(param interface{}) *UpdateEntityOptions {
    options.NewMetadata = param
    options.IsNewMetadataSet = true
    return options
}

// SetNewValues : Allow user to set NewValues
func (options *UpdateEntityOptions) SetNewValues(param []CreateValue) *UpdateEntityOptions {
    options.NewValues = param
    options.IsNewValuesSet = true
    return options
}

// SetNewDescription : Allow user to set NewDescription
func (options *UpdateEntityOptions) SetNewDescription(param string) *UpdateEntityOptions {
    options.NewDescription = param
    options.IsNewDescriptionSet = true
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
	WorkspaceID string `json:"workspace_id"`

	// The intent name.
	Intent string `json:"intent"`

	// The text of the user input example.
	Text string `json:"text"`

	// The text of the user input example. This string must conform to the following restrictions: - It cannot contain carriage return, newline, or tab characters. - It cannot consist of only whitespace characters. - It must be no longer than 1024 characters.
	NewText string `json:"text,omitempty"`

    // Indicates whether user set optional parameter NewText
    IsNewTextSet bool

	// An array of contextual entity mentions.
	NewMentions []Mentions `json:"mentions,omitempty"`

    // Indicates whether user set optional parameter NewMentions
    IsNewMentionsSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewUpdateExampleOptions : Instantiate UpdateExampleOptions
func NewUpdateExampleOptions(workspaceID string, intent string, text string) *UpdateExampleOptions {
    return &UpdateExampleOptions{
        WorkspaceID: workspaceID,
        Intent: intent,
        Text: text,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *UpdateExampleOptions) SetWorkspaceID(param string) *UpdateExampleOptions {
    options.WorkspaceID = param
    return options
}

// SetIntent : Allow user to set Intent
func (options *UpdateExampleOptions) SetIntent(param string) *UpdateExampleOptions {
    options.Intent = param
    return options
}

// SetText : Allow user to set Text
func (options *UpdateExampleOptions) SetText(param string) *UpdateExampleOptions {
    options.Text = param
    return options
}

// SetNewText : Allow user to set NewText
func (options *UpdateExampleOptions) SetNewText(param string) *UpdateExampleOptions {
    options.NewText = param
    options.IsNewTextSet = true
    return options
}

// SetNewMentions : Allow user to set NewMentions
func (options *UpdateExampleOptions) SetNewMentions(param []Mentions) *UpdateExampleOptions {
    options.NewMentions = param
    options.IsNewMentionsSet = true
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
	WorkspaceID string `json:"workspace_id"`

	// The intent name.
	Intent string `json:"intent"`

	// The name of the intent. This string must conform to the following restrictions: - It can contain only Unicode alphanumeric, underscore, hyphen, and dot characters. - It cannot begin with the reserved prefix `sys-`. - It must be no longer than 128 characters.
	NewIntent string `json:"intent,omitempty"`

    // Indicates whether user set optional parameter NewIntent
    IsNewIntentSet bool

	// An array of user input examples for the intent.
	NewExamples []CreateExample `json:"examples,omitempty"`

    // Indicates whether user set optional parameter NewExamples
    IsNewExamplesSet bool

	// The description of the intent.
	NewDescription string `json:"description,omitempty"`

    // Indicates whether user set optional parameter NewDescription
    IsNewDescriptionSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewUpdateIntentOptions : Instantiate UpdateIntentOptions
func NewUpdateIntentOptions(workspaceID string, intent string) *UpdateIntentOptions {
    return &UpdateIntentOptions{
        WorkspaceID: workspaceID,
        Intent: intent,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *UpdateIntentOptions) SetWorkspaceID(param string) *UpdateIntentOptions {
    options.WorkspaceID = param
    return options
}

// SetIntent : Allow user to set Intent
func (options *UpdateIntentOptions) SetIntent(param string) *UpdateIntentOptions {
    options.Intent = param
    return options
}

// SetNewIntent : Allow user to set NewIntent
func (options *UpdateIntentOptions) SetNewIntent(param string) *UpdateIntentOptions {
    options.NewIntent = param
    options.IsNewIntentSet = true
    return options
}

// SetNewExamples : Allow user to set NewExamples
func (options *UpdateIntentOptions) SetNewExamples(param []CreateExample) *UpdateIntentOptions {
    options.NewExamples = param
    options.IsNewExamplesSet = true
    return options
}

// SetNewDescription : Allow user to set NewDescription
func (options *UpdateIntentOptions) SetNewDescription(param string) *UpdateIntentOptions {
    options.NewDescription = param
    options.IsNewDescriptionSet = true
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
	WorkspaceID string `json:"workspace_id"`

	// The name of the entity.
	Entity string `json:"entity"`

	// The text of the entity value.
	Value string `json:"value"`

	// The text of the synonym.
	Synonym string `json:"synonym"`

	// The text of the synonym. This string must conform to the following restrictions: - It cannot contain carriage return, newline, or tab characters. - It cannot consist of only whitespace characters. - It must be no longer than 64 characters.
	NewSynonym string `json:"synonym,omitempty"`

    // Indicates whether user set optional parameter NewSynonym
    IsNewSynonymSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewUpdateSynonymOptions : Instantiate UpdateSynonymOptions
func NewUpdateSynonymOptions(workspaceID string, entity string, value string, synonym string) *UpdateSynonymOptions {
    return &UpdateSynonymOptions{
        WorkspaceID: workspaceID,
        Entity: entity,
        Value: value,
        Synonym: synonym,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *UpdateSynonymOptions) SetWorkspaceID(param string) *UpdateSynonymOptions {
    options.WorkspaceID = param
    return options
}

// SetEntity : Allow user to set Entity
func (options *UpdateSynonymOptions) SetEntity(param string) *UpdateSynonymOptions {
    options.Entity = param
    return options
}

// SetValue : Allow user to set Value
func (options *UpdateSynonymOptions) SetValue(param string) *UpdateSynonymOptions {
    options.Value = param
    return options
}

// SetSynonym : Allow user to set Synonym
func (options *UpdateSynonymOptions) SetSynonym(param string) *UpdateSynonymOptions {
    options.Synonym = param
    return options
}

// SetNewSynonym : Allow user to set NewSynonym
func (options *UpdateSynonymOptions) SetNewSynonym(param string) *UpdateSynonymOptions {
    options.NewSynonym = param
    options.IsNewSynonymSet = true
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
	WorkspaceID string `json:"workspace_id"`

	// The name of the entity.
	Entity string `json:"entity"`

	// The text of the entity value.
	Value string `json:"value"`

	// Specifies the type of value.
	ValueType string `json:"type,omitempty"`

    // Indicates whether user set optional parameter ValueType
    IsValueTypeSet bool

	// An array of synonyms for the entity value. You can provide either synonyms or patterns (as indicated by **type**), but not both. A synonym must conform to the following resrictions: - It cannot contain carriage return, newline, or tab characters. - It cannot consist of only whitespace characters. - It must be no longer than 64 characters.
	NewSynonyms []string `json:"synonyms,omitempty"`

    // Indicates whether user set optional parameter NewSynonyms
    IsNewSynonymsSet bool

	// Any metadata related to the entity value.
	NewMetadata interface{} `json:"metadata,omitempty"`

    // Indicates whether user set optional parameter NewMetadata
    IsNewMetadataSet bool

	// The text of the entity value. This string must conform to the following restrictions: - It cannot contain carriage return, newline, or tab characters. - It cannot consist of only whitespace characters. - It must be no longer than 64 characters.
	NewValue string `json:"value,omitempty"`

    // Indicates whether user set optional parameter NewValue
    IsNewValueSet bool

	// An array of patterns for the entity value. You can provide either synonyms or patterns (as indicated by **type**), but not both. A pattern is a regular expression no longer than 512 characters. For more information about how to specify a pattern, see the [documentation](https://console.bluemix.net/docs/services/conversation/entities.html#creating-entities).
	NewPatterns []string `json:"patterns,omitempty"`

    // Indicates whether user set optional parameter NewPatterns
    IsNewPatternsSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewUpdateValueOptions : Instantiate UpdateValueOptions
func NewUpdateValueOptions(workspaceID string, entity string, value string) *UpdateValueOptions {
    return &UpdateValueOptions{
        WorkspaceID: workspaceID,
        Entity: entity,
        Value: value,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *UpdateValueOptions) SetWorkspaceID(param string) *UpdateValueOptions {
    options.WorkspaceID = param
    return options
}

// SetEntity : Allow user to set Entity
func (options *UpdateValueOptions) SetEntity(param string) *UpdateValueOptions {
    options.Entity = param
    return options
}

// SetValue : Allow user to set Value
func (options *UpdateValueOptions) SetValue(param string) *UpdateValueOptions {
    options.Value = param
    return options
}

// SetValueType : Allow user to set ValueType
func (options *UpdateValueOptions) SetValueType(param string) *UpdateValueOptions {
    options.ValueType = param
    options.IsValueTypeSet = true
    return options
}

// SetNewSynonyms : Allow user to set NewSynonyms
func (options *UpdateValueOptions) SetNewSynonyms(param []string) *UpdateValueOptions {
    options.NewSynonyms = param
    options.IsNewSynonymsSet = true
    return options
}

// SetNewMetadata : Allow user to set NewMetadata
func (options *UpdateValueOptions) SetNewMetadata(param interface{}) *UpdateValueOptions {
    options.NewMetadata = param
    options.IsNewMetadataSet = true
    return options
}

// SetNewValue : Allow user to set NewValue
func (options *UpdateValueOptions) SetNewValue(param string) *UpdateValueOptions {
    options.NewValue = param
    options.IsNewValueSet = true
    return options
}

// SetNewPatterns : Allow user to set NewPatterns
func (options *UpdateValueOptions) SetNewPatterns(param []string) *UpdateValueOptions {
    options.NewPatterns = param
    options.IsNewPatternsSet = true
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
	WorkspaceID string `json:"workspace_id"`

	// The name of the workspace. This string cannot contain carriage return, newline, or tab characters, and it must be no longer than 64 characters.
	Name string `json:"name,omitempty"`

    // Indicates whether user set optional parameter Name
    IsNameSet bool

	// The description of the workspace. This string cannot contain carriage return, newline, or tab characters, and it must be no longer than 128 characters.
	Description string `json:"description,omitempty"`

    // Indicates whether user set optional parameter Description
    IsDescriptionSet bool

	// The language of the workspace.
	Language string `json:"language,omitempty"`

    // Indicates whether user set optional parameter Language
    IsLanguageSet bool

	// An array of objects defining the intents for the workspace.
	Intents []CreateIntent `json:"intents,omitempty"`

    // Indicates whether user set optional parameter Intents
    IsIntentsSet bool

	// An array of objects defining the entities for the workspace.
	Entities []CreateEntity `json:"entities,omitempty"`

    // Indicates whether user set optional parameter Entities
    IsEntitiesSet bool

	// An array of objects defining the nodes in the workspace dialog.
	DialogNodes []CreateDialogNode `json:"dialog_nodes,omitempty"`

    // Indicates whether user set optional parameter DialogNodes
    IsDialogNodesSet bool

	// An array of objects defining input examples that have been marked as irrelevant input.
	Counterexamples []CreateCounterexample `json:"counterexamples,omitempty"`

    // Indicates whether user set optional parameter Counterexamples
    IsCounterexamplesSet bool

	// Any metadata related to the workspace.
	Metadata interface{} `json:"metadata,omitempty"`

    // Indicates whether user set optional parameter Metadata
    IsMetadataSet bool

	// Whether training data from the workspace can be used by IBM for general service improvements. `true` indicates that workspace training data is not to be used.
	LearningOptOut bool `json:"learning_opt_out,omitempty"`

    // Indicates whether user set optional parameter LearningOptOut
    IsLearningOptOutSet bool

	// Global settings for the workspace.
	SystemSettings WorkspaceSystemSettings `json:"system_settings,omitempty"`

    // Indicates whether user set optional parameter SystemSettings
    IsSystemSettingsSet bool

	// Whether the new data is to be appended to the existing data in the workspace. If **append**=`false`, elements included in the new data completely replace the corresponding existing elements, including all subelements. For example, if the new data includes **entities** and **append**=`false`, all existing entities in the workspace are discarded and replaced with the new entities. If **append**=`true`, existing elements are preserved, and the new elements are added. If any elements in the new data collide with existing elements, the update request fails.
	AppendVar bool `json:"append,omitempty"`

    // Indicates whether user set optional parameter AppendVar
    IsAppendVarSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewUpdateWorkspaceOptions : Instantiate UpdateWorkspaceOptions
func NewUpdateWorkspaceOptions(workspaceID string) *UpdateWorkspaceOptions {
    return &UpdateWorkspaceOptions{
        WorkspaceID: workspaceID,
    }
}

// SetWorkspaceID : Allow user to set WorkspaceID
func (options *UpdateWorkspaceOptions) SetWorkspaceID(param string) *UpdateWorkspaceOptions {
    options.WorkspaceID = param
    return options
}

// SetName : Allow user to set Name
func (options *UpdateWorkspaceOptions) SetName(param string) *UpdateWorkspaceOptions {
    options.Name = param
    options.IsNameSet = true
    return options
}

// SetDescription : Allow user to set Description
func (options *UpdateWorkspaceOptions) SetDescription(param string) *UpdateWorkspaceOptions {
    options.Description = param
    options.IsDescriptionSet = true
    return options
}

// SetLanguage : Allow user to set Language
func (options *UpdateWorkspaceOptions) SetLanguage(param string) *UpdateWorkspaceOptions {
    options.Language = param
    options.IsLanguageSet = true
    return options
}

// SetIntents : Allow user to set Intents
func (options *UpdateWorkspaceOptions) SetIntents(param []CreateIntent) *UpdateWorkspaceOptions {
    options.Intents = param
    options.IsIntentsSet = true
    return options
}

// SetEntities : Allow user to set Entities
func (options *UpdateWorkspaceOptions) SetEntities(param []CreateEntity) *UpdateWorkspaceOptions {
    options.Entities = param
    options.IsEntitiesSet = true
    return options
}

// SetDialogNodes : Allow user to set DialogNodes
func (options *UpdateWorkspaceOptions) SetDialogNodes(param []CreateDialogNode) *UpdateWorkspaceOptions {
    options.DialogNodes = param
    options.IsDialogNodesSet = true
    return options
}

// SetCounterexamples : Allow user to set Counterexamples
func (options *UpdateWorkspaceOptions) SetCounterexamples(param []CreateCounterexample) *UpdateWorkspaceOptions {
    options.Counterexamples = param
    options.IsCounterexamplesSet = true
    return options
}

// SetMetadata : Allow user to set Metadata
func (options *UpdateWorkspaceOptions) SetMetadata(param interface{}) *UpdateWorkspaceOptions {
    options.Metadata = param
    options.IsMetadataSet = true
    return options
}

// SetLearningOptOut : Allow user to set LearningOptOut
func (options *UpdateWorkspaceOptions) SetLearningOptOut(param bool) *UpdateWorkspaceOptions {
    options.LearningOptOut = param
    options.IsLearningOptOutSet = true
    return options
}

// SetSystemSettings : Allow user to set SystemSettings
func (options *UpdateWorkspaceOptions) SetSystemSettings(param WorkspaceSystemSettings) *UpdateWorkspaceOptions {
    options.SystemSettings = param
    options.IsSystemSettingsSet = true
    return options
}

// SetAppendVar : Allow user to set AppendVar
func (options *UpdateWorkspaceOptions) SetAppendVar(param bool) *UpdateWorkspaceOptions {
    options.AppendVar = param
    options.IsAppendVarSet = true
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
	ValueText string `json:"value"`

	// Any metadata related to the entity value.
	Metadata interface{} `json:"metadata,omitempty"`

	// The timestamp for creation of the entity value.
	Created strfmt.DateTime `json:"created,omitempty"`

	// The timestamp for the last update to the entity value.
	Updated strfmt.DateTime `json:"updated,omitempty"`

	// An array containing any synonyms for the entity value.
	Synonyms []string `json:"synonyms,omitempty"`

	// An array containing any patterns for the entity value.
	Patterns []string `json:"patterns,omitempty"`

	// Specifies the type of value.
	ValueType string `json:"type"`
}

// ValueCollection : ValueCollection struct
type ValueCollection struct {

	// An array of entity values.
	Values []ValueExport `json:"values"`

	// An object defining the pagination data for the returned objects.
	Pagination Pagination `json:"pagination"`
}

// ValueExport : ValueExport struct
type ValueExport struct {

	// The text of the entity value.
	ValueText string `json:"value"`

	// Any metadata related to the entity value.
	Metadata interface{} `json:"metadata,omitempty"`

	// The timestamp for creation of the entity value.
	Created strfmt.DateTime `json:"created,omitempty"`

	// The timestamp for the last update to the entity value.
	Updated strfmt.DateTime `json:"updated,omitempty"`

	// An array containing any synonyms for the entity value.
	Synonyms []string `json:"synonyms,omitempty"`

	// An array containing any patterns for the entity value.
	Patterns []string `json:"patterns,omitempty"`

	// Specifies the type of value.
	ValueType string `json:"type"`
}

// Workspace : Workspace struct
type Workspace struct {

	// The name of the workspace.
	Name string `json:"name"`

	// The language of the workspace.
	Language string `json:"language"`

	// The timestamp for creation of the workspace.
	Created strfmt.DateTime `json:"created,omitempty"`

	// The timestamp for the last update to the workspace.
	Updated strfmt.DateTime `json:"updated,omitempty"`

	// The workspace ID.
	WorkspaceID string `json:"workspace_id"`

	// The description of the workspace.
	Description string `json:"description,omitempty"`

	// Any metadata related to the workspace.
	Metadata interface{} `json:"metadata,omitempty"`

	// Whether training data from the workspace (including artifacts such as intents and entities) can be used by IBM for general service improvements. `true` indicates that workspace training data is not to be used.
	LearningOptOut bool `json:"learning_opt_out,omitempty"`

	// Global settings for the workspace.
	SystemSettings WorkspaceSystemSettings `json:"system_settings,omitempty"`
}

// WorkspaceCollection : WorkspaceCollection struct
type WorkspaceCollection struct {

	// An array of objects describing the workspaces associated with the service instance.
	Workspaces []Workspace `json:"workspaces"`

	// An object defining the pagination data for the returned objects.
	Pagination Pagination `json:"pagination"`
}

// WorkspaceExport : WorkspaceExport struct
type WorkspaceExport struct {

	// The name of the workspace.
	Name string `json:"name"`

	// The description of the workspace.
	Description string `json:"description"`

	// The language of the workspace.
	Language string `json:"language"`

	// Any metadata that is required by the workspace.
	Metadata interface{} `json:"metadata"`

	// The timestamp for creation of the workspace.
	Created strfmt.DateTime `json:"created,omitempty"`

	// The timestamp for the last update to the workspace.
	Updated strfmt.DateTime `json:"updated,omitempty"`

	// The workspace ID.
	WorkspaceID string `json:"workspace_id"`

	// The current status of the workspace.
	Status string `json:"status"`

	// Whether training data from the workspace can be used by IBM for general service improvements. `true` indicates that workspace training data is not to be used.
	LearningOptOut bool `json:"learning_opt_out"`

	// Global settings for the workspace.
	SystemSettings WorkspaceSystemSettings `json:"system_settings,omitempty"`

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
	Tooling WorkspaceSystemSettingsTooling `json:"tooling,omitempty"`

	// Workspace settings related to the disambiguation feature. **Note:** This feature is available only to Premium users.
	Disambiguation WorkspaceSystemSettingsDisambiguation `json:"disambiguation,omitempty"`

	// For internal use only.
	HumanAgentAssist interface{} `json:"human_agent_assist,omitempty"`
}

// WorkspaceSystemSettingsDisambiguation : WorkspaceSystemSettingsDisambiguation struct
type WorkspaceSystemSettingsDisambiguation struct {

	// The text of the introductory prompt that accompanies disambiguation options presented to the user.
	Prompt string `json:"prompt,omitempty"`

	// The user-facing label for the option users can select if none of the suggested options is correct. If no value is specified for this property, this option does not appear.
	NoneOfTheAbovePrompt string `json:"none_of_the_above_prompt,omitempty"`

	// Whether the disambiguation feature is enabled for the workspace.
	Enabled bool `json:"enabled,omitempty"`

	// The sensitivity of the disambiguation feature to intent detection conflicts. Set to **high** if you want the disambiguation feature to be triggered more often. This can be useful for testing or demonstration purposes.
	Sensitivity string `json:"sensitivity,omitempty"`
}

// WorkspaceSystemSettingsTooling : WorkspaceSystemSettingsTooling struct
type WorkspaceSystemSettingsTooling struct {

	// Whether the dialog JSON editor displays text responses within the `output.generic` object.
	StoreGenericResponses bool `json:"store_generic_responses,omitempty"`
}
