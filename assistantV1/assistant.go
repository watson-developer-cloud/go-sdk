package assistantV1

import (
	"bytes"
	"fmt"
	req "github.com/parnurzeal/gorequest"
	watson "golang-sdk"
)

// AssistantV1 :
type AssistantV1 struct {
	client *watson.Client
}

// NewAssistantV1 :
func NewAssistantV1(creds watson.Credentials) (*AssistantV1, error) {
	client, clientErr := watson.NewClient(creds, "conversation")

	if clientErr != nil {
		return nil, clientErr
	}

	return &AssistantV1{client: client}, nil
}

func (assistant *AssistantV1) ListWorkspaces(params *ListWorkspacesRequest) (*ListWorkspacesResponse, []error) {
	path := "/v1/workspaces"
	creds := assistant.client.Creds()
	useTM := assistant.client.UseTM()
	tokenManager := assistant.client.TokenManager()

	request := req.New().Get(creds.ServiceURL+path).
		Set("Accept", "application/json").
		Query("version=" + creds.Version).
		Query(*params)

	if useTM {
		token, tokenErr := tokenManager.GetToken()

		if tokenErr != nil {
			return nil, tokenErr
		}

		request.Set("Authorization", "Bearer "+token)
	} else {
		request.SetBasicAuth(creds.Username, creds.Password)
	}

	response := new(ListWorkspacesResponse)
	res, _, err := request.EndStruct(&response)

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

func (assistant *AssistantV1) GetWorkspace(workspaceID string, params *GetWorkspaceRequest) (*GetWorkspaceResponse, []error) {
	path := "/v1/workspaces/" + workspaceID
	creds := assistant.client.Creds()
	useTM := assistant.client.UseTM()
	tokenManager := assistant.client.TokenManager()

	request := req.New().Get(creds.ServiceURL+path).
		Set("Accept", "application/json").
		Query("version=" + creds.Version).
		Query(*params)

	if useTM {
		token, tokenErr := tokenManager.GetToken()

		if tokenErr != nil {
			return nil, tokenErr
		}

		request.Set("Authorization", "Bearer "+token)
	} else {
		request.SetBasicAuth(creds.Username, creds.Password)
	}

	response := new(GetWorkspaceResponse)
	res, _, err := request.EndStruct(&response)

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

func (assistant *AssistantV1) DeleteWorkspace(workspaceID string) []error {
	path := "/v1/workspaces/" + workspaceID
	creds := assistant.client.Creds()
	useTM := assistant.client.UseTM()
	tokenManager := assistant.client.TokenManager()

	request := req.New().Delete(creds.ServiceURL+path).
		Set("Accept", "application/json").
		Query("version=" + creds.Version)

	if useTM {
		token, tokenErr := tokenManager.GetToken()

		if tokenErr != nil {
			return tokenErr
		}

		request.Set("Authorization", "Bearer "+token)
	} else {
		request.SetBasicAuth(creds.Username, creds.Password)
	}

	res, _, err := request.End()

	if err != nil {
		return err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		buff := new(bytes.Buffer)
		buff.ReadFrom(res.Body)
		errStr := buff.String()
		err = append(err, fmt.Errorf(errStr))
		return err
	}

	return nil
}

func (assistant *AssistantV1) CreateWorkspace(body *CreateWorkspace) (*Workspace, []error) {
	path := "/v1/workspaces"
	creds := assistant.client.Creds()
	useTM := assistant.client.UseTM()
	tokenManager := assistant.client.TokenManager()

	request := req.New().Post(creds.ServiceURL+path).
		Set("Accept", "application/json").
		Query("version=" + creds.Version).
		Send(body)

	if useTM {
		token, tokenErr := tokenManager.GetToken()

		if tokenErr != nil {
			return nil, tokenErr
		}

		request.Set("Authorization", "Bearer "+token)
	} else {
		request.SetBasicAuth(creds.Username, creds.Password)
	}

	response := new(Workspace)
	res, _, err := request.EndStruct(&response)

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

func (assistant *AssistantV1) UpdateWorkspace(workspaceID string, body *CreateWorkspace, params *UpdateWorkspaceRequest) (*Workspace, []error) {
	path := "/v1/workspaces/" + workspaceID
	creds := assistant.client.Creds()
	useTM := assistant.client.UseTM()
	tokenManager := assistant.client.TokenManager()

	request := req.New().Post(creds.ServiceURL+path).
		Set("Accept", "application/json").
		Query("version=" + creds.Version).
		Query(*params).
		Send(body)

	if useTM {
		token, tokenErr := tokenManager.GetToken()

		if tokenErr != nil {
			return nil, tokenErr
		}

		request.Set("Authorization", "Bearer "+token)
	} else {
		request.SetBasicAuth(creds.Username, creds.Password)
	}

	response := new(Workspace)
	res, _, err := request.EndStruct(&response)

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
