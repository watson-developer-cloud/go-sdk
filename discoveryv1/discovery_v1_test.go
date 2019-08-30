/**
 * (C) Copyright IBM Corp. 2019.
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
	"github.com/watson-developer-cloud/go-sdk/discoveryv1"
    "github.com/IBM/go-sdk-core/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"strings"
	"os"
)

var _ = Describe("DiscoveryV1", func() {
	Describe("CreateEnvironment(createEnvironmentOptions *CreateEnvironmentOptions)", func() {
		createEnvironmentPath := "/v1/environments"
		version := "exampleString"
		accessToken := "0ui9876453"
		name := "exampleString"
		Context("Successfully - Create an environment", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createEnvironmentPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(201)
			}))
			It("Succeed to call CreateEnvironment", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.CreateEnvironment(nil)
				Expect(returnValueErr).NotTo(BeNil())

				createEnvironmentOptions := testService.NewCreateEnvironmentOptions(name)
				returnValue, returnValueErr = testService.CreateEnvironment(createEnvironmentOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetCreateEnvironmentResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListEnvironments(listEnvironmentsOptions *ListEnvironmentsOptions)", func() {
		listEnvironmentsPath := "/v1/environments"
		version := "exampleString"
		accessToken := "0ui9876453"
		Context("Successfully - List environments", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listEnvironmentsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call ListEnvironments", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ListEnvironments(nil)
				Expect(returnValueErr).NotTo(BeNil())

				listEnvironmentsOptions := testService.NewListEnvironmentsOptions()
				returnValue, returnValueErr = testService.ListEnvironments(listEnvironmentsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListEnvironmentsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetEnvironment(getEnvironmentOptions *GetEnvironmentOptions)", func() {
		getEnvironmentPath := "/v1/environments/{environment_id}"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		getEnvironmentPath = strings.Replace(getEnvironmentPath, "{environment_id}", environmentID, 1)
		Context("Successfully - Get environment info", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getEnvironmentPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call GetEnvironment", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.GetEnvironment(nil)
				Expect(returnValueErr).NotTo(BeNil())

				getEnvironmentOptions := testService.NewGetEnvironmentOptions(environmentID)
				returnValue, returnValueErr = testService.GetEnvironment(getEnvironmentOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetEnvironmentResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("UpdateEnvironment(updateEnvironmentOptions *UpdateEnvironmentOptions)", func() {
		updateEnvironmentPath := "/v1/environments/{environment_id}"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		updateEnvironmentPath = strings.Replace(updateEnvironmentPath, "{environment_id}", environmentID, 1)
		Context("Successfully - Update an environment", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateEnvironmentPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call UpdateEnvironment", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.UpdateEnvironment(nil)
				Expect(returnValueErr).NotTo(BeNil())

				updateEnvironmentOptions := testService.NewUpdateEnvironmentOptions(environmentID)
				returnValue, returnValueErr = testService.UpdateEnvironment(updateEnvironmentOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetUpdateEnvironmentResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteEnvironment(deleteEnvironmentOptions *DeleteEnvironmentOptions)", func() {
		deleteEnvironmentPath := "/v1/environments/{environment_id}"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		deleteEnvironmentPath = strings.Replace(deleteEnvironmentPath, "{environment_id}", environmentID, 1)
		Context("Successfully - Delete environment", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteEnvironmentPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"environment_id": "fake EnvironmentID", "status": "fake Status"}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call DeleteEnvironment", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.DeleteEnvironment(nil)
				Expect(returnValueErr).NotTo(BeNil())

				deleteEnvironmentOptions := testService.NewDeleteEnvironmentOptions(environmentID)
				returnValue, returnValueErr = testService.DeleteEnvironment(deleteEnvironmentOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetDeleteEnvironmentResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListFields(listFieldsOptions *ListFieldsOptions)", func() {
		listFieldsPath := "/v1/environments/{environment_id}/fields"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		collectionIds := []string{}
		listFieldsPath = strings.Replace(listFieldsPath, "{environment_id}", environmentID, 1)
		Context("Successfully - List fields across collections", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listFieldsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call ListFields", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ListFields(nil)
				Expect(returnValueErr).NotTo(BeNil())

				listFieldsOptions := testService.NewListFieldsOptions(environmentID, collectionIds)
				returnValue, returnValueErr = testService.ListFields(listFieldsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListFieldsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateConfiguration(createConfigurationOptions *CreateConfigurationOptions)", func() {
		createConfigurationPath := "/v1/environments/{environment_id}/configurations"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		name := "exampleString"
		createConfigurationPath = strings.Replace(createConfigurationPath, "{environment_id}", environmentID, 1)
		Context("Successfully - Add configuration", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createConfigurationPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"name": "fake Name"}`)
				res.WriteHeader(201)
			}))
			It("Succeed to call CreateConfiguration", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.CreateConfiguration(nil)
				Expect(returnValueErr).NotTo(BeNil())

				createConfigurationOptions := testService.NewCreateConfigurationOptions(environmentID, name)
				returnValue, returnValueErr = testService.CreateConfiguration(createConfigurationOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetCreateConfigurationResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListConfigurations(listConfigurationsOptions *ListConfigurationsOptions)", func() {
		listConfigurationsPath := "/v1/environments/{environment_id}/configurations"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		listConfigurationsPath = strings.Replace(listConfigurationsPath, "{environment_id}", environmentID, 1)
		Context("Successfully - List configurations", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listConfigurationsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call ListConfigurations", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ListConfigurations(nil)
				Expect(returnValueErr).NotTo(BeNil())

				listConfigurationsOptions := testService.NewListConfigurationsOptions(environmentID)
				returnValue, returnValueErr = testService.ListConfigurations(listConfigurationsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListConfigurationsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetConfiguration(getConfigurationOptions *GetConfigurationOptions)", func() {
		getConfigurationPath := "/v1/environments/{environment_id}/configurations/{configuration_id}"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		configurationID := "exampleString"
		getConfigurationPath = strings.Replace(getConfigurationPath, "{environment_id}", environmentID, 1)
		getConfigurationPath = strings.Replace(getConfigurationPath, "{configuration_id}", configurationID, 1)
		Context("Successfully - Get configuration details", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getConfigurationPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"name": "fake Name"}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call GetConfiguration", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.GetConfiguration(nil)
				Expect(returnValueErr).NotTo(BeNil())

				getConfigurationOptions := testService.NewGetConfigurationOptions(environmentID, configurationID)
				returnValue, returnValueErr = testService.GetConfiguration(getConfigurationOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetConfigurationResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("UpdateConfiguration(updateConfigurationOptions *UpdateConfigurationOptions)", func() {
		updateConfigurationPath := "/v1/environments/{environment_id}/configurations/{configuration_id}"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		configurationID := "exampleString"
		name := "exampleString"
		updateConfigurationPath = strings.Replace(updateConfigurationPath, "{environment_id}", environmentID, 1)
		updateConfigurationPath = strings.Replace(updateConfigurationPath, "{configuration_id}", configurationID, 1)
		Context("Successfully - Update a configuration", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateConfigurationPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"name": "fake Name"}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call UpdateConfiguration", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.UpdateConfiguration(nil)
				Expect(returnValueErr).NotTo(BeNil())

				updateConfigurationOptions := testService.NewUpdateConfigurationOptions(environmentID, configurationID, name)
				returnValue, returnValueErr = testService.UpdateConfiguration(updateConfigurationOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetUpdateConfigurationResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteConfiguration(deleteConfigurationOptions *DeleteConfigurationOptions)", func() {
		deleteConfigurationPath := "/v1/environments/{environment_id}/configurations/{configuration_id}"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		configurationID := "exampleString"
		deleteConfigurationPath = strings.Replace(deleteConfigurationPath, "{environment_id}", environmentID, 1)
		deleteConfigurationPath = strings.Replace(deleteConfigurationPath, "{configuration_id}", configurationID, 1)
		Context("Successfully - Delete a configuration", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteConfigurationPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"configuration_id": "fake ConfigurationID", "status": "fake Status"}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call DeleteConfiguration", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.DeleteConfiguration(nil)
				Expect(returnValueErr).NotTo(BeNil())

				deleteConfigurationOptions := testService.NewDeleteConfigurationOptions(environmentID, configurationID)
				returnValue, returnValueErr = testService.DeleteConfiguration(deleteConfigurationOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetDeleteConfigurationResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("TestConfigurationInEnvironment(testConfigurationInEnvironmentOptions *TestConfigurationInEnvironmentOptions)", func() {
		testConfigurationInEnvironmentPath := "/v1/environments/{environment_id}/preview"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		testConfigurationInEnvironmentPath = strings.Replace(testConfigurationInEnvironmentPath, "{environment_id}", environmentID, 1)
		Context("Successfully - Test configuration", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(testConfigurationInEnvironmentPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call TestConfigurationInEnvironment", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.TestConfigurationInEnvironment(nil)
				Expect(returnValueErr).NotTo(BeNil())

				testConfigurationInEnvironmentOptions := testService.NewTestConfigurationInEnvironmentOptions(environmentID)
				testConfigurationInEnvironmentOptions.SetConfiguration("ConfigurationID")
				returnValue, returnValueErr = testService.TestConfigurationInEnvironment(testConfigurationInEnvironmentOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetTestConfigurationInEnvironmentResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateCollection(createCollectionOptions *CreateCollectionOptions)", func() {
		createCollectionPath := "/v1/environments/{environment_id}/collections"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		name := "exampleString"
		createCollectionPath = strings.Replace(createCollectionPath, "{environment_id}", environmentID, 1)
		Context("Successfully - Create a collection", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createCollectionPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(201)
			}))
			It("Succeed to call CreateCollection", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.CreateCollection(nil)
				Expect(returnValueErr).NotTo(BeNil())

				createCollectionOptions := testService.NewCreateCollectionOptions(environmentID, name)
				returnValue, returnValueErr = testService.CreateCollection(createCollectionOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetCreateCollectionResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListCollections(listCollectionsOptions *ListCollectionsOptions)", func() {
		listCollectionsPath := "/v1/environments/{environment_id}/collections"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		listCollectionsPath = strings.Replace(listCollectionsPath, "{environment_id}", environmentID, 1)
		Context("Successfully - List collections", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listCollectionsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call ListCollections", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ListCollections(nil)
				Expect(returnValueErr).NotTo(BeNil())

				listCollectionsOptions := testService.NewListCollectionsOptions(environmentID)
				returnValue, returnValueErr = testService.ListCollections(listCollectionsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListCollectionsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetCollection(getCollectionOptions *GetCollectionOptions)", func() {
		getCollectionPath := "/v1/environments/{environment_id}/collections/{collection_id}"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		getCollectionPath = strings.Replace(getCollectionPath, "{environment_id}", environmentID, 1)
		getCollectionPath = strings.Replace(getCollectionPath, "{collection_id}", collectionID, 1)
		Context("Successfully - Get collection details", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getCollectionPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call GetCollection", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.GetCollection(nil)
				Expect(returnValueErr).NotTo(BeNil())

				getCollectionOptions := testService.NewGetCollectionOptions(environmentID, collectionID)
				returnValue, returnValueErr = testService.GetCollection(getCollectionOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetCollectionResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("UpdateCollection(updateCollectionOptions *UpdateCollectionOptions)", func() {
		updateCollectionPath := "/v1/environments/{environment_id}/collections/{collection_id}"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		updateCollectionPath = strings.Replace(updateCollectionPath, "{environment_id}", environmentID, 1)
		updateCollectionPath = strings.Replace(updateCollectionPath, "{collection_id}", collectionID, 1)
		Context("Successfully - Update a collection", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateCollectionPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(201)
			}))
			It("Succeed to call UpdateCollection", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.UpdateCollection(nil)
				Expect(returnValueErr).NotTo(BeNil())

				updateCollectionOptions := testService.NewUpdateCollectionOptions(environmentID, collectionID)
				returnValue, returnValueErr = testService.UpdateCollection(updateCollectionOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetUpdateCollectionResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteCollection(deleteCollectionOptions *DeleteCollectionOptions)", func() {
		deleteCollectionPath := "/v1/environments/{environment_id}/collections/{collection_id}"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		deleteCollectionPath = strings.Replace(deleteCollectionPath, "{environment_id}", environmentID, 1)
		deleteCollectionPath = strings.Replace(deleteCollectionPath, "{collection_id}", collectionID, 1)
		Context("Successfully - Delete a collection", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteCollectionPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"collection_id": "fake CollectionID", "status": "fake Status"}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call DeleteCollection", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.DeleteCollection(nil)
				Expect(returnValueErr).NotTo(BeNil())

				deleteCollectionOptions := testService.NewDeleteCollectionOptions(environmentID, collectionID)
				returnValue, returnValueErr = testService.DeleteCollection(deleteCollectionOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetDeleteCollectionResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListCollectionFields(listCollectionFieldsOptions *ListCollectionFieldsOptions)", func() {
		listCollectionFieldsPath := "/v1/environments/{environment_id}/collections/{collection_id}/fields"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		listCollectionFieldsPath = strings.Replace(listCollectionFieldsPath, "{environment_id}", environmentID, 1)
		listCollectionFieldsPath = strings.Replace(listCollectionFieldsPath, "{collection_id}", collectionID, 1)
		Context("Successfully - List collection fields", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listCollectionFieldsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call ListCollectionFields", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ListCollectionFields(nil)
				Expect(returnValueErr).NotTo(BeNil())

				listCollectionFieldsOptions := testService.NewListCollectionFieldsOptions(environmentID, collectionID)
				returnValue, returnValueErr = testService.ListCollectionFields(listCollectionFieldsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListCollectionFieldsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListExpansions(listExpansionsOptions *ListExpansionsOptions)", func() {
		listExpansionsPath := "/v1/environments/{environment_id}/collections/{collection_id}/expansions"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		listExpansionsPath = strings.Replace(listExpansionsPath, "{environment_id}", environmentID, 1)
		listExpansionsPath = strings.Replace(listExpansionsPath, "{collection_id}", collectionID, 1)
		Context("Successfully - Get the expansion list", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listExpansionsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"expansions": []}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call ListExpansions", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ListExpansions(nil)
				Expect(returnValueErr).NotTo(BeNil())

				listExpansionsOptions := testService.NewListExpansionsOptions(environmentID, collectionID)
				returnValue, returnValueErr = testService.ListExpansions(listExpansionsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListExpansionsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateExpansions(createExpansionsOptions *CreateExpansionsOptions)", func() {
		createExpansionsPath := "/v1/environments/{environment_id}/collections/{collection_id}/expansions"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		expansions := []discoveryv1.Expansion{}
		createExpansionsPath = strings.Replace(createExpansionsPath, "{environment_id}", environmentID, 1)
		createExpansionsPath = strings.Replace(createExpansionsPath, "{collection_id}", collectionID, 1)
		Context("Successfully - Create or update expansion list", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createExpansionsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"expansions": []}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call CreateExpansions", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.CreateExpansions(nil)
				Expect(returnValueErr).NotTo(BeNil())

				createExpansionsOptions := testService.NewCreateExpansionsOptions(environmentID, collectionID, expansions)
				returnValue, returnValueErr = testService.CreateExpansions(createExpansionsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetCreateExpansionsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteExpansions(deleteExpansionsOptions *DeleteExpansionsOptions)", func() {
		deleteExpansionsPath := "/v1/environments/{environment_id}/collections/{collection_id}/expansions"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		deleteExpansionsPath = strings.Replace(deleteExpansionsPath, "{environment_id}", environmentID, 1)
		deleteExpansionsPath = strings.Replace(deleteExpansionsPath, "{collection_id}", collectionID, 1)
		Context("Successfully - Delete the expansion list", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteExpansionsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.WriteHeader(204)
			}))
			It("Succeed to call DeleteExpansions", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.DeleteExpansions(nil)
				Expect(returnValueErr).NotTo(BeNil())

				deleteExpansionsOptions := testService.NewDeleteExpansionsOptions(environmentID, collectionID)
				returnValue, returnValueErr = testService.DeleteExpansions(deleteExpansionsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("GetTokenizationDictionaryStatus(getTokenizationDictionaryStatusOptions *GetTokenizationDictionaryStatusOptions)", func() {
		getTokenizationDictionaryStatusPath := "/v1/environments/{environment_id}/collections/{collection_id}/word_lists/tokenization_dictionary"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		getTokenizationDictionaryStatusPath = strings.Replace(getTokenizationDictionaryStatusPath, "{environment_id}", environmentID, 1)
		getTokenizationDictionaryStatusPath = strings.Replace(getTokenizationDictionaryStatusPath, "{collection_id}", collectionID, 1)
		Context("Successfully - Get tokenization dictionary status", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getTokenizationDictionaryStatusPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call GetTokenizationDictionaryStatus", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.GetTokenizationDictionaryStatus(nil)
				Expect(returnValueErr).NotTo(BeNil())

				getTokenizationDictionaryStatusOptions := testService.NewGetTokenizationDictionaryStatusOptions(environmentID, collectionID)
				returnValue, returnValueErr = testService.GetTokenizationDictionaryStatus(getTokenizationDictionaryStatusOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetTokenizationDictionaryStatusResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateTokenizationDictionary(createTokenizationDictionaryOptions *CreateTokenizationDictionaryOptions)", func() {
		createTokenizationDictionaryPath := "/v1/environments/{environment_id}/collections/{collection_id}/word_lists/tokenization_dictionary"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		createTokenizationDictionaryPath = strings.Replace(createTokenizationDictionaryPath, "{environment_id}", environmentID, 1)
		createTokenizationDictionaryPath = strings.Replace(createTokenizationDictionaryPath, "{collection_id}", collectionID, 1)
		Context("Successfully - Create tokenization dictionary", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createTokenizationDictionaryPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(202)
			}))
			It("Succeed to call CreateTokenizationDictionary", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.CreateTokenizationDictionary(nil)
				Expect(returnValueErr).NotTo(BeNil())

				createTokenizationDictionaryOptions := testService.NewCreateTokenizationDictionaryOptions(environmentID, collectionID)
				returnValue, returnValueErr = testService.CreateTokenizationDictionary(createTokenizationDictionaryOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetCreateTokenizationDictionaryResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteTokenizationDictionary(deleteTokenizationDictionaryOptions *DeleteTokenizationDictionaryOptions)", func() {
		deleteTokenizationDictionaryPath := "/v1/environments/{environment_id}/collections/{collection_id}/word_lists/tokenization_dictionary"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		deleteTokenizationDictionaryPath = strings.Replace(deleteTokenizationDictionaryPath, "{environment_id}", environmentID, 1)
		deleteTokenizationDictionaryPath = strings.Replace(deleteTokenizationDictionaryPath, "{collection_id}", collectionID, 1)
		Context("Successfully - Delete tokenization dictionary", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteTokenizationDictionaryPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.WriteHeader(200)
			}))
			It("Succeed to call DeleteTokenizationDictionary", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.DeleteTokenizationDictionary(nil)
				Expect(returnValueErr).NotTo(BeNil())

				deleteTokenizationDictionaryOptions := testService.NewDeleteTokenizationDictionaryOptions(environmentID, collectionID)
				returnValue, returnValueErr = testService.DeleteTokenizationDictionary(deleteTokenizationDictionaryOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("GetStopwordListStatus(getStopwordListStatusOptions *GetStopwordListStatusOptions)", func() {
		getStopwordListStatusPath := "/v1/environments/{environment_id}/collections/{collection_id}/word_lists/stopwords"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		getStopwordListStatusPath = strings.Replace(getStopwordListStatusPath, "{environment_id}", environmentID, 1)
		getStopwordListStatusPath = strings.Replace(getStopwordListStatusPath, "{collection_id}", collectionID, 1)
		Context("Successfully - Get stopword list status", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getStopwordListStatusPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call GetStopwordListStatus", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.GetStopwordListStatus(nil)
				Expect(returnValueErr).NotTo(BeNil())

				getStopwordListStatusOptions := testService.NewGetStopwordListStatusOptions(environmentID, collectionID)
				returnValue, returnValueErr = testService.GetStopwordListStatus(getStopwordListStatusOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetStopwordListStatusResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateStopwordList(createStopwordListOptions *CreateStopwordListOptions)", func() {
		createStopwordListPath := "/v1/environments/{environment_id}/collections/{collection_id}/word_lists/stopwords"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		pwd, _ := os.Getwd()
		stopwordsFile, stopwordsFileErr := os.Open(pwd + "/../resources/stopwords.txt")
		if stopwordsFileErr != nil {
			panic(stopwordsFileErr)
		}
		stopwordFilename := "exampleString"
		createStopwordListPath = strings.Replace(createStopwordListPath, "{environment_id}", environmentID, 1)
		createStopwordListPath = strings.Replace(createStopwordListPath, "{collection_id}", collectionID, 1)
		Context("Successfully - Create stopword list", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createStopwordListPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call CreateStopwordList", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.CreateStopwordList(nil)
				Expect(returnValueErr).NotTo(BeNil())

				createStopwordListOptions := testService.NewCreateStopwordListOptions(environmentID, collectionID, stopwordsFile, stopwordFilename)
				returnValue, returnValueErr = testService.CreateStopwordList(createStopwordListOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetCreateStopwordListResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteStopwordList(deleteStopwordListOptions *DeleteStopwordListOptions)", func() {
		deleteStopwordListPath := "/v1/environments/{environment_id}/collections/{collection_id}/word_lists/stopwords"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		deleteStopwordListPath = strings.Replace(deleteStopwordListPath, "{environment_id}", environmentID, 1)
		deleteStopwordListPath = strings.Replace(deleteStopwordListPath, "{collection_id}", collectionID, 1)
		Context("Successfully - Delete a custom stopword list", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteStopwordListPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.WriteHeader(200)
			}))
			It("Succeed to call DeleteStopwordList", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.DeleteStopwordList(nil)
				Expect(returnValueErr).NotTo(BeNil())

				deleteStopwordListOptions := testService.NewDeleteStopwordListOptions(environmentID, collectionID)
				returnValue, returnValueErr = testService.DeleteStopwordList(deleteStopwordListOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("AddDocument(addDocumentOptions *AddDocumentOptions)", func() {
		addDocumentPath := "/v1/environments/{environment_id}/collections/{collection_id}/documents"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		addDocumentPath = strings.Replace(addDocumentPath, "{environment_id}", environmentID, 1)
		addDocumentPath = strings.Replace(addDocumentPath, "{collection_id}", collectionID, 1)
		Context("Successfully - Add a document", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(addDocumentPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(202)
			}))
			It("Succeed to call AddDocument", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.AddDocument(nil)
				Expect(returnValueErr).NotTo(BeNil())

				addDocumentOptions := testService.NewAddDocumentOptions(environmentID, collectionID).
					SetMetadata("Name:John Smith")
				returnValue, returnValueErr = testService.AddDocument(addDocumentOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetAddDocumentResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetDocumentStatus(getDocumentStatusOptions *GetDocumentStatusOptions)", func() {
		getDocumentStatusPath := "/v1/environments/{environment_id}/collections/{collection_id}/documents/{document_id}"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		documentID := "exampleString"
		getDocumentStatusPath = strings.Replace(getDocumentStatusPath, "{environment_id}", environmentID, 1)
		getDocumentStatusPath = strings.Replace(getDocumentStatusPath, "{collection_id}", collectionID, 1)
		getDocumentStatusPath = strings.Replace(getDocumentStatusPath, "{document_id}", documentID, 1)
		Context("Successfully - Get document details", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getDocumentStatusPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"document_id": "fake DocumentID", "status": "fake Status", "status_description": "fake StatusDescription", "notices": []}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call GetDocumentStatus", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.GetDocumentStatus(nil)
				Expect(returnValueErr).NotTo(BeNil())

				getDocumentStatusOptions := testService.NewGetDocumentStatusOptions(environmentID, collectionID, documentID)
				returnValue, returnValueErr = testService.GetDocumentStatus(getDocumentStatusOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetDocumentStatusResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("UpdateDocument(updateDocumentOptions *UpdateDocumentOptions)", func() {
		updateDocumentPath := "/v1/environments/{environment_id}/collections/{collection_id}/documents/{document_id}"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		documentID := "exampleString"
		updateDocumentPath = strings.Replace(updateDocumentPath, "{environment_id}", environmentID, 1)
		updateDocumentPath = strings.Replace(updateDocumentPath, "{collection_id}", collectionID, 1)
		updateDocumentPath = strings.Replace(updateDocumentPath, "{document_id}", documentID, 1)
		Context("Successfully - Update a document", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateDocumentPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(202)
			}))
			It("Succeed to call UpdateDocument", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.UpdateDocument(nil)
				Expect(returnValueErr).NotTo(BeNil())

				updateDocumentOptions := testService.NewUpdateDocumentOptions(environmentID, collectionID, documentID).
					SetMetadata("Name:John Smith")
				returnValue, returnValueErr = testService.UpdateDocument(updateDocumentOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetUpdateDocumentResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteDocument(deleteDocumentOptions *DeleteDocumentOptions)", func() {
		deleteDocumentPath := "/v1/environments/{environment_id}/collections/{collection_id}/documents/{document_id}"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		documentID := "exampleString"
		deleteDocumentPath = strings.Replace(deleteDocumentPath, "{environment_id}", environmentID, 1)
		deleteDocumentPath = strings.Replace(deleteDocumentPath, "{collection_id}", collectionID, 1)
		deleteDocumentPath = strings.Replace(deleteDocumentPath, "{document_id}", documentID, 1)
		Context("Successfully - Delete a document", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteDocumentPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call DeleteDocument", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.DeleteDocument(nil)
				Expect(returnValueErr).NotTo(BeNil())

				deleteDocumentOptions := testService.NewDeleteDocumentOptions(environmentID, collectionID, documentID)
				returnValue, returnValueErr = testService.DeleteDocument(deleteDocumentOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetDeleteDocumentResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("Query(queryOptions *QueryOptions)", func() {
		queryPath := "/v1/environments/{environment_id}/collections/{collection_id}/query"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		queryPath = strings.Replace(queryPath, "{environment_id}", environmentID, 1)
		queryPath = strings.Replace(queryPath, "{collection_id}", collectionID, 1)
		Context("Successfully - Query a collection", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(queryPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call Query", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.Query(nil)
				Expect(returnValueErr).NotTo(BeNil())

				queryOptions := testService.NewQueryOptions(environmentID, collectionID)
				returnValue, returnValueErr = testService.Query(queryOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetQueryResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("QueryNotices(queryNoticesOptions *QueryNoticesOptions)", func() {
		queryNoticesPath := "/v1/environments/{environment_id}/collections/{collection_id}/notices"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		queryNoticesPath = strings.Replace(queryNoticesPath, "{environment_id}", environmentID, 1)
		queryNoticesPath = strings.Replace(queryNoticesPath, "{collection_id}", collectionID, 1)
		Context("Successfully - Query system notices", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(queryNoticesPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call QueryNotices", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.QueryNotices(nil)
				Expect(returnValueErr).NotTo(BeNil())

				queryNoticesOptions := testService.NewQueryNoticesOptions(environmentID, collectionID)
				returnValue, returnValueErr = testService.QueryNotices(queryNoticesOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetQueryNoticesResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("FederatedQuery(federatedQueryOptions *FederatedQueryOptions)", func() {
		federatedQueryPath := "/v1/environments/{environment_id}/query"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		federatedQueryPath = strings.Replace(federatedQueryPath, "{environment_id}", environmentID, 1)
		Context("Successfully - Query multiple collections", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(federatedQueryPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call FederatedQuery", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.FederatedQuery(nil)
				Expect(returnValueErr).NotTo(BeNil())

				federatedQueryOptions := testService.NewFederatedQueryOptions(environmentID)
				returnValue, returnValueErr = testService.FederatedQuery(federatedQueryOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetFederatedQueryResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("FederatedQueryNotices(federatedQueryNoticesOptions *FederatedQueryNoticesOptions)", func() {
		federatedQueryNoticesPath := "/v1/environments/{environment_id}/notices"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		collectionIds := []string{}
		federatedQueryNoticesPath = strings.Replace(federatedQueryNoticesPath, "{environment_id}", environmentID, 1)
		Context("Successfully - Query multiple collection system notices", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(federatedQueryNoticesPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call FederatedQueryNotices", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.FederatedQueryNotices(nil)
				Expect(returnValueErr).NotTo(BeNil())

				federatedQueryNoticesOptions := testService.NewFederatedQueryNoticesOptions(environmentID, collectionIds)
				returnValue, returnValueErr = testService.FederatedQueryNotices(federatedQueryNoticesOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetFederatedQueryNoticesResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("QueryEntities(queryEntitiesOptions *QueryEntitiesOptions)", func() {
		queryEntitiesPath := "/v1/environments/{environment_id}/collections/{collection_id}/query_entities"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		queryEntitiesPath = strings.Replace(queryEntitiesPath, "{environment_id}", environmentID, 1)
		queryEntitiesPath = strings.Replace(queryEntitiesPath, "{collection_id}", collectionID, 1)
		Context("Successfully - Knowledge Graph entity query", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(queryEntitiesPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call QueryEntities", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.QueryEntities(nil)
				Expect(returnValueErr).NotTo(BeNil())

				queryEntitiesOptions := testService.NewQueryEntitiesOptions(environmentID, collectionID)
				returnValue, returnValueErr = testService.QueryEntities(queryEntitiesOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetQueryEntitiesResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("QueryRelations(queryRelationsOptions *QueryRelationsOptions)", func() {
		queryRelationsPath := "/v1/environments/{environment_id}/collections/{collection_id}/query_relations"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		queryRelationsPath = strings.Replace(queryRelationsPath, "{environment_id}", environmentID, 1)
		queryRelationsPath = strings.Replace(queryRelationsPath, "{collection_id}", collectionID, 1)
		Context("Successfully - Knowledge Graph relationship query", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(queryRelationsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call QueryRelations", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.QueryRelations(nil)
				Expect(returnValueErr).NotTo(BeNil())

				queryRelationsOptions := testService.NewQueryRelationsOptions(environmentID, collectionID)
				returnValue, returnValueErr = testService.QueryRelations(queryRelationsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetQueryRelationsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListTrainingData(listTrainingDataOptions *ListTrainingDataOptions)", func() {
		listTrainingDataPath := "/v1/environments/{environment_id}/collections/{collection_id}/training_data"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		listTrainingDataPath = strings.Replace(listTrainingDataPath, "{environment_id}", environmentID, 1)
		listTrainingDataPath = strings.Replace(listTrainingDataPath, "{collection_id}", collectionID, 1)
		Context("Successfully - List training data", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listTrainingDataPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call ListTrainingData", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ListTrainingData(nil)
				Expect(returnValueErr).NotTo(BeNil())

				listTrainingDataOptions := testService.NewListTrainingDataOptions(environmentID, collectionID)
				returnValue, returnValueErr = testService.ListTrainingData(listTrainingDataOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListTrainingDataResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("AddTrainingData(addTrainingDataOptions *AddTrainingDataOptions)", func() {
		addTrainingDataPath := "/v1/environments/{environment_id}/collections/{collection_id}/training_data"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		addTrainingDataPath = strings.Replace(addTrainingDataPath, "{environment_id}", environmentID, 1)
		addTrainingDataPath = strings.Replace(addTrainingDataPath, "{collection_id}", collectionID, 1)
		Context("Successfully - Add query to training data", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(addTrainingDataPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call AddTrainingData", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.AddTrainingData(nil)
				Expect(returnValueErr).NotTo(BeNil())

				addTrainingDataOptions := testService.NewAddTrainingDataOptions(environmentID, collectionID)
				returnValue, returnValueErr = testService.AddTrainingData(addTrainingDataOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetAddTrainingDataResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteAllTrainingData(deleteAllTrainingDataOptions *DeleteAllTrainingDataOptions)", func() {
		deleteAllTrainingDataPath := "/v1/environments/{environment_id}/collections/{collection_id}/training_data"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		deleteAllTrainingDataPath = strings.Replace(deleteAllTrainingDataPath, "{environment_id}", environmentID, 1)
		deleteAllTrainingDataPath = strings.Replace(deleteAllTrainingDataPath, "{collection_id}", collectionID, 1)
		Context("Successfully - Delete all training data", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteAllTrainingDataPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.WriteHeader(204)
			}))
			It("Succeed to call DeleteAllTrainingData", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.DeleteAllTrainingData(nil)
				Expect(returnValueErr).NotTo(BeNil())

				deleteAllTrainingDataOptions := testService.NewDeleteAllTrainingDataOptions(environmentID, collectionID)
				returnValue, returnValueErr = testService.DeleteAllTrainingData(deleteAllTrainingDataOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("GetTrainingData(getTrainingDataOptions *GetTrainingDataOptions)", func() {
		getTrainingDataPath := "/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		queryID := "exampleString"
		getTrainingDataPath = strings.Replace(getTrainingDataPath, "{environment_id}", environmentID, 1)
		getTrainingDataPath = strings.Replace(getTrainingDataPath, "{collection_id}", collectionID, 1)
		getTrainingDataPath = strings.Replace(getTrainingDataPath, "{query_id}", queryID, 1)
		Context("Successfully - Get details about a query", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getTrainingDataPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call GetTrainingData", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.GetTrainingData(nil)
				Expect(returnValueErr).NotTo(BeNil())

				getTrainingDataOptions := testService.NewGetTrainingDataOptions(environmentID, collectionID, queryID)
				returnValue, returnValueErr = testService.GetTrainingData(getTrainingDataOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetTrainingDataResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteTrainingData(deleteTrainingDataOptions *DeleteTrainingDataOptions)", func() {
		deleteTrainingDataPath := "/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		queryID := "exampleString"
		deleteTrainingDataPath = strings.Replace(deleteTrainingDataPath, "{environment_id}", environmentID, 1)
		deleteTrainingDataPath = strings.Replace(deleteTrainingDataPath, "{collection_id}", collectionID, 1)
		deleteTrainingDataPath = strings.Replace(deleteTrainingDataPath, "{query_id}", queryID, 1)
		Context("Successfully - Delete a training data query", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteTrainingDataPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.WriteHeader(204)
			}))
			It("Succeed to call DeleteTrainingData", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.DeleteTrainingData(nil)
				Expect(returnValueErr).NotTo(BeNil())

				deleteTrainingDataOptions := testService.NewDeleteTrainingDataOptions(environmentID, collectionID, queryID)
				returnValue, returnValueErr = testService.DeleteTrainingData(deleteTrainingDataOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("ListTrainingExamples(listTrainingExamplesOptions *ListTrainingExamplesOptions)", func() {
		listTrainingExamplesPath := "/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}/examples"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		queryID := "exampleString"
		listTrainingExamplesPath = strings.Replace(listTrainingExamplesPath, "{environment_id}", environmentID, 1)
		listTrainingExamplesPath = strings.Replace(listTrainingExamplesPath, "{collection_id}", collectionID, 1)
		listTrainingExamplesPath = strings.Replace(listTrainingExamplesPath, "{query_id}", queryID, 1)
		Context("Successfully - List examples for a training data query", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listTrainingExamplesPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call ListTrainingExamples", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ListTrainingExamples(nil)
				Expect(returnValueErr).NotTo(BeNil())

				listTrainingExamplesOptions := testService.NewListTrainingExamplesOptions(environmentID, collectionID, queryID)
				returnValue, returnValueErr = testService.ListTrainingExamples(listTrainingExamplesOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListTrainingExamplesResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateTrainingExample(createTrainingExampleOptions *CreateTrainingExampleOptions)", func() {
		createTrainingExamplePath := "/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}/examples"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		queryID := "exampleString"
		createTrainingExamplePath = strings.Replace(createTrainingExamplePath, "{environment_id}", environmentID, 1)
		createTrainingExamplePath = strings.Replace(createTrainingExamplePath, "{collection_id}", collectionID, 1)
		createTrainingExamplePath = strings.Replace(createTrainingExamplePath, "{query_id}", queryID, 1)
		Context("Successfully - Add example to training data query", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createTrainingExamplePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(201)
			}))
			It("Succeed to call CreateTrainingExample", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.CreateTrainingExample(nil)
				Expect(returnValueErr).NotTo(BeNil())

				createTrainingExampleOptions := testService.NewCreateTrainingExampleOptions(environmentID, collectionID, queryID)
				returnValue, returnValueErr = testService.CreateTrainingExample(createTrainingExampleOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetCreateTrainingExampleResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteTrainingExample(deleteTrainingExampleOptions *DeleteTrainingExampleOptions)", func() {
		deleteTrainingExamplePath := "/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}/examples/{example_id}"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		queryID := "exampleString"
		exampleID := "exampleString"
		deleteTrainingExamplePath = strings.Replace(deleteTrainingExamplePath, "{environment_id}", environmentID, 1)
		deleteTrainingExamplePath = strings.Replace(deleteTrainingExamplePath, "{collection_id}", collectionID, 1)
		deleteTrainingExamplePath = strings.Replace(deleteTrainingExamplePath, "{query_id}", queryID, 1)
		deleteTrainingExamplePath = strings.Replace(deleteTrainingExamplePath, "{example_id}", exampleID, 1)
		Context("Successfully - Delete example for training data query", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteTrainingExamplePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.WriteHeader(204)
			}))
			It("Succeed to call DeleteTrainingExample", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.DeleteTrainingExample(nil)
				Expect(returnValueErr).NotTo(BeNil())

				deleteTrainingExampleOptions := testService.NewDeleteTrainingExampleOptions(environmentID, collectionID, queryID, exampleID)
				returnValue, returnValueErr = testService.DeleteTrainingExample(deleteTrainingExampleOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("UpdateTrainingExample(updateTrainingExampleOptions *UpdateTrainingExampleOptions)", func() {
		updateTrainingExamplePath := "/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}/examples/{example_id}"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		queryID := "exampleString"
		exampleID := "exampleString"
		updateTrainingExamplePath = strings.Replace(updateTrainingExamplePath, "{environment_id}", environmentID, 1)
		updateTrainingExamplePath = strings.Replace(updateTrainingExamplePath, "{collection_id}", collectionID, 1)
		updateTrainingExamplePath = strings.Replace(updateTrainingExamplePath, "{query_id}", queryID, 1)
		updateTrainingExamplePath = strings.Replace(updateTrainingExamplePath, "{example_id}", exampleID, 1)
		Context("Successfully - Change label or cross reference for example", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateTrainingExamplePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call UpdateTrainingExample", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.UpdateTrainingExample(nil)
				Expect(returnValueErr).NotTo(BeNil())

				updateTrainingExampleOptions := testService.NewUpdateTrainingExampleOptions(environmentID, collectionID, queryID, exampleID)
				returnValue, returnValueErr = testService.UpdateTrainingExample(updateTrainingExampleOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetUpdateTrainingExampleResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetTrainingExample(getTrainingExampleOptions *GetTrainingExampleOptions)", func() {
		getTrainingExamplePath := "/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}/examples/{example_id}"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		collectionID := "exampleString"
		queryID := "exampleString"
		exampleID := "exampleString"
		getTrainingExamplePath = strings.Replace(getTrainingExamplePath, "{environment_id}", environmentID, 1)
		getTrainingExamplePath = strings.Replace(getTrainingExamplePath, "{collection_id}", collectionID, 1)
		getTrainingExamplePath = strings.Replace(getTrainingExamplePath, "{query_id}", queryID, 1)
		getTrainingExamplePath = strings.Replace(getTrainingExamplePath, "{example_id}", exampleID, 1)
		Context("Successfully - Get details for training data example", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getTrainingExamplePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call GetTrainingExample", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.GetTrainingExample(nil)
				Expect(returnValueErr).NotTo(BeNil())

				getTrainingExampleOptions := testService.NewGetTrainingExampleOptions(environmentID, collectionID, queryID, exampleID)
				returnValue, returnValueErr = testService.GetTrainingExample(getTrainingExampleOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetTrainingExampleResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteUserData(deleteUserDataOptions *DeleteUserDataOptions)", func() {
		deleteUserDataPath := "/v1/user_data"
		version := "exampleString"
		accessToken := "0ui9876453"
		customerID := "exampleString"
		Context("Successfully - Delete labeled data", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteUserDataPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				Expect(req.URL.Query()["customer_id"]).To(Equal([]string{customerID}))

				res.WriteHeader(200)
			}))
			It("Succeed to call DeleteUserData", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.DeleteUserData(nil)
				Expect(returnValueErr).NotTo(BeNil())

				deleteUserDataOptions := testService.NewDeleteUserDataOptions(customerID)
				returnValue, returnValueErr = testService.DeleteUserData(deleteUserDataOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("CreateEvent(createEventOptions *CreateEventOptions)", func() {
		createEventPath := "/v1/events"
		version := "exampleString"
		accessToken := "0ui9876453"
		typeVar := "exampleString"
		data := &discoveryv1.EventData{EnvironmentID: &typeVar, SessionToken: &typeVar, CollectionID: &typeVar, DocumentID: &typeVar}
		Context("Successfully - Create event", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createEventPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(201)
			}))
			It("Succeed to call CreateEvent", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.CreateEvent(nil)
				Expect(returnValueErr).NotTo(BeNil())

				createEventOptions := testService.NewCreateEventOptions(typeVar, data)
				returnValue, returnValueErr = testService.CreateEvent(createEventOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetCreateEventResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("QueryLog(queryLogOptions *QueryLogOptions)", func() {
		queryLogPath := "/v1/logs"
		version := "exampleString"
		accessToken := "0ui9876453"
		Context("Successfully - Search the query and event log", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(queryLogPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call QueryLog", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.QueryLog(nil)
				Expect(returnValueErr).NotTo(BeNil())

				queryLogOptions := testService.NewQueryLogOptions()
				returnValue, returnValueErr = testService.QueryLog(queryLogOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetQueryLogResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetMetricsQuery(getMetricsQueryOptions *GetMetricsQueryOptions)", func() {
		getMetricsQueryPath := "/v1/metrics/number_of_queries"
		version := "exampleString"
		accessToken := "0ui9876453"
		Context("Successfully - Number of queries over time", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getMetricsQueryPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call GetMetricsQuery", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.GetMetricsQuery(nil)
				Expect(returnValueErr).NotTo(BeNil())

				getMetricsQueryOptions := testService.NewGetMetricsQueryOptions()
				returnValue, returnValueErr = testService.GetMetricsQuery(getMetricsQueryOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetMetricsQueryResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetMetricsQueryEvent(getMetricsQueryEventOptions *GetMetricsQueryEventOptions)", func() {
		getMetricsQueryEventPath := "/v1/metrics/number_of_queries_with_event"
		version := "exampleString"
		accessToken := "0ui9876453"
		Context("Successfully - Number of queries with an event over time", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getMetricsQueryEventPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call GetMetricsQueryEvent", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.GetMetricsQueryEvent(nil)
				Expect(returnValueErr).NotTo(BeNil())

				getMetricsQueryEventOptions := testService.NewGetMetricsQueryEventOptions()
				returnValue, returnValueErr = testService.GetMetricsQueryEvent(getMetricsQueryEventOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetMetricsQueryEventResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetMetricsQueryNoResults(getMetricsQueryNoResultsOptions *GetMetricsQueryNoResultsOptions)", func() {
		getMetricsQueryNoResultsPath := "/v1/metrics/number_of_queries_with_no_search_results"
		version := "exampleString"
		accessToken := "0ui9876453"
		Context("Successfully - Number of queries with no search results over time", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getMetricsQueryNoResultsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call GetMetricsQueryNoResults", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.GetMetricsQueryNoResults(nil)
				Expect(returnValueErr).NotTo(BeNil())

				getMetricsQueryNoResultsOptions := testService.NewGetMetricsQueryNoResultsOptions()
				returnValue, returnValueErr = testService.GetMetricsQueryNoResults(getMetricsQueryNoResultsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetMetricsQueryNoResultsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetMetricsEventRate(getMetricsEventRateOptions *GetMetricsEventRateOptions)", func() {
		getMetricsEventRatePath := "/v1/metrics/event_rate"
		version := "exampleString"
		accessToken := "0ui9876453"
		Context("Successfully - Percentage of queries with an associated event", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getMetricsEventRatePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call GetMetricsEventRate", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.GetMetricsEventRate(nil)
				Expect(returnValueErr).NotTo(BeNil())

				getMetricsEventRateOptions := testService.NewGetMetricsEventRateOptions()
				returnValue, returnValueErr = testService.GetMetricsEventRate(getMetricsEventRateOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetMetricsEventRateResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetMetricsQueryTokenEvent(getMetricsQueryTokenEventOptions *GetMetricsQueryTokenEventOptions)", func() {
		getMetricsQueryTokenEventPath := "/v1/metrics/top_query_tokens_with_event_rate"
		version := "exampleString"
		accessToken := "0ui9876453"
		Context("Successfully - Most frequent query tokens with an event", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getMetricsQueryTokenEventPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call GetMetricsQueryTokenEvent", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.GetMetricsQueryTokenEvent(nil)
				Expect(returnValueErr).NotTo(BeNil())

				getMetricsQueryTokenEventOptions := testService.NewGetMetricsQueryTokenEventOptions()
				returnValue, returnValueErr = testService.GetMetricsQueryTokenEvent(getMetricsQueryTokenEventOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetMetricsQueryTokenEventResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListCredentials(listCredentialsOptions *ListCredentialsOptions)", func() {
		listCredentialsPath := "/v1/environments/{environment_id}/credentials"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		listCredentialsPath = strings.Replace(listCredentialsPath, "{environment_id}", environmentID, 1)
		Context("Successfully - List credentials", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listCredentialsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call ListCredentials", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ListCredentials(nil)
				Expect(returnValueErr).NotTo(BeNil())

				listCredentialsOptions := testService.NewListCredentialsOptions(environmentID)
				returnValue, returnValueErr = testService.ListCredentials(listCredentialsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListCredentialsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateCredentials(createCredentialsOptions *CreateCredentialsOptions)", func() {
		createCredentialsPath := "/v1/environments/{environment_id}/credentials"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		createCredentialsPath = strings.Replace(createCredentialsPath, "{environment_id}", environmentID, 1)
		Context("Successfully - Create credentials", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createCredentialsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call CreateCredentials", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.CreateCredentials(nil)
				Expect(returnValueErr).NotTo(BeNil())

				createCredentialsOptions := testService.NewCreateCredentialsOptions(environmentID)
				returnValue, returnValueErr = testService.CreateCredentials(createCredentialsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetCreateCredentialsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetCredentials(getCredentialsOptions *GetCredentialsOptions)", func() {
		getCredentialsPath := "/v1/environments/{environment_id}/credentials/{credential_id}"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		credentialID := "exampleString"
		getCredentialsPath = strings.Replace(getCredentialsPath, "{environment_id}", environmentID, 1)
		getCredentialsPath = strings.Replace(getCredentialsPath, "{credential_id}", credentialID, 1)
		Context("Successfully - View Credentials", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getCredentialsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call GetCredentials", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.GetCredentials(nil)
				Expect(returnValueErr).NotTo(BeNil())

				getCredentialsOptions := testService.NewGetCredentialsOptions(environmentID, credentialID)
				returnValue, returnValueErr = testService.GetCredentials(getCredentialsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetCredentialsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("UpdateCredentials(updateCredentialsOptions *UpdateCredentialsOptions)", func() {
		updateCredentialsPath := "/v1/environments/{environment_id}/credentials/{credential_id}"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		credentialID := "exampleString"
		updateCredentialsPath = strings.Replace(updateCredentialsPath, "{environment_id}", environmentID, 1)
		updateCredentialsPath = strings.Replace(updateCredentialsPath, "{credential_id}", credentialID, 1)
		Context("Successfully - Update credentials", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateCredentialsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call UpdateCredentials", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.UpdateCredentials(nil)
				Expect(returnValueErr).NotTo(BeNil())

				updateCredentialsOptions := testService.NewUpdateCredentialsOptions(environmentID, credentialID)
				returnValue, returnValueErr = testService.UpdateCredentials(updateCredentialsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetUpdateCredentialsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteCredentials(deleteCredentialsOptions *DeleteCredentialsOptions)", func() {
		deleteCredentialsPath := "/v1/environments/{environment_id}/credentials/{credential_id}"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		credentialID := "exampleString"
		deleteCredentialsPath = strings.Replace(deleteCredentialsPath, "{environment_id}", environmentID, 1)
		deleteCredentialsPath = strings.Replace(deleteCredentialsPath, "{credential_id}", credentialID, 1)
		Context("Successfully - Delete credentials", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteCredentialsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call DeleteCredentials", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.DeleteCredentials(nil)
				Expect(returnValueErr).NotTo(BeNil())

				deleteCredentialsOptions := testService.NewDeleteCredentialsOptions(environmentID, credentialID)
				returnValue, returnValueErr = testService.DeleteCredentials(deleteCredentialsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetDeleteCredentialsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListGateways(listGatewaysOptions *ListGatewaysOptions)", func() {
		listGatewaysPath := "/v1/environments/{environment_id}/gateways"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		listGatewaysPath = strings.Replace(listGatewaysPath, "{environment_id}", environmentID, 1)
		Context("Successfully - List Gateways", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listGatewaysPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call ListGateways", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ListGateways(nil)
				Expect(returnValueErr).NotTo(BeNil())

				listGatewaysOptions := testService.NewListGatewaysOptions(environmentID)
				returnValue, returnValueErr = testService.ListGateways(listGatewaysOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListGatewaysResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateGateway(createGatewayOptions *CreateGatewayOptions)", func() {
		createGatewayPath := "/v1/environments/{environment_id}/gateways"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		createGatewayPath = strings.Replace(createGatewayPath, "{environment_id}", environmentID, 1)
		Context("Successfully - Create Gateway", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createGatewayPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call CreateGateway", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.CreateGateway(nil)
				Expect(returnValueErr).NotTo(BeNil())

				createGatewayOptions := testService.NewCreateGatewayOptions(environmentID)
				returnValue, returnValueErr = testService.CreateGateway(createGatewayOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetCreateGatewayResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetGateway(getGatewayOptions *GetGatewayOptions)", func() {
		getGatewayPath := "/v1/environments/{environment_id}/gateways/{gateway_id}"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		gatewayID := "exampleString"
		getGatewayPath = strings.Replace(getGatewayPath, "{environment_id}", environmentID, 1)
		getGatewayPath = strings.Replace(getGatewayPath, "{gateway_id}", gatewayID, 1)
		Context("Successfully - List Gateway Details", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getGatewayPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call GetGateway", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.GetGateway(nil)
				Expect(returnValueErr).NotTo(BeNil())

				getGatewayOptions := testService.NewGetGatewayOptions(environmentID, gatewayID)
				returnValue, returnValueErr = testService.GetGateway(getGatewayOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetGatewayResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteGateway(deleteGatewayOptions *DeleteGatewayOptions)", func() {
		deleteGatewayPath := "/v1/environments/{environment_id}/gateways/{gateway_id}"
		version := "exampleString"
		accessToken := "0ui9876453"
		environmentID := "exampleString"
		gatewayID := "exampleString"
		deleteGatewayPath = strings.Replace(deleteGatewayPath, "{environment_id}", environmentID, 1)
		deleteGatewayPath = strings.Replace(deleteGatewayPath, "{gateway_id}", gatewayID, 1)
		Context("Successfully - Delete Gateway", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteGatewayPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + accessToken))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
				res.WriteHeader(200)
			}))
			It("Succeed to call DeleteGateway", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL: testServer.URL,
					Version: version,
                    Authenticator: &core.BearerTokenAuthenticator{
                        BearerToken: accessToken,
                    },
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.DeleteGateway(nil)
				Expect(returnValueErr).NotTo(BeNil())

				deleteGatewayOptions := testService.NewDeleteGatewayOptions(environmentID, gatewayID)
				returnValue, returnValueErr = testService.DeleteGateway(deleteGatewayOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetDeleteGatewayResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
})
