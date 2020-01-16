/**
 * (C) Copyright IBM Corp. 2018, 2020.
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

package discoveryv1_test

import (
	"fmt"
	"github.com/IBM/go-sdk-core/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/discoveryv1"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
)

var _ = Describe(`DiscoveryV1`, func() {
	Describe(`CreateEnvironment(createEnvironmentOptions *CreateEnvironmentOptions)`, func() {
		createEnvironmentPath := "/v1/environments"
		version := "exampleString"
		bearerToken := "0ui9876453"
		name := "exampleString"
		Context(`Successfully - Create an environment`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createEnvironmentPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call CreateEnvironment`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.CreateEnvironment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				createEnvironmentOptions := testService.NewCreateEnvironmentOptions(name)
				result, response, operationErr = testService.CreateEnvironment(createEnvironmentOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListEnvironments(listEnvironmentsOptions *ListEnvironmentsOptions)`, func() {
		listEnvironmentsPath := "/v1/environments"
		version := "exampleString"
		bearerToken := "0ui9876453"
		Context(`Successfully - List environments`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listEnvironmentsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call ListEnvironments`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListEnvironments(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listEnvironmentsOptions := testService.NewListEnvironmentsOptions()
				result, response, operationErr = testService.ListEnvironments(listEnvironmentsOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetEnvironment(getEnvironmentOptions *GetEnvironmentOptions)`, func() {
		getEnvironmentPath := "/v1/environments/{environment_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		getEnvironmentPath = strings.Replace(getEnvironmentPath, "{environment_id}", environmentID, 1)
		Context(`Successfully - Get environment info`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getEnvironmentPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call GetEnvironment`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetEnvironment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getEnvironmentOptions := testService.NewGetEnvironmentOptions(environmentID)
				result, response, operationErr = testService.GetEnvironment(getEnvironmentOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateEnvironment(updateEnvironmentOptions *UpdateEnvironmentOptions)`, func() {
		updateEnvironmentPath := "/v1/environments/{environment_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		updateEnvironmentPath = strings.Replace(updateEnvironmentPath, "{environment_id}", environmentID, 1)
		Context(`Successfully - Update an environment`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateEnvironmentPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call UpdateEnvironment`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.UpdateEnvironment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				updateEnvironmentOptions := testService.NewUpdateEnvironmentOptions(environmentID)
				result, response, operationErr = testService.UpdateEnvironment(updateEnvironmentOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteEnvironment(deleteEnvironmentOptions *DeleteEnvironmentOptions)`, func() {
		deleteEnvironmentPath := "/v1/environments/{environment_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		deleteEnvironmentPath = strings.Replace(deleteEnvironmentPath, "{environment_id}", environmentID, 1)
		Context(`Successfully - Delete environment`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteEnvironmentPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"environment_id": "fake_EnvironmentID", "status": "fake_Status"}`)
			}))
			It(`Succeed to call DeleteEnvironment`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.DeleteEnvironment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				deleteEnvironmentOptions := testService.NewDeleteEnvironmentOptions(environmentID)
				result, response, operationErr = testService.DeleteEnvironment(deleteEnvironmentOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListFields(listFieldsOptions *ListFieldsOptions)`, func() {
		listFieldsPath := "/v1/environments/{environment_id}/fields"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		collectionIds := []string{}
		listFieldsPath = strings.Replace(listFieldsPath, "{environment_id}", environmentID, 1)
		Context(`Successfully - List fields across collections`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listFieldsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call ListFields`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListFields(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listFieldsOptions := testService.NewListFieldsOptions(environmentID, collectionIds)
				result, response, operationErr = testService.ListFields(listFieldsOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateConfiguration(createConfigurationOptions *CreateConfigurationOptions)`, func() {
		createConfigurationPath := "/v1/environments/{environment_id}/configurations"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		name := "exampleString"
		createConfigurationPath = strings.Replace(createConfigurationPath, "{environment_id}", environmentID, 1)
		Context(`Successfully - Add configuration`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createConfigurationPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"name": "fake_Name"}`)
			}))
			It(`Succeed to call CreateConfiguration`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.CreateConfiguration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				createConfigurationOptions := testService.NewCreateConfigurationOptions(environmentID, name)
				result, response, operationErr = testService.CreateConfiguration(createConfigurationOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListConfigurations(listConfigurationsOptions *ListConfigurationsOptions)`, func() {
		listConfigurationsPath := "/v1/environments/{environment_id}/configurations"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		listConfigurationsPath = strings.Replace(listConfigurationsPath, "{environment_id}", environmentID, 1)
		Context(`Successfully - List configurations`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listConfigurationsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call ListConfigurations`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListConfigurations(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listConfigurationsOptions := testService.NewListConfigurationsOptions(environmentID)
				result, response, operationErr = testService.ListConfigurations(listConfigurationsOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetConfiguration(getConfigurationOptions *GetConfigurationOptions)`, func() {
		getConfigurationPath := "/v1/environments/{environment_id}/configurations/{configuration_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		configurationID := "exampleString"
		getConfigurationPath = strings.Replace(getConfigurationPath, "{environment_id}", environmentID, 1)
		getConfigurationPath = strings.Replace(getConfigurationPath, "{configuration_id}", configurationID, 1)
		Context(`Successfully - Get configuration details`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getConfigurationPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"name": "fake_Name"}`)
			}))
			It(`Succeed to call GetConfiguration`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetConfiguration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getConfigurationOptions := testService.NewGetConfigurationOptions(environmentID, configurationID)
				result, response, operationErr = testService.GetConfiguration(getConfigurationOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateConfiguration(updateConfigurationOptions *UpdateConfigurationOptions)`, func() {
		updateConfigurationPath := "/v1/environments/{environment_id}/configurations/{configuration_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		configurationID := "exampleString"
		name := "exampleString"
		updateConfigurationPath = strings.Replace(updateConfigurationPath, "{environment_id}", environmentID, 1)
		updateConfigurationPath = strings.Replace(updateConfigurationPath, "{configuration_id}", configurationID, 1)
		Context(`Successfully - Update a configuration`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateConfigurationPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"name": "fake_Name"}`)
			}))
			It(`Succeed to call UpdateConfiguration`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.UpdateConfiguration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				updateConfigurationOptions := testService.NewUpdateConfigurationOptions(environmentID, configurationID, name)
				result, response, operationErr = testService.UpdateConfiguration(updateConfigurationOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteConfiguration(deleteConfigurationOptions *DeleteConfigurationOptions)`, func() {
		deleteConfigurationPath := "/v1/environments/{environment_id}/configurations/{configuration_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		configurationID := "exampleString"
		deleteConfigurationPath = strings.Replace(deleteConfigurationPath, "{environment_id}", environmentID, 1)
		deleteConfigurationPath = strings.Replace(deleteConfigurationPath, "{configuration_id}", configurationID, 1)
		Context(`Successfully - Delete a configuration`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteConfigurationPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"configuration_id": "fake_ConfigurationID", "status": "fake_Status"}`)
			}))
			It(`Succeed to call DeleteConfiguration`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.DeleteConfiguration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				deleteConfigurationOptions := testService.NewDeleteConfigurationOptions(environmentID, configurationID)
				result, response, operationErr = testService.DeleteConfiguration(deleteConfigurationOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateCollection(createCollectionOptions *CreateCollectionOptions)`, func() {
		createCollectionPath := "/v1/environments/{environment_id}/collections"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		name := "exampleString"
		createCollectionPath = strings.Replace(createCollectionPath, "{environment_id}", environmentID, 1)
		Context(`Successfully - Create a collection`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createCollectionPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call CreateCollection`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.CreateCollection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				createCollectionOptions := testService.NewCreateCollectionOptions(environmentID, name)
				result, response, operationErr = testService.CreateCollection(createCollectionOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListCollections(listCollectionsOptions *ListCollectionsOptions)`, func() {
		listCollectionsPath := "/v1/environments/{environment_id}/collections"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		listCollectionsPath = strings.Replace(listCollectionsPath, "{environment_id}", environmentID, 1)
		Context(`Successfully - List collections`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listCollectionsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call ListCollections`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListCollections(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listCollectionsOptions := testService.NewListCollectionsOptions(environmentID)
				result, response, operationErr = testService.ListCollections(listCollectionsOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetCollection(getCollectionOptions *GetCollectionOptions)`, func() {
		getCollectionPath := "/v1/environments/{environment_id}/collections/{collection_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		getCollectionPath = strings.Replace(getCollectionPath, "{environment_id}", environmentID, 1)
		getCollectionPath = strings.Replace(getCollectionPath, "{collection_id}", collectionID, 1)
		Context(`Successfully - Get collection details`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getCollectionPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call GetCollection`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetCollection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getCollectionOptions := testService.NewGetCollectionOptions(environmentID, collectionID)
				result, response, operationErr = testService.GetCollection(getCollectionOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateCollection(updateCollectionOptions *UpdateCollectionOptions)`, func() {
		updateCollectionPath := "/v1/environments/{environment_id}/collections/{collection_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		name := "exampleString"
		updateCollectionPath = strings.Replace(updateCollectionPath, "{environment_id}", environmentID, 1)
		updateCollectionPath = strings.Replace(updateCollectionPath, "{collection_id}", collectionID, 1)
		Context(`Successfully - Update a collection`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateCollectionPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call UpdateCollection`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.UpdateCollection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				updateCollectionOptions := testService.NewUpdateCollectionOptions(environmentID, collectionID, name)
				result, response, operationErr = testService.UpdateCollection(updateCollectionOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteCollection(deleteCollectionOptions *DeleteCollectionOptions)`, func() {
		deleteCollectionPath := "/v1/environments/{environment_id}/collections/{collection_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		deleteCollectionPath = strings.Replace(deleteCollectionPath, "{environment_id}", environmentID, 1)
		deleteCollectionPath = strings.Replace(deleteCollectionPath, "{collection_id}", collectionID, 1)
		Context(`Successfully - Delete a collection`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteCollectionPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"collection_id": "fake_CollectionID", "status": "fake_Status"}`)
			}))
			It(`Succeed to call DeleteCollection`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.DeleteCollection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				deleteCollectionOptions := testService.NewDeleteCollectionOptions(environmentID, collectionID)
				result, response, operationErr = testService.DeleteCollection(deleteCollectionOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListCollectionFields(listCollectionFieldsOptions *ListCollectionFieldsOptions)`, func() {
		listCollectionFieldsPath := "/v1/environments/{environment_id}/collections/{collection_id}/fields"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		listCollectionFieldsPath = strings.Replace(listCollectionFieldsPath, "{environment_id}", environmentID, 1)
		listCollectionFieldsPath = strings.Replace(listCollectionFieldsPath, "{collection_id}", collectionID, 1)
		Context(`Successfully - List collection fields`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listCollectionFieldsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call ListCollectionFields`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListCollectionFields(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listCollectionFieldsOptions := testService.NewListCollectionFieldsOptions(environmentID, collectionID)
				result, response, operationErr = testService.ListCollectionFields(listCollectionFieldsOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListExpansions(listExpansionsOptions *ListExpansionsOptions)`, func() {
		listExpansionsPath := "/v1/environments/{environment_id}/collections/{collection_id}/expansions"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		listExpansionsPath = strings.Replace(listExpansionsPath, "{environment_id}", environmentID, 1)
		listExpansionsPath = strings.Replace(listExpansionsPath, "{collection_id}", collectionID, 1)
		Context(`Successfully - Get the expansion list`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listExpansionsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"expansions": []}`)
			}))
			It(`Succeed to call ListExpansions`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListExpansions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listExpansionsOptions := testService.NewListExpansionsOptions(environmentID, collectionID)
				result, response, operationErr = testService.ListExpansions(listExpansionsOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateExpansions(createExpansionsOptions *CreateExpansionsOptions)`, func() {
		createExpansionsPath := "/v1/environments/{environment_id}/collections/{collection_id}/expansions"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		expansions := []discoveryv1.Expansion{}
		createExpansionsPath = strings.Replace(createExpansionsPath, "{environment_id}", environmentID, 1)
		createExpansionsPath = strings.Replace(createExpansionsPath, "{collection_id}", collectionID, 1)
		Context(`Successfully - Create or update expansion list`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createExpansionsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"expansions": []}`)
			}))
			It(`Succeed to call CreateExpansions`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.CreateExpansions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				createExpansionsOptions := testService.NewCreateExpansionsOptions(environmentID, collectionID, expansions)
				result, response, operationErr = testService.CreateExpansions(createExpansionsOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteExpansions(deleteExpansionsOptions *DeleteExpansionsOptions)`, func() {
		deleteExpansionsPath := "/v1/environments/{environment_id}/collections/{collection_id}/expansions"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		deleteExpansionsPath = strings.Replace(deleteExpansionsPath, "{environment_id}", environmentID, 1)
		deleteExpansionsPath = strings.Replace(deleteExpansionsPath, "{collection_id}", collectionID, 1)
		Context(`Successfully - Delete the expansion list`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteExpansionsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(204)
			}))
			It(`Succeed to call DeleteExpansions`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.DeleteExpansions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				deleteExpansionsOptions := testService.NewDeleteExpansionsOptions(environmentID, collectionID)
				response, operationErr = testService.DeleteExpansions(deleteExpansionsOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetTokenizationDictionaryStatus(getTokenizationDictionaryStatusOptions *GetTokenizationDictionaryStatusOptions)`, func() {
		getTokenizationDictionaryStatusPath := "/v1/environments/{environment_id}/collections/{collection_id}/word_lists/tokenization_dictionary"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		getTokenizationDictionaryStatusPath = strings.Replace(getTokenizationDictionaryStatusPath, "{environment_id}", environmentID, 1)
		getTokenizationDictionaryStatusPath = strings.Replace(getTokenizationDictionaryStatusPath, "{collection_id}", collectionID, 1)
		Context(`Successfully - Get tokenization dictionary status`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getTokenizationDictionaryStatusPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call GetTokenizationDictionaryStatus`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetTokenizationDictionaryStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getTokenizationDictionaryStatusOptions := testService.NewGetTokenizationDictionaryStatusOptions(environmentID, collectionID)
				result, response, operationErr = testService.GetTokenizationDictionaryStatus(getTokenizationDictionaryStatusOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateTokenizationDictionary(createTokenizationDictionaryOptions *CreateTokenizationDictionaryOptions)`, func() {
		createTokenizationDictionaryPath := "/v1/environments/{environment_id}/collections/{collection_id}/word_lists/tokenization_dictionary"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		createTokenizationDictionaryPath = strings.Replace(createTokenizationDictionaryPath, "{environment_id}", environmentID, 1)
		createTokenizationDictionaryPath = strings.Replace(createTokenizationDictionaryPath, "{collection_id}", collectionID, 1)
		Context(`Successfully - Create tokenization dictionary`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createTokenizationDictionaryPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(202)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call CreateTokenizationDictionary`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.CreateTokenizationDictionary(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				createTokenizationDictionaryOptions := testService.NewCreateTokenizationDictionaryOptions(environmentID, collectionID)
				result, response, operationErr = testService.CreateTokenizationDictionary(createTokenizationDictionaryOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteTokenizationDictionary(deleteTokenizationDictionaryOptions *DeleteTokenizationDictionaryOptions)`, func() {
		deleteTokenizationDictionaryPath := "/v1/environments/{environment_id}/collections/{collection_id}/word_lists/tokenization_dictionary"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		deleteTokenizationDictionaryPath = strings.Replace(deleteTokenizationDictionaryPath, "{environment_id}", environmentID, 1)
		deleteTokenizationDictionaryPath = strings.Replace(deleteTokenizationDictionaryPath, "{collection_id}", collectionID, 1)
		Context(`Successfully - Delete tokenization dictionary`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteTokenizationDictionaryPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(200)
			}))
			It(`Succeed to call DeleteTokenizationDictionary`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.DeleteTokenizationDictionary(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				deleteTokenizationDictionaryOptions := testService.NewDeleteTokenizationDictionaryOptions(environmentID, collectionID)
				response, operationErr = testService.DeleteTokenizationDictionary(deleteTokenizationDictionaryOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetStopwordListStatus(getStopwordListStatusOptions *GetStopwordListStatusOptions)`, func() {
		getStopwordListStatusPath := "/v1/environments/{environment_id}/collections/{collection_id}/word_lists/stopwords"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		getStopwordListStatusPath = strings.Replace(getStopwordListStatusPath, "{environment_id}", environmentID, 1)
		getStopwordListStatusPath = strings.Replace(getStopwordListStatusPath, "{collection_id}", collectionID, 1)
		Context(`Successfully - Get stopword list status`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getStopwordListStatusPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call GetStopwordListStatus`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetStopwordListStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getStopwordListStatusOptions := testService.NewGetStopwordListStatusOptions(environmentID, collectionID)
				result, response, operationErr = testService.GetStopwordListStatus(getStopwordListStatusOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateStopwordList(createStopwordListOptions *CreateStopwordListOptions)`, func() {
		createStopwordListPath := "/v1/environments/{environment_id}/collections/{collection_id}/word_lists/stopwords"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		pwd, _ := os.Getwd()
		stopwordFile, stopwordsFileErr := os.Open(pwd + "/../resources/stopwords.txt")
		if stopwordsFileErr != nil {
			panic(stopwordsFileErr)
		}
		stopwordFilename := "exampleString"
		createStopwordListPath = strings.Replace(createStopwordListPath, "{environment_id}", environmentID, 1)
		createStopwordListPath = strings.Replace(createStopwordListPath, "{collection_id}", collectionID, 1)
		Context(`Successfully - Create stopword list`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createStopwordListPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call CreateStopwordList`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.CreateStopwordList(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				createStopwordListOptions := testService.NewCreateStopwordListOptions(environmentID, collectionID, stopwordFile, stopwordFilename)
				result, response, operationErr = testService.CreateStopwordList(createStopwordListOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteStopwordList(deleteStopwordListOptions *DeleteStopwordListOptions)`, func() {
		deleteStopwordListPath := "/v1/environments/{environment_id}/collections/{collection_id}/word_lists/stopwords"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		deleteStopwordListPath = strings.Replace(deleteStopwordListPath, "{environment_id}", environmentID, 1)
		deleteStopwordListPath = strings.Replace(deleteStopwordListPath, "{collection_id}", collectionID, 1)
		Context(`Successfully - Delete a custom stopword list`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteStopwordListPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(200)
			}))
			It(`Succeed to call DeleteStopwordList`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.DeleteStopwordList(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				deleteStopwordListOptions := testService.NewDeleteStopwordListOptions(environmentID, collectionID)
				response, operationErr = testService.DeleteStopwordList(deleteStopwordListOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`AddDocument(addDocumentOptions *AddDocumentOptions)`, func() {
		addDocumentPath := "/v1/environments/{environment_id}/collections/{collection_id}/documents"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		addDocumentPath = strings.Replace(addDocumentPath, "{environment_id}", environmentID, 1)
		addDocumentPath = strings.Replace(addDocumentPath, "{collection_id}", collectionID, 1)
		Context(`Successfully - Add a document`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(addDocumentPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(202)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call AddDocument`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.AddDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				addDocumentOptions := testService.NewAddDocumentOptions(environmentID, collectionID).
					SetMetadata("Name:John Smith")
				result, response, operationErr = testService.AddDocument(addDocumentOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetDocumentStatus(getDocumentStatusOptions *GetDocumentStatusOptions)`, func() {
		getDocumentStatusPath := "/v1/environments/{environment_id}/collections/{collection_id}/documents/{document_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		documentID := "exampleString"
		getDocumentStatusPath = strings.Replace(getDocumentStatusPath, "{environment_id}", environmentID, 1)
		getDocumentStatusPath = strings.Replace(getDocumentStatusPath, "{collection_id}", collectionID, 1)
		getDocumentStatusPath = strings.Replace(getDocumentStatusPath, "{document_id}", documentID, 1)
		Context(`Successfully - Get document details`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getDocumentStatusPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"document_id": "fake_DocumentID", "status": "fake_Status", "status_description": "fake_StatusDescription", "notices": []}`)
			}))
			It(`Succeed to call GetDocumentStatus`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetDocumentStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getDocumentStatusOptions := testService.NewGetDocumentStatusOptions(environmentID, collectionID, documentID)
				result, response, operationErr = testService.GetDocumentStatus(getDocumentStatusOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateDocument(updateDocumentOptions *UpdateDocumentOptions)`, func() {
		updateDocumentPath := "/v1/environments/{environment_id}/collections/{collection_id}/documents/{document_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		documentID := "exampleString"
		updateDocumentPath = strings.Replace(updateDocumentPath, "{environment_id}", environmentID, 1)
		updateDocumentPath = strings.Replace(updateDocumentPath, "{collection_id}", collectionID, 1)
		updateDocumentPath = strings.Replace(updateDocumentPath, "{document_id}", documentID, 1)
		Context(`Successfully - Update a document`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateDocumentPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(202)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call UpdateDocument`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.UpdateDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				updateDocumentOptions := testService.NewUpdateDocumentOptions(environmentID, collectionID, documentID).
					SetMetadata("Name:John Smith")
				result, response, operationErr = testService.UpdateDocument(updateDocumentOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteDocument(deleteDocumentOptions *DeleteDocumentOptions)`, func() {
		deleteDocumentPath := "/v1/environments/{environment_id}/collections/{collection_id}/documents/{document_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		documentID := "exampleString"
		deleteDocumentPath = strings.Replace(deleteDocumentPath, "{environment_id}", environmentID, 1)
		deleteDocumentPath = strings.Replace(deleteDocumentPath, "{collection_id}", collectionID, 1)
		deleteDocumentPath = strings.Replace(deleteDocumentPath, "{document_id}", documentID, 1)
		Context(`Successfully - Delete a document`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteDocumentPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call DeleteDocument`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.DeleteDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				deleteDocumentOptions := testService.NewDeleteDocumentOptions(environmentID, collectionID, documentID)
				result, response, operationErr = testService.DeleteDocument(deleteDocumentOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`Query(queryOptions *QueryOptions)`, func() {
		queryPath := "/v1/environments/{environment_id}/collections/{collection_id}/query"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		queryPath = strings.Replace(queryPath, "{environment_id}", environmentID, 1)
		queryPath = strings.Replace(queryPath, "{collection_id}", collectionID, 1)
		Context(`Successfully - Query a collection`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(queryPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call Query`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.Query(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				queryOptions := testService.NewQueryOptions(environmentID, collectionID)
				result, response, operationErr = testService.Query(queryOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`QueryNotices(queryNoticesOptions *QueryNoticesOptions)`, func() {
		queryNoticesPath := "/v1/environments/{environment_id}/collections/{collection_id}/notices"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		queryNoticesPath = strings.Replace(queryNoticesPath, "{environment_id}", environmentID, 1)
		queryNoticesPath = strings.Replace(queryNoticesPath, "{collection_id}", collectionID, 1)
		Context(`Successfully - Query system notices`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(queryNoticesPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call QueryNotices`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.QueryNotices(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				queryNoticesOptions := testService.NewQueryNoticesOptions(environmentID, collectionID)
				result, response, operationErr = testService.QueryNotices(queryNoticesOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`FederatedQuery(federatedQueryOptions *FederatedQueryOptions)`, func() {
		federatedQueryPath := "/v1/environments/{environment_id}/query"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		collectionIds := "exampleString"
		federatedQueryPath = strings.Replace(federatedQueryPath, "{environment_id}", environmentID, 1)
		Context(`Successfully - Query multiple collections`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(federatedQueryPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call FederatedQuery`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.FederatedQuery(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				federatedQueryOptions := testService.NewFederatedQueryOptions(environmentID, collectionIds)
				result, response, operationErr = testService.FederatedQuery(federatedQueryOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`FederatedQueryNotices(federatedQueryNoticesOptions *FederatedQueryNoticesOptions)`, func() {
		federatedQueryNoticesPath := "/v1/environments/{environment_id}/notices"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		collectionIds := []string{}
		federatedQueryNoticesPath = strings.Replace(federatedQueryNoticesPath, "{environment_id}", environmentID, 1)
		Context(`Successfully - Query multiple collection system notices`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(federatedQueryNoticesPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call FederatedQueryNotices`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.FederatedQueryNotices(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				federatedQueryNoticesOptions := testService.NewFederatedQueryNoticesOptions(environmentID, collectionIds)
				result, response, operationErr = testService.FederatedQueryNotices(federatedQueryNoticesOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetAutocompletion(getAutocompletionOptions *GetAutocompletionOptions)`, func() {
		getAutocompletionPath := "/v1/environments/{environment_id}/collections/{collection_id}/autocompletion"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		prefix := "exampleString"
		getAutocompletionPath = strings.Replace(getAutocompletionPath, "{environment_id}", environmentID, 1)
		getAutocompletionPath = strings.Replace(getAutocompletionPath, "{collection_id}", collectionID, 1)
		Context(`Successfully - Get Autocomplete Suggestions`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getAutocompletionPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				Expect(req.URL.Query()["prefix"]).To(Equal([]string{prefix}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call GetAutocompletion`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetAutocompletion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getAutocompletionOptions := testService.NewGetAutocompletionOptions(environmentID, collectionID, prefix)
				result, response, operationErr = testService.GetAutocompletion(getAutocompletionOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListTrainingData(listTrainingDataOptions *ListTrainingDataOptions)`, func() {
		listTrainingDataPath := "/v1/environments/{environment_id}/collections/{collection_id}/training_data"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		listTrainingDataPath = strings.Replace(listTrainingDataPath, "{environment_id}", environmentID, 1)
		listTrainingDataPath = strings.Replace(listTrainingDataPath, "{collection_id}", collectionID, 1)
		Context(`Successfully - List training data`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listTrainingDataPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call ListTrainingData`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListTrainingData(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listTrainingDataOptions := testService.NewListTrainingDataOptions(environmentID, collectionID)
				result, response, operationErr = testService.ListTrainingData(listTrainingDataOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`AddTrainingData(addTrainingDataOptions *AddTrainingDataOptions)`, func() {
		addTrainingDataPath := "/v1/environments/{environment_id}/collections/{collection_id}/training_data"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		addTrainingDataPath = strings.Replace(addTrainingDataPath, "{environment_id}", environmentID, 1)
		addTrainingDataPath = strings.Replace(addTrainingDataPath, "{collection_id}", collectionID, 1)
		Context(`Successfully - Add query to training data`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(addTrainingDataPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call AddTrainingData`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.AddTrainingData(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				addTrainingDataOptions := testService.NewAddTrainingDataOptions(environmentID, collectionID)
				result, response, operationErr = testService.AddTrainingData(addTrainingDataOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteAllTrainingData(deleteAllTrainingDataOptions *DeleteAllTrainingDataOptions)`, func() {
		deleteAllTrainingDataPath := "/v1/environments/{environment_id}/collections/{collection_id}/training_data"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		deleteAllTrainingDataPath = strings.Replace(deleteAllTrainingDataPath, "{environment_id}", environmentID, 1)
		deleteAllTrainingDataPath = strings.Replace(deleteAllTrainingDataPath, "{collection_id}", collectionID, 1)
		Context(`Successfully - Delete all training data`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteAllTrainingDataPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(204)
			}))
			It(`Succeed to call DeleteAllTrainingData`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.DeleteAllTrainingData(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				deleteAllTrainingDataOptions := testService.NewDeleteAllTrainingDataOptions(environmentID, collectionID)
				response, operationErr = testService.DeleteAllTrainingData(deleteAllTrainingDataOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetTrainingData(getTrainingDataOptions *GetTrainingDataOptions)`, func() {
		getTrainingDataPath := "/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		queryID := "exampleString"
		getTrainingDataPath = strings.Replace(getTrainingDataPath, "{environment_id}", environmentID, 1)
		getTrainingDataPath = strings.Replace(getTrainingDataPath, "{collection_id}", collectionID, 1)
		getTrainingDataPath = strings.Replace(getTrainingDataPath, "{query_id}", queryID, 1)
		Context(`Successfully - Get details about a query`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getTrainingDataPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call GetTrainingData`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetTrainingData(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getTrainingDataOptions := testService.NewGetTrainingDataOptions(environmentID, collectionID, queryID)
				result, response, operationErr = testService.GetTrainingData(getTrainingDataOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteTrainingData(deleteTrainingDataOptions *DeleteTrainingDataOptions)`, func() {
		deleteTrainingDataPath := "/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		queryID := "exampleString"
		deleteTrainingDataPath = strings.Replace(deleteTrainingDataPath, "{environment_id}", environmentID, 1)
		deleteTrainingDataPath = strings.Replace(deleteTrainingDataPath, "{collection_id}", collectionID, 1)
		deleteTrainingDataPath = strings.Replace(deleteTrainingDataPath, "{query_id}", queryID, 1)
		Context(`Successfully - Delete a training data query`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteTrainingDataPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(204)
			}))
			It(`Succeed to call DeleteTrainingData`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.DeleteTrainingData(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				deleteTrainingDataOptions := testService.NewDeleteTrainingDataOptions(environmentID, collectionID, queryID)
				response, operationErr = testService.DeleteTrainingData(deleteTrainingDataOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`ListTrainingExamples(listTrainingExamplesOptions *ListTrainingExamplesOptions)`, func() {
		listTrainingExamplesPath := "/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}/examples"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		queryID := "exampleString"
		listTrainingExamplesPath = strings.Replace(listTrainingExamplesPath, "{environment_id}", environmentID, 1)
		listTrainingExamplesPath = strings.Replace(listTrainingExamplesPath, "{collection_id}", collectionID, 1)
		listTrainingExamplesPath = strings.Replace(listTrainingExamplesPath, "{query_id}", queryID, 1)
		Context(`Successfully - List examples for a training data query`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listTrainingExamplesPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call ListTrainingExamples`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListTrainingExamples(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listTrainingExamplesOptions := testService.NewListTrainingExamplesOptions(environmentID, collectionID, queryID)
				result, response, operationErr = testService.ListTrainingExamples(listTrainingExamplesOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateTrainingExample(createTrainingExampleOptions *CreateTrainingExampleOptions)`, func() {
		createTrainingExamplePath := "/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}/examples"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		queryID := "exampleString"
		createTrainingExamplePath = strings.Replace(createTrainingExamplePath, "{environment_id}", environmentID, 1)
		createTrainingExamplePath = strings.Replace(createTrainingExamplePath, "{collection_id}", collectionID, 1)
		createTrainingExamplePath = strings.Replace(createTrainingExamplePath, "{query_id}", queryID, 1)
		Context(`Successfully - Add example to training data query`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createTrainingExamplePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call CreateTrainingExample`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.CreateTrainingExample(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				createTrainingExampleOptions := testService.NewCreateTrainingExampleOptions(environmentID, collectionID, queryID)
				result, response, operationErr = testService.CreateTrainingExample(createTrainingExampleOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteTrainingExample(deleteTrainingExampleOptions *DeleteTrainingExampleOptions)`, func() {
		deleteTrainingExamplePath := "/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}/examples/{example_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		queryID := "exampleString"
		exampleID := "exampleString"
		deleteTrainingExamplePath = strings.Replace(deleteTrainingExamplePath, "{environment_id}", environmentID, 1)
		deleteTrainingExamplePath = strings.Replace(deleteTrainingExamplePath, "{collection_id}", collectionID, 1)
		deleteTrainingExamplePath = strings.Replace(deleteTrainingExamplePath, "{query_id}", queryID, 1)
		deleteTrainingExamplePath = strings.Replace(deleteTrainingExamplePath, "{example_id}", exampleID, 1)
		Context(`Successfully - Delete example for training data query`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteTrainingExamplePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(204)
			}))
			It(`Succeed to call DeleteTrainingExample`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.DeleteTrainingExample(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				deleteTrainingExampleOptions := testService.NewDeleteTrainingExampleOptions(environmentID, collectionID, queryID, exampleID)
				response, operationErr = testService.DeleteTrainingExample(deleteTrainingExampleOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateTrainingExample(updateTrainingExampleOptions *UpdateTrainingExampleOptions)`, func() {
		updateTrainingExamplePath := "/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}/examples/{example_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		queryID := "exampleString"
		exampleID := "exampleString"
		updateTrainingExamplePath = strings.Replace(updateTrainingExamplePath, "{environment_id}", environmentID, 1)
		updateTrainingExamplePath = strings.Replace(updateTrainingExamplePath, "{collection_id}", collectionID, 1)
		updateTrainingExamplePath = strings.Replace(updateTrainingExamplePath, "{query_id}", queryID, 1)
		updateTrainingExamplePath = strings.Replace(updateTrainingExamplePath, "{example_id}", exampleID, 1)
		Context(`Successfully - Change label or cross reference for example`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateTrainingExamplePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call UpdateTrainingExample`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.UpdateTrainingExample(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				updateTrainingExampleOptions := testService.NewUpdateTrainingExampleOptions(environmentID, collectionID, queryID, exampleID)
				result, response, operationErr = testService.UpdateTrainingExample(updateTrainingExampleOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetTrainingExample(getTrainingExampleOptions *GetTrainingExampleOptions)`, func() {
		getTrainingExamplePath := "/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}/examples/{example_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		queryID := "exampleString"
		exampleID := "exampleString"
		getTrainingExamplePath = strings.Replace(getTrainingExamplePath, "{environment_id}", environmentID, 1)
		getTrainingExamplePath = strings.Replace(getTrainingExamplePath, "{collection_id}", collectionID, 1)
		getTrainingExamplePath = strings.Replace(getTrainingExamplePath, "{query_id}", queryID, 1)
		getTrainingExamplePath = strings.Replace(getTrainingExamplePath, "{example_id}", exampleID, 1)
		Context(`Successfully - Get details for training data example`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getTrainingExamplePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call GetTrainingExample`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetTrainingExample(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getTrainingExampleOptions := testService.NewGetTrainingExampleOptions(environmentID, collectionID, queryID, exampleID)
				result, response, operationErr = testService.GetTrainingExample(getTrainingExampleOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteUserData(deleteUserDataOptions *DeleteUserDataOptions)`, func() {
		deleteUserDataPath := "/v1/user_data"
		version := "exampleString"
		bearerToken := "0ui9876453"
		customerID := "exampleString"
		Context(`Successfully - Delete labeled data`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteUserDataPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				Expect(req.URL.Query()["customer_id"]).To(Equal([]string{customerID}))

				res.WriteHeader(200)
			}))
			It(`Succeed to call DeleteUserData`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.DeleteUserData(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				deleteUserDataOptions := testService.NewDeleteUserDataOptions(customerID)
				response, operationErr = testService.DeleteUserData(deleteUserDataOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateEvent(createEventOptions *CreateEventOptions)`, func() {
		createEventPath := "/v1/events"
		version := "exampleString"
		bearerToken := "0ui9876453"
		typeVar := "exampleString"
		data := &discoveryv1.EventData{EnvironmentID: core.StringPtr("exampleString"), SessionToken: core.StringPtr("exampleString"), CollectionID: core.StringPtr("exampleString"), DocumentID: core.StringPtr("exampleString")}
		Context(`Successfully - Create event`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createEventPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call CreateEvent`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.CreateEvent(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				createEventOptions := testService.NewCreateEventOptions(typeVar, data)
				result, response, operationErr = testService.CreateEvent(createEventOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`QueryLog(queryLogOptions *QueryLogOptions)`, func() {
		queryLogPath := "/v1/logs"
		version := "exampleString"
		bearerToken := "0ui9876453"
		Context(`Successfully - Search the query and event log`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(queryLogPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call QueryLog`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.QueryLog(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				queryLogOptions := testService.NewQueryLogOptions()
				result, response, operationErr = testService.QueryLog(queryLogOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetMetricsQuery(getMetricsQueryOptions *GetMetricsQueryOptions)`, func() {
		getMetricsQueryPath := "/v1/metrics/number_of_queries"
		version := "exampleString"
		bearerToken := "0ui9876453"
		Context(`Successfully - Number of queries over time`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getMetricsQueryPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call GetMetricsQuery`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetMetricsQuery(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getMetricsQueryOptions := testService.NewGetMetricsQueryOptions()
				result, response, operationErr = testService.GetMetricsQuery(getMetricsQueryOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetMetricsQueryEvent(getMetricsQueryEventOptions *GetMetricsQueryEventOptions)`, func() {
		getMetricsQueryEventPath := "/v1/metrics/number_of_queries_with_event"
		version := "exampleString"
		bearerToken := "0ui9876453"
		Context(`Successfully - Number of queries with an event over time`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getMetricsQueryEventPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call GetMetricsQueryEvent`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetMetricsQueryEvent(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getMetricsQueryEventOptions := testService.NewGetMetricsQueryEventOptions()
				result, response, operationErr = testService.GetMetricsQueryEvent(getMetricsQueryEventOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetMetricsQueryNoResults(getMetricsQueryNoResultsOptions *GetMetricsQueryNoResultsOptions)`, func() {
		getMetricsQueryNoResultsPath := "/v1/metrics/number_of_queries_with_no_search_results"
		version := "exampleString"
		bearerToken := "0ui9876453"
		Context(`Successfully - Number of queries with no search results over time`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getMetricsQueryNoResultsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call GetMetricsQueryNoResults`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetMetricsQueryNoResults(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getMetricsQueryNoResultsOptions := testService.NewGetMetricsQueryNoResultsOptions()
				result, response, operationErr = testService.GetMetricsQueryNoResults(getMetricsQueryNoResultsOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetMetricsEventRate(getMetricsEventRateOptions *GetMetricsEventRateOptions)`, func() {
		getMetricsEventRatePath := "/v1/metrics/event_rate"
		version := "exampleString"
		bearerToken := "0ui9876453"
		Context(`Successfully - Percentage of queries with an associated event`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getMetricsEventRatePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call GetMetricsEventRate`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetMetricsEventRate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getMetricsEventRateOptions := testService.NewGetMetricsEventRateOptions()
				result, response, operationErr = testService.GetMetricsEventRate(getMetricsEventRateOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetMetricsQueryTokenEvent(getMetricsQueryTokenEventOptions *GetMetricsQueryTokenEventOptions)`, func() {
		getMetricsQueryTokenEventPath := "/v1/metrics/top_query_tokens_with_event_rate"
		version := "exampleString"
		bearerToken := "0ui9876453"
		Context(`Successfully - Most frequent query tokens with an event`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getMetricsQueryTokenEventPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call GetMetricsQueryTokenEvent`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetMetricsQueryTokenEvent(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getMetricsQueryTokenEventOptions := testService.NewGetMetricsQueryTokenEventOptions()
				result, response, operationErr = testService.GetMetricsQueryTokenEvent(getMetricsQueryTokenEventOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListCredentials(listCredentialsOptions *ListCredentialsOptions)`, func() {
		listCredentialsPath := "/v1/environments/{environment_id}/credentials"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		listCredentialsPath = strings.Replace(listCredentialsPath, "{environment_id}", environmentID, 1)
		Context(`Successfully - List credentials`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listCredentialsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call ListCredentials`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListCredentials(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listCredentialsOptions := testService.NewListCredentialsOptions(environmentID)
				result, response, operationErr = testService.ListCredentials(listCredentialsOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateCredentials(createCredentialsOptions *CreateCredentialsOptions)`, func() {
		createCredentialsPath := "/v1/environments/{environment_id}/credentials"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		createCredentialsPath = strings.Replace(createCredentialsPath, "{environment_id}", environmentID, 1)
		Context(`Successfully - Create credentials`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createCredentialsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call CreateCredentials`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.CreateCredentials(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				createCredentialsOptions := testService.NewCreateCredentialsOptions(environmentID)
				result, response, operationErr = testService.CreateCredentials(createCredentialsOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetCredentials(getCredentialsOptions *GetCredentialsOptions)`, func() {
		getCredentialsPath := "/v1/environments/{environment_id}/credentials/{credential_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		credentialID := "exampleString"
		getCredentialsPath = strings.Replace(getCredentialsPath, "{environment_id}", environmentID, 1)
		getCredentialsPath = strings.Replace(getCredentialsPath, "{credential_id}", credentialID, 1)
		Context(`Successfully - View Credentials`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getCredentialsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call GetCredentials`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetCredentials(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getCredentialsOptions := testService.NewGetCredentialsOptions(environmentID, credentialID)
				result, response, operationErr = testService.GetCredentials(getCredentialsOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateCredentials(updateCredentialsOptions *UpdateCredentialsOptions)`, func() {
		updateCredentialsPath := "/v1/environments/{environment_id}/credentials/{credential_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		credentialID := "exampleString"
		updateCredentialsPath = strings.Replace(updateCredentialsPath, "{environment_id}", environmentID, 1)
		updateCredentialsPath = strings.Replace(updateCredentialsPath, "{credential_id}", credentialID, 1)
		Context(`Successfully - Update credentials`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateCredentialsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call UpdateCredentials`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.UpdateCredentials(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				updateCredentialsOptions := testService.NewUpdateCredentialsOptions(environmentID, credentialID)
				result, response, operationErr = testService.UpdateCredentials(updateCredentialsOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteCredentials(deleteCredentialsOptions *DeleteCredentialsOptions)`, func() {
		deleteCredentialsPath := "/v1/environments/{environment_id}/credentials/{credential_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		credentialID := "exampleString"
		deleteCredentialsPath = strings.Replace(deleteCredentialsPath, "{environment_id}", environmentID, 1)
		deleteCredentialsPath = strings.Replace(deleteCredentialsPath, "{credential_id}", credentialID, 1)
		Context(`Successfully - Delete credentials`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteCredentialsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call DeleteCredentials`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.DeleteCredentials(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				deleteCredentialsOptions := testService.NewDeleteCredentialsOptions(environmentID, credentialID)
				result, response, operationErr = testService.DeleteCredentials(deleteCredentialsOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListGateways(listGatewaysOptions *ListGatewaysOptions)`, func() {
		listGatewaysPath := "/v1/environments/{environment_id}/gateways"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		listGatewaysPath = strings.Replace(listGatewaysPath, "{environment_id}", environmentID, 1)
		Context(`Successfully - List Gateways`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listGatewaysPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call ListGateways`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListGateways(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listGatewaysOptions := testService.NewListGatewaysOptions(environmentID)
				result, response, operationErr = testService.ListGateways(listGatewaysOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateGateway(createGatewayOptions *CreateGatewayOptions)`, func() {
		createGatewayPath := "/v1/environments/{environment_id}/gateways"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		createGatewayPath = strings.Replace(createGatewayPath, "{environment_id}", environmentID, 1)
		Context(`Successfully - Create Gateway`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createGatewayPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call CreateGateway`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.CreateGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				createGatewayOptions := testService.NewCreateGatewayOptions(environmentID)
				result, response, operationErr = testService.CreateGateway(createGatewayOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetGateway(getGatewayOptions *GetGatewayOptions)`, func() {
		getGatewayPath := "/v1/environments/{environment_id}/gateways/{gateway_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		gatewayID := "exampleString"
		getGatewayPath = strings.Replace(getGatewayPath, "{environment_id}", environmentID, 1)
		getGatewayPath = strings.Replace(getGatewayPath, "{gateway_id}", gatewayID, 1)
		Context(`Successfully - List Gateway Details`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getGatewayPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call GetGateway`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getGatewayOptions := testService.NewGetGatewayOptions(environmentID, gatewayID)
				result, response, operationErr = testService.GetGateway(getGatewayOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteGateway(deleteGatewayOptions *DeleteGatewayOptions)`, func() {
		deleteGatewayPath := "/v1/environments/{environment_id}/gateways/{gateway_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		environmentID := "exampleString"
		gatewayID := "exampleString"
		deleteGatewayPath = strings.Replace(deleteGatewayPath, "{environment_id}", environmentID, 1)
		deleteGatewayPath = strings.Replace(deleteGatewayPath, "{gateway_id}", gatewayID, 1)
		Context(`Successfully - Delete Gateway`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteGatewayPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call DeleteGateway`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.DeleteGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				deleteGatewayOptions := testService.NewDeleteGatewayOptions(environmentID, gatewayID)
				result, response, operationErr = testService.DeleteGateway(deleteGatewayOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("Model constructor tests", func() {
		Context("with a sample service", func() {
			version := "1970-01-01"
			testService, _ := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
				URL:           "http://discoveryv1modelgenerator.com",
				Version:       version,
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It("should call NewConfiguration successfully", func() {
				name := "exampleString"
				model, err := testService.NewConfiguration(name)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It("should call NewEnrichment successfully", func() {
				destinationField := "exampleString"
				sourceField := "exampleString"
				enrichment := "exampleString"
				model, err := testService.NewEnrichment(destinationField, sourceField, enrichment)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It("should call NewEventData successfully", func() {
				environmentID := "exampleString"
				sessionToken := "exampleString"
				collectionID := "exampleString"
				documentID := "exampleString"
				model, err := testService.NewEventData(environmentID, sessionToken, collectionID, documentID)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It("should call NewExpansion successfully", func() {
				expandedTerms := []string{}
				model, err := testService.NewExpansion(expandedTerms)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It("should call NewExpansions successfully", func() {
				expansions := []discoveryv1.Expansion{}
				model, err := testService.NewExpansions(expansions)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It("should call NewSourceOptionsBuckets successfully", func() {
				name := "exampleString"
				model, err := testService.NewSourceOptionsBuckets(name)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It("should call NewSourceOptionsFolder successfully", func() {
				ownerUserID := "exampleString"
				folderID := "exampleString"
				model, err := testService.NewSourceOptionsFolder(ownerUserID, folderID)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It("should call NewSourceOptionsObject successfully", func() {
				name := "exampleString"
				model, err := testService.NewSourceOptionsObject(name)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It("should call NewSourceOptionsSiteColl successfully", func() {
				siteCollectionPath := "exampleString"
				model, err := testService.NewSourceOptionsSiteColl(siteCollectionPath)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It("should call NewSourceOptionsWebCrawl successfully", func() {
				url := "exampleString"
				model, err := testService.NewSourceOptionsWebCrawl(url)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It("should call NewTokenDictRule successfully", func() {
				text := "exampleString"
				tokens := []string{}
				partOfSpeech := "exampleString"
				model, err := testService.NewTokenDictRule(text, tokens, partOfSpeech)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
})
