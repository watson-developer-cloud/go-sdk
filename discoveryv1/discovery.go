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
func (discovery *DiscoveryV1) CreateEnvironment(options *CreateEnvironmentOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    body := map[string]interface{}{}
    body["name"] = options.Name
    if options.IsDescriptionSet {
        body["description"] = options.Description
    }
    if options.IsSizeSet {
        body["size"] = options.Size
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
func (discovery *DiscoveryV1) DeleteEnvironment(options *DeleteEnvironmentOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

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
func (discovery *DiscoveryV1) GetEnvironment(options *GetEnvironmentOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

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
func (discovery *DiscoveryV1) ListEnvironments(options *ListEnvironmentsOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    if options.IsNameSet {
        request.Query("name=" + fmt.Sprint(options.Name))
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
func (discovery *DiscoveryV1) ListFields(options *ListFieldsOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/fields"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("collection_ids=" + fmt.Sprint(options.CollectionIds))

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
func (discovery *DiscoveryV1) UpdateEnvironment(options *UpdateEnvironmentOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    request := req.New().Put(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

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
func (discovery *DiscoveryV1) CreateConfiguration(options *CreateConfigurationOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/configurations"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

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
    if options.IsConversionsSet {
        body["conversions"] = options.Conversions
    }
    if options.IsEnrichmentsSet {
        body["enrichments"] = options.Enrichments
    }
    if options.IsNormalizationsSet {
        body["normalizations"] = options.Normalizations
    }
    if options.IsSourceSet {
        body["source"] = options.Source
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
func (discovery *DiscoveryV1) DeleteConfiguration(options *DeleteConfigurationOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/configurations/{configuration_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    path = strings.Replace(path, "{configuration_id}", options.ConfigurationID, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

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
func (discovery *DiscoveryV1) GetConfiguration(options *GetConfigurationOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/configurations/{configuration_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    path = strings.Replace(path, "{configuration_id}", options.ConfigurationID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

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
func (discovery *DiscoveryV1) ListConfigurations(options *ListConfigurationsOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/configurations"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    if options.IsNameSet {
        request.Query("name=" + fmt.Sprint(options.Name))
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
func (discovery *DiscoveryV1) UpdateConfiguration(options *UpdateConfigurationOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/configurations/{configuration_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    path = strings.Replace(path, "{configuration_id}", options.ConfigurationID, 1)
    request := req.New().Put(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

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
    if options.IsConversionsSet {
        body["conversions"] = options.Conversions
    }
    if options.IsEnrichmentsSet {
        body["enrichments"] = options.Enrichments
    }
    if options.IsNormalizationsSet {
        body["normalizations"] = options.Normalizations
    }
    if options.IsSourceSet {
        body["source"] = options.Source
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
func (discovery *DiscoveryV1) TestConfigurationInEnvironment(options *TestConfigurationInEnvironmentOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/preview"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    if options.IsStepSet {
        request.Query("step=" + fmt.Sprint(options.Step))
    }
    if options.IsConfigurationIDSet {
        request.Query("configuration_id=" + fmt.Sprint(options.ConfigurationID))
    }
    request.Type("multipart")
    form := map[string]interface{}{}
    if options.IsConfigurationSet {
        form["configuration"] = options.Configuration
    }
    if options.IsFileSet {
        request.SendFile(options.File, "", "file")
    }
    if options.IsMetadataSet {
        form["metadata"] = options.Metadata
    }
    request.Send(form)

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
func (discovery *DiscoveryV1) CreateCollection(options *CreateCollectionOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    body := map[string]interface{}{}
    body["name"] = options.Name
    if options.IsDescriptionSet {
        body["description"] = options.Description
    }
    if options.IsConfigurationIDSet {
        body["configuration_id"] = options.ConfigurationID
    }
    if options.IsLanguageSet {
        body["language"] = options.Language
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
func (discovery *DiscoveryV1) DeleteCollection(options *DeleteCollectionOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    path = strings.Replace(path, "{collection_id}", options.CollectionID, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

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
func (discovery *DiscoveryV1) GetCollection(options *GetCollectionOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    path = strings.Replace(path, "{collection_id}", options.CollectionID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

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
func (discovery *DiscoveryV1) ListCollectionFields(options *ListCollectionFieldsOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/fields"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    path = strings.Replace(path, "{collection_id}", options.CollectionID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

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
func (discovery *DiscoveryV1) ListCollections(options *ListCollectionsOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    if options.IsNameSet {
        request.Query("name=" + fmt.Sprint(options.Name))
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
func (discovery *DiscoveryV1) UpdateCollection(options *UpdateCollectionOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    path = strings.Replace(path, "{collection_id}", options.CollectionID, 1)
    request := req.New().Put(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

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
    if options.IsConfigurationIDSet {
        body["configuration_id"] = options.ConfigurationID
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
func (discovery *DiscoveryV1) CreateExpansions(options *CreateExpansionsOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/expansions"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    path = strings.Replace(path, "{collection_id}", options.CollectionID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    body := map[string]interface{}{}
    if options.IsExpansionsSet {
        body["expansions"] = options.Expansions
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
func (discovery *DiscoveryV1) DeleteExpansions(options *DeleteExpansionsOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/expansions"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    path = strings.Replace(path, "{collection_id}", options.CollectionID, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

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
func (discovery *DiscoveryV1) ListExpansions(options *ListExpansionsOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/expansions"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    path = strings.Replace(path, "{collection_id}", options.CollectionID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

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
func (discovery *DiscoveryV1) AddDocument(options *AddDocumentOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/documents"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    path = strings.Replace(path, "{collection_id}", options.CollectionID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    request.Type("multipart")
    form := map[string]interface{}{}
    if options.IsFileSet {
        request.SendFile(options.File, "", "file")
    }
    if options.IsMetadataSet {
        form["metadata"] = options.Metadata
    }
    request.Send(form)

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
func (discovery *DiscoveryV1) DeleteDocument(options *DeleteDocumentOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/documents/{document_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    path = strings.Replace(path, "{collection_id}", options.CollectionID, 1)
    path = strings.Replace(path, "{document_id}", options.DocumentID, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

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
func (discovery *DiscoveryV1) GetDocumentStatus(options *GetDocumentStatusOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/documents/{document_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    path = strings.Replace(path, "{collection_id}", options.CollectionID, 1)
    path = strings.Replace(path, "{document_id}", options.DocumentID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

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
func (discovery *DiscoveryV1) UpdateDocument(options *UpdateDocumentOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/documents/{document_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    path = strings.Replace(path, "{collection_id}", options.CollectionID, 1)
    path = strings.Replace(path, "{document_id}", options.DocumentID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    request.Type("multipart")
    form := map[string]interface{}{}
    if options.IsFileSet {
        request.SendFile(options.File, "", "file")
    }
    if options.IsMetadataSet {
        form["metadata"] = options.Metadata
    }
    request.Send(form)

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
func (discovery *DiscoveryV1) FederatedQuery(options *FederatedQueryOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/query"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("collection_ids=" + fmt.Sprint(options.CollectionIds))
    if options.IsFilterSet {
        request.Query("filter=" + fmt.Sprint(options.Filter))
    }
    if options.IsQuerySet {
        request.Query("query=" + fmt.Sprint(options.Query))
    }
    if options.IsNaturalLanguageQuerySet {
        request.Query("natural_language_query=" + fmt.Sprint(options.NaturalLanguageQuery))
    }
    if options.IsAggregationSet {
        request.Query("aggregation=" + fmt.Sprint(options.Aggregation))
    }
    if options.IsCountSet {
        request.Query("count=" + fmt.Sprint(options.Count))
    }
    if options.IsReturnFieldsSet {
        request.Query("return=" + fmt.Sprint(options.ReturnFields))
    }
    if options.IsOffsetSet {
        request.Query("offset=" + fmt.Sprint(options.Offset))
    }
    if options.IsSortSet {
        request.Query("sort=" + fmt.Sprint(options.Sort))
    }
    if options.IsHighlightSet {
        request.Query("highlight=" + fmt.Sprint(options.Highlight))
    }
    if options.IsDeduplicateSet {
        request.Query("deduplicate=" + fmt.Sprint(options.Deduplicate))
    }
    if options.IsDeduplicateFieldSet {
        request.Query("deduplicate.field=" + fmt.Sprint(options.DeduplicateField))
    }
    if options.IsSimilarSet {
        request.Query("similar=" + fmt.Sprint(options.Similar))
    }
    if options.IsSimilarDocumentIdsSet {
        request.Query("similar.document_ids=" + fmt.Sprint(options.SimilarDocumentIds))
    }
    if options.IsSimilarFieldsSet {
        request.Query("similar.fields=" + fmt.Sprint(options.SimilarFields))
    }
    if options.IsPassagesSet {
        request.Query("passages=" + fmt.Sprint(options.Passages))
    }
    if options.IsPassagesFieldsSet {
        request.Query("passages.fields=" + fmt.Sprint(options.PassagesFields))
    }
    if options.IsPassagesCountSet {
        request.Query("passages.count=" + fmt.Sprint(options.PassagesCount))
    }
    if options.IsPassagesCharactersSet {
        request.Query("passages.characters=" + fmt.Sprint(options.PassagesCharacters))
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
func (discovery *DiscoveryV1) FederatedQueryNotices(options *FederatedQueryNoticesOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/notices"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("collection_ids=" + fmt.Sprint(options.CollectionIds))
    if options.IsFilterSet {
        request.Query("filter=" + fmt.Sprint(options.Filter))
    }
    if options.IsQuerySet {
        request.Query("query=" + fmt.Sprint(options.Query))
    }
    if options.IsNaturalLanguageQuerySet {
        request.Query("natural_language_query=" + fmt.Sprint(options.NaturalLanguageQuery))
    }
    if options.IsAggregationSet {
        request.Query("aggregation=" + fmt.Sprint(options.Aggregation))
    }
    if options.IsCountSet {
        request.Query("count=" + fmt.Sprint(options.Count))
    }
    if options.IsReturnFieldsSet {
        request.Query("return=" + fmt.Sprint(options.ReturnFields))
    }
    if options.IsOffsetSet {
        request.Query("offset=" + fmt.Sprint(options.Offset))
    }
    if options.IsSortSet {
        request.Query("sort=" + fmt.Sprint(options.Sort))
    }
    if options.IsHighlightSet {
        request.Query("highlight=" + fmt.Sprint(options.Highlight))
    }
    if options.IsDeduplicateFieldSet {
        request.Query("deduplicate.field=" + fmt.Sprint(options.DeduplicateField))
    }
    if options.IsSimilarSet {
        request.Query("similar=" + fmt.Sprint(options.Similar))
    }
    if options.IsSimilarDocumentIdsSet {
        request.Query("similar.document_ids=" + fmt.Sprint(options.SimilarDocumentIds))
    }
    if options.IsSimilarFieldsSet {
        request.Query("similar.fields=" + fmt.Sprint(options.SimilarFields))
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
func (discovery *DiscoveryV1) Query(options *QueryOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/query"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    path = strings.Replace(path, "{collection_id}", options.CollectionID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    if options.IsLoggingOptOutSet {
        request.Set("X-Watson-Logging-Opt-Out", fmt.Sprint(options.LoggingOptOut))
    }
    request.Query("version=" + creds.Version)
    if options.IsFilterSet {
        request.Query("filter=" + fmt.Sprint(options.Filter))
    }
    if options.IsQuerySet {
        request.Query("query=" + fmt.Sprint(options.Query))
    }
    if options.IsNaturalLanguageQuerySet {
        request.Query("natural_language_query=" + fmt.Sprint(options.NaturalLanguageQuery))
    }
    if options.IsPassagesSet {
        request.Query("passages=" + fmt.Sprint(options.Passages))
    }
    if options.IsAggregationSet {
        request.Query("aggregation=" + fmt.Sprint(options.Aggregation))
    }
    if options.IsCountSet {
        request.Query("count=" + fmt.Sprint(options.Count))
    }
    if options.IsReturnFieldsSet {
        request.Query("return=" + fmt.Sprint(options.ReturnFields))
    }
    if options.IsOffsetSet {
        request.Query("offset=" + fmt.Sprint(options.Offset))
    }
    if options.IsSortSet {
        request.Query("sort=" + fmt.Sprint(options.Sort))
    }
    if options.IsHighlightSet {
        request.Query("highlight=" + fmt.Sprint(options.Highlight))
    }
    if options.IsPassagesFieldsSet {
        request.Query("passages.fields=" + fmt.Sprint(options.PassagesFields))
    }
    if options.IsPassagesCountSet {
        request.Query("passages.count=" + fmt.Sprint(options.PassagesCount))
    }
    if options.IsPassagesCharactersSet {
        request.Query("passages.characters=" + fmt.Sprint(options.PassagesCharacters))
    }
    if options.IsDeduplicateSet {
        request.Query("deduplicate=" + fmt.Sprint(options.Deduplicate))
    }
    if options.IsDeduplicateFieldSet {
        request.Query("deduplicate.field=" + fmt.Sprint(options.DeduplicateField))
    }
    if options.IsSimilarSet {
        request.Query("similar=" + fmt.Sprint(options.Similar))
    }
    if options.IsSimilarDocumentIdsSet {
        request.Query("similar.document_ids=" + fmt.Sprint(options.SimilarDocumentIds))
    }
    if options.IsSimilarFieldsSet {
        request.Query("similar.fields=" + fmt.Sprint(options.SimilarFields))
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
func (discovery *DiscoveryV1) QueryEntities(options *QueryEntitiesOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/query_entities"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    path = strings.Replace(path, "{collection_id}", options.CollectionID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    body := map[string]interface{}{}
    if options.IsFeatureSet {
        body["feature"] = options.Feature
    }
    if options.IsEntitySet {
        body["entity"] = options.Entity
    }
    if options.IsContextSet {
        body["context"] = options.Context
    }
    if options.IsCountSet {
        body["count"] = options.Count
    }
    if options.IsEvidenceCountSet {
        body["evidence_count"] = options.EvidenceCount
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
func (discovery *DiscoveryV1) QueryNotices(options *QueryNoticesOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/notices"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    path = strings.Replace(path, "{collection_id}", options.CollectionID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    if options.IsFilterSet {
        request.Query("filter=" + fmt.Sprint(options.Filter))
    }
    if options.IsQuerySet {
        request.Query("query=" + fmt.Sprint(options.Query))
    }
    if options.IsNaturalLanguageQuerySet {
        request.Query("natural_language_query=" + fmt.Sprint(options.NaturalLanguageQuery))
    }
    if options.IsPassagesSet {
        request.Query("passages=" + fmt.Sprint(options.Passages))
    }
    if options.IsAggregationSet {
        request.Query("aggregation=" + fmt.Sprint(options.Aggregation))
    }
    if options.IsCountSet {
        request.Query("count=" + fmt.Sprint(options.Count))
    }
    if options.IsReturnFieldsSet {
        request.Query("return=" + fmt.Sprint(options.ReturnFields))
    }
    if options.IsOffsetSet {
        request.Query("offset=" + fmt.Sprint(options.Offset))
    }
    if options.IsSortSet {
        request.Query("sort=" + fmt.Sprint(options.Sort))
    }
    if options.IsHighlightSet {
        request.Query("highlight=" + fmt.Sprint(options.Highlight))
    }
    if options.IsPassagesFieldsSet {
        request.Query("passages.fields=" + fmt.Sprint(options.PassagesFields))
    }
    if options.IsPassagesCountSet {
        request.Query("passages.count=" + fmt.Sprint(options.PassagesCount))
    }
    if options.IsPassagesCharactersSet {
        request.Query("passages.characters=" + fmt.Sprint(options.PassagesCharacters))
    }
    if options.IsDeduplicateFieldSet {
        request.Query("deduplicate.field=" + fmt.Sprint(options.DeduplicateField))
    }
    if options.IsSimilarSet {
        request.Query("similar=" + fmt.Sprint(options.Similar))
    }
    if options.IsSimilarDocumentIdsSet {
        request.Query("similar.document_ids=" + fmt.Sprint(options.SimilarDocumentIds))
    }
    if options.IsSimilarFieldsSet {
        request.Query("similar.fields=" + fmt.Sprint(options.SimilarFields))
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
func (discovery *DiscoveryV1) QueryRelations(options *QueryRelationsOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/query_relations"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    path = strings.Replace(path, "{collection_id}", options.CollectionID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    body := map[string]interface{}{}
    if options.IsEntitiesSet {
        body["entities"] = options.Entities
    }
    if options.IsContextSet {
        body["context"] = options.Context
    }
    if options.IsSortSet {
        body["sort"] = options.Sort
    }
    if options.IsFilterSet {
        body["filter"] = options.Filter
    }
    if options.IsCountSet {
        body["count"] = options.Count
    }
    if options.IsEvidenceCountSet {
        body["evidence_count"] = options.EvidenceCount
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
func (discovery *DiscoveryV1) AddTrainingData(options *AddTrainingDataOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/training_data"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    path = strings.Replace(path, "{collection_id}", options.CollectionID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    body := map[string]interface{}{}
    if options.IsNaturalLanguageQuerySet {
        body["natural_language_query"] = options.NaturalLanguageQuery
    }
    if options.IsFilterSet {
        body["filter"] = options.Filter
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
func (discovery *DiscoveryV1) CreateTrainingExample(options *CreateTrainingExampleOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}/examples"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    path = strings.Replace(path, "{collection_id}", options.CollectionID, 1)
    path = strings.Replace(path, "{query_id}", options.QueryID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    body := map[string]interface{}{}
    if options.IsDocumentIDSet {
        body["document_id"] = options.DocumentID
    }
    if options.IsCrossReferenceSet {
        body["cross_reference"] = options.CrossReference
    }
    if options.IsRelevanceSet {
        body["relevance"] = options.Relevance
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
func (discovery *DiscoveryV1) DeleteAllTrainingData(options *DeleteAllTrainingDataOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/training_data"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    path = strings.Replace(path, "{collection_id}", options.CollectionID, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

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
func (discovery *DiscoveryV1) DeleteTrainingData(options *DeleteTrainingDataOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    path = strings.Replace(path, "{collection_id}", options.CollectionID, 1)
    path = strings.Replace(path, "{query_id}", options.QueryID, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

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
func (discovery *DiscoveryV1) DeleteTrainingExample(options *DeleteTrainingExampleOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}/examples/{example_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    path = strings.Replace(path, "{collection_id}", options.CollectionID, 1)
    path = strings.Replace(path, "{query_id}", options.QueryID, 1)
    path = strings.Replace(path, "{example_id}", options.ExampleID, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

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
func (discovery *DiscoveryV1) GetTrainingData(options *GetTrainingDataOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    path = strings.Replace(path, "{collection_id}", options.CollectionID, 1)
    path = strings.Replace(path, "{query_id}", options.QueryID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

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
func (discovery *DiscoveryV1) GetTrainingExample(options *GetTrainingExampleOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}/examples/{example_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    path = strings.Replace(path, "{collection_id}", options.CollectionID, 1)
    path = strings.Replace(path, "{query_id}", options.QueryID, 1)
    path = strings.Replace(path, "{example_id}", options.ExampleID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

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
func (discovery *DiscoveryV1) ListTrainingData(options *ListTrainingDataOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/training_data"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    path = strings.Replace(path, "{collection_id}", options.CollectionID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

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
func (discovery *DiscoveryV1) ListTrainingExamples(options *ListTrainingExamplesOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}/examples"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    path = strings.Replace(path, "{collection_id}", options.CollectionID, 1)
    path = strings.Replace(path, "{query_id}", options.QueryID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

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
func (discovery *DiscoveryV1) UpdateTrainingExample(options *UpdateTrainingExampleOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}/examples/{example_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    path = strings.Replace(path, "{collection_id}", options.CollectionID, 1)
    path = strings.Replace(path, "{query_id}", options.QueryID, 1)
    path = strings.Replace(path, "{example_id}", options.ExampleID, 1)
    request := req.New().Put(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    body := map[string]interface{}{}
    if options.IsCrossReferenceSet {
        body["cross_reference"] = options.CrossReference
    }
    if options.IsRelevanceSet {
        body["relevance"] = options.Relevance
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
func (discovery *DiscoveryV1) DeleteUserData(options *DeleteUserDataOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/user_data"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    request := req.New().Delete(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
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
func (discovery *DiscoveryV1) CreateCredentials(options *CreateCredentialsOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/credentials"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    body := map[string]interface{}{}
    if options.IsSourceTypeSet {
        body["source_type"] = options.SourceType
    }
    if options.IsCredentialDetailsSet {
        body["credential_details"] = options.CredentialDetails
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
func (discovery *DiscoveryV1) DeleteCredentials(options *DeleteCredentialsOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/credentials/{credential_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    path = strings.Replace(path, "{credential_id}", options.CredentialID, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

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
func (discovery *DiscoveryV1) GetCredentials(options *GetCredentialsOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/credentials/{credential_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    path = strings.Replace(path, "{credential_id}", options.CredentialID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

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
func (discovery *DiscoveryV1) ListCredentials(options *ListCredentialsOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/credentials"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

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
func (discovery *DiscoveryV1) UpdateCredentials(options *UpdateCredentialsOptions) (*watson.WatsonResponse, []error) {
    path := "/v1/environments/{environment_id}/credentials/{credential_id}"
    creds := discovery.client.Creds
    useTM := discovery.client.UseTM
    tokenManager := discovery.client.TokenManager

    path = strings.Replace(path, "{environment_id}", options.EnvironmentID, 1)
    path = strings.Replace(path, "{credential_id}", options.CredentialID, 1)
    request := req.New().Put(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    body := map[string]interface{}{}
    if options.IsSourceTypeSet {
        body["source_type"] = options.SourceType
    }
    if options.IsCredentialDetailsSet {
        body["credential_details"] = options.CredentialDetails
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


// AddDocumentOptions : The addDocument options.
type AddDocumentOptions struct {

	// The ID of the environment.
	EnvironmentID string `json:"environment_id"`

	// The ID of the collection.
	CollectionID string `json:"collection_id"`

	// The content of the document to ingest. The maximum supported file size is 50 megabytes. Files larger than 50 megabytes is rejected.
	File os.File `json:"file,omitempty"`

    // Indicates whether user set optional parameter File
    IsFileSet bool

	// If you're using the Data Crawler to upload your documents, you can test a document against the type of metadata that the Data Crawler might send. The maximum supported metadata file size is 1 MB. Metadata parts larger than 1 MB are rejected. Example:  ``` { "Creator": "Johnny Appleseed", "Subject": "Apples" } ```.
	Metadata string `json:"metadata,omitempty"`

    // Indicates whether user set optional parameter Metadata
    IsMetadataSet bool

	// The content type of File.
	FileContentType string `json:"file_content_type,omitempty"`

    // Indicates whether user set optional parameter FileContentType
    IsFileContentTypeSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewAddDocumentOptions : Instantiate AddDocumentOptions
func NewAddDocumentOptions(environmentID string, collectionID string) *AddDocumentOptions {
    return &AddDocumentOptions{
        EnvironmentID: environmentID,
        CollectionID: collectionID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *AddDocumentOptions) SetEnvironmentID(param string) *AddDocumentOptions {
    options.EnvironmentID = param
    return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *AddDocumentOptions) SetCollectionID(param string) *AddDocumentOptions {
    options.CollectionID = param
    return options
}

// SetFile : Allow user to set File
func (options *AddDocumentOptions) SetFile(param os.File) *AddDocumentOptions {
    options.File = param
    options.IsFileSet = true
    return options
}

// SetMetadata : Allow user to set Metadata
func (options *AddDocumentOptions) SetMetadata(param string) *AddDocumentOptions {
    options.Metadata = param
    options.IsMetadataSet = true
    return options
}

// SetFileContentType : Allow user to set FileContentType
func (options *AddDocumentOptions) SetFileContentType(param string) *AddDocumentOptions {
    options.FileContentType = param
    options.IsFileContentTypeSet = true
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
	EnvironmentID string `json:"environment_id"`

	// The ID of the collection.
	CollectionID string `json:"collection_id"`

	NaturalLanguageQuery string `json:"natural_language_query,omitempty"`

    // Indicates whether user set optional parameter NaturalLanguageQuery
    IsNaturalLanguageQuerySet bool

	Filter string `json:"filter,omitempty"`

    // Indicates whether user set optional parameter Filter
    IsFilterSet bool

	Examples []TrainingExample `json:"examples,omitempty"`

    // Indicates whether user set optional parameter Examples
    IsExamplesSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewAddTrainingDataOptions : Instantiate AddTrainingDataOptions
func NewAddTrainingDataOptions(environmentID string, collectionID string) *AddTrainingDataOptions {
    return &AddTrainingDataOptions{
        EnvironmentID: environmentID,
        CollectionID: collectionID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *AddTrainingDataOptions) SetEnvironmentID(param string) *AddTrainingDataOptions {
    options.EnvironmentID = param
    return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *AddTrainingDataOptions) SetCollectionID(param string) *AddTrainingDataOptions {
    options.CollectionID = param
    return options
}

// SetNaturalLanguageQuery : Allow user to set NaturalLanguageQuery
func (options *AddTrainingDataOptions) SetNaturalLanguageQuery(param string) *AddTrainingDataOptions {
    options.NaturalLanguageQuery = param
    options.IsNaturalLanguageQuerySet = true
    return options
}

// SetFilter : Allow user to set Filter
func (options *AddTrainingDataOptions) SetFilter(param string) *AddTrainingDataOptions {
    options.Filter = param
    options.IsFilterSet = true
    return options
}

// SetExamples : Allow user to set Examples
func (options *AddTrainingDataOptions) SetExamples(param []TrainingExample) *AddTrainingDataOptions {
    options.Examples = param
    options.IsExamplesSet = true
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
	Name string `json:"name,omitempty"`

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

// CreateCollectionOptions : The createCollection options.
type CreateCollectionOptions struct {

	// The ID of the environment.
	EnvironmentID string `json:"environment_id"`

	// The name of the collection to be created.
	Name string `json:"name"`

	// A description of the collection.
	Description string `json:"description,omitempty"`

    // Indicates whether user set optional parameter Description
    IsDescriptionSet bool

	// The ID of the configuration in which the collection is to be created.
	ConfigurationID string `json:"configuration_id,omitempty"`

    // Indicates whether user set optional parameter ConfigurationID
    IsConfigurationIDSet bool

	// The language of the documents stored in the collection, in the form of an ISO 639-1 language code.
	Language string `json:"language,omitempty"`

    // Indicates whether user set optional parameter Language
    IsLanguageSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewCreateCollectionOptions : Instantiate CreateCollectionOptions
func NewCreateCollectionOptions(environmentID string, name string) *CreateCollectionOptions {
    return &CreateCollectionOptions{
        EnvironmentID: environmentID,
        Name: name,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *CreateCollectionOptions) SetEnvironmentID(param string) *CreateCollectionOptions {
    options.EnvironmentID = param
    return options
}

// SetName : Allow user to set Name
func (options *CreateCollectionOptions) SetName(param string) *CreateCollectionOptions {
    options.Name = param
    return options
}

// SetDescription : Allow user to set Description
func (options *CreateCollectionOptions) SetDescription(param string) *CreateCollectionOptions {
    options.Description = param
    options.IsDescriptionSet = true
    return options
}

// SetConfigurationID : Allow user to set ConfigurationID
func (options *CreateCollectionOptions) SetConfigurationID(param string) *CreateCollectionOptions {
    options.ConfigurationID = param
    options.IsConfigurationIDSet = true
    return options
}

// SetLanguage : Allow user to set Language
func (options *CreateCollectionOptions) SetLanguage(param string) *CreateCollectionOptions {
    options.Language = param
    options.IsLanguageSet = true
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
	EnvironmentID string `json:"environment_id"`

	// The name of the configuration.
	Name string `json:"name,omitempty"`

    // Indicates whether user set optional parameter Name
    IsNameSet bool

	// The description of the configuration, if available.
	Description string `json:"description,omitempty"`

    // Indicates whether user set optional parameter Description
    IsDescriptionSet bool

	// The document conversion settings for the configuration.
	Conversions Conversions `json:"conversions,omitempty"`

    // Indicates whether user set optional parameter Conversions
    IsConversionsSet bool

	// An array of document enrichment settings for the configuration.
	Enrichments []Enrichment `json:"enrichments,omitempty"`

    // Indicates whether user set optional parameter Enrichments
    IsEnrichmentsSet bool

	// Defines operations that can be used to transform the final output JSON into a normalized form. Operations are executed in the order that they appear in the array.
	Normalizations []NormalizationOperation `json:"normalizations,omitempty"`

    // Indicates whether user set optional parameter Normalizations
    IsNormalizationsSet bool

	// Object containing source parameters for the configuration.
	Source Source `json:"source,omitempty"`

    // Indicates whether user set optional parameter Source
    IsSourceSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewCreateConfigurationOptions : Instantiate CreateConfigurationOptions
func NewCreateConfigurationOptions(environmentID string) *CreateConfigurationOptions {
    return &CreateConfigurationOptions{
        EnvironmentID: environmentID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *CreateConfigurationOptions) SetEnvironmentID(param string) *CreateConfigurationOptions {
    options.EnvironmentID = param
    return options
}

// SetName : Allow user to set Name
func (options *CreateConfigurationOptions) SetName(param string) *CreateConfigurationOptions {
    options.Name = param
    options.IsNameSet = true
    return options
}

// SetDescription : Allow user to set Description
func (options *CreateConfigurationOptions) SetDescription(param string) *CreateConfigurationOptions {
    options.Description = param
    options.IsDescriptionSet = true
    return options
}

// SetConversions : Allow user to set Conversions
func (options *CreateConfigurationOptions) SetConversions(param Conversions) *CreateConfigurationOptions {
    options.Conversions = param
    options.IsConversionsSet = true
    return options
}

// SetEnrichments : Allow user to set Enrichments
func (options *CreateConfigurationOptions) SetEnrichments(param []Enrichment) *CreateConfigurationOptions {
    options.Enrichments = param
    options.IsEnrichmentsSet = true
    return options
}

// SetNormalizations : Allow user to set Normalizations
func (options *CreateConfigurationOptions) SetNormalizations(param []NormalizationOperation) *CreateConfigurationOptions {
    options.Normalizations = param
    options.IsNormalizationsSet = true
    return options
}

// SetSource : Allow user to set Source
func (options *CreateConfigurationOptions) SetSource(param Source) *CreateConfigurationOptions {
    options.Source = param
    options.IsSourceSet = true
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
	EnvironmentID string `json:"environment_id"`

	// The source that this credentials object connects to. -  `box` indicates the credentials are used to connect an instance of Enterprise Box. -  `salesforce` indicates the credentials are used to connect to Salesforce. -  `sharepoint` indicates the credentials are used to connect to Microsoft SharePoint Online.
	SourceType string `json:"source_type,omitempty"`

    // Indicates whether user set optional parameter SourceType
    IsSourceTypeSet bool

	// Object containing details of the stored credentials. Obtain credentials for your source from the administrator of the source.
	CredentialDetails CredentialDetails `json:"credential_details,omitempty"`

    // Indicates whether user set optional parameter CredentialDetails
    IsCredentialDetailsSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewCreateCredentialsOptions : Instantiate CreateCredentialsOptions
func NewCreateCredentialsOptions(environmentID string) *CreateCredentialsOptions {
    return &CreateCredentialsOptions{
        EnvironmentID: environmentID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *CreateCredentialsOptions) SetEnvironmentID(param string) *CreateCredentialsOptions {
    options.EnvironmentID = param
    return options
}

// SetSourceType : Allow user to set SourceType
func (options *CreateCredentialsOptions) SetSourceType(param string) *CreateCredentialsOptions {
    options.SourceType = param
    options.IsSourceTypeSet = true
    return options
}

// SetCredentialDetails : Allow user to set CredentialDetails
func (options *CreateCredentialsOptions) SetCredentialDetails(param CredentialDetails) *CreateCredentialsOptions {
    options.CredentialDetails = param
    options.IsCredentialDetailsSet = true
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
	Name string `json:"name"`

	// Description of the environment.
	Description string `json:"description,omitempty"`

    // Indicates whether user set optional parameter Description
    IsDescriptionSet bool

	// **Deprecated**: Size of the environment.
	Size int64 `json:"size,omitempty"`

    // Indicates whether user set optional parameter Size
    IsSizeSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewCreateEnvironmentOptions : Instantiate CreateEnvironmentOptions
func NewCreateEnvironmentOptions(name string) *CreateEnvironmentOptions {
    return &CreateEnvironmentOptions{
        Name: name,
    }
}

// SetName : Allow user to set Name
func (options *CreateEnvironmentOptions) SetName(param string) *CreateEnvironmentOptions {
    options.Name = param
    return options
}

// SetDescription : Allow user to set Description
func (options *CreateEnvironmentOptions) SetDescription(param string) *CreateEnvironmentOptions {
    options.Description = param
    options.IsDescriptionSet = true
    return options
}

// SetSize : Allow user to set Size
func (options *CreateEnvironmentOptions) SetSize(param int64) *CreateEnvironmentOptions {
    options.Size = param
    options.IsSizeSet = true
    return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateEnvironmentOptions) SetHeaders(param map[string]string) *CreateEnvironmentOptions {
    options.Headers = param
    return options
}

// CreateExpansionsOptions : The createExpansions options.
type CreateExpansionsOptions struct {

	// The ID of the environment.
	EnvironmentID string `json:"environment_id"`

	// The ID of the collection.
	CollectionID string `json:"collection_id"`

	// An array of query expansion definitions. Each object in the **expansions** array represents a term or set of terms that will be expanded into other terms. Each expansion object can be configured as bidirectional or unidirectional. Bidirectional means that all terms are expanded to all other terms in the object. Unidirectional means that a set list of terms can be expanded into a second list of terms. To create a bi-directional expansion specify an **expanded_terms** array. When found in a query, all items in the **expanded_terms** array are then expanded to the other items in the same array. To create a uni-directional expansion, specify both an array of **input_terms** and an array of **expanded_terms**. When items in the **input_terms** array are present in a query, they are expanded using the items listed in the **expanded_terms** array.
	Expansions []Expansion `json:"expansions,omitempty"`

    // Indicates whether user set optional parameter Expansions
    IsExpansionsSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewCreateExpansionsOptions : Instantiate CreateExpansionsOptions
func NewCreateExpansionsOptions(environmentID string, collectionID string) *CreateExpansionsOptions {
    return &CreateExpansionsOptions{
        EnvironmentID: environmentID,
        CollectionID: collectionID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *CreateExpansionsOptions) SetEnvironmentID(param string) *CreateExpansionsOptions {
    options.EnvironmentID = param
    return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *CreateExpansionsOptions) SetCollectionID(param string) *CreateExpansionsOptions {
    options.CollectionID = param
    return options
}

// SetExpansions : Allow user to set Expansions
func (options *CreateExpansionsOptions) SetExpansions(param []Expansion) *CreateExpansionsOptions {
    options.Expansions = param
    options.IsExpansionsSet = true
    return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateExpansionsOptions) SetHeaders(param map[string]string) *CreateExpansionsOptions {
    options.Headers = param
    return options
}

// CreateTrainingExampleOptions : The createTrainingExample options.
type CreateTrainingExampleOptions struct {

	// The ID of the environment.
	EnvironmentID string `json:"environment_id"`

	// The ID of the collection.
	CollectionID string `json:"collection_id"`

	// The ID of the query used for training.
	QueryID string `json:"query_id"`

	DocumentID string `json:"document_id,omitempty"`

    // Indicates whether user set optional parameter DocumentID
    IsDocumentIDSet bool

	CrossReference string `json:"cross_reference,omitempty"`

    // Indicates whether user set optional parameter CrossReference
    IsCrossReferenceSet bool

	Relevance int64 `json:"relevance,omitempty"`

    // Indicates whether user set optional parameter Relevance
    IsRelevanceSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewCreateTrainingExampleOptions : Instantiate CreateTrainingExampleOptions
func NewCreateTrainingExampleOptions(environmentID string, collectionID string, queryID string) *CreateTrainingExampleOptions {
    return &CreateTrainingExampleOptions{
        EnvironmentID: environmentID,
        CollectionID: collectionID,
        QueryID: queryID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *CreateTrainingExampleOptions) SetEnvironmentID(param string) *CreateTrainingExampleOptions {
    options.EnvironmentID = param
    return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *CreateTrainingExampleOptions) SetCollectionID(param string) *CreateTrainingExampleOptions {
    options.CollectionID = param
    return options
}

// SetQueryID : Allow user to set QueryID
func (options *CreateTrainingExampleOptions) SetQueryID(param string) *CreateTrainingExampleOptions {
    options.QueryID = param
    return options
}

// SetDocumentID : Allow user to set DocumentID
func (options *CreateTrainingExampleOptions) SetDocumentID(param string) *CreateTrainingExampleOptions {
    options.DocumentID = param
    options.IsDocumentIDSet = true
    return options
}

// SetCrossReference : Allow user to set CrossReference
func (options *CreateTrainingExampleOptions) SetCrossReference(param string) *CreateTrainingExampleOptions {
    options.CrossReference = param
    options.IsCrossReferenceSet = true
    return options
}

// SetRelevance : Allow user to set Relevance
func (options *CreateTrainingExampleOptions) SetRelevance(param int64) *CreateTrainingExampleOptions {
    options.Relevance = param
    options.IsRelevanceSet = true
    return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateTrainingExampleOptions) SetHeaders(param map[string]string) *CreateTrainingExampleOptions {
    options.Headers = param
    return options
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
	SiteCollectionPath string `json:"site_collection.path,omitempty"`

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

// DeleteAllTrainingDataOptions : The deleteAllTrainingData options.
type DeleteAllTrainingDataOptions struct {

	// The ID of the environment.
	EnvironmentID string `json:"environment_id"`

	// The ID of the collection.
	CollectionID string `json:"collection_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewDeleteAllTrainingDataOptions : Instantiate DeleteAllTrainingDataOptions
func NewDeleteAllTrainingDataOptions(environmentID string, collectionID string) *DeleteAllTrainingDataOptions {
    return &DeleteAllTrainingDataOptions{
        EnvironmentID: environmentID,
        CollectionID: collectionID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *DeleteAllTrainingDataOptions) SetEnvironmentID(param string) *DeleteAllTrainingDataOptions {
    options.EnvironmentID = param
    return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *DeleteAllTrainingDataOptions) SetCollectionID(param string) *DeleteAllTrainingDataOptions {
    options.CollectionID = param
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
	EnvironmentID string `json:"environment_id"`

	// The ID of the collection.
	CollectionID string `json:"collection_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewDeleteCollectionOptions : Instantiate DeleteCollectionOptions
func NewDeleteCollectionOptions(environmentID string, collectionID string) *DeleteCollectionOptions {
    return &DeleteCollectionOptions{
        EnvironmentID: environmentID,
        CollectionID: collectionID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *DeleteCollectionOptions) SetEnvironmentID(param string) *DeleteCollectionOptions {
    options.EnvironmentID = param
    return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *DeleteCollectionOptions) SetCollectionID(param string) *DeleteCollectionOptions {
    options.CollectionID = param
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
	CollectionID string `json:"collection_id"`

	// The status of the collection. The status of a successful deletion operation is `deleted`.
	Status string `json:"status"`
}

// DeleteConfigurationOptions : The deleteConfiguration options.
type DeleteConfigurationOptions struct {

	// The ID of the environment.
	EnvironmentID string `json:"environment_id"`

	// The ID of the configuration.
	ConfigurationID string `json:"configuration_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewDeleteConfigurationOptions : Instantiate DeleteConfigurationOptions
func NewDeleteConfigurationOptions(environmentID string, configurationID string) *DeleteConfigurationOptions {
    return &DeleteConfigurationOptions{
        EnvironmentID: environmentID,
        ConfigurationID: configurationID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *DeleteConfigurationOptions) SetEnvironmentID(param string) *DeleteConfigurationOptions {
    options.EnvironmentID = param
    return options
}

// SetConfigurationID : Allow user to set ConfigurationID
func (options *DeleteConfigurationOptions) SetConfigurationID(param string) *DeleteConfigurationOptions {
    options.ConfigurationID = param
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

// DeleteCredentialsOptions : The deleteCredentials options.
type DeleteCredentialsOptions struct {

	// The ID of the environment.
	EnvironmentID string `json:"environment_id"`

	// The unique identifier for a set of source credentials.
	CredentialID string `json:"credential_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewDeleteCredentialsOptions : Instantiate DeleteCredentialsOptions
func NewDeleteCredentialsOptions(environmentID string, credentialID string) *DeleteCredentialsOptions {
    return &DeleteCredentialsOptions{
        EnvironmentID: environmentID,
        CredentialID: credentialID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *DeleteCredentialsOptions) SetEnvironmentID(param string) *DeleteCredentialsOptions {
    options.EnvironmentID = param
    return options
}

// SetCredentialID : Allow user to set CredentialID
func (options *DeleteCredentialsOptions) SetCredentialID(param string) *DeleteCredentialsOptions {
    options.CredentialID = param
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
	EnvironmentID string `json:"environment_id"`

	// The ID of the collection.
	CollectionID string `json:"collection_id"`

	// The ID of the document.
	DocumentID string `json:"document_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewDeleteDocumentOptions : Instantiate DeleteDocumentOptions
func NewDeleteDocumentOptions(environmentID string, collectionID string, documentID string) *DeleteDocumentOptions {
    return &DeleteDocumentOptions{
        EnvironmentID: environmentID,
        CollectionID: collectionID,
        DocumentID: documentID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *DeleteDocumentOptions) SetEnvironmentID(param string) *DeleteDocumentOptions {
    options.EnvironmentID = param
    return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *DeleteDocumentOptions) SetCollectionID(param string) *DeleteDocumentOptions {
    options.CollectionID = param
    return options
}

// SetDocumentID : Allow user to set DocumentID
func (options *DeleteDocumentOptions) SetDocumentID(param string) *DeleteDocumentOptions {
    options.DocumentID = param
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
	DocumentID string `json:"document_id,omitempty"`

	// Status of the document. A deleted document has the status deleted.
	Status string `json:"status,omitempty"`
}

// DeleteEnvironmentOptions : The deleteEnvironment options.
type DeleteEnvironmentOptions struct {

	// The ID of the environment.
	EnvironmentID string `json:"environment_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewDeleteEnvironmentOptions : Instantiate DeleteEnvironmentOptions
func NewDeleteEnvironmentOptions(environmentID string) *DeleteEnvironmentOptions {
    return &DeleteEnvironmentOptions{
        EnvironmentID: environmentID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *DeleteEnvironmentOptions) SetEnvironmentID(param string) *DeleteEnvironmentOptions {
    options.EnvironmentID = param
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
	EnvironmentID string `json:"environment_id"`

	// Status of the environment.
	Status string `json:"status"`
}

// DeleteExpansionsOptions : The deleteExpansions options.
type DeleteExpansionsOptions struct {

	// The ID of the environment.
	EnvironmentID string `json:"environment_id"`

	// The ID of the collection.
	CollectionID string `json:"collection_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewDeleteExpansionsOptions : Instantiate DeleteExpansionsOptions
func NewDeleteExpansionsOptions(environmentID string, collectionID string) *DeleteExpansionsOptions {
    return &DeleteExpansionsOptions{
        EnvironmentID: environmentID,
        CollectionID: collectionID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *DeleteExpansionsOptions) SetEnvironmentID(param string) *DeleteExpansionsOptions {
    options.EnvironmentID = param
    return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *DeleteExpansionsOptions) SetCollectionID(param string) *DeleteExpansionsOptions {
    options.CollectionID = param
    return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteExpansionsOptions) SetHeaders(param map[string]string) *DeleteExpansionsOptions {
    options.Headers = param
    return options
}

// DeleteTrainingDataOptions : The deleteTrainingData options.
type DeleteTrainingDataOptions struct {

	// The ID of the environment.
	EnvironmentID string `json:"environment_id"`

	// The ID of the collection.
	CollectionID string `json:"collection_id"`

	// The ID of the query used for training.
	QueryID string `json:"query_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewDeleteTrainingDataOptions : Instantiate DeleteTrainingDataOptions
func NewDeleteTrainingDataOptions(environmentID string, collectionID string, queryID string) *DeleteTrainingDataOptions {
    return &DeleteTrainingDataOptions{
        EnvironmentID: environmentID,
        CollectionID: collectionID,
        QueryID: queryID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *DeleteTrainingDataOptions) SetEnvironmentID(param string) *DeleteTrainingDataOptions {
    options.EnvironmentID = param
    return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *DeleteTrainingDataOptions) SetCollectionID(param string) *DeleteTrainingDataOptions {
    options.CollectionID = param
    return options
}

// SetQueryID : Allow user to set QueryID
func (options *DeleteTrainingDataOptions) SetQueryID(param string) *DeleteTrainingDataOptions {
    options.QueryID = param
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
	EnvironmentID string `json:"environment_id"`

	// The ID of the collection.
	CollectionID string `json:"collection_id"`

	// The ID of the query used for training.
	QueryID string `json:"query_id"`

	// The ID of the document as it is indexed.
	ExampleID string `json:"example_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewDeleteTrainingExampleOptions : Instantiate DeleteTrainingExampleOptions
func NewDeleteTrainingExampleOptions(environmentID string, collectionID string, queryID string, exampleID string) *DeleteTrainingExampleOptions {
    return &DeleteTrainingExampleOptions{
        EnvironmentID: environmentID,
        CollectionID: collectionID,
        QueryID: queryID,
        ExampleID: exampleID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *DeleteTrainingExampleOptions) SetEnvironmentID(param string) *DeleteTrainingExampleOptions {
    options.EnvironmentID = param
    return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *DeleteTrainingExampleOptions) SetCollectionID(param string) *DeleteTrainingExampleOptions {
    options.CollectionID = param
    return options
}

// SetQueryID : Allow user to set QueryID
func (options *DeleteTrainingExampleOptions) SetQueryID(param string) *DeleteTrainingExampleOptions {
    options.QueryID = param
    return options
}

// SetExampleID : Allow user to set ExampleID
func (options *DeleteTrainingExampleOptions) SetExampleID(param string) *DeleteTrainingExampleOptions {
    options.ExampleID = param
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
	EnrichmentName string `json:"enrichment"`

	// If true, then most errors generated during the enrichment process will be treated as warnings and will not cause the document to fail processing.
	IgnoreDownstreamErrors bool `json:"ignore_downstream_errors,omitempty"`

	// A list of options specific to the enrichment.
	Options EnrichmentOptions `json:"options,omitempty"`
}

// EnrichmentOptions : Options which are specific to a particular enrichment.
type EnrichmentOptions struct {

	// An object representing the enrichment features that will be applied to the specified field.
	Features NluEnrichmentFeatures `json:"features,omitempty"`

	// ISO 639-1 code indicating the language to use for the analysis. This code overrides the automatic language detection performed by the service. Valid codes are `ar` (Arabic), `en` (English), `fr` (French), `de` (German), `it` (Italian), `pt` (Portuguese), `ru` (Russian), `es` (Spanish), and `sv` (Swedish). **Note:** Not all features support all languages, automatic detection is recommended.
	Language string `json:"language,omitempty"`

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
	Expansions []Expansion `json:"expansions,omitempty"`
}

// FederatedQueryNoticesOptions : The federatedQueryNotices options.
type FederatedQueryNoticesOptions struct {

	// The ID of the environment.
	EnvironmentID string `json:"environment_id"`

	// A comma-separated list of collection IDs to be queried against.
	CollectionIds []string `json:"collection_ids"`

	// A cacheable query that limits the documents returned to exclude any documents that don't mention the query content. Filter searches are better for metadata type searches and when you are trying to get a sense of concepts in the data set.
	Filter string `json:"filter,omitempty"`

    // Indicates whether user set optional parameter Filter
    IsFilterSet bool

	// A query search returns all documents in your data set with full enrichments and full text, but with the most relevant documents listed first. Use a query search when you want to find the most relevant search results. You cannot use **natural_language_query** and **query** at the same time.
	Query string `json:"query,omitempty"`

    // Indicates whether user set optional parameter Query
    IsQuerySet bool

	// A natural language query that returns relevant documents by utilizing training data and natural language understanding. You cannot use **natural_language_query** and **query** at the same time.
	NaturalLanguageQuery string `json:"natural_language_query,omitempty"`

    // Indicates whether user set optional parameter NaturalLanguageQuery
    IsNaturalLanguageQuerySet bool

	// An aggregation search uses combinations of filters and query search to return an exact answer. Aggregations are useful for building applications, because you can use them to build lists, tables, and time series. For a full list of possible aggregrations, see the Query reference.
	Aggregation string `json:"aggregation,omitempty"`

    // Indicates whether user set optional parameter Aggregation
    IsAggregationSet bool

	// Number of documents to return.
	Count int64 `json:"count,omitempty"`

    // Indicates whether user set optional parameter Count
    IsCountSet bool

	// A comma separated list of the portion of the document hierarchy to return.
	ReturnFields []string `json:"return,omitempty"`

    // Indicates whether user set optional parameter ReturnFields
    IsReturnFieldsSet bool

	// The number of query results to skip at the beginning. For example, if the total number of results that are returned is 10, and the offset is 8, it returns the last two results.
	Offset int64 `json:"offset,omitempty"`

    // Indicates whether user set optional parameter Offset
    IsOffsetSet bool

	// A comma separated list of fields in the document to sort on. You can optionally specify a sort direction by prefixing the field with `-` for descending or `+` for ascending. Ascending is the default sort direction if no prefix is specified.
	Sort []string `json:"sort,omitempty"`

    // Indicates whether user set optional parameter Sort
    IsSortSet bool

	// When true a highlight field is returned for each result which contains the fields that match the query with `<em></em>` tags around the matching query terms. Defaults to false.
	Highlight bool `json:"highlight,omitempty"`

    // Indicates whether user set optional parameter Highlight
    IsHighlightSet bool

	// When specified, duplicate results based on the field specified are removed from the returned results. Duplicate comparison is limited to the current query only, **offset** is not considered. This parameter is currently Beta functionality.
	DeduplicateField string `json:"deduplicate.field,omitempty"`

    // Indicates whether user set optional parameter DeduplicateField
    IsDeduplicateFieldSet bool

	// When `true`, results are returned based on their similarity to the document IDs specified in the **similar.document_ids** parameter.
	Similar bool `json:"similar,omitempty"`

    // Indicates whether user set optional parameter Similar
    IsSimilarSet bool

	// A comma-separated list of document IDs that will be used to find similar documents. **Note:** If the **natural_language_query** parameter is also specified, it will be used to expand the scope of the document similarity search to include the natural language query. Other query parameters, such as **filter** and **query** are subsequently applied and reduce the query scope.
	SimilarDocumentIds []string `json:"similar.document_ids,omitempty"`

    // Indicates whether user set optional parameter SimilarDocumentIds
    IsSimilarDocumentIdsSet bool

	// A comma-separated list of field names that will be used as a basis for comparison to identify similar documents. If not specified, the entire document is used for comparison.
	SimilarFields []string `json:"similar.fields,omitempty"`

    // Indicates whether user set optional parameter SimilarFields
    IsSimilarFieldsSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewFederatedQueryNoticesOptions : Instantiate FederatedQueryNoticesOptions
func NewFederatedQueryNoticesOptions(environmentID string, collectionIds []string) *FederatedQueryNoticesOptions {
    return &FederatedQueryNoticesOptions{
        EnvironmentID: environmentID,
        CollectionIds: collectionIds,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *FederatedQueryNoticesOptions) SetEnvironmentID(param string) *FederatedQueryNoticesOptions {
    options.EnvironmentID = param
    return options
}

// SetCollectionIds : Allow user to set CollectionIds
func (options *FederatedQueryNoticesOptions) SetCollectionIds(param []string) *FederatedQueryNoticesOptions {
    options.CollectionIds = param
    return options
}

// SetFilter : Allow user to set Filter
func (options *FederatedQueryNoticesOptions) SetFilter(param string) *FederatedQueryNoticesOptions {
    options.Filter = param
    options.IsFilterSet = true
    return options
}

// SetQuery : Allow user to set Query
func (options *FederatedQueryNoticesOptions) SetQuery(param string) *FederatedQueryNoticesOptions {
    options.Query = param
    options.IsQuerySet = true
    return options
}

// SetNaturalLanguageQuery : Allow user to set NaturalLanguageQuery
func (options *FederatedQueryNoticesOptions) SetNaturalLanguageQuery(param string) *FederatedQueryNoticesOptions {
    options.NaturalLanguageQuery = param
    options.IsNaturalLanguageQuerySet = true
    return options
}

// SetAggregation : Allow user to set Aggregation
func (options *FederatedQueryNoticesOptions) SetAggregation(param string) *FederatedQueryNoticesOptions {
    options.Aggregation = param
    options.IsAggregationSet = true
    return options
}

// SetCount : Allow user to set Count
func (options *FederatedQueryNoticesOptions) SetCount(param int64) *FederatedQueryNoticesOptions {
    options.Count = param
    options.IsCountSet = true
    return options
}

// SetReturnFields : Allow user to set ReturnFields
func (options *FederatedQueryNoticesOptions) SetReturnFields(param []string) *FederatedQueryNoticesOptions {
    options.ReturnFields = param
    options.IsReturnFieldsSet = true
    return options
}

// SetOffset : Allow user to set Offset
func (options *FederatedQueryNoticesOptions) SetOffset(param int64) *FederatedQueryNoticesOptions {
    options.Offset = param
    options.IsOffsetSet = true
    return options
}

// SetSort : Allow user to set Sort
func (options *FederatedQueryNoticesOptions) SetSort(param []string) *FederatedQueryNoticesOptions {
    options.Sort = param
    options.IsSortSet = true
    return options
}

// SetHighlight : Allow user to set Highlight
func (options *FederatedQueryNoticesOptions) SetHighlight(param bool) *FederatedQueryNoticesOptions {
    options.Highlight = param
    options.IsHighlightSet = true
    return options
}

// SetDeduplicateField : Allow user to set DeduplicateField
func (options *FederatedQueryNoticesOptions) SetDeduplicateField(param string) *FederatedQueryNoticesOptions {
    options.DeduplicateField = param
    options.IsDeduplicateFieldSet = true
    return options
}

// SetSimilar : Allow user to set Similar
func (options *FederatedQueryNoticesOptions) SetSimilar(param bool) *FederatedQueryNoticesOptions {
    options.Similar = param
    options.IsSimilarSet = true
    return options
}

// SetSimilarDocumentIds : Allow user to set SimilarDocumentIds
func (options *FederatedQueryNoticesOptions) SetSimilarDocumentIds(param []string) *FederatedQueryNoticesOptions {
    options.SimilarDocumentIds = param
    options.IsSimilarDocumentIdsSet = true
    return options
}

// SetSimilarFields : Allow user to set SimilarFields
func (options *FederatedQueryNoticesOptions) SetSimilarFields(param []string) *FederatedQueryNoticesOptions {
    options.SimilarFields = param
    options.IsSimilarFieldsSet = true
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
	EnvironmentID string `json:"environment_id"`

	// A comma-separated list of collection IDs to be queried against.
	CollectionIds []string `json:"collection_ids"`

	// A cacheable query that limits the documents returned to exclude any documents that don't mention the query content. Filter searches are better for metadata type searches and when you are trying to get a sense of concepts in the data set.
	Filter string `json:"filter,omitempty"`

    // Indicates whether user set optional parameter Filter
    IsFilterSet bool

	// A query search returns all documents in your data set with full enrichments and full text, but with the most relevant documents listed first. Use a query search when you want to find the most relevant search results. You cannot use **natural_language_query** and **query** at the same time.
	Query string `json:"query,omitempty"`

    // Indicates whether user set optional parameter Query
    IsQuerySet bool

	// A natural language query that returns relevant documents by utilizing training data and natural language understanding. You cannot use **natural_language_query** and **query** at the same time.
	NaturalLanguageQuery string `json:"natural_language_query,omitempty"`

    // Indicates whether user set optional parameter NaturalLanguageQuery
    IsNaturalLanguageQuerySet bool

	// An aggregation search uses combinations of filters and query search to return an exact answer. Aggregations are useful for building applications, because you can use them to build lists, tables, and time series. For a full list of possible aggregrations, see the Query reference.
	Aggregation string `json:"aggregation,omitempty"`

    // Indicates whether user set optional parameter Aggregation
    IsAggregationSet bool

	// Number of documents to return.
	Count int64 `json:"count,omitempty"`

    // Indicates whether user set optional parameter Count
    IsCountSet bool

	// A comma separated list of the portion of the document hierarchy to return.
	ReturnFields []string `json:"return,omitempty"`

    // Indicates whether user set optional parameter ReturnFields
    IsReturnFieldsSet bool

	// The number of query results to skip at the beginning. For example, if the total number of results that are returned is 10, and the offset is 8, it returns the last two results.
	Offset int64 `json:"offset,omitempty"`

    // Indicates whether user set optional parameter Offset
    IsOffsetSet bool

	// A comma separated list of fields in the document to sort on. You can optionally specify a sort direction by prefixing the field with `-` for descending or `+` for ascending. Ascending is the default sort direction if no prefix is specified.
	Sort []string `json:"sort,omitempty"`

    // Indicates whether user set optional parameter Sort
    IsSortSet bool

	// When true a highlight field is returned for each result which contains the fields that match the query with `<em></em>` tags around the matching query terms. Defaults to false.
	Highlight bool `json:"highlight,omitempty"`

    // Indicates whether user set optional parameter Highlight
    IsHighlightSet bool

	// When `true` and used with a Watson Discovery News collection, duplicate results (based on the contents of the **title** field) are removed. Duplicate comparison is limited to the current query only; **offset** is not considered. This parameter is currently Beta functionality.
	Deduplicate bool `json:"deduplicate,omitempty"`

    // Indicates whether user set optional parameter Deduplicate
    IsDeduplicateSet bool

	// When specified, duplicate results based on the field specified are removed from the returned results. Duplicate comparison is limited to the current query only, **offset** is not considered. This parameter is currently Beta functionality.
	DeduplicateField string `json:"deduplicate.field,omitempty"`

    // Indicates whether user set optional parameter DeduplicateField
    IsDeduplicateFieldSet bool

	// When `true`, results are returned based on their similarity to the document IDs specified in the **similar.document_ids** parameter.
	Similar bool `json:"similar,omitempty"`

    // Indicates whether user set optional parameter Similar
    IsSimilarSet bool

	// A comma-separated list of document IDs that will be used to find similar documents. **Note:** If the **natural_language_query** parameter is also specified, it will be used to expand the scope of the document similarity search to include the natural language query. Other query parameters, such as **filter** and **query** are subsequently applied and reduce the query scope.
	SimilarDocumentIds []string `json:"similar.document_ids,omitempty"`

    // Indicates whether user set optional parameter SimilarDocumentIds
    IsSimilarDocumentIdsSet bool

	// A comma-separated list of field names that will be used as a basis for comparison to identify similar documents. If not specified, the entire document is used for comparison.
	SimilarFields []string `json:"similar.fields,omitempty"`

    // Indicates whether user set optional parameter SimilarFields
    IsSimilarFieldsSet bool

	// A passages query that returns the most relevant passages from the results.
	Passages bool `json:"passages,omitempty"`

    // Indicates whether user set optional parameter Passages
    IsPassagesSet bool

	// A comma-separated list of fields that passages are drawn from. If this parameter not specified, then all top-level fields are included.
	PassagesFields []string `json:"passages.fields,omitempty"`

    // Indicates whether user set optional parameter PassagesFields
    IsPassagesFieldsSet bool

	// The maximum number of passages to return. The search returns fewer passages if the requested total is not found. The default is `10`. The maximum is `100`.
	PassagesCount int64 `json:"passages.count,omitempty"`

    // Indicates whether user set optional parameter PassagesCount
    IsPassagesCountSet bool

	// The approximate number of characters that any one passage will have. The default is `400`. The minimum is `50`. The maximum is `2000`.
	PassagesCharacters int64 `json:"passages.characters,omitempty"`

    // Indicates whether user set optional parameter PassagesCharacters
    IsPassagesCharactersSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewFederatedQueryOptions : Instantiate FederatedQueryOptions
func NewFederatedQueryOptions(environmentID string, collectionIds []string) *FederatedQueryOptions {
    return &FederatedQueryOptions{
        EnvironmentID: environmentID,
        CollectionIds: collectionIds,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *FederatedQueryOptions) SetEnvironmentID(param string) *FederatedQueryOptions {
    options.EnvironmentID = param
    return options
}

// SetCollectionIds : Allow user to set CollectionIds
func (options *FederatedQueryOptions) SetCollectionIds(param []string) *FederatedQueryOptions {
    options.CollectionIds = param
    return options
}

// SetFilter : Allow user to set Filter
func (options *FederatedQueryOptions) SetFilter(param string) *FederatedQueryOptions {
    options.Filter = param
    options.IsFilterSet = true
    return options
}

// SetQuery : Allow user to set Query
func (options *FederatedQueryOptions) SetQuery(param string) *FederatedQueryOptions {
    options.Query = param
    options.IsQuerySet = true
    return options
}

// SetNaturalLanguageQuery : Allow user to set NaturalLanguageQuery
func (options *FederatedQueryOptions) SetNaturalLanguageQuery(param string) *FederatedQueryOptions {
    options.NaturalLanguageQuery = param
    options.IsNaturalLanguageQuerySet = true
    return options
}

// SetAggregation : Allow user to set Aggregation
func (options *FederatedQueryOptions) SetAggregation(param string) *FederatedQueryOptions {
    options.Aggregation = param
    options.IsAggregationSet = true
    return options
}

// SetCount : Allow user to set Count
func (options *FederatedQueryOptions) SetCount(param int64) *FederatedQueryOptions {
    options.Count = param
    options.IsCountSet = true
    return options
}

// SetReturnFields : Allow user to set ReturnFields
func (options *FederatedQueryOptions) SetReturnFields(param []string) *FederatedQueryOptions {
    options.ReturnFields = param
    options.IsReturnFieldsSet = true
    return options
}

// SetOffset : Allow user to set Offset
func (options *FederatedQueryOptions) SetOffset(param int64) *FederatedQueryOptions {
    options.Offset = param
    options.IsOffsetSet = true
    return options
}

// SetSort : Allow user to set Sort
func (options *FederatedQueryOptions) SetSort(param []string) *FederatedQueryOptions {
    options.Sort = param
    options.IsSortSet = true
    return options
}

// SetHighlight : Allow user to set Highlight
func (options *FederatedQueryOptions) SetHighlight(param bool) *FederatedQueryOptions {
    options.Highlight = param
    options.IsHighlightSet = true
    return options
}

// SetDeduplicate : Allow user to set Deduplicate
func (options *FederatedQueryOptions) SetDeduplicate(param bool) *FederatedQueryOptions {
    options.Deduplicate = param
    options.IsDeduplicateSet = true
    return options
}

// SetDeduplicateField : Allow user to set DeduplicateField
func (options *FederatedQueryOptions) SetDeduplicateField(param string) *FederatedQueryOptions {
    options.DeduplicateField = param
    options.IsDeduplicateFieldSet = true
    return options
}

// SetSimilar : Allow user to set Similar
func (options *FederatedQueryOptions) SetSimilar(param bool) *FederatedQueryOptions {
    options.Similar = param
    options.IsSimilarSet = true
    return options
}

// SetSimilarDocumentIds : Allow user to set SimilarDocumentIds
func (options *FederatedQueryOptions) SetSimilarDocumentIds(param []string) *FederatedQueryOptions {
    options.SimilarDocumentIds = param
    options.IsSimilarDocumentIdsSet = true
    return options
}

// SetSimilarFields : Allow user to set SimilarFields
func (options *FederatedQueryOptions) SetSimilarFields(param []string) *FederatedQueryOptions {
    options.SimilarFields = param
    options.IsSimilarFieldsSet = true
    return options
}

// SetPassages : Allow user to set Passages
func (options *FederatedQueryOptions) SetPassages(param bool) *FederatedQueryOptions {
    options.Passages = param
    options.IsPassagesSet = true
    return options
}

// SetPassagesFields : Allow user to set PassagesFields
func (options *FederatedQueryOptions) SetPassagesFields(param []string) *FederatedQueryOptions {
    options.PassagesFields = param
    options.IsPassagesFieldsSet = true
    return options
}

// SetPassagesCount : Allow user to set PassagesCount
func (options *FederatedQueryOptions) SetPassagesCount(param int64) *FederatedQueryOptions {
    options.PassagesCount = param
    options.IsPassagesCountSet = true
    return options
}

// SetPassagesCharacters : Allow user to set PassagesCharacters
func (options *FederatedQueryOptions) SetPassagesCharacters(param int64) *FederatedQueryOptions {
    options.PassagesCharacters = param
    options.IsPassagesCharactersSet = true
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
	FieldName string `json:"field,omitempty"`

	// The type of the field.
	FieldType string `json:"type,omitempty"`
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

// GetCollectionOptions : The getCollection options.
type GetCollectionOptions struct {

	// The ID of the environment.
	EnvironmentID string `json:"environment_id"`

	// The ID of the collection.
	CollectionID string `json:"collection_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewGetCollectionOptions : Instantiate GetCollectionOptions
func NewGetCollectionOptions(environmentID string, collectionID string) *GetCollectionOptions {
    return &GetCollectionOptions{
        EnvironmentID: environmentID,
        CollectionID: collectionID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *GetCollectionOptions) SetEnvironmentID(param string) *GetCollectionOptions {
    options.EnvironmentID = param
    return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *GetCollectionOptions) SetCollectionID(param string) *GetCollectionOptions {
    options.CollectionID = param
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
	EnvironmentID string `json:"environment_id"`

	// The ID of the configuration.
	ConfigurationID string `json:"configuration_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewGetConfigurationOptions : Instantiate GetConfigurationOptions
func NewGetConfigurationOptions(environmentID string, configurationID string) *GetConfigurationOptions {
    return &GetConfigurationOptions{
        EnvironmentID: environmentID,
        ConfigurationID: configurationID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *GetConfigurationOptions) SetEnvironmentID(param string) *GetConfigurationOptions {
    options.EnvironmentID = param
    return options
}

// SetConfigurationID : Allow user to set ConfigurationID
func (options *GetConfigurationOptions) SetConfigurationID(param string) *GetConfigurationOptions {
    options.ConfigurationID = param
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
	EnvironmentID string `json:"environment_id"`

	// The unique identifier for a set of source credentials.
	CredentialID string `json:"credential_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewGetCredentialsOptions : Instantiate GetCredentialsOptions
func NewGetCredentialsOptions(environmentID string, credentialID string) *GetCredentialsOptions {
    return &GetCredentialsOptions{
        EnvironmentID: environmentID,
        CredentialID: credentialID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *GetCredentialsOptions) SetEnvironmentID(param string) *GetCredentialsOptions {
    options.EnvironmentID = param
    return options
}

// SetCredentialID : Allow user to set CredentialID
func (options *GetCredentialsOptions) SetCredentialID(param string) *GetCredentialsOptions {
    options.CredentialID = param
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
	EnvironmentID string `json:"environment_id"`

	// The ID of the collection.
	CollectionID string `json:"collection_id"`

	// The ID of the document.
	DocumentID string `json:"document_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewGetDocumentStatusOptions : Instantiate GetDocumentStatusOptions
func NewGetDocumentStatusOptions(environmentID string, collectionID string, documentID string) *GetDocumentStatusOptions {
    return &GetDocumentStatusOptions{
        EnvironmentID: environmentID,
        CollectionID: collectionID,
        DocumentID: documentID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *GetDocumentStatusOptions) SetEnvironmentID(param string) *GetDocumentStatusOptions {
    options.EnvironmentID = param
    return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *GetDocumentStatusOptions) SetCollectionID(param string) *GetDocumentStatusOptions {
    options.CollectionID = param
    return options
}

// SetDocumentID : Allow user to set DocumentID
func (options *GetDocumentStatusOptions) SetDocumentID(param string) *GetDocumentStatusOptions {
    options.DocumentID = param
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
	EnvironmentID string `json:"environment_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewGetEnvironmentOptions : Instantiate GetEnvironmentOptions
func NewGetEnvironmentOptions(environmentID string) *GetEnvironmentOptions {
    return &GetEnvironmentOptions{
        EnvironmentID: environmentID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *GetEnvironmentOptions) SetEnvironmentID(param string) *GetEnvironmentOptions {
    options.EnvironmentID = param
    return options
}

// SetHeaders : Allow user to set Headers
func (options *GetEnvironmentOptions) SetHeaders(param map[string]string) *GetEnvironmentOptions {
    options.Headers = param
    return options
}

// GetTrainingDataOptions : The getTrainingData options.
type GetTrainingDataOptions struct {

	// The ID of the environment.
	EnvironmentID string `json:"environment_id"`

	// The ID of the collection.
	CollectionID string `json:"collection_id"`

	// The ID of the query used for training.
	QueryID string `json:"query_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewGetTrainingDataOptions : Instantiate GetTrainingDataOptions
func NewGetTrainingDataOptions(environmentID string, collectionID string, queryID string) *GetTrainingDataOptions {
    return &GetTrainingDataOptions{
        EnvironmentID: environmentID,
        CollectionID: collectionID,
        QueryID: queryID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *GetTrainingDataOptions) SetEnvironmentID(param string) *GetTrainingDataOptions {
    options.EnvironmentID = param
    return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *GetTrainingDataOptions) SetCollectionID(param string) *GetTrainingDataOptions {
    options.CollectionID = param
    return options
}

// SetQueryID : Allow user to set QueryID
func (options *GetTrainingDataOptions) SetQueryID(param string) *GetTrainingDataOptions {
    options.QueryID = param
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
	EnvironmentID string `json:"environment_id"`

	// The ID of the collection.
	CollectionID string `json:"collection_id"`

	// The ID of the query used for training.
	QueryID string `json:"query_id"`

	// The ID of the document as it is indexed.
	ExampleID string `json:"example_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewGetTrainingExampleOptions : Instantiate GetTrainingExampleOptions
func NewGetTrainingExampleOptions(environmentID string, collectionID string, queryID string, exampleID string) *GetTrainingExampleOptions {
    return &GetTrainingExampleOptions{
        EnvironmentID: environmentID,
        CollectionID: collectionID,
        QueryID: queryID,
        ExampleID: exampleID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *GetTrainingExampleOptions) SetEnvironmentID(param string) *GetTrainingExampleOptions {
    options.EnvironmentID = param
    return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *GetTrainingExampleOptions) SetCollectionID(param string) *GetTrainingExampleOptions {
    options.CollectionID = param
    return options
}

// SetQueryID : Allow user to set QueryID
func (options *GetTrainingExampleOptions) SetQueryID(param string) *GetTrainingExampleOptions {
    options.QueryID = param
    return options
}

// SetExampleID : Allow user to set ExampleID
func (options *GetTrainingExampleOptions) SetExampleID(param string) *GetTrainingExampleOptions {
    options.ExampleID = param
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

// ListCollectionFieldsOptions : The listCollectionFields options.
type ListCollectionFieldsOptions struct {

	// The ID of the environment.
	EnvironmentID string `json:"environment_id"`

	// The ID of the collection.
	CollectionID string `json:"collection_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewListCollectionFieldsOptions : Instantiate ListCollectionFieldsOptions
func NewListCollectionFieldsOptions(environmentID string, collectionID string) *ListCollectionFieldsOptions {
    return &ListCollectionFieldsOptions{
        EnvironmentID: environmentID,
        CollectionID: collectionID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *ListCollectionFieldsOptions) SetEnvironmentID(param string) *ListCollectionFieldsOptions {
    options.EnvironmentID = param
    return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *ListCollectionFieldsOptions) SetCollectionID(param string) *ListCollectionFieldsOptions {
    options.CollectionID = param
    return options
}

// SetHeaders : Allow user to set Headers
func (options *ListCollectionFieldsOptions) SetHeaders(param map[string]string) *ListCollectionFieldsOptions {
    options.Headers = param
    return options
}

// ListCollectionFieldsResponse : The list of fetched fields. The fields are returned using a fully qualified name format, however, the format differs slightly from that used by the query operations. * Fields which contain nested JSON objects are assigned a type of "nested". * Fields which belong to a nested object are prefixed with `.properties` (for example, `warnings.properties.severity` means that the `warnings` object has a property called `severity`). * Fields returned from the News collection are prefixed with `v{N}-fullnews-t3-{YEAR}.mappings` (for example, `v5-fullnews-t3-2016.mappings.text.properties.author`).
type ListCollectionFieldsResponse struct {

	// An array containing information about each field in the collections.
	Fields []Field `json:"fields,omitempty"`
}

// ListCollectionsOptions : The listCollections options.
type ListCollectionsOptions struct {

	// The ID of the environment.
	EnvironmentID string `json:"environment_id"`

	// Find collections with the given name.
	Name string `json:"name,omitempty"`

    // Indicates whether user set optional parameter Name
    IsNameSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewListCollectionsOptions : Instantiate ListCollectionsOptions
func NewListCollectionsOptions(environmentID string) *ListCollectionsOptions {
    return &ListCollectionsOptions{
        EnvironmentID: environmentID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *ListCollectionsOptions) SetEnvironmentID(param string) *ListCollectionsOptions {
    options.EnvironmentID = param
    return options
}

// SetName : Allow user to set Name
func (options *ListCollectionsOptions) SetName(param string) *ListCollectionsOptions {
    options.Name = param
    options.IsNameSet = true
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
	EnvironmentID string `json:"environment_id"`

	// Find configurations with the given name.
	Name string `json:"name,omitempty"`

    // Indicates whether user set optional parameter Name
    IsNameSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewListConfigurationsOptions : Instantiate ListConfigurationsOptions
func NewListConfigurationsOptions(environmentID string) *ListConfigurationsOptions {
    return &ListConfigurationsOptions{
        EnvironmentID: environmentID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *ListConfigurationsOptions) SetEnvironmentID(param string) *ListConfigurationsOptions {
    options.EnvironmentID = param
    return options
}

// SetName : Allow user to set Name
func (options *ListConfigurationsOptions) SetName(param string) *ListConfigurationsOptions {
    options.Name = param
    options.IsNameSet = true
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
	EnvironmentID string `json:"environment_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewListCredentialsOptions : Instantiate ListCredentialsOptions
func NewListCredentialsOptions(environmentID string) *ListCredentialsOptions {
    return &ListCredentialsOptions{
        EnvironmentID: environmentID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *ListCredentialsOptions) SetEnvironmentID(param string) *ListCredentialsOptions {
    options.EnvironmentID = param
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
	Name string `json:"name,omitempty"`

    // Indicates whether user set optional parameter Name
    IsNameSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewListEnvironmentsOptions : Instantiate ListEnvironmentsOptions
func NewListEnvironmentsOptions() *ListEnvironmentsOptions {
    return &ListEnvironmentsOptions{}
}

// SetName : Allow user to set Name
func (options *ListEnvironmentsOptions) SetName(param string) *ListEnvironmentsOptions {
    options.Name = param
    options.IsNameSet = true
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
	EnvironmentID string `json:"environment_id"`

	// The ID of the collection.
	CollectionID string `json:"collection_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewListExpansionsOptions : Instantiate ListExpansionsOptions
func NewListExpansionsOptions(environmentID string, collectionID string) *ListExpansionsOptions {
    return &ListExpansionsOptions{
        EnvironmentID: environmentID,
        CollectionID: collectionID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *ListExpansionsOptions) SetEnvironmentID(param string) *ListExpansionsOptions {
    options.EnvironmentID = param
    return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *ListExpansionsOptions) SetCollectionID(param string) *ListExpansionsOptions {
    options.CollectionID = param
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
	EnvironmentID string `json:"environment_id"`

	// A comma-separated list of collection IDs to be queried against.
	CollectionIds []string `json:"collection_ids"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewListFieldsOptions : Instantiate ListFieldsOptions
func NewListFieldsOptions(environmentID string, collectionIds []string) *ListFieldsOptions {
    return &ListFieldsOptions{
        EnvironmentID: environmentID,
        CollectionIds: collectionIds,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *ListFieldsOptions) SetEnvironmentID(param string) *ListFieldsOptions {
    options.EnvironmentID = param
    return options
}

// SetCollectionIds : Allow user to set CollectionIds
func (options *ListFieldsOptions) SetCollectionIds(param []string) *ListFieldsOptions {
    options.CollectionIds = param
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
	EnvironmentID string `json:"environment_id"`

	// The ID of the collection.
	CollectionID string `json:"collection_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewListTrainingDataOptions : Instantiate ListTrainingDataOptions
func NewListTrainingDataOptions(environmentID string, collectionID string) *ListTrainingDataOptions {
    return &ListTrainingDataOptions{
        EnvironmentID: environmentID,
        CollectionID: collectionID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *ListTrainingDataOptions) SetEnvironmentID(param string) *ListTrainingDataOptions {
    options.EnvironmentID = param
    return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *ListTrainingDataOptions) SetCollectionID(param string) *ListTrainingDataOptions {
    options.CollectionID = param
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
	EnvironmentID string `json:"environment_id"`

	// The ID of the collection.
	CollectionID string `json:"collection_id"`

	// The ID of the query used for training.
	QueryID string `json:"query_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewListTrainingExamplesOptions : Instantiate ListTrainingExamplesOptions
func NewListTrainingExamplesOptions(environmentID string, collectionID string, queryID string) *ListTrainingExamplesOptions {
    return &ListTrainingExamplesOptions{
        EnvironmentID: environmentID,
        CollectionID: collectionID,
        QueryID: queryID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *ListTrainingExamplesOptions) SetEnvironmentID(param string) *ListTrainingExamplesOptions {
    options.EnvironmentID = param
    return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *ListTrainingExamplesOptions) SetCollectionID(param string) *ListTrainingExamplesOptions {
    options.CollectionID = param
    return options
}

// SetQueryID : Allow user to set QueryID
func (options *ListTrainingExamplesOptions) SetQueryID(param string) *ListTrainingExamplesOptions {
    options.QueryID = param
    return options
}

// SetHeaders : Allow user to set Headers
func (options *ListTrainingExamplesOptions) SetHeaders(param map[string]string) *ListTrainingExamplesOptions {
    options.Headers = param
    return options
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
	TypeVar string `json:"type,omitempty"`

	Results []AggregationResult `json:"results,omitempty"`

	// Number of matching results.
	MatchingResults int64 `json:"matching_results,omitempty"`

	// Aggregations returned by the Discovery service.
	Aggregations []QueryAggregation `json:"aggregations,omitempty"`
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
	TypeVar string `json:"type,omitempty"`
}

// QueryEntitiesOptions : The queryEntities options.
type QueryEntitiesOptions struct {

	// The ID of the environment.
	EnvironmentID string `json:"environment_id"`

	// The ID of the collection.
	CollectionID string `json:"collection_id"`

	// The entity query feature to perform. Supported features are `disambiguate` and `similar_entities`.
	Feature string `json:"feature,omitempty"`

    // Indicates whether user set optional parameter Feature
    IsFeatureSet bool

	// A text string that appears within the entity text field.
	Entity QueryEntitiesEntity `json:"entity,omitempty"`

    // Indicates whether user set optional parameter Entity
    IsEntitySet bool

	// Entity text to provide context for the queried entity and rank based on that association. For example, if you wanted to query the city of London in England your query would look for `London` with the context of `England`.
	Context QueryEntitiesContext `json:"context,omitempty"`

    // Indicates whether user set optional parameter Context
    IsContextSet bool

	// The number of results to return. The default is `10`. The maximum is `1000`.
	Count int64 `json:"count,omitempty"`

    // Indicates whether user set optional parameter Count
    IsCountSet bool

	// The number of evidence items to return for each result. The default is `0`. The maximum number of evidence items per query is 10,000.
	EvidenceCount int64 `json:"evidence_count,omitempty"`

    // Indicates whether user set optional parameter EvidenceCount
    IsEvidenceCountSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewQueryEntitiesOptions : Instantiate QueryEntitiesOptions
func NewQueryEntitiesOptions(environmentID string, collectionID string) *QueryEntitiesOptions {
    return &QueryEntitiesOptions{
        EnvironmentID: environmentID,
        CollectionID: collectionID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *QueryEntitiesOptions) SetEnvironmentID(param string) *QueryEntitiesOptions {
    options.EnvironmentID = param
    return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *QueryEntitiesOptions) SetCollectionID(param string) *QueryEntitiesOptions {
    options.CollectionID = param
    return options
}

// SetFeature : Allow user to set Feature
func (options *QueryEntitiesOptions) SetFeature(param string) *QueryEntitiesOptions {
    options.Feature = param
    options.IsFeatureSet = true
    return options
}

// SetEntity : Allow user to set Entity
func (options *QueryEntitiesOptions) SetEntity(param QueryEntitiesEntity) *QueryEntitiesOptions {
    options.Entity = param
    options.IsEntitySet = true
    return options
}

// SetContext : Allow user to set Context
func (options *QueryEntitiesOptions) SetContext(param QueryEntitiesContext) *QueryEntitiesOptions {
    options.Context = param
    options.IsContextSet = true
    return options
}

// SetCount : Allow user to set Count
func (options *QueryEntitiesOptions) SetCount(param int64) *QueryEntitiesOptions {
    options.Count = param
    options.IsCountSet = true
    return options
}

// SetEvidenceCount : Allow user to set EvidenceCount
func (options *QueryEntitiesOptions) SetEvidenceCount(param int64) *QueryEntitiesOptions {
    options.EvidenceCount = param
    options.IsEvidenceCountSet = true
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
	Text string `json:"text,omitempty"`

	// The type of the result entity.
	TypeVar string `json:"type,omitempty"`

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
	TypeVar string `json:"type,omitempty"`

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

// QueryNoticesOptions : The queryNotices options.
type QueryNoticesOptions struct {

	// The ID of the environment.
	EnvironmentID string `json:"environment_id"`

	// The ID of the collection.
	CollectionID string `json:"collection_id"`

	// A cacheable query that limits the documents returned to exclude any documents that don't mention the query content. Filter searches are better for metadata type searches and when you are trying to get a sense of concepts in the data set.
	Filter string `json:"filter,omitempty"`

    // Indicates whether user set optional parameter Filter
    IsFilterSet bool

	// A query search returns all documents in your data set with full enrichments and full text, but with the most relevant documents listed first. Use a query search when you want to find the most relevant search results. You cannot use **natural_language_query** and **query** at the same time.
	Query string `json:"query,omitempty"`

    // Indicates whether user set optional parameter Query
    IsQuerySet bool

	// A natural language query that returns relevant documents by utilizing training data and natural language understanding. You cannot use **natural_language_query** and **query** at the same time.
	NaturalLanguageQuery string `json:"natural_language_query,omitempty"`

    // Indicates whether user set optional parameter NaturalLanguageQuery
    IsNaturalLanguageQuerySet bool

	// A passages query that returns the most relevant passages from the results.
	Passages bool `json:"passages,omitempty"`

    // Indicates whether user set optional parameter Passages
    IsPassagesSet bool

	// An aggregation search uses combinations of filters and query search to return an exact answer. Aggregations are useful for building applications, because you can use them to build lists, tables, and time series. For a full list of possible aggregrations, see the Query reference.
	Aggregation string `json:"aggregation,omitempty"`

    // Indicates whether user set optional parameter Aggregation
    IsAggregationSet bool

	// Number of documents to return.
	Count int64 `json:"count,omitempty"`

    // Indicates whether user set optional parameter Count
    IsCountSet bool

	// A comma separated list of the portion of the document hierarchy to return.
	ReturnFields []string `json:"return,omitempty"`

    // Indicates whether user set optional parameter ReturnFields
    IsReturnFieldsSet bool

	// The number of query results to skip at the beginning. For example, if the total number of results that are returned is 10, and the offset is 8, it returns the last two results.
	Offset int64 `json:"offset,omitempty"`

    // Indicates whether user set optional parameter Offset
    IsOffsetSet bool

	// A comma separated list of fields in the document to sort on. You can optionally specify a sort direction by prefixing the field with `-` for descending or `+` for ascending. Ascending is the default sort direction if no prefix is specified.
	Sort []string `json:"sort,omitempty"`

    // Indicates whether user set optional parameter Sort
    IsSortSet bool

	// When true a highlight field is returned for each result which contains the fields that match the query with `<em></em>` tags around the matching query terms. Defaults to false.
	Highlight bool `json:"highlight,omitempty"`

    // Indicates whether user set optional parameter Highlight
    IsHighlightSet bool

	// A comma-separated list of fields that passages are drawn from. If this parameter not specified, then all top-level fields are included.
	PassagesFields []string `json:"passages.fields,omitempty"`

    // Indicates whether user set optional parameter PassagesFields
    IsPassagesFieldsSet bool

	// The maximum number of passages to return. The search returns fewer passages if the requested total is not found. The default is `10`. The maximum is `100`.
	PassagesCount int64 `json:"passages.count,omitempty"`

    // Indicates whether user set optional parameter PassagesCount
    IsPassagesCountSet bool

	// The approximate number of characters that any one passage will have. The default is `400`. The minimum is `50`. The maximum is `2000`.
	PassagesCharacters int64 `json:"passages.characters,omitempty"`

    // Indicates whether user set optional parameter PassagesCharacters
    IsPassagesCharactersSet bool

	// When specified, duplicate results based on the field specified are removed from the returned results. Duplicate comparison is limited to the current query only, **offset** is not considered. This parameter is currently Beta functionality.
	DeduplicateField string `json:"deduplicate.field,omitempty"`

    // Indicates whether user set optional parameter DeduplicateField
    IsDeduplicateFieldSet bool

	// When `true`, results are returned based on their similarity to the document IDs specified in the **similar.document_ids** parameter.
	Similar bool `json:"similar,omitempty"`

    // Indicates whether user set optional parameter Similar
    IsSimilarSet bool

	// A comma-separated list of document IDs that will be used to find similar documents. **Note:** If the **natural_language_query** parameter is also specified, it will be used to expand the scope of the document similarity search to include the natural language query. Other query parameters, such as **filter** and **query** are subsequently applied and reduce the query scope.
	SimilarDocumentIds []string `json:"similar.document_ids,omitempty"`

    // Indicates whether user set optional parameter SimilarDocumentIds
    IsSimilarDocumentIdsSet bool

	// A comma-separated list of field names that will be used as a basis for comparison to identify similar documents. If not specified, the entire document is used for comparison.
	SimilarFields []string `json:"similar.fields,omitempty"`

    // Indicates whether user set optional parameter SimilarFields
    IsSimilarFieldsSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewQueryNoticesOptions : Instantiate QueryNoticesOptions
func NewQueryNoticesOptions(environmentID string, collectionID string) *QueryNoticesOptions {
    return &QueryNoticesOptions{
        EnvironmentID: environmentID,
        CollectionID: collectionID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *QueryNoticesOptions) SetEnvironmentID(param string) *QueryNoticesOptions {
    options.EnvironmentID = param
    return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *QueryNoticesOptions) SetCollectionID(param string) *QueryNoticesOptions {
    options.CollectionID = param
    return options
}

// SetFilter : Allow user to set Filter
func (options *QueryNoticesOptions) SetFilter(param string) *QueryNoticesOptions {
    options.Filter = param
    options.IsFilterSet = true
    return options
}

// SetQuery : Allow user to set Query
func (options *QueryNoticesOptions) SetQuery(param string) *QueryNoticesOptions {
    options.Query = param
    options.IsQuerySet = true
    return options
}

// SetNaturalLanguageQuery : Allow user to set NaturalLanguageQuery
func (options *QueryNoticesOptions) SetNaturalLanguageQuery(param string) *QueryNoticesOptions {
    options.NaturalLanguageQuery = param
    options.IsNaturalLanguageQuerySet = true
    return options
}

// SetPassages : Allow user to set Passages
func (options *QueryNoticesOptions) SetPassages(param bool) *QueryNoticesOptions {
    options.Passages = param
    options.IsPassagesSet = true
    return options
}

// SetAggregation : Allow user to set Aggregation
func (options *QueryNoticesOptions) SetAggregation(param string) *QueryNoticesOptions {
    options.Aggregation = param
    options.IsAggregationSet = true
    return options
}

// SetCount : Allow user to set Count
func (options *QueryNoticesOptions) SetCount(param int64) *QueryNoticesOptions {
    options.Count = param
    options.IsCountSet = true
    return options
}

// SetReturnFields : Allow user to set ReturnFields
func (options *QueryNoticesOptions) SetReturnFields(param []string) *QueryNoticesOptions {
    options.ReturnFields = param
    options.IsReturnFieldsSet = true
    return options
}

// SetOffset : Allow user to set Offset
func (options *QueryNoticesOptions) SetOffset(param int64) *QueryNoticesOptions {
    options.Offset = param
    options.IsOffsetSet = true
    return options
}

// SetSort : Allow user to set Sort
func (options *QueryNoticesOptions) SetSort(param []string) *QueryNoticesOptions {
    options.Sort = param
    options.IsSortSet = true
    return options
}

// SetHighlight : Allow user to set Highlight
func (options *QueryNoticesOptions) SetHighlight(param bool) *QueryNoticesOptions {
    options.Highlight = param
    options.IsHighlightSet = true
    return options
}

// SetPassagesFields : Allow user to set PassagesFields
func (options *QueryNoticesOptions) SetPassagesFields(param []string) *QueryNoticesOptions {
    options.PassagesFields = param
    options.IsPassagesFieldsSet = true
    return options
}

// SetPassagesCount : Allow user to set PassagesCount
func (options *QueryNoticesOptions) SetPassagesCount(param int64) *QueryNoticesOptions {
    options.PassagesCount = param
    options.IsPassagesCountSet = true
    return options
}

// SetPassagesCharacters : Allow user to set PassagesCharacters
func (options *QueryNoticesOptions) SetPassagesCharacters(param int64) *QueryNoticesOptions {
    options.PassagesCharacters = param
    options.IsPassagesCharactersSet = true
    return options
}

// SetDeduplicateField : Allow user to set DeduplicateField
func (options *QueryNoticesOptions) SetDeduplicateField(param string) *QueryNoticesOptions {
    options.DeduplicateField = param
    options.IsDeduplicateFieldSet = true
    return options
}

// SetSimilar : Allow user to set Similar
func (options *QueryNoticesOptions) SetSimilar(param bool) *QueryNoticesOptions {
    options.Similar = param
    options.IsSimilarSet = true
    return options
}

// SetSimilarDocumentIds : Allow user to set SimilarDocumentIds
func (options *QueryNoticesOptions) SetSimilarDocumentIds(param []string) *QueryNoticesOptions {
    options.SimilarDocumentIds = param
    options.IsSimilarDocumentIdsSet = true
    return options
}

// SetSimilarFields : Allow user to set SimilarFields
func (options *QueryNoticesOptions) SetSimilarFields(param []string) *QueryNoticesOptions {
    options.SimilarFields = param
    options.IsSimilarFieldsSet = true
    return options
}

// SetHeaders : Allow user to set Headers
func (options *QueryNoticesOptions) SetHeaders(param map[string]string) *QueryNoticesOptions {
    options.Headers = param
    return options
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

// QueryOptions : The query options.
type QueryOptions struct {

	// The ID of the environment.
	EnvironmentID string `json:"environment_id"`

	// The ID of the collection.
	CollectionID string `json:"collection_id"`

	// A cacheable query that limits the documents returned to exclude any documents that don't mention the query content. Filter searches are better for metadata type searches and when you are trying to get a sense of concepts in the data set.
	Filter string `json:"filter,omitempty"`

    // Indicates whether user set optional parameter Filter
    IsFilterSet bool

	// A query search returns all documents in your data set with full enrichments and full text, but with the most relevant documents listed first. Use a query search when you want to find the most relevant search results. You cannot use **natural_language_query** and **query** at the same time.
	Query string `json:"query,omitempty"`

    // Indicates whether user set optional parameter Query
    IsQuerySet bool

	// A natural language query that returns relevant documents by utilizing training data and natural language understanding. You cannot use **natural_language_query** and **query** at the same time.
	NaturalLanguageQuery string `json:"natural_language_query,omitempty"`

    // Indicates whether user set optional parameter NaturalLanguageQuery
    IsNaturalLanguageQuerySet bool

	// A passages query that returns the most relevant passages from the results.
	Passages bool `json:"passages,omitempty"`

    // Indicates whether user set optional parameter Passages
    IsPassagesSet bool

	// An aggregation search uses combinations of filters and query search to return an exact answer. Aggregations are useful for building applications, because you can use them to build lists, tables, and time series. For a full list of possible aggregrations, see the Query reference.
	Aggregation string `json:"aggregation,omitempty"`

    // Indicates whether user set optional parameter Aggregation
    IsAggregationSet bool

	// Number of documents to return.
	Count int64 `json:"count,omitempty"`

    // Indicates whether user set optional parameter Count
    IsCountSet bool

	// A comma separated list of the portion of the document hierarchy to return.
	ReturnFields []string `json:"return,omitempty"`

    // Indicates whether user set optional parameter ReturnFields
    IsReturnFieldsSet bool

	// The number of query results to skip at the beginning. For example, if the total number of results that are returned is 10, and the offset is 8, it returns the last two results.
	Offset int64 `json:"offset,omitempty"`

    // Indicates whether user set optional parameter Offset
    IsOffsetSet bool

	// A comma separated list of fields in the document to sort on. You can optionally specify a sort direction by prefixing the field with `-` for descending or `+` for ascending. Ascending is the default sort direction if no prefix is specified.
	Sort []string `json:"sort,omitempty"`

    // Indicates whether user set optional parameter Sort
    IsSortSet bool

	// When true a highlight field is returned for each result which contains the fields that match the query with `<em></em>` tags around the matching query terms. Defaults to false.
	Highlight bool `json:"highlight,omitempty"`

    // Indicates whether user set optional parameter Highlight
    IsHighlightSet bool

	// A comma-separated list of fields that passages are drawn from. If this parameter not specified, then all top-level fields are included.
	PassagesFields []string `json:"passages.fields,omitempty"`

    // Indicates whether user set optional parameter PassagesFields
    IsPassagesFieldsSet bool

	// The maximum number of passages to return. The search returns fewer passages if the requested total is not found. The default is `10`. The maximum is `100`.
	PassagesCount int64 `json:"passages.count,omitempty"`

    // Indicates whether user set optional parameter PassagesCount
    IsPassagesCountSet bool

	// The approximate number of characters that any one passage will have. The default is `400`. The minimum is `50`. The maximum is `2000`.
	PassagesCharacters int64 `json:"passages.characters,omitempty"`

    // Indicates whether user set optional parameter PassagesCharacters
    IsPassagesCharactersSet bool

	// When `true` and used with a Watson Discovery News collection, duplicate results (based on the contents of the **title** field) are removed. Duplicate comparison is limited to the current query only; **offset** is not considered. This parameter is currently Beta functionality.
	Deduplicate bool `json:"deduplicate,omitempty"`

    // Indicates whether user set optional parameter Deduplicate
    IsDeduplicateSet bool

	// When specified, duplicate results based on the field specified are removed from the returned results. Duplicate comparison is limited to the current query only, **offset** is not considered. This parameter is currently Beta functionality.
	DeduplicateField string `json:"deduplicate.field,omitempty"`

    // Indicates whether user set optional parameter DeduplicateField
    IsDeduplicateFieldSet bool

	// When `true`, results are returned based on their similarity to the document IDs specified in the **similar.document_ids** parameter.
	Similar bool `json:"similar,omitempty"`

    // Indicates whether user set optional parameter Similar
    IsSimilarSet bool

	// A comma-separated list of document IDs that will be used to find similar documents. **Note:** If the **natural_language_query** parameter is also specified, it will be used to expand the scope of the document similarity search to include the natural language query. Other query parameters, such as **filter** and **query** are subsequently applied and reduce the query scope.
	SimilarDocumentIds []string `json:"similar.document_ids,omitempty"`

    // Indicates whether user set optional parameter SimilarDocumentIds
    IsSimilarDocumentIdsSet bool

	// A comma-separated list of field names that will be used as a basis for comparison to identify similar documents. If not specified, the entire document is used for comparison.
	SimilarFields []string `json:"similar.fields,omitempty"`

    // Indicates whether user set optional parameter SimilarFields
    IsSimilarFieldsSet bool

	// If `true`, queries are not stored in the Discovery **Logs** endpoint.
	LoggingOptOut bool `json:"X-Watson-Logging-Opt-Out,omitempty"`

    // Indicates whether user set optional parameter LoggingOptOut
    IsLoggingOptOutSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewQueryOptions : Instantiate QueryOptions
func NewQueryOptions(environmentID string, collectionID string) *QueryOptions {
    return &QueryOptions{
        EnvironmentID: environmentID,
        CollectionID: collectionID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *QueryOptions) SetEnvironmentID(param string) *QueryOptions {
    options.EnvironmentID = param
    return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *QueryOptions) SetCollectionID(param string) *QueryOptions {
    options.CollectionID = param
    return options
}

// SetFilter : Allow user to set Filter
func (options *QueryOptions) SetFilter(param string) *QueryOptions {
    options.Filter = param
    options.IsFilterSet = true
    return options
}

// SetQuery : Allow user to set Query
func (options *QueryOptions) SetQuery(param string) *QueryOptions {
    options.Query = param
    options.IsQuerySet = true
    return options
}

// SetNaturalLanguageQuery : Allow user to set NaturalLanguageQuery
func (options *QueryOptions) SetNaturalLanguageQuery(param string) *QueryOptions {
    options.NaturalLanguageQuery = param
    options.IsNaturalLanguageQuerySet = true
    return options
}

// SetPassages : Allow user to set Passages
func (options *QueryOptions) SetPassages(param bool) *QueryOptions {
    options.Passages = param
    options.IsPassagesSet = true
    return options
}

// SetAggregation : Allow user to set Aggregation
func (options *QueryOptions) SetAggregation(param string) *QueryOptions {
    options.Aggregation = param
    options.IsAggregationSet = true
    return options
}

// SetCount : Allow user to set Count
func (options *QueryOptions) SetCount(param int64) *QueryOptions {
    options.Count = param
    options.IsCountSet = true
    return options
}

// SetReturnFields : Allow user to set ReturnFields
func (options *QueryOptions) SetReturnFields(param []string) *QueryOptions {
    options.ReturnFields = param
    options.IsReturnFieldsSet = true
    return options
}

// SetOffset : Allow user to set Offset
func (options *QueryOptions) SetOffset(param int64) *QueryOptions {
    options.Offset = param
    options.IsOffsetSet = true
    return options
}

// SetSort : Allow user to set Sort
func (options *QueryOptions) SetSort(param []string) *QueryOptions {
    options.Sort = param
    options.IsSortSet = true
    return options
}

// SetHighlight : Allow user to set Highlight
func (options *QueryOptions) SetHighlight(param bool) *QueryOptions {
    options.Highlight = param
    options.IsHighlightSet = true
    return options
}

// SetPassagesFields : Allow user to set PassagesFields
func (options *QueryOptions) SetPassagesFields(param []string) *QueryOptions {
    options.PassagesFields = param
    options.IsPassagesFieldsSet = true
    return options
}

// SetPassagesCount : Allow user to set PassagesCount
func (options *QueryOptions) SetPassagesCount(param int64) *QueryOptions {
    options.PassagesCount = param
    options.IsPassagesCountSet = true
    return options
}

// SetPassagesCharacters : Allow user to set PassagesCharacters
func (options *QueryOptions) SetPassagesCharacters(param int64) *QueryOptions {
    options.PassagesCharacters = param
    options.IsPassagesCharactersSet = true
    return options
}

// SetDeduplicate : Allow user to set Deduplicate
func (options *QueryOptions) SetDeduplicate(param bool) *QueryOptions {
    options.Deduplicate = param
    options.IsDeduplicateSet = true
    return options
}

// SetDeduplicateField : Allow user to set DeduplicateField
func (options *QueryOptions) SetDeduplicateField(param string) *QueryOptions {
    options.DeduplicateField = param
    options.IsDeduplicateFieldSet = true
    return options
}

// SetSimilar : Allow user to set Similar
func (options *QueryOptions) SetSimilar(param bool) *QueryOptions {
    options.Similar = param
    options.IsSimilarSet = true
    return options
}

// SetSimilarDocumentIds : Allow user to set SimilarDocumentIds
func (options *QueryOptions) SetSimilarDocumentIds(param []string) *QueryOptions {
    options.SimilarDocumentIds = param
    options.IsSimilarDocumentIdsSet = true
    return options
}

// SetSimilarFields : Allow user to set SimilarFields
func (options *QueryOptions) SetSimilarFields(param []string) *QueryOptions {
    options.SimilarFields = param
    options.IsSimilarFieldsSet = true
    return options
}

// SetLoggingOptOut : Allow user to set LoggingOptOut
func (options *QueryOptions) SetLoggingOptOut(param bool) *QueryOptions {
    options.LoggingOptOut = param
    options.IsLoggingOptOutSet = true
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

// QueryRelationsArgument : QueryRelationsArgument struct
type QueryRelationsArgument struct {

	Entities []QueryEntitiesEntity `json:"entities,omitempty"`
}

// QueryRelationsEntity : QueryRelationsEntity struct
type QueryRelationsEntity struct {

	// Entity text content.
	Text string `json:"text,omitempty"`

	// The type of the specified entity.
	TypeVar string `json:"type,omitempty"`

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

// QueryRelationsOptions : The queryRelations options.
type QueryRelationsOptions struct {

	// The ID of the environment.
	EnvironmentID string `json:"environment_id"`

	// The ID of the collection.
	CollectionID string `json:"collection_id"`

	// An array of entities to find relationships for.
	Entities []QueryRelationsEntity `json:"entities,omitempty"`

    // Indicates whether user set optional parameter Entities
    IsEntitiesSet bool

	// Entity text to provide context for the queried entity and rank based on that association. For example, if you wanted to query the city of London in England your query would look for `London` with the context of `England`.
	Context QueryEntitiesContext `json:"context,omitempty"`

    // Indicates whether user set optional parameter Context
    IsContextSet bool

	// The sorting method for the relationships, can be `score` or `frequency`. `frequency` is the number of unique times each entity is identified. The default is `score`.
	Sort string `json:"sort,omitempty"`

    // Indicates whether user set optional parameter Sort
    IsSortSet bool

	// Filters to apply to the relationship query.
	Filter QueryRelationsFilter `json:"filter,omitempty"`

    // Indicates whether user set optional parameter Filter
    IsFilterSet bool

	// The number of results to return. The default is `10`. The maximum is `1000`.
	Count int64 `json:"count,omitempty"`

    // Indicates whether user set optional parameter Count
    IsCountSet bool

	// The number of evidence items to return for each result. The default is `0`. The maximum number of evidence items per query is 10,000.
	EvidenceCount int64 `json:"evidence_count,omitempty"`

    // Indicates whether user set optional parameter EvidenceCount
    IsEvidenceCountSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewQueryRelationsOptions : Instantiate QueryRelationsOptions
func NewQueryRelationsOptions(environmentID string, collectionID string) *QueryRelationsOptions {
    return &QueryRelationsOptions{
        EnvironmentID: environmentID,
        CollectionID: collectionID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *QueryRelationsOptions) SetEnvironmentID(param string) *QueryRelationsOptions {
    options.EnvironmentID = param
    return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *QueryRelationsOptions) SetCollectionID(param string) *QueryRelationsOptions {
    options.CollectionID = param
    return options
}

// SetEntities : Allow user to set Entities
func (options *QueryRelationsOptions) SetEntities(param []QueryRelationsEntity) *QueryRelationsOptions {
    options.Entities = param
    options.IsEntitiesSet = true
    return options
}

// SetContext : Allow user to set Context
func (options *QueryRelationsOptions) SetContext(param QueryEntitiesContext) *QueryRelationsOptions {
    options.Context = param
    options.IsContextSet = true
    return options
}

// SetSort : Allow user to set Sort
func (options *QueryRelationsOptions) SetSort(param string) *QueryRelationsOptions {
    options.Sort = param
    options.IsSortSet = true
    return options
}

// SetFilter : Allow user to set Filter
func (options *QueryRelationsOptions) SetFilter(param QueryRelationsFilter) *QueryRelationsOptions {
    options.Filter = param
    options.IsFilterSet = true
    return options
}

// SetCount : Allow user to set Count
func (options *QueryRelationsOptions) SetCount(param int64) *QueryRelationsOptions {
    options.Count = param
    options.IsCountSet = true
    return options
}

// SetEvidenceCount : Allow user to set EvidenceCount
func (options *QueryRelationsOptions) SetEvidenceCount(param int64) *QueryRelationsOptions {
    options.EvidenceCount = param
    options.IsEvidenceCountSet = true
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
	TypeVar string `json:"type,omitempty"`

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

	// The session token for this query. The session token can be used to add events associated with this query to the query and event log. **Important:** Session tokens are case sensitive.
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

	// An unbounded measure of the relevance of a particular result, dependent on the query and matching document. A higher score indicates a greater match to the query parameters.
	Score float64 `json:"score,omitempty"`

	// The confidence score for the given result. Calculated based on how relevant the result is estimated to be, compared to a trained relevancy model. confidence can range from `0.0` to `1.0`. The higher the number, the more relevant the document.
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
	TypeVar string `json:"type,omitempty"`

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

// TestConfigurationInEnvironmentOptions : The testConfigurationInEnvironment options.
type TestConfigurationInEnvironmentOptions struct {

	// The ID of the environment.
	EnvironmentID string `json:"environment_id"`

	// The configuration to use to process the document. If this part is provided, then the provided configuration is used to process the document. If the **configuration_id** is also provided (both are present at the same time), then request is rejected. The maximum supported configuration size is 1 MB. Configuration parts larger than 1 MB are rejected. See the `GET /configurations/{configuration_id}` operation for an example configuration.
	Configuration string `json:"configuration,omitempty"`

    // Indicates whether user set optional parameter Configuration
    IsConfigurationSet bool

	// Specify to only run the input document through the given step instead of running the input document through the entire ingestion workflow. Valid values are `convert`, `enrich`, and `normalize`.
	Step string `json:"step,omitempty"`

    // Indicates whether user set optional parameter Step
    IsStepSet bool

	// The ID of the configuration to use to process the document. If the **configuration** form part is also provided (both are present at the same time), then the request will be rejected.
	ConfigurationID string `json:"configuration_id,omitempty"`

    // Indicates whether user set optional parameter ConfigurationID
    IsConfigurationIDSet bool

	// The content of the document to ingest. The maximum supported file size is 50 megabytes. Files larger than 50 megabytes is rejected.
	File os.File `json:"file,omitempty"`

    // Indicates whether user set optional parameter File
    IsFileSet bool

	// If you're using the Data Crawler to upload your documents, you can test a document against the type of metadata that the Data Crawler might send. The maximum supported metadata file size is 1 MB. Metadata parts larger than 1 MB are rejected. Example:  ``` { "Creator": "Johnny Appleseed", "Subject": "Apples" } ```.
	Metadata string `json:"metadata,omitempty"`

    // Indicates whether user set optional parameter Metadata
    IsMetadataSet bool

	// The content type of File.
	FileContentType string `json:"file_content_type,omitempty"`

    // Indicates whether user set optional parameter FileContentType
    IsFileContentTypeSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewTestConfigurationInEnvironmentOptions : Instantiate TestConfigurationInEnvironmentOptions
func NewTestConfigurationInEnvironmentOptions(environmentID string) *TestConfigurationInEnvironmentOptions {
    return &TestConfigurationInEnvironmentOptions{
        EnvironmentID: environmentID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *TestConfigurationInEnvironmentOptions) SetEnvironmentID(param string) *TestConfigurationInEnvironmentOptions {
    options.EnvironmentID = param
    return options
}

// SetConfiguration : Allow user to set Configuration
func (options *TestConfigurationInEnvironmentOptions) SetConfiguration(param string) *TestConfigurationInEnvironmentOptions {
    options.Configuration = param
    options.IsConfigurationSet = true
    return options
}

// SetStep : Allow user to set Step
func (options *TestConfigurationInEnvironmentOptions) SetStep(param string) *TestConfigurationInEnvironmentOptions {
    options.Step = param
    options.IsStepSet = true
    return options
}

// SetConfigurationID : Allow user to set ConfigurationID
func (options *TestConfigurationInEnvironmentOptions) SetConfigurationID(param string) *TestConfigurationInEnvironmentOptions {
    options.ConfigurationID = param
    options.IsConfigurationIDSet = true
    return options
}

// SetFile : Allow user to set File
func (options *TestConfigurationInEnvironmentOptions) SetFile(param os.File) *TestConfigurationInEnvironmentOptions {
    options.File = param
    options.IsFileSet = true
    return options
}

// SetMetadata : Allow user to set Metadata
func (options *TestConfigurationInEnvironmentOptions) SetMetadata(param string) *TestConfigurationInEnvironmentOptions {
    options.Metadata = param
    options.IsMetadataSet = true
    return options
}

// SetFileContentType : Allow user to set FileContentType
func (options *TestConfigurationInEnvironmentOptions) SetFileContentType(param string) *TestConfigurationInEnvironmentOptions {
    options.FileContentType = param
    options.IsFileContentTypeSet = true
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

// UpdateCollectionOptions : The updateCollection options.
type UpdateCollectionOptions struct {

	// The ID of the environment.
	EnvironmentID string `json:"environment_id"`

	// The ID of the collection.
	CollectionID string `json:"collection_id"`

	// The name of the collection.
	Name string `json:"name,omitempty"`

    // Indicates whether user set optional parameter Name
    IsNameSet bool

	// A description of the collection.
	Description string `json:"description,omitempty"`

    // Indicates whether user set optional parameter Description
    IsDescriptionSet bool

	// The ID of the configuration in which the collection is to be updated.
	ConfigurationID string `json:"configuration_id,omitempty"`

    // Indicates whether user set optional parameter ConfigurationID
    IsConfigurationIDSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewUpdateCollectionOptions : Instantiate UpdateCollectionOptions
func NewUpdateCollectionOptions(environmentID string, collectionID string) *UpdateCollectionOptions {
    return &UpdateCollectionOptions{
        EnvironmentID: environmentID,
        CollectionID: collectionID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *UpdateCollectionOptions) SetEnvironmentID(param string) *UpdateCollectionOptions {
    options.EnvironmentID = param
    return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *UpdateCollectionOptions) SetCollectionID(param string) *UpdateCollectionOptions {
    options.CollectionID = param
    return options
}

// SetName : Allow user to set Name
func (options *UpdateCollectionOptions) SetName(param string) *UpdateCollectionOptions {
    options.Name = param
    options.IsNameSet = true
    return options
}

// SetDescription : Allow user to set Description
func (options *UpdateCollectionOptions) SetDescription(param string) *UpdateCollectionOptions {
    options.Description = param
    options.IsDescriptionSet = true
    return options
}

// SetConfigurationID : Allow user to set ConfigurationID
func (options *UpdateCollectionOptions) SetConfigurationID(param string) *UpdateCollectionOptions {
    options.ConfigurationID = param
    options.IsConfigurationIDSet = true
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
	EnvironmentID string `json:"environment_id"`

	// The ID of the configuration.
	ConfigurationID string `json:"configuration_id"`

	// The name of the configuration.
	Name string `json:"name,omitempty"`

    // Indicates whether user set optional parameter Name
    IsNameSet bool

	// The description of the configuration, if available.
	Description string `json:"description,omitempty"`

    // Indicates whether user set optional parameter Description
    IsDescriptionSet bool

	// The document conversion settings for the configuration.
	Conversions Conversions `json:"conversions,omitempty"`

    // Indicates whether user set optional parameter Conversions
    IsConversionsSet bool

	// An array of document enrichment settings for the configuration.
	Enrichments []Enrichment `json:"enrichments,omitempty"`

    // Indicates whether user set optional parameter Enrichments
    IsEnrichmentsSet bool

	// Defines operations that can be used to transform the final output JSON into a normalized form. Operations are executed in the order that they appear in the array.
	Normalizations []NormalizationOperation `json:"normalizations,omitempty"`

    // Indicates whether user set optional parameter Normalizations
    IsNormalizationsSet bool

	// Object containing source parameters for the configuration.
	Source Source `json:"source,omitempty"`

    // Indicates whether user set optional parameter Source
    IsSourceSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewUpdateConfigurationOptions : Instantiate UpdateConfigurationOptions
func NewUpdateConfigurationOptions(environmentID string, configurationID string) *UpdateConfigurationOptions {
    return &UpdateConfigurationOptions{
        EnvironmentID: environmentID,
        ConfigurationID: configurationID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *UpdateConfigurationOptions) SetEnvironmentID(param string) *UpdateConfigurationOptions {
    options.EnvironmentID = param
    return options
}

// SetConfigurationID : Allow user to set ConfigurationID
func (options *UpdateConfigurationOptions) SetConfigurationID(param string) *UpdateConfigurationOptions {
    options.ConfigurationID = param
    return options
}

// SetName : Allow user to set Name
func (options *UpdateConfigurationOptions) SetName(param string) *UpdateConfigurationOptions {
    options.Name = param
    options.IsNameSet = true
    return options
}

// SetDescription : Allow user to set Description
func (options *UpdateConfigurationOptions) SetDescription(param string) *UpdateConfigurationOptions {
    options.Description = param
    options.IsDescriptionSet = true
    return options
}

// SetConversions : Allow user to set Conversions
func (options *UpdateConfigurationOptions) SetConversions(param Conversions) *UpdateConfigurationOptions {
    options.Conversions = param
    options.IsConversionsSet = true
    return options
}

// SetEnrichments : Allow user to set Enrichments
func (options *UpdateConfigurationOptions) SetEnrichments(param []Enrichment) *UpdateConfigurationOptions {
    options.Enrichments = param
    options.IsEnrichmentsSet = true
    return options
}

// SetNormalizations : Allow user to set Normalizations
func (options *UpdateConfigurationOptions) SetNormalizations(param []NormalizationOperation) *UpdateConfigurationOptions {
    options.Normalizations = param
    options.IsNormalizationsSet = true
    return options
}

// SetSource : Allow user to set Source
func (options *UpdateConfigurationOptions) SetSource(param Source) *UpdateConfigurationOptions {
    options.Source = param
    options.IsSourceSet = true
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
	EnvironmentID string `json:"environment_id"`

	// The unique identifier for a set of source credentials.
	CredentialID string `json:"credential_id"`

	// The source that this credentials object connects to. -  `box` indicates the credentials are used to connect an instance of Enterprise Box. -  `salesforce` indicates the credentials are used to connect to Salesforce. -  `sharepoint` indicates the credentials are used to connect to Microsoft SharePoint Online.
	SourceType string `json:"source_type,omitempty"`

    // Indicates whether user set optional parameter SourceType
    IsSourceTypeSet bool

	// Object containing details of the stored credentials. Obtain credentials for your source from the administrator of the source.
	CredentialDetails CredentialDetails `json:"credential_details,omitempty"`

    // Indicates whether user set optional parameter CredentialDetails
    IsCredentialDetailsSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewUpdateCredentialsOptions : Instantiate UpdateCredentialsOptions
func NewUpdateCredentialsOptions(environmentID string, credentialID string) *UpdateCredentialsOptions {
    return &UpdateCredentialsOptions{
        EnvironmentID: environmentID,
        CredentialID: credentialID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *UpdateCredentialsOptions) SetEnvironmentID(param string) *UpdateCredentialsOptions {
    options.EnvironmentID = param
    return options
}

// SetCredentialID : Allow user to set CredentialID
func (options *UpdateCredentialsOptions) SetCredentialID(param string) *UpdateCredentialsOptions {
    options.CredentialID = param
    return options
}

// SetSourceType : Allow user to set SourceType
func (options *UpdateCredentialsOptions) SetSourceType(param string) *UpdateCredentialsOptions {
    options.SourceType = param
    options.IsSourceTypeSet = true
    return options
}

// SetCredentialDetails : Allow user to set CredentialDetails
func (options *UpdateCredentialsOptions) SetCredentialDetails(param CredentialDetails) *UpdateCredentialsOptions {
    options.CredentialDetails = param
    options.IsCredentialDetailsSet = true
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
	EnvironmentID string `json:"environment_id"`

	// The ID of the collection.
	CollectionID string `json:"collection_id"`

	// The ID of the document.
	DocumentID string `json:"document_id"`

	// The content of the document to ingest. The maximum supported file size is 50 megabytes. Files larger than 50 megabytes is rejected.
	File os.File `json:"file,omitempty"`

    // Indicates whether user set optional parameter File
    IsFileSet bool

	// If you're using the Data Crawler to upload your documents, you can test a document against the type of metadata that the Data Crawler might send. The maximum supported metadata file size is 1 MB. Metadata parts larger than 1 MB are rejected. Example:  ``` { "Creator": "Johnny Appleseed", "Subject": "Apples" } ```.
	Metadata string `json:"metadata,omitempty"`

    // Indicates whether user set optional parameter Metadata
    IsMetadataSet bool

	// The content type of File.
	FileContentType string `json:"file_content_type,omitempty"`

    // Indicates whether user set optional parameter FileContentType
    IsFileContentTypeSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewUpdateDocumentOptions : Instantiate UpdateDocumentOptions
func NewUpdateDocumentOptions(environmentID string, collectionID string, documentID string) *UpdateDocumentOptions {
    return &UpdateDocumentOptions{
        EnvironmentID: environmentID,
        CollectionID: collectionID,
        DocumentID: documentID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *UpdateDocumentOptions) SetEnvironmentID(param string) *UpdateDocumentOptions {
    options.EnvironmentID = param
    return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *UpdateDocumentOptions) SetCollectionID(param string) *UpdateDocumentOptions {
    options.CollectionID = param
    return options
}

// SetDocumentID : Allow user to set DocumentID
func (options *UpdateDocumentOptions) SetDocumentID(param string) *UpdateDocumentOptions {
    options.DocumentID = param
    return options
}

// SetFile : Allow user to set File
func (options *UpdateDocumentOptions) SetFile(param os.File) *UpdateDocumentOptions {
    options.File = param
    options.IsFileSet = true
    return options
}

// SetMetadata : Allow user to set Metadata
func (options *UpdateDocumentOptions) SetMetadata(param string) *UpdateDocumentOptions {
    options.Metadata = param
    options.IsMetadataSet = true
    return options
}

// SetFileContentType : Allow user to set FileContentType
func (options *UpdateDocumentOptions) SetFileContentType(param string) *UpdateDocumentOptions {
    options.FileContentType = param
    options.IsFileContentTypeSet = true
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
	EnvironmentID string `json:"environment_id"`

	// Name that identifies the environment.
	Name string `json:"name,omitempty"`

    // Indicates whether user set optional parameter Name
    IsNameSet bool

	// Description of the environment.
	Description string `json:"description,omitempty"`

    // Indicates whether user set optional parameter Description
    IsDescriptionSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewUpdateEnvironmentOptions : Instantiate UpdateEnvironmentOptions
func NewUpdateEnvironmentOptions(environmentID string) *UpdateEnvironmentOptions {
    return &UpdateEnvironmentOptions{
        EnvironmentID: environmentID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *UpdateEnvironmentOptions) SetEnvironmentID(param string) *UpdateEnvironmentOptions {
    options.EnvironmentID = param
    return options
}

// SetName : Allow user to set Name
func (options *UpdateEnvironmentOptions) SetName(param string) *UpdateEnvironmentOptions {
    options.Name = param
    options.IsNameSet = true
    return options
}

// SetDescription : Allow user to set Description
func (options *UpdateEnvironmentOptions) SetDescription(param string) *UpdateEnvironmentOptions {
    options.Description = param
    options.IsDescriptionSet = true
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
	EnvironmentID string `json:"environment_id"`

	// The ID of the collection.
	CollectionID string `json:"collection_id"`

	// The ID of the query used for training.
	QueryID string `json:"query_id"`

	// The ID of the document as it is indexed.
	ExampleID string `json:"example_id"`

	CrossReference string `json:"cross_reference,omitempty"`

    // Indicates whether user set optional parameter CrossReference
    IsCrossReferenceSet bool

	Relevance int64 `json:"relevance,omitempty"`

    // Indicates whether user set optional parameter Relevance
    IsRelevanceSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewUpdateTrainingExampleOptions : Instantiate UpdateTrainingExampleOptions
func NewUpdateTrainingExampleOptions(environmentID string, collectionID string, queryID string, exampleID string) *UpdateTrainingExampleOptions {
    return &UpdateTrainingExampleOptions{
        EnvironmentID: environmentID,
        CollectionID: collectionID,
        QueryID: queryID,
        ExampleID: exampleID,
    }
}

// SetEnvironmentID : Allow user to set EnvironmentID
func (options *UpdateTrainingExampleOptions) SetEnvironmentID(param string) *UpdateTrainingExampleOptions {
    options.EnvironmentID = param
    return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *UpdateTrainingExampleOptions) SetCollectionID(param string) *UpdateTrainingExampleOptions {
    options.CollectionID = param
    return options
}

// SetQueryID : Allow user to set QueryID
func (options *UpdateTrainingExampleOptions) SetQueryID(param string) *UpdateTrainingExampleOptions {
    options.QueryID = param
    return options
}

// SetExampleID : Allow user to set ExampleID
func (options *UpdateTrainingExampleOptions) SetExampleID(param string) *UpdateTrainingExampleOptions {
    options.ExampleID = param
    return options
}

// SetCrossReference : Allow user to set CrossReference
func (options *UpdateTrainingExampleOptions) SetCrossReference(param string) *UpdateTrainingExampleOptions {
    options.CrossReference = param
    options.IsCrossReferenceSet = true
    return options
}

// SetRelevance : Allow user to set Relevance
func (options *UpdateTrainingExampleOptions) SetRelevance(param int64) *UpdateTrainingExampleOptions {
    options.Relevance = param
    options.IsRelevanceSet = true
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

	// Used to indicate that anomaly detection should be performed. Anomaly detection is used to locate unusual datapoints within a time series.
	Anomaly bool `json:"anomaly,omitempty"`
}

// TopHits : TopHits struct
type TopHits struct {

	// Number of top hits returned by the aggregation.
	Size int64 `json:"size,omitempty"`

	Hits TopHitsResults `json:"hits,omitempty"`
}
