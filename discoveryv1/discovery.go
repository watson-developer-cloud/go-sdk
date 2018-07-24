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
    "bytes"
    "fmt"
    "github.com/go-openapi/strfmt"
    "os"
    "strings"
    req "github.com/parnurzeal/gorequest"
    watson "golang-sdk"
)

// DiscoveryV1 : The DiscoveryV1 service
type DiscoveryV1 struct {
	client *watson.Client
}

// NewDiscoveryV1 : Instantiate DiscoveryV1
func NewDiscoveryV1(creds watson.Credentials) (*DiscoveryV1, error) {
    if creds.ServiceURL == "" {
        creds.ServiceURL = "https://gateway.watsonplatform.net/discovery/api"
    }

	client, clientErr := watson.NewClient(creds, "discovery")

	if clientErr != nil {
		return nil, clientErr
	}

	return &DiscoveryV1{ client: client }, nil
}

// CreateEnvironment : Create an environment
func (discovery *DiscoveryV1) CreateEnvironment(body *CreateEnvironmentRequest) (*watson.WatsonResponse, []error) {
    path := "/v1/environments"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

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

    response.Result = new(Environment)
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

// GetCreateEnvironmentResult : Cast result of CreateEnvironment operation
func GetCreateEnvironmentResult(response *watson.WatsonResponse) *Environment {
    result, ok := response.Result.(*Environment)

    if ok {
        return result
    }

    return nil
}

// DeleteEnvironment : Delete environment
func (discovery *DiscoveryV1) DeleteEnvironment(environmentID string) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
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

    response.Result = new(DeleteEnvironmentResponse)
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

// GetDeleteEnvironmentResult : Cast result of DeleteEnvironment operation
func GetDeleteEnvironmentResult(response *watson.WatsonResponse) *DeleteEnvironmentResponse {
    result, ok := response.Result.(*DeleteEnvironmentResponse)

    if ok {
        return result
    }

    return nil
}

// GetEnvironment : Get environment info
func (discovery *DiscoveryV1) GetEnvironment(environmentID string) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
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

    response.Result = new(Environment)
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

// GetGetEnvironmentResult : Cast result of GetEnvironment operation
func GetGetEnvironmentResult(response *watson.WatsonResponse) *Environment {
    result, ok := response.Result.(*Environment)

    if ok {
        return result
    }

    return nil
}

// ListEnvironments : List environments
func (discovery *DiscoveryV1) ListEnvironments(name string) (*watson.WatsonResponse, []error) {
    path := "/v1/environments"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("name=" + fmt.Sprint(name))

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

    response.Result = new(ListEnvironmentsResponse)
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

// GetListEnvironmentsResult : Cast result of ListEnvironments operation
func GetListEnvironmentsResult(response *watson.WatsonResponse) *ListEnvironmentsResponse {
    result, ok := response.Result.(*ListEnvironmentsResponse)

    if ok {
        return result
    }

    return nil
}

// ListFields : List fields across collections
func (discovery *DiscoveryV1) ListFields(environmentID string, collectionIds []string) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/fields"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("collection_ids=" + fmt.Sprint(collectionIds))

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

    response.Result = new(ListCollectionFieldsResponse)
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

// GetListFieldsResult : Cast result of ListFields operation
func GetListFieldsResult(response *watson.WatsonResponse) *ListCollectionFieldsResponse {
    result, ok := response.Result.(*ListCollectionFieldsResponse)

    if ok {
        return result
    }

    return nil
}

// UpdateEnvironment : Update an environment
func (discovery *DiscoveryV1) UpdateEnvironment(environmentID string, body *UpdateEnvironmentRequest) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    request := req.New().Put(creds.ServiceURL + path)

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

    response.Result = new(Environment)
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

// GetUpdateEnvironmentResult : Cast result of UpdateEnvironment operation
func GetUpdateEnvironmentResult(response *watson.WatsonResponse) *Environment {
    result, ok := response.Result.(*Environment)

    if ok {
        return result
    }

    return nil
}

// CreateConfiguration : Add configuration
func (discovery *DiscoveryV1) CreateConfiguration(environmentID string, body *Configuration) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/configurations"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
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

    response.Result = new(Configuration)
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

// GetCreateConfigurationResult : Cast result of CreateConfiguration operation
func GetCreateConfigurationResult(response *watson.WatsonResponse) *Configuration {
    result, ok := response.Result.(*Configuration)

    if ok {
        return result
    }

    return nil
}

// DeleteConfiguration : Delete a configuration
func (discovery *DiscoveryV1) DeleteConfiguration(environmentID string, configurationID string) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/configurations/{configuration_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    path = strings.Replace(path, "{configuration_id}", configurationID, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
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

    response.Result = new(DeleteConfigurationResponse)
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

// GetDeleteConfigurationResult : Cast result of DeleteConfiguration operation
func GetDeleteConfigurationResult(response *watson.WatsonResponse) *DeleteConfigurationResponse {
    result, ok := response.Result.(*DeleteConfigurationResponse)

    if ok {
        return result
    }

    return nil
}

// GetConfiguration : Get configuration details
func (discovery *DiscoveryV1) GetConfiguration(environmentID string, configurationID string) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/configurations/{configuration_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    path = strings.Replace(path, "{configuration_id}", configurationID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
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

    response.Result = new(Configuration)
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

// GetGetConfigurationResult : Cast result of GetConfiguration operation
func GetGetConfigurationResult(response *watson.WatsonResponse) *Configuration {
    result, ok := response.Result.(*Configuration)

    if ok {
        return result
    }

    return nil
}

// ListConfigurations : List configurations
func (discovery *DiscoveryV1) ListConfigurations(environmentID string, name string) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/configurations"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("name=" + fmt.Sprint(name))

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

    response.Result = new(ListConfigurationsResponse)
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

// GetListConfigurationsResult : Cast result of ListConfigurations operation
func GetListConfigurationsResult(response *watson.WatsonResponse) *ListConfigurationsResponse {
    result, ok := response.Result.(*ListConfigurationsResponse)

    if ok {
        return result
    }

    return nil
}

// UpdateConfiguration : Update a configuration
func (discovery *DiscoveryV1) UpdateConfiguration(environmentID string, configurationID string, body *Configuration) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/configurations/{configuration_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    path = strings.Replace(path, "{configuration_id}", configurationID, 1)
    request := req.New().Put(creds.ServiceURL + path)

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

    response.Result = new(Configuration)
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

// GetUpdateConfigurationResult : Cast result of UpdateConfiguration operation
func GetUpdateConfigurationResult(response *watson.WatsonResponse) *Configuration {
    result, ok := response.Result.(*Configuration)

    if ok {
        return result
    }

    return nil
}

// TestConfigurationInEnvironment : Test configuration
func (discovery *DiscoveryV1) TestConfigurationInEnvironment(environmentID string, configuration string, step string, configurationID string, file os.File, metadata string, fileContentType string) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/preview"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "multipart/form-data")
    request.Query("version=" + creds.Version)
    request.Query("step=" + fmt.Sprint(step))
    request.Query("configuration_id=" + fmt.Sprint(configurationID))
    request.Type("multipart")
    request.SendFile(file)

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

    response.Result = new(TestDocument)
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

// GetTestConfigurationInEnvironmentResult : Cast result of TestConfigurationInEnvironment operation
func GetTestConfigurationInEnvironmentResult(response *watson.WatsonResponse) *TestDocument {
    result, ok := response.Result.(*TestDocument)

    if ok {
        return result
    }

    return nil
}

// CreateCollection : Create a collection
func (discovery *DiscoveryV1) CreateCollection(environmentID string, body *CreateCollectionRequest) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
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

    response.Result = new(Collection)
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

// GetCreateCollectionResult : Cast result of CreateCollection operation
func GetCreateCollectionResult(response *watson.WatsonResponse) *Collection {
    result, ok := response.Result.(*Collection)

    if ok {
        return result
    }

    return nil
}

// DeleteCollection : Delete a collection
func (discovery *DiscoveryV1) DeleteCollection(environmentID string, collectionID string) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    path = strings.Replace(path, "{collection_id}", collectionID, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
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

    response.Result = new(DeleteCollectionResponse)
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

// GetDeleteCollectionResult : Cast result of DeleteCollection operation
func GetDeleteCollectionResult(response *watson.WatsonResponse) *DeleteCollectionResponse {
    result, ok := response.Result.(*DeleteCollectionResponse)

    if ok {
        return result
    }

    return nil
}

// GetCollection : Get collection details
func (discovery *DiscoveryV1) GetCollection(environmentID string, collectionID string) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    path = strings.Replace(path, "{collection_id}", collectionID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
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

    response.Result = new(Collection)
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

// GetGetCollectionResult : Cast result of GetCollection operation
func GetGetCollectionResult(response *watson.WatsonResponse) *Collection {
    result, ok := response.Result.(*Collection)

    if ok {
        return result
    }

    return nil
}

// ListCollectionFields : List collection fields
func (discovery *DiscoveryV1) ListCollectionFields(environmentID string, collectionID string) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/fields"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    path = strings.Replace(path, "{collection_id}", collectionID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
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

    response.Result = new(ListCollectionFieldsResponse)
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

// GetListCollectionFieldsResult : Cast result of ListCollectionFields operation
func GetListCollectionFieldsResult(response *watson.WatsonResponse) *ListCollectionFieldsResponse {
    result, ok := response.Result.(*ListCollectionFieldsResponse)

    if ok {
        return result
    }

    return nil
}

// ListCollections : List collections
func (discovery *DiscoveryV1) ListCollections(environmentID string, name string) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("name=" + fmt.Sprint(name))

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

    response.Result = new(ListCollectionsResponse)
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

// GetListCollectionsResult : Cast result of ListCollections operation
func GetListCollectionsResult(response *watson.WatsonResponse) *ListCollectionsResponse {
    result, ok := response.Result.(*ListCollectionsResponse)

    if ok {
        return result
    }

    return nil
}

// UpdateCollection : Update a collection
func (discovery *DiscoveryV1) UpdateCollection(environmentID string, collectionID string, body *UpdateCollectionRequest) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    path = strings.Replace(path, "{collection_id}", collectionID, 1)
    request := req.New().Put(creds.ServiceURL + path)

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

    response.Result = new(Collection)
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

// GetUpdateCollectionResult : Cast result of UpdateCollection operation
func GetUpdateCollectionResult(response *watson.WatsonResponse) *Collection {
    result, ok := response.Result.(*Collection)

    if ok {
        return result
    }

    return nil
}

// CreateExpansions : Create or update expansion list
func (discovery *DiscoveryV1) CreateExpansions(environmentID string, collectionID string, body *Expansions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/expansions"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    path = strings.Replace(path, "{collection_id}", collectionID, 1)
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

    response.Result = new(Expansions)
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

// GetCreateExpansionsResult : Cast result of CreateExpansions operation
func GetCreateExpansionsResult(response *watson.WatsonResponse) *Expansions {
    result, ok := response.Result.(*Expansions)

    if ok {
        return result
    }

    return nil
}

// DeleteExpansions : Delete the expansion list
func (discovery *DiscoveryV1) DeleteExpansions(environmentID string, collectionID string) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/expansions"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    path = strings.Replace(path, "{collection_id}", collectionID, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
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


// ListExpansions : Get the expansion list
func (discovery *DiscoveryV1) ListExpansions(environmentID string, collectionID string) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/expansions"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    path = strings.Replace(path, "{collection_id}", collectionID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
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

    response.Result = new(Expansions)
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

// GetListExpansionsResult : Cast result of ListExpansions operation
func GetListExpansionsResult(response *watson.WatsonResponse) *Expansions {
    result, ok := response.Result.(*Expansions)

    if ok {
        return result
    }

    return nil
}

// AddDocument : Add a document
func (discovery *DiscoveryV1) AddDocument(environmentID string, collectionID string, file os.File, metadata string, fileContentType string) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/documents"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    path = strings.Replace(path, "{collection_id}", collectionID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "multipart/form-data")
    request.Query("version=" + creds.Version)
    request.Type("multipart")
    request.SendFile(file)

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

    response.Result = new(DocumentAccepted)
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

// GetAddDocumentResult : Cast result of AddDocument operation
func GetAddDocumentResult(response *watson.WatsonResponse) *DocumentAccepted {
    result, ok := response.Result.(*DocumentAccepted)

    if ok {
        return result
    }

    return nil
}

// DeleteDocument : Delete a document
func (discovery *DiscoveryV1) DeleteDocument(environmentID string, collectionID string, documentID string) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/documents/{document_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    path = strings.Replace(path, "{collection_id}", collectionID, 1)
    path = strings.Replace(path, "{document_id}", documentID, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
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

    response.Result = new(DeleteDocumentResponse)
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

// GetDeleteDocumentResult : Cast result of DeleteDocument operation
func GetDeleteDocumentResult(response *watson.WatsonResponse) *DeleteDocumentResponse {
    result, ok := response.Result.(*DeleteDocumentResponse)

    if ok {
        return result
    }

    return nil
}

// GetDocumentStatus : Get document details
func (discovery *DiscoveryV1) GetDocumentStatus(environmentID string, collectionID string, documentID string) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/documents/{document_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    path = strings.Replace(path, "{collection_id}", collectionID, 1)
    path = strings.Replace(path, "{document_id}", documentID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
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

    response.Result = new(DocumentStatus)
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

// GetGetDocumentStatusResult : Cast result of GetDocumentStatus operation
func GetGetDocumentStatusResult(response *watson.WatsonResponse) *DocumentStatus {
    result, ok := response.Result.(*DocumentStatus)

    if ok {
        return result
    }

    return nil
}

// UpdateDocument : Update a document
func (discovery *DiscoveryV1) UpdateDocument(environmentID string, collectionID string, documentID string, file os.File, metadata string, fileContentType string) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/documents/{document_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    path = strings.Replace(path, "{collection_id}", collectionID, 1)
    path = strings.Replace(path, "{document_id}", documentID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "multipart/form-data")
    request.Query("version=" + creds.Version)
    request.Type("multipart")
    request.SendFile(file)

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

    response.Result = new(DocumentAccepted)
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

// GetUpdateDocumentResult : Cast result of UpdateDocument operation
func GetUpdateDocumentResult(response *watson.WatsonResponse) *DocumentAccepted {
    result, ok := response.Result.(*DocumentAccepted)

    if ok {
        return result
    }

    return nil
}

// FederatedQuery : Query documents in multiple collections
func (discovery *DiscoveryV1) FederatedQuery(environmentID string, collectionIds []string, filter string, query string, naturalLanguageQuery string, aggregation string, count int64, returnFields []string, offset int64, sort []string, highlight bool, deduplicate bool, deduplicateField string, similar bool, similarDocumentIds []string, similarFields []string, passages bool, passagesFields []string, passagesCount int64, passagesCharacters int64) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/query"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("collection_ids=" + fmt.Sprint(collectionIds))
    request.Query("filter=" + fmt.Sprint(filter))
    request.Query("query=" + fmt.Sprint(query))
    request.Query("natural_language_query=" + fmt.Sprint(naturalLanguageQuery))
    request.Query("aggregation=" + fmt.Sprint(aggregation))
    request.Query("count=" + fmt.Sprint(count))
    request.Query("return=" + fmt.Sprint(returnFields))
    request.Query("offset=" + fmt.Sprint(offset))
    request.Query("sort=" + fmt.Sprint(sort))
    request.Query("highlight=" + fmt.Sprint(highlight))
    request.Query("deduplicate=" + fmt.Sprint(deduplicate))
    request.Query("deduplicate.field=" + fmt.Sprint(deduplicateField))
    request.Query("similar=" + fmt.Sprint(similar))
    request.Query("similar.document_ids=" + fmt.Sprint(similarDocumentIds))
    request.Query("similar.fields=" + fmt.Sprint(similarFields))
    request.Query("passages=" + fmt.Sprint(passages))
    request.Query("passages.fields=" + fmt.Sprint(passagesFields))
    request.Query("passages.count=" + fmt.Sprint(passagesCount))
    request.Query("passages.characters=" + fmt.Sprint(passagesCharacters))

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

    response.Result = new(QueryResponse)
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

// GetFederatedQueryResult : Cast result of FederatedQuery operation
func GetFederatedQueryResult(response *watson.WatsonResponse) *QueryResponse {
    result, ok := response.Result.(*QueryResponse)

    if ok {
        return result
    }

    return nil
}

// FederatedQueryNotices : Query multiple collection system notices
func (discovery *DiscoveryV1) FederatedQueryNotices(environmentID string, collectionIds []string, filter string, query string, naturalLanguageQuery string, aggregation string, count int64, returnFields []string, offset int64, sort []string, highlight bool, deduplicateField string, similar bool, similarDocumentIds []string, similarFields []string) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/notices"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("collection_ids=" + fmt.Sprint(collectionIds))
    request.Query("filter=" + fmt.Sprint(filter))
    request.Query("query=" + fmt.Sprint(query))
    request.Query("natural_language_query=" + fmt.Sprint(naturalLanguageQuery))
    request.Query("aggregation=" + fmt.Sprint(aggregation))
    request.Query("count=" + fmt.Sprint(count))
    request.Query("return=" + fmt.Sprint(returnFields))
    request.Query("offset=" + fmt.Sprint(offset))
    request.Query("sort=" + fmt.Sprint(sort))
    request.Query("highlight=" + fmt.Sprint(highlight))
    request.Query("deduplicate.field=" + fmt.Sprint(deduplicateField))
    request.Query("similar=" + fmt.Sprint(similar))
    request.Query("similar.document_ids=" + fmt.Sprint(similarDocumentIds))
    request.Query("similar.fields=" + fmt.Sprint(similarFields))

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

    response.Result = new(QueryNoticesResponse)
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

// GetFederatedQueryNoticesResult : Cast result of FederatedQueryNotices operation
func GetFederatedQueryNoticesResult(response *watson.WatsonResponse) *QueryNoticesResponse {
    result, ok := response.Result.(*QueryNoticesResponse)

    if ok {
        return result
    }

    return nil
}

// Query : Query your collection
func (discovery *DiscoveryV1) Query(environmentID string, collectionID string, filter string, query string, naturalLanguageQuery string, passages bool, aggregation string, count int64, returnFields []string, offset int64, sort []string, highlight bool, passagesFields []string, passagesCount int64, passagesCharacters int64, deduplicate bool, deduplicateField string, similar bool, similarDocumentIds []string, similarFields []string) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/query"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    path = strings.Replace(path, "{collection_id}", collectionID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("filter=" + fmt.Sprint(filter))
    request.Query("query=" + fmt.Sprint(query))
    request.Query("natural_language_query=" + fmt.Sprint(naturalLanguageQuery))
    request.Query("passages=" + fmt.Sprint(passages))
    request.Query("aggregation=" + fmt.Sprint(aggregation))
    request.Query("count=" + fmt.Sprint(count))
    request.Query("return=" + fmt.Sprint(returnFields))
    request.Query("offset=" + fmt.Sprint(offset))
    request.Query("sort=" + fmt.Sprint(sort))
    request.Query("highlight=" + fmt.Sprint(highlight))
    request.Query("passages.fields=" + fmt.Sprint(passagesFields))
    request.Query("passages.count=" + fmt.Sprint(passagesCount))
    request.Query("passages.characters=" + fmt.Sprint(passagesCharacters))
    request.Query("deduplicate=" + fmt.Sprint(deduplicate))
    request.Query("deduplicate.field=" + fmt.Sprint(deduplicateField))
    request.Query("similar=" + fmt.Sprint(similar))
    request.Query("similar.document_ids=" + fmt.Sprint(similarDocumentIds))
    request.Query("similar.fields=" + fmt.Sprint(similarFields))

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

    response.Result = new(QueryResponse)
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

// GetQueryResult : Cast result of Query operation
func GetQueryResult(response *watson.WatsonResponse) *QueryResponse {
    result, ok := response.Result.(*QueryResponse)

    if ok {
        return result
    }

    return nil
}

// QueryEntities : Knowledge Graph entity query
func (discovery *DiscoveryV1) QueryEntities(environmentID string, collectionID string, body *QueryEntities) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/query_entities"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    path = strings.Replace(path, "{collection_id}", collectionID, 1)
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

    response.Result = new(QueryEntitiesResponse)
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

// GetQueryEntitiesResult : Cast result of QueryEntities operation
func GetQueryEntitiesResult(response *watson.WatsonResponse) *QueryEntitiesResponse {
    result, ok := response.Result.(*QueryEntitiesResponse)

    if ok {
        return result
    }

    return nil
}

// QueryNotices : Query system notices
func (discovery *DiscoveryV1) QueryNotices(environmentID string, collectionID string, filter string, query string, naturalLanguageQuery string, passages bool, aggregation string, count int64, returnFields []string, offset int64, sort []string, highlight bool, passagesFields []string, passagesCount int64, passagesCharacters int64, deduplicateField string, similar bool, similarDocumentIds []string, similarFields []string) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/notices"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    path = strings.Replace(path, "{collection_id}", collectionID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("filter=" + fmt.Sprint(filter))
    request.Query("query=" + fmt.Sprint(query))
    request.Query("natural_language_query=" + fmt.Sprint(naturalLanguageQuery))
    request.Query("passages=" + fmt.Sprint(passages))
    request.Query("aggregation=" + fmt.Sprint(aggregation))
    request.Query("count=" + fmt.Sprint(count))
    request.Query("return=" + fmt.Sprint(returnFields))
    request.Query("offset=" + fmt.Sprint(offset))
    request.Query("sort=" + fmt.Sprint(sort))
    request.Query("highlight=" + fmt.Sprint(highlight))
    request.Query("passages.fields=" + fmt.Sprint(passagesFields))
    request.Query("passages.count=" + fmt.Sprint(passagesCount))
    request.Query("passages.characters=" + fmt.Sprint(passagesCharacters))
    request.Query("deduplicate.field=" + fmt.Sprint(deduplicateField))
    request.Query("similar=" + fmt.Sprint(similar))
    request.Query("similar.document_ids=" + fmt.Sprint(similarDocumentIds))
    request.Query("similar.fields=" + fmt.Sprint(similarFields))

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

    response.Result = new(QueryNoticesResponse)
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

// GetQueryNoticesResult : Cast result of QueryNotices operation
func GetQueryNoticesResult(response *watson.WatsonResponse) *QueryNoticesResponse {
    result, ok := response.Result.(*QueryNoticesResponse)

    if ok {
        return result
    }

    return nil
}

// QueryRelations : Knowledge Graph relationship query
func (discovery *DiscoveryV1) QueryRelations(environmentID string, collectionID string, body *QueryRelations) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/query_relations"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    path = strings.Replace(path, "{collection_id}", collectionID, 1)
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

    response.Result = new(QueryRelationsResponse)
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

// GetQueryRelationsResult : Cast result of QueryRelations operation
func GetQueryRelationsResult(response *watson.WatsonResponse) *QueryRelationsResponse {
    result, ok := response.Result.(*QueryRelationsResponse)

    if ok {
        return result
    }

    return nil
}

// AddTrainingData : Add query to training data
func (discovery *DiscoveryV1) AddTrainingData(environmentID string, collectionID string, body *NewTrainingQuery) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/training_data"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    path = strings.Replace(path, "{collection_id}", collectionID, 1)
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

    response.Result = new(TrainingQuery)
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

// GetAddTrainingDataResult : Cast result of AddTrainingData operation
func GetAddTrainingDataResult(response *watson.WatsonResponse) *TrainingQuery {
    result, ok := response.Result.(*TrainingQuery)

    if ok {
        return result
    }

    return nil
}

// CreateTrainingExample : Add example to training data query
func (discovery *DiscoveryV1) CreateTrainingExample(environmentID string, collectionID string, queryID string, body *TrainingExample) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}/examples"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    path = strings.Replace(path, "{collection_id}", collectionID, 1)
    path = strings.Replace(path, "{query_id}", queryID, 1)
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

    response.Result = new(TrainingExample)
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

// GetCreateTrainingExampleResult : Cast result of CreateTrainingExample operation
func GetCreateTrainingExampleResult(response *watson.WatsonResponse) *TrainingExample {
    result, ok := response.Result.(*TrainingExample)

    if ok {
        return result
    }

    return nil
}

// DeleteAllTrainingData : Delete all training data
func (discovery *DiscoveryV1) DeleteAllTrainingData(environmentID string, collectionID string) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/training_data"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    path = strings.Replace(path, "{collection_id}", collectionID, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
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


// DeleteTrainingData : Delete a training data query
func (discovery *DiscoveryV1) DeleteTrainingData(environmentID string, collectionID string, queryID string) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    path = strings.Replace(path, "{collection_id}", collectionID, 1)
    path = strings.Replace(path, "{query_id}", queryID, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
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


// DeleteTrainingExample : Delete example for training data query
func (discovery *DiscoveryV1) DeleteTrainingExample(environmentID string, collectionID string, queryID string, exampleID string) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}/examples/{example_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    path = strings.Replace(path, "{collection_id}", collectionID, 1)
    path = strings.Replace(path, "{query_id}", queryID, 1)
    path = strings.Replace(path, "{example_id}", exampleID, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
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


// GetTrainingData : Get details about a query
func (discovery *DiscoveryV1) GetTrainingData(environmentID string, collectionID string, queryID string) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    path = strings.Replace(path, "{collection_id}", collectionID, 1)
    path = strings.Replace(path, "{query_id}", queryID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
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

    response.Result = new(TrainingQuery)
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

// GetGetTrainingDataResult : Cast result of GetTrainingData operation
func GetGetTrainingDataResult(response *watson.WatsonResponse) *TrainingQuery {
    result, ok := response.Result.(*TrainingQuery)

    if ok {
        return result
    }

    return nil
}

// GetTrainingExample : Get details for training data example
func (discovery *DiscoveryV1) GetTrainingExample(environmentID string, collectionID string, queryID string, exampleID string) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}/examples/{example_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    path = strings.Replace(path, "{collection_id}", collectionID, 1)
    path = strings.Replace(path, "{query_id}", queryID, 1)
    path = strings.Replace(path, "{example_id}", exampleID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
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

    response.Result = new(TrainingExample)
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

// GetGetTrainingExampleResult : Cast result of GetTrainingExample operation
func GetGetTrainingExampleResult(response *watson.WatsonResponse) *TrainingExample {
    result, ok := response.Result.(*TrainingExample)

    if ok {
        return result
    }

    return nil
}

// ListTrainingData : List training data
func (discovery *DiscoveryV1) ListTrainingData(environmentID string, collectionID string) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/training_data"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    path = strings.Replace(path, "{collection_id}", collectionID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
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

    response.Result = new(TrainingDataSet)
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

// GetListTrainingDataResult : Cast result of ListTrainingData operation
func GetListTrainingDataResult(response *watson.WatsonResponse) *TrainingDataSet {
    result, ok := response.Result.(*TrainingDataSet)

    if ok {
        return result
    }

    return nil
}

// ListTrainingExamples : List examples for a training data query
func (discovery *DiscoveryV1) ListTrainingExamples(environmentID string, collectionID string, queryID string) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}/examples"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    path = strings.Replace(path, "{collection_id}", collectionID, 1)
    path = strings.Replace(path, "{query_id}", queryID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
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

    response.Result = new(TrainingExampleList)
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

// GetListTrainingExamplesResult : Cast result of ListTrainingExamples operation
func GetListTrainingExamplesResult(response *watson.WatsonResponse) *TrainingExampleList {
    result, ok := response.Result.(*TrainingExampleList)

    if ok {
        return result
    }

    return nil
}

// UpdateTrainingExample : Change label or cross reference for example
func (discovery *DiscoveryV1) UpdateTrainingExample(environmentID string, collectionID string, queryID string, exampleID string, body *TrainingExamplePatch) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}/examples/{example_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    path = strings.Replace(path, "{collection_id}", collectionID, 1)
    path = strings.Replace(path, "{query_id}", queryID, 1)
    path = strings.Replace(path, "{example_id}", exampleID, 1)
    request := req.New().Put(creds.ServiceURL + path)

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

    response.Result = new(TrainingExample)
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

// GetUpdateTrainingExampleResult : Cast result of UpdateTrainingExample operation
func GetUpdateTrainingExampleResult(response *watson.WatsonResponse) *TrainingExample {
    result, ok := response.Result.(*TrainingExample)

    if ok {
        return result
    }

    return nil
}

// DeleteUserData : Delete labeled data
func (discovery *DiscoveryV1) DeleteUserData(customerID string) (*watson.WatsonResponse, []error) {
    path := "/v1/user_data"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    request := req.New().Delete(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
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


// CreateCredentials : Create credentials
func (discovery *DiscoveryV1) CreateCredentials(environmentID string, body *Credentials) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/credentials"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
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

    response.Result = new(Credentials)
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

// GetCreateCredentialsResult : Cast result of CreateCredentials operation
func GetCreateCredentialsResult(response *watson.WatsonResponse) *Credentials {
    result, ok := response.Result.(*Credentials)

    if ok {
        return result
    }

    return nil
}

// DeleteCredentials : Delete credentials
func (discovery *DiscoveryV1) DeleteCredentials(environmentID string, credentialID string) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/credentials/{credential_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    path = strings.Replace(path, "{credential_id}", credentialID, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
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

    response.Result = new(DeleteCredentials)
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

// GetDeleteCredentialsResult : Cast result of DeleteCredentials operation
func GetDeleteCredentialsResult(response *watson.WatsonResponse) *DeleteCredentials {
    result, ok := response.Result.(*DeleteCredentials)

    if ok {
        return result
    }

    return nil
}

// GetCredentials : View Credentials
func (discovery *DiscoveryV1) GetCredentials(environmentID string, credentialID string) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/credentials/{credential_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    path = strings.Replace(path, "{credential_id}", credentialID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
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

    response.Result = new(Credentials)
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

// GetGetCredentialsResult : Cast result of GetCredentials operation
func GetGetCredentialsResult(response *watson.WatsonResponse) *Credentials {
    result, ok := response.Result.(*Credentials)

    if ok {
        return result
    }

    return nil
}

// ListCredentials : List credentials
func (discovery *DiscoveryV1) ListCredentials(environmentID string) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/credentials"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
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

    response.Result = new(CredentialsList)
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

// GetListCredentialsResult : Cast result of ListCredentials operation
func GetListCredentialsResult(response *watson.WatsonResponse) *CredentialsList {
    result, ok := response.Result.(*CredentialsList)

    if ok {
        return result
    }

    return nil
}

// UpdateCredentials : Update credentials
func (discovery *DiscoveryV1) UpdateCredentials(environmentID string, credentialID string, body *Credentials) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/credentials/{credential_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", environmentID, 1)
    path = strings.Replace(path, "{credential_id}", credentialID, 1)
    request := req.New().Put(creds.ServiceURL + path)

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

    response.Result = new(Credentials)
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

// GetUpdateCredentialsResult : Cast result of UpdateCredentials operation
func GetUpdateCredentialsResult(response *watson.WatsonResponse) *Credentials {
    result, ok := response.Result.(*Credentials)

    if ok {
        return result
    }

    return nil
}


// AggregationResult : AggregationResult struct
type AggregationResult struct {

	// Key that matched the aggregation type.
	Key string `json:"key,omitempty"`

	// Number of matching results.
	MatchingResults int64 `json:"matching_results,omitempty"`

	// Aggregations returned in the case of chained aggregations.
	Aggregations []QueryAggregation `json:"aggregations,omitempty"`
}

// Collection : A collection for storing documents.
type Collection struct {

	// The unique identifier of the collection.
	CollectionID string `json:"collection_id,omitempty"`

	// The name of the collection.
	Name string `json:"name,omitempty"`

	// The description of the collection.
	Description string `json:"description,omitempty"`

	// The creation date of the collection in the format yyyy-MM-dd'T'HH:mmcon:ss.SSS'Z'.
	Created strfmt.DateTime `json:"created,omitempty"`

	// The timestamp of when the collection was last updated in the format yyyy-MM-dd'T'HH:mm:ss.SSS'Z'.
	Updated strfmt.DateTime `json:"updated,omitempty"`

	// The status of the collection.
	Status string `json:"status,omitempty"`

	// The unique identifier of the collection's configuration.
	ConfigurationID string `json:"configuration_id,omitempty"`

	// The language of the documents stored in the collection. Permitted values include `en` (English), `de` (German), and `es` (Spanish).
	Language string `json:"language,omitempty"`

	// The object providing information about the documents in the collection. Present only when retrieving details of a collection.
	DocumentCounts DocumentCounts `json:"document_counts,omitempty"`

	// The object providing information about the disk usage of the collection. Present only when retrieving details of a collection.
	DiskUsage CollectionDiskUsage `json:"disk_usage,omitempty"`

	// Provides information about the status of relevance training for collection.
	TrainingStatus TrainingStatus `json:"training_status,omitempty"`

	// Object containing source crawl status information.
	SourceCrawl SourceStatus `json:"source_crawl,omitempty"`
}

// CollectionDiskUsage : Summary of the disk usage statistics for this collection.
type CollectionDiskUsage struct {

	// Number of bytes used by the collection.
	UsedBytes int64 `json:"used_bytes,omitempty"`
}

// CollectionUsage : Summary of the collection usage in the environment.
type CollectionUsage struct {

	// Number of active collections in the environment.
	Available int64 `json:"available,omitempty"`

	// Total number of collections allowed in the environment.
	MaximumAllowed int64 `json:"maximum_allowed,omitempty"`
}

// Configuration : A custom configuration for the environment.
type Configuration struct {

	// The unique identifier of the configuration.
	ConfigurationID string `json:"configuration_id,omitempty"`

	// The name of the configuration.
	Name string `json:"name"`

	// The creation date of the configuration in the format yyyy-MM-dd'T'HH:mm:ss.SSS'Z'.
	Created strfmt.DateTime `json:"created,omitempty"`

	// The timestamp of when the configuration was last updated in the format yyyy-MM-dd'T'HH:mm:ss.SSS'Z'.
	Updated strfmt.DateTime `json:"updated,omitempty"`

	// The description of the configuration, if available.
	Description string `json:"description,omitempty"`

	// The document conversion settings for the configuration.
	Conversions Conversions `json:"conversions,omitempty"`

	// An array of document enrichment settings for the configuration.
	Enrichments []Enrichment `json:"enrichments,omitempty"`

	// Defines operations that can be used to transform the final output JSON into a normalized form. Operations are executed in the order that they appear in the array.
	Normalizations []NormalizationOperation `json:"normalizations,omitempty"`

	// Object containing source parameters for the configuration.
	Source Source `json:"source,omitempty"`
}

// Conversions : Document conversion settings.
type Conversions struct {

	// A list of PDF conversion settings.
	Pdf PdfSettings `json:"pdf,omitempty"`

	// A list of Word conversion settings.
	Word WordSettings `json:"word,omitempty"`

	// A list of HTML conversion settings.
	HTML HTMLSettings `json:"html,omitempty"`

	// A list of Document Segmentation settings.
	Segment SegmentSettings `json:"segment,omitempty"`

	// Defines operations that can be used to transform the final output JSON into a normalized form. Operations are executed in the order that they appear in the array.
	JSONNormalizations []NormalizationOperation `json:"json_normalizations,omitempty"`
}

// CreateCollectionRequest : CreateCollectionRequest struct
type CreateCollectionRequest struct {

	// The name of the collection to be created.
	Name string `json:"name"`

	// A description of the collection.
	Description string `json:"description,omitempty"`

	// The ID of the configuration in which the collection is to be created.
	ConfigurationID string `json:"configuration_id,omitempty"`

	// The language of the documents stored in the collection, in the form of an ISO 639-1 language code.
	Language string `json:"language,omitempty"`
}

// CreateEnvironmentRequest : CreateEnvironmentRequest struct
type CreateEnvironmentRequest struct {

	// Name that identifies the environment.
	Name string `json:"name"`

	// Description of the environment.
	Description string `json:"description,omitempty"`

	// **Deprecated**: Size of the environment.
	Size int64 `json:"size,omitempty"`
}

// CredentialDetails : Object containing details of the stored credentials. Obtain credentials for your source from the administrator of the source.
type CredentialDetails struct {

	// The authentication method for this credentials definition. The  **credential_type** specified must be supported by the **source_type**. The following combinations are possible: -  `"source_type": "box"` - valid `credential_type`s: `oauth2` -  `"source_type": "salesforce"` - valid `credential_type`s: `username_password` -  `"source_type": "sharepoint"` - valid `credential_type`s: `saml`.
	CredentialType string `json:"credential_type,omitempty"`

	// The **client_id** of the source that these credentials connect to. Only valid, and required, with a **credential_type** of `oauth2`.
	ClientID string `json:"client_id,omitempty"`

	// The **enterprise_id** of the Box site that these credentials connect to. Only valid, and required, with a **source_type** of `box`.
	EnterpriseID string `json:"enterprise_id,omitempty"`

	// The **url** of the source that these credentials connect to. Only valid, and required, with a **credential_type** of `username_password`.
	URL string `json:"url,omitempty"`

	// The **username** of the source that these credentials connect to. Only valid, and required, with a **credential_type** of `saml` and `username_password`.
	Username string `json:"username,omitempty"`

	// The **organization_url** of the source that these credentials connect to. Only valid, and required, with a **credential_type** of `saml`.
	OrganizationURL string `json:"organization_url,omitempty"`

	// The **site_collection.path** of the source that these credentials connect to. Only valid, and required, with a **source_type** of `sharepoint`.
	SiteCollectionPath string `json:"site_collection_path,omitempty"`

	// The **client_secret** of the source that these credentials connect to. Only valid, and required, with a **credential_type** of `oauth2`. This value is never returned and is only used when creating or modifying **credentials**.
	ClientSecret string `json:"client_secret,omitempty"`

	// The **public_key_id** of the source that these credentials connect to. Only valid, and required, with a **credential_type** of `oauth2`. This value is never returned and is only used when creating or modifying **credentials**.
	PublicKeyID string `json:"public_key_id,omitempty"`

	// The **private_key** of the source that these credentials connect to. Only valid, and required, with a **credential_type** of `oauth2`. This value is never returned and is only used when creating or modifying **credentials**.
	PrivateKey string `json:"private_key,omitempty"`

	// The **passphrase** of the source that these credentials connect to. Only valid, and required, with a **credential_type** of `oauth2`. This value is never returned and is only used when creating or modifying **credentials**.
	Passphrase string `json:"passphrase,omitempty"`

	// The **password** of the source that these credentials connect to. Only valid, and required, with **credential_type**s of `saml` and `username_password`. **Note:** When used with a **source_type** of `salesforce`, the password consists of the Salesforce password and a valid Salesforce security token concatenated. This value is never returned and is only used when creating or modifying **credentials**.
	Password string `json:"password,omitempty"`
}

// Credentials : Object containing credential information.
type Credentials struct {

	// Unique identifier for this set of credentials.
	CredentialID string `json:"credential_id,omitempty"`

	// The source that this credentials object connects to. -  `box` indicates the credentials are used to connect an instance of Enterprise Box. -  `salesforce` indicates the credentials are used to connect to Salesforce. -  `sharepoint` indicates the credentials are used to connect to Microsoft SharePoint Online.
	SourceType string `json:"source_type,omitempty"`

	// Object containing details of the stored credentials. Obtain credentials for your source from the administrator of the source.
	CredentialDetails CredentialDetails `json:"credential_details,omitempty"`
}

// CredentialsList : CredentialsList struct
type CredentialsList struct {

	// An array of credential definitions that were created for this instance.
	Credentials []Credentials `json:"credentials,omitempty"`
}

// DeleteCollectionResponse : DeleteCollectionResponse struct
type DeleteCollectionResponse struct {

	// The unique identifier of the collection that is being deleted.
	CollectionID string `json:"collection_id"`

	// The status of the collection. The status of a successful deletion operation is `deleted`.
	Status string `json:"status"`
}

// DeleteConfigurationResponse : DeleteConfigurationResponse struct
type DeleteConfigurationResponse struct {

	// The unique identifier for the configuration.
	ConfigurationID string `json:"configuration_id"`

	// Status of the configuration. A deleted configuration has the status deleted.
	Status string `json:"status"`

	// An array of notice messages, if any.
	Notices []Notice `json:"notices,omitempty"`
}

// DeleteCredentials : Object returned after credentials are deleted.
type DeleteCredentials struct {

	// The unique identifier of the credentials that have been deleted.
	CredentialID string `json:"credential_id,omitempty"`

	// The status of the deletion request.
	Status string `json:"status,omitempty"`
}

// DeleteDocumentResponse : DeleteDocumentResponse struct
type DeleteDocumentResponse struct {

	// The unique identifier of the document.
	DocumentID string `json:"document_id,omitempty"`

	// Status of the document. A deleted document has the status deleted.
	Status string `json:"status,omitempty"`
}

// DeleteEnvironmentResponse : DeleteEnvironmentResponse struct
type DeleteEnvironmentResponse struct {

	// The unique identifier for the environment.
	EnvironmentID string `json:"environment_id"`

	// Status of the environment.
	Status string `json:"status"`
}

// DiskUsage : Summary of the disk usage statistics for the environment.
type DiskUsage struct {

	// Number of bytes within the environment's disk capacity that are currently used to store data.
	UsedBytes int64 `json:"used_bytes,omitempty"`

	// Total number of bytes available in the environment's disk capacity.
	MaximumAllowedBytes int64 `json:"maximum_allowed_bytes,omitempty"`

	// **Deprecated**: Total number of bytes available in the environment's disk capacity.
	TotalBytes int64 `json:"total_bytes,omitempty"`

	// **Deprecated**: Amount of disk capacity used, in KB or GB format.
	Used string `json:"used,omitempty"`

	// **Deprecated**: Total amount of the environment's disk capacity, in KB or GB format.
	Total string `json:"total,omitempty"`

	// **Deprecated**: Percentage of the environment's disk capacity that is being used.
	PercentUsed float64 `json:"percent_used,omitempty"`
}

// DocumentAccepted : DocumentAccepted struct
type DocumentAccepted struct {

	// The unique identifier of the ingested document.
	DocumentID string `json:"document_id,omitempty"`

	// Status of the document in the ingestion process.
	Status string `json:"status,omitempty"`

	// Array of notices produced by the document-ingestion process.
	Notices []Notice `json:"notices,omitempty"`
}

// DocumentCounts : DocumentCounts struct
type DocumentCounts struct {

	// The total number of available documents in the collection.
	Available int64 `json:"available,omitempty"`

	// The number of documents in the collection that are currently being processed.
	Processing int64 `json:"processing,omitempty"`

	// The number of documents in the collection that failed to be ingested.
	Failed int64 `json:"failed,omitempty"`
}

// DocumentSnapshot : DocumentSnapshot struct
type DocumentSnapshot struct {

	Step string `json:"step,omitempty"`

	Snapshot interface{} `json:"snapshot,omitempty"`
}

// DocumentStatus : Status information about a submitted document.
type DocumentStatus struct {

	// The unique identifier of the document.
	DocumentID string `json:"document_id"`

	// The unique identifier for the configuration.
	ConfigurationID string `json:"configuration_id,omitempty"`

	// The creation date of the document in the format yyyy-MM-dd'T'HH:mm:ss.SSS'Z'.
	Created strfmt.DateTime `json:"created,omitempty"`

	// Date of the most recent document update, in the format yyyy-MM-dd'T'HH:mm:ss.SSS'Z'.
	Updated strfmt.DateTime `json:"updated,omitempty"`

	// Status of the document in the ingestion process.
	Status string `json:"status"`

	// Description of the document status.
	StatusDescription string `json:"status_description"`

	// Name of the original source file (if available).
	Filename string `json:"filename,omitempty"`

	// The type of the original source file.
	FileType string `json:"file_type,omitempty"`

	// The SHA-1 hash of the original source file (formatted as a hexadecimal string).
	Sha1 string `json:"sha1,omitempty"`

	// Array of notices produced by the document-ingestion process.
	Notices []Notice `json:"notices"`
}

// Enrichment : Enrichment struct
type Enrichment struct {

	// Describes what the enrichment step does.
	Description string `json:"description,omitempty"`

	// Field where enrichments will be stored. This field must already exist or be at most 1 level deeper than an existing field. For example, if `text` is a top-level field with no sub-fields, `text.foo` is a valid destination but `text.foo.bar` is not.
	DestinationField string `json:"destination_field"`

	// Field to be enriched.
	SourceField string `json:"source_field"`

	// Indicates that the enrichments will overwrite the destination_field field if it already exists.
	Overwrite bool `json:"overwrite,omitempty"`

	// Name of the enrichment service to call. Current options are `natural_language_understanding` and `elements`. When using `natual_language_understanding`, the **options** object must contain Natural Language Understanding options. When using `elements` the **options** object must contain Element Classification options. Additionally, when using the `elements` enrichment the configuration specified and files ingested must meet all the criteria specified in [the documentation](https://console.bluemix.net/docs/services/discovery/element-classification.html) Previous API versions also supported `alchemy_language`.
	EnrichmentName string `json:"enrichment_name"`

	// If true, then most errors generated during the enrichment process will be treated as warnings and will not cause the document to fail processing.
	IgnoreDownstreamErrors bool `json:"ignore_downstream_errors,omitempty"`

	// A list of options specific to the enrichment.
	Options EnrichmentOptions `json:"options,omitempty"`
}

// EnrichmentOptions : Options which are specific to a particular enrichment.
type EnrichmentOptions struct {

	// An object representing the enrichment features that will be applied to the specified field.
	Features NluEnrichmentFeatures `json:"features,omitempty"`

	// *For use with `elements` enrichments only.* The element extraction model to use. Models available are: `contract`.
	Model string `json:"model,omitempty"`
}

// Environment : Details about an environment.
type Environment struct {

	// Unique identifier for the environment.
	EnvironmentID string `json:"environment_id,omitempty"`

	// Name that identifies the environment.
	Name string `json:"name,omitempty"`

	// Description of the environment.
	Description string `json:"description,omitempty"`

	// Creation date of the environment, in the format `yyyy-MM-dd'T'HH:mm:ss.SSS'Z'`.
	Created strfmt.DateTime `json:"created,omitempty"`

	// Date of most recent environment update, in the format `yyyy-MM-dd'T'HH:mm:ss.SSS'Z'`.
	Updated strfmt.DateTime `json:"updated,omitempty"`

	// Status of the environment.
	Status string `json:"status,omitempty"`

	// If `true`, the environment contains read-only collections that are maintained by IBM.
	ReadOnly bool `json:"read_only,omitempty"`

	// **Deprecated**: Size of the environment.
	Size int64 `json:"size,omitempty"`

	// Details about the resource usage and capacity of the environment.
	IndexCapacity IndexCapacity `json:"index_capacity,omitempty"`
}

// EnvironmentDocuments : Summary of the document usage statistics for the environment.
type EnvironmentDocuments struct {

	// Number of documents indexed for the environment.
	Indexed int64 `json:"indexed,omitempty"`

	// Total number of documents allowed in the environment's capacity.
	MaximumAllowed int64 `json:"maximum_allowed,omitempty"`
}

// Expansion : An expansion definition. Each object respresents one set of expandable strings. For example, you could have expansions for the word `hot` in one object, and expansions for the word `cold` in another.
type Expansion struct {

	// A list of terms that will be expanded for this expansion. If specified, only the items in this list are expanded.
	InputTerms []string `json:"input_terms,omitempty"`

	// A list of terms that this expansion will be expanded to. If specified without **input_terms**, it also functions as the input term list.
	ExpandedTerms []string `json:"expanded_terms"`
}

// Expansions : The query expansion definitions for the specified collection.
type Expansions struct {

	// An array of query expansion definitions. Each object in the **expansions** array represents a term or set of terms that will be expanded into other terms. Each expansion object can be configured as bidirectional or unidirectional. Bidirectional means that all terms are expanded to all other terms in the object. Unidirectional means that a set list of terms can be expanded into a second list of terms. To create a bi-directional expansion specify an **expanded_terms** array. When found in a query, all items in the **expanded_terms** array are then expanded to the other items in the same array. To create a uni-directional expansion, specify both an array of **input_terms** and an array of **expanded_terms**. When items in the **input_terms** array are present in a query, they are expanded using the items listed in the **expanded_terms** array.
	Expansions []Expansion `json:"expansions"`
}

// Field : Field struct
type Field struct {

	// The name of the field.
	FieldName string `json:"field_name,omitempty"`

	// The type of the field.
	FieldType string `json:"field_type,omitempty"`
}

// FontSetting : FontSetting struct
type FontSetting struct {

	Level int64 `json:"level,omitempty"`

	MinSize int64 `json:"min_size,omitempty"`

	MaxSize int64 `json:"max_size,omitempty"`

	Bold bool `json:"bold,omitempty"`

	Italic bool `json:"italic,omitempty"`

	Name string `json:"name,omitempty"`
}

// HTMLSettings : A list of HTML conversion settings.
type HTMLSettings struct {

	ExcludeTagsCompletely []string `json:"exclude_tags_completely,omitempty"`

	ExcludeTagsKeepContent []string `json:"exclude_tags_keep_content,omitempty"`

	KeepContent XPathPatterns `json:"keep_content,omitempty"`

	ExcludeContent XPathPatterns `json:"exclude_content,omitempty"`

	KeepTagAttributes []string `json:"keep_tag_attributes,omitempty"`

	ExcludeTagAttributes []string `json:"exclude_tag_attributes,omitempty"`
}

// IndexCapacity : Details about the resource usage and capacity of the environment.
type IndexCapacity struct {

	// Summary of the document usage statistics for the environment.
	Documents EnvironmentDocuments `json:"documents,omitempty"`

	// Summary of the disk usage of the environment.
	DiskUsage DiskUsage `json:"disk_usage,omitempty"`

	// Summary of the collection usage in the environment.
	Collections CollectionUsage `json:"collections,omitempty"`

	// **Deprecated**: Summary of the memory usage of the environment.
	MemoryUsage MemoryUsage `json:"memory_usage,omitempty"`
}

// ListCollectionFieldsResponse : The list of fetched fields. The fields are returned using a fully qualified name format, however, the format differs slightly from that used by the query operations. * Fields which contain nested JSON objects are assigned a type of "nested". * Fields which belong to a nested object are prefixed with `.properties` (for example, `warnings.properties.severity` means that the `warnings` object has a property called `severity`). * Fields returned from the News collection are prefixed with `v{N}-fullnews-t3-{YEAR}.mappings` (for example, `v5-fullnews-t3-2016.mappings.text.properties.author`).
type ListCollectionFieldsResponse struct {

	// An array containing information about each field in the collections.
	Fields []Field `json:"fields,omitempty"`
}

// ListCollectionsResponse : ListCollectionsResponse struct
type ListCollectionsResponse struct {

	// An array containing information about each collection in the environment.
	Collections []Collection `json:"collections,omitempty"`
}

// ListConfigurationsResponse : ListConfigurationsResponse struct
type ListConfigurationsResponse struct {

	// An array of Configurations that are available for the service instance.
	Configurations []Configuration `json:"configurations,omitempty"`
}

// ListEnvironmentsResponse : ListEnvironmentsResponse struct
type ListEnvironmentsResponse struct {

	// An array of [environments] that are available for the service instance.
	Environments []Environment `json:"environments,omitempty"`
}

// MemoryUsage : **Deprecated**: Summary of the memory usage statistics for this environment.
type MemoryUsage struct {

	// **Deprecated**: Number of bytes used in the environment's memory capacity.
	UsedBytes int64 `json:"used_bytes,omitempty"`

	// **Deprecated**: Total number of bytes available in the environment's memory capacity.
	TotalBytes int64 `json:"total_bytes,omitempty"`

	// **Deprecated**: Amount of memory capacity used, in KB or GB format.
	Used string `json:"used,omitempty"`

	// **Deprecated**: Total amount of the environment's memory capacity, in KB or GB format.
	Total string `json:"total,omitempty"`

	// **Deprecated**: Percentage of the environment's memory capacity that is being used.
	PercentUsed float64 `json:"percent_used,omitempty"`
}

// NewTrainingQuery : NewTrainingQuery struct
type NewTrainingQuery struct {

	NaturalLanguageQuery string `json:"natural_language_query,omitempty"`

	Filter string `json:"filter,omitempty"`

	Examples []TrainingExample `json:"examples,omitempty"`
}

// NluEnrichmentCategories : An object that indicates the Categories enrichment will be applied to the specified field.
type NluEnrichmentCategories struct {
}

// NluEnrichmentEmotion : An object specifying the emotion detection enrichment and related parameters.
type NluEnrichmentEmotion struct {

	// When `true`, emotion detection is performed on the entire field.
	Document bool `json:"document,omitempty"`

	// A comma-separated list of target strings that will have any associated emotions detected.
	Targets []string `json:"targets,omitempty"`
}

// NluEnrichmentEntities : An object speficying the Entities enrichment and related parameters.
type NluEnrichmentEntities struct {

	// When `true`, sentiment analysis of entities will be performed on the specified field.
	Sentiment bool `json:"sentiment,omitempty"`

	// When `true`, emotion detection of entities will be performed on the specified field.
	Emotion bool `json:"emotion,omitempty"`

	// The maximum number of entities to extract for each instance of the specified field.
	Limit int64 `json:"limit,omitempty"`

	// When `true`, the number of mentions of each identified entity is recorded. The default is `false`.
	Mentions bool `json:"mentions,omitempty"`

	// When `true`, the types of mentions for each idetifieid entity is recorded. The default is `false`.
	MentionTypes bool `json:"mention_types,omitempty"`

	// When `true`, a list of sentence locations for each instance of each identified entity is recorded. The default is `false`.
	SentenceLocation bool `json:"sentence_location,omitempty"`

	// The enrichement model to use with entity extraction. May be a custom model provided by Watson Knowledge Studio, the public model for use with Knowledge Graph `en-news`, or the default public model `alchemy`.
	Model string `json:"model,omitempty"`
}

// NluEnrichmentFeatures : NluEnrichmentFeatures struct
type NluEnrichmentFeatures struct {

	// An object specifying the Keyword enrichment and related parameters.
	Keywords NluEnrichmentKeywords `json:"keywords,omitempty"`

	// An object speficying the Entities enrichment and related parameters.
	Entities NluEnrichmentEntities `json:"entities,omitempty"`

	// An object specifying the sentiment extraction enrichment and related parameters.
	Sentiment NluEnrichmentSentiment `json:"sentiment,omitempty"`

	// An object specifying the emotion detection enrichment and related parameters.
	Emotion NluEnrichmentEmotion `json:"emotion,omitempty"`

	// An object specifying the categories enrichment and related parameters.
	Categories NluEnrichmentCategories `json:"categories,omitempty"`

	// An object specifiying the semantic roles enrichment and related parameters.
	SemanticRoles NluEnrichmentSemanticRoles `json:"semantic_roles,omitempty"`

	// An object specifying the relations enrichment and related parameters.
	Relations NluEnrichmentRelations `json:"relations,omitempty"`
}

// NluEnrichmentKeywords : An object specifying the Keyword enrichment and related parameters.
type NluEnrichmentKeywords struct {

	// When `true`, sentiment analysis of keywords will be performed on the specified field.
	Sentiment bool `json:"sentiment,omitempty"`

	// When `true`, emotion detection of keywords will be performed on the specified field.
	Emotion bool `json:"emotion,omitempty"`

	// The maximum number of keywords to extract for each instance of the specified field.
	Limit int64 `json:"limit,omitempty"`
}

// NluEnrichmentRelations : An object specifying the relations enrichment and related parameters.
type NluEnrichmentRelations struct {

	// *For use with `natural_language_understanding` enrichments only.* The enrichement model to use with relationship extraction. May be a custom model provided by Watson Knowledge Studio, the public model for use with Knowledge Graph `en-news`, the default is`en-news`.
	Model string `json:"model,omitempty"`
}

// NluEnrichmentSemanticRoles : An object specifiying the semantic roles enrichment and related parameters.
type NluEnrichmentSemanticRoles struct {

	// When `true`, entities are extracted from the identified sentence parts.
	Entities bool `json:"entities,omitempty"`

	// When `true`, keywords are extracted from the identified sentence parts.
	Keywords bool `json:"keywords,omitempty"`

	// The maximum number of semantic roles enrichments to extact from each instance of the specified field.
	Limit int64 `json:"limit,omitempty"`
}

// NluEnrichmentSentiment : An object specifying the sentiment extraction enrichment and related parameters.
type NluEnrichmentSentiment struct {

	// When `true`, sentiment analysis is performed on the entire field.
	Document bool `json:"document,omitempty"`

	// A comma-separated list of target strings that will have any associated sentiment analyzed.
	Targets []string `json:"targets,omitempty"`
}

// NormalizationOperation : NormalizationOperation struct
type NormalizationOperation struct {

	// Identifies what type of operation to perform. **copy** - Copies the value of the **source_field** to the **destination_field** field. If the **destination_field** already exists, then the value of the **source_field** overwrites the original value of the **destination_field**. **move** - Renames (moves) the **source_field** to the **destination_field**. If the **destination_field** already exists, then the value of the **source_field** overwrites the original value of the **destination_field**. Rename is identical to copy, except that the **source_field** is removed after the value has been copied to the **destination_field** (it is the same as a _copy_ followed by a _remove_). **merge** - Merges the value of the **source_field** with the value of the **destination_field**. The **destination_field** is converted into an array if it is not already an array, and the value of the **source_field** is appended to the array. This operation removes the **source_field** after the merge. If the **source_field** does not exist in the current document, then the **destination_field** is still converted into an array (if it is not an array already). This conversion ensures the type for **destination_field** is consistent across all documents. **remove** - Deletes the **source_field** field. The **destination_field** is ignored for this operation. **remove_nulls** - Removes all nested null (blank) field values from the JSON tree. **source_field** and **destination_field** are ignored by this operation because _remove_nulls_ operates on the entire JSON tree. Typically, **remove_nulls** is invoked as the last normalization operation (if it is invoked at all, it can be time-expensive).
	Operation string `json:"operation,omitempty"`

	// The source field for the operation.
	SourceField string `json:"source_field,omitempty"`

	// The destination field for the operation.
	DestinationField string `json:"destination_field,omitempty"`
}

// Notice : A notice produced for the collection.
type Notice struct {

	// Identifies the notice. Many notices might have the same ID. This field exists so that user applications can programmatically identify a notice and take automatic corrective action.
	NoticeID string `json:"notice_id,omitempty"`

	// The creation date of the collection in the format yyyy-MM-dd'T'HH:mm:ss.SSS'Z'.
	Created strfmt.DateTime `json:"created,omitempty"`

	// Unique identifier of the document.
	DocumentID string `json:"document_id,omitempty"`

	// Unique identifier of the query used for relevance training.
	QueryID string `json:"query_id,omitempty"`

	// Severity level of the notice.
	Severity string `json:"severity,omitempty"`

	// Ingestion or training step in which the notice occurred.
	Step string `json:"step,omitempty"`

	// The description of the notice.
	Description string `json:"description,omitempty"`
}

// PdfHeadingDetection : PdfHeadingDetection struct
type PdfHeadingDetection struct {

	Fonts []FontSetting `json:"fonts,omitempty"`
}

// PdfSettings : A list of PDF conversion settings.
type PdfSettings struct {

	Heading PdfHeadingDetection `json:"heading,omitempty"`
}

// QueryAggregation : An aggregation produced by the Discovery service to analyze the input provided.
type QueryAggregation struct {

	// The type of aggregation command used. For example: term, filter, max, min, etc.
	TypeVar string `json:"type_var,omitempty"`

	Results []AggregationResult `json:"results,omitempty"`

	// Number of matching results.
	MatchingResults int64 `json:"matching_results,omitempty"`

	// Aggregations returned by the Discovery service.
	Aggregations []QueryAggregation `json:"aggregations,omitempty"`
}

// QueryEntities : QueryEntities struct
type QueryEntities struct {

	// The entity query feature to perform. Supported features are `disambiguate` and `similar_entities`.
	Feature string `json:"feature,omitempty"`

	// A text string that appears within the entity text field.
	Entity QueryEntitiesEntity `json:"entity,omitempty"`

	// Entity text to provide context for the queried entity and rank based on that association. For example, if you wanted to query the city of London in England your query would look for `London` with the context of `England`.
	Context QueryEntitiesContext `json:"context,omitempty"`

	// The number of results to return. The default is `10`. The maximum is `1000`.
	Count int64 `json:"count,omitempty"`

	// The number of evidence items to return for each result. The default is `0`. The maximum number of evidence items per query is 10,000.
	EvidenceCount int64 `json:"evidence_count,omitempty"`
}

// QueryEntitiesContext : Entity text to provide context for the queried entity and rank based on that association. For example, if you wanted to query the city of London in England your query would look for `London` with the context of `England`.
type QueryEntitiesContext struct {

	// Entity text to provide context for the queried entity and rank based on that association. For example, if you wanted to query the city of London in England your query would look for `London` with the context of `England`.
	Text string `json:"text,omitempty"`
}

// QueryEntitiesEntity : A text string that appears within the entity text field.
type QueryEntitiesEntity struct {

	// Entity text content.
	Text string `json:"text,omitempty"`

	// The type of the specified entity.
	TypeVar string `json:"type_var,omitempty"`
}

// QueryEntitiesResponse : An array of entities resulting from the query.
type QueryEntitiesResponse struct {

	Entities []QueryEntitiesResponseItem `json:"entities,omitempty"`
}

// QueryEntitiesResponseItem : Object containing Entity query response information.
type QueryEntitiesResponseItem struct {

	// Entity text content.
	Text string `json:"text,omitempty"`

	// The type of the result entity.
	TypeVar string `json:"type_var,omitempty"`

	// List of different evidentiary items to support the result.
	Evidence []QueryEvidence `json:"evidence,omitempty"`
}

// QueryEvidence : Description of evidence location supporting Knoweldge Graph query result.
type QueryEvidence struct {

	// The docuemnt ID (as indexed in Discovery) of the evidence location.
	DocumentID string `json:"document_id,omitempty"`

	// The field of the document where the supporting evidence was identified.
	Field string `json:"field,omitempty"`

	// The start location of the evidence in the identified field. This value is inclusive.
	StartOffset int64 `json:"start_offset,omitempty"`

	// The end location of the evidence in the identified field. This value is inclusive.
	EndOffset int64 `json:"end_offset,omitempty"`

	// An array of entity objects that show evidence of the result.
	Entities []QueryEvidenceEntity `json:"entities,omitempty"`
}

// QueryEvidenceEntity : Entity description and location within evidence field.
type QueryEvidenceEntity struct {

	// The entity type for this entity. Possible types vary based on model used.
	TypeVar string `json:"type_var,omitempty"`

	// The original text of this entity as found in the evidence field.
	Text string `json:"text,omitempty"`

	// The start location of the entity text in the identified field. This value is inclusive.
	StartOffset int64 `json:"start_offset,omitempty"`

	// The end location of the entity text in the identified field. This value is exclusive.
	EndOffset int64 `json:"end_offset,omitempty"`
}

// QueryFilterType : QueryFilterType struct
type QueryFilterType struct {

	// A comma-separated list of types to exclude.
	Exclude []string `json:"exclude,omitempty"`

	// A comma-separated list of types to include. All other types are excluded.
	Include []string `json:"include,omitempty"`
}

// QueryNoticesResponse : QueryNoticesResponse struct
type QueryNoticesResponse struct {

	MatchingResults int64 `json:"matching_results,omitempty"`

	Results []QueryNoticesResult `json:"results,omitempty"`

	Aggregations []QueryAggregation `json:"aggregations,omitempty"`

	Passages []QueryPassages `json:"passages,omitempty"`

	DuplicatesRemoved int64 `json:"duplicates_removed,omitempty"`
}

// QueryNoticesResult : QueryNoticesResult struct
type QueryNoticesResult struct {

	// The unique identifier of the document.
	ID string `json:"id,omitempty"`

	// *Deprecated* This field is now part of the **result_metadata** object.
	Score float64 `json:"score,omitempty"`

	// Metadata of the document.
	Metadata interface{} `json:"metadata,omitempty"`

	// The collection ID of the collection containing the document for this result.
	CollectionID string `json:"collection_id,omitempty"`

	// Metadata of the query result.
	ResultMetadata QueryResultMetadata `json:"result_metadata,omitempty"`

	// The internal status code returned by the ingestion subsystem indicating the overall result of ingesting the source document.
	Code int64 `json:"code,omitempty"`

	// Name of the original source file (if available).
	Filename string `json:"filename,omitempty"`

	// The type of the original source file.
	FileType string `json:"file_type,omitempty"`

	// The SHA-1 hash of the original source file (formatted as a hexadecimal string).
	Sha1 string `json:"sha1,omitempty"`

	// Array of notices for the document.
	Notices []Notice `json:"notices,omitempty"`
}

// QueryPassages : QueryPassages struct
type QueryPassages struct {

	// The unique identifier of the document from which the passage has been extracted.
	DocumentID string `json:"document_id,omitempty"`

	// The confidence score of the passages's analysis. A higher score indicates greater confidence.
	PassageScore float64 `json:"passage_score,omitempty"`

	// The content of the extracted passage.
	PassageText string `json:"passage_text,omitempty"`

	// The position of the first character of the extracted passage in the originating field.
	StartOffset int64 `json:"start_offset,omitempty"`

	// The position of the last character of the extracted passage in the originating field.
	EndOffset int64 `json:"end_offset,omitempty"`

	// The label of the field from which the passage has been extracted.
	Field string `json:"field,omitempty"`
}

// QueryRelations : A respresentation of a relationship query.
type QueryRelations struct {

	// An array of entities to find relationships for.
	Entities []QueryRelationsEntity `json:"entities,omitempty"`

	// Entity text to provide context for the queried entity and rank based on that association. For example, if you wanted to query the city of London in England your query would look for `London` with the context of `England`.
	Context QueryEntitiesContext `json:"context,omitempty"`

	// The sorting method for the relationships, can be `score` or `frequency`. `frequency` is the number of unique times each entity is identified. The default is `score`.
	Sort string `json:"sort,omitempty"`

	// Filters to apply to the relationship query.
	Filter QueryRelationsFilter `json:"filter,omitempty"`

	// The number of results to return. The default is `10`. The maximum is `1000`.
	Count int64 `json:"count,omitempty"`

	// The number of evidence items to return for each result. The default is `0`. The maximum number of evidence items per query is 10,000.
	EvidenceCount int64 `json:"evidence_count,omitempty"`
}

// QueryRelationsArgument : QueryRelationsArgument struct
type QueryRelationsArgument struct {

	Entities []QueryEntitiesEntity `json:"entities,omitempty"`
}

// QueryRelationsEntity : QueryRelationsEntity struct
type QueryRelationsEntity struct {

	// Entity text content.
	Text string `json:"text,omitempty"`

	// The type of the specified entity.
	TypeVar string `json:"type_var,omitempty"`

	// If false, implicit querying is performed. The default is `false`.
	Exact bool `json:"exact,omitempty"`
}

// QueryRelationsFilter : QueryRelationsFilter struct
type QueryRelationsFilter struct {

	// A list of relation types to include or exclude from the query.
	RelationTypes QueryFilterType `json:"relation_types,omitempty"`

	// A list of entity types to include or exclude from the query.
	EntityTypes QueryFilterType `json:"entity_types,omitempty"`

	// A comma-separated list of document IDs to include in the query.
	DocumentIds []string `json:"document_ids,omitempty"`
}

// QueryRelationsRelationship : QueryRelationsRelationship struct
type QueryRelationsRelationship struct {

	// The identified relationship type.
	TypeVar string `json:"type_var,omitempty"`

	// The number of times the relationship is mentioned.
	Frequency int64 `json:"frequency,omitempty"`

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

	MatchingResults int64 `json:"matching_results,omitempty"`

	Results []QueryResult `json:"results,omitempty"`

	Aggregations []QueryAggregation `json:"aggregations,omitempty"`

	Passages []QueryPassages `json:"passages,omitempty"`

	DuplicatesRemoved int64 `json:"duplicates_removed,omitempty"`

	// The session token for this query. The session token can be used to add events associated with this query to the query and event log.
	SessionToken string `json:"session_token,omitempty"`
}

// QueryResult : QueryResult struct
type QueryResult struct {

	// The unique identifier of the document.
	ID string `json:"id,omitempty"`

	// *Deprecated* This field is now part of the **result_metadata** object.
	Score float64 `json:"score,omitempty"`

	// Metadata of the document.
	Metadata interface{} `json:"metadata,omitempty"`

	// The collection ID of the collection containing the document for this result.
	CollectionID string `json:"collection_id,omitempty"`

	// Metadata of the query result.
	ResultMetadata QueryResultMetadata `json:"result_metadata,omitempty"`
}

// QueryResultMetadata : Metadata of a query result.
type QueryResultMetadata struct {

	// The raw score of the result. A higher score indicates a greater match to the query parameters.
	Score float64 `json:"score,omitempty"`

	// The confidence score of the result's analysis. A higher score indicates greater confidence.
	Confidence float64 `json:"confidence,omitempty"`
}

// SegmentSettings : A list of Document Segmentation settings.
type SegmentSettings struct {

	// Enables/disables the Document Segmentation feature.
	Enabled bool `json:"enabled,omitempty"`

	// Defines the heading level that splits into document segments. Valid values are h1, h2, h3, h4, h5, h6.
	SelectorTags []string `json:"selector_tags,omitempty"`
}

// Source : Object containing source parameters for the configuration.
type Source struct {

	// The type of source to connect to. -  `box` indicates the configuration is to connect an instance of Enterprise Box. -  `salesforce` indicates the configuration is to connect to Salesforce. -  `sharepoint` indicates the configuration is to connect to Microsoft SharePoint Online.
	TypeVar string `json:"type_var,omitempty"`

	// The **credential_id** of the credentials to use to connect to the source. Credentials are defined using the **credentials** method. The **source_type** of the credentials used must match the **type** field specified in this object.
	CredentialID string `json:"credential_id,omitempty"`

	// Object containing the schedule information for the source.
	Schedule SourceSchedule `json:"schedule,omitempty"`

	// The **options** object defines which items to crawl from the source system.
	Options SourceOptions `json:"options,omitempty"`
}

// SourceOptions : The **options** object defines which items to crawl from the source system.
type SourceOptions struct {

	// Array of folders to crawl from the Box source. Only valid, and required, when the **type** field of the **source** object is set to `box`.
	Folders []SourceOptionsFolder `json:"folders,omitempty"`

	// Array of Salesforce document object types to crawl from the Salesforce source. Only valid, and required, when the **type** field of the **source** object is set to `salesforce`.
	Objects []SourceOptionsObject `json:"objects,omitempty"`

	// Array of Microsoft SharePointoint Online site collections to crawl from the SharePoint source. Only valid and required when the **type** field of the **source** object is set to `sharepoint`.
	SiteCollections []SourceOptionsSiteColl `json:"site_collections,omitempty"`
}

// SourceOptionsFolder : Object that defines a box folder to crawl with this configuration.
type SourceOptionsFolder struct {

	// The Box user ID of the user who owns the folder to crawl.
	OwnerUserID string `json:"owner_user_id"`

	// The Box folder ID of the folder to crawl.
	FolderID string `json:"folder_id"`

	// The maximum number of documents to crawl for this folder. By default, all documents in the folder are crawled.
	Limit int64 `json:"limit,omitempty"`
}

// SourceOptionsObject : Object that defines a Salesforce document object type crawl with this configuration.
type SourceOptionsObject struct {

	// The name of the Salesforce document object to crawl. For example, `case`.
	Name string `json:"name"`

	// The maximum number of documents to crawl for this document object. By default, all documents in the document object are crawled.
	Limit int64 `json:"limit,omitempty"`
}

// SourceOptionsSiteColl : Object that defines a Microsoft SharePoint site collection to crawl with this configuration.
type SourceOptionsSiteColl struct {

	// The Microsoft SharePoint Online site collection path to crawl. The path must be be relative to the **organization_url** that was specified in the credentials associated with this source configuration.
	SiteCollectionPath string `json:"site_collection_path"`

	// The maximum number of documents to crawl for this site collection. By default, all documents in the site collection are crawled.
	Limit int64 `json:"limit,omitempty"`
}

// SourceSchedule : Object containing the schedule information for the source.
type SourceSchedule struct {

	// When `true`, the source is re-crawled based on the **frequency** field in this object. When `false` the source is not re-crawled; When `false` and connecting to Salesforce the source is crawled annually.
	Enabled bool `json:"enabled,omitempty"`

	// The time zone to base source crawl times on. Possible values correspond to the IANA (Internet Assigned Numbers Authority) time zones list.
	TimeZone string `json:"time_zone,omitempty"`

	// The crawl schedule in the specified **time_zone**. -  `daily`: Runs every day between 00:00 and 06:00. -  `weekly`: Runs every week on Sunday between 00:00 and 06:00. -  `monthly`: Runs the on the first Sunday of every month between 00:00 and 06:00.
	Frequency string `json:"frequency,omitempty"`
}

// SourceStatus : Object containing source crawl status information.
type SourceStatus struct {

	// The current status of the source crawl for this collection. This field returns `not_configured` if the default configuration for this source does not have a **source** object defined. -  `running` indicates that a crawl to fetch more documents is in progress. -  `complete` indicates that the crawl has completed with no errors. -  `complete_with_notices` indicates that some notices were generated during the crawl. Notices can be checked by using the **notices** query method. -  `stopped` indicates that the crawl has stopped but is not complete.
	Status string `json:"status,omitempty"`

	// Date in UTC format indicating when the last crawl was attempted. If `null`, no crawl was completed.
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`
}

// TestDocument : TestDocument struct
type TestDocument struct {

	// The unique identifier for the configuration.
	ConfigurationID string `json:"configuration_id,omitempty"`

	// Status of the preview operation.
	Status string `json:"status,omitempty"`

	// The number of 10-kB chunks of field data that were enriched. This can be used to estimate the cost of running a real ingestion.
	EnrichedFieldUnits int64 `json:"enriched_field_units,omitempty"`

	// Format of the test document.
	OriginalMediaType string `json:"original_media_type,omitempty"`

	// An array of objects that describe each step in the preview process.
	Snapshots []DocumentSnapshot `json:"snapshots,omitempty"`

	// An array of notice messages about the preview operation.
	Notices []Notice `json:"notices,omitempty"`
}

// TopHitsResults : TopHitsResults struct
type TopHitsResults struct {

	// Number of matching results.
	MatchingResults int64 `json:"matching_results,omitempty"`

	// Top results returned by the aggregation.
	Hits []QueryResult `json:"hits,omitempty"`
}

// TrainingDataSet : TrainingDataSet struct
type TrainingDataSet struct {

	EnvironmentID string `json:"environment_id,omitempty"`

	CollectionID string `json:"collection_id,omitempty"`

	Queries []TrainingQuery `json:"queries,omitempty"`
}

// TrainingExample : TrainingExample struct
type TrainingExample struct {

	DocumentID string `json:"document_id,omitempty"`

	CrossReference string `json:"cross_reference,omitempty"`

	Relevance int64 `json:"relevance,omitempty"`
}

// TrainingExampleList : TrainingExampleList struct
type TrainingExampleList struct {

	Examples []TrainingExample `json:"examples,omitempty"`
}

// TrainingExamplePatch : TrainingExamplePatch struct
type TrainingExamplePatch struct {

	CrossReference string `json:"cross_reference,omitempty"`

	Relevance int64 `json:"relevance,omitempty"`
}

// TrainingQuery : TrainingQuery struct
type TrainingQuery struct {

	QueryID string `json:"query_id,omitempty"`

	NaturalLanguageQuery string `json:"natural_language_query,omitempty"`

	Filter string `json:"filter,omitempty"`

	Examples []TrainingExample `json:"examples,omitempty"`
}

// TrainingStatus : TrainingStatus struct
type TrainingStatus struct {

	TotalExamples int64 `json:"total_examples,omitempty"`

	Available bool `json:"available,omitempty"`

	Processing bool `json:"processing,omitempty"`

	MinimumQueriesAdded bool `json:"minimum_queries_added,omitempty"`

	MinimumExamplesAdded bool `json:"minimum_examples_added,omitempty"`

	SufficientLabelDiversity bool `json:"sufficient_label_diversity,omitempty"`

	Notices int64 `json:"notices,omitempty"`

	SuccessfullyTrained strfmt.DateTime `json:"successfully_trained,omitempty"`

	DataUpdated strfmt.DateTime `json:"data_updated,omitempty"`
}

// UpdateCollectionRequest : UpdateCollectionRequest struct
type UpdateCollectionRequest struct {

	// The name of the collection.
	Name string `json:"name"`

	// A description of the collection.
	Description string `json:"description,omitempty"`

	// The ID of the configuration in which the collection is to be updated.
	ConfigurationID string `json:"configuration_id,omitempty"`
}

// UpdateEnvironmentRequest : UpdateEnvironmentRequest struct
type UpdateEnvironmentRequest struct {

	// Name that identifies the environment.
	Name string `json:"name,omitempty"`

	// Description of the environment.
	Description string `json:"description,omitempty"`
}

// WordHeadingDetection : WordHeadingDetection struct
type WordHeadingDetection struct {

	Fonts []FontSetting `json:"fonts,omitempty"`

	Styles []WordStyle `json:"styles,omitempty"`
}

// WordSettings : A list of Word conversion settings.
type WordSettings struct {

	Heading WordHeadingDetection `json:"heading,omitempty"`
}

// WordStyle : WordStyle struct
type WordStyle struct {

	Level int64 `json:"level,omitempty"`

	Names []string `json:"names,omitempty"`
}

// XPathPatterns : XPathPatterns struct
type XPathPatterns struct {

	Xpaths []string `json:"xpaths,omitempty"`
}

// Calculation : Calculation struct
type Calculation struct {

	// The field where the aggregation is located in the document.
	Field string `json:"field,omitempty"`

	// Value of the aggregation.
	Value float64 `json:"value,omitempty"`
}

// Filter : Filter struct
type Filter struct {

	// The match the aggregated results queried for.
	Match string `json:"match,omitempty"`
}

// Histogram : Histogram struct
type Histogram struct {

	// The field where the aggregation is located in the document.
	Field string `json:"field,omitempty"`

	// Interval of the aggregation. (For 'histogram' type).
	Interval int64 `json:"interval,omitempty"`
}

// Nested : Nested struct
type Nested struct {

	// The area of the results the aggregation was restricted to.
	Path string `json:"path,omitempty"`
}

// Term : Term struct
type Term struct {

	// The field where the aggregation is located in the document.
	Field string `json:"field,omitempty"`

	Count int64 `json:"count,omitempty"`
}

// Timeslice : Timeslice struct
type Timeslice struct {

	// The field where the aggregation is located in the document.
	Field string `json:"field,omitempty"`

	// Interval of the aggregation. Valid date interval values are second/seconds minute/minutes, hour/hours, day/days, week/weeks, month/months, and year/years.
	Interval string `json:"interval,omitempty"`

	// Used to inducate that anomaly detection should be performed. Anomaly detection is used to locate unusual datapoints within a time series.
	Anomaly bool `json:"anomaly,omitempty"`
}

// TopHits : TopHits struct
type TopHits struct {

	// Number of top hits returned by the aggregation.
	Size int64 `json:"size,omitempty"`

	Hits TopHitsResults `json:"hits,omitempty"`
}
