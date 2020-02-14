/**
 * (C) Copyright IBM Corp. 2019, 2020.
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

package discoveryv2_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	"github.com/IBM/go-sdk-core/v3/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/discoveryv2"
)

var _ = Describe(`DiscoveryV2`, func() {
	Describe(`ListCollections(listCollectionsOptions *ListCollectionsOptions)`, func() {
		listCollectionsPath := "/v2/projects/{project_id}/collections"
		version := "exampleString"
		bearerToken := "0ui9876453"
		projectID := "exampleString"
		listCollectionsPath = strings.Replace(listCollectionsPath, "{project_id}", projectID, 1)
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

				testService, testServiceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
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

				listCollectionsOptions := testService.NewListCollectionsOptions(projectID)
				result, response, operationErr = testService.ListCollections(listCollectionsOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`Query(queryOptions *QueryOptions)`, func() {
		queryPath := "/v2/projects/{project_id}/query"
		version := "exampleString"
		bearerToken := "0ui9876453"
		projectID := "exampleString"
		queryPath = strings.Replace(queryPath, "{project_id}", projectID, 1)
		Context(`Successfully - Query a project`, func() {
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

				testService, testServiceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
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

				queryOptions := testService.NewQueryOptions(projectID)
				result, response, operationErr = testService.Query(queryOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetAutocompletion(getAutocompletionOptions *GetAutocompletionOptions)`, func() {
		getAutocompletionPath := "/v2/projects/{project_id}/autocompletion"
		version := "exampleString"
		bearerToken := "0ui9876453"
		projectID := "exampleString"
		prefix := "exampleString"
		getAutocompletionPath = strings.Replace(getAutocompletionPath, "{project_id}", projectID, 1)
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

				testService, testServiceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
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

				getAutocompletionOptions := testService.NewGetAutocompletionOptions(projectID, prefix)
				result, response, operationErr = testService.GetAutocompletion(getAutocompletionOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`QueryNotices(queryNoticesOptions *QueryNoticesOptions)`, func() {
		queryNoticesPath := "/v2/projects/{project_id}/notices"
		version := "exampleString"
		bearerToken := "0ui9876453"
		projectID := "exampleString"
		queryNoticesPath = strings.Replace(queryNoticesPath, "{project_id}", projectID, 1)
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

				testService, testServiceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
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

				queryNoticesOptions := testService.NewQueryNoticesOptions(projectID)
				result, response, operationErr = testService.QueryNotices(queryNoticesOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListFields(listFieldsOptions *ListFieldsOptions)`, func() {
		listFieldsPath := "/v2/projects/{project_id}/fields"
		version := "exampleString"
		bearerToken := "0ui9876453"
		projectID := "exampleString"
		listFieldsPath = strings.Replace(listFieldsPath, "{project_id}", projectID, 1)
		Context(`Successfully - List fields`, func() {
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

				testService, testServiceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
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

				listFieldsOptions := testService.NewListFieldsOptions(projectID)
				result, response, operationErr = testService.ListFields(listFieldsOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetComponentSettings(getComponentSettingsOptions *GetComponentSettingsOptions)`, func() {
		getComponentSettingsPath := "/v2/projects/{project_id}/component_settings"
		version := "exampleString"
		bearerToken := "0ui9876453"
		projectID := "exampleString"
		getComponentSettingsPath = strings.Replace(getComponentSettingsPath, "{project_id}", projectID, 1)
		Context(`Successfully - Configuration settings for components`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getComponentSettingsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call GetComponentSettings`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetComponentSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getComponentSettingsOptions := testService.NewGetComponentSettingsOptions(projectID)
				result, response, operationErr = testService.GetComponentSettings(getComponentSettingsOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`AddDocument(addDocumentOptions *AddDocumentOptions)`, func() {
		addDocumentPath := "/v2/projects/{project_id}/collections/{collection_id}/documents"
		version := "exampleString"
		bearerToken := "0ui9876453"
		projectID := "exampleString"
		collectionID := "exampleString"
		pwd, _ := os.Getwd()
		file, fileErr := os.Open(pwd + "/../resources/simple.html")
		if fileErr != nil {
			panic(fileErr)
		}
		addDocumentPath = strings.Replace(addDocumentPath, "{project_id}", projectID, 1)
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

				testService, testServiceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
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

				addDocumentOptions := testService.NewAddDocumentOptions(projectID, collectionID)
				addDocumentOptions.File = file
				result, response, operationErr = testService.AddDocument(addDocumentOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateDocument(updateDocumentOptions *UpdateDocumentOptions)`, func() {
		updateDocumentPath := "/v2/projects/{project_id}/collections/{collection_id}/documents/{document_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		projectID := "exampleString"
		collectionID := "exampleString"
		documentID := "exampleString"
		pwd, _ := os.Getwd()
		file, fileErr := os.Open(pwd + "/../resources/simple.html")
		if fileErr != nil {
			panic(fileErr)
		}
		updateDocumentPath = strings.Replace(updateDocumentPath, "{project_id}", projectID, 1)
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

				testService, testServiceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
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

				updateDocumentOptions := testService.NewUpdateDocumentOptions(projectID, collectionID, documentID)
				updateDocumentOptions.File = file
				result, response, operationErr = testService.UpdateDocument(updateDocumentOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteDocument(deleteDocumentOptions *DeleteDocumentOptions)`, func() {
		deleteDocumentPath := "/v2/projects/{project_id}/collections/{collection_id}/documents/{document_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		projectID := "exampleString"
		collectionID := "exampleString"
		documentID := "exampleString"
		deleteDocumentPath = strings.Replace(deleteDocumentPath, "{project_id}", projectID, 1)
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

				testService, testServiceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
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

				deleteDocumentOptions := testService.NewDeleteDocumentOptions(projectID, collectionID, documentID)
				result, response, operationErr = testService.DeleteDocument(deleteDocumentOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListTrainingQueries(listTrainingQueriesOptions *ListTrainingQueriesOptions)`, func() {
		listTrainingQueriesPath := "/v2/projects/{project_id}/training_data/queries"
		version := "exampleString"
		bearerToken := "0ui9876453"
		projectID := "exampleString"
		listTrainingQueriesPath = strings.Replace(listTrainingQueriesPath, "{project_id}", projectID, 1)
		Context(`Successfully - List training queries`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listTrainingQueriesPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call ListTrainingQueries`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListTrainingQueries(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listTrainingQueriesOptions := testService.NewListTrainingQueriesOptions(projectID)
				result, response, operationErr = testService.ListTrainingQueries(listTrainingQueriesOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteTrainingQueries(deleteTrainingQueriesOptions *DeleteTrainingQueriesOptions)`, func() {
		deleteTrainingQueriesPath := "/v2/projects/{project_id}/training_data/queries"
		version := "exampleString"
		bearerToken := "0ui9876453"
		projectID := "exampleString"
		deleteTrainingQueriesPath = strings.Replace(deleteTrainingQueriesPath, "{project_id}", projectID, 1)
		Context(`Successfully - Delete training queries`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteTrainingQueriesPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(204)
			}))
			It(`Succeed to call DeleteTrainingQueries`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.DeleteTrainingQueries(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				deleteTrainingQueriesOptions := testService.NewDeleteTrainingQueriesOptions(projectID)
				response, operationErr = testService.DeleteTrainingQueries(deleteTrainingQueriesOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateTrainingQuery(createTrainingQueryOptions *CreateTrainingQueryOptions)`, func() {
		createTrainingQueryPath := "/v2/projects/{project_id}/training_data/queries"
		version := "exampleString"
		bearerToken := "0ui9876453"
		projectID := "exampleString"
		naturalLanguageQuery := "exampleString"
		examples := []discoveryv2.TrainingExample{}
		createTrainingQueryPath = strings.Replace(createTrainingQueryPath, "{project_id}", projectID, 1)
		Context(`Successfully - Create training query`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createTrainingQueryPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"natural_language_query": "fake_NaturalLanguageQuery", "examples": []}`)
			}))
			It(`Succeed to call CreateTrainingQuery`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.CreateTrainingQuery(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				createTrainingQueryOptions := testService.NewCreateTrainingQueryOptions(projectID, naturalLanguageQuery, examples)
				result, response, operationErr = testService.CreateTrainingQuery(createTrainingQueryOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetTrainingQuery(getTrainingQueryOptions *GetTrainingQueryOptions)`, func() {
		getTrainingQueryPath := "/v2/projects/{project_id}/training_data/queries/{query_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		projectID := "exampleString"
		queryID := "exampleString"
		getTrainingQueryPath = strings.Replace(getTrainingQueryPath, "{project_id}", projectID, 1)
		getTrainingQueryPath = strings.Replace(getTrainingQueryPath, "{query_id}", queryID, 1)
		Context(`Successfully - Get a training data query`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getTrainingQueryPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"natural_language_query": "fake_NaturalLanguageQuery", "examples": []}`)
			}))
			It(`Succeed to call GetTrainingQuery`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetTrainingQuery(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getTrainingQueryOptions := testService.NewGetTrainingQueryOptions(projectID, queryID)
				result, response, operationErr = testService.GetTrainingQuery(getTrainingQueryOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateTrainingQuery(updateTrainingQueryOptions *UpdateTrainingQueryOptions)`, func() {
		updateTrainingQueryPath := "/v2/projects/{project_id}/training_data/queries/{query_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		projectID := "exampleString"
		queryID := "exampleString"
		naturalLanguageQuery := "exampleString"
		examples := []discoveryv2.TrainingExample{}
		updateTrainingQueryPath = strings.Replace(updateTrainingQueryPath, "{project_id}", projectID, 1)
		updateTrainingQueryPath = strings.Replace(updateTrainingQueryPath, "{query_id}", queryID, 1)
		Context(`Successfully - Update a training query`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateTrainingQueryPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"natural_language_query": "fake_NaturalLanguageQuery", "examples": []}`)
			}))
			It(`Succeed to call UpdateTrainingQuery`, func() {
				defer testServer.Close()

				testService, testServiceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.UpdateTrainingQuery(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				updateTrainingQueryOptions := testService.NewUpdateTrainingQueryOptions(projectID, queryID, naturalLanguageQuery, examples)
				result, response, operationErr = testService.UpdateTrainingQuery(updateTrainingQueryOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("Model constructor tests", func() {
		Context("with a sample service", func() {
			version := "1970-01-01"
			testService, _ := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
				URL:           "http://discoveryv2modelgenerator.com",
				Version:       version,
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It("should call NewTrainingExample successfully", func() {
				documentID := "exampleString"
				collectionID := "exampleString"
				relevance := int64(1234)
				model, err := testService.NewTrainingExample(documentID, collectionID, relevance)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It("should call NewTrainingQuery successfully", func() {
				naturalLanguageQuery := "exampleString"
				examples := []discoveryv2.TrainingExample{}
				model, err := testService.NewTrainingQuery(naturalLanguageQuery, examples)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
})
