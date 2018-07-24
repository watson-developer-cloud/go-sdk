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
    "bytes"
    "fmt"
    "github.com/go-openapi/strfmt"
    "strings"
    req "github.com/parnurzeal/gorequest"
    watson "golang-sdk"
)

// AssistantV1 : The AssistantV1 service
type AssistantV1 struct {
	client *watson.Client
}

// NewAssistantV1 : Instantiate AssistantV1
func NewAssistantV1(creds watson.Credentials) (*AssistantV1, error) {
    if creds.ServiceURL == "" {
        creds.ServiceURL = "https://gateway.watsonplatform.net/assistant/api"
    }

	client, clientErr := watson.NewClient(creds, "conversation")

	if clientErr != nil {
		return nil, clientErr
	}

	return &AssistantV1{ client: client }, nil
}

// Message : Get response to user input
func (assistant *AssistantV1) Message(workspaceID string, body *MessageRequest, nodesVisitedDetails bool) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/message"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("nodes_visited_details=" + fmt.Sprint(nodesVisitedDetails))
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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) CreateWorkspace(body *CreateWorkspace) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) DeleteWorkspace(workspaceID string) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    request := req.New().Delete(creds.ServiceURL + path)

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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) GetWorkspace(workspaceID string, export bool, includeAudit bool) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("export=" + fmt.Sprint(export))
    request.Query("include_audit=" + fmt.Sprint(includeAudit))

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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) ListWorkspaces(pageLimit int64, includeCount bool, sort string, cursor string, includeAudit bool) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("page_limit=" + fmt.Sprint(pageLimit))
    request.Query("include_count=" + fmt.Sprint(includeCount))
    request.Query("sort=" + fmt.Sprint(sort))
    request.Query("cursor=" + fmt.Sprint(cursor))
    request.Query("include_audit=" + fmt.Sprint(includeAudit))

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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) UpdateWorkspace(workspaceID string, body *UpdateWorkspace, appendVar bool) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("append=" + fmt.Sprint(appendVar))
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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) CreateIntent(workspaceID string, body *CreateIntent) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/intents"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) DeleteIntent(workspaceID string, intent string) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/intents/{intent}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    path = strings.Replace(path, "{intent}", intent, 1)
    request := req.New().Delete(creds.ServiceURL + path)

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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) GetIntent(workspaceID string, intent string, export bool, includeAudit bool) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/intents/{intent}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    path = strings.Replace(path, "{intent}", intent, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("export=" + fmt.Sprint(export))
    request.Query("include_audit=" + fmt.Sprint(includeAudit))

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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) ListIntents(workspaceID string, export bool, pageLimit int64, includeCount bool, sort string, cursor string, includeAudit bool) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/intents"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("export=" + fmt.Sprint(export))
    request.Query("page_limit=" + fmt.Sprint(pageLimit))
    request.Query("include_count=" + fmt.Sprint(includeCount))
    request.Query("sort=" + fmt.Sprint(sort))
    request.Query("cursor=" + fmt.Sprint(cursor))
    request.Query("include_audit=" + fmt.Sprint(includeAudit))

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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) UpdateIntent(workspaceID string, intent string, body *UpdateIntent) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/intents/{intent}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    path = strings.Replace(path, "{intent}", intent, 1)
    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) CreateExample(workspaceID string, intent string, body *CreateExample) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/intents/{intent}/examples"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    path = strings.Replace(path, "{intent}", intent, 1)
    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) DeleteExample(workspaceID string, intent string, text string) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/intents/{intent}/examples/{text}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    path = strings.Replace(path, "{intent}", intent, 1)
    path = strings.Replace(path, "{text}", text, 1)
    request := req.New().Delete(creds.ServiceURL + path)

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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) GetExample(workspaceID string, intent string, text string, includeAudit bool) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/intents/{intent}/examples/{text}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    path = strings.Replace(path, "{intent}", intent, 1)
    path = strings.Replace(path, "{text}", text, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("include_audit=" + fmt.Sprint(includeAudit))

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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) ListExamples(workspaceID string, intent string, pageLimit int64, includeCount bool, sort string, cursor string, includeAudit bool) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/intents/{intent}/examples"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    path = strings.Replace(path, "{intent}", intent, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("page_limit=" + fmt.Sprint(pageLimit))
    request.Query("include_count=" + fmt.Sprint(includeCount))
    request.Query("sort=" + fmt.Sprint(sort))
    request.Query("cursor=" + fmt.Sprint(cursor))
    request.Query("include_audit=" + fmt.Sprint(includeAudit))

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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) UpdateExample(workspaceID string, intent string, text string, body *UpdateExample) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/intents/{intent}/examples/{text}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    path = strings.Replace(path, "{intent}", intent, 1)
    path = strings.Replace(path, "{text}", text, 1)
    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) CreateCounterexample(workspaceID string, body *CreateCounterexample) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/counterexamples"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) DeleteCounterexample(workspaceID string, text string) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/counterexamples/{text}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    path = strings.Replace(path, "{text}", text, 1)
    request := req.New().Delete(creds.ServiceURL + path)

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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) GetCounterexample(workspaceID string, text string, includeAudit bool) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/counterexamples/{text}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    path = strings.Replace(path, "{text}", text, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("include_audit=" + fmt.Sprint(includeAudit))

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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) ListCounterexamples(workspaceID string, pageLimit int64, includeCount bool, sort string, cursor string, includeAudit bool) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/counterexamples"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("page_limit=" + fmt.Sprint(pageLimit))
    request.Query("include_count=" + fmt.Sprint(includeCount))
    request.Query("sort=" + fmt.Sprint(sort))
    request.Query("cursor=" + fmt.Sprint(cursor))
    request.Query("include_audit=" + fmt.Sprint(includeAudit))

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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) UpdateCounterexample(workspaceID string, text string, body *UpdateCounterexample) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/counterexamples/{text}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    path = strings.Replace(path, "{text}", text, 1)
    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) CreateEntity(workspaceID string, body *CreateEntity) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/entities"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) DeleteEntity(workspaceID string, entity string) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/entities/{entity}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    path = strings.Replace(path, "{entity}", entity, 1)
    request := req.New().Delete(creds.ServiceURL + path)

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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) GetEntity(workspaceID string, entity string, export bool, includeAudit bool) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/entities/{entity}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    path = strings.Replace(path, "{entity}", entity, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("export=" + fmt.Sprint(export))
    request.Query("include_audit=" + fmt.Sprint(includeAudit))

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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) ListEntities(workspaceID string, export bool, pageLimit int64, includeCount bool, sort string, cursor string, includeAudit bool) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/entities"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("export=" + fmt.Sprint(export))
    request.Query("page_limit=" + fmt.Sprint(pageLimit))
    request.Query("include_count=" + fmt.Sprint(includeCount))
    request.Query("sort=" + fmt.Sprint(sort))
    request.Query("cursor=" + fmt.Sprint(cursor))
    request.Query("include_audit=" + fmt.Sprint(includeAudit))

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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) UpdateEntity(workspaceID string, entity string, body *UpdateEntity) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/entities/{entity}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    path = strings.Replace(path, "{entity}", entity, 1)
    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) ListEntityMentions(workspaceID string, entity string, export bool, includeAudit bool) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/entities/{entity}/mentions"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    path = strings.Replace(path, "{entity}", entity, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("export=" + fmt.Sprint(export))
    request.Query("include_audit=" + fmt.Sprint(includeAudit))

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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) CreateValue(workspaceID string, entity string, body *CreateValue) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/entities/{entity}/values"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    path = strings.Replace(path, "{entity}", entity, 1)
    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) DeleteValue(workspaceID string, entity string, value string) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    path = strings.Replace(path, "{entity}", entity, 1)
    path = strings.Replace(path, "{value}", value, 1)
    request := req.New().Delete(creds.ServiceURL + path)

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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) GetValue(workspaceID string, entity string, value string, export bool, includeAudit bool) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    path = strings.Replace(path, "{entity}", entity, 1)
    path = strings.Replace(path, "{value}", value, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("export=" + fmt.Sprint(export))
    request.Query("include_audit=" + fmt.Sprint(includeAudit))

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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) ListValues(workspaceID string, entity string, export bool, pageLimit int64, includeCount bool, sort string, cursor string, includeAudit bool) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/entities/{entity}/values"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    path = strings.Replace(path, "{entity}", entity, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("export=" + fmt.Sprint(export))
    request.Query("page_limit=" + fmt.Sprint(pageLimit))
    request.Query("include_count=" + fmt.Sprint(includeCount))
    request.Query("sort=" + fmt.Sprint(sort))
    request.Query("cursor=" + fmt.Sprint(cursor))
    request.Query("include_audit=" + fmt.Sprint(includeAudit))

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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) UpdateValue(workspaceID string, entity string, value string, body *UpdateValue) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    path = strings.Replace(path, "{entity}", entity, 1)
    path = strings.Replace(path, "{value}", value, 1)
    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) CreateSynonym(workspaceID string, entity string, value string, body *CreateSynonym) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}/synonyms"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    path = strings.Replace(path, "{entity}", entity, 1)
    path = strings.Replace(path, "{value}", value, 1)
    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) DeleteSynonym(workspaceID string, entity string, value string, synonym string) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}/synonyms/{synonym}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    path = strings.Replace(path, "{entity}", entity, 1)
    path = strings.Replace(path, "{value}", value, 1)
    path = strings.Replace(path, "{synonym}", synonym, 1)
    request := req.New().Delete(creds.ServiceURL + path)

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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) GetSynonym(workspaceID string, entity string, value string, synonym string, includeAudit bool) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}/synonyms/{synonym}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    path = strings.Replace(path, "{entity}", entity, 1)
    path = strings.Replace(path, "{value}", value, 1)
    path = strings.Replace(path, "{synonym}", synonym, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("include_audit=" + fmt.Sprint(includeAudit))

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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) ListSynonyms(workspaceID string, entity string, value string, pageLimit int64, includeCount bool, sort string, cursor string, includeAudit bool) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}/synonyms"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    path = strings.Replace(path, "{entity}", entity, 1)
    path = strings.Replace(path, "{value}", value, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("page_limit=" + fmt.Sprint(pageLimit))
    request.Query("include_count=" + fmt.Sprint(includeCount))
    request.Query("sort=" + fmt.Sprint(sort))
    request.Query("cursor=" + fmt.Sprint(cursor))
    request.Query("include_audit=" + fmt.Sprint(includeAudit))

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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) UpdateSynonym(workspaceID string, entity string, value string, synonym string, body *UpdateSynonym) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/entities/{entity}/values/{value}/synonyms/{synonym}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    path = strings.Replace(path, "{entity}", entity, 1)
    path = strings.Replace(path, "{value}", value, 1)
    path = strings.Replace(path, "{synonym}", synonym, 1)
    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) CreateDialogNode(workspaceID string, body *CreateDialogNode) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/dialog_nodes"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) DeleteDialogNode(workspaceID string, dialogNode string) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/dialog_nodes/{dialog_node}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    path = strings.Replace(path, "{dialog_node}", dialogNode, 1)
    request := req.New().Delete(creds.ServiceURL + path)

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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) GetDialogNode(workspaceID string, dialogNode string, includeAudit bool) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/dialog_nodes/{dialog_node}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    path = strings.Replace(path, "{dialog_node}", dialogNode, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("include_audit=" + fmt.Sprint(includeAudit))

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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) ListDialogNodes(workspaceID string, pageLimit int64, includeCount bool, sort string, cursor string, includeAudit bool) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/dialog_nodes"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("page_limit=" + fmt.Sprint(pageLimit))
    request.Query("include_count=" + fmt.Sprint(includeCount))
    request.Query("sort=" + fmt.Sprint(sort))
    request.Query("cursor=" + fmt.Sprint(cursor))
    request.Query("include_audit=" + fmt.Sprint(includeAudit))

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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) UpdateDialogNode(workspaceID string, dialogNode string, body *UpdateDialogNode) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/dialog_nodes/{dialog_node}"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    path = strings.Replace(path, "{dialog_node}", dialogNode, 1)
    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) ListAllLogs(filter string, sort string, pageLimit int64, cursor string) (*watson.WatsonResponse, []error) {
    path := "/v1/logs"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("filter=" + fmt.Sprint(filter))
    request.Query("sort=" + fmt.Sprint(sort))
    request.Query("page_limit=" + fmt.Sprint(pageLimit))
    request.Query("cursor=" + fmt.Sprint(cursor))

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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) ListLogs(workspaceID string, sort string, filter string, pageLimit int64, cursor string) (*watson.WatsonResponse, []error) {
    path := "/v1/workspaces/{workspace_id}/logs"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    path = strings.Replace(path, "{workspace_id}", workspaceID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("sort=" + fmt.Sprint(sort))
    request.Query("filter=" + fmt.Sprint(filter))
    request.Query("page_limit=" + fmt.Sprint(pageLimit))
    request.Query("cursor=" + fmt.Sprint(cursor))

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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
func (assistant *AssistantV1) DeleteUserData(customerID string) (*watson.WatsonResponse, []error) {
    path := "/v1/user_data"
    creds := assistant.client.Creds
    useTM := assistant.client.UseTM
    tokenManager := assistant.client.TokenManager

    request := req.New().Delete(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("customer_id=" + fmt.Sprint(customerID))

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

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

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
	NodeType string `json:"node_type,omitempty"`

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

// CreateExample : CreateExample struct
type CreateExample struct {

	// The text of a user input example. This string must conform to the following restrictions: - It cannot contain carriage return, newline, or tab characters. - It cannot consist of only whitespace characters. - It must be no longer than 1024 characters.
	Text string `json:"text"`

	// An array of contextual entity mentions.
	Mentions []Mentions `json:"mentions,omitempty"`
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

// CreateSynonym : CreateSynonym struct
type CreateSynonym struct {

	// The text of the synonym. This string must conform to the following restrictions: - It cannot contain carriage return, newline, or tab characters. - It cannot consist of only whitespace characters. - It must be no longer than 64 characters.
	Synonym string `json:"synonym"`
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
	ValueType string `json:"value_type,omitempty"`
}

// CreateWorkspace : CreateWorkspace struct
type CreateWorkspace struct {

	// The name of the workspace. This string cannot contain carriage return, newline, or tab characters, and it must be no longer than 64 characters.
	Name string `json:"name,omitempty"`

	// The description of the workspace. This string cannot contain carriage return, newline, or tab characters, and it must be no longer than 128 characters.
	Description string `json:"description,omitempty"`

	// The language of the workspace.
	Language string `json:"language,omitempty"`

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

	// Whether training data from the workspace can be used by IBM for general service improvements. `true` indicates that workspace training data is not to be used.
	LearningOptOut bool `json:"learning_opt_out,omitempty"`

	// Global settings for the workspace.
	SystemSettings WorkspaceSystemSettings `json:"system_settings,omitempty"`
}

// DialogNode : DialogNode struct
type DialogNode struct {

	// The dialog node ID.
	DialogNodeID string `json:"dialog_node_id"`

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
	NodeType string `json:"node_type,omitempty"`

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
	ActionType string `json:"action_type,omitempty"`

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

	// The type of response returned by the dialog node. The specified response type must be supported by the client application or channel.
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
}

// Entity : Entity struct
type Entity struct {

	// The name of the entity.
	EntityName string `json:"entity_name"`

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
	EntityName string `json:"entity_name"`

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
	ExampleText string `json:"example_text"`

	// The name of the intent.
	IntentName string `json:"intent_name"`

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
	ExampleText string `json:"example_text"`

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

// InputData : The user input.
type InputData struct {

	// The text of the user input. This string cannot contain carriage return, newline, or tab characters, and it must be no longer than 2048 characters.
	Text string `json:"text"`
}

// Intent : Intent struct
type Intent struct {

	// The name of the intent.
	IntentName string `json:"intent_name"`

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
	IntentName string `json:"intent_name"`

	// The timestamp for creation of the intent.
	Created strfmt.DateTime `json:"created,omitempty"`

	// The timestamp for the last update to the intent.
	Updated strfmt.DateTime `json:"updated,omitempty"`

	// The description of the intent.
	Description string `json:"description,omitempty"`

	// An array of objects describing the user input examples for the intent.
	Examples []Example `json:"examples,omitempty"`
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

// OutputData : An output object that includes the response to the user, the nodes that were hit, and messages from the log.
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
	SynonymText string `json:"synonym_text"`

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

// UpdateCounterexample : UpdateCounterexample struct
type UpdateCounterexample struct {

	// The text of a user input counterexample.
	Text string `json:"text,omitempty"`
}

// UpdateDialogNode : UpdateDialogNode struct
type UpdateDialogNode struct {

	// The dialog node ID. This string must conform to the following restrictions: - It can contain only Unicode alphanumeric, space, underscore, hyphen, and dot characters. - It must be no longer than 1024 characters.
	DialogNode string `json:"dialog_node,omitempty"`

	// The description of the dialog node. This string cannot contain carriage return, newline, or tab characters, and it must be no longer than 128 characters.
	Description string `json:"description,omitempty"`

	// The condition that will trigger the dialog node. This string cannot contain carriage return, newline, or tab characters, and it must be no longer than 2048 characters.
	Conditions string `json:"conditions,omitempty"`

	// The ID of the parent dialog node.
	Parent string `json:"parent,omitempty"`

	// The ID of the previous sibling dialog node.
	PreviousSibling string `json:"previous_sibling,omitempty"`

	// The output of the dialog node. For more information about how to specify dialog node output, see the [documentation](https://console.bluemix.net/docs/services/conversation/dialog-overview.html#complex).
	Output DialogNodeOutput `json:"output,omitempty"`

	// The context for the dialog node.
	Context interface{} `json:"context,omitempty"`

	// The metadata for the dialog node.
	Metadata interface{} `json:"metadata,omitempty"`

	// The next step to be executed in dialog processing.
	NextStep DialogNodeNextStep `json:"next_step,omitempty"`

	// The alias used to identify the dialog node. This string must conform to the following restrictions: - It can contain only Unicode alphanumeric, space, underscore, hyphen, and dot characters. - It must be no longer than 64 characters.
	Title string `json:"title,omitempty"`

	// How the dialog node is processed.
	NodeType string `json:"node_type,omitempty"`

	// How an `event_handler` node is processed.
	EventName string `json:"event_name,omitempty"`

	// The location in the dialog context where output is stored.
	Variable string `json:"variable,omitempty"`

	// An array of objects describing any actions to be invoked by the dialog node.
	Actions []DialogNodeAction `json:"actions,omitempty"`

	// Whether this top-level dialog node can be digressed into.
	DigressIn string `json:"digress_in,omitempty"`

	// Whether this dialog node can be returned to after a digression.
	DigressOut string `json:"digress_out,omitempty"`

	// Whether the user can digress to top-level nodes while filling out slots.
	DigressOutSlots string `json:"digress_out_slots,omitempty"`

	// A label that can be displayed externally to describe the purpose of the node to users.
	UserLabel string `json:"user_label,omitempty"`
}

// UpdateEntity : UpdateEntity struct
type UpdateEntity struct {

	// The name of the entity. This string must conform to the following restrictions: - It can contain only Unicode alphanumeric, underscore, and hyphen characters. - It cannot begin with the reserved prefix `sys-`. - It must be no longer than 64 characters.
	Entity string `json:"entity,omitempty"`

	// The description of the entity. This string cannot contain carriage return, newline, or tab characters, and it must be no longer than 128 characters.
	Description string `json:"description,omitempty"`

	// Any metadata related to the entity.
	Metadata interface{} `json:"metadata,omitempty"`

	// Whether to use fuzzy matching for the entity.
	FuzzyMatch bool `json:"fuzzy_match,omitempty"`

	// An array of entity values.
	Values []CreateValue `json:"values,omitempty"`
}

// UpdateExample : UpdateExample struct
type UpdateExample struct {

	// The text of the user input example. This string must conform to the following restrictions: - It cannot contain carriage return, newline, or tab characters. - It cannot consist of only whitespace characters. - It must be no longer than 1024 characters.
	Text string `json:"text,omitempty"`

	// An array of contextual entity mentions.
	Mentions []Mentions `json:"mentions,omitempty"`
}

// UpdateIntent : UpdateIntent struct
type UpdateIntent struct {

	// The name of the intent. This string must conform to the following restrictions: - It can contain only Unicode alphanumeric, underscore, hyphen, and dot characters. - It cannot begin with the reserved prefix `sys-`. - It must be no longer than 128 characters.
	Intent string `json:"intent,omitempty"`

	// The description of the intent.
	Description string `json:"description,omitempty"`

	// An array of user input examples for the intent.
	Examples []CreateExample `json:"examples,omitempty"`
}

// UpdateSynonym : UpdateSynonym struct
type UpdateSynonym struct {

	// The text of the synonym. This string must conform to the following restrictions: - It cannot contain carriage return, newline, or tab characters. - It cannot consist of only whitespace characters. - It must be no longer than 64 characters.
	Synonym string `json:"synonym,omitempty"`
}

// UpdateValue : UpdateValue struct
type UpdateValue struct {

	// The text of the entity value. This string must conform to the following restrictions: - It cannot contain carriage return, newline, or tab characters. - It cannot consist of only whitespace characters. - It must be no longer than 64 characters.
	Value string `json:"value,omitempty"`

	// Any metadata related to the entity value.
	Metadata interface{} `json:"metadata,omitempty"`

	// Specifies the type of value.
	ValueType string `json:"value_type,omitempty"`

	// An array of synonyms for the entity value. You can provide either synonyms or patterns (as indicated by **type**), but not both. A synonym must conform to the following resrictions: - It cannot contain carriage return, newline, or tab characters. - It cannot consist of only whitespace characters. - It must be no longer than 64 characters.
	Synonyms []string `json:"synonyms,omitempty"`

	// An array of patterns for the entity value. You can provide either synonyms or patterns (as indicated by **type**), but not both. A pattern is a regular expression no longer than 512 characters. For more information about how to specify a pattern, see the [documentation](https://console.bluemix.net/docs/services/conversation/entities.html#creating-entities).
	Patterns []string `json:"patterns,omitempty"`
}

// UpdateWorkspace : UpdateWorkspace struct
type UpdateWorkspace struct {

	// The name of the workspace. This string cannot contain carriage return, newline, or tab characters, and it must be no longer than 64 characters.
	Name string `json:"name,omitempty"`

	// The description of the workspace. This string cannot contain carriage return, newline, or tab characters, and it must be no longer than 128 characters.
	Description string `json:"description,omitempty"`

	// The language of the workspace.
	Language string `json:"language,omitempty"`

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

	// Whether training data from the workspace can be used by IBM for general service improvements. `true` indicates that workspace training data is not to be used.
	LearningOptOut bool `json:"learning_opt_out,omitempty"`

	// Global settings for the workspace.
	SystemSettings WorkspaceSystemSettings `json:"system_settings,omitempty"`
}

// Value : Value struct
type Value struct {

	// The text of the entity value.
	ValueText string `json:"value_text"`

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
	ValueType string `json:"value_type"`
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
	ValueText string `json:"value_text"`

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
	ValueType string `json:"value_type"`
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

	// Reserved for future use.
	Disambiguation interface{} `json:"disambiguation,omitempty"`

	// For internal use only.
	HumanAgentAssist interface{} `json:"human_agent_assist,omitempty"`
}

// WorkspaceSystemSettingsTooling : WorkspaceSystemSettingsTooling struct
type WorkspaceSystemSettingsTooling struct {

	// Whether the dialog JSON editor displays text responses within the `output.generic` object.
	StoreGenericResponses bool `json:"store_generic_responses,omitempty"`
}
