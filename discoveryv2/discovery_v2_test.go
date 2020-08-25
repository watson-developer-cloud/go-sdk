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

	"github.com/IBM/go-sdk-core/core"
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
	Describe(`CreateCollection(createCollectionOptions *CreateCollectionOptions)`, func() {
		createCollectionPath := "/v2/projects/{project_id}/collections"
		version := "exampleString"
		bearerToken := "0ui9876453"
		projectID := "exampleString"
		name := "exampleString"
		createCollectionPath = strings.Replace(createCollectionPath, "{project_id}", projectID, 1)
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
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"name": "fake_Name"}`)
			}))
			It(`Succeed to call CreateCollection`, func() {
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
				result, response, operationErr := testService.CreateCollection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				createCollectionOptions := testService.NewCreateCollectionOptions(projectID, name)
				result, response, operationErr = testService.CreateCollection(createCollectionOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetCollection(getCollectionOptions *GetCollectionOptions)`, func() {
		getCollectionPath := "/v2/projects/{project_id}/collections/{collection_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		projectID := "exampleString"
		collectionID := "exampleString"
		getCollectionPath = strings.Replace(getCollectionPath, "{project_id}", projectID, 1)
		getCollectionPath = strings.Replace(getCollectionPath, "{collection_id}", collectionID, 1)
		Context(`Successfully - Get collection`, func() {
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
				fmt.Fprintf(res, `{"name": "fake_Name"}`)
			}))
			It(`Succeed to call GetCollection`, func() {
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
				result, response, operationErr := testService.GetCollection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getCollectionOptions := testService.NewGetCollectionOptions(projectID, collectionID)
				result, response, operationErr = testService.GetCollection(getCollectionOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateCollection(updateCollectionOptions *UpdateCollectionOptions)`, func() {
		updateCollectionPath := "/v2/projects/{project_id}/collections/{collection_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		projectID := "exampleString"
		collectionID := "exampleString"
		updateCollectionPath = strings.Replace(updateCollectionPath, "{project_id}", projectID, 1)
		updateCollectionPath = strings.Replace(updateCollectionPath, "{collection_id}", collectionID, 1)
		Context(`Successfully - Update a collection`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateCollectionPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"name": "fake_Name"}`)
			}))
			It(`Succeed to call UpdateCollection`, func() {
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
				result, response, operationErr := testService.UpdateCollection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				updateCollectionOptions := testService.NewUpdateCollectionOptions(projectID, collectionID)
				result, response, operationErr = testService.UpdateCollection(updateCollectionOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteCollection(deleteCollectionOptions *DeleteCollectionOptions)`, func() {
		deleteCollectionPath := "/v2/projects/{project_id}/collections/{collection_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		projectID := "exampleString"
		collectionID := "exampleString"
		deleteCollectionPath = strings.Replace(deleteCollectionPath, "{project_id}", projectID, 1)
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
				res.WriteHeader(204)
			}))
			It(`Succeed to call DeleteCollection`, func() {
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
				response, operationErr := testService.DeleteCollection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				deleteCollectionOptions := testService.NewDeleteCollectionOptions(projectID, collectionID)
				response, operationErr = testService.DeleteCollection(deleteCollectionOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
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
		Context(`Successfully - List component settings`, func() {
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
	Describe(`ListEnrichments(listEnrichmentsOptions *ListEnrichmentsOptions)`, func() {
		listEnrichmentsPath := "/v2/projects/{project_id}/enrichments"
		version := "exampleString"
		bearerToken := "0ui9876453"
		projectID := "exampleString"
		listEnrichmentsPath = strings.Replace(listEnrichmentsPath, "{project_id}", projectID, 1)
		Context(`Successfully - List Enrichments`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listEnrichmentsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call ListEnrichments`, func() {
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
				result, response, operationErr := testService.ListEnrichments(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listEnrichmentsOptions := testService.NewListEnrichmentsOptions(projectID)
				result, response, operationErr = testService.ListEnrichments(listEnrichmentsOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateEnrichment(createEnrichmentOptions *CreateEnrichmentOptions)`, func() {
		createEnrichmentPath := "/v2/projects/{project_id}/enrichments"
		version := "exampleString"
		bearerToken := "0ui9876453"
		projectID := "exampleString"
		enrichment := new(discoveryv2.CreateEnrichment)
		createEnrichmentPath = strings.Replace(createEnrichmentPath, "{project_id}", projectID, 1)
		Context(`Successfully - Create an enrichment`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createEnrichmentPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call CreateEnrichment`, func() {
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
				result, response, operationErr := testService.CreateEnrichment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				createEnrichmentOptions := testService.NewCreateEnrichmentOptions(projectID, enrichment)
				result, response, operationErr = testService.CreateEnrichment(createEnrichmentOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetEnrichment(getEnrichmentOptions *GetEnrichmentOptions)`, func() {
		getEnrichmentPath := "/v2/projects/{project_id}/enrichments/{enrichment_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		projectID := "exampleString"
		enrichmentID := "exampleString"
		getEnrichmentPath = strings.Replace(getEnrichmentPath, "{project_id}", projectID, 1)
		getEnrichmentPath = strings.Replace(getEnrichmentPath, "{enrichment_id}", enrichmentID, 1)
		Context(`Successfully - Get enrichment`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getEnrichmentPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call GetEnrichment`, func() {
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
				result, response, operationErr := testService.GetEnrichment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getEnrichmentOptions := testService.NewGetEnrichmentOptions(projectID, enrichmentID)
				result, response, operationErr = testService.GetEnrichment(getEnrichmentOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateEnrichment(updateEnrichmentOptions *UpdateEnrichmentOptions)`, func() {
		updateEnrichmentPath := "/v2/projects/{project_id}/enrichments/{enrichment_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		projectID := "exampleString"
		enrichmentID := "exampleString"
		name := "exampleString"
		updateEnrichmentPath = strings.Replace(updateEnrichmentPath, "{project_id}", projectID, 1)
		updateEnrichmentPath = strings.Replace(updateEnrichmentPath, "{enrichment_id}", enrichmentID, 1)
		Context(`Successfully - Update an enrichment`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateEnrichmentPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call UpdateEnrichment`, func() {
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
				result, response, operationErr := testService.UpdateEnrichment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				updateEnrichmentOptions := testService.NewUpdateEnrichmentOptions(projectID, enrichmentID, name)
				result, response, operationErr = testService.UpdateEnrichment(updateEnrichmentOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteEnrichment(deleteEnrichmentOptions *DeleteEnrichmentOptions)`, func() {
		deleteEnrichmentPath := "/v2/projects/{project_id}/enrichments/{enrichment_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		projectID := "exampleString"
		enrichmentID := "exampleString"
		deleteEnrichmentPath = strings.Replace(deleteEnrichmentPath, "{project_id}", projectID, 1)
		deleteEnrichmentPath = strings.Replace(deleteEnrichmentPath, "{enrichment_id}", enrichmentID, 1)
		Context(`Successfully - Delete an enrichment`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteEnrichmentPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(204)
			}))
			It(`Succeed to call DeleteEnrichment`, func() {
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
				response, operationErr := testService.DeleteEnrichment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				deleteEnrichmentOptions := testService.NewDeleteEnrichmentOptions(projectID, enrichmentID)
				response, operationErr = testService.DeleteEnrichment(deleteEnrichmentOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`ListProjects(listProjectsOptions *ListProjectsOptions)`, func() {
		listProjectsPath := "/v2/projects"
		version := "exampleString"
		bearerToken := "0ui9876453"
		Context(`Successfully - List projects`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listProjectsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call ListProjects`, func() {
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
				result, response, operationErr := testService.ListProjects(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listProjectsOptions := testService.NewListProjectsOptions()
				result, response, operationErr = testService.ListProjects(listProjectsOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateProject(createProjectOptions *CreateProjectOptions)`, func() {
		createProjectPath := "/v2/projects"
		version := "exampleString"
		bearerToken := "0ui9876453"
		name := "exampleString"
		typeVar := "exampleString"
		Context(`Successfully - Create a Project`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createProjectPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call CreateProject`, func() {
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
				result, response, operationErr := testService.CreateProject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				createProjectOptions := testService.NewCreateProjectOptions(name, typeVar)
				result, response, operationErr = testService.CreateProject(createProjectOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetProject(getProjectOptions *GetProjectOptions)`, func() {
		getProjectPath := "/v2/projects/{project_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		projectID := "exampleString"
		getProjectPath = strings.Replace(getProjectPath, "{project_id}", projectID, 1)
		Context(`Successfully - Get project`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getProjectPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call GetProject`, func() {
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
				result, response, operationErr := testService.GetProject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getProjectOptions := testService.NewGetProjectOptions(projectID)
				result, response, operationErr = testService.GetProject(getProjectOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateProject(updateProjectOptions *UpdateProjectOptions)`, func() {
		updateProjectPath := "/v2/projects/{project_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		projectID := "exampleString"
		updateProjectPath = strings.Replace(updateProjectPath, "{project_id}", projectID, 1)
		Context(`Successfully - Update a project`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateProjectPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call UpdateProject`, func() {
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
				result, response, operationErr := testService.UpdateProject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				updateProjectOptions := testService.NewUpdateProjectOptions(projectID)
				result, response, operationErr = testService.UpdateProject(updateProjectOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteProject(deleteProjectOptions *DeleteProjectOptions)`, func() {
		deleteProjectPath := "/v2/projects/{project_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		projectID := "exampleString"
		deleteProjectPath = strings.Replace(deleteProjectPath, "{project_id}", projectID, 1)
		Context(`Successfully - Delete a project`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteProjectPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(204)
			}))
			It(`Succeed to call DeleteProject`, func() {
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
				response, operationErr := testService.DeleteProject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				deleteProjectOptions := testService.NewDeleteProjectOptions(projectID)
				response, operationErr = testService.DeleteProject(deleteProjectOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteUserData(deleteUserDataOptions *DeleteUserDataOptions)`, func() {
		deleteUserDataPath := "/v2/user_data"
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
	Describe("Model constructor tests", func() {
		Context("with a sample service", func() {
			version := "1970-01-01"
			testService, _ := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
				URL:           "http://discoveryv2modelgenerator.com",
				Version:       version,
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It("should call NewCollectionDetails successfully", func() {
				name := "exampleString"
				model, err := testService.NewCollectionDetails(name)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
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
