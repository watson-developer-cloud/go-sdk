package discoveryv1_test

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	"github.com/IBM/go-sdk-core/core"

	discoveryv1 "github.com/watson-developer-cloud/go-sdk/discoveryv1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DiscoveryV1", func() {
	Describe("CreateEnvironment(options *CreateEnvironmentOptions)", func() {
		createEnvironmentPath := "/v1/environments"
		version := "exampleString"
		name := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Create an environment", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(createEnvironmentPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(createEnvironmentPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"environment_id":"xxx"}`)
			}))
			It("Succeed to call CreateEnvironment", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// pass empty options
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
	Describe("DeleteEnvironment(options *DeleteEnvironmentOptions)", func() {
		deleteEnvironmentPath := "/v1/environments/{environment_id}"
		version := "exampleString"
		environmentID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		deleteEnvironmentPath = strings.Replace(deleteEnvironmentPath, "{environment_id}", environmentID, 1)
		Context("Successfully - Delete environment", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(deleteEnvironmentPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(deleteEnvironmentPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"environment_id":"xxx", "status":"active"}`)
			}))
			It("Succeed to call DeleteEnvironment", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				deleteEnvironmentOptions := testService.NewDeleteEnvironmentOptions(environmentID)
				returnValue, returnValueErr := testService.DeleteEnvironment(deleteEnvironmentOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetDeleteEnvironmentResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetEnvironment(options *GetEnvironmentOptions)", func() {
		getEnvironmentPath := "/v1/environments/{environment_id}"
		version := "exampleString"
		environmentID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		getEnvironmentPath = strings.Replace(getEnvironmentPath, "{environment_id}", environmentID, 1)
		Context("Successfully - Get environment info", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(getEnvironmentPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(getEnvironmentPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
			}))
			It("Succeed to call GetEnvironment", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				getEnvironmentOptions := testService.NewGetEnvironmentOptions(environmentID)
				returnValue, returnValueErr := testService.GetEnvironment(getEnvironmentOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetEnvironmentResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListEnvironments(options *ListEnvironmentsOptions)", func() {
		listEnvironmentsPath := "/v1/environments"
		version := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - List environments", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(listEnvironmentsPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(listEnvironmentsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `[]`)
			}))
			It("Succeed to call ListEnvironments", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				listEnvironmentsOptions := testService.NewListEnvironmentsOptions()
				returnValue, returnValueErr := testService.ListEnvironments(listEnvironmentsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListEnvironmentsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListFields(options *ListFieldsOptions)", func() {
		listFieldsPath := "/v1/environments/{environment_id}/fields"
		version := "exampleString"
		environmentID := "exampleString"
		collectionIds := []string{}
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		listFieldsPath = strings.Replace(listFieldsPath, "{environment_id}", environmentID, 1)
		Context("Successfully - List fields across collections", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(listFieldsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `[]`)
			}))
			It("Succeed to call ListFields", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				listFieldsOptions := testService.NewListFieldsOptions(environmentID, collectionIds)
				returnValue, returnValueErr := testService.ListFields(listFieldsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListFieldsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("UpdateEnvironment(options *UpdateEnvironmentOptions)", func() {
		updateEnvironmentPath := "/v1/environments/{environment_id}"
		version := "exampleString"
		environmentID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		updateEnvironmentPath = strings.Replace(updateEnvironmentPath, "{environment_id}", environmentID, 1)
		Context("Successfully - Update an environment", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(updateEnvironmentPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(updateEnvironmentPath))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"environment_id":"xxx", "name":"yyy"}`)
			}))
			It("Succeed to call UpdateEnvironment", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				updateEnvironmentOptions := testService.NewUpdateEnvironmentOptions(environmentID)
				returnValue, returnValueErr := testService.UpdateEnvironment(updateEnvironmentOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetUpdateEnvironmentResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateConfiguration(options *CreateConfigurationOptions)", func() {
		createConfigurationPath := "/v1/environments/{environment_id}/configurations"
		version := "exampleString"
		environmentID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		createConfigurationPath = strings.Replace(createConfigurationPath, "{environment_id}", environmentID, 1)
		Context("Successfully - Add configuration", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(createConfigurationPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(createConfigurationPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"configuration_id":"xxx"}`)
			}))
			It("Succeed to call CreateConfiguration", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				createConfigurationOptions := testService.
					NewCreateConfigurationOptions(environmentID, "test configuration id")
				returnValue, returnValueErr := testService.CreateConfiguration(createConfigurationOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetCreateConfigurationResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteConfiguration(options *DeleteConfigurationOptions)", func() {
		deleteConfigurationPath := "/v1/environments/{environment_id}/configurations/{configuration_id}"
		version := "exampleString"
		environmentID := "exampleString"
		configurationID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		deleteConfigurationPath = strings.Replace(deleteConfigurationPath, "{environment_id}", environmentID, 1)
		deleteConfigurationPath = strings.Replace(deleteConfigurationPath, "{configuration_id}", configurationID, 1)
		Context("Successfully - Delete a configuration", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(deleteConfigurationPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(deleteConfigurationPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"configuration_id":"xxx", "status": "active"}`)
			}))
			It("Succeed to call DeleteConfiguration", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				deleteConfigurationOptions := testService.NewDeleteConfigurationOptions(environmentID, configurationID)
				returnValue, returnValueErr := testService.DeleteConfiguration(deleteConfigurationOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetDeleteConfigurationResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetConfiguration(options *GetConfigurationOptions)", func() {
		getConfigurationPath := "/v1/environments/{environment_id}/configurations/{configuration_id}"
		version := "exampleString"
		environmentID := "exampleString"
		configurationID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		getConfigurationPath = strings.Replace(getConfigurationPath, "{environment_id}", environmentID, 1)
		getConfigurationPath = strings.Replace(getConfigurationPath, "{configuration_id}", configurationID, 1)
		Context("Successfully - Get configuration details", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(getConfigurationPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(getConfigurationPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
			}))
			It("Succeed to call GetConfiguration", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				getConfigurationOptions := testService.NewGetConfigurationOptions(environmentID, configurationID)
				returnValue, returnValueErr := testService.GetConfiguration(getConfigurationOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetConfigurationResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListConfigurations(options *ListConfigurationsOptions)", func() {
		listConfigurationsPath := "/v1/environments/{environment_id}/configurations"
		version := "exampleString"
		environmentID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		listConfigurationsPath = strings.Replace(listConfigurationsPath, "{environment_id}", environmentID, 1)
		Context("Successfully - List configurations", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(listConfigurationsPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(listConfigurationsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `[]`)
			}))
			It("Succeed to call ListConfigurations", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				listConfigurationsOptions := testService.NewListConfigurationsOptions(environmentID)
				returnValue, returnValueErr := testService.ListConfigurations(listConfigurationsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListConfigurationsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("UpdateConfiguration(options *UpdateConfigurationOptions)", func() {
		updateConfigurationPath := "/v1/environments/{environment_id}/configurations/{configuration_id}"
		version := "exampleString"
		environmentID := "exampleString"
		configurationID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		updateConfigurationPath = strings.Replace(updateConfigurationPath, "{environment_id}", environmentID, 1)
		updateConfigurationPath = strings.Replace(updateConfigurationPath, "{configuration_id}", configurationID, 1)
		Context("Successfully - Update a configuration", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(updateConfigurationPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(updateConfigurationPath))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
			}))
			It("Succeed to call UpdateConfiguration", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				updateConfigurationOptions := testService.
					NewUpdateConfigurationOptions(environmentID, configurationID, "new configuration")
				returnValue, returnValueErr := testService.UpdateConfiguration(updateConfigurationOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetUpdateConfigurationResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("TestConfigurationInEnvironment(options *TestConfigurationInEnvironmentOptions)", func() {
		testConfigurationInEnvironmentPath := "/v1/environments/{environment_id}/preview"
		version := "exampleString"
		environmentID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		testConfigurationInEnvironmentPath = strings.Replace(testConfigurationInEnvironmentPath, "{environment_id}", environmentID, 1)
		Context("Successfully - Test configuration", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(testConfigurationInEnvironmentPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(testConfigurationInEnvironmentPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"configuration_id":"xxx"}`)
			}))
			It("Succeed to call TestConfigurationInEnvironment", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				testConfigurationInEnvironmentOptions := testService.NewTestConfigurationInEnvironmentOptions(environmentID).SetMetadata("Name:John Smith")
				returnValue, returnValueErr := testService.TestConfigurationInEnvironment(testConfigurationInEnvironmentOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetTestConfigurationInEnvironmentResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateCollection(options *CreateCollectionOptions)", func() {
		createCollectionPath := "/v1/environments/{environment_id}/collections"
		version := "exampleString"
		environmentID := "exampleString"
		name := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		createCollectionPath = strings.Replace(createCollectionPath, "{environment_id}", environmentID, 1)
		Context("Successfully - Create a collection", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(createCollectionPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(createCollectionPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"collection_id":"xxx"}`)
			}))
			It("Succeed to call CreateCollection", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				createCollectionOptions := testService.NewCreateCollectionOptions(environmentID, name)
				returnValue, returnValueErr := testService.CreateCollection(createCollectionOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetCreateCollectionResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteCollection(options *DeleteCollectionOptions)", func() {
		deleteCollectionPath := "/v1/environments/{environment_id}/collections/{collection_id}"
		version := "exampleString"
		environmentID := "exampleString"
		collectionID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		deleteCollectionPath = strings.Replace(deleteCollectionPath, "{environment_id}", environmentID, 1)
		deleteCollectionPath = strings.Replace(deleteCollectionPath, "{collection_id}", collectionID, 1)
		Context("Successfully - Delete a collection", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(deleteCollectionPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(deleteCollectionPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"collection_id":"xxx"}`)
			}))
			It("Succeed to call DeleteCollection", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				deleteCollectionOptions := testService.NewDeleteCollectionOptions(environmentID, collectionID)
				returnValue, returnValueErr := testService.DeleteCollection(deleteCollectionOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetDeleteCollectionResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetCollection(options *GetCollectionOptions)", func() {
		getCollectionPath := "/v1/environments/{environment_id}/collections/{collection_id}"
		version := "exampleString"
		environmentID := "exampleString"
		collectionID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		getCollectionPath = strings.Replace(getCollectionPath, "{environment_id}", environmentID, 1)
		getCollectionPath = strings.Replace(getCollectionPath, "{collection_id}", collectionID, 1)
		Context("Successfully - Get collection details", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(getCollectionPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(getCollectionPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"collection_id":"xxx"}`)
			}))
			It("Succeed to call GetCollection", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				getCollectionOptions := testService.NewGetCollectionOptions(environmentID, collectionID)
				returnValue, returnValueErr := testService.GetCollection(getCollectionOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetCollectionResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListCollectionFields(options *ListCollectionFieldsOptions)", func() {
		listCollectionFieldsPath := "/v1/environments/{environment_id}/collections/{collection_id}/fields"
		version := "exampleString"
		environmentID := "exampleString"
		collectionID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		listCollectionFieldsPath = strings.Replace(listCollectionFieldsPath, "{environment_id}", environmentID, 1)
		listCollectionFieldsPath = strings.Replace(listCollectionFieldsPath, "{collection_id}", collectionID, 1)
		Context("Successfully - List collection fields", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(listCollectionFieldsPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(listCollectionFieldsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `[]`)
			}))
			It("Succeed to call ListCollectionFields", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				listCollectionFieldsOptions := testService.NewListCollectionFieldsOptions(environmentID, collectionID)
				returnValue, returnValueErr := testService.ListCollectionFields(listCollectionFieldsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListCollectionFieldsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListCollections(options *ListCollectionsOptions)", func() {
		listCollectionsPath := "/v1/environments/{environment_id}/collections"
		version := "exampleString"
		environmentID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		listCollectionsPath = strings.Replace(listCollectionsPath, "{environment_id}", environmentID, 1)
		Context("Successfully - List collections", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(listCollectionsPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(listCollectionsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `[]`)
			}))
			It("Succeed to call ListCollections", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				listCollectionsOptions := testService.NewListCollectionsOptions(environmentID)
				returnValue, returnValueErr := testService.ListCollections(listCollectionsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListCollectionsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("UpdateCollection(options *UpdateCollectionOptions)", func() {
		updateCollectionPath := "/v1/environments/{environment_id}/collections/{collection_id}"
		version := "exampleString"
		environmentID := "exampleString"
		collectionID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		updateCollectionPath = strings.Replace(updateCollectionPath, "{environment_id}", environmentID, 1)
		updateCollectionPath = strings.Replace(updateCollectionPath, "{collection_id}", collectionID, 1)
		Context("Successfully - Update a collection", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(updateCollectionPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(updateCollectionPath))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"collection_id":"xxx"}`)
			}))
			It("Succeed to call UpdateCollection", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				updateCollectionOptions := testService.NewUpdateCollectionOptions(environmentID, collectionID)
				returnValue, returnValueErr := testService.UpdateCollection(updateCollectionOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetUpdateCollectionResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateExpansions(options *CreateExpansionsOptions)", func() {
		createExpansionsPath := "/v1/environments/{environment_id}/collections/{collection_id}/expansions"
		version := "exampleString"
		environmentID := "exampleString"
		collectionID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		createExpansionsPath = strings.Replace(createExpansionsPath, "{environment_id}", environmentID, 1)
		createExpansionsPath = strings.Replace(createExpansionsPath, "{collection_id}", collectionID, 1)
		Context("Successfully - Create or update expansion list", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(createExpansionsPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(createExpansionsPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `[]`)
			}))
			It("Succeed to call CreateExpansions", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				expansions := []discoveryv1.Expansion{
					discoveryv1.Expansion{
						InputTerms: []string{"banana"},
						ExpandedTerms: []string{
							"banana",
							"plantain",
							"fruit",
						},
					},
				}
				createExpansionsOptions := testService.
					NewCreateExpansionsOptions(environmentID, collectionID, expansions)
				returnValue, returnValueErr := testService.CreateExpansions(createExpansionsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetCreateExpansionsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteExpansions(options *DeleteExpansionsOptions)", func() {
		deleteExpansionsPath := "/v1/environments/{environment_id}/collections/{collection_id}/expansions"
		version := "exampleString"
		environmentID := "exampleString"
		collectionID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		deleteExpansionsPath = strings.Replace(deleteExpansionsPath, "{environment_id}", environmentID, 1)
		deleteExpansionsPath = strings.Replace(deleteExpansionsPath, "{collection_id}", collectionID, 1)
		Context("Successfully - Delete the expansion list", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(deleteExpansionsPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(deleteExpansionsPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call DeleteExpansions", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				deleteExpansionsOptions := testService.NewDeleteExpansionsOptions(environmentID, collectionID)
				returnValue, returnValueErr := testService.DeleteExpansions(deleteExpansionsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("ListExpansions(options *ListExpansionsOptions)", func() {
		listExpansionsPath := "/v1/environments/{environment_id}/collections/{collection_id}/expansions"
		version := "exampleString"
		environmentID := "exampleString"
		collectionID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		listExpansionsPath = strings.Replace(listExpansionsPath, "{environment_id}", environmentID, 1)
		listExpansionsPath = strings.Replace(listExpansionsPath, "{collection_id}", collectionID, 1)
		Context("Successfully - Get the expansion list", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(listExpansionsPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(listExpansionsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `[]`)
			}))
			It("Succeed to call ListExpansions", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				listExpansionsOptions := testService.NewListExpansionsOptions(environmentID, collectionID)
				returnValue, returnValueErr := testService.ListExpansions(listExpansionsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListExpansionsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("AddDocument(options *AddDocumentOptions)", func() {
		addDocumentPath := "/v1/environments/{environment_id}/collections/{collection_id}/documents"
		version := "exampleString"
		environmentID := "exampleString"
		collectionID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		addDocumentPath = strings.Replace(addDocumentPath, "{environment_id}", environmentID, 1)
		addDocumentPath = strings.Replace(addDocumentPath, "{collection_id}", collectionID, 1)
		Context("Successfully - Add a document", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(addDocumentPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(addDocumentPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"document_id":"xxx"}`)
			}))
			It("Succeed to call AddDocument", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				addDocumentOptions := testService.NewAddDocumentOptions(environmentID, collectionID).
					SetMetadata("Name:John Smith")
				returnValue, returnValueErr := testService.AddDocument(addDocumentOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetAddDocumentResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteDocument(options *DeleteDocumentOptions)", func() {
		deleteDocumentPath := "/v1/environments/{environment_id}/collections/{collection_id}/documents/{document_id}"
		version := "exampleString"
		environmentID := "exampleString"
		collectionID := "exampleString"
		documentID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		deleteDocumentPath = strings.Replace(deleteDocumentPath, "{environment_id}", environmentID, 1)
		deleteDocumentPath = strings.Replace(deleteDocumentPath, "{collection_id}", collectionID, 1)
		deleteDocumentPath = strings.Replace(deleteDocumentPath, "{document_id}", documentID, 1)
		Context("Successfully - Delete a document", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(deleteDocumentPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(deleteDocumentPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"document_id":"xxx"}`)
			}))
			It("Succeed to call DeleteDocument", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				deleteDocumentOptions := testService.NewDeleteDocumentOptions(environmentID, collectionID, documentID)
				returnValue, returnValueErr := testService.DeleteDocument(deleteDocumentOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetDeleteDocumentResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetDocumentStatus(options *GetDocumentStatusOptions)", func() {
		getDocumentStatusPath := "/v1/environments/{environment_id}/collections/{collection_id}/documents/{document_id}"
		version := "exampleString"
		environmentID := "exampleString"
		collectionID := "exampleString"
		documentID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		getDocumentStatusPath = strings.Replace(getDocumentStatusPath, "{environment_id}", environmentID, 1)
		getDocumentStatusPath = strings.Replace(getDocumentStatusPath, "{collection_id}", collectionID, 1)
		getDocumentStatusPath = strings.Replace(getDocumentStatusPath, "{document_id}", documentID, 1)
		Context("Successfully - Get document details", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(getDocumentStatusPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(getDocumentStatusPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"document_id":"xxx"}`)
			}))
			It("Succeed to call GetDocumentStatus", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				getDocumentStatusOptions := testService.NewGetDocumentStatusOptions(environmentID, collectionID, documentID)
				returnValue, returnValueErr := testService.GetDocumentStatus(getDocumentStatusOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetDocumentStatusResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("UpdateDocument(options *UpdateDocumentOptions)", func() {
		updateDocumentPath := "/v1/environments/{environment_id}/collections/{collection_id}/documents/{document_id}"
		version := "exampleString"
		environmentID := "exampleString"
		collectionID := "exampleString"
		documentID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		updateDocumentPath = strings.Replace(updateDocumentPath, "{environment_id}", environmentID, 1)
		updateDocumentPath = strings.Replace(updateDocumentPath, "{collection_id}", collectionID, 1)
		updateDocumentPath = strings.Replace(updateDocumentPath, "{document_id}", documentID, 1)
		Context("Successfully - Update a document", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(updateDocumentPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(updateDocumentPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
			}))
			It("Succeed to call UpdateDocument", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				updateDocumentOptions := testService.NewUpdateDocumentOptions(environmentID, collectionID, documentID).
					SetMetadata("Name:John Smith")
				returnValue, returnValueErr := testService.UpdateDocument(updateDocumentOptions)
				fmt.Println(returnValueErr)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetUpdateDocumentResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("FederatedQuery(options *FederatedQueryOptions)", func() {
		federatedQueryPath := "/v1/environments/{environment_id}/query"
		version := "exampleString"
		environmentID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		federatedQueryPath = strings.Replace(federatedQueryPath, "{environment_id}", environmentID, 1)
		Context("Successfully - Query documents in multiple collections", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(federatedQueryPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"matching_results":12}`)
			}))
			It("Succeed to call FederatedQuery", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				federatedQueryOptions := testService.NewFederatedQueryOptions(environmentID)
				returnValue, returnValueErr := testService.FederatedQuery(federatedQueryOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetFederatedQueryResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("FederatedQueryNotices(options *FederatedQueryNoticesOptions)", func() {
		federatedQueryNoticesPath := "/v1/environments/{environment_id}/notices"
		version := "exampleString"
		environmentID := "exampleString"
		collectionIds := []string{}
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		federatedQueryNoticesPath = strings.Replace(federatedQueryNoticesPath, "{environment_id}", environmentID, 1)
		Context("Successfully - Query multiple collection system notices", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(federatedQueryNoticesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"matching_results":"23"}`)
			}))
			It("Succeed to call FederatedQueryNotices", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				federatedQueryNoticesOptions := testService.NewFederatedQueryNoticesOptions(environmentID, collectionIds)
				returnValue, returnValueErr := testService.FederatedQueryNotices(federatedQueryNoticesOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetFederatedQueryNoticesResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("Query(options *QueryOptions)", func() {
		queryPath := "/v1/environments/{environment_id}/collections/{collection_id}/query"
		version := "exampleString"
		environmentID := "exampleString"
		collectionID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		queryPath = strings.Replace(queryPath, "{environment_id}", environmentID, 1)
		queryPath = strings.Replace(queryPath, "{collection_id}", collectionID, 1)
		Context("Successfully - Query your collection", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(queryPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(queryPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
			}))
			It("Succeed to call Query", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				queryOptions := testService.NewQueryOptions(environmentID, collectionID)
				returnValue, returnValueErr := testService.Query(queryOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetQueryResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("QueryEntities(options *QueryEntitiesOptions)", func() {
		queryEntitiesPath := "/v1/environments/{environment_id}/collections/{collection_id}/query_entities"
		version := "exampleString"
		environmentID := "exampleString"
		collectionID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		queryEntitiesPath = strings.Replace(queryEntitiesPath, "{environment_id}", environmentID, 1)
		queryEntitiesPath = strings.Replace(queryEntitiesPath, "{collection_id}", collectionID, 1)
		Context("Successfully - Knowledge Graph entity query", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(queryEntitiesPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(queryEntitiesPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `[]`)
			}))
			It("Succeed to call QueryEntities", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				queryEntitiesOptions := testService.NewQueryEntitiesOptions(environmentID, collectionID)
				returnValue, returnValueErr := testService.QueryEntities(queryEntitiesOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetQueryEntitiesResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("QueryNotices(options *QueryNoticesOptions)", func() {
		queryNoticesPath := "/v1/environments/{environment_id}/collections/{collection_id}/notices"
		version := "exampleString"
		environmentID := "exampleString"
		collectionID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		queryNoticesPath = strings.Replace(queryNoticesPath, "{environment_id}", environmentID, 1)
		queryNoticesPath = strings.Replace(queryNoticesPath, "{collection_id}", collectionID, 1)
		Context("Successfully - Query system notices", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(queryNoticesPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(queryNoticesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
			}))
			It("Succeed to call QueryNotices", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				queryNoticesOptions := testService.NewQueryNoticesOptions(environmentID, collectionID)
				returnValue, returnValueErr := testService.QueryNotices(queryNoticesOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetQueryNoticesResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("QueryRelations(options *QueryRelationsOptions)", func() {
		queryRelationsPath := "/v1/environments/{environment_id}/collections/{collection_id}/query_relations"
		version := "exampleString"
		environmentID := "exampleString"
		collectionID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		queryRelationsPath = strings.Replace(queryRelationsPath, "{environment_id}", environmentID, 1)
		queryRelationsPath = strings.Replace(queryRelationsPath, "{collection_id}", collectionID, 1)
		Context("Successfully - Knowledge Graph relationship query", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(queryRelationsPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(queryRelationsPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `[]`)
			}))
			It("Succeed to call QueryRelations", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				queryRelationsOptions := testService.NewQueryRelationsOptions(environmentID, collectionID)
				returnValue, returnValueErr := testService.QueryRelations(queryRelationsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetQueryRelationsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("AddTrainingData(options *AddTrainingDataOptions)", func() {
		addTrainingDataPath := "/v1/environments/{environment_id}/collections/{collection_id}/training_data"
		version := "exampleString"
		environmentID := "exampleString"
		collectionID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		addTrainingDataPath = strings.Replace(addTrainingDataPath, "{environment_id}", environmentID, 1)
		addTrainingDataPath = strings.Replace(addTrainingDataPath, "{collection_id}", collectionID, 1)
		Context("Successfully - Add query to training data", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(addTrainingDataPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(addTrainingDataPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
			}))
			It("Succeed to call AddTrainingData", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				addTrainingDataOptions := testService.NewAddTrainingDataOptions(environmentID, collectionID)
				returnValue, returnValueErr := testService.AddTrainingData(addTrainingDataOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetAddTrainingDataResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateTrainingExample(options *CreateTrainingExampleOptions)", func() {
		createTrainingExamplePath := "/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}/examples"
		version := "exampleString"
		environmentID := "exampleString"
		collectionID := "exampleString"
		queryID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		createTrainingExamplePath = strings.Replace(createTrainingExamplePath, "{environment_id}", environmentID, 1)
		createTrainingExamplePath = strings.Replace(createTrainingExamplePath, "{collection_id}", collectionID, 1)
		createTrainingExamplePath = strings.Replace(createTrainingExamplePath, "{query_id}", queryID, 1)
		Context("Successfully - Add example to training data query", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(createTrainingExamplePath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(createTrainingExamplePath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"document_id":"xxx"}`)
			}))
			It("Succeed to call CreateTrainingExample", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				createTrainingExampleOptions := testService.NewCreateTrainingExampleOptions(environmentID, collectionID, queryID)
				returnValue, returnValueErr := testService.CreateTrainingExample(createTrainingExampleOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetCreateTrainingExampleResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteAllTrainingData(options *DeleteAllTrainingDataOptions)", func() {
		deleteAllTrainingDataPath := "/v1/environments/{environment_id}/collections/{collection_id}/training_data"
		version := "exampleString"
		environmentID := "exampleString"
		collectionID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		deleteAllTrainingDataPath = strings.Replace(deleteAllTrainingDataPath, "{environment_id}", environmentID, 1)
		deleteAllTrainingDataPath = strings.Replace(deleteAllTrainingDataPath, "{collection_id}", collectionID, 1)
		Context("Successfully - Delete all training data", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(deleteAllTrainingDataPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(deleteAllTrainingDataPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call DeleteAllTrainingData", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				deleteAllTrainingDataOptions := testService.NewDeleteAllTrainingDataOptions(environmentID, collectionID)
				returnValue, returnValueErr := testService.DeleteAllTrainingData(deleteAllTrainingDataOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteTrainingData(options *DeleteTrainingDataOptions)", func() {
		deleteTrainingDataPath := "/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}"
		version := "exampleString"
		environmentID := "exampleString"
		collectionID := "exampleString"
		queryID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		deleteTrainingDataPath = strings.Replace(deleteTrainingDataPath, "{environment_id}", environmentID, 1)
		deleteTrainingDataPath = strings.Replace(deleteTrainingDataPath, "{collection_id}", collectionID, 1)
		deleteTrainingDataPath = strings.Replace(deleteTrainingDataPath, "{query_id}", queryID, 1)
		Context("Successfully - Delete a training data query", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(deleteTrainingDataPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(deleteTrainingDataPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call DeleteTrainingData", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				deleteTrainingDataOptions := testService.NewDeleteTrainingDataOptions(environmentID, collectionID, queryID)
				returnValue, returnValueErr := testService.DeleteTrainingData(deleteTrainingDataOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteTrainingExample(options *DeleteTrainingExampleOptions)", func() {
		deleteTrainingExamplePath := "/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}/examples/{example_id}"
		version := "exampleString"
		environmentID := "exampleString"
		collectionID := "exampleString"
		queryID := "exampleString"
		exampleID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		deleteTrainingExamplePath = strings.Replace(deleteTrainingExamplePath, "{environment_id}", environmentID, 1)
		deleteTrainingExamplePath = strings.Replace(deleteTrainingExamplePath, "{collection_id}", collectionID, 1)
		deleteTrainingExamplePath = strings.Replace(deleteTrainingExamplePath, "{query_id}", queryID, 1)
		deleteTrainingExamplePath = strings.Replace(deleteTrainingExamplePath, "{example_id}", exampleID, 1)
		Context("Successfully - Delete example for training data query", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(deleteTrainingExamplePath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(deleteTrainingExamplePath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call DeleteTrainingExample", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				deleteTrainingExampleOptions := testService.NewDeleteTrainingExampleOptions(environmentID, collectionID, queryID, exampleID)
				returnValue, returnValueErr := testService.DeleteTrainingExample(deleteTrainingExampleOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("GetTrainingData(options *GetTrainingDataOptions)", func() {
		getTrainingDataPath := "/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}"
		version := "exampleString"
		environmentID := "exampleString"
		collectionID := "exampleString"
		queryID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		getTrainingDataPath = strings.Replace(getTrainingDataPath, "{environment_id}", environmentID, 1)
		getTrainingDataPath = strings.Replace(getTrainingDataPath, "{collection_id}", collectionID, 1)
		getTrainingDataPath = strings.Replace(getTrainingDataPath, "{query_id}", queryID, 1)
		Context("Successfully - Get details about a query", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(getTrainingDataPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(getTrainingDataPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"query_id":"xxx"}`)
			}))
			It("Succeed to call GetTrainingData", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				getTrainingDataOptions := testService.NewGetTrainingDataOptions(environmentID, collectionID, queryID)
				returnValue, returnValueErr := testService.GetTrainingData(getTrainingDataOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetTrainingDataResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetTrainingExample(options *GetTrainingExampleOptions)", func() {
		getTrainingExamplePath := "/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}/examples/{example_id}"
		version := "exampleString"
		environmentID := "exampleString"
		collectionID := "exampleString"
		queryID := "exampleString"
		exampleID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		getTrainingExamplePath = strings.Replace(getTrainingExamplePath, "{environment_id}", environmentID, 1)
		getTrainingExamplePath = strings.Replace(getTrainingExamplePath, "{collection_id}", collectionID, 1)
		getTrainingExamplePath = strings.Replace(getTrainingExamplePath, "{query_id}", queryID, 1)
		getTrainingExamplePath = strings.Replace(getTrainingExamplePath, "{example_id}", exampleID, 1)
		Context("Successfully - Get details for training data example", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(getTrainingExamplePath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(getTrainingExamplePath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
			}))
			It("Succeed to call GetTrainingExample", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				getTrainingExampleOptions := testService.NewGetTrainingExampleOptions(environmentID, collectionID, queryID, exampleID)
				returnValue, returnValueErr := testService.GetTrainingExample(getTrainingExampleOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetTrainingExampleResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListTrainingData(options *ListTrainingDataOptions)", func() {
		listTrainingDataPath := "/v1/environments/{environment_id}/collections/{collection_id}/training_data"
		version := "exampleString"
		environmentID := "exampleString"
		collectionID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		listTrainingDataPath = strings.Replace(listTrainingDataPath, "{environment_id}", environmentID, 1)
		listTrainingDataPath = strings.Replace(listTrainingDataPath, "{collection_id}", collectionID, 1)
		Context("Successfully - List training data", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(listTrainingDataPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(listTrainingDataPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"environment_id":"xxx"}`)
			}))
			It("Succeed to call ListTrainingData", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				listTrainingDataOptions := testService.NewListTrainingDataOptions(environmentID, collectionID)
				returnValue, returnValueErr := testService.ListTrainingData(listTrainingDataOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListTrainingDataResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListTrainingExamples(options *ListTrainingExamplesOptions)", func() {
		listTrainingExamplesPath := "/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}/examples"
		version := "exampleString"
		environmentID := "exampleString"
		collectionID := "exampleString"
		queryID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		listTrainingExamplesPath = strings.Replace(listTrainingExamplesPath, "{environment_id}", environmentID, 1)
		listTrainingExamplesPath = strings.Replace(listTrainingExamplesPath, "{collection_id}", collectionID, 1)
		listTrainingExamplesPath = strings.Replace(listTrainingExamplesPath, "{query_id}", queryID, 1)
		Context("Successfully - List examples for a training data query", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(listTrainingExamplesPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(listTrainingExamplesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `[]`)
			}))
			It("Succeed to call ListTrainingExamples", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				listTrainingExamplesOptions := testService.NewListTrainingExamplesOptions(environmentID, collectionID, queryID)
				returnValue, returnValueErr := testService.ListTrainingExamples(listTrainingExamplesOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListTrainingExamplesResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("UpdateTrainingExample(options *UpdateTrainingExampleOptions)", func() {
		updateTrainingExamplePath := "/v1/environments/{environment_id}/collections/{collection_id}/training_data/{query_id}/examples/{example_id}"
		version := "exampleString"
		environmentID := "exampleString"
		collectionID := "exampleString"
		queryID := "exampleString"
		exampleID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		updateTrainingExamplePath = strings.Replace(updateTrainingExamplePath, "{environment_id}", environmentID, 1)
		updateTrainingExamplePath = strings.Replace(updateTrainingExamplePath, "{collection_id}", collectionID, 1)
		updateTrainingExamplePath = strings.Replace(updateTrainingExamplePath, "{query_id}", queryID, 1)
		updateTrainingExamplePath = strings.Replace(updateTrainingExamplePath, "{example_id}", exampleID, 1)
		Context("Successfully - Change label or cross reference for example", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(updateTrainingExamplePath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(updateTrainingExamplePath))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
			}))
			It("Succeed to call UpdateTrainingExample", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				updateTrainingExampleOptions := testService.NewUpdateTrainingExampleOptions(environmentID, collectionID, queryID, exampleID)
				returnValue, returnValueErr := testService.UpdateTrainingExample(updateTrainingExampleOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetUpdateTrainingExampleResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteUserData(options *DeleteUserDataOptions)", func() {
		deleteUserDataPath := "/v1/user_data"
		version := "exampleString"
		customerID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Delete labeled data", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(deleteUserDataPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call DeleteUserData", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				deleteUserDataOptions := testService.NewDeleteUserDataOptions(customerID)
				returnValue, returnValueErr := testService.DeleteUserData(deleteUserDataOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("CreateEvent(options *CreateEventOptions)", func() {
		createEventPath := "/v1/events"
		version := "exampleString"
		typeVar := "exampleString"
		data := &discoveryv1.EventData{
			EnvironmentID: core.StringPtr("xxx"),
			SessionToken:  core.StringPtr("yyy"),
			CollectionID:  core.StringPtr("zzz"),
			DocumentID:    core.StringPtr("yyy"),
		}
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Create event", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(createEventPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(createEventPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"type":"xxx"}`)
			}))
			It("Succeed to call CreateEvent", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				createEventOptions := testService.NewCreateEventOptions(typeVar, data)
				returnValue, returnValueErr := testService.CreateEvent(createEventOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetCreateEventResult(returnValue)
				fmt.Println(result)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetMetricsEventRate(options *GetMetricsEventRateOptions)", func() {
		getMetricsEventRatePath := "/v1/metrics/event_rate"
		version := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Percentage of queries with an associated event", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(getMetricsEventRatePath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(getMetricsEventRatePath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `[]`)
			}))
			It("Succeed to call GetMetricsEventRate", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				getMetricsEventRateOptions := testService.NewGetMetricsEventRateOptions()
				returnValue, returnValueErr := testService.GetMetricsEventRate(getMetricsEventRateOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetMetricsEventRateResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetMetricsQuery(options *GetMetricsQueryOptions)", func() {
		getMetricsQueryPath := "/v1/metrics/number_of_queries"
		version := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Number of queries over time", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(getMetricsQueryPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(getMetricsQueryPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `[]`)
			}))
			It("Succeed to call GetMetricsQuery", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				getMetricsQueryOptions := testService.NewGetMetricsQueryOptions()
				returnValue, returnValueErr := testService.GetMetricsQuery(getMetricsQueryOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetMetricsQueryResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetMetricsQueryEvent(options *GetMetricsQueryEventOptions)", func() {
		getMetricsQueryEventPath := "/v1/metrics/number_of_queries_with_event"
		version := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Number of queries with an event over time", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(getMetricsQueryEventPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(getMetricsQueryEventPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `[]`)
			}))
			It("Succeed to call GetMetricsQueryEvent", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				getMetricsQueryEventOptions := testService.NewGetMetricsQueryEventOptions()
				returnValue, returnValueErr := testService.GetMetricsQueryEvent(getMetricsQueryEventOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetMetricsQueryEventResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetMetricsQueryNoResults(options *GetMetricsQueryNoResultsOptions)", func() {
		getMetricsQueryNoResultsPath := "/v1/metrics/number_of_queries_with_no_search_results"
		version := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Number of queries with no search results over time", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(getMetricsQueryNoResultsPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(getMetricsQueryNoResultsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `[]`)
			}))
			It("Succeed to call GetMetricsQueryNoResults", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				getMetricsQueryNoResultsOptions := testService.NewGetMetricsQueryNoResultsOptions()
				returnValue, returnValueErr := testService.GetMetricsQueryNoResults(getMetricsQueryNoResultsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetMetricsQueryNoResultsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetMetricsQueryTokenEvent(options *GetMetricsQueryTokenEventOptions)", func() {
		getMetricsQueryTokenEventPath := "/v1/metrics/top_query_tokens_with_event_rate"
		version := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Most frequent query tokens with an event", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(getMetricsQueryTokenEventPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(getMetricsQueryTokenEventPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `[]`)
			}))
			It("Succeed to call GetMetricsQueryTokenEvent", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				getMetricsQueryTokenEventOptions := testService.NewGetMetricsQueryTokenEventOptions()
				returnValue, returnValueErr := testService.GetMetricsQueryTokenEvent(getMetricsQueryTokenEventOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetMetricsQueryTokenEventResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("QueryLog(options *QueryLogOptions)", func() {
		queryLogPath := "/v1/logs"
		version := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Search the query and event log", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(queryLogPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(queryLogPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
			}))
			It("Succeed to call QueryLog", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				queryLogOptions := testService.NewQueryLogOptions()
				returnValue, returnValueErr := testService.QueryLog(queryLogOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetQueryLogResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateCredentials(options *CreateCredentialsOptions)", func() {
		createCredentialsPath := "/v1/environments/{environment_id}/credentials"
		version := "exampleString"
		environmentID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		createCredentialsPath = strings.Replace(createCredentialsPath, "{environment_id}", environmentID, 1)
		Context("Successfully - Create credentials", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(createCredentialsPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(createCredentialsPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
			}))
			It("Succeed to call CreateCredentials", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				createCredentialsOptions := testService.NewCreateCredentialsOptions(environmentID)
				returnValue, returnValueErr := testService.CreateCredentials(createCredentialsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetCreateCredentialsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteCredentials(options *DeleteCredentialsOptions)", func() {
		deleteCredentialsPath := "/v1/environments/{environment_id}/credentials/{credential_id}"
		version := "exampleString"
		environmentID := "exampleString"
		credentialID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		deleteCredentialsPath = strings.Replace(deleteCredentialsPath, "{environment_id}", environmentID, 1)
		deleteCredentialsPath = strings.Replace(deleteCredentialsPath, "{credential_id}", credentialID, 1)
		Context("Successfully - Delete credentials", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(deleteCredentialsPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(deleteCredentialsPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"credential_id":"xxxx"}`)
			}))
			It("Succeed to call DeleteCredentials", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				deleteCredentialsOptions := testService.NewDeleteCredentialsOptions(environmentID, credentialID)
				returnValue, returnValueErr := testService.DeleteCredentials(deleteCredentialsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetDeleteCredentialsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetCredentials(options *GetCredentialsOptions)", func() {
		getCredentialsPath := "/v1/environments/{environment_id}/credentials/{credential_id}"
		version := "exampleString"
		environmentID := "exampleString"
		credentialID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		getCredentialsPath = strings.Replace(getCredentialsPath, "{environment_id}", environmentID, 1)
		getCredentialsPath = strings.Replace(getCredentialsPath, "{credential_id}", credentialID, 1)
		Context("Successfully - View Credentials", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(getCredentialsPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(getCredentialsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"credential_id":"xxx"}`)
			}))
			It("Succeed to call GetCredentials", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				getCredentialsOptions := testService.NewGetCredentialsOptions(environmentID, credentialID)
				returnValue, returnValueErr := testService.GetCredentials(getCredentialsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetCredentialsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListCredentials(options *ListCredentialsOptions)", func() {
		listCredentialsPath := "/v1/environments/{environment_id}/credentials"
		version := "exampleString"
		environmentID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		listCredentialsPath = strings.Replace(listCredentialsPath, "{environment_id}", environmentID, 1)
		Context("Successfully - List credentials", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(listCredentialsPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(listCredentialsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
			}))
			It("Succeed to call ListCredentials", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				listCredentialsOptions := testService.NewListCredentialsOptions(environmentID)
				returnValue, returnValueErr := testService.ListCredentials(listCredentialsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListCredentialsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("UpdateCredentials(options *UpdateCredentialsOptions)", func() {
		updateCredentialsPath := "/v1/environments/{environment_id}/credentials/{credential_id}"
		version := "exampleString"
		environmentID := "exampleString"
		credentialID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		updateCredentialsPath = strings.Replace(updateCredentialsPath, "{environment_id}", environmentID, 1)
		updateCredentialsPath = strings.Replace(updateCredentialsPath, "{credential_id}", credentialID, 1)
		Context("Successfully - Update credentials", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(updateCredentialsPath + "?version=" + version))
				Expect(req.URL.Path).To(Equal(updateCredentialsPath))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
			}))
			It("Succeed to call UpdateCredentials", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				updateCredentialsOptions := testService.NewUpdateCredentialsOptions(environmentID, credentialID)
				returnValue, returnValueErr := testService.UpdateCredentials(updateCredentialsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetUpdateCredentialsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateTokenizationDictionary(createTokenizationDictionaryOptions *CreateTokenizationDictionaryOptions)", func() {
		createTokenizationDictionaryPath := "/v1/environments/{environment_id}/collections/{collection_id}/word_lists/tokenization_dictionary"
		version := "exampleString"
		environmentID := "exampleString"
		collectionID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(createTokenizationDictionaryPath, "{environment_id}", environmentID, 1)
		Path = strings.Replace(Path, "{collection_id}", collectionID, 1)
		Context("Successfully - Create tokenization dictionary", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"status": "ready"}`)
			}))
			It("Succeed to call CreateTokenizationDictionary", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
		environmentID := "exampleString"
		collectionID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(deleteTokenizationDictionaryPath, "{environment_id}", environmentID, 1)
		Path = strings.Replace(Path, "{collection_id}", collectionID, 1)
		Context("Successfully - Delete tokenization dictionary", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
			}))
			It("Succeed to call DeleteTokenizationDictionary", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("GetTokenizationDictionaryStatus(getTokenizationDictionaryStatusOptions *GetTokenizationDictionaryStatusOptions)", func() {
		getTokenizationDictionaryStatusPath := "/v1/environments/{environment_id}/collections/{collection_id}/word_lists/tokenization_dictionary"
		version := "exampleString"
		environmentID := "exampleString"
		collectionID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(getTokenizationDictionaryStatusPath, "{environment_id}", environmentID, 1)
		Path = strings.Replace(Path, "{collection_id}", collectionID, 1)
		Context("Successfully - Get tokenization dictionary status", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"status": "ready"}`)
			}))
			It("Succeed to call GetTokenizationDictionaryStatus", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("CreateStopwordList(createStopwordListOptions *CreateStopwordListOptions)", func() {
		createStopwordListPath := "/v1/environments/{environment_id}/collections/{collection_id}/word_lists/stopwords"
		version := "exampleString"
		environmentID := "exampleString"
		collectionID := "exampleString"
		pwd, _ := os.Getwd()
		stopwordsFile, stopwordsFileErr := os.Open(pwd + "/../resources/stopwords.txt")
		if stopwordsFileErr != nil {
			panic(stopwordsFileErr)
		}
		stopwordFilename := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(createStopwordListPath, "{environment_id}", environmentID, 1)
		Path = strings.Replace(Path, "{collection_id}", collectionID, 1)
		Context("Successfully - Create stopword list", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"status": "ready"}`)
			}))
			It("Succeed to call CreateStopwordList", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
		environmentID := "exampleString"
		collectionID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(deleteStopwordListPath, "{environment_id}", environmentID, 1)
		Path = strings.Replace(Path, "{collection_id}", collectionID, 1)
		Context("Successfully - Delete a custom stopword list", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
			}))
			It("Succeed to call DeleteStopwordList", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("CreateGateway(createGatewayOptions *CreateGatewayOptions)", func() {
		createGatewayPath := "/v1/environments/{environment_id}/gateways"
		version := "exampleString"
		environmentID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(createGatewayPath, "{environment_id}", environmentID, 1)
		Context("Successfully - Create Gateway", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"gateway_id":"1234"}`)
			}))
			It("Succeed to call CreateGateway", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("DeleteGateway(deleteGatewayOptions *DeleteGatewayOptions)", func() {
		deleteGatewayPath := "/v1/environments/{environment_id}/gateways/{gateway_id}"
		version := "exampleString"
		environmentID := "exampleString"
		gatewayID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(deleteGatewayPath, "{environment_id}", environmentID, 1)
		Path = strings.Replace(Path, "{gateway_id}", gatewayID, 1)
		Context("Successfully - Delete Gateway", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"gateway_id":"1234"}`)
			}))
			It("Succeed to call DeleteGateway", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("GetGateway(getGatewayOptions *GetGatewayOptions)", func() {
		getGatewayPath := "/v1/environments/{environment_id}/gateways/{gateway_id}"
		version := "exampleString"
		environmentID := "exampleString"
		gatewayID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(getGatewayPath, "{environment_id}", environmentID, 1)
		Path = strings.Replace(Path, "{gateway_id}", gatewayID, 1)
		Context("Successfully - Get Gateway", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"gateway_id":"1234"}`)
			}))
			It("Succeed to call GetGateway", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("ListGateway(listGatewayOptions *ListGatewayOptions)", func() {
		listGatewayPath := "/v1/environments/{environment_id}/gateways"
		version := "exampleString"
		environmentID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(listGatewayPath, "{environment_id}", environmentID, 1)
		Context("Successfully - List Gateways", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `["gateways": {"gateway_id":"1234"}]`)
			}))
			It("Succeed to call ListGateway", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
	Describe("GetStopwordListStatus(getStopwordListStatusOptions *GetStopwordListStatusOptions)", func() {
		getStopwordListStatusPath := "/v1/environments/{environment_id}/collections/{collection_id}/word_lists/stopwords"
		version := "exampleString"
		environmentID := "exampleString"
		collectionID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(getStopwordListStatusPath, "{environment_id}", environmentID, 1)
		Path = strings.Replace(Path, "{collection_id}", collectionID, 1)
		Context("Successfully - Get stopword list status", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"status": "ready"}`)
			}))
			It("Succeed to call GetStopwordListStatus", func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:      testServer.URL,
					Version:  version,
					Username: username,
					Password: password,
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
})
